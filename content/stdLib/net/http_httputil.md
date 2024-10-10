+++
title = "http/httputil"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/http/httputil@go1.23.0](https://pkg.go.dev/net/http/httputil@go1.23.0)

Package httputil provides HTTP utility functions, complementing the more common ones in the net/http package.

​	`httputil` 包提供了 HTTP 的实用函数，补充了 net/http 包中更常见的函数。


## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/httputil/persist.go;l=17)

``` go 
var (
	// 已弃用：不再使用。
	ErrPersistEOF = &http.ProtocolError{ErrorString: "persistent connection closed"}

	// 已弃用：不再使用。
	ErrClosed = &http.ProtocolError{ErrorString: "connection closed by user"}

	// 已弃用：不再使用。
	ErrPipeline = &http.ProtocolError{ErrorString: "pipeline error"}
)
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/httputil/httputil.go;l=41)

``` go 
var ErrLineTooLong = internal.ErrLineTooLong
```

ErrLineTooLong is returned when reading malformed chunked data with lines that are too long.

​	在读取具有过长行的格式错误的块数据时返回 ErrLineTooLong。

## 函数

#### func DumpRequest 

``` go 
func DumpRequest(req *http.Request, body bool) ([]byte, error)
```

DumpRequest returns the given request in its HTTP/1.x wire representation. It should only be used by servers to debug client requests. The returned representation is an approximation only; some details of the initial request are lost while parsing it into an http.Request. In particular, the order and case of header field names are lost. The order of values in multi-valued headers is kept intact. HTTP/2 requests are dumped in HTTP/1.x form, not in their original binary representations.

​	DumpRequest 函数返回给定请求的 HTTP/1.x 线路表示（wire representation）。它应仅由服务器用于调试客户端请求。返回的表示仅为近似值；将初始请求解析为 http.Request 时，将丢失一些细节。特别是，标头字段名称的顺序和大小写将丢失。多值标头中的值的顺序保持不变。HTTP/2 请求以 HTTP/1.x 形式转储，而不是以其原始二进制表示形式。

If body is true, DumpRequest also returns the body. To do so, it consumes req.Body and then replaces it with a new io.ReadCloser that yields the same bytes. If DumpRequest returns an error, the state of req is undefined.

​	如果 body 为 true，则 DumpRequest 还返回主体。为此，它会消耗 req.Body，然后将其替换为一个新的 io.ReadCloser，该 io.ReadCloser 产生相同的字节。如果 DumpRequest 函数返回错误，则 req 的状态未定义。

The documentation for http.Request.Write details which fields of req are included in the dump.

​	有关 http.Request.Write 的文档详细说明了在转储中包含的 req 的哪些字段。

##### DumpRequest Example
``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%q", dump)
	}))
	defer ts.Close()

	const body = "Go is a general-purpose language designed with systems programming in mind."
	req, err := http.NewRequest("POST", ts.URL, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.Host = "www.example.org"
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}

Output:

"POST / HTTP/1.1\r\nHost: www.example.org\r\nAccept-Encoding: gzip\r\nContent-Length: 75\r\nUser-Agent: Go-http-client/1.1\r\n\r\nGo is a general-purpose language designed with systems programming in mind."
```

#### func DumpRequestOut 

``` go 
func DumpRequestOut(req *http.Request, body bool) ([]byte, error)
```

DumpRequestOut is like DumpRequest but for outgoing client requests. It includes any headers that the standard http.Transport adds, such as User-Agent.

​	DumpRequestOut 函数类似于 DumpRequest 函数，但用于传出的客户端请求。它包括标准 http.Transport 添加的任何标头，例如 User-Agent。

##### Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	const body = "Go is a general-purpose language designed with systems programming in mind."
	req, err := http.NewRequest("PUT", "http://www.example.org", strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q", dump)

}

```

#### func DumpResponse 

``` go 
func DumpResponse(resp *http.Response, body bool) ([]byte, error)
```

DumpResponse is like DumpRequest but dumps a response.

​	DumpResponse 函数类似于 DumpRequest函数，但会转储响应。

##### Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func main() {
	const body = "Go is a general-purpose language designed with systems programming in mind."
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Date", "Wed, 19 Jul 1972 19:00:00 GMT")
		fmt.Fprintln(w, body)
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q", dump)

}

Output:

"PUT / HTTP/1.1\r\nHost: www.example.org\r\nUser-Agent: Go-http-client/1.1\r\nContent-Length: 75\r\nAccept-Encoding: gzip\r\n\r\nGo is a general-purpose language designed with systems programming in mind."
```

#### func NewChunkedReader 

``` go 
func NewChunkedReader(r io.Reader) io.Reader
```

NewChunkedReader returns a new chunkedReader that translates the data read from r out of HTTP "chunked" format before returning it. The chunkedReader returns io.EOF when the final 0-length chunk is read.

​	NewChunkedReader 函数返回一个新的 chunkedReader，该 chunkedReader 将从 r 中读取的数据从 HTTP "chunked" 格式转换为普通格式。当读取到最后的 0 长度块时，chunkedReader 返回 io.EOF。

NewChunkedReader is not needed by normal applications. The http package automatically decodes chunking when reading response bodies.

​	通常应用程序不需要 NewChunkedReader函数。在读取响应主体时，http 包会自动解码分块。

#### func NewChunkedWriter 

``` go 
func NewChunkedWriter(w io.Writer) io.WriteCloser
```

NewChunkedWriter returns a new chunkedWriter that translates writes into HTTP "chunked" format before writing them to w. Closing the returned chunkedWriter sends the final 0-length chunk that marks the end of the stream but does not send the final CRLF that appears after trailers; trailers and the last CRLF must be written separately.

​	NewChunkedWriter 返回一个新的 chunkedWriter，它会在将写入的数据转换成 HTTP 的 "chunked" 格式之后再写入到 w 中。关闭返回的 chunkedWriter 会发送表示流结束的最终长度为 0 的块，但不会发送出现在 trailers 之后的最终 CRLF；trailers 和最后的 CRLF 必须单独写入。

NewChunkedWriter is not needed by normal applications. The http package adds chunking automatically if handlers don't set a Content-Length header. Using NewChunkedWriter inside a handler would result in double chunking or chunking with a Content-Length length, both of which are wrong.

​	正常应用程序不需要 NewChunkedWriter函数。如果处理程序未设置 Content-Length 标头，http 包会自动添加分块。在处理程序内部使用 NewChunkedWriter 会导致双重分块或带有 Content-Length 长度的分块，这两种情况都是错误的。

## 类型

### type BufferPool  <- go1.6

``` go 
type BufferPool interface {
	Get() []byte
	Put([]byte)
}
```

A BufferPool is an interface for getting and returning temporary byte slices for use by io.CopyBuffer.

​	BufferPool 是一个接口，用于获取和返回供 io.CopyBuffer 使用的临时字节片。

### type ClientConn <- DEPRECATED

``` go
type ClientConn struct {
	// contains filtered or unexported fields
}
```

ClientConn is an artifact of Go's early HTTP implementation. It is low-level, old, and unused by Go's current HTTP stack. We should have deleted it before Go 1.

​	ClientConn 是 Go 早期 HTTP 实现的产物。它属于底层、旧版，并且 Go 当前的 HTTP 堆栈未使用它。我们应该在 Go 1 之前将其删除。

Deprecated: Use Client or Transport in package net/http instead.

​	已弃用：改用 net/http 包中的 Client 或 Transport。

#### func NewClientConn <- DEPRECATED

```go
func NewClientConn(c net.Conn, r *bufio.Reader) *ClientConn
```

NewClientConn is an artifact of Go's early HTTP implementation. It is low-level, old, and unused by Go's current HTTP stack. We should have deleted it before Go 1.

​	NewClientConn 是 Go 早期 HTTP 实现的产物。它属于底层、旧版，并且 Go 当前的 HTTP 堆栈未使用它。我们应该在 Go 1 之前将其删除。

Deprecated: Use the Client or Transport in package net/http instead.

​	已弃用：改用 net/http 包中的 Client 或 Transport。

#### func NewProxyClientConn <- DEPRECATED

```go
func NewProxyClientConn(c net.Conn, r *bufio.Reader) *ClientConn
```

NewProxyClientConn is an artifact of Go's early HTTP implementation. It is low-level, old, and unused by Go's current HTTP stack. We should have deleted it before Go 1.

​	NewProxyClientConn 是 Go 早期 HTTP 实现的产物。它属于底层、旧版，并且 Go 当前的 HTTP 堆栈未使用它。我们应该在 Go 1 之前将其删除。

Deprecated: Use the Client or Transport in package net/http instead.

​	已弃用：改用 net/http 包中的 Client 或 Transport。

#### func (*ClientConn)Close

``` go
func (cc *ClientConn) Close() error
```

Close calls Hijack and then also closes the underlying connection.

​	Close 调用 Hijack，然后还关闭底层连接。

#### func (*ClientConn) Do

``` go
func (cc *ClientConn) Do(req *http.Request) (*http.Response, error)
```

Do is convenience method that writes a request and reads a response.

​	Do 是一个便捷方法，用于编写请求并读取响应。

#### func (*ClientConn) Hijack

``` go
func (cc *ClientConn) Hijack() (c net.Conn, r *bufio.Reader)
```

Hijack detaches the ClientConn and returns the underlying connection as well as the read-side bufio which may have some left over data. Hijack may be called before the user or Read have signaled the end of the keep-alive logic. The user should not call Hijack while Read or Write is in progress.

​	Hijack 分离 ClientConn 并返回基础连接以及可能包含一些剩余数据的读取端 bufio。在用户或 Read 发出保持活动逻辑结束信号之前，可以调用 Hijack。用户在 Read 或 Write 正在进行时不应调用 Hijack。

#### func (*ClientConn) Pending

``` go
func (cc *ClientConn) Pending() int
```

Pending returns the number of unanswered requests that have been sent on the connection.

​	Pending 返回已在连接上发送的未应答请求数。

#### func (*ClientConn) Read

``` go
func (cc *ClientConn) Read(req *http.Request) (resp *http.Response, err error)
```

Read reads the next response from the wire. A valid response might be returned together with an ErrPersistEOF, which means that the remote requested that this be the last request serviced. Read can be called concurrently with Write, but not with another Read.

​	Read 从网络读取下一个响应。可能会返回一个有效的响应以及 ErrPersistEOF，这意味着远程请求这是要服务的最后一个请求。Read 可以与 Write 并发调用，但不能与另一个 Read 并发调用。

#### func (*ClientConn) Write

``` go
func (cc *ClientConn) Write(req *http.Request) error
```

Write writes a request. An ErrPersistEOF error is returned if the connection has been closed in an HTTP keep-alive sense. If req.Close equals true, the keep-alive connection is logically closed after this request and the opposing server is informed. An ErrUnexpectedEOF indicates the remote closed the underlying TCP connection, which is usually considered as graceful close.

​	Write 写入请求。如果连接已在 HTTP 保持活动意义上关闭，则返回 ErrPersistEOF 错误。如果 req.Close 等于 true，则在此请求后逻辑上关闭保持活动连接，并通知对端服务器。ErrUnexpectedEOF 表示远程关闭了底层 TCP 连接，这通常被视为正常关闭。

### type ProxyRequest  <- go1.20

``` go 
type ProxyRequest struct {
    // In 是代理接收的请求。
    // 在调用 Rewrite 函数后，不能修改 In。
	In *http.Request

    // Out 是代理将发送的请求。
    // 在调用 Rewrite 函数中可以修改或替换此请求。
    // 在调用 Rewrite 之前，从此请求中移除了逐跳标头。
	Out *http.Request
}
```

A ProxyRequest contains a request to be rewritten by a ReverseProxy.

​	ProxyRequest 包含要由 ReverseProxy 重写的请求。

#### (*ProxyRequest) SetURL  <- go1.20

``` go 
func (r *ProxyRequest) SetURL(target *url.URL)
```

SetURL routes the outbound request to the scheme, host, and base path provided in target. If the target's path is "/base" and the incoming request was for "/dir", the target request will be for "/base/dir".

​	SetURL 方法将出站请求路由到 target 提供的方案、主机和基本路径。如果 target 的路径是 "/base"，并且传入的请求是 "/dir"，则目标请求将为 /base/dir。

SetURL rewrites the outbound Host header to match the target's host. To preserve the inbound request's Host header (the default behavior of NewSingleHostReverseProxy):

​	SetURL 方法会将出站的 Host 头重写为与 target  主机相匹配。要保留入站请求的 Host 头（NewSingleHostReverseProxy 函数的默认行为）：

```go
rewriteFunc := func(r *httputil.ProxyRequest) {
	r.SetURL(url)
	r.Out.Host = r.In.Host
}
```

#### (*ProxyRequest) SetXForwarded  <- go1.20

``` go 
func (r *ProxyRequest) SetXForwarded()
```

SetXForwarded sets the X-Forwarded-For, X-Forwarded-Host, and X-Forwarded-Proto headers of the outbound request.

​	SetXForwarded 方法设置出站请求的 X-Forwarded-For、X-Forwarded-Host 和 X-Forwarded-Proto 标头。 

- The X-Forwarded-For header is set to the client IP address.
- X-Forwarded-For 标头设置为客户端 IP 地址。
- The X-Forwarded-Host header is set to the host name requested by the client.
- X-Forwarded-Host 标头设置为客户端请求的主机名。
- The X-Forwarded-Proto header is set to "http" or "https", depending on whether the inbound request was made on a TLS-enabled connection.
- X-Forwarded-Proto 标头设置为 "http" 或 "https"，取决于传入请求是否在启用 TLS 的连接上进行。

If the outbound request contains an existing X-Forwarded-For header, SetXForwarded appends the client IP address to it. To append to the inbound request's X-Forwarded-For header (the default behavior of ReverseProxy when using a Director function), copy the header from the inbound request before calling SetXForwarded:

​	如果出站请求包含现有的 X-Forwarded-For 标头，SetXForwarded 方法会将客户端 IP 地址追加到其中。要将入站请求的 X-Forwarded-For 标头追加（使用 Director 函数时的 ReverseProxy 的默认行为），请在调用 SetXForwarded 方法之前从传入请求中复制标头：

```go
rewriteFunc := func(r *httputil.ProxyRequest) {
	r.Out.Header["X-Forwarded-For"] = r.In.Header["X-Forwarded-For"]
	r.SetXForwarded()
}
```

### type ReverseProxy 

``` go 
type ReverseProxy struct {
    // Rewrite 必须是修改请求的函数，以便将其修改为使用 Transport 发送的新请求。
    // 然后将其响应不加修改地复制回原始客户端。
    // 在返回后，Rewrite 不得访问提供的 ProxyRequest 或其内容。
    //
    // 在调用 Rewrite 之前，将从出站请求中删除 Forwarded、X-Forwarded、X-Forwarded-Host
    // 和 X-Forwarded-Proto 标头。请参阅 ProxyRequest.SetXForwarded 方法。
    //
    // 在调用 Rewrite 之前，将从出站请求中删除不可解析的查询参数。
    // 如果 Rewrite 函数将入站 URL 的 RawQuery 复制到出站 URL 以保留原始参数字符串，
    // 则可能会导致安全问题。注意，如果代理的查询参数解释与下游服务器的解释不匹配，可能会导致安全问题。
    //
    // Rewrite 和 Director 最多只能设置一个。
	Rewrite func(*ProxyRequest)

    // Director 是一个函数，用于修改请求，以便将其修改为使用 Transport 发送的新请求。
    // 然后将其响应不加修改地复制回原始客户端。
    // 在返回后，Director 不得访问提供的 Request。
    //
    // 默认情况下，X-Forwarded-For 标头设置为客户端 IP 地址。如果已存在 X-Forwarded-For 标头，
    // 则将客户端 IP 追加到现有值中。作为特殊情况，如果在 Request.Header 映射中存在标头，
    // 但其值为 nil（例如，由 Director 函数设置），则不会修改 X-Forwarded-For 标头。
    //
    // Director 返回后，从请求中删除逐跳标头，这可能会删除 Director 添加的标头。
    // 若要确保请求的修改保留，请改用 Rewrite 函数。
    //
    // 如果在 Director 返回后设置了 Request.Form，将从出站请求中删除不可解析的查询参数。
    //
    // Rewrite 和 Director 最多只能设置一个。
	Director func(*http.Request)

    // 用于执行代理请求的传输。
    // 如果为 nil，则使用 http.DefaultTransport。
	Transport http.RoundTripper

    // FlushInterval 指定在复制响应体时向客户端刷新的间隔。
    // 如果为零，则不会进行定期刷新。
    // 负值意味着在每次写入客户端后立即刷新。
    // 当 ReverseProxy 将响应识别为流响应时，或者其 ContentLength 为 -1 时，将忽略 FlushInterval；
    // 对于这种响应，写入会立即刷新到客户端。
	FlushInterval time.Duration

    // ErrorLog 指定一个可选的用于处理尝试代理请求时出现的错误的记录器。
    // 如果为 nil，则通过 log 包的标准记录器进行记录。
	ErrorLog *log.Logger

    // BufferPool 可选地指定用于在复制 HTTP 响应体时为 io.CopyBuffer 获取字节片的缓冲池。
	BufferPool BufferPool

    // ModifyResponse 是一个可选的函数，用于修改来自后端的 Response。
    // 如果后端返回响应（无论具有任何 HTTP 状态码），则调用此函数。
    // 如果无法访问后端，将调用可选的 ErrorHandler，而不调用 ModifyResponse。
    //
    // 如果 ModifyResponse 返回错误，将使用其错误值调用 ErrorHandler。
    // 如果 ErrorHandler 为 nil，则使用其默认实现。
	ModifyResponse func(*http.Response) error

    // ErrorHandler 是一个可选的函数，用于处理到达后端的错误或来自 ModifyResponse 的错误。
    //
    // 如果为 nil，则默认行为是记录所提供的错误并返回 502 状态 Bad Gateway 响应。
	ErrorHandler func(http.ResponseWriter, *http.Request, error)
}
```

ReverseProxy is an HTTP Handler that takes an incoming request and sends it to another server, proxying the response back to the client.

​	ReverseProxy 是一个 HTTP 处理程序，它接收传入请求并将其发送到另一个服务器，然后将响应代理回客户端。

1xx responses are forwarded to the client if the underlying transport supports ClientTrace.Got1xxResponse.

​	如果底层传输支持 ClientTrace.Got1xxResponse （ClientTrace是在http.trace包中定义的结构体），则 1xx 响应会转发到客户端。

#### Example
``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

func main() {
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this call was relayed by the reverse proxy")
	}))
	defer backendServer.Close()

	rpURL, err := url.Parse(backendServer.URL)
	if err != nil {
		log.Fatal(err)
	}
	frontendProxy := httptest.NewServer(&httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetXForwarded()
			r.SetURL(rpURL)
		},
	})
	defer frontendProxy.Close()

	resp, err := http.Get(frontendProxy.URL)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}

Output:

this call was relayed by the reverse proxy
```

#### func NewSingleHostReverseProxy 

``` go 
func NewSingleHostReverseProxy(target *url.URL) *ReverseProxy
```

NewSingleHostReverseProxy returns a new ReverseProxy that routes URLs to the scheme, host, and base path provided in target. If the target's path is "/base" and the incoming request was for "/dir", the target request will be for /base/dir.

​	NewSingleHostReverseProxy 函数返回一个新的 ReverseProxy实例，该 ReverseProxy 实例将 URL 路由到 target 提供的方案、主机和基本路径。如果 target 的路径是 "/base"，并且传入的请求是 "/dir"，则目标请求将为 `/base/dir`。

NewSingleHostReverseProxy does not rewrite the Host header.

​	NewSingleHostReverseProxy 函数不会重写 Host 标头。

To customize the ReverseProxy behavior beyond what NewSingleHostReverseProxy provides, use ReverseProxy directly with a Rewrite function. The ProxyRequest SetURL method may be used to route the outbound request. (Note that SetURL, unlike NewSingleHostReverseProxy, rewrites the Host header of the outbound request by default.)

​	如果要在 NewSingleHostReverseProxy 函数提供的功能之外定制 ReverseProxy 行为，请直接使用 ReverseProxy 并使用 Rewrite 函数。ProxyRequest 结构体的 SetURL 方法可用于路由出站请求。（请注意，与 NewSingleHostReverseProxy 函数不同，SetURL 方法默认情况下会重写出站请求的 Host 标头。）

```go
proxy := &ReverseProxy{
	Rewrite: func(r *ProxyRequest) {
		r.SetURL(target)
		r.Out.Host = r.In.Host // if desired
	}
}
```

#### (*ReverseProxy) ServeHTTP 

``` go 
func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request)
```

### type ServerConn  DEPRECATED

``` go 
type ServerConn struct {
	// contains filtered or unexported fields
}
```

ServerConn is an artifact of Go's early HTTP implementation. It is low-level, old, and unused by Go's current HTTP stack. We should have deleted it before Go 1.

​	ServerConn 是 Go 早期 HTTP 实现的产物。它是低级别的、旧的，并且在当前 Go HTTP 栈中未被使用。在 Go 1 之前，我们应该已经删除它。

Deprecated: Use the Server in package net/http instead.


​	已弃用：请使用 net/http 包中的 Server。

#### func NewServerConn DEPRECATED

``` go 
func NewServerConn(c net.Conn, r *bufio.Reader) *ServerConn
```

NewServerConn is an artifact of Go's early HTTP implementation. It is low-level, old, and unused by Go's current HTTP stack. We should have deleted it before Go 1.

​	NewServerConn 是 Go 早期 HTTP 实现的产物。它是低级别的、旧的，并且在当前 Go HTTP 栈中未被使用。在 Go 1 之前，我们应该已经删除它。

Deprecated: Use the Server in package net/http instead.

​	已弃用：请使用 net/http 包中的 Server。

#### (*ServerConn) Close 

``` go 
func (sc *ServerConn) Close() error
```

Close calls Hijack and then also closes the underlying connection.

​	Close 函数调用 Hijack方法，然后也关闭底层连接。

#### (*ServerConn) Hijack 

``` go 
func (sc *ServerConn) Hijack() (net.Conn, *bufio.Reader)
```

Hijack detaches the ServerConn and returns the underlying connection as well as the read-side bufio which may have some left over data. Hijack may be called before Read has signaled the end of the keep-alive logic. The user should not call Hijack while Read or Write is in progress.

​	Hijack 方法分离 ServerConn，并返回底层连接以及可能有一些剩余数据的读取侧 bufio。在 Read 信号化保持活动逻辑结束之前，可以调用 Hijack。在进行 Read 或 Write 时，用户不应该调用 Hijack。

#### (*ServerConn) Pending 

``` go 
func (sc *ServerConn) Pending() int
```

Pending returns the number of unanswered requests that have been received on the connection.

​	Pending 方法返回在连接上接收到的未回复的请求数。

#### (*ServerConn) Read 

``` go 
func (sc *ServerConn) Read() (*http.Request, error)
```

Read returns the next request on the wire. An ErrPersistEOF is returned if it is gracefully determined that there are no more requests (e.g. after the first request on an HTTP/1.0 connection, or after a Connection:close on a HTTP/1.1 connection).

​	Read 方法返回下一个在传输线上的请求。如果在连接上优雅地确定没有更多请求（例如，在 HTTP/1.0 连接的第一个请求之后，或在 HTTP/1.1 连接上的 Connection:close 之后），则返回 ErrPersistEOF。

#### (*ServerConn) Write 

``` go 
func (sc *ServerConn) Write(req *http.Request, resp *http.Response) error
```

Write writes resp in response to req. To close the connection gracefully, set the Response.Close field to true. Write should be considered operational until it returns an error, regardless of any errors returned on the Read side.

​	Write 方法将 resp 作为对 req 的响应进行写入。要优雅地关闭连接，请将 Response.Close 字段设置为 true。在返回错误之前，无论在读取侧返回了什么错误，都应将 Write 方法视为可操作的。