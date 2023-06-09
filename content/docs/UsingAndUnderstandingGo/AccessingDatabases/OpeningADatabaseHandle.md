+++
title = "打开一个数据库句柄"
date = 2023-05-17T15:03:14+08:00
weight = 2
description = ""
isCJKLanguage = true
draft = false
+++
# Opening a database handle - 打开一个数据库句柄

> 原文：[https://go.dev/doc/database/open-handle](https://go.dev/doc/database/open-handle)

​	[database/sql]({{< ref "/stdLib/database/sql" >}})包通过减少您需要的管理连接来简化数据库访问。与许多数据访问API不同，使用`database/sql`，您不需要明确地打开一个连接，进行工作，然后关闭连接。相反，您的代码会打开一个代表连接池的数据库句柄，然后用这个句柄执行数据访问操作，仅在需要释放资源（比如那些被检索的行或预处理语句所持有的资源）时才调用`Close`方法。

​	换句话说，它是由[sql.DB]({{< ref "/stdLib/database/sql#type-db" >}})表示的数据库句柄，代表您的代码来处理连接、打开和关闭连接。当您的代码使用句柄来执行数据库操作时，这些操作对数据库有并发的访问。更多信息请参见[Managing connections （管理连接）](../ManagingConnections)。

注意：您也可以保留一个数据库连接。更多信息，请参见 [Using dedicated connections （使用专用连接）](../ManagingConnections#使用专用连接)。

​	除了`database/sql`包中的API之外，Go社区还为所有最常见的（以及许多不常见的）数据库管理系统（DBMSes）开发了驱动程序。

当打开一个数据库句柄时，请遵循以下高级步骤：

1. 找到一个驱动程序。

   驱动程序会在您的Go代码和数据库之间转换（translates）请求和响应。更多信息，请参见 [Locating and importing a database driver （找到并导入数据库驱动程序）](#locating-and-importing-a-database-driver-找到并导入数据库驱动程序)。

2. 打开一个数据库句柄。

   在您导入驱动程序后，您可以为特定的数据库打开一个句柄。更多信息，请参见[Opening a database handle （打开一个数据库句柄）](#打开一个数据库句柄)。

3. 确认连接。

   一旦您打开了一个数据库句柄，您的代码就可以检查是否有连接可用。更多信息，请参见[Confirming a connection （确认连接）](#确认一个连接)。

​	您的代码通常不会明确地打开或关闭数据库连接——那是由数据库句柄完成的。**然而，您的代码应该释放它沿途获得的资源**，例如包含查询结果的`sql.Rows`。更多信息请参见[Freeing resources（释放资源）](#释放资源)。

### Locating and importing a database driver 找到并导入数据库驱动程序

​	您需要一个支持您所使用的数据库管理系统的数据库驱动程序。要为您的数据库找到一个驱动，请看[SQLDrivers](https://github.com/golang/go/wiki/SQLDrivers)。

​	为了使您的代码能够使用该驱动程序，您可以像导入其他Go包一样导入它。下面是一个例子：

```go 
import "github.com/go-sql-driver/mysql"
```

​	注意，如果您没有直接从驱动包中调用任何函数——比如它被`sql`包隐式使用——您需要使用空白导入，它在导入路径前加了一个下划线：

```go 
import _ "github.com/go-sql-driver/mysql"
```

> 注意
>
> ​	作为一个最佳实践，请避免使用数据库驱动程序自己的API进行数据库操作。相反，使用`database/sql`包中的函数。这将有助于保持您的代码与DBMS的松耦合，使您在需要时更容易切换到不同的DBMS。

### 打开一个数据库句柄

​	`sql.DB`数据库句柄提供了从数据库读取和写入数据库的能力，无论是单独的还是在一个事务中。

​	您可以通过调用`sql.Open`（它接收一个连接字符串）或者`sql.OpenDB`（它接收一个`driver.Connector`）来获得一个数据库句柄。两者都返回一个指向[sql.DB]({{< ref "/stdLib/database/sql#type-db" >}})的指针。

注意：请确保您的数据库凭证不在您的Go源代码中。更多信息请参见[Storing database credentials （存储数据库凭证）](#存储数据库凭证) 。

#### 用连接字符串打开

​	当您想使用连接字符串进行连接时，请使用[sql.Open]({{< ref "/stdLib/database/sql#func-open" >}})函数。字符串的格式将根据您使用的驱动程序而有所不同。

下面是一个关于MySQL的例子：

```go 
db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/jazzrecords")
if err != nil {
    log.Fatal(err)
}
```

​	然而，您可能会发现，以更结构化的方式捕获连接属性会使您的代码更具可读性。这些细节会因驱动而异。

​	例如，您可以把前面的例子换成下面的例子，它使用MySQL驱动的[Config](https://pkg.go.dev/github.com/go-sql-driver/mysql#Config)来指定属性，并使用[FormatDSN](https://pkg.go.dev/github.com/go-sql-driver/mysql#Config.FormatDSN)方法来构建一个连接字符串。

```go 
// Specify connection properties.
cfg := mysql.Config{
    User:   username,
    Passwd: password,
    Net:    "tcp",
    Addr:   "127.0.0.1:3306",
    DBName: "jazzrecords",
}

// Get a database handle.
db, err = sql.Open("mysql", cfg.FormatDSN())
if err != nil {
    log.Fatal(err)
}
```

#### 用一个连接器打开

​	当您想利用连接字符串中没有的特定驱动程序的连接特性时，请使用[sql.OpenDB]({{< ref "/stdLib/database/sql#func-opendb----go110" >}})函数。每个驱动都支持自己的连接属性集，通常提供了定制特定于DBMS的连接请求的方式。

​	将前面的`sql.Open`例子改成使用`sql.OpenDB`，您可以用下面的代码创建一个句柄：

```go 
// Specify connection properties.
cfg := mysql.Config{
    User:   username,
    Passwd: password,
    Net:    "tcp",
    Addr:   "127.0.0.1:3306",
    DBName: "jazzrecords",
}

// Get a driver-specific connector.
connector, err := mysql.NewConnector(&cfg)
if err != nil {
    log.Fatal(err)
}

// Get a database handle.
db = sql.OpenDB(connector)
```

#### 处理错误

​	您的代码应该检查是否在尝试创建句柄时出错，比如用`sql.Open`。这不会是一个连接错误。相反，如果`sql.Open`无法初始化句柄，您会得到一个错误。这可能发生，例如，如果它无法解析您指定的`DSN`。

### 确认一个连接

​	当您打开一个数据库句柄时，`sql`包可能不会立即自己创建一个新的数据库连接。相反，它可能会在您的代码需要它时创建连接。如果您不会马上使用数据库，并且想确认连接可以被建立，可以调用[Ping]({{< ref "/stdLib/database/sql#db-ping----go11" >}})或[PingContext]({{< ref "/stdLib/database/sql#db-pingcontext----go18" >}})。

下面的例子中的代码对数据库进行ping，以确认连接。

```go 
db, err = sql.Open("mysql", connString)

// Confirm a successful connection.
if err := db.Ping(); err != nil {
    log.Fatal(err)
}
```

### 存储数据库凭证

​	避免在您的Go源代码中存储数据库凭证（credentials ），这可能会将您的数据库内容暴露给其他人。相反，要想办法将其存储在代码之外的位置，但对代码是可用的。例如，考虑一个保密应用程序，该应用程序存储凭据并提供一个 API，您的代码可以使用该 API 检索凭据，以便对 DBMS 进行身份验证。

​	一种流行的方法是在程序启动前将秘密存储在环境中，可能是从秘密管理器中加载，然后您的 Go 程序可以使用 [os.Getenv]({{< ref "/stdLib/os/os#func-getenv">}}) 读取这些秘密：

```go 
username := os.Getenv("DB_USER")
password := os.Getenv("DB_PASS")
```

​	这种方法还允许您为本地测试自己设置环境变量。

### 释放资源

​	虽然您没有显式地管理或关闭`database/sql`包提供连接，但您的代码应该在不再需要时释放它所获得的资源。这些可能包括由表示从查询返回的数据的`sql.Rows`或代表预处理语句的`sql.Stmt`持有的资源。

​	通常，您通过推迟对`Close`函数的调用来关闭资源，以便在外层函数退出之前释放资源。

​	在下面的例子中，代码延迟`Close`，以释放由[sql.Rows]({{< ref "/stdLib/database/sql#type-rows">}})持有的资源。

```go  hl_lines="5 5"
rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artist)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

// Loop through returned rows.
```