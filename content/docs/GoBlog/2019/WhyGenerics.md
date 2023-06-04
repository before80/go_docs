+++
title = "为什么泛型？"
weight = 12
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Why Generics? - 为什么泛型？

https://go.dev/blog/why-generics

Ian Lance Taylor
31 July 2019

## Introduction 简介

This is the blog post version of my talk last week at Gophercon 2019.

这是我上周在Gophercon 2019的演讲的博文版本。

<iframe src="https://www.youtube.com/embed/WzgLqE-3IhY?rel=0" width="560" height="315" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

This article is about what it would mean to add generics to Go, and why I think we should do it. I’ll also touch on an update to a possible design for adding generics to Go.

这篇文章是关于在Go中添加泛型的意义，以及我为什么认为我们应该这样做。我还会谈谈为Go添加泛型的可能设计的更新。

Go was released on November 10, 2009. Less than 24 hours later we saw the [first comment about generics](https://groups.google.com/d/msg/golang-nuts/70-pdwUUrbI/onMsQspcljcJ). (That comment also mentions exceptions, which we added to the language, in the form of `panic` and `recover`, in early 2010.)

Go是在2009年11月10日发布的。不到24小时后，我们看到了第一条关于泛型的评论。(那条评论还提到了例外，我们在2010年初以恐慌和恢复的形式将其添加到语言中）。

In three years of Go surveys, lack of generics has always been listed as one of the top three problems to fix in the language.

在三年的Go调查中，缺乏泛型总是被列为Go语言中需要解决的三大问题之一。

## Why generics? 为什么是泛型？

But what does it mean to add generics, and why would we want it?

但是增加泛型是什么意思，我们为什么要这样做？

To paraphrase [Jazayeri, et al](https://www.dagstuhl.de/en/program/calendar/semhp/?semnr=98171): generic programming enables the representation of functions and data structures in a generic form, with types factored out.

用Jazayeri等人的话说：泛型编程能够以泛型的形式表示函数和数据结构，而类型则被剔除。

What does that mean?

这意味着什么呢？

For a simple example, let’s assume we want to reverse the elements in a slice. It’s not something that many programs need to do, but it’s not all that unusual.

举个简单的例子，让我们假设我们想把一个片断中的元素倒过来。这不是很多程序需要做的事情，但也不是那么不寻常。

Let’s say it’s a slice of int.

比方说，这是一个int的片断。

```go linenums="1"
func ReverseInts(s []int) {
    first := 0
    last := len(s)
    for first < last {
        s[first], s[last] = s[last], s[first]
        first++
        last--
    }
}
```

Pretty simple, but even for a simple function like that you’d want to write a few test cases. In fact, when I did, I found a bug. I’m sure many readers have spotted it already.

很简单，但即使是这样一个简单的函数，您也会想写一些测试案例。事实上，当我这样做时，我发现了一个错误。我相信很多读者已经发现了它。

```go linenums="1"
func ReverseInts(s []int) {
    first := 0
    last := len(s) - 1
    for first < last {
        s[first], s[last] = s[last], s[first]
        first++
        last--
    }
}
```

We need to subtract 1 when we set the variable last.

我们在设置变量last的时候需要减去1。

Now let’s reverse a slice of string.

现在我们来反转字符串的一个片断。

```go linenums="1"
func ReverseStrings(s []string) {
    first := 0
    last := len(s) - 1
    for first < last {
        s[first], s[last] = s[last], s[first]
        first++
        last--
    }
}
```

If you compare `ReverseInts` and `ReverseStrings`, you’ll see that the two functions are exactly the same, except for the type of the parameter. I don’t think any reader is surprised by that.

如果您比较一下ReverseInts和ReverseStrings，您会发现这两个函数是完全一样的，除了参数的类型不同。我想没有任何读者会对此感到惊讶。

What some people new to Go find surprising is that there is no way to write a simple `Reverse` function that works for a slice of any type.

一些刚接触Go的人感到惊讶的是，没有办法编写一个简单的反向函数，对任何类型的片断都有效。

Most other languages do let you write that kind of function.

大多数其他语言都允许您写这种函数。

In a dynamically typed language like Python or JavaScript you can simply write the function, without bothering to specify the element type. This doesn’t work in Go because Go is statically typed, and requires you to write down the exact type of the slice and the type of the slice elements.

在像Python或JavaScript这样的动态类型语言中，您可以简单地编写函数，而不必费心去指定元素类型。这在Go中是行不通的，因为Go是静态类型的，它要求您写下分片的确切类型和分片元素的类型。

Most other statically typed languages, like C++ or Java or Rust or Swift, support generics to address exactly this kind of issue.

大多数其他静态类型的语言，如C++或Java或Rust或Swift，都支持泛型来解决这类问题。

## Go generic programming today 今天的Go泛型编程

So how do people write this kind of code in Go?

那么人们如何在Go中编写这种代码呢？

In Go you can write a single function that works for different slice types by using an interface type, and defining a method on the slice types you want to pass in. That is how the standard library’s `sort.Sort` function works.

在Go中，您可以通过使用一个接口类型，并在您想传入的片断类型上定义一个方法，来编写一个适用于不同片断类型的单一函数。这就是标准库的sort.Sort函数的工作方式。

In other words, interface types in Go are a form of generic programming. They let us capture the common aspects of different types and express them as methods. We can then write functions that use those interface types, and those functions will work for any type that implements those methods.

换句话说，Go中的接口类型是通用编程的一种形式。它们让我们捕捉不同类型的共同点，并将其表达为方法。然后我们可以编写使用这些接口类型的函数，这些函数将对任何实现这些方法的类型起作用。

But this approach falls short of what we want. With interfaces you have to write the methods yourself. It’s awkward to have to define a named type with a couple of methods just to reverse a slice. And the methods you write are exactly the same for each slice type, so in a sense we’ve just moved and condensed the duplicate code, we haven’t eliminated it. Although interfaces are a form of generics, they don’t give us everything we want from generics.

但这种方法并没有达到我们的要求。对于接口，您必须自己编写方法。要定义一个带有几个方法的命名类型来反转一个片断，这是很尴尬的。而且您所写的方法对每个片断类型都是完全一样的，所以从某种意义上说，我们只是移动和浓缩了重复的代码，并没有消除它。尽管接口是泛型的一种形式，但它并没有给我们提供我们想要的泛型的一切。

A different way of using interfaces for generics, which could get around the need to write the methods yourself, would be to have the language define methods for some kinds of types. That isn’t something the language supports today, but, for example, the language could define that every slice type has an Index method that returns an element. But in order to use that method in practice it would have to return an empty interface type, and then we lose all the benefits of static typing. More subtly, there would be no way to define a generic function that takes two different slices with the same element type, or that takes a map of one element type and returns a slice of the same element type. Go is a statically typed language because that makes it easier to write large programs; we don’t want to lose the benefits of static typing in order to gain the benefits of generics.

将接口用于泛型的另一种方式是让语言为某些类型定义方法，这样就可以避免自己编写方法的需要。这不是现在的语言所支持的，但是，比如说，语言可以定义每个片断类型都有一个返回元素的索引方法。但是为了在实践中使用这个方法，它必须返回一个空的接口类型，这样我们就失去了静态类型化的所有好处。更微妙的是，我们没有办法定义一个通用函数来接收两个具有相同元素类型的不同片断，或者接收一个元素类型的映射并返回一个相同元素类型的片断。Go是一种静态类型的语言，因为这使得编写大型程序更加容易；我们不希望为了获得泛型的好处而失去静态类型的好处。

Another approach would be to write a generic `Reverse` function using the reflect package, but that is so awkward to write and slow to run that few people do that. That approach also requires explicit type assertions and has no static type checking.

另一种方法是使用reflect包写一个泛型的Reverse函数，但是这样写起来很麻烦，运行起来也很慢，很少有人这样做。这种方法也需要明确的类型断言，而且没有静态类型检查。

Or, you could write a code generator that takes a type and generates a `Reverse` function for slices of that type. There are several code generators out there that do just that. But this adds another step to every package that needs `Reverse`, it complicates the build because all the different copies have to be compiled, and fixing a bug in the master source requires re-generating all the instances, some of which may be in different projects entirely.

或者，您可以写一个代码生成器，接受一个类型并为该类型的片断生成一个反向函数。现在有几个代码生成器就是这样做的。但这给每个需要Reverse的包增加了一个步骤，它使构建变得复杂，因为所有不同的副本都必须被编译，而且修复主源码中的错误需要重新生成所有的实例，其中一些可能完全在不同的项目中。

All these approaches are awkward enough that I think most people who have to reverse a slice in Go just write the function for the specific slice type that they need. Then they’ll need to write test cases for the function, to make sure they didn’t make a simple mistake like the one I made initially. And they’ll need to run those tests routinely.

所有这些方法都很笨拙，我认为大多数人在Go中要反转一个分片时，只需为他们需要的特定分片类型编写函数。然后他们需要为该函数编写测试用例，以确保他们没有犯像我最初犯的那样的简单错误。他们还需要定期地运行这些测试。

However we do it, it means a lot of extra work just for a function that looks exactly the same except for the element type. It’s not that it can’t be done. It clearly can be done, and Go programmers are doing it. It’s just that there ought to be a better way.

不管我们怎么做，这意味着仅仅为了一个除了元素类型之外看起来完全相同的函数，就需要做大量的额外工作。这并不是说它不能做。这显然是可以做到的，而且Go的程序员也在这样做。只是应该有一个更好的方法。

For a statically typed language like Go, that better way is generics. What I wrote earlier is that generic programming enables the representation of functions and data structures in a generic form, with types factored out. That’s exactly what we want here.

对于像Go这样的静态类型语言，更好的方法就是泛型。我在前面写到，泛型编程使函数和数据结构能够以泛型的形式表现出来，而类型则被剔除。这正是我们在这里想要的。

## What generics can bring to Go 泛型可以给Go带来什么

The first and most important thing we want from generics in Go is to be able to write functions like `Reverse` without caring about the element type of the slice. We want to factor out that element type. Then we can write the function once, write the tests once, put them in a go-gettable package, and call them whenever we want.

在Go中，我们希望从泛型中得到的第一件也是最重要的事情就是能够编写像Reverse这样的函数，而不关心分片的元素类型。我们想把这个元素类型剔除。这样我们就可以只写一次函数，只写一次测试，把它们放在go-gettable包里，然后随时调用它们。

Even better, since this is an open source world, someone else can write `Reverse` once, and we can use their implementation.

更好的是，由于这是一个开源的世界，别人可以写一次Reverse，而我们可以使用他们的实现。

At this point I should say that "generics" can mean a lot of different things. In this article, what I mean by "generics" is what I just described. In particular, I don’t mean templates as found in the C++ language, which support quite a bit more than what I’ve written here.

在这一点上，我应该说，"泛型 "可以有很多不同的含义。在这篇文章中，我所说的 "泛型 "就是我刚才描述的意思。特别是，我不是指C++语言中的模板，它支持的内容比我在这里写的多得多。

I went through `Reverse` in detail, but there are many other functions that we could write generically, such as:

我详细介绍了Reverse，但是还有很多其他的功能我们可以用泛型来写，比如说：

- Find smallest/largest element in slice 
- Find average/standard deviation of slice
- Compute union/intersection of maps
- Find shortest path in node/edge graph
- Apply transformation function to slice/map, returning new slice/map
- 查找切片中最小/最大的元素
  找出切片的平均/标准偏差
  计算地图的并集/交集
  寻找节点/边图中的最短路径
  对切片/地图应用转换函数，返回新的切片/地图

These examples are available in most other languages. In fact, I wrote this list by glancing at the C++ standard template library.

这些例子在大多数其他语言中都可以使用。事实上，我是通过浏览C++标准模板库来写这个列表的。

There are also examples that are specific to Go with its strong support for concurrency.

还有一些例子是Go特有的，它对并发性的支持很强。

- Read from a channel with a timeout
- Combine two channels into a single channel
- Call a list of functions in parallel, returning a slice of results
- Call a list of functions, using a Context, return the result of the first function to finish, canceling and cleaning up extra goroutines
- 从一个有超时的通道中读取数据
  将两个通道合并为一个通道
  并行调用一个函数列表，返回一个片断的结果
  调用一个函数列表，使用一个Context，返回第一个完成的函数的结果，取消并清理多余的goroutine

I’ve seen all of these functions written out many times with different types. It’s not hard to write them in Go. But it would be nice to be able to reuse an efficient and debugged implementation that works for any value type.

我已经看到所有这些函数用不同的类型写出来很多次了。用Go写这些东西并不难。但是，如果能够重复使用一个高效的、可调试的实现，并且适用于任何价值类型，那就更好了。

To be clear, these are just examples. There are many more general purpose functions that could be written more easily and safely using generics.

要清楚的是，这些只是例子。还有更多的通用函数可以使用泛型来更容易、更安全地编写。

Also, as I wrote earlier, it’s not just functions. It’s also data structures.

另外，正如我之前写的，这不仅仅是函数。它也是数据结构。

Go has two general purpose generic data structures built into the language: slices and maps. Slices and maps can hold values of any data type, with static type checking for values stored and retrieved. The values are stored as themselves, not as interface types. That is, when I have a `[]int`, the slice holds ints directly, not ints converted to an interface type.

Go语言中内置了两个通用的通用数据结构：切片和地图。分片和地图可以保存任何数据类型的值，对存储和检索的值进行静态类型检查。这些值是作为其本身而存储的，而不是作为接口类型。也就是说，当我有一个[]int的时候，切片直接持有ints，而不是转换为接口类型的ints。

Slices and maps are the most useful generic data structures, but they aren’t the only ones. Here are some other examples.

切片和地图是最有用的通用数据结构，但它们并不是唯一的。下面是一些其他的例子。

- Sets
- Self-balancing trees, with efficient insertion and traversal in sorted order
- Multimaps, with multiple instances of a key
- Concurrent hash maps, supporting parallel insertions and lookups with no single lock
- 集合
  自平衡树，具有高效的插入和遍历排序的功能
  多图，有一个键的多个实例
  并发哈希图，支持并行插入和查找，没有单锁。

If we can write generic types, we can define new data structures, like these, that have the same type-checking advantages as slices and maps: the compiler can statically type-check the types of the values that they hold, and the values can be stored as themselves, not as interface types.

如果我们可以编写通用类型，我们就可以定义新的数据结构，比如这些，它们具有与切片和地图相同的类型检查优势：编译器可以静态地对它们持有的值的类型进行检查，而且这些值可以作为它们自己而不是作为接口类型来存储。

It should also be possible to take algorithms like the ones mentioned earlier and apply them to generic data structures.

也应该可以采取像前面提到的那些算法，并将其应用于通用数据结构。

These examples should all be just like `Reverse`: generic functions and data structures written once, in a package, and reused whenever they are needed. They should work like slices and maps, in that they shouldn’t store values of empty interface type, but should store specific types, and those types should be checked at compile time.

这些例子都应该像Reverse一样：通用函数和数据结构只写一次，放在一个包里，在需要的时候重复使用。它们应该像切片和地图一样工作，因为它们不应该存储空接口类型的值，而应该存储特定的类型，而且这些类型应该在编译时被检查。

So that’s what Go can gain from generics. Generics can give us powerful building blocks that let us share code and build programs more easily.

所以这就是Go可以从泛型中获得的好处。泛型可以给我们提供强大的构件，让我们更容易地分享代码和构建程序。

I hope I’ve explained why this is worth looking into.

我希望我已经解释了为什么这值得研究。

## Benefits and costs 效益和成本

But generics don’t come from the [Big Rock Candy Mountain](https://mainlynorfolk.info/folk/songs/bigrockcandymountain.html), the land where the sun shines every day over the [lemonade springs](http://www.lat-long.com/Latitude-Longitude-773297-Montana-Lemonade_Springs.html). Every language change has a cost. There’s no doubt that adding generics to Go will make the language more complicated. As with any change to the language, we need to talk about maximizing the benefit and minimizing the cost.

但是，泛型并不是来自于大岩石糖果山，那片每天都有阳光照耀在柠檬水泉上的土地。每一种语言的改变都有成本。毫无疑问，在Go中添加泛型会使语言更加复杂。正如对语言的任何改变一样，我们需要讨论利益最大化和成本最小化的问题。

In Go, we’ve aimed to reduce complexity through independent, orthogonal language features that can be combined freely. We reduce complexity by making the individual features simple, and we maximize the benefit of the features by permitting their free combination. We want to do the same with generics.

在Go中，我们的目标是通过可以自由组合的独立、正交的语言特性来降低复杂性。我们通过使单个特征简单化来降低复杂度，并通过允许自由组合来最大化特征的利益。我们希望对泛型也这样做。

To make this more concrete I’m going to list a few guidelines we should follow.

为了使之更加具体，我将列出一些我们应该遵循的准则。

### Minimize new concepts 尽量减少新概念

We should add as few new concepts to the language as possible. That means a minimum of new syntax and a minimum of new keywords and other names.

我们应该尽可能少地在语言中添加新的概念。这意味着尽量少用新的语法，尽量少用新的关键字和其他名称。

### Complexity falls on the writer of generic code, not the user 复杂度落在通用代码的编写者身上，而不是用户身上

As much as possible the complexity should fall on the programmer writing the generic package. We don’t want the user of the package to have to worry about generics. This means that it should be possible to call generic functions in a natural way, and it means that any errors in using a generic package should be reported in a way that is easy to understand and to fix. It should also be easy to debug calls into generic code.

复杂性应尽可能地落在编写通用包的程序员身上。我们不希望包的用户不得不担心泛型的问题。这意味着应该能够以自然的方式调用泛型函数，也意味着使用泛型包的任何错误都应该以易于理解和修正的方式报告。调试对通用代码的调用也应该是容易的。

### Writer and user can work independently 编写者和使用者可以独立工作

Similarly, we should make it easy to separate the concerns of the writer of the generic code and its user, so that they can develop their code independently. They shouldn’t have to worry about what the other is doing, any more than the writer and caller of a normal function in different packages have to worry. This sounds obvious, but it’s not true of generics in every other programming language.

同样地，我们应该让泛型代码的编写者和它的使用者的关注点容易分离，这样他们就可以独立地开发他们的代码。他们不应该担心对方在做什么，就像不同包中的普通函数的编写者和调用者需要担心一样。这听起来很明显，但在其他每一种编程语言中的泛型都不是如此。

### Short build times, fast execution times 构建时间短，执行时间快

Naturally, as much as possible, we want to keep the short build times and fast execution time that Go gives us today. Generics tend to introduce a tradeoff between fast builds and fast execution. As much as possible, we want both.

自然地，我们希望尽可能地保持Go今天给我们带来的短构建时间和快执行时间。泛型往往会在快速构建和快速执行之间做出权衡。尽可能地，我们希望两者兼得。

### Preserve clarity and simplicity of Go 保留Go的清晰性和简单性

Most importantly, Go today is a simple language. Go programs are usually clear and easy to understand. A major part of our long process of exploring this space has been trying to understand how to add generics while preserving that clarity and simplicity. We need to find mechanisms that fit well into the existing language, without turning it into something quite different.

最重要的是，今天的Go是一种简单的语言。Go程序通常都很清晰，容易理解。我们在探索这一领域的漫长过程中，有一个主要部分是试图了解如何在保留这种清晰和简单的同时增加泛型。我们需要找到能够很好地适应现有语言的机制，而不至于把它变成完全不同的东西。

These guidelines should apply to any generics implementation in Go. That’s the most important message I want to leave you with today: **generics can bring a significant benefit to the language, but they are only worth doing if Go still feels like Go**.

这些准则应该适用于Go中任何泛型的实现。这就是我今天想留给大家的最重要的信息：泛型可以给语言带来很大的好处，但只有在Go仍然像Go的情况下才值得去做。

## Draft design 设计草案

Fortunately, I think it can be done. To finish up this article I’m going to shift from discussing why we want generics, and what the requirements on them are, to briefly discuss a design for how we think we can add them to the language.

幸运的是，我认为这是可以做到的。为了结束这篇文章，我将从讨论为什么我们要有泛型，以及对泛型的要求是什么，转到简要讨论我们认为可以如何将泛型加入语言的设计。

Note added January 2022: This blog post was written in 2019 and does not describe the version of generics that was finally adopted. For updated information please see [the language spec](https://go.dev/ref/spec) and [the generics design document] (https://go.dev/design/43651-type-parameters).

2022年1月添加的注释：这篇博文写于2019年，没有描述最终采用的泛型的版本。最新信息请参见语言规范和[泛型设计文档]（https://go.dev/design/43651-type-parameters）。

At this year’s Gophercon Robert Griesemer and I published [a design draft](https://github.com/golang/proposal/blob/master/design/go2draft-contracts.md) for adding generics to Go. See the draft for full details. I’ll go over some of the main points here.

在今年的Gophercon大会上，Robert Griesemer和我发表了一份在Go中加入泛型的设计草案。详细内容请见草案。我将在这里复述一些要点。

Here is the generic Reverse function in this design.

下面是这个设计中的通用反向函数。

```go linenums="1"
func Reverse (type Element) (s []Element) {
    first := 0
    last := len(s) - 1
    for first < last {
        s[first], s[last] = s[last], s[first]
        first++
        last--
    }
}
```

You’ll notice that the body of the function is exactly the same. Only the signature has changed.

您会注意到，这个函数的主体是完全一样的。只有签名发生了变化。

The element type of the slice has been factored out. It’s now named `Element` and has become what we call a *type parameter*. Instead of being part of the type of the slice parameter, it’s now a separate, additional, type parameter.

分片的元素类型已经被剔除。它现在被命名为元素，并成为我们所说的类型参数。它不再是分片参数类型的一部分，而是一个单独的、额外的类型参数。

To call a function with a type parameter, in the general case you pass a type argument, which is like any other argument except that it’s a type.

要调用一个带有类型参数的函数，在一般情况下，您要传递一个类型参数，它和其他参数一样，只是它是一个类型。

```go linenums="1"
func ReverseAndPrint(s []int) {
    Reverse(int)(s)
    fmt.Println(s)
}
```

That is the `(int)` seen after `Reverse` in this example.

这就是本例中Reverse后面看到的(int)。

Fortunately, in most cases, including this one, the compiler can deduce the type argument from the types of the regular arguments, and you don’t need to mention the type argument at all.

幸运的是，在大多数情况下，包括这个例子，编译器可以从常规参数的类型中推断出类型参数，而您根本不需要提及类型参数。

Calling a generic function just looks like calling any other function.

调用一个通用函数看起来就像调用其他函数一样。

```go linenums="1"
func ReverseAndPrint(s []int) {
    Reverse(s)
    fmt.Println(s)
}
```

In other words, although the generic `Reverse` function is slightly more complex than `ReverseInts` and `ReverseStrings`, that complexity falls on the writer of the function, not the caller.

换句话说，尽管通用的Reverse函数比ReverseInts和ReverseStrings稍微复杂一些，但这种复杂性落在函数的编写者身上，而不是调用者。

### Contracts 契约

Since Go is a statically typed language, we have to talk about the type of a type parameter. This *meta-type* tells the compiler what sorts of type arguments are permitted when calling a generic function, and what sorts of operations the generic function can do with values of the type parameter.

由于Go是一种静态类型的语言，我们必须谈论类型参数的类型。这个元类型告诉编译器在调用泛型函数时允许什么样的类型参数，以及泛型函数可以对类型参数的值做什么样的操作。



The `Reverse` function can work with slices of any type. The only thing it does with values of type `Element` is assignment, which works with any type in Go. For this kind of generic function, which is a very common case, we don’t need to say anything special about the type parameter.

反向函数可以对任何类型的片子进行操作。它对Element类型的值所做的唯一事情是赋值，这在Go中对任何类型都有效。对于这种通用函数，也就是非常常见的情况，我们不需要对类型参数做任何特殊说明。

Let’s take a quick look at a different function.

让我们快速看看一个不同的函数。

```go linenums="1"
func IndexByte (type T Sequence) (s T, b byte) int {
    for i := 0; i < len(s); i++ {
        if s[i] == b {
            return i
        }
    }
    return -1
}
```

Currently both the bytes package and the strings package in the standard library have an `IndexByte` function. This function returns the index of `b` in the sequence `s`, where `s` is either a `string` or a `[]byte`. We could use this single generic function to replace the two functions in the bytes and strings packages. In practice we may not bother doing that, but this is a useful simple example.

目前，标准库中的字节包和字符串包都有一个IndexByte函数。这个函数返回b在序列s中的索引，其中s是一个字符串或一个[]字节。我们可以用这个单一的通用函数来代替字节包和字符串包中的两个函数。在实践中，我们可能懒得这么做，但这是一个有用的简单例子。

Here we need to know that the type parameter `T` acts like a `string` or a `[]byte`. We can call `len` on it, and we can index to it, and we can compare the result of the index operation to a byte value.

在这里，我们需要知道类型参数T的行为就像一个字符串或[]字节。我们可以对它调用len，我们可以对它进行索引，我们可以将索引操作的结果与一个字节值进行比较。

To let this compile, the type parameter `T` itself needs a type. It’s a meta-type, but because we sometimes need to describe multiple related types, and because it describes a relationship between the implementation of the generic function and its callers, we actually call the type of `T` a contract. Here the contract is named `Sequence`. It appears after the list of type parameters.

为了让这个编译成功，类型参数T本身需要一个类型。它是一个元类型，但因为我们有时需要描述多个相关的类型，而且它描述了泛型函数的实现和它的调用者之间的关系，所以我们实际上把T的类型称为契约。这里的契约被命名为Sequence。它出现在类型参数列表的后面。

This is how the Sequence contract is defined for this example.

这就是本例中Sequence契约的定义方式。

```go linenums="1"
contract Sequence(T) {
    T string, []byte
}
```

It’s pretty simple, since this is a simple example: the type parameter `T` can be either `string` or `[]byte`. Here `contract` may be a new keyword, or a special identifier recognized in package scope; see the design draft for details.

这很简单，因为这是一个简单的例子：类型参数T可以是string或[]byte。这里的contract可以是一个新的关键字，也可以是一个在包范围内被认可的特殊标识符；详情见设计草案。

Anybody who remembers [the design we presented at Gophercon 2018](https://github.com/golang/proposal/blob/4a530dae40977758e47b78fae349d8e5f86a6c0a/design/go2draft-contracts.md) will see that this way of writing a contract is a lot simpler. We got a lot of feedback on that earlier design that contracts were too complicated, and we’ve tried to take that into account. The new contracts are much simpler to write, and to read, and to understand.

任何记得我们在Gophercon 2018上展示的设计的人都会发现，这种写契约的方式要简单得多。我们在早期的设计中得到了很多反馈，认为合同太复杂了，我们已经尽力考虑到了这一点。新的合同更容易写，更容易读，也更容易理解。

They let you specify the underlying type of a type parameter, and/or list the methods of a type parameter. They also let you describe the relationship between different type parameters.

它们让您指定一个类型参数的基本类型，和/或列出一个类型参数的方法。它们还可以让您描述不同类型参数之间的关系。

### Contracts with methods 带有方法的契约

Here is another simple example, of a function that uses the String method to return a `[]string` of the string representation of all the elements in `s`.

下面是另一个简单的例子，一个使用String方法来返回s中所有元素的字符串表示的[]字符串的函数。

```go linenums="1"
func ToStrings (type E Stringer) (s []E) []string {
    r := make([]string, len(s))
    for i, v := range s {
        r[i] = v.String()
    }
    return r
}
```

It’s pretty straightforward: walk through the slice, call the `String` method on each element, and return a slice of the resulting strings.

这是非常直接的：走过切片，在每个元素上调用String方法，并返回所得字符串的切片。

This function requires that the element type implement the `String` method. The Stringer contract ensures that.

这个函数要求元素类型实现String方法。Stringer合约确保了这一点。

```go linenums="1"
contract Stringer(T) {
    T String() string
}
```

The contract simply says that `T` has to implement the `String` method.

该契约简单地说，T必须实现String方法。

You may notice that this contract looks like the `fmt.Stringer` interface, so it’s worth pointing out that the argument of the `ToStrings` function is not a slice of `fmt.Stringer`. It’s a slice of some element type, where the element type implements `fmt.Stringer`. The memory representation of a slice of the element type and a slice of `fmt`.Stringer are normally different, and Go does not support direct conversions between them. So this is worth writing, even though `fmt.Stringer` exists.

您可能注意到这个契约看起来像fmt.Stringer接口，所以值得指出的是，ToStrings函数的参数不是fmt.Stringer的一个片断。它是某个元素类型的片断，其中的元素类型实现了fmt.Stringer。元素类型的片断和fmt.Stringer的片断的内存表示通常是不同的，并且Go不支持它们之间的直接转换。所以这值得一写，即使fmt.Stringer存在。

### Contracts with multiple types 具有多种类型的契约

Here is an example of a contract with multiple type parameters.

下面是一个具有多种类型参数的契约的例子。

```go linenums="1"
type Graph (type Node, Edge G) struct { ... }

contract G(Node, Edge) {
    Node Edges() []Edge
    Edge Nodes() (from Node, to Node)
}

func New (type Node, Edge G) (nodes []Node) *Graph(Node, Edge) {
    ...
}

func (g *Graph(Node, Edge)) ShortestPath(from, to Node) []Edge {
    ...
}
```

Here we’re describing a graph, built from nodes and edges. We’re not requiring a particular data structure for the graph. Instead, we’re saying that the `Node` type has to have an `Edges` method that returns the list of edges that connect to the `Node`. And the `Edge` type has to have a `Nodes` method that returns the two `Nodes` that the `Edge` connects.

这里我们描述的是一个图，由节点和边构成。我们并不要求该图有一个特定的数据结构。相反，我们说Node类型必须有一个Edges方法来返回连接到Node的边的列表。而Edge类型必须有一个Nodes方法来返回Edge所连接的两个Node。

I’ve skipped the implementation, but this shows the signature of a `New` function that returns a `Graph`, and the signature of a `ShortestPath` method on `Graph`.

我跳过了实现，但这显示了一个返回Graph的New函数的签名，以及Graph上ShortestPath方法的签名。

The important takeaway here is that a contract isn’t just about a single type. It can describe the relationships between two or more types.

这里的重要启示是，契约并不只是关于一个单一的类型。它可以描述两个或多个类型之间的关系。

### Ordered types 有序类型

One surprisingly common complaint about Go is that it doesn’t have a `Min` function. Or, for that matter, a `Max` function. That’s because a useful `Min` function should work for any ordered type, which means that it has to be generic.

关于Go的一个令人惊讶的普遍抱怨是它没有Min函数。或者，就这一点而言，没有Max函数。这是因为一个有用的Min函数应该适用于任何有序类型，这意味着它必须是通用的。

While `Min` is pretty trivial to write yourself, any useful generics implementation should let us add it to the standard library. This is what it looks like with our design.

虽然Min函数自己写起来很简单，但任何有用的泛型实现都应该让我们把它添加到标准库中。这就是我们的设计，看起来像这样。

```go linenums="1"
func Min (type T Ordered) (a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

The `Ordered` contract says that the type T has to be an ordered type, which means that it supports operators like less than, greater than, and so forth.

Ordered合约规定，类型T必须是一个有序类型，这意味着它支持小于、大于等运算符。

```go linenums="1"
contract Ordered(T) {
    T int, int8, int16, int32, int64,
        uint, uint8, uint16, uint32, uint64, uintptr,
        float32, float64,
        string
}
```

The `Ordered` contract is just a list of all the ordered types that are defined by the language. This contract accepts any of the listed types, or any named type whose underlying type is one of those types. Basically, any type you can use with the less than operator.

有序合约只是一个由语言定义的所有有序类型的列表。这个契约接受任何列出的类型，或者任何其底层类型是这些类型之一的命名类型。基本上，任何您可以使用小于运算符的类型。

It turns out that it’s much easier to simply enumerate the types that support the less than operator than it is to invent a new notation that works for all operators. After all, in Go, only built-in types support operators.

事实证明，简单列举支持小于运算符的类型要比发明一个适用于所有运算符的新符号容易得多。毕竟，在Go中，只有内置类型支持运算符。

This same approach can be used for any operator, or more generally to write a contract for any generic function intended to work with builtin types. It lets the writer of the generic function specify clearly the set of types the function is expected to be used with. It lets the caller of the generic function clearly see whether the function is applicable for the types being used.

这种方法也可以用于任何运算符，或者更广泛地用于为任何旨在与内置类型一起工作的通用函数编写契约。它可以让泛型函数的编写者清楚地指定该函数所要使用的类型集。它让泛型函数的调用者清楚地看到该函数是否适用于正在使用的类型。

In practice this contract would probably go into the standard library, and so really the `Min` function (which will probably also be in the standard library somewhere) will look like this. Here we’re just referring to the contract `Ordered` defined in the contracts package.

在实践中，这个契约可能会被放入标准库中，所以真正的Min函数（可能也会在标准库的某个地方）会是这样的。这里我们只是参考了定义在合约包中的Ordered合约。

```go linenums="1"
func Min (type T contracts.Ordered) (a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

### Generic data structures 通用数据结构

Finally, let’s look at a simple generic data structure, a binary tree. In this example the tree has a comparison function, so there are no requirements on the element type.

最后，我们来看看一个简单的通用数据结构，即二叉树。在这个例子中，树有一个比较函数，所以对元素的类型没有要求。

```go linenums="1"
type Tree (type E) struct {
    root    *node(E)
    compare func(E, E) int
}

type node (type E) struct {
    val         E
    left, right *node(E)
}
```

Here is how to create a new binary tree. The comparison function is passed to the `New` function.

这里是如何创建一个新的二叉树。比较函数被传递给New函数。

```go linenums="1"
func New (type E) (cmp func(E, E) int) *Tree(E) {
    return &Tree(E){compare: cmp}
}
```

An unexported method returns a pointer either to the slot holding v, or to the location in the tree where it should go.

一个未被导出的方法会返回一个指针，这个指针可以指向容纳v的槽，也可以指向树中它应该去的位置。

```go linenums="1"
func (t *Tree(E)) find(v E) **node(E) {
    pn := &t.root
    for *pn != nil {
        switch cmp := t.compare(v, (*pn).val); {
        case cmp < 0:
            pn = &(*pn).left
        case cmp > 0:
            pn = &(*pn).right
        default:
            return pn
        }
    }
    return pn
}
```

The details here don’t really matter, especially since I haven’t tested this code. I’m just trying to show what it looks like to write a simple generic data structure.

这里的细节其实并不重要，尤其是我还没有测试这段代码。我只是想说明写一个简单的通用数据结构是什么样子的。

This is the code for testing whether the tree contains a value.

这是测试树是否包含一个值的代码。

```go linenums="1"
func (t *Tree(E)) Contains(v E) bool {
    return *t.find(e) != nil
}
```

This is the code for inserting a new value.

这是插入一个新值的代码。

```go linenums="1"
func (t *Tree(E)) Insert(v E) bool {
    pn := t.find(v)
    if *pn != nil {
        return false
    }
    *pn = &node(E){val: v}
    return true
}
```

Notice that the type `node` has a type argument `E`. This is what it looks like to write a generic data structure. As you can see, it looks like writing ordinary Go code, except that some type arguments are sprinkled in here and there.

注意，类型节点有一个类型参数E，这就是写一个通用数据结构的样子。正如您所看到的，它看起来就像在写普通的Go代码，只是在这里和那里洒上了一些类型参数。

Using the tree is pretty simple.

使用树是非常简单的。

```go linenums="1"
var intTree = tree.New(func(a, b int) int { return a - b })

func InsertAndCheck(v int) {
    intTree.Insert(v)
    if !intTree.Contains(v) {
        log.Fatalf("%d not found after insertion", v)
    }
}
```

That’s as it should be. It’s a bit harder to write a generic data structure, because you often have to explicitly write out type arguments for supporting types, but as much as possible using one is no different from using an ordinary non-generic data structure.

这就是它应该有的样子。编写通用数据结构有点困难，因为您经常要明确地写出支持类型的参数，但尽可能地使用通用数据结构与使用普通的非通用数据结构没有区别。

### Next steps 接下来的步骤

We are working on actual implementations to allow us to experiment with this design. It’s important to be able to try out the design in practice, to make sure that we can write the kinds of programs we want to write. It hasn’t gone as fast as we’d hoped, but we’ll send out more detail on these implementations as they become available.

我们正在进行实际的实现工作，以使我们能够对这种设计进行实验。能够在实践中尝试这个设计是很重要的，以确保我们能够写出我们想写的那种程序。这并没有像我们希望的那样快，但是当这些实现出现时，我们会发出更多关于这些实现的细节。

Robert Griesemer has written a [preliminary CL](https://go.dev/cl/187317) that modifies the go/types package. This permits testing whether code using generics and contracts can type check. It’s incomplete right now, but it mostly works for a single package, and we’ll keep working on it.

Robert Griesemer写了一个初步的CL，修改了go/types包。这允许测试使用泛型和契约的代码是否可以进行类型检查。它现在还不完整，但对于一个单一的包来说，它基本上是有效的，我们会继续努力。

What we’d like people to do with this and future implementations is to try writing and using generic code and see what happens. We want to make sure that people can write the code they need, and that they can use it as expected. Of course not everything is going to work at first, and as we explore this space we may have to change things. And, to be clear, we’re much more interested in feedback on the semantics than on details of the syntax.

我们希望人们用这个和未来的实现来做的是，尝试编写和使用泛型代码，看看会发生什么。我们想确保人们能够写出他们需要的代码，并且能够按照预期使用。当然，一开始并不是所有的东西都能成功，随着我们对这个空间的探索，我们可能不得不改变一些东西。而且，明确地说，我们对语义的反馈比对语法的细节更感兴趣。

I’d like to thank everyone who commented on the earlier design, and everyone who has discussed what generics can look like in Go. We’ve read all of the comments, and we greatly appreciate the work that people have put into this. We would not be where we are today without that work.

我想感谢每一个对早期设计提出意见的人，以及每一个讨论过泛型在Go中可以是什么样子的人。我们已经阅读了所有的评论，并且我们非常感谢大家在这方面所做的工作。没有这些工作，我们就不会有今天的成就。

Our goal is to arrive at a design that makes it possible to write the kinds of generic code I’ve discussed today, without making the language too complex to use or making it not feel like Go anymore. We hope that this design is a step toward that goal, and we expect to continue to adjust it as we learn, from our experiences and yours, what works and what doesn’t. If we do reach that goal, then we’ll have something that we can propose for future versions of Go.

我们的目标是达成一个设计，使我今天所讨论的那种泛型代码的编写成为可能，同时又不会使语言的使用变得过于复杂，或者使它不再有Go的感觉。我们希望这个设计是朝着这个目标迈出的一步，并且我们希望在从我们和您们的经验中了解到哪些是可行的，哪些是不可行的时候继续调整它。如果我们真的达到了这个目标，那么我们就可以为未来的Go版本提出一些建议。
