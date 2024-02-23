+++
title = "高级 Go 并发模式"
weight = 12
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Advanced Go Concurrency Patterns - 高级 Go 并发模式

> 原文：[https://go.dev/blog/io2013-talk-concurrency](https://go.dev/blog/io2013-talk-concurrency)

Andrew Gerrand
23 May 2013

At Google I/O a year ago Rob Pike presented [*Go Concurrency Patterns*](https://go.dev/talks/2012/concurrency.slide), an introduction to Go’s concurrency model. Last week, at I/O 2013, Go team member Sameer Ajmani continued the story with [*Advanced Go Concurrency Patterns*](http://go.dev/talks/2013/advconc.slide), an in-depth look at a real concurrent programming problem. The talk shows how to detect and avoid deadlocks and race conditions, and demonstrates the implementation of deadlines, cancellation, and more. For those who want to take their Go programming to the next level, this is a must-see.

在一年前的谷歌I/O大会上，Rob Pike介绍了Go并发模式，这是一个关于Go并发模型的介绍。上周，在I/O 2013大会上，Go团队成员Sameer Ajmani继续讲述了高级Go并发模式，深入探讨了一个真实的并发编程问题。该讲座展示了如何检测和避免死锁和竞赛条件，并演示了最后期限、取消等的实现。对于那些想把Go编程提高到新水平的人来说，这是个必看的节目。

<iframe src="https://www.youtube.com/embed/QDDwwePbDtw?rel=0" width="560" height="315" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

The slides are [available here](https://go.dev/talks/2013/advconc.slide) (use the left and right arrows to navigate).

幻灯片可以在这里找到（使用左右箭头进行导航）。

The slides were produced with [the present tool](https://godoc.org/golang.org/x/tools/present), and the runnable code snippets are powered by the [Go Playground](http://play.golang.org/). The source code for this talk is in [the go.talks sub-repository](https://github.com/golang/talks/tree/master/content/2013/advconc).

幻灯片是用本工具制作的，可运行的代码片段是由Go Playground提供的。本讲座的源代码在go.talks子资源库中。
