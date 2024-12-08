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

​	`containers` 包为数据结构提供了核心接口和功能。

Container is the base interface for all data structures to implement.

​	**Container** 是所有数据结构必须实现的基础接口。

Iterators provide stateful iterators.

​	**Iterators** 提供有状态的迭代器。

Enumerable provides Ruby inspired (each, select, map, find, any?, etc.) container functions.

​	**Enumerable** 提供了受 Ruby 启发的容器函数（例如 each、select、map、find、any? 等）。

Serialization provides serializers (marshalers) and deserializers (unmarshalers).

​	**Serialization** 提供了序列化（marshalers）和反序列化（unmarshalers）功能。

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

​	`GetSortedValues` 根据传入的比较器返回容器中排序后的元素。不会影响容器内元素的原始顺序。

#### func GetSortedValuesFunc 

``` go
func GetSortedValuesFunc[T any](container Container[T], comparator utils.Comparator[T]) []T
```

GetSortedValuesFunc is the equivalent of GetSortedValues for containers of values that are not ordered.

​	`GetSortedValuesFunc` 是 `GetSortedValues` 的扩展，适用于值不可排序的容器。

## 类型 

### type Container 

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

### type EnumerableWithIndex 

``` go
type EnumerableWithIndex[T any] interface {
	// Each calls the given function once for each element, passing that element's index and value.
    // Each 调用传入的函数，对每个元素进行操作，传递该元素的索引和值。
	Each(func(index int, value T))

	// Any passes each element of the container to the given function and
	// returns true if the function ever returns true for any element.
    // Any 将容器的每个元素传递给传入的函数，
	// 如果函数对任何一个元素返回 true，则返回 true。
	Any(func(index int, value T) bool) bool

	// All passes each element of the container to the given function and
	// returns true if the function returns true for all elements.
    // All 将容器的每个元素传递给传入的函数，
	// 如果函数对所有元素都返回 true，则返回 true。
	All(func(index int, value T) bool) bool

	// Find passes each element of the container to the given function and returns
	// the first (index,value) for which the function is true or -1,nil otherwise
	// if no element matches the criteria.
    // Find 将容器的每个元素传递给传入的函数，返回第一个使函数返回 true 的 (索引, 值)，
	// 如果没有匹配的元素，则返回 -1, nil。
	Find(func(index int, value T) bool) (int, T)
}
```

EnumerableWithIndex provides functions for ordered containers whose values can be fetched by an index.

​	`EnumerableWithIndex` 提供了用于有序容器的函数，这些容器的值可以通过索引进行访问。

### type EnumerableWithKey 

``` go
type EnumerableWithKey[K, V any] interface {
	// Each calls the given function once for each element, passing that element's key and value.
    // Each 调用传入的函数，对每个元素进行操作，传递该元素的键和值。
	Each(func(key K, value V))

	// Any passes each element of the container to the given function and
	// returns true if the function ever returns true for any element.
    // Any 将容器的每个元素传递给传入的函数，
	// 如果函数对任何一个元素返回 true，则返回 true。
	Any(func(key K, value V) bool) bool

	// All passes each element of the container to the given function and
	// returns true if the function returns true for all elements.
    // All 将容器的每个元素传递给传入的函数，
	// 如果函数对所有元素都返回 true，则返回 true。
	All(func(key K, value V) bool) bool

	// Find passes each element of the container to the given function and returns
	// the first (key,value) for which the function is true or nil,nil otherwise if no element
	// matches the criteria.
    // Find 将容器的每个元素传递给传入的函数，返回第一个使函数返回 true 的 (键, 值)，
	// 如果没有匹配的元素，则返回 nil, nil。
	Find(func(key K, value V) bool) (K, V)
}
```

EnumerableWithKey provides functions for ordered containers whose values whose elements are key/value pairs.

​	`EnumerableWithKey` 提供了用于有序容器的函数，这些容器的元素是键值对。

### type IteratorWithIndex 

``` go
type IteratorWithIndex[T any] interface {
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
	// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
	// Modifies the state of the iterator.
    // Next 将迭代器移动到下一个元素，并返回容器中是否存在下一个元素。
	// 如果 Next() 返回 true，则可以通过 Index() 和 Value() 获取下一个元素的索引和值。
	// 如果首次调用 Next()，则会将迭代器指向第一个元素（如果存在）。
	// 修改迭代器的状态。
	Next() bool

	// Value returns the current element's value.
	// Does not modify the state of the iterator.
    // Value 返回当前元素的值。
	// 不会修改迭代器的状态。
	Value() T

	// Index returns the current element's index.
	// Does not modify the state of the iterator.
    // Index 返回当前元素的索引。
	// 不会修改迭代器的状态。
	Index() int

	// Begin resets the iterator to its initial state (one-before-first)
	// Call Next() to fetch the first element if any.
    // Begin 重置迭代器到其初始状态（第一个元素之前）。
	// 调用 Next() 以获取第一个元素（如果有）。
	Begin()

	// First moves the iterator to the first element and returns true if there was a first element in the container.
	// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
    // First 将迭代器移动到第一个元素，并返回容器中是否存在第一个元素。
	// 如果 First() 返回 true，则可以通过 Index() 和 Value() 获取第一个元素的索引和值。
	// 修改迭代器的状态。
	First() bool

	// NextTo moves the iterator to the next element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If NextTo() returns true, then next element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
    // NextTo 将迭代器从当前位置移动到满足传入函数条件的下一个元素，并返回容器中是否存在这样的下一个元素。
	// 如果 NextTo() 返回 true，则可以通过 Index() 和 Value() 获取该元素的索引和值。
	// 修改迭代器的状态。
	NextTo(func(index int, value T) bool) bool
}
```

IteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.

​	`IteratorWithIndex` 是一种用于有序容器的有状态迭代器，可以通过索引访问值。

### type IteratorWithKey 

``` go
type IteratorWithKey[K, V any] interface {
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
	// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
	// Modifies the state of the iterator.
    // Next 将迭代器移动到下一个元素，并返回容器中是否存在下一个元素。
	// 如果 Next() 返回 true，则可以通过 Key() 和 Value() 获取下一个元素的键和值。
	// 如果首次调用 Next()，则会将迭代器指向第一个元素（如果存在）。
	// 修改迭代器的状态。
	Next() bool

	// Value returns the current element's value.
	// Does not modify the state of the iterator.
    // Value 返回当前元素的值。
	// 不会修改迭代器的状态。
	Value() V

	// Key returns the current element's key.
	// Does not modify the state of the iterator.
    // Key 返回当前元素的键。
	// 不会修改迭代器的状态。
	Key() K

	// Begin resets the iterator to its initial state (one-before-first)
	// Call Next() to fetch the first element if any.
    // Begin 重置迭代器到其初始状态（第一个元素之前）。
	// 调用 Next() 以获取第一个元素（如果有）。
	Begin()

	// First moves the iterator to the first element and returns true if there was a first element in the container.
	// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
    // First 将迭代器移动到第一个元素，并返回容器中是否存在第一个元素。
	// 如果 First() 返回 true，则可以通过 Key() 和 Value() 获取第一个元素的键和值。
	// 修改迭代器的状态。
	First() bool

	// NextTo moves the iterator to the next element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
    // NextTo 将迭代器从当前位置移动到满足传入函数条件的下一个元素，并返回容器中是否存在这样的下一个元素。
	// 如果 NextTo() 返回 true，则可以通过 Key() 和 Value() 获取该元素的键和值。
	// 修改迭代器的状态。
	NextTo(func(key K, value V) bool) bool
}
```

IteratorWithKey is a stateful iterator for ordered containers whose elements are key value pairs.

​	`IteratorWithKey` 是一种用于键值对容器的有状态迭代器。

### type JSONDeserializer 

``` go
type JSONDeserializer interface {
	// FromJSON populates containers's elements from the input JSON representation.
    // FromJSON 从输入的 JSON 表示中填充容器的元素。
	FromJSON([]byte) error
	// UnmarshalJSON @implements json.Unmarshaler
    // UnmarshalJSON 实现了 json.Unmarshaler 接口。
	UnmarshalJSON([]byte) error
}
```

JSONDeserializer provides JSON deserialization

​	`JSONDeserializer` 提供 JSON 反序列化功能。

### type JSONSerializer 

``` go
type JSONSerializer interface {
	// ToJSON outputs the JSON representation of containers's elements.
    // ToJSON 输出容器的 JSON 表示。
	ToJSON() ([]byte, error)
	// MarshalJSON @implements json.Marshaler
    // MarshalJSON 实现了 json.Marshaler 接口。
	MarshalJSON() ([]byte, error)
}
```

JSONSerializer provides JSON serialization

​	`JSONSerializer` 提供 JSON 序列化功能。

### type ReverseIteratorWithIndex 

``` go
type ReverseIteratorWithIndex[T any] interface {
	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
    // Prev 将迭代器移动到上一个元素，并返回容器中是否存在上一个元素。
	// 如果 Prev() 返回 true，则可以通过 Index() 和 Value() 获取上一个元素的索引和值。
	// 修改迭代器的状态。
	Prev() bool

	// End moves the iterator past the last element (one-past-the-end).
	// Call Prev() to fetch the last element if any.
    // End 将迭代器移动到容器的末尾之后（超出位置）。
	// 调用 Prev() 以获取最后一个元素（如果有）。
	End()

	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
    // Last 将迭代器移动到最后一个元素，并返回容器中是否存在最后一个元素。
	// 如果 Last() 返回 true，则可以通过 Index() 和 Value() 获取最后一个元素的索引和值。
	// 修改迭代器的状态。
	Last() bool

	// PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If PrevTo() returns true, then next element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
    // PrevTo 将迭代器从当前位置移动到满足传入函数条件的上一个元素，并返回容器中是否存在这样的上一个元素。
	// 如果 PrevTo() 返回 true，则可以通过 Index() 和 Value() 获取该元素的索引和值。
	// 修改迭代器的状态。
	PrevTo(func(index int, value T) bool) bool

	IteratorWithIndex[T]
}
```

- [Prev() function to enable traversal in reverse](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/containers#hdr-Prev___function_to_enable_traversal_in_reverse-ReverseIteratorWithIndex)
  - 支持反向遍历的 Prev() 函数


ReverseIteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.

Essentially it is the same as IteratorWithIndex, but provides additional:

​	`ReverseIteratorWithIndex` 是一种用于有序容器的有状态反向迭代器，其功能与 `IteratorWithIndex` 类似，但提供了额外的以下功能：

#### Prev() function to enable traversal in reverse 

​	`Prev()` 支持反向遍历。

Last() function to move the iterator to the last element.

​	`Last()` 移动到最后一个元素。

End() function to move the iterator past the last element (one-past-the-end).

​	`End()` 移动到末尾之后的位置。

### type ReverseIteratorWithKey 

``` go
type ReverseIteratorWithKey[K, V any] interface {
	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
    // Prev 将迭代器移动到上一个元素，并返回容器中是否存在上一个元素。
	// 如果 Prev() 返回 true，则可以通过 Key() 和 Value() 获取上一个元素的键和值。
	// 修改迭代器的状态。
	Prev() bool

	// End moves the iterator past the last element (one-past-the-end).
	// Call Prev() to fetch the last element if any.
    // End 将迭代器移动到容器的末尾之后（超出位置）。
	// 调用 Prev() 以获取最后一个元素（如果有）。
	End()

	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
    // Last 将迭代器移动到最后一个元素，并返回容器中是否存在最后一个元素。
	// 如果 Last() 返回 true，则可以通过 Key() 和 Value() 获取最后一个元素的键和值。
	// 修改迭代器的状态。
	Last() bool

	// PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
    // PrevTo 将迭代器从当前位置移动到满足传入函数条件的上一个元素，并返回容器中是否存在这样的上一个元素。
	// 如果 PrevTo() 返回 true，则可以通过 Key() 和 Value() 获取该元素的键和值。
	// 修改迭代器的状态。
	PrevTo(func(key K, value V) bool) bool

	IteratorWithKey[K, V]
}
```

- [Prev() function to enable traversal in reverse](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/containers#hdr-Prev___function_to_enable_traversal_in_reverse-ReverseIteratorWithKey)
  - 支持反向遍历的 Prev() 函数


ReverseIteratorWithKey is a stateful iterator for ordered containers whose elements are key value pairs.	

Essentially it is the same as IteratorWithKey, but provides additional:

​	`ReverseIteratorWithKey` 是一种用于键值对容器的有状态反向迭代器，其功能与 `IteratorWithKey` 类似，但提供了额外的以下功能：

#### Prev() function to enable traversal in reverse 

​	`Prev()` 支持反向遍历。

Last() function to move the iterator to the last element.

​	`Last()` 移动到最后一个元素。
