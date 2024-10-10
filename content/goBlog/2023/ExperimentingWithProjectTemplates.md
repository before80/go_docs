+++
title = "尝试项目模板"
date = 2023-08-21T14:59:57+08:00
weight = 91
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Experimenting with project templates - 尝试项目模板

> 原文：[https://go.dev/blog/gonew](https://go.dev/blog/gonew)
>

Cameron Balahan
31 July 2023

Cameron Balahan 2023年7月31日

When you start a new project in Go, you might begin by cloning an existing project. That way, you can start with something that already works, making incremental changes instead of starting from scratch.

​	当您在Go中启动一个新项目时，您可能会首先克隆一个现有项目。这样，您可以从已经可以运行的东西开始，逐步进行更改，而不是从头开始。

For a long time now, we have heard from Go developers that getting started is often the hardest part. New developers coming from other languages expect guidance on a default project layout, experienced developers working on teams expect consistency in their projects’ dependencies, and developers of all kinds expect an easy way to try new products and services without having to copy and paste from samples on the web.

​	很长一段时间以来，我们从Go开发人员那里听到，开始是最困难的部分。来自其他语言的新开发人员期望在默认项目布局上得到指导，有经验的开发人员在团队中工作时期望在项目的依赖关系方面保持一致，所有类型的开发人员都期望以一种简单的方式尝试新的产品和服务，而不必从网上的示例中复制粘贴。

To that end, today we published `gonew`, an experimental tool for instantiating new projects in Go from predefined templates. Anyone can write templates, which are packaged and distributed as modules, leveraging the Go module proxy and checksum database for better security and availability.

​	为此，今天我们发布了 `gonew`，这是一个用于从预定义模板中实例化新项目的实验性工具。任何人都可以编写模板，这些模板被打包并分发为模块，利用Go模块代理和校验和数据库以获得更好的安全性和可用性。

The prototype `gonew` is intentionally minimal: what we have released today is an extremely limited prototype meant to provide a base from which we can gather feedback and community direction. Try it out, [tell us what you think](https://go.dev/s/gonew-feedback), and help us build a more useful tool for everyone.

​	这个原型 `gonew` 故意设计得很简单：我们今天发布的是一个极其有限的原型，旨在为我们收集反馈和社区方向提供基础。试用一下，[告诉我们您的想法](https://go.dev/s/gonew-feedback)，并帮助我们为所有人构建一个更有用的工具。

## 开始使用 Getting started

Start by installing `gonew` using [`go install`](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies):

​	首先，使用 [`go install`](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) 安装 `gonew`：

```
$ go install golang.org/x/tools/cmd/gonew@latest
```

To copy an existing template, run `gonew` in your new project’s parent directory with two arguments: first, the path to the template you wish to copy, and second, the module name of the project you are creating. For example:

​	要复制一个现有模板，请在您的新项目的父目录中运行 `gonew`，并提供两个参数：首先，要复制的模板路径，其次是要创建的项目的模块名称。例如：

```
$ gonew golang.org/x/example/helloserver example.com/myserver
$ cd ./myserver
```

And then you can read and edit the files in `./myserver` to customize.

​	然后，您可以阅读和编辑 `./myserver` 中的文件以进行自定义。

We’ve written two templates to get you started:

​	我们编写了两个模板供您参考：

- [hello](https://pkg.go.dev/golang.org/x/example/hello): A command line tool that prints a greeting, with customization flags.
- [helloserver](https://pkg.go.dev/golang.org/x/example/helloserver): An HTTP server that serves greetings.
- [hello](https://pkg.go.dev/golang.org/x/example/hello)：一个打印问候语的命令行工具，带有自定义标志。
- [helloserver](https://pkg.go.dev/golang.org/x/example/helloserver)：提供问候的HTTP服务器。

## 编写自己的模板 Write your own templates

Writing your own template is as easy as [creating any other module](https://go.dev/doc/tutorial/create-module) in Go. Check out the examples we linked above to get started.

​	编写自己的模板就像在Go中[创建任何其他模块](https://go.dev/doc/tutorial/create-module)一样简单。查看上面我们提供的示例，以便开始。

There are also examples available from the [Google Cloud](https://github.com/GoogleCloudPlatform/go-templates) and [Service Weaver](https://github.com/ServiceWeaver/template) teams.

​	还有来自[Google Cloud](https://github.com/GoogleCloudPlatform/go-templates)和[Service Weaver](https://github.com/ServiceWeaver/template)团队的示例。

## 下一步计划 Next steps

Please try out `gonew` and let us know how we can make it better and more useful. Remember, `gonew` is just an experiment for now; we need your [feedback to get it right](https://go.dev/s/gonew-feedback).

​	请尝试使用 `gonew`，并告诉我们如何使其更好更有用。请记住，`gonew` 目前只是一个实验；我们需要您的[反馈来完善它](https://go.dev/s/gonew-feedback)。