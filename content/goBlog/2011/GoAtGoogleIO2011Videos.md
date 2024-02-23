+++
title = "谷歌2011年I/O大会上的 Go：视频"
weight = 20
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go at Google I/O 2011: videos - 谷歌2011年I/O大会上的 go：视频

> 原文：[https://go.dev/blog/io2011](https://go.dev/blog/io2011)

Andrew Gerrand
23 May 2011

## Introduction 简介

The Go team had a great time at Google I/O 2011. It was a pleasure to meet so many programmers who share our enthusiasm for Go, and to share our work of the past few months. For those of you that couldn’t be there in person, you can now watch videos of our two Go presentations on YouTube.

Go团队在Google I/O 2011上度过了一段美好的时光。很高兴能见到这么多与我们一样热衷于Go的程序员，并分享我们过去几个月的工作。对于那些不能亲临现场的人来说，您现在可以在YouTube上观看我们两个Go演讲的视频。

## Writing Web Apps in Go 用Go编写网络应用程序

In "[Writing Web Apps in Go](http://www.youtube.com/watch?v=-i0hat7pdpk)" we announce the [Go runtime for Google App Engine](https://blog.golang.org/2011/05/go-and-google-app-engine.html) and walk through the development and deployment of [Moustachio](http://moustach-io.appspot.com/), the first Go App Engine app.

在 "用Go编写网络应用程序 "中，我们宣布了Google App Engine的Go运行时间，并介绍了Moustachio的开发和部署，这是第一个Go App Engine应用程序。

<iframe src="https://www.youtube.com/embed/-i0hat7pdpk" width="560" height="315" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

(See the [presentation slides](https://go.dev/doc/talks/io2011/Writing_Web_Apps_in_Go.pdf).)

(见演讲幻灯片)。

The [source code](https://code.google.com/p/appengine-go/source/browse/example/moustachio) for Moustachio is available as part of [the SDK](http://code.google.com/appengine/downloads.html#Google_App_Engine_SDK_for_Go) along with some other examples, such as this [Mandelbrot demo](http://mandelbrot-tiles.appspot.com/).

Moustachio的源代码可以作为SDK的一部分，还有其他一些例子，比如这个Mandelbrot演示。

Most important, this talk features the debut of the plush gopher.

最重要的是，这次演讲中首次出现了毛绒地鼠。

![img](GoAtGoogleIO2011Videos_img/gopher.jpg)

For those that didn’t get one at the conference, we hope to make him available for purchase online soon.

对于那些没能在会议上得到一个的人，我们希望能很快在网上买到他。

## Real World Go 真实世界的Go

"[Real World Go](http://www.youtube.com/watch?v=7QDVRowyUQA)", presented at [I/O Bootcamp](http://io-bootcamp.com/), gives a brief introduction to Go and four case studies of its use in solving real problems:

在I/O Bootcamp上发表的 "Real World Go"，简要介绍了Go的情况，并对其在解决实际问题中的应用进行了四个案例分析。

- [Heroku](http://heroku.com/) with [Doozer](https://github.com/ha/doozerd), a highly available consistent data store, Heroku与Doozer，一个高可用的一致性数据存储。
  - [MROffice Dialer](http://mroffice.org/telephony.html), a VOIP system for call centers, MROffice Dialer，一个用于呼叫中心的VOIP系统。
  - [Atlassian](http://www.atlassian.com/)’s virtual machine cluster management system, Atlassian的虚拟机集群管理系统。
  - [Camlistore](http://www.camlistore.org/), a content addressable storage system. Camlistore，一个内容可寻址存储系统。

<iframe src="https://www.youtube.com/embed/7QDVRowyUQA" width="560" height="315" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

(See the [presentation slides](https://go.dev/doc/talks/io2011/Real_World_Go.pdf).)

(见演讲幻灯片)。

Thanks to everyone who attended our talks and workshops. We look forward to seeing you again soon!

感谢所有参加我们讲座和研讨会的人。我们期待着很快再次见到您
