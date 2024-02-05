+++
title = "Grouping"
date = 2024-02-05T09:14:15+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/guide/grouping]({{< ref "/fiber/Guide/Grouping" >}})

# 🎭 Grouping

INFO
信息

In general, the Group functionality in Fiber behaves similarly to ExpressJS. Groups are declared virtually and all routes declared within the group are flattened into a single list with a prefix, which is then checked by the framework in the order it was declared. This means that the behavior of Group in Fiber is identical to that of ExpressJS.

​	通常，Fiber 中的 Group 功能的行为与 ExpressJS 类似。组是虚拟声明的，组内声明的所有路由都展平为一个带有前缀的单一列表，然后由框架按声明顺序进行检查。这意味着 Fiber 中 Group 的行为与 ExpressJS 中的行为相同。

## Paths 路径

Like **Routing**, groups can also have paths that belong to a cluster.

​	与路由一样，组也可以具有属于集群的路径。

```go
func main() {
  app := fiber.New()

  api := app.Group("/api", middleware) // /api

  v1 := api.Group("/v1", middleware)   // /api/v1
  v1.Get("/list", handler)             // /api/v1/list
  v1.Get("/user", handler)             // /api/v1/user

  v2 := api.Group("/v2", middleware)   // /api/v2
  v2.Get("/list", handler)             // /api/v2/list
  v2.Get("/user", handler)             // /api/v2/user

  log.Fatal(app.Listen(":3000"))
}
```



A **Group** of paths can have an optional handler.

​	一组路径可以具有一个可选的处理程序。

```go
func main() {
  app := fiber.New()

  api := app.Group("/api")      // /api

  v1 := api.Group("/v1")        // /api/v1
  v1.Get("/list", handler)      // /api/v1/list
  v1.Get("/user", handler)      // /api/v1/user

  v2 := api.Group("/v2")        // /api/v2
  v2.Get("/list", handler)      // /api/v2/list
  v2.Get("/user", handler)      // /api/v2/user

  log.Fatal(app.Listen(":3000"))
}
```



CAUTION
注意

Running **/api**, **/v1** or **/v2** will result in **404** error, make sure you have the errors set.

​	运行 /api、/v1 或 /v2 将导致 404 错误，请确保已设置错误。

## Group Handlers 组处理程序 

Group handlers can also be used as a routing path but they must have **Next** added to them so that the flow can continue.

​	组处理程序也可以用作路由路径，但必须向它们添加 Next，以便流可以继续。

```go
func main() {
    app := fiber.New()

    handler := func(c *fiber.Ctx) error {
        return c.SendStatus(fiber.StatusOK)
    }
    api := app.Group("/api") // /api

    v1 := api.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
        c.Set("Version", "v1")
        return c.Next()
    })
    v1.Get("/list", handler) // /api/v1/list
    v1.Get("/user", handler) // /api/v1/user

    log.Fatal(app.Listen(":3000"))
}
```
