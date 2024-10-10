+++
title = "pkg.go.dev的下一步工作"
weight = 15
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Next steps for pkg.go.dev - pkg.go.dev的下一步工作

> 原文：[https://go.dev/blog/pkg.go.dev-2020](https://go.dev/blog/pkg.go.dev-2020)

Julie Qiu
31 January 2020

## Introduction 简介

In 2019, we launched [go.dev](https://go.dev/), a new hub for Go developers.

在2019年，我们推出了go.dev，一个新的Go开发者中心。

As part of the site, we also launched [pkg.go.dev](https://pkg.go.dev/), a central source of information about Go packages and modules. Like [godoc.org](https://godoc.org/), pkg.go.dev serves Go documentation. However, it also understands modules and has information about past versions of a package!

作为网站的一部分，我们还推出了pkg.go.dev，这是一个关于Go软件包和模块的中央信息来源。与godoc.org一样，pkg.go.dev为Go文档提供服务。然而，它也了解模块，并有关于包的过去版本的信息!

Throughout this year, we will be adding features to [pkg.go.dev](https://pkg.go.dev/) to help our users better understand their dependencies and help them make better decisions around what libraries to import.

在今年，我们将为pkg.go.dev增加一些功能，以帮助我们的用户更好地了解他们的依赖关系，并帮助他们在导入哪些库方面做出更好的决定。

## Redirecting godoc.org requests to pkg.go.dev 将godoc.org的请求重定向到pkg.go.dev

To minimize confusion about which site to use, later this year we are planning to redirect traffic from [godoc.org](https://godoc.org/) to the corresponding page on [pkg.go.dev](https://pkg.go.dev/). We need your help to ensure that pkg.go.dev addresses all of our users' needs. We encourage everyone to begin using pkg.go.dev today for all of their needs and provide feedback.

为了减少对使用哪个网站的混淆，今年晚些时候，我们计划将来自godoc.org的流量重定向到pkg.go.dev的相应页面。我们需要您的帮助，以确保pkg.go.dev能满足我们所有用户的需求。我们鼓励大家今天开始使用pkg.go.dev来满足他们的所有需求，并提供反馈。

Your feedback will inform our transition plan, with the goal of making [pkg.go.dev](https://pkg.go.dev/) our primary source of information and documentation for packages and modules. We’re sure there are things that you want to see on pkg.go.dev, and we want to hear from you about what those features are!

您的反馈将为我们的过渡计划提供信息，目标是使pkg.go.dev成为我们的软件包和模块的主要信息和文档来源。我们确信在pkg.go.dev上有您想看到的东西，我们想听听您的意见，看看这些功能是什么。

You can share your feedback with us on these channels:

您可以在这些渠道与我们分享您的反馈：

- Post on the [Go issue tracker](https://go.dev/s/discovery-feedback).在 Go 问题跟踪器上发帖。
- Email [go-discovery-feedback@google.com](mailto:go-discovery-feedback@google.com). 电子邮件 go-discovery-feedback@google.com。
- Click "Share Feedback" or "Report an Issue" in the go.dev footer.点击go.dev页脚的 "分享反馈 "或 "报告问题"。

As part of this transition, we will also be discussing plans for API access to [pkg.go.dev](https://pkg.go.dev/). We will be posting updates on [Go issue 33654](https://go.dev/s/discovery-updates).

作为这一过渡的一部分，我们还将讨论对pkg.go.dev的API访问计划。我们将在Go问题33654上发布更新。

## Frequently asked questions 常见的问题

Since our launch in November, we’ve received tons of great feedback about [pkg.go.dev](https://pkg.go.dev/) from Go users. For the remainder of this post, we thought it would be helpful to answer some frequently asked questions.

自从我们在11月推出以来，我们收到了Go用户对pkg.go.dev的大量反馈。在这篇文章的剩余部分，我们认为回答一些常见问题会有帮助。

### My package doesn’t show up on pkg.go.dev! How do I add it? 我的软件包没有出现在 pkg.go.dev 上！我该如何添加？我怎样才能添加它？

We monitor the [Go Module Index](https://index.golang.org/index) regularly for new packages to add to [pkg.go.dev](https://pkg.go.dev/). If you don’t see a package on pkg.go.dev, you can add it by fetching the module version from [proxy.golang.org](https://proxy.golang.org/). See [go.dev/about](https://go.dev/about) for instructions.

我们会定期监控 Go 模块索引，以便将新的软件包添加到 pkg.go.dev 中。如果您在 pkg.go.dev 上没有看到某个软件包，您可以通过从 proxy.golang.org 获取模块版本来添加它。请参阅 go.dev/about 以了解说明。

### My package has license restrictions. What’s wrong with it? 我的软件包有许可证限制。这有什么问题吗？

We understand it can be a frustrating experience to not be able to see the package you want in its entirety on [pkg.go.dev](https://pkg.go.dev/). We appreciate your patience as we improve our license detection algorithm.

我们理解在 pkg.go.dev 上无法看到您想要的软件包的全部内容可能是一种令人沮丧的经历。我们感谢您的耐心，因为我们正在改进我们的许可证检测算法。

Since our launch in November, we’ve made the following improvements:

自从我们在11月推出以来，我们已经做了以下改进：

- Updated our [license policy](https://pkg.go.dev/license-policy) to include the list of licenses that we detect and recognize更新了我们的许可证政策，包括我们检测和识别的许可证清单。
- Worked with the [licensecheck](https://github.com/google/licensecheck) team to improve detection for copyright notices 与licensecheck团队合作，改进对版权声明的检测。
- Established a manual review process for special cases 建立了一个针对特殊情况的人工审查程序

As always, our license policy is at [pkg.go.dev/license-policy](https://pkg.go.dev/license-policy). If you are having issues, feel free to [file an issue on the Go issue tracker](https://go.dev/s/discovery-feedback), or email [go-discovery-feedback@google.com](mailto:go-discovery-feedback@google.com) so that we can work with you directly!

一如既往，我们的许可证政策在pkg.go.dev/license-policy。如果您有问题，请随时在Go问题跟踪器上提出问题，或发送电子邮件至 go-discovery-feedback@google.com，以便我们可以直接与您合作

### Will pkg.go.dev be open-sourced so I can run it at work for my private code? pkg.go.dev是否会被开源，以便我可以在工作中为我的私人代码运行它？

We understand that corporations with private code want to run a documentation server that provides module support. We want to help meet that need, but we feel we don’t yet understand it as well as we need to.

我们理解拥有私有代码的公司希望运行一个提供模块支持的文档服务器。我们想帮助满足这一需求，但我们觉得我们还没有充分理解这一需求。

We’ve heard from users that running the [godoc.org](https://godoc.org/) server is more complex than it should be, because it is designed for serving at public internet scale instead of just within a company. We believe the current [pkg.go.dev](https://pkg.go.dev/) server would have the same problem.

我们从用户那里听说，运行godoc.org服务器比它应该的更复杂，因为它是为公共互联网规模的服务而设计的，而不仅仅是在一个公司内部。我们相信目前的pkg.go.dev服务器也会有同样的问题。

We think a new server is more likely to be the right answer for use with private code, instead of exposing every company to the complexity of running the internet-scale [pkg.go.dev](https://pkg.go.dev/) codebase. In addition to serving documentation, a new server could also serve information to [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports?tab=doc) and [gopls](https://pkg.go.dev/golang.org/x/tools/gopls).

我们认为一个新的服务器更有可能是用于私人代码的正确答案，而不是让每个公司都暴露在运行互联网规模的pkg.go.dev代码库的复杂性中。除了提供文档，新的服务器还可以为 goimports 和 gopls 提供信息。

If you want to run such a server, please fill out this [**3-5 minute survey**](https://google.qualtrics.com/jfe/form/SV_6FHmaLveae6d8Bn) to help us better understand your needs. This survey will be available until March 1st, 2020.

如果您想运行这样一个服务器，请填写这个3-5分钟的调查，以帮助我们更好地了解您的需求。这项调查将持续到2020年3月1日。

We’re excited about the future of [pkg.go.dev](https://pkg.go.dev/) in 2020, and we hope you are too! We look forward to hearing your feedback and working with the Go community on this transition.

我们对2020年pkg.go.dev的未来感到兴奋，我们希望您也是如此！我们期待着听到您的反馈。我们期待着听到您的反馈，并与Go社区合作完成这一过渡。
