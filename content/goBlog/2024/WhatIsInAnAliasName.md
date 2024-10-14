+++
title = "名字（别名）中有什么？"
date = 2024-10-10T14:40:47+08:00
weight = 860
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# What's in an (Alias) Name? - 名字（别名）中有什么？

Robert Griesemer
17 September 2024

2024年9月17日

This post is about generic alias types, what they are, and why we need them.

​	这篇文章讨论了泛型别名类型，它们是什么，以及为什么我们需要它们。

## 背景 Background

Go was designed for programming at scale. Programming at scale means dealing with large amounts of data, but also large codebases, with many engineers working on those codebases over long periods of time.

​	Go 是为大规模编程设计的。大规模编程意味着处理大量的数据，但也意味着处理大规模的代码库，以及许多工程师在长期项目中的协作。

Go’s organization of code into packages enables programming at scale by splitting up large codebases into smaller, more manageable pieces, often written by different people, and connected through public APIs. In Go, these APIs consist of the identifiers exported by a package: the exported constants, types, variables, and functions. This includes the exported fields of structs and methods of types.

​	Go 通过将代码组织到包（package）中来实现大规模编程，这样可以将大型代码库拆分成更小、更易于管理的部分，这些部分通常由不同的人编写，并通过公共 API 连接。在 Go 中，这些 API 由包导出的标识符组成：导出的常量、类型、变量和函数。这包括结构体的导出字段和类型的方法。

As software projects evolve over time or requirements change, the original organization of the code into packages may turn out to be inadequate and require *refactoring*. Refactoring may involve moving exported identifiers and their respective declarations from an old package to a new package. This also requires that any references to the moved declarations must be updated so that they refer to the new location. In large codebases it may be unpractical or infeasible to make such a change atomically; or in other words, to do the move and update all clients in a single change. Instead, the change must happen incrementally: for instance, to “move” a function `F`, we add its declaration in a new package without deleting the original declaration in the old package. This way, clients can be updated incrementally, over time. Once all callers refer to `F` in the new package, the original declaration of `F` may be safely deleted (unless it must be retained indefinitely, for backward compatibility). Russ Cox describes refactoring in detail in his 2016 article on [Codebase Refactoring (with help from Go)](https://go.dev/talks/2016/refactor.article).

​	随着软件项目的发展或需求的变化，最初的代码包组织可能不再适合，且需要进行*重构*。重构可能包括将导出的标识符及其相应的声明从旧包移动到新包。这也要求对所有引用进行更新，以指向新位置。在大型代码库中，可能不切实际或无法一次性完成这些更改；换句话说，无法原子地完成移动和更新所有客户端。相反，变更必须逐步进行：例如，要将函数 `F` 从 `pkg1` 移动到 `pkg2`，可以在 `pkg2` 中添加一个新的函数声明 `F`（包装函数），而不删除旧包中的原始声明 `pkg1.F`。这样，客户端可以逐步更新。等到所有调用者都引用 `pkg2.F` 时，便可以安全地删除 `pkg1.F` 的原始声明（除非必须为向后兼容性保留该声明）。Russ Cox 在他2016年的文章[《通过 Go 进行代码库重构》](https://go.dev/talks/2016/refactor.article)中详细描述了重构过程。

Moving a function `F` from one package to another while also retaining it in the original package is easy: a wrapper function is all that’s needed. To move `F` from `pkg1` to `pkg2`, `pkg2` declares a new function `F` (the wrapper function) with the same signature as `pkg1.F`, and `pkg2.F` calls `pkg1.F`. New callers may call `pkg2.F`, old callers may call `pkg1.F`, yet in both cases the function eventually called is the same.

​	将函数 `F` 从一个包移动到另一个包，并保留在原始包中是很简单的：只需要一个包装函数即可。要将 `F` 从 `pkg1` 移动到 `pkg2`，`pkg2` 声明一个与 `pkg1.F` 相同签名的新函数 `F`，并且 `pkg2.F` 调用 `pkg1.F`。新调用者可以调用 `pkg2.F`，旧调用者可以调用 `pkg1.F`，但最终调用的函数是相同的。

Moving constants is similarly straightforward. Variables take a bit more work: one may have to introduce a pointer to the original variable in the new package or perhaps use accessor functions. This is less ideal, but at least it is workable. The point here is that for constants, variables, and functions, existing language features exist that permit incremental refactoring as described above.

​	移动常量同样简单。变量需要更多的工作：可能需要在新包中引入一个指向原始变量的指针，或使用访问器函数。这虽然不太理想，但至少可行。关键是对于常量、变量和函数，现有的语言特性允许如上所述的逐步重构。

But what about moving a type?

​	但如果是移动一个类型呢？

In Go, the [(qualified) identifier](https://go.dev/ref/spec#Qualified_identifiers), or just *name* for short, determines the *identity* of types: a type `T` [defined](https://go.dev/ref/spec#Type_definitions) and exported by a package `pkg1` is [different](https://go.dev/ref/spec#Type_identity) from an *otherwise identical* type definition of a type `T` exported by a package `pkg2`. This property complicates a move of `T` from one package to another while retaining a copy of it in the original package. For instance, a value of type `pkg2.T` is not [assignable](https://go.dev/ref/spec#Assignability) to a variable of type `pkg1.T` because their type names and thus their type identities are different. During an incremental update phase, clients may have values and variables of both types, even though the programmer’s intent is for them to have the same type.

​	在 Go 中，[(限定)标识符](https://go.dev/ref/spec#Qualified_identifiers)，或简称*名称*，决定了类型的*身份*：一个由包 `pkg1` [定义](https://go.dev/ref/spec#Type_definitions)并导出的类型 `T` 与包 `pkg2` 导出的*结构相同*的类型 `T` 是[不同的](https://go.dev/ref/spec#Type_identity)。这种性质使得在移动 `T` 的同时在原包中保留一份副本变得复杂。例如，类型 `pkg2.T` 的值不能[赋值](https://go.dev/ref/spec#Assignability)给类型 `pkg1.T` 的变量，因为它们的类型名称和类型身份是不同的。在逐步更新阶段，客户端可能会同时拥有两种类型的值和变量，即使程序员的意图是让它们具有相同的类型。

To solve this problem, [Go 1.9](https://go.dev/doc/go1.9) introduced the notion of a [*type alias*](https://go.dev/ref/spec#Alias_declarations). A type alias provides a new name for an existing type without introducing a new type that has a different identity.

​	为了解决这个问题，[Go 1.9](https://go.dev/doc/go1.9) 引入了[*类型别名*](https://go.dev/ref/spec#Alias_declarations)的概念。类型别名为现有类型提供了一个新名称，而不会引入具有不同身份的新类型。

In contrast to a regular [type definition](https://go.dev/ref/spec#Type_definitions)

​	与常规的[类型定义](https://go.dev/ref/spec#Type_definitions)不同：

```go
type T T0
```

which declares a *new type* that is never identical to the type on the right-hand side of the declaration, an [alias declaration](https://go.dev/ref/spec#Alias_declarations)

​	这种方式声明了一个*新类型*，它永远与声明右侧的类型不同。而[*别名声明*](https://go.dev/ref/spec#Alias_declarations)：

```go
type A = T  // the "=" indicates an alias declaration "=" 表示这是一个别名声明
```

declares only a *new name* `A` for the type on the right-hand side: here, `A` and `T` denote the same and thus identical type `T`.

​	只为右侧的类型 `T` 声明了一个*新名称* `A`：在这里，`A` 和 `T` 表示相同且完全相同的类型 `T`。

Alias declarations make it possible to provide a new name (in a new package!) for a given type while retaining type identity:

​	别名声明使得可以为给定类型提供一个新名称（甚至在一个新包中！），同时保留类型的身份：

```go
package pkg2

import "path/to/pkg1"

type T = pkg1.T
```

The type name has changed from `pkg1.T` to `pkg2.T` but values of type `pkg2.T` have the same type as variables of type `pkg1.T`.

​	类型名称从 `pkg1.T` 变为 `pkg2.T`，但类型 `pkg2.T` 的值与类型 `pkg1.T` 的变量具有相同的类型。

## 泛型别名类型 Generic alias types

[Go 1.18](https://go.dev/doc/go1.18) introduced generics. Since that release, type definitions and function declarations can be customized through type parameters. For technical reasons, alias types didn’t gain the same ability at that time. Obviously, there were also no large codebases exporting generic types and requiring refactoring.

​	[Go 1.18](https://go.dev/doc/go1.18) 引入了泛型。从那次发布起，类型定义和函数声明可以通过类型参数进行自定义。出于技术原因，别名类型当时没有获得相同的能力。当时也没有大规模代码库导出泛型类型并需要重构。

Today, generics have been around for a couple of years, and large codebases are making use of generic features. Eventually the need will arise to refactor these codebases, and with that the need to migrate generic types from one package to another.

​	今天，泛型已经存在了几年，并且大型代码库也在使用泛型功能。最终需要重构这些代码库，并因此需要将泛型类型从一个包迁移到另一个包。

To support incremental refactorings involving generic types, the future Go 1.24 release, planned for early February 2025, will fully support type parameters on alias types in accordance with proposal [#46477](https://go.dev/issue/46477). The new syntax follows the same pattern as it does for type definitions and function declarations, with an optional type parameter list following the identifier (the alias name) on the left-hand side. Before this change one could only write:

​	为了支持涉及泛型类型的逐步重构，计划于2025年2月初发布的 Go 1.24 将完全支持别名类型上的类型参数，符合提案[#46477](https://go.dev/issue/46477)。新语法遵循类型定义和函数声明的相同模式，别名声明左侧的标识符（别名名称）后面有一个可选的类型参数列表。在此更改之前，您只能写：

```go
type Alias = someType
```

but now we can also declare type parameters with the alias declaration:

​	但现在我们也可以在别名声明中声明类型参数：

```go
type Alias[P1 C1, P2 C2] = someType
```

Consider the previous example, now with generic types. The original package `pkg1` declared and exported a generic type `G` with a type parameter `P` that is suitably constrained:

​	考虑之前的例子，现在使用泛型类型。原包 `pkg1` 声明并导出一个带有类型参数 `P` 的泛型类型 `G`，该类型参数有适当的约束：

```go
package pkg1

type Constraint      someConstraint
type G[P Constraint] someType
```

If the need arises to provide access to the same type `G` from a new package `pkg2`, a generic alias type is just the ticket [(playground)](https://go.dev/play/p/wKOf6NbVtdw?v=gotip):

​	如果需要从新包 `pkg2` 提供对同一类型 `G` 的访问，泛型别名类型就是一个很好的解决方案[(playground)](https://go.dev/play/p/wKOf6NbVtdw?v=gotip)：

```go
package pkg2

import "path/to/pkg1"

type Constraint      = pkg1.Constraint  // pkg1.Constraint could also be used directly in G
type G[P Constraint] = pkg1.G[P]
```

Note that one **cannot** simply write

​	请注意，您**不能**简单地写：

```go
type G = pkg1.G
```

for a couple of reasons:

​	原因有两点：

1. Per [existing spec rules](https://go.dev/ref/spec#Type_definitions), generic types must be [instantiated](https://go.dev/ref/spec#Instantiations) when they are *used*. The right-hand side of the alias declaration uses the type `pkg1.G` and therefore type arguments must be provided. Not doing so would require an exception for this case, making the spec more complicated. It is not obvious that the minor convenience is worth the complication. 根据[现有的规范规则](https://go.dev/ref/spec#Type_definitions)，泛型类型在*使用*时必须[实例化](https://go.dev/ref/spec#Instantiations)。别名声明右侧使用了类型 `pkg1.G`，因此必须提供类型参数。不这样做将需要为此情况设置一个例外，这会使规范更加复杂。很明显，这点小方便不值得这种复杂化。
2. If the alias declaration doesn’t need to declare its own type parameters and instead simply “inherits” them from the aliased type `pkg1.G`, the declaration of `A` provides no indication that it is a generic type. Its type parameters and constraints would have to be retrieved from the declaration of `pkg1.G` (which itself might be an alias). Readability will suffer, yet readable code is one of the primary aims of the Go project. 如果别名声明不需要声明自己的类型参数，而是简单地“继承”了 `pkg1.G` 的类型参数，别名 `A` 的声明并未表明它是一个泛型类型。它的类型参数和约束必须从 `pkg1.G` 的声明中检索（它自己可能也是一个别名）。可读性会受到影响，而可读代码是 Go 项目的主要目标之一。

Writing down an explicit type parameter list may seem like an unnecessary burden at first, but it also provides additional flexibility. For one, the number of type parameters declared by the alias type doesn’t have to match the number of type parameters of the aliased type. Consider a generic map type:

​	显式地写出类型参数列表乍看之下似乎是一种不必要的负担，但它也提供了额外的灵活性。例如，别名类型声明的类型参数数量不必与被别名的类型的类型参数数量相同。考虑一个泛型映射类型：

```go
type Map[K comparable, V any] mapImplementation
```

If uses of `Map` as sets are common, the alias

​	如果 `Map` 被用作集合的情况很常见，那么别名：

```go
type Set[K comparable] = Map[K, bool]
```

might be useful [(playground)](https://go.dev/play/p/IxeUPGCztqf?v=gotip). Because it is an alias, types such as `Set[int]` and `Map[int, bool]` are identical. This would not be the case if `Set` were a [defined](https://go.dev/ref/spec#Type_definitions) (non-alias) type.

可能会很有用[(playground)](https://go.dev/play/p/IxeUPGCztqf?v=gotip)。因为它是一个别名，像 `Set[int]` 和 `Map[int, bool]` 这样的类型是相同的。如果 `Set` 是一个[定义](https://go.dev/ref/spec#Type_definitions)（非别名）类型，情况则不同。

Furthermore, the type constraints of a generic alias type don’t have to match the constraints of the aliased type, they only have to [satisfy](https://go.dev/ref/spec#Satisfying_a_type_constraint) them. For instance, reusing the set example above, one could define an `IntSet` as follows:

​	此外，泛型别名类型的类型约束不必与被别名的类型的约束相同，它们只需要[满足](https://go.dev/ref/spec#Satisfying_a_type_constraint)这些约束。例如，重用上面的集合示例，可以定义一个 `IntSet`，如下所示：

```go
type integers interface{ ~int | ~int8 | ~int16 | ~int32 | ~int64 }
type IntSet[K integers] = Set[K]
```

This map can be instantiated with any key type that satisfies the `integers` constraint [(playground)](https://go.dev/play/p/0f7hOAALaFb?v=gotip). Because `integers` satisfies `comparable`, the type parameter `K` may be used as type argument for the `K` parameter of `Set`, following the usual instantiation rules.

​	这个映射可以使用任何满足 `integers` 约束的键类型进行实例化[(playground)](https://go.dev/play/p/0f7hOAALaFb?v=gotip)。因为 `integers` 满足 `comparable`，因此类型参数 `K` 可以作为类型参数用于 `Set` 的 `K` 参数，遵循通常的实例化规则。

Finally, because an alias may also denote a type literal, parameterized aliases make it possible to create generic type literals [(playground)](https://go.dev/play/p/wql3NJaUs0o?v=gotip):

​	最后，因为别名也可以表示类型字面量，参数化别名使创建泛型类型字面量成为可能[(playground)](https://go.dev/play/p/wql3NJaUs0o?v=gotip)：

```go
type Point3D[E any] = struct{ x, y, z E }
```

To be clear, none of these examples are “special cases” or somehow require additional rules in the spec. They follow directly from the application of the existing rules put in place for generics. The only thing that changed in the spec is the ability to declare type parameters in an alias declaration.

​	需要明确的是，这些示例都不是“特殊情况”，也不需要在规范中增加额外的规则。它们直接遵循为泛型设定的现有规则。规范中唯一的变化是能够在别名声明中声明类型参数。

## 关于类型名称的插曲 An interlude about type names

Before the introduction of alias types, Go had only one form of type declarations:

​	在引入别名类型之前，Go 只有一种类型声明形式：

```
type TypeName existingType
```

This declaration creates a new and different type from an existing type and gives that new type a name. It was natural to call such types *named types* as they have a *type name* in contrast to unnamed [type literals](https://go.dev/ref/spec#Types) such as `struct{ x, y int }`.

​	这种声明从现有类型创建了一个新且不同的类型，并为该新类型赋予了一个名称。很自然地将这种类型称为*命名类型*，因为它们有一个*类型名称*，与未命名的[类型字面量](https://go.dev/ref/spec#Types)（如 `struct{ x, y int }`）相对。

With the introduction of alias types in Go 1.9 it became possible to give a name (an alias) to type literals, too. For instance, consider:

​	随着 Go 1.9 中别名类型的引入，现在可以为类型字面量赋予一个名称（别名）。例如，考虑：

```
type Point2D = struct{ x, y int }
```

Suddenly, the notion of a *named type* describing something that is different from a type literal didn’t make that much sense anymore, since an alias name clearly is a name for a type, and thus the denoted type (which might be a type literal, not a type name!) arguably could be called a “named type”.

​	突然之间，描述某些不同于类型字面量的东西的*命名类型*概念不再那么有意义了，因为别名显然是类型的一个名称，因此所指的类型（可能是一个类型字面量，不是类型名称！）可以被称为“命名类型”。

Because (proper) named types have special properties (one can bind methods to them, they follow different assignment rules, etc.), it seemed prudent to use a new term in order to avoid confusions. Thus, since Go 1.9, the spec calls the types formerly called named types *defined types*: only defined types have properties (methods, assignability restrictions, etc) that are tied to their names. Defined types are introduced through type definitions, and alias types are introduced through alias declarations. In both cases, names are given to types.

​	因为（适当的）命名类型有特殊属性（可以为它们绑定方法，它们遵循不同的赋值规则等），因此使用一个新术语似乎是明智的，以避免混淆。因此，自 Go 1.9 起，规范将以前称为命名类型的类型称为*定义类型*：只有定义类型才具有与其名称相关的属性（方法、可赋值性限制等）。定义类型是通过类型定义引入的，而别名类型是通过别名声明引入的。在这两种情况下，都是为类型赋予了名称。

The introduction of generics in Go 1.18 made things more complicated. Type parameters are types, too, they have a name, and they share rules with defined types. For instance, like defined types, two differently named type parameters denote different types. In other words, type parameters are named types, and furthermore, they behave similarly to Go’s original named types in some ways.

​	Go 1.18 中的泛型引入使事情变得更加复杂。类型参数也是类型，它们有一个名称，并且它们与定义类型共享规则。例如，像定义类型一样，两个具有不同名称的类型参数表示不同的类型。换句话说，类型参数是命名类型，并且在某些方面，它们的行为与Go的原始命名类型相似。

To top things off, Go’s predeclared types (`int`, `string` and so on) can only be accessed through their names, and like defined types and type parameters, are different if their names are different (ignoring for a moment the `byte` and `rune` alias types). The predeclared types truly are named types.

​	最后，Go的预声明类型（如 `int`、`string` 等）只能通过其名称访问，并且与定义类型和类型参数一样，如果它们的名称不同，则它们是不同的类型（暂时忽略 `byte` 和 `rune` 别名类型）。预声明类型确实是命名类型。

Therefore, with Go 1.18, the spec came full circle and formally re-introduced the notion of a [named type](https://go.dev/ref/spec#Types) which now comprises “predeclared types, defined types, and type parameters”. To correct for alias types denoting type literals the spec says: “An alias denotes a named type if the type given in the alias declaration is a named type.”

​	因此，随着 Go 1.18 的发布，规范正式重新引入了[命名类型](https://go.dev/ref/spec#Types)的概念，现在它包括“预声明类型、定义类型和类型参数”。为了纠正别名类型表示类型字面量的情况，规范规定：“如果别名声明中的类型是命名类型，则别名表示命名类型。”

Stepping back and outside the box of Go nomenclature for a moment, the correct technical term for a named type in Go is probably [*nominal type*](https://en.wikipedia.org/wiki/Nominal_type_system). A nominal type’s identity is explicitly tied to its name which is exactly what Go’s named types (now using the 1.18 terminology) are all about. A nominal type’s behavior is in contrast to a *structural type* which has behavior that only depends on its structure and not its name (if it has one in the first place). Putting it all together, Go’s predeclared, defined, and type parameter types are all nominal types, while Go’s type literals and aliases denoting type literals are structural types. Both nominal and structural types can have names, but having a name doesn’t mean the type is nominal, it just means it is named.

​	从 Go 的术语框架之外稍作思考，Go 中命名类型的正确技术术语可能是[*名义类型*](https://en.wikipedia.org/wiki/Nominal_type_system)。名义类型的身份明确与其名称相关，这正是 Go 的命名类型（现在使用1.18术语）所体现的。名义类型的行为与*结构类型*相对，结构类型的行为只依赖于其结构，而不是其名称（如果它有名称的话）。总而言之，Go的预声明、定义和类型参数都是名义类型，而Go的类型字面量和表示类型字面量的别名是结构类型。无论是名义类型还是结构类型都可以有名称，但拥有名称并不意味着该类型是名义的，它只是意味着它是命名的。

None of this matters for day-to-day use of Go and in practice the details can safely be ignored. But precise terminology matters in the spec because it makes it easier to describe the rules governing the language. So should the spec change its terminology one more time? It is probably not worth the churn: it is not just the spec that would need to be updated, but also a lot of supporting documentation. A fair number of books written on Go might become inaccurate. Furthermore, “named”, while less precise, is probably intuitively clearer than “nominal” for most people. It also matches the original terminology used in the spec, even if it now requires an exception for alias types denoting type literals.

​	这些对日常使用Go并没有实际影响，在实践中可以安全地忽略细节。但在规范中，精确的术语很重要，因为它使描述语言规则变得更加容易。那么规范是否应该再次更改其术语？可能不值得这样做：不仅是规范需要更新，还有很多支持文档也需要更新。大量关于Go的书籍可能会因此变得不准确。此外，“命名”虽然不够精确，但对于大多数人来说，可能比“名义”更直观。它也与规范中最初使用的术语一致，尽管现在需要对表示类型字面量的别名类型做出例外规定。

## 可用性 Availability

Implementing generic type aliases has taken longer than expected: the necessary changes required adding a new exported `Alias` type to [`go/types`](https://go.dev/pkg/go/types) and then adding the ability to record type parameters with that type. On the compiler side, the analogous changes also required modifications to the export data format, the file format that describes a package’s exports, which now needs to be able to describe type parameters for aliases. The impact of these changes is not confined to the compiler, but affects clients of `go/types` and thus many third-party packages. This was very much a change affecting a large code base; to avoid breaking things, an incremental roll-out over several releases was necessary.

​	实现泛型类型别名比预期花费了更长的时间：必要的更改要求将一个新的导出类型 `Alias` 添加到 [`go/types`](https://go.dev/pkg/go/types) 中，然后添加记录该类型参数的能力。在编译器方面，类似的更改也需要修改导出数据格式（描述包导出内容的文件格式），现在它需要能够描述别名的类型参数。这些更改的影响不仅限于编译器，还影响到 `go/types` 的客户端，从而影响到许多第三方包。这无疑是一次影响大规模代码库的更改；为了避免破坏，必须在多个版本中逐步推出。

After all this work, generic alias types will finally be available by default in Go 1.24.

​	经过这一切努力，泛型别名类型将在 Go 1.24 中默认可用。

To allow third-party clients to get their code ready, starting with Go 1.23, support for generic type aliases can be enabled by setting `GOEXPERIMENT=aliastypeparams` when invoking the `go` tool. However, be aware that support for exported generic aliases is still missing for that version.

​	为了让第三方客户端准备好他们的代码，从 Go 1.23 开始，可以通过在调用 `go` 工具时设置 `GOEXPERIMENT=aliastypeparams` 来启用对泛型类型别名的支持。但是，请注意，该版本仍然缺少对导出泛型别名的支持。

Full support (including export) is implemented at tip, and the default setting for `GOEXPERIMENT` will soon be switched so that generic type aliases are enabled by default. Thus, another option is to experiement with the latest version of Go at tip.

​	完整支持（包括导出）已在最新版本中实现，并且 `GOEXPERIMENT` 的默认设置很快将切换，以便默认启用泛型类型别名。因此，另一种选择是使用最新版本的 Go 进行实验。

As always, please let us know if you encounter any problems by filing an [issue](https://go.dev/issue/new); the better we test a new feature, the smoother the general roll-out.

​	一如既往，如果您遇到任何问题，请通过提交一个[问题](https://go.dev/issue/new)告诉我们；测试新功能越充分，正式推出的过程就会越顺利。

Thanks and happy refactoring!

​	感谢并祝您重构愉快！
