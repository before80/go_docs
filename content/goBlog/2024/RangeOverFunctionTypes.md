+++
title = "遍历函数类型"
date = 2024-09-06T12:19:40+08:00
weight = 910
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Range Over Function Types - 遍历函数类型

Ian Lance Taylor
20 August 2024

作者：Ian Lance Taylor
日期：2024年8月20日

## 简介 Introduction

This is the blog post version of my talk at GopherCon 2024.

​	这篇博客文章是我在 GopherCon 2024 上演讲的版本。

Range over function types is a new language feature in the Go 1.23 release. This blog post will explain why we are adding this new feature, what exactly it is, and how to use it.

​	遍历函数类型是 Go 1.23 版本中的一个新语言特性。本文将解释我们为什么要添加这个新特性、它到底是什么以及如何使用它。

## 为什么要引入这个特性？ Why?

Since Go 1.18 we’ve had the ability to write new generic container types in Go. For example, let’s consider this very simple `Set` type, a generic type implemented on top of a map.

​	自 Go 1.18 以来，我们已经可以在 Go 中编写新的泛型容器类型。举个例子，让我们来看一个非常简单的 `Set` 类型，这是一个基于 map 实现的泛型类型。

```go
// Set holds a set of elements.
// Set 保存一组元素。
type Set[E comparable] struct {
    m map[E]struct{}
}

// New returns a new [Set].
// New 返回一个新的 [Set]。
func New[E comparable]() *Set[E] {
    return &Set[E]{m: make(map[E]struct{})}
}
```

Naturally a set type has a way to add elements and a way to check whether elements are present. The details here don’t matter.

​	自然地，一个集合类型应该有添加元素和检查元素是否存在的方法。这里的细节并不重要。

```go
// Add adds an element to a set.
// Add 向集合中添加一个元素。
func (s *Set[E]) Add(v E) {
    s.m[v] = struct{}{}
}

// Contains reports whether an element is in a set.
// Contains 检查元素是否在集合中。
func (s *Set[E]) Contains(v E) bool {
    _, ok := s.m[v]
    return ok
}
```

And among other things we will want a function to return the union of two sets.

​	在其他功能中，我们可能需要一个函数来返回两个集合的并集。

```go
// Union returns the union of two sets.
// Union 返回两个集合的并集。
func Union[E comparable](s1, s2 *Set[E]) *Set[E] {
    r := New[E]()
    // Note for/range over internal Set field m.
    // We are looping over the maps in s1 and s2.
    // 注意在内部 Set 字段 m 上使用 for/range。
    // 我们在遍历 s1 和 s2 中的 map。
    for v := range s1.m {
        r.Add(v)
    }
    for v := range s2.m {
        r.Add(v)
    }
    return r
}
```

Let’s look at this implementation of the `Union` function for a minute. In order to compute the union of two sets, we need a way to get all the elements that are in each set. In this code we use a for/range statement over an unexported field of the set type. That only works if the `Union` function is defined in the set package.

​	让我们仔细看看 `Union` 函数的实现。为了计算两个集合的并集，我们需要获取每个集合中的所有元素。在这段代码中，我们使用了 for/range 语句来遍历集合类型的一个未导出字段。这只有在 `Union` 函数定义在集合包中时才有效。

But there are a lot of reasons why someone might want to loop over all the elements in a set. This set package has to provide some way for its users to do that.

​	但是，有很多理由表明可能有人想遍历集合中的所有元素。这个集合包必须为用户提供某种方式来实现这一点。

How should that work?

​	那么，这应该如何工作呢？

### 推送集合元素 Push Set elements

One approach is to provide a `Set` method that takes a function, and to call that function with every element in the Set. We’ll call this `Push`, because the `Set` pushes every value to the function. Here if the function returns false, we stop calling it.

​	一种方法是提供一个 `Set` 方法，该方法接受一个函数，并将集合中的每个元素传递给该函数。我们称之为 `Push`，因为集合将每个值推送给函数。这里如果函数返回 `false`，我们就停止调用它。

```go
func (s *Set[E]) Push(f func(E) bool) {
    for v := range s.m {
        if !f(v) {
            return
        }
    }
}
```

In the Go standard library, we see this general pattern used for cases like the [`sync.Map.Range`](https://pkg.go.dev/sync#Map.Range) method, the [`flag.Visit`](https://pkg.go.dev/flag#Visit) function, and the [`filepath.Walk`](https://pkg.go.dev/path/filepath#Walk) function. This is a general pattern, not an exact one; as it happens, none of those three examples work quite the same way.

​	在 Go 标准库中，我们可以看到这种通用模式用于类似 [`sync.Map.Range`](https://pkg.go.dev/sync#Map.Range) 方法、[`flag.Visit`](https://pkg.go.dev/flag#Visit) 函数和 [`filepath.Walk`](https://pkg.go.dev/path/filepath#Walk) 函数的场景。虽然这是一种通用模式，但并不是完全一致的；实际上，这三个例子都没有完全相同的工作方式。

This is what it looks like to use the `Push` method to print all the elements of a set: you call `Push` with a function that does what you want with the element.

​	这是使用 `Push` 方法打印集合中所有元素的示例：你调用 `Push` 并传递一个用于处理元素的函数。

```go
func PrintAllElementsPush[E comparable](s *Set[E]) {
    s.Push(func(v E) bool {
        fmt.Println(v)
        return true
    })
}
```

### 拉取集合元素 Pull Set elements

Another approach to looping over the elements of a `Set` is to return a function. Each time the function is called, it will return a value from the `Set`, along with a boolean that reports whether the value is valid. The boolean result will be false when the loop has gone through all the elements. In this case we also need a stop function that can be called when no more values are needed.

​	遍历 `Set` 元素的另一种方法是返回一个函数。每次调用该函数时，它都会返回集合中的一个值，并带有一个布尔值表示该值是否有效。当遍历了所有元素后，布尔值会变为 `false`。在这种情况下，我们还需要一个停止函数，在不再需要值时调用它。

This implementations uses a pair of channels, one for values in the set and one to stop returning values. We use a goroutine to send values on the channel. The `next` function returns an element from the set by reading from the element channel, and the `stop` function tells the goroutine to exit by closing the stop channel. We need the `stop` function to make sure that the goroutine exits when no more values are needed.

​	这个实现使用了一对通道，一个用于传递集合中的值，另一个用于停止返回值。我们使用一个 goroutine 在通道上发送值。`next` 函数通过读取元素通道返回集合中的一个元素，`stop` 函数通过关闭停止通道来告诉 goroutine 退出。我们需要 `stop` 函数来确保在不再需要值时，goroutine 能够退出。

```go
// Pull returns a next function that returns each
// element of s with a bool for whether the value
// is valid. The stop function should be called
// when finished calling the next function.
// Pull 返回一个 next 函数，该函数会依次返回集合中的每个元素，并带有一个布尔值表示该值是否有效。完成对 next 函数的调用后，应调用 stop 函数。
func (s *Set[E]) Pull() (func() (E, bool), func()) {
    ch := make(chan E)
    stopCh := make(chan bool)

    go func() {
        defer close(ch)
        for v := range s.m {
            select {
            case ch <- v:
            case <-stopCh:
                return
            }
        }
    }()

    next := func() (E, bool) {
        v, ok := <-ch
        return v, ok
    }

    stop := func() {
        close(stopCh)
    }

    return next, stop
}
```

Nothing in the standard library works exactly this way. Both [`runtime.CallersFrames`](https://pkg.go.dev/runtime#CallersFrames) and [`reflect.Value.MapRange`](https://pkg.go.dev/reflect#Value.MapRange) are similar, though they return values with methods rather than returning functions directly.

​	在标准库中没有任何东西是完全按照这种方式工作的。[`runtime.CallersFrames`](https://pkg.go.dev/runtime#CallersFrames) 和 [`reflect.Value.MapRange`](https://pkg.go.dev/reflect#Value.MapRange) 都类似，尽管它们返回值的方法是通过调用方法，而不是直接返回函数。

This is what it looks like to use the `Pull` method to print all the elements of a `Set`. You call `Pull` to get a function, and you repeatedly call that function in a for loop.

​	这是使用 `Pull` 方法打印集合中所有元素的示例：你调用 `Pull` 获取一个函数，然后在 for 循环中重复调用该函数。

```go
func PrintAllElementsPull[E comparable](s *Set[E]) {
    next, stop := s.Pull()
    defer stop()
    for v, ok := next(); ok; v, ok = next() {
        fmt.Println(v)
    }
}
```

## 规范化方法 Standardize the approach

We’ve now seen two different approaches to looping over all the elements of a set. Different Go packages use these approaches and several others. That means that when you start using a new Go container package you may have to learn a new looping mechanism. It also means that we can’t write one function that works with several different types of containers, as the container types will handle looping differently.

​	我们已经看到了两种遍历集合中所有元素的方法。不同的 Go 包使用这些方法以及其他几种方法。这意味着当你开始使用一个新的 Go 容器包时，可能需要学习一种新的遍历机制。这也意味着我们无法编写一个适用于多种容器类型的函数，因为不同的容器类型处理遍历的方式不同。

We want to improve the Go ecosystem by developing standard approaches for looping over containers.

​	我们希望通过开发用于遍历容器的标准方法来改善 Go 生态系统。

### 迭代器 Iterators

This is, of course, an issue that arises in many programming languages.

​	当然，这是许多编程语言中都会出现的问题。

The popular [Design Patterns book](https://en.wikipedia.org/wiki/Design_Patterns), first published in 1994, describes this as the iterator pattern. You use an iterator to “provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.” What this quote calls an aggregate object is what I’ve been calling a container. An aggregate object, or container, is just a value that holds other values, like the `Set` type we’ve been discussing.

​	1994 年首次出版的流行书籍 [《设计模式》](https://en.wikipedia.org/wiki/Design_Patterns) 将此称为迭代器模式。你使用迭代器来“提供一种方式来顺序访问聚合对象的元素，而不暴露其底层表示。”这个引用中的聚合对象就是我一直称之为容器的东西。聚合对象或容器只是保存其他值的值，就像我们讨论的 `Set` 类型一样。

Like many ideas in programming, iterators date back to Barbara Liskov’s [CLU language](https://en.wikipedia.org/wiki/CLU_(programming_language)), developed in the 1970’s.

​	像许多编程中的想法一样，迭代器可以追溯到 Barbara Liskov 在 1970 年代开发的 [CLU 语言](https://en.wikipedia.org/wiki/CLU_(programming_language))。

Today many popular languages provide iterators one way or another, including, among others, C++, Java, Javascript, Python, and Rust.

​	今天，许多流行的语言都以某种方式提供了迭代器，包括 C++、Java、Javascript、Python 和 Rust 等。

However, Go before version 1.23 did not.

​	然而，在 1.23 版本之前，Go 并没有提供迭代器。

### For/range

As we all know, Go has container types that are built in to the language: slices, arrays, and maps. And it has a way to access the elements of those values without exposing the underlying representation: the for/range statement. The for/range statement works for Go’s built-in container types (and also for strings, channels, and, as of Go 1.22, int).

​	众所周知，Go 语言内置了容器类型：切片、数组和映射（map）。它提供了一种访问这些容器元素的方式，而不暴露底层实现：即 for/range 语句。for/range 语句不仅适用于 Go 的内置容器类型（还适用于字符串、通道，以及从 Go 1.22 开始的 int）。

The for/range statement is iteration, but it is not iterators as they appear in today’s popular languages. Still, it would be nice to be able to use for/range to iterate over a user-defined container like the `Set` type.

​	for/range 语句是一种迭代方式，但它并不像当今流行语言中的迭代器。然而，能够使用 for/range 来遍历像 `Set` 类型这样的用户定义容器类型也是很有用的。

However, Go before version 1.23 did not support this.

​	然而，在 Go 1.23 之前，Go 并不支持这一功能。

### 本版本的改进 Improvements in this release

For Go 1.23 we’ve decided to support both for/range over user-defined container types, and a standardized form of iterators.

​	在 Go 1.23 中，我们决定支持 for/range 遍历用户定义的容器类型，并引入标准化的迭代器形式。

We extended the for/range statement to support ranging over function types. We’ll see below how this helps loop over user-defined containers.

​	我们扩展了 for/range 语句，使其能够支持遍历函数类型。稍后我们将看到这如何帮助遍历用户定义的容器。

We also added standard library types and functions to support using function types as iterators. A standard definition of iterators lets us write functions that work smoothly with different container types.

​	我们还在标准库中添加了类型和函数，以支持将函数类型用作迭代器。标准的迭代器定义让我们可以编写能够与不同容器类型无缝工作的函数。

### 遍历（某些）函数类型 Range over (some) function types

The improved for/range statement doesn’t support arbitrary function types. As of Go 1.23 it now supports ranging over functions that take a single argument. The single argument must itself be a function that takes zero to two arguments and returns a bool; by convention, we call it the yield function.

​	改进后的 for/range 语句不支持任意的函数类型。从 Go 1.23 开始，它支持遍历接收单个参数的函数。该参数本身必须是一个函数，且该函数接受零到两个参数并返回一个布尔值；按照惯例，我们称其为 yield 函数。

```go
func(yield func() bool)

func(yield func(V) bool)

func(yield func(K, V) bool)
```

When we speak of an iterator in Go, we mean a function with one of these three types. As we’ll discuss below, there is another kind of iterator in the standard library: a pull iterator. When it is necessary to distinguish between standard iterators and pull iterators, we call the standard iterators push iterators. That is because, as we will see, they push out a sequence of values by calling a yield function.

​	当我们谈到 Go 中的迭代器时，我们指的是具有上述三种类型之一的函数。正如我们将在下文讨论的那样，标准库中还有另一种迭代器：拉取迭代器（pull iterator）。当需要区分标准迭代器和拉取迭代器时，我们称标准迭代器为推送迭代器（push iterator），因为正如我们将看到的，它们通过调用 yield 函数推送一系列的值。

### 标准（推送）迭代器 Standard (push) iterators

To make iterators easier to use, the new standard library package iter defines two types: `Seq` and `Seq2`. These are names for the iterator function types, the types that can be used with the for/range statement. The name `Seq` is short for sequence, as iterators loop through a sequence of values.

​	为了使迭代器更易于使用，新的标准库包 iter 定义了两种类型：`Seq` 和 `Seq2`。这些名称表示可以与 for/range 语句一起使用的迭代器函数类型。`Seq` 这个名字是 sequence（序列）的缩写，因为迭代器循环遍历一系列的值。

```go
package iter

type Seq[V any] func(yield func(V) bool)

type Seq2[K, V any] func(yield func(K, V) bool)

// for now, no Seq0
// 目前没有 Seq0
```

The difference between `Seq` and `Seq2` is just that `Seq2` is a sequence of pairs, such as a key and a value from a map. In this post we’ll focus on `Seq` for simplicity, but most of what we say covers `Seq2` as well.

​	`Seq` 和 `Seq2` 之间的区别仅在于 `Seq2` 是一对对的序列，例如映射中的键值对。为了简化讨论，本文将重点放在 `Seq` 上，但我们所说的大多数内容同样适用于 `Seq2`。

It’s easiest to explain how iterators work with an example. Here the `Set` method `All` returns a function. The return type of `All` is `iter.Seq[E]`, so we know that it returns an iterator.

​	通过示例来解释迭代器的工作原理最为简单。这里，`Set` 的 `All` 方法返回一个函数。`All` 的返回类型是 `iter.Seq[E]`，因此我们知道它返回的是一个迭代器。

```go
// All is an iterator over the elements of s.
// All 是 s 中元素的迭代器。
func (s *Set[E]) All() iter.Seq[E] {
    return func(yield func(E) bool) {
        for v := range s.m {
            if !yield(v) {
                return
            }
        }
    }
}
```

The iterator function itself takes another function, the yield function, as an argument. The iterator calls the yield function with every value in the set. In this case the iterator, the function returned by `Set.All`, is a lot like the `Set.Push` function we saw earlier.

​	迭代器函数本身接收另一个函数，即 yield 函数，作为参数。迭代器会对集合中的每个值调用 yield 函数。在这种情况下，迭代器，即 `Set.All` 返回的函数，与我们之前看到的 `Set.Push` 函数非常相似。

This shows how iterators work: for some sequence of values, they call a yield function with each value in the sequence. If the yield function returns false, no more values are needed, and the iterator can just return, doing any cleanup that may be required. If the yield function never returns false, the iterator can just return after calling yield with all the values in the sequence.

​	这展示了迭代器的工作原理：对于某个值序列，它们通过 yield 函数传递序列中的每个值。如果 yield 函数返回 false，则不再需要其他值，迭代器可以直接返回，并执行任何必要的清理操作。如果 yield 函数从不返回 false，迭代器将在对序列中所有值调用 yield 之后返回。

That’s how they work, but let’s acknowledge that the first time you see one of these, your first reaction is probably “there are a lot of functions flying around here.” You’re not wrong about that. Let’s focus on two things.

​	这就是它们的工作原理，但我们要承认，第一次看到这个例子时，你的第一反应可能是“这里有很多函数在飞来飞去。”你对此的感觉并没有错。我们来关注两个方面。

The first is that once you get past the first line of this function’s code, the actual implementation of the iterator is pretty simple: call yield with every element of the set, stopping if yield returns false.

​	首先，一旦你越过了这个函数代码的第一行，迭代器的实际实现就相当简单：对集合中的每个元素调用 yield，如果 yield 返回 false 则停止。

```go
        for v := range s.m {
            if !yield(v) {
                return
            }
        }
```

The second is that using this is really easy. You call `s.All` to get an iterator, and then you use for/range to loop over all the elements in `s`. The for/range statement supports any iterator, and this shows how easy that is to use.

​	其次，使用这个功能真的非常简单。你调用 `s.All` 来获取一个迭代器，然后使用 for/range 来遍历 `s` 中的所有元素。for/range 语句支持任何迭代器，而这也展示了使用它是多么容易。

```go
func PrintAllElements[E comparable](s *Set[E]) {
    for v := range s.All() {
        fmt.Println(v)
    }
}
```

In this kind of code `s.All` is a method that returns a function. We are calling `s.All`, and then using for/range to range over the function that it returns. In this case we could have made `Set.All` be an iterator function itself, rather than having it return an iterator function. However, in some cases that won’t work, such as if the function that returns the iterator needs to take an argument, or needs to do some set up work. As a matter of convention, we encourage all container types to provide an `All` method that returns an iterator, so that programmers don’t have to remember whether to range over `All` directly or whether to call `All` to get a value they can range over. They can always do the latter.

​	在这种代码中，`s.All` 是一个返回函数的方法。我们调用 `s.All`，然后使用 for/range 遍历它返回的函数。在这种情况下，我们本可以让 `Set.All` 本身成为一个迭代器函数，而不是返回一个迭代器函数。然而，在某些情况下这并不奏效，比如如果返回迭代器的函数需要接收一个参数，或需要做一些设置工作。作为一种惯例，我们鼓励所有容器类型提供一个返回迭代器的 `All` 方法，这样程序员就不必记住是直接遍历 `All`，还是调用 `All` 以获得可以遍历的值。他们总是可以选择后者。

If you think about it, you’ll see that the compiler must be adjusting the loop to create a yield function to pass to the iterator returned by `s.All`. There’s a fair bit of complexity in the Go compiler and runtime to make this efficient, and to correctly handle things like `break` or `panic` in the loop. We’re not going to cover any of that in this blog post. Fortunately the implementation details are not important when it comes to actually using this feature.

​	如果你仔细想想，你会发现编译器必须调整循环以创建一个 yield 函数传递给由 `s.All` 返回的迭代器。Go 编译器和运行时在让这一功能高效运行以及正确处理循环中的 `break` 或 `panic` 等方面做了很多工作。我们不会在本文中讨论这些实现细节。幸运的是，这些实现细节在实际使用这个功能时并不重要。

### 拉取迭代器 Pull iterators

We’ve now seen how to use iterators in a for/range loop. But a simple loop is not the only way to use an iterator. For example, sometimes we may need to iterate over two containers in parallel. How do we do that?

​	现在我们已经看到了如何在 for/range 循环中使用迭代器。但简单循环并不是使用迭代器的唯一方式。例如，有时我们可能需要并行遍历两个容器。我们该怎么做呢？

The answer is that we use a different kind of iterator: a pull iterator. We’ve seen that a standard iterator, also known as a push iterator, is a function that takes a yield function as an argument and pushes each value in a sequence by calling the yield function.

​	答案是使用另一种迭代器：拉取迭代器。我们已经看到，标准迭代器（也称为推送迭代器）是一个接受 yield 函数作为参数并通过调用 yield 函数推送每个值的函数。

A pull iterator works the other way around: it is a function that is written such that each time you call it, it returns the next value in the sequence.

​	拉取迭代器则相反：它是一个函数，每次调用它时，它都会返回序列中的下一个值。

We’ll repeat the difference between the two types of iterators to help you remember:

​	我们将重复两种迭代器之间的区别，以帮助您记住：

- A push iterator pushes each value in a sequence to a yield function. Push iterators are standard iterators in the Go standard library, and are supported directly by the for/range statement.
- **推送迭代器（Push Iterator）** 将序列中的每个值推送到 `yield` 函数。推送迭代器是 Go 标准库中的标准迭代器，并且直接受到 `for/range` 语句的支持。

- A pull iterator works the other way around. Each time you call a pull iterator, it pulls another value from a sequence and returns it. Pull iterators are *not* supported directly by the for/range statement; however, it’s straightforward to write an ordinary for statement that loops through a pull iterator. In fact, we saw an example earlier when we looked at using the `Set.Pull` method.
- **拉取迭代器（Pull Iterator）** 则相反。每次调用拉取迭代器时，它会从序列中拉取另一个值并返回。拉取迭代器*不*被 `for/range` 语句直接支持；然而，编写一个普通的 `for` 语句来循环遍历拉取迭代器是很简单的。实际上，当我们之前查看使用 `Set.Pull` 方法时，看到了一个示例。

You could write a pull iterator yourself, but normally you don’t have to. The new standard library function [`iter.Pull`](https://pkg.go.dev/iter#Pull) takes a standard iterator, that is to say a function that is a push iterator, and returns a pair of functions. The first is a pull iterator: a function that returns the next value in the sequence each time it is called. The second is a stop function that should be called when we are done with the pull iterator. This is like the `Set.Pull` method we saw earlier.

​	您可以自己编写一个拉取迭代器，但通常您不需要。新的标准库函数 [`iter.Pull`](https://pkg.go.dev/iter#Pull) 接受一个标准迭代器，即推送迭代器函数，并返回一对函数。第一个是拉取迭代器：一个函数，每次调用它时返回序列中的下一个值。第二个是一个停止函数，应在我们完成使用拉取迭代器时调用。它类似于我们之前看到的 `Set.Pull` 方法。

The first function returned by `iter.Pull`, the pull iterator, returns a value and a boolean that reports whether that value is valid. The boolean will be false at the end of the sequence.

​	`iter.Pull` 返回的第一个函数，拉取迭代器，返回一个值和一个布尔值，布尔值指示该值是否有效。在序列结束时，布尔值将为 `false`。

`iter.Pull` returns a stop function in case we don’t read through the sequence to the end. In the general case the push iterator, the argument to `iter.Pull`, may start goroutines, or build new data structures that need to be cleaned up when iteration is complete. The push iterator will do any cleanup when the yield function returns false, meaning that no more values are required. When used with a for/range statement, the for/range statement will ensure that if the loop exits early, through a `break` statement or for any other reason, then the yield function will return false. With a pull iterator, on the other hand, there is no way to force the yield function to return false, so the stop function is needed.

​	`iter.Pull` 返回一个停止函数，以防我们没有读取完序列。在一般情况下，传递给 `iter.Pull` 的推送迭代器可能启动 goroutine，或者构建需要在迭代完成后清理的新数据结构。当 `yield` 函数返回 `false` 时，推送迭代器将执行任何清理操作，这意味着不再需要更多的值。当与 `for/range` 语句一起使用时，`for/range` 语句将确保如果循环提前退出，例如通过 `break` 语句或其他任何原因，`yield` 函数将返回 `false`。然而，对于拉取迭代器，没有办法强制 `yield` 函数返回 `false`，因此需要停止函数。

Another way to say this is that calling the stop function will cause the yield function to return false when it is called by the push iterator.

​	换句话说，调用停止函数将导致 `yield` 函数在被推送迭代器调用时返回 `false`。

Strictly speaking you don’t need to call the stop function if the pull iterator returns false to indicate that it has reached the end of the sequence, but it’s usually simpler to just always call it.

​	严格来说，如果拉取迭代器返回 `false` 以表明已到达序列末尾，则您不需要调用停止函数，但通常始终调用它会更简单。

Here is an example of using pull iterators to walk through two sequences in parallel. This function reports whether two arbitrary sequences contain the same elements in the same order.

​	下面是一个使用拉取迭代器并行遍历两个序列的示例。此函数报告两个任意序列是否包含相同的元素且顺序相同。

```go
// EqSeq reports whether two iterators contain the same
// elements in the same order.
// EqSeq 报告两个迭代器是否包含相同的元素且顺序相同。
func EqSeq[E comparable](s1, s2 iter.Seq[E]) bool {
    next1, stop1 := iter.Pull(s1)
    defer stop1()
    next2, stop2 := iter.Pull(s2)
    defer stop2()
    for {
        v1, ok1 := next1()
        v2, ok2 := next2()
        if !ok1 {
            return !ok2
        }
        if ok1 != ok2 || v1 != v2 {
            return false
        }
    }
}
```

The function uses `iter.Pull` to convert the two push iterators, `s1` and `s2`, into pull iterators. It uses `defer` statements to make sure that the pull iterators are stopped when we are done with them.

​	该函数使用 `iter.Pull` 将两个推送迭代器 `s1` 和 `s2` 转换为拉取迭代器。它使用 `defer` 语句确保在我们使用完它们时停止拉取迭代器。

Then the code loops, calling the pull iterators to retrieve values. If the first sequence is done, it returns true if the second sequence is also done, or false if it isn’t. If the values are different, it returns false. Then it loops to pull the next two values.

​	然后代码循环调用拉取迭代器以检索值。如果第一个序列完成，它将返回 `true`，如果第二个序列也完成，或者如果未完成则返回 `false`。如果值不同，它将返回 `false`。然后循环以拉取下两个值。

As with push iterators, there is some complexity in the Go runtime to make pull iterators efficient, but this does not affect code that actually uses the `iter.Pull` function.

​	与推送迭代器一样，Go 运行时有一些复杂性来使拉取迭代器变得高效，但这不会影响实际使用 `iter.Pull` 函数的代码。

## 迭代迭代器Iterating on iterators

Now you know everything there is to know about range over function types and about iterators. We hope you enjoy using them!

​	现在您已经了解了关于函数类型上的 range 以及迭代器的所有内容。我们希望您享受使用它们！

Still, there are a few more things worth mentioning.

​	不过，仍有一些值得一提的内容。

### 适配器 Adapters

An advantage of a standard definition of iterators is the ability to write standard adapter functions that use them.

​	标准迭代器定义的一个优点是能够编写使用它们的标准适配器函数。

For example, here is a function that filters a sequence of values, returning a new sequence. This `Filter` function takes an iterator as an argument and returns a new iterator. The other argument is a filter function that decides which values should be in the new iterator that `Filter` returns.

​	例如，下面是一个过滤值序列的函数，它返回一个新序列。这个 `Filter` 函数将一个迭代器作为参数并返回一个新迭代器。另一个参数是决定哪些值应该包含在 `Filter` 返回的新迭代器中的过滤器函数。

```go
// Filter returns a sequence that contains the elements
// of s for which f returns true.
// Filter 返回一个包含 f 返回 true 的 s 元素的序列。
func Filter[V any](f func(V) bool, s iter.Seq[V]) iter.Seq[V] {
    return func(yield func(V) bool) {
        for v := range s {
            if f(v) {
                if !yield(v) {
                    return
                }
            }
        }
    }
}
```

As with the earlier example, the function signatures look complicated when you first see them. Once you get past the signatures, the implementation is straightforward.

​	如同前面的示例一样，当您第一次看到函数签名时，可能看起来比较复杂。一旦您理解了签名，函数的实现就相对简单了。

```go
        for v := range s {
            if f(v) {
                if !yield(v) {
                    return
                }
            }
        }
```

The code ranges over the input iterator, checks the filter function, and calls yield with the values that should go into the output iterator.

​	代码遍历输入迭代器，检查过滤器函数，并将应该进入输出迭代器的值传递给 `yield`。

We’ll show an example of using `Filter` below.

​	我们将在下面展示一个使用 `Filter` 的示例。

(There is no version of `Filter` in the Go standard library today, but one may be added in future releases.)

​	（今天的 Go 标准库中没有 `Filter` 的版本，但在未来的版本中可能会添加。）

### 二叉树 Binary tree

As an example of how convenient a push iterator can be to loop over a container type, let’s consider this simple binary tree type.

​	作为推送迭代器用于循环遍历容器类型的便利性的一个示例，我们来看一下这个简单的二叉树类型。

```go
// Tree is a binary tree.
type Tree[E any] struct {
    val         E
    left, right *Tree[E]
}
```

We won’t show the code to insert values into the tree, but naturally there should be some way to range over all the values in the tree.

​	我们不会展示向树中插入值的代码，但显然应该有某种方式遍历树中的所有值。

It turns out that the iterator code is easier to write if it returns a bool. Since the function types supported by for/range don’t return anything, the `All` method here return a small function literal that calls the iterator itself, here called `push`, and ignores the bool result.

​	事实证明，编写返回布尔值的迭代器代码会更容易。由于 `for/range` 语句支持的函数类型不返回任何内容，这里的 `All` 方法返回一个小的函数文字，该文字调用迭代器本身（这里称为 `push`），并忽略布尔结果。

```go
// All returns an iterator over the values in t.
// All 返回一个遍历 t 中值的迭代器。
func (t *Tree[E]) All() iter.Seq[E] {
    return func(yield func(E) bool) {
        t.push(yield)
    }
}

// push pushes all elements to the yield function.
// push 将所有元素推送到 yield 函数。
func (t *Tree[E]) push(yield func(E) bool) bool {
    if t == nil {
        return true
    }
    return t.left.push(yield) &&
        yield(t.val) &&
        t.right.push(yield)
}
```

The `push` method uses recursion to walk over the whole tree, calling yield on each element. If the yield function returns false, the method returns false all the way up the stack. Otherwise it just returns once the iteration is complete.

​	`push` 方法使用递归遍历整棵树，对每个元素调用 `yield`。如果 `yield` 函数返回 `false`，则方法在整个调用栈中返回 `false`。否则，在迭代完成后它会直接返回。

This shows how straightforward it is to use this iterator approach to loop over even complex data structures. There is no need to maintain a separate stack to record the position within the tree; we can just use the goroutine call stack to do that for us.

​	这展示了使用这种迭代器方法遍历复杂数据结构是多么直接。我们不需要维护一个单独的堆栈来记录在树中的位置；我们可以直接使用 goroutine 调用栈来为我们完成这个任务。

### 新的迭代器函数 New iterator functions.

Also new in Go 1.23 are functions in the slices and maps packages that work with iterators.

​	Go 1.23 中还引入了 `slices` 和 `maps` 包中与迭代器配合使用的新函数。

Here are the new functions in the slices package. `All` and `Values` are functions that return iterators over the elements of a slice. `Collect` fetches the values out of an iterator and returns a slice holding those values. See the docs for the others.

​	下面是 `slices` 包中的新函数。`All` 和 `Values` 是返回切片元素迭代器的函数。`Collect` 从迭代器中提取值并返回一个包含这些值的切片。有关其他函数的详细信息，请参阅文档。

- [`All([\]E) iter.Seq2[int, E]`](https://pkg.go.dev/slices#All)

- [`Values([\]E) iter.Seq[E]`](https://pkg.go.dev/slices#Values)
- [`Collect(iter.Seq[E\]) []E`](https://pkg.go.dev/slices#Collect)
- [`AppendSeq([\]E, iter.Seq[E]) []E`](https://pkg.go.dev/slices#AppendSeq)
- [`Backward([\]E) iter.Seq2[int, E]`](https://pkg.go.dev/slices#Backward)
- [`Sorted(iter.Seq[E\]) []E`](https://pkg.go.dev/slices#Sorted)
- [`SortedFunc(iter.Seq[E\], func(E, E) int) []E`](https://pkg.go.dev/slices#SortedFunc)
- [`SortedStableFunc(iter.Seq[E\], func(E, E) int) []E`](https://pkg.go.dev/slices#SortedStableFunc)
- [`Repeat([\]E, int) []E`](https://pkg.go.dev/slices#Repeat)
- [`Chunk([\]E, int) iter.Seq([]E)`](https://pkg.go.dev/slices#Chunk)

Here are the new functions in the maps package. `All`, `Keys`, and `Values` returns iterators over the map contents. `Collect` fetches the keys and values out of an iterator and returns a new map.

​	以下是 maps 包中的新函数。`All`、`Keys` 和 `Values` 返回 map 内容的迭代器。`Collect` 从迭代器中提取键和值，并返回一个新 map。

- [`All(map[K\]V) iter.Seq2[K, V]`](https://pkg.go.dev/maps#All)

- [`Keys(map[K\]V) iter.Seq[K]`](https://pkg.go.dev/maps#Keys)
- [`Values(map[K\]V) iter.Seq[V]`](https://pkg.go.dev/maps#Values)
- [`Collect(iter.Seq2[K, V\]) map[K, V]`](https://pkg.go.dev/maps#Collect)
- [`Insert(map[K, V\], iter.Seq2[K, V])`](https://pkg.go.dev/maps#Insert)

### 标准库迭代器示例 Standard library iterator example

Here is an example of how you might use these new functions along with the `Filter` function we saw earlier. This function takes a map from int to string and returns a slice holding just the values in the map that are longer than some argument `n`.

​	这是一个如何使用这些新函数以及之前提到的 `Filter` 函数的示例。此函数接收一个从 int 到 string 的 map 并返回一个切片，其中仅包含 map 中长度大于某个参数 `n` 的值。

```go
// LongStrings returns a slice of just the values
// in m whose length is n or more.
// LongStrings 返回一个仅包含 m 中长度为 n 或更多的值的切片。
func LongStrings(m map[int]string, n int) []string {
    isLong := func(s string) bool {
        return len(s) >= n
    }
    return slices.Collect(Filter(isLong, maps.Values(m)))
}
```

The `maps.Values` function returns an iterator over the values in `m`. `Filter` reads that iterator and returns a new iterator that only contains the long strings. `slices.Collect` reads from that iterator into a new slice.

​	`maps.Values` 函数返回 `m` 中值的迭代器。`Filter` 读取该迭代器，并返回一个仅包含长字符串的新迭代器。`slices.Collect` 将从该迭代器中读取数据并放入一个新的切片中。

Of course, you could write a loop to do this easily enough, and in many cases a loop will be clearer. We don’t want to encourage everybody to write code in this style all the time. That said, the advantage of using iterators is that this kind of function works the same way with any sequence. In this example, notice how Filter is using a map as an input and a slice as an output, without having to change the code in Filter at all.

​	当然，您可以轻松编写一个循环来实现这一点，并且在许多情况下，循环可能更清晰。我们并不鼓励大家总是以这种风格编写代码。然而，使用迭代器的好处在于，这类函数可以在任何序列上以相同方式工作。在这个例子中，注意 `Filter` 使用了一个 map 作为输入并返回一个切片作为输出，而不需要对 `Filter` 中的代码进行任何更改。

### 遍历文件中的行 Looping over lines in a file

Although most of the examples we’ve seen have involved containers, iterators are flexible.

​	虽然我们看到的大多数示例都涉及容器，但迭代器是灵活的。

Consider this simple code, which doesn’t use iterators, to loop over the lines in a byte slice. This is easy to write and fairly efficient.

​	考虑这段不使用迭代器的简单代码，用于遍历字节切片中的每一行。这段代码编写起来很容易，并且效率相当高。

```go
    for _, line := range bytes.Split(data, []byte{'\n'}) {
        handleLine(line)
    }
```

However, `bytes.Split` does allocate and return a slice of byte slices to hold the lines. The garbage collector will have to do a bit of work to eventually free that slice.

​	然而，`bytes.Split` 确实会分配并返回一个字节切片的切片来保存这些行。垃圾收集器最终将不得不进行一些工作来释放这个切片。

Here is a function that returns an iterator over the lines of some byte slice. After the usual iterator signatures, the function is pretty simple. We keep picking lines out of data until there is nothing left, and we pass each line to the yield function.

​	这是一个返回字节切片中每一行的迭代器的函数。遵循常见的迭代器签名后，这个函数非常简单。我们不断从数据中获取行，直到没有剩余的内容，然后将每一行传递给 yield 函数。

```go
// Lines returns an iterator over lines in data.
// Lines 返回 data 中每一行的迭代器。
func Lines(data []byte) iter.Seq[[]byte] {
    return func(yield func([]byte) bool) {
        for len(data) > 0 {
            line, rest, _ := bytes.Cut(data, []byte{'\n'})
            if !yield(line) {
                return
            }
            data = rest
        }
    }
}
```

Now our code to loop over the lines of a byte slice looks like this.

​	现在，我们遍历字节切片中每一行的代码如下所示。

```go
    for _, line := range Lines(data) {
        handleLine(line)
    }
```

This is just as easy to write as the earlier code, and is a bit more efficient because it doesn’t have allocate a slice of lines.

​	这段代码编写起来与之前的代码一样容易，并且效率更高一些，因为它不需要分配一个行的切片。

### 将函数传递给推送迭代器 Passing a function to a push iterator

For our final example, we’ll see that you don’t have to use a push iterator in a range statement.

​	最后一个示例，我们将看到不一定要在 range 语句中使用推送迭代器。

Earlier we saw a `PrintAllElements` function that prints out each element of a set. Here is another way to print all the elements of a set: call `s.All` to get an iterator, then pass in a hand-written yield function. This yield function just prints a value and returns true. Note that there are two function calls here: we call `s.All` to get an iterator which is itself a function, and we call that function with our hand-written yield function.

​	之前我们看到一个 `PrintAllElements` 函数，它打印集合中的每个元素。这里是另一种打印集合中所有元素的方法：调用 `s.All` 以获取一个迭代器，然后传入一个手写的 yield 函数。这个 yield 函数只打印一个值并返回 true。注意，这里有两个函数调用：我们调用 `s.All` 以获取一个迭代器，它本身是一个函数，然后我们调用该函数并传入我们手写的 yield 函数。

```go
func PrintAllElements[E comparable](s *Set[E]) {
    s.All()(func(v E) bool {
        fmt.Println(v)
        return true
    })
}
```

There’s no particular reason to write this code this way. This is just an example to show that the yield function isn’t magic. It can be any function you like.

​	没有特别的理由以这种方式编写代码。这只是一个示例，说明 yield 函数并不是魔法。它可以是您喜欢的任何函数。

## 更新 go.mod - Update go.mod

A final note: every Go module specifies the language version that it uses. That means that in order to use new language features in an existing module you may need to update that version. This is true for all new language features; it’s not something specific to range over function types. As range over function types is new in the Go 1.23 release, using it requires specifying at least Go language version 1.23.

​	最后一点注意事项：每个 Go 模块都会指定其使用的语言版本。这意味着，要在现有模块中使用新的语言特性，您可能需要更新该版本。这对于所有新的语言特性都是如此；这并不是特定于函数类型 range 的要求。由于函数类型 range 是 Go 1.23 版本中的新功能，使用它需要指定至少 Go 语言版本 1.23。

There are (at least) four ways to set the language version:

​	有（至少）四种方法来设置语言版本：

- On the command line, run `go get go@1.23` (or `go mod edit -go=1.23` to only edit the `go` directive).
- 在命令行上运行 `go get go@1.23`（或者 `go mod edit -go=1.23` 来只编辑 `go` 指令）。
- Manually edit the `go.mod` file and change the `go` line.
- 手动编辑 `go.mod` 文件并更改 `go` 行。
- Keep the older language version for the module as a whole, but use a `//go:build go1.23` build tag to permit using range over function types in a specific file.
- 保持模块的较旧语言版本，但在特定文件中使用 `//go:build go1.23` 构建标记以允许使用函数类型 range。
