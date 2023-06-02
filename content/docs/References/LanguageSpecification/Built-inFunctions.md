+++
title = "内置函数"
date = 2023-05-17T09:59:21+08:00
weight = 13
description = ""
isCJKLanguage = true
draft = false
+++
## Built-in functions 内置函数

> 原文：[https://go.dev/ref/spec#Built-in_functions](https://go.dev/ref/spec#Built-in_functions)

​	内置函数是[预先声明的](../DeclarationsAndScope#predeclared-identifiers--预先声明的标识符)。它们像其他函数一样被调用，但其中一些函数接受一个类型而非表达式作为其第一个实参。

​	内置函数没有标准的Go类型，所以它们只能出现在[调用表达式](../Expressions#calls-调用)中；它们不能作为函数值使用。

### Close 

​	对于一个[核心类型](../PropertiesOfTypesAndValues#core-typess-核心类型)为[通道](../Types#channel-types-通道型)的参数`ch`，内置函数`close`记录了通道上将不再有任何值被发送。如果`ch`是一个仅接收的通道，那么（关闭它）是一个错误。发送到或关闭一个已关闭的通道会导致[运行时恐慌](../Run-timePanics)。关闭`nil`通道也会引起[运行时恐慌](../Run-timePanics)。在调用`close`后，并且在任何先前发送的值被接收后，接收操作将返回通道类型的`零值`而不阻塞。多值[接收操作](../Expressions#receive-operator-接收操作符)会返回一个接收值以及通道是否被关闭的指示。

### Length and capacity 长度和容量

​	内置函数`len`和`cap`接受各种类型的实参并返回`int`类型的结果。该实现保证结果总是适合于一个`int`。

```
Call      Argument type    Result
调用        实参类型          结果

len(s)    string type      string length in bytes
          [n]T, *[n]T      array length (== n)
          []T              slice length
          map[K]T          map length (number of defined keys)
          chan T           number of elements queued in channel buffer
          type parameter   see below

cap(s)    [n]T, *[n]T      array length (== n)
          []T              slice capacity
          chan T           channel buffer capacity
          type parameter   see below
```

​	如果参数类型是一个[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)`P`，调用`len(e)`（或`cap(e)`）必须对`P`的类型集中的每个类型有效。其结果是（类型对应`P`被[实例化](../Expressions#instantiations-实例化)时使用的类型实参的）实参的长度（或容量）。

​	切片的容量是底层数组中分配到的元素的数量。在任何时候，以下关系都是成立的：

```
0 <= len(s) <= cap(s)
```

​	`nil`切片、`nil`映射或`nil`通道的长度是`0`。`nil`切片或`nil`通道的容量是0。

​	如果`s`是一个字符串[常量](../Constants)，那么表达式`len(s)`就是常量。如果`s`的类型是一个数组或指向数组的指针，并且表达式`s`不包含[通道接收](../Expressions#receive-operator-接收操作符)或（非常量）[函数调用](../Expressions#calls-调用)，那么表达式`len(s)`和`cap(s)`是常量；在这种情况下，`s`不被求值。否则，`len`和`cap`的调用不是常量，`s`被求值。

```go 
const (
	c1 = imag(2i)                    // imag(2i) = 2.0 is a constant
	c2 = len([10]float64{2})         // [10]float64{2} contains no function calls
	c3 = len([10]float64{c1})        // [10]float64{c1} contains no function calls
	c4 = len([10]float64{imag(2i)})  // imag(2i) is a constant and no function call is issued
	c5 = len([10]float64{imag(z)})   // invalid: imag(z) is a (non-constant) function call
)
var z complex128
```

### Allocation 分配

​	内置函数`new`接收一个类型`T`，在运行时为该类型的[变量](../Variables)分配存储空间，并返回一个[指向](../Types#pointer-types-指针型)它的`*T`类型的值。该变量被初始化，如[初始值](../ProgramInitializationAndExecution#the-zero-value-零值)一节中所述。

```go 
new(T)
```

举例来说

```go 
type S struct { a int; b float64 }
new(S)
```

为一个`S`类型的变量分配存储空间，初始化它（`a=0`， `b=0.0`），并返回一个包含该位置地址的`*S`类型的值。

### Making slices, maps and channels 制作切片、映射和通道

​	内置函数`make`接收一个类型`T`，后面可以选择一个特定类型的表达式列表。`T`的[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)`必须是一个切片、映射或通道`。它返回一个类型为`T`（不是`*T`）的值。内存被初始化，如[初始值](../ProgramInitializationAndExecution#the-zero-value-零值)一节中所述。

```
Call             Core type    Result

make(T, n)       slice        slice of type T with length n and capacity n
make(T, n, m)    slice        slice of type T with length n and capacity m

make(T)          map          map of type T
make(T, n)       map          map of type T with initial space for approximately n elements

make(T)          channel      unbuffered channel of type T
make(T, n)       channel      buffered channel of type T, buffer size n
```

​	每个大小实参`n`和`m`必须是[整型](../Types#numeric-types-数值型)，或者是一个只包含整型的类型集，或者是一个无类型的[常量](../Constants)。一个常量大小参数必须是非负数，并且可以用`int`类型的值[表示](../PropertiesOfTypesAndValues#representability-可表示性)；如果它是一个无类型的常量，它被赋予`int`类型。如果`n`和`m`都被提供并且是常量，那么`n`必须**不大于**`m`。对于切片和通道，如果`n`在运行时是负数或者大于`m`，就会发生[运行时恐慌](../Run-timePanics)。

```go 
s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
s := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int => 非法的: len(s) 不能被 int 类型的值表示
s := make([]int, 10, 0)         // illegal: len(s) > cap(s) => 非法的: len(s) > cap(s)
c := make(chan int, 10)         // channel with a buffer size of 10
m := make(map[string]int, 100)  // map with initial space for approximately 100 elements
```

​	调用`map`类型和大小提示`n`的`make`将创建一个初始空间可容纳`n`个map元素的`map`。`具体的行为是依赖于实现的`。

### Appending to and copying slices 追加和复制切片

​	内置函数`append`和`copy`可以帮助进行常见的切片操作。对于这两个函数，其结果与实参所引用的内存是否重叠无关。

The [variadic](https://go.dev/ref/spec#Function_types) function `append` appends zero or more values `x` to a slice `s` and returns the resulting slice of the same type as `s`. The [core type](https://go.dev/ref/spec#Core_types) of `s` must be a slice of type `[]E`. The values `x` are passed to a parameter of type `...E` and the respective [parameter passing rules](https://go.dev/ref/spec#Passing_arguments_to_..._parameters) apply. As a special case, if the core type of `s` is `[]byte`, `append` also accepts a second argument with core type [`bytestring`](https://go.dev/ref/spec#Core_types) followed by `...`. This form appends the bytes of the byte slice or string.

​	[可变参数](../Types#function-types-函数型)函数`append`将**零个或多个值**`x`追加到一个切片`s`，并返回与`s`相同类型的结果切片。值`x`被传递给一个类型为`...E`的参数，各自的[参数传递规则](../Expressions#passing-arguments-to--parameters-向参数传递实参)适用。**作为一个特例**，如果`s`的[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)是`[]byte`，`append`也接受第二个参数，其核心类型是[bytestring](../PropertiesOfTypesAndValues#core-types-核心类型)，后面是`...` 。这种形式追加了字节切片或字符串的字节。

```go 
append(s S, x ...E) S  // core type of S is []E
```

​	如果`s`的容量不足以容纳额外的值，`append`会分配一个新的、足够大的底层数组，同时容纳现有的切片元素和额外的值。否则，`append`复用原来的底层数组。

```go 
s0 := []int{0, 0}
s1 := append(s0, 2)                // append a single element     s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")   //                             t == []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...)            // append string contents      b == []byte{'b', 'a', 'r' }
```

​	函数`copy`将切片元素从源`src`复制到目标`dst`，`并返回复制的元素数量`。两个参数的核心类型必须是具有[一致的](../PropertiesOfTypesAndValues#type-identity-类型一致性)元素类型的切片。复制的元素数是`len(src)`和`len(dst)`中的最小值。**作为一种特殊情况**，如果目标的[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)是`[]byte`，`copy`也接受一个核心类型为[bytestring](../PropertiesOfTypesAndValues##core-types-核心类型)的源参数。这种形式将字节切片或字符串中的字节复制到字节切片中。

```go 
copy(dst, src []T) int
copy(dst []byte, src string) int
```

例子：

```go 
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hello")
```

### Deletion of map elements 删除映射元素

​	内置函数`delete`可以从[映射](../Types#map-types-映射型)`m`中删除`键值`为`k`的元素，值`k`必须可以[分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`m`的键类型。

```go 
delete(m, k)  // remove element m[k] from map m
```

​	如果`m`的类型是[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)，那么该类型集合中的所有类型必须是映射，并且它们必须都有相同的`键类型`。

​	如果映射`m`是`nil`或者元素`m[k]`不存在，delete就是一个空操作。

### Manipulating complex numbers 操纵复数

​	有三个函数`组装`和`分解`复数。内置函数 `complex` 从浮点实部和虚部构造一个复数值，而 `real` 和 `imag` 则提取复数值的实部和虚部。

```go 
complex(realPart, imaginaryPart floatT) complexT
real(complexT) floatT
imag(complexT) floatT
```

​	实参的类型和返回值相对应。对于`complex`函数，两个实参必须是相同的[浮点类型](../Types#numeric-types-数值型)，返回类型是具有对应浮点成分的[复数类型](../Types#numeric-types-数值型)：`float32`类型的实参对应`comple64`复数类型，而`float64`类型的实参对应`comple128`复数类型。如果其中一个实参的值是一个无类型常量，那么它首先会被隐式[转换](../Expressions#conversions-转换)为另一个实参的类型。如果两个参数都求值为无类型常量，那么它们必须是`非复数`，或者它们的虚数部分必须为零，这样函数的返回值就是一个无类型的复数常量。

​	对于`real`和`imag`，实参必须是复数类型，返回类型是相应的浮点类型：`float32`对应`complex64`，`float64`对应`complex128`。如果实参的值是一个无类型的常量，它必须是一个数字，这样函数的返回值就是一个无类型的浮点常量。

​	`real`和`imag`函数一起构成了复数的逆运算，所以对于一个复数类型`Z`的值`z`来说，`z == Z(complex(real(z), imag(z)))`。

​	如果这些函数的操作数都是常量，返回值就是一个常量。

```go 
var a = complex(2, -2)             // complex128 <=仍有疑问？？怎么推导出来是 complex128 ？
const b = complex(1.0, -1.4)       // untyped complex constant 1 - 1.4i
x := float32(math.Cos(math.Pi/2))  // float32
var c64 = complex(5, -x)           // complex64
var s int = complex(1, 0)          // untyped complex constant 1 + 0i can be converted to int
_ = complex(1, 2<<s)               // illegal: 2 assumes floating-point type, cannot shift => 非法的：2 被认为是 浮点类型，不能移位 <=仍有疑问？？2为什么是浮点类型？怎么推导出来的
var rl = real(c64)                 // float32
var im = imag(a)                   // float64
const c = imag(b)                  // untyped constant -1.4
_ = imag(3 << s)                   // illegal: 3 assumes complex type, cannot shift => 非法的：3 被认为是复数类型，不能移位 <=仍有疑问？？3为什么是浮点类型？怎么推导出来的
```

Arguments of type parameter type are not permitted.

不允许使用参数类型的实参。

### Handling panics 处理恐慌

​	两个内置函数，`panic`和`recover`，协助报告和处理[运行时恐慌](../Run-timePanics)和程序定义的错误情况。

```go 
func panic(interface{})
func recover() interface{}
```

​	在执行函数`F`时，对`panic`的显式调用或[运行时恐慌](../Run-timePanics)终止了`F`的执行，然后被`F`延迟的任何函数会照常执行。接下来，任何被`F`的调用者[延迟](../Statements#defer-statements-语句-defer)的函数都会被运行，以此类推，直到被执行中的goroutine中的顶级函数所延迟的任何函数。此时，程序被终止，错误情况被报告，包括`panic`的实参值。这个终止过程被称为`panicking`。

```go 
panic(42)
panic("unreachable")
panic(Error("cannot parse"))
```

​	`recover`函数允许程序管理 panicking goroutine 的行为。假设函数`G`延迟了一个调用`recover`的函数`D`，并且在`G`执行的同一个goroutine上的一个函数发生了恐慌。当被延迟函数的运行到达`D`时，`D`调用`recover`的返回值将是传递给调用`panic`的值。如果`D`正常返回，没有启动新的`panic`，则 panicking goroutine 停止。在这种情况下，`G`和对`panic`的调用之间调用的函数的状态被丢弃，并恢复正常执行。`G`在`D`之前被延迟的任何函数随后被运行，`G`的执行通过返回给它的调用者而终止。

​	如果以下任何一个条件成立，`recover`的返回值为`nil`：

- `panic`的参数是空的。
- goroutine没有陷入panicking。
- `recover`不是由被延迟函数直接调用的。

下面的例子中的`protect`函数调用了函数实参`g`，并保护调用者免受`g`引发的运行时恐慌。

```go 
func protect(g func()) {
	defer func() {
		log.Println("done")  // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}
```

### Bootstrapping 引导

​	目前的实现提供了几个在引导（bootstrapping）过程中有用的内置函数。为了完整起见，这些函数被记录下来，但不保证会留在语言中。它们并不返回结果。

```
Function   Behavior

print      prints all arguments; formatting of arguments is implementation-specific => 打印所有实参；实参的格式化和实现有关
println    like print but prints spaces between arguments and a newline at the end => 和 print 类似，但是会在每个实参间打印空格，在结尾打印新行
```

实现限制：`print`和`println`不一定需要接受任意的实参类型，但必须支持布尔型、数字型和字符串型的打印。