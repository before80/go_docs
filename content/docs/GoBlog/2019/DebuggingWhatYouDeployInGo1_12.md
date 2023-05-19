+++
title = "调试你在go 1.12中部署的内容"
weight = 16
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Debugging what you deploy in Go 1.12 - 调试你在go 1.12中部署的内容

https://go.dev/blog/debug-opt

David Chase
21 March 2019

## Introduction 简介

Go 1.11 and Go 1.12 make significant progress toward allowing developers to debug the same optimized binaries that they deploy to production.

Go 1.11 和 Go 1.12 在允许开发人员调试他们部署到生产中的相同优化二进制文件方面取得了重大进展。

As the Go compiler has become increasingly aggressive in producing faster binaries, we’ve lost ground in debuggability. In Go 1.10, users needed to disable optimizations entirely in order to have a good debugging experience from interactive tools like Delve. But users shouldn’t have to trade performance for debuggability, especially when running production services. If your problem is occurring in production, you need to debug it in production, and that shouldn’t require deploying unoptimized binaries.

随着 Go 编译器在生成更快的二进制文件方面变得越来越积极，我们在调试性方面已经失去了优势。在Go 1.10中，用户需要完全禁用优化，以便通过Delve等交互式工具获得良好的调试体验。但是用户不应该用性能来换取调试性，尤其是在运行生产服务时。如果你的问题发生在生产中，你需要在生产中进行调试，而这不应该需要部署未经优化的二进制文件。

For Go 1.11 and 1.12, we focused on improving the debugging experience on optimized binaries (the default setting of the Go compiler). Improvements include

对于 Go 1.11 和 1.12，我们专注于改善优化二进制文件（Go 编译器的默认设置）的调试体验。改进之处包括

- More accurate value inspection, in particular for arguments at function entry;
- More precisely identifying statement boundaries so that stepping is less jumpy and breakpoints more often land where the programmer expects;
- And preliminary support for Delve to call Go functions (goroutines and garbage collection make this trickier than it is in C and C++).
- 更准确的值检查，特别是对函数入口处的参数。
- 更精确地识别语句的边界，使步进不那么跳跃，断点更多地落在程序员期望的地方。
- 初步支持Delve调用Go函数（goroutines和垃圾收集使其比C和C++中更棘手）。

## Debugging optimized code with Delve 用Delve调试优化代码

[Delve](https://github.com/go-delve/delve) is a debugger for Go on x86 supporting both Linux and macOS. Delve is aware of goroutines and other Go features and provides one of the best Go debugging experiences. Delve is also the debugging engine behind [GoLand](https://www.jetbrains.com/go/), [VS Code](https://code.visualstudio.com/), and [Vim](https://github.com/fatih/vim-go).

Delve是一个支持Linux和macOS的x86平台Go的调试器。Delve了解goroutines和其他Go特性，并提供最好的Go调试体验之一。Delve也是GoLand、VS Code和Vim背后的调试引擎。

Delve normally rebuilds the code it is debugging with `-gcflags "all=-N -l"`, which disables inlining and most optimizations. To debug optimized code with delve, first build the optimized binary, then use `dlv exec your_program` to debug it. Or, if you have a core file from a crash, you can examine it with `dlv core your_program your_core`. With 1.12 and the latest Delve releases, you should be able to examine many variables, even in optimized binaries.

Delve通常用-gcflags "all=-N -l "来重建它所调试的代码，这将禁用内联和大多数优化。要用delve调试优化的代码，首先要建立优化的二进制文件，然后用dlv exec your_program来调试它。或者，如果你有一个崩溃的核心文件，你可以用dlv core your_program your_core检查它。有了1.12和最新的Delve版本，你应该能够检查许多变量，甚至在优化的二进制文件中。

## Improved value inspection 改进的值检查

When debugging optimized binaries produced by Go 1.10, variable values were usually completely unavailable. In contrast, starting with Go 1.11, variables can usually be examined even in optimized binaries, unless they’ve been optimized away completely. In Go 1.11 the compiler began emitting DWARF location lists so debuggers can track variables as they move in and out of registers and reconstruct complex objects that are split across different registers and stack slots.

在调试由Go 1.10产生的优化二进制文件时，变量值通常是完全不可用的。相比之下，从Go 1.11开始，即使在优化的二进制文件中，通常也可以检查变量，除非它们被完全优化掉了。在Go 1.11中，编译器开始发出DWARF位置列表，因此调试器可以跟踪变量在寄存器中的进出，并重建分割在不同寄存器和堆栈槽的复杂对象。

## Improved stepping 改进的步进

This shows an example of stepping through a simple function in a debugger in 1.10, with flaws (skipped and repeated lines) highlighted by red arrows.

这是在1.10中调试器步进一个简单函数的例子，缺陷（跳过的和重复的行）用红色箭头标出。

![img](DebuggingWhatYouDeployInGo1_12_img/stepping.svg)

Flaws like this make it easy to lose track of where you are when stepping through a program and interfere with hitting breakpoints.

像这样的缺陷使你在步进程序时很容易失去方向，并影响到断点的判断。

Go 1.11 and 1.12 record statement boundary information and do a better job of tracking source line numbers through optimizations and inlining. As a result, in Go 1.12, stepping through this code stops on every line and does so in the order you would expect.

Go 1.11和1.12记录了语句边界信息，并通过优化和内联在跟踪源码行数方面做得更好。因此，在Go 1.12中，在这段代码中的每一行都会停止，而且是按照你所期望的顺序进行。

## Function calls 函数调用

Function call support in Delve is still under development, but simple cases work. For example:

Delve中的函数调用支持仍在开发中，但简单的情况下也可以。比如说：

```
(dlv) call fib(6)
> main.main() ./hello.go:15 (PC: 0x49d648)
Values returned:
    ~r1: 8
```

## The path forward 前进的道路

Go 1.12 is a step toward a better debugging experience for optimized binaries and we have plans to improve it even further.

Go 1.12是朝着为优化的二进制文件提供更好的调试体验迈出的一步，我们有计划进一步改进它。

There are fundamental tradeoffs between debuggability and performance, so we’re focusing on the highest-priority debugging defects, and working to collect automated metrics to monitor our progress and catch regressions.

在可调试性和性能之间存在着基本的权衡，因此我们将重点放在优先级最高的调试缺陷上，并努力收集自动化指标来监测我们的进展和捕捉回归。

We’re focusing on generating correct information for debuggers about variable locations, so if a variable can be printed, it is printed correctly. We’re also looking at making variable values available more of the time, particularly at key points like call sites, though in many cases improving this would require slowing down program execution. Finally, we’re working on improving stepping: we’re focusing on the order of stepping with panics, the order of stepping around loops, and generally trying to follow source order where possible.

我们专注于为调试器生成关于变量位置的正确信息，所以如果一个变量可以被打印出来，它就会被正确打印出来。我们还在考虑让变量值在更多的时间内可用，特别是在调用点等关键点上，尽管在很多情况下，改善这一点需要减慢程序的执行。最后，我们正在努力改进步进：我们专注于恐慌时的步进顺序，循环时的步进顺序，以及在可能的情况下尝试遵循源代码的顺序。

## A note on macOS support 关于macOS支持的说明

Go 1.11 started compressing debug information to reduce binary sizes. This is natively supported by Delve, but neither LLDB nor GDB support compressed debug info on macOS. If you are using LLDB or GDB, there are two workarounds: build binaries with `-ldflags=-compressdwarf=false`, or use [splitdwarf](https://godoc.org/golang.org/x/tools/cmd/splitdwarf) (`go get golang.org/x/tools/cmd/splitdwarf`) to decompress the debug information in an existing binary.

Go 1.11 开始压缩调试信息以减少二进制文件的大小。Delve原生支持这个功能，但LLDB和GDB在macOS上都不支持压缩的调试信息。如果你使用LLDB或GDB，有两个解决方法：在构建二进制文件时使用-ldflags=-compressdwarf=false，或者使用splitdwarf (go get golang.org/x/tools/cmd/splitdwarf) 来解压现有二进制文件中的调试信息。
