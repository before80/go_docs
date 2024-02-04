+++
title = "请求绑定"
date = 2024-02-04T21:08:10+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/request_handling/bind/]({{< ref "/buffalo/requestHandling/requestBinding" >}})

# Request Binding 请求绑定 

The `buffalo.Context` interface has a method named, `Bind`. This method allows for the binding of a request, such as a form or JSON body, to be mapped to a struct for easy form handling.

​	 `buffalo.Context` 接口有一个名为 `Bind` 的方法。此方法允许将请求（例如表单或 JSON 正文）的绑定映射到一个结构，以便轻松处理表单。

For more information on request binding in Buffalo, see the godoc for [github.com/gobuffalo/buffalo/binding](https://godoc.org/github.com/gobuffalo/buffalo/binding)

​	有关 Buffalo 中请求绑定的更多信息，请参阅 github.com/gobuffalo/buffalo/binding 的 godoc

## How Does It Work? 它是如何工作的？ 

The `Bind` function works by looking at the `Content-Type` or `Accept` header of a request and looking for a mapped implementation of the [`binding.Binder`](https://godoc.org/github.com/gobuffalo/buffalo/binding#Binder) interface.

​	 `Bind` 函数通过查看请求的 `Content-Type` 或 `Accept` 标头并查找 `binding.Binder` 接口的映射实现来工作。

#### Mapped Content Types (HTML) 映射的内容类型 (HTML) 

- `application/html`
- `text/html`
- `application/x-www-form-urlencoded`
- `multipart/form-data`
- `html`

#### Mapped Content Types (JSON) 映射的内容类型 (JSON) 

- `application/json`
- `text/json`
- `json`

#### Mapped Content Types (XML) 映射的内容类型 (XML) 

- `application/xml`
- `text/xml`
- `xml`

## HTML/Form Binding HTML/表单绑定 

Binding HTML forms, by default, uses the [github.com/monoculum/formam](https://github.com/monoculum/formam) package to bind HTML forms to a struct.

​	默认情况下，绑定 HTML 表单使用 github.com/monoculum/forma 包将 HTML 表单绑定到结构。 保留此符号 采用以下 结构和下面的 HTML 表单： 保留此符号 在操作中，我们可以将此 HTML 表单绑定到 结构，如下所示： 表单和结构之间的默认映射是结构上属性的名称，并且应与表单字段的 属性匹配。请注意，示例中的 字段同时匹配结构属性和表单字段上的 属性。 保留此符号 通过使用 结构标记，我们可以将 HTML 表单中的字段映射到 结构，包括通过使用 忽略 。 保留此符号 有关 结构标记的更多信息，请参阅 github.com/monoculum/forma 文档。 保留此符号 JSON 和 XML 绑定 # 保留此符号 默认情况下，绑定 JSON 和 XML 请求使用标准库中的 和 包。 保留此符号 采用以下 结构和下面的 JSON 有效负载（XML 完全相同，但用 代替 结构标记）。 保留此符号 在操作中，我们可以将此 JSON 有效负载绑定到 结构，如下所示：

Take the following `User` struct and the HTML form below.

```go
type User struct {
  Name     string `form:"name"`
  Email    string
  Password string `form:"-"`
}
<form>
  <input type="text" value="ringo" name="name"/>
  <input type="text" value="ringo@beatles.com" name="Email"/>
  <input type="text" value="starr" name="Password"/>
  <input type="submit"/>
</form>
```

In an action we can bind this HTML form to the `User` struct as follows:

```go
func MyAction(c buffalo.Context) error {
  u := &User{}
  if err := c.Bind(u); err != nil {
    return err
  }
  u.Name // "Ringo"
  u.Email // "ringo@beatles.com"
  u.Password // ""
  // do more work
}
```

The default mapping between the form and struct is the name of the attribute on the struct, and should match the `name` attribute of the form field. Notice the `Email` field in the examples matches both the struct attribute and the `name` attribute on the form field.

By using the `form` struct tags we can map the fields in the HTML form to the `User` struct, including ignoring the `Password` by using a `-`.

Please refer to the [github.com/monoculum/formam](https://github.com/monoculum/formam) docs for more information about the `form` struct tag.

## JSON and XML Binding

Binding JSON and XML requests, by default, uses the `encoding/json` and `encoding/xml` packages in the standard library.

Take the following `User` struct and the JSON payload below. (XML works exactly the same, but instead of `json` struct tags, substitute `xml` instead.)

```go
type User struct {
  Name     string `json:"name"`
  Email    string `json:"email"`
  Password string `json:"-"`
}
{
  "name": "Ringo",
  "email": "ringo@beatles.com",
  "password": "starr"
}
```

In an action we can bind this JSON payload to the `User` struct as follows:

```go
func MyAction(c buffalo.Context) error {
  u := &User{}
  if err := c.Bind(u); err != nil {
    return err
  }
  u.Name // "Ringo"
  u.Email // "ringo@beatles.com"
  u.Password // ""
  // do more work
}
```

By using the `json` or `xml` struct tags we can map the fields in the JSON payload to the `User` struct, including ignoring the `Password` by using a `-`.

​	通过使用 `json` 或 `xml` 结构体标签，我们可以将 JSON 有效负载中的字段映射到 `User` 结构体，包括通过使用 `-` 忽略 `Password` 。

Please refer to the standard library docs for more information about the `json` and `xml` struct tags.

​	有关 `json` 和 `xml` 结构体标签的更多信息，请参阅标准库文档。

## Registering a Custom Binder 注册自定义绑定器 

Perhaps you don’t want to use the default binders, or you have a different content type you want to be able to bind to, for example, let’s say you want to bind a YAML request.

​	也许您不想使用默认绑定器，或者您有想要能够绑定的不同内容类型，例如，假设您想绑定一个 YAML 请求。

```go
package actions

import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

func init() {
  binding.Register("text/yaml", func(req *http.Request, model interface{}) error {
    b, err := ioutil.ReadAll(req.Body)
    if err != nil {
      return err
    }
    return yaml.Unmarshal(b, model)
  })
}
```
