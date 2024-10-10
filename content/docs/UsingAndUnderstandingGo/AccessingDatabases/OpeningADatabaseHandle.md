+++
title = "打开数据库句柄"
date = 2023-05-17T15:03:14+08:00
weight = 2
description = ""
isCJKLanguage = true
draft = false

+++
# Opening a database handle - 打开数据库句柄

> 原文：[https://go.dev/doc/database/open-handle](https://go.dev/doc/database/open-handle)

​	[database/sql]({{< ref "/stdLib/database/sql" >}})包通过减少管理连接的需求来简化数据库访问。与许多数据访问API不同，使用`database/sql`，您不需要显式地打开连接、执行操作然后关闭连接。相反，您的代码会打开一个代表连接池的数据库句柄，然后使用该句柄执行数据访问操作，仅在需要释放资源（例如检索到的行或预处理语句持有的资源）时调用 `Close` 方法。

​	换句话说，由 [sql.DB]({{< ref "/stdLib/database/sql#type-db" >}}) 表示的数据库句柄处理连接，代表您的代码打开和关闭它们。当您的代码使用句柄执行数据库操作时，这些操作可以并发地访问数据库。有关更多信息，请参阅[Managing connections （管理连接）](../ManagingConnections)。

> 注意：您还可以保留一个数据库连接。有关更多信息，请参阅[Using dedicated connections （使用专用连接）](../ManagingConnections#使用专用连接)。

​	除了 `database/sql` 包中提供的 API 之外，Go 社区还为所有最常见的（以及许多不常见的）数据库管理系统（DBMS）开发了驱动程序。

​	打开数据库句柄时，您需要遵循以下高级步骤：

1. 找到驱动程序。

   驱动程序会在您的Go代码和数据库之间转换（translates）请求和响应。有关更多信息，请参阅 [Locating and importing a database driver （找到并导入数据库驱动程序）](#locating-and-importing-a-database-driver-找到并导入数据库驱动程序)。

2. 打开数据库句柄。

   导入驱动程序后，您可以为特定的数据库打开一个句柄。有关更多信息，请参阅[Opening a database handle （打开一个数据库句柄）](#打开一个数据库句柄)。

3. 确认连接。

   一旦打开了数据库句柄，您的代码就可以检查是否有可用连接。有关更多信息，请参阅[Confirming a connection （确认连接）](#确认一个连接)。

​	您的代码通常不会显式地打开或关闭数据库连接 —— 这是由数据库句柄完成的。但是，您的代码应该释放它沿途获得的资源，例如包含查询结果的 `sql.Rows`。有关更多信息，请参阅[Freeing resources（释放资源）](#释放资源)。

### Locating and importing a database driver 找到并导入数据库驱动程序

​	你需要一个支持你正在使用的DBMS的数据库驱动程序。要找到你的数据库的驱动程序，请参阅[SQLDrivers](https://github.com/golang/go/wiki/SQLDrivers)。

​	为了让驱动程序对你的代码可用，你可以像导入其他Go包一样导入它。这里有一个示例：

```go 
import "github.com/go-sql-driver/mysql"
```

> 注意，如果你没有直接从驱动程序包中调用任何函数（例如，当它被`sql`包隐式使用时），你需要使用空白导入，即在导入路径前加上下划线：
>
> ```go
> import _ "github.com/go-sql-driver/mysql"
> ```

> 注意：作为一个最佳实践，避免使用数据库驱动程序自己的API进行数据库操作。相反，使用`database/sql`包中的函数。这将有助于保持你的代码与DBMS松散耦合，使其更容易切换到不同的DBMS（如果需要的话）。

### 打开数据库句柄

​	一个`sql.DB`数据库句柄提供了从数据库读取和写入的能力，无论是单独的还是在一个事务中。

​	你可以通过调用`sql.Open`（接收一个连接字符串）或`sql.OpenDB`（接收一个`driver.Connector`）来获取数据库句柄。两者都返回一个指向[sql.DB]({{< ref "/stdLib/database/sql#type-db" >}})的指针。

> 注意：请确保将数据库凭据保留在你的Go源代码之外。有关更多信息，请参阅[Storing database credentials （存储数据库凭证）](#存储数据库凭证)。

#### 使用连接字符串打开

​	当你想使用连接字符串时，请使用[sql.Open]({{< ref "/stdLib/database/sql#func-open" >}})函数。字符串的格式将取决于你使用的驱动程序。

​	这是一个针对MySQL的示例：

```go 
db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/jazzrecords")
if err != nil {
    log.Fatal(err)
}
```

​	然而，您可能会发现，以更结构化的方式捕获连接属性会使您的代码更具可读性。这些细节会因驱动而异。

​	例如，您可以把前面的例子换成下面的例子，它使用MySQL驱动的[Config]({{< ref "/thirdPkg/go_sql_driver_mysql#type-config---130">}})来指定属性，并使用[FormatDSN]({{< ref "/thirdPkg/go_sql_driver_mysql#config-formatdsn---130">}})方法来构建一个连接字符串。  

​	然而，你可能会发现以更结构化的方式捕获连接属性会使你的代码更具可读性。详细信息将因驱动程序而异。

​	例如，你可以用以下方式替换前面的示例，该示例使用MySQL驱动程序的[Config]({{< ref "/thirdPkg/go_sql_driver_mysql#type-config---130">}})指定属性并使用其[FormatDSN]({{< ref "/thirdPkg/go_sql_driver_mysql#config-formatdsn---130">}})方法构建连接字符串。

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

#### 使用连接器打开

​	当你想利用连接字符串中不可用的特定于驱动程序的连接功能时，请使用[sql.OpenDB]({{< ref "/stdLib/database/sql#func-opendb----go110" >}})函数。每个驱动程序都支持自己的一组连接属性，通常提供自定义特定于DBMS的连接请求的方法。

​	将前面的`sql.Open`示例改编为使用`sql.OpenDB`，你可以使用以下代码创建一个句柄：

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

​	你的代码应该检查尝试创建句柄时的错误，例如使用`sql.Open`。这不是一个连接错误。相反，如果`sql.Open`无法初始化句柄，你会得到一个错误。例如，如果它无法解析你指定的`DSN`，这种情况可能会发生。

### 确认一个连接

​		当你打开数据库句柄时，`sql`包可能不会立即创建新的数据库连接本身。相反，它可能会在代码需要时创建连接。如果你不会立即使用数据库并希望确认可以建立连接，请调用[Ping]({{< ref "/stdLib/database/sql#db-ping----go11" >}})方法或[PingContext]({{< ref "/stdLib/database/sql#db-pingcontext----go18" >}})方法。

​	下面示例中的代码ping数据库以确认连接。

```go 
db, err = sql.Open("mysql", connString)

// Confirm a successful connection.
if err := db.Ping(); err != nil {
    log.Fatal(err)
}
```

### 存储数据库凭据

​	避免将数据库凭据（credentials ）存储在你的Go源代码中，这可能会将你的数据库内容暴露给他人。相反，找到一个方法将它们存储在你代码之外但可供其使用的位置。例如，考虑使用一个秘密管理器应用程序来存储凭据并提供一个API，你的代码可以使用该API检索用于与DBMS进行身份验证的凭据。

​	一种流行的方法是在程序启动之前将秘密存储在环境中，可能是从秘密管理器加载的，然后你的Go程序可以使用[os.Getenv]({{< ref "/stdLib/os/os#func-getenv">}})读取它们：

```go 
username := os.Getenv("DB_USER")
password := os.Getenv("DB_PASS")
```

​	这种方法还允许你在本地测试时自己设置环境变量。

### 释放资源

​	尽管你没有显式地使用`database/sql`包管理或关闭连接，但你的代码应该在不再需要时释放已获得的资源。这些资源可以包括由表示从查询返回的数据的`sql.Rows`和表示预处理语句的`sql.Stmt`持有的资源。

​	通常，你通过延迟调用`Close`函数来关闭资源，以便在外层函数退出之前释放资源。

​	下面示例中的代码将`Close`延迟以释放由[sql.Rows]({{< ref "/stdLib/database/sql#type-rows">}})持有的资源。

```go  hl_lines="5 5"
rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artist)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

// Loop through returned rows.
```