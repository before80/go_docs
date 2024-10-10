+++
title = "中间件"
date = 2024-02-04T21:08:25+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/request_handling/middleware/]({{< ref "/buffalo/requestHandling/middleware" >}})

# Middleware 中间件 

Middleware allows for the interjection of code in the request/response cycle. Common use cases for middleware are things like logging (which Buffalo already does), authentication requests, etc.

​	中间件允许在请求/响应周期中插入代码。中间件的常见用例包括日志记录（Buffalo 已经执行此操作）、身份验证请求等。

A list of “known” middleware packages can be found at https://toolkit.gobuffalo.io/tools?topic=middleware.

​	可以在 https://toolkit.gobuffalo.io/tools?topic=middleware 上找到“已知”中间件包的列表。

## Writing Your Own Middleware 编写您自己的中间件 

The [`buffalo.MiddlewareFunc`](https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareFunc) interface is any function that takes a `buffalo.Handler` and returns a `buffalo.Handler`.

​	 `buffalo.MiddlewareFunc` 接口是接受 `buffalo.Handler` 并返回 `buffalo.Handler` 的任何函数。

```go
func MyMiddleware(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
    // do some work before calling the next handler
    err := next(c)
    // do some work after calling the next handler
    return err
  }
}
```

By implementing the `buffalo.MiddlewareFunc` interface you are able to control the flow of execution in your application. Think an authorization middleware; send errors off to your favorite monitoring tool; load data on to the `buffalo.Context`, and more.

​	通过实现 `buffalo.MiddlewareFunc` 接口，您能够控制应用程序中的执行流。考虑授权中间件；将错误发送到您最喜欢的监控工具；将数据加载到 `buffalo.Context` 等。

### Example 示例 

```go
// UserIPMiddleware gets the user IP and sets it to the context.
func UserIPMiddleware(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
    if xRealIP := c.Request().Header.Get("X-Real-Ip"); len(xRealIP) > 0 {
      c.Set("user_ip", xRealIP)
      return next(c)
    }

    if xForwardedFor := c.Request().Header.Get("X-Forwarded-For"); len(xForwardedFor) > 0 {
      c.Set("user_ip", xForwardedFor)
      return next(c)
    }

    h, _, err := net.SplitHostPort(c.Request().RemoteAddr)
    if err != nil {
      return err
    }
    c.Set("user_ip", h)
    return next(c)
  }
}
```

## Using Middleware 使用中间件 

```go
a := buffalo.New(buffalo.Options{})

a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)
// or
a.Use(MyMiddleware, AnotherPieceOfMiddleware)
```

In the above example all requests will first go through the `MyMiddleware` middleware, and then through the `AnotherPieceOfMiddleware` middleware before first getting to their final handler.

​	在上述示例中，所有请求将首先通过 `MyMiddleware` 中间件，然后通过 `AnotherPieceOfMiddleware` 中间件，然后再到达其最终处理程序。

**NOTE**: Middleware defined on an application is automatically inherited by all routes and groups in that application.
注意：在应用程序上定义的中间件会自动继承该应用程序中的所有路由和组。

## Using Middleware with One Action 将中间件与一个操作一起使用 

Often there are cases when you want to use a piece of middleware on just one action, and not on the whole application or resource.

​	通常情况下，您可能只想对一个操作使用中间件，而不是整个应用程序或资源。

Since the definition of a piece of middleware is that it takes in a `buffalo.Handler` and returns a `buffalo.Handler` you can wrap any `buffalo.Handler` in a piece of middlware.

​	由于中间件的定义是它接收一个 `buffalo.Handler` 并返回一个 `buffalo.Handler` ，因此您可以用中间件包装任何 `buffalo.Handler` 。

```go
a := buffalo.New(buffalo.Options{})
a.GET("/foo", MyMiddleware(MyHandler))
```

This does not affect the rest of the middleware stack that is already in place, instead it appends to the middleware chain for just that one action.

​	这不会影响已经就位的其余中间件堆栈，而是仅针对该一个操作将中间件附加到中间件链中。

This can be taken a step further, by wrapping unlimited numbers of middleware around a `buffalo.Handler`.

​	可以通过在 `buffalo.Handler` 周围包装无限数量的中间件来进一步执行此操作。

```go
a := buffalo.New(buffalo.Options{})
a.GET("/foo", MyMiddleware(AnotherPieceOfMiddleware(MyHandler)))
```

## Group Middleware 组中间件 

```go
a := buffalo.New(buffalo.Options{})
a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)

g := a.Group("/api")
// authorize the API end-point
g.Use(AuthorizeAPIMiddleware)
g.GET("/users", UsersHandler)

a.GET("/foo", FooHandler)
```

In the above example the `MyMiddleware` and `AnotherPieceOfMiddleware` middlewares will be called on *all* requests, but the `AuthorizeAPIMiddleware` middleware will only be called on the `/api/*` routes.

​	在上面的示例中， `MyMiddleware` 和 `AnotherPieceOfMiddleware` 中间件将在所有请求上调用，但 `AuthorizeAPIMiddleware` 中间件将仅在 `/api/*` 路由上调用。

```text
GET /foo       -> MyMiddleware -> AnotherPieceOfMiddleware -> FooHandler
GET /api/users -> MyMiddleware -> AnotherPieceOfMiddleware -> AuthorizeAPIMiddleware -> UsersHandler
```

## Skipping Middleware 跳过中间件 

There are times when, in an application, you want to add middleware to the entire application, or a group, but not call that middleware on a few individual handlers. Buffalo allows you to create these sorts of mappings.

​	有时，您可能希望在应用程序中将中间件添加到整个应用程序或组，但不要在几个单独的处理程序上调用该中间件。Buffalo 允许您创建此类映射。

actions/app.go

OUTPUT
输出

```go
// actions/app.go
a := buffalo.New(buffalo.Options{})
a.Use(AuthorizeUser)

// skip the AuthorizeUser middleware for the NewUser and CreateUser handlers.
a.Middleware.Skip(AuthorizeUser, NewUser, CreateUser)

a.GET("/users/new", NewUser)
a.POST("/users", CreateUser)
a.GET("/users", ListUsers)
a.GET("/users/{id}", ShowUser)
```

------

**IMPORTANT:** The middleware function and the action functions you want to skip **MUST** be the same Go instance.
重要提示：要跳过的中间件函数和操作函数必须是同一个 Go 实例。

### Examples 示例 

EXAMPLE 1
示例 1

EXAMPLE 2
示例 2

```go
// EXAMPLE 1
m1 := MyMiddleware()
m2 := MyMiddleware()

app.Use(m1)

app.Skip(m2, Foo, Bar) // WON'T WORK m2 != m1
app.Skip(m1, Foo, Bar) // WORKS
```

See https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareStack.Skip for more details on the `Skip` function.

​	有关 `Skip` 函数的更多详细信息，请参阅 https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareStack.Skip。

## Skipping Resource Actions 跳过资源操作 

Often it is necessary to want to skip middleware for one or more actions. For example, allowing guest users to view the `List` and `Show` actions on a resource, but requiring authorization on the rest of the actions.

​	通常，需要跳过一个或多个操作的中间件。例如，允许访客用户查看资源上的 `List` 和 `Show` 操作，但要求对其余操作进行授权。

Understanding from the [Skipping Middleware](https://gobuffalo.io/documentation/request_handling/middleware/#skipping-middleware) section we need to make sure that we are using the same functions when we register the resource as we do when we want to skip the middleware on those functions later.

​	从跳过中间件部分了解到，我们需要确保在注册资源时使用与稍后在这些函数上跳过中间件时相同的函数。

The line that was generated in `actions/app.go` by `buffalo generate resource` will need to be changed to accommodate this requirement.

​	 `actions/app.go` 中由 `buffalo generate resource` 生成的行需要更改以满足此要求。

Before
之前

After
之后

```go
app.Resource("/widgets", WidgetResource{})
```

## Replace Middleware 替换中间件 

You can use the [`Middleware.Replace`](https://pkg.go.dev/github.com/gobuffalo/buffalo#MiddlewareStack.Replace) method that allows you to replace a middleware with another one keeping the same execution position.

​	您可以使用 `Middleware.Replace` 方法，该方法允许您用另一个中间件替换一个中间件，同时保持相同的执行位置。

actions/app.go

OUTPUT
输出

```go
// actions/app.go

app := buffalo.New(buffalo.Options{})
app.Use(Middleware1, Middleware2, Middleware3)

app.GET("/foo/", FooHandler)


g := app.Group("/group")
g.Middleware.Replace(Middleware1, Middleware4)

g.GET("/", GroupListHandler)
```

## Clearing Middleware 清除中间件 

Since middleware is [inherited](https://gobuffalo.io/documentation/request_handling/middleware/#using-middleware) from its parent, there maybe times when it is necessary to start with a “blank” set of middleware.

​	由于中间件是从其父级继承的，因此有时可能需要从一组“空白”中间件开始。

actions/app.go

OUTPUT
输出

```go
// actions/app.go
app := buffalo.New(buffalo.Options{})
app.Use(MyMiddleware)
app.Use(AnotherPieceOfMiddleware)

app.GET("/foo", FooHandler)

g := app.Group("/api")
// clear out any previously defined middleware
g.Middleware.Clear()
g.Use(AuthorizeAPIMiddleware)
g.GET("/users", UsersHandler)
```

## Listing an Application’s Middleware 列出应用程序的中间件 

To get a complete list of the middleware your application is using, broken down by grouping, can be found by running the `buffalo task middleware` command.

​	要获取应用程序正在使用的中间件的完整列表（按分组细分），可以运行 `buffalo task middleware` 命令。

actions/app.go

Middleware list
中间件列表

```go
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_coke_session",
		})

		app.Use(forceSSL())
		app.Use(paramlogger.ParameterLogger)
		app.Use(csrf.New)
		app.Use(translations())
		app.Use(Middleware1)
		app.Use(Middleware2)

		app.GET("/", HomeHandler)

		app.ServeFiles("/", http.FS(public.FS()))
	}

	return app
}
```
