+++
title = "trees"
date = 2024-12-07T11:09:11+08:00
weight = 70
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/trees](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/trees)
>
> 收录该文档时间： `2024-12-07T11:09:11+08:00`

## Overview 

Package trees provides an abstract Tree interface.

In computer science, a tree is a widely used abstract data type (ADT) or data structure implementing this ADT that simulates a hierarchical tree structure, with a root value and subtrees of children with a parent node, represented as a set of linked nodes.

Reference: [https://en.wikipedia.org/wiki/Tree_%28data_structure%29](https://en.wikipedia.org/wiki/Tree_(data_structure))

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Tree 

``` go
type Tree[V any] interface {
	containers.Container[V]
}
```

Tree interface that all trees implement
