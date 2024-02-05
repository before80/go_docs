+++
title = "Hooks"
date = 2024-02-05T09:14:15+08:00
weight = 50
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/guide/hooks]({{< ref "/fiber/Guide/Hooks" >}})

# 🎣 Hooks

With Fiber v2.30.0, you can execute custom user functions when to run some methods. Here is a list of this hooks:

​	在 Fiber v2.30.0 中，您可以在运行某些方法时执行自定义用户函数。以下是这些钩子的列表：

- [OnRoute](https://docs.gofiber.io/guide/hooks/#onroute)
- [OnName](https://docs.gofiber.io/guide/hooks/#onname)
- [OnGroup](https://docs.gofiber.io/guide/hooks/#ongroup)
- [OnGroupName](https://docs.gofiber.io/guide/hooks/#ongroupname)
- [OnListen](https://docs.gofiber.io/guide/hooks/#onlisten)
- [OnFork](https://docs.gofiber.io/guide/hooks/#onfork)
- [OnShutdown](https://docs.gofiber.io/guide/hooks/#onshutdown)
- [OnMount](https://docs.gofiber.io/guide/hooks/#onmount)

## Constants 常量 

```go
// Handlers define a function to create hooks for Fiber.
type OnRouteHandler = func(Route) error
type OnNameHandler = OnRouteHandler
type OnGroupHandler = func(Group) error
type OnGroupNameHandler = OnGroupHandler
type OnListenHandler = func(ListenData) error
type OnForkHandler = func(int) error
type OnShutdownHandler = func() error
type OnMountHandler = func(*App) error
```



## OnRoute

OnRoute is a hook to execute user functions on each route registeration. Also you can get route properties by **route** parameter.

​	OnRoute 是一个钩子，用于在每个路由注册上执行用户函数。您还可以通过路由参数获取路由属性。

Signature
签名

```go
func (h *Hooks) OnRoute(handler ...OnRouteHandler)
```



## OnName

OnName is a hook to execute user functions on each route naming. Also you can get route properties by **route** parameter.

​	OnName 是一个钩子，用于在每个路由命名上执行用户函数。您还可以通过路由参数获取路由属性。

CAUTION
注意

OnName only works with naming routes, not groups.

​	OnName 仅适用于命名路由，不适用于组。

Signature
签名

```go
func (h *Hooks) OnName(handler ...OnNameHandler)
```



- OnName Example
  OnName 示例

```go
package main

import (
    "fmt"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString(c.Route().Name)
    }).Name("index")

    app.Hooks().OnName(func(r fiber.Route) error {
        fmt.Print("Name: " + r.Name + ", ")

        return nil
    })

    app.Hooks().OnName(func(r fiber.Route) error {
        fmt.Print("Method: " + r.Method + "\n")

        return nil
    })

    app.Get("/add/user", func(c *fiber.Ctx) error {
        return c.SendString(c.Route().Name)
    }).Name("addUser")

    app.Delete("/destroy/user", func(c *fiber.Ctx) error {
        return c.SendString(c.Route().Name)
    }).Name("destroyUser")

    app.Listen(":5000")
}

// Results:
// Name: addUser, Method: GET
// Name: destroyUser, Method: DELETE
```



## OnGroup

OnGroup is a hook to execute user functions on each group registeration. Also you can get group properties by **group** parameter.

​	OnGroup 是一个钩子，用于在每个组注册上执行用户函数。您还可以通过组参数获取组属性。

Signature
签名

```go
func (h *Hooks) OnGroup(handler ...OnGroupHandler)
```



## OnGroupName

OnGroupName is a hook to execute user functions on each group naming. Also you can get group properties by **group** parameter.

​	OnGroupName 是一个钩子，用于在每个组命名上执行用户函数。您还可以通过组参数获取组属性。

CAUTION
注意

OnGroupName only works with naming groups, not routes.

​	OnGroupName 仅适用于命名组，不适用于路由。

Signature
签名

```go
func (h *Hooks) OnGroupName(handler ...OnGroupNameHandler)
```



## OnListen

OnListen is a hook to execute user functions on Listen, ListenTLS, Listener.

​	OnListen 是一个在 Listen、ListenTLS、Listener 上执行用户函数的钩子。

Signature
签名

```go
func (h *Hooks) OnListen(handler ...OnListenHandler)
```



- OnListen Example
  OnListen 示例

```go
app := fiber.New(fiber.Config{
  DisableStartupMessage: true,
})

app.Hooks().OnListen(func(listenData fiber.ListenData) error {
  if fiber.IsChild() {
      return nil
  }
  scheme := "http"
  if data.TLS {
    scheme = "https"
  }
  log.Println(scheme + "://" + listenData.Host + ":" + listenData.Port)
  return nil
})

app.Listen(":5000")
```



## OnFork

OnFork is a hook to execute user functions on Fork.

​	OnFork 是一个在 Fork 上执行用户函数的钩子。

Signature
签名

```go
func (h *Hooks) OnFork(handler ...OnForkHandler)
```



## OnShutdown

OnShutdown is a hook to execute user functions after Shutdown.

​	OnShutdown 是一个在 Shutdown 后执行用户函数的钩子。

Signature
签名

```go
func (h *Hooks) OnShutdown(handler ...OnShutdownHandler)
```



## OnMount

OnMount is a hook to execute user function after mounting process. The mount event is fired when sub-app is mounted on a parent app. The parent app is passed as a parameter. It works for app and group mounting.

​	OnMount 是一个在挂载进程后执行用户函数的钩子。当子应用挂载到父应用时，会触发挂载事件。父应用作为参数传递。它适用于应用和组挂载。

Signature
签名

```go
func (h *Hooks) OnMount(handler ...OnMountHandler) 
```



- OnMount Example
  OnMount 示例

```go
package main

import (
    "fmt"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := New()
    app.Get("/", testSimpleHandler).Name("x")

    subApp := New()
    subApp.Get("/test", testSimpleHandler)

    subApp.Hooks().OnMount(func(parent *fiber.App) error {
        fmt.Print("Mount path of parent app: "+parent.MountPath())
        // ...

        return nil
    })

    app.Mount("/sub", subApp)
}

// Result:
// Mount path of parent app: 
```



> CAUTION
> 注意
>
> OnName/OnRoute/OnGroup/OnGroupName hooks are mount-sensitive. If you use one of these routes on sub app and you mount it; paths of routes and groups will start with mount prefix.
>
> ​	OnName/OnRoute/OnGroup/OnGroupName 钩子对挂载敏感。如果您在子应用上使用其中一个路由并将其挂载；路由和组的路径将以挂载前缀开头。
