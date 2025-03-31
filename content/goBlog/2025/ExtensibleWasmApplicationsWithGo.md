+++
title = "使用Go的可扩展Wasm 应用程序"
date = 2025-03-31T14:20:54+08:00
weight = 990
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/wasmexport](https://go.dev/blog/wasmexport)

## Extensible Wasm Applications with Go - 使用Go的可扩展Wasm 应用程序 

Cherry Mui

13 February 2025

Go 1.24 enhances its WebAssembly (Wasm) capabilities with the addition of the `go:wasmexport` directive and the ability to build a reactor for WebAssembly System Interface (WASI). These features enable Go developers to export Go functions to Wasm, facilitating better integration with Wasm hosts and expanding the possibilities for Go-based Wasm applications.

​	Go 1.24 通过新增 `go:wasmexport` 指令以及构建 WebAssembly System Interface (WASI) 反应器的能力，增强了其 WebAssembly (Wasm) 功能。这些功能使 Go 开发者能够将 Go 函数导出到 Wasm，从而更好地与 Wasm 宿主集成，并扩展基于 Go 的 Wasm 应用程序的可能性。

## WebAssembly 与 WebAssembly 系统接口 WebAssembly and the WebAssembly System Interface

[WebAssembly (Wasm)](https://webassembly.org/) is a binary instruction format that was initially created for web browsers, providing the execution of high-performance, low-level code at speeds approaching native performance. Since then, Wasm’s utility has expanded, and it is now used in various environments beyond the browser. Notably, cloud providers offer services that directly execute Wasm executables, taking advantage of the [WebAssembly System Interface (WASI)](https://wasi.dev/) system call API. WASI allows these executables to interact with system resources.

​	[WebAssembly (Wasm)](https://webassembly.org/) 是一种二进制指令格式，最初为 Web 浏览器创建，能够以接近原生性能的速度执行高性能、低级代码。此后，Wasm 的用途不断扩展，现在已被用于浏览器之外的各种环境。尤其是云服务提供商提供直接执行 Wasm 可执行文件的服务，利用 [WebAssembly System Interface (WASI)](https://wasi.dev/) 系统调用 API，使这些可执行文件能够与系统资源交互。

Go first added support for compiling to Wasm in the 1.11 release, through the `js/wasm` port. Go 1.21 added a new port targeting the WASI preview 1 syscall API through the new `GOOS=wasip1` port.

​	Go 在 1.11 版本中首次通过 `js/wasm` 端口添加了对编译到 Wasm 的支持。Go 1.21 则通过全新的 `GOOS=wasip1` 端口，新增了一个针对 WASI preview 1 系统调用 API 的端口。

## 使用 `go:wasmexport` 导出 Go 函数到 Wasm - Exporting Go Functions to Wasm with `go:wasmexport`

Go 1.24 introduces a new compiler directive, `go:wasmexport`, which allows developers to export Go functions to be called from outside of the Wasm module, typically from a host application that runs the Wasm runtime. This directive instructs the compiler to make the annotated function available as a Wasm [export](https://webassembly.github.io/spec/core/valid/modules.html?highlight=export#exports) in the resulting Wasm binary.

​	Go 1.24 引入了一个新的编译器指令 `go:wasmexport`，它允许开发者将 Go 函数导出，使其可以从 Wasm 模块外部调用，通常由运行 Wasm 运行时的宿主应用调用。该指令指示编译器将被注解的函数在生成的 Wasm 二进制文件中作为 Wasm [导出](https://webassembly.github.io/spec/core/valid/modules.html?highlight=export#exports) 提供。

To use the `go:wasmexport` directive, simply add it to a function definition:

​	要使用 `go:wasmexport` 指令，只需将其添加到函数定义中：

```go
//go:wasmexport add
func add(a, b int32) int32 { return a + b }
```

With this, the Wasm module will have an exported function named `add` that can be called from the host.

​	这样，Wasm 模块将拥有一个名为 `add` 的导出函数，供宿主调用。

This is analogous to the [cgo `export` directive](https://go.dev/cmd/cgo#hdr-C_references_to_Go), which makes the function available to be called from C, though `go:wasmexport` uses a different, simpler mechanism.

​	这类似于 [cgo `export` 指令](https://go.dev/cmd/cgo#hdr-C_references_to_Go)，它使该函数可供 C 调用，不过 `go:wasmexport` 使用了不同且更简单的机制。

## 构建 WASI 反应器 Building a WASI Reactor

A WASI reactor is a WebAssembly module that operates continuously, and can be called upon multiple times to react on events or requests. Unlike a “command” module, which terminates after its main function finishes, a reactor instance remains live after initialization, and its exports remain accessible.

​	WASI 反应器是一种持续运行的 WebAssembly 模块，可以多次调用以响应事件或请求。与在主函数结束后终止的“命令”模块不同，反应器实例在初始化后仍然保持活跃，其导出内容也始终可访问。

With Go 1.24, one can build a WASI reactor with the `-buildmode=c-shared` build flag.

​	在 Go 1.24 中，可以使用 `-buildmode=c-shared` 构建标志构建 WASI 反应器。

```sh
$ GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o reactor.wasm
```

The build flag signals to the linker not to generate the `_start` function (the entry point for a command module), and instead generate an `_initialize` function, which performs runtime and package initialization, along with any exported functions and their dependencies. The `_initialize` function must be called before any other exported functions. The `main` function will not be automatically invoked.

​	该构建标志指示链接器不要生成 `_start` 函数（命令模块的入口点），而是生成 `_initialize` 函数，该函数执行运行时和包的初始化，以及任何导出函数及其依赖项。必须在调用任何其他导出函数之前调用 `_initialize` 函数。`main` 函数不会被自动调用。

To use a WASI reactor, the host application first initializes it by calling `_initialize`, then simply invoke the exported functions. Here is an example using [Wazero](https://wazero.io/), a Go-based Wasm runtime implementation:

​	要使用 WASI 反应器，宿主应用程序首先通过调用 `_initialize` 进行初始化，然后简单地调用导出函数。下面是一个使用 [Wazero](https://wazero.io/)（基于 Go 的 Wasm 运行时实现）的示例：

```go
// Create a Wasm runtime, set up WASI.
r := wazero.NewRuntime(ctx)
defer r.Close(ctx)
wasi_snapshot_preview1.MustInstantiate(ctx, r)

// Configure the module to initialize the reactor.
config := wazero.NewModuleConfig().WithStartFunctions("_initialize")

// Instantiate the module.
wasmModule, _ := r.InstantiateWithConfig(ctx, wasmFile, config)

// Call the exported function.
fn := wasmModule.ExportedFunction("add")
var a, b int32 = 1, 2
res, _ := fn.Call(ctx, api.EncodeI32(a), api.EncodeI32(b))
c := api.DecodeI32(res[0])
fmt.Printf("add(%d, %d) = %d\n", a, b, c)

// The instance is still alive. We can call the function again.
res, _ = fn.Call(ctx, api.EncodeI32(b), api.EncodeI32(c))
fmt.Printf("add(%d, %d) = %d\n", b, c, api.DecodeI32(res[0]))
```

The `go:wasmexport` directive and the reactor build mode allow applications to be extended by calling into Go-based Wasm code. This is particularly valuable for applications that have adopted Wasm as a plugin or extension mechanism with well-defined interfaces. By exporting Go functions, applications can leverage the Go Wasm modules to provide functionality without needing to recompile the entire application. Furthermore, building as a reactor ensures that the exported functions can be called multiple times without requiring reinitialization, making it suitable for long-running applications or services.

​	`go:wasmexport` 指令和反应器构建模式使得应用程序能够通过调用基于 Go 的 Wasm 代码来扩展功能。这对那些采用 Wasm 作为插件或扩展机制、具有明确定义接口的应用程序尤其有价值。通过导出 Go 函数，应用程序可以利用 Go Wasm 模块提供功能，而无需重新编译整个应用程序。此外，以反应器方式构建确保了导出函数可以被多次调用而无需重新初始化，使其适用于长时间运行的应用程序或服务。

## 支持宿主与客户端之间的丰富类型  Supporting rich types between the host and the client

Go 1.24 also relaxes the constraints on types that can be used as input and result parameters with `go:wasmimport` functions. For example, one can pass a bool, a string, a pointer to an `int32`, or a pointer to a struct which embeds `structs.HostLayout` and contains supported field types (see the [documentation](https://go.dev/cmd/compile#hdr-WebAssembly_Directives) for detail). This allows Go Wasm applications to be written in a more natural and ergonomic way, and removes some unnecessary type conversions.

​	Go 1.24 同时放宽了使用 `go:wasmimport` 函数时可作为输入和输出参数的类型限制。例如，可以传递一个 bool、一个字符串、一个指向 `int32` 的指针，或一个嵌入了 `structs.HostLayout` 并包含支持字段类型的结构体指针（详细信息请参阅 [文档](https://go.dev/cmd/compile#hdr-WebAssembly_Directives)）。这使得用 Go 编写 Wasm 应用程序更加自然、符合人体工学，并且消除了一些不必要的类型转换。

## 限制 Limitations

While Go 1.24 has made significant enhancements to its Wasm capabilities, there are still some notable limitations.

​	虽然 Go 1.24 在 Wasm 能力上做出了显著提升，但仍存在一些明显的限制。

Wasm is a single-threaded architecture with no parallelism. A `go:wasmexport` function can spawn new goroutines. But if a function creates a background goroutine, it will not continue executing when the `go:wasmexport` function returns, until calling back into the Go-based Wasm module.

​	Wasm 是一种单线程架构，没有并行性。一个 `go:wasmexport` 函数可以启动新的 goroutine，但如果一个函数创建了后台 goroutine，当 `go:wasmexport` 函数返回后，该 goroutine 将不会继续执行，直到再次调用基于 Go 的 Wasm 模块。

While some type restrictions have been relaxed in Go 1.24, there are still limitations on the types that can be used with `go:wasmimport` and `go:wasmexport` functions. Due to the unfortunate mismatch between the 64-bit architecture of the client and the 32-bit architecture of the host, it is not possible to pass pointers in memory. For example, a `go:wasmimport` function cannot take a pointer to a struct that contains a pointer-typed field.

​	尽管 Go 1.24 放宽了一些类型限制，但在使用 `go:wasmimport` 和 `go:wasmexport` 函数时，仍存在类型上的限制。由于客户端的 64 位架构与宿主的 32 位架构之间的不匹配，无法在内存中传递指针。例如，一个 `go:wasmimport` 函数不能接受一个包含指针类型字段的结构体指针。

## 结论 Conclusion

The addition of the ability to build a WASI reactor and export Go functions to Wasm in Go 1.24 represent a significant step forward for Go’s WebAssembly capabilities. These features empower developers to create more versatile and powerful Go-based Wasm applications, opening up new possibilities for Go in the Wasm ecosystem.

​	Go 1.24 中新增构建 WASI 反应器和将 Go 函数导出到 Wasm 的能力，代表了 Go 在 WebAssembly 能力上的一大步进。这些特性使开发者能够创建更灵活、更强大的基于 Go 的 Wasm 应用程序，为 Go 在 Wasm 生态系统中开辟了新的可能性.

------

Please read the [Go 1.24 release notes](https://go.dev/doc/go1.24) for the complete and detailed information. Don’t forget to watch for follow-up blog posts that will go in more depth on some of the topics mentioned here!

​	请阅读 [Go 1.24 发布说明](https://go.dev/doc/go1.24) 以获取完整详细的信息。别忘了关注后续博客文章，它们将更深入探讨此处提到的一些主题!

Thank you to everyone who contributed to this release by writing code and documentation, reporting bugs, sharing feedback, and testing the release candidates. Your efforts helped to ensure that Go 1.24 is as stable as possible. As always, if you notice any problems, please [file an issue](https://go.dev/issue/new).

​	感谢所有通过编写代码和文档、报告错误、分享反馈以及测试候选版本为此次发布做出贡献的人。您们的努力帮助确保了 Go 1.24 尽可能稳定。如往常一样，如果您发现任何问题，请 [提交问题](https://go.dev/issue/new).

Enjoy Go 1.24!

​	祝您享受 Go 1.24!