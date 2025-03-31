+++
title = "Go 1.24 已发布！"
date = 2025-03-31T14:05:28+08:00
weight = 1000
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/go1.24](https://go.dev/blog/go1.24)

## Go 1.24 is released! - Go 1.24 已发布！

Junyang Shao, on behalf of the Go team

 Junyang Shao，代表 Go 团队

11 February 2025

 2025 年 2 月 11 日

Today the Go team is excited to release Go 1.24, which you can get by visiting the [download page](https://go.dev/dl/).

​	今天，Go 团队很高兴发布 Go 1.24，您可以通过访问 [下载页面](https://go.dev/dl/) 获取。

Go 1.24 comes with many improvements over Go 1.23. Here are some of the notable changes; for the full list, refer to the [release notes](https://go.dev/doc/go1.24).

​	Go 1.24 相较于 Go 1.23 有许多改进。以下是一些显著变化；完整列表请参阅 [发布说明](https://go.dev/doc/go1.24).

## 语言变化 Language changes

Go 1.24 now fully supports [generic type aliases](https://go.dev/issue/46477): a type alias may be parameterized like a defined type. See the [language spec](https://go.dev/ref/spec#Alias_declarations) for details.

​	Go 1.24 现已完全支持 [泛型类型别名](https://go.dev/issue/46477)：类型别名可以像定义类型那样带参数。详细信息请参阅 [语言规范](https://go.dev/ref/spec#Alias_declarations).

## 性能改进 Performance improvements

Several performance improvements in the runtime have decreased CPU overhead by 2–3% on average across a suite of representative benchmarks. These improvements include a new builtin `map` implementation based on [Swiss Tables](https://abseil.io/about/design/swisstables), more efficient memory allocation of small objects, and a new runtime-internal mutex implementation.

​	 运行时的几项性能改进使得在一系列代表性基准测试中，CPU 开销平均降低了 2–3%。这些改进包括基于 [Swiss Tables](https://abseil.io/about/design/swisstables) 的全新内置 `map` 实现、更高效的小对象内存分配，以及全新的运行时内部互斥锁实现.

## 工具改进 Tool improvements

- The `go` command now provides a mechanism for tracking tool dependencies for a module. Use `go get -tool` to add a `tool` directive to the current module. Use `go tool [tool name]` to run the tools declared with the `tool` directive. Read more on the [go command](https://go.dev/doc/go1.24#go-command) in the release notes.
- `go` 命令现在提供了一种跟踪模块工具依赖的机制。使用 `go get -tool` 为当前模块添加 `tool` 指令，使用 `go tool [tool name]` 来运行通过 `tool` 指令声明的工具。更多信息请阅读发布说明中的 [go 命令](https://go.dev/doc/go1.24#go-command).
- The new `test` analyzer in `go vet` subcommand reports common mistakes in declarations of tests, fuzzers, benchmarks, and examples in test packages. Read more on [vet](https://go.dev/doc/go1.24#vet) in the release notes.
- `go vet` 子命令中的新 `test` 分析器会报告测试包中测试、模糊测试、基准测试和示例声明中的常见错误。更多信息请阅读发布说明中的 [vet](https://go.dev/doc/go1.24#vet).

## 标准库新增内容 Standard library additions

- The standard library now includes [a new set of mechanisms to facilitate FIPS 140-3 compliance](https://go.dev/doc/security/fips140). Applications require no source code changes to use the new mechanisms for approved algorithms. Read more on [FIPS 140-3 compliance](https://go.dev/doc/go1.24#fips140) in the release notes. Apart from FIPS 140, several packages that were previously in the [x/crypto](https://go.dev/pkg/golang.org/x/crypto) module are now available in the [standard library](https://go.dev/doc/go1.24#crypto-mlkem).
- 标准库现新增了一套 [促进 FIPS 140-3 合规性的机制](https://go.dev/doc/security/fips140)。使用这些新机制支持已批准的算法，无需修改源代码。更多信息请阅读发布说明中的 [FIPS 140-3 合规性](https://go.dev/doc/go1.24#fips140)。此外，之前位于 [x/crypto](https://go.dev/pkg/golang.org/x/crypto) 模块中的几个包现在已可在 [标准库](https://go.dev/doc/go1.24#crypto-mlkem) 中使用.
- Benchmarks may now use the faster and less error-prone [`testing.B.Loop`](https://go.dev/pkg/testing#B.Loop) method to perform benchmark iterations like `for b.Loop() { ... }` in place of the typical loop structures involving `b.N` like `for range b.N`. Read more on [the new benchmark function](https://go.dev/doc/go1.24#new-benchmark-function) in the release notes.
- 基准测试现在可以使用更快且更不易出错的 [`testing.B.Loop`](https://go.dev/pkg/testing#B.Loop) 方法来执行基准迭代，如 `for b.Loop() { ... }`，以替代传统涉及 `b.N` 的循环结构，如 `for range b.N`。更多信息请阅读发布说明中的 [新基准函数](https://go.dev/doc/go1.24#new-benchmark-function).
- The new [`os.Root`](https://go.dev/pkg/os#Root) type provides the ability to perform filesystem operations isolated under a specific directory. Read more on [filesystem access](https://go.dev/doc/go1.24#directory-limited-filesystem-access) in the release notes.
- 全新的 [`os.Root`](https://go.dev/pkg/os#Root) 类型提供了在特定目录下执行文件系统操作的能力。更多信息请阅读发布说明中的 [文件系统访问](https://go.dev/doc/go1.24#directory-limited-filesystem-access).
- The runtime provides a new finalization mechanism, [`runtime.AddCleanup`](https://go.dev/pkg/runtime#AddCleanup), that is more flexible, more efficient, and less error-prone than [`runtime.SetFinalizer`](https://go.dev/pkg/runtime#SetFinalizer). Read more on [cleanups](https://go.dev/doc/go1.24#improved-finalizers) in the release notes.
- 运行时提供了全新的清理机制 [`runtime.AddCleanup`](https://go.dev/pkg/runtime#AddCleanup)，比 [`runtime.SetFinalizer`](https://go.dev/pkg/runtime#SetFinalizer) 更灵活、更高效且不易出错。更多信息请阅读发布说明中的 [清理](https://go.dev/doc/go1.24#improved-finalizers).

## 改进的 WebAssembly 支持 Improved WebAssembly support

Go 1.24 adds a new `go:wasmexport` directive for Go programs to export functions to the WebAssembly host, and supports building a Go program as a WASI [reactor/library](https://github.com/WebAssembly/WASI/blob/63a46f61052a21bfab75a76558485cf097c0dbba/legacy/application-abi.md#current-unstable-abi). Read more on [WebAssembly](https://go.dev/doc/go1.24#wasm) in the release notes.

​	Go 1.24 新增了 `go:wasmexport` 指令，使 Go 程序可以将函数导出到 WebAssembly 宿主，并支持将 Go 程序构建为 WASI [反应器/库](https://github.com/WebAssembly/WASI/blob/63a46f61052a21bfab75a76558485cf097c0dbba/legacy/application-abi.md#current-unstable-abi)。更多信息请阅读发布说明中的 [WebAssembly](https://go.dev/doc/go1.24#wasm).

------

Please read the [Go 1.24 release notes](https://go.dev/doc/go1.24) for the complete and detailed information. Don’t forget to watch for follow-up blog posts that will go in more depth on some of the topics mentioned here!

​	请阅读 [Go 1.24 发布说明](https://go.dev/doc/go1.24) 以获取完整详细的信息。别忘了关注后续博客文章，它们将更深入地探讨此处提到的一些主题!

Thank you to everyone who contributed to this release by writing code and documentation, reporting bugs, sharing feedback, and testing the release candidates. Your efforts helped to ensure that Go 1.24 is as stable as possible. As always, if you notice any problems, please [file an issue](https://go.dev/issue/new).

​	 感谢所有通过编写代码和文档、报告错误、分享反馈以及测试候选版本为此次发布做出贡献的人。您们的努力帮助确保 Go 1.24 尽可能稳定。如往常一样，如果您发现任何问题，请 [提交问题](https://go.dev/issue/new).

Enjoy Go 1.24!

 祝您享受 Go 1.24!