+++
title = "go 变得更加稳定"
weight = 27
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go becomes more stable - go 变得更加稳定

https://go.dev/blog/stable-releases

Andrew Gerrand
16 March 2011

2011年3月16日

​	Go 项目进展迅速。随着我们对 Go 的了解越来越深入，我们不得不改变我们的工具、库甚至是语言本身。我们允许不兼容的变化，这样我们就能从我们的错误中吸取教训，而不是永远留下它们。我们相信在 Go 发展的这个阶段，灵活性对于项目的发展以及最终的生命周期至关重要。

​	自从 Go 推出以来，我们已经每周发布一次更新版本。每个版本都会伴随着说明文档，[描述了发生了哪些变化](https://go.dev/doc/devel/release.html)，并标注了任何不兼容的更改。我经常听到的问题是："Go 稳定吗？我怎样才能确定我不必每周更新我的 Go 代码？" 现在的答案是 "是的" 和 "您不需要"。

​	在本周的更新版本中，我们引入了一个新的发布标签方案。我们计划继续每周发布更新版本，但将现有的标签从"release"更名为"weekly"。将"release"标签应用于每个月或两个月中挑选出来的一个稳定版本。这种更加宽松的发布计划应该会让普通的 Go 程序员生活更轻松。

​	用户仍需要定期更新他们的代码（这是使用年轻语言的代价），但更新的频率会更少。另一个好处是，我们可以更少地标记稳定版本，从而将更多的精力投入到自动化更新中。为此，我们推出了 gofix 工具，可以帮助您更新您的代码。

​	过去被标记为 release.2011-03-07.1（现在是 weekly.2011-03-07.1）的版本已经被评为我们的第一个稳定版本，并被赋予了 release.r56 标签。随着我们标记每个稳定版本，我们会在新的 [golang-announce](http://groups.google.com/group/golang-announce) 邮件列表中发布公告。（现在就为什么不[订阅](http://groups.google.com/group/golang-announce/subscribe))呢？）

​	所有这些的结果是什么？您仍然可以使用 hg update release 来保持您的 Go 安装更新，但现在您只需要在我们标记一个新的稳定版本时进行更新。如果您想保持领先，您应该使用 hg update weekly 切换到 weekly 标签。

​	祝编码愉快！
