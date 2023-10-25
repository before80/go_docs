+++
title = "time"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/time@go1.20.1

​	time包提供了测量和显示时间的功能。

​	日历计算总是假定[格里高利日历](https://zh.wikipedia.org/wiki/%E5%85%AC%E5%8E%86)，没有[闰秒](https://zh.wikipedia.org/wiki/%E9%97%B0%E7%A7%92)。

#### 单调钟 Monotonic Clocks 

​	操作系统提供了“挂钟(wall clock)”和“单调钟(monotonic clock)”两种时钟，前者会受到时间同步的影响，后者则不会。一般规则是挂钟用于显示时间，而单调钟用于计时。为了不分割API，在这个包中，`time.Now`返回的时间包含了挂钟读数和单调钟读数；后续的时间显示操作使用挂钟读数，而后续的时间计算操作，特别是比较和减法操作，使用单调钟读数。

​	例如，即使在计时操作中挂钟被更改，以下代码始终计算出大约20毫秒的正计时：

```go
start := time.Now()
... operation that takes 20 milliseconds ...
t := time.Now()
elapsed := t.Sub(start)
```

​	其他的习语，比如`time.Since(start)`，`time.Until(deadline)`，和`time.Now().Before(deadline)`，同样能够很好地应对挂钟重置的情况。

​	本节的其余部分详细介绍了操作如何使用单调钟，但理解这些细节并不是使用该包的必要条件。

​	`time.Now`返回的Time包含一个单调钟读数。如果Time `t`具有单调钟读数，则`t.Add`会将相同的持续时间添加到挂钟和单调钟读数中，以计算结果。因为`t.AddDate(y, m, d)`，`t.Round(d)`，和`t.Truncate(d)`都是针对挂钟时间的计算，它们总是从结果中剥离任何单调钟读数。因为`t.In`，`t.Local`，和`t.UTC`用于影响挂钟时间解释的效果，它们也会从结果中剥离任何单调钟读数。剥离单调钟读数的规范方法是使用`t = t.Round(0)`。

​	如果Times `t`和u都包含单调钟读数，则操作`t.After(u)`，`t.Before(u)`，`t.Equal(u)`，`t.Compare(u)`，和`t.Sub(u)`仅使用单调钟读数进行计算，忽略挂钟读数。如果`t`或`u`中没有单调钟读数，则这些操作将回退到使用挂钟读数。

​	在某些系统上，如果计算机进入睡眠状态，单调钟会停止。在这种情况下，`t.Sub(u)`可能无法准确反映`t`和`u`之间经过的实际时间。

​	由于单调钟读数在当前进程之外没有意义，因此`t.GobEncode`、`t.MarshalBinary`、`t.MarshalJSON`和`t.MarshalText`生成的序列化形式都省略了单调钟读数，而`t.Format`则不提供任何格式。同样，构造函数`time.Date`、`time.Parse`、`time.ParseInLocation`和`time.Unix`，以及解码器`t.GobDecode`、`t.UnmarshalBinary`、`t.UnmarshalJSON`和`t.UnmarshalText`始终创建没有单调钟读数的时间。

​	**单调钟读数仅存在于Time值中**。它不是Duration值或`t.Unix`和相关函数返回的Unix时间的一部分。

​	请注意，Go的`==`运算符不仅比较时间瞬间，还比较Location (时区)和单调钟读数。有关Time值的相等性测试的讨论，请参阅Time类型的文档。

​	为了调试，如果存在单调钟读数，则`t.String`的结果会包含该读数。如果`t != u`，因为具有不同的单调钟读数，那么在打印`t.String()`和`u.String()`时将可见这种差异。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/time/format.go;l=101)

``` go 
const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	DateTime   = "2006-01-02 15:04:05"
	DateOnly   = "2006-01-02"
	TimeOnly   = "15:04:05"
)
```

​	这些是用于`Time.Format`和`time.Parse`的预定义布局。这些布局中使用的参考时间是特定的时间戳：

```
01/02 03:04:05PM '06 -0700
```

（2006年1月2日，15:04:05，位于GMT西七小时的时区）。该值被记录为名为Layout的常量，如下所示。作为Unix时间，这是1136239445。由于MST是GMT-0700，Unix date命令将打印出该参考值：

```
Mon Jan 2 15:04:05 MST 2006
```

​	令人遗憾的是，该日期使用将数字月份放在日前的美国惯例。

​	Time.Format 的示例详细演示了布局字符串的工作方式，是一个很好的参考。

​	请注意，RFC822、RFC850 和 RFC1123 格式应仅应用于本地时间。将它们应用于 UTC 时间将使用 "UTC" 作为时区缩写，而严格来说，这些 RFC 要求在这种情况下使用 "GMT"。通常应该使用 RFC1123Z 替代 RFC1123，用于坚持该格式的服务器，并且应该为新协议首选 RFC3339。RFC3339、RFC822、RFC822Z、RFC1123 和 RFC1123Z 适用于格式化；与 time.Parse 一起使用时，它们不接受 RFC 允许的所有时间格式，并且它们接受形式上未定义的时间格式。RFC3339Nano 格式从秒字段中删除尾随的零，因此一旦格式化，可能无法正确排序。

​	大多数程序可以使用定义的常量之一作为传递给 Format 或 Parse 的布局。除非您要创建自定义布局字符串，否则可以忽略此注释的其余部分。

​	要定义自己的格式，请编写引用时间以您的方式格式化的样子；请参阅 ANSIC、StampMicro 或 Kitchen 等常量的值。模型是演示引用时间的外观，以便 Format 和 Parse 方法可以将相同的转换应用于一般时间值。

​	以下是布局字符串的组件摘要。每个元素都以示例形式显示了参考时间的一个元素的格式。仅识别这些值。布局字符串中未被识别为参考时间的文本在 Format 中被回显，并期望在 Parse 的输入中以原样出现。



​	这是一个令人遗憾的历史错误，日期使用了美国的习惯，在日期之前放置数字月份。

​	`Time.Format`的示例详细演示了布局字符串的工作原理，是一个很好的参考。

​	请注意，**RFC822、RFC850和RFC1123格式仅适用于本地时间**。将它们应用于UTC时间将使用"UTC"作为时区缩写，严格来说，这些RFC要求在这种情况下使用"GMT"。一般应该使用RFC1123Z而不是RFC1123，以满足一些服务器对该格式的要求，并且对于新协议应该优先选择RFC3339。RFC3339、RFC822、RFC822Z、RFC1123和RFC1123Z适用于格式化；当与`time.Parse`一起使用时，它们不接受RFC允许的所有时间格式，但可以接受未正式定义的时间格式。RFC3339Nano格式从秒字段中去除尾随的零，因此一旦格式化可能无法正确排序。

​	大多数程序可以使用定义的常量之一作为传递给`Format`或`Parse`的布局。除非您正在创建自定义布局字符串，否则可以忽略此注释的其余部分。

​	为了定义自己的格式，请按照您的方式编写参考时间的格式；参考ANSIC、StampMicro或Kitchen等常量的值以获取示例。该模型旨在演示参考时间的样式，以便`Format`和`Parse`方法可以将相同的转换应用于一般时间值。

​	以下是布局字符串的组成部分的摘要。每个元素都通过示例显示了参考时间的格式化方式。只有这些值会被识别。布局字符串中未被识别为参考时间部分的文本在`Format`期间会被原样回显，并且期望在`Parse`的输入中原样出现。

```
Year: "2006" "06"
Month: "Jan" "January" "01" "1"
Day of the week: "Mon" "Monday"
Day of the month: "2" "_2" "02"
Day of the year: "__2" "002"
Hour: "15" "3" "03" (PM or AM)
Minute: "4" "04"
Second: "5" "05"
AM/PM mark: "PM"
```

​	数字时区偏移的格式如下：

```
"-0700"     ±hhmm
"-07:00"    ±hh:mm
"-07"       ±hh
"-070000"   ±hhmmss
"-07:00:00" ±hh:mm:ss
```

​	用 Z 替换格式中的符号会触发 ISO 8601 的行为，打印 Z 代替 UTC 时区的偏移。因此：

​	用Z替换格式中的符号会触发ISO 8601的行为，将Z打印为UTC时区的偏移量。因此：

```
"Z0700"      Z or ±hhmm
"Z07:00"     Z or ±hh:mm
"Z07"        Z or ±hh
"Z070000"    Z or ±hhmmss
"Z07:00:00"  Z or ±hh:mm:ss
```

​	在格式字符串中，"`_2`" 和 "`__2`"中的下划线表示空格，如果后面的数字有多个数字，则可以替换为空格，以与固定宽度的Unix时间格式兼容。前导零表示零填充值。

​	格式 `__2` 和 `002` 是以空格填充和以零填充的三位数年份的格式；没有未填充的年份格式。

​	逗号或小数点后面跟着一个或多个零表示小数秒，打印到指定的小数位数。逗号或小数点后面跟着一个或多个九表示小数秒，打印到指定的小数位数，并去掉尾随的零。例如，"15:04:05,000" 或"15:04:05.000"格式可以用毫秒精度进行打印或解析。

​	一些有效的布局对于 `time.Parse` 来说是无效的时间值，原因是格式中使用了 `_` 进行空格填充和 Z 进行时区信息的表示。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/time/time.go;l=631)

``` go 
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
```

​	常见的时间段。为了避免在夏令时区转换中产生混淆，没有对Day （天）或更长时间单位进行定义。

​	要计算Duration （持续时间）中的单位数量，请进行除法运算：

```go
second := time.Second
fmt.Print(int64(second/time.Millisecond)) // prints 1000
```

​	要将整数单位转换为 Duration，请使用乘法：

```go
seconds := 10
fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
```

## 变量

This section is empty.

## 函数

#### func After 

``` go 
func After(d Duration) <-chan Time
```

​	After函数等待指定的时间间隔过去后，将当前时间发送到返回的通道上。它等效于`NewTimer(d).C`。底层的计时器直到计时器触发后才会被垃圾回收器回收。如果效率是一个问题，应该使用`NewTimer`，并在不再需要计时器时调用Timer.Stop。

##### After Example
``` go 
package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) {}

func main() {
	fmt.Println("before:", time.Now())
	select {
	case m := <-c:
		handle(m)
	case t, ok := <-time.After(2 * time.Second):
		if ok {
			fmt.Println("timed out:", t)
		}
	}
}

// Output:
//before: 2023-10-22 20:02:07.1865854 +0800 CST m=+0.006234601
//timed out: 2023-10-22 20:02:09.2169465 +0800 CST m=+2.036595701

```

#### func Sleep 

``` go 
func Sleep(d Duration)
```

​	Sleep函数暂停当前 goroutine 至少持续时间 d。负或零持续时间会使 Sleep 立即返回。

##### Sleep Example
``` go 
package main

import (
	"time"
)

func main() {
	time.Sleep(100 * time.Millisecond)
}

```

#### func Tick 

``` go 
func Tick(d Duration) <-chan Time
```

​	Tick是对NewTicker的方便封装，仅提供对滴答通道的访问。尽管Tick对于不需要关闭Ticker的客户端很有用，但请注意，如果没有关闭它的方法，底层的Ticker将无法被垃圾回收器回收；它会"泄漏"。与NewTicker不同，如果d <= 0，Tick将返回`nil`。

##### Tick Example
``` go 
package main

import (
	"fmt"
	"time"
)

func statusUpdate() string { return "" }

func main() {
	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}
}

```

```go
package main

import (
	"fmt"
	"time"
)

func statusUpdate() string { return "" }

func main() {
	c1 := time.Tick(0)
	fmt.Printf("c1=%v\n", c1)
	time.Sleep(2 * time.Second)
	fmt.Println("slept 2 seconds")
	fmt.Printf("c1=%v\n", c1)
	c2 := time.Tick(-1)
	fmt.Printf("c2=%v\n", c2)
	time.Sleep(2 * time.Second)
	fmt.Println("slept 2 seconds")
	fmt.Printf("c2=%v\n", c2)
}

// Output:
//c1=<nil>
//slept 2 seconds
//c1=<nil>
//c2=<nil>
//slept 2 seconds
//c2=<nil>

```



## 类型

### type Duration 

``` go 
type Duration int64
```

​	Duration类型表示两个时间点之间经过的时间，以 int64 纳秒计数的方式表示。该表示方式将最大可表示的持续时间限制在大约 `290` 年左右。

##### Example
``` go 
package main

import (
	"fmt"
	"time"
)

func expensiveCall() {}

func main() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

```

#### func ParseDuration 

``` go 
func ParseDuration(s string) (Duration, error)
```

​	ParseDuration 函数解析一个持续时间字符串。持续时间字符串是一个可能带有符号的十进制数序列，每个数字都可以有小数部分和单位后缀，例如 "300ms"、"-1.5h" 或 "2h45m"。有效的时间单位有 "ns"、"us"(或 "µs")、"ms"、"s"、"m" 和 "h"。

##### ParseDuration Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")
	// The package also accepts the incorrect but common prefix u for micro.
	micro2, _ := time.ParseDuration("1us")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("There are %.0f seconds in %v.\n", complex.Seconds(), complex)
	fmt.Printf("There are %d nanoseconds in %v.\n", micro.Nanoseconds(), micro)
	fmt.Printf("There are %6.2e seconds in %v.\n", micro2.Seconds(), micro)
}
Output:

10h0m0s
1h10m10s
There are 4210 seconds in 1h10m10s.
There are 1000 nanoseconds in 1µs.
There are 1.00e-06 seconds in 1µs.
```

#### func Since 

``` go 
func Since(t Time) Duration
```

​	Since 函数返回自 t 以来经过的时间。它相当于 time.Now().Sub(t)。

#### func Until  <- go1.8

``` go 
func Until(t Time) Duration
```

​	Until 函数返回直到 t 的时间间隔。它相当于 t.Sub(time.Now())。

#### (Duration) Abs  <- go1.19

``` go 
func (d Duration) Abs() Duration
```

​	Abs方法返回 d 的绝对值。作为特殊情况，将 math.MinInt64 转换为 math.MaxInt64。

#### (Duration) Hours 

``` go 
func (d Duration) Hours() float64
```

​	Hours 方法返回持续时间作为小时数的浮点数。

##### Hours Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	h, _ := time.ParseDuration("4h30m")
	fmt.Printf("I've got %.1f hours of work left.", h.Hours())
}
Output:

I've got 4.5 hours of work left.
```

#### (Duration) Microseconds  <- go1.13

``` go 
func (d Duration) Microseconds() int64
```

​	Microseconds方法返回以整数微秒计算的持续时间。

##### Microseconds Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d microseconds.\n", u.Microseconds())
}
Output:

One second is 1000000 microseconds.
```

#### (Duration) Milliseconds  <- go1.13

``` go 
func (d Duration) Milliseconds() int64
```

​	Milliseconds方法返回以整数毫秒计算的持续时间。

##### Milliseconds Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d milliseconds.\n", u.Milliseconds())
}
Output:

One second is 1000 milliseconds.
```

#### (Duration) Minutes 

``` go 
func (d Duration) Minutes() float64
```

​	Minutes方法返回浮点数表示的分钟数持续时间。

##### Minutes Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	m, _ := time.ParseDuration("1h30m")
	fmt.Printf("The movie is %.0f minutes long.", m.Minutes())
}
Output:

The movie is 90 minutes long.
```

#### (Duration) Nanoseconds 

``` go 
func (d Duration) Nanoseconds() int64
```

​	Nanoseconds方法返回以整数纳秒计算的持续时间。

##### Nanoseconds Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	u, _ := time.ParseDuration("1µs")
	fmt.Printf("One microsecond is %d nanoseconds.\n", u.Nanoseconds())
}
Output:

One microsecond is 1000 nanoseconds.
```

#### (Duration) Round  <- go1.9

``` go 
func (d Duration) Round(m Duration) Duration
```

​	Round方法返回将d四舍五入为m的最近倍数的结果。如果结果超过可以存储在Duration中的最大(或最小)值，则Round返回最大(或最小)持续时间。如果m<=0，则Round返回未更改的d。

##### Round Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range round {
		fmt.Printf("d.Round(%6s) = %s\n", r, d.Round(r).String())
	}
}
Output:

d.Round(   1ns) = 1h15m30.918273645s
d.Round(   1µs) = 1h15m30.918274s
d.Round(   1ms) = 1h15m30.918s
d.Round(    1s) = 1h15m31s
d.Round(    2s) = 1h15m30s
d.Round(  1m0s) = 1h16m0s
d.Round( 10m0s) = 1h20m0s
d.Round(1h0m0s) = 1h0m0s
```

#### (Duration) Seconds 

``` go 
func (d Duration) Seconds() float64
```

​	Seconds方法返回浮点数表示的秒数持续时间。

##### Seconds Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	m, _ := time.ParseDuration("1m30s")
	fmt.Printf("Take off in t-%.0f seconds.", m.Seconds())
}
Output:

Take off in t-90 seconds.
```

#### (Duration) String 

``` go 
func (d Duration) String() string
```

​	String方法以 "72h3m0.5s" 的形式返回持续时间的字符串表示形式。省略前导零的单位。作为特殊情况，持续时间小于一秒的格式使用更小的单位(毫秒、微秒或纳秒)，以确保前导数字不为零。零持续时间格式化为0s。

##### String Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(1*time.Hour + 2*time.Minute + 300*time.Millisecond)
	fmt.Println(300 * time.Millisecond)
}
Output:

1h2m0.3s
300ms
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	seconds := 10
	fmt.Println(time.Duration(seconds) * time.Second)            // 10s
	fmt.Println((time.Duration(seconds) * time.Second).String()) // 10s

	seconds = -10
	fmt.Println(time.Duration(seconds) * time.Second)            // -10s
	fmt.Println((time.Duration(seconds) * time.Second).String()) // -10s
}

// Output:
//10s
//10s
//-10s
//-10s

```



#### (Duration) Truncate  <- go1.9

``` go 
func (d Duration) Truncate(m Duration) Duration
```

​	Truncate方法将 d 向零舍入为 m 的倍数并返回结果。如果 m <= 0，则 Truncate 返回未经更改的 d。

##### Truncate Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, t := range trunc {
		fmt.Printf("d.Truncate(%6s) = %s\n", t, d.Truncate(t).String())
	}
}
Output:

d.Truncate(   1ns) = 1h15m30.918273645s
d.Truncate(   1µs) = 1h15m30.918273s
d.Truncate(   1ms) = 1h15m30.918s
d.Truncate(    1s) = 1h15m30s
d.Truncate(    2s) = 1h15m30s
d.Truncate(  1m0s) = 1h15m0s
d.Truncate( 10m0s) = 1h10m0s
d.Truncate(1h0m0s) = 1h0m0s
```

### type Location 

``` go 
type Location struct {
	// contains filtered or unexported fields
}
```

​	Location结构体将时间时刻映射到使用的时区。通常，Location 表示在地理区域中使用的时间偏移集合。对于许多 Location，时间偏移量取决于在时间时刻是否使用夏令时。

##### Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// If the system has a timezone database present, it's possible to load a location
	// from that, e.g.:
	//    newYork, err := time.LoadLocation("America/New_York")

	// Creating a time requires a location. Common locations are time.Local and time.UTC.
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)

	// Although the UTC clock time is 1200 and the Beijing clock time is 2000, Beijing is
	// 8 hours ahead so the two dates actually represent the same instant.
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

}
Output:

true
```

``` go 
var Local *Location = &localLoc
```

​	Local表示系统的本地时区。在 Unix 系统上，Local 会查询 TZ 环境变量以找到要使用的时区。没有 TZ 意味着使用系统默认值 /etc/localtime。TZ="" 表示使用 UTC。TZ="foo" 表示使用系统时区目录中的文件 foo。

``` go 
var UTC *Location = &utcLoc
```

​	UTC 表示协调世界时 (UTC)。

#### func FixedZone 

``` go 
func FixedZone(name string, offset int) *Location
```

​	FixedZone函数返回始终使用给定区域名称和偏移量(相对于 UTC 的秒数)的 Location。

##### FixedZone Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	loc := time.FixedZone("UTC-8", -8*60*60)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	fmt.Println("The time is:", t.Format(time.RFC822))
}
Output:

The time is: 10 Nov 09 23:00 UTC-8
```

#### func LoadLocation 

``` go 
func LoadLocation(name string) (*Location, error)
```

​	LoadLocation函数返回具有给定名称的 Location。

​	如果名称为 "" 或 "UTC"，LoadLocation 返回 UTC。如果名称为 "Local"，LoadLocation 返回 Local。

​	否则，名称被视为对应于 IANA 时区数据库中的文件的位置名称，例如 "America/New_York"。

​	LoadLocation 按顺序在以下位置查找 IANA 时区数据库：

- the directory or uncompressed zip file named by the ZONEINFO environment variable 由 ZONEINFO 环境变量指定的目录或未压缩的 zip 文件 
- 在 Unix 系统上，是系统标准安装位置 
- `$GOROOT/lib/time/zoneinfo.zip` 
- 如果导入了 time/tzdata 包，它将用于查找时区信息。

##### LoadLocation Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}

	timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(timeInUTC.In(location))
}
Output:

2018-08-30 05:00:00 -0700 PDT
```

#### func LoadLocationFromTZData  <- go1.10

``` go 
func LoadLocationFromTZData(name string, data []byte) (*Location, error)
```

​	LoadLocationFromTZData 函数从 IANA 时区数据库格式的数据中返回一个指定名字的 Location。该数据应该采用标准的 IANA 时区文件格式(例如，在 Unix 系统上的 /etc/localtime 内容)。

#### (*Location) String 

``` go 
func (l *Location) String() string
```

​	String 方法返回一个描述性的字符串，表示时区信息，对应于 LoadLocation 或 FixedZone 中的 name 参数。

### type Month 

``` go 
type Month int
```

​	Month 类型表示一年中的月份(1 代表一月，……)。

##### Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	_, month, day := time.Now().Date()
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
}

```

``` go 
const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
```

#### (Month) String 

``` go 
func (m Month) String() string
```

​	String 方法返回该月份的英文名称("January"、"February" 等)。

### type ParseError 

``` go 
type ParseError struct {
	Layout     string
	Value      string
	LayoutElem string
	ValueElem  string
	Message    string
}
```

​	ParseError 类型描述解析时间字符串时的问题。

#### (*ParseError) Error 

``` go 
func (e *ParseError) Error() string
```

​	Error 方法返回 ParseError 的字符串表示形式。

### type Ticker 

``` go 
type Ticker struct {
	C <-chan Time // The channel on which the ticks are delivered.
	r runtimeTimer
}
```

​	Ticker 持有一个通道，该通道以时间间隔提供时钟"滴答声"。

#### func NewTicker 

``` go 
func NewTicker(d Duration) *Ticker
```

​	NewTicker函数返回一个包含通道的新的Ticker。每个tick后，通道将发送当前时间。ticks的周期由duration参数指定。为了弥补接收者的速度慢，ticker将调整时间间隔或者丢弃一些ticks。时间间隔`d`必须大于零；否则，NewTicker将引发panic。**停止ticker以释放相关资源**。

##### NewTicker Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

// Output:
//Current time:  2023-10-22 18:55:09.7151 +0800 CST m=+1.017119601
//Current time:  2023-10-22 18:55:10.710931 +0800 CST m=+2.012950601
//Current time:  2023-10-22 18:55:11.7104022 +0800 CST m=+3.012421801
//Current time:  2023-10-22 18:55:12.7095846 +0800 CST m=+4.011604201
//Current time:  2023-10-22 18:55:13.7084657 +0800 CST m=+5.010485301
//Current time:  2023-10-22 18:55:14.7109187 +0800 CST m=+6.012938301
//Current time:  2023-10-22 18:55:15.7088807 +0800 CST m=+7.010900301
//Current time:  2023-10-22 18:55:16.7066232 +0800 CST m=+8.008642801
//Current time:  2023-10-22 18:55:17.7201499 +0800 CST m=+9.022169501
//Current time:  2023-10-22 18:55:18.7178336 +0800 CST m=+10.019853201
//Done!

```

#### (*Ticker) Reset  <- go1.15

``` go 
func (t *Ticker) Reset(d Duration)
```

​	Reset方法停止一个 Ticker 并将其周期重置为指定的时间间隔。下一个 tick 将在新周期过去后到达。时间间隔`d` 必须大于零；否则，Reset 方法会 panic。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	i := 3
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			i--
			fmt.Println("Current time: ", t)
			if i == 0 {
				ticker.Reset(2 * time.Second)
				fmt.Println("--------After Reset------")
			}
		}
	}
}
// Output:
//Current time:  2023-10-22 19:01:52.641886 +0800 CST m=+1.017594201
//Current time:  2023-10-22 19:01:53.640631 +0800 CST m=+2.016339201
//Current time:  2023-10-22 19:01:54.6396148 +0800 CST m=+3.015323001
//--------After Reset------
//Current time:  2023-10-22 19:01:56.6480269 +0800 CST m=+5.023735101
//Current time:  2023-10-22 19:01:58.6463226 +0800 CST m=+7.022030801
//Current time:  2023-10-22 19:02:00.6546651 +0800 CST m=+9.030373301
//Done!
```



#### (*Ticker) Stop 

``` go 
func (t *Ticker) Stop()
```

​	Stop方法关闭一个 Ticker。Stop 后，将不会再发送任何 tick。Stop 不会关闭通道，以防止并发的 goroutine 从通道中读取到错误的"tick"。

### type Time 

``` go 
type Time struct {
	// contains filtered or unexported fields
}
```

​	Time 表示带有纳秒精度的时间点。

​	使用时间的程序通常应将其存储和传递为值，而不是指针。也就是说，时间变量和结构字段应为 time.Time 类型，而不是 *time.Time。

​	Time 值可以被多个 goroutine 同时使用，除了 GobDecode、UnmarshalBinary、UnmarshalJSON 和 UnmarshalText 方法之外。这些方法不支持并发安全性。

​	时间点可以使用 Before、After 和 Equal 方法进行比较。Sub 方法减去两个时间点，产生一个 Duration。Add 方法将一个时间点和一个 Duration 相加，产生一个时间点。

​	Time 类型的零值是 UTC 时间的 1 年 1 月 1 日 00:00:00.000000000。由于实际使用中不太可能遇到这个时间，IsZero 方法提供了一种简单的方式来检测未明确初始化的时间。

​	每个 Time 都有一个关联的 Location，在计算时间的表示形式时进行查询，例如在 Format、Hour 和 Year 方法中。Local、UTC 和 In 方法返回具有特定位置的时间。以这种方式更改位置只更改表示形式；它不会更改所表示的时间点，因此不会影响前面段落中描述的计算。

​	一个保存了 GobEncode，MarshalBinary，MarshalJSON 和 MarshalText 方法的 Time 值的表示形式存储了 Time.Location 的偏移量，但没有存储位置名称，因此它们会丢失关于夏令时的信息。

​	除了必须的"挂钟"读数之外，一个 Time 可能还包含当前进程单调钟的可选读数，以提供比较或减法的额外精度。有关详细信息，请参见包文档中的"单调钟"部分。

​	请注意，Go 的 == 操作符不仅比较时间瞬间，还比较 Location 和单调钟读数。因此，在未保证所有值都已设置相同的 Location 之前，不应将 Time 值用作映射或数据库键，可以通过使用 UTC 或 Local 方法来实现，同时剥离单调钟读数，设置 t = t.Round(0)。通常，优先使用 t.Equal(u) 而不是 t == u，因为 t.Equal 使用可用的最准确的比较，并正确处理仅有一个参数具有单调钟读数的情况。

#### func Date 

``` go 
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
```

​	Date函数返回与给定位置中该时间的适当区域内的"yyyy-mm-dd hh:mm:ss + nsec nanoseconds"相对应的 Time。

​	月份，日期，小时，分钟，秒和纳秒的值可能超出其通常范围，并在转换期间进行归一化。例如，10 月 32 日转换为 11 月 1 日。

​	夏令时转换会跳过或重复时间。例如，在美国，2011 年 3 月 13 日 2:15 上午从未发生过，而在 2011 年 11 月 6 日 1:15 上午则发生了两次。在这种情况下，时区和因此时间的选择并不明确。Date 返回一个在转换涉及的两个区域中一个区域内正确的时间，但不保证是哪一个。

​	如果 loc 为 nil，Date 将 panic。

##### Date Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}
Output:

Go launched at 2009-11-10 15:00:00 -0800 PST
```

#### func Now 

``` go 
func Now() Time
```

​	Now函数返回当前本地时间。

#### func Parse 

``` go 
func Parse(layout, value string) (Time, error)
```

​	Parse函数解析格式化的字符串并返回其代表的时间值。查看名为 Layout 的常量的文档以了解如何表示格式。第二个参数必须可用作第一个参数提供的格式字符串(layout)进行解析。

​	Time.Format 的示例详细演示了布局字符串的工作原理，并且是一个很好的参考。

​	在解析时，输入可能包含紧随秒字段后面的小数秒字段，即使布局没有表示它的存在也是如此。在这种情况下，逗号或点号后跟最大系列的数字将被解析为小数秒。小数秒将被截断为纳秒精度。

​	布局中省略的元素被假定为零或(当零不可能时)为一，因此解析"3:04pm"会返回对应于 0 年 1 月 1 日 15:04:00 UTC 的时间(请注意，因为年份是 0，所以此时间在零时间之前)。年份必须在 0000..9999 范围内。星期几的语法被检查，但是被忽略。

​	对于指定两位数年份 06 的布局，NN >= 69 的值将被视为 19NN，NN < 69 的值将被视为 20NN。

​	此注释的其余部分描述了时区的处理方式。

​	如果没有时区指示符，则 Parse 返回 UTC 时间。

​	当解析具有类似于 -0700 的区域偏移量的时间时，如果该偏移量对应于当前位置(Local)使用的时区，则 Parse 使用该位置和区域在返回的时间中。否则，它会将时间记录为处于一个虚构的位置，该位置的时间固定在给定的区域偏移量上。

​	当解析具有类似于 MST 的区域缩写的时间时，如果该区域缩写在当前位置具有定义的偏移量，则使用该偏移量。无论位置如何，"UTC"区域缩写都被认为是 UTC。如果区域缩写未知，则 Parse 记录该时间处于具有给定区域缩写和零偏移量的虚构位置。此选择意味着可以使用相同的布局无损地解析和重新格式化此类时间，但是表示中使用的确切时刻将与实际区域偏移量不同。为避免此类问题，请使用使用数字区域偏移量的时间布局，或使用 ParseInLocation。

##### Parse Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	// See the example for Time.Format for a thorough description of how
	// to define the layout string to parse a time.Time value; Parse and
	// Format use the same model to describe their input and output.

	// longForm shows by example how the reference time would be represented in
	// the desired layout.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	// Some valid layouts are invalid time values, due to format specifiers
	// such as _ for space padding and Z for zone information.
	// For example the RFC3339 layout 2006-01-02T15:04:05Z07:00
	// contains both Z and a time zone offset in order to handle both valid options:
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t)
	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err) // Returns an error as the layout is not a valid time value

}
Output:

2013-02-03 19:54:00 -0800 PST
2013-02-03 00:00:00 +0000 UTC
2006-01-02 15:04:05 +0000 UTC
2006-01-02 15:04:05 +0700 +0700
error parsing time "2006-01-02T15:04:05Z07:00": extra text: "07:00"
```

#### func ParseInLocation  <- go1.1

``` go 
func ParseInLocation(layout, value string, loc *Location) (Time, error)
```

​	ParseInLocation函数与Parse函数类似，但有两个重要区别。首先，在缺少时区信息的情况下，Parse将一个时间解释为UTC；而ParseInLocation则将时间解释为给定时区的时间。其次，当给定时区偏移量或缩写时，Parse会尝试将其与本地时区匹配；而ParseInLocation使用给定时区。

##### ParseInLocation Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Europe/Berlin")

	// This will look for the name CEST in the Europe/Berlin time zone.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)

	// Note: without explicit zone, returns time in given location.
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)

}
Output:

2012-07-09 05:02:00 +0200 CEST
2012-07-09 00:00:00 +0200 CEST
```

#### func Unix 

``` go 
func Unix(sec int64, nsec int64) Time
```

​	Unix函数返回与自1970年1月1日UTC以来的sec秒和nsec纳秒对应的本地时间。传递nsec超出范围[0, 999999999]是有效的。并非所有的sec值都有相应的时间值。其中一个这样的值是1<<63-1(最大的int64值)。



##### Unix Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	unixTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(unixTime.Unix())
	t := time.Unix(unixTime.Unix(), 0).UTC()
	fmt.Println(t)

}
Output:

1257894000
2009-11-10 23:00:00 +0000 UTC
```

#### func UnixMicro  <- go1.17

``` go 
func UnixMicro(usec int64) Time
```

​	UnixMicro函数返回与自1970年1月1日UTC以来的usec微秒对应的本地时间。

##### UnixMicro Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	umt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(umt.UnixMicro())
	t := time.UnixMicro(umt.UnixMicro()).UTC()
	fmt.Println(t)

}
Output:

1257894000000000
2009-11-10 23:00:00 +0000 UTC
```

#### func UnixMilli  <- go1.17

``` go 
func UnixMilli(msec int64) Time
```

​	UnixMilli函数返回与自1970年1月1日UTC以来的msec毫秒对应的本地时间。

##### UnixMilli Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	umt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(umt.UnixMilli())
	t := time.UnixMilli(umt.UnixMilli()).UTC()
	fmt.Println(t)

}
Output:

1257894000000
2009-11-10 23:00:00 +0000 UTC
```

#### (Time) Add 

``` go 
func (t Time) Add(d Duration) Time
```

​	Add方法返回t+d时间。

##### Add Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenSeconds := start.Add(time.Second * 10)
	afterTenMinutes := start.Add(time.Minute * 10)
	afterTenHours := start.Add(time.Hour * 10)
	afterTenDays := start.Add(time.Hour * 24 * 10)

	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds)
	fmt.Printf("start.Add(time.Minute * 10) = %v\n", afterTenMinutes)
	fmt.Printf("start.Add(time.Hour * 10) = %v\n", afterTenHours)
	fmt.Printf("start.Add(time.Hour * 24 * 10) = %v\n", afterTenDays)

}
Output:

start = 2009-01-01 12:00:00 +0000 UTC
start.Add(time.Second * 10) = 2009-01-01 12:00:10 +0000 UTC
start.Add(time.Minute * 10) = 2009-01-01 12:10:00 +0000 UTC
start.Add(time.Hour * 10) = 2009-01-01 22:00:00 +0000 UTC
start.Add(time.Hour * 24 * 10) = 2009-01-11 12:00:00 +0000 UTC
```

#### (Time) AddDate 

``` go 
func (t Time) AddDate(years int, months int, days int) Time
```

​	AddDate方法返回添加指定年、月、日后的时间 t。例如，将 AddDate(-1, 2, 3) 应用于 2011 年 1 月 1 日，会返回 2010 年 3 月 4 日。

​	AddDate方法会像 Date 方法一样对其结果进行规范化，因此例如将 10 月 31 日加一月会得到规范化后的 11 月 31 日，即 12 月 1 日。

##### AddDate Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)
	oneDayLater := start.AddDate(0, 0, 1)
	oneMonthLater := start.AddDate(0, 1, 0)
	oneYearLater := start.AddDate(1, 0, 0)

	fmt.Printf("oneDayLater: start.AddDate(0, 0, 1) = %v\n", oneDayLater)
	fmt.Printf("oneMonthLater: start.AddDate(0, 1, 0) = %v\n", oneMonthLater)
	fmt.Printf("oneYearLater: start.AddDate(1, 0, 0) = %v\n", oneYearLater)

}
Output:

oneDayLater: start.AddDate(0, 0, 1) = 2009-01-02 00:00:00 +0000 UTC
oneMonthLater: start.AddDate(0, 1, 0) = 2009-02-01 00:00:00 +0000 UTC
oneYearLater: start.AddDate(1, 0, 0) = 2010-01-01 00:00:00 +0000 UTC
```

#### (Time) After 

``` go 
func (t Time) After(u Time) bool
```

​	After方法报告时间点 t 是否在 u 之后。

##### After Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear3000AfterYear2000 := year3000.After(year2000) // True
	isYear2000AfterYear3000 := year2000.After(year3000) // False

	fmt.Printf("year3000.After(year2000) = %v\n", isYear3000AfterYear2000)
	fmt.Printf("year2000.After(year3000) = %v\n", isYear2000AfterYear3000)

}
Output:

year3000.After(year2000) = true
year2000.After(year3000) = false
```

#### (Time) AppendFormat  <- go1.5

``` go 
func (t Time) AppendFormat(b []byte, layout string) []byte
```

​	AppendFormat方法与 Format方法类似，但将文本表示附加到 b 中并返回扩展缓冲区。

##### AppendFormat Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2017, time.November, 4, 11, 0, 0, 0, time.UTC)
	text := []byte("Time: ")

	text = t.AppendFormat(text, time.Kitchen)
	fmt.Println(string(text))

}
Output:

Time: 11:00AM
```

#### (Time) Before 

``` go 
func (t Time) Before(u Time) bool
```

​	Before方法报告时间点 t 是否在 u 之前。

##### Before Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear2000BeforeYear3000 := year2000.Before(year3000) // True
	isYear3000BeforeYear2000 := year3000.Before(year2000) // False

	fmt.Printf("year2000.Before(year3000) = %v\n", isYear2000BeforeYear3000)
	fmt.Printf("year3000.Before(year2000) = %v\n", isYear3000BeforeYear2000)

}
Output:

year2000.Before(year3000) = true
year3000.Before(year2000) = false
```

#### (Time) Clock 

``` go 
func (t Time) Clock() (hour, min, sec int)
```

​	Clock方法返回指定 t 的所在日的小时、分钟和秒。

#### (Time) Compare  <- go1.20

``` go 
func (t Time) Compare(u Time) int
```

​	Compare方法比较时间点 t 和 u。如果 t 在 u 之前，则返回 -1；如果 t 在 u 之后，则返回 +1；如果它们相同，则返回 0。

#### (Time) Date 

``` go 
func (t Time) Date() (year int, month Month, day int)
```

​	Date方法返回 t 所代表的年、月、日。

##### Date Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()

	fmt.Printf("year = %v\n", year)
	fmt.Printf("month = %v\n", month)
	fmt.Printf("day = %v\n", day)

}
Output:

year = 2000
month = February
day = 1
```

#### (Time) Day 

``` go 
func (t Time) Day() int
```

​	Day方法返回 t 所代表的月份中的第几天。

##### Day Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	day := d.Day()

	fmt.Printf("day = %v\n", day)

}
Output:

day = 1
```

#### (Time) Equal 

``` go 
func (t Time) Equal(u Time) bool
```

​	Equal方法报告 t 和 u 是否代表同一时刻。即使两个时间处于不同的时区，它们也可以相等。例如，6:00 +0200 和 4:00 UTC 是相等的。有关使用 == 与 Time 值时遇到的问题，请参阅 Time 类型的文档；大多数代码应该使用 Equal。

##### Equal Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// Unlike the equal operator, Equal is aware that d1 and d2 are the
	// same instant but in different time zones.
	d1 := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	d2 := time.Date(2000, 2, 1, 20, 30, 0, 0, beijing)

	datesEqualUsingEqualOperator := d1 == d2
	datesEqualUsingFunction := d1.Equal(d2)

	fmt.Printf("datesEqualUsingEqualOperator = %v\n", datesEqualUsingEqualOperator)
	fmt.Printf("datesEqualUsingFunction = %v\n", datesEqualUsingFunction)

}
Output:

datesEqualUsingEqualOperator = false
datesEqualUsingFunction = true
```

#### (Time) Format 

``` go 
func (t Time) Format(layout string) string
```

​	Format方法根据提供的格式返回 t 的文本表示。有关如何表示布局格式的信息，请参阅名为 Layout 的常量的文档。

​	Time.Format 的可执行示例详细演示了布局字符串的工作方式，是一个很好的参考。

##### Format Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	// Parse a time value from a string in the standard Unix format.
    // 从标准Unix格式的字符串中解析时间值。
	t, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil { // Always check errors even if they should not happen. 始终检查错误，即使它们不应该发生。
		panic(err)
	}

	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil { // Always check errors even if they should not happen. 始终检查错误，即使它们不应该发生。
		panic(err)
	}

	// time.Time's Stringer method is useful without any format.
    // time.Time的Stringer方法在没有任何格式的情况下非常有用。
	fmt.Println("default format:", t)

	// Predefined constants in the package implement common layouts.
    // 该包中的预定义常量实现了常见的布局。
	fmt.Println("Unix format:", t.Format(time.UnixDate))

	// The time zone attached to the time value affects its output.
	fmt.Println("Same, in UTC:", t.UTC().Format(time.UnixDate))

	fmt.Println("in Shanghai with seconds:", t.In(tz).Format("2006-01-02T15:04:05 -070000"))

	fmt.Println("in Shanghai with colon seconds:", t.In(tz).Format("2006-01-02T15:04:05 -07:00:00"))

	// The rest of this function demonstrates the properties of the
	// layout string used in the format.

	// The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
	// There are some wrinkles illustrated below.

	// Most uses of Format and Parse use constant layout strings such as
	// the ones defined in this package, but the interface is flexible,
	// as these examples show.

	// Define a helper function to make the examples' output look nice.
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}

	// Print a header in our output.
	fmt.Printf("\nFormats:\n\n")

	// Simple starter examples.
	do("Basic full date", "Mon Jan 2 15:04:05 MST 2006", "Wed Feb 25 11:06:39 PST 2015")
	do("Basic short date", "2006/01/02", "2015/02/25")

	// The hour of the reference time is 15, or 3PM. The layout can express
	// it either way, and since our value is the morning we should see it as
	// an AM time. We show both in one format string. Lower case too.
	do("AM/PM", "3PM==3pm==15h", "11AM==11am==11h")

	// When parsing, if the seconds value is followed by a decimal point
	// and some digits, that is taken as a fraction of a second even if
	// the layout string does not represent the fractional second.
	// Here we add a fractional second to our time value used above.
	t, err = time.Parse(time.UnixDate, "Wed Feb 25 11:06:39.1234 PST 2015")
	if err != nil {
		panic(err)
	}
	// It does not appear in the output if the layout string does not contain
	// a representation of the fractional second.
	do("No fraction", time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")

	// Fractional seconds can be printed by adding a run of 0s or 9s after
	// a decimal point in the seconds value in the layout string.
	// If the layout digits are 0s, the fractional second is of the specified
	// width. Note that the output has a trailing zero.
	do("0s for fraction", "15:04:05.00000", "11:06:39.12340")

	// If the fraction in the layout is 9s, trailing zeros are dropped.
	do("9s for fraction", "15:04:05.99999999", "11:06:39.1234")

}
Output:

default format: 2015-02-25 11:06:39 -0800 PST
Unix format: Wed Feb 25 11:06:39 PST 2015
Same, in UTC: Wed Feb 25 19:06:39 UTC 2015
in Shanghai with seconds: 2015-02-26T03:06:39 +080000
in Shanghai with colon seconds: 2015-02-26T03:06:39 +08:00:00

Formats:

Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
Basic short date "2006/01/02" gives "2015/02/25"
AM/PM            "3PM==3pm==15h" gives "11AM==11am==11h"
No fraction      "Mon Jan _2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
0s for fraction  "15:04:05.00000" gives "11:06:39.12340"
9s for fraction  "15:04:05.99999999" gives "11:06:39.1234"
```

##### Format Example(Pad)
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	// Parse a time value from a string in the standard Unix format.
	t, err := time.Parse(time.UnixDate, "Sat Mar 7 11:06:39 PST 2015")
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}

	// Define a helper function to make the examples' output look nice.
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}

	// The predefined constant Unix uses an underscore to pad the day.
	do("Unix", time.UnixDate, "Sat Mar  7 11:06:39 PST 2015")

	// For fixed-width printing of values, such as the date, that may be one or
	// two characters (7 vs. 07), use an _ instead of a space in the layout string.
	// Here we print just the day, which is 2 in our layout string and 7 in our
	// value.
	do("No pad", "<2>", "<7>")

	// An underscore represents a space pad, if the date only has one digit.
	do("Spaces", "<_2>", "< 7>")

	// A "0" indicates zero padding for single-digit values.
	do("Zeros", "<02>", "<07>")

	// If the value is already the right width, padding is not used.
	// For instance, the second (05 in the reference time) in our value is 39,
	// so it doesn't need padding, but the minutes (04, 06) does.
	do("Suppressed pad", "04:05", "06:39")

}
Output:

Unix             "Mon Jan _2 15:04:05 MST 2006" gives "Sat Mar  7 11:06:39 PST 2015"
No pad           "<2>" gives "<7>"
Spaces           "<_2>" gives "< 7>"
Zeros            "<02>" gives "<07>"
Suppressed pad   "04:05" gives "06:39"
```

#### (Time) GoString  <- go1.17

``` go 
func (t Time) GoString() string
```

​	GoString方法实现了 fmt.GoStringer 接口，并将 t 格式化为可打印的 Go 源代码。

##### Example
``` go 
```

#### (*Time) GobDecode 

``` go 
func (t *Time) GobDecode(data []byte) error
```

​	GobDecode方法实现了 gob.GobDecoder 接口。

#### (Time) GobEncode 

``` go 
func (t Time) GobEncode() ([]byte, error)
```

​	GobEncode方法实现了 gob.GobEncoder 接口。

#### (Time) Hour 

``` go 
func (t Time) Hour() int
```

​	Hour方法返回 t 所在的一天中的小时数，范围为 [0, 23]。

#### (Time) ISOWeek 

``` go 
func (t Time) ISOWeek() (year, week int)
```

​	ISOWeek方法返回 t 所在的 ISO 8601 年份和周数。周数从 1 到 53。一月 1 日到 3 日属于年 n 的第 52 或 53 周，而 12 月 29 日到 31 日可能属于年 n+1 的第 1 周。

#### (Time) In 

``` go 
func (t Time) In(loc *Location) Time
```

​	In方法返回表示与 t 相同时间点的副本，但将副本的位置信息设置为 loc 以供显示。

​	如果 loc 为 nil，则会引发 panic。

#### (Time) IsDST  <- go1.17

``` go 
func (t Time) IsDST() bool
```

​	IsDST方法报告配置位置中的时间是否处于夏令时。

#### (Time) IsZero 

``` go 
func (t Time) IsZero() bool
```

​	IsZero方法报告t是否表示零时刻，即UTC时间的January 1, year 1, 00:00:00.

#### (Time) Local 

``` go 
func (t Time) Local() Time
```

​	Local方法返回t对应的本地时间。

#### (Time) Location 

``` go 
func (t Time) Location() *Location
```

​	Location方法返回t对应的时区信息。

#### (Time) MarshalBinary  <- go1.2

``` go 
func (t Time) MarshalBinary() ([]byte, error)
```

​	MarshalBinary方法实现encoding.BinaryMarshaler接口。

#### (Time) MarshalJSON 

``` go 
func (t Time) MarshalJSON() ([]byte, error)
```

​	MarshalJSON方法实现json.Marshaler接口。时间使用带有亚秒精度的[RFC 3339](https://rfc-editor.org/rfc/rfc3339.html)格式的带引号的字符串表示。如果时间戳无法表示为有效的[RFC 3339](https://rfc-editor.org/rfc/rfc3339.html)格式(例如，年份超出范围)，则报告错误。

#### (Time) MarshalText  <- go1.2

``` go 
func (t Time) MarshalText() ([]byte, error)
```

​	MarshalText方法实现encoding.TextMarshaler接口。时间使用带有亚秒精度的[RFC 3339](https://rfc-editor.org/rfc/rfc3339.html)格式表示。如果时间戳无法表示为有效的[RFC 3339](https://rfc-editor.org/rfc/rfc3339.html)格式(例如，年份超出范围)，则报告错误。

#### (Time) Minute 

``` go 
func (t Time) Minute() int
```

​	Minute方法返回t对应的小时内分钟数，范围为[0, 59]。

#### (Time) Month 

``` go 
func (t Time) Month() Month
```

​	Month方法返回t对应的月份。

#### (Time) Nanosecond 

``` go 
func (t Time) Nanosecond() int
```

​	Nanosecond方法返回t对应的秒内纳秒数，范围为[0, 999999999]。

#### (Time) Round  <- go1.1

``` go 
func (t Time) Round(d Duration) Time
```

​	Round方法返回将t四舍五入到离d的最近倍数后的结果(自零时间开始)。半数值的舍入行为为向上舍入。如果d<=0，则Round返回除单调钟读数以外其他方面都没有改变的t。

​	Round方法以自零时间以来的绝对持续时间为基础操作时间；它不以时间的表现形式为基础。因此，Round(Hour)可能会返回具有非零分钟的时间，这取决于时间的Location。

##### Round Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}
}
Output:

t.Round(   1ns) = 12:15:30.918273645
t.Round(   1µs) = 12:15:30.918274
t.Round(   1ms) = 12:15:30.918
t.Round(    1s) = 12:15:31
t.Round(    2s) = 12:15:30
t.Round(  1m0s) = 12:16:00
t.Round( 10m0s) = 12:20:00
t.Round(1h0m0s) = 12:00:00
```

#### (Time) Second 

``` go 
func (t Time) Second() int
```

​	Second方法返回 t 所代表的时间的秒数偏移量，范围在 [0, 59] 之间。

#### (Time) String 

``` go 
func (t Time) String() string
```

String returns the time formatted using the format string

```
"2006-01-02 15:04:05.999999999 -0700 MST"
```

​	String方法返回使用格式字符串 "2006-01-02 15:04:05.999999999 -0700 MST" 格式化的时间。如果时间具有单调钟读数，则返回的字符串包括最终字段 "`m=±<value>`"，其中 value 是单调钟读数，格式化为十进制秒数。

​	返回的字符串仅用于调试；对于稳定的序列化表示，使用 t.MarshalText、t.MarshalBinary 或带有显式格式字符串的 t.Format。

##### String Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	timeWithNanoseconds := time.Date(2000, 2, 1, 12, 13, 14, 15, time.UTC)
	withNanoseconds := timeWithNanoseconds.String()

	timeWithoutNanoseconds := time.Date(2000, 2, 1, 12, 13, 14, 0, time.UTC)
	withoutNanoseconds := timeWithoutNanoseconds.String()

	fmt.Printf("withNanoseconds = %v\n", string(withNanoseconds))
	fmt.Printf("withoutNanoseconds = %v\n", string(withoutNanoseconds))

}
Output:

withNanoseconds = 2000-02-01 12:13:14.000000015 +0000 UTC
withoutNanoseconds = 2000-02-01 12:13:14 +0000 UTC
```

#### (Time) Sub 

``` go 
func (t Time) Sub(u Time) Duration
```

​	Sub方法返回 t-u 的持续时间。如果结果超过可以存储在 Duration 中的最大(或最小)值，则将返回最大(或最小)持续时间。要计算 t-d 的持续时间 d，请使用 t.Add(-d)。

##### Sub Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	difference := end.Sub(start)
	fmt.Printf("difference = %v\n", difference)

}
Output:

difference = 12h0m0s
```

#### (Time) Truncate  <- go1.1

``` go 
func (t Time) Truncate(d Duration) Time
```

​	Truncate方法返回将 t 向下舍入为 d 的倍数的结果(自零时起)。如果 d <= 0，则 Truncate 返回 t 去除任何单调钟读数但其他不变。

​	Truncate方法在绝对时间自零时以来的持续时间上操作；它不会操作时间的表现形式。因此，Truncate(Hour) 可能会返回带有非零分钟的时间，具体取决于时间的位置。

##### Truncate Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")
	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
	}

	for _, d := range trunc {
		fmt.Printf("t.Truncate(%5s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
	}
	// To round to the last midnight in the local timezone, create a new Date.
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	_ = midnight

}
Output:

t.Truncate(  1ns) = 12:15:30.918273645
t.Truncate(  1µs) = 12:15:30.918273
t.Truncate(  1ms) = 12:15:30.918
t.Truncate(   1s) = 12:15:30
t.Truncate(   2s) = 12:15:30
t.Truncate( 1m0s) = 12:15:00
t.Truncate(10m0s) = 12:10:00
```

#### (Time) UTC 

``` go 
func (t Time) UTC() Time
```

​	UTC方法将位置设置为 UTC 并返回 t。

#### (Time) Unix 

``` go 
func (t Time) Unix() int64
```

​	Unix方法返回 t 作为 Unix 时间，即自 1970 年 1 月 1 日 UTC 以来经过的秒数。结果不依赖于与t关联的位置。类 Unix 操作系统通常将时间记录为 32 位秒计数，但由于此方法返回 64 位值，因此对过去或未来数十亿年都有效。

##### Unix Example
``` go 
package main

import (
	"fmt"
	"time"
)

func main() {
	// 1 billion seconds of Unix, three ways.
	fmt.Println(time.Unix(1e9, 0).UTC())     // 1e9 seconds
	fmt.Println(time.Unix(0, 1e18).UTC())    // 1e18 nanoseconds
	fmt.Println(time.Unix(2e9, -1e18).UTC()) // 2e9 seconds - 1e18 nanoseconds

	t := time.Date(2001, time.September, 9, 1, 46, 40, 0, time.UTC)
	fmt.Println(t.Unix())     // seconds since 1970
	fmt.Println(t.UnixNano()) // nanoseconds since 1970

}
Output:

2001-09-09 01:46:40 +0000 UTC
2001-09-09 01:46:40 +0000 UTC
2001-09-09 01:46:40 +0000 UTC
1000000000
1000000000000000000
```

#### (Time) UnixMicro  <- go1.17

``` go 
func (t Time) UnixMicro() int64
```

​	UnixMicro方法将 t 表示为 Unix 时间，即自 1970 年 1 月 1 日 UTC 起经过的微秒数。如果 Unix 时间不能被 int64 表示(即在年份 -290307 之前或 294246 之后)，则结果未定义。结果不取决于与 t 关联的位置。

#### (Time) UnixMilli  <- go1.17

``` go 
func (t Time) UnixMilli() int64
```

​	UnixMilli方法将 t 表示为 Unix 时间，即自 1970 年 1 月 1 日 UTC 起经过的毫秒数。如果 Unix 时间不能被 int64 表示(即在 1970 年之前或之后的 29.2 亿年)，则结果未定义。结果不取决于与 t 关联的位置。

#### (Time) UnixNano 

``` go 
func (t Time) UnixNano() int64
```

​	UnixNano方法将 t 表示为 Unix 时间，即自 1970 年 1 月 1 日 UTC 起经过的纳秒数。如果 Unix 时间不能被 int64 表示(即在 1678 年之前或 2262 年之后的日期)，则结果未定义。请注意，这意味着在零时间上调用 UnixNano 的结果是未定义的。结果不取决于与 t 关联的位置。

#### (*Time) UnmarshalBinary  <- go1.2

``` go 
func (t *Time) UnmarshalBinary(data []byte) error
```

​	UnmarshalBinary方法实现了 encoding.BinaryUnmarshaler 接口。

#### (*Time) UnmarshalJSON 

``` go 
func (t *Time) UnmarshalJSON(data []byte) error
```

​	UnmarshalJSON方法实现了 json.Unmarshaler 接口。时间必须是[RFC 3339](https://rfc-editor.org/rfc/rfc3339.html) 格式的引号字符串。

#### (*Time) UnmarshalText  <- go1.2

``` go 
func (t *Time) UnmarshalText(data []byte) error
```

​	UnmarshalText方法实现了 encoding.TextUnmarshaler 接口。时间必须是 [RFC 3339](https://rfc-editor.org/rfc/rfc3339.html)格式。

#### (Time) Weekday 

``` go 
func (t Time) Weekday() Weekday
```

​	Weekday方法返回 t 指定的星期几。

#### (Time) Year 

``` go 
func (t Time) Year() int
```

​	Year方法返回 t 所在的年份。

#### (Time) YearDay  <- go1.1

``` go 
func (t Time) YearDay() int
```

​	YearDay方法返回 t 指定的年份的天数，在非闰年中为 [1,365] 范围内的数，闰年中为 [1,366]。

#### (Time) Zone 

``` go 
func (t Time) Zone() (name string, offset int)
```

​	Zone 方法计算 t 所在时区，返回时区的缩写名称(例如"CET")和其相对于 UTC 的偏移量(以秒为单位)。

#### (Time) ZoneBounds  <- go1.19

``` go 
func (t Time) ZoneBounds() (start, end Time)
```

​	ZoneBounds方法返回t时刻所在时区的起始时间和终止时间。该时区从start开始，下一个时区从end开始。如果该时区从时间开始时刻开始，则start将作为零值Time返回。如果时区一直持续下去，则end将作为零值Time返回。返回的时间Location与t相同。

### type Timer 

``` go 
type Timer struct {
	C <-chan Time
	r runtimeTimer
}
```

​	Timer类型表示单个事件。当Timer过期时，当前时间将发送到C通道，除非Timer是由AfterFunc创建的。Timer必须使用NewTimer或AfterFunc创建。

#### func AfterFunc 

``` go 
func AfterFunc(d Duration, f func()) *Timer
```

​	AfterFunc函数等待持续时间过去，然后在其自己的goroutine中调用f。它返回一个Timer，可以使用它的Stop方法取消调用。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("before:", time.Now())
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("It's been 5 seconds and Now time is \n", time.Now())
	})
	time.Sleep(10 * time.Second)
}

// Output:
//before: 2023-10-22 08:15:47.1432426 +0800 CST m=+0.005348101
//It's been 5 seconds and Now time is
//2023-10-22 08:15:52.1694514 +0800 CST m=+5.031556901

```



#### func NewTimer 

``` go 
func NewTimer(d Duration) *Timer
```

​	NewTimer函数创建一个新的Timer，它将在至少持续时间d之后在其通道上发送当前时间。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("before:", time.Now())
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("after:", <-timer.C)
}

// Output:
//before: 2023-10-22 08:04:39.0256614 +0800 CST m=+0.005701401
//after: 2023-10-22 08:04:44.0665367 +0800 CST m=+5.046576701
```



#### (*Timer) Reset  <- go1.1

``` go 
func (t *Timer) Reset(d Duration) bool
```

​	Reset方法将计时器更改为在持续时间`d`后到期。如果计时器处于活动状态，则返回`true`，如果计时器已过期或已停止，则返回`false`。 

​	**对于使用NewTimer创建的计时器，Reset只应在已停止或已过期且通道已排空的情况下调用**。 

> 个人注释
>
> ​	Reset只应在（已停止）或（已过期且通道已排空）的情况下调用。
>
> ​	怎么判断计时器已停止？
>
> ​	调用Stop方法返回true，可以说明计时器已停止。

​	如果程序已经从t.C接收到值，则已知该计时器已过期并且通道已排空，因此可以直接使用t.Reset。**但是，如果程序尚未从t.C接收到值（即未排空），则必须停止该计时器**，并且如果Stop报告该计时器在停止之前已过期，则必须显式排空通道。

```go
if !t.Stop() {
	<-t.C
}
t.Reset(d)
```

​	这不应与来自该计时器通道的其他接收同时进行。 

​	请注意，正确使用Reset的返回值是不可能的，因为在清空通道和新计时器到期之间存在竞态条件。**应如上所述始终在已停止或已到期的通道上调用Reset，以保持与现有程序的兼容性**。返回值存在是为了保持与现有程序的兼容性。 

​	对于使用`AfterFunc(d, f)`创建的计时器，Reset方法要么重新安排f的运行时间，此时Reset方法返回`true`，要么安排`f`再次运行，此时Reset方法返回`false`。当Reset方法返回`false`时，Reset方法既不等待之前的`f`完成再返回，也不能保证后续运行`f`的goroutine不与之前的同时运行。如果调用方需要知道之前的`f`执行是否完成，它必须与`f`进行显式的协调。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("--------------------\n timer1 \n --------------------\n")
	fmt.Println("before:", time.Now())
	timer1 := time.AfterFunc(5*time.Second, func() {
		fmt.Println("timer1 has been 5+ seconds and Now time is \n", time.Now())
	})
	time.Sleep(2 * time.Second)
	fmt.Println("had sleep 2 second:", time.Now())

	// 应始终在已停止或到期的通道上调用Reset，以保持与现有程序的兼容性
	if !timer1.Stop() {
		<-timer1.C
	}

	fmt.Println("reset:", time.Now())
	fmt.Printf("Reset timer1 return %t\n", timer1.Reset(5*time.Second))
	time.Sleep(10 * time.Second)

	fmt.Printf("--------------------\n timer2 \n --------------------\n")
	fmt.Println("before:", time.Now())
	timer2 := time.NewTimer(5 * time.Second)
	time.Sleep(2 * time.Second)
	fmt.Println("had sleep 2 second:", time.Now())
	if !timer2.Stop() {
		fmt.Println("执行Stop方法ing...")
		<-timer2.C
		fmt.Println("已执行了Stop方法")
	}
	fmt.Printf("Reset timer2 return %t\n", timer2.Reset(5*time.Second))

	for {
		select {
		case _, ok := <-timer2.C:
			if ok {
				fmt.Println("timer2 has been 5+ seconds and Now time is \n", time.Now())
				goto END
			}
		default:
		}
	}

END:
	fmt.Println("END")
}
// 请编译后在执行，直接使用go run 可能结果会不同！！！
// Output:
//--------------------
//timer1
//--------------------
//before: 2023-10-22 18:21:18.6134724 +0800 CST m=+0.006476501
//had sleep 2 second: 2023-10-22 18:21:20.6517914 +0800 CST m=+2.044795501
//reset: 2023-10-22 18:21:20.6520226 +0800 CST m=+2.045026701
//Reset timer1 return false
//timer1 has been 5+ seconds and Now time is
//2023-10-22 18:21:25.6589852 +0800 CST m=+7.051989301
//--------------------
//timer2
//--------------------
//before: 2023-10-22 18:21:30.6569628 +0800 CST m=+12.049966901
//had sleep 2 second: 2023-10-22 18:21:32.6692213 +0800 CST m=+14.062225401
//Reset timer2 return false
//timer2 has been 5+ seconds and Now time is
//2023-10-22 18:21:37.6838445 +0800 CST m=+19.076848601
//END
```



#### (*Timer) Stop 

``` go 
func (t *Timer) Stop() bool
```

​	Stop方法用于阻止计时器触发。如果该调用成功停止了计时器，则返回`true`；如果计时器已过期或已停止，则返回`false`。Stop方法不会关闭通道，以防止从通道中成功读取不正确的数据。

​	为了确保在调用Stop方法后通道为空，需要检查返回值并排空通道。例如，假设程序尚未从`t.C`接收到数据：

```go
if !t.Stop() {
	<-t.C
}
```

​	这不能与从该计时器的通道或其他对该计时器的Stop方法的调用并发执行。

​	对于使用`AfterFunc(d, f)`创建的计时器，如果t.Stop返回`false`，则计时器已过期，并且函数f已经在其自己的goroutine中启动；Stop方法在返回之前不会等待`f`完成。如果调用者需要知道`f`是否完成，它必须与`f`显式协调。

### type Weekday 

``` go 
type Weekday int
```

​	Weekday类型指定一周中的某一天(星期日= 0，...)。

``` go 
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
```

#### (Weekday) String 

``` go 
func (d Weekday) String() string
```

​	String方法返回一周中指定的星期几的英文名称("Sunday"、"Monday"，等等)。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 7; i++ {
		wd := time.Weekday(i)
		fmt.Printf("%s\n", wd.String())
	}
}

// Output:
//Sunday
//Monday
//Tuesday
//Wednesday
//Thursday
//Friday
//Saturday 
```

