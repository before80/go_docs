+++
title = "array"
date = 2024-07-13T14:04:33+08:00
weight = 400
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 说明
>
> ​	mfp包来自：github.com/before80/utils/mfp

## 数组的底层数据结构



## C创建

### 一维数组

#### 直接创建

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

#### 是否可以通过make创建？

​	=> 不可以

```go
//a4 := make([3]int{1, 2, 3}, 3, 3) // 报错：[3]int{…} is not a type
//a4 := make([3]int{}, 3, 3) // 报错：[3]int{} is not a type
//a4 := make([3]int, 3, 3) // 报错：invalid argument: cannot make [3]int; type must be slice, map, or channel
//a4 := make([3]int) // 报错： invalid argument: cannot make [3]int; type must be slice, map, or channel	
```

#### 是否可以通过new创建？

​	=> 可以

```go
a5 := new([3]int)
mfp.PrintFmtVal("a5", a5, verbs)
mfp.PrintFmtVal("*a5", *a5, verbs)
```

```
a5:     %T -> *[3]int | %v -> &[0 0 0] | %#v -> &[3]int{0, 0, 0}
*a5:    %T -> [3]int | %v -> [0 0 0] | %#v -> [3]int{0, 0, 0}
```



### 多维数组

​	从一维数组的创建中可以推测出，多维数组的创建也是不可以通过`make`进行创建，`new`可以进行创建。

​	根据以下代码和运行结果，我们还可以发现在使用数组字面量直接创建多维数组的时候，只能在第一维的位置上使用`...`，而不能在后续维度上使用，否则编译时报错。

```go
fmt.Println("多维数组")
fmt.Println("二维数组")
var a6 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
var a7 [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
var a8 = [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
//var a9 = [...][...]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // 报错：invalid use of [...] array (outside a composite literal)
ad2 := [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
ad20 := new([3][3]int)
mfp.PrintFmtVal("a6", a6, verbs)
mfp.PrintFmtVal("a7", a7, verbs)
mfp.PrintFmtVal("a8", a8, verbs)
//mfp.PrintFmtVal("a9", a9, verbs)
mfp.PrintFmtVal("ad2", ad2, verbs)
mfp.PrintFmtVal("ad20", ad20, verbs)
mfp.PrintFmtVal("*ad20", *ad20, verbs)

fmt.Println("三维数组")
var a10 = [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
var a11 [2][2][2]int = [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
var a12 = [...][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
//var a13 = [...][...][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}} // 报错：invalid use of [...] array (outside a composite literal) 以及 missing type in composite literal
ad3 := [...][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
//ad3x1 := [2][...][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}} // 报错：invalid use of [...] array (outside a composite literal) 以及 missing type in composite literal
//ad3x2 := [2][2][...]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}// 报错：invalid use of [...] array (outside a composite literal)
//ad3x3 := [2][...][...]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}// 报错：invalid use of [...] array (outside a composite literal) 以及 missing type in composite literal
ad30 := new([2][2][2]int)
mfp.PrintFmtVal("a10", a10, verbs)
mfp.PrintFmtVal("a11", a11, verbs)
mfp.PrintFmtVal("a12", a12, verbs)
//mfp.PrintFmtVal("a13", a13, verbs)
mfp.PrintFmtVal("ad3", ad3, verbs)
//mfp.PrintFmtVal("ad3x1", ad3x1, verbs)
//mfp.PrintFmtVal("ad3x2", ad3x2, verbs)
//mfp.PrintFmtVal("ad3x3", ad3x3, verbs)
mfp.PrintFmtVal("ad30", ad30, verbs)
mfp.PrintFmtVal("*ad30", *ad30, verbs)
```

```
多维数组
二维数组
a6:     %T -> [3][3]int | %v -> [[1 2 3] [4 5 6] [7 8 9]] | %#v -> [3][3]int{[3]
int{1, 2, 3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}}
a7:     %T -> [3][3]int | %v -> [[1 2 3] [4 5 6] [7 8 9]] | %#v -> [3][3]int{[3]
int{1, 2, 3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}}
a8:     %T -> [3][3]int | %v -> [[1 2 3] [4 5 6] [7 8 9]] | %#v -> [3][3]int{[3]
int{1, 2, 3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}}
ad2:    %T -> [3][3]int | %v -> [[1 2 3] [4 5 6] [7 8 9]] | %#v -> [3][3]int{[3]
int{1, 2, 3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}}
ad20:   %T -> *[3][3]int | %v -> &[[0 0 0] [0 0 0] [0 0 0]] | %#v -> &[3][3]int{
[3]int{0, 0, 0}, [3]int{0, 0, 0}, [3]int{0, 0, 0}}
*ad20:  %T -> [3][3]int | %v -> [[0 0 0] [0 0 0] [0 0 0]] | %#v -> [3][3]int{[3]
int{0, 0, 0}, [3]int{0, 0, 0}, [3]int{0, 0, 0}}
三维数组
a10:    %T -> [2][2][2]int | %v -> [[[1 2] [3 4]] [[5 6] [7 8]]] | %#v -> [2][2]
[2]int{[2][2]int{[2]int{1, 2}, [2]int{3, 4}}, [2][2]int{[2]int{5, 6}, [2]int{7, 
8}}}
a11:    %T -> [2][2][2]int | %v -> [[[1 2] [3 4]] [[5 6] [7 8]]] | %#v -> [2][2]
[2]int{[2][2]int{[2]int{1, 2}, [2]int{3, 4}}, [2][2]int{[2]int{5, 6}, [2]int{7, 
8}}}
a12:    %T -> [2][2][2]int | %v -> [[[1 2] [3 4]] [[5 6] [7 8]]] | %#v -> [2][2]
[2]int{[2][2]int{[2]int{1, 2}, [2]int{3, 4}}, [2][2]int{[2]int{5, 6}, [2]int{7, 
8}}}
ad3:    %T -> [2][2][2]int | %v -> [[[1 2] [3 4]] [[5 6] [7 8]]] | %#v -> [2][2]
[2]int{[2][2]int{[2]int{1, 2}, [2]int{3, 4}}, [2][2]int{[2]int{5, 6}, [2]int{7, 
8}}}
ad30:   %T -> *[2][2][2]int | %v -> &[[[0 0] [0 0]] [[0 0] [0 0]]] | %#v -> &[2]
[2][2]int{[2][2]int{[2]int{0, 0}, [2]int{0, 0}}, [2][2]int{[2]int{0, 0}, [2]int{
0, 0}}}
*ad30:  %T -> [2][2][2]int | %v -> [[[0 0] [0 0]] [[0 0] [0 0]]] | %#v -> [2][2]
[2]int{[2][2]int{[2]int{0, 0}, [2]int{0, 0}}, [2][2]int{[2]int{0, 0}, [2]int{0, 
0}}}
```

## U修改

### 修改元素	

​	根据以下代码和运行结果，我们可以发现在使用`a[index]`修改元素的时候，若编译时发现`index`已经超过`a`数组的长度，则编译时报错，若在运行时发现`index`已经超过`a`数组的长度，则程序引发panic。

```go
fmt.Println("一维数组")

a14 := [...]int{1, 2, 3}
mfp.PrintFmtVal("a14", a14, verbs)

a14[0] = 11
mfp.PrintFmtVal("a14", a14, verbs)

a14[len(a14)-1] = 33
mfp.PrintFmtVal("a14", a14, verbs)

//a14[len(a14)] = 44   // 报错：invalid argument: index 3 out of bounds [0:3]
//a14[len(a14)+1] = 55 // 报错：invalid argument: index 3 out of bounds [0:3]

pa141 := &a14[0]
*pa141 = 111
mfp.PrintFmtVal("a14", a14, verbs)
fmt.Println("二维数组")

a15 := [...][2]int{{1, 2}, {3, 4}}
a15[0][0] = 11
mfp.PrintFmtVal("a15", a15, verbs)

a15[len(a15)-1][0] = 33
mfp.PrintFmtVal("a15", a15, verbs)
//a15[len(a15)][0] = 11   // 报错：invalid argument: index 2 out of bounds [0:2]
//a15[len(a15)+1][0] = 11 // 报错：invalid argument: index 3 out of bounds [0:2]

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
a15:    %T -> [2][2]int | %v -> [[11 2] [3 4]] | %#v -> [2][2]int{[2]int{11, 2},
 [2]int{3, 4}}
a15:    %T -> [2][2]int | %v -> [[11 2] [33 4]] | %#v -> [2][2]int{[2]int{11, 2}
, [2]int{33, 4}}
a15:    %T -> [2][2]int | %v -> [[111 2] [33 4]] | %#v -> [2][2]int{[2]int{111, 
2}, [2]int{33, 4}}
三维数组和二维数组类似
```

```go
package main

import "fmt"

var a = [3]int{1, 2, 3}

func editA(i, v int) {
	a[i] = v
}

func main() {
	fmt.Println(a)
	editA(len(a), 4)
	fmt.Println(a)
}
```

```
panic: runtime error: index out of range [3] with length 3
```



### 用整个数组赋值

​	根据以下代码和运行结果，我们可以发现在使用整个数组赋值的时候，若新数组的长度与原数组的长度不一致，或者新数组的元素类型与原数组的元素类型不一致，则编译时报错。

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

## A访问

​	访问数组中的某一元素，可通过索引下标，索引下标范围`[0, len(数组名) - 1]`，即从0开始到数组的长度减去1。

### 直接访问指定索引下标的元素

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

### 遍历数组

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

### 获取相关数组属性

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

## D删除

### 是否可以删除某一元素呢？

​	=> 不可以

​	通过上面的创建、修改、访问，我们知道数组有两个重要的属性：长度和元素类型。假设真能删除某一元素，那么新旧数组的长度就不一样了，这样就导致了前后两个数组不一致，故Go语言的设计中也没有提供删除数组元素的操作。

### 作为实参传递给函数或方法

​	因数组在Go语言中是`值类型`，数组作为实参传递给函数，将发生完整复制数组，若数组很大，对于内存和性能将会是一个大开销。

## 其他问题

### 两个数组能进行比较吗？

​	根据以下代码和运行结果，我们可以发现只有在两个数组的类型一致（数组元素类型和数组长度都一致）的情况下，可以使用`==`比较两个数组，但不可以使用`>`、`<`、`>=`、`<=`等比较符号进行比较，且在两个数组的元素值完全相同的情况下，`==`的结果才是`true`。

```go
a1 := [3]int{1, 2, 3}
//a2 := [4]int{1, 2, 3, 4}
a3 := [3]int{1, 2, 3}
//a4 := [3]string{"1", "2", "3"}
a5 := [3]int{2, 3, 4}
// 报错：invalid operation: a1 == a2 (mismatched types [3]int and [4]int)
//if a1 == a2 {
//	fmt.Println("a1 == a2")
//}

if a1 == a3 {
    fmt.Println("a1 == a3")
} else {
    fmt.Println("a1 != a3")
}

// 报错：invalid operation: a1 < a3 (operator < not defined on array)
//if a1 < a3 {
//	fmt.Println("a1 < a3")
//}

// 报错：invalid operation: a1 > a3 (operator > not defined on array)
//if a1 > a3 {
//	fmt.Println("a1 > a3")
//}

// 报错：invalid operation: a1 <= a3 (operator <= not defined on array)
//if a1 <= a3 {
//	fmt.Println("a1 <= a3")
//}

// 报错：invalid operation: a1 >= a3 (operator >= not defined on array)
//if a1 >= a3 {
//	fmt.Println("a1 >= a3")
//}

// 报错：invalid operation: a1 == a4 (mismatched types [3]int and [3]string)
//if a1 == a4 {
//	fmt.Println("a1 == a4")
//}

if a1 == a5 {
    fmt.Println("a1 == a5")
} else {
    fmt.Println("a1 != a5")
}
```

```
a1 == a3
a1 != a5
```

### 数组的元素类型可以是哪些?

​	根据以下代码和运行结果，我们可以发现数组的元素类型可以任意类型，包括数组（即变成多维数组）。

```go
package main

import "github.com/before80/utils/mfp"

type St struct {
	a int
	b string
}

type Itf interface {
	M1()
	M2()
}

var verbs = []string{"T", "v", "#v"}

func main() {
	var aby [3]byte
	var abl [3]bool
	var as [3]string
	var ar [3]rune
	var autr [3]uintptr
	var ai8 [3]int8
	var ai16 [3]int16
	var ai32 [3]int32
	var ai64 [3]int64
	var ai [3]int
	var aui8 [3]uint8
	var aui16 [3]uint16
	var aui64 [3]uint64
	var aui32 [3]uint32
	var aui [3]uint
	var af32 [3]float32
	var af64 [3]float64
	var acplx64 [3]complex64
	var acplx128 [3]complex128
	var asli [3][]int
	var ast [3]St
	var aitf [3]Itf
	var af [3]func(int) int
	var am [3]map[string]int
	var ach [3]chan int
	mfp.PrintFmtVal("aby", aby, verbs)
	mfp.PrintFmtVal("abl", abl, verbs)
	mfp.PrintFmtVal("as", as, verbs)
	mfp.PrintFmtVal("ar", ar, verbs)
	mfp.PrintFmtVal("autr", autr, verbs)
	mfp.PrintFmtVal("ai8", ai8, verbs)
	mfp.PrintFmtVal("ai16", ai16, verbs)
	mfp.PrintFmtVal("ai32", ai32, verbs)
	mfp.PrintFmtVal("ai64", ai64, verbs)
	mfp.PrintFmtVal("ai", ai, verbs)
	mfp.PrintFmtVal("aui8", aui8, verbs)
	mfp.PrintFmtVal("aui16", aui16, verbs)
	mfp.PrintFmtVal("aui32", aui32, verbs)
	mfp.PrintFmtVal("aui64", aui64, verbs)
	mfp.PrintFmtVal("aui", aui, verbs)
	mfp.PrintFmtVal("af32", af32, verbs)
	mfp.PrintFmtVal("af64", af64, verbs)
	mfp.PrintFmtVal("acplx64", acplx64, verbs)
	mfp.PrintFmtVal("acplx128", acplx128, verbs)
	mfp.PrintFmtVal("asli", asli, verbs)
	mfp.PrintFmtVal("ast", ast, verbs)
	mfp.PrintFmtVal("asli", asli, verbs)
	mfp.PrintFmtVal("aitf", aitf, verbs)
	mfp.PrintFmtVal("af", af, verbs)
	mfp.PrintFmtVal("am", am, verbs)
	mfp.PrintFmtVal("ach", ach, verbs)
}
```

```
aby:    %T -> [3]uint8 | %v -> [0 0 0] | %#v -> [3]uint8{0x0, 0x0, 0x0}
abl:    %T -> [3]bool | %v -> [false false false] | %#v -> [3]bool{false, false, false}
as:     %T -> [3]string | %v -> [  ] | %#v -> [3]string{"", "", ""}
ar:     %T -> [3]int32 | %v -> [0 0 0] | %#v -> [3]int32{0, 0, 0}
autr:   %T -> [3]uintptr | %v -> [0 0 0] | %#v -> [3]uintptr{0x0, 0x0, 0x0}     
ai8:    %T -> [3]int8 | %v -> [0 0 0] | %#v -> [3]int8{0, 0, 0}
ai16:   %T -> [3]int16 | %v -> [0 0 0] | %#v -> [3]int16{0, 0, 0}
ai32:   %T -> [3]int32 | %v -> [0 0 0] | %#v -> [3]int32{0, 0, 0}
ai64:   %T -> [3]int64 | %v -> [0 0 0] | %#v -> [3]int64{0, 0, 0}
ai:     %T -> [3]int | %v -> [0 0 0] | %#v -> [3]int{0, 0, 0}
aui8:   %T -> [3]uint8 | %v -> [0 0 0] | %#v -> [3]uint8{0x0, 0x0, 0x0}
aui16:  %T -> [3]uint16 | %v -> [0 0 0] | %#v -> [3]uint16{0x0, 0x0, 0x0}       
aui32:  %T -> [3]uint32 | %v -> [0 0 0] | %#v -> [3]uint32{0x0, 0x0, 0x0}       
aui64:  %T -> [3]uint64 | %v -> [0 0 0] | %#v -> [3]uint64{0x0, 0x0, 0x0}       
aui:    %T -> [3]uint | %v -> [0 0 0] | %#v -> [3]uint{0x0, 0x0, 0x0}
af32:   %T -> [3]float32 | %v -> [0 0 0] | %#v -> [3]float32{0, 0, 0}
af64:   %T -> [3]float64 | %v -> [0 0 0] | %#v -> [3]float64{0, 0, 0}
acplx64:        %T -> [3]complex64 | %v -> [(0+0i) (0+0i) (0+0i)] | %#v -> [3]co
mplex64{(0+0i), (0+0i), (0+0i)}
acplx128:       %T -> [3]complex128 | %v -> [(0+0i) (0+0i) (0+0i)] | %#v -> [3]c
omplex128{(0+0i), (0+0i), (0+0i)}
asli:   %T -> [3][]int | %v -> [[] [] []] | %#v -> [3][]int{[]int(nil), []int(nil), []int(nil)}
ast:    %T -> [3]main.St | %v -> [{0 } {0 } {0 }] | %#v -> [3]main.St{main.St{a:0, b:""}, main.St{a:0, b:""}, main.St{a:0, b:""}}
asli:   %T -> [3][]int | %v -> [[] [] []] | %#v -> [3][]int{[]int(nil), []int(nil), []int(nil)}
aitf:   %T -> [3]main.Itf | %v -> [<nil> <nil> <nil>] | %#v -> [3]main.Itf{main.Itf(nil), main.Itf(nil), main.Itf(nil)}
af:     %T -> [3]func(int) int | %v -> [<nil> <nil> <nil>] | %#v -> [3]func(int)
 int{(func(int) int)(nil), (func(int) int)(nil), (func(int) int)(nil)}
am:     %T -> [3]map[string]int | %v -> [map[] map[] map[]] | %#v -> [3]map[stri
ng]int{map[string]int(nil), map[string]int(nil), map[string]int(nil)}
ach:    %T -> [3]chan int | %v -> [<nil> <nil> <nil>] | %#v -> [3]chan int{(chan int)(nil), (chan int)(nil), (chan int)(nil)}

```



## 易混淆的知识点

### 数组指针和指针数组

​	根据以下代码和运行结果，我们可以发现对数组指针进行赋值时，若`&a`所在的数组与`ptr`数组指针在定义时所指向的数组类型不同，则编译时报错。

​	数组指针和指针数组，怎么区分呢，直接看最后两字是什么！

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



## 易错点

### 访问最后一个数组元素

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

### 以为`[...]T`可以用来作为形参

```go
package main

func editA1(a [...]int) {
	_ = a
}

func main() {
}
```

​	然而，编译时直接报错：

```go
syntax error: unexpected ..., expected expression
```

### 以为数组的零值是nil

```go

```



## 数组的特点

​	数组中的元素在内存中的存储是连续的，故检索数组非常快，但定义后数组的大小不能再修改。
