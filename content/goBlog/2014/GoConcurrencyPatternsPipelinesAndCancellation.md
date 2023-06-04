+++
title = "go 并发模式：Pipelines和取消"
weight = 13
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go Concurrency Patterns: Pipelines and cancellation - go 并发模式：Pipelines和取消

https://go.dev/blog/pipelines

Sameer Ajmani
13 March 2014

## Introduction 简介

Go’s concurrency primitives make it easy to construct streaming data pipelines that make efficient use of I/O and multiple CPUs. This article presents examples of such pipelines, highlights subtleties that arise when operations fail, and introduces techniques for dealing with failures cleanly.

Go的并发基元使其能够轻松构建流式数据管道，有效利用I/O和多个CPU。本文介绍了这种管道的例子，强调了操作失败时出现的微妙情况，并介绍了干净地处理失败的技术。

## What is a pipeline? 什么是流水线？

There’s no formal definition of a pipeline in Go; it’s just one of many kinds of concurrent programs. Informally, a pipeline is a series of *stages* connected by channels, where each stage is a group of goroutines running the same function. In each stage, the goroutines

Go中没有管道的正式定义；它只是众多并发程序中的一种。非正式地讲，流水线是一系列由通道连接的阶段，每个阶段是一组运行相同函数的goroutines。在每个阶段中，goroutines

- receive values from *upstream* via *inbound* channels 通过入站通道接收来自上游的数值
- perform some function on that data, usually producing new values 对该数据执行一些功能，通常产生新的值
- send values *downstream* via *outbound* channels 通过出站通道向下游发送数值

Each stage has any number of inbound and outbound channels, except the first and last stages, which have only outbound or inbound channels, respectively. The first stage is sometimes called the *source* or *producer*; the last stage, the *sink* or *consumer*.

每个阶段都有任意数量的入站和出站通道，除了第一和最后一个阶段，它们分别只有出站或入站通道。第一阶段有时被称为源或生产者；最后一个阶段是汇或消费者。

We’ll begin with a simple example pipeline to explain the ideas and techniques. Later, we’ll present a more realistic example.

我们将从一个简单的管道例子开始，解释这些想法和技术。稍后，我们将介绍一个更现实的例子。

## Squaring numbers 数字的平方

Consider a pipeline with three stages.

考虑一个有三个阶段的流水线。

The first stage, `gen`, is a function that converts a list of integers to a channel that emits the integers in the list. The `gen` function starts a goroutine that sends the integers on the channel and closes the channel when all the values have been sent:

第一阶段，gen，是一个将一个整数列表转换为一个通道的函数，该通道将列表中的整数发射出来。gen函数启动一个goroutine，在通道上发送整数，当所有的值都发送完后关闭通道：

```go linenums="1"
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}
```

The second stage, `sq`, receives integers from a channel and returns a channel that emits the square of each received integer. After the inbound channel is closed and this stage has sent all the values downstream, it closes the outbound channel:

第二阶段，sq，从一个通道接收整数，并返回一个通道，发射每个接收的整数的平方。在入站通道关闭后，这个阶段将所有的值发送到下游，它将关闭出站通道：

```go linenums="1"
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}
```

The `main` function sets up the pipeline and runs the final stage: it receives values from the second stage and prints each one, until the channel is closed:

main函数设置了管道并运行了最后阶段：它从第二阶段接收数值并打印每一个数值，直到通道关闭：

```go linenums="1"
func main() {
    // Set up the pipeline.
    c := gen(2, 3)
    out := sq(c)

    // Consume the output.
    fmt.Println(<-out) // 4
    fmt.Println(<-out) // 9
}
```

Since `sq` has the same type for its inbound and outbound channels, we can compose it any number of times. We can also rewrite `main` as a range loop, like the other stages:

由于sq的入站和出站通道具有相同的类型，我们可以对它进行任意次数的组合。我们也可以把main改写成一个范围循环，就像其他阶段一样：

```go linenums="1"
func main() {
    // Set up the pipeline and consume the output.
    for n := range sq(sq(gen(2, 3))) {
        fmt.Println(n) // 16 then 81
    }
}
```

## Fan-out, fan-in 扇出、扇入

Multiple functions can read from the same channel until that channel is closed; this is called *fan-out*. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.

多个函数可以从同一个通道中读取数据，直到该通道关闭；这被称为扇出。这提供了一种在一组工作者之间分配工作的方法，以并行化CPU的使用和I/O。

A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that’s closed when all the inputs are closed. This is called *fan-in*.

一个函数可以从多个输入中读取，并通过将输入通道复用到一个单一的通道上，直到所有的输入都被关闭为止。这就是所谓的扇入。

We can change our pipeline to run two instances of `sq`, each reading from the same input channel. We introduce a new function, *merge*, to fan in the results:

我们可以改变我们的管道来运行两个sq的实例，每个实例从同一个输入通道读取数据。我们引入一个新的函数，merge，来扇入结果：

```go linenums="1"
func main() {
    in := gen(2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    c1 := sq(in)
    c2 := sq(in)

    // Consume the merged output from c1 and c2.
    for n := range merge(c1, c2) {
        fmt.Println(n) // 4 then 9, or 9 then 4
    }
}
```

The `merge` function converts a list of channels to a single channel by starting a goroutine for each inbound channel that copies the values to the sole outbound channel. Once all the `output` goroutines have been started, `merge` starts one more goroutine to close the outbound channel after all sends on that channel are done.

merge函数通过为每个入站通道启动一个goroutine，将数值复制到唯一的出站通道，从而将一个通道列表转换为一个通道。一旦所有的输出goroutine被启动，merge再启动一个goroutine，在该通道的所有发送完成后关闭出站通道。

Sends on a closed channel panic, so it’s important to ensure all sends are done before calling close. The [`sync.WaitGroup`](https://go.dev/pkg/sync/#WaitGroup) type provides a simple way to arrange this synchronization:

在一个关闭的通道上的发送会恐慌，所以在调用close之前确保所有的发送都完成是很重要的。sync.WaitGroup类型提供了一种简单的方法来安排这种同步：

```go linenums="1"
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```

## Stopping short 短暂停止

There is a pattern to our pipeline functions:

我们的流水线函数有一个模式：

- stages close their outbound channels when all the send operations are done. 当所有的发送操作完成后，阶段性地关闭其出站通道。
- stages keep receiving values from inbound channels until those channels are closed. 阶段一直从入站通道接收数值，直到这些通道被关闭。

This pattern allows each receiving stage to be written as a `range` loop and ensures that all goroutines exit once all values have been successfully sent downstream.

这种模式允许将每个接收阶段写成一个范围循环，并确保所有的goroutines在所有的值被成功发送到下游时退出。

But in real pipelines, stages don’t always receive all the inbound values. Sometimes this is by design: the receiver may only need a subset of values to make progress. More often, a stage exits early because an inbound value represents an error in an earlier stage. In either case the receiver should not have to wait for the remaining values to arrive, and we want earlier stages to stop producing values that later stages don’t need.

但是在真实的管道中，各阶段并不总是收到所有的入站值。有时这是设计好的：接收器可能只需要一个子集的值来取得进展。更常见的情况是，一个阶段提前退出，因为一个入站值代表了一个早期阶段的错误。无论哪种情况，接收者都不应该等待剩余的值到来，我们希望早期阶段停止产生后期阶段不需要的值。

In our example pipeline, if a stage fails to consume all the inbound values, the goroutines attempting to send those values will block indefinitely:

在我们的例子管道中，如果一个阶段不能消耗所有的入站值，试图发送这些值的goroutines将无限期地阻塞：

```go linenums="1"
    // Consume the first value from the output.
    out := merge(c1, c2)
    fmt.Println(<-out) // 4 or 9
    return
    // Since we didn't receive the second value from out,
    // one of the output goroutines is hung attempting to send it.
}
```

This is a resource leak: goroutines consume memory and runtime resources, and heap references in goroutine stacks keep data from being garbage collected. Goroutines are not garbage collected; they must exit on their own.

这是一个资源泄漏：goroutines消耗内存和运行时资源，goroutine栈中的堆引用使数据不被垃圾收集。goroutines不被垃圾收集，它们必须自己退出。

We need to arrange for the upstream stages of our pipeline to exit even when the downstream stages fail to receive all the inbound values. One way to do this is to change the outbound channels to have a buffer. A buffer can hold a fixed number of values; send operations complete immediately if there’s room in the buffer:

我们需要安排管道的上游阶段退出，即使下游阶段未能接收所有的入站值。做到这一点的一个方法是将出站通道改为有一个缓冲区。缓冲区可以容纳固定数量的值；如果缓冲区有空间，发送操作立即完成：

```go linenums="1"
c := make(chan int, 2) // buffer size 2
c <- 1  // succeeds immediately
c <- 2  // succeeds immediately
c <- 3  // blocks until another goroutine does <-c and receives 1
```

When the number of values to be sent is known at channel creation time, a buffer can simplify the code. For example, we can rewrite `gen` to copy the list of integers into a buffered channel and avoid creating a new goroutine:

当要发送的值的数量在通道创建时就已经知道，缓冲区可以简化代码。例如，我们可以重写gen，将整数列表复制到一个缓冲通道中，避免创建一个新的goroutine：

```go linenums="1"
func gen(nums ...int) <-chan int {
    out := make(chan int, len(nums))
    for _, n := range nums {
        out <- n
    }
    close(out)
    return out
}
```

Returning to the blocked goroutines in our pipeline, we might consider adding a buffer to the outbound channel returned by `merge`:

回到我们的管道中被阻塞的goroutines，我们可能会考虑给merge返回的出站通道添加一个缓冲区：

```go linenums="1"
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int, 1) // enough space for the unread inputs
    // ... the rest is unchanged ...
```

While this fixes the blocked goroutine in this program, this is bad code. The choice of buffer size of 1 here depends on knowing the number of values `merge` will receive and the number of values downstream stages will consume. This is fragile: if we pass an additional value to `gen`, or if the downstream stage reads any fewer values, we will again have blocked goroutines.

虽然这修复了这个程序中的阻塞的goroutine，但这是坏的代码。这里选择缓冲区大小为1，取决于知道merge将收到的值的数量和下游阶段将消耗的值的数量。这是很脆弱的：如果我们传递一个额外的值给gen，或者如果下游阶段读取的值再少，我们又会出现阻塞的goroutines。

Instead, we need to provide a way for downstream stages to indicate to the senders that they will stop accepting input.

相反，我们需要提供一种方法，让下游阶段向发送者表明，他们将停止接受输入。

## Explicit cancellation 明确取消

When `main` decides to exit without receiving all the values from `out`, it must tell the goroutines in the upstream stages to abandon the values they’re trying to send. It does so by sending values on a channel called `done`. It sends two values since there are potentially two blocked senders:

当main决定在没有收到所有来自out的值的情况下退出时，它必须告诉上游阶段的goroutines放弃它们试图发送的值。它通过在一个叫做done的通道上发送数值来实现这一目的。它发送两个值，因为可能有两个阻塞的发送者：

```go linenums="1"
func main() {
    in := gen(2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    c1 := sq(in)
    c2 := sq(in)

    // Consume the first value from output.
    done := make(chan struct{}, 2)
    out := merge(done, c1, c2)
    fmt.Println(<-out) // 4 or 9

    // Tell the remaining senders we're leaving.
    done <- struct{}{}
    done <- struct{}{}
}
```

The sending goroutines replace their send operation with a `select` statement that proceeds either when the send on `out` happens or when they receive a value from `done`. The value type of `done` is the empty struct because the value doesn’t matter: it is the receive event that indicates the send on `out` should be abandoned. The `output` goroutines continue looping on their inbound channel, `c`, so the upstream stages are not blocked. (We’ll discuss in a moment how to allow this loop to return early.)

发送goroutines用一个select语句来代替他们的发送操作，该语句要么在out上的发送发生时，要么在他们从doed上收到一个值时进行。done的值类型是空结构，因为它的值并不重要：它是表明应该放弃发送的接收事件。输出的goroutines继续在他们的入站通道c上循环，所以上游阶段不会被阻断。 我们稍后会讨论如何让这个循环提前返回。

```go linenums="1"
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed or it receives a value
    // from done, then output calls wg.Done.
    output := func(c <-chan int) {
        for n := range c {
            select {
            case out <- n:
            case <-done:
            }
        }
        wg.Done()
    }
    // ... the rest is unchanged ...
```

This approach has a problem: *each* downstream receiver needs to know the number of potentially blocked upstream senders and arrange to signal those senders on early return. Keeping track of these counts is tedious and error-prone.

这种方法有一个问题：每个下游的接收者需要知道可能被屏蔽的上游发送者的数量，并安排在提前返回时向这些发送者发出信号。追踪这些计数是很繁琐的，而且容易出错。

We need a way to tell an unknown and unbounded number of goroutines to stop sending their values downstream. In Go, we can do this by closing a channel, because [a receive operation on a closed channel can always proceed immediately, yielding the element type’s zero value.](https://go.dev/ref/spec#Receive_operator)

我们需要一种方法来告诉未知的、不受约束的goroutines停止向下游发送其数值。在Go中，我们可以通过关闭一个通道来做到这一点，因为在一个关闭的通道上的接收操作总是可以立即进行，产生元素类型的零值。

This means that `main` can unblock all the senders simply by closing the `done` channel. This close is effectively a broadcast signal to the senders. We extend *each* of our pipeline functions to accept `done` as a parameter and arrange for the close to happen via a `defer` statement, so that all return paths from `main` will signal the pipeline stages to exit.

这意味着main可以通过关闭已完成的通道来解除对所有发送者的封锁。这种关闭实际上是对发送者的一种广播信号。我们将每个管道函数扩展为接受doed作为参数，并通过defer语句安排关闭，这样，所有来自main的返回路径都会向管道阶段发出退出信号。

```go linenums="1"
func main() {
    // Set up a done channel that's shared by the whole pipeline,
    // and close that channel when this pipeline exits, as a signal
    // for all the goroutines we started to exit.
    done := make(chan struct{})
    defer close(done)          

    in := gen(done, 2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    c1 := sq(done, in)
    c2 := sq(done, in)

    // Consume the first value from output.
    out := merge(done, c1, c2)
    fmt.Println(<-out) // 4 or 9

    // done will be closed by the deferred call.      
}
```

Each of our pipeline stages is now free to return as soon as `done` is closed. The `output` routine in `merge` can return without draining its inbound channel, since it knows the upstream sender, `sq`, will stop attempting to send when `done` is closed. `output` ensures `wg.Done` is called on all return paths via a `defer` statement:

我们的每个管道阶段现在都可以自由地返回，只要doed被关闭。merge中的输出例程可以在不耗尽其入站通道的情况下返回，因为它知道上游的发送者sq会在done关闭时停止尝试发送。output通过defer语句确保wg.Done在所有返回路径上被调用：

```go linenums="1"
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c or done is closed, then calls
    // wg.Done.
    output := func(c <-chan int) {
        defer wg.Done()
        for n := range c {
            select {
            case out <- n:
            case <-done:
                return
            }
        }
    }
    // ... the rest is unchanged ...
```

Similarly, `sq` can return as soon as `done` is closed. `sq` ensures its `out` channel is closed on all return paths via a `defer` statement:

sq通过defer语句确保它的输出通道在所有的返回路径上都是关闭的，同样地，sq可以在doed关闭后立即返回：

```go linenums="1"
func sq(done <-chan struct{}, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            select {
            case out <- n * n:
            case <-done:
                return
            }
        }
    }()
    return out
}
```

Here are the guidelines for pipeline construction:

以下是流水线构建的准则：

- stages close their outbound channels when all the send operations are done. 当所有的发送操作完成后，阶段会关闭其出站通道。
- stages keep receiving values from inbound channels until those channels are closed or the senders are unblocked. 阶段继续从入站通道接收值，直到这些通道被关闭或发送者被解封。

Pipelines unblock senders either by ensuring there’s enough buffer for all the values that are sent or by explicitly signalling senders when the receiver may abandon the channel.

管道通过确保有足够的缓冲区容纳所有发送的值，或者通过明确的信号通知发送者接收者可以放弃通道来解除对发送者的封锁。

## Digesting a tree 消化一棵树

Let’s consider a more realistic pipeline.

让我们考虑一个更现实的管道。

MD5 is a message-digest algorithm that’s useful as a file checksum. The command line utility `md5sum` prints digest values for a list of files.

MD5是一种消息摘要算法，作为文件校验很有用。命令行工具md5sum可以打印出一列文件的摘要值。

```shell linenums="1"
% md5sum *.go
d47c2bbc28298ca9befdfbc5d3aa4e65  bounded.go
ee869afd31f83cbb2d10ee81b2b831dc  parallel.go
b88175e65fdcbc01ac08aaf1fd9b5e96  serial.go
```

Our example program is like `md5sum` but instead takes a single directory as an argument and prints the digest values for each regular file under that directory, sorted by path name.

我们的例子程序与md5sum类似，但它将一个目录作为参数，并打印出该目录下每个常规文件的摘要值，按路径名称排序。

```shell linenums="1"
% go run serial.go .
d47c2bbc28298ca9befdfbc5d3aa4e65  bounded.go
ee869afd31f83cbb2d10ee81b2b831dc  parallel.go
b88175e65fdcbc01ac08aaf1fd9b5e96  serial.go
```

The main function of our program invokes a helper function `MD5All`, which returns a map from path name to digest value, then sorts and prints the results:

我们程序的主函数调用了一个辅助函数MD5All，它返回一个从路径名到摘要值的映射，然后对结果进行排序和打印：

```go linenums="1"
func main() {
    // Calculate the MD5 sum of all files under the specified directory,
    // then print the results sorted by path name.
    m, err := MD5All(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    var paths []string
    for path := range m {
        paths = append(paths, path)
    }
    sort.Strings(paths)
    for _, path := range paths {
        fmt.Printf("%x  %s\n", m[path], path)
    }
}
```

The `MD5All` function is the focus of our discussion. In [serial.go](https://go.dev/blog/pipelines/serial.go), the implementation uses no concurrency and simply reads and sums each file as it walks the tree.

MD5All函数是我们讨论的重点。在serial.go中，该实现没有使用并发，只是在行走树的过程中对每个文件进行读取和计算。

```go linenums="1"
// MD5All reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.
func MD5All(root string) (map[string][md5.Size]byte, error) {
    m := make(map[string][md5.Size]byte)
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.Mode().IsRegular() {
            return nil
        }
        data, err := ioutil.ReadFile(path)
        if err != nil {
            return err
        }
        m[path] = md5.Sum(data)
        return nil
    })
    if err != nil {
        return nil, err
    }
    return m, nil
}
```

## Parallel digestion 并行消化

In [parallel.go](https://go.dev/blog/pipelines/parallel.go), we split `MD5All` into a two-stage pipeline. The first stage, `sumFiles`, walks the tree, digests each file in a new goroutine, and sends the results on a channel with value type `result`:

在parallel.go中，我们将MD5All分成了一个两阶段的流水线。第一阶段，sumFiles，行走树，在一个新的goroutine中消化每个文件，并将结果发送到一个具有价值类型result的通道上：

```go linenums="1"
type result struct {
    path string
    sum  [md5.Size]byte
    err  error
}
```

`sumFiles` returns two channels: one for the `results` and another for the error returned by `filepath.Walk`. The walk function starts a new goroutine to process each regular file, then checks `done`. If `done` is closed, the walk stops immediately:

sumFiles返回两个通道：一个是结果，另一个是由filepath.Walk返回的错误。walk函数启动一个新的goroutine来处理每个常规文件，然后检查doed。如果doed是关闭的，那么walk就会立即停止：

```go linenums="1"
func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
    // For each regular file, start a goroutine that sums the file and sends
    // the result on c.  Send the result of the walk on errc.
    c := make(chan result)
    errc := make(chan error, 1)
    go func() {
        var wg sync.WaitGroup
        err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.Mode().IsRegular() {
                return nil
            }
            wg.Add(1)
            go func() {
                data, err := ioutil.ReadFile(path)
                select {
                case c <- result{path, md5.Sum(data), err}:
                case <-done:
                }
                wg.Done()
            }()
            // Abort the walk if done is closed.
            select {
            case <-done:
                return errors.New("walk canceled")
            default:
                return nil
            }
        })
        // Walk has returned, so all calls to wg.Add are done.  Start a
        // goroutine to close c once all the sends are done.
        go func() {
            wg.Wait()
            close(c)
        }()
        // No select needed here, since errc is buffered.
        errc <- err
    }()
    return c, errc
}
```

`MD5All` receives the digest values from `c`. `MD5All` returns early on error, closing `done` via a `defer`:

MD5All从c中接收摘要值。MD5All在错误时提前返回，通过defer完成关闭：

```go linenums="1"
func MD5All(root string) (map[string][md5.Size]byte, error) {
    // MD5All closes the done channel when it returns; it may do so before
    // receiving all the values from c and errc.
    done := make(chan struct{})
    defer close(done)          

    c, errc := sumFiles(done, root)

    m := make(map[string][md5.Size]byte)
    for r := range c {
        if r.err != nil {
            return nil, r.err
        }
        m[r.path] = r.sum
    }
    if err := <-errc; err != nil {
        return nil, err
    }
    return m, nil
}
```

## Bounded parallelism 有限制的并行性

The `MD5All` implementation in [parallel.go](https://go.dev/blog/pipelines/parallel.go) starts a new goroutine for each file. In a directory with many large files, this may allocate more memory than is available on the machine.

parallel.go中的MD5All实现为每个文件启动一个新的goroutine。在一个有许多大文件的目录中，这可能会分配比机器上可用的更多内存。

We can limit these allocations by bounding the number of files read in parallel. In [bounded.go](https://go.dev/blog/pipelines/bounded.go), we do this by creating a fixed number of goroutines for reading files. Our pipeline now has three stages: walk the tree, read and digest the files, and collect the digests.

我们可以通过限定并行读取的文件数量来限制这些分配。在bounded.go中，我们通过为读取文件创建固定数量的goroutines来做到这一点。我们的管道现在有三个阶段：行走树，读取和消化文件，以及收集摘要。

The first stage, `walkFiles`, emits the paths of regular files in the tree:

第一阶段，walkFiles，发出树中常规文件的路径：

```go linenums="1"
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
    paths := make(chan string)
    errc := make(chan error, 1)
    go func() {
        // Close the paths channel after Walk returns.
        defer close(paths)
        // No select needed for this send, since errc is buffered.
        errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.Mode().IsRegular() {
                return nil
            }
            select {
            case paths <- path:
            case <-done:
                return errors.New("walk canceled")
            }
            return nil
        })
    }()
    return paths, errc
}
```

The middle stage starts a fixed number of `digester` goroutines that receive file names from `paths` and send `results` on channel `c`:

中间阶段启动固定数量的消化器goroutines，从paths接收文件名并在通道c上发送结果：

```go linenums="1"
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
    for path := range paths {
        data, err := ioutil.ReadFile(path)
        select {
        case c <- result{path, md5.Sum(data), err}:
        case <-done:
            return
        }
    }
}
```

Unlike our previous examples, `digester` does not close its output channel, as multiple goroutines are sending on a shared channel. Instead, code in `MD5All` arranges for the channel to be closed when all the `digesters` are done:

与我们之前的例子不同，消化器并没有关闭它的输出通道，因为多个goroutines在一个共享通道上发送。相反，MD5All中的代码安排在所有消化器完成后关闭该通道：

```go linenums="1"
    // Start a fixed number of goroutines to read and digest files.
    c := make(chan result)
    var wg sync.WaitGroup
    const numDigesters = 20
    wg.Add(numDigesters)
    for i := 0; i < numDigesters; i++ {
        go func() {
            digester(done, paths, c)
            wg.Done()
        }()
    }
    go func() {
        wg.Wait()
        close(c)
    }()
```

We could instead have each digester create and return its own output channel, but then we would need additional goroutines to fan-in the results.

我们可以让每个消化器创建并返回自己的输出通道，但这样我们就需要额外的goroutines来对结果进行fan-in。

The final stage receives all the `results` from `c` then checks the error from `errc`. This check cannot happen any earlier, since before this point, `walkFiles` may block sending values downstream:

最后阶段从c接收所有的结果，然后从errc检查错误。这个检查不能再早了，因为在这之前，walkFiles可能会阻止向下游发送值：

```go linenums="1"
    m := make(map[string][md5.Size]byte)
    for r := range c {
        if r.err != nil {
            return nil, r.err
        }
        m[r.path] = r.sum
    }
    // Check whether the Walk failed.
    if err := <-errc; err != nil {
        return nil, err
    }
    return m, nil
}
```

## Conclusion 结论

This article has presented techniques for constructing streaming data pipelines in Go. Dealing with failures in such pipelines is tricky, since each stage in the pipeline may block attempting to send values downstream, and the downstream stages may no longer care about the incoming data. We showed how closing a channel can broadcast a "done" signal to all the goroutines started by a pipeline and defined guidelines for constructing pipelines correctly.

本文介绍了在Go中构建流式数据管道的技术。处理这种管道中的故障是很棘手的，因为管道中的每个阶段都可能阻断向下游发送数值的尝试，而且下游阶段可能不再关心传入的数据。我们展示了关闭一个通道如何向管道启动的所有goroutine广播一个 "完成 "信号，并定义了正确构建管道的准则。

Further reading:

进一步阅读：

- [Go Concurrency Patterns](https://go.dev/talks/2012/concurrency.slide#1) ([video](https://www.youtube.com/watch?v=f6kdp27TYZs)) presents the basics of Go’s concurrency primitives and several ways to apply them. Go并发模式（视频）介绍了Go并发基元的基础知识以及应用它们的几种方法。
- [Advanced Go Concurrency Patterns](https://blog.golang.org/advanced-go-concurrency-patterns) ([video](http://www.youtube.com/watch?v=QDDwwePbDtw)) covers more complex uses of Go’s primitives, especially `select`. 高级 Go 并发模式（视频）涵盖了 Go 基元更复杂的应用，特别是选择。
- Douglas McIlroy’s paper [Squinting at Power Series](https://swtch.com/~rsc/thread/squint.pdf) shows how Go-like concurrency provides elegant support for complex calculations. Douglas McIlroy的论文《Squinting at Power Series》展示了类似Go的并发性如何为复杂的计算提供优雅的支持。
