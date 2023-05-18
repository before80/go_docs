+++
title = "ring"
date = 2023-05-17T09:59:21+08:00
weight = 3
description = ""
isCJKLanguage = true
draft = false
+++
# ring

https://pkg.go.dev/container/ring@go1.20.1



Package ring implements operations on circular lists.

ring 包实现了对循环列表的操作。











## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type [Ring](https://cs.opensource.google/go/go/+/go1.20.1:src/container/ring/ring.go;l=13) 

``` go 
type Ring struct {
	Value any // for use by client; untouched by this library// 供客户端使用；本库不触及。
	// contains filtered or unexported fields
}
```

A Ring is an element of a circular list, or ring. Rings do not have a beginning or end; a pointer to any ring element serves as reference to the entire ring. Empty rings are represented as nil Ring pointers. The zero value for a Ring is a one-element ring with a nil Value.

Ring是一个循环列表的一个元素，或者说是环。环没有开始或结束；任何环元素的指针都可以作为整个环的参考。空的环表示为nil Ring指针。一个环的零值是一个具有nil值的单元素环。

#### func [New](https://cs.opensource.google/go/go/+/go1.20.1:src/container/ring/ring.go;l=60) 

``` go 
func New(n int) *Ring
```

New creates a ring of n elements.

New创建一个有n个元素的环。

#### (*Ring) [Do](https://cs.opensource.google/go/go/+/go1.20.1:src/container/ring/ring.go;l=129) 

``` go 
func (r *Ring) Do(f func(any))
```

Do calls function f on each element of the ring, in forward order. The behavior of Do is undefined if f changes *r.

Do在环的每个元素上以正向顺序调用函数f。如果f改变了*r，Do的行为是未定义的。

##### Example
``` go 
```

#### (*Ring) [Len](https://cs.opensource.google/go/go/+/go1.20.1:src/container/ring/ring.go;l=116) 

``` go 
func (r *Ring) Len() int
```

Len computes the number of elements in ring r. It executes in time proportional to the number of elements.

Len计算环中元素的数量，执行时间与元素的数量成正比。

##### Example
``` go 
```

#### (*Ring) [Link](https://cs.opensource.google/go/go/+/go1.20.1:src/container/ring/ring.go;l=90) 

``` go 
func (r *Ring) Link(s *Ring) *Ring
```

Link connects ring r with ring s such that r.Next() becomes s and returns the original value for r.Next(). r must not be empty.

Link将r环和s环连接起来，这样r.Next()就变成了s，并返回r.Next()的原始值，r不能为空。

If r and s point to the same ring, linking them removes the elements between r and s from the ring. The removed elements form a subring and the result is a reference to that subring (if no elements were removed, the result is still the original value for r.Next(), and not nil).

如果r和s指向同一个环，连接它们会从环中移除r和s之间的元素。被移除的元素形成一个子环，结果是对该子环的引用(如果没有元素被移除，结果仍然是r.Next()的原始值，而不是nil)。

If r and s point to different rings, linking them creates a single ring with the elements of s inserted after r. The result points to the element following the last element of s after insertion.

如果r和s指向不同的环，连接它们会创建一个单一的环，其中s的元素插入到r之后。

##### Example
``` go 
```

#### (*Ring) [Move](https://cs.opensource.google/go/go/+/go1.20.1:src/container/ring/ring.go;l=42) 

``` go 
func (r *Ring) Move(n int) *Ring
```

Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0) in the ring and returns that ring element. r must not be empty.

Move在环中向后(n < 0)或向前(n >= 0)移动n % r.Len()元素，并返回该环元素。

##### Example
``` go 
```

#### (*Ring) [Next](https://cs.opensource.google/go/go/+/go1.20.1:src/container/ring/ring.go;l=25) 

``` go 
func (r *Ring) Next() *Ring
```

Next returns the next ring element. r must not be empty.

Next 返回下一个环状元素，r必须不是空的。

##### Example
``` go 
```

#### (*Ring) [Prev](https://cs.opensource.google/go/go/+/go1.20.1:src/container/ring/ring.go;l=33) 

``` go 
func (r *Ring) Prev() *Ring
```

Prev returns the previous ring element. r must not be empty.

Prev返回上一个环状元素，r不能为空。

##### Example
``` go 
```

#### (*Ring) [Unlink](https://cs.opensource.google/go/go/+/go1.20.1:src/container/ring/ring.go;l=107) 

``` go 
func (r *Ring) Unlink(n int) *Ring
```

Unlink removes n % r.Len() elements from the ring r, starting at r.Next(). If n % r.Len() == 0, r remains unchanged. The result is the removed subring. r must not be empty.

Unlink从r环中移除n % r.Len()元素，从r.Next()开始。如果n % r.Len() == 0，r保持不变。结果是移除的子环。

```go 
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 6
	r := ring.New(6)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Unlink three elements from r, starting from r.Next()
	r.Unlink(3)

	// Iterate through the remaining ring and print its contents
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

}
Output:

0
4
5
```

