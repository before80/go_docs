+++
title = "快速入门"
date = 2024-12-09T07:59:19+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/cli/getting_started.md](https://github.com/go-delve/delve/blob/master/Documentation/cli/getting_started.md)
>
> 收录该文档时间： `2024-12-09T07:59:19+08:00`

# Getting Started - 快速入门



Delve aims to be a very simple and powerful tool, but can be confusing if you're not used to using a source level debugger in a compiled language. This document will provide all the information you need to get started debugging your Go programs.

​	Delve 旨在成为一个非常简单而强大的工具，但如果你不习惯使用编译语言中的源代码级调试器，它可能会让人感到困惑。本文档将提供所有你需要的信息，帮助你开始调试 Go 程序。

## 调试 'main' 包 Debugging 'main' packages



The first CLI subcommand we will explore is `debug`. This subcommand can be run without arguments if you're in the same directory as your `main` package, otherwise it optionally accepts a package path.

​	我们将首先探索的 CLI 子命令是 `debug`。如果你在与 `main` 包相同的目录中，可以不带参数运行此子命令，否则它会接受一个可选的包路径。

For example given this project layout:

​	例如，给定如下的项目结构：

```
github.com/me/foo
├── cmd
│   └── foo
│       └── main.go
└── pkg
    └── baz
        ├── bar.go
        └── bar_test.go
```



If you are in the directory `github.com/me/foo/cmd/foo` you can simply run `dlv debug` from the command line. From anywhere else, say the project root, you can simply provide the package: `dlv debug github.com/me/foo/cmd/foo`. To pass flags to your program separate them with `--`: `dlv debug github.com/me/foo/cmd/foo -- -arg1 value`.

​	如果你在目录 `github.com/me/foo/cmd/foo` 中，你可以直接从命令行运行 `dlv debug`。如果你在其他地方，例如项目根目录，你可以提供包路径：`dlv debug github.com/me/foo/cmd/foo`。要向程序传递标志，可以使用 `--` 来分隔：`dlv debug github.com/me/foo/cmd/foo -- -arg1 value`。

Invoking that command will cause Delve to compile the program in a way most suitable for debugging, then it will execute and attach to the program and begin a debug session. Now, when the debug session has first started you are at the very beginning of the program's initialization. To get to someplace more useful you're going to want to set a breakpoint or two and continue execution to that point.

​	执行该命令将导致 Delve 以最适合调试的方式编译程序，然后执行并附加到程序并开始调试会话。现在，当调试会话首次启动时，你正处于程序初始化的最开始阶段。为了到达更有用的地方，你需要设置一个或多个断点，并继续执行到那个位置。

For example, to continue execution to your program's `main` function:

​	例如，要继续执行到程序的 `main` 函数：

```
$ dlv debug github.com/me/foo/cmd/foo
Type 'help' for list of commands.
(dlv) break main.main
Breakpoint 1 set at 0x49ecf3 for main.main() ./test.go:5
(dlv) continue
> main.main() ./test.go:5 (hits goroutine(1):1 total:1) (PC: 0x49ecf3)
     1:	package main
     2:	
     3:	import "fmt"
     4:	
=>   5:	func main() {
     6:		fmt.Println("delve test")
     7:	}
(dlv) 
```



## 调试测试 Debugging tests



Given the same directory structure as above you can debug your code by executing your test suite. For this you can use the `dlv test` subcommand, which takes the same optional package path as `dlv debug`, and will also build the current package if not given any argument.

​	给定与上面相同的目录结构，你可以通过执行测试套件来调试你的代码。为此，你可以使用 `dlv test` 子命令，它接受与 `dlv debug` 相同的可选包路径，并且如果未提供任何参数，它还会构建当前包。

```
$ dlv test github.com/me/foo/pkg/baz
Type 'help' for list of commands.
(dlv) funcs test.Test*
/home/me/go/src/github.com/me/foo/pkg/baz/test.TestHi
(dlv) break TestHi
Breakpoint 1 set at 0x536513 for /home/me/go/src/github.com/me/foo/pkg/baz/test.TestHi() ./test_test.go:5
(dlv) continue
> /home/me/go/src/github.com/me/foo/pkg/baz/test.TestHi() ./bar_test.go:5 (hits goroutine(5):1 total:1) (PC: 0x536513)
     1:	package baz
     2:	
     3:	import "testing"
     4:	
=>   5:	func TestHi(t *testing.T) {
     6:		t.Fatal("implement me!")
     7:	}
(dlv) 
```



As you can see, we began debugging the test binary, found our test function via the `funcs` command which takes a regexp to filter the list of functions, set a breakpoint and then continued execution until we hit that breakpoint.

​	如你所见，我们开始调试测试二进制文件，通过 `funcs` 命令（它使用正则表达式来过滤函数列表）找到我们的测试函数，设置了一个断点，然后继续执行，直到我们击中该断点。

For more information on subcommands you can use, type `dlv help`, and once in a debug session you can see all of the commands available to you by typing `help` at any time.

​	有关可用子命令的更多信息，可以输入 `dlv help`，在调试会话中随时输入 `help`，你可以查看所有可用的命令。
