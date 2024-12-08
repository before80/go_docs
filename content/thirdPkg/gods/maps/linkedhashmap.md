+++
title = "linkedhashmap"
date = 2024-12-07T11:03:41+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps/linkedhashmap](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps/linkedhashmap)
>
> 收录该文档时间： `2024-12-07T11:03:41+08:00`

## Overview 

Package linkedhashmap is a map that preserves insertion-order.

It is backed by a hash table to store values and doubly-linked list to store ordering.

Structure is not thread safe.

Reference: http://en.wikipedia.org/wiki/Associative_array

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

#### type Iterator 

``` go
type Iterator[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Iterator holding the iterator's state

#### func (*Iterator[K, V]) [Begin](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L59) 

``` go
func (iterator *Iterator[K, V]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

#### func (*Iterator[K, V]) [End](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L65) 

``` go
func (iterator *Iterator[K, V]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

#### func (*Iterator[K, V]) [First](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L72) 

``` go
func (iterator *Iterator[K, V]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator

#### func (*Iterator[K, V]) [Key](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L53) 

``` go
func (iterator *Iterator[K, V]) Key() K
```

Key returns the current element's key. Does not modify the state of the iterator.

#### func (*Iterator[K, V]) [Last](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L79) 

``` go
func (iterator *Iterator[K, V]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Next](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L33) 

``` go
func (iterator *Iterator[K, V]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's key and value can be retrieved by Key() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

#### func (*Iterator[K, V]) [NextTo](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L87) 

``` go
func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Prev](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L40) 

``` go
func (iterator *Iterator[K, V]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [PrevTo](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L101) 

``` go
func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Value](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L46) 

``` go
func (iterator *Iterator[K, V]) Value() V
```

Value returns the current element's value. Does not modify the state of the iterator.

#### type Map 

``` go
type Map[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Map holds the elements in a regular hash table, and uses doubly-linked list to store key ordering.

#### func New 

``` go
func New[K comparable, V any]() *Map[K, V]
```

New instantiates a linked-hash-map.

#### func (*Map[K, V]) [All](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/enumerable.go#L58) 

``` go
func (m *Map[K, V]) All(f func(key K, value V) bool) bool
```

All passes each element of the container to the given function and returns true if the function returns true for all elements.

#### func (*Map[K, V]) [Any](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/enumerable.go#L46) 

``` go
func (m *Map[K, V]) Any(f func(key K, value V) bool) bool
```

Any passes each element of the container to the given function and returns true if the function ever returns true for any element.

#### func (*Map[K, V]) [Clear](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/linkedhashmap.go#L94) 

``` go
func (m *Map[K, V]) Clear()
```

Clear removes all elements from the map.

#### func (*Map[K, V]) [Each](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/enumerable.go#L13) 

``` go
func (m *Map[K, V]) Each(f func(key K, value V))
```

Each calls the given function once for each element, passing that element's key and value.

#### func (*Map[K, V]) [Empty](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/linkedhashmap.go#L67) 

``` go
func (m *Map[K, V]) Empty() bool
```

Empty returns true if map does not contain any elements

#### func (*Map[K, V]) [Find](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/enumerable.go#L71) 

``` go
func (m *Map[K, V]) Find(f func(key K, value V) bool) (k K, v V)
```

Find passes each element of the container to the given function and returns the first (key,value) for which the function is true or nil,nil otherwise if no element matches the criteria.

#### func (*Map[K, V]) [FromJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/serialization.go#L72) 

``` go
func (m *Map[K, V]) FromJSON(data []byte) error
```

FromJSON populates map from the input JSON representation.

#### func (*Map[K, V]) [Get](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/linkedhashmap.go#L51) 

``` go
func (m *Map[K, V]) Get(key K) (value V, found bool)
```

Get searches the element in the map by key and returns its value or nil if key is not found in tree. Second return parameter is true if key was found, otherwise false. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Map[K, V]) [Iterator](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/iterator.go#L22) 

``` go
func (m *Map[K, V]) Iterator() *Iterator[K, V]
```

Iterator returns a stateful iterator whose elements are key/value pairs.

#### func (*Map[K, V]) [Keys](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/linkedhashmap.go#L77) 

``` go
func (m *Map[K, V]) Keys() []K
```

Keys returns all keys in-order

#### func (*Map[K, V]) [Map](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/enumerable.go#L22) 

``` go
func (m *Map[K, V]) Map(f func(key1 K, value1 V) (K, V)) *Map[K, V]
```

Map invokes the given function once for each element and returns a container containing the values returned by the given function as key/value pairs.

#### func (*Map[K, V]) [MarshalJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/serialization.go#L108) 

``` go
func (m *Map[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### func (*Map[K, V]) [Put](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/linkedhashmap.go#L41) 

``` go
func (m *Map[K, V]) Put(key K, value V)
```

Put inserts key-value pair into the map. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Map[K, V]) [Remove](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/linkedhashmap.go#L58) 

``` go
func (m *Map[K, V]) Remove(key K)
```

Remove removes the element from the map by key. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Map[K, V]) [Select](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/enumerable.go#L33) 

``` go
func (m *Map[K, V]) Select(f func(key K, value V) bool) *Map[K, V]
```

Select returns a new container containing all elements for which the given function returns a true value.

#### func (*Map[K, V]) [Size](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/linkedhashmap.go#L72) 

``` go
func (m *Map[K, V]) Size() int
```

Size returns number of elements in the map.

#### func (*Map[K, V]) [String](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/linkedhashmap.go#L100) 

``` go
func (m *Map[K, V]) String() string
```

String returns a string representation of container

#### func (*Map[K, V]) [ToJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/serialization.go#L21) 

``` go
func (m *Map[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of map.

#### func (*Map[K, V]) [UnmarshalJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/serialization.go#L103) 

``` go
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### func (*Map[K, V]) [Values](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/linkedhashmap/linkedhashmap.go#L82) 

``` go
func (m *Map[K, V]) Values() []V
```

Values returns all values in-order based on the key.
