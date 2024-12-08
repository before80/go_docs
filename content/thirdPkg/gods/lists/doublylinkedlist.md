+++
title = "doublylinkedlist"
date = 2024-12-07T11:02:08+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/lists/doublylinkedlist](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/lists/doublylinkedlist)
>
> 收录该文档时间： `2024-12-07T11:02:08+08:00`

## Overview 

Package doublylinkedlist implements the doubly-linked list.

Structure is not thread safe.

Reference: [https://en.wikipedia.org/wiki/List_%28abstract_data_type%29](https://en.wikipedia.org/wiki/List_(abstract_data_type))

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

Iterator holding the iterator's state

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

### type List 

``` go
type List[T comparable] struct {
	// contains filtered or unexported fields
}
```

List holds the elements, where each element points to the next and previous element

#### func New 

``` go
func New[T comparable](values ...T) *List[T]
```

New instantiates a new list and adds the passed values, if any, to the list

#### (*List[T]) Add 

``` go
func (list *List[T]) Add(values ...T)
```

Add appends a value (one or more) at the end of the list (same as Append())

#### (*List[T]) All 

``` go
func (list *List[T]) All(f func(index int, value T) bool) bool
```

All passes each element of the container to the given function and returns true if the function returns true for all elements.

#### (*List[T]) Any 

``` go
func (list *List[T]) Any(f func(index int, value T) bool) bool
```

Any passes each element of the container to the given function and returns true if the function ever returns true for any element.

#### (*List[T]) Append 

``` go
func (list *List[T]) Append(values ...T)
```

Append appends a value (one or more) at the end of the list (same as Add())

#### (*List[T]) Clear 

``` go
func (list *List[T]) Clear()
```

Clear removes all elements from the list.

#### (*List[T]) Contains 

``` go
func (list *List[T]) Contains(values ...T) bool
```

Contains check if values (one or more) are present in the set. All values have to be present in the set for the method to return true. Performance time complexity of n^2. Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.

#### (*List[T]) Each 

``` go
func (list *List[T]) Each(f func(index int, value T))
```

Each calls the given function once for each element, passing that element's index and value.

#### (*List[T]) Empty 

``` go
func (list *List[T]) Empty() bool
```

Empty returns true if list does not contain any elements.

#### (*List[T]) Find 

``` go
func (list *List[T]) Find(f func(index int, value T) bool) (index int, value T)
```

Find passes each element of the container to the given function and returns the first (index,value) for which the function is true or -1,nil otherwise if no element matches the criteria.

#### (*List[T]) FromJSON 

``` go
func (list *List[T]) FromJSON(data []byte) error
```

FromJSON populates list's elements from the input JSON representation.

#### (*List[T]) Get 

``` go
func (list *List[T]) Get(index int) (T, bool)
```

Get returns the element at index. Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.

#### (*List[T]) IndexOf 

``` go
func (list *List[T]) IndexOf(value T) int
```

IndexOf returns index of provided element

#### (*List[T]) Insert 

``` go
func (list *List[T]) Insert(index int, values ...T)
```

Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right. Does not do anything if position is negative or bigger than list's size Note: position equal to list's size is valid, i.e. append.

#### (*List[T]) Iterator 

``` go
func (list *List[T]) Iterator() Iterator[T]
```

Iterator returns a stateful iterator whose values can be fetched by an index.

#### (*List[T]) Map 

``` go
func (list *List[T]) Map(f func(index int, value T) T) *List[T]
```

Map invokes the given function once for each element and returns a container containing the values returned by the given function.

#### (*List[T]) MarshalJSON 

``` go
func (list *List[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### (*List[T]) Prepend 

``` go
func (list *List[T]) Prepend(values ...T)
```

Prepend prepends a values (or more)

#### (*List[T]) Remove 

``` go
func (list *List[T]) Remove(index int)
```

Remove removes the element at the given index from the list.

#### (*List[T]) Select 

``` go
func (list *List[T]) Select(f func(index int, value T) bool) *List[T]
```

Select returns a new container containing all elements for which the given function returns a true value.

#### (*List[T]) Set 

``` go
func (list *List[T]) Set(index int, value T)
```

Set value at specified index position Does not do anything if position is negative or bigger than list's size Note: position equal to list's size is valid, i.e. append.

#### (*List[T]) Size 

``` go
func (list *List[T]) Size() int
```

Size returns number of elements within the list.

#### (*List[T]) Sort 

``` go
func (list *List[T]) Sort(comparator utils.Comparator[T])
```

Sort sorts values (in-place) using.

#### (*List[T]) String 

``` go
func (list *List[T]) String() string
```

String returns a string representation of container

#### (*List[T]) Swap 

``` go
func (list *List[T]) Swap(i, j int)
```

Swap swaps values of two elements at the given indices.

#### (*List[T]) ToJSON 

``` go
func (list *List[T]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of list's elements.

#### (*List[T]) UnmarshalJSON 

``` go
func (list *List[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### (*List[T]) Values 

``` go
func (list *List[T]) Values() []T
```

Values returns all elements in the list.
