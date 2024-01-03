+++
title = "十四年的 Go"
date = 2024-01-03T21:02:57+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Fourteen Years of Go 十四年的 Go

Russ Cox, for the Go team

​	Russ Cox，Go 团队
10 November 2023

​	2023 年 11 月 10 日

![img](./FourteenYearsOfGo_img/gopherdrink.png)

Today we celebrate the fourteenth birthday of the Go open source release! Go has had a great year, with two feature-filled releases and other important milestones.

​	今天，我们庆祝 Go 开源发布的第十四个生日！Go 经历了伟大的一年，发布了两个功能丰富的版本以及其他重要里程碑。

We released [Go 1.20 in February](https://go.dev/blog/go1.20) and [Go 1.21 in August](https://go.dev/blog/go1.21), focusing more on implementation improvements than new language changes.

​	我们在 2 月发布了 Go 1.20，在 8 月发布了 Go 1.21，重点关注实现改进，而不是新的语言更改。

Profile-guided optimization (PGO), [previewed in Go 1.20](https://go.dev/blog/pgo-preview) and [released in Go 1.21](https://go.dev/blog/pgo), allows the Go compiler to read a profile of your program and then spend more time optimizing the parts of your program that run most often. In Go 1.21, workloads typically get between 2% and 7% CPU usage improvements from enabling PGO. See “[Profile-guided optimization in Go 1.21](https://go.dev/blog/pgo)” for an overview and the [profile-guided optimization user guide](https://go.dev/doc/pgo) for complete documentation.

​	Go 1.20 中预览并在 Go 1.21 中发布的 Profile-guided optimization (PGO) 允许 Go 编译器读取程序的配置文件，然后花费更多时间优化程序中最常运行的部分。在 Go 1.21 中，工作负载通常通过启用 PGO 获得 2% 到 7% 的 CPU 使用率改进。请参阅“Go 1.21 中的 Profile-guided optimization”以获取概述，并参阅 profile-guided optimization 用户指南以获取完整文档。

Go has provided support for gathering coverage profiles during `go test` [since Go 1.2](https://go.dev/blog/cover). Go 1.20 added support for gathering coverage profiles in binaries built by `go build`, allowing you to gather coverage during larger integration tests as well. See “[Code coverage for Go integration tests](https://go.dev/blog/integration-test-coverage)” for details.

​	自 Go 1.2 以来，Go 一直提供在 `go test` 期间收集覆盖率配置文件的支持。Go 1.20 添加了在由 `go build` 构建的二进制文件中收集覆盖率配置文件的支持，从而允许您在更大的集成测试期间收集覆盖率。有关详细信息，请参阅“Go 集成测试的代码覆盖率”。

Compatibility has been an important part of Go since “[Go 1 and the Future of Go Programs](https://go.dev/doc/go1compat)”. Go 1.21 improved compatibility further by expanding the conventions for use of GODEBUG in situations where we need to make a change, such as an important bug fix, that must be permitted but may still break existing programs. See the blog post “[Backward Compatibility, Go 1.21, and Go 2](https://go.dev/blog/compat)” for an overview and the documentation “[Go, Backwards Compatibility, and GODEBUG](https://go.dev/doc/godebug)” for details.

​	自“Go 1 和 Go 程序的未来”以来，兼容性一直是 Go 的重要组成部分。Go 1.21 通过扩展 GODEBUG 的使用惯例进一步改进了兼容性，适用于我们需要进行更改的情况，例如必须允许但仍可能破坏现有程序的重要错误修复。请参阅博文“向后兼容性、Go 1.21 和 Go 2”以获取概述，以及文档“Go、向后兼容性和 GODEBUG”以获取详细信息。

Go 1.21 also shipped support for built-in toolchain management, allowing you to change which version of the Go toolchain you use in a specific module as easily as you change the versions of other dependencies. See the blog post “[Forward Compatibility and Toolchain Management in Go 1.21](https://go.dev/blog/toolchain)” for an overview and the documentation “[Go Toolchains](https://go.dev/doc/toolchain)” for details.

​	Go 1.21 还附带对内置工具链管理的支持，允许您像更改其他依赖项的版本一样轻松地更改在特定模块中使用的 Go 工具链版本。请参阅博文“Go 1.21 中的前向兼容性和工具链管理”以获取概述，以及文档“Go 工具链”以获取详细信息。

Another important tooling achievement was the integration of on-disk indexes into gopls, the Go LSP server. This cut gopls’s startup latency and memory usage by 3-5X in typical use cases. “[Scaling gopls for the growing Go ecosystem](https://go.dev/blog/gopls-scalability)” explains the technical details. You can make sure you’re running the latest gopls by running:

​	另一个重要的工具成就是将磁盘索引集成到 gopls（Go LSP 服务器）中。在典型用例中，这将 gopls 的启动延迟和内存使用量减少了 3-5 倍。“为不断增长的 Go 生态系统扩展 gopls”解释了技术细节。您可以通过运行以下命令确保您运行的是最新的 gopls：

```
go install golang.org/x/tools/gopls@latest
```

Go 1.21 introduced new [cmp](https://go.dev/pkg/cmp/), [maps](https://go.dev/pkg/maps/), and [slices](https://go.dev/pkg/slices/) packages — Go’s first generic standard libraries — as well as expanding the set of comparable types. For details about that, see the blog post “[All your comparable types](https://go.dev/blog/comparable)”.

​	Go 1.21 引入了新的 cmp、maps 和 slices 包（Go 的第一个泛型标准库），并扩展了可比较类型的集合。有关详细信息，请参阅博文“所有可比较类型”。

Overall, we continue to refine generics and to write talks and blog posts explaining important details. Two notable posts this year were “[Deconstructing Type Parameters](https://go.dev/blog/deconstructing-type-parameters)”, and “[Everything You Always Wanted to Know About Type Inference – And a Little Bit More](https://go.dev/blog/type-inference)”.

​	总体而言，我们继续完善泛型，并撰写演讲和博文来解释重要细节。今年有两篇值得注意的博文是“解构类型参数”和“您一直想了解的有关类型推断的一切——以及更多内容”。

Another important new package in Go 1.21 is [log/slog](https://go.dev/pkg/log/slog/), which adds an official API for structured logging to the standard library. See “[Structured logging with slog](https://go.dev/blog/slog)” for an overview.

​	Go 1.21 中的另一个重要新包是 log/slog，它为标准库添加了一个用于结构化日志记录的官方 API。有关概述，请参阅“使用 slog 进行结构化日志记录”。

For the WebAssembly (Wasm) port, Go 1.21 shipped support for running on WebAssembly System Interface (WASI) preview 1. WASI preview 1 is a new “operating system” interface for Wasm that is supported by most server-side Wasm environments. See “[WASI support in Go](https://go.dev/blog/wasi)” for a walkthrough.

​	对于 WebAssembly (Wasm) 端口，Go 1.21 提供了对在 WebAssembly 系统接口 (WASI) 预览版 1 上运行的支持。WASI 预览版 1 是 Wasm 的一个新的“操作系统”接口，受大多数服务器端 Wasm 环境支持。有关演练，请参阅“[Go 中的 WASI 支持](https://go.dev/blog/rebuild)”。

On the security side, we are continuing to make sure Go leads the way in helping developers understand their dependencies and vulnerabilities, with [Govulncheck 1.0 launching in July](https://go.dev/blog/govulncheck). If you use VS Code, you can run govulncheck directly in your editor using the Go extension: see [this tutorial](https://go.dev/doc/tutorial/govulncheck-ide) to get started. And if you use GitHub, you can run govulncheck as part of your CI/CD, with the [GitHub Action for govulncheck](https://github.com/marketplace/actions/golang-govulncheck-action). For more about checking your dependencies for vulnerability problems, see this year’s Google I/O talk, “[Build more secure apps with Go and Google](https://www.youtube.com/watch?v=HSt6FhsPT8c&ab_channel=TheGoProgrammingLanguage)”.)

​	在安全性方面，我们继续确保 Go 在帮助开发者了解其依赖项和漏洞方面处于领先地位，[Govulncheck 1.0 于 7 月发布](https://go.dev/blog/govulncheck)。如果您使用 VS Code，可以使用 Go 扩展直接在编辑器中运行 govulncheck：请参阅[本教程](https://go.dev/doc/tutorial/govulncheck-ide)以开始使用。如果您使用 GitHub，可以使用 [GitHub Action for govulncheck](https://github.com/marketplace/actions/golang-govulncheck-action) 在 CI/CD 中运行 govulncheck。有关检查依赖项是否存在漏洞问题的更多信息，请参阅今年的 Google I/O 演讲“[使用 Go 和 Google 构建更安全的应用](https://www.youtube.com/watch?v=HSt6FhsPT8c&ab_channel=TheGoProgrammingLanguage)”。

Another important security milestone was Go 1.21’s highly reproducible toolchain builds. See “[Perfectly Reproducible, Verified Go Toolchains](https://go.dev/blog/rebuild)” for details, including a demonstration of reproducing an Ubuntu Linux Go toolchain on a Mac without using any Linux tools at all.

​	另一个重要的安全里程碑是 Go 1.21 高度可重现的工具链构建。有关详细信息，请参阅“[完全可重现、经过验证的 Go 工具链](https://go.dev/blog/rebuild)”，其中包括在不使用任何 Linux 工具的情况下在 Mac 上重现 Ubuntu Linux Go 工具链的演示。

It has been a busy year!

​	这一年过得真充实！

In Go’s 15th year, we’ll keep working to make Go the best environment for software engineering at scale. One change we’re particularly excited about is redefining for loop `:=` semantics to remove the potential for accidental aliasing bugs. See “[Fixing For Loops in Go 1.22](https://go.dev/blog/loopvar-preview)” for details, including instructions for previewing this change in Go 1.21.

​	在 Go 的第 15 年，我们将继续努力，让 Go 成为最适合大规模软件工程的环境。我们特别兴奋的一项更改是重新定义 for 循环 `:=` 语义，以消除意外别名错误的可能性。有关详细信息，包括在 Go 1.21 中预览此更改的说明，请参阅“[修复 Go 1.22 中的 For 循环](../FixingForLoopsInGo1_22)”。

## 感谢您！ Thank You! 

The Go project has always been far more than just us on the Go team at Google. Thank you to all our contributors and everyone in the Go community for making Go what it is today. We wish you all the best in the year ahead.

​	Go 项目一直不仅仅是 Google Go 团队的我们。感谢所有贡献者和 Go 社区中的每个人，让 Go 成为今天的模样。我们祝愿大家在新的一年里一切顺利。