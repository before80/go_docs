+++
title = "访问关系型数据库"
date = 2023-05-17T15:03:14+08:00
weight = 1
description = ""
isCJKLanguage = true
draft = false
+++
# Accessing relational databases - 访问关系型数据库

> 原文：[https://go.dev/doc/database/](https://go.dev/doc/database/)

​	使用 Go，您可以将各种数据库和数据访问方法纳入您的应用程序中。本节的主题描述了如何使用标准库的[database/sql]({{< ref "/docs/StdLib/database/sql" >}})包来访问关系型数据库。

​	关于Go的数据访问的介绍性教程，请看 [Tutorial: Accessing a relational database](../../../GettingStarted/TutorialAccessingARelationalDatabase)。

​	Go 也支持其他数据访问技术，包括用于更高级别的访问关系型数据库的 ORM 库，以及非关系型 NoSQL 数据存储。

- **对象-关系映射（ORM）库（Object-relational mapping (ORM) libraries）**。虽然`database/sql` 包包括用于低级数据访问逻辑的函数，但您也可以使用 Go 来访问更高抽象级别的数据存储。关于Go的两个流行的对象-关系映射（ORM）库的更多信息，请参见[GORM](https://gorm.io/index.html)和[ent](https://entgo.io/) ([package reference](https://pkg.go.dev/entgo.io/ent))。
- **NoSQL 数据存储.** Go 社区已经为大多数 NoSQL 数据存储开发了驱动程序，包括 [MongoDB](https://docs.mongodb.com/drivers/go/)和 [Couchbase](https://docs.couchbase.com/go-sdk/current/hello-world/overview.html)。您可以搜索[pkg.go.dev](https://pkg.go.dev/) 获取更多信息。

### 支持的数据库管理系统

xxxxxxxxxx1 1SELECT * FROM user WHERE id = 1 OR 1=1;mysql

​	你可以在 [SQLDrivers](https://github.com/golang/go/wiki/SQLDrivers)页面找到完整的驱动列表。

### 执行查询或更改数据库的函数

​	`database/sql`包包含专门为你正在执行的数据库操作设计的函数。例如，虽然你可以使用`Query`或`QueryRow`来执行查询，但`QueryRow`是为只需要一行的情况而设计的，省去了返回只包括一行记录的`sql.Rows`的开销。你可以使用`Exec`函数用SQL语句对数据库进行修改，如`INSERT`, `UPDATE`, 或`DELETE`。

​	更多内容，请参见以下内容：

- [Executing SQL statements that don’t return data （执行不返回数据的SQL语句）](../ExecutingSQLStatementsThatDoNotReturnData)
- [Querying for data （查询数据）](../QueryingForData)

### 事务

​	通过`sql.Tx`，您可以编写代码来执行事务中的数据库操作。在一个事务中，多个操作可以一起执行，并以最后的提交（commit）来结束，以便在一个原子步骤中应用所有的更改，或者以回滚（rollback）来丢弃（discard ）它们。

​	关于事务的更多信息，请参见[Executing transactions （执行事务）](../ExecutingTransactions)。

### 查询的取消

​	当你希望能够取消一个数据库操作时，你可以使用`context.Context`，例如，当客户端的连接关闭或操作运行的时间超过期望时。

​	对于任何数据库操作，你可以使用一个`database/sql`包函数，该函数将`Context`作为一个实参。使用`Context`，你可以为操作指定一个超时或最后期限。你还可以使用 `Context` 将取消请求通过应用程序传播到执行 SQL 语句的函数，确保在不再需要资源时释放资源。

​	更多信息请参见[Canceling in-progress operations （取消正在进行的操作）](../CancelingIn-progressDatabaseOperations)。

### 管理连接池

​	当你使用`sql.DB`数据库句柄时，你正在与一个内置的连接池连接，该连接池根据你代码的需要创建和处置连接。通过`sql.DB`的句柄是用Go进行数据库访问的最常见方式。更多信息请参见[Opening a database handle （打开数据库句柄）](../OpeningADatabaseHandle) 。

​	`database/sql`包为你管理连接池。然而，对于更高级的需求，可以按照[Setting connection pool properties （设置连接池属性）](../ManagingConnections#设置连接池属性)中的说明设置连接池属性。

​	对于那些需要单一保留连接的操作，`database/sql`包提供了[sql.Conn](https://pkg.go.dev/database/sql#Conn)。当使用`sql.Tx`的事务是一个糟糕的选择时，`Conn`就特别有用。

​	例如，你的代码可能需要：

- 通过`DDL`进行模式更改，包括包含其自身事务语义的逻辑。将`sql`包的事务函数与SQL事务语句混合在一起是一种不好的做法，正如在[Executing transactions （执行事务）](../ExecutingTransactions) 中所描述的那样。
- 执行创建临时表的查询锁定操作。



​	更多内容请参见 [Using dedicated connections （使用专用连接）](../ManagingConnections#使用专用连接)。

