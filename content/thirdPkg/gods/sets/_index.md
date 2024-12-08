+++
title = "sets"
date = 2024-12-07T11:07:52+08:00
weight = 50
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/sets](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/sets)
>
> 收录该文档时间： `2024-12-07T11:07:52+08:00`

## Overview 

Package sets provides an abstract Set interface.

​	`sets` 包提供了抽象的集合（Set）接口。

In computer science, a set is an abstract data type that can store certain values and no repeated values. It is a computer implementation of the mathematical concept of a finite set. Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests a value for membership in a set.

​	在计算机科学中，集合是一种抽象数据类型，可以存储某些值且值不重复。它是数学上有限集合概念的计算机实现。与大多数其他集合类型不同，集合通常不检索特定的元素，而是测试某个值是否属于集合。

Reference: [https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29](https://en.wikipedia.org/wiki/Set_(abstract_data_type))

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Set 

``` go
type Set[T comparable] interface {
	Add(elements ...T)
	Remove(elements ...T)
	Contains(elements ...T) bool

	containers.Container[T]
}
```

Set interface that all sets implement

​	集合接口，所有集合实现此接口
