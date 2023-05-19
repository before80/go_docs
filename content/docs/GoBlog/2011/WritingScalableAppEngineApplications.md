+++
title = "编写可扩展的App Engine应用程序"
weight = 5
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Writing scalable App Engine applications - 编写可扩展的App Engine应用程序

https://go.dev/blog/appengine-scalable

David Symonds
1 November 2011

Back in May, we [announced](https://blog.golang.org/2011/05/go-and-google-app-engine.html) the Go runtime for App Engine. Since then, we’ve opened it up for everyone to use, added many new APIs, and improved performance. We have been thrilled by all the interesting ways that people are using Go on App Engine. One of the key benefits of the Go runtime, apart from working in a fantastic language, is that it has high performance. Go applications compile to native code, with no interpreter or virtual machine getting between your program and the machine.

早在五月，我们就宣布了App Engine的Go运行时。从那时起，我们将它开放给所有人使用，增加了许多新的API，并提高了性能。我们对人们在App Engine上使用Go的所有有趣方式感到非常兴奋。Go运行时的主要好处之一，除了在一种奇妙的语言中工作之外，就是它的高性能。Go应用程序被编译为本地代码，在你的程序和机器之间没有解释器或虚拟机。

Making your web application fast is important because it is well known that a web site’s latency has a measurable impact on user happiness, and [Google web search uses it as a ranking factor](https://googlewebmastercentral.blogspot.com/2010/04/using-site-speed-in-web-search-ranking.html). Also announced in May was that App Engine would be [leaving its Preview status](http://googleappengine.blogspot.com/2011/05/year-ahead-for-google-app-engine.html) and transitioning to a [new pricing model](https://www.google.com/enterprise/cloud/appengine/pricing.html), providing another reason to write efficient App Engine applications.

使你的网络应用程序快速是很重要的，因为众所周知，一个网站的延迟对用户的幸福感有可衡量的影响，而谷歌网络搜索也将其作为一个排名因素。5月份还宣布，App Engine将脱离预览状态，过渡到新的定价模式，为编写高效的App Engine应用程序提供了另一个理由。

To make it easier for Go developers using App Engine to write highly efficient, scalable applications, we recently updated some existing App Engine articles to include snippets of Go source code and to link to relevant Go documentation.

为了使使用App Engine的Go开发者更容易编写高效、可扩展的应用程序，我们最近更新了一些现有的App Engine文章，以包括Go源代码的片段，并链接到相关的Go文档。

- [Best practices for writing scalable applications 编写可扩展应用程序的最佳实践](http://code.google.com/appengine/articles/scaling/overview.html)
- [Managing Your App’s Resource Usage 管理你的应用程序的资源使用情况](http://code.google.com/appengine/articles/managing-resources.html)
