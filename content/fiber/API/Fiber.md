+++
title = "Fiber"
date = 2024-02-05T09:14:15+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/fiber]({{< ref "/fiber/API/Fiber" >}})

# 📦 Fiber

## New 新

This method creates a new **App** named instance. You can pass optional [config ](https://docs.gofiber.io/api/fiber/#config)when creating a new instance.

​	此方法创建一个名为实例的新应用。您可以在创建新实例时传递可选配置。

Signature
签名

```go
func New(config ...Config) *App
```



Example
示例

```go
// Default config
app := fiber.New()

// ...
```



## Config 配置

You can pass an optional Config when creating a new Fiber instance.

​	在创建新 Fiber 实例时，您可以传递一个可选的配置。

Example
示例

```go
// Custom config
app := fiber.New(fiber.Config{
    Prefork:       true,
    CaseSensitive: true,
    StrictRouting: true,
    ServerHeader:  "Fiber",
    AppName: "Test App v1.0.1",
})

// ...
```



**Config fields
配置字段**

| Property 属性                                                | Type 输入                                                    | Description 说明                                             | Default 默认                                                 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| AppName                                                      | `string`                                                     | This allows to setup app name for the app 这允许为应用设置应用名称 | `""`                                                         |
| BodyLimit                                                    | `int`                                                        | Sets the maximum allowed size for a request body, if the size exceeds the configured limit, it sends `413 - Request Entity Too Large` response. 设置请求正文允许的最大大小，如果大小超过配置的限制，它将发送 `413 - Request Entity Too Large` 响应。 | `4 * 1024 * 1024`                                            |
| CaseSensitive                                                | `bool`                                                       | When enabled, `/Foo` and `/foo` are different routes. When disabled, `/Foo`and `/foo` are treated the same. 启用时， `/Foo` 和 `/foo` 是不同的路由。禁用时， `/Foo` 和 `/foo` 被视为相同。 | `false`                                                      |
| ColorScheme                                                  | [`Colors`](https://github.com/gofiber/fiber/blob/master/color.go) | You can define custom color scheme. They'll be used for startup message, route list and some middlewares. 您可以定义自定义配色方案。它们将用于启动消息、路由列表和一些中间件。 | [`DefaultColors`](https://github.com/gofiber/fiber/blob/master/color.go) |
| CompressedFileSuffix                                         | `string`                                                     | Adds a suffix to the original file name and tries saving the resulting compressed file under the new file name. 向原始文件名添加后缀，并尝试以新文件名保存生成的压缩文件。 | `".fiber.gz"`                                                |
| Concurrency 并发性                                           | `int`                                                        | Maximum number of concurrent connections. 最大并发连接数。   | `256 * 1024`                                                 |
| DisableDefaultContentType                                    | `bool`                                                       | When set to true, causes the default Content-Type header to be excluded from the Response. 设置为 true 时，将导致默认 Content-Type 标头从响应中排除。 | `false`                                                      |
| DisableDefaultDate                                           | `bool`                                                       | When set to true causes the default date header to be excluded from the response. 设置为 true 时，将从响应中排除默认日期标头。 | `false`                                                      |
| DisableHeaderNormalizing                                     | `bool`                                                       | By default all header names are normalized: conteNT-tYPE -> Content-Type 默认情况下，所有标头名称都已规范化：conteNT-tYPE -> Content-Type | `false`                                                      |
| DisableKeepalive                                             | `bool`                                                       | Disable keep-alive connections, the server will close incoming connections after sending the first response to the client 禁用保持活动连接，服务器在向客户端发送第一个响应后将关闭传入连接 | `false`                                                      |
| DisablePreParseMultipartForm                                 | `bool`                                                       | Will not pre parse Multipart Form data if set to true. This option is useful for servers that desire to treat multipart form data as a binary blob, or choose when to parse the data. 如果设置为 true，则不会预先解析多部分表单数据。此选项对于希望将多部分表单数据视为二进制 blob 或选择何时解析数据的服务器很有用。 | `false`                                                      |
| DisableStartupMessage                                        | `bool`                                                       | When set to true, it will not print out debug information 设置为 true 时，它不会打印出调试信息 | `false`                                                      |
| ETag                                                         | `bool`                                                       | Enable or disable ETag header generation, since both weak and strong etags are generated using the same hashing method (CRC-32). Weak ETags are the default when enabled. 启用或禁用 ETag 头生成，因为弱 ETag 和强 ETag 都使用相同的哈希方法 (CRC-32) 生成。启用时，弱 ETag 为默认值。 | `false`                                                      |
| EnableIPValidation                                           | `bool`                                                       | If set to true, `c.IP()` and `c.IPs()` will validate IP addresses before returning them. Also, `c.IP()` will return only the first valid IP rather than just the raw header value that may be a comma seperated string. 如果设置为 true， `c.IP()` 和 `c.IPs()` 将在返回 IP 地址之前验证它们。此外， `c.IP()` 将仅返回第一个有效 IP，而不是可能为逗号分隔字符串的原始头值。  **WARNING:** There is a small performance cost to doing this validation. Keep disabled if speed is your only concern and your application is behind a trusted proxy that already validates this header. 警告：执行此验证会产生少量性能开销。如果速度是您唯一关注的问题，并且您的应用程序位于已验证此头的受信任代理之后，请保持禁用状态。 | `false`                                                      |
| EnablePrintRoutes                                            | `bool`                                                       | EnablePrintRoutes enables print all routes with their method, path, name and handler.. EnablePrintRoutes 启用打印所有具有其方法、路径、名称和处理程序的路由。 | `false`                                                      |
| EnableSplittingOnParsers                                     | `bool`                                                       | EnableSplittingOnParsers splits the query/body/header parameters by comma when it's true. 当 EnableSplittingOnParsers 为 true 时，它会按逗号拆分查询/正文/头参数。  For example, you can use it to parse multiple values from a query parameter like this: `/api?foo=bar,baz == foo[]=bar&foo[]=baz` 例如，您可以使用它来解析查询参数中的多个值，如下所示： `/api?foo=bar,baz == foo[]=bar&foo[]=baz` | `false`                                                      |
| EnableTrustedProxyCheck                                      | `bool`                                                       | When set to true, fiber will check whether proxy is trusted, using TrustedProxies list. 如果设置为 true，fiber 将使用 TrustedProxies 列表检查代理是否受信任。 默认情况下， 将从 X-Forwarded-Proto、X-Forwarded-Protocol、X-Forwarded-Ssl 或 X-Url-Scheme 标头获取值， 将从 标头获取值， 将从 X-Forwarded-Host 标头获取值。 如果 为 true，并且 在 、 和 的列表中，则在禁用 时，、 和 将具有相同的行为，如果 不在列表中，则 将在应用程序处理 tls 连接时返回 https，否则返回 http， 将从 fasthttp 上下文返回 RemoteIP()， 将返回 。 ErrorHandler 当从 fiber 返回错误时执行 ErrorHandler。由顶级应用程序保留已安装的 fiber 错误处理程序，并将其应用于前缀关联的请求。 GETOnly 如果设置为 true，则拒绝所有非 GET 请求。此选项对于仅接受 GET 请求的服务器来说非常有用，可以作为抗拒绝服务攻击的保护。如果设置了 GETOnly，则请求大小受 ReadBufferSize 限制。 IdleTimeout  By default `c.Protocol()` will get value from X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl or X-Url-Scheme header, `c.IP()` will get value from `ProxyHeader` header, `c.Hostname()` will get value from X-Forwarded-Host header. If `EnableTrustedProxyCheck` is true, and `RemoteIP` is in the list of `TrustedProxies` `c.Protocol()`, `c.IP()`, and `c.Hostname()` will have the same behaviour when `EnableTrustedProxyCheck` disabled, if `RemoteIP` isn't in the list, `c.Protocol()` will return https in case when tls connection is handled by the app, or http otherwise, `c.IP()` will return RemoteIP() from fasthttp context, `c.Hostname()` will return `fasthttp.Request.URI().Host()` | `false`                                                      |
| ErrorHandler                                                 | `ErrorHandler`                                               | ErrorHandler is executed when an error is returned from fiber.Handler. Mounted fiber error handlers are retained by the top-level app and applied on prefix associated requests. | `DefaultErrorHandler`                                        |
| GETOnly                                                      | `bool`                                                       | Rejects all non-GET requests if set to true. This option is useful as anti-DoS protection for servers accepting only GET requests. The request size is limited by ReadBufferSize if GETOnly is set. | `false`                                                      |
| IdleTimeout                                                  | `time.Duration`                                              | The maximum amount of time to wait for the next request when keep-alive is enabled. If IdleTimeout is zero, the value of ReadTimeout is used. 启用保持活动时，等待下一个请求的最大时间。如果 IdleTimeout 为零，则使用 ReadTimeout 的值。 | `nil`                                                        |
| Immutable 不可变                                             | `bool`                                                       | When enabled, all values returned by context methods are immutable. By default, they are valid until you return from the handler; see issue [#185](https://github.com/gofiber/fiber/issues/185). 启用后，上下文方法返回的所有值都是不可变的。默认情况下，它们在您从处理程序返回之前有效；请参阅问题 #185。 | `false`                                                      |
| JSONDecoder                                                  | `utils.JSONUnmarshal`                                        | Allowing for flexibility in using another json library for decoding. 允许灵活使用另一个 json 库进行解码。 | `json.Unmarshal`                                             |
| JSONEncoder                                                  | `utils.JSONMarshal`                                          | Allowing for flexibility in using another json library for encoding. 允许灵活使用另一个 json 库进行编码。 | `json.Marshal`                                               |
| Network 网络                                                 | `string`                                                     | Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only) 已知网络为“tcp”、“tcp4”（仅限 IPv4）、“tcp6”（仅限 IPv6）  **WARNING:** When prefork is set to true, only "tcp4" and "tcp6" can be chosen. 警告：当 prefork 设置为 true 时，只能选择“tcp4”和“tcp6”。 | `NetworkTCP4`                                                |
| PassLocalsToViews                                            | `bool`                                                       | PassLocalsToViews Enables passing of the locals set on a fiber.Ctx to the template engine. See our **Template Middleware** for supported engines. PassLocalsToViews 允许将 fiber.Ctx 上设置的本地变量传递给模板引擎。有关支持的引擎，请参阅我们的模板中间件。 | `false`                                                      |
| Prefork                                                      | `bool`                                                       | Enables use of the[`SO_REUSEPORT`](https://lwn.net/Articles/542629/)socket option. This will spawn multiple Go processes listening on the same port. learn more about [socket sharding](https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/). **NOTE: if enabled, the application will need to be ran through a shell because prefork mode sets environment variables. If you're using Docker, make sure the app is ran with `CMD ./app` or `CMD ["sh", "-c", "/app"]`. For more info, see** [**this**](https://github.com/gofiber/fiber/issues/1021#issuecomment-730537971) **issue comment.** 启用 `SO_REUSEPORT` 套接字选项。这将生成多个侦听同一端口的 Go 进程。详细了解套接字分片。注意：如果启用，则需要通过 shell 运行应用程序，因为 prefork 模式会设置环境变量。如果您使用的是 Docker，请确保使用 `CMD ./app` 或 `CMD ["sh", "-c", "/app"]` 运行该应用程序。有关更多信息，请参阅此问题评论。 | `false`                                                      |
| ProxyHeader                                                  | `string`                                                     | This will enable `c.IP()` to return the value of the given header key. By default `c.IP()`will return the Remote IP from the TCP connection, this property can be useful if you are behind a load balancer e.g. *X-Forwarded-**. 这将启用 `c.IP()` 以返回给定标头键的值。默认情况下， `c.IP()` 将从 TCP 连接返回远程 IP，如果您位于负载均衡器之后，此属性可能很有用，例如 X-Forwarded-*. | `""`                                                         |
| ReadBufferSize                                               | `int`                                                        | per-connection buffer size for requests' reading. This also limits the maximum header size. Increase this buffer if your clients send multi-KB RequestURIs and/or multi-KB headers (for example, BIG cookies). 用于读取请求的每个连接的缓冲区大小。这也限制了最大标头大小。如果您的客户端发送多 KB 的 RequestURI 和/或多 KB 的标头（例如，大型 Cookie），请增加此缓冲区。 | `4096`                                                       |
| ReadTimeout                                                  | `time.Duration`                                              | The amount of time allowed to read the full request, including the body. The default timeout is unlimited. 允许读取完整请求（包括正文）的时间量。默认超时时间不限。 | `nil`                                                        |
| RequestMethods                                               | `[]string`                                                   | RequestMethods provides customizibility for HTTP methods. You can add/remove methods as you wish. RequestMethods 提供了 HTTP 方法的可定制性。您可以根据需要添加/删除方法。 | `DefaultMethods`                                             |
| ServerHeader                                                 | `string`                                                     | Enables the `Server` HTTP header with the given value. 使用给定值启用 `Server` HTTP 标头。 | `""`                                                         |
| StreamRequestBody                                            | `bool`                                                       | StreamRequestBody enables request body streaming, and calls the handler sooner when given body is larger then the current limit. StreamRequestBody 启用请求正文流式传输，并在给定正文大于当前限制时尽早调用处理程序。 | `false`                                                      |
| StrictRouting                                                | `bool`                                                       | When enabled, the router treats `/foo` and `/foo/` as different. Otherwise, the router treats `/foo` and `/foo/` as the same. 启用后，路由器将 `/foo` 和 `/foo/` 视为不同。否则，路由器将 `/foo` 和 `/foo/` 视为相同。 | `false`                                                      |
| TrustedProxies                                               | `[]string`                                                   | Contains the list of trusted proxy IP's. Look at `EnableTrustedProxyCheck` doc. 包含受信任的代理 IP 列表。请参阅 `EnableTrustedProxyCheck` 文档。  It can take IP or IP range addresses. If it gets IP range, it iterates all possible addresses. 它可以获取 IP 或 IP 范围地址。如果它获取 IP 范围，它会迭代所有可能的地址。 | `[]string*__*`                                               |
| UnescapePath                                                 | `bool`                                                       | Converts all encoded characters in the route back before setting the path for the context, so that the routing can also work with URL encoded special characters 在为上下文设置路径之前，将路由中的所有编码字符转换回来，以便路由也可以使用 URL 编码的特殊字符 | `false`                                                      |
| Views 视图 Beego 的 MVC 简介 MVC 简介 MVC 简介 Beego 的 MVC 简介 Beego 使用典型的模型-视图-控制器 (MVC) 框架。此图说明了如何处理请求处理逻辑： 整个逻辑处理过程如下所述： 数据从侦听端口接收。侦听端口默认设置为 8080。 请求到达端口 8080 后，Beego 开始处理请求的数据 | `Views`                                                      | Views is the interface that wraps the Render function. See our **Template Middleware** for supported engines. Views 是包装 Render 函数的接口。有关支持的引擎，请参阅我们的模板中间件。 | `nil`                                                        |
| ViewsLayout                                                  | `string`                                                     | Views Layout is the global layout for all template render until override on Render function. See our **Template Middleware** for supported engines. Views Layout 是所有模板渲染的全局布局，直到在 Render 函数中覆盖。有关支持的引擎，请参阅我们的模板中间件。 | `""`                                                         |
| WriteBufferSize                                              | `int`                                                        | Per-connection buffer size for responses' writing. 用于响应写入的每个连接的缓冲区大小。 | `4096`                                                       |
| WriteTimeout                                                 | `time.Duration`                                              | The maximum duration before timing out writes of the response. The default timeout is unlimited. 响应写入超时之前的最长持续时间。默认超时时间不限。 | `nil`                                                        |
| XMLEncoder                                                   | `utils.XMLMarshal`                                           | Allowing for flexibility in using another XML library for encoding. 允许使用另一个 XML 库进行编码，从而提高灵活性。 | `xml.Marshal`                                                |

## NewError

NewError creates a new HTTPError instance with an optional message.

​	NewError 使用可选消息创建一个新的 HTTPError 实例。

Signature
签名

```go
func NewError(code int, message ...string) *Error
```



Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
    return fiber.NewError(782, "Custom error message")
})
```



## IsChild

IsChild determines if the current process is a result of Prefork.

​	IsChild 确定当前进程是否是 Prefork 的结果。

Signature
签名

```go
func IsChild() bool
```



Example
示例

```go
// Prefork will spawn child processes
app := fiber.New(fiber.Config{
    Prefork: true,
})

if !fiber.IsChild() {
    fmt.Println("I'm the parent process")
} else {
    fmt.Println("I'm a child process")
}

// ...
```
