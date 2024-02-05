+++
title = "ring"
date = 2023-05-17T09:59:21+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/container/ring@go1.21.3](https://pkg.go.dev/container/ring@go1.21.3)

Package ring implements operations on circular lists.

​	`ring` 包实现了对循环列表（circular lists）的操作。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Ring 

``` go 
type Ring struct {
	Value any // for use by client; untouched by this library// 供客户端使用；本库不触及。
	// contains filtered or unexported fields
}
```

A Ring is an element of a circular list, or ring. Rings do not have a beginning or end; a pointer to any ring element serves as reference to the entire ring. Empty rings are represented as nil Ring pointers. The zero value for a Ring is a one-element ring with a nil Value.

​	Ring结构体是一个循环列表的一个元素，或者说是环。环没有开始或结束；任何环元素的指针都可以作为整个环的参考。空的环表示为nil Ring指针。一个环的零值是一个具有nil值的单元素环。

#### func New 

``` go 
func New(n int) *Ring
```

New creates a ring of n elements.

​	`New`函数创建一个有`n`个元素的环。

#### (*Ring) Do 

``` go 
func (r *Ring) Do(f func(any))
```

​	`Do`方法在环的每个元素上以正向顺序调用函数`f`。如果`f`改变了`*r`，Do方法的行为是未定义的。

##### Do Example
``` go 
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
    // 创建一个大小为5的新的环
	r := ring.New(5)

	// Get the length of the ring
    // 获取该环的长度
	n := r.Len()

	// Initialize the ring with some integer values
    // 使用一些整数值来初始化该环
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
    // 迭代该环并打印其内容
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

}
Output:

0
1
2
3
4
```

#### (*Ring) Len 

``` go 
func (r *Ring) Len() int
```

Do calls function f on each element of the ring, in forward order. The behavior of Do is undefined if f changes *r.

​	`Len`方法计算环中元素的数量，执行时间与元素的数量成正比。

##### Len Example
``` go 
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 4
    // 创建一个大小为4的新的环
	r := ring.New(4)

	// Print out its length
    // 打印其长度
	fmt.Println(r.Len())

}
Output:

4
```

#### (*Ring) Link 

``` go 
func (r *Ring) Link(s *Ring) *Ring
```

Link connects ring r with ring s such that r.Next() becomes s and returns the original value for r.Next(). r must not be empty.

​	Link方法将`r`环和`s`环连接起来，这样`r.Next()`就变成了`s`，并返回`r.Next()`的原始值，`r`不能为空。

If r and s point to the same ring, linking them removes the elements between r and s from the ring. The removed elements form a subring and the result is a reference to that subring (if no elements were removed, the result is still the original value for r.Next(), and not nil).

​	如果`r`和`s`指向同一个环，连接它们会从环中移除`r`和`s`之间的元素。被移除的元素形成一个子环，结果是对该子环的引用(如果没有元素被移除，结果仍然是`r.Next()`的原始值，而不是nil)。

If r and s point to different rings, linking them creates a single ring with the elements of s inserted after r. The result points to the element following the last element of s after insertion.

​	如果`r`和`s`指向不同的环，连接它们会创建一个单一的环，其中`s`的元素插入到`r`之后。

##### Link Example
``` go 
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create two rings, r and s, of size 2
    // 创建大小都为2的两个环：r 和 s
	r := ring.New(2)
	s := ring.New(2)

	// Get the length of the ring
    // 获取环的长度
	lr := r.Len()
	ls := s.Len()

	// Initialize r with 0s
    // 使用 0 来初始化 r 环
	for i := 0; i < lr; i++ {
		r.Value = 0
		r = r.Next()
	}

	// Initialize s with 1s
    // 使用 1 来初始化 s 环
	for j := 0; j < ls; j++ {
		s.Value = 1
		s = s.Next()
	}

	// Link ring r and ring s
    // Link r 和 s 环
	rs := r.Link(s)

	// Iterate through the combined ring and print its contents
    // 迭代联合后的环并打印其内容
	rs.Do(func(p any) {
		fmt.Println(p.(int))
	})

}
Output:

0
0
1
1
```

#### (*Ring) Move 

``` go 
func (r *Ring) Move(n int) *Ring
```

Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0) in the ring and returns that ring element. r must not be empty.

​	Move方法在环中向后(`n < 0`)或向前(`n >= 0`)移动`n % r.Len()`个元素，并返回该环元素。

##### Move Example
``` go 
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Move the pointer forward by three steps
	r = r.Move(3)

	// Iterate through the ring and print its contents
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

}
Output:

3
4
0
1
2
```

#### (*Ring) Next 

``` go 
func (r *Ring) Next() *Ring
```

Next returns the next ring element. r must not be empty.

​	`Next`方法返回下一个环元素，`r`必须不是空的。

##### Next Example
``` go 
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	for j := 0; j < n; j++ {
		fmt.Println(r.Value)
		r = r.Next()
	}

}
Output:

0
1
2
3
4
```

#### (*Ring) Prev 

``` go 
func (r *Ring) Prev() *Ring
```

Prev returns the previous ring element. r must not be empty.

​	`Prev`方法返回上一个环元素，`r`不能为空。

##### Prev Example
``` go 
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring backwards and print its contents
	for j := 0; j < n; j++ {
		r = r.Prev()
		fmt.Println(r.Value)
	}

}
Output:

4
3
2
1
0
```

#### (*Ring) Unlink 

``` go 
func (r *Ring) Unlink(n int) *Ring
```

Unlink removes n % r.Len() elements from the ring r, starting at r.Next(). If n % r.Len() == 0, r remains unchanged. The result is the removed subring. r must not be empty.

​	`Unlink`方法从`r`环中移除`n % r.Len()`个元素，从`r.Next()`开始。如果`n % r.Len() == 0`，`r`保持不变。结果是移除的子环。`r`不能为空。

##### Unlink Example

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

	fmt.Println("r.Value=", r.Value)
	// Unlink three elements from r, starting from r.Next()
	ra := r.Unlink(3)

	fmt.Println("ra")
	ra.Do(func(p any) {
		fmt.Println(p.(int))
	})

	// Iterate through the remaining ring and print its contents
	fmt.Println("r")
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

}

r.Value= 0
ra
1
2
3
r
0
4
5

```

