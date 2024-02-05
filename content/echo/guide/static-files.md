+++
title = "静态文件"
weight = 120
date = 2023-07-09T21:51:46+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Static Files - 静态文件

> 原文：[https://echo.labstack.com/docs/static-files](https://echo.labstack.com/docs/static-files)

​	图片、JavaScript、CSS、PDF、字体等等...

## 使用静态中间件

[参见](https://echo.labstack.com/docs/middleware/static)。

## 使用 Echo#Static()

​	`Echo#Static(prefix, root string)` 注册一个新的路由，以路径前缀来提供静态文件，这些文件存放在指定的根目录下。

*用法 1*

```go
e := echo.New()
e.Static("/static", "assets")
```



​	上述示例将为路径 `/static/*` 下的任何文件提供来自 assets 目录的文件。例如，对 `/static/js/main.js` 的请求将获取并提供 `assets/js/main.js` 文件。

*用法 2*

```go
e := echo.New()
e.Static("/", "assets")
```



​	上述示例将为路径 `/*` 下的任何文件提供来自 assets 目录的文件。例如，对 `/js/main.js` 的请求将获取并提供 `assets/js/main.js` 文件。

## 使用 Echo#File()

​	`Echo#File(path, file string)` 注册一个新的路由，以提供指定的静态文件。

*用法1*

​	从 `public/index.html` 提供一个索引页面：

```go
e.File("/", "public/index.html")
```



*用法 2*

​	从 `images/favicon.ico` 提供一个网站图标

```go
e.File("/favicon.ico", "images/favicon.ico")
```