+++
title = "Go 1.21 发布了！"
date = 2023-08-21T15:01:16+08:00
weight = 90
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Go 1.21 is released! - Go 1.21 发布了！

> 原文：[https://go.dev/blog/go1.21](https://go.dev/blog/go1.21)
>

Eli Bendersky, on behalf of the Go team
8 August 2023

Eli Bendersky，代表Go团队 

2023年8月8日

Today the Go team is thrilled to release Go 1.21, which you can get by visiting the [download page](https://go.dev/dl/).

​	今天，Go团队非常高兴地发布了Go 1.21版本，您可以通过访问[下载页面](https://go.dev/dl/)来获取。

Go 1.21 is packed with new features and improvements. Here are some of the notable changes; for the full list, refer to the [release notes](https://go.dev/doc/go1.21).

​	Go 1.21版本充满了新功能和改进。以下是一些值得注意的变化；有关完整列表，请参阅[发布说明](https://go.dev/doc/go1.21)。

## 工具改进 Tool improvements

- The Profile Guided Optimization (PGO) feature we [announced for preview in 1.20](https://go.dev/blog/pgo-preview) is now generally available! If a file named `default.pgo` is present in the main package’s directory, the `go` command will use it to enable a PGO build. See the [PGO documentation](https://go.dev/doc/pgo) for more details. We’ve measured the impact of PGO on a wide set of Go programs and see performance improvements of 2-7%.
- 我们在1.20版本中[预览发布的概要指导优化（PGO）功能](https://go.dev/blog/pgo-preview)现在已经正式发布！如果主包目录中存在名为 `default.pgo` 的文件，`go` 命令将使用它来启用PGO构建。有关更多详情，请参阅[PGO文档](https://go.dev/doc/pgo)。我们在各种Go程序中测量了PGO的影响，并看到了2-7%的性能提升。
- The [`go` tool](https://go.dev/cmd/go) now supports [backward](https://go.dev/doc/godebug) and [forward](https://go.dev/doc/toolchain) language compatibility.
- [`go` 工具](https://go.dev/cmd/go) 现在支持[向后](https://go.dev/doc/godebug)和[向前](https://go.dev/doc/toolchain)的语言兼容性。

## 语言变更 Language changes

- New built-in functions: [min, max](https://go.dev/ref/spec#Min_and_max) and [clear](https://go.dev/ref/spec#Clear).
- 新的内建函数：[min, max](https://go.dev/ref/spec#Min_and_max) 和 [clear](https://go.dev/ref/spec#Clear)。
- Several improvements to type inference for generic functions. The description of [type inference in the spec](https://go.dev/ref/spec#Type_inference) has been expanded and clarified.
- 对于泛型函数，进行了几处类型推断的改进。规范中[类型推断的描述](https://go.dev/ref/spec#Type_inference)已扩展并明确。
- In a future version of Go we’re planning to address one of the most common gotchas of Go programming: [loop variable capture](https://go.dev/wiki/CommonMistakes). Go 1.21 comes with a preview of this feature that you can enable in your code using an environment variable. See [the LoopvarExperiment wiki page](https://go.dev/wiki/LoopvarExperiment) for more details.
- 在将来的Go版本中，我们计划解决Go编程中最常见的问题之一：[循环变量捕获](https://go.dev/wiki/CommonMistakes)。Go 1.21版本提供了这一特性的预览，您可以通过环境变量在您的代码中启用它。详细信息请参阅[LoopvarExperiment wiki页面](https://go.dev/wiki/LoopvarExperiment)。

## 标准库新增 Standard library additions

- New [log/slog](https://go.dev/pkg/log/slog) package for structured logging.
- 新的 [log/slog](https://go.dev/pkg/log/slog) 包，用于结构化日志记录。
- New [slices](https://go.dev/pkg/slices) package for common operations on slices of any element type. This includes sorting functions that are generally faster and more ergonomic than the [sort](https://go.dev/pkg/sort) package.
- 新的 [slices](https://go.dev/pkg/slices) 包，用于对任何元素类型的切片进行常见操作。其中包括比 [sort](https://go.dev/pkg/sort) 包通常更快且更符合人体工程学的排序函数。
- New [maps](https://go.dev/pkg/maps) package for common operations on maps of any key or element type.
- 新的 [maps](https://go.dev/pkg/maps) 包，用于对任何键或元素类型的映射进行常见操作。
- New [cmp](https://go.dev/pkg/cmp) package with new utilities for comparing ordered values.
- 新的 [cmp](https://go.dev/pkg/cmp) 包，提供用于比较有序值的新实用程序。

## 性能改进 Improved performance

In addition to the performance improvements when enabling PGO:

​	除了启用PGO时的性能改进： 

- The Go compiler itself has been rebuilt with PGO enabled for 1.21, and as a result it builds Go programs 2-4% faster, depending on the host architecture.
- Go编译器本身已使用PGO重新构建，因此构建Go程序的速度提高了2-4%，具体取决于主机架构。
- Due to tuning of the garbage collector, some applications may see up to a 40% reduction in tail latency.
- 由于垃圾回收器的调优，某些应用程序的尾延迟可能减少了高达40%。
- Collecting traces with [runtime/trace](https://go.dev/pkg/runtime/trace) now incurs a substantially smaller CPU cost on amd64 and arm64.
- 使用[runtime/trace](https://go.dev/pkg/runtime/trace)收集跟踪现在在amd64和arm64上的CPU成本大大降低。

## 新的WASI端口 A new port to WASI

Go 1.21 adds an experimental port for [WebAssembly System Interface (WASI)](https://wasi.dev/), Preview 1 (`GOOS=wasip1`, `GOARCH=wasm`).

​	Go 1.21为[WebAssembly系统接口（WASI）](https://wasi.dev/)添加了一个实验性端口，预览1 (`GOOS=wasip1`, `GOARCH=wasm`)。

To facilitate writing more general WebAssembly (Wasm) code, the compiler also supports a new directive for importing functions from the Wasm host: `go:wasmimport`.

​	为了便于编写更通用的WebAssembly（Wasm）代码，编译器还支持从Wasm主机导入函数的新指令：`go:wasmimport`。

------

Thanks to everyone who contributed to this release by writing code, filing bugs, sharing feedback, and testing the release candidates. Your efforts helped to ensure that Go 1.21 is as stable as possible. As always, if you notice any problems, please [file an issue](https://go.dev/issue/new).

​	感谢所有通过编写代码、提交错误、分享反馈和测试发布候选版本来为此版本做出贡献的人。您的努力确保了Go 1.21的稳定性。如有问题，请随时[提交问题](https://go.dev/issue/new)。

Enjoy Go 1.21!

​	尽情享受Go 1.21吧！