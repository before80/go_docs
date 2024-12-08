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

​	包 `trees` 提供了一个抽象的树（Tree）接口。

In computer science, a tree is a widely used abstract data type (ADT) or data structure implementing this ADT that simulates a hierarchical tree structure, with a root value and subtrees of children with a parent node, represented as a set of linked nodes.

​	在计算机科学中，树是一种广泛使用的抽象数据类型（ADT）或实现此 ADT 的数据结构，用于模拟具有根值及子树（由父节点连接的子节点）的层级树结构，通常表示为一组链接节点。

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

​	Tree 接口是所有树实现的基础接口。
