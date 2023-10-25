+++
title = "unsafe"
linkTitle = "unsafe"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/unsafe@go1.20.1

​	unsafe包包含了绕过Go程序类型安全的操作。

​	导入unsafe包的程序可能是非可移植的，并且不受Go 1兼容性指南的保护。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func Alignof 

``` go 
func Alignof(x ArbitraryType) uintptr
```

​	Alignof函数取任何类型的表达式x，并返回假定变量v通过var v = x声明后的所需对齐方式。它是最大的值m，使得v的地址始终为0 mod m。它与reflect.TypeOf(x).Align()返回的值相同。作为特例，如果变量s是结构体类型，f是该结构体中的字段，则Alignof(s.f)将返回该类型字段在结构体中所需的对齐方式。该情况与reflect.TypeOf(s.f).FieldAlign()返回的值相同。如果参数的类型不具有可变大小，则Alignof函数的返回值是一个Go常量。(有关可变大小类型的定义，请参见Sizeof函数的描述。)

#### func Offsetof 

``` go 
func Offsetof(x ArbitraryType) uintptr
```

​	Offsetof函数返回由x表示的字段在结构体内的偏移量，x必须是structValue.field的形式。换句话说，它返回结构体开始和字段开始之间的字节数。如果参数x的类型不具有可变大小，则Offsetof函数的返回值是一个Go常量。(有关可变大小类型的定义，请参见[Sizeof](#func-sizeof)函数的描述。)

#### func Sizeof 

``` go 
func Sizeof(x ArbitraryType) uintptr
```

​	Sizeof函数取任何类型的表达式x，并返回假定变量v通过var v = x声明后的大小(以字节为单位)。该大小不包括x可能引用的任何内存。例如，如果x是一个切片，则Sizeof函数返回切片描述符的大小，而不是切片引用的内存的大小。对于结构体，大小包括由字段对齐引入的任何填充。如果参数x的类型不具有可变大小，则Sizeof函数的返回值是一个Go常量。(类型具有可变大小，如果它是类型参数，或者它是具有可变大小元素的数组或结构体类型)。

#### func String  <- go1.20

``` go 
func String(ptr *byte, len IntegerType) string
```

​	String函数返回一个字符串值，其底层字节从ptr开始，长度为len。

​	len参数必须是整数类型或未命名常量。常量len参数必须是非负数，并且可以表示为int类型的值；如果它是未命名常量，则给定类型为int。在运行时，如果len为负数，或者ptr为nil且len不为零，则会发生运行时恐慌。

​	由于Go字符串是不可变的，因此在调用String函数后不得修改传递的字节。

#### func StringData  <- go1.20

``` go 
func StringData(str string) *byte
```

​	StringData函数返回一个指向str底层字节的指针。对于空字符串，返回值是未指定的，可能为nil。

​	由于Go字符串是不可变的，因此不能修改StringData函数返回的字节。

## 类型

Arbitrary `adj.任意的，随心所欲的；专横的，武断的`

- 英*/*ˈɑːbɪtrəri*/*
- 美*/*ˈɑːrbɪtreri*/*

### type ArbitraryType 

``` go 
type ArbitraryType int
```

​	ArbitraryType仅用于文档目的，实际上不属于unsafe包。它表示任意Go表达式的类型。

#### func Slice  <- go1.17

``` go 
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
```

​	Slice函数返回一个切片，其底层数组从ptr开始，长度和容量为len。Slice(ptr，len)等效于

```
(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
```

​	特殊情况下，如果ptr为nil且len为零，则Slice返回nil。

​	len参数必须是整数类型或无类型常量。常量len参数必须为非负数，并且可以由int类型的值表示；如果它是无类型常量，则给定类型为int。运行时，如果len为负数，或者ptr为nil且len不为零，则会发生运行时恐慌。

#### func SliceData  <- go1.20

``` go 
func SliceData(slice []ArbitraryType) *ArbitraryType
```

​	SliceData返回参数切片的底层数组的指针。

- 如果cap(slice)> 0，则SliceData返回`&slice[:1][0]`。
- 如果slice == nil，SliceData返回nil。 如果slice == nil，则SliceData返回nil。
- 否则，SliceData返回一个非nil指向未指定内存地址的指针。

### type IntegerType  <- go1.17

``` go 
type IntegerType int
```

​	IntegerType仅用于文档目的，实际上不属于unsafe包。它表示任意整数类型。

### type Pointer 

``` go 
type Pointer *ArbitraryType
```

​	Pointer 表示指向任意类型的指针。Pointer类型具有其他类型没有的四个特殊操作：

- 任何类型的指针值都可以转换为 Pointer。 
- Pointer 可以转换为任何类型的指针值。 
- uintptr 可以转换为 Pointer。 
- Pointer 可以转换为 uintptr。 

​	因此，Pointer 允许程序打破类型系统，并读写任意内存。需要非常小心地使用。

​	以下是 Pointer 的有效使用方式。不使用这些方式的代码可能是无效的，今天或未来可能会变得无效。即使是下面的有效模式也有重要的警告。

​	运行 "go vet" 可以帮助找到未符合这些模式的 Pointer 使用情况，但 "go vet" 的沉默并不能保证代码是有效的。

(1) Conversion of a *T1 to Pointer to *T2.

(1) 将 `*T1` 转换为 Pointer 到 `*T2`。

​	假设 T2 不比 T1 大，并且两者具有相同的内存布局，此转换允许将一种类型的数据重新解释为另一种类型的数据。一个例子是 math.Float64bits 的实现：

``` go 
func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
```

(2) 将 Pointer 转换为 uintptr(但不转换回 Pointer)。

​	将 Pointer 转换为 uintptr 可以产生所指向值的内存地址，作为整数。这样的 uintptr 通常用于打印。

​	通常情况下，不应将 uintptr 转换回 Pointer。

​	uintptr 是一个整数，而不是引用。将 Pointer 转换为 uintptr 会创建一个没有指针语义的整数值。即使 uintptr 持有某个对象的地址，垃圾回收器也不会在对象移动时更新 uintptr 的值，uintptr 也无法防止该对象被回收。

​	下面的模式列举了从 uintptr 到 Pointer 的唯一有效转换。

(3) 将 Pointer 转换为 uintptr，然后进行算术操作，再转换回 Pointer。

​	如果 p 指向分配的对象，则可以通过将其转换为 uintptr、添加偏移量，然后再将其转换回 Pointer 来使 p 在对象中前进。

```go 
p = unsafe.Pointer(uintptr(p) + offset)
```

​	此模式最常见的用途是访问结构体的字段或数组的元素：

```go 
// 等同于 f := unsafe.Pointer(&s.f)
f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f))

// 等同于 e := unsafe.Pointer(&x[i])
e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0]))
```

​	在此方式下，可以对指针进行偏移量的加减操作。通常用 `&^` 对指针进行取整，通常是为了对齐。在所有情况下，结果必须继续指向原始分配的对象。

Unlike in C, it is not valid to advance a pointer just beyond the end of its original allocation:

与C语言不同的是，将一个指针推进到其原始分配的末端是无效的：

​	与 C 语言不同的是，将指针移动到其原始分配空间的边界之外是无效的：

```go 
// 不正确：end 指向分配空间之外。
var s thing
end = unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s))

// 不正确：end 在分配的空间之外。
b := make([]byte, n)
end = unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(n))
```

​	请注意，两个转换必须出现在同一表达式中，只能在它们之间进行算术运算：

```go 
// 不正确：uintptr在转换回Pointer之前不能被存储在变量中。
u := uintptr(p)
p = unsafe.Pointer(u + offset)
```

​	请注意，这个指针必须指向已分配的对象，因此它可能不是 nil。

```go 
// INVALID: conversion of nil pointer
// 不合法：转换为零的指针
u := unsafe.Pointer(nil)
p := unsafe.Pointer(uintptr(u) + offset)
```

(4) 当调用 syscall.Syscall 时，将指针转换为 uintptr。

​	在 syscall 包中，Syscall 函数将它们的 uintptr 参数直接传递给操作系统，然后根据调用的细节重新解释其中的一些参数为指针。也就是说，系统调用的实现隐式地将某些参数从 uintptr 转换回指针。

​	如果必须将指针参数转换为 uintptr 以便用作参数，那么该转换必须出现在调用表达式本身中：

```go 
syscall.Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))
```

​	编译器通过在汇编实现的函数调用的参数列表中处理将 Pointer 转换为 uintptr，安排保留引用的分配对象(如果有的话)，即使从类型上看，在调用期间似乎不再需要该对象。

​	为了让编译器识别这种模式，转换必须出现在参数列表中：

```go 
// INVALID: uintptr在系统调用过程中隐式转换回指针前不能被存储在变量中。
u := uintptr(unsafe.Pointer(p))
syscall.Syscall(SYS_READ, uintptr(fd), u, uintptr(n))
```

(5) 将 reflect.Value.Pointer 或 reflect.Value.UnsafeAddr 的结果从 uintptr 转换为 Pointer。

​	reflect包名为 Pointer 和 UnsafeAddr 的 Value 方法返回类型为 uintptr，而不是 unsafe.Pointer，以防止调用者在未导入 "unsafe" 的情况下将结果更改为任意类型。但是，这意味着结果是脆弱的，并且必须在调用后立即将其转换为 Pointer，即在同一表达式中：

```go
p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))
```

​	与以上情况类似，存储转换之前的结果是无效的：

```go
// INVALID: uintptr在转换回Pointer之前不能被存储在变量中。
u := reflect.ValueOf(new(int)).Pointer()
p := (*int)(unsafe.Pointer(u))
```

(6) 将reflect.SliceHeader或reflect.StringHeader数据字段转换为Pointer，或将Pointer转换为reflect.SliceHeader或reflect.StringHeader数据字段。

​	与之前的情况类似，reflect.SliceHeader和reflect.StringHeader将Data字段声明为uintptr，以防止调用者在未导入"unsafe"包之前将结果更改为任意类型。但是，这意味着SliceHeader和StringHeader仅在解释实际切片或字符串值的内容时才有效。

``` go 
var s string
hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // case 1
hdr.Data = uintptr(unsafe.Pointer(p))              // case 6 (this case)
hdr.Len = n
```

​	在这种用法中，hdr.Data实际上是一种替代方式，用于引用字符串头中的底层指针，而不是uintptr变量本身。

​	通常情况下，reflect.SliceHeader和reflect.StringHeader应仅作为`*reflect.SliceHeader`和`*reflect.StringHeader`使用，指向实际的切片或字符串，而不是作为普通结构体使用。程序不应声明或分配这些结构体类型的变量。

```go 
// INVALID：直接声明的标头不会将Data作为引用保存。
var hdr reflect.StringHeader
hdr.Data = uintptr(unsafe.Pointer(p))
hdr.Len = n
s := *(*string)(unsafe.Pointer(&hdr)) // p可能已经丢失
```

#### func Add  <- go1.17

``` go 
func Add(ptr Pointer, len IntegerType) Pointer
```

​	Add函数将len添加到ptr并返回更新后的指针Pointer(uintptr(ptr) + uintptr(len))。len参数必须是整数类型或无类型常量。常量len参数必须可由int类型的值表示；如果它是无类型常量，则它会被赋予int类型。仍然适用指针的有效用法规则。