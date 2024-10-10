+++
title = "类型"
date = 2024-07-13T10:57:15+08:00
weight = 200
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

## 预先声明的类型（或内置类型）

​	合计27种。

```go
布尔型【共1种】：
bool

数值型【共16种】：
uint8 uint16 uint32 uint64 uint
int6  int16  int32  int64  int
byte【uint8的别名】   rune【int32的别名】
float32      float64
complex64    complex128

字符串型【共1种】
string

数组型【共1种】
[N]T【N为元素个数，T为元素类型】

切片型【共1种】
[]T【T为元素类型】

结构体型【共1种】：
struct{...fields fieldsType}

指针型【共1种】：
*T

指针地址型/指针整数型【共1种】：
uintptr

函数型【共1种】：
func(...params paramsType) ...returns returnsType 

接口型【共1种】：
interface{...methods}

map型【共1种】：
map[keyType]valueType

chan型【共1种】：
chan T【T为通道中所传递的值类型】

```



### Go的内置数据类型相关信息

| 序号 | 数据类型    | 名称       | 别名 | 默认值                      | 占用字节数 | 数据最小值                                                   | 数据最大值                                                   | 备注                                                         | 链接 |
| ---- | ----------- | ---------- | ---- | --------------------------- | ---------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ---- |
| 1    | bool        | 布尔型     |      | false                       | 1          | 无                                                           | 无                                                           | 只有true和false两种值                                        |      |
| 2    | byte        | 字节型     |      | '\x00'（使用%q）0（使用%d） | 1          | '\x00'（使用%q）0（使用%d）                                  | '\x7f'（使用%q）127（使用%d）                                | 是uint8的别名                                                |      |
| 3    | complex64   | 复数型     |      | (0+0i)                      | 8          | 无                                                           | 无                                                           |                                                              |      |
| 4    | complex128  | 复数型     |      | (0+0i)                      | 16         | 无                                                           | 无                                                           |                                                              |      |
| 5    | float32     | 浮点型     |      | 0                           | 4          | 1.401298464324817070923729583289916131280e-45 最小正非零值（使用%.39e保留39位小数） | 3.40282346638528859811704183484516925440e+38 （使用%.38e保留38位小数） |                                                              |      |
| 6    | float64     | 浮点型     |      | 0                           | 8          | 4.94065645841246544176568792868221372365059803e-324 最小正非零值（使用%.44e保留44位小数） | 1.797693134862315708145274237317043567980706e+308 （使用%.42e保留42位小数） |                                                              |      |
| 7    | int8        | 有符号整型 |      | 0                           | 1          | -128                                                         | 127                                                          |                                                              |      |
| 8    | int16       | 有符号整型 |      | 0                           | 2          | -32768                                                       | 32767                                                        |                                                              |      |
| 9    | int32       | 有符号整型 | rune | 0                           | 4          | -2147483648                                                  | 2147483647                                                   |                                                              |      |
| 10   | int64       | 有符号整型 |      | 0                           | 8          | -9223372036854775808                                         | 9223372036854775807（>922亿亿）                              |                                                              |      |
| 11   | int         | 有符号整型 |      | 0                           | 8          | -9223372036854775808                                         | 9223372036854775807（>922亿亿）                              | 请注意：这里给出的是64位系统的情况！                         |      |
| 12   | uint8       | 无符号整型 | byte | 0                           | 1          | 0                                                            | 255                                                          |                                                              |      |
| 13   | uint16      | 无符号整型 |      | 0                           | 2          | 0                                                            | 65535                                                        |                                                              |      |
| 14   | uint32      | 无符号整型 |      | 0                           | 4          | 0                                                            | 4294967295                                                   |                                                              |      |
| 15   | uint64      | 无符号整型 |      | 0                           | 8          | 0                                                            | 18446744073709551615                                         |                                                              |      |
| 16   | uint        | 无符号整型 |      | 0                           | 8          | 0                                                            | 18446744073709551615                                         |                                                              |      |
| 17   | rune        | 符文型     |      | '\x00'（使用%q）0（使用%d） | 4          | '\x00'（使用%q）0（使用%d）                                  |                                                              | 是int32的别名，而非uint32的别名                              |      |
| 18   | uintptr     | 指针整数型 |      | 0                           |            |                                                              |                                                              | uintptr 是一个整数类型，它足够大，可以容纳任何指针的比特模式 |      |
| 19   | string      | 字符串型   |      | ""                          |            |                                                              |                                                              |                                                              |      |
| 20   | [n]T        | 数组       |      | 空数组                      |            |                                                              |                                                              |                                                              |      |
| 21   | []T         | 切片       |      | nil                         |            |                                                              |                                                              |                                                              |      |
| 22   | map[K]V     | 映射       |      | nil                         |            |                                                              |                                                              |                                                              |      |
| 23   | struct{...} | 结构体     |      | 各自字段的零值              |            |                                                              |                                                              |                                                              |      |
| 24   | chan T      | 通道       |      | nil                         |            |                                                              |                                                              |                                                              |      |
| 25   | *T          | 指针       |      | nil                         |            |                                                              |                                                              |                                                              |      |
| 26   | interface   | 接口       |      | nil                         |            |                                                              |                                                              |                                                              |      |
| 27   | error       | 错误类型   |      | 无                          |            |                                                              |                                                              |                                                              |      |

## 获取内置类型的默认值

```go
package main

import "fmt"

func main() {
	fmt.Println("default value:")
	var b bool
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64
	var ui uint
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var i int
	var s string
	var ai [3]int
	var as [3]string
	var ab [3]bool
	var sli []int
	var sls []string
	var slb []bool
	type St struct {
		a int
		b string
		c *int
		d chan int
	}
	var st St
	var pti *int
	var pts *string
	var ptb *bool
	var uptr uintptr
	var fti func(int, int) int
	var fts func(string, string) string
	var ftb func(bool, bool) bool
	type ITF interface {
		M1(int) int
		M2(string) string
		M3(bool) bool
	}
	var itf ITF
	var mii map[int]int
	var mss map[string]string
	var mbb map[bool]bool
	var chi chan int
	var chs chan string
	var chb chan bool

	fmt.Printf("bool -> %t\n", b)
	fmt.Printf("uint8 -> %d\n", ui8)
	fmt.Printf("uint16 -> %d\n", ui16)
	fmt.Printf("uint32 -> %d\n", ui32)
	fmt.Printf("uint64 -> %d\n", ui64)
	fmt.Printf("uint -> %d\n", ui)
	fmt.Printf("int8 -> %d\n", i8)
	fmt.Printf("int16 -> %d\n", i16)
	fmt.Printf("int32 -> %d\n", i32)
	fmt.Printf("int64 -> %d\n", i64)
	fmt.Printf("int -> %d\n", i)
	fmt.Printf("string -> %q\n", s)
	fmt.Printf("array [3]int -> %#v\n", ai)
	fmt.Printf("array [3]string -> %#v\n", as)
	fmt.Printf("array [3]bool -> %#v\n", ab)
	fmt.Printf("slice []int -> %#v\n", sli)
	fmt.Printf("slice []int -> %v\n", sli)
	if sli == nil {
		fmt.Println("sli is nil.")
	}
	fmt.Printf("slice []string -> %#v\n", sls)
	fmt.Printf("slice []string -> %v\n", sls)
	if sls == nil {
		fmt.Println("sls is nil.")
	}
	fmt.Printf("slice []bool -> %#v\n", slb)
	fmt.Printf("slice []bool -> %v\n", slb)
	if slb == nil {
		fmt.Println("slb is nil.")
	}
	fmt.Printf("struct st -> %#v\n", st)
	fmt.Printf("struct st -> %v\n", st)
	fmt.Printf("point *int -> %v\n", pti)
	if pti == nil {
		fmt.Println("pti is nil.")
	}
	fmt.Printf("point *string -> %v\n", pts)
	if pts == nil {
		fmt.Println("pts is nil.")
	}
	fmt.Printf("point *bool -> %v\n", ptb)
	if ptb == nil {
		fmt.Println("ptb is nil.")
	}
	fmt.Printf("uintptr -> %v\n", uptr)
	fmt.Printf("func(int,int) int -> %v\n", fti)
	if fti == nil {
		fmt.Println("fti is nil.")
	}
	fmt.Printf("func(string,string) string -> %v\n", fts)
	if fts == nil {
		fmt.Println("fts is nil.")
	}

	fmt.Printf("func(bool,bool) bool -> %v\n", ftb)
	if ftb == nil {
		fmt.Println("ftb is nil.")
	}

	fmt.Printf("interface -> %v\n", itf)
	if itf == nil {
		fmt.Println("itf is nil.")
	}

	fmt.Printf("map[int]int -> %#v\n", mii)
	fmt.Printf("map[int]int -> %v\n", mii)
	if mii == nil {
		fmt.Println("mii is nil.")
	}
	fmt.Printf("map[string]string -> %#v\n", mss)
	fmt.Printf("map[string]string -> %v\n", mss)
	if mss == nil {
		fmt.Println("mss is nil.")
	}
	fmt.Printf("map[bool]bool -> %#v\n", mbb)
	fmt.Printf("map[bool]bool -> %v\n", mbb)
	if mbb == nil {
		fmt.Println("mbb is nil.")
	}

	fmt.Printf("chan int -> %#v\n", chi)
	fmt.Printf("chan int -> %v\n", chi)
	if chi == nil {
		fmt.Println("chi is nil.")
	}
	fmt.Printf("chan string -> %#v\n", chs)
	fmt.Printf("chan string -> %v\n", chs)
	if chs == nil {
		fmt.Println("chs is nil.")
	}
	fmt.Printf("chan bool -> %#v\n", chb)
	fmt.Printf("chan bool -> %v\n", chb)
	if chb == nil {
		fmt.Println("chb is nil.")
	}
}
//default value:
//bool -> false
//uint8 -> 0
//uint16 -> 0
//uint32 -> 0
//uint64 -> 0
//uint -> 0
//int8 -> 0
//int16 -> 0
//int32 -> 0
//int64 -> 0
//int -> 0
//string -> ""
//array [3]int -> [3]int{0, 0, 0}
//array [3]string -> [3]string{"", "", ""}
//array [3]bool -> [3]bool{false, false, false}
//slice []int -> []int(nil)
//slice []int -> []
//sli is nil.
//slice []string -> []string(nil)
//slice []string -> []
//sls is nil.
//slice []bool -> []bool(nil)
//slice []bool -> []
//slb is nil.
//struct st -> main.St{a:0, b:"", c:(*int)(nil), d:(chan int)(nil)}
//struct st -> {0  <nil> <nil>}
//point *int -> <nil>
//pti is nil.
//point *string -> <nil>
//pts is nil.
//point *bool -> <nil>
//ptb is nil.
//uintptr -> 0
//func(int,int) int -> <nil>
//fti is nil.
//func(string,string) string -> <nil>
//fts is nil.
//func(bool,bool) bool -> <nil>
//ftb is nil.
//interface -> <nil>
//itf is nil.
//map[int]int -> map[int]int(nil)
//map[int]int -> map[]
//mii is nil.
//map[string]string -> map[string]string(nil)
//map[string]string -> map[]
//mss is nil.
//map[bool]bool -> map[bool]bool(nil)
//map[bool]bool -> map[]
//mbb is nil.
//chan int -> (chan int)(nil)
//chan int -> <nil>
//chi is nil.
//chan string -> (chan string)(nil)
//chan string -> <nil>
//chs is nil.
//chan bool -> (chan bool)(nil)
//chan bool -> <nil>
//chb is nil.
```

