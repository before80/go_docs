+++
title = "glist"
date = 2024-03-21T17:44:28+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/glist](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/glist)

Package glist provides most commonly used doubly linked list container which also supports concurrent-safe/unsafe switch feature.

​	软件包 glist 提供了最常用的双链表容器，同时支持并发安全/不安全交换功能。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Element

```go
type Element = list.Element
```

Element the item type of the list.

​	元素列表的项类型。

### type List

```go
type List struct {
	// contains filtered or unexported fields
}
```

List is a doubly linked list containing a concurrent-safe/unsafe switch. The switch should be set when its initialization and cannot be changed then.

​	List 是包含并发安全/不安全交换机的双向链表。开关应在初始化时设置，然后无法更改。

#### func New

```go
func New(safe ...bool) *List
```

New creates and returns a new empty doubly linked list.

​	new 创建并返回一个新的空双向链表。

##### Example

``` go
```

#### func NewFrom

```go
func NewFrom(array []interface{}, safe ...bool) *List
```

NewFrom creates and returns a list from a copy of given slice `array`. The parameter `safe` is used to specify whether using list in concurrent-safety, which is false in default.

​	NewFrom 从给定切片 `array` 的副本创建并返回一个列表。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 list，默认为 false。

##### Example

``` go
```

#### (*List) Back

```go
func (l *List) Back() (e *Element)
```

Back returns the last element of list `l` or nil if the list is empty.

​	如果列表为空，则返回 list `l` 的最后一个元素或 nil。

##### Example

``` go
```

#### (*List) BackAll

```go
func (l *List) BackAll() (values []interface{})
```

BackAll copies and returns values of all elements from back of `l` as slice.

​	BackAll 复制并返回 `l` as slice 背面所有元素的值。

##### Example

``` go
```

#### (*List) BackValue

```go
func (l *List) BackValue() (value interface{})
```

BackValue returns value of the last element of `l` or nil if the list is empty.

​	如果列表为空，BackValue 返回 或 `l` nil 的最后一个元素的值。

##### Example

``` go
```

#### (*List) Clear

```go
func (l *List) Clear()
```

Clear is alias of RemoveAll.

​	Clear 是 RemoveAll 的别名。

#### (*List) DeepCopy

```go
func (l *List) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*List) Front

```go
func (l *List) Front() (e *Element)
```

Front returns the first element of list `l` or nil if the list is empty.

​	如果列表为空，则 Front 返回 list `l` 或 nil 的第一个元素。

##### Example

``` go
```

#### (*List) FrontAll

```go
func (l *List) FrontAll() (values []interface{})
```

FrontAll copies and returns values of all elements from front of `l` as slice.

​	FrontAll 复制并返回 `l` as slice 前面所有元素的值。

##### Example

``` go
```

#### (*List) FrontValue

```go
func (l *List) FrontValue() (value interface{})
```

FrontValue returns value of the first element of `l` or nil if the list is empty.

​	如果列表为空，则 FrontValue 返回 或 `l` nil 的第一个元素的值。

##### Example

``` go
```

#### (*List) InsertAfter

```go
func (l *List) InsertAfter(p *Element, v interface{}) (e *Element)
```

InsertAfter inserts a new element `e` with value `v` immediately after `p` and returns `e`. If `p` is not an element of `l`, the list is not modified. The `p` must not be nil.

​	InsertAfter 紧跟其后 `p` 插入一个值 `v` 为新元素 `e` ，并返回 `e` 。如果不是 `p` 的 `l` 元素，则不会修改列表。不能 `p` 为零。

##### Example

``` go
```

#### (*List) InsertBefore

```go
func (l *List) InsertBefore(p *Element, v interface{}) (e *Element)
```

InsertBefore inserts a new element `e` with value `v` immediately before `p` and returns `e`. If `p` is not an element of `l`, the list is not modified. The `p` must not be nil.

​	InsertBefore 在紧接之前 `p` 插入一个值 `v` 的新元素 `e` ，并返回 `e` 。如果不是 `p` 的 `l` 元素，则不会修改列表。不能 `p` 为零。

##### Example

``` go
```

#### (*List) Iterator

```go
func (l *List) Iterator(f func(e *Element) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

#### (*List) IteratorAsc

```go
func (l *List) IteratorAsc(f func(e *Element) bool)
```

IteratorAsc iterates the list readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序迭代列表只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*List) IteratorDesc

```go
func (l *List) IteratorDesc(f func(e *Element) bool)
```

IteratorDesc iterates the list readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序迭代列表只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*List) Join

```go
func (l *List) Join(glue string) string
```

Join joins list elements with a string `glue`.

​	Join 使用字符串 `glue` 连接列表元素。

##### Example

``` go
```

#### (*List) Len

```go
func (l *List) Len() (length int)
```

Len returns the number of elements of list `l`. The complexity is O(1).

​	Len 返回 list `l` 的元素数。复杂度为 O（1）。

##### Example

``` go
```

#### (*List) LockFunc

```go
func (l *List) LockFunc(f func(list *list.List))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

​	LockFunc 在 RWMutex.Lock 中使用给定的回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (List) MarshalJSON

```go
func (l List) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*List) MoveAfter

```go
func (l *List) MoveAfter(e, p *Element)
```

MoveAfter moves element `e` to its new position after `p`. If `e` or `p` is not an element of `l`, or `e` == `p`, the list is not modified. The element and `p` must not be nil.

​	MoveAfter 将元素 `e` 移动到其 之后 `p` 的新位置。如果 `e` or `p` 不是 `l` 的元素，或者 `e` == `p` ，则不会修改列表。元素 和 `p` 不能为 nil。

##### Example

``` go
```

#### (*List) MoveBefore

```go
func (l *List) MoveBefore(e, p *Element)
```

MoveBefore moves element `e` to its new position before `p`. If `e` or `p` is not an element of `l`, or `e` == `p`, the list is not modified. The element and `p` must not be nil.

​	MoveBefore 将元素 `e` 移动到其之前 `p` 的新位置。如果 `e` or `p` 不是 `l` 的元素，或者 `e` == `p` ，则不会修改列表。元素 和 `p` 不能为零。

##### Example

``` go
```

#### (*List) MoveToBack

```go
func (l *List) MoveToBack(e *Element)
```

MoveToBack moves element `e` to the back of list `l`. If `e` is not an element of `l`, the list is not modified. The element must not be nil.

​	MoveToBack 将元素 `e` 移动到列表 `l` 的后面。如果不是 `e` 的 `l` 元素，则不会修改列表。元素不能为零。

##### Example

``` go
```

#### (*List) MoveToFront

```go
func (l *List) MoveToFront(e *Element)
```

MoveToFront moves element `e` to the front of list `l`. If `e` is not an element of `l`, the list is not modified. The element must not be nil.

​	MoveToFront 将元素 `e` 移动到列表 `l` 的前面。如果不是 `e` 的 `l` 元素，则不会修改列表。元素不能为零。

##### Example

``` go
```

#### (*List) PopBack

```go
func (l *List) PopBack() (value interface{})
```

PopBack removes the element from back of `l` and returns the value of the element.

​	PopBack 从 的 `l` 后面删除元素并返回元素的值。

##### Example

``` go
```

#### (*List) PopBackAll

```go
func (l *List) PopBackAll() []interface{}
```

PopBackAll removes all elements from back of `l` and returns values of the removed elements as slice.

​	PopBackAll 从背面 `l` 删除所有元素，并将删除元素的值作为切片返回。

##### Example

``` go
```

#### (*List) PopBacks

```go
func (l *List) PopBacks(max int) (values []interface{})
```

PopBacks removes `max` elements from back of `l` and returns values of the removed elements as slice.

​	PopBacks从背面 `l` 删除 `max` 元素，并将删除元素的值作为切片返回。

##### Example

``` go
```

#### (*List) PopFront

```go
func (l *List) PopFront() (value interface{})
```

PopFront removes the element from front of `l` and returns the value of the element.

​	PopFront 从前面 `l` 删除元素并返回元素的值。

##### Example

``` go
```

#### (*List) PopFrontAll

```go
func (l *List) PopFrontAll() []interface{}
```

PopFrontAll removes all elements from front of `l` and returns values of the removed elements as slice.

​	PopFrontAll 从前面 `l` 删除所有元素，并将删除的元素的值作为切片返回。

##### Example

``` go
```

#### (*List) PopFronts

```go
func (l *List) PopFronts(max int) (values []interface{})
```

PopFronts removes `max` elements from front of `l` and returns values of the removed elements as slice.

​	PopFronts 从前面 `l` 删除 `max` 元素，并将删除元素的值作为切片返回。

##### Example

``` go
```

#### (*List) PushBack

```go
func (l *List) PushBack(v interface{}) (e *Element)
```

PushBack inserts a new element `e` with value `v` at the back of list `l` and returns `e`.

​	PushBack 在列表 `l` 后面插入一个值 `v` 为的新元素 `e` 并返回 `e` 。

##### Example

``` go
```

#### (*List) PushBackList

```go
func (l *List) PushBackList(other *List)
```

PushBackList inserts a copy of an other list at the back of list `l`. The lists `l` and `other` may be the same, but they must not be nil.

​	PushBackList 在列表 `l` 的后面插入另一个列表的副本。列表 `l` 和 `other` 可能相同，但不能为零。

##### Example

``` go
```

#### (*List) PushBacks

```go
func (l *List) PushBacks(values []interface{})
```

PushBacks inserts multiple new elements with values `values` at the back of list `l`.

​	PushBacks 插入多个新元素，其值 `values` 位于列表 `l` 的后面。

##### Example

``` go
```

#### (*List) PushFront

```go
func (l *List) PushFront(v interface{}) (e *Element)
```

PushFront inserts a new element `e` with value `v` at the front of list `l` and returns `e`.

​	PushFront 在列表 `l` 前面插入一个值 `v` 为新元素 `e` 并返回 `e` 。

##### Example

``` go
```

#### (*List) PushFrontList

```go
func (l *List) PushFrontList(other *List)
```

PushFrontList inserts a copy of an other list at the front of list `l`. The lists `l` and `other` may be the same, but they must not be nil.

​	PushFrontList 在列表 `l` 的前面插入另一个列表的副本。列表 `l` 和 `other` 可能相同，但不能为零。

##### Example

``` go
```

#### (*List) PushFronts

```go
func (l *List) PushFronts(values []interface{})
```

PushFronts inserts multiple new elements with values `values` at the front of list `l`.

​	PushFronts 插入多个新元素，其值 `values` 位于列表 `l` 的前面。

##### Example

``` go
```

#### (*List) RLockFunc

```go
func (l *List) RLockFunc(f func(list *list.List))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

​	RLockFunc 在 RWMutex.RLock 中使用给定的回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*List) Remove

```go
func (l *List) Remove(e *Element) (value interface{})
```

Remove removes `e` from `l` if `e` is an element of list `l`. It returns the element value e.Value. The element must not be nil.

​	Remove `e` remove from `l` if `e` 是 list `l` 的一个元素。它返回元素值 e.Value。元素不能为零。

##### Example

``` go
```

#### (*List) RemoveAll

```go
func (l *List) RemoveAll()
```

RemoveAll removes all elements from list `l`.

​	RemoveAll 从列表中 `l` 删除所有元素。

##### Example

``` go
```

#### (*List) Removes

```go
func (l *List) Removes(es []*Element)
```

Removes removes multiple elements `es` from `l` if `es` are elements of list `l`.

​	Removes从 `l` 列表中 `l` 的if `es` 是元素中删除多个元素 `es` 。

##### Example

``` go
```

#### (*List) Size

```go
func (l *List) Size() int
```

Size is alias of Len.

​	size 是 Len 的别名。

##### Example

``` go
```

#### (*List) String

```go
func (l *List) String() string
```

String returns current list as a string.

​	String 以字符串形式返回当前列表。

#### (*List) UnmarshalJSON

```go
func (l *List) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*List) UnmarshalValue

```go
func (l *List) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for list.

​	UnmarshalValue 是一个接口实现，用于为 list 设置任何类型的值。