+++
title = "2014年FOSDEM上的 Go 讲座"
weight = 14
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go talks at FOSDEM 2014 - 2014年FOSDEM上的 Go 讲座

> 原文：[https://go.dev/blog/fosdem14](https://go.dev/blog/fosdem14)

Andrew Gerrand
24 February 2014

## Introduction 简介

At [FOSDEM](http://fosdem.org/) on the 2nd of February 2014 members of the Go community presented a series of talks in the Go Devroom. The day was a huge success, with 13 great talks presented to a consistently jam-packed room.

在2014年2月2日的FOSDEM上，Go社区的成员们在Go Devroom上进行了一系列的会谈。这一天取得了巨大的成功，有13场精彩的演讲，房间里一直挤满了人。

Video recordings of the talks are now available, and a selection of these videos are presented below.

讲座的录像现在已经可以获得，下面是这些录像的一部分。

The complete series of talks is available [as a YouTube playlist](http://www.youtube.com/playlist?list=PLtLJO5JKE5YDKG4WcaNts3IVZqhDmmuBH). (You can also get them directly at the [FOSDEM video archive](http://video.fosdem.org/2014/K4601/Sunday/).)

完整的会谈系列可作为YouTube播放列表。(您也可以在FOSDEM的视频档案中直接获得这些视频）。

## Scaling with Go: YouTube’s Vitess 用Go进行扩展：YouTube的Vitess

Google Engineer Sugu Sougoumarane described how he and his team built [Vitess](https://github.com/youtube/vitess) in Go to help scale [YouTube](https://youtube.com/).

谷歌工程师Sugu Sougoumarane介绍了他和他的团队如何在Go中建立Vitess来帮助扩展YouTube。

Vitess is a set of servers and tools primarily developed in Go. It helps scale MySQL databases for the web, and is currently used as a fundamental component of YouTube’s MySQL infrastructure.

Vitess是一套主要用Go开发的服务器和工具。它有助于为网络扩展MySQL数据库，目前被用作YouTube的MySQL基础设施的一个基本组成部分。

The talk covers some history about how and why the team chose Go, and how it paid off. Sugu also talks abou tips and techniques used to scale Vitess using Go.

讲座涵盖了一些关于团队如何和为什么选择Go的历史，以及它是如何得到回报的。Sugu还谈到了使用Go来扩展Vitess的技巧和方法。

{{< youtube "qATTTSg6zXk">}}

The slides for the talk are [available here](https://github.com/youtube/vitess/blob/master/doc/Vitess2014.pdf?raw=true).

讲座的幻灯片可以在这里找到。

## Camlistore

[Camlistore](http://camlistore.org/) is designed to be "your personal storage system for life, putting you in control, and designed to last." It’s open source, under nearly 4 years of active development, and extremely flexible. In this talk, Brad Fitzpatrick and Mathieu Lonjaret explain why they built it, what it does, and talk about its design.

Camlistore被设计成 "您的个人生活存储系统，让您掌控，并且设计得很持久"。它是开源的，经过近4年的积极开发，而且非常灵活。在这个讲座中，Brad Fitzpatrick和Mathieu Lonjaret解释了他们为什么要建造它，它的作用，并谈到了它的设计。

{{< youtube "yvjeIZgykiA">}}

## Write your own Go compiler 编写您自己的Go编译器

Elliot Stoneham explains the potential for Go as a portable language and reviews the Go tools that make that such an exciting possibility.

Elliot Stoneham解释了Go作为一种可移植语言的潜力，并回顾了使这种可能性变得如此激动人心的Go工具。

He said: "Based on my experiences writing an experimental Go to Haxe translator, I’ll talk about the practical issues of code generation and runtime emulation required. I’ll compare some of my design decisions with those of two other Go compiler/translators that build on the go.tools library. My aim is to encourage you to try one of these new ‘mutant’ Go compilers. I hope some of you will be inspired to contribute to one of them or even to write a new one of your own."

他说："根据我编写实验性Go到Haxe翻译器的经验，我将谈论代码生成和运行时仿真所需的实际问题。我将把我的一些设计决定与其他两个建立在go.tools库上的Go编译器/翻译器进行比较。我的目的是鼓励您尝试这些新的 "变种 "Go编译器中的一个。我希望您们中的一些人会受到启发，为其中一个编译器做出贡献，甚至自己编写一个新的编译器"。

{{< youtube "Qe8Dq7V3hXY">}}

## More 更多

There were many more great talks, so please check out the complete series [as a YouTube playlist](http://www.youtube.com/playlist?list=PLtLJO5JKE5YDKG4WcaNts3IVZqhDmmuBH). In particular, the [lightning talks](http://www.youtube.com/watch?v=cwpI5ONWGxc&list=PLtLJO5JKE5YDKG4WcaNts3IVZqhDmmuBH&index=7) were a lot of fun.

还有更多精彩的演讲，所以请查看YouTube上的完整系列播放列表。特别是，闪电式会谈是非常有趣的。

I would like to give my personal thanks to the excellent speakers, Mathieu Lonjaret for managing the video gear, and to the FOSDEM staff for making all this possible.

我个人要感谢这些优秀的演讲者，感谢Mathieu Lonjaret对视频设备的管理，感谢FOSDEM的工作人员使这一切成为可能。
