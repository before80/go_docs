+++
title = "context"
date = 2024-02-04T21:07:58+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/request_handling/context/]({{< ref "/buffalo/requestHandling/context" >}})

# Context 上下文 

At the heart of every Buffalo request handler is the `Context`. This context gives handlers a simple, and clean, function definition, while being immensely powerful.

​	每个 Buffalo 请求处理程序的核心是 `Context` 。此上下文为处理程序提供了一个简单且清晰的功能定义，同时功能非常强大。

## The Context Interface 上下文接口 

The `buffalo.Context` interface supports `context.Context` so it can be passed around and used as a “standard” Go Context.

​	 `buffalo.Context` 接口支持 `context.Context` ，因此可以传递它并将其用作“标准”Go Context。

Since `buffalo.Context` is an interface it is possible to create an application specific implementation that is tailored to the needs of the application being built.

​	由于 `buffalo.Context` 是一个接口，因此可以创建针对正在构建的应用程序的需求量身定制的特定于应用程序的实现。

Since **0.12.0**
自 0.12.0 起



```go
type Context interface {
	context.Context
	Response() http.ResponseWriter
	Request() *http.Request
	Session() *Session
	Cookies() *Cookies
	Params() ParamValues
	Param(string) string
	Set(string, interface{})
	LogField(string, interface{})
	LogFields(map[string]interface{})
	Logger() Logger
	Bind(interface{}) error
	Render(int, render.Renderer) error
	Error(int, error) error
	Redirect(int, string, ...interface{}) error
	Data() map[string]interface{}
	Flash() *Flash
	File(string) (binding.File, error)
}
```

The `Websocket() (*websocket.Conn, error)` function was removed from `buffalo.Context` in version `v0.12.0`. Use the http://www.gorillatoolkit.org/pkg/websocket package directly instead

​	 `Websocket() (*websocket.Conn, error)` 函数已从 `buffalo.Context` 中的版本 `v0.12.0` 中移除。直接使用 http://www.gorillatoolkit.org/pkg/websocket 包

## Context and Rendering 上下文和渲染 

As part of the context interface, there is a `Render` function that takes a type of `render.Renderer`. See [rendering]({{< ref "/buffalo/frontend/rendering" >}}) for more information.

​	作为上下文接口的一部分，有一个 `Render` 函数，它采用 `render.Renderer` 类型。有关更多信息，请参阅渲染。

Any values that are “set” on the context will automatically be available to the `render.Renderer` that is passed into the `Render` function.

​	在上下文中“设置”的任何值都将自动提供给传递给 `Render` 函数的 `render.Renderer` 。

```go
func Hello(c buffalo.Context) error {
  c.Set("name", "Mark")

  return c.Render(http.StatusOK, render.String("Hi <%= name %>"))
}
```

## Implementing the Interface 实现接口 

The `buffalo.Context` is never meant to be “fully” implemented. Instead it is recommended that you use [composition](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html) and implement only the functions that you want to provide custom implementations of.

​	 `buffalo.Context` 永远不会被“完全”实现。相反，建议您使用组合并仅实现您想要提供自定义实现的函数。

Below is an example of changing the `Error` function to log the error and kill application:

​	下面是一个将 `Error` 函数更改为记录错误并终止应用程序的示例：

```go
// actions/context.go
type MyContext struct {
  buffalo.Context
}

func (my MyContext) Error(status int, err error) error {
  my.Logger().Fatal(err)
  return err
}
// actions/app.go
// ...
func App() *buffalo.App {
  if app != nil {
    // ...
    app.Use(func (next buffalo.Handler) buffalo.Handler {
      return func(c buffalo.Context) error {
      // change the context to MyContext
      return next(MyContext{c})
      }
    })
    // ...
  }
  return app
}
// ...
```

## Ranging Over Parameters 遍历参数 

The `buffalo.Context#Params` method returns [`buffalo.ParamValues`](https://godoc.org/github.com/gobuffalo/buffalo#ParamValues) which is an interface around [`url.Values`](https://golang.org/pkg/net/url/#Values). You can cast to this type in a handler to range over the parameter values.

​	 `buffalo.Context#Params` 方法返回 `buffalo.ParamValues` ，它是 `url.Values` 周围的接口。您可以在处理程序中将此类型转换为范围超过参数值的类型。

```go
import "net/url"

func HomeHandler(c buffalo.Context) error {
  if m, ok := c.Params().(url.Values); ok {
    for k, v := range m {
      fmt.Println(k, v)
    }
  }

  return c.Render(http.StatusOK, r.HTML("index.html"))
}
```

## What’s in the Context 上下文中有什么 

Buffalo stuffs the context of each request with a lot of information that could be useful in your application, such as the `current_route` or the `session`. Below is a list of what Buffalo adds to the context on each request that you can access from in your actions or templates.

​	Buffalo 会在每个请求的上下文中填充大量信息，这些信息可能对您的应用程序很有用，例如 `current_route` 或 `session` 。以下是 Buffalo 在每个请求中添加到上下文中您可以从操作或模板中访问的内容列表。

| Key             | Type 类型                                                    | Usage 用法                                                   |
| --------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `app`           | [`*buffalo.App`](https://godoc.org/github.com/gobuffalo/buffalo#App) | The current Buffalo application that’s running. 正在运行的当前 Buffalo 应用程序。 |
| `env`           | `string`                                                     | The current environment the app is running in. Example: `test`, `development`, `production` 应用程序正在运行的当前环境。示例： `test` 、 `development` 、 `production` |
| `routes`        | [`buffalo.RouteList`](https://godoc.org/github.com/gobuffalo/buffalo#RouteList) | A list of all of the routes mapped on the application. 应用程序上映射的所有路由的列表。 |
| `current_route` | [`buffalo.RouteInfo`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo) | The current route that is being accessed. 正在访问的当前路由。 |
| `current_path`  | `string`                                                     | The current path being requested. Example: `/users/1/edit` 正在请求的当前路径。示例： `/users/1/edit` |
| `*Path`         | [`RouteHelperFunc`](https://godoc.org/github.com/gobuffalo/buffalo#RouteHelperFunc) | Helpers to create paths based off of mapped routes. Example: `editUserPath`. Run `buffalo task routes` to see a full list for your app. 基于映射的路由创建路径的帮助器。示例： `editUserPath` 。运行 `buffalo task routes` 以查看应用程序的完整列表。 |
| `params`        | `map[string]string`                                          | Query parameters for the requested page. 请求的页面的查询参数。 |
| `flash`         | `map[string][]string`                                        | A map of messages set using `buffalo.Context#Flash`. 使用 `buffalo.Context#Flash` 设置的消息映射。 |
| `session`       | [`*buffalo.Session`](https://godoc.org/github.com/gobuffalo/buffalo#Session) | The current user’s session. 当前用户会话。                   |
| `request`       | [`*http.Request`](https://godoc.org/net/http#Request)        | The current request. 当前请求。                              |
| `tx`            | [`*pop.Connection`](https://godoc.org/github.com/gobuffalo/pop#Connection) | Only set if using the `github.com/gobuffalo/buffalo/middleware.PopTransaction` middleware (on by default). 仅在使用 `github.com/gobuffalo/buffalo/middleware.PopTransaction` 中间件（默认情况下打开）时设置。 |

See [Helpers]({{< ref "/buffalo/frontend/helpers#builtin-helpers" >}}) for a list of built-in helper functions available inside of templates.

​	请参阅帮助程序，以获取模板内部可用的内置帮助程序函数列表。
