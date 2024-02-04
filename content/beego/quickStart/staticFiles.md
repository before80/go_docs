+++
title = "静态文件"
date = 2024-02-04T09:10:59+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/quickstart/static/](https://beego.wiki/docs/quickstart/static/)

# Static files 静态文件



## Handling static files 处理静态文件

Most applications include numerous static files such as images, js, css and more. To support this requirement the Beego project skeleton incorporates folders for these files by default.

&zeroWidthSpace;大多数应用程序包含许多静态文件，例如图像、js、css 等。为了支持此要求，Beego 项目框架默认情况下为这些文件合并文件夹。

```
├── static
│   ├── css
│   ├── img
│   └── js
```

Beego registers the static directory as the static path by default. Registered rule: URL prefix with directory mapping

&zeroWidthSpace;Beego 默认将静态目录注册为静态路径。已注册规则：带目录映射的 URL 前缀

```
StaticDir["/static"] = "static"
```

You can register multiple static directories. For example two different download directories, `download1` and `download2`, can be set using:

&zeroWidthSpace;您可以注册多个静态目录。例如，可以使用以下方式设置两个不同的下载目录 `download1` 和 `download2` ：

```
web.SetStaticPath("/down1", "download1")
web.SetStaticPath("/down2", "download2")
```

Visiting the URL `http://localhost/down1/123.txt` will request the file `123.txt` in the `download1` directory. To remove the default `/static -> static` mapping use `web.DelStaticPath("/static")`.

&zeroWidthSpace;访问 URL `http://localhost/down1/123.txt` 将请求 `download1` 目录中的文件 `123.txt` 。要删除默认 `/static -> static` 映射，请使用 `web.DelStaticPath("/static")` 。

# Implementation 实现

To implement this in a Web Application register the Static directory to your `routes.go` files

&zeroWidthSpace;要在 Web 应用程序中实现此功能，请将静态目录注册到您的 `routes.go` 文件

```
web.SetStaticPath("/down1", "download1")
```

Once the file is save it can be accessed from the browser.

&zeroWidthSpace;保存文件后，即可从浏览器访问该文件。