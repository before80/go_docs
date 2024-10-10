+++
title = "Go1_19IsReleased"
weight = 7
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.19 is released! - Go 1.19发布了!

> 原文：[https://go.dev/blog/go1.19](https://go.dev/blog/go1.19)

The Go Team
2 August 2022

Today the Go team is thrilled to release Go 1.19, which you can get by visiting the [download page](https://go.dev/dl/).

今天，Go团队很高兴地发布了Go 1.19，您可以通过访问下载页面获得该版本。

Go 1.19 refines and improves our massive [Go 1.18 release](https://go.dev/blog/go1.18) earlier this year. We focused Go 1.19’s generics development on addressing the subtle issues and corner cases reported to us by the community, as well as important performance improvements (up to 20% for some generic programs).

Go 1.19完善并改进了我们今年早些时候发布的大规模Go 1.18。我们将Go 1.19的泛型开发集中在解决社区向我们报告的细微问题和角落案例，以及重要的性能改进（对一些泛型程序来说，性能改进最高可达20%）。

Doc comments now support [links, lists, and clearer heading syntax](https://go.dev/doc/comment). This change helps users write clearer, more navigable doc comments, especially in packages with large APIs. As part of this change `gofmt` now reformats doc comments to apply a standard formatting to uses of these features. See "[Go Doc Comments](https://go.dev/doc/comment)" for all the details.

文档注释现在支持链接、列表和更清晰的标题语法。这一变化有助于用户写出更清晰、更容易浏览的文档注释，特别是在有大型API的软件包中。作为这一变化的一部分，gofmt 现在对文档注释进行了重新格式化，以便在使用这些功能时采用标准的格式。请参阅 "Go 文档注释 "了解所有细节。

[Go’s memory model](https://go.dev/ref/mem) now explicitly defines the behavior of the [sync/atomic package](https://go.dev/pkg/sync/atomic/). The formal definition of the happens-before relation has been revised to align with the memory models used by C, C++, Java, JavaScript, Rust, and Swift. Existing programs are unaffected. Along with the memory model update, there are [new types in the sync/atomic package](https://go.dev/doc/go1.19#atomic_types), such as [atomic.Int64](https://go.dev/pkg/sync/atomic/#Int64) and [atomic.Pointer[T\]](https://go.dev/pkg/sync/atomic/#Pointer), to make it easier to use atomic values.

Go 的内存模型现在明确定义了 sync/atomic 包的行为。happens-before 关系的正式定义已被修订，以与 C、C++、Java、JavaScript、Rust 和 Swift 使用的内存模型保持一致。现有的程序不受影响。随着内存模型的更新，在sync/atomic包中出现了新的类型，如atomic.Int64和atomic.Pointer[T]，以使使用原子值更加容易。

For [security reasons](https://go.dev/blog/path-security), the os/exec package no longer respects relative paths in PATH lookups. See the [package documentation](https://go.dev/pkg/os/exec/#hdr-Executables_in_the_current_directory) for details. Existing uses of [golang.org/x/sys/execabs](https://pkg.go.dev/golang.org/x/sys/execabs) can be moved back to os/exec in programs that only build using Go 1.19 or later.

出于安全考虑，os/exec包不再尊重PATH查找中的相对路径。详情见包的文档。现有的 golang.org/x/sys/execabs 的使用可以在只使用 Go 1.19 或更高版本构建的程序中移回 os/exec。

The garbage collector has added support for a soft memory limit, discussed in detail in [the new garbage collection guide](https://go.dev/doc/gc-guide#Memory_limit). The limit can be particularly helpful for optimizing Go programs to run as efficiently as possible in containers with dedicated amounts of memory.

垃圾收集器增加了对软内存限制的支持，在新的垃圾收集指南中有详细的讨论。该限制对于优化 Go 程序以尽可能有效地在具有专用内存数量的容器中运行特别有帮助。

The new build constraint `unix` is satisfied when the target operating system (`GOOS`) is any Unix-like system. Today, Unix-like means all of Go’s target operating systems except `js`, `plan9`, `windows`, and `zos`.

当目标操作系统（GOOS）是任何类似Unix的系统时，新的构建约束unix就得到满足。今天，类Unix是指所有Go的目标操作系统，除了js、plan9、windows和zos。

Finally, Go 1.19 includes a wide variety of performance and implementation improvements, including dynamic sizing of initial goroutine stacks to reduce stack copying, automatic use of additional file descriptors on most Unix systems, jump tables for large switch statements on x86-64 and ARM64, support for debugger-injected function calls on ARM64, register ABI support on RISC-V, and experimental support for Linux running on Loongson 64-bit architecture LoongArch (`GOARCH=loong64`).

最后，Go 1.19包括各种各样的性能和实现方面的改进，包括初始goroutine堆栈的动态大小以减少堆栈复制，在大多数Unix系统上自动使用额外的文件描述符，在x86-64和ARM64上为大型switch语句提供跳转表，在ARM64上支持调试器注入的函数调用，在RISC-V上支持寄存器ABI，以及对运行在Loongson 64位架构LoongArch（GOARCH=loong64）的Linux实验性支持。

Thanks to everyone who contributed to this release by writing code, filing bugs, sharing feedback, and testing the beta and release candidates. Your efforts helped to ensure that Go 1.19 is as stable as possible. As always, if you notice any problems, please [file an issue](https://go.dev/issue/new).

感谢所有通过编写代码、提交错误、分享反馈以及测试测试版和候选版而对该版本作出贡献的人。您的努力有助于确保Go 1.19尽可能的稳定。一如既往，如果您发现任何问题，请提交问题。

Enjoy Go 1.19!

享受Go 1.19!
