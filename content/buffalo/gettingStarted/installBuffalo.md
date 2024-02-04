+++
title = "安装 Buffalo"
date = 2024-02-04T21:05:33+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/getting_started/installation/]({{< ref "/buffalo/gettingStarted/installBuffalo" >}})

# Install Buffalo  

In this chapter, you’ll learn how to install Buffalo, either from pre-built binaries or from source.

​	在本章中，您将学习如何安装 Buffalo，无论是从预构建的二进制文件还是从源代码安装。

Buffalo provides **two major components**:

​	Buffalo 提供两个主要组件：

- The `buffalo` tool, a powerful toolbox to help you develop in a fast and efficient way.
  `buffalo` 工具，一个强大的工具箱，可以帮助您快速高效地进行开发。
- The buffalo framework, a collection of pieces to construct your app.
  buffalo 框架，用于构建应用程序的组件集合。

Buffalo is currently available and tested on the following platforms:

​	Buffalo 目前可在以下平台上使用并经过测试：

- GNU/Linux
- Mac OSX
- Windows

## Requirements 要求 

Before installing make sure you have the required dependencies installed:

​	在安装之前，请确保您已安装所需的依赖项：

- [A working Go environment
  一个可用的 Go 环境](http://gopherguides.com/before-you-come-to-class)
- [Go](https://golang.org/) version `v1.16.0`.
  Go 版本 `v1.16.0` 。

##### Frontend Requirements 前端要求 

The following requirements are optional. You don’t need them if you want to build an API or if you prefer to build your app in an old-fashioned way.

​	以下要求是可选的。如果您想构建 API 或更喜欢以老式方式构建应用，则不需要它们。

- [node](https://github.com/nodejs/node) version `8` or greater
  node 版本 `8` 或更高版本
- either [yarn](https://yarnpkg.com/en/) or [npm](https://github.com/npm/npm) for the [asset pipeline]({{< ref "/buffalo/frontend/assets" >}}) built upon [webpack](https://github.com/webpack/webpack).
  yarn 或 npm 用于基于 webpack 构建的资产管道。

##### Database Specific Requirements 数据库特定要求 

Again, if you don’t need a database, you won’t need these.

​	同样，如果您不需要数据库，则不需要这些。

- **SQLite 3**: GCC, or equivalent C compiler for [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).
  SQLite 3：GCC 或 mattn/go-sqlite3 的等效 C 编译器。

## Installation from a Release Archive - 64 bits 从发行存档安装 - 64 位 

The release packages contain Buffalo without SQLite support.
发行包包含不带 SQLite 支持的 Buffalo。

Since `v0.10.3`, pre-compiled archives are provided with each release. If you don’t need the latest cutting-edge version, you’ll probably prefer to install this version.

​	自 `v0.10.3` 起，每个发行版都提供预编译存档。如果您不需要最新的前沿版本，您可能更喜欢安装此版本。

### GNU / Linux

```sh
$ wget https://github.com/gobuffalo/cli/releases/download/v0.18.14/buffalo_0.18.14_Linux_x86_64.tar.gz
$ tar -xvzf buffalo_0.18.14_Linux_x86_64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
```

### MacOS

```sh
$ curl -OL https://github.com/gobuffalo/cli/releases/download/v0.18.14/buffalo_0.18.14_Darwin_x86_64.tar.gz
$ tar -xvzf buffalo_0.18.14_Darwin_x86_64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
# or if you have ~/bin folder setup in the environment PATH variable
$ mv buffalo ~/bin/buffalo
```

## Scoop (Windows) Scoop（Windows）

Buffalo can be installed using the [Scoop](http://scoop.sh/) package manager:

​	可以使用 Scoop 包管理器安装 Buffalo：

```powershell
PS C:\> scoop install buffalo
```

## Chocolatey (Windows) Chocolatey（Windows）

Buffalo can be installed using the [Chocolatey](https://chocolatey.org/packages/buffalo) package manager. Versions on Chocolatey are published with a potential delay and must go through moderation before they are available:

​	可以使用 Chocolatey 包管理器安装 Buffalo。Chocolatey 上的版本可能会延迟发布，并且必须经过审核才能使用：

```powershell
PS C:\> choco install buffalo
```

## Homebrew (macOS) Homebrew（macOS）

On macOS, you can also install Buffalo with [Homebrew](https://brew.sh/). After you have Homebrew [installed](https://docs.brew.sh/Installation), you can easily install Buffalo:

​	在 macOS 上，您还可以使用 Homebrew 安装 Buffalo。安装 Homebrew 后，您可以轻松安装 Buffalo：

```sh
brew install gobuffalo/tap/buffalo
```

## GoFish (Cross-Platforms) GoFish（跨平台）

[GoFish](https://gofi.sh/) is a cross-platform systems package manager, that works across Windows, MacOSX and Linux.

​	GoFish 是一个跨平台系统包管理器，适用于 Windows、MacOSX 和 Linux。

After you have GoFish [installed](https://gofi.sh/#install), you can very simply install Buffalo:

​	安装 GoFish 后，您可以非常简单地安装 Buffalo：

```sh
$ gofish install buffalo
==> Installing buffalo...
🐠  buffalo v0.18.14: installed in 3.223672926s
```

## Custom Installation **with** SQLite3 Support 自定义安装，支持 SQLite3

**SQLite 3** requires a GCC, or equivalent C compiler for [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) to compile. You **must** have a GCC installed **first** before installing Buffalo.

​	SQLite 3 需要一个 GCC 或同等的 C 编译器才能编译 mattn/go-sqlite3。在安装 Buffalo 之前，您必须先安装一个 GCC。

```sh
$ go install -tags sqlite github.com/gobuffalo/cli/cmd/buffalo@v0.18.14
```

**Windows Users**: Follow the installation guide at https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3 to install a GCC for Windows 10. Alternatively, GCC can be installed with the [Scoop](http://scoop.sh/) package manager:

​	Windows 用户：按照 https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3 上的安装指南在 Windows 10 上安装 GCC。或者，可以使用 Scoop 包管理器安装 GCC：

```powershell
PS C:\> scoop install gcc
```

These instructions can also be used for upgrading to a newer version of Buffalo.
这些说明也可用于升级到 Buffalo 的较新版本。

## Custom Installation **without** SQLite3 Support 自定义安装，不支持 SQLite3

```sh
$ go install github.com/gobuffalo/cli/cmd/buffalo@v0.18.14
```

These instructions can also be used for upgrading to a newer version of Buffalo.
这些说明也可以用于升级到 Buffalo 的较新版本。

## Verify Your Installation 验证您的安装 

You can check if your installation is working, by executing the `buffalo` command in a terminal/command prompt:

​	您可以通过在终端/命令提示符中执行 `buffalo` 命令来检查您的安装是否正常工作：

```sh
$ buffalo
Build Buffalo applications with ease

Usage:
  buffalo [command]

Available Commands:
  build       Build the application binary, including bundling of webpack assets
  completion  Generate the autocompletion script for the specified shell
  db          [PLUGIN] [DEPRECATED] please use `buffalo pop` instead.
  destroy     Destroy generated components
  dev         Run the Buffalo app in 'development' mode
  fix         Attempt to fix a Buffalo applications API to match version v0.18.6
  generate    Generate application components
  help        Help about any command
  info        Print diagnostic information (useful for debugging)
  new         Creates a new Buffalo application
  plugins     tools for working with buffalo plugins
  pop         [PLUGIN] A tasty treat for all your database needs
  routes      Print all defined routes
  setup       Setup a newly created, or recently checked out application.
  task        Run grift tasks
  test        Run the tests for the Buffalo app. Use --force-migrations to skip schema load.
  version     Print the version information

Flags:
  -h, --help   help for buffalo

Use "buffalo [command] --help" for more information about a command.
```

If you have a similar output, your Buffalo toolbox is ready to work!

​	如果您有类似的输出，那么您的 Buffalo 工具箱就可以工作了！

## Next Steps 后续步骤 

- [Generate a New Project]({{< ref "/buffalo/gettingStarted/generatingANewProject" >}}) - Create your first Buffalo project!
  生成一个新项目 - 创建您的第一个 Buffalo 项目！
