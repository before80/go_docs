+++
title = "list"
date = 2023-05-17T09:59:21+08:00
weight = 2
description = ""
isCJKLanguage = true
draft = false
+++
# list

https://pkg.go.dev/container/list@go1.20.1



Package list implements a doubly linked list.

包list实现了一个双链表。

To iterate over a list (where l is a *List):

要在一个列表上进行迭代(其中l是一个*List)：

```
for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
	//对e.Value做一些处理
}
```

##### Example
``` go 
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

Element是链接列表的一个元素。

#### (*Element) Next 

``` go 
func (e *Element) Next() *Element
```

Next returns the next list element or nil.

Next 返回下一个列表元素或nil。

#### (*Element) Prev 

``` go 
func (e *Element) Prev() *Element
```

Prev returns the previous list element or nil.

Prev返回前一个列表元素或nil。

### type List 

``` go 
type List struct {
	// contains filtered or unexported fields
}
```

List represents a doubly linked list. The zero value for List is an empty list ready to use.

List表示一个双链表。List的零值是一个准备使用的空列表。

#### func New 

``` go 
func New() *List
```

New returns an initialized list.

New返回一个初始化的列表。

#### (*List) Back 

``` go 
func (l *List) Back() *Element
```

Back returns the last element of list l or nil if the list is empty.

Back返回列表l的最后一个元素，如果列表为空，则返回nil。

#### (*List) Front 

``` go 
func (l *List) Front() *Element
```

Front returns the first element of list l or nil if the list is empty.

Front返回列表l的第一个元素，如果列表为空则返回nil。

#### (*List) Init 

``` go 
func (l *List) Init() *List
```

Init initializes or clears list l.

Init 初始化或清除列表l。

#### (*List) InsertAfter 

``` go 
func (l *List) InsertAfter(v any, mark *Element) *Element
```

InsertAfter inserts a new element e with value v immediately after mark and returns e. If mark is not an element of l, the list is not modified. The mark must not be nil.

InsertAfter在mark之后插入一个新的元素e，其值为v，并返回e。如果mark不是l的一个元素，列表不会被修改。mark不能是nil。

#### (*List) InsertBefore 

``` go 
func (l *List) InsertBefore(v any, mark *Element) *Element
```

InsertBefore inserts a new element e with value v immediately before mark and returns e. If mark is not an element of l, the list is not modified. The mark must not be nil.

InsertBefore在mark之前插入一个新的元素e，其值为v，并返回e，如果mark不是l的一个元素，列表就不会被修改。mark不能是nil。

#### (*List) Len 

``` go 
func (l *List) Len() int
```

Len returns the number of elements of list l. The complexity is O(1).

Len返回列表l的元素数，其复杂度为O(1)。

#### (*List) MoveAfter  <- go1.2

``` go 
func (l *List) MoveAfter(e, mark *Element)
```

MoveAfter moves element e to its new position after mark. If e or mark is not an element of l, or e == mark, the list is not modified. The element and mark must not be nil.

MoveAfter将元素e移动到mark之后的新位置。如果e或mark不是l的一个元素，或者e == mark，列表不会被修改。元素和mark不能是nil。

#### (*List) MoveBefore  <- go1.2

``` go 
func (l *List) MoveBefore(e, mark *Element)
```

MoveBefore moves element e to its new position before mark. If e or mark is not an element of l, or e == mark, the list is not modified. The element and mark must not be nil.

MoveBefore将元素e移动到mark之前的新位置。如果e或mark不是l的一个元素，或者e == mark，列表不会被修改。元素和mark不能是nil。

#### (*List) MoveToBack 

``` go 
func (l *List) MoveToBack(e *Element)
```

MoveToBack moves element e to the back of list l. If e is not an element of l, the list is not modified. The element must not be nil.

MoveToBack把元素e移到列表l的后面。如果e不是l的一个元素，列表不会被修改。该元素不能是nil。

#### (*List) MoveToFront 

``` go 
func (l *List) MoveToFront(e *Element)
```

MoveToFront moves element e to the front of list l. If e is not an element of l, the list is not modified. The element must not be nil.

MoveToFront把元素e移到列表l的前面，如果e不是l的元素，列表不被修改。该元素不能是nil。

#### (*List) PushBack 

``` go 
func (l *List) PushBack(v any) *Element
```

PushBack inserts a new element e with value v at the back of list l and returns e.

PushBack在列表l的后面插入一个新元素e，其值为v，并返回e。

#### (*List) PushBackList 

``` go 
func (l *List) PushBackList(other *List)
```

PushBackList inserts a copy of another list at the back of list l. The lists l and other may be the same. They must not be nil.

PushBackList在列表l的后面插入一个另一个列表的副本。它们不能是nil。

#### (*List) PushFront 

``` go 
func (l *List) PushFront(v any) *Element
```

PushFront inserts a new element e with value v at the front of list l and returns e.

PushFront在列表l的前面插入一个值为v的新元素e，并返回e。

#### (*List) PushFrontList 

``` go 
func (l *List) PushFrontList(other *List)
```

PushFrontList inserts a copy of another list at the front of list l. The lists l and other may be the same. They must not be nil.

PushFrontList在列表l的前面插入一个另一个列表的副本。它们不能是nil。

#### (*List) Remove 

``` go 
func (l *List) Remove(e *Element) any
```

Remove removes e from l if e is an element of list l. It returns the element value e.Value. The element must not be nil.

如果e是列表l的一个元素，Remove将e从l中移除，并返回元素值e.Value。该元素不能是nil。