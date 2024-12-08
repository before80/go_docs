+++
title = "btree"
date = 2024-12-07T11:09:51+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/trees/btree](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/trees/btree)
>
> 收录该文档时间： `2024-12-07T11:09:51+08:00`

## Overview 

Package btree implements a B tree.

​	包 btree 实现了一个 B 树。

According to Knuth's definition, a B-tree of order m is a tree which satisfies the following properties: 

​	根据 Knuth 的定义，m 阶 B 树满足以下性质：

- Every node has at most m children. 
  - 每个节点最多有 m 个子节点。
- Every non-leaf node (except root) has at least ⌈m/2⌉ children. 
  - 每个非叶子节点（除根节点外）至少有 ⌈m/2⌉ 个子节点。
- The root has at least two children if it is not a leaf node. 
  - 如果根节点不是叶子节点，则至少有两个子节点。
- A non-leaf node with k children contains k−1 keys. 
  - 拥有 k 个子节点的非叶子节点包含 k-1 个键。
- All leaves appear in the same level
  - 所有叶子节点都位于同一层。

Structure is not thread safe.

​	此结构是非线程安全的。

References: https://en.wikipedia.org/wiki/B-tree

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Entry 

``` go
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
```

Entry represents the key-value pair contained within nodes

​	`Entry` 表示节点中包含的键值对。

#### (*Entry[K, V]) String 

``` go
func (entry *Entry[K, V]) String() string
```

### type Iterator 

``` go
type Iterator[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Iterator holding the iterator's state

​	`Iterator` 保存迭代器的状态。

#### (*Iterator[K, V]) Begin 

``` go
func (iterator *Iterator[K, V]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

​	`Begin` 将迭代器重置为初始状态（第一个元素之前）。调用 `Next()` 获取第一个元素（如果有）。

#### (*Iterator[K, V]) End 

``` go
func (iterator *Iterator[K, V]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

​	`End` 将迭代器移动到最后一个元素之后（超出末尾）。调用 `Prev()` 获取最后一个元素（如果有）。

#### (*Iterator[K, V]) First 

``` go
func (iterator *Iterator[K, V]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator

​	`First` 将迭代器移动到第一个元素并返回 `true`（如果容器中有第一个元素）。如果返回 `true`，可以通过 `Key()` 和 `Value()` 获取第一个元素的键和值。修改迭代器的状态。

#### (*Iterator[K, V]) Key 

``` go
func (iterator *Iterator[K, V]) Key() K
```

Key returns the current element's key. Does not modify the state of the iterator.

​	`Key` 返回当前元素的键。不会修改迭代器的状态。

#### (*Iterator[K, V]) Last 

``` go
func (iterator *Iterator[K, V]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`Last` 将迭代器移动到最后一个元素并返回 `true`（如果容器中有最后一个元素）。如果返回 `true`，可以通过 `Key()` 和 `Value()` 获取最后一个元素的键和值。修改迭代器的状态。

#### (*Iterator[K, V]) Next 

``` go
func (iterator *Iterator[K, V]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's key and value can be retrieved by Key() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

​	`Next` 将迭代器移动到下一个元素并返回 `true`（如果容器中有下一个元素）。如果返回 `true`，可以通过 `Key()` 和 `Value()` 获取下一个元素的键和值。如果是第一次调用 `Next()`，则会将迭代器指向第一个元素（如果存在）。修改迭代器的状态。

#### (*Iterator[K, V]) NextTo 

``` go
func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`NextTo` 将迭代器移动到当前位置之后第一个满足条件的元素，并返回 `true`（如果容器中有满足条件的下一个元素）。如果返回 `true`，可以通过 `Key()` 和 `Value()` 获取下一个元素的键和值。修改迭代器的状态。

#### (*Iterator[K, V]) Node 

``` go
func (iterator *Iterator[K, V]) Node() *Node[K, V]
```

Node returns the current element's node. Does not modify the state of the iterator.

​	`Node` 返回当前元素的节点。不会修改迭代器的状态。

#### (*Iterator[K, V]) Prev 

``` go
func (iterator *Iterator[K, V]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`Prev` 将迭代器移动到上一个元素并返回 `true`（如果容器中有上一个元素）。如果返回 `true`，可以通过 `Key()` 和 `Value()` 获取上一个元素的键和值。修改迭代器的状态。

#### (*Iterator[K, V]) PrevTo 

``` go
func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`PrevTo` 将迭代器移动到当前位置之前第一个满足条件的元素，并返回 `true`（如果容器中有满足条件的上一个元素）。如果返回 `true`，可以通过 `Key()` 和 `Value()` 获取上一个元素的键和值。修改迭代器的状态。

#### (*Iterator[K, V]) Value 

``` go
func (iterator *Iterator[K, V]) Value() V
```

Value returns the current element's value. Does not modify the state of the iterator.

​	`Value` 返回当前元素的值。不会修改迭代器的状态。

### type Node 

``` go
type Node[K comparable, V any] struct {
	Parent   *Node[K, V]
	Entries  []*Entry[K, V] // Contained keys in node
	Children []*Node[K, V]  // Children nodes
}
```

Node is a single element within the tree

​	`Node` 表示树中的一个节点。

#### (*Node[K, V]) Size 

``` go
func (node *Node[K, V]) Size() int
```

Size returns the number of elements stored in the subtree. Computed dynamically on each call, i.e. the subtree is traversed to count the number of the nodes.

​	`Size` 返回子树中存储的元素数量。每次调用时动态计算，即遍历子树以统计节点数。

### type Tree 

``` go
type Tree[K comparable, V any] struct {
	Root       *Node[K, V]         // Root node
	Comparator utils.Comparator[K] // Key comparator
	// contains filtered or unexported fields
}
```

Tree holds elements of the B-tree

​	`Tree` 包含 B 树的元素。

#### func New 

``` go
func New[K cmp.Ordered, V any](order int) *Tree[K, V]
```

New instantiates a B-tree with the order (maximum number of children) and the built-in comparator for K

​	`New` 使用指定的阶（最大子节点数）和 `K` 的内置比较器实例化一个 B 树。

#### func NewWith 

``` go
func NewWith[K comparable, V any](order int, comparator utils.Comparator[K]) *Tree[K, V]
```

NewWith instantiates a B-tree with the order (maximum number of children) and a custom key comparator.

​	`NewWith` 使用指定的阶（最大子节点数）和自定义键比较器实例化一个 B 树。

#### (*Tree[K, V]) Clear 

``` go
func (tree *Tree[K, V]) Clear()
```

Clear removes all nodes from the tree.

​	`Clear` 移除树中的所有节点。

#### (*Tree[K, V]) Empty 

``` go
func (tree *Tree[K, V]) Empty() bool
```

Empty returns true if tree does not contain any nodes

​	`Empty` 如果树中没有任何节点，则返回 `true`。

#### (*Tree[K, V]) FromJSON 

``` go
func (tree *Tree[K, V]) FromJSON(data []byte) error
```

FromJSON populates the tree from the input JSON representation.

​	FromJSON 从输入的 JSON 表示中填充树。

#### (*Tree[K, V]) Get 

``` go
func (tree *Tree[K, V]) Get(key K) (value V, found bool)
```

Get searches the node in the tree by key and returns its value or nil if key is not found in tree. Second return parameter is true if key was found, otherwise false. Key should adhere to the comparator's type assertion, otherwise method panics.

​	Get 通过键在树中搜索节点，并返回其值，如果键在树中没有找到则返回 nil。第二个返回值是布尔值，表示是否找到了键。如果键未找到，返回 false；如果找到，返回 true。键应该符合比较器的类型断言，否则该方法会发生 panic。

#### (*Tree[K, V]) GetNode 

``` go
func (tree *Tree[K, V]) GetNode(key K) *Node[K, V]
```

GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree. Key should adhere to the comparator's type assertion, otherwise method panics.

​	GetNode 通过键在树中搜索节点，并返回其节点，如果键在树中没有找到则返回 nil。键应该符合比较器的类型断言，否则该方法会发生 panic。

#### (*Tree[K, V]) Height 

``` go
func (tree *Tree[K, V]) Height() int
```

Height returns the height of the tree.

​	Height 返回树的高度。

#### (*Tree[K, V]) Iterator 

``` go
func (tree *Tree[K, V]) Iterator() *Iterator[K, V]
```

Iterator returns a stateful iterator whose elements are key/value pairs.

​	Iterator 返回一个有状态的迭代器，其元素为键/值对。

#### (*Tree[K, V]) Keys 

``` go
func (tree *Tree[K, V]) Keys() []K
```

Keys returns all keys in-order

​	Keys 按顺序返回所有键。

#### (*Tree[K, V]) Left 

``` go
func (tree *Tree[K, V]) Left() *Node[K, V]
```

Left returns the left-most (min) node or nil if tree is empty.

​	Left 返回最左边（最小）的节点，如果树为空则返回 nil。

#### (*Tree[K, V]) LeftKey 

``` go
func (tree *Tree[K, V]) LeftKey() interface{}
```

LeftKey returns the left-most (min) key or nil if tree is empty.

​	LeftKey 返回最左边（最小）的键，如果树为空则返回 nil。

#### (*Tree[K, V]) LeftValue 

``` go
func (tree *Tree[K, V]) LeftValue() interface{}
```

LeftValue returns the left-most value or nil if tree is empty.

​	LeftValue 返回最左边的值，如果树为空则返回 nil。

#### (*Tree[K, V]) MarshalJSON 

``` go
func (tree *Tree[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### (*Tree[K, V]) Put 

``` go
func (tree *Tree[K, V]) Put(key K, value V)
```

Put inserts key-value pair node into the tree. If key already exists, then its value is updated with the new value. Key should adhere to the comparator's type assertion, otherwise method panics.

​	Put 将键值对节点插入树中。如果键已存在，则其值将更新为新值。键应该符合比较器的类型断言，否则该方法会发生 panic。

#### (*Tree[K, V]) Remove 

``` go
func (tree *Tree[K, V]) Remove(key K)
```

Remove remove the node from the tree by key. Key should adhere to the comparator's type assertion, otherwise method panics.

​	Remove 根据键从树中移除节点。键应该符合比较器的类型断言，否则该方法会发生 panic。

#### (*Tree[K, V]) Right 

``` go
func (tree *Tree[K, V]) Right() *Node[K, V]
```

Right returns the right-most (max) node or nil if tree is empty.

​	Right 返回最右边（最大）的节点，如果树为空则返回 nil。

#### (*Tree[K, V]) RightKey 

``` go
func (tree *Tree[K, V]) RightKey() interface{}
```

RightKey returns the right-most (max) key or nil if tree is empty.

​	RightKey 返回最右边（最大）的键，如果树为空则返回 nil。

#### (*Tree[K, V]) RightValue 

``` go
func (tree *Tree[K, V]) RightValue() interface{}
```

RightValue returns the right-most value or nil if tree is empty.

​	RightValue 返回最右边的值，如果树为空则返回 nil。

#### (*Tree[K, V]) Size 

``` go
func (tree *Tree[K, V]) Size() int
```

Size returns number of nodes in the tree.

​	Size 返回树中的节点数量。

#### (*Tree[K, V]) String 

``` go
func (tree *Tree[K, V]) String() string
```

String returns a string representation of container (for debugging purposes)

​	String 返回容器的字符串表示（用于调试目的）。

#### (*Tree[K, V]) ToJSON 

``` go
func (tree *Tree[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the tree.

​	ToJSON 输出树的 JSON 表示。

#### (*Tree[K, V]) UnmarshalJSON 

``` go
func (tree *Tree[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### (*Tree[K, V]) Values 

``` go
func (tree *Tree[K, V]) Values() []V
```

Values returns all values in-order based on the key.

​	Values 根据键按顺序返回所有值。
