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

​	包 `linkedhashmap` 实现了一个保留插入顺序的映射。

It is backed by a hash table to store values and doubly-linked list to store ordering.

​	它由哈希表支持，用于存储值，并且使用双向链表来存储键的顺序。

Structure is not thread safe.

​	结构不是线程安全的。

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

​	`Iterator` 持有迭代器的状态。

#### (*Iterator[K, V]) Begin 

``` go
func (iterator *Iterator[K, V]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

​	`Begin` 重置迭代器到其初始状态（即第一个元素之前）。调用 `Next()` 以获取第一个元素（如果存在）。

#### (*Iterator[K, V]) End 

``` go
func (iterator *Iterator[K, V]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

​	`End` 将迭代器移动到最后一个元素之后的位置（即最后一个元素之后）。调用 `Prev()` 可以获取最后一个元素（如果存在）。

#### (*Iterator[K, V]) First 

``` go
func (iterator *Iterator[K, V]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator

​	`First` 将迭代器移到第一个元素，并返回 `true` 如果存在第一个元素。调用 `First()` 返回 `true` 后，可以通过 `Key()` 和 `Value()` 获取第一个元素的键和值。此操作会修改迭代器的状态。

#### (*Iterator[K, V]) Key 

``` go
func (iterator *Iterator[K, V]) Key() K
```

Key returns the current element's key. Does not modify the state of the iterator.

​	`Key` 返回当前元素的键。此操作不会修改迭代器的状态。

#### (*Iterator[K, V]) Last 

``` go
func (iterator *Iterator[K, V]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`Last` 将迭代器移到最后一个元素，并返回 `true` 如果存在最后一个元素。调用 `Last()` 返回 `true` 后，可以通过 `Key()` 和 `Value()` 获取最后一个元素的键和值。此操作会修改迭代器的状态。

#### (*Iterator[K, V]) Next 

``` go
func (iterator *Iterator[K, V]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's key and value can be retrieved by Key() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

​	`Next` 将迭代器移到下一个元素，并返回 `true` 如果存在下一个元素。调用 `Next()` 返回 `true` 后，可以通过 `Key()` 和 `Value()` 获取下一个元素的键和值。第一次调用 `Next()` 会将迭代器指向第一个元素（如果存在）。此操作会修改迭代器的状态。

#### (*Iterator[K, V]) NextTo 

``` go
func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`NextTo` 将迭代器移到下一个满足给定条件（由传递的函数 `f` 指定）的元素，并返回 `true` 如果存在下一个满足条件的元素。调用 `NextTo()` 返回 `true` 后，可以通过 `Key()` 和 `Value()` 获取该元素的键和值。此操作会修改迭代器的状态。

#### (*Iterator[K, V]) Prev 

``` go
func (iterator *Iterator[K, V]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`Prev` 将迭代器移到前一个元素，并返回 `true` 如果存在前一个元素。调用 `Prev()` 返回 `true` 后，可以通过 `Key()` 和 `Value()` 获取前一个元素的键和值。此操作会修改迭代器的状态。

#### (*Iterator[K, V]) PrevTo 

``` go
func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`PrevTo` 将迭代器移到前一个满足给定条件（由传递的函数 `f` 指定）的元素，并返回 `true` 如果存在前一个满足条件的元素。调用 `PrevTo()` 返回 `true` 后，可以通过 `Key()` 和 `Value()` 获取该元素的键和值。此操作会修改迭代器的状态。

#### (*Iterator[K, V]) Value 

``` go
func (iterator *Iterator[K, V]) Value() V
```

Value returns the current element's value. Does not modify the state of the iterator.

​	`Value` 返回当前元素的值。此操作不会修改迭代器的状态。

### type Map 

``` go
type Map[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Map holds the elements in a regular hash table, and uses doubly-linked list to store key ordering.

​	`Map` 使用哈希表存储元素，并利用双向链表存储键的顺序。

#### func New 

``` go
func New[K comparable, V any]() *Map[K, V]
```

New instantiates a linked-hash-map.

​	`New` 创建并返回一个新的链式哈希映射。

#### (*Map[K, V]) All 

``` go
func (m *Map[K, V]) All(f func(key K, value V) bool) bool
```

All passes each element of the container to the given function and returns true if the function returns true for all elements.

​	`All` 将容器中的每个元素传递给给定函数，并返回 `true` 如果函数对所有元素返回 `true`。

#### (*Map[K, V]) Any 

``` go
func (m *Map[K, V]) Any(f func(key K, value V) bool) bool
```

Any passes each element of the container to the given function and returns true if the function ever returns true for any element.

​	`Any` 将容器中的每个元素传递给给定函数，并返回 `true` 如果函数对任何元素返回 `true`。

#### (*Map[K, V]) Clear 

``` go
func (m *Map[K, V]) Clear()
```

Clear removes all elements from the map.

​	`Clear` 清空映射中的所有元素。

#### (*Map[K, V]) Each 

``` go
func (m *Map[K, V]) Each(f func(key K, value V))
```

Each calls the given function once for each element, passing that element's key and value.

​	`Each` 将给定函数应用于每个元素，传递该元素的键和值。

#### (*Map[K, V]) Empty 

``` go
func (m *Map[K, V]) Empty() bool
```

Empty returns true if map does not contain any elements

​	`Empty` 返回布尔值，表示映射是否为空。

#### (*Map[K, V]) Find 

``` go
func (m *Map[K, V]) Find(f func(key K, value V) bool) (k K, v V)
```

Find passes each element of the container to the given function and returns the first (key,value) for which the function is true or nil,nil otherwise if no element matches the criteria.

​	`Find` 将容器中的每个元素传递给给定函数，并返回第一个满足条件的元素的 `(key, value)`，如果没有元素匹配条件，则返回 `nil, nil`。

#### (*Map[K, V]) FromJSON 

``` go
func (m *Map[K, V]) FromJSON(data []byte) error
```

FromJSON populates map from the input JSON representation.

​	`FromJSON` 从输入的 JSON 表示填充映射。

#### (*Map[K, V]) Get 

``` go
func (m *Map[K, V]) Get(key K) (value V, found bool)
```

Get searches the element in the map by key and returns its value or nil if key is not found in tree. Second return parameter is true if key was found, otherwise false. Key should adhere to the comparator's type assertion, otherwise method panics.

​	`Get` 根据键从映射中查找元素，并返回其值（如果找到）或 `nil`（如果未找到）。第二个返回值为 `true` 如果键存在，否则为 `false`。

#### (*Map[K, V]) Iterator 

``` go
func (m *Map[K, V]) Iterator() *Iterator[K, V]
```

Iterator returns a stateful iterator whose elements are key/value pairs.

​	`Iterator` 返回一个有状态的迭代器，其中的元素是键/值对。

#### (*Map[K, V]) Keys 

``` go
func (m *Map[K, V]) Keys() []K
```

Keys returns all keys in-order

​	`Keys` 按顺序返回所有的键。

#### (*Map[K, V]) Map 

``` go
func (m *Map[K, V]) Map(f func(key1 K, value1 V) (K, V)) *Map[K, V]
```

Map invokes the given function once for each element and returns a container containing the values returned by the given function as key/value pairs.

​	`Map` 将给定函数应用于每个元素，并返回一个包含返回的键/值对的新容器。

#### (*Map[K, V]) MarshalJSON 

``` go
func (m *Map[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

​	`MarshalJSON` 实现了 `json.Marshaler` 接口。

#### (*Map[K, V]) Put 

``` go
func (m *Map[K, V]) Put(key K, value V)
```

Put inserts key-value pair into the map. Key should adhere to the comparator's type assertion, otherwise method panics.

​	`Put` 插入一个键/值对到映射中。键必须符合比较器的类型声明，否则方法会触发恐慌。

#### (*Map[K, V]) Remove 

``` go
func (m *Map[K, V]) Remove(key K)
```

Remove removes the element from the map by key. Key should adhere to the comparator's type assertion, otherwise method panics.

​	`Remove` 根据键从映射中删除元素。键必须符合比较器的类型声明，否则方法会触发恐慌。

#### (*Map[K, V]) Select 

``` go
func (m *Map[K, V]) Select(f func(key K, value V) bool) *Map[K, V]
```

Select returns a new container containing all elements for which the given function returns a true value.

​	`Select` 返回一个新容器，包含所有符合给定函数条件的元素。

#### (*Map[K, V]) Size 

``` go
func (m *Map[K, V]) Size() int
```

Size returns number of elements in the map.

​	`Size` 返回映射中元素的数量。

#### (*Map[K, V]) String 

``` go
func (m *Map[K, V]) String() string
```

String returns a string representation of container

​	`String` 返回容器的字符串表示。

#### (*Map[K, V]) ToJSON 

``` go
func (m *Map[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of map.

​	`ToJSON` 输出映射的 JSON 表示。

#### (*Map[K, V]) UnmarshalJSON 

``` go
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

​	`UnmarshalJSON` 实现了 `json.Unmarshaler` 接口。

#### (*Map[K, V]) Values 

``` go
func (m *Map[K, V]) Values() []V
```

Values returns all values in-order based on the key.

​	`Values` 返回所有按键顺序排列的值。
