+++
title = "关于 go 命令"
weight = 2
date = 2023-05-18T17:31:23+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# About the go command - 关于 go 命令

> 原文：[https://go.dev/doc/articles/go_command](https://go.dev/doc/articles/go_command)

​	Go 发行版包含一个名为 "`go`"的命令，该命令自动下载、构建、安装和测试Go包和命令。本文讨论了我们为什么编写了一个新的命令，它是什么，它不是什么，以及如何使用它。

## 动机

​	您可能已经看到早期的Go演讲，Rob Pike在其中开玩笑说Go的想法是在等待大型Google服务器进行编译时产生的。这确实是Go的动机：构建一种适用于编写和运行Google大型软件的语言。从一开始就清楚，这种语言必须提供一种明确表达代码库之间依赖关系的方式，因此有了包组织和显式导入块。从一开始就清楚，您可能需要任意的语法来描述导入的代码；这就是为什么导入路径是字符串文字的原因。

​	Go的一个明确目标是只使用源代码中找到的信息构建Go代码，而不需要编写makefile或现代替代makefile的许多工具。如果Go需要一个配置文件来解释如何构建您的程序，那么Go将失败。

​	起初，没有Go编译器，最初的开发重点是构建一个编译器，然后为其构建库。为了迅速，我们推迟了自动化构建Go代码的工作，使用make和编写makefile。当编译单个包涉及多次调用Go编译器时，我们甚至使用一个程序来为我们编写makefile。如果您查看存储库历史记录，可以找到它。

​	新go命令的目的是回归这个理想，即Go程序应该无需配置或开发人员除了编写必要的导入语句外，不需要额外的努力即可编译。

## 配置与约定

​		实现无需配置的简单系统的方法是建立约定。系统只有在遵循这些约定的程度上才能发挥作用。当 Go 首次发布时，许多人发布的软件包必须安装在特定的位置，使用特定的构建工具才能使用，这是可以理解的，因为大多数其他编程语言都是这样的。在过去的几年中，我们一直提醒人们 `goinstall` 命令(现在已被 [go get](../References/CommandDocumentation/go) 命令取代)及其约定，即导入路径的推导方法是从源代码的 URL 中获取的，本地文件系统存储源代码的位置是从导入路径的推导方法中获取的，源代码树中的每个目录对应一个软件包，且该软件包只使用源代码中的信息进行构建。今天，绝大多数软件包都遵循这些约定。结果，Go 生态系统变得更加简单和强大。

​	我们收到了许多允许在软件包目录中提供 makefile 以提供比源代码更少的额外配置的请求。但这会引入新的规则。由于我们没有同意这样的请求，我们能够编写 go 命令并消除我们对 make 或任何其他构建系统的使用。

​	重要的是要理解，go 命令不是一个通用的构建工具。它不能进行配置，也不尝试构建除 Go 软件包之外的任何东西。这些是重要的简化假设：它们不仅简化了实现，而且更重要的是简化了工具本身的使用。

## Go 的约定

​	`go` 命令要求代码遵循一些关键的、成熟的约定。	

​	首先，导入路径是从源代码的 URL 中以已知的方式推导出来的。对于 Bitbucket、GitHub、Google Code 和 Launchpad，仓库的根目录由仓库的主 URL 识别，不包括 `https://` 前缀。子目录是通过在路径上添加来命名的。例如，可以通过运行以下命令来获取 Google 的 glog 日志软件包的源代码：

```shell
git clone https://github.com/golang/glog
```

因此，[glog ](https://pkg.dev.go/github.com/golang/glog)包的导入路径是 "`github.com/golang/glog`"。

​	这些路径有些冗长，但换取的是导入路径的自动管理名称空间和像 go 命令这样的工具可以查看陌生的导入路径并推断出从哪里获取源代码的能力。

​	其次，本地文件系统中存储源代码的位置是从导入路径派生的，具体地说是`$GOPATH/src/<import-path>`。如果未设置，则`$GOPATH`默认为用户主目录中名为go的子目录。如果`$GOPATH`设置为路径列表，则go命令尝试对该列表中的每个目录使用`<dir>/src/<import-path>`。

​	按照惯例，这些树的每个目录都包含一个名为"bin"的顶级目录，用于保存编译后的可执行文件，以及一个名为"pkg"的顶级目录，用于保存可导入的已编译包，以及包含包源文件的"src"目录。强制使用此结构使我们可以使每个目录树都是自包含的：编译形式和源文件始终在彼此附近。

​	这些命名约定还使我们能够从目录名反向进行映射到其导入路径。这个映射对于许多go命令的子命令非常重要，如下面所述。	

​	第三，源树中的每个目录对应一个单独的包。通过将一个目录限制为单个包，我们不必创建混合导入路径，这些导入路径首先指定目录，然后指定该目录中的包。此外，大多数文件管理工具和用户界面都使用目录作为基本单位。将基本的Go单元——包——与文件系统结构相结合，意味着文件系统工具成为Go包工具。复制、移动或删除一个包对应于复制、移动或删除一个目录。

​	第四，每个包都是使用源文件中存在的信息构建的。这使得该工具更有可能适应不断变化的构建环境和条件。例如，如果允许额外的配置，例如编译器标志或命令行配方，那么每次构建工具发生变化时，该配置都需要更新；它也固有地与使用特定工具链有关。

## 使用go命令开始入门

​	最后，快速浏览如何使用go命令。如上所述，Unix上的默认`$GOPATH`是`$HOME/go`。我们将在那里存储我们的程序。要使用不同的位置，您可以设置`$GOPATH`；有关详细信息，请参阅[如何编写Go代码](../GettingStarted/HowToWriteGoCode)。

​	首先添加一些源代码。假设我们想要使用`codesearch`项目中的索引库以及左倾红黑树。我们可以使用"`go get`"子命令安装两者：

```shell
$ go get github.com/google/codesearch/index
$ go get github.com/petar/GoLLRB/llrb
$
```

​	这两个项目现在都被下载并安装到了 `$HOME/go` 目录下，其中包含两个目录 `src/github.com/google/codesearch/index/` 和 `src/github.com/petar/GoLLRB/llrb/`，以及它们的依赖库和编译后的包(在 `pkg/` 目录下)。

​	由于我们使用版本控制系统(`Mercurial` 和 `Git`)来检出源代码，因此源代码树中还包含了对应代码库中的其他文件，例如相关的包。 "`go list`" 命令列出了其参数所对应的导入路径，而模式 "`./...`" 表示从当前目录 ("`./`") 开始，向下找到所有包 ("`...`")：

```shell
$ cd $HOME/go/src
$ go list ./...
github.com/google/codesearch/cmd/cgrep
github.com/google/codesearch/cmd/cindex
github.com/google/codesearch/cmd/csearch
github.com/google/codesearch/index
github.com/google/codesearch/regexp
github.com/google/codesearch/sparse
github.com/petar/GoLLRB/example
github.com/petar/GoLLRB/llrb
$
```

我们也可以测试这些包：

```shell
$ go test ./...
?   	github.com/google/codesearch/cmd/cgrep	[no test files]
?   	github.com/google/codesearch/cmd/cindex	[no test files]
?   	github.com/google/codesearch/cmd/csearch	[no test files]
ok  	github.com/google/codesearch/index	0.203s
ok  	github.com/google/codesearch/regexp	0.017s
?   	github.com/google/codesearch/sparse	[no test files]
?       github.com/petar/GoLLRB/example          [no test files]
ok      github.com/petar/GoLLRB/llrb             0.231s
$
```

如果 go 子命令在没有列出路径的情况下被调用，则该命令将在当前目录上运行：

```shell
$ cd github.com/google/codesearch/regexp
$ go list
github.com/google/codesearch/regexp
$ go test -v
=== RUN   TestNstateEnc
--- PASS: TestNstateEnc (0.00s)
=== RUN   TestMatch
--- PASS: TestMatch (0.00s)
=== RUN   TestGrep
--- PASS: TestGrep (0.00s)
PASS
ok  	github.com/google/codesearch/regexp	0.018s
$ go install
$
```

​	该 "`go install`" 子命令将该包的最新副本安装到 pkg 目录中。由于 go 命令能够分析依赖关系图，"`go install`" 还会递归地安装此包导入但过时的任何包。

​	注意，"`go install`" 能够通过目录命名的惯例来确定当前目录中包的导入路径名称。如果我们可以选择源代码所在目录的名称，那么会更方便一些，而且我们可能不会选择如此冗长的名称。但这种能力需要工具进行额外的配置和复杂性。在输入一个或两个额外的目录名称的情况下，获取到的简单性和功能性是一个很小的代价。

## 限制

​	如上所述，go 命令不是通用的构建工具。特别是，它没有任何生成 Go 源文件的功能，尽管它提供了 [go generate](../References/CommandDocumentation/go)，可以在构建之前自动创建 Go 文件。对于更高级的构建设置，您可能需要编写一个 makefile(或您选择的构建工具的配置文件)，以运行创建 Go 文件的任何工具，然后将这些生成的源文件检入到您的仓库中。这对于包的作者来说可能需要更多的工作，但对于您的用户来说，他们可以使用 "`go get`" 而无需获取和构建任何其他工具，这显然是更简单的。

## 更多信息

​	更多信息，请阅读[How to Write Go Code](../GettingStarted/HowToWriteGoCode)并查看[go命令文档](../References/CommandDocumentation/go)。

