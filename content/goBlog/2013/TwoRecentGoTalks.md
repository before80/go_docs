+++
title = "最近的两次 go 讲座"
weight = 21
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Two recent Go talks - 最近的两次 go 讲座

https://go.dev/blog/two-recent-go-talks

Andrew Gerrand
2 January 2013

## Introduction 简介

Late last year I wrote a couple of Go talks and presented them at [Strange Loop](http://thestrangeloop.com/), [Øredev](http://oredev.com/), and various other venues. The talks are designed to give insight into the practice of Go programming, each describing the construction of a real program and demonstrating the power and depth of the Go language and its libraries and tools.

去年年底，我写了几篇 Go 讲座，并在 Strange Loop、Øredev 和其他各种场合发表。这些讲座旨在让大家深入了解Go编程的实践，每个讲座都描述了一个真实程序的构建，并展示了Go语言及其库和工具的力量和深度。

The following videos are, in my opinion, the best recordings of these talks.

在我看来，以下视频是这些讲座的最佳记录。

## Go: a simple programming environment - Go：一个简单的编程环境

Go is a general-purpose language that bridges the gap between efficient statically typed languages and productive dynamic language. But it’s not just the language that makes Go special – Go has broad and consistent standard libraries and powerful but simple tools.

Go是一种通用语言，它在高效的静态类型语言和高效的动态语言之间架起了一座桥梁。但Go的特别之处不仅仅在于它的语言--Go拥有广泛而一致的标准库和强大而简单的工具。

This talk gives an introduction to Go, followed by a tour of some real programs that demonstrate the power, scope, and simplicity of the Go programming environment.

本讲座对Go进行了介绍，随后参观了一些真实的程序，展示了Go编程环境的力量、范围和简单性。

<iframe src="https://player.vimeo.com/video/53221558?badge=0" width="500" height="281" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

See the [slide deck](https://go.dev/talks/2012/simple.slide) (use the left and right arrows to navigate).

请看幻灯片（使用左右箭头进行导航）。

## Go: code that grows with grace - Go：优雅成长的代码

One of Go’s key design goals is code adaptability; that it should be easy to take a simple design and build upon it in a clean and natural way. In this talk I describe a simple "chat roulette" server that matches pairs of incoming TCP connections, and then use Go’s concurrency mechanisms, interfaces, and standard library to extend it with a web interface and other features. While the function of the program changes dramatically, Go’s flexibility preserves the original design as it grows.

Go的关键设计目标之一是代码的适应性；即应该很容易采用一个简单的设计，并以一种干净自然的方式在此基础上进行构建。在这次演讲中，我描述了一个简单的 "聊天轮盘 "服务器，它可以匹配进入的TCP连接对，然后使用Go的并发机制、接口和标准库，用网络接口和其他功能来扩展它。虽然程序的功能发生了巨大的变化，但Go的灵活性在其发展过程中保留了原有的设计。

<iframe src="https://player.vimeo.com/video/53221560?badge=0" width="500" height="281" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

See the [slide deck](https://go.dev/talks/2012/chat.slide) (use the left and right arrows to navigate).

请看幻灯片（使用左右箭头进行导航）。
