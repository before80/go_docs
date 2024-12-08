+++
title = "hashset"
date = 2024-12-07T11:08:05+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/sets/hashset](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/sets/hashset)
>
> 收录该文档时间： `2024-12-07T11:08:05+08:00`

## Overview 

Package hashset implements a set backed by a hash table.

​	`hashset` 包实现了基于哈希表的集合（Set）。

Structure is not thread safe.

​	该结构不是线程安全的。

References: [http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29](http://en.wikipedia.org/wiki/Set_(abstract_data_type))

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Set 

``` go
type Set[T comparable] struct {
	// contains filtered or unexported fields
}
```

Set holds elements in go's native map

​	`Set` 使用 Go 的原生 `map` 存储元素。

#### func New 

``` go
func New[T comparable](values ...T) *Set[T]
```

New instantiates a new empty set and adds the passed values, if any, to the set

​	**New** 实例化一个新的空集合，并添加传入的值（如果有）。

#### (*Set[T]) Add 

``` go
func (set *Set[T]) Add(items ...T)
```

Add adds the items (one or more) to the set.

​	**Add** 将一个或多个元素添加到集合中。

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

Contains check if items (one or more) are present in the set. All items have to be present in the set for the method to return true. Returns true if no arguments are passed at all, i.e. set is always superset of empty set.

​	**Contains** 检查一个或多个元素是否存在于集合中。所有指定的元素都必须存在才能返回 `true`。如果没有传入任何参数，则始终返回 `true`，即集合总是空集的超集。

#### (*Set[T]) Difference 

``` go
func (set *Set[T]) Difference(another *Set[T]) *Set[T]
```

Difference returns the difference between two sets. The new set consists of all elements that are in "set" but not in "another". 

​	**Difference** 返回两个集合的差集。新集合由 "set" 中存在但 "another" 中不存在的所有元素组成。

Ref: [https://proofwiki.org/wiki/Definition:Set_Difference](https://proofwiki.org/wiki/Definition:Set_Difference)

#### (*Set[T]) Empty 

``` go
func (set *Set[T]) Empty() bool
```

Empty returns true if set does not contain any elements.

​	**Empty** 如果集合不包含任何元素，返回 `true`。

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

Intersection returns the intersection between two sets. The new set consists of all elements that are both in "set" and "another". 

​	**Intersection** 返回两个集合的交集。新集合由同时存在于 "set" 和 "another" 中的所有元素组成。

Ref: [https://en.wikipedia.org/wiki/Intersection_(set_theory)](https://en.wikipedia.org/wiki/Intersection_(set_theory))

#### (*Set[T]) MarshalJSON 

``` go
func (set *Set[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

​	**MarshalJSON** @实现接口 `json.Marshaler`

#### (*Set[T]) Remove 

``` go
func (set *Set[T]) Remove(items ...T)
```

Remove removes the items (one or more) from the set.

​	**Remove** 从集合中移除一个或多个元素。

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

​	**String** 返回集合的字符串表示。

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

Union returns the union of two sets. The new set consists of all elements that are in "set" or "another" (possibly both). 

​	**Union** 返回两个集合的并集。新集合由 "set" 或 "another"（可能两者都有）的所有元素组成。

Ref: https://en.wikipedia.org/wiki/Union_(set_theory)

#### (*Set[T]) UnmarshalJSON 

``` go
func (set *Set[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

​	**UnmarshalJSON** @实现接口 `json.Unmarshaler`

#### (*Set[T]) Values 

``` go
func (set *Set[T]) Values() []T
```

Values returns all items in the set.

​	**Values** 返回集合中的所有元素。
