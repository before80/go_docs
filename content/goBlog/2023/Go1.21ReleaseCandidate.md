+++
title = "go1.21 候选版本"
date = 2023-07-12T20:42:31+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Go 1.21 Release Candidate - go1.21 候选版本

https://go.dev/blog/go1.21rc

Eli Bendersky, on behalf of the Go team 代表 Go 团队
21 June 2023

2023年6月21日

The Go 1.21 first Release Candidate (RC) is available today on the [download page](https://go.dev/dl/#go1.21rc2)! Go 1.21 is packed with new features and improvements. Getting the RC (release candidate) allows you to experiment with it early, try it on your workloads, and report any issues before the final release (scheduled for August). Here are some notable changes and features in Go 1.21; for the full list, refer to the [full release notes](https://tip.golang.org/doc/go1.21).

​	今天发布了 Go 1.21 第一个候选版本（RC）！您可以在[下载页面](https://go.dev/dl/#go1.21rc2)获取 Go 1.21。Go 1.21 带来了许多新功能和改进。获取 RC（候选版本）可以让您尽早进行尝试，在您的工作负载上测试，并在正式发布之前（计划于8月发布）报告任何问题。以下是 Go 1.21 中一些值得注意的变化和功能；完整列表请参阅[完整的发布说明](https://tip.golang.org/doc/go1.21)。

*(Please note that the first RC for Go 1.21 is called `go1.21rc2` because a bug was found and fixed after tagging `go1.21rc1`)*

​	*（请注意，Go 1.21 的第一个候选版本名为 `go1.21rc2`，因为在打标签 `go1.21rc1` 后发现并修复了一个错误）*

## 工具改进 Tool improvements

- The Profile Guided Optimization (PGO) feature we [announced for preview in 1.20](https://go.dev/blog/pgo-preview) is now generally available! If a file named `default.pgo` is present in the main package’s directory, the `go` command will use it to enable a PGO build. See the [PGO documentation](https://go.dev/doc/pgo) for more details. We’ve measured the impact of PGO on a wide set of Go programs and see performance improvements of 2-7%.
- 我们在 1.20 中[宣布的预览版](https://go.dev/blog/pgo-preview)的 Profile Guided Optimization（PGO）功能现已正式可用！如果主包目录中存在名为 `default.pgo` 的文件，`go` 命令将使用它来启用 PGO 构建。有关详细信息，请参阅 [PGO 文档](https://go.dev/doc/pgo)。我们对大量 Go 程序的 PGO 影响进行了评估，发现性能提升为 2-7%。
- The [`go` tool](https://go.dev/cmd/go) now supports [backward](https://tip.golang.org/doc/godebug) and [forward](https://go.dev/doc/toolchain) language compatibility.
- [`go` 工具](https://go.dev/cmd/go)现在支持[向后](https://tip.golang.org/doc/godebug)和[向前](https://go.dev/doc/toolchain)的语言兼容性。

## 语言变更 Language changes

- New built-in functions: [min, max](https://tip.golang.org/ref/spec#Min_and_max) and [clear](https://tip.golang.org/ref/spec#Clear).
- 新增内置函数：[min、max](https://tip.golang.org/ref/spec#Min_and_max)和 [clear](https://tip.golang.org/ref/spec#Clear)。
- Several improvements to type inference for generic functions. The description of [type inference in the spec](https://tip.golang.org/ref/spec#Type_inference) has been expanded and clarified.
- 对泛型函数的类型推断进行了多项改进。规范中关于[类型推断](https://tip.golang.org/ref/spec#Type_inference)的描述已进行了扩展和澄清。
- In a future version of Go we’re planning to address one of the most common gotchas of Go programming: [loop variable capture](https://go.dev/wiki/CommonMistakes). Go 1.21 comes with a preview of this feature that you can enable in your code using an environment variable. See [this LoopvarExperiment wiki page](https://go.dev/wiki/LoopvarExperiment) for more details.
- 在未来的 Go 版本中，我们计划解决 Go 编程中最常见的一个问题：[循环变量捕获](https://go.dev/wiki/CommonMistakes)。Go 1.21 引入了此功能的预览版，您可以使用环境变量在代码中启用它。有关详细信息，请参阅[此 LoopvarExperiment wiki 页面](https://go.dev/wiki/LoopvarExperiment)。

## 标准库新增 Standard library additions

- New [log/slog](https://tip.golang.org/pkg/log/slog) package for structured logging.
- 新增用于结构化日志记录的 [log/slog](https://tip.golang.org/pkg/log/slog) 包。
- New [slices](https://tip.golang.org/pkg/slices) package for common operations on slices of any element type. This includes sorting functions that are generally faster and more ergonomic than the [sort](https://tip.golang.org/pkg/sort) package.
- 新增用于任意元素类型的切片常见操作的 [slices](https://tip.golang.org/pkg/slices) 包。其中包括通常比 [sort](https://tip.golang.org/pkg/sort) 包更快且更易用的排序函数。
- New [maps](https://tip.golang.org/pkg/maps) package for common operations on maps of any key or element type.
- 新增用于任意键或元素类型的映射常见操作的 [maps](https://tip.golang.org/pkg/maps) 包。
- New [cmp](https://tip.golang.org/pkg/cmp) package with new utilities for comparing ordered values.
- 新增 [cmp](https://tip.golang.org/pkg/cmp) 包，提供用于比较有序值的新实用工具。

## 改进性能 Improved performance

In addition to the performance improvements when enabling PGO:

​	除了启用 PGO 时的性能改进之外： 

- The Go compiler itself has been rebuilt with PGO enabled for 1.21, and as a result it builds Go programs 2-4% faster, depending on the host architecture.
- Go 编译器本身已经使用 PGO 重新构建为 1.21 版本，结果是构建 Go 程序的速度提高了 2-4%，具体取决于主机架构。
- Due to tuning of the garbage collector, some applications may see up to a 40% reduction in tail latency.
- 由于垃圾收集器的调整，某些应用程序的尾部延迟可能会减少高达 40%。
- Collecting traces with [runtime/trace](https://pkg.go.dev/runtime/trace) now incurs a substantially smaller CPU cost on amd64 and arm64.
- 在 amd64 和 arm64 上，使用 [runtime/trace](https://pkg.go.dev/runtime/trace) 收集跟踪信息的 CPU 开销显著减少。

## 针对 WASI 的新端口 A new port to WASI

Go 1.21 adds an experimental port for [WebAssembly System Interface (WASI)](https://wasi.dev/), Preview 1 (`GOOS=wasip1`, `GOARCH=wasm`).

​	Go 1.21 添加了一个实验性的 WebAssembly System Interface (WASI) 端口，预览版 1 (`GOOS=wasip1`，`GOARCH=wasm`)。

To facilitate writing more general WebAssembly (WASM) code, the compiler also supports a new directive for importing functions from the WASM host: `go:wasmimport`.

​	为了更方便编写通用的 WebAssembly（WASM）代码，编译器还支持从 WASM 主机导入函数的新指令：`go:wasmimport`。

Please [download the Go 1.21 RC](https://go.dev/dl/#go1.21rc2) and try it! If you notice any problems, please [file an issue](https://go.dev/issue/new).

​	请[下载 Go 1.21 RC](https://go.dev/dl/#go1.21rc2) 并尝试使用！如果您注意到任何问题，请[提交问题](https://go.dev/issue/new)。