+++
title = "Go 的 runtime: 4年之后"
weight = 88
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go runtime: 4 years later - go 的 runtime: 4年之后

> 原文：[https://go.dev/blog/go119runtime](https://go.dev/blog/go119runtime)

Michael Knyszek
26 September 2022

Since our [last blog post about the Go GC in 2018](https://go.dev/blog/ismmkeynote) the Go GC, and the Go runtime more broadly, has been steadily improving. We’ve tackled some large projects, motivated by real-world Go programs and real challenges facing Go users. Let’s catch you up on the highlights!

自从我们在2018年发表关于Go GC的最后一篇博文以来，Go GC，以及更广泛的Go运行时，一直在稳步改进。我们已经处理了一些大型项目，其动机是真实世界的Go程序和Go用户面临的真实挑战。让我们来了解一下其中的亮点吧

### What’s new? 什么是新的？

- `sync.Pool`, a GC-aware tool for reusing memory, has a [lower latency impact](https://go.dev/cl/166960) and [recycles memory much more effectively](https://go.dev/cl/166961) than before. (Go 1.13)sync.Pool，一个有GC意识的重用内存的工具，具有更低的延迟影响，并且比以前更有效地回收内存。(Go 1.13)
- The Go runtime returns unneeded memory back to the operating system [much more proactively](https://go.dev/issue/30333), reducing excess memory consumption and the chance of out-of-memory errors. This reduces idle memory consumption by up to 20%. (Go 1.13 and 1.14)Go运行时更主动地将不需要的内存返回给操作系统，减少了多余的内存消耗和内存外错误的发生。这减少了高达20%的闲置内存消耗。(Go 1.13和1.14)
- The Go runtime is able to preempt goroutines more readily in many cases, reducing stop-the-world latencies up to 90%. [Watch the talk from Gophercon 2020 here.](https://www.youtube.com/watch?v=1I1WmeSjRSw) (Go 1.14)Go运行时在许多情况下能够更容易地抢占goroutines，将stop-the-world的延迟减少到90%。在此观看Gophercon 2020的演讲。(Go 1.14)
- The Go runtime [manages timers more efficiently than before](https://go.dev/cl/171883), especially on machines with many CPU cores. (Go 1.14)Go运行时比以前更有效地管理定时器，特别是在有许多CPU核心的机器上。(Go 1.14)
- Function calls that have been deferred with the `defer` statement now cost as little as a regular function call in most cases. [Watch the talk from Gophercon 2020 here.](https://www.youtube.com/watch?v=DHVeUsrKcbM) (Go 1.14)用defer语句延迟的函数调用，现在在大多数情况下与普通函数调用的成本一样低。在这里观看Gophercon 2020的演讲。(Go 1.14)
- The memory allocator’s slow path [scales](https://go.dev/issue/35112) [better](https://go.dev/issue/37487) with CPU cores, increasing throughput up to 10% and decreasing tail latencies up to 30%, especially in highly-parallel programs. (Go 1.14 and 1.15)内存分配器的慢速路径随着CPU内核的增加而扩展，吞吐量增加了10%，尾部延迟减少了30%，特别是在高度并行的程序中。(Go 1.14 和 1.15)
- Go memory statistics are now accessible in a more granular, flexible, and efficient API, the [runtime/metrics](https://pkg.go.dev/runtime/metrics) package. This reduces latency of obtaining runtime statistics by two orders of magnitude (milliseconds to microseconds). (Go 1.16)Go内存统计数据现在可以通过一个更细化、更灵活、更高效的API，即运行时/指标包来访问。这将获取运行时统计数据的延迟降低了两个数量级（从毫秒到微秒）。(Go 1.16)
- The Go scheduler spends up to [30% less CPU time spinning to find new work](https://go.dev/issue/43997). (Go 1.17)Go调度器花在寻找新工作上的CPU时间减少了30%。(Go 1.17)
- Go code now follows a [register-based calling convention](https://go.dev/issues/40724) on amd64, arm64, and ppc64, improving CPU efficiency by up to 15%. (Go 1.17 and Go 1.18)Go代码现在在amd64、arm64和ppc64上遵循基于寄存器的调用惯例，将CPU效率提高了15%。(Go 1.17 和 Go 1.18)
- The Go GC’s internal accounting and scheduling has been [redesigned](https://go.dev/issue/44167), resolving a variety of long-standing issues related to efficiency and robustness. This results in a significant decrease in application tail latency (up to 66%) for applications where goroutines stacks are a substantial portion of memory use. (Go 1.18)Go GC的内部核算和调度已被重新设计，解决了与效率和健壮性有关的各种长期存在的问题。这使得goroutines堆栈在内存使用中占很大比重的应用程序的尾部延迟大幅下降（最高达66%）。(Go 1.18)
- The Go GC now limits [its own CPU use when the application is idle](https://go.dev/issue/44163). This results in 75% lower CPU utilization during a GC cycle in very idle applications, reducing CPU spikes that can confuse job shapers. (Go 1.19)Go GC 现在限制其在应用程序空闲时的 CPU 使用。这使得非常空闲的应用程序在GC周期内的CPU利用率降低了75%，减少了可能混淆作业整形器的CPU峰值。(Go 1.19)

These changes have been mostly invisible to users: the Go code they’ve come to know and love runs better, just by upgrading Go.

这些变化对用户来说大多是不可见的：他们所熟悉和喜爱的Go代码运行得更好，只是通过升级Go。

### A new knob 一个新的旋钮

With Go 1.19 comes an long-requested feature that requires a little extra work to use, but carries a lot of potential: [the Go runtime’s soft memory limit](https://pkg.go.dev/runtime/debug#SetMemoryLimit).

随着Go 1.19的到来，一个长期以来被要求的功能出现了，它需要一点额外的工作来使用，但具有很大的潜力：Go运行时间的软内存限制。

For years, the Go GC has had only one tuning parameter: `GOGC`. `GOGC` lets the user adjust [the trade-off between CPU overhead and memory overhead made by the Go GC](https://pkg.go.dev/runtime/debug#SetGCPercent). For years, this "knob" has served the Go community well, capturing a wide variety of use-cases.

多年来，Go GC 只有一个调整参数。GOGC。GOGC让用户调整Go GC在CPU开销和内存开销之间的权衡。多年来，这个 "旋钮 "为Go社区提供了很好的服务，捕捉到了各种各样的使用情况。

The Go runtime team has been reluctant to add new knobs to the Go runtime, with good reason: every new knob represents a new *dimension* in the space of configurations that we need to test and maintain, potentially forever. The proliferation of knobs also places a burden on Go developers to understand and use them effectively, which becomes more difficult with more knobs. Hence, the Go runtime has always leaned into behaving reasonably with minimal configuration.

Go运行时团队一直不愿意在Go运行时中添加新的旋钮，这是有原因的：每一个新的旋钮都代表了我们需要测试和维护的配置空间的一个新维度，而且可能是永远。旋钮的激增也给Go开发者带来了理解和有效使用它们的负担，随着旋钮的增多，这也变得更加困难。因此，Go运行时一直倾向于在最小配置的情况下合理行事。

So why add a memory limit knob?

那么为什么要增加一个内存限制旋钮呢？

Memory is not as fungible as CPU time. With CPU time, there’s always more of it in the future, if you just wait a bit. But with memory, there’s a limit to what you have.

内存并不像CPU时间那样可以被替换。对于CPU时间，只要您稍加等待，未来总会有更多的时间。但是对于内存来说，您所拥有的东西是有限制的。

The memory limit solves two problems.

内存限制解决了两个问题。

The first is that when the peak memory use of an application is unpredictable, `GOGC` alone offers virtually no protection from running out of memory. With just `GOGC`, the Go runtime is simply unaware of how much memory it has available to it. Setting a memory limit enables the runtime to be robust against transient, recoverable load spikes by making it aware of when it needs to work harder to reduce memory overhead.

第一个问题是，当一个应用程序的内存使用峰值是不可预测的时候，单靠GOGC几乎不能提供保护，以免内存耗尽。只有GOGC，Go运行时根本不知道它有多少可用的内存。设置内存限制可以使运行时对瞬时的、可恢复的负载高峰保持稳健，让它知道何时需要更努力地工作以减少内存开销。

The second is that to avoid out-of-memory errors without using the memory limit, `GOGC` must be tuned according to peak memory, resulting in higher GC CPU overheads to maintain low memory overheads, even when the application is not at peak memory use and there is plenty of memory available. This is especially relevant in our containerized world, where programs are placed in boxes with specific and isolated memory reservations; we might as well make use of them! By offering protection from load spikes, setting a memory limit allows for `GOGC` to be tuned much more aggressively with respect to CPU overheads.

其次，为了避免在不使用内存限制的情况下出现内存不足的错误，GOGC必须根据峰值内存进行调整，从而导致更高的GC CPU开销，以维持低内存开销，即使应用程序不处于内存使用的峰值，并且有大量的内存可用。这在我们的容器化世界中尤其重要，程序被放置在具有特定和隔离的内存预留的盒子里；我们不妨利用它们！"。通过提供对负载高峰的保护，设置内存限制允许GOGC在CPU开销方面进行更积极的调整。

The memory limit is designed to be easy to adopt and robust. For example, it’s a limit on the whole memory footprint of the Go parts of an application, not just the Go heap, so users don’t have to worry about accounting for Go runtime overheads. The runtime also adjusts its memory scavenging policy in response to the memory limit so it returns memory to the OS more proactively in response to memory pressure.

内存限制被设计为易于采用和稳健。例如，它是对应用程序中Go部分的整个内存足迹的限制，而不仅仅是对Go堆的限制，因此用户不必担心考虑Go运行时的开销问题。运行时还会根据内存限制调整其内存清扫策略，以便在应对内存压力时更主动地将内存返回给操作系统。

But while the memory limit is a powerful tool, it must still be used with some care. One big caveat is that it opens up your program to GC thrashing: a state in which a program spends too much time running the GC, resulting in not enough time spent making meaningful progress. For example, a Go program might thrash if the memory limit is set too low for how much memory the program actually needs. GC thrashing is something that was unlikely previously, unless `GOGC` was explicitly tuned heavily in favor of memory use. We chose to favor running out of memory over thrashing, so as a mitigation, the runtime will limit the GC to 50% of total CPU time, even if this means exceeding the memory limit.

但是，尽管内存限制是一个强大的工具，它仍然必须被谨慎地使用。一个很大的问题是，它使您的程序有可能出现GC thrashing：一种状态，即程序花太多时间运行GC，导致没有足够的时间取得有意义的进展。例如，如果内存限制设置得太低，与程序实际需要的内存相比，Go程序可能会发生激动。除非GOGC被明确地调整为有利于内存的使用，否则GC激动是以前不太可能发生的。我们选择了倾向于耗尽内存而不是激动，所以作为一种缓解措施，运行时将把GC限制在总CPU时间的50%，即使这意味着超过内存限制。

All of this is a lot to consider, so as a part of this work, we released [a shiny new GC guide](https://go.dev/doc/gc-guide), complete with interactive visualizations to help you understand GC costs and how to manipulate them.

所有这些都是需要考虑的，所以作为这项工作的一部分，我们发布了一个闪亮的新的GC指南，其中包括互动的可视化，以帮助您理解GC成本和如何操作它们。

### Conclusion 结论

Try out the memory limit! Use it in production! Read the [GC guide](https://go.dev/doc/gc-guide)!

尝试一下内存限制! 在生产中使用它! 阅读GC指南!

We’re always looking for feedback on how to improve Go, but it also helps to hear about when it just works for you. [Send us feedback](https://groups.google.com/g/golang-dev)!

我们一直在寻找关于如何改进 Go 的反馈，但也希望能听到关于它对您有用的信息。请给我们发送反馈意见!
