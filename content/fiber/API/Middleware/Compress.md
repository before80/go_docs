+++
title = "Compress"
date = 2024-02-05T09:14:15+08:00
weight = 30
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/compress]({{< ref "/fiber/API/Middleware/Compress" >}})

# Compress

Compression middleware for [Fiber](https://github.com/gofiber/fiber) that will compress the response using `gzip`, `deflate` and `brotli` compression depending on the [Accept-Encoding](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding) header.

​	用于 Fiber 的压缩中间件，它将使用 `gzip` 、 `deflate` 和 `brotli` 压缩来压缩响应，具体取决于 Accept-Encoding 标头。

NOTE
注意

The compression middleware refrains from compressing bodies that are smaller than 200 bytes. This decision is based on the observation that, in such cases, the compressed size is likely to exceed the original size, making compression inefficient. [more](https://github.com/valyala/fasthttp/blob/497922a21ef4b314f393887e9c6147b8c3e3eda4/http.go#L1713-L1715)

​	压缩中间件不会压缩小于 200 字节的主体。此决定基于以下观察：在这种情况下，压缩后的尺寸很可能超过原始尺寸，从而使压缩效率低下。更多

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
  "github.com/gofiber/fiber/v2/middleware/compress"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(compress.New())

// Or extend your config for customization
app.Use(compress.New(compress.Config{
    Level: compress.LevelBestSpeed, // 1
}))

// Skip middleware for specific routes
app.Use(compress.New(compress.Config{
  Next:  func(c *fiber.Ctx) bool {
    return c.Path() == "/dont_compress"
  },
  Level: compress.LevelBestSpeed, // 1
}))
```



## Config 配置

### Config 配置

| Property 属性 | Type 输入               | Description 说明                                             | Default 默认       |
| ------------- | ----------------------- | ------------------------------------------------------------ | ------------------ |
| Next 下一步   | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`              |
| Level 级别    | `Level`                 | Level determines the compression algorithm. 级别确定压缩算法。 | `LevelDefault (0)` |

Possible values for the "Level" field are:

​	“级别”字段的可能值为：

- `LevelDisabled (-1)`: Compression is disabled.
  `LevelDisabled (-1)` ：禁用压缩。
- `LevelDefault (0)`: Default compression level.
  `LevelDefault (0)` ：默认压缩级别。
- `LevelBestSpeed (1)`: Best compression speed.
  `LevelBestSpeed (1)` ：最佳压缩速度。
- `LevelBestCompression (2)`: Best compression.
  `LevelBestCompression (2)` ：最佳压缩。

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next:  nil,
    Level: LevelDefault,
}
```



## Constants 常量 

```go
// Compression levels
const (
    LevelDisabled        = -1
    LevelDefault         = 0
    LevelBestSpeed       = 1
    LevelBestCompression = 2
)
```
