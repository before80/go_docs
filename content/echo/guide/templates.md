+++
title = "模板"
weight = 130
date = 2023-07-09T21:52:14+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Templates - 模板

> 原文：[https://echo.labstack.com/docs/templates](https://echo.labstack.com/docs/templates)

## 渲染

​	`Context#Render(code int, name string, data interface{}) error` 使用数据渲染模板，并发送一个带有指定状态码的 text/html 响应。我们可以通过设置 `Echo.Renderer` 来注册模板引擎，以便使用任何模板引擎。

​	下面的示例演示了如何使用 Go 的 `html/template`：

1. 实现 `echo.Renderer` 接口

   ```go
   type Template struct {
       templates *template.Template
   }
   
   func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
       return t.templates.ExecuteTemplate(w, name, data)
   }
   ```

   

2. 预编译模板

   `public/views/hello.html`

   ```html
   {{define "hello"}}Hello, {{.}}!{{end}}
   ```

   

   ```go
   t := &Template{
       templates: template.Must(template.ParseGlob("public/views/*.html")),
   }
   ```

   

3. 注册模板

   ```go
   e := echo.New()
   e.Renderer = t
   e.GET("/hello", Hello)
   ```

   

4. 在处理程序中渲染模板

   ```go
   func Hello(c echo.Context) error {
       return c.Render(http.StatusOK, "hello", "World")
   }
   ```

   

## 高级 —— 在模板中调用 Echo

​	在某些情况下，从模板中生成URI可能非常有用。为了做到这一点，你需要在模板本身中调用`Echo#Reverse`。Golang的`html/template`包并不是最适合这个任务的，但可以通过两种方式来实现：为传递给模板的所有对象提供一个公共方法，或者通过传递`map[string]interface{}`并在自定义渲染器中增强此对象。鉴于后一种方法的灵活性，以下是一个示例程序：

template.html

```html
<html>
    <body>
        <h1>Hello {{index . "name"}}</h1>

        <p>{{ with $x := index . "reverse" }}
           {{ call $x "foobar" }} &lt;-- this will call the $x with parameter "foobar"
           {{ end }}
        </p>
    </body>
</html>
```

server.go

```go
package main

import (
    "html/template"
    "io"
    "net/http"

    "github.com/labstack/echo/v4"
)

// TemplateRenderer 是用于 Echo 框架的自定义 html/template 渲染器
type TemplateRenderer struct {
    templates *template.Template
}

// Render 渲染模板文档
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

    // 如果数据是一个 map，添加全局方法
    if viewContext, isMap := data.(map[string]interface{}); isMap {
        viewContext["reverse"] = c.Echo().Reverse
    }

    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
  e := echo.New()
  renderer := &TemplateRenderer{
      templates: template.Must(template.ParseGlob("*.html")),
  }
  e.Renderer = renderer

  // 命名路由为 "foobar"
  e.GET("/something", func(c echo.Context) error {
      return c.Render(http.StatusOK, "template.html", map[string]interface{}{
          "name": "Dolly!",
      })
  }).Name = "foobar"

  e.Logger.Fatal(e.Start(":8000"))
}
```



