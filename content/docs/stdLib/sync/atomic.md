+++
title = "atomic"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# atomic

[https://pkg.go.dev/sync/atomic@go1.20.1](https://pkg.go.dev/sync/atomic@go1.20.1)

​	atomic包提供了低级别的原子内存原语，用于实现同步算法。

​	这些函数需要极其小心地使用才能正确地使用。除非是特殊的低级别应用，否则使用通道或sync包的工具进行同步更好。通过通信共享内存，而不是通过共享内存进行通信。

​	SwapT函数实现的交换操作是以下操作的原子等效操作：

```
old = *addr
*addr = new
return old
```

​	CompareAndSwapT函数实现的比较和交换操作是以下操作的原子等效操作：

```
if *addr == old {
	*addr = new
	return true
}
return false
```

​	AddT函数实现的添加操作是以下操作的原子等效操作：

```
*addr += delta
return *addr
```

​	LoadT和StoreT函数实现的加载和存储操作是"return *addr"和"*addr = val"的原子等效操作。

​	在Go内存模型的术语中，如果原子操作A的效果被原子操作B观察到，则A"在"B之前同步。此外，程序中执行的所有原子操作都像按某些顺序连续执行一样运行。这个定义提供了与C++的顺序一致原子和Java的易失性变量相同的语义。


## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [AddInt32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=118) 

``` go linenums="1"
func AddInt32(addr *int32, delta int32) (new int32)
```

AddInt32函数原子地将delta添加到`*addr`，并返回新值。考虑使用更符合人体工程学和更不容易出错的Int32.Add代替。

#### func [AddInt64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=129) 

``` go linenums="1"
func AddInt64(addr *int64, delta int64) (new int64)
```

AddInt64函数原子地将delta添加到`*addr`，并返回新值。如果您的目标是32位平台，请考虑使用更符合人体工程学和更不容易出错的Int64.Add代替(请参见错误部分)。

#### func [AddUint32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=124) 

``` go linenums="1"
func AddUint32(addr *uint32, delta uint32) (new uint32)
```

​	AddUint32函数原子地将delta添加到`*addr`，并返回新值。要从x减去已知正常数值c，请执行AddUint32(&x，^uint32(c-1))。特别地，要将x减少1，请执行AddUint32(&x，^uint32(0))。考虑使用更符合人体工程学和更不容易出错的Uint32.Add代替。

#### func [AddUint64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=136) 

``` go linenums="1"
func AddUint64(addr *uint64, delta uint64) (new uint64)
```

​	AddUint64函数原子地将delta添加到`*addr`，并返回新值。要从x减去已知正常数值c，请执行AddUint64(&x，^uint64(c-1))。特别地，要将x减少1，请执行AddUint64(&x，^uint64(0))。如果您的目标是32位平台，请考虑使用更符合人体工程学和更不容易出错的Uint64.Add代替(请参见错误部分)。

#### func [AddUintptr](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=140) 

``` go linenums="1"
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
```

​	AddUintptr函数原子地将delta添加到`*addr`，并返回新值。考虑使用更符合人体工程学和更不容易出错的Uintptr.Add代替。

#### func [CompareAndSwapInt32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=92) 

``` go linenums="1"
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
```

​	CompareAndSwapInt32函数执行 int32 值的比较并交换操作。建议使用更符合人体工程学和更少容易出错的 Int32.CompareAndSwap。

#### func [CompareAndSwapInt64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=97) 

``` go linenums="1"
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
```

​	CompareAndSwapInt64函数执行 int64 值的比较并交换操作。建议使用更符合人体工程学和更少容易出错的 Int64.CompareAndSwap (特别是针对 32 位平台；参见错误部分)。

#### func [CompareAndSwapPointer](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=114) 

``` go linenums="1"
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
```

​	CompareAndSwapPointer函数执行 unsafe.Pointer 值的比较并交换操作。建议使用更符合人体工程学和更少容易出错的 Pointer.CompareAndSwap。

#### func [CompareAndSwapUint32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=101) 

``` go linenums="1"
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
```

​	CompareAndSwapUint32函数执行 uint32 值的比较并交换操作。建议使用更符合人体工程学和更少容易出错的 Uint32.CompareAndSwap。

#### func [CompareAndSwapUint64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=106) 

``` go linenums="1"
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
```

​	CompareAndSwapUint64函数执行 uint64 值的比较并交换操作。建议使用更符合人体工程学和更少容易出错的 Uint64.CompareAndSwap (特别是针对 32 位平台；参见错误部分)。

#### func [CompareAndSwapUintptr](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=110) 

``` go linenums="1"
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
```

​	CompareAndSwapUintptr函数执行 uintptr 值的比较并交换操作。建议使用更符合人体工程学和更少容易出错的 Uintptr.CompareAndSwap。

#### func [LoadInt32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=144) 

``` go linenums="1"
func LoadInt32(addr *int32) (val int32)
```

​	LoadInt32函数原子地加载 `*addr`。考虑使用更符合人体工程学且不易出错的 Int32.Load。

#### func [LoadInt64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=149) 

``` go linenums="1"
func LoadInt64(addr *int64) (val int64)
```

​	LoadInt64函数原子地加载 `*addr`。考虑使用更符合人体工程学且不易出错的 Int64.Load(特别是如果您的目标是32位平台；请参见错误部分)。

#### func [LoadPointer](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=166) 

``` go linenums="1"
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
```

​	LoadPointer函数原子地加载 `*addr`。考虑使用更符合人体工程学且不易出错的 Pointer.Load。

#### func [LoadUint32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=153) 

``` go linenums="1"
func LoadUint32(addr *uint32) (val uint32)
```

​	LoadUint32函数原子地加载 `*addr`。考虑使用更符合人体工程学且不易出错的 Uint32.Load。

#### func [LoadUint64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=158) 

``` go linenums="1"
func LoadUint64(addr *uint64) (val uint64)
```

​	LoadUint64函数原子地加载 `*addr`。考虑使用更符合人体工程学且不易出错的 Uint64.Load(特别是如果您的目标是32位平台；请参见错误部分)。

#### func [LoadUintptr](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=162) 

``` go linenums="1"
func LoadUintptr(addr *uintptr) (val uintptr)
```

​	LoadUintptr函数原子地加载 `*addr`。考虑使用更符合人体工程学且不易出错的 Uintptr.Load。

#### func [StoreInt32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=170) 

``` go linenums="1"
func StoreInt32(addr *int32, val int32)
```

​	StoreInt32函数原子地将 val 存储到 `*addr`。考虑使用更符合人体工程学且不易出错的 Int32.Store。

#### func [StoreInt64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=175) 

``` go linenums="1"
func StoreInt64(addr *int64, val int64)
```

​	StoreInt64函数原子地将 val 存储到 `*addr`。考虑使用更符合人体工程学且不易出错的 Int64.Store(特别是如果您的目标是32位平台；请参见错误部分)。

#### func [StorePointer](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=192) 

``` go linenums="1"
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
```

​	StorePointer函数原子地将 val 存储到 `*addr`。考虑使用更符合人体工程学且不易出错的 Pointer.Store。

#### func [StoreUint32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=179) 

``` go linenums="1"
func StoreUint32(addr *uint32, val uint32)
```

​	StoreUint32函数原子性地将val存储到`*addr`。考虑使用更符合人体工程学且更少出错的Uint32.Store。

#### func [StoreUint64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=184) 

``` go linenums="1"
func StoreUint64(addr *uint64, val uint64)
```

​	StoreUint64函数原子性地将val存储到`*addr`。考虑使用更符合人体工程学且更少出错的Uint64.Store(特别是针对32位平台，参见bugs一节)。

#### func [StoreUintptr](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=188) 

``` go linenums="1"
func StoreUintptr(addr *uintptr, val uintptr)
```

​	StoreUintptr函数原子性地将val存储到`*addr`。考虑使用更符合人体工程学且更少出错的Uintptr.Store。

#### func [SwapInt32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=66)  <- go1.2

``` go linenums="1"
func SwapInt32(addr *int32, new int32) (old int32)
```

​	SwapInt32函数原子性地将new存储到`*addr`并返回先前的`*addr`值。考虑使用更符合人体工程学且更少出错的Int32.Swap。

#### func [SwapInt64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=71)  <- go1.2

``` go linenums="1"
func SwapInt64(addr *int64, new int64) (old int64)
```

​	SwapInt64函数原子性地将new存储到`*addr`并返回先前的`*addr`值。考虑使用更符合人体工程学且更少出错的Int64.Swap(特别是针对32位平台，参见bugs一节)。

#### func [SwapPointer](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=88)  <- go1.2

``` go linenums="1"
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
```

​	SwapPointer函数原子性地将new存储到`*addr`并返回先前的`*addr`值。考虑使用更符合人体工程学且更少出错的Pointer.Swap。

#### func [SwapUint32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=75)  <- go1.2

``` go linenums="1"
func SwapUint32(addr *uint32, new uint32) (old uint32)
```

​	SwapUint32函数原子性地将new存储到`*addr`并返回先前的`*addr`值。考虑使用更符合人体工程学且更少出错的Uint32.Swap。

#### func [SwapUint64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=80)  <- go1.2

``` go linenums="1"
func SwapUint64(addr *uint64, new uint64) (old uint64)
```

​	SwapUint64函数原子性地将new存储到`*addr`并返回先前的`*addr`值。考虑使用更符合人体工程学且更少出错的Uint64.Swap(特别是针对32位平台，参见bugs一节)。

#### func [SwapUintptr](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/doc.go;l=84)  <- go1.2

``` go linenums="1"
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
```

​	SwapUintptr函数原子性地将new存储到`*addr`并返回先前的`*addr`值。考虑使用更符合人体工程学且更少出错的Uintptr.Swap。

## 类型

### type [Bool](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=11)  <- go1.19

``` go linenums="1"
type Bool struct {
	// contains filtered or unexported fields
}
```

​	Bool结构体是一个原子布尔值。 零值为false。

#### (*Bool) [CompareAndSwap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=26)  <- go1.19

``` go linenums="1"
func (x *Bool) CompareAndSwap(old, new bool) (swapped bool)
```

​	CompareAndSwap方法执行布尔值x的比较和交换操作。

#### (*Bool) [Load](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=17)  <- go1.19

``` go linenums="1"
func (x *Bool) Load() bool
```

​	Load方法以原子方式加载并返回存储在x中的值。

#### (*Bool) [Store](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=20)  <- go1.19

``` go linenums="1"
func (x *Bool) Store(val bool)
```

​	Store方法原子地将val存储到x中。

#### (*Bool) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=23)  <- go1.19

``` go linenums="1"
func (x *Bool) Swap(new bool) (old bool)
```

​	Swap方法以原子方式将new存储到x中并返回先前的值。

### type [Int32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=68)  <- go1.19

``` go linenums="1"
type Int32 struct {
	// contains filtered or unexported fields
}
```

​	Int32方法是一个原子int32。 零值为零。

#### (*Int32) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=88)  <- go1.19

``` go linenums="1"
func (x *Int32) Add(delta int32) (new int32)
```

​	Add方法原子地将delta添加到x并返回新值。

#### (*Int32) [CompareAndSwap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=83)  <- go1.19

``` go linenums="1"
func (x *Int32) CompareAndSwap(old, new int32) (swapped bool)
```

​	CompareAndSwap方法执行x的比较和交换操作。

#### (*Int32) [Load](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=74)  <- go1.19

``` go linenums="1"
func (x *Int32) Load() int32
```

​	Load方法原子地加载并返回存储在x中的值。

#### (*Int32) [Store](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=77)  <- go1.19

``` go linenums="1"
func (x *Int32) Store(val int32)
```

​	Store方法原子地将val存储到x中。

#### (*Int32) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=80)  <- go1.19

``` go linenums="1"
func (x *Int32) Swap(new int32) (old int32)
```

​	Swap方法原子地将new存储到x中并返回先前的值。

### type [Int64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=91)  <- go1.19

``` go linenums="1"
type Int64 struct {
	// contains filtered or unexported fields
}
```

​	Int64方法是原子int64。零值为零。

#### (*Int64) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=112)  <- go1.19

``` go linenums="1"
func (x *Int64) Add(delta int64) (new int64)
```

​	Add方法原子地将delta添加到x并返回新值。

#### (*Int64) [CompareAndSwap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=107)  <- go1.19

``` go linenums="1"
func (x *Int64) CompareAndSwap(old, new int64) (swapped bool)
```

​	CompareAndSwap方法执行x的比较并交换操作。

#### (*Int64) [Load](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=98)  <- go1.19

``` go linenums="1"
func (x *Int64) Load() int64
```

​	Load方法原子地加载并返回存储在x中的值。

#### (*Int64) [Store](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=101)  <- go1.19

``` go linenums="1"
func (x *Int64) Store(val int64)
```

​	Store方法原子地将val存储到x中。

#### (*Int64) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=104)  <- go1.19

``` go linenums="1"
func (x *Int64) Swap(new int64) (old int64)
```

​	Swap方法原子地将new存储到x中并返回先前的值。

### type [Pointer](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=43)  <- go1.19

``` go linenums="1"
type Pointer[T any] struct {
	// contains filtered or unexported fields
}
```

​	Pointer 是一个类型为 *T 的原子指针。零值是 nil 的 *T。

#### (*Pointer[T]) [CompareAndSwap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=63)  <- go1.19

``` go linenums="1"
func (x *Pointer[T]) CompareAndSwap(old, new *T) (swapped bool)
```

​	CompareAndSwap方法执行 x 的比较并交换操作。

#### (*Pointer[T]) [Load](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=54)  <- go1.19

``` go linenums="1"
func (x *Pointer[T]) Load() *T
```

​	Load方法原子地加载并返回存储在 x 中的值。

#### (*Pointer[T]) [Store](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=57)  <- go1.19

``` go linenums="1"
func (x *Pointer[T]) Store(val *T)
```

​	Store方法原子地将 val 存储到 x 中。

#### (*Pointer[T]) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=60)  <- go1.19

``` go linenums="1"
func (x *Pointer[T]) Swap(new *T) (old *T)
```

​	Swap方法原子地将 new 存储到 x 中，并返回先前的值。

### type [Uint32](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=115)  <- go1.19

``` go linenums="1"
type Uint32 struct {
	// contains filtered or unexported fields
}
```

​	Uint32方法是一个原子 uint32。零值是零。

#### (*Uint32) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=135)  <- go1.19

``` go linenums="1"
func (x *Uint32) Add(delta uint32) (new uint32)
```

​	Add方法原子地将 delta 添加到 x 并返回新值。

#### (*Uint32) [CompareAndSwap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=130)  <- go1.19

``` go linenums="1"
func (x *Uint32) CompareAndSwap(old, new uint32) (swapped bool)
```

​	CompareAndSwap方法执行 x 的比较并交换操作。

#### (*Uint32) [Load](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=121)  <- go1.19

``` go linenums="1"
func (x *Uint32) Load() uint32
```

​	Load方法原子地加载并返回存储在x中的值。

#### (*Uint32) [Store](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=124)  <- go1.19

``` go linenums="1"
func (x *Uint32) Store(val uint32)
```

​	Store方法原子地将val存储到x中。

#### (*Uint32) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=127)  <- go1.19

``` go linenums="1"
func (x *Uint32) Swap(new uint32) (old uint32)
```

​	Swap方法原子地将new存储到x中并返回先前的值。

### type [Uint64](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=138)  <- go1.19

``` go linenums="1"
type Uint64 struct {
	// contains filtered or unexported fields
}
```

​	Uint64方法是一个原子性的uint64类型。零值为零。

#### (*Uint64) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=159)  <- go1.19

``` go linenums="1"
func (x *Uint64) Add(delta uint64) (new uint64)
```

​	Add方法原子性地将delta添加到x中并返回新值。

#### (*Uint64) [CompareAndSwap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=154)  <- go1.19

``` go linenums="1"
func (x *Uint64) CompareAndSwap(old, new uint64) (swapped bool)
```

​	CompareAndSwap方法在x上执行比较并交换操作。

#### (*Uint64) [Load](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=145)  <- go1.19

``` go linenums="1"
func (x *Uint64) Load() uint64
```

​	Load方法原子地加载并返回存储在x中的值。

#### (*Uint64) [Store](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=148)  <- go1.19

``` go linenums="1"
func (x *Uint64) Store(val uint64)
```

​	Store方法原子地将val存储到x中。

#### (*Uint64) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=151)  <- go1.19

``` go linenums="1"
func (x *Uint64) Swap(new uint64) (old uint64)
```

​	Swap方法原子地将new存储到x中并返回先前的值。

### type [Uintptr](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=162)  <- go1.19

``` go linenums="1"
type Uintptr struct {
	// contains filtered or unexported fields
}
```

​	Uintptr是一个原子的uintptr类型。零值为零。

#### (*Uintptr) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=182)  <- go1.19

``` go linenums="1"
func (x *Uintptr) Add(delta uintptr) (new uintptr)
```

​	Add方法原子地将delta添加到x并返回新值。

#### (*Uintptr) [CompareAndSwap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=177)  <- go1.19

``` go linenums="1"
func (x *Uintptr) CompareAndSwap(old, new uintptr) (swapped bool)
```

​	CompareAndSwap方法执行x的比较和交换操作。

#### (*Uintptr) [Load](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=168)  <- go1.19

``` go linenums="1"
func (x *Uintptr) Load() uintptr
```

​	Load方法以原子方式加载并返回存储在x中的值。

#### (*Uintptr) [Store](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=171)  <- go1.19

``` go linenums="1"
func (x *Uintptr) Store(val uintptr)
```

​	Store方法以原子方式将val存储到x中。

#### (*Uintptr) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/type.go;l=174)  <- go1.19

``` go linenums="1"
func (x *Uintptr) Swap(new uintptr) (old uintptr)
```

​	Swap方法以原子方式将new存储到x中并返回先前的值。

### type [Value](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/value.go;l=16)  <- go1.4

``` go linenums="1"
type Value struct {
	// contains filtered or unexported fields
}
```

​	Value类型提供了一个一致类型值的原子加载和存储。Value类型的零值从Load返回nil。一旦调用了Store，就不能再复制Value。

​	使用后不能再复制Value。

##### Example (Config)

The following example shows how to use Value for periodic program config updates and propagation of the changes to worker goroutines.

``` go linenums="1"
package main

import (
	"sync/atomic"
	"time"
)

func loadConfig() map[string]string {
	return make(map[string]string)
}

func requests() chan int {
	return make(chan int)
}

func main() {
	var config atomic.Value // holds current server configuration
	// Create initial config value and store into config.
	config.Store(loadConfig())
	go func() {
		// Reload config every 10 seconds
		// and update config value with the new version.
		for {
			time.Sleep(10 * time.Second)
			config.Store(loadConfig())
		}
	}()
	// Create worker goroutines that handle incoming requests
	// using the latest config value.
	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				c := config.Load()
				// Handle request r using config c.
				_, _ = r, c
			}
		}()
	}
}

```

##### Example (ReadMostly) 

The following example shows how to maintain a scalable frequently read, but infrequently updated data structure using copy-on-write idiom.

``` go linenums="1"
package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))
	var mu sync.Mutex // used only by writers
	// read function can be used to read the data without further synchronization
	read := func(key string) (val string) {
		m1 := m.Load().(Map)
		return m1[key]
	}
	// insert function can be used to update the data without further synchronization
	insert := func(key, val string) {
		mu.Lock() // synchronize with other potential writers
		defer mu.Unlock()
		m1 := m.Load().(Map) // load current value of the data structure
		m2 := make(Map)      // create a new value
		for k, v := range m1 {
			m2[k] = v // copy all data from the current object to the new one
		}
		m2[key] = val // do the update that we need
		m.Store(m2)   // atomically replace the current object with the new one
		// At this point all new readers start working with the new version.
		// The old version will be garbage collected once the existing readers
		// (if any) are done with it.
	}
	_, _ = read, insert
}

```

#### (*Value) [CompareAndSwap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/value.go;l=135)  <- go1.17

``` go linenums="1"
func (v *Value) CompareAndSwap(old, new any) (swapped bool)
```

CompareAndSwap方法为Value执行比较和交换操作。

​	所有对给定Value的CompareAndSwap调用必须使用相同具体类型的值。类型不一致的CompareAndSwap会导致panic，CompareAndSwap(old, nil)也是如此。

#### (*Value) [Load](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/value.go;l=28)  <- go1.4

``` go linenums="1"
func (v *Value) Load() (val any)
```

​	Load方法返回最近一次Store设置的值。如果没有为此Value调用Store，则返回nil。

#### (*Value) [Store](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/value.go;l=47)  <- go1.4

``` go linenums="1"
func (v *Value) Store(val any)
```

​	Store方法将Value v的值设置为val。所有对给定Value的Store调用必须使用相同具体类型的值。类型不一致的Store会导致panic，Store(nil)也是如此。

#### (*Value) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sync/atomic/value.go;l=90)  <- go1.17

``` go linenums="1"
func (v *Value) Swap(new any) (old any)
```

​	Swap方法操作将new值存储到Value中，并返回旧值。如果Value为空，则返回nil。

​	对于同一个 Value，所有 Swap方法的调用必须使用相同的具体类型的值。如果使用不一致的类型进行 Swap方法，会导致 panic，就像 Swap(nil) 一样。

## Notes

## Bugs

- 在386上，64位函数使用Pentium MMX之前不可用的指令。

  在非Linux ARM上，64位函数使用ARMv6k核心之前不可用的指令。

  在ARM，386和32位MIPS上，通过原子函数访问64位字(类型Int64和Uint64自动对齐)的调用方有责任安排64位对齐。可以依靠分配的结构体，数组或切片中的第一个字，全局变量中的第一个字或局部变量中的第一个字(因为所有原子操作的主题都将逃逸到堆中)为64位对齐。