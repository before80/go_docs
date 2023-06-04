+++
title = "go 12 岁了"
weight = 85
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Twelve Years of Go - go 12 岁了

https://go.dev/blog/12years

Russ Cox, for the Go team
10 November 2021

Today we celebrate the twelfth birthday of the Go open source release. We have had an eventful year and have a lot to look forward to next year.

今天我们庆祝Go开源版本的十二岁生日。我们经历了多事之秋，明年也有很多值得期待的事情。

The most visible change here on the blog is our [new home on go.dev](https://go.dev/blog/tidy-web), part of consolidating all our Go web sites into a single, coherent site. Another part of that consolidation was [replacing godoc.org with pkg.go.dev](https://go.dev/blog/godoc.org-redirect).

博客上最明显的变化是我们在go.dev上的新家，这是将我们所有的Go网站整合成一个统一的网站的一部分。整合的另一部分是用pkg.go.dev取代godoc.org。

In February, the [Go 1.16 release](https://go.dev/blog/go1.16) added [macOS ARM64 support](https://go.dev/blog/ports), added [a file system interface](https://go.dev/pkg/io/fs) and [embedded files](https://go.dev/pkg/embed), and [enabled modules by default](https://go.dev/blog/go116-module-changes), along with the usual assortment of improvements and optimizations.

2月，Go 1.16版本增加了对macOS ARM64的支持，增加了文件系统接口和嵌入式文件，并默认启用了模块，同时还进行了一系列的改进和优化。

In August, the [Go 1.17 release](https://go.dev/blog/go1.17) added Windows ARM64 support, made [TLS cipher suite decisions easier and more secure](https://go.dev/blog/tls-cipher-suites), introduced [pruned module graphs](https://go.dev/doc/go1.17#go-command) to make modules even more efficient in large projects, and added [new, more readable build constraint syntax](https://pkg.go.dev/cmd/go#hdr-Build_constraints). Under the hood, Go 1.17 also switched to a register-based calling convention for Go functions on x86-64, improving performance in CPU-bound applications by 5–15%.

8月，Go 1.17版本增加了对Windows ARM64的支持，使TLS密码套件的决定更加简单和安全，引入了修剪模块图，使模块在大型项目中更加有效，并增加了新的、更易读的构建约束语法。在系统内部，Go 1.17还为x86-64上的Go函数切换到了基于寄存器的调用约定，使受CPU约束的应用程序的性能提高了5-15%。

Over the course of the year, we published [many new tutorials](https://go.dev/doc/tutorial/), a [guide to databases in Go](https://go.dev/doc/database/), a [guide to developing modules](https://go.dev/doc/#developing-modules), and a [Go modules reference](https://go.dev/ref/mod). One highlight is the new tutorial "[Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)", which is also available in [interactive form using Google Cloud Shell](https://go.dev/s/cloud-shell-web-tutorial).

在这一年里，我们发布了许多新的教程、Go中的数据库指南、开发模块指南和Go模块参考。其中一个亮点是新的教程 "用Go和Gin开发RESTful API"，该教程也可以通过Google Cloud Shell以互动的形式获得。

We’ve been busy on the IDE side, [enabling gopls by default in VS Code Go](https://go.dev/blog/gopls-vscode-go) and delivering countless improvements to both `gopls` and VS Code Go, including a [powerful debugging experience](https://github.com/golang/vscode-go/blob/master/docs/debugging.md) powered by Delve.

我们在IDE方面一直很忙，在VS Code Go中默认启用了gopls，并对gopls和VS Code Go进行了无数次的改进，包括由Delve提供的强大的调试体验。

We also launched the [Go fuzzing beta](https://go.dev/blog/fuzz-beta) and [officially proposed adding generics to Go](https://go.dev/blog/generics-proposal), both of which are now expected in Go 1.18.

我们还推出了Go模糊测试版，并正式提议在Go中加入泛型，现在这两项都有望在Go 1.18中实现。

Continuing to adapt to "virtual-first", the Go team hosted our second annual [Go day at Google Open Source Live](https://opensourcelive.withgoogle.com/events/go-day-2021). You can watch the talks on YouTube:

为了继续适应 "虚拟优先"，Go团队在Google Open Source Live上举办了我们的第二个年度Go日。您可以在YouTube上观看这些讲座：

- "[Using Generics in Go](https://www.youtube.com/watch?v=nr8EpUO9jhw)", by Ian Lance Taylor, introduces generics and how to use them effectively.Ian Lance Taylor的 "在Go中使用泛型"，介绍了泛型以及如何有效地使用它们。
- "[Modern Enterprise Applications](https://www.youtube.com/watch?v=5fgG1qZaV4w)", by Steve Francia, shows how Go plays a role in enterprise modernization."现代企业应用"，由Steve Francia主讲，展示了Go如何在企业现代化中发挥作用。
- "[Building Better Projects with the Go Editor](https://www.youtube.com/watch?v=jMyzsp2E_0U)", by Suzy Mueller, demonstrates how VS Code Go’s integrated tooling helps you navigate code, debug tests, and more.Suzy Mueller的 "用Go编辑器构建更好的项目"，展示了VS Code Go的集成工具如何帮助您浏览代码、调试测试等。
- "[From Proof of Concept to Production](https://www.youtube.com/watch?v=e7PtBOsTpXE)", by Benjamin Cane, a Distinguished Engineer at American Express, explains how American Express came to use Go for its payments and rewards platforms.美国运通公司的杰出工程师Benjamin Cane的 "从概念验证到生产"，解释了美国运通公司如何在其支付和奖励平台中使用Go。

## Going Forward 向前迈进

We’re incredibly excited about what’s in store for Go’s 13th year. Next month, we will have two talks at [GopherCon 2021](https://www.gophercon.com/), along with [many talented speakers from across the Go community](https://www.gophercon.com/agenda). Register for free and mark your calendars!

我们对Go的第13年的发展感到非常兴奋。下个月，我们将在GopherCon 2021上举办两场讲座，还有许多来自Go社区的天才演讲者。请免费注册，并在您的日历上做个记号。

- "Why and How to Use Go Generics", by Robert Griesemer and Ian Lance Taylor, who led the design and implementation of this new feature."为什么和如何使用Go泛型"，由Robert Griesemer和Ian Lance Taylor主讲，他们领导了这项新功能的设计和实施。
  [Dec 8, 11:00 AM (US Eastern)](https://www.gophercon.com/agenda/session/593015). 12月8日，上午11:00（美国东部）。
- "Debugging Go Code Using the Debug Adapter Protocol (DAP)", by Suzy Mueller, show how to use VS Code Go’s advanced debugging features with Delve."使用调试适配器协议（DAP）调试Go代码"，作者是Suzy Mueller，展示如何使用VS Code Go的高级调试功能与Delve。
  [Dec 9, 3:20 PM (US Eastern)](https://www.gophercon.com/agenda/session/593029). 12月9日，下午3:20（美国东部时间）。

In February, the Go 1.18 release will expand the new register-based calling convention to non-x86 architectures, bringing dramatic performance improvements with it. It will include the new Go fuzzing support. And it will be the first release to include support for generics.

2月，Go 1.18版本将把新的基于寄存器的调用约定扩展到非x86架构，并带来巨大的性能改进。它将包括新的Go模糊支持。这也将是第一个包含泛型支持的版本。

Generics will be one of our focuses for 2022. The initial release in Go 1.18 is only the beginning. We need to spend time using generics and learning what works and what doesn’t, so that we can write best practices and decide what should be added to the standard library and other libraries. We expect that Go 1.19 (expected in August 2022) and later releases will further refine the design and implementation of generics as well as integrating them further into the overall Go experience.

泛型将是我们2022年的重点之一。Go 1.18中的初始版本只是一个开始。我们需要花时间使用泛型，了解哪些是有效的，哪些是无效的，这样我们才能写出最佳实践，并决定哪些应该被添加到标准库和其他库中。我们期望Go 1.19（预计在2022年8月）及以后的版本将进一步完善泛型的设计和实现，并将其进一步整合到整个Go体验中。

Another focus for 2022 is supply chain security. We have been talking for years about the [problems of dependencies](https://research.swtch.com/deps). The design of Go modules provides [reproducible, verifiable, verified builds](https://research.swtch.com/vgo-repro), but there is still more work to be done. Starting in Go 1.18, the `go` command will embed more information in binaries about their build configurations, both to make reproducibility easier and to help projects that need to [generate an SBOM](https://en.wikipedia.org/wiki/Software_bill_of_materials) for Go binaries. We have also started work on a [Go vulnerability database](https://pkg.go.dev/golang.org/x/vuln) and an associated tool to report vulnerabilities in a program’s dependencies. One of our goals in this work is to significantly improve the signal-to-noise ratio of this kind of tool: if a program doesn’t use the vulnerable function, we don’t want to report that. Over the course of 2022 we plan to make this available as a standalone tool but also to add it to existing tooling, including `gopls` and VS Code Go, and to [pkg.go.dev](https://pkg.go.dev/). There is also more to do to improve other aspects of Go’s supply chain security posture. Stay tuned for details.

2022年的另一个重点是供应链安全。我们多年来一直在讨论依赖项的问题。Go模块的设计提供了可复制、可验证、可核实的构建，但仍有更多工作要做。从Go 1.18开始，go命令将在二进制文件中嵌入更多关于其构建配置的信息，这既是为了使可重复性更容易，也是为了帮助那些需要为Go二进制文件生成SBOM的项目。我们也已经开始了Go漏洞数据库和相关工具的工作，以报告程序依赖中的漏洞。我们在这项工作中的目标之一是大幅提高这种工具的信噪比：如果一个程序不使用有漏洞的功能，我们就不想报告。在2022年期间，我们计划将其作为一个独立的工具，但也将其添加到现有的工具中，包括gopls和VS Code Go，以及pkg.go.dev中。还有更多工作要做，以改善Go的供应链安全态势的其他方面。请继续关注细节。

Overall, we expect 2022 to be an eventful year for Go, and we will continue to deliver the timely releases and improvements you’ve come to expect.

总的来说，我们预计2022年将是Go的多事之秋，我们将继续及时发布和改进您所期待的内容。

## Thank You! 谢谢您!

Go is far more than just us on the Go team at Google. Thank you for your help making Go a success and joining us on this adventure. We hope you are all staying safe and wish you all the best.

Go不仅仅是我们谷歌的Go团队的事情。感谢您们帮助Go取得了成功，并加入我们的冒险。我们希望您们都能保持安全，并祝您们一切顺利。
