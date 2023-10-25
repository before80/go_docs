+++
title = "list"
date = 2023-05-17T09:59:21+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/container/list@go1.20.1

​	list包实现了一个双链表（a doubly linked list.）。

​	迭代一个列表(其中`l`是一个`*List`)：

```
for e := l.Front(); e != nil; e = e.Next() {
	//对e.Value做一些处理
}
```

## Example
``` go 
package main

import (
	"container/list"
	"fmt"
)

func main() {
    // 创建一个新列表并在其中放入一些数字
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

    // 迭代列表并打印其内容
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
Output:

1
2
3
4
```

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Element 

``` go 
type Element struct {

	// The value stored with this element.
    // 这个元素存储的值。
	Value any
	// contains filtered or unexported fields
}
```

​	Element结构体是链表的一个元素。

#### (*Element) Next 

``` go 
func (e *Element) Next() *Element
```

​	Next方法返回下一个列表元素或nil。

#### (*Element) Prev 

``` go 
func (e *Element) Prev() *Element
```

​	Prev方法返回前一个列表元素或nil。

### type List 

``` go 
type List struct {
	// contains filtered or unexported fields
}
```

​	List结构体表示一个双链表。List的零值是一个准备使用的空列表。

#### func New 

``` go 
func New() *List
```

​	New函数返回一个初始化的列表。

#### (*List) Back 

``` go 
func (l *List) Back() *Element
```

​	Back方法返回列表`l`的最后一个元素，如果列表为空，则返回nil。

#### (*List) Front 

``` go 
func (l *List) Front() *Element
```

​	Front方法返回列表`l`的第一个元素，如果列表为空则返回nil。

#### (*List) Init 

``` go 
func (l *List) Init() *List
```

​	Init方法初始化或清除列表`l`。

#### (*List) InsertAfter 

``` go 
func (l *List) InsertAfter(v any, mark *Element) *Element
```

​	InsertAfter方法在`mark`之后插入一个新的元素e，其值为v，并返回e。如果`mark`不是`l`的一个元素，该列表不会被修改。`mark`不能是nil。

#### (*List) InsertBefore 

``` go 
func (l *List) InsertBefore(v any, mark *Element) *Element
```

​	InsertBefore方法在`mark`之前插入一个新的元素e，其值为v，并返回e，如果`mark`不是`l`的一个元素，该列表就不会被修改。`mark`不能是nil。

#### (*List) Len 

``` go 
func (l *List) Len() int
```

​	Len方法返回列表`l`的元素数，其复杂度为O(1)。

#### (*List) MoveAfter  <- go1.2

``` go 
func (l *List) MoveAfter(e, mark *Element)
```

​	MoveAfter方法将元素`e`移动到`mark`之后的新位置。如果`e`或`mark`不是`l`的一个元素，或者`e == mark`，该列表不会被修改。`e`和`mark`元素不能是nil。

#### (*List) MoveBefore  <- go1.2

``` go 
func (l *List) MoveBefore(e, mark *Element)
```

​	MoveBefore方法将元素`e`移动到`mark`之前的新位置。如果`e`或`mark`不是`l`的一个元素，或者`e == mark`，该列表不会被修改。`e`和`mark`元素不能是nil。

#### (*List) MoveToBack 

``` go 
func (l *List) MoveToBack(e *Element)
```

​	MoveToBack方法把元素`e`移到列表`l`的后面。如果`e`不是`l`的一个元素，该列表不会被修改。该元素不能是nil。

#### (*List) MoveToFront 

``` go 
func (l *List) MoveToFront(e *Element)
```

​	MoveToFront方法把元素`e`移到列表`l`的前面，如果`e`不是`l`的元素，该列表不会被修改。该元素不能是nil。

#### (*List) PushBack 

``` go 
func (l *List) PushBack(v any) *Element
```

​	PushBack方法在列表`l`的后面插入一个新元素e，其值为v，并返回e。

#### (*List) PushBackList 

``` go 
func (l *List) PushBackList(other *List)
```

​	PushBackList方法在列表`l`的后面插入一个另一个列表的副本。它们不能是nil。

#### (*List) PushFront 

``` go 
func (l *List) PushFront(v any) *Element
```

​	PushFront方法在列表`l`的前面插入一个值为v的新元素e，并返回e。

#### (*List) PushFrontList 

``` go 
func (l *List) PushFrontList(other *List)
```

​	PushFrontList方法在列表`l`的前面插入一个另一个列表的副本。它们不能是nil。

#### (*List) Remove 

``` go 
func (l *List) Remove(e *Element) any
```

​	如果`e`是列表`l`的一个元素，Remove方法将`e`从`l`中移除，并返回元素值`e.Value`。该元素不能是nil。