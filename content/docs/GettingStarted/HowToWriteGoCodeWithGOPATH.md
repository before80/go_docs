+++
title = "如何编写 Go 代码（使用GOPATH）"
weight = 15
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false

+++
# How to Write Go Code (with GOPATH) - 如何编写 Go 代码（使用GOPATH）

> 原文：[https://go.dev/doc/gopath_code](https://go.dev/doc/gopath_code)

## 简介

​	如果您是Go的新手，请看最近的[How to Write Go Code](../HowToWriteGoCode)。

​	本文演示了一个简单的 Go 包的开发，并介绍了 [go tool]({{< ref "/cmd/go">}})，这是获取、构建和安装 Go 包和命令的标准方法。

​	`go tool`要求您以特定的方式组织您的代码。请仔细阅读本文档。它解释了最简单的方法来启动和运行您的Go安装。

​	类似的解释可以以[screencast](https://www.youtube.com/watch?v=XCsL89YtqCs)的形式提供。

## 代码组织

### 概述

- Go 程序员通常将所有 Go 代码放在一个工作区中。
- 一个工作区包含许多版本控制代码库（例如由 Git 管理）。
- 每个仓库包含一个或多个包。
- 每个包由单个目录下的一个或多个 Go 源文件组成。
- 包的目录路径决定了它的导入路径。

请注意，这与其他编程环境不同，在其他编程环境中，每个项目都有一个独立的工作区，工作区与版本控制代码库紧密相连。

### 工作区

​	工作区是一个目录层次结构，其根目录下有两个目录：

- `src` —— 包含 Go 源文件
- `bin` —— 包含可执行命令

`go`工具会在`bin`目录下构建和安装二进制文件。

`src`子目录通常包含多个版本控制代码库库（如为Git或Mercurial），跟踪一个或多个源代码包的开发。

为了让您了解工作区的实际情况，这里有一个例子：

```text linenums="1"
bin/
    hello                          # command executable
    outyet                         # command executable
src/
    golang.org/x/example/
        .git/                      # Git repository metadata
		hello/
	    	hello.go               # command source
		outyet/
	    	main.go                # command source
	    	main_test.go           # test source
		stringutil/
	    	reverse.go             # package source
	    	reverse_test.go        # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
		bmp/
	    	reader.go              # package source
	    	writer.go              # package source
    ... (many more repositories and packages omitted) ...
```

​	上面的树显示了一个包含两个仓库（`example`和`image`）的工作区。`example` 仓库包含两个命令（`hello` 和 `outyet`）和一个库（`stringutil`）。`image`代码库包含`bmp`包和[several others](https://pkg.go.dev/golang.org/x/image)几个包。

​	典型的工作区包含许多源码代码库，其中包含许多包和命令。大多数 Go 程序员将所有 Go 源代码和依赖关系保存在一个工作区中。

​	请注意，不应该使用`符号链接`将文件或目录链接到您的工作区。

​	命令和库是由不同种类的源码包构建的。我们将在[后面](#package-names)讨论这种区别。

### `GOPATH`环境变量

​	`GOPATH`环境变量指定了您的工作区的位置。它默认为您的主目录内的一个名为`go`的目录，如Unix中的`$HOME/go`，Plan 9中的`$home/go`，以及Windows中的`%USERPROFILE%\go`（通常是`C:\Users\YourName\go`）。

​	如果您希望在不同的位置工作，您需要把`GOPATH`设置为该目录的路径。(另一种常见的设置是设置 `GOPATH=$HOME`。)注意 `GOPATH` 不能与您的 Go 安装路径相同。

​	命令 `go env GOPATH` 打印当前有效的 `GOPATH`；如果环境变量未设置，则打印默认位置。

​	为方便起见，将工作区的 `bin` 子目录添加到您的 `PATH` 中。

```
$ export PATH=$PATH:$(go env GOPATH)/bin
```

​	为了简洁起见，本文其余部分的脚本使用`$GOPATH`而不是`$(go env GOPATH)`。如果您没有设置`GOPATH`，要想让脚本按照写好的内容运行，您可以在这些命令中用`$HOME/go`代替，否则就运行：

```
$ export GOPATH=$(go env GOPATH)
```

​	要了解更多关于`GOPATH`环境变量的信息，请参阅 "`go help gopath`"。

### 导入路径

​	导入路径是一个字符串，它唯一地标识了一个包。包的导入路径与它在工作区或远程代码库中的位置相对应（解释如下）。

​	来自标准库的包被赋予简短的导入路径，如 "`fmt` "和 "`net/http`"。对于您自己的包，您必须选择一个不太可能与未来添加到标准库或其他外部库相冲突的基本路径。

​	如果您把您的代码保存在某个源码库中，那么您应该使用该源码库的根作为您的基本路径。例如，如果您有一个[GitHub](https://github.com/)账户，地址是`github.com/user`，这应该是您的基本路径。

​	注意，您不需要在构建代码之前将其发布到远程仓库。这只是一个好习惯，把您的代码组织起来，就像您有一天会发布它一样。在实践中，您可以选择任何任意的路径名称，只要它对标准库和更大的Go生态系统来说是唯一的。

​	我们将使用`github.com/user`作为我们的基本路径。在您的工作区内创建一个目录，用来保存源代码。

```shell
$ mkdir -p $GOPATH/src/github.com/user
```

### 您的第一个程序

​	要编译和运行一个简单的程序，首先选择一个包的路径（我们将使用`github.com/user/hello`），并在您的工作区中创建一个相应的包目录：

```shell
$ mkdir $GOPATH/src/github.com/user/hello
```

​	接下来，在该目录下创建一个名为`hello.go`的文件，包含以下Go代码。

```go linenums="1" title="hello.go"
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```

现在您可以用`go`工具构建和安装该程序。

```shell
$ go install github.com/user/hello
```

​	注意，您可以在您系统的任何地方运行这个命令。`go`工具通过在`GOPATH`指定的工作区内寻找`github.com/user/hello`包来找到源代码。

​	如果您在包目录下运行`go install`，您也可以省略包路径：

```shell
$ cd $GOPATH/src/github.com/user/hello
$ go install
```

​	该命令构建了`hello`命令，产生了一个可执行的二进制文件。然后，它将该二进制文件安装到工作区的`bin`目录下，称为`hello`（或者，在Windows下，`hello.exe`）。在我们的例子中，这将是`$GOPATH/bin/hello`，也就是`$HOME/go/bin/hello`。

​	`go`工具只有在发生错误时才会打印输出，所以如果这些命令没有产生输出，它们就已经成功执行了。

​	现在您可以通过在命令行中输入程序的完整路径来运行该程序。

```shell
$ $GOPATH/bin/hello
Hello, world.
```

​	或者，由于您已经将`$GOPATH/bin`添加到您的`PATH`中，只需输入二进制名称：

```shell
$ hello
Hello, world.
```

​	如果您使用的是源码控制系统，现在是一个很好的时机来初始化一个仓库，添加文件，并提交您的第一个修改。同样，这一步是可选的：您不需要使用源码控制来编写Go代码。

```shell
$ cd $GOPATH/src/github.com/user/hello
$ git init
Initialized empty Git repository in /home/user/go/src/github.com/user/hello/.git/
$ git add hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 7 insertion(+)
 create mode 100644 hello.go
```

将代码推送到远程仓库是留给读者的一个练习。

### 您的第一个库

​	让我们写一个库并在 `hello` 程序中使用它。

​	同样，第一步是选择一个包路径（我们将使用`github.com/user/stringutil`）并创建包目录：

```shell
$ mkdir $GOPATH/src/github.com/user/stringutil
```

接下来，在该目录下创建一个名为`reverse.go`的文件，内容如下。

```go linenums="1" title="reverse.go"
// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

现在，用`go build`测试软件包的编译情况：

```shell
$ go build github.com/user/stringutil
```

或者，如果您是在包源码目录下工作，只需：

```shell
$ go build
```

这不会产生一个输出文件。相反，它把编译好的包保存在本地构建缓存中。

​	在确认`stringutil`包构建完成后，修改您原来的`hello.go`（它在`$GOPATH/src/github.com/user/hello`中）以使用它。

```go linenums="1" hl_lines="6 6"
package main

import (
	"fmt"

	"github.com/user/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
}
```

安装`hello`程序：

```shell
$ go install github.com/user/hello
```

运行新版本的程序，您应该看到一个新的、反转的信息：

```shell
$ hello
Hello, Go!
```

完成上述步骤后，您的工作区应该是这样的：

```
bin/
    hello                 # command executable
src/
    github.com/user/
        hello/
            hello.go      # command source
        stringutil/
            reverse.go    # package source
```

### 包名

Go源文件中的第一条语句必须是

```go linenums="1"
package name
```

其中`name`是导入包的默认名称。(一个包中的所有文件必须使用相同的`name`)。

​	Go的惯例是，包名是导入路径的最后一个元素：导入为 "`crypto/rot13` "的包应被命名为`rot13`。

​	可执行命令必须始终使用`package main`。

There is no requirement that package names be unique across all packages linked into a single binary, only that the import paths (their full file names) be unique.

​	没有要求包名在链接到一个二进制文件的所有包中是唯一的，只要求导入路径（它们的完整文件名）是唯一的。 =>仍有疑问？？

​	请参阅[Effective Go](../../UsingAndUnderstandingGo/EffectiveGo)以了解更多关于Go的命名规则。

## 测试

​	Go 有一个由 `go test` 命令和`testing`包组成的轻量级测试框架。

​	您可以通过创建一个名称以`_test.go`结尾的文件来编写测试，该文件包含名为`TestXXX`的函数，其签名为`func（t *testing.T）`。测试框架运行每个这样的函数；如果在该函数调用一个失败的函数，如`t.Error`或`t.Fail`，则认为测试失败。

​	通过创建包含以下Go代码的`$GOPATH/src/github.com/user/stringutil/reverse_test.go`文件，向`stringutil`包添加一个测试。

```go linenums="1" title="reverse_test.go"
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

然后用`go test`运行该测试：

```shell
$ go test github.com/user/stringutil
ok  	github.com/user/stringutil 0.165s
```

像往常一样，如果您从包目录中运行`go`工具，您可以省略包路径：

```shell
$ go test
ok  	github.com/user/stringutil 0.165s
```

运行`go help test`，查看[testing package documentation](https://go.dev/pkg/testing/)以了解更多细节。

## 远程包

​	导入路径可以描述如何使用`Git`或`Mercurial`等修订控制系统获得软件包的源代码。`go`工具使用这个属性来自动从远程代码库获取包。例如，本文档中描述的例子也保存在GitHub `golang.org/x/example`的Git代码库中。如果您在包的导入路径中包含代码库的URL，`go get`将自动获取、构建和安装它：

```shell
$ go get golang.org/x/example/hello
$ $GOPATH/bin/hello
Hello, Go examples!
```

​	如果指定的包不存在于工作区，`go get`会把它放在`GOPATH`指定的第一个工作区中。(如果该包已经存在，`go get`将跳过远程获取，其行为与`go install`相同)。

​	在发出上述`go get`命令后，工作区的目录树现在应该是这样的：

```
bin/
    hello                           # command executable
src/
    golang.org/x/example/
	.git/                       # Git repository metadata
        hello/
            hello.go                # command source
        stringutil/
            reverse.go              # package source
            reverse_test.go         # test source
    github.com/user/
        hello/
            hello.go                # command source
        stringutil/
            reverse.go              # package source
            reverse_test.go         # test source
```

The `hello` command hosted at GitHub depends on the `stringutil` package within the same repository. The imports in `hello.go` file use the same import path convention, so the `go get` command is able to locate and install the dependent package, too.

​	在GitHub上托管的`hello`命令依赖于同一代码库中的`stringutil`包。`hello.go`文件中的导入使用了相同的导入路径约定，因此`go get`命令也能定位并安装依赖的包。

```go linenums="1"
import "golang.org/x/example/stringutil"
```

​	这个约定是让您的 Go 包供别人使用的最简单的方法。[pkg.go.dev](https://pkg.go.dev/) 和 [Go Wiki](https://go.dev/wiki/Projects) 提供了外部 Go 项目的列表。

​	关于使用 `go` 工具的远程代码库的更多信息，请参见 `go help importpath`。

## 下一步

Subscribe to the [golang-announce](https://groups.google.com/group/golang-announce) mailing list to be notified when a new stable version of Go is released.

订阅 golang-announce 邮件列表，以便在 Go 的新稳定版本发布时获得通知。

See [Effective Go](https://go.dev/doc/effective_go.html) for tips on writing clear, idiomatic Go code.

请参阅 Effective Go 以了解编写清晰、简洁的 Go 代码的技巧。

Take [A Tour of Go](https://go.dev/tour/) to learn the language proper.

参加 Go 之旅来学习这门语言。

Visit the [documentation page](https://go.dev/doc/#articles) for a set of in-depth articles about the Go language and its libraries and tools.

访问文档页面，了解有关 Go 语言及其库和工具的一系列深度文章。

## 获得帮助

For real-time help, ask the helpful gophers in `#go-nuts` on the [Libera.Chat](https://libera.chat/) IRC server.

要获得实时帮助，请向Libera.Chat IRC服务器上的#go-nuts提问。

The official mailing list for discussion of the Go language is [Go Nuts](https://groups.google.com/group/golang-nuts).

用于讨论 Go 语言的官方邮件列表是 Go Nuts。

Report bugs using the [Go issue tracker](https://go.dev/issue).

使用 Go 问题跟踪器报告错误。