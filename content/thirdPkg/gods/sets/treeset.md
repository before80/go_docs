+++
title = "treeset"
date = 2024-12-07T11:08:25+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/sets/treeset](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/sets/treeset)
>
> 收录该文档时间： `2024-12-07T11:08:25+08:00`

## Overview 

Package treeset implements a tree backed by a red-black tree.

​	`treeset` 包实现了一个基于红黑树的集合（Set）。

Structure is not thread safe.

​	该结构不是线程安全的。

Reference: [http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29](http://en.wikipedia.org/wiki/Set_(abstract_data_type))

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

​	一个状态化迭代器，支持按索引获取值。

#### (*Iterator[T]) Begin 

``` go
func (iterator *Iterator[T]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

​	**Begin** 将迭代器重置到初始状态（第一个元素之前）。调用 `Next()` 获取第一个元素（如果存在）。

#### (*Iterator[T]) End 

``` go
func (iterator *Iterator[T]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

​	**End** 将迭代器移动到最后一个元素之后（超出范围）。调用 `Prev()` 获取最后一个元素（如果存在）。

#### (*Iterator[T]) First 

``` go
func (iterator *Iterator[T]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	**First** 将迭代器移动到第一个元素并返回 `true`，如果容器中存在第一个元素。若返回 `true`，则可通过 `Index()` 和 `Value()` 获取该元素的索引和值。修改迭代器的状态。

#### (*Iterator[T]) Index 

``` go
func (iterator *Iterator[T]) Index() int
```

Index returns the current element's index. Does not modify the state of the iterator.

​	**Index** 返回当前元素的索引，不修改迭代器的状态。

#### (*Iterator[T]) Last 

``` go
func (iterator *Iterator[T]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	**Last** 将迭代器移动到最后一个元素并返回 `true`，如果容器中存在最后一个元素。若返回 `true`，则可通过 `Index()` 和 `Value()` 获取该元素的索引和值。修改迭代器的状态。

#### (*Iterator[T]) Next 

``` go
func (iterator *Iterator[T]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's index and value can be retrieved by Index() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

​	**Next** 将迭代器移动到下一个元素并返回 `true`，如果容器中存在下一个元素。若返回 `true`，则可以通过 `Index()` 和 `Value()` 获取下一个元素的索引和值。首次调用 `Next()` 时，如果存在元素，迭代器将指向第一个元素。此方法会修改迭代器的状态。

#### (*Iterator[T]) NextTo 

``` go
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	**NextTo** 将迭代器移动到满足给定条件的下一个元素，并返回 `true`，如果容器中存在这样的元素。若返回 `true`，则可以通过 `Index()` 和 `Value()` 获取该元素的索引和值。此方法会修改迭代器的状态。

#### (*Iterator[T]) Prev 

``` go
func (iterator *Iterator[T]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	**Prev** 将迭代器移动到上一个元素并返回 `true`，如果容器中存在上一个元素。若返回 `true`，则可以通过 `Index()` 和 `Value()` 获取上一个元素的索引和值。此方法会修改迭代器的状态。

#### (*Iterator[T]) PrevTo 

``` go
func (iterator *Iterator[T]) PrevTo(f func(index int, value T) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	**PrevTo** 将迭代器移动到满足给定条件的上一个元素，并返回 `true`，如果容器中存在这样的元素。若返回 `true`，则可以通过 `Index()` 和 `Value()` 获取该元素的索引和值。此方法会修改迭代器的状态。

#### (*Iterator[T]) Value 

``` go
func (iterator *Iterator[T]) Value() T
```

Value returns the current element's value. Does not modify the state of the iterator.

​	**Value** 返回当前元素的值，不修改迭代器的状态。

### type Set 

``` go
type Set[T comparable] struct {
	// contains filtered or unexported fields
}
```

Set holds elements in a red-black tree

​	`Set` 使用红黑树存储元素。

#### func New 

``` go
func New[T cmp.Ordered](values ...T) *Set[T]
```

#### func NewWith 

``` go
func NewWith[T comparable](comparator utils.Comparator[T], values ...T) *Set[T]
```

NewWith instantiates a new empty set with the custom comparator.

​	**NewWith** 使用自定义比较器实例化一个新的空集合。

#### (*Set[T]) Add 

``` go
func (set *Set[T]) Add(items ...T)
```

Add adds the items (one or more) to the set.

​	**Add** 将一个或多个元素添加到集合中。

#### (*Set[T]) All 

``` go
func (set *Set[T]) All(f func(index int, value T) bool) bool
```

All passes each element of the container to the given function and returns true if the function returns true for all elements.

​	**All** 遍历集合中的每个元素，将其传递给给定函数。如果函数对所有元素返回 `true`，则返回 `true`。

#### (*Set[T]) Any 

``` go
func (set *Set[T]) Any(f func(index int, value T) bool) bool
```

Any passes each element of the container to the given function and returns true if the function ever returns true for any element.

​	**Any** 遍历集合中的每个元素，将其传递给给定函数。如果函数对任意元素返回 `true`，则返回 `true`。

#### (*Set[T]) Clear 

``` go
func (set *Set[T]) Clear()
```

Clear clears all values in the set.

​	**Clear** 清空集合中的所有元素。

#### (*Set[T]) Contains 

``` go
func (set *Set[T]) Contains(items ...T) bool
```

Contains checks weather items (one or more) are present in the set. All items have to be present in the set for the method to return true. Returns true if no arguments are passed at all, i.e. set is always superset of empty set.

​	**Contains** 检查一个或多个元素是否存在于集合中。方法仅在所有指定的元素都存在于集合中时返回 `true`。若未传入任何参数，则返回 `true`，即集合始终是空集合的超集。

#### (*Set[T]) Difference 

``` go
func (set *Set[T]) Difference(another *Set[T]) *Set[T]
```

Difference returns the difference between two sets. The two sets should have the same comparators, otherwise the result is empty set. The new set consists of all elements that are in "set" but not in "another". 

​	**Difference** 返回两个集合的差集。两个集合应具有相同的比较器，否则结果为空集。新集合包含所有在当前集合中但不在另一个集合中的元素。

Ref: https://proofwiki.org/wiki/Definition:Set_Difference

#### (*Set[T]) Each 

``` go
func (set *Set[T]) Each(f func(index int, value T))
```

Each calls the given function once for each element, passing that element's index and value.

​	**Each** 遍历集合中的每个元素，并将其索引和值传递给指定的函数。

#### (*Set[T]) Empty 

``` go
func (set *Set[T]) Empty() bool
```

Empty returns true if set does not contain any elements.

​	**Empty** 如果集合不包含任何元素，则返回 `true`。

#### (*Set[T]) Find 

``` go
func (set *Set[T]) Find(f func(index int, value T) bool) (int, T)
```

Find passes each element of the container to the given function and returns the first (index,value) for which the function is true or -1,nil otherwise if no element matches the criteria.

​	**Find** 遍历集合中的每个元素，将其传递给指定函数，并返回第一个使函数返回 `true` 的元素的索引和值。如果没有匹配的元素，则返回 `-1, nil`。

#### (*Set[T]) FromJSON 

``` go
func (set *Set[T]) FromJSON(data []byte) error
```

FromJSON populates the set from the input JSON representation.

​	**FromJSON** 根据输入的 JSON 表示填充集合。

#### (*Set[T]) Intersection 

``` go
func (set *Set[T]) Intersection(another *Set[T]) *Set[T]
```

Intersection returns the intersection between two sets. The new set consists of all elements that are both in "set" and "another". The two sets should have the same comparators, otherwise the result is empty set. 

​	**Intersection** 返回两个集合的交集。新集合包含当前集合和另一个集合中共同存在的所有元素。两个集合应具有相同的比较器，否则结果为空集。

Ref: https://en.wikipedia.org/wiki/Intersection_(set_theory)

#### (*Set[T]) Iterator 

``` go
func (set *Set[T]) Iterator() Iterator[T]
```

Iterator holding the iterator's state

​	**Iterator** 返回一个状态化的迭代器，用于遍历集合中的元素。

#### (*Set[T]) Map 

``` go
func (set *Set[T]) Map(f func(index int, value T) T) *Set[T]
```

Map invokes the given function once for each element and returns a container containing the values returned by the given function.

​	**Map** 遍历集合中的每个元素，调用指定的函数，并返回一个新集合，包含函数返回的所有值。

#### (*Set[T]) MarshalJSON 

``` go
func (set *Set[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

​	**MarshalJSON** @实现接口 `json.Marshaler`。

#### (*Set[T]) Remove 

``` go
func (set *Set[T]) Remove(items ...T)
```

Remove removes the items (one or more) from the set.

​	**Remove** 从集合中移除一个或多个指定的元素。

#### (*Set[T]) Select 

``` go
func (set *Set[T]) Select(f func(index int, value T) bool) *Set[T]
```

Select returns a new container containing all elements for which the given function returns a true value.

​	**Select** 返回一个新集合，包含所有使指定函数返回 `true` 的元素。

#### (*Set[T]) Size 

``` go
func (set *Set[T]) Size() int
```

Size returns number of elements within the set.

​	**Size** 返回集合中的元素数量。

#### (*Set[T]) String 

``` go
func (set *Set[T]) String() string
```

String returns a string representation of container

​	**String** 返回集合的字符串表示形式。

#### (*Set[T]) ToJSON 

``` go
func (set *Set[T]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the set.

​	**ToJSON** 输出集合的 JSON 表示。

#### (*Set[T]) Union 

``` go
func (set *Set[T]) Union(another *Set[T]) *Set[T]
```

Union returns the union of two sets. The new set consists of all elements that are in "set" or "another" (possibly both). The two sets should have the same comparators, otherwise the result is empty set. 

​	**Union** 返回两个集合的并集。新集合包含当前集合和另一个集合中的所有元素（可能同时存在于两个集合中）。两个集合应具有相同的比较器，否则结果为空集。

Ref: https://en.wikipedia.org/wiki/Union_(set_theory)

#### (*Set[T]) UnmarshalJSON 

``` go
func (set *Set[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

​	**UnmarshalJSON** @实现接口 `json.Unmarshaler`。

#### (*Set[T]) Values 

``` go
func (set *Set[T]) Values() []T
```

Values returns all items in the set.

​	**Values** 返回集合中的所有元素。
