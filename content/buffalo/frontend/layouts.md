+++
title = "布局"
date = 2024-02-04T21:10:17+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/frontend-layer/layouts/]({{< ref "/buffalo/frontend/layouts" >}})

# Layouts 布局 

Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).

​	此文档仅适用于使用 github.com/gobuffalo/buffalo/render 时。

Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para más detalles sobre el paquete de plantillas.

​	有关模板包的更多详细信息，请查阅 github.com/gobuffalo/plush。

## Using a Standard Layout 使用标准布局 

It is quite common to want to use the same layout across most, if not all of an application. When creating a new `render.Engine` the `HTMLLayout` property can be set to a file that will automatically be used by the `render.HTML` function.

​	通常需要在大多数（如果不是全部）应用程序中使用相同的布局。创建新的 `render.Engine` 时，可以将 `HTMLLayout` 属性设置为一个文件，该文件将自动被 `render.HTML` 函数使用。

actions/render.go

templates/application.plush.html

templates/hello.plush.html

actions/hello.go

Output
输出

```go
var r *render.Engine

func init() {
  r = render.New(render.Options{
    // ...
    HTMLLayout: "application.plush.html",
    // ...
  })
}
```

## Using a Custom Layout 使用自定义布局 

Sometimes, on certain requests, a different layout is needed. This alternate layout can be passed in as the second parameter to `render.HTML`.

​	有时，在某些请求中，需要不同的布局。此备用布局可以作为第二个参数传递给 `render.HTML` 。

Custom layouts do **NOT** work with **`render.Auto`**.
自定义布局不适用于 `render.Auto` 。

actions/render.go

templates/custom.plush.html

templates/hello.plush.html

actions/hello.go

Output
输出

```go
var r *render.Engine

func init() {
  r = render.New(render.Options{
    // ...
    HTMLLayout: "application.plush.html", // You can realize that render continues using the application.plush.html
    // ...
  })
}
```
