+++
title = "list"
date = 2023-05-17T09:59:21+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/container/list@go1.21.3](https://pkg.go.dev/container/list@go1.21.3)

Package list implements a doubly linked list.

​	list包实现了一个双链表（a doubly linked list.）。

To iterate over a list (where l is a *List):

​	迭代一个列表(其中`l`是一个`*List`)：

```go
for e := l.Front(); e != nil; e = e.Next() {
    // do something with e.Value
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
    // Create a new list and put some numbers in it.
    // 创建一个新列表并在其中放入一些数字
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

    // Iterate through list and print its contents.
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

Element is an element of a linked list.

​	`Element`结构体是链表的一个元素。

#### (*Element) Next 

``` go 
func (e *Element) Next() *Element
```

Next returns the next list element or nil.

​	`Next`方法返回下一个列表元素或`nil`。

#### (*Element) Prev 

``` go 
func (e *Element) Prev() *Element
```

Prev returns the previous list element or nil.

​	`Prev`方法返回前一个列表元素或`nil`。

### type List 

``` go 
type List struct {
	// contains filtered or unexported fields
}
```

List represents a doubly linked list. The zero value for List is an empty list ready to use.

​	`List`结构体表示一个双链表。`List`的零值是一个准备使用的空列表。

#### func New 

``` go 
func New() *List
```

New returns an initialized list.

​	`New`函数返回一个初始化的列表。

#### (*List) Back 

``` go 
func (l *List) Back() *Element
```

Back returns the last element of list l or nil if the list is empty.

​	`Back`方法返回列表`l`的最后一个元素，如果列表为空，则返回`nil`。

#### (*List) Front 

``` go 
func (l *List) Front() *Element
```

Front returns the first element of list l or nil if the list is empty.

​	`Front`方法返回列表`l`的第一个元素，如果列表为空则返回`nil`。

#### (*List) Init 

``` go 
func (l *List) Init() *List
```

Init initializes or clears list l.

​	`Init`方法初始化或清除列表`l`。

#### (*List) InsertAfter 

``` go 
func (l *List) InsertAfter(v any, mark *Element) *Element
```

InsertAfter inserts a new element e with value v immediately after mark and returns e. If mark is not an element of l, the list is not modified. The mark must not be nil.

​	`InsertAfter`方法在`mark`之后插入一个新的元素`e`，其值为`v`，并返回`e`。如果`mark`不是`l`的一个元素，该列表不会被修改。`mark`不能是`nil`。

#### (*List) InsertBefore 

``` go 
func (l *List) InsertBefore(v any, mark *Element) *Element
```

InsertBefore inserts a new element e with value v immediately before mark and returns e. If mark is not an element of l, the list is not modified. The mark must not be nil.

​	`InsertBefore`方法在`mark`之前插入一个新的元素`e`，其值为`v`，并返回`e`，如果`mark`不是`l`的一个元素，该列表就不会被修改。`mark`不能是`nil`。

#### (*List) Len 

``` go 
func (l *List) Len() int
```

Len returns the number of elements of list l. The complexity is O(1).

​	`Len`方法返回列表`l`的元素数，其复杂度为`O(1)`。

#### (*List) MoveAfter  <- go1.2

``` go 
func (l *List) MoveAfter(e, mark *Element)
```

MoveAfter moves element e to its new position after mark. If e or mark is not an element of l, or e == mark, the list is not modified. The element and mark must not be nil.

​	`MoveAfter`方法将元素`e`移动到`mark`之后的新位置。如果`e`或`mark`不是`l`的一个元素，或者`e == mark`，该列表不会被修改。`e`和`mark`元素不能是`nil`。

#### (*List) MoveBefore  <- go1.2

``` go 
func (l *List) MoveBefore(e, mark *Element)
```

MoveBefore moves element e to its new position before mark. If e or mark is not an element of l, or e == mark, the list is not modified. The element and mark must not be nil.

​	`MoveBefore`方法将元素`e`移动到`mark`之前的新位置。如果`e`或`mark`不是`l`的一个元素，或者`e == mark`，该列表不会被修改。`e`和`mark`元素不能是`nil`。

#### (*List) MoveToBack 

``` go 
func (l *List) MoveToBack(e *Element)
```

MoveToBack moves element e to the back of list l. If e is not an element of l, the list is not modified. The element must not be nil.

​	`MoveToBack`方法把元素`e`移到列表`l`的后面。如果`e`不是`l`的一个元素，该列表不会被修改。该元素不能是`nil`。

#### (*List) MoveToFront 

``` go 
func (l *List) MoveToFront(e *Element)
```

MoveToFront moves element e to the front of list l. If e is not an element of l, the list is not modified. The element must not be nil.

​	`MoveToFront`方法把元素`e`移到列表`l`的前面，如果`e`不是`l`的元素，该列表不会被修改。该元素不能是`nil`。

#### (*List) PushBack 

``` go 
func (l *List) PushBack(v any) *Element
```

PushBack inserts a new element e with value v at the back of list l and returns e.

​	`PushBack`方法在列表`l`的后面插入一个新元素`e`，其值为`v`，并返回`e`。

#### (*List) PushBackList 

``` go 
func (l *List) PushBackList(other *List)
```

PushBackList inserts a copy of another list at the back of list l. The lists l and other may be the same. They must not be nil.

​	`PushBackList`方法在列表`l`的后面插入一个另一个列表的副本。它们不能是`nil`。

#### (*List) PushFront 

``` go 
func (l *List) PushFront(v any) *Element
```

PushFront inserts a new element e with value v at the front of list l and returns e.

​	`PushFront`方法在列表`l`的前面插入一个值为`v`的新元素`e`，并返回`e`。

#### (*List) PushFrontList 

``` go 
func (l *List) PushFrontList(other *List)
```

PushFrontList inserts a copy of another list at the front of list l. The lists l and other may be the same. They must not be nil.

​	`PushFrontList`方法在列表`l`的前面插入一个另一个列表的副本。它们不能是`nil`。

#### (*List) Remove 

``` go 
func (l *List) Remove(e *Element) any
```

Remove removes e from l if e is an element of list l. It returns the element value e.Value. The element must not be nil.

​	如果`e`是列表`l`的一个元素，`Remove`方法将`e`从`l`中移除，并返回元素值`e.Value`。该元素不能是`nil`。