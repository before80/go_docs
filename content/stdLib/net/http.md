+++
title = "http"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/http@go1.21.3](https://pkg.go.dev/net/http@go1.21.3)

Package http provides HTTP client and server implementations.

​	`http`包提供了 HTTP 客户端和服务端的实现。

Get, Head, Post, and PostForm make HTTP (or HTTPS) requests:

​	Get、Head、Post 和 PostForm 方法可用于发起 HTTP(或 HTTPS)请求：

```go
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
```

The caller must close the response body when finished with it:

​	在使用完响应后，客户端必须关闭响应主体：

```go
resp, err := http.Get("http://example.com/")
if err != nil {
	// 处理错误
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
// ...
```

For control over HTTP client headers, redirect policy, and other settings, create a Client:

​	为了控制 HTTP 客户端的标头、重定向策略和其他设置，可以创建一个 Client：

```go
client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
```

For control over proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport:

​	为了控制代理、TLS 配置、保持连接、压缩和其他设置，可以创建一个 Transport：

```go
tr := &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
```

Clients and Transports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.

​	Clients 和 Transports 可以被多个 goroutine 并发使用，并且为了效率应该只创建一次并重复使用。

ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:

​	ListenAndServe 方法可启动一个使用给定地址和处理程序的 HTTP 服务器。处理程序通常为 nil，这意味着使用 DefaultServeMux。Handle 和 HandleFunc 方法可将处理程序添加到 DefaultServeMux：

```go
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

More control over the server's behavior is available by creating a custom Server:

​	通过创建自定义服务器可以获得更多对服务器行为的控制：

```go
s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```

Starting with Go 1.6, the http package has transparent support for the HTTP/2 protocol when using HTTPS. Programs that must disable HTTP/2 can do so by setting Transport.TLSNextProto (for clients) or Server.TLSNextProto (for servers) to a non-nil, empty map. Alternatively, the following GODEBUG environment variables are currently supported:

​	从 Go 1.6 开始，当使用 HTTPS 时，http 包对 HTTP/2 协议具有透明支持。必须禁用 HTTP/2 的程序可以通过将 Transport.TLSNextProto(用于客户端)或 Server.TLSNextProto(用于服务器)设置为non-nil的空映射来实现。或者，目前支持以下 GODEBUG 环境变量：

```go
GODEBUG=http2client=0  # disable HTTP/2 client support
GODEBUG=http2server=0  # disable HTTP/2 server support
GODEBUG=http2debug=1   # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2   # ... even more verbose, with frame dumps
```

The GODEBUG variables are not covered by Go's API compatibility promise. Please report any issues before disabling HTTP/2 support: https://golang.org/s/http2bug

​	GODEBUG 变量不受 Go 的 API 兼容性承诺的保护。在禁用 HTTP/2 支持之前，请报告任何问题：https://golang.org/s/http2bug

The http package's Transport and Server both automatically enable HTTP/2 support for simple configurations. To enable HTTP/2 for more complex configurations, to use lower-level HTTP/2 features, or to use a newer version of Go's http2 package, import "golang.org/x/net/http2" directly and use its ConfigureTransport and/or ConfigureServer functions. Manually configuring HTTP/2 via the golang.org/x/net/http2 package takes precedence over the net/http package's built-in HTTP/2 support.

​	http 包的 Transport 和 Server 都会自动为简单配置启用 HTTP/2 支持。为了启用更复杂的配置的 HTTP/2，使用更低级别的 HTTP/2 功能或使用 Go 的 http2 包的更新版本，请直接导入 "golang.org/x/net/http2" 并使用其 ConfigureTransport 和/或 ConfigureServer 函数。手动配置 HTTP/2 通过 golang.org/x/net/http2 包优先于 net/http 包内置的 HTTP/2 支持。




## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/method.go;l=10)

``` go 
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)
```

Common HTTP methods.

​	常见的HTTP方法。

Unless otherwise noted, these are defined in [RFC 7231 section 4.3](https://rfc-editor.org/rfc/rfc7231.html#section-4.3).

​	除非另有说明，这些方法定义在[RFC 7231第4.3节](https://rfc-editor.org/rfc/rfc7231.html#section-4.3)中。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/status.go;l=9)

``` go 
const (
	StatusContinue           = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols = 101 // RFC 9110, 15.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 9110, 15.3.1
	StatusCreated              = 201 // RFC 9110, 15.3.2
	StatusAccepted             = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 9110, 15.3.4
	StatusNoContent            = 204 // RFC 9110, 15.3.5
	StatusResetContent         = 205 // RFC 9110, 15.3.6
	StatusPartialContent       = 206 // RFC 9110, 15.3.7
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices  = 300 // RFC 9110, 15.4.1
	StatusMovedPermanently = 301 // RFC 9110, 15.4.2
	StatusFound            = 302 // RFC 9110, 15.4.3
	StatusSeeOther         = 303 // RFC 9110, 15.4.4
	StatusNotModified      = 304 // RFC 9110, 15.4.5
	StatusUseProxy         = 305 // RFC 9110, 15.4.6

	StatusTemporaryRedirect = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              = 402 // RFC 9110, 15.5.3
	StatusForbidden                    = 403 // RFC 9110, 15.5.4
	StatusNotFound                     = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               = 408 // RFC 9110, 15.5.9
	StatusConflict                     = 409 // RFC 9110, 15.5.10
	StatusGone                         = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            = 417 // RFC 9110, 15.5.18
	StatusTeapot                       = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          = 422 // RFC 9110, 15.5.21
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)
```

HTTP status codes as registered with IANA. See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml

​	HTTP状态码在IANA注册。请参见：[https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml)

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=857)

``` go 
const DefaultMaxHeaderBytes = 1 << 20 // 1 MB
```

DefaultMaxHeaderBytes is the maximum permitted size of the headers in an HTTP request. This can be overridden by setting Server.MaxHeaderBytes.

​	DefaultMaxHeaderBytes是HTTP请求中标头的最大允许大小。这可以通过设置Server.MaxHeaderBytes来覆盖。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transport.go;l=58)

``` go 
const DefaultMaxIdleConnsPerHost = 2
```

DefaultMaxIdleConnsPerHost is the default value of Transport's MaxIdleConnsPerHost.

​	DefaultMaxIdleConnsPerHost是Transport的MaxIdleConnsPerHost的默认值。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=934)

``` go 
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
```

TimeFormat is the time format to use when generating times in HTTP headers. It is like time.RFC1123 but hard-codes GMT as the time zone. The time being formatted must be in UTC for Format to generate the correct format.

​	TimeFormat是生成HTTP标头中时间时要使用的时间格式。它类似于time.RFC1123，但将GMT硬编码为时区。要格式化的时间必须在UTC中，Format才能生成正确的格式。

For parsing this time format, see ParseTime.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=513)

``` go 
const TrailerPrefix = "Trailer:"
```

TrailerPrefix is a magic prefix for ResponseWriter.Header map keys that, if present, signals that the map entry is actually for the response trailers, and not the response headers. The prefix is stripped after the ServeHTTP call finishes and the values are sent in the trailers.

​	TrailerPrefix是ResponseWriter.Header映射键的魔术前缀，如果存在，则表示该映射条目实际上是响应trailer，而不是响应标头。 ServeHTTP调用完成后，将删除前缀并将值发送到trailers中。

This mechanism is intended only for trailers that are not known prior to the headers being written. If the set of trailers is fixed or known before the header is written, the normal Go trailers mechanism is preferred:

​	该机制仅用于不在编写标头之前未知的trailers。如果trailers的集合是固定的或在编写标头之前已知，则首选常规的Go trailers机制：

```
> 原文：[https://pkg.go.dev/net/http#ResponseWriter](https://pkg.go.dev/net/http#ResponseWriter)
> 原文：[https://pkg.go.dev/net/http#example-ResponseWriter-Trailers](https://pkg.go.dev/net/http#example-ResponseWriter-Trailers)
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=51)

``` go 
var (
    // ErrNotSupported indicates that a feature is not supported.
	//
	// It is returned by ResponseController methods to indicate that
	// the handler does not support the method, and by the Push method
	// of Pusher implementations to indicate that HTTP/2 Push support
	// is not available.
	// ErrNotSupported 表示不支持某一特性。
	//
	// 它由 ResponseController 方法返回，表示处理程序不支持该方法，
	// 并且由 Pusher 实现的 Push 方法返回，表示不支持 HTTP/2 推送功能。
	ErrNotSupported = &ProtocolError{"feature not supported"}

    // Deprecated: ErrUnexpectedTrailer is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	// 已弃用：ErrUnexpectedTrailer 
    // 不再由 net/http 包中的任何内容返回。
	// 调用者不应将错误与此变量进行比较。
	ErrUnexpectedTrailer = &ProtocolError{"trailer header without chunked transfer encoding"}

    // ErrMissingBoundary is returned by Request.MultipartReader when the
	// request's Content-Type does not include a "boundary" parameter.
	// 当请求的 Content-Type 不包含 "boundary" 参数时，
	// Request.MultipartReader 将返回 ErrMissingBoundary。
	ErrMissingBoundary = &ProtocolError{"no multipart boundary param in Content-Type"}

    // ErrNotMultipart is returned by Request.MultipartReader when the
	// request's Content-Type is not multipart/form-data.
	// 当请求的 Content-Type 不是 multipart/form-data 时，
	// Request.MultipartReader 将返回 ErrNotMultipart。
	ErrNotMultipart = &ProtocolError{"request Content-Type isn't multipart/form-data"}

    // Deprecated: ErrHeaderTooLong is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	// 已弃用：ErrHeaderTooLong 不再由 net/http 包中的任何内容返回。
	// 调用者不应将错误与此变量进行比较。
	ErrHeaderTooLong = &ProtocolError{"header too long"}

    // Deprecated: ErrShortBody is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	// 已弃用：ErrShortBody 不再由 net/http 包中的任何内容返回。
	// 调用者不应将错误与此变量进行比较。
	ErrShortBody = &ProtocolError{"entity body too short"}

    // Deprecated: ErrMissingContentLength is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	// 已弃用：ErrMissingContentLength 
    // 不再由 net/http 包中的任何内容返回。
	// 调用者不应将错误与此变量进行比较。
	ErrMissingContentLength = &ProtocolError{"missing ContentLength in HEAD response"}
)
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=37)

``` go 
var (
    // ErrBodyNotAllowed is returned by ResponseWriter.Write calls
	// when the HTTP method or response code does not permit a
	// body.
	// ErrBodyNotAllowed 在 ResponseWriter.Write 调用时，
    // 当 HTTP 方法或响应状态码不允许 body 时返回。
	ErrBodyNotAllowed = errors.New("http: request method or response status code does not allow body")

    // ErrHijacked is returned by ResponseWriter.Write calls when
	// the underlying connection has been hijacked using the
	// Hijacker interface. A zero-byte write on a hijacked
	// connection will return ErrHijacked without any other side
	// effects.
	// ErrHijacked 在 ResponseWriter.Write 调用时，
    // 当使用 Hijacker 接口劫持了底层连接时返回。
    // 在劫持的连接上进行零字节写操作会返回 ErrHijacked，没有其他副作用。
	ErrHijacked = errors.New("http: connection has been hijacked")

    // ErrContentLength is returned by ResponseWriter.Write calls
	// when a Handler set a Content-Length response header with a
	// declared size and then attempted to write more bytes than
	// declared.
	// ErrContentLength 在 ResponseWriter.Write 调用时，
    // 当处理程序设置了一个声明大小的 Content-Length 响应头
    // 并尝试写入比声明的更多字节时返回。
	ErrContentLength = errors.New("http: wrote more than the declared Content-Length")

    // Deprecated: ErrWriteAfterFlush is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	// Deprecated: ErrWriteAfterFlush 
    // 不再由 net/http 包中的任何内容返回。
    // 调用方不应将错误与此变量进行比较。
	ErrWriteAfterFlush = errors.New("unused")
)
```

​	HTTP 服务器使用的错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=238)

``` go 
var (
    // ServerContextKey is a context key. It can be used in HTTP
	// handlers with Context.Value to access the server that
	// started the handler. The associated value will be of
	// type *Server.
	// ServerContextKey 是一个 context key。
    // 它可在 HTTP 处理程序中使用 Context.Value 
    // 来访问启动处理程序的服务器。相关的值的类型为 *Server。

	ServerContextKey = &contextKey{"http-server"}

    // LocalAddrContextKey is a context key. It can be used in
	// HTTP handlers with Context.Value to access the local
	// address the connection arrived on.
	// The associated value will be of type net.Addr.
	// LocalAddrContextKey 是一个 context key。
    // 它可在 HTTP 处理程序中使用 Context.Value 
    // 来访问连接到达的本地地址。相关的值的类型为 net.Addr。
	LocalAddrContextKey = &contextKey{"local-addr"}
)
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=110)

``` go 
var DefaultClient = &Client{}
```

DefaultClient is the default Client and is used by Get, Head, and Post.

​	DefaultClient 是默认的 Client，并被 Get、Head 和 Post 使用。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2322)

``` go 
var DefaultServeMux = &defaultServeMux
```

DefaultServeMux is the default ServeMux used by Serve.

​	DefaultServeMux 是 Serve 使用的默认 ServeMux。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=1826)

``` go 
var ErrAbortHandler = errors.New("net/http: abort Handler")
```

ErrAbortHandler is a sentinel panic value to abort a handler. While any panic from ServeHTTP aborts the response to the client, panicking with ErrAbortHandler also suppresses logging of a stack trace to the server's error log.

​	ErrAbortHandler 是一个特殊的 panic 值，用于中止处理程序。虽然任何从 ServeHTTP 中发生的 panic 都会中止向客户端的响应，但是使用 ErrAbortHandler 发生 panic 还会阻止将堆栈跟踪记录到服务器的错误日志中。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transfer.go;l=823)

``` go 
var ErrBodyReadAfterClose = errors.New("http: invalid Read on closed Body")
```

ErrBodyReadAfterClose is returned when reading a Request or Response Body after the body has been closed. This typically happens when the body is read after an HTTP Handler calls WriteHeader or Write on its ResponseWriter.

​	当请求或响应主体关闭后继续读取主体时，返回 ErrBodyReadAfterClose。这通常发生在 HTTP 处理程序在其 ResponseWriter 上调用 WriteHeader 或 Write 后读取主体时。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3356)

``` go 
var ErrHandlerTimeout = errors.New("http: Handler timeout")
```

ErrHandlerTimeout is returned on ResponseWriter Write calls in handlers which have timed out.

​	在处理超时的处理程序的 ResponseWriter Write 调用上返回 ErrHandlerTimeout。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transfer.go;l=29)

``` go 
var ErrLineTooLong = internal.ErrLineTooLong
```

ErrLineTooLong is returned when reading request or response bodies with malformed chunked encoding.

​	当使用格式错误的分块编码读取请求或响应主体时返回 ErrLineTooLong。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=39)

``` go 
var ErrMissingFile = errors.New("http: no such file")
```

ErrMissingFile is returned by FormFile when the provided file field name is either not present in the request or not a file field.

​	当请求中提供的文件字段名称不存在或不是文件字段时，由 FormFile 返回 ErrMissingFile。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=415)

``` go 
var ErrNoCookie = errors.New("http: named cookie not present")
```

ErrNoCookie is returned by Request's Cookie method when a cookie is not found.

​	当找不到 cookie 时，Request 的 Cookie 方法会返回 ErrNoCookie。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/response.go;l=131)

``` go 
var ErrNoLocation = errors.New("http: no Location header in response")
```

ErrNoLocation is returned by Response's Location method when no Location header is present.

​	当响应中没有 Location 头时，Response 的 Location 方法会返回 ErrNoLocation。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3017)

``` go 
var ErrServerClosed = errors.New("http: Server closed")
```

ErrServerClosed is returned by the Server's Serve, ServeTLS, ListenAndServe, and ListenAndServeTLS methods after a call to Shutdown or Close.

​	ErrServerClosed在Server的Serve、ServeTLS、ListenAndServe和ListenAndServeTLS方法调用Shutdown或Close后返回。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transport.go;l=741)

``` go 
var ErrSkipAltProtocol = errors.New("net/http: skip alternate protocol")
```

ErrSkipAltProtocol is a sentinel error value defined by Transport.RegisterProtocol.

​	ErrSkipAltProtocol是Transport.RegisterProtocol定义的标志性错误值。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=489)

``` go 
var ErrUseLastResponse = errors.New("net/http: use last response")
```

ErrUseLastResponse can be returned by Client.CheckRedirect hooks to control how redirects are processed. If returned, the next request is not sent and the most recent response is returned with its body unclosed.

​	如果Client.CheckRedirect钩子返回ErrUseLastResponse，则可以控制如何处理重定向。 如果返回此值，则不会发送下一个请求，并使用其正文未关闭的最近响应返回。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/http.go;l=104)

``` go 
var NoBody = noBody{}
```

NoBody is an io.ReadCloser with no bytes. Read always returns EOF and Close always returns nil. It can be used in an outgoing client request to explicitly signal that a request has zero bytes. An alternative, however, is to simply set Request.Body to nil.

​	NoBody是一个没有字节的io.ReadCloser。Read始终返回EOF，Close始终返回nil。 它可以在传出的客户端请求中使用，以显式地表示请求没有字节。 但是，另一种方法是将Request.Body设置为nil。

## 函数

### func CanonicalHeaderKey 

``` go 
func CanonicalHeaderKey(s string) string
```

CanonicalHeaderKey returns the canonical format of the header key s. The canonicalization converts the first letter and any letter following a hyphen to upper case; the rest are converted to lowercase. For example, the canonical key for "accept-encoding" is "Accept-Encoding". If s contains a space or invalid header field bytes, it is returned without modifications.

​	CanonicalHeaderKey返回标头键s的规范格式。规范化将第一个字母和任何连字符后面的字母转换为大写字母；其余字母转换为小写字母。例如，"accept-encoding"的规范键是"Accept-Encoding"。如果s包含空格或无效的标头字段字节，则返回不带修改的s。

### func DetectContentType 

``` go 
func DetectContentType(data []byte) string
```

DetectContentType implements the algorithm described at https://mimesniff.spec.whatwg.org/ to determine the Content-Type of the given data. It considers at most the first 512 bytes of data. DetectContentType always returns a valid MIME type: if it cannot determine a more specific one, it returns "application/octet-stream".

​	DetectContentType实现在https://mimesniff.spec.whatwg.org/上描述的算法，以确定给定数据的Content-Type。它最多考虑前512个字节的数据。DetectContentType始终返回有效的MIME类型：如果无法确定更具体的类型，则返回"application/octet-stream"。

### func Error 

``` go 
func Error(w ResponseWriter, error string, code int)
```

Error replies to the request with the specified error message and HTTP code. It does not otherwise end the request; the caller should ensure no further writes are done to w. The error message should be plain text.

​	Error使用指定的错误消息和HTTP代码回复请求。它不会以其他方式结束请求；调用者应确保不会对w进行进一步的写入。错误消息应为纯文本。

### func Handle 

``` go 
func Handle(pattern string, handler Handler)
```

Handle registers the handler for the given pattern in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched.

​	Handle 在 DefaultServeMux 中为给定的 pattern 注册 handler。ServeMux 的文档解释了如何匹配 pattern。

#### Handle Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

### func HandleFunc 

``` go 
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

HandleFunc registers the handler function for the given pattern in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched.

​	HandleFunc 在 DefaultServeMux 中为给定的 pattern 注册 handler function。ServeMux 的文档解释了如何匹配 pattern。

#### HandleFunc Example
``` go 
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

### func ListenAndServe 

``` go 
func ListenAndServe(addr string, handler Handler) error
```

ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.

​	ListenAndServe 监听 TCP 网络地址 addr，并在传入连接上调用 Serve 以处理请求。已接受的连接已配置为启用 TCP keep-alives。

The handler is typically nil, in which case the DefaultServeMux is used.

​	通常 handler 是 nil，在这种情况下会使用 DefaultServeMux。

ListenAndServe always returns a non-nil error.

​	ListenAndServe 总是返回非 nil 错误。

#### ListenAndServe Example
``` go 
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

### func ListenAndServeTLS 

``` go 
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```

ListenAndServeTLS acts identically to ListenAndServe, except that it expects HTTPS connections. Additionally, files containing a certificate and matching private key for the server must be provided. If the certificate is signed by a certificate authority, the certFile should be the concatenation of the server's certificate, any intermediates, and the CA's certificate.

​	ListenAndServeTLS 的行为与 ListenAndServe 相同，除了它预期 HTTPS 连接。另外，必须提供包含服务器证书和匹配私钥的文件。如果证书由证书颁发机构签署，则 certFile 应该是服务器证书、任何中间文件以及 CA 的证书的连接。

#### ListenAndServeTLS Example
``` go 
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, TLS!\n")
	})

	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/")
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	log.Fatal(err)
}

```

### func MaxBytesReader 

``` go 
func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser
```

MaxBytesReader is similar to io.LimitReader but is intended for limiting the size of incoming request bodies. In contrast to io.LimitReader, MaxBytesReader's result is a ReadCloser, returns a non-nil error of type *MaxBytesError for a Read beyond the limit, and closes the underlying reader when its Close method is called.

​	MaxBytesReader 类似于 io.LimitReader，但旨在限制传入请求主体的大小。与 io.LimitReader 不同，MaxBytesReader 的结果是一个 ReadCloser，在超出限制的情况下返回类型为 *MaxBytesError 的非 nil 错误，并在调用其 Close 方法时关闭底层读取器。

MaxBytesReader prevents clients from accidentally or maliciously sending a large request and wasting server resources. If possible, it tells the ResponseWriter to close the connection after the limit has been reached.

​	MaxBytesReader 防止客户端意外或恶意发送大型请求并浪费服务器资源。如果可能，它会告诉 ResponseWriter 在达到限制后关闭连接。

### func NotFound 

``` go 
func NotFound(w ResponseWriter, r *Request)
```

NotFound replies to the request with an HTTP 404 not found error.

​	NotFound 函数返回一个 HTTP 404 错误响应。

### func ParseHTTPVersion 

``` go 
func ParseHTTPVersion(vers string) (major, minor int, ok bool)
```

ParseHTTPVersion parses an HTTP version string according to [RFC 7230, section 2.6](https://rfc-editor.org/rfc/rfc7230.html#section-2.6). "HTTP/1.0" returns (1, 0, true). Note that strings without a minor version, such as "HTTP/2", are not valid.

​	ParseHTTPVersion 函数按照 [RFC 7230 第 2.6 节](https://rfc-editor.org/rfc/rfc7230.html#section-2.6)解析 HTTP 版本字符串。例如 "HTTP/1.0" 会返回 (1, 0, true)。请注意，不带次要版本号的字符串，例如 "HTTP/2"，是无效的。

### func ParseTime  <- go1.1

``` go 
func ParseTime(text string) (t time.Time, err error)
```

ParseTime parses a time header (such as the Date: header), trying each of the three formats allowed by HTTP/1.1: TimeFormat, time.RFC850, and time.ANSIC.

​	ParseTime 函数解析时间头(如 Date 头)，尝试 HTTP/1.1 所允许的三种格式：TimeFormat、time.RFC850 和 time.ANSIC。

### func ProxyFromEnvironment 

``` go 
func ProxyFromEnvironment(req *Request) (*url.URL, error)
```

ProxyFromEnvironment returns the URL of the proxy to use for a given request, as indicated by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions thereof). Requests use the proxy from the environment variable matching their scheme, unless excluded by NO_PROXY.

​	ProxyFromEnvironment 函数返回给定请求使用的代理 URL，该 URL 由环境变量 HTTP_PROXY、HTTPS_PROXY 和 NO_PROXY(或它们的小写形式)指示。请求使用与其方案匹配的环境变量中的代理，除非被 NO_PROXY 排除。

The environment values may be either a complete URL or a "host[:port]", in which case the "http" scheme is assumed. The schemes "http", "https", and "socks5" are supported. An error is returned if the value is a different form.

​	环境变量的值可以是完整的 URL，也可以是 "host[:port]" 形式，在这种情况下会假定为 "http" 方案。支持 "http"、"https" 和 "socks5" 方案。如果值是其他形式，则返回错误。

A nil URL and nil error are returned if no proxy is defined in the environment, or a proxy should not be used for the given request, as defined by NO_PROXY.

​	如果环境中没有定义代理，或者请求不应使用代理，则返回 nil URL 和 nil 错误。

As a special case, if req.URL.Host is "localhost" (with or without a port number), then a nil URL and nil error will be returned.

​	作为特例，如果 req.URL.Host 是 "localhost"(带或不带端口号)，则会返回 nil URL 和 nil 错误。

### func ProxyURL 

``` go 
func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error)
```

ProxyURL returns a proxy function (for use in a Transport) that always returns the same URL.

​	ProxyURL 函数返回一个代理函数(供 Transport 使用)，该函数始终返回相同的 URL。

### func Redirect 

``` go 
func Redirect(w ResponseWriter, r *Request, url string, code int)
```

Redirect replies to the request with a redirect to url, which may be a path relative to the request path.

​	Redirect 函数将请求重定向到 url，url 可以是相对于请求路径的路径。

The provided code should be in the 3xx range and is usually StatusMovedPermanently, StatusFound or StatusSeeOther.

​	提供的 code 应该在 3xx 范围内，通常为 StatusMovedPermanently、StatusFound 或 StatusSeeOther。

If the Content-Type header has not been set, Redirect sets it to "text/html; charset=utf-8" and writes a small HTML body. Setting the Content-Type header to any value, including nil, disables that behavior.

​	如果尚未设置 Content-Type 标头，则 Redirect 函数会将其设置为 "text/html; charset=utf-8" 并写入小的 HTML 主体。设置 Content-Type 标头为任何值(包括 nil)都会禁用该行为。

### func Serve 

``` go 
func Serve(l net.Listener, handler Handler) error
```

Serve accepts incoming HTTP connections on the listener l, creating a new service goroutine for each. The service goroutines read requests and then call handler to reply to them.

​	Serve 函数在监听器 l 上接受传入的 HTTP 连接，为每个连接创建一个新的服务协程。服务协程读取请求，然后调用 handler 进行响应。

The handler is typically nil, in which case the DefaultServeMux is used.

​	handler 通常为 nil，此时会使用 DefaultServeMux。

HTTP/2 support is only enabled if the Listener returns *tls.Conn connections and they were configured with "h2" in the TLS Config.NextProtos.

​	只有当 Listener 返回 `*tls.Conn` 连接且它们在 TLS Config.NextProtos 中配置为 "h2" 时，才启用 HTTP/2 支持。

Serve always returns a non-nil error.

​	Serve 函数始终返回非 nil 错误。

### func ServeContent 

``` go 
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
```

ServeContent replies to the request using the content in the provided ReadSeeker. The main benefit of ServeContent over io.Copy is that it handles Range requests properly, sets the MIME type, and handles If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since, and If-Range requests.

​	ServeContent 使用 io.ReadSeeker 中提供的内容回复请求。ServeContent 相对于 io.Copy 的主要好处在于正确处理范围请求、设置 MIME 类型以及处理 If-Match、If-Unmodified-Since、If-None-Match、If-Modified-Since 和 If-Range 请求。

If the response's Content-Type header is not set, ServeContent first tries to deduce the type from name's file extension and, if that fails, falls back to reading the first block of the content and passing it to DetectContentType. The name is otherwise unused; in particular it can be empty and is never sent in the response.

​	如果响应的 Content-Type 头未设置，ServeContent 首先尝试从 name 的文件扩展名推断类型，如果失败，则返回到读取内容的第一个块并将其传递给 DetectContentType。否则，不使用 name；特别地，它可以为空，并且永远不会在响应中发送。

If modtime is not the zero time or Unix epoch, ServeContent includes it in a Last-Modified header in the response. If the request includes an If-Modified-Since header, ServeContent uses modtime to decide whether the content needs to be sent at all.

​	如果 modtime 不是零时间或 Unix 纪元，则 ServeContent 在响应中包括一个 Last-Modified 头。如果请求包括一个 If-Modified-Since 头，则 ServeContent 使用 modtime 来决定是否需要发送内容。

The content's Seek method must work: ServeContent uses a seek to the end of the content to determine its size.

​	内容的 Seek 方法必须工作：ServeContent 使用 seek 到内容的末尾以确定其大小。

If the caller has set w's ETag header formatted per [RFC 7232, section 2.3](https://rfc-editor.org/rfc/rfc7232.html#section-2.3), ServeContent uses it to handle requests using If-Match, If-None-Match, or If-Range.

​	如果调用方按 [RFC 7232，第 2.3 节](https://rfc-editor.org/rfc/rfc7232.html#section-2.3)格式化了 w 的 ETag 头，则 ServeContent 使用它来处理使用 If-Match、If-None-Match 或 If-Range 的请求。

Note that *os.File implements the io.ReadSeeker interface.

​	请注意，`*os.File` 实现了 io.ReadSeeker 接口。

### func ServeFile 

``` go 
func ServeFile(w ResponseWriter, r *Request, name string)
```

ServeFile replies to the request with the contents of the named file or directory.

​	ServeFile 用指定的文件或目录的内容回复请求。

If the provided file or directory name is a relative path, it is interpreted relative to the current directory and may ascend to parent directories. If the provided name is constructed from user input, it should be sanitized before calling ServeFile.

​	如果提供的文件或目录名是相对路径，则相对于当前目录进行解释，可能升到父目录。如果提供的名称由用户输入构造，则在调用 ServeFile 之前应进行清理。

As a precaution, ServeFile will reject requests where r.URL.Path contains a ".." path element; this protects against callers who might unsafely use filepath.Join on r.URL.Path without sanitizing it and then use that filepath.Join result as the name argument.

​	作为一项预防措施，ServeFile 将拒绝 r.URL.Path 包含 ".." 路径元素的请求；这可保护免受可能会不安全地对 r.URL.Path 使用 filepath.Join 而未对其进行清理的调用者的侵害，并使用 filepath.Join 结果作为名称参数。

As another special case, ServeFile redirects any request where r.URL.Path ends in "/index.html" to the same path, without the final "index.html". To avoid such redirects either modify the path or use ServeContent.

​	作为另一个特殊情况，ServeFile 将任何以 "/index.html" 结尾的 r.URL.Path 的请求重定向到相同的路径，不包括最后的 "index.html"。为避免此类重定向，请修改路径或使用 ServeContent。

Outside of those two special cases, ServeFile does not use r.URL.Path for selecting the file or directory to serve; only the file or directory provided in the name argument is used.

​	除了这两个特殊情况外，ServeFile 不使用 r.URL.Path 来选择要提供的文件或目录；只使用名称参数中提供的文件或目录。

### func ServeFileFS <- go1.22.0

```
func ServeFileFS(w ResponseWriter, r *Request, fsys fs.FS, name string)
```

ServeFileFS replies to the request with the contents of the named file or directory from the file system fsys. The files provided by fsys must implement [io.Seeker](https://pkg.go.dev/io#Seeker).

​	ServeFileFS 使用文件系统 `fsys` 中指定文件或目录的内容来回复请求。`fsys` 提供的文件必须实现 [io.Seeker](https://pkg.go.dev/io#Seeker) 接口。

If the provided name is constructed from user input, it should be sanitized before calling [ServeFileFS](https://pkg.go.dev/net/http@go1.23.0#ServeFileFS).

​	如果提供的文件名是由用户输入生成的，在调用 [ServeFileFS](https://pkg.go.dev/net/http@go1.23.0#ServeFileFS) 之前应对其进行清理。

As a precaution, ServeFileFS will reject requests where r.URL.Path contains a ".." path element; this protects against callers who might unsafely use [filepath.Join](https://pkg.go.dev/path/filepath#Join) on r.URL.Path without sanitizing it and then use that filepath.Join result as the name argument.

​	作为预防措施，ServeFileFS 将拒绝 `r.URL.Path` 中包含 ".." 路径元素的请求；这可以防止调用者在未清理 `r.URL.Path` 的情况下使用 [filepath.Join](https://pkg.go.dev/path/filepath#Join) 并将该 `filepath.Join` 结果作为 `name` 参数使用。

As another special case, ServeFileFS redirects any request where r.URL.Path ends in "/index.html" to the same path, without the final "index.html". To avoid such redirects either modify the path or use [ServeContent](https://pkg.go.dev/net/http@go1.23.0#ServeContent).

​	作为另一种特殊情况，ServeFileFS 会将 `r.URL.Path` 以 "/index.html" 结尾的请求重定向到相同路径，但去掉最后的 "index.html"。要避免这种重定向，可以修改路径或使用 [ServeContent](https://pkg.go.dev/net/http@go1.23.0#ServeContent)。

Outside of those two special cases, ServeFileFS does not use r.URL.Path for selecting the file or directory to serve; only the file or directory provided in the name argument is used.

​	除了这两种特殊情况之外，ServeFileFS 不会使用 `r.URL.Path` 来选择要提供的文件或目录；只会使用 `name` 参数中提供的文件或目录。

### func ServeTLS  <- go1.9

``` go 
func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error
```

ServeTLS accepts incoming HTTPS connections on the listener l, creating a new service goroutine for each. The service goroutines read requests and then call handler to reply to them.

​	ServeTLS 函数在监听器 l 上接受传入的 HTTPS 连接，为每个连接创建一个新的服务协程。服务协程读取请求，然后调用处理程序来回复请求。

The handler is typically nil, in which case the DefaultServeMux is used.

​	处理程序通常为 nil，此时将使用 DefaultServeMux。

Additionally, files containing a certificate and matching private key for the server must be provided. If the certificate is signed by a certificate authority, the certFile should be the concatenation of the server's certificate, any intermediates, and the CA's certificate.

​	另外，必须提供包含服务器证书和匹配的私钥的文件。如果证书由证书颁发机构签名，则 certFile 应该是服务器证书、任何中间证书和 CA 证书的连接。

ServeTLS always returns a non-nil error.

​	ServeTLS 函数总是返回非 nil 的错误。

### func SetCookie 

``` go 
func SetCookie(w ResponseWriter, cookie *Cookie)
```

SetCookie adds a Set-Cookie header to the provided ResponseWriter's headers. The provided cookie must have a valid Name. Invalid cookies may be silently dropped.

​	SetCookie 函数将 Set-Cookie 头添加到提供的 ResponseWriter 的头中。提供的 cookie 必须有一个有效的名称。无效的 cookie 可能会被静默丢弃。

### func StatusText 

``` go 
func StatusText(code int) string
```

StatusText returns a text for the HTTP status code. It returns the empty string if the code is unknown.

​	StatusText 函数返回 HTTP 状态码的文本。如果状态码未知，则返回空字符串。

## 类型

### type Client 

``` go 
type Client struct {
    // Transport specifies the mechanism by which individual
	// HTTP requests are made.
	// If nil, DefaultTransport is used.
	// Transport 指定单个 HTTP 请求的执行机制。
	// 如果为 nil，则使用 DefaultTransport。
	Transport RoundTripper

    // CheckRedirect specifies the policy for handling redirects.
	// If CheckRedirect is not nil, the client calls it before
	// following an HTTP redirect. The arguments req and via are
	// the upcoming request and the requests made already, oldest
	// first. If CheckRedirect returns an error, the Client's Get
	// method returns both the previous Response (with its Body
	// closed) and CheckRedirect's error (wrapped in a url.Error)
	// instead of issuing the Request req.
	// As a special case, if CheckRedirect returns ErrUseLastResponse,
	// then the most recent response is returned with its body
	// unclosed, along with a nil error.
	//
	// If CheckRedirect is nil, the Client uses its default policy,
	// which is to stop after 10 consecutive requests.
	// CheckRedirect 指定处理重定向的策略。
	// 如果 CheckRedirect 不为 nil，
    // 则客户端在跟随 HTTP 重定向之前调用它。
	// 参数 req 和 via 分别是即将到来的请求和已经发出的请求，
    // 最旧的请求在最前面。
	// 如果 CheckRedirect 返回一个错误，
    // 则 Get 方法返回之前的 Response(其 Body 已关闭)和 	
    // CheckRedirect 的错误(封装在 url.Error 中)，
    // 而不是发出请求 req。
	// 作为特殊情况，如果 CheckRedirect 返回 ErrUseLastResponse，
    // 则返回最近的 Response，并保持其 Body 未关闭，同时返回空错误。
	//
	// 如果 CheckRedirect 为 nil，则客户端使用其默认策略，
    // 即在连续 10 次请求后停止。
	CheckRedirect func(req *Request, via []*Request) error

    // Jar specifies the cookie jar.
	//
	// The Jar is used to insert relevant cookies into every
	// outbound Request and is updated with the cookie values
	// of every inbound Response. The Jar is consulted for every
	// redirect that the Client follows.
	//
	// If Jar is nil, cookies are only sent if they are explicitly
	// set on the Request.
	// Jar 指定 cookie 存储。
	//
	// Jar 用于在每个出站请求中插入相关的 cookies，
    // 并使用每个入站 Response 的 cookie 值进行更新。
    // Jar 用于处理客户端跟随的每个重定向。
	//
	// 如果 Jar 为 nil，则仅在请求上显式设置了 cookie 时才发送它们。
	Jar CookieJar

    // Timeout specifies a time limit for requests made by this
	// Client. The timeout includes connection time, any
	// redirects, and reading the response body. The timer remains
	// running after Get, Head, Post, or Do return and will
	// interrupt reading of the Response.Body.
	//
	// A Timeout of zero means no timeout.
	//
	// The Client cancels requests to the underlying Transport
	// as if the Request's Context ended.
	//
	// For compatibility, the Client will also use the deprecated
	// CancelRequest method on Transport if found. New
	// RoundTripper implementations should use the Request's Context
	// for cancellation instead of implementing CancelRequest.
	// Timeout 指定此客户端发出的请求的时间限制。
	// 超时包括连接时间、任何重定向和读取响应正文。
    // 定时器在 Get、Head、Post 或 Do 返回之后仍在运行，
    // 并将中断读取 Response.Body。
	//
	// Timeout 为零表示没有超时。
	//
	// 如同请求的 Context 结束一样，客户端会取消底层传输的请求。
	//
	// 为了兼容性，
    // 如果 Transport 中找到了废弃的 CancelRequest 方法，
    // 则客户端也会使用它。
    // 新的 RoundTripper 实现应使用请求的 Context 进行取消，
    // 而不是实现 CancelRequest。
	Timeout time.Duration
}
```

A Client is an HTTP client. Its zero value (DefaultClient) is a usable client that uses DefaultTransport.

​	一个 Client 是一个 HTTP 客户端。它的零值(DefaultClient)是一个可用的客户端，它使用 DefaultTransport。

The Client's Transport typically has internal state (cached TCP connections), so Clients should be reused instead of created as needed. Clients are safe for concurrent use by multiple goroutines.

​	Client 的 Transport 通常具有内部状态(缓存的 TCP 连接)，因此应该重复使用 Client，而不是按需创建。Client 可以被多个 goroutine 并发使用。

A Client is higher-level than a RoundTripper (such as Transport) and additionally handles HTTP details such as cookies and redirects.

​	Client 比 RoundTripper(如 Transport)更高级，还处理诸如 cookie 和重定向之类的 HTTP 细节。

When following redirects, the Client will forward all headers set on the initial Request except:

​	在遵循重定向时，Client 将转发在初始请求上设置的所有标头，但除了以下情况：

- when forwarding sensitive headers like "Authorization", "WWW-Authenticate", and "Cookie" to untrusted targets. These headers will be ignored when following a redirect to a domain that is not a subdomain match or exact match of the initial domain. For example, a redirect from "foo.com" to either "foo.com" or "sub.foo.com" will forward the sensitive headers, but a redirect to "bar.com" will not.

- 将敏感标头(例如"Authorization"、"WWW-Authenticate"和"Cookie")转发到不受信任的目标时。在重定向到与初始域不是子域匹配或精确匹配的域时，将忽略这些标头。例如，从"foo.com"重定向到"foo.com"或"sub.foo.com"将转发敏感标头，但是重定向到"bar.com"将不会。

- when forwarding the "Cookie" header with a non-nil cookie Jar. Since each redirect may mutate the state of the cookie jar, a redirect may possibly alter a cookie set in the initial request. When forwarding the "Cookie" header, any mutated cookies will be omitted, with the expectation that the Jar will insert those mutated cookies with the updated values (assuming the origin matches). If Jar is nil, the initial cookies are forwarded without change.

-  使用非空 cookie Jar 转发"Cookie"标头时。由于每个重定向可能会更改 cookie jar 的状态，重定向可能会更改在初始请求中设置的 cookie。在转发"Cookie"标头时，任何已更改的 cookie 都将被省略，预期 Jar 将使用更新后的值插入这些已更改的 cookie(假设原点匹配)。如果 Jar 为 nil，则初始 cookie 会不加更改地转发。

#### (*Client) CloseIdleConnections  <- go1.12

``` go 
func (c *Client) CloseIdleConnections()
```

CloseIdleConnections closes any connections on its Transport which were previously connected from previous requests but are now sitting idle in a "keep-alive" state. It does not interrupt any connections currently in use.

​	CloseIdleConnections 关闭 Transport 上任何连接，这些连接先前从先前的请求连接，但现在处于"keep-alive"状态。它不会中断当前正在使用的任何连接。

If the Client's Transport does not have a CloseIdleConnections method then this method does nothing.

​	如果 Client 的 Transport 没有 CloseIdleConnections 方法，则此方法不执行任何操作。

#### (*Client) Do 

``` go 
func (c *Client) Do(req *Request) (*Response, error)
```

Do sends an HTTP request and returns an HTTP response, following policy (such as redirects, cookies, auth) as configured on the client.

​	Do 发送 HTTP 请求并返回 HTTP 响应，遵循客户端配置的策略(例如重定向、cookie、身份验证)。

An error is returned if caused by client policy (such as CheckRedirect), or failure to speak HTTP (such as a network connectivity problem). A non-2xx status code doesn't cause an error.

​	如果由客户端策略(例如 CheckRedirect)或无法进行 HTTP 通信(例如网络连接问题)导致错误，则返回错误。非 2xx 状态码不会导致错误。

If the returned error is nil, the Response will contain a non-nil Body which the user is expected to close. If the Body is not both read to EOF and closed, the Client's underlying RoundTripper (typically Transport) may not be able to re-use a persistent TCP connection to the server for a subsequent "keep-alive" request.

​	如果返回的错误是 nil，那么 Response 将包含一个非 nil 的 Body，用户应该在读取完后关闭它。如果 Body 没有读取到 EOF 并且关闭，Client 的底层 RoundTripper(通常是 Transport)可能无法重用持久的 TCP 连接，用于后续的 "keep-alive" 请求。

The request Body, if non-nil, will be closed by the underlying Transport, even on errors.

​	如果请求的 Body 非 nil，则在底层的 Transport 上即使发生错误，也会关闭它。

On error, any Response can be ignored. A non-nil Response with a non-nil error only occurs when CheckRedirect fails, and even then the returned Response.Body is already closed.

​	在出现错误时，可以忽略任何 Response。当 CheckRedirect 失败时，非 nil 的 Response 和非 nil 的 error 仅会出现一次。即使出现这种情况，返回的 Response.Body 也已经关闭。

Generally Get, Post, or PostForm will be used instead of Do.

​	通常会使用 Get、Post 或 PostForm 而不是 Do。

If the server replies with a redirect, the Client first uses the CheckRedirect function to determine whether the redirect should be followed. If permitted, a 301, 302, or 303 redirect causes subsequent requests to use HTTP method GET (or HEAD if the original request was HEAD), with no body. A 307 or 308 redirect preserves the original HTTP method and body, provided that the Request.GetBody function is defined. The NewRequest function automatically sets GetBody for common standard library body types.

​	如果服务器回复重定向，则客户端首先使用 CheckRedirect 函数来确定是否应该跟随重定向。如果被允许，则 301、302 或 303 重定向会导致后续请求使用 HTTP 方法 GET(或 HEAD，如果原始请求是 HEAD)，而不使用 Body。307 或 308 重定向会保留原始的 HTTP 方法和 Body，前提是定义了 Request.GetBody 函数。NewRequest 函数会自动为常见的标准库 Body 类型设置 GetBody。

Any returned error will be of type *url.Error. The url.Error value's Timeout method will report true if the request timed out.

​	任何返回的错误都将是 `*url.Error` 类型。url.Error 值的 Timeout 方法将在请求超时时报告 true。

#### (*Client) Get 

``` go 
func (c *Client) Get(url string) (resp *Response, err error)
```

Get issues a GET to the specified URL. If the response is one of the following redirect codes, Get follows the redirect after calling the Client's CheckRedirect function:

​	Get 向指定的 URL 发送一个 GET 请求。如果响应是以下重定向代码之一，则 Get 在调用 Client 的 CheckRedirect 函数后跟随重定向：

```
301 (Moved Permanently)(永久移动)
302 (Found)(找到)
303 (See Other)(参见其他)
307 (Temporary Redirect)(临时重定向)
308 (Permanent Redirect)(永久重定向)
```

An error is returned if the Client's CheckRedirect function fails or if there was an HTTP protocol error. A non-2xx response doesn't cause an error. Any returned error will be of type *url.Error. The url.Error value's Timeout method will report true if the request timed out.

​	如果 Client 的 CheckRedirect 函数失败或存在 HTTP 协议错误，则返回错误。非 2xx 响应不会导致错误。任何返回的错误都将是 `*url.Error` 类型。url.Error 值的 Timeout 方法将在请求超时时报告 true。

When err is nil, resp always contains a non-nil resp.Body. Caller should close resp.Body when done reading from it.

​	当 err 为 nil 时，resp 总是包含一个非 nil 的 resp.Body。在读取完它后，调用者应该关闭 resp.Body。

To make a request with custom headers, use NewRequest and Client.Do.

​	要使用自定义标头发出请求，请使用 NewRequest 和 Client.Do。

To make a request with a specified context.Context, use NewRequestWithContext and Client.Do.

​	要使用指定的 context.Context 发出请求，请使用 NewRequestWithContext 和 Client.Do。

#### (*Client) Head 

``` go 
func (c *Client) Head(url string) (resp *Response, err error)
```

Head issues a HEAD to the specified URL. If the response is one of the following redirect codes, Head follows the redirect after calling the Client's CheckRedirect function:

​	Head 向指定的 URL 发出一个 HEAD 请求。如果响应是以下重定向代码之一，则在调用 Client 的 CheckRedirect 函数后，Head 跟随重定向：

```
301 (Moved Permanently)(永久移动)
302 (Found)(找到)
303 (See Other)(参见其他)
307 (Temporary Redirect)(临时重定向)
308 (Permanent Redirect)(永久重定向)
```

To make a request with a specified context.Context, use NewRequestWithContext and Client.Do.

​	使用NewRequestWithContext和Client.Do可以指定context.Context进行请求。

#### (*Client) Post 

``` go 
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
```

Post issues a POST to the specified URL.

​	Post向指定的URL发出POST请求。

Caller should close resp.Body when done reading from it.

​	调用者在完成读取后应关闭resp.Body。

If the provided body is an io.Closer, it is closed after the request.

​	如果提供的body是io.Closer，则在请求后关闭它。

To set custom headers, use NewRequest and Client.Do.

​	要设置自定义标头，请使用NewRequest和Client.Do。

To make a request with a specified context.Context, use NewRequestWithContext and Client.Do.

​	要使用指定的context.Context进行请求，请使用NewRequestWithContext和Client.Do。

See the Client.Do method documentation for details on how redirects are handled.

​	有关如何处理重定向的详细信息，请参阅Client.Do方法文档。

#### (*Client) PostForm 

``` go 
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
```

PostForm issues a POST to the specified URL, with data's keys and values URL-encoded as the request body.

​	PostForm将数据的键和值作为请求正文进行URL编码，并向指定的URL发出POST请求。

The Content-Type header is set to application/x-www-form-urlencoded. To set other headers, use NewRequest and Client.Do.

​	Content-Type标头设置为application/x-www-form-urlencoded。要设置其他标头，请使用NewRequest和Client.Do。

When err is nil, resp always contains a non-nil resp.Body. Caller should close resp.Body when done reading from it.

​	当err为nil时，resp始终包含非nil resp.Body。完成读取后，调用者应关闭resp.Body。

See the Client.Do method documentation for details on how redirects are handled.

​	有关如何处理重定向的详细信息，请参阅Client.Do方法文档。

To make a request with a specified context.Context, use NewRequestWithContext and Client.Do.

​	要使用指定的context.Context进行请求，请使用NewRequestWithContext和Client.Do。

### type CloseNotifier <- DEPRECATED

```go
type CloseNotifier interface {
	// CloseNotify returns a channel that receives at most a
	// single value (true) when the client connection has gone
	// away.
	//
	// CloseNotify may wait to notify until Request.Body has been
	// fully read.
	//
	// After the Handler has returned, there is no guarantee
	// that the channel receives a value.
	//
	// If the protocol is HTTP/1.1 and CloseNotify is called while
	// processing an idempotent request (such a GET) while
	// HTTP/1.1 pipelining is in use, the arrival of a subsequent
	// pipelined request may cause a value to be sent on the
	// returned channel. In practice HTTP/1.1 pipelining is not
	// enabled in browsers and not seen often in the wild. If this
	// is a problem, use HTTP/2 or only use CloseNotify on methods
	// such as POST.
	CloseNotify() <-chan bool
}
```

The CloseNotifier interface is implemented by ResponseWriters which allow detecting when the underlying connection has gone away.

This mechanism can be used to cancel long operations on the server if the client has disconnected before the response is ready.

Deprecated: the CloseNotifier interface predates Go's context package. New code should use Request.Context instead.

### type ConnState  <- go1.3

``` go 
type ConnState int
```

A ConnState represents the state of a client connection to a server. It's used by the optional Server.ConnState hook.

​	ConnState表示与服务器的客户端连接状态。它由可选的Server.ConnState钩子使用。

``` go 
const (
    // StateNew represents a new connection that is expected to
	// send a request immediately. Connections begin at this
	// state and then transition to either StateActive or
	// StateClosed.
	// StateNew 表示一个新的连接，该连接预计会立即发送请求。
    // 连接从此状态开始，然后转换为 StateActive 或 StateClosed。
	StateNew ConnState = iota

    // StateActive represents a connection that has read 1 or more
	// bytes of a request. The Server.ConnState hook for
	// StateActive fires before the request has entered a handler
	// and doesn't fire again until the request has been
	// handled. After the request is handled, the state
	// transitions to StateClosed, StateHijacked, or StateIdle.
	// For HTTP/2, StateActive fires on the transition from zero
	// to one active request, and only transitions away once all
	// active requests are complete. That means that ConnState
	// cannot be used to do per-request work; ConnState only notes
	// the overall state of the connection.
	// StateActive 表示已读取一个或多个请求字节的连接。
    // 对于 StateActive 的 Server.ConnState 钩子
    // 会在请求进入处理程序之前触发，
    // 然后在请求被处理之前不再触发。请求处理完成后，
    // 状态转换为 StateClosed、StateHijacked 或 StateIdle。
    // 对于 HTTP/2，StateActive 在从零到一个活动请求的转换时触发，
    // 并且仅在所有活动请求完成后再次转换。
    // 这意味着 ConnState 不能用于执行每个请求的工作；
    // ConnState 只记录连接的整体状态。
	StateActive

    // StateIdle represents a connection that has finished
	// handling a request and is in the keep-alive state, waiting
	// for a new request. Connections transition from StateIdle
	// to either StateActive or StateClosed.
	// StateIdle 表示已完成请求处理并处于保持活动状态的连接，
    // 等待新的请求。
    // 连接从 StateIdle 转换为 StateActive 或 StateClosed。
	StateIdle

    // StateHijacked represents a hijacked connection.
	// This is a terminal state. It does not transition to StateClosed.
	// StateHijacked 表示被劫持的连接。
    // 这是一个终止状态。它不会转换为 StateClosed。
	StateHijacked

    // StateClosed represents a closed connection.
	// This is a terminal state. Hijacked connections do not
	// transition to StateClosed.
	// StateClosed 表示已关闭的连接。
    // 这是一个终止状态。
    // 被劫持的连接不会转换为 StateClosed。
	StateClosed
)
```

#### (ConnState) String  <- go1.3

``` go 
func (c ConnState) String() string
```

### type Cookie 

``` go 
type Cookie struct {
	Name  string
	Value string

	Path       string    // 可选
	Domain     string    // 可选
	Expires    time.Time // 可选
	RawExpires string    // 仅用于读取cookie

    // MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	// MaxAge=0 表示没有指定"Max-Age"属性。
	// MaxAge<0 表示立即删除 cookie，相当于"Max-Age: 0"
	// MaxAge>0 表示指定了以秒为单位的"Max-Age"属性
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
	Raw      string
	Unparsed []string // 未解析的属性-值对的原始文本
}
```

A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an HTTP response or the Cookie header of an HTTP request.

​	Cookie 表示 HTTP cookie，其在 HTTP 响应的 Set-Cookie 标头或 HTTP 请求的 Cookie 标头中发送。

See https://tools.ietf.org/html/rfc6265 for details.

​	详见[https://tools.ietf.org/html/rfc6265](https://tools.ietf.org/html/rfc6265)。

#### func ParseCookie <- go1.23.0

```
func ParseCookie(line string) ([]*Cookie, error)
```

ParseCookie parses a Cookie header value and returns all the cookies which were set in it. Since the same cookie name can appear multiple times the returned Values can contain more than one value for a given key.

​	ParseCookie 解析一个 Cookie 头部的值，并返回其中设置的所有 Cookie。由于同一 Cookie 名称可能会出现多次，因此返回的值可能包含相同键的多个值。

#### func ParseSetCookie <- go1.23.0

```
func ParseSetCookie(line string) (*Cookie, error)
```

ParseSetCookie parses a Set-Cookie header value and returns a cookie. It returns an error on syntax error.

​	ParseSetCookie 解析一个 Set-Cookie 头部的值并返回一个 Cookie。如果语法错误，它将返回错误。

#### (*Cookie) String 

``` go 
func (c *Cookie) String() string
```

String returns the serialization of the cookie for use in a Cookie header (if only Name and Value are set) or a Set-Cookie response header (if other fields are set). If c is nil or c.Name is invalid, the empty string is returned.

​	String 返回 cookie 的序列化字符串，以用于 Cookie 标头(如果仅设置了 Name 和 Value)或 Set-Cookie 响应标头(如果设置了其他字段)。如果 c 为 nil 或 c.Name 无效，则返回空字符串。

#### (*Cookie) Valid  <- go1.18

``` go 
func (c *Cookie) Valid() error
```

Valid reports whether the cookie is valid.

​	Valid 函数用于判断 cookie 是否有效。

### type CookieJar 

``` go 
type CookieJar interface {
    // SetCookies handles the receipt of the cookies in a reply for the
	// given URL.  It may or may not choose to save the cookies, depending
	// on the jar's policy and implementation.
	// SetCookies 处理收到的回复中的 cookie。
    // 根据 jar 的策略和实现，它可能会选择保存 cookie。
	SetCookies(u *url.URL, cookies []*Cookie)

    // Cookies returns the cookies to send in a request for the given URL.
	// It is up to the implementation to honor the standard cookie use
	// restrictions such as in RFC 6265.
	// Cookies 返回指定 URL 的请求中应发送的 cookie。
    // 实现必须遵守 RFC 6265 等标准 cookie 使用限制。
	Cookies(u *url.URL) []*Cookie
}
```

A CookieJar manages storage and use of cookies in HTTP requests.

​	CookieJar 管理 HTTP 请求中的 cookie 的存储和使用。

Implementations of CookieJar must be safe for concurrent use by multiple goroutines.

​	CookieJar 的实现必须支持多 goroutine 并发使用。

The net/http/cookiejar package provides a CookieJar implementation.

​	net/http/cookiejar 包提供了 CookieJar 的实现。

### type Dir 

``` go 
type Dir string
```

A Dir implements FileSystem using the native file system restricted to a specific directory tree.

​	Dir 类型使用本地文件系统来实现 FileSystem 接口，且仅限于特定的目录树。

While the FileSystem.Open method takes '/'-separated paths, a Dir's string value is a filename on the native file system, not a URL, so it is separated by filepath.Separator, which isn't necessarily '/'.

​	尽管 FileSystem.Open 方法接收以 / 分隔的路径，但 Dir 的 string 值是本地文件系统上的文件名，而不是 URL，因此它由 filepath.Separator 分隔，不一定是 `/`。

Note that Dir could expose sensitive files and directories. Dir will follow symlinks pointing out of the directory tree, which can be especially dangerous if serving from a directory in which users are able to create arbitrary symlinks. Dir will also allow access to files and directories starting with a period, which could expose sensitive directories like .git or sensitive files like .htpasswd. To exclude files with a leading period, remove the files/directories from the server or create a custom FileSystem implementation.

​	请注意，Dir 可能会暴露敏感文件和目录。Dir 将遵循指向目录树之外的符号链接，如果从用户可以创建任意符号链接的目录提供服务，这可能会特别危险。Dir 还将允许访问以句点开头的文件和目录，这可能会暴露像 .git 这样的敏感目录或像 `.htpasswd` 这样的敏感文件。要排除以句点开头的文件，请从服务器删除这些文件/目录，或者创建自定义 FileSystem 实现。

An empty Dir is treated as ".".

​	一个空的 Dir 被视为 "`.`"。

#### (Dir) Open 

``` go 
func (d Dir) Open(name string) (File, error)
```

Open implements FileSystem using os.Open, opening files for reading rooted and relative to the directory d.

​	Open 实现 FileSystem 接口，使用 os.Open 打开文件进行读取，根据目录 d 来确定其路径。

### type File 

``` go 
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]fs.FileInfo, error)
	Stat() (fs.FileInfo, error)
}
```

A File is returned by a FileSystem's Open method and can be served by the FileServer implementation.

​	该接口的 Open 方法返回一个 File，可由 FileServer 实现进行服务。

The methods should behave the same as those on an *os.File.

​	该接口的方法应与 `*os.File` 上的方法表现相同。

### type FileSystem 

``` go 
type FileSystem interface {
	Open(name string) (File, error)
}
```

A FileSystem implements access to a collection of named files. The elements in a file path are separated by slash ('/', U+002F) characters, regardless of host operating system convention. See the FileServer function to convert a FileSystem to a Handler.

​	FileSystem 实现对命名文件集合的访问。文件路径中的元素使用斜杠('/'，U+002F)字符分隔，而不管主机操作系统惯例如何。请参见 FileServer 函数，将 FileSystem 转换为处理程序。

This interface predates the fs.FS interface, which can be used instead: the FS adapter function converts an fs.FS to a FileSystem.

​	该接口早于 fs.FS 接口，可以使用 fs.FS 代替：FS 适配器函数将 fs.FS 转换为 FileSystem。

#### func FS  <- go1.16

``` go 
func FS(fsys fs.FS) FileSystem
```

FS converts fsys to a FileSystem implementation, for use with FileServer and NewFileTransport. The files provided by fsys must implement io.Seeker.

​	FS 将 fsys 转换为 FileSystem 实现，供 FileServer 和 NewFileTransport 使用。由 fsys 提供的文件必须实现 io.Seeker。

### type Flusher 

``` go 
type Flusher interface {
	// Flush sends any buffered data to the client.
	Flush()
}
```

The Flusher interface is implemented by ResponseWriters that allow an HTTP handler to flush buffered data to the client.

​	Flusher 接口由 ResponseWriter 实现，允许 HTTP 处理程序将缓冲的数据刷新到客户端。

The default HTTP/1.x and HTTP/2 ResponseWriter implementations support Flusher, but ResponseWriter wrappers may not. Handlers should always test for this ability at runtime.

​	默认的 HTTP/1.x 和 HTTP/2 ResponseWriter 实现支持 Flusher，但 ResponseWriter 包装器可能不支持。处理程序应始终在运行时测试此功能。

Note that even for ResponseWriters that support Flush, if the client is connected through an HTTP proxy, the buffered data may not reach the client until the response completes.

​	请注意，即使对于支持 Flush 的 ResponseWriter，如果客户端通过 HTTP 代理连接，则缓冲的数据可能要等到响应完成后才能到达客户端。

### type Handler 

``` go 
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

A Handler responds to an HTTP request.

​	Handler 响应 HTTP 请求。

ServeHTTP should write reply headers and data to the ResponseWriter and then return. Returning signals that the request is finished; it is not valid to use the ResponseWriter or read from the Request.Body after or concurrently with the completion of the ServeHTTP call.

​	ServeHTTP 应将响应标头和数据写入 ResponseWriter，然后返回。返回信号表示请求已完成；在完成 ServeHTTP 调用之后或与之同时使用 ResponseWriter 或从 Request.Body 中读取是无效的。

Depending on the HTTP client software, HTTP protocol version, and any intermediaries between the client and the Go server, it may not be possible to read from the Request.Body after writing to the ResponseWriter. Cautious handlers should read the Request.Body first, and then reply.

​	根据HTTP客户端软件、HTTP协议版本以及客户端和Go服务器之间的任何中介，可能无法在写入ResponseWriter之后从Request.Body中读取。谨慎的处理程序应该先读取Request.Body，然后再回复。

Except for reading the body, handlers should not modify the provided Request.

​	除了读取body外，处理程序不应修改提供的Request。

If ServeHTTP panics, the server (the caller of ServeHTTP) assumes that the effect of the panic was isolated to the active request. It recovers the panic, logs a stack trace to the server error log, and either closes the network connection or sends an HTTP/2 RST_STREAM, depending on the HTTP protocol. To abort a handler so the client sees an interrupted response but the server doesn't log an error, panic with the value ErrAbortHandler.

​	如果ServeHTTP恐慌，服务器(ServeHTTP的调用者)假定恐慌的影响仅限于活动请求。它会恢复panic，将堆栈跟踪记录到服务器错误日志，并关闭网络连接或发送HTTP/2 RST_STREAM，具体取决于HTTP协议。为了中止处理程序，以便客户端看到中断的响应但服务器不记录错误，请使用值为ErrAbortHandler的panic。

#### func AllowQuerySemicolons  <- go1.17

``` go 
func AllowQuerySemicolons(h Handler) Handler
```

AllowQuerySemicolons returns a handler that serves requests by converting any unescaped semicolons in the URL query to ampersands, and invoking the handler h.

​	AllowQuerySemicolons返回一个处理程序，它通过将URL查询中的任何未转义分号转换为"&"并调用处理程序h来为请求提供服务。

This restores the pre-Go 1.17 behavior of splitting query parameters on both semicolons and ampersands. (See golang.org/issue/25192). Note that this behavior doesn't match that of many proxies, and the mismatch can lead to security issues.

​	这恢复了分割查询参数的分号和"&"的Go 1.17之前的行为。(请参阅golang.org/issue/25192)。请注意，此行为与许多代理的行为不匹配，不匹配可能导致安全问题。

AllowQuerySemicolons should be invoked before Request.ParseForm is called.

​	在调用Request.ParseForm之前，应调用AllowQuerySemicolons。

#### func FileServer 

``` go 
func FileServer(root FileSystem) Handler
```

FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.

​	FileServer返回一个处理程序，它使用以root为根的文件系统的内容为HTTP请求提供服务。

As a special case, the returned file server redirects any request ending in "/index.html" to the same path, without the final "index.html".

​	作为一个特殊情况，返回的文件服务器将以"/index.html"结尾的任何请求重定向到相同的路径，而不包括最后的"index.html"。

To use the operating system's file system implementation, use http.Dir:

​	要使用操作系统的文件系统实现，请使用http.Dir：

```
http.Handle("/", http.FileServer(http.Dir("/tmp")))
```

To use an fs.FS implementation, use http.FS to convert it:

​	要使用fs.FS实现，请使用http.FS将其转换：

```
http.Handle("/", http.FileServer(http.FS(fsys)))
```

##### FileServer Example
``` go 
package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}

```

##### FileServer Example (DotFileHiding)
``` go 
package main

import (
	"io/fs"
	"log"
	"net/http"
	"strings"
)

// containsDotFile reports whether name contains a path element starting with a period.
// The name is assumed to be a delimited by forward slashes, as guaranteed
// by the http.FileSystem interface.
func containsDotFile(name string) bool {
	parts := strings.Split(name, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ".") {
			return true
		}
	}
	return false
}

// dotFileHidingFile is the http.File use in dotFileHidingFileSystem.
// It is used to wrap the Readdir method of http.File so that we can
// remove files and directories that start with a period from its output.
type dotFileHidingFile struct {
	http.File
}

// Readdir is a wrapper around the Readdir method of the embedded File
// that filters out all files that start with a period in their name.
func (f dotFileHidingFile) Readdir(n int) (fis []fs.FileInfo, err error) {
	files, err := f.File.Readdir(n)
	for _, file := range files { // Filters out the dot files
		if !strings.HasPrefix(file.Name(), ".") {
			fis = append(fis, file)
		}
	}
	return
}

// dotFileHidingFileSystem is an http.FileSystem that hides
// hidden "dot files" from being served.
type dotFileHidingFileSystem struct {
	http.FileSystem
}

// Open is a wrapper around the Open method of the embedded FileSystem
// that serves a 403 permission error when name has a file or directory
// with whose name starts with a period in its path.
func (fsys dotFileHidingFileSystem) Open(name string) (http.File, error) {
	if containsDotFile(name) { // If dot file, return 403 response
		return nil, fs.ErrPermission
	}

	file, err := fsys.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return dotFileHidingFile{file}, err
}

func main() {
	fsys := dotFileHidingFileSystem{http.Dir(".")}
	http.Handle("/", http.FileServer(fsys))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

##### FileServer Example (StripPrefix) 
``` go 
package main

import (
	"net/http"
)

func main() {
	// To serve a directory on disk (/tmp) under an alternate URL
	// path (/tmpfiles/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
}

```

#### func FileServerFS <- go1.22.0

```
func FileServerFS(root fs.FS) Handler
```

FileServerFS returns a handler that serves HTTP requests with the contents of the file system fsys. The files provided by fsys must implement [io.Seeker](https://pkg.go.dev/io#Seeker).

​	FileServerFS 返回一个处理器，用于通过文件系统 `fsys` 的内容来处理 HTTP 请求。`fsys` 提供的文件必须实现 [io.Seeker](https://pkg.go.dev/io#Seeker) 接口。

As a special case, the returned file server redirects any request ending in "/index.html" to the same path, without the final "index.html".

​	作为一种特殊情况，返回的文件服务器会将任何以 "/index.html" 结尾的请求重定向到相同路径，但去掉最后的 "index.html"。

```
http.Handle("/", http.FileServerFS(fsys))
```

#### func MaxBytesHandler  <- go1.18

``` go 
func MaxBytesHandler(h Handler, n int64) Handler
```

MaxBytesHandler returns a Handler that runs h with its ResponseWriter and Request.Body wrapped by a MaxBytesReader.

​	MaxBytesHandler返回一个处理程序，该处理程序使用其ResponseWriter和Request.Body包装了一个MaxBytesReader。

#### func NotFoundHandler 

``` go 
func NotFoundHandler() Handler
```

NotFoundHandler returns a simple request handler that replies to each request with a “404 page not found” reply.

​	NotFoundHandler返回一个简单的请求处理程序，它回复每个请求的"404页面未找到"回复。

##### NotFoundHandler Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/http"
)

func newPeopleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the people handler.")
	})
}

func main() {
	mux := http.NewServeMux()

	// Create sample handler to returns 404
	mux.Handle("/resources", http.NotFoundHandler())

	// Create sample handler that returns 200
	mux.Handle("/resources/people/", newPeopleHandler())

	log.Fatal(http.ListenAndServe(":8080", mux))
}

```

#### func RedirectHandler 

``` go 
func RedirectHandler(url string, code int) Handler
```

RedirectHandler returns a request handler that redirects each request it receives to the given url using the given status code.

​	RedirectHandler 函数返回一个请求处理程序，该程序会使用给定的状态码将收到的每个请求重定向到给定的URL。

The provided code should be in the 3xx range and is usually StatusMovedPermanently, StatusFound or StatusSeeOther.

​	提供的状态码应该在 3xx 范围内，通常为 StatusMovedPermanently、StatusFound 或 StatusSeeOther。

#### func StripPrefix 

``` go 
func StripPrefix(prefix string, h Handler) Handler
```

StripPrefix returns a handler that serves HTTP requests by removing the given prefix from the request URL's Path (and RawPath if set) and invoking the handler h. StripPrefix handles a request for a path that doesn't begin with prefix by replying with an HTTP 404 not found error. The prefix must match exactly: if the prefix in the request contains escaped characters the reply is also an HTTP 404 not found error.

​	StripPrefix 函数返回一个处理程序，它通过从请求 URL 的路径(如果设置，则为 RawPath)中删除给定前缀并调用处理程序 h 来服务 HTTP 请求。如果路径不以前缀开头，则 StripPrefix 会用 HTTP 404 未找到错误回复。前缀必须精确匹配：如果请求中的前缀包含转义字符，则回复也是 HTTP 404 未找到错误。

##### StripPrefix Example
``` go 
package main

import (
	"net/http"
)

func main() {
	// To serve a directory on disk (/tmp) under an alternate URL
	// path (/tmpfiles/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
}

```

#### func TimeoutHandler 

``` go 
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
```

TimeoutHandler returns a Handler that runs h with the given time limit.

​	TimeoutHandler 函数返回一个运行限时的处理程序。

The new Handler calls h.ServeHTTP to handle each request, but if a call runs for longer than its time limit, the handler responds with a 503 Service Unavailable error and the given message in its body. (If msg is empty, a suitable default message will be sent.) After such a timeout, writes by h to its ResponseWriter will return ErrHandlerTimeout.

​	新的处理程序调用 h.ServeHTTP 来处理每个请求，但如果一个调用的运行时间超过其时间限制，则处理程序会使用一个 HTTP 503 Service Unavailable 错误并在其主体中给出给定的消息进行回复。 (如果 msg 为空，则将发送适当的默认消息。)在此类超时之后，h 对其 ResponseWriter 的写入将返回 ErrHandlerTimeout。

TimeoutHandler supports the Pusher interface but does not support the Hijacker or Flusher interfaces.

​	TimeoutHandler 支持 Pusher 接口，但不支持 Hijacker 或 Flusher 接口。

### type HandlerFunc 

``` go 
type HandlerFunc func(ResponseWriter, *Request)
```

The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

​	HandlerFunc 类型是一个适配器，允许使用普通函数作为 HTTP 处理程序。如果 f 是具有适当签名的函数，则 HandlerFunc(f) 是调用 f 的处理程序。

#### (HandlerFunc) ServeHTTP 

``` go 
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

ServeHTTP calls f(w, r).

​	ServeHTTP 调用 f(w, r)。

### type Header 

``` go 
type Header map[string][]string
```

A Header represents the key-value pairs in an HTTP header.

​	Header 表示 HTTP 头中的键值对。

The keys should be in canonical form, as returned by CanonicalHeaderKey.

​	键应该是规范化形式，即由 CanonicalHeaderKey函数返回的形式。

#### (Header) Add 

``` go 
func (h Header) Add(key, value string)
```

Add adds the key, value pair to the header. It appends to any existing values associated with key. The key is case insensitive; it is canonicalized by CanonicalHeaderKey.

​	Add方法向 Header 中添加一个键值对，它会将值附加到与键相关联的任何现有值的末尾。键的大小写不敏感，因此会通过 CanonicalHeaderKey 进行规范化。

#### (Header) Clone  <- go1.13

``` go 
func (h Header) Clone() Header
```

Clone returns a copy of h or nil if h is nil.

​	Clone方法返回 h 的一个副本，如果 h 为 nil，则返回 nil。

#### (Header) Del 

``` go 
func (h Header) Del(key string)
```

Del deletes the values associated with key. The key is case insensitive; it is canonicalized by CanonicalHeaderKey.

​	Del方法删除与 key 关联的值。键的大小写不敏感，因此会通过 CanonicalHeaderKey函数进行规范化。

#### (Header) Get 

``` go 
func (h Header) Get(key string) string
```

Get gets the first value associated with the given key. If there are no values associated with the key, Get returns "". It is case insensitive; textproto.CanonicalMIMEHeaderKey is used to canonicalize the provided key. Get assumes that all keys are stored in canonical form. To use non-canonical keys, access the map directly.

​	Get方法获取与给定键关联的第一个值。如果没有与该键关联的值，则 Get 返回 ""。键的大小写不敏感，因此使用 textproto.CanonicalMIMEHeaderKey 来规范化提供的键。Get 假定所有键都以规范形式存储。要使用非规范键，请直接访问映射。

#### (Header) Set 

``` go 
func (h Header) Set(key, value string)
```

Set sets the header entries associated with key to the single element value. It replaces any existing values associated with key. The key is case insensitive; it is canonicalized by textproto.CanonicalMIMEHeaderKey. To use non-canonical keys, assign to the map directly.

​	Set方法将与 key 关联的 header 条目设置为单个元素值。它将替换与 key 关联的任何现有值。键的大小写不敏感，因此会通过 textproto.CanonicalMIMEHeaderKey 进行规范化。要使用非规范键，请直接分配到映射中。

#### (Header) Values  <- go1.14

``` go 
func (h Header) Values(key string) []string
```

Values returns all values associated with the given key. It is case insensitive; textproto.CanonicalMIMEHeaderKey is used to canonicalize the provided key. To use non-canonical keys, access the map directly. The returned slice is not a copy.

​	Values方法返回与给定键关联的所有值。键的大小写不敏感，因此使用 textproto.CanonicalMIMEHeaderKey 来规范化提供的键。要使用非规范键，请直接访问映射。返回的切片不是副本。

#### (Header) Write 

``` go 
func (h Header) Write(w io.Writer) error
```

Write writes a header in wire format.

​	Write 方法以 wire format 写入 Header。

#### (Header) WriteSubset 

``` go 
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```

WriteSubset writes a header in wire format. If exclude is not nil, keys where exclude[key] == true are not written. Keys are not canonicalized before checking the exclude map.

​	WriteSubset 方法以 wire format 写入 Header。如果 exclude 不为 nil，则不会写入其中 exclude[key] == true 的键。写入之前不会对键进行规范化。

### type Hijacker 

``` go 
type Hijacker interface {
    // Hijack lets the caller take over the connection.
	// After a call to Hijack the HTTP server library
	// will not do anything else with the connection.
	//
	// It becomes the caller's responsibility to manage
	// and close the connection.
	//
	// The returned net.Conn may have read or write deadlines
	// already set, depending on the configuration of the
	// Server. It is the caller's responsibility to set
	// or clear those deadlines as needed.
	//
	// The returned bufio.Reader may contain unprocessed buffered
	// data from the client.
	//
	// After a call to Hijack, the original Request.Body must not
	// be used. The original Request's Context remains valid and
	// is not canceled until the Request's ServeHTTP method
	// returns.
	// Hijack允许调用方接管连接。
	// 调用Hijack后，HTTP服务器库将不会再对连接进行任何操作。
	//
	// 管理和关闭连接将成为调用方的责任。
	//
	// 返回的net.Conn可能已经设置了读取或写入期限，
    // 具体取决于服务器的配置。
    // 调用方有责任根据需要设置或清除这些期限。
    //
    // 返回的bufio.Reader可能包含来自客户端的未处理缓冲数据。
    //
    // 在调用Hijack后，原始Request.Body不能再使用。
    // 原始请求的上下文保持有效，直到请求的ServeHTTP方法返回为止。
	Hijack() (net.Conn, *bufio.ReadWriter, error)
}
```

The Hijacker interface is implemented by ResponseWriters that allow an HTTP handler to take over the connection.

​	Hijacker接口由允许HTTP处理程序接管连接的ResponseWriters实现。

The default ResponseWriter for HTTP/1.x connections supports Hijacker, but HTTP/2 connections intentionally do not. ResponseWriter wrappers may also not support Hijacker. Handlers should always test for this ability at runtime.

​	默认的HTTP/1.x连接的ResponseWriter支持Hijacker，但是HTTP/2连接有意不支持它。 ResponseWriter包装器也可能不支持Hijacker。处理程序应始终在运行时测试此能力。

##### Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hijack", func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Don't forget to close the connection:
		defer conn.Close()
		bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
		bufrw.Flush()
		s, err := bufrw.ReadString('\n')
		if err != nil {
			log.Printf("error reading string: %v", err)
			return
		}
		fmt.Fprintf(bufrw, "You said: %q\nBye.\n", s)
		bufrw.Flush()
	})
}

```

### type MaxBytesError  <- go1.19

``` go 
type MaxBytesError struct {
	Limit int64
}
```

MaxBytesError is returned by MaxBytesReader when its read limit is exceeded.

​	MaxBytesReader读取的字节数超过其读取限制时返回MaxBytesError。

#### (*MaxBytesError) Error  <- go1.19

``` go 
func (e *MaxBytesError) Error() string
```

### type ProtocolError <- DEPRECATED

```go
type ProtocolError struct {
	ErrorString string
}
```

ProtocolError represents an HTTP protocol error.

​	ProtocolError 表示一个 HTTP 协议错误。

Deprecated: Not all errors in the http package related to protocol errors are of type ProtocolError.

​	已弃用：并非所有与协议错误相关的 `http` 包中的错误都属于 `ProtocolError` 类型。

#### (*ProtocolError) Error

```
func (pe *ProtocolError) Error() string
```

####  (*ProtocolError) Is <-go1.21.0

```go
func (pe *ProtocolError) Is(err error) bool
```

Is lets http.ErrNotSupported match errors.ErrUnsupported.

​	`Is` 允许 `http.ErrNotSupported` 与 `errors.ErrUnsupported` 匹配。

### type PushOptions  <- go1.8

``` go 
type PushOptions struct {
    // Method specifies the HTTP method for the promised request.
	// If set, it must be "GET" or "HEAD". Empty means "GET".
	// Method指定承诺请求的HTTP方法。
    // 如果设置，则必须是"GET"或"HEAD"。空表示"GET"。
	Method string

    // Header specifies additional promised request headers. This cannot
	// include HTTP/2 pseudo header fields like ":path" and ":scheme",
	// which will be added automatically.
	// Header指定额外的承诺请求标头。
    // 这不能包括HTTP/2伪标头字段，如"：path"和"：scheme"，
    // 它们将自动添加。
	Header Header
}
```

PushOptions describes options for Pusher.Push.

### type Pusher  <- go1.8

``` go 
type Pusher interface {
    // Push initiates an HTTP/2 server push. This constructs a synthetic
	// request using the given target and options, serializes that request
	// into a PUSH_PROMISE frame, then dispatches that request using the
	// server's request handler. If opts is nil, default options are used.
	//
	// The target must either be an absolute path (like "/path") or an absolute
	// URL that contains a valid host and the same scheme as the parent request.
	// If the target is a path, it will inherit the scheme and host of the
	// parent request.
	//
	// The HTTP/2 spec disallows recursive pushes and cross-authority pushes.
	// Push may or may not detect these invalid pushes; however, invalid
	// pushes will be detected and canceled by conforming clients.
	//
	// Handlers that wish to push URL X should call Push before sending any
	// data that may trigger a request for URL X. This avoids a race where the
	// client issues requests for X before receiving the PUSH_PROMISE for X.
	//
	// Push will run in a separate goroutine making the order of arrival
	// non-deterministic. Any required synchronization needs to be implemented
	// by the caller.
	//
	// Push returns ErrNotSupported if the client has disabled push or if push
	// is not supported on the underlying connection.
	// Push启动HTTP/2服务器推送。
    // 这将使用给定的目标和选项构造合成请求，
    // 将该请求序列化为PUSH_PROMISE帧，
    // 然后使用服务器的请求处理程序分派该请求。
    // 如果opts为nil，则使用默认选项。
	//
	// 目标必须是绝对路径(如"/path")或
    // 包含有效主机和与父请求相同的方案的绝对URL。
    // 如果目标是路径，则会继承父请求的方案和主机。
	//
	// HTTP/2规范禁止递归推送和跨权威推送。
    // Push可以或可以不检测这些无效推送；
    // 但是，符合规范的客户端将检测并取消无效的推送。
	//
	// 希望推送URL X的处理程序应在发送可能触发
    // 对URL X的请求的任何数据之前调用Push。
    // 这避免了客户端在收到X的PUSH_PROMISE之前就发出X的请求的竞争条件。
	//
	// Push将在单独的goroutine中运行，使到达顺序是不确定的。
    // 任何必需的同步需要由调用者实现。
	//
	// 如果客户端禁用推送或底层连接不支持推送，
    // 则Push返回ErrNotSupported。
	Push(target string, opts *PushOptions) error
}
```

Pusher is the interface implemented by ResponseWriters that support HTTP/2 server push. For more background, see https://tools.ietf.org/html/rfc7540#section-8.2.

​	Pusher是ResponseWriters实现的接口，用于支持HTTP/2服务器推送。有关更多背景信息，请参见https://tools.ietf.org/html/rfc7540#section-8.2。

### type Request 

``` go 
type Request struct {
    // Method specifies the HTTP method (GET, POST, PUT, etc.).
	// For client requests, an empty string means GET.
	//
	// Go's HTTP client does not support sending a request with
	// the CONNECT method. See the documentation on Transport for
	// details.
	// Method 指定HTTP方法(GET，POST，PUT等)。
	// 对于客户端请求，空字符串表示GET。
	//
	// Go的HTTP客户端不支持使用CONNECT方法发送请求。
    // 有关详细信息，请参阅传输的文档。
	Method string

    // URL specifies either the URI being requested (for server
	// requests) or the URL to access (for client requests).
	//
	// For server requests, the URL is parsed from the URI
	// supplied on the Request-Line as stored in RequestURI.  For
	// most requests, fields other than Path and RawQuery will be
	// empty. (See RFC 7230, Section 5.3)
	//
	// For client requests, the URL's Host specifies the server to
	// connect to, while the Request's Host field optionally
	// specifies the Host header value to send in the HTTP
	// request.
	// URL指定要请求的URI(用于服务器请求)
    // 或要访问的URL(用于客户端请求)。
	//
	// 对于服务器请求，
    // URL从存储在RequestURI中的Request-Line中提供的URI中解析。
    // 对于大多数请求，除Path和RawQuery之外的字段都将为空。 
    // (请参见RFC 7230，第5.3节)
	//
	// 对于客户端请求，URL的Host指定要连接的服务器，
    // 而Request的Host字段可选地指定要在HTTP请求中发送的Host标头值。
	URL *url.URL

    // The protocol version for incoming server requests.
	//
	// For client requests, these fields are ignored. The HTTP
	// client code always uses either HTTP/1.1 or HTTP/2.
	// See the docs on Transport for details.
	// 传入服务器请求的协议版本。
	//
	// 对于客户端请求，将忽略这些字段。
    // HTTP客户端代码始终使用HTTP/1.1或HTTP/2。
    // 有关详细信息，请参阅传输文档。
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0

    // Header contains the request header fields either received
	// by the server or to be sent by the client.
	//
	// If a server received a request with header lines,
	//
	//	Host: example.com
	//	accept-encoding: gzip, deflate
	//	Accept-Language: en-us
	//	fOO: Bar
	//	foo: two
	//
	// then
	//
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
	//
	// For incoming requests, the Host header is promoted to the
	// Request.Host field and removed from the Header map.
	//
	// HTTP defines that header names are case-insensitive. The
	// request parser implements this by using CanonicalHeaderKey,
	// making the first character and any characters following a
	// hyphen uppercase and the rest lowercase.
	//
	// For client requests, certain headers such as Content-Length
	// and Connection are automatically written when needed and
	// values in Header may be ignored. See the documentation
	// for the Request.Write method.
	// Header 包含服务器收到的请求头字段或客户端将要发送的请求头字段。
    //
    // 如果服务器收到了以下请求头行：
    //
    // Host: example.com
    // accept-encoding: gzip, deflate
    // Accept-Language: en-us
    // fOO: Bar
    // foo: two
    //
    // 那么
    //
    // Header = map[string][]string{
    // "Accept-Encoding": {"gzip, deflate"},
    // "Accept-Language": {"en-us"},
    // "Foo": {"Bar", "two"},
    // }
    //
    // 对于传入的请求，
    // Host 头部将提升到 Request.Host 字段并从 Header 映射中删除。
    //
    // HTTP 定义了头部名称不区分大小写。
    // 请求解析器通过使用 CanonicalHeaderKey 实现这一点，
    // 将连字符后的第一个字符和任何后续字符转换为大写字母，
    // 其余字母转换为小写字母。
    //
    // 对于客户端请求，当需要时会自动写入某些头部字段，
    // 例如 Content-Length 和 Connection，
    // 并且 Header 中的值可能会被忽略。
    // 请参阅 Request.Write 方法的文档。
	Header Header

    // Body is the request's body.
	//
	// For client requests, a nil body means the request has no
	// body, such as a GET request. The HTTP Client's Transport
	// is responsible for calling the Close method.
	//
	// For server requests, the Request Body is always non-nil
	// but will return EOF immediately when no body is present.
	// The Server will close the request body. The ServeHTTP
	// Handler does not need to.
	//
	// Body must allow Read to be called concurrently with Close.
	// In particular, calling Close should unblock a Read waiting
	// for input.
	// Body 是请求的主体。
    //
    // 对于客户端请求，nil 主体表示请求没有主体，例如 GET 请求。
    // HTTP 客户端的 Transport 负责调用 Close 方法。
    //
    // 对于服务器请求，请求 Body 总是非空的，
    // 但在没有主体的情况下将立即返回 EOF。
    // 服务器将关闭请求主体。ServeHTTP 处理程序不需要。
    //
    // Body 必须允许同时调用 Read 和 Close。
    // 特别是，调用 Close 应该取消阻塞等待输入的 Read。
	Body io.ReadCloser

    // GetBody defines an optional func to return a new copy of
	// Body. It is used for client requests when a redirect requires
	// reading the body more than once. Use of GetBody still
	// requires setting Body.
	//
	// For server requests, it is unused.
	// GetBody 定义了一个可选的函数，
    // 用于返回 Body 的新副本。
    // 当重定向需要多次读取主体时，它用于客户端请求。
    // 仍然需要设置 Body 才能使用 GetBody。
	//
	// 对于服务器请求，它没有使用。
	GetBody func() (io.ReadCloser, error)

    // ContentLength records the length of the associated content.
	// The value -1 indicates that the length is unknown.
	// Values >= 0 indicate that the given number of bytes may
	// be read from Body.
	//
	// For client requests, a value of 0 with a non-nil Body is
	// also treated as unknown.
	// ContentLength 记录相关内容的长度。
    // 值-1表示长度未知。值>= 0表示可以从Body读取给定的字节数。
	//
	// 对于客户端请求，具有非零Body的值为0也被视为未知。
	ContentLength int64

    // TransferEncoding lists the transfer encodings from outermost to
	// innermost. An empty list denotes the "identity" encoding.
	// TransferEncoding can usually be ignored; chunked encoding is
	// automatically added and removed as necessary when sending and
	// receiving requests.
	// TransferEncoding列出了从最外层到最内层的传输编码。
    // 空列表表示"identity"编码。
    // 在发送和接收请求时，可以通常忽略TransferEncoding。
    // 当需要时，块编码会自动添加和删除。
	TransferEncoding []string

    // Close indicates whether to close the connection after
	// replying to this request (for servers) or after sending this
	// request and reading its response (for clients).
	//
	// For server requests, the HTTP server handles this automatically
	// and this field is not needed by Handlers.
	//
	// For client requests, setting this field prevents re-use of
	// TCP connections between requests to the same hosts, as if
	// Transport.DisableKeepAlives were set.
	// Close指示是否在回复此请求后(对于服务器)
    // 或在发送此请求并读取其响应后(对于客户端)关闭连接。
	//
	// 对于服务器请求，HTTP服务器会自动处理此操作，处理程序不需要此字段。
	//
	// 对于客户端请求，
    // 设置此字段会阻止在向相同主机的请求之间重新使用TCP连接，
    // 就像Transport.DisableKeepAlives已设置一样。
	Close bool

    // For server requests, Host specifies the host on which the
	// URL is sought. For HTTP/1 (per RFC 7230, section 5.4), this
	// is either the value of the "Host" header or the host name
	// given in the URL itself. For HTTP/2, it is the value of the
	// ":authority" pseudo-header field.
	// It may be of the form "host:port". For international domain
	// names, Host may be in Punycode or Unicode form. Use
	// golang.org/x/net/idna to convert it to either format if
	// needed.
	// To prevent DNS rebinding attacks, server Handlers should
	// validate that the Host header has a value for which the
	// Handler considers itself authoritative. The included
	// ServeMux supports patterns registered to particular host
	// names and thus protects its registered Handlers.
	//
	// For client requests, Host optionally overrides the Host
	// header to send. If empty, the Request.Write method uses
	// the value of URL.Host. Host may contain an international
	// domain name.
	// 对于服务器请求，Host指定寻找URL的主机。
    // 对于HTTP / 1(根据RFC 7230，第5.4节)，
    // 这是"Host"标头的值或URL本身中给定的主机名。
    // 对于HTTP / 2，它是"：authority"伪标头字段的值。
	// 它可以采用"host：port"的形式。
    // 对于国际域名，Host可以为Punycode或Unicode形式。
    // 如果需要，可以使用golang.org/x/net/idna将其转换为任一格式。
	// 为防止DNS重绑定攻击，
    // 服务器处理程序应验证Host标头具有其认为自己具有权威的值。
    // 包括ServeMux支持为特定主机名注册的模式，因此保护其注册的处理程序。
	//
	// 对于客户端请求，Host可选择覆盖要发送的Host标头。
    // 如果为空，则Request.Write方法使用URL.Host的值。
    // 主机可能包含国际域名。
	Host string

    // Form contains the parsed form data, including both the URL
	// field's query parameters and the PATCH, POST, or PUT form data.
	// This field is only available after ParseForm is called.
	// The HTTP client ignores Form and uses Body instead.
	// Form 包含已解析的表单数据，
    // 包括 URL 字段的查询参数以及 PATCH、POST 或 PUT 表单数据。
	// 只有在调用 ParseForm 之后才能使用此字段。
    // HTTP 客户端会忽略 Form 并使用 Body。
	Form url.Values

    // PostForm contains the parsed form data from PATCH, POST
	// or PUT body parameters.
	//
	// This field is only available after ParseForm is called.
	// The HTTP client ignores PostForm and uses Body instead.
	// PostForm 包含来自 PATCH、POST 或 PUT 请求体参数的解析表单数据。
	// 只有在调用 ParseForm 之后才能使用此字段。
    // HTTP 客户端会忽略 PostForm 并使用 Body。
	PostForm url.Values

    // MultipartForm is the parsed multipart form, including file uploads.
	// This field is only available after ParseMultipartForm is called.
	// The HTTP client ignores MultipartForm and uses Body instead.
	// MultipartForm 是已解析的多部分表单，包括文件上传。
	// 只有在调用 ParseMultipartForm 之后才能使用此字段。
    // HTTP 客户端会忽略 MultipartForm 并使用 Body。
	MultipartForm *multipart.Form

    // Trailer specifies additional headers that are sent after the request
	// body.
	//
	// For server requests, the Trailer map initially contains only the
	// trailer keys, with nil values. (The client declares which trailers it
	// will later send.)  While the handler is reading from Body, it must
	// not reference Trailer. After reading from Body returns EOF, Trailer
	// can be read again and will contain non-nil values, if they were sent
	// by the client.
	//
	// For client requests, Trailer must be initialized to a map containing
	// the trailer keys to later send. The values may be nil or their final
	// values. The ContentLength must be 0 or -1, to send a chunked request.
	// After the HTTP request is sent the map values can be updated while
	// the request body is read. Once the body returns EOF, the caller must
	// not mutate Trailer.
	//
	// Few HTTP clients, servers, or proxies support HTTP trailers.
	// Trailer 指定在请求体之后发送的其他标头。
	// 对于服务器请求，Trailer 映射最初仅包含 trailer 键和 nil 值。
    // (客户端声明将稍后发送哪些 trailers。)
	// 当处理程序从 Body 读取时，它不能引用 Trailer。
    // 从 Body 返回 EOF 后，Trailer 可以再次读取，
    // 并且如果客户端发送了这些值，则将包含非 nil 的值。
	// 对于客户端请求，
    // Trailer 必须初始化为一个包含稍后发送的 trailer 键的映射。
    // 值可以为 nil 或它们的最终值。
	// ContentLength 必须为 0 或 -1，以发送分块请求。
    // 在发送 HTTP 请求后，可以在读取请求体时更新映射值。
	// 请求体返回 EOF 后，调用者不能改变 Trailer。
    // 很少有 HTTP 客户端、服务器或代理支持 HTTP trailers。
	Trailer Header

    // RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. This field is not filled in by ReadRequest and
	// has no defined format. The HTTP server in this package
	// sets RemoteAddr to an "IP:port" address before invoking a
	// handler.
	// This field is ignored by the HTTP client.
	// RemoteAddr 允许 HTTP 服务器和其他软件记录发送请求的网络地址，
    // 通常用于记录日志。
	// ReadRequest 不会填充此字段，也没有定义的格式。
    // 该包中的 HTTP 服务器在调用处理程序之前将 RemoteAddr 
    // 设置为 "IP:port" 地址。
	// HTTP 客户端会忽略此字段。
	RemoteAddr string

    // RequestURI is the unmodified request-target of the
	// Request-Line (RFC 7230, Section 3.1.1) as sent by the client
	// to a server. Usually the URL field should be used instead.
	// It is an error to set this field in an HTTP client request.
	// RequestURI 是客户端发送到服务器的请求行
    //(RFC 7230，第 3.1.1 节)的未修改请求目标。
	// 通常应该使用 URL 字段而不是此字段。
    // 在 HTTP 客户端请求中设置此字段是错误的。
	RequestURI string

    // TLS allows HTTP servers and other software to record
	// information about the TLS connection on which the request
	// was received. This field is not filled in by ReadRequest.
	// The HTTP server in this package sets the field for
	// TLS-enabled connections before invoking a handler;
	// otherwise it leaves the field nil.
	// This field is ignored by the HTTP client.
	// TLS 允许 HTTP 服务器和其他软件记录接收请求的 TLS 连接的信息。
    // 该字段不由 ReadRequest 填充。
	// 在调用处理程序之前，
    // 此包中的 HTTP 服务器在启用 TLS 的连接上设置字段；
    // 否则它将保留字段为空。
	// HTTP 客户端忽略此字段。
	TLS *tls.ConnectionState

    // Cancel is an optional channel whose closure indicates that the client
	// request should be regarded as canceled. Not all implementations of
	// RoundTripper may support Cancel.
	//
	// For server requests, this field is not applicable.
	//
	// Deprecated: Set the Request's context with NewRequestWithContext
	// instead. If a Request's Cancel field and context are both
	// set, it is undefined whether Cancel is respected.
	// Cancel 是一个可选通道，其关闭表示应将客户端请求视为已取消。
    // 并非所有 RoundTripper 的实现都支持 Cancel。
    //
    // 对于服务器请求，此字段不适用。
    //
    // 已弃用：应使用 NewRequestWithContext 
    // 将请求的上下文设置为上下文，而不是设置请求的 Cancel 字段。
    // 如果请求的 Cancel 字段和上下文都已设置，则未定义是否尊重 Cancel。
	Cancel <-chan struct{}

    // Response is the redirect response which caused this request
	// to be created. This field is only populated during client
	// redirects.
	// Response 是重定向响应，导致创建此请求。
    // 此字段仅在客户端重定向期间填充。
	Response *Response
	// 包含过滤或未导出的字段
}
```

A Request represents an HTTP request received by a server or to be sent by a client.

​	一个Request代表一个由服务器接收或由客户端发送的HTTP请求。

The field semantics differ slightly between client and server usage. In addition to the notes on the fields below, see the documentation for Request.Write and RoundTripper.

​	客户端和服务器使用中字段语义略有不同。除了下面字段的注释外，请参阅Request.Write和RoundTripper的文档。

#### func NewRequest 

``` go 
func NewRequest(method, url string, body io.Reader) (*Request, error)
```

NewRequest wraps NewRequestWithContext using context.Background.

​	NewRequest使用context.Background包装NewRequestWithContext。

#### func NewRequestWithContext  <- go1.13

``` go 
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error)
```

NewRequestWithContext returns a new Request given a method, URL, and optional body.

​	NewRequestWithContext给定方法、URL和可选的body，返回一个新的Request。

If the provided body is also an io.Closer, the returned Request.Body is set to body and will be closed by the Client methods Do, Post, and PostForm, and Transport.RoundTrip.

NewRequestWithContext returns a Request suitable for use with Client.Do or Transport.RoundTrip. To create a request for use with testing a Server Handler, either use the NewRequest function in the net/http/httptest package, use ReadRequest, or manually update the Request fields. For an outgoing client request, the context controls the entire lifetime of a request and its response: obtaining a connection, sending the request, and reading the response headers and body. See the Request type's documentation for the difference between inbound and outbound request fields.

​	NewRequestWithContext返回一个适用于使用Client.Do或Transport.RoundTrip的Request。要为测试服务器处理程序创建一个请求，请使用net/http/httptest包中的NewRequest函数、使用ReadRequest或手动更新Request字段。对于传出的客户端请求，上下文控制请求及其响应的整个生命周期：获取连接、发送请求以及读取响应标头和正文。请参阅Request类型的文档，了解入站和出站请求字段之间的差异。

If body is of type *bytes.Buffer, *bytes.Reader, or *strings.Reader, the returned request's ContentLength is set to its exact value (instead of -1), GetBody is populated (so 307 and 308 redirects can replay the body), and Body is set to NoBody if the ContentLength is 0.

​	如果body的类型为*bytes.Buffer、*bytes.Reader或*strings.Reader，则返回的请求的ContentLength将设置为其确切值(而不是-1)，GetBody将被填充(因此307和308重定向可以重放body)，如果ContentLength为0，则Body将设置为NoBody。

#### func ReadRequest 

``` go 
func ReadRequest(b *bufio.Reader) (*Request, error)
```

ReadRequest reads and parses an incoming request from b.

​	ReadRequest从b中读取并解析一个传入的请求。

ReadRequest is a low-level function and should only be used for specialized applications; most code should use the Server to read requests and handle them via the Handler interface. ReadRequest only supports HTTP/1.x requests. For HTTP/2, use golang.org/x/net/http2.

​	ReadRequest是一个低级函数，只应用于专用应用程序；大多数代码应该使用Server来读取请求并通过Handler接口处理它们。ReadRequest仅支持HTTP/1.x请求。对于HTTP/2，请使用golang.org/x/net/http2。

#### (*Request) AddCookie 

``` go 
func (r *Request) AddCookie(c *Cookie)
```

AddCookie adds a cookie to the request. Per [RFC 6265 section 5.4](https://rfc-editor.org/rfc/rfc6265.html#section-5.4), AddCookie does not attach more than one Cookie header field. That means all cookies, if any, are written into the same line, separated by semicolon. AddCookie only sanitizes c's name and value, and does not sanitize a Cookie header already present in the request.

​	AddCookie向请求中添加一个cookie。根据[RFC 6265第5.4节](https://rfc-editor.org/rfc/rfc6265.html#section-5.4)，AddCookie不会添加超过一个Cookie头字段，这意味着所有cookie(如果有)都将写入同一行，由分号分隔。AddCookie仅对c的名称和值进行清理，不对已经存在于请求中的Cookie头进行清理。

#### (*Request) BasicAuth  <- go1.4

``` go 
func (r *Request) BasicAuth() (username, password string, ok bool)
```

BasicAuth returns the username and password provided in the request's Authorization header, if the request uses HTTP Basic Authentication. See [RFC 2617, Section 2](https://rfc-editor.org/rfc/rfc2617.html#section-2).

​	BasicAuth返回请求头中提供的用户名和密码，如果请求使用HTTP基本身份验证。请参阅[RFC 2617第2节](https://rfc-editor.org/rfc/rfc2617.html#section-2)。

#### (*Request) Clone  <- go1.13

``` go 
func (r *Request) Clone(ctx context.Context) *Request
```

Clone returns a deep copy of r with its context changed to ctx. The provided ctx must be non-nil.

​	Clone方法返回r的深层副本，其上下文已更改为ctx。提供的ctx必须非nil。

For an outgoing client request, the context controls the entire lifetime of a request and its response: obtaining a connection, sending the request, and reading the response headers and body.

​	对于出站客户端请求，上下文控制请求及其响应的整个生命周期：获取连接，发送请求和读取响应标头和主体。

#### (*Request) Context  <- go1.7

``` go 
func (r *Request) Context() context.Context
```

Context returns the request's context. To change the context, use Clone or WithContext.

​	Context方法返回请求的上下文。要更改上下文，请使用Clone方法或WithContext方法。

The returned context is always non-nil; it defaults to the background context.

​	返回的上下文始终非零；默认为后台上下文。

For outgoing client requests, the context controls cancellation.

​	对于输出客户端请求，上下文控制取消。

For incoming server requests, the context is canceled when the client's connection closes, the request is canceled (with HTTP/2), or when the ServeHTTP method returns.

​	对于传入的服务器请求，当客户端连接关闭，请求被取消(使用HTTP/2)或ServeHTTP方法返回时，上下文被取消。

#### (*Request) Cookie 

``` go 
func (r *Request) Cookie(name string) (*Cookie, error)
```

Cookie returns the named cookie provided in the request or ErrNoCookie if not found. If multiple cookies match the given name, only one cookie will be returned.

​	Cookie方法返回请求中提供的指定cookie或ErrNoCookie(如果没有找到)。如果有多个cookie与给定名称匹配，则仅返回一个cookie。

#### (*Request) Cookies 

``` go 
func (r *Request) Cookies() []*Cookie
```

Cookies parses and returns the HTTP cookies sent with the request.

​	Cookies方法解析并返回发送请求的HTTP cookie。

#### (*Request) FormFile 

``` go 
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
```

FormFile returns the first file for the provided form key. FormFile calls ParseMultipartForm and ParseForm if necessary.

​	FormFile方法返回所提供的表单键的第一个文件。如果需要，FormFile方法调用ParseMultipartForm方法和ParseForm方法。

#### (*Request) FormValue 

``` go 
func (r *Request) FormValue(key string) string
```

FormValue returns the first value for the named component of the query. POST and PUT body parameters take precedence over URL query string values. FormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions. If key is not present, FormValue returns the empty string. To access multiple values of the same key, call ParseForm and then inspect Request.Form directly.

​	FormValue方法返回查询的指定组件的第一个值。POST和PUT请求体参数优先于URL查询字符串值。FormValue如果需要，调用ParseMultipartForm方法和ParseForm方法，并忽略这些函数返回的任何错误。如果键不存在，则FormValue方法返回空字符串。要访问同一键的多个值，请调用ParseForm，然后直接检查Request.Form方法。

#### (*Request) MultipartReader 

``` go 
func (r *Request) MultipartReader() (*multipart.Reader, error)
```

MultipartReader returns a MIME multipart reader if this is a multipart/form-data or a multipart/mixed POST request, else returns nil and an error. Use this function instead of ParseMultipartForm to process the request body as a stream.

​	MultipartReader方法如果这是multipart/form-data或multipart/mixed POST请求，则返回MIME多部分读取器；否则返回nil和错误。使用此方法而不是ParseMultipartForm方法处理请求正文作为流。

#### (*Request) ParseForm 

``` go 
func (r *Request) ParseForm() error
```

ParseForm populates r.Form and r.PostForm.

​	ParseForm方法解析请求中的表单，并更新`r.Form`和`r.PostForm`。

For all requests, ParseForm parses the raw query from the URL and updates r.Form.

​	对于所有请求，ParseForm方法解析URL中的原始查询，并更新`r.Form`。

For POST, PUT, and PATCH requests, it also reads the request body, parses it as a form and puts the results into both r.PostForm and r.Form. Request body parameters take precedence over URL query string values in r.Form.

​	对于POST、PUT和PATCH请求，该方法还读取请求体，将其解析为表单，并将结果放入`r.PostForm`和`r.Form`中。请求体参数优先于URL查询字符串值在`r.Form`中。

If the request Body's size has not already been limited by MaxBytesReader, the size is capped at 10MB.

​	如果请求体大小还没有被`MaxBytesReader`限制，则最大为10MB。

For other HTTP methods, or when the Content-Type is not application/x-www-form-urlencoded, the request Body is not read, and r.PostForm is initialized to a non-nil, empty value.

​	对于其他HTTP方法或`Content-Type`不为`application/x-www-form-urlencoded`的情况，不读取请求体，`r.PostForm`初始化为非零但为空的值。

ParseMultipartForm calls ParseForm automatically. ParseForm is idempotent.

​	ParseMultipartForm方法会自动调用ParseForm方法。ParseForm方法是幂等的。

#### (*Request) ParseMultipartForm 

``` go 
func (r *Request) ParseMultipartForm(maxMemory int64) error
```

ParseMultipartForm parses a request body as multipart/form-data. The whole request body is parsed and up to a total of maxMemory bytes of its file parts are stored in memory, with the remainder stored on disk in temporary files. ParseMultipartForm calls ParseForm if necessary. If ParseForm returns an error, ParseMultipartForm returns it but also continues parsing the request body. After one call to ParseMultipartForm, subsequent calls have no effect.

​	ParseMultipartForm方法将请求体解析为multipart/form-data。整个请求体被解析，并且它的文件部分的最多maxMemory字节数在内存中存储，其余部分在临时文件中存储。ParseMultipartForm方法在必要时调用ParseForm方法。如果ParseForm方法返回错误，则ParseMultipartForm方法返回该错误，但也会继续解析请求体。调用一次ParseMultipartForm方法之后，随后的调用没有任何效果。

#### (*Request) PathValue <- go1.22.0

```
func (r *Request) PathValue(name string) string
```

PathValue returns the value for the named path wildcard in the [ServeMux](https://pkg.go.dev/net/http@go1.23.0#ServeMux) pattern that matched the request. It returns the empty string if the request was not matched against a pattern or there is no such wildcard in the pattern.

​	`PathValue` 返回在匹配请求的 [ServeMux](https://pkg.go.dev/net/http@go1.23.0#ServeMux) 模式中指定路径通配符的值。如果请求没有匹配到模式或模式中没有这样的通配符，则返回空字符串。

#### (*Request) PostFormValue  <- go1.1

``` go 
func (r *Request) PostFormValue(key string) string
```

PostFormValue returns the first value for the named component of the POST, PATCH, or PUT request body. URL query parameters are ignored. PostFormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions. If key is not present, PostFormValue returns the empty string.

​	PostFormValue方法返回POST、PATCH或PUT请求体中命名组件的第一个值。忽略URL查询参数。PostFormValue方法在必要时调用ParseMultipartForm和ParseForm方法，并忽略这些函数返回的任何错误。如果key不存在，则PostFormValue方法返回空字符串。

#### (*Request) ProtoAtLeast 

``` go 
func (r *Request) ProtoAtLeast(major, minor int) bool
```

ProtoAtLeast reports whether the HTTP protocol used in the request is at least major.minor.

​	ProtoAtLeast方法报告请求中使用的HTTP协议是否至少为major.minor。

#### (*Request) Referer 

``` go 
func (r *Request) Referer() string
```

Referer returns the referring URL, if sent in the request.

​	Referer方法返回引用的URL(如果在请求中发送了)。

Referer is misspelled as in the request itself, a mistake from the earliest days of HTTP. This value can also be fetched from the Header map as Header["Referer"]; the benefit of making it available as a method is that the compiler can diagnose programs that use the alternate (correct English) spelling req.Referrer() but cannot diagnose programs that use Header["Referrer"].

​	Referer方法在请求本身中就被错误地拼写为Referer，这是HTTP早期的一个错误。这个值也可以从Header映射中获取，如`Header["Referer"]`；将其作为方法可用的好处是编译器可以诊断使用替代(正确的英文)拼写`req.Referrer()`的程序，但无法诊断使用`Header["Referrer"]`的程序。

#### (*Request) SetBasicAuth 

``` go 
func (r *Request) SetBasicAuth(username, password string)
```

SetBasicAuth sets the request's Authorization header to use HTTP Basic Authentication with the provided username and password.

​	SetBasicAuth方法将请求的Authorization头设置为使用HTTP基本身份验证，提供用户名和密码。

With HTTP Basic Authentication the provided username and password are not encrypted. It should generally only be used in an HTTPS request.

​	使用HTTP基本身份验证时，提供的用户名和密码未加密。通常只应在HTTPS请求中使用它。

The username may not contain a colon. Some protocols may impose additional requirements on pre-escaping the username and password. For instance, when used with OAuth2, both arguments must be URL encoded first with url.QueryEscape.

​	用户名不能包含冒号。某些协议可能会对预转义用户名和密码有额外的要求。例如，当与OAuth2一起使用时，必须首先使用url.QueryEscape对两个参数进行URL编码。

#### (*Request) SetPathValue <- go1.22.0

```
func (r *Request) SetPathValue(name, value string)
```

SetPathValue sets name to value, so that subsequent calls to r.PathValue(name) return value.

​	`SetPathValue` 设置名称为 `name` 的通配符值为 `value`，这样后续对 `r.PathValue(name)` 的调用将返回 `value`。

#### (*Request) UserAgent 

``` go 
func (r *Request) UserAgent() string
```

UserAgent returns the client's User-Agent, if sent in the request.

​	UserAgent方法返回客户端的User-Agent，如果在请求中发送。

#### (*Request) WithContext  <- go1.7

``` go 
func (r *Request) WithContext(ctx context.Context) *Request
```

WithContext returns a shallow copy of r with its context changed to ctx. The provided ctx must be non-nil.

​	WithContext方法返回r的浅层副本，并将其上下文更改为ctx。提供的ctx必须非nil。

For outgoing client request, the context controls the entire lifetime of a request and its response: obtaining a connection, sending the request, and reading the response headers and body.

​	对于出站客户端请求，上下文控制请求及其响应的整个生命周期：获取连接、发送请求和读取响应头和主体。

To create a new request with a context, use NewRequestWithContext. To make a deep copy of a request with a new context, use Request.Clone.

​	要使用上下文创建新请求，请使用NewRequestWithContext函数。要使用新上下文对请求进行深层复制，请使用Request.Clone方法。

#### (*Request) Write 

``` go 
func (r *Request) Write(w io.Writer) error
```

Write writes an HTTP/1.1 request, which is the header and body, in wire format. This method consults the following fields of the request:

​	Write方法以线路格式编写HTTP/1.1请求，即头部和主体。该方法会查看请求的以下字段：

```
Host
URL
Method (defaults to "GET")
Header
ContentLength
TransferEncoding
Body
```

If Body is present, Content-Length is <= 0 and TransferEncoding hasn't been set to "identity", Write adds "Transfer-Encoding: chunked" to the header. Body is closed after it is sent.

​	如果Body存在，Content-Length小于等于0且TransferEncoding未设置为"identity"，Write将"Transfer-Encoding: chunked"添加到头部。Body在发送后被关闭。

#### (*Request) WriteProxy 

``` go 
func (r *Request) WriteProxy(w io.Writer) error
```

WriteProxy is like Write but writes the request in the form expected by an HTTP proxy. In particular, WriteProxy writes the initial Request-URI line of the request with an absolute URI, per section 5.3 of [RFC 7230](https://rfc-editor.org/rfc/rfc7230.html), including the scheme and host. In either case, WriteProxy also writes a Host header, using either r.Host or r.URL.Host.

​	WriteProxy方法类似于 Write 方法，但是将请求写成 HTTP 代理所期望的格式。特别是，WriteProxy 使用绝对 URI 写入请求的初始 Request-URI 行，根据 [RFC 7230 第 5.3 节](https://rfc-editor.org/rfc/rfc7230.html)，包括方案和主机。无论哪种情况，WriteProxy 还会使用 r.Host 或 r.URL.Host 写入 Host 标头。

### type Response 

``` go 
type Response struct {
	Status     string // Status 表示响应状态，例如 "200 OK"。
	StatusCode int    // StatusCode 表示响应状态码，例如 200
	Proto      string // Proto 表示协议，例如 "HTTP/1.0"。
	ProtoMajor int    // ProtoMajor 表示协议的主要版本，例如 1。
	ProtoMinor int    // ProtoMinor 表示协议的次要版本，例如 0。

    // Header maps header keys to values. If the response had multiple
	// headers with the same key, they may be concatenated, with comma
	// delimiters.  (RFC 7230, section 3.2.2 requires that multiple headers
	// be semantically equivalent to a comma-delimited sequence.) When
	// Header values are duplicated by other fields in this struct (e.g.,
	// ContentLength, TransferEncoding, Trailer), the field values are
	// authoritative.
	//
	// Keys in the map are canonicalized (see CanonicalHeaderKey).
	// Header 将头键映射到值。
    // 如果响应包含相同的键，则它们可能会连接，用逗号分隔。
    //(RFC 7230，第3.2.2节要求多个标题在语义上等同于逗号分隔的序列。)
    // 当 Header 值被此结构中的其他字段
    //(例如 ContentLength、TransferEncoding、Trailer)重复时，
    // 字段值是权威的。
	// 映射中的键已规范化(请参阅 CanonicalHeaderKey)。
	Header Header

    // Body represents the response body.
	//
	// The response body is streamed on demand as the Body field
	// is read. If the network connection fails or the server
	// terminates the response, Body.Read calls return an error.
	//
	// The http Client and Transport guarantee that Body is always
	// non-nil, even on responses without a body or responses with
	// a zero-length body. It is the caller's responsibility to
	// close Body. The default HTTP client's Transport may not
	// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
	// not read to completion and closed.
	//
	// The Body is automatically dechunked if the server replied
	// with a "chunked" Transfer-Encoding.
	//
	// As of Go 1.12, the Body will also implement io.Writer
	// on a successful "101 Switching Protocols" response,
	// as used by WebSockets and HTTP/2's "h2c" mode.
	// Body 表示响应正文。
	//
	// Body 字段读取时，响应正文是按需流式传输的。
    // 如果网络连接失败或服务器终止响应，
    // 则 Body.Read 调用会返回一个错误。
	//
	// http Client 和 Transport 保证 Body 始终是非 nil 的，
    // 即使响应没有正文或正文长度为零。
    // 关闭 Body 是调用方的责任。
    // 如果 Body 没有被完整读取并关闭，
    // 则默认的 HTTP 客户端的 Transport 可能
    // 不会重用 HTTP/1.x "keep-alive" TCP 连接。
	//
	// 如果服务器使用"chunked"传输编码，Body 会自动去块化。
	//
	// 从 Go 1.12 开始，
    // Body 还将在成功的 "101 Switching Protocols" 
    // 响应上实现 io.Writer，
    // 例如 WebSocket 和 HTTP/2 的 "h2c" 模式。
	Body io.ReadCloser

    // ContentLength records the length of the associated content. The
	// value -1 indicates that the length is unknown. Unless Request.Method
	// is "HEAD", values >= 0 indicate that the given number of bytes may
	// be read from Body.
	// ContentLength 记录相关内容的长度。
    // 值 -1 表示长度未知。
    // 除非 Request.Method 是 "HEAD"，
    // 否则值 >= 0 表示可以从 Body 读取给定数量的字节。
	ContentLength int64

    // Contains transfer encodings from outer-most to inner-most. Value is
	// nil, means that "identity" encoding is used.
	// 包含最外层到最内层的传输编码。
    // 如果值为 nil，则表示使用"identity"编码。
	TransferEncoding []string

    // Close records whether the header directed that the connection be
	// closed after reading Body. The value is advice for clients: neither
	// ReadResponse nor Response.Write ever closes a connection.
	// Close 记录了该响应头指示在读取 Body 后是否关闭连接。
    // 该值是对客户端的建议：
	// 无论是 ReadResponse 还是 Response.Write 都不会关闭连接。
	Close bool

    // Uncompressed reports whether the response was sent compressed but
	// was decompressed by the http package. When true, reading from
	// Body yields the uncompressed content instead of the compressed
	// content actually set from the server, ContentLength is set to -1,
	// and the "Content-Length" and "Content-Encoding" fields are deleted
	// from the responseHeader. To get the original response from
	// the server, set Transport.DisableCompression to true.
	// Uncompressed 报告该响应是否被压缩，
    // 但是被 http 包解压缩。当为 true 时，
	// 从 Body 读取会返回解压缩后的内容，
    // 而不是实际从服务器接收到的压缩内容，ContentLength 被设置为 -1，
	// 并且从响应头中删除 "Content-Length" 和 
    // "Content-Encoding" 字段。
	// 要获取来自服务器的原始响应，
    // 请将 Transport.DisableCompression 设置为 true。
	Uncompressed bool

    // Trailer maps trailer keys to values in the same
	// format as Header.
	//
	// The Trailer initially contains only nil values, one for
	// each key specified in the server's "Trailer" header
	// value. Those values are not added to Header.
	//
	// Trailer must not be accessed concurrently with Read calls
	// on the Body.
	//
	// After Body.Read has returned io.EOF, Trailer will contain
	// any trailer values sent by the server.
	// Trailer 以与 Header 相同的格式，将 trailer 的键映射为值。
	//
	// Trailer 最初仅包含 nil 值，
    // 每个值都与服务器的 "Trailer" 头值中指定的每个键一一对应。
	// 这些值不会添加到 Header 中。
	//
	// Trailer 不得与对 Body 的读取调用同时访问。
	//
	// 在 Body.Read 返回 io.EOF 后，
    // Trailer 将包含服务器发送的任何 trailer 值。
	Trailer Header

    // Request is the request that was sent to obtain this Response.
	// Request's Body is nil (having already been consumed).
	// This is only populated for Client requests.
	// Request 是用于获取该响应的请求。
    // Request 的 Body 为 nil(已被使用)。
	// 这仅用于 Client 请求。
	Request *Request

    // TLS contains information about the TLS connection on which the
	// response was received. It is nil for unencrypted responses.
	// The pointer is shared between responses and should not be
	// modified.
	// TLS 包含有关收到响应的 TLS 连接的信息。
    // 对于未加密的响应，它为 nil。
	// 该指针在响应之间共享，不应更改。
	TLS *tls.ConnectionState
}
```

Response represents the response from an HTTP request.

​	Response表示HTTP请求的响应。

The Client and Transport return Responses from servers once the response headers have been received. The response body is streamed on demand as the Body field is read.

​	一旦接收到响应头，客户端和传输机制就会从服务器返回响应。随着读取Body字段的增长，响应体会被按需流式传输。

#### func Get 

``` go 
func Get(url string) (resp *Response, err error)
```

Get issues a GET to the specified URL. If the response is one of the following redirect codes, Get follows the redirect, up to a maximum of 10 redirects:

​	Get方法向指定的URL发出GET请求。如果响应是以下重定向代码之一，则Get会遵循重定向，最多重定向10次：

```
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
308 (Permanent Redirect)
```

An error is returned if there were too many redirects or if there was an HTTP protocol error. A non-2xx response doesn't cause an error. Any returned error will be of type *url.Error. The url.Error value's Timeout method will report true if the request timed out.

​	如果重定向过多或存在HTTP协议错误，则返回错误。非2xx响应不会引起错误。任何返回的错误都将是*url.Error类型。如果请求超时，则url.Error值的Timeout方法将返回true。

When err is nil, resp always contains a non-nil resp.Body. Caller should close resp.Body when done reading from it.

​	当err为nil时，resp始终包含一个非nil resp.Body。调用者在读取完毕后应该关闭resp.Body。

Get is a wrapper around DefaultClient.Get.

​	Get方法是DefaultClient.Get的包装。

To make a request with custom headers, use NewRequest and DefaultClient.Do.

​	要使用自定义标题进行请求，请使用NewRequest和DefaultClient.Do。

To make a request with a specified context.Context, use NewRequestWithContext and DefaultClient.Do.

​	要使用指定的context.Context进行请求，请使用NewRequestWithContext函数和DefaultClient.Do方法。

##### Get Example
``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

```

#### func Head 

``` go 
func Head(url string) (resp *Response, err error)
```

Head issues a HEAD to the specified URL. If the response is one of the following redirect codes, Head follows the redirect, up to a maximum of 10 redirects:

​	Head方法向指定的URL发出HEAD请求。如果响应是以下重定向代码之一，则Head方法会遵循重定向，最多重定向10次：

```
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
308 (Permanent Redirect)
```

Head is a wrapper around DefaultClient.Head.

​	Head方法是DefaultClient.Head的包装函数。

To make a request with a specified context.Context, use NewRequestWithContext and DefaultClient.Do.

​	要使用指定的context.Context发出请求，请使用NewRequestWithContext函数和DefaultClient.Do方法。

#### func Post 

``` go 
func Post(url, contentType string, body io.Reader) (resp *Response, err error)
```

Post issues a POST to the specified URL.

​	`Post`函数向指定的URL发出POST请求。

Caller should close resp.Body when done reading from it.

​	调用者在读取完`resp.Body`后应该关闭它。

If the provided body is an io.Closer, it is closed after the request.

​	如果提供的body实现了`io.Closer`接口，则在请求之后会关闭它。

Post is a wrapper around DefaultClient.Post.

​	`Post`函数是`DefaultClient.Post`的包装函数。

To set custom headers, use NewRequest and DefaultClient.Do.

​	要设置自定义头，请使用`NewRequest`函数和`DefaultClient.Do`方法。

See the Client.Do method documentation for details on how redirects are handled.

​	有关重定向处理方式的详细信息，请参见`Client.Do`方法文档。

To make a request with a specified context.Context, use NewRequestWithContext and DefaultClient.Do.

​	要使用指定的context.Context发出请求，请使用`NewRequestWithContext`函数和`DefaultClient.Do`方法。

#### func PostForm 

``` go 
func PostForm(url string, data url.Values) (resp *Response, err error)
```

PostForm issues a POST to the specified URL, with data's keys and values URL-encoded as the request body.

​	`PostForm`函数向指定的URL发出POST请求，使用`data`的键和值作为请求体进行URL编码。

The Content-Type header is set to application/x-www-form-urlencoded. To set other headers, use NewRequest and DefaultClient.Do.

​	Content-Type标头设置为application/x-www-form-urlencoded。要设置其他标头，请使用`NewRequest`函数和`DefaultClient.Do`方法。

When err is nil, resp always contains a non-nil resp.Body. Caller should close resp.Body when done reading from it.

​	当`err`为nil时，resp始终包含一个非空的`resp.Body`。在读取完毕后，调用者应该关闭resp.Body。

PostForm is a wrapper around DefaultClient.PostForm.

​	`PostForm`函数是`DefaultClient.PostForm`的包装函数。

See the Client.Do method documentation for details on how redirects are handled.

​	有关重定向处理方式的详细信息，请参见Client.Do方法文档。

To make a request with a specified context.Context, use NewRequestWithContext and DefaultClient.Do.

​	要使用指定的context.Context发出请求，请使用`NewRequestWithContext`和`DefaultClient.Do`。

#### func ReadResponse 

``` go 
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
```

ReadResponse reads and returns an HTTP response from r. The req parameter optionally specifies the Request that corresponds to this Response. If nil, a GET request is assumed. Clients must call resp.Body.Close when finished reading resp.Body. After that call, clients can inspect resp.Trailer to find key/value pairs included in the response trailer.

​	`ReadResponse`函数从r中读取并返回HTTP响应。 req参数可选地指定与此Response对应的Request。如果为nil，则假定为GET请求。客户端必须在完成读取resp.Body后调用resp.Body.Close。在该调用之后，客户端可以检查resp.Trailer以查找包含在响应trailer中的键/值对。

#### (*Response) Cookies 

``` go 
func (r *Response) Cookies() []*Cookie
```

Cookies parses and returns the cookies set in the Set-Cookie headers.

​	`Cookies`方法解析并返回在Set-Cookie头中设置的Cookie。

#### (*Response) Location 

``` go 
func (r *Response) Location() (*url.URL, error)
```

Location returns the URL of the response's "Location" header, if present. Relative redirects are resolved relative to the Response's Request. ErrNoLocation is returned if no Location header is present.

​	`Location`方法返回响应头"Location"字段的URL，如果存在的话。相对URL会根据响应的请求进行解析。如果没有Location头，会返回ErrNoLocation。

#### (*Response) ProtoAtLeast 

``` go 
func (r *Response) ProtoAtLeast(major, minor int) bool
```

ProtoAtLeast reports whether the HTTP protocol used in the response is at least major.minor.

​	`ProtoAtLeast`方法报告响应中使用的HTTP协议是否至少为major.minor。

#### (*Response) Write 

``` go 
func (r *Response) Write(w io.Writer) error
```

Write writes r to w in the HTTP/1.x server response format, including the status line, headers, body, and optional trailer.

​	`Write`方法以HTTP/1.x服务器响应格式将r写入w，包括状态行、头部、正文和可选的尾随部分。

This method consults the following fields of the response r:

​	此方法会查询响应r的以下字段：

```
StatusCode
ProtoMajor
ProtoMinor
Request.Method
TransferEncoding
Trailer
Body
ContentLength
Header, 非规范键的值将具有不可预测的行为
```

The Response Body is closed after it is sent.

​	响应主体在发送后将被关闭。

### type ResponseController  <- go1.20

``` go 
type ResponseController struct {
	// contains filtered or unexported fields
}
```

A ResponseController is used by an HTTP handler to control the response.

​	`ResponseController`用于在HTTP处理程序中控制响应。

A ResponseController may not be used after the Handler.ServeHTTP method has returned.

​	在`Handler.ServeHTTP`方法返回之后，不得再使用`ResponseController`。

#### func NewResponseController  <- go1.20

``` go 
func NewResponseController(rw ResponseWriter) *ResponseController
```

NewResponseController creates a ResponseController for a request.

​	`NewResponseController`函数为请求创建一个ResponseController。

The ResponseWriter should be the original value passed to the Handler.ServeHTTP method, or have an Unwrap method returning the original ResponseWriter.

​	`ResponseWriter`应该是传递给Handler.ServeHTTP方法的原始值，或者具有返回原始ResponseWriter的Unwrap方法。

If the ResponseWriter implements any of the following methods, the ResponseController will call them as appropriate:

​	如果`ResponseWriter`实现了以下任何方法，则ResponseController将根据需要调用它们：

```
Flush()
FlushError() error // 替代Flush，返回错误
Hijack() (net.Conn, *bufio.ReadWriter, error)
SetReadDeadline(deadline time.Time) error
SetWriteDeadline(deadline time.Time) error
```

If the ResponseWriter does not support a method, ResponseController returns an error matching ErrNotSupported.

​	如果`ResponseWriter`不支持某个方法，则ResponseController返回与ErrNotSupported相匹配的错误。

#### (*ResponseController) EnableFullDuplex <-go1.21.0

```go
func (c *ResponseController) EnableFullDuplex() error
```

EnableFullDuplex indicates that the request handler will interleave reads from Request.Body with writes to the ResponseWriter.

For HTTP/1 requests, the Go HTTP server by default consumes any unread portion of the request body before beginning to write the response, preventing handlers from concurrently reading from the request and writing the response. Calling EnableFullDuplex disables this behavior and permits handlers to continue to read from the request while concurrently writing the response.

For HTTP/2 requests, the Go HTTP server always permits concurrent reads and responses.



#### (*ResponseController) Flush  <- go1.20

``` go 
func (c *ResponseController) Flush() error
```

Flush flushes buffered data to the client.

​	`Flush`方法将缓冲数据刷新到客户端。

#### (*ResponseController) Hijack  <- go1.20

``` go 
func (c *ResponseController) Hijack() (net.Conn, *bufio.ReadWriter, error)
```

Hijack lets the caller take over the connection. See the Hijacker interface for details.

​	`Hijack`方法允许调用者接管连接。有关详细信息，请参见Hijacker接口。

#### (*ResponseController) SetReadDeadline  <- go1.20

``` go 
func (c *ResponseController) SetReadDeadline(deadline time.Time) error
```

SetReadDeadline sets the deadline for reading the entire request, including the body. Reads from the request body after the deadline has been exceeded will return an error. A zero value means no deadline.

​	`SetReadDeadline`方法设置读取整个请求(包括正文)的截止日期。超过截止日期后从请求正文中读取会返回错误。零值表示没有截止日期。

Setting the read deadline after it has been exceeded will not extend it.

​	在超过期限后设置读取期限将不会延长它。

#### (*ResponseController) SetWriteDeadline  <- go1.20

``` go 
func (c *ResponseController) SetWriteDeadline(deadline time.Time) error
```

SetWriteDeadline sets the deadline for writing the response. Writes to the response body after the deadline has been exceeded will not block, but may succeed if the data has been buffered. A zero value means no deadline.

​	`SetWriteDeadline`方法为写入响应设置截止日期。在截止日期之后向响应正文写入不会阻塞，但如果数据已被缓冲，则可能成功。零值表示没有截止日期。

Setting the write deadline after it has been exceeded will not extend it.

​	在超过期限后设置写入期限将不会延长它。

### type ResponseWriter 

``` go 
type ResponseWriter interface {
    // Header returns the header map that will be sent by
	// WriteHeader. The Header map also is the mechanism with which
	// Handlers can set HTTP trailers.
	//
	// Changing the header map after a call to WriteHeader (or
	// Write) has no effect unless the HTTP status code was of the
	// 1xx class or the modified headers are trailers.
	//
	// There are two ways to set Trailers. The preferred way is to
	// predeclare in the headers which trailers you will later
	// send by setting the "Trailer" header to the names of the
	// trailer keys which will come later. In this case, those
	// keys of the Header map are treated as if they were
	// trailers. See the example. The second way, for trailer
	// keys not known to the Handler until after the first Write,
	// is to prefix the Header map keys with the TrailerPrefix
	// constant value. See TrailerPrefix.
	//
	// To suppress automatic response headers (such as "Date"), set
	// their value to nil.
	// Header 返回将由 WriteHeader 发送的标头映射。
    // Header 映射也是处理程序设置 HTTP 尾随项的机制。
	//
	// 在调用 WriteHeader(或 Write)
    // 之后更改标头映射对状态码为 1xx 类的 
    // HTTP 状态码或修改后的标头为尾随项时，才会生效。
	//
	// 有两种方法可以设置 Trailers。
    // 首选的方式是在标头中预先声明，您稍后将发送哪些拖车，
    // 通过将"Trailer"标头设置为稍后将到来的拖车键的名称来完成。
    // 在这种情况下，Header 映射的那些键被视为拖车。
    // 请参阅示例。
    // 第二种方式是针对处理程序直到第一次写入后才知道的尾随键，
    // 在 Header 映射键中添加 TrailerPrefix 常量值前缀。
    // 请参阅 TrailerPrefix。
	//
	// 要抑制自动响应标头(例如"Date")，请将其值设置为 nil。
	Header() Header

    // Write writes the data to the connection as part of an HTTP reply.
	//
	// If WriteHeader has not yet been called, Write calls
	// WriteHeader(http.StatusOK) before writing the data. If the Header
	// does not contain a Content-Type line, Write adds a Content-Type set
	// to the result of passing the initial 512 bytes of written data to
	// DetectContentType. Additionally, if the total size of all written
	// data is under a few KB and there are no Flush calls, the
	// Content-Length header is added automatically.
	//
	// Depending on the HTTP protocol version and the client, calling
	// Write or WriteHeader may prevent future reads on the
	// Request.Body. For HTTP/1.x requests, handlers should read any
	// needed request body data before writing the response. Once the
	// headers have been flushed (due to either an explicit Flusher.Flush
	// call or writing enough data to trigger a flush), the request body
	// may be unavailable. For HTTP/2 requests, the Go HTTP server permits
	// handlers to continue to read the request body while concurrently
	// writing the response. However, such behavior may not be supported
	// by all HTTP/2 clients. Handlers should read before writing if
	// possible to maximize compatibility.
	// Write 在 HTTP 响应的一部分将数据写入连接。
	//
	// 如果尚未调用 WriteHeader，
    // 则 Write 在写入数据之前调用 WriteHeader(http.StatusOK)。
    // 如果标头不包含 Content-Type 行，
    // 则 Write 添加一个 Content-Type 集，
    // 将初始 512 字节的写入数据传递给 DetectContentType 的结果。
    // 此外，如果所有写入数据的总大小小于几 KB，
    // 且没有 Flush 调用，则会自动添加 Content-Length 标头。
	//
	// 根据 HTTP 协议版本和客户端，
    // 调用 Write 或 WriteHeader 可能会阻止 Request.Body 
    // 上的未来读取。对于 HTTP/1.x 请求，
    // 处理程序应在编写响应之前读取任何所需的请求正文数据。
    // 一旦标头已被刷新(由于显式调用 Flusher.Flush 
    // 或写入足够的数据以触发刷新)，请求正文可能不可用。
    // 对于 HTTP/2 请求，Go HTTP 服务器允许处理程序
    // 在并发写入响应时继续读取请求正文。
    // 但是，并非所有 HTTP/2 客户端都支持这种行为。
    // 如果可能的话，处理程序应在编写之前进行读取，以最大程度地提高兼容性。
	Write([]byte) (int, error)

    // WriteHeader sends an HTTP response header with the provided
	// status code.
	//
	// If WriteHeader is not called explicitly, the first call to Write
	// will trigger an implicit WriteHeader(http.StatusOK).
	// Thus explicit calls to WriteHeader are mainly used to
	// send error codes or 1xx informational responses.
	//
	// The provided code must be a valid HTTP 1xx-5xx status code.
	// Any number of 1xx headers may be written, followed by at most
	// one 2xx-5xx header. 1xx headers are sent immediately, but 2xx-5xx
	// headers may be buffered. Use the Flusher interface to send
	// buffered data. The header map is cleared when 2xx-5xx headers are
	// sent, but not with 1xx headers.
	//
	// The server will automatically send a 100 (Continue) header
	// on the first read from the request body if the request has
	// an "Expect: 100-continue" header.
	// WriteHeader 使用提供的状态码发送 HTTP 响应头。
	//
	// 如果没有显式调用 WriteHeader，
    // 第一次调用 Write 将会触发一个隐式的 
    // WriteHeader(http.StatusOK)。
	// 因此，显式调用 WriteHeader 主要用于发送错误代码
    // 或 1xx 的信息响应。
	//
	// 提供的代码必须是有效的 HTTP 1xx-5xx 状态码。
	// 可以写入任意数量的 1xx 头，其后最多一个 2xx-5xx 头。
    // 1xx 头会立即发送，但 2xx-5xx 头可能会被缓冲。
	// 使用 Flusher 接口发送缓冲的数据。
    // 在发送 2xx-5xx 头时，头映射将被清除，但在 1xx 头时不会被清除。
	//
	// 如果请求有一个"Expect: 100-continue"头，
    // 服务器将在第一次从请求正文读取时自动发送一个 100(Continue)头。
	WriteHeader(statusCode int)
}
```

A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.

​	`ResponseWriter`接口由HTTP处理程序用于构建HTTP响应。

A ResponseWriter may not be used after the Handler.ServeHTTP method has returned.

​	在Handler.ServeHTTP方法返回之后，不得再使用ResponseWriter。

##### Example (Trailers)

HTTP Trailers are a set of key/value pairs like headers that come after the HTTP response, instead of before.

​	HTTP 尾部标签(HTTP Trailers)是一组键/值对，类似于报头(headers)，但是它们出现在 HTTP 响应的末尾而不是前面。

``` go 
package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/sendstrailers", func(w http.ResponseWriter, req *http.Request) {
		// Before any call to WriteHeader or Write, declare
		// the trailers you will set during the HTTP
		// response. These three headers are actually sent in
		// the trailer.
		w.Header().Set("Trailer", "AtEnd1, AtEnd2")
		w.Header().Add("Trailer", "AtEnd3")

		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)

		w.Header().Set("AtEnd1", "value 1")
		io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
		w.Header().Set("AtEnd2", "value 2")
		w.Header().Set("AtEnd3", "value 3") // These will appear as trailers.
	})
}

```

### type RoundTripper 

``` go 
type RoundTripper interface {
    // RoundTrip executes a single HTTP transaction, returning
	// a Response for the provided Request.
	//
	// RoundTrip should not attempt to interpret the response. In
	// particular, RoundTrip must return err == nil if it obtained
	// a response, regardless of the response's HTTP status code.
	// A non-nil err should be reserved for failure to obtain a
	// response. Similarly, RoundTrip should not attempt to
	// handle higher-level protocol details such as redirects,
	// authentication, or cookies.
	//
	// RoundTrip should not modify the request, except for
	// consuming and closing the Request's Body. RoundTrip may
	// read fields of the request in a separate goroutine. Callers
	// should not mutate or reuse the request until the Response's
	// Body has been closed.
	//
	// RoundTrip must always close the body, including on errors,
	// but depending on the implementation may do so in a separate
	// goroutine even after RoundTrip returns. This means that
	// callers wanting to reuse the body for subsequent requests
	// must arrange to wait for the Close call before doing so.
	//
	// The Request's URL and Header fields must be initialized.
	// RoundTrip执行单个HTTP事务，为提供的Request返回一个Response。
	//
	// RoundTrip不应该试图解释响应。
    // 特别是，无论响应的HTTP状态码如何，
    // RoundTrip必须返回err == nil。
	// 非nil err 应该保留给无法获取响应的情况。
    // 同样，RoundTrip不应尝试处理高级协议细节，
    // 例如重定向，身份验证或cookie。
	//
	// RoundTrip不应修改请求，除了消耗和关闭请求的Body。 
    // RoundTrip可以在单独的goroutine中读取请求的字段。
    // 调用方不应在关闭响应的Body之前更改或重用请求。
	//
	// RoundTrip必须始终关闭Body，即使出现错误，
    // 但是根据实现可能会在RoundTrip返回后在单独的goroutine中执行关闭。
    // 这意味着，调用者想要重用Body进行后续请求必须
    // 在进行此操作之前安排等待Close调用。
	//
	// Request的URL和Header字段必须初始化。
	RoundTrip(*Request) (*Response, error)
}
```

RoundTripper is an interface representing the ability to execute a single HTTP transaction, obtaining the Response for a given Request.

​	`RoundTripper`是一个接口，表示执行单个HTTP事务的能力，获得给定请求的响应。

A RoundTripper must be safe for concurrent use by multiple goroutines.

​	`RoundTripper`必须对多个goroutine进行并发使用的安全。

``` go 
var DefaultTransport RoundTripper = &Transport{
	Proxy: ProxyFromEnvironment,
	DialContext: defaultTransportDialContext(&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}),
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
```

DefaultTransport is the default implementation of Transport and is used by DefaultClient. It establishes network connections as needed and caches them for reuse by subsequent calls. It uses HTTP proxies as directed by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions thereof).

​	`DefaultTransport`是Transport的默认实现，并由DefaultClient使用。它根据需要建立网络连接并将其缓存以供后续调用重用。它按照环境变量HTTP_PROXY、HTTPS_PROXY和NO_PROXY(或其小写版本)的指示使用HTTP代理。

#### func NewFileTransport 

``` go 
func NewFileTransport(fs FileSystem) RoundTripper
```

NewFileTransport returns a new RoundTripper, serving the provided FileSystem. The returned RoundTripper ignores the URL host in its incoming requests, as well as most other properties of the request.

​	`NewFileTransport`函数返回一个新的`RoundTripper`，为所提供的`FileSystem`提供服务。返回的RoundTripper忽略其传入请求中的URL主机，以及请求的大多数其他属性。

The typical use case for NewFileTransport is to register the "file" protocol with a Transport, as in:

​	`NewFileTransport`函数的典型用例是向Transport注册"file"协议，例如：

```
t := &http.Transport{}
t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
c := &http.Client{Transport: t}
res, err := c.Get("file:///etc/passwd")
...
```

#### func NewFileTransportFS <- go1.22.0

```
func NewFileTransportFS(fsys fs.FS) RoundTripper
```

NewFileTransportFS returns a new [RoundTripper](https://pkg.go.dev/net/http@go1.23.0#RoundTripper), serving the provided file system fsys. The returned RoundTripper ignores the URL host in its incoming requests, as well as most other properties of the request. The files provided by fsys must implement [io.Seeker](https://pkg.go.dev/io#Seeker).

​	`NewFileTransportFS` 返回一个新的 [RoundTripper](https://pkg.go.dev/net/http@go1.23.0#RoundTripper)，用来处理提供的文件系统 `fsys`。返回的 `RoundTripper` 会忽略其传入请求中的 URL 主机以及请求的大多数其他属性。`fsys` 提供的文件必须实现 [io.Seeker](https://pkg.go.dev/io#Seeker) 接口。

The typical use case for NewFileTransportFS is to register the "file" protocol with a [Transport](https://pkg.go.dev/net/http@go1.23.0#Transport), as in:

​	`NewFileTransportFS` 的典型用例是将 "file" 协议与 [Transport](https://pkg.go.dev/net/http@go1.23.0#Transport) 绑定注册，如下所示：

```
fsys := os.DirFS("/")
t := &http.Transport{}
t.RegisterProtocol("file", http.NewFileTransportFS(fsys))
c := &http.Client{Transport: t}
res, err := c.Get("file:///etc/passwd")
...
```

### type SameSite  <- go1.11

``` go 
type SameSite int
```

SameSite allows a server to define a cookie attribute making it impossible for the browser to send this cookie along with cross-site requests. The main goal is to mitigate the risk of cross-origin information leakage, and provide some protection against cross-site request forgery attacks.

​	`SameSite` 允许服务器定义一个 cookie 属性，使浏览器无法将该 cookie 与跨站请求一起发送。主要目的是减少跨源信息泄露的风险，并提供一定的保护，以防止跨站请求伪造攻击。

See https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00 for details.

​	详见 https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00。

``` go 
const (
	SameSiteDefaultMode SameSite = iota + 1
	SameSiteLaxMode
	SameSiteStrictMode
	SameSiteNoneMode
)
```

### type ServeMux 

``` go 
type ServeMux struct {
	// 包含过滤或未公开的字段
}
```

ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

​	`ServeMux` 是一个 HTTP 请求多路复用器。它会根据已注册的模式列表匹配每个传入请求的 URL，并调用与该 URL 最匹配的模式对应的处理器。

#### 模式 Patterns 

Patterns can match the method, host and path of a request. Some examples:

​	模式可以匹配请求的方法、主机和路径。例如：

- "/index.html" matches the path "/index.html" for any host and method.
- "/index.html" 匹配任何主机和方法下的路径 "/index.html"。

- "GET /static/" matches a GET request whose path begins with "/static/".
- "GET /static/" 匹配路径以 "/static/" 开头的 GET 请求。
- "example.com/" matches any request to the host "example.com".
- "example.com/" 匹配主机为 "example.com" 的任何请求。
- "example.com/{$}" matches requests with host "example.com" and path "/".
- "example.com/{$}" 匹配主机为 "example.com" 且路径为 "/" 的请求。
- "/b/{bucket}/o/{objectname...}" matches paths whose first segment is "b" and whose third segment is "o". The name "bucket" denotes the second segment and "objectname" denotes the remainder of the path.
- "/b/{bucket}/o/{objectname...}" 匹配路径中第一个片段为 "b" 且第三个片段为 "o" 的请求。名称 "bucket" 表示第二个片段，"objectname" 表示路径的其余部分。

In general, a pattern looks like

​	一般来说，一个模式的格式如下：

```
[METHOD ][HOST]/[PATH]
```

All three parts are optional; "/" is a valid pattern. If METHOD is present, it must be followed by at least one space or tab.

​	这三个部分都是可选的；"/" 是一个有效的模式。如果指定了 METHOD，则它后面必须至少有一个空格或制表符。

Literal (that is, non-wildcard) parts of a pattern match the corresponding parts of a request case-sensitively.

​	模式中字面的部分（即非通配符的部分）与请求的相应部分是区分大小写的。

A pattern with no method matches every method. A pattern with the method GET matches both GET and HEAD requests. Otherwise, the method must match exactly.

​	没有方法的模式匹配所有方法。带有 GET 方法的模式同时匹配 GET 和 HEAD 请求。否则，方法必须精确匹配。

A pattern with no host matches every host. A pattern with a host matches URLs on that host only.

​	没有主机的模式匹配所有主机。带有主机的模式仅匹配该主机上的 URL。

A path can include wildcard segments of the form {NAME} or {NAME...}. For example, "/b/{bucket}/o/{objectname...}". The wildcard name must be a valid Go identifier. Wildcards must be full path segments: they must be preceded by a slash and followed by either a slash or the end of the string. For example, "/b_{bucket}" is not a valid pattern.

​	路径可以包含形如 `{NAME}` 或 `{NAME...}` 的通配符片段。例如，"/b/{bucket}/o/{objectname...}"。通配符名称必须是有效的 Go 标识符。通配符必须是完整的路径片段：它们必须以斜杠开头，并且以斜杠或字符串结尾。例如，"/b_{bucket}" 不是有效的模式。

Normally a wildcard matches only a single path segment, ending at the next literal slash (not %2F) in the request URL. But if the "..." is present, then the wildcard matches the remainder of the URL path, including slashes. (Therefore it is invalid for a "..." wildcard to appear anywhere but at the end of a pattern.) The match for a wildcard can be obtained by calling [Request.PathValue](https://pkg.go.dev/net/http@go1.23.0#Request.PathValue) with the wildcard's name. A trailing slash in a path acts as an anonymous "..." wildcard.

​	通常情况下，通配符只匹配单个路径片段，匹配到下一个字面斜杠（不是 %2F）时结束。但如果通配符中包含 "..."，则该通配符匹配 URL 路径的其余部分，包括斜杠。（因此，"..." 通配符只能出现在模式的末尾。）可以通过调用 [Request.PathValue](https://pkg.go.dev/net/http@go1.23.0#Request.PathValue) 并提供通配符的名称来获取通配符的匹配结果。路径中的尾随斜杠相当于一个匿名的 "..." 通配符。

The special wildcard {$} matches only the end of the URL. For example, the pattern "/{$}" matches only the path "/", whereas the pattern "/" matches every path.

​	特殊通配符 `{$}` 仅匹配 URL 的结尾。例如，模式 "/{$}" 只匹配路径 "/"，而模式 "/" 则匹配所有路径。

For matching, both pattern paths and incoming request paths are unescaped segment by segment. So, for example, the path "/a%2Fb/100%25" is treated as having two segments, "a/b" and "100%". The pattern "/a%2fb/" matches it, but the pattern "/a/b/" does not.

​	对于匹配来说，模式路径和传入的请求路径都逐段进行解码。因此，例如路径 "/a%2Fb/100%25" 会被视为两个片段："a/b" 和 "100%"。模式 "/a%2fb/" 可以匹配该路径，但模式 "/a/b/" 不能。

#### 优先级 Precedence 

If two or more patterns match a request, then the most specific pattern takes precedence. A pattern P1 is more specific than P2 if P1 matches a strict subset of P2’s requests; that is, if P2 matches all the requests of P1 and more. If neither is more specific, then the patterns conflict. There is one exception to this rule, for backwards compatibility: if two patterns would otherwise conflict and one has a host while the other does not, then the pattern with the host takes precedence. If a pattern passed to [ServeMux.Handle](https://pkg.go.dev/net/http@go1.23.0#ServeMux.Handle) or [ServeMux.HandleFunc](https://pkg.go.dev/net/http@go1.23.0#ServeMux.HandleFunc) conflicts with another pattern that is already registered, those functions panic.

​	如果两个或多个模式匹配一个请求，则最具体的模式优先。模式 P1 比 P2 更具体是指，P1 匹配的请求严格是 P2 的子集；也就是说，P2 匹配 P1 的所有请求以及更多的请求。如果没有一个模式更具体，那么这些模式之间存在冲突。有一个例外，为了向后兼容：如果两个模式发生冲突，并且其中一个有主机而另一个没有，那么带有主机的模式优先。如果传递给 [ServeMux.Handle](https://pkg.go.dev/net/http@go1.23.0#ServeMux.Handle) 或 [ServeMux.HandleFunc](https://pkg.go.dev/net/http@go1.23.0#ServeMux.HandleFunc) 的模式与已经注册的模式冲突，这些函数会引发 panic。

As an example of the general rule, "/images/thumbnails/" is more specific than "/images/", so both can be registered. The former matches paths beginning with "/images/thumbnails/" and the latter will match any other path in the "/images/" subtree.

​	例如，"/images/thumbnails/" 比 "/images/" 更具体，因此两者都可以注册。前者匹配以 "/images/thumbnails/" 开头的路径，后者则匹配 "/images/" 子树中的任何其他路径。

As another example, consider the patterns "GET /" and "/index.html": both match a GET request for "/index.html", but the former pattern matches all other GET and HEAD requests, while the latter matches any request for "/index.html" that uses a different method. The patterns conflict.

​	另一个例子，考虑模式 "GET /" 和 "/index.html"：两者都匹配路径为 "/index.html" 的 GET 请求，但前者模式匹配所有其他 GET 和 HEAD 请求，而后者则匹配使用不同方法的 "/index.html" 请求。模式之间存在冲突。

#### 尾部斜杠重定向 Trailing-slash redirection 

Consider a [ServeMux](https://pkg.go.dev/net/http@go1.23.0#ServeMux) with a handler for a subtree, registered using a trailing slash or "..." wildcard. If the ServeMux receives a request for the subtree root without a trailing slash, it redirects the request by adding the trailing slash. This behavior can be overridden with a separate registration for the path without the trailing slash or "..." wildcard. For example, registering "/images/" causes ServeMux to redirect a request for "/images" to "/images/", unless "/images" has been registered separately.

​	考虑一个 [ServeMux](https://pkg.go.dev/net/http@go1.23.0#ServeMux)，它为一个子树注册了一个带有尾部斜杠或 "..." 通配符的处理器。如果 ServeMux 接收到对该子树根路径的请求，但请求路径没有尾部斜杠，它会通过添加尾部斜杠来重定向请求。可以通过单独为没有尾部斜杠或 "..." 通配符的路径进行注册来覆盖这种行为。例如，注册 "/images/" 会导致 ServeMux 将 "/images" 的请求重定向到 "/images/"，除非 "/images" 已单独注册。

#### 请求清理 Request sanitizing 

ServeMux also takes care of sanitizing the URL request path and the Host header, stripping the port number and redirecting any request containing . or .. segments or repeated slashes to an equivalent, cleaner URL.

​	ServeMux 还负责清理 URL 请求路径和 Host 头部，去除端口号，并将包含 `.` 或 `..` 片段或重复斜杠的请求重定向到等效的、更简洁的 URL。

#### 兼容性 Compatibility 

The pattern syntax and matching behavior of ServeMux changed significantly in Go 1.22. To restore the old behavior, set the GODEBUG environment variable to "httpmuxgo121=1". This setting is read once, at program startup; changes during execution will be ignored.

​	`ServeMux` 的模式语法和匹配行为在 Go 1.22 中发生了显著变化。要恢复旧行为，可以将 GODEBUG 环境变量设置为 "httpmuxgo121=1"。此设置在程序启动时读取；执行期间的更改将被忽略。

The backwards-incompatible changes include:

​	向后不兼容的更改包括：

- Wildcards are just ordinary literal path segments in 1.21. For example, the pattern "/{x}" will match only that path in 1.21, but will match any one-segment path in 1.22.
- 通配符在 1.21 中只是普通的字面路径片段。例如，模式 "/{x}" 在 1.21 中只会匹配该路径，但在 1.22 中会匹配任何单片段路径。
- In 1.21, no pattern was rejected, unless it was empty or conflicted with an existing pattern. In 1.22, syntactically invalid patterns will cause [ServeMux.Handle](https://pkg.go.dev/net/http@go1.23.0#ServeMux.Handle) and [ServeMux.HandleFunc](https://pkg.go.dev/net/http@go1.23.0#ServeMux.HandleFunc) to panic. For example, in 1.21, the patterns "/{" and "/a{x}" match themselves, but in 1.22 they are invalid and will cause a panic when registered.
- 在 1.21 中，除非模式为空或与现有模式冲突，否则不会拒绝任何模式。在 1.22 中，语法无效的模式会导致 [ServeMux.Handle](https://pkg.go.dev/net/http@go1.23.0#ServeMux.Handle) 和 [ServeMux.HandleFunc](https://pkg.go.dev/net/http@go1.23.0#ServeMux.HandleFunc) 引发 panic。例如，在 1.21 中，模式 "/{" 和 "/a{x}" 匹配它们自身，但在 1.22 中它们是无效的，并且注册时会引发 panic。
- In 1.22, each segment of a pattern is unescaped; this was not done in 1.21. For example, in 1.22 the pattern "/%61" matches the path "/a" ("%61" being the URL escape sequence for "a"), but in 1.21 it would match only the path "/%2561" (where "%25" is the escape for the percent sign).
- 在 1.22 中，模式的每个片段都会被解码；而在 1.21 中不会。例如，在 1.22 中，模式 "/%61" 匹配路径 "/a"（"%61" 是 "a" 的 URL 转义序列），但在 1.21 中它只匹配路径 "/%2561"（其中 "%25" 是百分号的转义序列）。
- When matching patterns to paths, in 1.22 each segment of the path is unescaped; in 1.21, the entire path is unescaped. This change mostly affects how paths with %2F escapes adjacent to slashes are treated. See https://go.dev/issue/21955 for details.
- 在模式与路径匹配时，在 1.22 中路径的每个片段都会被解码；而在 1.21 中整个路径会被解码。此更改主要影响带有紧邻斜杠的 %2F 转义的路径的处理方式。有关详细信息，请参阅 https://go.dev/issue/21955。

#### func NewServeMux 

``` go 
func NewServeMux() *ServeMux
```

NewServeMux allocates and returns a new ServeMux.

​	`NewServeMux`函数分配并返回一个新的 ServeMux。

#### (*ServeMux) Handle 

``` go 
func (mux *ServeMux) Handle(pattern string, handler Handler)
```

Handle registers the handler for the given pattern. If a handler already exists for pattern, Handle panics.

​	`Handle`方法为给定的模式注册处理程序。如果已存在处理程序，则 Handle 方法会引发 panic。

##### Handle Example
``` go 
package main

import (
	"fmt"
	"net/http"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})
}

```

#### (*ServeMux) HandleFunc 

``` go 
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

HandleFunc registers the handler function for the given pattern.

​	`HandleFunc`方法为给定的模式注册处理函数。

#### (*ServeMux) Handler  <- go1.1

``` go 
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
```

Handler returns the handler to use for the given request, consulting r.Method, r.Host, and r.URL.Path. It always returns a non-nil handler. If the path is not in its canonical form, the handler will be an internally-generated handler that redirects to the canonical path. If the host contains a port, it is ignored when matching handlers.

​	`Handler`方法根据 r.Method、r.Host 和 r.URL.Path 来获取给定请求使用的处理程序。它始终返回一个非 nil 的处理程序。如果路径不在其规范形式中，则处理程序将是一个内部生成的处理程序，用于重定向到规范路径。如果主机包含端口，则匹配处理程序时将忽略该端口。

The path and host are used unchanged for CONNECT requests.

​	对于 CONNECT 请求，路径和主机都不会被修改。

Handler also returns the registered pattern that matches the request or, in the case of internally-generated redirects, the pattern that will match after following the redirect.

​	`Handler`方法还返回与请求匹配的已注册模式，或在生成重定向的情况下，将在遵循重定向后匹配的模式。

If there is no registered handler that applies to the request, Handler returns a “page not found” handler and an empty pattern.

​	如果没有已注册的处理程序适用于请求，则 Handler方法返回一个"页面未找到"的处理程序和一个空模式。

#### (*ServeMux) ServeHTTP 

``` go 
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL.

​	`ServeHTTP`方法将请求分派到模式最接近请求 URL 的处理程序。

### type Server 

``` go 
type Server struct {
    // Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	// Addr 可选地指定服务器监听的 TCP 地址，
    // 格式为 "host:port"。如果为空，则使用 ":http"(端口80)。
	// 服务名称在 RFC 6335 中定义，
    // 并由 IANA 分配。有关地址格式的详细信息，请参见 net.Dial。
	Addr string

    // handler to invoke, http.DefaultServeMux if nil
    // 处理程序，如果为nil，则使用http.DefaultServeMux
	Handler Handler 

    // DisableGeneralOptionsHandler, if true, passes "OPTIONS *" requests to the Handler,
	// otherwise responds with 200 OK and Content-Length: 0.
	// 如果为 true，则禁用通用 OPTIONS 处理程序，
    // 并将 "OPTIONS *" 请求传递给处理程序，
    // 否则响应为 200 OK 并带有 Content-Length：0。
	DisableGeneralOptionsHandler bool

    // TLSConfig optionally provides a TLS configuration for use
	// by ServeTLS and ListenAndServeTLS. Note that this value is
	// cloned by ServeTLS and ListenAndServeTLS, so it's not
	// possible to modify the configuration with methods like
	// tls.Config.SetSessionTicketKeys. To use
	// SetSessionTicketKeys, use Server.Serve with a TLS Listener
	// instead.
	// TLSConfig 可选地提供用于 ServeTLS 和 
    // ListenAndServeTLS 的 TLS 配置。
    // 请注意，此值将由 ServeTLS 和 ListenAndServeTLS 克隆，
    // 因此无法使用 tls.Config.SetSessionTicketKeys 
    // 等方法修改配置。要使用 SetSessionTicketKeys，
    // 请改用具有 TLS 监听器的 Server.Serve。
	TLSConfig *tls.Config

    // ReadTimeout is the maximum duration for reading the entire
	// request, including the body. A zero or negative value means
	// there will be no timeout.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	// ReadTimeout 是读取整个请求(包括正文)的最长时间。
    // 零值或负值表示没有超时。
	//
	// 由于 ReadTimeout 不允许处理程序对每个请求正文的
    // 可接受截止时间或上传速率做出每个请求的决策，
    // 因此大多数用户将更喜欢使用 ReadHeaderTimeout。
    // 使用这两个值是有效的。
	ReadTimeout time.Duration

    // ReadHeaderTimeout is the amount of time allowed to read
	// request headers. The connection's read deadline is reset
	// after reading the headers and the Handler can decide what
	// is considered too slow for the body. If ReadHeaderTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	// ReadHeaderTimeout 是允许读取请求标头的时间量。
    // 读取标头后，连接的读取截止时间将重置，
    // 处理程序可以决定什么被认为是请求主体的过慢的响应时间。
    // 如果 ReadHeaderTimeout 为零，
    // 则使用 ReadTimeout 的值。如果两者都为零，则没有超时。
	ReadHeaderTimeout time.Duration

    // WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	// A zero or negative value means there will be no timeout.
	// WriteTimeout 是在超时写入响应之前的最长持续时间。
    // 每次读取新请求的标头时都会重置它。
    // 与 ReadTimeout 一样，它不允许处理程序对每个请求做出决策。
    // 零值或负值表示没有超时。
	WriteTimeout time.Duration

    // IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	// IdleTimeout 是在启用保持连接的情况下等待下一个请求的最长时间。
    // 如果 IdleTimeout 为零，则使用 ReadTimeout 的值。
    // 如果两者都为零，则没有超时时间。
	IdleTimeout time.Duration

    // MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	// If zero, DefaultMaxHeaderBytes is used.
	// MaxHeaderBytes 控制服务器在解析请求标头的键和值
    //(包括请求行)时读取的最大字节数。它不限制请求主体的大小。
    // 如果为零，则使用 DefaultMaxHeaderBytes。
	MaxHeaderBytes int

    // TLSNextProto optionally specifies a function to take over
	// ownership of the provided TLS connection when an ALPN
	// protocol upgrade has occurred. The map key is the protocol
	// name negotiated. The Handler argument should be used to
	// handle HTTP requests and will initialize the Request's TLS
	// and RemoteAddr if not already set. The connection is
	// automatically closed when the function returns.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	// TLSNextProto 可选地指定一个函数，
    // 以在 ALPN 协议升级发生时接管提供的 TLS 连接所有权。
    // 映射键是协议协商的协议名称。
    // Handler 参数应用于处理 HTTP 请求，
    // 并将初始化请求的 TLS 和 RemoteAddr(如果尚未设置)。
    // 该函数返回时连接将自动关闭。
    // 如果 TLSNextProto 不为 nil，则不会自动启用 HTTP/2 支持。
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

    // ConnState specifies an optional callback function that is
	// called when a client connection changes state. See the
	// ConnState type and associated constants for details.
	// ConnState 指定一个可选的回调函数，
    // 当客户端连接更改状态时调用。
    // 有关详细信息，请参见 ConnState 类型和相关常量。
	ConnState func(net.Conn, ConnState)

    // ErrorLog specifies an optional logger for errors accepting
	// connections, unexpected behavior from handlers, and
	// underlying FileSystem errors.
	// If nil, logging is done via the log package's standard logger.
	// ErrorLog 指定一个可选的记录器，
    // 用于记录接受连接时的错误、处理程序的意外行为以及
    // 基础 FileSystem 的错误。
    // 如果为 nil，则通过 log 包的标准记录器进行记录。
	ErrorLog *log.Logger

    // BaseContext optionally specifies a function that returns
	// the base context for incoming requests on this server.
	// The provided Listener is the specific Listener that's
	// about to start accepting requests.
	// If BaseContext is nil, the default is context.Background().
	// If non-nil, it must return a non-nil context.
	// BaseContext 可选地指定一个函数，
    // 该函数返回此服务器上的传入请求的基本上下文。
    // 提供的 Listener 是即将开始接受请求的特定 Listener。
    // 如果 BaseContext 为 nil，
    // 则默认为 context.Background()。
    // 如果非 nil，则必须返回非 nil 上下文。
	BaseContext func(net.Listener) context.Context

    // ConnContext optionally specifies a function that modifies
	// the context used for a new connection c. The provided ctx
	// is derived from the base context and has a ServerContextKey
	// value.
	// ConnContext 可选地指定一个函数，
    // 该函数修改用于新连接 c 的上下文。
    // 提供的 ctx 派生自基本上下文并具有 ServerContextKey 值。
	ConnContext func(ctx context.Context, c net.Conn) context.Context
	// 包含已过滤或未公开的字段
}
```

A Server defines parameters for running an HTTP server. The zero value for Server is a valid configuration.

​	`Server`定义了运行HTTP服务器的参数。`Server`的零值是有效的配置。

#### (*Server) Close  <- go1.8

``` go 
func (srv *Server) Close() error
```

Close immediately closes all active net.Listeners and any connections in state StateNew, StateActive, or StateIdle. For a graceful shutdown, use Shutdown.

​	`Close`方法会立即关闭所有活动的`net.Listeners`和状态为`StateNew`、`StateActive`或`StateIdle`的任何连接。为了优雅的关闭，请使用Shutdown方法。

Close does not attempt to close (and does not even know about) any hijacked connections, such as WebSockets.

​	`Close`方法不会尝试关闭(甚至不知道)任何被劫持的连接，例如WebSockets。

Close returns any error returned from closing the Server's underlying Listener(s).

​	`Close`方法返回从关闭Server的底层`Listener(s)`返回的任何错误。

#### (*Server) ListenAndServe 

``` go 
func (srv *Server) ListenAndServe() error
```

ListenAndServe listens on the TCP network address srv.Addr and then calls Serve to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.

​	`ListenAndServe`方法在TCP网络地址srv.Addr上监听并调用Serve来处理传入连接的请求。接受的连接将被配置为启用TCP keep-alives。

If srv.Addr is blank, ":http" is used.

​	如果`srv.Addr`为空，则使用"：http"。

ListenAndServe always returns a non-nil error. After Shutdown or Close, the returned error is ErrServerClosed.

​	`ListenAndServe`方法始终返回非零错误。在Shutdown方法或Close方法之后，返回的错误是ErrServerClosed。

#### (*Server) ListenAndServeTLS 

``` go 
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
```

ListenAndServeTLS listens on the TCP network address srv.Addr and then calls ServeTLS to handle requests on incoming TLS connections. Accepted connections are configured to enable TCP keep-alives.

​	`ListenAndServeTLS`方法在TCP网络地址`srv.Addr`上监听并调用`ServeTLS`来处理传入TLS连接的请求。接受的连接将被配置为启用TCP keep-alives。

Filenames containing a certificate and matching private key for the server must be provided if neither the Server's TLSConfig.Certificates nor TLSConfig.GetCertificate are populated. If the certificate is signed by a certificate authority, the certFile should be the concatenation of the server's certificate, any intermediates, and the CA's certificate.

​	如果Server的`TLSConfig.Certificates`或`TLSConfig.GetCertificate`都未填充，则必须提供包含服务器证书和匹配私钥的文件名。如果证书由证书颁发机构签名，则certFile应该是服务器证书、任何中间证书和CA的证书的串联。

If srv.Addr is blank, ":https" is used.

​	如果`srv.Addr`为空，则使用"：https"。

ListenAndServeTLS always returns a non-nil error. After Shutdown or Close, the returned error is ErrServerClosed.

​	`ListenAndServeTLS`方法始终返回非零错误。在Shutdown方法或Close方法之后，返回的错误是ErrServerClosed。

#### (*Server) RegisterOnShutdown  <- go1.9

``` go 
func (srv *Server) RegisterOnShutdown(f func())
```

RegisterOnShutdown registers a function to call on Shutdown. This can be used to gracefully shutdown connections that have undergone ALPN protocol upgrade or that have been hijacked. This function should start protocol-specific graceful shutdown, but should not wait for shutdown to complete.

​	`RegisterOnShutdown`方法注册一个在`Shutdown`时调用的函数。这可用于优雅地关闭已经进行了ALPN协议升级或已被劫持的连接。此函数应启动协议特定的优雅关闭，但不应等待关闭完成。

#### (*Server) Serve 

``` go 
func (srv *Server) Serve(l net.Listener) error
```

Serve accepts incoming connections on the Listener l, creating a new service goroutine for each. The service goroutines read requests and then call srv.Handler to reply to them.

​	`Serve`方法在监听器l上接受传入的连接，为每个连接创建一个新的服务goroutine。服务goroutine读取请求，然后调用`srv.Handler`回复请求。

HTTP/2 support is only enabled if the Listener returns *tls.Conn connections and they were configured with "h2" in the TLS Config.NextProtos.

​	仅当`Listener`返回`*tls.Conn`连接并且它们使用`TLS` `Config.NextProtos`配置为"h2"时才启用HTTP/2支持。

Serve always returns a non-nil error and closes l. After Shutdown or Close, the returned error is ErrServerClosed.

​	`Serve`方法总是返回一个非nil错误并关闭l。在`Shutdown`方法或`Close`方法之后，返回的错误是`ErrServerClosed`。

#### (*Server) ServeTLS  <- go1.9

``` go 
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error
```

ServeTLS accepts incoming connections on the Listener l, creating a new service goroutine for each. The service goroutines perform TLS setup and then read requests, calling srv.Handler to reply to them.

​	`ServeTLS`方法在监听器l上接受传入的连接，为每个连接创建一个新的服务goroutine。服务goroutine执行TLS设置，然后读取请求，调用`srv.Handler`回复请求。

Files containing a certificate and matching private key for the server must be provided if neither the Server's TLSConfig.Certificates nor TLSConfig.GetCertificate are populated. If the certificate is signed by a certificate authority, the certFile should be the concatenation of the server's certificate, any intermediates, and the CA's certificate.

​	如果未填充Server的`TLSConfig.Certificates`或`TLSConfig.GetCertificate`，则必须提供包含服务器证书和匹配私钥的文件。如果证书由证书颁发机构签名，则certFile应该是服务器证书、任何中间证书和CA证书的连接。

ServeTLS always returns a non-nil error. After Shutdown or Close, the returned error is ErrServerClosed.

​	`ServeTLS`方法总是返回一个非nil错误。在`Shutdown`方法或`Close`方法之后，返回的错误是ErrServerClosed。

#### (*Server) SetKeepAlivesEnabled  <- go1.3

``` go 
func (srv *Server) SetKeepAlivesEnabled(v bool)
```

SetKeepAlivesEnabled controls whether HTTP keep-alives are enabled. By default, keep-alives are always enabled. Only very resource-constrained environments or servers in the process of shutting down should disable them.

​	`SetKeepAlivesEnabled`方法控制是否启用HTTP keep-alives。默认情况下，keep-alives始终启用。只有非常资源受限的环境或正在关闭的服务器才应禁用它们。

#### (*Server) Shutdown  <- go1.8

``` go 
func (srv *Server) Shutdown(ctx context.Context) error
```

Shutdown gracefully shuts down the server without interrupting any active connections. Shutdown works by first closing all open listeners, then closing all idle connections, and then waiting indefinitely for connections to return to idle and then shut down. If the provided context expires before the shutdown is complete, Shutdown returns the context's error, otherwise it returns any error returned from closing the Server's underlying Listener(s).

​	`Shutdown`方法优雅地关闭服务器，而不会中断任何活动连接。`Shutdown`的工作原理是首先关闭所有打开的侦听器，然后关闭所有空闲连接，最后无限期地等待连接返回空闲状态，然后关闭。如果提供的上下文在关闭完成之前过期，则`Shutdown`方法返回上下文的错误，否则它返回从关闭服务器的基础侦听器(s)返回的任何错误。

When Shutdown is called, Serve, ListenAndServe, and ListenAndServeTLS immediately return ErrServerClosed. Make sure the program doesn't exit and waits instead for Shutdown to return.

​	当调用 `Shutdown`方法时，`Serve`函数、`ListenAndServe`函数和 `ListenAndServeTLS`函数立即返回 `ErrServerClosed`。确保程序不会退出，而是等待 `Shutdown` 返回。

Shutdown does not attempt to close nor wait for hijacked connections such as WebSockets. The caller of Shutdown should separately notify such long-lived connections of shutdown and wait for them to close, if desired. See RegisterOnShutdown for a way to register shutdown notification functions.

​	`Shutdown`方法不会尝试关闭也不会等待诸如 WebSocket 等被劫持的连接。如果需要，`Shutdown` 的调用者应该单独通知这些长期运行的连接关闭并等待它们关闭。有关注册关闭通知函数的方法，请参见 `RegisterOnShutdown`。

Once Shutdown has been called on a server, it may not be reused; future calls to methods such as Serve will return ErrServerClosed.

​	一旦在服务器上调用了 `Shutdown`方法，就不能重用它；诸如 `Serve` 等方法的未来调用将返回 `ErrServerClosed`。

##### Shutdown Example
``` go 
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	var srv http.Server

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

```

### type Transport 

``` go 
type Transport struct {

    // Proxy specifies a function to return a proxy for a given
	// Request. If the function returns a non-nil error, the
	// request is aborted with the provided error.
	//
	// The proxy type is determined by the URL scheme. "http",
	// "https", and "socks5" are supported. If the scheme is empty,
	// "http" is assumed.
	//
	// If Proxy is nil or returns a nil *URL, no proxy is used.
	// Proxy 指定返回给定请求的代理函数。
    // 如果函数返回非空错误，则使用提供的错误中止请求。
	//
	// 代理类型由 URL 方案确定。支持"http"、"https"和"socks5"。
    // 如果方案为空，则假定为"http"。
	//
	// 如果 Proxy 为 nil 或返回 nil *URL，则不使用代理。
	Proxy func(*Request) (*url.URL, error)

    // OnProxyConnectResponse is called when the Transport gets an HTTP response from
	// a proxy for a CONNECT request. It's called before the check for a 200 OK response.
	// If it returns an error, the request fails with that error.
	// OnProxyConnectResponse 在 Transport 获取代理的 
    // CONNECT 请求的 HTTP 响应时调用。
    // 它在检查是否为 200 OK 响应之前被调用。
    // 如果它返回错误，则请求将失败并返回该错误。

	OnProxyConnectResponse func(ctx context.Context, proxyURL *url.URL, connectReq *Request, connectRes *Response) error

    // DialContext specifies the dial function for creating unencrypted TCP connections.
	// If DialContext is nil (and the deprecated Dial below is also nil),
	// then the transport dials using package net.
	//
	// DialContext runs concurrently with calls to RoundTrip.
	// A RoundTrip call that initiates a dial may end up using
	// a connection dialed previously when the earlier connection
	// becomes idle before the later DialContext completes.
	// DialContext 指定用于创建未加密 TCP 连接的 dial 函数。
    // 如果 DialContext 为 nil(并且下面的 Dial 已弃用)，
    // 则 transport 使用 package net 进行拨号。
	//
	// DialContext 与调用 RoundTrip 并发运行。
    // 发起 dial 的 RoundTrip 调用可能会在 DialContext 
    // 完成之前变得空闲，从而使用之前拨打的连接。
	DialContext func(ctx context.Context, network, addr string) (net.Conn, error)

    // Dial specifies the dial function for creating unencrypted TCP connections.
	//
	// Dial runs concurrently with calls to RoundTrip.
	// A RoundTrip call that initiates a dial may end up using
	// a connection dialed previously when the earlier connection
	// becomes idle before the later Dial completes.
	//
	// Deprecated: Use DialContext instead, which allows the transport
	// to cancel dials as soon as they are no longer needed.
	// If both are set, DialContext takes priority.
	// Dial 指定用于创建未加密 TCP 连接的 dial 函数。
	//
	// Dial 与调用 RoundTrip 并发运行。
    // 发起 dial 的 RoundTrip 调用可能会在 Dial 完成之前变得空闲，
    // 从而使用之前拨打的连接。
	//
	// 已弃用：使用 DialContext 代替，
    // 它允许 transport 在不再需要 dial 时立即取消它们。
    // 如果两者都设置，则 DialContext 优先。
	Dial func(network, addr string) (net.Conn, error)

    // DialTLSContext specifies an optional dial function for creating
	// TLS connections for non-proxied HTTPS requests.
	//
	// If DialTLSContext is nil (and the deprecated DialTLS below is also nil),
	// DialContext and TLSClientConfig are used.
	//
	// If DialTLSContext is set, the Dial and DialContext hooks are not used for HTTPS
	// requests and the TLSClientConfig and TLSHandshakeTimeout
	// are ignored. The returned net.Conn is assumed to already be
	// past the TLS handshake.// DialTLS specifies an optional dial function for creating
	// TLS connections for non-proxied HTTPS requests.
	//
	// Deprecated: Use DialTLSContext instead, which allows the transport
	// to cancel dials as soon as they are no longer needed.
	// If both are set, DialTLSContext takes priority.
	// DialTLSContext 指定了一个可选的拨号函数，
    // 用于为非代理的 HTTPS 请求创建 TLS 连接。
	//
	// 如果 DialTLSContext 为 nil(并且已弃用的 DialTLS 也是 nil)，
    // 则使用 DialContext 和 TLSClientConfig。
	//
	// 如果设置了 DialTLSContext，
    // 则 HTTPS 请求不使用 Dial 和 DialContext 钩子，
    // 并且忽略 TLSClientConfig 和 TLSHandshakeTimeout。
    // 假定返回的 net.Conn 已经经过了 TLS 握手。
	DialTLSContext func(ctx context.Context, network, addr string) (net.Conn, error)

    // DialTLS specifies an optional dial function for creating
	// TLS connections for non-proxied HTTPS requests.
	//
	// Deprecated: Use DialTLSContext instead, which allows the transport
	// to cancel dials as soon as they are no longer needed.
	// If both are set, DialTLSContext takes priority.
	// DialTLS 指定了一个可选的拨号函数，
    // 用于为非代理的 HTTPS 请求创建 TLS 连接。
	//
	// 已弃用：使用 DialTLSContext 代替，
    // 它允许传输在不再需要时取消拨号。
    // 如果两者都设置了，DialTLSContext 优先。
	DialTLS func(network, addr string) (net.Conn, error)

    // TLSClientConfig specifies the TLS configuration to use with
	// tls.Client.
	// If nil, the default configuration is used.
	// If non-nil, HTTP/2 support may not be enabled by default.
	// TLSClientConfig 指定用于 tls.Client 的 TLS 配置。
	// 如果为 nil，则使用默认配置。
	// 如果为非 nil，则默认情况下可能未启用 HTTP/2 支持。
	TLSClientConfig *tls.Config

    // TLSHandshakeTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake. Zero means no timeout.
	// TLSHandshakeTimeout 指定等待 TLS 握手的最长时间。零表示无超时。
	TLSHandshakeTimeout time.Duration

    // DisableKeepAlives, if true, disables HTTP keep-alives and
	// will only use the connection to the server for a single
	// HTTP request.
	//
	// This is unrelated to the similarly named TCP keep-alives.
	// DisableKeepAlives 为 true 时禁用 HTTP keep-alives，
    // 并且仅使用与服务器的连接进行单个 HTTP 请求。
	//
	// 这与同名的 TCP keep-alives 无关。
	DisableKeepAlives bool

    
    // DisableCompression, if true, prevents the Transport from
	// requesting compression with an "Accept-Encoding: gzip"
	// request header when the Request contains no existing
	// Accept-Encoding value. If the Transport requests gzip on
	// its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body. However, if the user
	// explicitly requested gzip it is not automatically
	// uncompressed.
	// DisableCompression 为 true 时，
    // 阻止 Transport 在 Request 不包含任何现有的 
    // Accept-Encoding 值时，
    // 使用"Accept-Encoding：gzip"请求标头请求压缩。
    // 如果 Transport 自行请求 gzip 并获得 gzipped 响应，
    // 则 Response.Body 中的响应会被透明解码。
    // 但是，如果用户明确请求 gzip，则不会自动解压缩。
	DisableCompression bool

    // MaxIdleConns controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	// MaxIdleConns 控制所有主机上闲置(keep-alive)连接的最大数量。
    // 零表示没有限制。
	MaxIdleConns int

    // MaxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) connections to keep per-host. If zero,
	// DefaultMaxIdleConnsPerHost is used.
	// MaxIdleConnsPerHost(如果非零)控制每个主机
    // 保留的最大闲置(keep-alive)连接数。
    // 如果为零，则使用 DefaultMaxIdleConnsPerHost。
	MaxIdleConnsPerHost int

    // MaxConnsPerHost optionally limits the total number of
	// connections per host, including connections in the dialing,
	// active, and idle states. On limit violation, dials will block.
	//
	// Zero means no limit.
	// MaxConnsPerHost 可选择限制每个主机的总连接数，
    // 包括拨号、活动和空闲状态的连接。
    // 当超过限制时，拨号将阻塞。
	// 零表示没有限制。
	MaxConnsPerHost int

    // IdleConnTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	// Zero means no limit.
	// IdleConnTimeout 是空闲(keep-alive)连接保持
    // 空闲状态的最长时间。
	// 零表示没有限制。
	IdleConnTimeout time.Duration

    // ResponseHeaderTimeout, if non-zero, specifies the amount of
	// time to wait for a server's response headers after fully
	// writing the request (including its body, if any). This
	// time does not include the time to read the response body.
	// ResponseHeaderTimeout(如果非零)指定完全写入请求
    //(包括其正文，如果有)后等待服务器的响应标头的时间。
    // 此时间不包括读取响应正文的时间。
	ResponseHeaderTimeout time.Duration

    // ExpectContinueTimeout, if non-zero, specifies the amount of
	// time to wait for a server's first response headers after fully
	// writing the request headers if the request has an
	// "Expect: 100-continue" header. Zero means no timeout and
	// causes the body to be sent immediately, without
	// waiting for the server to approve.
	// This time does not include the time to send the request header.
	// ExpectContinueTimeout(如果非零)指定发送
    // 带有"Expect: 100-continue"标头的请求标头后，
    // 等待服务器的第一个响应标头的时间。
	// 零表示没有超时，会立即发送请求正文，无需等待服务器批准。
	// 此时间不包括发送请求标头的时间。
	ExpectContinueTimeout time.Duration

    // TLSNextProto specifies how the Transport switches to an
	// alternate protocol (such as HTTP/2) after a TLS ALPN
	// protocol negotiation. If Transport dials an TLS connection
	// with a non-empty protocol name and TLSNextProto contains a
	// map entry for that key (such as "h2"), then the func is
	// called with the request's authority (such as "example.com"
	// or "example.com:1234") and the TLS connection. The function
	// must return a RoundTripper that then handles the request.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	// TLSNextProto 指定如何在 TLS ALPN 协议协商后切换到
    // 替代协议(例如 HTTP/2)。
    // 如果 Transport 用非空协议名称拨号 TLS 连接，
    // 并且 TLSNextProto 包含该键(例如 "h2")的映射条目，
	// 则将使用请求的授权(例如"example.com"或"example.com:1234")
    // 和 TLS 连接调用该函数。
    // 该函数必须返回一个 RoundTripper，然后处理该请求。
	// 如果 TLSNextProto 不为 nil，则不会自动启用 HTTP/2 支持。
	TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper

    // ProxyConnectHeader optionally specifies headers to send to
	// proxies during CONNECT requests.
	// To set the header dynamically, see GetProxyConnectHeader.
	// ProxyConnectHeader 可选地指定在 CONNECT 
    // 请求期间发送到代理的标头。
	// 要动态设置标头，请参见 GetProxyConnectHeader。
	ProxyConnectHeader Header

    // GetProxyConnectHeader optionally specifies a func to return
	// headers to send to proxyURL during a CONNECT request to the
	// ip:port target.
	// If it returns an error, the Transport's RoundTrip fails with
	// that error. It can return (nil, nil) to not add headers.
	// If GetProxyConnectHeader is non-nil, ProxyConnectHeader is
	// ignored.
	// GetProxyConnectHeader 可选地指定要在针对 ip:port 
    // 目标的 CONNECT 请求期间发送到 proxyURL 的标头的函数。
	// 如果返回错误，则 Transport 的 RoundTrip 会失败并返回该错误。
    // 它可以返回 (nil, nil) 以不添加标头。
	// 如果 GetProxyConnectHeader 不是 nil，
    // 则 ProxyConnectHeader 将被忽略。
	GetProxyConnectHeader func(ctx context.Context, proxyURL *url.URL, target string) (Header, error)

    // MaxResponseHeaderBytes specifies a limit on how many
	// response bytes are allowed in the server's response
	// header.
	//
	// Zero means to use a default limit.
	// MaxResponseHeaderBytes 指定在服务器的响应标头中
    // 允许多少响应字节的限制。
	// 如果为零，则使用默认限制。
	MaxResponseHeaderBytes int64

    // WriteBufferSize specifies the size of the write buffer used
	// when writing to the transport.
	// If zero, a default (currently 4KB) is used.
	// WriteBufferSize 指定在写入传输时使用的写缓冲区的大小。
	// 如果为零，则使用默认值(当前为 4KB)。
	WriteBufferSize int

    // ReadBufferSize specifies the size of the read buffer used
	// when reading from the transport.
	// If zero, a default (currently 4KB) is used.
	// ReadBufferSize 指定在从传输中读取时使用的读缓冲区的大小。
	// 如果为零，则使用默认值(当前为 4KB)。
	ReadBufferSize int

    // ForceAttemptHTTP2 controls whether HTTP/2 is enabled when a non-zero
	// Dial, DialTLS, or DialContext func or TLSClientConfig is provided.
	// By default, use of any those fields conservatively disables HTTP/2.
	// To use a custom dialer or TLS config and still attempt HTTP/2
	// upgrades, set this to true.
	// ForceAttemptHTTP2 控制在提供了非零的 Dial、DialTLS 或 
    // DialContext 函数或 TLSClientConfig 时是否启用 HTTP/2。
	// 默认情况下，使用任何这些字段都会保守地禁用 HTTP/2。
	// 要使用自定义拨号器或 TLS 配置并仍然尝试 HTTP/2 升级，
    // 请将其设置为 true。
	ForceAttemptHTTP2 bool
	// 包含过滤或不导出的字段
}
```

Transport is an implementation of RoundTripper that supports HTTP, HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).

​	`Transport` 是一个实现了 `RoundTripper` 接口的结构体，支持 HTTP、HTTPS 和 HTTP 代理(使用 `CONNECT` 实现 HTTP 或 HTTPS 代理)。

By default, Transport caches connections for future re-use. This may leave many open connections when accessing many hosts. This behavior can be managed using Transport's CloseIdleConnections method and the MaxIdleConnsPerHost and DisableKeepAlives fields.

​	默认情况下，Transport 会缓存连接以便后续重用。这会在访问多个主机时导致许多开放的连接。可以使用 `Transport` 的 `CloseIdleConnections` 方法以及 `MaxIdleConnsPerHost` 和 `DisableKeepAlives` 字段来管理这种行为。

Transports should be reused instead of created as needed. Transports are safe for concurrent use by multiple goroutines.

​	应该重复使用 Transport 而不是按需创建它们。Transport 可以被多个 goroutine 并发使用。

A Transport is a low-level primitive for making HTTP and HTTPS requests. For high-level functionality, such as cookies and redirects, see Client.

​	`Transport` 是用于进行 HTTP 和 HTTPS 请求的低级原语。对于高级功能(如 cookie 和重定向)，请参阅 Client。

Transport uses HTTP/1.1 for HTTP URLs and either HTTP/1.1 or HTTP/2 for HTTPS URLs, depending on whether the server supports HTTP/2, and how the Transport is configured. The DefaultTransport supports HTTP/2. To explicitly enable HTTP/2 on a transport, use golang.org/x/net/http2 and call ConfigureTransport. See the package docs for more about HTTP/2.

​	`Transport` 在 HTTP URL 上使用 HTTP/1.1，在 HTTPS URL 上使用 HTTP/1.1 或 HTTP/2，具体取决于服务器是否支持 HTTP/2，以及 Transport 的配置方式。DefaultTransport 支持 HTTP/2。要在传输中显式启用 HTTP/2，请使用 golang.org/x/net/http2 并调用 `ConfigureTransport`。有关 HTTP/2 的更多信息，请参见软件包文档。

Responses with status codes in the 1xx range are either handled automatically (100 expect-continue) or ignored. The one exception is HTTP status code 101 (Switching Protocols), which is considered a terminal status and returned by RoundTrip. To see the ignored 1xx responses, use the httptrace trace package's ClientTrace.Got1xxResponse.

​	状态码为 1xx 的响应要么被自动处理(100 expect-continue)，要么被忽略。唯一的例外是 HTTP 状态码 101(切换协议)，它被视为终端状态并由 `RoundTrip` 返回。要查看被忽略的 1xx 响应，请使用 httptrace 跟踪包的 `ClientTrace.Got1xxResponse`。

Transport only retries a request upon encountering a network error if the request is idempotent and either has no body or has its Request.GetBody defined. HTTP requests are considered idempotent if they have HTTP methods GET, HEAD, OPTIONS, or TRACE; or if their Header map contains an "Idempotency-Key" or "X-Idempotency-Key" entry. If the idempotency key value is a zero-length slice, the request is treated as idempotent but the header is not sent on the wire.

​	仅当请求是幂等的并且没有主体或其 Request.GetBody 已定义时，Transport 才在遇到网络错误时重试请求。如果 HTTP 请求具有 HTTP 方法 GET、HEAD、OPTIONS 或 TRACE，或者它们的 Header 映射包含 "Idempotency-Key" 或 "X-Idempotency-Key" 条目，则被视为幂等的。如果幂等键值为零长度切片，则将该请求视为幂等，但不会将头部发送到网络。

#### (*Transport) CancelRequest <- DEPRECATED

```go
func (t *Transport) CancelRequest(req *Request)
```

CancelRequest cancels an in-flight request by closing its connection. CancelRequest should only be called after RoundTrip has returned.

Deprecated: Use Request.WithContext to create a request with a cancelable context instead. CancelRequest cannot cancel HTTP/2 requests.

#### (*Transport) Clone  <- go1.13

``` go 
func (t *Transport) Clone() *Transport
```

Clone returns a deep copy of t's exported fields.

​	`Clone`方法返回 `t` 的导出字段的深度副本。

#### (*Transport) CloseIdleConnections 

``` go 
func (t *Transport) CloseIdleConnections()
```

CloseIdleConnections closes any connections which were previously connected from previous requests but are now sitting idle in a "keep-alive" state. It does not interrupt any connections currently in use.

​	`CloseIdleConnections`方法关闭之前从先前的请求连接到的但现在处于"保持活动"状态的空闲连接。它不会中断当前正在使用的任何连接。

#### (*Transport) RegisterProtocol 

``` go 
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)
```

RegisterProtocol registers a new protocol with scheme. The Transport will pass requests using the given scheme to rt. It is rt's responsibility to simulate HTTP request semantics.

​	`RegisterProtocol`方法注册一个新的协议，使用指定的协议名称(`scheme`)和`RoundTripper`处理程序(rt)。`Transport`会将使用给定`scheme`的请求传递给`rt`。它是rt的责任来模拟HTTP请求语义。

RegisterProtocol can be used by other packages to provide implementations of protocol schemes like "ftp" or "file".

​	`RegisterProtocol`方法可以被其他包用来提供协议方案(scheme)的实现，例如"ftp"或"file"。

If rt.RoundTrip returns ErrSkipAltProtocol, the Transport will handle the RoundTrip itself for that one request, as if the protocol were not registered.

​	如果`rt.RoundTrip`返回`ErrSkipAltProtocol`，则`Transport`将自己处理该请求的`RoundTrip`，就好像未注册该协议一样。

#### (*Transport) RoundTrip 

``` go 
func (t *Transport) RoundTrip(req *Request) (*Response, error)
```

RoundTrip implements the RoundTripper interface.

​	`RoundTrip`方法实现`RoundTripper`接口。

For higher-level HTTP client support (such as handling of cookies and redirects), see Get, Post, and the Client type.

​	对于更高级的HTTP客户端支持(例如处理cookie和重定向)，请参见`Get`，`Post`和`Client`类型。

Like the RoundTripper interface, the error types returned by RoundTrip are unspecified.

​	与`RoundTripper`接口一样，RoundTrip返回的错误类型是未指定的。