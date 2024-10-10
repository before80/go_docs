+++
title = "staticFiles"
date = 2024-02-04T10:04:34+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/view/static/]({{< ref "/beego/mvcIntroduction/views/staticFiles" >}})

# Static files 静态文件



## Static Files 静态文件

Go already has the built-in `http.ServeFile` package to serve static files. Beego made a wrapper for it. To register static files use:

​	Go 已经内置了 `http.ServeFile` 包来提供静态文件。Beego 为其制作了一个包装器。要注册静态文件，请使用：

```
web.SetStaticPath("/static","public")
```

- The first parameter is the url path
  第一个参数是 URL 路径
- The second parameter is the static file directory path. (relative to the application directory)
  第二个参数是静态文件目录路径。（相对于应用程序目录）

Beego supports multiple static file directories:

​	Beego 支持多个静态文件目录：

```
web.SetStaticPath("/images","images")
web.SetStaticPath("/css","css")
web.SetStaticPath("/js","js")
```

With the above settings, request `/images/login/login.png` will find `application_path/images/login/login.png` and request `/static/img/logo.png` will find `public/img/logo.png` file.

​	使用上述设置，请求 `/images/login/login.png` 将找到 `application_path/images/login/login.png` ，请求 `/static/img/logo.png` 将找到 `public/img/logo.png` 文件。

By default Beego will check if the file exists, if not it will return a 404 page. If the request is for `index.html`, because `http.ServeFile` will redirect and doesn’t display this page by default, you can set `web.BConfig.WebConfig.DirectoryIndex = true` to show `index.html` page. If this is enabled, users can see the file list while visit the directory.

​	默认情况下，Beego 将检查文件是否存在，如果不存在，它将返回 404 页面。如果请求是针对 `index.html` ，因为 `http.ServeFile` 将重定向并且默认情况下不会显示此页面，您可以设置 `web.BConfig.WebConfig.DirectoryIndex = true` 以显示 `index.html` 页面。如果启用此功能，用户在访问目录时可以看到文件列表。
