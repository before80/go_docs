+++
title = "Pprof"
date = 2024-02-05T09:14:15+08:00
weight = 200
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/pprof]({{< ref "/fiber/API/Middleware/Pprof" >}})

# Pprof

Pprof middleware for [Fiber](https://github.com/gofiber/fiber) that serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool. The package is typically only imported for the side effect of registering its HTTP handlers. The handled paths all begin with /debug/pprof/.

​	Fiber 的 Pprof 中间件，通过其 HTTP 服务器运行时配置文件数据提供 pprof 可视化工具预期的格式。通常仅导入该软件包以产生注册其 HTTP 处理程序的副作用。处理的路径均以 /debug/pprof/ 开头。

## Signatures 签名

```go
func New() fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/pprof"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(pprof.New())

// Or extend your config for customization

// For example, in systems where you have multiple ingress endpoints, it is common to add a URL prefix, like so:
app.Use(pprof.New(pprof.Config{Prefix: "/endpoint-prefix"}))

// This prefix will be added to the default path of "/debug/pprof/", for a resulting URL of: "/endpoint-prefix/debug/pprof/".
```



## Config 配置

| Property 属性 | Type 输入               | Description 说明                                             | Default 默认 |
| ------------- | ----------------------- | ------------------------------------------------------------ | ------------ |
| Next 下一步   | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`        |
| Prefix 前缀   | `string`                | Prefix defines a URL prefix added before "/debug/pprof". Note that it should start with (but not end with) a slash. Example: "/federated-fiber" 前缀定义了在 "/debug/pprof" 之前添加的 URL 前缀。请注意，它应以斜杠开头（但不以斜杠结尾）。示例："/federated-fiber" | ""           |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next: nil,
}
```
