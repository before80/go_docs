+++
title = "泛型的下一步"
weight = 8
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# The Next Step for Generics - 泛型的下一步

https://go.dev/blog/generics-next-step

Ian Lance Taylor and Robert Griesemer
16 June 2020

## Introduction 简介

It’s been almost a year since we [last wrote about the possibility of adding generics to Go](https://blog.golang.org/why-generics). It’s time for an update.

距离我们上次写关于在Go中加入泛型的可能性已经过去了将近一年。现在是更新的时候了。

## Updated design 更新的设计

We’ve been continuing to refine the [generics design draft](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-contracts.md). We’ve written a type checker for it: a program that can parse Go code that uses generics as described in the design draft and report any type errors. We’ve written example code. And we’ve collected feedback from many, many people—thanks for providing it!

我们一直在继续完善泛型的设计草案。我们已经为它写了一个类型检查器：一个可以解析使用设计草案中描述的泛型的 Go 代码并报告任何类型错误的程序。我们已经编写了示例代码。我们还收集了许多人的反馈--感谢他们提供的反馈!

Based on what we’ve learned, we’re releasing an [updated design draft](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md). The biggest change is that we are dropping the idea of contracts. The difference between contracts and interface types was confusing, so we’re eliminating that difference. Type parameters are now constrained by interface types. Interface types are now permitted to include type lists, though only when used as constraints; in the previous design draft type lists were a feature of contracts. More complex cases will use a parameterized interface type.

根据我们所了解的情况，我们发布了一个更新的设计草案。最大的变化是，我们放弃了合约的概念。契约和接口类型之间的区别令人困惑，所以我们要消除这种区别。类型参数现在受制于接口类型。接口类型现在被允许包括类型列表，尽管只在作为约束条件时使用；在以前的设计草案中，类型列表是契约的一个特征。更复杂的情况将使用参数化的接口类型。

We hope that people will find this design draft simpler and easier to understand.

我们希望人们会发现这个设计草案更简单，更容易理解。

## Experimentation tool 实验工具

To help decide how to further refine the design draft, we are releasing a translation tool. This is a tool that permits people to type check and run code written using the version of generics described in the design draft. It works by translating generic code into ordinary Go code. This translation process imposes some limitations, but we hope that it will be good enough for people to get a feel for what generic Go code might look like. The real implementation of generics, if they are accepted into the language, will work differently. (We have only just begun to sketch out what a direct compiler implementation would look like.)

为了帮助决定如何进一步完善设计草案，我们发布了一个翻译工具。这是一个允许人们对使用设计草案中描述的泛型版本编写的代码进行类型检查和运行的工具。它的工作原理是将泛型代码翻译成普通Go代码。这个翻译过程有一些限制，但我们希望它足以让人们感受到泛型Go代码可能是什么样子。如果泛型被接受到语言中，泛型的真正实现将以不同的方式工作。(我们才刚刚开始勾勒出直接编译器实现的模样）。

The tool is available on a variant of the Go playground at [https://go2goplay.golang.org](https://go2goplay.golang.org/). This playground works just like the usual Go playground, but it supports generic code.

这个工具可以在 Go playground 的一个变体上使用，网址是 https://go2goplay.golang.org。这个操场的工作方式与通常的Go操场一样，但它支持通用代码。

You can also build and use the tool yourself. It is available in a branch of the master Go repo. Follow the [instructions on installing Go from source](https://go.dev/doc/install/source). Where those instructions direct you to check out the latest release tag, instead run `git checkout dev.go2go`. Then build the Go toolchain as directed.

您也可以自己构建和使用这个工具。它在 Go 主版本的一个分支中可用。请遵循从源代码安装 Go 的说明。如果这些说明指示您查看最新的发布标签，请运行 git checkout dev.go2go。然后按照指示构建 Go 工具链。

The translation tool is documented in [README.go2go](https://go.googlesource.com/go/+/refs/heads/dev.go2go/README.go2go.md).

翻译工具在 README.go2go 中有记录。

## Next steps 接下来的步骤

We hope that the tool will give the Go community a chance to experiment with generics. There are two main things that we hope to learn.

我们希望这个工具能给Go社区一个实验泛型的机会。我们希望能学到两个主要的东西。

First, does generic code make sense? Does it feel like Go? What surprises do people encounter? Are the error messages useful?

首先，泛型代码有意义吗？它感觉像Go吗？人们遇到了什么惊喜？错误信息有用吗？

Second, we know that many people have said that Go needs generics, but we don’t necessarily know exactly what that means. Does this draft design address the problem in a useful way? If there is a problem that makes you think "I could solve this if Go had generics," can you solve the problem when using this tool?

第二，我们知道很多人都说Go需要泛型，但我们不一定知道这到底是什么意思。这个设计草案是否以一种有用的方式解决了这个问题？如果有一个问题让您觉得 "如果Go有泛型，我就能解决这个问题"，那么在使用这个工具时，您能解决这个问题吗？

We will use the feedback we gather from the Go community to decide how to move forward. If the draft design is well received and doesn’t need significant changes, the next step would be a [formal language change proposal](https://go.dev/s/proposal). To set expectations, if everybody is completely happy with the design draft and it does not require any further adjustments, the earliest that generics could be added to Go would be the Go 1.17 release, scheduled for August 2021. In reality, of course, there may be unforeseen problems, so this is an optimistic timeline; we can’t make any definite prediction.

我们将利用从Go社区收集到的反馈来决定如何继续前进。如果设计草案受到好评，并且不需要重大修改，下一步将是正式的语言修改提案。为了设定预期，如果大家对设计草案完全满意，并且不需要任何进一步的调整，那么最早可以在Go中加入泛型的是Go 1.17版本，计划在2021年8月发布。当然，在现实中，可能会有不可预见的问题，所以这是一个乐观的时间表；我们不能做出任何明确的预测。

## Feedback 反馈意见

The best way to provide feedback for the language changes will be on the mailing list `golang-nuts@googlegroups.com`. Mailing lists are imperfect, but they seem like our best option for initial discussion. When writing about the design draft, please put `[generics]` at the start of the Subject line and to start different threads for different specific topics.

对语言变化提供反馈的最好方式是在邮件列表golang-nuts@googlegroups.com。邮件列表并不完美，但它们似乎是我们进行初步讨论的最佳选择。当写到设计草案时，请在主题行的开头写上[泛型]，并为不同的特定主题开始不同的主题。

If you find bugs in the generics type checker or the translation tool, they should be filed in the standard Go issue tracker at [https://golang.org/issue](https://go.dev/issue). Please start the issue title with `cmd/go2go:`. Note that the issue tracker is not the best place to discuss changes to the language, because it does not provide threading and it is not well suited to lengthy conversations.

如果您发现泛型类型检查器或翻译工具中的错误，应将其归档到标准的Go问题跟踪器中，网址是https://golang.org/issue。请以cmd/go2go:作为问题标题的开头。请注意，问题跟踪器并不是讨论语言变化的最佳场所，因为它不提供线程，也不太适合冗长的对话。

We look forward to your feedback.

我们期待着您的反馈。

## Acknowledgements 鸣谢

We’re not finished, but we’ve come a long way. We would not be here without a lot of help.

我们还没有完成，但我们已经走了很长的路。没有很多人的帮助，我们就不会有今天。

We’d like to thank Philip Wadler and his collaborators for thinking formally about generics in Go and helping us clarify the theoretical aspects of the design. Their paper [Featherweight Go](https://arxiv.org/abs/2005.11710) analyzes generics in a restricted version of Go, and they have developed a prototype [on GitHub](https://github.com/rhu1/fgg).

我们要感谢Philip Wadler和他的合作者，他们对Go中的泛型进行了正式的思考，帮助我们澄清了设计的理论方面。他们的论文Featherweight Go分析了Go的限制性版本中的泛型，并且他们在GitHub上开发了一个原型。

We would also like to thank [the people](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md#acknowledgements) who provided detailed feedback on an earlier version of the design draft.

我们还要感谢那些对早期版本的设计草案提供详细反馈的人。

And last but definitely not least, we’d like to thank many people on the Go team, many contributors to the Go issue tracker, and everybody else who shared ideas and feedback on earlier design drafts. We read all of it, and we’re grateful. We wouldn’t be here without you.

最后但绝对不是最不重要的，我们要感谢Go团队的许多人，Go问题跟踪器的许多贡献者，以及其他所有对早期设计草案分享想法和反馈的人。我们阅读了所有的反馈，并且非常感激。没有您们，我们就不会有今天。
