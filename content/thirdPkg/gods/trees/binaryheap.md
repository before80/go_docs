+++
title = "binaryheap"
date = 2024-12-07T11:09:41+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/trees/binaryheap](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/trees/binaryheap)
>
> 收录该文档时间： `2024-12-07T11:09:41+08:00`

## Overview 

Package binaryheap implements a binary heap backed by array list.

​	包 `binaryheap` 实现了一个基于数组列表的二叉堆。

Comparator defines this heap as either min or max heap.

​	由 `Heap` 类型管理堆的操作，可以是最小堆或最大堆，取决于 `Comparator`。

Structure is not thread safe.

​	结构不具备线程安全。

References: http://en.wikipedia.org/wiki/Binary_heap

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Heap 

``` go
type Heap[T comparable] struct {
	Comparator utils.Comparator[T]
	// contains filtered or unexported fields
}
```

Heap holds elements in an array-list

​	Heap 存储元素在一个数组列表中。

#### func New 

``` go
func New[T cmp.Ordered]() *Heap[T]
```

New instantiates a new empty heap tree with the built-in comparator for T

​	New 使用内置的比较器实例化一个新的空堆树。

#### func NewWith 

``` go
func NewWith[T comparable](comparator utils.Comparator[T]) *Heap[T]
```

NewWith instantiates a new empty heap tree with the custom comparator.

​	NewWith 使用自定义比较器实例化一个新的空堆树。

#### (*Heap[T]) Clear 

``` go
func (heap *Heap[T]) Clear()
```

Clear removes all elements from the heap.

​	Clear 移除堆中的所有元素。

#### (*Heap[T]) Empty 

``` go
func (heap *Heap[T]) Empty() bool
```

Empty returns true if heap does not contain any elements.

​	Empty 如果堆中没有任何元素则返回 `true`。

#### (*Heap[T]) FromJSON 

``` go
func (heap *Heap[T]) FromJSON(data []byte) error
```

FromJSON populates the heap from the input JSON representation.

​	FromJSON 从输入的 JSON 表示中填充堆。

#### (*Heap[T]) Iterator 

``` go
func (heap *Heap[T]) Iterator() *Iterator[T]
```

Iterator returns a stateful iterator whose values can be fetched by an index.

​	Iterator 返回一个有状态的迭代器，索引可以用来获取值。

#### (*Heap[T]) MarshalJSON 

``` go
func (heap *Heap[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

​	MarshalJSON 实现了 `json.Marshaler` 接口。

#### (*Heap[T]) Peek 

``` go
func (heap *Heap[T]) Peek() (value T, ok bool)
```

Peek returns top element on the heap without removing it, or nil if heap is empty. Second return parameter is true, unless the heap was empty and there was nothing to peek.

​	Peek 返回堆顶元素而不移除它，如果堆为空则返回 `nil`。第二个返回参数为 `true`，除非堆为空，没有元素可供查看。

#### (*Heap[T]) Pop 

``` go
func (heap *Heap[T]) Pop() (value T, ok bool)
```

Pop removes top element on heap and returns it, or nil if heap is empty. Second return parameter is true, unless the heap was empty and there was nothing to pop.

​	Pop 移除堆顶元素并返回它，如果堆为空则返回 `nil`。第二个返回参数为 `true`，除非堆为空，没有元素可供移除。

#### (*Heap[T]) Push 

``` go
func (heap *Heap[T]) Push(values ...T)
```

Push adds a value onto the heap and bubbles it up accordingly.

​	Push 将一个或多个元素添加到堆中，并根据堆的规则进行上浮操作。

#### (*Heap[T]) Size 

``` go
func (heap *Heap[T]) Size() int
```

Size returns number of elements within the heap.

​	Size 返回堆中的元素个数。

#### (*Heap[T]) String 

``` go
func (heap *Heap[T]) String() string
```

String returns a string representation of container

​	String 返回容器的字符串表示。

#### (*Heap[T]) ToJSON 

``` go
func (heap *Heap[T]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the heap.

​	ToJSON 输出堆的 JSON 表示。

#### (*Heap[T]) UnmarshalJSON 

``` go
func (heap *Heap[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

​	UnmarshalJSON 实现了 `json.Unmarshaler` 接口。

#### (*Heap[T]) Values 

``` go
func (heap *Heap[T]) Values() []T
```

Values returns all elements in the heap.

​	Values 返回堆中所有元素。

### type Iterator 

``` go
type Iterator[T comparable] struct {
	// contains filtered or unexported fields
}
```

Iterator returns a stateful iterator whose values can be fetched by an index.

​	Iterator 返回一个有状态的迭代器，索引可以用来获取值。

#### (*Iterator[T]) Begin 

``` go
func (iterator *Iterator[T]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

​	Begin 将迭代器重置为初始状态（即“第一个元素之前”）。调用 `Next()` 获取第一个元素（如果有）。

#### (*Iterator[T]) End 

``` go
func (iterator *Iterator[T]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

​	End 将迭代器移到最后一个元素之后（即“最后一个元素之后”）。调用 `Prev()` 获取最后一个元素（如果有）。

#### (*Iterator[T]) First 

``` go
func (iterator *Iterator[T]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	First 将迭代器移到第一个元素，并返回 `true` 如果容器中存在第一个元素。如果 `First()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取第一个元素的索引和值。会修改迭代器的状态。

#### (*Iterator[T]) Index 

``` go
func (iterator *Iterator[T]) Index() int
```

Index returns the current element's index. Does not modify the state of the iterator.

​	Index 返回当前元素的索引。不会修改迭代器的状态。

#### (*Iterator[T]) Last 

``` go
func (iterator *Iterator[T]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	Last 将迭代器移到最后一个元素，并返回 `true` 如果容器中存在最后一个元素。如果 `Last()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取最后一个元素的索引和值。会修改迭代器的状态。

#### (*Iterator[T]) Next 

``` go
func (iterator *Iterator[T]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's index and value can be retrieved by Index() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

​	Next 将迭代器移到下一个元素，并返回 `true` 如果容器中存在下一个元素。如果 `Next()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取下一个元素的索引和值。如果是第一次调用 `Next()`，它将使迭代器指向第一个元素（如果存在）。会修改迭代器的状态。

#### (*Iterator[T]) NextTo 

``` go
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	NextTo 将迭代器移到当前元素位置之后的下一个满足传入函数条件的元素，并返回 `true` 如果容器中存在下一个元素。如果 `NextTo()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取下一个元素的索引和值。会修改迭代器的状态。

#### (*Iterator[T]) Prev 

``` go
func (iterator *Iterator[T]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	NextTo 将迭代器移到当前元素位置之后的下一个满足传入函数条件的元素，并返回 `true` 如果容器中存在下一个元素。如果 `NextTo()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取下一个元素的索引和值。会修改迭代器的状态。

#### (*Iterator[T]) PrevTo 

``` go
func (iterator *Iterator[T]) PrevTo(f func(index int, value T) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	PrevTo 将迭代器移到当前元素位置之前的上一个满足传入函数条件的元素，并返回 `true` 如果容器中存在上一个元素。如果 `PrevTo()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取上一个元素的索引和值。会修改迭代器的状态。

​	

#### (*Iterator[T]) Value 

``` go
func (iterator *Iterator[T]) Value() T
```

Value returns the current element's value. Does not modify the state of the iterator.

​	Value 返回当前元素的值。不会修改迭代器的状态。
