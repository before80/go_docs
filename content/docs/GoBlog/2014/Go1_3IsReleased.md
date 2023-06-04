+++
title = "go 1.3发布了"
weight = 10
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.3 is released - go 1.3发布了

https://go.dev/blog/go1.3

Andrew Gerrand
18 June 2014

Today we are happy to announce the release of [Go 1.3](https://go.dev/doc/go1.3). This release comes six months after our last major release and provides better performance, improved tools, support for running Go in new environments, and more. All Go users should upgrade to Go 1.3. You can grab the release from our [downloads page](https://go.dev/dl/) and find the full list of improvements and fixes in the [release notes](https://go.dev/doc/go1.3). What follows are some highlights.

今天我们很高兴地宣布Go 1.3的发布。这个版本是在我们上次发布主要版本六个月后发布的，它提供了更好的性能、改进的工具、对在新环境中运行Go的支持，以及更多。所有Go用户都应该升级到Go 1.3。您可以从我们的下载页面获取该版本，并在发布说明中找到完整的改进和修复清单。下面是一些亮点。

[Godoc](https://godoc.org/code.google.com/p/go.tools/cmd/godoc), the Go documentation server, now performs static analysis. When enabled with the -analysis flag, analysis results are presented in both the source and package documentation views, making it easier than ever to navigate and understand Go programs. See [the documentation](https://go.dev/lib/godoc/analysis/help.html) for the details.

Go文档服务器Godoc现在可以执行静态分析。当使用-分析标志时，分析结果会同时呈现在源码和包的文档视图中，使得浏览和理解Go程序比以前更容易。详细内容请参见文档。

The gc toolchain now supports the Native Client (NaCl) execution sandbox on the 32- and 64-bit Intel architectures. This permits the safe execution of untrusted code, useful in environments such as the [Playground](https://blog.golang.org/playground). To set up NaCl on your system see the [NativeClient wiki page](https://go.dev/wiki/NativeClient).

gc工具链现在支持32和64位英特尔架构上的本地客户端（NaCl）执行沙盒。这允许安全地执行不受信任的代码，在诸如Playground的环境中非常有用。要在您的系统上设置NaCl，请看NativeClient wiki页面。

Also included in this release is experimental support for the DragonFly BSD, Plan 9, and Solaris operating systems. To use Go on these systems you must [install from source](https://go.dev/doc/install/source).

这个版本还包括对DragonFly BSD、Plan 9和Solaris操作系统的实验性支持。要在这些系统上使用Go，您必须从源码安装。

Changes to the runtime have improved the [performance](https://go.dev/doc/go1.3#performance) of Go binaries, with an improved garbage collector, a new ["contiguous" goroutine stack management strategy](https://go.dev/s/contigstacks), a faster race detector, and improvements to the regular expression engine.

对运行时的修改提高了 Go 二进制文件的性能，改进了垃圾收集器，新的 "连续 "goroutine 堆栈管理策略，更快的竞赛检测器，以及对正则表达式引擎的改进。

As part of the general [overhaul](https://go.dev/s/go13linker) of the Go linker, the compilers and linkers have been refactored. The instruction selection phase that was part of the linker has been moved to the compiler. This can speed up incremental builds for large projects.

作为Go链接器总体大修的一部分，编译器和链接器都进行了重构。原本属于链接器的指令选择阶段已被移至编译器。这可以加快大型项目的增量构建。

The [garbage collector](https://go.dev/doc/go1.3#garbage_collector) is now precise when examining stacks (collection of the heap has been precise since Go 1.1), meaning that a non-pointer value such as an integer will never be mistaken for a pointer and prevent unused memory from being reclaimed. This change affects code that uses package unsafe; if you have unsafe code you should read the [release notes](https://go.dev/doc/go1.3#garbage_collector) carefully to see if your code needs updating.

垃圾收集器在检查堆栈时现在是精确的（从Go 1.1开始，对堆的收集是精确的），这意味着一个非指针值（如整数）绝不会被误认为是指针，并阻止未使用的内存被回收。这一变化影响了使用包unsafe的代码；如果您有不安全的代码，您应该仔细阅读发行说明，看看您的代码是否需要更新。

We would like to thank the many people who contributed to this release; it would not have been possible without your help.

我们要感谢为这个版本做出贡献的许多人；没有您们的帮助，这个版本是不可能完成的。

So, what are you waiting for? Head on over to the [downloads page](https://go.dev/dl/) and start hacking.

那么，您还在等什么呢？快到下载页面开始黑客行动吧。
