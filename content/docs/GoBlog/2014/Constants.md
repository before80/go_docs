+++
title = "常量"
weight = 6
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Constants - 常量

https://go.dev/blog/constants

Rob Pike
25 August 2014

## Introduction 简介

Go is a statically typed language that does not permit operations that mix numeric types. You can’t add a `float64` to an `int`, or even an `int32` to an `int`. Yet it is legal to write `1e6*time.Second` or `math.Exp(1)` or even `1<<(' '+2.0)`. In Go, constants, unlike variables, behave pretty much like regular numbers. This post explains why that is and what it means.

Go是一种静态类型的语言，不允许混合数字类型的操作。你不能把float64加到int中，甚至不能把int32加到int中。但是写1e6*time.Second或math.Exp(1)甚至1<<(' '+2.0)都是合法的。在Go中，常量与变量不同，它的行为与普通数字差不多。这篇文章解释了为什么会这样，以及它意味着什么。

## Background: C 背景：C

In the early days of thinking about Go, we talked about a number of problems caused by the way C and its descendants let you mix and match numeric types. Many mysterious bugs, crashes, and portability problems are caused by expressions that combine integers of different sizes and “signedness”. Although to a seasoned C programmer the result of a calculation like

在思考Go的早期，我们谈到了一些由C及其后代让你混合和匹配数字类型的方式引起的问题。许多神秘的错误、崩溃和可移植性问题都是由组合不同大小和 "有符号 "的整数的表达式引起的。尽管对一个经验丰富的C程序员来说，像这样的计算结果是

```c linenums="1"
unsigned int u = 1e9;
long signed int i = -1;
... i + u ...
```

may be familiar, it isn’t *a priori* obvious. How big is the result? What is its value? Is it signed or unsigned?

可能很熟悉，但这并不是先验的。结果有多大？它的值是什么？它是有符号还是无符号的？

Nasty bugs lurk here.

这里潜伏着可怕的错误。

C has a set of rules called “the usual arithmetic conversions” and it is an indicator of their subtlety that they have changed over the years (introducing yet more bugs, retroactively).

C语言有一套被称为 "通常的算术转换 "的规则，这些规则多年来一直在变化，这说明了它们的微妙之处（引入了更多的错误，追溯性的）。

When designing Go, we decided to avoid this minefield by mandating that there is *no* mixing of numeric types. If you want to add `i` and `u`, you must be explicit about what you want the result to be. Given

在设计Go的时候，我们决定通过规定数字类型不能混合的方式来避免这个雷区。如果你想把i和u加起来，你必须明确说明你想要的结果是什么。鉴于

```go linenums="1"
var u uint
var i int
```

you can write either `uint(i)+u` or `i+int(u)`, with both the meaning and type of the addition clearly expressed, but unlike in C you cannot write `i+u`. You can’t even mix `int` and `int32`, even when `int` is a 32-bit type.

你可以写uint(i)+u或i+int(u)，明确表达加法的含义和类型，但与C语言不同，你不能写i+u。你甚至不能混合使用int和int32，即使int是一个32位类型。

This strictness eliminates a common cause of bugs and other failures. It is a vital property of Go. But it has a cost: it sometimes requires programmers to decorate their code with clumsy numeric conversions to express their meaning clearly.

这种严格性消除了一个常见的错误和其他故障的原因。这是Go的一个重要属性。但是它有一个代价：它有时需要程序员用笨拙的数字转换来装饰他们的代码，以明确表达他们的意思。

And what about constants? Given the declarations above, what would make it legal to write `i` `=` `0` or `u` `=` `0`? What is the *type* of `0`? It would be unreasonable to require constants to have type conversions in simple contexts such as `i` `=` `int(0)`.

那么常量呢？考虑到上面的声明，怎样才能使写i = 0或u = 0合法？0的类型是什么？如果要求常量在简单的上下文中进行类型转换是不合理的，比如i = int(0)。

We soon realized the answer lay in making numeric constants work differently from how they behave in other C-like languages. After much thinking and experimentation, we came up with a design that we believe feels right almost always, freeing the programmer from converting constants all the time yet being able to write things like `math.Sqrt(2)` without being chided by the compiler.

我们很快意识到，答案在于使数字常量的工作方式与它们在其他类C语言中的行为方式不同。经过大量的思考和实验，我们想出了一种设计，我们认为这种设计几乎总是正确的，使程序员不必一直转换常量，但又能写出像math.Sqrt(2)这样的东西而不被编译器责骂。

In short, constants in Go just work, most of the time anyway. Let’s see how that happens.

简而言之，Go中的常量就是有效的，反正大多数时候都是如此。让我们来看看这是如何发生的。

## Terminology 术语

First, a quick definition. In Go, `const` is a keyword introducing a name for a scalar value such as `2` or `3.14159` or `"scrumptious"`. Such values, named or otherwise, are called *constants* in Go. Constants can also be created by expressions built from constants, such as `2+3` or `2+3i` or `math.Pi/2` or `("go"+"pher")`.

首先，快速定义一下。在Go中，const是一个关键字，它为标量值引入了一个名称，如2或3.14159或 "scrumptious"。这样的值，不管是命名的还是其他的，在Go中都称为常量。常量也可以由常量构建的表达式创建，如2+3或2+3i或math.Pi/2或（"go "+"pher"）。

Some languages don’t have constants, and others have a more general definition of constant or application of the word `const`. In C and C++, for instance, `const` is a type qualifier that can codify more intricate properties of more intricate values.

有些语言没有常量，有些语言对常量的定义或对常量一词的应用更为宽泛。例如，在C和C++中，const是一个类型限定符，可以编码更复杂的值的更复杂的属性。

But in Go, a constant is just a simple, unchanging value, and from here on we’re talking only about Go.

但是在Go中，常量只是一个简单的、不变的值，从这里开始我们只谈Go。

## String constants 字符串常量

There are many kinds of numeric constants—integers, floats, runes, signed, unsigned, imaginary, complex—so let’s start with a simpler form of constant: strings. String constants are easy to understand and provide a smaller space in which to explore the type issues of constants in Go.

数字常量有很多种--整数、浮点数、符码、有符号、无符号、虚数、复数--所以让我们从一种更简单的常量形式开始：字符串。字符串常量很容易理解，并提供了一个较小的空间来探索Go中常量的类型问题。

A string constant encloses some text between double quotes. (Go also has raw string literals, enclosed by backquotes ````, but for the purpose of this discussion they have all the same properties.) Here is a string constant:

字符串常量在双引号之间包含了一些文本。(Go 也有原始的字符串字元，由反引号``括起来，但为了讨论的目的，它们具有所有相同的属性。) 下面是一个字符串常量：

```
"Hello, 世界"
```

(For much more detail about the representation and interpretation of strings, see [this blog post](https://blog.golang.org/strings).)

(关于字符串的表示和解释的更多细节，请看这篇博文）。

What type does this string constant have? The obvious answer is `string`, but that is *wrong*.

这个字符串常量的类型是什么？显然答案是字符串，但这是错误的。

This is an *untyped string constant*, which is to say it is a constant textual value that does not yet have a fixed type. Yes, it’s a string, but it’s not a Go value of type `string`. It remains an untyped string constant even when given a name:

这是一个未定型的字符串常量，也就是说，它是一个还没有固定类型的文本常量。是的，它是一个字符串，但它不是一个字符串类型的Go值。它仍然是一个未定型的字符串常量，即使被赋予了一个名字：

```go linenums="1"
const hello = "Hello, 世界"
```

After this declaration, `hello` is also an untyped string constant. An untyped constant is just a value, one not yet given a defined type that would force it to obey the strict rules that prevent combining differently typed values.

在这个声明之后，hello也是一个未定型的字符串常量。一个未定型的常量只是一个值，它还没有被赋予一个定义好的类型，这将迫使它遵守防止不同类型的值组合的严格规则。

It is this notion of an *untyped* constant that makes it possible for us to use constants in Go with great freedom.

正是这种未定型常量的概念，使我们能够在Go中自由地使用常量。

So what, then, is a *typed* string constant? It’s one that’s been given a type, like this:

那么，什么是类型化的字符串常量呢？它是一个被赋予了类型的常量，就像这样：

```go linenums="1"
const typedHello string = "Hello, 世界"
```

Notice that the declaration of `typedHello` has an explicit `string` type before the equals sign. This means that `typedHello` has Go type `string`, and cannot be assigned to a Go variable of a different type. That is to say, this code works:

注意typedHello的声明在等号前有一个明确的字符串类型。这意味着typedHello的Go类型是字符串，并且不能被分配给不同类型的Go变量。也就是说，这段代码是有效的：

```go linenums="1"
    var s string
    s = typedHello
    fmt.Println(s)
```

Run 运行

but this does not:

但这样就不行了：

```go linenums="1"
    type MyString string
    var m MyString
    m = typedHello // Type error
    fmt.Println(m)
```

Run 运行

The variable `m` has type `MyString` and cannot be assigned a value of a different type. It can only be assigned values of type `MyString`, like this:

变量m的类型是MyString，不能被分配一个不同类型的值。它只能被分配为MyString类型的值，像这样：

```go linenums="1"
    const myStringHello MyString = "Hello, 世界"
    m = myStringHello // OK
    fmt.Println(m)
```

Run 运行

or by forcing the issue with a conversion, like this:

或者通过转换来强制解决这个问题，比如这样：

```go linenums="1"
    m = MyString(typedHello)
    fmt.Println(m)
```

Run 运行

Returning to our *untyped* string constant, it has the helpful property that, since it has no type, assigning it to a typed variable does not cause a type error. That is, we can write

回到我们的无类型字符串常量，它有一个有用的特性，即由于它没有类型，将其分配给一个有类型的变量不会导致类型错误。也就是说，我们可以写

```go linenums="1"
m = "Hello, 世界"
```

or 或

```go linenums="1"
m = hello
```

because, unlike the typed constants `typedHello` and `myStringHello`, the untyped constants `"Hello, 世界"` and `hello` *have no type*. Assigning them to a variable of any type compatible with strings works without error.

因为，与类型化的常量typedHello和myStringHello不同，未类型化的常量 "Hello, 世界 "和hello没有类型。将它们赋值给与字符串兼容的任何类型的变量都不会出错。

These untyped string constants are strings, of course, so they can only be used where a string is allowed, but they do not have *type* `string`.

当然，这些未定型的字符串常量是字符串，所以它们只能在允许使用字符串的地方使用，但它们没有字符串类型。

## Default type 默认类型

As a Go programmer, you have certainly seen many declarations like

作为一个Go程序员，你肯定见过很多类似于以下的声明

```go linenums="1"
str := "Hello, 世界"
```

and by now you might be asking, “if the constant is untyped, how does `str` get a type in this variable declaration?” The answer is that an untyped constant has a default type, an implicit type that it transfers to a value if a type is needed where none is provided. For untyped string constants, that default type is obviously `string`, so

的声明，现在你可能会问，"如果常量是未定型的，str在这个变量声明中是如何得到一个类型的？" 答案是，一个没有类型的常量有一个默认的类型，一个隐含的类型，如果在没有提供类型的情况下需要一个类型，它就会转移到一个值。对于未定型的字符串常量，默认类型显然是字符串，所以

```go linenums="1"
str := "Hello, 世界"
```

or 或

```go linenums="1"
var str = "Hello, 世界"
```

means exactly the same as

的意思与下列情况完全相同

```go linenums="1"
var str string = "Hello, 世界"
```

One way to think about untyped constants is that they live in a kind of ideal space of values, a space less restrictive than Go’s full type system. But to do anything with them, we need to assign them to variables, and when that happens the *variable* (not the constant itself) needs a type, and the constant can tell the variable what type it should have. In this example, `str` becomes a value of type `string` because the untyped string constant gives the declaration its default type, `string`.

思考无类型常量的一种方式是，它们生活在一种理想的数值空间中，这个空间比Go的完整类型系统的限制要少。但是要对它们做任何事情，我们需要把它们分配给变量，当这种情况发生时，变量（而不是常量本身）需要一个类型，而常量可以告诉变量它应该有什么类型。在这个例子中，str变成了string类型的值，因为未定型的string常量给了声明一个默认的类型，即string。

In such a declaration, a variable is declared with a type and initial value. Sometimes when we use a constant, however, the destination of the value is not so clear. For instance consider this statement:

在这样的声明中，一个变量被声明为具有类型和初始值。然而，有时当我们使用一个常量时，值的目的地并不那么明确。例如，考虑这个语句：

```go linenums="1"
    fmt.Printf("%s", "Hello, 世界")
```

Run 运行

The signature of `fmt.Printf` is

fmt.Printf的签名是

```go linenums="1"
func Printf(format string, a ...interface{}) (n int, err error)
```

which is to say its arguments (after the format string) are interface values. What happens when `fmt.Printf` is called with an untyped constant is that an interface value is created to pass as an argument, and the concrete type stored for that argument is the default type of the constant. This process is analogous to what we saw earlier when declaring an initialized value using an untyped string constant.

这就是说，它的参数（在格式字符串之后）是接口值。当fmt.Printf被调用时，发生的情况是，一个接口值被创建为参数，并且为该参数存储的具体类型是常量的默认类型。这个过程类似于我们前面看到的使用未定型的字符串常量声明初始化值的过程。

You can see the result in this example, which uses the format `%v` to print the value and `%T` to print the type of the value being passed to `fmt.Printf`:

你可以在这个例子中看到结果，它使用格式%v来打印值，%T来打印被传递给fmt.Printf的值的类型：

```go linenums="1"
    fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")
    fmt.Printf("%T: %v\n", hello, hello)
```

Run 运行

If the constant has a type, that goes into the interface, as this example shows:

如果常量有一个类型，就会进入接口，如本例所示：

```go linenums="1"
    fmt.Printf("%T: %v\n", myStringHello, myStringHello)
```

Run 运行

(For more information about how interface values work, see the first sections of [this blog post](https://blog.golang.org/laws-of-reflection).)

(关于接口值如何工作的更多信息，请参见本博文的第一部分)。

In summary, a typed constant obeys all the rules of typed values in Go. On the other hand, an untyped constant does not carry a Go type in the same way and can be mixed and matched more freely. It does, however, have a default type that is exposed when, and only when, no other type information is available.

综上所述，一个类型化的常量遵守Go中类型化值的所有规则。另一方面，未类型化的常量不以同样的方式携带Go类型，可以更自由地混合和匹配。然而，它确实有一个默认的类型，当且仅当没有其他类型信息可用时，才会暴露出来。

## Default type determined by syntax 默认类型由语法决定

The default type of an untyped constant is determined by its syntax. For string constants, the only possible implicit type is `string`. For [numeric constants](https://go.dev/ref/spec#Numeric_types), the implicit type has more variety. Integer constants default to `int`, floating-point constants `float64`, rune constants to `rune` (an alias for `int32`), and imaginary constants to `complex128`. Here’s our canonical print statement used repeatedly to show the default types in action:

非类型常量的默认类型由其语法决定。对于字符串常量，唯一可能的隐式类型是字符串。对于数字常量来说，隐含类型有更多种类。整数常量默认为int，浮点常量为float64，符码常量为符码（int32的别名），虚数常量为complex128。下面是我们重复使用的典型打印语句，以显示默认类型的作用：

```go linenums="1"
    fmt.Printf("%T %v\n", 0, 0)
    fmt.Printf("%T %v\n", 0.0, 0.0)
    fmt.Printf("%T %v\n", 'x', 'x')
    fmt.Printf("%T %v\n", 0i, 0i)
```

Run 运行

(Exercise: Explain the result for `'x'`.)

(练习：解释'x'的结果。)

## Booleans 布尔类型

Everything we said about untyped string constants can be said for untyped boolean constants. The values `true` and `false` are untyped boolean constants that can be assigned to any boolean variable, but once given a type, boolean variables cannot be mixed:

我们所说的关于非类型化字符串常量的一切都可以用于非类型化布尔常量。值true和false是未定型的布尔常量，可以分配给任何布尔变量，但是一旦给定了一个类型，布尔变量就不能混合。

```go linenums="1"
    type MyBool bool
    const True = true
    const TypedTrue bool = true
    var mb MyBool
    mb = true      // OK
    mb = True      // OK
    mb = TypedTrue // Bad
    fmt.Println(mb)
```

Run 运行

Run the example and see what happens, then comment out the “Bad” line and run it again. The pattern here follows exactly that of string constants.

运行这个例子，看看会发生什么，然后注释掉 "Bad "行，再运行一次。这里的模式完全遵循字符串常数的模式。

## Floats 浮点型

Floating-point constants are just like boolean constants in most respects. Our standard example works as expected in translation:

浮点常量在大多数方面与布尔常量一样。我们的标准例子在翻译时也是如此：

```go linenums="1"
    type MyFloat64 float64
    const Zero = 0.0
    const TypedZero float64 = 0.0
    var mf MyFloat64
    mf = 0.0       // OK
    mf = Zero      // OK
    mf = TypedZero // Bad
    fmt.Println(mf)
```

Run 运行

One wrinkle is that there are *two* floating-point types in Go: `float32` and `float64`. The default type for a floating-point constant is `float64`, although an untyped floating-point constant can be assigned to a `float32` value just fine:

一个问题是，Go中有两种浮点类型：float32和float64。浮点常量的默认类型是float64，尽管一个没有类型的浮点常量可以很好地分配给float32值：

```go linenums="1"
    var f32 float32
    f32 = 0.0
    f32 = Zero      // OK: Zero is untyped
    f32 = TypedZero // Bad: TypedZero is float64 not float32.
    fmt.Println(f32)
```

Run 运行

Floating-point values are a good place to introduce the concept of overflow, or the range of values.

浮点数值是介绍溢出概念的好地方，也就是数值范围。

Numeric constants live in an arbitrary-precision numeric space; they are just regular numbers. But when they are assigned to a variable the value must be able to fit in the destination. We can declare a constant with a very large value:

数字常数生活在一个任意精度的数字空间中，它们只是普通的数字。但是当它们被分配到一个变量时，其值必须能够适应目的地。我们可以声明一个具有非常大数值的常量：

```go linenums="1"
    const Huge = 1e1000
```

—that’s just a number, after all—but we can’t assign it or even print it. This statement won’t even compile:

-这毕竟只是一个数字，但我们不能分配它，甚至不能打印它。这条语句甚至不会被编译。

```go linenums="1"
    fmt.Println(Huge)
```

Run 运行

The error is, “constant 1.00000e+1000 overflows float64”, which is true. But `Huge` might be useful: we can use it in expressions with other constants and use the value of those expressions if the result can be represented in the range of a `float64`. The statement,

错误是："常数1.00000e+1000溢出了float64"，这是真的。但是Huge可能是有用的：我们可以在有其他常数的表达式中使用它，如果结果可以在float64的范围内表示，则使用这些表达式的值。的语句。

```go linenums="1"
    fmt.Println(Huge / 1e999)
```

Run 运行

prints `10`, as one would expect.

打印出10，正如人们所期望的那样。

In a related way, floating-point constants may have very high precision, so that arithmetic involving them is more accurate. The constants defined in the [math](https://go.dev/pkg/math) package are given with many more digits than are available in a `float64`. Here is the definition of `math.Pi`:

与此相关的是，浮点常量可以有很高的精度，这样涉及到它们的算术就更精确了。math包中定义的常数比float64中的数字多很多。下面是math.Pi的定义：

```go linenums="1"
Pi  = 3.14159265358979323846264338327950288419716939937510582097494459
```

When that value is assigned to a variable, some of the precision will be lost; the assignment will create the `float64` (or `float32`) value closest to the high-precision value. This snippet

当这个值被赋值给一个变量时，一些精度将被丢失；赋值将创建最接近高精度值的float64（或float32）值。这个片段

```go linenums="1"
    pi := math.Pi
    fmt.Println(pi)
```

Run 运行

prints `3.141592653589793`.

打印出3.141592653589793。

Having so many digits available means that calculations like `Pi/2` or other more intricate evaluations can carry more precision until the result is assigned, making calculations involving constants easier to write without losing precision. It also means that there is no occasion in which the floating-point corner cases like infinities, soft underflows, and `NaNs` arise in constant expressions. (Division by a constant zero is a compile-time error, and when everything is a number there’s no such thing as “not a number”.)

有这么多可用的数字意味着像Pi/2这样的计算或其他更复杂的计算可以携带更多的精度，直到结果被分配，使涉及常数的计算更容易编写而不损失精度。这也意味着在常数表达式中不会出现浮点的角落情况，比如无穷大、软下溢和NaN。(除以常数0是一个编译时的错误，而且当所有东西都是数字时，就没有 "不是数字 "这样的东西)。

## Complex numbers 复数

Complex constants behave a lot like floating-point constants. Here’s a version of our now-familiar litany translated into complex numbers:

复数常数的行为很像浮点常数。下面是我们现在熟悉的复数的版本：

```go linenums="1"
    type MyComplex128 complex128
    const I = (0.0 + 1.0i)
    const TypedI complex128 = (0.0 + 1.0i)
    var mc MyComplex128
    mc = (0.0 + 1.0i) // OK
    mc = I            // OK
    mc = TypedI       // Bad
    fmt.Println(mc)
```

Run 运行

The default type of a complex number is `complex128`, the larger-precision version composed of two `float64` values.

复数的默认类型是complex128，是由两个float64值组成的较大精度版本。

For clarity in our example, we wrote out the full expression `(0.0+1.0i)`, but this value can be shortened to `0.0+1.0i`, `1.0i` or even `1i`.

在我们的例子中，为了清晰起见，我们写出了完整的表达式（0.0+1.0i），但是这个值可以缩短为0.0+1.0i、1.0i甚至1i。

Let’s play a trick. We know that in Go, a numeric constant is just a number. What if that number is a complex number with no imaginary part, that is, a real? Here’s one:

我们来玩个花样。我们知道，在Go中，数字常数只是一个数字。如果这个数字是一个没有虚部的复数，也就是一个实数呢？这里有一个。

```go linenums="1"
    const Two = 2.0 + 0i
```

That’s an untyped complex constant. Even though it has no imaginary part, the *syntax* of the expression defines it to have default type `complex128`. Therefore, if we use it to declare a variable, the default type will be `complex128`. The snippet

这是一个没有类型的复数常数。尽管它没有虚部，但表达式的语法将其定义为默认的复数128类型。因此，如果我们用它来声明一个变量，默认类型将是complex128。这段话

```go linenums="1"
    s := Two
    fmt.Printf("%T: %v\n", s, s)
```

Run 运行

prints `complex128:` `(2+0i)`. But numerically, `Two` can be stored in a scalar floating-point number, a `float64` or `float32`, with no loss of information. Thus we can assign `Two` to a `float64`, either in an initialization or an assignment, without problems:

打印复数128：（2+0i）。但从数值上看，Two可以存储在一个标量浮点数中，即float64或float32，而不会有任何信息损失。因此，我们可以在初始化或赋值时将二赋给浮点64，而不会出现问题：

```go linenums="1"
    var f float64
    var g float64 = Two
    f = Two
    fmt.Println(f, "and", g)
```

Run 运行

The output is `2` `and` `2`. Even though `Two` is a complex constant, it can be assigned to scalar floating-point variables. This ability for a constant to “cross” types like this will prove useful.

输出结果是2和2。尽管Two是一个复合常数，但它可以被分配给标量浮点变量。这种常数 "跨越 "类型的能力将被证明是有用的。

## Integers 整数

At last we come to integers. They have more moving parts—[many sizes, signed or unsigned, and more](https://go.dev/ref/spec#Numeric_types)—but they play by the same rules. For the last time, here is our familiar example, using just `int` this time:

最后，我们来看看整数。它们有更多的活动部件--许多尺寸，有符号或无符号，等等，但它们遵循同样的规则。最后，这里是我们熟悉的例子，这次只使用int：

```go linenums="1"
    type MyInt int
    const Three = 3
    const TypedThree int = 3
    var mi MyInt
    mi = 3          // OK
    mi = Three      // OK
    mi = TypedThree // Bad
    fmt.Println(mi)
```

Run 运行

The same example could be built for any of the integer types, which are:

同样的例子可以为任何一种整数类型建立，它们是：

```
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64
uintptr
```

(plus the aliases `byte` for `uint8` and `rune` for `int32`). That’s a lot, but the pattern in the way constants work should be familiar enough by now that you can see how things will play out.

(加上uint8的别名byte和int32的别名rune）。这是一个很大的问题，但是常量工作方式的模式现在应该足够熟悉，你可以看到事情将如何发展。

As mentioned above, integers come in a couple of forms and each form has its own default type: `int` for simple constants like `123` or `0xFF` or `-14` and `rune` for quoted characters like ‘a’, ‘世’ or ‘\r’.

如上所述，整数有几种形式，每种形式都有自己的默认类型：int用于简单的常数，如123或0xFF或-14；rune用于带引号的字符，如'a'、'世'或'/r'。

No constant form has as its default type an unsigned integer type. However, the flexibility of untyped constants means we can initialize unsigned integer variables using simple constants as long as we are clear about the type. It’s analogous to how we can initialize a `float64` using a complex number with zero imaginary part. Here are several different ways to initialize a `uint`; all are equivalent, but all must mention the type explicitly for the result to be unsigned.

没有常数形式的默认类型是无符号整数类型。然而，无类型常量的灵活性意味着我们可以使用简单的常量来初始化无符号整数变量，只要我们清楚地知道其类型。这就好比我们可以用一个虚部为零的复数来初始化float64。这里有几种不同的方法来初始化一个uint；所有的方法都是等价的，但是所有的方法都必须明确地提到类型，这样结果才是无符号的。

```go linenums="1"
var u uint = 17
var u = uint(17)
u := uint(17)
```

Similarly to the range issue mentioned in the section on floating-point values, not all integer values can fit in all integer types. There are two problems that might arise: the value might be too large, or it might be a negative value being assigned to an unsigned integer type. For instance, `int8` has range -128 through 127, so constants outside of that range can never be assigned to a variable of type `int8`:

与浮点值部分提到的范围问题类似，不是所有的整数值都能适合所有的整数类型。可能会出现两个问题：数值可能太大，或者是一个负值被分配到一个无符号的整数类型。例如，int8的范围是-128到127，所以这个范围之外的常量永远不能分配给int8类型的变量：

```go linenums="1"
    var i8 int8 = 128 // Error: too large.
```

Run 运行

Similarly, `uint8`, also known as `byte`, has range 0 through 255, so a large or negative constant cannot be assigned to a `uint8`:

同样，uint8，也被称为byte，其范围是0到255，所以一个大的或负的常数不能被分配给uint8：

```go linenums="1"
    var u8 uint8 = -1 // Error: negative value.
```

Run 运行

This type-checking can catch mistakes like this one:

这种类型检查可以抓住像这样的错误：

```go linenums="1"
    type Char byte
    var c Char = '世' // Error: '世' has value 0x4e16, too large.
```

Run 运行

If the compiler complains about your use of a constant, it’s likely a real bug like this.

如果编译器抱怨你使用了一个常数，那很可能是像这样的一个真正的错误。

## An exercise: The largest unsigned int 一个练习。最大的无符号int

Here is an informative little exercise. How do we express a constant representing the largest value that fits in a `uint`? If we were talking about `uint32` rather than `uint`, we could write

这里有一个内容丰富的小练习。我们如何表达一个代表适合于uint的最大值的常数？如果我们谈论的是uint32而不是uint，我们可以这样写

```go linenums="1"
const MaxUint32 = 1<<32 - 1
```

but we want `uint`, not `uint32`. The `int` and `uint` types have equal unspecified numbers of bits, either 32 or 64. Since the number of bits available depends on the architecture, we can’t just write down a single value.

但我们要的是uint，而不是uint32。int和uint类型有相等的未指定的位数，要么是32，要么是64。由于可用的位数取决于架构，我们不能只写下一个单一的值。

Fans of [two’s-complement arithmetic](http://en.wikipedia.org/wiki/Two's_complement), which Go’s integers are defined to use, know that the representation of `-1` has all its bits set to 1, so the bit pattern of `-1` is internally the same as that of the largest unsigned integer. We therefore might think we could write

Go的整数被定义为使用的二元互补算术的爱好者知道，-1的表示方法是将其所有位设置为1，因此-1的位模式在内部与最大的无符号整数相同。因此，我们可能认为我们可以写

```go linenums="1"
    const MaxUint uint = -1 // Error: negative value
```

Run 运行

but that is illegal because -1 cannot be represented by an unsigned variable; `-1` is not in the range of unsigned values. A conversion won’t help either, for the same reason:

但这是非法的，因为-1不能用无符号变量表示；-1不在无符号值的范围内。由于同样的原因，转换也无济于事。

```
    const MaxUint uint = uint(-1) // Error: negative value
```

Run 运行

Even though at run-time a value of -1 can be converted to an unsigned integer, the rules for constant [conversions](https://go.dev/ref/spec#Conversions) forbid this kind of coercion at compile time. That is to say, this works:

尽管在运行时，-1的值可以转换为无符号的整数，但常数转换的规则禁止在编译时进行这种强制操作。也就是说，这样做是可行的。

```go linenums="1"
    var u uint
    var v = -1
    u = uint(v)
```

Run  运行

but only because `v` is a variable; if we made `v` a constant, even an untyped constant, we’d be back in forbidden territory:

但只是因为v是一个变量；如果我们把v变成一个常数，甚至是一个未定型的常数，我们就会回到禁区：

```go linenums="1"
    var u uint
    const v = -1
    u = uint(v) // Error: negative value
```

Run运行

We return to our previous approach, but instead of `-1` we try `^0`, the bitwise negation of an arbitrary number of zero bits. But that fails too, for a similar reason: In the space of numeric values, `^0` represents an infinite number of ones, so we lose information if we assign that to any fixed-size integer:

我们回到之前的方法，但是我们尝试用^0来代替-1，即任意数量的零位的比特化否定。但这也失败了，原因类似：在数值空间中，^0代表无限多的1，所以如果我们把它分配给任何固定大小的整数，就会失去信息：

```go linenums="1"
    const MaxUint uint = ^0 // Error: overflow
```

Run 运行

How then do we represent the largest unsigned integer as a constant?

那么我们如何将最大的无符号整数表示为一个常数呢？

The key is to constrain the operation to the number of bits in a `uint` and avoiding values, such as negative numbers, that are not representable in a `uint`. The simplest `uint` value is the typed constant `uint(0)`. If `uints` have 32 or 64 bits, `uint(0)` has 32 or 64 zero bits accordingly. If we invert each of those bits, we’ll get the correct number of one bits, which is the largest `uint` value.

关键是将操作限制在uint的位数上，并避免用uint无法表示的数值，如负数。最简单的uint值是类型化的常数uint(0)。如果uint有32或64位，那么uint(0)也相应地有32或64个零位。如果我们对这些位进行反转，我们会得到正确的1位数，也就是最大的uint值。

Therefore we don’t flip the bits of the untyped constant `0`, we flip the bits of the typed constant `uint(0)`. Here, then, is our constant:

因此我们不翻转非类型常数0的位，而是翻转类型常数uint(0)的位。那么，这里就是我们的常数。

```go linenums="1"
    const MaxUint = ^uint(0)
    fmt.Printf("%x\n", MaxUint)
```

Run 运行

Whatever the number of bits it takes to represent a `uint` in the current execution environment (on the [playground](https://blog.golang.org/playground), it’s 32), this constant correctly represents the largest value a variable of type `uint` can hold.

无论在当前的执行环境中用多少位来表示一个uint（在操场上是32位），这个常数都正确地表示了一个uint类型的变量所能容纳的最大值。

If you understand the analysis that got us to this result, you understand all the important points about constants in Go.

如果你理解了让我们得到这个结果的分析，你就理解了Go中关于常数的所有重要观点。

## Numbers 数字

The concept of untyped constants in Go means that all the numeric constants, whether integer, floating-point, complex, or even character values, live in a kind of unified space. It’s when we bring them to the computational world of variables, assignments, and operations that the actual types matter. But as long as we stay in the world of numeric constants, we can mix and match values as we like. All these constants have numeric value 1:

Go中的非类型常量概念意味着所有的数字常量，无论是整数、浮点、复数，甚至是字符值，都生活在一种统一的空间中。当我们把它们带到变量、赋值和操作的计算世界中时，实际的类型才是重要的。但只要我们呆在数字常量的世界里，我们就可以随意混合和匹配数值。所有这些常量都有数字值1：

```
1
1.000
1e3-99.0*10-9
'\x01'
'\u0001'
'b' - 'a'
1.0+3i-3.0i
```

Therefore, although they have different implicit default types, written as untyped constants they can be assigned to a variable of any numeric type:

因此，尽管它们有不同的隐含默认类型，但写成非类型常量，它们可以被分配给任何数字类型的变量：

```go linenums="1"
    var f float32 = 1
    var i int = 1.000
    var u uint32 = 1e3 - 99.0*10.0 - 9
    var c float64 = '\x01'
    var p uintptr = '\u0001'
    var r complex64 = 'b' - 'a'
    var b byte = 1.0 + 3i - 3.0i

    fmt.Println(f, i, u, c, p, r, b)
```

Run 运行

The output from this snippet is: `1 1 1 1 1 (1+0i) 1`.

这个片段的输出是。1 1 1 1 1 (1+0i) 1.

You can even do nutty stuff like

你甚至可以做一些疯狂的事情，比如

```go linenums="1"
    var f = 'a' * 1.5
    fmt.Println(f)
```

Run 运行

which yields 145.5, which is pointless except to prove a point.

产生145.5，除了证明一个观点外，这毫无意义。

But the real point of these rules is flexibility. That flexibility means that, despite the fact that in Go it is illegal in the same expression to mix floating-point and integer variables, or even `int` and `int32` variables, it is fine to write

但这些规则的真正意义在于灵活性。这种灵活性意味着，尽管在Go中，在同一个表达式中混合使用浮点变量和整数变量，甚至是int和int32变量都是不合法的，但写成以下情况是可以的

```go linenums="1"
sqrt2 := math.Sqrt(2)
```

or 或

```go linenums="1"
const millisecond = time.Second/1e3
```

or 或

```go linenums="1"
bigBufferWithHeader := make([]byte, 512+1e6)
```

and have the results mean what you expect.

并使结果达到你所期望的效果。

Because in Go, numeric constants work as you expect: like numbers.

因为在Go中，数字常量的工作方式和你期望的一样：像数字一样。
