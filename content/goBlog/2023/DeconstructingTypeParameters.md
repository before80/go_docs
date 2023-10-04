+++
title = "解构类型参数"
date = 2023-10-04T14:33:20+08:00
type = "docs"
weight = 79
description = ""
isCJKLanguage = true
draft = false

+++

# Deconstructing Type Parameters - 解构类型参数

> 原文：[https://go.dev/blog/deconstructing-type-parameters](https://go.dev/blog/deconstructing-type-parameters)

Ian Lance Taylor
26 September 2023

Ian Lance Taylor
2023年9月26日

## slices包函数签名 slices package function signatures

The [`slices.Clone`](https://pkg.go.dev/slices#Clone) function is pretty simple: it makes a copy of a slice of any type.

[	 `slices.Clone` ](https://pkg.go.dev/slices#Clone)函数非常简单：它复制任意类型的切片。

```Go
func Clone[S ~[]E, E any](s S) S {
    return append(s[:0:0], s...)
}
```

This works because appending to a slice with zero capacity will allocate a new backing array. The function body winds up being shorter than the function signature, which is in part because the body is short, but also because the signature is long. In this blog post we’ll explain why the signature is written the way that it is.

​	这个函数的工作原理是，将一个容量为零的切片追加到切片中会分配一个新的后备数组。函数体比函数签名要短，部分原因是函数体很短，但也因为函数签名很长。在本博文中，我们将解释为什么要以这种方式编写函数签名。

## 简单的Clone函数 Simple Clone

We’ll start by writing a simple generic `Clone` function. This is not the one in the `slices` package. We want to take a slice of any element type, and return a new slice.

​	我们首先编写一个简单的泛型 `Clone` 函数。这不是 `slices` 包中的函数。我们想要获取任意元素类型的切片，并返回一个新的切片。

```Go
func Clone1[E any](s []E) []E {
    // body omitted
}
```

The generic function `Clone1` has a single type parameter `E`. It takes a single argument `s` which is a slice of type `E`, and it returns a slice of the same type. This signature is straightforward for anybody familiar with generics in Go.

​	泛型函数 `Clone1` 有一个类型参数 `E` 。它接受一个名为 `s` 的参数，该参数是类型为 `E` 的切片，并返回一个相同类型的切片。对于熟悉Go中泛型的人来说，这个签名很简单明了。

However, there is a problem. Named slice types are not common in Go, but people do use them.

​	然而，这里存在一个问题。在Go中，命名的切片类型并不常见，但人们确实在使用它们。

```Go
// MySlice is a slice of strings with a special String method.
// MySlice是一个具有特殊String方法的字符串切片。
type MySlice []string

// String returns the printable version of a MySlice value.
// String返回MySlice值的可打印版本。
func (s MySlice) String() string {
    return strings.Join(s, "+")
}
```

Let’s say that we want to make a copy of a `MySlice` and then get the printable version, but with the strings in sorted order.

​	假设我们想要复制一个 `MySlice` ，然后按排序顺序获取可打印版本的字符串切片。

```Go
func PrintSorted(ms MySlice) string {
    c := Clone1(ms)
    slices.Sort(c)
    return c.String() // FAILS TO COMPILE 无法编译通过
}
```

Unfortunately, this doesn’t work. The compiler reports an error:

​	不幸的是，这样做不起作用。编译器报告了一个错误：

```
c.String undefined (type []string has no field or method String)
```

We can see the problem if we manually instantiate `Clone1` by replacing the type parameter with the type argument.

​	如果我们手动实例化 `Clone1` ，将类型参数替换为类型参数，我们就可以看到问题。

```Go
func InstantiatedClone1(s []string) []string
```

The [Go assignment rules](https://go.dev/ref/spec#Assignability) allow us to pass a value of type `MySlice` to a parameter of type `[]string`, so calling `Clone1` is fine. But `Clone1` will return a value of type `[]string`, not a value of type `MySlice`. The type `[]string` doesn’t have a `String` method, so the compiler reports an error.

​	[Go的赋值规则](https://go.dev/ref/spec#Assignability)允许我们将类型为 `MySlice` 的值传递给类型为 `[]string` 的参数，因此调用 `Clone1` 是可以的。但是， `Clone1` 将返回一个类型为 `[]string` 的值，而不是 `MySlice` 类型的值。类型 `[]string` 没有 `String` 方法，因此编译器报告了一个错误。

## 灵活的Clone函数 Flexible Clone

To fix this problem, we have to write a version of `Clone` that returns the same type as its argument. If we can do that, then when we call `Clone` with a value of type `MySlice`, it will return a result of type `MySlice`.

​	为了解决这个问题，我们必须编写一个返回与其参数相同类型的 `Clone` 版本。如果我们能做到这一点，那么当我们使用类型为 `MySlice` 的值调用 `Clone` 时，它将返回一个类型为 `MySlice` 的结果。

We know that it has to look something like this.

​	我们知道它必须看起来像这样。

```Go
func Clone2[S ?](s S) S // INVALID
```

This `Clone2` function returns a value that is the same type as its argument.

​	这个 `Clone2` 函数返回一个与其参数类型相同的值。

Here I’ve written the constraint as `?`, but that’s just a placeholder. To make this work we need to write a constraint that will let us write the body of the function. For `Clone1` we could just use a constraint of `any` for the element type. For `Clone2` that won’t work: we want to require that `s` be a slice type.

​	在这里，我将约束写为 `?` ，但那只是一个占位符。为了使其工作，我们需要编写一个约束，使我们能够编写函数的主体。对于 `Clone1` ，我们可以只使用 `any` 作为元素类型的约束。对于 `Clone2` ，这样做行不通：我们希望要求 `s` 是一个切片类型。

Since we know we want a slice, the constraint of `S` has to be a slice. We don’t care what the slice element type is, so let’s just call it `E`, as we did with `Clone1`.

​	由于我们知道我们想要一个切片，约束 `S` 必须是一个切片。我们不关心切片元素类型是什么，所以让我们像在 `Clone1` 中一样称之为 `E` 。

```Go
func Clone3[S []E](s S) S // INVALID
```

This is still invalid, because we haven’t declared `E`. The type argument for `E` can be any type, which means it also has to be a type parameter itself. Since it can be any type, its constraint is `any`.

​	这仍然是无效的，因为我们没有声明 `E` 。对于 `E` 的类型参数可以是任何类型，这意味着它本身也必须是一个类型参数。由于它可以是任何类型，所以它的约束是 `any` 。

```Go
func Clone4[S []E, E any](s S) S
```

This is getting close, and at least it will compile, but we’re not quite there yet. If we compile this version, we get an error when we call `Clone4(ms)`.

​	这接近了，至少它可以编译通过，但我们还没有到达目标。如果我们编译这个版本，当我们调用 `Clone4(ms)` 时会出现错误。

```
MySlice does not satisfy []string (possibly missing ~ for []string in []string)
```

The compiler is telling us that we can’t use the type argument `MySlice` for the type parameter `S`, because `MySlice` does not satisfy the constraint `[]E`. That’s because `[]E` as a constraint only permits a slice type literal, like `[]string`. It doesn’t permit a named type like `MySlice`.

​	编译器告诉我们，我们不能将类型参数 `MySlice` 用于类型参数 `S` ，因为 `MySlice` 不满足约束 `[]E` 。这是因为 `[]E` 作为约束只允许切片类型字面量，比如 `[]string` 。它不允许像 `MySlice` 这样的命名类型。

## 底层类型约束 Underlying type constraints

As the error message hints, the answer is to add a `~`.

​	正如错误消息所暗示的，答案是添加一个 `~` 。

```Go
func Clone5[S ~[]E, E any](s S) S
```

To repeat, writing type parameters and constraints `[S []E, E any]` means that the type argument for `S` can be any unnamed slice type, but it can’t be a named type defined as a slice literal. Writing `[S ~[]E, E any]`, with a `~`, means that the type argument for `S` can be any type whose underlying type is a slice type.

​	重申一遍，编写类型参数和约束 `[S []E, E any]` 的意思是，类型参数 `S` 的类型参数可以是任何未命名的切片类型，但不能是命名类型，定义为切片字面量。写 `[S ~[]E, E any]` ，带有一个 `~` ，意味着类型参数 `S` 的类型参数可以是任何底层类型为切片类型的类型。

For any named type `type T1 T2` the underlying type of `T1` is the underlying type of `T2`. The underlying type of a predeclared type like `int` or a type literal like `[]string` is just the type itself. For the exact details, [see the language spec](https://go.dev/ref/spec#Underlying_types). In our example, the underlying type of `MySlice` is `[]string`.

​	对于任何命名类型 `type T1 T2` ， `T1` 的底层类型是 `T2` 的底层类型。预声明类型的底层类型，如 `int` 或类型字面量 `[]string` ，就是类型本身。有关详细信息，请参阅[语言规范](https://go.dev/ref/spec#Underlying_types)。在我们的示例中， `MySlice` 的底层类型是 `[]string` 。

Since the underlying type of `MySlice` is a slice, we can pass an argument of type `MySlice` to `Clone5`. As you may have noticed, the signature of `Clone5` is the same as the signature of `slices.Clone`. We’ve finally gotten to where we want to be.

​	由于 `MySlice` 的底层类型是切片，我们可以将类型为 `MySlice` 的参数传递给 `Clone5` 。正如你可能已经注意到的， `Clone5` 的签名与 `slices.Clone` 的签名相同。我们终于达到了我们想要的地方。

Before we move on, let’s discuss why the Go syntax requires a `~`. It might seem that we would always want to permit passing `MySlice`, so why not make that the default? Or, if we need to support exact matching, why not flip things around, so that a constraint of `[]E` permits a named type while a constraint of, say, `=[]E`, only permits slice type literals?

​	在我们继续之前，让我们讨论一下为什么Go语法需要一个 `~` 。它可能看起来我们总是希望允许传递 `MySlice` ，那么为什么不将其作为默认值呢？或者，如果我们需要支持精确匹配，为什么不颠倒一下，使得约束 `[]E` 允许命名类型，而约束 `=[]E` 只允许切片类型字面量？

To explain this, let’s first observe that a type parameter list like `[T ~MySlice]` doesn’t make sense. That’s because `MySlice` is not the underlying type of any other type. For instance, if we have a definition like `type MySlice2 MySlice`, the underlying type of `MySlice2` is `[]string`, not `MySlice`. So either `[T ~MySlice]` would permit no types at all, or it would be the same as `[T MySlice]` and only match `MySlice`. Either way, `[T ~MySlice]` isn’t useful. To avoid this confusion, the language prohibits `[T ~MySlice]`, and the compiler produces an error like

​	为了解释这个问题，首先让我们观察到一个类型参数列表，比如 `[T ~MySlice]` 是没有意义的。那是因为 `MySlice` 不是任何其他类型的底层类型。例如，如果我们有一个定义 `type MySlice2 MySlice` ，那么 `MySlice2` 的底层类型是 `[]string` ，而不是 `MySlice` 。因此，要么 `[T ~MySlice]` 不允许任何类型，要么它与 `[T MySlice]` 相同，只匹配 `MySlice` 。无论哪种方式， `[T ~MySlice]` 都没有用。为了避免这种混淆，语言禁止了 `[T ~MySlice]` ，编译器会产生类似于

```
invalid use of ~ (underlying type of MySlice is []string)
```

的错误。

If Go didn’t require the tilde, so that `[S []E]` would match any type whose underlying type is `[]E`, then we would have to define the meaning of `[S MySlice]`.

​	如果Go不要求波浪号，那么 `[S []E]` 将匹配任何底层类型为 `[]E` 的类型，我们将不得不定义 `[S MySlice]` 的含义。

We could prohibit `[S MySlice]`, or we could say that `[S MySlice]` only matches `MySlice`, but either approach runs into trouble with predeclared types. A predeclared type, like `int` is its own underlying type. We want to permit people to be able to write constraints that accept any type argument whose underlying type is `int`. In the language today, they can do that by writing `[T ~int]`. If we don’t require the tilde we would still need a way to say “any type whose underlying type is `int`”. The natural way to say that would be `[T int]`. That would mean that `[T MySlice]` and `[T int]` would behave differently, although they look very similar.

​	我们可以禁止 `[S MySlice]` ，或者我们可以说 `[S MySlice]` 只匹配 `MySlice` ，但是这两种方法都会遇到预声明类型的问题。预声明类型，例如 `int` ，其底层类型就是它自己。我们希望允许人们编写接受任何底层类型为 `int` 的类型参数的约束条件。在当前的语言中，他们可以通过写 `[T ~int]` 来实现这一点。如果我们不需要波浪符，我们仍然需要一种方式来表示“任何底层类型为 `int` 的类型”。自然的表示方式将是 `[T int]` 。这意味着 `[T MySlice]` 和 `[T int]` 将表现不同，尽管它们看起来非常相似。

We could perhaps say that `[S MySlice]` matches any type whose underlying type is the underlying type of `MySlice`, but that makes `[S MySlice]` unnecessary and confusing.

​	也许我们可以说 `[S MySlice]` 匹配任何底层类型为 `MySlice` 的类型，但这会使 `[S MySlice]` 变得不必要和混乱。

We think it’s better to require the `~` and be very clear about when we are matching the underlying type rather than the type itself.

​	我们认为要求使用 `~` 并且非常明确地指出我们何时匹配底层类型而不是类型本身更好。

## 类型推断 Type inference

Now that we’ve explained the signature of `slices.Clone`, let’s see how actually using `slices.Clone` is simplified by type inference. Remember, the signature of `Clone` is

​	现在我们已经解释了 `slices.Clone` 的签名，让我们看看类型推断如何简化实际使用 `slices.Clone` 。记住， `Clone` 的签名是：

```Go
func Clone[S ~[]E, E any](s S) S
```

A call of `slices.Clone` will pass a slice to the parameter `s`. Simple type inference will let the compiler infer that the type argument for the type parameter `S` is the type of the slice being passed to `Clone`. Type inference is then powerful enough to see that the type argument for `E` is the element type of the type argument passed to `S`.

​	调用 `slices.Clone` 将把一个切片传递给参数 `s` 。简单的类型推断将使编译器推断类型参数 `S` 的类型参数是传递给 `Clone` 的切片的类型。类型推断足够强大，可以看到类型参数 `E` 的类型参数是传递给 `S` 的类型参数的元素类型。

This means that we can write

​	这意味着我们可以这样写：

```Go
    c := Clone(ms)
```

without having to write

而不必写成：

```Go
    c := Clone[MySlice, string](ms)
```

If we refer to `Clone` without calling it, we do have to specify a type argument for `S`, as the compiler has nothing it can use to infer it. Fortunately, in that case, type inference is able to infer the type argument for `E` from the argument for `S`, and we don’t have to specify it separately.

​	如果我们引用 `Clone` 而不调用它，我们必须为 `S` 指定一个类型参数，因为编译器没有可以用来推断的东西。幸运的是，在这种情况下，类型推断能够从 `S` 的参数中推断出类型参数 `E` 的类型参数，我们不必单独指定它。

That is, we can write

​	也就是说，我们可以这样写：

```Go
    myClone := Clone[MySlice]
```

without having to write

而不必写成：

```Go
    myClone := Clone[MySlice, string]
```

## 解构类型参数 Deconstructing type parameters

The general technique we’ve used here, in which we define one type parameter `S` using another type parameter `E`, is a way to deconstruct types in generic function signatures. By deconstructing a type, we can name, and constrain, all aspects of the type.

​	我们在这里使用的一般技术，即使用另一个类型参数 `E` 来定义一个类型参数 `S` ，是一种在通用函数签名中解构类型的方法。通过解构类型，我们可以命名并约束类型的所有方面。

For example, here is the signature for `maps.Clone`.

​	例如，这是 `maps.Clone` 的签名：

```Go
func Clone[M ~map[K]V, K comparable, V any](m M) M
```

Just as with `slices.Clone`, we use a type parameter for the type of the parameter `m`, and then deconstruct the type using two other type parameters `K` and `V`.

​	与 `slices.Clone` 一样，我们使用一个类型参数来表示参数 `m` 的类型，然后使用另外两个类型参数 `K` 和 `V` 来解构类型。

In `maps.Clone` we constrain `K` to be comparable, as is required for a map key type. We can constrain the component types any way we like.

​	在 `maps.Clone` 中，我们将 `K` 约束为可比较的，这是映射键类型所需的。我们可以以任何我们喜欢的方式约束组成类型。

```Go
func WithStrings[S ~[]E, E interface { String() string }](s S) (S, []string)
```

This says that the argument of `WithStrings` must be a slice type for which the element type has a `String` method.

​	这表示 `WithStrings` 的参数必须是一个切片类型，其元素类型具有 `String` 方法。

Since all Go types can be built up from component types, we can always use type parameters to deconstruct those types and constrain them as we like.

​	由于所有的Go类型都可以从组成类型构建，我们始终可以使用类型参数来解构这些类型并根据需要对其进行约束。