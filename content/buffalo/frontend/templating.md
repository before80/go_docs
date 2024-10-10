+++
title = "模板化"
date = 2024-02-04T21:09:47+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/frontend-layer/templating/]({{< ref "/buffalo/frontend/templating" >}})

# Templating 模板化 

Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).

​	此文档仅适用于使用 github.com/gobuffalo/buffalo/render 时。

Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para más detalles sobre el paquete de plantillas.

​	有关模板包的更多详细信息，请查阅 github.com/gobuffalo/plush。

Buffalo defaults to using [plush](https://github.com/gobuffalo/plush) as its template engine.

​	Buffalo 默认使用 plush 作为其模板引擎。

## Introduction to Plush Plush 简介 

{{< vimeo "207200621">}}

## Plush - Tips, Tricks and Testing Plush - 技巧、窍门和测试 

{{< vimeo "267643588">}}

## General Usage 常规用法 

Plush allows you to capture the `context` variables to use anywhere in your templates.

​	Plush 允许您捕获 `context` 变量，以便在模板中的任何位置使用。

actions/index.go

templates/index.plush.html

Output
输出

```go
func myHandler(c buffalo.Context) error {
  c.Set("name", "John Smith")
  return c.Render(http.StatusOK, r.HTML("index.plush.html"))
}
```

## Plush Examples Plush 示例 

#### Conditional Statements 条件语句 

IF

ELSE

ELSE IF

Multiple Conditions
多个条件

```erb
<%= if (true) { %>
  <!-- some template content -->
<% } %>
```

actions/main.go

templates/index.plush.html

Output
输出

```go
func MyHandler(c buffalo.Context) error {
	// ...
	c.Set("userName", "John Smith")
	return c.Render(http.StatusOK, r.HTML("templates/index.plush.html"))
}
```

### Iterating 迭代 

#### Through Slices 遍历切片 

When looping through `slices`, the block being looped through will have access to the “global” context.

​	当遍历 `slices` 时，被遍历的块将能够访问“全局”上下文。

The `for` statement takes 1 or 2 arguments. When using the two arguments version, the first argument is the “index” of the loop and the second argument is the value from the slice.

​	语句 `for` 采用 1 个或 2 个参数。使用两个参数版本时，第一个参数是循环的“索引”，第二个参数是切片中的值。

actions/main.go

Loop using 2 Arguments
使用 2 个参数的循环

Loop using 1 Arguments
使用 1 个参数的循环

```go
func MyHandler(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  return c.Render(http.StatusOK, r.HTML("index.plush.html"))
}
```

#### Through Maps 通过映射 

Looping through `maps` using the `each` helper is also supported, and follows very similar guidelines to looping through `arrays`.

​	使用 `each` 帮助器循环遍历 `maps` 也受支持，并且遵循与循环遍历 `arrays` 非常相似的准则。

When using the two argument version, the first argument is the key of the map and the second argument is the value from the map:

​	使用两个参数版本时，第一个参数是映射的键，第二个参数是映射中的值：

actions/main.go

Loop using 2 Arguments
使用 2 个参数的循环

Loop using 1 Arguments
使用 1 个参数的循环

```go
func ColorsHandler(c buffalo.Context) error {
	colors := map[string]interface{}{
		"White":  "#FFFFFF",
		"Maroon": "#800000",
		"Red":    "#FF0000",
		"Purple": "#800080",
	}

	c.Set("colors", colors)
	return c.Render(http.StatusOK, r.HTML("home/colors.plush.html"))
}
```

You can see more examples in [plush repository](https://github.com/gobuffalo/plush).
更多示例请见 plush 代码库。
