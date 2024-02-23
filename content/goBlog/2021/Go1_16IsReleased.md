+++
title = "Go 1.16版发布了"
weight = 96
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.16 is released - go 1.16版发布了

> 原文：[https://go.dev/blog/go1.16](https://go.dev/blog/go1.16)

Matt Pearring and Dmitri Shuralyov
16 February 2021

Today the Go team is very happy to announce the release of Go 1.16. You can get it from the [download page](https://go.dev/dl/).

今天，Go团队非常高兴地宣布Go 1.16的发布。您可以从下载页面获得它。

The new [embed package](https://go.dev/doc/go1.16#library-embed) provides access to files embedded at compile time using the new `//go:embed` directive. Now it is easy to bundle supporting data files into your Go programs, making developing with Go even smoother. You can get started using the [embed package documentation](https://pkg.go.dev/embed). Carl Johnson has also written a nice tutorial, "[How to use Go embed](https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/)".

新的嵌入包提供了对使用新的//go:embed指令在编译时嵌入文件的访问。现在可以很容易地将支持性的数据文件捆绑到您的Go程序中，使Go的开发更加顺畅。您可以使用embed包的文档开始学习。卡尔-约翰逊也写了一个很好的教程，"如何使用 Go 嵌入"。

Go 1.16 also adds [macOS ARM64 support](https://go.dev/doc/go1.16#darwin) (also known as Apple silicon). Since Apple’s announcement of their new arm64 architecture, we have been working closely with them to ensure Go is fully supported; see our blog post "[Go on ARM and Beyond](https://blog.golang.org/ports)" for more.

Go 1.16还增加了对macOS ARM64的支持（也被称为苹果硅）。自从苹果公司宣布其新的 arm64 架构以来，我们一直在与他们密切合作，以确保 Go 得到全面支持；更多信息请参见我们的博文 "Go on ARM and Beyond"。

Note that Go 1.16 [requires use of Go modules by default](https://go.dev/doc/go1.16#modules), now that, according to our 2020 Go Developer Survey, 96% of Go developers have made the switch. We recently added official documentation for [developing and publishing modules](https://go.dev/doc/modules/developing).

请注意，Go 1.16要求默认使用Go模块，现在，根据我们的2020年Go开发者调查，96%的Go开发者已经进行了转换。我们最近增加了开发和发布模块的官方文档。

Finally, there are many other improvements and bug fixes, including builds that are up to 25% faster and use as much as 15% less memory. For the complete list of changes and more information about the improvements above, see the [Go 1.16 release notes](https://go.dev/doc/go1.16).

最后，还有许多其他改进和错误修复，包括构建速度提高了25%，内存使用量减少了15%。关于完整的变化列表和有关上述改进的更多信息，请参阅Go 1.16发布说明。

We want to thank everyone who contributed to this release by writing code filing bugs, providing feedback, and testing the beta and release candidate.

我们要感谢所有为这个版本做出贡献的人，他们编写代码填补漏洞，提供反馈，并测试测试版和候选版。

Your contributions and diligence helped to ensure that Go 1.16 is as stable as possible. That said, if you notice any problems, please [file an issue](https://go.dev/issue/new).

您的贡献和勤奋有助于确保Go 1.16尽可能的稳定。也就是说，如果您发现任何问题，请提出问题。

We hope you enjoy the new release!

我们希望您喜欢这个新版本
