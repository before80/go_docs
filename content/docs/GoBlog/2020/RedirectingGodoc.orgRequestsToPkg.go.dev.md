+++
title = "将godoc.org的请求重定向到pkg.go.dev"
weight = 2
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Redirecting godoc.org requests to pkg.go.dev - 将godoc.org的请求重定向到pkg.go.dev

https://go.dev/blog/godoc.org-redirect

Julie Qiu
15 December 2020

With the introduction of Go modules and the growth of the Go ecosystem, [pkg.go.dev](https://pkg.go.dev/) was [launched in 2019](https://blog.golang.org/go.dev) to provide a central place where developers can discover and evaluate Go packages and modules. Like godoc.org, pkg.go.dev serves Go documentation, but it also supports modules, better search functionality, and signals to help Go users to find the right packages.

随着Go模块的引入和Go生态系统的发展，pkg.go.dev于2019年推出，为开发者提供一个发现和评估Go包和模块的中心场所。与godoc.org一样，pkg.go.dev为Go文档服务，但它也支持模块、更好的搜索功能和信号，以帮助Go用户找到正确的包。

As [we shared in January 2020](https://blog.golang.org/pkg.go.dev-2020), our goal is to eventually redirect traffic from godoc.org to the corresponding page on pkg.go.dev. We’ve also made it possible for users to opt in to redirecting their own requests from godoc.org to pkg.go.dev.

正如我们在2020年1月所分享的，我们的目标是最终将godoc.org的流量重定向到pkg.go.dev的相应页面上。我们也让用户有可能选择将他们自己的请求从godoc.org重定向到pkg.go.dev。

We’ve received a lot of great feedback this year, which has been tracked and resolved through the [pkgsite/godoc.org-redirect](https://github.com/golang/go/milestone/157?closed=1) and [pkgsite/design-2020](https://github.com/golang/go/milestone/159?closed=1) milestones on the Go issue tracker. Your feedback resulted in support for popular feature requests on pkg.go.dev, [open sourcing pkgsite](https://blog.golang.org/pkgsite), and most recently, a [redesign of pkg.go.dev](https://blog.golang.org/pkgsite-redesign).

今年我们收到了很多很好的反馈，这些反馈已经通过Go问题跟踪器上的pkgsite/godoc.org-redirect和pkgsite/design-2020里程碑进行跟踪和解决。您的反馈导致了对pkg.go.dev上流行的功能请求的支持，开放了pkgsite的资源，以及最近对pkg.go.dev的重新设计。

## Next Steps 接下来的步骤

The next step in this migration is to redirect all requests from godoc.org to the corresponding page on pkg.go.dev.

这次迁移的下一步是将所有请求从godoc.org重定向到pkg.go.dev上的相应页面。

This will happen in early 2021, once the work tracked at the [pkgsite/godoc.org-redirect milestone](https://github.com/golang/go/milestone/157) is complete.

这将在2021年初发生，一旦pkgsite/godoc.org-redirect里程碑的工作完成。

During this migration, updates will be posted to [Go issue 43178](https://go.dev/issue/43178).

在这个迁移过程中，更新将被发布到Go问题43178。

We encourage everyone to begin using pkg.go.dev today. You can do so by visiting [godoc.org?redirect=on](https://godoc.org/?redirect=on), or clicking "Always use pkg.go.dev" in the top right corner of any godoc.org page.

我们鼓励大家今天就开始使用pkg.go.dev。您可以通过访问godoc.org?redirect=on，或在任何godoc.org页面的右上角点击 "永远使用pkg.go.dev "来实现。

## FAQs 常见问题

**Will godoc.org URLs continue to work? **godoc.org的URL是否能继续使用？

Yes! We will redirect all requests arriving at godoc.org to the equivalent page on pkg.go.dev, so all your bookmarks and links will continue to take you to the documentation you need.

是的！我们将重定向所有进入godoc.org的请求。我们会将所有到达godoc.org的请求重定向到pkg.go.dev的相应页面，所以您所有的书签和链接都会继续带您到您需要的文档。

**What will happen to the golang/gddo repository? **golang/gddo资源库会发生什么？

The [gddo repository](http://go.googlesource.com/gddo) will remain available for anyone who wants to keep running it themselves, or even fork and improve it. We will mark it archived to make clear that we will no longer accept contributions. However, you will be able to continue forking the repository.

gddo 仓库将继续提供给任何想继续运行它的人，甚至是分叉和改进它。我们会把它归档，以表明我们将不再接受贡献。然而，您将能够继续分叉这个仓库。

**Will api.godoc.org continue to work? **api.godoc.org 还能继续工作吗？

This transition will have no impact on api.godoc.org. Until an API is available for pkg.go.dev, api.godoc.org will continue to serve traffic. See [Go issue 36785](https://go.dev/issue/36785) for updates on an API for pkg.go.dev.

这次过渡对 api.godoc.org 没有影响。在为pkg.go.dev提供API之前，api.godoc.org将继续提供流量。有关pkg.go.dev的API的更新，请参见Go问题36785。

**Will my godoc.org badges keep working? **我的godoc.org徽章会继续使用吗？

Yes! Badge URLs will redirect to the equivalent URL on pkg.go.dev too. Your page will automatically get a new pkg.go.dev badge. You can also generate a new badge at [pkg.go.dev/badge](https://pkg.go.dev/badge) if you would like to update your badge link.

是的！徽章的URL也会重定向到pkg.go.dev的相应URL。您的页面将自动获得一个新的pkg.go.dev徽章。如果您想更新您的徽章链接，您也可以在pkg.go.dev/badge生成一个新的徽章。

## Feedback 反馈意见

As always, feel free to [file an issue](https://go.dev/s/pkgsite-feedback) on the Go issue tracker for any feedback.

如同往常一样，如果有任何反馈，请随时在 Go 问题跟踪器上提交问题。

## Contributing 贡献

Pkg.go.dev is an [open source project](https://go.googlesource.com/pkgsite). If you’re interested in contributing to the pkgsite project, check out the [contribution guidelines](https://go.googlesource.com/pkgsite/+/refs/heads/master/CONTRIBUTING.md) or join the [#pkgsite channel](https://gophers.slack.com/messages/pkgsite) on Gophers Slack to learn more.

pkg.go.dev 是一个开源项目。如果您对pkgsite项目的贡献感兴趣，请查看贡献指南或加入Gophers Slack的#pkgsite频道以了解更多信息。
