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

In computer science, a queue is a collection of entities that are maintained in a sequence and can be modified by the addition of entities at one end of the sequence and the removal of entities from the other end of the sequence. By convention, the end of the sequence at which elements are added is called the back, tail, or rear of the queue, and the end at which elements are removed is called the head or front of the queue, analogously to the words used when people line up to wait for goods or services. The operation of adding an element to the rear of the queue is known as enqueue, and the operation of removing an element from the front is known as dequeue. Other operations may also be allowed, often including a peek or front operation that returns the value of the next element to be dequeued without remove it.

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
