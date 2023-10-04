+++
title = "go 1.20 发布了！"
weight = 99
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Go 1.20 is released! - go 1.20 发布了！

> 原文：[https://go.dev/blog/go1.20](https://go.dev/blog/go1.20)https://go.dev/blog/go1.20
>

Robert Griesemer, on behalf of the Go team  - 代表 Go 团队
1 February 2023

2023 年 2 月 1 日

Today the Go team is thrilled to release Go 1.20, which you can get by visiting the [download page](https://go.dev/dl/).

​	今天，Go 团队非常高兴地发布了 Go 1.20 版本，您可以通过访问[下载页面](https://go.dev/dl/)获取该版本。

Go 1.20 benefited from an extended development phase, made possible by earlier broad testing and improved overall stability of the code base.

​	Go 1.20 版本经历了较长的开发阶段，这得益于之前的广泛测试和代码库整体稳定性的提高。

We’re particularly excited to launch a preview of [profile-guided optimization](https://go.dev/doc/pgo) (PGO), which enables the compiler to perform application- and workload-specific optimizations based on run-time profile information. Providing a profile to `go build` enables the compiler to speed up typical applications by around 3–4%, and we expect future releases to benefit even more from PGO. Since this is a preview release of PGO support, we encourage folks to try it out, but there are still rough edges which may preclude production use.

​	我们特别激动地推出了[profile-guided optimization](https://go.dev/doc/pgo)，它使编译器能够根据运行时的配置文件信息执行应用程序和工作负载特定的优化。通过为 `go build` 提供配置文件，编译器可以将典型应用程序的速度提高约 3-4%，我们预计未来的版本将更多地受益于 PGO。由于这是 PGO 支持的预览版本，我们鼓励大家尝试使用，但仍存在一些可能阻止其在生产环境中使用的问题。

Go 1.20 also includes a handful of language changes, many improvements to tooling and the library, and better overall performance.

​	Go 1.20 还包括一些语言更改，对工具和库进行了许多改进，并提高了整体性能。

## 语言变更 - Language changes

- The predeclared [`comparable`](https://go.dev/ref/spec#Type_constraints) constraint is now also [satisfied](https://go.dev/ref/spec#Satisfying_a_type_constraint) by ordinary [comparable types](https://go.dev/ref/spec#Comparison_operators), such as interfaces, which will simplify generic code.
- 预声明的 [`comparable`](https://go.dev/ref/spec#Type_constraints) 约束现在也由普通的[可比类型](https://go.dev/ref/spec#Comparison_operators)，例如接口等，[满足](https://go.dev/ref/spec#Satisfying_a_type_constraint)，这将简化通用代码的编写。
- The functions `SliceData`, `String`, and `StringData` have been added to package [`unsafe`](https://go.dev/ref/spec#Package_unsafe). They complete the set of functions for implementation-independent slice and string manipulation.
- 在 [`unsafe`](https://go.dev/ref/spec#Package_unsafe) 包中新增了函数 `SliceData`、`String` 和 `StringData`，它们完善了与实现无关的切片和字符串操作的函数集。
- Go’s type conversion rules have been extended to permit direct conversion [from a slice to an array](https://go.dev/ref/spec#Conversions_from_slice_to_array_or_array_pointer).
- Go 的类型转换规则已扩展，允许直接将[切片转换为数组](https://go.dev/ref/spec#Conversions_from_slice_to_array_or_array_pointer)。
- The language specification now defines the exact order in which array elements and struct fields are [compared](https://go.dev/ref/spec#Comparison_operators). This clarifies what happens in case of panics during comparisons.
- 语言规范现在定义了数组元素和结构字段的[比较顺序](https://go.dev/ref/spec#Comparison_operators)。这澄清了在比较过程中发生 panic 的情况。

## 工具改进 - Tool improvements

- The [`cover` tool](https://go.dev/testing/coverage) now can collect coverage profiles of whole programs, not just of unit tests.
- [`cover` 工具](https://go.dev/testing/coverage)现在可以收集整个程序的覆盖率profiles，而不仅仅是单元测试的覆盖率。
- The [`go` tool](https://go.dev/cmd/go) no longer relies on pre-compiled standard library package archives in the `$GOROOT/pkg` directory, and they are no longer shipped with the distribution, resulting in smaller downloads. Instead, packages in the standard library are built as needed and cached in the build cache, like other packages.
- [`go` 工具](https://go.dev/cmd/go)不再依赖于 `$GOROOT/pkg` 目录中预编译的标准库包档案，它们也不再随发行版一起提供，从而使下载文件更小。相反，标准库中的包将根据需要构建并缓存在构建缓存中，就像其他包一样。
- The implementation of `go test -json` has been improved to make it more robust in the presence of stray writes to `stdout`.
- `go test -json` 的实现改进，使其在存在对 `stdout` 的杂散写入时更加稳定。
- The `go build`, `go install`, and other build-related commands now accept a `-pgo` flag enabling profile-guided optimizations as well as a `-cover` flag for whole-program coverage analysis.
- `go build`、`go install` 和其他与构建相关的命令现在接受 `-pgo` 标志以启用基于优化指导的优化，并接受 `-cover` 标志以进行整个程序的覆盖分析。
- The `go` command now disables `cgo` by default on systems without a C toolchain. Consequently, when Go is installed on a system without a C compiler, it will now use pure Go builds for packages in the standard library that optionally use cgo, instead of using pre-distributed package archives (which have been removed, as noted above).
- `go` 命令现在在没有 C 工具链的系统上默认禁用 `cgo`。因此，当在没有 C 编译器的系统上安装 Go 时，它现在将使用纯 Go 构建标准库中可选使用 cgo 的包，而不是使用预先分发的包档案（如上所述已被移除）。
- The [`vet` tool](https://go.dev/cmd/vet) reports more loop variable reference mistakes that may occur in tests running in parallel.
- [`vet` 工具](https://go.dev/cmd/vet)报告更多在并行运行的测试中可能发生的循环变量引用错误。

## 标准库新增内容 - Standard library additions

- The new [`crypto/ecdh`](https://go.dev/pkg/crypto/ecdh) package provides explicit support for Elliptic Curve Diffie-Hellman key exchanges over NIST curves and Curve25519.
- 新的 [`crypto/ecdh`](https://go.dev/pkg/crypto/ecdh) 包提供了对 NIST 曲线和 Curve25519 上的椭圆曲线迪菲-赫尔曼密钥交换的显式支持。
- The new function [`errors.Join`](https://go.dev/pkg/errors#Join) returns an error wrapping a list of errors which may be obtained again if the error type implements the `Unwrap() []error` method.
- 新的函数 [`errors.Join`](https://go.dev/pkg/errors#Join) 返回一个错误，其中包装了一系列错误，如果错误类型实现了 `Unwrap() []error` 方法，则可以再次获取这些错误。
- The new [`http.ResponseController`](https://go.dev/pkg/net/http#ResponseController) type provides access to extended per-request functionality not handled by the [`http.ResponseWriter`](https://go.dev/pkg/net/http#ResponseWriter) interface.
- 新的 [`http.ResponseController`](https://go.dev/pkg/net/http#ResponseController) 类型提供了对 [`http.ResponseWriter`](https://go.dev/pkg/net/http#ResponseWriter) 接口未处理的扩展请求功能的访问。
- The [`httputil.ReverseProxy`](https://go.dev/pkg/net/http/httputil#ReverseProxy) forwarding proxy includes a new `Rewrite` hook function, superseding the previous `Director` hook.
- [`httputil.ReverseProxy`](https://go.dev/pkg/net/http/httputil#ReverseProxy) 转发代理包含一个新的 `Rewrite` 钩子函数，取代了以前的 `Director` 钩子函数。
- The new [`context.WithCancelCause`](https://go.dev/pkg/context#WithCancelCause) function provides a way to cancel a context with a given error. That error can be retrieved by calling the new [`context.Cause`](https://go.dev/pkg/context#Cause) function.
- 新的 [`context.WithCancelCause`](https://go.dev/pkg/context#WithCancelCause) 函数提供了一种使用给定错误取消上下文的方式。可以通过调用新的 [`context.Cause`](https://go.dev/pkg/context#Cause) 函数来检索该错误。
- The new [`os/exec.Cmd`](https://go.dev/pkg/os/exec#Cmd) fields [`Cancel`](https://go.dev/pkg/os/exec#Cmd.Cancel) and [`WaitDelay`](https://go.dev/pkg/os/exec#Cmd.WaitDelay) specify the behavior of the `Cmd` when its associated `Context` is canceled or its process exits.
- 新的 [`os/exec.Cmd`](https://go.dev/pkg/os/exec#Cmd) 字段 [`Cancel`](https://go.dev/pkg/os/exec#Cmd.Cancel) 和 [`WaitDelay`](https://go.dev/pkg/os/exec#Cmd.WaitDelay) 指定了 `Cmd` 在其相关联的 `Context` 被取消或其进程退出时的行为。

## 改进的性能 - Improved performance

- Compiler and garbage collector improvements have reduced memory overhead and improved overall CPU performance by up to 2%.
- 编译器和垃圾收集器的改进降低了内存开销，并将整体 CPU 性能提高了最多 2%。
- Work specifically targeting compilation times led to build improvements by up to 10%. This brings build speeds back in line with Go 1.17.
- 针对编译时间的特定工作导致构建时间改进了最多 10%。这使得构建速度与 Go 1.17 保持一致。

When [building a Go release from source](https://go.dev/doc/install/source), Go 1.20 requires a Go 1.17.13 or newer release. In the future, we plan to move the bootstrap toolchain forward approximately once a year. Also, starting with Go 1.21, some older operating systems will no longer be supported: this includes Windows 7, 8, Server 2008 and Server 2012, macOS 10.13 High Sierra, and 10.14 Mojave. On the other hand, Go 1.20 adds experimental support for FreeBSD on RISC-V.

​	在[从源代码构建 Go 发行版](https://go.dev/doc/install/source)时，Go 1.20 需要 Go 1.17.13 或更新的版本。未来，我们计划每年前进一次引导工具链。此外，从 Go 1.21 开始，将不再支持一些较旧的操作系统，包括 Windows 7、8、Server 2008 和 Server 2012，以及 macOS 10.13 High Sierra 和 10.14 Mojave。另一方面，Go 1.20 添加了对 RISC-V 上 FreeBSD 的实验性支持。

For a complete and more detailed list of all changes see the [full release notes](https://go.dev/doc/go1.20).

​	有关所有更改的完整和更详细列表，请参阅[完整的发布说明](https://go.dev/doc/go1.20)。

Thanks to everyone who contributed to this release by writing code, filing bugs, sharing feedback, and testing the release candidates. Your efforts helped to ensure that Go 1.20 is as stable as possible. As always, if you notice any problems, please [file an issue](https://go.dev/issue/new).

​	感谢所有通过编写代码、提交错误、提供反馈和测试发布候选版本而为此版本做出贡献的人。您的努力有助于确保 Go 1.20 尽可能稳定。如果您注意到任何问题，请[提交问题](https://go.dev/issue/new)。

Enjoy Go 1.20!

​	尽情享受 Go 1.20！
