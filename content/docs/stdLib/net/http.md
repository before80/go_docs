+++
title = "http"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# http

[https://pkg.go.dev/net/http@go1.20.1](https://pkg.go.dev/net/http@go1.20.1)

​	http包提供了 HTTP 客户端和服务端的实现。

​	Get、Head、Post 和 PostForm 方法可用于发起 HTTP(或 HTTPS)请求：

```
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
```

​	在使用完响应后，客户端必须关闭响应主体：

```
resp, err := http.Get("http://example.com/")
if err != nil {
	// 处理错误
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
// ...
```

​	为了控制 HTTP 客户端的标头、重定向策略和其他设置，可以创建一个 Client：

```
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

​	为了控制代理、TLS 配置、保持连接、压缩和其他设置，可以创建一个 Transport：

```
tr := &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
```

​	Clients 和 Transports 可以被多个 goroutine 并发使用，并且为了效率应该只创建一次并重复使用。

​	ListenAndServe 方法可启动一个使用给定地址和处理程序的 HTTP 服务器。处理程序通常为 nil，这意味着使用 DefaultServeMux。Handle 和 HandleFunc 方法可将处理程序添加到 DefaultServeMux：

```
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

​	通过创建自定义服务器可以获得更多对服务器行为的控制：

```
s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```

​	从 Go 1.6 开始，当使用 HTTPS 时，http 包对 HTTP/2 协议具有透明支持。必须禁用 HTTP/2 的程序可以通过将 Transport.TLSNextProto(用于客户端)或 Server.TLSNextProto(用于服务器)设置为non-nil的空映射来实现。或者，目前支持以下 GODEBUG 环境变量：

```
GODEBUG=http2client=0  # disable HTTP/2 client support
GODEBUG=http2server=0  # disable HTTP/2 server support
GODEBUG=http2debug=1   # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2   # ... even more verbose, with frame dumps
```

​	GODEBUG 变量不受 Go 的 API 兼容性承诺的保护。在禁用 HTTP/2 支持之前，请报告任何问题：https://golang.org/s/http2bug

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

​	常见的HTTP方法。

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

​	HTTP状态代码在IANA注册。请参见：[https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml)

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=857)

``` go 
const DefaultMaxHeaderBytes = 1 << 20 // 1 MB
```

​	DefaultMaxHeaderBytes是HTTP请求中标头的最大允许大小。这可以通过设置Server.MaxHeaderBytes来覆盖。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transport.go;l=58)

``` go 
const DefaultMaxIdleConnsPerHost = 2
```

​	DefaultMaxIdleConnsPerHost是Transport的MaxIdleConnsPerHost的默认值。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=934)

``` go 
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
```

​	TimeFormat是生成HTTP标头中时间时要使用的时间格式。它类似于time.RFC1123，但将GMT硬编码为时区。要格式化的时间必须在UTC中，Format才能生成正确的格式。

For parsing this time format, see ParseTime.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=513)

``` go 
const TrailerPrefix = "Trailer:"
```

​	TrailerPrefix是ResponseWriter.Header映射键的魔术前缀，如果存在，则表示该映射条目实际上是响应trailer，而不是响应标头。 ServeHTTP调用完成后，将删除前缀并将值发送到trailers中。

​	该机制仅用于不在编写标头之前未知的trailers。如果trailers的集合是固定的或在编写标头之前已知，则首选常规的Go trailers机制：

```
https://pkg.go.dev/net/http#ResponseWriter
https://pkg.go.dev/net/http#example-ResponseWriter-Trailers
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=51)

``` go 
var (
	// ErrNotSupported 表示不支持某一特性。
	//
	// 它由 ResponseController 方法返回，表示处理程序不支持该方法，
	// 并且由 Pusher 实现的 Push 方法返回，表示不支持 HTTP/2 推送功能。
	ErrNotSupported = &ProtocolError{"feature not supported"}

	// 已弃用：ErrUnexpectedTrailer 
    // 不再由 net/http 包中的任何内容返回。
	// 调用者不应将错误与此变量进行比较。
	ErrUnexpectedTrailer = &ProtocolError{"trailer header without chunked transfer encoding"}

	// 当请求的 Content-Type 不包含 "boundary" 参数时，
	// Request.MultipartReader 将返回 ErrMissingBoundary。
	ErrMissingBoundary = &ProtocolError{"no multipart boundary param in Content-Type"}

	// 当请求的 Content-Type 不是 multipart/form-data 时，
	// Request.MultipartReader 将返回 ErrNotMultipart。
	ErrNotMultipart = &ProtocolError{"request Content-Type isn't multipart/form-data"}

	// 已弃用：ErrHeaderTooLong 不再由 net/http 包中的任何内容返回。
	// 调用者不应将错误与此变量进行比较。
	ErrHeaderTooLong = &ProtocolError{"header too long"}

	// 已弃用：ErrShortBody 不再由 net/http 包中的任何内容返回。
	// 调用者不应将错误与此变量进行比较。
	ErrShortBody = &ProtocolError{"entity body too short"}

	// 已弃用：ErrMissingContentLength 
    // 不再由 net/http 包中的任何内容返回。
	// 调用者不应将错误与此变量进行比较。
	ErrMissingContentLength = &ProtocolError{"missing ContentLength in HEAD response"}
)
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=37)

``` go 
var (
	// ErrBodyNotAllowed 在 ResponseWriter.Write 调用时，
    // 当 HTTP 方法或响应状态码不允许 body 时返回。
	ErrBodyNotAllowed = errors.New("http: request method or response status code does not allow body")

	// ErrHijacked 在 ResponseWriter.Write 调用时，
    // 当使用 Hijacker 接口劫持了底层连接时返回。
    // 在劫持的连接上进行零字节写操作会返回 ErrHijacked，没有其他副作用。
	ErrHijacked = errors.New("http: connection has been hijacked")

	// ErrContentLength 在 ResponseWriter.Write 调用时，
    // 当处理程序设置了一个声明大小的 Content-Length 响应头
    // 并尝试写入比声明的更多字节时返回。

	ErrContentLength = errors.New("http: wrote more than the declared Content-Length")

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
	// ServerContextKey 是一个 context key。
    // 它可在 HTTP 处理程序中使用 Context.Value 
    // 来访问启动处理程序的服务器。相关的值的类型为 *Server。

	ServerContextKey = &contextKey{"http-server"}

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

​	DefaultClient 是默认的 Client，并被 Get、Head 和 Post 使用。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2322)

``` go 
var DefaultServeMux = &defaultServeMux
```

​	DefaultServeMux 是 Serve 使用的默认 ServeMux。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=1826)

``` go 
var ErrAbortHandler = errors.New("net/http: abort Handler")
```

​	ErrAbortHandler 是一个特殊的 panic 值，用于中止处理程序。虽然任何从 ServeHTTP 中发生的 panic 都会中止向客户端的响应，但是使用 ErrAbortHandler 发生 panic 还会阻止将堆栈跟踪记录到服务器的错误日志中。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transfer.go;l=823)

``` go 
var ErrBodyReadAfterClose = errors.New("http: invalid Read on closed Body")
```

​	当请求或响应主体关闭后继续读取主体时，返回 ErrBodyReadAfterClose。这通常发生在 HTTP 处理程序在其 ResponseWriter 上调用 WriteHeader 或 Write 后读取主体时。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3356)

``` go 
var ErrHandlerTimeout = errors.New("http: Handler timeout")
```

​	在处理超时的处理程序的 ResponseWriter Write 调用上返回 ErrHandlerTimeout。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transfer.go;l=29)

``` go 
var ErrLineTooLong = internal.ErrLineTooLong
```

​	当使用格式错误的分块编码读取请求或响应主体时返回 ErrLineTooLong。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=39)

``` go 
var ErrMissingFile = errors.New("http: no such file")
```

​	当请求中提供的文件字段名称不存在或不是文件字段时，由 FormFile 返回 ErrMissingFile。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=415)

``` go 
var ErrNoCookie = errors.New("http: named cookie not present")
```

​	当找不到 cookie 时，Request 的 Cookie 方法会返回 ErrNoCookie。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/response.go;l=131)

``` go 
var ErrNoLocation = errors.New("http: no Location header in response")
```

​	当响应中没有 Location 头时，Response 的 Location 方法会返回 ErrNoLocation。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3017)

``` go 
var ErrServerClosed = errors.New("http: Server closed")
```

​	ErrServerClosed在Server的Serve、ServeTLS、ListenAndServe和ListenAndServeTLS方法调用Shutdown或Close后返回。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transport.go;l=741)

``` go 
var ErrSkipAltProtocol = errors.New("net/http: skip alternate protocol")
```

​	ErrSkipAltProtocol是Transport.RegisterProtocol定义的标志性错误值。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=489)

``` go 
var ErrUseLastResponse = errors.New("net/http: use last response")
```

​	如果Client.CheckRedirect钩子返回ErrUseLastResponse，则可以控制如何处理重定向。 如果返回此值，则不会发送下一个请求，并使用其正文未关闭的最近响应返回。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/http.go;l=104)

``` go 
var NoBody = noBody{}
```

​	NoBody是一个没有字节的io.ReadCloser。Read始终返回EOF，Close始终返回nil。 它可以在传出的客户端请求中使用，以显式地表示请求没有字节。 但是，另一种方法是将Request.Body设置为nil。

## 函数

#### func [CanonicalHeaderKey](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=240) 

``` go 
func CanonicalHeaderKey(s string) string
```

​	CanonicalHeaderKey返回标头键s的规范格式。规范化将第一个字母和任何连字符后面的字母转换为大写字母；其余字母转换为小写字母。例如，"accept-encoding"的规范键是"Accept-Encoding"。如果s包含空格或无效的标头字段字节，则返回不带修改的s。

#### func [DetectContentType](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/sniff.go;l=21) 

``` go 
func DetectContentType(data []byte) string
```

​	DetectContentType实现在https://mimesniff.spec.whatwg.org/上描述的算法，以确定给定数据的Content-Type。它最多考虑前512个字节的数据。DetectContentType始终返回有效的MIME类型：如果无法确定更具体的类型，则返回"application/octet-stream"。

#### func [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2131) 

``` go 
func Error(w ResponseWriter, error string, code int)
```

​	Error使用指定的错误消息和HTTP代码回复请求。它不会以其他方式结束请求；调用者应确保不会对w进行进一步的写入。错误消息应为纯文本。

#### func [Handle](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2559) 

``` go 
func Handle(pattern string, handler Handler)
```

​	Handle 在 DefaultServeMux 中为给定的 pattern 注册 handler。ServeMux 的文档解释了如何匹配 pattern。

##### Handle Example
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

#### func [HandleFunc](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2564) 

``` go 
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

​	HandleFunc 在 DefaultServeMux 中为给定的 pattern 注册 handler function。ServeMux 的文档解释了如何匹配 pattern。

##### HandleFunc Example
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

#### func [ListenAndServe](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3240) 

``` go 
func ListenAndServe(addr string, handler Handler) error
```

​	ListenAndServe 监听 TCP 网络地址 addr，并在传入连接上调用 Serve 以处理请求。已接受的连接已配置为启用 TCP keep-alives。

​	通常 handler 是 nil，在这种情况下会使用 DefaultServeMux。

​	ListenAndServe 总是返回非 nil 错误。

##### ListenAndServe Example
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

#### func [ListenAndServeTLS](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3250) 

``` go 
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```

​	ListenAndServeTLS 的行为与 ListenAndServe 相同，除了它预期 HTTPS 连接。另外，必须提供包含服务器证书和匹配私钥的文件。如果证书由证书颁发机构签署，则 certFile 应该是服务器证书、任何中间文件以及 CA 的证书的连接。

##### ListenAndServeTLS Example
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

#### func [MaxBytesReader](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=1141) 

``` go 
func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser
```

​	MaxBytesReader 类似于 io.LimitReader，但旨在限制传入请求主体的大小。与 io.LimitReader 不同，MaxBytesReader 的结果是一个 ReadCloser，在超出限制的情况下返回类型为 *MaxBytesError 的非 nil 错误，并在调用其 Close 方法时关闭底层读取器。

​	MaxBytesReader 防止客户端意外或恶意发送大型请求并浪费服务器资源。如果可能，它会告诉 ResponseWriter 在达到限制后关闭连接。

#### func [NotFound](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2139) 

``` go 
func NotFound(w ResponseWriter, r *Request)
```

​	NotFound 函数返回一个 HTTP 404 错误响应。

#### func [ParseHTTPVersion](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=794) 

``` go 
func ParseHTTPVersion(vers string) (major, minor int, ok bool)
```

​	ParseHTTPVersion 函数按照 [RFC 7230 第 2.6 节](https://rfc-editor.org/rfc/rfc7230.html#section-2.6)解析 HTTP 版本字符串。例如 "HTTP/1.0" 会返回 (1, 0, true)。请注意，不带次要版本号的字符串，例如 "HTTP/2"，是无效的。

#### func [ParseTime](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=129)  <- go1.1

``` go 
func ParseTime(text string) (t time.Time, err error)
```

​	ParseTime 函数解析时间头(如 Date 头)，尝试 HTTP/1.1 所允许的三种格式：TimeFormat、time.RFC850 和 time.ANSIC。

#### func [ProxyFromEnvironment](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transport.go;l=447) 

``` go 
func ProxyFromEnvironment(req *Request) (*url.URL, error)
```

​	ProxyFromEnvironment 函数返回给定请求使用的代理 URL，该 URL 由环境变量 HTTP_PROXY、HTTPS_PROXY 和 NO_PROXY(或它们的小写形式)指示。请求使用与其方案匹配的环境变量中的代理，除非被 NO_PROXY 排除。

​	环境变量的值可以是完整的 URL，也可以是 "host[:port]" 形式，在这种情况下会假定为 "http" 方案。支持 "http"、"https" 和 "socks5" 方案。如果值是其他形式，则返回错误。

​	如果环境中没有定义代理，或者请求不应使用代理，则返回 nil URL 和 nil 错误。

​	作为特例，如果 req.URL.Host 是 "localhost"(带或不带端口号)，则会返回 nil URL 和 nil 错误。

#### func [ProxyURL](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transport.go;l=453) 

``` go 
func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error)
```

​	ProxyURL 函数返回一个代理函数(供 Transport 使用)，该函数始终返回相同的 URL。

#### func [Redirect](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2182) 

``` go 
func Redirect(w ResponseWriter, r *Request, url string, code int)
```

​	Redirect 函数将请求重定向到 url，url 可以是相对于请求路径的路径。

​	提供的 code 应该在 3xx 范围内，通常为 StatusMovedPermanently、StatusFound 或 StatusSeeOther。

​	如果尚未设置 Content-Type 标头，则 Redirect 函数会将其设置为 "text/html; charset=utf-8" 并写入小的 HTML 主体。设置 Content-Type 标头为任何值(包括 nil)都会禁用该行为。

#### func [Serve](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2579) 

``` go 
func Serve(l net.Listener, handler Handler) error
```

​	Serve 函数在监听器 l 上接受传入的 HTTP 连接，为每个连接创建一个新的服务协程。服务协程读取请求，然后调用 handler 进行响应。

​	handler 通常为 nil，此时会使用 DefaultServeMux。

​	只有当 Listener 返回 `*tls.Conn` 连接且它们在 TLS Config.NextProtos 中配置为 "h2" 时，才启用 HTTP/2 支持。

​	Serve 函数始终返回非 nil 错误。

#### func [ServeContent](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/fs.go;l=194) 

``` go 
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
```

​	ServeContent 使用 io.ReadSeeker 中提供的内容回复请求。ServeContent 相对于 io.Copy 的主要好处在于正确处理范围请求、设置 MIME 类型以及处理 If-Match、If-Unmodified-Since、If-None-Match、If-Modified-Since 和 If-Range 请求。

​	如果响应的 Content-Type 头未设置，ServeContent 首先尝试从 name 的文件扩展名推断类型，如果失败，则返回到读取内容的第一个块并将其传递给 DetectContentType。否则，不使用 name；特别地，它可以为空，并且永远不会在响应中发送。

​	如果 modtime 不是零时间或 Unix 纪元，则 ServeContent 在响应中包括一个 Last-Modified 头。如果请求包括一个 If-Modified-Since 头，则 ServeContent 使用 modtime 来决定是否需要发送内容。

​	内容的 Seek 方法必须工作：ServeContent 使用 seek 到内容的末尾以确定其大小。

​	如果调用方按 [RFC 7232，第 2.3 节](https://rfc-editor.org/rfc/rfc7232.html#section-2.3)格式化了 w 的 ETag 头，则 ServeContent 使用它来处理使用 If-Match、If-None-Match 或 If-Range 的请求。

​	请注意，`*os.File` 实现了 io.ReadSeeker 接口。

#### func [ServeFile](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/fs.go;l=730) 

``` go 
func ServeFile(w ResponseWriter, r *Request, name string)
```

​	ServeFile 用指定的文件或目录的内容回复请求。

​	如果提供的文件或目录名是相对路径，则相对于当前目录进行解释，可能升到父目录。如果提供的名称由用户输入构造，则在调用 ServeFile 之前应进行清理。

​	作为一项预防措施，ServeFile 将拒绝 r.URL.Path 包含 ".." 路径元素的请求；这可保护免受可能会不安全地对 r.URL.Path 使用 filepath.Join 而未对其进行清理的调用者的侵害，并使用 filepath.Join 结果作为名称参数。

​	作为另一个特殊情况，ServeFile 将任何以 "/index.html" 结尾的 r.URL.Path 的请求重定向到相同的路径，不包括最后的 "index.html"。为避免此类重定向，请修改路径或使用 ServeContent。

​	除了这两个特殊情况外，ServeFile 不使用 r.URL.Path 来选择要提供的文件或目录；只使用名称参数中提供的文件或目录。

#### func [ServeTLS](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2596)  <- go1.9

``` go 
func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error
```

​	ServeTLS 函数在监听器 l 上接受传入的 HTTPS 连接，为每个连接创建一个新的服务协程。服务协程读取请求，然后调用处理程序来回复请求。

​	处理程序通常为 nil，此时将使用 DefaultServeMux。

​	另外，必须提供包含服务器证书和匹配的私钥的文件。如果证书由证书颁发机构签名，则 certFile 应该是服务器证书、任何中间证书和 CA 证书的连接。

​	ServeTLS 函数总是返回非 nil 的错误。

#### func [SetCookie](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookie.go;l=169) 

``` go 
func SetCookie(w ResponseWriter, cookie *Cookie)
```

​	SetCookie 函数将 Set-Cookie 头添加到提供的 ResponseWriter 的头中。提供的 cookie 必须有一个有效的名称。无效的 cookie 可能会被静默丢弃。

#### func [StatusText](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/status.go;l=81) 

``` go 
func StatusText(code int) string
```

​	StatusText 函数返回 HTTP 状态码的文本。如果状态码未知，则返回空字符串。

## 类型

### type [Client](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=58) 

``` go 
type Client struct {
	// Transport 指定单个 HTTP 请求的执行机制。
	// 如果为 nil，则使用 DefaultTransport。
	Transport RoundTripper

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

	// Jar 指定 cookie 存储。
	//
	// Jar 用于在每个出站请求中插入相关的 cookies，
    // 并使用每个入站 Response 的 cookie 值进行更新。
    // Jar 用于处理客户端跟随的每个重定向。
	//
	// 如果 Jar 为 nil，则仅在请求上显式设置了 cookie 时才发送它们。
	Jar CookieJar

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

​	一个 Client 是一个 HTTP 客户端。它的零值(DefaultClient)是一个可用的客户端，它使用 DefaultTransport。

​	Client 的 Transport 通常具有内部状态(缓存的 TCP 连接)，因此应该重复使用 Client，而不是按需创建。Client 可以被多个 goroutine 并发使用。

​	Client 比 RoundTripper(如 Transport)更高级，还处理诸如 cookie 和重定向之类的 HTTP 细节。

​	在遵循重定向时，Client 将转发在初始请求上设置的所有标头，但除了以下情况：

- 将敏感标头(例如"Authorization"、"WWW-Authenticate"和"Cookie")转发到不受信任的目标时。在重定向到与初始域不是子域匹配或精确匹配的域时，将忽略这些标头。例如，从"foo.com"重定向到"foo.com"或"sub.foo.com"将转发敏感标头，但是重定向到"bar.com"将不会。

-  使用非空 cookie Jar 转发"Cookie"标头时。由于每个重定向可能会更改 cookie jar 的状态，重定向可能会更改在初始请求中设置的 cookie。在转发"Cookie"标头时，任何已更改的 cookie 都将被省略，预期 Jar 将使用更新后的值插入这些已更改的 cookie(假设原点匹配)。如果 Jar 为 nil，则初始 cookie 会不加更改地转发。

#### (*Client) [CloseIdleConnections](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=935)  <- go1.12

``` go 
func (c *Client) CloseIdleConnections()
```

​	CloseIdleConnections 关闭 Transport 上任何连接，这些连接先前从先前的请求连接，但现在处于"keep-alive"状态。它不会中断当前正在使用的任何连接。

​	如果 Client 的 Transport 没有 CloseIdleConnections 方法，则此方法不执行任何操作。

#### (*Client) [Do](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=581) 

``` go 
func (c *Client) Do(req *Request) (*Response, error)
```

​	Do 发送 HTTP 请求并返回 HTTP 响应，遵循客户端配置的策略(例如重定向、cookie、身份验证)。

​	如果由客户端策略(例如 CheckRedirect)或无法进行 HTTP 通信(例如网络连接问题)导致错误，则返回错误。非 2xx 状态代码不会导致错误。

​	如果返回的错误是 nil，那么 Response 将包含一个非 nil 的 Body，用户应该在读取完后关闭它。如果 Body 没有读取到 EOF 并且关闭，Client 的底层 RoundTripper(通常是 Transport)可能无法重用持久的 TCP 连接，用于后续的 "keep-alive" 请求。

​	如果请求的 Body 非 nil，则在底层的 Transport 上即使发生错误，也会关闭它。

​	在出现错误时，可以忽略任何 Response。当 CheckRedirect 失败时，非 nil 的 Response 和非 nil 的 error 仅会出现一次。即使出现这种情况，返回的 Response.Body 也已经关闭。

​	通常会使用 Get、Post 或 PostForm 而不是 Do。

​	如果服务器回复重定向，则客户端首先使用 CheckRedirect 函数来确定是否应该跟随重定向。如果被允许，则 301、302 或 303 重定向会导致后续请求使用 HTTP 方法 GET(或 HEAD，如果原始请求是 HEAD)，而不使用 Body。307 或 308 重定向会保留原始的 HTTP 方法和 Body，前提是定义了 Request.GetBody 函数。NewRequest 函数会自动为常见的标准库 Body 类型设置 GetBody。

​	任何返回的错误都将是 `*url.Error` 类型。url.Error 值的 Timeout 方法将在请求超时时报告 true。

#### (*Client) [Get](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=475) 

``` go 
func (c *Client) Get(url string) (resp *Response, err error)
```

​	Get 向指定的 URL 发送一个 GET 请求。如果响应是以下重定向代码之一，则 Get 在调用 Client 的 CheckRedirect 函数后跟随重定向：

```
301 (Moved Permanently)(永久移动)
302 (Found)(找到)
303 (See Other)(参见其他)
307 (Temporary Redirect)(临时重定向)
308 (Permanent Redirect)(永久重定向)
```

​	如果 Client 的 CheckRedirect 函数失败或存在 HTTP 协议错误，则返回错误。非 2xx 响应不会导致错误。任何返回的错误都将是 `*url.Error` 类型。url.Error 值的 Timeout 方法将在请求超时时报告 true。

​	当 err 为 nil 时，resp 总是包含一个非 nil 的 resp.Body。在读取完它后，调用者应该关闭 resp.Body。

​	要使用自定义标头发出请求，请使用 NewRequest 和 Client.Do。

​	要使用指定的 context.Context 发出请求，请使用 NewRequestWithContext 和 Client.Do。

#### (*Client) [Head](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=920) 

``` go 
func (c *Client) Head(url string) (resp *Response, err error)
```

​	Head 向指定的 URL 发出一个 HEAD 请求。如果响应是以下重定向代码之一，则在调用 Client 的 CheckRedirect 函数后，Head 跟随重定向：

```
301 (Moved Permanently)(永久移动)
302 (Found)(找到)
303 (See Other)(参见其他)
307 (Temporary Redirect)(临时重定向)
308 (Permanent Redirect)(永久重定向)
```

​	使用NewRequestWithContext和Client.Do可以指定context.Context进行请求。

#### (*Client) [Post](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=843) 

``` go 
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
```

​	Post向指定的URL发出POST请求。

​	调用者在完成读取后应关闭resp.Body。

​	如果提供的body是io.Closer，则在请求后关闭它。

​	要设置自定义标头，请使用NewRequest和Client.Do。

​	要使用指定的context.Context进行请求，请使用NewRequestWithContext和Client.Do。

​	有关如何处理重定向的详细信息，请参阅Client.Do方法文档。

#### (*Client) [PostForm](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=886) 

``` go 
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
```

​	PostForm将数据的键和值作为请求正文进行URL编码，并向指定的URL发出POST请求。

​	Content-Type标头设置为application/x-www-form-urlencoded。要设置其他标头，请使用NewRequest和Client.Do。

​	当err为nil时，resp始终包含非nil resp.Body。完成读取后，调用者应关闭resp.Body。

​	有关如何处理重定向的详细信息，请参阅Client.Do方法文档。

​	要使用指定的context.Context进行请求，请使用NewRequestWithContext和Client.Do。

### type [ConnState](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2859)  <- go1.3

``` go 
type ConnState int
```

​	ConnState表示与服务器的客户端连接状态。它由可选的Server.ConnState钩子使用。

``` go 
const (
	// StateNew 表示一个新的连接，该连接预计会立即发送请求。
    // 连接从此状态开始，然后转换为 StateActive 或 StateClosed。
	StateNew ConnState = iota

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

	// StateIdle 表示已完成请求处理并处于保持活动状态的连接，
    // 等待新的请求。
    // 连接从 StateIdle 转换为 StateActive 或 StateClosed。
	StateIdle

	// StateHijacked 表示被劫持的连接。
    // 这是一个终止状态。它不会转换为 StateClosed。
	StateHijacked

	// StateClosed 表示已关闭的连接。
    // 这是一个终止状态。
    // 被劫持的连接不会转换为 StateClosed。
	StateClosed
)
```

#### (ConnState) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2905)  <- go1.3

``` go 
func (c ConnState) String() string
```

### type [Cookie](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookie.go;l=23) 

``` go 
type Cookie struct {
	Name  string
	Value string

	Path       string    // 可选
	Domain     string    // 可选
	Expires    time.Time // 可选
	RawExpires string    // 仅用于读取cookie

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

​	Cookie 表示 HTTP cookie，其在 HTTP 响应的 Set-Cookie 标头或 HTTP 请求的 Cookie 标头中发送。

​	详见[https://tools.ietf.org/html/rfc6265](https://tools.ietf.org/html/rfc6265)。

#### (*Cookie) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookie.go;l=179) 

``` go 
func (c *Cookie) String() string
```

​	String 返回 cookie 的序列化字符串，以用于 Cookie 标头(如果仅设置了 Name 和 Value)或 Set-Cookie 响应标头(如果设置了其他字段)。如果 c 为 nil 或 c.Name 无效，则返回空字符串。

#### (*Cookie) [Valid](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookie.go;l=243)  <- go1.18

``` go 
func (c *Cookie) Valid() error
```

​	Valid 函数用于判断 cookie 是否有效。

### type [CookieJar](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/jar.go;l=17) 

``` go 
type CookieJar interface {
	// SetCookies 处理收到的回复中的 cookie。
    // 根据 jar 的策略和实现，它可能会选择保存 cookie。
	SetCookies(u *url.URL, cookies []*Cookie)

	// Cookies 返回指定 URL 的请求中应发送的 cookie。
    // 实现必须遵守 RFC 6265 等标准 cookie 使用限制。
	Cookies(u *url.URL) []*Cookie
}
```

​	CookieJar 管理 HTTP 请求中的 cookie 的存储和使用。

​	CookieJar 的实现必须支持多 goroutine 并发使用。

​	net/http/cookiejar 包提供了 CookieJar 的实现。

### type [Dir](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/fs.go;l=44) 

``` go 
type Dir string
```

​	Dir 类型使用本地文件系统来实现 FileSystem 接口，且仅限于特定的目录树。

​	尽管 FileSystem.Open 方法接收以 / 分隔的路径，但 Dir 的 string 值是本地文件系统上的文件名，而不是 URL，因此它由 filepath.Separator 分隔，不一定是 `/`。

​	请注意，Dir 可能会暴露敏感文件和目录。Dir 将遵循指向目录树之外的符号链接，如果从用户可以创建任意符号链接的目录提供服务，这可能会特别危险。Dir 还将允许访问以句点开头的文件和目录，这可能会暴露像 .git 这样的敏感目录或像 `.htpasswd` 这样的敏感文件。要排除以句点开头的文件，请从服务器删除这些文件/目录，或者创建自定义 FileSystem 实现。

​	一个空的 Dir 被视为 "`.`"。

#### (Dir) [Open](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/fs.go;l=72) 

``` go 
func (d Dir) Open(name string) (File, error)
```

​	Open 实现 FileSystem 接口，使用 os.Open 打开文件进行读取，根据目录 d 来确定其路径。

### type [File](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/fs.go;l=104) 

``` go 
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]fs.FileInfo, error)
	Stat() (fs.FileInfo, error)
}
```

​	该接口的 Open 方法返回一个 File，可由 FileServer 实现进行服务。

​	该接口的方法应与 `*os.File` 上的方法表现相同。

### type [FileSystem](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/fs.go;l=96) 

``` go 
type FileSystem interface {
	Open(name string) (File, error)
}
```

​	FileSystem 实现对命名文件集合的访问。文件路径中的元素使用斜杠('/'，U+002F)字符分隔，而不管主机操作系统惯例如何。请参见 FileServer 函数，将 FileSystem 转换为处理程序。

​	该接口早于 fs.FS 接口，可以使用 fs.FS 代替：FS 适配器函数将 fs.FS 转换为 FileSystem。

#### func [FS](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/fs.go;l=837)  <- go1.16

``` go 
func FS(fsys fs.FS) FileSystem
```

​	FS 将 fsys 转换为 FileSystem 实现，供 FileServer 和 NewFileTransport 使用。由 fsys 提供的文件必须实现 io.Seeker。

### type [Flusher](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=173) 

``` go 
type Flusher interface {
	// Flush sends any buffered data to the client.
	Flush()
}
```

​	Flusher 接口由 ResponseWriter 实现，允许 HTTP 处理程序将缓冲的数据刷新到客户端。

​	默认的 HTTP/1.x 和 HTTP/2 ResponseWriter 实现支持 Flusher，但 ResponseWriter 包装器可能不支持。处理程序应始终在运行时测试此功能。

​	请注意，即使对于支持 Flush 的 ResponseWriter，如果客户端通过 HTTP 代理连接，则缓冲的数据可能要等到响应完成后才能到达客户端。

### type [Handler](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=86) 

``` go 
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

​	Handler 响应 HTTP 请求。

​	ServeHTTP 应将响应标头和数据写入 ResponseWriter，然后返回。返回信号表示请求已完成；在完成 ServeHTTP 调用之后或与之同时使用 ResponseWriter 或从 Request.Body 中读取是无效的。

​	根据HTTP客户端软件、HTTP协议版本以及客户端和Go服务器之间的任何中介，可能无法在写入ResponseWriter之后从Request.Body中读取。谨慎的处理程序应该先读取Request.Body，然后再回复。

​	除了读取body外，处理程序不应修改提供的Request。

​	如果ServeHTTP恐慌，服务器(ServeHTTP的调用者)假定恐慌的影响仅限于活动请求。它会恢复panic，将堆栈跟踪记录到服务器错误日志，并关闭网络连接或发送HTTP/2 RST_STREAM，具体取决于HTTP协议。为了中止处理程序，以便客户端看到中断的响应但服务器不记录错误，请使用值为ErrAbortHandler的panic。

#### func [AllowQuerySemicolons](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2950)  <- go1.17

``` go 
func AllowQuerySemicolons(h Handler) Handler
```

​	AllowQuerySemicolons返回一个处理程序，它通过将URL查询中的任何未转义分号转换为"&"并调用处理程序h来为请求提供服务。

​	这恢复了分割查询参数的分号和"&"的Go 1.17之前的行为。(请参阅golang.org/issue/25192)。请注意，此行为与许多代理的行为不匹配，不匹配可能导致安全问题。

​	在调用Request.ParseForm之前，应调用AllowQuerySemicolons。

#### func [FileServer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/fs.go;l=856) 

``` go 
func FileServer(root FileSystem) Handler
```

​	FileServer返回一个处理程序，它使用以root为根的文件系统的内容为HTTP请求提供服务。

​	作为一个特殊情况，返回的文件服务器将以"/index.html"结尾的任何请求重定向到相同的路径，而不包括最后的"index.html"。

​	要使用操作系统的文件系统实现，请使用http.Dir：

```
http.Handle("/", http.FileServer(http.Dir("/tmp")))
```

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

#### func [MaxBytesHandler](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3638)  <- go1.18

``` go 
func MaxBytesHandler(h Handler, n int64) Handler
```

​	MaxBytesHandler返回一个处理程序，该处理程序使用其ResponseWriter和Request.Body包装了一个MaxBytesReader。

#### func [NotFoundHandler](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2143) 

``` go 
func NotFoundHandler() Handler
```

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

#### func [RedirectHandler](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2267) 

``` go 
func RedirectHandler(url string, code int) Handler
```

​	RedirectHandler 函数返回一个请求处理程序，该程序会使用给定的状态码将收到的每个请求重定向到给定的URL。

​	提供的状态码应该在 3xx 范围内，通常为 StatusMovedPermanently、StatusFound 或 StatusSeeOther。

#### func [StripPrefix](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2151) 

``` go 
func StripPrefix(prefix string, h Handler) Handler
```

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

#### func [TimeoutHandler](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3346) 

``` go 
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
```

​	TimeoutHandler 函数返回一个运行限时的处理程序。

​	新的处理程序调用 h.ServeHTTP 来处理每个请求，但如果一个调用的运行时间超过其时间限制，则处理程序会使用一个 HTTP 503 Service Unavailable 错误并在其主体中给出给定的消息进行回复。 (如果 msg 为空，则将发送适当的默认消息。)在此类超时之后，h 对其 ResponseWriter 的写入将返回 ErrHandlerTimeout。

​	TimeoutHandler 支持 Pusher 接口，但不支持 Hijacker 或 Flusher 接口。

### type [HandlerFunc](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2118) 

``` go 
type HandlerFunc func(ResponseWriter, *Request)
```

​	HandlerFunc 类型是一个适配器，允许使用普通函数作为 HTTP 处理程序。如果 f 是具有适当签名的函数，则 HandlerFunc(f) 是调用 f 的处理程序。

#### (HandlerFunc) [ServeHTTP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2121) 

``` go 
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

​	ServeHTTP 调用 f(w, r)。

### type [Header](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=24) 

``` go 
type Header map[string][]string
```

​	Header 表示 HTTP 头中的键值对。

​	键应该是规范化形式，即由 CanonicalHeaderKey函数返回的形式。

#### (Header) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=30) 

``` go 
func (h Header) Add(key, value string)
```

​	Add方法向 Header 中添加一个键值对，它会将值附加到与键相关联的任何现有值的末尾。键的大小写不敏感，因此会通过 CanonicalHeaderKey 进行规范化。

#### (Header) [Clone](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=94)  <- go1.13

``` go 
func (h Header) Clone() Header
```

​	Clone方法返回 h 的一个副本，如果 h 为 nil，则返回 nil。

#### (Header) [Del](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=80) 

``` go 
func (h Header) Del(key string)
```

​	Del方法删除与 key 关联的值。键的大小写不敏感，因此会通过 CanonicalHeaderKey函数进行规范化。

#### (Header) [Get](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=49) 

``` go 
func (h Header) Get(key string) string
```

​	Get方法获取与给定键关联的第一个值。如果没有与该键关联的值，则 Get 返回 ""。键的大小写不敏感，因此使用 textproto.CanonicalMIMEHeaderKey 来规范化提供的键。Get 假定所有键都以规范形式存储。要使用非规范键，请直接访问映射。

#### (Header) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=39) 

``` go 
func (h Header) Set(key, value string)
```

​	Set方法将与 key 关联的 header 条目设置为单个元素值。它将替换与 key 关联的任何现有值。键的大小写不敏感，因此会通过 textproto.CanonicalMIMEHeaderKey 进行规范化。要使用非规范键，请直接分配到映射中。

#### (Header) [Values](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=58)  <- go1.14

``` go 
func (h Header) Values(key string) []string
```

​	Values方法返回与给定键关联的所有值。键的大小写不敏感，因此使用 textproto.CanonicalMIMEHeaderKey 来规范化提供的键。要使用非规范键，请直接访问映射。返回的切片不是副本。

#### (Header) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=85) 

``` go 
func (h Header) Write(w io.Writer) error
```

​	Write 方法以 wire format 写入 Header。

#### (Header) [WriteSubset](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/header.go;l=192) 

``` go 
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```

​	WriteSubset 方法以 wire format 写入 Header。如果 exclude 不为 nil，则不会写入其中 exclude[key] == true 的键。写入之前不会对键进行规范化。

### type [Hijacker](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=185) 

``` go 
type Hijacker interface {
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

​	Hijacker接口由允许HTTP处理程序接管连接的ResponseWriters实现。

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

### type [MaxBytesError](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=1149)  <- go1.19

``` go 
type MaxBytesError struct {
	Limit int64
}
```

​	MaxBytesReader读取的字节数超过其读取限制时返回MaxBytesError。

#### (*MaxBytesError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=1153)  <- go1.19

``` go 
func (e *MaxBytesError) Error() string
```

### type [PushOptions](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/http.go;l=119)  <- go1.8

``` go 
type PushOptions struct {
	// Method指定承诺请求的HTTP方法。
    // 如果设置，则必须是"GET"或"HEAD"。空表示"GET"。
	Method string

	// Header指定额外的承诺请求标头。
    // 这不能包括HTTP/2伪标头字段，如"：path"和"：scheme"，
    // 它们将自动添加。
	Header Header
}
```

PushOptions describes options for Pusher.Push.

### type [Pusher](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/http.go;l=133)  <- go1.8

``` go 
type Pusher interface {
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

​	Pusher是ResponseWriters实现的接口，用于支持HTTP/2服务器推送。有关更多背景信息，请参见https://tools.ietf.org/html/rfc7540#section-8.2。

### type [Request](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=106) 

``` go 
type Request struct {
	// Method 指定HTTP方法(GET，POST，PUT等)。
	// 对于客户端请求，空字符串表示GET。
	//
	// Go的HTTP客户端不支持使用CONNECT方法发送请求。
    // 有关详细信息，请参阅传输的文档。
	Method string

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

	// 传入服务器请求的协议版本。
	//
	// 对于客户端请求，将忽略这些字段。
    // HTTP客户端代码始终使用HTTP/1.1或HTTP/2。
    // 有关详细信息，请参阅传输文档。
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0

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

	// GetBody 定义了一个可选的函数，
    // 用于返回 Body 的新副本。
    // 当重定向需要多次读取主体时，它用于客户端请求。
    // 仍然需要设置 Body 才能使用 GetBody。
	//
	// 对于服务器请求，它没有使用。
	GetBody func() (io.ReadCloser, error)

	// ContentLength 记录相关内容的长度。
    // 值-1表示长度未知。值>= 0表示可以从Body读取给定的字节数。
	//
	// 对于客户端请求，具有非零Body的值为0也被视为未知。
	ContentLength int64

	// TransferEncoding列出了从最外层到最内层的传输编码。
    // 空列表表示"identity"编码。
    // 在发送和接收请求时，可以通常忽略TransferEncoding。
    // 当需要时，块编码会自动添加和删除。
	TransferEncoding []string

	// Close指示是否在回复此请求后(对于服务器)
    // 或在发送此请求并读取其响应后(对于客户端)关闭连接。
	//
	// 对于服务器请求，HTTP服务器会自动处理此操作，处理程序不需要此字段。
	//
	// 对于客户端请求，
    // 设置此字段会阻止在向相同主机的请求之间重新使用TCP连接，
    // 就像Transport.DisableKeepAlives已设置一样。
	Close bool

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

	// Form 包含已解析的表单数据，
    // 包括 URL 字段的查询参数以及 PATCH、POST 或 PUT 表单数据。
	// 只有在调用 ParseForm 之后才能使用此字段。
    // HTTP 客户端会忽略 Form 并使用 Body。
	Form url.Values

	// PostForm 包含来自 PATCH、POST 或 PUT 请求体参数的解析表单数据。
	// 只有在调用 ParseForm 之后才能使用此字段。
    // HTTP 客户端会忽略 PostForm 并使用 Body。
	PostForm url.Values

	// MultipartForm 是已解析的多部分表单，包括文件上传。
	// 只有在调用 ParseMultipartForm 之后才能使用此字段。
    // HTTP 客户端会忽略 MultipartForm 并使用 Body。
	MultipartForm *multipart.Form

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

	// RemoteAddr 允许 HTTP 服务器和其他软件记录发送请求的网络地址，
    // 通常用于记录日志。
	// ReadRequest 不会填充此字段，也没有定义的格式。
    // 该包中的 HTTP 服务器在调用处理程序之前将 RemoteAddr 
    // 设置为 "IP:port" 地址。
	// HTTP 客户端会忽略此字段。
	RemoteAddr string

	// RequestURI 是客户端发送到服务器的请求行
    //(RFC 7230，第 3.1.1 节)的未修改请求目标。
	// 通常应该使用 URL 字段而不是此字段。
    // 在 HTTP 客户端请求中设置此字段是错误的。
	RequestURI string

	// TLS 允许 HTTP 服务器和其他软件记录接收请求的 TLS 连接的信息。
    // 该字段不由 ReadRequest 填充。
	// 在调用处理程序之前，
    // 此包中的 HTTP 服务器在启用 TLS 的连接上设置字段；
    // 否则它将保留字段为空。
	// HTTP 客户端忽略此字段。
	TLS *tls.ConnectionState

	// Cancel 是一个可选通道，其关闭表示应将客户端请求视为已取消。
    // 并非所有 RoundTripper 的实现都支持 Cancel。
    //
    // 对于服务器请求，此字段不适用。
    //
    // 已弃用：应使用 NewRequestWithContext 
    // 将请求的上下文设置为上下文，而不是设置请求的 Cancel 字段。
    // 如果请求的 Cancel 字段和上下文都已设置，则未定义是否尊重 Cancel。
	Cancel <-chan struct{}

	// Response 是重定向响应，导致创建此请求。
    // 此字段仅在客户端重定向期间填充。
	Response *Response
	// 包含过滤或未导出的字段
}
```

​	一个Request代表一个由服务器接收或由客户端发送的HTTP请求。

​	客户端和服务器使用中字段语义略有不同。除了下面字段的注释外，请参阅Request.Write和RoundTripper的文档。

#### func [NewRequest](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=839) 

``` go 
func NewRequest(method, url string, body io.Reader) (*Request, error)
```

​	NewRequest使用context.Background包装NewRequestWithContext。

#### func [NewRequestWithContext](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=865)  <- go1.13

``` go 
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error)
```

​	NewRequestWithContext给定方法、URL和可选的body，返回一个新的Request。

​	NewRequestWithContext返回一个适用于使用Client.Do或Transport.RoundTrip的Request。要为测试服务器处理程序创建一个请求，请使用net/http/httptest包中的NewRequest函数、使用ReadRequest或手动更新Request字段。对于传出的客户端请求，上下文控制请求及其响应的整个生命周期：获取连接、发送请求以及读取响应标头和正文。请参阅Request类型的文档，了解入站和出站请求字段之间的差异。

​	如果body的类型为*bytes.Buffer、*bytes.Reader或*strings.Reader，则返回的请求的ContentLength将设置为其确切值(而不是-1)，GetBody将被填充(因此307和308重定向可以重放body)，如果ContentLength为0，则Body将设置为NoBody。

#### func [ReadRequest](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=1024) 

``` go 
func ReadRequest(b *bufio.Reader) (*Request, error)
```

​	ReadRequest从b中读取并解析一个传入的请求。

​	ReadRequest是一个低级函数，只应用于专用应用程序；大多数代码应该使用Server来读取请求并通过Handler接口处理它们。ReadRequest仅支持HTTP/1.x请求。对于HTTP/2，请使用golang.org/x/net/http2。

#### (*Request) [AddCookie](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=437) 

``` go 
func (r *Request) AddCookie(c *Cookie)
```

​	AddCookie向请求中添加一个cookie。根据[RFC 6265第5.4节](https://rfc-editor.org/rfc/rfc6265.html#section-5.4)，AddCookie不会添加超过一个Cookie头字段，这意味着所有cookie(如果有)都将写入同一行，由分号分隔。AddCookie仅对c的名称和值进行清理，不对已经存在于请求中的Cookie头进行清理。

#### (*Request) [BasicAuth](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=949)  <- go1.4

``` go 
func (r *Request) BasicAuth() (username, password string, ok bool)
```

​	BasicAuth返回请求头中提供的用户名和密码，如果请求使用HTTP基本身份验证。请参阅[RFC 2617第2节](https://rfc-editor.org/rfc/rfc2617.html#section-2)。

#### (*Request) [Clone](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=372)  <- go1.13

``` go 
func (r *Request) Clone(ctx context.Context) *Request
```

​	Clone方法返回r的深层副本，其上下文已更改为ctx。提供的ctx必须非nil。

​	对于出站客户端请求，上下文控制请求及其响应的整个生命周期：获取连接，发送请求和读取响应标头和主体。

#### (*Request) [Context](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=340)  <- go1.7

``` go 
func (r *Request) Context() context.Context
```

​	Context方法返回请求的上下文。要更改上下文，请使用Clone方法或WithContext方法。

​	返回的上下文始终非零；默认为后台上下文。

​	对于输出客户端请求，上下文控制取消。

​	对于传入的服务器请求，当客户端连接关闭，请求被取消(使用HTTP/2)或ServeHTTP方法返回时，上下文被取消。

#### (*Request) [Cookie](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=421) 

``` go 
func (r *Request) Cookie(name string) (*Cookie, error)
```

​	Cookie方法返回请求中提供的指定cookie或ErrNoCookie(如果没有找到)。如果有多个cookie与给定名称匹配，则仅返回一个cookie。

#### (*Request) [Cookies](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=410) 

``` go 
func (r *Request) Cookies() []*Cookie
```

​	Cookies方法解析并返回发送请求的HTTP cookie。

#### (*Request) [FormFile](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=1397) 

``` go 
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
```

​	FormFile方法返回所提供的表单键的第一个文件。如果需要，FormFile方法调用ParseMultipartForm方法和ParseForm方法。

#### (*Request) [FormValue](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=1370) 

``` go 
func (r *Request) FormValue(key string) string
```

​	FormValue方法返回查询的指定组件的第一个值。POST和PUT请求体参数优先于URL查询字符串值。FormValue如果需要，调用ParseMultipartForm方法和ParseForm方法，并忽略这些函数返回的任何错误。如果键不存在，则FormValue方法返回空字符串。要访问同一键的多个值，请调用ParseForm，然后直接检查Request.Form方法。

#### (*Request) [MultipartReader](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=470) 

``` go 
func (r *Request) MultipartReader() (*multipart.Reader, error)
```

​	MultipartReader方法如果这是multipart/form-data或multipart/mixed POST请求，则返回MIME多部分读取器；否则返回nil和错误。使用此方法而不是ParseMultipartForm方法处理请求正文作为流。

#### (*Request) [ParseForm](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=1282) 

``` go 
func (r *Request) ParseForm() error
```

​	ParseForm方法解析请求中的表单，并更新`r.Form`和`r.PostForm`。

​	对于所有请求，ParseForm方法解析URL中的原始查询，并更新`r.Form`。

​	对于POST、PUT和PATCH请求，该方法还读取请求体，将其解析为表单，并将结果放入`r.PostForm`和`r.Form`中。请求体参数优先于URL查询字符串值在`r.Form`中。

​	如果请求体大小还没有被`MaxBytesReader`限制，则最大为10MB。

​	对于其他HTTP方法或`Content-Type`不为`application/x-www-form-urlencoded`的情况，不读取请求体，`r.PostForm`初始化为非零但为空的值。

​	ParseMultipartForm方法会自动调用ParseForm方法。ParseForm方法是幂等的。

#### (*Request) [ParseMultipartForm](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=1325) 

``` go 
func (r *Request) ParseMultipartForm(maxMemory int64) error
```

​	ParseMultipartForm方法将请求体解析为multipart/form-data。整个请求体被解析，并且它的文件部分的最多maxMemory字节数在内存中存储，其余部分在临时文件中存储。ParseMultipartForm方法在必要时调用ParseForm方法。如果ParseForm方法返回错误，则ParseMultipartForm方法返回该错误，但也会继续解析请求体。调用一次ParseMultipartForm方法之后，随后的调用没有任何效果。

#### (*Request) [PostFormValue](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=1385)  <- go1.1

``` go 
func (r *Request) PostFormValue(key string) string
```

​	PostFormValue方法返回POST、PATCH或PUT请求体中命名组件的第一个值。忽略URL查询参数。PostFormValue方法在必要时调用ParseMultipartForm和ParseForm方法，并忽略这些函数返回的任何错误。如果key不存在，则PostFormValue方法返回空字符串。

#### (*Request) [ProtoAtLeast](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=399) 

``` go 
func (r *Request) ProtoAtLeast(major, minor int) bool
```

​	ProtoAtLeast方法报告请求中使用的HTTP协议是否至少为major.minor。

#### (*Request) [Referer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=454) 

``` go 
func (r *Request) Referer() string
```

​	Referer方法返回引用的URL(如果在请求中发送了)。

​	Referer方法在请求本身中就被错误地拼写为Referer，这是HTTP早期的一个错误。这个值也可以从Header映射中获取，如`Header["Referer"]`；将其作为方法可用的好处是编译器可以诊断使用替代(正确的英文)拼写`req.Referrer()`的程序，但无法诊断使用`Header["Referrer"]`的程序。

#### (*Request) [SetBasicAuth](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=988) 

``` go 
func (r *Request) SetBasicAuth(username, password string)
```

​	SetBasicAuth方法将请求的Authorization头设置为使用HTTP基本身份验证，提供用户名和密码。

​	使用HTTP基本身份验证时，提供的用户名和密码未加密。通常只应在HTTPS请求中使用它。

​	用户名不能包含冒号。某些协议可能会对预转义用户名和密码有额外的要求。例如，当与OAuth2一起使用时，必须首先使用url.QueryEscape对两个参数进行URL编码。

#### (*Request) [UserAgent](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=405) 

``` go 
func (r *Request) UserAgent() string
```

​	UserAgent方法返回客户端的User-Agent，如果在请求中发送。

#### (*Request) [WithContext](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=356)  <- go1.7

``` go 
func (r *Request) WithContext(ctx context.Context) *Request
```

​	WithContext方法返回r的浅层副本，并将其上下文更改为ctx。提供的ctx必须非nil。

​	对于出站客户端请求，上下文控制请求及其响应的整个生命周期：获取连接、发送请求和读取响应头和主体。

​	要使用上下文创建新请求，请使用NewRequestWithContext函数。要使用新上下文对请求进行深层复制，请使用Request.Clone方法。

#### (*Request) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=534) 

``` go 
func (r *Request) Write(w io.Writer) error
```

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

​	如果Body存在，Content-Length小于等于0且TransferEncoding未设置为"identity"，Write将"Transfer-Encoding: chunked"添加到头部。Body在发送后被关闭。

#### (*Request) [WriteProxy](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/request.go;l=544) 

``` go 
func (r *Request) WriteProxy(w io.Writer) error
```

​	WriteProxy方法类似于 Write 方法，但是将请求写成 HTTP 代理所期望的格式。特别是，WriteProxy 使用绝对 URI 写入请求的初始 Request-URI 行，根据 [RFC 7230 第 5.3 节](https://rfc-editor.org/rfc/rfc7230.html)，包括方案和主机。无论哪种情况，WriteProxy 还会使用 r.Host 或 r.URL.Host 写入 Host 标头。

### type [Response](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/response.go;l=35) 

``` go 
type Response struct {
	Status     string // Status 表示响应状态，例如 "200 OK"。
	StatusCode int    // StatusCode 表示响应状态码，例如 200
	Proto      string // Proto 表示协议，例如 "HTTP/1.0"。
	ProtoMajor int    // ProtoMajor 表示协议的主要版本，例如 1。
	ProtoMinor int    // ProtoMinor 表示协议的次要版本，例如 0。

	// Header 将头键映射到值。
    // 如果响应包含相同的键，则它们可能会连接，用逗号分隔。
    //(RFC 7230，第3.2.2节要求多个标题在语义上等同于逗号分隔的序列。)
    // 当 Header 值被此结构中的其他字段
    //(例如 ContentLength、TransferEncoding、Trailer)重复时，
    // 字段值是权威的。
	// 映射中的键已规范化(请参阅 CanonicalHeaderKey)。
	Header Header

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

	// ContentLength 记录相关内容的长度。
    // 值 -1 表示长度未知。
    // 除非 Request.Method 是 "HEAD"，
    // 否则值 >= 0 表示可以从 Body 读取给定数量的字节。
	ContentLength int64

	// 包含最外层到最内层的传输编码。
    // 如果值为 nil，则表示使用"identity"编码。
	TransferEncoding []string

	// Close 记录了该响应头指示在读取 Body 后是否关闭连接。
    // 该值是对客户端的建议：
	// 无论是 ReadResponse 还是 Response.Write 都不会关闭连接。
	Close bool

	// Uncompressed 报告该响应是否被压缩，
    // 但是被 http 包解压缩。当为 true 时，
	// 从 Body 读取会返回解压缩后的内容，
    // 而不是实际从服务器接收到的压缩内容，ContentLength 被设置为 -1，
	// 并且从响应头中删除 "Content-Length" 和 
    // "Content-Encoding" 字段。
	// 要获取来自服务器的原始响应，
    // 请将 Transport.DisableCompression 设置为 true。
	Uncompressed bool

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

	// Request 是用于获取该响应的请求。
    // Request 的 Body 为 nil(已被使用)。
	// 这仅用于 Client 请求。
	Request *Request

	// TLS 包含有关收到响应的 TLS 连接的信息。
    // 对于未加密的响应，它为 nil。
	// 该指针在响应之间共享，不应更改。
	TLS *tls.ConnectionState
}
```

​	Response表示HTTP请求的响应。

​	一旦接收到响应头，客户端和传输机制就会从服务器返回响应。随着读取Body字段的增长，响应体会被按需流式传输。

#### func [Get](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=448) 

``` go 
func Get(url string) (resp *Response, err error)
```

​	Get方法向指定的URL发出GET请求。如果响应是以下重定向代码之一，则Get会遵循重定向，最多重定向10次：

```
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
308 (Permanent Redirect)
```

​	如果重定向过多或存在HTTP协议错误，则返回错误。非2xx响应不会引起错误。任何返回的错误都将是*url.Error类型。如果请求超时，则url.Error值的Timeout方法将返回true。

​	当err为nil时，resp始终包含一个非nil resp.Body。调用者在读取完毕后应该关闭resp.Body。

​	Get方法是DefaultClient.Get的包装。

​	要使用自定义标题进行请求，请使用NewRequest和DefaultClient.Do。

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

#### func [Head](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=904) 

``` go 
func Head(url string) (resp *Response, err error)
```

​	Head方法向指定的URL发出HEAD请求。如果响应是以下重定向代码之一，则Head方法会遵循重定向，最多重定向10次：

```
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
308 (Permanent Redirect)
```

​	Head方法是DefaultClient.Head的包装函数。

​	要使用指定的context.Context发出请求，请使用NewRequestWithContext函数和DefaultClient.Do方法。

#### func [Post](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=825) 

``` go 
func Post(url, contentType string, body io.Reader) (resp *Response, err error)
```

​	Post函数向指定的URL发出POST请求。

​	调用者在读取完resp.Body后应该关闭它。

​	如果提供的body实现了io.Closer接口，则在请求之后会关闭它。

​	Post函数是DefaultClient.Post的包装函数。

​	要设置自定义头，请使用NewRequest函数和DefaultClient.Do方法。

​	有关重定向处理方式的详细信息，请参见Client.Do方法文档。

​	要使用指定的context.Context发出请求，请使用NewRequestWithContext函数和DefaultClient.Do方法。

#### func [PostForm](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=868) 

``` go 
func PostForm(url string, data url.Values) (resp *Response, err error)
```

​	PostForm函数向指定的URL发出POST请求，使用data的键和值作为请求体进行URL编码。

​	Content-Type标头设置为application/x-www-form-urlencoded。要设置其他标头，请使用NewRequest函数和DefaultClient.Do方法。

​	当err为nil时，resp始终包含一个非空的resp.Body。在读取完毕后，调用者应该关闭resp.Body。

​	PostForm函数是DefaultClient.PostForm的包装函数。

​	有关重定向处理方式的详细信息，请参见Client.Do方法文档。

​	要使用指定的context.Context发出请求，请使用NewRequestWithContext和DefaultClient.Do。

#### func [ReadResponse](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/response.go;l=154) 

``` go 
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
```

​	ReadResponse函数从r中读取并返回HTTP响应。 req参数可选地指定与此Response对应的Request。如果为nil，则假定为GET请求。客户端必须在完成读取resp.Body后调用resp.Body.Close。在该调用之后，客户端可以检查resp.Trailer以查找包含在响应trailer中的键/值对。

#### (*Response) [Cookies](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/response.go;l=125) 

``` go 
func (r *Response) Cookies() []*Cookie
```

​	Cookies方法解析并返回在Set-Cookie头中设置的Cookie。

#### (*Response) [Location](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/response.go;l=137) 

``` go 
func (r *Response) Location() (*url.URL, error)
```

​	Location方法返回响应头"Location"字段的URL，如果存在的话。相对URL会根据响应的请求进行解析。如果没有Location头，会返回ErrNoLocation。

#### (*Response) [ProtoAtLeast](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/response.go;l=224) 

``` go 
func (r *Response) ProtoAtLeast(major, minor int) bool
```

​	ProtoAtLeast方法报告响应中使用的HTTP协议是否至少为major.minor。

#### (*Response) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/response.go;l=245) 

``` go 
func (r *Response) Write(w io.Writer) error
```

​	Write方法以HTTP/1.x服务器响应格式将r写入w，包括状态行、头部、正文和可选的尾随部分。

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

​	响应主体在发送后将被关闭。

### type [ResponseController](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/responsecontroller.go;l=17)  <- go1.20

``` go 
type ResponseController struct {
	// contains filtered or unexported fields
}
```

​	ResponseController用于在HTTP处理程序中控制响应。

​	在Handler.ServeHTTP方法返回之后，不得再使用ResponseController。

#### func [NewResponseController](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/responsecontroller.go;l=37)  <- go1.20

``` go 
func NewResponseController(rw ResponseWriter) *ResponseController
```

​	NewResponseController函数为请求创建一个ResponseController。

​	ResponseWriter应该是传递给Handler.ServeHTTP方法的原始值，或者具有返回原始ResponseWriter的Unwrap方法。

​	如果ResponseWriter实现了以下任何方法，则ResponseController将根据需要调用它们：

```
Flush()
FlushError() error // 替代Flush，返回错误
Hijack() (net.Conn, *bufio.ReadWriter, error)
SetReadDeadline(deadline time.Time) error
SetWriteDeadline(deadline time.Time) error
```

​	如果ResponseWriter不支持某个方法，则ResponseController返回与ErrNotSupported相匹配的错误。

#### (*ResponseController) [Flush](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/responsecontroller.go;l=46)  <- go1.20

``` go 
func (c *ResponseController) Flush() error
```

​	Flush方法将缓冲数据刷新到客户端。

#### (*ResponseController) [Hijack](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/responsecontroller.go;l=65)  <- go1.20

``` go 
func (c *ResponseController) Hijack() (net.Conn, *bufio.ReadWriter, error)
```

​	Hijack方法允许调用者接管连接。有关详细信息，请参见Hijacker接口。

#### (*ResponseController) [SetReadDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/responsecontroller.go;l=84)  <- go1.20

``` go 
func (c *ResponseController) SetReadDeadline(deadline time.Time) error
```

​	SetReadDeadline方法设置读取整个请求(包括正文)的截止日期。超过截止日期后从请求正文中读取会返回错误。零值表示没有截止日期。

​	在超过期限后设置读取期限将不会延长它。

#### (*ResponseController) [SetWriteDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/responsecontroller.go;l=104)  <- go1.20

``` go 
func (c *ResponseController) SetWriteDeadline(deadline time.Time) error
```

​	SetWriteDeadline方法为写入响应设置截止日期。在截止日期之后向响应正文写入不会阻塞，但如果数据已被缓冲，则可能成功。零值表示没有截止日期。

​	在超过期限后设置写入期限将不会延长它。

### type [ResponseWriter](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=95) 

``` go 
type ResponseWriter interface {
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

​	ResponseWriter接口由HTTP处理程序用于构建HTTP响应。

​	在Handler.ServeHTTP方法返回之后，不得再使用ResponseWriter。

##### Example

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

### type [RoundTripper](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/client.go;l=117) 

``` go 
type RoundTripper interface {
	// RoundTrip执行单个HTTP事务，为提供的Request返回一个Response。
	//
	// RoundTrip不应该试图解释响应。
    // 特别是，无论响应的HTTP状态代码如何，
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

​	RoundTripper是一个接口，表示执行单个HTTP事务的能力，获得给定请求的响应。

​	RoundTripper必须对多个goroutine进行并发使用的安全。

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

​	DefaultTransport是Transport的默认实现，并由DefaultClient使用。它根据需要建立网络连接并将其缓存以供后续调用重用。它按照环境变量HTTP_PROXY、HTTPS_PROXY和NO_PROXY(或其小写版本)的指示使用HTTP代理。

#### func [NewFileTransport](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/filetransport.go;l=30) 

``` go 
func NewFileTransport(fs FileSystem) RoundTripper
```

​	NewFileTransport函数返回一个新的RoundTripper，为所提供的FileSystem提供服务。返回的RoundTripper忽略其传入请求中的URL主机，以及请求的大多数其他属性。

​	NewFileTransport函数的典型用例是向Transport注册"file"协议，例如：

```
t := &http.Transport{}
t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
c := &http.Client{Transport: t}
res, err := c.Get("file:///etc/passwd")
...
```

### type [SameSite](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookie.go;l=49)  <- go1.11

``` go 
type SameSite int
```

​	SameSite 允许服务器定义一个 cookie 属性，使浏览器无法将该 cookie 与跨站请求一起发送。主要目的是减少跨源信息泄露的风险，并提供一定的保护，以防止跨站请求伪造攻击。

详见 https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00。

``` go 
const (
	SameSiteDefaultMode SameSite = iota + 1
	SameSiteLaxMode
	SameSiteStrictMode
	SameSiteNoneMode
)
```

### type [ServeMux](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2306) 

``` go 
type ServeMux struct {
	// 包含过滤或未公开的字段
}
```

​	ServeMux 是一个 HTTP 请求多路复用器。它将每个传入请求的 URL 与已注册模式列表进行匹配，并调用与 URL 最匹配的模式的处理程序。

​	模式命名了固定的根路径，如 "/favicon.ico"，或根子树，如 "/images/"(请注意尾随斜杠)。较长的模式优先于较短的模式，因此，如果为 "/images/" 和 "/images/thumbnails/" 都注册了处理程序，则后者处理程序将用于以 "/images/thumbnails/" 开头的路径，并且前者将接收任何 "/images/" 子树中的其他路径的请求。

​	请注意，由于以斜杠结尾的模式命名了一个根子树，因此模式 "/" 匹配所有未被其他已注册模式匹配的路径，而不仅仅是 Path == "/" 的 URL。

​	如果已经注册了子树，并收到了以其末尾没有斜杠的根子树命名的请求，ServeMux 会将该请求重定向到根子树(添加尾随斜杠)。此行为可以通过单独为没有尾随斜杠的路径注册来覆盖。例如，注册 "/images/" 会导致 ServeMux 将对 "/images" 的请求重定向到 "/images/"，除非 "/images" 已经单独注册。

​	模式可以可选地以主机名开头，将匹配限制为仅在该主机的 URL。特定于主机的模式优先于通用模式，因此处理程序可能会注册两个模式 "/codesearch" 和 "codesearch.google.com/"，而不会同时接管 "http://www.google.com/" 的请求。

​	ServeMux 还负责对 URL 请求路径和 Host 标头进行清理，删除端口号，并将包含 . 或 .. 元素或重复斜杠的任何请求重定向到等效且更清晰的 URL。

#### func [NewServeMux](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2319) 

``` go 
func NewServeMux() *ServeMux
```

​	NewServeMux函数分配并返回一个新的 ServeMux。

#### (*ServeMux) [Handle](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2505) 

``` go 
func (mux *ServeMux) Handle(pattern string, handler Handler)
```

​	Handle方法为给定的模式注册处理程序。如果已存在处理程序，则 Handle 方法会引发 panic。

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

#### (*ServeMux) [HandleFunc](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2549) 

``` go 
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

​	HandleFunc方法为给定的模式注册处理函数。

#### (*ServeMux) [Handler](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2436)  <- go1.1

``` go 
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
```

​	Handler方法根据 r.Method、r.Host 和 r.URL.Path 来获取给定请求使用的处理程序。它始终返回一个非 nil 的处理程序。如果路径不在其规范形式中，则处理程序将是一个内部生成的处理程序，用于重定向到规范路径。如果主机包含端口，则匹配处理程序时将忽略该端口。

​	对于 CONNECT 请求，路径和主机都不会被修改。

​	Handler方法还返回与请求匹配的已注册模式，或在生成重定向的情况下，将在遵循重定向后匹配的模式。

​	如果没有已注册的处理程序适用于请求，则 Handler方法返回一个"页面未找到"的处理程序和一个空模式。

#### (*ServeMux) [ServeHTTP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2491) 

``` go 
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

​	ServeHTTP方法将请求分派到模式最接近请求 URL 的处理程序。

### type [Server](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2603) 

``` go 
type Server struct {
	// Addr 可选地指定服务器监听的 TCP 地址，
    // 格式为 "host:port"。如果为空，则使用 ":http"(端口80)。
	// 服务名称在 RFC 6335 中定义，
    // 并由 IANA 分配。有关地址格式的详细信息，请参见 net.Dial。
	Addr string

    // 处理程序，如果为nil，则使用http.DefaultServeMux
	Handler Handler 

	// 如果为 true，则禁用通用 OPTIONS 处理程序，
    // 并将 "OPTIONS *" 请求传递给处理程序，
    // 否则响应为 200 OK 并带有 Content-Length：0。
	DisableGeneralOptionsHandler bool

	// TLSConfig 可选地提供用于 ServeTLS 和 
    // ListenAndServeTLS 的 TLS 配置。
    // 请注意，此值将由 ServeTLS 和 ListenAndServeTLS 克隆，
    // 因此无法使用 tls.Config.SetSessionTicketKeys 
    // 等方法修改配置。要使用 SetSessionTicketKeys，
    // 请改用具有 TLS 监听器的 Server.Serve。
	TLSConfig *tls.Config

	// ReadTimeout 是读取整个请求(包括正文)的最长时间。
    // 零值或负值表示没有超时。
	//
	// 由于 ReadTimeout 不允许处理程序对每个请求正文的
    // 可接受截止时间或上传速率做出每个请求的决策，
    // 因此大多数用户将更喜欢使用 ReadHeaderTimeout。
    // 使用这两个值是有效的。
	ReadTimeout time.Duration

	// ReadHeaderTimeout 是允许读取请求标头的时间量。
    // 读取标头后，连接的读取截止时间将重置，
    // 处理程序可以决定什么被认为是请求主体的过慢的响应时间。
    // 如果 ReadHeaderTimeout 为零，
    // 则使用 ReadTimeout 的值。如果两者都为零，则没有超时。
	ReadHeaderTimeout time.Duration

	// WriteTimeout 是在超时写入响应之前的最长持续时间。
    // 每次读取新请求的标头时都会重置它。
    // 与 ReadTimeout 一样，它不允许处理程序对每个请求做出决策。
    // 零值或负值表示没有超时。
	WriteTimeout time.Duration

	// IdleTimeout 是在启用保持连接的情况下等待下一个请求的最长时间。
    // 如果 IdleTimeout 为零，则使用 ReadTimeout 的值。
    // 如果两者都为零，则没有超时时间。
	IdleTimeout time.Duration

	// MaxHeaderBytes 控制服务器在解析请求标头的键和值
    //(包括请求行)时读取的最大字节数。它不限制请求主体的大小。
    // 如果为零，则使用 DefaultMaxHeaderBytes。
	MaxHeaderBytes int

	// TLSNextProto 可选地指定一个函数，
    // 以在 ALPN 协议升级发生时接管提供的 TLS 连接所有权。
    // 映射键是协议协商的协议名称。
    // Handler 参数应用于处理 HTTP 请求，
    // 并将初始化请求的 TLS 和 RemoteAddr(如果尚未设置)。
    // 该函数返回时连接将自动关闭。
    // 如果 TLSNextProto 不为 nil，则不会自动启用 HTTP/2 支持。
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

	// ConnState 指定一个可选的回调函数，
    // 当客户端连接更改状态时调用。
    // 有关详细信息，请参见 ConnState 类型和相关常量。
	ConnState func(net.Conn, ConnState)

	// ErrorLog 指定一个可选的记录器，
    // 用于记录接受连接时的错误、处理程序的意外行为以及
    // 基础 FileSystem 的错误。
    // 如果为 nil，则通过 log 包的标准记录器进行记录。
	ErrorLog *log.Logger

	// BaseContext 可选地指定一个函数，
    // 该函数返回此服务器上的传入请求的基本上下文。
    // 提供的 Listener 是即将开始接受请求的特定 Listener。
    // 如果 BaseContext 为 nil，
    // 则默认为 context.Background()。
    // 如果非 nil，则必须返回非 nil 上下文。
	BaseContext func(net.Listener) context.Context

	// ConnContext 可选地指定一个函数，
    // 该函数修改用于新连接 c 的上下文。
    // 提供的 ctx 派生自基本上下文并具有 ServerContextKey 值。
	ConnContext func(ctx context.Context, c net.Conn) context.Context
	// 包含已过滤或未公开的字段
}
```

​	Server定义了运行HTTP服务器的参数。Server的零值是有效的配置。

#### (*Server) [Close](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2722)  <- go1.8

``` go 
func (srv *Server) Close() error
```

​	Close方法会立即关闭所有活动的net.Listeners和状态为StateNew、StateActive或StateIdle的任何连接。为了优雅的关闭，请使用Shutdown方法。

​	Close方法不会尝试关闭(甚至不知道)任何被劫持的连接，例如WebSockets。

​	Close方法返回从关闭Server的底层Listener(s)返回的任何错误。

#### (*Server) [ListenAndServe](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2976) 

``` go 
func (srv *Server) ListenAndServe() error
```

​	ListenAndServe方法在TCP网络地址srv.Addr上监听并调用Serve来处理传入连接的请求。接受的连接将被配置为启用TCP keep-alives。

​	如果srv.Addr为空，则使用"：http"。

​	ListenAndServe方法始终返回非零错误。在Shutdown方法或Close方法之后，返回的错误是ErrServerClosed。

#### (*Server) [ListenAndServeTLS](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3270) 

``` go 
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
```

​	ListenAndServeTLS方法在TCP网络地址srv.Addr上监听并调用ServeTLS来处理传入TLS连接的请求。接受的连接将被配置为启用TCP keep-alives。

​	如果Server的TLSConfig.Certificates或TLSConfig.GetCertificate都未填充，则必须提供包含服务器证书和匹配私钥的文件名。如果证书由证书颁发机构签名，则certFile应该是服务器证书、任何中间证书和CA的证书的串联。

​	如果srv.Addr为空，则使用"：https"。

​	ListenAndServeTLS方法始终返回非零错误。在Shutdown方法或Close方法之后，返回的错误是ErrServerClosed。

#### (*Server) [RegisterOnShutdown](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2815)  <- go1.9

``` go 
func (srv *Server) RegisterOnShutdown(f func())
```

​	RegisterOnShutdown方法注册一个在Shutdown时调用的函数。这可用于优雅地关闭已经进行了ALPN协议升级或已被劫持的连接。此函数应启动协议特定的优雅关闭，但不应等待关闭完成。

#### (*Server) [Serve](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3029) 

``` go 
func (srv *Server) Serve(l net.Listener) error
```

​	Serve方法在监听器l上接受传入的连接，为每个连接创建一个新的服务goroutine。服务goroutine读取请求，然后调用srv.Handler回复请求。

​	仅当Listener返回*tls.Conn连接并且它们使用TLS Config.NextProtos配置为"h2"时才启用HTTP/2支持。

​	Serve方法总是返回一个非nil错误并关闭l。在Shutdown方法或Close方法之后，返回的错误是ErrServerClosed。

#### (*Server) [ServeTLS](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3106)  <- go1.9

``` go 
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error
```

​	ServeTLS方法在监听器l上接受传入的连接，为每个连接创建一个新的服务goroutine。服务goroutine执行TLS设置，然后读取请求，调用srv.Handler回复请求。

​	如果未填充Server的TLSConfig.Certificates或TLSConfig.GetCertificate，则必须提供包含服务器证书和匹配私钥的文件。如果证书由证书颁发机构签名，则certFile应该是服务器证书、任何中间证书和CA证书的连接。

​	ServeTLS方法总是返回一个非nil错误。在Shutdown方法或Close方法之后，返回的错误是ErrServerClosed。

#### (*Server) [SetKeepAlivesEnabled](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=3200)  <- go1.3

``` go 
func (srv *Server) SetKeepAlivesEnabled(v bool)
```

​	SetKeepAlivesEnabled方法控制是否启用HTTP keep-alives。默认情况下，keep-alives始终启用。只有非常资源受限的环境或正在关闭的服务器才应禁用它们。

#### (*Server) [Shutdown](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/server.go;l=2772)  <- go1.8

``` go 
func (srv *Server) Shutdown(ctx context.Context) error
```

​	Shutdown方法优雅地关闭服务器，而不会中断任何活动连接。Shutdown的工作原理是首先关闭所有打开的侦听器，然后关闭所有空闲连接，最后无限期地等待连接返回空闲状态，然后关闭。如果提供的上下文在关闭完成之前过期，则Shutdown方法返回上下文的错误，否则它返回从关闭服务器的基础侦听器(s)返回的任何错误。

​	当调用 Shutdown方法时，Serve函数、ListenAndServe函数和 ListenAndServeTLS函数立即返回 ErrServerClosed。确保程序不会退出，而是等待 Shutdown 返回。

​	Shutdown方法不会尝试关闭也不会等待诸如 WebSocket 等被劫持的连接。如果需要，Shutdown 的调用者应该单独通知这些长期运行的连接关闭并等待它们关闭。有关注册关闭通知函数的方法，请参见 RegisterOnShutdown。

​	一旦在服务器上调用了 Shutdown方法，就不能重用它；诸如 Serve 等方法的未来调用将返回 ErrServerClosed。

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

### type [Transport](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transport.go;l=95) 

``` go 
type Transport struct {

	// Proxy 指定返回给定请求的代理函数。
    // 如果函数返回非空错误，则使用提供的错误中止请求。
	//
	// 代理类型由 URL 方案确定。支持"http"、"https"和"socks5"。
    // 如果方案为空，则假定为"http"。
	//
	// 如果 Proxy 为 nil 或返回 nil *URL，则不使用代理。
	Proxy func(*Request) (*url.URL, error)

	// OnProxyConnectResponse 在 Transport 获取代理的 
    // CONNECT 请求的 HTTP 响应时调用。
    // 它在检查是否为 200 OK 响应之前被调用。
    // 如果它返回错误，则请求将失败并返回该错误。

	OnProxyConnectResponse func(ctx context.Context, proxyURL *url.URL, connectReq *Request, connectRes *Response) error

	// DialContext 指定用于创建未加密 TCP 连接的 dial 函数。
    // 如果 DialContext 为 nil(并且下面的 Dial 已弃用)，
    // 则 transport 使用 package net 进行拨号。
	//
	// DialContext 与调用 RoundTrip 并发运行。
    // 发起 dial 的 RoundTrip 调用可能会在 DialContext 
    // 完成之前变得空闲，从而使用之前拨打的连接。
	DialContext func(ctx context.Context, network, addr string) (net.Conn, error)

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

	// DialTLS 指定了一个可选的拨号函数，
    // 用于为非代理的 HTTPS 请求创建 TLS 连接。
	//
	// 已弃用：使用 DialTLSContext 代替，
    // 它允许传输在不再需要时取消拨号。
    // 如果两者都设置了，DialTLSContext 优先。
	DialTLS func(network, addr string) (net.Conn, error)

	// TLSClientConfig 指定用于 tls.Client 的 TLS 配置。
	// 如果为 nil，则使用默认配置。
	// 如果为非 nil，则默认情况下可能未启用 HTTP/2 支持。
	TLSClientConfig *tls.Config

	// TLSHandshakeTimeout 指定等待 TLS 握手的最长时间。零表示无超时。
	TLSHandshakeTimeout time.Duration

	// DisableKeepAlives 为 true 时禁用 HTTP keep-alives，
    // 并且仅使用与服务器的连接进行单个 HTTP 请求。
	//
	// 这与同名的 TCP keep-alives 无关。
	DisableKeepAlives bool

	// DisableCompression 为 true 时，
    // 阻止 Transport 在 Request 不包含任何现有的 
    // Accept-Encoding 值时，
    // 使用"Accept-Encoding：gzip"请求标头请求压缩。
    // 如果 Transport 自行请求 gzip 并获得 gzipped 响应，
    // 则 Response.Body 中的响应会被透明解码。
    // 但是，如果用户明确请求 gzip，则不会自动解压缩。
	DisableCompression bool

	// MaxIdleConns 控制所有主机上闲置(keep-alive)连接的最大数量。
    // 零表示没有限制。
	MaxIdleConns int

	// MaxIdleConnsPerHost(如果非零)控制每个主机
    // 保留的最大闲置(keep-alive)连接数。
    // 如果为零，则使用 DefaultMaxIdleConnsPerHost。
	MaxIdleConnsPerHost int

	// MaxConnsPerHost 可选择限制每个主机的总连接数，
    // 包括拨号、活动和空闲状态的连接。
    // 当超过限制时，拨号将阻塞。
	// 零表示没有限制。
	MaxConnsPerHost int

	// IdleConnTimeout 是空闲(keep-alive)连接保持
    // 空闲状态的最长时间。
	// 零表示没有限制。
	IdleConnTimeout time.Duration

	// ResponseHeaderTimeout(如果非零)指定完全写入请求
    //(包括其正文，如果有)后等待服务器的响应标头的时间。
    // 此时间不包括读取响应正文的时间。
	ResponseHeaderTimeout time.Duration

	// ExpectContinueTimeout(如果非零)指定发送
    // 带有"Expect: 100-continue"标头的请求标头后，
    // 等待服务器的第一个响应标头的时间。
	// 零表示没有超时，会立即发送请求正文，无需等待服务器批准。
	// 此时间不包括发送请求标头的时间。
	ExpectContinueTimeout time.Duration

	// TLSNextProto 指定如何在 TLS ALPN 协议协商后切换到
    // 替代协议(例如 HTTP/2)。
    // 如果 Transport 用非空协议名称拨号 TLS 连接，
    // 并且 TLSNextProto 包含该键(例如 "h2")的映射条目，
	// 则将使用请求的授权(例如"example.com"或"example.com:1234")
    // 和 TLS 连接调用该函数。
    // 该函数必须返回一个 RoundTripper，然后处理该请求。
	// 如果 TLSNextProto 不为 nil，则不会自动启用 HTTP/2 支持。
	TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper

	// ProxyConnectHeader 可选地指定在 CONNECT 
    // 请求期间发送到代理的标头。
	// 要动态设置标头，请参见 GetProxyConnectHeader。
	ProxyConnectHeader Header

	// GetProxyConnectHeader 可选地指定要在针对 ip:port 
    // 目标的 CONNECT 请求期间发送到 proxyURL 的标头的函数。
	// 如果返回错误，则 Transport 的 RoundTrip 会失败并返回该错误。
    // 它可以返回 (nil, nil) 以不添加标头。
	// 如果 GetProxyConnectHeader 不是 nil，
    // 则 ProxyConnectHeader 将被忽略。
	GetProxyConnectHeader func(ctx context.Context, proxyURL *url.URL, target string) (Header, error)

	// MaxResponseHeaderBytes 指定在服务器的响应标头中
    // 允许多少响应字节的限制。
	// 如果为零，则使用默认限制。
	MaxResponseHeaderBytes int64

	// WriteBufferSize 指定在写入传输时使用的写缓冲区的大小。
	// 如果为零，则使用默认值(当前为 4KB)。
	WriteBufferSize int

	// ReadBufferSize 指定在从传输中读取时使用的读缓冲区的大小。
	// 如果为零，则使用默认值(当前为 4KB)。
	ReadBufferSize int

	// ForceAttemptHTTP2 控制在提供了非零的 Dial、DialTLS 或 
    // DialContext 函数或 TLSClientConfig 时是否启用 HTTP/2。
	// 默认情况下，使用任何这些字段都会保守地禁用 HTTP/2。
	// 要使用自定义拨号器或 TLS 配置并仍然尝试 HTTP/2 升级，
    // 请将其设置为 true。
	ForceAttemptHTTP2 bool
	// 包含过滤或不导出的字段
}
```

​	Transport 是一个实现了 RoundTripper 接口的结构体，支持 HTTP、HTTPS 和 HTTP 代理(使用 CONNECT 实现 HTTP 或 HTTPS 代理)。

​	默认情况下，Transport 会缓存连接以便后续重用。这会在访问多个主机时导致许多开放的连接。可以使用 Transport 的 CloseIdleConnections 方法以及 MaxIdleConnsPerHost 和 DisableKeepAlives 字段来管理这种行为。

​	应该重复使用 Transport 而不是按需创建它们。Transport 可以被多个 goroutine 并发使用。

​	Transport 是用于进行 HTTP 和 HTTPS 请求的低级原语。对于高级功能(如 cookie 和重定向)，请参阅 Client。

​	Transport 在 HTTP URL 上使用 HTTP/1.1，在 HTTPS URL 上使用 HTTP/1.1 或 HTTP/2，具体取决于服务器是否支持 HTTP/2，以及 Transport 的配置方式。DefaultTransport 支持 HTTP/2。要在传输中显式启用 HTTP/2，请使用 golang.org/x/net/http2 并调用 ConfigureTransport。有关 HTTP/2 的更多信息，请参见软件包文档。

​	状态代码为 1xx 的响应要么被自动处理(100 expect-continue)，要么被忽略。唯一的例外是 HTTP 状态码 101(切换协议)，它被视为终端状态并由 RoundTrip 返回。要查看被忽略的 1xx 响应，请使用 httptrace 跟踪包的 ClientTrace.Got1xxResponse。

​	仅当请求是幂等的并且没有主体或其 Request.GetBody 已定义时，Transport 才在遇到网络错误时重试请求。如果 HTTP 请求具有 HTTP 方法 GET、HEAD、OPTIONS 或 TRACE，或者它们的 Header 映射包含 "Idempotency-Key" 或 "X-Idempotency-Key" 条目，则被视为幂等的。如果幂等键值为零长度切片，则将该请求视为幂等，但不会将头部发送到网络。

#### (*Transport) [Clone](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transport.go;l=313)  <- go1.13

``` go 
func (t *Transport) Clone() *Transport
```

​	Clone方法返回 t 的导出字段的深度副本。

#### (*Transport) [CloseIdleConnections](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transport.go;l=772) 

``` go 
func (t *Transport) CloseIdleConnections()
```

​	CloseIdleConnections方法关闭之前从先前的请求连接到的但现在处于"保持活动"状态的空闲连接。它不会中断当前正在使用的任何连接。

#### (*Transport) [RegisterProtocol](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/transport.go;l=753) 

``` go 
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)
```

​	RegisterProtocol方法注册一个新的协议，使用指定的协议名称(scheme)和RoundTripper处理程序(rt)。Transport会将使用给定scheme的请求传递给rt。它是rt的责任来模拟HTTP请求语义。

​	RegisterProtocol方法可以被其他包用来提供协议方案(scheme)的实现，例如"ftp"或"file"。

​	如果rt.RoundTrip返回ErrSkipAltProtocol，则Transport将自己处理该请求的RoundTrip，就好像未注册该协议一样。

#### (*Transport) [RoundTrip](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/roundtrip.go;l=16) 

``` go 
func (t *Transport) RoundTrip(req *Request) (*Response, error)
```

​	RoundTrip方法实现RoundTripper接口。

​	对于更高级的HTTP客户端支持(例如处理cookie和重定向)，请参见Get，Post和Client类型。

​	与RoundTripper接口一样，RoundTrip返回的错误类型是未指定的。