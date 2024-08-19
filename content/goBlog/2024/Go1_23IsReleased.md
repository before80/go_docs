+++
title = "Go 1.23 发布！"
date = 2024-08-19T20:07:23+08:00
weight = 920
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Go 1.23 is released

Dmitri Shuralyov, on behalf of the Go team
13 August 2024

代表 Go 团队，Dmitri Shuralyov
2024年8月13日

Today the Go team is happy to release Go 1.23, which you can get by visiting the [download page](https://go.dev/dl/).

​	今天，Go 团队很高兴发布 Go 1.23，你可以通过访问[下载页面](https://go.dev/dl/)获取它。

If you already have Go 1.22 or Go 1.21 installed on your machine, you can also try `go get toolchain@go1.23.0` in an existing module. This will download the new toolchain and let you begin using it in your module right away. At some later point, you can follow up with `go get go@1.23.0` when you’re ready to fully switch to Go 1.23 and have that be your module’s minimum required Go version. See [Managing Go version module requirements with go get](https://go.dev/doc/toolchain#get) for more information on this functionality.

​	如果你已经在机器上安装了 Go 1.22 或 Go 1.21，也可以在现有模块中尝试使用 `go get toolchain@go1.23.0`。这将下载新的工具链，并让你立即在模块中使用它。稍后，当你准备完全切换到 Go 1.23 并将其设为模块的最低 Go 版本时，可以使用 `go get go@1.23.0`。有关此功能的更多信息，请参阅[使用 go get 管理 Go 版本模块要求]({{< ref "/docs/GoToolchains#使用go-get管理go版本模块要求-managing-go-version-module-requirements-with-go-get">}})。

Go 1.23 comes with many improvements over Go 1.22. Some of the highlights include:

​	Go 1.23 带来了许多相较于 Go 1.22 的改进。以下是一些亮点：

## Language changes 语言变化

- Range expressions in a “for-range” loop may now be iterator functions, such as `func(func(K) bool)`. This supports user-defined iterators over arbitrary sequences. There are several additions to the standard `slices` and `maps` packages that work with iterators, as well as a new `iter` package. As an example, if you wish to collect the keys of a map `m` into a slice and then sort its values, you can do that in Go 1.23 with `slices.Sorted(maps.Keys(m))`.

- “for-range” 循环中的范围表达式现在可以是迭代器函数，例如 `func(func(K) bool)`。这支持用户定义的迭代器遍历任意序列。标准库中的 `slices` 和 `maps` 包添加了多个与迭代器相关的功能，还有一个新的 `iter` 包。例如，如果你想将一个 map `m` 的键收集到一个 slice 中，然后对其值进行排序，你可以在 Go 1.23 中使用 `slices.Sorted(maps.Keys(m))`。

  Go 1.23 also includes preview support for generic type aliases.

  ​	Go 1.23 还包括对泛型类型别名的预览支持。

  Read more about [language changes](https://go.dev/doc/go1.23#language) and [iterators](https://go.dev/doc/go1.23#iterators) in the release notes.

  ​	阅读发行说明，了解更多关于[语言变化](https://go.dev/doc/go1.23#language)和[迭代器](https://go.dev/doc/go1.23#iterators)的内容。

  

## Tool improvements 工具改进

- Starting with Go 1.23, it’s possible for the Go toolchain to collect usage and breakage statistics to help understand how the Go toolchain is used, and how well it is working. This is Go telemetry, an *opt-in system*. Please consider opting in to help us keep Go working well and better understand Go usage. Read more on [Go telemetry](https://go.dev/doc/go1.23#telemetry) in the release notes.
- 从 Go 1.23 开始，Go 工具链可以收集使用和错误统计数据，以帮助理解 Go 工具链的使用情况及其工作效果。这是一个*可选系统*的 Go 遥测功能。请考虑选择加入，以帮助我们保持 Go 的良好运行并更好地了解 Go 的使用情况。阅读发行说明，了解更多关于[Go 遥测](https://go.dev/doc/go1.23#telemetry)的信息。
- The `go` command has new conveniences. For example, running `go env -changed` makes it easier to see only those settings whose effective value differs from the default value, and `go mod tidy -diff` helps determine the necessary changes to the go.mod and go.sum files without modifying them. Read more on the [Go command](https://go.dev/doc/go1.23#go-command) in the release notes.
- `go` 命令有了新的便捷功能。例如，运行 `go env -changed` 可以更容易地看到那些有效值与默认值不同的设置，而 `go mod tidy -diff` 有助于确定对 go.mod 和 go.sum 文件所需的更改，而不对它们进行修改。阅读发行说明，了解更多[Go 命令](https://go.dev/doc/go1.23#go-command)的改进。
- The `go vet` subcommand now reports symbols that are too new for the intended Go version. Read more on [tools](https://go.dev/doc/go1.23#tools) in the release notes.
- `go vet` 子命令现在报告符号是否过新而不适用于目标 Go 版本。阅读发行说明，了解更多[工具](https://go.dev/doc/go1.23#tools)的改进。

## Standard library improvements 标准库改进

- Go 1.23 improves the implementation of `time.Timer` and `time.Ticker`. Read more on [timer changes](https://go.dev/doc/go1.23#timer-changes) in the release notes.
- Go 1.23 改进了 `time.Timer` 和 `time.Ticker` 的实现。阅读发行说明，了解更多关于[计时器变化](https://go.dev/doc/go1.23#timer-changes)的内容。
- There are a total of 3 new packages in the Go 1.23 standard library: `iter`, `structs`, and `unique`. Package `iter` is mentioned above. Package `structs` defines marker types to modify the properties of a struct. Package `unique` provides facilities for canonicalizing (“interning”) comparable values. Read more on [new standard library packages](https://go.dev/doc/go1.23#new-unique-package) in the release notes.
- Go 1.23 标准库中新增了 3 个包：`iter`、`structs` 和 `unique`。前面提到的 `iter` 包。`structs` 包定义了修改结构体属性的标记类型。`unique` 包提供了将可比较值标准化（“内存化”）的工具。阅读发行说明，了解更多关于[新标准库包](https://go.dev/doc/go1.23#new-unique-package)的信息。
- There are many improvements and additions to the standard library enumerated in the [minor changes to the library](https://go.dev/doc/go1.23#minor_library_changes) section of the release notes. The “Go, Backwards Compatibility, and GODEBUG” documentation enumerates [new to Go 1.23 GODEBUG settings](https://go.dev/doc/godebug#go-123).
- 在发行说明的[标准库小变化](https://go.dev/doc/go1.23#minor_library_changes)部分列举了许多对标准库的改进和新增内容。“Go、向后兼容性和 GODEBUG” 文档列举了[Go 1.23 中新增的 GODEBUG 设置](https://go.dev/doc/godebug#go-123)。
- Go 1.23 supports the new `godebug` directive in `go.mod` and `go.work` files to allow separate control of the default GODEBUGs and the “go” directive of `go.mod`, in addition to `//go:debug` directive comments made available two releases ago (Go 1.21). See the updated documentation on [Default GODEBUG Values](https://go.dev/doc/godebug#default).
- Go 1.23 支持在 go.mod 和 go.work 文件中使用新的 `godebug` 指令，以单独控制默认的 GODEBUG 设置和 go.mod 的 “go” 指令，以及在两次发布（Go 1.21）之前引入的 `//go:debug` 指令注释。请参阅更新的[默认 GODEBUG 值](https://go.dev/doc/godebug#default)文档。

## More improvements and changes 更多改进和变化

- Go 1.23 adds experimental support for OpenBSD on 64-bit RISC-V (`openbsd/riscv64`). There are several minor changes relevant to Linux, macOS, ARM64, RISC-V, and WASI. Read more on [ports](https://go.dev/doc/go1.23#ports) in the release notes.
- Go 1.23 增加了对 64 位 RISC-V（`openbsd/riscv64`）上的 OpenBSD 的实验性支持。还有一些与 Linux、macOS、ARM64、RISC-V 和 WASI 相关的小变化。阅读发行说明，了解更多关于[端口](https://go.dev/doc/go1.23#ports)的信息。

- Build time when using profile-guided optimization (PGO) is reduced, and performance with PGO on 386 and amd64 architectures is improved. Read more on [runtime, compiler, and linker](https://go.dev/doc/go1.23#runtime) in the release notes.
- 使用配置文件引导优化（PGO）时的构建时间缩短，并且在 386 和 amd64 架构上的 PGO 性能得到了提高。阅读发行说明，了解更多关于[运行时、编译器和链接器](https://go.dev/doc/go1.23#runtime)的信息。

We encourage everyone to read the [Go 1.23 release notes](https://go.dev/doc/go1.23) for the complete and detailed information on these changes, and everything else that’s new to Go 1.23.

​	我们鼓励大家阅读[Go 1.23 发行说明](https://go.dev/doc/go1.23)，以获取关于这些变化以及 Go 1.23 中所有新增内容的完整详细信息。

Over the next few weeks, look out for follow-up blog posts that will go in more depth on some of the topics mentioned here, including “range-over-func”, the new `unique` package, Go 1.23 timer implementation changes, and more.

​	在接下来的几周内，请留意后续的博客文章，这些文章将更深入地讨论这里提到的一些主题，包括“range-over-func”、新的 `unique` 包、Go 1.23 计时器实现的变化等。

------

Thank you to everyone who contributed to this release by writing code and documentation, reporting bugs, sharing feedback, and testing the release candidates. Your efforts helped to ensure that Go 1.23 is as stable as possible. As always, if you notice any problems, please [file an issue](https://go.dev/issue/new).

​	感谢所有通过编写代码和文档、报告错误、分享反馈和测试候选版本为此次发布做出贡献的人。你们的努力帮助确保了 Go 1.23 尽可能稳定。和往常一样，如果你发现任何问题，请[提交 issue](https://go.dev/issue/new)。

Enjoy Go 1.23!

​	享受 Go 1.23 吧！
