+++
title = "pkg.go.dev是开源的!"
weight = 9
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Pkg.go.dev is open source! - pkg.go.dev是开源的!

> 原文：[https://go.dev/blog/pkgsite](https://go.dev/blog/pkgsite)

Julie Qiu
15 June 2020

We’re excited to announce that the codebase for [pkg.go.dev](https://pkg.go.dev/) is now open source.

我们很高兴地宣布，pkg.go.dev的代码库已经开源了。

The repository lives at [go.googlesource.com/pkgsite](https://go.googlesource.com/pkgsite) and is mirrored to [github.com/golang/pkgsite](https://github.com/golang/pkgsite). We will continue using the Go issue tracker to track [feedback](https://github.com/golang/go/labels/go.dev) related to pkg.go.dev.

代码库位于go.googlesource.com/pkgsite，并被镜像到github.com/golang/pkgsite。我们将继续使用 Go 问题跟踪器来跟踪与 pkg.go.dev 有关的反馈。

## Contributing 贡献

If you are interested in contributing to any [issues related to pkg.go.dev](https://github.com/golang/go/labels/go.dev), check out our [contribution guidelines](https://go.googlesource.com/pkgsite/+/refs/heads/master/CONTRIBUTING.md). We also encourage you to continue [filing issues](https://go.dev/s/discovery-feedback) if you run into problems or have feedback.

如果您有兴趣对任何与 pkg.go.dev 有关的问题做出贡献，请查看我们的贡献指南。我们也鼓励您在遇到问题或有反馈时继续提交问题。

## What’s Next 下一步工作

We really appreciate all the feedback we’ve received so far. It has been a big help in shaping our [roadmap](https://go.googlesource.com/pkgsite#roadmap) for the coming year. Now that pkg.go.dev is open source, here’s what we’ll be working on next:

我们非常感谢到目前为止收到的所有反馈。这对我们制定来年的路线图有很大帮助。既然pkg.go.dev已经开源，下面是我们接下来要做的工作：

- We have some design changes planned for pkg.go.dev, to address [UX feedback](https://github.com/golang/go/issues?q=is%3Aissue+is%3Aopen+label%3Ago.dev+label%3AUX) that we have received. You can expect a more cohesive search and navigation experience. We plan to share these designs for feedback once they are ready.我们计划对pkg.go.dev进行一些设计修改，以解决我们收到的用户体验反馈。您可以期待一个更有凝聚力的搜索和导航体验。我们计划在这些设计完成后与大家分享，以获得反馈。
- We know that there are features available on godoc.org that users want to see on pkg.go.dev. We’ve been keeping track of them on [Go issue #39144](https://go.dev/issue/39144), and will prioritize adding them in the next few months. We also plan to continue improving our license detection algorithm based on feedback.我们知道，godoc.org上有一些用户希望在pkg.go.dev上看到的功能。我们一直在Go问题#39144上跟踪它们，并将在未来几个月内优先添加它们。我们还计划根据反馈继续改进我们的许可证检测算法。
- We’ll be improving our search experience based on feedback in [Go issue #37810](https://go.dev/issue/37810), to make it easier for users to find the dependencies they are looking for and make better decisions around which ones to import.我们将根据Go问题#37810中的反馈改进我们的搜索体验，以使用户更容易找到他们正在寻找的依赖项，并围绕哪些依赖项做出更好的决定。

Thanks for being patient with us in the process of open sourcing pkg.go.dev. We’re looking forward to receiving your contributions and working with you on the future of the project.

感谢您在开源pkg.go.dev的过程中对我们保持耐心。我们期待着收到您的贡献，并与您一起为项目的未来而努力。
