+++
title = "go GC：优先考虑低延迟和简单性"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go GC: Prioritizing low latency and simplicity - go GC：优先考虑低延迟和简单性

https://go.dev/blog/go15gc

Richard Hudson
31 August 2015

## The Setup 设置

Go is building a garbage collector (GC) not only for 2015 but for 2025 and beyond: A GC that supports today’s software development and scales along with new software and hardware throughout the next decade. Such a future has no place for stop-the-world GC pauses, which have been an impediment to broader uses of safe and secure languages such as Go.

Go不仅为2015年，而且为2025年及以后建立了一个垃圾收集器（GC）。一个支持今天的软件开发，并在未来十年与新的软件和硬件一起扩展的GC。在这样的未来，没有停止的地方，因为GC的暂停一直阻碍着像Go这样的安全语言的广泛使用。

Go 1.5, the first glimpse of this future, achieves GC latencies well below the 10 millisecond goal we set a year ago. We presented some impressive numbers in [a talk at Gophercon](https://go.dev/talks/2015/go-gc.pdf). The latency improvements have generated a lot of attention; Robin Verlangen’s blog post [*Billions of requests per day meet Go 1.5*](https://medium.com/@robin.verlangen/billions-of-request-per-day-meet-go-1-5-362bfefa0911) validates our direction with end to end results. We also particularly enjoyed [Alan Shreve’s production server graphs](https://twitter.com/inconshreveable/status/620650786662555648) and his “Holy 85% reduction” comment.

Go 1.5是这个未来的第一道曙光，它实现了远低于我们一年前设定的10毫秒目标的GC延迟。我们在Gophercon的演讲中展示了一些令人印象深刻的数字。延迟的改善引起了很多人的关注；Robin Verlangen的博文《每天数十亿的请求满足Go 1.5》用端到端的结果验证了我们的方向。我们也特别喜欢Alan Shreve的生产服务器图和他的 "神圣的减少85%"的评论。

Today 16 gigabytes of RAM costs $100 and CPUs come with many cores, each with multiple hardware threads. In a decade this hardware will seem quaint but the software being built in Go today will need to scale to meet expanding needs and the next big thing. Given that hardware will provide the power to increase throughput, Go’s garbage collector is being designed to favor low latency and tuning via only a single knob. Go 1.5 is the first big step down this path and these first steps will forever influence Go and the applications it best supports. This blog post gives a high-level overview of what we have done for the Go 1.5 collector.

今天，16G的内存要100美元，CPU有很多核心，每个核心都有多个硬件线程。十年后，这种硬件将显得古板，但今天用Go构建的软件将需要扩展，以满足不断扩大的需求和下一个大事件。鉴于硬件将提供提高吞吐量的能力，Go的垃圾收集器被设计成有利于低延迟和只通过一个旋钮进行调整。Go 1.5是在这条道路上迈出的第一步，这些第一步将永远影响Go和它所支持的应用程序。这篇博文对我们为 Go 1.5 采集器所做的工作进行了高层次的概述。

## The Embellishment 润色

To create a garbage collector for the next decade, we turned to an algorithm from decades ago. Go’s new garbage collector is a *concurrent*, *tri-color*, *mark-sweep* collector, an idea first proposed by [Dijkstra in 1978](http://dl.acm.org/citation.cfm?id=359655). This is a deliberate divergence from most “enterprise” grade garbage collectors of today, and one that we believe is well suited to the properties of modern hardware and the latency requirements of modern software.

为了创建一个面向未来十年的垃圾收集器，我们转向了几十年前的算法。Go的新垃圾收集器是一个并发的、三色的、标记扫除的收集器，这是Dijkstra在1978年首次提出的想法。这是与当今大多数 "企业 "级垃圾收集器的特意区别，我们认为它很适合现代硬件的特性和现代软件的延迟要求。

In a tri-color collector, every object is either white, grey, or black and we view the heap as a graph of connected objects. At the start of a GC cycle all objects are white. The GC visits all *roots*, which are objects directly accessible by the application such as globals and things on the stack, and colors these grey. The GC then chooses a grey object, blackens it, and then scans it for pointers to other objects. When this scan finds a pointer to a white object, it turns that object grey. This process repeats until there are no more grey objects. At this point, white objects are known to be unreachable and can be reused.

在三色收集器中，每个对象要么是白色的，要么是灰色的，要么是黑色的，我们把堆看作是一个连接对象的图。在GC周期的开始，所有对象都是白色的。GC会访问所有的根，也就是应用程序可以直接访问的对象，如globals和堆栈上的东西，并将其染成灰色。然后GC选择一个灰色的对象，将其涂黑，然后扫描它以寻找指向其他对象的指针。当这个扫描找到一个指向白色对象的指针时，它就把这个对象变成灰色。这个过程重复进行，直到不再有灰色对象。在这一点上，白色对象被认为是不可触及的，可以被重新使用。

This all happens concurrently with the application, known as the *mutator*, changing pointers while the collector is running. Hence, the mutator must maintain the invariant that no black object points to a white object, lest the garbage collector lose track of an object installed in a part of the heap it has already visited. Maintaining this invariant is the job of the *write barrier*, which is a small function run by the mutator whenever a pointer in the heap is modified. Go’s write barrier colors the now-reachable object grey if it is currently white, ensuring that the garbage collector will eventually scan it for pointers.

这一切都与被称为变体的应用程序同时发生，在收集器运行时改变指针。因此，突变器必须保持一个不变性，即没有黑对象指向白对象，以免垃圾收集器失去对安装在它已经访问过的堆的某个部分的对象的跟踪。维护这一不变性是写屏障的工作，它是一个小函数，每当堆中的指针被修改时，由变体运行。Go的写屏障将现在可访问的对象染成灰色，如果它目前是白色的话，确保垃圾收集器最终会扫描它的指针。

Deciding when the job of finding all grey objects is done is subtle and can be expensive and complicated if we want to avoid blocking the mutators. To keep things simple Go 1.5 does as much work as it can concurrently and then briefly stops the world to inspect all potential sources of grey objects. Finding the sweet spot between the time needed for this final stop-the-world and the total amount of work that this GC does is a major deliverable for Go 1.6.

决定何时完成寻找所有灰色对象的工作是很微妙的，如果我们想避免阻塞突变器，可能会很昂贵和复杂。为了保持简单，Go 1.5 尽可能多地并发工作，然后短暂地停止世界，检查所有潜在的灰色对象的来源。在最后停止世界所需的时间和这个GC所做的总工作量之间找到一个甜蜜的点，是Go 1.6的一个主要交付成果。

Of course the devil is in the details. When do we start a GC cycle? What metrics do we use to make that decision? How should the GC interact with the Go scheduler? How do we pause a mutator thread long enough to scan its stack?  How do we represent white, grey, and black so we can efficiently find and scan grey objects? How do we know where the roots are? How do we know where in an object pointers are located? How do we minimize memory fragmentation? How do we deal with cache performance issues? How big should the heap be? And on and on, some related to allocation, some to finding reachable objects, some related to scheduling, but many related to performance. Low-level discussions of each of these areas are beyond the scope of this blog post.

当然，魔鬼是在细节中的。我们什么时候开始一个GC周期？我们用什么指标来做这个决定？GC应该如何与Go的调度器互动？我们如何让突变器线程暂停足够长的时间来扫描其栈？ 我们如何表示白色、灰色和黑色，以便我们能够有效地找到和扫描灰色对象？我们如何知道根部在哪里？我们如何知道在一个对象中指针的位置？我们如何尽量减少内存碎片？我们如何处理缓冲区的性能问题？堆应该有多大？等等，有些与分配有关，有些与寻找可达对象有关，有些与调度有关，但许多与性能有关。对这些领域的低层次讨论超出了本博文的范围。

At a higher level, one approach to solving performance problems is to add GC knobs, one for each performance issue. The programmer can then turn the knobs in search of appropriate settings for their application. The downside is that after a decade with one or two new knobs each year you end up with the GC Knobs Turner Employment Act. Go is not going down that path. Instead we provide a single knob, called GOGC. This value controls the total size of the heap relative to the size of reachable objects. The default value of 100 means that total heap size is now 100% bigger than (i.e., twice) the size of the reachable objects after the last collection. 200 means total heap size is 200% bigger than (i.e., three times) the size of the reachable objects. If you want to lower the total time spent in GC, increase GOGC. If you want to trade more GC time for less memory, lower GOGC.

在更高的层次上，解决性能问题的一种方法是增加GC旋钮，每个性能问题都有一个。然后，程序员可以转动这些旋钮，为他们的应用寻找合适的设置。缺点是在十年后，每年都有一两个新的旋钮，最后你会发现GC旋钮特纳就业法案。Go不会走这条路。相反，我们提供一个单一的旋钮，称为GOGC。这个值控制堆的总大小，相对于可触及对象的大小。默认值为100，意味着现在的总堆大小比上次收集后的可达对象大小大100%（也就是两倍）。200意味着总堆的大小比可达对象的大小大200%（即三倍）。如果你想降低花在GC上的总时间，增加GOGC。如果你想用更多的GC时间换取更少的内存，就降低GOGC。

More importantly as RAM doubles with the next generation of hardware, simply doubling GOGC will halve the number of GC cycles. On the other hand since GOGC is based on reachable object size, doubling the load by doubling the reachable objects requires no retuning. The application just scales. Furthermore, unencumbered by ongoing support for dozens of knobs, the runtime team can focus on improving the runtime based on feedback from real customer applications.

更重要的是，随着下一代硬件的RAM翻倍，简单地将GOGC翻倍将使GC周期的数量减半。另一方面，由于GOGC是基于可触及对象的大小，通过增加可触及对象的一倍来增加负载，不需要重新调整。应用程序只是扩展了。此外，由于没有对几十个旋钮的持续支持，运行时团队可以专注于根据实际客户应用的反馈来改进运行时。

## The Punchline 冲锋号

Go 1.5’s GC ushers in a future where stop-the-world pauses are no longer a barrier to moving to a safe and secure language. It is a future where applications scale effortlessly along with hardware and as hardware becomes more powerful the GC will not be an impediment to better, more scalable software. It’s a good place to be for the next decade and beyond. For more details about the 1.5 GC and how we eliminated latency issues see the [Go GC: Latency Problem Solved presentation](https://www.youtube.com/watch?v=aiv1JOfMjm0) or [the slides](https://go.dev/talks/2015/go-gc.pdf).

Go 1.5 的 GC 带来了这样一个未来，即停止世界的暂停不再是转向安全和可靠语言的障碍。在这个未来，应用程序可以毫不费力地与硬件一起扩展，随着硬件变得更加强大，GC将不会成为更好、更可扩展的软件的障碍。这是在未来十年及以后的一个好地方。关于1.5 GC的更多细节，以及我们如何消除延迟问题，请参见Go GC：延迟问题解决的演讲或幻灯片。
