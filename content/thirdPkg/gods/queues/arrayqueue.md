+++
title = "arrayqueue"
date = 2024-12-07T11:06:50+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/queues/arrayqueue](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/queues/arrayqueue)
>
> 收录该文档时间： `2024-12-07T11:06:50+08:00`

## Overview 

Package arrayqueue implements a queue backed by array list.

​	包 `arrayqueue` 实现了一个基于数组列表的队列。

Structure is not thread safe.

​	该结构体不是线程安全的。

Reference: https://en.wikipedia.org/wiki/Queue_(abstract_data_type)

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Iterator 

``` go
type Iterator[T comparable] struct {
	// contains filtered or unexported fields
}
```

Iterator returns a stateful iterator whose values can be fetched by an index.

​	`Iterator` 返回一个状态化的迭代器，其值可以通过索引获取。

#### (*Iterator[T]) Begin 

``` go
func (iterator *Iterator[T]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

​	`Begin` 将迭代器重置到初始状态（第一个元素之前）。调用 `Next()` 以获取第一个元素（如果存在）。

#### (*Iterator[T]) End 

``` go
func (iterator *Iterator[T]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

​	`End` 将迭代器移至最后一个元素之后（超出范围）。调用 `Prev()` 以获取最后一个元素（如果存在）。

#### (*Iterator[T]) First 

``` go
func (iterator *Iterator[T]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`First` 将迭代器移至第一个元素并返回 `true`，如果容器中存在第一个元素。如果返回 `true`，可以通过 `Index()` 和 `Value()` 获取第一个元素的索引和值。此操作会修改迭代器的状态。

#### (*Iterator[T]) Index 

``` go
func (iterator *Iterator[T]) Index() int
```

Index returns the current element's index. Does not modify the state of the iterator.

​	`Index` 返回当前元素的索引。不会修改迭代器的状态。

#### (*Iterator[T]) Last 

``` go
func (iterator *Iterator[T]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`Last` 将迭代器移至最后一个元素并返回 `true`，如果容器中存在最后一个元素。如果返回 `true`，可以通过 `Index()` 和 `Value()` 获取最后一个元素的索引和值。此操作会修改迭代器的状态。

#### (*Iterator[T]) Next 

``` go
func (iterator *Iterator[T]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's index and value can be retrieved by Index() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

​	`Next` 将迭代器移至下一个元素并返回 `true`，如果容器中存在下一个元素。如果返回 `true`，可以通过 `Index()` 和 `Value()` 获取下一个元素的索引和值。若首次调用 `Next()`，迭代器将指向第一个元素（如果存在）。此操作会修改迭代器的状态。

#### (*Iterator[T]) NextTo 

``` go
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`NextTo` 将迭代器移至满足传递函数条件的下一个元素，并返回 `true`，如果容器中存在下一个元素。如果返回 `true`，可以通过 `Index()` 和 `Value()` 获取该元素的索引和值。此操作会修改迭代器的状态。

#### (*Iterator[T]) Prev 

``` go
func (iterator *Iterator[T]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`Prev` 将迭代器移至前一个元素并返回 `true`，如果容器中存在前一个元素。如果返回 `true`，可以通过 `Index()` 和 `Value()` 获取前一个元素的索引和值。此操作会修改迭代器的状态。

#### (*Iterator[T]) PrevTo 

``` go
func (iterator *Iterator[T]) PrevTo(f func(index int, value T) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`PrevTo` 将迭代器移至满足传递函数条件的前一个元素，并返回 `true`，如果容器中存在前一个元素。如果返回 `true`，可以通过 `Index()` 和 `Value()` 获取该元素的索引和值。此操作会修改迭代器的状态。

#### (*Iterator[T]) Value 

``` go
func (iterator *Iterator[T]) Value() T
```

Value returns the current element's value. Does not modify the state of the iterator.

​	`Value` 返回当前元素的值。不会修改迭代器的状态。type Queue 

``` go
type Queue[T comparable] struct {
	// contains filtered or unexported fields
}
```

Queue holds elements in an array-list

​	`Queue` 使用数组列表存储元素。

#### func New 

``` go
func New[T comparable]() *Queue[T]
```

New instantiates a new empty queue

​	`New` 实例化一个新的空队列。

#### (*Queue[T]) Clear 

``` go
func (queue *Queue[T]) Clear()
```

Clear removes all elements from the queue.

​	`Clear` 移除队列中的所有元素。

#### (*Queue[T]) Dequeue 

``` go
func (queue *Queue[T]) Dequeue() (value T, ok bool)
```

Dequeue removes first element of the queue and returns it, or nil if queue is empty. Second return parameter is true, unless the queue was empty and there was nothing to dequeue.

​	`Dequeue` 移除队列中的第一个元素并返回。如果队列为空，则返回 `nil`，且第二个返回值为 `false`。

#### (*Queue[T]) Empty 

``` go
func (queue *Queue[T]) Empty() bool
```

Empty returns true if queue does not contain any elements.

​	`Empty` 返回 `true`，如果队列中没有元素。

#### (*Queue[T]) Enqueue 

``` go
func (queue *Queue[T]) Enqueue(value T)
```

Enqueue adds a value to the end of the queue

​	`Enqueue` 将一个值添加到队列的末尾。

#### (*Queue[T]) FromJSON 

``` go
func (queue *Queue[T]) FromJSON(data []byte) error
```

FromJSON populates the queue from the input JSON representation.

​	`FromJSON` 从输入的 JSON 表示中填充队列。

#### (*Queue[T]) Iterator 

``` go
func (queue *Queue[T]) Iterator() *Iterator[T]
```

Iterator returns a stateful iterator whose values can be fetched by an index.

​	`Iterator` 返回一个状态化的迭代器，其值可以通过索引获取。

#### (*Queue[T]) MarshalJSON 

``` go
func (queue *Queue[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

​	`MarshalJSON` 实现了 `json.Marshaler` 接口。

#### (*Queue[T]) Peek 

``` go
func (queue *Queue[T]) Peek() (value T, ok bool)
```

Peek returns first element of the queue without removing it, or nil if queue is empty. Second return parameter is true, unless the queue was empty and there was nothing to peek.

​	`Peek` 返回队列中的第一个元素，但不移除它。如果队列为空，则返回 `nil`，且第二个返回值为 `false`。

#### (*Queue[T]) Size 

``` go
func (queue *Queue[T]) Size() int
```

Size returns number of elements within the queue.

​	`Size` 返回队列中的元素数量。

#### (*Queue[T]) String 

``` go
func (queue *Queue[T]) String() string
```

String returns a string representation of container

​	`String` 返回容器的字符串表示。

#### (*Queue[T]) ToJSON 

``` go
func (queue *Queue[T]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the queue.

​	`ToJSON` 输出队列的 JSON 表示。

#### (*Queue[T]) UnmarshalJSON 

``` go
func (queue *Queue[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

​	`UnmarshalJSON` 实现了 `json.Unmarshaler` 接口。

#### (*Queue[T]) Values 

``` go
func (queue *Queue[T]) Values() []T
```

Values returns all elements in the queue (FIFO order).

​	`Values` 返回队列中的所有元素（FIFO 顺序）。
