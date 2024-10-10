+++
title = "Favicon"
date = 2024-02-05T09:14:15+08:00
weight = 110
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/favicon]({{< ref "/fiber/API/Middleware/Favicon" >}})

# Favicon 网站图标

Favicon middleware for [Fiber](https://github.com/gofiber/fiber) that ignores favicon requests or caches a provided icon in memory to improve performance by skipping disk access. User agents request favicon.ico frequently and indiscriminately, so you may wish to exclude these requests from your logs by using this middleware before your logger middleware.

​	Favicon 中间件，用于 Fiber，它忽略 favicon 请求或将提供的图标缓存到内存中，以通过跳过磁盘访问来提高性能。用户代理经常且不加区别地请求 favicon.ico，因此您可能希望通过在记录器中间件之前使用此中间件来从日志中排除这些请求。

NOTE
注意

This middleware is exclusively for serving the default, implicit favicon, which is GET /favicon.ico or [custom favicon URL](https://docs.gofiber.io/api/middleware/favicon/#config).

​	此中间件专门用于提供默认的隐式 favicon，即 GET /favicon.ico 或自定义 favicon URL。

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
  "github.com/gofiber/fiber/v2/middleware/favicon"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(favicon.New())

// Or extend your config for customization
app.Use(favicon.New(favicon.Config{
    File: "./favicon.ico",
    URL: "/favicon.ico",
}))
```



## Config 配置

| Property 属性       | Type 输入               | Description 说明                                             | Default 默认               |
| ------------------- | ----------------------- | ------------------------------------------------------------ | -------------------------- |
| Next 下一步         | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                      |
| Data 数据           | `[]byte`                | Raw data of the favicon file. This can be used instead of `File`. favicon 文件的原始数据。可以使用此数据代替 `File` 。 | `nil`                      |
| File                | `string`                | File holds the path to an actual favicon that will be cached. 文件保存将被缓存的实际 favicon 的路径。 | ""                         |
| URL                 | `string`                | URL for favicon handler. favicon 处理程序的 URL。            | "/favicon.ico"             |
| FileSystem 文件系统 | `http.FileSystem`       | FileSystem is an optional alternate filesystem to search for the favicon in. 文件系统是用于搜索 favicon 的可选备用文件系统。 | `nil`                      |
| CacheControl        | `string`                | CacheControl defines how the Cache-Control header in the response should be set. CacheControl 定义如何设置响应中的 Cache-Control 头。 | "public, max-age=31536000" |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next:         nil,
    File:         "",
    URL:          fPath,
    CacheControl: "public, max-age=31536000",
}
```
