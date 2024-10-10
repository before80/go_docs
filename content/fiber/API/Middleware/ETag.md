+++
title = "ETag"
date = 2024-02-05T09:14:15+08:00
weight = 90
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/etag]({{< ref "/fiber/API/Middleware/ETag" >}})

# ETag

ETag middleware for [Fiber](https://github.com/gofiber/fiber) that lets caches be more efficient and save bandwidth, as a web server does not need to resend a full response if the content has not changed.

​	ETag 中间件，用于 Fiber，它让缓存更高效并节省带宽，因为如果内容没有改变，Web 服务器不需要重新发送完整响应。

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
  "github.com/gofiber/fiber/v2/middleware/etag"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(etag.New())

// Get / receives Etag: "13-1831710635" in response header
app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})

// Or extend your config for customization
app.Use(etag.New(etag.Config{
    Weak: true,
}))

// Get / receives Etag: "W/"13-1831710635" in response header
app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})
```



## Config 配置

| Property 属性 | Type 输入               | Description 说明                                             | Default 默认 |
| ------------- | ----------------------- | ------------------------------------------------------------ | ------------ |
| Weak 弱       | `bool`                  | Weak indicates that a weak validator is used. Weak etags are easy to generate but are less useful for comparisons. 弱表示使用的是弱验证器。弱 etag 易于生成，但对于比较而言不太有用。 | `false`      |
| Next 下一步   | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`        |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next: nil,
    Weak: false,
}
```
