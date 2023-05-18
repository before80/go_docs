+++
title = "Go 2, here we come!"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go 2, here we come!

https://go.dev/blog/go2-here-we-come

Robert Griesemer
29 November 2018

## Background 背景介绍

At GopherCon 2017, Russ Cox officially started the thought process on the next big version of Go with his talk [The Future of Go](https://www.youtube.com/watch?v=0Zbh_vmAKvk) ([blog post](https://blog.golang.org/toward-go2)). We have called this future language informally Go 2, even though we understand now that it will arrive in incremental steps rather than with a big bang and a single major release. Still, Go 2 is a useful moniker, if only to have a way to talk about that future language, so let’s keep using it for now.

在2017年的GopherCon上，Russ Cox以他的演讲《Go的未来》（博文）正式开始了对Go下一个大版本的思考。我们将这种未来的语言非正式地称为Go 2，尽管我们现在明白，它将以渐进的方式到来，而不是大爆炸和单一的主要版本。尽管如此，Go 2仍然是一个有用的名称，如果只是为了有一个谈论未来语言的方法的话，那么我们现在就继续使用它。

A major difference between Go 1 and Go 2 is who is going to influence the design and how decisions are made. Go 1 was a small team effort with modest outside influence; Go 2 will be much more community-driven. After almost 10 years of exposure, we have learned a lot about the language and libraries that we didn’t know in the beginning, and that was only possible through feedback from the Go community.

Go 1和Go 2的一个主要区别是谁将影响设计以及如何做出决定。Go 1是一个小团队的努力，外部影响不大；Go 2将更多地由社区驱动。经过近10年的接触，我们学到了很多关于语言和库的知识，这些都是我们一开始不知道的，而这只有通过Go社区的反馈才能实现。

In 2015 we introduced the [proposal process](https://go.dev/s/proposal) to gather a specific kind of feedback: proposals for language and library changes. A committee composed of senior Go team members has been reviewing, categorizing, and deciding on incoming proposals on a regular basis. That has worked pretty well, but as part of that process we have ignored all proposals that are not backward-compatible, simply labeling them Go 2 instead. In 2017 we also stopped making any kind of incremental backward-compatible language changes, however small, in favor of a more comprehensive plan that takes the bigger picture of Go 2 into account.

2015年，我们引入了提案程序来收集一种特殊的反馈：关于语言和库变化的提案。一个由Go团队高级成员组成的委员会一直在定期审查、分类和决定收到的提案。这样做效果很好，但作为该过程的一部分，我们忽略了所有不向后兼容的提案，而只是简单地将其标记为Go 2。在2017年，我们也停止了任何形式的渐进式向后兼容的语言变化，无论多么微小，都是为了支持一个更全面的计划，将Go 2的大局考虑在内。

It is now time to act on the Go 2 proposals, but to do this we first need a plan.

现在是对Go 2提案采取行动的时候了，但要做到这一点，我们首先需要一个计划。

## Status 状况

At the time of writing, there are around 120 [open issues labeled Go 2 proposal](https://github.com/golang/go/issues?page=1&q=is%3Aissue+is%3Aopen+label%3Aproposal+label%3AGo2&utf8=✓). Each of them proposes a significant library or language change, often one that does not satisfy the existing [Go 1 compatibility guarantee](https://go.dev/doc/go1compat). Ian Lance Taylor and I have been working through these proposals and categorized them ([Go2Cleanup](https://github.com/golang/go/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Aproposal+label%3AGo2+label%3AGo2Cleanup), [NeedsDecision](https://github.com/golang/go/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Aproposal+label%3AGo2+label%3ANeedsDecision), etc.) to get an idea of what’s there and to make it easier to proceed with them. We also merged related proposals and closed the ones which seemed clearly out of the scope of Go, or were otherwise unactionable.

在写这篇文章的时候，大约有120个标记为Go 2提案的开放问题。每一个问题都提出了一个重要的库或语言的变化，通常不满足现有的Go 1兼容性保证。Ian Lance Taylor和我一直在研究这些提案，并将其分类（Go2Cleanup、NeedsDecision等），以了解其中的内容，并使其更容易进行。我们还合并了相关的提案，并关闭了那些明显超出Go范围的提案，或者其他无法行动的提案。

Ideas from the remaining proposals will likely influence Go 2’s libraries and languages. Two major themes have emerged early on: support for better error handling, and generics. [Draft designs](https://blog.golang.org/go2draft) for these two areas have been published at this year’s GopherCon, and more exploration is needed.

剩余提案中的想法可能会影响Go 2的库和语言。两个主要的主题在早期就已经出现了：支持更好的错误处理和泛型。这两个领域的设计草案已经在今年的GopherCon上公布，还需要更多的探索。

But what about the rest? We are [constrained](https://blog.golang.org/toward-go2) by the fact that we now have millions of Go programmers and a large body of Go code, and we need to bring it all along, lest we risk a split ecosystem. That means we cannot make many changes, and the changes we are going to make need to be chosen carefully. To make progress, we are implementing a new proposal evaluation process for these significant potential changes.

但是其他方面呢？我们受制于这样一个事实：我们现在有数以百万计的Go程序员和大量的Go代码，我们需要把它们都带过来，以免我们有分裂生态系统的风险。这意味着我们不能做很多改变，而且我们要做的改变需要谨慎选择。为了取得进展，我们正在为这些重大的潜在变化实施一个新的提案评估过程。

## Proposal evaluation process 提案评估过程

The purpose of the proposal evaluation process is to collect feedback on a small number of select proposals such that a final decision can be made. The process runs more or less in parallel to a release cycle and consists of the following steps:

提案评估过程的目的是收集对少数精选提案的反馈意见，以便做出最终决定。这个过程或多或少与发布周期平行，包括以下步骤：

1. *Proposal selection*. The Go team selects a small number of [Go 2 proposals](https://github.com/golang/go/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3AGo2+label%3AProposal) that seem worth considering for acceptance, without making a final decision. See below for more on the selection criteria. 提案选择。Go团队选择少量似乎值得考虑接受的Go 2提案，但不做最终决定。关于选择标准的更多信息，见下文。
2. *Proposal feedback*. The Go team sends out an announcement listing the selected proposals. The announcement explains to the community the tentative intent to move forward with the selected proposals and to collect feedback for each of them. This gives the community a chance to make suggestions and express concerns. 提案反馈。Go团队会发出公告，列出被选中的提案。该公告向社区解释了推进所选提案的暂定意图，并收集每个提案的反馈。这让社区有机会提出建议和表达关切。
3. *Implementation*. Based on that feedback, the proposals are implemented. The target for these significant language and library changes is to have them ready to submit on day 1 of an upcoming release cycle. 实施。根据这些反馈意见，这些建议被实施。对于这些重要的语言和库的变化，目标是在即将到来的发布周期的第一天就可以提交。
4. *Implementation feedback*. During the development cycle, the Go team and community have a chance to experiment with the new features and collect further feedback. 实施反馈。在开发周期内，Go团队和社区有机会对新功能进行试验，并收集进一步的反馈。
5. *Launch decision*. At the end of the three month [development cycle](https://github.com/golang/go/wiki/Go-Release-Cycle) (just when starting the three month repo freeze before a release), and based on the experience and feedback gathered during the release cycle, the Go team makes the final decision about whether to ship each change. This provides an opportunity to consider whether the change has delivered the expected benefits or created any unexpected costs. Once shipped, the changes become part of the language and libraries. Excluded proposals may go back to the drawing board or may be declined for good.发布决定。在三个月的开发周期结束时（就在开始发布前的三个月 repo 冻结时），根据在发布周期中收集到的经验和反馈，Go 团队会对是否发布每个变更做出最终决定。这提供了一个机会来考虑该变化是否带来了预期的收益或创造了任何意外的成本。一旦出货，这些变化就成为语言和库的一部分。被排除的建议可能会回到绘图板上，也可能被永远地拒绝。

With two rounds of feedback, this process is slanted towards declining proposals, which will hopefully prevent feature creep and help with keeping the language small and clean.

通过两轮反馈，这个过程倾向于拒绝提案，这将有望防止功能蠕变，并有助于保持语言的小而干净。

We can’t go through this process for each of the open Go 2 proposals, there are simply too many of them. That’s where the selection criteria come into play.

我们不可能对每一个开放的Go 2提案都进行这个过程，因为它们实在是太多了。这就是选择标准发挥作用的地方。

## Proposal selection criteria 提案选择标准

A proposal must at the very least:

一项提案至少要做到以下几点：

1. *address an important issue for many people*, 解决许多人的重要问题。
2. *have minimal impact on everybody else*, and 对其他人的影响最小，并且
3. *come with a clear and well-understood solution*. 有一个明确的、被充分理解的解决方案。

Requirement 1 ensures that any changes we make help as many Go developers as possible (make their code more robust, easier to write, more likely to be correct, and so on), while requirement 2 ensures we are careful to hurt as few developers as possible, whether by breaking their programs or causing other churn. As a rule of thumb, we should aim to help at least ten times as many developers as we hurt with a given change. Changes that don’t affect real Go usage are a net zero benefit put up against a significant implementation cost and should be avoided.

要求1确保我们所做的任何改变都能帮助尽可能多的Go开发者（使他们的代码更健壮、更容易编写、更可能正确，等等），而要求2确保我们小心翼翼地尽可能少地伤害开发者，无论是破坏他们的程序还是造成其他流失。作为一个经验法则，我们的目标应该是帮助至少十倍于我们在特定的变化中伤害的开发者。那些不影响真正的Go使用的变化，在巨大的实施成本面前，是一种净零收益，应该被避免。

Without requirement 3 we don’t have an implementation of the proposal. For instance, we believe that some form of genericity might solve an important issue for a lot of people, but we don’t yet have a clear and well-understood solution. That’s fine, it just means that the proposal needs to go back to the drawing board before it can be considered.

如果没有要求3，我们就没有建议的实施。例如，我们相信某种形式的通用性可能会解决很多人的一个重要问题，但我们还没有一个明确的、被充分理解的解决方案。这很好，这只是意味着在考虑这个提案之前，需要回到绘图板。

## Proposals 提案

We feel that this is a good plan that should serve us well but it is important to understand that this is only a starting point. As the process is used we will discover the ways in which it fails to work well and we will refine it as needed. The critical part is that until we use it in practice we won’t know how to improve it.

我们觉得这是一个很好的计划，应该能很好地为我们服务，但必须明白这只是一个起点。随着这个过程的使用，我们将发现它不能很好地发挥作用的方式，我们将根据需要对其进行完善。最关键的是，直到我们在实践中使用它，我们才会知道如何改进它。

A safe place to start is with a small number of backward-compatible language proposals. We haven’t done language changes for a long time, so this gets us back into that mode. Also, the changes won’t require us worrying about breaking existing code, and thus they serve as a perfect trial balloon.

一个安全的起点是少量的向后兼容的语言建议。我们已经很久没有做过语言上的改变了，所以这让我们重新回到了这种模式。而且，这些变化不需要我们担心破坏现有的代码，因此它们可以作为一个完美的试验气球。

With all that said, we propose the following selection of Go 2 proposals for the Go 1.13 release (step 1 in the proposal evaluation process):

综上所述，我们建议为Go 1.13版本选择以下Go 2提案（提案评估过程的第一步）：

1. [*#20706*](https://github.com/golang/go/issues/20706) *General Unicode identifiers based on* [*Unicode TR31*](http://unicode.org/reports/tr31/): This addresses an important issue for Go programmers using non-Western alphabets and should have little if any impact on anyone else. There are normalization questions which we need to answer and where community feedback will be important, but after that the implementation path is well understood. Note that identifier export rules will not be affected by this.#20706 基于Unicode TR31的通用Unicode标识符：这解决了使用非西方字母的Go程序员的一个重要问题，对其他人的影响很小。我们需要回答一些规范化的问题，社区的反馈也很重要，但之后的实施路径就很好理解了。请注意，标识符的导出规则不会受此影响。
2. [*#19308*](https://github.com/golang/go/issues/19308), [*#28493*](https://github.com/golang/go/issues/28493) *Binary integer literals and support for _ in number literals*: These are relatively minor changes that seem hugely popular among many programmers. They may not quite reach the threshold of solving an “important issue” (hexadecimal numbers have worked well so far) but they bring Go up to par with most other languages in this respect and relieve a pain point for some programmers. They have minimal impact on others who don’t care about binary integer literals or number formatting, and the implementation is well understood.#19308, #28493 二进制整数字头和数字字头中对_的支持。这些都是相对较小的变化，但在许多程序员中似乎非常受欢迎。它们可能还没有达到解决 "重要问题 "的门槛（到目前为止，十六进制的数字工作得很好），但它们使Go在这方面达到了大多数其他语言的水平，缓解了一些程序员的痛苦。它们对其他不关心二进制整数字或数字格式的人影响很小，而且实现起来也很容易理解。
3. [*#19113*](https://github.com/golang/go/issues/19113) *Permit signed integers as shift counts*: An estimated 38% of all non-constant shifts require an (artificial) uint conversion (see the issue for a more detailed break-down). This proposal will clean up a lot of code, get shift expressions better in sync with index expressions and the built-in functions cap and len. It will mostly have a positive impact on code. The implementation is well understood.#19113 允许有符号的整数作为移位计数。据估计，38%的非常数移位需要（人为的）uint转换（见该问题的更详细的分解）。这个建议将清理大量的代码，使移位表达式与索引表达式以及内置函数cap和len更好地同步。它将对代码产生很大的积极影响。这个实现是很好理解的。

## Next steps 接下来的步骤

With this blog post we have executed the first step and started the second step of the proposal evaluation process. It’s now up to you, the Go community, to provide feedback on the issues listed above.

通过这篇博文，我们已经执行了第一步，开始了提案评估过程的第二步。现在就看你了，Go社区，对上面列出的问题提供反馈。

For each proposal for which we have clear and approving feedback, we will move forward with the implementation (step 3 in the process). Because we want the changes implemented on the first day of the next release cycle (tentatively Feb. 1, 2019) we may start the implementation a bit early this time to leave time for two full months of feedback (Dec. 2018, Jan. 2019).

对于每一个我们有明确和认可的反馈的建议，我们将推进实施（过程中的第三步）。因为我们希望在下一个发布周期（暂定2019年2月1日）的第一天实施这些变化，所以这次我们可能会提前一点开始实施，以便留出时间来听取两个完整的反馈（2018年12月，2019年1月）。

For the 3-month development cycle (Feb. to May 2019) the chosen features are implemented and available at tip and everybody will have a chance to gather experience with them. This provides another opportunity for feedback (step 4 in the process).

在3个月的开发周期（2019年2月至5月），所选择的功能被实施，并在小费中可用，每个人都将有机会收集使用经验。这提供了另一个反馈的机会（过程中的第四步）。

Finally, shortly after the repo freeze (May 1, 2019), the Go team makes the final decision whether to keep the new features for good (and include them in the Go 1 compatibility guarantee), or whether to abandon them (final step in the process).

最后，在repo冻结后不久（2019年5月1日），Go团队做出最终决定，是否永远保留新功能（并将其纳入Go 1的兼容性保证），或是否放弃它们（流程中的最后一步）。

(Since there is a real chance that a feature may need to be removed just when we freeze the repo, the implementation will need to be such that the feature can be disabled without destabilizing the rest of the system. For language changes that may mean that all feature-related code is guarded by an internal flag.)

(由于确实有可能在我们冻结 repo 时需要删除某项功能，所以实现时需要做到在不破坏系统其他部分稳定的情况下禁用该功能）。对于语言变化来说，这可能意味着所有与特性相关的代码都要有一个内部标志来保护。）

This will be the first time that we have followed this process, hence the repo freeze will also be a good moment to reflect on the process and to adjust it if necessary. Let’s see how it goes.

这将是我们第一次遵循这个过程，因此，版本冻结也将是一个很好的时机来反思这个过程，并在必要时进行调整。让我们看看进展如何。

Happy evaluating!

评估愉快！
