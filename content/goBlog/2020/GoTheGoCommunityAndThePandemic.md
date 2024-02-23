+++
title = "Go、go社区界和大流行病"
weight = 12
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go, the Go Community, and the Pandemic - go、go社区界和大流行病

> 原文：[https://go.dev/blog/pandemic](https://go.dev/blog/pandemic)

Carmen Andoh, Russ Cox, and Steve Francia
25 March 2020

Go always comes second to more basic concerns like personal and family health and safety. Around the world, the past couple months have been terrible, and we are still at the start of this awful pandemic. There are days when it seems like working on anything related to Go should be considered a serious priority inversion.

Go总是排在个人和家庭健康和安全等更基本的关切之后。在全世界范围内，过去几个月的情况很糟糕，我们仍然处于这场可怕的大流行病的开始阶段。有那么几天，似乎从事任何与Go有关的工作都应该被认为是严重的优先倒置。

But after we’ve done all we can to prepare ourselves and our families for whatever is coming, getting back to some approximation of a familiar routine and normal work is a helpful coping mechanism. In that spirit, we intend to keep working on Go and trying to help the Go community as much as we can.

但是，在我们竭尽所能为自己和家人准备好应对即将到来的一切之后，回到某种近似于熟悉的常规和正常工作的状态是一种有益的应对机制。本着这种精神，我们打算继续从事Go工作，并尽可能地帮助Go界。

In this post we want to share a few important notes about how the pandemic is affecting the Go community, a few things we’re doing to help, what you can do to help, and our plans for Go itself.

在这篇文章中，我们想与大家分享一些重要的注意事项，包括大流行病对Go社区的影响，我们正在做的一些事情，您可以做什么来帮助我们，以及我们对Go本身的计划。

## Conferences and Meetups 会议和聚会

The Go community thrives on in-person conferences and meetups. We had anticipated 35 conferences this year and thousands of meetups, nearly all of which have now changed, been postponed, or been canceled. We’ll keep the [conferences wiki page](https://github.com/golang/go/wiki/Conferences) updated as plans change.

Go社区的发展离不开亲自参加的会议和聚会。我们预计今年会有35个会议和数千个聚会，现在几乎所有的会议都已经改变了，或者被推迟了，或者被取消了。随着计划的变化，我们会随时更新会议的维基页面。

We want to do everything we can to help support impacted Go conferences. We also want to support efforts to explore new ways for gophers to connect in the time of social distancing. In addition to honoring Google’s existing sponsorships, we are interested to offer support to people planning virtual conference alternatives through the rest of the year. If you are organizing a Go conference and have been impacted, or if you are considering holding a virtual alternative, please reach out to Carmen Andoh at *candoh@google.com*.

我们希望尽我们所能，帮助支持受影响的Go会议。我们也希望支持在社会疏远的时代，为地鼠探索新的联系方式的努力。除了尊重谷歌现有的赞助外，我们有兴趣在今年余下的时间里为策划虚拟会议的人提供支持。如果您正在组织Go会议并受到影响，或者您正在考虑举办一个虚拟的替代方案，请联系Carmen Andoh，candoh@google.com。

For conference organizers, the [Gophers slack](https://gophers.slack.com/) [#conf-organizers](https://app.slack.com/client/T029RQSE6/C97B0NCVD) channel is a place to discuss contingency plans, best practices, cancellation, and postponement support. It’s also a place to share idea for virtual events, to continue to connect and support the Go community.

对于会议组织者来说，Gophers slack #conf-organizers频道是一个讨论应急计划、最佳实践、取消和推迟支持的地方。它也是一个分享虚拟活动想法的地方，以继续连接和支持Go社区。

For meetup organizers, the [Go Developer Network](https://www.meetup.com/pro/go) can provide Zoom for Education licensing to meetups that want to start holding virtual meetings. If you host a meetup, or you’d like to, we encourage you to use this opportunity to get speakers from outside your region to present to your group. For more information, and to get involved, please join the [Gophers slack](https://gophers.slack.com/) [#remotemeetup](https://app.slack.com/client/T029RQSE6/C152YB9UZ) channel.

对于聚会的组织者，Go开发者网络可以为想要开始举行虚拟会议的聚会提供Zoom for Education的许可。如果您举办聚会，或者您想举办，我们鼓励您利用这个机会，让您所在地区以外的演讲者为您的小组做演讲。欲了解更多信息，并参与其中，请加入Gophers slack #remotemeetup频道。

## Online Training 在线培训

The Go trainers you meet at conferences also travel the globe doing [in-person training](https://learn.go.dev/) for companies that want help adopting Go. That in-person teaching is crucial to bringing new gophers into the community; we’re incredibly grateful to the trainers for the work they do. Unfortunately, on-site training contracts have all been canceled for the next few months, and the trainers in our community have lost their primary (or sole) source of income. We encourage companies to consider virtual training and workshops during this difficult time. Most trainers are being flexible with pricing, scheduling, and class structure.

您在会议上遇到的Go培训师也在全球范围内为那些需要帮助采用Go的公司进行面对面的培训。这种面对面的教学对于将新的地鼠带入社区至关重要；我们非常感谢培训师们所做的工作。不幸的是，未来几个月的现场培训合同都被取消了，我们社区的培训师也失去了他们主要（或唯一）的收入来源。我们鼓励公司在这个困难时期考虑虚拟培训和研讨会。大多数培训师在价格、时间安排和课程结构方面都很灵活。

## Job Postings 招聘信息

We know that the current downturn means that some gophers are looking for new jobs. The Go community has built a number of Go-specific job-posting sites, including [Golang Cafe](https://golang.cafe/), [Golang Projects](https://www.golangprojects.com/), and [We Love Go](https://www.welovegolang.com/). The [Gophers slack](https://gophers.slack.com/) also has many job-hunting channels: search for "job" in the channel list. We encourage employers with any new openings to post in as many appropriate places as possible.

我们知道，目前的经济衰退意味着一些地鼠正在寻找新的工作。Go社区已经建立了一些专门针对Go的工作发布网站，包括Golang Cafe、Golang Projects和We Love Go。Gophers slack也有很多求职频道：在频道列表中搜索 "工作"。我们鼓励有新的职位空缺的雇主在尽可能多的适当地方发布。

## FOSS Responders - FOSS响应者

We are proud that Go is part of the broader open-source ecosystem. [FOSS Responders](https://fossresponders.com/) is one effort to help the open-source ecosystem deal with the impacts of the pandemic. If you want to do something to help affected open-source communities, they are coordinating efforts and also have links to other efforts. And if you know of other open-source communities that need help, let them know about FOSS Responders.

我们为Go是更广泛的开源生态系统的一部分而感到自豪。FOSS Responders是帮助开源生态系统应对大流行病影响的一项努力。如果您想做些什么来帮助受影响的开源社区，他们正在协调各种努力，也有其他努力的链接。如果您知道其他需要帮助的开源社区，让他们知道FOSS Responders。

## COVID-19 Open-Source Help Desk - COVID-19开源服务台

The [COVID-19 Open-Source Help Desk](https://covid-oss-help.org/) aims to help virologists, epidemiologists, and other domain experts find quick answers to any problems they are having with open-source scientific computing software, from experts in that software, so they can focus their time on what they know best. If you are a developer or a scientific computing expert willing to help by answering the posts of the domain experts, visit the site to learn how to help.

COVID-19开源帮助台旨在帮助病毒学家、流行病学家和其他领域的专家从开源科学计算软件的专家那里快速找到他们遇到的任何问题的答案，这样他们就可以把时间集中在他们最熟悉的领域。如果您是一名开发人员或科学计算专家，愿意通过回答领域专家的帖子来提供帮助，请访问该网站，了解如何提供帮助。

## U.S. Digital Response 美国的数字响应

For our gophers in the United States, the [U.S. Digital Response](https://www.usdigitalresponse.org/) is working to connect qualified volunteers to state and local governments that need digital help during this crisis. Quoting the web page, "If you have relevant experience (healthcare, data, engineering & product development, general management, operations, supply chain/procurement and more), can work autonomously through ambiguity, and are ready to jump into a high-intensity environment," see the site for how to volunteer.

对于我们在美国的gophers ，美国数字响应正在努力将合格的志愿者与在这次危机中需要数字帮助的州和地方政府联系起来。引用网页上的话，"如果您有相关的经验（医疗保健、数据、工程和产品开发、一般管理、运营、供应链/采购等），能够在模糊不清的情况下自主工作，并准备好跳入一个高强度的环境，"请参阅该网站，了解如何成为志愿者。

## Plans for Go - Go的计划

Here on the Go team at Google, we recognize that the world around us is changing rapidly and that plans beyond the next couple weeks are not much more than hopeful guesses. That said, right now we are working on what we think are the most important projects for 2020. Like all of you, we’re at reduced capacity, so the work continues slower than planned.

在谷歌的Go团队中，我们认识到我们周围的世界正在迅速变化，未来几周之后的计划也不过是充满希望的猜测。也就是说，现在我们正致力于我们认为是2020年最重要的项目。像您们所有人一样，我们的能力下降了，所以工作继续的速度比计划的慢。

Our analysis of the Go 2019 user survey is almost complete, and we hope to post it soon.

我们对Go 2019年用户调查的分析已基本完成，我们希望能很快发布。

At least for now, we intend to keep to our timeline for Go 1.15, with the understanding that it will probably have fewer new features and improvements than we originally planned. We continue to do code reviews, issue triage, and [proposal review](https://go.dev/s/proposal-minutes).

至少现在，我们打算遵守Go 1.15的时间表，但有一点可以理解，即它的新功能和改进可能会比我们最初计划的少。我们将继续进行代码审查、问题分流和提案审查。

[Gopls](https://go.googlesource.com/tools/+/refs/heads/master/gopls/README.md) is the language-aware backend supporting most Go editors today, and we continue to work toward its 1.0 release.

Gopls是目前支持大多数Go编辑器的语言识别后端，我们将继续努力实现其1.0版本。

The new Go package and module site [pkg.go.dev](https://pkg.go.dev/) keeps getting better. We’ve been working on usability improvements and new features to better help users find and evaluate Go packages. We’ve also expanded the set of recognized licenses and improved the license detector, with more improvements to come.

新的Go软件包和模块网站pkg.go.dev不断完善。我们一直在努力改进可用性和新功能，以更好地帮助用户寻找和评估Go软件包。我们还扩大了可识别的许可证集，并改进了许可证检测器，还有更多的改进将陆续推出。

Our [Gopher values](https://go.dev/conduct#values) are what ground us, now more than ever. We are working extra hard to be friendly, welcoming, patient, thoughtful, respectful, and charitable. We hope everyone in the Go community will try to do the same.

我们的Gopher价值观是我们的基础，现在比以往任何时候都要重要。我们正在努力做到友好、热情、耐心、周到、尊重和慈善。我们希望Go社区的每一个人都能努力做到这一点。

We’ll continue to use this blog to let you know about important news for the Go ecosystem. In those moments when you’ve taken care of the much more important things going on in your life, we hope you’ll check in and see what we’ve been up to.

我们将继续通过这个博客让您了解Go生态系统的重要消息。在您处理完您生活中更重要的事情的时候，我们希望您能来看看我们在做什么。

Thank you, as always, for using Go and being part of the Go community. We wish you all the best in these difficult times.

感谢您们使用Go，感谢您们成为Go社区的一员。在这个困难时期，我们祝愿您一切顺利。
