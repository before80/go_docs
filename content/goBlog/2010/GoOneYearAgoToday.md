+++
title = "go: 一年前的今天"
weight = 1
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go: one year ago today - go: 一年前的今天

> 原文：[https://go.dev/blog/1year](https://go.dev/blog/1year)

Andrew Gerrand
10 November 2010

2010年11月10日



​	2009年11月10日，我们启动了Go项目：一个注重简单和效率的开源编程语言。在接下来的一年中，Go项目本身和它的社区都发生了许多进展。

​	我们开始构建一种系统编程语言——通常可以用C或C++编写的程序类型——但我们惊讶地发现Go作为通用编程语言的效用。我们预计会吸引C、C++和Java程序员的兴趣，但是来自Python和JavaScript等动态类型语言的用户的兴趣激增，这是出乎意料的。Go语言的本地编译、静态类型、内存管理和轻量级语法似乎与广泛的编程社区产生了共鸣。

​	这个广泛的编程社区成长为了一个热情洋溢的Go编程者社区。我们的[邮件列表](http://groups.google.com/group/golang-nuts)有超过3,800名成员，每月约有1,500篇帖子。该项目有超过130位贡献者（提交代码或文档的人），自启动以来的2,800个提交中，将近三分之一是由核心团队之外的程序员贡献的。为了整理所有的代码，我们的[开发邮件列表](http://groups.google.com/group/golang-dev)上交换了近14,000封电子邮件。

​	这些数字反映了一项劳动的成果，这个劳动的成果在项目的代码库中得以体现。编译器已经得到了大幅改进，代码生成更快、更高效，修复了超过一百个已报告的错误，并支持越来越多的操作系统和架构。得益于一群热心的贡献者（其中一位成为了我们的第一位非Google提交者），Windows端口即将完成。ARM端口也取得了巨大进展，最近达到了通过所有测试的里程碑。

​	Go工具集已经扩展和改进。Go文档工具[godoc](https://go.dev/cmd/godoc/)现在支持其他源树的文档（您可以浏览和搜索自己的代码），并提供了一个"[代码漫步](https://go.dev/doc/codewalk/)"界面，用于呈现教程材料（还有许多其他改进）。新的包管理工具[Goinstall](https://go.dev/cmd/goinstall/)允许用户使用单个命令安装和更新外部包。Go格式化程序[Gofmt](https://go.dev/cmd/gofmt/)现在会尽可能地进行语法简化。[Goplay](https://go.dev/misc/goplay/)是一个基于Web的"编译时即时编译"工具，是一种方便的Go实验方式，特别是在您无法访问[Go Playground](https://go.dev/doc/play/)的时候。

​	标准库已增加了超过42,000行代码，包括20个新[包](https://go.dev/pkg/)。其中新增的包括[jpeg](https://go.dev/pkg/image/jpeg/)、[jsonrpc](https://go.dev/pkg/rpc/jsonrpc/)、[mime](https://go.dev/pkg/mime/)、[netchan](https://go.dev/pkg/netchan/)和[smtp](https://go.dev/pkg/smtp/)包，以及许多新的[加密](https://go.dev/pkg/crypto/)包。总的来说，随着我们对Go的语法习惯的理解加深，标准库一直在不断地完善和修订。

​	调试故事也变得更好了。最近对gc编译器的DWARF输出进行了改进，使得GNU调试器GDB在Go二进制文件中变得有用，我们正在积极努力使该调试信息更加完整。（详见[最近的博客文章](https://blog.golang.org/2010/11/debugging-go-code-status-report.html)。）

​	现在，与Go语言不同的语言编写的现有库比以往更容易连接。Go支持最新的[SWIG](http://www.swig.org/)版本2.0.1，使得连接C和C++代码更容易，并且我们的[cgo](https://go.dev/cmd/cgo/)工具已经得到了许多修复和改进。

​	[Gccgo](https://go.dev/doc/install/gccgo)，GNU C编译器的Go前端，作为一种并行的Go实现，已经跟上了gc编译器的步伐。现在它有了一个可工作的垃圾回收器，并已被GCC核心接受。我们现在正在努力使[gofrontend](http://code.google.com/p/gofrontend/)成为一个基于BSD许可证的Go编译器前端，与GCC完全分离。

​	除了Go项目本身外，Go开始被用于构建真正的软件。我们的[项目仪表板](http://godashboard.appspot.com/project)上列出了200多个Go程序和库，[Google Code](http://code.google.com/hosting/search?q=label:Go)和[Github](https://github.com/search?q=language:Go)上还有数百个。在我们的邮件列表和IRC频道上，您可以找到来自世界各地的编程人员，他们将Go用于他们的编程项目中。（请参见我们上个月的[客座博客文章](https://blog.golang.org/2010/10/real-go-projects-smarttwitter-and-webgo.html)以获取一个实际例子。）在Google内部，有几个团队选择使用Go来构建生产软件，并且我们收到了其他公司正在使用Go开发大型系统的报告。我们还与一些使用Go作为教学语言的教育工作者取得了联系。

​	语言本身也在增长和成熟。在过去的一年中，我们收到了许多功能请求。但Go是一种小型语言，我们努力确保任何新功能在简单性和实用性之间取得了正确的平衡。自启动以来，我们已经进行了许多语言更改，其中许多是由社区反馈驱动的。

- 在几乎所有情况下，分号现在是可选的。[规范]({{< ref "/langSpec/LexicalElements#Semicolons">}})
- 新的内置函数`copy`和`append`使得片段管理更加高效和简单。[规范]({{< ref "/langSpec/Built-inFunctions#appending-and-copying-slices">}})   
- 在制作子切片时可以省略上下界。这意味着`s[:]`是`s[0:len(s)]`的简写。[规范]({{< ref "/langSpec/types#slice-types">}}) 
- 新的内置函数recover作为一个错误处理机制，补充了panic和defer。[博客](../DeferPanicAandRecover)，[规范]({{< ref "/langSpec/Built-inFunctions#handling-panics">}}) 
- 新的复数类型（`complex`、`complex64`和`complex128`）简化了某些数学操作。[规范]({{< ref "/langSpec/Built-inFunctions#manipulating-complex-numbers">}})，[规范]({{< ref "/langSpec/LexicalElements#imaginary-literals">}})  
- 复合字面语法允许省略冗余的类型信息（例如，在指定二维数组时）。[发布.2010-10-27](https://go.dev/doc/devel/release.html#2010-10-27)，[规范]({{< ref "/langSpec/Expressions#composite-literals">}})  
- 现在规定了一种变量函数参数（`...T`）及其传播（`v...`）的通用语法。[规范]({{< ref "/langSpec/types#function-types">}})，[规范]({{< ref "/langSpec/Expressions#passing-arguments-to-parameters">}})，[发布.2010-09-29](https://go.dev/doc/devel/release.html#2010-09-29) 

​	Go肯定已经准备好进行生产使用，但仍有改进的余地。我们未来的重点是在高性能系统的背景下使Go程序更快、更高效。这意味着改进垃圾收集器、优化生成的代码和改进核心库。我们还在探索一些进一步的类型系统添加，以使通用编程更加容易。一年中发生了很多事情；这既令人激动，又令人满意。我们希望这一年会比去年更加丰硕。

​	如果您一直想（重新）使用Go，现在是一个很好的时机！查看[文档](https://go.dev/doc/docs.html)和[入门页面](https://go.dev/doc/install.html)以获取更多信息，或者在[*Go Playground*](https://go.dev/doc/play/)中进行实验。
