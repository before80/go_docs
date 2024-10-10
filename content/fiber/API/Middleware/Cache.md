+++
title = "Cache"
date = 2024-02-05T09:14:15+08:00
weight = 20
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/cache]({{< ref "/fiber/API/Middleware/Cache" >}})

# Cache 缓存

Cache middleware for [Fiber](https://github.com/gofiber/fiber) designed to intercept responses and cache them. This middleware will cache the `Body`, `Content-Type` and `StatusCode` using the `c.Path()` as unique identifier. Special thanks to [@codemicro](https://github.com/codemicro/fiber-cache) for creating this middleware for Fiber core!

​	专为拦截响应并缓存响应而设计的 Fiber 缓存中间件。此中间件将使用 `Body` 作为唯一标识符来缓存 `Content-Type` 和 `StatusCode` 。特别感谢 @codemicro 为 Fiber 内核创建此中间件！

Request Directives

​	请求指令
`Cache-Control: no-cache` will return the up-to-date response but still caches it. You will always get a `miss` cache status.

​	 `Cache-Control: no-cache` 将返回最新的响应，但仍会缓存它。您将始终获得 `miss` 缓存状态。
`Cache-Control: no-store` will refrain from caching. You will always get the up-to-date response.

​	 `Cache-Control: no-store` 将不进行缓存。您将始终获得最新的响应。

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
    "github.com/gofiber/fiber/v2/middleware/cache"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(cache.New())

// Or extend your config for customization
app.Use(cache.New(cache.Config{
    Next: func(c *fiber.Ctx) bool {
        return c.Query("noCache") == "true"
    },
    Expiration: 30 * time.Minute,
    CacheControl: true,
}))
```



Or you can custom key and expire time like this:

​	或者，您可以像这样自定义键和过期时间：

```go
app.Use(cache.New(cache.Config{
    ExpirationGenerator: func(c *fiber.Ctx, cfg *cache.Config) time.Duration {
        newCacheTime, _ := strconv.Atoi(c.GetRespHeader("Cache-Time", "600"))
        return time.Second * time.Duration(newCacheTime)
    },
    KeyGenerator: func(c *fiber.Ctx) string {
        return utils.CopyString(c.Path())
    },
}))

app.Get("/", func(c *fiber.Ctx) error {
    c.Response().Header.Add("Cache-Time", "6000")
    return c.SendString("hi")
})
```



## Config 配置

| Property 属性                      | Type 输入                                       | Description 说明                                             | Default 默认                                                 |
| ---------------------------------- | ----------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Next 下一步                        | `func(*fiber.Ctx) bool`                         | Next defines a function that is executed before creating the cache entry and can be used to execute the request without cache creation. If an entry already exists, it will be used. If you want to completely bypass the cache functionality in certain cases, you should use the [skip middleware]({{< ref "/fiber/API/Middleware/Skip" >}}). 接下来定义一个在创建缓存条目之前执行的函数，可用于在不创建缓存的情况下执行请求。如果条目已存在，则将使用该条目。如果您想在某些情况下完全绕过缓存功能，则应使用 skip 中间件。 | `nil`                                                        |
| Expiration 过期                    | `time.Duration`                                 | Expiration is the time that a cached response will live. 过期是缓存响应的生存时间。 | `1 * time.Minute`                                            |
| CacheHeader                        | `string`                                        | CacheHeader is the header on the response header that indicates the cache status, with the possible return values "hit," "miss," or "unreachable." CacheHeader 是响应头上的头，它指示缓存状态，可能的返回值为“命中”、“未命中”或“不可达”。 | `X-Cache`                                                    |
| CacheControl                       | `bool`                                          | CacheControl enables client-side caching if set to true. 如果设置为 true，CacheControl 将启用客户端缓存。 | `false`                                                      |
| KeyGenerator                       | `func(*fiber.Ctx) string`                       | Key allows you to generate custom keys. Key 允许您生成自定义密钥。 | `func(c *fiber.Ctx) string { return utils.CopyString(c.Path()) }` |
| ExpirationGenerator                | `func(*fiber.Ctx, *cache.Config) time.Duration` | ExpirationGenerator allows you to generate custom expiration keys based on the request. ExpirationGenerator 允许您根据请求生成自定义过期密钥。 | `nil`                                                        |
| Storage                            | `fiber.Storage`                                 | Store is used to store the state of the middleware. Store 用于存储中间件的状态。 | In-memory store 内存存储                                     |
| Store (Deprecated) Store（已弃用） | `fiber.Storage`                                 | Deprecated: Use Storage instead. 已弃用：请改用存储。        | In-memory store 内存存储                                     |
| Key (Deprecated) 密钥（已弃用）    | `func(*fiber.Ctx) string`                       | Deprecated: Use KeyGenerator instead. 已弃用：请改用 KeyGenerator。 | `nil`                                                        |
| StoreResponseHeaders               | `bool`                                          | StoreResponseHeaders allows you to store additional headers generated by next middlewares & handler. StoreResponseHeaders 允许您存储由下一个中间件和处理程序生成的附加标头。 | `false`                                                      |
| MaxBytes                           | `uint`                                          | MaxBytes is the maximum number of bytes of response bodies simultaneously stored in cache. MaxBytes 是同时存储在缓存中的响应正文的最大字节数。 | `0` (No limit) `0` （无限制）                                |
| Methods 方法                       | `[]string`                                      | Methods specifies the HTTP methods to cache. Methods 指定要缓存的 HTTP 方法。 | `[]string{fiber.MethodGet, fiber.MethodHead}`                |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next:         nil,
    Expiration:   1 * time.Minute,
    CacheHeader:  "X-Cache",
    CacheControl: false,
    KeyGenerator: func(c *fiber.Ctx) string {
        return utils.CopyString(c.Path())
    },
    ExpirationGenerator:  nil,
    StoreResponseHeaders: false,
    Storage:              nil,
    MaxBytes:             0,
    Methods: []string{fiber.MethodGet, fiber.MethodHead},
}
```
