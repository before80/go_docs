+++
title = "URL 构建"
date = 2024-02-04T09:58:26+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/urlbuilding/](https://beego.wiki/docs/mvc/controller/urlbuilding/)

# URL Building - URL 构建



## URL Building URL 构建

If it can match URLs, can Beego also generate them? Of course it can. To build a URL to a specific function you can use the URLFor() function. It accepts the name of the function of Controller as first argument and a number of keyword arguments, each corresponding to the variable part of the URL rule. Unknown variable parts are appended to the URL as query parameters. Here are some examples:

&zeroWidthSpace;如果它可以匹配 URL，那么 Beego 是否也可以生成它们？当然可以。要构建到特定函数的 URL，可以使用 URLFor() 函数。它接受控制器函数的名称作为第一个参数，以及许多关键字参数，每个参数都对应于 URL 规则的可变部分。未知的可变部分作为查询参数附加到 URL。这里有一些示例：

Here is the controller definition:

&zeroWidthSpace;这是控制器定义：

```
type TestController struct {
	web.Controller
}

func (this *TestController) Get() {
	this.Data["Username"] = "astaxie"
	this.Ctx.Output.Body([]byte("ok"))
}

func (this *TestController) List() {
	this.Ctx.Output.Body([]byte("i am list"))
}

func (this *TestController) Params() {
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Params["0"] + this.Ctx.Input.Params["1"] + this.Ctx.Input.Params["2"]))
}

func (this *TestController) Myext() {
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":ext")))
}

func (this *TestController) GetUrl() {
	this.Ctx.Output.Body([]byte(this.URLFor(".Myext")))
}
```

This is how you register the router:

&zeroWidthSpace;这是注册路由器的方式：

```
web.Router("/api/list", &TestController{}, "*:List")
web.Router("/person/:last/:first", &TestController{})
web.AutoRouter(&TestController{})
```

This is how you generate the url:

&zeroWidthSpace;这是生成 URL 的方式：

```
web.URLFor("TestController.List")
// Output /api/list

web.URLFor("TestController.Get", ":last", "xie", ":first", "asta")
// Output /person/xie/asta

web.URLFor("TestController.Myext")
// Output /Test/Myext

web.URLFor("TestController.GetUrl")
// Output /Test/GetUrl
```

## This is how you use it in a template 这是在模板中使用它的方式

beego has already registered the template function `urlfor`. You can use it like this:

&zeroWidthSpace;beego 已经注册了模板函数 `urlfor` 。您可以像这样使用它：

```
{{urlfor "TestController.List"}}
// Output /api/list

{{urlfor "TestController.Get" ":last" "xie" ":first" "asta"}}
// Output /person/xie/asta
```

Why would you want to build URLs instead of hard-coding them into your templates? There are three good reasons for this:

&zeroWidthSpace;为什么要构建 URL 而不是将它们硬编码到模板中？这样做有三个很好的理由：

1. Reversing is often more descriptive than hard-coding the URLs. More importantly, it allows you to change URLs in one go, without having to remember to change URLs all over the place.
   反向解析通常比硬编码 URL 更具描述性。更重要的是，它允许您一次更改 URL，而无需记住在所有地方更改 URL。
2. URL building will handle escaping of special characters and Unicode data transparently for you, so you don’t have to deal with them.
   URL 构建将透明地处理特殊字符和 Unicode 的转义，因此您不必处理它们。