+++
title = "并发不是并行"
weight = 19
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Concurrency is not parallelism - 并发不是并行

https://go.dev/blog/waza-talk

Andrew Gerrand
16 January 2013

If there’s one thing most people know about Go, is that it is designed for concurrency. No introduction to Go is complete without a demonstration of its goroutines and channels.

如果说大多数人对Go有什么了解，那就是它是为并发而设计的。如果不展示Go的goroutines和channel，对Go的介绍就不完整。

But when people hear the word *concurrency* they often think of *parallelism*, a related but quite distinct concept. In programming, concurrency is the *composition* of independently executing processes, while parallelism is the simultaneous *execution* of (possibly related) computations. Concurrency is about *dealing with* lots of things at once. Parallelism is about *doing* lots of things at once.

但是，当人们听到并发这个词时，他们往往会想到并行，这是一个相关的但又相当不同的概念。在编程中，并发是独立执行进程的组合，而并行是同时执行（可能相关的）计算。并发是指一次处理很多事情。平行性是指一次做很多事情。

To clear up this conflation, Rob Pike gave a talk at [Heroku](http://heroku.com/)’s Waza conference entitled [*Concurrency is not parallelism*](https://blog.heroku.com/concurrency_is_not_parallelism), and a video recording of the talk was released a few months ago.

为了澄清这种混淆，Rob Pike在Heroku的Waza会议上发表了题为 "并发性不是并行性 "的演讲，几个月前发布了该演讲的视频记录。

<iframe src="https://www.youtube.com/embed/oV9rvDllKEg" width="500" height="281" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

The slides are available at [go.dev/talks](https://go.dev/talks/2012/waza.slide) (use the left and right arrow keys to navigate).

幻灯片可以在go.dev/talks上找到（使用左右方向键进行导航）。

To learn about Go’s concurrency primitives, watch [Go concurrency patterns](http://www.youtube.com/watch?v=f6kdp27TYZs) ([slides](https://go.dev/talks/2012/concurrency.slide)).

要了解Go的并发原语，请观看Go并发模式（幻灯片）。
