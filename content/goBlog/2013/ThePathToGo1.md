+++
title = "通往 go 1的道路"
weight = 14
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# The path to Go 1 - 通往 go 1的道路

> 原文：[https://go.dev/blog/go1-path](https://go.dev/blog/go1-path)

Andrew Gerrand
14 March 2013

In July 2012, Rob Pike and I presented a talk at OSCON titled *The path to Go 1*. In it we explain how Go 1 came to be, and outline the process by which Go was refined and stabilized to become the clean, consistent programming environment that it is today. We present the major highlights of the release and discuss the details behind some specific libraries and tools.

2012年7月，Rob Pike和我在OSCON上发表了题为《Go 1之路》的演讲。在演讲中，我们解释了Go 1是如何诞生的，并概述了Go是如何被完善和稳定下来，成为今天这个干净、一致的编程环境的过程。我们介绍了该版本的主要亮点，并讨论了一些特定库和工具背后的细节。

<iframe src="https://www.youtube.com/embed/bj9T2c2Xk_s" width="560" height="315" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

The slides for the talk are [available here](https://go.dev/talks/2012/go1.slide).

讲座的幻灯片可以在这里找到。

It’s almost a year since we cut Go 1.0 and we are now busy preparing Go 1.1. The release will include performance improvements to the gc compiler, garbage collector, and goroutine scheduler, some standard library additions, and many bug fixes and other improvements. Stay tuned, as we hope to release Go 1.1 in the coming weeks.

自从我们砍掉Go 1.0以来，已经快一年了，我们现在正忙着准备Go 1.1。该版本将包括对gc编译器、垃圾收集器和goroutine调度器的性能改进，一些标准库的添加，以及许多bug修复和其他改进。敬请关注，我们希望在未来几周内发布Go 1.1。
