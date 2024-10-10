+++
title = "介绍 Go 竞争检测器"
weight = 9
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Introducing the Go Race Detector - 介绍 Go 竞争检测器

> 原文：[https://go.dev/blog/race-detector](https://go.dev/blog/race-detector)

Dmitry Vyukov and Andrew Gerrand
26 June 2013

## Introduction 简介

[Race conditions](http://en.wikipedia.org/wiki/Race_condition) are among the most insidious and elusive programming errors. They typically cause erratic and mysterious failures, often long after the code has been deployed to production. While Go’s concurrency mechanisms make it easy to write clean concurrent code, they don’t prevent race conditions. Care, diligence, and testing are required. And tools can help.

竞赛条件是最隐蔽、最难以捉摸的编程错误之一。它们通常会导致不稳定和神秘的故障，往往是在代码部署到生产中很久之后。虽然 Go 的并发机制使编写干净的并发代码变得容易，但它们并不能防止出现竞赛条件。我们需要谨慎、勤奋和测试。而工具可以提供帮助。

We’re happy to announce that Go 1.1 includes a [race detector](https://go.dev/doc/articles/race_detector.html), a new tool for finding race conditions in Go code. It is currently available for Linux, OS X, and Windows systems with 64-bit x86 processors.

我们很高兴地宣布，Go 1.1包含了一个竞赛检测器，这是一个用于查找Go代码中竞赛条件的新工具。它目前可用于Linux、OS X和Windows系统的64位x86处理器。

The race detector is based on the C/C++ [ThreadSanitizer runtime library](https://github.com/google/sanitizers), which has been used to detect many errors in Google’s internal code base and in [Chromium](http://www.chromium.org/). The technology was integrated with Go in September 2012; since then it has detected [42 races](https://github.com/golang/go/issues?utf8=✓&q=ThreadSanitizer) in the standard library. It is now part of our continuous build process, where it continues to catch race conditions as they arise.

该竞赛检测器是基于C/C++ ThreadSanitizer运行时库，该库已被用于检测谷歌内部代码库和Chromium中的许多错误。该技术于2012年9月被集成到Go中；从那时起，它已经在标准库中检测出42个竞赛。它现在是我们持续构建过程的一部分，在那里它继续捕捉出现的竞赛情况。

## How it works 它是如何工作的

The race detector is integrated with the go tool chain. When the `-race` command-line flag is set, the compiler instruments all memory accesses with code that records when and how the memory was accessed, while the runtime library watches for unsynchronized accesses to shared variables. When such "racy" behavior is detected, a warning is printed. (See [this article](https://github.com/google/sanitizers/wiki/ThreadSanitizerAlgorithm) for the details of the algorithm.)

竞赛检测器与Go工具链集成。当设置-race命令行标志时，编译器会用代码记录所有内存访问的时间和方式，而运行时库会观察对共享变量的非同步访问。当检测到这种 "淫秽 "行为时，就会打印出一个警告。(该算法的细节见本文）。

Because of its design, the race detector can detect race conditions only when they are actually triggered by running code, which means it’s important to run race-enabled binaries under realistic workloads. However, race-enabled binaries can use ten times the CPU and memory, so it is impractical to enable the race detector all the time. One way out of this dilemma is to run some tests with the race detector enabled. Load tests and integration tests are good candidates, since they tend to exercise concurrent parts of the code. Another approach using production workloads is to deploy a single race-enabled instance within a pool of running servers.

由于其设计，竞赛检测器只有在运行的代码实际触发竞赛条件时才能检测到，这意味着在现实的工作负载下运行支持竞赛的二进制文件非常重要。然而，启用了竞赛的二进制文件可以使用十倍的CPU和内存，所以一直启用竞赛检测器是不现实的。摆脱这种困境的一个方法是在启用竞赛检测器的情况下运行一些测试。负载测试和集成测试是很好的选择，因为它们倾向于锻炼代码的并发部分。另一个使用生产工作负载的方法是在一个运行的服务器池中部署一个启用了竞赛的实例。

## Using the race detector 使用竞赛检测器

The race detector is fully integrated with the Go tool chain. To build your code with the race detector enabled, just add the `-race` flag to the command line:

竞赛检测器与Go工具链完全集成。要在启用竞赛检测器的情况下构建代码，只需在命令行中添加-race标志。

```shell linenums="1"
$ go test -race mypkg    // test the package
$ go run -race mysrc.go  // compile and run the program
$ go build -race mycmd   // build the command
$ go install -race mypkg // install the package
```

To try out the race detector for yourself, copy this example program into `racy.go`:

要想自己尝试一下竞赛检测器，请将此示例程序复制到racy.go中：

```go
package main

import "fmt"

func main() {
    done := make(chan bool)
    m := make(map[string]string)
    m["name"] = "world"
    go func() {
        m["name"] = "data race"
        done <- true
    }()
    fmt.Println("Hello,", m["name"])
    <-done
}
```

Then run it with the race detector enabled:

然后在启用竞赛检测器的情况下运行它：

```shell linenums="1"
$ go run -race racy.go
```

## Examples 例子

Here are two examples of real issues caught by the race detector.

下面是两个被竞赛检测器发现的真实问题的例子。

### Example 1: Timer.Reset 例子1：Timer.Reset

The first example is a simplified version of an actual bug found by the race detector. It uses a timer to print a message after a random duration between 0 and 1 second. It does so repeatedly for five seconds. It uses [`time.AfterFunc`](https://go.dev/pkg/time/#AfterFunc) to create a [`Timer`](https://go.dev/pkg/time/#Timer) for the first message and then uses the [`Reset`](https://go.dev/pkg/time/#Timer.Reset) method to schedule the next message, re-using the `Timer` each time.

第一个例子是竞赛检测器发现的一个实际错误的简化版本。它使用一个定时器，在0到1秒之间的随机时间后打印一条信息。它重复这样做了五秒钟。它使用time.AfterFunc为第一条消息创建一个Timer，然后使用Reset方法来安排下一条消息，每次都重新使用Timer。

```go
10  func main() {
11      start := time.Now()
12      var t *time.Timer
13      t = time.AfterFunc(randomDuration(), func() {
14          fmt.Println(time.Now().Sub(start))
15          t.Reset(randomDuration())
16      })
17      time.Sleep(5 * time.Second)
18  }
19  
20  func randomDuration() time.Duration {
21      return time.Duration(rand.Int63n(1e9))
22  }
23  
```

Run 运行

This looks like reasonable code, but under certain circumstances it fails in a surprising way:

这看起来是合理的代码，但在某些情况下，它以一种令人惊讶的方式失败：

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xb code=0x1 addr=0x8 pc=0x41e38a]

goroutine 4 [running]:
time.stopTimer(0x8, 0x12fe6b35d9472d96)
    src/pkg/runtime/ztime_linux_amd64.c:35 +0x25
time.(*Timer).Reset(0x0, 0x4e5904f, 0x1)
    src/pkg/time/sleep.go:81 +0x42
main.func·001()
    race.go:14 +0xe3
created by time.goFunc
    src/pkg/time/sleep.go:122 +0x48
```

What’s going on here? Running the program with the race detector enabled is more illuminating:

这里发生了什么？在启用竞赛检测器的情况下运行该程序更有启示意义：

```
==================
WARNING: DATA RACE
Read by goroutine 5:
  main.func·001()
     race.go:16 +0x169

Previous write by goroutine 1:
  main.main()
      race.go:14 +0x174

Goroutine 5 (running) created at:
  time.goFunc()
      src/pkg/time/sleep.go:122 +0x56
  timerproc()
     src/pkg/runtime/ztime_linux_amd64.c:181 +0x189
==================
```

The race detector shows the problem: an unsynchronized read and write of the variable `t` from different goroutines. If the initial timer duration is very small, the timer function may fire before the main goroutine has assigned a value to `t` and so the call to `t.Reset` is made with a nil `t`.

竞赛检测器显示了问题所在：从不同的goroutine对变量t进行了不同步的读写。如果初始定时器的持续时间非常小，定时器函数可能在主goroutine给t赋值之前就启动了，因此对t.Reset的调用是以nil t进行的。

To fix the race condition we change the code to read and write the variable `t` only from the main goroutine:

为了解决这个竞赛条件，我们改变了代码，只从主goroutine中读写变量t：

```go
10  func main() {
11      start := time.Now()
12      reset := make(chan bool)
13      var t *time.Timer
14      t = time.AfterFunc(randomDuration(), func() {
15          fmt.Println(time.Now().Sub(start))
16          reset <- true
17      })
18      for time.Since(start) < 5*time.Second {
19          <-reset
20          t.Reset(randomDuration())
21      }
22  }
23  
```

Run 运行

Here the main goroutine is wholly responsible for setting and resetting the `Timer` `t` and a new reset channel communicates the need to reset the timer in a thread-safe way.

在这里，主程序完全负责设置和重置定时器t，一个新的重置通道以线程安全的方式传达了重置定时器的需求。

A simpler but less efficient approach is to [avoid reusing timers](http://play.golang.org/p/kuWTrY0pS4).

一个更简单但效率较低的方法是避免重复使用定时器。

### Example 2: ioutil.Discard 例2：ioutil.Discard

The second example is more subtle.

第二个例子更微妙。

The `ioutil` package’s [`Discard`](https://go.dev/pkg/io/ioutil/#Discard) object implements [`io.Writer`](https://go.dev/pkg/io/#Writer), but discards all the data written to it. Think of it like `/dev/null`: a place to send data that you need to read but don’t want to store. It is commonly used with [`io.Copy`](https://go.dev/pkg/io/#Copy) to drain a reader, like this:

ioutil包的Discard对象实现了io.Writer，但是丢弃了所有写给它的数据。可以把它想象成/dev/null：一个用来发送您需要读取但不想存储的数据的地方。它通常与io.Copy一起使用，以耗尽一个阅读器，像这样。

```go
io.Copy(ioutil.Discard, reader)
```

Back in July 2011 the Go team noticed that using `Discard` in this way was inefficient: the `Copy` function allocates an internal 32 kB buffer each time it is called, but when used with `Discard` the buffer is unnecessary since we’re just throwing the read data away. We thought that this idiomatic use of `Copy` and `Discard` should not be so costly.

早在2011年7月，Go团队就注意到以这种方式使用Discard是低效的：Copy函数在每次调用时都会分配一个32 kB的内部缓冲区，但当与Discard一起使用时，缓冲区就没有必要了，因为我们只是把读取的数据扔掉了。我们认为Copy和Discard的这种习惯性使用不应该是如此昂贵的。

The fix was simple. If the given `Writer` implements a `ReadFrom` method, a `Copy` call like this:

修复方法很简单。如果给定的Writer实现了ReadFrom方法，像这样的Copy调用：

```go
io.Copy(writer, reader)
```

is delegated to this potentially more efficient call:

被委托给这个潜在的更有效的调用：

```go
writer.ReadFrom(reader)
```

We [added a ReadFrom method](https://go.dev/cl/4817041) to Discard’s underlying type, which has an internal buffer that is shared between all its users. We knew this was theoretically a race condition, but since all writes to the buffer should be thrown away we didn’t think it was important.

我们给Discard的底层类型添加了一个ReadFrom方法，它有一个内部缓冲区，在所有用户之间共享。我们知道这在理论上是一个竞赛条件，但由于所有对缓冲区的写入都应该被丢弃，我们认为这并不重要。

When the race detector was implemented it immediately [flagged this code](https://go.dev/issue/3970) as racy. Again, we considered that the code might be problematic, but decided that the race condition wasn’t "real". To avoid the "false positive" in our build we implemented [a non-racy version](https://go.dev/cl/6624059) that is enabled only when the race detector is running.

当竞赛检测器被实施时，它立即将这段代码标记为狂妄。我们再次考虑到这段代码可能有问题，但决定这个竞赛条件不是 "真正的"。为了避免在我们的构建中出现 "假阳性"，我们实现了一个只有在竞赛检测器运行时才会启用的非狂热版本。

But a few months later [Brad](https://bradfitz.com/) encountered a [frustrating and strange bug](https://go.dev/issue/4589). After a few days of debugging, he narrowed it down to a real race condition caused by `ioutil.Discard`.

但几个月后，布拉德遇到了一个令人沮丧的奇怪的错误。经过几天的调试，他把范围缩小到由ioutil.Discard引起的一个真正的竞赛条件。

Here is the known-racy code in `io/ioutil`, where `Discard` is a `devNull` that shares a single buffer between all of its users.

下面是io/ioutil中已知的错误代码，其中Discard是一个devNull，在所有用户之间共享一个缓冲区。

```go
var blackHole [4096]byte // shared buffer

func (devNull) ReadFrom(r io.Reader) (n int64, err error) {
    readSize := 0
    for {
        readSize, err = r.Read(blackHole[:])
        n += int64(readSize)
        if err != nil {
            if err == io.EOF {
                return n, nil
            }
            return
        }
    }
}
```

Brad’s program includes a `trackDigestReader` type, which wraps an `io.Reader` and records the hash digest of what it reads.

Brad的程序包括一个trackDigestReader类型，它包装了一个io.Reader，并记录了它所读取的哈希摘要。

```go
type trackDigestReader struct {
    r io.Reader
    h hash.Hash
}

func (t trackDigestReader) Read(p []byte) (n int, err error) {
    n, err = t.r.Read(p)
    t.h.Write(p[:n])
    return
}
```

For example, it could be used to compute the SHA-1 hash of a file while reading it:

例如，它可以用来在读取文件时计算文件的SHA-1哈希值：

```go
tdr := trackDigestReader{r: file, h: sha1.New()}
io.Copy(writer, tdr)
fmt.Printf("File hash: %x", tdr.h.Sum(nil))
```

In some cases there would be nowhere to write the data—but still a need to hash the file—and so `Discard` would be used:

在某些情况下，没有地方可以写入数据，但仍然需要对文件进行哈希处理，因此将使用Discard：

```go
io.Copy(ioutil.Discard, tdr)
```

But in this case the `blackHole` buffer isn’t just a black hole; it is a legitimate place to store the data between reading it from the source `io.Reader` and writing it to the `hash.Hash`. With multiple goroutines hashing files simultaneously, each sharing the same `blackHole` buffer, the race condition manifested itself by corrupting the data between reading and hashing. No errors or panics occurred, but the hashes were wrong. Nasty!

但在这种情况下，blackHole缓冲区不仅仅是一个黑洞；它是一个合法的地方，在从源io.Reader读取数据和将其写入hash.Hash之间存储数据。在多个goroutines同时对文件进行哈希运算的情况下，每个人都共享同一个blackHole缓冲区，竞赛条件表现为在读取和哈希运算之间损坏数据。没有发生错误或恐慌，但哈希值是错误的。糟糕的是!

```go
func (t trackDigestReader) Read(p []byte) (n int, err error) {
    // the buffer p is blackHole
    n, err = t.r.Read(p)
    // p may be corrupted by another goroutine here,
    // between the Read above and the Write below
    t.h.Write(p[:n])
    return
}
```

The bug was finally [fixed](https://go.dev/cl/7011047) by giving a unique buffer to each use of `ioutil.Discard`, eliminating the race condition on the shared buffer.

通过给ioutil.Discard的每次使用提供一个唯一的缓冲区，消除了共享缓冲区上的竞赛条件，这个错误最终得到了解决。

## Conclusions 结论

The race detector is a powerful tool for checking the correctness of concurrent programs. It will not issue false positives, so take its warnings seriously. But it is only as good as your tests; you must make sure they thoroughly exercise the concurrent properties of your code so that the race detector can do its job.

竞赛检测器是检查并发程序正确性的一个强大工具。它不会发出误报，所以要认真对待它的警告。但是，它只和您的测试一样好；您必须确保它们彻底锻炼了您的代码的并发特性，这样，竞赛检测器才能完成它的工作。

What are you waiting for? Run `"go test -race"` on your code today!

您还在等什么呢？今天就在您的代码上运行 "go test -race "吧!
