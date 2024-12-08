+++
title = "linkedlistqueue"
date = 2024-12-07T11:07:24+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/queues/linkedlistqueue](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/queues/linkedlistqueue)
>
> 收录该文档时间： `2024-12-07T11:07:24+08:00`

## Overview 

Package linkedlistqueue implements a queue backed by a singly-linked list.

Structure is not thread safe.

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

#### (*Iterator[T]) Begin 

``` go
func (iterator *Iterator[T]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

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

#### (*Iterator[T]) Value 

``` go
func (iterator *Iterator[T]) Value() T
```

Value returns the current element's value. Does not modify the state of the iterator.

### type Queue 

``` go
type Queue[T comparable] struct {
	// contains filtered or unexported fields
}
```

Queue holds elements in a singly-linked-list

#### func New 

``` go
func New[T comparable]() *Queue[T]
```

New instantiates a new empty queue

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

Peek returns first element of the queue without removing it, or nil if queue is empty. Second return parameter is true, unless the queue was empty and there was nothing to peek.

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

Values returns all elements in the queue (FIFO order).
