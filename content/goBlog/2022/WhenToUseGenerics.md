+++
title = "何时使用泛型"
weight = 92
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# When To Use Generics - 何时使用泛型

https://go.dev/blog/when-generics

Ian Lance Taylor
12 April 2022

## 简介 Introduction 

This is the blog post version of my talks at Google Open Source Live:

​	这是我在Google Open Source Live

<iframe src="https://www.youtube.com/embed/nr8EpUO9jhw" width="560" height="315" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

和GopherCon 2021的演讲的博文版本：

<iframe src="https://www.youtube.com/embed/Pa_e9EeCdy8?start=1244" width="560" height="315" frameborder="0" allowfullscreen="" mozallowfullscreen="" webkitallowfullscreen="" style="box-sizing: border-box;"></iframe>

The Go 1.18 release adds a major new language feature: support for generic programming. In this article I’m not going to describe what generics are nor how to use them. This article is about when to use generics in Go code, and when not to use them.

​	Go 1.18版本增加了一个重要的新语言特性：对泛型编程的支持。在这篇文章中，我不打算描述什么是泛型，也不打算描述如何使用它们。这篇文章是关于在Go代码中何时使用泛型，以及何时不使用泛型。

To be clear, I’ll provide general guidelines, not hard and fast rules. Use your own judgement. But if you aren’t sure, I recommend using the guidelines discussed here.

​	为了明确起见，我将提供一般性的指导方针，而不是硬性规定。请根据自己的判断进行决策。但如果你不确定，我建议遵循这里讨论的指导方针。

## 编写代码 Write code 

Let’s start with a general guideline for programming Go: write Go programs by writing code, not by defining types. When it comes to generics, if you start writing your program by defining type parameter constraints, you are probably on the wrong path. Start by writing functions. It’s easy to add type parameters later when it’s clear that they will be useful.

​	让我们从编写 Go 程序的一般性指导方针开始：通过编写代码来编写 Go 程序，而不是通过定义类型。在使用泛型时，如果你开始通过定义类型参数约束来编写程序，那么你可能走错了方向。应该先编写函数。当确定它们会有用时，稍后再添加类型参数是很容易的。

## 类型参数什么时候有用？ When are type parameters useful? 

That said, let’s look at cases for which type parameters can be useful.

​	话虽如此，让我们来看看类型参数可能有用的情况。

### 使用语言定义的容器类型时 When using language-defined container types 

One case is when writing functions that operate on the special container types that are defined by the language: slices, maps, and channels. If a function has parameters with those types, and the function code doesn’t make any particular assumptions about the element types, then it may be useful to use a type parameter.

​	其中一种情况是在编写操作语言定义的特殊容器类型的函数时，如切片、映射和通道。如果一个函数具有这些类型的参数，并且函数代码对元素类型没有特定的假设，那么使用类型参数可能会很有用。

For example, here is a function that returns a slice of all the keys in a map of any type:

​	例如，下面是一个函数，它返回任意类型的映射中所有键的切片：

```go
// MapKeys returns a slice of all the keys in m.
// The keys are not returned in any particular order.
func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
    s := make([]Key, 0, len(m))
    for k := range m {
        s = append(s, k)
    }
    return s
}
```

This code doesn’t assume anything about the map key type, and it doesn’t use the map value type at all. It works for any map type. That makes it a good candidate for using type parameters.

​	这段代码对映射的键类型不做任何假设，也完全不使用映射的值类型。它适用于任何映射类型。因此，它是使用类型参数的一个很好的候选项。

The alternative to type parameters for this kind of function is typically to use reflection, but that is a more awkward programming model, is not statically typechecked at build time, and is often slower at run time.

​	对于这种类型的函数，使用类型参数的替代方案通常是使用反射，但这是一种更为繁琐的编程模型，在构建时没有静态类型检查，并且在运行时通常更慢。

### 通用的数据结构 General purpose data structures 

Another case where type parameters can be useful is for general purpose data structures. A general purpose data structure is something like a slice or map, but one that is not built into the language, such as a linked list, or a binary tree.

​	类型参数还可以在通用数据结构中发挥作用。通用数据结构指的是类似于切片或映射的结构，但并非（类似链表或二叉树）内置于语言中。

Today, programs that need such data structures typically do one of two things: write them with a specific element type, or use an interface type. Replacing a specific element type with a type parameter can produce a more general data structure that can be used in other parts of the program, or by other programs. Replacing an interface type with a type parameter can permit data to be stored more efficiently, saving memory resources; it can also permit the code to avoid type assertions, and to be fully type checked at build time.

​	目前，需要使用此类数据结构的程序通常有两种做法：使用特定的元素类型编写它们，或者使用接口类型。用类型参数替换特定的元素类型可以产生更通用的数据结构，可以在程序的其他部分或其他程序中使用。用类型参数替换接口类型可以使数据存储更高效，节省内存资源；它还可以避免类型断言，并且在构建时进行完全的类型检查。

For example, here is part of what a binary tree data structure might look like using type parameters:

​	例如，使用类型参数的二叉树数据结构的一部分可能如下所示：

```go
// Tree is a binary tree.
type Tree[T any] struct {
    cmp  func(T, T) int
    root *node[T]
}

// A node in a Tree.
type node[T any] struct {
    left, right  *node[T]
    val          T
}

// find returns a pointer to the node containing val,
// or, if val is not present, a pointer to where it
// would be placed if added.
func (bt *Tree[T]) find(val T) **node[T] {
    pl := &bt.root
    for *pl != nil {
        switch cmp := bt.cmp(val, (*pl).val); {
        case cmp < 0:
            pl = &(*pl).left
        case cmp > 0:
            pl = &(*pl).right
        default:
            return pl
        }
    }
    return pl
}

// Insert inserts val into bt if not already there,
// and reports whether it was inserted.
func (bt *Tree[T]) Insert(val T) bool {
    pl := bt.find(val)
    if *pl != nil {
        return false
    }
    *pl = &node[T]{val: val}
    return true
}
```

Each node in the tree contains a value of the type parameter `T`. When the tree is instantiated with a particular type argument, values of that type will be stored directly in the nodes. They will not be stored as interface types.

​	树中的每个节点都包含类型参数 `T` 的值。当使用特定的类型参数实例化树时，该类型的值将直接存储在节点中，而不会作为接口类型存储。

This is a reasonable use of type parameters because the `Tree` data structure, including the code in the methods, is largely independent of the element type `T`.

​	这是一种合理的类型参数使用方式，因为`Tree` 数据结构，包括方法中的代码，大部分与元素类型 `T` 无关。

The `Tree` data structure does need to know how to compare values of the element type `T`; it uses a passed-in comparison function for that. You can see this on the fourth line of the `find` method, in the call to `bt.cmp`. Other than that, the type parameter doesn’t matter at all.

​	`Tree` 数据结构需要知道如何比较元素类型 `T` 的值；它使用传入的比较函数来实现。可以在 `find` 方法的第四行看到这一点，在调用 `bt.cmp` 时。除此之外，类型参数并不重要。

### 对于类型参数，更倾向于使用函数而不是方法 For type parameters, prefer functions to methods 

The `Tree` example illustrates another general guideline: when you need something like a comparison function, prefer a function to a method.

​	Tree的例子说明了另一条通用准则：当您需要类似于比较函数的东西时，最好使用函数而不是方法。

We could have defined the `Tree` type such that the element type is required to have a `Compare` or `Less` method. This would be done by writing a constraint that requires the method, meaning that any type argument used to instantiate the `Tree` type would need to have that method.

​	我们可以定义Tree类型，要求元素类型有一个Compare或Less方法。这将通过编写一个需要该方法的约束来实现，这意味着任何用于实例化Tree类型的类型参数都需要有该方法。

A consequence would be that anybody who wants to use `Tree` with a simple data type like `int` would have to define their own integer type and write their own comparison method. If we define `Tree` to take a comparison function, as in the code shown above, then it is easy to pass in the desired function. It’s just as easy to write that comparison function as it is to write a method.

​	这样做的后果是，任何想用int这样的简单数据类型来使用Tree的人都必须定义他们自己的整数类型并编写他们自己的比较方法。如果我们将Tree定义为接受一个比较函数，如上图所示的代码，那么就很容易传入想要的函数。写那个比较函数和写一个方法一样容易。

If the `Tree` element type happens to already have a `Compare` method, then we can simply use a method expression like `ElementType.Compare` as the comparison function.

​	如果Tree元素类型刚好已经有一个Compare方法，那么我们可以简单地使用ElementType.Compare这样的方法表达式作为比较函数。

To put it another way, it is much simpler to turn a method into a function than it is to add a method to a type. So for general purpose data types, prefer a function rather than writing a constraint that requires a method.

​	换句话说，把一个方法变成一个函数比给一个类型添加一个方法要简单得多。所以对于通用的数据类型，宁愿选择一个函数而不是写一个需要方法的约束。

### 实现一个通用方法 Implementing a common method 

Another case where type parameters can be useful is when different types need to implement some common method, and the implementations for the different types all look the same.

​	类型参数有用的另一种情况是，不同的类型需要实现一些共同的方法，而不同类型的实现看起来都是一样的。

For example, consider the standard library’s `sort.Interface`. It requires that a type implement three methods: `Len`, `Swap`, and `Less`.

​	例如，考虑标准库的sort.Interface。它要求一个类型实现三个方法。Len, Swap, and Less。

Here is an example of a generic type `SliceFn` that implements `sort.Interface` for any slice type:

​	下面是一个通用类型 SliceFn 的例子，它为任何片断类型实现了 sort.Interface：

```go
// SliceFn implements sort.Interface for a slice of T.
type SliceFn[T any] struct {
    s    []T
    less func(T, T) bool
}

func (s SliceFn[T]) Len() int {
    return len(s.s)
}
func (s SliceFn[T]) Swap(i, j int) {
    s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s SliceFn[T]) Less(i, j int) bool {
    return s.less(s.s[i], s.s[j])
}
```

For any slice type, the `Len` and `Swap` methods are exactly the same. The `Less` method requires a comparison, which is the `Fn` part of the name `SliceFn`. As with the earlier `Tree` example, we will pass in a function when we create a `SliceFn`.

​	对于任何切片类型，Len和Swap方法都是完全一样的。Less方法需要进行比较，也就是SliceFn这个名字中的Fn部分。和前面的Tree例子一样，我们在创建SliceFn时将传入一个函数。

Here is how to use `SliceFn` to sort any slice using a comparison function:

​	下面是如何使用SliceFn来对任何切片进行比较函数的排序：

```go
// SortFn sorts s in place using a comparison function.
func SortFn[T any](s []T, less func(T, T) bool) {
    sort.Sort(SliceFn[T]{s, less})
}
```

This is similar to the standard library function `sort.Slice`, but the comparison function is written using values rather than slice indexes.

​	这与标准库函数sort.Slice类似，但比较函数是用值而不是片状索引来写的。

Using type parameters for this kind of code is appropriate because the methods look exactly the same for all slice types.

​	在这种代码中使用类型参数是合适的，因为这些方法对于所有的片断类型看起来都是一样的。

(I should mention that Go 1.19–not 1.18–will most likely include a generic function to sort a slice using a comparison function, and that generic function will most likely not use `sort.Interface`. See [proposal #47619](https://go.dev/issue/47619). But the general point is still true even if this specific example will most likely not be useful: it’s reasonable to use type parameters when you need to implement methods that look the same for all the relevant types.)

​	(我应该提到，Go 1.19--而不是1.18--很可能包括一个使用比较函数对切片进行排序的通用函数，而这个通用函数很可能不会使用sort.Interface。参见提议#47619。但是，即使这个具体的例子很可能没有用，一般的观点仍然是正确的：当您需要实现对所有相关类型看起来都一样的方法时，使用类型参数是合理的）。

## 什么时候类型参数没有用？ When are type parameters not useful? 

Now let’s talk about the other side of the question: when not to use type parameters.

现在我们来谈谈问题的另一面：什么时候不使用类型参数。

### 不要用类型参数代替接口类型 Don’t replace interface types with type parameters 

As we all know, Go has interface types. Interface types permit a kind of generic programming.

​	我们都知道，Go有接口类型。接口类型允许一种通用编程。

For example, the widely used `io.Reader` interface provides a generic mechanism for reading data from any value that contains information (for example, a file) or that produces information (for example, a random number generator). If all you need to do with a value of some type is call a method on that value, use an interface type, not a type parameter. `io.Reader` is easy to read, efficient, and effective. There is no need to use a type parameter to read data from a value by calling the `Read` method.

​	例如，广泛使用的io.Reader接口提供了一种通用机制，可以从任何包含信息（例如文件）或产生信息（例如随机数发生器）的值中读取数据。如果您需要对某个类型的值进行处理，只是在该值上调用一个方法，那么就使用一个接口类型，而不是一个类型参数。io.Reader很容易阅读，效率高，效果好。没有必要使用类型参数来通过调用Read方法从一个值中读取数据。

For example, it might be tempting to change the first function signature here, which uses just an interface type, into the second version, which uses a type parameter.

​	例如，把这里的第一个函数签名（只使用接口类型）改成第二个版本（使用类型参数），可能会很有吸引力。

```go
func ReadSome(r io.Reader) ([]byte, error)

func ReadSome[T io.Reader](r T) ([]byte, error)
```

Don’t make that kind of change. Omitting the type parameter makes the function easier to write, easier to read, and the execution time will likely be the same.

​	不要做这样的改变。省略类型参数使函数更容易编写，更容易阅读，而且执行时间很可能是一样的。

It’s worth emphasizing the last point. While it’s possible to implement generics in several different ways, and implementations will change and improve over time, the implementation used in Go 1.18 will in many cases treat values whose type is a type parameter much like values whose type is an interface type. What this means is that using a type parameter will generally not be faster than using an interface type. So don’t change from interface types to type parameters just for speed, because it probably won’t run any faster.

值得强调的是最后一点。虽然有可能以几种不同的方式实现泛型，而且实现方式会随着时间的推移而改变和改进，但Go 1.18中使用的实现方式在很多情况下会将类型为类型参数的值与类型为接口类型的值一样对待。这意味着，使用类型参数通常不会比使用接口类型更快。所以不要仅仅为了速度而从接口类型改为类型参数，因为它可能不会运行得更快。

### 如果方法的实现不同，不要使用类型参数 Don’t use type parameters if method implementations differ 

When deciding whether to use a type parameter or an interface type, consider the implementation of the methods. Earlier we said that if the implementation of a method is the same for all types, use a type parameter. Inversely, if the implementation is different for each type, then use an interface type and write different method implementations, don’t use a type parameter.

​	在决定是使用类型参数还是接口类型时，要考虑方法的实现。前面我们说过，如果一个方法的实现对所有类型都是一样的，那么就使用类型参数。反之，如果每种类型的实现都不同，那么就使用接口类型，编写不同的方法实现，不要使用类型参数。

For example, the implementation of `Read` from a file is nothing like the implementation of `Read` from a random number generator. That means that we should write two different `Read` methods, and use an interface type like `io.Reader`.

例如，从文件中读取的实现与从随机数发生器中读取的实现完全不同。这意味着我们应该写两个不同的Read方法，并使用一个接口类型，如io.Reader。

### 在适当的地方使用反射 Use reflection where appropriate 

Go has [run time reflection](https://pkg.go.dev/reflect). Reflection permits a kind of generic programming, in that it permits you to write code that works with any type.

​	Go有运行时反射。反射允许一种通用编程，因为它允许您编写适用于任何类型的代码。

If some operation has to support even types that don’t have methods (so that interface types don’t help), and if the operation is different for each type (so that type parameters aren’t appropriate), use reflection.

​	如果某些操作必须支持没有方法的类型（因此接口类型没有帮助），并且如果操作对每个类型都不同（因此类型参数不合适），请使用反射。

An example of this is the [encoding/json](https://pkg.go.dev/encoding/json) package. We don’t want to require that every type that we encode have a `MarshalJSON` method, so we can’t use interface types. But encoding an interface type is nothing like encoding a struct type, so we shouldn’t use type parameters. Instead, the package uses reflection. The code is not simple, but it works. For details, see [the source code](https://go.dev/src/encoding/json/encode.go).

​	这方面的一个例子是编码/json包。我们不想要求我们编码的每个类型都有一个MarshalJSON方法，所以我们不能使用接口类型。但对接口类型的编码与对结构类型的编码完全不同，所以我们不应该使用类型参数。相反，该包使用反射。这段代码并不简单，但它是有效的。详情请看源代码。

## 一个简单的指导原则 One simple guideline 

In closing, this discussion of when to use generics can be reduced to one simple guideline.

​	最后，关于何时使用泛型的讨论可以简化为一个简单的准则。

If you find yourself writing the exact same code multiple times, where the only difference between the copies is that the code uses different types, consider whether you can use a type parameter.

​	如果您发现自己多次编写完全相同的代码，其中唯一的区别是代码使用了不同的类型，请考虑是否可以使用一个类型参数。

Another way to say this is that you should avoid type parameters until you notice that you are about to write the exact same code multiple times.

​	另一种说法是，您应该避免使用类型参数，直到您注意到您即将多次编写完全相同的代码。
