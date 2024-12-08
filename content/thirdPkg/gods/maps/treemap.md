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

​	包 `treemap` 实现了一个由红黑树支持的映射（map）。

Elements are ordered by key in the map.

​	元素按键在映射中排序。

Structure is not thread safe.

​	该结构不是线程安全的。

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

​	Iterator 持有迭代器的状态。

#### (*Iterator[K, V]) Begin 

``` go
func (iterator *Iterator[K, V]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

​	Begin 将迭代器重置为其初始状态（即在第一个元素之前）。如果有元素，调用 Next() 来获取第一个元素。

#### (*Iterator[K, V]) End 

``` go
func (iterator *Iterator[K, V]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

​	End 将迭代器移动到最后一个元素之后（即最后一个元素之后的位置）。如果有元素，调用 Prev() 来获取最后一个元素。

#### (*Iterator[K, V]) First 

``` go
func (iterator *Iterator[K, V]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator

​	First 将迭代器移动到第一个元素并返回 `true`，如果容器中存在第一个元素。若 First() 返回 `true`，则可以通过 Key() 和 Value() 获取第一个元素的键和值。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) Key 

``` go
func (iterator *Iterator[K, V]) Key() K
```

Key returns the current element's key. Does not modify the state of the iterator.

​	Key 返回当前元素的键。此方法不会修改迭代器的状态。

#### (*Iterator[K, V]) Last 

``` go
func (iterator *Iterator[K, V]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	Last 将迭代器移动到最后一个元素并返回 `true`，如果容器中存在最后一个元素。若 Last() 返回 `true`，则可以通过 Key() 和 Value() 获取最后一个元素的键和值。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) Next 

``` go
func (iterator *Iterator[K, V]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's key and value can be retrieved by Key() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

​	Next 将迭代器移动到下一个元素并返回 `true`，如果容器中存在下一个元素。如果 Next() 第一次被调用，则它会将迭代器指向第一个元素（如果存在）。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) NextTo 

``` go
func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	NextTo 将迭代器移动到满足给定条件的下一个元素，并返回 `true`，如果容器中存在下一个元素。如果 NextTo() 返回 `true`，则可以通过 Key() 和 Value() 获取下一个元素的键和值。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) Prev 

``` go
func (iterator *Iterator[K, V]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	Prev 将迭代器移动到前一个元素并返回 `true`，如果容器中存在前一个元素。如果 Prev() 返回 `true`，则可以通过 Key() 和 Value() 获取前一个元素的键和值。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) PrevTo 

``` go
func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	PrevTo 将迭代器移动到满足给定条件的前一个元素，并返回 `true`，如果容器中存在前一个元素。如果 PrevTo() 返回 `true`，则可以通过 Key() 和 Value() 获取前一个元素的键和值。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) Value 

``` go
func (iterator *Iterator[K, V]) Value() V
```

Value returns the current element's value. Does not modify the state of the iterator.

​	Value 返回当前元素的值。此方法不会修改迭代器的状态。

### type Map 

``` go
type Map[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Map holds the elements in a red-black tree

​	Map 将元素保存在红黑树中。

#### func New 

``` go
func New[K cmp.Ordered, V any]() *Map[K, V]
```

New instantiates a tree map with the built-in comparator for K

​	New 使用内置比较器为 K 创建一个树映射。

#### func NewWith 

``` go
func NewWith[K comparable, V any](comparator utils.Comparator[K]) *Map[K, V]
```

NewWith instantiates a tree map with the custom comparator.

​	NewWith 使用自定义比较器创建一个树映射。

#### (*Map[K, V]) All 

``` go
func (m *Map[K, V]) All(f func(key K, value V) bool) bool
```

All passes each element of the container to the given function and returns true if the function returns true for all elements.

​	All 将容器中的每个元素传递给给定的函数，并且只有当函数对所有元素返回 `true` 时，才返回 `true`。

#### (*Map[K, V]) Any 

``` go
func (m *Map[K, V]) Any(f func(key K, value V) bool) bool
```

Any passes each element of the container to the given function and returns true if the function ever returns true for any element.

​	Any 将容器中的每个元素传递给给定的函数，并且如果该函数对任何元素返回 `true`，则返回 `true`。

#### (*Map[K, V]) Ceiling 

``` go
func (m *Map[K, V]) Ceiling(key K) (foundKey K, foundValue V, ok bool)
```

Ceiling finds the ceiling key-value pair for the input key. In case that no ceiling is found, then both returned values will be nil. It's generally enough to check the first value (key) for nil, which determines if ceiling was found.

​	Ceiling 查找输入键的天花板键值对。如果没有找到天花板，则返回的两个值将是 `nil`。通常，只需检查第一个值（键）是否为 `nil`，这可以确定是否找到了天花板。

Ceiling key is defined as the smallest key that is larger than or equal to the given key. A ceiling key may not be found, either because the map is empty, or because all keys in the map are smaller than the given key.

​	天花板键定义为大于或等于给定键的最小键。如果没有找到天花板键，可能是因为映射为空，或者映射中的所有键都小于给定的键。

Key should adhere to the comparator's type assertion, otherwise method panics.

​	键应该符合比较器的类型断言，否则该方法会引发恐慌。

#### (*Map[K, V]) Clear 

``` go
func (m *Map[K, V]) Clear()
```

Clear removes all elements from the map.

​	Clear 从映射中移除所有元素。

#### (*Map[K, V]) Each 

``` go
func (m *Map[K, V]) Each(f func(key K, value V))
```

Each calls the given function once for each element, passing that element's key and value.

​	Each 对容器中的每个元素调用给定的函数，传递该元素的键和值。

#### (*Map[K, V]) Empty 

``` go
func (m *Map[K, V]) Empty() bool
```

Empty returns true if map does not contain any elements

​	Empty 返回 `true`，如果映射不包含任何元素。

#### (*Map[K, V]) Find 

``` go
func (m *Map[K, V]) Find(f func(key K, value V) bool) (k K, v V)
```

Find passes each element of the container to the given function and returns the first (key,value) for which the function is true or nil,nil otherwise if no element matches the criteria.

​	Find 将容器中的每个元素传递给给定的函数，并返回第一个满足条件的键值对，如果没有元素符合条件，则返回 `nil, nil`。

#### (*Map[K, V]) Floor 

``` go
func (m *Map[K, V]) Floor(key K) (foundKey K, foundValue V, ok bool)
```

Floor finds the floor key-value pair for the input key. In case that no floor is found, then both returned values will be nil. It's generally enough to check the first value (key) for nil, which determines if floor was found.

​	Floor 查找输入键的下限键值对。如果没有找到下限，则返回的两个值将是 `nil`。通常，只需检查第一个值（键）是否为 `nil`，这可以确定是否找到了下限。

Floor key is defined as the largest key that is smaller than or equal to the given key. A floor key may not be found, either because the map is empty, or because all keys in the map are larger than the given key.

​	下限键定义为小于或等于给定键的最大键。如果没有找到下限键，可能是因为映射为空，或者映射中的所有键都大于给定的键。

Key should adhere to the comparator's type assertion, otherwise method panics.

​	键应该符合比较器的类型断言，否则该方法会引发恐慌。

#### (*Map[K, V]) FromJSON 

``` go
func (m *Map[K, V]) FromJSON(data []byte) error
```

FromJSON populates the map from the input JSON representation.

​	FromJSON 从输入的 JSON 表示填充映射。

#### (*Map[K, V]) Get 

``` go
func (m *Map[K, V]) Get(key K) (value V, found bool)
```

Get searches the element in the map by key and returns its value or nil if key is not found in tree. Second return parameter is true if key was found, otherwise false. Key should adhere to the comparator's type assertion, otherwise method panics.

​	Get 通过键在映射中查找元素并返回其值，如果没有找到该键，则返回 `nil`。第二个返回值为 `true` 表示找到了该键，否则为 `false`。键应符合比较器的类型断言，否则该方法会引发恐慌。

#### (*Map[K, V]) Iterator 

``` go
func (m *Map[K, V]) Iterator() *Iterator[K, V]
```

Iterator returns a stateful iterator whose elements are key/value pairs.

​	Iterator 返回一个有状态的迭代器，其元素是键值对。

#### (*Map[K, V]) Keys 

``` go
func (m *Map[K, V]) Keys() []K
```

Keys returns all keys in-order

​	Keys 返回所有按键排序的键。

#### (*Map[K, V]) Map 

``` go
func (m *Map[K, V]) Map(f func(key1 K, value1 V) (K, V)) *Map[K, V]
```

Map invokes the given function once for each element and returns a container containing the values returned by the given function as key/value pairs.

​	Map 对每个元素调用给定的函数，并返回一个包含函数返回的键值对的容器。

#### (*Map[K, V]) MarshalJSON 

``` go
func (m *Map[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

​	MarshalJSON 实现了 `json.Marshaler` 接口。

#### (*Map[K, V]) Max 

``` go
func (m *Map[K, V]) Max() (key K, value V, ok bool)
```

Max returns the maximum key and its value from the tree map. Returns 0-value, 0-value, false if map is empty.

​	Max 返回树映射中的最大键及其值。如果映射为空，则返回 0 值、0 值和 `false`。

#### (*Map[K, V]) Min 

``` go
func (m *Map[K, V]) Min() (key K, value V, ok bool)
```

Min returns the minimum key and its value from the tree map. Returns 0-value, 0-value, false if map is empty.

​	Min 返回树映射中的最小键及其值。如果映射为空，则返回 0 值、0 值和 `false`。

#### (*Map[K, V]) Put 

``` go
func (m *Map[K, V]) Put(key K, value V)
```

Put inserts key-value pair into the map. Key should adhere to the comparator's type assertion, otherwise method panics.

​	Put 将键值对插入映射中。键应符合比较器的类型断言，否则该方法会引发恐慌。

#### (*Map[K, V]) Remove 

``` go
func (m *Map[K, V]) Remove(key K)
```

Remove removes the element from the map by key. Key should adhere to the comparator's type assertion, otherwise method panics.

​	Remove 通过键从映射中移除元素。键应符合比较器的类型断言，否则该方法会引发恐慌。

#### (*Map[K, V]) Select 

``` go
func (m *Map[K, V]) Select(f func(key K, value V) bool) *Map[K, V]
```

Select returns a new container containing all elements for which the given function returns a true value.

​	Select 返回一个新的容器，包含所有函数返回 `true` 的元素。

#### (*Map[K, V]) Size 

``` go
func (m *Map[K, V]) Size() int
```

Size returns number of elements in the map.

​	Size 返回映射中元素的数量。

#### (*Map[K, V]) String 

``` go
func (m *Map[K, V]) String() string
```

String returns a string representation of container

​	String 返回容器的字符串表示。

#### (*Map[K, V]) ToJSON 

``` go
func (m *Map[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the map.

​	ToJSON 输出映射的 JSON 表示。

#### (*Map[K, V]) UnmarshalJSON 

``` go
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

​	UnmarshalJSON 实现了 `json.Unmarshaler` 接口。

#### (*Map[K, V]) Values 

``` go
func (m *Map[K, V]) Values() []V
```

Values returns all values in-order based on the key.

​	Values 返回所有按键排序的值。
