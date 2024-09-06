+++
title = "sql/driver"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/database/sql/driver@go1.23.0](https://pkg.go.dev/database/sql/driver@go1.23.0)

Package driver defines interfaces to be implemented by database drivers as used by package sql.

​	`driver`包定义了数据库驱动需要实现的接口，供`sql`包使用。

Most code should use package sql.

​	大多数代码应该使用 `sql` 包。

The driver interface has evolved over time. Drivers should implement Connector and DriverContext interfaces. The Connector.Connect and Driver.Open methods should never return ErrBadConn. ErrBadConn should only be returned from Validator, SessionResetter, or a query method if the connection is already in an invalid (e.g. closed) state.

​	随着时间的推移，驱动程序接口已经发生了演变。驱动程序应实现Connector和DriverContext接口。 Connector.Connect和Driver.Open方法永远不应该返回ErrBadConn。如果连接已经处于无效状态(例如已关闭)，则 ErrBadConn 只能从 Validator、SessionResetter 或查询方法中返回。

All Conn implementations should implement the following interfaces: Pinger, SessionResetter, and Validator.

​	所有 Conn 实现应实现以下接口：Pinger、SessionResetter和Validator。

If named parameters or context are supported, the driver's Conn should implement: ExecerContext, QueryerContext, ConnPrepareContext, and ConnBeginTx.

​	如果支持命名参数或上下文，则驱动程序的Conn应实现：ExecerContext、QueryerContext、ConnPrepareContext和ConnBeginTx。

To support custom data types, implement NamedValueChecker. NamedValueChecker also allows queries to accept per-query options as a parameter by returning ErrRemoveArgument from CheckNamedValue.

​	为了支持自定义数据类型，则需实现NamedValueChecker。NamedValueChecker还允许查询通过返回ErrRemoveArgument来接受每个查询选项作为参数。

If multiple result sets are supported, Rows should implement RowsNextResultSet. If the driver knows how to describe the types present in the returned result it should implement the following interfaces: RowsColumnTypeScanType, RowsColumnTypeDatabaseTypeName, RowsColumnTypeLength, RowsColumnTypeNullable, and RowsColumnTypePrecisionScale. A given row value may also return a Rows type, which may represent a database cursor value.

​	如果支持多个结果集，则Rows应实现RowsNextResultSet。如果驱动程序知道如何描述返回结果中存在的类型，则应实现以下接口：RowsColumnTypeScanType、RowsColumnTypeDatabaseTypeName、RowsColumnTypeLength、RowsColumnTypeNullable和RowsColumnTypePrecisionScale。给定的行值还可以返回一个Rows类型，它可以表示数据库游标值。

Before a connection is returned to the connection pool after use, IsValid is called if implemented. Before a connection is reused for another query, ResetSession is called if implemented. If a connection is never returned to the connection pool but immediately reused, then ResetSession is called prior to reuse but IsValid is not called.

​	在将连接归还给连接池之前，如果实现了 IsValid，则会调用 IsValid。在连接被重用于另一个查询之前，如果实现了 ResetSession，则会调用 ResetSession。如果连接从未返回到连接池而立即被重用，则在重用之前将调用ResetSession，但不会调用IsValid。

## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=55)

``` go 
var Bool boolType
```

Bool is a ValueConverter that converts input values to bools.

​	Bool是一个ValueConverter，用于将输入值转换为布尔值

The conversion rules are:

转换规则如下：

- booleans are returned unchanged

- 布尔值保持不变
- for integer types, 1 is true 0 is false, other integers are an error
- 对于整数类型，1为true，0为false，其他整数为错误
- for strings and []byte, same rules as strconv.ParseBool
- 对于字符串和[]byte，与strconv.ParseBool相同的规则 
- all other types are an error
- 所有其他类型均为错误

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=208)

``` go 
var DefaultParameterConverter defaultConverter
```

DefaultParameterConverter is the default implementation of ValueConverter that's used when a Stmt doesn't implement ColumnConverter.

​	DefaultParameterConverter是ValueConverter的默认实现，当Stmt未实现ColumnConverter时使用。

DefaultParameterConverter returns its argument directly if IsValue(arg). Otherwise, if the argument implements Valuer, its Value method is used to return a Value. As a fallback, the provided argument's underlying type is used to convert it to a Value: underlying integer types are converted to int64, floats to float64, bool, string, and []byte to themselves. If the argument is a nil pointer, ConvertValue returns a nil Value. If the argument is a non-nil pointer, it is dereferenced and ConvertValue is called recursively. Other types are an error.

​	DefaultParameterConverter直接返回其实参（如果IsValue(arg)）。否则，如果实参实现了Valuer接口，则使用其Value方法返回一个Value。作为后备方案，使用提供的实参的底层类型将其转换为Value：底层整数类型转换为int64，浮点数转换为float64，bool、string和[]byte保持不变。如果实参是nil指针，则ConvertValue返回一个nil Value。如果实参是非nil指针，则取消引用并递归调用ConvertValue。其他类型为错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=162)

``` go 
var ErrBadConn = errors.New("driver: bad connection")
```

ErrBadConn should be returned by a driver to signal to the sql package that a driver.Conn is in a bad state (such as the server having earlier closed the connection) and the sql package should retry on a new connection.

​	ErrBadConn应由驱动程序返回，以向sql包发出信号，表明driver.Conn处于不良状态（例如服务器早些时候关闭了连接），sql包应在新的连接上重试。

To prevent duplicate operations, ErrBadConn should NOT be returned if there's a possibility that the database server might have performed the operation. Even if the server sends back an error, you shouldn't return ErrBadConn.

​	为防止重复操作，如果数据库服务器可能已执行操作，则不应返回ErrBadConn。即使服务器返回错误，也不应返回ErrBadConn。

Errors will be checked using errors.Is. An error may wrap ErrBadConn or implement the Is(error) bool method.

​	错误将使用errors.Is进行检查。错误可能会包装ErrBadConn或实现Is(error) bool方法。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=385)

``` go 
var ErrRemoveArgument = errors.New("driver: remove argument from query")
```

ErrRemoveArgument may be returned from NamedValueChecker to instruct the sql package to not pass the argument to the driver query interface. Return when accepting query specific options or structures that aren't SQL query arguments.

​	ErrRemoveArgument可能从NamedValueChecker返回，指示sql包不将实参传递给驱动程序查询接口。在接受非SQL查询实参的特定查询选项或结构时返回。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=148)

``` go 
var ErrSkip = errors.New("driver: skip fast-path; continue as if unimplemented")
```

ErrSkip may be returned by some optional interfaces' methods to indicate at runtime that the fast path is unavailable and the sql package should continue as if the optional interface was not implemented. ErrSkip is only supported where explicitly documented.

​	ErrSkip可能由某些可选接口的方法在运行时返回，以指示快速路径不可用，sql包应继续执行，就像可选接口未实现一样。ErrSkip仅在显式文档中支持。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=102)

``` go 
var Int32 int32Type
```

Int32 is a ValueConverter that converts input values to int64, respecting the limits of an int32 value.

​	Int32是一个ValueConverter，将输入值转换为int64，同时考虑int32值的限制。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/driver.go;l=540)

``` go 
var ResultNoRows noRows
```

ResultNoRows is a pre-defined Result for drivers to return when a DDL command (such as a CREATE TABLE) succeeds. It returns an error for both LastInsertId and RowsAffected.

​	ResultNoRows是预定义的Result，供驱动程序在DDL命令（例如CREATE TABLE）成功时返回。它同时返回LastInsertId和RowsAffected的错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/database/sql/driver/types.go;l=137)

``` go 
var String stringType
```

String is a ValueConverter that converts its input to a string. If the value is already a string or []byte, it's unchanged. If the value is of another type, conversion to string is done with fmt.Sprintf("%v", v).

​	String是一个ValueConverter，将其输入转换为字符串。如果值已经是字符串或[]byte，则保持不变。如果值是其他类型，则使用`fmt.Sprintf("%v", v)`将其转换为字符串。

## 函数

### func IsScanValue 

``` go 
func IsScanValue(v any) bool
```

IsScanValue is equivalent to IsValue. It exists for compatibility.

​	IsScanValue函数和IsValue函数是等价的，为了兼容性而存在。

### func IsValue 

``` go 
func IsValue(v any) bool
```

IsValue reports whether v is a valid Value parameter type.

​	IsValue函数用于判断参数`v`是否是有效的Value参数类型。

## 类型

### type ColumnConverter <- DEPRECATED

```go
type ColumnConverter interface {
	// ColumnConverter returns a ValueConverter for the provided
	// column index. If the type of a specific column isn't known
	// or shouldn't be handled specially, DefaultValueConverter
	// can be returned.
	ColumnConverter(idx int) ValueConverter
}
```

ColumnConverter may be optionally implemented by Stmt if the statement is aware of its own columns' types and can convert from any type to a driver Value.

Deprecated: Drivers should implement NamedValueChecker.

### type Conn 

``` go 
type Conn interface {
    // Prepare returns a prepared statement, bound to this connection.
	// Prepare返回绑定到此连接的预处理语句。
	Prepare(query string) (Stmt, error)

    // Close invalidates and potentially stops any current
	// prepared statements and transactions, marking this
	// connection as no longer in use.
	//
	// Because the sql package maintains a free pool of
	// connections and only calls Close when there's a surplus of
	// idle connections, it shouldn't be necessary for drivers to
	// do their own connection caching.
	//
	// Drivers must ensure all network calls made by Close
	// do not block indefinitely (e.g. apply a timeout).
	// Close 使当前预处理语句和事务无效，标记此连接不再使用。
	//
	// 因为sql包维护一个空闲连接池，并且仅在空闲连接过多时调用Close，
    // 所以驱动程序不需要执行自己的连接缓存。
	//
	// 驱动程序必须确保由Close进行的所有网络调用都不会无限期阻塞（例如应用超时）。
	Close() error

    // Begin starts and returns a new transaction.
	//
	// Deprecated: Drivers should implement ConnBeginTx instead (or additionally).
    // Begin 启动并返回一个新的事务。
    //
    // 已弃用：驱动程序应该实现 ConnBeginTx 来代替(或同时实现)。
	Begin() (Tx, error)
}
```

Conn is a connection to a database. It is not used concurrently by multiple goroutines.

​	Conn 是到数据库的连接。它不允许多个goroutine并发使用。

Conn is assumed to be stateful.

​	Conn 被认为是有状态的。

### type ConnBeginTx  <- go1.8

``` go 
type ConnBeginTx interface {
    // BeginTx starts and returns a new transaction.
	// If the context is canceled by the user the sql package will
	// call Tx.Rollback before discarding and closing the connection.
    // BeginTx 启动并返回一个新的事务。
    // 如果用户取消了上下文，sql包将在丢弃和关闭连接之前调用Tx.Rollback。
    //
    // This must check opts.Isolation to determine if there is a set
	// isolation level. If the driver does not support a non-default
	// level and one is set or if there is a non-default isolation level
	// that is not supported, an error must be returned.
    // 这必须检查opts.Isolation来确定是否有设置的隔离级别。
    // 如果驱动程序不支持非默认级别并且设置了一个非默认级别，
    // 或者如果驱动程序不支持非默认隔离级别，则必须返回一个错误。
    //
    // This must also check opts.ReadOnly to determine if the read-only
	// value is true to either set the read-only transaction property if supported
	// or return an error if it is not supported.
    // 这还必须检查opts.ReadOnly来确定只读值是否为true，
    // 如果支持则设置只读事务属性，如果不可支持则返回错误。
	BeginTx(ctx context.Context, opts TxOptions) (Tx, error)
}
```

ConnBeginTx enhances the Conn interface with context and TxOptions.

​	ConnBeginTx 使用 context 和 TxOptions 增强了 Conn 接口。

### type ConnPrepareContext  <- go1.8

``` go 
type ConnPrepareContext interface {
    // PrepareContext returns a prepared statement, bound to this connection.
	// context is for the preparation of the statement,
	// it must not store the context within the statement itself.
    // PrepareContext 返回一个绑定到此连接的预处理语句。
    // context 是用于准备语句的，它不能在语句本身中存储上下文。
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}
```

ConnPrepareContext enhances the Conn interface with context.

​	ConnPrepareContext 通过context增强 Conn 接口。

### type Connector  <- go1.10

``` go 
type Connector interface {
    // Connect returns a connection to the database.
	// Connect may return a cached connection (one previously
	// closed), but doing so is unnecessary; the sql package
	// maintains a pool of idle connections for efficient re-use.
    // Connect 返回到数据库的连接。
    // Connect 可以返回缓存的连接(先前关闭的连接)，但这样做是不必要的；
    // sql 包维护一个空闲连接池，以便有效地重复使用。
    //
    // The provided context.Context is for dialing purposes only
	// (see net.DialContext) and should not be stored or used for
	// other purposes. A default timeout should still be used
	// when dialing as a connection pool may call Connect
	// asynchronously to any query.
    // 提供的 context.Context 仅用于拨号目的
    //(请参阅 net.DialContext)，不应存储或用于其他目的。
    // 在拨号时仍应使用默认超时时间，因为连接池可能异步调用Connect来执行任何查询。
    //
    // The returned connection is only used by one goroutine at a
	// time.
    // 返回的连接一次只能由一个 goroutine 使用。
	Connect(context.Context) (Conn, error)

    // Driver returns the underlying Driver of the Connector,
	// mainly to maintain compatibility with the Driver method
	// on sql.DB.
    // Driver 返回 Connector 的基础 Driver，
    // 主要是为了与 sql.DB 上的 Driver 方法保持兼容性。
	Driver() Driver
}
```

A Connector represents a driver in a fixed configuration and can create any number of equivalent Conns for use by multiple goroutines.

​	Connector 表示一个具有固定配置的驱动程序，可以为多个 goroutine 创建任意数量的等效 Conns。

A Connector can be passed to sql.OpenDB, to allow drivers to implement their own sql.DB constructors, or returned by DriverContext's OpenConnector method, to allow drivers access to context and to avoid repeated parsing of driver configuration.

​	可以将 Connector 传递给 sql.OpenDB，以允许驱动程序实现自己的 sql.DB 构造函数，或者通过 DriverContext 的 OpenConnector 方法返回，以允许驱动程序访问上下文并避免重复解析驱动程序配置。

If a Connector implements io.Closer, the sql package's DB.Close method will call Close and return error (if any).

​	如果 Connector 实现了 io.Closer，则 sql 包的 DB.Close 方法将调用 Close 并返回错误（如果有）。	

### type Driver 

``` go 
type Driver interface {
    // Open returns a new connection to the database.
	// The name is a string in a driver-specific format.
    // Open 返回到数据库的新连接。
    // name 是一个驱动程序特定的字符串格式。
    //
    // Open may return a cached connection (one previously
	// closed), but doing so is unnecessary; the sql package
	// maintains a pool of idle connections for efficient re-use.
    // Open 可以返回缓存的连接(先前关闭的连接)，但这样做是不必要的；
    // sql 包维护一个空闲连接池，以便高效地重用。
    //
    // The returned connection is only used by one goroutine at a
	// time.
    // 返回的连接一次只能由一个 goroutine 使用。
	Open(name string) (Conn, error)
}
```

Driver is the interface that must be implemented by a database driver.

​	Driver 是必须由数据库驱动程序实现的接口。

Database drivers may implement DriverContext for access to contexts and to parse the name only once for a pool of connections, instead of once per connection.

​	数据库驱动程序可以实现对DriverContext的访问，以便访问上下文，并将 name 解析为连接池中的连接，而不是每个连接解析一次。

### type DriverContext  <- go1.10

``` go 
type DriverContext interface {
	// OpenConnector must parse the name in the same format that Driver.Open
	// parses the name parameter.
    // OpenConnector 必须以与 Driver.Open 解析name参数相同的格式解析名称
	OpenConnector(name string) (Connector, error)
}
```

If a Driver implements DriverContext, then sql.DB will call OpenConnector to obtain a Connector and then invoke that Connector's Connect method to obtain each needed connection, instead of invoking the Driver's Open method for each connection. The two-step sequence allows drivers to parse the name just once and also provides access to per-Conn contexts.

​	如果一个 Driver 实现了 DriverContext 接口，则 sql.DB 将调用 OpenConnector 方法来获取 Connector，然后调用该 Connector 的 Connect 方法来获取每个所需的连接，而不是为每个连接调用Driver的 Open 方法。这种两步序列允许驱动程序只解析name一次，并且还提供对每个Conn上下文的访问。

### type Execer <- DEPRECATED

``` go
type Execer interface {
	Exec(query string, args []Value) (Result, error)
}
```

Execer is an optional interface that may be implemented by a Conn.

​	Execer是一个可选的接口，可以由Conn（这里是泛指的意思下Conn）实现。

If a Conn implements neither ExecerContext nor Execer, the sql package's DB.Exec will first prepare a query, execute the statement, and then close the statement.

​	如果一个Conn（这里是泛指的意思下Conn）既没有实现ExecerContext也没有实现Execer，则sql包的DB.Exec将首先准备一个查询，执行语句，然后关闭语句。

Exec may return ErrSkip.

​	Exec可以返回ErrSkip。

Deprecated: Drivers should implement ExecerContext instead.

​	已弃用：驱动程序应改为实现ExecerContext接口。



### type ExecerContext  <- go1.8

``` go 
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args []NamedValue) (Result, error)
}
```

ExecerContext is an optional interface that may be implemented by a Conn.

​	ExecerContext 是一个可选接口，可以由 Conn （这里是泛指的意思下Conn）实现。

If a Conn does not implement ExecerContext, the sql package's DB.Exec will fall back to Execer; if the Conn does not implement Execer either, DB.Exec will first prepare a query, execute the statement, and then close the statement.

​	如果 Conn （这里是泛指的意思下Conn）没有实现 ExecerContext，则 sql 包的 DB.Exec 将回退到 Execer；如果 Conn （这里是泛指的意思下Conn）也没有实现 Execer，则 DB.Exec 将首先准备查询，执行语句，然后关闭语句。

ExecContext may return ErrSkip.

​	ExecContext 方法可以返回 ErrSkip。

ExecContext must honor the context timeout and return when the context is canceled.

​	ExecContext 方法必须遵守上下文超时，并在上下文取消时返回。

### type IsolationLevel  <- go1.8

``` go 
type IsolationLevel int
```

IsolationLevel is the transaction isolation level stored in TxOptions.

​	IsolationLevel 是存储在 TxOptions 中的事务隔离级别。

This type should be considered identical to sql.IsolationLevel along with any values defined on it.

​	该类型应被视为与sql.IsolationLevel类型相同以及任何在其上定义的值。

### type NamedValue  <- go1.8

``` go 
type NamedValue struct {
    // If the Name is not empty it should be used for the parameter identifier and
	// not the ordinal position.
    // 如果 Name 不为空，应该用于参数标识符，而不是位置。
    //
    // Name will not have a symbol prefix.
    // Name 不会有符号前缀。
	Name string

    // Ordinal position of the parameter starting from one and is always set.
	// 参数的序号位置，从一开始计数，并始终设置。
	Ordinal int	

    // Value is the parameter value.
    // Value是参数的值。
	Value Value
}
```

NamedValue holds both the value name and value.

​	NamedValue 持有值的名称和值。

### type NamedValueChecker  <- go1.9

``` go 
type NamedValueChecker interface {
    // CheckNamedValue is called before passing arguments to the driver
	// and is called in place of any ColumnConverter. CheckNamedValue must do type
	// validation and conversion as appropriate for the driver.
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

### type NotNull 

``` go 
type NotNull struct {
	Converter ValueConverter
}
```

​	NotNull结构体是一种类型，通过禁止nil值但否则委托给另一个ValueConverter来实现ValueConverter。

#### (NotNull) ConvertValue 

``` go 
func (n NotNull) ConvertValue(v any) (Value, error)
```

### type Null 

``` go 
type Null struct {
	Converter ValueConverter
}
```

​	Null 结构体是一个实现了 ValueConverter 接口的类型，它允许 nil 值，但除此之外会委托给另一个 ValueConverter。

#### (Null) ConvertValue 

``` go 
func (n Null) ConvertValue(v any) (Value, error)
```

### type Pinger  <- go1.8

``` go 
type Pinger interface {
	Ping(ctx context.Context) error
}
```

​	Pinger 是一个可选的接口，可以由 Conn（这里是泛指的意思下Conn） 实现。

​	如果一个 Conn（这里是泛指的意思下Conn） 没有实现 Pinger接口，那么 sql 包中的 DB.Ping 和 DB.PingContext 将检查是否至少有一个可用的 Conn。

​	如果 sql.Conn.Ping 方法返回 ErrBadConn，则sql.DB.Ping 方法和 sql.DB.PingContext 方法将从池中移除该 Conn。

### type Queryer <-DEPRECATED

``` go
type Queryer interface {
	Query(query string, args []Value) (Rows, error)
}
```

Queryer is an optional interface that may be implemented by a Conn.

If a Conn implements neither QueryerContext nor Queryer, the sql package's DB.Query will first prepare a query, execute the statement, and then close the statement.

Query may return ErrSkip.

Deprecated: Drivers should implement QueryerContext instead.

### type QueryerContext  <- go1.8

``` go 
type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args []NamedValue) (Rows, error)
}
```

​	QueryerContext是一个可选的接口，可以由Conn（这里是泛指的意思下Conn）实现。

​	如果Conn（这里是泛指的意思下Conn）未实现QueryerContext方法，则sql包的DB.Query方法将回退到Queryer；如果Conn（这里是泛指的意思下Conn）也未实现Queryer，则DB.Query方法将首先准备一个查询，执行语句，然后关闭语句。

​	QueryContext方法可能会返回ErrSkip。

​	QueryContext方法必须遵守上下文超时并在取消上下文时返回。

### type Result 

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

### type Rows 

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

### type RowsAffected 

``` go 
type RowsAffected int64
```

​	RowsAffected 表示执行 INSERT 或 UPDATE 操作所影响的行数，实现了 Result 接口。

#### (RowsAffected) LastInsertId 

``` go 
func (RowsAffected) LastInsertId() (int64, error)
```

#### (RowsAffected) RowsAffected 

``` go 
func (v RowsAffected) RowsAffected() (int64, error)
```

### type RowsColumnTypeDatabaseTypeName  <- go1.8

``` go 
type RowsColumnTypeDatabaseTypeName interface {
	Rows
	ColumnTypeDatabaseTypeName(index int) string
}
```

​	RowsColumnTypeDatabaseTypeName接口可以被Rows（这里是泛指的意思下Rows）实现。它应该返回数据库系统类型名称，但不包括长度信息。**类型名称应大写**。以下是各种类型的返回示例："VARCHAR"、"NVARCHAR"、"VARCHAR2"、"CHAR"、"TEXT"、"DECIMAL"、"SMALLINT"、"INT"、"BIGINT"、"BOOL"、"[]BIGINT"、"JSONB"、"XML"、"TIMESTAMP"。

### type RowsColumnTypeLength  <- go1.8

``` go 
type RowsColumnTypeLength interface {
	Rows
	ColumnTypeLength(index int) (length int64, ok bool)
}
```

​	RowsColumnTypeLength接口可以被Rows（这里是泛指的意思下Rows）实现。如果列是变长类型，则它应返回列类型的长度。如果列不是变长类型，则应该返回 `false`。如果长度除系统限制外不受限制，则应返回 math.MaxInt64。以下是各种类型的返回值示例：

```
TEXT          (math.MaxInt64, true)
varchar(10)   (10, true)
nvarchar(10)  (10, true)
decimal       (0, false)
int           (0, false)
bytea(30)     (30, true)
```

### type RowsColumnTypeNullable  <- go1.8

``` go 
type RowsColumnTypeNullable interface {
	Rows
	ColumnTypeNullable(index int) (nullable, ok bool)
}
```

​	RowsColumnTypeNullable 接口可以由Rows（这里是泛指的意思下Rows）实现。如果已知某列为可能为null，则该可空值的值为 `true`；如果已知该列不可能为null，则该可空值的值为 `false`。如果该列的可空值属性未知，则 `ok` 应为 `false`。

### type RowsColumnTypePrecisionScale  <- go1.8

``` go 
type RowsColumnTypePrecisionScale interface {
	Rows
	ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool)
}
```

​	RowsColumnTypePrecisionScale 接口可以由Rows（这里是泛指的意思下Rows）实现。它应返回十进制类型的精度和小数位数。如果不适用，则应将`ok`设置为`false`。以下是各种类型的返回值示例：

```
decimal(38, 4)    (38, 4, true)
int               (0, 0, false)
decimal           (math.MaxInt64, math.MaxInt64, true)
```

### type RowsColumnTypeScanType  <- go1.8

``` go 
type RowsColumnTypeScanType interface {
	Rows
	ColumnTypeScanType(index int) reflect.Type
}
```

​	RowsColumnTypeScanType 接口可以由Rows（这里是泛指的意思下Rows） 实现。它应返回可以用于扫描类型的值类型。例如，对于数据库列类型"bigint"，这应返回"`reflect.TypeOf(int64(0))`"。

### type RowsNextResultSet  <- go1.8

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

​	RowsNextResultSet 接口通过提供一种方式来向驱动程序发出信号，使其前进到下一个结果集，从而扩展了 Rows 接口。

### type SessionResetter  <- go1.10

``` go 
type SessionResetter interface {
	// 如果在连接被使用过之前，ResetSession 会在连接上执行查询之前被调用。
    // 如果驱动程序返回 ErrBadConn，连接将被丢弃。
	ResetSession(ctx context.Context) error
}
```

​	SessionResetter 可以由 Conn （这里是泛指的意思下Conn ）实现，以允许驱动程序重置与连接关联的会话状态并发出不良连接信号。

### type Stmt 

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

​	Stmt接口是预处理语句。它绑定到一个Conn （这里是泛指的意思下Conn ），并且不能同时被多个goroutine使用。

### type StmtExecContext  <- go1.8

``` go 
type StmtExecContext interface {
	// ExecContext方法执行不返回行的查询，例如INSERT或UPDATE。
	// ExecContext方法必须遵守上下文超时并在取消时返回。
	ExecContext(ctx context.Context, args []NamedValue) (Result, error)
}
```

​	StmtExecContext接口通过提供带有上下文的Exec方法来增强Stmt接口。

### type StmtQueryContext  <- go1.8

``` go 
type StmtQueryContext interface {
	// QueryContext方法执行可能返回行的查询，例如SELECT。
	// QueryContext方法必须遵守上下文超时并在取消时返回。
	QueryContext(ctx context.Context, args []NamedValue) (Rows, error)
}
```

​	StmtQueryContext接口通过提供带有上下文的Query方法来增强Stmt接口。

### type Tx 

``` go 
type Tx interface {
	Commit() error
	Rollback() error
}
```

​	Tx 是事务。

### type TxOptions  <- go1.8

``` go 
type TxOptions struct {
	Isolation IsolationLevel
	ReadOnly  bool
}
```

​	TxOptions结构体存储事务的选项。

​	此类型应被认为与 sql.TxOptions 相同。

### type Validator  <- go1.15

``` go 
type Validator interface {
	// IsValid方法在将连接放入连接池之前调用。
    // 如果返回false，则连接将被丢弃。
	IsValid() bool
}
```

​	Validator接口可以被 （这里是泛指的意思下Conn ）实现，允许驱动程序发出信号，指示连接是否有效或是否应该被丢弃。

​	如果实现了 Validator，即使连接应该被丢弃，驱动程序也可以返回查询的基本错误。

### type Value 

``` go 
type Value any
```

​	Value类型是驱动程序必须能够处理的值。它可以是nil，也可以是数据库驱动程序的NamedValueChecker接口处理的类型，或者是这些类型之一的实例：

```
int64
float64
bool
[]byte
string
time.Time
```

​	如果驱动程序支持游标，则返回的 Value 还可以在此包中实现 Rows 接口。例如，当用户选择游标时，如 "`select cursor(select * from my_table) from dual`"，就会使用到这种情况。如果 select 中的 Rows （这里是泛指的意思下Rows ）关闭了，游标 Rows （这里是泛指的意思下Rows ）也会被关闭。

### type ValueConverter 

``` go 
type ValueConverter interface {
	// ConvertValue方法将一个值转换为driver.Value。
	ConvertValue(v any) (Value, error)
}
```

​	ValueConverter接口是提供 ConvertValue 方法的接口。

​	驱动程序包提供了 ValueConverter 的各种实现，以提供驱动程序之间的一致的转换实现。ValueConverter 有以下几个用途：

- 将 sql 包提供的 Value 类型转换为数据库表的特定列类型，并确保其适合，例如确保特定 int64 适合于表的 uint16 列。
- 将数据库中给定的值转换为驱动程序的一种 Value 类型。
-  由 sql 包使用，将驱动程序的 Value 类型转换为用户的类型进行扫描。

### type Valuer 

``` go 
type Valuer interface {
    // Value方法返回一个driver.Value。
    // Value方法不得panic。
	Value() (Value, error)
}
```

​	Valuer接口是提供 Value 方法的接口。

​	实现 Valuer 接口的类型能够将自身转换为驱动程序 Value。