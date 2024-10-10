+++
title = "教程：创建一个 Go 模块"
weight = 7
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false

+++
# Tutorial: Create a Go module - 教程：创建一个 Go 模块

> 原文：[https://go.dev/doc/tutorial/create-module.html](https://go.dev/doc/tutorial/create-module.html)

​	这是教程的第一部分，介绍了Go语言的一些基本特性。如果您刚开始使用Go，请务必看一下[Tutorial: Get started with Go](../TutorialGetStartedWithGo)，其中介绍了`go`命令、Go模块和非常简单的Go代码。

​	在本教程中，您将创建两个模块。第一个是一个库，旨在被其他库或应用程序导入。第二个是一个调用方应用程序，它将使用第一个模块。

​	本教程包括七个简短的主题，每个主题说明了语言的不同部分。

1. 创建一个模块 —— 编写一个小模块，其中包含可以从另一个模块调用的函数。
2. 从另一个模块调用代码 —— 导入并使用新模块。
3. 返回并处理错误 —— 添加简单的错误处理。
4. 返回一个随机的问候语 —— 以切片（Go的动态大小的数组）处理数据。
5. 返回多个人的问候语 —— 在映射中存储键/值对。
6. 添加测试 —— 使用 Go 的内置单元测试功能来测试代码。
7. 编译并安装应用程序 —— 在本地编译并安装代码。

注意：关于其他教程，请参见 [Tutorials](../Tutorials)。

## 前提条件

- 一些编程经验。这里的代码非常简单，但对函数有所了解是有帮助的。
- 编辑代码的工具。您拥有的任何文本编辑器都可以工作。大多数文本编辑器都对Go有很好的支持。最受欢迎的是`VSCode`（免费）、`GoLand`（付费）和`Vim`（免费）。
- 命令终端。在Linux和Mac上使用任何终端，以及在Windows上使用`PowerShell`或`cmd`，Go都能很好地工作。

## 启动一个别人可以使用的模块

​	首先创建一个Go模块。在模块中，您集和了一个或多个相关的包，以实现一组离散而有用的功能。例如，您可以创建一个模块，其中的包具有进行财务分析的功能，这样其他编写财务应用程序的人就可以使用您的工作。关于开发模块的更多信息，请参见[开发和发布模块](../../UsingAndUnderstandingGo/DevelopingModules/DevelopingAndPublishingModules)。

​	Go代码被分组进包，而包又被分组进模块。模块指定了运行您的代码所需的依赖项，包括Go版本和它所需的其他模块集。

​	在模块中增加或改进功能时，您会发布模块的新版本。开发人员在编写调用您的模块中的函数的代码时，可以导入模块的更新包，并在投入生产使用之前用新版本进行测试。

a. 打开一个命令提示符，`cd`到您的主目录。

在Linux或Mac上：

```shell
cd
```

在Windows上：

```shell
cd %HOMEPATH%
```

b. 为您的Go模块源代码创建一个`greetings`目录。

例如，从您的主目录使用以下命令：

```shell
mkdir greetings
cd greetings
```

c. 使用`go mod init`命令启动您的模块。

​	运行`go mod init`命令，给它您的模块路径 —— 这里用`example.com/greetings`。如果您发布了一个模块，这必须是一个Go工具可以下载您模块的路径。这将是您的代码库。

更多关于用模块路径命名模块的信息，请看[管理依赖项](../../UsingAndUnderstandingGo/ManagingDependencies)。

```shell
$ go mod init example.com/greetings
go: creating new go.mod: module example.com/greetings
```

`go mod init`命令创建了一个`go.mod`文件来跟踪您的代码的依赖项。到目前为止，该文件只包括您的模块名称和您的代码支持的Go版本。但当您添加依赖项时，`go.mod`文件将列出您的代码所依赖的版本。这可以保持构建的可重复性，并让您直接控制使用哪些模块的版本。

d. 在您的文本编辑器中，创建一个文件来编写您的代码，并将其称为`greetings.go`。

e. 把下面的代码粘贴到您的`greetings.go`文件中并保存该文件。

```go title="greeting.go" linenums="1"
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

​	这是您的模块的第一段代码。它向任何要求问候的调用方返回问候语。您将在下一步编写调用该函数的代码。

在这段代码中，您：

- 声明一个`greetings`包，以收集相关的函数。

- 实现一个`Hello`函数来返回问候语。

  ​	这个函数接收一个`name`参数，其类型为字符串。该函数还返回一个字符串。在Go中，一个名字以大写字母开头的函数可以被不在同一个包中的函数调用。这在 Go 中被称为`导出名称`。关于导出名称的更多信息，请参见Go之旅中的[导出名称](../../GoTour/Basics/PackagesVariablesAndFunctions#exported-names-导出名)。

  ![img](TutorialCreateAGoModule_img/function-syntax.png)
  
- 声明一个`message`变量来保存您的问候语。

  ​	在Go中，`:=`操作符是在一行中`声明和初始化`一个变量的快捷方式（Go使用右边的值来确定变量的类型）。如果从长计议，您可以这样写：

  ```go linenums="1"
  var message string
  message = fmt.Sprintf("Hi, %v. Welcome!", name)
  ```
  
- 使用`fmt`包的[`Sprintf`函数]({{< ref "/stdLib/fmt#func-sprintf">}})来创建一个问候信息。第一个参数是一个格式字符串，`Sprintf`将`name`参数的值替换`%v`格式动词。插入`name`参数的值就完成了问候语的文本。

- 将格式化的问候语文本返回给调用方。


在下一步，您将从另一个模块调用这个函数。