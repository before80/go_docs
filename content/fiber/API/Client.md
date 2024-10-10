+++
title = "Client"
date = 2024-02-05T09:14:15+08:00
weight = 40
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/client]({{< ref "/fiber/API/Client" >}})

# 🌎 Client

## Start request 开始请求

Start a http request with http method and url.
使用 http 方法和 url 启动 http 请求。

Signatures
签名

```go
// Client http methods
func (c *Client) Get(url string) *Agent
func (c *Client) Head(url string) *Agent
func (c *Client) Post(url string) *Agent
func (c *Client) Put(url string) *Agent
func (c *Client) Patch(url string) *Agent
func (c *Client) Delete(url string) *Agent
```



Here we present a brief example demonstrating the simulation of a proxy using our `*fiber.Agent` methods.
这里我们提供了一个简短的示例，演示如何使用我们的 `*fiber.Agent` 方法模拟代理。

```go
// Get something
func getSomething(c *fiber.Ctx) (err error) {
    agent := fiber.Get("<URL>")
    statusCode, body, errs := agent.Bytes()
    if len(errs) > 0 {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "errs": errs,
        })
    }

    var something fiber.Map
    err = json.Unmarshal(body, &something)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "err": err,
        })
    }

    return c.Status(statusCode).JSON(something)
}

// Post something
func createSomething(c *fiber.Ctx) (err error) {
    agent := fiber.Post("<URL>")
    agent.Body(c.Body()) // set body received by request
    statusCode, body, errs := agent.Bytes()
    if len(errs) > 0 {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "errs": errs,
        })
    }

    // pass status code and body received by the proxy
    return c.Status(statusCode).Send(body)
}
```



Based on this short example, we can perceive that using the `*fiber.Client` is very straightforward and intuitive.
基于这个简短的示例，我们可以看出使用 `*fiber.Client` 非常简单直观。

## ✨ Agent

`Agent` is built on top of FastHTTP's [`HostClient`](https://github.com/valyala/fasthttp/blob/master/client.go#L603) which has lots of convenient helper methods such as dedicated methods for request methods.
`Agent` 建立在 FastHTTP 的 `HostClient` 之上，它有很多方便的帮助程序方法，例如针对请求方法的专用方法。

### Parse

Parse initializes a HostClient.
Parse 初始化一个 HostClient。

Parse

```go
a := AcquireAgent()
req := a.Request()
req.Header.SetMethod(MethodGet)
req.SetRequestURI("http://example.com")

if err := a.Parse(); err != nil {
    panic(err)
}

code, body, errs := a.Bytes() // ...
```



### Set

Set sets the given `key: value` header.
Set 设置给定的 `key: value` 头。

Signature
签名

```go
func (a *Agent) Set(k, v string) *Agent
func (a *Agent) SetBytesK(k []byte, v string) *Agent
func (a *Agent) SetBytesV(k string, v []byte) *Agent
func (a *Agent) SetBytesKV(k []byte, v []byte) *Agent
```



Example
示例

```go
agent.Set("k1", "v1").
    SetBytesK([]byte("k1"), "v1").
    SetBytesV("k1", []byte("v1")).
    SetBytesKV([]byte("k2"), []byte("v2"))
// ...
```



### Add

Add adds the given `key: value` header. Multiple headers with the same key may be added with this function.
Add 添加给定的 `key: value` 标头。可以使用此函数添加具有相同键的多个标头。

Signature
签名

```go
func (a *Agent) Add(k, v string) *Agent
func (a *Agent) AddBytesK(k []byte, v string) *Agent
func (a *Agent) AddBytesV(k string, v []byte) *Agent
func (a *Agent) AddBytesKV(k []byte, v []byte) *Agent
```



Example
示例

```go
agent.Add("k1", "v1").
    AddBytesK([]byte("k1"), "v1").
    AddBytesV("k1", []byte("v1")).
    AddBytesKV([]byte("k2"), []byte("v2"))
// Headers:
// K1: v1
// K1: v1
// K1: v1
// K2: v2
```



### ConnectionClose

ConnectionClose adds the `Connection: close` header.
ConnectionClose 添加 `Connection: close` 标头。

Signature
签名

```go
func (a *Agent) ConnectionClose() *Agent
```



Example
示例

```go
agent.ConnectionClose()
// ...
```



### UserAgent

UserAgent sets `User-Agent` header value.
UserAgent 设置 `User-Agent` 标头值。

Signature
签名

```go
func (a *Agent) UserAgent(userAgent string) *Agent
func (a *Agent) UserAgentBytes(userAgent []byte) *Agent
```



Example
示例

```go
agent.UserAgent("fiber")
// ...
```



### Cookie

Cookie sets a cookie in `key: value` form. `Cookies` can be used to set multiple cookies.
Cookie 以 `key: value` 形式设置 cookie。 `Cookies` 可用于设置多个 cookie。

Signature
签名

```go
func (a *Agent) Cookie(key, value string) *Agent
func (a *Agent) CookieBytesK(key []byte, value string) *Agent
func (a *Agent) CookieBytesKV(key, value []byte) *Agent
func (a *Agent) Cookies(kv ...string) *Agent
func (a *Agent) CookiesBytesKV(kv ...[]byte) *Agent
```



Example
示例

```go
agent.Cookie("k", "v")
agent.Cookies("k1", "v1", "k2", "v2")
// ...
```



### Referer

Referer sets the Referer header value.
Referer 设置 Referer 标头值。

Signature
签名

```go
func (a *Agent) Referer(referer string) *Agent
func (a *Agent) RefererBytes(referer []byte) *Agent
```



Example
示例

```go
agent.Referer("https://docs.gofiber.io")
// ...
```



### ContentType

ContentType sets Content-Type header value.
ContentType 设置 Content-Type 标头值。

Signature
签名

```go
func (a *Agent) ContentType(contentType string) *Agent
func (a *Agent) ContentTypeBytes(contentType []byte) *Agent
```



Example
示例

```go
agent.ContentType("custom-type")
// ...
```



### Host

Host sets the Host header.
Host 设置 Host 头。

Signature
签名

```go
func (a *Agent) Host(host string) *Agent
func (a *Agent) HostBytes(host []byte) *Agent
```



Example
示例

```go
agent.Host("example.com")
// ...
```



### QueryString

QueryString sets the URI query string.
QueryString 设置 URI 查询字符串。

Signature
签名

```go
func (a *Agent) QueryString(queryString string) *Agent
func (a *Agent) QueryStringBytes(queryString []byte) *Agent
```



Example
示例

```go
agent.QueryString("foo=bar")
// ...
```



### BasicAuth

BasicAuth sets the URI username and password using HTTP Basic Auth.
BasicAuth 使用 HTTP 基本身份验证设置 URI 用户名和密码。

Signature
签名

```go
func (a *Agent) BasicAuth(username, password string) *Agent
func (a *Agent) BasicAuthBytes(username, password []byte) *Agent
```



Example
示例

```go
agent.BasicAuth("foo", "bar")
// ...
```



### Body

There are several ways to set request body.
有几种方法可以设置请求正文。

Signature
签名

```go
func (a *Agent) BodyString(bodyString string) *Agent
func (a *Agent) Body(body []byte) *Agent

// BodyStream sets request body stream and, optionally body size.
//
// If bodySize is >= 0, then the bodyStream must provide exactly bodySize bytes
// before returning io.EOF.
//
// If bodySize < 0, then bodyStream is read until io.EOF.
//
// bodyStream.Close() is called after finishing reading all body data
// if it implements io.Closer.
//
// Note that GET and HEAD requests cannot have body.
func (a *Agent) BodyStream(bodyStream io.Reader, bodySize int) *Agent
```



Example
示例

```go
agent.BodyString("foo=bar")
agent.Body([]byte("bar=baz"))
agent.BodyStream(strings.NewReader("body=stream"), -1)
// ...
```



### JSON JSON

JSON sends a JSON request by setting the Content-Type header to the `ctype` parameter. If no `ctype` is passed in, the header is set to `application/json`.
JSON 通过将 Content-Type 头设置为 `ctype` 参数来发送 JSON 请求。如果未传入 `ctype` ，则将头设置为 `application/json` 。

Signature
签名

```go
func (a *Agent) JSON(v interface{}, ctype ...string) *Agent
```



Example
示例

```go
agent.JSON(fiber.Map{"success": true})
// ...
```



### XML

XML sends an XML request by setting the Content-Type header to `application/xml`.
XML 通过将 Content-Type 头设置为 `application/xml` 来发送 XML 请求。

Signature
签名

```go
func (a *Agent) XML(v interface{}) *Agent
```



Example
示例

```go
agent.XML(fiber.Map{"success": true})
// ...
```



### Form

Form sends a form request by setting the Content-Type header to `application/x-www-form-urlencoded`.
Form 通过将 Content-Type 头设置为 `application/x-www-form-urlencoded` 来发送表单请求。

Signature
签名

```go
// Form sends form request with body if args is non-nil.
//
// It is recommended obtaining args via AcquireArgs and release it
// manually in performance-critical code.
func (a *Agent) Form(args *Args) *Agent
```



Example
示例

```go
args := AcquireArgs()
args.Set("foo", "bar")

agent.Form(args)
// ...
ReleaseArgs(args)
```



### MultipartForm

MultipartForm sends multipart form request by setting the Content-Type header to `multipart/form-data`. These requests can include key-value's and files.
MultipartForm 通过将 Content-Type 头设置为 `multipart/form-data` 来发送多部分表单请求。这些请求可以包括键值和文件。

Signature
签名

```go
// MultipartForm sends multipart form request with k-v and files.
//
// It is recommended to obtain args via AcquireArgs and release it
// manually in performance-critical code.
func (a *Agent) MultipartForm(args *Args) *Agent
```



Example
示例

```go
args := AcquireArgs()
args.Set("foo", "bar")

agent.MultipartForm(args)
// ...
ReleaseArgs(args)
```



Fiber provides several methods for sending files. Note that they must be called before `MultipartForm`.
Fiber 提供了多种发送文件的方法。请注意，它们必须在 `MultipartForm` 之前调用。

#### Boundary

Boundary sets boundary for multipart form request.
Boundary 为多部分表单请求设置边界。

Signature
签名

```go
func (a *Agent) Boundary(boundary string) *Agent
```



Example
示例

```go
agent.Boundary("myBoundary")
    .MultipartForm(nil)
// ...
```



#### SendFile(s)

SendFile read a file and appends it to a multipart form request. Sendfiles can be used to append multiple files.
SendFile 读取文件并将其追加到多部分表单请求中。Sendfiles 可用于追加多个文件。

Signature
签名

```go
func (a *Agent) SendFile(filename string, fieldname ...string) *Agent
func (a *Agent) SendFiles(filenamesAndFieldnames ...string) *Agent
```



Example
示例

```go
agent.SendFile("f", "field name")
    .SendFiles("f1", "field name1", "f2").
    .MultipartForm(nil)
// ...
```



#### FileData

FileData appends file data for multipart form request.
FileData 追加文件数据以进行多部分表单请求。

```go
// FormFile represents multipart form file
type FormFile struct {
    // Fieldname is form file's field name
    Fieldname string
    // Name is form file's name
    Name string
    // Content is form file's content
    Content []byte
}
```



Signature
签名

```go
// FileData appends files for multipart form request.
//
// It is recommended obtaining formFile via AcquireFormFile and release it
// manually in performance-critical code.
func (a *Agent) FileData(formFiles ...*FormFile) *Agent
```



Example
示例

```go
ff1 := &FormFile{"filename1", "field name1", []byte("content")}
ff2 := &FormFile{"filename2", "field name2", []byte("content")}
agent.FileData(ff1, ff2).
    MultipartForm(nil)
// ...
```



### Debug

Debug mode enables logging request and response detail to `io.writer`(default is `os.Stdout`).
调试模式将请求和响应详细信息记录到 `io.writer` （默认值为 `os.Stdout` ）。

Signature
签名

```go
func (a *Agent) Debug(w ...io.Writer) *Agent
```



Example
示例

```go
agent.Debug()
// ...
```



### Timeout 超时 

Timeout sets request timeout duration.
超时设置请求超时持续时间。

Signature
签名

```go
func (a *Agent) Timeout(timeout time.Duration) *Agent
```



Example
示例

```go
agent.Timeout(time.Second)
// ...
```



### Reuse 重用 

Reuse enables the Agent instance to be used again after one request. If agent is reusable, then it should be released manually when it is no longer used.
重用使 Agent 实例在一次请求后可再次使用。如果 agent 可重用，则应在不再使用时手动释放它。

Signature
签名

```go
func (a *Agent) Reuse() *Agent
```



Example
示例

```go
agent.Reuse()
// ...
```



### InsecureSkipVerify

InsecureSkipVerify controls whether the Agent verifies the server certificate chain and host name.
InsecureSkipVerify 控制 Agent 是否验证服务器证书链和主机名。

Signature
签名

```go
func (a *Agent) InsecureSkipVerify() *Agent
```



Example
示例

```go
agent.InsecureSkipVerify()
// ...
```



### TLSConfig

TLSConfig sets tls config.
TLSConfig 设置 tls 配置。

Signature
签名

```go
func (a *Agent) TLSConfig(config *tls.Config) *Agent
```



Example
示例

```go
// Create tls certificate
cer, _ := tls.LoadX509KeyPair("pem", "key")

config := &tls.Config{
    Certificates: []tls.Certificate{cer},
}

agent.TLSConfig(config)
// ...
```



### MaxRedirectsCount

MaxRedirectsCount sets max redirect count for GET and HEAD.
MaxRedirectsCount 设置 GET 和 HEAD 的最大重定向计数。

Signature
签名

```go
func (a *Agent) MaxRedirectsCount(count int) *Agent
```



Example
示例

```go
agent.MaxRedirectsCount(7)
// ...
```



### JSONEncoder

JSONEncoder sets custom json encoder.
JSONEncoder 设置自定义 JSON 编码器。

Signature
签名

```go
func (a *Agent) JSONEncoder(jsonEncoder utils.JSONMarshal) *Agent
```



Example
示例

```go
agent.JSONEncoder(json.Marshal)
// ...
```



### JSONDecoder

JSONDecoder sets custom json decoder.
JSONDecoder 设置自定义 JSON 解码器。

Signature
签名

```go
func (a *Agent) JSONDecoder(jsonDecoder utils.JSONUnmarshal) *Agent
```



Example
示例

```go
agent.JSONDecoder(json.Unmarshal)
// ...
```



### Request

Request returns Agent request instance.
Request 返回 Agent 请求实例。

Signature
签名

```go
func (a *Agent) Request() *Request
```



Example
示例

```go
req := agent.Request()
// ...
```



### SetResponse

SetResponse sets custom response for the Agent instance. It is recommended obtaining custom response via AcquireResponse and release it manually in performance-critical code.
SetResponse 为 Agent 实例设置自定义响应。建议通过 AcquireResponse 获取自定义响应，并在性能关键代码中手动释放它。

Signature
签名

```go
func (a *Agent) SetResponse(customResp *Response) *Agent
```



Example
示例

```go
resp := AcquireResponse()
agent.SetResponse(resp)
// ...
ReleaseResponse(resp)
```



<details class="details_lb9f isBrowser_bmU9 alert alert--info details_b_Ee" data-collapsed="true" data-immersive-translate-walked="180ddf95-b199-453f-9777-d54368f1701b" style="box-sizing: border-box; color: var(--ifm-alert-foreground-color); --ifm-alert-background-color: var(--ifm-color-info-contrast-background); --ifm-alert-background-color-highlight: #54c7ec26; --ifm-alert-foreground-color: var(--ifm-color-info-contrast-foreground); --ifm-alert-border-color: var(--ifm-color-info-dark); --ifm-code-background: var(--ifm-alert-background-color-highlight); --ifm-link-color: var(--ifm-alert-foreground-color); --ifm-link-hover-color: var(--ifm-alert-foreground-color); --ifm-link-decoration: underline; --ifm-tabs-color: var(--ifm-alert-foreground-color); --ifm-tabs-color-active: var(--ifm-alert-foreground-color); --ifm-tabs-color-active-border: var(--ifm-alert-border-color); background-color: var(--ifm-alert-background-color); border: 1px solid var(--ifm-alert-border-color); border-radius: var(--ifm-alert-border-radius); box-shadow: var(--ifm-alert-shadow); padding: var(--ifm-alert-padding-vertical) var(--ifm-alert-padding-horizontal); --docusaurus-details-summary-arrow-size: 0.38rem; --docusaurus-details-transition: transform var(--ifm-transition-fast) ease; --docusaurus-details-decoration-color: var(--ifm-alert-border-color); margin: 0 0 var(--ifm-spacing-vertical); font-family: system-ui, -apple-system, &quot;Segoe UI&quot;, Roboto, Ubuntu, Cantarell, &quot;Noto Sans&quot;, sans-serif, BlinkMacSystemFont, &quot;Segoe UI&quot;, Helvetica, Arial, sans-serif, &quot;Apple Color Emoji&quot;, &quot;Segoe UI Emoji&quot;, &quot;Segoe UI Symbol&quot;; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;"><summary data-immersive-translate-walked="180ddf95-b199-453f-9777-d54368f1701b" data-immersive-translate-paragraph="1" style="box-sizing: border-box; list-style: none; cursor: pointer; padding-left: 1rem; position: relative;">Example handling for response values<font class="notranslate immersive-translate-target-wrapper" lang="zh-CN" data-immersive-translate-translation-element-mark="1" style="box-sizing: border-box;"><br style="box-sizing: border-box;"><font class="notranslate immersive-translate-target-translation-theme-none immersive-translate-target-translation-block-wrapper-theme-none immersive-translate-target-translation-block-wrapper" data-immersive-translate-translation-element-mark="1" style="box-sizing: border-box; display: inline-block; margin: 8px 0px !important;"><font class="notranslate immersive-translate-target-inner immersive-translate-target-translation-theme-none-inner" data-immersive-translate-translation-element-mark="1" style="box-sizing: border-box; font-family: inherit;">响应值的示例处理</font></font></font></summary></details>

### Dest

Dest sets custom dest. The contents of dest will be replaced by the response body, if the dest is too small a new slice will be allocated.
Dest 设置自定义 dest。如果 dest 太小，则将用响应正文替换 dest 的内容，并将分配一个新的切片。

Signature
签名

```go
func (a *Agent) Dest(dest []byte) *Agent {
```



Example
示例

```go
agent.Dest(nil)
// ...
```



### Bytes 字节 

Bytes returns the status code, bytes body and errors of url.
Bytes 返回 url 的状态代码、字节正文和错误。

Signature
签名

```go
func (a *Agent) Bytes() (code int, body []byte, errs []error)
```



Example
示例

```go
code, body, errs := agent.Bytes()
// ...
```



### String String

String returns the status code, string body and errors of url.
String 返回 url 的状态码、字符串主体和错误。

Signature
签名

```go
func (a *Agent) String() (int, string, []error)
```



Example
示例

```go
code, body, errs := agent.String()
// ...
```



### Struct Struct

Struct returns the status code, bytes body and errors of url. And bytes body will be unmarshalled to given v.
Struct 返回 url 的状态码、字节主体和错误。字节主体将被解析为给定的 v。

Signature
签名

```go
func (a *Agent) Struct(v interface{}) (code int, body []byte, errs []error)
```



Example
示例

```go
var d data
code, body, errs := agent.Struct(&d)
// ...
```



### RetryIf RetryIf

RetryIf controls whether a retry should be attempted after an error. By default, will use isIdempotent function from fasthttp
RetryIf 控制是否应在错误后尝试重试。默认情况下，将使用 fasthttp 的 isIdempotent 函数

Signature
签名

```go
func (a *Agent) RetryIf(retryIf RetryIfFunc) *Agent
```



Example
示例

```go
agent.Get("https://example.com").RetryIf(func (req *fiber.Request) bool {
    return req.URI() == "https://example.com"
})
// ...
```
