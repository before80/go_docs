+++
title = "containers"
date = 2024-12-07T11:01:14+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/containers](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/containers)
>
> 收录该文档时间： `2024-12-07T11:01:14+08:00`

## Overview 

Package containers provides core interfaces and functions for data structures.

Container is the base interface for all data structures to implement.

Iterators provide stateful iterators.

Enumerable provides Ruby inspired (each, select, map, find, any?, etc.) container functions.

Serialization provides serializers (marshalers) and deserializers (unmarshalers).

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

#### func GetSortedValues 

``` go
func GetSortedValues[T cmp.Ordered](container Container[T]) []T
```

GetSortedValues returns sorted container's elements with respect to the passed comparator. Does not affect the ordering of elements within the container.

#### func GetSortedValuesFunc 

``` go
func GetSortedValuesFunc[T any](container Container[T], comparator utils.Comparator[T]) []T
```

GetSortedValuesFunc is the equivalent of GetSortedValues for containers of values that are not ordered.

## 类型 

#### type Container 

``` go
type Container[T any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}
```

Container is base interface that all data structures implement.

#### type EnumerableWithIndex 

``` go
type EnumerableWithIndex[T any] interface {
	// Each calls the given function once for each element, passing that element's index and value.
	Each(func(index int, value T))

	// Any passes each element of the container to the given function and
	// returns true if the function ever returns true for any element.
	Any(func(index int, value T) bool) bool

	// All passes each element of the container to the given function and
	// returns true if the function returns true for all elements.
	All(func(index int, value T) bool) bool

	// Find passes each element of the container to the given function and returns
	// the first (index,value) for which the function is true or -1,nil otherwise
	// if no element matches the criteria.
	Find(func(index int, value T) bool) (int, T)
}
```

EnumerableWithIndex provides functions for ordered containers whose values can be fetched by an index.

#### type EnumerableWithKey 

``` go
type EnumerableWithKey[K, V any] interface {
	// Each calls the given function once for each element, passing that element's key and value.
	Each(func(key K, value V))

	// Any passes each element of the container to the given function and
	// returns true if the function ever returns true for any element.
	Any(func(key K, value V) bool) bool

	// All passes each element of the container to the given function and
	// returns true if the function returns true for all elements.
	All(func(key K, value V) bool) bool

	// Find passes each element of the container to the given function and returns
	// the first (key,value) for which the function is true or nil,nil otherwise if no element
	// matches the criteria.
	Find(func(key K, value V) bool) (K, V)
}
```

EnumerableWithKey provides functions for ordered containers whose values whose elements are key/value pairs.

#### type IteratorWithIndex 

``` go
type IteratorWithIndex[T any] interface {
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
	// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
	// Modifies the state of the iterator.
	Next() bool

	// Value returns the current element's value.
	// Does not modify the state of the iterator.
	Value() T

	// Index returns the current element's index.
	// Does not modify the state of the iterator.
	Index() int

	// Begin resets the iterator to its initial state (one-before-first)
	// Call Next() to fetch the first element if any.
	Begin()

	// First moves the iterator to the first element and returns true if there was a first element in the container.
	// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	First() bool

	// NextTo moves the iterator to the next element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If NextTo() returns true, then next element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	NextTo(func(index int, value T) bool) bool
}
```

IteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.

#### type IteratorWithKey 

``` go
type IteratorWithKey[K, V any] interface {
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
	// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
	// Modifies the state of the iterator.
	Next() bool

	// Value returns the current element's value.
	// Does not modify the state of the iterator.
	Value() V

	// Key returns the current element's key.
	// Does not modify the state of the iterator.
	Key() K

	// Begin resets the iterator to its initial state (one-before-first)
	// Call Next() to fetch the first element if any.
	Begin()

	// First moves the iterator to the first element and returns true if there was a first element in the container.
	// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	First() bool

	// NextTo moves the iterator to the next element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	NextTo(func(key K, value V) bool) bool
}
```

IteratorWithKey is a stateful iterator for ordered containers whose elements are key value pairs.

#### type JSONDeserializer 

``` go
type JSONDeserializer interface {
	// FromJSON populates containers's elements from the input JSON representation.
	FromJSON([]byte) error
	// UnmarshalJSON @implements json.Unmarshaler
	UnmarshalJSON([]byte) error
}
```

JSONDeserializer provides JSON deserialization

#### type JSONSerializer 

``` go
type JSONSerializer interface {
	// ToJSON outputs the JSON representation of containers's elements.
	ToJSON() ([]byte, error)
	// MarshalJSON @implements json.Marshaler
	MarshalJSON() ([]byte, error)
}
```

JSONSerializer provides JSON serialization

#### type ReverseIteratorWithIndex 

``` go
type ReverseIteratorWithIndex[T any] interface {
	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	Prev() bool

	// End moves the iterator past the last element (one-past-the-end).
	// Call Prev() to fetch the last element if any.
	End()

	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	Last() bool

	// PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If PrevTo() returns true, then next element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	PrevTo(func(index int, value T) bool) bool

	IteratorWithIndex[T]
}
```

- [Prev() function to enable traversal in reverse](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/containers#hdr-Prev___function_to_enable_traversal_in_reverse-ReverseIteratorWithIndex)

ReverseIteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.

Essentially it is the same as IteratorWithIndex, but provides additional:

#### Prev() function to enable traversal in reverse 

Last() function to move the iterator to the last element.

End() function to move the iterator past the last element (one-past-the-end).

#### type ReverseIteratorWithKey 

``` go
type ReverseIteratorWithKey[K, V any] interface {
	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	Prev() bool

	// End moves the iterator past the last element (one-past-the-end).
	// Call Prev() to fetch the last element if any.
	End()

	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	Last() bool

	// PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	PrevTo(func(key K, value V) bool) bool

	IteratorWithKey[K, V]
}
```

- [Prev() function to enable traversal in reverse](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/containers#hdr-Prev___function_to_enable_traversal_in_reverse-ReverseIteratorWithKey)

ReverseIteratorWithKey is a stateful iterator for ordered containers whose elements are key value pairs.

Essentially it is the same as IteratorWithKey, but provides additional:

#### Prev() function to enable traversal in reverse 

Last() function to move the iterator to the last element.
