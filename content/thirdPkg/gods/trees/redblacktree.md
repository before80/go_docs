+++
title = "redblacktree"
date = 2024-12-07T11:10:05+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/trees/redblacktree](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/trees/redblacktree)
>
> 收录该文档时间： `2024-12-07T11:10:05+08:00`

## Overview 

Package redblacktree implements a red-black tree.

Used by TreeSet and TreeMap.

Structure is not thread safe.

References: [http://en.wikipedia.org/wiki/Red%E2%80%93black_tree](http://en.wikipedia.org/wiki/Red–black_tree)

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

This section is empty.

## 类型 

#### type Iterator 

``` go
type Iterator[K comparable, V any] struct {
	// contains filtered or unexported fields
}
```

Iterator holding the iterator's state

#### func (*Iterator[K, V]) [Begin](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L136) 

``` go
func (iterator *Iterator[K, V]) Begin()
```

Begin resets the iterator to its initial state (one-before-first) Call Next() to fetch the first element if any.

#### func (*Iterator[K, V]) [End](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L143) 

``` go
func (iterator *Iterator[K, V]) End()
```

End moves the iterator past the last element (one-past-the-end). Call Prev() to fetch the last element if any.

#### func (*Iterator[K, V]) [First](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L151) 

``` go
func (iterator *Iterator[K, V]) First() bool
```

First moves the iterator to the first element and returns true if there was a first element in the container. If First() returns true, then first element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator

#### func (*Iterator[K, V]) [Key](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L124) 

``` go
func (iterator *Iterator[K, V]) Key() K
```

Key returns the current element's key. Does not modify the state of the iterator.

#### func (*Iterator[K, V]) [Last](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L159) 

``` go
func (iterator *Iterator[K, V]) Last() bool
```

Last moves the iterator to the last element and returns true if there was a last element in the container. If Last() returns true, then last element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Next](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L39) 

``` go
func (iterator *Iterator[K, V]) Next() bool
```

Next moves the iterator to the next element and returns true if there was a next element in the container. If Next() returns true, then next element's key and value can be retrieved by Key() and Value(). If Next() was called for the first time, then it will point the iterator to the first element if it exists. Modifies the state of the iterator.

#### func (*Iterator[K, V]) [NextTo](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L168) 

``` go
func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool
```

NextTo moves the iterator to the next element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Node](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L130) 

``` go
func (iterator *Iterator[K, V]) Node() *Node[K, V]
```

Node returns the current element's node. Does not modify the state of the iterator.

#### func (*Iterator[K, V]) [Prev](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L79) 

``` go
func (iterator *Iterator[K, V]) Prev() bool
```

Prev moves the iterator to the previous element and returns true if there was a previous element in the container. If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [PrevTo](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L182) 

``` go
func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool
```

PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the passed function, and returns true if there was a next element in the container. If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.

#### func (*Iterator[K, V]) [Value](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L118) 

``` go
func (iterator *Iterator[K, V]) Value() V
```

Value returns the current element's value. Does not modify the state of the iterator.

#### type Node 

``` go
type Node[K comparable, V any] struct {
	Key   K
	Value V

	Left   *Node[K, V]
	Right  *Node[K, V]
	Parent *Node[K, V]
	// contains filtered or unexported fields
}
```

Node is a single element within the tree

#### func (*Node[K, V]) [Size](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L162) 

``` go
func (node *Node[K, V]) Size() int
```

Size returns the number of elements stored in the subtree. Computed dynamically on each call, i.e. the subtree is traversed to count the number of the nodes.

#### func (*Node[K, V]) [String](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L291) 

``` go
func (node *Node[K, V]) String() string
```

#### type Tree 

``` go
type Tree[K comparable, V any] struct {
	Root *Node[K, V]

	Comparator utils.Comparator[K]
	// contains filtered or unexported fields
}
```

Tree holds elements of the red-black tree

#### func New 

``` go
func New[K cmp.Ordered, V any]() *Tree[K, V]
```

New instantiates a red-black tree with the built-in comparator for K

#### func NewWith 

``` go
func NewWith[K comparable, V any](comparator utils.Comparator[K]) *Tree[K, V]
```

NewWith instantiates a red-black tree with the custom comparator.

#### func (*Tree[K, V]) [Ceiling](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L255) 

``` go
func (tree *Tree[K, V]) Ceiling(key K) (ceiling *Node[K, V], found bool)
```

Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found. Second return parameter is true if ceiling was found, otherwise false.

Ceiling node is defined as the smallest node that is larger than or equal to the given node. A ceiling node may not be found, either because the tree is empty, or because all nodes in the tree are smaller than the given node.

Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Tree[K, V]) [Clear](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L277) 

``` go
func (tree *Tree[K, V]) Clear()
```

Clear removes all nodes from the tree.

#### func (*Tree[K, V]) [Empty](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L151) 

``` go
func (tree *Tree[K, V]) Empty() bool
```

Empty returns true if tree does not contain any nodes

#### func (*Tree[K, V]) [Floor](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L226) 

``` go
func (tree *Tree[K, V]) Floor(key K) (floor *Node[K, V], found bool)
```

Floor Finds floor node of the input key, return the floor node or nil if no floor is found. Second return parameter is true if floor was found, otherwise false.

Floor node is defined as the largest node that is smaller than or equal to the given node. A floor node may not be found, either because the tree is empty, or because all nodes in the tree are larger than the given node.

Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Tree[K, V]) [FromJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/serialization.go#L28) 

``` go
func (tree *Tree[K, V]) FromJSON(data []byte) error
```

FromJSON populates the tree from the input JSON representation.

#### func (*Tree[K, V]) [Get](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L104) 

``` go
func (tree *Tree[K, V]) Get(key K) (value V, found bool)
```

Get searches the node in the tree by key and returns its value or nil if key is not found in tree. Second return parameter is true if key was found, otherwise false. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Tree[K, V]) [GetNode](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L114) 

``` go
func (tree *Tree[K, V]) GetNode(key K) *Node[K, V]
```

GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Tree[K, V]) [Iterator](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L26) 

``` go
func (tree *Tree[K, V]) Iterator() *Iterator[K, V]
```

Iterator returns a stateful iterator whose elements are key/value pairs.

#### func (*Tree[K, V]) [IteratorAt](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/iterator.go#L31) 

``` go
func (tree *Tree[K, V]) IteratorAt(node *Node[K, V]) *Iterator[K, V]
```

IteratorAt returns a stateful iterator whose elements are key/value pairs that is initialised at a particular node.

#### func (*Tree[K, V]) [Keys](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L177) 

``` go
func (tree *Tree[K, V]) Keys() []K
```

Keys returns all keys in-order

#### func (*Tree[K, V]) [Left](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L197) 

``` go
func (tree *Tree[K, V]) Left() *Node[K, V]
```

Left returns the left-most (min) node or nil if tree is empty.

#### func (*Tree[K, V]) [MarshalJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/serialization.go#L46) 

``` go
func (tree *Tree[K, V]) MarshalJSON() ([]byte, error)
```

MarshalJSON @implements json.Marshaler

#### func (*Tree[K, V]) [Put](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L60) 

``` go
func (tree *Tree[K, V]) Put(key K, value V)
```

Put inserts node into the tree. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Tree[K, V]) [Remove](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L120) 

``` go
func (tree *Tree[K, V]) Remove(key K)
```

Remove remove the node from the tree by key. Key should adhere to the comparator's type assertion, otherwise method panics.

#### func (*Tree[K, V]) [Right](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L208) 

``` go
func (tree *Tree[K, V]) Right() *Node[K, V]
```

Right returns the right-most (max) node or nil if tree is empty.

#### func (*Tree[K, V]) [Size](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L156) 

``` go
func (tree *Tree[K, V]) Size() int
```

Size returns number of nodes in the tree.

#### func (*Tree[K, V]) [String](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L283) 

``` go
func (tree *Tree[K, V]) String() string
```

String returns a string representation of container

#### func (*Tree[K, V]) [ToJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/serialization.go#L18) 

``` go
func (tree *Tree[K, V]) ToJSON() ([]byte, error)
```

ToJSON outputs the JSON representation of the tree.

#### func (*Tree[K, V]) [UnmarshalJSON](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/serialization.go#L41) 

``` go
func (tree *Tree[K, V]) UnmarshalJSON(bytes []byte) error
```

UnmarshalJSON @implements json.Unmarshaler

#### func (*Tree[K, V]) [Values](https://github.com/emirpasic/gods/blob/v2.0.0-alpha/trees/redblacktree/redblacktree.go#L187) 

``` go
func (tree *Tree[K, V]) Values() []V
```

Values returns all values in-order based on the key.
