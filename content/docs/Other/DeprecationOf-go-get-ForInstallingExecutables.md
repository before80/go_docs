+++
title = "废弃用于安装可执行文件的 go get"
weight = 1
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Deprecation of 'go get' for installing executable - 废弃用于安装可执行文件的 "go get"

> 原文：[https://go.dev/doc/go-get-install-deprecation](https://go.dev/doc/go-get-install-deprecation)

## Overview 概述

​	从 Go 1.17 开始，不推荐使用 `go get` 安装可执行文件。可以使用 `go install` 来代替。

​	在 Go 1.18 中，`go get` 将不再构建包；它将只用于添加、更新或删除 `go.mod` 中的依赖项。具体来说，`go get`将始终像启用`-d`标志那样行事。

## What to use instead 用什么来代替

​	要在当前模块的上下文中安装一个可执行文件，请使用`go install`，不带使用版本后缀，如下所示。这将应用当前目录或父目录中`go.mod`文件的版本要求和其他指令。

```
go install example.com/cmd
```

​	要在忽略当前模块的情况下安装一个可执行文件，请使用带有[版本后缀](../../References/GoModulesReference/Module-awareCommands#version-queries)的`go install`，例如`@v1.2.3`或`@latest`，如下所示。当使用版本后缀时，`go install` 不会读取或更新当前目录或父目录下的 `go.mod` 文件。

```
# 安装特定版本
go install example.com/cmd@v1.2.3

# 安装可用的最高版本
go install example.com/cmd@latest
```

​	为了避免歧义，当使用带有版本后缀的`go install`时，所有参数必须引用同一模块中同一版本的`main`包。如果该模块有一个`go.mod`文件，它不能包含像`replace`或`exclude`这样的指令，如果它是主模块，会导致它被不同的解释。该模块的`vendor`目录不被使用。

​	有关详细信息，请参阅[go install](../../References/GoModulesReference/Module-awareCommands#go-install)。

## Why this is happening 为什么会出现这种情况

​	自从引入模块以来，`go get`命令既被用来更新`go.mod`中的依赖项，也被用来安装命令。这种组合经常造成混乱和不便：在大多数情况下，开发人员希望更新依赖项或安装命令，但不能同时更新两者。

​	从Go 1.16开始，`go install`可以按照命令行指定的版本安装命令，同时忽略当前目录下的`go.mod`文件（如果存在的话）。现在大多数情况下都应该使用`go install`来安装命令。

​	`go get` 构建和安装命令的功能现在已被弃用，因为这个功能对于`go install`来说是多余的。删除这一功能将使 `go get` 更加快速，因为它默认不会编译或链接包。当更新一个无法为当前平台构建的包时，`go get` 也不会报告错误。

​	完整的讨论见提议[#40276](https://go.dev/issue/40276)。