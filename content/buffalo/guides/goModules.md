+++
title = "Go 模块"
date = 2024-02-04T21:18:12+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/gomods/]({{< ref "/buffalo/guides/goModules" >}})

# Go Modules Go 模块 

**NOTE**: Please read https://github.com/golang/go/wiki/Modules to understand more about Go Modules **before** using them.

​	注意：在使用 Go 模块之前，请阅读 https://github.com/golang/go/wiki/Modules 以便更多地了解 Go 模块。

## Enabling Go Module Support 启用 Go 模块支持 

Since **v0.13.0**
自 v0.13.0 起



The support for Go Modules in “Buffalo” packages is experimental, as are Go Modules (as of `v1.11.x`). To “opt-in” to using Go Modules you need to turn them using the `GO111MODULE` environment variable and setting it to `on`.

​	“Buffalo” 包中的 Go 模块支持是实验性的，就像 Go 模块一样（截至 `v1.11.x` ）。要“选择加入”使用 Go 模块，您需要使用 `GO111MODULE` 环境变量并将其设置为 `on` 来启用它们。

This is **REQUIRED** to use Go Modules with “Buffalo” packages. The `auto` setting for `GO111MODULE` is **NOT** supported.

​	这对于将 Go 模块与“Buffalo”包一起使用是必需的。不支持 `GO111MODULE` 的 `auto` 设置。

```bash
$ export GO111MODULE=on
```

## Working Outside of the `GOPATH` 在 `GOPATH` 外部工作 

In addition to repeatable builds, Go Modules, allows you to easily work outside of the `GOPATH`.

​	除了可重复构建之外，Go 模块还允许您轻松地在 `GOPATH` 外部工作。

With `GO111MODULE=on` the `buffalo` command should work as it previously did *inside* the `GOPATH`.

​	使用 `GO111MODULE=on` 时， `buffalo` 命令应像以前在 `GOPATH` 内部一样工作。

```bash
$ export GO111MODULE=on
$ buffalo new -h
```

## Working Inside the `GOPATH` 在 `GOPATH` 内部工作 

Because Go Modules are still experimental, and not complete, it is recommended to continue to work **INSIDE** the `GOPATH`. This will allow you to easily move between using, and not using modules.

​	由于 Go 模块仍处于实验阶段，尚未完成，因此建议继续在 `GOPATH` 内部工作。这将允许您轻松地在使用和不使用模块之间切换。

When working inside the `GOPATH` you should continue to use `GOPATH` style module names.

​	在 `GOPATH` 内部工作时，您应继续使用 `GOPATH` 样式的模块名称。

#### Recommended 推荐 

This style of module name works both inside, and outside, of the `GOPATH` easily. It also makes your projects work with `go get`.

​	这种样式的模块名称在 `GOPATH` 内部和外部都可以轻松使用。它还使您的项目能够与 `go get` 配合使用。

```go
module github.com/markbates/coke
```

#### Not-Recommended 不推荐 

This style of module, can work inside of the `GOPATH`, but it is less flexible, although shorter, than the longer format module name.

​	这种样式的模块可以在 `GOPATH` 内部工作，但它不如较长的格式模块名称灵活，尽管较短。

```go
module coke
```

Regardless of which module name style you pick, you **MUST** be consistent within your application.

​	无论您选择哪种模块名称样式，您都必须在您的应用程序中保持一致。

For example, if your module name is `coke` your actions package is `coke/actions`. If you module name is `github.com/markbates/coke` your actions package is `github.com/markbates/coke/actions`.

​	例如，如果您的模块名称是 `coke` ，则您的操作包是 `coke/actions` 。如果您的模块名称是 `github.com/markbates/coke` ，则您的操作包是 `github.com/markbates/coke/actions` 。

## FAQs 常见问题解答 

### I Get `invalid import` 我了解 `invalid import` 

When I run `buffalo build` I get strange errors like this when I run **outside** of my `GOPATH`:

​	当我运行 `buffalo build` 时，在我 `GOPATH` 外部运行时会收到类似这样的奇怪错误：

```text
invalid import path: "D:/projects/testBuffalo/src/my-project/actions"
```

Make sure you have `GO111MODULE=on`. If you don’t, Buffalo, tries to use your `GOPATH` to determine your package locations. Enable Go Modules support and try again.

​	确保您拥有 `GO111MODULE=on` 。如果没有，Buffalo 会尝试使用您的 `GOPATH` 来确定您的软件包位置。启用 Go 模块支持并重试。

### How Do I Migrate From Dep? 如何从 Dep 迁移？

The `go mod init` tool can read your `Gopkg.toml` files and create a new `go.mod` for you. https://github.com/golang/go/wiki/Modules

​	 `go mod init` 工具可以读取您的 `Gopkg.toml` 文件并为您创建一个新的 `go.mod` 。https://github.com/golang/go/wiki/Modules

### How Do I Use The `development` Branch? 如何使用 `development` 分支？

If you want to live on the “edge” and use the latest, bleeding edge, version of Buffalo you can tell Go Modules to get that version:

​	如果您想走在“前沿”，并使用 Buffalo 的最新、最前沿版本，您可以告诉 Go 模块获取该版本：

```bash
$ go get -u github.com/gobuffalo/buffalo@development
$ go mod tidy
```
