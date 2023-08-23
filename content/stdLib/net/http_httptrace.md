+++
title = "http/httptrace"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# httptrace

https://pkg.go.dev/net/http/httptrace@go1.20.1

​	httptrace 包提供了在 HTTP 客户端请求内部跟踪事件的机制。

## Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}

```



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func WithClientTrace 

``` go 
func WithClientTrace(ctx context.Context, trace *ClientTrace) context.Context
```

​	WithClientTrace 函数基于提供的父上下文（ctx）返回一个新的上下文。使用返回的上下文进行的 HTTP 客户端请求将使用提供的跟踪钩子（trace hooks），除此之外，还会使用之前在 ctx 中注册的任何钩子。在提供的 trace 中定义的任何钩子都将首先被调用。

## 类型

### type ClientTrace 

``` go 
type ClientTrace struct {	
    // GetConn 在创建连接或从空闲连接池中获取连接之前调用。hostPort 是目标或代理的“host:port”。
	// 即使已经有一个可用的空闲缓存连接，也会调用 GetConn。
	GetConn func(hostPort string)

    // GotConn 在成功获取连接后调用。对于无法获取连接的情况，没有钩子；而是使用 Transport.RoundTrip 中的错误。
	GotConn func(GotConnInfo)

    // PutIdleConn 在连接返回到空闲连接池时调用。如果 err 为 nil，则连接成功返回到空闲连接池。
	// 如果 err 不为 nil，则表示未成功返回连接。
    // 如果通过 Transport.DisableKeepAlives 禁用连接重用，则不会调用 PutIdleConn。
	// 在调用方的 Response.Body.Close 调用返回之前，不会调用 PutIdleConn。
	// 对于 HTTP/2，此钩子当前不使用。
	PutIdleConn func(err error)

    // GotFirstResponseByte 在响应头的第一个字节可用时调用。
	GotFirstResponseByte func()

    // Got100Continue 在服务器回复“100 Continue”响应时调用。
	Got100Continue func()

    // Got1xxResponse 在最终的非 1xx 响应之前，对每个 1xx 信息响应头进行调用。
	// 对于返回的“100 Continue”响应，即使定义了 Got100Continue，也会调用 Got1xxResponse。
	// 如果返回错误，则客户端请求将中止并带有该错误值。
	Got1xxResponse func(code int, header textproto.MIMEHeader) error

    // DNSStart 在开始 DNS 查找时调用。
	DNSStart func(DNSStartInfo)

	// DNSDone is called when a DNS lookup ends.
    // DNSDone 在 DNS 查找结束时调用。
	DNSDone func(DNSDoneInfo)

    // ConnectStart 在开始新连接的拨号时调用。
    // 如果启用了 net.Dialer.DualStack（IPv6 “Happy Eyeballs”）支持，则可能会多次调用此函数。
	ConnectStart func(network, addr string)

    // ConnectDone 在新连接的拨号完成时调用。提供的 err 指示连接是否成功完成。
	// 如果启用了 net.Dialer.DualStack（“Happy Eyeballs”）支持，则可能会多次调用此函数。
	ConnectDone func(network, addr string, err error)

    // TLSHandshakeStart 在开始 TLS 握手时调用。
    // 当通过 HTTP 代理连接到 HTTPS 站点时，握手将在代理处理 CONNECT 请求后进行。
	TLSHandshakeStart func()

    // TLSHandshakeDone 在 TLS 握手完成后调用，
    // 其中包含成功握手的连接状态，或握手失败时的非 nil 错误。
	TLSHandshakeDone func(tls.ConnectionState, error)

    // WroteHeaderField 在 Transport 写入每个请求头之后调用。
    // 在此调用时，值可能已缓冲并尚未写入网络。
	WroteHeaderField func(key string, value []string)

    // WroteHeaders 在 Transport 写入所有请求头后调用。
	WroteHeaders func()

    // Wait100Continue 如果请求指定了“Expect: 100-continue”，
    // 且 Transport 已写入请求头但在写入请求主体之前等待服务器的“100 Continue”响应。
	Wait100Continue func()

    // WroteRequest 在写入请求和任何正文的结果上调用。
    // 在重试请求的情况下可能会被多次调用。
	WroteRequest func(WroteRequestInfo)
}
```

ClientTrace is a set of hooks to run at various stages of an outgoing HTTP request. Any particular hook may be nil. Functions may be called concurrently from different goroutines and some may be called after the request has completed or failed.

​	ClientTrace 是一组在传出的 HTTP 请求的不同阶段运行的钩子。任何特定的钩子都可能为 nil。函数可以从不同的 goroutine 并发地调用，有些可能在请求已完成或失败后调用。

ClientTrace currently traces a single HTTP request & response during a single round trip and has no hooks that span a series of redirected requests.

​	ClientTrace 当前在单个往返期间跟踪单个 HTTP 请求和响应，并且没有跨越一系列重定向请求的钩子。

​	有关更多信息，请参见 https://blog.golang.org/http-tracing。

#### func ContextClientTrace 

``` go 
func ContextClientTrace(ctx context.Context) *ClientTrace
```

​	ContextClientTrace 函数返回与提供的上下文关联的 ClientTrace。如果没有关联的，则返回 nil。

### type DNSDoneInfo 

``` go 
type DNSDoneInfo struct {
    // Addrs 是 DNS 查找中找到的 IPv4 和/或 IPv6 地址。切片的内容不应被修改。
	Addrs []net.IPAddr

    // Err 是在 DNS 查找期间出现的任何错误。
	Err error

    // Coalesced 是 Addrs 是否与同时进行相同 DNS 查找的另一个调用者共享。
	Coalesced bool
}
```

​	DNSDoneInfo 包含关于 DNS 查找结果的信息。

### type DNSStartInfo 

``` go 
type DNSStartInfo struct {
	Host string
}
```

​	DNSStartInfo 包含关于 DNS 请求的信息。

### type GotConnInfo 

``` go 
type GotConnInfo struct {
    // Conn 是获取的连接。它由 http.Transport 拥有，
    // 不应由 ClientTrace 的用户读取、写入或关闭。
	Conn net.Conn

    // Reused 是此连接是否已经用于另一个 HTTP 请求。
	Reused bool

    // WasIdle 是此连接是否是从空闲连接池获取的。
	WasIdle bool

    // IdleTime 在 WasIdle 为 true 时报告连接先前闲置的时间。
	IdleTime time.Duration
}
```

​	GotConnInfo 是传递给 ClientTrace.GotConn 函数的实参，包含有关获取的连接的信息。

### type WroteRequestInfo 

``` go 
type WroteRequestInfo struct {
    // Err 是在写入请求时遇到的任何错误。
	Err error
}
```

​	WroteRequestInfo 包含提供给 WroteRequest （ClientTrace结构体中的字段，ClientTrace#WroteRequest 类型为函数）钩子的信息。