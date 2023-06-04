+++
title = "管理连接"
date = 2023-05-17T15:03:14+08:00
weight = 8
description = ""
isCJKLanguage = true
draft = false
+++
# Managing connections - 管理连接

> 原文：[https://go.dev/doc/database/manage-connections](https://go.dev/doc/database/manage-connections)

​	对于绝大多数程序，您不需要调整`sql.DB`连接池的默认值。但是对于一些高级程序，您可能需要优化连接池的参数或者显式地处理连接。本主题解释了如何做。

​	[sql.DB]({{< ref "/docs/StdLib/database/sql#type-db">}})数据库句柄对于多个goroutine的并发使用是安全的（意味着该句柄是其他语言可能称之为 "`线程安全（thread-safe）` "的）。其他一些数据库访问库是基于一次只能用于一个操作的连接。为了弥合这一差距，每个`sql.DB`都管理着一个与底层数据库的活动连接池，根据您 Go 程序的并行性需要创建新连接池。

​	连接池适用于大多数数据访问需求。当您调用一个`sql.DB`的`Query`或`Exec`方法时，`sql.DB`实现从池中检索一个可用的连接，或者在需要时创建一个连接。当不再需要连接时，包会将其返回到池中。这支持了数据库访问的高度并行性。

### 设置连接池属性

​	您可以设置属性来指导`sql`包如何管理连接池。要获得关于这些属性的效果的统计信息，请使用[DB.Stats]({{< ref "/docs/StdLib/database/sql#db-stats----go15">}})。

#### 设置开放连接的最大数量

​	[DB.SetMaxOpenConns]({{< ref "/docs/StdLib/database/sql#db-setmaxopenconns----go12">}})对开放连接的数量进行了限制。超过这个限制，新的数据库操作将等待一个现有的操作完成，这时`sql.DB`将创建另一个连接。默认情况下，当需要一个连接的时候，`sql.DB`会在所有现有的连接都处于使用状态时创建一个新的连接。

​	请记住，设置限制使得数据库的使用类似于获取一个锁或信号，其结果是您的应用程序可能会在等待新的数据库连接时发生死锁（deadlock ）。

#### 设置空闲（idle）连接的最大数量

​	[DB.SetMaxIdleConns]({{< ref "/docs/StdLib/database/sql#db-setmaxidleconns----go11">}})更改`sql.DB`维护的最大空闲连接数的限制。

​	当一个SQL操作在一个给定的数据库连接上完成后，它通常不会立即关闭：应用程序可能很快就会再次需要一个连接，保持开放的连接可以避免在下一个操作中重新连接到数据库。默认情况下，`sql.DB`在任何时候都保持两个空闲连接。提高这个限制可以避免在有大量并行性的程序中频繁的重新连接。

#### 设置一个连接可以空闲（idle）的最大时间量

​	[DB.SetConnMaxIdleTime]({{< ref "/docs/StdLib/database/sql#db-setconnmaxidletime----go115">}})设置了一个连接在被关闭之前可以空闲的最大时间。这将导致`sql.DB`关闭那些空闲时间超过给定时间的连接。

​	默认情况下，当将空闲连接添加到连接池时，它将一直保持在那里，直到再次需要它为止。当使用`DB.SetMaxIdleConns`来增加并行活动爆发期间允许的空闲连接数时，也可以使用`DB.SetConnMaxIdleTime`安排在系统安静时释放这些连接。

#### 设置连接的最大生存期

​	使用[DB.SetConnMaxLifetime]({{< ref "/docs/StdLib/database/sql#db-setconnmaxlifetime----go16">}})可以设置一个连接在被关闭之前可以保持开放的最大时间长度。

​	默认情况下，一个连接可以在任意长的时间内被使用和重复使用，但要遵守上述的限制。在一些系统中，比如那些使用负载平衡的数据库服务器的系统，确保应用程序在不重新连接的情况下不会使用某个特定的连接太久是很有帮助的。

### 使用专用连接

​	`database/sql`包包含一些函数，当数据库可能为在特定连接上执行的操作序列赋予隐式含义时，您可以使用这些函数。

​	最常见的例子是事务，它通常以`BEGIN`命令开始，以`COMMIT`或`ROLLBACK`命令结束，并在整个事务中包括这些命令之间的连接上发出的所有命令。对于这种使用情况，请使用`sql`包的事务支持。参见 在较大的事务中运行一个查询。更多信息，请参阅[Executing transactions （执行事务）](../ExecutingTransactions) 。

​	对于其他必须在同一连接上执行一系列单独操作的使用情况，`sql`包提供了专用连接。[DB.Conn]({{< ref "/docs/StdLib/database/sql#db-conn----go19">}})获得一个专用连接，即[sql.Conn]({{< ref "/docs/StdLib/database/sql#type-conn----go19">}})。`sql.Conn`有`BeginTx`、`ExecContext`、`PingContext`、`PrepareContext`、`QueryContext`和`QueryRowContext`等方法，这些方法的行为与DB上的同类方法类似，但只使用专用连接。当完成专用连接后，您的代码必须使用`Conn.Close`释放它。