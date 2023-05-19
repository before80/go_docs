+++
title = "反射的法则"
weight = 12
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# The Laws of Reflection - 反射的法则

https://go.dev/blog/laws-of-reflection

Rob Pike
6 September 2011

## Introduction 简介

Reflection in computing is the ability of a program to examine its own structure, particularly through types; it’s a form of metaprogramming. It’s also a great source of confusion.

计算机中的反思是指程序检查其自身结构的能力，特别是通过类型；这是元编程的一种形式。它也是混乱的一个重要来源。

In this article we attempt to clarify things by explaining how reflection works in Go. Each language’s reflection model is different (and many languages don’t support it at all), but this article is about Go, so for the rest of this article the word “reflection” should be taken to mean “reflection in Go”.

在这篇文章中，我们试图通过解释反射在Go中的作用来澄清一些事情。每种语言的反射模型都是不同的（许多语言根本不支持反射），但本文是关于Go的，所以在本文的其余部分，"反射 "一词应该被理解为 "Go中的反射"。

Note added January 2022: This blog post was written in 2011 and predates parametric polymorphism (a.k.a. generics) in Go. Although nothing important in the article has become incorrect as a result of that development in the language, it has been tweaked in a few places to avoid confusing someone familiar with modern Go.

2022年1月添加的注释：这篇博文写于2011年，比Go中的参数化多态性（又称泛型）还要早。虽然文章中没有什么重要的内容因为语言的发展而变得不正确，但在一些地方进行了调整，以避免让熟悉现代Go的人感到困惑。

## Types and interfaces 类型和接口

Because reflection builds on the type system, let’s start with a refresher about types in Go.

因为反射是建立在类型系统之上的，所以我们先来复习一下Go的类型。

Go is statically typed. Every variable has a static type, that is, exactly one type known and fixed at compile time: `int`, `float32`, `*MyType`, `[]byte`, and so on. If we declare

Go是静态类型的。每个变量都有一个静态类型，也就是说，在编译时有一个已知的固定类型：int, float32, *MyType, []byte，等等。如果我们声明

```go linenums="1"
type MyInt int

var i int
var j MyInt
```

then `i` has type `int` and `j` has type `MyInt`. The variables `i` and `j` have distinct static types and, although they have the same underlying type, they cannot be assigned to one another without a conversion.

那么i的类型是int，j的类型是MyInt。变量i和j有不同的静态类型，虽然它们有相同的底层类型，但如果不进行转换，它们就不能相互分配。

One important category of type is interface types, which represent fixed sets of methods. (When discussing reflection, we can ignore the use of interface definitions as constraints within polymorphic code.) An interface variable can store any concrete (non-interface) value as long as that value implements the interface’s methods. A well-known pair of examples is `io.Reader` and `io.Writer`, the types `Reader` and `Writer` from the [io package](https://go.dev/pkg/io/):

类型的一个重要类别是接口类型，它代表了固定的方法集。(在讨论反射时，我们可以忽略接口定义作为多态代码中的约束的使用）。一个接口变量可以存储任何具体（非接口）的值，只要该值实现了接口的方法。一对著名的例子是io.Reader和io.Writer，它们是io包中的Reader和Writer类型：

```go linenums="1"
// Reader is the interface that wraps the basic Read method.
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer is the interface that wraps the basic Write method.
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Any type that implements a `Read` (or `Write`) method with this signature is said to implement `io.Reader` (or `io.Writer`). For the purposes of this discussion, that means that a variable of type `io.Reader` can hold any value whose type has a `Read` method:

任何用这个签名实现读（或写）方法的类型都被称为实现了io.Reader（或io.Writer）。在本讨论中，这意味着一个io.Reader类型的变量可以持有任何类型具有Read方法的值：

```go linenums="1"
var r io.Reader
r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)
// and so on
```

It’s important to be clear that whatever concrete value `r` may hold, `r`’s type is always `io.Reader`: Go is statically typed and the static type of `r` is `io.Reader`.

重要的是要清楚，无论r的具体值是什么，r的类型始终是io.Reader。Go是静态类型的，r的静态类型是io.Reader。

An extremely important example of an interface type is the empty interface:

接口类型的一个极其重要的例子是空接口：

```go linenums="1"
interface{}
```

or its equivalent alias,

或其等效别名，

```go linenums="1"
any
```

It represents the empty set of methods and is satisfied by any value at all, since every value has zero or more methods.

它代表了方法的空集合，并且可以被任何值所满足，因为每个值都有零或更多的方法。

Some people say that Go’s interfaces are dynamically typed, but that is misleading. They are statically typed: a variable of interface type always has the same static type, and even though at run time the value stored in the interface variable may change type, that value will always satisfy the interface.

有些人说Go的接口是动态类型的，但这是一种误导。它们是静态类型的：接口类型的变量总是具有相同的静态类型，即使在运行时存储在接口变量中的值可能会改变类型，但该值将始终满足接口。

We need to be precise about all this because reflection and interfaces are closely related.

我们需要对所有这些进行精确的分析，因为反射和接口是密切相关的。

## The representation of an interface 接口的表示

Russ Cox has written a [detailed blog post](https://research.swtch.com/2009/12/go-data-structures-interfaces.html) about the representation of interface values in Go. It’s not necessary to repeat the full story here, but a simplified summary is in order.

Russ Cox写了一篇关于Go中接口值表示的详细博文。这里没有必要重复完整的故事，但简化的总结是有必要的。

A variable of interface type stores a pair: the concrete value assigned to the variable, and that value’s type descriptor. To be more precise, the value is the underlying concrete data item that implements the interface and the type describes the full type of that item. For instance, after

一个接口类型的变量存储了一对：分配给该变量的具体数值，以及该数值的类型描述符。更准确地说，值是实现接口的底层具体数据项，而类型则描述了该数据项的完整类型。例如，在

```go linenums="1"
var r io.Reader
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
if err != nil {
    return nil, err
}
r = tty
```

`r` contains, schematically, the (value, type) pair, (`tty`, `*os.File`). Notice that the type `*os.File` implements methods other than `Read`; even though the interface value provides access only to the `Read` method, the value inside carries all the type information about that value. That’s why we can do things like this:

r包含，示意性的，（值，类型）对，（tty，*os.File）。请注意，*os.File类型实现了除Read之外的其他方法；尽管接口值只提供了对Read方法的访问，但里面的值却携带了关于该值的所有类型信息。这就是为什么我们可以做这样的事情：

```go linenums="1"
var w io.Writer
w = r.(io.Writer)
```

The expression in this assignment is a type assertion; what it asserts is that the item inside `r` also implements `io.Writer`, and so we can assign it to `w`. After the assignment, `w` will contain the pair (`tty`, `*os.File`). That’s the same pair as was held in `r`. The static type of the interface determines what methods may be invoked with an interface variable, even though the concrete value inside may have a larger set of methods.

这个赋值中的表达式是一个类型断言；它断言r里面的项目也实现了io.Writer，所以我们可以把它赋给w。接口的静态类型决定了哪些方法可以被接口变量调用，即使里面的具体值可能有更大的方法集。

Continuing, we can do this:

继续下去，我们可以这样做：

```go linenums="1"
var empty interface{}
empty = w
```

and our empty interface value `empty` will again contain that same pair, (`tty`, `*os.File`). That’s handy: an empty interface can hold any value and contains all the information we could ever need about that value.

而我们的空接口值empty将再次包含同样的一对，（tty，*os.File）。这很方便：一个空接口可以容纳任何值，并包含我们可能需要的关于这个值的所有信息。

(We don’t need a type assertion here because it’s known statically that `w` satisfies the empty interface. In the example where we moved a value from a `Reader` to a `Writer`, we needed to be explicit and use a type assertion because `Writer`’s methods are not a subset of `Reader`’s.)

(我们在这里不需要类型断言，因为静态地知道w满足空接口的要求。在我们把一个值从Reader移到Writer的例子中，我们需要明确地使用一个类型断言，因为Writer的方法不是Reader的子集）。

One important detail is that the pair inside an interface variable always has the form (value, concrete type) and cannot have the form (value, interface type). Interfaces do not hold interface values.

一个重要的细节是，接口变量内部的一对总是具有（值，具体类型）的形式，而不能具有（值，接口类型）的形式。接口不持有接口值。

Now we’re ready to reflect.

现在我们准备反思了。

## The first law of reflection 反射的第一个定律

## 1. Reflection goes from interface value to reflection object. 1. 反射从接口值到反射对象。

At the basic level, reflection is just a mechanism to examine the type and value pair stored inside an interface variable. To get started, there are two types we need to know about in [package reflect](https://go.dev/pkg/reflect/): [Type](https://go.dev/pkg/reflect/#Type) and [Value](https://go.dev/pkg/reflect/#Value). Those two types give access to the contents of an interface variable, and two simple functions, called `reflect.TypeOf` and `reflect.ValueOf`, retrieve `reflect.Type` and `reflect.Value` pieces out of an interface value. (Also, from a `reflect.Value` it’s easy to get to the corresponding `reflect.Type`, but let’s keep the `Value` and `Type` concepts separate for now.)

在基本层面上，反射只是一种检查存储在接口变量内的类型和值对的机制。为了入门，在包反射中我们需要了解两种类型。类型和值。这两种类型可以访问接口变量的内容，两个简单的函数，叫做 reflect.TypeOf 和 reflect.ValueOf，可以从一个接口值中检索 reflect.Type 和 reflect.Value 片段。(另外，从一个 reflect.Value 可以很容易地得到相应的 reflect.Type，但是现在让我们把 Value 和 Type 的概念分开。）

Let’s start with `TypeOf`:

让我们从 TypeOf 开始：

```go linenums="1"
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    fmt.Println("type:", reflect.TypeOf(x))
}
```

This program prints

这个程序打印出

```go linenums="1"
type: float64
```

You might be wondering where the interface is here, since the program looks like it’s passing the `float64` variable `x`, not an interface value, to `reflect.TypeOf`. But it’s there; as [godoc reports](https://go.dev/pkg/reflect/#TypeOf), the signature of `reflect.TypeOf` includes an empty interface:

你可能想知道这里的接口在哪里，因为这个程序看起来是在向reflect.TypeOf传递float64变量x，而不是一个接口值。但它就在那里；正如 godoc 所报告的，reflect.TypeOf 的签名包括一个空接口：

```go linenums="1"
// TypeOf returns the reflection Type of the value in the interface{}.
func TypeOf(i interface{}) Type
```

When we call `reflect.TypeOf(x)`, `x` is first stored in an empty interface, which is then passed as the argument; `reflect.TypeOf` unpacks that empty interface to recover the type information.

当我们调用reflect.TypeOf(x)时，x首先被存储在一个空接口中，然后被作为参数传递；reflect.TypeOf解包这个空接口以恢复类型信息。

The `reflect.ValueOf` function, of course, recovers the value (from here on we’ll elide the boilerplate and focus just on the executable code):

当然，reflect.ValueOf函数会恢复值（从这里开始，我们将省略这些模板，只关注可执行代码）：

```go linenums="1"
var x float64 = 3.4
fmt.Println("value:", reflect.ValueOf(x).String())
```

prints打印



```
value: <float64 Value>
```

(We call the `String` method explicitly because by default the `fmt` package digs into a `reflect.Value` to show the concrete value inside. The `String` method does not.)

(我们明确地调用String方法，因为默认情况下，fmt包会挖掘reflect.Value来显示里面的具体数值。而String方法不会这样做）。

Both `reflect.Type` and `reflect.Value` have lots of methods to let us examine and manipulate them. One important example is that `Value` has a `Type` method that returns the `Type` of a `reflect.Value`. Another is that both `Type` and `Value` have a `Kind` method that returns a constant indicating what sort of item is stored: `Uint`, `Float64`, `Slice`, and so on. Also methods on `Value` with names like `Int` and `Float` let us grab values (as `int64` and `float64`) stored inside:

reflect.Type 和 reflect.Value 都有很多方法让我们检查和操作它们。一个重要的例子是，Value 有一个 Type 方法来返回 reflect.Value 的 Type。另一个例子是Type和Value都有一个Kind方法，该方法返回一个常数，表明存储的是什么类型的项目。Uint, Float64, Slice, 等等。另外，Value上的方法有Int和Float这样的名字，让我们可以抓取里面存储的值（如int64和float64）。

```go linenums="1"
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
```

prints 打印

```go linenums="1"
type: float64
kind is float64: true
value: 3.4
```

There are also methods like `SetInt` and `SetFloat` but to use them we need to understand settability, the subject of the third law of reflection, discussed below.

还有一些方法，如SetInt和SetFloat，但要使用它们，我们需要了解settingability，即下面讨论的反射第三定律的主题。

The reflection library has a couple of properties worth singling out. First, to keep the API simple, the “getter” and “setter” methods of `Value` operate on the largest type that can hold the value: `int64` for all the signed integers, for instance. That is, the `Int` method of `Value` returns an `int64` and the `SetInt` value takes an `int64`; it may be necessary to convert to the actual type involved:

反射库有几个属性值得特别指出。首先，为了保持API的简单性，Value的 "getter "和 "setter "方法操作的是可以容纳该值的最大类型：例如，所有有符号的整数的int64。也就是说，Value的Int方法返回一个int64，SetInt值取一个int64；可能需要转换为实际涉及的类型：

```go linenums="1"
var x uint8 = 'x'
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())                            // uint8.
fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
x = uint8(v.Uint())                                       // v.Uint returns a uint64.
```

The second property is that the `Kind` of a reflection object describes the underlying type, not the static type. If a reflection object contains a value of a user-defined integer type, as in

第二个属性是，反射对象的Kind描述的是底层类型，而不是静态类型。如果一个反射对象包含一个用户定义的整数类型的值，如

```go linenums="1"
type MyInt int
var x MyInt = 7
v := reflect.ValueOf(x)
```

the `Kind` of `v` is still `reflect.Int`, even though the static type of `x` is `MyInt`, not `int`. In other words, the `Kind` cannot discriminate an `int` from a `MyInt` even though the `Type` can.

v的Kind仍然是reflect.Int，尽管x的静态类型是MyInt而不是int。换句话说，尽管类型可以区分int和MyInt，但Kind不能区分int。

## The second law of reflection 反射的第二定律

## 2. Reflection goes from reflection object to interface value.  2. 反射从反射对象到接口值。

Like physical reflection, reflection in Go generates its own inverse.

像物理反射一样，Go中的反射也会产生它自己的逆向。

Given a `reflect.Value` we can recover an interface value using the `Interface` method; in effect the method packs the type and value information back into an interface representation and returns the result:

给定一个reflect.Value，我们可以使用Interface方法恢复一个接口值；实际上，该方法将类型和值信息打包回一个接口表示，并返回结果：

```go linenums="1"
// Interface returns v's value as an interface{}.
func (v Value) Interface() interface{}
```

As a consequence we can say

因此，我们可以说

```go linenums="1"
y := v.Interface().(float64) // y will have type float64.
fmt.Println(y)
```

to print the `float64` value represented by the reflection object `v`.

来打印由反射对象v代表的float64值。

We can do even better, though. The arguments to `fmt.Println`, `fmt.Printf` and so on are all passed as empty interface values, which are then unpacked by the `fmt` package internally just as we have been doing in the previous examples. Therefore all it takes to print the contents of a `reflect.Value` correctly is to pass the result of the `Interface` method to the formatted print routine:

不过，我们可以做得更好。fmt.Println、fmt.Printf 等的参数都是作为空的接口值传递的，然后由 fmt 包内部解包，就像我们在前面的例子中做的那样。因此，正确打印reflect.Value的内容只需要将接口方法的结果传递给格式化的打印例程：

```go linenums="1"
fmt.Println(v.Interface())
```

(Since this article was first written, a change was made to the `fmt` package so that it automatically unpacks a `reflect.Value` like this, so we could just say

(自从这篇文章第一次写完后，fmt包做了一个修改，这样它就会像这样自动解包一个reflect.Value，所以我们可以直接说

```go linenums="1"
fmt.Println(v)
```

for the same result, but for clarity we’ll keep the `.Interface()` calls here.)

得到同样的结果，但为了清楚起见，我们将在这里保留.Interface()的调用）。

Since our value is a `float64`, we can even use a floating-point format if we want:

由于我们的值是 float64，如果我们想的话，甚至可以使用浮点格式：

```go linenums="1"
fmt.Printf("value is %7.1e\n", v.Interface())
```

and get in this case

并在本例中得到

```
3.4e+00
```

Again, there’s no need to type-assert the result of `v.Interface()` to `float64`; the empty interface value has the concrete value’s type information inside and `Printf` will recover it.

同样，不需要对v.Interface()的结果进行type-assert到float64；空的接口值里面有具体值的type信息，Printf会恢复它。

In short, the `Interface` method is the inverse of the `ValueOf` function, except that its result is always of static type `interface{}`.

简而言之，Interface方法是ValueOf函数的倒数，只是它的结果总是静态类型interface{}。

Reiterating: Reflection goes from interface values to reflection objects and back again.

重申一下。反射从接口值到反射对象，然后再返回。

## The third law of reflection 反射的第三定律

## 3. To modify a reflection object, the value must be settable. 3. 要修改一个反射对象，其值必须是可设置的。

The third law is the most subtle and confusing, but it’s easy enough to understand if we start from first principles.

第三条定律是最微妙和令人困惑的，但如果我们从第一条原则出发，就很容易理解。

Here is some code that does not work, but is worth studying.

下面是一些不起作用的代码，但值得研究。

```go linenums="1"
var x float64 = 3.4
v := reflect.ValueOf(x)
v.SetFloat(7.1) // Error: will panic.
```

If you run this code, it will panic with the cryptic message

如果你运行这段代码，它就会出现恐慌，并发出神秘的信息

```
panic: reflect.Value.SetFloat using unaddressable value
```

The problem is not that the value `7.1` is not addressable; it’s that `v` is not settable. Settability is a property of a reflection `Value`, and not all reflection `Values` have it.

问题不在于7.1这个值不可寻址，而在于v不可设置。可设置性是反射值的一个属性，并不是所有的反射值都有这个属性。

The `CanSet` method of `Value` reports the settability of a `Value`; in our case,

Value的CanSet方法报告一个Value的可设置性；在我们的例子中。

```go linenums="1"
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("settability of v:", v.CanSet())
```

prints

打印

```
settability of v: false
```

It is an error to call a `Set` method on a non-settable `Value`. But what is settability?

在一个不可设置的值上调用一个设置方法是一个错误。但什么是可设置性？

Settability is a bit like addressability, but stricter. It’s the property that a reflection object can modify the actual storage that was used to create the reflection object. Settability is determined by whether the reflection object holds the original item. When we say

可设置性有点像可寻址性，但更严格。它是一个反射对象可以修改用于创建反射对象的实际存储的属性。可设置性是由反射对象是否持有原始项目决定的。当我们说

```go linenums="1"
var x float64 = 3.4
v := reflect.ValueOf(x)
```

we pass a copy of `x` to `reflect.ValueOf`, so the interface value created as the argument to `reflect.ValueOf` is a copy of `x`, not `x` itself. Thus, if the statement

我们把 x 的副本传递给 reflect.ValueOf，所以作为 reflect.ValueOf 的参数创建的接口值是 x 的副本，而不是 x 本身。因此，如果语句

```go linenums="1"
v.SetFloat(7.1)
```

were allowed to succeed, it would not update `x`, even though `v` looks like it was created from `x`. Instead, it would update the copy of `x` stored inside the reflection value and `x` itself would be unaffected. That would be confusing and useless, so it is illegal, and settability is the property used to avoid this issue.

允许成功的话，它不会更新x，尽管v看起来是由x创建的。相反，它将更新存储在反射值中的x的副本，而x本身不会受到影响。这将是混乱和无用的，所以它是非法的，而settability是用来避免这个问题的属性。

If this seems bizarre, it’s not. It’s actually a familiar situation in unusual garb. Think of passing `x` to a function:

如果这看起来很怪异，其实不然。它实际上是一个熟悉的情况，穿着不寻常的衣服。想想把x传给一个函数：

```go linenums="1"
f(x)
```

We would not expect `f` to be able to modify `x` because we passed a copy of `x`’s value, not `x` itself. If we want `f` to modify `x` directly we must pass our function the address of `x` (that is, a pointer to `x`):

我们不会期望f能够修改x，因为我们传递的是x的值的拷贝，而不是x本身。如果我们想让f直接修改x，我们必须把x的地址（也就是一个指向x的指针）传给我们的函数：

```go linenums="1"
f(&x)
```

This is straightforward and familiar, and reflection works the same way. If we want to modify `x` by reflection, we must give the reflection library a pointer to the value we want to modify.

这是直接的和熟悉的，反射的工作方式也是如此。如果我们想通过反射来修改x，我们必须给反射库一个指向我们想修改的值的指针。

```go linenums="1"
var x float64 = 3.4
p := reflect.ValueOf(&x) // Note: take the address of x.
fmt.Println("type of p:", p.Type())
fmt.Println("settability of p:", p.CanSet())
```

Let’s do that. First we initialize `x` as usual and then create a reflection value that points to it, called `p`.

让我们来做这件事。首先我们像往常一样初始化x，然后创建一个指向它的反射值，叫做p。

```go linenums="1"
var x float64 = 3.4
p := reflect.ValueOf(&x) // Note: take the address of x.
fmt.Println("type of p:", p.Type())
fmt.Println("settability of p:", p.CanSet())
```

The output so far is

到目前为止的输出是

```go linenums="1"
type of p: *float64
settability of p: false
```

The reflection object `p` isn’t settable, but it’s not `p` we want to set, it’s (in effect) `*p`. To get to what `p` points to, we call the `Elem` method of `Value`, which indirects through the pointer, and save the result in a reflection `Value` called `v`:

反射对象p是不可设置的，但我们想设置的不是p，而是（实际上）*p。为了得到p所指向的东西，我们调用Value的Elem方法，它通过指针进行间接操作，并将结果保存在一个叫做v的反射Value中：

```go linenums="1"
v := p.Elem()
fmt.Println("settability of v:", v.CanSet())
```

Now `v` is a settable reflection object, as the output demonstrates,

现在v是一个可设置的反射对象，正如输出所演示的那样。

```
settability of v: true
```

and since it represents `x`, we are finally able to use `v.SetFloat` to modify the value of `x`:

并且由于它代表x，我们最终能够使用v.SetFloat来修改x的值：

```go linenums="1"
v.SetFloat(7.1)
fmt.Println(v.Interface())
fmt.Println(x)
```

The output, as expected, is

正如预期的那样，输出结果是

```
7.1
7.1
```

Reflection can be hard to understand but it’s doing exactly what the language does, albeit through reflection `Types` and `Values` that can disguise what’s going on. Just keep in mind that reflection Values need the address of something in order to modify what they represent.

反射可能很难理解，但它所做的正是语言所做的，尽管通过反射类型和值可以掩盖正在发生的事情。请记住，反射值需要某个东西的地址，以便修改它们所代表的东西。

## Structs 结构体

In our previous example `v` wasn’t a pointer itself, it was just derived from one. A common way for this situation to arise is when using reflection to modify the fields of a structure. As long as we have the address of the structure, we can modify its fields.

在我们前面的例子中，v本身并不是一个指针，它只是从一个指针派生出来的。出现这种情况的一个常见方法是使用反射来修改结构的字段。只要我们有结构的地址，我们就可以修改它的字段。

Here’s a simple example that analyzes a struct value, `t`. We create the reflection object with the address of the struct because we’ll want to modify it later. Then we set `typeOfT` to its type and iterate over the fields using straightforward method calls (see [package reflect](https://go.dev/pkg/reflect/) for details). Note that we extract the names of the fields from the struct type, but the fields themselves are regular `reflect.Value` objects.

下面是一个分析结构值t的简单例子。我们用结构的地址创建反射对象，因为我们以后会想修改它。然后我们将typeOfT设置为它的类型，并使用直接的方法调用遍历字段（详见包reflect）。请注意，我们从结构类型中提取字段的名称，但字段本身是普通的 reflect.Value 对象。

```go linenums="1"
type T struct {
    A int
    B string
}
t := T{23, "skidoo"}
s := reflect.ValueOf(&t).Elem()
typeOfT := s.Type()
for i := 0; i < s.NumField(); i++ {
    f := s.Field(i)
    fmt.Printf("%d: %s %s = %v\n", i,
        typeOfT.Field(i).Name, f.Type(), f.Interface())
}
```

The output of this program is

这个程序的输出是

```
0: A int = 23
1: B string = skidoo
```

There’s one more point about settability introduced in passing here: the field names of `T` are upper case (exported) because only exported fields of a struct are settable.

这里还顺便介绍了一个关于可设置性的观点：T的字段名是大写的（导出的），因为只有结构的导出字段是可设置的。

Because `s` contains a settable reflection object, we can modify the fields of the structure.

因为s包含一个可设置的反射对象，我们可以修改结构的字段。

```
s.Field(0).SetInt(77)
s.Field(1).SetString("Sunset Strip")
fmt.Println("t is now", t)
```

And here’s the result:

这里是结果：

```
t is now {77 Sunset Strip}
```

If we modified the program so that `s` was created from `t`, not `&t`, the calls to `SetInt` and `SetString` would fail as the fields of `t` would not be settable.

如果我们修改程序，使s是从t而不是&t创建的，那么对SetInt和SetString的调用就会失败，因为t的字段不能被设置。

## Conclusion 结论

Here again are the laws of reflection:

这里又是反射的规律：

- Reflection goes from interface value to reflection object. 反射从接口值到反射对象。
- Reflection goes from reflection object to interface value. 反射从反射对象到接口值。
- To modify a reflection object, the value must be settable. 要修改一个反射对象，其值必须是可设置的。

Once you understand these laws reflection in Go becomes much easier to use, although it remains subtle. It’s a powerful tool that should be used with care and avoided unless strictly necessary.

一旦你理解了这些定律，Go中的反射就会变得更容易使用，尽管它仍然很微妙。它是一个强大的工具，应该谨慎使用，除非绝对必要，否则应避免使用。

There’s plenty more to reflection that we haven’t covered — sending and receiving on channels, allocating memory, using slices and maps, calling methods and functions — but this post is long enough. We’ll cover some of those topics in a later article.

还有很多关于反射的内容我们没有涉及--在通道上发送和接收，分配内存，使用分片和映射，调用方法和函数--但是这篇文章已经够长了。我们将在以后的文章中介绍其中的一些主题。
