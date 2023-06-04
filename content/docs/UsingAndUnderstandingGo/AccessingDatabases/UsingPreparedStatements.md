+++
title = "使用预处理语句"
date = 2023-05-17T15:03:14+08:00
weight = 5
description = ""
isCJKLanguage = true
draft = false
+++
# Using prepared statements - 使用预处理语句

> 原文：[https://go.dev/doc/database/prepared-statements](https://go.dev/doc/database/prepared-statements)

​	您可以为重复使用定义一个预处理语句。这可以避免每次代码执行数据库操作时`重新创建语句的开销`，从而帮助代码更快地运行。

注意：预处理语句中的参数占位符根据您所使用的`DBMS`和驱动而不同。例如，`Postgres`的[pq driver](https://pkg.go.dev/github.com/lib/pq)需要一个像`$1`这样的占位符，而不是`?`。

###  什么是预处理语句？

​	预处理语句是由`DBMS`解析并保存的SQL，通常包含占位符，但没有实际参数值。之后，可以用一组参数值来执行该语句。

### 您如何使用预处理语句

​	当您希望重复执行相同的SQL时，您可以使用一个`sql.Stmt`来提前准备SQL语句，然后根据需要执行它。

​	下面的例子创建了一个预处理语句，从数据库中选择一个特定的相册。[DB.Prepare]({{< ref "/docs/StdLib/database/sql#db-prepare">}})返回一个`sql.Stmt`，代表一个给定的SQL文本的预处理语句。您可以将SQL语句的参数传递给`Stmt.Exec`、`Stmt.QueryRow`或`Stmt.Query`来运行该语句。

```go  hl_lines="14 14"
// AlbumByID retrieves the specified album.
func AlbumByID(id int) (Album, error) {
    // Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    stmt, err := db.Prepare("SELECT * FROM album WHERE id = ?")
    if err != nil {
        log.Fatal(err)
    }

    var album Album

    // Execute the prepared statement, passing in an id value for the
    // parameter whose placeholder is ?
    err := stmt.QueryRow(id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.Quantity)
    if err != nil {
        if err == sql.ErrNoRows {
            // Handle the case of no rows returned.
        }
        return album, err
    }
    return album, nil
}
```

### 预处理语句的行为

​	一个准备好的[sql.Stmt]({{< ref "/docs/StdLib/database/sql#type-stmt">}})提供了常用的`Exec`、`QueryRow`和`Query`方法来调用其语句。关于使用这些方法的更多信息，请参阅[Querying for data （查询数据）](../QueryingForData) 和 [Executing SQL statements that don’t return data（执行不返回数据的SQL语句）](../ExecutingSQLStatementsThatDoNotReturnData)。

​	然而，由于一个`sql.Stmt`已经代表了一个预设的SQL语句，它的`Exec`、`QueryRow`和`Query`方法只接受与占位符对应的SQL参数值，而忽略了SQL文本。

​	您可以用不同的方式定义一个新的`sql.Stmt`，这取决于您将如何使用它。

- `DB.Prepare`和`DB.PrepareContext`创建了一个预处理语句，该语句可以在事务外单独执行，就像`DB.Exec`和`DB.Query`一样。
- `Tx.Prepare`、`Tx.PrepareContext`、`Tx.Stmt`和`Tx.StmtContext`创建一个预处理语句，以便在一个特定的事务中使用。`Prepare`和`PrepareContext`使用SQL文本来定义语句。`Stmt`和`StmtContext`使用`DB.Prepare`或`DB.PrepareContext`的结果。也就是说，它们将一个非事务用的`sql.Stmt`转换为这个事务用的`sql.Stmt`。
- `Conn.PrepareContext`从一个代表保留连接的`sql.Conn`创建一个预处理语句。

​	一定要记住，在代码使用语句完成时，调用`stmt.Close`。这将释放任何可能与之相关的数据库资源（如底层连接）。对于只是一个函数中的局部变量的语句，`defer stmt.Close()`就足够了。

#### 创建预处理语句的函数

| Function 函数                                               | Description 描述                                             |
| ----------------------------------------------------------- | ------------------------------------------------------------ |
| `DB.Prepare` `DB.PrepareContext`                            | 准备独立执行的语句，或者使用 `Tx.Stmt` 将其转换为事务内准备的语句。 |
| `Tx.Prepare` `Tx.PrepareContext` `Tx.Stmt` `Tx.StmtContext` | 准备一个在特定事务中使用的语句。更多信息，请参阅[Executing transactions （执行事务）](../ExecutingTransactions) 。 |
| `Conn.PrepareContext`                                       | 用于保留连接。更多信息，请参见[Managing connections（ 管理连接）](../ManagingConnections)。 |