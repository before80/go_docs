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

​	对于绝大多数程序来说，你不需要调整[sql.DB]({{< ref "/stdLib/database/sql#type-db">}})连接池的默认值。但对于一些高级程序，你可能需要调整连接池参数或显式处理连接。本主题将解释如何操作。

​	[sql.DB]({{< ref "/stdLib/database/sql#type-db">}})数据库句柄可以安全地被多个goroutine并发使用（意味着该句柄在其他语言中可能被称为"线程安全（thread-safe）"）。其他一些数据库访问库基于只能一次用于一个操作的连接。为了弥补这一差距，每个`sql.DB`都管理着一个与底层数据库的活动连接池，根据Go程序中的并行性需要创建新的连接。

​	连接池适用于大多数数据访问需求。当你调用`sql.DB`的`Query`或`Exec`方法时，`sql.DB`实现会从连接池中检索可用的连接，如果需要，则创建一个新连接。当不再需要连接时，该包将其返回到连接池中。这支持对数据库访问的高度并行性。

### 设置连接池属性

​	你可以设置一些属性来指导`sql`包如何管理连接池。要获取有关这些属性效果的统计信息，请使用[DB.Stats]({{< ref "/stdLib/database/sql#db-stats----go15">}})。

#### 设置开放连接的最大数量

​	[DB.SetMaxOpenConns]({{< ref "/stdLib/database/sql#db-setmaxopenconns----go12">}})会对打开的连接数施加限制。超过此限制后，新的数据库操作将等待现有操作完成，此时`sql.DB`将创建另一个连接。默认情况下，当需要一个连接时，如果所有现有连接都在使用中，sql.DB将创建一个新的连接。

​	请注意，设置限制会使数据库使用类似于获取锁或信号量，结果你的应用程序可能会因为等待新的数据库连接而发生死锁（deadlock ）。

#### 设置空闲（idle）连接的最大数量

​	[DB.SetMaxIdleConns]({{< ref "/stdLib/database/sql#db-setmaxidleconns----go11">}})更改sql.DB维护的最大空闲连接数的限制。

​	当给定数据库连接上的SQL操作完成时，通常不会立即关闭连接：应用程序可能很快再次需要它，保持打开的连接可以避免为下一个操作重新连接到数据库。默认情况下，任何时候`sql.DB`都会保留两个空闲连接。在具有显著并行性的程序中，提高限制可以避免频繁重新连接。

#### 设置连接可以空闲（idle）的最长时间

​	[DB.SetConnMaxIdleTime]({{< ref "/stdLib/database/sql#db-setconnmaxidletime----go115">}})设置连接可以在关闭之前保持空闲的最长时间。这将导致`sql.DB`关闭那些空闲时间超过给定持续时间的连接。

​	默认情况下，当空闲连接添加到连接池中时，它将一直留在那里，直到再次需要它。当使用`DB.SetMaxIdleConns`增加允许的空闲连接数量以应对突发的并行活动时，同时使用`DB.SetConnMaxIdleTime`可以安排在系统平静时释放这些连接。

#### 设置连接的最大生命周期

​	使用[DB.SetConnMaxLifetime]({{< ref "/stdLib/database/sql#db-setconnmaxlifetime----go16">}})设置连接可以在关闭之前保持打开的最长时间。

​	默认情况下，连接可以任意长时间地被使用和重复使用，受到上述限制的约束。在一些系统中，如使用负载均衡数据库服务器的系统，确保应用程序永远不会在没有重新连接的情况下使用特定的连接太久是有帮助的。

### 使用专用连接

​	`database/sql`包包括一些函数，当数据库可能对特定连接上执行的一系列操作赋予隐式含义时可以使用这些函数。

​	最常见的例子是事务，通常以`BEGIN`命令开始，以`COMMIT`或`ROLLBACK`命令结束，并在这两个命令之间的连接上包含整个事务中发出的所有命令。对于这种用例，请使用`sql`包的事务支持。请参阅[Executing transactions （执行事务）](../ExecutingTransactions)。

​	对于其他必须在同一连接上执行一系列单独操作的情况，`sql`包提供了专用连接。[DB.Conn]({{< ref "/stdLib/database/sql#db-conn----go19">}})获取一个专用连接，即sql.Conn。[sql.Conn]({{< ref "/stdLib/database/sql#type-conn----go19">}})具有类似于DB上的等效方法的方法`BeginTx`、`ExecContext`、`PingContext`、`PrepareContext`、`QueryContext`和`QueryRowContext`，但它们仅使用专用连接。完成专用连接后，代码必须使用`Conn.Close`释放它。