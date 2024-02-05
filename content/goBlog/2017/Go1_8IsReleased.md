+++
title = "go 1.8发布了"
weight = 11
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.8 is released - go 1.8发布了

> 原文：[https://go.dev/blog/go1.8](https://go.dev/blog/go1.8)

Chris Broadfoot
16 February 2017

Today the Go team is happy to announce the release of Go 1.8. You can get it from the [download page](https://go.dev/dl/). There are significant performance improvements and changes across the standard library.

今天Go团队很高兴地宣布Go 1.8的发布。您可以从下载页面获得它。整个标准库有明显的性能改进和变化。

The compiler back end introduced in [Go 1.7](https://blog.golang.org/go1.7) for 64-bit x86 is now used on all architectures, and those architectures should see significant [performance improvements](https://go.dev/doc/go1.8#compiler). For instance, the CPU time required by our benchmark programs was reduced by 20-30% on 32-bit ARM systems. There are also some modest performance improvements in this release for 64-bit x86 systems. The compiler and linker have been made faster. Compile times should be improved by about 15% over Go 1.7. There is still more work to be done in this area: expect faster compilation speeds in future releases.

Go 1.7中引入的用于64位x86的编译器后端现在在所有架构上使用，这些架构应该看到明显的性能改进。例如，在32位ARM系统上，我们的基准程序所需的CPU时间减少了20-30%。在这个版本中，对于64位x86系统也有一些适度的性能改进。编译器和链接器的速度变快了。编译时间应该比Go 1.7提高了约15%。在这个领域还有更多的工作要做：期望在未来的版本中能有更快的编译速度。

Garbage collection pauses should be [significantly shorter](https://go.dev/doc/go1.8#gc), usually under 100 microseconds and often as low as 10 microseconds.

垃圾收集的停顿时间应该大大缩短，通常在100微秒以下，经常低至10微秒。

The HTTP server adds support for [HTTP/2 Push](https://go.dev/doc/go1.8#h2push), allowing servers to preemptively send responses to a client. This is useful for minimizing network latency by eliminating roundtrips. The HTTP server also adds support for [graceful shutdown](https://go.dev/doc/go1.8#http_shutdown), allowing servers to minimize downtime by shutting down only after serving all requests that are in flight.

HTTP服务器增加了对HTTP/2推送的支持，允许服务器抢先向客户端发送响应。这对于通过消除往返而最大限度地减少网络延迟是非常有用的。HTTP服务器还增加了对优雅关机的支持，允许服务器在服务完所有正在运行的请求后才关闭，从而最大限度地减少停机时间。

[Contexts](https://go.dev/pkg/context/) (added to the standard library in Go 1.7) provide a cancellation and timeout mechanism. Go 1.8 [adds](https://go.dev/doc/go1.8#more_context) support for contexts in more parts of the standard library, including the [`database/sql`](https://go.dev/pkg/database/sql) and [`net`](https://go.dev/pkg/net) packages and [`Server.Shutdown`](http://beta.golang.org/pkg/net/http/#Server.Shutdown) in the `net/http` package.

Contexts（在Go 1.7中被添加到标准库中）提供了一个取消和超时机制。Go 1.8在标准库的更多部分增加了对上下文的支持，包括数据库/sql和net包以及net/http包的Server.Shutdown。

It’s now much simpler to sort slices using the newly added [`Slice`](https://go.dev/pkg/sort/#Slice) function in the `sort` package. For example, to sort a slice of structs by their `Name` field:

现在使用排序包中新增加的 Slice 函数对切片进行排序要简单得多。例如，通过名称字段对结构的切片进行排序：

```go
sort.Slice(s, func(i, j int) bool { return s[i].Name < s[j].Name })
```

Go 1.8 includes many more additions, improvements, and fixes. Find the complete set of changes, and more information about the improvements listed above, in the [Go 1.8 release notes](https://go.dev/doc/go1.8.html).

Go 1.8 包括更多的新增内容、改进和修复。在 Go 1.8 发行说明中，可以找到完整的变化，以及有关上述改进的更多信息。

To celebrate the release, Go User Groups around the world are holding [release parties](https://github.com/golang/go/wiki/Go-1.8-release-party) this week. Release parties have become a tradition in the Go community, so if you missed out this time, keep an eye out when 1.9 nears.

为了庆祝发布，世界各地的Go用户组将在本周举行发布派对。发布会已经成为Go社区的一个传统，如果您这次错过了，请在1.9来临之际留意。

Thank you to over 200 contributors who helped with this release.

感谢超过200位帮助过这次发布的贡献者。
