+++
title = "Welcome"
date = 2024-02-05T09:14:15+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/]({{< ref "/fiber/Welcome" >}})

# 👋 Welcome

An online API documentation with examples so you can start building web apps with Fiber right away!

​	一个在线 API 文档，其中包含示例，以便您可以立即开始使用 Fiber 构建 Web 应用！

**Fiber** is an [Express](https://github.com/expressjs/express) inspired **web framework** built on top of [Fasthttp](https://github.com/valyala/fasthttp), the **fastest** HTTP engine for [Go](https://go.dev/doc/). Designed to **ease** things up for **fast** development with **zero memory allocation** and **performance** in mind.

​	Fiber 是一个基于 Fasthttp（Go 最快的 HTTP 引擎）构建的受 Express 启发的 Web 框架。旨在简化快速开发，同时兼顾零内存分配和性能。

These docs are for **Fiber v2**, which was released on **September 15th, 2020**.

​	这些文档适用于 Fiber v2，该版本于 2020 年 9 月 15 日发布。

### Installation 安装

First of all, [download](https://go.dev/dl/) and install Go. `1.17` or higher is required.

​	首先，下载并安装 Go。需要 `1.17` 或更高版本。

Installation is done using the [`go get`](https://pkg.go.dev/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) command:

​	安装使用 `go get` 命令完成：

```bash
go get github.com/gofiber/fiber/v2
```



### Zero Allocation 零分配

Some values returned from ***fiber.Ctx** are **not** immutable by default.

​	从 *fiber.Ctx 返回的某些值默认情况下不是不可变的。

Because fiber is optimized for **high-performance**, values returned from **fiber.Ctx** are **not** immutable by default and **will** be re-used across requests. As a rule of thumb, you **must** only use context values within the handler, and you **must not** keep any references. As soon as you return from the handler, any values you have obtained from the context will be re-used in future requests and will change below your feet. Here is an example:

​	由于 fiber 针对高性能进行了优化，因此从 fiber.Ctx 返回的值默认情况下不是不可变的，并且将在请求之间重复使用。根据经验，您只能在处理程序中使用上下文值，并且不能保留任何引用。一旦您从处理程序返回，您从上下文中获取的任何值都将在将来的请求中重复使用，并且会在您脚下发生变化。这里有一个例子：

```go
func handler(c *fiber.Ctx) error {
    // Variable is only valid within this handler
    result := c.Params("foo") 

    // ...
}
```



If you need to persist such values outside the handler, make copies of their **underlying buffer** using the [copy](https://pkg.go.dev/builtin/#copy) builtin. Here is an example for persisting a string:

​	如果您需要在处理程序外部保留此类值，请使用 copy 内置函数复制其底层缓冲区。这里有一个持久化字符串的示例：

```go
func handler(c *fiber.Ctx) error {
    // Variable is only valid within this handler
    result := c.Params("foo")

    // Make a copy
    buffer := make([]byte, len(result))
    copy(buffer, result)
    resultCopy := string(buffer) 
    // Variable is now valid forever

    // ...
}
```



We created a custom `CopyString` function that does the above and is available under [gofiber/utils](https://github.com/gofiber/fiber/tree/master/utils).

​	我们创建了一个自定义 `CopyString` 函数来执行上述操作，该函数在 gofiber/utils 下可用。

```go
app.Get("/:foo", func(c *fiber.Ctx) error {
    // Variable is now immutable
    result := utils.CopyString(c.Params("foo")) 

    // ...
})
```



Alternatively, you can also use the `Immutable` setting. It will make all values returned from the context immutable, allowing you to persist them anywhere. Of course, this comes at the cost of performance.

​	或者，您也可以使用 `Immutable` 设置。它会使从上下文返回的所有值不可变，允许您将它们持久存储在任何地方。当然，这是以牺牲性能为代价的。

```go
app := fiber.New(fiber.Config{
    Immutable: true,
})
```



For more information, please check [**#426**](https://github.com/gofiber/fiber/issues/426) and [**#185**](https://github.com/gofiber/fiber/issues/185).

​	有关更多信息，请查看 #426 和 #185。

### Hello, World! 你好，世界！

Embedded below is essentially the most straightforward **Fiber** app you can create:

​	下面嵌入的是您可以创建的最简单的 Fiber 应用程序：

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}
```



```text
go run server.go
```



Browse to `http://localhost:3000` and you should see `Hello, World!` on the page.

​	浏览到 `http://localhost:3000` ，您应该会在页面上看到 `Hello, World!` 。

### Basic routing 基本路由

Routing refers to determining how an application responds to a client request to a particular endpoint, which is a URI (or path) and a specific HTTP request method (`GET`, `PUT`, `POST`, etc.).

​	路由是指确定应用程序如何响应客户端对特定端点的请求，该端点是一个 URI（或路径）和一个特定的 HTTP 请求方法（ `GET` 、 `PUT` 、 `POST` 等）。

Each route can have **multiple handler functions** that are executed when the route is matched.

​	每个路由可以有多个处理程序函数，当路由匹配时执行这些函数。

Route definition takes the following structures:

​	路由定义采用以下结构：

```go
// Function signature
app.Method(path string, ...func(*fiber.Ctx) error)
```



- `app` is an instance of **Fiber**
  `app` 是 Fiber 的一个实例
- `Method` is an [HTTP request method]({{< ref "/fiber/API/App#route-handlers" >}}): `GET`, `PUT`, `POST`, etc.
  `Method` 是一种HTTP请求方法： `GET` 、 `PUT` 、 `POST` 等。
- `path` is a virtual path on the server
  `path` 是服务器上的虚拟路径
- `func(*fiber.Ctx) error` is a callback function containing the [Context]({{< ref "/fiber/API/Ctx" >}}) executed when the route is matched
  `func(*fiber.Ctx) error` 是一个回调函数，其中包含在匹配路由时执行的上下文

**Simple route
简单路由**

```go
// Respond with "Hello, World!" on root path, "/"
app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})
```



**Parameters
参数**

```go
// GET http://localhost:8080/hello%20world

app.Get("/:value", func(c *fiber.Ctx) error {
    return c.SendString("value: " + c.Params("value"))
    // => Get request with value: hello world
})
```



**Optional parameter
可选参数**

```go
// GET http://localhost:3000/john

app.Get("/:name?", func(c *fiber.Ctx) error {
    if c.Params("name") != "" {
        return c.SendString("Hello " + c.Params("name"))
        // => Hello john
    }
    return c.SendString("Where is john?")
})
```



**Wildcards
通配符**

```go
// GET http://localhost:3000/api/user/john

app.Get("/api/*", func(c *fiber.Ctx) error {
    return c.SendString("API path: " + c.Params("*"))
    // => API path: user/john
})
```



### Static files 静态文件

To serve static files such as **images**, **CSS**, and **JavaScript** files, replace your function handler with a file or directory string.

​	要提供图像、CSS 和 JavaScript 文件等静态文件，请将您的函数处理程序替换为文件或目录字符串。

Function signature:

​	函数签名：

```go
app.Static(prefix, root string, config ...Static)
```



Use the following code to serve files in a directory named `./public`:

​	使用以下代码在名为 `./public` 的目录中提供文件：

```go
app := fiber.New()

app.Static("/", "./public") 

app.Listen(":3000")
```



Now, you can load the files that are in the `./public` directory:

​	现在，您可以加载位于 `./public` 目录中的文件：

```bash
http://localhost:3000/hello.html
http://localhost:3000/js/jquery.js
http://localhost:3000/css/style.css
```



### Note 注意

For more information on how to build APIs in Go with Fiber, please check out this excellent article [on building an express-style API in Go with Fiber](https://blog.logrocket.com/express-style-api-go-fiber/).

​	有关如何在 Go 中使用 Fiber 构建 API 的更多信息，请查看这篇关于使用 Fiber 在 Go 中构建 express 风格 API 的优秀文章。
