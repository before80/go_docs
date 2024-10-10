+++
title = "系统考虑"
date = 2023-05-17T09:59:21+08:00
weight = 18
description = ""
isCJKLanguage = true
type = "docs"
draft = false
+++
## System considerations 系统考虑

> 原文：[https://go.dev/ref/spec#System considerations](https://go.dev/ref/spec#System considerations)

### Package `unsafe ` - unsafe 包

The built-in package `unsafe`, known to the compiler and accessible through the [import path](https://go.dev/ref/spec#Import_declarations) `"unsafe"`, provides facilities for low-level programming including operations that violate the type system. A package using `unsafe` must be vetted manually for type safety and may not be portable. The package provides the following interface:

​	编译器知道的内置包`unsafe`，可以通过[导入路径](../Packages#import-declarations-导入声明) "`unsafe` "访问，它为低级编程提供了便利，包括违反（violate）类型系统的操作。使用`unsafe`的包必须手动审查其类型安全，并且可能无法移植。该包提供以下接口：

```go 
package unsafe

type ArbitraryType int  // shorthand for an arbitrary Go type; it is not a real type
type Pointer *ArbitraryType

func Alignof(variable ArbitraryType) uintptr
func Offsetof(selector ArbitraryType) uintptr
func Sizeof(variable ArbitraryType) uintptr

type IntegerType int  // shorthand for an integer type; it is not a real type
func Add(ptr Pointer, len IntegerType) Pointer
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
```

A `Pointer` is a [pointer type](https://go.dev/ref/spec#Pointer_types) but a `Pointer` value may not be [dereferenced](https://go.dev/ref/spec#Address_operators). Any pointer or value of [core type](https://go.dev/ref/spec#Core_types) `uintptr` can be [converted](https://go.dev/ref/spec#Conversions) to a type of core type `Pointer` and vice versa. The effect of converting between `Pointer` and `uintptr` is implementation-defined.

​	`Pointer`是一个[指针类型](../Types#pointer-types-指针型)，但是`Pointer`的值不能被[解除引用](../Expressions#address-operators-地址运算符)。任何[底层类型](../Types)为`uintptr`的指针或值都可以被[转换](../Expressions#conversions-转换)为底层类型`Pointer`的类型，反之亦然。在`Pointer`和`uintptr`之间转换的效果是由实现定义的。

```go 
var f float64
bits = *(*uint64)(unsafe.Pointer(&f))

type ptr unsafe.Pointer
bits = *(*uint64)(ptr(&f))

var p ptr = nil
```

The functions `Alignof` and `Sizeof` take an expression `x` of any type and return the alignment or size, respectively, of a hypothetical variable `v` as if `v` was declared via `var v = x`.

​	`Alignof`和`Sizeof`函数接收任意类型的表达式`x`，并分别返回假设变量`v`的对齐方式或大小，就像`v`通过`var v = x`声明一样。

The function `Offsetof` takes a (possibly parenthesized) [selector](https://go.dev/ref/spec#Selectors) `s.f`, denoting a field `f` of the struct denoted by `s` or `*s`, and returns the field offset in bytes relative to the struct's address. If `f` is an [embedded field](https://go.dev/ref/spec#Struct_types), it must be reachable without pointer indirections through fields of the struct. For a struct `s` with field `f`:

​	`Offsetof`函数接收一个（可能是括号内的）[选择器](../Expressions#selectorss-选择器)`s.f`，表示由`s`或`*s`表示的结构体中的字段`f`，并返回相对于该结构体地址的字段偏移量（字节）。如果`f`是一个[嵌入字段](../Types#struct-types-结构体型)，它必须可以通过结构体的字段在没有指针间接的情况下访问。对于带有字段`f`的结构体`s`：

```go 
uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f) == uintptr(unsafe.Pointer(&s.f))
```

Computer architectures may require memory addresses to be *aligned*; that is, for addresses of a variable to be a multiple of a factor, the variable's type's *alignment*. The function `Alignof` takes an expression denoting a variable of any type and returns the alignment of the (type of the) variable in bytes. For a variable `x`:

​	计算机体系结构可能要求内存地址被对齐；也就是说，变量的地址必须是一个因子的倍数，即变量类型的对齐方式。函数`Alignof`接收一个表示任何类型的变量的表达式，并返回（该类型的）变量的对齐方式，单位为字节。对于变量`x`：

```go 
uintptr(unsafe.Pointer(&x)) % unsafe.Alignof(x) == 0
```

A (variable of) type `T` has *variable size* if `T` is a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), or if it is an array or struct type containing elements or fields of variable size. Otherwise the size is *constant*. Calls to `Alignof`, `Offsetof`, and `Sizeof` are compile-time [constant expressions](https://go.dev/ref/spec#Constant_expressions) of type `uintptr` if their arguments (or the struct `s` in the selector expression `s.f` for `Offsetof`) are types of constant size.

​	如果`T`是一个[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)，或者它是一个包含可变大小的元素或字段的数组或结构体类型，则`T`类型的（变量）具有可变大小。否则，大小是常量。如果对 `Alignof`、`Offsetof` 和 `Sizeof` 的调用的实参（或 `Offsetof` 的选择器表达式 `s.f` 中的结构体 `s`）是恒定大小的类型，则它们是 `uintptr` 类型的编译时[常量表达式](../Expressions#constant-expressions-常量表达式)。

The function `Add` adds `len` to `ptr` and returns the updated pointer `unsafe.Pointer(uintptr(ptr) + uintptr(len))` [[Go 1.17]({{< ref "/langSpec/Appendix#go-117">}})]. The `len` argument must be of [integer type](https://go.dev/ref/spec#Numeric_types) or an untyped [constant](https://go.dev/ref/spec#Constants). A constant `len` argument must be [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; if it is an untyped constant it is given type `int`. The rules for [valid uses](https://go.dev/pkg/unsafe#Pointer) of `Pointer` still apply.

​	 `Add` 函数将 `len` 添加到 `ptr` 并返回更新后的指针 `unsafe.Pointer(uintptr(ptr) + uintptr(len))` 。`len`实参必须是[整数类型](../Types#numeric-types-数值型)或无类型的[常量](../Constants)。一个常量`len`实参必须可以用`int`类型的值[表示](../PropertiesOfTypesAndValues#representability-可表示性)；如果它是一个无类型的常量，则它会被赋予`int`类型。Pointer的[有效使用](https://go.dev/pkg/unsafe#Pointer)规则仍然适用。

The function `Slice` returns a slice whose underlying array starts at `ptr` and whose length and capacity are `len`. `Slice(ptr, len)` is equivalent to

​	`Slice` 函数返回一个切片，其底层数组从 `ptr` 开始，其长度和容量为 `len`。`Slice(ptr, len)` 等于

```go 
(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
```

except that, as a special case, if `ptr` is `nil` and `len` is zero, `Slice` returns `nil` [[Go 1.17]({{< ref "/langSpec/Appendix#go-117">}})].

**除了这样，还有一种特殊情况**，如果 `ptr` 是 `nil` 并且 `len` 是零，`Slice` 返回 `nil`。

The `len` argument must be of [integer type](https://go.dev/ref/spec#Numeric_types) or an untyped [constant](https://go.dev/ref/spec#Constants). A constant `len` argument must be non-negative and [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; if it is an untyped constant it is given type `int`. At run time, if `len` is negative, or if `ptr` is `nil` and `len` is not zero, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs [[Go 1.17]({{< ref "/langSpec/Appendix#go-117">}})].

​	`len`实参必须是[整数类型](../Types#numeric-types-数值型)或无类型的[常量](../Constants)。一个常量`len`实参必须是非负的，并且可以用`int`类型的值[表示](../PropertiesOfTypesAndValues#representability-可表示性)；如果它是一个无类型的常量，则它会被赋予`int`类型。在运行时，如果`len`是负的，或者如果`ptr`是`nil`而`len`不是0，会发生[运行时恐慌](../Run-timePanics) [[Go 1.17]({{< ref "/langSpec/Appendix#go-117">}})]。

The function `SliceData` returns a pointer to the underlying array of the `slice` argument. If the slice's capacity `cap(slice)` is not zero, that pointer is `&slice[:1][0]`. If `slice` is `nil`, the result is `nil`. Otherwise it is a non-`nil` pointer to an unspecified memory address [[Go 1.20]({{< ref "/langSpec/Appendix#go-120">}})].

​	函数 `SliceData` 返回一个指向 `slice` 参数的底层数组的指针。如果切片的容量 `cap(slice)` 不为零，则该指针为 `&slice[:1][0]` 。如果 `slice` 为 `nil` ，则结果为 `nil` 。否则，它是一个指向未指定内存地址的非 `nil` 指针 [ [Go 1.20]({{< ref "/langSpec/Appendix#go-120">}})]。

The function `String` returns a `string` value whose underlying bytes start at `ptr` and whose length is `len`. The same requirements apply to the `ptr` and `len` argument as in the function `Slice`. If `len` is zero, the result is the empty string `""`. Since Go strings are immutable, the bytes passed to `String` must not be modified afterwards. [[Go 1.20]({{< ref "/langSpec/Appendix#go-120">}})]

​	函数 `String` 返回一个 `string` 值，其底层字节从 `ptr` 开始，长度为 `len` 。 `ptr` 和 `len` 参数与函数 `Slice` 中的要求相同。如果 `len` 为零，则结果为空字符串 `""` 。由于 Go 字符串是不可变的，因此传递给 `String` 的字节之后不能被修改。[[Go 1.20]({{< ref "/langSpec/Appendix#go-120">}})]

The function `StringData` returns a pointer to the underlying bytes of the `str` argument. For an empty string the return value is unspecified, and may be `nil`. Since Go strings are immutable, the bytes returned by `StringData` must not be modified [[Go 1.20]({{< ref "/langSpec/Appendix#go-120">}})].

​	函数 `StringData` 返回一个指向 `str` 参数的底层字节的指针。对于空字符串，返回值未指定，可能为 `nil` 。由于 Go 字符串是不可变的，因此 `StringData` 返回的字节不能被修改 [[Go 1.20]({{< ref "/langSpec/Appendix#go-120">}})]。

### Size and alignment guarantees 大小和对齐保证

For the [numeric types](https://go.dev/ref/spec#Numeric_types), the following sizes are guaranteed:

​	对于数值型，以下大小是有保证的：

```go 
type                                 size in bytes

byte, uint8, int8                     1
uint16, int16                         2
uint32, int32, float32                4
uint64, int64, float64, complex64     8
complex128                           16
```

The following minimal alignment properties are guaranteed:

​	以下最小对齐特性得到了保证：

1. For a variable `x` of any type: `unsafe.Alignof(x)` is at least 1.
2. 对于任何类型的变量`x`：`unsafe.Alignof(x)`至少是1。
3. For a variable `x` of struct type: `unsafe.Alignof(x)` is the largest of all the values `unsafe.Alignof(x.f)` for each field `f` of `x`, but at least 1.
4. 对于结构体类型的变量`x`：`unsafe.Alignof(x)`是`x`的每个字段`f`的所有值`unsafe.Alignof(x.f)`中最大的一个，但至少是1。
5. For a variable `x` of array type: `unsafe.Alignof(x)` is the same as the alignment of a variable of the array's element type.
6. 对于数组类型的变量`x`：`unsafe.Alignof(x)`与数组元素类型的变量的对齐方式相同。

A struct or array type has size zero if it contains no fields (or elements, respectively) that have a size greater than zero. Two distinct zero-size variables may have the same address in memory.

​	如果结构体或数组类型不包含大小大于0的字段（或元素），那么它的大小就是零。两个不同的zero-size变量在内存中可能有相同的地址。