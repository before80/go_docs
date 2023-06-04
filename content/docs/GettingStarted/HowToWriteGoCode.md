+++
title = "如何编写 go 代码"
weight = 14
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# How to Write Go Code - 如何编写 Go 代码

> 原文：[https://go.dev/doc/code](https://go.dev/doc/code)

## 简介

​	本文档演示了在模块中开发一个简单的Go包，并介绍了[go tool](../../References/CommandDocumentation/go)，这是获取、构建和安装Go模块、包和命令的标准方式。

!!! warning "注意"

	注意：本文档假设您使用的是`Go 1.13`或更高版本，并且没有设置`GO111MODULE`环境变量。如果你要找的是本文档的旧版，即模块之前的版本，它被存档在[这里](../HowToWriteGoCodeWithGOPATH)。

## 代码组织

​	Go程序被组织成包。包是同一目录下的源文件的集合，这些文件被编译在一起。在一个源文件中定义的函数、类型、变量和常量对同一包内的所有其他源文件都是可见的。

​	代码库包含一个或多个模块。模块是相关 Go 包的集合，它们被一起发布。一个 Go 代码库通常只包含一个模块，位于代码库的根部。那里有一个名为`go.mod`的文件声明了模块路径：模块内所有包的导入路径前缀。该模块包含其`go.mod`文件的所在目录中的软件包，以及该目录的子目录（的软件包），直到包含另一个`go.mod`文件的所在的子目录（的软件包）（如果有的话）。

​	请注意，你不需要在构建之前将你的代码发布到一个远程仓库。一个模块可以在本地定义而不属于一个版本库。然而，组织你的代码是一个好习惯，就像你有一天会发布它一样。

​	每个模块的路径不仅作为其包的导入路径前缀，而且还指出`go`命令应该在哪里下载它。例如，为了下载`golang.org/x/tools`模块，go命令会查阅[https://golang.org/x/tools](https://golang.org/x/tools)（[这里](../../References/CommandDocumentation/go#remote-import-paths)有更多描述）所指示的仓库。

​	导入路径是一个用于导入软件包的字符串。一个包的导入路径是它的模块路径与模块中的子目录相连接。例如，模块`github.com/google/go-cmp`在`cmp/`目录下包含一个包。该包的导入路径是`github.com/google/go-cmp/cmp`。**标准库中的包没有模块路径前缀**。

## 你的第一个程序

​	要编译和运行一个简单的程序，首先选择一个模块路径（我们将使用 `example/user/hello`），并创建一个声明它的 `go.mod` 文件：

```shell
$ mkdir hello # Alternatively, clone it if it already exists in version control.
$ cd hello
$ go mod init example/user/hello
go: creating new go.mod: module example/user/hello
$ cat go.mod
module example/user/hello

go 1.16
$
```

​	Go源文件中的第一条语句必须是`package name`。可执行命令必须始终使用`package main`。

​	接下来，在该目录下创建一个名为`hello.go`的文件，包含以下Go代码：

```go linenums="1" title="hello.go"
package main

import "fmt"

func main() {
    fmt.Println("Hello, world.")
}
```

现在你可以用`go`工具构建和安装该程序：

```shell
$ go install example/user/hello
$
```

​	这个命令构建了`hello`命令，产生了一个可执行的二进制文件。然后，它将该二进制文件安装为`$HOME/go/bin/hello`（或者，在Windows下，`%USERPROFILE%\gobinhello.exe`）。

​	安装目录是由`GOPATH`和`GOBIN`[环境变量](../../References/CommandDocumentation/go#environment-variables)控制的。如果设置了`GOBIN`，二进制文件将被安装到该目录。如果设置了`GOPATH`，二进制文件将被安装到`GOPATH`列表中第一个目录的`bin`子目录中。否则，二进制文件将被安装到默认的 `GOPATH`（`$HOME/go` 或 `%USERPROFILE%\go`）的 `bin` 子目录。

​	你可以使用`go env`命令为未来的`go`命令可移植地设置环境变量的默认值：

```shell
$ go env -w GOBIN=/somewhere/else/bin
$
```

​	要取消先前由`go env -w`设置的变量，请使用`go env -u`。

```shell
$ go env -u GOBIN
$
```

​	像`go install`这样的命令在包含当前工作目录的模块上下文中应用。如果工作目录不在`example/user/hello`模块内，`go install`可能会失败。

​	为了方便起见，`go`命令接受相对于工作目录的路径，如果没有给出其他路径，则默认为当前工作目录下的软件包。因此在我们的工作目录中，以下命令都是等价的。

```shell
$ go install example/user/hello
$ go install .
$ go install
```

​	接下来，让我们运行该程序，以确保它能工作。为了方便起见，我们将安装目录添加到我们的`PATH`中，以使运行二进制文件更加容易：

```shell
# Windows users should consult https://github.com/golang/go/wiki/SettingGOPATH
# for setting %PATH%.
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ hello
Hello, world.
$
```

​	如果你使用的是源码控制系统，现在是初始化代码库的好时机，添加文件，并提交你的第一个改动。同样，这一步是可选的：你不需要使用源码控制来编写Go代码。

```shell
$ git init
Initialized empty Git repository in /home/user/hello/.git/
$ git add go.mod hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 7 insertion(+)
 create mode 100644 go.mod hello.go
$
```

`go` 命令通过请求相应的 HTTPS URL 并读取嵌入在 HTML 响应中的元数据来定位包含给定模块路径的代码库（见 `go help importpath`）。许多托管服务已经为包含Go代码库提供了元数据，所以让你的模块供他人使用的最简单方法通常是使其模块路径与代码库的URL相匹配。

### 在你的模块中导入包

​	让我们写一个`morestrings`包并从`hello`程序中使用它。首先，为该包创建一个名为`$HOME/hello/morestrings`的目录，然后在该目录下创建一个名为`reverse.go`的文件，内容如下：

```go linenums="1" title="reverse.go"
// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

​	因为我们的`ReverseRunes`函数以大写字母开头，所以它是[可导出](../../References/LanguageSpecification/DeclarationsAndScope#exported-identifiers)的，可以在其他包中通过导入`morestrings`包来使用该函数。

让我们用`go build`来测试一下这个包的编译情况：

```shell
$ cd $HOME/hello/morestrings
$ go build
$
```

​	这不会产生一个输出文件。相反，它会把编译好的包保存在本地的构建缓存中。

​	在确认了`morestrings`包的构建之后，让我们在`hello`程序中使用它。要做到这一点，请修改你原来的`$HOME/hello/hello.go`以使用`morestrings`包：

```go linenums="1"
package main

import (
    "fmt"

    "example/user/hello/morestrings"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
```

安装`hello`程序：

```shell
$ go install example/user/hello
```

运行新版本的程序，你应该看到一个新的、反转的信息：

```shell
$ hello
Hello, Go!
```

### 从远程模块导入包

​	导入路径可以描述如何使用`Git`或`Mercurial`等修订控制系统获得软件包的源代码。`go`工具使用这个属性来自动从远程代码库获取包。例如，要在你的程序中使用`github.com/google/go-cmp/cmp`：

```go linenums="1"
package main

import (
    "fmt"

    "example/user/hello/morestrings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
```

​	现在你有了对外部模块的依赖，你需要`下载该模块`并在你的`go.mod`文件中记录其版本。`go mod tidy`命令为导入的软件包添加缺少的模块需求，并删除不再使用的模块需求。

```shell
$ go mod tidy
go: finding module for package github.com/google/go-cmp/cmp
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.4
$ go install example/user/hello
$ hello
Hello, Go!
  string(
-     "Hello World",
+     "Hello Go",
  )
$ cat go.mod
module example/user/hello

go 1.16

require github.com/google/go-cmp v0.5.4
$
```

​	模块的依赖项被自动下载到`GOPATH`环境变量所指示的目录的`pkg/mod`子目录中。一个特定版本的模块的下载内容在所有需要（`require`）该版本的其他模块中共享，因此`go`命令将这些文件和目录标记为只读。要删除`所有下载的模块`，你可以在`go clean`中传递`-modcache`标志：

```shell
$ go clean -modcache
$
```

## 测试

​	Go有一个由`go test`命令和`testing`包组成的轻量级测试框架。

​	你可以通过创建一个名称以`_test.go`结尾的文件来编写测试，该文件包含名为`TestXXX`的函数，其签名为`func（t *testing.T）`。测试框架运行每个这样的函数；如果在该函数调用一个失败的函数，如`t.Error`或`t.Fail`，则认为测试失败。

​	通过创建包含以下Go代码的`$HOME/hello/morestrings/reverse_test.go`文件，向`morestrings`包添加一个测试。

```go linenums="1" title="reverse_test.go"
package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := ReverseRunes(c.in)
        if got != c.want {
            t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
```

然后用`go test`运行该测试：

```shell
$ cd $HOME/hello/morestrings
$ go test
PASS
ok  	example/user/hello/morestrings 0.165s
$
```

运行`go help test`，更多细节请看[testing package documentation](https://go.dev/pkg/testing/)。

## 下一步

​	订阅 [golang-announce](https://groups.google.com/group/golang-announce) 邮件列表，以便在 Go 的新稳定版本发布时获得通知。

​	请参阅 [Effective Go](../../UsingAndUnderstandingGo/EffectiveGo) 以了解编写清晰、简洁的 Go 代码的技巧。

​	参加 [A Tour of Go](../../GoTour/UsingTheTour/Welcome)来学习这门语言。

​	访问[文档页面]()，了解有关 Go 语言及其库和工具的一系列深度文章。

## 获得帮助

​	要获得实时帮助，请向社区管理的 [gophers Slack server](https://gophers.slack.com/messages/general/) （[在此](https://invite.slack.golangbridge.org/)获取邀请）中愿意帮忙的 gophers 咨询。

​	用于讨论 Go 语言的官方邮件列表是 [Go Nuts](https://groups.google.com/group/golang-nuts)。

​	使用 [Go issue tracker](https://go.dev/issue)报告 bugs。