+++
title = "CORS"
date = 2024-02-05T09:14:15+08:00
weight = 40
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/cors]({{< ref "/fiber/API/Middleware/CORS" >}})

# CORS

CORS middleware for [Fiber](https://github.com/gofiber/fiber) that can be used to enable [Cross-Origin Resource Sharing](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) with various options.

​	可用于启用跨源资源共享（具有各种选项）的 Fiber 的 CORS 中间件。

The middleware conforms to the `access-control-allow-origin` specification by parsing `AllowOrigins`. First, the middleware checks if there is a matching allowed origin for the requesting 'origin' header. If there is a match, it returns exactly one matching domain from the list of allowed origins.

​	中间件通过解析 `AllowOrigins` 来符合 `access-control-allow-origin` 规范。首先，中间件检查请求的“origin”标头是否有匹配的允许来源。如果有匹配项，它将从允许来源列表中返回一个完全匹配的域。

For more control, `AllowOriginsFunc` can be used to programatically determine if an origin is allowed. If no match was found in `AllowOrigins` and if `AllowOriginsFunc` returns true then the 'access-control-allow-origin' response header is set to the 'origin' request header.

​	为了更好地控制，可以使用 `AllowOriginsFunc` 来以编程方式确定是否允许来源。如果在 `AllowOrigins` 中未找到匹配项，并且如果 `AllowOriginsFunc` 返回 true，则将“access-control-allow-origin”响应标头设置为“origin”请求标头。

## Signatures 签名

```go
func New(config ...Config) fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(cors.New())

// Or extend your config for customization
app.Use(cors.New(cors.Config{
    AllowOrigins: "https://gofiber.io, https://gofiber.net",
    AllowHeaders:  "Origin, Content-Type, Accept",
}))
```



Using the `AllowOriginsFunc` function. In this example any origin will be allowed via CORS.

​	使用 `AllowOriginsFunc` 函数。在此示例中，将通过 CORS 允许任何来源。

For example, if a browser running on `http://localhost:3000` sends a request, this will be accepted and the `access-control-allow-origin` response header will be set to `http://localhost:3000`.

​	例如，如果运行在 `http://localhost:3000` 上的浏览器发送请求，这将被接受，并且 `access-control-allow-origin` 响应标头将设置为 `http://localhost:3000` 。

**Note: Using this feature is discouraged in production and it's best practice to explicitly set CORS origins via `AllowOrigins`.
注意：不建议在生产中使用此功能，最好通过 `AllowOrigins` 显式设置 CORS 来源。**

```go
app.Use(cors.New())

app.Use(cors.New(cors.Config{
    AllowOriginsFunc: func(origin string) bool {
        return os.Getenv("ENVIRONMENT") == "development"
    },
}))
```



## Config 配置

| Property 属性    | Type 输入                  | Description 说明                                             | Default 默认                       |
| ---------------- | -------------------------- | ------------------------------------------------------------ | ---------------------------------- |
| Next 下一步      | `func(*fiber.Ctx) bool`    | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                              |
| AllowOriginsFunc | `func(origin string) bool` | AllowOriginsFunc defines a function that will set the 'access-control-allow-origin' response header to the 'origin' request header when returned true. AllowOriginsFunc 定义一个函数，当返回 true 时，该函数将“access-control-allow-origin”响应标头设置为“origin”请求标头。 | `nil`                              |
| AllowOrigins     | `string`                   | AllowOrigin defines a comma separated list of origins that may access the resource. AllowOrigin 定义一个逗号分隔的来源列表，这些来源可以访问资源。 | `"*"`                              |
| AllowMethods     | `string`                   | AllowMethods defines a list of methods allowed when accessing the resource. This is used in response to a preflight request. AllowMethods 定义访问资源时允许使用的方法列表。这是对预检请求的响应。 | `"GET,POST,HEAD,PUT,DELETE,PATCH"` |
| AllowHeaders     | `string`                   | AllowHeaders defines a list of request headers that can be used when making the actual request. This is in response to a preflight request. AllowHeaders 定义在发出实际请求时可以使用的一个请求头列表。这是对预检请求的响应。 | `""`                               |
| AllowCredentials | `bool`                     | AllowCredentials indicates whether or not the response to the request can be exposed when the credentials flag is true. AllowCredentials 指示当 credentials 标志为 true 时，是否可以公开对请求的响应。 | `false`                            |
| ExposeHeaders    | `string`                   | ExposeHeaders defines a whitelist headers that clients are allowed to access. ExposeHeaders 定义客户端允许访问的白名单头。 | `""`                               |
| MaxAge           | `int`                      | MaxAge indicates how long (in seconds) the results of a preflight request can be cached. If you pass MaxAge 0, Access-Control-Max-Age header will not be added and browser will use 5 seconds by default. To disable caching completely, pass MaxAge value negative. It will set the Access-Control-Max-Age header 0. MaxAge 指示预检请求的结果可以缓存多长时间（以秒为单位）。如果您传递 MaxAge 0，则不会添加 Access-Control-Max-Age 头，浏览器将默认使用 5 秒。要完全禁用缓存，请传递 MaxAge 值为负数。它将把 Access-Control-Max-Age 头设置为 0。 | `0`                                |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next:         nil,
    AllowOriginsFunc: nil,
    AllowOrigins: "*",
    AllowMethods: strings.Join([]string{
        fiber.MethodGet,
        fiber.MethodPost,
        fiber.MethodHead,
        fiber.MethodPut,
        fiber.MethodDelete,
        fiber.MethodPatch,
    }, ","),
    AllowHeaders:     "",
    AllowCredentials: false,
    ExposeHeaders:    "",
    MaxAge:           0,
}
```
