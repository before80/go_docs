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

​	包 `hashbidimap` 实现了一个由两个哈希映射支持的双向映射。

A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence. Thus the binary relation is functional in each direction: value can also act as a key to key. A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.

​	双向映射，或称哈希包，是一种关联数据结构，其中 (键, 值) 对形成一一对应关系。因此，这种二元关系在每个方向上都是功能性的：值也可以作为键来查找键。一个 (a, b) 对提供了 'a' 和 'b' 之间的唯一耦合关系，因此当 'a' 被用作键时，可以找到 'b'，而当 'b' 被用作键时，也可以找到 'a'。

Elements are unordered in the map.

​	映射中的元素是无序的。

Structure is not thread safe.

​	结构不是线程安全的。

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

​	`Map` 类型通过两个哈希映射来存储元素。

#### func New 

``` go
func New[K, V comparable]() *Map[K, V]
```

New instantiates a bidirectional map.

​	`New` 创建并返回一个新的双向映射。

#### (*Map[K, V]) Clear 

``` go
func (m *Map[K, V]) Clear()
```

Clear removes all elements from the map.

​	`Clear` 清空映射中的所有元素。

#### (*Map[K, V]) Empty 

``` go
func (m *Map[K, V]) Empty() bool
```

Empty returns true if map does not contain any elements

​	`Empty` 返回一个布尔值，表示映射是否为空。

#### (*Map[K, V]) FromJSON 

``` go
func (m *Map[K, V]) FromJSON(data []byte) error
```

FromJSON populates the map from the input JSON representation.

​	`FromJSON` 从输入的 JSON 表示填充映射。

#### (*Map[K, V]) Get 

``` go
func (m *Map[K, V]) Get(key K) (value V, found bool)
```

Get searches the element in the map by key and returns its value or nil if key is not found in map. Second return parameter is true if key was found, otherwise false.

​	`Get` 根据键从映射中查找元素并返回其值。如果键未找到，返回 `nil`，第二个返回值为 `false`，否则为 `true`。

#### (*Map[K, V]) GetKey 

``` go
func (m *Map[K, V]) GetKey(value V) (key K, found bool)
```

GetKey searches the element in the map by value and returns its key or nil if value is not found in map. Second return parameter is true if value was found, otherwise false.

​	`GetKey` 根据值从映射中查找元素并返回其键。如果值未找到，返回 `nil`，第二个返回值为 `false`，否则为 `true`。

#### (*Map[K, V]) Keys 

``` go
func (m *Map[K, V]) Keys() []K
```

Keys returns all keys (random order).

​	`Keys` 返回映射中的所有键，顺序是随机的。

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

Put inserts element into the map.

​	`Put` 将键值对插入到映射中。

#### (*Map[K, V]) Remove 

``` go
func (m *Map[K, V]) Remove(key K)
```

Remove removes the element from the map by key.

​	`Remove` 根据键从映射中删除元素。

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

​	`String` 返回映射的字符串表示。

#### (*Map[K, V]) ToJSON 

``` go
func (m *Map[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the map.

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

Values returns all values (random order).

​	`Values` 返回映射中的所有值，顺序是随机的。
