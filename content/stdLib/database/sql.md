+++
title = "sql"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# sql

https://pkg.go.dev/database/sql@go1.20.1

​	sql包提供了一个围绕SQL(或类SQL)数据库的通用接口。

​	sql包必须与数据库驱动程序一起使用。参见[https://golang.org/s/sqldrivers](https://golang.org/s/sqldrivers)，获取驱动程序的列表。

​	不支持上下文取消的驱动程序将等到查询完成后才会返回。

​	关于用法的例子，请参见wiki页面[https://golang.org/s/sqlwiki](https://golang.org/s/sqlwiki)。

### Example (OpenDBCLI) 

```go 
package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"
)

var pool *sql.DB // 数据库连接池。

func main() {
	id := flag.Int64("id", 0, "person ID to find")
	dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	flag.Parse()

	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	if *id == 0 {
		log.Fatal("missing person ID")
	}
	var err error

   // 打开一个驱动程序通常不会尝试连接到数据库。
	pool, err = sql.Open("driver-name", *dsn)
	if err != nil {
       // 这将不是一个连接错误，而是一个DSN解析错误或其他初始化错误。
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		<-appSignal
		stop()
	}()

	Ping(ctx)

	Query(ctx, *id)
}

// Ping数据库，以验证用户提供的DSN是否有效，服务器是否可以访问。如果Ping失败，则以错误退出程序。
func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

// 在数据库中查询所要求的信息并打印出结果。
// 如果查询失败，则以错误退出程序。
func Query(ctx context.Context, id int64) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var name string
	err := pool.QueryRowContext(ctx, "select p.name from people as p where p.id = :id;", sql.Named("id", id)).Scan(&name)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}
	log.Println("name=", name)
}

```

### Example (OpenDBService) 

```go 
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// 打开驱动程序通常不会尝试连接数据库。这将不是连接错误，而是DSN解析错误或其他初始化错误。
	db, err := sql.Open("driver-name", "database=test1")
	if err != nil {
		// 打开驱动程序通常不会尝试连接数据库。
        这不会是连接错误，而是DSN解析错误或其他初始化错误。
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	s := &Service{db: db}

	http.ListenAndServe(":8080", s)
}

type Service struct {
	db *sql.DB
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := s.db
	switch r.URL.Path {
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	case "/healthz":
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()

		err := s.db.PingContext(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("db down: %v", err), http.StatusFailedDependency)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	case "/quick-action":
		// 这是一个短查询语句。使用请求上下文作为上下文超时的基础。
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		id := 5
		org := 10
		var name string
		err := db.QueryRowContext(ctx, `
select
	p.name
from
	people as p
	join organization as o on p.organization = o.id
where
	p.id = :id
	and o.id = :org
;`,
			sql.Named("id", id),
			sql.Named("org", org),
		).Scan(&name)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		io.WriteString(w, name)
		return
	case "/long-action":
		// 这是一个长查询语句。使用请求上下文作为上下文超时的基础，
        // 但给它一些时间来完成。如果客户端在查询完成之前取消了操作，
        // 则该查询也将被取消。
		ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
		defer cancel()

		var names []string
		rows, err := db.QueryContext(ctx, "select p.name from people as p where p.active = true;")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for rows.Next() {
			var name string
			err = rows.Scan(&name)
			if err != nil {
				break
			}
			names = append(names, name)
		}
		// 检查行"Close"时是否有错误。
        // 如果在单个批处理中执行了多个语句并且写入了行，那么这可能更重要。
		if closeErr := rows.Close(); closeErr != nil {
			http.Error(w, closeErr.Error(), http.StatusInternalServerError)
			return
		}

		// 检查行扫描错误。
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 检查行迭代过程中是否有错误。
		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(names)
		return
	case "/async-action":
		// 这个操作有副作用，我们希望即使客户端在HTTP请求过程中部分取消了，
        // 它也能够保留。为此，我们不使用http请求上下文作为超时的基础。
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var orderRef = "ABC123"
		tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
		_, err = tx.ExecContext(ctx, "stored_proc_name", orderRef)

		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tx.Commit()
		if err != nil {
			http.Error(w, "action in unknown state, check state before attempting again", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}

```


## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/sql.go;l=1898)

``` go 
var ErrConnDone = errors.New("sql: connection is already closed")
```

​	ErrConnDone 被任何在已经返回到连接池的连接上执行的操作返回。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/sql.go;l=441)

``` go 
var ErrNoRows = errors.New("sql: no rows in result set")
```

​	ErrNoRows 在 Scan 时由 QueryRow 不返回行时返回。在这种情况下，QueryRow 返回一个占位符 *Row 值，直到扫描时才会出现此错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/sql.go;l=2185)

``` go 
var ErrTxDone = errors.New("sql: transaction has already been committed or rolled back")
```

​	ErrTxDone 被任何在已经提交或回滚的事务上执行的操作返回。

## 函数

#### func Drivers  <- go1.4

``` go 
func Drivers() []string
```

​	Drivers函数返回已注册驱动程序名称的排序列表。

#### func Register 

``` go 
func Register(name string, driver driver.Driver)
```

​	Register函数使用提供的名称使数据库驱动程序可用。如果使用相同的名称两次调用Register或者驱动程序为nil，则会引发panic。

## 类型

### type ColumnType  <- go1.8

``` go 
type ColumnType struct {
	// contains filtered or unexported fields
	// 包含已过滤或未导出的字段
}
```

​	ColumnType结构体包含列的名称和类型。

#### (*ColumnType) DatabaseTypeName  <- go1.8

``` go 
func (ci *ColumnType) DatabaseTypeName() string
```

​	DatabaseTypeName方法返回列类型的数据库系统名称。如果返回空字符串，则不支持驱动程序类型名称。请查阅您的驱动程序文档以获取驱动程序数据类型列表。不包括长度说明符。常见类型名称包括"VARCHAR"，"TEXT"，"NVARCHAR"，"DECIMAL"，"BOOL"，"INT"和"BIGINT"。

#### (*ColumnType) DecimalSize  <- go1.8

``` go 
func (ci *ColumnType) DecimalSize() (precision, scale int64, ok bool)
```

​	DecimalSize方法返回十进制类型的比例和精度。如果不适用或不受支持，则为假。

#### (*ColumnType) Length  <- go1.8

``` go 
func (ci *ColumnType) Length() (length int64, ok bool)
```

​	Length方法返回变量长度列类型(如文本和二进制字段类型)的列类型长度。如果类型长度不受限制，则值将是math.MaxInt64(任何数据库限制仍将适用)。如果列类型不是可变长度，例如int，或者驱动程序不支持，则为false。

#### (*ColumnType) Name  <- go1.8

``` go 
func (ci *ColumnType) Name() string
```

​	Name方法返回列的名称或别名。

#### (*ColumnType) Nullable  <- go1.8

``` go 
func (ci *ColumnType) Nullable() (nullable, ok bool)
```

​	Nullable方法报告列是否可以为null。如果驱动程序不支持此属性，则ok将为false。

#### (*ColumnType) ScanType  <- go1.8

``` go 
func (ci *ColumnType) ScanType() reflect.Type
```

​	ScanType方法返回适合使用Rows.Scan进行扫描的Go类型。如果驱动程序不支持此属性，则ScanType将返回空接口的类型。

### type Conn  <- go1.9

```go 
type Conn struct {
	// contains filtered or unexported fields
	// 包含已过滤或未导出的字段
}
```

​	Conn结构体表示一个单独的数据库连接，而不是数据库连接池。除非有特定需要一个连续的单个数据库连接，否则请优先使用 DB 来运行查询。

​	在调用 Close方法以将连接返回到数据库池之前，Conn 必须调用 Close方法，可能与正在运行的查询同时进行。

​	调用 Close方法后，连接上的所有操作都将失败，并返回 ErrConnDone。

#### (*Conn) BeginTx  <- go1.9

``` go 
func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)
```

​	BeginTx方法开始一个事务。

​	提供的上下文会一直使用到事务提交或回滚。如果上下文被取消，sql 包将回滚事务。如果传递给 BeginTx 的上下文被取消，Tx.Commit 将返回一个错误。

​	提供的 TxOptions 是可选的，如果应使用默认值，则可以为 nil。如果使用了非默认的隔离级别，并且该驱动程序不支持，将返回一个错误。

#### (*Conn) Close  <- go1.9

``` go 
func (c *Conn) Close() error
```

​	Close方法将连接返回到连接池。Close方法后的所有操作都将返回 ErrConnDone。Close 方法可以与其他操作并发调用，它将阻塞，直到所有其他操作完成。在调用 Close 前，先取消任何使用的上下文，然后直接调用它可能会很有用。

#### (*Conn) ExecContext  <- go1.9

``` go 
func (c *Conn) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
```

​	ExecContext方法执行一条查询，而不返回任何行。args 用于查询中的任何占位符参数。

##### ExecContext Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	// A *DB is a pool of connections. Call Conn to reserve a connection for
	// exclusive use.
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // Return the connection to the pool.
	id := 41
	result, err := conn.ExecContext(ctx, `UPDATE balances SET balance = balance + 10 WHERE user_id = ?;`, id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected single row affected, got %d rows affected", rows)
	}
}

```

#### (*Conn) PingContext  <- go1.9

``` go 
func (c *Conn) PingContext(ctx context.Context) error
```

​	PingContext方法验证与数据库的连接是否仍然存活。

#### (*Conn) PrepareContext  <- go1.9

``` go 
func (c *Conn) PrepareContext(ctx context.Context, query string) (*Stmt, error)
```

​	PrepareContext方法为以后的查询或执行创建一个准备好的语句。可以从返回的语句中并发运行多个查询或执行。调用者在不再需要语句时必须调用语句的 Close 方法。

​	提供的上下文用于准备语句，而不是执行语句的上下文。

#### (*Conn) QueryContext  <- go1.9

``` go 
func (c *Conn) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
```

​	QueryContext方法执行一个返回行的查询，通常是SELECT。args参数用于查询中的任何占位符参数。

#### (*Conn) QueryRowContext  <- go1.9

``` go 
func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...any) *Row
```

​	QueryRowContext方法执行一个预期返回最多一行的查询。QueryRowContext总是返回一个非nil值。错误会被推迟到调用Row的Scan方法时返回。如果查询没有选择任何行，则`* Row`的Scan将返回ErrNoRows。否则，`*Row`的Scan扫描第一个选择的行并丢弃其余的行。

#### (*Conn) Raw  <- go1.13

``` go 
func (c *Conn) Raw(f func(driverConn any) error) (err error)
```

​	Raw方法执行f，暴露底层驱动程序连接。driverConn不能在f之外使用。

​	一旦f返回且err不是driver.ErrBadConn，Conn将继续可用，直到调用Conn.Close为止。

### type DB 

``` go 
type DB struct {
	// contains filtered or unexported fields
}
```

​	DB结构体是表示零个或多个底层连接池的数据库句柄。它可以被多个goroutine并发使用。

​	sql包会自动创建和释放连接；它还维护一个空闲连接的自由池。如果数据库有每个连接状态的概念，则可以在事务(Tx)或连接(Conn)中可靠地观察到此状态。一旦调用DB.Begin，返回的Tx将绑定到单个连接。一旦在事务上调用Commit或Rollback，该事务的连接就会返回到DB的空闲连接池中。池的大小可以使用SetMaxIdleConns进行控制。

#### func Open 

``` go 
func Open(driverName, dataSourceName string) (*DB, error)
```

​	Open函数通过其数据库驱动程序名称和驱动程序特定的数据源名称打开一个数据库，通常至少包含数据库名称和连接信息。

​	大多数用户将通过返回`*DB`的特定于驱动程序的连接助手函数打开数据库。Go标准库中不包含任何数据库驱动程序。请参阅https://golang.org/s/sqldrivers以获取第三方驱动程序的列表。

​	Open函数可能只验证其参数而不创建到数据库的连接。要验证数据源名称是否有效，请调用Ping。

​	返回的DB可安全地由多个goroutine并发使用，并维护其自己的空闲连接池。因此，应该只调用一次Open函数。很少需要关闭DB。

#### func OpenDB  <- go1.10

``` go 
func OpenDB(c driver.Connector) *DB
```

​	OpenDB函数使用Connector打开数据库，允许驱动程序绕过基于字符串的数据源名称。

​	大多数用户将通过特定于驱动程序的连接助手函数打开数据库，该函数返回*DB。Go标准库中不包含任何数据库驱动程序。有关第三方驱动程序的列表，请参见https://golang.org/s/sqldrivers。

​	OpenDB函数可能只是验证其参数而没有创建到数据库的连接。要验证数据源名称是否有效，请调用Ping。

​	返回的DB安全地支持多个goroutine的并发使用，并维护其自己的空闲连接池。因此，OpenDB函数应该只调用一次。很少需要关闭DB。

#### (*DB) Begin 

``` go 
func (db *DB) Begin() (*Tx, error)
```

​	Begin方法开始一个事务。默认的隔离级别取决于驱动程序。

​	Begin方法在内部使用context.Background；要指定上下文，请使用BeginTx方法。

#### (*DB) BeginTx  <- go1.8

``` go 
func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)
```

​	BeginTx方法开始一个事务。

​	提供的上下文用于直到事务提交或回滚。如果上下文被取消，则sql包将回滚事务。如果提供给BeginTx方法的上下文被取消，Tx.Commit将返回错误。

​	提供的TxOptions是可选的，如果应使用默认值，则可以为nil。如果使用驱动程序不支持的非默认隔离级别，则会返回错误。

##### BeginTx Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	id := 37
	_, execErr := tx.Exec(`UPDATE users SET status = ? WHERE id = ?`, "paid", id)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

```

#### (*DB) Close 

``` go 
func (db *DB) Close() error
```

​	Close方法关闭数据库并防止启动新查询。然后，Close方法等待在服务器上已经开始处理的所有查询完成。

​	很少需要关闭DB，因为DB句柄应该是长期存在的并在许多goroutine之间共享。

#### (*DB) Conn  <- go1.9

``` go 
func (db *DB) Conn(ctx context.Context) (*Conn, error)
```

​	Conn方法通过打开一个新连接或从连接池中返回一个现有连接来返回单个连接。 Conn 将阻塞直到返回连接或取消 ctx。在同一 Conn 上运行的查询将在同一个数据库会话中运行。

​	每个 Conn方法必须通过调用 Conn.Close 返回到数据库池中。

#### (*DB) Driver 

``` go 
func (db *DB) Driver() driver.Driver
```

​	Driver方法返回数据库的底层驱动程序。

#### (*DB) Exec 

``` go 
func (db *DB) Exec(query string, args ...any) (Result, error)
```

​	Exec方法执行一个查询，而不返回任何行。 args 是查询中任何占位符参数的参数。

​	Exec方法内部使用 context.Background。要指定上下文，请使用 ExecContext方法。

#### (*DB) ExecContext  <- go1.8

``` go 
func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
```

​	ExecContext方法执行一个查询，而不返回任何行。 args 是查询中任何占位符参数的参数。

##### ExecContext Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	id := 47
	result, err := db.ExecContext(ctx, "UPDATE balances SET balance = balance + 10 WHERE user_id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}
}

```

#### (*DB) Ping  <- go1.1

``` go 
func (db *DB) Ping() error
```

​	Ping方法验证与数据库的连接仍然存在，必要时建立连接。

​	Ping方法内部使用 context.Background。要指定上下文，请使用 PingContext方法。

#### (*DB) PingContext  <- go1.8

``` go 
func (db *DB) PingContext(ctx context.Context) error
```

​	PingContext方法验证与数据库的连接仍然存在，必要时建立连接。

##### PingContext Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
	"time"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	// Ping和PingContext可用于确定是否仍然可以与数据库服务器通信。
	//
	// 在命令行应用程序中使用Ping可以建立进一步查询是可能的；DSN是有效的。
	//
	// 在长时间运行的服务中，Ping可以是健康检查系统的一部分。
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	status := "up"
	if err := db.PingContext(ctx); err != nil {
		status = "down"
	}
	log.Println(status)
}

```

#### (*DB) Prepare 

``` go 
func (db *DB) Prepare(query string) (*Stmt, error)
```

​	Prepare方法创建一个准备好的语句，以便以后进行查询或执行操作。可以从返回的语句中并发运行多个查询或执行。调用者在语句不再需要时必须调用语句的 Close 方法。

​	Prepare方法内部使用 context.Background。要指定上下文，请使用 PrepareContext方法。

##### Prepare Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
)

var db *sql.DB

func main() {
	projects := []struct {
		mascot  string
		release int
	}{
		{"tux", 1991},
		{"duke", 1996},
		{"gopher", 2009},
		{"moby dock", 2013},
	}

	stmt, err := db.Prepare("INSERT INTO projects(id, mascot, release, category) VALUES( ?, ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for id, project := range projects {
		if _, err := stmt.Exec(id+1, project.mascot, project.release, "open source"); err != nil {
			log.Fatal(err)
		}
	}
}

```

#### (*DB) PrepareContext  <- go1.8

``` go 
func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error)
```

​	PrepareContext方法创建一个准备好的语句，以便以后进行查询或执行操作。可以从返回的语句中并发运行多个查询或执行。调用者在语句不再需要时必须调用语句的 Close 方法。

​	提供的上下文用于语句的准备，而不是语句的执行。

#### (*DB) Query 

``` go 
func (db *DB) Query(query string, args ...any) (*Rows, error)
```

​	Query方法执行返回行的查询，通常是 SELECT。args 是查询中的任何占位符参数。

​	Query方法内部使用 context.Background。要指定上下文，请使用 QueryContext方法。

##### Query Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
)

var db *sql.DB

func main() {
	age := 27
	q := `
create temp table uid (id bigint); -- Create temp table for queries.
insert into uid
select id from users where age < ?; -- Populate temp table.

-- First result set.
select
	users.id, name
from
	users
	join uid on users.id = uid.id
;

-- Second result set.
select 
	ur.user, ur.role
from
	user_roles as ur
	join uid on uid.id = ur.user
;
	`
	rows, err := db.Query(q, age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Printf("id %d name is %s\n", id, name)
	}
	if !rows.NextResultSet() {
		log.Fatalf("expected more result sets: %v", rows.Err())
	}
	var roleMap = map[int64]string{
		1: "user",
		2: "admin",
		3: "gopher",
	}
	for rows.Next() {
		var (
			id   int64
			role int64
		)
		if err := rows.Scan(&id, &role); err != nil {
			log.Fatal(err)
		}
		log.Printf("id %d has role %s\n", id, roleMap[role])
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

```

#### (*DB) QueryContext  <- go1.8

```go 
func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
```

​	QueryContext方法执行返回行的查询，通常是 SELECT。args 是查询中的任何占位符参数。

##### QueryContext Example

```go 
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	age := 27
	rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age=?", age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	names := make([]string, 0)

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		names = append(names, name)
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s are %d years old", strings.Join(names, ", "), age)
}

```

#### (*DB) QueryRow 

``` go 
func (db *DB) QueryRow(query string, args ...any) *Row
```

​	QueryRow方法执行一个预期返回最多一行的查询。QueryRow方法总是返回一个非空值。错误将被延迟直到调用 Row 的 Scan 方法。如果查询未选择任何行，则 `*Row` 的 Scan 将返回 ErrNoRows。否则，`*Row` 的 Scan 扫描第一行并丢弃其余行。

​	QueryRow方法内部使用 context.Background。要指定上下文，请使用 QueryRowContext方法。

#### (*DB) QueryRowContext  <- go1.8

``` go 
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row
```

​	QueryRowContext方法执行一个预期返回最多一行的查询。QueryRowContext 总是返回一个非空值。错误将被延迟直到调用 Row 的 Scan 方法。如果查询未选择任何行，则 `*Row` 的 Scan 将返回 ErrNoRows。否则，`*Row` 的 Scan 扫描第一行并丢弃其余行。

##### QueryRowContext Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
	"time"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	id := 123
	var username string
	var created time.Time
	err := db.QueryRowContext(ctx, "SELECT username, created_at FROM users WHERE id=?", id).Scan(&username, &created)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", id)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created on %s\n", username, created)
	}
}

```

#### (*DB) SetConnMaxIdleTime  <- go1.15

``` go 
func (db *DB) SetConnMaxIdleTime(d time.Duration)
```

​	SetConnMaxIdleTime方法设置连接的最大空闲时间。

​	过期的连接可能会在重新使用之前被懒惰地关闭。

​	如果d <= 0，则连接不会因连接的空闲时间而关闭。

#### (*DB) SetConnMaxLifetime  <- go1.6

``` go 
func (db *DB) SetConnMaxLifetime(d time.Duration)
```

​	SetConnMaxLifetime方法设置连接可以被重用的最长时间。

​	到期的连接可能会在重用之前被懒惰地关闭。

​	如果d <= 0，则不会因连接的年龄而关闭连接。

#### (*DB) SetMaxIdleConns  <- go1.1

``` go 
func (db *DB) SetMaxIdleConns(n int)
```

​	SetMaxIdleConns方法设置空闲连接池中的最大连接数。

​	如果MaxOpenConns大于0但小于新的MaxIdleConns，则新的MaxIdleConns将被减少以匹配MaxOpenConns限制。

​	如果n <= 0，则不会保留空闲连接。

​	默认的最大空闲连接数当前为2。这可能会在将来的版本中更改。

#### (*DB) SetMaxOpenConns  <- go1.2

``` go 
func (db *DB) SetMaxOpenConns(n int)
```

​	SetMaxOpenConns方法设置打开到数据库的最大连接数。

​	如果MaxIdleConns大于0且新的MaxOpenConns小于MaxIdleConns，则MaxIdleConns将被减少以匹配新的MaxOpenConns限制。

​	如果n <= 0，则没有打开连接的数量限制。默认值为0(无限制)。

#### (*DB) Stats  <- go1.5

``` go 
func (db *DB) Stats() DBStats
```

​	Stats方法返回数据库统计信息。

### type DBStats  <- go1.5

```go 
type DBStats struct {
	MaxOpenConnections int //打开到数据库的最大连接数。

	// 池状态
	OpenConnections int // 在使用和空闲中建立的连接数。
	InUse           int // 当前正在使用的连接数。
	Idle            int // 空闲连接数。

	// 计数器
	WaitCount         int64         // 等待新连接的总数。
	WaitDuration      time.Duration // 阻止等待新连接的总时间。
	MaxIdleClosed     int64 // 由于SetMaxIdleConns而关闭的连接总数。
	MaxIdleTimeClosed int64 // 由于SetConnMaxIdleTime而关闭的连接总数。
	MaxLifetimeClosed int64 // 由于SetConnMaxLifetime而关闭的连接总数。
}
```

​	DBStats包含数据库统计信息。

### type IsolationLevel  <- go1.8

``` go 
type IsolationLevel int
```

​	IsolationLevel是在TxOptions中使用的事务隔离级别。

```go 
const (
	LevelDefault IsolationLevel = iota
	LevelReadUncommitted
	LevelReadCommitted
	LevelWriteCommitted
	LevelRepeatableRead
	LevelSnapshot
	LevelSerializable
	LevelLinearizable
)
```

​	各种隔离级别，驱动程序可以在BeginTx中支持。如果驱动程序不支持给定的隔离级别，则可能返回错误。

​	参见https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels。

#### (IsolationLevel) String  <- go1.11

``` go 
func (i IsolationLevel) String() string
```

​	String 函数返回事务隔离级别的名称。

### type NamedArg  <- go1.8

```go 
type NamedArg struct {

	// Name 是参数占位符的名称。
	//
	// 如果为空，则使用参数列表中的序号。
	//
	// Name 必须省略任何符号前缀。	
	Name string

	// Value 是参数的值。
	// 它可以被赋予与查询参数相同的值类型。
	Value any
    // 包含已过滤或未导出的字段
}
```

​	NamedArg 是一个命名参数。NamedArg 值可以作为参数传递给 Query 或 Exec 并绑定到 SQL 语句中相应的命名参数。

​	为了更简洁地创建 NamedArg 值，请参考 Named 函数。

#### func Named  <- go1.8

``` go 
func Named(name string, value any) NamedArg
```

​	Named 函数提供了一种更简洁的方式来创建 NamedArg 值。

Example usage:

使用示例：

```go 
db.ExecContext(ctx, `
    delete from Invoice
    where
        TimeCreated < @end
        and TimeCreated >= @start;`,
    sql.Named("start", startTime),
    sql.Named("end", endTime),
)
```

### type NullBool 

```go 
type NullBool struct {
	Bool  bool
	Valid bool // 如果Bool不是NULL，Valid为真。
}
```

​	NullBool结构体表示可能为空的 bool 值。NullBool 实现了 Scanner 接口，因此它可以用作扫描目标，类似于 NullString。

#### (*NullBool) Scan 

``` go 
func (n *NullBool) Scan(value any) error
```

​	Scan方法实现了 Scanner 接口。

#### (NullBool) Value 

``` go 
func (n NullBool) Value() (driver.Value, error)
```

​	Value方法实现了 driver Valuer 接口。

### type NullByte  <- go1.17

```go 
type NullByte struct {
	Byte  byte
	Valid bool // 如果 Byte 不为 NULL，则 Valid 为 true
}
```

​	NullByte结构体表示一个可能为空的 byte。NullByte 实现了 Scanner 接口，因此它可以像 NullString 一样用作扫描目标。

#### (*NullByte) Scan  <- go1.17

``` go 
func (n *NullByte) Scan(value any) error
```

​	Scan方法实现了 Scanner 接口。

#### (NullByte) Value  <- go1.17

``` go 
func (n NullByte) Value() (driver.Value, error)
```

​	Value方法实现了 driver Valuer 接口。

### type NullFloat64 

```go 
type NullFloat64 struct {
	Float64 float64
	Valid   bool // 如果 Float64 不为 NULL，则 Valid 为 true
}
```

​	NullFloat64结构体表示一个可能为空的 float64。NullFloat64 实现了 Scanner 接口，因此它可以像 NullString 一样用作扫描目标。

#### (*NullFloat64) Scan 

``` go 
func (n *NullFloat64) Scan(value any) error
```

​	Scan方法实现了 Scanner 接口。

#### (NullFloat64) Value 

``` go 
func (n NullFloat64) Value() (driver.Value, error)
```

​	Value方法实现了 driver Valuer 接口。

### type NullInt16  <- go1.17

```go 
type NullInt16 struct {
	Int16 int16
	Valid bool // 如果Int16不是NULL，Valid为true
}
```

​	NullInt16结构体表示可能为null的int16。NullInt16实现Scanner接口，因此它可以用作扫描目标，类似于NullString。

#### (*NullInt16) Scan  <- go1.17

``` go 
func (n *NullInt16) Scan(value any) error
```

​	Scan方法实现Scanner接口。

#### (NullInt16) Value  <- go1.17

``` go 
func (n NullInt16) Value() (driver.Value, error)
```

​	Value方法实现driver Valuer接口。

### type NullInt32  <- go1.13

```go 
type NullInt32 struct {
	Int32 int32
	Valid bool //如果Int32不是NULL，Valid为true
}
```

​	NullInt32结构体表示可能为null的int32。NullInt32实现Scanner接口，因此它可以用作扫描目标，类似于NullString。

#### (*NullInt32) Scan  <- go1.13

``` go 
func (n *NullInt32) Scan(value any) error
```

​	Scan方法实现Scanner接口。

#### (NullInt32) Value  <- go1.13

``` go 
func (n NullInt32) Value() (driver.Value, error)
```

​	Value方法实现driver Valuer接口。

### type NullInt64 

``` go 
type NullInt64 struct {
	Int64 int64
	Valid bool // 如果Int64不是NULL，Valid为true
}
```

​	NullInt64结构体表示可能为null的int64。NullInt64实现Scanner接口，因此它可以用作扫描目标，类似于NullString。

#### (*NullInt64) Scan 

``` go 
func (n *NullInt64) Scan(value any) error
```

​	Scan方法实现Scanner接口。

#### (NullInt64) Value 

``` go 
func (n NullInt64) Value() (driver.Value, error)
```

​	Value方法实现driver Valuer接口。

### type NullString 

```go 
type NullString struct {
	String string
	Valid  bool // 如果 String 不为 NULL，则 Valid 为 true
}
```

​	NullString结构体表示可能为 null 的字符串。NullString 实现了 Scanner 接口，因此可以用作扫描目标：

```go 
var s NullString
err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&s)
...
if s.Valid {
   // use s.String
} else {
   // NULL value
}
```

#### (*NullString) Scan 

``` go 
func (ns *NullString) Scan(value any) error
```

​	Scan方法实现了 Scanner 接口。

#### (NullString) Value 

``` go 
func (ns NullString) Value() (driver.Value, error)
```

​	Value方法实现了 driver Valuer 接口。

### type NullTime  <- go1.13

```go 
type NullTime struct {
	Time  time.Time
	Valid bool // 如果 Time 不为 NULL，则 Valid 为 true
}
```

​	NullTime结构体表示可能为 null 的 time.Time。NullTime 实现了 Scanner 接口，因此可以用作扫描目标，类似于 NullString。

#### (*NullTime) Scan  <- go1.13

``` go 
func (n *NullTime) Scan(value any) error
```

​	Scan方法实现了 Scanner 接口。

#### (NullTime) Value  <- go1.13

``` go 
func (n NullTime) Value() (driver.Value, error)
```

​	Value方法实现了 driver Valuer 接口。

### type Out  <- go1.9

```go 
type Out struct {

	// Dest is a pointer to the value that will be set to the result of the
	// stored procedure's OUTPUT parameter.
    // Dest是一个指向将被设置为存储过程OUTPUT参数结果的值的指针。
	Dest any

	// In is whether the parameter is an INOUT parameter. If so, the input value to the stored
	// procedure is the dereferenced value of Dest's pointer, which is then replaced with
	// the output value.
    // In是指该参数是否为INOUT参数。如果是的话，存储过程的输入值是Dest指针的解引用值，然后用输出值替换。
	In bool
	// contains filtered or unexported fields
    //包含过滤过的或未导出的字段
}
```

​	Out结构体可用于从存储过程中检索 OUTPUT 值参数。

​	并非所有驱动程序和数据库都支持 OUTPUT 值参数。

使用示例：

```go 
var outArg string
_, err := db.ExecContext(ctx, "ProcName", sql.Named("Arg1", sql.Out{Dest: &outArg}))
```

### type RawBytes 

``` go 
type RawBytes []byte
```

​	RawBytes 是一个字节切片，它持有数据库本身拥有的内存引用。将 RawBytes扫描后，该切片仅在下一次调用 Next方法、Scan方法或 Close方法之前有效。

### type Result 

```go 
type Result interface {
	// LastInsertId 返回数据库在响应命令时生成的整数。
	// 通常这将来自插入新行时的"自增"列。
    // 不是所有数据库都支持此功能，并且此类语句的语法各不相同。
	LastInsertId() (int64, error)

	// RowsAffected 返回受更新、插入或删除影响的行数。
    // 并非所有数据库或数据库驱动程序都支持此功能。
	RowsAffected() (int64, error)
}
```

​	Result 是对执行的 SQL 命令的摘要。

### type Row 

```go 
type Row struct {
	// 包含已过滤或未导出的字段
}
```

​	Row结构体是调用 QueryRow方法选择单行时的结果。

#### (*Row) Err  <- go1.15

``` go 
func (r *Row) Err() error
```

​	Err方法提供了一种方法，使封装的包在不调用 Scan方法的情况下检查查询错误。如果在运行查询时遇到错误，Err 返回错误(如果有)。如果此错误不是 nil，则还将从 Scan 返回此错误。

#### (*Row) Scan 

``` go 
func (r *Row) Scan(dest ...any) error
```

​	Scan方法将匹配的行中的列复制到 dest 指向的值中。有关详细信息，请参见 Rows.Scan 的文档。如果有多行与查询匹配，则 Scan方法使用第一行并且忽略其余的行。如果没有行与查询匹配，则 Scan方法返回 ErrNoRows。

### type Rows 

```go 
type Rows struct {
	// contains filtered or unexported fields
	// 包含已过滤或未导出的字段
}
```

​	Rowsjgt是查询的结果。它的游标在结果集的第一行之前。使用 Next方法来从一行移到另一行。

##### Rows Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	age := 27
	rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age=?", age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s are %d years old", strings.Join(names, ", "), age)
}

```

#### (*Rows) Close 

``` go 
func (rs *Rows) Close() error
```

​	Close方法方法关闭 Rows，防止进一步枚举。如果调用 Next 并且返回 false，并且没有其他结果集，则 Rows 将自动关闭，并且检查 Err 的结果就足够了。Close方法是幂等的，不会影响 Err 的结果。

#### (*Rows) ColumnTypes  <- go1.8

``` go 
func (rs *Rows) ColumnTypes() ([]*ColumnType, error)
```

​	ColumnTypes方法返回列信息，例如列类型、长度和可空性。某些信息可能不适用于某些驱动程序。

#### (*Rows) Columns 

``` go 
func (rs *Rows) Columns() ([]string, error)
```

​	Columns方法返回列名。如果行已关闭，则 Columns方法返回错误。

#### (*Rows) Err 

``` go 
func (rs *Rows) Err() error
```

​	Err方法返回迭代过程中遇到的错误(如果有)。Err方法可以在显式或隐式 Close 后调用。

#### (*Rows) Next 

``` go 
func (rs *Rows) Next() bool
```

​	Next方法准备下一个结果行以供 Scan 方法读取。它返回 true 表示成功，false 表示没有下一个结果行或准备它时发生错误。应该使用 Err 来区分这两种情况。

​	每次调用 Scan方法，即使是第一次调用，也必须先调用 Next方法。

#### (*Rows) NextResultSet  <- go1.8

``` go 
func (rs *Rows) NextResultSet() bool
```

​	NextResultSet方法准备下一个结果集以供读取。它返回 true 表示还有其他结果集，或者 false 表示没有其他结果集，或者无法提前到其结果集。应该使用 Err 方法来区分这两种情况。

​	在调用 NextResultSet方法后，应始终在扫描之前调用 Next 方法。如果还有其他结果集，则它们可能没有结果集中的行。

#### (*Rows) Scan 

``` go 
func (rs *Rows) Scan(dest ...any) error
```

​	Scan方法将当前行的列复制到 dest 指向的值中。dest 中的值数量必须与 Rows 中的列数相同。

​	Scan方法将从数据库中读取的列转换为 sql 包提供的以下常见 Go 类型和特殊类型：

```
*string
*[]byte
*int, *int8, *int16, *int32, *int64
*uint, *uint8, *uint16, *uint32, *uint64
*bool
*float32, *float64
*interface{}
*RawBytes
*Rows (cursor value)
any type implementing Scanner (see Scanner docs)
```

​	实现 Scanner 接口的任何类型(请参见 Scanner 文档) 在最简单的情况下，如果源列的值类型是整数、布尔或字符串类型 T，而 dest 的类型是 *T，则 Scan 只需通过指针分配值。

​	Scan 还可以在字符串和数字类型之间进行转换，只要不会丢失信息即可。虽然 Scan 将从数值数据库列中读取的所有数字字符串化为 *string，但会检查是否存在数字类型的扫描溢出。例如，值为 300 的 float64 或值为 "300" 的字符串可以扫描到 uint16，但不能扫描到 uint8，尽管 float64(255) 或 "255" 可以扫描到 uint8。有一个例外，即某些 float64 数字扫描为字符串时可能会丢失信息。通常，将浮点列扫描到 *float64 中。

​	如果 dest 参数的类型为 *[]byte，则 Scan 会在该参数中保存相应数据的副本。该副本由调用方拥有，可以进行修改并无限期保留。可以通过使用类型 *RawBytes 的参数来避免复制；请参阅 RawBytes 的文档以了解其使用限制。

​	如果一个参数的类型是 *interface{}，Scan 会复制由底层驱动程序提供的值，而不进行转换。从类型 []byte 到 *interface{} 扫描时，会复制切片并且调用方拥有结果。

​	类型 time.Time 的源值可以扫描到类型 *time.Time、*interface{}、*string 或 *[]byte 的值中。转换为后两者时，使用 time.RFC3339Nano。

​	布尔类型的源值可以扫描到类型 *bool、*interface{}、*string、*[]byte 或 *RawBytes 中。

​	对于扫描到 *bool，源可以是 true、false、1、0 或可由 strconv.ParseBool 解析的字符串输入。

​	Scan 还可以将查询返回的游标(例如 "select cursor(select * from my_table) from dual")转换为 *Rows 值，该值本身可以进行扫描。如果父选择查询关闭了任何游标 *Rows，则父选择查询将关闭它。

​	如果实现 Scanner 接口的第一个参数返回错误，则该错误将包装在返回的错误中。

### type Scanner 

```go 
type Scanner interface {
	// Scan从数据库驱动程序中分配一个值。
	//
	// src值将是以下类型之一：
	//
	// int64
	// float64
	// bool
	// []byte
	// string
	// time.Time
	// nil - 对于NULL值
	//
	// 如果不能存储值而不丢失信息，则应返回错误。
	//
	// 诸如[]byte之类的引用类型仅在下一次调用Scan之前有效，不应保留。
	// 它们的底层内存由驱动程序拥有。
	// 如果需要保留，则在下一次调用Scan之前复制其值。
	Scan(src any) error
}
```

​	Scanner是由Scan使用的接口。

### type Stmt 

```go 
type Stmt struct {
	// 包含已过滤或未导出的字段
}
```

​	Stmt结构体是一个预处理语句。Stmt对于多个goroutine并发使用是安全的。

​	如果在Tx或Conn上准备了Stmt，则将永远绑定到单个底层连接。 如果Tx或Conn关闭，则Stmt将变得无法使用，所有操作都将返回错误。 如果在DB上准备了Stmt，则在DB的生命周期内将保持可用。当Stmt需要在新的底层连接上执行时， 它将自动在新连接上准备自己。

##### Stmt Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	// In normal use, create one Stmt when your process starts.
	stmt, err := db.PrepareContext(ctx, "SELECT username FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Then reuse it each time you need to issue the query.
	id := 43
	var username string
	err = stmt.QueryRowContext(ctx, id).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		log.Fatalf("no user with id %d", id)
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("username is %s\n", username)
	}
}

```

#### (*Stmt) Close 

``` go 
func (s *Stmt) Close() error
```

​	Close方法关闭该语句。

#### (*Stmt) Exec 

``` go 
func (s *Stmt) Exec(args ...any) (Result, error)
```

​	Exec方法使用给定的参数执行预处理语句，并返回概括语句效果的Result。

​	Exec方法在内部使用context.Background。要指定上下文，请使用ExecContext方法。

#### (*Stmt) ExecContext  <- go1.8

``` go 
func (s *Stmt) ExecContext(ctx context.Context, args ...any) (Result, error)
```

​	ExecContext方法使用给定的参数执行预处理语句，并返回概括语句效果的Result。

#### (*Stmt) Query 

``` go 
func (s *Stmt) Query(args ...any) (*Rows, error)
```

​	Query方法使用给定的参数执行预处理查询语句，并将查询结果作为`*Rows`返回。

​	Query方法在内部使用context.Background。要指定上下文，请使用QueryContext方法。

#### (*Stmt) QueryContext  <- go1.8

``` go 
func (s *Stmt) QueryContext(ctx context.Context, args ...any) (*Rows, error)
```

​	QueryContext方法使用给定的参数执行预处理查询语句，并将查询结果作为`*Rows`返回。

#### (*Stmt) QueryRow 

``` go 
func (s *Stmt) QueryRow(args ...any) *Row
```

​	QueryRow方法使用给定的参数执行预处理的查询语句。如果在执行语句期间出现错误，则通过对返回的`*Row调`用Scan返回该错误，该`*Row`始终非零。如果查询未选择任何行，则`*Row`的Scan将返回ErrNoRows。否则，`*Row`的Scan将扫描第一个选定的行并丢弃其余行。

使用示例：

```go 
var name string
err := nameByUseridStmt.QueryRow(id).Scan(&name)
```

​	QueryRow方法在内部使用context.Background；要指定上下文，请使用QueryRowContext方法。

#### (*Stmt) QueryRowContext  <- go1.8

``` go 
func (s *Stmt) QueryRowContext(ctx context.Context, args ...any) *Row
```

​	QueryRowContext方法使用给定的参数执行预处理的查询语句。如果在执行语句期间出现错误，则通过对返回的`*Row`调用Scan返回该错误，该`*Row`始终非零。如果查询未选择任何行，则`*Row`的Scan将返回ErrNoRows。否则，`*Row`的Scan将扫描第一个选定的行并丢弃其余行。

##### QueryRowContext Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	// In normal use, create one Stmt when your process starts.
	stmt, err := db.PrepareContext(ctx, "SELECT username FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Then reuse it each time you need to issue the query.
	id := 43
	var username string
	err = stmt.QueryRowContext(ctx, id).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		log.Fatalf("no user with id %d", id)
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("username is %s\n", username)
	}
}

```

### type Tx 

```go 
type Tx struct {
	// contains filtered or unexported fields
	// 包含已过滤或未导出的字段
}
```

​	Tx结构体是数据库事务中的一个过程。

​	事务必须以Commit方法或Rollback方法的调用结束。

​	在调用Commit方法或Rollback方法之后，所有事务上的操作都会失败并返回ErrTxDone。

​	通过调用事务的Prepare方法或Stmt方法方法准备的语句将在调用Commit方法或Rollback方法时关闭。

#### (*Tx) Commit 

``` go 
func (tx *Tx) Commit() error
```

​	Commit方法提交事务。

#### (*Tx) Exec 

``` go 
func (tx *Tx) Exec(query string, args ...any) (Result, error)
```

​	Exec方法执行不返回行的查询。例如：INSERT和UPDATE。

​	Exec方法在内部使用context.Background；要指定上下文，请使用ExecContext方法。

#### (*Tx) ExecContext  <- go1.8

``` go 
func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
```

​	ExecContext方法执行不返回行的查询。例如：INSERT和UPDATE。

##### ExecContext Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	id := 37
	_, execErr := tx.ExecContext(ctx, "UPDATE users SET status = ? WHERE id = ?", "paid", id)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update failed: %v, unable to rollback: %v\n", execErr, rollbackErr)
		}
		log.Fatalf("update failed: %v", execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

```

#### (*Tx) Prepare 

``` go 
func (tx *Tx) Prepare(query string) (*Stmt, error)
```

​	Prepare方法创建一个准备好的语句以在事务中使用。

​	返回的语句在事务中运行，并将在事务提交或回滚时关闭。

​	要在此事务上使用现有的准备好的语句，请参见Tx.Stmt。

​	Prepare方法在内部使用context.Background; 要指定上下文，请使用PrepareContext。

##### Prepare Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
)

var db *sql.DB

func main() {
	projects := []struct {
		mascot  string
		release int
	}{
		{"tux", 1991},
		{"duke", 1996},
		{"gopher", 2009},
		{"moby dock", 2013},
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare("INSERT INTO projects(id, mascot, release, category) VALUES( ?, ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for id, project := range projects {
		if _, err := stmt.Exec(id+1, project.mascot, project.release, "open source"); err != nil {
			log.Fatal(err)
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

```

#### (*Tx) PrepareContext  <- go1.8

``` go 
func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error)
```

​	PrepareContext方法创建一个准备好的语句以在事务中使用。

​	返回的语句在事务中运行，并将在事务提交或回滚时关闭。

​	要在此事务上使用现有的准备好的语句，请参见Tx.Stmt。

​	提供的上下文将用于准备上下文，而不是用于执行返回的语句。返回的语句将在事务上下文中运行。

#### (*Tx) Query 

``` go 
func (tx *Tx) Query(query string, args ...any) (*Rows, error)
```

​	Query方法执行返回行的查询，通常是 SELECT。

​	Query方法在内部使用context.Background；要指定上下文，请使用QueryContext方法。

#### (*Tx) QueryContext  <- go1.8

``` go 
func (tx *Tx) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
```

​	QueryContext方法执行返回行的查询，通常是 SELECT。

#### (*Tx) QueryRow 

```go 
func (tx *Tx) QueryRow(query string, args ...any) *Row
```

​	QueryRow方法执行预期最多返回一行的查询。QueryRow方法总是返回一个非空值。错误被延迟到调用 Row 的 Scan 方法时才会返回。如果查询未选择行，则 *Row 的 Scan 将返回 ErrNoRows。否则，*Row 的 Scan 扫描第一个选择的行并丢弃其余的行。

​	QueryRow方法在内部使用context.Background；要指定上下文，请使用QueryRowContext。

#### (*Tx) QueryRowContext  <- go1.8

```go 
func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *Row
```

​	QueryRowContext方法执行预期最多返回一行的查询。QueryRowContext方法总是返回一个非空值。错误被延迟到调用 Row 的 Scan 方法时才会返回。如果查询未选择行，则 `*Row` 的 Scan 将返回 ErrNoRows。否则，`*Row` 的 Scan 扫描第一个选择的行并丢弃其余的行。

#### (*Tx) Rollback 

```go 
func (tx *Tx) Rollback() error
```

​	Rollback方法中止事务。

##### Rollback Example

```go 
package main

import (
	"context"
	"database/sql"
	"log"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	id := 53
	_, err = tx.ExecContext(ctx, "UPDATE drivers SET status = ? WHERE id = ?;", "assigned", id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
		}
		log.Fatal(err)
	}
	_, err = tx.ExecContext(ctx, "UPDATE pickups SET driver_id = $1;", id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update failed: %v, unable to back: %v", err, rollbackErr)
		}
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

```



#### (*Tx) Stmt 

```go 
func (tx *Tx) Stmt(stmt *Stmt) *Stmt
```

​	Stmt方法从现有语句返回一个事务特定的准备好的语句。

示例：

```go 
updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
...
tx, err := db.Begin()
...
res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)
```

​	返回的语句在事务中运行，并在事务提交或回滚后关闭。

​	Stmt在内部使用context.Background；要指定上下文，请使用StmtContext。

#### (*Tx) StmtContext  <- go1.8

```go 
func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt
```

​	StmtContext方法从现有语句返回一个事务特定的准备语句。

Example:

```go 
updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
...
tx, err := db.Begin()
...
res, err := tx.StmtContext(ctx, updateMoney).Exec(123.45, 98293203)
```

​	提供的上下文用于准备语句，而不是执行语句。

​	返回的语句在事务中运行，并在事务提交或回滚后关闭。

### type TxOptions  <- go1.8

```go 
type TxOptions struct {
	// Isolation是事务隔离级别。
	// 如果为零，则使用驱动程序或数据库的默认级别。
	Isolation IsolationLevel
	ReadOnly  bool
}
```

​	TxOptions保存在DB.BeginTx中使用的事务选项。

