+++
title = "Limiter"
date = 2024-02-05T09:14:15+08:00
weight = 170
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/limiter]({{< ref "/fiber/API/Middleware/Limiter" >}})

# Limiter

Limiter middleware for [Fiber](https://github.com/gofiber/fiber) that is used to limit repeat requests to public APIs and/or endpoints such as password reset. It is also useful for API clients, web crawling, or other tasks that need to be throttled.

​	Fiber 的限流器中间件用于限制对公共 API 和/或端点的重复请求，例如密码重置。它还适用于 API 客户、网络爬虫或需要节流的其他任务。

NOTE
注意

This middleware uses our [Storage](https://github.com/gofiber/storage) package to support various databases through a single interface. The default configuration for this middleware saves data to memory, see the examples below for other databases.

​	此中间件使用我们的存储包通过单个接口支持各种数据库。此中间件的默认配置将数据保存到内存中，请参阅以下示例以了解其他数据库。

NOTE
注意

This module does not share state with other processes/servers by default.

​	默认情况下，此模块不与其他进程/服务器共享状态。

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
  "github.com/gofiber/fiber/v2/middleware/limiter"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(limiter.New())

// Or extend your config for customization
app.Use(limiter.New(limiter.Config{
    Next: func(c *fiber.Ctx) bool {
        return c.IP() == "127.0.0.1"
    },
    Max:          20,
    Expiration:     30 * time.Second,
    KeyGenerator:          func(c *fiber.Ctx) string {
        return c.Get("x-forwarded-for")
    },
    LimitReached: func(c *fiber.Ctx) error {
        return c.SendFile("./toofast.html")
    },
    Storage: myCustomStorage{},
}))
```



## Sliding window 滑动窗口

Instead of using the standard fixed window algorithm, you can enable the [sliding window](https://en.wikipedia.org/wiki/Sliding_window_protocol) algorithm.

​	您可以启用滑动窗口算法，而不是使用标准固定窗口算法。

A example of such configuration is:

​	此类配置的一个示例是：

```go
app.Use(limiter.New(limiter.Config{
    Max:            20,
    Expiration:     30 * time.Second,
    LimiterMiddleware: limiter.SlidingWindow{},
}))
```



This means that every window will take into account the previous window(if there was any). The given formula for the rate is:

​	这意味着每个窗口都会考虑前一个窗口（如果有的话）。给定的速率公式为：

```text
weightOfPreviousWindpw = previous window's amount request * (whenNewWindow / Expiration)
rate = weightOfPreviousWindpw + current window's amount request.
```



## Config 配置

| Property 属性                            | Type 输入                 | Description 说明                                             | Default 默认                                                 |
| ---------------------------------------- | ------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Next 下一步                              | `func(*fiber.Ctx) bool`   | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                                                        |
| Max                                      | `int`                     | Max number of recent connections during `Expiration` seconds before sending a 429 response. 在发送 429 响应之前， `Expiration` 秒内最近连接的最大数量。 | 5                                                            |
| KeyGenerator                             | `func(*fiber.Ctx) string` | KeyGenerator allows you to generate custom keys, by default c.IP() is used. KeyGenerator 允许您生成自定义键，默认情况下使用 c.IP()。 | A function using c.IP() as the default 使用 c.IP() 作为默认值的一个函数 |
| Expiration 过期                          | `time.Duration`           | Expiration is the time on how long to keep records of requests in memory. 过期时间是将请求记录保存在内存中的时间长度。 | 1 * time.Minute                                              |
| LimitReached                             | `fiber.Handler`           | LimitReached is called when a request hits the limit. 当请求达到限制时调用 LimitReached。 | A function sending 429 response 发送 429 响应的函数          |
| SkipFailedRequests                       | `bool`                    | When set to true, requests with StatusCode >= 400 won't be counted. 设置为 true 时，不会计算状态代码 >= 400 的请求。 | false                                                        |
| SkipSuccessfulRequests                   | `bool`                    | When set to true, requests with StatusCode < 400 won't be counted. 设置为 true 时，不会计算状态代码 < 400 的请求。 | false                                                        |
| Storage                                  | `fiber.Storage`           | Store is used to store the state of the middleware. Store 用于存储中间件的状态。 | An in-memory store for this process only 仅针对此进程的内存存储 |
| LimiterMiddleware                        | `LimiterHandler`          | LimiterMiddleware is the struct that implements a limiter middleware. LimiterMiddleware 是实现限流中间件的结构。 | A new Fixed Window Rate Limiter 新的固定窗口速率限制器       |
| Duration (Deprecated) 持续时间（已弃用） | `time.Duration`           | Deprecated: Use Expiration instead 已弃用：改用 Expiration   | -                                                            |
| Store (Deprecated) Store（已弃用）       | `fiber.Storage`           | Deprecated: Use Storage instead 已弃用：改用 Storage         | -                                                            |
| Key (Deprecated) 密钥（已弃用）          | `func(*fiber.Ctx) string` | Deprecated: Use KeyGenerator instead 已弃用：改用 KeyGenerator | -                                                            |

NOTE
注意

A custom store can be used if it implements the `Storage` interface - more details and an example can be found in `store.go`.

​	如果自定义存储实现了 `Storage` 接口，则可以使用它 - 更多详细信息和示例可以在 `store.go` 中找到。

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Max:        5,
    Expiration: 1 * time.Minute,
    KeyGenerator: func(c *fiber.Ctx) string {
        return c.IP()
    },
    LimitReached: func(c *fiber.Ctx) error {
        return c.SendStatus(fiber.StatusTooManyRequests)
    },
    SkipFailedRequests: false,
    SkipSuccessfulRequests: false,
    LimiterMiddleware: FixedWindow{},
}
```



### Custom Storage/Database 自定义存储/数据库

You can use any storage from our [storage](https://github.com/gofiber/storage/) package.

​	您可以使用存储包中的任何存储。

```go
storage := sqlite3.New() // From github.com/gofiber/storage/sqlite3
app.Use(limiter.New(limiter.Config{
    Storage: storage,
}))
```
