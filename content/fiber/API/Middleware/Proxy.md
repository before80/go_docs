+++
title = "Proxy"
date = 2024-02-05T09:14:15+08:00
weight = 210
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/proxy]({{< ref "/fiber/API/Middleware/Proxy" >}})

# Proxy 代理

Proxy middleware for [Fiber](https://github.com/gofiber/fiber) that allows you to proxy requests to multiple servers.

​	Fiber 的代理中间件，允许您将请求代理到多个服务器。

## Signatures 签名

```go
// Balancer create a load balancer among multiple upstrem servers.
func Balancer(config Config) fiber.Handler
// Forward performs the given http request and fills the given http response.
func Forward(addr string, clients ...*fasthttp.Client) fiber.Handler
// Do performs the given http request and fills the given http response.
func Do(c *fiber.Ctx, addr string, clients ...*fasthttp.Client) error
// DoRedirects performs the given http request and fills the given http response while following up to maxRedirectsCount redirects.
func DoRedirects(c *fiber.Ctx, addr string, maxRedirectsCount int, clients ...*fasthttp.Client) error
// DoDeadline performs the given request and waits for response until the given deadline.
func DoDeadline(c *fiber.Ctx, addr string, deadline time.Time, clients ...*fasthttp.Client) error
// DoTimeout performs the given request and waits for response during the given timeout duration.
func DoTimeout(c *fiber.Ctx, addr string, timeout time.Duration, clients ...*fasthttp.Client) error
// DomainForward the given http request based on the given domain and fills the given http response
func DomainForward(hostname string, addr string, clients ...*fasthttp.Client) fiber.Handler
// BalancerForward performs the given http request based round robin balancer and fills the given http response
func BalancerForward(servers []string, clients ...*fasthttp.Client) fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/proxy"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// if target https site uses a self-signed certificate, you should
// call WithTlsConfig before Do and Forward
proxy.WithTlsConfig(&tls.Config{
    InsecureSkipVerify: true,
})
// if you need to use global self-custom client, you should use proxy.WithClient.
proxy.WithClient(&fasthttp.Client{
    NoDefaultUserAgentHeader: true, 
    DisablePathNormalizing:   true,
})

// Forward to url
app.Get("/gif", proxy.Forward("https://i.imgur.com/IWaBepg.gif"))

// If you want to forward with a specific domain. You have to use proxy.DomainForward.
app.Get("/payments", proxy.DomainForward("docs.gofiber.io", "http://localhost:8000"))

// Forward to url with local custom client
app.Get("/gif", proxy.Forward("https://i.imgur.com/IWaBepg.gif", &fasthttp.Client{
    NoDefaultUserAgentHeader: true, 
    DisablePathNormalizing:   true,
}))

// Make request within handler
app.Get("/:id", func(c *fiber.Ctx) error {
    url := "https://i.imgur.com/"+c.Params("id")+".gif"
    if err := proxy.Do(c, url); err != nil {
        return err
    }
    // Remove Server header from response
    c.Response().Header.Del(fiber.HeaderServer)
    return nil
})

// Make proxy requests while following redirects
app.Get("/proxy", func(c *fiber.Ctx) error {
    if err := proxy.DoRedirects(c, "http://google.com", 3); err != nil {
        return err
    }
    // Remove Server header from response
    c.Response().Header.Del(fiber.HeaderServer)
    return nil
})

// Make proxy requests and wait up to 5 seconds before timing out
app.Get("/proxy", func(c *fiber.Ctx) error {
    if err := proxy.DoTimeout(c, "http://localhost:3000", time.Second * 5); err != nil {
        return err
    }
    // Remove Server header from response
    c.Response().Header.Del(fiber.HeaderServer)
    return nil
})

// Make proxy requests, timeout a minute from now
app.Get("/proxy", func(c *fiber.Ctx) error {
    if err := proxy.DoDeadline(c, "http://localhost", time.Now().Add(time.Minute)); err != nil {
        return err
    }
    // Remove Server header from response
    c.Response().Header.Del(fiber.HeaderServer)
    return nil
})

// Minimal round robin balancer
app.Use(proxy.Balancer(proxy.Config{
    Servers: []string{
        "http://localhost:3001",
        "http://localhost:3002",
        "http://localhost:3003",
    },
}))

// Or extend your balancer for customization
app.Use(proxy.Balancer(proxy.Config{
    Servers: []string{
        "http://localhost:3001",
        "http://localhost:3002",
        "http://localhost:3003",
    },
    ModifyRequest: func(c *fiber.Ctx) error {
        c.Request().Header.Add("X-Real-IP", c.IP())
        return nil
    },
    ModifyResponse: func(c *fiber.Ctx) error {
        c.Response().Header.Del(fiber.HeaderServer)
        return nil
    },
}))

// Or this way if the balancer is using https and the destination server is only using http.
app.Use(proxy.BalancerForward([]string{
    "http://localhost:3001",
    "http://localhost:3002",
    "http://localhost:3003",
}))
```



## Config 配置

| Property 属性           | Type 输入                                                    | Description 说明                                             | Default 默认               |
| ----------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | -------------------------- |
| Next 下一步             | `func(*fiber.Ctx) bool`                                      | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                      |
| Servers 服务器          | `[]string`                                                   | Servers defines a list of `<scheme>://<host>` HTTP servers, which are used in a round-robin manner. i.e.: "[https://foobar.com](https://foobar.com/), [http://www.foobar.com"](http://www.foobar.com"/) 服务器定义了一个 `<scheme>://<host>` HTTP 服务器列表，这些服务器以循环方式使用。例如：“https://foobar.com, http://www.foobar.com” | (Required) （必需）        |
| ModifyRequest 修改请求  | `fiber.Handler`                                              | ModifyRequest allows you to alter the request. 修改请求允许您更改请求。 | `nil`                      |
| ModifyResponse 修改响应 | `fiber.Handler`                                              | ModifyResponse allows you to alter the response. 修改响应允许您更改响应。 | `nil`                      |
| Timeout 超时            | `time.Duration`                                              | Timeout is the request timeout used when calling the proxy client. 超时是在调用代理客户端时使用的请求超时。 | 1 second 1 秒              |
| ReadBufferSize          | `int`                                                        | Per-connection buffer size for requests' reading. This also limits the maximum header size. Increase this buffer if your clients send multi-KB RequestURIs and/or multi-KB headers (for example, BIG cookies). 用于读取请求的每个连接的缓冲区大小。这也限制了最大标头大小。如果您的客户端发送多 KB 的 RequestURI 和/或多 KB 的标头（例如，BIG cookie），请增加此缓冲区。 | (Not specified) （未指定） |
| WriteBufferSize         | `int`                                                        | Per-connection buffer size for responses' writing. 用于响应写入的每个连接的缓冲区大小。 | (Not specified) （未指定） |
| TlsConfig               | `*tls.Config` (or `*fasthttp.TLSConfig` in v3) `*tls.Config` （或 v3 中的 `*fasthttp.TLSConfig` ） | TLS config for the HTTP client. HTTP 客户端的 TLS 配置。     | `nil`                      |
| Client                  | `*fasthttp.LBClient`                                         | Client is a custom client when client config is complex. 当客户端配置复杂时，Client 是一个自定义客户端。 | `nil`                      |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next:           nil,
    ModifyRequest:  nil,
    ModifyResponse: nil,
    Timeout:        fasthttp.DefaultLBClientTimeout,
}
```
