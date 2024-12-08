+++
title = "singlylinkedlist"
date = 2024-12-07T11:02:30+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/lists/singlylinkedlist](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/lists/singlylinkedlist)
>
> 收录该文档时间： `2024-12-07T11:02:30+08:00`

## Overview 

Package singlylinkedlist implements the singly-linked list.

​	`singlylinkedlist` 包实现了单链表。

Structure is not thread safe.

​	此结构非线程安全。

Reference: [https://en.wikipedia.org/wiki/List_%28abstract_data_type%29](https://en.wikipedia.org/wiki/List_(abstract_data_type))

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Iterator 

``` go
type Iterator[T comparable] struct {
	// contains filtered or unexported fields
}
```

Iterator holding the iterator's state

​	`Iterator` 保存迭代器的状态。

#### (*Iterator[T]) Begin 

``` go
func (iterator *Iterator[T]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

​	`Begin` 将迭代器重置为初始状态（位于第一个元素之前）。调用 `Next()` 获取第一个元素（如果存在）。

#### (*Iterator[T]) First 

``` go
func (iterator *Iterator[T]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`First` 将迭代器移动到第一个元素，并返回 true（如果容器中存在第一个元素）。如果 `First()` 返回 true，则可以通过 `Index()` 和 `Value()` 获取第一个元素的索引和值。修改了迭代器的状态。

#### (*Iterator[T]) Index 

``` go
func (iterator *Iterator[T]) Index() int
```

Index returns the current element's index. Does not modify the state of the iterator.

​	`Index` 返回当前元素的索引，不会修改迭代器的状态。

#### (*Iterator[T]) Next 

``` go
func (iterator *Iterator[T]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's index and value can be retrieved by Index() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

​	`Next` 将迭代器移动到下一个元素，并返回 true（如果容器中存在下一个元素）。如果 `Next()` 返回 true，则可以通过 `Index()` 和 `Value()` 获取下一个元素的索引和值。如果 `Next()` 是第一次调用，则迭代器将指向第一个元素（如果存在）。修改了迭代器的状态。

#### (*Iterator[T]) NextTo 

``` go
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's index and value can be retrieved by Index() and Value(). Modifies the state of the iterator.

​	`NextTo` 将迭代器移动到当前位置之后满足传入函数条件的下一个元素，并返回 true（如果容器中存在这样的元素）。如果 `NextTo()` 返回 true，则可以通过 `Index()` 和 `Value()` 获取该元素的索引和值。修改了迭代器的状态。

#### (*Iterator[T]) Value 

``` go
func (iterator *Iterator[T]) Value() T
```

Value returns the current element's value. Does not modify the state of the iterator.

​	`Value` 返回当前元素的值，不会修改迭代器的状态。

### type List 

``` go
type List[T comparable] struct {
	// contains filtered or unexported fields
}
```

List holds the elements, where each element points to the next element

​	`List` 保存元素，其中每个元素指向下一个元素。

#### func New 

``` go
func New[T comparable](values ...T) *List[T]
```

New instantiates a new list and adds the passed values, if any, to the list

​	`New` 实例化一个新列表，并添加传入的值（如果有）。

#### (*List[T]) Add 

``` go
func (list *List[T]) Add(values ...T)
```

Add appends a value (one or more) at the end of the list (same as Append())

​	`Add` 在列表末尾追加一个或多个值（等同于 `Append()`）。

#### (*List[T]) All 

``` go
func (list *List[T]) All(f func(index int, value T) bool) bool
```

All passes each element of the container to the given function and returns true if the function returns true for all elements.

​	`All` 将容器的每个元素传递给指定函数，如果函数对所有元素都返回 true，则返回 true。

#### (*List[T]) Any 

``` go
func (list *List[T]) Any(f func(index int, value T) bool) bool
```

Any passes each element of the container to the given function and returns true if the function ever returns true for any element.

​	`Any` 将容器的每个元素传递给指定函数，如果函数对任意元素返回 true，则返回 true。

#### (*List[T]) Append 

``` go
func (list *List[T]) Append(values ...T)
```

Append appends a value (one or more) at the end of the list (same as Add())

​	`Append` 在列表末尾追加一个或多个值（等同于 `Add()`）。

#### (*List[T]) Clear 

``` go
func (list *List[T]) Clear()
```

Clear removes all elements from the list.

​	`Clear` 移除列表中的所有元素。

#### (*List[T]) Contains 

``` go
func (list *List[T]) Contains(values ...T) bool
```

Contains checks if values (one or more) are present in the set. All values have to be present in the set for the method to return true. Performance time complexity of n^2. Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.

​	`Contains` 检查一个或多个值是否存在于集合中。所有值都必须存在于集合中时才返回 true。时间复杂度为 n^2。如果没有传入任何参数，则返回 true，因为集合始终是空集合的超集。

#### (*List[T]) Each 

``` go
func (list *List[T]) Each(f func(index int, value T))
```

Each calls the given function once for each element, passing that element's index and value.

​	`Each` 调用传入函数一次，传递每个元素的索引和值。

#### (*List[T]) Empty 

``` go
func (list *List[T]) Empty() bool
```

Empty returns true if list does not contain any elements.

​	`Empty` 如果列表中没有元素，则返回 true。

#### (*List[T]) Find 

``` go
func (list *List[T]) Find(f func(index int, value T) bool) (index int, value T)
```

Find passes each element of the container to the given function and returns the first (index,value) for which the function is true or -1,nil otherwise if no element matches the criteria.

​	`Find` 将每个元素传递给指定函数，返回第一个匹配条件的元素的索引和值；如果没有匹配的元素，则返回 -1 和 nil。

#### (*List[T]) FromJSON 

``` go
func (list *List[T]) FromJSON(data []byte) error
```

FromJSON populates list's elements from the input JSON representation.

​	`FromJSON` 使用输入 JSON 表示填充列表的元素。

#### (*List[T]) Get 

``` go
func (list *List[T]) Get(index int) (T, bool)
```

Get returns the element at index. Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.

​	`Get` 返回指定索引处的元素。如果索引在数组范围内且数组非空，则返回第二个参数为 true，否则返回 false。

#### (*List[T]) IndexOf 

``` go
func (list *List[T]) IndexOf(value T) int
```

IndexOf returns index of provided element

​	`IndexOf` 返回指定元素的索引。

#### (*List[T]) Insert 

``` go
func (list *List[T]) Insert(index int, values ...T)
```

Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right. Does not do anything if position is negative or bigger than list's size Note: position equal to list's size is valid, i.e. append.

​	`Insert` 在指定的索引位置插入值，并将该位置的值（如果有）及后续元素右移。如果索引为负或超出列表大小，则不执行任何操作。注意：索引等于列表大小是有效的，即追加。

#### (*List[T]) Iterator 

``` go
func (list *List[T]) Iterator() *Iterator[T]
```

Iterator returns a stateful iterator whose values can be fetched by an index.

​	`Iterator` 返回一个状态化的迭代器，其值可以通过索引获取。

#### (*List[T]) Map 

``` go
func (list *List[T]) Map(f func(index int, value T) T) *List[T]
```

Map invokes the given function once for each element and returns a container containing the values returned by the given function.

​	`Map` 调用传入函数一次，对每个元素执行操作，并返回一个包含函数返回值的容器。

#### (*List[T]) MarshalJSON 

``` go
func (list *List[T]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

​	`MarshalJSON` 实现了 `json.Marshaler` 接口。

#### (*List[T]) Prepend 

``` go
func (list *List[T]) Prepend(values ...T)
```

Prepend prepends a values (or more)

​	`Prepend` 在列表开头追加一个或多个值。

#### (*List[T]) Remove 

``` go
func (list *List[T]) Remove(index int)
```

Remove removes the element at the given index from the list.

​	`Remove` 移除指定索引处的元素。

#### (*List[T]) Select 

``` go
func (list *List[T]) Select(f func(index int, value T) bool) *List[T]
```

Select returns a new container containing all elements for which the given function returns a true value.

​	`Select` 返回一个新容器，包含传入函数返回值为 true 的所有元素。

#### (*List[T]) Set 

``` go
func (list *List[T]) Set(index int, value T)
```

Set value at specified index Does not do anything if position is negative or bigger than list's size Note: position equal to list's size is valid, i.e. append.

​	`Set` 在指定索引处设置值。如果索引为负或超出列表大小，则不执行任何操作。注意：索引等于列表大小是有效的，即追加。

#### (*List[T]) Size 

``` go
func (list *List[T]) Size() int
```

Size returns number of elements within the list.

​	`Size` 返回列表中的元素数量。

#### (*List[T]) Sort 

``` go
func (list *List[T]) Sort(comparator utils.Comparator[T])
```

Sort sort values (in-place) using.

​	`Sort` 使用指定的比较器对值进行排序（原地）。

#### (*List[T]) String 

``` go
func (list *List[T]) String() string
```

String returns a string representation of container

​	`String` 返回容器的字符串表示。

#### (*List[T]) Swap 

``` go
func (list *List[T]) Swap(i, j int)
```

Swap swaps values of two elements at the given indices.

​	`Swap` 交换指定位置的两个值。

#### (*List[T]) ToJSON 

``` go
func (list *List[T]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of list's elements.

​	`ToJSON` 输出列表元素的 JSON 表示。

#### (*List[T]) UnmarshalJSON 

``` go
func (list *List[T]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

​	`UnmarshalJSON` 实现了 `json.Unmarshaler` 接口。

#### (*List[T]) Values 

``` go
func (list *List[T]) Values() []T
```

Values returns all elements in the list.

​	`Values` 返回列表中的所有元素。
