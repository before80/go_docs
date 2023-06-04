+++
title = "go 1.11发布了"
weight = 9
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.11 is released - go 1.11发布了

https://go.dev/blog/go1.11

Andrew Bonventre
24 August 2018

Who says releasing on Friday is a bad idea?

谁说在周五发布是个坏主意？

Today the Go team is happy to announce the release of Go 1.11. You can get it from the [download page](https://go.dev/dl/).

今天Go团队很高兴地宣布Go 1.11的发布。您可以从下载页面获得它。

There are many changes and improvements to the toolchain, runtime, and libraries, but two features stand out as being especially exciting: modules and WebAssembly support.

在工具链、运行时和库方面有许多变化和改进，但有两个功能特别令人激动：模块和WebAssembly支持。

This release adds preliminary support for a [new concept called "modules,"](https://go.dev/doc/go1.11#modules) an alternative to GOPATH with integrated support for versioning and package distribution. Module support is considered experimental, and there are still a few rough edges to smooth out, so please make liberal use of the [issue tracker](https://go.dev/issue/new).

这个版本增加了对一个新概念的初步支持，称为 "模块"，是对GOPATH的替代，集成了对版本和包分发的支持。模块支持被认为是实验性的，仍有一些粗糙的边缘需要磨平，所以请自由使用问题跟踪器。

Go 1.11 also adds an experimental port to [WebAssembly](https://go.dev/doc/go1.11#wasm) (`js/wasm`). This allows programmers to compile Go programs to a binary format compatible with four major web browsers. You can read more about WebAssembly (abbreviated "Wasm") at [webassembly.org](https://webassembly.org/) and see [this wiki page](https://go.dev/wiki/WebAssembly) on how to get started with using Wasm with Go. Special thanks to [Richard Musiol](https://github.com/neelance) for contributing the WebAssembly port!

Go 1.11还增加了对WebAssembly（js/wasm）的实验性移植。这使得程序员可以将Go程序编译成与四个主要网络浏览器兼容的二进制格式。您可以在webassembly.org上阅读更多关于WebAssembly（缩写为 "Wasm"）的信息，并查看这个wiki页面，了解如何开始使用Wasm与Go。特别感谢Richard Musiol对WebAssembly移植的贡献!

We also want to thank everyone who contributed to this release by writing code, filing bugs, providing feedback, and/or testing the betas and release candidates. Your contributions and diligence helped to ensure that Go 1.11 is as bug-free as possible. That said, if you do notice any problems, please [file an issue](https://go.dev/issues/new).

我们也要感谢每一个通过编写代码、提交错误、提供反馈和/或测试测试版和候选版而为这个版本做出贡献的人。您的贡献和勤奋有助于确保Go 1.11尽可能的没有错误。也就是说，如果您发现任何问题，请提出问题。

For more detail about the changes in Go 1.11, see the [release notes](https://go.dev/doc/go1.11).

关于Go 1.11的更多细节变化，请参见发布说明。

Have a wonderful weekend and enjoy the release!

祝您周末愉快，享受发布的乐趣!
