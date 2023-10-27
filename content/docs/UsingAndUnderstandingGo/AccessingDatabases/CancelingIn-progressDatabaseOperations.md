+++
title = "取消正在进行的操作"
date = 2023-05-17T15:03:14+08:00
weight = 7
description = ""
isCJKLanguage = true
draft = false
+++
# Canceling in-progress operations - 取消正在进行的操作

> 原文：[https://go.dev/doc/database/cancel-operations](https://go.dev/doc/database/cancel-operations)

​	你可以使用Go的[context.Context]({{< ref "/stdLib/context#type-context">}})来管理进行中的操作。`Context`是一个标准的Go数据值，可以报告它所代表的整个操作是否被取消，是否不再需要。通过在应用程序中的函数调用和服务之间传递`context.Context`，当它们的处理不再需要时，它们可以提前停止工作并返回错误。有关`Context`的更多信息，请参阅[Go并发模式：Context]({{< ref "/goBlog/2014/GoConcurrencyPatternsContext" >}})。

​	例如，你可能想要：

- 结束长时间运行的操作，包括完成时间过长的数据库操作。
- 传播来自其他地方的取消请求，例如当客户端关闭连接时。 

​	许多面向Go开发人员的API都包含接受`Context`实参的方法，使你更容易在整个应用程序中使用`Context`。

### 在超时后取消数据库操作

​	您可以使用 `Context` 来设置一个超时或最后期限，超时后的操作将被取消。要派生一个带有超时或最后期限的`Context`，请调用[context.WithTimeout]({{< ref "/stdLib/context#func-withtimeout">}})或[context.WithDeadline]({{< ref "/stdLib/context#func-withdeadline">}})。

​	下面的超时例子中的代码派生了一个Context，并将其传递给`sql.DB` [QueryContext]({{< ref "/stdLib/database/sql#db-querycontext----go18">}})方法。

​	你可以使用一个`Context`来设置超时或截止时间，超过该时间后将取消操作。要派生出具有超时或截止时间的`Context`，可以调用[context.WithTimeout]({{< ref "/stdLib/context#func-withtimeout">}})或[context.WithDeadline]({{< ref "/stdLib/context#func-withdeadline">}})。

​	以下超时示例中的代码派生出一个Context，并将其传递给`sql.DB`的[QueryContext]({{< ref "/stdLib/database/sql#db-querycontext----go18">}})方法。

```go 
func QueryWithTimeout(ctx context.Context) {
    // Create a Context with a timeout.
    queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    // Pass the timeout Context with a query.
    rows, err := db.QueryContext(queryCtx, "SELECT * FROM album")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Handle returned rows.
}
```

​	当一个上下文从外部上下文派生出来时，就像`queryCtx`在这个例子中从`ctx`派生出来一样，如果外部上下文被取消，那么派生的上下文也会自动被取消。例如，在HTTP服务器中，`http.Request.Context`方法返回与请求关联的上下文。如果HTTP客户端断开连接或取消HTTP请求（在HTTP/2中是可能的），则该上下文将被取消。将HTTP请求的上下文传递给上面的`QueryWithTimeout`将导致数据库查询提前停止，无论是**整个HTTP请求被取消**还是**查询花费超过五秒钟**。

> 注意：在创建带有超时或截止日期的新的`Context`时，始终延迟调用返回的`cancel`函数。这将在外层函数退出时释放由新`Context`持有的资源。它还会取消`queryCtx`，但是在函数返回时，不应该再使用`queryCtx`了。
