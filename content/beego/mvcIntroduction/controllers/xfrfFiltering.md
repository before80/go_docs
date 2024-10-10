+++
title = "XSRF 过滤"
date = 2024-02-04T09:57:42+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/xsrf/]({{< ref "/beego/mvcIntroduction/controllers/xfrfFiltering" >}})

# XSRF filtering - XSRF 过滤



## Cross-Site Request Forgery 跨站请求伪造

XSRF, [Cross-Site Request Forgery](http://en.wikipedia.org/wiki/Cross-site_request_forgery), is an important security concern for web development. Beego has built in XSRF protection which assigns each user a randomized cookie that is used to verify requests. XSRF protection can be activated by setting `EnableXSRF = true` in the configuration file:

​	XSRF（跨站请求伪造）是 Web 开发中一个重要的安全问题。Beego 内置了 XSRF 保护，它为每个用户分配一个随机 cookie，用于验证请求。XSRF 保护可以通过在配置文件中设置 `EnableXSRF = true` 来激活：

```
EnableXSRF = true
XSRFKey = 61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o
XSRFExpire = 3600 // set cookie expire in 3600 seconds, default to 60 seconds if not specified
```

XSRF protection can also be enabled in the main application entry function:

​	XSRF 保护也可以在主应用程序入口函数中启用：

```
web.BConfig.WebConfig.EnableXSRF = true
web.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
web.BConfig.WebConfig.XSRFExpire = 3600
```

When XSRF is enabled Beego will set a cookie `_xsrf` for every user. Beego will refuse any `POST`, `PUT`, or `DELETE` request that does not include this cookie. If XSRF protection is enabled a field must be added to provide an `_xsrf` value to every form. This can be added directly in the template with `XSRFFormHTML()`.

​	启用 XSRF 后，Beego 将为每个用户设置一个 cookie `_xsrf` 。Beego 将拒绝任何不包含此 cookie 的 `POST` 、 `PUT` 或 `DELETE` 请求。如果启用了 XSRF 保护，则必须向每个表单添加一个字段以提供 `_xsrf` 值。这可以直接在模板中使用 `XSRFFormHTML()` 添加。

A global expiration time should be set using `web.XSRFExpire`. This value can be also be set for individual logic functions:

​	应使用 `web.XSRFExpire` 设置全局过期时间。此值也可以为各个逻辑函数设置：

```go
func (this *HomeController) Get(){
	this.XSRFExpire = 7200
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
}
```

**XSRF** works with HTTPS protocol. In Beego 2.x, the cookie storing XSRF token has two flag: [secure](https://en.wikipedia.org/wiki/Secure_cookie) and [http-only](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies).

​	XSRF 与 HTTPS 协议配合使用。在 Beego 2.x 中，存储 XSRF 令牌的 cookie 有两个标志：secure 和 http-only。

In Beego 1.x (<=1.12.2), we don’t have this two flags, so it’s not safe because attackers is able to steal the XSRF token.

​	在 Beego 1.x（<=1.12.2）中，我们没有这两个标志，因此不安全，因为攻击者能够窃取 XSRF 令牌。
