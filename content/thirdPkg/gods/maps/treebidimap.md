+++
title = "treebidimap"
date = 2024-12-07T11:04:17+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps/treebidimap](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps/treebidimap)
>
> 收录该文档时间： `2024-12-07T11:04:17+08:00`

## Overview 

Package treebidimap implements a bidirectional map backed by two red-black tree.

This structure guarantees that the map will be in both ascending key and value order.

Other than key and value ordering, the goal with this structure is to avoid duplication of elements, which can be significant if contained elements are large.

A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence. Thus the binary relation is functional in each direction: value can also act as a key to key. A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.

Structure is not thread safe.

Reference: https://en.wikipedia.org/wiki/Bidirectional_map



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

#### func (*Iterator[K, V]) [Begin](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L54) 

``` go
func (iterator *Iterator[K, V]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

#### func (*Iterator[K, V]) [End](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L60) 

``` go
func (iterator *Iterator[K, V]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

#### func (*Iterator[K, V]) [First](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L67) 

``` go
func (iterator *Iterator[K, V]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator

#### func (*Iterator[K, V]) [Key](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L48) 

``` go
func (iterator *Iterator[K, V]) Key() K
```

Key returns the current element's key. Does not modify the state of the iterator.

#### func (*Iterator[K, V]) [Last](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L74) 

``` go
func (iterator *Iterator[K, V]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Next](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L29) 

``` go
func (iterator *Iterator[K, V]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's key and value can be retrieved by Key() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

#### func (*Iterator[K, V]) [NextTo](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L82) 

``` go
func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Prev](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L36) 

``` go
func (iterator *Iterator[K, V]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [PrevTo](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L96) 

``` go
func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Value](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L42) 

``` go
func (iterator *Iterator[K, V]) Value() V
```

Value returns the current element's value. Does not modify the state of the iterator.

#### type Map 

``` go
type Map[K, V comparable] struct {
	// contains filtered or unexported fields
}
```

Map holds the elements in two red-black trees.

#### func New 

``` go
func New[K, V cmp.Ordered]() *Map[K, V]
```

New instantiates a bidirectional map.

#### func NewWith 

``` go
func NewWith[K, V comparable](keyComparator utils.Comparator[K], valueComparator utils.Comparator[V]) *Map[K, V]
```

NewWith instantiates a bidirectional map.

#### func (*Map[K, V]) [All](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/enumerable.go#L58) 

``` go
func (m *Map[K, V]) All(f func(key K, value V) bool) bool
```

All passes each element of the container to the given function and returns true if the function returns true for all elements.

#### func (*Map[K, V]) [Any](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/enumerable.go#L46) 

``` go
func (m *Map[K, V]) Any(f func(key K, value V) bool) bool
```

Any passes each element of the container to the given function and returns true if the function ever returns true for any element.

#### func (*Map[K, V]) [Clear](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/treebidimap.go#L109) 

``` go
func (m *Map[K, V]) Clear()
```

Clear removes all elements from the map.

#### func (*Map[K, V]) [Each](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/enumerable.go#L13) 

``` go
func (m *Map[K, V]) Each(f func(key K, value V))
```

Each calls the given function once for each element, passing that element's key and value.

#### func (*Map[K, V]) [Empty](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/treebidimap.go#L89) 

``` go
func (m *Map[K, V]) Empty() bool
```

Empty returns true if map does not contain any elements

#### func (*Map[K, V]) [Find](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/enumerable.go#L71) 

``` go
func (m *Map[K, V]) Find(f func(key K, value V) bool) (k K, v V)
```

Find passes each element of the container to the given function and returns the first (key,value) for which the function is true or nil,nil otherwise if no element matches the criteria.

#### func (*Map[K, V]) [FromJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/serialization.go#L23) 

``` go
func (m *Map[K, V]) FromJSON(data []byte) error
```

FromJSON populates the map from the input JSON representation.

#### func (*Map[K, V]) [Get](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/treebidimap.go#L70) 

``` go
func (m *Map[K, V]) Get(key K) (value V, found bool)
```

Get searches the element in the map by key and returns its value or nil if key is not found in map. Second return parameter is true if key was found, otherwise false.

#### func (*Map[K, V]) [GetKey](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/treebidimap.go#L76) 

``` go
func (m *Map[K, V]) GetKey(value V) (key K, found bool)
```

GetKey searches the element in the map by value and returns its key or nil if value is not found in map. Second return parameter is true if value was found, otherwise false.

#### func (*Map[K, V]) [Iterator](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/iterator.go#L21) 

``` go
func (m *Map[K, V]) Iterator() *Iterator[K, V]
```

Iterator returns a stateful iterator whose elements are key/value pairs.

#### func (*Map[K, V]) [Keys](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/treebidimap.go#L99) 

``` go
func (m *Map[K, V]) Keys() []K
```

Keys returns all keys (ordered).

#### func (*Map[K, V]) [Map](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/enumerable.go#L22) 

``` go
func (m *Map[K, V]) Map(f func(key1 K, value1 V) (K, V)) *Map[K, V]
```

Map invokes the given function once for each element and returns a container containing the values returned by the given function as key/value pairs.

#### func (*Map[K, V]) [MarshalJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/serialization.go#L44) 

``` go
func (m *Map[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### func (*Map[K, V]) [Put](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/treebidimap.go#L57) 

``` go
func (m *Map[K, V]) Put(key K, value V)
```

Put inserts element into the map.

#### func (*Map[K, V]) [Remove](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/treebidimap.go#L81) 

``` go
func (m *Map[K, V]) Remove(key K)
```

Remove removes the element from the map by key.

#### func (*Map[K, V]) [Select](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/enumerable.go#L33) 

``` go
func (m *Map[K, V]) Select(f func(key K, value V) bool) *Map[K, V]
```

Select returns a new container containing all elements for which the given function returns a true value.

#### func (*Map[K, V]) [Size](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/treebidimap.go#L94) 

``` go
func (m *Map[K, V]) Size() int
```

Size returns number of elements in the map.

#### func (*Map[K, V]) [String](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/treebidimap.go#L115) 

``` go
func (m *Map[K, V]) String() string
```

String returns a string representation of container

#### func (*Map[K, V]) [ToJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/serialization.go#L18) 

``` go
func (m *Map[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the map.

#### func (*Map[K, V]) [UnmarshalJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/serialization.go#L39) 

``` go
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### func (*Map[K, V]) [Values](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/treebidimap/treebidimap.go#L104) 

``` go
func (m *Map[K, V]) Values() []V
```

Values returns all values (ordered).
