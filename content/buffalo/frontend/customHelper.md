+++
title = "自定义助手"
date = 2024-02-04T21:10:56+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/frontend-layer/custom-helpers/]({{< ref "/buffalo/frontend/customHelper" >}})

# Custom Helpers 自定义助手 

This document only applies when using https://github.com/gobuffalo/buffalo/tree/main/render.

​	仅当使用 https://github.com/go/tree/main/render 时，此文档才适用。

Please see [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) for more details on the underlying templating package.

​	有关基础模板包的详细信息，请参阅 github.com/go/plush。

No templating package would be complete without allowing for you to build your own, custom, helper functions.

​	如果没有允许您构建自己的自定义助手函数，那么任何模板包都是不完整的。

{{< vimeo "229572343">}}

## Registering Helpers 注册助手 

Helper functions can be registered in two different places, depending on how they are to be used.

​	助手函数可以在两个不同的地方注册，具体取决于它们的使用方式。

### Global Helpers 全局助手 

*Most* helpers will be global helpers, meaning that they should be included in every template. The types of helpers can can be set in `actions/render.go`:

​	大多数助手将是全局助手，这意味着它们应该包含在每个模板中。助手类型可以在 `actions/render.go` 中设置：

```go
func init() {
  r = render.New(render.Options{
    // ...
    Helpers: render.Helpers{
      "myHelper": func() string {
        return "hello"
      },
    },
    // ...
  })
}
```

### Per Request Helpers 按请求助手 

Other helpers, that are specific to a certain request can be added to the `buffalo.Context` for that request.

​	其他特定于某个请求的助手可以添加到该请求的 `buffalo.Context` 中。

```go
func HomeHandler(c buffalo.Context) error {
  // ...
  c.Set("myHelper", func() string {
    return "hello"
  })
  // ...
}
```

## Return Values 返回值 

Plush allows you to return any values you would like from a helper function. This guide will focus on helpers that are designed to generate “output”.

​	Plush 允许您从辅助函数返回任何您想要的值。本指南将重点介绍旨在生成“输出”的辅助函数。

When returning multiple values from a function, the first value will be the one used for rendering in the template. If the last return value is an `error`, Plush will handle that error.

​	从函数返回多个值时，第一个值将用于在模板中渲染。如果最后一个返回值是 `error` ，Plush 将处理该错误。

------

#### `string`

Return just a `string`. The `string` will be HTML escaped, and deemed “not”-safe.

​	仅返回一个 `string` 。 `string` 将进行 HTML 转义，并被视为“不”安全。

```go
func() string {
  return ""
}
```

------

#### `template.HTML`

https://golang.org/pkg/html/template/#HTML

Return a `template.HTML` string. The `template.HTML` will **not** be HTML escaped, and will be deemed safe.

​	返回一个 `template.HTML` 字符串。 `template.HTML` 将不会进行 HTML 转义，并将被视为安全。

```go
func() template.HTML {
  return template.HTML("")
}
```

## Input Values 输入值 

Custom helper functions can take any type, and any number of arguments. You can even use variadic functions. There is an optional last argument, [`plush.HelperContext`](https://godoc.org/github.com/gobuffalo/plush#HelperContext), that can be received. It’s quite useful, and I would recommend taking it, as it provides you access to things like the context of the call, the block associated with the helper, etc…

​	自定义辅助函数可以采用任何类型和任意数量的参数。您甚至可以使用可变参数函数。有一个可选的最后一个参数 `plush.HelperContext` ，可以接收。它非常有用，我建议您使用它，因为它可以让您访问诸如调用上下文、与辅助函数关联的块等内容。

## Simple Helpers 简单辅助函数 

```go
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "greet": func(name string) string {
      return fmt.Sprintf("Hi %s!", name)
    },
  },
  // ...
})
```

The `greet` function is now available to all templates that use that `render.Engine`.

​	 `greet` 函数现在可用于所有使用该 `render.Engine` 的模板。

```go
// actions/greet.go
func Greeter(c buffalo.Context) error {
  c.Set("name", "Mark")
  return c.Render(200, r.String("<h1><%= greet(name) %></h1>"))
}
// output
<h1>Hi Mark!</h1>
```

## Block Helpers 块辅助函数 

Like the `if` or `for` statements, block helpers take a “block” of text that can be evaluated and potentially rendered, manipulated, or whatever you would like. To write a block helper, you have to take the `plush.HelperContext` as the last argument to your helper function. This will give you access to the block associated with that call.

​	与 `if` 或 `for` 语句一样，块助手采用可以评估并可能呈现、操作或您希望执行任何操作的文本“块”。要编写块助手，您必须将 `plush.HelperContext` 作为助手函数的最后一个参数。这将使您能够访问与该调用关联的块。

actions/render.go

helper
助手

actions/upper.go

templates/up.html

Output
输出

```go
// actions/render.go
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "upblock": upblock,
  },
  // ...
})
```

## Getting Values From the Context 从上下文获取值 

actions/render.go

helper
助手

action
操作

template
模板

Output
输出

```go
// actions/render.go
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "is_logged_in": isLoggedIn,
  },
  // ...
})
```
