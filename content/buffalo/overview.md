+++
title = "概述"
date = 2024-02-04T21:04:20+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/overview/]({{< ref "/buffalo/overview" >}})

# Overview 概述 

Welcome aboard!

​	欢迎 aboard！

While Buffalo can be considered as a framework, it’s mostly an ecosystem of Go and Javascript libraries curated to fit together. Most of these components can be switched for another, but we’ll only provide support for this default mix.

​	虽然 Buffalo 可以被认为是一个框架，但它主要是一个 Go 和 Javascript 库的生态系统，经过精心策划以相互配合。这些组件中的大多数都可以切换为另一个组件，但我们只为这种默认组合提供支持。

In this chapter, we’ll make a tour of the default bricks shipped with your Buffalo app.

​	在本章中，我们将介绍随 Buffalo 应用程序一起提供的默认构建块。

## Backend libraries 后端库 

### buffalo

Buffalo is the “glue” between all the provided components. It wraps the libraries and manages the workflow.

​	Buffalo 是所有提供的组件之间的“粘合剂”。它封装了库并管理工作流。

### gorilla/mux

[gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) is one of the most used routers in Go. While some routers are faster (like [httprouter](https://github.com/julienschmidt/httprouter)), gorilla/mux is the one providing the most features while being fast enough.

​	gorilla/mux 是 Go 中使用最多的路由器之一。虽然有些路由器更快（如 httprouter），但 gorilla/mux 是提供最多功能且足够快的路由器。

### pop

[pop](https://github.com/gobuffalo/pop) is the default ORM for Buffalo. It provides the `soda` toolbox to help you with your database needs and supports several databases, such as PostgreSQL, MySQL and SQLite.

​	pop 是 Buffalo 的默认 ORM。它提供了 `soda` 工具箱来帮助您满足数据库需求，并支持多种数据库，例如 PostgreSQL、MySQL 和 SQLite。

### plush

[plush](https://github.com/gobuffalo/plush) is the default templating engine for Buffalo. Its syntax is close to ERB templates (in Ruby).

​	plush 是 Buffalo 的默认模板引擎。它的语法接近 ERB 模板（在 Ruby 中）。

## Frontend libraries 前端库 

### Bootstrap

[Bootstrap](https://getbootstrap.com/) is one of the most famous frontend toolkit library. It helps to build responsive interfaces using common components like tables, carousels or grid layouts.

​	Bootstrap 是最著名的前端工具包库之一。它有助于使用常见组件（如表格、旋转木马或网格布局）构建响应式界面。

### jQuery

[jQuery](https://jquery.com/) is a rich library aiming to make DOM manipulation and AJAX queries simple. While it’s less used now, many projects still have it as a side-companion to help supporting all the browsers.

​	jQuery 是一个丰富的库，旨在简化 DOM 操作和 AJAX 查询。虽然现在使用较少，但许多项目仍然将其作为辅助工具来帮助支持所有浏览器。

### Webpack

[Webpack](https://webpack.js.org/) is a well-known Javascript assets bundler. It will take care of your Javascript, CSS, images and static assets files.

​	Webpack 是一个著名的 Javascript 资产捆绑器。它将负责处理您的 Javascript、CSS、图像和静态资产文件。

Webpack is configured by default to hash and minify your assets.

​	Webpack 默认配置为对您的资产进行哈希和压缩。

## Next Steps 后续步骤 

- [Installation]({{< ref "/buffalo/gettingStarted/installBuffalo">}}) - Install Buffalo!
  安装 - 安装 Buffalo！
