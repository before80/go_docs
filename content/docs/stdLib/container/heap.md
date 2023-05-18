+++
title = "heap"
date = 2023-05-17T09:59:21+08:00
weight = 1
description = ""
isCJKLanguage = true
draft = false
+++
# heap

https://pkg.go.dev/container/heap@go1.20.1



Package heap provides heap operations for any type that implements heap.Interface. A heap is a tree with the property that each node is the minimum-valued node in its subtree.

包heap为任何实现heap.Interface的类型提供堆操作。堆是一棵树，其属性是每个节点都是其子树中的最小值节点。

The minimum element in the tree is the root, at index 0.

树中的最小元素是根，索引为0。

A heap is a common way to implement a priority queue. To build a priority queue, implement the Heap interface with the (negative) priority as the ordering for the Less method, so Push adds items while Pop removes the highest-priority item from the queue. The Examples include such an implementation; the file example_pq_test.go has the complete source.

heap 是实现优先级队列的一种常见方式。要建立一个优先级队列，要用(负)优先级作为Less方法的排序来实现Heap接口，因此Push增加项目，而Pop从队列中删除优先级最高的项目。实例中包括这样的实现；文件example_pq_test.go中有完整的源代码。

##### Example
``` go 
```

##### Example
``` go 
```








## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func Fix  <- go1.2

``` go 
func Fix(h Interface, i int)
```

Fix re-establishes the heap ordering after the element at index i has changed its value. Changing the value of the element at index i and then calling Fix is equivalent to, but less expensive than, calling Remove(h, i) followed by a Push of the new value. The complexity is O(log n) where n = h.Len().

Fix在索引i处的元素改变其值后重新建立堆的顺序。改变索引i处元素的值，然后调用Fix，相当于调用Remove(h, i)，然后推送新的值，但成本较低。复杂度是O(log n)，其中n = h.Len()。

#### func Init 

``` go 
func Init(h Interface)
```

Init establishes the heap invariants required by the other routines in this package. Init is idempotent with respect to the heap invariants and may be called whenever the heap invariants may have been invalidated. The complexity is O(n) where n = h.Len().

Init建立了本包中其他例程所要求的堆不变性。Init对于堆不变性来说是等价的，只要堆不变性可能已经失效，就可以调用它。复杂度为O(n)，其中n = h.Len()。

#### func Pop 

``` go 
func Pop(h Interface) any
```

Pop removes and returns the minimum element (according to Less) from the heap. The complexity is O(log n) where n = h.Len(). Pop is equivalent to Remove(h, 0).

Pop删除并返回堆中的最小元素(根据Less)。复杂度为O(log n)，其中n = h.Len()。Pop等同于Remove(h, 0)。

#### func Push 

``` go 
func Push(h Interface, x any)
```

Push pushes the element x onto the heap. The complexity is O(log n) where n = h.Len().

Push将元素x推到堆上。复杂度为O(log n)，其中n = h.Len()。

#### func Remove 

``` go 
func Remove(h Interface, i int) any
```

Remove removes and returns the element at index i from the heap. The complexity is O(log n) where n = h.Len().

Remove从堆中删除并返回索引为i的元素。复杂度为O(log n)，其中n = h.Len()。

## 类型

### type Interface 

``` go 
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()  //添加x为元素Len()
	Pop() any   // remove and return element Len() - 1.  // 删除并返回元素Len() - 1.
}
```

The Interface type describes the requirements for a type using the routines in this package. Any type that implements it may be used as a min-heap with the following invariants (established after Init has been called or if the data is empty or sorted):

接口类型描述了一个使用本包中的例程的类型的要求。任何实现它的类型都可以作为最小堆使用，并具有以下不变性(在调用Init后或数据为空或被排序后建立)：

```
!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
```

Note that Push and Pop in this interface are for package heap's implementation to call. To add and remove things from the heap, use heap.Push and heap.Pop.

注意，这个接口中的Push和Pop是供包堆的实现调用的。要从堆中添加和删除东西，请使用heap.Push和heap.Pop。