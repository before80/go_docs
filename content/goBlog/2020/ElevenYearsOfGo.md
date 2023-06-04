+++
title = "go 11 岁了"
weight = 3
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Eleven Years of Go - go 11 岁了

https://go.dev/blog/11years

Russ Cox, for the Go team
10 November 2020

Today we celebrate the eleventh birthday of the Go open source release. The parties we had for [Go turning 10](https://blog.golang.org/10years) seem like a distant memory. It’s been a tough year, but we’ve kept Go development moving forward and accumulated quite a few highlights.

今天我们庆祝Go开源版本的11岁生日。我们为Go十岁所举办的聚会似乎已经成为了遥远的记忆。这是艰难的一年，但我们让Go的开发不断向前推进，并积累了不少亮点。

In November, we launched [go.dev and pkg.go.dev](https://blog.golang.org/go.dev) shortly after Go’s 10th birthday.

11月，我们在Go的10岁生日后不久推出了go.dev和pkg.go.dev。

In February, the [Go 1.14 release](https://blog.golang.org/go1.14) delivered the first officially "production-ready" implementation of Go modules, along with many performance improvements, including [faster defers](https://go.dev/design/34481-opencoded-defers) and [non-cooperative goroutine preemption](https://go.dev/design/24543/conservative-inner-frame) to reduce scheduling and garbage collection latency.

2月，Go 1.14版本推出了第一个正式的 "生产就绪 "的Go模块实现，以及许多性能改进，包括更快的延迟和非合作的goroutine抢占，以减少调度和垃圾收集延迟。

In early March, we launched a [new API for protocol buffers](https://blog.golang.org/protobuf-apiv2), [google.golang.org/protobuf](https://pkg.go.dev/google.golang.org/protobuf), with much-improved support for protocol buffer reflection and custom messages.

3月初，我们推出了一个新的协议缓冲区的API，google.golang.org/protobuf，对协议缓冲区反射和自定义消息的支持有了很大的改进。

![img](ElevenYearsOfGo_img/gophermask.jpg)

When the pandemic hit, we decided to pause any public announcements or launches in the spring, recognizing that everyone’s attention rightly belonged elsewhere. But we kept working, and one of our team members joined the Apple/Google collaboration on [privacy-preserving exposure notifications](https://www.google.com/covid19/exposurenotifications/) to support contact tracing efforts all over the world. In May, that group launched the [reference backend server](https://github.com/google/exposure-notifications-server), written in Go.

当大流行病袭来时，我们决定暂停在春季的任何公开声明或发布，因为我们认识到每个人的注意力理应属于其他地方。但我们继续工作，我们的一个团队成员加入了苹果/谷歌合作的保护隐私的曝光通知，以支持世界各地的联系人追踪工作。5月，该小组推出了用Go编写的参考后端服务器。

We continued to improve [gopls](https://www.youtube.com/watch?v=EFJfdWzBHwE), which enables advanced [Go-aware support](https://github.com/golang/tools/blob/master/gopls/doc/user.md) in many editors. In June, the [VSCode Go extension officially joined the Go project](https://blog.golang.org/vscode-go) and is now maintained by the same developers who work on gopls.

我们继续改进gopls，它可以在许多编辑器中实现高级的Go-aware支持。6月，VSCode Go扩展正式加入了Go项目，现在由从事gopls工作的同一批开发人员维护。

Also in June, thanks to your feedback, we open-sourced [the code behind pkg.go.dev](https://blog.golang.org/pkgsite) as part of the Go project as well.

同样在6月，由于您们的反馈，我们将pkg.go.dev背后的代码也开源为Go项目的一部分。

Later in June, we [released the latest design draft for generics](https://blog.golang.org/generics-next-step), along with a prototype tool and [generics playground](https://go2goplay.golang.org/).

6月下旬，我们发布了泛型的最新设计草案，以及一个原型工具和泛型操场。

In July, we published and discussed three new design drafts for future changes: [new `//go:build` lines for file selection](https://go.dev/design/draft-gobuild), [file system interfaces](https://go.dev/design/draft-iofs), and [build-time file embedding](https://go.dev/design/draft-embed). (We’ll see all of those in 2021, as noted below.)

7月，我们发布并讨论了三个新的设计草案，用于未来的修改：用于文件选择的新的//go:build行，文件系统接口，以及构建时文件嵌入。(如下文所述，我们将在2021年看到所有这些东西）。

In August, the [Go 1.15 release](https://blog.golang.org/go1.15) delivered mainly optimizations and bug fixes rather than new features. The most significant was the start of a rewrite of the linker, making it run 20% faster and use 30% less memory on average for large builds.

在8月，Go 1.15版本主要提供了优化和错误修复，而不是新功能。最重要的是开始重写链接器，使其运行速度提高了20%，在大型构建中平均使用的内存减少了30%。

Last month, we ran our [annual Go user survey](https://blog.golang.org/survey2020). We will post results on the blog once we’ve analyzed them.

上个月，我们进行了年度Go用户调查。一旦我们分析了结果，我们会在博客上公布。

The Go community has adapted to "virtual-first" along with everyone else, and we saw many virtual meetups and over a dozen virtual Go conferences this year. Last week, the Go team hosted [Go day at Google Open Source Live](https://opensourcelive.withgoogle.com/events/go) (videos at the link).

Go社区已经和大家一起适应了 "虚拟优先"，今年我们看到了许多虚拟聚会和十多个虚拟Go会议。上周，Go团队在Google Open Source Live举办了Go日（视频见链接）。

## Going Forward 向前迈进

We’re also incredibly excited about what’s in store for Go’s 12th year. Most immediately, this week Go team members will be presenting eight events at [GopherCon 2020](https://www.gophercon.com/). Mark your calendars!

我们也对Go的第12年的发展感到无比的兴奋。最直接的是，本周Go团队成员将在GopherCon 2020上展示八项活动。请在您的日历上做个记号。

- "Typing [Generic] Go", a talk by Robert Griesemer, 
  [Nov 11, 10:00 AM (US Eastern)](https://www.gophercon.com/agenda/session/233094); [Q&A at 10:30 AM](https://www.gophercon.com/agenda/session/417935).
- "What to Expect When You’re NOT Expecting", a live taping of the Go time podcast with a panel of expert debuggers, including Hana Kim,
  [Nov 11 12:00 PM](https://www.gophercon.com/agenda/session/2334490).
- "Evolving the Go Memory Manager’s RAM and CPU Efficiency", a talk by Michael Knyszek,
  [Nov 11 1:00 PM](https://www.gophercon.com/agenda/session/233086); [Q&A at 1:50 PM](https://www.gophercon.com/agenda/session/417940).
- "Implementing Faster Defers", a talk by Dan Scales,
  [Nov 11 5:10 PM](https://www.gophercon.com/agenda/session/233397); [Q&A at 5:40 PM](https://www.gophercon.com/agenda/session/417941).
- "Go Team - Ask Me Anything", a live Q&A with Julie Qiu, Rebecca Stambler, Russ Cox, Sameer Ajmani, and Van Riper,
  [Nov 12 3:00 PM](https://www.gophercon.com/agenda/session/420539).
- "Pardon the Interruption: Loop Preemption in Go 1.14", a talk by Austin Clements,
  [Nov 12 4:45 PM](https://www.gophercon.com/agenda/session/233441); [Q&A at 5:15 PM](https://www.gophercon.com/agenda/session/417943).
- "Working with Errors", a talk by Jonathan Amsterdam,
  [Nov 13 1:00 PM](https://www.gophercon.com/agenda/session/233432); [Q&A at 1:50 PM](https://www.gophercon.com/agenda/session/417945).
- "Crossing the Chasm for Go: Two Million Users and Growing", a talk by Carmen Andoh,
  [Nov 13 5:55 PM](https://www.gophercon.com/agenda/session/233426).
- 11月11日上午10:00（美国东部时间）；上午10:30进行问答。
  "What to Expect When You're NOT Expecting"，Go time播客的现场录制，包括Hana Kim在内的专家调试员参加。
  11月11日 12:00 PM.
  "进化Go内存管理器的内存和CPU效率"，由Michael Knyszek演讲。
  11月11日下午1:00；下午1:50的问答。
  "实现更快的延迟"，由Dan Scales演讲。
  11月11日下午5:10；下午5:40进行问答。
  "Go Team - Ask Me Anything"，由Julie Qiu、Rebecca Stambler、Russ Cox、Sameer Ajmani和Van Riper现场问答。
  11月12日下午3:00。
  "请原谅我的打断。Go 1.14中的循环抢占"，Austin Clements的演讲。
  11月12日下午4:45；下午5:15的问答。
  "与错误打交道"，由Jonathan Amsterdam主讲。
  11月13日下午1:00；下午1:50时有问答。
  "跨越Go的鸿沟。200万用户和增长"，Carmen Andoh的演讲。
  11月13日下午5:55。

## Go Releases Go的发布

In February, the Go 1.16 release will include the new [file system interfaces](https://tip.golang.org/pkg/io/fs/) and [build-time file embedding](https://tip.golang.org/pkg/embed/). It will complete the linker rewrite, bringing additional performance improvements. And it will include support for the new Apple Silicon (`GOARCH=arm64`) Macs.

在2月份，Go 1.16版本将包括新的文件系统接口和构建时文件嵌入。它将完成链接器的重写，带来额外的性能改进。它还将包括对新的苹果硅（GOARCH=arm64）Mac的支持。

In August, the Go 1.17 release will no doubt bring more features and improvements, although it’s far enough out that the exact details remain up in the air. It will include a new register-based calling convention for x86-64 (without breaking existing assembly!), which will make programs faster across the board. (Other architectures will follow in later releases.) One nice feature that will definitely be included is the [new `//go:build` lines](https://go.dev/design/draft-gobuild), which are far less error-prone than the [current `//` `+build` lines](https://go.dev/cmd/go/#hdr-Build_constraints). Another highly anticipated feature we hope will be ready for beta testing next year is [support for fuzzing in the `go test` command](https://go.dev/design/draft-fuzzing).

8月，Go 1.17版本无疑将带来更多的功能和改进，尽管它还很遥远，确切的细节仍然是未知的。它将包括一个针对x86-64的新的基于寄存器的调用约定（不破坏现有的汇编！），这将使整个程序更快。(其他架构将在以后的版本中跟进。)一个肯定会包括的好功能是新的//go:build行，它比目前的//+build行更少出错。另一个备受期待的功能，我们希望能在明年准备好进行测试，那就是支持go test命令中的fuzzing。

## Go Modules Go模块

Over the next year, we will continue to work on developing support for Go modules and integrating them well into the entire Go ecosystem. Go 1.16 will include our smoothest Go modules experience yet. One preliminary result from our recent survey is that 96% of users have now adopted Go modules (up from 90% a year ago).

在接下来的一年里，我们将继续致力于开发对Go模块的支持，并将其很好地整合到整个Go生态系统中。Go 1.16将包括我们迄今为止最顺畅的Go模块体验。我们最近调查的一个初步结果是，96%的用户现在已经采用了Go模块（一年前为90%）。

We will also finally wind down support for GOPATH-based development: any programs using dependencies other than the standard library will need a `go.mod`. (If you haven’t switched to modules yet, see the [GOPATH wiki page](https://go.dev/wiki/GOPATH) for details about this final step in the journey from GOPATH to modules.)

我们也将最终结束对基于GOPATH的开发的支持：任何使用标准库以外的依赖关系的程序都需要一个go.mod。(如果您还没有切换到模块，请参阅GOPATH维基页面，了解从GOPATH到模块的最后一步的细节）。

From the start, the [goal for Go modules](https://research.swtch.com/vgo-intro) has been "to add the concept of package versions to the working vocabulary of both Go developers and our tools," to enable deep support for modules and versions throughout the Go ecosystem. The [Go module mirror, checksum database, and index](https://blog.golang.org/modules2019) were made possible by this ecosystem-wide understanding of what a package version is. Over the next year, we will see rich module support added to more tools and systems. For example, we plan to investigate new tooling to help module authors publish new versions (`go release`) as well as to help module consumers update their code to migrate away from deprecated APIs (a new `go fix`).

从一开始，Go模块的目标就是 "在Go开发者和我们的工具的工作词汇中加入软件包版本的概念"，以便在整个Go生态系统中实现对模块和版本的深度支持。Go模块镜像、校验和数据库和索引都是由整个生态系统对软件包版本的理解所促成的。在接下来的一年里，我们将看到更多的工具和系统加入丰富的模块支持。例如，我们计划研究新的工具，以帮助模块作者发布新的版本（go release），以及帮助模块消费者更新他们的代码，以便从废弃的API中迁移出来（一个新的go fix）。

As a larger example, [we created gopls](https://github.com/golang/tools/blob/master/gopls/README.md) to reduce many tools used by editors for Go support, none of which supported modules, down to a single one that did. Over the next year, we’ll be ready to make the VSCode Go extension use `gopls` by default, for an excellent module experience out of the box, and we’ll release gopls 1.0. Of course, one of the best things about gopls is that it is editor-neutral: any editor that understands the [language server protocol](https://langserver.org/) can use it.

作为一个更大的例子，我们创建了gopls来减少编辑者用于Go支持的许多工具，这些工具都不支持模块，只有一个工具支持。明年，我们将准备让VSCode Go扩展默认使用gopls，以获得开箱即用的出色的模块体验，并且我们将发布gopls 1.0。当然，gopls最好的一点是，它是编辑器中立的：任何理解语言服务器协议的编辑器都可以使用它。

Another important use of version information is tracking whether any package in a build has a known vulnerability. Over the next year, we plan to develop a database of known vulnerabilities as well as tools to check your programs against that database.

版本信息的另一个重要用途是跟踪构建中的任何软件包是否有已知的漏洞。在接下来的一年里，我们计划开发一个已知漏洞的数据库，以及根据该数据库检查您的程序的工具。

The Go package discovery site [pkg.go.dev](https://pkg.go.dev/) is another example of a version-aware system enabled by Go modules. We’ve been focused on getting the core functionality and user experience right, including a [redesign launching today](https://blog.golang.org/pkgsite-redesign). Over the next year, we will be unifying godoc.org into pkg.go.dev. We will also be expanding the version timeline for each package, showing important changes in each version, known vulnerabilities, and more, following the overall goal of surfacing what you need to make [informed decisions about adding dependencies](https://research.swtch.com/deps).

Go软件包发现网站pkg.go.dev是另一个由Go模块启用的版本意识系统的例子。我们一直专注于将核心功能和用户体验做好，包括今天推出的重新设计。在接下来的一年里，我们将把godoc.org统一到pkg.go.dev。我们还将扩大每个软件包的版本时间线，显示每个版本的重要变化、已知的漏洞等等，其总体目标是展示您在添加依赖关系时需要做出的知情决定。

We’re excited to see this journey from GOPATH to Go modules nearing completion and all the excellent dependency-aware tools that Go modules are enabling.

我们很高兴看到从GOPATH到Go模块的旅程即将完成，以及Go模块所带来的所有优秀的依赖意识工具。

## Generics 泛型

The next feature on everyone’s minds is of course generics. As we mentioned above, we published the [latest design draft for generics](https://blog.golang.org/generics-next-step) back in June. Since then, we’ve continued to refine rough edges and have turned our attention to the details of implementing a production-ready version. We will be working on that throughout 2021, with a goal of having something for people to try out by the end of the year, perhaps a part of the Go 1.18 betas.

每个人都在考虑的下一个功能当然是泛型。正如我们上面提到的，我们早在6月份就发布了泛型的最新设计草案。从那时起，我们继续完善粗糙的边缘，并将注意力转移到实现生产就绪版本的细节上。我们将在整个2021年进行这项工作，目标是在年底前为人们提供一些尝试，也许是Go 1.18测试版的一部分。

## Thank You! 谢谢您!

Go is far more than just us on the Go team at Google. We are indebted to the contributors who work with us with the Go releases and tools. Beyond that, Go only succeeds because of all of you who work in and contribute to Go’s thriving ecosystem. It has been a difficult year in the world outside Go. More than ever, we appreciate you taking the time to join us and help make Go such a success. Thank you. We hope you are all staying safe and wish you all the best.

Go不仅仅是我们谷歌的Go团队的事情。我们要感谢那些与我们一起发布Go版本和工具的贡献者们。除此之外，Go的成功还在于所有在Go的繁荣生态系统中工作并为之做出贡献的人们。在Go之外的世界，这是艰难的一年。我们比以往任何时候都更感谢您抽出时间加入我们，帮助Go取得如此的成功。谢谢您们。我们希望您们都能保持安全，并祝您们一切顺利。
