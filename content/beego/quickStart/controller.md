+++
title = "控制器"
date = 2024-02-04T09:10:26+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/quickstart/controller/]({{< ref "/beego/quickStart/controller" >}})

# Controller 控制器



## Controller logic 控制器逻辑

The previous section covered user requests to controllers. This section will explain how to write a controller. Let’s start with some code:

​	上一节介绍了用户对控制器的请求。本节将说明如何编写控制器。我们从一些代码开始：

```
package controllers

import (
        "github.com/beego/beego/v2/server/web"
)

type MainController struct {
        web.Controller
}

func (this *MainController) Get() {
        this.Data["Website"] = "beego.wiki"
        this.Data["Email"] = "astaxie@gmail.com"
        this.TplName = "index.tpl" // version 1.6 use this.TplName = "index.tpl"
}
```

The following is a breakdown of the different sections of this code.

​	以下是此代码不同部分的细分。

## How Beego dispatches requests Beego 如何分派请求

The `MainController` is the first thing created. It contains an anonymous struct field of type `web.Controller`. This is called struct embedding and is the way that Go mimics inheritance. Because of this `MainController` automatically acquires all the methods of `web.Controller`.

​	 `MainController` 是创建的第一个内容。它包含一个类型为 `web.Controller` 的匿名结构字段。这称为结构嵌入，是 Go 模仿继承的方式。因此， `MainController` 自动获取 `web.Controller` 的所有方法。

`web.Controller` has several functions such as `Init`, `Prepare`, `Post`, `Get`, `Delete` and `Head`. These functions can be overwritten by implementing them. In this example the `Get` method was overwritten.

​	 `web.Controller` 具有多个函数，例如 `Init` 、 `Prepare` 、 `Post` 、 `Get` 、 `Delete` 和 `Head` 。可以通过实现这些函数来覆盖它们。在此示例中，覆盖了 `Get` 方法。

We talked about the fact that Beego is a RESTful framework so our requests will run the related `req.Method` method by default. For example, if the browser sends a `GET` request, it will execute the `Get` method in `MainController`. Therefore the `Get` method and the logic we defined above will be executed.

​	我们讨论了 Beego 是一个 RESTful 框架这一事实，因此我们的请求将默认运行相关的 `req.Method` 方法。例如，如果浏览器发送 `GET` 请求，它将在 `MainController` 中执行 `Get` 方法。因此，将执行 `Get` 方法和我们上面定义的逻辑。

## The `Get` method `Get` 方法

The logic of the `Get` method only outputs data. This data will be stored in `this.Data`, a `map[interface{}]interface{}`. Any type of data can be assigned here. In this case only two strings are assigned.

​	方法 `Get` 的逻辑仅输出数据。此数据将存储在 `this.Data` 中，这是一个 `map[interface{}]interface{}` 。可以在这里分配任何类型的数据。在这种情况下，仅分配两个字符串。

Finally the template will be rendered. `this.TplName` (v1.6 uses `this.TplName`) specifies the template which will be rendered. In this case it is `index.tpl`. If a template is not set it will default to `controller/method_name.tpl`. For example, in this case it would try to find `maincontroller/get.tpl`.

​	最后将渲染模板。 `this.TplName` （v1.6 使用 `this.TplName` ）指定要渲染的模板。在这种情况下，它是 `index.tpl` 。如果未设置模板，则默认为 `controller/method_name.tpl` 。例如，在这种情况下，它将尝试查找 `maincontroller/get.tpl` 。

There is no need to render manually. Beego will call the `Render` function (which is implemented in `web.Controller`) automatically if it is set up in the template.

​	无需手动渲染。如果在模板中设置了 `Render` 函数（在 `web.Controller` 中实现），Beego 将自动调用该函数。

Check the controller section in the [MVC Introduction]({{< ref "/beego/mvcIntroduction" >}}) to learn more about these functions. [The next section]({{< ref "/beego/quickStart/models" >}}) will describe model writing.

​	查看 MVC 简介中的控制器部分以了解有关这些函数的更多信息。下一节将介绍模型编写。
