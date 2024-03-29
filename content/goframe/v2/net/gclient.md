+++
title = "gclient"
date = 2024-03-21T17:52:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gclient](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gclient)

Package gclient provides convenient http client functionalities.

​	软件包 gclient 提供了方便的 http 客户端功能。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func LoadKeyCrt

```go
func LoadKeyCrt(crtFile, keyFile string) (*tls.Config, error)
```

LoadKeyCrt creates and returns a TLS configuration object with given certificate and key files.

​	LoadKeyCrt 创建并返回具有给定证书和密钥文件的 TLS 配置对象。

## 类型

### type Client

```go
type Client struct {
	http.Client // Underlying HTTP Client.
	// contains filtered or unexported fields
}
```

Client is the HTTP client for HTTP request management.

​	客户端是用于 HTTP 请求管理的 HTTP 客户端。

#### func New

```go
func New() *Client
```

New creates and returns a new HTTP client object.

​	new 创建并返回新的 HTTP 客户端对象。

##### Example

``` go
```

#### (*Client) BasicAuth

```go
func (c *Client) BasicAuth(user, pass string) *Client
```

BasicAuth is a chaining function, which sets HTTP basic authentication information for next request.

​	BasicAuth 是一个链接函数，用于设置下一个请求的 HTTP 基本身份验证信息。

#### (*Client) Clone

```go
func (c *Client) Clone() *Client
```

Clone deeply clones current client and returns a new one.

​	克隆会深度克隆当前客户端并返回一个新客户端。

##### Example

``` go
```

#### (*Client) Connect

```go
func (c *Client) Connect(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Connect send CONNECT request and returns the response object. Note that the response object MUST be closed if it’ll never be used.

​	连接：发送 CONNECT 请求并返回响应对象。请注意，如果响应对象永远不会被使用，则必须将其关闭。

##### Example

``` go
```

#### (*Client) ConnectBytes

```go
func (c *Client) ConnectBytes(ctx context.Context, url string, data ...interface{}) []byte
```

ConnectBytes sends a CONNECT request, retrieves and returns the result content as bytes.

​	ConnectBytes 发送 CONNECT 请求，检索结果内容并以字节形式返回。

##### Example

``` go
```

#### (*Client) ConnectContent

```go
func (c *Client) ConnectContent(ctx context.Context, url string, data ...interface{}) string
```

ConnectContent is a convenience method for sending CONNECT request, which retrieves and returns the result content and automatically closes response object.

​	ConnectContent 是一种发送 CONNECT 请求的便捷方式，用于检索并返回结果内容并自动关闭响应对象。

##### Example

``` go
```

#### (*Client) ConnectVar

```go
func (c *Client) ConnectVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

ConnectVar sends a CONNECT request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

​	ConnectVar 发送 CONNECT 请求，检索结果内容并将其转换为 *gvar.Var。客户端在内部自动读取和关闭响应对象。结果 *gvar.Var 可以方便地转换为您想要的任何类型。

##### Example

``` go
```

#### (*Client) ContentJson

```go
func (c *Client) ContentJson() *Client
```

ContentJson is a chaining function, which sets the HTTP content type as “application/json” for the next request.

​	ContentJson 是一个链接函数，它将下一个请求的 HTTP 内容类型设置为“application/json”。

Note that it also checks and encodes the parameter to JSON format automatically.

​	请注意，它还会自动检查参数并将其编码为 JSON 格式。

##### Example

``` go
```

#### (*Client) ContentType

```go
func (c *Client) ContentType(contentType string) *Client
```

ContentType is a chaining function, which sets HTTP content type for the next request.

​	ContentType 是一个链接函数，用于设置下一个请求的 HTTP 内容类型。

#### (*Client) ContentXml

```go
func (c *Client) ContentXml() *Client
```

ContentXml is a chaining function, which sets the HTTP content type as “application/xml” for the next request.

​	ContentXml 是一个链接函数，它将下一个请求的 HTTP 内容类型设置为“application/xml”。

Note that it also checks and encodes the parameter to XML format automatically.

​	请注意，它还会自动检查参数并将其编码为 XML 格式。

#### (*Client) Cookie

```go
func (c *Client) Cookie(m map[string]string) *Client
```

Cookie is a chaining function, which sets cookie items with map for next request.

​	Cookie 是一种链接功能，它为下一个请求设置带有映射的 cookie 项。

##### Example

``` go
```

#### (*Client) Delete

```go
func (c *Client) Delete(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Delete send DELETE request and returns the response object. Note that the response object MUST be closed if it’ll never be used.

​	删除发送 DELETE 请求并返回响应对象。请注意，如果响应对象永远不会被使用，则必须将其关闭。

##### Example

``` go
```

#### (*Client) DeleteBytes

```go
func (c *Client) DeleteBytes(ctx context.Context, url string, data ...interface{}) []byte
```

DeleteBytes sends a DELETE request, retrieves and returns the result content as bytes.

​	DeleteBytes 发送 DELETE 请求，检索结果内容并以字节形式返回。

##### Example

``` go
```

#### (*Client) DeleteContent

```go
func (c *Client) DeleteContent(ctx context.Context, url string, data ...interface{}) string
```

DeleteContent is a convenience method for sending DELETE request, which retrieves and returns the result content and automatically closes response object.

​	DeleteContent 是一种发送 DELETE 请求的便捷方式，用于检索并返回结果内容并自动关闭响应对象。

##### Example

``` go
```

#### (*Client) DeleteVar

```go
func (c *Client) DeleteVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

DeleteVar sends a DELETE request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

​	DeleteVar 发送 DELETE 请求，检索结果内容并将其转换为 *gvar.Var。客户端在内部自动读取和关闭响应对象。结果 *gvar.Var 可以方便地转换为您想要的任何类型。

##### Example

``` go
```

#### (*Client) Discovery

```go
func (c *Client) Discovery(discovery gsvc.Discovery) *Client
```

Discovery is a chaining function, which sets the discovery for client. You can use `Discovery(nil)` to disable discovery feature for current client.

​	发现是一个链接函数，用于设置客户端的发现。您可以使用 `Discovery(nil)` 禁用当前客户端的发现功能。

#### (*Client) DoRequest

```go
func (c *Client) DoRequest(ctx context.Context, method, url string, data ...interface{}) (resp *Response, err error)
```

DoRequest sends request with given HTTP method and data and returns the response object. Note that the response object MUST be closed if it’ll never be used.

​	DoRequest 使用给定的 HTTP 方法和数据发送请求，并返回响应对象。请注意，如果响应对象永远不会被使用，则必须将其关闭。

Note that it uses “multipart/form-data” as its Content-Type if it contains file uploading, else it uses “application/x-www-form-urlencoded”. It also automatically detects the post content for JSON format, and for that it automatically sets the Content-Type as “application/json”.

​	请注意，如果它包含文件上传，则它使用“multipart/form-data”作为其内容类型，否则它使用“application/x-www-form-urlencoded”。它还会自动检测 JSON 格式的帖子内容，并为此自动将 Content-Type 设置为“application/json”。

#### (*Client) DoRequestObj

```go
func (c *Client) DoRequestObj(ctx context.Context, req, res interface{}) error
```

DoRequestObj does HTTP request using standard request/response object. The request object `req` is defined like:

​	DoRequestObj 使用标准请求/响应对象执行 HTTP 请求。请求对象 `req` 的定义如下：

```go
type UseCreateReq struct {
    g.Meta `path:"/user" method:"put"`
    // other fields....
}
```

The response object `res` should be a pointer type. It automatically converts result to given object `res` is success.

​	响应对象 `res` 应为指针类型。它会自动将结果转换为给定对象 `res` 即成功。

Example: var (

​	示例：var （

```
req = UseCreateReq{}
res *UseCreateRes
```

)

err := DoRequestObj(ctx, req, &res)

​	err ：= DoRequestObj（ctx， req， &res）

#### (*Client) Get

```go
func (c *Client) Get(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Get send GET request and returns the response object. Note that the response object MUST be closed if it’ll never be used.

​	获取发送 GET 请求并返回响应对象。请注意，如果响应对象永远不会被使用，则必须将其关闭。

##### Example

``` go
```

#### (*Client) GetBytes

```go
func (c *Client) GetBytes(ctx context.Context, url string, data ...interface{}) []byte
```

GetBytes sends a GET request, retrieves and returns the result content as bytes.

​	GetBytes 发送 GET 请求，检索结果内容并以字节形式返回。

##### Example

``` go
```

#### (*Client) GetContent

```go
func (c *Client) GetContent(ctx context.Context, url string, data ...interface{}) string
```

GetContent is a convenience method for sending GET request, which retrieves and returns the result content and automatically closes response object.

​	GetContent 是一种发送 GET 请求的便捷方式，它检索并返回结果内容并自动关闭响应对象。

##### Example

``` go
```

#### (*Client) GetVar

```go
func (c *Client) GetVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

GetVar sends a GET request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

​	GetVar 发送 GET 请求，检索结果内容并将其转换为 *gvar.Var。客户端在内部自动读取和关闭响应对象。结果 *gvar.Var 可以方便地转换为您想要的任何类型。

##### Example

``` go
```

#### (*Client) Head

```go
func (c *Client) Head(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Head send HEAD request and returns the response object. Note that the response object MUST be closed if it’ll never be used.

​	head 发送 HEAD 请求并返回响应对象。请注意，如果响应对象永远不会被使用，则必须将其关闭。

##### Example

``` go
```

#### (*Client) HeadBytes

```go
func (c *Client) HeadBytes(ctx context.Context, url string, data ...interface{}) []byte
```

HeadBytes sends a HEAD request, retrieves and returns the result content as bytes.

​	HeadBytes 发送 HEAD 请求，检索结果内容并以字节形式返回。

##### Example

``` go
```

#### (*Client) HeadContent

```go
func (c *Client) HeadContent(ctx context.Context, url string, data ...interface{}) string
```

HeadContent is a convenience method for sending HEAD request, which retrieves and returns the result content and automatically closes response object.

​	HeadContent 是一种发送 HEAD 请求的便捷方式，它检索并返回结果内容，并自动关闭响应对象。

##### Example

``` go
```

#### (*Client) HeadVar

```go
func (c *Client) HeadVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

HeadVar sends a HEAD request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

​	HeadVar 发送 HEAD 请求，检索结果内容并将其转换为 *gvar.Var。客户端在内部自动读取和关闭响应对象。结果 *gvar.Var 可以方便地转换为您想要的任何类型。

##### Example

``` go
```

#### (*Client) Header

```go
func (c *Client) Header(m map[string]string) *Client
```

Header is a chaining function, which sets custom HTTP headers with map for next request.

​	Header 是一个链接函数，它为下一个请求设置带有 map 的自定义 HTTP 标头。

##### Example

``` go
```

#### (*Client) HeaderRaw

```go
func (c *Client) HeaderRaw(headers string) *Client
```

HeaderRaw is a chaining function, which sets custom HTTP header using raw string for next request.

​	HeaderRaw 是一个链接函数，它使用原始字符串为下一个请求设置自定义 HTTP 标头。

##### Example

``` go
```

#### (*Client) Next

```go
func (c *Client) Next(req *http.Request) (*Response, error)
```

Next calls the next middleware. This should only be call in HandlerFunc.

​	接下来调用下一个中间件。这应该只在 HandlerFunc 中调用。

#### (*Client) NoUrlEncode

```go
func (c *Client) NoUrlEncode() *Client
```

NoUrlEncode sets the mark that do not encode the parameters before sending request.

​	NoUrlEncode 设置在发送请求之前不对参数进行编码的标记。

#### (*Client) Options

```go
func (c *Client) Options(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Options send OPTIONS request and returns the response object. Note that the response object MUST be closed if it’ll never be used.

​	Options 发送 OPTIONS 请求并返回响应对象。请注意，如果响应对象永远不会被使用，则必须将其关闭。

##### Example

``` go
```

#### (*Client) OptionsBytes

```go
func (c *Client) OptionsBytes(ctx context.Context, url string, data ...interface{}) []byte
```

OptionsBytes sends an OPTIONS request, retrieves and returns the result content as bytes.

​	OptionsBytes 发送 OPTIONS 请求，检索结果内容并以字节形式返回。

##### Example

``` go
```

#### (*Client) OptionsContent

```go
func (c *Client) OptionsContent(ctx context.Context, url string, data ...interface{}) string
```

OptionsContent is a convenience method for sending OPTIONS request, which retrieves and returns the result content and automatically closes response object.

​	OptionsContent 是一种发送 OPTIONS 请求的便捷方式，它检索并返回结果内容并自动关闭响应对象。

##### Example

``` go
```

#### (*Client) OptionsVar

```go
func (c *Client) OptionsVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

OptionsVar sends an OPTIONS request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

​	OptionsVar 发送 OPTIONS 请求，检索结果内容并将其转换为 *gvar.Var。客户端在内部自动读取和关闭响应对象。结果 *gvar.Var 可以方便地转换为您想要的任何类型。

##### Example

``` go
```

#### (*Client) Patch

```go
func (c *Client) Patch(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Patch send PATCH request and returns the response object. Note that the response object MUST be closed if it’ll never be used.

​	Patch 发送 PATCH 请求并返回响应对象。请注意，如果响应对象永远不会被使用，则必须将其关闭。

##### Example

``` go
```

#### (*Client) PatchBytes

```go
func (c *Client) PatchBytes(ctx context.Context, url string, data ...interface{}) []byte
```

PatchBytes sends a PATCH request, retrieves and returns the result content as bytes.

​	PatchBytes 发送 PATCH 请求，检索结果内容并以字节形式返回。

##### Example

``` go
```

#### (*Client) PatchContent

```go
func (c *Client) PatchContent(ctx context.Context, url string, data ...interface{}) string
```

PatchContent is a convenience method for sending PATCH request, which retrieves and returns the result content and automatically closes response object.

​	PatchContent 是一种发送 PATCH 请求的便捷方式，它检索并返回结果内容并自动关闭响应对象。

##### Example

``` go
```

#### (*Client) PatchVar

```go
func (c *Client) PatchVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

PatchVar sends a PATCH request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

​	PatchVar 发送 PATCH 请求，检索结果内容并将其转换为 *gvar.Var。客户端在内部自动读取和关闭响应对象。结果 *gvar.Var 可以方便地转换为您想要的任何类型。

##### Example

``` go
```

#### (*Client) Post

```go
func (c *Client) Post(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Post sends request using HTTP method POST and returns the response object. Note that the response object MUST be closed if it’ll never be used.

​	POST 使用 HTTP 方法 POST 发送请求并返回响应对象。请注意，如果响应对象永远不会被使用，则必须将其关闭。

##### Example

``` go
```

#### (*Client) PostBytes

```go
func (c *Client) PostBytes(ctx context.Context, url string, data ...interface{}) []byte
```

PostBytes sends a POST request, retrieves and returns the result content as bytes.

​	PostBytes 发送 POST 请求，检索结果内容并以字节形式返回。

##### Example

``` go
```

#### (*Client) PostContent

```go
func (c *Client) PostContent(ctx context.Context, url string, data ...interface{}) string
```

PostContent is a convenience method for sending POST request, which retrieves and returns the result content and automatically closes response object.

​	PostContent 是一种发送 POST 请求的便捷方式，用于检索并返回结果内容并自动关闭响应对象。

##### Example

``` go
```

#### (*Client) PostForm

```go
func (c *Client) PostForm(ctx context.Context, url string, data map[string]string) (resp *Response, err error)
```

PostForm is different from net/http.PostForm. It’s a wrapper of Post method, which sets the Content-Type as “multipart/form-data;”. and It will automatically set boundary characters for the request body and Content-Type.

​	PostForm 与 net/http 不同。PostForm的。它是 Post 方法的包装器，它将 Content-Type 设置为“multipart/form-data;”。它将自动为请求正文和 Content-Type 设置边界字符。

It’s Seem like the following case:

​	这似乎是以下情况：

Content-Type: multipart/form-data; boundary=—-Boundarye4Ghaog6giyQ9ncN

​	内容类型：multipart/form-data;boundary=—-Boundarye4Ghaog6giyQ9ncN

And form data is like: ——Boundarye4Ghaog6giyQ9ncN Content-Disposition: form-data; name=“checkType”

​	而表单数据是这样的： ——Boundarye4Ghaog6giyQ9ncN Content-Disposition： form-data;名称=“checkType”

none

​	没有

It’s used for sending form data. Note that the response object MUST be closed if it’ll never be used.

​	它用于发送表单数据。请注意，如果响应对象永远不会被使用，则必须将其关闭。

#### (*Client) PostVar

```go
func (c *Client) PostVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

PostVar sends a POST request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

​	PostVar 发送 POST 请求，检索结果内容并将其转换为 *gvar.Var。客户端在内部自动读取和关闭响应对象。结果 *gvar.Var 可以方便地转换为您想要的任何类型。

##### Example

``` go
```

#### (*Client) Prefix

```go
func (c *Client) Prefix(prefix string) *Client
```

Prefix is a chaining function, which sets the URL prefix for next request of this client. Eg: Prefix(“http://127.0.0.1:8199/api/v1”) Prefix(“http://127.0.0.1:8199/api/v2”)

​	Prefix 是一个链接函数，它为该客户端的下一个请求设置 URL 前缀。例如：prefix（“http://127.0.0.1:8199/api/v1”） prefix（“http://127.0.0.1:8199/api/v2”）

##### Example

``` go
```

#### (*Client) Proxy

```go
func (c *Client) Proxy(proxyURL string) *Client
```

Proxy is a chaining function, which sets proxy for next request. Make sure you pass the correct `proxyURL`. The correct pattern is like `[http://USER:PASSWORD@IP:PORT](http://USER:PASSWORD@ip:PORT/)` or `socks5://USER:PASSWORD@IP:PORT`. Only `http` and `socks5` proxies are supported currently.

​	Proxy 是一个链接函数，它为下一个请求设置代理。确保通过正确的 `proxyURL` .正确的模式是 like `[http://USER:PASSWORD@IP:PORT](http://USER:PASSWORD@ip:PORT/)` 或 `socks5://USER:PASSWORD@IP:PORT` 。目前仅 `http` 支持代理和 `socks5` 代理。

##### Example

``` go
```

#### (*Client) Put

```go
func (c *Client) Put(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Put send PUT request and returns the response object. Note that the response object MUST be closed if it’ll never be used.

​	put send PUT 请求并返回响应对象。请注意，如果响应对象永远不会被使用，则必须将其关闭。

##### Example

``` go
```

#### (*Client) PutBytes

```go
func (c *Client) PutBytes(ctx context.Context, url string, data ...interface{}) []byte
```

PutBytes sends a PUT request, retrieves and returns the result content as bytes.

​	PutBytes 发送 PUT 请求，检索结果内容并以字节形式返回。

##### Example

``` go
```

#### (*Client) PutContent

```go
func (c *Client) PutContent(ctx context.Context, url string, data ...interface{}) string
```

PutContent is a convenience method for sending PUT request, which retrieves and returns the result content and automatically closes response object.

​	PutContent 是一种发送 PUT 请求的便捷方式，用于检索并返回结果内容并自动关闭响应对象。

##### Example

``` go
```

#### (*Client) PutVar

```go
func (c *Client) PutVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

PutVar sends a PUT request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

​	PutVar 发送 PUT 请求，检索结果内容并将其转换为 *gvar.Var。客户端在内部自动读取和关闭响应对象。结果 *gvar.Var 可以方便地转换为您想要的任何类型。

##### Example

``` go
```

#### (*Client) RedirectLimit

```go
func (c *Client) RedirectLimit(redirectLimit int) *Client
```

RedirectLimit is a chaining function, which sets the redirect limit the number of jumps for the request.

​	RedirectLimit 是一个链接函数，它设置重定向限制请求的跳转次数。

##### Example

``` go
```

#### (*Client) RequestBytes

```go
func (c *Client) RequestBytes(ctx context.Context, method string, url string, data ...interface{}) []byte
```

RequestBytes sends request using given HTTP method and data, retrieves returns the result as bytes. It reads and closes the response object internally automatically.

​	RequestBytes 使用给定的 HTTP 方法和数据发送请求，检索以字节形式返回结果。它会在内部自动读取和关闭响应对象。

#### (*Client) RequestContent

```go
func (c *Client) RequestContent(ctx context.Context, method string, url string, data ...interface{}) string
```

RequestContent is a convenience method for sending custom http method request, which retrieves and returns the result content and automatically closes response object.

​	RequestContent 是一种发送自定义 http 方法请求的便捷方式，用于检索并返回结果内容并自动关闭响应对象。

##### Example

``` go
```

#### (*Client) RequestVar

```go
func (c *Client) RequestVar(ctx context.Context, method string, url string, data ...interface{}) *gvar.Var
```

RequestVar sends request using given HTTP method and data, retrieves converts the result to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

​	RequestVar 使用给定的 HTTP 方法和数据发送请求，检索将结果转换为 *gvar.Var。客户端在内部自动读取和关闭响应对象。结果 *gvar.Var 可以方便地转换为您想要的任何类型。

#### (*Client) Retry

```go
func (c *Client) Retry(retryCount int, retryInterval time.Duration) *Client
```

Retry is a chaining function, which sets retry count and interval when failure for next request.

​	重试是一个链接函数，用于设置下一个请求失败时的重试计数和间隔。

##### Example

``` go
```

#### (*Client) SetAgent

```go
func (c *Client) SetAgent(agent string) *Client
```

SetAgent sets the User-Agent header for client.

​	SetAgent 为客户端设置 User-Agent 标头。

#### (*Client) SetBasicAuth

```go
func (c *Client) SetBasicAuth(user, pass string) *Client
```

SetBasicAuth sets HTTP basic authentication information for the client.

​	SetBasicAuth 设置客户端的 HTTP 基本身份验证信息。

#### (*Client) SetBrowserMode

```go
func (c *Client) SetBrowserMode(enabled bool) *Client
```

SetBrowserMode enables browser mode of the client. When browser mode is enabled, it automatically saves and sends cookie content from and to server.

​	SetBrowserMode 启用客户端的浏览器模式。启用浏览器模式后，它会自动保存和发送来自服务器和发送到服务器的 cookie 内容。

##### Example

``` go
```

#### (*Client) SetBuilder

```go
func (c *Client) SetBuilder(builder gsel.Builder)
```

SetBuilder sets the load balance builder for client.

​	SetBuilder 为客户端设置负载平衡生成器。

#### (*Client) SetContentType

```go
func (c *Client) SetContentType(contentType string) *Client
```

SetContentType sets HTTP content type for the client.

​	SetContentType 设置客户端的 HTTP 内容类型。

#### (*Client) SetCookie

```go
func (c *Client) SetCookie(key, value string) *Client
```

SetCookie sets a cookie pair for the client.

​	SetCookie 为客户端设置 Cookie 对。

#### (*Client) SetCookieMap

```go
func (c *Client) SetCookieMap(m map[string]string) *Client
```

SetCookieMap sets cookie items with map.

​	SetCookieMap 使用 map 设置 cookie 项。

#### (*Client) SetDiscovery

```go
func (c *Client) SetDiscovery(discovery gsvc.Discovery)
```

SetDiscovery sets the load balance builder for client.

​	SetDiscovery 为客户端设置负载平衡生成器。

#### (*Client) SetHeader

```go
func (c *Client) SetHeader(key, value string) *Client
```

SetHeader sets a custom HTTP header pair for the client.

​	SetHeader 为客户端设置自定义 HTTP 标头对。

##### Example

``` go
```

#### (*Client) SetHeaderMap

```go
func (c *Client) SetHeaderMap(m map[string]string) *Client
```

SetHeaderMap sets custom HTTP headers with map.

​	SetHeaderMap 使用 map 设置自定义 HTTP 标头。

#### (*Client) SetHeaderRaw

```go
func (c *Client) SetHeaderRaw(headers string) *Client
```

SetHeaderRaw sets custom HTTP header using raw string.

​	SetHeaderRaw 使用原始字符串设置自定义 HTTP 标头。

#### (*Client) SetNoUrlEncode

```go
func (c *Client) SetNoUrlEncode(noUrlEncode bool) *Client
```

SetNoUrlEncode sets the mark that do not encode the parameters before sending request.

​	SetNoUrlEncode 设置在发送请求之前不对参数进行编码的标记。

#### (*Client) SetPrefix

```go
func (c *Client) SetPrefix(prefix string) *Client
```

SetPrefix sets the request server URL prefix.

​	SetPrefix 设置请求服务器 URL 前缀。

#### (*Client) SetProxy

```go
func (c *Client) SetProxy(proxyURL string)
```

SetProxy set proxy for the client. This func will do nothing when the parameter `proxyURL` is empty or in wrong pattern. The correct pattern is like `[http://USER:PASSWORD@IP:PORT](http://USER:PASSWORD@ip:PORT/)` or `socks5://USER:PASSWORD@IP:PORT`. Only `http` and `socks5` proxies are supported currently.

​	SetProxy 为客户端设置代理。当参数 `proxyURL` 为空或模式错误时，此函数将不执行任何操作。正确的模式是 like `[http://USER:PASSWORD@IP:PORT](http://USER:PASSWORD@ip:PORT/)` 或 `socks5://USER:PASSWORD@IP:PORT` 。目前仅 `http` 支持代理和 `socks5` 代理。

##### Example

``` go
```

#### (*Client) SetRedirectLimit

```go
func (c *Client) SetRedirectLimit(redirectLimit int) *Client
```

SetRedirectLimit limits the number of jumps.

​	SetRedirectLimit 限制跳转次数。

##### Example

``` go
```

#### (*Client) SetRetry

```go
func (c *Client) SetRetry(retryCount int, retryInterval time.Duration) *Client
```

SetRetry sets retry count and interval.

​	SetRetry 设置重试计数和间隔。

#### (*Client) SetTLSConfig

```go
func (c *Client) SetTLSConfig(tlsConfig *tls.Config) error
```

SetTLSConfig sets the TLS configuration of client.

​	SetTLSConfig 设置客户端的 TLS 配置。

##### Example

``` go
```

#### (*Client) SetTLSKeyCrt

```go
func (c *Client) SetTLSKeyCrt(crtFile, keyFile string) error
```

SetTLSKeyCrt sets the certificate and key file for TLS configuration of client.

​	SetTLSKeyCrt 设置客户端 TLS 配置的证书和密钥文件。

##### Example

``` go
```

#### (*Client) SetTimeout

```go
func (c *Client) SetTimeout(t time.Duration) *Client
```

SetTimeout sets the request timeout for the client.

​	SetTimeout 设置客户端的请求超时。

#### (*Client) Timeout

```go
func (c *Client) Timeout(t time.Duration) *Client
```

Timeout is a chaining function, which sets the timeout for next request.

​	Timeout 是一个链接函数，用于设置下一个请求的超时。

#### (*Client) Trace

```go
func (c *Client) Trace(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Trace send TRACE request and returns the response object. Note that the response object MUST be closed if it’ll never be used.

​	Trace 发送 TRACE 请求并返回响应对象。请注意，如果响应对象永远不会被使用，则必须将其关闭。

##### Example

``` go
```

#### (*Client) TraceBytes

```go
func (c *Client) TraceBytes(ctx context.Context, url string, data ...interface{}) []byte
```

TraceBytes sends a TRACE request, retrieves and returns the result content as bytes.

​	TraceBytes 发送 TRACE 请求，检索结果内容并以字节形式返回。

##### Example

``` go
```

#### (*Client) TraceContent

```go
func (c *Client) TraceContent(ctx context.Context, url string, data ...interface{}) string
```

TraceContent is a convenience method for sending TRACE request, which retrieves and returns the result content and automatically closes response object.

​	TraceContent 是一种发送 TRACE 请求的便捷方式，用于检索并返回结果内容并自动关闭响应对象。

##### Example

``` go
```

#### (*Client) TraceVar

```go
func (c *Client) TraceVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

TraceVar sends a TRACE request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

​	TraceVar 发送 TRACE 请求，检索结果内容并将其转换为 *gvar.Var。客户端在内部自动读取和关闭响应对象。结果 *gvar.Var 可以方便地转换为您想要的任何类型。

##### Example

``` go
```

#### (*Client) Use

```go
func (c *Client) Use(handlers ...HandlerFunc) *Client
```

Use adds one or more middleware handlers to client.

​	Use 将一个或多个中间件处理程序添加到客户端。

### type HandlerFunc

```go
type HandlerFunc = func(c *Client, r *http.Request) (*Response, error)
```

HandlerFunc middleware handler func

​	HandlerFunc 中间件处理程序 func

### type Response

```go
type Response struct {
	*http.Response // Response is the underlying http.Response object of certain request.
	// contains filtered or unexported fields
}
```

Response is the struct for client request response.

​	响应是客户端请求响应的结构。

#### (*Response) Close

```go
func (r *Response) Close() error
```

Close closes the response when it will never be used.

​	关闭 （Close） 在永远不会使用的响应时关闭响应。

#### (*Response) GetCookie

```go
func (r *Response) GetCookie(key string) string
```

GetCookie retrieves and returns the cookie value of specified `key`.

​	GetCookie 检索并返回指定的 `key` cookie 值。

#### (*Response) GetCookieMap

```go
func (r *Response) GetCookieMap() map[string]string
```

GetCookieMap retrieves and returns a copy of current cookie values map.

​	GetCookieMap 检索并返回当前 Cookie 值映射的副本。

#### (*Response) Raw

```go
func (r *Response) Raw() string
```

Raw returns the raw text of the request and the response.

​	Raw 返回请求和响应的原始文本。

#### (*Response) RawDump

```go
func (r *Response) RawDump()
```

RawDump outputs the raw text of the request and the response to stdout.

​	RawDump 输出请求的原始文本和对 stdout 的响应。

#### (*Response) RawRequest

```go
func (r *Response) RawRequest() string
```

RawRequest returns the raw content of the request.

​	RawRequest 返回请求的原始内容。

#### (*Response) RawResponse

```go
func (r *Response) RawResponse() string
```

RawResponse returns the raw content of the response.

​	RawResponse 返回响应的原始内容。

#### (*Response) ReadAll

```go
func (r *Response) ReadAll() []byte
```

ReadAll retrieves and returns the response content as []byte.

​	ReadAll 检索响应内容并以 []byte 的形式返回。

#### (*Response) ReadAllString

```go
func (r *Response) ReadAllString() string
```

ReadAllString retrieves and returns the response content as string.

​	ReadAllString 检索响应内容并将其作为字符串返回。

#### (*Response) SetBodyContent

```go
func (r *Response) SetBodyContent(content []byte)
```

SetBodyContent overwrites response content with custom one.

​	SetBodyContent 使用自定义响应内容覆盖响应内容。

### type WebSocketClient

```go
type WebSocketClient struct {
	*websocket.Dialer
}
```

WebSocketClient wraps the underlying websocket client connection and provides convenient functions.

​	WebSocketClient 包装底层 websocket 客户端连接，提供便捷的功能。

#### func NewWebSocket

```go
func NewWebSocket() *WebSocketClient
```

NewWebSocket creates and returns a new WebSocketClient object.

​	NewWebSocket 创建并返回一个新的 WebSocketClient 对象。