+++
title = "gring"
date = 2024-03-21T17:44:58+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gring](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gring)

Package gring provides a concurrent-safe/unsafe ring(circular lists).

​	软件包 gring 提供并发安全/不安全环（循环列表）。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Ring

```go
type Ring struct {
	// contains filtered or unexported fields
}
```

Ring is a struct of ring structure.

​	环是环结构的结构。

#### func New

```go
func New(cap int, safe ...bool) *Ring
```

New creates and returns a Ring structure of `cap` elements. The optional parameter `safe` specifies whether using this structure in concurrent safety, which is false in default.

​	New 创建并返回元素的 `cap` Ring 结构。可选参数 `safe` 指定是否在并发安全中使用此结构，默认为 false。

##### Example

``` go
```

#### (*Ring) Cap

```go
func (r *Ring) Cap() int
```

Cap returns the capacity of ring.

​	Cap 返回环的容量。

##### Example

``` go
```

#### (*Ring) Len

```go
func (r *Ring) Len() int
```

Len returns the size of ring.

​	Len 返回戒指的大小。

##### Example

``` go
```

#### (*Ring) Link

```go
func (r *Ring) Link(s *Ring) *Ring
```

Link connects ring r with ring s such that r.Next() becomes s and returns the original value for r.Next(). r must not be empty.

​	Link 将环 r 与环 s 连接起来，使得 r.Next（） 变为 s，并返回 r.Next（） 的原始值。r 不能为空。

If r and s point to the same ring, linking them removes the elements between r and s from the ring. The removed elements form a sub-ring and the result is a reference to that sub-ring (if no elements were removed, the result is still the original value for r.Next(), and not nil).

​	如果 r 和 s 指向同一个环，则将它们链接起来会从环中删除 r 和 s 之间的元素。删除的元素形成一个子环，结果是对该子环的引用（如果未删除任何元素，则结果仍然是 r.Next（） 的原始值，而不是 nil）。

If r and s point to different rings, linking them creates a single ring with the elements of s inserted after r. The result points to the element following the last element of s after insertion.

​	如果 r 和 s 指向不同的环，则将它们链接起来会创建一个环，其中 s 的元素插入到 r 之后。结果指向插入后 s 的最后一个元素后面的元素。

#### (*Ring) Move

```go
func (r *Ring) Move(n int) *Ring
```

Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0) in the ring and returns that ring element. r must not be empty.

​	Move 在环中向后移动 （n < 0） 或向前移动 （n >= 0） 元素，并返回该环元素。 r 不能为空。

##### Example

``` go
```

#### (*Ring) Next

```go
func (r *Ring) Next() *Ring
```

Next returns the next ring element. r must not be empty.

​	Next 返回下一个环元素。r 不能为空。

##### Example

``` go
```

#### (*Ring) Prev

```go
func (r *Ring) Prev() *Ring
```

Prev returns the previous ring element. r must not be empty.

​	Prev 返回上一个环元素。r 不能为空。

##### Example

``` go
```

#### (*Ring) Put

```go
func (r *Ring) Put(value interface{}) *Ring
```

Put sets `value` to current item of ring and moves position to next item.

​	将设置 `value` 到戒指的当前项目，并将位置移动到下一个项目。

##### Example

``` go
```

#### (*Ring) RLockIteratorNext

```go
func (r *Ring) RLockIteratorNext(f func(value interface{}) bool)
```

RLockIteratorNext iterates and locks reading forward with given callback function `f` within RWMutex.RLock. If `f` returns true, then it continues iterating; or false to stop.

​	RLockIteratorNext 使用 RWMutex.RLock 中的给定回调函数 `f` 迭代并锁定前向读取。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*Ring) RLockIteratorPrev

```go
func (r *Ring) RLockIteratorPrev(f func(value interface{}) bool)
```

RLockIteratorPrev iterates and locks writing backward with given callback function `f` within RWMutex.RLock. If `f` returns true, then it continues iterating; or false to stop.

​	RLockIteratorPrev 使用 RWMutex.RLock 中的给定回调函数 `f` 迭代并锁定向后写入。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*Ring) Set

```go
func (r *Ring) Set(value interface{}) *Ring
```

Set sets value to the item of current position.

​	将设置值设置为当前位置的项目。

##### Example

``` go
```

#### (*Ring) SliceNext

```go
func (r *Ring) SliceNext() []interface{}
```

SliceNext returns a copy of all item values as slice forward from current position.

​	SliceNext 返回所有项值的副本，作为从当前位置向前切片。

##### Example

``` go
```

#### (*Ring) SlicePrev

```go
func (r *Ring) SlicePrev() []interface{}
```

SlicePrev returns a copy of all item values as slice backward from current position.

​	SlicePrev 从当前位置向后返回所有项值的副本。

##### Example

``` go
```

#### (*Ring) Unlink

```go
func (r *Ring) Unlink(n int) *Ring
```

Unlink removes n % r.Len() elements from the ring r, starting at r.Next(). If n % r.Len() == 0, r remains unchanged. The result is the removed sub-ring. r must not be empty.

​	Unlink 从环 r 中删除 n % r.Len（） 元素，从 r.Next（） 开始。如果 n % r.Len（） == 0，则 r 保持不变。结果是移除的子环。r 不能为空。

##### Example

``` go
```

#### (*Ring) Val

```go
func (r *Ring) Val() interface{}
```

Val returns the item’s value of current position.

​	Val 返回项目当前位置的值。

Example Val

​	示例 Val

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/gring"
)

func main() {
	r := gring.New(10)
	r.Set(1)
	fmt.Println("Val:", r.Val())

	r.Next().Set("GoFrame")
	fmt.Println("Val:", r.Val())

}

Output:

Val: 1
Val: GoFrame
```







