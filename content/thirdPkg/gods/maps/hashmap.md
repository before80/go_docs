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

​	包 `hashmap` 实现了一个由哈希表支持的映射。

Elements are unordered in the map.

​	映射中的元素是无序的。

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

### type Map 

``` go
type Map[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Map holds the elements in go's native map

​	`Map` 类型通过 Go 的原生映射来存储元素。

#### func New 

``` go
func New[K comparable, V any]() *Map[K, V]
```

New instantiates a hash map.

​	`New` 创建并返回一个新的哈希映射。

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
