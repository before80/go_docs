+++
title = "http/httptest"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/http/httptest@go1.21.3](https://pkg.go.dev/net/http/httptest@go1.21.3)

Package httptest provides utilities for HTTP testing.

​	`httptest` 包提供了用于 HTTP 测试的实用工具。


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/httptest/recorder.go;l=61)

``` go 
const DefaultRemoteAddr = "1.2.3.4"
```

DefaultRemoteAddr is the default remote address to return in RemoteAddr if an explicit DefaultRemoteAddr isn't set on ResponseRecorder.

​	DefaultRemoteAddr 常量是在未显式设置 ResponseRecorder （结构体） 的 DefaultRemoteAddr 时，用于在 RemoteAddr （http.Request结构体中的字段）中返回的默认远程地址。

## 变量

This section is empty.

## 函数

### func NewRequest  <- go1.7

``` go 
func NewRequest(method, target string, body io.Reader) *http.Request
```

NewRequest returns a new incoming server Request, suitable for passing to an http.Handler for testing.

​	NewRequest 函数返回一个新的服务器入站请求，适用于传递给 http.Handler 进行测试。

The target is the [RFC 7230](https://rfc-editor.org/rfc/rfc7230.html) "request-target": it may be either a path or an absolute URL. If target is an absolute URL, the host name from the URL is used. Otherwise, "example.com" is used.

​	目标是 [RFC 7230](https://rfc-editor.org/rfc/rfc7230.html) 中的 "request-target"：它可以是路径或绝对 URL。如果目标是绝对 URL，则使用 URL 中的主机名。否则，使用 "example.com"。

The TLS field is set to a non-nil dummy value if target has scheme "https".

​	如果目标的 scheme 是 "https"，则会将 TLS 字段设置为非 nil 的虚拟值。

The Request.Proto is always HTTP/1.1.

​	Request.Proto 总是 HTTP/1.1。

An empty method means "GET".

​	空方法表示 "GET"。

The provided body may be nil. If the body is of type *bytes.Reader, *strings.Reader, or *bytes.Buffer, the Request.ContentLength is set.

​	提供的 body 可以为 nil。如果 body 的类型是 *bytes.Reader、*strings.Reader 或 `*bytes.Buffer`，则会设置 Request.ContentLength。

NewRequest panics on error for ease of use in testing, where a panic is acceptable.

​	为了在测试中更容易使用，NewRequest 函数在出错时会引发 panic，这在测试中是可以接受的。

To generate a client HTTP request instead of a server request, see the NewRequest function in the net/http package.

​	要生成客户端的 HTTP 请求而不是服务器请求，请参见 net/http 包中的 NewRequest 函数。

## 类型

### type ResponseRecorder 

``` go 
type ResponseRecorder struct {
    // Code is the HTTP response code set by WriteHeader.
	//
	// Note that if a Handler never calls WriteHeader or Write,
	// this might end up being 0, rather than the implicit
	// http.StatusOK. To get the implicit value, use the Result
	// method.
    // Code 是由 WriteHeader 设置的 HTTP 响应码。
	//
	// 请注意，如果 Handler 从未调用 WriteHeader 或 Write，
	// 这可能会变为 0，而不是隐式的 http.StatusOK。要获取隐式值，请使用 Result 方法。
	Code int
	
    // HeaderMap contains the headers explicitly set by the Handler.
	// It is an internal detail.
	//
	// Deprecated: HeaderMap exists for historical compatibility
	// and should not be used. To access the headers returned by a handler,
	// use the Response.Header map as returned by the Result method.
    // HeaderMap 包含由 Handler 显式设置的头部信息。
	// 这是一个内部细节。
	//
	// 已弃用：HeaderMap 存在于历史兼容性，
	// 不应使用。要访问处理程序返回的头部信息，请使用 Result 方法返回的 Response 值的 Header。
	HeaderMap http.Header

    // Body is the buffer to which the Handler's Write calls are sent.
	// If nil, the Writes are silently discarded.
	// Body 是 Handler 的 Write 调用发送到的缓冲区。
	// 如果为 nil，则 Writes 会被静默丢弃。
	Body *bytes.Buffer

    // Flushed is whether the Handler called Flush.
    // Flushed 是 Handler 是否调用了 Flush。
	Flushed bool
    // 包含已过滤或未公开的字段
}
```

ResponseRecorder is an implementation of http.ResponseWriter that records its mutations for later inspection in tests.

​	ResponseRecorder 是 http.ResponseWriter 的一种实现，它记录了其变化，以便以后在测试中进行检查。

#### Example
``` go 
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

}

```

#### func NewRecorder 

``` go 
func NewRecorder() *ResponseRecorder
```

NewRecorder returns an initialized ResponseRecorder.

​	NewRecorder 函数返回一个初始化后的 ResponseRecorder实例。

#### (*ResponseRecorder) Flush 

``` go 
func (rw *ResponseRecorder) Flush()
```

Flush implements http.Flusher. To test whether Flush was called, see rw.Flushed.

​	Flush 方法实现了 http.Flusher。要测试是否调用了 Flush，请查看 rw.Flushed。

#### (*ResponseRecorder) Header 

``` go 
func (rw *ResponseRecorder) Header() http.Header
```

Header implements http.ResponseWriter. It returns the response headers to mutate within a handler. To test the headers that were written after a handler completes, use the Result method and see the returned Response value's Header.

​	Header 方法实现了 http.ResponseWriter。它返回响应头部，以在处理程序内进行修改。要测试处理程序完成后写入的头部，请使用 Result 方法，并查看返回的 Response 值的 Header。

#### (*ResponseRecorder) Result  <- go1.7

``` go 
func (rw *ResponseRecorder) Result() *http.Response
```

Result returns the response generated by the handler.

​	Result 方法返回由处理程序生成的响应。

The returned Response will have at least its StatusCode, Header, Body, and optionally Trailer populated. More fields may be populated in the future, so callers should not DeepEqual the result in tests.

​	返回的 Response 至少会填充其 StatusCode、Header、Body，以及可能填充 Trailer。未来可能填充更多字段，因此调用者不应在测试中使用 DeepEqual 比较结果。

The Response.Header is a snapshot of the headers at the time of the first write call, or at the time of this call, if the handler never did a write.

​	Response.Header 是在第一次写入调用时的头部的快照，或在此调用时的时间，如果处理程序从未进行过写入。

The Response.Body is guaranteed to be non-nil and Body.Read call is guaranteed to not return any error other than io.EOF.

​	Response.Body 保证为非 nil，并且 Body.Read 调用保证不会返回除 io.EOF 之外的任何错误。

The Response.Body is guaranteed to be non-nil and Body.Read call is guaranteed to not return any error other than io.EOF.

​	只有在处理程序完成运行后才能调用 Result。

#### (*ResponseRecorder) Write 

``` go 
func (rw *ResponseRecorder) Write(buf []byte) (int, error)
```

Write implements http.ResponseWriter. The data in buf is written to rw.Body, if not nil.

​	Write 方法实现了 http.ResponseWriter。如果 buf 中的数据不为 nil，则将数据写入 rw.Body。

#### (*ResponseRecorder) WriteHeader 

``` go 
func (rw *ResponseRecorder) WriteHeader(code int)
```

WriteHeader implements http.ResponseWriter.

​	WriteHeader 方法实现了 http.ResponseWriter。

#### (*ResponseRecorder) WriteString  <- go1.6

``` go 
func (rw *ResponseRecorder) WriteString(str string) (int, error)
```

WriteString implements io.StringWriter. The data in str is written to rw.Body, if not nil.

​	WriteString 方法实现了 io.StringWriter。如果 str 中的数据不为 nil，则将数据写入 rw.Body。

### type Server 

``` go 
type Server struct {
	URL      string // 基本 URL，格式为 http://ipaddr:port，不带尾随斜杠 base URL of form http://ipaddr:port with no trailing slash
	Listener net.Listener
	
    // EnableHTTP2 controls whether HTTP/2 is enabled
	// on the server. It must be set between calling
	// NewUnstartedServer and calling Server.StartTLS.
    // EnableHTTP2 控制是否在服务器上启用 HTTP/2。
	// 必须在调用 NewUnstartedServer 和调用 Server.StartTLS 之间设置它。
	EnableHTTP2 bool

    
    // TLS 是可选的 TLS 配置，在启动 TLS 后会填充一个新配置。
	// 如果在调用 StartTLS 之前在未启动的服务器上设置，
	// 则会将现有字段复制到新配置中。
	TLS *tls.Config

    // Config may be changed after calling NewUnstartedServer and
	// before Start or StartTLS.
    // Config 可以在调用 NewUnstartedServer 之后、调用 Start 或 StartTLS 之前更改。
	Config *http.Server
    // 包含已过滤或未公开的字段
}
```

A Server is an HTTP server listening on a system-chosen port on the local loopback interface, for use in end-to-end HTTP tests.

​	Server 是一个 HTTP 服务器，监听本地环回接口上的系统选择端口，用于进行端到端的 HTTP 测试。

#### Example
``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}

```

#### Example (HTTP2)
``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", r.Proto)
	}))
	ts.EnableHTTP2 = true
	ts.StartTLS()
	defer ts.Close()

	res, err := ts.Client().Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", greeting)

}

```

#### func NewServer 

``` go 
func NewServer(handler http.Handler) *Server
```

NewServer starts and returns a new Server. The caller should call Close when finished, to shut it down.

​	NewServer 函数启动并返回一个新的 Server实例。调用者应在完成后调用 Close方法，以关闭它。

#### func NewTLSServer 

``` go 
func NewTLSServer(handler http.Handler) *Server
```

NewTLSServer starts and returns a new Server using TLS. The caller should call Close when finished, to shut it down.

​	NewTLSServer 格式启动并返回一个使用 TLS 的新 Server实例。调用者应在完成后调用 Close方法，以关闭它。

##### NewTLSServer  Example
``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	client := ts.Client()
	res, err := client.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}

```

#### func NewUnstartedServer 

``` go 
func NewUnstartedServer(handler http.Handler) *Server
```

NewUnstartedServer returns a new Server but doesn't start it.

​	NewUnstartedServer 函数返回一个新的 Server实例，但不会启动它。

After changing its configuration, the caller should call Start or StartTLS.

​	在更改其配置后，调用者应在调用 Start 方法或 StartTLS 方法之前调用它。

The caller should call Close when finished, to shut it down.

​	调用者应在完成后调用 Close方法，以关闭它。

#### (*Server) Certificate  <- go1.9

``` go 
func (s *Server) Certificate() *x509.Certificate
```

Certificate returns the certificate used by the server, or nil if the server doesn't use TLS.

​	Certificate 方法返回服务器使用的证书，如果服务器不使用 TLS，则返回 nil。

#### (*Server) Client  <- go1.9

``` go 
func (s *Server) Client() *http.Client
```

Client returns an HTTP client configured for making requests to the server. It is configured to trust the server's TLS test certificate and will close its idle connections on Server.Close.

​	Client 方法返回配置为向服务器发出请求的 HTTP 客户端。它配置为信任服务器的 TLS 测试证书，并且将在 Server.Close 时关闭其空闲连接（idle connections）。

#### (*Server) Close 

``` go 
func (s *Server) Close()
```

Close shuts down the server and blocks until all outstanding requests on this server have completed.

​	Close 方法关闭服务器并阻塞，直到此服务器上的所有未完成请求完成。

#### (*Server) CloseClientConnections 

``` go 
func (s *Server) CloseClientConnections()
```

CloseClientConnections closes any open HTTP connections to the test Server.

​	CloseClientConnections 方法关闭到测试 Server 的任何打开的 HTTP 连接。

#### (*Server) Start 

``` go 
func (s *Server) Start()
```

Start starts a server from NewUnstartedServer.

​	Start 方法启动一个从 NewUnstartedServer方法（返回的）服务器。

#### (*Server) StartTLS 

``` go 
func (s *Server) StartTLS()
```

StartTLS starts TLS on a server from NewUnstartedServer.

​	StartTLS 方法启动一个从 NewUnstartedServer方法（返回的）服务器，并在其上启动 TLS。