+++
title = "模板"
date = 2024-02-05T09:14:15+08:00
weight = 20
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/guide/templates]({{< ref "/fiber/Guide/Templates" >}})

# 📝 Templates  模板

## Template interfaces 模板接口 

Fiber provides a Views interface to provide your own template engine:

​	Fiber 提供了一个 Views 接口，以便您提供自己的模板引擎：

- Views
  视图 Beego 的 MVC 简介 MVC 简介 MVC 简介 Beego 的 MVC 简介 Beego 使用典型的模型-视图-控制器 (MVC) 框架。此图说明了如何处理请求处理逻辑： 整个逻辑处理过程如下所述： 数据从侦听端口接收。侦听端口默认设置为 8080。 请求到达端口 8080 后，Beego 开始处理请求的数据

```go
type Views interface {
    Load() error
    Render(io.Writer, string, interface{}, ...string) error
}
```



`Views` interface contains a `Load` and `Render` method, `Load` is executed by Fiber on app initialization to load/parse the templates.

​	 `Views` 接口包含一个 `Load` 和 `Render` 方法， `Load` 由 Fiber 在应用程序初始化时执行，以加载/解析模板。

```go
// Pass engine to Fiber's Views Engine
app := fiber.New(fiber.Config{
    Views: engine,
    // Views Layout is the global layout for all template render until override on Render function.
    ViewsLayout: "layouts/main"
})
```



The `Render` method is linked to the [**ctx.Render()**]({{< ref "/fiber/API/Ctx#render" >}}) function that accepts a template name and binding data. It will use global layout if layout is not being defined in `Render` function. If the Fiber config option `PassLocalsToViews` is enabled, then all locals set using `ctx.Locals(key, value)` will be passed to the template.

​	 `Render` 方法链接到 ctx.Render() 函数，该函数接受模板名称和绑定数据。如果在 `Render` 函数中未定义布局，它将使用全局布局。如果启用了 Fiber 配置选项 `PassLocalsToViews` ，则使用 `ctx.Locals(key, value)` 设置的所有本地变量都将传递给模板。

```go
app.Get("/", func(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{
        "hello": "world",
    });
})
```



## Engines 引擎 

Fiber team maintains [templates](https://docs.gofiber.io/template) package that provides wrappers for multiple template engines:

​	Fiber 团队维护模板包，该包为多个模板引擎提供包装器：

- [ace](https://docs.gofiber.io/template/ace/)
- [amber](https://docs.gofiber.io/template/amber/)
- [django](https://docs.gofiber.io/template/django/)
- [handlebars](https://docs.gofiber.io/template/handlebars)
- [html](https://docs.gofiber.io/template/html)
- [jet](https://docs.gofiber.io/template/jet)
- [mustache](https://docs.gofiber.io/template/mustache)
- [pug](https://docs.gofiber.io/template/pug)
- [slim](https://docs.gofiber.io/template/slim)

- Example
  示例
- views/index.html

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
)

func main() {
    // Initialize standard Go html template engine
    engine := html.New("./views", ".html")
    // If you want other engine, just replace with following
    // Create a new engine with django
    // engine := django.New("./views", ".django")

    app := fiber.New(fiber.Config{
        Views: engine,
    })
    app.Get("/", func(c *fiber.Ctx) error {
        // Render index template
        return c.Render("index", fiber.Map{
            "Title": "Hello, World!",
        })
    })

    log.Fatal(app.Listen(":3000"))
}
```
