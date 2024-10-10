+++
title = "绑定"
weight = 30
date = 2023-07-09T21:50:01+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Binding - 绑定

> 原文：[https://echo.labstack.com/docs/binding](https://echo.labstack.com/docs/binding)

​	解析请求数据是 Web 应用程序的关键部分。在 Echo 中，这是通过一种称为 *绑定* 的过程来完成的。绑定使用客户端在 HTTP 请求的以下部分中传递的信息来进行： 

- URL Path parameter
- URL Query parameter
- Header
- Request body

​	Echo 提供了不同的方法来执行绑定，每种方法在下面的部分中进行了描述。

## 结构体标签绑定

​	通过结构体绑定，您可以定义一个带有标签的 Go 结构体，指定数据源和相应的键。在您的请求处理程序中，只需调用 `Context#Bind(i interface{})` 并提供指向结构体的指针。标签告诉绑定器从请求中加载数据所需的一切。

​	在这个示例中，一个名为 `User` 的结构体类型告诉绑定器将查询字符串参数 `id` 绑定到其字符串字段 `ID`：

```go
type User struct {
  ID string `query:"id"`
}

// in the handler for /users?id=<userID>
// 在处理程序中用于 /users?id=<userID>
var user User
err := c.Bind(&user); if err != nil {
    return c.String(http.StatusBadRequest, "bad request")
}
```



### 数据源

​	Echo 支持以下指定数据源的标签：

- `query` —— 查询参数
- `param` —— 路径参数（也称为路由）
- `header` —— 标头参数
- `json` —— 请求体。使用内置的 Go [json](https://golang.org/pkg/encoding/json/) 包进行反序列化。
- `xml` —— 请求体。使用内置的 Go [xml](https://golang.org/pkg/encoding/xml/) 包进行反序列化。
- `form` —— 表单数据。值来自查询字符串和请求体。使用 Go 标准库的表单解析。

### 数据类型

​	在解码请求体时，根据 `Content-Type` 标头，支持以下数据类型：

- `application/json`
- `application/xml`
- `application/x-www-form-urlencoded`

​	当绑定路径参数、查询参数、标头或表单数据时，必须在每个结构字段上明确设置标签。但是，如果标签被省略，则会根据结构字段名进行 JSON 和 XML 绑定。这符合 [Go 的 json 包](https://pkg.go.dev/encoding/json#Unmarshal) 的行为。

​	对于表单数据，Echo 使用 Go 标准库的表单解析。如果内容类型不是 `MIMEMultipartForm`，则会从请求 URL 和请求体解析表单数据。请参阅 [non-MIMEMultipartForm](https://golang.org/pkg/net/http/#Request.ParseForm) 和 [MIMEMultipartForm](https://golang.org/pkg/net/http/#Request.ParseMultipartForm) 的文档。

### 多个数据源

​	可以在同一个字段上指定多个数据源。在这种情况下，请求数据按照以下顺序绑定： 

1. 路径参数
2. 查询参数（仅适用于 GET/DELETE 方法）
3. 请求体

```go
type User struct {
  ID string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
}
```



​	请注意，每个阶段的绑定将覆盖先前阶段绑定的数据。这意味着，如果您的 JSON 请求包含查询参数 `name=query` 和请求体 `{"name": "body"}`，那么结果将是 `User{Name: "body"}`。

### 直接数据源

​	还可以直接从特定的数据源中绑定数据：

请求体：

```go
err := (&DefaultBinder{}).BindBody(c, &payload)
```



查询参数：

```go
err := (&DefaultBinder{}).BindQueryParams(c, &payload)
```



路径参数：

```go
err := (&DefaultBinder{}).BindPathParams(c, &payload)
```



标头参数：

```go
err := (&DefaultBinder{}).BindHeaders(c, &payload)
```



Note that headers is not one of the included sources with `Context#Bind`. The only way to bind header data is by calling `BindHeaders` directly.

​	请注意，`Context#Bind`方法不包括标头（headers）作为绑定数据的源。绑定头部数据的唯一方法是直接调用`BindHeaders`方法。

### 安全性

​	为了保证应用程序的安全性，如果这些结构体包含不应绑定的字段，请避免直接将绑定的结构体传递给其他方法。建议为绑定创建一个单独的结构体，并将其显式映射到业务结构体。

​	考虑一下，如果绑定的结构体具有一个名为 `IsAdmin bool` 的导出字段，并且请求体包含 `{IsAdmin: true, Name: "hacker"}`。

### 示例

​	在这个示例中，我们定义了一个名为 `User` 的结构体类型，其中的字段标签用于绑定来自 `json`、`form` 或 `query` 请求数据：

```go
type User struct {
  Name  string `json:"name" form:"name" query:"name"`
  Email string `json:"email" form:"email" query:"email"`
}

type UserDTO struct {
  Name    string
  Email   string
  IsAdmin bool
}
```



​	并且在 POST `/users` 路由的处理程序中将请求数据绑定到该结构体：

```go
e.POST("/users", func(c echo.Context) (err error) {
  u := new(User)
  if err = c.Bind(u); err != nil {
    return c.String(http.StatusBadRequest, "bad request")
  }

  // 为了安全起见，加载到单独的结构体中
  user := UserDTO{
    Name: u.Name,
    Email: u.Email,
    IsAdmin: false // 避免绑定不应该绑定的字段
  }

  executeSomeBusinessLogic(user)
  
  return c.JSON(http.StatusOK, u)
})
```



#### JSON 数据

```sh
curl -X POST http://localhost:1323/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe","email":"joe@labstack"}'
```



#### Form 数据

```sh
curl -X POST http://localhost:1323/users \
  -d 'name=Joe' \
  -d 'email=joe@labstack.com'
```



#### 查询参数

```sh
curl -X GET 'http://localhost:1323/users?name=Joe&email=joe@labstack.com'
```



## 流畅绑定

​	Echo 提供了一个接口，用于从指定的数据源绑定显式的数据类型。它使用方法链，也称为[流畅接口](https://en.wikipedia.org/wiki/Fluent_interface)。

​	以下方法为绑定到 Go 数据类型提供了一些有用的方法。这些绑定器提供了流畅的语法，可以链接在一起配置和执行绑定，并处理错误。 

- `echo.QueryParamsBinder(c)` —— 绑定查询参数（源自 URL）
- `echo.PathParamsBinder(c)` —— 绑定路径参数（源自 URL）
- `echo.FormFieldBinder(c)` —— 绑定表单字段（源自 URL + 请求体）。也可参考 [Request.ParseForm](https://golang.org/pkg/net/http/#Request.ParseForm)。

### 错误处理

​	绑定器通常通过调用 `BindError()` 或 `BindErrors()` 来完成。如果发生任何错误，`BindError()` 返回遇到的第一个错误，而`BindErrors()` 返回所有绑定错误。绑定器中存储的任何错误也将被重置。

​	使用 `FailFast(true)` 可以配置绑定器，在遇到第一个错误时停止绑定，或者使用 `FailFast(false)` 执行整个绑定器调用链。快速失败默认启用，但在使用 `BindErrors()` 时应禁用。

### 示例

```go
// url =  "/api/search?active=true&id=1&id=2&id=3&length=25"
var opts struct {
  IDs []int64
  Active bool
}
length := int64(50) // 默认长度为 50

// 创建查询参数绑定器，它在遇到第一个错误时停止绑定
err := echo.QueryParamsBinder(c).
  Int64("length", &length).
  Int64s("ids", &opts.IDs).
  Bool("active", &opts.Active).
  BindError() // 返回第一个绑定错误
```



### 支持的数据类型

| 数据类型            | 说明                                                         |
| ------------------- | ------------------------------------------------------------ |
| `bool`              |                                                              |
| `float32`           |                                                              |
| `float64`           |                                                              |
| `int`               |                                                              |
| `int8`              |                                                              |
| `int16`             |                                                              |
| `int32`             |                                                              |
| `int64`             |                                                              |
| `uint`              |                                                              |
| `uint8/byte`        | 不支持 `bytes()`。请使用 `BindUnmarshaler`/`CustomFunc` 将值从 base64 等转换为 `[]byte{}`。 |
| `uint16`            |                                                              |
| `uint32`            |                                                              |
| `uint64`            |                                                              |
| `string`            |                                                              |
| `time`              |                                                              |
| `duration`          |                                                              |
| `BindUnmarshaler()` | 绑定到实现了 BindUnmarshaler 接口的类型                      |
| `TextUnmarshaler()` | 绑定到实现了 encoding.TextUnmarshaler 接口的类型             |
| `JsonUnmarshaler()` | 绑定到实现了 json.Unmarshaler 接口的类型                     |
| `UnixTime()`        | 将 Unix 时间（整数）转换为 `time.Time`                       |
| `UnixTimeMilli()`   | 将带有毫秒精度的 Unix 时间（整数）转换为 `time.Time`         |
| `UnixTimeNano()`    | 将带有纳秒精度的 Unix 时间（整数）转换为 `time.Time`         |
| `CustomFunc()`      | 自定义转换逻辑的回调函数                                     |

​	每种支持的类型都有以下方法：

- `<Type>("param", &destination)` —— 如果参数值存在，则将其绑定到给定类型的目标，例如 `Int64(...)`。
- `Must<Type>("param", &destination)` —— 参数值必须存在，将其绑定到给定类型的目标，例如 `MustInt64(...)`。
- `<Type>s("param", &destination)` —— （对于切片）如果参数值存在，则将其绑定到给定类型的目标，例如 `Int64s(...)`。
- `Must<Type>s("param", &destination)` ——（对于切片）参数值必须存在，将其绑定到给定类型的目标，例如 `MustInt64s(...)`。

​	对于某些切片类型，`BindWithDelimiter("param", &dest, ",")` 支持在进行类型转换之前拆分参数值。例如，从 URL `/api/search?id=1,2,3&id=1` 绑定一个整数切片将得到 `[]int64{1,2,3,1}`。

## 自定义绑定

​	可以使用 `Echo#Binder` 注册自定义绑定器。

```go
type CustomBinder struct {}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
  // 你可以使用默认的绑定器
  db := new(echo.DefaultBinder)
  if err := db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
    return
  }

  // 在这里定义你的自定义实现
  return
}
```