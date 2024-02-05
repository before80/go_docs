+++
title = "Ctx"
date = 2024-02-05T09:14:15+08:00
weight = 20
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/ctx]({{< ref "/fiber/API/Ctx" >}})

# 🧠 Ctx

## Accepts

Checks, if the specified **extensions** or **content** **types** are acceptable.

​	检查指定的扩展名或内容类型是否可接受。

INFO
信息

Based on the request’s [Accept](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept) HTTP header.

​	基于请求的 Accept HTTP 头。

Signature
签名

```go
func (c *Ctx) Accepts(offers ...string)          string
func (c *Ctx) AcceptsCharsets(offers ...string)  string
func (c *Ctx) AcceptsEncodings(offers ...string) string
func (c *Ctx) AcceptsLanguages(offers ...string) string
```



Example
示例

```go
// Accept: text/html, application/json; q=0.8, text/plain; q=0.5; charset="utf-8"

app.Get("/", func(c *fiber.Ctx) error {
  c.Accepts("html")             // "html"
  c.Accepts("text/html")        // "text/html"
  c.Accepts("json", "text")     // "json"
  c.Accepts("application/json") // "application/json"
  c.Accepts("text/plain", "application/json") // "application/json", due to quality
  c.Accepts("image/png")        // ""
  c.Accepts("png")              // ""
  // ...
})
```



Example 2
示例 2

```go
// Accept: text/html, text/*, application/json, */*; q=0

app.Get("/", func(c *fiber.Ctx) error {
  c.Accepts("text/plain", "application/json") // "application/json", due to specificity
  c.Accepts("application/json", "text/html") // "text/html", due to first match
  c.Accepts("image/png")        // "", due to */* without q factor 0 is Not Acceptable
  // ...
})
```



Media-Type parameters are supported.

​	支持媒体类型参数。

Example 3
示例 3

```go
// Accept: text/plain, application/json; version=1; foo=bar

app.Get("/", func(c *fiber.Ctx) error {
  // Extra parameters in the accept are ignored
  c.Accepts("text/plain;format=flowed") // "text/plain;format=flowed"
  
  // An offer must contain all parameters present in the Accept type
  c.Accepts("application/json") // ""

  // Parameter order and capitalization does not matter. Quotes on values are stripped.
  c.Accepts(`application/json;foo="bar";VERSION=1`) // "application/json;foo="bar";VERSION=1"
})
```



Example 4
示例 4

```go
// Accept: text/plain;format=flowed;q=0.9, text/plain
// i.e., "I prefer text/plain;format=flowed less than other forms of text/plain"
app.Get("/", func(c *fiber.Ctx) error {
  // Beware: the order in which offers are listed matters.
  // Although the client specified they prefer not to receive format=flowed,
  // the text/plain Accept matches with "text/plain;format=flowed" first, so it is returned.
  c.Accepts("text/plain;format=flowed", "text/plain") // "text/plain;format=flowed"

  // Here, things behave as expected:
  c.Accepts("text/plain", "text/plain;format=flowed") // "text/plain"
})
```



Fiber provides similar functions for the other accept headers.

​	Fiber 为其他接受头提供类似的功能。

```go
// Accept-Charset: utf-8, iso-8859-1;q=0.2
// Accept-Encoding: gzip, compress;q=0.2
// Accept-Language: en;q=0.8, nl, ru

app.Get("/", func(c *fiber.Ctx) error {
  c.AcceptsCharsets("utf-16", "iso-8859-1")
  // "iso-8859-1"

  c.AcceptsEncodings("compress", "br")
  // "compress"

  c.AcceptsLanguages("pt", "nl", "ru")
  // "nl"
  // ...
})
```



## AllParams

Params is used to get all route parameters. Using Params method to get params.

​	Params 用于获取所有路由参数。使用 Params 方法获取参数。

Signature
签名

```go
func (c *Ctx) AllParams() map[string]string
```



Example
示例

```go
// GET http://example.com/user/fenny
app.Get("/user/:name", func(c *fiber.Ctx) error {
  c.AllParams() // "{"name": "fenny"}"

  // ...
})

// GET http://example.com/user/fenny/123
app.Get("/user/*", func(c *fiber.Ctx) error {
  c.AllParams()  // "{"*1": "fenny/123"}"

  // ...
})
```



## App App

Returns the [*App]({{< ref "/fiber/API/Ctx" >}}) reference so you could easily access all application settings.

​	返回 *App 引用，以便您可以轻松访问所有应用程序设置。

Signature
签名

```go
func (c *Ctx) App() *App
```



Example
示例

```go
app.Get("/stack", func(c *fiber.Ctx) error {
  return c.JSON(c.App().Stack())
})
```



## Append 追加

Appends the specified **value** to the HTTP response header field.

​	将指定值追加到 HTTP 响应头字段。

CAUTION
注意

If the header is **not** already set, it creates the header with the specified value.

​	如果尚未设置头，则使用指定值创建头。

Signature
签名

```go
func (c *Ctx) Append(field string, values ...string)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.Append("Link", "http://google.com", "http://localhost")
  // => Link: http://localhost, http://google.com

  c.Append("Link", "Test")
  // => Link: http://localhost, http://google.com, Test

  // ...
})
```



## Attachment 附件

Sets the HTTP response [Content-Disposition](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition) header field to `attachment`.

​	将 HTTP 响应 Content-Disposition 头字段设置为 `attachment` 。

Signature
签名

```go
func (c *Ctx) Attachment(filename ...string)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.Attachment()
  // => Content-Disposition: attachment

  c.Attachment("./upload/images/logo.png")
  // => Content-Disposition: attachment; filename="logo.png"
  // => Content-Type: image/png

  // ...
})
```



## BaseURL 基本网址

Returns the base URL (**protocol** + **host**) as a `string`.

​	将基本网址（协议 + 主机）作为 `string` 返回。

Signature
签名

```go
func (c *Ctx) BaseURL() string
```



Example
示例

```go
// GET https://example.com/page#chapter-1

app.Get("/", func(c *fiber.Ctx) error {
  c.BaseURL() // https://example.com
  // ...
})
```



## Bind 绑定

Add vars to default view var map binding to template engine. Variables are read by the Render method and may be overwritten.

​	将变量添加到默认视图变量映射以绑定到模板引擎。变量由 Render 方法读取，并且可能会被覆盖。

Signature
签名

```go
func (c *Ctx) Bind(vars Map) error
```



Example
示例

```go
app.Use(func(c *fiber.Ctx) error {
  c.Bind(fiber.Map{
    "Title": "Hello, World!",
  })
})

app.Get("/", func(c *fiber.Ctx) error {
  return c.Render("xxx.tmpl", fiber.Map{}) // Render will use Title variable
})
```



## BodyRaw

Returns the raw request **body**.

​	返回原始请求正文。

Signature
签名

```go
func (c *Ctx) BodyRaw() []byte
```



Example
示例

```go
// curl -X POST http://localhost:8080 -d user=john

app.Post("/", func(c *fiber.Ctx) error {
  // Get raw body from POST request:
  return c.Send(c.BodyRaw()) // []byte("user=john")
})
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回值仅在处理程序内有效。不要存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置。* [*Read more...
> 了解更多...*](https://docs.gofiber.io/#zero-allocation)

## Body

As per the header `Content-Encoding`, this method will try to perform a file decompression from the **body** bytes. In case no `Content-Encoding` header is sent, it will perform as [BodyRaw](https://docs.gofiber.io/api/ctx/#bodyraw).

​	根据标头 `Content-Encoding` ，此方法将尝试对正文字节执行文件解压缩。如果没有发送 `Content-Encoding` 标头，它将执行 BodyRaw。

Signature
签名

```go
func (c *Ctx) Body() []byte
```



Example
示例

```go
// echo 'user=john' | gzip | curl -v -i --data-binary @- -H "Content-Encoding: gzip" http://localhost:8080

app.Post("/", func(c *fiber.Ctx) error {
  // Decompress body from POST request based on the Content-Encoding and return the raw content:
  return c.Send(c.Body()) // []byte("user=john")
})
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回值仅在处理程序内有效。不要存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## BodyParser

Binds the request body to a struct.

​	将请求正文绑定到结构。

It is important to specify the correct struct tag based on the content type to be parsed. For example, if you want to parse a JSON body with a field called Pass, you would use a struct field of `json:"pass"`.

​	根据要解析的内容类型指定正确的结构标记非常重要。例如，如果您想使用名为 Pass 的字段解析 JSON 正文，则可以使用 `json:"pass"` 的结构字段。

| content-type 内容类型               | struct tag 结构标记 |
| ----------------------------------- | ------------------- |
| `application/x-www-form-urlencoded` | form 表单           |
| `multipart/form-data`               | form 表单           |
| `application/json`                  | json                |
| `application/xml`                   | xml                 |
| `text/xml`                          | xml                 |

Signature
签名

```go
func (c *Ctx) BodyParser(out interface{}) error
```



Example
示例

```go
// Field names should start with an uppercase letter
type Person struct {
    Name string `json:"name" xml:"name" form:"name"`
    Pass string `json:"pass" xml:"pass" form:"pass"`
}

app.Post("/", func(c *fiber.Ctx) error {
        p := new(Person)

        if err := c.BodyParser(p); err != nil {
            return err
        }

        log.Println(p.Name) // john
        log.Println(p.Pass) // doe

        // ...
})

// Run tests with the following curl commands

// curl -X POST -H "Content-Type: application/json" --data "{\"name\":\"john\",\"pass\":\"doe\"}" localhost:3000

// curl -X POST -H "Content-Type: application/xml" --data "<login><name>john</name><pass>doe</pass></login>" localhost:3000

// curl -X POST -H "Content-Type: application/x-www-form-urlencoded" --data "name=john&pass=doe" localhost:3000

// curl -X POST -F name=john -F pass=doe http://localhost:3000

// curl -X POST "http://localhost:3000/?name=john&pass=doe"
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回值仅在处理程序中有效。请勿存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## ClearCookie

Expire a client cookie (*or all cookies if left empty)*

​	使客户端 cookie 过期（如果留空，则使所有 cookie 过期）

Signature
签名

```go
func (c *Ctx) ClearCookie(key ...string)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  // Clears all cookies:
  c.ClearCookie()

  // Expire specific cookie by name:
  c.ClearCookie("user")

  // Expire multiple cookies by names:
  c.ClearCookie("token", "session", "track_id", "version")
  // ...
})
```



CAUTION
注意

Web browsers and other compliant clients will only clear the cookie if the given options are identical to those when creating the cookie, excluding expires and maxAge. ClearCookie will not set these values for you - a technique similar to the one shown below should be used to ensure your cookie is deleted.

​	只有在给定的选项与创建 cookie 时的选项（不包括 expires 和 maxAge）相同的情况下，Web 浏览器和其他兼容客户端才会清除 cookie。ClearCookie 不会为您设置这些值 - 应使用类似于下面所示的技术来确保您的 cookie 已删除。

Example
示例

```go
app.Get("/set", func(c *fiber.Ctx) error {
    c.Cookie(&fiber.Cookie{
        Name:     "token",
        Value:    "randomvalue",
        Expires:  time.Now().Add(24 * time.Hour),
        HTTPOnly: true,
        SameSite: "lax",
    })

    // ...
})

app.Get("/delete", func(c *fiber.Ctx) error {
    c.Cookie(&fiber.Cookie{
        Name:     "token",
        // Set expiry date to the past
        Expires:  time.Now().Add(-(time.Hour * 2)),
        HTTPOnly: true,
        SameSite: "lax",
    })

    // ...
})
```



## ClientHelloInfo

ClientHelloInfo contains information from a ClientHello message in order to guide application logic in the GetCertificate and GetConfigForClient callbacks. You can refer to the [ClientHelloInfo](https://golang.org/pkg/crypto/tls/#ClientHelloInfo) struct documentation for more information on the returned struct.

​	ClientHelloInfo 包含 ClientHello 消息中的信息，以便在 GetCertificate 和 GetConfigForClient 回调中指导应用程序逻辑。您可以参考 ClientHelloInfo 结构文档以获取有关返回结构的更多信息。

Signature
签名

```go
func (c *Ctx) ClientHelloInfo() *tls.ClientHelloInfo
```



Example
示例

```go
// GET http://example.com/hello
app.Get("/hello", func(c *fiber.Ctx) error {
  chi := c.ClientHelloInfo()
  // ...
})
```



## Context

Returns [*fasthttp.RequestCtx](https://godoc.org/github.com/valyala/fasthttp#RequestCtx) that is compatible with the context.Context interface that requires a deadline, a cancellation signal, and other values across API boundaries.

​	返回与 context.Context 接口兼容的 *fasthttp.RequestCtx，该接口需要一个截止时间、一个取消信号以及跨 API 边界的其他值。

Signature
签名

```go
func (c *Ctx) Context() *fasthttp.RequestCtx
```



INFO
信息

Please read the [Fasthttp Documentation](https://pkg.go.dev/github.com/valyala/fasthttp?tab=doc) for more information.

​	请阅读 Fasthttp 文档了解更多信息。

## Cookie

Set cookie

​	设置 cookie

Signature
签名

```go
func (c *Ctx) Cookie(cookie *Cookie)
```



```go
type Cookie struct {
    Name        string    `json:"name"`
    Value       string    `json:"value"`
    Path        string    `json:"path"`
    Domain      string    `json:"domain"`
    MaxAge      int       `json:"max_age"`
    Expires     time.Time `json:"expires"`
    Secure      bool      `json:"secure"`
    HTTPOnly    bool      `json:"http_only"`
    SameSite    string    `json:"same_site"`
    SessionOnly bool      `json:"session_only"`
}
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  // Create cookie
  cookie := new(fiber.Cookie)
  cookie.Name = "john"
  cookie.Value = "doe"
  cookie.Expires = time.Now().Add(24 * time.Hour)

  // Set cookie
  c.Cookie(cookie)
  // ...
})
```



## CookieParser

This method is similar to [BodyParser]({{< ref "/fiber/API/Ctx#bodyparser" >}}), but for cookie parameters. It is important to use the struct tag "cookie". For example, if you want to parse a cookie with a field called Age, you would use a struct field of `cookie:"age"`.

​	此方法类似于 BodyParser，但适用于 cookie 参数。重要的是使用结构标记“cookie”。例如，如果您想解析一个名为 Age 的字段的 cookie，您将使用 `cookie:"age"` 的结构字段。

Signature
签名

```go
func (c *Ctx) CookieParser(out interface{}) error
```



Example
示例

```go
// Field names should start with an uppercase letter
type Person struct {
    Name     string  `cookie:"name"`
    Age      int     `cookie:"age"`
    Job      bool    `cookie:"job"`
}

app.Get("/", func(c *fiber.Ctx) error {
        p := new(Person)

        if err := c.CookieParser(p); err != nil {
            return err
        }

        log.Println(p.Name)     // Joseph
        log.Println(p.Age)      // 23
        log.Println(p.Job)      // true
})
// Run tests with the following curl command
// curl.exe --cookie "name=Joseph; age=23; job=true" http://localhost:8000/
```



## Cookies

Get cookie value by key, you could pass an optional default value that will be returned if the cookie key does not exist.

​	按键获取 cookie 值，您可以传递一个可选的默认值，如果 cookie 键不存在，则将返回该默认值。

Signature
签名

```go
func (c *Ctx) Cookies(key string, defaultValue ...string) string
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  // Get cookie by key:
  c.Cookies("name")         // "john"
  c.Cookies("empty", "doe") // "doe"
  // ...
})
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回值仅在处理程序中有效。不要存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## Download 下载 

Transfers the file from path as an `attachment`.

​	将文件从路径作为 `attachment` 传输。

Typically, browsers will prompt the user to download. By default, the [Content-Disposition](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition) header `filename=` parameter is the file path (*this typically appears in the browser dialog*).

​	通常，浏览器会提示用户下载。默认情况下，Content-Disposition 头 `filename=` 参数是文件路径（这通常显示在浏览器对话框中）。

Override this default with the **filename** parameter.

​	使用 filename 参数覆盖此默认设置。

Signature
签名

```go
func (c *Ctx) Download(file string, filename ...string) error
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  return c.Download("./files/report-12345.pdf");
  // => Download report-12345.pdf

  return c.Download("./files/report-12345.pdf", "report.pdf");
  // => Download report.pdf
})
```



## Format 格式 

Performs content-negotiation on the [Accept](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept) HTTP header. It uses [Accepts]({{< ref "/fiber/API/Ctx#accepts" >}}) to select a proper format.

​	对 Accept HTTP 头执行内容协商。它使用 Accepts 选择适当的格式。

INFO
信息

If the header is **not** specified or there is **no** proper format, **text/plain** is used.

​	如果未指定头或没有适当的格式，则使用 text/plain。

Signature
签名

```go
func (c *Ctx) Format(body interface{}) error
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  // Accept: text/plain
  c.Format("Hello, World!")
  // => Hello, World!

  // Accept: text/html
  c.Format("Hello, World!")
  // => <p>Hello, World!</p>

  // Accept: application/json
  c.Format("Hello, World!")
  // => "Hello, World!"
  // ..
})
```



## FormFile

MultipartForm files can be retrieved by name, the **first** file from the given key is returned.

​	可以通过名称检索 MultipartForm 文件，将返回给定键的第一个文件。

Signature
签名

```go
func (c *Ctx) FormFile(key string) (*multipart.FileHeader, error)
```



Example
示例

```go
app.Post("/", func(c *fiber.Ctx) error {
  // Get first file from form field "document":
  file, err := c.FormFile("document")

  // Save file to root directory:
  return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
})
```



## FormValue

Any form values can be retrieved by name, the **first** value from the given key is returned.

​	可以通过名称检索任何表单值，将返回给定键的第一个值。

Signature
签名

```go
func (c *Ctx) FormValue(key string, defaultValue ...string) string
```



Example
示例

```go
app.Post("/", func(c *fiber.Ctx) error {
  // Get first value from form field "name":
  c.FormValue("name")
  // => "john" or "" if not exist

  // ..
})
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回的值仅在处理程序中有效。不要存储任何引用。
> Make copies or use the
> 改用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置进行复制或使用。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## Fresh

When the response is still **fresh** in the client's cache **true** is returned, otherwise **false** is returned to indicate that the client cache is now stale and the full response should be sent.

​	当响应在客户端缓存中仍然新鲜时，将返回 true，否则将返回 false 以指示客户端缓存现在已过期，应发送完整响应。

When a client sends the Cache-Control: no-cache request header to indicate an end-to-end reload request, `Fresh` will return false to make handling these requests transparent.

​	当客户端发送 Cache-Control: no-cache 请求头以指示端到端重新加载请求时， `Fresh` 将返回 false 以使处理这些请求变得透明。

Read more on https://expressjs.com/en/4x/api.html#req.fresh

​	阅读更多内容，网址为 https://expressjs.com/en/4x/api.html#req.fresh

Signature
签名

```go
func (c *Ctx) Fresh() bool
```



## Get

Returns the HTTP request header specified by the field.

​	返回字段指定的 HTTP 请求头。

TIP

The match is **case-insensitive**.

​	匹配不区分大小写。

Signature
签名

```go
func (c *Ctx) Get(key string, defaultValue ...string) string
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.Get("Content-Type")       // "text/plain"
  c.Get("CoNtEnT-TypE")       // "text/plain"
  c.Get("something", "john")  // "john"
  // ..
})
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回值仅在处理程序内有效。请勿存储任何引用。
> Make copies or use the
> 改用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置进行复制或使用。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## GetReqHeaders

Returns the HTTP request headers as a map. Since a header can be set multiple times in a single request, the values of the map are slices of strings containing all the different values of the header.

​	将 HTTP 请求头作为映射返回。由于可以在单个请求中多次设置头，因此映射的值是包含头的所有不同值的字符串切片。

Signature
签名

```go
func (c *Ctx) GetReqHeaders() map[string][]string
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回值仅在处理程序内有效。请勿存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## GetRespHeader

Returns the HTTP response header specified by the field.

​	返回字段指定的 HTTP 响应头。

TIP

The match is **case-insensitive**.

​	匹配不区分大小写。

Signature
签名

```go
func (c *Ctx) GetRespHeader(key string, defaultValue ...string) string
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.GetRespHeader("X-Request-Id")       // "8d7ad5e3-aaf3-450b-a241-2beb887efd54"
  c.GetRespHeader("Content-Type")       // "text/plain"
  c.GetRespHeader("something", "john")  // "john"
  // ..
})
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回的值仅在处理程序中有效。不要存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## GetRespHeaders

Returns the HTTP response headers as a map. Since a header can be set multiple times in a single request, the values of the map are slices of strings containing all the different values of the header.

​	以映射的形式返回 HTTP 响应头。由于在单个请求中可以多次设置头，因此映射的值是包含头的所有不同值的字符串切片。

Signature
签名

```go
func (c *Ctx) GetRespHeaders() map[string][]string
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回的值仅在处理程序中有效。不要存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## GetRouteURL

Generates URLs to named routes, with parameters. URLs are relative, for example: "/user/1831"

​	生成带参数的命名路由的 URL。URL 是相对的，例如：“/user/1831”

Signature
签名

```go
func (c *Ctx) GetRouteURL(routeName string, params Map) (string, error)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Home page")
}).Name("home")

app.Get("/user/:id", func(c *fiber.Ctx) error {
    return c.SendString(c.Params("id"))
}).Name("user.show")

app.Get("/test", func(c *fiber.Ctx) error {
    location, _ := c.GetRouteURL("user.show", fiber.Map{"id": 1})
    return c.SendString(location)
})

// /test returns "/user/1"
```



## Hostname 主机名 

Returns the hostname derived from the [Host](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Host) HTTP header.

​	返回从 Host HTTP 标头派生的主机名。

Signature
签名

```go
func (c *Ctx) Hostname() string
```



Example
示例

```go
// GET http://google.com/search

app.Get("/", func(c *fiber.Ctx) error {
  c.Hostname() // "google.com"

  // ...
})
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回值仅在处理程序中有效。不要存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## IP IP

Returns the remote IP address of the request.

​	返回请求的远程 IP 地址。

Signature
签名

```go
func (c *Ctx) IP() string
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.IP() // "127.0.0.1"

  // ...
})
```



When registering the proxy request header in the fiber app, the ip address of the header is returned [(Fiber configuration)]({{< ref "/fiber/API/Fiber#config" >}})

​	在 Fiber 应用中注册代理请求头时，将返回头部的 IP 地址（Fiber 配置）

```go
app := fiber.New(fiber.Config{
  ProxyHeader: fiber.HeaderXForwardedFor,
})
```



## IPs IP

Returns an array of IP addresses specified in the [X-Forwarded-For](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For) request header.

​	返回 X-Forwarded-For 请求头中指定的一系列 IP 地址。

Signature
签名

```go
func (c *Ctx) IPs() []string
```



Example
示例

```go
// X-Forwarded-For: proxy1, 127.0.0.1, proxy3

app.Get("/", func(c *fiber.Ctx) error {
  c.IPs() // ["proxy1", "127.0.0.1", "proxy3"]

  // ...
})
```



CAUTION
注意

Improper use of the X-Forwarded-For header can be a security risk. For details, see the [Security and privacy concerns](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For#security_and_privacy_concerns) section.

​	不当使用 X-Forwarded-For 头部可能存在安全风险。有关详细信息，请参阅安全和隐私问题部分。

## Is IsMime

Returns the matching **content type**, if the incoming request’s [Content-Type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type) HTTP header field matches the [MIME type](https://developer.mozilla.org/ru/docs/Web/HTTP/Basics_of_HTTP/MIME_types) specified by the type parameter.

​	如果传入请求的 Content-Type 头字段与 type 参数指定的内容类型匹配，则返回该内容类型。

INFO
信息

If the request has **no** body, it returns **false**.

​	如果请求没有内容，则返回 False。

Signature
签名

```go
func (c *Ctx) Is(extension string) bool
```



Example
示例

```go
// Content-Type: text/html; charset=utf-8

app.Get("/", func(c *fiber.Ctx) error {
  c.Is("html")  // true
  c.Is(".html") // true
  c.Is("json")  // false

  // ...
})
```



## IsFromLocal IsFromLocal

Returns true if request came from localhost

​	如果请求来自

Signature
签名

```go
func (c *Ctx) IsFromLocal() bool {
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  // If request came from localhost, return true else return false
  c.IsFromLocal()

  // ...
})
```



## JSON JSON

Converts any **interface** or **string** to JSON using the [encoding/json](https://pkg.go.dev/encoding/json) package.

​	使用 /json 包将任何接口或字符串转换为 JSON。

INFO
信息

JSON also sets the content header to the `ctype` parameter. If no `ctype` is passed in, the header is set to `application/json`.

​	JSON 还会将内容头设置为 `ctype` 参数。如果没有传入 `ctype` ，则将头设置为 `application/json` 。

Signature
签名

```go
func (c *Ctx) JSON(data interface{}, ctype ...string) error
```



Example
示例

```go
type SomeStruct struct {
  Name string
  Age  uint8
}

app.Get("/json", func(c *fiber.Ctx) error {
  // Create data struct:
  data := SomeStruct{
    Name: "Grame",
    Age:  20,
  }

  return c.JSON(data)
  // => Content-Type: application/json
  // => "{"Name": "Grame", "Age": 20}"

  return c.JSON(fiber.Map{
    "name": "Grame",
    "age": 20,
  })
  // => Content-Type: application/json
  // => "{"name": "Grame", "age": 20}"

  return c.JSON(fiber.Map{
    "type": "https://example.com/probs/out-of-credit",
    "title": "You do not have enough credit.",
    "status": 403,
    "detail": "Your current balance is 30, but that costs 50.",
    "instance": "/account/12345/msgs/abc",
  }, "application/problem+json")
  // => Content-Type: application/problem+json
  // => "{
  // =>     "type": "https://example.com/probs/out-of-credit",
  // =>     "title": "You do not have enough credit.",
  // =>     "status": 403,
  // =>     "detail": "Your current balance is 30, but that costs 50.",
  // =>     "instance": "/account/12345/msgs/abc",
  // => }"
})
```



## JSONP JSONP

Sends a JSON response with JSONP support. This method is identical to [JSON]({{< ref "/fiber/API/Ctx#json" >}}), except that it opts-in to JSONP callback support. By default, the callback name is simply callback.

​	发送带有 JSONP支持的 JSON 响应。此方法与 JSON 相同，但它启用了 JSONP

Override this by passing a **named string** in the method.

​	通过在方法中传递一个命名字符串来覆盖此项。

Signature
签名

```go
func (c *Ctx) JSONP(data interface{}, callback ...string) error
```



Example
示例

```go
type SomeStruct struct {
  name string
  age  uint8
}

app.Get("/", func(c *fiber.Ctx) error {
  // Create data struct:
  data := SomeStruct{
    name: "Grame",
    age:  20,
  }

  return c.JSONP(data)
  // => callback({"name": "Grame", "age": 20})

  return c.JSONP(data, "customFunc")
  // => customFunc({"name": "Grame", "age": 20})
})
```



## Links 链接 

Joins the links followed by the property to populate the response’s [Link](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Link) HTTP header field.

​	连接属性后面的链接以填充响应的 Link HTTP 标头字段。

Signature
签名

```go
func (c *Ctx) Links(link ...string)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.Links(
    "http://api.example.com/users?page=2", "next",
    "http://api.example.com/users?page=5", "last",
  )
  // Link: <http://api.example.com/users?page=2>; rel="next",
  //       <http://api.example.com/users?page=5>; rel="last"

  // ...
})
```



## Locals 本地 

A method that stores variables scoped to the request and, therefore, are available only to the routes that match the request.

​	一种存储变量的方法，这些变量的作用域限定为请求，因此仅对与请求匹配的路由可用。

TIP

This is useful if you want to pass some **specific** data to the next middleware.

​	如果您想将一些特定数据传递给下一个中间件，这将非常有用。

Signature
签名

```go
func (c *Ctx) Locals(key interface{}, value ...interface{}) interface{}
```



Example
示例

```go
app.Use(func(c *fiber.Ctx) error {
  c.Locals("user", "admin")
  return c.Next()
})

app.Get("/admin", func(c *fiber.Ctx) error {
  if c.Locals("user") == "admin" {
    return c.Status(fiber.StatusOK).SendString("Welcome, admin!")
  }
  return c.SendStatus(fiber.StatusForbidden)

})
```



## Location 位置 

Sets the response [Location](https://developer.mozilla.org/ru/docs/Web/HTTP/Headers/Location) HTTP header to the specified path parameter.

​	将响应 Location HTTP 标头设置为指定的路径参数。

Signature
签名

```go
func (c *Ctx) Location(path string)
```



Example
示例

```go
app.Post("/", func(c *fiber.Ctx) error {
  c.Location("http://example.com")

  c.Location("/foo/bar")

  return nil
})
```



## Method 方法 

Returns a string corresponding to the HTTP method of the request: `GET`, `POST`, `PUT`, and so on.

​	返回一个字符串，该字符串对应于请求的 HTTP 方法： `GET` 、 `POST` 、 `PUT` 等。
Optionally, you could override the method by passing a string.

​	您可以选择通过传递字符串来覆盖该方法。

Signature
签名

```go
func (c *Ctx) Method(override ...string) string
```



Example
示例

```go
app.Post("/", func(c *fiber.Ctx) error {
  c.Method() // "POST"

  c.Method("GET")
  c.Method() // GET

  // ...
})
```



## MultipartForm

To access multipart form entries, you can parse the binary with `MultipartForm()`. This returns a `map[string][]string`, so given a key, the value will be a string slice.

​	要访问多部分表单条目，您可以使用 `MultipartForm()` 解析二进制文件。这会返回一个 `map[string][]string` ，因此给定一个键，值将是一个字符串切片。

Signature
签名

```go
func (c *Ctx) MultipartForm() (*multipart.Form, error)
```



Example
示例

```go
app.Post("/", func(c *fiber.Ctx) error {
  // Parse the multipart form:
  if form, err := c.MultipartForm(); err == nil {
    // => *multipart.Form

    if token := form.Value["token"]; len(token) > 0 {
      // Get key value:
      fmt.Println(token[0])
    }

    // Get all files from "documents" key:
    files := form.File["documents"]
    // => []*multipart.FileHeader

    // Loop through files:
    for _, file := range files {
      fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
      // => "tutorial.pdf" 360641 "application/pdf"

      // Save the files to disk:
      if err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); err != nil {
        return err
      }
    }
  }

  return err
})
```



## Next

When **Next** is called, it executes the next method in the stack that matches the current route. You can pass an error struct within the method that will end the chaining and call the [error handler]({{< ref "/fiber/Guide/ErrorHandling" >}}).

​	调用 Next 时，它会执行与当前路由匹配的堆栈中的下一个方法。您可以在方法中传递一个错误结构，该结构将结束链接并调用错误处理程序。

Signature
签名

```go
func (c *Ctx) Next() error
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  fmt.Println("1st route!")
  return c.Next()
})

app.Get("*", func(c *fiber.Ctx) error {
  fmt.Println("2nd route!")
  return c.Next()
})

app.Get("/", func(c *fiber.Ctx) error {
  fmt.Println("3rd route!")
  return c.SendString("Hello, World!")
})
```



## OriginalURL

Returns the original request URL.

​	返回原始请求 URL。

Signature
签名

```go
func (c *Ctx) OriginalURL() string
```



Example
示例

```go
// GET http://example.com/search?q=something

app.Get("/", func(c *fiber.Ctx) error {
  c.OriginalURL() // "/search?q=something"

  // ...
})
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回值仅在处理程序中有效。不要存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## Params 参数

Method can be used to get the route parameters, you could pass an optional default value that will be returned if the param key does not exist.

​	方法可用于获取路由参数，您可以传递一个可选的默认值，如果参数键不存在，则将返回该默认值。

INFO
信息

Defaults to empty string (`""`), if the param **doesn't** exist.

​	如果参数不存在，则默认为空字符串 ( `""` )。

Signature
签名

```go
func (c *Ctx) Params(key string, defaultValue ...string) string
```



Example
示例

```go
// GET http://example.com/user/fenny
app.Get("/user/:name", func(c *fiber.Ctx) error {
  c.Params("name") // "fenny"

  // ...
})

// GET http://example.com/user/fenny/123
app.Get("/user/*", func(c *fiber.Ctx) error {
  c.Params("*")  // "fenny/123"
  c.Params("*1") // "fenny/123"

  // ...
})
```



Unnamed route parameters(*, +) can be fetched by the **character** and the **counter** in the route.

​	可以通过路由中的字符和计数器获取未命名的路由参数(*, +)。

Example
示例

```go
// ROUTE: /v1/*/shop/*
// GET:   /v1/brand/4/shop/blue/xs
c.Params("*1")  // "brand/4"
c.Params("*2")  // "blue/xs"
```



For reasons of **downward compatibility**, the first parameter segment for the parameter character can also be accessed without the counter.

​	出于向下兼容性的原因，也可以在没有计数器的情况下访问参数字符的第一个参数段。

Example
示例

```go
app.Get("/v1/*/shop/*", func(c *fiber.Ctx) error {
  c.Params("*") // outputs the values of the first wildcard segment
})
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回值仅在处理程序中有效。不要存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置代替。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## ParamsInt

Method can be used to get an integer from the route parameters. Please note if that parameter is not in the request, zero will be returned. If the parameter is NOT a number, zero and an error will be returned

​	方法可用于从路由参数中获取整数。请注意，如果请求中没有该参数，则会返回零。如果参数不是数字，则会返回零和错误

INFO
信息

Defaults to the integer zero (`0`), if the param **doesn't** exist.

​	如果参数不存在，则默认为整数零 ( `0` )。

Signature
签名

```go
func (c *Ctx) ParamsInt(key string) (int, error)
```



Example
示例

```go
// GET http://example.com/user/123
app.Get("/user/:id", func(c *fiber.Ctx) error {
  id, err := c.ParamsInt("id") // int 123 and no error

  // ...
})
```



This method is equivalent of using `atoi` with ctx.Params

​	此方法相当于在 ctx.Params 中使用 `atoi`

## ParamsParser

This method is similar to BodyParser, but for path parameters. It is important to use the struct tag "params". For example, if you want to parse a path parameter with a field called Pass, you would use a struct field of params:"pass"

​	此方法类似于 BodyParser，但适用于路径参数。重要的是使用结构标记“params”。例如，如果您想使用名为 Pass 的字段解析路径参数，则可以使用 params:“pass” 的结构字段

Signature
签名

```go
func (c *Ctx) ParamsParser(out interface{}) error
```



Example
示例

```go
// GET http://example.com/user/111
app.Get("/user/:id", func(c *fiber.Ctx) error {
  param := struct {ID uint `params:"id"`}{}

  c.ParamsParser(&param) // "{"id": 111}"

  // ...
})
```



## Path

Contains the path part of the request URL. Optionally, you could override the path by passing a string. For internal redirects, you might want to call [RestartRouting]({{< ref "/fiber/API/Ctx#restartrouting" >}}) instead of [Next]({{< ref "/fiber/API/Ctx#next" >}}).

​	包含请求 URL 的路径部分。或者，您可以通过传递字符串来覆盖路径。对于内部重定向，您可能希望调用 RestartRouting 而不是 Next。

Signature
签名

```go
func (c *Ctx) Path(override ...string) string
```



Example
示例

```go
// GET http://example.com/users?sort=desc

app.Get("/users", func(c *fiber.Ctx) error {
  c.Path() // "/users"

  c.Path("/john")
  c.Path() // "/john"

  // ...
})
```



## Protocol

Contains the request protocol string: `http` or `https` for **TLS** requests.

​	包含请求协议字符串：对于 TLS 请求，为 `http` 或 `https` 。

Signature
签名

```go
func (c *Ctx) Protocol() string
```



Example
示例

```go
// GET http://example.com

app.Get("/", func(c *fiber.Ctx) error {
  c.Protocol() // "http"

  // ...
})
```



## Queries 查询 

Queries is a function that returns an object containing a property for each query string parameter in the route.

​	查询是一个函数，它返回一个对象，其中包含一个属性，用于路由中的每个查询字符串参数。

Signature
签名

```go
func (c *Ctx) Queries() map[string]string
```



Example
示例

```go
// GET http://example.com/?name=alex&want_pizza=false&id=

app.Get("/", func(c *fiber.Ctx) error {
    m := c.Queries()
    m["name"] // "alex"
    m["want_pizza"] // "false"
    m["id"] // ""
    // ...
})
```



Example
示例

```go
// GET http://example.com/?field1=value1&field1=value2&field2=value3

app.Get("/", func (c *fiber.Ctx) error {
    m := c.Queries()
    m["field1"] // "value2"
    m["field2"] // value3
})
```



Example
示例

```go
// GET http://example.com/?list_a=1&list_a=2&list_a=3&list_b[]=1&list_b[]=2&list_b[]=3&list_c=1,2,3

app.Get("/", func(c *fiber.Ctx) error {
    m := c.Queries()
    m["list_a"] // "3"
    m["list_b[]"] // "3"
    m["list_c"] // "1,2,3"
})
```



Example
示例

```go
// GET /api/posts?filters.author.name=John&filters.category.name=Technology

app.Get("/", func(c *fiber.Ctx) error {
    m := c.Queries()
    m["filters.author.name"] // John
    m["filters.category.name"] // Technology
})
```



Example
示例

```go
// GET /api/posts?tags=apple,orange,banana&filters[tags]=apple,orange,banana&filters[category][name]=fruits&filters.tags=apple,orange,banana&filters.category.name=fruits

app.Get("/", func(c *fiber.Ctx) error {
    m := c.Queries()
    m["tags"] // apple,orange,banana
    m["filters[tags]"] // apple,orange,banana
    m["filters[category][name]"] // fruits
    m["filters.tags"] // apple,orange,banana
    m["filters.category.name"] // fruits
})
```



## Query 查询 

This property is an object containing a property for each query string parameter in the route, you could pass an optional default value that will be returned if the query key does not exist.

​	此属性是一个对象，其中包含一个属性，用于路由中的每个查询字符串参数，您可以传递一个可选的默认值，如果查询键不存在，则将返回该值。

INFO
信息

If there is **no** query string, it returns an **empty string**.

​	如果没有查询字符串，则返回一个空字符串。

Signature
签名

```go
func (c *Ctx) Query(key string, defaultValue ...string) string
```



Example
示例

```go
// GET http://example.com/?order=desc&brand=nike

app.Get("/", func(c *fiber.Ctx) error {
  c.Query("order")         // "desc"
  c.Query("brand")         // "nike"
  c.Query("empty", "nike") // "nike"

  // ...
})
```



> *Returned value is only valid within the handler. Do not store any references.
> 返回的值仅在处理程序中有效。不要存储任何引用。
> Make copies or use the
> 制作副本或使用* [***`Immutable`***]({{< ref "/fiber/API/Ctx" >}}) *setting instead.
> 设置。* [*Read more...
> 阅读更多...*](https://docs.gofiber.io/#zero-allocation)

## QueryBool

This property is an object containing a property for each query boolean parameter in the route, you could pass an optional default value that will be returned if the query key does not exist.

​	此属性是一个对象，其中包含路由中每个查询布尔参数的属性，您可以传递一个可选的默认值，如果查询键不存在，则将返回该值。

CAUTION
注意

Please note if that parameter is not in the request, false will be returned. If the parameter is not a boolean, it is still tried to be converted and usually returned as false.

​	请注意，如果请求中没有该参数，则将返回 false。如果该参数不是布尔值，仍会尝试进行转换，通常返回 false。

Signature
签名

```go
func (c *Ctx) QueryBool(key string, defaultValue ...bool) bool
```



Example
示例

```go
// GET http://example.com/?name=alex&want_pizza=false&id=

app.Get("/", func(c *fiber.Ctx) error {
    c.QueryBool("want_pizza")           // false
    c.QueryBool("want_pizza", true) // false
    c.QueryBool("name")                 // false
    c.QueryBool("name", true)           // true
    c.QueryBool("id")                   // false
    c.QueryBool("id", true)             // true

  // ...
})
```



## QueryFloat

This property is an object containing a property for each query float64 parameter in the route, you could pass an optional default value that will be returned if the query key does not exist.

​	此属性是一个对象，其中包含路由中每个查询 float64 参数的属性，您可以传递一个可选的默认值，如果查询键不存在，则将返回该值。

CAUTION
注意

Please note if that parameter is not in the request, zero will be returned. If the parameter is not a number, it is still tried to be converted and usually returned as 1.

​	请注意，如果请求中没有该参数，则将返回零。如果该参数不是数字，仍会尝试进行转换，通常返回 1。

INFO
信息

Defaults to the float64 zero (`0`), if the param **doesn't** exist.

​	如果参数不存在，则默认为 float64 零 ( `0` )。

Signature
签名

```go
func (c *Ctx) QueryFloat(key string, defaultValue ...float64) float64
```



Example
示例

```go
// GET http://example.com/?name=alex&amount=32.23&id=

app.Get("/", func(c *fiber.Ctx) error {
    c.QueryFloat("amount")      // 32.23
    c.QueryFloat("amount", 3)   // 32.23
    c.QueryFloat("name", 1)     // 1
    c.QueryFloat("name")        // 0
    c.QueryFloat("id", 3)       // 3

  // ...
})
```



## QueryInt

This property is an object containing a property for each query integer parameter in the route, you could pass an optional default value that will be returned if the query key does not exist.

​	此属性是一个对象，其中包含路由中每个查询整数参数的属性，您可以传递一个可选的默认值，如果查询键不存在，则将返回该值。

CAUTION
注意

Please note if that parameter is not in the request, zero will be returned. If the parameter is not a number, it is still tried to be converted and usually returned as 1.

​	请注意，如果请求中没有该参数，则将返回零。如果该参数不是数字，仍会尝试进行转换，通常返回 1。

INFO
信息

Defaults to the integer zero (`0`), if the param **doesn't** exist.

​	如果参数不存在，则默认为整数零 ( `0` )。

Signature
签名

```go
func (c *Ctx) QueryInt(key string, defaultValue ...int) int
```



Example
示例

```go
// GET http://example.com/?name=alex&wanna_cake=2&id=

app.Get("/", func(c *fiber.Ctx) error {
    c.QueryInt("wanna_cake", 1) // 2
    c.QueryInt("name", 1)       // 1
    c.QueryInt("id", 1)         // 1
    c.QueryInt("id")            // 0

  // ...
})
```



## QueryParser

This method is similar to [BodyParser]({{< ref "/fiber/API/Ctx#bodyparser" >}}), but for query parameters. It is important to use the struct tag "query". For example, if you want to parse a query parameter with a field called Pass, you would use a struct field of `query:"pass"`.

​	此方法类似于 BodyParser，但适用于查询参数。重要的是使用结构标记“query”。例如，如果您想使用名为 Pass 的字段解析查询参数，则可以使用 `query:"pass"` 的结构字段。

Signature
签名

```go
func (c *Ctx) QueryParser(out interface{}) error
```



Example
示例

```go
// Field names should start with an uppercase letter
type Person struct {
    Name     string     `query:"name"`
    Pass     string     `query:"pass"`
    Products []string   `query:"products"`
}

app.Get("/", func(c *fiber.Ctx) error {
        p := new(Person)

        if err := c.QueryParser(p); err != nil {
            return err
        }

        log.Println(p.Name)     // john
        log.Println(p.Pass)     // doe
        log.Println(p.Products) // [shoe, hat]

        // ...
})
// Run tests with the following curl command

// curl "http://localhost:3000/?name=john&pass=doe&products=shoe,hat"
```



## Range

A struct containing the type and a slice of ranges will be returned.

​	将返回包含类型和范围切片的结构。

Signature
签名

```go
func (c *Ctx) Range(size int) (Range, error)
```



Example
示例

```go
// Range: bytes=500-700, 700-900
app.Get("/", func(c *fiber.Ctx) error {
  b := c.Range(1000)
  if b.Type == "bytes" {
      for r := range r.Ranges {
      fmt.Println(r)
      // [500, 700]
    }
  }
})
```



## Redirect

Redirects to the URL derived from the specified path, with specified status, a positive integer that corresponds to an HTTP status code.

​	重定向到从指定路径派生的 URL，具有指定的正整数状态，该正整数对应于 HTTP 状态代码。

INFO
信息

If **not** specified, status defaults to **302 Found**.

​	如果未指定，则状态默认为 302 Found。

Signature
签名

```go
func (c *Ctx) Redirect(location string, status ...int) error
```



Example
示例

```go
app.Get("/coffee", func(c *fiber.Ctx) error {
  return c.Redirect("/teapot")
})

app.Get("/teapot", func(c *fiber.Ctx) error {
  return c.Status(fiber.StatusTeapot).Send("🍵 short and stout 🍵")
})
```



More examples
更多示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  return c.Redirect("/foo/bar")
  return c.Redirect("../login")
  return c.Redirect("http://example.com")
  return c.Redirect("http://example.com", 301)
})
```



## RedirectToRoute

Redirects to the specific route along with the parameters and with specified status, a positive integer that corresponds to an HTTP status code.

​	重定向到特定路由以及参数，并具有指定状态，一个对应于 HTTP 状态代码的正整数。

INFO
信息

If **not** specified, status defaults to **302 Found**.

​	如果未指定，则状态默认为 302 Found。

INFO
信息

If you want to send queries to route, you must add **"queries"** key typed as **map[string]string** to params.

​	如果要将查询发送到路由，则必须将键入为 map[string]string 的“queries”键添加到 params。

Signature
签名

```go
func (c *Ctx) RedirectToRoute(routeName string, params fiber.Map, status ...int) error
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  // /user/fiber
  return c.RedirectToRoute("user", fiber.Map{
    "name": "fiber"
  })
})

app.Get("/with-queries", func(c *fiber.Ctx) error {
  // /user/fiber?data[0][name]=john&data[0][age]=10&test=doe
  return c.RedirectToRoute("user", fiber.Map{
    "name": "fiber",
    "queries": map[string]string{"data[0][name]": "john", "data[0][age]": "10", "test": "doe"},
  })
})

app.Get("/user/:name", func(c *fiber.Ctx) error {
  return c.SendString(c.Params("name"))
}).Name("user")
```



## RedirectBack

Redirects back to refer URL. It redirects to fallback URL if refer header doesn't exists, with specified status, a positive integer that corresponds to an HTTP status code.

​	重定向回引用 URL。如果引用头不存在，则重定向到后备 URL，并具有指定状态，一个对应于 HTTP 状态代码的正整数。

INFO
信息

If **not** specified, status defaults to **302 Found**.

​	如果未指定，则状态默认为 302 Found。

Signature
签名

```go
func (c *Ctx) RedirectBack(fallback string, status ...int) error
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  return c.SendString("Home page")
})
app.Get("/test", func(c *fiber.Ctx) error {
  c.Set("Content-Type", "text/html")
  return c.SendString(`<a href="/back">Back</a>`)
})

app.Get("/back", func(c *fiber.Ctx) error {
  return c.RedirectBack("/")
})
```



## Render

Renders a view with data and sends a `text/html` response. By default `Render` uses the default [**Go Template engine**](https://pkg.go.dev/html/template/). If you want to use another View engine, please take a look at our [**Template middleware**](https://docs.gofiber.io/template).

​	使用数据呈现视图并发送 `text/html` 响应。默认情况下， `Render` 使用默认的 Go Template 引擎。如果您想使用其他视图引擎，请查看我们的模板中间件。

Signature
签名

```go
func (c *Ctx) Render(name string, bind interface{}, layouts ...string) error
```



## Request

Request return the [*fasthttp.Request](https://godoc.org/github.com/valyala/fasthttp#Request) pointer

​	Request 返回 *fasthttp.Request 指针

Signature
签名

```go
func (c *Ctx) Request() *fasthttp.Request
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.Request().Header.Method()
  // => []byte("GET")
})
```



## ReqHeaderParser

This method is similar to [BodyParser]({{< ref "/fiber/API/Ctx#bodyparser" >}}), but for request headers. It is important to use the struct tag "reqHeader". For example, if you want to parse a request header with a field called Pass, you would use a struct field of `reqHeader:"pass"`.

​	此方法类似于 BodyParser，但适用于请求头。重要的是使用结构标记“reqHeader”。例如，如果您想解析一个带有名为 Pass 的字段的请求头，您将使用 `reqHeader:"pass"` 的结构字段。

Signature
签名

```go
func (c *Ctx) ReqHeaderParser(out interface{}) error
```



Example
示例

```go
// Field names should start with an uppercase letter
type Person struct {
    Name     string     `reqHeader:"name"`
    Pass     string     `reqHeader:"pass"`
    Products []string   `reqHeader:"products"`
}

app.Get("/", func(c *fiber.Ctx) error {
        p := new(Person)

        if err := c.ReqHeaderParser(p); err != nil {
            return err
        }

        log.Println(p.Name)     // john
        log.Println(p.Pass)     // doe
        log.Println(p.Products) // [shoe, hat]

        // ...
})
// Run tests with the following curl command

// curl "http://localhost:3000/" -H "name: john" -H "pass: doe" -H "products: shoe,hat"
```



## Response

Response return the [*fasthttp.Response](https://godoc.org/github.com/valyala/fasthttp#Response) pointer

​	Response 返回 *fasthttp.Response 指针

Signature
签名

```go
func (c *Ctx) Response() *fasthttp.Response
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.Response().BodyWriter().Write([]byte("Hello, World!"))
  // => "Hello, World!"
  return nil
})
```



## RestartRouting

Instead of executing the next method when calling [Next]({{< ref "/fiber/API/Ctx#next" >}}), **RestartRouting** restarts execution from the first method that matches the current route. This may be helpful after overriding the path, i. e. an internal redirect. Note that handlers might be executed again which could result in an infinite loop.

​	在调用 Next 时，RestartRouting 从匹配当前路由的第一个方法重新开始执行，而不是执行下一个方法。这在覆盖路径后可能会有所帮助，即内部重定向。请注意，处理程序可能会再次执行，这可能会导致无限循环。

Signature
签名

```go
func (c *Ctx) RestartRouting() error
```



Example
示例

```go
app.Get("/new", func(c *fiber.Ctx) error {
  return c.SendString("From /new")
})

app.Get("/old", func(c *fiber.Ctx) error {
  c.Path("/new")
  return c.RestartRouting()
})
```



## Route 路由 

Returns the matched [Route](https://pkg.go.dev/github.com/gofiber/fiber?tab=doc#Route) struct.

​	返回匹配的 Route 结构。

Signature
签名

```go
func (c *Ctx) Route() *Route
```



Example
示例

```go
// http://localhost:8080/hello


app.Get("/hello/:name", func(c *fiber.Ctx) error {
  r := c.Route()
  fmt.Println(r.Method, r.Path, r.Params, r.Handlers)
  // GET /hello/:name handler [name]

  // ...
})
```



CAUTION
注意

Do not rely on `c.Route()` in middlewares **before** calling `c.Next()` - `c.Route()` returns the **last executed route**.

​	在调用 `c.Next()` 之前，不要依赖中间件中的 `c.Route()` - `c.Route()` 返回最后执行的路由。

Example
示例

```go
func MyMiddleware() fiber.Handler {
  return func(c *fiber.Ctx) error {
    beforeNext := c.Route().Path // Will be '/'
    err := c.Next()
    afterNext := c.Route().Path // Will be '/hello/:name'
    return err
  }
}
```



## SaveFile

Method is used to save **any** multipart file to disk.

​	方法用于将任何多部分文件保存到磁盘。

Signature
签名

```go
func (c *Ctx) SaveFile(fh *multipart.FileHeader, path string) error
```



Example
示例

```go
app.Post("/", func(c *fiber.Ctx) error {
  // Parse the multipart form:
  if form, err := c.MultipartForm(); err == nil {
    // => *multipart.Form

    // Get all files from "documents" key:
    files := form.File["documents"]
    // => []*multipart.FileHeader

    // Loop through files:
    for _, file := range files {
      fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
      // => "tutorial.pdf" 360641 "application/pdf"

      // Save the files to disk:
      if err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); err != nil {
        return err
      }
    }
    return err
  }
})
```



## SaveFileToStorage

Method is used to save **any** multipart file to an external storage system.

​	方法用于将任何多部分文件保存到外部存储系统。

Signature
签名

```go
func (c *Ctx) SaveFileToStorage(fileheader *multipart.FileHeader, path string, storage Storage) error
```



Example
示例

```go
storage := memory.New()

app.Post("/", func(c *fiber.Ctx) error {
  // Parse the multipart form:
  if form, err := c.MultipartForm(); err == nil {
    // => *multipart.Form

    // Get all files from "documents" key:
    files := form.File["documents"]
    // => []*multipart.FileHeader

    // Loop through files:
    for _, file := range files {
      fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
      // => "tutorial.pdf" 360641 "application/pdf"

      // Save the files to storage:
      if err := c.SaveFileToStorage(file, fmt.Sprintf("./%s", file.Filename), storage); err != nil {
        return err
      }
    }
    return err
  }
})
```



## Secure

A boolean property that is `true` , if a **TLS** connection is established.

​	一个布尔属性，如果建立了 TLS 连接，则为 `true` 。

Signature
签名

```go
func (c *Ctx) Secure() bool
```



Example
示例

```go
// Secure() method is equivalent to:
c.Protocol() == "https"
```



## Send

Sets the HTTP response body.

​	设置 HTTP 响应主体。

Signature
签名

```go
func (c *Ctx) Send(body []byte) error
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  return c.Send([]byte("Hello, World!")) // => "Hello, World!"
})
```



Fiber also provides `SendString` and `SendStream` methods for raw inputs.

​	Fiber 还为原始输入提供 `SendString` 和 `SendStream` 方法。

TIP

Use this if you **don't need** type assertion, recommended for **faster** performance.

​	如果您不需要类型断言，请使用此方法，建议使用此方法以获得更快的性能。

Signature
签名

```go
func (c *Ctx) SendString(body string) error
func (c *Ctx) SendStream(stream io.Reader, size ...int) error
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  return c.SendString("Hello, World!")
  // => "Hello, World!"

  return c.SendStream(bytes.NewReader([]byte("Hello, World!")))
  // => "Hello, World!"
})
```



## SendFile

Transfers the file from the given path. Sets the [Content-Type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type) response HTTP header field based on the **filenames** extension.

​	从给定路径传输文件。根据文件名扩展名设置 Content-Type 响应 HTTP 头字段。

CAUTION
注意

Method doesn´t use **gzipping** by default, set it to **true** to enable.

​	方法默认不使用 gzip，将其设置为 true 以启用。

Signature
签名

```go
func (c *Ctx) SendFile(file string, compress ...bool) error
```



Example
示例

```go
app.Get("/not-found", func(c *fiber.Ctx) error {
  return c.SendFile("./public/404.html");

  // Disable compression
  return c.SendFile("./static/index.html", false);
})
```



INFO
信息

If the file contains an url specific character you have to escape it before passing the file path into the `sendFile` function.

​	如果文件包含特定于 url 的字符，则必须在将文件路径传递给 `sendFile` 函数之前对其进行转义。

Example
示例

```go
app.Get("/file-with-url-chars", func(c *fiber.Ctx) error {
  return c.SendFile(url.PathEscape("hash_sign_#.txt"))
})
```



INFO
信息

For sending files from embedded file system [this functionality]({{< ref "/fiber/API/Middleware/FileSystem#sendfile" >}}) can be used

​	可用于从嵌入式文件系统发送文件

## SendStatus

Sets the status code and the correct status message in the body, if the response body is **empty**.

​	如果响应体为空，则在正文中设置状态代码和正确状态消息。

TIP

You can find all used status codes and messages [here](https://github.com/gofiber/fiber/blob/dffab20bcdf4f3597d2c74633a7705a517d2c8c2/utils.go#L183-L244).

​	您可以在此处找到所有已使用状态代码和消息。

Signature
签名

```go
func (c *Ctx) SendStatus(status int) error
```



Example
示例

```go
app.Get("/not-found", func(c *fiber.Ctx) error {
  return c.SendStatus(415)
  // => 415 "Unsupported Media Type"

  c.SendString("Hello, World!")
  return c.SendStatus(415)
  // => 415 "Hello, World!"
})
```



## Set

Sets the response’s HTTP header field to the specified `key`, `value`.

​	将响应的 HTTP 标头字段设置为指定的 `key` 、 `value` 。

Signature
签名

```go
func (c *Ctx) Set(key string, val string)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.Set("Content-Type", "text/plain")
  // => "Content-type: text/plain"

  // ...
})
```



## SetParserDecoder

Allow you to config BodyParser/QueryParser decoder, base on schema's options, providing possibility to add custom type for parsing.

​	允许您配置 BodyParser/QueryParser 解码器，基于架构的选项，提供添加自定义类型以进行解析的可能性。

Signature
签名

```go
func SetParserDecoder(parserConfig fiber.ParserConfig{
  IgnoreUnknownKeys bool,
  ParserType        []fiber.ParserType{
      Customtype interface{},
      Converter  func(string) reflect.Value,
  },
  ZeroEmpty         bool,
  SetAliasTag       string,
})
```



Example
示例

```go
type CustomTime time.Time

// String() returns the time in string
func (ct *CustomTime) String() string {
    t := time.Time(*ct).String()
    return t
}

// Register the converter for CustomTime type format as 2006-01-02
var timeConverter = func(value string) reflect.Value {
  fmt.Println("timeConverter", value)
  if v, err := time.Parse("2006-01-02", value); err == nil {
    return reflect.ValueOf(v)
  }
  return reflect.Value{}
}

customTime := fiber.ParserType{
  Customtype: CustomTime{},
  Converter:  timeConverter,
}

// Add setting to the Decoder
fiber.SetParserDecoder(fiber.ParserConfig{
  IgnoreUnknownKeys: true,
  ParserType:        []fiber.ParserType{customTime},
  ZeroEmpty:         true,
})

// Example to use CustomType, you pause custom time format not in RFC3339
type Demo struct {
    Date  CustomTime `form:"date" query:"date"`
    Title string     `form:"title" query:"title"`
    Body  string     `form:"body" query:"body"`
}

app.Post("/body", func(c *fiber.Ctx) error {
    var d Demo
    c.BodyParser(&d)
    fmt.Println("d.Date", d.Date.String())
    return c.JSON(d)
})

app.Get("/query", func(c *fiber.Ctx) error {
    var d Demo
    c.QueryParser(&d)
    fmt.Println("d.Date", d.Date.String())
    return c.JSON(d)
})

// curl -X POST -F title=title -F body=body -F date=2021-10-20 http://localhost:3000/body

// curl -X GET "http://localhost:3000/query?title=title&body=body&date=2021-10-20"
```



## SetUserContext

Sets the user specified implementation for context interface.

​	设置上下文接口的用户指定实现。

Signature
签名

```go
func (c *Ctx) SetUserContext(ctx context.Context)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  ctx := context.Background()
  c.SetUserContext(ctx)
  // Here ctx could be any context implementation

  // ...
})
```



## Stale 陈旧 

https://expressjs.com/en/4x/api.html#req.stale

Signature
签名

```go
func (c *Ctx) Stale() bool
```



## Status 状态 

Sets the HTTP status for the response.

​	设置响应的 HTTP 状态。

INFO
信息

Method is a **chainable**.

​	方法是可链接的。

Signature
签名

```go
func (c *Ctx) Status(status int) *Ctx
```



Example
示例

```go
app.Get("/fiber", func(c *fiber.Ctx) error {
  c.Status(fiber.StatusOK)
  return nil
}

app.Get("/hello", func(c *fiber.Ctx) error {
  return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
}

app.Get("/world", func(c *fiber.Ctx) error {
  return c.Status(fiber.StatusNotFound).SendFile("./public/gopher.png")
})
```



## Subdomains 子域 

Returns a string slice of subdomains in the domain name of the request.

​	返回请求域名中的子域字符串切片。

The application property subdomain offset, which defaults to `2`, is used for determining the beginning of the subdomain segments.

​	应用程序属性子域偏移量，默认为 `2` ，用于确定子域段的开始。

Signature
签名

```go
func (c *Ctx) Subdomains(offset ...int) []string
```



Example
示例

```go
// Host: "tobi.ferrets.example.com"

app.Get("/", func(c *fiber.Ctx) error {
  c.Subdomains()  // ["ferrets", "tobi"]
  c.Subdomains(1) // ["tobi"]

  // ...
})
```



## Type 类型 

Sets the [Content-Type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type) HTTP header to the MIME type listed [here](https://github.com/nginx/nginx/blob/master/conf/mime.types) specified by the file **extension**.

​	将 Content-Type HTTP 标头设置为此处列出的由文件扩展名指定的 MIME 类型。

Signature
签名

```go
func (c *Ctx) Type(ext string, charset ...string) *Ctx
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.Type(".html") // => "text/html"
  c.Type("html")  // => "text/html"
  c.Type("png")   // => "image/png"

  c.Type("json", "utf-8")  // => "application/json; charset=utf-8"

  // ...
})
```



## UserContext

UserContext returns a context implementation that was set by user earlier or returns a non-nil, empty context, if it was not set earlier.

​	UserContext 返回由用户较早设置的上下文实现，或者如果未较早设置，则返回非零空上下文。

Signature
签名

```go
func (c *Ctx) UserContext() context.Context
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  ctx := c.UserContext()
  // ctx is context implementation set by user

  // ...
})
```



## Vary

Adds the given header field to the [Vary](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Vary) response header. This will append the header, if not already listed, otherwise leaves it listed in the current location.

​	将给定的标头字段添加到 Vary 响应标头。如果尚未列出，这将追加标头，否则将其保留在当前位置中列出。

INFO
信息

Multiple fields are **allowed**.

​	允许多个字段。

Signature
签名

```go
func (c *Ctx) Vary(fields ...string)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.Vary("Origin")     // => Vary: Origin
  c.Vary("User-Agent") // => Vary: Origin, User-Agent

  // No duplicates
  c.Vary("Origin") // => Vary: Origin, User-Agent

  c.Vary("Accept-Encoding", "Accept")
  // => Vary: Origin, User-Agent, Accept-Encoding, Accept

  // ...
})
```



## Write

Write adopts the Writer interface

​	Write 采用 Writer 接口

Signature
签名

```go
func (c *Ctx) Write(p []byte) (n int, err error)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.Write([]byte("Hello, World!")) // => "Hello, World!"

  fmt.Fprintf(c, "%s\n", "Hello, World!") // "Hello, World!Hello, World!"
})
```



## Writef

Writef adopts the string with variables

​	Writef 采用带变量的字符串

Signature
签名

```go
func (c *Ctx) Writef(f string, a ...interface{}) (n int, err error)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  world := "World!"
  c.Writef("Hello, %s", world) // => "Hello, World!"

  fmt.Fprintf(c, "%s\n", "Hello, World!") // "Hello, World!Hello, World!"
})
```



## WriteString

WriteString adopts the string

​	WriteString 采用字符串

Signature
签名

```go
func (c *Ctx) WriteString(s string) (n int, err error)
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
  c.WriteString("Hello, World!") // => "Hello, World!"

  fmt.Fprintf(c, "%s\n", "Hello, World!") // "Hello, World!Hello, World!"
})
```



## XHR

A Boolean property, that is `true`, if the request’s [X-Requested-With](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers) header field is [XMLHttpRequest](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest), indicating that the request was issued by a client library (such as [jQuery](https://api.jquery.com/jQuery.ajax/)).

​	一个布尔属性，即 `true` ，如果请求的 X-Requested-With 头字段是 XMLHttpRequest，则表示该请求是由客户端库（例如 jQuery）发出的。

Signature
签名

```go
func (c *Ctx) XHR() bool
```



Example
示例

```go
// X-Requested-With: XMLHttpRequest

app.Get("/", func(c *fiber.Ctx) error {
  c.XHR() // true

  // ...
})
```



## XML

Converts any **interface** or **string** to XML using the standard `encoding/xml` package.

​	使用标准 `encoding/xml` 包将任何接口或字符串转换为 XML。

INFO
信息

XML also sets the content header to **application/xml**.

​	XML 还将内容头设置为 application/xml。

Signature
签名

```go
func (c *Ctx) XML(data interface{}) error
```



Example
示例

```go
type SomeStruct struct {
  XMLName xml.Name `xml:"Fiber"`
  Name    string   `xml:"Name"`
  Age     uint8    `xml:"Age"`
}

app.Get("/", func(c *fiber.Ctx) error {
  // Create data struct:
  data := SomeStruct{
    Name: "Grame",
    Age:  20,
  }

  return c.XML(data)
  // <Fiber>
  //     <Name>Grame</Name>
  //    <Age>20</Age>
  // </Fiber>
})
```
