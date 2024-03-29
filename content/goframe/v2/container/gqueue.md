+++
title = "gqueue"
date = 2024-03-21T17:44:52+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gqueue](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gqueue)

Package gqueue provides dynamic/static concurrent-safe queue.

​	软件包 gqueue 提供动态/静态并发安全队列。

Features:

​	特征：

1. FIFO queue(data -> list -> chan);
   FIFO 队列（数据 -> 列表 -> chan）;
2. Fast creation and initialization;
   快速创建和初始化;
3. Support dynamic queue size(unlimited queue size);
   支持动态队列大小（无限队列大小）;
4. Blocking when reading data from queue;
   从队列中读取数据时阻塞;

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Queue

```go
type Queue struct {
	C chan interface{} // Underlying channel for data reading.
	// contains filtered or unexported fields
}
```

Queue is a concurrent-safe queue built on doubly linked list and channel.

​	队列是基于双链表和通道构建的并发安全队列。

#### func New

```go
func New(limit ...int) *Queue
```

New returns an empty queue object. Optional parameter `limit` is used to limit the size of the queue, which is unlimited in default. When `limit` is given, the queue will be static and high performance which is comparable with stdlib channel.

​	New 返回一个空队列对象。可选参数 `limit` 用于限制队列的大小，默认情况下是无限的。当给定时 `limit` ，队列将是静态的和高性能的，这与 stdlib 通道相当。

##### Example

``` go
```

#### (*Queue) Close

```go
func (q *Queue) Close()
```

Close closes the queue. Notice: It would notify all goroutines return immediately, which are being blocked reading using Pop method.

​	关闭 关闭队列。注意：它会立即通知所有 goroutines 返回，这些 goroutine 正在使用 Pop 方法阻止读取。

##### Example

``` go
```

#### (*Queue) Len

```go
func (q *Queue) Len() (length int64)
```

Len returns the length of the queue. Note that the result might not be accurate if using unlimited queue size as there’s an asynchronous channel reading the list constantly.

​	Len 返回队列的长度。请注意，如果使用无限的队列大小，则结果可能不准确，因为有一个异步通道不断读取列表。

##### Example

``` go
```

#### (*Queue) Pop

```go
func (q *Queue) Pop() interface{}
```

Pop pops an item from the queue in FIFO way. Note that it would return nil immediately if Pop is called after the queue is closed.

​	Pop 以 FIFO 方式从队列中弹出一个项目。请注意，如果在队列关闭后调用 Pop，它将立即返回 nil。

##### Example

``` go
```

#### (*Queue) Push

```go
func (q *Queue) Push(v interface{})
```

Push pushes the data `v` into the queue. Note that it would panic if Push is called after the queue is closed.

​	Push 将数据 `v` 推送到队列中。请注意，如果在队列关闭后调用 Push，则会出现恐慌。

##### Example

``` go
```

#### (*Queue) Size

```go
func (q *Queue) Size() int64
```

Size is alias of Len. Deprecated: use Len instead.

​	size 是 Len 的别名。已弃用：请改用 Len。

##### Example

``` go
```

