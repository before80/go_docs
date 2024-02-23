+++
title = "Go 并发模式：Pipelines和取消"
weight = 13
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go Concurrency Patterns: Pipelines and cancellation - Go 并发模式：管道和取消

> 原文：[https://go.dev/blog/pipelines](https://go.dev/blog/pipelines)

Sameer Ajmani
13 March 2014

2014年3月13日

## 简介 Introduction 

Go’s concurrency primitives make it easy to construct streaming data pipelines that make efficient use of I/O and multiple CPUs. This article presents examples of such pipelines, highlights subtleties that arise when operations fail, and introduces techniques for dealing with failures cleanly.

​	Go 的并发原语使得构建流式数据管道变得容易，这些管道能够高效地利用 I/O 和多个 CPU。本文介绍了这种管道的示例，强调了操作失败时出现的细微之处，并介绍了处理失败的技巧。

## 什么是管道？- What is a pipeline? 

There’s no formal definition of a pipeline in Go; it’s just one of many kinds of concurrent programs. Informally, a pipeline is a series of *stages* connected by channels, where each stage is a group of goroutines running the same function. In each stage, the goroutines

​	在 Go 中，没有管道的正式定义；它只是众多并发程序中的一种。非正式地讲，管道是由通道连接的一系列*阶段*，其中每个阶段都是运行相同函数的一组 goroutine。在每个阶段中，goroutine

- receive values from *upstream* via *inbound* channels 
- 通过*入站* 通道从*上游* 接收值
- perform some function on that data, usually producing new values 
- 对该数据执行某些函数，通常生成新的值
- send values *downstream* via *outbound* channels 
- 通过*出站* 通道将值发送到 *下游*

Each stage has any number of inbound and outbound channels, except the first and last stages, which have only outbound or inbound channels, respectively. The first stage is sometimes called the *source* or *producer*; the last stage, the *sink* or *consumer*.

​	每个阶段有任意数量的入站和出站通道，除了第一个和最后一个阶段，它们只有出站或入站通道。第一个阶段有时称为*源* 或*生产者*；最后一个阶段是 *接收者* 或*消费者* 。

We’ll begin with a simple example pipeline to explain the ideas and techniques. Later, we’ll present a more realistic example.

​	我们将从一个简单的示例管道开始，以解释这些思想和技术。随后，我们将呈现一个更实际的示例。

## 数字的平方 - Squaring numbers

Consider a pipeline with three stages.

​	考虑一个具有三个阶段的管道。

The first stage, `gen`, is a function that converts a list of integers to a channel that emits the integers in the list. The `gen` function starts a goroutine that sends the integers on the channel and closes the channel when all the values have been sent:

​	第一个阶段，`gen`，是一个将整数列表转换为发出列表中整数的通道的函数。`gen` 函数启动一个 goroutine，将整数发送到通道上，并在所有值都被发送后关闭通道：

```go
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

​	第二个阶段，`sq`，从通道接收整数，并返回一个发出每个接收到的整数的平方的通道。在入站通道关闭并且此阶段已将所有值发送到下游之后，它关闭出站通道：

```go
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

​	`main` 函数设置管道并运行最终阶段：它从第二个阶段接收值并打印每个值，直到通道关闭：

```go
func main() {
    // Set up the pipeline.
    // 设置管道并消耗输出。
    c := gen(2, 3)
    out := sq(c)

    // Consume the output.
    // 消耗输出。
    fmt.Println(<-out) // 4
    fmt.Println(<-out) // 9
}
```

Since `sq` has the same type for its inbound and outbound channels, we can compose it any number of times. We can also rewrite `main` as a range loop, like the other stages:

​	由于 `sq` 的入站和出站通道具有相同的类型，我们可以多次组合它。我们还可以将 `main` 重写为一个范围循环，就像其他阶段一样：

```go
func main() {
    // Set up the pipeline and consume the output.
    // 设置管道并消耗输出。
    for n := range sq(sq(gen(2, 3))) {
        fmt.Println(n) // 16 then 81
    }
}
```

## 扇出、扇入 - Fan-out, fan-in 

Multiple functions can read from the same channel until that channel is closed; this is called *fan-out*. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.

​	多个函数可以从同一个通道读取，直到该通道关闭；这称为*扇出*。这为将工作分配给一组工作者以并行化 CPU 使用和 I/O 提供了一种方法。

A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that’s closed when all the inputs are closed. This is called *fan-in*.

​	函数可以从多个输入读取并在所有输入关闭时进行多路复用，将输入通道复用到单个通道上，该通道在所有输入关闭时关闭。这称为*扇入*。

We can change our pipeline to run two instances of `sq`, each reading from the same input channel. We introduce a new function, *merge*, to fan in the results:

​	我们可以改变我们的管道，使其运行两个 `sq` 实例，每个实例都从相同的输入通道读取。我们引入了一个名为*merge*的新函数，用于扇入结果：

```go
func main() {
    in := gen(2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    // 将 sq 工作分布到两个同时从 in 读取的 goroutine 中。
    c1 := sq(in)
    c2 := sq(in)

    // Consume the merged output from c1 and c2.
    // 消耗来自 c1 和 c2 的合并输出。
    for n := range merge(c1, c2) {
        fmt.Println(n) // 4 then 9, or 9 then 4
    }
}
```

The `merge` function converts a list of channels to a single channel by starting a goroutine for each inbound channel that copies the values to the sole outbound channel. Once all the `output` goroutines have been started, `merge` starts one more goroutine to close the outbound channel after all sends on that channel are done.

​	`merge` 函数通过为每个入站通道启动一个 goroutine 来将通道列表转换为单个通道，这些 goroutine 将值复制到唯一的出站通道。一旦启动了所有 `output` goroutine，`merge` 就会再启动一个 goroutine，以在该通道上的所有发送完成后关闭出站通道。

Sends on a closed channel panic, so it’s important to ensure all sends are done before calling close. The [`sync.WaitGroup`](https://go.dev/pkg/sync/#WaitGroup) type provides a simple way to arrange this synchronization:

​	在关闭的通道上发送会导致 panic，因此在调用 close 之前要确保所有发送都已完成。[`sync.WaitGroup`](https://go.dev/pkg/sync/#WaitGroup) 类型提供了一种简单的方式来安排这种同步：

```go
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    // 为 cs 中的每个输入通道启动一个输出 goroutine。
    // output 从 c 复制值到 out，直到 c 被关闭，然后调用 wg.Done。
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
    // 在所有输出 goroutine 完成后启动一个 goroutine 来关闭 out。
    // 这必须在 wg.Add 调用之后启动。
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```

## 突然停止 - Stopping short 

There is a pattern to our pipeline functions:

​	我们的管道函数有一个模式：

- stages close their outbound channels when all the send operations are done. 
- 阶段在所有发送操作完成后关闭其出站通道。
- stages keep receiving values from inbound channels until those channels are closed. 
- 阶段继续从入站通道接收值，直到这些通道关闭。

This pattern allows each receiving stage to be written as a `range` loop and ensures that all goroutines exit once all values have been successfully sent downstream.

​	此模式允许每个接收阶段都编写为 `range` 循环，并确保一旦所有值都成功发送到下游，所有 goroutine 都会退出。

But in real pipelines, stages don’t always receive all the inbound values. Sometimes this is by design: the receiver may only need a subset of values to make progress. More often, a stage exits early because an inbound value represents an error in an earlier stage. In either case the receiver should not have to wait for the remaining values to arrive, and we want earlier stages to stop producing values that later stages don’t need.

​	但是在实际的管道中，阶段并不总是接收所有入站值。有时这是有意为之的：接收方可能只需要子集值来取得进展。更常见的是，阶段提前退出，因为入站值表示前面阶段的错误。无论哪种情况，接收方都不应等待剩余的值到达，我们希望较早的阶段停止生成后续阶段不需要的值。

In our example pipeline, if a stage fails to consume all the inbound values, the goroutines attempting to send those values will block indefinitely:

​	在我们的示例管道中，如果阶段未能消耗所有入站值，尝试发送这些值的 goroutine 将无限期地阻塞：

```go
    // Consume the first value from the output.
	// 从输出中消费第一个值。
    out := merge(c1, c2)
    fmt.Println(<-out) // 4 or 9
    return
    // Since we didn't receive the second value from out,
    // one of the output goroutines is hung attempting to send it.
	// 由于我们没有从 out 接收第二个值，
    // 一个输出 goroutine 因试图发送第二个值而被阻塞。
}
```

This is a resource leak: goroutines consume memory and runtime resources, and heap references in goroutine stacks keep data from being garbage collected. Goroutines are not garbage collected; they must exit on their own.

​	这是一个资源泄漏：goroutine 占用内存和运行时资源，而 goroutine 堆栈中的堆引用会阻止数据被垃圾回收。Goroutine 不会被垃圾回收；它们必须自己退出。

We need to arrange for the upstream stages of our pipeline to exit even when the downstream stages fail to receive all the inbound values. One way to do this is to change the outbound channels to have a buffer. A buffer can hold a fixed number of values; send operations complete immediately if there’s room in the buffer:

​	我们需要安排管道的上游阶段在下游阶段未能接收所有入站值时也退出。一种方法是将出站通道更改为具有缓冲区。缓冲区可以容纳固定数量的值；如果缓冲区中有空间，发送操作会立即完成：

```go
c := make(chan int, 2) // buffer size 2 缓冲区大小为 2
c <- 1  // succeeds immediately 立即成功
c <- 2  // succeeds immediately 立即成功
c <- 3  // blocks until another goroutine does <-c and receives 1 阻塞，直到另一个 goroutine 执行 <-c 并接收 1
```

When the number of values to be sent is known at channel creation time, a buffer can simplify the code. For example, we can rewrite `gen` to copy the list of integers into a buffered channel and avoid creating a new goroutine:

​	当在通道创建时知道要发送的值的数量时，缓冲区可以简化代码。例如，我们可以重写 `gen` 将整数列表复制到带缓冲区的通道中，并避免创建新的 goroutine：

```go
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

​	回到我们管道中的阻塞 goroutine，我们可能会考虑为 `merge` 返回的出站通道添加一个缓冲区：

```go
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int, 1) // enough space for the unread inputs 有足够的空间来存储未读取的输入
    // ... the rest is unchanged ...  ... 其余部分不变 ...
```

While this fixes the blocked goroutine in this program, this is bad code. The choice of buffer size of 1 here depends on knowing the number of values `merge` will receive and the number of values downstream stages will consume. This is fragile: if we pass an additional value to `gen`, or if the downstream stage reads any fewer values, we will again have blocked goroutines.

​	虽然这修复了此程序中的阻塞 goroutine，但这是不好的代码。这里缓冲区大小为 1 的选择取决于知道 `merge` 将接收的值的数量以及下游阶段将消耗的值的数量。这是脆弱的：如果我们向 `gen` 传递了额外的值，或者如果下游阶段读取的值更少，我们将再次有阻塞的 goroutine。

Instead, we need to provide a way for downstream stages to indicate to the senders that they will stop accepting input.

​	相反，我们需要一种方法来告诉下游阶段，它们将停止接受输入。

## 显式取消 - Explicit cancellation

When `main` decides to exit without receiving all the values from `out`, it must tell the goroutines in the upstream stages to abandon the values they’re trying to send. It does so by sending values on a channel called `done`. It sends two values since there are potentially two blocked senders:

​	当 `main` 决定退出而不接收所有来自 `out` 的值时，它必须告诉上游阶段的 goroutine 放弃它们正在尝试发送的值。它通过在一个称为 `done` 的通道上发送值来实现。它发送两个值，因为可能有两个被阻塞的发送方：

```go
func main() {
    in := gen(2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    // 将 sq 工作分布到两个从 in 读取的 goroutine 中。
    c1 := sq(in)
    c2 := sq(in)

    // Consume the first value from output.
    // 从输出中消费第一个值。
    done := make(chan struct{}, 2)
    out := merge(done, c1, c2)
    fmt.Println(<-out) // 4 or 9

    // Tell the remaining senders we're leaving.
    // 告诉剩余的发送方我们要离开。
    done <- struct{}{}
    done <- struct{}{}
}
```

The sending goroutines replace their send operation with a `select` statement that proceeds either when the send on `out` happens or when they receive a value from `done`. The value type of `done` is the empty struct because the value doesn’t matter: it is the receive event that indicates the send on `out` should be abandoned. The `output` goroutines continue looping on their inbound channel, `c`, so the upstream stages are not blocked. (We’ll discuss in a moment how to allow this loop to return early.)

​	发送方的 goroutine 用一个 `select` 语句替换其发送操作，当在 `out` 上的发送发生或从 `done` 接收到值时，该语句继续执行。`done` 的值类型是空结构，因为值不重要：它是接收事件，指示在 `out` 上的发送应该被放弃。`output` goroutine 在其入站通道 `c` 上继续循环，因此上游阶段不会被阻塞。 （我们马上将讨论如何让此循环提前返回。）

```go
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed or it receives a value
    // from done, then output calls wg.Done.
    // 为每个输入通道在 cs 中启动一个输出 goroutine。
    // output 从 c 复制值到 out，直到 c 关闭或从 done 接收到值，然后 output 调用 wg.Done。
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
    // ... 其余部分不变 ...
```

This approach has a problem: *each* downstream receiver needs to know the number of potentially blocked upstream senders and arrange to signal those senders on early return. Keeping track of these counts is tedious and error-prone.

​	这种方法存在一个问题：每个下游接收方都需要知道可能被阻塞的上游发送方的数量，并在提前返回时安排向这些发送方发出信号。跟踪这些计数是繁琐且容易出错的。

We need a way to tell an unknown and unbounded number of goroutines to stop sending their values downstream. In Go, we can do this by closing a channel, because [a receive operation on a closed channel can always proceed immediately, yielding the element type’s zero value.](https://go.dev/ref/spec#Receive_operator)

​	我们需要一种方法来告诉未知且无限的数量的 goroutine 停止向下游发送其值。在 Go 中，我们可以通过关闭通道来做到这一点，因为[关闭的通道上的接收操作始终可以立即进行，产生元素类型的零值。](https://go.dev/ref/spec#Receive_operator)

This means that `main` can unblock all the senders simply by closing the `done` channel. This close is effectively a broadcast signal to the senders. We extend *each* of our pipeline functions to accept `done` as a parameter and arrange for the close to happen via a `defer` statement, so that all return paths from `main` will signal the pipeline stages to exit.

​	这意味着 `main` 可以通过关闭 `done` 通道来解除所有发送方的阻塞。这个关闭实际上是对发送方的广播信号。我们通过将 `done` 作为参数扩展我们的每个管道函数，通过 `defer` 语句安排关闭，以便从 `main` 的所有返回路径都会发出信号，以使管道阶段退出。

```go
func main() {
    // Set up a done channel that's shared by the whole pipeline,
    // and close that channel when this pipeline exits, as a signal
    // for all the goroutines we started to exit.
    // 设置一个 done 通道，整个管道共享它，
    // 并在此管道退出时关闭该通道，作为我们启动的所有 goroutine 退出的信号。
    done := make(chan struct{})
    defer close(done)          

    in := gen(done, 2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    // 将 sq 工作分布到两个从 in 读取的 goroutine 中。
    c1 := sq(done, in)
    c2 := sq(done, in)

    // Consume the first value from output.
    // 从输出中消费第一个值。
    out := merge(done, c1, c2)
    fmt.Println(<-out) // 4 or 9

    // done will be closed by the deferred call.  
    // 通过延迟调用将会关闭 done。
}
```

Each of our pipeline stages is now free to return as soon as `done` is closed. The `output` routine in `merge` can return without draining its inbound channel, since it knows the upstream sender, `sq`, will stop attempting to send when `done` is closed. `output` ensures `wg.Done` is called on all return paths via a `defer` statement:

​	我们的每个管道阶段现在可以在 `done` 关闭时立即返回。在 `merge` 中，`output` 例程可以在未能完全处理其入站通道的情况下返回，因为它知道上游发送方 `sq` 会在 `done` 关闭时停止尝试发送。通过 `defer` 语句，`output` 确保在所有返回路径上都调用了 `wg.Done`：

```go
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c or done is closed, then calls
    // wg.Done.
    // 为每个输入通道在 cs 中启动一个输出 goroutine。
    // output 从 c 复制值到 out，直到 c 或 done 关闭，然后调用 wg.Done。
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
    // ... 其余部分不变 ...
```

Similarly, `sq` can return as soon as `done` is closed. `sq` ensures its `out` channel is closed on all return paths via a `defer` statement:

​	同样地，当 `done` 关闭时，`sq` 也可以立即返回。通过 `defer` 语句，`sq` 确保其 `out` 通道在所有返回路径上都被关闭：

```go
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

​	以下是管道构建的准则：

- stages close their outbound channels when all the send operations are done. 
- 阶段在所有发送操作完成后关闭其出站通道。
- stages keep receiving values from inbound channels until those channels are closed or the senders are unblocked. 
- 阶段从入站通道接收值，直到这些通道关闭或发送方被解除阻塞。

Pipelines unblock senders either by ensuring there’s enough buffer for all the values that are sent or by explicitly signalling senders when the receiver may abandon the channel.

​	通过确保有足够的缓冲区来发送所有值，或者通过显式地在接收者可能放弃通道时向发送者发送信号，管道可以解除发送者的阻塞。

## 处理树的摘要 - Digesting a tree

Let’s consider a more realistic pipeline.

​	让我们考虑一个更现实的管道示例。

MD5 is a message-digest algorithm that’s useful as a file checksum. The command line utility `md5sum` prints digest values for a list of files.

​	MD5是一种用作文件校验和的消息摘要算法。命令行实用程序 `md5sum` 会为一组文件打印摘要值。

```shell linenums="1"
% md5sum *.go
d47c2bbc28298ca9befdfbc5d3aa4e65  bounded.go
ee869afd31f83cbb2d10ee81b2b831dc  parallel.go
b88175e65fdcbc01ac08aaf1fd9b5e96  serial.go
```

Our example program is like `md5sum` but instead takes a single directory as an argument and prints the digest values for each regular file under that directory, sorted by path name.

​	我们的示例程序类似于 `md5sum`，但它接受一个目录作为参数，并为该目录下的每个常规文件打印路径名称排序后的摘要值。

```shell linenums="1"
% go run serial.go .
d47c2bbc28298ca9befdfbc5d3aa4e65  bounded.go
ee869afd31f83cbb2d10ee81b2b831dc  parallel.go
b88175e65fdcbc01ac08aaf1fd9b5e96  serial.go
```

The main function of our program invokes a helper function `MD5All`, which returns a map from path name to digest value, then sorts and prints the results:

​	我们的程序的主函数调用名为 `MD5All` 的帮助函数，该函数返回从路径名称到摘要值的映射，然后对结果进行排序并打印：

```go
func main() {
    // Calculate the MD5 sum of all files under the specified directory,
    // then print the results sorted by path name.
    // 计算指定目录下所有文件的 MD5 校验和，
    // 然后按路径名称排序并打印结果。
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

​	`MD5All` 函数是我们讨论的重点。在 [serial.go](https://go.dev/blog/pipelines/serial.go) 中的实现没有使用并发，只是在遍历树的过程中读取和计算每个文件的摘要。

```go
// MD5All reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.
// MD5All 会读取根目录下的所有文件，并返回从文件路径到文件内容的 MD5 摘要的映射。
// 如果目录遍历失败或任何读取操作失败，MD5All 会返回一个错误。
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

## 并行摘要 - Parallel digestion

In [parallel.go](https://go.dev/blog/pipelines/parallel.go), we split `MD5All` into a two-stage pipeline. The first stage, `sumFiles`, walks the tree, digests each file in a new goroutine, and sends the results on a channel with value type `result`:

​	在 [parallel.go](https://go.dev/blog/pipelines/parallel.go) 中，我们将 `MD5All` 拆分为一个两阶段的管道。第一阶段 `sumFiles` 遍历树，对每个文件进行摘要计算并将结果发送到一个值类型为 `result` 的通道中：

```go
type result struct {
    path string
    sum  [md5.Size]byte
    err  error
}
```

`sumFiles` returns two channels: one for the `results` and another for the error returned by `filepath.Walk`. The walk function starts a new goroutine to process each regular file, then checks `done`. If `done` is closed, the walk stops immediately:

​	`sumFiles` 返回两个通道：一个用于 `results`，另一个用于 `filepath.Walk` 返回的错误。遍历函数会启动一个新的 goroutine 来处理每个常规文件，然后检查 `done`。如果 `done` 被关闭，遍历会立即停止：

```go
func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
    // For each regular file, start a goroutine that sums the file and sends
    // the result on c.  Send the result of the walk on errc.
    // 为每个常规文件，启动一个 goroutine 计算文件的摘要并将结果发送到 c。
    // 将遍历结果发送到 errc。
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
            // 如果 done 被关闭，终止遍历。
            select {
            case <-done:
                return errors.New("walk canceled")
            default:
                return nil
            }
        })
        // Walk has returned, so all calls to wg.Add are done.  Start a
        // goroutine to close c once all the sends are done.
         // 遍历已返回，所以所有对 wg.Add 的调用都已完成。
        // 启动一个 goroutine 在所有发送完成后关闭 c。
        go func() {
            wg.Wait()
            close(c)
        }()
        // No select needed here, since errc is buffered.
        // 这里不需要 select，因为 errc 是有缓冲的。
        errc <- err
    }()
    return c, errc
}
```

`MD5All` receives the digest values from `c`. `MD5All` returns early on error, closing `done` via a `defer`:

​	`MD5All` 从 `c` 中接收摘要值。`MD5All` 会在出现错误时提前返回，并通过 `defer` 关闭 `done`：

```go
func MD5All(root string) (map[string][md5.Size]byte, error) {
    // MD5All closes the done channel when it returns; it may do so before
    // receiving all the values from c and errc.
     // 当 MD5All 返回时，它会关闭 done 通道；
    // 它可能在接收到所有值之前就这样做。
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

## 有限并行性 - Bounded parallelism

The `MD5All` implementation in [parallel.go](https://go.dev/blog/pipelines/parallel.go) starts a new goroutine for each file. In a directory with many large files, this may allocate more memory than is available on the machine.

​	在 [parallel.go](https://go.dev/blog/pipelines/parallel.go) 中，`MD5All` 的实现为每个文件启动了一个新的 goroutine。在包含许多大文件的目录中，这可能会分配超过机器可用内存的内存。

We can limit these allocations by bounding the number of files read in parallel. In [bounded.go](https://go.dev/blog/pipelines/bounded.go), we do this by creating a fixed number of goroutines for reading files. Our pipeline now has three stages: walk the tree, read and digest the files, and collect the digests.

​	我们可以通过限制并行读取的文件数量来限制这些分配。在 [bounded.go](https://go.dev/blog/pipelines/bounded.go) 中，我们通过为读取文件创建一定数量的固定数量的 goroutines 来实现。我们的管道现在有三个阶段：遍历树，读取和计算摘要，以及收集摘要。

The first stage, `walkFiles`, emits the paths of regular files in the tree:

​	第一阶段 `walkFiles` 发出树中常规文件的路径：

```go
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
    paths := make(chan string)
    errc := make(chan error, 1)
    go func() {
        // Close the paths channel after Walk returns.
        // 在 Walk 返回后关闭 paths 通道。
        defer close(paths)
        // No select needed for this send, since errc is buffered.
        // 对于这次发送，不需要 select，因为 errc 是有缓冲的。
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

​	中间阶段启动了固定数量的 `digester` goroutines，这些 goroutines 从 `paths` 接收文件名，并在通道 `c` 上发送 `results`：

```go
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

​	与之前的示例不同，`digester` 不会关闭其输出通道，因为多个 goroutines 正在共享一个通道发送数据。相反，在 `MD5All` 中的代码会安排在所有 `digesters` 完成时关闭通道：

```go
    // Start a fixed number of goroutines to read and digest files.
	// 启动固定数量的 goroutines 来读取和计算摘要。
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

​	我们也可以让每个 `digester` 创建并返回自己的输出通道，但这样我们就需要额外的 goroutines 来汇集（fan-in）结果。

The final stage receives all the `results` from `c` then checks the error from `errc`. This check cannot happen any earlier, since before this point, `walkFiles` may block sending values downstream:

​	最终阶段从 `c` 接收所有的 `results`，然后检查来自 `errc` 的错误。在此之前无法进行此检查，因为在此之前，`walkFiles` 可能会阻止将值发送到下游：

```go
    m := make(map[string][md5.Size]byte)
    for r := range c {
        if r.err != nil {
            return nil, r.err
        }
        m[r.path] = r.sum
    }
    // Check whether the Walk failed.
	// 检查 Walk 是否失败。
    if err := <-errc; err != nil {
        return nil, err
    }
    return m, nil
}
```

## 结论 Conclusion

This article has presented techniques for constructing streaming data pipelines in Go. Dealing with failures in such pipelines is tricky, since each stage in the pipeline may block attempting to send values downstream, and the downstream stages may no longer care about the incoming data. We showed how closing a channel can broadcast a "done" signal to all the goroutines started by a pipeline and defined guidelines for constructing pipelines correctly.

​	本文介绍了在 Go 中构建流数据管道的技术。在这种管道中处理故障是棘手的，因为管道中的每个阶段可能会阻止尝试将值发送到下游，而下游阶段可能不再关心传入的数据。我们展示了如何通过关闭通道来向由管道启动的所有 goroutines 广播“完成”信号，并定义了构建管道的正确指南。

Further reading:

​	进一步阅读：

- [Go Concurrency Patterns](https://go.dev/talks/2012/concurrency.slide#1) ([video](https://www.youtube.com/watch?v=f6kdp27TYZs)) presents the basics of Go’s concurrency primitives and several ways to apply them. 
- [Go 并发模式](https://go.dev/talks/2012/concurrency.slide#1)（[视频](https://www.youtube.com/watch?v=f6kdp27TYZs)）介绍了 Go 并发原语的基础知识以及几种应用方法。
- [Advanced Go Concurrency Patterns](https://blog.golang.org/advanced-go-concurrency-patterns) ([video](http://www.youtube.com/watch?v=QDDwwePbDtw)) covers more complex uses of Go’s primitives, especially `select`. 
- [高级 Go 并发模式](https://blog.golang.org/advanced-go-concurrency-patterns)（[视频](http://www.youtube.com/watch?v=QDDwwePbDtw)）介绍了更复杂的 Go 原语使用方式，特别是 `select`。
- Douglas McIlroy’s paper [Squinting at Power Series](https://swtch.com/~rsc/thread/squint.pdf) shows how Go-like concurrency provides elegant support for complex calculations. 
- Douglas McIlroy 的论文 [Squinting at Power Series](https://swtch.com/~rsc/thread/squint.pdf) 展示了类似 Go 的并发是如何优雅地支持复杂计算的。
