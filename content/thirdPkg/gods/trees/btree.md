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

According to Knuth's definition, a B-tree of order m is a tree which satisfies the following properties: - Every node has at most m children. - Every non-leaf node (except root) has at least ⌈m/2⌉ children. - The root has at least two children if it is not a leaf node. - A non-leaf node with k children contains k−1 keys. - All leaves appear in the same level

Structure is not thread safe.

References: https://en.wikipedia.org/wiki/B-tree

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

#### type Entry 

``` go
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
```

Entry represents the key-value pair contained within nodes

#### func (*Entry[K, V]) [String](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L217) 

``` go
func (entry *Entry[K, V]) String() string
```

#### type Iterator 

``` go
type Iterator[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Iterator holding the iterator's state

#### func (*Iterator[K, V]) [Begin](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L170) 

``` go
func (iterator *Iterator[K, V]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

#### func (*Iterator[K, V]) [End](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L178) 

``` go
func (iterator *Iterator[K, V]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

#### func (*Iterator[K, V]) [First](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L187) 

``` go
func (iterator *Iterator[K, V]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator

#### func (*Iterator[K, V]) [Key](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L158) 

``` go
func (iterator *Iterator[K, V]) Key() K
```

Key returns the current element's key. Does not modify the state of the iterator.

#### func (*Iterator[K, V]) [Last](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L195) 

``` go
func (iterator *Iterator[K, V]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Next](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L35) 

``` go
func (iterator *Iterator[K, V]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's key and value can be retrieved by Key() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

#### func (*Iterator[K, V]) [NextTo](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L204) 

``` go
func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Node](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L164) 

``` go
func (iterator *Iterator[K, V]) Node() *Node[K, V]
```

Node returns the current element's node. Does not modify the state of the iterator.

#### func (*Iterator[K, V]) [Prev](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L94) 

``` go
func (iterator *Iterator[K, V]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [PrevTo](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L218) 

``` go
func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Value](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L152) 

``` go
func (iterator *Iterator[K, V]) Value() V
```

Value returns the current element's value. Does not modify the state of the iterator.

#### type Node 

``` go
type Node[K comparable, V any] struct {
	Parent   *Node[K, V]
	Entries  []*Entry[K, V] // Contained keys in node
	Children []*Node[K, V]  // Children nodes
}
```

Node is a single element within the tree

#### func (*Node[K, V]) [Size](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L123) 

``` go
func (node *Node[K, V]) Size() int
```

Size returns the number of elements stored in the subtree. Computed dynamically on each call, i.e. the subtree is traversed to count the number of the nodes.

#### type Tree 

``` go
type Tree[K comparable, V any] struct {
	Root       *Node[K, V]         // Root node
	Comparator utils.Comparator[K] // Key comparator
	// contains filtered or unexported fields
}
```

Tree holds elements of the B-tree

#### func New 

``` go
func New[K cmp.Ordered, V any](order int) *Tree[K, V]
```

New instantiates a B-tree with the order (maximum number of children) and the built-in comparator for K

#### func NewWith 

``` go
func NewWith[K comparable, V any](order int, comparator utils.Comparator[K]) *Tree[K, V]
```

NewWith instantiates a B-tree with the order (maximum number of children) and a custom key comparator.

#### func (*Tree[K, V]) [Clear](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L155) 

``` go
func (tree *Tree[K, V]) Clear()
```

Clear removes all nodes from the tree.

#### func (*Tree[K, V]) [Empty](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L112) 

``` go
func (tree *Tree[K, V]) Empty() bool
```

Empty returns true if tree does not contain any nodes

#### func (*Tree[K, V]) [FromJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/serialization.go#L28) 

``` go
func (tree *Tree[K, V]) FromJSON(data []byte) error
```

FromJSON populates the tree from the input JSON representation.

#### func (*Tree[K, V]) [Get](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L86) 

``` go
func (tree *Tree[K, V]) Get(key K) (value V, found bool)
```

Get searches the node in the tree by key and returns its value or nil if key is not found in tree. Second return parameter is true if key was found, otherwise false. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Tree[K, V]) [GetNode](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L96) 

``` go
func (tree *Tree[K, V]) GetNode(key K) *Node[K, V]
```

GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Tree[K, V]) [Height](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L161) 

``` go
func (tree *Tree[K, V]) Height() int
```

Height returns the height of the tree.

#### func (*Tree[K, V]) [Iterator](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/iterator.go#L27) 

``` go
func (tree *Tree[K, V]) Iterator() *Iterator[K, V]
```

Iterator returns a stateful iterator whose elements are key/value pairs.

#### func (*Tree[K, V]) [Keys](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L135) 

``` go
func (tree *Tree[K, V]) Keys() []K
```

Keys returns all keys in-order

#### func (*Tree[K, V]) [Left](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L166) 

``` go
func (tree *Tree[K, V]) Left() *Node[K, V]
```

Left returns the left-most (min) node or nil if tree is empty.

#### func (*Tree[K, V]) [LeftKey](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L171) 

``` go
func (tree *Tree[K, V]) LeftKey() interface{}
```

LeftKey returns the left-most (min) key or nil if tree is empty.

#### func (*Tree[K, V]) [LeftValue](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L179) 

``` go
func (tree *Tree[K, V]) LeftValue() interface{}
```

LeftValue returns the left-most value or nil if tree is empty.

#### func (*Tree[K, V]) [MarshalJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/serialization.go#L49) 

``` go
func (tree *Tree[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### func (*Tree[K, V]) [Put](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L69) 

``` go
func (tree *Tree[K, V]) Put(key K, value V)
```

Put inserts key-value pair node into the tree. If key already exists, then its value is updated with the new value. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Tree[K, V]) [Remove](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L103) 

``` go
func (tree *Tree[K, V]) Remove(key K)
```

Remove remove the node from the tree by key. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Tree[K, V]) [Right](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L187) 

``` go
func (tree *Tree[K, V]) Right() *Node[K, V]
```

Right returns the right-most (max) node or nil if tree is empty.

#### func (*Tree[K, V]) [RightKey](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L192) 

``` go
func (tree *Tree[K, V]) RightKey() interface{}
```

RightKey returns the right-most (max) key or nil if tree is empty.

#### func (*Tree[K, V]) [RightValue](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L200) 

``` go
func (tree *Tree[K, V]) RightValue() interface{}
```

RightValue returns the right-most value or nil if tree is empty.

#### func (*Tree[K, V]) [Size](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L117) 

``` go
func (tree *Tree[K, V]) Size() int
```

Size returns number of nodes in the tree.

#### func (*Tree[K, V]) [String](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L208) 

``` go
func (tree *Tree[K, V]) String() string
```

String returns a string representation of container (for debugging purposes)

#### func (*Tree[K, V]) [ToJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/serialization.go#L18) 

``` go
func (tree *Tree[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the tree.

#### func (*Tree[K, V]) [UnmarshalJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/serialization.go#L44) 

``` go
func (tree *Tree[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### func (*Tree[K, V]) [Values](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/btree/btree.go#L145) 

``` go
func (tree *Tree[K, V]) Values() []V
```

Values returns all values in-order based on the key.
