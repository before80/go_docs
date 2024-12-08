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

​	包 `maps` 提供了一个抽象的 `Map` 接口。

In computer science, an associative array, map, symbol table, or dictionary is an abstract data type composed of a collection of (key, value) pairs, such that each possible key appears just once in the collection.

​	在计算机科学中，关联数组、映射、符号表或字典是一种抽象数据类型，由一组 (键, 值) 对组成，使得每个可能的键在集合中只出现一次。

Operations associated with this data type allow: 

​	与此数据类型相关的操作包括：

- the addition of a pair to the collection 
  - 向集合中添加一对键值对
- the removal of a pair from the collection 
  - 从集合中删除一对键值对
- the modification of an existing pair 
  - 修改现有的键值对
- the lookup of a value associated with a particular key
  - 查找与特定键相关联的值

Reference: https://en.wikipedia.org/wiki/Associative_array

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

​	`BidiMap` 接口是所有双向映射实现的接口（扩展自 `Map` 接口）。

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

​	`Map` 接口是所有映射实现的接口。
