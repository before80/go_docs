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

​	你可以为重复使用定义一个预处理语句。这可以帮助你的代码运行得更快，避免每次执行数据库操作时重新创建语句的开销。

> 注意：预处理语句中的参数占位符根据你使用的`DBMS`和驱动程序而异。例如，Postgres的[pq driver](https://pkg.go.dev/github.com/lib/pq)需要一个像`$1`这样的占位符，而不是`?`。

###  什么是预处理语句？

​	预处理语句是由`DBMS`解析和保存的SQL，通常包含占位符，但没有实际的参数值。稍后，可以使用一组参数值执行该语句。

### 如何使用预处理语句？

​	当您希望重复执行相同的SQL时，您可以使用一个`sql.Stmt`来提前准备SQL语句，然后根据需要执行它。

​	下面的例子创建了一个预处理语句，从数据库中选择一个特定的相册。[DB.Prepare]({{< ref "/stdLib/database/sql#db-prepare">}})返回一个`sql.Stmt`，代表一个给定的SQL文本的预处理语句。您可以将SQL语句的参数传递给`Stmt.Exec`、`Stmt.QueryRow`或`Stmt.Query`来运行该语句。

​	当你期望重复执行相同的SQL时，可以使用`sql.Stmt`预先准备SQL语句，然后根据需要执行它。

​	以下示例创建一个从数据库中选择特定专辑的预处理语句。[DB.Prepare]({{< ref "/stdLib/database/sql#db-prepare">}})返回一个表示给定SQL文本的预处理语句的`sql.Stmt`。你可以将SQL语句的参数传递给`Stmt.Exec`、`Stmt.QueryRow`或`Stmt.Query`来运行该语句。

```go  hl_lines="14 14"
// AlbumByID retrieves the specified album.
// AlbumByID检索指定的专辑。
func AlbumByID(id int) (Album, error) {
    // Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    // 定义一个预处理语句。你通常会在其他地方定义这个语句，并将其保存起来以供像这样的函数使用。
    stmt, err := db.Prepare("SELECT * FROM album WHERE id = ?")
    if err != nil {
        log.Fatal(err)
    }

    var album Album

    // Execute the prepared statement, passing in an id value for the
    // parameter whose placeholder is ?
    // 执行预处理语句，为占位符为?的参数传递一个id值。
    err := stmt.QueryRow(id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.Quantity)
    if err != nil {
        if err == sql.ErrNoRows {
            // Handle the case of no rows returned.
            // 处理没有返回行的情形。
        }
        return album, err
    }
    return album, nil
}
```

### 预处理语句的行为

​	一个准备好的[sql.Stmt]({{< ref "/stdLib/database/sql#type-stmt">}})提供了常用的`Exec`、`QueryRow`和`Query`方法来调用其语句。关于使用这些方法的更多信息，请参阅[Querying for data （查询数据）](../QueryingForData) 和 [Executing SQL statements that don’t return data（执行不返回数据的SQL语句）](../ExecutingSQLStatementsThatDoNotReturnData)。

​	然而，由于一个`sql.Stmt`已经代表了一个预设的SQL语句，它的`Exec`、`QueryRow`和`Query`方法只接受与占位符对应的SQL参数值，而忽略了SQL文本。

​	您可以用不同的方式定义一个新的`sql.Stmt`，这取决于您将如何使用它。

- `DB.Prepare`和`DB.PrepareContext`创建了一个预处理语句，该语句可以在事务外单独执行，就像`DB.Exec`和`DB.Query`一样。
- `Tx.Prepare`、`Tx.PrepareContext`、`Tx.Stmt`和`Tx.StmtContext`创建一个预处理语句，以便在一个特定的事务中使用。`Prepare`和`PrepareContext`使用SQL文本来定义语句。`Stmt`和`StmtContext`使用`DB.Prepare`或`DB.PrepareContext`的结果。也就是说，它们将一个非事务用的`sql.Stmt`转换为这个事务用的`sql.Stmt`。
- `Conn.PrepareContext`从一个代表保留连接的`sql.Conn`创建一个预处理语句。

​	一定要记住，在代码使用语句完成时，调用`stmt.Close`。这将释放任何可能与之相关的数据库资源（如底层连接）。对于只是一个函数中的局部变量的语句，`defer stmt.Close()`就足够了。



​	一个预处理的[sql.Stmt]({{< ref "/stdLib/database/sql#type-stmt">}})提供了通常的`Exec`、`QueryRow`和`Query`方法来调用该语句。有关使用这些方法的更多信息，请参阅[Querying for data （查询数据）](../QueryingForData)和[Executing SQL statements that don’t return data（执行不返回数据的SQL语句）](../ExecutingSQLStatementsThatDoNotReturnData)。

​	然而，由于`sql.Stmt`已经表示了一个预设的SQL语句，它的`Exec`、`QueryRow`和`Query`方法只接受与占位符对应的SQL参数值，省略了SQL文本。

​	你可以根据如何使用它来以不同的方式定义一个新的`sql.Stmt`。

- `DB.Prepare`和`DB.PrepareContext`创建一个可以独立执行的预处理语句，就像`DB.Exec`和`DB.Query`一样，可以在事务之外单独使用。 

- `Tx.Prepare`、`Tx.PrepareContext`、`Tx.Stmt`和`Tx.StmtContext`创建一个特定事务中使用的预处理语句。`Prepare`和`PrepareContext`使用SQL文本定义语句。`Stmt`和`StmtContext`使用`DB.Prepare`或`DB.PrepareContext`的结果。也就是说，它们将一个非事务性的`sql.Stmt`转换为一个针对此事务的`sql.Stmt`。 

- `Conn.PrepareContext`从一个表示保留连接的`sql.Conn`中创建一个预处理语句。 

  

​	确保在代码完成对一个语句的处理时调用`stmt.Close`。这将释放可能与其关联的任何数据库资源（如底层连接）。对于仅作为函数局部变量的语句，只需`defer stmt.Close()`就足够了。

#### 创建预处理语句的函数

| Function 函数                                                | Description 描述                                             |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [DB.Prepare]({{< ref "/stdLib/database/sql#">}}) [DB.PrepareContext]({{< ref "/stdLib/database/sql#db-preparecontext----go18">}}) | 准备独立执行的语句，或者使用 `Tx.Stmt` 将其转换为事务内准备的语句。 |
| [Tx.Prepare]({{< ref "/stdLib/database/sql#tx-prepare">}}) [Tx.PrepareContext]({{< ref "/stdLib/database/sql#tx-preparecontext----go18">}}) [Tx.Stmt]({{< ref "/stdLib/database/sql#tx-stmt">}}) [Tx.StmtContext]({{< ref "/stdLib/database/sql#tx-stmtcontext----go18">}}) | 准备一个在特定事务中使用的语句。更多信息，请参阅[Executing transactions （执行事务）](../ExecutingTransactions) 。 |
| [Conn.PrepareContext]({{< ref "/stdLib/database/sql#conn-preparecontext----go19">}}) | 用于保留连接。更多信息，请参见[Managing connections（ 管理连接）](../ManagingConnections)。 |