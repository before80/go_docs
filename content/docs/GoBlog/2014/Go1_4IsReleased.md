+++
title = "go 1.4发布了"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.4 is released - go 1.4发布了

https://go.dev/blog/go1.4

Andrew Gerrand
10 December 2014

Today we announce Go 1.4, the fifth major stable release of Go, arriving six months after our previous major release [Go 1.3](https://blog.golang.org/go1.3). It contains a small language change, support for more operating systems and processor architectures, and improvements to the tool chain and libraries. As always, Go 1.4 keeps the promise of compatibility, and almost everything will continue to compile and run without change when moved to 1.4. For the full details, see the [Go 1.4 release notes](https://go.dev/doc/go1.4).

今天我们宣布Go 1.4，这是Go的第五个主要稳定版本，在我们上一个主要版本Go 1.3的六个月后到达。它包含一个小的语言变化，对更多操作系统和处理器架构的支持，以及对工具链和库的改进。一如既往，Go 1.4保持了兼容性的承诺，当转移到1.4时，几乎所有的东西都会继续编译和运行而不发生变化。完整的细节，请参见Go 1.4发布说明。

The most notable new feature in this release is official support for Android. Using the support in the core and the libraries in the [golang.org/x/mobile](https://godoc.org/golang.org/x/mobile) repository, it is now possible to write simple Android apps using only Go code. At this stage, the support libraries are still nascent and under heavy development. Early adopters should expect a bumpy ride, but we welcome the community to get involved.

这个版本中最值得注意的新功能是对Android的官方支持。使用核心中的支持和 golang.org/x/mobile 仓库中的库，现在可以只用 Go 代码编写简单的 Android 应用程序。在这个阶段，支持库仍然是新生的，并且正在大力开发。早期采用者应该期待一个颠簸的旅程，但我们欢迎社区参与。

The language change is a tweak to the syntax of for-range loops. You may now write “for range s {” to loop over each item from s, without having to assign the value, loop index, or map key. See the [release notes](https://go.dev/doc/go1.4#forrange) for details.

语言的变化是对for-range循环的语法进行了调整。你现在可以写 "for range s {"来循环s中的每一个项目，而不需要指定值、循环索引或地图键。详情请见发布说明。

The go command has a new subcommand, go generate, to automate the running of tools to generate source code before compilation. For example, it can be used to automate the generation of String methods for typed constants using the [new stringer tool](https://godoc.org/golang.org/x/tools/cmd/stringer/). For more information, see the [design document](https://go.dev/s/go1.4-generate).

go命令有一个新的子命令，go generate，用于自动运行工具，在编译前生成源代码。例如，它可以用来自动生成字符串方法，用于使用新的stringer工具生成类型化的常量。更多信息请参见设计文档。

Most programs will run about the same speed or slightly faster in 1.4 than in 1.3; some will be slightly slower. There are many changes, making it hard to be precise about what to expect. See the [release notes](https://go.dev/doc/go1.4#performance) for more discussion.

大多数程序在1.4中的运行速度与1.3基本相同或稍快；有些程序会稍慢一些。有许多变化，使我们很难精确地预期到什么。更多的讨论见发行说明。

And, of course, there are many more improvements and bug fixes.

当然，还有更多的改进和错误修复。

In case you missed it, a few weeks ago the sub-repositories were moved to new locations. For example, the go.tools packages are now imported from “golang.org/x/tools”. See the [announcement post](https://groups.google.com/d/msg/golang-announce/eD8dh3T9yyA/HDOEU_ZSmvAJ) for details.

如果你错过了，几周前，子库被移到了新的位置。例如，go.tools包现在从 "golang.org/x/tools "导入。详情请见公告帖子。

This release also coincides with the project’s move from Mercurial to Git (for source control), Rietveld to Gerrit (for code review), and Google Code to Github (for issue tracking and wiki). The move affects the core Go repository and its sub-repositories. You can find the canonical Git repositories at [go.googlesource.com](https://go.googlesource.com/), and the issue tracker and wiki at the [golang/go GitHub repo](https://github.com/golang/go).

这个版本也与项目从Mercurial转移到Git（用于源代码控制），从Rietveld转移到Gerrit（用于代码审查），以及从Google Code转移到Github（用于问题跟踪和wiki）相吻合。此举影响了核心Go仓库及其子仓库。你可以在go.googlesource.com找到规范的Git仓库，在golang/go GitHub repo找到问题跟踪器和维基。

While development has already moved over to the new infrastructure, for the 1.4 release we still recommend that users who [install from source](https://go.dev/doc/install/source) use the Mercurial repositories.

虽然开发工作已经转移到新的基础设施，但对于1.4版本，我们仍然建议从源代码安装的用户使用Mercurial仓库。

For App Engine users, Go 1.4 is now available for beta testing. See [the announcement](https://groups.google.com/d/msg/google-appengine-go/ndtQokV3oFo/25wV1W9JtywJ) for details.

对于App Engine用户，Go 1.4现在可以进行测试了。详情请见公告。

From all of us on the Go team, please enjoy Go 1.4, and have a happy holiday season.

来自Go团队的所有成员，请享受Go 1.4，并祝您节日快乐。
