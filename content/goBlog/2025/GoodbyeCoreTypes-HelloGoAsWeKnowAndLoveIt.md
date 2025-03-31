+++
title = "告别核心类型 - 迎接我们熟悉和喜爱的 Go！"
date = 2025-03-31T14:29:10+08:00
weight = 940
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/coretypes](https://go.dev/blog/coretypes)

# Goodbye core types - Hello Go as we know and love it! 告别核心类型 - 迎接我们熟悉和喜爱的 Go！

Robert Griesemer  
26 March 2025  

The Go 1.18 release introduced generics and with that a number of new features, including type parameters, type constraints, and new concepts such as type sets. It also introduced the notion of a *core type*. While the former provide concrete new functionality, a core type is an abstract construct that was introduced for expediency and to simplify dealing with generic operands (operands whose types are type parameters). In the Go compiler, code that in the past relied on the [underlying type](https://go.dev/ref/spec/#Underlying_types) of an operand, now instead had to call a function computing the operand’s core type. In the language spec, in many places we just needed to replace “underlying type” with “core type”. What’s not to like?

​	Go 1.18 版本引入了泛型，随之而来的是许多新功能，包括类型参数、类型约束以及类型集等新概念。它还引入了 *核心类型* 的概念。前者提供了具体的新功能，而核心类型则是一个为了方便和简化处理泛型操作数（其类型是类型参数的操作数）而引入的抽象构造。在 Go 编译器中，过去依赖操作数的 [底层类型](https://go.dev/ref/spec/#Underlying_types) 的代码，现在需要调用一个计算操作数核心类型的函数。在语言规范中，许多地方只需将“底层类型”替换为“核心类型”。这有什么不好呢？

Quite a few things, as it turns out! To understand how we got here, it’s useful to briefly revisit how type parameters and type constraints work.

​	结果证明，有不少问题！为了理解我们是如何走到这一步的，简要回顾一下类型参数和类型约束的工作原理很有帮助。

## Type parameters and type constraints 类型参数和类型约束

A type parameter is a placeholder for a future type argument; it acts like a *type variable* whose value is known at compile time, similar to how a named constant stands for a number, string, or bool whose value is known at compile time. Like ordinary variables, type parameters have a type. That type is described by their *type constraint* which determines what operations are permitted on operands whose type is the respective type parameter.

​	类型参数是未来类型实参的占位符；它就像一个 *类型变量*，其值在编译时已知，类似于命名常量代表一个在编译时已知的数字、字符串或布尔值。与普通变量一样，类型参数也有类型。该类型由其 *类型约束* 描述，类型约束决定了在类型为相应类型参数的操作数上允许执行哪些操作。

Any concrete type that instantiates a type parameter must satisfy the type parameter’s constraint. This ensures that an operand whose type is a type parameter possesses all of the respective type constraint’s properties, no matter what concrete type is used to instantiate the type parameter.

​	实例化类型参数的任何具体类型都必须满足该类型参数的约束。这确保了类型为类型参数的操作数具备相应类型约束的所有属性，无论使用何种具体类型来实例化该类型参数。

In Go, type constraints are described through a mixture of method and type requirements which together define a *type set*: this is the set of all the types that satisfy all the requirements. Go uses a generalized form of interfaces for this purpose. An interface enumerates a set of methods and types, and the type set described by such an interface consists of all the types that implement those methods and that are included in the enumerated types.

​	在 Go 中，类型约束通过方法和类型需求的混合来描述，这些共同定义了一个 *类型集*：这是满足所有需求的类型集合。Go 为此使用了接口的广义形式。接口枚举了一组方法和类型，由此类接口描述的类型集包括实现这些方法并包含在枚举类型中的所有类型。

For instance, the type set described by the interface

​	例如，由以下接口描述的类型集

```Go
type Constraint interface {
    ~[]byte | ~string
    Hash() uint64
}
```

consists of all the types whose representation is `[]byte` or `string` and whose method set includes the `Hash` method.

​	包括所有表示为 `[]byte` 或 `string` 且方法集包含 `Hash` 方法的类型。

With this we can now write down the rules that govern operations on generic operands. For instance, the [rules for index expressions](https://go.dev/ref/spec#Index_expressions) state that (among other things) for an operand `a` of type parameter type `P`:

​	有了这个，我们现在可以写下管理泛型操作数的规则。例如，[索引表达式的规则](https://go.dev/ref/spec#Index_expressions) 规定（除其他外），对于类型参数类型 `P` 的操作数 `a`：

> The index expression `a[x]` must be valid for values of all types in `P`’s type set. The element types of all types in `P`’s type set must be identical. (In this context, the element type of a string type is `byte`.)
>
> ​	索引表达式 `a[x]` 必须对 `P` 类型集中所有类型的值有效。`P` 类型集中所有类型的元素类型必须相同。（在此上下文中，字符串类型的元素类型是 `byte`。）

These rules make it possible to index the generic variable `s` below ([playground](https://go.dev/play/p/M1LYKm3x3IB)):

​	这些规则使得以下泛型变量 `s` 的索引成为可能（[playground](https://go.dev/play/p/M1LYKm3x3IB)）：

```Go
func at[bytestring Constraint](s bytestring, i int) byte {
    return s[i]
}
```

The indexing operation `s[i]` is permitted because the type of `s` is `bytestring`, and the type constraint (type set) of `bytestring` contains `[]byte` and `string` types for which indexing with `i` is valid.

​	索引操作 `s[i]` 是允许的，因为 `s` 的类型是 `bytestring`，而 `bytestring` 的类型约束（类型集）包含 `[]byte` 和 `string` 类型，对于这些类型，使用 `i` 进行索引是有效的。

## Core types 核心类型

This type set-based approach is very flexible and in line with the intentions of the [original generics proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md): an operation involving operands of generic type should be valid if it is valid for any type permitted by the respective type constraint. To simplify matters with respect to the implementation, knowing that we would be able to relax rules later, this approach was *not* chosen universally. Instead, for instance, for [Send statements](https://go.dev/ref/spec#Send_statements), the spec states that

​	这种基于类型集的方法非常灵活，符合 [原始泛型提案](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md) 的意图：涉及泛型类型操作数的操作如果对相应类型约束允许的任何类型有效，则应有效。为了简化实现相关的事项，知道我们以后可以放宽规则，这种方法并未被普遍采用。例如，对于 [发送语句](https://go.dev/ref/spec#Send_statements)，规范规定：

> The channel expression’s *core type* must be a channel, the channel direction must permit send operations, and the type of the value to be sent must be assignable to the channel’s element type.
>
> ​	通道表达式的 *核心类型* 必须是通道，通道方向必须允许发送操作，且要发送的值的类型必须可赋值给通道的元素类型。

These rules are based on the notion of a core type which is defined roughly as follows:

​	这些规则基于核心类型的概念，其定义大致如下：

- If a type is not a type parameter, its core type is just its [underlying type](https://go.dev/ref/spec#Underlying_types). 如果类型不是类型参数，其核心类型就是其 [底层类型](https://go.dev/ref/spec#Underlying_types)。
- If the type is a type parameter, the core type is the single underlying type of all the types in the type parameter’s type set. If the type set has *different* underlying types, the core type doesn’t exist. 如果类型是类型参数，则核心类型是该类型参数类型集中所有类型的单一底层类型。如果类型集有 *不同* 的底层类型，则核心类型不存在。


For instance, `interface{ ~[]int }` has a core type (`[]int`), but the `Constraint` interface above does not have a core type. To make things more complicated, when it comes to channel operations and certain built-in calls (`append`, `copy`) the above definition of core types is too restrictive. The actual rules have adjustments that allow for differing channel directions and type sets containing both `[]byte` and `string` types.

​	例如，`interface{ ~[]int }` 具有核心类型（`[]int`），但上面的 `Constraint` 接口没有核心类型。更复杂的是，对于通道操作和某些内置调用（`append`、`copy`），上述核心类型的定义过于严格。实际规则进行了调整，允许不同的通道方向和包含 `[]byte` 和 `string` 类型的类型集。

There are various problems with this approach:

​	这种方法存在各种问题：

- Because the definition of core type must lead to sound type rules for different language features, it is overly restrictive for specific operations. For instance, the Go 1.24 rules for [slice expressions](https://go.dev/ref/spec#Slice_expressions) do rely on core types, and as a consequence slicing an operand of type `S` constrained by `Constraint` is not permitted, even though it could be valid.
  - 因为核心类型的定义必须为不同语言特性带来合理的类型规则，它对特定操作过于限制。例如，Go 1.24 的 [切片表达式规则](https://go.dev/ref/spec#Slice_expressions) 依赖核心类型，因此对受 `Constraint` 约束的类型 `S` 的操作数进行切片是不允许的，尽管它可能是有效的。

- When trying to understand a specific language feature, one may have to learn the intricacies of core types even when considering non-generic code. Again, for slice expressions, the language spec talks about the core type of the sliced operand, rather then just stating that the operand must be an array, slice, or string. The latter is more direct, simpler, and clearer, and doesn’t require knowing another concept that may be irrelevant in the concrete case.
  - 在试图理解特定语言特性时，即使考虑非泛型代码，也可能需要学习核心类型的复杂细节。同样，对于切片表达式，语言规范讨论的是被切片操作数的核心类型，而不是简单说明操作数必须是数组、切片或字符串。后者更直接、更简单、更清晰，且不需要了解在具体情况下可能无关的另一个概念。

- Because the notion of core types exists, the rules for index expressions, and `len` and `cap` (and others), which all eschew core types, appear as exceptions in the language rather than the norm. In turn, core types cause proposals such as [issue #48522](https://go.dev/issue/48522) which would permit a selector `x.f` to access a field `f` shared by all elements of `x`’s type set, to appear to add more exceptions to the language. Without core types, that feature becomes a natural and useful consequence of the ordinary rules for non-generic field access.
  - 因为核心类型的概念存在，索引表达式、`len` 和 `cap`（以及其他）的规则都避开了核心类型，在语言中显得像是例外而非常态。反过来，核心类型导致像 [issue #48522](https://go.dev/issue/48522) 这样的提案——允许选择器 `x.f` 访问 `x` 类型集中所有元素共享的字段 `f`——看似为语言增加了更多例外。如果没有核心类型，这一特性将成为非泛型字段访问普通规则的自然且有用的结果。


## Go 1.25

For the upcoming Go 1.25 release (August 2025) we decided to remove the notion of core types from the language spec in favor of explicit (and equivalent!) prose where needed. This has multiple benefits:

​	对于即将发布的 Go 1.25 版本（2025 年 8 月），我们决定从语言规范中移除核心类型的概念，转而使用必要的显式（且等效的！）文字说明。这有多个好处：

- The Go spec presents fewer concepts, making it easier to learn the language.
  - Go 规范呈现的概念更少，使得学习语言更容易。

- The behavior of non-generic code can be understood without reference to generics concepts.
  - 非泛型代码的行为无需参考泛型概念即可理解。

- The individualized approach (specific rules for specific operations) opens the door for more flexible rules. We already mentioned [issue #48522](https://go.dev/issue/48522), but there are also ideas for more powerful slice operations, and [improved type inference](https://go.dev/issue/69153).
  - 个性化方法（特定操作的特定规则）为更灵活的规则打开了大门。我们已经提到了 [issue #48522](https://go.dev/issue/48522)，但还有更强大的切片操作和 [改进的类型推断](https://go.dev/issue/69153) 的想法。


The respective [proposal issue #70128](https://go.dev/issue/70128) was recently approved and the relevant changes are already implemented. Concretely this means that a lot of prose in the language spec was reverted to its original, pre-generics form, and new paragraphs were added where needed to explain the rules as they pertain to generic operands. Importantly, no behavior was changed. The entire section on core types was removed. The compiler’s error messages were updated to not mention “core type” anymore, and in many cases error messages are now more specific by pointing out exactly which type in a type set is causing a problem.

​	相应的 [提案 issue #70128](https://go.dev/issue/70128) 最近已获批准，相关更改已实施。具体来说，这意味着语言规范中的许多文字恢复到了其原始的、前泛型形式，并在需要时添加了新段落，以解释与泛型操作数相关的规则。重要的是，没有改变任何行为。关于核心类型的整个部分已被移除。编译器的错误消息已更新，不再提及“核心类型”，在许多情况下，错误消息现在更具体，指出类型集中究竟哪个类型引发了问题。

Here is a sample of the changes made. For the built-in function `close`, starting with Go 1.18 the spec began as follows:

​	以下是所做更改的一个示例。对于内置函数 `close`，从 Go 1.18 开始，规范开头如下：

> For an argument `ch` with core type that is a channel, the built-in function `close` records that no more values will be sent on the channel.
>
> ​	对于核心类型为通道的参数 `ch`，内置函数 `close` 记录不再向该通道发送值。

A reader who simply wanted to know how `close` works, had to first learn about core types. Starting with Go 1.25, this section will again begin the same way it began before Go 1.18:

​	一个只想知道 `close` 如何工作的读者，必须首先了解核心类型。从 Go 1.25 开始，这一节将再次以 Go 1.18 之前的方式开头：

> For a channel `ch`, the built-in function `close(ch)` records that no more values will be sent on the channel.
>
> ​	对于通道 `ch`，内置函数 `close(ch)` 记录不再向该通道发送值。

This is shorter and easier to understand. Only when the reader is dealing with a generic operand will they have to contemplate the newly added paragraph:

​	这更短且更容易理解。只有当读者处理泛型操作数时，他们才需要考虑新增的段落：

> If the type of the argument to `close` is a type parameter all types in its type set must be channels with the same element type. It is an error if any of those channels is a receive-only channel.
>
> ​	如果 `close` 的参数类型是类型参数，其类型集中的所有类型必须是具有相同元素类型的通道。如果其中任何通道是只接收通道，则会出错。

We made similar changes to each place that mentioned core types. In summary, although this spec update does not affect any current Go program, it opens the door to future language improvements while making the language as it is today easier to learn and its spec simpler.

​	我们对每个提到核心类型的地方进行了类似的更改。总之，尽管这次规范更新不会影响任何当前的 Go 程序，但它为未来的语言改进打开了大门，同时使今天的语言更容易学习，其规范更简单。