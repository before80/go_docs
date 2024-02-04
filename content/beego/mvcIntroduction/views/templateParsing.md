+++
title = "模板解析"
date = 2024-02-04T10:03:09+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/view/view/](https://beego.wiki/docs/mvc/view/view/)

# Template Parsing 模板解析



## Template Parsing 模板解析

Beego uses Go’s built-in package `html/template` as the template parser. Upon startup, it will compile and cache the templates into a map for efficient rendering.

&zeroWidthSpace;Beego 使用 Go 内置的包 `html/template` 作为模板解析器。在启动时，它会将模板编译并缓存到一个映射中，以便高效渲染。

## Template Directory 模板目录

The default template directory for Beego is `views`. Template files can be put into this directory and Beego will parse and cache them automatically. However if the development mode is enabled, Beego parses templates every time without caching. Beego can only have one template directory which can be customized:

&zeroWidthSpace;Beego 的默认模板目录是 `views` 。模板文件可以放入此目录，Beego 将自动解析并缓存它们。但是，如果启用了开发模式，Beego 每次都会解析模板，而不会缓存。Beego 只能有一个可以自定义的模板目录：

```
web.BConfig.WebConfig.ViewsPath = "myviewpath"
```

You can add alternative template directories by calling

&zeroWidthSpace;您可以通过调用添加备用模板目录

```
web.AddViewPath("moreViews")
```

This will parse and cache template files in this directory and allow you to use them by setting ViewPath on a Controller:

&zeroWidthSpace;这将在该目录中解析并缓存模板文件，并允许您通过在控制器上设置 ViewPath 来使用它们：

```
this.ViewPath = "moreViews"
```

Setting a ViewPath to a directory which was not previously registered with AddViewPath() will fail with “Unknown view path”

&zeroWidthSpace;将 ViewPath 设置为先前未使用 AddViewPath() 注册的目录将失败，并显示“未知视图路径”

## Auto Rendering 自动渲染

You don’t need to render and output templates manually. Beego will call Render automatically after finishing the method. You can disable auto rendering in the configuration file or in `main.go` if you don’t need it.

&zeroWidthSpace;您无需手动渲染和输出模板。Beego 将在完成方法后自动调用 Render。如果您不需要，可以在配置文件或 `main.go` 中禁用自动渲染。

In configuration file: 
&zeroWidthSpace;在配置文件中：

```
autorender = false
```

In main.go: 
&zeroWidthSpace;在 main.go 中：

```
web.BConfig.WebConfig.AutoRender = false
```

## Template Tags 模板标记

Go uses `{{` and `}}` as the default template tags. In the case that these tags conflict with other template tags as in AngularJS, we can use other tags. To do so, In configuration file:

&zeroWidthSpace;Go 使用 `{{` 和 `}}` 作为默认模板标记。如果这些标记与其他模板标记（如 AngularJS 中的标记）冲突，我们可以使用其他标记。为此，在配置文件中：

```
templateleft = <<<
templateright = >>>
```

Or, add these to the main.go:

&zeroWidthSpace;或者，将这些添加到 main.go 中：

```
web.BConfig.WebConfig.TemplateLeft = "<<<"
web.BConfig.WebConfig.TemplateRight = ">>>"
```

## Template Data 模板数据

Template gets its data from `this.Data` in Controller. So for example if you need `{{.Content}}` in the template, you need to assign it in the Controller first:

&zeroWidthSpace;模板从 Controller 中的 `this.Data` 获取数据。因此，例如，如果您需要在模板中使用 `{{.Content}}` ，则需要先在 Controller 中分配它：

```
this.Data["Content"] = "value"
```

Different rendering types:

&zeroWidthSpace;不同的渲染类型：

- struct 
  &zeroWidthSpace;结构体

  Struct variable: 
  &zeroWidthSpace;结构体变量：

  ```
    type A struct{
    	Name string
    	Age  int
    }
  ```

  Assign value in the Controller:

  &zeroWidthSpace;在 Controller 中分配值：

  ```
    this.Data["a"]=&A{Name:"astaxie",Age:25}
  ```

  Render it in the template:

  &zeroWidthSpace;在模板中渲染它：

  ```
    the username is {{.a.Name}}
    the age is {{.a.Age}}
  ```

- map

  Assign value in the Controller:

  &zeroWidthSpace;在 Controller 中分配值：

  ```
    mp["name"]="astaxie"
    mp["nickname"] = "haha"
    this.Data["m"]=mp
  ```

  Render it in the template:

  &zeroWidthSpace;在模板中渲染它：

  ```
    the username is {{.m.name}}
    the username is {{.m.nickname}}
  ```

- slice 
  &zeroWidthSpace;切片

  Assign value in the Controller:

  &zeroWidthSpace;在控制器中分配值：

  ```
    ss :=[]string{"a","b","c"}
    this.Data["s"]=ss
  ```

  Render it in the template:

  &zeroWidthSpace;在模板中呈现它：

  ```
    {{range $key, $val := .s}}
    {{$key}}
    {{$val}}
    {{end}}
  ```

## Template Name 模板名称

> From version 1.6: this.TplNames is this.TplName
>
> &zeroWidthSpace;从版本 1.6 开始：this.TplNames 为 this.TplName

Beego uses Go’s built-in template engine, so the syntax is same as Go. To learn more about template see [Templates](https://github.com/Unknwon/build-web-application-with-golang_EN/blob/master/eBook/07.4.md).

&zeroWidthSpace;Beego 使用 Go 内置的模板引擎，因此语法与 Go 相同。要详细了解模板，请参阅模板。

You can set the template name in Controller and Beego will find the template file under the viewpath and render it automatically. In the config below, Beego will find add.tpl under admin and render it.

&zeroWidthSpace;您可以在控制器中设置模板名称，Beego 将在 viewpath 下找到模板文件并自动呈现它。在下面的配置中，Beego 将在 admin 下找到 add.tpl 并呈现它。

```go
this.TplName = "admin/add.tpl"
```

Beego supports `.tpl` and `.html` file extensions by default. If you’re using other extensions, you must set it in the configuration first:

&zeroWidthSpace;Beego 默认支持 `.tpl` 和 `.html` 文件扩展名。如果您使用其他扩展名，则必须首先在配置中设置它：

```go
beego.AddTemplateExt("file_extension_you_need")
```

If `TplName` is not set in the Controller while `autorender` is enabled, Beego will use `TplName` as below by default:

&zeroWidthSpace;如果在启用 `autorender` 时未在控制器中设置 `TplName` ，则 Beego 将默认使用 `TplName` ，如下所示：

```go
c.TplName = strings.ToLower(c.controllerName) + "/" + strings.ToLower(c.actionName) + "." + c.TplExt
```

It is Controller name + “/” + request method name + “.” + template extension. So if the Controller name is `AddController`, request method is `POST` and the default template extension is `tpl`, Beego will render `/viewpath/addcontroller/post.tpl` template file.

&zeroWidthSpace;它是控制器名称 + “/” + 请求方法名称 + “.” + 模板扩展名。因此，如果控制器名称为 `AddController` ，请求方法为 `POST` ，默认模板扩展名为 `tpl` ，则 Beego 将呈现 `/viewpath/addcontroller/post.tpl` 模板文件。

## Layout Design 布局设计

Beego supports layout design. For example, if in your application the main navigation and footer does not change and only the content part is different, you can use a layout like this:

&zeroWidthSpace;Beego 支持布局设计。例如，如果在您的应用程序中，主导航和页脚不会改变，只有内容部分不同，您可以使用这样的布局：

```go
this.Layout = "admin/layout.html"
this.TplName = "admin/add.tpl"
```

In `layout.html` you must set a variable like this:

&zeroWidthSpace;在 `layout.html` 中，您必须设置一个这样的变量：

```
{{.LayoutContent}}
```

Beego will parse the file named `TplName` and assign it to `LayoutContent` then render `layout.html`.

&zeroWidthSpace;Beego 将解析名为 `TplName` 的文件，并将其分配给 `LayoutContent` ，然后渲染 `layout.html` 。

Beego will cache all the template files. You can also implement a layout this way:

&zeroWidthSpace;Beego 将缓存所有模板文件。您还可以通过这种方式实现布局：

```
{{template "header.html"}}
Logic code
{{template "footer.html"}}
```

## LayoutSection 布局部分

`LayoutContent` is a little complicated, as it can include Javascript and CSS. Since in most situations having only one `LayoutContent` is not enough, there is an attribute called `LayoutSections` in `Controller`. It allows us to set multiple `section` in `Layout` page and each `section` can contain its own sub-template page.

&zeroWidthSpace; `LayoutContent` 有点复杂，因为它可以包含 Javascript 和 CSS。由于在大多数情况下，仅有一个 `LayoutContent` 是不够的，因此在 `Controller` 中有一个称为 `LayoutSections` 的属性。它允许我们在 `Layout` 页面中设置多个 `section` ，并且每个 `section` 都可以包含其自己的子模板页面。

layout_blog.tpl: 
&zeroWidthSpace;layout_blog.tpl：

```
<!DOCTYPE html>
<html>
<head>
    <title>Lin Li</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css">
    <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap-theme.min.css">
    {{.HtmlHead}}
</head>
<body>

    <div class="container">
        {{.LayoutContent}}
    </div>
    <div>
        {{.SideBar}}
    </div>
    <script type="text/javascript" src="http://code.jquery.com/jquery-2.0.3.min.js"></script>
    <script src="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
    {{.Scripts}}
</body>
</html>
```

html_head.tpl: 
&zeroWidthSpace;html_head.tpl：

```
<style>
     h1 {
        color: red;
     }
</style>
```

scripts.tpl：

```
<script type="text/javascript">
    $(document).ready(function() {
        // bla bla bla
    });
</script>
```

Here is the logic in the Controller:

&zeroWidthSpace;以下是控制器中的逻辑：

```go
type BlogsController struct {
    web.Controller
}

func (this *BlogsController) Get() {
    this.Layout = "layout_blog.tpl"
    this.TplName = "blogs/index.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["HtmlHead"] = "blogs/html_head.tpl"
    this.LayoutSections["Scripts"] = "blogs/scripts.tpl"
    this.LayoutSections["Sidebar"] = ""
}
```

## Another approach 另一种方法

We can also just specify the template the controller is going to use and let the template system handle the layout:

&zeroWidthSpace;我们还可以只指定控制器将要使用的模板，并让模板系统处理布局：

for example: 
&zeroWidthSpace;例如：

controller: 
&zeroWidthSpace;控制器：

```go
this.TplName = "blog/add.tpl"
this.Data["SomeVar"] = "SomeValue"
this.Data["Title"] = "Add"
```

template add.tpl: 
&zeroWidthSpace;模板 add.tpl：

```
{{ template "layout_blog.tpl" . }}
{{ define "css" }}
		<link rel="stylesheet" href="/static/css/current.css">
{{ end}}


{{ define "content" }}
		<h2>{{ .Title }}</h2>
		<p> This is SomeVar: {{ .SomeVar }}</p>
{{ end }}

{{ define "js" }}
	<script src="/static/js/current.js"></script>
{{ end}}
```

layout_blog.tpl: 
&zeroWidthSpace;layout_blog.tpl：

```
<!DOCTYPE html>
<html>
<head>
    <title>{{ .Title }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css">
    <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap-theme.min.css">
     {{ block "css" . }}{{ end }}
</head>
<body>

    <div class="container">
        {{ block "content" . }}{{ end }}
    </div>
    <script type="text/javascript" src="http://code.jquery.com/jquery-2.0.3.min.js"></script>
    <script src="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
     {{ block "js" . }}{{ end }}
</body>
</html>
```

Using `block` action instead of `template` allows us to have default block content and skipping blocks that we don’t need in every template (for example, if we don’t need css block in `add.tpl` template - we will not define it and that won’t raise an error)

&zeroWidthSpace;使用 `block` 操作代替 `template` 允许我们拥有默认块内容，并跳过我们不需要的块（例如，如果我们不需要 `add.tpl` 模板中的 css 块 - 我们将不会定义它，并且不会引发错误）

## renderform

Define struct: 
&zeroWidthSpace;定义结构：

```go
type User struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age,text,age:"`
	Sex   string
	Intro string `form:",textarea"`
}
```

- StructTag definition uses `form` as tag. It uses the same tags as [Parse Form](https://beego.wiki/docs/mvc/controller/params#parse-to-struct). There are three optional params separated by ‘,’:

  &zeroWidthSpace;StructTag 定义使用 `form` 作为标签。它使用与解析表单相同的标签。有三个可选参数，用“,”分隔：

  - The first param is `name` attribute of the form field. If empty, it will use `struct field name` as the value.
    第一个参数是表单字段的 `name` 属性。如果为空，它将使用 `struct field name` 作为值。
  - The second param is the form field type. If empty, it is assumed as `text`.
    第二个参数是表单字段类型。如果为空，则假定为 `text` 。
  - The third param is the tag of form field. If empty, it will use `struct field name` as the tag name.
    第三个参数是表单字段的标签。如果为空，它将使用 `struct field name` 作为标签名。

- If the `form` tag only has one value, it is the `name` attribute of the form field. Except last value can be ignore all the other place must be separated by ‘,’. E.g.: `form:",,username:"`

  &zeroWidthSpace;如果 `form` 标签只有一个值，则它是表单字段的 `name` 属性。除了最后一个值可以忽略之外，其他所有地方都必须用“,”分隔。例如： `form:",,username:"`

- To ignore a field there are two ways:

  &zeroWidthSpace;忽略字段有两种方法：

  - The first way is to use lowercase for the field name in the struct.
    第一种方法是将结构体中的字段名使用小写。
  - The second way is to set `-` as the value of `form` tag.
    第二种方法是将 `-` 作为 `form` 标签的值。

controller：

```go
func (this *AddController) Get() {
    this.Data["Form"] = &User{}
    this.TplName = "index.tpl"
}
```

The param of Form must be a pointer to a struct.

&zeroWidthSpace;Form 的参数必须是结构体的指针。

template: 
&zeroWidthSpace;template：

```
<form action="" method="post">
{{.Form | renderform}}
</form>
```

The code above will generate the form below:

&zeroWidthSpace;以上编码会生成如下表格：

```
Name: <input name="username" type="text" value="test"></br>
Age: <input name="age" type="text" value="0"></br>
Gender: <input name="Sex" type="text" value=""></br>
Intro: <input name="Intro" type="textarea" value="">
```