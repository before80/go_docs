+++
title = "Go 1.2发布了"
weight = 3
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.2 is released - Go 1.2发布了

> 原文：[https://go.dev/blog/go12](https://go.dev/blog/go12)

Andrew Gerrand
1 December 2013

We are pleased to announce the release of Go 1.2, the latest stable version of the Go Programming Language.

我们很高兴地宣布Go 1.2的发布，这是Go编程语言的最新稳定版本。

Binary distributions may be downloaded from the [usual place](https://go.dev/doc/install) or if you prefer to [compile from source](https://go.dev/doc/install/source) you should use the `release` or `go1.2` tags.

二进制发行版可以从通常的地方下载，如果您喜欢从源代码编译，您应该使用 release 或 go1.2 标签。

This new release comes nearly seven months after the release of Go 1.1 in May, a much shorter period than the 14 months between 1.1 and 1.0. We anticipate a comparable interval between future major releases.

这个新版本是在5月发布Go 1.1后的近7个月内发布的，比1.1和1.0之间的14个月要短很多。我们预计未来的主要版本之间也会有类似的时间间隔。

[Go 1.2](https://go.dev/doc/go1.2) includes a couple of minor language changes, several improvements to the language implementation and tools, some performance improvements, and many additions and (backward-compatible) changes to the standard library.

Go 1.2包括一些小的语言变化，对语言实现和工具的一些改进，一些性能改进，以及对标准库的许多补充和（向后兼容的）变化。

Please read the [release notes](https://go.dev/doc/go1.2) for all the details, as some changes may affect the behavior of existing (buggy) programs. What follows is the highlights of the release.

请阅读发行说明以了解所有细节，因为一些变化可能会影响到现有（有问题）程序的行为。以下是该版本的亮点。

A new [three-index slice syntax](https://go.dev/doc/go1.2#three_index) adds the ability to specify capacity as well as length. This allows the programmer to pass a slice value that can only access a limited portion of the underlying array, a technique that previously required the use of the unsafe package.

一个新的三索引切片语法增加了指定容量和长度的能力。这允许程序员传递一个只能访问底层数组中有限部分的分片值，这种技术以前需要使用不安全包。

A major new feature of the tool chain is the facility to compute and display [test coverage results](https://go.dev/doc/go1.2#cover). See the [`go test`](https://go.dev/cmd/go/#hdr-Description_of_testing_flags) and [cover tool](https://golang.org/x/tools/cmd/cover) documentation for details. Later this week we will publish an article that discusses this new feature in detail.

该工具链的一个主要新功能是计算和显示测试覆盖率结果的设施。详情请见go测试和覆盖工具文档。本周晚些时候，我们将发表一篇文章，详细讨论这个新功能。

Goroutines are now [pre-emptively scheduled](https://go.dev/doc/go1.2#preemption), in that the scheduler is invoked occasionally upon entry to a function. This can prevent busy goroutines from starving other goroutines on the same thread.

现在，Goroutines被预先调度了，在进入一个函数时，调度器被偶尔调用。这可以防止繁忙的goroutines使同一线程上的其他goroutines陷入饥饿。

An increase to the default goroutine stack size should improve the performance of some programs. (The old size had a tendency to introduce expensive stack-segment switching in performance-critical sections.) On the other end, new restrictions on [stack sizes](https://go.dev/doc/go1.2#stack_size) and [the number of operating system threads](https://go.dev/doc/go1.2#thread_limit) should prevent misbehaving programs from consuming all the resources of a machine. (These limits may be adjusted using new functions in the [`runtime/debug` package](https://go.dev/pkg/runtime/debug).)

增加默认的goroutine堆栈大小应该可以改善一些程序的性能。(旧的大小有一种倾向，即在性能关键部分引入昂贵的堆栈段切换。） 在另一端，对堆栈大小和操作系统线程数量的新限制应该可以防止行为不端的程序消耗机器的所有资源。(这些限制可以通过运行时/调试包中的新函数来调整）。

Finally, among the [many changes to the standard library](https://go.dev/doc/go1.2#library), significant changes include the new [`encoding` package](https://go.dev/doc/go1.2#encoding), [indexed arguments](https://go.dev/doc/go1.2#fmt_indexed_arguments) in `Printf` format strings, and some [convenient additions](https://go.dev/doc/go1.2#text_template) to the template packages.

最后，在标准库的众多变化中，重要的变化包括新的编码包、Printf格式字符串中的索引参数，以及模板包的一些方便的补充。

As part of the release, the [Go Playground](http://play.golang.org/) has been updated to Go 1.2. This also affects services that use the Playground, such as [the Go Tour](https://go.dev/tour/) and this blog. The update also adds the ability to use threads and the `os`, `net`, and `unsafe` packages inside the sandbox, making it more like a real Go environment.

作为该版本的一部分，Go Playground已经更新到Go 1.2。这也影响了使用Playground的服务，如Go Tour和本博客。这次更新还增加了在沙盒内使用线程和os、net、unsafe包的能力，使其更像一个真正的Go环境。

To everyone who helped make this release possible, from the many users who submitted bug reports to the 116 (!) contributors who committed more than 1600 changes to the core: Your help is invaluable to the project. Thank you!

对于每一个帮助实现这个版本的人，从提交错误报告的许多用户到对核心部分进行了1600多处修改的116位（！）贡献者：您们的帮助对于这个项目是非常宝贵的。谢谢您们!

*This blog post is the first of the* [Go Advent Calendar](http://blog.gopheracademy.com/day-01-go-1.2), *a series of daily articles presented by the* [Gopher Academy](http://gopheracademy.com/) *from December 1 to 25.*

这篇博文是Go Advent Calendar的第一篇，这是Gopher Academy从12月1日至25日推出的一系列每日文章。
