+++
title = "go 在ARM和其他"
weight = 1
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go on ARM and Beyond - go 在ARM和其他

> 原文：[https://go.dev/blog/ports](https://go.dev/blog/ports)

Russ Cox
17 December 2020

The industry is abuzz about non-x86 processors recently, so we thought it would be worth a brief post about Go’s support for them.

最近，业界对非x86处理器议论纷纷，所以我们认为值得发表一篇关于Go对其支持的简短文章。

It has always been important to us for Go to be portable, not overfitting to any particular operating system or architecture. The [initial open source release of Go](https://opensource.googleblog.com/2009/11/hey-ho-lets-go.html) included support for two operating systems (Linux and Mac OS X) and three architectures (64-bit x86, 32-bit x86, and 32-bit ARM).

对我们来说，Go的可移植性一直很重要，它不会过度适应任何特定的操作系统或架构。Go的最初开源版本包括对两种操作系统（Linux和Mac OS X）和三种架构（64位x86、32位x86和32位ARM）的支持。

Over the years, we’ve added support for many more operating systems and architecture combinations:

多年来，我们已经增加了对更多操作系统和架构组合的支持：

- Go 1 (March 2012) supported the original systems as well as FreeBSD, NetBSD, and OpenBSD on 64-bit and 32-bit x86, and Plan 9 on 32-bit x86.Go 1（2012年3月）支持原始系统以及64位和32位x86的FreeBSD、NetBSD和OpenBSD，以及32位x86的Plan 9。
- Go 1.3 (June 2014) added support for Solaris on 64-bit x86.Go 1.3（2014年6月）增加了对64位x86的Solaris的支持。
- Go 1.4 (December 2014) added support for Android on 32-bit ARM and Plan 9 on 64-bit x86.Go 1.4（2014年12月）增加了对32位ARM的Android和64位x86的Plan 9的支持。
- Go 1.5 (August 2015) added support for Linux on 64-bit ARM and 64-bit PowerPC, as well as iOS on 32-bit and 64-bit ARM.Go 1.5（2015年8月）增加了对64位ARM和64位PowerPC上的Linux的支持，以及32位和64位ARM上的iOS。
- Go 1.6 (February 2016) added support for Linux on 64-bit MIPS, as well as Android on 32-bit x86. It also added an official binary download for Linux on 32-bit ARM, primarily for Raspberry Pi systems.Go 1.6（2016年2月）增加了对64位MIPS上的Linux，以及32位x86上的Android的支持。它还为32位ARM上的Linux增加了官方二进制下载，主要用于Raspberry Pi系统。
- Go 1.7 (August 2016) added support for Linux on z Systems (S390x) and Plan 9 on 32-bit ARM.Go 1.7（2016年8月）增加了对z系统（S390x）上的Linux和32位ARM上的Plan 9的支持。
- Go 1.8 (February 2017) added support for Linux on 32-bit MIPS, and it added official binary downloads for Linux on 64-bit PowerPC and z Systems.Go 1.8（2017年2月）增加了对32位MIPS上的Linux的支持，并且它增加了对64位PowerPC和z系统上的Linux的官方二进制下载。
- Go 1.9 (August 2017) added official binary downloads for Linux on 64-bit ARM.Go 1.9（2017年8月）为64位ARM上的Linux增加了官方二进制下载。
- Go 1.12 (February 2018) added support for Windows 10 IoT Core on 32-bit ARM, such as the Raspberry Pi 3. It also added support for AIX on 64-bit PowerPC.Go 1.12（2018年2月）增加了对32位ARM上的Windows 10 IoT Core的支持，如Raspberry Pi 3。它还增加了对64位PowerPC上AIX的支持。
- Go 1.14 (February 2019) added support for Linux on 64-bit RISC-V.Go 1.14（2019年2月）增加了对64位RISC-V上的Linux的支持。

Although the x86-64 port got most of the attention in the early days of Go, today all our target architectures are well supported by our [SSA-based compiler back end](https://www.youtube.com/watch?v=uTMvKVma5ms) and produce excellent code. We’ve been helped along the way by many contributors, including engineers from Amazon, ARM, Atos, IBM, Intel, and MIPS.

虽然x86-64端口在Go的早期得到了大部分的关注，但今天我们所有的目标架构都得到了我们基于SSA的编译器后端良好的支持，并产生了优秀的代码。一路走来，我们得到了许多贡献者的帮助，包括来自Amazon、ARM、Atos、IBM、Intel和MIPS的工程师。

Go supports cross-compiling for all these systems out of the box with minimal effort. For example, to build an app for 32-bit x86-based Windows from a 64-bit Linux system:

Go支持所有这些系统的交叉编译，开箱即用。例如，要从64位的Linux系统为32位的x86系统的Windows建立一个应用程序：

```
GOARCH=386 GOOS=windows go build myapp  # writes myapp.exe
```

In the past year, several major vendors have made announcements of new ARM64 hardware for servers, laptops and developer machines. Go was well-positioned for this. For years now, Go has been powering Docker, Kubernetes, and the rest of the Go ecosystem on ARM64 Linux servers, as well as mobile apps on ARM64 Android and iOS devices.

在过去的一年中，几个主要的供应商已经宣布了用于服务器、笔记本电脑和开发者机器的新的ARM64硬件。Go在这方面处于有利的地位。多年来，Go一直在ARM64 Linux服务器上为Docker、Kubernetes和Go生态系统的其他部分提供动力，并在ARM64 Android和iOS设备上提供移动应用程序。

Since Apple’s announcement of the Mac transitioning to Apple Silicon this summer, Apple and Google have been working together to ensure that Go and the broader Go ecosystem work well on them, both running Go x86 binaries under Rosetta 2 and running native Go ARM64 binaries. Earlier this week, we released the first Go 1.16 beta, which includes native support for Macs using the M1 chip. You can download and try the Go 1.16 beta for M1 Macs and all your other systems on [the Go download page](https://go.dev/dl/#go1.16beta1). (Of course, this is a beta release and, like all betas, it certainly has bugs we don’t know about. If you run into any problems, please report them at [golang.org/issue/new](https://go.dev/issue/new).)

自从今年夏天苹果宣布Mac过渡到Apple Silicon以来，苹果和谷歌一直在合作，以确保Go和更广泛的Go生态系统在其上运行良好，既在Rosetta 2下运行Go x86二进制文件，又运行原生Go ARM64二进制文件。本周早些时候，我们发布了第一个Go 1.16测试版，其中包括对使用M1芯片的Mac的本地支持。您可以在Go下载页面上下载并试用适用于M1 Macs和所有其他系统的Go 1.16测试版。(当然，这是一个测试版，像所有的测试版一样，它肯定有我们不知道的错误）。如果您遇到任何问题，请在golang.org/issue/new上报告）。

It’s always nice to use the same CPU architecture for local development as in production, to remove one variation between the two environments. If you deploy to ARM64 production servers, Go makes it easy to develop on ARM64 Linux and Mac systems too. But of course, it’s still as easy as ever to work on one system and cross-compile for deployment to another, whether you’re working on an x86 system and deploying to ARM, working on Windows and deploying to Linux, or some other combination.

在本地开发中使用与生产中相同的CPU架构总是好的，这样可以消除两种环境之间的一个差异。如果您部署到ARM64的生产服务器上，Go也可以让您在ARM64的Linux和Mac系统上轻松开发。但当然，在一个系统上工作并交叉编译以部署到另一个系统上仍然和以前一样容易，无论您是在x86系统上工作并部署到ARM，在Windows上工作并部署到Linux，还是其他一些组合。

The next target we’d like to add support for is ARM64 Windows 10 systems. If you have expertise and would like to help, we’re coordinating work on [golang.org/issue/36439](https://github.com/golang/go/issues/36439).

我们希望增加对ARM64 Windows 10系统的支持，这是我们的下一个目标。如果您有专业知识并愿意提供帮助，我们正在golang.org/issue/36439上协调工作。
