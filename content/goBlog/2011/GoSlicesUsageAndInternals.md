+++
title = "Go 切片：用法和内部机制"
weight = 29
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go Slices: usage and internals - Go 切片：用法和内部机制

> 原文：[https://go.dev/blog/slices-intro](https://go.dev/blog/slices-intro)

Andrew Gerrand
5 January 2011

2011年1月5日

## 简介

​	Go的切片类型提供了一种方便和高效的方式来处理有类型数据的序列。切片类似于其他语言中的数组，但具有一些不寻常的属性。本文将介绍切片是什么以及如何使用它们。

## 数组

​	切片类型是建立在Go的数组类型之上的一种抽象，因此要理解切片，我们必须先了解数组。

​	数组类型定义指定了长度和元素类型。例如，类型[4]int表示四个整数的数组。数组的大小是固定的；其长度是其类型的一部分（[4]int和[5]int是不同的不兼容类型）。可以按照通常的方式对数组进行索引，因此表达式s[n]访问从零开始的第n个元素。

```go
var a [4]int
a[0] = 1
i := a[0]
// i == 1
```

​	数组不需要显式初始化；数组的零值是一个已准备好使用的数组，其元素本身被清零：

```go
// a[2] == 0, int类型的零值
```

​	[4]int在内存中的表示只是顺序排列的四个整数值：

![img](GoSlicesUsageAndInternals_img/slice-array.png)

​	Go的数组是值。数组变量表示整个数组；它不是指向第一个数组元素的指针（如在C中的情况）。这意味着当您分配或传递数组值时，将复制其内容。（为避免复制，可以传递指向数组的指针，但这样就是指向数组的指针，而不是数组。）关于数组的一种思考方式是一种类似于结构体但具有索引而不是命名字段的固定大小的复合值。

​	可以像这样指定数组字面量：

```go
b := [2]string{"Penn", "Teller"}
```

或者，您可以让编译器为您计算数组元素：

```go
b := [...]string{"Penn", "Teller"}
```

​	在两种情况下，b的类型都是[2]string。

## 切片

​	数组有它们的用处，但它们有点不灵活，所以在Go代码中很少看到它们。相比之下，切片无处不在。它们基于数组提供了极大的强大和便利。

​	切片的类型规范是[]T，其中T是切片元素的类型。与数组类型不同，切片类型没有指定长度。

​	切片字面量声明的方式与字面量声明的方式相同，只需省略元素计数：

```go
letters := []string{"a", "b", "c", "d"}
```

​	可以使用内置函数`make`创建切片，其签名为

```go
func make([]T, len, cap) []T
```

​	其中T代表要创建的切片的元素类型。make函数需要一个类型、一个长度和一个可选容量。调用make函数会分配一个数组并返回一个引用该数组的切片。

```go
var s []byte
s = make([]byte, 5, 5)
// s == []byte{0, 0, 0, 0, 0}
```

​	当省略容量参数时，它默认为指定的长度。下面是同样代码的更简洁版本：

```go
s := make([]byte, 5)
```

​	可以使用内置的`len`和`cap`函数检查切片的长度和容量。

```go
len(s) == 5
cap(s) == 5
```

​	接下来的两个部分讨论了长度和容量之间的关系。

​	切片的零值是nil。对于nil切片，len和cap函数都将返回0。

​	切片也可以通过"切片"现有的切片或数组来形成。切片是通过用两个由冒号分隔的索引指定的半开范围来完成的。例如，表达式b[1:4]创建了一个包括b的元素1到3的切片（结果切片的索引将是0到2）。

```go
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
// b[1:4] == []byte{'o', 'l', 'a'}, 共享b的存储空间
```

​	切片表达式的起始和结束索引是可选的；它们分别默认为零和切片的长度：

```go
// b[:2] == []byte{'g', 'o'}
// b[2:] == []byte{'l', 'a', 'n', 'g'}
// b[:] == b
```

这也是使用数组创建切片的语法：

```go
x := [3]string{"Лайка", "Белка", "Стрелка"}
s := x[:] // 引用x的存储空间的切片
```

## 切片内部

​	切片是数组段的描述符。它包含一个指向数组的指针、段的长度和其容量（段的最大长度）。

![img](GoSlicesUsageAndInternals_img/slice-struct.png)

我们之前使用 make([]byte, 5) 创建的变量 s 结构如下：

![img](GoSlicesUsageAndInternals_img/slice-1.png)

​	长度是片段中由该片段引用的元素数。容量是底层数组中的元素数（从由该片段指针引用的元素开始）。在接下来的几个示例中，我们将清楚地说明长度和容量之间的区别。

​	当我们对 s 进行切片时，请观察切片数据结构的变化及其与底层数组的关系：

```go
s = s[2:4]
```

![img](GoSlicesUsageAndInternals_img/slice-2.png)

​	切片不会复制切片的数据。它创建一个指向原始数组的新切片值。这使得切片操作像操作数组索引一样高效。因此，修改重新切片的元素（而不是切片本身）将修改原始切片的元素：

```go
d := []byte{'r', 'o', 'a', 'd'}
e := d[2:]
// e == []byte{'a', 'd'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}
```

​	我们之前对 s 进行了一个比容量更短的长度的切片。我们可以通过再次切片将 s 增长到其容量：

```go
s = s[:cap(s)]
```

![img](GoSlicesUsageAndInternals_img/slice-3.png)

​	切片的长度不能超过其容量。尝试这样做将导致运行时恐慌，就像在切片或数组范围之外索引一样。同样，切片不能重新切片到零以下以访问数组中的早期元素。

## 增长切片（copy 和 append 函数）

​	要增加切片的容量，必须创建一个新的更大的切片，并将原始切片的内容复制到其中。这种技术是其他语言中的动态数组实现在幕后的工作方式。下一个示例通过创建一个新切片 t，将 s 的内容复制到 t 中，然后将切片值 t 赋给 s，将 s 的容量加倍：

```go
t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
for i := range s {
        t[i] = s[i]
}
s = t
```

​	通过内置的 copy 函数，可以简化此常见操作的循环部分。正如名称所示，copy 将数据从源切片复制到目标切片。它返回复制的元素数。

```go
func copy(dst, src []T) int
```

​	copy 函数支持在不同长度的切片之间进行复制（仅复制较小数量的元素）。此外，copy 可以处理共享相同底层数组的源和目标切片，正确处理重叠切片。

​	使用 copy，我们可以简化上面的代码段：

```go
t := make([]byte, len(s), (cap(s)+1)*2)
copy(t, s)
s = t
```

​	一个常见的操作是将数据追加到一个切片的末尾。这个函数会将字节元素追加到一个字节切片中，如果需要的话增长切片，并返回更新后的切片值：

```go
func AppendByte(slice []byte, data ...byte) []byte {
    m := len(slice)
    n := m + len(data)
    if n > cap(slice) { // 如果需要，重新分配
        // 分配比所需的两倍还多，以备将来增长。
        newSlice := make([]byte, (n+1)*2)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:n]
    copy(slice[m:n], data)
    return slice
}
```

可以像这样使用 AppendByte：

```go
p := []byte{2, 3, 5}
p = AppendByte(p, 7, 11, 13)
// p == []byte{2, 3, 5, 7, 11, 13}
```

像 AppendByte 这样的函数非常有用，因为它们可以完全控制切片增长的方式。根据程序的特性，可能需要以更小或更大的块分配内存，或者限制重新分配的大小。

​	但是大多数程序不需要完全控制，因此 Go 提供了一个内置的 append 函数，适用于大多数情况；它的签名如下：

```go
func append(s []T, x ...T) []T
```

​	append 函数将元素 x 追加到切片 s 的末尾，并在需要更大的容量时增长切片。

```go
a := make([]int, 1)
// a == []int{0}
a = append(a, 1, 2, 3)
// a == []int{0, 1, 2, 3}
```

​	要将一个切片附加到另一个切片，请使用 ... 将第二个参数扩展为参数列表。

```go
a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...) // 相当于 "append(a, b[0], b[1], b[2])"
// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
```

​	由于切片的零值（nil）就像零长度的切片，因此您可以声明一个切片变量，然后在循环中将其附加：

```go
// Filter 返回一个仅包含满足 fn() 的 s 元素的新切片。
func Filter(s []int, fn func(int) bool) []int {
    var p []int // == nil
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}
```

## 一个可能会出现的"坑点"

​	如前所述，重新切片一个切片并不会复制底层数组。完整的底层数组将一直保留在内存中，直到不再被引用为止。偶尔这会导致程序在只需要一小部分数据时，占用整个数据在内存中的空间。

​	例如，FindDigits函数将一个文件加载到内存中，并搜索其中连续的数字，将它们作为一个新的切片返回。

```go
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return digitRegexp.Find(b)
}
```

​	这段代码的行为与预期相符，但返回的[]byte指向包含整个文件的数组。由于切片引用了原始数组，只要切片被保留下来，垃圾回收器就无法释放整个数组，文件中少量有用的字节会占用整个内容在内存中的空间。

​	为了解决这个问题，在返回之前，我们可以将感兴趣的数据复制到一个新的切片中：

```go
func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}
```

​	还可以通过使用append函数来构造这个函数的更简洁版本。这留作读者的练习。

## 进一步阅读

​	[Effective Go](../../../UsingAndUnderstandingGo/EffectiveGo)对[切片](../../../UsingAndUnderstandingGo/EffectiveGo#slices)和[数组](../../../UsingAndUnderstandingGo/EffectiveGo#arrays)进行了深入的介绍，Go语言规范定义了[切片]({{< ref "/langSpec/Types#slice-types">}})及其[相关的辅助函数]({{< ref "/langSpec/Built-inFunctions#making-slices-maps-and-channels">}})。

