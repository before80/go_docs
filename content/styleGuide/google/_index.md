+++
title = "Go 风格"
date = 2024-01-22T09:32:42+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go Style - Go 风格

> 原文：[https://google.github.io/styleguide/go](https://google.github.io/styleguide/go)

## About 关于

The Go Style Guide and accompanying documents codify the current best approaches for writing readable and idiomatic Go. Adherence to the Style Guide is not intended to be absolute, and these documents will never be exhaustive. Our intention is to minimize the guesswork of writing readable Go so that newcomers to the language can avoid common mistakes. The Style Guide also serves to unify the style guidance given by anyone reviewing Go code at Google.

​	Go 风格指南和随附文档编纂了编写可读且符合惯例的 Go 代码的当前最佳方法。遵守风格指南并非绝对的，并且这些文档永远不会详尽无遗。我们的目的是最大限度地减少编写可读 Go 代码的猜测，以便语言新手可以避免常见错误。风格指南还用于统一 Google 中任何审阅 Go 代码的人员提供的风格指导。

| Document 文档                | Link 链接                                             | Primary Audience 主要受众        | [Normative 规范性](https://google.github.io/styleguide/go/#normative) | [Canonical 权威性](https://google.github.io/styleguide/go/#canonical) |
| ---------------------------- | ----------------------------------------------------- | -------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **Style Guide 风格指南**     | https://google.github.io/styleguide/go/guide          | Everyone 每个人                  | Yes                                                          | Yes                                                          |
| **Style Decisions 风格决策** | https://google.github.io/styleguide/go/decisions      | Readability Mentors 可读性导师   | Yes                                                          | No                                                           |
| **Best Practices 最佳实践**  | https://google.github.io/styleguide/go/best-practices | Anyone interested 任何感兴趣的人 | No                                                           | No                                                           |



### Documents 文档

1. The **[Style Guide](https://google.github.io/styleguide/go/guide)** outlines the foundation of Go style at Google. This document is definitive and is used as the basis for the recommendations in Style Decisions and Best Practices.

   ​	风格指南概述了 Google 中 Go 风格的基础。此文档是明确的，并用作风格决策和最佳实践中建议的基础。

2. **[Style Decisions](https://google.github.io/styleguide/go/decisions)** is a more verbose document that summarizes decisions on specific style points and discusses the reasoning behind the decisions where appropriate.

   ​	风格决策是一份更冗长的文档，它总结了有关特定风格点的决策，并在适当的情况下讨论了决策背后的原因。

   These decisions may occasionally change based on new data, new language features, new libraries, or emerging patterns, but it is not expected that individual Go programmers at Google should keep up-to-date with this document.

   ​	这些决策可能会根据新数据、新语言特性、新库或新兴模式偶尔发生变化，但预计 Google 中的各个 Go 程序员不必随时了解此文档。

3. **[Best Practices](https://google.github.io/styleguide/go/best-practices)** documents some of the patterns that have evolved over time that solve common problems, read well, and are robust to code maintenance needs.

   ​	最佳实践记录了一些随着时间推移而演变的模式，这些模式解决了常见问题，可读性好，并且对代码维护需求具有鲁棒性。

   These best practices are not canonical, but Go programmers at Google are encouraged to use them where possible to keep the codebase uniform and consistent.

   ​	这些最佳实践并非规范，但鼓励 Google 的 Go 程序员尽可能使用它们，以保持代码库的统一和一致性。

These documents intend to:

​	这些文档旨在：

- Agree on a set of principles for weighing alternate styles
  就权衡备选风格达成一套原则
- Codify settled matters of Go style
  编纂已解决的 Go 风格问题
- Document and provide canonical examples for Go idioms
  记录并提供 Go 惯用法的规范示例
- Document the pros and cons of various style decisions
  记录各种风格决策的优缺点
- Help minimize surprises in Go readability reviews
  帮助最大程度地减少 Go 可读性审查中的意外情况
- Help readability mentors use consistent terminology and guidance
  帮助可读性导师使用一致的术语和指导

These documents do **not** intend to:

​	这些文档无意于：

- Be an exhaustive list of comments that can be given in a readability review
  成为可读性审查中可以给出的评论的详尽列表
- List all of the rules everyone is expected to remember and follow at all times
  列出每个人在任何时候都应该记住并遵循的所有规则
- Replace good judgment in the use of language features and style
  取代在使用语言特性和样式时的良好判断
- Justify large-scale changes to get rid of style differences
  证明进行大规模更改以消除样式差异是合理的

There will always be differences from one Go programmer to another and from one team’s codebase to another. However, it is in the best interest of Google and Alphabet that our codebase be as consistent as possible. (See [guide](https://google.github.io/styleguide/go/guide#consistency) for more on consistency.) To that end, feel free to make style improvements as you see fit, but you do not need to nit-pick every violation of the Style Guide that you find. In particular, these documents may change over time, and that is no reason to cause extra churn in existing codebases; it suffices to write new code using the latest best practices and address nearby issues over time.

​	不同的 Go 程序员和不同的团队代码库之间总是存在差异。但是，我们的代码库尽可能保持一致最符合 Google 和 Alphabet 的利益。（有关一致性的更多信息，请参阅指南。）为此，您可以随时根据需要进行样式改进，但您不必对您发现的每个违反样式指南的行为吹毛求疵。特别是，这些文档可能会随着时间而改变，这并不是在现有代码库中造成额外混乱的理由；只需使用最新的最佳实践编写新代码，并随着时间的推移解决附近的问题即可。

It is important to recognize that issues of style are inherently personal and that there are always inherent trade-offs. Much of the guidance in these documents is subjective, but just like with `gofmt`, there is significant value in the uniformity they provide. As such, style recommendations will not be changed without due discourse, Go programmers at Google are encouraged to follow the style guide even where they might disagree.

​	重要的是要认识到风格问题本质上是个人问题，并且总是存在固有的权衡。这些文档中的许多指导都是主观的，但就像 `gofmt` 一样，它们提供的统一性具有重大价值。因此，在没有适当讨论的情况下，不会更改风格建议，鼓励谷歌的 Go 程序员遵循风格指南，即使他们可能不同意。



## Definitions 定义

The following words, which are used throughout the style documents, are defined below:

​	以下在整个风格文档中使用的单词在下面定义：

- **Canonical**: Establishes prescriptive and enduring rules

  ​	权威性：建立规范性和持久性规则

  Within these documents, “canonical” is used to describe something that is considered a standard that all code (old and new) should follow and that is not expected to change substantially over time. Principles in the canonical documents should be understood by authors and reviewers alike, so everything included within a canonical document must meet a high bar. As such, canonical documents are generally shorter and prescribe fewer elements of style than non-canonical documents.

  ​	在这些文档中，“规范”用于描述被认为是所有代码（旧代码和新代码）都应遵循的标准，并且预计不会随着时间的推移而发生重大变化。规范文档中的原则应为作者和审阅者所理解，因此规范文档中包含的所有内容都必须达到很高的标准。因此，规范文档通常较短，并且比非规范文档规定更少的风格元素。

  https://google.github.io/styleguide/go#canonical

- **Normative**: Intended to establish consistency

  ​	规范性：旨在建立一致性

  Within these documents, “normative” is used to describe something that is an agreed-upon element of style for use by Go code reviewers, in order that the suggestions, terminology, and justifications are consistent. These elements may change over time, and these documents will reflect such changes so that reviewers can remain consistent and up-to-date. Authors of Go code are not expected to be familiar with the normative documents, but the documents will frequently be used as a reference by reviewers in readability reviews.

  ​	在这些文档中，“规范”用于描述 Go 代码审阅者使用的一致风格的约定元素，以便建议、术语和理由保持一致。这些元素可能会随着时间的推移而发生变化，这些文档将反映这些变化，以便审阅者能够保持一致和最新。Go 代码的作者不必熟悉规范性文档，但审阅者在可读性审阅中经常会将这些文档用作参考。

  https://google.github.io/styleguide/go#normative

- **Idiomatic**: Common and familiar

  ​	惯用：常见且熟悉

  Within these documents, “idiomatic” is used to refer to something that is prevalent in Go code and has become a familiar pattern that is easy to recognize. In general, an idiomatic pattern should be preferred to something unidiomatic if both serve the same purpose in context, as this is what will be the most familiar to readers.

  ​	在这些文档中，“惯用”用于指代 Go 代码中普遍存在且已成为易于识别的熟悉模式的内容。一般来说，如果惯用模式和非惯用模式在上下文中具有相同的作用，则应优先选择惯用模式，因为这对于读者来说最熟悉。

  https://google.github.io/styleguide/go#idiomatic



## Additional references 其他参考

This guide assumes the reader is familiar with [Effective Go](https://go.dev/doc/effective_go), as it provides a common baseline for Go code across the entire Go community.

​	本指南假定读者熟悉《高效的 Go》，因为它为整个 Go 社区中的 Go 代码提供了共同的基础线。

Below are some additional resources for those looking to self-educate about Go style and for reviewers looking to provide further linkable context in their reviews. Participants in the Go readability process are not expected to be familiar with these resources, but they may arise as context in readability reviews.

​	以下是一些其他资源，供那些希望自学 Go 风格的人以及希望在他们的评论中提供更多可链接上下文的审阅者使用。Go 可读性流程的参与者不必熟悉这些资源，但它们可能会作为可读性评论中的上下文出现。

**External References
外部参考**

- [Go Language Specification
  Go 语言规范](https://go.dev/ref/spec)

- [Go FAQ
  Go 常见问题解答](https://go.dev/doc/faq)

- [Go Memory Model
  Go 内存模型](https://go.dev/ref/mem)

- [Go Data Structures
  Go 数据结构](https://research.swtch.com/godata)

- [Go Interfaces
  Go 接口](https://research.swtch.com/interfaces)

- [Go Proverbs
  Go 格言](https://go-proverbs.github.io/)

- Go Tip Episodes - stay tuned.

  ​	Go Tip 剧集 - 敬请期待。

- Unit Testing Practices - stay tuned.
  单元测试实践 - 敬请期待。

**Relevant Testing-on-the-Toilet articles
相关测试文章**

- [TotT: Identifier Naming
  TotT：标识符命名](https://testing.googleblog.com/2017/10/code-health-identifiernamingpostforworl.html)
- [TotT: Testing State vs. Testing Interactions
  TotT：测试状态与测试交互](https://testing.googleblog.com/2013/03/testing-on-toilet-testing-state-vs.html)
- [TotT: Effective Testing
  TotT：有效测试](https://testing.googleblog.com/2014/05/testing-on-toilet-effective-testing.html)
- [TotT: Risk-driven Testing
  TotT：风险驱动测试](https://testing.googleblog.com/2014/05/testing-on-toilet-risk-driven-testing.html)
- [TotT: Change-detector Tests Considered Harmful
  TotT：认为变更检测器测试有害](https://testing.googleblog.com/2015/01/testing-on-toilet-change-detector-tests.html)

**Additional External Writings
其他外部文章**

- [Go and Dogma
  Go 和教条](https://research.swtch.com/dogma)
- [Less is exponentially more
  少即是多](https://commandcenter.blogspot.com/2012/06/less-is-exponentially-more.html)
- [Esmerelda’s Imagination
  埃斯梅雷达的想象力](https://commandcenter.blogspot.com/2011/12/esmereldas-imagination.html)
- [Regular expressions for parsing
  用于解析的正则表达式](https://commandcenter.blogspot.com/2011/08/regular-expressions-in-lexing-and.html)
- [Gofmt’s style is no one’s favorite, yet Gofmt is everyone’s favorite](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=8m43s) (YouTube)
  Gofmt 的风格不是任何人的最爱，但 Gofmt 是每个人的最爱 (YouTube)