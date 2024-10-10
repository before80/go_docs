+++
title = "Go 1.18 Beta 1已经推出，有泛型的功能"
weight = 84
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.18 Beta 1 is available, with generics - go 1.18 Beta 1已经推出，有泛型的功能

> 原文：[https://go.dev/blog/go1.18beta1](https://go.dev/blog/go1.18beta1)

Russ Cox, for the Go team
14 December 2021

We have just released Go 1.18 Beta 1, which you can get by visiting the [download page](https://go.dev/dl/#go1.18beta1).

我们刚刚发布了Go 1.18 Beta 1，您可以通过访问下载页面获得该版本。

The official Go 1.18 release won’t happen for a couple of months yet. This is the first preview release of Go 1.18, to let you kick the tires, take it for a spin, and let us know what problems you encounter. Go 1.18 Beta 1 represents an enormous amount of work by the entire Go team at Google and Go contributors around the world, and we’re excited to hear what you think.

Go 1.18的正式发布还需要几个月的时间。这是Go 1.18的第一个预览版，目的是让您试一试，转一转，并让我们知道您遇到了什么问题。Go 1.18 Beta 1代表了Google的整个Go团队和世界各地的Go贡献者的大量工作，我们很高兴听到您的想法。

Go 1.18 Beta 1 is the first preview release containing Go’s new support for [generic code using parameterized types](https://go.dev/blog/why-generics). Generics are the most significant change to Go since the release of Go 1, and certainly the largest single language change we’ve ever made. With any large, new feature, it is common for new users to discover new bugs, and we don’t expect generics to be an exception to this rule; be sure to approach them with appropriate caution. Also, certain subtle cases, such as specific kinds of recursive generic types, have been postponed to future releases. That said, we know of early adopters who have been quite happy, and if you have use cases that you think are particularly suited to generics, we hope you will give them a try. We’ve published a [brief tutorial about how to get started with generics](https://go.dev/doc/tutorial/generics) and gave a [talk at GopherCon last week](https://www.youtube.com/watch?v=35eIxI_n5ZM&t=1755s). You can even try it on the [Go playground in Go dev branch mode](https://go.dev/play/?v=gotip).

Go 1.18 Beta 1是第一个预览版，包含Go对使用参数化类型的通用代码的新支持。泛型是Go 1发布以来最重要的变化，当然也是我们有史以来最大的单一语言变化。对于任何大型的新功能，新用户发现新的错误是很常见的，我们不希望泛型是这个规则的例外；一定要以适当的谨慎态度对待它们。另外，某些微妙的情况，例如特定种类的递归泛型，已经被推迟到未来的版本。也就是说，我们知道一些早期采用者已经相当满意，如果您有您认为特别适合泛型的用例，我们希望您能试一试。我们已经发布了一个关于如何开始使用泛型的简短教程，并在上周的GopherCon上做了一个演讲。您甚至可以在Go开发分支模式下的Go操场上进行尝试。

Go 1.18 Beta 1 adds built-in support for writing [fuzzing-based tests](https://go.dev/blog/fuzz-beta), to automatically find inputs that cause your program to crash or return invalid answers.

Go 1.18 Beta 1 增加了对编写基于模糊测试的内置支持，以自动查找导致程序崩溃或返回无效答案的输入。

Go 1.18 Beta 1 adds a new "[Go workspace mode](https://go.dev/design/45713-workspace)", which lets you work with multiple Go modules simultaneously, an important use case for larger projects.

Go 1.18 Beta 1增加了一个新的 "Go工作区模式"，让您可以同时处理多个Go模块，这对大型项目来说是一个重要的使用案例。

Go 1.18 Beta 1 contains an expanded `go version -m` command, which now records build details such as compiler flags. A program can query its own build details using [debug.ReadBuildInfo](https://pkg.go.dev/runtime/debug@master#BuildInfo), and it can now read build details from other binaries using the new [debug/buildinfo](https://pkg.go.dev/debug/buildinfo@master) package. This functionality is meant to be the foundation for any tool that needs to produce a software bill of materials (SBOM) for Go binaries.

Go 1.18 Beta 1包含一个扩展的go版本-m命令，它现在可以记录编译器标志等构建细节。程序可以使用debug.ReadBuildInfo查询自己的构建细节，现在也可以使用新的debug/buildinfo包从其他二进制文件读取构建细节。这一功能旨在为任何需要为Go二进制文件生成软件材料清单（SBOM）的工具奠定基础。

Earlier this year, Go 1.17 added a new register-based calling convention to speed up Go code on x86-64 systems. Go 1.18 Beta 1 expands that feature to ARM64 and PPC64, resulting in as much as 20% speed-ups.

今年早些时候，Go 1.17增加了一个新的基于寄存器的调用约定，以加快X86-64系统上的Go代码。Go 1.18 Beta 1将这一功能扩展到了ARM64和PPC64，使其速度提高了20%之多。

Thanks to everyone who contributed to this beta release, and especially to the team here at Google who has been working tirelessly for years on making generics a reality. It’s been a long road, we’re very happy with the result, and we hope you like it too.

感谢所有为这个测试版做出贡献的人，尤其要感谢谷歌的团队，他们多年来一直在为实现泛型而不懈努力。这是一条漫长的道路，我们对结果非常满意，我们希望您也喜欢它。

See the full [draft release notes for Go 1.18](https://tip.golang.org/doc/go1.18) for more details.

更多细节请参见Go 1.18的完整发布说明草案。

As always, especially for beta releases, if you notice any problems, please [file an issue](https://go.dev/issue/new).

像往常一样，特别是对于测试版，如果您发现任何问题，请提交一个问题。

We hope you enjoy testing the beta, and we hope you all have a restful remainder of 2021. Happy holidays!

我们希望您喜欢测试这个测试版，并希望您在2021年的剩余时间里有一个安宁的生活。节日快乐!
