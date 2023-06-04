+++
title = "go 1.7发布了"
weight = 7
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.7 is released - go 1.7发布了

https://go.dev/blog/go1.7

Chris Broadfoot
15 August 2016

Today we are happy to announce the release of Go 1.7. You can get it from the [download page](https://go.dev/dl/). There are several significant changes in this release: a port for [Linux on IBM z Systems](https://en.wikipedia.org/wiki/IBM_System_z) (s390x), compiler improvements, the addition of the [context](https://go.dev/pkg/context/) package, and support for [hierarchical tests and benchmarks](https://go.dev/pkg/testing/#hdr-Subtests_and_Sub_benchmarks).

今天我们很高兴地宣布Go 1.7的发布。您可以从下载页面获得它。这个版本有几个重要的变化：在IBM z系统（s390x）上的Linux移植，编译器改进，增加了上下文包，并支持分层测试和基准。

A new compiler back end, based on [static single-assignment](https://en.wikipedia.org/wiki/Static_single_assignment_form) form (SSA), has been under development for the past year. By representing a program in SSA form, a compiler may perform advanced optimizations more easily. This new back end generates more compact, more efficient code that includes optimizations like [bounds check elimination](https://en.wikipedia.org/wiki/Bounds-checking_elimination) and [common subexpression elimination](https://en.wikipedia.org/wiki/Common_subexpression_elimination). We observed a 5–35% speedup across our [benchmarks](https://go.dev/test/bench/go1/). For now, the new backend is only available for the 64-bit x86 platform ("amd64"), but we’re planning to convert more architecture backends to SSA in future releases.

一个新的编译器后端，基于静态单一赋值形式（SSA），在过去一年里一直在开发。通过以SSA形式表示一个程序，编译器可以更容易地进行高级优化。这个新的后端生成了更紧凑、更有效的代码，包括边界检查消除和普通子表达式消除等优化。我们观察到，在我们的基准测试中，速度提高了5-35%。目前，新的后端只适用于64位x86平台（"amd64"），但我们计划在未来的版本中将更多的架构后端转换为SSA。

The compiler front end uses a new, more compact export data format, and processes import declarations more efficiently. While these [changes across the compiler toolchain](https://go.dev/doc/go1.7#compiler) are mostly invisible, users have [observed](http://dave.cheney.net/2016/04/02/go-1-7-toolchain-improvements) a significant speedup in compile time and a reduction in binary size by as much as 20–30%.

编译器前端使用一个新的、更紧凑的导出数据格式，并更有效地处理导入声明。虽然这些在整个编译器工具链中的变化大多是不可见的，但用户已经观察到编译时间明显加快，二进制大小减少了20-30%。

Programs should run a bit faster due to speedups in the garbage collector and optimizations in the standard library. Programs with many idle goroutines will experience much shorter garbage collection pauses than in Go 1.6.

由于垃圾收集器的加速和标准库的优化，程序的运行速度应该更快一些。与Go 1.6相比，有许多空闲goroutine的程序将经历更短的垃圾收集暂停时间。

Over the past few years, the [golang.org/x/net/context](https://godoc.org/golang.org/x/net/context/) package has proven to be essential to many Go applications. Contexts are used to great effect in applications related to networking, infrastructure, and microservices (such as [Kubernetes](http://kubernetes.io/) and [Docker](https://www.docker.com/)). They make it easy to enable cancellation, timeouts, and passing request-scoped data. To make use of contexts within the standard library and to encourage more extensive use, the package has been moved from the [x/net](https://godoc.org/golang.org/x/net/context/) repository to the standard library as the [context](https://go.dev/pkg/context/) package. Support for contexts has been added to the [net](https://go.dev/pkg/net/), [net/http](https://go.dev/pkg/net/http/), and [os/exec](https://go.dev/pkg/os/exec/) packages. For more information about contexts, see the [package documentation](https://go.dev/pkg/context) and the Go blog post [*Go Concurrency Patterns: Context*](https://blog.golang.org/context).

在过去的几年中，golang.org/x/net/context包已经被证明是许多Go应用程序的关键。上下文在与网络、基础设施和微服务（如Kubernetes和Docker）相关的应用中发挥了巨大作用。它们使取消、超时和传递请求范围的数据变得容易。为了在标准库中使用上下文，并鼓励更广泛的使用，该包已经从x/net仓库转移到标准库中，成为上下文包。net、net/http 和 os/exec 包都加入了对上下文的支持。关于上下文的更多信息，请参见包的文档和 Go 博客中的 Go 并发模式。Context。

Go 1.5 introduced experimental support for a ["vendor" directory](https://go.dev/cmd/go/#hdr-Vendor_Directories), enabled by the `GO15VENDOREXPERIMENT` environment variable. Go 1.6 enabled this behavior by default, and in Go 1.7, this switch has been removed and the "vendor" behavior is always enabled.

Go 1.5 引入了对 "供应商 "目录的试验性支持，由 GO15VENDOREXPERIMENT 环境变量启用。Go 1.6默认启用了这一行为，而在Go 1.7中，这一开关被移除，"供应商 "行为始终被启用。

Go 1.7 includes many more additions, improvements, and fixes. Find the complete set of changes, and details of the points above, in the [Go 1.7 release notes](https://go.dev/doc/go1.7.html).

Go 1.7 包括了更多的新增内容、改进和修复。在Go 1.7发布说明中可以找到完整的变化，以及以上几点的细节。

Finally, the Go team would like thank everyone who contributed to the release. 170 people contributed to this release, including 140 from the Go community. These contributions ranged from changes to the compiler and linker, to the standard library, to documentation, and code reviews. We welcome contributions; if you’d like to get involved, check out the [contribution guidelines](https://go.dev/doc/contribute.html).

最后，Go团队要感谢所有为该版本做出贡献的人。有170人为这个版本做出了贡献，其中140人来自Go社区。这些贡献包括对编译器和链接器的修改，对标准库的修改，对文档的修改，以及代码审查。我们欢迎大家的贡献；如果您想参与，请查看贡献指南。
