+++
title = "Skip"
date = 2024-02-05T09:14:15+08:00
weight = 270
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/skip]({{< ref "/fiber/API/Middleware/Skip" >}})

# Skip 跳过

Skip middleware for [Fiber](https://github.com/gofiber/fiber) that skips a wrapped handler if a predicate is true.

​	Fiber 的跳过中间件，如果谓词为真，则跳过包装的处理程序。

## Signatures 签名

```go
func New(handler fiber.Handler, exclude func(c *fiber.Ctx) bool) fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/skip"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
func main() {
    app := fiber.New()

    app.Use(skip.New(BasicHandler, func(ctx *fiber.Ctx) bool {
        return ctx.Method() == fiber.MethodGet
    }))

    app.Get("/", func(ctx *fiber.Ctx) error {
        return ctx.SendString("It was a GET request!")
    })

    log.Fatal(app.Listen(":3000"))
}

func BasicHandler(ctx *fiber.Ctx) error {
    return ctx.SendString("It was not a GET request!")
}
```



TIP

app.Use will handle requests from any route, and any method. In the example above, it will only skip if the method is GET.

​	app.Use 将处理来自任何路由和任何方法的请求。在上面的示例中，它仅在方法为 GET 时才会跳过。
