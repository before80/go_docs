+++
title = "Go, 开源, 社区"
weight = 6
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go, Open Source, Community - Go, 开源, 社区

> 原文：[https://go.dev/blog/open-source](https://go.dev/blog/open-source)

Russ Cox
8 July 2015

## Welcome 欢迎

[This is the text of my opening keynote at Gophercon 2015. [The video is available here](https://www.youtube.com/watch?v=XvZOdpd_9tc).]

[这是我在2015年Gophercon大会上的开幕式主题演讲的内容。视频可以在这里找到。］

Thank you all for traveling to Denver to be here, and thank you to everyone watching on video. If this is your first Gophercon, welcome. If you were here last year, welcome back. Thank you to the organizers for all the work it takes to make a conference like this happen. I am thrilled to be here and to be able to talk to all of you.

感谢大家来到丹佛参加会议，也感谢通过视频观看的所有人。如果这是您第一次参加Gophercon，欢迎。如果您去年来过这里，欢迎回来。感谢组织者为实现这样的会议所做的所有工作。我很高兴来到这里，并能与您们所有人交谈。

I am the tech lead for the Go project and the Go team at Google. I share that role with Rob Pike. In that role, I spend a lot of time thinking about the overall Go open source project, in particular the way it runs, what it means to be open source, and the interaction between contributors inside and outside Google. Today I want to share with you how I see the Go project as a whole and then based on that explain how I see the Go open source project evolving.

我是谷歌的Go项目和Go团队的技术负责人。我和罗伯-派克共同担任这一角色。在这个职位上，我花了很多时间来思考整个Go开源项目，特别是它的运行方式，开源的意义，以及谷歌内部和外部贡献者之间的互动。今天我想和大家分享一下我对Go项目的整体看法，然后在此基础上解释我对Go开源项目的发展看法。

## Why Go? 为什么是Go？

To get started, we have to go back to the beginning. Why did we start working on Go?

为了开始，我们必须回到起点。我们为什么要开始研究Go？

Go is an attempt to make programmers more productive. We wanted to improve the software development process at Google, but the problems Google has are not unique to Google.

Go是一个让程序员更有效率的尝试。我们想改善谷歌的软件开发过程，但谷歌的问题并不是谷歌独有的。

There were two overarching goals.

有两个首要的目标。

The first goal is to make a better language to meet the challenges of scalable concurrency. By scalable concurrency I mean software that deals with many concerns simultaneously, such as coordinating a thousand back end servers by sending network traffic back and forth.

第一个目标是做一个更好的语言来应对可扩展并发的挑战。我所说的可扩展并发性是指同时处理许多问题的软件，例如通过来回发送网络流量来协调一千个后端服务器。

Today, that kind of software has a shorter name: we call it cloud software. It’s fair to say that Go was designed for the cloud before clouds ran software.

今天，这种软件有一个更简短的名字：我们称之为云软件。可以说，在云计算运行软件之前，Go是为云计算设计的。

The larger goal is to make a better environment to meet the challenges of scalable software development, software worked on and used by many people, with limited coordination between them, and maintained for years. At Google we have thousands of engineers writing and sharing their code with each other, trying to get their work done, reusing the work of others as much as possible, and working in a code base with a history dating back over ten years. Engineers often work on or at least look at code originally written by someone else, or that they wrote years ago, which often amounts to the same thing.

更大的目标是制造一个更好的环境，以应对可扩展软件开发的挑战，这些软件由许多人工作和使用，他们之间的协调有限，并维护多年。在谷歌，我们有成千上万的工程师在写代码，并相互分享他们的代码，努力完成他们的工作，尽可能地重复使用其他人的工作，并在一个历史可以追溯到十年以上的代码库中工作。工程师们经常工作在或至少看一看最初由别人写的代码，或他们多年前写的代码，这往往相当于同样的事情。

That situation inside Google has a lot in common with large scale, modern open source development as practiced on sites like GitHub. Because of this, Go is a great fit for open source projects, helping them accept and manage contributions from a large community over a long period of time.

谷歌内部的这种情况与GitHub等网站上实行的大规模现代开放源代码开发有很多共同之处。正因为如此，Go很适合开源项目，帮助它们接受和管理来自大型社区的长期贡献。

I believe much of Go’s success is explained by the fact that Go is a great fit for cloud software, Go is a great fit for open source projects, and, serendipitously, both of those are growing in popularity and importance in the software industry.

我相信Go的成功在很大程度上可以解释为：Go非常适合云计算软件，Go非常适合开源项目，而且，偶然的是，这两样东西在软件行业都越来越受欢迎，越来越重要。

Other people have made similar observations. Here are two. Last year, on RedMonk.com, Donnie Berkholz wrote about "[Go as the emerging language of cloud infrastructure](http://redmonk.com/dberkholz/2014/03/18/go-the-emerging-language-of-cloud-infrastructure/)," observing that "[Go’s] marquee projects … are cloud-centric or otherwise made for dealing with distributed systems or transient environments."

其他人也提出了类似的看法。这里有两个。去年，在RedMonk.com上，Donnie Berkholz写到 "Go是新兴的云基础设施语言"，他观察到"[Go的]主要项目......是以云为中心的，或者是为处理分布式系统或瞬时环境而做的。"

This year, on Texlution.com, the author wrote an article titled "[Why Golang is doomed to succeed](https://texlution.com/post/why-go-is-doomed-to-succeed/)," pointing out that this focus on large-scale development was possibly even better suited to open source than to Google itself: "This open source fitness is why I think you are about to see more and more Go around …"

今年，在Texlution.com上，作者写了一篇题为 "为什么Golang注定要成功 "的文章，指出这种对大规模开发的关注可能甚至比谷歌本身更适合于开源。"这种开源的健身方式是我认为您即将看到周围越来越多的Go的原因..."

## The Go Balance - Go 的平衡

How does Go accomplish those things?

Go是如何完成这些事情的？

How does it make scalable concurrency and scalable software development easier?

它是如何使可扩展的并发性和可扩展的软件开发更容易的？

Most people answer this question by talking about channels and goroutines, and interfaces, and fast builds, and the go command, and good tool support. Those are all important parts of the answer, but I think there is a broader idea behind them.

大多数人在回答这个问题时都会谈论通道、goroutines、接口、快速构建、go命令和良好的工具支持。这些都是答案的重要部分，但我认为它们背后有一个更广泛的想法。

I think of that idea as Go’s balance. There are competing concerns in any software design, and there is a very natural tendency to try to solve all the problems you foresee. In Go, we have explicitly tried not to solve everything. Instead, we’ve tried to do just enough that you can build your own custom solutions easily.

我认为这个想法就是Go的平衡。在任何软件设计中都存在着相互竞争的问题，而且有一种非常自然的倾向，就是试图解决您所预见的所有问题。在Go中，我们明确地尝试不解决所有问题。相反，我们试图做得足够多，以便您可以轻松建立自己的定制解决方案。

The way I would summarize Go’s chosen balance is this: **Do Less. Enable More.**

我对Go所选择的平衡方式的总结是这样的。做得更少。启用更多。

Do less, but enable more.

做得少，但能做得更多。

Go can’t do everything. We shouldn’t try. But if we work at it, Go can probably do a few things well. If we select those things carefully, we can lay a foundation on which developers can *easily* build the solutions and tools they need, and ideally can interoperate with the solutions and tools built by others.

Go不可能做所有事情。我们不应该尝试。但是，如果我们努力工作，Go 可能会做好几件事。如果我们仔细选择这些事情，我们就可以打下一个基础，使开发人员可以轻松地建立他们所需要的解决方案和工具，并且最好能够与其他人建立的解决方案和工具进行互操作。

### Examples 例子

Let me illustrate this with some examples.

让我用一些例子来说明这一点。

First, the size of the Go language itself. We worked hard to put in as few concepts as possible, to avoid the problem of mutually incomprehensible dialects forming in different parts of a large developer community. No idea went into Go until it had been simplified to its essence and then had clear benefits that justified the complexity being added.

首先，Go语言本身的规模。我们努力将尽可能少的概念纳入其中，以避免在一个庞大的开发者社区的不同部分形成相互无法理解的方言问题。在Go语言中没有任何概念，直到它被简化到其本质，然后有明确的好处来证明所增加的复杂性是合理的。

In general, if we have 100 things we want Go to do well, we can’t make 100 separate changes. Instead, we try to research and understand the design space and then identify a few changes that work well together and that enable maybe 90 of those things. We’re willing to sacrifice the remaining 10 to avoid bloating the language, to avoid adding complexity only to address specific use cases that seem important today but might be gone tomorrow.

一般来说，如果我们有100件希望Go做得好的事情，我们不能做100个单独的改变。相反，我们试图研究和理解设计空间，然后找出一些能很好地配合的变化，这些变化可能能实现其中的90件事情。我们愿意牺牲剩下的10个，以避免语言的臃肿，避免仅仅为了解决今天看起来很重要但明天就可能消失的特定用例而增加复杂性。

Keeping the language small enables more important goals. Being small makes Go easier to learn, easier to understand, easier to implement, easier to reimplement, easier to debug, easier to adjust, and easier to evolve. Doing less enables more.

保持语言的小规模可以实现更重要的目标。小而精使Go更容易学习，更容易理解，更容易实现，更容易重新实现，更容易调试，更容易调整，更容易发展。少做多得。

I should point out that this means we say no to a lot of other people’s ideas, but I assure you we’ve said no to even more of our own ideas.

我应该指出，这意味着我们拒绝了很多别人的想法，但我向您保证，我们拒绝了更多自己的想法。

Next, channels and goroutines. How should we structure and coordinate concurrent and parallel computations? Mutexes and condition variables are very general but so low-level that they’re difficult to use correctly. Parallel execution frameworks like OpenMP are so high-level that they can only be used to solve a narrow range of problems. Channels and goroutines sit between these two extremes. By themselves, they aren’t a solution to much. But they are powerful enough to be easily arranged to enable solutions to many common problems in concurrent software. Doing less—really doing just enough—enables more.

接下来是通道和goroutines。我们应该如何构造和协调并发和并行计算？Mutexes和条件变量是非常通用的，但却非常低级，以至于难以正确使用。像OpenMP这样的并行执行框架是如此高级，以至于它们只能用于解决狭窄的问题。通道和goroutines位于这两个极端之间。就其本身而言，它们并不能解决很多问题。但是它们足够强大，可以很容易地被安排来解决并发软件中许多常见的问题。少做一点--真正做得足够多--就可以做得更多。

Next, types and interfaces. Having static types enables useful compile-time checking, something lacking in dynamically-typed languages like Python or Ruby. At the same time, Go’s static typing avoids much of the repetition of traditional statically typed languages, making it feel more lightweight, more like the dynamically-typed languages. This was one of the first things people noticed, and many of Go’s early adopters came from dynamically-typed languages.

接下来是类型和接口。拥有静态类型可以实现有用的编译时检查，这是Python或Ruby等动态类型语言所缺乏的。同时，Go的静态类型避免了传统静态类型语言的许多重复，使其感觉更轻盈，更像动态类型语言。这是人们最先注意到的一点，Go的许多早期采用者都来自动态类型的语言。

Go’s interfaces are a key part of that. In particular, omitting the ``implements'' declarations of Java or other languages with static hierarchy makes interfaces lighter weight and more flexible. Not having that rigid hierarchy enables idioms such as test interfaces that describe existing, unrelated production implementations. Doing less enables more.

Go的接口是其中的一个关键部分。特别是，省略了Java或其他具有静态层次结构的语言的 "实现 "声明，使得接口的重量更轻、更灵活。没有那种僵硬的层次结构，就会产生一些习惯性的做法，例如描述现有的、不相关的生产实现的测试接口。少做多得。

Next, testing and benchmarking. Is there any shortage of testing and benchmarking frameworks in most languages? Is there any agreement between them?

接下来是测试和基准测试。在大多数语言中，测试和基准测试框架有什么不足之处吗？他们之间有什么共识吗？

Go’s testing package is not meant to address every possible facet of these topics. Instead, it is meant to provide the basic concepts necessary for most higher-level tooling. Packages have test cases that pass, fail, or are skipped. Packages have benchmarks that run and can be measured by various metrics.

Go的测试包并不是要解决这些主题的每个可能的方面。相反，它是为了提供大多数高层工具所需的基本概念。测试包有通过、失败或被跳过的测试案例。软件包有运行的基准，可以用各种指标来衡量。

Doing less here is an attempt to reduce these concepts to their essence, to create a shared vocabulary so that richer tools can interoperate. That agreement enables higher-level testing software like Miki Tebeka’s go2xunit converter, or the benchcmp and benchstat benchmark analysis tools.

在这里做得比较少，是试图把这些概念减少到它们的本质，创建一个共享的词汇，以便更丰富的工具可以互操作。这种协议使更高层次的测试软件，如Miki Tebeka的go2xunit转换器，或 benchcmp和 benchstat基准分析工具。

Because there *is* agreement about the representation of the basic concepts, these higher-level tools work for all Go packages, not just ones that make the effort to opt in, and they interoperate with each other, in that using, say, go2xunit does not preclude also using benchstat, the way it would if these tools were, say, plugins for competing testing frameworks. Doing less enables more.

因为对基本概念的表述达成了一致，所以这些高级工具适用于所有Go软件包，而不仅仅是那些努力选择加入的软件包，而且它们可以相互操作，比如说，使用go2xunit并不妨碍同时使用benchstat，如果这些工具是竞争性测试框架的插件，就会出现这种情况。少做多得。

Next, refactoring and program analysis. Because Go is for large code bases, we knew it would need to support automatic maintenance and updating of source code. We also knew that this topic was too large to build in directly. But we knew one thing that we had to do. In our experience attempting automated program changes in other settings, the most significant barrier we hit was actually writing the modified program out in a format that developers can accept.

接下来是重构和程序分析。因为Go是针对大型代码库的，我们知道它需要支持源代码的自动维护和更新。我们也知道，这个话题太大，无法直接构建。但我们知道有一件事我们必须要做。根据我们在其他环境中尝试自动修改程序的经验，我们遇到的最重要的障碍是将修改后的程序实际写成开发者可以接受的格式。

In other languages, it’s common for different teams to use different formatting conventions. If an edit by a program uses the wrong convention, it either writes a section of the source file that looks nothing like the rest of the file, or it reformats the entire file, causing unnecessary and unwanted diffs.

在其他语言中，不同的团队使用不同的格式约定是很常见的。如果一个程序的编辑使用了错误的约定，它要么写出的源文件的一部分看起来与文件的其他部分完全不同，要么就会重新格式化整个文件，造成不必要的和不需要的差异。

Go does not have this problem. We designed the language to make gofmt possible, we worked hard to make gofmt’s formatting acceptable for all Go programs, and we made sure gofmt was there from day one of the original public release. Gofmt imposes such uniformity that automated changes blend into the rest of the file. You can’t tell whether a particular change was made by a person or a computer. We didn’t build explicit refactoring support. Establishing an agreed-upon formatting algorithm was enough of a shared base for independent tools to develop and to interoperate. Gofmt enabled gofix, goimports, eg, and other tools. I believe the work here is only just getting started. Even more can be done.

Go就没有这个问题。我们设计的语言使gofmt成为可能，我们努力使gofmt的格式化为所有Go程序所接受，并且我们确保gofmt从最初公开发布的第一天起就存在。Gofmt强加了这样的统一性，使自动化的修改与文件的其他部分融为一体。您无法分辨某个特定的修改是由人还是由计算机完成的。我们并没有建立明确的重构支持。建立一个约定俗成的格式化算法，就足以成为独立工具开发和互操作的共享基础。Gofmt使gofix、goimports、eg和其他工具成为可能。我相信这里的工作才刚刚开始。甚至可以做得更多。

Last, building and sharing software. In the run up to Go 1, we built goinstall, which became what we all know as "go get". That tool defined a standard zero-configuration way to resolve import paths on sites like github.com, and later a way to resolve paths on other sites by making HTTP requests. This agreed-upon resolution algorithm enabled other tools that work in terms of those paths, most notably Gary Burd’s creation of godoc.org. In case you haven’t used it, you go to godoc.org/the-import-path for any valid "go get" import path, and the web site will fetch the code and show you the documentation for it. A nice side effect of this has been that godoc.org serves as a rough master list of the Go packages publicly available. All we did was give import paths a clear meaning. Do less, enable more.

最后，建设和分享软件。在Go 1之前，我们建立了goinstall，也就是我们都知道的 "go get"。该工具定义了一种标准的零配置方式来解决github.com等网站的导入路径，后来又定义了一种通过HTTP请求来解决其他网站路径的方式。这种约定俗成的解析算法使其他工具能够根据这些路径工作，最值得一提的是Gary Burd创建的godoc.org。如果您没有使用过它，您可以到godoc.org/the-import-path去寻找任何有效的 "go get "导入路径，该网站会获取代码并向您展示它的文档。这样做的一个很好的副作用是，godoc.org可以作为公开的Go软件包的一个粗略的主列表。我们所做的就是给导入路径一个明确的含义。少做一点，多做一点。

You’ll notice that many of these tooling examples are about establishing a shared convention. Sometimes people refer to this as Go being "opinionated," but there’s something deeper going on. Agreeing to the limitations of a shared convention is a way to enable a broad class of tools that interoperate, because they all speak the same base language. This is a very effective way to do less but enable more. Specifically, in many cases we can do the minimum required to establish a shared understanding of a particular concept, like remote imports, or the proper formatting of a source file, and thereby enable the creation of packages and tools that work together because they all agree about those core details.

您会注意到这些工具实例中有许多是关于建立一个共享的惯例。有时人们会把这说成是Go的 "意见"，但其实还有更深层次的东西在里面。同意共同约定的局限性是一种方法，可以使一大类工具相互操作，因为它们都使用相同的基础语言。这是一种非常有效的方式，可以做得更少，但却可以实现更多。具体来说，在许多情况下，我们可以做最低限度的工作来建立对某个特定概念的共同理解，比如远程导入，或者源文件的正确格式，从而使创建的包和工具能够一起工作，因为他们都同意这些核心细节。

I’m going to return to that idea later.

我将在后面回到这个想法。

## Why is Go open source? 为什么Go是开源的？

But first, as I said earlier, I want to explain how I see the balance of Do Less and Enable More guiding our work on the broader Go open source project. To do that, I need to start with why Go is open source at all.

但首先，正如我之前所说的，我想解释一下我是如何看待 "少做多得 "的平衡点来指导我们在更广泛的Go开源项目上的工作的。要做到这一点，我需要从Go为什么要开放源代码开始。

Google pays me and others to work on Go, because, if Google’s programmers are more productive, Google can build products faster, maintain them more easily, and so on. But why open source Go? Why should Google share this benefit with the world?

谷歌付钱给我和其他人，让他们在Go上工作，因为如果谷歌的程序员有更高的生产力，谷歌就能更快地建立产品，更容易地维护它们，等等。但为什么要开放Go的源代码？为什么谷歌要与世界分享这种好处？

Of course, many of us worked on open source projects before Go, and we naturally wanted Go to be part of that open source world. But our preferences are not a business justification. The business justification is that Go is open source because that’s the only way that Go can succeed. We, the team that built Go within Google, knew this from day one. We knew that Go had to be made available to as many people as possible for it to succeed.

当然，我们中的许多人在Go之前就从事开源项目，我们自然希望Go能成为这个开源世界的一部分。但我们的偏好并不是一个商业理由。商业理由是，Go是开源的，因为这是Go成功的唯一途径。我们，在谷歌内部建立Go的团队，从第一天起就知道这一点。我们知道Go必须让尽可能多的人使用，才能成功。

Closed languages die.

封闭的语言会死亡。

A language needs large, broad communities.

一种语言需要大型、广泛的社区。

A language needs lots of people writing lots of software, so that when you need a particular tool or library, there’s a good chance it has already been written, by someone who knows the topic better than you, and who spent more time than you have to make it great.

一种语言需要大量的人编写大量的软件，这样当您需要一个特定的工具或库时，很有可能它已经被写出来了，而且是由比您更了解这个主题的人写的，他花了比您更多的时间来使它变得伟大。

A language needs lots of people reporting bugs, so that problems are identified and fixed quickly. Because of the much larger user base, the Go compilers are much more robust and spec-compliant than the Plan 9 C compilers they’re loosely based on ever were.

一门语言需要很多人报告错误，这样问题才能被发现并迅速修复。因为有了更大的用户群，Go编译器比它们松散地基于的Plan 9 C编译器更加强大和符合规范。

A language needs lots of people using it for lots of different purposes, so that the language doesn’t overfit to one use case and end up useless when the technology landscape changes.

一门语言需要有很多人将其用于不同的目的，这样语言就不会过度适应一个使用案例，而在技术环境发生变化时最终失去作用。

A language needs lots of people who want to learn it, so that there is a market for people to write books or teach courses, or run conferences like this one.

一门语言需要有很多人想学习它，这样就会有一个市场，让人们去写书、教课程，或者举办像这样的会议。

None of this could have happened if Go had stayed within Google. Go would have suffocated inside Google, or inside any single company or closed environment.

如果Go留在谷歌内部，这一切都不可能发生。Go会在谷歌内部，或在任何单一的公司或封闭的环境中窒息而死。

Fundamentally, Go must be open, and Go needs you. Go can’t succeed without all of you, without all the people using Go for all different kinds of projects all over the world.

从根本上说，Go必须是开放的，Go需要您。没有您们，没有全世界所有在不同项目中使用Go的人，Go不可能成功。

In turn, the Go team at Google could never be large enough to support the entire Go community. To keep scaling, we need to enable all this ``more'' while doing less. Open source is a huge part of that.

反过来，谷歌的Go团队也不可能大到足以支持整个Go社区。为了继续扩大规模，我们需要在减少工作的同时实现所有这些 "更多"。开源是其中的一个重要部分。

## Go’s open source Go的开放源码

What does open source mean? The minimum requirement is to open the source code, making it available under an open source license, and we’ve done that.

开源是什么意思？最起码的要求是开放源代码，在开放源码许可下提供，我们已经做到了。

But we also opened our development process: since announcing Go, we’ve done all our development in public, on public mailing lists open to all. We accept and review source code contributions from anyone. The process is the same whether you work for Google or not. We maintain our bug tracker in public, we discuss and develop proposals for changes in public, and we work toward releases in public. The public source tree is the authoritative copy. Changes happen there first. They are only brought into Google’s internal source tree later. For Go, being open source means that this is a collective effort that extends beyond Google, open to all.

但是我们也开放了我们的开发过程：自从发布Go以来，我们在公开的邮件列表中进行了所有的开发工作，对所有人开放。我们接受并审查来自任何人的源代码贡献。无论您是否为谷歌工作，这个过程都是一样的。我们公开维护我们的bug跟踪器，我们公开讨论和开发修改建议，我们公开为发布而努力。公共源码树是权威的副本。更改首先发生在那里。它们只是后来才被带入谷歌的内部源代码树。对Go来说，开放源代码意味着这是一个超越Google的集体努力，对所有人开放。

Any open source project starts with a few people, often just one, but with Go it was three: Robert Griesemer, Rob Pike, and Ken Thompson. They had a vision of what they wanted Go to be, what they thought Go could do better than existing languages, and Robert will talk more about that tomorrow morning. I was the next person to join the team, and then Ian Taylor, and then, one by one, we’ve ended up where we are today, with hundreds of contributors.

任何开源项目都是从几个人开始的，通常只有一个人，但Go是三个人。Robert Griesemer, Rob Pike, 和Ken Thompson。他们有一个愿景，那就是他们希望Go成为什么样的语言，他们认为Go可以比现有的语言做得更好，明天早上Robert会详细介绍这个问题。我是下一个加入团队的人，然后是Ian Taylor，然后，一个接一个，我们最终取得了今天的成就，有数百名贡献者。

Thank You to the many people who have contributed code or ideas or bug reports to the Go project so far. We tried to list everyone we could in our space in the program today. If your name is not there, I apologize, but thank you.

感谢到目前为止为Go项目贡献代码或想法或错误报告的许多人。我们试图在今天的节目中列出所有我们可以列出的人。如果您的名字不在那里，我很抱歉，但是谢谢您。

I believe the hundreds of contributors so far are working toward a shared vision of what Go can be. It’s hard to put words to these things, but I did my best to explain one part of the vision earlier: Do Less, Enable More.

我相信到目前为止，数以百计的贡献者正在努力实现Go的共同愿景。很难用语言来表达这些东西，但我已经尽力解释了这一愿景的一个部分。少做一点，多做一点。

## Google’s role 谷歌的作用

A natural question is: What is the role of the Go team at Google, compared to other contributors? I believe that role has changed over time, and it continues to change. The general trend is that over time the Go team at Google should be doing less and enabling more.

一个自然的问题是：与其他贡献者相比，Google的Go团队的角色是什么？我相信这个角色已经随着时间的推移发生了变化，而且还在继续变化。总的趋势是，随着时间的推移，Google的Go团队应该做得更少，促成更多。

In the very early days, before Go was known to the public, the Go team at Google was obviously working by itself. We wrote the first draft of everything: the specification, the compiler, the runtime, the standard library.

在很早的时候，在Go被公众所知之前，谷歌的Go团队显然是自己在工作。我们写了所有东西的初稿：规范、编译器、运行时、标准库。

Once Go was open sourced, though, our role began to change. The most important thing we needed to do was communicate our vision for Go. That’s difficult, and we’re still working at it. The initial implementation was an important way to communicate that vision, as was the development work we led that resulted in Go 1, and the various blog posts, and articles, and talks we’ve published.

不过，一旦Go被开源，我们的角色就开始改变了。我们需要做的最重要的事情是传达我们对Go的愿景。这很难，我们还在努力。最初的实施是传达这一愿景的重要方式，我们领导的开发工作也是如此，最终产生了Go 1，以及我们发表的各种博文、文章和讲座。

But as Rob said at Gophercon last year, "the language is done." Now we need to see how it works, to see how people use it, to see what people build. The focus now is on expanding the kind of work that Go can help with.

但正如Rob在去年的Gophercon上所说，"语言已经完成了"。现在我们需要看看它是如何工作的，看看人们是如何使用它的，看看人们建立了什么。现在的重点是扩大Go所能帮助的工作种类。

Google’s primarily role is now to enable the community, to coordinate, to make sure changes work well together, and to keep Go true to the original vision.

谷歌现在的主要作用是使社区能够发挥作用，进行协调，确保各种变化能够很好地结合起来，并使Go忠实于最初的愿景。

Google’s primary role is: Do Less. Enable More.

谷歌的主要作用是：少做事。启用更多。

I mentioned earlier that we’d rather have a small number of features that enable, say, 90% of the target use cases, and avoid the orders of magnitude more features necessary to reach 99 or 100%. We’ve been successful in applying that strategy to the areas of software that we know well. But if Go is to become useful in many new domains, we need experts in those areas to bring their expertise to our discussions, so that together we can design small adjustments that enable many new applications for Go.

我在前面提到，我们宁愿拥有少量的功能来实现，比如，90%的目标用例，而避免达到99%或100%所需的数量级的功能。在我们熟悉的软件领域中，我们已经成功地应用了这一策略。但是，如果Go要在许多新的领域变得有用，我们需要这些领域的专家把他们的专业知识带到我们的讨论中来，这样我们就可以一起设计一些小的调整，使Go有许多新的应用。

This shift applies not just to design but also to development. The role of the Go team at Google continues to shift more to one of guidance and less of pure development. I certainly spend much more time doing code reviews than writing code, more time processing bug reports than filing bug reports myself. We need to do less and enable more.

这种转变不仅适用于设计，也适用于开发。Google的Go团队的角色继续向指导性的转变，而不是纯粹的开发。我当然会花更多的时间做代码审查而不是写代码，花更多的时间处理错误报告而不是自己提交错误报告。我们需要做得更少，更有能力。

As design and development shift to the broader Go community, one of the most important things we the original authors of Go can offer is consistency of vision, to help keep Go Go. The balance that we must strike is certainly subjective. For example, a mechanism for extensible syntax would be a way to enable more ways to write Go code, but that would run counter to our goal of having a consistent language without different dialects.

随着设计和开发工作向更广泛的Go社区转移，我们这些Go的原作者能够提供的最重要的东西之一就是愿景的一致性，以帮助保持Go的发展。我们必须取得的平衡当然是主观的。例如，可扩展语法的机制可以使更多的方式来编写Go代码，但这与我们的目标相悖，即拥有一个没有不同方言的一致的语言。

We have to say no sometimes, perhaps more than in other language communities, but when we do, we aim to do so constructively and respectfully, to take that as an opportunity to clarify the vision for Go.

我们有时不得不说 "不"，也许比其他语言社区更多，但当我们这样做时，我们的目标是以建设性和尊重的方式进行，并将此作为澄清Go愿景的一个机会。

Of course, it’s not all coordination and vision. Google still funds Go development work. Rick Hudson is going to talk later today about his work on reducing garbage collector latency, and Hana Kim is going to talk tomorrow about her work on bringing Go to mobile devices. But I want to make clear that, as much as possible, we aim to treat development funded by Google as equal to development funded by other companies or contributed by individuals using their spare time. We do this because we don’t know where the next great idea will come from. Everyone contributing to Go should have the opportunity to be heard.

当然，这并不全是协调和愿景。谷歌仍然资助Go的开发工作。Rick Hudson今天晚些时候将谈论他在减少垃圾收集器延迟方面的工作，Hana Kim明天将谈论她在将Go引入移动设备方面的工作。但我想说明的是，在可能的情况下，我们的目标是将谷歌资助的开发与其他公司资助的或个人利用业余时间贡献的开发同等对待。我们这样做是因为我们不知道下一个伟大的想法会从哪里来。每个为Go做出贡献的人都应该有机会被听到。

### Examples 例子

I want to share some evidence for this claim that, over time, the original Go team at Google is focusing more on coordination than direct development.

我想为这一说法分享一些证据，随着时间的推移，谷歌的原始Go团队更注重于协调而不是直接开发。

First, the sources of funding for Go development are expanding. Before the open source release, obviously Google paid for all Go development. After the open source release, many individuals started contributing their time, and we’ve slowly but steadily been growing the number of contributors supported by other companies to work on Go at least part-time, especially as it relates to making Go more useful for those companies. Today, that list includes Canonical, Dropbox, Intel, Oracle, and others. And of course Gophercon and the other regional Go conferences are organized entirely by people outside Google, and they have many corporate sponsors besides Google.

首先，Go开发的资金来源正在扩大。在开源发布之前，显然谷歌支付了所有Go的开发费用。在开源发布之后，许多人开始贡献他们的时间，而且我们已经缓慢但稳定地增加了由其他公司支持的贡献者的数量，他们至少是兼职从事Go的工作，尤其是涉及到使Go对这些公司更有用。今天，这个名单包括Canonical、Dropbox、英特尔、甲骨文和其他公司。当然，Gophercon和其他地区的Go会议也完全是由谷歌以外的人组织的，除了谷歌以外，他们还有很多企业赞助商。

Second, the conceptual depth of Go development done outside the original team is expanding.

第二，在原始团队之外完成的Go开发的概念深度正在扩大。

Immediately after the open source release, one of the first large contributions was the port to Microsoft Windows, started by Hector Chu and completed by Alex Brainman and others. More contributors ported Go to other operating systems. Even more contributors rewrote most of our numeric code to be faster or more precise or both. These were all important contributions, and very much appreciated, but for the most part they did not involve new designs.

在开源发布之后，最早的大型贡献之一是对微软Windows的移植，由Hector Chu开始，由Alex Brainman和其他人完成。更多的贡献者将Go移植到其他操作系统上。更多的贡献者重写了我们大部分的数字代码，使之更快、更精确或两者兼而有之。这些都是重要的贡献，而且非常值得赞赏，但在大多数情况下，它们并不涉及新的设计。

More recently, a group of contributors led by Aram Hăvărneanu ported Go to the ARM 64 architecture, This was the first architecture port by contributors outside Google. This is significant, because in general support for a new architecture requires more design work than support for a new operating system. There is more variation between architectures than between operating systems.

最近，由Aram Hăvărneanu领导的一组贡献者将Go移植到了ARM 64架构上，这是第一个由Google以外的贡献者移植的架构。这很重要，因为一般来说，支持一个新的架构比支持一个新的操作系统需要更多的设计工作。架构之间的差异比操作系统之间的差异更大。

Another example is the introduction over the past few releases of preliminary support for building Go programs using shared libraries. This feature is important for many Linux distributions but not as important for Google, because we deploy static binaries. We have been helping guide the overall strategy, but most of the design and nearly all of the implementation has been done by contributors outside Google, especially Michael Hudson-Doyle.

另一个例子是在过去几个版本中引入了对使用共享库构建Go程序的初步支持。这个功能对许多Linux发行版来说是很重要的，但对Google来说却不那么重要，因为我们部署的是静态二进制文件。我们一直在帮助指导整体策略，但大部分的设计和几乎所有的实现都是由Google以外的贡献者完成的，特别是Michael Hudson-Doyle。

My last example is the go command’s approach to vendoring. I define vendoring as copying source code for external dependencies into your tree to make sure that they don’t disappear or change underfoot.

我的最后一个例子是go命令的销售方法。我把vendoring定义为把外部依赖的源代码复制到您的树上，以确保它们不会在脚下消失或改变。

Vendoring is not a problem Google suffers, at least not the way the rest of the world does. We copy open source libraries we want to use into our shared source tree, record what version we copied, and only update the copy when there is a need to do so. We have a rule that there can only be one version of a particular library in the source tree, and it’s the job of whoever wants to upgrade that library to make sure it keeps working as expected by the Google code that depends on it. None of this happens often. This is the lazy approach to vendoring.

摊销并不是谷歌所面临的问题，至少不是像世界其他地方那样。我们把我们想要使用的开源库复制到我们的共享源码树中，记录我们复制的版本，并且只在有需要的时候更新副本。我们有一个规则，在源码树中只能有一个特定的库的版本，不管是谁想升级这个库，都要确保它能按照谷歌代码的预期运行。这一切都不经常发生。这就是懒惰的销售方法。

In contrast, most projects outside Google take a more eager approach, importing and updating code using automated tools and making sure that they are always using the latest versions.

相比之下，谷歌以外的大多数项目采取了一种更急切的方法，使用自动化工具导入和更新代码，并确保他们总是在使用最新的版本。

Because Google has relatively little experience with this vendoring problem, we left it to users outside Google to develop solutions. Over the past five years, people have built a series of tools. The main ones in use today are Keith Rarick’s godep, Owen Ou’s nut, and the gb-vendor plugin for Dave Cheney’s gb,

由于谷歌在这个销售问题上的经验相对较少，我们把它留给了谷歌以外的用户来开发解决方案。在过去的五年里，人们建立了一系列的工具。现在使用的主要是Keith Rarick的godep，Owen Ou的nut，以及Dave Cheney的gb-vendor插件。

There are two problems with the current situation. The first is that these tools are not compatible out of the box with the go command’s "go get". The second is that the tools are not even compatible with each other. Both of these problems fragment the developer community by tool.

目前的情况有两个问题。第一个问题是，这些工具与go命令的 "go get "不兼容。第二是这些工具甚至不能相互兼容。这两个问题都使开发者社区被工具分割开来。

Last fall, we started a public design discussion to try to build consensus on some basics about how these tools all operate, so that they can work alongside "go get" and each other.

去年秋天，我们开始了一个公开的设计讨论，试图在这些工具如何运作的一些基本问题上建立共识，以便它们能够与 "go get "和彼此一起工作。

Our basic proposal was that all tools agree on the approach of rewriting import paths during vendoring, to fit with "go get"’s model, and also that all tools agree on a file format describing the source and version of the copied code, so that the different vendoring tools can be used together even by a single project. If you use one today, you should still be able to use another tomorrow.

我们的基本建议是，所有的工具都同意在销售过程中重写导入路径的方法，以适应 "go get "的模式，并且所有的工具都同意描述复制的代码的来源和版本的文件格式，以便不同的销售工具可以一起使用，即使是在一个项目中。如果您今天使用一个，明天应该仍然能够使用另一个。

Finding common ground in this way was very much in the spirit of Do Less, Enable More. If we could build consensus about these basic semantic aspects, that would enable "go get" and all these tools to interoperate, and it would enable switching between tools, the same way that agreement about how Go programs are stored in text files enables the Go compiler and all text editors to interoperate. So we sent out our proposal for common ground.

以这种方式找到共同点，非常符合 "少做多得 "的精神。如果我们能够就这些基本的语义方面达成共识，这将使 "go get "和所有这些工具能够互操作，并且能够在不同的工具之间进行切换，就像就Go程序如何存储在文本文件中达成的协议使Go编译器和所有文本编辑器能够互操作一样。所以我们发出了我们的建议，以寻求共同点。

Two things happened.

发生了两件事。

First, Daniel Theophanes started a vendor-spec project on GitHub with a new proposal and took over coordination and design of the spec for vendoring metadata.

首先，Daniel Theophanes在GitHub上启动了一个供应商规范项目，并接管了供应商元数据规范的协调和设计。

Second, the community spoke with essentially one voice to say that rewriting import paths during vendoring was not tenable. Vendoring works much more smoothly if code can be copied without changes.

第二，社区基本上用一个声音说，在销售过程中重写导入路径是不可行的。如果代码可以不经修改就被复制，那么销售工作就会顺利得多。

Keith Rarick posted an alternate proposal for a minimal change to the go command to support vendoring without rewriting import paths. Keith’s proposal was configuration-free and fit in well with the rest of the go command’s approach. That proposal will ship as an experimental feature in Go 1.5 and likely enabled by default in Go 1.6. And I believe that the various vendoring tool authors have agreed to adopt Daniel’s spec once it is finalized.

Keith Rarick发表了另一个建议，即对go命令进行最小的修改，以支持销售而不重写导入路径。Keith的建议是无配置的，并且与go命令的其他方法很相配。该建议将作为实验性功能在Go 1.5中推出，并可能在Go 1.6中默认启用。而且我相信，一旦Daniel的规范最终确定下来，各个销售工具的作者都同意采用该规范。

The result is that at the next Gophercon we should have broad interoperability between vendoring tools and the go command, and the design to make that happen was done entirely by contributors outside the original Go team.

其结果是，在下一次Gophercon会议上，我们应该在销售工具和go命令之间实现广泛的互操作性，而实现这一目标的设计完全是由Go团队以外的贡献者完成。

Not only that, the Go team’s proposal for how to do this was essentially completely wrong. The Go community told us that very clearly. We took that advice, and now there’s a plan for vendoring support that I believe everyone involved is happy with.

不仅如此，Go团队关于如何做到这一点的建议基本上是完全错误的。Go社区非常清楚地告诉我们这一点。我们采纳了这个建议，现在有了一个我相信每个人都会满意的销售支持计划。

This is also a good example of our general approach to design. We try not to make any changes to Go until we feel there is broad consensus on a well-understood solution. For vendoring, feedback and design from the Go community was critical to reaching that point.

这也是我们一般设计方法的一个好例子。我们尽量不对Go做任何改动，直到我们觉得对一个被充分理解的解决方案有广泛的共识。对于vendoring来说，来自Go社区的反馈和设计是达到这一点的关键。

This general trend toward both code and design coming from the broader Go community is important for Go. You, the broader Go community, know what is working and what is not in the environments where you use Go. We at Google don’t. More and more, we will rely on your expertise, and we will try to help you develop designs and code that extend Go to be useful in more settings and fit well with Go’s original vision. At the same time, we will continue to wait for broad consensus on well-understood solutions.

这种代码和设计都来自于更广泛的Go社区的大趋势对Go来说非常重要。您们，更广泛的 Go 社区，知道在您们使用 Go 的环境中哪些是有效的，哪些是无效的。而我们 Google 则不知道。我们将越来越多地依靠您们的专业知识，我们将努力帮助您们开发设计和代码，使 Go 在更多的环境中发挥作用，并与 Go 的原始愿景相一致。同时，我们将继续等待人们对公认的解决方案达成广泛的共识。

This brings me to my last point.

这就引出了我的最后一点。

## Code of Conduct 行为准则

I’ve argued that Go must be open, and that Go needs your help.

我已经论证了Go必须是开放的，Go需要您的帮助。

But in fact Go needs everyone’s help. And everyone isn’t here.

但事实上，Go需要所有人的帮助。而每个人都不在这里。

Go needs ideas from as many people as possible.

Go需要来自尽可能多的人的想法。

To make that a reality, the Go community needs to be as inclusive, welcoming, helpful, and respectful as possible.

为了实现这一目标，Go社区需要尽可能的包容、欢迎、帮助和尊重。

The Go community is large enough now that, instead of assuming that everyone involved knows what is expected, I and others believe that it makes sense to write down those expectations explicitly. Much like the Go spec sets expectations for all Go compilers, we can write a spec setting expectations for our behavior in online discussions and in offline meetings like this one.

Go社区现在已经足够大了，我和其他人认为，与其假设每个人都知道期望是什么，不如明确地写下这些期望。就像Go规范为所有Go编译器设定的期望一样，我们也可以写一个规范，为我们在网上讨论和像这样的线下会议中的行为设定期望。

Like any good spec, it must be general enough to allow many implementations but specific enough that it can identify important problems. When our behavior doesn’t meet the spec, people can point that out to us, and we can fix the problem. At the same time, it’s important to understand that this kind of spec cannot be as precise as a language spec. We must start with the assumption that we will all be reasonable in applying it.

就像任何好的规范一样，它必须有足够的通用性，以允许许多实现，但也要有足够的特殊性，以便它能识别重要的问题。当我们的行为不符合规范时，人们可以向我们指出这一点，我们就可以解决这个问题。同时，重要的是要明白，这种规范不可能像语言规范那样精确。我们必须以我们在应用它时都会合理的假设为出发点。

This kind of spec is often referred to as a Code of Conduct. Gophercon has one, which we’ve all agreed to follow by being here, but the Go community does not. I and others believe the Go community needs a Code of Conduct.

这种规范通常被称为行为准则。Gophercon有一个，我们在这里都同意遵守，但Go社区却没有。我和其他人认为Go界需要一个行为准则。

But what should it say?

但它应该怎么说呢？

I believe the most important overall statement we can make is that if you want to use or discuss Go, then you are welcome here, in our community. That is the standard I believe we aspire to.

我相信我们可以做出的最重要的总体声明是：如果您想使用或讨论Go，那么我们就欢迎您来这里，来我们的社区。这就是我认为我们所追求的标准。

If for no other reason (and, to be clear, there are excellent other reasons), Go needs as large a community as possible. To the extent that behavior limits the size of the community, it holds Go back. And behavior can easily limit the size of the community.

如果没有其他原因的话（说白了，有很好的其他原因），Go需要一个尽可能大的社区。如果行为限制了社区的规模，它就会阻碍Go的发展。而行为可以轻易地限制社区的规模。

The tech community in general and the Go community in particular is skewed toward people who communicate bluntly. I don’t believe this is fundamental. I don’t believe this is necessary. But it’s especially easy to do in online discussions like email and IRC, where plain text is not supplemented by the other cues and signals we have in face-to-face interactions.

一般的技术社区，尤其是Go社区，都是偏向于直截了当地交流的人。我不相信这是根本。我也不相信这是必要的。但在电子邮件和IRC这样的在线讨论中，特别容易出现这种情况，因为纯文本没有我们在面对面交流时的其他线索和信号作为补充。

For example, I have learned that when I am pressed for time I tend to write fewer words, with the end result that my emails seem not just hurried but blunt, impatient, even dismissive. That’s not how I feel, but it’s how I can come across, and that impression can be enough to make people think twice about using or contributing to Go. I realized I was doing this when some Go contributors sent me private email to let me know. Now, when I am pressed for time, I pay extra attention to what I’m writing, and I often write more than I naturally would, to make sure I’m sending the message I intend.

例如，我了解到，当我时间紧迫时，我倾向于写更少的字，最终的结果是我的电子邮件看起来不仅是匆忙的，而且是直率的、不耐烦的，甚至是轻蔑的。这不是我的感觉，但这是我的表现，而这种印象足以让人们对使用Go或为Go做贡献三思而后行。当一些Go贡献者给我发来私人邮件让我知道时，我意识到我正在这样做。现在，当我时间紧迫的时候，我会格外注意我所写的东西，而且我经常写得比我自然会写的更多，以确保我发出的是我想要的信息。

I believe that correcting the parts of our everyday interactions, intended or not, that drive away potential users and contributors is one of the most important things we can all do to make sure the Go community continues to grow. A good Code of Conduct can help us do that.

我相信，纠正我们日常互动中有意或无意地驱赶潜在用户和贡献者的部分，是我们所有人可以做的最重要的事情之一，以确保Go社区继续发展。一个好的行为准则可以帮助我们做到这一点。

We have no experience writing a Code of Conduct, so we have been reading existing ones, and we will probably adopt an existing one, perhaps with minor adjustments. The one I like the most is the Django Code of Conduct, which originated with another project called SpeakUp! It is structured as an elaboration of a list of reminders for everyday interaction.

我们没有编写行为准则的经验，所以我们一直在阅读现有的行为准则，我们可能会采用现有的行为准则，也许会稍加调整。我最喜欢的是Django行为准则，它起源于另一个叫SpeakUp的项目。它的结构是对日常互动的提醒清单的阐述。

"Be friendly and patient. Be welcoming. Be considerate. Be respectful. Be careful in the words that you choose. When we disagree, try to understand why."

"要友好和耐心。要热情好客。要考虑周到。要尊重他人。要注意您所选择的词语。当我们有不同意见时，试着理解原因。

I believe this captures the tone we want to set, the message we want to send, the environment we want to create for new contributors. I certainly want to be friendly, patient, welcoming, considerate, and respectful. I won’t get it exactly right all the time, and I would welcome a helpful note if I’m not living up to that. I believe most of us feel the same way.

我相信这抓住了我们想要设定的基调，我们想要传递的信息，我们想要为新的贡献者创造的环境。我当然希望能做到友好、耐心、热情、周到和尊重。我不会每次都做得很好，如果我没有达到这个要求，我欢迎大家给我提供帮助。我相信我们大多数人都有同样的感觉。

I haven’t mentioned active exclusion based on or disproportionately affecting race, gender, disability, or other personal characteristics, and I haven’t mentioned harassment. For me, it follows from what I just said that exclusionary behavior or explicit harassment is absolutely unacceptable, online and offline. Every Code of Conduct says this explicitly, and I expect that ours will too. But I believe the SpeakUp! reminders about everyday interactions are an equally important statement. I believe that setting a high standard for those everyday interactions makes extreme behavior that much clearer and easier to deal with.

我没有提到基于种族、性别、残疾或其他个人特征的主动排斥或不成比例的影响，我也没有提到骚扰。对我来说，从我刚才所说的内容来看，排他性行为或明确的骚扰是绝对不能接受的，无论是在线还是离线。每一个行为准则都明确规定了这一点，我希望我们的行为准则也会如此。但我相信，SpeakUp！对日常互动的提醒是一个同样重要的声明。我相信，为这些日常互动设定一个高标准，会使极端行为变得更加清晰和容易处理。

I have no doubts that the Go community can be one of the most friendly, welcoming, considerate, and respectful communities in the tech industry. We can make that happen, and it will be a benefit and credit to us all.

我毫不怀疑，Go社区可以成为科技行业中最友好、最受欢迎、最体贴和最尊重的社区之一。我们可以实现这一目标，这将是我们所有人的利益和功劳。

Andrew Gerrand has been leading the effort to adopt an appropriate Code of Conduct for the Go community. If you have suggestions, or concerns, or experience with Codes of Conduct, or want to be involved, please find Andrew or me during the conference. If you’ll still be here on Friday, Andrew and I are going to block off some time for Code of Conduct discussions during Hack Day.

Andrew Gerrand一直在领导为Go社区制定适当的行为准则的工作。如果您有建议，或担忧，或对行为准则有经验，或想参与，请在会议期间找到安德鲁或我。如果您周五还在这里，安德鲁和我将在黑客日期间为行为准则的讨论留出一些时间。

Again, we don’t know where the next great idea will come from. We need all the help we can get. We need a large, diverse Go community.

同样，我们不知道下一个伟大的想法会从哪里来。我们需要所有我们能得到的帮助。我们需要一个庞大的、多样化的Go社区。

## Thank You 感谢您

I consider the many people releasing software for download using "go get," sharing their insights via blog posts, or helping others on the mailing lists or IRC to be part of this broad open source effort, part of the Go community. Everyone here today is also part of that community.

我认为许多人使用 "go get "发布软件供下载，通过博客文章分享他们的见解，或在邮件列表或IRC上帮助他人，都是这个广泛的开源努力的一部分，是Go社区的一部分。今天在座的各位也是这个社区的一部分。

Thank you in advance to the presenters who over the next few days will take time to share their experiences using and extending Go.

提前感谢各位演讲者，他们将在接下来的几天里抽出时间来分享他们使用和扩展Go的经验。

Thank you in advance to all of you in the audience for taking the time to be here, to ask questions, and to let us know how Go is working for you. When you go back home, please continue to share what you’ve learned. Even if you don’t use Go for daily work, we’d love to see what’s working for Go adopted in other contexts, just as we’re always looking for good ideas to bring back into Go.

提前感谢在座的各位，感谢您们花时间来到这里，提出问题，并让我们知道Go对您们的作用。当您们回家的时候，请继续分享您们所学到的东西。即使您不在日常工作中使用Go，我们也希望看到Go在其他环境中的应用，就像我们一直在寻找好的想法带回Go中一样。

Thank you all again for making the effort to be here and for being part of the Go community.

再次感谢大家努力来到这里，成为Go社区的一员。

For the next few days, please: tell us what we’re doing right, tell us what we’re doing wrong, and help us all work together to make Go even better.

在接下来的几天里，请告诉我们我们做对了什么，告诉我们做错了什么，并帮助我们一起努力使Go变得更好。

Remember to be friendly, patient, welcoming, considerate, and respectful.

记住要友好、耐心、热情、周到和尊重。

Above all, enjoy the conference.

最重要的是，享受这个会议。
