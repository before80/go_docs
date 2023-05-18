+++
title = "整理 go 的Web体验"
weight = 90
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Tidying up the Go web experience - 整理 go 的Web体验

[https://go.dev/blog/tidy-web](https://go.dev/blog/tidy-web)

Russ Cox
18 August 2021

In 2019, which seems like a decade ago, we [launched go.dev](https://go.dev/blog/go.dev), a new hub for Go developers, along with the [companion site pkg.go.dev](https://pkg.go.dev/), providing information about Go packages and modules.

​	在2019年，似乎是十年前，我们推出了[go.dev](../../2021/GodevANewHubForGoDevelopers)，一个新的Go开发者中心，以及提供Go包和模块信息的配套网站[pkg.go.dev](https://pkg.go.dev/)。

The go.dev web site contains useful information for people evaluating Go, but golang.org continued to serve distribution downloads, documentation, and a package reference for the standard library. Other sites — blog.golang.org, play.golang.org, talks.golang.org, and tour.golang.org — hold additional material. It’s all a bit fragmented and confusing.

​	`go.dev`网站包含了对评估Go的人有用的信息，但`golang.org`继续提供分发下载、文档和标准库的包参考。其他网站 —— `blog.golang.org`、`play.golang.org`、`talk.golang.org`和`tour.golang.org` —— 都有额外的材料。这一切都有点零散和混乱。

Over the next month or two we will be merging the golang.org sites into a single coherent web presence, here on go.dev. You may have already noticed that links to the package reference docs for the standard library on golang.org/pkg now redirect to their [equivalents on pkg.go.dev](https://pkg.go.dev/std), which is a better experience today and will continue to improve. As the next step, the Go blog has moved to go.dev/blog, starting with the post you are reading right now. (Of course, all the old blog posts are here too.)

​	在接下来的一两个月里，我们将把`golang.org`的网站合并成一个统一的网站，在`go.dev`上。你可能已经注意到，`golang.org/pkg`上的标准库参考文档的链接现在会重定向到[pkg.go.dev](https://pkg.go.dev/std)上的相应内容，这在今天是一个更好的体验，并将继续改进。作为下一步，Go博客已经搬到了`go.dev/blog`，从你现在正在阅读的这篇文章开始。(当然，所有的旧博客文章也在这里）。

As we move the content to its new home on go.dev, rest assured that all existing URLs will redirect to their new homes: no links will be broken.

​	当我们把内容转移到`go.dev`的新家时，请放心，所有现有的URL都会重定向到它们的新家：没有链接会被破坏。

We are excited to have a single coherent web site where everyone can find what they need to know about Go. It’s a small detail, but one long overdue.

​	我们很高兴能有一个统一的网站，每个人都能在这里找到他们需要了解的关于Go的内容。这是个小细节，但早该如此了。

If you have any ideas or suggestions, or you run into problems, please let us know via the “Report an Issue” link at the bottom of every page. Thanks!

​	如果你有任何想法或建议，或遇到问题，请通过每个页面底部的 "Report an Issue（报告问题）"链接让我们知道。谢谢!
