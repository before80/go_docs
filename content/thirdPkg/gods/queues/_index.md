+++
title = "queues"
date = 2024-12-07T11:05:02+08:00
weight = 40
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/queues](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/queues)
>
> 收录该文档时间： `2024-12-07T11:05:02+08:00`

## Overview 

Package queues provides an abstract Queue interface.

​	包 `queues` 提供了一个抽象的队列接口。

In computer science, a queue is a collection of entities that are maintained in a sequence and can be modified by the addition of entities at one end of the sequence and the removal of entities from the other end of the sequence. By convention, the end of the sequence at which elements are added is called the back, tail, or rear of the queue, and the end at which elements are removed is called the head or front of the queue, analogously to the words used when people line up to wait for goods or services. The operation of adding an element to the rear of the queue is known as enqueue, and the operation of removing an element from the front is known as dequeue. Other operations may also be allowed, often including a peek or front operation that returns the value of the next element to be dequeued without remove it.

​	在计算机科学中，队列是一个维护实体顺序的集合，可以通过在序列的一端添加实体以及在另一端移除实体来进行修改。按照惯例，添加元素的一端称为队列的尾部（`back`、`tail` 或 `rear`），移除元素的一端称为队列的头部（`head` 或 `front`），类似于人们排队等候商品或服务时所使用的术语。将元素添加到队列尾部的操作称为 `enqueue`，从队列头部移除元素的操作称为 `dequeue`。通常还允许其他操作，例如返回下一个将被移除的元素值但不移除它的 `peek` 或 `front` 操作。

Reference: https://en.wikipedia.org/wiki/Queue_(abstract_data_type)

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Queue 

``` go
type Queue[T comparable] interface {
	Enqueue(value T)
	Dequeue() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
}
```

Queue interface that all queues implement

​	**Queue** 是所有队列实现的接口
