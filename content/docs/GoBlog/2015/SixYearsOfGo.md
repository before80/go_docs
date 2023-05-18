+++
title = "go 六岁了"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Six years of Go - go 六岁了

https://go.dev/blog/6years

Andrew Gerrand
10 November 2015

Six years ago today the Go language was released as an open source project. Since then, more than 780 contributors have made over 30,000 commits to the project’s 22 repositories. The ecosystem continues to grow, with GitHub reporting more than 90,000 Go repositories. And, offline, we see new Go events and user groups pop up [around](https://blog.golang.org/gophercon2015) [the](https://blog.golang.org/gouk15) [world](https://blog.golang.org/gopherchina) with regularity.

6年前的今天，Go语言作为一个开源项目被发布。从那时起，超过780位贡献者向项目的22个存储库提交了超过3万条信息。该生态系统继续增长，GitHub报告了超过90,000个Go存储库。而且，在线下，我们看到新的Go活动和用户组经常在世界各地出现。

![img](SixYearsOfGo_img/6years-gopher-16687730581303.png)

In August we [released Go 1.5](https://blog.golang.org/go1.5), the most significant release since Go 1. It features a completely [redesigned garbage collector](https://go.dev/doc/go1.5#gc) that makes the language more suitable for latency-sensitive applications; it marks the transition from a C-based compiler tool chain to one [written entirely in Go](https://go.dev/doc/go1.5#c); and it includes ports to [new architectures](https://go.dev/doc/go1.5#ports), with better support for ARM processors (the chips that power most smartphones). These improvements make Go better suited to a broader range of tasks, a trend that we hope will continue over the coming years.

8月，我们发布了Go 1.5，这是自Go 1以来最重要的版本。它具有完全重新设计的垃圾收集器，使该语言更适合于对延迟敏感的应用；它标志着从基于C的编译器工具链过渡到完全用Go编写的工具链；它包括对新架构的移植，对ARM处理器（大多数智能手机的芯片）的更好支持。这些改进使Go更适合于更广泛的任务，我们希望这一趋势将在未来几年持续下去。

Improvements to tools continue to boost developer productivity. We introduced the [execution tracer](https://go.dev/cmd/trace/) and the “[go doc](https://go.dev/cmd/go/#hdr-Show_documentation_for_package_or_symbol)” command, as well as more enhancements to our various [static analysis tools](https://go.dev/talks/2014/static-analysis.slide). We are also working on an [official Go plugin for Sublime Text](https://groups.google.com/forum/#!topic/Golang-nuts/8oCSjAiKXUQ), with better support for other editors in the pipeline.

对工具的改进继续提高开发者的生产力。我们引入了执行跟踪器和 "go doc "命令，以及对我们各种静态分析工具的更多改进。我们还在为Sublime Text开发一个官方的Go插件，并且正在开发对其他编辑器的更好支持。

Early next year we will release more improvements in Go 1.6, including HTTP/2 support for [net/http](https://go.dev/pkg/net/http/) servers and clients, an official package vendoring mechanism, support for blocks in text and HTML templates, a memory sanitizer that checks both Go and C/C++ code, and the usual assortment of other improvements and fixes.

明年年初，我们将在Go 1.6中发布更多的改进，包括对net/http服务器和客户端的HTTP/2支持，官方软件包销售机制，对文本和HTML模板中的块的支持，检查Go和C/C++代码的内存净化器，以及其他各种常见的改进和修复。

This is the sixth time we have had the pleasure of writing a birthday blog post for Go, and we would not be doing so if not for the wonderful and passionate people in our community. The Go team would like to thank everyone who has contributed code, written an open source library, authored a blog post, helped a new gopher, or just given Go a try. Without you, Go would not be as complete, useful, or successful as it is today. Thank you, and celebrate!

这是我们第六次有幸为Go写生日博文，如果不是我们社区中优秀而热情的人们，我们是不会这样做的。Go团队要感谢每一个贡献过代码、编写过开源库、撰写过博文、帮助过新的打地鼠，或者只是尝试过Go的人。没有你们，Go就不会像今天这样完整、有用或成功。感谢你们，并庆祝吧
