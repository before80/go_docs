+++
title = "Helmet"
date = 2024-02-05T09:14:15+08:00
weight = 140
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/helmet]({{< ref "/fiber/API/Middleware/Helmet" >}})

# Helmet

Helmet middleware helps secure your apps by setting various HTTP headers.

​	Helmet 中间件通过设置各种 HTTP 头来帮助保护您的应用。

## Signatures 签名

```go
func New(config ...Config) fiber.Handler
```



## Examples 示例 

```go
package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {
  app := fiber.New()

  app.Use(helmet.New())

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Welcome!")
  })

  app.Listen(":3000")
}
```



**Test:
测试：**

```curl
curl -I http://localhost:3000
```



## Config 配置

| Property 属性             | Type 输入               | Description 说明                                             | Default 默认              |
| ------------------------- | ----------------------- | ------------------------------------------------------------ | ------------------------- |
| Next 下一步               | `func(*fiber.Ctx) bool` | Next defines a function to skip middleware. 接下来定义一个跳过中间件的函数。 | `nil`                     |
| XSSProtection             | `string`                | XSSProtection                                                | "0"                       |
| ContentTypeNosniff        | `string`                | ContentTypeNosniff                                           | "nosniff" “nosniff”       |
| XFrameOptions             | `string`                | XFrameOptions                                                | "SAMEORIGIN" “SAMEORIGIN” |
| HSTSMaxAge                | `int`                   | HSTSMaxAge HSTS 最大生存期                                   | 0                         |
| HSTSExcludeSubdomains     | `bool`                  | HSTSExcludeSubdomains HSTS 排除子域                          | false                     |
| ContentSecurityPolicy     | `string`                | ContentSecurityPolicy 内容安全策略                           | ""                        |
| CSPReportOnly             | `bool`                  | CSPReportOnly CSP 仅报告                                     | false                     |
| HSTSPreloadEnabled        | `bool`                  | HSTSPreloadEnabled 启用 HSTS 预加载                          | false                     |
| ReferrerPolicy            | `string`                | ReferrerPolicy                                               | "ReferrerPolicy"          |
| PermissionPolicy          | `string`                | Permissions-Policy                                           | ""                        |
| CrossOriginEmbedderPolicy | `string`                | Cross-Origin-Embedder-Policy                                 | "require-corp"            |
| CrossOriginOpenerPolicy   | `string`                | Cross-Origin-Opener-Policy                                   | "same-origin" “同源”      |
| CrossOriginResourcePolicy | `string`                | Cross-Origin-Resource-Policy                                 | "same-origin" “同源”      |
| OriginAgentCluster        | `string`                | Origin-Agent-Cluster                                         | "?1" “?1”                 |
| XDNSPrefetchControl       | `string`                | X-DNS-Prefetch-Control                                       | "off" “off”               |
| XDownloadOptions          | `string`                | X-Download-Options                                           | "noopen"                  |
| XPermittedCrossDomain     | `string`                | X-Permitted-Cross-Domain-Policies                            | "none"                    |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    XSSProtection:             "0",
    ContentTypeNosniff:        "nosniff",
    XFrameOptions:             "SAMEORIGIN",
    ReferrerPolicy:            "no-referrer",
    CrossOriginEmbedderPolicy: "require-corp",
    CrossOriginOpenerPolicy:   "same-origin",
    CrossOriginResourcePolicy: "same-origin",
    OriginAgentCluster:        "?1",
    XDNSPrefetchControl:       "off",
    XDownloadOptions:          "noopen",
    XPermittedCrossDomain:     "none",
}
```
