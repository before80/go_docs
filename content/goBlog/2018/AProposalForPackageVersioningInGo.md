+++
title = "关于 Go 中包的版本管理的建议"
weight = 14
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# A Proposal for Package Versioning in Go - 关于 go 中包的版本管理的建议

> 原文：[https://go.dev/blog/versioning-proposal](https://go.dev/blog/versioning-proposal)

Russ Cox
26 March 2018

## Introduction 引言

Eight years ago, the Go team introduced `goinstall` (which led to `go get`) and with it the decentralized, URL-like import paths that Go developers are familiar with today. After we released `goinstall`, one of the first questions people asked was how to incorporate version information. We admitted we didn’t know. For a long time, we believed that the problem of package versioning would be best solved by an add-on tool, and we encouraged people to create one. The Go community created many tools with different approaches. Each one helped us all better understand the problem, but by mid-2016 it was clear that there were now too many solutions. We needed to adopt a single, official tool.

八年前，Go团队推出了goinstall（后来又推出了go get），随之而来的是分散的、类似于URL的导入路径，这也是Go开发者今天所熟悉的。在我们发布goinstall之后，人们问的第一个问题就是如何纳入版本信息。我们承认我们不知道。在很长一段时间里，我们认为软件包的版本问题最好由一个附加工具来解决，并且我们鼓励人们去创造一个。Go社区创造了许多具有不同方法的工具。每一个工具都帮助我们更好地理解这个问题，但到2016年年中，很明显，现在有太多的解决方案了。我们需要采用一个单一的、官方的工具。

After a community discussion started at GopherCon in July 2016 and continuing into the fall, we all believed the answer would be to follow the package versioning approach exemplified by Rust’s Cargo, with tagged semantic versions, a manifest, a lock file, and a [SAT solver](https://research.swtch.com/version-sat) to decide which versions to use. Sam Boyer led a team to create Dep, which followed this rough plan, and which we intended to serve as the model for `go` command integration. But as we learned more about the implications of the Cargo/Dep approach, it became clear to me that Go would benefit from changing some of the details, especially concerning backwards compatibility.

在2016年7月GopherCon开始的社区讨论后，一直持续到秋天，我们都认为答案是遵循Rust的Cargo所示范的软件包版本管理方法，有标记的语义版本、清单、锁文件和SAT解算器来决定使用哪些版本。Sam Boyer带领一个团队创建了Dep，它遵循这个粗略的计划，我们打算将其作为go命令集成的模型。但随着我们对Cargo/Dep方法的影响有了更多的了解，我清楚地认识到，Go将从改变一些细节中受益，特别是关于向后兼容性的细节。

## The Impact of Compatibility 兼容性的影响

The most important new feature of [Go 1](https://blog.golang.org/preview-of-go-version-1) was not a language feature. It was Go 1’s emphasis on backwards compatibility. Until that point we’d issued stable release snapshots approximately monthly, each with significant incompatible changes. We observed significant acceleration in interest and adoption immediately after the release of Go 1. We believe that the [promise of compatibility](https://go.dev/doc/go1compat.html) made developers feel much more comfortable relying on Go for production use and is a key reason that Go is popular today. Since 2013 the [Go FAQ](https://go.dev/doc/faq#get_version) has encouraged package developers to provide their own users with similar expectations of compatibility. We call this the *import compatibility rule*: "If an old package and a new package have the same import path, the new package must be backwards compatible with the old package."

Go 1 的最重要的新特性不是语言特性。它是 Go 1 对向后兼容性的强调。在这之前，我们大约每个月都会发布稳定版快照，每次都有重大的不兼容变化。我们观察到，在 Go 1 发布后，人们对 Go 1 的兴趣和采用速度明显加快。我们相信，兼容性的承诺使开发人员在生产中更容易依赖 Go，这也是 Go 今天流行的一个重要原因。自2013年以来，Go FAQ鼓励软件包开发者为自己的用户提供类似的兼容性期望。我们称之为导入兼容性规则："如果一个旧包和一个新包有相同的导入路径，新包必须向后兼容旧包"。

Independently, [semantic versioning](http://semver.org/) has become the *de facto* standard for describing software versions in many language communities, including the Go community. Using semantic versioning, later versions are expected to be backwards-compatible with earlier versions, but only within a single major version: v1.2.3 must be compatible with v1.2.1 and v1.1.5, but v2.3.4 need not be compatible with any of those.

独立地讲，语义版本学已经成为许多语言社区（包括Go社区）描述软件版本的事实上的标准。使用语义版本学，后来的版本有望向后兼容早期的版本，但只在一个主要的版本内：v1.2.3必须与v1.2.1和v1.1.5兼容，但是v2.3.4不需要与其中任何一个兼容。

If we adopt semantic versioning for Go packages, as most Go developers expect, then the import compatibility rule requires that different major versions must use different import paths. This observation led us to *semantic import versioning*, in which versions starting at v2.0.0 include the major version in the import path: `my/thing/v2/sub/pkg`.

如果我们像大多数Go开发者所期望的那样，为Go包采用语义版本划分，那么导入兼容性规则要求不同的主要版本必须使用不同的导入路径。这个观察结果导致我们采用了语义导入版本管理，在这个版本中，从v2.0.0开始的版本在导入路径中包括主要版本：my/thing/v2/sub/pkg。

A year ago I strongly believed that whether to include version numbers in import paths was largely a matter of taste, and I was skeptical that having them was particularly elegant. But the decision turns out to be a matter not of taste but of logic: import compatibility and semantic versioning together require semantic import versioning. When I realized this, the logical necessity surprised me.

一年前，我坚信，是否在导入路径中包含版本号在很大程度上是一个品味问题，而且我对拥有它们是否特别优雅持怀疑态度。但是，事实证明，这个决定不是一个品味问题，而是一个逻辑问题：导入兼容性和语义版本化共同要求语义导入版本化。当我意识到这一点的时候，逻辑上的必要性让我感到惊讶。

I was also surprised to realize that there is a second, independent logical route to semantic import versioning: [gradual code repair](https://go.dev/talks/2016/refactor.article) or partial code upgrades. In a large program, it’s unrealistic to expect all packages in the program to update from v1 to v2 of a particular dependency at the same time. Instead, it must be possible for some of the program to keep using v1 while other parts have upgraded to v2. But then the program’s build, and the program’s final binary, must include both v1 and v2 of the dependency. Giving them the same import path would lead to confusion, violating what we might call the *import uniqueness rule*: different packages must have different import paths. The only way to have partial code upgrades, import uniqueness, *and* semantic versioning is to adopt semantic import versioning as well.

我还惊讶地意识到，还有第二条独立的逻辑路线来实现语义导入版本化：逐步的代码修复或部分代码升级。在一个大型程序中，期望程序中的所有软件包都同时从某个特定的依赖关系的v1版更新到v2版是不现实的。相反，必须让程序的某些部分继续使用v1版，而其他部分则升级到v2版。 但这样一来，程序的构建和程序的最终二进制文件必须同时包括依赖关系的v1版和v2版。给它们提供相同的导入路径会导致混乱，违反了我们可以称之为导入唯一性的规则：不同的包必须有不同的导入路径。拥有部分代码升级、导入唯一性和语义版本控制的唯一方法就是同时采用语义导入版本控制。

It is of course possible to build systems that use semantic versioning without semantic import versioning, but only by giving up either partial code upgrades or import uniqueness. Cargo allows partial code upgrades by giving up import uniqueness: a given import path can have different meanings in different parts of a large build. Dep ensures import uniqueness by giving up partial code upgrades: all packages involved in a large build must find a single agreed-upon version of a given dependency, raising the possibility that large programs will be unbuildable. Cargo is right to insist on partial code upgrades, which are critical to large-scale software development. Dep is equally right to insist on import uniqueness. Complex uses of Go’s current vendoring support can violate import uniqueness. When they have, the resulting problems have been quite challenging for both developers and tools to understand. Deciding between partial code upgrades and import uniqueness requires predicting which will hurt more to give up. Semantic import versioning lets us avoid the choice and keep both instead.

当然，在没有语义导入版本管理的情况下，也可以构建使用语义版本管理的系统，但是只有通过放弃部分代码升级或者导入唯一性的方式才行。Cargo通过放弃导入的唯一性而允许部分代码升级：一个给定的导入路径在一个大型构建的不同部分可能具有不同的含义。Dep通过放弃部分代码升级来确保导入的唯一性：所有参与大型构建的软件包都必须找到一个特定依赖的单一约定的版本，这就提高了大型程序无法构建的可能性。Cargo坚持部分代码升级是正确的，这对大规模软件开发至关重要。Dep坚持导入的唯一性也是正确的。Go当前的销售支持的复杂使用可能会违反导入的唯一性。当它们发生时，所产生的问题对于开发者和工具来说都是相当具有挑战性的。在部分代码升级和导入唯一性之间做出决定，需要预测放弃哪一个会更痛苦。语义导入版本管理让我们避免了这种选择，而将两者都保留下来。

I was also surprised to discover how much import compatibility simplifies version selection, which is the problem of deciding which package versions to use for a given build. The constraints of Cargo and Dep make version selection equivalent to [solving Boolean satisfiability](https://research.swtch.com/version-sat), meaning it can be very expensive to determine whether a valid version configuration even exists. And then there may be many valid configurations, with no clear criteria for choosing the "best" one. Relying on import compatibility can instead let Go use a trivial, linear-time algorithm to find the single best configuration, which always exists. This algorithm, which I call [*minimal version selection*](https://research.swtch.com/vgo-mvs), in turn eliminates the need for separate lock and manifest files. It replaces them with a single, short configuration file, edited directly by both developers and tools, that still supports reproducible builds.

我还惊讶地发现，导入兼容性在多大程度上简化了版本选择，也就是决定在特定构建中使用哪些软件包版本的问题。Cargo和Dep的约束使版本选择等同于解决布尔可满足性问题，这意味着要确定一个有效的版本配置是否存在是非常昂贵的。然后，可能会有许多有效的配置，但没有明确的标准来选择 "最佳 "的配置。依靠导入兼容性可以让Go使用一个微不足道的线性时间算法来找到单一的最佳配置，这个配置总是存在的。这种算法，我称之为最小版本选择，反过来又消除了对单独的锁和清单文件的需求。它用一个简短的配置文件取代了它们，由开发人员和工具直接编辑，仍然支持可重复的构建。

Our experience with Dep demonstrates the impact of compatibility. Following the lead of Cargo and earlier systems, we designed Dep to give up import compatibility as part of adopting semantic versioning. I don’t believe we decided this deliberately; we just followed those other systems. The first-hand experience of using Dep helped us better understand exactly how much complexity is created by permitting incompatible import paths. Reviving the import compatibility rule by introducing semantic import versioning eliminates that complexity, leading to a much simpler system.

我们在Dep方面的经验表明了兼容性的影响。按照Cargo和早期系统的做法，我们设计Dep时放弃了进口兼容性，作为采用语义版本的一部分。我不相信我们是故意这样决定的；我们只是跟随那些其他系统。使用Dep的第一手经验帮助我们更好地理解，允许不兼容的导入路径到底会造成多大的复杂性。通过引入语义导入版本，恢复了导入兼容性规则，消除了这种复杂性，导致了一个更简单的系统。

## Progress, a Prototype, and a Proposal 进展、原型和建议

Dep was released in January 2017. Its basic model—code tagged with semantic versions, along with a configuration file that specified dependency requirements—was a clear step forward from most of the Go vendoring tools, and converging on Dep itself was also a clear step forward. I wholeheartedly encouraged its adoption, especially to help developers get used to thinking about Go package versions, both for their own code and their dependencies. While Dep was clearly moving us in the right direction, I had lingering concerns about the complexity devil in the details. I was particularly concerned about Dep lacking support for gradual code upgrades in large programs. Over the course of 2017, I talked to many people, including Sam Boyer and the rest of the package management working group, but none of us could see any clear way to reduce the complexity. (I did find many approaches that added to it.) Approaching the end of the year, it still seemed like SAT solvers and unsatisfiable builds might be the best we could do.

Dep于2017年1月发布。它的基本模型--用语义版本标记的代码，以及指定依赖项要求的配置文件--与大多数Go销售工具相比，是一个明显的进步，而汇聚在Dep本身也是一个明显的进步。我全心全意地鼓励采用它，特别是帮助开发人员习惯于考虑Go软件包的版本，包括他们自己的代码和他们的依赖关系。虽然Dep显然在朝着正确的方向发展，但我对细节中的复杂性问题仍有担忧。我特别担心Dep缺乏对大型程序中的渐进式代码升级的支持。在2017年的过程中，我和很多人谈过，包括Sam Boyer和软件包管理工作组的其他成员，但我们都看不到任何明确的方法来减少复杂性。(我确实找到了许多增加复杂性的方法。)在接近年底时，似乎SAT求解器和不可满足的构建可能是我们能做的最好的。

In mid-November, trying once again to work through how Dep could support gradual code upgrades, I realized that our old advice about import compatibility implied semantic import versioning. That seemed like a real breakthrough. I wrote a first draft of what became my [semantic import versioning](https://research.swtch.com/vgo-import) blog post, concluding it by suggesting that Dep adopt the convention. I sent the draft to the people I’d been talking to, and it elicited very strong responses: everyone loved it or hated it. I realized that I needed to work out more of the implications of semantic import versioning before circulating the idea further, and I set out to do that.

11月中旬，我再次尝试解决Dep如何支持渐进式代码升级的问题，我意识到，我们以前关于导入兼容性的建议意味着语义导入版本化。这似乎是一个真正的突破。我写了一篇后来成为我的语义导入版本管理博文的初稿，最后建议Dep采用这一惯例。我把这篇博文的初稿发给了与我交谈过的人，结果引起了非常强烈的反应：每个人都喜欢它，或者讨厌它。我意识到，在进一步传播这个想法之前，我需要弄清楚语义导入版本控制的更多含义，于是我开始着手进行这项工作。

In mid-December, I discovered that import compatibility and semantic import versioning together allowed cutting version selection down to [minimal version selection](https://research.swtch.com/vgo-mvs). I wrote a basic implementation to be sure I understood it, I spent a while learning the theory behind why it was so simple, and I wrote a draft of the post describing it. Even so, I still wasn’t sure the approach would be practical in a real tool like Dep. It was clear that a prototype was needed.

12月中旬，我发现，进口兼容性和语义进口版本管理一起允许将版本选择削减到最小的版本选择。我写了一个基本的实现，以确保我理解它，我花了一段时间来学习为什么它如此简单背后的理论，并且我写了一个描述它的帖子的草稿。即便如此，我仍然不确定这种方法在像Dep这样的真正的工具中是否实用，显然需要一个原型。

In January, I started work on a simple `go` command wrapper that implemented semantic import versioning and minimal version selection. Trivial tests worked well. Approaching the end of the month, my simple wrapper could build Dep, a real program that made use of many versioned packages. The wrapper still had no command-line interface—the fact that it was building Dep was hard-coded in a few string constants—but the approach was clearly viable.

在一月份，我开始了一个简单的go命令包装器的工作，它实现了语义导入的版本管理和最小的版本选择。琐碎的测试效果不错。到了月底，我的简单包装器可以构建Dep，一个使用许多版本包的真实程序。这个包装器仍然没有命令行界面--它正在构建Dep的事实是在几个字符串常量中硬编码的--但这个方法显然是可行的。

I spent the first three weeks of February turning the wrapper into a full versioned `go` command, `vgo`; writing drafts of a [blog post series introducing `vgo`](https://research.swtch.com/vgo); and discussing them with Sam Boyer, the package management working group, and the Go team. And then I spent the last week of February finally sharing `vgo` and the ideas behind it with the whole Go community.

我在二月份的前三周里，把包装器变成了一个完整的版本化go命令，即vgo；写了一系列介绍vgo的博文草稿；并与Sam Boyer、软件包管理工作组和Go团队讨论了这些问题。然后我花了2月的最后一个星期，终于与整个Go社区分享了vgo和它背后的想法。

In addition to the core ideas of import compatibility, semantic import versioning, and minimal version selection, the `vgo` prototype introduces a number of smaller but significant changes motivated by eight years of experience with `goinstall` and `go get`: the new concept of a [Go module](https://research.swtch.com/vgo-module), which is a collection of packages versioned as a unit; [verifiable and verified builds](https://research.swtch.com/vgo-repro); and [version-awareness throughout the `go` command](https://research.swtch.com/vgo-cmd), enabling work outside `$GOPATH` and the elimination of (most) `vendor` directories.

除了导入兼容性、语义导入版本和最小化版本选择的核心思想外，vgo原型还引入了一些较小但重要的变化，这些变化是由八年来goinstall和go get的经验促成的：Go模块的新概念，它是作为一个单元的软件包的集合；可验证和验证的构建；以及整个go命令的版本意识，使工作在$GOPATH之外并消除了（大多数）供应商目录。

The result of all of this is the [official Go proposal](https://go.dev/design/24301-versioned-go), which I filed last week. Even though it might look like a complete implementation, it’s still just a prototype, one that we will all need to work together to complete. You can download and try the `vgo` prototype from [golang.org/x/vgo](https://golang.org/x/vgo), and you can read the [Tour of Versioned Go](https://research.swtch.com/vgo-tour) to get a sense of what using `vgo` is like.

所有这些的结果就是我上周提交的官方Go提案。尽管它看起来像是一个完整的实现，但它仍然只是一个原型，一个需要我们共同完成的原型。您可以从golang.org/x/vgo下载并试用vgo原型，您也可以阅读《版本化Go之旅》以了解使用vgo的情况。

## The Path Forward 前进的道路

The proposal I filed last week is exactly that: an initial proposal. I know there are problems with it that the Go team and I can’t see, because Go developers use Go in many clever ways that we don’t know about. The goal of the proposal feedback process is for us all to work together to identify and address the problems in the current proposal, to make sure that the final implementation that ships in a future Go release works well for as many developers as possible. Please point out problems on the [proposal discussion issue](https://go.dev/issue/24301). I will keep the [discussion summary](https://go.dev/issue/24301#issuecomment-371228742) and [FAQ](https://go.dev/issue/24301#issuecomment-371228664) updated as feedback arrives.

我上周提交的提案正是如此：一个初步提案。我知道它有一些Go团队和我都看不到的问题，因为Go开发者以许多我们不知道的巧妙方式使用Go。提案反馈过程的目的是让我们大家一起努力，找出并解决当前提案中的问题，以确保在未来的Go版本中推出的最终实现能够为尽可能多的开发者带来良好的效果。请在提案的讨论问题上指出问题。随着反馈的到来，我将不断更新讨论摘要和常见问题。

For this proposal to succeed, the Go ecosystem as a whole—and in particular today’s major Go projects—will need to adopt the import compatibility rule and semantic import versioning. To make sure that can happen smoothly, we will also be conducting user feedback sessions by video conference with projects that have questions about how to incorporate the new versioning proposal into their code bases or have feedback about their experiences. If you are interested in participating in such a session, please email Steve Francia at spf@golang.org.

为了使这项提议取得成功，整个Go生态系统，特别是今天的主要Go项目，都需要采用导入兼容性规则和语义导入版本控制。为了确保这种情况能够顺利发生，我们还将通过视频会议与那些对如何将新的版本管理建议纳入其代码库有疑问的项目进行用户反馈会议，或者对其经验进行反馈。如果您有兴趣参加这样的会议，请发邮件给Steve Francia，spf@golang.org。

We’re looking forward to (finally!) providing the Go community with a single, official answer to the question of how to incorporate package versioning into `go get`. Thanks to everyone who helped us get this far, and to everyone who will help us going forward. We hope that, with your help, we can ship something that Go developers will love.

我们期待着（最终！）为Go社区提供一个单一的、官方的答案，以解决如何将包的版本管理纳入Go get中的问题。感谢每一个帮助我们走到这一步的人，也感谢每一个将帮助我们前进的人。我们希望，在您们的帮助下，我们能够推出Go开发者喜欢的东西。
