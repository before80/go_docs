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

var verbs = []string{"T", "v", "+v", "#v"}

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
全局变量 gb1:   %T -> bool | %v -> true | %+v -> true | %#v -> true |    
全局变量 gb2:   %T -> bool | %v -> false | %+v -> false | %#v -> false | 
---init 执行完成后---                                                    
全局变量 gb1:   %T -> bool | %v -> false | %+v -> false | %#v -> false | 
---局部变量---                                                           
声明方式1 b1:   %T -> bool | %v -> false | %+v -> false | %#v -> false | 
赋值后:         %T -> bool | %v -> true | %+v -> true | %#v -> true |    
赋值后:         %T -> bool | %v -> false | %+v -> false | %#v -> false | 
声明方式2 b2:   %T -> bool | %v -> true | %+v -> true | %#v -> true |    
声明方式3 b3:   %T -> bool | %v -> true | %+v -> true | %#v -> true | 
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
	var bt1 byte // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 false
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
	var c1281 complex128 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 false
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

```

{{< /tab  >}}

{{< tab header="int/int8/16/32/64" >}}

```go

```

{{< /tab  >}}

{{< tab header="uint/uint8/16/32/64" >}}

```go

```

{{< /tab  >}}

{{< tab header="rune" >}}

```go

```

{{< /tab  >}}

{{< tab header="uintptr" >}}

```go

```

{{< /tab  >}}

{{< tab header="string" >}}

```go

```

{{< /tab  >}}

{{< tab header="数组" >}}

```go

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



