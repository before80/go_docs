+++
title = "FileSystem"
date = 2024-02-05T09:14:15+08:00
weight = 120
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/filesystem]({{< ref "/fiber/API/Middleware/FileSystem" >}})

# FileSystem 文件系统

Filesystem middleware for [Fiber](https://github.com/gofiber/fiber) that enables you to serve files from a directory.

​	Filesystem 中间件，用于 Fiber，它使您能够从目录提供文件。

CAUTION
注意

**`:params` & `:optionals?` within the prefix path are not supported!
前缀路径中的 `:params` 和 `:optionals?` 不受支持！**

**To handle paths with spaces (or other url encoded values) make sure to set `fiber.Config{ UnescapePath: true }`
要处理带有空格（或其他 url 编码值）的路径，请确保设置 `fiber.Config{ UnescapePath: true }`**

## Signatures 签名

```go
func New(config Config) fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Provide a minimal config
app.Use(filesystem.New(filesystem.Config{
    Root: http.Dir("./assets"),
}))

// Or extend your config for customization
app.Use(filesystem.New(filesystem.Config{
    Root:         http.Dir("./assets"),
    Browse:       true,
    Index:        "index.html",
    NotFoundFile: "404.html",
    MaxAge:       3600,
}))
```



> If your environment (Go 1.16+) supports it, we recommend using Go Embed instead of the other solutions listed as this one is native to Go and the easiest to use.
>
> ​	如果您的环境（Go 1.16+）支持它，我们建议使用 Go Embed 而不是列出的其他解决方案，因为它是 Go 的原生解决方案，也是最容易使用的。

## embed embed Embed 是将文件嵌入到 Golang 可执行文件中的原生方法。在 Go 1.16 中引入。

[Embed](https://golang.org/pkg/embed/) is the native method to embed files in a Golang excecutable. Introduced in Go 1.16.

​	pkger 

```go
package main

import (
    "embed"
    "io/fs"
    "log"
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"
)

// Embed a single file
//go:embed index.html
var f embed.FS

// Embed a directory
//go:embed static/*
var embedDirStatic embed.FS

func main() {
    app := fiber.New()

    app.Use("/", filesystem.New(filesystem.Config{
        Root: http.FS(f),
    }))

    // Access file "image.png" under `static/` directory via URL: `http://<server>/static/image.png`.
    // Without `PathPrefix`, you have to access it via URL:
    // `http://<server>/static/static/image.png`.
    app.Use("/static", filesystem.New(filesystem.Config{
        Root: http.FS(embedDirStatic),
        PathPrefix: "static",
        Browse: true,
    }))

    log.Fatal(app.Listen(":3000"))
}
```



## pkger packr 

https://github.com/markbates/pkger

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"

    "github.com/markbates/pkger"
)

func main() {
    app := fiber.New()

    app.Use("/assets", filesystem.New(filesystem.Config{
        Root: pkger.Dir("/assets"),
    }))

    log.Fatal(app.Listen(":3000"))
}
```



## packr go.rice 

https://github.com/gobuffalo/packr

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"

    "github.com/gobuffalo/packr/v2"
)

func main() {
    app := fiber.New()

    app.Use("/assets", filesystem.New(filesystem.Config{
        Root: packr.New("Assets Box", "/assets"),
    }))

    log.Fatal(app.Listen(":3000"))
}
```



## go.rice fileb0x 

https://github.com/GeertJohan/go.rice

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"

    "github.com/GeertJohan/go.rice"
)

func main() {
    app := fiber.New()

    app.Use("/assets", filesystem.New(filesystem.Config{
        Root: rice.MustFindBox("assets").HTTPBox(),
    }))

    log.Fatal(app.Listen(":3000"))
}
```



## fileb0x statik 

https://github.com/UnnoTed/fileb0x

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"

    "<Your go module>/myEmbeddedFiles"
)

func main() {
    app := fiber.New()

    app.Use("/assets", filesystem.New(filesystem.Config{
        Root: myEmbeddedFiles.HTTP,
    }))

    log.Fatal(app.Listen(":3000"))
}
```



## statik

https://github.com/rakyll/statik

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"

    // Use blank to invoke init function and register data to statik
    _ "<Your go module>/statik" 
    "github.com/rakyll/statik/fs"
)

func main() {
    statikFS, err := fs.New()
    if err != nil {
        panic(err)
    }

    app := fiber.New()

    app.Use("/", filesystem.New(filesystem.Config{
        Root: statikFS,
    }))

    log.Fatal(app.Listen(":3000"))
}
```



## Config 配置

| Property 属性       | Type 输入               | Description 说明                                             | Default 默认              |
| ------------------- | ----------------------- | ------------------------------------------------------------ | ------------------------- |
| Next 下一步         | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                     |
| Root 根             | `http.FileSystem`       | Root is a FileSystem that provides access to a collection of files and directories. 根是一个文件系统，它提供对文件和目录集合的访问。 | `nil`                     |
| PathPrefix 路径前缀 | `string`                | PathPrefix defines a prefix to be added to a filepath when reading a file from the FileSystem. 路径前缀定义在从文件系统读取文件时要添加到文件路径的前缀。 | ""                        |
| Browse 浏览         | `bool`                  | Enable directory browsing. 启用目录浏览。                    | `false`                   |
| Index 索引          | `string`                | Index file for serving a directory. 用于提供目录的索引文件。 | "index.html" “index.html” |
| MaxAge              | `int`                   | The value for the Cache-Control HTTP-header that is set on the file response. MaxAge is defined in seconds. 在文件响应中设置的 Cache-Control HTTP 标头值。MaxAge 以秒为单位定义。 | 0                         |
| NotFoundFile        | `string`                | File to return if the path is not found. Useful for SPA's. 如果未找到路径，则返回文件。适用于 SPA。 | ""                        |
| ContentTypeCharset  | `string`                | The value for the Content-Type HTTP-header that is set on the file response. 在文件响应中设置的 Content-Type HTTP 标头的值。 | ""                        |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next:   nil,
    Root:   nil,
    PathPrefix: "",
    Browse: false,
    Index:  "/index.html",
    MaxAge: 0,
    ContentTypeCharset: "",
}
```



## Utils

### SendFile

Serves a file from an [HTTP file system](https://pkg.go.dev/net/http#FileSystem) at the specified path.

​	从指定路径的 HTTP 文件系统提供文件。

Signature
签名

```go
func SendFile(c *fiber.Ctx, filesystem http.FileSystem, path string) error
```



Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"
)
```



Example
示例

```go
// Define a route to serve a specific file
app.Get("/download", func(c *fiber.Ctx) error {
    // Serve the file using SendFile function
    err := filesystem.SendFile(c, http.Dir("your/filesystem/root"), "path/to/your/file.txt")
    if err != nil {
        // Handle the error, e.g., return a 404 Not Found response
        return c.Status(fiber.StatusNotFound).SendString("File not found")
    }
    
    return nil
})
```



Example
示例

```go
// Serve static files from the "build" directory using Fiber's built-in middleware.
app.Use("/", filesystem.New(filesystem.Config{
    Root:       http.FS(f),         // Specify the root directory for static files.
    PathPrefix: "build",            // Define the path prefix where static files are served.
}))

// For all other routes (wildcard "*"), serve the "index.html" file from the "build" directory.
app.Use("*", func(ctx *fiber.Ctx) error {
    return filesystem.SendFile(ctx, http.FS(f), "build/index.html")
})
```
