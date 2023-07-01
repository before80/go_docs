+++
title = "类型"
date = 2023-05-17T09:59:21+08:00
weight = 7
description = ""
isCJKLanguage = true
type = "docs"
draft = false

+++
## Types 类型

> 原文：[https://go.dev/ref/spec#Types ](https://go.dev/ref/spec#Types )

​	类型确定了一组值，以及针对这些值的特定操作和方法。如果一个类型有类型名称，可以通过类型名称来表示该类型，如果该类型是泛型的，则必须在类型名称后面跟上[类型参数](../Expressions#instantiations-实例化)。还可以使用`类型字面量`来指定类型，该类型由现有类型组成

``` go
type      = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
TypeName  = identifier | QualifiedIdent .
TypeArgs  = "[" TypeList [ "," ] "]" .
TypeList  = Type { "," Type } .
TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
            SliceType | MapType | ChannelType .
```

> 个人注释
>
> 怎么理解“如果该类型是泛型的，则必须在类型名称后面跟上[类型参数](../Expressions#instantiations-实例化)。”，请看如下示例：
>
> > 此示例来自[Go Tour的泛型]({{< ref "/docs/GoTour/Generics">}})
>
> ```go
> package main
> 
> import "fmt"
> 
> // Index 返回x在s中的索引，若没有找到x，则返回-1
> func Index[T comparable](s []T, x T) int {
> 	for i, v := range s {
> 		// v和x都是类型T，该类型具有可比较约束，因此我们可以在这里使用 ==。
> 		if v == x {
> 			return i
> 		}
> 	}
> 	return -1
> }
> 
> func main() {
> 	// 向Index传入整型切片
> 	si := []int{10, 20, 15, -10}
> 	fmt.Println(Index(si, 15)) // 2
> 
> 	// 向Index传入字符串切片
> 	ss := []string{"foo", "bar", "baz"}
> 	fmt.Println(Index(ss, "hello")) // -1
> }
> 
> ```
>
> ​	该示例中，Index 函数中的s和x参数即为类型参数，T类型即为泛型。
>
> 类型字面量是什么？有哪些类型字面量？
>
> ​	请看ChatGPT是如何回答的：
>
> > 在Go语言中，类型字面量是用于直接表示和描述类型的语法结构。它可以用来明确指定变量的类型或创建新的类型。
> >
> > 以下是Go语言中常见的类型字面量：
> >
> > 1. 基本类型字面量：
> >    - `int`：整数类型
> >    - `int8`：8位整数类型
> >    - `int16`：16位整数类型
> >    - `int32`：32位整数类型
> >    - `int64`：64位整数类型
> >    - `uint`：无符号整数类型
> >    - `uint8`：8位无符号整数类型
> >    - `uint16`：16位无符号整数类型
> >    - `uint32`：32位无符号整数类型
> >    - `uint64`：64位无符号整数类型
> >    - `uintptr`：指针整数类型
> >    - `float32`：32位浮点数类型
> >    - `float64`：64位浮点数类型
> >    - `complex64`：64位复数类型
> >    - `complex128`：128位复数类型
> >    - `bool`：布尔类型
> >    - `byte`：字节类型（等同于`uint8`）
> >    - `rune`：Unicode字符类型（等同于`int32`）
> >    - `string`：字符串类型
> > 2. 复合类型字面量：
> >    - 数组类型字面量：`[n]T`，其中 `n` 是数组的长度，`T` 是数组元素的类型。例如：`[5]int` 表示包含5个整数元素的数组类型。
> >    - 结构体类型字面量：`struct{}`，用于定义结构体类型。结构体类型由多个字段组成，每个字段都有自己的类型和名称。例如：`struct{ name string; age int }` 定义了一个包含 `name` 和 `age` 字段的匿名结构体类型。
> >    - 指针类型字面量：`*T`，其中 `T` 是指针指向的类型。例如：`*int` 表示整数类型的指针。
> >    - 函数类型字面量：`func(args) returnType`，用于定义函数类型。例如：`func(int) bool` 表示接受一个整数参数并返回布尔值的函数类型。
> >    - 接口类型字面量：`interface{}`，用于定义接口类型。接口类型定义了一组方法的集合，一个类型只要实现了接口中定义的所有方法，就被视为实现了该接口。
> >    - 切片类型字面量：`[]T`，其中 `T` 是切片元素的类型。例如：`[]int` 表示整数类型的切片。
> >    - 映射类型字面量：`map[T]U`，其中 `T` 是键的类型，`U` 是值的类型。例如：`map[string]int` 表示字符串到整数的映射类型。
> >    - 通道类型字面量：`chan T`，其中 `T` 是通道元素的类型。例如：`chan int` 表示整数类型的通道。
> >
> > 这些是Go语言中常见的类型字面量示例。通过组合和嵌套这些类型字面量，可以创建更复杂的类型。需要注意的是，还有一些其他的类型字面量，如函数字面量和接口字面量，用于直接定义匿名函数和匿名接口。

​	该语言[预先声明](../DeclarationsAndScope#predeclared-identifiers--预先声明的标识符)了某些类型的名称。其他类型是通过[类型声明](../DeclarationsAndScope#type-declarations-类型声明)或[类型参数列表](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)引入的。`复合类型`：数组、结构体、指针、函数、接口、切片、映射和通道类型 —— 可以用类型字面量来构造。

​	预先声明的类型、[已定义的类型](../DeclarationsAndScope#type-declarations-类型声明)和类型参数被称为`命名类型`。如果别名声明中给出的类型是命名类型，则别名也表示一个（新的）命名类型。

> 个人注释
>
> ​	在《Go语言精进之路》第24条 第226页，有这么一句话：
>
> 已有的类型（比如上面的I、T）被称为underlying类型，而新类型被称为defined类型。新定义的defined类型与原underlying类型是完全不同的类型，那么它们的方法集合上又会有什么关系呢？它们通过

### Boolean types 布尔型

​	布尔型表示由预先声明的常量`true`和`false`表示的一组布尔真值。预先声明的布尔类型是`bool`；它是一个[已定义的类型]({{<ref "/langSpec/DeclarationsAndScope#type-definitions-类型定义">}})。

### Numeric types 数值型

​	整数类型、浮点类型或复数类型分别表示整数、浮点或复数的值的集合。它们被统称为`数值类型`。预先声明的与体系结构无关的数值类型有：

```go 
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

​	The value of an *n*-bit integer is *n* bits wide and represented using [two's complement arithmetic](https://en.wikipedia.org/wiki/Two's_complement). =>仍有疑问？？

​	一个n bit整数的值是n bit宽，并用[二进制补码运算法（two's complement arithmetic）](https://en.wikipedia.org/wiki/Two's_complement)表示。

​	还有一组预先声明的整数类型，其具体实现的大小因实现而异：

```go 
uint     要么32位要么64位
int      与uint大小相同
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value => 一个足够大的无符号整数，用于存储指针值的未解释位
```

​	为了避免可移植性问题，所有的数值类型都是[已定义的类型](../DeclarationsAndScope#type-definitions-类型定义)，因此除了 `byte` (它是`uint8`的别名)和 `rune` (它是`int32`的别名)之外，它们是截然不同的。 当不同的数值类型在表达式或赋值中混合使用时，需要进行显式转换。例如，int32和int不是相同类型，尽管它们在一个特定的体系结构上可能具有相同的大小。

### String types 字符串型

​	字符串类型表示字符串值的集合。字符串值是（可能为空的）字节序列。`字节数`被称为字符串的`长度`，并且永远不会是负数。字符串是不可变的：一旦创建，就无法改变字符串的内容。预先声明的字符串类型是`string`；它是一种[已定义的类型](../DeclarationsAndScope#type-declarations-类型定义)。

​	可以使用内置函数 `len` 查找字符串 `s` 的长度。如果字符串是常量，那么长度就是编译时常量。字符串的字节可以通过整数[索引](../Expressions#index-expressions-索引表达式)0到`len(s)-1`来访问。`取这样一个元素的地址是非法的`；如果`s[i]`是字符串的第`i`个字节，那么`&s[i]`是无效的。

> 个人注释
>
> ​	什么是“`&s[i]`是无效的”？=>  `invalid operation`
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	s := "abcd"
> 	fmt.Println(&s[0]) // invalid operation: cannot take address of s[0] (value of type byte)
> 	fmt.Println(&s[1]) // invalid operation: cannot take address of s[1] (value of type byte)
> }
> 
> ```
>
> 

### Array types 数组型

​	数组是单类型的元素组成的编号序列，称为元素类型。`元素的数量`被称为数组的`长度`，并且永远不会是负数。

```
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```

​	长度是数组类型的一部分；它必须求值为一个非负[常数](../Constants)，该常数可由 int 类型的值[表示](../PropertiesOfTypesAndValues#representability)。数组`a`的长度可以用内置函数`len`发现。元素可以通过整数[索引](../Expressions#index-expressions-索引表达式)0到`len(a)-1`进行寻址。数组类型总是一维的，但是可以组成多维类型。

```go 
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

> 个人注释
>
> ​	解释下“`[2][2][2]float64`  // same as `[2]([2]([2]float64))`”？
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	arr1 := [2][2][2]float64{
> 		{
> 			{1, 2},
> 			{1, 2},
> 		},
> 		{
> 			{1, 2},
> 			{1, 2},
> 		},
> 	}
> 	arr2 := [2]([2]([2]float64)){
> 		{
> 			{1, 2},
> 			{1, 2},
> 		},
> 		{
> 			{1, 2},
> 			{1, 2},
> 		},
> 	}
> 
> 	fmt.Printf("arr1=%v\n", arr1) // arr1=[[[1 2] [1 2]] [[1 2] [1 2]]]
> 	fmt.Printf("arr2=%v\n", arr2) // arr2=[[[1 2] [1 2]] [[1 2] [1 2]]]
> }
> 
> ```
>
> 
>
> ​	数组的长度是在常量吗？是否可以在声明常量时作为常量的值？=> 是常量，可以作为常量的值！
>
> ```go
> package main
> 
> import "fmt"
> 
> var arr = [3]int{1, 2, 3}
> 
> const LEN = len(arr)
> 
> func main() {
> 	fmt.Println(LEN) // 3
> }
> 
> ```
>
> ​	是否可以对数组中的元素进行取地址操作？=> 可以
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	var arr = [3]int{1, 2, 3}
> 	fmt.Println(&arr[0]) // 0xc000010120
> 	fmt.Println(&arr[1]) // 0xc000010128
> 	fmt.Println(&arr[2]) // 0xc000010130
> }
> 
> ```
>
> 

### Slice types 切片型

​	切片是底层数组的连续段的描述符，并提供对该数组中编号的元素序列的访问。切片类型表示其元素类型的所有数组切片的集合。`元素的数量`被称为切片的`长度`，并且永远不会是负数。一个未初始化的切片的值是`nil`。

```
SliceType = "[" "]" ElementType .
```

​	切片`s`的长度可以通过内置函数`len`发现；与数组不同，它在运行过程中可能会发生变化。元素可以通过整数[索引](../Expressions#index-expressions-索引表达式)0到`len(s)-1`进行寻址。给定元素的切片索引可能小于底层数组中同一元素的索引。

> 个人注释
>
> ​	切片的长度是常量吗？是否可以在声明常量时作为常量的值？=> 不是常量，不可以作为常量的值！
>
> ```go
> package main
> 
> import "fmt"
> 
> var sli = []int{1, 2, 3}
> 
> const LEN = len(sli) // len(sli) (value of type int) is not constant
> 
> func main() {
> 	fmt.Println(LEN)
> }
> 
> ```
>
> ​	解释下“给定元素的切片索引可能小于底层数组中同一元素的索引。”？
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	arr := []int{0, 1, 2, 3, 4, 5}
> 
> 	sli := arr[2:]
> 
> 	fmt.Printf("arr=%v\n", arr)                                 // arr=[0 1 2 3 4 5]
> 	fmt.Printf("sli=%v\n", sli)                                 // sli=[2 3 4 5]
> 	fmt.Printf("len(arr)=%d,cap(arr)=%d\n", len(arr), cap(arr)) // len(arr)=6,cap(arr)=6
> 	fmt.Printf("len(sli)=%d,cap(sli)=%d\n", len(sli), cap(sli)) // len(sli)=4,cap(sli)=4
> 
> 	for i, v := range arr {
> 		if v == 2 {
> 			fmt.Printf("2在arr中的索引是%d\n", i) // 2
> 		}
> 	}
> 
> 	for i, v := range sli {
> 		if v == 2 {
> 			fmt.Printf("2在sli中的索引是%d\n", i) // 0
> 		}
> 	}
> 
> }
> 
> ```
>
> 

​	切片一旦被初始化，总是与保存其元素的底层数组相关联。因此，一个切片与它的底层数组和同一数组的其他切片共享存储；相反，不同的数组总是表示不同的存储。

> 个人注释
>
> ​	解释下“一个切片与它的底层数组和同一数组的其他切片共享存储”？
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	arr := []int{0, 1, 2, 3, 4, 5}
> 
> 	sli1 := arr[1:]
> 	sli2 := arr[2:]
> 
> 	fmt.Printf("arr=%v\n", arr)   // arr=[0 1 2 3 4 5]
> 	fmt.Printf("sli1=%v\n", sli1) // sli1=[1 2 3 4 5]
> 	fmt.Printf("sli2=%v\n", sli2) // sli2=[2 3 4 5]
> 
> 	sli1[1] = 22
> 	fmt.Printf("arr=%v\n", arr)   // arr=[0 1 22 3 4 5]
> 	fmt.Printf("sli1=%v\n", sli1) // sli1=[1 22 3 4 5]
> 	fmt.Printf("sli2=%v\n", sli2) // sli2=[22 3 4 5]
> 
> 	sli2[1] = 33
> 	fmt.Printf("arr=%v\n", arr)   // arr=[0 1 22 33 4 5]
> 	fmt.Printf("sli1=%v\n", sli1) // sli1=[1 22 33 4 5]
> 	fmt.Printf("sli2=%v\n", sli2) // sli2=[22 33 4 5]
> }
> 
> ```
>
> 

​	切片的底层数组可以超过切片的末端。容量是对这一范围的衡量：它是切片的长度和切片之外的数组长度之和；可以通过从原始切片切割一个新的切片来创建一个达到这个容量的切片。使用内置函数 `cap(a)`可以发现切片 `a` 的容量。

> 个人注释
>
> ​	解释下“容量是对这一范围的衡量：它是切片的长度和切片之外的数组长度之和”中的“切片之外的数组长度”是什么意思？以及“可以通过从原始切片切割一个新的切片来创建一个达到这个容量的切片”？
>
> => 可以想象成向右方向还不是（切片中的数组元素的）长度。
>
> => 只有从最左边的切割原始切片， 产生的切片的容量才和原始切片的容量一致！
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
> 
> 	sli1 := arr[1:4]
> 	sli2 := arr[1:5]
> 	sli3 := arr[1:6]
> 
> 	fmt.Printf("arr=%v,len(arr)=%d,cap(arr)=%d\n", arr, len(arr), cap(arr))       // arr=[0 1 2 3 4 5 6 7],len(arr)=8,cap(arr)=8
> 	fmt.Printf("sli1=%v,len(sli1)=%d,cap(sli1)=%d\n", sli1, len(sli1), cap(sli1)) // sli1=[1 2 3],len(sli1)=3,cap(sli1)=7
> 	fmt.Printf("sli2=%v,len(sli2)=%d,cap(sli2)=%d\n", sli2, len(sli2), cap(sli2)) // sli2=[1 2 3 4],len(sli2)=4,cap(sli2)=7
> 	fmt.Printf("sli3=%v,len(sli3)=%d,cap(sli3)=%d\n", sli3, len(sli3), cap(sli3)) // sli3=[1 2 3 4 5],len(sli3)=5,cap(sli3)=7
> 
> 	sli10 := sli1[0:1]
> 	sli20 := sli2[0:1]
> 	sli30 := sli3[0:1]
> 	fmt.Printf("sli10=%v,len(sli10)=%d,cap(sli10)=%d\n", sli10, len(sli10), cap(sli10)) // sli10=[1],len(sli10)=1,cap(sli10)=7
> 	fmt.Printf("sli20=%v,len(sli20)=%d,cap(sli20)=%d\n", sli20, len(sli20), cap(sli20)) // sli20=[1],len(sli20)=1,cap(sli20)=7
> 	fmt.Printf("sli30=%v,len(sli30)=%d,cap(sli30)=%d\n", sli30, len(sli30), cap(sli30)) // sli30=[1],len(sli30)=1,cap(sli30)=7
> 
> 	sli11 := sli1[1:2]
> 	sli21 := sli2[1:2]
> 	sli31 := sli3[1:2]
> 	fmt.Printf("sli11=%v,len(sli11)=%d,cap(sli11)=%d\n", sli11, len(sli11), cap(sli11)) // sli11=[2],len(sli11)=1,cap(sli11)=6
> 	fmt.Printf("sli21=%v,len(sli21)=%d,cap(sli21)=%d\n", sli21, len(sli21), cap(sli21)) // sli21=[2],len(sli21)=1,cap(sli21)=6
> 	fmt.Printf("sli31=%v,len(sli31)=%d,cap(sli31)=%d\n", sli31, len(sli31), cap(sli31)) // sli31=[2],len(sli31)=1,cap(sli31)=6
> 	// 从 sli10、sli11、sli20、sli21、sli30、sli31等可以看出，
> 	// 只有从最左边的切割原始切片，
> 	// 产生的切片的容量才和原始切片的容量一致！
> 
> 	sli4 := arr[2:4]
> 	sli5 := arr[2:5]
> 	sli6 := arr[2:6]
> 
> 	fmt.Printf("arr=%v,len(arr)=%d,cap(arr)=%d\n", arr, len(arr), cap(arr))       // arr=[0 1 2 3 4 5 6 7],len(arr)=8,cap(arr)=8
> 	fmt.Printf("sli4=%v,len(sli4)=%d,cap(sli4)=%d\n", sli4, len(sli4), cap(sli4)) // sli4=[2 3],len(sli4)=2,cap(sli4)=6
> 	fmt.Printf("sli5=%v,len(sli5)=%d,cap(sli5)=%d\n", sli5, len(sli5), cap(sli5)) // sli5=[2 3 4],len(sli5)=3,cap(sli5)=6
> 	fmt.Printf("sli6=%v,len(sli6)=%d,cap(sli6)=%d\n", sli6, len(sli6), cap(sli6)) // sli6=[2 3 4 5],len(sli6)=4,cap(sli6)=6
> 
> }
> 
> ```
>
> 



​	可以使用内置函数`make`来创建一个给定元素类型`T`的新的、初始化的切片值，该函数接受一个切片类型和指定长度和可选容量的参数。用`make`创建的切片总是分配一个新的、隐藏的数组，返回的切片值指向该数组。也就是说，执行

```go 
make([]T, length, capacity)
```

产生的切片与分配一个数组并对其进行[切片](../Expressions#slice-expressions-切片表达式)是一样的，所以这两个表达式是等同的：

```go 
make([]int, 50, 100)
new([100]int)[0:50]
```

> 个人注释
>
> ​	make和new函数返回的分别是什么类型？是相同的吗？
>
> => 类型不同！
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	sli1 := make([]int, 6, 6)
> 
> 	sli2 := new([]int)
> 
> 	fmt.Printf("make返回的类型是%T\n", sli1) // make返回的类型是[]int
> 	fmt.Printf("new返回的类型是%T\n", sli2)  // new返回的类型是*[]int
> 
> 	// arr1 := make([6]int, 6) // invalid argument: cannot make [6]int; type must be slice, map, or channel
> 	arr2 := new([6]int)
> 	fmt.Printf("new返回的类型是%T\n", arr2) // new返回的类型是*[6]int
> }
> 
> ```
>
> ​	那为什么说 make([]int, 50, 100)和new([100]int)[0:50] 等同？
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	sli1 := make([]int, 50, 100)
> 
> 	sli2 := new([100]int)[0:50]
> 
> 	fmt.Printf("sli1的类型是%T\n", sli1) // sli1的类型是[]int
> 	fmt.Printf("sli2的类型是%T\n", sli2) // sli2的类型是[]int
> 
> }
> 
> ```
>
> ​	奇怪了，难道是Go的做了什么特殊处理？TODO

​	和数组一样，切片总是一维的，但可以通过组合来构造更高维的对象。对于数组的数组，内部数组在结构上总是相同的长度；但是对于切片的切片（或切片的数组），内部长度可以动态变化。此外，`内部切片必须被单独初始化`。

> 个人注释
>
> ​	以下例子，应该可以解释“`内部切片必须被单独初始化`”：
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	sli1 := make([][]int, 2, 2)
> 	// sli1的类型是[][]int,sli1=[[] []],len(sli1)=2,cap(sli1)=2
> 	fmt.Printf("sli1的类型是%T,sli1=%+v,len(sli1)=%d,cap(sli1)=%d\n", sli1, sli1, len(sli1), cap(sli1))
> 	//sli1[0][0] = 11 // panic: runtime error: index out of range [0] with length 0
> 	sli1[1] = []int{1, 2}
> 	// sli1的类型是[][]int,sli1=[[] [1 2]],len(sli1)=2,cap(sli1)=2
> 	fmt.Printf("sli1的类型是%T,sli1=%+v,len(sli1)=%d,cap(sli1)=%d\n", sli1, sli1, len(sli1), cap(sli1))
> 	sli1[1][0] = 11
> 	// sli1的类型是[][]int,sli1=[[] [11 2]],len(sli1)=2,cap(sli1)=2
> 	fmt.Printf("sli1的类型是%T,sli1=%+v,len(sli1)=%d,cap(sli1)=%d\n", sli1, sli1, len(sli1), cap(sli1))
> 
> 	sli2 := *new([][]int)
> 	// sli2的类型是[][]int,sli2=[],len(sli2)=0,cap(sli2)=0
> 	fmt.Printf("sli2的类型是%T,sli2=%+v,len(sli2)=%d,cap(sli2)=%d\n", sli2, sli2, len(sli2), cap(sli2))
> 
> 	//sli2[0][0] = 11 // panic: runtime error: index out of range [0] with length 0
> 	//sli2[0] = []int{1, 2}     // panic: runtime error: index out of range [0] with length 0                                                                      //
> 	//sli2[1] = []int{1, 2, 3}     // panic: runtime error: index out of range [1] with length 0                                                                   //
> 
> 	var sli3 [][]int
> 	//sli3的类型是[][]int,sli3=[],len(sli3)=0,cap(sli3)=0
> 	fmt.Printf("sli3的类型是%T,sli3=%+v,len(sli3)=%d,cap(sli3)=%d\n", sli3, sli3, len(sli3), cap(sli3))
> 
> 	//sli3[0][0] = 11 // panic: runtime error: index out of range [0] with length 0
> 	//sli3[1] = []int{1, 2, 3} // panic: runtime error: index out of range [1] with length 0
> 
> 	sli4 := make([][2]int, 2, 2)
> 	// sli4的类型是[][2]int,sli4=[[0 0] [0 0]],len(sli4)=2,cap(sli4)=2
> 	fmt.Printf("sli4的类型是%T,sli4=%+v,len(sli4)=%d,cap(sli4)=%d\n", sli4, sli4, len(sli4), cap(sli4))
> 	sli4[0][0] = 11
> 	// sli4的类型是[][2]int,sli4=[[11 0] [0 0]],len(sli4)=2,cap(sli4)=2
> 	fmt.Printf("sli4的类型是%T,sli4=%+v,len(sli4)=%d,cap(sli4)=%d\n", sli4, sli4, len(sli4), cap(sli4))
> 	//sli4[0] = [3]int{1, 2, 3} // cannot use [3]int{…} (value of type [3]int) as [2]int value in assignment
> 	sli4[0] = [2]int{111, 222}
> 	// sli4的类型是[][2]int,sli4=[[111 222] [0 0]],len(sli4)=2,cap(sli4)=2
> 	fmt.Printf("sli4的类型是%T,sli4=%+v,len(sli4)=%d,cap(sli4)=%d\n", sli4, sli4, len(sli4), cap(sli4))
> 
> 	sli5 := *new([][2]int)
> 	// sli5的类型是[][2]int,sli5=[],len(sli5)=0,cap(sli5)=0
> 	fmt.Printf("sli5的类型是%T,sli5=%+v,len(sli5)=%d,cap(sli5)=%d\n", sli5, sli5, len(sli5), cap(sli5))
> 
> 	//sli5[0][0] = 11 // panic: runtime error: index out of range [0] with length 0                                                                          // panic: runtime error: index out of range [0] with length 0
> 	//sli5[0] = [2]int{1, 2} // panic: runtime error: index out of range [0] with length 0
> 
> 	sli6 := [][2]int{{1, 2}, {2, 3}}
> 	// sli6的类型是[][2]int,sli6=[[1 2] [2 3]],len(sli6)=2,cap(sli6)=2
> 	fmt.Printf("sli6的类型是%T,sli6=%+v,len(sli6)=%d,cap(sli6)=%d\n", sli6, sli6, len(sli6), cap(sli6))
> 
> 	sli6[1] = [2]int{22, 33}
> 	// sli6的类型是[][2]int,sli6=[[1 2] [22 33]],len(sli6)=2,cap(sli6)=2
> 	fmt.Printf("sli6的类型是%T,sli6=%+v,len(sli6)=%d,cap(sli6)=%d\n", sli6, sli6, len(sli6), cap(sli6))
> 
> 	sli6[1] = [...]int{222, 333}
> 	// sli6的类型是[][2]int,sli6=[[1 2] [222 333]],len(sli6)=2,cap(sli6)=2
> 	fmt.Printf("sli6的类型是%T,sli6=%+v,len(sli6)=%d,cap(sli6)=%d\n", sli6, sli6, len(sli6), cap(sli6))
> 
> 	sli7 := [][]int{{1}, {2, 3}, {4, 5, 6}}
> 	// sli7的类型是[][]int,sli7=[[1] [2 3] [4 5 6]],len(sli7)=3,cap(sli7)=3
> 	fmt.Printf("sli7的类型是%T,sli7=%+v,len(sli7)=%d,cap(sli7)=%d\n", sli7, sli7, len(sli7), cap(sli7))
> 	sli7[0] = []int{1, 11, 111}
> 	// sli7的类型是[][]int,sli7=[[1 11 111] [2 3] [4 5 6]],len(sli7)=3,cap(sli7)=3
> 	fmt.Printf("sli7的类型是%T,sli7=%+v,len(sli7)=%d,cap(sli7)=%d\n", sli7, sli7, len(sli7), cap(sli7))
> 	sli7[0][2] = 1111
> 	// sli7的类型是[][]int,sli7=[[1 11 1111] [2 3] [4 5 6]],len(sli7)=3,cap(sli7)=3
> 	fmt.Printf("sli7的类型是%T,sli7=%+v,len(sli7)=%d,cap(sli7)=%d\n", sli7, sli7, len(sli7), cap(sli7))
> 	sli7[0][3] = 11111 // panic: runtime error: index out of range [3] with length 3
> 
> }
> 
> ```
>
> 

### Struct types 结构体型

​	结构体是一系列具有名称和类型的命名元素，称为`字段`。字段名可以显式指定（IdentifierList）或隐式指定（EmbeddedField）。在结构体内部，非[空白](../DeclarationsAndScope#blank-identifier-空白标识符)字段名必须是[唯一](../DeclarationsAndScope#uniqueness-of-identifiers-标识符的唯一性)的。

```
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName [ TypeArgs ] .
Tag           = string_lit .
```


```go
// 一个空结构体
struct {}

// 一个带有6个字段的结构体
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

​	使用类型但没有显式字段名声明的字段被称为`嵌入字段`。嵌入字段必须被指定为一个类型名`T`或一个指向非接口类型名`*T`的指针，而且`T`本身不能是一个指针类型。未限定类型名作为字段名。

```go 
// A struct with four embedded fields of types T1, *T2, P.T3 and *P.T4
struct {
	T1        // field name is T1
	*T2       // field name is T2
	P.T3      // field name is T3
	*P.T4     // field name is T4
	x, y int  // field names are x and y
}
```

> 个人注释
>
> ​	什么是“未限定类型名作为字段名”？结构体变量可以使用for range中多为被迭代对象吗？

{{< tabpane text=true >}}

{{< tab header="main.go" >}}

```go
package main

import (
	"example.com/101/dftype"
	"fmt"
)

type MySt1 struct {
	dftype.T1
	dftype.T2
	x, y int
}

func main() {
	st1 := MySt1{1, 2, 3, 4}
	fmt.Printf("st1的类型是%T,st=%+v\n", st1, st1)
	fmt.Printf("T1字段的值%v\n", st1.T1) //T1字段的值1
	fmt.Printf("T2字段的值%v\n", st1.T2) //T2字段的值2
	fmt.Printf("x字段的值%v\n", st1.x)   //x字段的值3
	fmt.Printf("y字段的值%v\n", st1.y)   //y字段的值4

	// implicit assignment to unexported field x in struct literal of type dftype.MySt1
	// implicit assignment to unexported field y in struct literal of type dftype.MySt1
	// st2 := dftype.MySt1{1, 2, 3, 4}
	// fmt.Printf("st2的类型是%T,st2=%+v\n", st2, st2)
	st2 := dftype.MySt1{T1: 1, T2: 2}
	fmt.Printf("st2的类型是%T,st2=%+v\n", st2, st2) //st2的类型是dftype.MySt1,st2={T1:1 T2:2 x:0 y:0}
	fmt.Printf("T1字段的值%v\n", st2.T1)            //T1字段的值1
	fmt.Printf("T2字段的值%v\n", st2.T2)            //T2字段的值2

	st3 := dftype.MySt2{1, 2, 3, 4}
	fmt.Printf("st3的类型是%T,st3=%+v\n", st3, st3) // st3的类型是dftype.MySt2,st3={T1:1 T2:2 X:3 Y:4}
	fmt.Printf("T1字段的值%v\n", st3.T1)            //T1字段的值1
	fmt.Printf("T2字段的值%v\n", st3.T2)            //T2字段的值2
	fmt.Printf("X字段的值%v\n", st3.X)              //X字段的值3
	fmt.Printf("Y字段的值%v\n", st3.Y)              //Y字段的值4

	// cannot range over st3 (variable of type dftype.MySt2)
	//for i, v := range st3 {
	//	fmt.Println(i, ":", v, "\n")
	//}
}

```

{{< /tab >}}

{{< tab header="dftype.go" >}}

```go
package dftype

type T1 int
type T2 int8

type MySt1 struct {
	T1
	T2
	x, y int
}

type MySt2 struct {
	T1
	T2
	X, Y int
}

```

{{< /tab >}}

{{< /tabpane >}}	

> 个人注释
>
> ​	相信以上示例，已经给出了答案：未限定类型名作为字段名，即是将类型名直接作为字段名；结构体变量不能用于 for range语句中作为被迭代对象。

​	下面的声明是非法的，`因为字段名在一个结构体类型中必须是唯一的`。

```go 
struct {
	T     // conflicts with embedded field *T and *P.T
	*T    // conflicts with embedded field T and *P.T
	*P.T  // conflicts with embedded field T and *T
}
```

> ​	个人注释
>
> ​	若嵌入字段是非接口类型名`*T`的指针，那结构体变量使用%v、%+v或%#v打印时，字段名是什么？输出是怎样的？
>
> ```go
> package main
> 
> import "fmt"
> 
> type T1 int
> type T2 int8
> 
> type MySt1 struct {
> 	*T1
> 	T2
> }
> 
> func main() {
> 	t := T1(1)
> 	st := MySt1{&t, 2}
> 	fmt.Printf("st=%v\n", st)  // st={0xc00001a0a8 2}
> 	fmt.Printf("st=%+v\n", st) // st={T1:0xc00001a0a8 T2:2}
> 	fmt.Printf("st=%#v\n", st) // st=main.MySt1{T1:(*main.T1)(0xc00001a0a8), T2:2}
> }
> 
> ```
>
> 

​	如果`x.f`是表示字段或[方法](../DeclarationsAndScope#function-declarations-方法声明)`f`的合法[选择器](../Expressions#selectors-选择器)，那么结构体`x`中的嵌入字段或方法`f`被称为（自动）提升（的字段或方法）。

> 个人注释
>
> ​	请给出提升的字段或方法的示例。
>
> ```go
> package main
> 
> import "fmt"
> 
> type MySt1 struct {
> 	Size int
> }
> 
> // 报错：field and method with the same name Size
> //func (mySt1 MySt1) Size() int {
> //	return mySt1.Size
> //}
> 
> func (mySt1 MySt1) Size1() int {
> 	return mySt1.Size
> }
> 
> type MySt2 struct {
> 	MySt1
> 	Name string
> 	Age  int
> }
> 
> func main() {
> 	st := MySt2{MySt1: MySt1{Size: 20}, Name: "zlongx", Age: 32}
> 	fmt.Printf("st=%v\n", st)  // st={{20} zlongx 32}
> 	fmt.Printf("st=%+v\n", st) // st={MySt1:{Size:20} Name:zlongx Age:32}
> 	fmt.Printf("st=%#v\n", st) // st=main.MySt2{MySt1:main.MySt1{Size:20}, Name:"zlongx", Age:32}
> 
> 	// 被提升的字段 Size
> 	fmt.Printf("Size=%d\n", st.Size) // Size=20
> 
> 	// 被提升的方法 Size1()
> 	fmt.Printf("Size=%d\n", st.Size1()) // Size=20
> }
> 
> ```
>
> 

​	被提升的字段与结构体中的普通字段一样，只是它们不能在结构体的[复合字面量](../Expressions#composite-literals-复合字面量)中作为字段名使用。

> 个人注释
>
> ​	解释下“被提升的字段与结构体中的普通字段一样，只是它们不能在结构体的[复合字面量](../Expressions#composite-literals-复合字面量)中作为字段名使用。”？
>
> ```go
> package main
> 
> import "fmt"
> 
> type MySt1 struct {
> 	Size int
> }
> 
> // 报错：field and method with the same name Size
> //func (mySt1 MySt1) Size() int {
> //	return mySt1.Size
> //}
> 
> func (mySt1 MySt1) Size1() int {
> 	return mySt1.Size
> }
> 
> type MySt2 struct {
> 	MySt1
> 	Name string
> 	Age  int
> }
> 
> func main() {
> 	// 报错：unknown field Size in struct literal of type MySt2
> 	// st := MySt2{Size: 20, Name: "zlongx", Age: 32}
> 	st1 := MySt2{MySt1: MySt1{Size: 20}, Name: "zlongx", Age: 32}
> 	fmt.Printf("st1=%#v\n", st1) // st1=main.MySt2{MySt1:main.MySt1{Size:20}, Name:"zlongx", Age:32}
> 	st2 := MySt2{MySt1{Size: 20}, "zlongx", 32}
> 	fmt.Printf("st2=%#v\n", st2) // st2=main.MySt2{MySt1:main.MySt1{Size:20}, Name:"zlongx", Age:32}
> 
> }
> ```
>
> ​	在创建结构体 `MySt2` 的实例时，我们可以使用复合字面量来初始化普通字段 `Name`和`Age`，但无法直接使用复合字面量来初始化提升字段 `Size`。我们需要通过指定嵌入字段 `MySt1` 的普通字段 `Size` 来初始化它。
>
> ​	即Size这个字段不能直接在结构体字面量的最外围花括号中直接使用。

​	给定一个结构体类型`S`和一个[命名类型](../Types)`T`，提升的方法按以下方式包含在结构体的方法集中：

(aa）如果`S`包含一个嵌入式字段`T`，那么`S`和`*S`的方法集都包括带有接收器`T`的提升方法，`*S`的方法集也包括带有接收器`*T`的提升方法。

(bb) 如果`S`包含一个嵌入式字段`*T`，那么`S`和`*S`的方法集都包括带有接收器`T`或`*T`的提升方法。

> ​	个人注释
>
> ​	如下示例中的 a、b、d 可以解释(aa)这一点。
>
> 但示例中的c，又该如何解释呢？这应该是go语言编译器的特殊处理或者称为go语法糖！TODO 待找出出处！以及给出编译后汇编代码？
>
> ​	如下示例中的 e、f、g、h 可以解释(bb)这一点。
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"reflect"
> )
> 
> func DumpMethodSet(i interface{}) {
> 	v := reflect.TypeOf(i)
> 	elemTyp := v.Elem()
> 
> 	n := elemTyp.NumMethod()
> 	if n == 0 {
> 		fmt.Printf("%s's method set is empty!\n", elemTyp)
> 	}
> 
> 	fmt.Printf("%s's method set:\n", elemTyp)
> 	for j := 0; j < n; j++ {
> 		fmt.Println("-", elemTyp.Method(j).Name)
> 	}
> 	fmt.Println()
> }
> 
> type T1 struct {
> }
> 
> func (t T1) Method1A() {
> 	fmt.Println("Method1A called")
> }
> 
> func (t *T1) Method1B() {
> 	fmt.Println("Method1B called")
> }
> 
> type T2 struct {
> }
> 
> func (t T2) Method2A() {
> 	fmt.Println("Method2A called")
> }
> 
> func (t *T2) Method2B() {
> 	fmt.Println("Method2B called")
> }
> 
> type S struct {
> 	T1
> 	*T2
> }
> 
> func main() {
> 	s := S{T1: T1{}, T2: &T2{}}
> 
> 	//a. 可以直接调用 接收器为T1 的提升方法 Method1A
> 	s.Method1A() // Method1A called
> 	//b. 可以通过指针调用 接收器为T1 的提升方法 Method1A
> 	(&s).Method1A() // Method1A called
> 
> 	//c. 可以直接调用 接收器为*T1 的提升方法 Method1B
> 	s.Method1B() // Method1B called
> 	//d. 可以通过指针调用 接收器为*T1 的提升方法 Method1B
> 	(&s).Method1B() // Method1B called
> 
> 	//e. 可以直接调用 接收器为T2 的提升方法 Method2A
> 	s.Method2A() // Method2A called
> 	//f. 可以通过指针调用 接收器为T2 的提升方法 Method2A
> 	(&s).Method2A() // Method2A called
> 
> 	//g. 可以直接调用 接收器为*T2 的提升方法 Method2B
> 	s.Method2B() // Method2B called
> 	//h. 可以通过指针调用 接收器为*T2 的提升方法 Method2B
> 	(&s).Method2B() // Method2B called
> 
> 	var t = s
> 	var pt = &s
> 	//main.S's method set:
> 	//- Method1A
> 	//- Method2A
> 	//- Method2B
> 	DumpMethodSet(&t)
> 	//*main.S's method set:
> 	//- Method1A
> 	//- Method1B
> 	//- Method2A
> 	//- Method2B
> 	DumpMethodSet(&pt)
> 	fmt.Printf("%T\n", pt)  // *main.S
> 	fmt.Printf("%T\n", &pt) // **main.S
> 
> 
> 	var tt S
> 	var ppt *S
> 	//main.S's method set:
> 	//- Method1A
> 	//- Method2A
> 	//- Method2B
> 	DumpMethodSet(&tt)
> 	//*main.S's method set:
> 	//- Method1A
> 	//- Method1B
> 	//- Method2A
> 	//- Method2B
> 	DumpMethodSet(&ppt)
> }
> 
> ```
>
> 

​	一个字段声明后面可以有一个可选的`字符串字面量标签`，它成为相应字段声明中所有字段的属性。一个空的标签字符串等同于一个不存在标签。标签通过[反射接口](https://pkg.go.dev/reflect#StructTag)可见，并参与结构体的[类型标识](../PropertiesOfTypesAndValues#type-identity-类型一致性)，但在其他情况下被忽略。

```go 
struct {
	x, y float64 ""  // an empty tag string is like an absent tag
	name string  "any string is permitted as a tag"
	_    [4]byte "ceci n'est pas un champ de structure"
}

// A struct corresponding to a TimeStamp protocol buffer. 
// 一个对应于 TimeStamp 协议缓冲区的结构。
// The tag strings define the protocol buffer field numbers;
// 标签字符串定义协议缓冲区字段编号
// they follow the convention outlined by the reflect package.
// 它们遵循反射包描述的约定。
struct {
	microsec  uint64 `protobuf:"1"`
	serverIP6 uint64 `protobuf:"2"`
}
```

> 个人注释
>
> ​	请给出怎么通过反射接口给出标签的示例？
>
> ```go
> 
> ```
>
> 

### Pointer types 指针型

​	指针类型表示指向给定类型（称为指针的基本类型）[变量](../Variables)的所有指针的集合。一个未初始化的指针的值是`nil`。

```go 
PointerType = "*" BaseType .
BaseType    = Type .
*Point
*[4]int
```

### Function types 函数型

​	函数类型表示具有相同参数类型和结果类型的所有函数的集合。一个函数类型的未初始化变量的值是`nil`。

``` go
functionType   = "func" Signature .
Signature      = Parameters [ Result ] .
Result         = Parameters | Type .
Parameters     = "(" [ ParameterList [ "," ] ] ")" .
ParameterList  = ParameterDecl { "," ParameterDecl } .
ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
```

​	在参数或结果的列表中，名称（IdentifierList）必须全部存在或全部不存在。如果存在（名称），每个名称代表指定类型的一个项（参数或结果），并且签名中所有非[空白](../DeclarationsAndScope#blank-identifier-空白标识符)的名称必须是唯一的。如果不存在（名称），每个类型代表该类型的一个项。参数和结果列表总是用括号表示，但如果正好仅有一个未命名的结果，则可以写成未括号的类型。

​	在函数签名中的最后一个传入参数可以有一个`类型前缀` `...`。有这样一个参数的函数被称为 variadic （`可变参数函数`），可以用零个或多个参数来调用该函数。

```go 
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

### Interface types 接口型

​	接口类型定义了一个类型集。接口类型的变量可以存储该接口类型集中的任何类型的值。这样的类型被称为[实现了该接口](#implementing-an-interface-实现一个接口)。未初始化的接口类型变量的值是`nil`。

```
InterfaceType  = "interface" "{" { InterfaceElem ";" } "}" .
InterfaceElem  = MethodElem | TypeElem .
MethodElem     = MethodName Signature .
MethodName     = identifier .
TypeElem       = TypeTerm { "|" TypeTerm } .
TypeTerm       = Type | UnderlyingType .
UnderlyingType = "~" Type .
```

​	接口类型由接口元素列表指定。接口元素是一个方法或一个类型元素，其中类型元素是一个或多个类型项的联合。类型项可以是一个单一类型，也可以是一个单一的底层类型。

#### Basic interfaces 基本接口

​	在其最基本的形式中，接口指定了一个（可能是空的）方法列表。由这样一个接口定义的类型集是实现了所有这些方法的类型集，而相应的方法集则完全由这个接口指定的方法组成。那些类型集可以`完全由一个方法列表`来定义的接口被称为`基本接口`。

```go 
// A simple File interface.
interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}
```

每个显式指定的方法的名称必须是[唯一](../DeclarationsAndScope#uniqueness-of-identifiers)的，不能是[空白](../PropertiesOfTypesAndValues#blank-identity)。

```go 
interface {
	String() string
	String() string  // illegal: String not unique
	_(x int)         // illegal: method must have non-blank name
}
```

​	多个类型可以实现一个（相同的）接口。例如，如果两个类型`S1`和`S2`的方法设置为

```go 
func (p T) Read(p []byte) (n int, err error)
func (p T) Write(p []byte) (n int, err error)
func (p T) Close() error
```

(其中`T`代表`S1`或`S2`），那么`File`接口就由`S1`和`S2`实现，而不管`S1`和`S2`可能有其他方法或共享什么其他方法。

​	作为接口类型集成员的每个类型都实现了该接口。任何给定的类型都可以实现几个不同的接口。例如，所有类型都实现`空接口` （interface {}），它代表所有（非接口）类型的集合：

```go 
interface{}
```

为了方便，预先声明的类型`any`是`空接口的别名`。

​	类似地，考虑这个接口规范，它出现在定义名为 `Locker` 的接口的类型声明中：

```go 
type Locker interface {
	Lock()
	Unlock()
}
```

如果`S1`和`S2`也实现了

```go 
func (p T) Lock() { … }
func (p T) Unlock() { … }
```

他们就实现了`Locker`接口和`File`接口。

#### Embedded interfaces 嵌入接口

​	接口`T`可以使用（可能是限定的）接口类型名称`E`作为接口元素。这就是在 `T` 中嵌入接口 `E`。`T`的类型集是由`T`的显式声明方法定义的类型集和`T`的嵌入接口的类型集的`交集`。换句话说，`T`的类型集是实现`T`的所有显式声明的方法以及`E`的所有方法的所有类型的集合。

```go 
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter's methods are Read, Write, and Close.
type ReadWriter interface {
	Reader  // includes methods of Reader in ReadWriter's method set
	Writer  // includes methods of Writer in ReadWriter's method set
}
```

在嵌入接口时，具有[相同](../DeclarationsAndScope#uniqueness-of-identifiers)名称的方法必须具有[相同](../PropertiesOfTypesAndValues#type-identity)的签名。

```go 
type ReadCloser interface {
	Reader   // includes methods of Reader in ReadCloser's method set
	Close()  // illegal: signatures of Reader.Close and Close are different
}
```

#### General interfaces 通用接口

​	在最通用的形式下，接口元素也可以是一个任意类型项`T`，或者是一个指定底层类型`T`的`~T`形式的项，或者是一系列项`t1|t2|...|tn`的联合。结合方法规范，这些元素能够精确地定义一个接口的类型集，如下所示：

- 空接口的类型集是`所有非接口类型的集合`。
- 非空接口的类型集是其接口元素的类型集的交集。
- 方法规范的类型集是包含该方法的`所有非接口类型的集合`。
- 非接口类型项的类型集是仅由该类型组成的集合。
- 形式为`~T`的项的类型集是底层类型为`T`的所有类型的集合。
- 一系列项`t1|t2|...|tn`的类型集是这些项的类型集的并集。

​	量化 "`所有非接口类型的集合` "不仅指当前程序中声明的所有（非接口）类型，还指所有可能程序中的所有可能类型，因此是无限的。类似地，给定实现某个特定方法的`所有非接口类型的集合`，这些类型的方法集的`交集`将正好包含该方法，即使当前程序中的所有类型总是将该方法与另一个方法配对。

​	根据定义，`一个接口的类型集永远不会包含一个接口类型`。

```go 
// An interface representing only the type int.
// 仅表示 int 类型的接口
interface {
	int
}

// An interface representing all types with underlying type int.
// 表示底层类型为 int 的所有类型的接口
interface {
	~int
}

// An interface representing all types with underlying type int that implement the String method.
// 表示具有实现 String() 方法 和 底层类型为 int 的所有类型的接口。
interface {
	~int
	String() string
}

// An interface representing an empty type set: there is no type that is both an int and a string.
// 表示空类型集的接口: 没有既是 int 又是 string 的类型。
interface {
	int
	string
}
```

​	在形式为`~T`的项中，`T`的底层类型必须是它自己，而且`T`不能是一个接口。

```go 
type MyInt int

interface {
	~[]byte  // the underlying type of []byte is itself => []byte 的底层类型是其本身
	~MyInt   // illegal: the underlying type of MyInt is not MyInt => 非法的: MyInt的底层类型不是MyInt
	~error   // illegal: error is an interface => 非法的: error 是一个接口
}
```

联合元素表示类型集的并集：

```go 
// The Float interface represents all floating-point types
// (including any named types whose underlying types are
// either float32 or float64).
// Float接口表示所有的浮点类型（包括底层类型为float32或float64的命名类型）。
type Float interface {
	~float32 | ~float64
}
```

​	形式为`T`或`~T`的项中的类型`T`不能是[类型参数](../DeclarationsAndScope#type-parameter-declarations)，所有非接口项的类型集必须是成对不相交的（类型集的成对交集必须为空）。给定一个类型参数P：

```go 
interface {
	P                // illegal: P is a type parameter => 非法的: P 是一个类型参数
	int | ~P         // illegal: P is a type parameter => 非法的: P 是一个类型参数
	~int | MyInt     // illegal: the type sets for ~int and MyInt are not disjoint (~int includes MyInt)  => 非法的: ~int 和 MyInt 的类型集是相交的(~int 包括 MyInt)
	float32 | Float  // overlapping type sets but Float is an interface  => 重叠的类型集，更进一步说Float也是一个接口
}
```

实现限制：一个联合(有多个项)不能包含[预先声明的标识符中的](../DeclarationsAndScope#predeclared-identifiers--预先声明的标识符)`comparable`或指定了方法的接口，或嵌入`comparable`或指定了方法的接口。

​	非[基本接口](#basic-interfaces-基本接口)只能作为类型约束使用，或者作为其他接口的元素作为约束使用。它们不能作为值或变量的类型，**也不能作为其他非接口类型的组成部分**。

```go 
var x Float                     // illegal: Float is not a basic interface => 非法: Float 不是一个基本接口

var x interface{} = Float(nil)  // illegal => 非法

type Floatish struct {
	f Float                 // illegal => 非法
}
```

接口类型 `T` 不能嵌入任何递归地包含或嵌入 `T` 的类型元素。

```go 
// illegal: Bad cannot embed itself => 非法: Bad 不能嵌入自己
type Bad interface {
	Bad
}

// illegal: Bad1 cannot embed itself using Bad2 => 非法: Bad1不能使用 Bad2嵌入自己
type Bad1 interface {
	Bad2
}
type Bad2 interface {
	Bad1
}

// illegal: Bad3 cannot embed a union containing Bad3 => Bad3 不能嵌入包含 Bad3的联合
type Bad3 interface {
	~int | ~string | Bad3
}
```

#### Implementing an interface 实现一个接口

如果类型`T`实现了接口`I`，则

​	(a)`T`不是接口，并且是`I`类型集的元素；或者

​	(b) `T`是接口，并且`T`的类型集是`I`的类型集的子集。

如果`T`实现了一个接口，那么`T`类型的值就实现了该接口。

> 个人注释
>
> ​	针对(b)给出示例说明下：
>
> ```go
> package main
> 
> import (
> 	"fmt"
> )
> 
> type Reader interface {
> 	Read()
> }
> 
> type Writer interface {
> 	Write(data string)
> }
> 
> type ReadWriter interface {
> 	Read()
> 	Write(data string)
> }
> 
> type Document struct{}
> 
> func (d Document) Read() {
> 	fmt.Println("Reading document")
> }
> 
> func (d Document) Write(data string) {
> 	fmt.Println("Writing \"" + data + "\" to document")
> }
> 
> func main() {
> 	var r Reader
> 	var w Writer
> 	var rw ReadWriter
> 	doc := Document{}
> 	doc.Read()         // Reading document
> 	doc.Write("hello") //Writing "hello" to document
> 
> 	r = doc
> 	w = doc
> 	rw = doc
> 
> 	r.Read()         // Reading document
> 	w.Write("world") // Writing "world" to document
> 	rw.Read()        // Reading document
> 	rw.Write("hi")   // Writing "hi" to document
> }
> 
> ```
>
> 

### Map types 映射型

​	映射是一个无序的元素组，由一种类型的元素（称为`元素类型`）组成，由另一种类型的唯一键集（称为`键类型`）进行索引。一个未初始化的映射的值是`nil`。

```
MapType     = "map" "[" KeyType "]" ElementType .
KeyType     = Type .
```

​	[比较运算符](../Expressions#comparison-operators-比较运算符)`==`和`!=`必须为键类型的操作数完全定义；`因此键类型不能是函数、映射或切片`。如果键类型是接口类型，则必须为动态键值定义这些比较运算符；失败将导致[运行时恐慌（run-time panic）](../Run-timePanics)。

> 个人注释
>
> ​	map的键可以是数组吗？=> 可以
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	var m0 map[[3]int]string
> 	fmt.Printf("m0=%v\n", m0)  // m0=map[]
> 	fmt.Printf("m0=%+v\n", m0) // m0=map[]
> 	fmt.Printf("m0=%#v\n", m0) // m0=map[[3]int]string(nil)
> 	//m0[[...]int{0, 1, 2}] = "A0" //报错： panic: assignment to entry in nil map
> 	//m0[[...]int{2, 3, 4}] = "B0" //报错： panic: assignment to entry in nil map
> 
> 	m1 := map[[3]int]string{}
> 	fmt.Printf("m1=%v\n", m1)  // m1=map[]
> 	fmt.Printf("m1=%+v\n", m1) // m1=map[]
> 	fmt.Printf("m1=%#v\n", m1) // m1=map[[3]int]string{}
> 
> 	m1[[...]int{0, 1, 2}] = "A1"
> 	m1[[...]int{2, 3, 4}] = "B1"
> 	fmt.Printf("m1=%v\n", m1)  // m1=map[[0 1 2]:A1 [2 3 4]:B1]
> 	fmt.Printf("m1=%+v\n", m1) // m1=map[[0 1 2]:A1 [2 3 4]:B1]
> 	fmt.Printf("m1=%#v\n", m1) // m1=map[[3]int]string{[3]int{0, 1, 2}:"A1", [3]int{2, 3, 4}:"B1"}
> 
> 	m2 := map[[3]int]string{[...]int{0, 1, 2}: "A2", [...]int{2, 3, 4}: "B2"}
> 	fmt.Printf("m2=%v\n", m2)  // m2=map[[0 1 2]:A2 [2 3 4]:B2]
> 	fmt.Printf("m2=%+v\n", m2) // m2=map[[0 1 2]:A2 [2 3 4]:B2]
> 	fmt.Printf("m2=%#v\n", m2) // m2=map[[3]int]string{[3]int{0, 1, 2}:"A2", [3]int{2, 3, 4}:"B2"}
> 
> 	m3 := make(map[[3]int]string)
> 	fmt.Printf("m3=%v\n", m3)  // m3=map[]
> 	fmt.Printf("m3=%+v\n", m3) // m3=map[]
> 	fmt.Printf("m3=%#v\n", m3) // m3=map[[3]int]string{}
> 	m3[[...]int{0, 1, 2}] = "A3"
> 	m3[[...]int{2, 3, 4}] = "B3"
> 	fmt.Printf("m3=%v\n", m3)  // m3=map[[0 1 2]:A3 [2 3 4]:B3]
> 	fmt.Printf("m3=%+v\n", m3) // m3=map[[0 1 2]:A3 [2 3 4]:B3]
> 	fmt.Printf("m3=%#v\n", m3) // m3=map[[3]int]string{[3]int{0, 1, 2}:"A3", [3]int{2, 3, 4}:"B3"}
> 
> 	m4 := *new(map[[3]int]string)
> 	fmt.Printf("m4=%v\n", m4)  // m4=map[]
> 	fmt.Printf("m4=%+v\n", m4) // m4=map[]
> 	fmt.Printf("m4=%#v\n", m4) // m4=map[[3]int]string(nil)
> 	//m4[[...]int{0, 1, 2}] = "A3" //报错 panic: assignment to entry in nil map
> }
> 
> ```
>
> ​	若map的键是接口类型，怎么为该动态键值定义比较运算符？TODO

```go 
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

​	映射元素的数量被称为它的`长度`。对于一个map `m`来说，它可以用内置函数`len`来发现，并且在运行过程中可能会改变。在运行过程中可以用[赋值](../Statements#assignment-statements)添加元素，用[索引表达式](../Expressions##index-expressions-索引表达式)检索元素；可以用内置函数`delete`删除元素。

​	使用内置函数 `make` 创建一个新的空 map 值，它使用 map 类型和一个可选的容量提示作为参数：

```go 
make(map[string]int)
make(map[string]int, 100)
```

​	初始容量不限制其大小：映射会增长以容纳其中存储的项数，但`nil`映射除外。`nil`映射等同于空映射，`只是不能添加任何元素`。

### Channel types 通道型

​	通道为[并发执行函数](../Statements#go-statements----go-语句)提供了一种机制，通过[发送](../Statements#send-statements-发送语句)和[接收](../Expressions#receive-operator-接收操作符)指定元素类型的值进行通信。未初始化的通道的值是`nil`。

```
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .
```

​	可选的`<-`操作符指定了通道的方向：发送或接收。如果指定了方向，则该通道是定向的，否则是双向的。通过[赋值](../Statements#assignment-statements-赋值语句)或显式[转换](../Expressions#conversions-转换)，通道可以被限制为仅发送或仅接收。

```go 
chan T          // can be used to send and receive values of type T => 可用于发送或接收类型为 T 的值
chan<- float64  // can only be used to send float64s => 仅用于发送 float64 类型
<-chan int      // can only be used to receive ints => 仅用于接收 int 类型
```

`<-` 操作符尽可能与最左边的 `chan` 相关联：

```go 
chan<- chan int    // same as chan<- (chan int) => 与 chan<- (chan int) 相同
chan<- <-chan int  // same as chan<- (<-chan int) =>与 chan<- (<-chan int) 相同
<-chan <-chan int  // same as <-chan (<-chan int) => 与 <-chan (<-chan int) 相同
chan (<-chan int)
```

​	可以使用内置函数 `make` 创建一个新的、初始化的 channel 值，它以channel 类型和可选的容量作为参数：

```go 
make(chan int, 100)
```

​	容量(以元素数量为单位)设置通道中缓冲区的大小。如果容量为零或没有指定，则通道是无缓冲的，只有当发送方和接收方都准备好时，通信才会成功。否则，如果缓冲区不满(可继续发送)或不是空的(可继续接收) ，通道会将数据缓冲起来，并且通信在没有阻塞的情况下成功。一个`nil`通道不能用于通信。

> 引用其他书籍
>
> ​    以下摘自《Go语言精进之路》第34条 了解channel的妙用 第348页。
> ​    与无缓冲channel 不同，带缓冲channel 可以通过带有 capacity 参数的内置make 函数创建：c:= make(chan  T, capctity)
> ​    由于带缓冲channel 的运行时层实现带有缓冲区，因此对带有缓冲channel的发送操作在缓冲区未满、接收操作在缓冲区非空的情况下是异步的（发送或接收无需阻塞等待）。也就是说，对一个带缓冲channel，在缓冲区无数据或有数据但未满的情况下，对其进行发送操作的goroutine不会阻塞；在缓冲区已满的情况下，对其进行发送操作的goroutine会阻塞；在缓冲区为空的情况下，对其进行接收操作的goroutine亦会阻塞。



​	通道可以用内置函数`close`来关闭。[接收操作符](../Expressions#receive-operator-接收操作符)的多值赋值形式可以用来判断数据是否在通道关闭之前发送出去。

> ​	个人注释
>
> ​	什么是“[接收操作符](../Expressions#receive-operator-接收操作符)的多值赋值”？
>
> ```go
> value, ok := <-ch
> ```
>
> ​	

​	任意数量的goroutines都可以通过[发送语句](../Statements#send-statements-发送语句)、[接收操作](../Expressions#receive-operator-接收操作符)以及对内置函数`cap`和`len`的调用，来操作一个通道。通道是一个先入先出的队列。例如，如果一个goroutine在通道上发送数据，第二个goroutine接收这些数据，那么这些数据将按照发送的顺序被接收。