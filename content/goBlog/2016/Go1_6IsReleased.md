+++
title = "Go 1.6发布了"
weight = 8
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.6 is released - Go 1.6发布了

> 原文：[https://go.dev/blog/go1.6](https://go.dev/blog/go1.6)

Andrew Gerrand
17 February 2016

Today we release [Go version 1.6](https://go.dev/doc/go1.6), the seventh major stable release of Go. You can grab it right now from the [download page](https://go.dev/dl/). Although [the release of Go 1.5](https://blog.golang.org/go1.5) six months ago contained dramatic implementation changes, this release is more incremental.

今天我们发布了Go的1.6版本，这是Go的第七个主要稳定版本。您现在就可以从下载页面抓取它。虽然6个月前发布的Go 1.5版本包含了巨大的实现变化，但这个版本更多是增量的。

The most significant change is support for [HTTP/2](https://http2.github.io/) in the [net/http package](https://go.dev/pkg/net/http/). HTTP/2 is a new protocol, a follow-on to HTTP that has already seen widespread adoption by browser vendors and major websites. In Go 1.6, support for HTTP/2 is [enabled by default](https://go.dev/doc/go1.6#http2) for both servers and clients when using HTTPS, bringing [the benefits](https://http2.github.io/faq/) of the new protocol to a wide range of Go projects, such as the popular [Caddy web server](https://caddyserver.com/download).

最重要的变化是对net/http包中的HTTP/2的支持。HTTP/2是一个新协议，是HTTP的后续协议，已经被浏览器供应商和主要网站广泛采用。在Go 1.6中，当使用HTTPS时，服务器和客户端都默认启用对HTTP/2的支持，将新协议的好处带给广泛的Go项目，如流行的Caddy网络服务器。

The template packages have learned some new tricks, with support for [trimming spaces around template actions](https://go.dev/pkg/text/template/#hdr-Text_and_spaces) to produce cleaner template output, and the introduction of the [`{{block}}` action](https://go.dev/pkg/text/template/#hdr-Actions) that can be used to create templates that build on other templates. A [new template example program](https://cs.opensource.google/go/x/example/+/master:template) demonstrates these new features.

模板包学会了一些新的技巧，支持修剪模板动作周围的空格，以产生更干净的模板输出，并引入了{{block}}动作，可用于创建建立在其他模板上的模板。一个新的模板示例程序展示了这些新功能。

Go 1.5 introduced [experimental support](https://go.dev/s/go15vendor) for a "vendor" directory that was enabled by an environment variable. In Go 1.6, the feature is now [enabled by default](https://go.dev/doc/go1.6#go_command). Source trees that contain a directory named "vendor" that is not used in accordance with the new feature will require changes to avoid broken builds (the simplest fix is to rename the directory).

Go 1.5引入了对 "供应商 "目录的试验性支持，通过环境变量启用。在 Go 1.6 中，该功能现在是默认启用的。包含名为 "vendor "的目录的源代码树，如果没有按照新功能使用，则需要进行修改以避免构建失败（最简单的修复方法是重命名该目录）。

The runtime has added lightweight, best-effort detection of concurrent misuse of maps. As always, if one goroutine is writing to a map, no other goroutine should be reading or writing the map concurrently. If the runtime detects this condition, it prints a diagnosis and crashes the program. The best way to find out more about the problem is to run it under the [race detector](https://blog.golang.org/race-detector), which will more reliably identify the race and give more detail.

运行时增加了轻量级的、尽最大努力的对地图的并发误用的检测。像往常一样，如果一个goroutine正在向一个地图写东西，其他goroutine不应该同时读或写这个地图。如果运行时检测到这种情况，它就会打印出诊断书并使程序崩溃。发现问题的最好方法是在竞赛检测器下运行，这将更可靠地识别竞赛并给出更多细节。

The runtime has also changed how it prints program-ending panics. It now prints only the stack of the panicking goroutine, rather than all existing goroutines. This behavior can be configured using the [GOTRACEBACK](https://go.dev/pkg/runtime/#hdr-Environment_Variables) environment variable or by calling the [debug.SetTraceback](https://go.dev/pkg/runtime/debug/#SetTraceback) function.

运行时也改变了它打印程序结束的恐慌的方式。它现在只打印恐慌的goroutine的堆栈，而不是所有现有的goroutine。这种行为可以通过 GOTRACEBACK 环境变量或调用 debug.SetTraceback 函数来配置。

Users of cgo should be aware of major changes to the rules for sharing pointers between Go and C code. The rules are designed to ensure that such C code can coexist with Go’s garbage collector and are checked during program execution, so code may require changes to avoid crashes. See the [release notes](https://go.dev/doc/go1.6#cgo) and [cgo documentation](https://go.dev/cmd/cgo/#hdr-Passing_pointers) for the details.

cgo的用户应该注意在Go和C代码之间共享指针规则的重大变化。这些规则旨在确保此类 C 代码能够与 Go 的垃圾收集器共存，并在程序执行期间进行检查，因此代码可能需要修改以避免崩溃。详情请参见发行说明和cgo文档。

The compiler, linker, and go command have a new `-msan` flag analogous to `-race` and only available on linux/amd64, that enables interoperation with the [Clang MemorySanitizer](http://clang.llvm.org/docs/MemorySanitizer.html). This is useful for testing a program containing suspect C or C++ code. You might like to try it while testing your cgo code with the new pointer rules.

编译器、链接器和 Go 命令有一个新的 -msan 标志，类似于 -race，仅在 linux/amd64 上可用，它可以与 Clang MemorySanitizer 进行互操作。这对于测试包含可疑的C或C++代码的程序很有用。在用新的指针规则测试您的cgo代码时，您可能想试试它。

Performance of Go programs built with Go 1.6 remains similar to those built with Go 1.5. Garbage-collection pauses are even lower than with Go 1.5, but this is particularly noticeable for programs using large amounts of memory. With regard to the performance of the compiler tool chain, build times should be similar to those of Go 1.5.

用 Go 1.6 构建的 Go 程序的性能仍然与用 Go 1.5 构建的程序相似。垃圾收集的暂停时间比Go 1.5还要低，但这对于使用大量内存的程序来说尤其明显。关于编译器工具链的性能，构建时间应该与Go 1.5的类似。

The algorithm inside [sort.Sort](https://go.dev/pkg/sort/#Sort) was improved to run about 10% faster, but the change may break programs that expect a specific ordering of equal but distinguishable elements. Such programs should refine their `Less` methods to indicate the desired ordering or use [sort.Stable](https://go.dev/pkg/sort/#Stable) to preserve the input order for equal values.

sort.Sort里面的算法经过改进，运行速度提高了10%左右，但这个变化可能会破坏那些期望对相等但可区分的元素进行特定排序的程序。这样的程序应该完善他们的Less方法，以表明所需的排序，或者使用sort.Stable来保留等值的输入顺序。

And, of course, there are many more additions, improvements, and fixes. You can find them all in the comprehensive [release notes](https://go.dev/doc/go1.6).

当然，还有更多的补充、改进和修正。您可以在全面的发布说明中找到它们。

To celebrate the release, [Go User Groups around the world](https://github.com/golang/go/wiki/Go-1.6-release-party) are holding release parties on the 17th of February. Online, the Go contributors are hosting a question and answer session on the [golang subreddit](https://reddit.com/r/golang) for the next 24 hours. If you have questions about the project, the release, or just Go in general, then please [join the discussion](https://www.reddit.com/r/golang/comments/46bd5h/ama_we_are_the_go_contributors_ask_us_anything/).

为了庆祝该版本，世界各地的Go用户组将在2月17日举行发布派对。在网上，Go贡献者们正在golang subreddit上主持未来24小时的问答会议。如果您有关于项目、发布的问题，或者只是关于Go的一般问题，那么请加入讨论。

Thanks to everyone that contributed to the release. Happy hacking.

感谢所有为该版本做出贡献的人。快乐的黑客行为。
