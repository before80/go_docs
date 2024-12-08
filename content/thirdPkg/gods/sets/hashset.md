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

Structure is not thread safe.

References: [http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29](http://en.wikipedia.org/wiki/Set_(abstract_data_type))

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

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

Add adds the items (one or more) to the set.

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

#### (*Set[T]) Empty 

``` go
func (set *Set[T]) Empty() bool
```

Empty returns true if set does not contain any elements.

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

#### (*Set[T]) MarshalJSON 

``` go
func (set *Set[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### (*Set[T]) Remove 

``` go
func (set *Set[T]) Remove(items ...T)
```

Remove removes the items (one or more) from the set.

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
