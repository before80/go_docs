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

Comparator defines this heap as either min or max heap.

Structure is not thread safe.

References: http://en.wikipedia.org/wiki/Binary_heap

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

#### type Heap 

``` go
type Heap[T comparable] struct {
	Comparator utils.Comparator[T]
	// contains filtered or unexported fields
}
```

Heap holds elements in an array-list

#### func New 

``` go
func New[T cmp.Ordered]() *Heap[T]
```

New instantiates a new empty heap tree with the built-in comparator for T

#### func NewWith 

``` go
func NewWith[T comparable](comparator utils.Comparator[T]) *Heap[T]
```

NewWith instantiates a new empty heap tree with the custom comparator.

#### (*Heap[T]) Clear 

``` go
func (heap *Heap[T]) Clear()
```

Clear removes all elements from the heap.

#### (*Heap[T]) Empty 

``` go
func (heap *Heap[T]) Empty() bool
```

Empty returns true if heap does not contain any elements.

#### (*Heap[T]) FromJSON 

``` go
func (heap *Heap[T]) FromJSON(data []byte) error
```

FromJSON populates the heap from the input JSON representation.

#### (*Heap[T]) Iterator 

``` go
func (heap *Heap[T]) Iterator() *Iterator[T]
```

Iterator returns a stateful iterator whose values can be fetched by an index.

#### (*Heap[T]) MarshalJSON 

``` go
func (heap *Heap[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### (*Heap[T]) Peek 

``` go
func (heap *Heap[T]) Peek() (value T, ok bool)
```

Peek returns top element on the heap without removing it, or nil if heap is empty. Second return parameter is true, unless the heap was empty and there was nothing to peek.

#### (*Heap[T]) Pop 

``` go
func (heap *Heap[T]) Pop() (value T, ok bool)
```

Pop removes top element on heap and returns it, or nil if heap is empty. Second return parameter is true, unless the heap was empty and there was nothing to pop.

#### (*Heap[T]) Push 

``` go
func (heap *Heap[T]) Push(values ...T)
```

Push adds a value onto the heap and bubbles it up accordingly.

#### (*Heap[T]) Size 

``` go
func (heap *Heap[T]) Size() int
```

Size returns number of elements within the heap.

#### (*Heap[T]) String 

``` go
func (heap *Heap[T]) String() string
```

String returns a string representation of container

#### (*Heap[T]) ToJSON 

``` go
func (heap *Heap[T]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the heap.

#### (*Heap[T]) UnmarshalJSON 

``` go
func (heap *Heap[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### (*Heap[T]) Values 

``` go
func (heap *Heap[T]) Values() []T
```

Values returns all elements in the heap.

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
