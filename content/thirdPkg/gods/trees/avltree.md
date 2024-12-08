+++
title = "avltree"
date = 2024-12-07T11:09:26+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/trees/avltree](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/trees/avltree)
>
> 收录该文档时间： `2024-12-07T11:09:26+08:00`

## Overview 

Package avltree implements an AVL balanced binary tree.

​	包 `avltree` 实现了一个 AVL 平衡二叉树。

Structure is not thread safe.

​	结构体不是线程安全的。

References: https://en.wikipedia.org/wiki/AVL_tree

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

### type Iterator 

``` go
type Iterator[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Iterator holding the iterator's state

​	保存迭代器状态的 `Iterator`

#### (*Iterator[K, V]) Begin 

``` go
func (iterator *Iterator[K, V]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

​	`Begin` 将迭代器重置到其初始状态（第一个元素之前）。调用 `Next()` 来获取第一个元素（如果存在）。

#### (*Iterator[K, V]) End 

``` go
func (iterator *Iterator[K, V]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

​	`End` 将迭代器移到最后一个元素之后（超出范围）。调用 `Prev()` 来获取最后一个元素（如果存在）。

#### (*Iterator[K, V]) First 

``` go
func (iterator *Iterator[K, V]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator

​	`First` 将迭代器移到第一个元素，并返回 `true` 如果容器中存在第一个元素。若返回 `true`，可以通过 `Key()` 和 `Value()` 获取第一个元素的键和值。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) Key 

``` go
func (iterator *Iterator[K, V]) Key() (k K)
```

Key returns the current element's key. Does not modify the state of the iterator.

​	`Key` 返回当前元素的键。此方法不会修改迭代器的状态。

#### (*Iterator[K, V]) Last 

``` go
func (iterator *Iterator[K, V]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`Last` 将迭代器移到最后一个元素，并返回 `true` 如果容器中存在最后一个元素。若返回 `true`，可以通过 `Key()` 和 `Value()` 获取最后一个元素的键和值。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) Next 

``` go
func (iterator *Iterator[K, V]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's key and value can be retrieved by Key() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

​	`Next` 将迭代器移到下一个元素，并返回 `true` 如果容器中存在下一个元素。若返回 `true`，可以通过 `Key()` 和 `Value()` 获取下一个元素的键和值。首次调用 `Next()` 时，会指向第一个元素（如果存在）。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) NextTo 

``` go
func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`NextTo` 将迭代器移到当前元素位置之后满足传入条件的下一个元素，并返回 `true` 如果容器中存在此类元素。若返回 `true`，可以通过 `Key()` 和 `Value()` 获取该元素的键和值。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) Node 

``` go
func (iterator *Iterator[K, V]) Node() *Node[K, V]
```

Node returns the current element's node. Does not modify the state of the iterator.

​	`Node` 返回当前元素的节点。此方法不会修改迭代器的状态。

#### (*Iterator[K, V]) Prev 

``` go
func (iterator *Iterator[K, V]) Prev() bool
```

Prev moves the iterator to the next element and returns true if there was a previous element in the container. If Prev() returns true, then next element's key and value can be retrieved by Key() and Value(). If Prev() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

​	`Prev` 将迭代器移到前一个元素，并返回 `true` 如果容器中存在前一个元素。若返回 `true`，可以通过 `Key()` 和 `Value()` 获取前一个元素的键和值。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) PrevTo 

``` go
func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

​	`PrevTo` 将迭代器移到当前元素位置之前满足传入条件的上一个元素，并返回 `true` 如果容器中存在此类元素。若返回 `true`，可以通过 `Key()` 和 `Value()` 获取该元素的键和值。此方法会修改迭代器的状态。

#### (*Iterator[K, V]) Value 

``` go
func (iterator *Iterator[K, V]) Value() (v V)
```

Value returns the current element's value. Does not modify the state of the iterator.

​	`Value` 返回当前元素的值。此方法不会修改迭代器的状态。

### type Node 

``` go
type Node[K comparable, V any] struct {
	Key      K
	Value    V
	Parent   *Node[K, V]    // Parent node
	Children [2]*Node[K, V] // Children nodes
	// contains filtered or unexported fields
}
```

Node is a single element within the tree

​	树中的单个节点

#### (*Node[K, V]) Next 

``` go
func (n *Node[K, V]) Next() *Node[K, V]
```

Next returns the next element in an inorder walk of the AVL tree.

​	`Next` 返回 AVL 树中中序遍历的下一个节点。

#### (*Node[K, V]) Prev 

``` go
func (n *Node[K, V]) Prev() *Node[K, V]
```

Prev returns the previous element in an inorder walk of the AVL tree.

​	`Prev` 返回 AVL 树中中序遍历的前一个节点。

#### (*Node[K, V]) Size 

``` go
func (n *Node[K, V]) Size() int
```

Size returns the number of elements stored in the subtree. Computed dynamically on each call, i.e. the subtree is traversed to count the number of the nodes.

​	`Size` 返回子树中存储的元素数量。此方法每次调用都会动态计算（遍历子树以统计节点数量）。

#### (*Node[K, V]) String 

``` go
func (n *Node[K, V]) String() string
```

### type Tree 

``` go
type Tree[K comparable, V any] struct {
	Root       *Node[K, V]         // Root node
	Comparator utils.Comparator[K] // Key comparator
	// contains filtered or unexported fields
}
```

Tree holds elements of the AVL tree.

​	AVL 树的存储结构

#### func New 

``` go
func New[K cmp.Ordered, V any]() *Tree[K, V]
```

New instantiates an AVL tree with the built-in comparator for K

​	`New` 使用内置比较器为 `K` 实例化一个新的 AVL 树。

#### func NewWith 

``` go
func NewWith[K comparable, V any](comparator utils.Comparator[K]) *Tree[K, V]
```

NewWith instantiates an AVL tree with the custom comparator.

​	`NewWith` 使用自定义比较器实例化一个新的 AVL 树。

#### (*Tree[K, V]) Ceiling 

``` go
func (tree *Tree[K, V]) Ceiling(key K) (floor *Node[K, V], found bool)
```

Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found. Second return parameter is true if ceiling was found, otherwise false.

​	`Ceiling` 查找输入键的天花板节点，返回天花板节点或 `nil`（如果未找到天花板）。第二个返回值为 `true` 表示找到了天花板节点，否则为 `false`。

Ceiling node is defined as the smallest node that is larger than or equal to the given node. A ceiling node may not be found, either because the tree is empty, or because all nodes in the tree is smaller than the given node.

​	天花板节点定义为大于或等于给定键的最小节点。如果树为空，或者树中所有节点均小于给定键，则可能找不到天花板节点。

Key should adhere to the comparator's type assertion, otherwise method panics.

​	键应符合比较器的类型断言，否则方法将触发 panic。

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

Empty returns true if tree does not contain any nodes.

​	`Empty` 返回 `true` 如果树中不包含任何节点。

#### (*Tree[K, V]) Floor 

``` go
func (tree *Tree[K, V]) Floor(key K) (floor *Node[K, V], found bool)
```

Floor Finds floor node of the input key, return the floor node or nil if no floor is found. Second return parameter is true if floor was found, otherwise false.

​	`Floor` 查找输入键的地板节点，返回地板节点或 `nil`（如果未找到地板）。第二个返回值为 `true` 表示找到了地板节点，否则为 `false`。

Floor node is defined as the largest node that is smaller than or equal to the given node. A floor node may not be found, either because the tree is empty, or because all nodes in the tree is larger than the given node.

​	地板节点定义为小于或等于给定键的最大节点。如果树为空，或者树中所有节点均大于给定键，则可能找不到地板节点。

Key should adhere to the comparator's type assertion, otherwise method panics.

​	键应符合比较器的类型断言，否则方法将触发 panic。

#### (*Tree[K, V]) FromJSON 

``` go
func (tree *Tree[K, V]) FromJSON(data []byte) error
```

FromJSON populates the tree from the input JSON representation.

​	`FromJSON` 从输入的 JSON 表示中填充树。

#### (*Tree[K, V]) Get 

``` go
func (tree *Tree[K, V]) Get(key K) (value V, found bool)
```

Get searches the node in the tree by key and returns its value or nil if key is not found in tree. Second return parameter is true if key was found, otherwise false. Key should adhere to the comparator's type assertion, otherwise method panics.

​	`Get` 根据键在树中搜索节点，返回其值或 `nil`（如果键未在树中找到）。第二个返回值为 `true` 表示找到了键，否则为 `false`。键应符合比较器的类型断言，否则方法将触发 panic。

#### (*Tree[K, V]) GetNode 

``` go
func (tree *Tree[K, V]) GetNode(key K) *Node[K, V]
```

GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree. Key should adhere to the comparator's type assertion, otherwise method panics.

​	`GetNode` 根据键在树中搜索节点，返回其节点或 `nil`（如果键未在树中找到）。键应符合比较器的类型断言，否则方法将触发 panic。

#### (*Tree[K, V]) Iterator 

``` go
func (tree *Tree[K, V]) Iterator() *Iterator[K, V]
```

Iterator returns a stateful iterator whose elements are key/value pairs.

​	`Iterator` 返回一个状态化的迭代器，其元素为键/值对。

#### (*Tree[K, V]) Keys 

``` go
func (tree *Tree[K, V]) Keys() []K
```

Keys returns all keys in-order

​	`Keys` 返回按顺序排列的所有键。

#### (*Tree[K, V]) Left 

``` go
func (tree *Tree[K, V]) Left() *Node[K, V]
```

Left returns the minimum element of the AVL tree or nil if the tree is empty.

​	`Left` 返回 AVL 树中的最小元素或 `nil`（如果树为空）。

#### (*Tree[K, V]) MarshalJSON 

``` go
func (tree *Tree[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

​	`MarshalJSON` 实现了 `json.Marshaler` 接口。

#### (*Tree[K, V]) Put 

``` go
func (tree *Tree[K, V]) Put(key K, value V)
```

Put inserts node into the tree. Key should adhere to the comparator's type assertion, otherwise method panics.

​	`Put` 将节点插入到树中。键应符合比较器的类型断言，否则方法将触发 panic。

#### (*Tree[K, V]) Remove 

``` go
func (tree *Tree[K, V]) Remove(key K)
```

Remove remove the node from the tree by key. Key should adhere to the comparator's type assertion, otherwise method panics.

​	`Remove` 根据键从树中移除节点。键应符合比较器的类型断言，否则方法将触发 panic。

#### (*Tree[K, V]) Right 

``` go
func (tree *Tree[K, V]) Right() *Node[K, V]
```

Right returns the maximum element of the AVL tree or nil if the tree is empty.

​	`Right` 返回 AVL 树中的最大元素或 `nil`（如果树为空）。

#### (*Tree[K, V]) Size 

``` go
func (tree *Tree[K, V]) Size() int
```

Size returns the number of elements stored in the tree.

​	`Size` 返回树中存储的元素数量。

#### (*Tree[K, V]) String 

``` go
func (tree *Tree[K, V]) String() string
```

String returns a string representation of container

​	`String` 返回容器的字符串表示。

#### (*Tree[K, V]) ToJSON 

``` go
func (tree *Tree[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the tree.

​	`ToJSON` 输出树的 JSON 表示。

#### (*Tree[K, V]) UnmarshalJSON 

``` go
func (tree *Tree[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

​	`UnmarshalJSON` 实现了 `json.Unmarshaler` 接口。

#### (*Tree[K, V]) Values 

``` go
func (tree *Tree[K, V]) Values() []V
```

Values returns all values in-order based on the key.

​	`Values` 返回根据键的顺序排列的所有值。
