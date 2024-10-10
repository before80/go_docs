+++
title = "关于在 Go 中添加泛型的建议"
weight = 100
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# A Proposal for Adding Generics to go - 关于在 go 中添加泛型的建议

> 原文：[https://go.dev/blog/generics-proposal](https://go.dev/blog/generics-proposal)

Ian Lance Taylor
12 January 2021

## Generics proposal 泛型建议

We’ve filed [a Go language change proposal](https://go.dev/issue/43651) to add support for type parameters for types and functions, permitting a form of generic programming.

我们已经提交了一份Go语言变更提案，以增加对类型和函数的类型参数的支持，允许某种形式的泛型编程。

## Why generics? 为什么是泛型？

Generics can give us powerful building blocks that let us share code and build programs more easily. Generic programming means writing functions and data structures where some types are left to be specified later. For example, you can write a function that operates on a slice of some arbitrary data type, where the actual data type is only specified when the function is called. Or, you can define a data structure that stores values of any type, where the actual type to be stored is specified when you create an instance of the data structure.

泛型可以为我们提供强大的构建模块，让我们可以更容易地分享代码和构建程序。泛型编程意味着编写函数和数据结构，其中一些类型留待以后指定。例如，您可以写一个函数，对某个任意数据类型的片断进行操作，而实际的数据类型只有在调用该函数时才被指定。或者，您可以定义一个存储任何类型的值的数据结构，在创建数据结构的实例时指定要存储的实际类型。

Since Go was first released in 2009, support for generics has been one of the most commonly requested language features. You can read more about why generics are useful in [an earlier blog post](https://blog.golang.org/why-generics).

自从Go在2009年首次发布以来，对泛型的支持一直是人们最常要求的语言功能之一。您可以在之前的一篇博文中读到更多关于泛型有用的原因。

Although generics have clear use cases, fitting them cleanly into a language like Go is a difficult task. One of the [first (flawed) attempts to add generics to Go](https://go.dev/design/15292/2010-06-type-functions) dates back all the way to 2010. There have been several others over the last decade.

虽然泛型有明确的用例，但将其干净利落地融入Go这样的语言是一项艰巨的任务。最早在Go中加入泛型的尝试之一（有缺陷）可以追溯到2010年。在过去的十年里，也有过几次这样的尝试。

For the last couple of years we’ve been working on a series of design drafts that have culminated in [a design based on type parameters](https://go.dev/design/go2draft-type-parameters). This design draft has had a lot of input from the Go programming community, and many people have experimented with it using the [generics playground](https://go2goplay.golang.org/) described in [an earlier blog post](https://blog.golang.org/generics-next-step). Ian Lance Taylor gave [a talk at GopherCon 2019](https://www.youtube.com/watch?v=WzgLqE-3IhY) about why to add generics and the strategy we are now following. Robert Griesemer gave [a follow-up talk about changes in the design, and the implementation, at GopherCon 2020](https://www.youtube.com/watch?v=TborQFPY2IM). The language changes are fully backward compatible, so existing Go programs will continue to work exactly as they do today. We have reached the point where we think that the design draft is good enough, and simple enough, to propose adding it to Go.

在过去的几年里，我们一直在进行一系列的设计草案，最终形成了一个基于类型参数的设计。这个设计草案得到了Go编程社区的大量意见，许多人使用早先博文中描述的泛型操场进行了实验。Ian Lance Taylor在GopherCon 2019上发表了关于为什么要添加泛型以及我们现在遵循的策略的演讲。Robert Griesemer在GopherCon 2020上发表了关于设计和实现的变化的后续演讲。语言的变化是完全向后兼容的，所以现有的Go程序将继续完全按照今天的方式工作。我们已经达到了这样的程度：我们认为设计草案足够好，也足够简单，可以提议将其加入Go中。

## What happens now? 现在会发生什么？

The [language change proposal process](https://go.dev/s/proposal) is how we make changes to the Go language. We have now [started this process](https://go.dev/issue/43651) to add generics to a future version of Go. We invite substantive criticisms and comments, but please try to avoid repeating earlier comments, and please try to [avoid simple plus-one and minus-one comments](https://go.dev/wiki/NoPlusOne). Instead, add thumbs-up/thumbs-down emoji reactions to comments with which you agree or disagree, or to the proposal as a whole.

语言修改建议过程是我们对Go语言进行修改的方式。我们现在已经开始了这个过程，将泛型添加到Go的未来版本中。我们邀请大家提出实质性的批评和意见，但请尽量避免重复先前的意见，也请尽量避免简单的加一和减一的意见。相反，请对您同意或不同意的评论，或对整个提案添加大拇指向上/向下的表情符号反应。

As with all language change proposals, our goal is to drive toward a consensus to either add generics to the language or let the proposal drop. We understand that for a change of this magnitude it will be impossible to make everybody in the Go community happy, but we intend to get to a decision that everybody is willing to accept.

就像所有的语言修改提案一样，我们的目标是推动达成共识，要么在语言中增加通用语，要么让提案放弃。我们知道，对于如此大规模的改变，不可能让Go社区的每个人都满意，但我们打算做出一个大家都愿意接受的决定。

If the proposal is accepted, our goal will be to have a complete, though perhaps not fully optimized, implementation for people to try by the end of the year, perhaps as part of the Go 1.18 betas.

如果提案被接受，我们的目标将是在年底前为人们提供一个完整的，尽管可能不是完全优化的实现，也许是作为Go 1.18测试版的一部分。
