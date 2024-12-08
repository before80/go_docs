+++
title = "treemap"
date = 2024-12-07T11:04:27+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps/treemap](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps/treemap)
>
> 收录该文档时间： `2024-12-07T11:04:27+08:00`

## Overview 

Package treemap implements a map backed by red-black tree.

Elements are ordered by key in the map.

Structure is not thread safe.

Reference: http://en.wikipedia.org/wiki/Associative_array

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Iterator 

``` go
type Iterator[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Iterator holding the iterator's state

#### (*Iterator[K, V]) Begin 

``` go
func (iterator *Iterator[K, V]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

#### (*Iterator[K, V]) End 

``` go
func (iterator *Iterator[K, V]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

#### (*Iterator[K, V]) First 

``` go
func (iterator *Iterator[K, V]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator

#### (*Iterator[K, V]) Key 

``` go
func (iterator *Iterator[K, V]) Key() K
```

Key returns the current element's key. Does not modify the state of the iterator.

#### (*Iterator[K, V]) Last 

``` go
func (iterator *Iterator[K, V]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### (*Iterator[K, V]) Next 

``` go
func (iterator *Iterator[K, V]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's key and value can be retrieved by Key() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

#### (*Iterator[K, V]) NextTo 

``` go
func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### (*Iterator[K, V]) Prev 

``` go
func (iterator *Iterator[K, V]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### (*Iterator[K, V]) PrevTo 

``` go
func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### (*Iterator[K, V]) Value 

``` go
func (iterator *Iterator[K, V]) Value() V
```

Value returns the current element's value. Does not modify the state of the iterator.

### type Map 

``` go
type Map[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Map holds the elements in a red-black tree

#### func New 

``` go
func New[K cmp.Ordered, V any]() *Map[K, V]
```

New instantiates a tree map with the built-in comparator for K

#### func NewWith 

``` go
func NewWith[K comparable, V any](comparator utils.Comparator[K]) *Map[K, V]
```

NewWith instantiates a tree map with the custom comparator.

#### (*Map[K, V]) All 

``` go
func (m *Map[K, V]) All(f func(key K, value V) bool) bool
```

All passes each element of the container to the given function and returns true if the function returns true for all elements.

#### (*Map[K, V]) Any 

``` go
func (m *Map[K, V]) Any(f func(key K, value V) bool) bool
```

Any passes each element of the container to the given function and returns true if the function ever returns true for any element.

#### (*Map[K, V]) Ceiling 

``` go
func (m *Map[K, V]) Ceiling(key K) (foundKey K, foundValue V, ok bool)
```

Ceiling finds the ceiling key-value pair for the input key. In case that no ceiling is found, then both returned values will be nil. It's generally enough to check the first value (key) for nil, which determines if ceiling was found.

Ceiling key is defined as the smallest key that is larger than or equal to the given key. A ceiling key may not be found, either because the map is empty, or because all keys in the map are smaller than the given key.

Key should adhere to the comparator's type assertion, otherwise method panics.

#### (*Map[K, V]) Clear 

``` go
func (m *Map[K, V]) Clear()
```

Clear removes all elements from the map.

#### (*Map[K, V]) Each 

``` go
func (m *Map[K, V]) Each(f func(key K, value V))
```

Each calls the given function once for each element, passing that element's key and value.

#### (*Map[K, V]) Empty 

``` go
func (m *Map[K, V]) Empty() bool
```

Empty returns true if map does not contain any elements

#### (*Map[K, V]) Find 

``` go
func (m *Map[K, V]) Find(f func(key K, value V) bool) (k K, v V)
```

Find passes each element of the container to the given function and returns the first (key,value) for which the function is true or nil,nil otherwise if no element matches the criteria.

#### (*Map[K, V]) Floor 

``` go
func (m *Map[K, V]) Floor(key K) (foundKey K, foundValue V, ok bool)
```

Floor finds the floor key-value pair for the input key. In case that no floor is found, then both returned values will be nil. It's generally enough to check the first value (key) for nil, which determines if floor was found.

Floor key is defined as the largest key that is smaller than or equal to the given key. A floor key may not be found, either because the map is empty, or because all keys in the map are larger than the given key.

Key should adhere to the comparator's type assertion, otherwise method panics.

#### (*Map[K, V]) FromJSON 

``` go
func (m *Map[K, V]) FromJSON(data []byte) error
```

FromJSON populates the map from the input JSON representation.

#### (*Map[K, V]) Get 

``` go
func (m *Map[K, V]) Get(key K) (value V, found bool)
```

Get searches the element in the map by key and returns its value or nil if key is not found in tree. Second return parameter is true if key was found, otherwise false. Key should adhere to the comparator's type assertion, otherwise method panics.

#### (*Map[K, V]) Iterator 

``` go
func (m *Map[K, V]) Iterator() *Iterator[K, V]
```

Iterator returns a stateful iterator whose elements are key/value pairs.

#### (*Map[K, V]) Keys 

``` go
func (m *Map[K, V]) Keys() []K
```

Keys returns all keys in-order

#### (*Map[K, V]) Map 

``` go
func (m *Map[K, V]) Map(f func(key1 K, value1 V) (K, V)) *Map[K, V]
```

Map invokes the given function once for each element and returns a container containing the values returned by the given function as key/value pairs.

#### (*Map[K, V]) MarshalJSON 

``` go
func (m *Map[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### (*Map[K, V]) Max 

``` go
func (m *Map[K, V]) Max() (key K, value V, ok bool)
```

Max returns the maximum key and its value from the tree map. Returns 0-value, 0-value, false if map is empty.

#### (*Map[K, V]) Min 

``` go
func (m *Map[K, V]) Min() (key K, value V, ok bool)
```

Min returns the minimum key and its value from the tree map. Returns 0-value, 0-value, false if map is empty.

#### (*Map[K, V]) Put 

``` go
func (m *Map[K, V]) Put(key K, value V)
```

Put inserts key-value pair into the map. Key should adhere to the comparator's type assertion, otherwise method panics.

#### (*Map[K, V]) Remove 

``` go
func (m *Map[K, V]) Remove(key K)
```

Remove removes the element from the map by key. Key should adhere to the comparator's type assertion, otherwise method panics.

#### (*Map[K, V]) Select 

``` go
func (m *Map[K, V]) Select(f func(key K, value V) bool) *Map[K, V]
```

Select returns a new container containing all elements for which the given function returns a true value.

#### (*Map[K, V]) Size 

``` go
func (m *Map[K, V]) Size() int
```

Size returns number of elements in the map.

#### (*Map[K, V]) String 

``` go
func (m *Map[K, V]) String() string
```

String returns a string representation of container

#### (*Map[K, V]) ToJSON 

``` go
func (m *Map[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the map.

#### (*Map[K, V]) UnmarshalJSON 

``` go
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### (*Map[K, V]) Values 

``` go
func (m *Map[K, V]) Values() []V
```

Values returns all values in-order based on the key.
