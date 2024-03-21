+++
title = "glist"
date = 2024-03-21T17:44:28+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/glist

Package glist provides most commonly used doubly linked list container which also supports concurrent-safe/unsafe switch feature.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type Element 

``` go
type Element = list.Element
```

Element the item type of the list.

#### type List 

``` go
type List struct {
	// contains filtered or unexported fields
}
```

List is a doubly linked list containing a concurrent-safe/unsafe switch. The switch should be set when its initialization and cannot be changed then.

##### func New 

``` go
func New(safe ...bool) *List
```

New creates and returns a new empty doubly linked list.

##### Example

``` go
```
##### func NewFrom 

``` go
func NewFrom(array []interface{}, safe ...bool) *List
```

NewFrom creates and returns a list from a copy of given slice `array`. The parameter `safe` is used to specify whether using list in concurrent-safety, which is false in default.

##### Example

``` go
```
##### (*List) Back 

``` go
func (l *List) Back() (e *Element)
```

Back returns the last element of list `l` or nil if the list is empty.

##### Example

``` go
```
##### (*List) BackAll 

``` go
func (l *List) BackAll() (values []interface{})
```

BackAll copies and returns values of all elements from back of `l` as slice.

##### Example

``` go
```
##### (*List) BackValue 

``` go
func (l *List) BackValue() (value interface{})
```

BackValue returns value of the last element of `l` or nil if the list is empty.

##### Example

``` go
```
##### (*List) Clear 

``` go
func (l *List) Clear()
```

Clear is alias of RemoveAll.

##### (*List) DeepCopy <-2.1.0

``` go
func (l *List) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*List) Front 

``` go
func (l *List) Front() (e *Element)
```

Front returns the first element of list `l` or nil if the list is empty.

##### Example

``` go
```
##### (*List) FrontAll 

``` go
func (l *List) FrontAll() (values []interface{})
```

FrontAll copies and returns values of all elements from front of `l` as slice.

##### Example

``` go
```
##### (*List) FrontValue 

``` go
func (l *List) FrontValue() (value interface{})
```

FrontValue returns value of the first element of `l` or nil if the list is empty.

##### Example

``` go
```
##### (*List) InsertAfter 

``` go
func (l *List) InsertAfter(p *Element, v interface{}) (e *Element)
```

InsertAfter inserts a new element `e` with value `v` immediately after `p` and returns `e`. If `p` is not an element of `l`, the list is not modified. The `p` must not be nil.

##### Example

``` go
```
##### (*List) InsertBefore 

``` go
func (l *List) InsertBefore(p *Element, v interface{}) (e *Element)
```

InsertBefore inserts a new element `e` with value `v` immediately before `p` and returns `e`. If `p` is not an element of `l`, the list is not modified. The `p` must not be nil.

##### Example

``` go
```
##### (*List) Iterator 

``` go
func (l *List) Iterator(f func(e *Element) bool)
```

Iterator is alias of IteratorAsc.

##### (*List) IteratorAsc 

``` go
func (l *List) IteratorAsc(f func(e *Element) bool)
```

IteratorAsc iterates the list readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*List) IteratorDesc 

``` go
func (l *List) IteratorDesc(f func(e *Element) bool)
```

IteratorDesc iterates the list readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*List) Join 

``` go
func (l *List) Join(glue string) string
```

Join joins list elements with a string `glue`.

##### Example

``` go
```
##### (*List) Len 

``` go
func (l *List) Len() (length int)
```

Len returns the number of elements of list `l`. The complexity is O(1).

##### Example

``` go
```
##### (*List) LockFunc 

``` go
func (l *List) LockFunc(f func(list *list.List))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

##### Example

``` go
```
##### (List) MarshalJSON 

``` go
func (l List) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*List) MoveAfter 

``` go
func (l *List) MoveAfter(e, p *Element)
```

MoveAfter moves element `e` to its new position after `p`. If `e` or `p` is not an element of `l`, or `e` == `p`, the list is not modified. The element and `p` must not be nil.

##### Example

``` go
```
##### (*List) MoveBefore 

``` go
func (l *List) MoveBefore(e, p *Element)
```

MoveBefore moves element `e` to its new position before `p`. If `e` or `p` is not an element of `l`, or `e` == `p`, the list is not modified. The element and `p` must not be nil.

##### Example

``` go
```
##### (*List) MoveToBack 

``` go
func (l *List) MoveToBack(e *Element)
```

MoveToBack moves element `e` to the back of list `l`. If `e` is not an element of `l`, the list is not modified. The element must not be nil.

##### Example

``` go
```
##### (*List) MoveToFront 

``` go
func (l *List) MoveToFront(e *Element)
```

MoveToFront moves element `e` to the front of list `l`. If `e` is not an element of `l`, the list is not modified. The element must not be nil.

##### Example

``` go
```
##### (*List) PopBack 

``` go
func (l *List) PopBack() (value interface{})
```

PopBack removes the element from back of `l` and returns the value of the element.

##### Example

``` go
```
##### (*List) PopBackAll 

``` go
func (l *List) PopBackAll() []interface{}
```

PopBackAll removes all elements from back of `l` and returns values of the removed elements as slice.

##### Example

``` go
```
##### (*List) PopBacks 

``` go
func (l *List) PopBacks(max int) (values []interface{})
```

PopBacks removes `max` elements from back of `l` and returns values of the removed elements as slice.

##### Example

``` go
```
##### (*List) PopFront 

``` go
func (l *List) PopFront() (value interface{})
```

PopFront removes the element from front of `l` and returns the value of the element.

##### Example

``` go
```
##### (*List) PopFrontAll 

``` go
func (l *List) PopFrontAll() []interface{}
```

PopFrontAll removes all elements from front of `l` and returns values of the removed elements as slice.

##### Example

``` go
```
##### (*List) PopFronts 

``` go
func (l *List) PopFronts(max int) (values []interface{})
```

PopFronts removes `max` elements from front of `l` and returns values of the removed elements as slice.

##### Example

``` go
```
##### (*List) PushBack 

``` go
func (l *List) PushBack(v interface{}) (e *Element)
```

PushBack inserts a new element `e` with value `v` at the back of list `l` and returns `e`.

##### Example

``` go
```
##### (*List) PushBackList 

``` go
func (l *List) PushBackList(other *List)
```

PushBackList inserts a copy of an other list at the back of list `l`. The lists `l` and `other` may be the same, but they must not be nil.

##### Example

``` go
```
##### (*List) PushBacks 

``` go
func (l *List) PushBacks(values []interface{})
```

PushBacks inserts multiple new elements with values `values` at the back of list `l`.

##### Example

``` go
```
##### (*List) PushFront 

``` go
func (l *List) PushFront(v interface{}) (e *Element)
```

PushFront inserts a new element `e` with value `v` at the front of list `l` and returns `e`.

##### Example

``` go
```
##### (*List) PushFrontList 

``` go
func (l *List) PushFrontList(other *List)
```

PushFrontList inserts a copy of an other list at the front of list `l`. The lists `l` and `other` may be the same, but they must not be nil.

##### Example

``` go
```
##### (*List) PushFronts 

``` go
func (l *List) PushFronts(values []interface{})
```

PushFronts inserts multiple new elements with values `values` at the front of list `l`.

##### Example

``` go
```
##### (*List) RLockFunc 

``` go
func (l *List) RLockFunc(f func(list *list.List))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

##### Example

``` go
```
##### (*List) Remove 

``` go
func (l *List) Remove(e *Element) (value interface{})
```

Remove removes `e` from `l` if `e` is an element of list `l`. It returns the element value e.Value. The element must not be nil.

##### Example

``` go
```
##### (*List) RemoveAll 

``` go
func (l *List) RemoveAll()
```

RemoveAll removes all elements from list `l`.

##### Example

``` go
```
##### (*List) Removes 

``` go
func (l *List) Removes(es []*Element)
```

Removes removes multiple elements `es` from `l` if `es` are elements of list `l`.

##### Example

``` go
```
##### (*List) Size 

``` go
func (l *List) Size() int
```

Size is alias of Len.

##### Example

``` go
```
##### (*List) String 

``` go
func (l *List) String() string
```

String returns current list as a string.

##### (*List) UnmarshalJSON 

``` go
func (l *List) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*List) UnmarshalValue 

``` go
func (l *List) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for list.