+++
title = "类型统一规则"
date = 2024-02-27T20:00:50+08:00
weight = 20
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/ref/spec#Type_unification_rules](https://go.dev/ref/spec#Type_unification_rules)

### Type unification rules 类型统一规则

The type unification rules describe if and how two types unify. The precise details are relevant for Go implementations, affect the specifics of error messages (such as whether a compiler reports a type inference or other error), and may explain why type inference fails in unusual code situations. But by and large these rules can be ignored when writing Go code: type inference is designed to mostly "work as expected", and the unification rules are fine-tuned accordingly.

​	类型统一规则描述了两个类型是否统一以及如何统一。确切的细节与 Go 实现相关，会影响错误消息的具体内容（例如编译器报告类型推断或其他错误），并可能解释为什么类型推断在不寻常的代码情况下会失败。但总体而言，在编写 Go 代码时可以忽略这些规则：类型推断旨在大部分情况下“按预期工作”，并且统一规则也进行了相应调整。

Type unification is controlled by a *matching mode*, which may be *exact* or *loose*. As unification recursively descends a composite type structure, the matching mode used for elements of the type, the *element matching mode*, remains the same as the matching mode except when two types are unified for [assignability](https://go.dev/ref/spec#Assignability) (`≡A`): in this case, the matching mode is *loose* at the top level but then changes to *exact* for element types, reflecting the fact that types don't have to be identical to be assignable.

​	类型统一由匹配模式控制，该模式可以是精确的或宽松的。随着统一递归地向下遍历复合类型结构，用于类型元素的匹配模式（元素匹配模式）保持与匹配模式相同，但当两个类型统一以实现可赋值性时除外（ `≡A` ）：在这种情况下，匹配模式在顶层是宽松的，但随后更改为元素类型的精确匹配，这反映了类型不必相同即可赋值的事实。

Two types that are not bound type parameters unify exactly if any of following conditions is true:

​	如果满足以下任一条件，则两个未绑定类型参数完全统一：

- Both types are [identical](https://go.dev/ref/spec#Type_identity).
- 两个类型相同。
- Both types have identical structure and their element types unify exactly.
- 两个类型具有相同的结构，并且它们的元素类型完全统一。
- Exactly one type is an [unbound](https://go.dev/ref/spec#Type_inference) type parameter with a [core type](https://go.dev/ref/spec#Core_types), and that core type unifies with the other type per the unification rules for `≡A` (loose unification at the top level and exact unification for element types).
- 只有一个类型是具有核心类型的未绑定类型参数，并且该核心类型按照 `≡A` 的统一规则与另一个类型统一（顶层松散统一，元素类型完全统一）。

If both types are bound type parameters, they unify per the given matching modes if:
如果两个类型都是绑定类型参数，则它们按照给定的匹配模式统一，前提是：

- Both type parameters are identical.
- 两个类型参数相同。
- At most one of the type parameters has a known type argument. In this case, the type parameters are *joined*: they both stand for the same type argument. If neither type parameter has a known type argument yet, a future type argument inferred for one the type parameters is simultaneously inferred for both of them.
- 最多只有一个类型参数具有已知类型参数。在这种情况下，类型参数将被连接：它们都代表相同的类型参数。如果两个类型参数都没有已知类型参数，则同时为两个类型参数推断出的未来类型参数。
- Both type parameters have a known type argument and the type arguments unify per the given matching modes.
- 两个类型参数都有已知类型参数，并且类型参数按照给定的匹配模式统一。

A single bound type parameter `P` and another type `T` unify per the given matching modes if:
单个绑定类型参数 `P` 和另一个类型 `T` 按照给定的匹配模式统一，前提是：

- `P` doesn't have a known type argument. In this case, `T` is inferred as the type argument for `P`.
  
- `P` 没有已知类型参数。在这种情况下， `T` 被推断为 `P` 的类型参数。
  
- `P`  does have a known type argument  `A`  ,  `A`  and  `T`  unify per the given matching modes, and one of the following conditions is true:

- `P` 确实具有已知类型参数 `A` ， `A` 和 `T` 按照给定的匹配模式统一，并且以下条件之一为真：
  - Both `A` and `T` are interface types: In this case, if both `A` and `T` are also [defined](https://go.dev/ref/spec#Type_definitions) types, they must be [identical](https://go.dev/ref/spec#Type_identity). Otherwise, if neither of them is a defined type, they must have the same number of methods (unification of `A` and `T` already established that the methods match).
  - `A` 和 `T` 都是接口类型：在这种情况下，如果 `A` 和 `T` 也是已定义类型，则它们必须相同。否则，如果它们都不是已定义类型，则它们必须具有相同数量的方法（ `A` 和 `T` 的统一已经确定这些方法匹配）。
  - Neither `A` nor `T` are interface types: In this case, if `T` is a defined type, `T` replaces `A` as the inferred type argument for `P`.
  - `A` 和 `T` 都不是接口类型：在这种情况下，如果 `T` 是已定义类型，则 `T` 将替换 `A` 作为 `P` 的推断类型参数。

Finally, two types that are not bound type parameters unify loosely (and per the element matching mode) if:

​	最后，如果满足以下条件，则两个未绑定类型参数的类型将松散统一（并按照元素匹配模式）：

- Both types unify exactly.
- 两种类型完全统一。
- One type is a [defined type](https://go.dev/ref/spec#Type_definitions), the other type is a type literal, but not an interface, and their underlying types unify per the element matching mode.
- 一种类型是已定义类型，另一种类型是类型文字，但不是接口，并且它们的底层类型按照元素匹配模式统一。
- Both types are interfaces (but not type parameters) with identical [type terms](https://go.dev/ref/spec#Interface_types), both or neither embed the predeclared type [comparable](https://go.dev/ref/spec#Predeclared_identifiers), corresponding method types unify exactly, and the method set of one of the interfaces is a subset of the method set of the other interface.
- 两种类型都是接口（但不是类型参数），具有相同的类型项，两者都嵌入预声明的类型可比较，或者两者都不嵌入，相应的方法类型完全统一，并且其中一个接口的方法集是另一个接口的方法集的子集。
- Only one type is an interface (but not a type parameter), corresponding methods of the two types unify per the element matching mode, and the method set of the interface is a subset of the method set of the other type.
- 仅有一种类型是接口（但不是类型参数），两种类型相应的方法按照元素匹配模式统一，并且接口的方法集是另一种类型的方法集的子集。
- Both types have the same structure and their element types unify per the element matching mode.
- 两种类型具有相同的结构，并且它们的元素类型按照元素匹配模式统一。