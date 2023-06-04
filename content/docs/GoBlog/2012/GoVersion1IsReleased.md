+++
title = "go 第一版发布"
weight = 6
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go version 1 is released - go 第一版发布

https://go.dev/blog/go1

Andrew Gerrand
28 March 2012

![img](GoVersion1IsReleased_img/gophermega.jpg)

Today marks a major milestone in the development of the Go programming language. We’re announcing Go version 1, or Go 1 for short, which defines a language and a set of core libraries to provide a stable foundation for creating reliable products, projects, and publications.

今天是Go编程语言发展中的一个重要里程碑。我们宣布Go版本1，简称Go 1，它定义了一种语言和一套核心库，为创建可靠的产品、项目和出版物提供了稳定的基础。

Go 1 is the first release of Go that is available in supported binary distributions. They are available for Linux, FreeBSD, Mac OS X and, we are thrilled to announce, Windows.

Go 1是Go的第一个版本，可用于支持的二进制发行版。它们可用于Linux、FreeBSD、Mac OS X以及我们很高兴地宣布的Windows。

The driving motivation for Go 1 is stability for its users. People who write Go 1 programs can be confident that those programs will continue to compile and run without change, in many environments, on a time scale of years. Similarly, authors who write books about Go 1 can be sure that their examples and explanations will be helpful to readers today and into the future.

Go 1的驱动力是其用户的稳定性。编写Go 1程序的人可以确信，这些程序将在许多环境中继续编译和运行，不会有任何变化，而且时间尺度为数年。同样地，编写Go 1书籍的作者也可以确信，他们的例子和解释在今天和未来都会对读者有所帮助。

Forward compatibility is part of stability. Code that compiles in Go 1 should, with few exceptions, continue to compile and run throughout the lifetime of that version, even as we issue updates and bug fixes such as Go version 1.1, 1.2, and so on. The [Go 1 compatibility document](https://go.dev/doc/go1compat.html) explains the compatibility guidelines in more detail.

前向兼容性是稳定性的一部分。在 Go 1 中编译的代码，除了少数例外，应该在该版本的整个生命周期内继续编译和运行，即使我们发布了更新和错误修复，如 Go 1.1、1.2 等版本。Go 1 兼容性文件更详细地解释了兼容性准则。

Go 1 is a representation of Go as it is used today, not a major redesign. In its planning, we focused on cleaning up problems and inconsistencies and improving portability. There had long been many changes to Go that we had designed and prototyped but not released because they were backwards-incompatible. Go 1 incorporates these changes, which provide significant improvements to the language and libraries but sometimes introduce incompatibilities for old programs. Fortunately, the [go fix](https://go.dev/cmd/go/#Run_go_tool_fix_on_packages) tool can automate much of the work needed to bring programs up to the Go 1 standard.

Go 1是对目前使用的Go的代表，而不是一个重大的重新设计。在其规划中，我们专注于清理问题和不一致之处，并提高可移植性。长期以来，我们对 Go 进行了许多修改，并设计了原型，但因为向后不兼容而没有发布。Go 1包含了这些变化，这些变化为语言和库提供了显著的改进，但有时会给旧程序带来不兼容的问题。幸运的是，go fix工具可以自动完成大部分的工作，使程序达到Go 1的标准。

Go 1 introduces changes to the language (such as new types for [Unicode characters](https://go.dev/doc/go1.html#rune) and [errors](https://go.dev/doc/go1.html#errors)) and the standard library (such as the new [time package](https://go.dev/doc/go1.html#time) and renamings in the [strconv package](https://go.dev/doc/go1.html#strconv)). Also, the package hierarchy has been rearranged to group related items together, such as moving the networking facilities, for instance the [rpc package](https://go.dev/pkg/net/rpc/), into subdirectories of net. A complete list of changes is documented in the [Go 1 release notes](https://go.dev/doc/go1.html). That document is an essential reference for programmers migrating code from earlier versions of Go.

Go 1对语言（如Unicode字符和错误的新类型）和标准库（如新的时间包和strconv包的重命名）进行了修改。另外，包的层次结构也被重新安排，以便将相关的项目放在一起，例如将网络设施，例如rpc包，移到net的子目录中。完整的变化列表见 Go 1 发行说明。该文件是程序员从Go早期版本迁移代码的重要参考。

We also restructured the Go tool suite around the new [go command](https://go.dev/doc/go1.html#cmd_go), a program for fetching, building, installing and maintaining Go code. The go command eliminates the need for Makefiles to write Go code because it uses the Go program source itself to derive the build instructions. No more build scripts!

我们还围绕新的 go 命令重组了 Go 工具套件，这是一个用于获取、构建、安装和维护 Go 代码的程序。go命令消除了编写Go代码时对Makefile的需求，因为它使用Go程序源代码本身来获取构建指令。没有更多的构建脚本!

Finally, the release of Go 1 triggers a new release of the [Google App Engine SDK](https://developers.google.com/appengine/docs/go). A similar process of revision and stabilization has been applied to the App Engine libraries, providing a base for developers to build programs for App Engine that will run for years.

最后，Go 1的发布引发了Google App Engine SDK的新版本。类似的修订和稳定过程已经应用于App Engine库，为开发者提供了一个基础，以便为App Engine构建可以运行多年的程序。

Go 1 is the result of a major effort by the core Go team and our many contributors from the open source community. We thank everyone who helped make this happen.

Go 1是Go核心团队和来自开源社区的众多贡献者的重大努力的结果。我们感谢所有帮助实现这一目标的人。

There has never been a better time to be a Go programmer. Everything you need to get started is at [golang.org](https://go.dev/).

现在是成为Go程序员的最佳时机。您所需要的一切都在golang.org上。
