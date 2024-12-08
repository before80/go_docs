+++
title = "maps"
date = 2024-12-07T11:02:55+08:00
weight = 30
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：
>
> 收录该文档时间： `2024-12-07T11:02:55+08:00`

## Overview 

Package maps provides an abstract Map interface.

In computer science, an associative array, map, symbol table, or dictionary is an abstract data type composed of a collection of (key, value) pairs, such that each possible key appears just once in the collection.

Operations associated with this data type allow: - the addition of a pair to the collection - the removal of a pair from the collection - the modification of an existing pair - the lookup of a value associated with a particular key

Reference: https://en.wikipedia.org/wiki/Associative_array

### Index 

- [type BidiMap](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps#BidiMap)
- [type Map](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/maps#Map)

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type BidiMap 

``` go
type BidiMap[K comparable, V comparable] interface {
	GetKey(value V) (key K, found bool)

	Map[K, V]
}
```

BidiMap interface that all bidirectional maps implement (extends the Map interface)

### type Map 

``` go
type Map[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Remove(key K)
	Keys() []K

	containers.Container[V]
}
```

Map interface that all maps implement
