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

​	您可以使用表示事务的[sql.Tx]({{< ref "/stdLib/database/sql#type-tx">}})来执行数据库事务。除了表示特定于事务语义的`Commit`和`Rollback`方法之外，`sql.Tx`还有所有用来执行普通数据库操作的方法。要获取`sql.Tx`，可以调用`DB.Begin`或`DB.BeginTx`。

​	一个[数据库事务](https://en.wikipedia.org/wiki/Database_transaction)将多个操作作为更大目标的一部分进行分组。所有的操作都必须成功，或者都不能成功，在这两种情况下都要保持数据的完整性。通常，一个事务的工作流程包括：

​	你可以使用一个表示事务的[sql.Tx]({{< ref "/stdLib/database/sql#type-tx">}})来执行数据库事务。除了`Commit`和`Rollback`方法代表特定于事务的语义之外，`sql.Tx`还具有你用于执行常见数据库操作的所有方法。要获取`sql.Tx`，你可以调用`DB.Begin`或`DB.BeginTx`。

​	[数据库事务](https://en.wikipedia.org/wiki/Database_transaction)将多个操作分组为更大目标的一部分。所有操作都必须成功，或者都不能成功，并且无论哪种情况都保留数据的完整性。通常，事务工作流程包括：

1. 开始事务。
2. 执行一系列的数据库操作。
3. 如果没有发生错误，提交事务以进行数据库更改。
4. 如果发生错误，回滚事务使数据库保持不变。

​	`sql`包提供了开始和结束事务的方法，以及执行中间的数据库操作的方法。这些方法与上述工作流程中的四个步骤相对应。

- 开始事务

  [DB.Begin]({{< ref "/stdLib/database/sql#db-begin">}})或[DB.BeginTx]({{< ref "/stdLib/database/sql#db-begintx----go18">}})开始一个新的数据库事务，返回一个代表它的`sql.Tx`。

- 执行数据库操作。

  使用一个`sql.Tx`，您可以在一系列使用单一连接的操作中查询或更新数据库。为了支持这一点，`Tx`导出了以下方法：

  使用`sql.Tx`，你可以在单个连接中使用一系列操作来查询或更新数据库。为了支持这一点，`Tx`导出了以下方法：

  - [Exec]({{< ref "/stdLib/database/sql#tx-exec">}}) 和 [ExecContext](https://pkg.go.dev/database/sql#Tx.ExecContext  {{< ref "/stdLib/database/sql#tx-execcontext----go18">}}) ，用于通过SQL语句（如`INSERT`、`UPDATE`和`DELETE`）进行数据库更改。

    有关更多信息，请参阅[Executing statements that don’t return data（执行不返回数据的语句）](../ExecutingSQLStatementsThatDoNotReturnData)。

  - [Query]({{< ref "/stdLib/database/sql#tx-query">}})，[QueryContext]({{< ref "/stdLib/database/sql#tx-querycontext----go18">}})，[QueryRow]({{< ref "/stdLib/database/sql#tx-queryrow">}}) ，[QueryRowContext]({{< ref "/stdLib/database/sql#tx-queryrowcontext----go18">}}) 用于返回行的操作。

    有关更多信息，请参阅 [Querying for data （查询数据）](../QueryingForData)。

  - [Prepare]({{< ref "/stdLib/database/sql#tx-prepare">}})，[PrepareContext]({{< ref "/stdLib/database/sql#tx-preparecontext----go18">}})，[Stmt]({{< ref "/stdLib/database/sql#tx-stmt">}})，[StmtContext]({{< ref "/stdLib/database/sql#tx-stmtcontext----go18">}})用于预定义预处理语句。

    有关更多信息，请参阅[Using prepared statements （使用预处理语句）](../UsingPreparedStatements)。

- 用以下方法之一结束事务：

  - 使用[Tx.Commit]({{< ref "/stdLib/database/sql#tx-commit">}})提交事务。

    如果`Commit`成功（返回`nil`错误），那么所有的查询结果都被确认为有效，并且所有已执行的更新作为单个原子更改应用于数据库。如果`Commit`失败，那么`Tx`上的所有`Query`和`Exec`的结果都应该被视为无效而丢弃。

  - 使用[Tx.Rollback]({{< ref "/stdLib/database/sql#tx-rollback">}})回滚事务。

    即使`Tx.Rollback`失败，该事务也不再有效，也不会被提交到数据库。

### 最佳实践

​	遵循以下最佳实践，以更好地处理事务有时所需的复杂语义和连接管理。

- 使用本节中描述的API来管理事务。**不要直接使用与事务相关的SQL语句**，如`BEGIN`和`COMMIT`，这样做可能会使数据库处于不可预测的状态，特别是在并发程序中。 
- 在使用事务时，注意**不要直接调用非事务性的**`sql.DB`方法，因为这些方法将在事务之外执行，将会给您的代码提供一个不一致的数据库状态视图，甚至可能导致死锁。

### 示例

​	以下示例中的代码使用事务为客户订购的专辑创建新的客户订单。在这个过程中，代码将：

1. 开始一个事务。
2. 延迟该事务的回滚。如果事务成功，它将在函数退出前被提交，使延迟的回滚调用成为no-op（空操作）。如果事务失败，它将不会被提交，这意味着当函数退出时会调用回滚。
3. 确认客户订购的专辑有足够的库存。
4. 如果有足够的存货，更新存货计数，用订购的专辑数量来减少它。
5. 创建一个新的订单，并为客户检索新订单的生成ID。
6. 提交事务并返回ID。

​	这个例子使用了`Tx`方法，这些方法需要一个`context.Context`实参。这使得函数的执行——包括数据库操作——在其运行时间过长或客户端连接关闭的情况下有可能被取消。更多信息请参见[Canceling in-progress operations（取消正在进行的操作）](../CancelingIn-progressDatabaseOperations)。

​	这个示例使用了接受`context.Context`实参的`Tx`方法。这使得函数的执行（包括数据库操作）可以在运行时间过长或客户端连接关闭时被取消。有关更多信息，请参阅[Canceling in-progress operations（取消正在进行的操作）](../CancelingIn-progressDatabaseOperations)。

```go 
// CreateOrder 为 album 创建一个订单，并返回新的订单 ID。
func CreateOrder(ctx context.Context, albumID, quantity, custID int) (orderID int64, err error) {

    // 创建一个助手函数来准备失败结果。
    fail := func(err error) (int64, error) {
        return 0, fmt.Errorf("CreateOrder: %v", err)
    }

    // 获取一个Tx用于进行事务请求。
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return fail(err)
    }
    // 如果有任何操作失败，延迟回滚。
    defer tx.Rollback()

    // 确认专辑库存足够满足订单需求。
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

    // 更新专辑库存以减去订单中的数量。
    _, err = tx.ExecContext(ctx, "UPDATE album SET quantity = quantity - ? WHERE id = ?",
        quantity, albumID)
    if err != nil {
        return fail(err)
    }

    // 在album_order表中创建一个新的行。
    result, err := tx.ExecContext(ctx, "INSERT INTO album_order (album_id, cust_id, quantity, date) VALUES (?, ?, ?, ?)",
        albumID, custID, quantity, time.Now())
    if err != nil {
        return fail(err)
    }
    // 获取刚刚创建的订单项的ID。
    orderID, err := result.LastInsertId()
    if err != nil {
        return fail(err)
    }

    // 提交事务。
    if err = tx.Commit(); err != nil {
        return fail(err)
    }

    // 返回订单ID。
    return orderID, nil
}
```