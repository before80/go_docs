+++
title = "教程：开始使用go"
weight = 6
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Tutorial: Get started with Go - 教程：开始使用Go

> 原文：[https://go.dev/doc/tutorial/getting-started](https://go.dev/doc/tutorial/getting-started)

​	在本教程中，您将获得Go编程的简要介绍。在这一过程中，您将：

- 安装Go（如果您还没有安装）。
- 编写一些简单的 "Hello, world "代码。
- 使用`go`命令来运行您的代码。
- 使用`Go包的发现工具`来寻找可以在自己的代码中使用的包。
- 调用外部模块的函数。

> 注意：关于其他教程，请参见[Tutorials](../Tutorials)。

## 前提条件

- 一些编程经验。这里的代码非常简单，但对函数有所了解是有帮助的。
- 一个编辑代码的工具。您拥有的任何文本编辑器都可以工作。大多数文本编辑器都对Go有很好的支持。最受欢迎的是`VSCode`（免费）、`GoLand`（付费）和`Vim`（免费）。
- 一个命令终端。在Linux和Mac上使用任何终端，以及在Windows上使用`PowerShell`或`cmd`，Go都能很好地工作。

## 安装Go

只需使用[下载和安装](../InstallingGo)页面的步骤。

## 编写一些代码

从Hello, World开始。

a. 打开一个命令提示符，`cd`到您的主目录。

在Linux或Mac上：

```shell
cd
```

在Windows上：

```shell
cd %HOMEPATH%
```

b. 为您的第一个Go源代码创建一个Hello目录。

例如，使用以下命令：

```shell
mkdir hello
cd hello
```

c. 为您的代码启用依赖项跟踪。

当您的代码导入其他模块中的包时，可以通过代码自己的模块管理这些依赖项。该模块由一个`go.mod`文件定义，该文件跟踪提供这些包的模块。这个`go.mod`文件与您的代码在一起，包括在您的源代码库里。

要通过创建`go.mod`文件来启用您的代码的依赖项跟踪，请运行[`go mod init`命令](../../References/GoModulesReference/Module-awareCommands/#go-mod-init)，并将代码所在的模块名称指定给它。这个名称就是该模块的模块路径。

在实际开发中，模块路径通常是保存您源代码的库位置。例如，模块路径可能是`github.com/mymodule`。如果您打算发布您的模块供他人使用，模块路径必须是Go工具可以下载您的模块的位置。关于用模块路径命名模块的更多信息，请参见[管理依赖项](../../UsingAndUnderstandingGo/ManagingDependencies)。

在本教程中，只需使用 `example/hello`。

```shell
$ go mod init example/hello
go: creating new go.mod: module example/hello
```

d. 在您的文本编辑器中，创建一个`hello.go`文件，在其中编写您的代码。

e. 将以下代码粘贴到`hello.go`文件中，并保存该文件。

```go title="hello.go" linenums="1"
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

这是您的Go代码。在这段代码中，您：

- 声明一个`main`包（包是对函数进行分组的一种方式，它是由同一目录下的所有文件组成的）。
- 导入流行的[`fmt`包]({{< ref "/stdLib/fmt">}})，其中包含用于格式化文本的函数，包括打印到控制台。这个包是安装Go时获得的[标准库]({{< ref "/stdLib" >}})包之一。
- 实现一个`main`函数，将信息打印到控制台。当您运行`main`包时，默认执行一个`main`函数。

f. 运行代码，查看问候语。

```shell
$ go run .
Hello, World!
```

[go run 命令]({{< ref "/cmd/go#编译并运行Go程序">}})是您用Go完成任务的众多`go`命令之一。使用下面的命令来获得其他命令的列表：

```shell
$ go help
```

## 调用外部包中的代码

​	当您需要您的代码做一些可能已经被别人实现的事情时，可以查找包含可以在代码中使用的函数的包。

a. 用一个外部模块的函数使您打印的信息更有趣一些。

1. 访问pkg.go.dev，[搜索 "quote"包](https://pkg.go.dev/search?q=quote)。
2. 在搜索结果中找到并点击`rsc.io/quote`包（如果您看到`rsc.io/quote/v3`，暂时忽略它）。
3. 在`Documentation`部分的 `Index` 下，请注意这些是可以在您的代码中调用的函数列表。您将使用到Go函数。
4. 在这个页面的顶部，请注意`quote`包包含在`rsc.io/quote`模块中。

​	您可以使用 `pkg.go.dev` 网站来查找已发布的模块，这些模块的包中有您可以在自己的代码中使用的函数。软件包以模块的形式发布 ——如`rsc.io/quote` —— 其他人可以使用它们。模块随着时间的推移会有新版本的改进，您可以升级您的代码以使用改进后的版本。

b. 在您的 Go 代码中，导入 `rsc.io/quote` 包并添加对其 Go 函数的调用。

在添加了高亮的行之后，您的代码应该包括以下内容：

```go hl_lines="5 5" title="hello.go" linenums="1"
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

c. 添加新的模块requirement 和 sums。

​	Go将添加`quote`模块作为需求，以及一个用于验证模块的`go.sum`文件。更多信息请参见 `Go模块参考` 中的[认证模块](../../References/GoModulesReference/AuthenticatingModules)。

```shell
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
```

d. 运行您的代码，看看您调用的函数所产生的信息。

```shell
$ go run .
Don't communicate by sharing memory, share memory by communicating.
```

注意您的代码调用了Go函数，打印了一条关于通信的巧妙信息。

当您运行`go mod tidy`这条命令时，它`定位`并`下载`了包含您导入的包的`rsc.io/quote`模块。默认情况下，它下载的是最新版本。

## 编写更多的代码

​	通过这个快速介绍，您已经安装了Go，并学会了一些基本知识。要用另一个教程写更多的代码，请看[创建Go模块](../TutorialCreateAGoModule)。