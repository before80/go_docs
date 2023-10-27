+++
title = "执行不返回数据的SQL语句"
date = 2023-05-17T15:03:14+08:00
weight = 3
description = ""
isCJKLanguage = true
draft = false
+++
# Executing SQL statements that don't return data - 执行不返回数据的SQL语句

> 原文：[https://go.dev/doc/database/change-data](https://go.dev/doc/database/change-data)

​	当你执行不返回数据的数据库操作时，使用`database/sql`包中的`Exec`或`ExecContext`方法。你可以通过这种方式执行的SQL语句包括`INSERT`、`DELETE`和`UPDATE`。

​	当你的查询可能会返回行时，请改用`Query`或`QueryContext`方法。更多信息，请参阅[Querying a database （查询数据库）](../QueryingForData)。

​	`ExecContext`方法与`Exec`方法的工作方式相同，但有一个额外的`context.Context`实参，如[Canceling in-progress operations （取消正在进行的操作）](../CancelingIn-progressDatabaseOperations)中所述。

​	以下示例代码使用[DB.Exec]({{< ref "/stdLib/database/sql#db-exec">}})执行一条语句，将一张新的专辑添加到`album`（专辑）表中。

```go 
func AddAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist) VALUES (?, ?)", alb.Title, alb.Artist)
    if err != nil {
        return 0, fmt.Errorf("AddAlbum: %v", err)
    }

    // Get the new album's generated ID for the client.
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("AddAlbum: %v", err)
    }
    // Return the new album's ID.
    return id, nil
}
```

​	`DB.Exec`返回值：一个[sql.Result]({{< ref "/stdLib/database/sql#type-result">}})和一个错误。当错误为`nil`时，你可以使用`Result`获取最后一个插入项的ID（如示例所示）或检索操作影响的行数。

> 注意：参数占位符在预处理语句中会因你使用的`DBMS`和驱动程序而异。例如，`Postgres`的[pq driver](https://pkg.go.dev/github.com/lib/pq)需要像`$1`这样的占位符，而不是`?`。

​	如果你的代码将重复执行相同的SQL语句，请考虑使用`sql.Stmt`从SQL语句创建一个可重用的预处理语句。更多信息，请参阅[Using prepared statements（使用预处理语句）](../UsingPreparedStatements)。

> 警告：不要使用字符串格式化函数（如`fmt.Sprintf`）来组装SQL语句！你可能会引入[SQL注入的风险](https://go.dev/doc/database/sql-injection)。更多信息，请参阅[Avoiding SQL injection risk（避免SQL注入风险）](../AvoidingSQLInjectionRisk)。

#### 用于执行不返回行的SQL语句的函数

| Function 函数                                                | Description 描述                                             |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [DB.Exec]({{< ref "/stdLib/database/sql#db-exec">}}) [DB.ExecContext]({{< ref "/stdLib/database/sql#db-execcontext----go18">}}) | 单独执行一条 SQL 语句。                                      |
| [Tx.Exec]({{< ref "/stdLib/database/sql#tx-exec">}}) [Tx.ExecContext]({{< ref "/stdLib/database/sql#tx-execcontext----go18">}}) | 在较大的事务中执行 SQL 语句。有关详细信息，请参阅请参阅[Executing transactions （执行事务）](../ExecutingTransactions) 。 |
| [Stmt.Exec]({{< ref "/stdLib/database/sql#stmt-exec">}}) [Stmt.ExecContext]({{< ref "/stdLib/database/sql#stmt-execcontext----go18">}}) | 执行一个预处理SQL语句。更多信息，请参见 [Using prepared statements（使用预处理语句）](../UsingPreparedStatements)。 |
| [Conn.ExecContext]({{< ref "/stdLib/database/sql#conn-execcontext----go19">}}) | 用于保留连接。更多信息，请参见[Managing connections（ 管理连接）](../ManagingConnections)。 |