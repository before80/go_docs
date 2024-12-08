+++
title = "lists"
date = 2024-12-07T11:01:33+08:00
weight = 20
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/lists](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/lists)
>
> 收录该文档时间： `2024-12-07T11:01:33+08:00`

## Overview 

Package lists provides an abstract List interface.

​	`lists` 包提供了一个抽象的 `List` 接口。

In computer science, a list or sequence is an abstract data type that represents an ordered sequence of values, where the same value may occur more than once. An instance of a list is a computer representation of the mathematical concept of a finite sequence; the (potentially) infinite analog of a list is a stream. Lists are a basic example of containers, as they contain other values. If the same value occurs multiple times, each occurrence is considered a distinct item.

​	在计算机科学中，列表或序列是一种抽象数据类型，表示值的有序序列，其中相同的值可以出现多次。列表的一个实例是数学上有限序列概念的计算机表示；列表的（可能）无限类似物是流（stream）。列表是容器的一个基本示例，因为它们包含其他值。如果相同的值多次出现，每次出现都被视为一个独立的项。

Reference: [https://en.wikipedia.org/wiki/List_%28abstract_data_type%29](https://en.wikipedia.org/wiki/List_(abstract_data_type))

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type List 

``` go
type List[T comparable] interface {
	Get(index int) (T, bool)
	Remove(index int)
	Add(values ...T)
	Contains(values ...T) bool
	Sort(comparator utils.Comparator[T])
	Swap(index1, index2 int)
	Insert(index int, values ...T)
	Set(index int, value T)

	containers.Container[T]
}
```

List interface that all lists implement

​	所有列表实现的 `List` 接口。
