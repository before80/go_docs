+++
title = "关于Go 1.15的建议"
weight = 16
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Proposals for Go 1.15 - 关于go 1.15的建议

> 原文：[https://go.dev/blog/go1.15-proposals](https://go.dev/blog/go1.15-proposals)

Robert Griesemer, for the Go team
28 January 2020

## Status 状态

We are close to the Go 1.14 release, planned for February assuming all goes well, with an RC1 candidate almost ready. Per the process outlined in the [Go 2, here we come!](https://blog.golang.org/go2-here-we-come) blog post, it is again the time in our development and release cycle to consider if and what language or library changes we might want to include for our next release, Go 1.15, scheduled for August of this year.

我们已经接近Go 1.14的发布了，如果一切顺利的话，计划在2月份发布，RC1候选版也快准备好了。根据Go 2，我们来了！博文中概述的过程，现在又到了我们的开发和发布周期，考虑我们是否要在下一个版本，即计划在今年8月发布的Go 1.15中包含哪些语言或库的变化。

The primary goals for Go remain package and version management, better error handling support, and generics. Module support is in good shape and getting better with each day, and we are also making progress on the generics front (more on that later this year). Our attempt seven months ago at providing a better error handling mechanism, the [`try` proposal](https://go.dev/issue/32437), met good support but also strong opposition and we decided to abandon it. In its aftermath there were many follow-up proposals, but none of them seemed convincing enough, clearly superior to the `try` proposal, or less likely to cause similar controversy. Thus, we have not further pursued changes to error handling for now. Perhaps some future insight will help us to improve upon the status quo.

Go的主要目标仍然是包和版本管理，更好的错误处理支持，以及泛型。模块支持情况良好，并且每天都在改善，我们在泛型方面也取得了进展（今年晚些时候会有更多的报道）。七个月前，我们试图提供一个更好的错误处理机制，即try提案，得到了良好的支持，但也遭到了强烈的反对，我们决定放弃它。在它之后，有许多后续建议，但没有一个看起来足够有说服力，明显优于try建议，或者不太可能引起类似的争论。因此，我们暂时没有进一步追求对错误处理的改变。也许未来的一些洞察力会帮助我们改善现状。

## Proposals 提议

Given that modules and generics are actively being worked on, and with error handling changes out of the way for the time being, what other changes should we pursue, if any? There are some perennial favorites such as requests for enums and immutable types, but none of those ideas are sufficiently developed yet, nor are they urgent enough to warrant a lot of attention by the Go team, especially when also considering the cost of making a language change.

鉴于模块和泛型的工作正在积极进行中，而且错误处理的修改也暂时停止了，如果有的话，我们还应该进行哪些修改？有一些多年来的最爱，比如对枚举和不可变类型的要求，但这些想法都还没有得到充分的发展，也没有迫切到需要Go团队的大量关注，特别是考虑到进行语言改变的成本。

After reviewing all potentially viable proposals, and more importantly, because we don’t want to incrementally add new features without a long-term plan, we concluded that it is better to hold off with major changes this time. Instead we concentrate on a couple of new `vet` checks and a minor adjustment to the language. We have selected the following three proposals:

在审查了所有潜在的可行建议后，更重要的是，由于我们不想在没有长期计划的情况下逐步增加新的功能，我们的结论是，这次最好还是不要进行重大的改变。相反，我们集中在几个新的审查检查和对语言的小调整上。我们选择了以下三个提案。

[#32479](https://go.dev/issue/32479). Diagnose `string(int)` conversion in `go vet`.

#32479. 诊断go vet中的字符串（int）转换。

We were planning to get this done for the upcoming Go 1.14 release but we didn’t get around to it, so here it is again. The `string(int)` conversion was introduced early in Go for convenience, but it is confusing to newcomers (`string(10)` is `"\n"` not `"10"`) and not justified anymore now that the conversion is available in the `unicode/utf8` package. Since [removing this conversion](https://go.dev/issue/3939) is not a backwards-compatible change, we propose to start with a `vet` error instead.

我们计划在即将发布的Go 1.14版本中完成这项工作，但我们没有来得及做，所以在此再次提出。string(int)转换是Go中早期引入的，目的是为了方便，但对于新来的人来说，它是混乱的（string(10)是"\n "而不是 "10"），而且现在unicode/utf8包中已经有了这种转换，所以不再是合理的。由于删除这个转换不是一个向后兼容的变化，我们建议用一个审核错误来代替。

[#4483](https://go.dev/issue/4483). Diagnose impossible interface-interface type assertions in `go vet`.

#4483. 诊断go vet中不可能的interface-interface类型断言。

Currently, Go permits any type assertion `x.(T)` (and corresponding type switch case) where the type of `x` and `T` are interfaces. Yet, if both `x` and `T` have a method with the same name but different signatures it is impossible for any value assigned to `x` to also implement `T`; such type assertions will always fail at runtime (panic or evaluate to `false`). Since we know this at compile time, the compiler might as well report an error. Reporting a compiler error in this case is not a backwards-compatible change, thus we also propose to start with a `vet` error instead.

目前，Go允许任何类型断言x.(T)（以及相应的类型转换情况），其中x和T的类型是接口。然而，如果x和T都有一个名字相同但签名不同的方法，那么分配给x的任何值都不可能同时实现T；这样的类型断言在运行时总是会失败（恐慌或评估为错误）。既然我们在编译时就知道这一点，编译器不妨报告一个错误。在这种情况下，报告一个编译器错误并不是一个向后兼容的变化，因此我们也建议用一个vet错误来代替。

[#28591](https://go.dev/issue/28591). Constant-evaluate index and slice expressions with constant strings and indices.

#28591. 用恒定的字符串和索引对索引和切片表达式进行恒定评价。

Currently, indexing or slicing a constant string with a constant index, or indices, produces a non-constant `byte` or `string` value, respectively. But if all operands are constant, the compiler can constant-evaluate such expressions and produce a constant (possibly untyped) result. This is a fully backward-compatible change and we propose to make the necessary adjustments to the spec and compilers.

目前，用一个或多个常数索引对一个常数字符串进行索引或切分，分别产生一个非常数的字节或字符串值。但是如果所有的操作数都是常数，编译器可以对这种表达式进行常数评估，并产生一个常数（可能是未定型的）结果。这是一个完全向后兼容的变化，我们建议对规范和编译器做必要的调整。

(Correction: We found out after posting that this change is not backward-compatible; see [comment](https://go.dev/issue/28591#issuecomment-579993684) for details.)

(更正一下。我们在发帖后发现，这一变化并不向后兼容；详情见评论）。

## Timeline 时间轴

We believe that none of these three proposals are controversial but there’s always a chance that we missed something important. For that reason we plan to have the proposals implemented at the beginning of the Go 1.15 release cycle (at or shortly after the Go 1.14 release) so that there is plenty of time to gather experience and provide feedback. Per the [proposal evaluation process](https://blog.golang.org/go2-here-we-come), the final decision will be made at the end of the development cycle, at the beginning of May, 2020.

我们认为这三个建议都没有争议，但总有可能我们会错过一些重要的东西。出于这个原因，我们计划在Go 1.15发布周期之初（Go 1.14发布之时或之后不久）实施这些提案，这样就有足够的时间来收集经验和提供反馈。根据提案的评估过程，最终决定将在开发周期结束时，即2020年5月初做出。

## And one more thing… 还有一件事...

We receive many more language change proposals ([issues labeled LanguageChange](https://github.com/golang/go/labels/LanguageChange)) than we can review thoroughly. For instance, just for error handling alone, there are 57 issues, of which five are currently still open. Since the cost of making a language change, no matter how small, is high and the benefits are often unclear, we must err on the side of caution. Consequently, most language change proposals get rejected sooner or later, sometimes with minimal feedback. This is unsatisfactory for all parties involved. If you have spent a lot of time and effort outlining your idea in detail, it would be nice to not have it immediately rejected. On the flip side, because the general [proposal process](https://github.com/golang/proposal/blob/master/README.md) is deliberately simple, it is very easy to create language change proposals that are only marginally explored, causing the review committee significant amounts of work. To improve this experience for everybody we are adding a new [questionnaire](https://github.com/golang/proposal/blob/master/go2-language-changes.md) for language changes: filling out that template will help reviewers evaluate proposals more efficiently because they don’t need to try to answer those questions themselves. And hopefully it will also provide better guidance for proposers by setting expectations right from the start. This is an experiment that we will refine over time as needed.

我们收到的语言变更提案（标有LanguageChange的问题）比我们能彻底审查的还要多。例如，仅就错误处理而言，就有57个问题，其中5个目前仍未解决。由于进行语言改变的成本很高，无论多么小，而且好处往往不明确，我们必须谨慎行事。因此，大多数语言修改建议迟早会被否决，有时反馈很少。这对所有相关方来说都是不满意的。如果您花费了大量的时间和精力来详细概述您的想法，那么不被立即拒绝就好了。反过来说，由于一般的提案程序是刻意简单的，所以很容易产生一些只是稍作探讨的语言修改提案，给审查委员会带来大量的工作。为了改善每个人的这种体验，我们正在为语言变化增加一份新的调查问卷：填写该模板将帮助审查员更有效地评估提案，因为他们不需要自己尝试回答这些问题。而且，希望它也能通过从一开始就设定预期，为提案人提供更好的指导。这是一个实验，我们将根据需要逐步完善。

Thank you for helping us improve the Go experience!

谢谢您帮助我们改善Go体验
