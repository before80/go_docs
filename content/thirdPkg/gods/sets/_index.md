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

In computer science, a set is an abstract data type that can store certain values and no repeated values. It is a computer implementation of the mathematical concept of a finite set. Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests a value for membership in a set.

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
