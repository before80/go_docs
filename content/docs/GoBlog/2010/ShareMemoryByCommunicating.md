+++
title = "通过通信共享内存"
weight = 8
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Share Memory By Communicating - 通过通信共享内存

https://go.dev/blog/codelab-share

Andrew Gerrand
13 July 2010

2010年7月13日

​	传统的线程模型（例如在编写Java、C++和Python程序时通常使用）要求程序员使用共享内存在线程之间进行通信。通常，共享数据结构由锁保护，线程将争用这些锁以访问数据。在某些情况下，通过使用线程安全的数据结构（如Python的队列）可以使这个过程更容易。

​	Go的并发原语——goroutines和channels——提供了一种优雅而独特的方式来构建并发软件。 （这些概念的[历史](https://swtch.com/~rsc/thread/)始于C. A. R. Hoare的"[通信顺序处理](http://www.usingcsp.com/)"。）Go鼓励使用通道（channel）来在goroutines之间传递对数据的引用，而不是显式地使用锁来协调对共享数据的访问。这种方法确保在任何时候只有一个goroutine能够访问数据。这个概念在[《Effective Go》](https://go.dev/doc/effective_go.html)（任何Go程序员必读的文档）中总结如下：

Go的并发原语--goroutines和channel--提供了一种优雅而独特的结构化并发软件的方法。(这些概念有一个有趣的历史，始于C. A. R. Hoare的Communicating Sequential Processes)。Go鼓励使用通道在goroutine之间传递对数据的引用，而不是明确地使用锁来调解对共享数据的访问。这种方法可以确保在给定时间内只有一个goroutine可以访问数据。这个概念在Effective Go文件中得到了总结（任何Go程序员都必须阅读）。

​	***不要通过共享内存来通信；通过通信来共享内存。***

​	考虑一个轮询URL列表的程序。在传统的线程环境中，可能会像这样构造它的数据：

```go linenums="1"
type Resource struct {
    url        string
    polling    bool
    lastPolled int64
}

type Resources struct {
    data []*Resource
    lock *sync.Mutex
}
```

然后Poller函数（许多这样的函数将在单独的线程中运行）可能是这样的：

```go linenums="1"
func Poller(res *Resources) {
    for {
        // 获取最近未被轮询的资源，并将其标记为正在轮询
        res.lock.Lock()
        var r *Resource
        for _, v := range res.data {
            if v.polling {
                continue
            }
            if r == nil || v.lastPolled < r.lastPolled {
                r = v
            }
        }
        if r != nil {
            r.polling = true
        }
        res.lock.Unlock()
        if r == nil {
            continue
        }

        // 对 URL 进行轮询

        // 更新Resource的轮询状态和最后轮询时间
        res.lock.Lock()
        r.polling = false
        r.lastPolled = time.Nanoseconds()
        res.lock.Unlock()
    }
}
```

​	这个函数大约有一页长，并需要更多细节才能完整。它甚至不包括URL轮询逻辑（本身只有几行），也不会优雅地处理资源池耗尽的情况。

​	现在我们来看一下使用Go风格实现相同功能的代码。在这个例子中，Poller是一个函数，它从一个输入通道接收要轮询的Resources，并在它们完成时将它们发送到一个输出通道。

```go linenums="1"
type Resource string

func Poller(in, out chan *Resource) {
    for r := range in {
        // 对 URL 进行轮询

        // 将处理完的资源发送到 out
        out <- r
    }
}
```

​	前面示例中的复杂逻辑已经不见了，我们的 Resource 数据结构也不再包含记录数据。实际上，只剩下了最重要的部分。这应该能让您领略到这些简单语言特性的威力。

​	上面的代码片段省略了很多内容。如果您想看一个完整的、符合 Go 语言惯例的程序示例，涉及这些想法，请参见"[共享内存通过通信](https://go.dev/doc/codewalk/sharemem/)"（Share Memory By Communicating）的代码漫步。
