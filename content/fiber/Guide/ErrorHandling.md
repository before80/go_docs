+++
title = "错误处理"
date = 2024-02-05T09:14:15+08:00
weight = 30
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/guide/error-handling]({{< ref "/fiber/Guide/ErrorHandling" >}})

# 🐛 Error Handling  处理错误

## Catching Errors 捕获错误

It’s essential to ensure that Fiber catches all errors that occur while running route handlers and middleware. You must return them to the handler function, where Fiber will catch and process them.

​	确保 Fiber 捕获在运行路由处理程序和中间件时发生的错误非常重要。您必须将它们返回给处理程序函数，Fiber 将捕获并处理它们。

- Example
  示例

```go
app.Get("/", func(c *fiber.Ctx) error {
    // Pass error to Fiber
    return c.SendFile("file-does-not-exist")
})
```



Fiber does not handle [panics](https://go.dev/blog/defer-panic-and-recover) by default. To recover from a panic thrown by any handler in the stack, you need to include the `Recover` middleware below:

​	默认情况下，Fiber 不处理恐慌。要从堆栈中任何处理程序抛出的恐慌中恢复，您需要包含以下 `Recover` 中间件：

Example
示例

```go
package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
    app := fiber.New()

    app.Use(recover.New())

    app.Get("/", func(c *fiber.Ctx) error {
        panic("This panic is caught by fiber")
    })

    log.Fatal(app.Listen(":3000"))
}
```



You could use Fiber's custom error struct to pass an additional `status code` using `fiber.NewError()`. It's optional to pass a message; if this is left empty, it will default to the status code message (`404` equals `Not Found`).

​	您可以使用 Fiber 的自定义错误结构来传递附加的 `status code` ，方法是使用 `fiber.NewError()` 。传递消息是可选的；如果留空，它将默认为状态代码消息（ `404` 等于 `Not Found` ）。

Example
示例

```go
app.Get("/", func(c *fiber.Ctx) error {
    // 503 Service Unavailable
    return fiber.ErrServiceUnavailable

    // 503 On vacation!
    return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
})
```



## Default Error Handler 默认错误处理程序

Fiber provides an error handler by default. For a standard error, the response is sent as **500 Internal Server Error**. If the error is of type [fiber.Error](https://godoc.org/github.com/gofiber/fiber#Error), the response is sent with the provided status code and message.

​	Fiber 默认提供错误处理程序。对于标准错误，响应将作为 500 内部服务器错误发送。如果错误的类型是 fiber.Error，则响应将随提供的状态代码和消息一起发送。

Example
示例

```go
// Default error handler
var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
    // Status code defaults to 500
    code := fiber.StatusInternalServerError

    // Retrieve the custom status code if it's a *fiber.Error
    var e *fiber.Error
    if errors.As(err, &e) {
        code = e.Code
    }

    // Set Content-Type: text/plain; charset=utf-8
    c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

    // Return status code with error message
    return c.Status(code).SendString(err.Error())
}
```



## Custom Error Handler 自定义错误处理程序 

A custom error handler can be set using a [Config ]({{< ref "/fiber/API/Fiber#config" >}})when initializing a [Fiber instance]({{< ref "/fiber/API/Fiber#new" >}}).

​	在初始化 Fiber 实例时，可以使用 Config 设置自定义错误处理程序。

In most cases, the default error handler should be sufficient. However, a custom error handler can come in handy if you want to capture different types of errors and take action accordingly e.g., send a notification email or log an error to the centralized system. You can also send customized responses to the client e.g., error page or just a JSON response.

​	在大多数情况下，默认错误处理程序应该足够了。但是，如果您想捕获不同类型的错误并相应地采取措施（例如，发送通知电子邮件或将错误记录到集中式系统），那么自定义错误处理程序会派上用场。您还可以向客户端发送自定义响应，例如错误页面或仅 JSON 响应。

The following example shows how to display error pages for different types of errors.

​	以下示例演示如何为不同类型的错误显示错误页面。

Example
示例

```go
// Create a new fiber instance with custom config
app := fiber.New(fiber.Config{
    // Override default error handler
    ErrorHandler: func(ctx *fiber.Ctx, err error) error {
        // Status code defaults to 500
        code := fiber.StatusInternalServerError

        // Retrieve the custom status code if it's a *fiber.Error
        var e *fiber.Error
        if errors.As(err, &e) {
            code = e.Code
        }

        // Send custom error page
        err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
        if err != nil {
            // In case the SendFile fails
            return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
        }

        // Return from handler
        return nil
    },
})

// ...
```



> Special thanks to the [Echo]({{< ref "/echo">}}) & [Express](https://expressjs.com/) framework for inspiration regarding error handling.
>
> ​	特别感谢 Echo & Express 框架在错误处理方面的启发。
