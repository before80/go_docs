+++
title = "quick-start"
date = 2023-07-09T21:49:33+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Quick Start

https://echo.labstack.com/docs/quick-start

## Installation

### Requirements

To install Echo [Go](https://go.dev/doc/install) 1.13 or higher is required. Go 1.12 has limited support and some middlewares will not be available. Make sure your project folder is outside your $GOPATH.

```sh
$ mkdir myapp && cd myapp
$ go mod init myapp
$ go get github.com/labstack/echo/v4
```



If you are working with Go v1.14 or earlier use:

```sh
$ GO111MODULE=on go get github.com/labstack/echo/v4
```



## Hello, World!

Create `server.go`

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



Start server

```sh
$ go run server.go
```



Browse to [http://localhost:1323](http://localhost:1323/) and you should see Hello, World! on the page.

## Routing

```go
e.POST("/users", saveUser)
e.GET("/users/:id", getUser)
e.PUT("/users/:id", updateUser)
e.DELETE("/users/:id", deleteUser)
```



## Path Parameters

```go
// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
    // User ID from path `users/:id`
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}
```



Browse to http://localhost:1323/users/joe and you should see 'joe' on the page.

## Query Parameters

```
/show?team=x-men&member=wolverine
//e.GET("/show", show)
func show(c echo.Context) error {
    // Get team and member from the query string
    team := c.QueryParam("team")
    member := c.QueryParam("member")
    return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
```



Browse to http://localhost:1323/show?team=x-men&member=wolverine and you should see 'team:x-men, member:wolverine' on the page.

## Form application/x-www-form-urlencoded

```
POST` `/save
```

| name  | value                                       |
| ----- | ------------------------------------------- |
| name  | Joe Smith                                   |
| email | [joe@labstack.com](mailto:joe@labstack.com) |

```go
// e.POST("/save", save)
func save(c echo.Context) error {
    // Get name and email
    name := c.FormValue("name")
    email := c.FormValue("email")
    return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}
```



Run the following command:

```sh
$ curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:1323/save
// => name:Joe Smith, email:joe@labstack.com
```



## Form multipart/form-data

```
POST` `/save
```

| name   | value     |
| ------ | --------- |
| name   | Joe Smith |
| avatar | avatar    |

```go
func save(c echo.Context) error {
    // Get name
    name := c.FormValue("name")
    // Get avatar
    avatar, err := c.FormFile("avatar")
    if err != nil {
        return err
    }
 
    // Source
    src, err := avatar.Open()
    if err != nil {
        return err
    }
    defer src.Close()
 
    // Destination
    dst, err := os.Create(avatar.Filename)
    if err != nil {
        return err
    }
    defer dst.Close()
 
    // Copy
    if _, err = io.Copy(dst, src); err != nil {
        return err
    }

    return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}
```



Run the following command.

```sh
$ curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/save
// => <b>Thank you! Joe Smith</b>
```



For checking uploaded image, run the following command.

```sh
cd <project directory>
ls avatar.png
// => avatar.png
```



## Handling Request

- Bind `json`, `xml`, `form` or `query` payload into Go struct based on `Content-Type` request header.
- Render response as `json` or `xml` with status code.

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
    // or
    // return c.XML(http.StatusCreated, u)
})
```



## Static Content

Serve any file from static directory for path `/static/*`.

```go
e.Static("/static", "static")
```



[Learn More](https://echo.labstack.com/docs/static-files)

## [Template Rendering](https://echo.labstack.com/docs/templates)

## Middleware

```go
// Root level middleware
e.Use(middleware.Logger())
e.Use(middleware.Recover())

// Group level middleware
g := e.Group("/admin")
g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
  if username == "joe" && password == "secret" {
    return true, nil
  }
  return false, nil
}))

// Route level middleware
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