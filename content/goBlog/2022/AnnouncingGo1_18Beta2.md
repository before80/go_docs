+++
title = "宣布Go 1.18 Beta 2"
weight = 99
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Announcing Go 1.18 Beta 2 - 宣布Go 1.18 Beta 2

> 原文：[https://go.dev/blog/go1.18beta2](https://go.dev/blog/go1.18beta2)

Jeremy Faller and Steve Francia, for the Go team
31 January 2022

We are encouraged by all the excitement around Go’s upcoming 1.18 release, which adds support for [generics](https://go.dev/blog/why-generics), [fuzzing](https://go.dev/blog/fuzz-beta), and the new [Go workspace mode](https://go.dev/design/45713-workspace).

Go的1.18版本即将发布，它增加了对泛型的支持、模糊处理和新的Go工作区模式，我们为大家的兴奋感到鼓舞。

We released Go 1.18 beta 1 two months ago, and it is now the most downloaded Go beta ever, with twice as many downloads as any previous release. Beta 1 has also proved very reliable; in fact, we are already running it in production here at Google.

我们在两个月前发布了Go 1.18 beta1，它现在是有史以来下载量最大的Go测试版，其下载量是以前任何版本的两倍。Beta 1也被证明是非常可靠的；事实上，我们已经在谷歌这里的生产中运行它。

Your feedback on Beta 1 helped us identify obscure bugs in the new support for generics and ensure a more stable final release. We’ve resolved these issues in today’s release of Go 1.18 Beta 2, and we encourage everyone to try it out. The easiest way to install it alongside your existing Go toolchain is to run:

您对Beta 1的反馈帮助我们发现了新的泛型支持中的一些不明显的错误，并确保最终版本更加稳定。在今天发布的Go 1.18 Beta 2中，我们已经解决了这些问题，我们鼓励大家尝试一下。在您现有的Go工具链旁边安装它的最简单方法是运行：

```go
go install golang.org/dl/go1.18beta2@latest
go1.18beta2 download
```

After that, you can run `go1.18beta2` as a drop-in replacement for `go`. For more download options, visit https://go.dev/dl/#go1.18beta2.

之后，您可以运行go1.18beta2作为go的直接替代品。有关更多的下载选项，请访问https://go.dev/dl/#go1.18beta2。

Because we are taking the time to issue a second beta, we now expect that the Go 1.18 release candidate will be issued in February, with the final Go 1.18 release in March.

由于我们正在花时间发布第二个测试版，我们现在预计Go 1.18候选版将在2月发布，最终的Go 1.18版将在3月发布。

The Go language server `gopls` and the VS Code Go extension now support generics. To install `gopls` with generics, see [this documentation](https://github.com/golang/tools/blob/master/gopls/doc/advanced.md#working-with-generic-code), and to configure the VS Code Go extension, follow [this instruction](https://github.com/golang/vscode-go/blob/master/docs/advanced.md#using-go118).

Go语言服务器gopls和VS Code Go扩展现在支持泛型。要安装带有泛型的gopls，请参见此文档，要配置VS Code Go扩展，请遵循此说明。

As always, especially for beta releases, if you notice any problems, please [file an issue](https://go.dev/issue/new).

像往常一样，特别是对于测试版，如果您发现任何问题，请提交一个问题。
