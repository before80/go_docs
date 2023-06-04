+++
title = "Heroku 上的 go"
weight = 22
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go at Heroku - Heroku 上的 go

https://go.dev/blog/heroku

Keith Rarick and Blake Mizerany
21 April 2011

*This week’s blog post is written by* [*Keith Rarick*](http://xph.us/) *and* [*Blake Mizerany*](http://itsbonus.heroku.com/), *systems engineers at* [Heroku](http://www.heroku.com/). *In their own words, they "eat, drink, and sleep distributed systems." Here they discuss their experiences using Go.*

本周的博文是由Heroku的系统工程师Keith Rarick和Blake Mizerany撰写的。用他们自己的话说，他们 "吃喝拉撒睡都是分布式系统"。在这里，他们讨论了他们使用Go的经验。

A big problem that comes with building distributed systems is the coordination of physical servers. Each server needs to know various facts about the system as a whole. This critical data includes locks, configuration data, and so on, and it must be consistent and available even during data store failures, so we need a data store with solid consistency guarantees. Our solution to this problem is [Doozer](http://xph.us/2011/04/13/introducing-doozer.html), a new, consistent, highly-available data store written in Go.

构建分布式系统的一个大问题是物理服务器的协调问题。每台服务器都需要知道关于整个系统的各种事实。这些关键数据包括锁、配置数据等，即使在数据存储发生故障时也必须保持一致和可用，因此我们需要一个具有坚实一致性保证的数据存储。我们对这个问题的解决方案是Doozer，一个用Go编写的新的、一致的、高可用的数据存储。

At Doozer’s core is [Paxos](http://en.wikipedia.org/wiki/Paxos_(computer_science)), a family of protocols for solving consensus in an unreliable network of unreliable nodes. While Paxos is essential to running a fault-tolerant system, it is notorious for being difficult to implement. Even example implementations that can be found online are complex and hard to follow, despite being simplified for educational purposes. Existing production systems have a reputation for being worse.

Doozer的核心是Paxos，这是一个用于解决不可靠节点的不可靠网络中的共识问题的协议系列。虽然Paxos对运行容错系统至关重要，但它因难以实现而臭名昭著。即使是在网上可以找到的实例实现也很复杂，很难遵循，尽管是为了教育目的而简化。现有的生产系统也以更糟糕而闻名。

Fortunately, Go’s concurrency primitives made the task much easier. Paxos is defined in terms of independent, concurrent processes that communicate via passing messages. In Doozer, these processes are implemented as goroutines, and their communications as channel operations. In the same way that garbage collectors improve upon malloc and free, we found that [goroutines and channels](https://blog.golang.org/2010/07/share-memory-by-communicating.html) improve upon the lock-based approach to concurrency. These tools let us avoid complex bookkeeping and stay focused on the problem at hand. We are still amazed at how few lines of code it took to achieve something renowned for being difficult.

幸运的是，Go的并发原语使这项任务变得更加容易。Paxos是以独立的、并发的进程来定义的，这些进程通过传递消息进行通信。在Doozer中，这些进程被实现为goroutines，而它们的通信则是通道操作。就像垃圾收集器改进了malloc和free一样，我们发现goroutines和通道改进了基于锁的并发方法。这些工具让我们避免了复杂的簿记工作，而将注意力集中在手头的问题上。我们仍然惊讶于只用了几行代码就实现了以困难著称的事情。

The standard packages in Go were another big win for Doozer. The Go team is very pragmatic about what goes into them. For instance, a package we quickly found useful was [websocket](https://go.dev/pkg/websocket/). Once we had a working data store, we needed an easy way to introspect it and visualize activity. Using the websocket package, Keith was able to add the web viewer on his train ride home and without requiring external dependencies. This is a real testament to how well Go mixes systems and application programming.

Go中的标准包是Doozer的另一大胜利。Go团队对进入它们的内容非常务实。例如，我们很快就发现websocket是一个有用的包。一旦我们有了一个工作的数据存储，我们就需要一个简单的方法来反省它并将活动可视化。使用websocket包，Keith能够在回家的火车上添加网络查看器，而且不需要外部依赖性。这真正证明了Go将系统和应用编程结合得多么好。

One of our favorite productivity gains was provided by Go’s source formatter: [gofmt](https://go.dev/cmd/gofmt/). We never argued over where to put a curly-brace, tabs vs. spaces, or if we should align assignments. We simply agreed that the buck stopped at the default output from gofmt.

我们最喜欢的生产力提升之一是由Go的源码格式化器提供的：gofmt。我们从来没有争论过要把大括号放在哪里，制表符与空格的关系，或者我们是否应该对齐赋值。我们只是简单地同意在gofmt的默认输出中停止。

Deploying Doozer was satisfyingly simple. Go builds statically linked binaries which means Doozer has no external dependencies; it’s a single file that can be copied to any machine and immediately launched to join a cluster of running Doozers.

部署Doozer是非常简单的，令人满意。Go构建静态链接的二进制文件，这意味着Doozer没有外部依赖性；它是一个单一的文件，可以被复制到任何一台机器上，并立即启动，加入运行中的Doozer集群。

Finally, Go’s maniacal focus on simplicity and orthogonality aligns with our view of software engineering. Like the Go team, we are pragmatic about what features go into Doozer. We sweat the details, preferring to change an existing feature instead of introducing a new one. In this sense, Go is a perfect match for Doozer.

最后，Go对简单性和正交性的狂热关注与我们对软件工程的看法是一致的。和Go团队一样，我们对Doozer的功能也是很务实的。我们关注细节，宁愿改变现有的功能而不是引入新的功能。从这个意义上说，Go是与Doozer完美匹配的。

We already have future projects in mind for Go. Doozer is just the start of much bigger system.

我们已经有了关于Go的未来项目。Doozer只是一个更大的系统的开始。
