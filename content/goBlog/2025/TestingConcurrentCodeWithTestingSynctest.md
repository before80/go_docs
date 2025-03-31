+++
title = "使用testing/synctest测试并发代码"
date = 2025-03-31T14:23:28+08:00
weight = 980
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/synctest](https://go.dev/blog/synctest)

# Testing concurrent code with testing/synctest - 使用testing/synctest测试并发代码



Damien Neil

19 February 2025

 2025 年 2 月 19 日

One of Go’s signature features is built-in support for concurrency. Goroutines and channels are simple and effective primitives for writing concurrent programs.

​	Go 的标志性特性之一是内置并发支持。Goroutine 和通道是编写并发程序时简单而高效的原语。

However, testing concurrent programs can be difficult and error prone.

​	然而，测试并发程序可能既困难又容易出错。

In Go 1.24, we are introducing a new, experimental [`testing/synctest`](https://go.dev/pkg/testing/synctest) package to support testing concurrent code. This post will explain the motivation behind this experiment, demonstrate how to use the synctest package, and discuss its potential future.

​	在 Go 1.24 中，我们推出了一个新的实验性 [`testing/synctest`](https://go.dev/pkg/testing/synctest) 包来支持并发代码的测试。本文将解释这一实验背后的动机，演示如何使用 synctest 包，并讨论其未来潜力。

In Go 1.24, the `testing/synctest` package is experimental and not subject to the Go compatibility promise. It is not visible by default. To use it, compile your code with `GOEXPERIMENT=synctest` set in your environment.

​	在 Go 1.24 中，`testing/synctest` 包是实验性的，不受 Go 兼容性承诺的约束。默认情况下该包不可见。要使用它，请在环境中设置 `GOEXPERIMENT=synctest` 后编译代码。

## 测试并发程序是困难的 Testing concurrent programs is difficult

To begin with, let us consider a simple example.

​	首先，让我们看一个简单的例子。

The [`context.AfterFunc`](https://go.dev/pkg/context#AfterFunc) function arranges for a function to be called in its own goroutine after a context is canceled. Here is a possible test for `AfterFunc`:

​	[`context.AfterFunc`](https://go.dev/pkg/context#AfterFunc) 函数安排在上下文被取消后，在其自己的 goroutine 中调用一个函数。下面是 `AfterFunc` 的一个可能的测试案例：

```go
func TestAfterFunc(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())

    calledCh := make(chan struct{}) // closed when AfterFunc is called  当 AfterFunc 被调用时关闭该通道
    context.AfterFunc(ctx, func() {
        close(calledCh)
    })
	// TODO: 断言 AfterFunc 在上下文取消之前未被调用。
    // TODO: Assert that the AfterFunc has not been called.

    cancel()

    // TODO: Assert that the AfterFunc has been called.
    // TODO: 断言 AfterFunc 在上下文取消之后已被调用。
}
```

We want to check two conditions in this test: The function is not called before the context is canceled, and the function *is* called after the context is canceled.

​	我们希望在这个测试中检查两个条件：函数在上下文取消前未被调用，而在上下文取消后确实被调用。

Checking a negative in a concurrent system is difficult. We can easily test that the function has not been called *yet*, but how do we check that it *will not* be called?

​	在并发系统中检查否定条件很困难。我们可以轻易测试函数是否 *尚未* 被调用，但如何验证它 *永远不会* 被调用呢？

A common approach is to wait for some amount of time before concluding that an event will not happen. Let’s try introducing a helper function to our test which does this.

​	一种常见的方法是等待一段时间，然后断定某个事件不会发生。让我们试着在测试中引入一个辅助函数来实现这一点。

```
// funcCalled reports whether the function was called.
// funcCalled 报告函数是否被调用。
funcCalled := func() bool {
    select {
    case <-calledCh:
        return true
    case <-time.After(10 * time.Millisecond):
        return false
    }
}

if funcCalled() {
    t.Fatalf("AfterFunc function called before context is canceled")
}

cancel()

if !funcCalled() {
    t.Fatalf("AfterFunc function not called after context is canceled")
}
```

This test is slow: 10 milliseconds isn’t a lot of time, but it adds up over many tests.

​	这个测试运行较慢：10 毫秒可能不长，但在许多测试中累计起来会很耗时。

This test is also flaky: 10 milliseconds is a long time on a fast computer, but it isn’t unusual to see pauses lasting several seconds on shared and overloaded [CI](https://en.wikipedia.org/wiki/Continuous_integration) systems.

​	这个测试也容易不稳定：10 毫秒在快速计算机上可能很长，但在共享和负载过重的 [CI](https://en.wikipedia.org/wiki/Continuous_integration) 系统上看到暂停持续几秒钟并不罕见。

We can make the test less flaky at the expense of making it slower, and we can make it less slow at the expense of making it flakier, but we can’t make it both fast and reliable.

​	我们可以以牺牲速度为代价使测试更稳定，也可以以牺牲稳定性为代价使其更快，但无法同时做到又快又可靠。

## 介绍 testing/synctest 包 Introducing the testing/synctest package

The `testing/synctest` package solves this problem. It allows us to rewrite this test to be simple, fast, and reliable, without any changes to the code being tested.

​	`testing/synctest` 包解决了这个问题。它允许我们重写测试，使之简单、快速且可靠，而无需更改被测试的代码。

The package contains only two functions: `Run` and `Wait`.

​	该包仅包含两个函数：`Run` 和 `Wait`。

`Run` calls a function in a new goroutine. This goroutine and any goroutines started by it exist in an isolated environment which we call a *bubble*. `Wait` waits for every goroutine in the current goroutine’s bubble to block on another goroutine in the bubble.

​	`Run` 在一个新的 goroutine 中调用一个函数。该 goroutine 以及由它启动的任何 goroutine 都存在于我们称之为 *bubble* 的隔离环境中。`Wait` 则等待当前 goroutine 的 bubble 中的每个 goroutine 都被另一个 goroutine 阻塞。

Let’s rewrite our test above using the `testing/synctest` package.

​	让我们使用 `testing/synctest` 包重写上面的测试。

```
func TestAfterFunc(t *testing.T) {
    synctest.Run(func() {
        ctx, cancel := context.WithCancel(context.Background())

        funcCalled := false
        context.AfterFunc(ctx, func() {
            funcCalled = true
        })

        synctest.Wait()
        if funcCalled {
            t.Fatalf("AfterFunc function called before context is canceled")
        }

        cancel()

        synctest.Wait()
        if !funcCalled {
            t.Fatalf("AfterFunc function not called after context is canceled")
        }
    })
}
```

This is almost identical to our original test, but we have wrapped the test in a `synctest.Run` call and we call `synctest.Wait` before asserting that the function has been called or not.

​	这与我们原始测试几乎完全相同，只不过我们将测试包装在 `synctest.Run` 调用中，并在断言函数是否被调用前调用了 `synctest.Wait`。

The `Wait` function waits for every goroutine in the caller’s bubble to block. When it returns, we know that the context package has either called the function, or will not call it until we take some further action.

​	`Wait` 函数会等待调用者的 bubble 中的每个 goroutine 都处于阻塞状态。当它返回时，我们就知道 context 包要么已经调用了该函数，要么在我们采取进一步措施之前不会调用它。

This test is now both fast and reliable.

​	这个测试现在既快速又可靠。

The test is simpler, too: we have replaced the `calledCh` channel with a boolean. Previously we needed to use a channel to avoid a data race between the test goroutine and the `AfterFunc` goroutine, but the `Wait` function now provides that synchronization.

​	测试也更简单了：我们用一个布尔值取代了 `calledCh` 通道。之前为了避免测试 goroutine 和 `AfterFunc` goroutine 之间的数据竞争，我们需要使用通道，但现在 `Wait` 函数提供了这种同步机制。

The race detector understands `Wait` calls, and this test passes when run with `-race`. If we remove the second `Wait` call, the race detector will correctly report a data race in the test.

​	竞态检测器能够识别 `Wait` 调用，并且该测试在使用 `-race` 参数运行时能通过。如果我们移除第二个 `Wait` 调用，竞态检测器将会正确报告测试中的数据竞争。

## 测试时间 Testing time

Concurrent code often deals with time.

​	并发代码常常涉及时间问题。

Testing code that works with time can be difficult. Using real time in tests causes slow and flaky tests, as we have seen above. Using fake time requires avoiding `time` package functions, and designing the code under test to work with an optional fake clock.

​	测试处理时间的代码可能很困难。正如我们上面看到的，使用真实时间会导致测试缓慢且不稳定。使用虚拟时间需要避免使用 `time` 包函数，并设计被测试代码以适应可选的虚拟时钟。

The `testing/synctest` package makes it simpler to test code that uses time.

​	`testing/synctest` 包简化了对使用时间的代码的测试。

Goroutines in the bubble started by `Run` use a fake clock. Within the bubble, functions in the `time` package operate on the fake clock. Time advances in the bubble when all goroutines are blocked.

​	在 `Run` 启动的 bubble 中的 goroutine 使用虚拟时钟。在 bubble 内，`time` 包的函数在虚拟时钟上运行。当所有 goroutine 都被阻塞时，bubble 中的时间才会推进。

To demonstrate, let’s write a test for the [`context.WithTimeout`](https://go.dev/pkg/context#WithTimeout) function. `WithTimeout` creates a child of a context, which expires after a given timeout.

​	为了演示，让我们为 [`context.WithTimeout`](https://go.dev/pkg/context#WithTimeout) 函数编写一个测试。`WithTimeout` 创建一个上下文的子上下文，该子上下文在给定的超时后失效。

```
func TestWithTimeout(t *testing.T) {
    synctest.Run(func() {
        const timeout = 5 * time.Second
        ctx, cancel := context.WithTimeout(context.Background(), timeout)
        defer cancel()

        // Wait just less than the timeout.
        // 等待的时间略短于超时。
        time.Sleep(timeout - time.Nanosecond)
        synctest.Wait()
        if err := ctx.Err(); err != nil {
            t.Fatalf("before timeout, ctx.Err() = %v; want nil", err)
        }

        // Wait the rest of the way until the timeout.
        // 再等待剩余的时间直到超时。
        time.Sleep(time.Nanosecond)
        synctest.Wait()
        if err := ctx.Err(); err != context.DeadlineExceeded {
            t.Fatalf("after timeout, ctx.Err() = %v; want DeadlineExceeded", err)
        }
    })
}
```

We write this test just as if we were working with real time. The only difference is that we wrap the test function in `synctest.Run`, and call `synctest.Wait` after each `time.Sleep` call to wait for the context package’s timers to finish running.

​	我们编写这个测试就像在使用真实时间一样。唯一的区别是我们将测试函数包装在 `synctest.Run` 中，并在每次调用 `time.Sleep` 后调用 `synctest.Wait`，以等待 context 包的定时器完成运行。

## Blocking and the bubble

A key concept in `testing/synctest` is the bubble becoming *durably blocked*. This happens when every goroutine in the bubble is blocked, and can only be unblocked by another goroutine in the bubble.

​	`testing/synctest` 的一个关键概念是 bubble 变得 *持久阻塞*。这发生在 bubble 中的每个 goroutine 都被阻塞，并且只能由 bubble 中的另一个 goroutine 解阻时。

When a bubble is durably blocked:

​	当 bubble 持久阻塞时：

- If there is an outstanding `Wait` call, it returns.
  - 如果有未完成的 `Wait` 调用，则返回。
- Otherwise, time advances to the next time that could unblock a goroutine, if any.
  - 否则，时间推进到下一个可能解阻某个 goroutine 的时刻（如果有的话）。
- Otherwise, the bubble is deadlocked and `Run` panics.
  - 否则，bubble 处于死锁状态，`Run` 将触发 panic。

A bubble is not durably blocked if any goroutine is blocked but might be woken by some event from outside the bubble.

​	如果有任何 goroutine 被阻塞但可能会被来自 bubble 外部的某个事件唤醒，则该 bubble 不被视为持久阻塞。

The complete list of operations which durably block a goroutine is:

​	 持久阻塞 goroutine 的操作完整列表如下：

- a send or receive on a nil channel
  - 对 nil 通道进行发送或接收
- a send or receive blocked on a channel created within the same bubble
  - 对同一 bubble 内创建的通道进行发送或接收而被阻塞
- a select statement where every case is durably blocking
  - 一个 select 语句，其中每个 case 都是持久阻塞的
- `time.Sleep`
- `sync.Cond.Wait`
- `sync.WaitGroup.Wait`

### Mutexes

Operations on a `sync.Mutex` are not durably blocking.

​	对 `sync.Mutex` 的操作不是持久阻塞的。

It is common for functions to acquire a global mutex. For example, a number of functions in the reflect package use a global cache guarded by a mutex. If a goroutine in a synctest bubble blocks while acquiring a mutex held by a goroutine outside the bubble, it is not durably blocked—it is blocked, but will be unblocked by a goroutine from outside its bubble.

​	函数通常会获取全局互斥锁。例如，reflect 包中的多个函数使用受互斥锁保护的全局缓存。如果一个 synctest bubble 中的 goroutine 在获取由 bubble 外部的 goroutine 持有的互斥锁时阻塞，它不会被视为持久阻塞——它是阻塞的，但会被来自 bubble 外部的 goroutine 解阻。

Since mutexes are usually not held for long periods of time, we simply exclude them from `testing/synctest`’s consideration.

​	由于互斥锁通常不会长时间被持有，我们简单地将它们排除在 `testing/synctest` 的考虑之外。

### Channels

Channels created within a bubble behave differently from ones created outside.

​	在 bubble 内创建的通道的行为与在外部创建的通道不同。

Channel operations are durably blocking only if the channel is bubbled (created in the bubble). Operating on a bubbled channel from outside the bubble panics.

​	只有当通道在 bubble 内创建（bubbled）时，其操作才会持久阻塞。从 bubble 外部操作一个 bubbled 通道会触发 panic。

These rules ensure that a goroutine is durably blocked only when communicating with goroutines within its bubble.

​	这些规则确保只有当 goroutine 与其 bubble 内的其他 goroutine 通信时，才会被视为持久阻塞。

### I/O

External I/O operations, such as reading from a network connection, are not durably blocking.

​	外部 I/O 操作，例如从网络连接中读取数据，并非持久阻塞。

Network reads may be unblocked by writes from outside the bubble, possibly even from other processes. Even if the only writer to a network connection is also in the same bubble, the runtime cannot distinguish between a connection waiting for more data to arrive and one where the kernel has received data and is in the process of delivering it.

​	网络读取可能会被来自 bubble 外部的写操作（甚至其他进程）解阻。即使网络连接的唯一写入者也在同一个 bubble 中，运行时也无法区分连接是在等待更多数据到达，还是内核已接收到数据并正在传递过程中。

Testing a network server or client with synctest will generally require supplying a fake network implementation. For example, the [`net.Pipe`](https://go.dev/pkg/net#Pipe) function creates a pair of `net.Conn`s that use an in-memory network connection and can be used in synctest tests.

​	使用 synctest 测试网络服务器或客户端通常需要提供一个假的网络实现。例如，[`net.Pipe`](https://go.dev/pkg/net#Pipe) 函数创建了一对使用内存网络连接的 `net.Conn`，可用于 synctest 测试。

## Bubble lifetime

The `Run` function starts a goroutine in a new bubble. It returns when every goroutine in the bubble has exited. It panics if the bubble is durably blocked and cannot be unblocked by advancing time.

​	`Run` 函数在新的 bubble 中启动一个 goroutine。当 bubble 中的每个 goroutine 都退出后，它才返回。如果 bubble 持久阻塞且无法通过推进时间解阻，则会触发 panic。

The requirement that every goroutine in the bubble exit before Run returns means that tests must be careful to clean up any background goroutines before completing.

​	要求 bubble 中的每个 goroutine 在 Run 返回前退出，这意味着测试必须在完成之前仔细清理所有后台 goroutine。

## Testing networked code

Let’s look at another example, this time using the `testing/synctest` package to test a networked program. For this example, we’ll test the `net/http` package’s handling of the 100 Continue response.

​	让我们看另一个例子，这次使用 `testing/synctest` 包来测试一个网络程序。在此例中，我们将测试 `net/http` 包对 100 Continue 响应的处理。

An HTTP client sending a request can include an “Expect: 100-continue” header to tell the server that the client has additional data to send. The server may then respond with a 100 Continue informational response to request the rest of the request, or with some other status to tell the client that the content is not needed. For example, a client uploading a large file might use this feature to confirm that the server is willing to accept the file before sending it.

​	发送请求的 HTTP 客户端可以包含 “Expect: 100-continue” 头，以告知服务器客户端还有额外数据要发送。服务器随后可能会响应 100 Continue 信息响应，请求发送剩余部分，或以其他状态告知客户端内容不需要。例如，一个上传大文件的客户端可能会利用此功能在发送文件前确认服务器是否愿意接收该文件。

Our test will confirm that when sending an “Expect: 100-continue” header the HTTP client does not send a request’s content before the server requests it, and that it does send the content after receiving a 100 Continue response.

​	我们的测试将确认，当发送 “Expect: 100-continue” 头时，HTTP 客户端在服务器请求之前不会发送请求内容，并且在收到 100 Continue 响应后会发送内容。

Often tests of a communicating client and server can use a loopback network connection. When working with `testing/synctest`, however, we will usually want to use a fake network connection to allow us to detect when all goroutines are blocked on the network. We’ll start this test by creating an `http.Transport` (an HTTP client) that uses an in-memory network connection created by [`net.Pipe`](https://go.dev/pkg/net#Pipe).

​	通常，测试通信的客户端和服务器可以使用回环网络连接。但是在使用 `testing/synctest` 时，我们通常希望使用一个伪网络连接，以便检测何时所有 goroutine 都在网络上阻塞。我们将通过创建一个使用 [`net.Pipe`](https://go.dev/pkg/net#Pipe) 创建的内存网络连接的 `http.Transport`（HTTP 客户端）来开始这个测试。

```
func Test(t *testing.T) {
    synctest.Run(func() {
        srvConn, cliConn := net.Pipe()
        defer srvConn.Close()
        defer cliConn.Close()
        tr := &http.Transport{
            DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
                return cliConn, nil
            },
            // Setting a non-zero timeout enables "Expect: 100-continue" handling.
            // Since the following test does not sleep,
            // we will never encounter this timeout,
            // even if the test takes a long time to run on a slow machine.
            // 设置一个非零超时以启用 "Expect: 100-continue" 处理。
            // 因为下面的测试不会 sleep，
            // 即使在慢速机器上测试耗时较长，我们也不会遇到此超时。
            ExpectContinueTimeout: 5 * time.Second,
        }
```

We send a request on this transport with the “Expect: 100-continue” header set. The request is sent in a new goroutine, since it won’t complete until the end of the test.

​	我们在此传输上发送一个带有 “Expect: 100-continue” 头的请求。该请求在新的 goroutine 中发送，因为它将在测试结束前无法完成。

```
        body := "request body"
        go func() {
            req, _ := http.NewRequest("PUT", "http://test.tld/", strings.NewReader(body))
            req.Header.Set("Expect", "100-continue")
            resp, err := tr.RoundTrip(req)
            if err != nil {
                t.Errorf("RoundTrip: unexpected error %v", err)
            } else {
                resp.Body.Close()
            }
        }()
```

We read the request headers sent by the client.

​	我们读取客户端发送的请求头。

```
       req, err := http.ReadRequest(bufio.NewReader(srvConn))
        if err != nil {
            t.Fatalf("ReadRequest: %v", err)
        }
```

Now we come to the heart of the test. We want to assert that the client will not send the request body yet.

​	现在我们进入测试的核心部分。我们希望断言客户端尚未发送请求体。

We start a new goroutine copying the body sent to the server into a `strings.Builder`, wait for all goroutines in the bubble to block, and verify that we haven’t read anything from the body yet.

​	我们启动一个新的 goroutine，将发送到服务器的请求体复制到 `strings.Builder` 中，等待 bubble 中所有 goroutine 阻塞，并验证我们还没有从请求体中读取任何内容。

If we forget the `synctest.Wait` call, the race detector will correctly complain about a data race, but with the `Wait` this is safe.

​	如果我们忘记调用 `synctest.Wait`，竞态检测器会正确报告数据竞争，但使用 `Wait` 后则是安全的。

```
        var gotBody strings.Builder
        go io.Copy(&gotBody, req.Body)
        synctest.Wait()
        if got := gotBody.String(); got != "" {
            t.Fatalf("before sending 100 Continue, unexpectedly read body: %q", got)
        }
```

We write a “100 Continue” response to the client and verify that it now sends the request body.

​	我们向客户端写入 “100 Continue” 响应，并验证它现在发送了请求体。

```
        srvConn.Write([]byte("HTTP/1.1 100 Continue\r\n\r\n"))
        synctest.Wait()
        if got := gotBody.String(); got != body {
            t.Fatalf("after sending 100 Continue, read body %q, want %q", got, body)
        }
```

And finally, we finish up by sending the “200 OK” response to conclude the request.

​	最后，我们通过发送 “200 OK” 响应来结束该请求。

We have started several goroutines during this test. The `synctest.Run` call will wait for all of them to exit before returning.

​	在此测试中，我们启动了多个 goroutine。`synctest.Run` 调用会等待所有 goroutine 退出后再返回。

```
      srvConn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
    })
}
```

This test can be easily extended to test other behaviors, such as verifying that the request body is not sent if the server does not ask for it, or that it is sent if the server does not respond within a timeout.

​	这个测试可以很容易地扩展到测试其他行为，例如验证如果服务器不请求，则请求体不会被发送；或者如果服务器在超时内未响应，则请求体会被发送。

## 实验状态 Status of the experiment

We are introducing [`testing/synctest`](https://go.dev/pkg/testing/synctest) in Go 1.24 as an *experimental* package. Depending on feedback and experience we may release it with or without amendments, continue the experiment, or remove it in a future version of Go.

​	我们在 Go 1.24 中引入 [`testing/synctest`](https://go.dev/pkg/testing/synctest) 作为一个 *实验性* 包。根据反馈和使用经验，我们可能会在修改后或不修改后发布它，继续实验，或在未来的 Go 版本中移除它。

The package is not visible by default. To use it, compile your code with `GOEXPERIMENT=synctest` set in your environment.

​	该包默认情况下不可见。要使用它，请在环境中设置 `GOEXPERIMENT=synctest` 后编译代码.

We want to hear your feedback! If you try out `testing/synctest`, please report your experiences, positive or negative, on [go.dev/issue/67434](https://go.dev/issue/67434).

​	我们希望听到您的反馈！如果您试用 `testing/synctest`，请将您的正面或负面体验报告至 [go.dev/issue/67434](https://go.dev/issue/67434).