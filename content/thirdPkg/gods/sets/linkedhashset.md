+++
title = "linkedhashset"
date = 2024-12-07T11:08:14+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/sets/linkedhashset](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/sets/linkedhashset)
>
> 收录该文档时间： `2024-12-07T11:08:14+08:00`

## Overview 

Package linkedhashset is a set that preserves insertion-order.

It is backed by a hash table to store values and doubly-linked list to store ordering.

Note that insertion-order is not affected if an element is re-inserted into the set.

Structure is not thread safe.

References: [http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29](http://en.wikipedia.org/wiki/Set_(abstract_data_type))

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

#### type Set 

``` go
type Set[T comparable] struct {
	// contains filtered or unexported fields
}
```

Set holds elements in go's native map

#### func New 

``` go
func New[T comparable](values ...T) *Set[T]
```

New instantiates a new empty set and adds the passed values, if any, to the set

#### (*Set[T]) Add 

``` go
func (set *Set[T]) Add(items ...T)
```

Add adds the items (one or more) to the set. Note that insertion-order is not affected if an element is re-inserted into the set.

#### (*Set[T]) All 

``` go
func (set *Set[T]) All(f func(index int, value T) bool) bool
```

All passes each element of the container to the given function and returns true if the function returns true for all elements.

#### (*Set[T]) Any 

``` go
func (set *Set[T]) Any(f func(index int, value T) bool) bool
```

Any passes each element of the container to the given function and returns true if the function ever returns true for any element.

#### (*Set[T]) Clear 

``` go
func (set *Set[T]) Clear()
```

Clear clears all values in the set.

#### (*Set[T]) Contains 

``` go
func (set *Set[T]) Contains(items ...T) bool
```

Contains check if items (one or more) are present in the set. All items have to be present in the set for the method to return true. Returns true if no arguments are passed at all, i.e. set is always superset of empty set.

#### (*Set[T]) Difference 

``` go
func (set *Set[T]) Difference(another *Set[T]) *Set[T]
```

Difference returns the difference between two sets. The new set consists of all elements that are in "set" but not in "another". Ref: https://proofwiki.org/wiki/Definition:Set_Difference

#### (*Set[T]) Each 

``` go
func (set *Set[T]) Each(f func(index int, value T))
```

Each calls the given function once for each element, passing that element's index and value.

#### (*Set[T]) Empty 

``` go
func (set *Set[T]) Empty() bool
```

Empty returns true if set does not contain any elements.

#### (*Set[T]) Find 

``` go
func (set *Set[T]) Find(f func(index int, value T) bool) (int, T)
```

Find passes each element of the container to the given function and returns the first (index,value) for which the function is true or -1,nil otherwise if no element matches the criteria.

#### (*Set[T]) FromJSON 

``` go
func (set *Set[T]) FromJSON(data []byte) error
```

FromJSON populates the set from the input JSON representation.

#### (*Set[T]) Intersection 

``` go
func (set *Set[T]) Intersection(another *Set[T]) *Set[T]
```

Intersection returns the intersection between two sets. The new set consists of all elements that are both in "set" and "another". Ref: https://en.wikipedia.org/wiki/Intersection_(set_theory)

#### (*Set[T]) Iterator 

``` go
func (set *Set[T]) Iterator() Iterator[T]
```

Iterator returns a stateful iterator whose values can be fetched by an index.

#### (*Set[T]) Map 

``` go
func (set *Set[T]) Map(f func(index int, value T) T) *Set[T]
```

Map invokes the given function once for each element and returns a container containing the values returned by the given function.

#### (*Set[T]) MarshalJSON 

``` go
func (set *Set[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### (*Set[T]) Remove 

``` go
func (set *Set[T]) Remove(items ...T)
```

Remove removes the items (one or more) from the set. Slow operation, worst-case O(n^2).

#### (*Set[T]) Select 

``` go
func (set *Set[T]) Select(f func(index int, value T) bool) *Set[T]
```

Select returns a new container containing all elements for which the given function returns a true value.

#### (*Set[T]) Size 

``` go
func (set *Set[T]) Size() int
```

Size returns number of elements within the set.

#### (*Set[T]) String 

``` go
func (set *Set[T]) String() string
```

String returns a string representation of container

#### (*Set[T]) ToJSON 

``` go
func (set *Set[T]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the set.

#### (*Set[T]) Union 

``` go
func (set *Set[T]) Union(another *Set[T]) *Set[T]
```

Union returns the union of two sets. The new set consists of all elements that are in "set" or "another" (possibly both). Ref: https://en.wikipedia.org/wiki/Union_(set_theory)

#### (*Set[T]) UnmarshalJSON 

``` go
func (set *Set[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### (*Set[T]) Values 

``` go
func (set *Set[T]) Values() []T
```

Values returns all items in the set.
