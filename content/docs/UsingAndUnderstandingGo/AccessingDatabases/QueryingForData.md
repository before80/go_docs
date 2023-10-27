+++
title = "查询数据"
date = 2023-05-17T15:03:14+08:00
weight = 4
description = ""
isCJKLanguage = true
draft = false
+++
# Querying for data - 查询数据

> 原文：[https://go.dev/doc/database/querying](https://go.dev/doc/database/querying)

​	当执行返回数据的SQL语句时，使用`database/sql`包中提供的`Query`方法之一。这些方法中的每个都返回一个或多个`Row`，你可以使用`Scan`方法将这些数据复制到变量中。例如，你将使用这些方法来执行`SELECT`语句。

​	当执行不返回数据的语句时，可以使用`Exec`或`ExecContext`方法代替。更多信息，请参阅[Executing statements that don’t return data（执行不返回数据的语句）](../ExecutingSQLStatementsThatDoNotReturnData)。

​	`database/sql`包提供了两种执行查询以获取结果的方法。

- **查询单行**  —— `QueryRow`最多从数据库返回一个`Row`。更多信息，请参阅[Querying for a single row （查询单行）](#查询单行)。
- 查询多行 —— `Query`将所有匹配的行作为`Rows`结构体返回，你的代码可以循环遍历它。更多信息，请参阅[查询多行](#查询多行)。

​	如果你的代码将重复执行相同的SQL语句，请考虑使用预处理语句。更多信息，请参阅[Using prepared statements （使用预处理语句）](../UsingPreparedStatements)。

> 注意：不要使用字符串格式化函数（如fmt.Sprintf）来组装SQL语句！你可能会引入SQL注入风险。更多信息，请参阅[避免SQL注入风险](https://go.dev/doc/database/sql-injection)。

### 查询单行

​	`QueryRow`检索最多一个数据库行，例如当你想通过唯一ID查找数据时。如果查询返回多个行，Scan方法将丢弃除第一个之外的所有行。

​	`QueryRowContext`的工作方式与`QueryRow`相同，但有一个`context.Context`参数。更多信息，请参阅[Canceling in-progress operations（取消正在进行的操作）](../CancelingIn-progressDatabaseOperations)。

​	以下示例使用查询来确定是否有足够的库存来支持购买。如果有足够的库存，该SQL语句返回`true`，如果没有则返回`false`。[Row.Scan]({{< ref "/stdLib/database/sql#rows-scan">}})通过指针将布尔返回值复制到`enough`变量中。

```go
func canPurchase(id int, quantity int) (bool, error) {
    var enough bool
    // Query for a value based on a single row.
    if err := db.QueryRow("SELECT (quantity >= ?) from album where id = ?",
        quantity, id).Scan(&enough); err != nil {
        if err == sql.ErrNoRows {
            return false, fmt.Errorf("canPurchase %d: unknown album", id)
        }
        return false, fmt.Errorf("canPurchase %d: %v", id)
    }
    return enough, nil
}
```

> 注意：预处理语句中的参数占位符根据你使用的`DBMS`和驱动程序而异。例如，`Postgres`的[pq driver](https://pkg.go.dev/github.com/lib/pq)需要像`$1`这样的占位符，而不是`?`。

#### 处理错误

​	`QueryRow`本身不返回错误。相反，`Scan`报告组合查找和扫描的任何错误。当查询找不到任何行时，它返回[sql.ErrNoRows]({{< ref "/stdLib/database/sql#变量">}})。  

#### Functions for returning a single row 用于返回单行的函数

| Function 函数                                                | Description **描述**                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [DB.QueryRow]({{< ref "/stdLib/database/sql#db-queryrow">}}) [DB.QueryRowContext]({{< ref "/stdLib/database/sql#db-queryrowcontext----go18">}}) | 单独运行一个单行查询。                                       |
| [Tx.QueryRow]({{< ref "/stdLib/database/sql#tx-queryrow">}}) [Tx.QueryRowContext]({{< ref "/stdLib/database/sql#tx-queryrowcontext----go18">}}) | 在较大的事务中运行一个单行查询。更多信息，请参阅[Executing transactions （执行事务）](../ExecutingTransactions) 。 |
| [Stmt.QueryRow]({{< ref "/stdLib/database/sql#stmt-queryrow">}}) [Stmt.QueryRowContext]({{< ref "/stdLib/database/sql#stmt-querycontext----go18">}}) | 使用预处理语句运行一个单行查询。更多信息，请参见 [Using prepared statements（使用预处理语句）](../UsingPreparedStatements)。 |
| [Conn.QueryRowContext]({{< ref "/stdLib/database/sql#conn-queryrowcontext----go19">}}) | 用于保留连接。更多信息，请参见[Managing connections（ 管理连接）](../ManagingConnections)。 |

### 查询多行

​	您可以使用`Query`或`QueryContext`查询多条记录，它们返回一个代表查询结果的`Rows`。您的代码使用[Rows.Next]({{< ref "/stdLib/database/sql#rows-next">}})对返回的行进行迭代。每次迭代都会调用`Scan`来将列值复制到变量中。

​	`QueryContext`的工作方式与`Query`类似，但有一个`context.Context`实参。更多信息请参见 [Canceling in-progress operations （取消正在进行的操作）](../CancelingIn-progressDatabaseOperations)。

​	下面的例子执行了一个查询，返回指定艺术家的专辑。这些专辑被返回到一个`sql.Rows`中。该代码使用[Rows.Scan](https://pkg.go.dev/database/sql#Rows.Scan)将列值复制到由指针表示的变量中。

```go 
func albumsByArtist(artist string) ([]Album, error) {
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artist)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // An album slice to hold data from returned rows.
    var albums []Album

    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist,
            &alb.Price, &alb.Quantity); err != nil {
            return albums, err
        }
        albums = append(albums, album)
    }
    if err = rows.Err(); err != nil {
        return albums, err
    }
    return albums, nil
}
```

> 注意
>
> 注意：延迟调用的[rows.Close](https://pkg.go.dev/database/sql#Rows.Close)。无论函数如何返回，这都会释放rows持有的任何资源。通过遍历所有行也会隐式地关闭它，但最好使用`defer`来确保无论发生什么情况，`rows`都会被关闭。

> 注意：预处理语句中的参数占位符根据你使用的DBMS和驱动程序而异。例如，`Postgres`的[pq driver](https://pkg.go.dev/github.com/lib/pq)需要一个类似于`$1`的占位符，而不是`?`。
>



#### 处理错误

Be sure to check for an error from `sql.Rows` after looping over query results. If the query failed, this is how your code finds out.	

​	在循环查询结果后，一定要检查`sql.Rows`的错误。如果查询失败，这就是你的代码如何发现的方式。

#### 返回多行记录的函数

| Function 函数                    | Description 描述                                             |
| -------------------------------- | ------------------------------------------------------------ |
| [DB.Query]({{< ref "/stdLib/database/sql#db-query">}}) [DB.QueryContext]({{< ref "/stdLib/database/sql#db-querycontext----go18">}}) | 单独运行一个查询。                                           |
| [Tx.Query]({{< ref "/stdLib/database/sql#tx-query">}}) [Tx.QueryContext]({{< ref "/stdLib/database/sql#tx-querycontext----go18">}}) | 在较大的事务中运行一个查询。更多信息，请参阅[Executing transactions （执行事务）](../ExecutingTransactions) 。 |
| [Stmt.Query]({{< ref "stdLib/database/sql#stmt-query">}}) [Stmt.QueryContext]({{< ref "/stdLib/database/sql#stmt-querycontext----go18">}}) | 使用预处理语句运行一个查询。更多信息，请参见[Using prepared statements（使用预处理语句）](../UsingPreparedStatements)。 |
| [Conn.QueryContext]({{< ref "/stdLib/database/sql#conn-querycontext----go19">}}) | 用于保留连接。更多信息，请参见[Managing connections（ 管理连接）](../ManagingConnections)。 |

### 处理可为null的列值

​	`database/sql`包提供了几种特殊类型，你可以在`Scan`方法中使用它们作为参数，当列的值可能为`null`时。每个类型都包含一个`Valid`字段，报告值是否非`null`，以及一个字段（如果值为空），则持有该值。

​	以下示例中的代码查询客户名称。如果名称值为`null`，则代码将另一个值替换为应用程序中使用的值。

```go 
var s sql.NullString
err := db.QueryRow("SELECT name FROM customer WHERE id = ?", id).Scan(&s)
if err != nil {
    log.Fatal(err)
}

// Find customer name, using placeholder if not present.
name := "Valued Customer"
if s.Valid {
    name = s.String
}
```

​	在`sql`包参考中可以查看每种类型的更多信息：

- [NullBool]({{< ref "/stdLib/database/sql#type-nullbool">}})
- [NullFloat64]({{< ref "/stdLib/database/sql#type-nullfloat64">}})
- [NullInt32]({{< ref "/stdLib/database/sql#type-nullint32----go113">}})
- [NullInt64]({{< ref "/stdLib/database/sql#type-nullint64">}})
- [NullString]({{< ref "/stdLib/database/sql#type-nullstring">}})
- [NullTime]({{< ref "/stdLib/database/sql#type-nulltime----go113">}})

### 从列中获取数据

​	当遍历查询返回的行时，你可以使用`Scan`将一行的列值复制到Go值中，如[Rows.Scan]({{< ref "/stdLib/database/sql#rows-scan">}})参考中所述。

​	所有驱动程序都支持一组基本的数据转换，例如将SQL `INT`转换为Go `int`。一些驱动程序扩展了这组转换；有关详细信息，请参阅每个驱动程序的文档。

​	正如你可能期望的那样，`Scan`将从与Go类型相似的列类型进行转换。例如，`Scan`将从SQL `CHAR`、`VARCHAR`和`TEXT`转换为Go `string`。然而，`Scan`还将执行到另一个适合列值的Go类型的转换。例如，如果列始终包含一个数字的`VARCHAR`，你可以指定一个数值型的Go类型（如`int`）来接收该值，然后`Scan`将使用`strconv.Atoi`为你进行转换。

​	有关`Scan`方法进行的转换的更多详细信息，请参阅[Rows.Scan]({{< ref "/stdLib/database/sql#rows-scan">}})参考。

### 处理多个结果集

​	当你的数据库操作可能返回多个结果集时，你可以使用[Rows.NextResultSet]({{< ref "/stdLib/database/sql#rows-nextresultset----go18">}})来检索它们。例如，当你分别查询多个表并返回每个表的结果集时，这可能会很有用。

​	`Rows.NextResultSet`准备下一个结果集，以便对`Rows.Next`的调用可以从下一个集中检索第一行。它返回一个布尔值，表示是否确实存在下一个结果集。

​	以下示例中的代码使用`DB.Query`执行两个SQL语句。第一个结果集来自存储过程的第一个查询，检索了`album`表中的所有行。下一个结果集来自第二个查询，从`song`表中检索行。

```go
rows, err := db.Query("SELECT * from album; SELECT * from song;")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

// Loop through the first result set.
for rows.Next() {
    // Handle result set.
}

// Advance to next result set.
rows.NextResultSet()

// Loop through the second result set.
for rows.Next() {
    // Handle second set.
}

// Check for any error in either result set.
if err := rows.Err(); err != nil {
    log.Fatal(err)
}
```