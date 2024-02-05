+++
title = "App"
date = 2024-02-05T09:14:15+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/app]({{< ref "/fiber/API/App" >}})

# 🚀 App

## Static 静态

Use the **Static** method to serve static files such as **images**, **CSS,** and **JavaScript**.

​	使用 Static 方法来提供静态文件，例如图像、CSS 和 JavaScript。

INFO
信息

By default, **Static** will serve `index.html` files in response to a request on a directory.

​	默认情况下，Static 将在响应目录上的请求时提供 `index.html` 文件。

Signature
签名

```go
func (app *App) Static(prefix, root string, config ...Static) Router
```



Use the following code to serve files in a directory named `./public`

​	使用以下代码来提供名为 `./public` 的目录中的文件

```go
app.Static("/", "./public")

// => http://localhost:3000/hello.html
// => http://localhost:3000/js/jquery.js
// => http://localhost:3000/css/style.css
```



Examples
示例

```go
// Serve files from multiple directories
app.Static("/", "./public")

// Serve files from "./files" directory:
app.Static("/", "./files")
```



You can use any virtual path prefix (*where the path does not actually exist in the file system*) for files that are served by the **Static** method, specify a prefix path for the static directory, as shown below:

​	您可以对通过 Static 方法提供的文件使用任何虚拟路径前缀（其中路径实际上并不存在于文件系统中），为静态目录指定前缀路径，如下所示：

Examples
示例

```go
app.Static("/static", "./public")

// => http://localhost:3000/static/hello.html
// => http://localhost:3000/static/js/jquery.js
// => http://localhost:3000/static/css/style.css
```



If you want to have a little bit more control regarding the settings for serving static files. You could use the `fiber.Static` struct to enable specific settings.

​	如果您想对提供静态文件的设置有更多控制权。您可以使用 `fiber.Static` 结构来启用特定设置。

fiber.Static{}

```go
// Static defines configuration options when defining static assets.
type Static struct {
    // When set to true, the server tries minimizing CPU usage by caching compressed files.
    // This works differently than the github.com/gofiber/compression middleware.
    // Optional. Default value false
    Compress bool `json:"compress"`

    // When set to true, enables byte range requests.
    // Optional. Default value false
    ByteRange bool `json:"byte_range"`

    // When set to true, enables directory browsing.
    // Optional. Default value false.
    Browse bool `json:"browse"`

    // When set to true, enables direct download.
    // Optional. Default value false.
    Download bool `json:"download"`

    // The name of the index file for serving a directory.
    // Optional. Default value "index.html".
    Index string `json:"index"`

    // Expiration duration for inactive file handlers.
    // Use a negative time.Duration to disable it.
    //
    // Optional. Default value 10 * time.Second.
    CacheDuration time.Duration `json:"cache_duration"`

    // The value for the Cache-Control HTTP-header
    // that is set on the file response. MaxAge is defined in seconds.
    //
    // Optional. Default value 0.
    MaxAge int `json:"max_age"`

    // ModifyResponse defines a function that allows you to alter the response.
    //
    // Optional. Default: nil
    ModifyResponse Handler

    // Next defines a function to skip this middleware when returned true.
    //
    // Optional. Default: nil
    Next func(c *Ctx) bool
}
```



Example
示例

```go
// Custom config
app.Static("/", "./public", fiber.Static{
  Compress:      true,
  ByteRange:     true,
  Browse:        true,
  Index:         "john.html",
  CacheDuration: 10 * time.Second,
  MaxAge:        3600,
})
```



## Route Handlers 路由处理程序 

Registers a route bound to a specific [HTTP method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods).

​	注册绑定到特定 HTTP 方法的路由。

Signatures
签名

```go
// HTTP methods
func (app *App) Get(path string, handlers ...Handler) Router
func (app *App) Head(path string, handlers ...Handler) Router
func (app *App) Post(path string, handlers ...Handler) Router
func (app *App) Put(path string, handlers ...Handler) Router
func (app *App) Delete(path string, handlers ...Handler) Router
func (app *App) Connect(path string, handlers ...Handler) Router
func (app *App) Options(path string, handlers ...Handler) Router
func (app *App) Trace(path string, handlers ...Handler) Router
func (app *App) Patch(path string, handlers ...Handler) Router

// Add allows you to specifiy a method as value
func (app *App) Add(method, path string, handlers ...Handler) Router

// All will register the route on all HTTP methods
// Almost the same as app.Use but not bound to prefixes
func (app *App) All(path string, handlers ...Handler) Router
```



Examples
示例

```go
// Simple GET handler
app.Get("/api/list", func(c *fiber.Ctx) error {
  return c.SendString("I'm a GET request!")
})

// Simple POST handler
app.Post("/api/register", func(c *fiber.Ctx) error {
  return c.SendString("I'm a POST request!")
})
```



**Use** can be used for middleware packages and prefix catchers. These routes will only match the beginning of each path i.e. `/john` will match `/john/doe`, `/johnnnnn` etc

​	Use 可用于中间件包和前缀捕获器。这些路由将仅匹配每个路径的开头，即 `/john` 将匹配 `/john/doe` 、 `/johnnnnn` 等

Signature
签名

```go
func (app *App) Use(args ...interface{}) Router
```



Examples
示例

```go
// Match any request
app.Use(func(c *fiber.Ctx) error {
    return c.Next()
})

// Match request starting with /api
app.Use("/api", func(c *fiber.Ctx) error {
    return c.Next()
})

// Match requests starting with /api or /home (multiple-prefix support)
app.Use([]string{"/api", "/home"}, func(c *fiber.Ctx) error {
    return c.Next()
})

// Attach multiple handlers 
app.Use("/api", func(c *fiber.Ctx) error {
  c.Set("X-Custom-Header", random.String(32))
    return c.Next()
}, func(c *fiber.Ctx) error {
    return c.Next()
})
```



## Mount 挂载 

You can Mount Fiber instance by creating a `*Mount`

​	您可以通过创建 `*Mount` 来挂载 Fiber 实例

Signature
签名

```go
func (a *App) Mount(prefix string, app *App) Router
```



Examples
示例

```go
func main() {
    app := fiber.New()
    micro := fiber.New()
    app.Mount("/john", micro) // GET /john/doe -> 200 OK

    micro.Get("/doe", func(c *fiber.Ctx) error {
        return c.SendStatus(fiber.StatusOK)
    })

    log.Fatal(app.Listen(":3000"))
}
```



## MountPath

The `MountPath` property contains one or more path patterns on which a sub-app was mounted.

​	 `MountPath` 属性包含一个或多个子应用程序挂载的路径模式。

Signature
签名

```go
func (app *App) MountPath() string
```



Examples
示例

```go
func main() {
    app := fiber.New()
    one := fiber.New()
    two := fiber.New()
    three := fiber.New()

    two.Mount("/three", three)
    one.Mount("/two", two)
    app.Mount("/one", one)
  
    one.MountPath()   // "/one"
    two.MountPath()   // "/one/two"
    three.MountPath() // "/one/two/three"
    app.MountPath()   // ""
}
```



CAUTION
注意

Mounting order is important for MountPath. If you want to get mount paths properly, you should start mounting from the deepest app.

​	挂载顺序对于 MountPath 非常重要。如果您想正确获取挂载路径，则应从最深的应用程序开始挂载。

## Group 组 

You can group routes by creating a `*Group` struct.

​	您可以通过创建 `*Group` 结构来对路由进行分组。

Signature
签名

```go
func (app *App) Group(prefix string, handlers ...Handler) Router
```



Examples
示例

```go
func main() {
  app := fiber.New()

  api := app.Group("/api", handler)  // /api

  v1 := api.Group("/v1", handler)   // /api/v1
  v1.Get("/list", handler)          // /api/v1/list
  v1.Get("/user", handler)          // /api/v1/user

  v2 := api.Group("/v2", handler)   // /api/v2
  v2.Get("/list", handler)          // /api/v2/list
  v2.Get("/user", handler)          // /api/v2/user

  log.Fatal(app.Listen(":3000"))
}
```



## Route 路由 

You can define routes with a common prefix inside the common function.

​	您可以在 common 函数中定义具有公共前缀的路由。

Signature
签名

```go
func (app *App) Route(prefix string, fn func(router Router), name ...string) Router
```



Examples
示例

```go
func main() {
  app := fiber.New()

  app.Route("/test", func(api fiber.Router) {
      api.Get("/foo", handler).Name("foo") // /test/foo (name: test.foo)
      api.Get("/bar", handler).Name("bar") // /test/bar (name: test.bar)
  }, "test.")

  log.Fatal(app.Listen(":3000"))
}
```



## Server

Server returns the underlying [fasthttp server](https://godoc.org/github.com/valyala/fasthttp#Server)

​	Server 返回底层的 fasthttp 服务器

Signature
签名

```go
func (app *App) Server() *fasthttp.Server
```



Examples
示例

```go
func main() {
    app := fiber.New()

    app.Server().MaxConnsPerIP = 1

    // ...
}
```



## Server Shutdown

Shutdown gracefully shuts down the server without interrupting any active connections. Shutdown works by first closing all open listeners and then waits indefinitely for all connections to return to idle before shutting down.

​	Shutdown 优雅地关闭服务器，不会中断任何活动连接。Shutdown 的工作原理是首先关闭所有打开的侦听器，然后无限期地等待所有连接返回空闲状态，然后再关闭。

ShutdownWithTimeout will forcefully close any active connections after the timeout expires.

​	ShutdownWithTimeout 将在超时到期后强制关闭所有活动连接。

ShutdownWithContext shuts down the server including by force if the context's deadline is exceeded.

​	ShutdownWithContext 将关闭服务器，包括在超出上下文的截止日期时强制关闭。

```go
func (app *App) Shutdown() error
func (app *App) ShutdownWithTimeout(timeout time.Duration) error
func (app *App) ShutdownWithContext(ctx context.Context) error
```



## HandlersCount

This method returns the amount of registered handlers.

​	此方法返回已注册处理程序的数量。

Signature
签名

```go
func (app *App) HandlersCount() uint32
```



## Stack

This method returns the original router stack

​	此方法返回原始路由堆栈

Signature
签名

```go
func (app *App) Stack() [][]*Route
```



Examples
示例

```go
var handler = func(c *fiber.Ctx) error { return nil }

func main() {
    app := fiber.New()

    app.Get("/john/:age", handler)
    app.Post("/register", handler)

    data, _ := json.MarshalIndent(app.Stack(), "", "  ")
    fmt.Println(string(data))

    app.Listen(":3000")
}
```



Result
结果

```javascript
[
  [
    {
      "method": "GET",
      "path": "/john/:age",
      "params": [
        "age"
      ]
    }
  ],
  [
    {
      "method": "HEAD",
      "path": "/john/:age",
      "params": [
        "age"
      ]
    }
  ],
  [
    {
      "method": "POST",
      "path": "/register",
      "params": null
    }
  ]
]
```



## Name 名称 

This method assigns the name of latest created route.

​	此方法分配最新创建的路由的名称。

Signature
签名

```go
func (app *App) Name(name string) Router
```



Examples
示例

```go
var handler = func(c *fiber.Ctx) error { return nil }

func main() {
    app := fiber.New()

    app.Get("/", handler)
    app.Name("index")

    app.Get("/doe", handler).Name("home")

    app.Trace("/tracer", handler).Name("tracert")

    app.Delete("/delete", handler).Name("delete")

    a := app.Group("/a")
    a.Name("fd.")

    a.Get("/test", handler).Name("test")

    data, _ := json.MarshalIndent(app.Stack(), "", "  ")
    fmt.Print(string(data))

    app.Listen(":3000")

}
```



Result
结果

```javascript
[
  [
    {
      "method": "GET",
      "name": "index",
      "path": "/",
      "params": null
    },
    {
      "method": "GET",
      "name": "home",
      "path": "/doe",
      "params": null
    },
    {
      "method": "GET",
      "name": "fd.test",
      "path": "/a/test",
      "params": null
    }
  ],
  [
    {
      "method": "HEAD",
      "name": "",
      "path": "/",
      "params": null
    },
    {
      "method": "HEAD",
      "name": "",
      "path": "/doe",
      "params": null
    },
    {
      "method": "HEAD",
      "name": "",
      "path": "/a/test",
      "params": null
    }
  ],
  null,
  null,
  [
    {
      "method": "DELETE",
      "name": "delete",
      "path": "/delete",
      "params": null
    }
  ],
  null,
  null,
  [
    {
      "method": "TRACE",
      "name": "tracert",
      "path": "/tracer",
      "params": null
    }
  ],
  null
]
```



## GetRoute

This method gets the route by name.

​	此方法按名称获取路由。

Signature
签名

```go
func (app *App) GetRoute(name string) Route
```



Examples
示例

```go
var handler = func(c *fiber.Ctx) error { return nil }

func main() {
    app := fiber.New()

    app.Get("/", handler).Name("index")
    
    data, _ := json.MarshalIndent(app.GetRoute("index"), "", "  ")
    fmt.Print(string(data))


    app.Listen(":3000")

}
```



Result
结果

```javascript
{
  "method": "GET",
  "name": "index",
  "path": "/",
  "params": null
}
```



## GetRoutes

This method gets all routes.

​	此方法获取所有路由。

Signature
签名

```go
func (app *App) GetRoutes(filterUseOption ...bool) []Route
```



When filterUseOption equal to true, it will filter the routes registered by the middleware.

​	当 filterUseOption 等于 true 时，它将过滤由中间件注册的路由。

Examples
示例

```go
func main() {
    app := fiber.New()
    app.Post("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    }).Name("index")
    data, _ := json.MarshalIndent(app.GetRoutes(true), "", "  ")
    fmt.Print(string(data))
}
```



Result
结果

```javascript
[
    {
        "method": "POST",
        "name": "index",
        "path": "/",
        "params": null
    }
]
```



## Config 配置

Config returns the app config as value ( read-only ).

​	Config 返回应用程序配置作为值（只读）。

Signature
签名

```go
func (app *App) Config() Config
```



## Handler 处理程序 

Handler returns the server handler that can be used to serve custom *fasthttp.RequestCtx requests.

​	处理程序返回可用于处理自定义 *fasthttp.RequestCtx 请求的服务器处理程序。

Signature
签名

```go
func (app *App) Handler() fasthttp.RequestHandler
```



## Listen 监听 

Listen serves HTTP requests from the given address.

​	监听从给定地址提供 HTTP 请求。

Signature
签名

```go
func (app *App) Listen(addr string) error
```



Examples
示例

```go
// Listen on port :8080 
app.Listen(":8080")

// Custom host
app.Listen("127.0.0.1:8080")
```



## ListenTLS

ListenTLS serves HTTPs requests from the given address using certFile and keyFile paths to as TLS certificate and key file.

​	ListenTLS 使用 certFile 和 keyFile 路径作为 TLS 证书和密钥文件，从给定地址提供 HTTPs 请求。

Signature
签名

```go
func (app *App) ListenTLS(addr, certFile, keyFile string) error
```



Examples
示例

```go
app.ListenTLS(":443", "./cert.pem", "./cert.key");
```



Using `ListenTLS` defaults to the following config ( use `Listener` to provide your own config )

​	使用 `ListenTLS` 默认为以下配置（使用 `Listener` 提供您自己的配置）

Default *tls.Config
默认 *tls.Config

```go
&tls.Config{
    MinVersion:               tls.VersionTLS12,
    Certificates: []tls.Certificate{
        cert,
    },
}
```



## ListenTLSWithCertificate

Signature
签名

```go
func (app *App) ListenTLS(addr string, cert tls.Certificate) error
```



Examples
示例

```go
app.ListenTLSWithCertificate(":443", cert);
```



Using `ListenTLSWithCertificate` defaults to the following config ( use `Listener` to provide your own config )

​	使用 `ListenTLSWithCertificate` 默认为以下配置（使用 `Listener` 提供您自己的配置）

Default *tls.Config
默认 *tls.Config

```go
&tls.Config{
    MinVersion:               tls.VersionTLS12,
    Certificates: []tls.Certificate{
        cert,
    },
}
```



## ListenMutualTLS

ListenMutualTLS serves HTTPs requests from the given address using certFile, keyFile and clientCertFile are the paths to TLS certificate and key file

​	ListenMutualTLS 使用 certFile、keyFile 和 clientCertFile 从给定地址提供 HTTPs 请求，这些路径指向 TLS 证书和密钥文件

Signature
签名

```go
func (app *App) ListenMutualTLS(addr, certFile, keyFile, clientCertFile string) error
```



Examples
示例

```go
app.ListenMutualTLS(":443", "./cert.pem", "./cert.key", "./ca-chain-cert.pem");
```



Using `ListenMutualTLS` defaults to the following config ( use `Listener` to provide your own config )

​	使用 `ListenMutualTLS` 默认为以下配置（使用 `Listener` 提供您自己的配置）

Default *tls.Config
默认 *tls.Config

```go
&tls.Config{
    MinVersion: tls.VersionTLS12,
    ClientAuth: tls.RequireAndVerifyClientCert,
    ClientCAs:  clientCertPool,
    Certificates: []tls.Certificate{
        cert,
    },
}
```



## ListenMutualTLSWithCertificate

ListenMutualTLSWithCertificate serves HTTPs requests from the given address using certFile, keyFile and clientCertFile are the paths to TLS certificate and key file

​	ListenMutualTLSWithCertificate 使用 certFile、keyFile 和 clientCertFile 从给定地址提供 HTTPs 请求，这些路径指向 TLS 证书和密钥文件

Signature
签名

```go
func (app *App) ListenMutualTLSWithCertificate(addr string, cert tls.Certificate, clientCertPool *x509.CertPool) error
```



Examples
示例

```go
app.ListenMutualTLSWithCertificate(":443", cert, clientCertPool);
```



Using `ListenMutualTLSWithCertificate` defaults to the following config ( use `Listener` to provide your own config )

​	使用 `ListenMutualTLSWithCertificate` 默认为以下配置（使用 `Listener` 提供您自己的配置）

Default *tls.Config
默认 *tls.Config

```go
&tls.Config{
    MinVersion: tls.VersionTLS12,
    ClientAuth: tls.RequireAndVerifyClientCert,
    ClientCAs:  clientCertPool,
    Certificates: []tls.Certificate{
        cert,
    },
}
```



## Listener 侦听器 

You can pass your own [`net.Listener`](https://pkg.go.dev/net/#Listener) using the `Listener` method. This method can be used to enable **TLS/HTTPS** with a custom tls.Config.

​	您可以使用 `Listener` 方法传递您自己的 `net.Listener` 。此方法可用于使用自定义 tls.Config 启用 TLS/HTTPS。

Signature
签名

```go
func (app *App) Listener(ln net.Listener) error
```



Examples
示例

```go
ln, _ := net.Listen("tcp", ":3000")

cer, _:= tls.LoadX509KeyPair("server.crt", "server.key")

ln = tls.NewListener(ln, &tls.Config{Certificates: []tls.Certificate{cer}})

app.Listener(ln)
```



## Test 测试 

Testing your application is done with the **Test** method. Use this method for creating `_test.go` files or when you need to debug your routing logic. The default timeout is `1s` if you want to disable a timeout altogether, pass `-1` as a second argument.

​	使用 Test 方法测试您的应用程序。使用此方法创建 `_test.go` 文件或当您需要调试路由逻辑时。默认超时为 `1s` ，如果您想完全禁用超时，请将 `-1` 作为第二个参数传递。

Signature
签名

```go
func (app *App) Test(req *http.Request, msTimeout ...int) (*http.Response, error)
```



Examples
示例

```go
// Create route with GET method for test:
app.Get("/", func(c *fiber.Ctx) error {
  fmt.Println(c.BaseURL())              // => http://google.com
  fmt.Println(c.Get("X-Custom-Header")) // => hi

  return c.SendString("hello, World!")
})

// http.Request
req := httptest.NewRequest("GET", "http://google.com", nil)
req.Header.Set("X-Custom-Header", "hi")

// http.Response
resp, _ := app.Test(req)

// Do something with results:
if resp.StatusCode == fiber.StatusOK {
  body, _ := io.ReadAll(resp.Body)
  fmt.Println(string(body)) // => Hello, World!
}
```



## Hooks 挂钩 

Hooks is a method to return [hooks]({{< ref "/fiber/Guide/Hooks" >}}) property.

​	Hooks 是返回挂钩属性的方法。

Signature
签名

```go
func (app *App) Hooks() *Hooks
```
