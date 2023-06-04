+++
title = "程序初始化和执行"
date = 2023-05-17T09:59:21+08:00
weight = 15
description = ""
isCJKLanguage = true
draft = false
+++
## Program initialization and execution 程序初始化和执行

> 原文：[https://go.dev/ref/spec#Program_initialization_and_execution](https://go.dev/ref/spec#Program_initialization_and_execution)

### The zero value 零值

​	当通过声明或调用`new`为一个[变量](../Variables)分配存储空间时，或者通过复合字面量或调用`make`创建一个新的值时，如果没有提供明确的初始化，该变量或值将被赋予一个默认值。这种变量或值的每个元素都被设置为其类型的零值：布尔类型为`false`，数值型为`0`，字符串为`""`，指针、函数、接口、切片、通道和映射为`nil`。`这种初始化是递归进行的`，因此，举例来说，如果没有指定值，结构体数组的每个元素字段都将被设为零值。

这两个简单的声明是等同的：

```go 
var i int
var i int = 0
```

在

```go 
type T struct { i int; f float64; next *T }
t := new(T)
```

之后，以下情况成立：

```go 
t.i == 0
t.f == 0.0
t.next == nil
```

 同样的情况，在以下情况下也会成立

```go 
var t T
```

### Package initialization 包的初始化

​	在包内，包级别变量的初始化是逐步进行的，每一步都会选择声明顺序中最早的变量，该变量与未初始化的变量没有依赖关系。

​	更确切地说，如果包级别变量还没有被初始化，并且没有初始化表达式，或者它的[初始化表达式](../DeclarationsAndScope#variable-declarations-变量声明)与未初始化的变量没有依赖关系，那么这个包级别变量就被认为可以被初始化。初始化的过程是重复初始化在声明顺序中最早的、准备好初始化的下一个包级别变量，直到没有准备好初始化的变量。

​	如果在这个过程结束时，仍有任何变量未被初始化，那么这些变量就是一个或多个初始化循环的一部分，程序是无效的。

​	变量声明左侧的多个变量（由右侧的单个（多值）表达式来初始化）是一起被初始化的：如果左侧的任何一个变量被初始化，那么所有这些变量都在同一步骤中被初始化。

```go 
var x = a
var a, b = f() // a and b are initialized together, before x is initialized => a 和 b 是在 x 被初始化之前一起被初始化的
```

为了实现包的初始化，[空白](../DeclarationsAndScope#blank-identifier-空白标识符)变量与声明中的其他变量一样被处理。

​	在多个文件中声明的变量的声明顺序是由文件呈现给编译器的顺序决定的。在第一个文件中声明的变量要在第二个文件中声明的任何变量之前声明，以此类推。

​	依赖项分析不依赖于变量的实际值，只依赖于源文件中对它们的词法引用，并进行传递性分析。例如，如果一个变量`x`的初始化表达式引用了一个函数，该函数的主体引用了变量`y`，那么`x`就依赖于`y`：

- 对变量或函数的引用是表示该变量或函数的标识符。
- 对一个方法`m`的引用是一个[方法值](../Expressions#method-values-方法值)或者形式为`t.m`的[方法表达式](../Expressions#method-expressions-方法表达式)，其中`t`的（静态）类型不是一个接口类型，并且方法`m`在`t`的[方法集](../PropertiesOfTypesAndValues#method-sets-方法集)中。是否调用结果函数值 `t.m` 并不重要。
- 如果`x`的初始化表达式或主体（对于函数和方法）包含对`y`的引用，或者对依赖于`y`的函数或方法的引用，那么变量、函数或方法`x`就依赖于变量`y`。

例如，给定的声明有

```go 
var (
	a = c + b  // == 9
	b = f()    // == 4
	c = f()    // == 5
	d = 3      // == 5 after initialization has finished
)

func f() int {
	d++
	return d
}
```

初始化顺序是 `d`，`b`，`c`，`a`。注意初始化表达式中子表达式的顺序是不相关的：在这个例子中`a = c + b`和`a = b + c`的初始化顺序是一样的。

​	依赖项分析是按包进行的；只考虑引用当前包中声明的变量、函数和（非接口）方法。如果变量之间存在其他隐藏的数据依赖关系，那么这些变量之间的初始化顺序是未指定的。

例如，给定的声明有

```go 
var x = I(T{}).ab()   // x has an undetected, hidden dependency on a and b => x 存在 在 a 和 b 上的未被发现的隐藏依赖
var _ = sideEffect()  // unrelated to x, a, or b => 与 x, a, 或 b 无关
var a = b
var b = 42

type I interface      { ab() []int }
type T struct{}
func (T) ab() []int   { return []int{a, b} }
```

变量`a`将在`b`之后被初始化，但`x`是在`b`之前、`b`和`a`之间还是在`a`之后被初始化，以及`sideEffect()`被调用的时刻（在`x`被初始化之前或之后）都是未指定的。

​	变量也可以使用在包块中声明的（不带实参，也没有结果参数）的名为`init`的函数来初始化。

```go 
func init() { … }
```

​	每个包`可以定义多个这样的（init）函数`，`甚至在一个源文件中也可以（有多个init函数）`。在包块中，`init`标识符只能用于声明`init`函数，但标识符本身并没有被[声明](../DeclarationsAndScope)。因此，`init`函数不能在程序中的任何地方被引用。

​	一个没有导入的包（即没有使用导入声明的包）的初始化方法是给它所有的包级别变量分配初始值，然后按照它们在源文件中出现的顺序调用所有的`init`函数，可能是在多个文件中出现，就像提交给编译器一样。如果一个包有导入（即使用了导入声明），那么在初始化包本身之前，被导入的包会先被初始化。如果多个包导入了一个包，那么被导入的包将只被初始化一次。 by construction，导入包可以保证不存在循环初始化依赖。

​	包的初始化 —— 变量的初始化和`init`函数的调用 —— 发生在单一goroutine中，按顺序，一次一个包。`init`函数可以启动其他goroutines，这些goroutines 可以与初始化代码同时运行。不过，初始化过程总是对`init`函数进行排序：在前一个函数返回之前，它不会调用下一个函数。

​	为了保证初始化行为的可重复性，我们鼓励构建系统以词法文件名的顺序向编译器提交属于同一包的多个文件。

### Program execution 程序执行

​	一个完整的程序是通过将一个名为（单个非导入包）`main`包与它所导入的所有包链接起来而创建的。主包必须有包名`main`，并声明一个不接受实参且不返回值的`main`函数。

```go 
func main() { … }
```

程序的执行从初始化`main`包开始，然后调用函数`main`。当该函数调用返回时，程序退出。它`不会等待`其他（非`main`）goroutine完成。