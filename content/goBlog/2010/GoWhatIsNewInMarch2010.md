+++
title = "go： 2010年3月有什么新动态"
weight = 16
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go: What's New in March 2010- go： 2010年3月有什么新动态

> 原文：[https://go.dev/blog/hello-world](https://go.dev/blog/hello-world)

Andrew Gerrand
18 March 2010

2010年3月18日

​	欢迎来到官方Go博客。我们，Go团队，希望通过这个博客向世界展示Go编程语言和其周围日益增长的库和应用的发展情况。

​	自我们上次发布（去年11月）以来已经过了几个月，所以让我们谈谈在Go世界中发生的事情。

​	Google的核心团队继续开发语言、编译器、包、工具和文档。编译器现在产生的代码在某些情况下比发布时快2倍到一个数量级。我们已经制作了一些选定的[基准测试（Benchmarks）](http://godashboard.appspot.com/benchmarks)的图表，并且[构建状态 （Build Status）](http://godashboard.appspot.com/)页面跟踪了提交到存储库的每个更改集的可靠性。

​	我们进行了语法更改，使语言更加简洁、规则化和灵活。语言中的分号[几乎完全被删除](http://groups.google.com/group/golang-nuts/t/5ee32b588d10f2e9)。[…T语法]({{< ref "/langSpec/Types#function-types">}})使得处理任意数量的带类型的函数参数更加简单。`x[lo:]`的语法现在是`x[lo:len(x)]`的缩写。Go现在本地支持复数。请参阅[发布说明](https://go.dev/doc/devel/release.html)以了解更多信息。

​	[Godoc](https://go.dev/cmd/godoc/)现在为第三方库提供更好的支持，还发布了一个新工具[goinstall](https://go.dev/cmd/goinstall)，使其易于安装。此外，我们已经开始研究一个包跟踪系统，以便更容易地找到所需内容。您可以在[Packages页面](http://godashboard.appspot.com/package)上查看其起始页。

​	[标准库](https://go.dev/pkg/)已经添加了超过40,000行代码，包括许多全新的包，其中相当大一部分由外部贡献者编写。

​	说到第三方，自从发布以来，一个充满活力的社区已经在我们的[邮件列表](http://groups.google.com/group/golang-nuts/)和irc频道（freenode上的#go-nuts）中蓬勃发展。我们已经有50多人正式加入了这个项目。他们的贡献从错误修复和文档修正到核心软件包和对其他操作系统的支持（Go现在支持FreeBSD，[Windows移植](http://code.google.com/p/go/wiki/WindowsPort)也正在进行中）。我们认为这些社区贡献是我们迄今为止最大的成功。

​		我们也收到了一些好评。这篇[最近在PC World上发表的文章](http://www.pcworld.idg.com.au/article/337773/google_go_captures_developers_imaginations/)总结了该项目周围的热情。一些博客作者已经开始记录他们在该语言中的经验（例如[这里](http://golang.tumblr.com/)、[这里](http://www.infi.nl/blog/view/id/47)和[这里](http://freecella.blogspot.com/2010/01/gospecify-basic-setup-of-projects.html)）。我们的用户的一般反应非常积极；一位初学者评论说，"[我对此印象非常深刻。Go在简单性和强大性之间走了一条优美的路线。](https://groups.google.com/group/golang-nuts/browse_thread/thread/5fabdd59f8562ed2)"

​	关于未来的计划：我们已经倾听了众多声音，听取了他们的需求，现在我们的重点是让 Go 准备好在主流应用中使用。我们正在改进垃圾回收器、运行时调度程序、工具和标准库，同时探索新的语言特性。2010 年将是 Go 的一个充满期待的年份，我们期待着与社区合作，使其成为一个成功的年份。

