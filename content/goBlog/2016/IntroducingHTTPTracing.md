+++
title = "介绍一下HTTP追踪"
weight = 4
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Introducing HTTP Tracing - 介绍一下HTTP追踪

https://go.dev/blog/http-tracing

Jaana Burcu Dogan
4 October 2016

## Introduction 简介

In Go 1.7 we introduced HTTP tracing, a facility to gather fine-grained information throughout the lifecycle of an HTTP client request. Support for HTTP tracing is provided by the [`net/http/httptrace`](https://go.dev/pkg/net/http/httptrace/) package. The collected information can be used for debugging latency issues, service monitoring, writing adaptive systems, and more.

在Go 1.7中，我们引入了HTTP跟踪，这是一种在HTTP客户端请求的整个生命周期中收集细粒度信息的工具。net/http/httptrace包提供了对HTTP追踪的支持。收集到的信息可用于调试延迟问题、服务监控、编写自适应系统等。

## HTTP events HTTP事件

The `httptrace` package provides a number of hooks to gather information during an HTTP round trip about a variety of events. These events include:

httptrace包提供了一些钩子来收集HTTP往返过程中各种事件的信息。这些事件包括：

- Connection creation 连接创建
- Connection reuse 连接重复使用
- DNS lookups DNS查询
- Writing the request to the wire 将请求写到电线上
- Reading the response 读取响应

## Tracing events 追踪事件

You can enable HTTP tracing by putting an [`*httptrace.ClientTrace`](https://go.dev/pkg/net/http/httptrace/#ClientTrace) containing hook functions into a request’s [`context.Context`](https://go.dev/pkg/context/#Context). Various [`http.RoundTripper`](https://go.dev/pkg/net/http/#RoundTripper) implementations report the internal events by looking for context’s `*httptrace.ClientTrace` and calling the relevant hook functions.

您可以通过把包含钩子函数的*httptrace.ClientTrace放到请求的context.Context中来启用HTTP跟踪。各种http.RoundTripper实现通过寻找context的*httptrace.ClientTrace并调用相关钩子函数来报告内部事件。

The tracing is scoped to the request’s context and users should put a `*httptrace.ClientTrace` to the request context before they start a request.

追踪的范围是请求的上下文，用户应该在开始请求之前把*httptrace.ClientTrace放到请求的上下文中。

```go
    req, _ := http.NewRequest("GET", "http://example.com", nil)
    trace := &httptrace.ClientTrace{
        DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
            fmt.Printf("DNS Info: %+v\n", dnsInfo)
        },
        GotConn: func(connInfo httptrace.GotConnInfo) {
            fmt.Printf("Got Conn: %+v\n", connInfo)
        },
    }
    req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
    if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
        log.Fatal(err)
    }
```

During a round trip, `http.DefaultTransport` will invoke each hook as an event happens. The program above will print the DNS information as soon as the DNS lookup is complete. It will similarly print connection information when a connection is established to the request’s host.

在往返过程中，http.DefaultTransport将在事件发生时调用每个钩子。上面的程序将在DNS查询完成后立即打印DNS信息。同样，当与请求的主机建立连接时，它也会打印连接信息。

## Tracing with http.Client 用http.Client追踪

The tracing mechanism is designed to trace the events in the lifecycle of a single `http.Transport.RoundTrip`. However, a client may make multiple round trips to complete an HTTP request. For example, in the case of a URL redirection, the registered hooks will be called as many times as the client follows HTTP redirects, making multiple requests. Users are responsible for recognizing such events at the `http.Client` level. The program below identifies the current request by using an `http.RoundTripper` wrapper.

追踪机制被设计为追踪单个http.Transport.RoundTrip生命周期内的事件。然而，一个客户端可能会进行多次往返以完成一个HTTP请求。例如，在URL重定向的情况下，注册的钩子将被多次调用，因为客户端跟随HTTP重定向，提出多个请求。用户负责在http.Client级别识别此类事件。下面的程序通过使用一个http.RoundTripper包装器来识别当前请求。

```go
// +build OMIT

package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httptrace"
)

// transport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type transport struct {
    current *http.Request
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
    t.current = req
    return http.DefaultTransport.RoundTrip(req)
}

// GotConn prints whether the connection has been used previously
// for the current request.
func (t *transport) GotConn(info httptrace.GotConnInfo) {
    fmt.Printf("Connection reused for %v? %v\n", t.current.URL, info.Reused)
}

func main() {
    t := &transport{}

    req, _ := http.NewRequest("GET", "https://google.com", nil)
    trace := &httptrace.ClientTrace{
        GotConn: t.GotConn,
    }
    req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

    client := &http.Client{Transport: t}
    if _, err := client.Do(req); err != nil {
        log.Fatal(err)
    }
}
```

The program will follow the redirect of google.com to [www.google.com](http://www.google.com/) and will output:

该程序将跟随google.com的重定向到www.google.com，并将输出：

```
Connection reused for https://google.com? false
Connection reused for https://www.google.com/? false
```

The Transport in the `net/http` package supports tracing of both HTTP/1 and HTTP/2 requests.

net/http包中的Transport支持对HTTP/1和HTTP/2请求进行追踪。

If you are an author of a custom `http.RoundTripper` implementation, you can support tracing by checking the request context for an `*httptest.ClientTrace` and invoking the relevant hooks as the events occur.

如果您是自定义http.RoundTripper实现的作者，您可以通过检查请求上下文的*httptest.ClientTrace并在事件发生时调用相关钩子来支持追踪。

## Conclusion 总结

HTTP tracing is a valuable addition to Go for those who are interested in debugging HTTP request latency and writing tools for network debugging for outbound traffic. By enabling this new facility, we hope to see HTTP debugging, benchmarking and visualization tools from the community — such as [httpstat](https://github.com/davecheney/httpstat).

对于那些对调试HTTP请求延迟和编写网络调试工具的出站流量感兴趣的人来说，HTTP跟踪是Go的一个重要补充。通过启用这一新设施，我们希望看到来自社区的HTTP调试、基准测试和可视化工具--比如httpstat。
