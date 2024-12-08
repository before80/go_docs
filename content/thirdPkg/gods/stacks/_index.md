+++
title = "stacks"
date = 2024-12-07T11:54:58+08:00
weight = 60
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/stacks](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/stacks)
>
> 收录该文档时间： `2024-12-07T11:54:58+08:00`

## Overview 

Package stacks provides an abstract Stack interface.

​	包 `stacks` 提供了一个抽象的 `Stack` 接口。

In computer science, a stack is an abstract data type that serves as a collection of elements, with two principal operations: 

​	在计算机科学中，栈是一种抽象数据类型，它作为元素的集合，有两个主要操作：

push, which adds an element to the collection, and pop, which removes the most recently added element that was not yet removed. 

​	`push`，将一个元素添加到集合中；`pop`，移除最先被添加但尚未移除的元素。

The order in which elements come off a stack gives rise to its alternative name, LIFO (for last in, first out). Additionally, a peek operation may give access to the top without modifying the stack.

​	栈中元素的移除顺序引出了它的另一个名称，LIFO（后进先出）。此外，还可以通过 `peek` 操作访问栈顶元素而不修改栈的内容。

Reference: [https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29](https://en.wikipedia.org/wiki/Stack_(abstract_data_type))

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Stack 

``` go
type Stack[T any] interface {
	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
}
```

Stack interface that all stacks implement

​	`Stack` 接口是所有栈实现的基础。
