+++
title = "方法和接口"
linkTitle = "方法和接口"
weight = 3
date = 2023-05-17T12:10:24+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Methods and interfaces

##  Methods 方法

> 原文：[https://go.dev/tour/methods/1](https://go.dev/tour/methods/1)

​	Go没有类。不过，您可以在类型上定义方法。

​	方法是一种带有特殊*接收器*参数的函数。

​	方法接收器出现在它自己的参数列表中，位于`func`关键字和方法名称之间。

​	在这个例子中，`Abs` 方法有一个名为 `v` 的`Vertex`类型的接收器。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

```

##  Methods are functions 方法即函数

> 原文：[https://go.dev/tour/methods/2](https://go.dev/tour/methods/2)

​	记住：方法只是一个带有接收器参数的函数。

​	下面是`Abs`写成的普通函数，功能上并没有变化。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

```

##  Methods continued 方法（续）

> 原文：[https://go.dev/tour/methods/3](https://go.dev/tour/methods/3)

​	您也可以在非结构体类型上声明方法。

​	在这个例子中，我们看到一个数值类型`MyFloat`有一个`Abs`方法。

​	您只能声明一个带有`接收器`的方法，其类型与该方法定义在同一个包中。您不能用一个类型定义在另一个包中的接收器来声明一个方法（其中包括内置类型，如`int`）。（即：接收器的类型定义和方法声明必须在同一包内，不能为内建类型声明方法）

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

```

##  Pointer receivers 指针接收器

> 原文：[https://go.dev/tour/methods/4](https://go.dev/tour/methods/4)

​	您可以声明带有指针接收器的方法。

This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`.)

​	这意味着对于某个类型的`T`，接收器的类型可以用`*T`字面量语法。（此外，`T`本身不能是一个指针，如`*int`）。

​	例如，这里为`*Vertex`定义了`Scale`方法。

​	具有指针接收器的方法可以修改接收器所指向的值（如这里的`Scale`）。由于方法经常需要修改它们的接收器，`指针接收器比值接收器更常见`。

​	试着从第16行的`Scale`函数的声明中去掉`*`，观察程序的行为如何变化。

​	若使用值接收器，`Scale`方法在原始`Vertex`值的副本上操作。(这与其他函数实参的行为相同。)`Scale`方法必须有一个指针接收器来改变`main`函数中声明的`Vertex`值。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

```

##  Pointers and functions 指针和函数

> 原文：[https://go.dev/tour/methods/5](https://go.dev/tour/methods/5)

​	这里我们看到`Abs`和`Scale`方法被改写成了函数。

​	再一次，试着把第16行中的`*`去掉。您能明白为什么行为发生了变化吗？为了编译这个例子，您还需要改变什么？

(如果您不确定，继续看下一页）。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

```

##  Methods and pointer indirection 方法和指针重定向

> 原文：[https://go.dev/tour/methods/6](https://go.dev/tour/methods/6)

​	对比前面两个程序，您可能会注意到，带有指针参数的函数必须接受一个指针。

```go 
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

​	而`带有指针接收器的方法`在被调用时，接收器既能为值，也能为指针。

```go 
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

​	对于`v.Scale(5)`这个语句，即使`v`是一个值而不是一个指针，也会自动调用带有指针接收器的方法。也就是说，为了方便起见，Go将语句`v.Scale(5)`解释为`(&v).Scale(5)`，因为`Scale`方法有一个指针接收器。

```go title="main.go" 
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

```

## Methods and pointer indirection (2) 方法和指针重定向 (2)

> 原文：[https://go.dev/tour/methods/7](https://go.dev/tour/methods/7)

​	同样的事情也发生在相反的方向上。

​	接受一个值形参的函数必须接受一个特定类型的值：

```go 
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

​	而以值为接收器的方法在被调用时，接收器既能为值，也能为指针：

```go 
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

在这种情况下，方法调用`p.Abs()`被解释为`(*p).Abs()`。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

```

##  Choosing a value or pointer receiver 选择值或指针接收器

> 原文：[https://go.dev/tour/methods/8](https://go.dev/tour/methods/8)

​	使用指针接收器有两个原因：

​	首先，方法可以修改其接收器所指向的值。

​	其次，这样可以避免在每次方法调用时复制值。若值的类型为一个大的结构，这可能会更有效率。

​	在这个例子中，`Scale`和`Abs`的接收器类型都是`*Vertex`，尽管`Abs`方法不需要修改其接收器。

​	通常来说，给定类型的所有方法都应该有值接收器或指针接收器，而不是两者的混合。(我们将在接下来的几页中看到原因）。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

```

##  Interfaces 接口

> 原文：[https://go.dev/tour/methods/9](https://go.dev/tour/methods/9)

​	接口类型被定义为一组方法签名的集合。

​	接口类型的变量可以保存任何实现这些方法的值。

注意：在示例代码的第22行有一个错误。`Vertex`（值类型）没有实现`Abser`，因为`Abs`方法只定义在`*Vertex`（指针类型）。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```

## Interfaces are implemented implicitly 接口与隐式实现

> 原文：[https://go.dev/tour/methods/10](https://go.dev/tour/methods/10)

​	类型通过实现其方法来实现一个接口。既无需专门显式声明，也没有"implements"关键字。

​	`隐式接口`将接口的定义与它的实现解耦，这样接口的实现就可以在任何包中出现，而无需提前准备。

```go title="main.go" 
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// => 这个方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

```

## Interface values 接口值

> 原文：[https://go.dev/tour/methods/11](https://go.dev/tour/methods/11)

​	在底层，接口值可以被认为是一个值和具体类型的元组:

```
(value, type)
```

​	接口值保存了一个特定底层具体类型的值。

​	在接口值上调用一个方法会在其底层类型上执行同名的方法。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

##  Interface values with nil underlying values 底层值为nil的接口值

> 原文：[https://go.dev/tour/methods/12](https://go.dev/tour/methods/12)

​	即便接口本身的具体值是`nil`，方法仍将被调用，其接收器为`nil`。

​	在某些语言中，这将引发空指针异常，但在Go中，通常会编写一些方法，以优雅的方式处理被调用的`nil`接收器（如本例中的方法`M`）。

​	请注意，保存了`nil`具体值的接口值本身并不是`nil`。

```go title="main.go" 
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

##  Nil interface values  - nil 接口值

> 原文：[https://go.dev/tour/methods/13](https://go.dev/tour/methods/13)

​	`nil`接口值既不保存值也不保存具体类型。

​	在`nil`接口上调用方法会产生运行时错误，因为在接口元组里面并未包含能够指明调用哪个具体方法的类型。

```go title="main.go" 
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

##  The empty interface 空接口

> 原文：[https://go.dev/tour/methods/14](https://go.dev/tour/methods/14)

​	未指定任何方法的接口类型被称为`空接口`。

```go 
interface{}
```

​	空接口可以保存任何类型的值。(每个类型都至少实现了零个方法）。

​	空接口被用来处理未知类型的值。例如，`fmt.Print`接受任何数量的`interface{}`类型的形参。

```go title="main.go" 
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

##  Type assertions 类型断言

> 原文：[https://go.dev/tour/methods/15](https://go.dev/tour/methods/15)

​	类型断言提供了对一个接口值的底层具体值的访问。

```go 
t := i.(T)
```

​	该语句断言接口值`i`保存了具体类型`T`的值，并将底层类型`T`的值分配给变量`t`。

​	如果`i`并不保存`T`类型的值，该语句将触发一个恐慌。

​	为了测试一个接口值是否持有一个特定的类型，类型断言可以返回两个值：`底层值`和一个报告断言是否成功的`布尔值`。

```go 
t, ok := i.(T)
```

​	如果`i`保存了一个`T`类型的值，那么`t`将是其底层值，`ok`将为`true`。

​	否则，`ok`将是`false`，`t`将是`T`类型的`零值`，`并且程序不会发生恐慌。`

​	注意这种语法和读取一个映射的语法有相似之处。

```go title="main.go" 
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

```

##  Type switches 类型选择

> 原文：[https://go.dev/tour/methods/16](https://go.dev/tour/methods/16)

​	`类型选择`是一种按顺序从几个`类型断言`中选择分支的结构。

​	`类型选择`与一般的switch语句相似，但是类型选择中的`case分支为类型（而不是值）`，这些值与给定的接口值所保存的值类型进行比较。

```go 
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

​	`类型选择`中的声明与类型断言`i.(T)`的语法相同，但具体的类型`T`被替换为`关键字type`。

​	这个选择语句测试接口值`i`保存值的类型是`T`还是`S`类型。在 `T` 或 `S` 的情况下，变量 `v` 会分别按 `T` 或 `S` 类型保存 `i` 拥有的值。**在默认情况下（没有匹配），变量`v`与`i`的接口类型和值相同**。

```go title="main.go" 
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}

```

##  Stringers 串联器

> 原文：[https://go.dev/tour/methods/17](https://go.dev/tour/methods/17)

​	`fmt`包中定义的[Stringer](https://go.dev/pkg/fmt/#Stringer)是最普遍的接口之一。

```go 
type Stringer interface {
    String() string
}
```

​	`Stringer`是一个可以字符串描述自己的类型。`fmt`包（和其他许多包）都通过此接口来打印数值。

```go title="main.go" 
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

```

## Exercise: Stringers 练习：Stringer

> 原文：[https://go.dev/tour/methods/18](https://go.dev/tour/methods/18)

​	通过让`IPAddr`类型实现`fmt.Stringer`，来打印为点号分隔的地址。

​	例如，`IPAddr{1, 2, 3, 4}`应打印为 `"1.2.3.4"`。

```go title="main.go" 
package main

import (
    "fmt"
    "strings"
	"strconv"    
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
   s := make([]string, len(ip))
	for i, val := range ip {
		s[i] = strconv.Itoa(int(val))
	}
	return fmt.Sprintf(strings.Join(s, "."))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

```

## Errors 错误

> 原文：[https://go.dev/tour/methods/19](https://go.dev/tour/methods/19)

​	Go程序用`error`值来表达错误状态。

​	`error`类型是一个类似于`fmt.Stringer`的内置接口。

```go 
type error interface {
    Error() string
}
```

(与`fmt.Stringer`类似，`fmt`包在打印值时也会寻找`error`接口)。

​	函数通常返回一个`error`值，调用它的代码应该通过测试错误是否等于`nil`来处理错误。

```go 
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

`error`为`nil`表示成功；非`nil`的 `error`表示失败。

```go title="main.go" 
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

```

##  Exercise: Errors 练习：错误

> 原文：[https://go.dev/tour/methods/20](https://go.dev/tour/methods/20)

​	从[之前的练习](#choosing-a-value-or-pointer-receiver)中复制`Sqrt`函数，并修改它以返回一个`error`值。

​	`Sqrt`接收到一个负数时，应当返回一个非`nil`的`error` 值，复数同样也不被不支持。

创建一个新的类型

```go 
type ErrNegativeSqrt float64
```

并为其实现：

```go 
func (e ErrNegativeSqrt) Error() string
```

方法使其拥有`error`值，这样`ErrNegativeSqrt(-2).Error()`返回 "`cannot Sqrt negative number: -2`"。

注意：在`Error`方法中调用`fmt.Sprint(e)`将`使程序进入死循环`。您可以通过先转换`e`来避免这种情况：`fmt.Sprint(float64(e))`。为什么？

​	修改您的`Sqrt`函数，当给定一个负数时返回一个`ErrNegativeSqrt`值。

```go title="main.go" 
package main

import (
	"fmt"
    "math"
)

type ErrNegativeSqrt float64


func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if (x > 0) {
		z:= x/2
		for i:= 0; math.Abs(z*z - x) > 0.0000000001; i++ {
			z -= (z*z - x) / (2*z)
			fmt.Println(i, "z:", z, "z^2 -x:", z*z - x)
		}
		return z, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

```

##  Readers 

> 原文：[https://go.dev/tour/methods/21](https://go.dev/tour/methods/21)

​	`io` 包指定了 `io.Reader` 接口，该接口表示数据流的读取端。

​	Go标准库包含了这个接口的[许多实现](https://cs.opensource.google/search?q=Read\(\w%2B\s\[\]byte\)&ss=go%2Fgo)，包括文件、网络连接、压缩、加密等。

​	`io.Reader` 接口有一个 `Read` 方法：

```go 
func (T) Read(b []byte) (n int, err error)
```

`Read`用数据填充给定的`字节切片`，并返回填充的字节数和错误值。当数据流结束时，它返回一个`io.EOF`错误。

​	该示例代码创建了一个[strings.Reader](https://go.dev/pkg/strings/#Reader)，并以每次8个字节的速度读取它的输出。

```go title="main.go" 
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

```

##  Exercise: Readers 练习：Readers 

> 原文：[https://go.dev/tour/methods/22](https://go.dev/tour/methods/22)

​	实现一个`Reader`类型，它可以产生一个ASCII字符'A'的无限流。

```go title="main.go" 
package main

import (
	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

type ErrEmptyBuffer []byte

func (b ErrEmptyBuffer) Error() string {
	return fmt.Sprintf("cannot read an empty buffer: %v", b)
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(b []byte) (int, error) {
	bLen := len(b)
	if (bLen == 0) {
		return 0, ErrEmptyBuffer(b)
	}
	for i := range b {
		b[i] = 'A'
	}
	return bLen, nil
}


func main() {
	reader.Validate(MyReader{})
}

```

## Exercise: rot13Reader 练习：Rot13Reader

> 原文：[https://go.dev/tour/methods/23](https://go.dev/tour/methods/23)

​	一个常见的模式是一个[io.Reader](https://go.dev/pkg/io/#Reader)包装另一个`io.Reader`，然后通过某种方式修改其数据流。

​	例如，[gzip.NewReader](https://go.dev/pkg/compress/gzip/#NewReader)函数接收一个 `io.Reader`（压缩数据流）并返回一个同样实现了`io.Reader`的 `*gzip.Reader`（解压缩后的数据流）。

​	编写一个实现`io.Reader`并从另一个`io.Reader`中读取数据的`rot13Reader`，通过应用[rot13](https://en.wikipedia.org/wiki/ROT13)代换密码对数据流进行修改。

​	`rot13Reader`的类型是为您提供了。通过实现它的`Read`方法使其成为实现`io.Reader`。

```go title="main.go" 
package main

import (
	"io"
	"os"

	"strings"
)

type rot13Reader struct{ r io.Reader }

func rot13(c byte) byte {
	switch {
	case (c >= 'A' && c <= 'M') || (c >= 'a' && c <= 'm'):
		c += 13
	case (c >= 'N' && c <= 'Z') || (c >= 'n' && c <= 'z'):
		c -= 13
	}
	return c
}

func (reader *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = reader.r.Read(b)
	for i := range b {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

```

##  Images 图像

> 原文：[https://go.dev/tour/methods/24](https://go.dev/tour/methods/24)

​	[image](https://go.dev/pkg/image/#Image)包定义了`Image`接口。

```go 
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

注意：`Bounds`方法的返回值`Rectangle`实际上是一个[image.Rectangle](https://go.dev/pkg/image/#Rectangle)，它在`image`包中声明。

(所有细节见[文档](https://go.dev/pkg/image/#Image))。

​	`color.Color` 和 `color.Model` 类型也是接口，但是通常因为直接使用预定义的实现 `image.RGBA` 和 `image.RGBAModel` 而被忽视了。这些接口和类型由 [image/color](https://go-zh.org/pkg/image/color/) 包定义。

```go title="main.go" 
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

```

##  Exercise: Images 练习：图片

> 原文：[https://go.dev/tour/methods/25](https://go.dev/tour/methods/25)

​	还记得您之前写的[图片生成器](#exercise-stringers-stringers)吗？让我们再写一个，但这次它将返回`image.Image`的实现，而不是一个数据切片。

​	定义您自己的`Image`类型，实现[必要的方法](https://go.dev/pkg/image/#Image)，然后调用`pic.ShowImage`。

​	`Bounds`应当返回一个`image.Rectangle`，例如`image.Rect(0, 0, w, h)`。

`ColorModel`应当返回`color.RGBAModel`。

`At`应当返回一种颜色；上一个图片生成器中的值`v`对应于此次的`color.RGBA{v, v, 255, 255}`。

```go title="main.go" 
package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	W int
	H int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

func (i Image) At(x, y int)  color.Color {
	v := uint8(x*y + y*y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{200, 200}
	pic.ShowImage(m)
}

```

## Congratulations! 祝贺您!

> 原文：[https://go.dev/tour/methods/26](https://go.dev/tour/methods/26)

您完成了这一课!

您可以回到`模块`列表中寻找下一步要学习的内容，或者继续学习下一课。