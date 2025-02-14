+++
title = "Adaptor"
date = 2024-02-05T09:14:15+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/adaptor]({{< ref "/fiber/API/Middleware/Adaptor" >}})

# Adaptor 适配器

Converter for net/http handlers to/from Fiber request handlers, special thanks to [@arsmn](https://github.com/arsmn)!

​	net/http 处理程序到/从 Fiber 请求处理程序的转换器，特别感谢 @arsmn！

## Signatures 签名

| Name 名称                 | Signature 签名                                               | Description 说明                                 |
| ------------------------- | ------------------------------------------------------------ | ------------------------------------------------ |
| HTTPHandler               | `HTTPHandler(h http.Handler) fiber.Handler`                  | http.Handler -> fiber.Handler                    |
| HTTPHandlerFunc           | `HTTPHandlerFunc(h http.HandlerFunc) fiber.Handler`          | http.HandlerFunc -> fiber.Handler                |
| HTTPMiddleware            | `HTTPHandlerFunc(mw func(http.Handler) http.Handler) fiber.Handler` | func(http.Handler) http.Handler -> fiber.Handler |
| FiberHandler              | `FiberHandler(h fiber.Handler) http.Handler`                 | fiber.Handler -> http.Handler                    |
| FiberHandlerFunc          | `FiberHandlerFunc(h fiber.Handler) http.HandlerFunc`         | fiber.Handler -> http.HandlerFunc                |
| FiberApp                  | `FiberApp(app *fiber.App) http.HandlerFunc`                  | Fiber app -> http.HandlerFunc                    |
| ConvertRequest            | `ConvertRequest(c *fiber.Ctx, forServer bool) (*http.Request, error)` | fiber.Ctx -> http.Request                        |
| CopyContextToFiberContext | `CopyContextToFiberContext(context interface{}, requestContext *fasthttp.RequestCtx)` | context.Context -> fasthttp.RequestCtx           |

## Examples 示例 

### net/http to Fiber net/http 到 Fiber 

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
    // New fiber app
    app := fiber.New()

    // http.Handler -> fiber.Handler
    app.Get("/", adaptor.HTTPHandler(handler(greet)))

    // http.HandlerFunc -> fiber.Handler
    app.Get("/func", adaptor.HTTPHandlerFunc(greet))

    // Listen on port 3000
    app.Listen(":3000")
}

func handler(f http.HandlerFunc) http.Handler {
    return http.HandlerFunc(f)
}

func greet(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World!")
}
```



### net/http middleware to Fiber net/http 中间件到 Fiber Handler Fiber Handler 到 net/http Fiber App 到 net/http Fiber 到 (net/http).Request BasicAuth

```go
package main

import (
    "log"
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
    // New fiber app
    app := fiber.New()

    // http middleware -> fiber.Handler
    app.Use(adaptor.HTTPMiddleware(logMiddleware))

    // Listen on port 3000
    app.Listen(":3000")
}

func logMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("log middleware")
        next.ServeHTTP(w, r)
    })
}
```



### Fiber Handler to net/http

```go
package main

import (
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
    // fiber.Handler -> http.Handler
    http.Handle("/", adaptor.FiberHandler(greet))

    // fiber.Handler -> http.HandlerFunc
    http.HandleFunc("/func", adaptor.FiberHandlerFunc(greet))

    // Listen on port 3000
    http.ListenAndServe(":3000", nil)
}

func greet(c *fiber.Ctx) error {
    return c.SendString("Hello World!")
}
```



### Fiber App to net/http

```go
package main

import (
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
    app := fiber.New()

    app.Get("/greet", greet)

    // Listen on port 3000
    http.ListenAndServe(":3000", adaptor.FiberApp(app))
}

func greet(c *fiber.Ctx) error {
    return c.SendString("Hello World!")
}
```



### Fiber Context to (net/http).Request

```go
package main

import (
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
    app := fiber.New()

    app.Get("/greet", greetWithHTTPReq)

    // Listen on port 3000
    http.ListenAndServe(":3000", adaptor.FiberApp(app))
}

func greetWithHTTPReq(c *fiber.Ctx) error {
    httpReq, err := adaptor.ConvertRequest(c, false)
    if err != nil {
        return err
    }

    return c.SendString("Request URL: " + httpReq.URL.String())
}
```
