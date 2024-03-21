+++
title = "gring"
date = 2024-03-21T17:44:58+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gring

Package gring provides a concurrent-safe/unsafe ring(circular lists).

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type Ring 

``` go
type Ring struct {
	// contains filtered or unexported fields
}
```

Ring is a struct of ring structure.

##### func New 

``` go
func New(cap int, safe ...bool) *Ring
```

New creates and returns a Ring structure of `cap` elements. The optional parameter `safe` specifies whether using this structure in concurrent safety, which is false in default.

##### Example

``` go
```
##### (*Ring) Cap 

``` go
func (r *Ring) Cap() int
```

Cap returns the capacity of ring.

##### Example

``` go
```
##### (*Ring) Len 

``` go
func (r *Ring) Len() int
```

Len returns the size of ring.

##### Example

``` go
```
##### (*Ring) Link 

``` go
func (r *Ring) Link(s *Ring) *Ring
```

Link connects ring r with ring s such that r.Next() becomes s and returns the original value for r.Next(). r must not be empty.

If r and s point to the same ring, linking them removes the elements between r and s from the ring. The removed elements form a sub-ring and the result is a reference to that sub-ring (if no elements were removed, the result is still the original value for r.Next(), and not nil).

If r and s point to different rings, linking them creates a single ring with the elements of s inserted after r. The result points to the element following the last element of s after insertion.

##### (*Ring) Move 

``` go
func (r *Ring) Move(n int) *Ring
```

Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0) in the ring and returns that ring element. r must not be empty.

##### Example

``` go
```
##### (*Ring) Next 

``` go
func (r *Ring) Next() *Ring
```

Next returns the next ring element. r must not be empty.

##### Example

``` go
```
##### (*Ring) Prev 

``` go
func (r *Ring) Prev() *Ring
```

Prev returns the previous ring element. r must not be empty.

##### Example

``` go
```
##### (*Ring) Put 

``` go
func (r *Ring) Put(value interface{}) *Ring
```

Put sets `value` to current item of ring and moves position to next item.

##### Example

``` go
```
##### (*Ring) RLockIteratorNext 

``` go
func (r *Ring) RLockIteratorNext(f func(value interface{}) bool)
```

RLockIteratorNext iterates and locks reading forward with given callback function `f` within RWMutex.RLock. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*Ring) RLockIteratorPrev 

``` go
func (r *Ring) RLockIteratorPrev(f func(value interface{}) bool)
```

RLockIteratorPrev iterates and locks writing backward with given callback function `f` within RWMutex.RLock. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*Ring) Set 

``` go
func (r *Ring) Set(value interface{}) *Ring
```

Set sets value to the item of current position.

##### Example

``` go
```
##### (*Ring) SliceNext 

``` go
func (r *Ring) SliceNext() []interface{}
```

SliceNext returns a copy of all item values as slice forward from current position.

##### Example

``` go
```
##### (*Ring) SlicePrev 

``` go
func (r *Ring) SlicePrev() []interface{}
```

SlicePrev returns a copy of all item values as slice backward from current position.

##### Example

``` go
```
##### (*Ring) Unlink 

``` go
func (r *Ring) Unlink(n int) *Ring
```

Unlink removes n % r.Len() elements from the ring r, starting at r.Next(). If n % r.Len() == 0, r remains unchanged. The result is the removed sub-ring. r must not be empty.

##### Example

``` go
```
##### (*Ring) Val 

``` go
func (r *Ring) Val() interface{}
```

Val returns the item's value of current position.



Example Val

``` go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/gring"
)

func main() {
	r := gring.New(10)
	r.Set(1)
	fmt.Println("Val:", r.Val())

	r.Next().Set("GoFrame")
	fmt.Println("Val:", r.Val())

}

Output:

Val: 1
Val: GoFrame
```







