+++
title = "视图"
date = 2024-02-04T09:10:48+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/quickstart/view/]({{< ref "/beego/quickStart/view" >}})

# View 视图



## Creating views 创建视图

In the previous example, when creating the Controller the line `this.TplName = "index.tpl"` was used to declare the template to be rendered. By default `beego.Controller` supports `tpl` and `html` extensions. Other extensions can be added by calling `beego.AddTemplateExt`.

​	在前面的示例中，创建 Controller 时，使用行 `this.TplName = "index.tpl"` 声明要呈现的模板。默认情况下， `beego.Controller` 支持 `tpl` 和 `html` 扩展。可通过调用 `beego.AddTemplateExt` 添加其他扩展。

Beego uses the default `html/template` engine built into Go, so view displays show data using standard Go templates. You can find more information about using Go templates at [*Building Web Applications with Golang*](https://github.com/astaxie/build-web-application-with-golang/blob/master/en/07.4.md).

​	Beego 使用内置于 Go 中的默认 `html/template` 引擎，因此视图显示使用标准 Go 模板显示数据。您可以在使用 Golang 构建 Web 应用程序中找到有关使用 Go 模板的更多信息。

Let’s look at an example:

​	我们来看一个示例：

```
<!DOCTYPE html>

<html>
    <head>
        <title>Beego</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    </head>
    <body>
        <header class="hero-unit" style="background-color:#A9F16C">
            <div class="container">
                <div class="row">
                    <div class="hero-text">
                        <h1>Welcome to Beego!</h1>
                        <p class="description">
                            Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
                            <br />
                            Official website: <a href="http://{{.Website}}">{{.Website}}</a>
                            <br />
                            Contact me: {{.Email}}
                        </p>
                    </div>
                </div>
            </div>
        </header>
    </body>
</html>
```

The data was assigned into a map type variable `Data` in the controller, which is used as the rendering context. The data can now be accessed and output by using the keys `.Website` and `.Email`.

​	数据已分配给控制器的映射类型变量 `Data` ，该变量用作呈现上下文。现在可以使用键 `.Website` 和 `.Email` 访问和输出数据。

[The next section]({{< ref "/beego/quickStart/staticFiles" >}}). will describe how to serve static files.

​	下一节将介绍如何提供静态文件。
