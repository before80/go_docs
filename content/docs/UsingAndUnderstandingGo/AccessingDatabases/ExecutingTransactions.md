+++
title = "执行事务"
date = 2023-05-17T15:03:14+08:00
weight = 6
description = ""
isCJKLanguage = true
draft = false
+++
# Executing transactions - 执行事务

> 原文：[https://go.dev/doc/database/execute-transactions](https://go.dev/doc/database/execute-transactions)

​	你可以使用表示事务的[sql.Tx](https://pkg.go.dev/database/sql#Tx)来执行数据库事务。除了表示特定于事务语义的`Commit`和`Rollback`方法之外，`sql.Tx`还有所有用来执行普通数据库操作的方法。要获取`sql.Tx`，可以调用`DB.Begin`或`DB.BeginTx`。

​	一个[数据库事务](https://en.wikipedia.org/wiki/Database_transaction)将多个操作作为更大目标的一部分进行分组。所有的操作都必须成功，或者都不能成功，在这两种情况下都要保持数据的完整性。通常，一个事务的工作流程包括：

1. 开始事务。
2. 执行一系列的数据库操作。
3. 如果没有发生错误，提交事务以进行数据库更改。
4. 如果发生错误，回滚事务使数据库保持不变。

​	`sql`包提供了开始和结束事务的方法，以及执行中间的数据库操作的方法。这些方法与上述工作流程中的四个步骤相对应。

- 开始一个事务

  [DB.Begin](https://pkg.go.dev/database/sql#DB.Begin)或[DB.BeginTx](https://pkg.go.dev/database/sql#DB.BeginTx)开始一个新的数据库事务，返回一个代表它的`sql.Tx`。

- 执行数据库操作。

  使用一个`sql.Tx`，你可以在一系列使用单一连接的操作中查询或更新数据库。为了支持这一点，`Tx`导出了以下方法：

  - [Exec](https://pkg.go.dev/database/sql#Tx.Exec) 和 [ExecContext](https://pkg.go.dev/database/sql#Tx.ExecContext) ，用于通过SQL语句（如`INSERT`、`UPDATE`和`DELETE`）进行数据库更改。

    更多信息请参见[Executing statements that don’t return data（执行不返回数据的语句）](../ExecutingSQLStatementsThatDoNotReturnData)。

  - [Query](https://pkg.go.dev/database/sql#Tx.Query)，[QueryContext](https://pkg.go.dev/database/sql#Tx.QueryContext)，[QueryRow](https://pkg.go.dev/database/sql#Tx.QueryRow) ，[QueryRowContext](https://pkg.go.dev/database/sql#Tx.QueryRowContext) 用于返回行的操作。

    更多信息，请参见 [Querying for data （查询数据）](../QueryingForData)。

  - [Prepare](https://pkg.go.dev/database/sql#Tx.Prepare)，[PrepareContext](https://pkg.go.dev/database/sql#Tx.PrepareContext)，[Stmt](https://pkg.go.dev/database/sql#Tx.Stmt)，[StmtContext](https://pkg.go.dev/database/sql#Tx.StmtContext)用于预先定义预处理语句。

    更多信息，请参见[Using prepared statements （使用预处理语句）](../UsingPreparedStatements)。

- 用以下方法之一结束事务：

  - 使用[Tx.Commit](https://pkg.go.dev/database/sql#Tx.Commit)提交事务。

    如果`Commit`成功（返回`nil`错误），那么所有的查询结果都被确认为有效，所有执行的更新都作为一个单一的原子变化应用到数据库中。如果`Commit`失败，那么`Tx`上的所有`Query`和`Exec`的结果都应该被视为无效而丢弃。

  - 使用[Tx.Rollback](https://pkg.go.dev/database/sql#Tx.Rollback)来回滚事务。

    即使`Tx.Rollback`失败，该事务也不再有效，也不会被提交到数据库。

### Best practices 最佳实践

​	遵循下面的最佳实践，可以更好地了解事务有时需要的复杂语义和连接管理。

- 使用本节中描述的API来管理事务。**不要直接使用与事务相关的SQL语句**，如`BEGIN`和`COMMIT`——这样做会使你的数据库处于不可预测的状态，尤其是在并发程序中。
- 当使用事务时，注意**不要直接调用非事务的**`sql.DB`方法，因为这些方法会在事务之外执行，将会给你的代码提供一个不一致的数据库状态，甚至导致**死锁**。

### Example 例子

​	下面的例子中的代码使用一个事务来创建新专辑的客户订单。在这一过程中，代码将：

1. 开始一个事务。
2. 延迟该事务的回滚。如果事务成功，它将在函数退出前被提交，使延迟的回滚调用成为no-op。如果事务失败，它将不会被提交，这意味着回滚将在函数退出时被调用。
3. 确认客户订购的专辑有足够的库存。
4. 如果有足够的存货，更新存货数量，用订购的专辑数量来减少它。
5. 创建一个新的订单，并为客户检索新订单的生成ID。
6. 提交事务并返回ID。

​	这个例子使用了`Tx`方法，这些方法需要一个`context.Context`实参。这使得函数的执行——包括数据库操作——在其运行时间过长或客户端连接关闭的情况下有可能被取消。更多信息请参见[Canceling in-progress operations（取消正在进行的操作）](../CancelingIn-progressDatabaseOperations)。

```go 
// CreateOrder creates an order for an album and returns the new order ID. => CreateOrder 为 album 创建一个订单，并返回新的订单 ID。
func CreateOrder(ctx context.Context, albumID, quantity, custID int) (orderID int64, err error) {

    // Create a helper function for preparing failure results.
    fail := func(err error) (int64, error) {
        return fmt.Errorf("CreateOrder: %v", err)
    }

    // Get a Tx for making transaction requests.
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return fail(err)
    }
    // Defer a rollback in case anything fails.
    defer tx.Rollback()

    // Confirm that album inventory is enough for the order. => 确认专辑有足够的库存。
    var enough bool
    if err = tx.QueryRowContext(ctx, "SELECT (quantity >= ?) from album where id = ?",
        quantity, albumID).Scan(&enough); err != nil {
        if err == sql.ErrNoRows {
            return fail(fmt.Errorf("no such album"))
        }
        return fail(err)
    }
    if !enough {
        return fail(fmt.Errorf("not enough inventory"))
    }

    // Update the album inventory to remove the quantity in the order.
    _, err = tx.ExecContext(ctx, "UPDATE album SET quantity = quantity - ? WHERE id = ?",
        quantity, albumID)
    if err != nil {
        return fail(err)
    }

    // Create a new row in the album_order table.
    result, err := tx.ExecContext(ctx, "INSERT INTO album_order (album_id, cust_id, quantity, date) VALUES (?, ?, ?, ?)",
        albumID, custID, quantity, time.Now())
    if err != nil {
        return fail(err)
    }
    // Get the ID of the order item just created.
    orderID, err := result.LastInsertId()
    if err != nil {
        return fail(err)
    }

    // Commit the transaction.
    if err = tx.Commit(); err != nil {
        return fail(err)
    }

    // Return the order ID.
    return orderID, nil
}
```