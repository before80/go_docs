+++
title = "go 7 岁了"
weight = 3
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Seven years of Go - go 7 岁了

https://go.dev/blog/7years

The Go Team
10 November 2016

![img](SevenYearsOfGo_img/gopherbelly300.jpg){ align=left }

Today marks seven years since we open-sourced our preliminary sketch of Go. With the help of the open source community, including more than a thousand individual contributors to the Go source repositories, Go has matured into a language used all over the world.

今天是我们将Go的初步草图开源的七周年。在开源社区的帮助下，包括一千多名Go源码库的个人贡献者，Go已经成熟为一种在世界各地使用的语言。

The most significant user-facing changes to Go over the past year are the addition of built-in support for [HTTP/2](https://www.youtube.com/watch?v=FARQMJndUn0#t=0m0s) in [Go 1.6](https://go.dev/doc/go1.6) and the integration of the [context package](https://blog.golang.org/context) into the standard library in [Go 1.7](https://go.dev/doc/go1.7). But we’ve been making many less visible improvements. Go 1.7 changed the x86-64 compiler to use a new SSA-based back end, improving the performance of most Go programs by 10–20%. For Go 1.8, planned for release next February, we have changed the compilers for the other architectures to use the new back end too. We’ve also added new ports, to Android on 32-bit x86, Linux on 64-bit MIPS, and Linux on IBM z Systems. And we’ve developed new garbage-collection techniques that reduce typical "stop the world" pauses to [under 100 microseconds](https://go.dev/design/17503-eliminate-rescan). (Contrast that with Go 1.5’s big news of [10 milliseconds or less](https://blog.golang.org/go15gc).)

在过去的一年中，Go最重要的面向用户的变化是在Go 1.6中增加了对HTTP/2的内置支持，以及在Go 1.7中将上下文包整合到标准库中。但我们一直在进行许多不太明显的改进。Go 1.7改变了x86-64编译器，使用新的基于SSA的后端，使大多数Go程序的性能提高了10-20%。在计划于明年二月发布的 Go 1.8 中，我们已经将其他架构的编译器也改为使用新的后端。我们还增加了新的端口，用于32位x86的Android，64位MIPS的Linux，以及IBM z系统的Linux。我们还开发了新的垃圾收集技术，将典型的 "停止世界 "的暂停时间减少到100微秒以下。(与Go 1.5的10毫秒或更少的大新闻形成鲜明对比）。

This year kicked off with a global Go hackathon, the [Gopher Gala](https://blog.golang.org/gophergala), in January. Then there were [Go conferences](https://go.dev/wiki/Conferences) in India and Dubai in February, China and Japan in April, San Francisco in May, Denver in July, London in August, Paris last month, and Brazil this past weekend. And GothamGo in New York is next week. This year also saw more than 30 new [Go user groups](https://go.dev/wiki/GoUserGroups), eight new [Women Who Go](http://www.womenwhogo.org/) chapters, and four [GoBridge](https://golangbridge.org/) workshops around the world.

今年1月的全球Go黑客马拉松（Gopher Gala）拉开了序幕。然后，2月在印度和迪拜，4月在中国和日本，5月在旧金山，7月在丹佛，8月在伦敦，上个月在巴黎，以及上周末在巴西都有Go会议。而纽约的GothamGo将在下周举行。今年还出现了30多个新的Go用户群，8个新的 "Go女性 "分会，以及世界各地的4个GoBridge研讨会。

We continue to be overwhelmed by and grateful for the enthusiasm and support of the Go community. Whether you participate by contributing changes, reporting bugs, sharing your expertise in design discussions, writing blog posts or books, running meetups, helping others learn or improve, open sourcing Go packages you wrote, or just being part of the Go community, the Go team thanks you for your help, your time, and your energy. Go would not be the success it is today without you.

我们继续被Go社区的热情和支持所折服，并对其表示感谢。无论您是通过贡献修改、报告错误、在设计讨论中分享您的专业知识、撰写博文或书籍、举办聚会、帮助他人学习或改进、开源您编写的Go软件包，还是仅仅作为Go社区的一部分，Go团队都感谢您的帮助、时间和精力。没有您们，Go就不会有今天的成功。

Thank you, and here’s to another year of fun and success with Go!

谢谢您们，祝愿Go在新的一年里获得更多的乐趣和成功。
