+++
title = "简介"
date = 2024-01-31T19:02:59+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-chi.io/#/README](https://go-chi.io/#/README)

# Introduction 简介

## Hi, Let's Get You Started With chi 👋 嗨，让我们开始使用 chi

`chi` is a lightweight, idiomatic and composable router for building Go HTTP services. It's especially good at helping you write large REST API services that are kept maintainable as your project grows and changes. `chi` is built on the new `context` package introduced in Go 1.7 to handle signaling, cancelation and request-scoped values across a handler chain.

​	 `chi` 是一个轻量级、惯用且可组合的路由器，用于构建 Go HTTP 服务。它特别擅长帮助您编写大型 REST API 服务，随着项目的增长和变化，这些服务可以保持可维护性。 `chi` 基于 Go 1.7 中引入的新 `context` 包构建，用于处理跨处理程序链的信令、取消和请求范围的值。

The focus of the project has been to seek out an elegant and comfortable design for writing REST API servers, written during the development of the Pressly API service that powers our public API service, which in turn powers all of our client-side applications.

​	该项目的重点是寻求一种优雅且舒适的设计来编写 REST API 服务器，该设计是在开发 Pressly API 服务期间编写的，该服务为我们的公共 API 服务提供支持，进而为我们所有的客户端应用程序提供支持。

The key considerations of chi's design are: project structure, maintainability, standard http handlers (stdlib-only), developer productivity, and deconstructing a large system into many small parts. The core router `github.com/go-chi/chi` is quite small (less than 1000 LOC), but we've also included some useful/optional subpackages: [middleware](https://github.com/go-chi/chi/tree/master/middleware), [render](https://github.com/go-chi/render) and [docgen](https://github.com/go-chi/docgen). We hope you enjoy it too!

​	chi 设计的关键考虑因素是：项目结构、可维护性、标准 http 处理程序（仅限 stdlib）、开发人员生产力和将大型系统分解为许多小部分。核心路由器 `github.com/go-chi/chi` 非常小（不到 1000 LOC），但我们还包含了一些有用/可选的子包：中间件、渲染器和 docgen。我们希望您也喜欢它！

## Features 功能

- **Lightweight** - cloc'd in ~1000 LOC for the chi router
  轻量级 - chi 路由器约为 1000 LOC
- **Fast** - yes, see [benchmarks](https://github.com/go-chi/chi#benchmarks)
  快速 - 是的，请参阅基准
- **100% compatible with net/http** - use any http or middleware pkg in the ecosystem that is also compatible with `net/http`
  100% 兼容 net/http - 在生态系统中使用任何与 `net/http` 兼容的 http 或中间件软件包
- **Designed for modular/composable APIs** - middlewares, inline middlewares, route groups and sub-router mounting
  专为模块化/可组合 API 而设计 - 中间件、内联中间件、路由组和子路由挂载
- **Context control** - built on new `context` package, providing value chaining, cancellations and timeouts
  上下文控制 - 基于新的 `context` 包构建，提供值链接、取消和超时
- **Robust** - in production at Pressly, CloudFlare, Heroku, 99Designs, and many others (see [discussion](https://github.com/go-chi/chi/issues/91))
  强大 - 在 Pressly、CloudFlare、Heroku、99Designs 和许多其他公司投入生产（请参阅讨论）
- **Doc generation** - `docgen` auto-generates routing documentation from your source to JSON or Markdown
  文档生成 - `docgen` 从您的源代码自动生成 JSON 或 Markdown 格式的路由文档
- **Go.mod support** - as of v5, go.mod support (see [CHANGELOG](https://github.com/go-chi/chi/blob/master/CHANGELOG.md))
  支持 Go.mod - 从 v5 开始，支持 go.mod（请参阅 CHANGELOG）
- **No external dependencies** - plain ol' Go stdlib + net/http
  无外部依赖项 - 纯粹的 Go stdlib + net/http

## Examples 示例

See [examples](https://github.com/go-chi/chi/blob/master/_examples/) for a variety of examples.

​	请参阅示例，了解各种示例。

## License 许可证

Copyright (c) 2015-present [Peter Kieltyka](https://github.com/pkieltyka)

​	版权所有 (c) 2015-至今 Peter Kieltyka

Licensed under [MIT License](https://github.com/go-chi/chi/blob/master/LICENSE)

​	根据 MIT 许可证授权