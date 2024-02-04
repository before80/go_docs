+++
title = "部分"
date = 2024-02-04T21:10:26+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/frontend-layer/partials/](https://gobuffalo.io/documentation/frontend-layer/partials/)

# Partials 部分 

This document only applies when using https://github.com/gobuffalo/buffalo/tree/main/render.

​	仅当使用 https://github.com/go/tree/main/render 时，此文档才适用。

Please see [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) for more details on the underlying templating package.

​	有关基础模板包的详细信息，请参阅 github.com/go/plush。

## Usage 用法 

You can call your partials using `partial` plush helper:

​	您可以使用 `partial` plush 帮助器调用您的部分：

templates/users/form.plush.html

templates/users/new.plush.html

Output
输出

```html
<form action="/users/" method="POST">
<!-- form content here  -->
<form>
```

## Context 上下文 

All [rendering context](https://gobuffalo.io/documentation/frontend-layer/rendering) from the parent template will automatically pass through to the partial, and any partials that partial may call. (see also [Context](https://gobuffalo.io/documentation/request_handling/context))

​	父模板的所有渲染上下文将自动传递到部分，以及该部分可能调用的任何部分。（另请参阅上下文）

actions/users.go

templates/users/edit.plush.html

templates/users/form.plush.html

Output
输出

```go
func UsersEdit(c buffalo.Context) error {
	user := User{
		Name: "John Smith",
	}
	// ...
	c.Set("user", user)
	return c.Render(http.StatusOK, render.HTML("users/edit.plush.html"))
}
```

## Local Context 局部上下文 

In addition to have the global [context](https://gobuffalo.io/documentation/request_handling/context), you can set additional variable only for partials as “local” variables.

​	除了拥有全局上下文外，您还可以仅为局部变量设置其他变量，作为“局部”变量。

actions/colors.go

templates/colors/index.plush.html

templates/colors/details.plush.html

Output
输出

```go
func ColorsHandler(c buffalo.Context) error {
  colors := map[string]interface{}{
		"White":  "#FFFFFF",
		"Maroon": "#800000",
		"Red":    "#FF0000",
		"Purple": "#800080",
	}

	c.Set("colors", colors)
	return c.Render(http.StatusOK, r.HTML("colors/index.plush.html"))
}
```

## Helpers 帮助程序 

Partials are not much different from standard [templates](https://gobuffalo.io/documentation/frontend-layer/templating) in Buffalo. They include all of the same [helpers](https://gobuffalo.io/documentation/frontend-layer/helpers) as well.

​	局部模板与 Buffalo 中的标准模板没有太大区别。它们也包含所有相同的帮助程序。