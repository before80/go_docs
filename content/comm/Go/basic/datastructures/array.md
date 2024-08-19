+++
title = "数组"
date = 2024-08-19T19:26:32+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 说明
>
> ​	以下实例代码中：
>
> ​	`verbs`的定义是：`var verbs = []string{"T", "v", "#v"} `
>
> ​	`mfp`来自：`"github.com/before80/utils/mfp"`

## C创建

### 一维数组

#### 1 直接创建

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

#### 2 是否可以通过make创建？

​	=> 不可以

```go
a4 := make([3]int{1, 2, 3}) // 报错：[3]int{…} is not a type
```

#### 3 是否可以通过new创建？

​	=> 可以

```go
a5 := new([3]int) 
```

### 多维数组

​	从一维数组的创建中可以推测出，多维数组的创建也是不能通过make进行创建，new可以进行创建。

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



## U修改

### 1 修改元素

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

### 2 用整个数组赋值

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

### 1 直接访问指定索引下标的元素

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

### 2 遍历数组

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

### 3 获取相关数组属性

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

### 1 是否可以删除某一元素呢？

​	=> 不可以

​	通过上面的创建、修改、访问，我们知道数组有两个重要的属性：长度和元素类型。假设真能删除某一元素，那么新旧数组的长度就不一样了，这样就导致了前后两个数组不一致，故Go语言的设计中也没有提供删除数组元素的操作。

## 作为实参传递给函数或方法

​	因数组在Go语言中是`值类型`，数组作为实参传递给函数，将发生完整复制数组，若数组很大，对于内存和性能将会是一个大开销。

## 易混淆的知识点

### 1 数组指针和指针数组

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

### 1 访问最后一个数组元素

​	直接用`a[len(a)]`访问数组a的最后一个元素 => 肯定报错

```go
fmt.Println("访问数组的最后一个元素")
a21 := [...]int{1, 2, 3}
//fmt.Println(a21[len(a21)]) // 报错：invalid argument: index 3 out of bounds [0:3]
fmt.Println(a21[len(a21)-1]) // 正确方式
```

```
3
```

## 数组的特点

​	数组中的元素在内存中的存储是连续的，故检索数组非常快，但定义后数组的大小不能再修改。	
