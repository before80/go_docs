+++
title = "go 1.1发布了"
weight = 13
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.1 is released - go 1.1发布了

https://go.dev/blog/go1.1

Andrew Gerrand
13 May 2013

It is our great pleasure to announce the release of Go 1.1.

我们非常高兴地宣布Go 1.1的发布。

![img](Go1_1IsReleased_img/gopherbiplane5.jpg)

In March last year we released Go 1.0, and since then we have released three minor "point releases". The point releases were made to fix only critical issues, so the Go 1.0.3 you use today is still, in essence, the Go 1.0 we released in March 2012.

去年3月，我们发布了Go 1.0，此后我们又发布了三个小的 "点 "版本。这些点发布只是为了修复关键问题，所以您今天使用的Go 1.0.3本质上仍然是我们在2012年3月发布的Go 1.0。

Go 1.1 includes many improvements over 1.0.

与1.0相比，Go 1.1包括许多改进。

The most significant improvements are performance-related. We have made optimizations in the compiler and linker, garbage collector, goroutine scheduler, map implementation, and parts of the standard library. It is likely that your Go code will run noticeably faster when built with Go 1.1.

最重要的改进是与性能有关的。我们在编译器和链接器、垃圾收集器、goroutine调度器、map实现以及部分标准库中进行了优化。当使用 Go 1.1 构建时，您的 Go 代码很可能会明显加快运行速度。

There are some minor changes to the language itself, two of which are worth singling out here: the [changes to return requirements](https://go.dev/doc/go1.1#return) will lead to more succinct and correct programs, and the introduction of [method values](https://go.dev/doc/go1.1#method_values) provides an expressive way to bind a method to its receiver as a function value.

语言本身也有一些细微的变化，其中有两点值得在此特别指出：对返回要求的改变将导致更简洁和正确的程序，而方法值的引入提供了一种表达方式，将一个方法与其接收器绑定为一个函数值。

Concurrent programming is safer in Go 1.1 with the addition of a race detector for finding memory synchronization errors in your programs. We will discuss the race detector more in an upcoming article, but for now [the manual](https://go.dev/doc/articles/race_detector.html) is a great place to get started.

Go1.1中的并发编程更加安全，增加了一个竞赛检测器，用于查找程序中的内存同步错误。我们将在接下来的文章中进一步讨论竞赛检测器，但现在手册是一个很好的开始。

The tools and standard library have been improved and expanded. You can read the full story in the [release notes](https://go.dev/doc/go1.1).

工具和标准库也得到了改进和扩展。您可以在发布说明中阅读完整的故事。

As per our [compatibility guidelines](https://go.dev/doc/go1compat.html), Go 1.1 remains compatible with Go 1.0 and we recommend all Go users upgrade to the new release.

根据我们的兼容性准则，Go 1.1仍然与Go 1.0兼容，我们建议所有Go用户升级到新版本。

All this would not have been possible without the help of our contributors from the open source community. Since Go 1.0, the core received more than 2600 commits from 161 people outside Google. Thank you everyone for your time and effort. In particular, we would like to thank Shenghou Ma, Rémy Oudompheng, Dave Cheney, Mikio Hara, Alex Brainman, Jan Ziak, and Daniel Morsing for their outstanding contributions.

如果没有我们来自开源社区的贡献者的帮助，这一切是不可能实现的。自Go 1.0以来，该核心收到了来自Google以外161人的2600多个提交。感谢大家的时间和努力。特别是，我们要感谢马胜侯、Rémy Oudompheng、Dave Cheney、Mikio Hara、Alex Brainman、Jan Ziak和Daniel Morsing的杰出贡献。

To grab the new release, follow the usual [installation instructions](https://go.dev/doc/install). Happy hacking!

要抓取新版本，请按照通常的安装说明进行。快乐的黑客行为!

*Thanks to Renée French for the gopher!*

感谢Renée French的打地鼠!
