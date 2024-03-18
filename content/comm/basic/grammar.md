+++
title = "Go和Python的基础语法"
date = 2024-03-01T15:08:03+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Go和Python的基础语法

## 数据类型

​	既然说，计算机是用来处理人给的数据，那么我们就首先来看看这两种语言各自能处理哪些类型的数据吧。

### Go的内置数据类型

| 序号 | 数据类型   | 名称       | 别名 | 默认值                      | 占用字节数 | 数据最小值                                                   | 数据最大值                                                   | 备注                                                         | 链接 |
| ---- | ---------- | ---------- | ---- | --------------------------- | ---------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ---- |
| 1    | bool       | 布尔型     |      | false                       | 1          | 无                                                           | 无                                                           | 只有true和false两种值                                        |      |
| 2    | byte       | 字节型     |      | '\x00'（使用%q）0（使用%d） | 1          | '\x00'（使用%q）0（使用%d）                                  | '\x7f'（使用%q）127（使用%d）                                | 是uint8的别名                                                |      |
| 3    | complex64  | 复数型     |      | (0+0i)                      | 8          | 无                                                           | 无                                                           |                                                              |      |
| 4    | complex128 | 复数型     |      | (0+0i)                      | 16         | 无                                                           | 无                                                           |                                                              |      |
| 5    | float32    | 浮点型     |      | 0                           | 4          | 1.401298464324817070923729583289916131280e-45 最小正非零值（使用%.39e保留39位小数） | 3.40282346638528859811704183484516925440e+38 （使用%.38e保留38位小数） |                                                              |      |
| 6    | float64    | 浮点型     |      | 0                           | 8          | 4.94065645841246544176568792868221372365059803e-324 最小正非零值（使用%.44e保留44位小数） | 1.797693134862315708145274237317043567980706e+308 （使用%.42e保留42位小数） |                                                              |      |
| 7    | int8       | 有符号整型 |      | 0                           | 1          | -128                                                         | 127                                                          |                                                              |      |
| 8    | int16      | 有符号整型 |      | 0                           | 2          | -32768                                                       | 32767                                                        |                                                              |      |
| 9    | int32      | 有符号整型 | rune | 0                           | 4          | -2147483648                                                  | 2147483647                                                   |                                                              |      |
| 10   | int64      | 有符号整型 |      | 0                           | 8          | -9223372036854775808                                         | 9223372036854775807（>922亿亿）                              |                                                              |      |
| 11   | int        | 有符号整型 |      | 0                           | 8          | -9223372036854775808                                         | 9223372036854775807（>922亿亿）                              | 请注意：这里给出的是64位系统的情况！                         |      |
| 12   | uint8      | 无符号整型 | byte | 0                           | 1          | 0                                                            | 255                                                          |                                                              |      |
| 13   | uint16     | 无符号整型 |      | 0                           | 2          | 0                                                            | 65535                                                        |                                                              |      |
| 14   | uint32     | 无符号整型 |      | 0                           | 4          | 0                                                            | 4294967295                                                   |                                                              |      |
| 15   | uint64     | 无符号整型 |      | 0                           | 8          | 0                                                            | 18446744073709551615                                         |                                                              |      |
| 16   | uint       | 无符号整型 |      | 0                           | 8          | 0                                                            | 18446744073709551615                                         |                                                              |      |
| 17   | rune       | 符文型     |      | '\x00'（使用%q）0（使用%d） | 4          | '\x00'（使用%q）0（使用%d）                                  |                                                              | 是int32的别名，而非uint32的别名                              |      |
| 18   | uintptr    | 指针整数型 |      | 无                          |            |                                                              |                                                              | uintptr 是一个整数类型，它足够大，可以容纳任何指针的比特模式 |      |
| 19   | string     | 字符串型   |      | ""                          |            |                                                              |                                                              |                                                              |      |
| 20   | [n]T       | 数组       |      | 空数组                      |            |                                                              |                                                              |                                                              |      |
| 21   | []T        | 切片       |      | 空切片                      |            |                                                              |                                                              |                                                              |      |
| 22   | map[K]V    | 映射       |      | 无                          |            |                                                              |                                                              |                                                              |      |
| 23   | struct{}   | 结构体     |      | 各自字段的零值              |            |                                                              |                                                              |                                                              |      |
| 24   | chan       | 通道       |      | nil                         |            |                                                              |                                                              |                                                              |      |
| 25   | *T         | 指针       |      | 无                          |            |                                                              |                                                              |                                                              |      |
| 26   | interface  | 接口       |      | 无                          |            |                                                              |                                                              |                                                              |      |
| 27   | error      | 错误类型   |      | 无                          |            |                                                              |                                                              |                                                              |      |

#### 用法

​	现在已经知道Go语言有这么多内置类型，那怎么使用呢？那我们就从这些类型的变量和常量如何声明、赋值（获取）、运算入手吧。

{{< tabpane text=true >}}

{{< tab header="bool" >}}

```go
package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gb1 = true
var gb2 bool = false

var verbs = []string{"T", "v", "+v", "#v", "t"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gb1", gb1, verbs)
	mfp.PrintFmtVal("全局变量 gb2", gb2, verbs)
	// 对部分全局变量进行修改
	gb1 = false
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gb1", gb1, verbs)
	fmt.Println("---局部变量---")
	// 声明方式1
	var b1 bool // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 false
	mfp.PrintFmtVal("声明方式1 b1", b1, verbs)
	// 赋值
	b1 = true
	mfp.PrintFmtVal("赋值后", b1, verbs)

	b1 = false
	mfp.PrintFmtVal("赋值后", b1, verbs)

	// 声明方式2
	var b2 = true
	mfp.PrintFmtVal("声明方式2 b2", b2, verbs)

	//短变量声明，仅用于局部变量
	b3 := true
	mfp.PrintFmtVal("声明方式3 b3", b3, verbs)

	b4 := false
	_ = b4 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错
}

---init 修改前---
全局变量 gb1:   %T -> bool | %v -> true | %+v -> true | %#v -> true | %t -> true | 
全局变量 gb2:   %T -> bool | %v -> false | %+v -> false | %#v -> false | %t -> false | 
---init 执行完成后---
全局变量 gb1:   %T -> bool | %v -> false | %+v -> false | %#v -> false | %t -> false | 
---局部变量---
声明方式1 b1:   %T -> bool | %v -> false | %+v -> false | %#v -> false | %t -> false | 
赋值后:         %T -> bool | %v -> true | %+v -> true | %#v -> true | %t -> true | 
赋值后:         %T -> bool | %v -> false | %+v -> false | %#v -> false | %t -> false | 
声明方式2 b2:   %T -> bool | %v -> true | %+v -> true | %#v -> true | %t -> true | 
声明方式3 b3:   %T -> bool | %v -> true | %+v -> true | %#v -> true | %t -> true | 
```

{{< /tab  >}}

{{< tab header="byte" >}}

```go
package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gbt1 = byte('i') // 注意这里需要使用byte()进行类型转换，这里的byte()并非函数，仅仅是一个类型+一对()而已
var gbt2 byte = 'j'

var verbs = []string{"T", "v", "+v", "#v", "q", "+q", "#q", "c"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gbt1", gbt1, verbs)
	mfp.PrintFmtVal("全局变量 gbt2", gbt2, verbs)

	// 对部分全局变量进行修改
	gbt1 = 'n'
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gbt1", gbt1, verbs)
	fmt.Println("---局部变量---")
	// 声明方式1
	var bt1 byte // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 '\x00'
	mfp.PrintFmtVal("声明方式1 bt1", bt1, verbs)
	// 赋值
	bt1 = 'A'
	mfp.PrintFmtVal("赋值后", bt1, verbs)
	bt1 = '\a' // 执行时会响铃
	mfp.PrintFmtVal("赋值后", bt1, verbs)

	// 声明方式2
	var bt2 = byte('h')
	mfp.PrintFmtVal("声明方式2 bt2", bt2, verbs)

	//短变量声明，仅用于局部变量
	bt3 := 'x'
	mfp.PrintFmtVal("声明方式3 bt3", bt3, verbs)

	bt4 := byte('\x00')
	_ = bt4 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错
}

---init 修改前---
全局变量 gbt1:  %T -> uint8 | %v -> 105 | %+v -> 105 | %#v -> 0x69 | %q -> 'i' | %+q -> 'i' | %#q -> 'i' | %c -> i |    
全局变量 gbt2:  %T -> uint8 | %v -> 106 | %+v -> 106 | %#v -> 0x6a | %q -> 'j' | %+q -> 'j' | %#q -> 'j' | %c -> j |    
---init 执行完成后---                                                                                                   
全局变量 gbt1:  %T -> uint8 | %v -> 110 | %+v -> 110 | %#v -> 0x6e | %q -> 'n' | %+q -> 'n' | %#q -> 'n' | %c -> n |    
---局部变量---                                                                                                          
声明方式1 bt1:  %T -> uint8 | %v -> 0 | %+v -> 0 | %#v -> 0x0 | %q -> '\x00' | %+q -> '\x00' | %#q -> '\x00' | %c ->  | 
赋值后:         %T -> uint8 | %v -> 65 | %+v -> 65 | %#v -> 0x41 | %q -> 'A' | %+q -> 'A' | %#q -> 'A' | %c -> A |      
赋值后:         %T -> uint8 | %v -> 7 | %+v -> 7 | %#v -> 0x7 | %q -> '\a' | %+q -> '\a' | %#q -> '\a' | %c ->  |       
声明方式2 bt2:  %T -> uint8 | %v -> 104 | %+v -> 104 | %#v -> 0x68 | %q -> 'h' | %+q -> 'h' | %#q -> 'h' | %c -> h |    
声明方式3 bt3:  %T -> int32 | %v -> 120 | %+v -> 120 | %#v -> 120 | %q -> 'x' | %+q -> 'x' | %#q -> 'x' | %c -> x | 
```

{{< /tab  >}}

{{< tab header="complex64/128" >}}

```go
package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gc641 = complex(float32(1), float32(2)) // 注意这里需要使用byte()进行类型转换，这里的byte()并非函数，仅仅是一个类型+一对()而已
var gc642 complex64 = 1 + 2i
var gc1281 = complex(float64(1), float64(2))
var gc1282 = complex(1, 2)
var gc1283 complex128 = 1 + 2i

var verbs = []string{"T", "v", "+v", "#v"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gc641", gc641, verbs)
	mfp.PrintFmtVal("全局变量 gc642", gc642, verbs)
	mfp.PrintFmtVal("全局变量 gc1281", gc1281, verbs)
	mfp.PrintFmtVal("全局变量 gc1282", gc1282, verbs)
	mfp.PrintFmtVal("全局变量 gc1283", gc1283, verbs)
	// 对部分全局变量进行修改
	gc641 = complex(float32(1.1), float32(2.2))
	gc1281 = complex(1.1, 2.2)
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gc641", gc641, verbs)
	mfp.PrintFmtVal("全局变量 gc1281", gc1281, verbs)
	fmt.Println("---局部变量---")
	fmt.Println("---complex64---")
	// 声明方式1
	var c641 complex64 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 false
	mfp.PrintFmtVal("声明方式1 c641", c641, verbs)
	// 赋值
	c641 = complex(float32(1), float32(2))
	mfp.PrintFmtVal("赋值后", c641, verbs)
	c641 = complex(float32(1.1), float32(2.2))
	mfp.PrintFmtVal("赋值后", c641, verbs)

	// 声明方式2
	var c642 complex64 = 1 + 2i
	mfp.PrintFmtVal("声明方式2 c642", c642, verbs)

	//短变量声明，仅用于局部变量
	c643 := complex(float32(1), float32(2))
	mfp.PrintFmtVal("声明方式3（短变量声明） c643", c643, verbs)

	x6431 := imag(c643)
	mfp.PrintFmtVal("调用imag函数 x6431", x6431, verbs)
	x6432 := real(c643)
	mfp.PrintFmtVal("调用real函数 x6432", x6432, verbs)

	c644 := complex(float32(1), float32(2))
	_ = c644 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错

	fmt.Println("---complex128---")
	// 声明方式1
    var c1281 complex128 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 (0+0i)
	mfp.PrintFmtVal("声明方式1 c1281", c1281, verbs)
	// 赋值
	c1281 = complex(1, 2)
	mfp.PrintFmtVal("赋值后", c1281, verbs)
	c1281 = complex(float64(1.1), float64(2.2))
	mfp.PrintFmtVal("赋值后", c1281, verbs)

	// 声明方式2
	var c1282 complex128 = 1 + 2i
	mfp.PrintFmtVal("声明方式2 c1282", c1282, verbs)

	//短变量声明，仅用于局部变量
	c1283 := complex(1, 2)
	mfp.PrintFmtVal("声明方式3（短变量声明） c1283", c1283, verbs)

	c1284 := complex(1, 2)
	_ = c1284
}

---init 修改前---
全局变量 gc641:         %T -> complex64 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |              
全局变量 gc642:         %T -> complex64 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |              
全局变量 gc1281:        %T -> complex128 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |             
全局变量 gc1282:        %T -> complex128 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |             
全局变量 gc1283:        %T -> complex128 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |             
---init 执行完成后---                                                                                 
全局变量 gc641:         %T -> complex64 | %v -> (1.1+2.2i) | %+v -> (1.1+2.2i) | %#v -> (1.1+2.2i) |  
全局变量 gc1281:        %T -> complex128 | %v -> (1.1+2.2i) | %+v -> (1.1+2.2i) | %#v -> (1.1+2.2i) | 
---局部变量---                                                                                        
---complex64---                                                                                       
声明方式1 c641:         %T -> complex64 | %v -> (0+0i) | %+v -> (0+0i) | %#v -> (0+0i) |              
赋值后:         %T -> complex64 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |                      
赋值后:         %T -> complex64 | %v -> (1.1+2.2i) | %+v -> (1.1+2.2i) | %#v -> (1.1+2.2i) |          
声明方式2 c642:         %T -> complex64 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |              
声明方式3（短变量声明） c643:   %T -> complex64 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |      
调用imag函数 x6431:     %T -> float32 | %v -> 2 | %+v -> 2 | %#v -> 2 |                               
调用real函数 x6432:     %T -> float32 | %v -> 1 | %+v -> 1 | %#v -> 1 |
---complex128---
声明方式1 c1281:        %T -> complex128 | %v -> (0+0i) | %+v -> (0+0i) | %#v -> (0+0i) |
赋值后:         %T -> complex128 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |
赋值后:         %T -> complex128 | %v -> (1.1+2.2i) | %+v -> (1.1+2.2i) | %#v -> (1.1+2.2i) |
声明方式2 c1282:        %T -> complex128 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |
声明方式3（短变量声明） c1283:  %T -> complex128 | %v -> (1+2i) | %+v -> (1+2i) | %#v -> (1+2i) |

```

{{< /tab  >}}

{{< tab header="float32/64" >}}

```go
package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gf321 = float32(1.1) // 注意这里需要使用byte()进行类型转换，这里的byte()并非函数，仅仅是一个类型+一对()而已
var gf322 float32 = float32(2.2)
var gf641 = 1.1
var gf642 float64 = 2.2

var verbs = []string{"T", "v", "+v", "#v", "b", "e", "E", "f", "F", "g", "G", "x", "X"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gf321", gf321, verbs)
	mfp.PrintFmtVal("全局变量 gf322", gf322, verbs)
	mfp.PrintFmtVal("全局变量 gf641", gf641, verbs)
	mfp.PrintFmtVal("全局变量 gf642", gf642, verbs)
	// 对部分全局变量进行修改
	gf321 = 1234567890.123456789
	gf641 = 1234567890.123456789
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gf321", gf321, verbs)
	mfp.PrintFmtVal("全局变量 gf641", gf641, verbs)
	fmt.Println("---局部变量---")
	fmt.Println("---float32---")
	// 声明方式1
	var f321 float32 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 f321", f321, verbs)
	// 赋值
	f321 = 12
	mfp.PrintFmtVal("赋值后", f321, verbs)
	f321 = 1234567890.123456789
	mfp.PrintFmtVal("赋值后", f321, verbs)

	// 声明方式2
	var f322 float32 = 1.1
	mfp.PrintFmtVal("声明方式2 f322", f322, verbs)

	//短变量声明，仅用于局部变量
	f323 := float32(2.2)
	mfp.PrintFmtVal("声明方式3（短变量声明） f323", f323, verbs)

	f324 := float32(1.1)
	_ = f324 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错

	fmt.Println("---float64---")
	// 声明方式1
	var f641 float64 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 f641", f641, verbs)
	// 赋值
	f641 = 13
	mfp.PrintFmtVal("赋值后", f641, verbs)
	f641 = 1234567890.123456789
	mfp.PrintFmtVal("赋值后", f641, verbs)

	// 声明方式2
	var f642 float64 = 1.1
	mfp.PrintFmtVal("声明方式2 f642", f642, verbs)

	//短变量声明，仅用于局部变量
	f643 := 2.2
	mfp.PrintFmtVal("声明方式3（短变量声明） f643", f643, verbs)

	f644 := 1.1
	_ = f644
}

---init 修改前---
全局变量 gf321:         %T -> float32 | %v -> 1.1 | %+v -> 1.1 | %#v -> 1.1 | %b -> 9227469p-23 | %e -> 1.100000e+00 | %E -> 1.100000E+00 | %f -> 1.100000 | %F -> 1.100000 | %g -> 1.1 | %G -> 1.1 | %x -> 0x1.19999ap+00 | %X -> 0X1.19999AP+00 | 
全局变量 gf322:         %T -> float32 | %v -> 2.2 | %+v -> 2.2 | %#v -> 2.2 | %b -> 9227469p-22 | %e -> 2.200000e+00 | %E -> 2.200000E+00 | %f -> 2.200000 | %F -> 2.200000 | %g -> 2.2 | %G -> 2.2 | %x -> 0x1.19999ap+01 | %X -> 0X1.19999AP+01 | 
全局变量 gf641:         %T -> float64 | %v -> 1.1 | %+v -> 1.1 | %#v -> 1.1 | %b -> 4953959590107546p-52 | %e -> 1.100000e+00 | %E -> 1.100000E+00 | %f -> 1.100000 | %F -> 1.100000 | %g -> 1.1 | %G -> 1.1 | %x -> 0x1.199999999999ap+00 | %X -> 0X1.199999999999AP+00 | 
全局变量 gf642:         %T -> float64 | %v -> 2.2 | %+v -> 2.2 | %#v -> 2.2 | %b -> 4953959590107546p-51 | %e -> 2.200000e+00 | %E -> 2.200000E+00 | %f -> 2.200000 | %F -> 2.200000 | %g -> 2.2 | %G -> 2.2 | %x -> 0x1.199999999999ap+01 | %X -> 0X1.199999999999AP+01 | 
---init 执行完成后---
全局变量 gf321:         %T -> float32 | %v -> 1.234568e+09 | %+v -> 1.234568e+09 | %#v -> 1.234568e+09 | %b -> 9645062p+7 | %e -> 1.234568e+09 | %E -> 1.234568E+09 | %f -> 1234567936.000000 | %F -> 1234567936.000000 | %g -> 1.234568e+09 | %G -> 1.234568E+09 | %x -> 0x1.26580cp+30 | %X -> 0X1.26580CP+30 | 
全局变量 gf641:         %T -> float64 | %v -> 1.2345678901234567e+09 | %+v -> 1.2345678901234567e+09 | %#v -> 1.2345678901234567e+09 | %b -> 5178153039816375p-22 | %e -> 1.234568e+09 | %E -> 1.234568E+09 | %f -> 1234567890.123457 | %F -> 1234567890.123457 | %g -> 1.2345678901234567e+09 | %G -> 1.2345678901234567E+09 | %x -> 0x1.26580b487e6b7p+30 | %X -> 0X1.26580B487E6B7P+30 | 
---局部变量---
---float32---
声明方式1 f321:         %T -> float32 | %v -> 0 | %+v -> 0 | %#v -> 0 | %b -> 0p-149 | %e -> 0.000000e+00 | %E -> 0.000000E+00 | %f -> 0.000000 | %F -> 0.000000 | %g -> 0 | %G -> 0 | %x -> 0x0p+00 | %X -> 0X0P+00 | 
赋值后:         %T -> float32 | %v -> 12 | %+v -> 12 | %#v -> 12 | %b -> 12582912p-20 | %e -> 1.200000e+01 | %E -> 1.200000E+01 | %f -> 12.000000 | %F -> 12.000000 | %g -> 12 | %G -> 12 | %x -> 0x1.8p+03 | %X -> 0X1.8P+03 | 
赋值后:         %T -> float32 | %v -> 1.234568e+09 | %+v -> 1.234568e+09 | %#v -> 1.234568e+09 | %b -> 9645062p+7 | %e -> 1.234568e+09 | %E -> 1.234568E+09 | %f -> 1234567936.000000 | %F -> 1234567936.000000 | %g -> 1.234568e+09 | %G -> 1.234568E+09 | %x -> 0x1.26580cp+30 | %X -> 0X1.26580CP+30 | 
声明方式2 f322:         %T -> float32 | %v -> 1.1 | %+v -> 1.1 | %#v -> 1.1 | %b -> 9227469p-23 | %e -> 1.100000e+00 | %E -> 1.100000E+00 | %f -> 1.100000 | %F -> 1.100000 | %g -> 1.1 | %G -> 1.1 | %x -> 0x1.19999ap+00 | %X -> 0X1.19999AP+00 | 
声明方式3（短变量声明） f323:   %T -> float32 | %v -> 2.2 | %+v -> 2.2 | %#v -> 2.2 | %b -> 9227469p-22 | %e -> 2.200000e+00 | %E -> 2.200000E+00 | %f -> 2.200000 | %F -> 2.200000 | %g -> 2.2 | %G -> 2.2 | %x -> 0x1.19999ap+01 | %X -> 0X1.19999AP+01 | 
---float64---
声明方式1 f641:         %T -> float64 | %v -> 0 | %+v -> 0 | %#v -> 0 | %b -> 0p-1074 | %e -> 0.000000e+00 | %E -> 0.000000E+00 | %f -> 0.000000 | %F -> 0.000000 | %g -> 0 | %G -> 0 | %x -> 0x0p+00 | %X -> 0X0P+00 | 
赋值后:         %T -> float64 | %v -> 13 | %+v -> 13 | %#v -> 13 | %b -> 7318349394477056p-49 | %e -> 1.300000e+01 | %E -> 1.300000E+01 | %f -> 13.000000 | %F -> 13.000000 | %g -> 13 | %G -> 13 | %x -> 0x1.ap+03 | %X -> 0X1.AP+03 | 
赋值后:         %T -> float64 | %v -> 1.2345678901234567e+09 | %+v -> 1.2345678901234567e+09 | %#v -> 1.2345678901234567e+09 | %b -> 5178153039816375p-22 | %e -> 1.234568e+09 | %E -> 1.234568E+09 | %f -> 1234567890.123457 | %F -> 1234567890.123457 | %g -> 1.2345678901234567e+09 | %G -> 1.2345678901234567E+09 | %x -> 0x1.26580b487e6b7p+30 | %X -> 0X1.26580B487E6B7P+30 | 
声明方式2 f642:         %T -> float64 | %v -> 1.1 | %+v -> 1.1 | %#v -> 1.1 | %b -> 4953959590107546p-52 | %e -> 1.100000e+00 | %E -> 1.100000E+00 | %f -> 1.100000 | %F -> 1.100000 | %g -> 1.1 | %G -> 1.1 | %x -> 0x1.199999999999ap+00 | %X -> 0X1.199999999999AP+00 | 
声明方式3（短变量声明） f643:   %T -> float64 | %v -> 2.2 | %+v -> 2.2 | %#v -> 2.2 | %b -> 4953959590107546p-51 | %e -> 2.200000e+00 | %E -> 2.200000E+00 | %f -> 2.200000 | %F -> 2.200000 | %g -> 2.2 | %G -> 2.2 | %x -> 0x1.199999999999ap+01 | %X -> 0X1.199999999999AP+01 | 
```

{{< /tab  >}}

{{< tab header="int/int8/16/32/64" >}}

```go
package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gi81 = int8(1) // 注意这里需要使用byte()进行类型转换，这里的byte()并非函数，仅仅是一个类型+一对()而已
var gi82 int8 = -2
var gi161 = int16(1)
var gi162 int16 = -2
var gi321 = int32(1)
var gi322 int32 = -2
var gi641 = int64(1)
var gi642 int64 = -2
var gi1 = 1
var gi2 int = -2
var verbs = []string{"T", "v", "+v", "#v", "b", "c", "d", "o", "O", "q", "x", "X", "U"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gi81", gi81, verbs)
	mfp.PrintFmtVal("全局变量 gi82", gi82, verbs)
	mfp.PrintFmtVal("全局变量 gi161", gi161, verbs)
	mfp.PrintFmtVal("全局变量 gi162", gi162, verbs)
	mfp.PrintFmtVal("全局变量 gi321", gi321, verbs)
	mfp.PrintFmtVal("全局变量 gi322", gi322, verbs)
	mfp.PrintFmtVal("全局变量 gi641", gi641, verbs)
	mfp.PrintFmtVal("全局变量 gi642", gi642, verbs)
	mfp.PrintFmtVal("全局变量 gi1", gi1, verbs)
	mfp.PrintFmtVal("全局变量 gi2", gi2, verbs)
	// 对部分全局变量进行修改
	gi81 = -12
	gi161 = -12
	gi321 = -12
	gi641 = -12
	gi1 = -12
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gi81", gi81, verbs)
	mfp.PrintFmtVal("全局变量 gi161", gi161, verbs)
	mfp.PrintFmtVal("全局变量 gi321", gi321, verbs)
	mfp.PrintFmtVal("全局变量 gi641", gi641, verbs)
	mfp.PrintFmtVal("全局变量 gi1", gi1, verbs)
	fmt.Println("---局部变量---")
	fmt.Println("---int8---")
	// 声明方式1
	var i81 int8 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i81", i81, verbs)
	// 赋值
	i81 = 1
	mfp.PrintFmtVal("赋值后", i81, verbs)
	i81 = 11
	mfp.PrintFmtVal("赋值后", i81, verbs)

	// 声明方式2
	var i82 int8 = 20
	mfp.PrintFmtVal("声明方式2 i82", i82, verbs)

	//短变量声明，仅用于局部变量
	i83 := int8(30)
	mfp.PrintFmtVal("声明方式3（短变量声明） i83", i83, verbs)

	i84 := int8(40)
	_ = i84 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错

	fmt.Println("---int16---")
	// 声明方式1
	var i161 int16 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i161", i161, verbs)
	// 赋值
	i161 = 1
	mfp.PrintFmtVal("赋值后", i161, verbs)
	i161 = 11
	mfp.PrintFmtVal("赋值后", i161, verbs)

	// 声明方式2
	var i162 int16 = 12
	mfp.PrintFmtVal("声明方式2 i162", i162, verbs)

	//短变量声明，仅用于局部变量
	i163 := int16(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） i163", i163, verbs)

	i164 := int16(1234)
	_ = i164

	fmt.Println("---int32---")
	// 声明方式1
	var i321 int32 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i321", i321, verbs)
	// 赋值
	i321 = 1
	mfp.PrintFmtVal("赋值后", i321, verbs)
	i321 = 11
	mfp.PrintFmtVal("赋值后", i321, verbs)

	// 声明方式2
	var i322 int32 = 12
	mfp.PrintFmtVal("声明方式2 i322", i322, verbs)

	//短变量声明，仅用于局部变量
	i323 := int32(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） i323", i323, verbs)

	i324 := int32(1234)
	_ = i324

	fmt.Println("---int64---")
	// 声明方式1
	var i641 int64 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i641", i641, verbs)
	// 赋值
	i641 = 1
	mfp.PrintFmtVal("赋值后", i641, verbs)
	i641 = 11
	mfp.PrintFmtVal("赋值后", i641, verbs)

	// 声明方式2
	var i642 int64 = 12
	mfp.PrintFmtVal("声明方式2 i642", i642, verbs)

	//短变量声明，仅用于局部变量
	i643 := int64(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） i643", i643, verbs)

	i644 := int64(1234)
	_ = i644

	fmt.Println("---int---")
	// 声明方式1
	var i1 int // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i1", i1, verbs)
	// 赋值
	i1 = 1
	mfp.PrintFmtVal("赋值后", i1, verbs)
	i1 = 11
	mfp.PrintFmtVal("赋值后", i1, verbs)

	// 声明方式2
	var i2 int = 12
	mfp.PrintFmtVal("声明方式2 i2", i2, verbs)

	//短变量声明，仅用于局部变量
	i3 := 123
	mfp.PrintFmtVal("声明方式3（短变量声明） i3", i3, verbs)

	i4 := 1234
	_ = i4
}
---init 修改前---
全局变量 gi81:  %T -> int8 | %v -> 1 | %+v -> 1 | %#v -> 1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
全局变量 gi82:  %T -> int8 | %v -> -2 | %+v -> -2 | %#v -> -2 | %b -> -10 | %c -> � | %d -> -2 | %o -> -2 | %O -> -0o2 | %q -> '�' | %x -> -2 | %X -> -2 | %U -> U+FFFFFFFFFFFFFFFE | 
全局变量 gi161:         %T -> int16 | %v -> 1 | %+v -> 1 | %#v -> 1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
全局变量 gi162:         %T -> int16 | %v -> -2 | %+v -> -2 | %#v -> -2 | %b -> -10 | %c -> � | %d -> -2 | %o -> -2 | %O -> -0o2 | %q -> '�' | %x -> -2 | %X -> -2 | %U -> U+FFFFFFFFFFFFFFFE | 
全局变量 gi321:         %T -> int32 | %v -> 1 | %+v -> 1 | %#v -> 1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
全局变量 gi322:         %T -> int32 | %v -> -2 | %+v -> -2 | %#v -> -2 | %b -> -10 | %c -> � | %d -> -2 | %o -> -2 | %O -> -0o2 | %q -> '�' | %x -> -2 | %X -> -2 | %U -> U+FFFFFFFFFFFFFFFE | 
全局变量 gi641:         %T -> int64 | %v -> 1 | %+v -> 1 | %#v -> 1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
全局变量 gi642:         %T -> int64 | %v -> -2 | %+v -> -2 | %#v -> -2 | %b -> -10 | %c -> � | %d -> -2 | %o -> -2 | %O -> -0o2 | %q -> '�' | %x -> -2 | %X -> -2 | %U -> U+FFFFFFFFFFFFFFFE | 
全局变量 gi1:   %T -> int | %v -> 1 | %+v -> 1 | %#v -> 1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
全局变量 gi2:   %T -> int | %v -> -2 | %+v -> -2 | %#v -> -2 | %b -> -10 | %c -> � | %d -> -2 | %o -> -2 | %O -> -0o2 | %q -> '�' | %x -> -2 | %X -> -2 | %U -> U+FFFFFFFFFFFFFFFE | 
---init 执行完成后---
全局变量 gi81:  %T -> int8 | %v -> -12 | %+v -> -12 | %#v -> -12 | %b -> -1100 | %c -> � | %d -> -12 | %o -> -14 | %O -> -0o14 | %q -> '�' | %x -> -c | %X -> -C | %U -> U+FFFFFFFFFFFFFFF4 | 
全局变量 gi161:         %T -> int16 | %v -> -12 | %+v -> -12 | %#v -> -12 | %b -> -1100 | %c -> � | %d -> -12 | %o -> -14 | %O -> -0o14 | %q -> '�' | %x -> -c | %X -> -C | %U -> U+FFFFFFFFFFFFFFF4 | 
全局变量 gi321:         %T -> int32 | %v -> -12 | %+v -> -12 | %#v -> -12 | %b -> -1100 | %c -> � | %d -> -12 | %o -> -14 | %O -> -0o14 | %q -> '�' | %x -> -c | %X -> -C | %U -> U+FFFFFFFFFFFFFFF4 | 
全局变量 gi641:         %T -> int64 | %v -> -12 | %+v -> -12 | %#v -> -12 | %b -> -1100 | %c -> � | %d -> -12 | %o -> -14 | %O -> -0o14 | %q -> '�' | %x -> -c | %X -> -C | %U -> U+FFFFFFFFFFFFFFF4 | 
全局变量 gi1:   %T -> int | %v -> -12 | %+v -> -12 | %#v -> -12 | %b -> -1100 | %c -> � | %d -> -12 | %o -> -14 | %O -> -0o14 | %q -> '�' | %x -> -c | %X -> -C | %U -> U+FFFFFFFFFFFFFFF4 | 
---局部变量---
---int8---
声明方式1 i81:  %T -> int8 | %v -> 0 | %+v -> 0 | %#v -> 0 | %b -> 0 | %c ->  | %d -> 0 | %o -> 0 | %O -> 0o0 | %q -> '\x00' | %x -> 0 | %X -> 0 | %U -> U+0000 | 
赋值后:         %T -> int8 | %v -> 1 | %+v -> 1 | %#v -> 1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
赋值后:         %T -> int8 | %v -> 11 | %+v -> 11 | %#v -> 11 | %b -> 1011 | %c -> 
                                                                                    | %d -> 11 | %o -> 13 | %O -> 0o13 | %q -> '\v' | %x -> b | %X -> B | %U -> U+000B | 
声明方式2 i82:  %T -> int8 | %v -> 20 | %+v -> 20 | %#v -> 20 | %b -> 10100 | %c ->  | %d -> 20 | %o -> 24 | %O -> 0o24 | %q -> '\x14' | %x -> 14 | %X -> 14 | %U -> U+0014 | 
声明方式3（短变量声明） i83:    %T -> int8 | %v -> 30 | %+v -> 30 | %#v -> 30 | %b -> 11110 | %c ->  | %d -> 30 | %o -> 36 | %O -> 0o36 | %q -> '\x1e' | %x -> 1e | %X -> 1E | %U -> U+001E | 
---int16---
声明方式1 i161:         %T -> int16 | %v -> 0 | %+v -> 0 | %#v -> 0 | %b -> 0 | %c ->  | %d -> 0 | %o -> 0 | %O -> 0o0 | %q -> '\x00' | %x -> 0 | %X -> 0 | %U -> U+0000 | 
赋值后:         %T -> int16 | %v -> 1 | %+v -> 1 | %#v -> 1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
赋值后:         %T -> int16 | %v -> 11 | %+v -> 11 | %#v -> 11 | %b -> 1011 | %c -> 
                                                                                     | %d -> 11 | %o -> 13 | %O -> 0o13 | %q -> '\v' | %x -> b | %X -> B | %U -> U+000B | 
声明方式2 i162:         %T -> int16 | %v -> 12 | %+v -> 12 | %#v -> 12 | %b -> 1100 | %c -> 
                                                                                             | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
声明方式3（短变量声明） i163:   %T -> int16 | %v -> 123 | %+v -> 123 | %#v -> 123 | %b -> 1111011 | %c -> { | %d -> 123 | %o -> 173 | %O -> 0o173 | %q -> '{' | %x -> 7b | %X -> 7B | %U -> U+007B | 
---int32---
声明方式1 i321:         %T -> int32 | %v -> 0 | %+v -> 0 | %#v -> 0 | %b -> 0 | %c ->  | %d -> 0 | %o -> 0 | %O -> 0o0 | %q -> '\x00' | %x -> 0 | %X -> 0 | %U -> U+0000 | 
赋值后:         %T -> int32 | %v -> 1 | %+v -> 1 | %#v -> 1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
赋值后:         %T -> int32 | %v -> 11 | %+v -> 11 | %#v -> 11 | %b -> 1011 | %c -> 
                                                                                     | %d -> 11 | %o -> 13 | %O -> 0o13 | %q -> '\v' | %x -> b | %X -> B | %U -> U+000B | 
声明方式2 i322:         %T -> int32 | %v -> 12 | %+v -> 12 | %#v -> 12 | %b -> 1100 | %c -> 
                                                                                             | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
声明方式3（短变量声明） i323:   %T -> int32 | %v -> 123 | %+v -> 123 | %#v -> 123 | %b -> 1111011 | %c -> { | %d -> 123 | %o -> 173 | %O -> 0o173 | %q -> '{' | %x -> 7b | %X -> 7B | %U -> U+007B | 
---int64---
声明方式1 i641:         %T -> int64 | %v -> 0 | %+v -> 0 | %#v -> 0 | %b -> 0 | %c ->  | %d -> 0 | %o -> 0 | %O -> 0o0 | %q -> '\x00' | %x -> 0 | %X -> 0 | %U -> U+0000 | 
赋值后:         %T -> int64 | %v -> 1 | %+v -> 1 | %#v -> 1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
赋值后:         %T -> int64 | %v -> 11 | %+v -> 11 | %#v -> 11 | %b -> 1011 | %c -> 
                                                                                     | %d -> 11 | %o -> 13 | %O -> 0o13 | %q -> '\v' | %x -> b | %X -> B | %U -> U+000B | 
声明方式2 i642:         %T -> int64 | %v -> 12 | %+v -> 12 | %#v -> 12 | %b -> 1100 | %c -> 
                                                                                             | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
声明方式3（短变量声明） i643:   %T -> int64 | %v -> 123 | %+v -> 123 | %#v -> 123 | %b -> 1111011 | %c -> { | %d -> 123 | %o -> 173 | %O -> 0o173 | %q -> '{' | %x -> 7b | %X -> 7B | %U -> U+007B | 
---int---
声明方式1 i1:   %T -> int | %v -> 0 | %+v -> 0 | %#v -> 0 | %b -> 0 | %c ->  | %d -> 0 | %o -> 0 | %O -> 0o0 | %q -> '\x00' | %x -> 0 | %X -> 0 | %U -> U+0000 | 
赋值后:         %T -> int | %v -> 1 | %+v -> 1 | %#v -> 1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
赋值后:         %T -> int | %v -> 11 | %+v -> 11 | %#v -> 11 | %b -> 1011 | %c -> 
                                                                                   | %d -> 11 | %o -> 13 | %O -> 0o13 | %q -> '\v' | %x -> b | %X -> B | %U -> U+000B | 
声明方式2 i2:   %T -> int | %v -> 12 | %+v -> 12 | %#v -> 12 | %b -> 1100 | %c -> 
                                                                                   | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
声明方式3（短变量声明） i3:     %T -> int | %v -> 123 | %+v -> 123 | %#v -> 123 | %b -> 1111011 | %c -> { | %d -> 123 | %o -> 173 | %O -> 0o173 | %q -> '{' | %x -> 7b | %X -> 7B | %U -> U+007B |
```

{{< /tab  >}}

{{< tab header="uint/uint8/16/32/64" >}}

```go
package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gui81 = uint8(1) // 注意这里需要使用byte()进行类型转换，这里的byte()并非函数，仅仅是一个类型+一对()而已
var gui82 uint8 = 2
var gui161 = uint16(1)
var gui162 uint16 = 2
var gui321 = uint32(1)
var gui322 uint32 = 2
var gui641 = uint64(1)
var gui642 uint64 = 2
var gui1 = uint(1)
var gui2 uint = 2
var verbs = []string{"T", "v", "+v", "#v", "b", "c", "d", "o", "O", "q", "x", "X", "U"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gui81", gui81, verbs)
	mfp.PrintFmtVal("全局变量 gui82", gui82, verbs)
	mfp.PrintFmtVal("全局变量 gui161", gui161, verbs)
	mfp.PrintFmtVal("全局变量 gui162", gui162, verbs)
	mfp.PrintFmtVal("全局变量 gui321", gui321, verbs)
	mfp.PrintFmtVal("全局变量 gui322", gui322, verbs)
	mfp.PrintFmtVal("全局变量 gui641", gui641, verbs)
	mfp.PrintFmtVal("全局变量 gui642", gui642, verbs)
	mfp.PrintFmtVal("全局变量 gui1", gui1, verbs)
	mfp.PrintFmtVal("全局变量 gui2", gui2, verbs)
	// 对部分全局变量进行修改
	gui81 = 12
	gui161 = 12
	gui321 = 12
	gui641 = 12
	gui1 = 12
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gui81", gui81, verbs)
	mfp.PrintFmtVal("全局变量 gui161", gui161, verbs)
	mfp.PrintFmtVal("全局变量 gui321", gui321, verbs)
	mfp.PrintFmtVal("全局变量 gui641", gui641, verbs)
	mfp.PrintFmtVal("全局变量 gui1", gui1, verbs)
	fmt.Println("---局部变量---")
	fmt.Println("---uint8---")
	// 声明方式1
	var ui81 uint8 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 ui81", ui81, verbs)
	// 赋值
	ui81 = 1
	mfp.PrintFmtVal("赋值后", ui81, verbs)
	ui81 = 11
	mfp.PrintFmtVal("赋值后", ui81, verbs)

	// 声明方式2
	var ui82 uint8 = 20
	mfp.PrintFmtVal("声明方式2 ui82", ui82, verbs)

	//短变量声明，仅用于局部变量
	ui83 := uint8(30)
	mfp.PrintFmtVal("声明方式3（短变量声明） ui83", ui83, verbs)

	ui84 := uint8(40)
	_ = ui84 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错

	fmt.Println("---uint16---")
	// 声明方式1
	var ui161 uint16 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 ui161", ui161, verbs)
	// 赋值
	ui161 = 1
	mfp.PrintFmtVal("赋值后", ui161, verbs)
	ui161 = 11
	mfp.PrintFmtVal("赋值后", ui161, verbs)

	// 声明方式2
	var ui162 uint16 = 12
	mfp.PrintFmtVal("声明方式2 ui162", ui162, verbs)

	//短变量声明，仅用于局部变量
	ui163 := uint16(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） ui163", ui163, verbs)

	ui164 := uint16(1234)
	_ = ui164

	fmt.Println("---uint32---")
	// 声明方式1
	var ui321 uint32 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 ui321", ui321, verbs)
	// 赋值
	ui321 = 1
	mfp.PrintFmtVal("赋值后", ui321, verbs)
	ui321 = 11
	mfp.PrintFmtVal("赋值后", ui321, verbs)

	// 声明方式2
	var ui322 uint32 = 12
	mfp.PrintFmtVal("声明方式2 ui322", ui322, verbs)

	//短变量声明，仅用于局部变量
	ui323 := uint32(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） ui323", ui323, verbs)

	ui324 := uint32(1234)
	_ = ui324

	fmt.Println("---uint64---")
	// 声明方式1
	var ui641 uint64 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 ui641", ui641, verbs)
	// 赋值
	ui641 = 1
	mfp.PrintFmtVal("赋值后", ui641, verbs)
	ui641 = 11
	mfp.PrintFmtVal("赋值后", ui641, verbs)

	// 声明方式2
	var ui642 uint64 = 12
	mfp.PrintFmtVal("声明方式2 ui642", ui642, verbs)

	//短变量声明，仅用于局部变量
	ui643 := uint64(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） ui643", ui643, verbs)

	ui644 := uint64(1234)
	_ = ui644

	fmt.Println("---uint---")
	// 声明方式1
	var i1 uint // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i1", i1, verbs)
	// 赋值
	i1 = 1
	mfp.PrintFmtVal("赋值后", i1, verbs)
	i1 = 11
	mfp.PrintFmtVal("赋值后", i1, verbs)

	// 声明方式2
	var i2 uint = 12
	mfp.PrintFmtVal("声明方式2 i2", i2, verbs)

	//短变量声明，仅用于局部变量
	i3 := uint(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） i3", i3, verbs)

	i4 := uint(1234)
	_ = i4
}

---init 修改前---
全局变量 gui81:         %T -> uint8 | %v -> 1 | %+v -> 1 | %#v -> 0x1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
全局变量 gui82:         %T -> uint8 | %v -> 2 | %+v -> 2 | %#v -> 0x2 | %b -> 10 | %c ->  | %d -> 2 | %o -> 2 | %O -> 0o2 | %q -> '\x02' | %x -> 2 | %X -> 2 | %U -> U+0002 | 
全局变量 gui161:        %T -> uint16 | %v -> 1 | %+v -> 1 | %#v -> 0x1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
全局变量 gui162:        %T -> uint16 | %v -> 2 | %+v -> 2 | %#v -> 0x2 | %b -> 10 | %c ->  | %d -> 2 | %o -> 2 | %O -> 0o2 | %q -> '\x02' | %x -> 2 | %X -> 2 | %U -> U+0002 | 
全局变量 gui321:        %T -> uint32 | %v -> 1 | %+v -> 1 | %#v -> 0x1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
全局变量 gui322:        %T -> uint32 | %v -> 2 | %+v -> 2 | %#v -> 0x2 | %b -> 10 | %c ->  | %d -> 2 | %o -> 2 | %O -> 0o2 | %q -> '\x02' | %x -> 2 | %X -> 2 | %U -> U+0002 | 
全局变量 gui641:        %T -> uint64 | %v -> 1 | %+v -> 1 | %#v -> 0x1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
全局变量 gui642:        %T -> uint64 | %v -> 2 | %+v -> 2 | %#v -> 0x2 | %b -> 10 | %c ->  | %d -> 2 | %o -> 2 | %O -> 0o2 | %q -> '\x02' | %x -> 2 | %X -> 2 | %U -> U+0002 | 
全局变量 gui1:  %T -> uint | %v -> 1 | %+v -> 1 | %#v -> 0x1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
全局变量 gui2:  %T -> uint | %v -> 2 | %+v -> 2 | %#v -> 0x2 | %b -> 10 | %c ->  | %d -> 2 | %o -> 2 | %O -> 0o2 | %q -> '\x02' | %x -> 2 | %X -> 2 | %U -> U+0002 | 
---init 执行完成后---
全局变量 gui81:         %T -> uint8 | %v -> 12 | %+v -> 12 | %#v -> 0xc | %b -> 1100 | %c -> 
                                                                                              | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
全局变量 gui161:        %T -> uint16 | %v -> 12 | %+v -> 12 | %#v -> 0xc | %b -> 1100 | %c -> 
                                                                                               | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
全局变量 gui321:        %T -> uint32 | %v -> 12 | %+v -> 12 | %#v -> 0xc | %b -> 1100 | %c -> 
                                                                                               | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
全局变量 gui641:        %T -> uint64 | %v -> 12 | %+v -> 12 | %#v -> 0xc | %b -> 1100 | %c -> 
                                                                                               | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
全局变量 gui1:  %T -> uint | %v -> 12 | %+v -> 12 | %#v -> 0xc | %b -> 1100 | %c -> 
                                                                                     | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
---局部变量---
---uint8---
声明方式1 ui81:         %T -> uint8 | %v -> 0 | %+v -> 0 | %#v -> 0x0 | %b -> 0 | %c ->  | %d -> 0 | %o -> 0 | %O -> 0o0 | %q -> '\x00' | %x -> 0 | %X -> 0 | %U -> U+0000 | 
赋值后:         %T -> uint8 | %v -> 1 | %+v -> 1 | %#v -> 0x1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
赋值后:         %T -> uint8 | %v -> 11 | %+v -> 11 | %#v -> 0xb | %b -> 1011 | %c -> 
                                                                                      | %d -> 11 | %o -> 13 | %O -> 0o13 | %q -> '\v' | %x -> b | %X -> B | %U -> U+000B | 
声明方式2 ui82:         %T -> uint8 | %v -> 20 | %+v -> 20 | %#v -> 0x14 | %b -> 10100 | %c ->  | %d -> 20 | %o -> 24 | %O -> 0o24 | %q -> '\x14' | %x -> 14 | %X -> 14 | %U -> U+0014 | 
声明方式3（短变量声明） ui83:   %T -> uint8 | %v -> 30 | %+v -> 30 | %#v -> 0x1e | %b -> 11110 | %c ->  | %d -> 30 | %o -> 36 | %O -> 0o36 | %q -> '\x1e' | %x -> 1e | %X -> 1E | %U -> U+001E | 
---uint16---
声明方式1 ui161:        %T -> uint16 | %v -> 0 | %+v -> 0 | %#v -> 0x0 | %b -> 0 | %c ->  | %d -> 0 | %o -> 0 | %O -> 0o0 | %q -> '\x00' | %x -> 0 | %X -> 0 | %U -> U+0000 | 
赋值后:         %T -> uint16 | %v -> 1 | %+v -> 1 | %#v -> 0x1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
赋值后:         %T -> uint16 | %v -> 11 | %+v -> 11 | %#v -> 0xb | %b -> 1011 | %c -> 
                                                                                       | %d -> 11 | %o -> 13 | %O -> 0o13 | %q -> '\v' | %x -> b | %X -> B | %U -> U+000B | 
声明方式2 ui162:        %T -> uint16 | %v -> 12 | %+v -> 12 | %#v -> 0xc | %b -> 1100 | %c -> 
                                                                                               | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
声明方式3（短变量声明） ui163:  %T -> uint16 | %v -> 123 | %+v -> 123 | %#v -> 0x7b | %b -> 1111011 | %c -> { | %d -> 123 | %o -> 173 | %O -> 0o173 | %q -> '{' | %x -> 7b | %X -> 7B | %U -> U+007B | 
---uint32---
声明方式1 ui321:        %T -> uint32 | %v -> 0 | %+v -> 0 | %#v -> 0x0 | %b -> 0 | %c ->  | %d -> 0 | %o -> 0 | %O -> 0o0 | %q -> '\x00' | %x -> 0 | %X -> 0 | %U -> U+0000 | 
赋值后:         %T -> uint32 | %v -> 1 | %+v -> 1 | %#v -> 0x1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
赋值后:         %T -> uint32 | %v -> 11 | %+v -> 11 | %#v -> 0xb | %b -> 1011 | %c -> 
                                                                                       | %d -> 11 | %o -> 13 | %O -> 0o13 | %q -> '\v' | %x -> b | %X -> B | %U -> U+000B | 
声明方式2 ui322:        %T -> uint32 | %v -> 12 | %+v -> 12 | %#v -> 0xc | %b -> 1100 | %c -> 
                                                                                               | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
声明方式3（短变量声明） ui323:  %T -> uint32 | %v -> 123 | %+v -> 123 | %#v -> 0x7b | %b -> 1111011 | %c -> { | %d -> 123 | %o -> 173 | %O -> 0o173 | %q -> '{' | %x -> 7b | %X -> 7B | %U -> U+007B | 
---uint64---
声明方式1 ui641:        %T -> uint64 | %v -> 0 | %+v -> 0 | %#v -> 0x0 | %b -> 0 | %c ->  | %d -> 0 | %o -> 0 | %O -> 0o0 | %q -> '\x00' | %x -> 0 | %X -> 0 | %U -> U+0000 | 
赋值后:         %T -> uint64 | %v -> 1 | %+v -> 1 | %#v -> 0x1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
赋值后:         %T -> uint64 | %v -> 11 | %+v -> 11 | %#v -> 0xb | %b -> 1011 | %c -> 
                                                                                       | %d -> 11 | %o -> 13 | %O -> 0o13 | %q -> '\v' | %x -> b | %X -> B | %U -> U+000B | 
声明方式2 ui642:        %T -> uint64 | %v -> 12 | %+v -> 12 | %#v -> 0xc | %b -> 1100 | %c -> 
                                                                                               | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
声明方式3（短变量声明） ui643:  %T -> uint64 | %v -> 123 | %+v -> 123 | %#v -> 0x7b | %b -> 1111011 | %c -> { | %d -> 123 | %o -> 173 | %O -> 0o173 | %q -> '{' | %x -> 7b | %X -> 7B | %U -> U+007B | 
---uint---
声明方式1 i1:   %T -> uint | %v -> 0 | %+v -> 0 | %#v -> 0x0 | %b -> 0 | %c ->  | %d -> 0 | %o -> 0 | %O -> 0o0 | %q -> '\x00' | %x -> 0 | %X -> 0 | %U -> U+0000 | 
赋值后:         %T -> uint | %v -> 1 | %+v -> 1 | %#v -> 0x1 | %b -> 1 | %c ->  | %d -> 1 | %o -> 1 | %O -> 0o1 | %q -> '\x01' | %x -> 1 | %X -> 1 | %U -> U+0001 | 
赋值后:         %T -> uint | %v -> 11 | %+v -> 11 | %#v -> 0xb | %b -> 1011 | %c -> 
                                                                                     | %d -> 11 | %o -> 13 | %O -> 0o13 | %q -> '\v' | %x -> b | %X -> B | %U -> U+000B | 
声明方式2 i2:   %T -> uint | %v -> 12 | %+v -> 12 | %#v -> 0xc | %b -> 1100 | %c -> 
                                                                                     | %d -> 12 | %o -> 14 | %O -> 0o14 | %q -> '\f' | %x -> c | %X -> C | %U -> U+000C | 
声明方式3（短变量声明） i3:     %T -> uint | %v -> 123 | %+v -> 123 | %#v -> 0x7b | %b -> 1111011 | %c -> { | %d -> 123 | %o -> 173 | %O -> 0o173 | %q -> '{' | %x -> 7b | %X -> 7B | %U -> U+007B |
```

{{< /tab  >}}

{{< tab header="rune" >}}

```go
package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gr1 = 'A' 
var gr2 rune = 'j'

var verbs = []string{"T", "v", "+v", "#v", "q", "+q", "#q", "c"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gr1", gr1, verbs)
	mfp.PrintFmtVal("全局变量 gr2", gr2, verbs)

	// 对部分全局变量进行修改
	gr1 = 'n'
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gr1", gr1, verbs)
	fmt.Println("---局部变量---")
	// 声明方式1
	var r1 rune // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 '\x00'
	mfp.PrintFmtVal("声明方式1 r1", r1, verbs)
	// 赋值
	r1 = 'A'
	mfp.PrintFmtVal("赋值后", r1, verbs)
	r1 = '\a' // 执行时会响铃
	mfp.PrintFmtVal("赋值后", r1, verbs)

	// 声明方式2
	var r2 = '中'
	mfp.PrintFmtVal("声明方式2 r2", r2, verbs)

	//短变量声明，仅用于局部变量
	r3 := '国'
	mfp.PrintFmtVal("声明方式3 r3", r3, verbs)

	r4 := '\x00'
	_ = r4 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错
}

---init 修改前---
全局变量 gr1:   %T -> int32 | %v -> 65 | %+v -> 65 | %#v -> 65 | %q -> 'A' | %+q -> 'A' | %#q -> 'A' | %c -> A | 
全局变量 gr2:   %T -> int32 | %v -> 106 | %+v -> 106 | %#v -> 106 | %q -> 'j' | %+q -> 'j' | %#q -> 'j' | %c -> j | 
---init 执行完成后---
全局变量 gr1:   %T -> int32 | %v -> 110 | %+v -> 110 | %#v -> 110 | %q -> 'n' | %+q -> 'n' | %#q -> 'n' | %c -> n | 
---局部变量---
声明方式1 r1:   %T -> int32 | %v -> 0 | %+v -> 0 | %#v -> 0 | %q -> '\x00' | %+q -> '\x00' | %#q -> '\x00' | %c ->  | 
赋值后:         %T -> int32 | %v -> 65 | %+v -> 65 | %#v -> 65 | %q -> 'A' | %+q -> 'A' | %#q -> 'A' | %c -> A | 
赋值后:         %T -> int32 | %v -> 7 | %+v -> 7 | %#v -> 7 | %q -> '\a' | %+q -> '\a' | %#q -> '\a' | %c ->  | 
声明方式2 r2:   %T -> int32 | %v -> 20013 | %+v -> 20013 | %#v -> 20013 | %q -> '中' | %+q -> '\u4e2d' | %#q -> '中' | %c -> 中 | 
声明方式3 r3:   %T -> int32 | %v -> 22269 | %+v -> 22269 | %#v -> 22269 | %q -> '国' | %+q -> '\u56fd' | %#q -> '国' | %c -> 国 | 
```

{{< /tab  >}}

{{< tab header="uintptr" >}}

```go

```

{{< /tab  >}}

{{< tab header="string" >}}

```go
package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gs1 = "Hello World"
var gs2 string = "勇敢前行"

var verbs = []string{"T", "v", "+v", "#v", "s", "q", "+q", "#q", "x", "X"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gs1", gs1, verbs)
	mfp.PrintFmtVal("全局变量 gs2", gs2, verbs)
	// 对部分全局变量进行修改
	gs1 = "Hello 中国！"
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gs1", gs1, verbs)
	fmt.Println("---局部变量---")
	// 声明方式1
	var s1 string // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 ""
	mfp.PrintFmtVal("声明方式1 s1", s1, verbs)
	// 赋值
	s1 = "你好"
	mfp.PrintFmtVal("赋值后", s1, verbs)

	s1 = "Hello 你好"
	mfp.PrintFmtVal("赋值后", s1, verbs)

	// 声明方式2
	var b2 = "真诚勤勇"
	mfp.PrintFmtVal("声明方式2 b2", b2, verbs)
	mfp.PrintFmtVal("声明方式2 b2", b2, []string{"x"})
	mfp.PrintFmtVal("声明方式2 b2", b2, []string{"x", "X"})
	mfp.PrintFmtVal("声明方式2 b2", b2, []string{"#q", "x", "X"})

	//短变量声明，仅用于局部变量
	b3 := "Welcome to Go"
	mfp.PrintFmtVal("声明方式3 b3", b3, verbs)

	b4 := "Nice to meet you!很高兴见到你！"
	_ = b4 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错
}

---init 修改前---
全局变量 gs1:   %T -> string | %v -> Hello World | %+v -> Hello World | %#v -> "Hello World" | %s -> Hello World | %q -> "Hello World" | %+q -> "Hello World" | %#q -> `Hello World` | %x -> 48656c6c6f20576f726c64 | %X -> 48656C6C6F20576F726C64 | 
全局变量 gs2:   %T -> string | %v -> 勇敢前行 | %+v -> 勇敢前行 | %#v -> "勇敢前行" | %s -> 勇敢前行 | %q -> "勇敢前行" | %+q ->\u884c" | %#q -> `勇敢前行` | %x -> e58b87e695a2e5898de8a18c | %X -> E58B87E695A2E5898DE8A18C | 
---init 执行完成后---
全局变量 gs1:   %T -> string | %v -> Hello 中国！ | %+v -> Hello 中国！ | %#v -> "Hello 中国！" | %s -> Hello 中国！ | %q -> "He+q -> "Hello \u4e2d\u56fd\uff01" | %#q -> `Hello 中国！` | %x -> 48656c6c6f20e4b8ade59bbdefbc81 | %X -> 48656C6C6F20E4B8ADE59BBDC81 | 
---局部变量---
声明方式1 s1:   %T -> string | %v ->  | %+v ->  | %#v -> "" | %s ->  | %q -> "" | %+q -> "" | %#q -> `` | %x ->  | %X ->  | 
赋值后:         %T -> string | %v -> 你好 | %+v -> 你好 | %#v -> "你好" | %s -> 你好 | %q -> "你好" | %+q -> "\u4f60\u597d" | %#| %x -> e4bda0e5a5bd | %X -> E4BDA0E5A5BD | 
赋值后:         %T -> string | %v -> Hello 你好 | %+v -> Hello 你好 | %#v -> "Hello 你好" | %s -> Hello 你好 | %q -> "Hello 你好 "Hello \u4f60\u597d" | %#q -> `Hello 你好` | %x -> 48656c6c6f20e4bda0e5a5bd | %X -> 48656C6C6F20E4BDA0E5A5BD | 
声明方式2 b2:   %T -> string | %v -> 真诚勤勇 | %+v -> 真诚勤勇 | %#v -> "真诚勤勇" | %s -> 真诚勤勇 | %q -> "真诚勤勇" | %+q ->\u52c7" | %#q -> `真诚勤勇` | %x -> e79c9fe8af9ae58ba4e58b87 | %X -> E79C9FE8AF9AE58BA4E58B87 | 
声明方式2 b2:   %x -> e79c9fe8af9ae58ba4e58b87 | 
声明方式2 b2:   %x -> e79c9fe8af9ae58ba4e58b87 | %X -> E79C9FE8AF9AE58BA4E58B87 | 
声明方式2 b2:   %#q -> `真诚勤勇` | %x -> e79c9fe8af9ae58ba4e58b87 | %X -> E79C9FE8AF9AE58BA4E58B87 | 
声明方式3 b3:   %T -> string | %v -> Welcome to Go | %+v -> Welcome to Go | %#v -> "Welcome to Go" | %s -> Welcome to Go | %q -> "Welcome to Go" | %+q -> "Welcome to Go" | %#q -> `Welcome to Go` | %x -> 57656c636f6d6520746f20476f | %X -> 57656C636F6D6520746F20476F | 

```

{{< /tab  >}}

{{< tab header="数组" >}}

```go
package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var ga1 = [3]int{1, 2, 3}
var ga2 [3]int = [3]int{1, 2, 3}
var ga3 = [...]int{1, 2, 3}
var ga4 = [3]string{"a", "b", "c"}
var ga5 [3]string = [3]string{"a", "b", "c"}
var ga6 = [...]string{"a", "b", "c"}
var ga7 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
var ga8 [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

//var ga9 = [3][...]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}    // invalid use of [...] array (outside a composite literal)
//var ga10 = [...][...]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // invalid use of [...] array (outside a composite literal)

var verbs = []string{"T", "v", "+v", "#v"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 ga1", ga1, verbs)
	mfp.PrintFmtVal("全局变量 ga2", ga2, verbs)
	// 对部分全局变量进行修改
	ga1[0] = 11
	ga7[0][0] = 11
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 ga1", ga1, verbs)
	mfp.PrintFmtVal("全局变量 ga7", ga7, verbs)
	fmt.Println("---局部变量---")
	var a1 = [3]int{1, 2, 3}
	mfp.PrintFmtVal("a1", a1, []string{"T", "v"})
	fmt.Println("a1=", a1)
	fmt.Println("获取a1的长度 -> len(a1)=", len(a1))
	fmt.Println("获取a1的容量 -> cap(a1)=", cap(a1))
	// 修改某一元素的值
	a1[0] = 11
	mfp.PrintFmtVal("1 修改后a1", a1, []string{"T", "v"})

	p := &a1[0]
	mfp.PrintFmtVal("p是什么？", p, []string{"T", "v", "+v", "#v", "p", "P"})
	*p = 111
	mfp.PrintFmtVal("2 修改后a1", a1, []string{"T", "v"})

	fmt.Println("错误赋值：a1 = [4]int{1, 2, 3, 4} // cannot use [4]int{…} (value of type [4]int) as [3]int value in assignment")
	//a1 = [4]int{1, 2, 3, 4} // cannot use [4]int{…} (value of type [4]int) as [3]int value in assignment

	a2 := [3]int{1, 2, 3}
	a3 := [...]int{1, 2, 3}
	a4 := [...]int{1, 2, 33}

	if a2 == a3 {
		fmt.Println("a2和a3竟然相等！")
	} else {
		fmt.Println("a2和a3不相等！")
	}

	if a2 == a4 {
		fmt.Println("a2和a4竟然相等！")
	} else {
		fmt.Println("a2和a4不相等！")
	}

	// 产生切片
	sl1 := a2[:]
	sl2 := a2[0:]
	sl3 := a2[:len(a1)]
	sl4 := a2[0:len(a1)]
	sl5 := a2[1:2]

	mfp.PrintFmtVal("sl1", sl1, []string{"T", "v"})
	mfp.PrintFmtVal("sl2", sl2, []string{"T", "v"})
	mfp.PrintFmtVal("sl3", sl3, []string{"T", "v"})
	mfp.PrintFmtVal("sl4", sl4, []string{"T", "v"})
	mfp.PrintFmtVal("sl5", sl5, []string{"T", "v"})
}

---init 修改前---
全局变量 ga1:   %T -> [3]int | %v -> [1 2 3] | %+v -> [1 2 3] | %#v -> [3]int{1, 2, 3} | 
全局变量 ga2:   %T -> [3]int | %v -> [1 2 3] | %+v -> [1 2 3] | %#v -> [3]int{1, 2, 3} | 
---init 执行完成后---
全局变量 ga1:   %T -> [3]int | %v -> [11 2 3] | %+v -> [11 2 3] | %#v -> [3]int{11, 2, 3} | 
全局变量 ga7:   %T -> [3][3]int | %v -> [[11 2 3] [4 5 6] [7 8 9]] | %+v -> [[11 2 3] [4 5 6] [7 8 9]] | %#v -> [3][3]int{[3]int{11, 2, 3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}} | 
---局部变量---
a1:     %T -> [3]int | %v -> [1 2 3] | 
a1= [1 2 3]
获取a1的长度 -> len(a1)= 3
获取a1的容量 -> cap(a1)= 3
1 修改后a1:     %T -> [3]int | %v -> [11 2 3] | 
p是什么？:      %T -> *int | %v -> 0xc0000a40c0 | %+v -> 0xc0000a40c0 | %#v -> (*int)(0xc0000a40c0) | %p -> 0xc0000a40c0 | %P -> %!P(*int=0xc0000a40c0) | 
2 修改后a1:     %T -> [3]int | %v -> [111 2 3] | 
错误赋值：a1 = [4]int{1, 2, 3, 4} // cannot use [4]int{…} (value of type [4]int) as [3]int value in assignment
a2和a3竟然相等！
a2和a4不相等！
sl1:    %T -> []int | %v -> [1 2 3] | 
sl2:    %T -> []int | %v -> [1 2 3] | 
sl3:    %T -> []int | %v -> [1 2 3] | 
sl4:    %T -> []int | %v -> [1 2 3] | 
sl5:    %T -> []int | %v -> [2] |
```

{{< /tab  >}}

{{< tab header="切片" >}}

```go

```

{{< /tab  >}}

{{< tab header="映射" >}}

```go

```

{{< /tab  >}}

{{< tab header="结构体" >}}

```go

```

{{< /tab  >}}

{{< tab header="通道" >}}

```go

```

{{< /tab  >}}

{{< tab header="指针" >}}

```go

```

{{< /tab  >}}

{{< tab header="接口" >}}

```go

```

{{< /tab  >}}

{{< tab header="错误类型" >}}

```go

```

{{< /tab  >}}

{{< /tabpane >}}



### 数组

#### C创建

##### 一维数组

###### 直接创建

```go
var verbs = []string{"T", "v", "#v"}
var a0 [3]int
var a1 = [3]int{1, 2, 3}
var a2 [3]int = [3]int{1, 2, 3}
var a3 = [...]int{1, 2, 3}
ad1 := [...]int{1, 2, 3}
mfp.PrintFmtVal("a0", a0, verbs)
mfp.PrintFmtVal("a1", a1, verbs)
mfp.PrintFmtVal("a2", a2, verbs)
mfp.PrintFmtVal("a3", a3, verbs)
mfp.PrintFmtVal("ad1", ad1, verbs)
```

```
a0:     %T -> [3]int | %v -> [0 0 0] | %#v -> [3]int{0, 0, 0}
a1:     %T -> [3]int | %v -> [1 2 3] | %#v -> [3]int{1, 2, 3}
a2:     %T -> [3]int | %v -> [1 2 3] | %#v -> [3]int{1, 2, 3}
a3:     %T -> [3]int | %v -> [1 2 3] | %#v -> [3]int{1, 2, 3}
ad1:    %T -> [3]int | %v -> [1 2 3] | %#v -> [3]int{1, 2, 3}
```

###### 是否可以通过make创建？

​	=> 不可以

```go
a4 := make([3]int{1, 2, 3}) // 报错：[3]int{…} is not a type
```

###### 是否可以通过new创建？

​	=> 不可以

```go
a5 := new([3]int{1, 2, 3}) // 报错：[3]int{…} is not a type
```

##### 多维数组

​	从一维数组的创建中可以推测出，多维数组的创建也是不能通过make和new内置函数进行创建。

```go
fmt.Println("二维数组")
var a6 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
var a7 [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
var a8 = [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
//var a9 = [...][...]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // 报错：invalid use of [...] array (outside a composite literal)
ad2 := [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

fmt.Println("三维数组")
var a10 = [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
var a11 [2][2][2]int = [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
var a12 = [...][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
//var a13 = [...][...][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}} // 报错：invalid use of [...] array (outside a composite literal) 和 missing type in composite literal
ad3 := [...][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
```



#### U修改

##### 修改元素

```go
fmt.Println("一维数组")

a14 := [...]int{1, 2, 3}
mfp.PrintFmtVal("a14", a14, verbs)

a14[0] = 11
mfp.PrintFmtVal("a14", a14, verbs)

a14[len(a14)-1] = 33
mfp.PrintFmtVal("a14", a14, verbs)

pa141 := &a14[0]
*pa141 = 111
mfp.PrintFmtVal("a14", a14, verbs)
fmt.Println("二维数组")

a15 := [...][2]int{{1, 2}, {3, 4}}
a15[0][0] = 11
mfp.PrintFmtVal("a15", a15, verbs)

a15[len(a15)-1][0] = 33
mfp.PrintFmtVal("a15", a15, verbs)

pa151 := &a15[0][0]
*pa151 = 111
mfp.PrintFmtVal("a15", a15, verbs)
fmt.Println("三维数组和二维数组类似")
```

```
一维数组
a14:    %T -> [3]int | %v -> [1 2 3] | %#v -> [3]int{1, 2, 3}
a14:    %T -> [3]int | %v -> [11 2 3] | %#v -> [3]int{11, 2, 3}
a14:    %T -> [3]int | %v -> [11 2 33] | %#v -> [3]int{11, 2, 33}
a14:    %T -> [3]int | %v -> [111 2 33] | %#v -> [3]int{111, 2, 33}
二维数组
a15:    %T -> [2][2]int | %v -> [[11 2] [3 4]] | %#v -> [2][2]int{[2]int{11, 2}, [2]int{3, 4}}
a15:    %T -> [2][2]int | %v -> [[11 2] [33 4]] | %#v -> [2][2]int{[2]int{11, 2}, [2]int{33, 4}}
a15:    %T -> [2][2]int | %v -> [[111 2] [33 4]] | %#v -> [2][2]int{[2]int{111, 2}, [2]int{33, 4}}
三维数组和二维数组类似
```

##### 用整个数组赋值

```go
a16 := [...]int{1, 2, 3}
mfp.PrintFmtVal("a16", a16, verbs)
a16 = [...]int{2, 3, 4}
mfp.PrintFmtVal("赋值后 a16", a16, verbs)
//a16 = [...]int{2, 3, 4, 5} // 报错：cannot use [...]int{…} (value of type [4]int) as [3]int value in assignment
//a16 = [...]string{"a", "b", "c"} // 报错：cannot use [...]string{…} (value of type [3]string) as [3]int value in assignment	
```

```
a16:    %T -> [3]int | %v -> [1 2 3] | %#v -> [3]int{1, 2, 3}
赋值后 a16:    %T -> [3]int | %v -> [2 3 4] | %#v -> [3]int{2, 3, 4}
```

​	可以看出，整个数组赋值时，新旧两个数组的元素个数（长度）和数组元素类型一定要都一致，否则将报错。

#### A访问

​	访问数组中的某一元素，可通过索引下标，索引下标范围`[0, len(数组名) - 1]`，即从0开始到数组的长度减去1。

##### 直接访问指定索引下标的元素

```go
a17 := [...]int{1, 2, 3}
fmt.Println("直接访问指定索引下标的元素")
fmt.Println(a17[0])
fmt.Println(a17[1])
fmt.Println(a17[len(a17)-1])
```

```
1
2
3
```

##### 遍历数组

​	通过遍历的方式访问所需索引下标或全部索引下标的元素：

```go
for k, v := range a17 {
    if k%2 == 0 {
        fmt.Println(k, "->", v)
    }
}
mfp.PrintHr()
for k, v := range a17 {
    fmt.Println(k, "->", v)
}
```

```
0 -> 1
2 -> 3
------------------
0 -> 1
1 -> 2
2 -> 3
```

##### 获取相关数组属性

```go
a22 := [...]int{1, 2, 3}
fmt.Println("a22数组的长度 len(a22)=", len(a22))
fmt.Println("a22数组的容量 cap(a22)=", cap(a22))
```

```
a22数组的长度 len(a22)= 3
a22数组的容量 cap(a22)= 3
```

​	我们会发现任何数组的长度和容量是相等的。

#### D删除

##### 是否可以删除某一元素呢？

​	=> 不可以

​	通过上面的创建、修改、访问，我们知道数组有两个重要的属性：长度和元素类型。假设真能删除某一元素，那么新旧数组的长度就不一样了，这样就导致了前后两个数组不一致，故Go语言的设计中也没有提供删除数组元素的操作。

#### 作为实参传递给函数或方法

​	因数组在Go语言中是`值类型`，数组作为实参传递给函数，将发生完整复制数组，若数组很大，对于内存和性能将会是一个大开销。

#### 易混淆的知识点

##### 数组指针和指针数组

```go
fmt.Println("数组指针")
a18 := [...]int{1, 2, 3}
a19 := [...]int{1, 2, 3, 4}
_ = a19
var ptrA181 *[3]int
ptrA181 = &a18
mfp.PrintFmtVal("ptrA181", ptrA181, []string{"T", "v", "#v"})
mfp.PrintFmtVal("*ptrA181", *ptrA181, []string{"T", "v", "#v"})
//ptrA181 = &a19 // 报错：cannot use &a19 (value of type *[4]int) as *[3]int value in assignment

mfp.PrintHr()
fmt.Println("指针数组")
xa201, xa202, xa203 := 1, 2, 3
a20 := [...]*int{&xa201, &xa202, &xa203}
mfp.PrintFmtVal("a20", a20, []string{"T", "v", "#v"})
for k, v := range a20 {
    fmt.Println(k, "->", *v)
}
```

```
数组指针
ptrA181:        %T -> *[3]int | %v -> &[1 2 3] | %#v -> &[3]int{1, 2, 3}
*ptrA181:       %T -> [3]int | %v -> [1 2 3] | %#v -> [3]int{1, 2, 3}
------------------
指针数组
a20:    %T -> [3]*int | %v -> [0xc000012340 0xc000012348 0xc000012350] | %#v -> [3]*int{(*int)(0xc000012340), (*int)(0xc000012348), (*int)(0xc000012350)}
0 -> 1
1 -> 2
2 -> 3
```



#### 易错点

##### 访问最后一个数组元素

​	直接用a[len(a)]访问数组a的最后一个元素 =》肯定报错

```go
fmt.Println("访问数组的最后一个元素")
a21 := [...]int{1, 2, 3}
//fmt.Println(a21[len(a21)]) // 报错：invalid argument: index 3 out of bounds [0:3]
fmt.Println(a21[len(a21)-1]) // 正确方式
```

```
3
```

#### 数组的特点

​	数组中的元素在内存中的存储是连续的，故检索数组非常快，但定义后数组的大小不能再修改。	

### 切片

#### C创建

##### 直接创建

```go
var sl1 []int
var sl2 []int = []int{1, 2, 3}
var sl3 = []int{1, 2, 3}
sl4 := []int{1, 2, 3}
mfp.PrintFmtValWithLC("sl1", sl1, verbs)
mfp.PrintFmtValWithLC("sl2", sl2, verbs)
mfp.PrintFmtValWithLC("sl3", sl3, verbs)
mfp.PrintFmtValWithLC("sl4", sl4, verbs
```

```
sl1:    %T -> []int | %v -> [] | %#v -> []int(nil) | len=0 | cap=0
sl2:    %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl3:    %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl4:    %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
```



##### 基于数组创建

```go
a1 := [...]int{1, 2, 3, 4, 5, 6}
mfp.PrintFmtValWithLC("a1", a1, verbs)
sl5 := a1[:]
sl6 := a1[0:]
sl7 := a1[:len(a1)]
sl8 := a1[0:len(a1)]
//sl9 := a1[0:3:2] // 报错：invalid slice indices: 2 < 3
sl10 := a1[0:3:3]
sl11 := a1[0:3:4]
sl12 := a1[0:3:5]
sl13 := a1[0:3:6]
//sl14 := a1[0:3:7] // 报错：invalid argument: index 7 out of bounds [0:7]
mfp.PrintFmtValWithLC("sl5", sl5, verbs)
mfp.PrintFmtValWithLC("sl6", sl6, verbs)
mfp.PrintFmtValWithLC("sl7", sl7, verbs)
mfp.PrintFmtValWithLC("sl8", sl8, verbs)
//mfp.PrintFmtValWithLC("sl9", sl9, verbs)
mfp.PrintFmtValWithLC("sl10", sl10, verbs)
mfp.PrintFmtValWithLC("sl11", sl11, verbs)
mfp.PrintFmtValWithLC("sl12", sl12, verbs)
mfp.PrintFmtValWithLC("sl13", sl13, verbs)
//mfp.PrintFmtValWithLC("sl14", sl14, verbs)
```

```
a1:     %T -> [6]int | %v -> [1 2 3 4 5 6] | %#v -> [6]int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
sl5:    %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
sl6:    %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
sl7:    %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
sl8:    %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
sl10:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl11:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=4
sl12:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=5
sl13:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=6
```

​	由以上示例在使用 `a[low:high:max]`获取新切片时可以看出：`max >= high >= low` ；`max`不得大于底层数组的上边界所在的索引； 新切片的长度为`high - low`，而容量为`max - low`。


##### 用make创建

```go
sl15 := make([]int, 3)
//sl16 := make([]int, 3, 2) // 报错：invalid argument: length and capacity swapped
sl17 := make([]int, 3, 3)
sl18 := make([]int, 3, 4)
mfp.PrintFmtValWithLC("sl15", sl15, verbs)
//mfp.PrintFmtValWithLC("sl16", sl16, verbs)
mfp.PrintFmtValWithLC("sl17", sl17, verbs)
mfp.PrintFmtValWithLC("sl18", sl18, verbs)
```

```
sl15:   %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=3
sl17:   %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=3
sl18:   %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=4
```



##### 用new创建

```go
sl19 := *new([]int) // 注意此时 sl19 为空切片，其长度和容量都为0
mfp.PrintFmtValWithLC("sl19", sl19, verbs)
sl19 = append(sl19, 1)
mfp.PrintFmtValWithLC("sl19", sl19, verbs)
sl19 = append(sl19, 2)
mfp.PrintFmtValWithLC("sl19", sl19, verbs)
sl19 = append(sl19, 3)
mfp.PrintFmtValWithLC("sl19", sl19, verbs)
```

```
sl19:   %T -> []int | %v -> [] | %#v -> []int(nil) | len=0 | cap=0
sl19:   %T -> []int | %v -> [1] | %#v -> []int{1} | len=1 | cap=1
sl19:   %T -> []int | %v -> [1 2] | %#v -> []int{1, 2} | len=2 | cap=2
sl19:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=4
```

##### 基于已有切片创建

```go
a2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
mfp.PrintFmtValWithLC("已有数组 a2", a2, verbs)

sl20 := a2[0:6]
mfp.PrintFmtValWithLC("已有切片 sl20", sl20, verbs)

sl21 := sl20[:]
sl22 := sl20[0:]
sl23 := sl20[:len(sl20)]
sl24 := sl20[:cap(sl20)]
sl25 := sl20[0:len(sl20)]
sl26 := sl20[0:cap(sl20)]
//sl27 := sl20[0:cap(sl20)+1] // 报错：panic: runtime error: slice bounds out of range [:11] with capacity 10
sl28 := sl20[1:3]
sl29 := sl20[1:4]
sl30 := sl20[2:4]
//sl31 := sl20[2:4:2] // 报错：invalid slice indices: 2 < 4
//sl32 := sl20[2:4:3] // 报错：invalid slice indices: 3 < 4
sl33 := sl20[2:4:4]
sl34 := sl20[2:4:5]
sl35 := sl20[2:4:6]
sl36 := sl20[2:4:7]

mfp.PrintFmtValWithLC("sl21=sl20[:]", sl21, verbs)
mfp.PrintFmtValWithLC("sl22=sl20[0:]", sl22, verbs)
mfp.PrintFmtValWithLC("sl23=sl20[:len(sl20)]", sl23, verbs)
mfp.PrintFmtValWithLC("sl24=sl20[:cap(sl20)]", sl24, verbs)
mfp.PrintFmtValWithLC("sl25=[0:len(sl20)]", sl25, verbs)
mfp.PrintFmtValWithLC("sl26=[0:cap(sl20)]", sl26, verbs)
//mfp.PrintFmtValWithLC("sl27=sl20[0:cap(sl20)+1]", sl27, verbs)
mfp.PrintFmtValWithLC("sl28=sl20[1:3]", sl28, verbs)
mfp.PrintFmtValWithLC("sl29=sl20[1:4]", sl29, verbs)
mfp.PrintFmtValWithLC("sl30=sl20[2:4]", sl30, verbs)
//mfp.PrintFmtValWithLC("sl31=sl20[2:4:2]", sl31, verbs)
//mfp.PrintFmtValWithLC("sl32=sl20[2:4:3]", sl32, verbs)
mfp.PrintFmtValWithLC("sl33=sl20[2:4:4]", sl33, verbs)
mfp.PrintFmtValWithLC("sl34=sl20[2:4:5]", sl34, verbs)
mfp.PrintFmtValWithLC("sl35=sl20[2:4:6]", sl35, verbs)
mfp.PrintFmtValWithLC("sl36=sl20[2:4:7]", sl36, verbs)
```

```
已有数组 a2:    %T -> [10]int | %v -> [1 2 3 4 5 6 7 8 9 10] | %#v -> [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} | len=10 | cap=10
已有切片 sl20:  %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
sl21=sl20[:]:   %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
sl22=sl20[0:]:  %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
sl23=sl20[:len(sl20)]:  %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
sl24=sl20[:cap(sl20)]:  %T -> []int | %v -> [1 2 3 4 5 6 7 8 9 10] | %#v -> []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} | len=10 | cap=10
sl25=[0:len(sl20)]:     %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
sl26=[0:cap(sl20)]:     %T -> []int | %v -> [1 2 3 4 5 6 7 8 9 10] | %#v -> []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} | len=10 | cap=10
sl28=sl20[1:3]:         %T -> []int | %v -> [2 3] | %#v -> []int{2, 3} | len=2 | cap=9
sl29=sl20[1:4]:         %T -> []int | %v -> [2 3 4] | %#v -> []int{2, 3, 4} | len=3 | cap=9
sl30=sl20[2:4]:         %T -> []int | %v -> [3 4] | %#v -> []int{3, 4} | len=2 | cap=8
sl33=sl20[2:4:4]:       %T -> []int | %v -> [3 4] | %#v -> []int{3, 4} | len=2 | cap=2
sl34=sl20[2:4:5]:       %T -> []int | %v -> [3 4] | %#v -> []int{3, 4} | len=2 | cap=3
sl35=sl20[2:4:6]:       %T -> []int | %v -> [3 4] | %#v -> []int{3, 4} | len=2 | cap=4
sl36=sl20[2:4:7]:       %T -> []int | %v -> [3 4] | %#v -> []int{3, 4} | len=2 | cap=5
```

​	由上面给出的示例代码中的`sl24`和`sl26`，我们可以知道`sl20`这个切片的底层数组实际上就是`a2`。同时`a2`也是`sl21`到`sl36`的底层数组。

​	由以上示例在使用 `sl[low:high:max]`获取新切片时可以看出：`max >= high >= low` ；`max`不得大于底层数组的上边界所在的索引； 新切片的长度为`high - low`，而容量为`max - low`。

#### U修改

##### 修改元素

```go
sl37 := []int{1, 2, 3}
mfp.PrintFmtValWithLC("sl37", sl37, verbs)
sl37[0] = 11
mfp.PrintFmtValWithLC("sl37", sl37, verbs)
sl37[len(sl37)-1] = 33
mfp.PrintFmtValWithLC("sl37", sl37, verbs)
// 修改不存在的元素
//sl37[3] = 4 // 报错：panic: runtime error: index out of range [3] with length 3
```

```
sl37:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl37:   %T -> []int | %v -> [11 2 3] | %#v -> []int{11, 2, 3} | len=3 | cap=3
sl37:   %T -> []int | %v -> [11 2 33] | %#v -> []int{11, 2, 33} | len=3 | cap=3
```



##### 用整个切片赋值

```go
sl38 := []int{1, 2, 3}
mfp.PrintFmtValWithLC("1 sl38", sl38, verbs)
sl38 = []int{1, 2, 3, 4}
mfp.PrintFmtValWithLC("2 sl38", sl38, verbs)
sl38 = make([]int, 5, 10)
mfp.PrintFmtValWithLC("3 sl38", sl38, verbs)
sl38 = *new([]int)
mfp.PrintFmtValWithLC("4 sl38", sl38, verbs)
sl39 := []int{1, 2, 3, 4, 5, 6}
mfp.PrintFmtValWithLC("5 sl39", sl39, verbs)
sl38 = sl39
mfp.PrintFmtValWithLC("6 sl38", sl38, verbs)
```

```
1 sl38:         %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
2 sl38:         %T -> []int | %v -> [1 2 3 4] | %#v -> []int{1, 2, 3, 4} | len=4 | cap=4
3 sl38:         %T -> []int | %v -> [0 0 0 0 0] | %#v -> []int{0, 0, 0, 0, 0} | len=5 | cap=10
4 sl38:         %T -> []int | %v -> [] | %#v -> []int(nil) | len=0 | cap=0
5 sl39:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
6 sl38:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
```

##### 替换

###### 使用for循环

```go
sl73 := make([]int, 6, 10)
mfp.PrintFmtValWithLC("1 sl73", sl73, verbs)
// 将 sl73[0]~sl73[6]依次替换为 1~6
for k, _ := range sl73 {
    if k <= 6 {
        sl73[k] = k + 1
    }
}
mfp.PrintFmtValWithLC("2 sl73", sl73, verbs)
```

```
1 sl73:         %T -> []int | %v -> [0 0 0 0 0 0] | %#v -> []int{0, 0, 0, 0, 0, 0} | len=6 | cap=10
2 sl73:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
```

###### 使用slices.Replace函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl74 := make([]int, 6, 10)
mfp.PrintFmtValWithLC("1 sl74", sl74, verbs)
sl74 = slices.Replace(sl74, 0, 6, []int{1, 2, 3, 4, 5, 6}...)
mfp.PrintFmtValWithLC("2 sl74", sl74, verbs)
//sl74 = slices.Replace(sl74, 0, 7, []int{1, 2, 3, 4, 5, 6}...) // 报错：panic: runtime error: slice bounds out of range [7:6]
//mfp.PrintFmtValWithLC("3 sl74", sl74, verbs)
//sl74 = slices.Replace(sl74, 0, 7, []int{1, 2, 3, 4, 5, 6, 7}...) // 报错：panic: runtime error: slice bounds out of range [7:6]
//mfp.PrintFmtValWithLC("4 sl74", sl74, verbs)
```

```
1 sl74:         %T -> []int | %v -> [0 0 0 0 0 0] | %#v -> []int{0, 0, 0, 0, 0, 0} | len=6 | cap=10
2 sl74:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
```

​	这里的Replace函数的定义为`func Replace[S ~[]E, E any](s S, i, j int, v ...E) S`，结合以上示例，可以发现， `i`和`j` 的必须是在`[0, len(S)]` （包含`0`和`len(S)`）的范围内，否则报错。

##### 反转

###### 使用for循环

```go
func reverseSlice(slice []int) {
	length := len(slice)
    for i := 0; i < length/2; i++ {
        j := length - 1 - i
        slice[i], slice[j] = slice[j], slice[i]
    }
}
sl76 := []int{1, 2, 3, 4, 5, 6}
mfp.PrintFmtValWithLC("1 sl76", sl76, verbs)
reverseSlice(sl76)
mfp.PrintFmtValWithLC("2 sl76", sl76, verbs)
```

```
1 sl76:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
2 sl76:         %T -> []int | %v -> [6 5 4 3 2 1] | %#v -> []int{6, 5, 4, 3, 2, 1} | len=6 | cap=6
```



###### 使用slices.Reverse函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl77 := []int{1, 2, 3, 4, 5, 6}
mfp.PrintFmtValWithLC("1 sl77", sl77, verbs)
slices.Reverse(sl77)
mfp.PrintFmtValWithLC("2 sl77", sl77, verbs)
```

```
1 sl77:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
2 sl77:         %T -> []int | %v -> [6 5 4 3 2 1] | %#v -> []int{6, 5, 4, 3, 2, 1} | len=6 | cap=6
```

##### 移除

###### 移除未使用的容量

​	这里使用了`slices.Clip`函数，需要注意，`Clip的返回值`才是移除未使用的容量后的切片。

```go
fmt.Println("使用slices.Clip函数")
fmt.Println("从go1.21版本开始才可以使用")
sl78 := make([]int, 3, 6)
mfp.PrintFmtValWithLC("1 sl78", sl78, verbs)
sl78 = slices.Clip(sl78)
mfp.PrintFmtValWithLC("2 sl78", sl78, verbs)
```

```
1 sl78:         %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=6
2 sl78:         %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=3
```



##### 排序

```go

```



#### A访问

##### 直接访问指定索引下标的元素

```go
sl40 := []int{1, 2, 3}
fmt.Println("sl40[0]=", sl40[0])
fmt.Println("sl40[1]=", sl40[1])
fmt.Println("sl40[2]=", sl40[2])
```

```
sl40[0]= 1
sl40[1]= 2
sl40[2]= 3
```

##### 遍历切片

​	通过遍历的方式访问所需索引下标或全部索引下标的元素：

```go
for k, v := range sl40 {
    if k%2 == 0 {
        fmt.Println(k, "->", v)
    }
}
mfp.PrintHr()
for k, v := range sl40 {
    fmt.Println(k, "->", v)
}
```

```
0 -> 1
2 -> 3
------------------
0 -> 1
1 -> 2
2 -> 3
```

##### 复制切片

```go
slSrc43 := []int{1, 2, 3}
mfp.PrintFmtValWithLC("slSrc43", slSrc43, verbs)
slDst44 := make([]int, len(slSrc43))
mfp.PrintFmtValWithLC("slDst44", slDst44, verbs)

copy(slDst44, slSrc43) // func copy(dst []Type, src []Type) int
fmt.Println("使用copy函数")
slDst44[0] = 11
fmt.Println("slDst44[0] = 11 之后")
mfp.PrintFmtValWithLC("slDst43", slSrc43, verbs)
mfp.PrintFmtValWithLC("slDst44", slDst44, verbs)
slSrc43[1] = 22
fmt.Println("slSrc43[1] = 22 之后")
mfp.PrintFmtValWithLC("slDst43", slSrc43, verbs)
mfp.PrintFmtValWithLC("slDst44", slDst44, verbs)
```

```
slSrc43:        %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
slDst44:        %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=3
使用copy函数
slDst44[0] = 11 之后
slDst43:        %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
slDst44:        %T -> []int | %v -> [11 2 3] | %#v -> []int{11, 2, 3} | len=3 | cap=3
slSrc43[1] = 22 之后
slDst43:        %T -> []int | %v -> [1 22 3] | %#v -> []int{1, 22, 3} | len=3 | cap=3
slDst44:        %T -> []int | %v -> [11 2 3] | %#v -> []int{11, 2, 3} | len=3 | cap=3
```

​	可见，`copy`函数复制后产生的切片和源切片不共用底层数组！

##### 获取相关切片属性

```go
sl41 := []int{1, 2, 3}
fmt.Println("sl41切片的长度 len(sl41)=", len(sl41))
fmt.Println("sl41切片的容量 cap(sl41)=", cap(sl41))
```

```
sl41切片的长度 len(sl41)= 3
sl41切片的容量 cap(sl41)= 3
```

##### 判断相等

###### 是否可以使用`==`或`!=`?

​	=> 不可以！

```go
sl46 := []int{1, 2, 3}
sl47 := []int{1, 2, 3}
//fmt.Println("sl46 == sl47 -> ", sl46 == sl47) // 报错：invalid operation: sl46 == sl47 (slice can only be compared to nil)
//fmt.Println("sl46 != sl47 -> ", sl46 != sl47)// 报错：invalid operation: sl46 != sl47 (slice can only be compared to nil)
```

​	以上示例显示，在使用`==` 或 `!=` 时 切片 只可以和 `nil` 进行比较。

###### 使用slices.Equal函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl48 := []int{1, 2, 3}
sl49 := []int{1, 2, 3}
sl50 := []int{11, 2, 3}
sl51 := []int{1, 2, 3, 4}
fmt.Println("sl48 == sl49 -> ", slices.Equal(sl48, sl49))
fmt.Println("sl48 == sl50 -> ", slices.Equal(sl48, sl50))
fmt.Println("sl48 == sl51 -> ", slices.Equal(sl48, sl51))
```

```
sl48 == sl49 ->  true
sl48 == sl50 ->  false
sl48 == sl51 ->  false
```



###### 使用slices.EqualFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl52 := []int{1, 15, 8}
sl53 := []int{1, 15, 8}
sl54 := []int{11, 15, 8}
sl55 := []string{"01", "0x0f", "0o10"}

feq1 := func(e1, e2 int) bool {
    return e1 == e2
}
feq2 := func(e1 int, e2 string) bool {
    sn, err := strconv.ParseInt(e2, 0, 64)
    if err != nil {
        return false
    }
    return e1 == int(sn)
}
fmt.Println("sl52 == sl53 -> ", slices.EqualFunc(sl52, sl53, feq1))
fmt.Println("sl52 == sl54 -> ", slices.EqualFunc(sl52, sl54, feq1))
fmt.Println("sl52 == sl55 -> ", slices.EqualFunc(sl52, sl55, feq2))
```

```
sl52 == sl53 ->  true
sl52 == sl54 ->  false
sl52 == sl55 ->  true
```

##### 判断是否存在

###### 	使用for循环

```go
sl56 := []int{1, 2, 3}
forFunc := func(src []int, target int) bool {
    for _, v := range src {
        if v == target {
            return true
        }
    }
    return false
}

fmt.Println("1 在 sl56中 -> ", forFunc(sl56, 1))
fmt.Println("4 在 sl56中 -> ", forFunc(sl56, 4))
```

```
1 在 sl56中 ->  true
4 在 sl56中 ->  false
```

###### 使用slices.Contains函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl57 := []int{1, 2, 3}
fmt.Println("1 在 sl57中 -> ", slices.Contains(sl57, 1))
fmt.Println("4 在 sl57中 -> ", slices.Contains(sl57, 4))
```

```
1 在 sl57中 ->  true
4 在 sl57中 ->  false
```

###### 使用slices.ContainsFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl58 := []int{0, 42, -10, 8}

fmt.Println("sl58中存在负数 -> ", slices.ContainsFunc(sl58, func(e int) bool {
    return e < 0
}))
fmt.Println("sl58中存在奇数 -> ", slices.ContainsFunc(sl58, func(e int) bool {
    return e%2 == 1
}))
fmt.Println("sl58中存在 8 -> ", slices.ContainsFunc(sl58, func(e int) bool {
    return e == 8
}))
```

```
sl58中存在负数 ->  true
sl58中存在奇数 ->  false
sl58中存在 8 ->  true
```



##### 获取最大值

###### 使用for循环

```go
sl59 := []int{0, 42, -10, 8}

maxK := 0
maxV := sl59[0]
for k, v := range sl59 {
    if maxV < v {
        maxK = k
        maxV = v
    }
}
fmt.Printf("sl59中的最大值是sl59[%d]=%d\n", maxK, maxV)
```

```
sl59中的最大值是sl59[1]=42
```



###### 使用slices.Max函数

```go
fmt.Println("从go1.21版本开始才可以使用")

sl60 := []int{0, 42, -10, 8}
IamNaN := math.NaN()
sl61 := []float64{0, 42.12, -10.123, 8, IamNaN}
//sl62 := []int{0, 42, -10, 8, IamNaN} // 报错：cannot use IamNaN (variable of type float64) as int value in array or slice literal
fmt.Printf("sl60中的最大值是%d\n", slices.Max(sl60))

maxV2 := slices.Max(sl61)
fmt.Printf("sl61中的最大值是%f（%T）\n", maxV2, maxV2)
```

```
sl60中的最大值是42
sl61中的最大值是NaN（float64）
```



###### 使用slices.MaxFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl64 := []int{0, 42, -10, 8}
IamNaN2 := math.NaN()
sl65 := []float64{0, 42.12, -10.123, 8, IamNaN2}
fmt.Printf("sl64中最大值是%d\n", slices.MaxFunc(sl64, func(e1, e2 int) int {
    return cmp.Compare(e1, e2)
}))

fmt.Printf("sl65中最大值是%f\n", slices.MaxFunc(sl65, func(e1, e2 float64) int {
    return cmp.Compare(e1, e2)
}))

//sl66 := []int{}
//fmt.Printf("sl66中最大值是%d\n", slices.MaxFunc(sl66, func(e1, e2 int) int {
//	return cmp.Compare(e1, e2)
//})) // 报错：panic: slices.Max: empty list
```

```
sl64中最大值是42
sl65中最大值是42.120000
```



##### 获取最小值

###### 使用for循环

```go
func findMin[T1, T2 cmp.Ordered](minK T1, minV T2, src []T2) (T1, T2) {
	for k, v := range src {
		if minV > v {
			minK = T1(k)
			minV = v
		}
	}
	return minK, minV
}

sl67 := []int{0, 42, -10, 8}
minK1, minV1 := findMin(0, sl67[0], sl67)
fmt.Printf("sl67中的最小值是sl67[%d]=%d\n", minK1, minV1)

sl68 := []float64{0, 42.12, -10.123, 8}
minK2, minV2 := findMin(0, sl68[0], sl68)
fmt.Printf("sl68中的最小值是sl68[%d]=%f\n", minK2, minV2)

IamNaN3 := math.NaN()
sl69 := []float64{0, 42.12, -10.123, 8, IamNaN3}
minK3, minV3 := findMin(0, sl69[0], sl69)
fmt.Printf("sl69中的最小值是sl69[%d]=%f\n", minK3, minV3)
```

```
sl67中的最小值是sl67[2]=-10
sl68中的最小值是sl68[2]=-10.123000
sl69中的最小值是sl69[2]=-10.123000
```

###### 使用slices.Min函数

```go
fmt.Println("从go1.21版本开始才可以使用")

sl70 := []int{0, 42, -10, 8}
sl71 := []float64{0, 42.12, -10.123, 8}
IamNaN4 := math.NaN()
sl72 := []float64{0, 42.12, -10.123, 8, IamNaN4}
fmt.Println("sl70中的最小值是", slices.Min(sl70))
fmt.Println("sl71中的最小值是", slices.Min(sl71))
fmt.Println("sl72中的最小值是", slices.Min(sl72))
```

```
sl70中的最小值是 -10
sl71中的最小值是 -10.123
sl72中的最小值是 NaN
```



###### 使用slices.MinFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")

sl70 := []int{0, 42, -10, 8}
sl71 := []float64{0, 42.12, -10.123, 8}
IamNaN4 := math.NaN()
sl72 := []float64{0, 42.12, -10.123, 8, IamNaN4}
fmt.Println("sl70中的最小值是", slices.MinFunc(sl70, func(a, b int) int {
	return cmp.Compare(a, b)
}))
fmt.Println("sl71中的最小值是", slices.MinFunc(sl71, func(a, b float64) int {
	return cmp.Compare(a, b)
}))
fmt.Println("sl72中的最小值是", slices.MinFunc(sl72, func(a, b float64) int {
	return cmp.Compare(a, b)
}))
```

```
sl70中的最小值是 -10
sl71中的最小值是 -10.123
sl72中的最小值是 NaN
```





###### 使用slices.Replace函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl74 := make([]int, 6, 10)
mfp.PrintFmtValWithLC("1 sl74", sl74, verbs)
sl74 = slices.Replace(sl74, 0, 6, []int{1, 2, 3, 4, 5, 6}...)
mfp.PrintFmtValWithLC("2 sl74", sl74, verbs)
//sl74 = slices.Replace(sl74, 0, 7, []int{1, 2, 3, 4, 5, 6}...) // 报错：panic: runtime error: slice bounds out of range [7:6]
//mfp.PrintFmtValWithLC("3 sl74", sl74, verbs)
//sl74 = slices.Replace(sl74, 0, 7, []int{1, 2, 3, 4, 5, 6, 7}...) // 报错：panic: runtime error: slice bounds out of range [7:6]
//mfp.PrintFmtValWithLC("4 sl74", sl74, verbs)
```

```
1 sl74:         %T -> []int | %v -> [0 0 0 0 0 0] | %#v -> []int{0, 0, 0, 0, 0, 0} | len=6 | cap=10
2 sl74:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
```



##### 排序



###### 



#### D删除

##### 是否可以删除某一元素呢？

​	=> 可以！

```go
sl42 := []int{1, 2, 3, 4, 5, 6}
i := 3 // 需要删除元素的索引下标
mfp.PrintFmtValWithLC("1 sl42", sl42, verbs)
sl42 = append(sl42[0:i], sl42[i+1:]...) // 删除 索引为3的元素
mfp.PrintFmtValWithLC("2 sl42", sl42, verbs)
sl42 = append(sl42[0:i], sl42[i+1:]...) // 删除 当前索引为3的元素
mfp.PrintFmtValWithLC("3 sl42", sl42, verbs)
i = 0
sl42 = append(sl42[0:i], sl42[i+1:]...) // 删除 当前索引为0的元素
mfp.PrintFmtValWithLC("4 sl42", sl42, verbs)
sl42 = sl42[:len(sl42) - 1] // 删除当前的最后一个元素
mfp.PrintFmtValWithLC("5 sl42", sl42, verbs)
sl42 = sl42[1:] // 删除当前的第一个元素
mfp.PrintFmtValWithLC("6 sl42", sl42, verbs)
```

```
1 sl42:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
2 sl42:         %T -> []int | %v -> [1 2 3 5 6] | %#v -> []int{1, 2, 3, 5, 6} | len=5 | cap=6
3 sl42:         %T -> []int | %v -> [1 2 3 6] | %#v -> []int{1, 2, 3, 6} | len=4 | cap=6
4 sl42:         %T -> []int | %v -> [2 3 6] | %#v -> []int{2, 3, 6} | len=3 | cap=6
4 sl42:         %T -> []int | %v -> [2 3 6] | %#v -> []int{2, 3, 6} | len=3 | cap=6
5 sl42:         %T -> []int | %v -> [2 3] | %#v -> []int{2, 3} | len=2 | cap=6
6 sl42:         %T -> []int | %v -> [3] | %#v -> []int{3} | len=1 | cap=5
```

#### 作为实参传递给函数或方法

​	在 Go 语言中，`切片是引用类型`。切片本身是一个包含指向底层数组的指针、长度和容量的数据结构。当你将一个切片赋值给另一个切片，或者将一个切片作为函数参数传递时，实际上是传递了切片的引用，而不是切片的副本。因此，对切片的修改会影响到原始切片以及引用同一底层数组的其他切片。

​	切片作为函数参数传递时，由于只是传递了切片的引用，而不是整个切片的副本，因此`在性能和内存上并不会有大开销`。即使切片的长度很大，传递切片的引用也只是传递了指向底层数组的指针、长度和容量这几个值，并不会复制整个底层数组。因此，切片作为实参通常不会产生额外的内存开销。

​	需要注意的是，如果在函数内部修改了切片的长度或容量，可能会导致底层数组重新分配内存，从而产生额外的内存开销。但是这种情况并不是切片本身作为实参引起的，而是对切片的修改引起的。

#### 易混淆的知识点

​	

#### 易错点

##### 访问最后一个切片元素

​	直接用`sl[len(sl)]`访问切片`sl`的最后一个元素 => 肯定报错！

```go
sl45 := []int{1, 2, 3}
//fmt.Println(sl45[len(sl45)])   // 报错：panic: runtime error: index out of range [3] with length 3
fmt.Println(sl45[len(sl45)-1]) // 正确方式
```

```
3
```

##### 长度和容量不一致的切片

​	长度和容量不一致时，给索引`i`的范围是`len(sl) <= i <= cap(sl)` 的元素赋值，以为可以增加切片的长度，实际却是 `panic`。 

```go
sl75 := make([]int, 3, 6)
mfp.PrintFmtValWithLC("1 sl75", sl75, verbs)
//sl75[3] = 4 // 报错：panic: runtime error: index out of range [3] with length 3
//mfp.PrintFmtValWithLC("2 sl75", sl75, verbs)
//sl75[4] = 5 // 报错：panic: runtime error: index out of range [4] with length 3
//mfp.PrintFmtValWithLC("3 sl75", sl75, verbs)
```

```
1 sl75:         %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=6
```

##### 使用slices.Replace函数

​	使用slices.Replace函数：`func Replace[S ~[]E, E any](s S, i, j int, v ...E) S ` 时，将 `i`和 `j`设置成一样，以为只会替换索引`i`这一处的元素值，而实际上却是往索引`i`前面插入一个新的元素值`v`。

```go

```



### map

#### C创建

##### 直接创建

```go
var m1 map[int]int
var m2 map[string]int = map[string]int{"A": 1, "B": 2}
var m3 = map[string]int{"A": 1, "B": 2}
m4 := map[string]int{"A": 1, "B": 2}
mfp.PrintFmtValWithL("m1", m1, verbs)
mfp.PrintFmtValWithL("m2", m2, verbs)
mfp.PrintFmtValWithL("m3", m3, verbs)
mfp.PrintFmtValWithL("m4", m4, verbs)
```

```
m1:     %T -> map[int]int | %v -> map[] | %#v -> map[int]int(nil) | len=0
m2:     %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
m3:     %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
m4:     %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
```



##### 用make创建

```go
m5 := make(map[string]int)
m6 := make(map[string]int, 3)
//m7 := make(map[string]int, 3, 3) // 报错：invalid operation: make(map[string]int, 3, 3) expects 1 or 2 arguments; found 3
mfp.PrintFmtValWithL("1 m5", m5, verbs)
mfp.PrintFmtValWithL("2 m6", m6, verbs)
//mfp.PrintFmtValWithL("m7", m7, verbs)
```

```
1 m5:   %T -> map[string]int | %v -> map[] | %#v -> map[string]int{} | len=0
2 m6:   %T -> map[string]int | %v -> map[] | %#v -> map[string]int{} | len=0
```



##### 用new创建

```go
m7 := *new(map[string]int)
mfp.PrintFmtValWithL("m7", m7, verbs)

//m7["A"] = 1 // 报错：panic: assignment to entry in nil map
//mfp.PrintFmtValWithL("m7", m7, verbs)

m7 = map[string]int{"A": 1}
mfp.PrintFmtValWithL("m7", m7, verbs)
```

```
m7:     %T -> map[string]int | %v -> map[] | %#v -> map[string]int(nil) | len=0
m7:     %T -> map[string]int | %v -> map[A:1] | %#v -> map[string]int{"A":1} | len=1
```



#### U修改

##### 修改元素

```
m9 := map[string]int{"A": 1, "B": 2, "C": 3}
mfp.PrintFmtValWithL("1 m9", m9, verbs)
m9["A"] = 11
mfp.PrintFmtValWithL("2 m9", m9, verbs)
m9["D"] = 4 // 修改不存在的Key
mfp.PrintFmtValWithL("3 m9", m9, verbs)
```

```
1 m9:   %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
2 m9:   %T -> map[string]int | %v -> map[A:11 B:2 C:3] | %#v -> map[string]int{"A":11, "B":2, "C":3} | len=3
3 m9:   %T -> map[string]int | %v -> map[A:11 B:2 C:3 D:4] | %#v -> map[string]int{"A":11, "B":2, "C":3, "D":4} | len=4
```



##### 用整个map赋值

```go
m10 := map[string]int{"A": 1, "B": 2, "C": 3}
mfp.PrintFmtValWithL("1 m10", m10, verbs)
m10 = map[string]int{"A": 11, "B": 22, "C": 33, "D": 44}
mfp.PrintFmtValWithL("2 m10", m10, verbs)
m11 := map[string]int{"A": 111, "B": 222, "C": 333, "D": 444}
m10 = m11
mfp.PrintFmtValWithL("3 m10", m10, verbs)
m11["A"] = 1
mfp.PrintFmtValWithL("4 m10", m10, verbs)
```

```
1 m10:  %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
2 m10:  %T -> map[string]int | %v -> map[A:11 B:22 C:33 D:44] | %#v -> map[string]int{"A":11, "B":22, "C":33, "D":44} | len=4
3 m10:  %T -> map[string]int | %v -> map[A:111 B:222 C:333 D:444] | %#v -> map[string]int{"A":111, "B":222, "C":333, "D":444} | len=4
4 m10:  %T -> map[string]int | %v -> map[A:1 B:222 C:333 D:444] | %#v -> map[string]int{"A":1, "B":222, "C":333, "D":444} | len=4
```



#### A访问

##### 直接访问指定Key的元素

```go
m12 := map[string]int{"A": 1, "B": 2, "C": 3}
fmt.Println(m12["A"])
fmt.Println(m12["B"])
fmt.Println(m12["C"])
fmt.Println(m12["D"])// 访问不存在的Key
```

```
1
2
3
0
```



##### 遍历map

```go
for k,v := range m12 {
    fmt.Println(k,"->", v)
}
```

```
A -> 1
B -> 2
C -> 3
```

​	需要注意的是，遍历是无序的，每一次的遍历顺序都有可能不同！

##### 复制map

```go
fmt.Println("从go1.21版本开始才可以使用")

fmt.Println("使用maps.Clone函数")
m13 := map[string]int{"A": 1, "B": 2, "C": 3}
mfp.PrintFmtValWithL("1 m13", m13, verbs)
m14 := maps.Clone(m13)
mfp.PrintFmtValWithL("2 m14", m14, verbs)

m13["A"] = 11
fmt.Println(`修改 m13["A"] = 11`)
mfp.PrintFmtValWithL("3 m13", m13, verbs)
mfp.PrintFmtValWithL("4 m14", m14, verbs)

m14["B"] = 22
fmt.Println(`修改 m14["B"] = 22`)
mfp.PrintFmtValWithL("5 m13", m13, verbs)
mfp.PrintFmtValWithL("6 m14", m14, verbs)
mfp.PrintHr()

fmt.Println("使用maps.Copy函数")
m15 := map[string]int{"A": 1, "B": 2}
m16 := map[string]int{"A": 11, "C": 33}

fmt.Println(`使用Copy函数前`)
mfp.PrintFmtValWithL("m15", m15, verbs)
mfp.PrintFmtValWithL("m16", m16, verbs)
maps.Copy(m16, m15) // func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2)

fmt.Println(`使用Copy函数后`)
mfp.PrintFmtValWithL("m15", m15, verbs)
mfp.PrintFmtValWithL("m16", m16, verbs)

m15["A"] = 111
fmt.Println(`修改 m15["A"] = 111`)
mfp.PrintFmtValWithL("m15", m15, verbs)
mfp.PrintFmtValWithL("m16", m16, verbs)

m16["B"] = 222
fmt.Println(`修改 m16["B"] = 222`)
mfp.PrintFmtValWithL("m15", m15, verbs)
mfp.PrintFmtValWithL("m16", m16, verbs)
```

```
从go1.21版本开始可使用
使用maps.Clone函数
1 m13:  %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
2 m14:  %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
修改 m13["A"] = 11
3 m13:  %T -> map[string]int | %v -> map[A:11 B:2 C:3] | %#v -> map[string]int{"A":11, "B":2, "C":3} | len=3
4 m14:  %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
修改 m14["B"] = 22
5 m13:  %T -> map[string]int | %v -> map[A:11 B:2 C:3] | %#v -> map[string]int{"A":11, "B":2, "C":3} | len=3
6 m14:  %T -> map[string]int | %v -> map[A:1 B:22 C:3] | %#v -> map[string]int{"A":1, "B":22, "C":3} | len=3
------------------
使用maps.Copy函数
使用Copy函数前
m15:    %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
m16:    %T -> map[string]int | %v -> map[A:11 C:33] | %#v -> map[string]int{"A":11, "C":33} | len=2
使用Copy函数后
m15:    %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
m16:    %T -> map[string]int | %v -> map[A:1 B:2 C:33] | %#v -> map[string]int{"A":1, "B":2, "C":33} | len=3
修改 m15["A"] = 111
m15:    %T -> map[string]int | %v -> map[A:111 B:2] | %#v -> map[string]int{"A":111, "B":2} | len=2
m16:    %T -> map[string]int | %v -> map[A:1 B:2 C:33] | %#v -> map[string]int{"A":1, "B":2, "C":33} | len=3
修改 m16["B"] = 222
m15:    %T -> map[string]int | %v -> map[A:111 B:2] | %#v -> map[string]int{"A":111, "B":2} | len=2
m16:    %T -> map[string]int | %v -> map[A:1 B:222 C:33] | %#v -> map[string]int{"A":1, "B":222, "C":33} | len=3
```



##### 获取相关map属性

```go
fmt.Println("m12 map的长度 len(m12)=", len(m12))
```

```
m12 map的长度 len(m12)= 3
```

##### 判断相等

###### 是否可以使用==或 !=？

​	=> 不可以！

```go
m18 := map[string]int{"A": 1, "B": 2, "C": 3}
m19 := map[string]int{"A": 1, "B": 2, "C": 3}
//fmt.Println("m18 == m19 -> ", m18 == m19) // 报错：invalid operation: m18 == m19 (map can only be compared to nil)
//fmt.Println("m18 != m19 -> ", m18 != m19) // 报错：invalid operation: m18 != m19 (map can only be compared to nil)
```

​	以上示例显示，在使用`==` 或 `!=` 时 map 只可以和 `nil` 进行比较。

###### 使用maps.Equal函数

```go
fmt.Println("从go1.21版本开始才可以使用")

m20 := map[string]int{"A": 1, "B": 2}
m21 := map[string]int{"A": 1, "B": 2}
fmt.Println("m20 == m21 ->", maps.Equal(m20, m21))

m22 := map[string]int{"A": 11, "B": 2}
fmt.Println("m20 == m22 ->", maps.Equal(m20, m22))

m23 := map[string]int{"A": 1, "B": 2, "C": 3}
fmt.Println("m20 == m23 ->", maps.Equal(m20, m23))
```

```
m20 == m21 -> true
m20 == m22 -> false
m20 == m23 -> false
```

###### 使用maps.EqualFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")
m24 := map[string]int{"A": 1, "B": 2}
m25 := map[string]int{"A": 1, "B": 2}
fmt.Println("m24 == m25 -> ", maps.EqualFunc(m24, m25, func(v1 int, v2 int) bool {
    if v1 == v2 {
        return true
    }
    return false
}))
```

```
m24 == m25 ->  true
```



#### D删除

##### 是否可以删除map中的某一元素？

​	=> 可以！

###### 使用delete函数

```go
m8 := map[string]int{"A": 1, "B": 2, "C": 3}
mfp.PrintFmtValWithL("m8", m8, verbs)
delete(m8, "A")
mfp.PrintFmtValWithL("m8", m8, verbs)
delete(m8, "A") // 重复删除，也不会报错
mfp.PrintFmtValWithL("m8", m8, verbs)
delete(m8, "B")
mfp.PrintFmtValWithL("m8", m8, verbs)
delete(m8, "C")
mfp.PrintFmtValWithL("m8", m8, verbs)
```

```
m8:     %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
m8:     %T -> map[string]int | %v -> map[B:2 C:3] | %#v -> map[string]int{"B":2, "C":3} | len=2
m8:     %T -> map[string]int | %v -> map[B:2 C:3] | %#v -> map[string]int{"B":2, "C":3} | len=2
m8:     %T -> map[string]int | %v -> map[C:3] | %#v -> map[string]int{"C":3} | len=1
m8:     %T -> map[string]int | %v -> map[] | %#v -> map[string]int{} | len=0
```

###### 使用maps.DeleteFunc函数

```go
m17 := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}
fmt.Println("使用maps.DeleteFunc函数前")
mfp.PrintFmtValWithL("m17", m17, verbs)
maps.DeleteFunc(m17, func(k string, v int) bool {
    if v%2 == 1 {
        return true
    }
    return false
})

fmt.Println("使用maps.DeleteFunc函数后")
mfp.PrintFmtValWithL("m17", m17, verbs)
```

```
使用maps.DeleteFunc函数前
m17:    %T -> map[string]int | %v -> map[A:1 B:2 C:3 D:4] | %#v -> map[string]int{"A":1, "B":2, "C":3, "D":4} | len=4
使用maps.DeleteFunc函数后
m17:    %T -> map[string]int | %v -> map[B:2 D:4] | %#v -> map[string]int{"B":2, "D":4} | len=2
```



#### 作为实参传递给函数或方法

​	在 Go 语言中，`map 是引用类型`。当你将一个 map 赋值给另一个变量，或者将一个 map 作为函数参数传递时，实际上是传递了 map 的引用，而不是整个 map 的副本。因此，对 map 的修改会影响到原始 map 以及引用同一个 map 的其他变量。

​	`map 作为函数参数传递时，并不会产生大的性能和内存开销`。与切片类似，虽然 map 可能包含大量的键值对，但传递 map 的引用只是传递了指向底层数据结构的指针，而不是复制整个底层数据结构。因此，map 作为实参传递通常不会产生额外的内存开销。

​	需要注意的是，在并发编程中，对 map 的并发访问可能会导致竞态条件，因此在多个 goroutine 中共享 map 时，需要使用适当的同步机制（例如 sync.Mutex 或 sync.RWMutex）来保护 map 的访问。

#### 易混淆的知识点



#### 易错点

new函数创建的map直接进行

##### 以为可以使用copy函数来复制一个map





















































































































