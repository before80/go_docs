+++
title = "与 go 团队的对话"
weight = 11
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# A conversation with the Go team - 与 go 团队的对话

https://go.dev/blog/io2013-chat

6 June 2013

At Google I/O 2013, several members of the Go team hosted a "Fireside chat." Robert Griesemer, Rob Pike, David Symonds, Andrew Gerrand, Ian Lance Taylor, Sameer Ajmani, Brad Fitzpatrick, and Nigel Tao took questions from the audience and people around the world about various aspects of the Go project.

在2013年谷歌I/O大会上，Go团队的几位成员主持了一场 "炉边谈话"。Robert Griesemer、Rob Pike、David Symonds、Andrew Gerrand、Ian Lance Taylor、Sameer Ajmani、Brad Fitzpatrick和Nigel Tao接受了来自观众和世界各地的人们关于Go项目各个方面的问题。

<iframe src="https://www.youtube.com/embed/p9VUCp98ay4" width="560" height="315" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

We also hosted a similar session at I/O last year: [*Meet the Go team*](http://www.youtube.com/watch?v=sln-gJaURzk).

我们在去年的I/O大会上也举办了类似的会议。认识Go团队。

There were many more questions from Google Moderator than we were able to answer in the short 40 minute session. Here we answer some of those we missed in the live session.

在短短40分钟的会议中，来自谷歌版主的问题比我们能够回答的多得多。在这里，我们回答了一些我们在现场会议上错过的问题。

*Linking speed (and memory usage) for the gc toolchain are a known problem.* *Are there any plans to address this during the 1.2 cycle?*

gc工具链的链接速度（和内存使用）是一个已知的问题。是否有计划在1.2周期内解决这个问题？

**Rob:** Yes. We are always thinking about ways to improve performance of the tools as well as the language and libraries.

Rob：是的。我们一直在考虑如何提高工具以及语言和库的性能。

*I have been very pleased to see how quickly Go appears to be gaining traction.* *Can you talk about the reactions you have experienced working with other* *developers inside and outside Google? Are there any major sticking points remaining?*

我非常高兴地看到，Go似乎正在迅速地获得吸引力。您能谈谈您在与谷歌内部和外部的其他开发者合作中所经历的反应吗？是否还有什么主要的症结所在？

**Robert:** A lot of developers that seriously tried Go are very happy with it. Many of them report a much smaller, more readable and thus maintainable code base: A 50% code size reduction or more when coming from C++ seems common. Developers that switched to Go from Python are invariably pleased with the performance gain. The typical complaints are about small inconsistencies in the language (some of which we might iron out at some point). What surprises me is that almost nobody complains about the lack of generics.

Robert：很多认真尝试过Go的开发者都对它非常满意。他们中的许多人报告说，Go的代码库更小、更易读，因而更容易维护。从C++转到Go，代码量减少了50%或更多，这似乎很常见。从Python转到Go的开发者无一例外地对其性能的提高感到满意。典型的抱怨是关于语言中的一些小的不一致（其中一些我们可能会在某个时间点上解决）。让我感到惊讶的是，几乎没有人抱怨缺乏泛型的问题。

*When will Go be a first-class language for Android development?*

Go什么时候能成为Android开发的一流语言？

**Andrew:** This would be great, but we don’t have anything to announce.

安德鲁：这将是一件好事，但我们没有任何东西可以宣布。

*Is there a roadmap for the next version of Go?*

Go的下一个版本有什么路线图吗？

**Andrew:** We have no feature roadmap as such. The contributors tend to work on what interests them. Active areas of development include the gc and gccgo compilers, the garbage collector and runtime, and many others. We expect the majority of exciting new additions will be in the form of improvements to our tools. You can find design discussions and code reviews on the [golang-dev mailing list](http://groups.google.com/group/golang-dev).

安德鲁：我们没有这样的功能路线图。贡献者们倾向于在他们感兴趣的地方工作。活跃的开发领域包括gc和gccgo编译器，垃圾收集器和运行时间，以及其他许多领域。我们希望大多数令人振奋的新功能将以改进我们的工具的形式出现。您可以在golang-dev邮件列表中找到设计讨论和代码审查。

As for the timeline, we do have [concrete plans](https://docs.google.com/document/d/106hMEZj58L9nq9N9p7Zll_WKfo-oyZHFyI6MttuZmBU/edit?usp=sharing): we expect to release Go 1.2 on December 1, 2013.

至于时间表，我们确实有具体的计划：我们预计在2013年12月1日发布Go 1.2。

*Where do you guys want to see Go used externally?* *What would you consider a big win for Go adoption outside Google?* *Where do you think Go has the potential to make a significant impact?*

您们想在哪里看到Go的外部应用？您认为在谷歌之外，Go的采用会有什么大的胜利？您认为Go在哪里有可能产生重大影响？

**Rob:** Where Go is deployed is up to its users, not to us. We’re happy to see it gain traction anywhere it helps. It was designed with server-side software in mind, and is showing promise there, but has also shown strengths in many other areas and the story is really just beginning. There are many surprises to come.

Rob：Go的部署由其用户决定，而不是由我们决定。我们很高兴看到它在任何有帮助的地方获得牵引力。它在设计时考虑到了服务器端软件，并在那里显示了前景，但也在许多其他领域显示了优势，故事才刚刚开始。还有很多惊喜要发生。

**Ian:** It’s easier for startups to use Go, because they don’t have an entrenched code base that they need to work with. So I see two future big wins for Go. One would be a significant use of Go by an existing large software company other than Google. Another would be a significant IPO or acquisition of a startup that primarily uses Go. These are both indirect: clearly choice of programming language is a very small factor in the success of a company. But it would be another way to show that Go can be part of a successful software system.

伊恩：初创公司使用Go更容易，因为他们没有一个根深蒂固的代码库需要使用。所以我认为Go有两个未来的大赢家。一个是谷歌以外的现有大型软件公司对Go的大量使用。另一个是一个主要使用Go的创业公司的重大IPO或收购。这两者都是间接的：显然，编程语言的选择对一个公司的成功是一个非常小的因素。但这将是显示Go可以成为成功软件系统的一部分的另一种方式。

*Have you thought any (more) about the potential of dynamically loading* *Go packages or objects and how it could work in Go?* *I think this could enable some really interesting and expressive constructs,* *especially coupled with interfaces.*

您有没有想过（更多）动态加载Go包或对象的潜力，以及它如何在Go中工作？我认为这可以实现一些非常有趣和有表现力的结构，特别是与接口相结合。

**Rob:** This is an active topic of discussion. We appreciate how powerful the concept can be and hope we can find a way to implement it before too long. There are serious challenges in the design approach to take and the need to make it work portably.

Rob: 这是个活跃的讨论话题。我们很欣赏这个概念的强大，并希望我们能在不久之后找到一种方法来实现它。在设计方法上存在严重的挑战，而且需要使其可移植地工作。

*There was a discussion a while ago about collecting some best-of-breed* `database/sql` *drivers in a more central place.* *Some people had strong opinions to the contrary though.* *Where is* `database/sql` *and its drivers going in the next year?*

不久前，有一个关于将一些最好的数据库/sql驱动收集在一个更集中的地方的讨论。但有些人有强烈的反对意见。明年，数据库/sql及其驱动将走向何方？

**Brad:** While we could create an official subrepo ("go.db") for database drivers, we fear that would unduly bless certain drivers. At this point we’d still rather see healthy competition between different drivers. The [SQLDrivers wiki page](https://go.dev/wiki/SQLDrivers) lists some good ones.

Brad：虽然我们可以为数据库驱动创建一个官方子程序（"go.db"），但我们担心这将不适当地保护某些驱动。在这一点上，我们仍然希望看到不同驱动之间的良性竞争。SQLDrivers维基页面列出了一些好的驱动。

The `database/sql` package didn’t get much attention for a while, due to lack of drivers. Now that drivers exist, usage of the package is increasing and correctness and performance bugs are now being reported (and fixed). Fixes will continue, but no major changes to the interface of `database/sql` are planned.  There might be small extensions here and there as needed for performance or to assist some drivers.

由于缺乏驱动，数据库/sql包有一段时间没有得到很多关注。现在有了驱动，包的使用量正在增加，正确性和性能方面的错误现在也被报告（和修复）。修复工作将继续进行，但没有计划对数据库/sql的界面进行重大改变。 在这里和那里可能会有一些小的扩展，以满足性能或协助一些驱动程序的需要。

*What is the status of versioning?* *Is importing some code from GitHub a best practice recommended by the Go team?* *What happens when we publish our code that is dependent on a GitHub repo and* *the API of the dependee changes?*

版本管理的情况如何？从GitHub导入一些代码是Go团队推荐的最佳做法吗？当我们发布了依赖GitHub repo的代码，而被依赖者的API发生变化时，会发生什么？

**Ian:** This is frequently discussed on the mailing list. What we do internally is take a snapshot of the imported code, and update that snapshot from time to time. That way, our code base won’t break unexpectedly if the API changes. But we understand that that approach doesn’t work very well for people who are themselves providing a library. We’re open to good suggestions in this area. Remember that this is an aspect of the tools that surround the language rather than the language itself; the place to fix this is in the tools, not the language.

Ian: 这个问题经常在邮件列表中讨论。我们在内部所做的是对导入的代码进行快照，并不时地更新该快照。这样一来，如果API发生变化，我们的代码库就不会意外中断。但我们明白，这种方法对于那些自己提供一个库的人来说并不是很有效。我们愿意接受这方面的好建议。请记住，这是围绕语言的工具的一个方面，而不是语言本身；解决这个问题的地方是在工具中，而不是在语言中。

*What about Go and Graphical User Interfaces?*

Go和图形用户界面怎么样？

**Rob:** This is a subject close to my heart. Newsqueak, a very early precursor language, was designed specifically for writing graphics programs (that’s what we used to call apps). The landscape has changed a lot but I think Go’s concurrency model has much to offer in the field of interactive graphics.

Rob：这是我最关心的一个问题。Newsqueak是一种非常早期的前驱语言，是专门为编写图形程序而设计的（这就是我们过去所说的应用程序）。现在的情况已经发生了很大的变化，但我认为Go的并发模型在交互式图形领域有很多值得借鉴的地方。

**Andrew:** There are many [bindings for existing graphics libraries](https://go.dev/wiki/Projects#Graphics_and_Audio) out there, and a few Go-specific projects. One of the more promising ones is [go.uik](https://github.com/skelterjohn/go.uik), but it’s still in its early days. I think there’s a lot of potential for a great Go-specific UI toolkit for writing native applications (consider handling user events by receiving from a channel), but developing a production-quality package is a significant undertaking. I have no doubt one will come in time.

Andrew：现在有很多现有图形库的绑定，还有一些专门针对Go的项目。其中一个比较有前途的项目是go.uik，但它仍然处于早期阶段。我认为为编写本地应用程序提供一个伟大的Go专用UI工具包有很大的潜力（考虑通过从通道接收来处理用户事件），但开发一个高质量的软件包是一项重要的工作。我毫不怀疑，随着时间的推移，会有这样的工具包出现。

In the meantime, the web is the most broadly available platform for user interfaces. Go provides great support for building web apps, albeit only on the back end.

同时，网络是最广泛使用的用户界面平台。Go为构建网络应用提供了巨大的支持，尽管只是在后端。

*In the mailing lists Adam Langley has stated that the TLS code has not been* *reviewed by outside groups, and thus should not be used in production.* *Are there plans to have the code reviewed?* *A good secure implementation of concurrent TLS would be very nice.*

在邮件列表中，Adam Langley表示，TLS代码没有经过外部团体的审查，因此不应该在生产中使用。是否有计划对代码进行审查？一个好的安全的并发TLS的实现将是非常好的。

**Adam**: Cryptography is notoriously easy to botch in subtle and surprising ways and I’m only human. I don’t feel that I can warrant that Go’s TLS code is flawless and I wouldn’t want to misrepresent it.

Adam: 密码学是出了名的容易出错，而且是以微妙和令人惊讶的方式出错，我也只是个普通人。我不觉得我可以保证Go的TLS代码是完美无缺的，我也不想歪曲它。

There are a couple of places where the code is known to have side-channel issues: the RSA code is blinded but not constant time, elliptic curves other than P-224 are not constant time and the Lucky13 attack might work. I hope to address the latter two in the Go 1.2 timeframe with a constant-time P-256 implementation and AES-GCM.

有几个地方的代码已知有侧信道问题：RSA代码是盲目的，但不是恒定时间的，P-224以外的椭圆曲线不是恒定时间的，Lucky13攻击可能有效。我希望在Go 1.2的时间框架内用恒定时间的P-256实现和AES-GCM来解决后两者。

Nobody has stepped forward to do a review of the TLS stack however and I’ve not investigated whether we could get Matasano or the like to do it. That depends on whether Google wishes to fund it.

然而，没有人站出来对TLS堆栈进行审查，我也没有调查过我们是否可以让Matasano或类似的人做这件事。这取决于谷歌是否愿意资助它。

*What do you think about* [*GopherCon 2014*](http://www.gophercon.com/)*?* *Does anyone from the team plan to attend?*

您对GopherCon 2014有什么看法？团队中有人打算参加吗？

**Andrew:** It’s very exciting. I’m sure some of us will be there.

安德鲁：这非常令人兴奋。我相信我们中的一些人将会去那里。
