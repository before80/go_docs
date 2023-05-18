+++
title = "最近的两篇 go 文章"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Two recent Go articles - 最近的两篇 go 文章

https://go.dev/blog/two-recent-go-articles

Andrew Gerrand
6 March 2013

## Introduction 简介

In today’s blog post I’d like to highlight a couple of recent articles about Go.

在今天的博文中，我想强调最近几篇关于Go的文章。

## Go at Google 谷歌的Go

In October last year, Rob Pike presented a keynote at the ACM [SPLASH](http://splashcon.org/2012/) conference in Tucson. The talk, titled [Go at Google](https://go.dev/talks/2012/splash.slide), was a comprehensive discussion of the motivations behind Go. Rob later expanded on his talk to produce an essay titled [Go at Google: Language Design in the Service of Software Engineering](http://go.dev/talks/2012/splash.article). Here is the abstract:

去年10月，Rob Pike在图森举行的ACM SPLASH会议上做了主题演讲。这次演讲的题目是Go at Google，全面讨论了Go背后的动机。后来，罗伯对他的演讲进行了扩展，写了一篇题为《Go at Google》的文章。为软件工程服务的语言设计。以下是其摘要：

```
The Go programming language was conceived in late 2007 as an answer to some of the problems we were seeing developing software infrastructure at Google. The computing landscape today is almost unrelated to the environment in which the languages being used, mostly C++, Java, and Python, had been
created. The problems introduced by multicore processors,networked systems, massive computation clusters, and the web programming model were being worked around rather than addressed head-on. Moreover, the scale has changed: today's server programs comprise tens of millions of lines of code,
are worked on by hundreds or even thousands of programmers,and are updated literally every day.  To make matters worse,build times, even on large compilation clusters, have stretched to many minutes, even hours.

Go was designed and developed to make working in this environment more productive. Besides its better-known aspects such as built-in concurrency and garbage collection,Go's design considerations include rigorous dependency management, the adaptability of software architecture as systems grow, and robustness across the boundaries between components.

Go编程语言是在2007年末被设想为 是为了解决我们在谷歌开发软件基础设施时遇到的一些问题。在谷歌开发软件基础设施时遇到的一些问题。今天的计算环境 今天的计算环境与以前使用的语言（主要是C++、Java和Python）几乎毫无关系。语言（主要是C++、Java和Python）所处的环境。
创造的环境。多核处理器所带来的问题。 多核处理器、网络化系统、大规模计算集群和网络 编程模型所带来的问题正在被解决，而不是 而不是正面解决。此外，规模已经改变：今天的 服务器程序由数千万行代码组成。
由数百甚至数千名程序员共同完成，并且每天都在更新。 更糟糕的是，即使在大型编译集群中，构建时间也延长到了几分钟，甚至几小时。

Go的设计和开发是为了使在这种环境下工作更有成效。除了内置的并发性和垃圾收集等众所周知的方面，Go的设计考虑包括严格的依赖性管理，随着系统的发展，软件架构的适应性，以及跨越组件之间边界的健壮性。
```

This article explains how these issues were addressed while building an efficient, compiled programming language that feels lightweight and pleasant. Examples and explanations will be taken from the real-world problems faced at Google.

本文解释了这些问题是如何被解决的，同时建立一个高效的、可编译的编程语言，并让人感到轻盈和愉快。例子和解释将取自Google所面临的真实世界的问题。

If you have wondered about the design decisions behind Go, you may find your questions answered by [the essay](https://go.dev/talks/2012/splash.article). It is recommended reading for both new and experienced Go programmers.

如果你想知道Go背后的设计决策，你可能会发现你的问题在文中得到了解答。它是推荐给新的和有经验的Go程序员阅读的。

## Go at the Google Developers Academy 谷歌开发者学院的Go

At Google I/O 2012 the Google Developers team [launched](http://googledevelopers.blogspot.com.au/2012/06/google-launches-new-developer-education.html) the [Google Developers Academy](https://developers.google.com/academy/), a program that provides training materials on Google technologies. Go is one of those technologies and we’re pleased to announce the first GDA article featuring Go front and center:

在2012年谷歌I/O大会上，谷歌开发者团队推出了谷歌开发者学院，这是一个提供谷歌技术培训材料的项目。Go是这些技术中的一种，我们很高兴地宣布第一篇以Go为中心的GDA文章：

[Getting Started with Go, App Engine and Google+ API](https://developers.google.com/appengine/training/go-plus-appengine/) is an introduction to writing web applications in Go. It demonstrates how to build and deploy App Engine applications and make calls to the Google+ API using the Google APIs Go Client. This is a great entry point for Go programmers eager to get started with Google’s developer ecosystem.

Go、App Engine和Google+ API入门》是一篇关于用Go编写网络应用的介绍。它演示了如何构建和部署App Engine应用程序，并使用Google APIs Go客户端调用Google+ API。对于渴望开始使用谷歌的开发者生态系统的Go程序员来说，这是一个很好的入门点。
