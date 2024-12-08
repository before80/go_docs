+++
title = "arraystack"
date = 2024-12-07T11:55:05+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/stacks/arraystack](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/stacks/arraystack)
>
> 收录该文档时间： `2024-12-07T11:55:05+08:00`

## Overview 

Package arraystack implements a stack backed by array list.

​	包 `arraystack` 实现了一个基于数组列表的栈。

Structure is not thread safe.

​	该结构不是线程安全的。

Reference: [https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array](https://en.wikipedia.org/wiki/Stack_(abstract_data_type)#Array)

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

​	`Iterator` 提供了一个可通过索引获取值的状态化迭代器。(*Iterator[T]) Begin 

``` go
func (iterator *Iterator[T]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

​	`Begin` 将迭代器重置为其初始状态（位于第一个元素之前）。调用 `Next()` 获取第一个元素（如果有）。

#### (*Iterator[T]) End 

``` go
func (iterator *Iterator[T]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

​	`End` 将迭代器移动到最后一个元素之后（超出范围）。调用 `Prev()` 获取最后一个元素（如果有）。

#### (*Iterator[T]) First 

``` go
func (iterator *Iterator[T]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`First` 将迭代器移动到第一个元素，并返回 `true`，如果存在第一个元素。如果 `First()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取第一个元素的索引和值。修改了迭代器的状态。

#### (*Iterator[T]) Index 

``` go
func (iterator *Iterator[T]) Index() int
```

Index returns the current element's index. Does not modify the state of the iterator.

​	`Index` 返回当前元素的索引。不修改迭代器的状态。

#### (*Iterator[T]) Last 

``` go
func (iterator *Iterator[T]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`Last` 将迭代器移动到最后一个元素，并返回 `true`，如果存在最后一个元素。如果 `Last()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取最后一个元素的索引和值。修改了迭代器的状态。

#### (*Iterator[T]) Next 

``` go
func (iterator *Iterator[T]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's index and value can be retrieved by Index() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

​	`Next` 将迭代器移动到下一个元素，并返回 `true`，如果存在下一个元素。如果 `Next()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取下一个元素的索引和值。如果这是首次调用 `Next()`，迭代器将指向第一个元素（如果存在）。修改了迭代器的状态。

#### (*Iterator[T]) NextTo 

``` go
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`NextTo` 将迭代器从当前位置移动到满足给定条件的下一个元素，并返回 `true`，如果存在下一个元素。如果 `NextTo()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取下一个元素的索引和值。修改了迭代器的状态。	

#### (*Iterator[T]) Prev 

``` go
func (iterator *Iterator[T]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`Prev` 将迭代器移动到前一个元素，并返回 `true`，如果存在前一个元素。如果 `Prev()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取前一个元素的索引和值。修改了迭代器的状态。

#### (*Iterator[T]) PrevTo 

``` go
func (iterator *Iterator[T]) PrevTo(f func(index int, value T) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`PrevTo` 将迭代器从当前位置移动到满足给定条件的前一个元素，并返回 `true`，如果存在前一个元素。如果 `PrevTo()` 返回 `true`，则可以通过 `Index()` 和 `Value()` 获取前一个元素的索引和值。修改了迭代器的状态。

#### (*Iterator[T]) Value 

``` go
func (iterator *Iterator[T]) Value() T
```

Value returns the current element's value. Does not modify the state of the iterator.

​	`Value` 返回当前元素的值。不修改迭代器的状态。

### type Stack 

``` go
type Stack[T comparable] struct {
	// contains filtered or unexported fields
}
```

Stack holds elements in an array-list

​	`Stack` 在一个数组列表中保存元素。

#### func New 

``` go
func New[T comparable]() *Stack[T]
```

New instantiates a new empty stack

​	`New` 实例化一个新的空栈。

#### (*Stack[T]) Clear 

``` go
func (stack *Stack[T]) Clear()
```

Clear removes all elements from the stack.

​	`Clear` 移除栈中的所有元素。

#### (*Stack[T]) Empty 

``` go
func (stack *Stack[T]) Empty() bool
```

Empty returns true if stack does not contain any elements.

​	`Empty` 如果栈中没有任何元素，返回 `true`。

#### (*Stack[T]) FromJSON 

``` go
func (stack *Stack[T]) FromJSON(data []byte) error
```

FromJSON populates the stack from the input JSON representation.

​	`FromJSON` 根据输入的 JSON 表示填充栈。

#### (*Stack[T]) Iterator 

``` go
func (stack *Stack[T]) Iterator() *Iterator[T]
```

Iterator returns a stateful iterator whose values can be fetched by an index.

​	`Iterator` 返回一个状态化迭代器，可通过索引获取值。

#### (*Stack[T]) MarshalJSON 

``` go
func (stack *Stack[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

​	`MarshalJSON` @实现了 `json.Marshaler` 接口。

#### (*Stack[T]) Peek 

``` go
func (stack *Stack[T]) Peek() (value T, ok bool)
```

Peek returns top element on the stack without removing it, or nil if stack is empty. Second return parameter is true, unless the stack was empty and there was nothing to peek.

​	`Peek` 返回栈顶元素但不移除它，或者如果栈为空则返回 `nil`。第二个返回值为 `true`，除非栈为空。

#### (*Stack[T]) Pop 

``` go
func (stack *Stack[T]) Pop() (value T, ok bool)
```

Pop removes top element on stack and returns it, or nil if stack is empty. Second return parameter is true, unless the stack was empty and there was nothing to pop.

​	`Pop` 移除栈顶元素并返回它，或者如果栈为空则返回 `nil`。第二个返回值为 `true`，除非栈为空。

#### (*Stack[T]) Push 

``` go
func (stack *Stack[T]) Push(value T)
```

Push adds a value onto the top of the stack

​	`Push` 将一个值添加到栈顶。

#### (*Stack[T]) Size 

``` go
func (stack *Stack[T]) Size() int
```

Size returns number of elements within the stack.

​	`Size` 返回栈中元素的数量。

#### (*Stack[T]) String 

``` go
func (stack *Stack[T]) String() string
```

String returns a string representation of container

​	`String` 返回容器的字符串表示。

#### (*Stack[T]) ToJSON 

``` go
func (stack *Stack[T]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the stack.

​	`ToJSON` 输出栈的 JSON 表示。

#### (*Stack[T]) UnmarshalJSON 

``` go
func (stack *Stack[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

​	`UnmarshalJSON` @实现了 `json.Unmarshaler` 接口。

#### (*Stack[T]) Values 

``` go
func (stack *Stack[T]) Values() []T
```

Values returns all elements in the stack (LIFO order).

​	`Values` 返回栈中的所有元素（按 LIFO 顺序）。
