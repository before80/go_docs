+++
title = "EarlyData"
date = 2024-02-05T09:14:15+08:00
weight = 60
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/earlydata]({{< ref "/fiber/API/Middleware/EarlyData" >}})

# EarlyData

The Early Data middleware for [Fiber](https://github.com/gofiber/fiber) adds support for TLS 1.3's early data ("0-RTT") feature. Citing [RFC 8446](https://datatracker.ietf.org/doc/html/rfc8446#section-2-3), when a client and server share a PSK, TLS 1.3 allows clients to send data on the first flight ("early data") to speed up the request, effectively reducing the regular 1-RTT request to a 0-RTT request.

​	Fiber 的 Early Data 中间件增加了对 TLS 1.3 早期数据（“0-RTT”）功能的支持。引用 RFC 8446，当客户端和服务器共享 PSK 时，TLS 1.3 允许客户端在第一次传输（“早期数据”）中发送数据，以加快请求速度，有效地将常规 1-RTT 请求减少到 0-RTT 请求。

Make sure to enable fiber's `EnableTrustedProxyCheck` config option before using this middleware in order to not trust bogus HTTP request headers of the client.

​	在使用此中间件之前，请务必启用 fiber 的 `EnableTrustedProxyCheck` 配置选项，以免信任客户端的虚假 HTTP 请求头。

Also be aware that enabling support for early data in your reverse proxy (e.g. nginx, as done with a simple `ssl_early_data on;`) makes requests replayable. Refer to the following documents before continuing:

​	还要注意，在反向代理中启用对早期数据的支持（例如 nginx，通过简单的 `ssl_early_data on;` 完成）会使请求可重放。在继续之前，请参阅以下文档：

- https://datatracker.ietf.org/doc/html/rfc8446#section-8
- https://blog.trailofbits.com/2019/03/25/what-application-developers-need-to-know-about-tls-early-data-0rtt/

By default, this middleware allows early data requests on safe HTTP request methods only and rejects the request otherwise, i.e. aborts the request before executing your handler. This behavior can be controlled by the `AllowEarlyData` config option. Safe HTTP methods — `GET`, `HEAD`, `OPTIONS` and `TRACE` — should not modify a state on the server.

​	默认情况下，此中间件仅允许安全 HTTP 请求方法的早期数据请求，否则拒绝该请求，即在执行处理程序之前中止该请求。此行为可由 `AllowEarlyData` 配置选项控制。安全 HTTP 方法 — `GET` 、 `HEAD` 、 `OPTIONS` 和 `TRACE` — 不应修改服务器上的状态。

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
    "github.com/gofiber/fiber/v2/middleware/earlydata"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(earlydata.New())

// Or extend your config for customization
app.Use(earlydata.New(earlydata.Config{
    Error: fiber.ErrTooEarly,
    // ...
}))
```



## Config 配置

| Property 属性  | Type 输入               | Description 说明                                             | Default 默认                                                 |
| -------------- | ----------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Next 下一步    | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                                                        |
| IsEarlyData    | `func(*fiber.Ctx) bool` | IsEarlyData returns whether the request is an early-data request. IsEarlyData 返回请求是否为早期数据请求。 | Function checking if "Early-Data" header equals "1" 检查“Early-Data”标头是否等于“1”的函数 |
| AllowEarlyData | `func(*fiber.Ctx) bool` | AllowEarlyData returns whether the early-data request should be allowed or rejected. AllowEarlyData 返回是否应允许或拒绝早期数据请求。 | Function rejecting on unsafe and allowing safe methods 拒绝不安全方法并允许安全方法的函数 |
| Error 错误     | `error`                 | Error is returned in case an early-data request is rejected. 如果拒绝早期数据请求，则返回错误。 | `fiber.ErrTooEarly`                                          |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    IsEarlyData: func(c *fiber.Ctx) bool {
        return c.Get(DefaultHeaderName) == DefaultHeaderTrueValue
    },

    AllowEarlyData: func(c *fiber.Ctx) bool {
        return fiber.IsMethodSafe(c.Method())
    },

    Error: fiber.ErrTooEarly,
}
```



## Constants 常量 

```go
const (
    DefaultHeaderName      = "Early-Data"
    DefaultHeaderTrueValue = "1"
)
```
