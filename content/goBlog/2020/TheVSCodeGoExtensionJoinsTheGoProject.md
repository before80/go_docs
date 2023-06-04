+++
title = "VS Code Go扩展加入了go项目"
weight = 10
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# The VS Code Go extension joins the Go project - VS Code Go扩展加入了go项目

https://go.dev/blog/vscode-go

The Go team
9 June 2020

When the Go project began, "an overarching goal was that Go do more to help the working programmer by enabling tooling, automating mundane tasks such as code formatting, and removing obstacles to working on large code bases" ([Go FAQ](https://go.dev/doc/faq#What_is_the_purpose_of_the_project)). Today, more than a decade later, we continue to be guided by that same goal, especially as it pertains to the programmer’s most critical tool: their editor.

当Go项目开始时，"一个首要的目标是，Go通过启用工具、自动化代码格式化等平凡的任务以及消除在大型代码库中工作的障碍，来帮助工作中的程序员"（Go FAQ）。十多年后的今天，我们仍然以这一目标为指导，尤其是涉及到程序员最关键的工具：他们的编辑器。

Throughout the past decade, Go developers have relied on a variety of editors and dozens of independently authored tools and plugins. Much of Go’s early success can be attributed to the fantastic development tools created by the Go community. The [VS Code extension for Go](https://github.com/microsoft/vscode-go), built using many of these tools, is now used by 41 percent of Go developers ([Go developer survey](https://blog.golang.org/survey2019-results)).

在过去的十年中，Go 的开发者依赖各种编辑器和几十个独立编写的工具和插件。Go的早期成功在很大程度上可以归功于Go社区创造的神奇的开发工具。使用许多这些工具构建的VS Code Go扩展，现在有41%的Go开发者在使用（Go开发者调查）。

As the VS Code Go extension grows in popularity and as [the ecosystem expands](https://www.youtube.com/watch?v=EFJfdWzBHwE), it requires [more maintenance and support](https://twitter.com/ramyanexus/status/1154470078978486272). Over the past few years, the Go team has collaborated with the VS Code team to help the Go extension maintainers. The Go team also began a new initiative to improve the tools powering all Go editor extensions, with a focus on supporting the [Language Server Protocol](https://microsoft.github.io/language-server-protocol/) with [`gopls`](https://go.dev/s/gopls) and [the Debug Adapter Protocol with Delve](https://github.com/go-delve/delve/issues/1515).

随着VS Code Go扩展的普及和生态系统的扩大，它需要更多的维护和支持。在过去的几年里，Go团队与VS Code团队合作，帮助Go扩展的维护者。Go团队还开始了一项新的计划，以改善支持所有Go编辑器扩展的工具，重点是支持gopls的语言服务器协议和Delve的调试适配器协议。

Through this collaborative work between the VS Code and Go teams, we realized that the Go team is uniquely positioned to evolve the Go development experience alongside the Go language.

通过 VS Code 和 Go 团队之间的合作，我们意识到 Go 团队在与 Go 语言一起发展 Go 开发体验方面具有独特的优势。

As a result, we’re happy to announce the next phase in the Go team’s partnership with the VS Code team: **The VS Code extension for Go is officially joining the Go project**. With this come two critical changes:

因此，我们很高兴地宣布 Go 团队与 VS Code 团队合作的下一个阶段。VS Code 的 Go 扩展将正式加入 Go 项目。这带来了两个关键的变化：

1. The publisher of the plugin is shifting from "Microsoft" to "Go Team at Google".该插件的发布者从 "微软 "转为 "谷歌的Go团队"。
2. The project’s repository is moving to join the rest of the Go project at https://github.com/golang/vscode-go.该项目仓库将加入Go项目的其他部分，网址是https://github.com/golang/vscode-go。

We cannot overstate our gratitude to those who have helped build and maintain this beloved extension. We know that innovative ideas and features come from you, our users. The Go team’s primary aim as owners of the extension is to reduce the burden of maintenance work on the Go community. We’ll make sure the builds stay green, the issues get triaged, and the docs get updated. Go team members will keep contributors abreast of relevant language changes, and we’ll smooth the rough edges between the extension’s different dependencies.

我们对那些帮助建立和维护这个心爱的扩展的人的感激之情怎么说都不过分。我们知道，创新的想法和功能来自于您们，我们的用户。作为扩展的所有者，Go团队的首要目标是减少Go社区的维护工作负担。我们将确保构建工作保持绿色，问题得到分流，文档得到更新。Go团队成员会让贡献者了解相关的语言变化，我们也会抚平扩展的不同依赖关系之间的粗糙边缘。

Please continue to share your thoughts with us by filing [issues](https://github.com/golang/vscode-go/issues) and making [contributions](https://github.com/golang/vscode-go/blob/master/docs/contributing.md) to the project. The process for contributing will now be the same as for the [rest of the Go project](https://go.dev/doc/contribute.html). Go team members will offer general help in the #vscode channel on [Gophers Slack](https://invite.slack.golangbridge.org/), and we’ve also created a #vscode-dev channel to discuss issues and brainstorm ideas with contributors.

请继续与我们分享您的想法，提出问题并为项目做出贡献。现在，贡献的过程将与Go项目的其他部分相同。Go团队成员将在Gophers Slack的#vscode频道中提供一般帮助，我们还创建了一个#vscode-dev频道，与贡献者讨论问题并集思广益。

We’re excited about this new step forward, and we hope you are too. By maintaining a major Go editor extension, as well as the Go tooling and language, the Go team will be able to provide all Go users, regardless of their editor, a more cohesive and refined development experience.

我们对这一新的进展感到兴奋，我们希望您也是如此。通过维护一个主要的Go编辑器扩展，以及Go工具和语言，Go团队将能够为所有的Go用户，无论他们使用何种编辑器，提供一个更有凝聚力和更完善的开发体验。

As always, our goal remains the same: Every user should have an excellent experience writing Go code.

一如既往，我们的目标仍然是：每个用户都应该有一个优秀的Go代码编写体验。

*See the accompanying post from the [Visual Studio Code team](https://aka.ms/go-blog-vscode-202006).*

请看Visual Studio Code团队的附文。
