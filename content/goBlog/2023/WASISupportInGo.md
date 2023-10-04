+++
title = "Go中的WASI支持"
date = 2023-10-04T14:32:35+08:00
type = "docs"
weight = 81
description = ""
isCJKLanguage = true
draft = false

+++

# WASI support in Go - Go中的WASI支持

> 原文：[https://go.dev/blog/wasi](https://go.dev/blog/wasi)

Johan Brandhorst-Satzkorn, Julien Fabre, Damian Gryski, Evan Phoenix, and Achille Roussel
13 September 2023

Johan Brandhorst-Satzkorn、Julien Fabre、Damian Gryski、Evan Phoenix和Achille Roussel
2023年9月13日

Go 1.21 adds a new port targeting the WASI preview 1 syscall API through the new `GOOS` value `wasip1`. This port builds on the existing WebAssembly port introduced in Go 1.11.

​	Go 1.21通过新的 `GOOS` 值 `wasip1` 添加了对WASI预览1系统调用API的支持。该端口建立在Go 1.11引入的现有WebAssembly端口的基础上。

## 什么是WebAssembly？ What is WebAssembly?

[WebAssembly (Wasm)](https://webassembly.org/) is a binary instruction format originally designed for the web. It represents a standard that allows developers to run high-performance, low-level code directly in web browsers at near-native speeds.

​	[WebAssembly（Wasm）](https://webassembly.org/)是一种最初为Web设计的二进制指令格式。它代表了一种标准，允许开发人员以接近本机速度在Web浏览器中直接运行高性能、低级别的代码。

Go first added support for compiling to Wasm in the 1.11 release, through the `js/wasm` port. This allowed Go code compiled using the Go compiler to be executed in web browsers, but it required a JavaScript execution environment.

​	Go在1.11版本中首次添加了对编译为Wasm的支持，通过 `js/wasm` 端口实现。这允许使用Go编译器编译的Go代码在Web浏览器中执行，但需要一个JavaScript执行环境。

As the use of Wasm has grown, so have use cases outside of the browser. Many cloud providers are now offering services that allow the user to execute Wasm executables directly, leveraging the new [WebAssembly System Interface (WASI)](https://wasi.dev/) syscall API.

​	随着Wasm的使用增加，它在浏览器之外的用例也越来越多。许多云提供商现在提供的服务允许用户直接执行Wasm可执行文件，利用新的[WebAssembly系统接口（WASI）](https://wasi.dev/)系统调用API。

## WebAssembly系统接口 The WebAssembly System Interface

WASI defines a syscall API for Wasm executables, allowing them to interact with system resources such as the filesystem, the system clock, random data utilities, and more. The latest release of the WASI spec is called `wasi_snapshot_preview1`, from which we derive the `GOOS` name `wasip1`. New versions of the API are being developed, and supporting them in the Go compiler in the future will likely mean adding a new `GOOS`.

​	WASI为Wasm可执行文件定义了系统调用API，允许它们与文件系统、系统时钟、随机数据工具等系统资源进行交互。WASI规范的最新版本称为 `wasi_snapshot_preview1` ，我们从中派生出 `GOOS` 名称 `wasip1` 。正在开发新版本的API，未来在Go编译器中支持它们可能意味着添加一个新的 `GOOS` 。

The creation of WASI has allowed a number of Wasm runtimes (hosts) to standardize their syscall API around it. Examples of Wasm/WASI hosts include [Wasmtime](https://wasmtime.dev/), [Wazero](https://wazero.io/), [WasmEdge](https://wasmedge.org/), [Wasmer](https://wasmer.io/), and [NodeJS](https://nodejs.org/). There are also a number of cloud providers offering hosting of Wasm/WASI executables.

​	WASI的创建使得许多Wasm运行时（宿主）可以围绕它标准化其系统调用API。Wasm/WASI宿主的示例包括[Wasmtime](https://wasmtime.dev/)、[Wazero](https://wazero.io/)、[WasmEdge](https://wasmedge.org/)、[Wasmer](https://wasmer.io/)和[NodeJS](https://nodejs.org/)。还有许多云提供商提供Wasm/WASI可执行文件的托管。

## 如何在Go中使用它？ How can we use it with Go?

Make sure that you have installed at least version 1.21 of Go. For this demo, we’ll use [the Wasmtime host](https://docs.wasmtime.dev/cli-install.html) to execute our binary. Let’s start with a simple `main.go`:

​	确保您已安装至少1.21版本的Go。对于此演示，我们将使用[Wasmtime宿主](https://docs.wasmtime.dev/cli-install.html)来执行我们的二进制文件。让我们从一个简单的 `main.go` 开始：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```

We can build it for `wasip1` using the command:

​	我们可以使用以下命令为 `wasip1` 构建它：

```shell
$ GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go
```

This will produce a file, `main.wasm` which we can execute with `wasmtime`:

​	这将生成一个名为 `main.wasm` 的文件，我们可以使用 `wasmtime` 执行它：

```shell
$ wasmtime main.wasm
Hello world!
```

That’s all it takes to get started with Wasm/WASI! You can expect almost all the features of Go to just work with `wasip1`. To learn more about the details of how WASI works with Go, please see [the proposal](https://go.dev/issue/58141).

​	这就是开始使用Wasm/WASI所需的全部！您可以期望几乎所有Go的功能在 `wasip1` 上都能正常工作。要了解有关WASI与Go的详细信息，请参阅[提案](https://go.dev/issue/58141)。

## 使用wasip1运行go测试 Running go tests with wasip1

Building and running a binary is easy, but sometimes we want to be able to run `go test` directly without having to build and execute the binary manually. Similar to the `js/wasm` port, the standard library distribution included in your Go installation comes with a file that makes this very easy. Add the `misc/wasm` directory to your `PATH` when running Go tests and it will run the tests using the Wasm host of your choice. This works by `go test` [automatically executing](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program) `misc/wasm/go_wasip1_wasm_exec` when it finds this file in the `PATH`.

​	构建和运行二进制文件很容易，但有时我们希望能够直接运行 `go test` ，而无需手动构建和执行二进制文件。与 `js/wasm` 端口类似，您安装的Go的标准库分发版中附带了一个文件，使这一点非常容易。在运行Go测试时，将 `misc/wasm` 目录添加到 `PATH` 中，它将使用您选择的Wasm宿主运行测试。这通过在 `PATH` 中找到 `misc/wasm/go_wasip1_wasm_exec` 文件时， `go test` 会[自动执行](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)。

```shell
$ export PATH=$PATH:$(go env GOROOT)/misc/wasm
$ GOOS=wasip1 GOARCH=wasm go test ./...
```

This will run `go test` using Wasmtime. The Wasm host used can be controlled using the environment variable `GOWASIRUNTIME`. Currently supported values for this variable are `wazero`, `wasmedge`, `wasmtime`, and `wasmer`. This script is subject to breaking changes between Go versions. Note that Go `wasip1` binaries don’t execute perfectly on all hosts yet (see [#59907](https://go.dev/issue/59907) and [#60097](https://go.dev/issue/60097)).

​	这将使用Wasmtime运行 `go test` 。可以使用环境变量 `GOWASIRUNTIME` 控制使用的Wasm宿主。目前支持的变量值有 `wazero` 、 `wasmedge` 、 `wasmtime` 和 `wasmer` 。该脚本在Go版本之间可能会有重大变化。请注意，Go的 `wasip1` 二进制文件还不能在所有宿主上完美执行（参见[#59907](https://go.dev/issue/59907)和[#60097](https://go.dev/issue/60097)）。

This functionality also works when using `go run`:

​	在使用 `go run` 时，此功能也适用：

```shell
$ GOOS=wasip1 GOARCH=wasm go run ./main.go
Hello world!
```

## 使用go:wasmimport在Go中包装Wasm函数 Wrapping Wasm functions in Go with go:wasmimport

In addition to the new `wasip1/wasm` port, Go 1.21 introduces a new compiler directive: `go:wasmimport`. It instructs the compiler to translate calls to the annotated function into a call to the function specified by the host module name and function name. This new compiler functionality is what allowed us to define the `wasip1` syscall API in Go to support the new port, but it isn’t limited to being used in the standard library.

​	除了新的 `wasip1/wasm` 端口，Go 1.21还引入了一个新的编译器指令： `go:wasmimport` 。它指示编译器将对带注释的函数的调用转换为对指定的宿主模块名称和函数名称的函数的调用。这个新的编译器功能使我们能够在Go中定义 `wasip1` 系统调用API，以支持新的端口，但它不限于在标准库中使用。

For example, the wasip1 syscall API defines the [`random_get` function](https://github.com/WebAssembly/WASI/blob/a51a66df5b1db01cf9e873f5537bc5bd552cf770/legacy/preview1/docs.md#-random_getbuf-pointeru8-buf_len-size---result-errno), and it is exposed to the Go standard library through [a function wrapper](https://cs.opensource.google/go/go/+/refs/tags/go1.21.0:src/runtime/os_wasip1.go;l=73-75) defined in the runtime package. It looks like this:

​	例如， `wasip1` 系统调用API定义了[ `random_get` 函数](https://github.com/WebAssembly/WASI/blob/a51a66df5b1db01cf9e873f5537bc5bd552cf770/legacy/preview1/docs.md#-random_getbuf-pointeru8-buf_len-size---result-errno)，并且通过运行时包中定义的[函数包装器](https://cs.opensource.google/go/go/+/refs/tags/go1.21.0:src/runtime/os_wasip1.go;l=73-75)暴露给Go标准库。它看起来像这样：

```go
//go:wasmimport wasi_snapshot_preview1 random_get
//go:noescape
func random_get(buf unsafe.Pointer, bufLen size) errno
```

This function wrapper is then wrapped in [a more ergonomic function](https://cs.opensource.google/go/go/+/refs/tags/go1.21.0:src/runtime/os_wasip1.go;l=183-187) for use in the standard library:

​	然后，在标准库中使用更符合人体工程学的函数进行封装：

```go
func getRandomData(r []byte) {
    if random_get(unsafe.Pointer(&r[0]), size(len(r))) != 0 {
        throw("random_get failed")
    }
}
```

This way, a user can call `getRandomData` with a byte slice and it will eventually make its way to the host-defined `random_get` function. In the same way, users can define their own wrappers for host functions.

​	这样，用户可以使用字节切片调用 `getRandomData` ，它最终会传递到宿主定义的 `random_get` 函数。同样，用户可以为宿主函数定义自己的包装器。

To learn more about the intricacies of wrapping Wasm functions in Go, please see [the `go:wasmimport` proposal](https://go.dev/issue/59149).

​	要了解有关在Go中封装Wasm函数的细节，请参阅[ `go:wasmimport` 提案](https://go.dev/issue/59149)。

## 限制 Limitations

While the `wasip1` port passes all standard library tests, there are some notable fundamental limitations of the Wasm architecture that may surprise users.

​	尽管 `wasip1` 端口通过了所有标准库测试，但Wasm架构存在一些值得注意的基本限制，这可能会让用户感到惊讶。

Wasm is a single threaded architecture with no parallelism. The scheduler can still schedule goroutines to run concurrently, and standard in/out/error is non-blocking, so a goroutine can execute while another reads or writes, but any host function calls (such as requesting random data using the example above) will cause all goroutines to block until the host function call has returned.

​	Wasm是一种单线程架构，没有并行性。调度器仍然可以调度并发运行的goroutine，并且标准输入/输出/错误是非阻塞的，因此一个goroutine可以在另一个读取或写入时执行，但是任何宿主函数调用（例如，使用上面的示例请求随机数据）将导致所有goroutine阻塞，直到宿主函数调用返回。

A notable missing feature in the `wasip1` API is a full implementation of network sockets. `wasip1` only defines functions that operate on already opened sockets, making it impossible to support some of the most popular features of the Go standard library, such as HTTP servers. Hosts like Wasmer and WasmEdge implement extensions to the `wasip1` API, allowing the opening of network sockets. While these extensions are not implemented by the Go compiler, there exists a third party library, [`github.com/stealthrocket/net`](https://github.com/stealthrocket/net), which uses `go:wasmimport` to allow the use of `net.Dial` and `net.Listen` on supported Wasm hosts. This enables the creation of `net/http` servers and other network related functionality when using this package.

​	 `wasip1`  API中一个值得注意的缺失功能是对网络套接字的完整实现。 `wasip1` 仅定义了对已打开套接字的操作的函数，这使得不可能支持Go标准库的一些最受欢迎的功能，例如HTTP服务器。像Wasmer和WasmEdge这样的宿主实现了 `wasip1`  API的扩展，允许打开网络套接字。虽然Go编译器没有实现这些扩展，但存在第三方库[ `github.com/stealthrocket/net` ](https://github.com/stealthrocket/net)，它使用 `go:wasmimport` 允许在受支持的Wasm宿主上使用 `net.Dial` 和 `net.Listen` 。这使得在使用此包时可以创建 `net/http` 服务器和其他与网络相关的功能。

## Go中Wasm的未来 The future of Wasm in Go

The addition of the `wasip1/wasm` port is just the beginning of the Wasm capabilities we would like to bring to Go. Please keep an eye out on [the issue tracker](https://github.com/golang/go/issues?q=is%3Aopen+is%3Aissue+label%3Aarch-wasm) for proposals around exporting Go functions to Wasm (`go:wasmexport`), a 32-bit port and future WASI API compatibility.

​	 `wasip1/wasm` 端口的添加只是我们希望为Go带来的Wasm功能的开始。请密切关注[问题跟踪器](https://github.com/golang/go/issues?q=is%3Aopen+is%3Aissue+label%3Aarch-wasm)上关于将Go函数导出到Wasm（ `go:wasmexport` ）、32位端口和未来WASI API兼容性的提案。 

## 参与其中 Get involved

If you are experimenting with and want to contribute to Wasm and Go, please get involved! The Go issue tracker tracks all in-progress work and the #webassembly channel on [the Gophers Slack](https://invite.slack.golangbridge.org/) is a great place to discuss Go and WebAssembly. We look forward to hearing from you!

​	如果您正在尝试并希望为Wasm和Go做出贡献，请参与其中！Go问题跟踪器跟踪所有正在进行的工作，[Gophers Slack](https://invite.slack.golangbridge.org/)上的 #webassembly 频道是讨论Go和WebAssembly的好地方。我们期待您的参与！