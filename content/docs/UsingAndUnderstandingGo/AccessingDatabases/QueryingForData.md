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

​	当执行一个返回数据的SQL语句时，使用`database/sql`包中提供的`Query`方法之一。每个方法都会返回一个行（`Row`）或多个行（`Rows`），你可以使用`Scan`方法将其数据复制到变量。你会使用这些方法，例如，执行`SELECT`语句。

​	当执行一个不返回数据的语句时，你可以改用`Exec`或`ExecContext`方法。更多信息请参见 [Executing statements that don’t return data（执行不返回数据的语句）](../ExecutingSQLStatementsThatDoNotReturnData)。

​	`database/sql`包提供了两种执行结果查询的方法：

- **查询单行**  —— `QueryRow` 最多只能从数据库中返回一个单行。更多信息请参见 [Querying for a single row （查询单行）](#查询单行)。
- 查询多行 —— `Query` 将所有匹配的行作为一个`Rows`结构体（你的代码可以循环遍历）返回，。更多信息，请参见 [查询多行](#查询多行)。

​	如果你的代码将重复执行相同的SQL语句，请考虑使用预处理语句。更多信息，请参见 [Using prepared statements （使用预处理语句）](../UsingPreparedStatements) 。

!!! warning "注意"

	注意：不要使用字符串格式化函数，如`fmt.Sprintf`来组合一个SQL语句！你可能会引入一个SQL注入的风险。更多信息，请参见避免[SQL注入风险](https://go.dev/doc/database/sql-injection)。

### 查询单行

​	`QueryRow`最多只能检索一条数据库记录，例如当你想通过一个唯一的ID来查询数据。如果查询返回多条记录，`Scan`方法会丢弃除第一条以外的所有记录。

​	`QueryRowContext`的工作方式与`QueryRow`类似，但有一个`context.Context`实参。更多信息请参见 [Canceling in-progress operations（取消正在进行的操作）](../CancelingIn-progressDatabaseOperations)。

​	下面的例子使用一个查询来找出是否有足够的库存来支持购买。如果有足够的库存，该SQL语句返回`true`，如果没有则返回`false`。[Row.Scan]({{< ref "/docs/StdLib/database/sql#rows-scan">}})通过一个指针将布尔型的返回值复制到`enough`变量中。

```go  hl_lines="5 5"
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

注意：准备预处理语句中的参数占位符根据你所使用的`DBMS`和驱动而不同。例如，`Postgres`的[pq driver](https://pkg.go.dev/github.com/lib/pq)需要一个类似于`$1`的占位符，而不是`?`。

#### 处理错误

​	`QueryRow`本身不返回错误。相反，`Scan`报告来自组合查询和扫描的任何错误。当查询没有找到记录时，它返回[sql.ErrNoRows]({{< ref "/docs/StdLib/database/sql#变量">}})。

#### Functions for returning a single row 用于返回单行的函数

| Function 函数                          | Description **描述**                                         |
| -------------------------------------- | ------------------------------------------------------------ |
| `DB.QueryRow` `DB.QueryRowContext`     | 单独运行一个单行查询。                                       |
| `Tx.QueryRow` `Tx.QueryRowContext`     | 在较大的事务中运行一个单行查询。更多信息，请参阅[Executing transactions （执行事务）](../ExecutingTransactions) 。 |
| `Stmt.QueryRow` `Stmt.QueryRowContext` | 使用预处理语句运行一个单行查询。更多信息，请参见 [Using prepared statements（使用预处理语句）](../UsingPreparedStatements)。 |
| `Conn.QueryRowContext`                 | 用于保留连接。更多信息，请参见[Managing connections（ 管理连接）](../ManagingConnections)。 |

### 查询多行

​	你可以使用`Query`或`QueryContext`查询多条记录，它们返回一个代表查询结果的`Rows`。你的代码使用[Rows.Next]({{< ref "/docs/StdLib/database/sql#rows-next">}})对返回的行进行迭代。每次迭代都会调用`Scan`来将列值复制到变量中。

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

!!! warning "注意"

	注意对[rows.Close](https://pkg.go.dev/database/sql#Rows.Close)的延迟调用。无论函数如何返回，这都会释放rows所持有的任何资源。循环处理所有的行也会隐式地关闭它，但最好使用`defer`来确保无论如何都会关闭`rows`。

!!! warning "注意"

	注意：预处理语句中的参数占位符根据你所使用的`DBMS`和驱动而不同。例如，`Postgres`的[pq driver](https://pkg.go.dev/github.com/lib/pq)需要一个类似于`$1`的占位符，而不是`?`。



#### 处理错误

Be sure to check for an error from `sql.Rows` after looping over query results. If the query failed, this is how your code finds out.

​	在循环查询结果后，一定要从`sql.Rows`中检查是否有错误。如果查询失败了，代码就是这样查找的。

#### 返回多行记录的函数

| Function 函数                    | Description 描述                                             |
| -------------------------------- | ------------------------------------------------------------ |
| `DB.Query` `DB.QueryContext`     | 单独运行一个查询。                                           |
| `Tx.Query` `Tx.QueryContext`     | 在较大的事务中运行一个查询。更多信息，请参阅[Executing transactions （执行事务）](../ExecutingTransactions) 。 |
| `Stmt.Query` `Stmt.QueryContext` | 使用预处理语句运行一个查询。更多信息，请参见[Using prepared statements（使用预处理语句）](../UsingPreparedStatements)。 |
| `Conn.QueryContext`              | 用于保留连接。更多信息，请参见[Managing connections（ 管理连接）](../ManagingConnections)。 |

### 处理可为null的列值

​	`database/sql`包提供了几种特殊的类型，当一个列的值可能为`null`时，你可以作为`Scan`函数的实参使用。每种类型都包括一个`Valid`字段，用于报告值是否为非`null`，如果是的话，还包括一个持有该值的字段。

​	下面的例子中的代码查询了一个客户名称。如果名字的值是`null`的，代码会替换另一个值在应用程序中使用。

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

在`sql`包参考资料中可以看到更多关于每种类型的信息：

- [`NullBool`](https://pkg.go.dev/database/sql#NullBool)
- [`NullFloat64`](https://pkg.go.dev/database/sql#NullFloat64)
- [`NullInt32`](https://pkg.go.dev/database/sql#NullInt32)
- [`NullInt64`](https://pkg.go.dev/database/sql#NullInt64)
- [`NullString`](https://pkg.go.dev/database/sql#NullString)
- [`NullTime`](https://pkg.go.dev/database/sql#NullTime)

### 从列中获取数据

​	在循环查询返回的行时，您可以使用`Scan`将行的列值复制到Go值，如[Rows.Scan]({{< ref "/docs/StdLib/database/sql#rows-scan">}})参考中所述。

​	所有驱动程序都支持一组基本的数据转换，例如将SQL `INT`转换为Go `int`。一些驱动程序扩展了这一转换集；详情请参见各个驱动程序的文档。

​	正如你所期望的，`Scan`将从列类型转换为类似的Go类型。例如，`Scan`将从SQL `CHAR`、`VARCHAR`和`TEXT`转换为Go `string`。但是，`Scan`也会执行转换为另一种适合列值的Go类型。例如，如果列是一个总是包含数字的`VARCHAR`，你可以指定一个数值Go类型，比如`int`，来接收这个值，`Scan`将使用`strconv.Atoi`对其进行转换。

​	关于`Scan`函数进行转换的更多细节，请参见[Rows.Scan]({{< ref "/docs/StdLib/database/sql#rows-scan">}})参考。

### 处理多个结果集

​	当你的数据库操作可能返回多个结果集时，你可以通过使用[Rows.NextResultSet]({{< ref "/docs/StdLib/database/sql#rows-nextresultset----go18">}})来检索这些结果。这可能很有用，例如，当你发送分别查询多个表的 SQL 时，为每个表返回一个结果集。

​	`Rows.NextResultSet`准备好下一个结果集，以便调用`Rows.Next`检索下一个结果集的第一条记录。它返回一个布尔值，表明是否存在下一个结果集。

​	下面的例子中的代码使用`DB.Query`来执行两个SQL语句。第一个结果集来自过程中的第一个查询，检索`album`表中的所有记录。下一个结果集是来自于第二个查询，从`song`表中检索记录。

```go  hl_lines="13 13"
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