+++
title = "gqueue"
date = 2024-03-21T17:44:52+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gqueue

Package gqueue provides dynamic/static concurrent-safe queue.

Features:

1. FIFO queue(data -> list -> chan);
2. Fast creation and initialization;
3. Support dynamic queue size(unlimited queue size);
4. Blocking when reading data from queue;

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type Queue 

``` go
type Queue struct {
	C chan interface{} // Underlying channel for data reading.
	// contains filtered or unexported fields
}
```

Queue is a concurrent-safe queue built on doubly linked list and channel.

##### func New 

``` go
func New(limit ...int) *Queue
```

New returns an empty queue object. Optional parameter `limit` is used to limit the size of the queue, which is unlimited in default. When `limit` is given, the queue will be static and high performance which is comparable with stdlib channel.

##### Example

``` go
```
##### (*Queue) Close 

``` go
func (q *Queue) Close()
```

Close closes the queue. Notice: It would notify all goroutines return immediately, which are being blocked reading using Pop method.

##### Example

``` go
```
##### (*Queue) Len 

``` go
func (q *Queue) Len() (length int64)
```

Len returns the length of the queue. Note that the result might not be accurate if using unlimited queue size as there's an asynchronous channel reading the list constantly.

##### Example

``` go
```
##### (*Queue) Pop 

``` go
func (q *Queue) Pop() interface{}
```

Pop pops an item from the queue in FIFO way. Note that it would return nil immediately if Pop is called after the queue is closed.

##### Example

``` go
```
##### (*Queue) Push 

``` go
func (q *Queue) Push(v interface{})
```

Push pushes the data `v` into the queue. Note that it would panic if Push is called after the queue is closed.

##### Example

``` go
```
##### (*Queue) Size 

``` go
func (q *Queue) Size() int64
```

Size is alias of Len. Deprecated: use Len instead.

##### Example

``` go
```






<iframe allowtransparency="true" frameborder="0" scrolling="no" class="sk_ui" src="chrome-extension://gfbliohnnapiefjpjlpjnehglfpaknnc/pages/frontend.html" title="Surfingkeys" style="left: 0px; bottom: 0px; width: 1555px; height: 0px; z-index: 2147483647;"></iframe>