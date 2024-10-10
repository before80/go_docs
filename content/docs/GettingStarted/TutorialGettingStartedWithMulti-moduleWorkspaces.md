+++
title = "教程：开始使用多模块工作区"
weight = 8
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Tutorial: Getting started with multi-module workspaces - 教程：开始使用多模块工作区

> 原文：[https://go.dev/doc/tutorial/workspaces.html](https://go.dev/doc/tutorial/workspaces.html)

​	本教程介绍了Go中多模块工作区的基本知识。通过多模块工作区，您可以告诉Go命令您同时在多个模块中编写代码，并在这些模块中轻松构建和运行代码。

​	在本教程中，您将在一个共享的多模块工作区中创建两个模块，对这些模块进行修改，并在构建中看到这些修改的结果。

注意：关于其他教程，请看[Tutorials](../Tutorials)。

## 前提条件

- 安装 `Go 1.18` 或更高版本。
- 编辑代码的工具。任何文本编辑器都可以使用。
- 命令终端。Go在Linux和Mac上使用任何终端，在Windows上使用`PowerShell`或`cmd`，都能很好地工作。

​	本教程需要`go1.18`或更高版本。请确保您已经使用 [go.dev/dl](https://go.dev/dl) 的链接安装了`Go 1.18`或更高版本。

## 为您的代码创建一个模块

首先，为您将要写的代码创建一个模块。

a. 打开一个命令提示符，换到您的主目录。

在Linux或Mac上：

```shell
$ cd
```

在Windows上：

```shell
C:\> cd %HOMEPATH%
```

本教程的其余部分将显示一个`$`作为提示符。您使用的命令在Windows上也可以使用。

b. 在命令提示符下，为您的代码创建一个名为`workspace`的目录。

```shell
$ mkdir workspace
$ cd workspace
```

c. 初始化模块。

​	我们的示例将创建一个新的模块 `hello`，它将依赖于 `golang.org/x/example` 模块。

创建`hello`模块：

```shell
$ mkdir hello
$ cd hello
$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello
```

通过使用`go get`添加对`golang.org/x/example`模块的依赖。

```shell
$ go get golang.org/x/example
```

在`hello`目录下创建`hello.go`，内容如下：

```go title="hello.go" linenums="1"
package main

import (
    "fmt"

    "golang.org/x/example/stringutil"
)

func main() {
    fmt.Println(stringutil.Reverse("Hello"))
}
```

现在，运行`hello`程序：

```shell
$ go run example.com/hello
olleH
```

## 创建工作区

在这一步骤中，我们将创建一个`go.work`文件来指定模块的工作区。

#### 初始化工作区

在`workspace`目录下，运行：

```shell
$ go work init ./hello
```

​	`go work init`命令告诉`go`为包含`./hello`目录中的模块的工作区创建一个`go.work`文件。

`go`命令生成的`go.work`文件看起来像这样：

```
go 1.18

use ./hello
```

`go.work`文件的语法与`go.mod`相似。

​	`go`指令告诉Go应该用哪个版本的Go来解释该文件。它与`go.mod`文件中的`go`指令类似。

​	`use`指令告诉Go，在进行构建时，`hello`目录下的模块应该是`主模块`。

因此，在`workspace`的任何子目录中，模块都会被激活。

#### 在`workspace` 目录下运行程序

在`workspace`目录下，运行：

```shell
$ go run example.com/hello
olleH
```

The Go command includes all the modules in the workspace as main modules. This allows us to refer to a package in the module, even outside the module. Running the `go run` command outside the module or the workspace would result in an error because the `go` command wouldn’t know which modules to use.

​	Go命令将`workspace` 中的所有模块都作为`主模块`。这使得我们可以引用模块中的包，甚至在模块之外。在模块或`workspace`之外运行`go run`命令会导致错误，因为`go`命令不知道要使用哪些模块。=> 仍有疑问？？

​	接下来，我们将在`workspace`中添加一份 `golang.org/x/example` 模块的本地副本。然后我们将在`stringutil`包中添加一个新的函数，我们可以用它代替`Reverse`。

## 下载并修改`golang.org/x/example`模块

​	在这一步骤中，我们将下载一份包含 `golang.org/x/example` 模块的 Git repo 的副本，将其添加到`workspace`，然后为其添加一个新函数，我们将在 `hello` 程序中使用它。

a. 克隆仓库。

在 `workspace` 目录下，运行`git`命令来克隆版本库：

```shell
$ git clone https://go.googlesource.com/example
Cloning into 'example'...
remote: Total 165 (delta 27), reused 165 (delta 27)
Receiving objects: 100% (165/165), 434.18 KiB | 1022.00 KiB/s, done.
Resolving deltas: 100% (27/27), done.
```

b. 将该模块添加到`workspace`

```shell
$ go work use ./example
```

​	`go work use`命令在`go.work`文件中添加了一个新模块。现在它看起来像这样：

```
go 1.18

use (
    ./hello
    ./example
)
```

该模块现在同时包括`example.com/hello`模块和`golang.org/x/example`模块。

​	这将使我们能够使用我们将在`stringutil`模块副本中编写的新代码，而不是我们用`go get`命令下载的模块缓存中的模块版本。

c. 添加新函数。

​	我们将在 `golang.org/x/example/stringutil` 包中添加一个新函数来对字符串进行大写。

​	在`workspace/example/stringutil`目录下创建一个名为`toupper.go`的新文件，包含以下内容：

```go title="toupper.go" linenums="1"
package stringutil

import "unicode"

// ToUpper uppercases all the runes in its argument string.
func ToUpper(s string) string {
    r := []rune(s)
    for i := range r {
        r[i] = unicode.ToUpper(r[i])
    }
    return string(r)
}
```

d. 修改`hello`程序以使用该函数。

修改`workspace/hello/hello.go`的内容，使其包含以下内容：

```go title="hello.go" linenums="1"
package main

import (
    "fmt"

    "golang.org/x/example/stringutil"
)

func main() {
    fmt.Println(stringutil.ToUpper("Hello"))
}
```

#### 在`workspace` 运行代码

从 `workspace` 目录中，运行

```shell
$ go run example.com/hello
HELLO
```

​	Go命令在`go.work`文件指定的`hello`目录下找到了命令行中指定的`example.com/hello`模块，并同样使用`go.work`文件解析了`golang.org/x/example`的导入。

​	`go.work`可以用来代替跨多模块工作时使用的 [replace](../../References/GoModulesReference/gomodFiles#replace-directive-replace) 指令。

​	由于这两个模块在同一个`workspace` 中，所以很容易在一个模块中做出改变，并在另一个模块中使用它。

#### 未来的步骤

​	现在，为了正确地发布这些模块，我们需要对 `golang.org/x/example` 模块进行发布，例如在 `v0.1.0`。这通常是通过在模块的版本控制库中标记一个提交来完成的。更多细节请参见[模块发布工作流程文档](../../UsingAndUnderstandingGo/DevelopingModules/ModuleReleaseAndVersioningWorkflow)。发布完成后，我们可以在 `hello/go.mod` 中增加对 `golang.org/x/example` 模块的要求：

```shell
cd hello
go get golang.org/x/example@v0.1.0
```

这样一来，`go`命令就可以正确解决`workspace`外的模块了。

## 了解更多关于工作区的信息

​	除了我们在教程前面看到的`go work init`外，`go`命令还有几个子命令用于处理工作区：

- `go work use [-r] [dir]` 在`go.work`文件中为`dir`添加一个`use`指令（如果该`dir`存在的话），如果参数目录不存在，则删除使用目录。`-r`标志会递归地检查`dir`的子目录。

- `go work edit` 编辑 `go.work`文件，与`go mod edit`类似。

- `go work sync` 将工作区构建列表中的依赖项同步到每个工作区模块中。

关于 workspaces 和`go.work`文件的更多细节，请参见Go模块参考中的[Workspaces](../../References/GoModulesReference/Workspaces)。