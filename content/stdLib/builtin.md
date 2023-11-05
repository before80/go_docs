+++
title = "builtin"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/builtin@go1.20.1

​	builtin包提供了 Go 的预声明标识符的文档。`这里所记录的条目实际上并不在 builtin 包中`，但它们在这里的描述使得 godoc 能够展示语言特殊标识符的文档。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/builtin/builtin.go;l=17)

``` go 
const (
	true  = 0 == 0 // 无类型布尔值。
	false = 0 != 0 // 无类型布尔值。
)
```

true 和 false 是两个无类型布尔值。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/builtin/builtin.go;l=107)

``` go 
const iota = 0 // // 无类型整数。
```

​	iota 是一个预声明标识符，表示当前 const 声明中(通常是在括号内)的无类型整数顺序编号。编号从零开始。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/builtin/builtin.go;l=111)

``` go 
var nil Type // Type 必须是指针、通道、函数、接口、映射或切片类型。
```

​	nil 是一个预声明标识符，表示指针、通道、函数、接口、映射或切片类型的零值。

## 函数

#### func append 

``` go 
func append(slice []Type, elems ...Type) []Type
```

​	append 内置函数将元素追加到切片的末尾。如果切片具有足够的容量，则目标切片就会被扩展以容纳新元素。如果没有足够的容量，将会分配一个新的底层数组。append 返回更新后的切片。因此，有必要将 append 的结果存储在变量中，通常是保存切片本身的变量：

```go 
slice = append(slice, elem1, elem2)
slice = append(slice, anotherSlice...)
```

​	`作为特殊情况，可以将字符串附加到字节切片`，如下所示：

```go 
slice = append([]byte("hello "), "world"...)
```

#### func cap 

``` go 
func cap(v Type) int
```

​	cap 内置函数返回 v 的容量，根据其类型不同而有所不同：

```
Array: v 中的元素数(与 len(v) 相同)。
Pointer to array: *v 中的元素数(与 len(v) 相同)。
Slice: 当重新分片时，切片可以达到的最大长度；如果 v 为 nil，则 cap(v) 为零。
Channel: 通道缓冲区的容量，以元素为单位；如果 v 为 nil，则 cap(v) 为零。
```

​	对于某些参数，例如简单的数组表达式，结果可以是常量。有关详细信息，请参见 Go 语言规范的 ["长度和容量" 部分]({{< ref "/langSpec/Built-inFunctions#length-and-capacity">}})。

#### func close 

``` go 
func close(c chan<- Type)
```

​	close 内置函数关闭一个通道，该通道必须是双向的或只能发送。它只应该由发送者执行，而不是接收者，并且在最后一个发送的值被接收后，会导致通道关闭。从关闭的信道 c 接收到最后一个值后，任何从 c 接收操作都将成功而不会被阻塞，返回信道元素的零值。如下形式：

```go 
x, ok := <-c
```

will also set ok to false for a closed and empty channel.

#### func complex 

``` go 
func complex(r, i FloatType) ComplexType
```

​	complex 内建函数使用两个浮点值构造一个复数值。实部和虚部必须具有相同的大小，可以是 float32 或 float64(或可分配给它们)，返回值将是相应的复数类型(对于 float32 是 complex64，对于 float64 是 complex128)。

#### func copy 

``` go 
func copy(dst, src []Type) int
```

​	copy 内建函数将源切片中的元素复制到目标切片中。(作为一个特殊情况，它也将字符串中的字节复制到字节切片中。)源和目标可能重叠。`Copy 返回复制的元素数量，这将是 len(src) 和 len(dst) 的最小值`。

#### func delete 

``` go 
func delete(m map[Type]Type1, key Type)
```

​	delete 内建函数从 map 中删除具有指定键(m[key])的元素。如果 m 为 nil 或没有这样的元素，则 delete 不执行任何操作。

#### func imag 

``` go 
func imag(c ComplexType) FloatType
```

​	imag 内建函数返回复数 c 的虚部。返回值将是与 c 的类型对应的浮点类型。

#### func len 

``` go 
func len(v Type) int
```

​	len 内建函数根据v类型返回 v 的长度：

```
Array: v 中的元素数。
Pointer to array: *v 中的元素数(即使 v 为 nil)。
Slice, or map: v 中的元素数；如果 v 为 nil，则 len(v) 为零。
String: v 中的字节数。
Channel: 通道缓冲区中排队的元素数(未读)；如果 v 为 nil，则 len(v) 为零。
```

​	对于某些参数，如字符串字面量或简单数组表达式，结果可以是常量。请参见 Go 语言规范的["长度和容量" 一节]({{< ref "/langSpec/Built-inFunctions#length-and-capacity">}})以获取详细信息。

#### func make 

``` go 
func make(t Type, size ...IntegerType) Type
```

​	make内建函数`用于分配`并`初始化`一个slice、map或channel类型的对象。与new函数不同的是，make的第一个参数是类型而不是值，并且返回值的类型与其参数的类型相同，而不是指向它的指针。其结果的规范取决于类型：

```
Slice: size参数指定其长度。该slice的容量等于其长度。可以提供第二个整数参数来指定不同的容量，它必须不小于长度。例如，make([]int, 0, 10)分配了一个大小为10的底层数组，并返回一个长度为0且容量为10的slice，该slice由该底层数组支持。
Map: 分配一个空的map，并预留足够空间来容纳指定数量的元素。可以省略size参数，此时会分配一个小的起始size。
Channel: 初始化通道的缓冲区以具有指定的缓冲区容量。如果size是0，或者省略了size，则通道是无缓冲的。
```

> ​	在 Go 编程语言中，使用 `make` 内置函数来创建 map 时，可以通过第二个可选参数来指定 map 可以容纳的元素数量。如果忽略这个参数，那么将会分配一个小的起始大小的空 map。
>
> ​	具体地，当指定了这个参数时，Go 会分配一个足够容纳指定数量元素的空 map。如果未指定该参数，则分配一个小的起始大小，这个大小通常是实现相关的，这意味着如果需要添加更多的元素到 map 中，那么 map 的大小将会根据需要自动增长。因此，如果您事先不知道要在 map 中存储多少个元素，则可以忽略第二个参数，让 Go 自动分配一个合适的初始大小。

#### func new 

``` go 
func new(Type) *Type
```

​	new内建函数用于分配内存。第一个参数是类型而不是值，并且返回的值是该类型的新分配的零值的指针。

#### func panic 

``` go 
func panic(v any)
```

​	panic内建函数会停止当前goroutine的正常执行。当一个函数`F`调用panic时，F的正常执行立即停止。由`F`推迟执行的任何函数都会以通常的方式运行，然后`F`返回其调用者。对调用者`G`来说，`F`的调用行为就像调用panic一样，终止`G`的执行并运行任何推迟的函数。这将继续进行，直到在执行goroutine中的所有函数都停止，按相反的顺序执行。此时，程序以非零的退出代码终止。这个终止序列被称为`panicking`，可以通过内建的`recover`函数来控制。

#### func print  <- go1.2

``` go 
func print(args ...Type)
```

​	print内建函数以实现特定的方式格式化其参数，并`将结果写入标准错误`。print函数对于引导和调试非常有用，但不能保证它会一直留在语言中。

#### func println  <- go1.2

``` go 
func println(args ...Type)
```

​	println内建函数以实现特定的方式格式化其参数，并`将结果写入标准错误`。参数之间始终添加空格，并附加一个换行符。println函数对于引导和调试非常有用，但不能保证它会一直留在语言中。

#### func real 

``` go 
func real(c ComplexType) FloatType
```

​	real内建函数返回复数c的实部。返回值将是与c类型相对应的浮点类型。

#### func recover 

``` go 
func recover() any
```

​	recover内置函数允许程序管理发生panic的goroutine的行为。`在延迟函数内部`执行调用recover，(但不包括由其调用的任何函数)会通过恢复正常执行并检索传递给panic调用的错误值来停止panic序列。`如果在延迟函数之外`调用recover，则不会停止panic序列。在这种情况下，或者当goroutine没有发生panic或者传递给panic的参数为nil时，recover返回nil。因此，recover的返回值报告goroutine是否发生panic。

```go 
package main

import (
	"fmt"
)

func main() {
	// 使用panic函数引发异常
	panicFunc()

	// 这句话不会被执行
	fmt.Println("程序结束")
}

// panicFunc 函数引发异常
func panicFunc() {
	defer func() {
		if r := recover(); r != nil {
			// 处理异常
			fmt.Println("捕获到异常：", r)
		}
	}()

	// 引发异常
	panic("出现了一个异常")
}
Output:

捕获到异常： 出现了一个异常
```



## 类型

### type ComplexType 

``` go 
type ComplexType complex64
```

​	ComplexType仅用于文档目的。它是一个占位符，代表复数类型：complex64或complex128。

### type FloatType 

``` go 
type FloatType float32
```

​	FloatType仅用于文档目的。它是一个占位符，代表浮点类型：float32或float64。

### type IntegerType 

``` go 
type IntegerType int
```

​	IntegerType仅用于文档目的。它是一个占位符，代表任何整数类型：int、uint、int8等。

### type Type 

``` go 
type Type int
```

​	Type仅用于文档目的。它是一个占位符，代表任何Go类型，但对于任何给定的函数调用表示相同的类型。

### type Type1 

``` go 
type Type1 int
```

​	Type1仅用于文档目的。它是一个占位符，代表任何Go类型，但对于任何给定的函数调用表示相同的类型。

### type any  <- go1.18

``` go 
type any = interface{}
```

​	any是interface{}的别名，在所有方面都等价于interface{}。

### type bool 

``` go 
type bool bool
```

​	bool是布尔值true和false的集合。

### type byte 

``` go 
type byte = uint8
```

​	byte是uint8的别名，与uint8在所有方面等效。它通常用于将字节值与8位无符号整数值区分开来。

### type comparable  <- go1.18

``` go 
type comparable interface{ comparable }
```

​	comparable是由所有可比较类型(布尔值、数字、字符串、指针、通道、由可比较类型组成的数组、字段均为可比较类型的结构体)实现的接口。`可比较接口只能用作类型参数约束，而不能作为变量类型`。

### type complex128 

``` go 
type complex128 complex128
```

​	complex128是具有float64实部和虚部的所有复数的集合。

### type complex64 

``` go 
type complex64 complex64
```

​	complex64是具有float32实部和虚部的所有复数的集合。

### type error 

``` go 
type error interface {
	Error() string
}
```

​	error内置接口类型是表示错误条件的常规接口类型，nil值表示没有错误。

### type float32 

``` go 
type float32 float32
```

​	float32是所有IEEE-754 32位浮点数的集合。

### type float64 

``` go 
type float64 float64
```

​	float64 是所有 IEEE-754 64 位浮点数的集合。

### type int 

``` go 
type int int
```

​	int 是至少为 32 位的带符号整数类型。然而，它是一个不同的类型，而不是 int32 的别名。

### type int16 

``` go 
type int16 int16
```

​	int16 是所有有符号 16 位整数的集合。范围：-32768 到 32767。

### type int32 

``` go 
type int32 int32
```

​	int32 是所有有符号 32 位整数的集合。范围：-2147483648 到 2147483647。

### type int64 

``` go 
type int64 int64
```

​	int64 是所有有符号 64 位整数的集合。范围：-9223372036854775808 到 9223372036854775807。

### type int8 

``` go 
type int8 int8
```

​	int8 是所有有符号 8 位整数的集合。范围：-128 到 127。

### type rune 

``` go 
type rune = int32
```

​	rune 是 int32 的别名，与 int32 在所有方面都是等价的。按照惯例，它用于区分字符值和整数值。

### type string 

``` go 
type string string
```

​	string 是所有由 8 位字节组成的字符串的集合，通常但不一定表示 UTF-8 编码的文本。字符串可以为空，但不能为 nil。string 类型的值是不可变的。

### type uint 

``` go 
type uint uint
```

​	类型 uint 是无符号整型，它至少占用 32 位。它是一个不同的类型，不是 uint32 的别名。

### type uint16 

``` go 
type uint16 uint16
```

​	类型 uint16 是所有 16 位无符号整型的集合。范围：0 到 65535。

### type uint32 

``` go 
type uint32 uint32
```

​	类型 uint32 是所有 32 位无符号整型的集合。范围：0 到 4294967295。

### type uint64 

``` go 
type uint64 uint64
```

​	类型 uint64 是所有 64 位无符号整型的集合。范围：0 到 18446744073709551615。

### type uint8 

``` go 
type uint8 uint8
```

​	类型 uint8 是所有 8 位无符号整型的集合。范围：0 到 255。

### type uintptr 

``` go 
type uintptr uintptr
```

​	类型 uintptr 是一个整数类型，它足够大，可以容纳任何指针的比特模式。