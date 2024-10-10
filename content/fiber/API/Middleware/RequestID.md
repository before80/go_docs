+++
title = "RequestID"
date = 2024-02-05T09:14:15+08:00
weight = 240
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/requestid]({{< ref "/fiber/API/Middleware/RequestID" >}})

# RequestID

RequestID middleware for [Fiber](https://github.com/gofiber/fiber) that adds an identifier to the response.

​	RequestID 中间件，为 Fiber 添加一个标识符到响应中。

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
  "github.com/gofiber/fiber/v2/middleware/requestid"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(requestid.New())

// Or extend your config for customization
app.Use(requestid.New(requestid.Config{
    Header:    "X-Custom-Header",
    Generator: func() string {
        return "static-id"
    },
}))
```



## Config 配置

| Property 属性    | Type 输入               | Description 说明                                             | Default 默认   |
| ---------------- | ----------------------- | ------------------------------------------------------------ | -------------- |
| Next 下一步      | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`          |
| Header 标题      | `string`                | Header is the header key where to get/set the unique request ID. Header 是获取/设置唯一请求 ID 的标头键。 | "X-Request-ID" |
| Generator 生成器 | `func() string`         | Generator defines a function to generate the unique identifier. 生成器定义一个生成唯一标识符的函数。 | utils.UUID     |
| ContextKey       | `interface{}`           | ContextKey defines the key used when storing the request ID in the locals for a specific request. ContextKey 定义在本地存储特定请求的请求 ID 时使用的键。 | "requestid"    |

## Default Config 默认配置 

The default config uses a fast UUID generator which will expose the number of requests made to the server. To conceal this value for better privacy, use the `utils.UUIDv4` generator.

​	默认配置使用一个快速的 UUID 生成器，它将公开发送到服务器的请求数。为了更好地保护隐私而隐藏此值，请使用 `utils.UUIDv4` 生成器。

```go
var ConfigDefault = Config{
    Next:       nil,
    Header:     fiber.HeaderXRequestID,
    Generator:  utils.UUID,
    ContextKey: "requestid",
}
```
