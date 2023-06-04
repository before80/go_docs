+++
title = "go 1.18发布了!"
weight = 97
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.18 is released!  - go 1.18发布了!

https://go.dev/blog/go1.18

The Go Team
15 March 2022

Today the Go team is thrilled to release Go 1.18, which you can get by visiting the [download page](https://go.dev/dl/).

今天，Go团队很高兴地发布了Go 1.18，您可以通过访问下载页面获得该版本。

Go 1.18 is a massive release that includes new features, performance improvements, and our biggest change ever to the language. It isn’t a stretch to say that the design for parts of Go 1.18 started over a decade ago when we first released Go.

Go 1.18是一个巨大的版本，包括新的功能、性能改进，以及我们对语言有史以来最大的改变。可以说，Go 1.18的部分设计始于十年前我们首次发布Go时，这并不夸张。

## Generics 泛型

In Go 1.18, we’re introducing new support for [generic code using parameterized types](https://go.dev/blog/why-generics). Supporting generics has been Go’s most often requested feature, and we’re proud to deliver the generic support that the majority of users need today. Subsequent releases will provide additional support for some of the more complicated generic use cases. We encourage you to get to know this new feature using our [generics tutorial](https://go.dev/doc/tutorial/generics), and to explore the best ways to use generics to optimize and simplify your code today. The [release notes](https://go.dev/doc/go1.18) have more details about using generics in Go 1.18.

在 Go 1.18 中，我们引入了对使用参数化类型的通用代码的新支持。支持泛型是 Go 最常被要求的功能，我们很自豪能够提供大多数用户目前需要的泛型支持。随后的版本将为一些更复杂的泛型用例提供额外支持。我们鼓励您使用我们的泛型教程来了解这个新功能，并探索使用泛型来优化和简化您的代码的最佳方法。发布说明中有关于在 Go 1.18 中使用泛型的更多细节。

## Fuzzing 模糊处理

With Go 1.18, Go is the first major language with fuzzing fully integrated into its standard toolchain. Like generics, fuzzing has been in design for a long time, and we’re delighted to share it with the Go ecosystem with this release. Please check out our [fuzzing tutorial](https://go.dev/doc/tutorial/fuzz) to help you get started with this new feature.

在 Go 1.18 中，Go 是第一个将模糊处理完全集成到其标准工具链中的主要语言。就像泛型一样，模糊处理在设计上已经有很长时间了，我们很高兴能在这个版本中与 Go 生态系统分享它。请查看我们的模糊处理教程，以帮助您开始使用这个新功能。

## Workspaces 工作区

Go modules have been almost universally adopted, and Go users have reported very high satisfaction scores in our annual surveys. In our 2021 user survey, the most common challenge users identified with modules was working across multiple modules. In Go 1.18, we’ve addressed this with a new [Go workspace mode](https://go.dev/doc/tutorial/workspaces), which makes it simple to work with multiple modules.

Go模块几乎已被普遍采用，Go用户在我们的年度调查中报告了非常高的满意度分数。在我们2021年的用户调查中，用户对模块最常见的挑战是跨多个模块工作。在Go 1.18中，我们通过新的Go工作区模式解决了这一问题，这使得在多个模块中工作变得简单。

## 20% Performance Improvements 20%的性能改进

Apple M1, ARM64, and PowerPC64 users rejoice! Go 1.18 includes CPU performance improvements of up to 20% due to the expansion of Go 1.17’s register ABI calling convention to these architectures. Just to underscore how big this release is, a 20% performance improvement is the fourth most important headline!

苹果M1、ARM64和PowerPC64用户欢欣鼓舞! 由于 Go 1.17 的寄存器 ABI 调用约定扩展到这些架构，Go 1.18 包括 CPU 性能的改进，幅度高达 20%。为了强调这个版本有多大，20%的性能改进是第四个最重要的标题

For a more detailed description of everything that’s in 1.18, please consult the [release notes](https://go.dev/doc/go1.18).

关于1.18中的所有内容的更详细描述，请查阅发布说明。

Go 1.18 is a huge milestone for the entire Go community. We want to thank every Go user who filed a bug, sent in a change, wrote a tutorial, or helped in any way to make Go 1.18 a reality. We couldn’t do it without you. Thank you.

Go 1.18是整个Go社区的一个巨大的里程碑。我们要感谢每一位提交错误、发送修改、编写教程或以任何方式帮助Go 1.18成为现实的Go用户。没有您们，我们无法做到这一点。谢谢您们。

Enjoy Go 1.18!

享受Go 1.18!
