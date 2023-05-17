+++
title = "Packages, variables, and functions"
weight = 1
date = 2023-05-17T12:10:24+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Packages, variables, and functions

## Packages 包

[https://go.dev/tour/basics/1](https://go.dev/tour/basics/1)

​	每个Go程序都是由包组成的。

​	程序从`main`包中开始运行。

​	本程序正在使用导入路径为 `"fmt"`和 `"math/rand "`的包。

​	按照约定，包名与导入路径的最后一个元素一致。例如，`"math/rand "`包中的源码均以`package rand`语句开始。

!!! waring "注意"

	注意：这些程序的执行环境是固定的，所以每次运行示例程序`rand.Intn`都会返回同一个数字。

(要想看到一个不同的数字，请给数字生成器不同的种子数；见[rand.Seed](https://go.dev/pkg/math/rand/#Seed)。在练习场中的时间是常量，因此你需要使用其他的值作为种子数）。

```go title="main.go" linenums="1"
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

## Imports 导入

[https://go.dev/tour/basics/2](https://go.dev/tour/basics/2)

​	此代码用圆括号组合了导入，这是 "分组 "形式的导入语句。

​	你也可以写多个导入语句，例如：

```go linenums="1"
import "fmt"
import "math"
```

​	不过，使用分组导入语句是更好的形式。

```go title="main.go" linenums="1"
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

```

## Exported names 导出名

[https://go.dev/tour/basics/3](https://go.dev/tour/basics/3)

​	在Go中，如果一个名字以大写字母开头，那么它就是已导出的。例如，`Pizza`是就是一个已导出名，`Pi`也是一个已导出名，它导出自`math`包。

​	`pizza` and `pi` do not start with a capital letter, so they are not exported.

​	`pizza`和`pi`并未以大写字母开头，所以它们是未导出的。

​	在导入一个包时，你只能引用其中已导出的名字。任何 "未导出的 "名字在该包外均无法访问。

​	运行这段代码。注意错误信息。

​	为了解决这个错误，将`math.pi`重命名为`math.Pi`，然后再试一次。

```go title="main.go" linenums="1"
package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.pi) // 错误
    fmt.Println(math.Pi) // 正确
}

```

## Functions 函数

[https://go.dev/tour/basics/4](https://go.dev/tour/basics/4)

​	函数可以接受零个或多个参数。

​	在本例中，`add`接受两个`int`类型的参数。

​	请注意，类型是在变量名之后。

​	(更多关于这种声明形式出现的原因，请参见[Go的声明语法](../../../GoBlog/2010/GosDeclarationSyntax)一文）。

```go title="main.go" linenums="1"
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

```

## Functions continued 函数（续）

[https://go.dev/tour/basics/5](https://go.dev/tour/basics/5)

​	当连续两个或更多的函数的已命名形参类型相同时，除了最后一个以外，其它都可以省略。

In this example, we shortened

在本例子中，

```
x int, y int
```

被缩写为：

```
x, y int
```

```go title="main.go" linenums="1" hl_lines="5 5"
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

```

##  Multiple results 多值返回

[https://go.dev/tour/basics/6](https://go.dev/tour/basics/6)

函数可以返回任意数量的返回值。

`swap`函数返回两个字符串。

```go title="main.go" linenums="1" hl_lines="5 5"
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

```

##  Named return values 命名返回值

[https://go.dev/tour/basics/7](https://go.dev/tour/basics/7)

​	Go 的返回值可被命名。它们被视作定义在函数顶部的变量。

​	返回值的名称应当具有一定的意义，它可以作为文档使用。

​	没有参数的`return`语句会返回已命名的返回值。也就是直接返回。

​	直接返回语句应当仅用在下面这样的短函数中。在较长的函数中，它们会影响代码的可读性。

```go title="main.go" linenums="1" hl_lines="8 8"
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

```

##  Variables 变量

[https://go.dev/tour/basics/8](https://go.dev/tour/basics/8)

​	`var`语句用于声明一个变量列表，和函数参数列表一样，类型在最后。

​	就像本例子中看到的一样，`var` 语句可以出现在包或函数级别。

```go title="main.go" linenums="1"
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

```

## Variables with initializers 变量初始化

[https://go.dev/tour/basics/9](https://go.dev/tour/basics/9)

​	var声明可以包含初始化值，每个变量对应一个。

​	如果初始化值已经存在，可以省略类型；变量会从初始化值中获得类型。

```go title="main.go" linenums="1"
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

```

##  Short variable declarations 短变量声明

[https://go.dev/tour/basics/10](https://go.dev/tour/basics/10)

​	在函数中，可用 `:=` 短赋值语句来代替隐含类型的`var`声明。

​	在函数外，每个语句都以关键字开始（`var`、`func`等），因此`:=`结构不能在函数外使用。

```go title="main.go" linenums="1"
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

```

## Basic types 

[https://go.dev/tour/basics/11](https://go.dev/tour/basics/11)

​	Go的基本类型有：

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

​	本例显示了几种类型的变量，同导入语句一样，变量声明也可被 "分组 "成块。

​	`int`、`uint`和`uintptr`类型在32位系统上通常是32位宽，在64位系统上是64位宽。当你需要一个整数值时，你应该使用`int`，除非你有特别的理由使用一个固定大小或无符号的整数类型。

```go title="main.go" linenums="1"
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

```

##  Zero values 零值

[https://go.dev/tour/basics/12](https://go.dev/tour/basics/12)

​	在没有明确初始值的变量声明中，变量会被赋予零值。

​	零值是：

- 数值类型为`0`，
- 布尔类型为`false`，而
- 字符串为`""`（空字符串）。

```go title="main.go" linenums="1"
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

```

##  Type conversions 类型转换

[https://go.dev/tour/basics/13](https://go.dev/tour/basics/13)

​	表达式`T(v)`将值`v`转换为类型`T`。

​	一些关于数值的转换：

```go linenums="1"
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

或者，更简单：

```go linenums="1"
i := 42
f := float64(i)
u := uint(f)
```

​	与C语言不同，Go语言中不同类型的项之间的赋值需要显示转换。尝试去掉例子中的`float64`或`uint`转换，看看会发生什么。

```go title="main.go" linenums="1"
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

```

##  Type inference 类型推导

[https://go.dev/tour/basics/14](https://go.dev/tour/basics/14)

​	当声明一个变量而不指定明确的类型时（无论是使用`:=`语法还是`var=`表达式语法），变量的类型是由右侧的值推导出来的。

​	当右值声明了类型时，新变量的类型与其相同：

```go linenums="1"
var i int
j := i // j is an int
```

​	但是当右边包含一个无类型的数值常量时，新变量可能是`int`、`float64`或者`complex128`，这取决于常量的精度：

```go linenums="1"
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

​	试着改变示例代码中`v`的初始值，观察它是如何影响类型的。

```go title="main.go" linenums="1"
package main

import "fmt"

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}

```

##  Constants 常量

[https://go.dev/tour/basics/15](https://go.dev/tour/basics/15)

​	常量的声明与变量类似，但使用`const`关键字。

​	常量可以是字符、字符串、布尔值或数字值。

​	常量不能用 `:=` 语法来声明。

```go title="main.go" linenums="1"
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

```

##  Numeric Constants 数值常量

[https://go.dev/tour/basics/16](https://go.dev/tour/basics/16)

​	数值常量是高精度的数值。

​	一个`无类型`的常量`由上下文来决定其类型`。

​	再尝试一下输出`needInt(Big)`。

(一个`int`最多能存储64位的整数，（根据所用的系统平台）有时甚至更少。)

```go title="main.go" linenums="1"
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

```

##  Congratulations! 祝贺你!

[https://go.dev/tour/basics/17](https://go.dev/tour/basics/17)

​	你完成了这一课!

​	你可以返回`模块`列表中寻找下一步要学习的内容，或者继续学习下一课。

