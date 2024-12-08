+++
title = "priorityqueue"
date = 2024-12-07T11:07:42+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/queues/priorityqueue](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/queues/priorityqueue)
>
> 收录该文档时间： `2024-12-07T11:07:42+08:00`

## Overview 

Package priorityqueue implements a priority queue backed by binary queue.

An unbounded priority queue based on a priority queue. The elements of the priority queue are ordered by a comparator provided at queue construction time.

The heap of this queue is the least/smallest element with respect to the specified ordering. If multiple elements are tied for least value, the heap is one of those elements arbitrarily.

Structure is not thread safe.

References: https://en.wikipedia.org/wiki/Priority_queue

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

#### type Iterator 

``` go
type Iterator[T comparable] struct {
	// contains filtered or unexported fields
}
```

Iterator returns a stateful iterator whose values can be fetched by an index.

#### (*Iterator[T]) Begin 

``` go
func (iterator *Iterator[T]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

#### (*Iterator[T]) End 

``` go
func (iterator *Iterator[T]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

#### (*Iterator[T]) First 

``` go
func (iterator *Iterator[T]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

#### (*Iterator[T]) Index 

``` go
func (iterator *Iterator[T]) Index() int
```

Index returns the current element's index. Does not modify the state of the iterator.

#### (*Iterator[T]) Last 

``` go
func (iterator *Iterator[T]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

#### (*Iterator[T]) Next 

``` go
func (iterator *Iterator[T]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's index and value can be retrieved by Index() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

#### (*Iterator[T]) NextTo 

``` go
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

#### (*Iterator[T]) Prev 

``` go
func (iterator *Iterator[T]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

#### (*Iterator[T]) PrevTo 

``` go
func (iterator *Iterator[T]) PrevTo(f func(index int, value T) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

#### (*Iterator[T]) Value 

``` go
func (iterator *Iterator[T]) Value() T
```

Value returns the current element's value. Does not modify the state of the iterator.

#### type Queue 

``` go
type Queue[T comparable] struct {
	Comparator utils.Comparator[T]
	// contains filtered or unexported fields
}
```

Queue holds elements in an array-list

#### func New 

``` go
func New[T cmp.Ordered]() *Queue[T]
```

#### func NewWith 

``` go
func NewWith[T comparable](comparator utils.Comparator[T]) *Queue[T]
```

NewWith instantiates a new empty queue with the custom comparator.

#### (*Queue[T]) Clear 

``` go
func (queue *Queue[T]) Clear()
```

Clear removes all elements from the queue.

#### (*Queue[T]) Dequeue 

``` go
func (queue *Queue[T]) Dequeue() (value T, ok bool)
```

Dequeue removes first element of the queue and returns it, or nil if queue is empty. Second return parameter is true, unless the queue was empty and there was nothing to dequeue.

#### (*Queue[T]) Empty 

``` go
func (queue *Queue[T]) Empty() bool
```

Empty returns true if queue does not contain any elements.

#### (*Queue[T]) Enqueue 

``` go
func (queue *Queue[T]) Enqueue(value T)
```

Enqueue adds a value to the end of the queue

#### (*Queue[T]) FromJSON 

``` go
func (queue *Queue[T]) FromJSON(data []byte) error
```

FromJSON populates the queue from the input JSON representation.

#### (*Queue[T]) Iterator 

``` go
func (queue *Queue[T]) Iterator() *Iterator[T]
```

Iterator returns a stateful iterator whose values can be fetched by an index.

#### (*Queue[T]) MarshalJSON 

``` go
func (queue *Queue[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### (*Queue[T]) Peek 

``` go
func (queue *Queue[T]) Peek() (value T, ok bool)
```

Peek returns top element on the queue without removing it, or nil if queue is empty. Second return parameter is true, unless the queue was empty and there was nothing to peek.

#### (*Queue[T]) Size 

``` go
func (queue *Queue[T]) Size() int
```

Size returns number of elements within the queue.

#### (*Queue[T]) String 

``` go
func (queue *Queue[T]) String() string
```

String returns a string representation of container

#### (*Queue[T]) ToJSON 

``` go
func (queue *Queue[T]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the queue.

#### (*Queue[T]) UnmarshalJSON 

``` go
func (queue *Queue[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### (*Queue[T]) Values 

``` go
func (queue *Queue[T]) Values() []T
```

Values returns all elements in the queue.
