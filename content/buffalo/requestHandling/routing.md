+++
title = "路由"
date = 2024-02-04T21:07:11+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/request_handling/routing/](https://gobuffalo.io/documentation/request_handling/routing/)

# Routing 路由 

Buffalo uses the [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) package under the covers, to handle routing within Buffalo applications. With that said, Buffalo wraps the `mux` API with its own. This guide walks you through all you’ll need to know about how Buffalo handles routing.

​	Buffalo 在后台使用 github.com/gorilla/mux 包来处理 Buffalo 应用程序中的路由。话虽如此，Buffalo 用自己的 API 包装了 `mux` 。本指南将引导您了解 Buffalo 如何处理路由的所有信息。

## We need to have the buffalo.App configuration created 我们需要创建 buffalo.App 配置 

The app configuration is located in the `actions/app.go` file.

​	应用程序配置位于 `actions/app.go` 文件中。

```go
// actions/app.go
app = buffalo.New(buffalo.Options{
    Env:         ENV,
    SessionName: "_coke_session",
  })
```

By default, buffalo requires only 2 options for its app setup:

​	默认情况下，buffalo 只需要 2 个选项来设置应用程序：

- `Env`: The enviroment where the application will run. Default value: `development`.
  `Env` ：应用程序将运行的环境。默认值： `development` 。
- `SessionName`: Is the session cookie that is set. Default value: `_buffalo_session`.
  `SessionName` ：设置的会话 cookie。默认值： `_buffalo_session` 。

You are free to customize it to fit your use case.

​	您可以自由定制以适合您的用例。

You can check the available options list here: https://godoc.org/github.com/gobuffalo/buffalo#Options

​	您可以在此处查看可用选项列表：https://godoc.org/github.com/gobuffalo/buffalo#Options

## Buffalo.Handler

If you already know about **MVC pattern**, `buffalo.Handler` functions manages the Controller part. Its signature looks like this:

​	如果您已经了解 MVC 模式， `buffalo.Handler` 函数管理控制器部分。其签名如下所示：

```go
func (c buffalo.Context) error {
  // do some work
}
```

This is the place where all the app logic goes. The handler function takes a `buffalo.Context` param, which contains everything you need about the current request.

​	这是所有应用程序逻辑所在的位置。处理程序函数采用 `buffalo.Context` 参数，其中包含有关当前请求所需的一切。

See the [Context](https://gobuffalo.io/documentation/request_handling/context) to understand the `buffalo.Context` interface.
请参阅上下文以了解 `buffalo.Context` 接口。

## Mapping Handlers 映射处理程序 

To map a `buffalo.Handler`, you’ll need to associate it with an specific path with an HTTP method.

​	要映射 `buffalo.Handler` ，您需要使用 HTTP 方法将其与特定路径相关联。

##### Supported HTTP Methods 支持的 HTTP 方法 

Buffalo supports the following HTTP methods:

​	Buffalo 支持以下 HTTP 方法：

GET

POST

PUT

PATCH

DELETE

OPTIONS

HEAD

```go
app.GET("/your/path", buffalo.Handler)
```

You can also match all HTTP methods using `ANY`.

​	您还可以使用 `ANY` 匹配所有 HTTP 方法。

As default, Buffalo sets a root path inside bufalo.App setup:

​	默认情况下，Buffalo 在 bufalo.App 设置中设置根路径：

```go
// actions/app.go
func App() *buffalo.App {
  // ...
  app.GET("/", HomeHandler)
  // ...
}
```

Mapping multiple `buffalo.Handlers` to HTTP methods take the form of:

​	将多个 `buffalo.Handlers` 映射到 HTTP 方法的形式为：

```go
// actions/app.go
app.GET("/", HomeHandler)
app.GET("/some/path", SomeHandler)
app.POST("/another/path", func (c buffalo.Context) error {
  // do some work
})
// etc...
```

As you can see, you can use inline buffalo.Handlers if you want. For more readability though, it’s often better to separate your handlers into multiple files. For example, if you have many handlers managing users stuff, you can group them into a `users.go` file in the [`actions`](https://gobuffalo.io/documentation/getting_started/directory-structure#actions) folder, for instance.

​	如您所见，如果您愿意，可以使用内联 buffalo.Handlers。但为了提高可读性，通常最好将处理程序分成多个文件。例如，如果您有许多处理程序管理用户内容，可以将它们分组到 `users.go` 文件夹中的 `actions` 文件中。

## Named Routes 命名路由 

By default, Buffalo will name your routes for you in the form of `<pathName>Path`.

​	默认情况下，Buffalo 会以 `<pathName>Path` 的形式为您命名路由。

For example: `a.GET("/coke", CokeHandler)` will result in a route named `cokePath`.

​	例如： `a.GET("/coke", CokeHandler)` 将导致一个名为 `cokePath` 的路由。

```go
a.GET("/coke", CokeHandler) // cokePath()
```

These names become the name of the route helpers in your templates.

​	这些名称成为模板中路由帮助程序的名称。

```html
<a href="<%= cokePath() %>">Coke</a>
```

## Custom Named Routes 自定义命名路由 

Buffalo also provides you a way to set a custom name for your route, The [`buffalo.RouteInfo#Name`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo.Name) method allows you to set a custom name for route helpers. To customize your route name, just use the Name method after mapping the HTTP Method.

​	Buffalo 还为您提供了一种为路由设置自定义名称的方法， `buffalo.RouteInfo#Name` 方法允许您为路由帮助程序设置自定义名称。要自定义路由名称，只需在映射 HTTP 方法后使用 Name 方法。

```go
app.GET("/coke", CokeHandler).Name("customCoke") // customCokePath()
```

This route is now called `customCokePath` and you can reference it as such in your templates.

​	此路由现在称为 `customCokePath` ，您可以在模板中引用它。

```html
<a href="<%= customCokePath() %>">Coke</a>
```

## Route list 路由列表 

You can inspect all of your paths by running `buffalo routes` from the command line.

​	您可以通过从命令行运行 `buffalo routes` 来检查所有路径。

```bash
$ buffalo routes

METHOD | HOST                  | PATH                    | ALIASES | NAME                 | HANDLER
------ | ----                  | ----                    | ------- | ----                 | -------
GET    | http://127.0.0.1:3000 | /                       |         | rootPath             | coke/actions.HomeHandler
GET    | http://127.0.0.1:3000 | /some/path/             |         | somePath             | coke/actions.SomeHandler
POST   | http://127.0.0.1:3000 | /another/path/          |         | anotherPath          | coke/actions.App.func1
GET    | http://127.0.0.1:3000 | /coke/                  |         | customCokePath       | coke/actions.CokeHandler
```

**IMPORTANT:** Because route helper names are calculated using the **`path`** pe. **`/widgets/new -> newWidgetsPath`**; if path changes, then the route helper name **also** changes.
重要提示：因为路由辅助名称是使用 `path` pe. `/widgets/new -> newWidgetsPath` 计算的；如果路径发生更改，则路由辅助名称也会更改。

Example:

​	示例：

Mapping `WidgetResource` in `/widgets` path:

​	映射 `WidgetResource` 中的 `/widgets` 路径：

```go
app.Resource("/widgets", WidgetsResource{})
```

You will get the following route path names:

​	您将获得以下路由路径名称：

```bash
$ buffalo routes

METHOD | HOST                  | PATH                       | ALIASES | NAME                 | HANDLER
------ | ----                  | ----                       | ------- | ----                 | -------
GET    | http://127.0.0.1:3000 | /                          |         | rootPath             | coke/actions.HomeHandler
GET    | http://127.0.0.1:3000 | /some/path/                |         | somePath             | coke/actions.SomeHandler
POST   | http://127.0.0.1:3000 | /another/path/             |         | anotherPath          | coke/actions.App.func1
GET    | http://127.0.0.1:3000 | /coke/                     |         | customCokePath       | coke/actions.CokeHandler

GET    | http://127.0.0.1:3000 | /widgets/                  |         | widgetsPath    | coke/actions.WidgetResource.List
POST   | http://127.0.0.1:3000 | /widgets/                  |         | widgetsPath    | coke/actions.WidgetResource.Create
GET    | http://127.0.0.1:3000 | /widgets/new/              |         | newWidgetsPath | coke/actions.WidgetResource.New
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/      |         | widgetPath     | coke/actions.WidgetResource.Show
PUT    | http://127.0.0.1:3000 | /widgets/{widget_id}/      |         | widgetPath     | coke/actions.WidgetResource.Update
DELETE | http://127.0.0.1:3000 | /widgets/{widget_id}/      |         | widgetPath     | coke/actions.WidgetResource.Destroy
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/edit/ |         | editWidgetPath | coke/actions.WidgetResource.Edit
```

But, if you rename the route path to `/fooz`:

​	但是，如果您将路由路径重命名为 `/fooz` ：

```go
app.Resource("/fooz", WidgetsResource{})
```

The route names will be renamed to:

​	路由名称将重命名为：

```bash
$ buffalo routes

METHOD | HOST                  | PATH                    | ALIASES | NAME                 | HANDLER
------ | ----                  | ----                    | ------- | ----                 | -------
GET    | http://127.0.0.1:3000 | /                       |         | rootPath             | coke/actions.HomeHandler
GET    | http://127.0.0.1:3000 | /some/path/             |         | somePath             | coke/actions.SomeHandler
POST   | http://127.0.0.1:3000 | /another/path/          |         | anotherPath          | coke/actions.App.func1
GET    | http://127.0.0.1:3000 | /coke/                  |         | customCokePath       | coke/actions.CokeHandler

GET    | http://127.0.0.1:3000 | /fooz/                  |         | foozPath             | coke/actions.WidgetResource.List
POST   | http://127.0.0.1:3000 | /fooz/                  |         | foozPath             | coke/actions.WidgetResource.Create
GET    | http://127.0.0.1:3000 | /fooz/new/              |         | newFoozPath          | coke/actions.WidgetResource.New
GET    | http://127.0.0.1:3000 | /fooz/{widget_id}/      |         | foozWidgetIDPath     | coke/actions.WidgetResource.Show
PUT    | http://127.0.0.1:3000 | /fooz/{widget_id}/      |         | foozWidgetIDPath     | coke/actions.WidgetResource.Update
DELETE | http://127.0.0.1:3000 | /fooz/{widget_id}/      |         | foozWidgetIDPath     | coke/actions.WidgetResource.Destroy
GET    | http://127.0.0.1:3000 | /fooz/{widget_id}/edit/ |         | editFoozWidgetIDPath | coke/actions.WidgetResource.Edit
```

See [`Custom Named Routes`](https://gobuffalo.io/documentation/request_handling/routing/#custom-named-routes) for details on how to change the generated name.

​	请参阅 `Custom Named Routes` ，了解有关如何更改生成名称的详细信息。

## Using Route Helpers in Templates 在模板中使用路由辅助 

Route helpers can be used directly in templates using the name of the helper:

​	路由助手可以使用助手的名称直接在模板中使用：

```erb
<%= widgetsPath() %> // /widgets
```

Routes that require named parameters, must be fed a map of those parameters.

​	需要命名参数的路由必须提供这些参数的映射。

```erb
<%= editWidgetPath({widget_id: 1}) %> --> /widgets/1/edit
```

## The `pathFor` Helper `pathFor` 助手 

The `pathFor` helper takes an `interface{}`, or a `slice` of them, and tries to convert it to a `/foos/{id}` style URL path.

​	 `pathFor` 助手接受一个 `interface{}` 或多个 `interface{}`，并尝试将其转换为 `/foos/{id}` 样式的 URL 路径。

Rules:

​	规则：

- if `string` it is returned as is
  如果 `string` ，则按原样返回
- if `Pathable` the `ToPath` method is returned
  如果 `Pathable` ，则返回 `ToPath` 方法
- if `slice` or an `array` each element is run through the helper then joined
  如果 `slice` 或 `array` ，则每个元素都通过助手运行，然后连接
- if `struct` the name of the struct, pluralized is used for the name
  如果 `struct` ，则使用结构的名称（复数形式）作为名称
- if `Paramable` the `ToParam` method is used to fill the `{id}` slot
  如果 `Paramable` 方法用于填充 `{id}` 插槽
- if `struct.Slug` the slug is used to fill the `{id}` slot of the URL
  如果 `struct.Slug` slug 用于填充 URL 的 `{id}` 插槽
- if `struct.ID` the ID is used to fill the `{id}` slot of the URL
  如果 `struct.ID` ID 用于填充 URL 的 `{id}` 插槽

```go
// Car{1} => "/cars/1"
// Car{} => "/cars"
// &Car{} => "/cars"
type Car struct {
  ID int
}

// Boat{"titanic"} => "/boats/titanic"
type Boat struct {
  Slug string
}

// Plane{} => "/planes/aeroPlane"
type Plane struct{}

func (Plane) ToParam() string {
  return "aeroPlane"
}

// Truck{} => "/a/Truck"
// {[]interface{}{Truck{}, Plane{}} => "/a/Truck/planes/aeroPlane"
type Truck struct{}

func (Truck) ToPath() string {
  return "/a/Truck"
}
```

## Using Route Helpers in Actions 在操作中使用路由助手 

### Redirecting with Route Helpers 使用路由助手重定向 

You can also use route names when redirecting to another url.

​	您还可以在重定向到另一个 URL 时使用路由名称。

```go
func MyHandler(c buffalo.Context) error {
  return c.Redirect(http.StatusSeeOther, "widgetsPath()")
  // Or with parameters
  return c.Redirect(http.StatusSeeOther, "widgetPath()", render.Data{"widget_id": "1"})
}
```

------

### Finding/Calling a Route Helper 查找/调用路由助手 

Since **0.13.0-beta.1**
自 0.13.0-beta.1 起



The [`buffalo.RouteList#Lookup`](https://godoc.org/github.com/gobuffalo/buffalo#RouteList.Lookup) allows you to look up a route by its name from the application. With the `RouteInfo` value for the given route you can generate the path for the route.

​	 `buffalo.RouteList#Lookup` 允许您通过其名称从应用程序中查找路由。使用给定路由的 `RouteInfo` 值，您可以生成该路由的路径。

```go
func MyHandler(c buffalo.Context) error {
  ri, err := App().Routes().Lookup("widgetPath")
  if err != nil {
    return errors.WithStack(err)
  }
  h := ri.BuildPathHelper()
  u, err := h(render.Data{"widget_id": 1})
  if err != nil {
    return errors.WithStack(err)
  }
  return c.Redirect(307, string(u))
}
```

## Parameters 参数 

Query string and other parameters are available from the [`buffalo.Context`](https://gobuffalo.io/documentation/request_handling/context) that is passed into the `buffalo.Handler`.

​	查询字符串和其他参数可从传递给 `buffalo.Handler` 的 `buffalo.Context` 中获得。

```go
a.GET("/users", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Given the above code sample, if we make a request with `GET /users?name=ringo`, the response should be `200: ringo`.

​	根据上面的代码示例，如果我们使用 `GET /users?name=ringo` 发出请求，则响应应为 `200: ringo` 。

## Named Parameters 命名参数 

Since Buffalo is the [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) under the covers, it means we can get access to some of the goodness it provides. In this case, the ability to create pseudo-regular expression patterns in the mapped path that will get converted into parameters that can be accessed from a [`buffalo.Context`](https://gobuffalo.io/documentation/request_handling/context).

​	由于 Buffalo 是 github.com/gorilla/mux 的底层，这意味着我们可以访问它提供的一些优点。在这种情况下，能够在映射路径中创建伪正则表达式模式，该模式将转换为可从 `buffalo.Context` 访问的参数。

```go
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Given the above code sample, if we make a request with `GET /users/ringo`, the response should be `200: ringo`.

​	根据上面的代码示例，如果我们使用 `GET /users/ringo` 发出请求，则响应应为 `200: ringo` 。

```go
a.GET("/users/new", func (c buffalo.Context) error {
  return c.Render(200, r.String("new"))
})
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

You may map seemingly similar paths, like `/users/new` and `/users/{name}` without any issues. The router will make sure they get to the same place.

​	您可以映射看似相似的路径，例如 `/users/new` 和 `/users/{name}` ，而不会出现任何问题。路由器将确保它们到达相同的位置。

### Regular expressions 正则表达式 

[github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) provides a way to use regular expressions, so you can pre-filter queries:

​	github.com/gorilla/mux 提供了一种使用正则表达式的方法，因此您可以预先过滤查询：

```go
a.GET("/articles/{id:[0-9]+}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("id")))
})
```

## Groups 组 

Buffalo apps allow for the grouping of end-points. This allows for common functionality, such as [middleware](https://gobuffalo.io/documentation/request_handling/middleware) to be collected together. A great example of this would be an API end-point.

​	Buffalo 应用程序允许对端点进行分组。这允许将常见的功能（例如中间件）收集在一起。一个很好的例子就是 API 端点。

```go
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
g.GET("/users", func (c buffalo.Context) error {
  // responds to GET /api/v1/users
})
```

By default a group will inherit any middleware from its parent app.

​	默认情况下，组将继承其父应用程序的任何中间件。

```go
a.Use(SomeMiddleware)
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
```

In the above example the `/api/v1` group will use both `SomeMiddleware` and `APIAuthorizer`. See [middleware](https://gobuffalo.io/documentation/request_handling/middleware) for more information about using, skipping, and clearing middleware.

​	在上面的示例中， `/api/v1` 组将同时使用 `SomeMiddleware` 和 `APIAuthorizer` 。有关使用、跳过和清除中间件的更多信息，请参阅中间件。

## Virtual Hosts 虚拟主机 

Since **0.18.2**
自 0.18.2 起



Buffalo apps also support grouping of end-points by host. `VirtualHost` creates a new group that matches the domain passed. This is useful for creating groups of end-points for different domains or subdomains.

​	Buffalo 应用程序还支持按主机对端点进行分组。 `VirtualHost` 创建一个与传递的域匹配的新组。这对于为不同域或子域创建端点组非常有用。

```go
app := buffalo.New(buffalo.Options{
    Env:         envy.Get("GO_ENV", "development"),
    SessionName: "_coke_session",
})

subApp := app.VirtualHost("docs.domain.com")
subApp.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("docs.domain.com Homepage"))
})

domainApp := app.VirtualHost("example.com")
domainApp.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("example.com Homepage"))
})

app.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("Main App Homepage"))
})
```

Variables mapped to parameters are also supported:

​	还支持映射到参数的变量：

```go
app.VirtualHost("{subdomain}.example.com")
app.VirtualHost("{subdomain:[a-z]+}.example.com")
```

## Mounting http.Handler Apps 挂载 http.Handler 应用程序 

Since **0.9.4**
自 0.9.4 起



Sometimes, you’ll want to reuse some components from other apps. Using the [`Mount`](https://godoc.org/github.com/gobuffalo/buffalo#App.Mount) method, you can bind a standard [`http.Handler`](https://golang.org/pkg/net/http/#Handler) to a route, just like you’ll do with a normal route handler.

​	有时，您会希望重用其他应用程序中的一些组件。使用 `Mount` 方法，您可以将标准 `http.Handler` 绑定到路由，就像使用普通路由处理程序一样。

```go
func muxer() http.Handler {
  f := func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "%s - %s", req.Method, req.URL.String())
  }
  mux := mux.NewRouter()
  mux.HandleFunc("/foo/", f).Methods("GET")
  mux.HandleFunc("/bar/", f).Methods("POST")
  mux.HandleFunc("/baz/baz/", f).Methods("DELETE")
  return mux
}

a.Mount("/admin", muxer())
```

Since Buffalo `App` implements the `http.Handler` interface, you can also mount another Buffalo app and build modular apps.

​	由于 Buffalo `App` 实现 `http.Handler` 接口，您还可以挂载另一个 Buffalo 应用程序并构建模块化应用程序。