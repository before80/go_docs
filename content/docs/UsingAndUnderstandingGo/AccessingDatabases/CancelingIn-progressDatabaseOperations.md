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

​	您可以通过使用Go [context.Context]({{< ref "/stdLib/context#type-context">}})来管理正在进行的操作。`Context`是一个标准的Go数据值，可以报告它所表示的整体操作是否已经被取消，是否不再需要。通过在应用程序中的函数调用和服务中传递`context.Context`，这些（函数调用和服务）可以提前停止工作，并在不再需要其处理时返回一个错误。关于`Context`的更多信息，请参阅 [Go并发模式：Context]({{< ref "/goBlog/2014/GoConcurrencyPatternsContext" >}})。

例如，您可能想：

- 结束长期运行的操作，包括需要太长时间才能完成的数据库操作。
- 传播来自其他地方的取消请求，例如当客户端关闭连接时。

​	许多为Go开发者提供的API都包含接受`Context`实参的方法，这使得在整个应用程序中更容易使用`Context`。

### 在超时后取消数据库操作

​	您可以使用 `Context` 来设置一个超时或最后期限，超时后的操作将被取消。要派生一个带有超时或最后期限的`Context`，请调用[context.WithTimeout]({{< ref "/stdLib/context#func-withtimeout">}})或[context.WithDeadline]({{< ref "/stdLib/context#func-withdeadline">}})。

​	下面的超时例子中的代码派生了一个Context，并将其传递给`sql.DB` [QueryContext]({{< ref "/stdLib/database/sql#db-querycontext----go18">}})方法。

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

​	当一个上下文派生自一个外部上下文时，就像本例中`queryCtx`派生自`ctx`一样，如果外部上下文被取消，那么派生的上下文也会被自动取消。例如，在HTTP服务器中，`http.Request.Context`方法返回一个与请求相关的上下文。如果HTTP客户端断开连接或取消HTTP请求（在HTTP/2中可能），则该上下文就会被取消。将HTTP请求的上下文传递给上面的`QueryWithTimeout`会导致数据库查询提前停止，如果整个HTTP请求被取消或者查询时间超过5秒的话。

> 注意
>
> ​	当您创建一个有超时或截止日期的新`Context`时，（一定记得）始终推迟对`cancel`函数的调用。这将在外层函数退出时释放新`Context`所持有的资源。它也会取消`queryCtx`，但是当函数返回时，就不会再有任何东西使用 `queryCtx` 了。

