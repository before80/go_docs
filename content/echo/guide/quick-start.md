+++
title = "快速入门"
weight = 10
date = 2023-07-09T21:49:33+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Quick Start - 快速入门

> 原文：[https://echo.labstack.com/docs/quick-start](https://echo.labstack.com/docs/quick-start)

## 安装

### 要求

​	安装 Echo 需要 Go 1.13 或更高版本。Go 1.12 的有限支持，某些中间件可能不可用。请确保您的项目文件夹位于 `$GOPATH` 之外。

```sh
$ mkdir myapp && cd myapp
$ go mod init myapp
$ go get github.com/labstack/echo/v4
```

​	如果您正在使用 Go v1.14 或更早的版本，请使用以下命令：

```sh
$ GO111MODULE=on go get github.com/labstack/echo/v4
```



## Hello, World!

​	创建 `server.go` 文件

```go
package main

import (
    "net/http"
    
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}
```

​	启动服务器

```sh
$ go run server.go
```

​	在浏览器中打开 [http://localhost:1323](http://localhost:1323/)，您将在页面上看到 Hello, World!。

## 路由

```go
e.POST("/users", saveUser)
e.GET("/users/:id", getUser)
e.PUT("/users/:id", updateUser)
e.DELETE("/users/:id", deleteUser)
```



## 路径参数

```go
// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
    // 从路径 `users/:id` 获取用户 ID
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}
```

​	在浏览器中打开 [http://localhost:1323/users/joe](http://localhost:1323/users/joe)，您将在页面上看到 'joe'。

## 查询参数

`/show?team=x-men&member=wolverine`

```go
//e.GET("/show", show)
func show(c echo.Context) error {
    // 从查询字符串中获取team和member信息
    team := c.QueryParam("team")
    member := c.QueryParam("member")
    return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
```

​	在浏览器中打开 [http://localhost:1323/show?team=x-men&member=wolverine](http://localhost:1323/show?team=x-men&member=wolverine)，您将在页面上看到 'team:x-men, member:wolverine'。

## 表单application/x-www-form-urlencoded

`POST` `/save`

| name  | value                                       |
| ----- | ------------------------------------------- |
| name  | Joe Smith                                   |
| email | [joe@labstack.com](mailto:joe@labstack.com) |

```go
// e.POST("/save", save)
func save(c echo.Context) error {
    // 获取 name 和 email
    name := c.FormValue("name")
    email := c.FormValue("email")
    return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}
```

​	运行以下命令：

```sh
$ curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:1323/save
// => name:Joe Smith, email:joe@labstack.com
```



## 表单multipart/form-data

`POST` `/save`

| name   | value     |
| ------ | --------- |
| name   | Joe Smith |
| avatar | avatar    |

```go
func save(c echo.Context) error {
    // 获取 name
    name := c.FormValue("name")
    // 获取 avatar
    avatar, err := c.FormFile("avatar")
    if err != nil {
        return err
    }
 
    // 源文件
    src, err := avatar.Open()
    if err != nil {
        return err
    }
    defer src.Close()
 
    // 目标文件
    dst, err := os.Create(avatar.Filename)
    if err != nil {
        return err
    }
    defer dst.Close()
 
    // 复制文件
    if _, err = io.Copy(dst, src); err != nil {
        return err
    }

    return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}
```

​	运行以下命令：

```sh
$ curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/save
// => <b>Thank you! Joe Smith</b>
```

​	为了查看已上传的图片，请运行以下命令：

```sh
cd <project directory>
ls avatar.png
// => avatar.png
```



## 处理请求

- 将 `json`、`xml`、`form` 或 `query` 载荷绑定到基于 `Content-Type` 请求头的 Go 结构体中。
- 以带有状态码的 `json` 或 `xml` 形式渲染响应。

```go
type User struct {
    Name  string `json:"name" xml:"name" form:"name" query:"name"`
    Email string `json:"email" xml:"email" form:"email" query:"email"`
}

e.POST("/users", func(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    return c.JSON(http.StatusCreated, u)
    // 或者
    // return c.XML(http.StatusCreated, u)
})
```



## 静态内容

​	为路径 `/static/*` 从静态目录中提供任何文件：

```go
e.Static("/static", "static")
```

[了解更多]({{< ref "/echo/guide/static-files">}})

## 模板渲染

## 中间件

```go
// 根级别（Root level）中间件
e.Use(middleware.Logger())
e.Use(middleware.Recover())

// 分组级别（Group level）中间件
g := e.Group("/admin")
g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
  if username == "joe" && password == "secret" {
    return true, nil
  }
  return false, nil
}))

// 路由级别（Route level）中间件
track := func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        println("request to /users")
        return next(c)
    }
}
e.GET("/users", func(c echo.Context) error {
    return c.String(http.StatusOK, "/users")
}, track)
```



