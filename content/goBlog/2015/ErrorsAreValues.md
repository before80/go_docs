+++
title = "errors 是值"
weight = 11
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Errors are values - errors 是值

> 原文：[https://go.dev/blog/errors-are-values](https://go.dev/blog/errors-are-values)

Rob Pike
12 January 2015

A common point of discussion among Go programmers, especially those new to the language, is how to handle errors. The conversation often turns into a lament at the number of times the sequence

​	Go程序员，尤其是那些刚接触该语言的程序员，经常讨论的一个问题是如何处理错误。对话往往变成了对下面序列的多次出现而哀叹连连。

```go
if err != nil {
    return err
}
```

shows up. We recently scanned all the open source projects we could find and discovered that this snippet occurs only once per page or two, less often than some would have you believe. Still, if the perception persists that one must type

​	我们最近扫描了所有我们能找到的开源项目，发现此片段只在每页出现一到两次，比一些人认为的要少。尽管如此，如果人们仍然认为必须一直输入

```go
if err != nil
```

all the time, something must be wrong, and the obvious target is Go itself.

一定有什么地方出了问题，而明显的目标就是Go本身。

This is unfortunate, misleading, and easily corrected. Perhaps what is happening is that programmers new to Go ask, "How does one handle errors?", learn this pattern, and stop there. In other languages, one might use a try-catch block or other such mechanism to handle errors. Therefore, the programmer thinks, when I would have used a try-catch in my old language, I will just type `if` `err` `!=` `nil` in Go. Over time the Go code collects many such snippets, and the result feels clumsy.

这是不幸的，误导性的，而且很容易被纠正。也许现在的情况是，刚接触Go的程序员会问："如何处理错误？"，学习这种模式，然后停在那里。在其他语言中，人们可能会使用try-catch块或其他类似机制来处理错误。因此，程序员认为，当我在以前的语言中会使用try-catch时，我在Go中只需输入if err != nil。随着时间的推移，Go代码中收集了许多这样的片段，结果感觉很笨拙。

Regardless of whether this explanation fits, it is clear that these Go programmers miss a fundamental point about errors: *Errors are values.*

不管这种解释是否合适，很明显，这些Go程序员错过了关于错误的一个基本点。错误就是值。

Values can be programmed, and since errors are values, errors can be programmed.

值可以被编程，既然错误是值，那么错误也可以被编程。

Of course a common statement involving an error value is to test whether it is nil, but there are countless other things one can do with an error value, and application of some of those other things can make your program better, eliminating much of the boilerplate that arises if every error is checked with a rote if statement.

当然，涉及错误值的常见语句是测试它是否为nil，但还有无数其他的事情可以用错误值来做，应用其中一些其他的事情可以使您的程序变得更好，消除了如果每个错误都用死记硬背的if语句来检查而产生的许多模板。

Here’s a simple example from the `bufio` package’s [`Scanner`](https://go.dev/pkg/bufio/#Scanner) type. Its [`Scan`](https://go.dev/pkg/bufio/#Scanner.Scan) method performs the underlying I/O, which can of course lead to an error. Yet the `Scan` method does not expose an error at all. Instead, it returns a boolean, and a separate method, to be run at the end of the scan, reports whether an error occurred. Client code looks like this:

下面是一个来自`bufio`包的[Scanner](https://go.dev/pkg/bufio/#Scanner)类型的简单例子。它的[Scan](https://go.dev/pkg/bufio/#Scanner.Scan)方法执行了底层的I/O，这当然会导致错误的发生。然而，扫描方法根本就没有暴露出错误。相反，它返回一个布尔值，并在扫描结束时运行一个单独的方法，报告是否发生错误。客户端代码看起来像这样：

```go
scanner := bufio.NewScanner(input)
for scanner.Scan() {
    token := scanner.Text()
    // process token
}
if err := scanner.Err(); err != nil {
    // process the error
}
```

Sure, there is a nil check for an error, but it appears and executes only once. The `Scan` method could instead have been defined as

当然，有一个错误的nil检查，但它只出现和执行一次。扫描方法可以被定义为

```go
func (s *Scanner) Scan() (token []byte, error)
```

and then the example user code might be (depending on how the token is retrieved),

然后，用户代码的例子可能是（取决于如何检索令牌）。

```go
scanner := bufio.NewScanner(input)
for {
    token, err := scanner.Scan()
    if err != nil {
        return err // or maybe break
    }
    // process token
}
```

This isn’t very different, but there is one important distinction. In this code, the client must check for an error on every iteration, but in the real `Scanner` API, the error handling is abstracted away from the key API element, which is iterating over tokens. With the real API, the client’s code therefore feels more natural: loop until done, then worry about errors. Error handling does not obscure the flow of control.

这并没有什么不同，但有一个重要的区别。在这段代码中，客户端必须在每次迭代中检查错误，但在真正的Scanner API中，错误处理被抽象出关键的API元素，即对令牌进行迭代。因此，在真正的API中，客户的代码感觉更自然：循环直到完成，然后担心错误。错误处理不会掩盖控制的流程。

Under the covers what’s happening, of course, is that as soon as `Scan` encounters an I/O error, it records it and returns `false`. A separate method, [`Err`](https://go.dev/pkg/bufio/#Scanner.Err), reports the error value when the client asks. Trivial though this is, it’s not the same as putting

当然，掩盖起来的是，一旦Scan遇到I/O错误，它就会记录下来并返回false。一个单独的方法，Err，在客户询问时报告错误值。尽管这很微不足道，但这与把

```go
if err != nil
```

everywhere or asking the client to check for an error after every token. It’s programming with error values. Simple programming, yes, but programming nonetheless.

或要求客户端在每个标记后检查错误。这是用错误值编程。简单的编程，是的，但仍然是编程。

It’s worth stressing that whatever the design, it’s critical that the program check the errors however they are exposed. The discussion here is not about how to avoid checking errors, it’s about using the language to handle errors with grace.

值得强调的是，不管是什么设计，关键是程序要检查错误，不管它们如何暴露。这里的讨论不是关于如何避免检查错误，而是关于使用语言来优雅地处理错误。

The topic of repetitive error-checking code arose when I attended the autumn 2014 GoCon in Tokyo. An enthusiastic gopher, who goes by [`@jxck_`](https://twitter.com/jxck_) on Twitter, echoed the familiar lament about error checking. He had some code that looked schematically like this:

当我参加东京的2014年秋季GoCon时，出现了重复检查错误代码的话题。一位热心的地鼠，在Twitter上的名字是@jxck_，回应了人们熟悉的关于错误检查的哀叹。他有一些代码，从原理上看是这样的：

```go
_, err = fd.Write(p0[a:b])
if err != nil {
    return err
}
_, err = fd.Write(p1[c:d])
if err != nil {
    return err
}
_, err = fd.Write(p2[e:f])
if err != nil {
    return err
}
// and so on
```

It is very repetitive. In the real code, which was longer, there is more going on so it’s not easy to just refactor this using a helper function, but in this idealized form, a function literal closing over the error variable would help:

这是很重复的。在真正的代码中，这段代码比较长，有更多的事情要做，所以不容易只是用一个辅助函数来重构这段代码，但是在这种理想化的形式中，在错误变量上关闭一个函数字面会有帮助：

```go
var err error
write := func(buf []byte) {
    if err != nil {
        return
    }
    _, err = w.Write(buf)
}
write(p0[a:b])
write(p1[c:d])
write(p2[e:f])
// and so on
if err != nil {
    return err
}
```

This pattern works well, but requires a closure in each function doing the writes; a separate helper function is clumsier to use because the `err` variable needs to be maintained across calls (try it).

这种模式效果很好，但需要在每个进行写操作的函数中都有一个闭包；单独的辅助函数使用起来更笨拙，因为err变量需要在不同的调用中进行维护（试试看）。

We can make this cleaner, more general, and reusable by borrowing the idea from the `Scan` method above. I mentioned this technique in our discussion but `@jxck_` didn’t see how to apply it. After a long exchange, hampered somewhat by a language barrier, I asked if I could just borrow his laptop and show him by typing some code.

我们可以通过借用上面的扫描方法的想法来使这个方法更简洁、更通用、更可重复使用。我在我们的讨论中提到了这个技术，但是@jxck_没有看到如何应用它。经过长时间的交流，在语言不通的情况下，我问能不能借他的笔记本，打一些代码给他看。

I defined an object called an `errWriter`, something like this:

我定义了一个叫做errWriter的对象，大概是这样的：

```go
type errWriter struct {
    w   io.Writer
    err error
}
```

and gave it one method, `write.` It doesn’t need to have the standard `Write` signature, and it’s lower-cased in part to highlight the distinction. The `write` method calls the `Write` method of the underlying `Writer` and records the first error for future reference:

并给了它一个方法，写。它不需要有标准的Write签名，它的小写部分是为了突出这种区别。Write方法调用底层Writer的Write方法，并记录第一个错误供将来参考：

```go
func (ew *errWriter) write(buf []byte) {
    if ew.err != nil {
        return
    }
    _, ew.err = ew.w.Write(buf)
}
```

As soon as an error occurs, the `write` method becomes a no-op but the error value is saved.

一旦发生错误，写入方法就变成了无操作，但错误值被保存。

Given the `errWriter` type and its `write` method, the code above can be refactored:

鉴于errWriter类型和它的写法，上面的代码可以被重构：

```go
ew := &errWriter{w: fd}
ew.write(p0[a:b])
ew.write(p1[c:d])
ew.write(p2[e:f])
// and so on
if ew.err != nil {
    return ew.err
}
```

This is cleaner, even compared to the use of a closure, and also makes the actual sequence of writes being done easier to see on the page. There is no clutter anymore. Programming with error values (and interfaces) has made the code nicer.

这样做比较干净，即使与使用闭包相比也是如此，而且也使实际的写操作顺序在页面上更容易看到。不再有任何杂乱无章的东西。用错误值（和接口）编程，使代码变得更漂亮。

It’s likely that some other piece of code in the same package can build on this idea, or even use `errWriter` directly.

很可能同一包中的其他一些代码可以建立在这个想法上，甚至直接使用errWriter。

Also, once `errWriter` exists, there’s more it could do to help, especially in less artificial examples. It could accumulate the byte count. It could coalesce writes into a single buffer that can then be transmitted atomically. And much more.

另外，一旦errWriter存在，它还可以做更多的事情来帮助，特别是在不太人性化的例子中。它可以积累字节数。它可以把写的内容凝聚成一个缓冲区，然后以原子方式传输。还有更多。

In fact, this pattern appears often in the standard library. The [`archive/zip`](https://go.dev/pkg/archive/zip/) and [`net/http`](https://go.dev/pkg/net/http/) packages use it. More salient to this discussion, the [`bufio` package’s `Writer`](https://go.dev/pkg/bufio/) is actually an implementation of the `errWriter` idea. Although `bufio.Writer.Write` returns an error, that is mostly about honoring the [`io.Writer`](https://go.dev/pkg/io/#Writer) interface. The `Write` method of `bufio.Writer` behaves just like our `errWriter.write` method above, with `Flush` reporting the error, so our example could be written like this:

事实上，这种模式经常出现在标准库中。archive/zip和net/http包都使用了它。对这次讨论来说，bufio包的Writer实际上是errWriter思想的一个实现。尽管bufio.Writer.Write返回一个错误，但这主要是为了尊重io.Writer接口。bufio.Writer的Write方法的行为就像我们上面的errWriter.write方法，Flush报告错误，所以我们的例子可以这样写：

```go
b := bufio.NewWriter(fd)
b.Write(p0[a:b])
b.Write(p1[c:d])
b.Write(p2[e:f])
// and so on
if b.Flush() != nil {
    return b.Flush()
}
```

There is one significant drawback to this approach, at least for some applications: there is no way to know how much of the processing completed before the error occurred. If that information is important, a more fine-grained approach is necessary. Often, though, an all-or-nothing check at the end is sufficient.

这种方法有一个明显的缺点，至少对某些应用来说是这样的：没有办法知道在错误发生之前完成了多少处理。如果这个信息很重要，就需要一个更精细的方法。不过，通常情况下，在最后进行全有或全无的检查就足够了。

We’ve looked at just one technique for avoiding repetitive error handling code. Keep in mind that the use of `errWriter` or `bufio.Writer` isn’t the only way to simplify error handling, and this approach is not suitable for all situations. The key lesson, however, is that errors are values and the full power of the Go programming language is available for processing them.

我们只看了一种避免重复错误处理代码的技术。请记住，使用errWriter或bufio.Writer并不是简化错误处理的唯一方法，而且这种方法并不适合所有情况。然而，关键的一课是，错误是价值，Go编程语言的全部功能都可以用来处理它们。

Use the language to simplify your error handling.

使用该语言来简化您的错误处理。

But remember: Whatever you do, always check your errors!

但请记住。无论您做什么，都要检查您的错误!

Finally, for the full story of my interaction with @jxck_, including a little video he recorded, visit [his blog](http://jxck.hatenablog.com/entry/golang-error-handling-lesson-by-rob-pike).

最后，关于我与@jxck_互动的完整故事，包括他录制的一个小视频，请访问他的博客。
