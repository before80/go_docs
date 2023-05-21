+++
title = "All your comparable types"
weight = 97
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# All your comparable types - 所有可比较的类型

[https://go.dev/blog/comparable](https://go.dev/blog/comparable)

Robert Griesemer
17 February 2023

On February 1 we released our latest Go version, 1.20, which included a few language changes. Here we’ll discuss one of those changes: the predeclared `comparable` type constraint is now satisfied by all [comparable types](https://go.dev/ref/spec#Comparison_operators). Surprisingly, before Go 1.20, some comparable types did not satisfy `comparable`!

​	在2023年2月1日，我们发布了最新的Go版本1.20，其中包含了一些语言变更。在这里，我们将讨论其中的一个变更：预定义的可比较类型约束现在适用于所有可比较的类型。令人惊讶的是，在Go 1.20之前，一些可比较的类型并不满足可比较的要求！

If you’re confused, you’ve come to the right place. Consider the valid map declaration

​	如果你感到困惑，你来对地方了。考虑以下有效的映射声明

```Go
var lookupTable map[any]string
```

where the map’s key type is `any` (which is a [comparable type](https://go.dev/ref/spec#Comparison_operators)). This works perfectly fine in Go. On the other hand, before Go 1.20, the seemingly equivalent generic map type

其中该映射的键类型为`any`（[可比较类型]({{< ref "/docs/References/LanguageSpecification/Expressions#comparison-operators-比较运算符" >}})）。这在Go中运行得很好。另一方面，在Go 1.20之前，看似等价的泛型映射类型

```Go
type genericLookupTable[K comparable, V any] map[K]V
```

could be used just like a regular map type, but produced a compile-time error when `any` was used as the key type:

可以像普通映射类型一样使用，但在将`any`作为键类型时会产生编译时错误：

```Go
var lookupTable genericLookupTable[any, string] // ERROR: any does not implement comparable (Go 1.18 and Go 1.19)
```

Starting with Go 1.20 this code will compile just fine.

从Go 1.20开始，这段代码将编译通过。

The pre-Go 1.20 behavior of `comparable` was particularly annoying because it prevented us from writing the kind of generic libraries we were hoping to write with generics in the first place. The proposed [`maps.Clone`](https://go.dev/issue/57436) function

`comparable`在Go 1.20之前的行为特别令人烦恼，因为它阻止我们编写我们最初希望使用泛型编写的通用库。提议中的[maps.Clone](https://go.dev/issue/57436)函数

```Go
func Clone[M ~map[K]V, K comparable, V any](m M) M { … }
```

can be written but could not be used for a map such as `lookupTable` for the same reason our `genericLookupTable` could not be used with `any` as key type.

可以编写，但不能用于像`lookupTable`这样的映射，原因与我们的`genericLookupTable`不能用`any`作为键类型一样。

In this blog post, we hope to shine some light on the language mechanics behind all this. In order to do so, we start with a bit of background information.

​	在本博客文章中，我们希望为这一切背后的语言机制提供一些解释。为了做到这一点，我们从一些背景信息开始。

## Type parameters and constraints 类型参数和约束 

Go 1.18 introduced generics and, with that, [*type parameters*](https://go.dev/ref/spec#Type_parameter_declarations) as a new language construct.

​	Go 1.18引入了泛型和[类型参数]({{< ref "/docs/GoTour/Generics#type-parameters-类型参数" >}})作为一种新的语言构造。

In an ordinary function, a parameter ranges over a set of values that is restricted by its type. Analogously, in a generic function (or type), a type parameter ranges over a set of types that is restricted by its [*type constraint*](https://go.dev/ref/spec#Type_constraints). Thus, a type constraint defines the *set of types* that are permissible as type arguments.

​	在普通函数中，参数范围限定于其类型限制的一组值。类似地，在泛型函数（或类型）中，类型参数范围限定于其**类型约束**的一组类型。因此，类型约束定义了可作为类型参数的**类型集合**。

Go 1.18 also changed how we view interfaces: while in the past an interface defined a set of methods, now an interface defines a set of types. This new view is completely backward compatible: for any given set of methods defined by an interface, we can imagine the (infinite) set of all types that implement those methods. For instance, given an [`io.Writer`](https://go.dev/pkg/io#Writer) interface, we can imagine the infinite set of all types that have a `Write` method with the appropriate signature. All of these types *implement* the interface because they all have the required `Write` method.

​	Go 1.18还改变了我们对接口的理解方式：过去，接口定义了一组方法，现在接口定义了一组类型。这种新的视角完全向后兼容：对于接口定义的一组方法，我们可以想象实现这些方法的所有类型的（无限）集合。例如，给定一个`io.Writer`接口，我们可以想象具有适当签名的`Write`方法的所有类型的无限集合。所有这些类型都**实现了**该接口，因为它们都具有所需的`Write`方法。

But the new type set view is more powerful than the old method set one: we can describe a set of types explicitly, not only indirectly through methods. This gives us new ways to control a type set. Starting with Go 1.18, an interface may embed not just other interfaces, but any type, a union of types, or an infinite set of types that share the same [underlying type](https://go.dev/ref/spec#Underlying_types). These types are then included in the [type set computation](https://go.dev/ref/spec#General_interfaces): the union notation `A|B` means “type `A` or type `B`”, and the `~T` notation stands for “all types that have the underlying type `T`”. For instance, the interface

​	但是，新的类型集合视角比旧的方法集合视角更强大：我们可以明确地描述一组类型，而不仅仅通过方法间接描述。这为我们提供了控制类型集合的新方法。从Go 1.18开始，接口可以嵌入不仅是其他接口，还包括任何类型、类型的并集，或者具有相同[底层类型]({{< ref "/docs/References/LanguageSpecification/PropertiesOfTypesAndValues#underlying-types-底层类型基本类型">}})的无限类型集合。这些类型将包含在[类型集合的计算]({{< ref "/docs/References/LanguageSpecification/Types#general-interfaces-通用接口">}})中：并集符号`A|B`表示“类型`A`或类型`B`”，`~T`符号表示“所有具有底层类型`T`的类型”。例如，接口

```Go
interface {
    ~int | ~string
    io.Writer
}
```

defines the set of all types whose underlying types are either `int` or `string` and that also implement `io.Writer`’s `Write` method.

定义了所有底层类型为`int`或`string`且实现了`io.Writer`的`Write`方法的类型集合。

Such generalized interfaces can’t be used as variable types. But because they describe type sets they are used as type constraints, which are sets of types. For instance, we can write a generic `min` function

​	这种广义接口不能用作变量类型。但是，因为它们描述了类型集合，它们被用作类型约束，即类型的集合。例如，我们可以编写一个泛型的`min`函数

```Go
func min[P interface{ ~int64 | ~float64 }](x, y P) P
```

which accepts any `int64` or `float64` argument. (Of course, a more realistic implementation would use a constraint that enumerates all basic types with an `<` operator.)

​	它接受任何`int64`或`float64`类型的实参。（当然，更实际的实现会使用一个枚举了所有具有`<`运算符的基本类型的约束。）

As an aside, because enumerating explicit types without methods is common, a little bit of [syntactic sugar](https://en.wikipedia.org/wiki/Syntactic_sugar) allows us to [omit the enclosing `interface{}`](https://go.dev/ref/spec#General_interfaces), leading to the compact and more idiomatic

​	顺便提一下，由于枚举没有方法的显式类型是常见的，一些[语法糖](https://en.wikipedia.org/wiki/Syntactic_sugar)允许我们[省略包围的`interface{}`]({{< ref "/docs/References/LanguageSpecification/Types#general-interfaces-通用接口">}})，从而导致更紧凑和更符合惯例的写法

```Go
func min[P ~int64 | ~float64](x, y P) P { … }
```

With the new type set view we also need a new way to explain what it means to [*implement*](https://go.dev/ref/spec#Implementing_an_interface) an interface. We say that a (non-interface) type `T` implements an interface `I` if `T` is an element of the interface’s type set. If `T` is an interface itself, it describes a type set. Every single type in that set must also be in the type set of `I`, otherwise `T` would contain types that do not implement `I`. Thus, if `T` is an interface, it implements interface `I` if the type set of `T` is a subset of the type set of `I`.

使用新的类型集合视角，我们还需要一种新的方式来解释实现接口的含义。我们说，如果一个（非接口）类型`T`是接口`I`的类型集合的元素，那么`T`就实现了接口`I`。如果`T`本身是一个接口，它描述了一个类型集合。该集合中的每个类型都必须也属于`I`的类型集合，否则`T`将包含未实现`I`的类型。因此，如果`T`是一个接口，当`T`的类型集合是`I`的类型集合的子集时，它就实现了接口`I`。

Now we have all the ingredients in place to understand constraint satisfaction. As we have seen earlier, a type constraint describes the set of acceptable argument types for a type parameter. A type argument satisfies the corresponding type parameter constraint if the type argument is in the set described by the constraint interface. This is another way of saying that the type argument implements the constraint. In Go 1.18 and Go 1.19, constraint satisfaction meant constraint implementation. As we’ll see in a bit, in Go 1.20 constraint satisfaction is not quite constraint implementation anymore.

​	现在我们已经掌握了所有必要的要素来理解约束满足。正如我们之前所见，类型约束描述了类型参数的可接受实参类型集合。如果某一类型实参满足对应的类型约束，那么该类型实参就在该约束接口所描述的集合中。这又可以说成该类型参数实现了该约束。在Go 1.18和Go 1.19中，约束满足意味着约束的实现。正如我们稍后将看到的，*从Go 1.20开始，约束满足不再完全等同于约束的实现。*

## Operations on type parameter values 类型参数值的操作 

A type constraint does not just specify what type arguments are acceptable for a type parameter, it also determines the operations that are possible on values of a type parameter. As we would expect, if a constraint defines a method such as `Write`, the `Write` method can be called on a value of the respective type parameter. More generally, an operation such as `+` or `*` that is supported by all types in the type set defined by a constraint is permitted with values of the corresponding type parameter.

​	类型约束不仅指定了哪些类型实参对于类型参数是可接受的，还确定了可以对类型参数值进行的操作。如我们所预期的，如果约束定义了一个名为`Write`的方法，那么可以在相应类型参数的值上调用`Write`方法。更一般地说，只要约束定义的类型集合中的所有类型都支持操作（如`+`或`*`），那么可以在相应类型参数的值上执行该操作。

For instance, given the `min` example, in the function body any operation that is supported by `int64` and `float64` types is permitted on values of the type parameter `P`. That includes all the basic arithmetic operations, but also comparisons such as `<`. But it does not include bitwise operations such as `&` or `|` because those operations are not defined on `float64` values.

​	例如，以`min`的例子为例，在函数体内，任何`int64`和`float64`类型支持的操作都可以在类型参数`P`的值上执行。这包括所有基本的算术操作，但也包括比较操作，如`<`。但它不包括按位操作，如`&`或|，因为这些操作在`float64`值上没有定义。

## Comparable types 可比较类型 

In contrast to other unary and binary operations, `==` is defined on not just a limited set of [predeclared types](https://go.dev/ref/spec#Types), but on an infinite variety of types, including arrays, structs, and interfaces. It is impossible to enumerate all these types in a constraint. We need a different mechanism to express that a type parameter must support `==` (and `!=`, of course) if we care about more than predeclared types.

​	与其他一元和二元操作不同，`==` 不仅定义在一组有限的[预定义类型]({{< ref "/docs/References/LanguageSpecification/Types">}})上，还定义在各种类型上，包括数组、结构体和接口。在约束中列举所有这些类型是不可能的。如果我们关心的不仅仅是预定义类型，我们需要一种不同的机制来表示类型参数必须支持 `==`（当然也包括 `!=`）。

We solve this problem through the predeclared type [`comparable`](https://go.dev/ref/spec#Predeclared_identifiers), introduced with Go 1.18. `comparable` is an interface type whose type set is the infinite set of comparable types, and that may be used as a constraint whenever we require a type argument to support `==`.

​	我们通过引入Go 1.18中的预定义类型 [comparable]({{< ref "/docs/References/LanguageSpecification/DeclarationsAndScope#predeclared-identifiers--预先声明的标识符" >}}) 来解决这个问题。`comparable` 是一个接口类型，其类型集合是可比较类型的无限集合，并且可以在我们需要一个类型实参支持 `==` 的情况下用作约束。

Yet, the set of types comprised by `comparable` is not the same as the set of all [comparable types](https://go.dev/ref/spec#Comparison_operators) defined by the Go spec. [By construction](https://go.dev/ref/spec#Interface_types), a type set specified by an interface (including `comparable`) does not contain the interface itself (or any other interface). Thus, an interface such as `any` is not included in `comparable`, even though all interfaces support `==`. What gives?

​	然而，`comparable` 包含的类型集合与Go规范中定义的所有[可比较类型]({{< ref "/docs/References/LanguageSpecification/Expressions#comparison-operators-比较运算符" >}})的集合并不相同。[通过构造]({{< ref "/docs/References/LanguageSpecification/Types#interface-types-接口型" >}})，由接口（包括 `comparable`）指定的类型集合不包含接口本身（或任何其他接口）。因此，即使所有接口都支持 `==`，接口类型（如 `any`）也不包含在 `comparable` 中。这是为什么呢？

Comparison of interfaces (and of composite types containing them) may panic at run time: this happens when the dynamic type, the type of the actual value stored in the interface variable, is not comparable. Consider our original `lookupTable` example: it accepts arbitrary values as keys. But if we try to enter a value with a key that does not support `==`, say a slice value, we get a run-time panic:

​	在运行时，接口类型的比较（以及包含接口的复合类型的比较）可能会引发 panic：当动态类型（即接口变量中存储的实际值的类型）不可比较时会发生这种情况。考虑我们之前的 `lookupTable` 示例：它接受任意类型的值作为键。但是，如果我们尝试使用不支持 `==` 的键值（例如切片值），就会发生运行时 panic：

```Go
lookupTable[[]int{}] = "slice"  // PANIC: runtime error: hash of unhashable type []int
```

By contrast, `comparable` contains only types that the compiler guarantees will not panic with `==`. We call these types *strictly comparable*.

​	相比之下，`comparable` 只包含编译器保证不会通过 `==` 引发 panic 的类型。我们称这些类型为**严格可比较**类型。

Most of the time this is exactly what we want: it’s comforting to know that `==` in a generic function won’t panic if the operands are constrained by `comparable`, and it is what we would intuitively expect.

​	大多数情况下，这正是我们想要的：在泛型函数中，当操作数受 `comparable` 约束时，我们很安心地知道 `==` 不会引发 panic，这也是我们直观地期望的。

Unfortunately, this definition of `comparable` together with the rules for constraint satisfaction prevented us from writing useful generic code, such as the `genericLookupTable` type shown earlier: for `any` to be an acceptable argument type, `any` must satisfy (and therefore implement) `comparable`. But the type set of `any` is larger than (not a subset of) the type set of `comparable` and therefore does not implement `comparable`.

​	不幸的是，`comparable` 的这个定义以及约束满足规则阻止了我们编写有用的泛型代码，例如前面展示的 `genericLookupTable` 类型：为了使 `any` 成为可接受的参数类型，`any` 必须满足（并且实现）`comparable`。但是 `any` 的类型集合比 `comparable` 的类型集合更大（不是它的子集），因此没有实现 `comparable`。

```Go
var lookupTable GenericLookupTable[any, string] // ERROR: any does not implement comparable (Go 1.18 and Go 1.19)
```

​	用户早早地认识到了这个问题，并迅速提出了大量的问题和建议（[#51338](https://go.dev/issue/51338)、[#52474](https://go.dev/issue/52474)、[#52531](https://go.dev/issue/52531)、[#52614](https://go.dev/issue/52614)、 [#52624](https://go.dev/issue/52624)、 [#53734](https://go.dev/issue/53734)等）。显然，这是一个我们需要解决的问题。

The “obvious” solution was simply to include even non-strictly comparable types in the `comparable` type set. But this leads to inconsistencies with the type set model. Consider the following example:

​	“显而易见”的解决方案是将非严格可比较类型包含在 `comparable` 类型集合中。但这会导致与类型集合模型的不一致。考虑以下示例：

```Go linenums="1"
func f[Q comparable]() { … }

func g[P any]() {
        _ = f[int] // (1) ok: int 实现了 comparable
        _ = f[P]   // (2) error: 类型参数 P 没有实现 comparable
        _ = f[any] // (3) error: any 没有实现 comparable（Go 1.18，Go 1.19）
}
```

Function `f` requires a type argument that is strictly comparable. Obviously it is ok to instantiate `f` with `int`: `int` values never panic on `==` and thus `int` implements `comparable` (case 1). On the other hand, instantiating `f` with `P` is not permitted: `P`’s type set is defined by its constraint `any`, and `any` stands for the set of all possible types. This set includes types that are not comparable at all. Hence, `P` doesn’t implement `comparable` and thus cannot be used to instantiate `f` (case 2). And finally, using the type `any` (rather than a type parameter constrained by `any`) doesn’t work either, because of exactly the same problem (case 3).

​	函数 `f` 要求传入一个严格可比较的类型实参。显然，使用 `int` 来实例化 `f` 是可以的：`int` 值在 `==` 操作上不会引发 panic，因此 `int` 实现了 `comparable`（情况1）。另一方面，使用 `P` 来实例化 `f` 是不允许的：`P` 的类型集合由其约束 `any` 定义，而 `any` 代表所有可能的类型的集合。这个集合中包括根本不可比较的类型。由于`P` 没有实现 `comparable`，因此无法用于实例化 `f`（情况2）。最后，使用类型 `any`（而不是由 `any` 约束的类型参数）也不起作用，原因正是同样的问题（情况3）。

Yet, we do want to be able to use the type `any` as type argument in this case. The only way out of this dilemma was to change the language somehow. But how?

​	然而，我们确实希望在这种情况下能够将类型 `any` 用作类型参数。摆脱这个困境的唯一方法是以某种方式改变该语言。但是如何做呢？

## Interface implementation vs constraint satisfaction - 接口实现 vs 约束满足 

As mentioned earlier, constraint satisfaction is interface implementation: a type argument `T` satisfies a constraint `C` if `T` implements `C`. This makes sense: `T` must be in the type set expected by `C` which is exactly the definition of interface implementation.

​	如前所述，约束满足即是接口实现：如果类型参数 T 实现了约束 C，则类型参数 T 满足约束 C。这是有道理的：T 必须在 C 期望的类型集合中，这正是接口实现的定义。

But this is also the problem because it prevents us from using non-strictly comparable types as type arguments for `comparable`.

​	但这也是问题所在，因为它阻止我们将非严格可比较类型用作 comparable 的类型参数。

So for Go 1.20, after almost a year of publicly discussing numerous alternatives (see the issues mentioned above), we decided to introduce an exception for just this case. To avoid the inconsistency, rather than changing what `comparable` means, we differentiated between *interface implementation*, which is relevant for passing values to variables, and *constraint satisfaction*, which is relevant for passing type arguments to type parameters. Once separated, we could give each of those concepts (slightly) different rules, and that is exactly what we did with proposal [#56548](https://go.dev/issue/56548).

​	因此，在经过近一年的公开讨论了许多替代方案之后（参见上述问题），我们决定在 Go 1.20 中为这种情况引入一个例外。为了避免不一致性，我们不改变 comparable 的含义，而是区分接口实现（适用于向变量传递值）和约束满足（适用于向类型参数传递类型参数）。一旦分开，我们可以给这两个概念（稍微）不同的规则，这正是我们在提案 #56548 中所做的。

The good news is that the exception is quite localized in the [spec](https://go.dev/ref/spec#Satisfying_a_type_constraint). Constraint satisfaction remains almost the same as interface implementation, with a caveat:

​	好消息是，在规范中，这个例外是相当局限的。约束满足几乎与接口实现保持一致，但有一个注意事项：

> A type `T` satisfies a constraint `C` if
>
> 如果满足以下条件之一，类型 T 就满足约束 C：
>
> - `T` implements `C`; or
> - T 实现了 C；或者 
> - `C` can be written in the form `interface{ comparable; E }`, where `E` is a basic interface and `T` is [comparable](https://go.dev/ref/spec#Comparison_operators) and implements `E`.
> - C 可以写成 interface{ comparable; E } 的形式，其中 E 是一个基本接口，T 是可比较的并且实现了 E。 

The second bullet point is the exception. Without going too much into the formalism of the spec, what the exception says is the following: a constraint `C` that expects strictly comparable types (and which may also have other requirements such as methods `E`) is satisfied by any type argument `T` that supports `==` (and which also implements the methods in `E`, if any). Or even shorter: a type that supports `==` also satisfies `comparable` (even though it may not implement it).

​	第二个要点是例外情况。不过我们不会深入规范的形式主义，例外的含义如下：一个期望严格可比较类型的约束 C（还可能有其他要求，如方法 E）将被任何支持 == 的类型参数 T 所满足（如果有的话，还需要实现 E 中的方法）。或者更简洁地说：支持 == 的类型也满足 comparable（即使可能没有实现它）。

We can immediately see that this change is backward-compatible: before Go 1.20, constraint satisfaction was the same as interface implementation, and we still have that rule (1st bullet point). All code that relied on that rule continues to work as before. Only if that rule fails do we need to consider the exception.

​	我们可以立即看到，这个改变是向后兼容的：在 Go 1.20 之前，约束满足与接口实现是相同的，我们仍然保留了这个规则（第一个要点）。所有依赖于该规则的代码都会继续正常工作。只有在该规则失败时，我们才需要考虑例外情况。

Let’s revisit our previous example:

​	让我们重新看一下之前的例子：

```Go linenums="1"
func f[Q comparable]() { … }

func g[P any]() {
        _ = f[int] // (1) ok: int satisfies comparable
        _ = f[P]   // (2) error: type parameter P does not satisfy comparable
        _ = f[any] // (3) ok: satisfies comparable (Go 1.20)
}
```

Now, `any` does satisfy (but not implement!) `comparable`. Why? Because Go permits `==` to be used with values of type `any` (which corresponds to the type `T` in the spec rule), and because the constraint `comparable` (which corresponds to the constraint `C` in the rule) can be written as `interface{ comparable; E }` where `E` is simply the empty interface in this example (case 3).

​	现在，any 满足（但不实现）comparable。为什么呢？因为 Go 允许将 == 用于 any 类型的值（对应规范规则中的类型 T），并且因为约束 comparable（对应规则中的约束 C）可以写成 interface{ comparable; E } 的形式，其中 E 在这个例子中是空接口（情况3）。

Interestingly, `P` still does not satisfy `comparable` (case 2). The reason is that `P` is a type parameter constrained by `any` (it *is not* `any`). The operation `==` is *not* available with all types in the type set of `P` and thus not available on `P`; it is not a [comparable type](https://go.dev/ref/spec#Comparison_operators). Therefore the exception doesn’t apply. But this is ok: we do like to know that `comparable`, the strict comparability requirement, is enforced most of the time. We just need an exception for Go types that support `==`, essentially for historical reasons: we always had the ability to compare non-strictly comparable types.

​	有趣的是，P 仍然不满足 comparable（情况2）。原因是 P 是由 any 约束的类型参数（而不是 any 本身）。操作符 == 并不适用于 P 类型集合中的所有类型，因此在 P 上也不可用；它不是一个可比较的类型。因此例外并不适用。但这是可以的：我们确实希望大部分时间都强制执行 strict comparability 要求的 comparable。我们只是出于历史原因，对支持 == 的 Go 类型需要一个例外：我们一直可以比较非严格可比较的类型。

## Consequences and remedies

We gophers take pride in the fact that language-specific behavior can be explained and reduced to a fairly compact set of rules, spelled out in the language spec. Over the years we have refined these rules, and when possible made them simpler and often more general. We also have been careful to keep the rules orthogonal, always on the lookout for unintended and unfortunate consequences. Disputes are resolved by consulting the spec, not by decree. That is what we have aspired to since the inception of Go.

​	作为 Go 语言开发者，我们自豪于将特定于语言的行为解释为一组相当简洁的规则，并在语言规范中详细说明。多年来，我们不断完善这些规则，尽可能地使它们更简单、更通用。我们还小心地保持这些规则的正交性，时刻警惕意外和不幸的后果。争议通过查阅规范来解决，而不是通过命令来解决。这是我们自 Go 诞生以来一直追求的目标。

*One does not simply add an exception to a carefully crafted type system without consequences!*

​	在不带来后果的情况下，不可能简单地为精心构建的类型系统添加例外！

So where’s the catch? There’s an obvious (if mild) drawback, and a less obvious (and more severe) one. Obviously, we now have a more complex rule for constraint satisfaction which is arguably less elegant than what we had before. This is unlikely to affect our day-to-day work in any significant way.

那么问题在哪里呢？有一个明显的（尽管轻微）缺点，还有一个不太明显的（但更严重）缺点。显然，我们现在对约束满足有了一个更复杂的规则，可以说不如之前的那个优雅。这不太可能对我们的日常工作产生重大影响。

But we do pay a price for the exception: in Go 1.20, generic functions that rely on `comparable` are not statically type-safe anymore. The `==` and `!=` operations may panic if applied to operands of `comparable` type parameters, even though the declaration says that they are strictly comparable. A single non-comparable value may sneak its way through multiple generic functions or types by way of a single non-strictly comparable type argument and cause a panic. In Go 1.20 we can now declare

​	但是，我们为例外付出了代价：在 Go 1.20 中，依赖 comparable 的泛型函数不再具备静态类型安全性。尽管声明表示它们是严格可比较的，但 == 和 != 操作在应用于可比较类型参数的操作数时可能导致 panic。通过一个非严格可比较的类型参数，单个非可比较的值可能通过多个泛型函数或类型，并导致 panic。在 Go 1.20 中，我们现在可以声明：

```Go
var lookupTable genericLookupTable[any, string]
```

without compile-time error, but we will get a run-time panic if we ever use a non-strictly comparable key type in this case, exactly like we would with the built-in `map` type. We have given up static type safety for a run-time check.

​	而没有编译时错误，但如果我们在这种情况下使用非严格可比较的键类型，就会出现运行时 panic，就像我们使用内置的 map 类型一样。我们为了运行时检查而放弃了静态类型安全性。

There may be situations where this is not good enough, and where we want to enforce strict comparability. The following observation allows us to do exactly that, at least in limited form: type parameters do not benefit from the exception that we added to the constraint satisfaction rule. For instance, in our earlier example, the type parameter `P` in the function `g` is constrained by `any` (which by itself is comparable but not strictly comparable) and so `P` does not satisfy `comparable`. We can use this knowledge to craft a compile-time assertion of sorts for a given type `T`:

​	可能存在一些情况下这还不够好，我们希望强制执行严格的可比较性。以下观察结果允许我们以有限的形式做到这一点：类型参数不受我们添加到约束满足规则的例外的影响。例如，在我们之前的例子中，函数 g 中的类型参数 P 受到 any 的约束（它本身是可比较的，但不是严格可比较的），因此 P 不满足 comparable。我们可以利用这个知识来为给定的类型 T 创建一个类似于编译时断言的东西：

```Go
type T struct { … }
```

We want to assert that `T` is strictly comparable. It’s tempting to write something like:

​	我们想要断言 T 是严格可比较的。可能会尝试写出类似以下的代码：

```Go linenums="1"
// isComparable may be instantiated with any type that supports ==
// including types that are not strictly comparable because of the
// exception for constraint satisfaction.
// isComparable 可以实例化为任何支持 == 的类型，
// 包括那些由于约束满足的例外而不是严格可比较的类型。
func isComparable[_ comparable]() {}

// Tempting but not quite what we want: this declaration is also
// valid for types T that are not strictly comparable.
// 诱人但并不完全符合我们的要求：这个声明也适用于不是严格可比较的类型 T。

var _ = isComparable[T] // compile-time error if T does not support == 如果 T 不支持 ==，会在编译时出错
```

The dummy (blank) variable declaration serves as our “assertion”. But because of the exception in the constraint satisfaction rule, `isComparable[T]` only fails if `T` is not comparable at all; it will succeed if `T` supports `==`. We can work around this problem by using `T` not as a type argument, but as a type constraint:

​	这个虚拟（空白）变量声明作为我们的“断言”。但是由于约束满足规则中的例外，isComparable[T] 仅在 T 完全不可比较时才失败；如果 T 支持 ==，则会成功。我们可以通过将 T 用作类型约束而不是类型参数来解决这个问题：

```Go linenums="1"
func _[P T]() {
    _ = isComparable[P] // P supports == only if T is strictly comparable 如果 T 是严格可比较的，那么 P 支持 ==
}
```

Here is a [passing](https://go.dev/play/p/9i9iEto3TgE) and [failing](https://go.dev/play/p/5d4BeKLevPB) playground example illustrating this mechanism.

​	这里有一个展示这种机制的示例，包括了通过和不通过的情况。

## Final observations 最后的观察 

Interestingly, until two months before the Go 1.18 release, the compiler implemented constraint satisfaction exactly as we do now in Go 1.20. But because at that time constraint satisfaction meant interface implementation, we did have an implementation that was inconsistent with the language specification. We were alerted to this fact with [issue #50646](https://go.dev/issue/50646). We were extremely close to the release and had to make a decision quickly. In the absence of a convincing solution, it seemed safest to make the implementation consistent with the spec. A year later, and with plenty of time to consider different approaches, it seems that the implementation we had was the implementation we want in the first place. We have come full circle.

​	有趣的是，在 Go 1.18 发布前的两个月，编译器的约束满足实现方式与我们现在在 Go 1.20 中的实现方式完全相同。但是由于当时约束满足意味着接口实现，我们的实现与语言规范存在不一致。我们通过问题 #50646 得知了这个事实。我们离发布的时间非常近，必须迅速做出决策。在没有令人信服的解决方案的情况下，使实现与规范一致似乎是最安全的选择。一年后，经过充分的时间考虑不同的方法，我们发现最初的实现方式正是我们想要的实现方式。我们已经走了一圈。

As always, please let us know if anything doesn’t work as expected by filing issues at https://go.dev/issue/new.

​	如往常一样，如果有任何意外情况，请通过在 https://go.dev/issue/new 提交问题报告让我们知道。

Thank you!
