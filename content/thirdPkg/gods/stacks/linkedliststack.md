+++
title = "linkedliststack"
date = 2024-12-07T11:55:20+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/stacks/linkedliststack](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/stacks/linkedliststack)
>
> 收录该文档时间： `2024-12-07T11:55:20+08:00`

## Overview 

Package linkedliststack implements a stack backed by a singly-linked list.

Structure is not thread safe.

Reference:[https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Linked_list](https://en.wikipedia.org/wiki/Stack_(abstract_data_type)#Linked_list)

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

### type Stack 

``` go
type Stack[T comparable] struct {
	// contains filtered or unexported fields
}
```

Stack holds elements in a singly-linked-list

#### func New 

``` go
func New[T comparable]() *Stack[T]
```

New nnstantiates a new empty stack

#### (*Stack[T]) Clear 

``` go
func (stack *Stack[T]) Clear()
```

Clear removes all elements from the stack.

#### (*Stack[T]) Empty 

``` go
func (stack *Stack[T]) Empty() bool
```

Empty returns true if stack does not contain any elements.

#### (*Stack[T]) FromJSON 

``` go
func (stack *Stack[T]) FromJSON(data []byte) error
```

FromJSON populates the stack from the input JSON representation.

#### (*Stack[T]) Iterator 

``` go
func (stack *Stack[T]) Iterator() *Iterator[T]
```

Iterator returns a stateful iterator whose values can be fetched by an index.

#### (*Stack[T]) MarshalJSON 

``` go
func (stack *Stack[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### (*Stack[T]) Peek 

``` go
func (stack *Stack[T]) Peek() (value T, ok bool)
```

Peek returns top element on the stack without removing it, or nil if stack is empty. Second return parameter is true, unless the stack was empty and there was nothing to peek.

#### (*Stack[T]) Pop 

``` go
func (stack *Stack[T]) Pop() (value T, ok bool)
```

Pop removes top element on stack and returns it, or nil if stack is empty. Second return parameter is true, unless the stack was empty and there was nothing to pop.

#### (*Stack[T]) Push 

``` go
func (stack *Stack[T]) Push(value T)
```

Push adds a value onto the top of the stack

#### (*Stack[T]) Size 

``` go
func (stack *Stack[T]) Size() int
```

Size returns number of elements within the stack.

#### (*Stack[T]) String 

``` go
func (stack *Stack[T]) String() string
```

String returns a string representation of container

#### (*Stack[T]) ToJSON 

``` go
func (stack *Stack[T]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the stack.

#### (*Stack[T]) UnmarshalJSON 

``` go
func (stack *Stack[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### (*Stack[T]) Values 

``` go
func (stack *Stack[T]) Values() []T
```

Values returns all elements in the stack (LIFO order).
