+++
title = "渲染"
date = 2024-02-04T21:09:28+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/frontend-layer/rendering/]({{< ref "/buffalo/frontend/rendering" >}})

# Rendering 渲染 

The [https://github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render) [[godoc\]](https://pkg.go.dev/github.com/gobuffalo/buffalo/render) package implements that interface, and has a collection of useful render types already defined. It is recommended that you use this package, but feel free and write your own renderers!

​	https://github.com/go/render [godoc] 程序包实现了该界面，并且已经定义了一组有用的渲染器。建议您使用此包，但也可以随时编写自己的渲染器！

This document only applies when using https://github.com/gobuffalo/buffalo/tree/main/render.

​	仅当使用 https://github.com/go/tree/main/render 时，此文档才适用。

Please see [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) for more details on the underlying templating package.

​	有关基础模板包的详细信息，请参阅 github.com/go/plush。

## Render Auto 渲染自动 

Since **0.11.0**
自 0.11.0 起



In many cases, you’ll have to provide the same contents in different formats: JSON, XML, HTML… Buffalo provides an easy way to do that using a single statement.

​	在许多情况下，您将不得不以不同的格式提供相同的内容：JSON、XML、HTML……Buffalo 提供了一种使用单个语句轻松做到这一点的方法。

```go
func Beatles(c buffalo.Context) error {
  members := models.Members{}
  // ...
  return c.Render(http.StatusOK, r.Auto(c, members))
}
```

{{< vimeo "257736901">}}

## JSON and XML JSON 与 XMl 

When rendering JSON, or XML, using the [`render.JSON`](https://pkg.go.dev/github.com/gobuffalo/buffalo/render#JSON) or [`render.XML`](https://pkg.go.dev/github.com/gobuffalo/buffalo/render#XML), you pass the value that you would like to be marshaled and the appropriate marshaler will encode the value you passed and write it to the response with the correct content/type.

​	在使用 `render.JSON` 或 `render.XML` 渲染 JSON 或 XMl 时，您传递要编组的值，相应的编组器会对您传递的值进行编码，并以正确的 content/type 将其写入响应。

**NOTE:** If you already have a string that contains JSON or XML, do **NOT** use these methods as they will attempt to marshal the string into JSON or XML causing strange responses. What you could do instead is write a **custom render** function as explained in the [Custom Rendering](https://gobuffalo.io/documentation/frontend-layer/rendering/#custom-rendering) section.
注意：如果您已经拥有包含 JSON 或 XML 的字符串，请勿使用这些方法，因为它们会尝试将字符串编组到 JSON 或 XML 中，从而导致奇怪的响应。相反，您可以按照自定义渲染部分中的说明编写自定义渲染函数。

```go
// models/user.go

type User struct {
	FirstName string
	LastName  string
	Gender    string
}
```

JSON

XML

```go
func MyHandler(c buffalo.Context) error {
  user := models.User{
		FirstName: "John",
		LastName:  "Smith",
		Gender:    "Male",
	}

  return c.Render(http.StatusOK, r.JSON(user))
}
// output
{
  "FirstName": "John",
  "LastName": "Smith",
  "Gender": "Male"
}
```

## Markdown

Files passed into the [`render.HTML`](https://pkg.go.dev/github.com/gobuffalo/buffalo/render#Engine.HTML) or [`render.Template`](https://pkg.go.dev/github.com/gobuffalo/buffalo/render#Engine.Template) methods, that have an extension of `.plush.md`, will be converted from Markdown (using GitHub flavored Markdown) to HTML before being run through the templating engine. This makes for incredibly easy templating for simpler pages.

​	传递给 `render.HTML` 或 `render.Template` 方法的文件，其扩展名为 `.plush.md` ，将在通过模板引擎运行之前从 Markdown（使用 GitHub 风格的 Markdown）转换为 HTML。这使得为更简单的页面进行模板化变得非常容易。

```md
<!-- beatles.plush.md -->

# The Beatles

<%= for (name) in names { %>
* <%= name %>
<% } %>
// actions/beatles.go

func Beatles(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})

  return c.Render(http.StatusOK, r.HTML("beatles.plush.md"))
}
<!-- output -->
<h1>The Beatles</h1>

<ul>
  <li><p>John</p></li>
  <li><p>Paul</p></li>
  <li><p>George</p></li>
  <li><p>Ringo</p></li>
</ul>
```

## JavaScript

Since **0.10.0**
自 0.10.0 起



The [`render`](https://godoc.org/github.com/gobuffalo/buffalo/render) package has a new implementation of [`render.Renderer`](https://godoc.org/github.com/gobuffalo/buffalo/render#Renderer), [`render.JavaScript`](https://godoc.org/github.com/gobuffalo/buffalo/render#JavaScript).

​	 `render` 包具有 `render.Renderer` 、 `render.JavaScript` 的新实现。

This means inside of an action you can do the following:

​	这意味着您可以在操作中执行以下操作：

```go
func HomeHandler(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.JavaScript("index.js"))
}
```

The [`render.Options`](https://godoc.org/github.com/gobuffalo/buffalo/render#Options) type now has a new attribute, `JavaScriptLayout`. This new option is similar to the `HTMLLayout` option in that it will wrap `*.js` files inside of another `*.js`.

​	 `render.Options` 类型现在具有一个新属性 `JavaScriptLayout` 。此新选项类似于 `HTMLLayout` 选项，因为它会将 `*.js` 文件包装在另一个 `*.js` 中。

The new JavaScript renderer also has it’s own implementation of the `partial` function. This new implementation behaves almost the same as the original implementation, but is smart enough to know that if you are rendering an `*.html` file inside of a `*.js` file that it will need to be escaped properly, and so it does it for you.

​	新的 JavaScript 渲染器也有自己的 `partial` 函数实现。此新实现的行为几乎与原始实现相同，但足够智能，可以知道如果您正在 `*.js` 文件中渲染 `*.html` 文件，则需要正确转义，因此它会为您执行此操作。

```javascript
$("#new-goal-form").replaceWith("<%= partial("goals/new.html") %>");
```

## Automatic Extensions 自动扩展 

Since **0.10.2**
自 0.10.2 起



You can use `HTML`, `Javascript` and `Markdown` renderers without specifying the file extension:

​	您可以使用 `HTML` 、 `Javascript` 和 `Markdown` 渲染器，而无需指定文件扩展名：

```go
// actions/beatles.go
func Beatles(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  // Render beatles.html
  return c.Render(http.StatusOK, r.HTML("beatles"))
}
```

This works with [partials]({{< ref "/buffalo/frontend/partials" >}}) too.
这也适用于部分内容。

## Download files 下载文件 

The [`r.Download`](https://pkg.go.dev/github.com/gobuffalo/buffalo/render#Engine.Download) method allows you to download files in your application easily.

​	 `r.Download` 方法允许您轻松地在应用程序中下载文件。

```go
func DownloadHandler(c buffalo.Context) error {
	// ...
	f, err := os.Open("your/path/file_name.extension")
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, r.Download(c, "file_name.extension", f))
}
```

## Custom Rendering 自定义渲染 

For another type of rendering, the [`r.Func`](https://godoc.org/github.com/gobuffalo/buffalo/render#Func) method allows you to pass in a content type and a function to render your data to the provided `io.Writer`, which is commonly, the HTTP response, in particular, a [`*buffalo.Response`](https://godoc.org/github.com/gobuffalo/buffalo#Response).

​	对于其他类型的渲染， `r.Func` 方法允许您传入内容类型和函数，以将数据渲染到提供的 `io.Writer` ，通常是 HTTP 响应，特别是 `*buffalo.Response` 。

```go
func MyHandler(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.Func("application/csv", csvWriter))
}

func csvWriter(w io.Writer, d render.Data) error {
  cw := csv.NewWriter(w)
  if err := cw.Write([]string{"a", "b", "c"}); err != nil {
    return errors.WithStack(err)
  }
  cw.Flush()
  return nil
}
```

For smaller, or one off situations, using an anonymous function can be even easier. In this example you can see how to use an anonymous function to render a string that already contains JSON.

​	对于较小或一次性情况，使用匿名函数甚至可以更轻松。在此示例中，您可以看到如何使用匿名函数来渲染已包含 JSON 的字符串。

```go
var myJSONString string
return c.Render(http.StatusOK, r.Func("application/json", func(w io.Writer, d render.Data) error {
  _, err := w.Write([]byte(myJSONString))
  return err
}))
```

## Renderer Interface 渲染器接口 

In order for a renderer to be able to be used with [`Context#Render`]({{< ref "/buffalo/requestHandling/context#context-and-rendering" >}}) it must implement the following interface:

​	为了使渲染器能够与 `Context#Render` 一起使用，它必须实现以下接口：

```go
// Renderer interface that must be satisified to be used with
// buffalo.Context.Render
type Renderer interface {
  ContentType() string
  Render(io.Writer, Data) error
}

// Data type to be provided to the Render function on the
// Renderer interface.

type Data map[string]interface{}
```

The [https://github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/master/render) [[godoc\]](https://godoc.org/github.com/gobuffalo/buffalo/render) package implements that interface, and has a collection of useful render types already defined. It is recommended that you use this package, but feel free and write your own renderers!

​	https://github.com/gobuffalo/buffalo/render [godoc] 包实现了该接口，并且已经定义了一组有用的渲染器类型。建议您使用此包，但也可以随意编写自己的渲染器！
