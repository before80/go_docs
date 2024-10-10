+++
title = "附录"
date = 2024-02-27T20:00:28+08:00
weight = 19
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/ref/spec#Appendix](https://go.dev/ref/spec#Appendix)

## Appendix 附录

### Language versions 语言版本

The [Go 1 compatibility guarantee](https://go.dev/doc/go1compat) ensures that programs written to the Go 1 specification will continue to compile and run correctly, unchanged, over the lifetime of that specification. More generally, as adjustments are made and features added to the language, the compatibility guarantee ensures that a Go program that works with a specific Go language version will continue to work with any subsequent version.

​	Go 1 兼容性保证确保按照 Go 1 规范编写的程序在该规范的生命周期内将继续编译并正确运行，且保持不变。更普遍地说，随着对语言进行调整并添加功能，兼容性保证确保与特定 Go 语言版本一起使用的 Go 程序将继续与任何后续版本一起使用。

For instance, the ability to use the prefix `0b` for binary integer literals was introduced with Go 1.13, indicated by [[Go 1.13]({{< ref "/langSpec/Appendix#go-113">}})] in the section on [integer literals](https://go.dev/ref/spec#Integer_literals). Source code containing an integer literal such as `0b1011` will be rejected if the implied or required language version used by the compiler is older than Go 1.13.

​	例如，使用前缀 `0b` 表示二进制整数字面量的功能在 Go 1.13 中引入，在整数字面量部分用 [Go 1.13] 表示。如果编译器使用的隐含或必需的语言版本低于 Go 1.13，则包含整数字面量（例如 `0b1011` ）的源代码将被拒绝。

The following table describes the minimum language version required for features introduced after Go 1.

​	下表描述了 Go 1 之后引入的功能所需的最低语言版本。

#### Go 1.9

- An [alias declaration](https://go.dev/ref/spec#Alias_declarations) may be used to declare an alias name for a type.
- 别名声明可用于为类型声明别名。

#### Go 1.13

- [Integer literals](https://go.dev/ref/spec#Integer_literals) may use the prefixes `0b`, `0B`, `0o`, and `0O` for binary, and octal literals, respectively.
- 整数字面量可以使用前缀 `0b` 、 `0B` 、 `0o` 和 `0O` 分别表示二进制和八进制字面量。
- Hexadecimal [floating-point literals](https://go.dev/ref/spec#Floating-point_literals) may be written using the prefixes `0x` and `0X`.
- 十六进制浮点字面量可以使用前缀 `0x` 和 `0X` 编写。
- The [imaginary suffix](https://go.dev/ref/spec#Imaginary_literals) `i` may be used with any (binary, decimal, hexadecimal) integer or floating-point literal, not just decimal literals.
- 虚部后缀 `i` 可用于任何（二进制、十进制、十六进制）整数或浮点字面量，而不仅仅是十进制字面量。
- The digits of any number literal may be [separated](https://go.dev/ref/spec#Integer_literals) (grouped) using underscores `_`.
- 任何数字字面量的数字都可以使用下划线 `_` 分隔（分组）。
- The shift count in a [shift operation](https://go.dev/ref/spec#Operators) may be a signed integer type.
- 移位操作中的移位计数可以是有符号整数类型。

#### Go 1.14

- Emdedding a method more than once through different [embedded interfaces](https://go.dev/ref/spec#Embedded_interfaces) is not an error.
- 通过不同的嵌入式接口多次嵌入方法不是错误。

#### Go 1.17

- A slice may be [converted](https://go.dev/ref/spec#Conversions) to an array pointer if the slice and array element types match, and the array is not longer than the slice.
- 如果切片和数组元素类型匹配，并且数组不长于切片，则可以将切片转换为数组指针。
- The built-in [package `unsafe`](https://go.dev/ref/spec#Package_unsafe) includes the new functions `Add` and `Slice`.
- 内置包 `unsafe` 包括新函数 `Add` 和 `Slice` 。

#### Go 1.18

The 1.18 release adds polymorphic functions and types ("generics") to the language. Specifically:

1.18 版本向语言添加了多态函数和类型（“泛型”）。具体来说：

- The set of [operators and punctuation](https://go.dev/ref/spec#Operators_and_punctuation) includes the new token `~`.
- 运算符和标点符号集包括新标记 `~` 。
- Function and type declarations may declare [type parameters](https://go.dev/ref/spec#Type_parameter_declarations).
- 函数和类型声明可以声明类型参数。
- Interface types may [embed arbitrary types](https://go.dev/ref/spec#General_interfaces) (not just type names of interfaces) as well as union and `~T` type elements.
- 接口类型可以嵌入任意类型（不仅仅是接口的类型名称），以及联合和 `~T` 类型元素。
- The set of [predeclared](https://go.dev/ref/spec#Predeclared_identifiers) types includes the new types `any` and `comparable`.
- 预声明类型集合包括新类型 `any` 和 `comparable` 。

#### Go 1.20

- A slice may be [converted](https://go.dev/ref/spec#Conversions) to an array if the slice and array element types match and the array is not longer than the slice.
- 如果切片和数组元素类型匹配并且数组不长于切片，则可以将切片转换为数组。
- The built-in [package `unsafe`](https://go.dev/ref/spec#Package_unsafe) includes the new functions `SliceData`, `String`, and `StringData`.
- 内置包 `unsafe` 包括新函数 `SliceData` 、 `String` 和 `StringData` 。
- [Comparable types](https://go.dev/ref/spec#Comparison_operators) (such as ordinary interfaces) may satisfy `comparable` constraints, even if the type arguments are not strictly comparable.
- 可比较类型（例如普通接口）可以满足 `comparable` 约束，即使类型参数不是严格可比较的。

#### Go 1.21

- The set of [predeclared](https://go.dev/ref/spec#Predeclared_identifiers) functions includes the new functions `min`, `max`, and `clear`.
- 预声明函数集合包括新函数 `min` 、 `max` 和 `clear` 。
- [Type inference](https://go.dev/ref/spec#Type_inference) uses the types of interface methods for inference. It also infers type arguments for generic functions assigned to variables or passed as arguments to other (possibly generic) functions.
- 类型推断使用接口方法的类型进行推断。它还推断分配给变量或作为参数传递给其他（可能是泛型）函数的泛型函数的类型参数。

#### Go 1.22

- In a ["for" statement](https://go.dev/ref/spec#For_statements), each iteration has its own set of iteration variables rather than sharing the same variables in each iteration.
- 在“for”语句中，每次迭代都有自己的一组迭代变量，而不是在每次迭代中共享相同的变量。
- A "for" statement with ["range" clause](https://go.dev/ref/spec#For_range) may iterate over integer values from zero to an upper limit.
- 带有 "range" 子句的 "for" 语句可以迭代从零到上限的整数值。