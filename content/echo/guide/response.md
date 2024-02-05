+++
title = "响应"
weight = 100
date = 2023-07-09T21:51:30+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Response - 响应

> 原文：[https://echo.labstack.com/docs/response](https://echo.labstack.com/docs/response)

## 发送字符串

​	可以使用 `Context#String(code int, s string)` 发送带有状态码的纯文本响应。

*示例*

```go
func(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
```



## 发送HTML（参考模板）

​	可以使用 `Context#HTML(code int, html string)` 发送带有状态码的简单HTML响应。如果您想要发送动态生成的HTML，请参阅[模板](https://echo.labstack.com/docs/templates)。

*示例*

```go
func(c echo.Context) error {
  return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
}
```



### 发送HTML Blob

​	可以使用 `Context#HTMLBlob(code int, b []byte)` 发送带有状态码的HTML blob。在使用输出 `[]byte` 的模板引擎时，这可能会非常方便。

## 渲染模板

[了解更多](https://echo.labstack.com/docs/templates)

## 发送JSON

​	可以使用 `Context#JSON(code int, i interface{})` 将提供的Go类型编码为JSON，并以带有状态码的响应发送它。

*示例*

```go
// User
type User struct {
  Name  string `json:"name" xml:"name"`
  Email string `json:"email" xml:"email"`
}

// Handler
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  return c.JSON(http.StatusOK, u)
}
```



### Stream JSON

​	`Context#JSON()` 内部使用 `json.Marshal`，对于大型JSON可能效率不高，此时可以直接流式传输（directly stream）JSON。

*示例*

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return json.NewEncoder(c.Response()).Encode(u)
}
```



### JSON 美化

​	`Context#JSONPretty(code int, i interface{}, indent string)` 可以用于发送基于缩进进行漂亮打印的JSON响应，缩进可以是空格或制表符。

​	下面的示例发送了一个使用空格缩进的漂亮打印JSON：

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "joe@labstack.com",
  }
  return c.JSONPretty(http.StatusOK, u, "  ")
}
```



```js
{
  "email": "joe@labstack.com",
  "name": "Jon"
}
```



> 提示
>
> ​	您还可以使用`Context#JSON()`在请求URL查询字符串中附加`pretty`以输出漂亮打印的JSON（使用空格缩进）。

*示例*

```sh
curl http://localhost:1323/users/1?pretty
```



### JSON Blob

​	`Context#JSONBlob(code int, b []byte)` 可以用于直接从外部源（例如数据库）发送预编码的（pre-encoded）JSON blob。

*示例*

```go
func(c echo.Context) error {
  encodedJSON := []byte{} // Encoded JSON from external source
  return c.JSONBlob(http.StatusOK, encodedJSON)
}
```



## 发送JSONP

​	`Context#JSONP(code int, callback string, i interface{})` 可以将提供的Go类型编码为JSON，并使用回调作为JSONP有效负载发送，带有状态码。

[示例](https://echo.labstack.com/docs/cookbook/jsonp)

## 发送XML

​	`Context#XML(code int, i interface{})` 可以将提供的Go类型编码为XML，并以带有状态码的响应发送它。

*示例*

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  return c.XML(http.StatusOK, u)
}
```



### Stream XML

​	`Context#XML` 内部使用 `xml.Marshal`，对于大型XML可能效率不高，此时可以直接流式（directly stream）传输XML。

*示例*

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return xml.NewEncoder(c.Response()).Encode(u)
}
```



### XML Pretty

​	`Context#XMLPretty(code int, i interface{}, indent string)` 可以用于发送基于缩进进行漂亮打印的XML响应，缩进可以是空格或制表符。

​	下面的示例发送了一个使用空格缩进的漂亮打印XML：

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "joe@labstack.com",
  }
  return c.XMLPretty(http.StatusOK, u, "  ")
}
```



```xml
<?xml version="1.0" encoding="UTF-8"?>
<User>
  <Name>Jon</Name>
  <Email>joe@labstack.com</Email>
</User>
```



> 提示
>
> ​	您还可以使用`Context#XML()`在请求URL查询字符串中附加`pretty`以输出漂亮打印的XML（使用空格缩进）。

*示例*

```sh
curl http://localhost:1323/users/1?pretty
```



### XML Blob

​	`Context#XMLBlob(code int, b []byte)` 可以用于直接从外部源（例如数据库）发送预编码的（pre-encoded）XML blob。

*示例*

```go
func(c echo.Context) error {
  encodedXML := []byte{} //  从外部源编码的XML
  return c.XMLBlob(http.StatusOK, encodedXML)
}
```



## 发送文件

​	`Context#File(file string)` 可以用于将文件内容作为响应发送。它会**自动设置正确的内容类型**并优雅地处理缓存。

*示例*

```go
func(c echo.Context) error {
  return c.File("<PATH_TO_YOUR_FILE>")
}
```



## 发送附件

​	`Context#Attachment(file, name string)` 类似于 `File()`，但用于将文件作为附件发送，并提供名称。

*示例*

```go
func(c echo.Context) error {
  return c.Attachment("<PATH_TO_YOUR_FILE>", "<ATTACHMENT_NAME>")
}
```



## 发送 Inline

​	`Context#Inline(file, name string)` 类似于 `File()`，但用于将文件作为内联（inline）发送，并提供名称。

*示例*

```go
func(c echo.Context) error {
  return c.Inline("<PATH_TO_YOUR_FILE>")
}
```



## Send Blob

​	`Context#Blob(code int, contentType string, b []byte)` 可以用于使用提供的内容类型和状态码发送任意数据响应。

*示例*

```go
func(c echo.Context) (err error) {
  data := []byte(`0306703,0035866,NO_ACTION,06/19/2006
      0086003,"0005866",UPDATED,06/19/2006`)
    return c.Blob(http.StatusOK, "text/csv", data)
}
```



## Send Stream

​	`Context#Stream(code int, contentType string, r io.Reader)` 可以用于使用提供的内容类型、`io.Reader` 和状态码发送任意（arbitrary）数据流响应。

*示例*

```go
func(c echo.Context) error {
  f, err := os.Open("<PATH_TO_IMAGE>")
  if err != nil {
    return err
  }
  return c.Stream(http.StatusOK, "image/png", f)
}
```



## Send No Content

​	`Context#NoContent(code int)` 可以用于发送空的响应体，并带有状态码。

*示例*

```go
func(c echo.Context) error {
  return c.NoContent(http.StatusOK)
}
```



## 重定向请求

​	`Context#Redirect(code int, url string)` 可以用于将请求重定向到提供的URL，并带有状态码。

*示例*

```go
func(c echo.Context) error {
  return c.Redirect(http.StatusMovedPermanently, "<URL>")
}
```



## 钩子函数

### 在响应之前

​	`Context#Response#Before(func())` 可以用于注册在响应被写入之前调用的函数。

### 在响应之后

​	`Context#Response#After(func())` 可以用于注册在响应被写入之后调用的函数。如果 "Content-Length" 未知，则不执行任何后续函数。

*示例*

```go
func(c echo.Context) error {
  c.Response().Before(func() {
    println("before response")
  })
  c.Response().After(func() {
    println("after response")
  })
  return c.NoContent(http.StatusNoContent)
}
```



> 提示
>
> ​	可以注册多个 Before 和 After 函数。