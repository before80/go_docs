+++
title = "连接到数据库"
date = 2023-10-28T14:24:49+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/connecting_to_the_database.html](https://gorm.io/docs/connecting_to_the_database.html)

GORM officially supports the databases MySQL, PostgreSQL, SQLite, SQL Server, and TiDB

​	GORM官方支持的数据库有`MySQL`、`PostgreSQL`、`SQLite`、`SQL Server`和`TiDB`。

## MySQL

```go
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详细信息
  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
```

> **NOTE:**
>
> To handle `time.Time` correctly, you need to include `parseTime` as a parameter. ([more parameters](https://github.com/go-sql-driver/mysql#parameters))
> To fully support UTF-8 encoding, you need to change `charset=utf8` to `charset=utf8mb4`. See [this 	2.article](https://mathiasbynens.be/notes/mysql-utf8mb4) for a detailed explanation
>
> **注意：** 要正确处理`time.Time`，需要将`parseTime`作为参数包含在内。（[更多参数]({{< ref "/thirdPkg/go_sql_driver_mysql#parameters">}})) 要完全支持UTF-8编码，需要将`charset=utf8`更改为`charset=utf8mb4`。参见这篇文章以获得详细解释

MySQL Driver provides a [few advanced configurations](https://github.com/go-gorm/mysql) which can be used during initialization, for example:

​	MySQL驱动程序提供了[一些高级配置](https://github.com/go-gorm/mysql)，可以在初始化时使用，例如：

```go
db, err := gorm.Open(mysql.New(mysql.Config{
  DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // 数据源名称 data source name
  DefaultStringSize: 256, // 字符串字段的默认大小 default size for string fields
  DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持 precision, which not supported before MySQL 5.6
  DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引 drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
  DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列 `change` when rename column, rename column not supported before MySQL 8, MariaDB
  SkipInitializeWithVersion: false, //  根据当前MySQL版本自动配置 auto configure based on currently MySQL version
}), &gorm.Config{})
```

### 自定义驱动 Customize Driver

GORM allows to customize the MySQL driver with the `DriverName` option, for example:

​	GORM允许使用`DriverName`选项自定义MySQL驱动程序，例如：

```go
import (
  _ "example.com/my_mysql_driver"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

db, err := gorm.Open(mysql.New(mysql.Config{
  DriverName: "my_mysql_driver",
  DSN: "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
}), &gorm.Config{})
```

### 现有数据库连接 Existing database connection

GORM allows to initialize `*gorm.DB` with an existing database connection

​	GORM允许使用现有的数据库连接初始化`*gorm.DB`

```go
import (
  "database/sql"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

sqlDB, err := sql.Open("mysql", "mydb_dsn")
gormDB, err := gorm.Open(mysql.New(mysql.Config{
  Conn: sqlDB,
}), &gorm.Config{})
```

## PostgreSQL

```go
import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

We are using [pgx](https://github.com/jackc/pgx) as postgres’s database/sql driver, it enables prepared statement cache by default, to disable it:

​	我们使用[pgx](https://github.com/jackc/pgx)作为PostgreSQL的 database/sql 驱动程序，默认启用了预处理语句缓存，要禁用它：

```go
// https://github.com/go-gorm/postgres
db, err := gorm.Open(postgres.New(postgres.Config{
  DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
  PreferSimpleProtocol: true, // 禁用隐式预处理语句使用 disables implicit prepared statement usage
}), &gorm.Config{})
```

### 自定义驱动 Customize Driver

GORM allows to customize the PostgreSQL driver with the `DriverName` option, for example:

​	GORM允许使用`DriverName`选项自定义PostgreSQL驱动程序，例如：

```go
import (
  _ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
  "gorm.io/gorm"
)

db, err := gorm.Open(postgres.New(postgres.Config{
  DriverName: "cloudsqlpostgres",
  DSN: "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable",
})
```

### 现有数据库连接 Existing database connection

GORM allows to initialize `*gorm.DB` with an existing database connection

​	GORM允许使用现有数据库连接初始化`*gorm.DB`

```go
import (
  "database/sql"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

sqlDB, err := sql.Open("pgx", "mydb_dsn")
gormDB, err := gorm.Open(postgres.New(postgres.Config{
  Conn: sqlDB,
}), &gorm.Config{})
```

## SQLite

```go
import (
  "gorm.io/driver/sqlite" // Sqlite driver based on CGO
  // "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
  "gorm.io/gorm"
)

// github.com/mattn/go-sqlite3
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
```

> **NOTE:** You can also use `file::memory:?cache=shared` instead of a path to a file. This will tell SQLite to use a temporary database in system memory. (See [SQLite docs](https://www.sqlite.org/inmemorydb.html) for this)
>
> **注意：** 你也可以使用 `file::memory:?cache=shared` 来代替一个文件路径。这将告诉 SQLite 在系统内存中使用一个临时数据库。（关于这个，请参阅 [SQLite 文档](https://www.sqlite.org/inmemorydb.html)）

## SQL Server

```go
import (
  "gorm.io/driver/sqlserver"
  "gorm.io/gorm"
)

// github.com/denisenkom/go-mssqldb
dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
```

## TiDB

TiDB is compatible with MySQL protocol. You can follow the [MySQL](https://gorm.io/docs/connecting_to_the_database.html#mysql) part to create a connection to TiDB.

​	TiDB 兼容 MySQL 协议。你可以按照 [MySQL](https://gorm.io/docs/connecting_to_the_database.html#mysql) 部分创建一个连接到 TiDB 的连接。

There are some points noteworthy for TiDB:

​	对于 TiDB 的一些值得注意的点：

- You can use `gorm:"primaryKey;default:auto_random()"` tag to use [`AUTO_RANDOM`](https://docs.pingcap.com/tidb/stable/auto-random) feature for TiDB.
- 你可以使用 `gorm:"primaryKey;default:auto_random()"` 标签来使用 [AUTO_RANDOM](https://docs.pingcap.com/tidb/stable/auto-random) 功能。
- TiDB supported [`SAVEPOINT`](https://docs.pingcap.com/tidb/stable/sql-statement-savepoint) from `v6.2.0`, please notice the version of TiDB when you use this feature.
- TiDB 从 `v6.2.0` 开始支持 [SAVEPOINT](https://docs.pingcap.com/tidb/stable/sql-statement-savepoint)，请在使用此功能时注意 TiDB 的版本。
- TiDB supported [`FOREIGN KEY`](https://docs.pingcap.com/tidb/dev/foreign-key) from `v6.6.0`, please notice the version of TiDB when you use this feature.
- TiDB 从 `v6.6.0` 开始支持 [FOREIGN KEY](https://docs.pingcap.com/tidb/dev/foreign-key)，请在使用此功能时注意 TiDB 的版本。

```go
import (
  "fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

type Product struct {
  ID    uint `gorm:"primaryKey;default:auto_random()"`
  Code  string
  Price uint
}

func main() {
  db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:4000)/test"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&Product{})

  insertProduct := &Product{Code: "D42", Price: 100}

  db.Create(insertProduct)
  fmt.Printf("insert ID: %d, Code: %s, Price: %d\n",
    insertProduct.ID, insertProduct.Code, insertProduct.Price)

  readProduct := &Product{}
  db.First(&readProduct, "code = ?", "D42") // find product with code D42

  fmt.Printf("read ID: %d, Code: %s, Price: %d\n",
    readProduct.ID, readProduct.Code, readProduct.Price)
}
```

## Clickhouse

> 原文：[https://github.com/go-gorm/clickhouse](https://github.com/go-gorm/clickhouse)

```go
import (
  "gorm.io/driver/clickhouse"
  "gorm.io/gorm"
)

func main() {
  dsn := "tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20"
  db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

  // 自动迁移 Auto Migrate
  db.AutoMigrate(&User{})
  // 设置表选项 Set table options
  db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&User{})

  // 插入 Insert
  db.Create(&user)

  // 查询 Select
  db.Find(&user, "id = ?", 10)

  // 批量插入 Batch Insert
  var users = []User{user1, user2, user3}
  db.Create(&users)
  // ...
}
```

## 连接池 Connection Pool

GORM using [database/sql](https://pkg.go.dev/database/sql) to maintain connection pool

​	GORM 使用 [database/sql]({{< ref "/stdLib/database/sql">}}) 来维护连接池

```go
sqlDB, err := db.DB()

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
// SetMaxIdleConns 设置空闲连接池中的最大连接数。
sqlDB.SetMaxIdleConns(10)

// SetMaxOpenConns sets the maximum number of open connections to the database.
// SetMaxOpenConns 设置到数据库的打开连接的最大数量。
sqlDB.SetMaxOpenConns(100)

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
// SetMaxOpenConns 设置到数据库的打开连接的最大数量。
sqlDB.SetConnMaxLifetime(time.Hour)
```

Refer [Generic Interface](https://gorm.io/docs/generic_interface.html) for details

​	详细信息请参考 [通用接口]({{< ref "/gorm/Tutorials/genericDatabaseInterface">}})

## 不支持的数据库 Unsupported Databases

Some databases may be compatible with the `mysql` or `postgres` dialect, in which case you could just use the dialect for those databases.

​	一些数据库可能与`mysql`或`postgres`方言兼容，在这种情况下，您只需为这些数据库使用该方言。

For others, [you are encouraged to make a driver, pull request welcome!](https://gorm.io/docs/write_driver.html)

​	对于其他数据库，[鼓励您制作驱动程序，欢迎拉取请求！]({{< ref "/gorm/AdancedTopics/writeDriver">}})