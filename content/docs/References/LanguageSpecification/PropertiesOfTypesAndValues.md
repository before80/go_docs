+++
title = "类型和值的属性"
date = 2023-05-17T09:59:21+08:00
weight = 8
description = ""
isCJKLanguage = true
draft = false
+++
## Properties of types and values 类型和值的属性

> 原文：[https://go.dev/ref/spec#Properties_of_types_and_values](https://go.dev/ref/spec#Properties_of_types_and_values)

### Underlying types 底层类型/基本类型

​	每个类型`T`都有一个底层类型。如果`T`是预先声明的布尔型、数值型或字符串型之一，或者是一个类型字面量，那么对应的底层类型就是`T`本身。否则，`T`的底层类型是`T`在其声明中所指的类型的底层类型。对于类型参数，则是其[类型约束](../DeclarationsAndScope#type-constraints)的底层类型，它总是一个接口。

```
type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)

func f[P any](x P) { … }
```

​	`string`、 `A1`、 `A2`、 `B1`和 `B2`的底层类型是 string。`[]B1`、`B3`和`B4`的底层类型是`[]B1`。`P`的底层类型是`interface{}`。

### Core types 核心类型

​	每个非接口类型`T`都有一个核心类型，它与`T`的[底层类型](#underlying-types)相同。

​	如果满足以下条件之一，那么接口`T`就有一个核心类型：

1. 存在一个单一的类型`U`，它是`T`的[类型集](../Types#interface-types)中所有类型的[底层类型](#underlying-types)；或者
2. `T`的类型集只包含具有相同元素类型`E`的[通道类型](../Types#channel-types)，并且所有定向通道具有相同的方向。

其他接口都没有核心类型。

接口的核心类型取决于满足的条件：

1. 类型`U`；或者
2. 如果`T`只包含双向通道，则为类型`chan E`；或者为 `chan<- E`或`<-chan E`类型，这取决于现存定向信道的方向。

根据定义，核心类型绝不是[已定义的类型](../DeclarationsAndScope#type-definitions)、[类型参数](../DeclarationsAndScope#type-parameter-declarations )或[接口类型](../Types#interface-types)。

具有核心类型的接口的示例：

```
type Celsius float32
type Kelvin  float32

interface{ int }                          // int
interface{ Celsius|Kelvin }               // float32
interface{ ~chan int }                    // chan int
interface{ ~chan int|~chan<- int }        // chan<- int
interface{ ~[]*data; String() string }    // []*data
```

没有核心类型的接口的示例：

```
interface{}                               // no single underlying type
interface{ Celsius|float64 }              // no single underlying type
interface{ chan int | chan<- string }     // channels have different element types
interface{ <-chan int | chan<- int }      // directional channels have different directions
```

​	一些操作（[切片表达式](../Expressions#slice-expressions)、[追加和复制](../Built-inFunctions#appending-to-copying-slices)）依赖于稍微宽松的核心类型形式，这些核心类型接受字节切片和字符串。具体来说，如果正好有两种类型，`[]byte`和`string`，它们是接口`T`的类型集中所有类型的底层类型，那么`T`的核心类型就被称为`bytestring`。

具有`bytestring`核心类型的接口的例子：

```
interface{ int }                          // int (same as ordinary core type) => int （与普通核心类型相同）
interface{ []byte | string }              // bytestring
interface{ ~[]byte | myString }           // bytestring
```

注意`bytestring`不是一个真正的类型；它不能用来声明变量(是由其他类型组成的)。它的存在只是为了描述一些从字节序列中读取的操作的行为，这些字节序列可能是字节切片或字符串。

### Type identity 类型一致性

两种类型要么相同，要么不同。

​	[命名类型](../Types)总是与任何其他类型不同。否则，如果两个类型的[底层类型](../Types)字面量在结构上是一致的，那么这两个类型就是相同的；也就是说，它们有相同的字面量结构，相应的组成部分拥有一致的类型。详细来说：

- 如果两个数组类型有一致的元素类型和相同的数组长度，那么它们就是一致的。
- 如果两个切片类型有一致的元素类型，那么它们就是一致的。
- 如果两个结构体类型具有相同的字段序列，并且相应的字段具有一致的名称、一致的类型和一致的标签，那么它们就是一致的。来自不同包的不可导出的字段名总是不同的。
- 如果两个指针类型有一致的基本类型，那么它们就是一致的。
- 如果两个函数类型有相同数量的参数和结果值，并且相应的参数和结果类型是相同的，并且两个函数要么都是可变的，要么都不是。参数和结果名称不需要匹配。
- 如果两个接口类型定义了相同的类型集，那么它们就是一致的。
- 如果两个映射类型有一致的键和元素类型，它们就是一致的。
- 如果两个通道类型有一致的元素类型和相同的方向，那么它们是一致的。
- 如果两个[实例化](../Expressions#instantiated)的类型的定义类型和所有类型参数都是一致的，那么它们就是一致的。

给出声明：

```
type (
	A0 = []string
	A1 = A0
	A2 = struct{ a, b int }
	A3 = int
	A4 = func(A3, float64) *A0
	A5 = func(x int, _ float64) *[]string

	B0 A0
	B1 []string
	B2 struct{ a, b int }
	B3 struct{ a, c int }
	B4 func(int, float64) *B0
	B5 func(x int, y float64) *A1

	C0 = B0
	D0[P1, P2 any] struct{ x P1; y P2 }
	E0 = D0[int, string]
)
```

这些类型是一致的：

```
A0, A1, and []string
A2 and struct{ a, b int }
A3 and int
A4, func(int, float64) *[]string, and A5

B0 and C0
D0[int, string] and E0
[]int and []int
struct{ a, b *B5 } and struct{ a, b *B5 }
func(x int, y float64) *[]string, func(int, float64) (result *[]string), and A5
```

​	`B0`和`B1`是不同的，因为它们是由不同的[类型定义](../DeclarationsAndScope#type-definitions)所创建的新类型；`func(int, float64) *B0`和`func(x int, y float64) *[]string`是不同的，因为`B0`与`[]string`是不同的；`P1`和`P2`是不同，因为它们是不同的类型参数。`D0[int, string]`和`struct{ x int; y string }`是不同的，因为前者是一个[实例化](../Expressions#instantiated)的定义类型，而后者是一个类型字面量（但它们仍然是[可分配的](#assignability)）。

### Assignability 可分配性

​	在以下这些情况中，`V`类型的值`x`是可以分配给`T`类型的[变量](../Variables)（"`x`可以分配给`T`"）：

- `V`和`T`是一致的。
- `V`和`T`有一致的[底层类型](#underlying-types)，但不是类型参数，并且`V`或`T`中至少有一个不是[命名类型](../Types)。
- `V`和`T`是具有一致元素类型的通道类型，`V`是一个双向通道，并且`V`或`T`中至少有一个不是[命名类型](../Types)。
- `T`是接口类型，但不是一个类型参数，并且`x`[实现](../Types#implementing-an-interface)了`T`。
- `x`是预先声明的标识符`nil`，并且`T`是一个指针、函数、切片、映射、通道或接口类型，但不是一个类型参数。
- `x`是可由`T`类型的值[表示](#representability )的非类型化的[常量]()。

​	除此之外，如果`x`的类型`V`或`T`是类型参数，并且满足以下条件之一，那么`x`也可以分配给类型`T`的变量：

- `x`是预先声明的标识符`nil`，`T`是类型参数，并且`x`可以分配给`T`的类型集中的每个类型。
- `V`不是[命名类型](../Types)，`T`是一个类型参数，并且`x`可以分配给`T`的类型集中的每个类型。
- `V`是类型参数，`T`不是[命名类型](../Types)，而`V`的类型集中的每个类型的值都可以分配给`T`。

### Representability 可表示性

​	如果满足以下条件之一，[常量](../Constants)`x`就可以被`T`类型的值所表示，其中`T`不是[类型参数](../DeclarationsAndScope#type-parameter-declarations )：

- `x`在由`T`[所确定的](../Types)值的集合中。
- `T`是[浮点类型](../Types/numeric-types)，并且`x`可以被舍入到`T`的精度而不会溢出。四舍五入使用的是IEEE 754的四舍五入到偶数的规则，但IEEE的负0被进一步简化为无符号0。请注意，常量值绝不会出现IEEE负零、NaN或无穷大。
- `T`是复数类型，`x`的[组成](../Built-inFunctions#manipulating-complex-numbers)`real(x)`和`imag(x)`可以用`T`的组成类型（`float32`或`float64`）的值表示。

​	如果`T`是类型参数，并且`x`可以由`T`的类型集中的每个类型的值来表示，那么`x`就可以由`T`类型的值来表示。

```
x                   T           x is representable by a value of T because

'a'                 byte        97 is in the set of byte values
97                  rune        rune is an alias for int32, and 97 is in the set of 32-bit integers
"foo"               string      "foo" is in the set of string values
1024                int16       1024 is in the set of 16-bit integers
42.0                byte        42 is in the set of unsigned 8-bit integers
1e10                uint64      10000000000 is in the set of unsigned 64-bit integers
2.718281828459045   float32     2.718281828459045 rounds to 2.7182817 which is in the set of float32 values
-1e-1000            float64     -1e-1000 rounds to IEEE -0.0 which is further simplified to 0.0
0i                  int         0 is an integer value
(42 + 0i)           float32     42.0 (with zero imaginary part) is in the set of float32 values


x                   T           x is not representable by a value of T because

0                   bool        0 is not in the set of boolean values
'a'                 string      'a' is a rune, it is not in the set of string values
1024                byte        1024 is not in the set of unsigned 8-bit integers
-1                  uint16      -1 is not in the set of unsigned 16-bit integers
1.1                 int         1.1 is not an integer value
42i                 float32     (0 + 42i) is not in the set of float32 values
1e1000              float64     1e1000 overflows to IEEE +Inf after rounding
```

### Method sets 方法集

​	类型的方法集确定了该类型的[操作数](../Expressions#operands)可以[调用](../Expressions#calls)的方法。每个类型都有一个与之相关的（可能是空的）方法集。

- [定义类型]()`T`的方法集包括所有用接收器类型`T`声明的[方法](../DeclarationsAndScope#method-declarations)。
- 指向[定义类型]()`T`的指针（`T`既不是指针也不是接口）的方法集是与接收器`*T`或`T`一起声明的所有方法的集合。
- [接口类型](../Types#interface-types)的方法集是该接口[类型集](../Types/interface-types)中每个类型的方法集的交集（最终的方法集通常只是接口中声明的方法集）。

​	进一步的规则，应用于包含嵌入字段的结构体（和结构体指针），会在关于[结构体类型](../Types#struct-types)的章节中描述。任何其他类型都有一个空的方法集。

​	在方法集中，每个方法必须有一个[唯一的](../DeclarationsAndScope#uniqueness-of-identifiers)非[空白](../DeclarationsAndScope#blank-identifier)[方法名](../Types#interface-types)。

