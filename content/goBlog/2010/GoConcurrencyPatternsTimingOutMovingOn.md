+++
title = "go 并发模式：超时和继续"
weight = 4
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go Concurrency Patterns: Timing out, moving on - go 并发模式：超时和继续

https://go.dev/blog/concurrency-timeouts

Andrew Gerrand
23 September 2010

2010年9月23日

​	并发编程有其自己的习语。一个很好的例子就是超时。虽然Go的通道不支持超时，但实现它们很容易。假设我们想要从通道ch中接收数据，但是最多只想等待一秒钟以获取该值。我们首先要创建一个信号通道并启动一个协程，在发送到该通道之前使其休眠：

```go
timeout := make(chan bool, 1)
go func() {
    time.Sleep(1 * time.Second)
    timeout <- true
}()
```

​	然后我们可以使用`select`语句从ch或timeout中接收数据。如果一秒钟后ch上没有数据，将选择timeout，放弃从ch中读取。

```go
select {
case <-ch:
    // 从ch中读取到数据
case <-timeout:
    // the read from ch has timed out
}
```

​	timeout通道被缓冲为1个值，允许timeout协程发送到通道然后退出。该协程不知道（也不关心）接收到的值是否被使用。这意味着如果在超时之前从ch接收到数据，协程将不会一直等待下去。timeout通道最终将被垃圾收集器释放。

（在此示例中，我们使用time.Sleep来演示协程和通道的机制。在实际程序中，应该使用[time.After](https://pkg.go.dev/pkg/time/#After)函数，该函数返回一个通道，并在指定的持续时间后在该通道上发送值。）

​	让我们来看看这种模式的另一种变体。在这个例子中，我们有一个程序，同时从多个复制的数据库中读取数据。该程序只需要一个答案，并且它应该接受第一个到达的答案。

​	函数Query接受一个数据库连接的切片和一个查询字符串。它并行查询每个数据库，并返回它收到的第一个响应：

```go
func Query(conns []Conn, query string) Result {
    ch := make(chan Result)
    for _, conn := range conns {
        go func(c Conn) {
            select {
            case ch <- c.DoQuery(query):
            default:
            }
        }(conn)
    }
    return <-ch
}
```

​	在这个示例中，闭包执行了一个非阻塞的发送，它通过在 select 语句中使用带有默认情况的发送操作来实现。如果发送不能立即完成，将选择默认情况。使发送非阻塞可确保在循环中启动的任何 goroutine 都不会挂起。但是，如果结果在主函数到达接收之前到达，则发送可能失败，因为没有人准备接收。

​	这个问题是所谓的[竞争条件](https://en.wikipedia.org/wiki/Race_condition)的一个经典例子，但是解决方法很简单。我们只需要确保缓冲通道 ch（通过将缓冲区长度作为 [make](https://go.dev/pkg/builtin/#make) 的第二个参数添加），以保证第一个发送有一个值的放置位置。这确保发送将始终成功，并且无论执行顺序如何，都将检索到到达的第一个值。

​	这两个示例演示了 Go 可以表达 goroutine 之间复杂交互的简单性。
