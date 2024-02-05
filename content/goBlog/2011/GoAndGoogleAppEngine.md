+++
title = "go 和 Google App Engine"
weight = 21
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go and Google App Engine - go 和 Google App Engine

> 原文：[https://go.dev/blog/appengine](https://go.dev/blog/appengine)

David Symonds, Nigel Tao, and Andrew Gerrand
10 May 2011

Google’s App Engine provides a reliable, scalable, easy way to build and deploy applications for the web. Over a hundred thousand apps are hosted at appspot.com and custom domains using the App Engine infrastructure. Originally written for Python apps, in 2009 the system added a Java runtime. And today, at Google I/O, we’re thrilled to announce that Go will be next. It’s marked as an experimental App Engine feature for now, because it’s early days, but both the App Engine and Go teams are very excited about this milestone.

Google的App Engine为构建和部署网络应用提供了可靠、可扩展、简便的方式。超过十万个应用程序被托管在appspot.com和使用App Engine基础设施的自定义域。该系统最初是为Python应用程序编写的，在2009年增加了一个Java运行时。今天，在谷歌I/O大会上，我们很高兴地宣布Go将是下一个。目前它被标记为一个实验性的App Engine功能，因为它还处于早期阶段，但App Engine和Go团队都对这个里程碑感到非常兴奋。

By early days, we mean that it’s still rolling out. As of today, the App Engine SDK for Go is [available for download](http://code.google.com/p/googleappengine/downloads/list), and we will soon enable deployment of Go apps into the App Engine hosting infrastructure. Today, through the SDK, you’ll be able to write web apps, learn about the APIs (and the language, if it’s new to you), and run your web app locally. Once full deployment is enabled, it’ll be easy to push your app to Google’s cloud.

我们所说的早期，是指它仍在推出。从今天起，适用于Go的App Engine SDK可供下载，而且我们很快就能将Go应用部署到App Engine的托管基础设施中。今天，通过SDK，您将能够编写Web应用程序，了解API（以及语言，如果它对您是新的），并在本地运行您的Web应用程序。一旦启用全面部署，就可以很容易地将您的应用程序推到谷歌的云端。

One of the cool but less obvious things about this news is that it provides a very easy way to play with Go. You don’t even need to have Go installed beforehand because the SDK is fully self-contained. Just download the SDK, unzip it, and start coding. Moreover, the SDK’s "dev app server" means you don’t even need to run the compiler yourself; everything is delightfully automatic.

这个消息的一个很酷但不太明显的地方是，它提供了一个非常简单的方法来玩Go。您甚至不需要事先安装Go，因为SDK是完全独立的。只要下载SDK，解压，然后开始编码。此外，SDK的 "开发应用服务器 "意味着您甚至不需要自己运行编译器；一切都令人愉快地自动进行。

What you’ll find in the SDK is many of the standard App Engine APIs, custom designed in good Go style, including Datastore, Blobstore, URL Fetch, Mail, Users, and so on. More APIs will be added as the environment develops. The runtime provides the full Go language and almost all the standard libraries, except for a few things that don’t make sense in the App Engine environment. For instance, there is no `unsafe` package and the `syscall` package is trimmed. (The implementation uses an expanded version of the setup in the [Go Playground](https://go.dev/doc/play/) on [golang.org](https://go.dev/).)

您会在SDK中发现许多标准的App Engine API，以良好的Go风格定制设计，包括Datastore、Blobstore、URL Fetch、Mail、Users等。随着环境的发展，更多的API将被加入。该运行时提供了完整的Go语言和几乎所有的标准库，除了一些在App Engine环境中没有意义的东西。例如，没有不安全包，系统调用包也被裁减了。(这个实现使用了golang.org上Go Playground的一个扩展版本的设置)。

Also, although goroutines and channels are present, when a Go app runs on App Engine only one thread is run in a given instance. That is, all goroutines run in a single operating system thread, so there is no CPU parallelism available for a given client request. We expect this restriction will be lifted at some point.

此外，尽管存在goroutines和通道，但当Go应用在App Engine上运行时，在一个给定的实例中只运行一个线程。也就是说，所有的goroutines都在一个操作系统线程中运行，所以对于一个给定的客户端请求来说，没有CPU并行性可言。我们希望这个限制在某个时候会被取消。

Despite these minor restrictions, it’s the real language: Code is deployed in source form and compiled in the cloud using the 64-bit x86 compiler (6g), making it the first true compiled language that runs on App Engine. Go on App Engine makes it possible to deploy efficient, CPU-intensive web applications.

尽管有这些小限制，但这是真正的语言。代码以源码形式部署，并在云端使用64位x86编译器（6g）进行编译，使其成为第一种在App Engine上运行的真正编译语言。App Engine上的Go使部署高效、CPU密集型的网络应用成为可能。

If you want to know more, read the [documentation](http://code.google.com/appengine/docs/go/) (start with "[Getting Started](http://code.google.com/appengine/docs/go/gettingstarted/)"). The libraries and SDK are open source, hosted at http://code.google.com/p/appengine-go/. We’ve created a new [google-appengine-go](http://groups.google.com/group/google-appengine-go) mailing list; feel free to contact us there with App Engine-specific questions. The [issue tracker for App Engine](http://code.google.com/p/googleappengine/issues/list) is the place for reporting issues related to the new Go SDK.

如果您想了解更多，请阅读文档（从 "入门 "开始）。这些库和SDK是开源的，托管在http://code.google.com/p/appengine-go/。我们已经创建了一个新的google-appengine-go邮件列表；如果有App Engine的具体问题，请随时与我们联系。App Engine的问题跟踪器是报告与新的Go SDK有关的问题的地方。

The Go App Engine SDK is [available](http://code.google.com/p/googleappengine/downloads/list) for Linux and Mac OS X (10.5 or greater); we hope a Windows version will also be available soon.

Go App Engine SDK可用于Linux和Mac OS X（10.5或更高版本）；我们希望Windows版本也能很快推出。

We’d like to offer our thanks for all the help and enthusiasm we received from Google’s App Engine team in making this happen.

我们要感谢Google App Engine团队在实现这一目标过程中给予我们的所有帮助和热情。
