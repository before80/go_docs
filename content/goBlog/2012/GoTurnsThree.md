+++
title = "Go 三周年"
weight = 1
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go turns three - Go 三周年

> 原文：[https://go.dev/blog/3years](https://go.dev/blog/3years)

Russ Cox
10 November 2012

The Go open source project is [three years old today](http://google-opensource.blogspot.com/2009/11/hey-ho-lets-go.html).

Go开源项目今天已经三岁了。

It’s great to look at how far Go has come in those three years. When we launched, Go was an idea backed by two implementations that worked on Linux and OS X. The syntax, semantics, and libraries changed regularly as we reacted to feedback from users and experience with the language.

看看Go在这三年里所取得的成就，真是太棒了。当我们启动时，Go只是一个由两个在Linux和OS X上运行的实现支持的想法。语法、语义和库经常发生变化，因为我们对用户的反馈和语言的经验做出了反应。

Since the open source launch, we’ve been joined by hundreds of external contributors, who have extended and improved Go in myriad ways, including writing a Windows port from scratch. We added a package management system [goinstall](https://groups.google.com/d/msg/golang-nuts/8JFwR3ESjjI/cy7qZzN7Lw4J), which eventually became the [go command](https://go.dev/cmd/go/). We also added [support for Go on App Engine](https://blog.golang.org/2011/07/go-for-app-engine-is-now-generally.html). Over the past year we’ve also given [many talks](https://go.dev/doc/#talks), created an [interactive introductory tour](https://go.dev/tour/) and recently we added support for [executable examples in package documentation](https://go.dev/pkg/strings/#pkg-examples).

自从开放源代码以来，我们已经有数百名外部贡献者加入，他们以各种方式扩展和改进 Go，包括从头开始编写 Windows 端口。我们增加了一个软件包管理系统goinstall，它最终成为go命令。我们还在 App Engine 上增加了对 Go 的支持。在过去的一年里，我们还举办了许多讲座，创建了一个互动的介绍性旅游，最近我们在包的文档中增加了对可执行实例的支持。

Perhaps the most important development in the past year was the launch of the first stable version, [Go 1](https://blog.golang.org/2012/03/go-version-1-is-released.html). People who write Go 1 programs can now be confident that their programs will continue to compile and run without change, in many environments, on a time scale of years. As part of the Go 1 launch we spent months cleaning up the [language and libraries](https://go.dev/doc/go1.html) to make it something that will age well.

也许去年最重要的发展是推出了第一个稳定版本--Go 1。编写Go 1程序的人现在可以确信，他们的程序将在许多环境中继续编译和运行，不会有任何变化，而且时间尺度长达数年。作为Go 1发布的一部分，我们花了几个月的时间来清理语言和库，以使其能够很好地使用。

We’re working now toward the release of Go 1.1 in 2013. There will be some new functionality, but that release will focus primarily on making Go perform even better than it does today.

我们现在正在努力争取在2013年发布Go 1.1。将会有一些新的功能，但该版本将主要关注于使Go的性能比现在更好。

We’re especially happy about the community that has grown around Go: the mailing list and IRC channels seem like they are overflowing with discussion, and a handful of Go books were published this year. The community is thriving. Use of Go in production environments has also taken off, especially since Go 1.

我们对围绕Go成长起来的社区感到特别高兴：邮件列表和IRC频道似乎都充斥着讨论，今年还出版了一些Go书籍。这个社区正在茁壮成长。Go在生产环境中的使用也在飞速发展，尤其是在Go 1之后。

We use Go at Google in a variety of ways, many of them invisible to the outside world. A few visible ones include [serving Chrome and other downloads](https://groups.google.com/d/msg/golang-nuts/BNUNbKSypE0/E4qSfpx9qI8J), [scaling MySQL database at YouTube](http://code.google.com/p/vitess/), and of course running the [Go home page](https://go.dev/) on [App Engine](https://developers.google.com/appengine/docs/go/overview). Last year’s [Thanksgiving Doodle](https://blog.golang.org/2011/12/from-zero-to-go-launching-on-google.html) and the recent [Jam with Chrome](http://www.jamwithchrome.com/technology) site are also served by Go programs.

我们在 Google 使用 Go 的方式多种多样，其中许多是外界看不到的。一些可见的方式包括为Chrome和其他下载提供服务，在YouTube扩展MySQL数据库，当然还有在App Engine上运行Go主页。去年的感恩节涂鸦和最近的Jam with Chrome网站也是由Go程序提供服务。

Other companies and projects are using Go too, including [BBC Worldwide](http://www.quora.com/Go-programming-language/Is-Google-Go-ready-for-production-use/answer/Kunal-Anand), [Canonical](http://dave.cheney.net/wp-content/uploads/2012/08/august-go-meetup.pdf), [CloudFlare](http://blog.cloudflare.com/go-at-cloudflare), [Heroku](https://blog.golang.org/2011/04/go-at-heroku.html), [Novartis](https://plus.google.com/114945221884326152379/posts/d1SVaqkRyTL), [SoundCloud](http://backstage.soundcloud.com/2012/07/go-at-soundcloud/), [SmugMug](http://sorcery.smugmug.com/2012/04/06/deriving-json-types-in-go/), [StatHat](https://blog.golang.org/2011/12/building-stathat-with-go.html), [Tinkercad](https://tinkercad.com/about/jobs), and [many others](https://go.dev/wiki/GoUsers).

其他公司和项目也在使用Go，包括BBC Worldwide、Canonical、CloudFlare、Heroku、Novartis、SoundCloud、SmugMug、StatHat、Tinkercad以及其他许多公司。



Here’s to many more years of productive programming in Go.

祝愿用Go进行编程的日子越来越好。
