+++
title = "hashmap"
date = 2024-12-07T11:03:32+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps/hashmap](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps/hashmap)
>
> 收录该文档时间： `2024-12-07T11:03:32+08:00`

## Overview 

Package hashmap implements a map backed by a hash table.

Elements are unordered in the map.

Structure is not thread safe.

Reference: http://en.wikipedia.org/wiki/Associative_array



## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

#### type Map 

``` go
type Map[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Map holds the elements in go's native map

#### func New 

``` go
func New[K comparable, V any]() *Map[K, V]
```

New instantiates a hash map.

#### func (*Map[K, V]) [Clear](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/hashmap.go#L83) 

``` go
func (m *Map[K, V]) Clear()
```

Clear removes all elements from the map.

#### func (*Map[K, V]) [Empty](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/hashmap.go#L51) 

``` go
func (m *Map[K, V]) Empty() bool
```

Empty returns true if map does not contain any elements

#### func (*Map[K, V]) [FromJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/serialization.go#L23) 

``` go
func (m *Map[K, V]) FromJSON(data []byte) error
```

FromJSON populates the map from the input JSON representation.

#### func (*Map[K, V]) [Get](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/hashmap.go#L40) 

``` go
func (m *Map[K, V]) Get(key K) (value V, found bool)
```

Get searches the element in the map by key and returns its value or nil if key is not found in map. Second return parameter is true if key was found, otherwise false.

#### func (*Map[K, V]) [Keys](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/hashmap.go#L61) 

``` go
func (m *Map[K, V]) Keys() []K
```

Keys returns all keys (random order).

#### func (*Map[K, V]) [MarshalJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/serialization.go#L33) 

``` go
func (m *Map[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### func (*Map[K, V]) [Put](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/hashmap.go#L34) 

``` go
func (m *Map[K, V]) Put(key K, value V)
```

Put inserts element into the map.

#### func (*Map[K, V]) [Remove](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/hashmap.go#L46) 

``` go
func (m *Map[K, V]) Remove(key K)
```

Remove removes the element from the map by key.

#### func (*Map[K, V]) [Size](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/hashmap.go#L56) 

``` go
func (m *Map[K, V]) Size() int
```

Size returns number of elements in the map.

#### func (*Map[K, V]) [String](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/hashmap.go#L88) 

``` go
func (m *Map[K, V]) String() string
```

String returns a string representation of container

#### func (*Map[K, V]) [ToJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/serialization.go#L18) 

``` go
func (m *Map[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the map.

#### func (*Map[K, V]) [UnmarshalJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/serialization.go#L28) 

``` go
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### func (*Map[K, V]) [Values](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/maps/hashmap/hashmap.go#L72) 

``` go
func (m *Map[K, V]) Values() []V
```

Values returns all values (random order).
