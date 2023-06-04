+++
title = "Flow control statements: for, if, else, switch and defer"
weight = 2
date = 2023-05-17T12:10:24+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Flow control statements: for, if, else, switch and defer

##  For

[https://go.dev/tour/flowcontrol/1](https://go.dev/tour/flowcontrol/1)

​	Go 只有一个循环结构：`for` 循环。

​	基本的for循环由三个部分组成，它们用分号隔开：

- 初始化语句：在第一次迭代前执行
- 条件表达式：在每次迭代前进行求值
- 后置语句：在每次迭代的结尾执行。

​	初始化语句通常是一个短变量声明，其中声明的变量只在for语句的作用域中可见。

​	一旦条件表达式求值为`false`，循环迭代将停止。

注意：与其他语言如C、Java或JavaScript不同的是，`for`语句的三个部分没有小括号，大括号`{ }`总是必需的。

```go title="main.go" 
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

```

## For continued - for (续)

[https://go.dev/tour/flowcontrol/2](https://go.dev/tour/flowcontrol/2)

初始化语句和后置语句都是可选的。

```go title="main.go" 
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

```

##  For is Go's "while"

[https://go.dev/tour/flowcontrol/3](https://go.dev/tour/flowcontrol/3)

​	此时，您可以放弃分号。C的`while`在Go中叫做`for`。

```go title="main.go" 
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

```

##  Forever 无限循环

[https://go.dev/tour/flowcontrol/4](https://go.dev/tour/flowcontrol/4)

​	如果您省略了循环条件，它就会永远循环下去。因此无限循环可以写得很紧凑。

```go title="main.go" 
package main

func main() {
	for {
	}
}

```

##  If

[https://go.dev/tour/flowcontrol/5](https://go.dev/tour/flowcontrol/5)

​	Go的`if`语句与它的`for`循环类似；表达式`不需要`用小括号`()`包围，`但需要`用大括号`{}`。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

```

## If with a short statement - if 的简短语句

[https://go.dev/tour/flowcontrol/6](https://go.dev/tour/flowcontrol/6)

​	和`for`一样，`if`语句可以在条件表达式之前执行一个简短语句。

​	该语句所声明变量的作用域仅在`if`之内。

​	(尝试在最后的`return`语句中使用`v`)。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

```

##  If and else

[https://go.dev/tour/flowcontrol/7](https://go.dev/tour/flowcontrol/7)

​	在`if`简短语句中声明的变量`也可以`在任何对应的`else`块中使用。

​	(对`pow`的两次调用都在`main`中对`fmt.Println`的调用开始之前返回其结果）。

```go title="main.go" 
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

```

##  Exercise: Loops and Functions 练习：循环和函数

[https://go.dev/tour/flowcontrol/8](https://go.dev/tour/flowcontrol/8)

​	为了练习函数和循环，我们来实现一个平方根函数：给定一个数字`x`，我们想找到`z²`最接近`x`的数字`z`。

​	计算机通常使用循环来计算`x`的平方根。从某个猜测的`z`开始，我们可以根据`z²`与`x`的接近程度来调整`z`，产生一个更好的猜测：

```
z -= (z*z - x) / (2*z)
```

​	重复这种调整使猜测越来越精确，直到我们得到一个尽可能接近实际平方根的答案。

​	在提供的函数`Sqrt`中实现这一点。无论输入什么，对`z`的一个合适的起始猜测是`1`。首先，重复计算`10`次，并打印每次的`z`值。观察对于不同的`x`值（1，2，3，......），您得到的答案是如何接近结果的，以及猜测提升的速度有多快。

提示：要声明和初始化一个浮点值，要给它以浮点的语法或使用转换：

```go 
z := 1.0
z := float64(1)
```

​	接下来，修改循环条件，一旦数值停止变化（或者只变化了很小的量）就停止。观察迭代次数多于还是少于10次。尝试其他的`z`的初始猜测，比如`x`，或者`x/2`。您的函数的结果与标准库中的`math.Sqrt`有多接近？

(注：如果您对算法的细节感兴趣，上面的`z²-x`是指`z²`到它所要到达的值（即x）的距离，除以`2z`为`z²`的导数，以缩放我们通过`z²`的变化速度调整`z`的程度。这种一般方法被称为[牛顿法](https://en.wikipedia.org/wiki/Newton's_method)。它对许多函数都很有效，但对平方根尤其有效)。

```go title="main.go" 
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
}

func main() {
	fmt.Println(Sqrt(2))
}

```

##  Switch

[https://go.dev/tour/flowcontrol/9](https://go.dev/tour/flowcontrol/9)

​	`switch`语句是写一连串`if-else`语句的一种更简短的方式。它运行第一个值等于条件表达式的`case`语句。

​	Go的`switch`与C、C++、Java、JavaScript和PHP中的`switch`一样，只是Go只运行选定的`case`，而不是之后的所有情况。实际上，Go中是自动提供了在这些语言中每个case结尾所需的break语句。【除非在每个case语句结尾加上 `fallthrough`，否则case 分支会自动终止】。另一个重要的区别是，Go的`switch` 的 `case`无需是常量，且值也不必是整数。

```go title="main.go" 
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

```

## Switch evaluation order

[https://go.dev/tour/flowcontrol/10](https://go.dev/tour/flowcontrol/10)

`switch` 的 `case`语句从上到下顺序执行，直到匹配成功时停止。

(例如：

```go  
switch i {
case 0:
case f():
}
```

在  `i == 0`时，`f`不会被调用。）

**Note:** Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader.

注意：Go练习场上的时间总是从`2009-11-10 23:00:00 UTC`开始，这个数值的意义留给读者发现。

```go title="main.go" 
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

```

##  Switch with no condition 没有条件的 switch

[https://go.dev/tour/flowcontrol/11](https://go.dev/tour/flowcontrol/11)

​	没有条件的`switch`和`switch true`是一样的。

​	这种形式能将一长串的`if-then-else`写得更加简洁。

```go title="main.go" 
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

```

## Defer

[https://go.dev/tour/flowcontrol/12](https://go.dev/tour/flowcontrol/12)

​	`defer` 语句将一个函数的执行推迟到外层函数返回的最后执行。

​	推迟调用的函数的参数会被立即求值，但直到外层函数返回之前，该函数都不会被调用。

```go title="main.go" 
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

```

##  Stacking defers

[https://go.dev/tour/flowcontrol/13](https://go.dev/tour/flowcontrol/13)

​	推迟的函数调用被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。

​	要了解更多关于`defer`语句的信息，请阅读[这篇博文]({{< ref "/goBlog/2010/DeferPanicAandRecover" >}})。

```go title="main.go" 
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

```

##  Congratulations!

[https://go.dev/tour/flowcontrol/14](https://go.dev/tour/flowcontrol/14)

​	您完成了这一课!

​	您可以回到`模块`列表中寻找下一步要学习的内容，或者继续学习下一课。