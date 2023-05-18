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

​	当你执行不返回数据的数据库操作时，请使用`database/sql`包的`Exec`或`ExecContext`方法。以这种方式执行的SQL语句包括`INSERT`, `DELETE`, 和`UPDATE`。

​	当你的查询可能返回行时，请改为使用`Query`或`QueryContext`方法。更多信息请参见 [Querying a database （查询数据库）](../QueryingForData)。

​	`ExecContext`方法的工作原理与`Exec`方法相同，但有一个额外的`context.Context`参数，如 [Canceling in-progress operations （取消正在进行的操作）](../CancelingIn-progressDatabaseOperations) 中所述。

​	下面的例子中的代码使用[DB.Exec](https://pkg.go.dev/database/sql#DB.Exec)来执行一条语句，将一个新的记录专辑添加到一个`album`表中。

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

​	`DB.Exec`返回值：一个[sql.Result](https://pkg.go.dev/database/sql#Result)和一个错误。当错误为`nil`时，你可以使用`Result`来获得最后插入的项目的ID（如在例子中）或检索受操作影响的行数。

!!! warning "注意"

	注意：预处理语句中的参数占位符根据你所使用的`DBMS`和驱动而不同。例如，`Postgres`的[pq driver](https://pkg.go.dev/github.com/lib/pq)需要一个类似于`$1`的占位符，而不是`?`.

​	如果你的代码将重复执行相同的SQL语句，考虑使用一个`sql.Stmt`来从SQL语句中创建一个可重复使用的预处理语句。更多信息请参见 [Using prepared statements（使用预处理语句）](../UsingPreparedStatements)。

!!! warning "注意"

	注意：不要使用字符串格式化函数，如`fmt.Sprintf`来组合一个SQL语句！ 你可能会引入一个[SQL注入的风险](https://go.dev/doc/database/sql-injection)。更多信息，请参见[Avoiding SQL injection risk（避免SQL注入风险）](../AvoidingSQLInjectionRisk)。

#### Functions for executing SQL statements that don’t return rows 用于执行不返回行的SQL语句的函数

| Function 函数                  | Description 描述                                             |
| ------------------------------ | ------------------------------------------------------------ |
| `DB.Exec` `DB.ExecContext`     | 单独执行一条 SQL 语句。                                      |
| `Tx.Exec` `Tx.ExecContext`     | 在较大的事务中执行 SQL 语句。有关详细信息，请参阅请参阅[Executing transactions （执行事务）](../ExecutingTransactions) 。 |
| `Stmt.Exec` `Stmt.ExecContext` | 执行一个预处理SQL语句。更多信息，请参见 [Using prepared statements（使用预处理语句）](../UsingPreparedStatements)。 |
| `Conn.ExecContext`             | 用于保留连接。更多信息，请参见[Managing connections（ 管理连接）](../ManagingConnections)。 |