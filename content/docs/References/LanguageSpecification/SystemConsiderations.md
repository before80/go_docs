+++
title = "系统考虑"
date = 2023-05-17T09:59:21+08:00
weight = 18
description = ""
isCJKLanguage = true
draft = false
+++
## System considerations 系统考虑

> 原文：[https://go.dev/ref/spec#System considerations](https://go.dev/ref/spec#System considerations)

### Package `unsafe ` - unsafe 包

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

​	`Pointer`是一个[指针类型](../Types#pointer-types-指针型)，但是`Pointer`的值不能被[解除引用](../Expressions#address-operators-地址运算符)。任何[底层类型](../Types)为`uintptr`的指针或值都可以被[转换](../Expressions#conversions-转换)为底层类型`Pointer`的类型，反之亦然。在`Pointer`和`uintptr`之间转换的效果是由实现定义的。

```go 
var f float64
bits = *(*uint64)(unsafe.Pointer(&f))

type ptr unsafe.Pointer
bits = *(*uint64)(ptr(&f))

var p ptr = nil
```

​		`Alignof`和`Sizeof`函数接收任意类型的表达式`x`，并分别返回假设变量`v`的对齐方式或大小，就像`v`通过`var v = x`声明一样。

​	`Offsetof`函数接收一个（可能是括号内的）[选择器](../Expressions#selectorss-选择器)`s.f`，表示由`s`或`*s`表示的结构体中的字段`f`，并返回相对于该结构体地址的字段偏移量（字节）。如果`f`是一个[嵌入字段](../Types#struct-types-结构体型)，它必须可以通过结构体的字段在没有指针间接的情况下访问。对于带有字段`f`的结构体`s`：

```go 
uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f) == uintptr(unsafe.Pointer(&s.f))
```

​	计算机体系结构可能要求内存地址被对齐；也就是说，变量的地址必须是一个因子的倍数，即变量类型的对齐方式。函数`Alignof`接收一个表示任何类型的变量的表达式，并返回（该类型的）变量的对齐方式，单位为字节。对于变量`x`：

```go 
uintptr(unsafe.Pointer(&x)) % unsafe.Alignof(x) == 0
```

​	如果`T`是一个[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)，或者它是一个包含可变大小的元素或字段的数组或结构体类型，则`T`类型的（变量）具有可变大小。否则，大小是常量。如果对 `Alignof`、`Offsetof` 和 `Sizeof` 的调用的实参（或 `Offsetof` 的选择器表达式 `s.f` 中的结构体 `s`）是恒定大小的类型，则它们是 `uintptr` 类型的编译时[常量表达式](../Expressions#constant-expressions-常量表达式)。

​	 `Add` 函数将 `len` 添加到 `ptr` 并返回更新后的指针 `unsafe.Pointer(uintptr(ptr) + uintptr(len))` 。`len`实参必须是[整数类型](../Types#numeric-types-数值型)或无类型的[常量](../Constants)。一个常量`len`实参必须可以用`int`类型的值[表示](../PropertiesOfTypesAndValues#representability-可表示性)；如果它是一个无类型的常量，则它会被赋予`int`类型。Pointer的[有效使用](https://go.dev/pkg/unsafe#Pointer)规则仍然适用。

​	`Slice` 函数返回一个切片，其底层数组从 `ptr` 开始，其长度和容量为 `len`。`Slice(ptr, len)` 等于

```go 
(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
```

**除了这样，还有一种特殊情况**，如果 `ptr` 是 `nil` 并且 `len` 是零，`Slice` 返回 `nil`。

​	`len`实参必须是[整数类型](../Types#numeric-types-数值型)或无类型的[常量](../Constants)。一个常量`len`实参必须是非负的，并且可以用`int`类型的值[表示](../PropertiesOfTypesAndValues#representability-可表示性)；如果它是一个无类型的常量，则它会被赋予`int`类型。在运行时，如果`len`是负的，或者如果`ptr`是`nil`而`len`不是0，会发生[运行时恐慌](../Run-timePanics)。

### Size and alignment guarantees 大小和对齐保证

对于数值型，以下大小是有保证的：

```go 
type                                 size in bytes

byte, uint8, int8                     1
uint16, int16                         2
uint32, int32, float32                4
uint64, int64, float64, complex64     8
complex128                           16
```

以下最小对齐特性得到了保证：

1. 对于任何类型的变量`x`：`unsafe.Alignof(x)`至少是1。
2. 对于结构体类型的变量`x`：`unsafe.Alignof(x)`是`x`的每个字段`f`的所有值`unsafe.Alignof(x.f)`中最大的一个，但至少是1。
3. 对于数组类型的变量`x`：`unsafe.Alignof(x)`与数组元素类型的变量的对齐方式相同。

​	如果结构体或数组类型不包含大小大于0的字段（或元素），那么它的大小就是零。两个不同的zero-size变量在内存中可能有相同的地址。