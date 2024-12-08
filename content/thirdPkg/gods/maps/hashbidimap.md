+++
title = "hashbidimap"
date = 2024-12-07T11:03:19+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps/hashbidimap](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps/hashbidimap)
>
> 收录该文档时间： `2024-12-07T11:03:19+08:00`

## Overview 

Package hashbidimap implements a bidirectional map backed by two hashmaps.

A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence. Thus the binary relation is functional in each direction: value can also act as a key to key. A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.

Elements are unordered in the map.

Structure is not thread safe.

Reference: https://en.wikipedia.org/wiki/Bidirectional_map

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Map 

``` go
type Map[K, V comparable] struct {
	// contains filtered or unexported fields
}
```

Map holds the elements in two hashmaps.

#### func New 

``` go
func New[K, V comparable]() *Map[K, V]
```

New instantiates a bidirectional map.

#### (*Map[K, V]) Clear 

``` go
func (m *Map[K, V]) Clear()
```

Clear removes all elements from the map.

#### (*Map[K, V]) Empty 

``` go
func (m *Map[K, V]) Empty() bool
```

Empty returns true if map does not contain any elements

#### (*Map[K, V]) FromJSON 

``` go
func (m *Map[K, V]) FromJSON(data []byte) error
```

FromJSON populates the map from the input JSON representation.

#### (*Map[K, V]) Get 

``` go
func (m *Map[K, V]) Get(key K) (value V, found bool)
```

Get searches the element in the map by key and returns its value or nil if key is not found in map. Second return parameter is true if key was found, otherwise false.

#### (*Map[K, V]) GetKey 

``` go
func (m *Map[K, V]) GetKey(value V) (key K, found bool)
```

GetKey searches the element in the map by value and returns its key or nil if value is not found in map. Second return parameter is true if value was found, otherwise false.

#### (*Map[K, V]) Keys 

``` go
func (m *Map[K, V]) Keys() []K
```

Keys returns all keys (random order).

#### (*Map[K, V]) MarshalJSON 

``` go
func (m *Map[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### (*Map[K, V]) Put 

``` go
func (m *Map[K, V]) Put(key K, value V)
```

Put inserts element into the map.

#### (*Map[K, V]) Remove 

``` go
func (m *Map[K, V]) Remove(key K)
```

Remove removes the element from the map by key.

#### (*Map[K, V]) Size 

``` go
func (m *Map[K, V]) Size() int
```

Size returns number of elements in the map.

#### (*Map[K, V]) String 

``` go
func (m *Map[K, V]) String() string
```

String returns a string representation of container

#### (*Map[K, V]) ToJSON 

``` go
func (m *Map[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the map.

#### (*Map[K, V]) UnmarshalJSON 

``` go
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### (*Map[K, V]) Values 

``` go
func (m *Map[K, V]) Values() []V
```

Values returns all values (random order).
