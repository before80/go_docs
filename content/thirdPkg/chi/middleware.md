+++
title = "中间件"
date = 2024-01-31T19:06:28+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-chi.io/#/pages/middleware](https://go-chi.io/#/pages/middleware)

# Middleware 🧬 中间件

## Introduction 简介

> Middleware performs some specific function on the HTTP request or response at a specific stage in the HTTP pipeline before or after the user defined controller. Middleware is a design pattern to eloquently add cross cutting concerns like logging, handling authentication without having many code contact points.
>
> ​	中间件在 HTTP 管道中的特定阶段对 HTTP 请求或响应执行一些特定功能，该阶段位于用户定义的控制器之前或之后。中间件是一种设计模式，可以巧妙地添加跨切关注点，例如日志记录、处理身份验证，而无需许多代码接触点。

`chi's` middlewares are just stdlib net/http middleware handlers. There is nothing special about them, which means the router and all the tooling is designed to be compatible and friendly with any middleware in the community. This offers much better extensibility and reuse of packages and is at the heart of chi's purpose.

​	`chi's` 中间件只是 stdlib net/http 中间件处理程序。它们没有什么特别之处，这意味着路由器和所有工具都旨在与社区中的任何中间件兼容并友好。这提供了更好的可扩展性和软件包的重用性，并且是 chi 目的的核心。

Here is an example of a standard net/http middleware where we assign a context key `"user"` the value of `"123"`. This middleware sets a hypothetical user identifier on the request context and calls the next handler in the chain.

​	以下是一个标准 net/http 中间件的示例，其中我们将上下文键 `"user"` 分配给值 `"123"` 。此中间件在请求上下文中设置一个假设的用户标识符，并调用链中的下一个处理程序。

```go
// HTTP middleware setting a value on the request context
func MyMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // create new context from `r` request context, and assign key `"user"`
    // to value of `"123"`
    ctx := context.WithValue(r.Context(), "user", "123")

    // call the next handler in the chain, passing the response writer and
    // the updated request object with the new context value.
    //
    // note: context.Context values are nested, so any previously set
    // values will be accessible as well, and the new `"user"` key
    // will be accessible from this point forward.
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}Copy to clipboardErrorCopied
```

We can now take these values from the context in our Handlers like this:

​	我们现在可以像这样从处理程序中的上下文中获取这些值：

```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    // here we read from the request context and fetch out `"user"` key set in
    // the MyMiddleware example above.
    user := r.Context().Value("user").(string)

    // respond to the client
    w.Write([]byte(fmt.Sprintf("hi %s", user)))
}Copy to clipboardErrorCopied
```

## AllowContentEncoding

AllowContentEncoding enforces a whitelist of request Content-Encoding otherwise responds with a `415 Unsupported Media Type status`.

​	AllowContentEncoding 强制执行请求 Content-Encoding 的白名单，否则会响应 `415 Unsupported Media Type status` 。

Content-Encoding Parameters: `gzip`, `deflate`, `gzip, deflate`, `deflate, gzip`

​	Content-Encoding 参数： `gzip` , `deflate` , `gzip, deflate` , `deflate, gzip`

***This Middleware Doesn't Support `br` encoding
此中间件不支持 `br` 编码\***

Refer [Content-Encoding](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding)

​	参考 Content-Encoding

#### Usage 用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main() {
  r := chi.NewRouter()
  r.Use(middleware.AllowContentEncoding("deflate", "gzip"))
  r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## AllowContentType

AllowContentType enforces a whitelist of request Content-Types otherwise responds with a `415 Unsupported Media Type status`.

​	AllowContentType 强制执行请求 Content-Types 的白名单，否则会响应 `415 Unsupported Media Type status` 。

Content-Type Parameters: `application/json`, `text/xml`, `application/json, text/xml`

​	Content-Type 参数： `application/json` 、 `text/xml` 、 `application/json, text/xml`

Refer [Content-Type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type)

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
  r := chi.NewRouter()
  r.Use(middleware.AllowContentType("application/json","text/xml"))
  r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## CleanPath

CleanPath middleware will clean out double slash mistakes from a user's request path. For example, if a user requests /users//1 or //users////1 will both be treated as: /users/1

​	CleanPath 中间件将清除用户请求路径中的双斜杠错误。例如，如果用户请求 /users//1 或 //users////1，它们都将被视为：/users/1

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
  r := chi.NewRouter()
  r.Use(middleware.CleanPath)
  r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## Compress

Compress is a middleware that compresses response body of a given content types to a data format based on Accept-Encoding request header. It uses a given compression level.

​	Compress 是一个中间件，它将给定内容类型的响应正文压缩为基于 Accept-Encoding 请求头的格式。它使用给定的压缩级别。

**NOTE:** *make sure to set the Content-Type header on your response otherwise this middleware will not compress the response body. For ex, in your handler you should set w.Header().Set("Content-Type", http.DetectContentType(yourBody)) or set it manually.*

​	注意：确保在响应中设置 Content-Type 头，否则此中间件不会压缩响应正文。例如，在处理程序中，您应该设置 w.Header().Set("Content-Type", http.DetectContentType(yourBody)) 或手动设置。

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
  r := chi.NewRouter()
  r.Use(middleware.Compress(5, "text/html", "text/css"))
  r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## ContentCharset

ContentCharset generates a handler that writes a 415 Unsupported Media Type response if none of the charsets match. An empty charset will allow requests with no Content-Type header or no specified charset.

​	ContentCharset 生成一个处理程序，如果没有任何字符集匹配，则会写入 415 不支持的媒体类型响应。空字符集将允许没有 Content-Type 标头或未指定字符集的请求。

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
  r := chi.NewRouter()
  allowedCharsets := []string{"UTF-8", "Latin-1", ""}
  r.Use(middleware.ContentCharset(allowedCharsets...))
  r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## CORS

To Implement CORS in `chi` we can use [go-chi/cors](https://github.com/go-chi/cors)

​	要在 `chi` 中实现 CORS，我们可以使用 go-chi/cors

This middleware is designed to be used as a top-level middleware on the chi router. Applying with within a `r.Group()` or using `With()` **will not work without routes matching OPTIONS added**.

​	此中间件设计为在 chi 路由器上用作顶级中间件。在 `r.Group()` 中应用或使用 `With()` 将不起作用，除非添加了匹配 OPTIONS 的路由。

#### Usage 用法

```go
func main() {
  r := chi.NewRouter()

  // Basic CORS
  // for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
  r.Use(cors.Handler(cors.Options{
    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("welcome"))
  })

  http.ListenAndServe(":3000", r)
}Copy to clipboardErrorCopied
```

## GetHead

GetHead automatically route undefined HEAD requests to GET handlers.

​	GetHead 自动将未定义的 HEAD 请求路由到 GET 处理程序。

Reference: [HEAD](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/HEAD)

​	参考：HEAD

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
  r := chi.NewRouter()
  r.Use(middleware.GetHead)
  r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## Heartbeat 心跳

Heartbeat endpoint middleware useful to setting up a path like `/ping` that load balancers or uptime testing external services can make a request before hitting any routes. It's also convenient to place this above ACL middlewares as well.

​	心跳端点中间件可用于设置路径，例如 `/ping` ，负载均衡器或正常运行时间测试外部服务可以在访问任何路由之前发出请求。将其放在 ACL 中间件之上也很方便。

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
  r := chi.NewRouter()
  r.Use(middleware.Heartbeat("/"))
}Copy to clipboardErrorCopied
Get -> http://api_address/ 

Response -> ".", Status 200Copy to clipboardErrorCopied
```

## Logger 日志记录器

Logger is a middleware that logs the start and end of each request, along with some useful data about what was requested, what the response status was, and how long it took to return. When standard output is a TTY, Logger will print in color, otherwise it will print in black and white. Logger prints a request ID if one is provided.

​	日志记录器是一个中间件，它记录每个请求的开始和结束，以及一些有关请求内容、响应状态以及返回所用时间的有用数据。当标准输出为 TTY 时，日志记录器将以彩色打印，否则将以黑白打印。如果提供了请求 ID，日志记录器将打印该 ID。

Alternatively, look at https://github.com/goware/httplog for a more in-depth http logger with structured logging support.

​	或者，请参阅 https://github.com/goware/httplog，了解具有结构化日志记录支持的更深入的 http 日志记录器。

**IMPORTANT NOTE**: *Logger should go before any other middleware that may change the response, such as `middleware.Recoverer`*.

​	重要说明：日志记录器应在可能更改响应的任何其他中间件（例如 `middleware.Recoverer` ）之前。

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
  r := chi.NewRouter()
  r.Use(middleware.Logger)        // <--<< Logger should come before Recoverer
  r.Use(middleware.Recoverer)
  r.Get("/", handler)
}Copy to clipboardErrorCopied
```

## NoCache

NoCache is a simple piece of middleware that sets a number of HTTP headers to prevent a router (or subrouter) from being cached by an upstream proxy and/or client.

​	NoCache 是一段简单的中间件，它设置了许多 HTTP 标头，以防止路由器（或子路由器）被上游代理和/或客户端缓存。

As per http://wiki.nginx.org/HttpProxyModule - NoCache sets:

​	根据 http://wiki.nginx.org/HttpProxyModule - NoCache 设置：

```
Expires: Thu, 01 Jan 1970 00:00:00 UTC
Cache-Control: no-cache, private, max-age=0
X-Accel-Expires: 0
Pragma: no-cache (for HTTP/1.0 proxies/clients)Copy to clipboardErrorCopied
```

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
  r := chi.NewRouter()
  r.Use(middleware.NoCache)
  r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## Oauth 2.0

### Authorization Server 授权服务器

We can make an Authorization Server, which generates tokens for three scopes

​	我们可以创建一个授权服务器，它为三个范围生成令牌

1. username & password
   用户名和密码
2. clientID & Secret
   clientID 和 Secret
3. RefreshTokenGrant

Example:

​	示例：

```go
package main

import(
    "errors"
    "net/http"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
    "github.com/go-chi/oauth"
)

func main() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"*"},
    AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTION"},
    AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))
  registerAPI(r)
  _ = http.ListenAndServe(":8080", r)
}

func registerAPI(r *chi.Mux) {
  s := oauth.NewBearerServer(
    "mySecretKey-10101",
    time.Second*120,
    &TestUserVerifier{},
    nil)
  r.Post("/token", s.UserCredentials)
  r.Post("/auth", s.ClientCredentials)
}Copy to clipboardErrorCopied
```

#### Generate Token using username & password 使用用户名和密码生成令牌

```
    POST http://localhost:3000/token
    User-Agent: Fiddler
    Host: localhost:3000
    Content-Length: 50
    Content-Type: application/x-www-form-urlencoded

    grant_type=password&username=user01&password=12345Copy to clipboardErrorCopied
```

#### Generate Token using clientID & secret 使用 clientID 和 secret 生成令牌

```
    POST http://localhost:3000/auth
    User-Agent: Fiddler
    Host: localhost:3000
    Content-Length: 66
    Content-Type: application/x-www-form-urlencoded

    grant_type=client_credentials&client_id=abcdef&client_secret=12345Copy to clipboardErrorCopied
```

#### RefreshTokenGrant Token RefreshTokenGrant 令牌

```
    POST http://localhost:3000/token
    User-Agent: Fiddler
    Host: localhost:3000
    Content-Length: 50
    Content-Type: application/x-www-form-urlencoded

    grant_type=refresh_token&refresh_token={the refresh_token obtained in the previous response}
Copy to clipboardErrorCopied
```

Refer [Example](https://github.com/go-chi/oauth/blob/master/example/authserver/main.go) For the full Example...

​	有关完整示例，请参阅示例...

### Resource Server 资源服务器

Here we can implement oauth2 authentication and verification

​	我们可以在此处实现 oauth2 身份验证和验证

Example:

​	示例：

```go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"

    "github.com/go-chi/oauth"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTION"},
        AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
    }))
    registerAPI(r)
    _ = http.ListenAndServe(":8081", r)
}

func registerAPI(r *chi.Mux) {
    r.Route("/", func(r chi.Router) {
        // use the Bearer Authentication middleware
        r.Use(oauth.Authorize("mySecretKey-10101", nil))
        r.Get("/customers", GetCustomers)
        r.Get("/customers/{id}/orders", GetOrders)
    })
}Copy to clipboardErrorCopied
   Resource Server Example

    Get Customers

        GET http://localhost:3200/customers
        User-Agent: Fiddler
        Host: localhost:3200
        Content-Length: 0
        Content-Type: application/json
        Authorization: Bearer {access_token}

    Get Orders

        GET http://localhost:3200/customers/12345/orders
        User-Agent: Fiddler
        Host: localhost:3200
        Content-Length: 0
        Content-Type: application/json
        Authorization: Bearer {access_token}

    {access_token} is produced by the Authorization Server response (see example /test/authserver).Copy to clipboardErrorCopied
```

Refer [Example](https://github.com/go-chi/oauth/blob/master/example/resourceserver/main.go) For the full Example...

​	有关完整示例，请参阅示例...

## Profiler 分析器

Profiler is a convenient subrouter used for mounting net/http/pprof. ie. Usage

​	分析器是一个用于挂载 net/http/pprof 的便捷子路由器。即用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

 func main(){
   r := chi.NewRouter()
   // ..middlewares
   r.Mount("/debug", middleware.Profiler())
   // ..routes
}Copy to clipboardErrorCopied
```

Now you can request @ /debug for pprof profiles

​	现在，您可以请求 @ /debug 以获取 pprof 配置文件

## RealIP

RealIP is a middleware that sets a http.Request's RemoteAddr to the results of parsing either the X-Real-IP header or the X-Forwarded-For header (in that order).

​	RealIP 是一个中间件，它将 http.Request 的 RemoteAddr 设置为解析 X-Real-IP 头或 X-Forwarded-For 头（按此顺序）的结果。

This middleware should be inserted fairly early in the middleware stack to ensure that subsequent layers (e.g., request loggers) which examine the RemoteAddr will see the intended value.

​	此中间件应尽早插入中间件堆栈，以确保检查 RemoteAddr 的后续层（例如请求记录器）将看到预期值。

You should only use this middleware if you can trust the headers passed to you (in particular, the two headers this middleware uses), for example because you have placed a reverse proxy like HAProxy or nginx in front of chi. If your reverse proxies are configured to pass along arbitrary header values from the client, or if you use this middleware without a reverse proxy, malicious clients will be able to cause harm (or, depending on how you're using RemoteAddr, vulnerable to an attack of some sort).

​	您仅应在可信赖传递给您的标头（尤其是此中间件使用的两个标头）时才使用此中间件，例如，因为您已将 HAProxy 或 nginx 等反向代理放置在 chi 之前。如果您的反向代理配置为从客户端传递任意标头值，或者您在没有反向代理的情况下使用此中间件，恶意客户端将能够造成危害（或根据您使用 RemoteAddr 的方式，容易受到某种攻击）。

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

 func main(){
   r := chi.NewRouter()
   // ..middlewares
   r.Use(middleware.RealIP)
   // ..routes
}Copy to clipboardErrorCopied
```

## Recoverer

Recoverer is a middleware that recovers from panics, logs the panic (and a backtrace), and returns a HTTP 500 (Internal Server Error) status if possible. Recoverer prints a request ID if one is provided.

​	Recoverer 是一个中间件，它从 panic 中恢复，记录 panic（和回溯），并在可能的情况下返回 HTTP 500（内部服务器错误）状态。如果提供了请求 ID，Recoverer 会打印该 ID。

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

 func main(){
   r := chi.NewRouter()
   // ..middlewares
   r.Use(middleware.Recoverer)
   // ..routes
   r.Get("/", func(http.ResponseWriter, *http.Request) { panic("foo") })
}Copy to clipboardErrorCopied
```

## RedirectSlashes

RedirectSlashes is a middleware that will match request paths with a trailing slash and redirect to the same path, less the trailing slash.

​	RedirectSlashes 是一个中间件，它将匹配带有尾部斜杠的请求路径，并重定向到相同的路径，减去尾部斜杠。

NOTE: RedirectSlashes middleware is *incompatible* with http.FileServer, see [Issue 343](https://github.com/go-chi/chi/issues/343)

​	注意：RedirectSlashes 中间件与 http.FileServer 不兼容，请参阅问题 343

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
   r := chi.NewRouter()
   r.Use(middleware.RedirectSlashes)
   r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## RouteHeaders

RouteHeaders is a neat little header-based router that allows you to direct the flow of a request through a middleware stack based on a request header.

​	RouteHeaders 是一个简洁的小型基于标头的路由器，它允许您根据请求标头通过中间件堆栈来指导请求流。

For example, lets say you'd like to setup multiple routers depending on the request Host header, you could then do something as so:

​	例如，假设您想根据请求 Host 标头设置多个路由器，那么您可以执行以下操作：

```go
r := chi.NewRouter()
rSubdomain := chi.NewRouter()

r.Use(middleware.RouteHeaders().
  Route("Host", "example.com", middleware.New(r)).
  Route("Host", "*.example.com", middleware.New(rSubdomain)).
  Handler)

r.Get("/", h)
rSubdomain.Get("/", h2)
Copy to clipboardErrorCopied
```

Another example, imagine you want to setup multiple CORS handlers, where for your origin servers you allow authorized requests, but for third-party public requests, authorization is disabled.

​	另一个示例，假设您想设置多个 CORS 处理程序，在其中，对于您的源服务器，您允许授权请求，但对于第三方公共请求，则禁用授权。

```go
r := chi.NewRouter()

r.Use(middleware.RouteHeaders().
  Route("Origin", "https://app.skyweaver.net", cors.Handler(cors.Options{
     AllowedOrigins:   []string{"https://api.skyweaver.net"},
     AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
     AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
     AllowCredentials: true, // <----------<<< allow credentials
  })).
  Route("Origin", "*", cors.Handler(cors.Options{
     AllowedOrigins:   []string{"*"},
     AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
     AllowedHeaders:   []string{"Accept", "Content-Type"},
     AllowCredentials: false, // <----------<<< do not allow credentials
  })).
  Handler)Copy to clipboardErrorCopied
```

## StripSlashes

StripSlashes is a middleware that will match request paths with a trailing slash, strip it from the path and continue routing through the mux, if a route matches, then it will serve the handler.

​	StripSlashes 是一个中间件，它将匹配带有尾部斜杠的请求路径，从路径中将其剥离，并继续通过 mux 进行路由，如果匹配到路由，那么它将提供处理程序。

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
   r := chi.NewRouter()
   r.Use(middleware.StripSlashes)
   r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## Throttle 限流

Throttle is a middleware that limits number of currently processed requests at a time across all users. Note: Throttle is not a rate-limiter per user, instead it just puts a ceiling on the number of currentl in-flight requests being processed from the point from where the Throttle middleware is mounted.

​	限流是一个中间件，它限制所有用户当前处理的请求数。注意：限流不是针对每个用户的速率限制器，而是对从安装限流中间件的点开始处理的当前正在进行的请求数设置上限。

Throttle has a BacklogTimeout of 60 seconds by default

​	限流的 BacklogTimeout 默认值为 60 秒

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
    r := chi.NewRouter()
    r.Use(middleware.Throttle(15))
    r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## ThrottleBacklog 限流积压

ThrottleBacklog is a middleware that limits number of currently processed requests at a time and provides a backlog for holding a finite number of pending requests.

​	限流积压是一个中间件，它限制当前处理的请求数，并提供一个积压队列来保存有限数量的待处理请求。

Usage

​	用法

```go
import (
  "time"

  "github.com/go-chi/chi/v5/middleware"
)

func main(){
    r := chi.NewRouter()
    r.Use(ThrottleBacklog(10, 50, time.Second*10))
    r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
}Copy to clipboardErrorCopied
```

## Timeout 超时

Timeout is a middleware that cancels ctx after a given timeout and return a 504 Gateway Timeout error to the client.

​	超时是一个中间件，它在给定超时后取消 ctx 并向客户端返回 504 网关超时错误。

It's required that you select the ctx.Done() channel to check for the signal if the context has reached its deadline and return, otherwise the timeout signal will be just ignored.

​	您需要选择 ctx.Done() 通道来检查信号，如果上下文已达到其截止时间，则返回，否则将忽略超时信号。

ie. a route/handler may look like:

​	即，路由/处理程序可能如下所示：

```go
 r.Get("/long", func(w http.ResponseWriter, r *http.Request) {
   ctx := r.Context()
   processTime := time.Duration(rand.Intn(4)+1) * time.Second

   select {
   case <-ctx.Done():
     return

   case <-time.After(processTime):
      // The above channel simulates some hard work.
   }

   w.Write([]byte("done"))
 })Copy to clipboardErrorCopied
```

Usage

​	用法

```go
import (
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
    r := chi.NewRouter()
    r.Use(middleware.Timeout(time.Second*60))
    // handlers ...
}Copy to clipboardErrorCopied
```

## JWT Authentication JWT 身份验证

For Implementing JWT Authentication we can use `go-chi/jwtauth` It is a middleware built upon lestrrat-go/jwx

​	为了实现 JWT 身份验证，我们可以使用 `go-chi/jwtauth` ，它是一个基于 lestrrat-go/jwx 构建的中间件

The `jwtauth` http middleware package provides a simple way to verify a JWT token from a http request and send the result down the request context (`context.Context`).

​	`jwtauth` http 中间件包提供了一种简单的方法，可以从 http 请求中验证 JWT 令牌并将结果发送到请求上下文 ( `context.Context` )。

In a complete JWT-authentication flow, you'll first capture the token from a http request, decode it, verify it and then validate that its correctly signed and hasn't expired - the `jwtauth.Verifier` middleware handler takes care of all of that. The `jwtauth.Verifier` will set the context values on keys `jwtauth.TokenCtxKey` and `jwtauth.ErrorCtxKey`.

​	在完整的 JWT 身份验证流程中，您首先从 http 请求中捕获令牌，对其进行解码、验证，然后验证其是否已正确签名且尚未过期 - `jwtauth.Verifier` 中间件处理程序负责所有这些操作。 `jwtauth.Verifier` 会在键 `jwtauth.TokenCtxKey` 和 `jwtauth.ErrorCtxKey` 上设置上下文值。

Next, it's up to an authentication handler to respond or continue processing after the `jwtauth.Verifier`. The `jwtauth.Authenticator` middleware responds with a 401 Unauthorized plain-text payload for all unverified tokens and passes the good ones through. You can also copy the Authenticator and customize it to handle invalid tokens to better fit your flow (ie. with a JSON error response body).

​	接下来，由身份验证处理程序在 `jwtauth.Verifier` 之后做出响应或继续处理。 `jwtauth.Authenticator` 中间件会对所有未验证的令牌以纯文本有效负载形式做出 401 未授权的响应，并通过良好的令牌。您还可以复制 Authenticator 并对其进行自定义，以处理无效令牌，以便更好地适应您的流程（即使用 JSON 错误响应主体）。

By default, the `Verifier` will search for a JWT token in a http request, in the order:

​	默认情况下， `Verifier` 将按以下顺序在 http 请求中搜索 JWT 令牌：

1. 'Authorization: BEARER T' request header
   'Authorization: BEARER T' 请求头
2. 'jwt' Cookie value
   'jwt' Cookie 值

The first JWT string that is found as an authorization header or cookie header is then decoded by the `lestrrat-go/jwx` library and a jwt.Token object is set on the request context. In the case of a signature decoding error the Verifier will also set the error on the request context.

​	第一个作为授权头或 cookie 头找到的 JWT 字符串随后由 `lestrrat-go/jwx` 库解码，并且在请求上下文中设置一个 jwt.Token 对象。如果出现签名解码错误，Verifier 还会在请求上下文中设置错误。

The Verifier always calls the next http handler in sequence, which can either be the generic `jwtauth.Authenticator` middleware or your own custom handler which checks the request context jwt token and error to prepare a custom http response.

​	Verifier 始终按顺序调用下一个 http 处理程序，它可以是通用的 `jwtauth.Authenticator` 中间件，也可以是您自己的自定义处理程序，它检查请求上下文 jwt 令牌和错误以准备自定义 http 响应。

Note: jwtauth supports custom verification sequences for finding a token from a request by using the `Verify` middleware instantiator directly. The default `Verifier` is instantiated by calling `Verify(ja, TokenFromHeader, TokenFromCookie)`.

​	注意：jwtauth 支持自定义验证序列，通过直接使用 `Verify` 中间件实例化程序从请求中查找令牌。默认 `Verifier` 通过调用 `Verify(ja, TokenFromHeader, TokenFromCookie)` 实例化。

Usage

​	用法

See the full [example](https://github.com/go-chi/jwtauth/blob/master/_example/main.go).

​	请参阅完整示例。

```go
package main

import (
  "fmt"
  "net/http"

  "github.com/go-chi/chi/v5"
  "github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
  tokenAuth = jwtauth.New("HS256", []byte("secret"), nil) // replace with secret key

  // For debugging/example purposes, we generate and print
  // a sample jwt token with claims `user_id:123` here:
  _, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
  fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func main() {
  addr := ":3333"
  fmt.Printf("Starting server on %v\n", addr)
  http.ListenAndServe(addr, router())
}

func router() http.Handler {
  r := chi.NewRouter()

  // Protected routes
  r.Group(func(r chi.Router) {
    // Seek, verify and validate JWT tokens
    r.Use(jwtauth.Verifier(tokenAuth))

    // Handle valid / invalid tokens. In this example, we use
    // the provided authenticator middleware, but you can write your
    // own very easily, look at the Authenticator method in jwtauth.go
    // and tweak it, its not scary.
    r.Use(jwtauth.Authenticator)

    r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
      _, claims, _ := jwtauth.FromContext(r.Context())
      w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
    })
  })

  // Public routes
  r.Group(func(r chi.Router) {
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("welcome anonymous"))
    })
  })

  return r
}Copy to clipboardErrorCopied
```

## Http Rate Limiting Middleware Http 速率限制中间件

To implement this we can use [go-chi/httprate](https://github.com/go-chi/httprate)

​	要实现此目的，我们可以使用 go-chi/httprate

#### Usage 用法

```go
package main

import (
  "net/http"

  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/httprate"
)

func main() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)

  // Enable httprate request limiter of 100 requests per minute.
  //
  // In the code example below, rate-limiting is bound to the request IP address
  // via the LimitByIP middleware handler.
  //
  // To have a single rate-limiter for all requests, use httprate.LimitAll(..).
  //
  // Please see _example/main.go for other more, or read the library code.
  r.Use(httprate.LimitByIP(100, 1*time.Minute))

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("."))
  })

  http.ListenAndServe(":3333", r)
}
```