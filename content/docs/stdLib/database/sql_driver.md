+++
title = "sql/driver"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# driver

[https://pkg.go.dev/database/sql/driver@go1.20.1](https://pkg.go.dev/database/sql/driver@go1.20.1)

​	driver包定义了数据库驱动需要实现的接口，供sql包使用。

​	大多数代码应该使用 sql 包。

​	驱动程序接口随着时间的推移而发展。驱动程序应该实现 Connector 和 DriverContext 接口。Connector.Connect 和 Driver.Open 方法不应返回 ErrBadConn。如果连接已经处于无效状态(例如已关闭)，则 ErrBadConn 只能从 Validator、SessionResetter 或查询方法中返回。

​	所有 Conn 实现都应该实现以下接口：Pinger、SessionResetter 和 Validator。

​	如果支持命名参数或上下文，则驱动程序的 Conn 应实现：ExecerContext、QueryerContext、ConnPrepareContext 和 ConnBeginTx。

​	为支持自定义数据类型，实现 NamedValueChecker。NamedValueChecker 还允许查询通过返回 ErrRemoveArgument 接受每个查询选项作为参数。

​	如果支持多个结果集，则 Rows 应实现 RowsNextResultSet。如果驱动程序知道返回结果中存在的类型，它应该实现以下接口：RowsColumnTypeScanType、RowsColumnTypeDatabaseTypeName、RowsColumnTypeLength、RowsColumnTypeNullable 和 RowsColumnTypePrecisionScale。给定行值也可以返回 Rows 类型，它可能表示数据库游标值。

​	在将连接归还给连接池之前，如果实现了 IsValid，则会调用 IsValid。在将连接用于另一个查询之前，如果实现了 ResetSession，则会调用 ResetSession。如果连接从未返回给连接池，而是立即重用，则在重用之前调用 ResetSession，但不调用 IsValid。

## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=55)

``` go 
var Bool boolType
```

​	Bool是ValueConverter，用于将输入值转换为布尔值。

转换规则如下：

- 布尔值不变 
- 对于整数类型，1为真，0为假，其他整数为错误 
- 对于字符串和[]byte，与strconv.ParseBool相同的规则 
- 所有其他类型都是一个错误

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=208)

``` go 
var DefaultParameterConverter defaultConverter
```

​	DefaultParameterConverter是ValueConverter的默认实现，当Stmt未实现ColumnConverter时使用。

​	DefaultParameterConverter如果IsValue(arg)则直接返回其参数。否则，如果参数实现了Valuer，则使用其Value方法返回一个Value。作为后备，提供的参数的基础类型用于将其转换为Value：基础整数类型转换为int64，浮点数转换为float64，bool，string和[]byte转换为自身。如果参数是nil指针，则ConvertValue返回nil Value。如果参数是非nil指针，则对其进行取消引用，并递归调用ConvertValue。其他类型是一个错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=162)

``` go 
var ErrBadConn = errors.New("driver: bad connection")
```

​	如果驱动程序处于错误状态(例如服务器早期关闭连接)，则驱动程序应返回ErrBadConn，以向sql包发信号。sql包应该在新连接上重试。

​	为了防止重复操作，如果数据库服务器可能执行操作，则不应返回ErrBadConn。即使服务器发送错误，您也不应返回ErrBadConn。

​	使用errors.Is检查错误。一个错误可能包装ErrBadConn或实现Is(error) bool方法。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=385)

``` go 
var ErrRemoveArgument = errors.New("driver: remove argument from query")
```

​	NamedValueChecker可能会返回ErrRemoveArgument，以指示sql包不将参数传递给驱动程序查询接口。在接受特定于查询的选项或不是SQL查询参数的结构时返回。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=148)

``` go 
var ErrSkip = errors.New("driver: skip fast-path; continue as if unimplemented")
```

​	某些可选接口的方法可能返回ErrSkip，以表示在运行时快速路径不可用，sql包应继续，就像未实现可选接口一样。仅在明确文档中记录了ErrSkip。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=102)

``` go 
var Int32 int32Type
```

​	Int32是ValueConverter，用于将输入值转换为int64，并考虑int32值的限制。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=540)

``` go 
var ResultNoRows noRows
```

​	ResultNoRows是预定义结果，驱动程序在DDL命令(如CREATE TABLE)成功时返回。它对于LastInsertId和RowsAffected都返回错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=137)

``` go 
var String stringType
```

​	String是ValueConverter，用于将其输入转换为字符串。如果该值已经是字符串或[]byte，则不变。如果值是其他类型，则使用fmt.Sprintf("%v"，v)进行字符串转换。

## 函数

#### func [IsScanValue](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=191) 

``` go 
func IsScanValue(v any) bool
```

​	IsScanValue函数等价于 IsValue。它存在是为了兼容性。

#### func [IsValue](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=176) 

``` go 
func IsValue(v any) bool
```

​	IsValue函数报告v是否为有效的 Value 参数类型。

## 类型

### type [Conn](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=233) 

``` go 
type Conn interface {
	// Prepare 返回一个准备好的语句，与此连接绑定。
	Prepare(query string) (Stmt, error)

	// Close 使任何当前的准备语句和事务无效，标记此连接不再使用。
	//
	// 因为sql包维护一个空闲连接池，并且仅在空闲连接过多时调用Close，
    // 所以驱动程序不需要执行自己的连接缓存。
	//
	// 驱动程序必须确保Close调用所做的所有网络调用不会无限期地阻塞(例如，应用超时)。
	Close() error

    // Begin 启动并返回一个新的事务。
    //
    // 已弃用：驱动程序应该实现 ConnBeginTx 代替(或同时实现)。
	Begin() (Tx, error)
}
```

​	Conn 是到数据库的连接。它不会被多个goroutine同时使用。

​	Conn 被认为是有状态的。

### type [ConnBeginTx](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=279)  <- go1.8

``` go 
type ConnBeginTx interface {
    // BeginTx 启动并返回一个新的事务。
    // 如果用户取消了上下文，sql包将在丢弃和关闭连接之前调用Tx.Rollback。
    //
    // 这必须检查opts.Isolation以确定是否有设置隔离级别。如果驱动程序不支持非默认级别，并且已设置一个级别，
    // 或者存在一个不支持的非默认隔离级别，则必须返回一个错误。
    //
    // 这还必须检查opts.ReadOnly以确定只读值是否为true，以设置只读事务属性(如果受支持)，
    // 或者如果不支持，则返回错误。
	BeginTx(ctx context.Context, opts TxOptions) (Tx, error)
}
```

​	ConnBeginTx 使用上下文和 TxOptions 增强了 Conn 接口。

### type [ConnPrepareContext](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=257)  <- go1.8

``` go 
type ConnPrepareContext interface {
    // PrepareContext 返回一个准备好的语句，与此连接绑定。
    // context 是用于准备语句的，它不能在语句本身中存储上下文。
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}
```

​	ConnPrepareContext 通过上下文增强 Conn 接口。

### type [Connector](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=121)  <- go1.10

``` go 
type Connector interface {
    // Connect 返回到数据库的连接。
    // Connect 可以返回缓存的连接(先前关闭的连接)，但这样做是不必要的；
    // sql 包维护一个空闲连接池，以便有效地重复使用。
    //
    // 提供的 context.Context 仅用于拨号目的
    //(请参阅 net.DialContext)，不应存储或用于其他目的。
    // 在拨号时仍应使用默认超时，因为连接池可以异步地调用 Connect 到任何查询。
    //
    // 返回的连接只由一个 goroutine 使用。
	Connect(context.Context) (Conn, error)

    // Driver 返回 Connector 的基础 Driver，
    // 主要是为了与 sql.DB 上的 Driver 方法保持兼容性。
	Driver() Driver
}
```

​	Connector 表示一个具有固定配置的驱动程序，并可以创建任意数量的等效 Conns 供多个 goroutine 使用。

​	可以将 Connector 传递给 sql.OpenDB，以允许驱动程序实现其自己的 sql.DB 构造函数，或者由 DriverContext 的 OpenConnector 方法返回，以允许驱动程序访问上下文并避免重复解析驱动程序配置。

​	如果 Connector 实现了 io.Closer，则 sql 包的 DB.Close 方法将调用 Close 并返回错误(如果有)。

### type [Driver](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=84) 

``` go 
type Driver interface {
    // Open 返回到数据库的新连接。
    // 名称是驱动程序特定格式的字符串。
    //
    // Open 可以返回缓存的连接(先前关闭的连接)，但这样做是不必要的；
    // sql 包维护一个空闲连接池，以便有效地重复使用。
    //
    // 返回的连接只由一个 goroutine 使用。
	Open(name string) (Conn, error)
}
```

​	Driver 是必须由数据库驱动程序实现的接口。

​	数据库驱动程序可以实现 DriverContext，以便访问上下文并仅解析一次名称以获得连接池，而不是每个连接都解析一次。

### type [DriverContext](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=103)  <- go1.10

``` go 
type DriverContext interface {
	// OpenConnector must parse the name in the same format that Driver.Open
	// parses the name parameter.
    // OpenConnector必须以Driver.Open解析名称参数的相同格式解析名称。
	OpenConnector(name string) (Connector, error)
}
```

​	如果 Driver 实现了 DriverContext 接口，那么 sql.DB 将调用 OpenConnector 来获取 Connector，并调用该 Connector 的 Connect 方法来获取每个所需的连接，而不是为每个连接调用 Driver 的 Open 方法。这个两步的过程允许驱动程序只解析一次名称，并且提供对每个 Conn 的上下文访问。

### type [ExecerContext](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=198)  <- go1.8

``` go 
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args []NamedValue) (Result, error)
}
```

​	ExecerContext 是一个可选接口，可以由 Conn 实现。

​	如果 Conn 没有实现 ExecerContext，则 sql 包的 DB.Exec 将退回到 Execer；如果 Conn 也没有实现 Execer，则 DB.Exec 将首先准备查询，执行语句，然后关闭语句。

​	ExecContext 可能返回 ErrSkip。

​	ExecContext 必须遵守上下文超时，并在上下文取消时返回。

### type [IsolationLevel](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=268)  <- go1.8

``` go 
type IsolationLevel int
```

​	IsolationLevel 是存储在 TxOptions 中的事务隔离级别。

​	这个类型应该被认为与 sql.IsolationLevel 相同，以及其上定义的任何值。

### type [NamedValue](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=64)  <- go1.8

``` go 
type NamedValue struct {
    // 如果 Name 不为空，应该用于参数标识符，而不是位置。
    //
    // Name 不会有符号前缀。
	Name string

	// 参数的序号位置，从一开始计数，并始终设置。
	Ordinal int	

    // Value是参数的值。
	Value Value
}
```

​	NamedValue 持有值的名称和值。

### type [NamedValueChecker](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=402)  <- go1.9

``` go 
type NamedValueChecker interface {
    // CheckNamedValue 在将参数传递给驱动程序之前调用，
    // 并在任何 ColumnConverter 的位置调用。
    // CheckNamedValue 必须做类型验证和转换，以适合驱动程序。
	CheckNamedValue(*NamedValue) error
}
```

​	NamedValueChecker 可以由 Conn 或 Stmt 可选实现。它为驱动程序提供了更多的控制权，以处理超出默认的 Value 类型允许的 Go 和数据库类型。

​	sql包会按照以下顺序检查值检查器，并在找到第一个匹配项时停止：Stmt.NamedValueChecker、Conn.NamedValueChecker、Stmt.ColumnConverter和DefaultParameterConverter。

​	如果CheckNamedValue返回ErrRemoveArgument，则NamedValue不会包含在最终查询参数中。这可用于向查询本身传递特殊选项。

​	如果返回ErrSkip，则对于该参数将使用列转换器错误检查路径。驱动程序可能会在耗尽其自己的特殊情况后返回ErrSkip。

### type [NotNull](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=164) 

``` go 
type NotNull struct {
	Converter ValueConverter
}
```

​	NotNull结构体是一种类型，通过禁止nil值但否则委托给另一个ValueConverter来实现ValueConverter。

#### (NotNull) [ConvertValue](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=168) 

``` go 
func (n NotNull) ConvertValue(v any) (Value, error)
```

### type [Null](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=151) 

``` go 
type Null struct {
	Converter ValueConverter
}
```

​	Null结构体是一种类型，通过允许nil值但否则委托给另一个ValueConverter来实现ValueConverter。

#### (Null) [ConvertValue](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=155) 

``` go 
func (n Null) ConvertValue(v any) (Value, error)
```

### type [Pinger](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=171)  <- go1.8

``` go 
type Pinger interface {
	Ping(ctx context.Context) error
}
```

​	Pinger是Conn可能实现的一个可选接口。

​	如果Conn未实现Pinger，则sql包的DB.Ping和DB.PingContext将检查是否至少有一个Conn可用。

​	如果Conn.Ping返回ErrBadConn，则DB.Ping和DB.PingContext将从池中删除Conn。

### type [QueryerContext](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=225)  <- go1.8

``` go 
type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args []NamedValue) (Rows, error)
}
```

​	QueryerContext是Conn可能实现的一个可选接口。

​	如果Conn未实现QueryerContext，则sql包的DB.Query将回退到Queryer；如果Conn也未实现Queryer，则DB.Query将首先准备查询，执行语句，然后关闭语句。

​	QueryContext可能会返回ErrSkip。

​	QueryContext必须遵守上下文超时并在取消上下文时返回。

### type [Result](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=316) 

``` go 
type Result interface {
	// LastInsertId 返回数据库生成的自增 ID，
	// 比如在插入一个带有主键的表时。
	LastInsertId() (int64, error)

	// RowsAffected 返回由查询所影响的行数。
	RowsAffected() (int64, error)
}
```

​	Result接口表示一个查询操作的结果。

### type [Rows](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=423) 

``` go 
type Rows interface {
	// Columns 返回列的名称。
    // 结果集中列的数量从切片的长度中推断出来。如果特定列的名称未知，则该项应返回空字符串。
	Columns() []string

	// Close 关闭行迭代器。
	Close() error

    // Next 用来将下一行数据填充到提供的切片中。
    // 提供的切片将与 Columns() 的宽度相同。
    //
    // Next 应在没有更多行时返回 io.EOF。
    //
    // 在 Next 之外不应写入 dest。
    // 在关闭 Rows 时要小心，以免修改 dest 中保存的缓冲区。
	Next(dest []Value) error
}
```

​	Rows接口是一个已执行查询的结果集迭代器。

### type [RowsAffected](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=525) 

``` go 
type RowsAffected int64
```

​	RowsAffected 表示执行 INSERT 或 UPDATE 操作所影响的行数，实现了 Result 接口。

#### (RowsAffected) [LastInsertId](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=529) 

``` go 
func (RowsAffected) LastInsertId() (int64, error)
```

#### (RowsAffected) [RowsAffected](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=533) 

``` go 
func (v RowsAffected) RowsAffected() (int64, error)
```

### type [RowsColumnTypeDatabaseTypeName](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=474)  <- go1.8

``` go 
type RowsColumnTypeDatabaseTypeName interface {
	Rows
	ColumnTypeDatabaseTypeName(index int) string
}
```

​	RowsColumnTypeDatabaseTypeName接口可以被 Rows 实现。它应该返回数据库系统类型名称，但不包括长度信息。类型名称应大写。以下是各种类型的返回示例："VARCHAR"、"NVARCHAR"、"VARCHAR2"、"CHAR"、"TEXT"、"DECIMAL"、"SMALLINT"、"INT"、"BIGINT"、"BOOL"、"[]BIGINT"、"JSONB"、"XML"、"TIMESTAMP"。

### type [RowsColumnTypeLength](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=491)  <- go1.8

``` go 
type RowsColumnTypeLength interface {
	Rows
	ColumnTypeLength(index int) (length int64, ok bool)
}
```

​	RowsColumnTypeLength接口可以被 Rows 实现。如果列是变长类型，则它应返回列类型的长度。如果列不是变长类型，则应该返回 false。如果长度除系统限制外不受限制，则应返回 math.MaxInt64。以下是各种类型的返回示例：

```
TEXT          (math.MaxInt64, true)
varchar(10)   (10, true)
nvarchar(10)  (10, true)
decimal       (0, false)
int           (0, false)
bytea(30)     (30, true)
```

### type [RowsColumnTypeNullable](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=500)  <- go1.8

``` go 
type RowsColumnTypeNullable interface {
	Rows
	ColumnTypeNullable(index int) (nullable, ok bool)
}
```

​	RowsColumnTypeNullable接口是由Rows实现的可选接口。如果已知该列可以为null，则可为空值应为true；如果已知该列不可为空，则为空值为false。如果列的可空性未知，则ok应为false。

### type [RowsColumnTypePrecisionScale](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=512)  <- go1.8

``` go 
type RowsColumnTypePrecisionScale interface {
	Rows
	ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool)
}
```

​	RowsColumnTypePrecisionScale接口是由Rows实现的可选接口。它应返回十进制类型的精度和刻度。如果不适用，则应将ok设置为false。以下是各种类型的返回值示例：

```
decimal(38, 4)    (38, 4, true)
int               (0, 0, false)
decimal           (math.MaxInt64, math.MaxInt64, true)
```

### type [RowsColumnTypeScanType](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=464)  <- go1.8

``` go 
type RowsColumnTypeScanType interface {
	Rows
	ColumnTypeScanType(index int) reflect.Type
}
```

​	RowsColumnTypeScanType接口是由Rows实现的可选接口。它应返回可以用于扫描类型的值类型。例如，对于数据库列类型"bigint"，这应返回"reflect.TypeOf(int64(0))"。

### type [RowsNextResultSet](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=447)  <- go1.8

``` go 
type RowsNextResultSet interface {
	Rows

	// HasNextResultSet在当前结果集结束时调用，
    // 报告当前结果集之后是否还有另一个结果集。
	HasNextResultSet() bool

	// NextResultSet将驱动程序推进到下一个结果集，
    // 即使当前结果集还有剩余行。
    //
    // 当没有更多的结果集时，NextResultSet应该返回io.EOF。
	NextResultSet() error
}
```

​	RowsNextResultSet接口通过提供一种信号方式来扩展Rows接口，以使驱动程序前进到下一个结果集。

### type [SessionResetter](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=297)  <- go1.10

``` go 
type SessionResetter interface {
	// ResetSession在连接之前执行查询时调用，
    // 如果连接之前已被使用。
    // 如果驱动程序返回ErrBadConn，则连接将被丢弃。
	ResetSession(ctx context.Context) error
}
```

​	SessionResetter接口可以由Conn实现，以允许驱动程序重置与连接相关的会话状态并发出坏连接信号。

### type [Stmt](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=329) 

``` go 
type Stmt interface {
	// Close关闭预处理语句。
    //
    // 从 Go 1.1 开始，如果预处理语句正在被任何查询使用，
    // 则不会关闭预处理语句。
    //
    // 驱动程序必须确保 Close 所做的所有网络调用不会无限期地阻塞(例如，应用超时)。
	Close() error

	// NumInput 返回占位符参数的数量。
    //
    // 如果 NumInput 返回 >= 0，
    // 则 sql 包将检查调用者的参数计数，
    // 并在调用预处理语句的 Exec 或 Query 方法之前向调用者返回错误。
    //
    // 如果驱动程序不知道其占位符的数量，
    // 则 NumInput 也可以返回-1。在这种情况下，
    // sql 包将不会检查 Exec 或 Query 参数计数。
	NumInput() int

	// Exec 执行不返回行的查询，例如 INSERT 或 UPDATE。
	//
	// 已弃用：驱动程序应该实现 StmtExecContext(或另外实现)。
	Exec(args []Value) (Result, error)

	// Query 执行可能返回行的查询，例如 SELECT。
	//
	// 已弃用：驱动程序应该实现 StmtQueryContext(或另外实现)。
	Query(args []Value) (Rows, error)
}
```

​	Stmt接口是预处理语句。它绑定到Conn，不能被多个goroutine同时使用。

### type [StmtExecContext](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=364)  <- go1.8

``` go 
type StmtExecContext interface {
	// ExecContext方法执行不返回行的查询，例如INSERT或UPDATE。
	// ExecContext方法必须遵守上下文超时并在取消时返回。
	ExecContext(ctx context.Context, args []NamedValue) (Result, error)
}
```

​	StmtExecContext接口通过提供带有上下文的Exec来增强Stmt接口。

### type [StmtQueryContext](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=373)  <- go1.8

``` go 
type StmtQueryContext interface {
	// QueryContext方法执行可能返回行的查询，例如SELECT。
	// QueryContext方法必须遵守上下文超时并在取消时返回。
	QueryContext(ctx context.Context, args []NamedValue) (Rows, error)
}
```

​	StmtQueryContext接口通过提供带有上下文的Query来增强Stmt接口。

### type [Tx](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=518) 

``` go 
type Tx interface {
	Commit() error
	Rollback() error
}
```

​	Tx 是事务。

### type [TxOptions](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=273)  <- go1.8

``` go 
type TxOptions struct {
	Isolation IsolationLevel
	ReadOnly  bool
}
```

​	TxOptions结构体存储事务的选项。

​	此类型应被认为与 sql.TxOptions 相同。

### type [Validator](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=309)  <- go1.15

``` go 
type Validator interface {
	// IsValid方法在将连接放入连接池之前调用。
    // 如果返回false，则连接将被丢弃。
	IsValid() bool
}
```

​	Validator接口可以被 Conn 实现，以允许驱动程序表明连接是否有效或是否应丢弃。

​	如果实现了 Validator，则驱动程序可能会从查询中返回基础错误，即使连接应该由连接池丢弃。

### type [Value](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=61) 

``` go 
type Value any
```

​	Value类型是驱动程序必须能够处理的值。它可以是 nil、由数据库驱动程序的 NamedValueChecker 接口处理的类型，或是以下类型的实例

```
int64
float64
bool
[]byte
string
time.Time
```

​	如果驱动程序支持游标，则返回的 Value 还可以在此包中实现 Rows 接口。例如，当用户选择类似于"select cursor(select * from my_table) from dual"这样的光标时。如果从选择中的 Rows 被关闭，则光标 Rows 也将被关闭。

### type [ValueConverter](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=30) 

``` go 
type ValueConverter interface {
	// ConvertValue方法将一个值转换为driver.Value。
	ConvertValue(v any) (Value, error)
}
```

ValueConverter接口是提供 ConvertValue 方法的接口。

驱动程序包提供了 ValueConverter 的各种实现，以提供驱动程序之间的一致的转换实现。ValueConverter 有以下几个用途：

- 将 sql 包提供的 Value 类型转换为数据库表的特定列类型，并确保其适合，例如确保特定 int64 适合于表的 uint16 列。
- 将从数据库中给出的值转换为驱动程序 Value 类型之一。
- 由 sql 包，在扫描中将驱动程序的 Value 类型转换为用户的类型。

### type [Valuer](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=39) 

``` go 
type Valuer interface {
    // Value方法返回一个driver.Value。
    // Value方法不得panic。
	Value() (Value, error)
}
```

​	Valuer接口是提供 Value 方法的接口。

​	实现 Valuer 接口的类型能够将自身转换为驱动程序 Value。