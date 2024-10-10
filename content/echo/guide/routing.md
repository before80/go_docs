+++
title = "路由"
weight = 110
date = 2023-07-09T21:52:01+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Routing - 路由

> 原文：[https://echo.labstack.com/docs/routing](https://echo.labstack.com/docs/routing)

​	Echo的路由器基于[基数树（radix tree）](http://en.wikipedia.org/wiki/Radix_tree)，使路由查找非常快速。它利用[同步池（sync pool）](https://golang.org/pkg/sync/#Pool)来重用内存，实现零动态内存分配且没有垃圾回收开销。

​	可以通过指定HTTP方法、路径和匹配的处理程序来注册路由。例如，下面的代码注册了一个处理方法为`GET`，路径为`/hello`，发送`Hello, World!` HTTP响应的路由。

```go
// Handler
func hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}

// Route
e.GET("/hello", hello)
```



​	您可以使用`Echo.Any(path string, h Handler)`为所有HTTP方法注册处理程序。如果您只想为某些方法注册它，请使用`Echo.Match(methods []string, path string, h Handler)`。

​	Echo将处理程序函数定义为`func(echo.Context) error`，其中`echo.Context`主要保存了HTTP请求和响应的接口。

## Match-any

​	在路径中匹配零个或多个字符。例如，模式`/users/*`将匹配以下路径： 

- `/users/`
- `/users/1`
- `/users/1/files/1`
- `/users/anything...`

## 路径匹配顺序

- Static
- Param
- Match any

*示例*

```go
e.GET("/users/:id", func(c echo.Context) error {
    return c.String(http.StatusOK, "/users/:id")
})

e.GET("/users/new", func(c echo.Context) error {
    return c.String(http.StatusOK, "/users/new")
})

e.GET("/users/1/files/*", func(c echo.Context) error {
    return c.String(http.StatusOK, "/users/1/files/*")
})
```



​	上述路由将按照以下顺序解析：

- `/users/new`
- `/users/:id`
- `/users/1/files/*`

> 提示
>
> ​	可以以任意顺序编写路由。

## 分组

```
Echo#Group(prefix string, m ...Middleware) *Group
```

​	具有共同前缀的路由可以分组，以定义具有可选中间件的新子路由器（new sub-router）。除了指定的中间件之外，该组还继承父级中间件。要在组中稍后添加中间件，可以使用`Group.Use(m ...Middleware)`。分组也可以嵌套。

​	在下面的代码中，我们创建了一个需要基本HTTP身份验证的管理员组，用于路由`/admin/*`。

```go
g := e.Group("/admin")
g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
    if username == "joe" && password == "secret" {
        return true, nil
    }
    return false, nil
}))
```



## 路由命名

​	每个注册方法返回一个`Route`对象，可以在注册后使用它来为路由命名。例如：

```go
route := e.POST("/users", func(c echo.Context) error {
})
route.Name = "create-user"

// or using the inline syntax
e.GET("/users/:id", func(c echo.Context) error {
}).Name = "get-user"
```



​	当从模板生成URI时，路由名称非常有用，特别是在无法访问处理程序引用或者存在多个使用相同处理程序的路由时。

## 构建URI

​	`Echo#URI(handler HandlerFunc, params ...interface{})` 可以用于为任何处理程序生成具有指定路径参数的URI。这有助于集中管理您的URI模式，从而方便您重构应用程序。

​	例如，对于下面注册的路由，`e.URI(h, 1)`将生成`/users/1`。

```go
// Handler
h := func(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
}

// Route
e.GET("/users/:id", h)
```



​	除了`Echo#URI`之外，还有`Echo#Reverse(name string, params ...interface{})`，它用于基于路由名称生成URI。例如，如果像下面这样注册了名为`foobar`的路由，则调用`Echo#Reverse("foobar", 1234)`将生成URI`/users/1234`。

```go
// Handler
h := func(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
}

// Route
e.GET("/users/:id", h).Name = "foobar"
```



## 列出路由

​	`Echo#Routes() []*Route` 可以用于按照定义的顺序列出所有注册的路由。每个路由包含HTTP方法、路径和关联的处理程序。

*示例*

```go
// Handlers
func createUser(c echo.Context) error {
}

func findUser(c echo.Context) error {
}

func updateUser(c echo.Context) error {
}

func deleteUser(c echo.Context) error {
}

// Routes
e.POST("/users", createUser)
e.GET("/users", findUser)
e.PUT("/users", updateUser)
e.DELETE("/users", deleteUser)
```



​	使用以下代码，您可以将所有路由输出到JSON文件中：

```go
data, err := json.MarshalIndent(e.Routes(), "", "  ")
if err != nil {
    return err
}
os.WriteFile("routes.json", data, 0644)
```

`routes.json`

```json
[
  {
    "method": "POST",
    "path": "/users",
    "name": "main.createUser"
  },
  {
    "method": "GET",
    "path": "/users",
    "name": "main.findUser"
  },
  {
    "method": "PUT",
    "path": "/users",
    "name": "main.updateUser"
  },
  {
    "method": "DELETE",
    "path": "/users",
    "name": "main.deleteUser"
  }
]
```