+++
title = "请求"
weight = 90
date = 2023-07-09T21:51:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Request - 请求

https://echo.labstack.com/docs/request

## 检索数据 - Retrieve Data

### 表单数据

​	可以通过名称使用 `Context#FormValue(name string)` 来检索表单数据：

```go
// 处理函数
func(c echo.Context) error {
  name := c.FormValue("name")
  return c.String(http.StatusOK, name)
}
```



```sh
curl -X POST http://localhost:1323 -d 'name=Joe'
```



​	要绑定自定义数据类型（的话），可以（通过）实现 `Echo#BindUnmarshaler` 接口（来实现）。

```go
type Timestamp time.Time

func (t *Timestamp) UnmarshalParam(src string) error {
  ts, err := time.Parse(time.RFC3339, src)
  *t = Timestamp(ts)
  return err
}
```



### 查询参数

​	可以通过名称使用 `Context#QueryParam(name string)` 来检索查询参数：

```go
// 处理函数
func(c echo.Context) error {
  name := c.QueryParam("name")
  return c.String(http.StatusOK, name)
})
```



```sh
curl \
  -X GET \
  http://localhost:1323\?name\=Joe
```



​	与表单数据类似，可以使用 `Context#QueryParam(name string)` 绑定自定义数据类型。

### 路径参数

​	注册的路径参数可以通过名称使用 `Context#Param(name string) string` 来检索：

```go
e.GET("/users/:name", func(c echo.Context) error {
  name := c.Param("name")
  return c.String(http.StatusOK, name)
})
```



```sh
curl http://localhost:1323/users/Joe
```



### 绑定数据

​		还支持将请求数据绑定到原生的 Go 结构体和变量中。请参阅 [绑定数据]({{< ref "/echo/guide/binding">}})。

## 验证数据

​	Echo 没有内置的数据验证功能，但是您可以使用 `Echo#Validator` 注册自定义验证器，并利用第三方的 [库](https://github.com/avelino/awesome-go#validation)。

​	下面的示例使用了 [https://github.com/go-playground/validator](https://github.com/go-playground/validator) 框架进行验证：

```go
package main

import (
  "net/http"

  "github.com/go-playground/validator"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

type (
  User struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
  }

  CustomValidator struct {
    validator *validator.Validate
  }
)

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
    // 可选地，您可以返回错误以使每个路由对状态码有更多的控制
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return nil
}

func main() {
  e := echo.New()
  e.Validator = &CustomValidator{validator: validator.New()}
  e.POST("/users", func(c echo.Context) (err error) {
    u := new(User)
    if err = c.Bind(u); err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    if err = c.Validate(u); err != nil {
      return err
    }
    return c.JSON(http.StatusOK, u)
  })
  e.Logger.Fatal(e.Start(":1323"))
}
```



```sh
curl -X POST http://localhost:1323/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe","email":"joe@invalid-domain"}'
{"message":"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag"}
```