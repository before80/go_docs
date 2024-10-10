+++
title = "Go 的第一版的预览"
weight = 8
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# A preview of Go version 1 - Go 的第一版的预览

> 原文：[https://go.dev/blog/go1-preview](https://go.dev/blog/go1-preview)

Russ Cox
5 October 2011

We want to be able to provide a stable base for people using Go. People should be able to write Go programs and expect that they will continue to compile and run without change, on a timescale of years. Similarly, people should be able to write books about Go, be able to say which version of Go the book is describing, and have that version number still be meaningful much later. None of these properties is true for Go today.

我们希望能够为使用Go的人提供一个稳定的基础。人们应该能够编写Go程序，并期望这些程序能够继续编译和运行而不发生变化，而且时间尺度为数年。同样的，人们应该能够写出关于Go的书，能够说出书中所描述的Go的版本，并且这个版本号在很久以后仍然是有意义的。这些特性在今天的Go中都不存在。

We propose to issue a Go release early next year that will be called "Go version 1", Go 1 for short, that will be the first Go release to be stable in this way. Code that compiles in Go version 1 should, with few exceptions, continue to compile throughout the lifetime of that version, as we issue updates and bug fixes such as Go version 1.1, 1.2, and so on. It will also be maintained with fixes for bugs and security flaws even as other versions may evolve. Also, production environments such as Google App Engine will support it for an extended time.

我们建议在明年初发布一个名为 "Go 1版 "的Go版本，简称Go 1，这将是第一个以这种方式稳定的Go版本。在Go版本1中编译的代码，除了少数例外，应该会在该版本的整个生命周期中继续编译，因为我们会发布更新和错误修复，如Go版本1.1、1.2等等。即使在其他版本可能发展的情况下，也会对其进行维护，修复错误和安全缺陷。另外，生产环境如Google App Engine将在较长的时间内支持它。

Go version 1 will be a stable language with stable libraries. Other than critical fixes, changes made to the library and packages for versions 1.1, 1.2 and so on may add functionality but will not break existing Go version 1 programs.

Go版本1将是一种具有稳定库的稳定语言。除了关键的修复之外，对1.1、1.2等版本的库和包所做的修改可能会增加功能，但不会破坏现有的Go 1版本程序。

Our goal is for Go 1 to be a stable version of today’s Go, not a wholesale rethinking of the language. In particular, we are explicitly resisting any efforts to design new language features "by committee."

我们的目标是让Go 1成为当今Go的稳定版本，而不是对语言进行全面的重新思考。特别是，我们明确抵制任何 "通过委员会 "来设计新语言功能的努力。

However, there are various changes to the Go language and packages that we have intended for some time and prototyped but have not deployed yet, primarily because they are significant and backwards-incompatible. If Go 1 is to be long-lasting, it is important that we plan, announce, implement, and test these changes as part of the preparation of Go 1, rather than delay them until after it is released and thereby introduce divergence that contradicts our goals.

然而，有一些对Go语言和包的改变，我们已经打算了一段时间，并制作了原型，但还没有部署，主要是因为这些改变很重要，而且向后不兼容。如果Go 1要长期存在，我们就必须在准备Go 1的过程中计划、宣布、实施和测试这些变化，而不是拖到Go 1发布之后，从而引入与我们的目标相矛盾的分歧。

Today, we are publishing our preliminary [plan for Go 1](https://docs.google.com/document/pub?id=1ny8uI-_BHrDCZv_zNBSthNKAMX_fR_0dc6epA6lztRE) for feedback from the Go community. If you have feedback, please reply to the [thread on the golang-nuts mailing list](http://groups.google.com/group/golang-nuts/browse_thread/thread/badc4f323431a4f6).

今天，我们公布Go 1的初步计划，以征求Go社区的反馈。如果您有反馈意见，请在golang-nuts邮件列表中回复该主题。
