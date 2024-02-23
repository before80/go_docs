+++
title = "Gccgo in GCC 4.7.1"
weight = 4
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Gccgo in GCC 4.7.1

> 原文：[https://go.dev/blog/gccgo-in-gcc-471](https://go.dev/blog/gccgo-in-gcc-471)

Ian Lance Taylor
11 July 2012

The Go language has always been defined by a [spec](https://go.dev/ref/spec), not an implementation. The Go team has written two different compilers that implement that spec: gc and gccgo. Having two different implementations helps ensure that the spec is complete and correct: when the compilers disagree, we fix the spec, and change one or both compilers accordingly. Gc is the original compiler, and the Go tool uses it by default. Gccgo is a different implementation with a different focus, and in this post we’ll take a closer look at it.

Go语言一直是由一个规范来定义的，而不是一个实现。Go团队编写了两种不同的编译器来实现该规范：GCC和Gccgo。有两种不同的实现方式有助于确保规范的完整性和正确性：当编译器出现分歧时，我们会修正规范，并相应地改变一个或两个编译器。Gc是原始的编译器，go工具默认使用它。Gccgo是一个不同的实现，有不同的重点，在这篇文章中，我们将对它进行仔细研究。

Gccgo is distributed as part of GCC, the GNU Compiler Collection. GCC supports several different frontends for different languages; gccgo is a Go frontend connected to the GCC backend. The Go frontend is separate from the GCC project and is designed to be able to connect to other compiler backends, but currently only supports GCC.

Gccgo是作为GCC的一部分发布的，GCC是GNU编译器集合。GCC为不同的语言支持几个不同的前台；gccgo是一个连接到GCC后台的Go前台。Go前台是独立于GCC项目的，被设计为能够连接到其他编译器后端，但目前只支持GCC。

Compared to gc, gccgo is slower to compile code but supports more powerful optimizations, so a CPU-bound program built by gccgo will usually run faster. All the optimizations implemented in GCC over the years are available, including inlining, loop optimizations, vectorization, instruction scheduling, and more. While it does not always produce better code, in some cases programs compiled with gccgo can run 30% faster.

与GCC相比，gccgo编译代码的速度较慢，但支持更强大的优化功能，因此由gccgo构建的由CPU绑定的程序通常会运行得更快。多年来在GCC中实现的所有优化都是可用的，包括内联、循环优化、矢量化、指令调度，等等。虽然它并不总是产生更好的代码，但在某些情况下，用gccgo编译的程序可以运行快30%。

The gc compiler supports only the most popular processors: x86 (32-bit and 64-bit) and ARM. Gccgo, however, supports all the processors that GCC supports. Not all those processors have been thoroughly tested for gccgo, but many have, including x86 (32-bit and 64-bit), SPARC, MIPS, PowerPC and even Alpha. Gccgo has also been tested on operating systems that the gc compiler does not support, notably Solaris.

gc编译器只支持最流行的处理器：X86（32位和64位）和ARM。而Gccgo则支持GCC所支持的所有处理器。并非所有这些处理器都经过了gccgo的彻底测试，但许多处理器都经过了测试，包括x86（32位和64位）、SPARC、MIPS、PowerPC甚至是Alpha。Gccgo还在gc编译器不支持的操作系统上进行了测试，特别是Solaris。

Gccgo provides the standard, complete Go library. Many of the core features of the Go runtime are the same in both gccgo and gc, including the goroutine scheduler, channels, the memory allocator, and the garbage collector. Gccgo supports splitting goroutine stacks as the gc compiler does, but currently only on x86 (32-bit or 64-bit) and only when using the gold linker (on other processors, each goroutine will have a large stack, and a deep series of function calls may run past the end of the stack and crash the program).

Gccgo提供了标准、完整的Go库。Go运行时的许多核心功能在gccgo和gc中都是一样的，包括goroutine调度器、通道、内存分配器和垃圾收集器。Gccgo像gc编译器一样支持分割goroutine堆栈，但目前只在x86（32位或64位）上支持，而且只在使用gold链接器时支持（在其他处理器上，每个goroutine都会有一个大堆栈，一连串的深层函数调用可能会跑过堆栈的末端而使程序崩溃）。

Gccgo distributions do not yet include a version of the go command. However, if you install the go command from a standard Go release, it already supports gccgo via the `-compiler` option: go build `-compiler gccgo myprog`. The tools used for calls between Go and C/C++, cgo and SWIG, also support gccgo.

Gccgo发行版还不包括go命令的版本。然而，如果您从标准Go版本中安装go命令，它已经通过-编译器选项支持gccgo了：go build -compiler gccgo myprog。用于Go和C/C++之间调用的工具，cgo和SWIG，也支持gccgo。

We have put the Go frontend under the same BSD license as the rest of the Go tools. You can download the source code for the frontend at the [gofrontend project](https://github.com/golang/gofrontend). Note that when the Go frontend is linked with the GCC backend to make gccgo, GCC’s GPL license takes precedence.

我们已经把Go前端放在与其他Go工具相同的BSD许可下。您可以在 gofrontend 项目中下载该前端的源代码。请注意，当 Go 前端与 GCC 后端连接以形成 gccgo 时，GCC 的 GPL 许可优先。

The latest release of GCC, 4.7.1, includes gccgo with support for Go 1. If you need better performance for CPU-bound Go programs, or you need to support processors or operating systems that the gc compiler does not support, gccgo might be the answer.

GCC的最新版本4.7.1包括了支持Go 1的gccgo。 如果您需要为绑定CPU的Go程序提供更好的性能，或者您需要支持gc编译器不支持的处理器或操作系统，gccgo可能就是答案了。
