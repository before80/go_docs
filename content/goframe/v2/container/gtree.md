+++
title = "gtree"
date = 2024-03-21T17:45:11+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gtree

Package gtree provides concurrent-safe/unsafe tree containers.

Some implements are from: https://github.com/emirpasic/gods

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type AVLTree 

``` go
type AVLTree struct {
	// contains filtered or unexported fields
}
```

AVLTree holds elements of the AVL tree.

##### func NewAVLTree 

``` go
func NewAVLTree(comparator func(v1, v2 interface{}) int, safe ...bool) *AVLTree
```

NewAVLTree instantiates an AVL tree with the custom key comparator. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewAVLTreeFrom 

``` go
func NewAVLTreeFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *AVLTree
```

NewAVLTreeFrom instantiates an AVL tree with the custom key comparator and data map. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

##### Example

``` go
```
##### (*AVLTree) Ceiling 

``` go
func (tree *AVLTree) Ceiling(key interface{}) (ceiling *AVLTreeNode, found bool)
```

Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling node is found. Second return parameter is true if ceiling was found, otherwise false.

Ceiling node is defined as the smallest node that is larger than or equal to the given node. A ceiling node may not be found, either because the tree is empty, or because all nodes in the tree is smaller than the given node.

Key should adhere to the comparator's type assertion, otherwise method panics.

##### Example

``` go
```
##### (*AVLTree) Clear 

``` go
func (tree *AVLTree) Clear()
```

Clear removes all nodes from the tree.

##### Example

``` go
```
##### (*AVLTree) Clone 

``` go
func (tree *AVLTree) Clone() *AVLTree
```

Clone returns a new tree with a copy of current tree.

##### Example

``` go
```
##### (*AVLTree) Contains 

``` go
func (tree *AVLTree) Contains(key interface{}) bool
```

Contains checks whether `key` exists in the tree.

##### Example

``` go
```
##### (*AVLTree) Flip 

``` go
func (tree *AVLTree) Flip(comparator ...func(v1, v2 interface{}) int)
```

Flip exchanges key-value of the tree to value-key. Note that you should guarantee the value is the same type as key, or else the comparator would panic.

If the type of value is different with key, you pass the new `comparator`.

##### Example

``` go
```
##### (*AVLTree) Floor 

``` go
func (tree *AVLTree) Floor(key interface{}) (floor *AVLTreeNode, found bool)
```

Floor Finds floor node of the input key, return the floor node or nil if no floor node is found. Second return parameter is true if floor was found, otherwise false.

Floor node is defined as the largest node that is smaller than or equal to the given node. A floor node may not be found, either because the tree is empty, or because all nodes in the tree is larger than the given node.

Key should adhere to the comparator's type assertion, otherwise method panics.

##### Example

``` go
```
##### (*AVLTree) Get 

``` go
func (tree *AVLTree) Get(key interface{}) (value interface{})
```

Get searches the node in the tree by `key` and returns its value or nil if key is not found in tree.

##### Example

``` go
```
##### (*AVLTree) GetOrSet 

``` go
func (tree *AVLTree) GetOrSet(key interface{}, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*AVLTree) GetOrSetFunc 

``` go
func (tree *AVLTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*AVLTree) GetOrSetFuncLock 

``` go
func (tree *AVLTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*AVLTree) GetVar 

``` go
func (tree *AVLTree) GetVar(key interface{}) *gvar.Var
```

GetVar returns a gvar.Var with the value by given `key`. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*AVLTree) GetVarOrSet 

``` go
func (tree *AVLTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var
```

GetVarOrSet returns a gvar.Var with result from GetVarOrSet. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*AVLTree) GetVarOrSetFunc 

``` go
func (tree *AVLTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*AVLTree) GetVarOrSetFuncLock 

``` go
func (tree *AVLTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*AVLTree) IsEmpty 

``` go
func (tree *AVLTree) IsEmpty() bool
```

IsEmpty returns true if tree does not contain any nodes.

##### Example

``` go
```
##### (*AVLTree) Iterator 

``` go
func (tree *AVLTree) Iterator(f func(key, value interface{}) bool)
```

Iterator is alias of IteratorAsc.

##### Example

``` go
```
##### (*AVLTree) IteratorAsc 

``` go
func (tree *AVLTree) IteratorAsc(f func(key, value interface{}) bool)
```

IteratorAsc iterates the tree readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*AVLTree) IteratorAscFrom 

``` go
func (tree *AVLTree) IteratorAscFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorAscFrom iterates the tree readonly in ascending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

##### (*AVLTree) IteratorDesc 

``` go
func (tree *AVLTree) IteratorDesc(f func(key, value interface{}) bool)
```

IteratorDesc iterates the tree readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*AVLTree) IteratorDescFrom 

``` go
func (tree *AVLTree) IteratorDescFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorDescFrom iterates the tree readonly in descending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*AVLTree) IteratorFrom 

``` go
func (tree *AVLTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorFrom is alias of IteratorAscFrom.

##### Example

``` go
```
##### (*AVLTree) Keys 

``` go
func (tree *AVLTree) Keys() []interface{}
```

Keys returns all keys in asc order.

##### Example

``` go
```
##### (*AVLTree) Left 

``` go
func (tree *AVLTree) Left() *AVLTreeNode
```

Left returns the minimum element of the AVL tree or nil if the tree is empty.

##### Example

``` go
```
##### (*AVLTree) Map 

``` go
func (tree *AVLTree) Map() map[interface{}]interface{}
```

Map returns all key-value items as map.

##### Example

``` go
```
##### (*AVLTree) MapStrAny 

``` go
func (tree *AVLTree) MapStrAny() map[string]interface{}
```

MapStrAny returns all key-value items as map[string]interface{}.

##### Example

``` go
```
##### (AVLTree) MarshalJSON 

``` go
func (tree AVLTree) MarshalJSON() (jsonBytes []byte, err error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*AVLTree) Print 

``` go
func (tree *AVLTree) Print()
```

Print prints the tree to stdout.

##### Example

``` go
```
##### (*AVLTree) Remove 

``` go
func (tree *AVLTree) Remove(key interface{}) (value interface{})
```

Remove removes the node from the tree by key. Key should adhere to the comparator's type assertion, otherwise method panics.

##### Example

``` go
```
##### (*AVLTree) Removes 

``` go
func (tree *AVLTree) Removes(keys []interface{})
```

Removes batch deletes values of the tree by `keys`.

##### Example

``` go
```
##### (*AVLTree) Replace 

``` go
func (tree *AVLTree) Replace(data map[interface{}]interface{})
```

Replace the data of the tree with given `data`.

##### Example

``` go
```
##### (*AVLTree) Right 

``` go
func (tree *AVLTree) Right() *AVLTreeNode
```

Right returns the maximum element of the AVL tree or nil if the tree is empty.

##### Example

``` go
```
##### (*AVLTree) Search 

``` go
func (tree *AVLTree) Search(key interface{}) (value interface{}, found bool)
```

Search searches the tree with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### Example

``` go
```
##### (*AVLTree) Set 

``` go
func (tree *AVLTree) Set(key interface{}, value interface{})
```

Set inserts node into the tree.

##### Example

``` go
```
##### (*AVLTree) SetIfNotExist 

``` go
func (tree *AVLTree) SetIfNotExist(key interface{}, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*AVLTree) SetIfNotExistFunc 

``` go
func (tree *AVLTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*AVLTree) SetIfNotExistFuncLock 

``` go
func (tree *AVLTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*AVLTree) Sets 

``` go
func (tree *AVLTree) Sets(data map[interface{}]interface{})
```

Sets batch sets key-values to the tree.

##### Example

``` go
```
##### (*AVLTree) Size 

``` go
func (tree *AVLTree) Size() int
```

Size returns number of nodes in the tree.

##### Example

``` go
```
##### (*AVLTree) String 

``` go
func (tree *AVLTree) String() string
```

String returns a string representation of container

##### Example

``` go
```
##### (*AVLTree) Values 

``` go
func (tree *AVLTree) Values() []interface{}
```

Values returns all values in asc order based on the key.

##### Example

``` go
```
#### type AVLTreeNode 

``` go
type AVLTreeNode struct {
	Key   interface{}
	Value interface{}
	// contains filtered or unexported fields
}
```

AVLTreeNode is a single element within the tree.

##### (*AVLTreeNode) Next 

``` go
func (node *AVLTreeNode) Next() *AVLTreeNode
```

Next returns the next element in an inorder walk of the AVL tree.

##### (*AVLTreeNode) Prev 

``` go
func (node *AVLTreeNode) Prev() *AVLTreeNode
```

Prev returns the previous element in an inorder walk of the AVL tree.

#### type BTree 

``` go
type BTree struct {
	// contains filtered or unexported fields
}
```

BTree holds elements of the B-tree.

##### func NewBTree 

``` go
func NewBTree(m int, comparator func(v1, v2 interface{}) int, safe ...bool) *BTree
```

NewBTree instantiates a B-tree with `m` (maximum number of children) and a custom key comparator. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default. Note that the `m` must be greater or equal than 3, or else it panics.

##### Example

``` go
```
##### func NewBTreeFrom 

``` go
func NewBTreeFrom(m int, comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *BTree
```

NewBTreeFrom instantiates a B-tree with `m` (maximum number of children), a custom key comparator and data map. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

##### Example

``` go
```
##### (*BTree) Clear 

``` go
func (tree *BTree) Clear()
```

Clear removes all nodes from the tree.

##### Example

``` go
```
##### (*BTree) Clone 

``` go
func (tree *BTree) Clone() *BTree
```

Clone returns a new tree with a copy of current tree.

##### Example

``` go
```
##### (*BTree) Contains 

``` go
func (tree *BTree) Contains(key interface{}) bool
```

Contains checks whether `key` exists in the tree.

##### Example

``` go
```
##### (*BTree) Get 

``` go
func (tree *BTree) Get(key interface{}) (value interface{})
```

Get searches the node in the tree by `key` and returns its value or nil if key is not found in tree.

##### Example

``` go
```
##### (*BTree) GetOrSet 

``` go
func (tree *BTree) GetOrSet(key interface{}, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*BTree) GetOrSetFunc 

``` go
func (tree *BTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*BTree) GetOrSetFuncLock 

``` go
func (tree *BTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*BTree) GetVar 

``` go
func (tree *BTree) GetVar(key interface{}) *gvar.Var
```

GetVar returns a gvar.Var with the value by given `key`. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*BTree) GetVarOrSet 

``` go
func (tree *BTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var
```

GetVarOrSet returns a gvar.Var with result from GetVarOrSet. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*BTree) GetVarOrSetFunc 

``` go
func (tree *BTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*BTree) GetVarOrSetFuncLock 

``` go
func (tree *BTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*BTree) Height 

``` go
func (tree *BTree) Height() int
```

Height returns the height of the tree.

##### Example

``` go
```
##### (*BTree) IsEmpty 

``` go
func (tree *BTree) IsEmpty() bool
```

IsEmpty returns true if tree does not contain any nodes

##### Example

``` go
```
##### (*BTree) Iterator 

``` go
func (tree *BTree) Iterator(f func(key, value interface{}) bool)
```

Iterator is alias of IteratorAsc.

##### Example

``` go
```
##### (*BTree) IteratorAsc 

``` go
func (tree *BTree) IteratorAsc(f func(key, value interface{}) bool)
```

IteratorAsc iterates the tree readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*BTree) IteratorAscFrom 

``` go
func (tree *BTree) IteratorAscFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorAscFrom iterates the tree readonly in ascending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

##### (*BTree) IteratorDesc 

``` go
func (tree *BTree) IteratorDesc(f func(key, value interface{}) bool)
```

IteratorDesc iterates the tree readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*BTree) IteratorDescFrom 

``` go
func (tree *BTree) IteratorDescFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorDescFrom iterates the tree readonly in descending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*BTree) IteratorFrom 

``` go
func (tree *BTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorFrom is alias of IteratorAscFrom.

##### Example

``` go
```
##### (*BTree) Keys 

``` go
func (tree *BTree) Keys() []interface{}
```

Keys returns all keys in asc order.

##### Example

``` go
```
##### (*BTree) Left 

``` go
func (tree *BTree) Left() *BTreeEntry
```

Left returns the left-most (min) entry or nil if tree is empty.

##### Example

``` go
```
##### (*BTree) Map 

``` go
func (tree *BTree) Map() map[interface{}]interface{}
```

Map returns all key-value items as map.

##### Example

``` go
```
##### (*BTree) MapStrAny 

``` go
func (tree *BTree) MapStrAny() map[string]interface{}
```

MapStrAny returns all key-value items as map[string]interface{}.

##### Example

``` go
```
##### (BTree) MarshalJSON 

``` go
func (tree BTree) MarshalJSON() (jsonBytes []byte, err error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*BTree) Print 

``` go
func (tree *BTree) Print()
```

Print prints the tree to stdout.

##### Example

``` go
```
##### (*BTree) Remove 

``` go
func (tree *BTree) Remove(key interface{}) (value interface{})
```

Remove removes the node from the tree by `key`.

##### Example

``` go
```
##### (*BTree) Removes 

``` go
func (tree *BTree) Removes(keys []interface{})
```

Removes batch deletes values of the tree by `keys`.

##### Example

``` go
```
##### (*BTree) Replace 

``` go
func (tree *BTree) Replace(data map[interface{}]interface{})
```

Replace the data of the tree with given `data`.

##### Example

``` go
```
##### (*BTree) Right 

``` go
func (tree *BTree) Right() *BTreeEntry
```

Right returns the right-most (max) entry or nil if tree is empty.

##### Example

``` go
```
##### (*BTree) Search 

``` go
func (tree *BTree) Search(key interface{}) (value interface{}, found bool)
```

Search searches the tree with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### Example

``` go
```
##### (*BTree) Set 

``` go
func (tree *BTree) Set(key interface{}, value interface{})
```

Set inserts key-value item into the tree.

##### Example

``` go
```
##### (*BTree) SetIfNotExist 

``` go
func (tree *BTree) SetIfNotExist(key interface{}, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*BTree) SetIfNotExistFunc 

``` go
func (tree *BTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*BTree) SetIfNotExistFuncLock 

``` go
func (tree *BTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*BTree) Sets 

``` go
func (tree *BTree) Sets(data map[interface{}]interface{})
```

Sets batch sets key-values to the tree.

##### Example

``` go
```
##### (*BTree) Size 

``` go
func (tree *BTree) Size() int
```

Size returns number of nodes in the tree.

##### Example

``` go
```
##### (*BTree) String 

``` go
func (tree *BTree) String() string
```

String returns a string representation of container (for debugging purposes)

##### Example

``` go
```
##### (*BTree) Values 

``` go
func (tree *BTree) Values() []interface{}
```

Values returns all values in asc order based on the key.

##### Example

``` go
```
#### type BTreeEntry 

``` go
type BTreeEntry struct {
	Key   interface{}
	Value interface{}
}
```

BTreeEntry represents the key-value pair contained within nodes.

#### type BTreeNode 

``` go
type BTreeNode struct {
	Parent   *BTreeNode
	Entries  []*BTreeEntry // Contained keys in node
	Children []*BTreeNode  // Children nodes
}
```

BTreeNode is a single element within the tree.

#### type RedBlackTree 

``` go
type RedBlackTree struct {
	// contains filtered or unexported fields
}
```

RedBlackTree holds elements of the red-black tree.

##### func NewRedBlackTree 

``` go
func NewRedBlackTree(comparator func(v1, v2 interface{}) int, safe ...bool) *RedBlackTree
```

NewRedBlackTree instantiates a red-black tree with the custom key comparator. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewRedBlackTreeFrom 

``` go
func NewRedBlackTreeFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *RedBlackTree
```

NewRedBlackTreeFrom instantiates a red-black tree with the custom key comparator and `data` map. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

##### Example

``` go
```
##### (*RedBlackTree) Ceiling 

``` go
func (tree *RedBlackTree) Ceiling(key interface{}) (ceiling *RedBlackTreeNode, found bool)
```

Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling node is found. Second return parameter is true if ceiling was found, otherwise false.

Ceiling node is defined as the smallest node that its key is larger than or equal to the given `key`. A ceiling node may not be found, either because the tree is empty, or because all nodes in the tree are smaller than the given node.

##### Example

``` go
```
##### (*RedBlackTree) Clear 

``` go
func (tree *RedBlackTree) Clear()
```

Clear removes all nodes from the tree.

##### Example

``` go
```
##### (*RedBlackTree) Clone 

``` go
func (tree *RedBlackTree) Clone() *RedBlackTree
```

Clone returns a new tree with a copy of current tree.

##### Example

``` go
```
##### (*RedBlackTree) Contains 

``` go
func (tree *RedBlackTree) Contains(key interface{}) bool
```

Contains checks whether `key` exists in the tree.

##### Example

``` go
```
##### (*RedBlackTree) Flip 

``` go
func (tree *RedBlackTree) Flip(comparator ...func(v1, v2 interface{}) int)
```

Flip exchanges key-value of the tree to value-key. Note that you should guarantee the value is the same type as key, or else the comparator would panic.

If the type of value is different with key, you pass the new `comparator`.

##### Example

``` go
```
##### (*RedBlackTree) Floor 

``` go
func (tree *RedBlackTree) Floor(key interface{}) (floor *RedBlackTreeNode, found bool)
```

Floor Finds floor node of the input key, return the floor node or nil if no floor node is found. Second return parameter is true if floor was found, otherwise false.

Floor node is defined as the largest node that its key is smaller than or equal to the given `key`. A floor node may not be found, either because the tree is empty, or because all nodes in the tree are larger than the given node.

##### Example

``` go
```
##### (*RedBlackTree) Get 

``` go
func (tree *RedBlackTree) Get(key interface{}) (value interface{})
```

Get searches the node in the tree by `key` and returns its value or nil if key is not found in tree.

##### Example

``` go
```
##### (*RedBlackTree) GetOrSet 

``` go
func (tree *RedBlackTree) GetOrSet(key interface{}, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*RedBlackTree) GetOrSetFunc 

``` go
func (tree *RedBlackTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*RedBlackTree) GetOrSetFuncLock 

``` go
func (tree *RedBlackTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*RedBlackTree) GetVar 

``` go
func (tree *RedBlackTree) GetVar(key interface{}) *gvar.Var
```

GetVar returns a gvar.Var with the value by given `key`. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*RedBlackTree) GetVarOrSet 

``` go
func (tree *RedBlackTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var
```

GetVarOrSet returns a gvar.Var with result from GetVarOrSet. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*RedBlackTree) GetVarOrSetFunc 

``` go
func (tree *RedBlackTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*RedBlackTree) GetVarOrSetFuncLock 

``` go
func (tree *RedBlackTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock. The returned gvar.Var is un-concurrent safe.

##### Example

``` go
```
##### (*RedBlackTree) IsEmpty 

``` go
func (tree *RedBlackTree) IsEmpty() bool
```

IsEmpty returns true if tree does not contain any nodes.

##### Example

``` go
```
##### (*RedBlackTree) Iterator 

``` go
func (tree *RedBlackTree) Iterator(f func(key, value interface{}) bool)
```

Iterator is alias of IteratorAsc.

##### Example

``` go
```
##### (*RedBlackTree) IteratorAsc 

``` go
func (tree *RedBlackTree) IteratorAsc(f func(key, value interface{}) bool)
```

IteratorAsc iterates the tree readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*RedBlackTree) IteratorAscFrom 

``` go
func (tree *RedBlackTree) IteratorAscFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorAscFrom iterates the tree readonly in ascending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

##### (*RedBlackTree) IteratorDesc 

``` go
func (tree *RedBlackTree) IteratorDesc(f func(key, value interface{}) bool)
```

IteratorDesc iterates the tree readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*RedBlackTree) IteratorDescFrom 

``` go
func (tree *RedBlackTree) IteratorDescFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorDescFrom iterates the tree readonly in descending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*RedBlackTree) IteratorFrom 

``` go
func (tree *RedBlackTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorFrom is alias of IteratorAscFrom.

##### Example

``` go
```
##### (*RedBlackTree) Keys 

``` go
func (tree *RedBlackTree) Keys() []interface{}
```

Keys returns all keys in asc order.

##### Example

``` go
```
##### (*RedBlackTree) Left 

``` go
func (tree *RedBlackTree) Left() *RedBlackTreeNode
```

Left returns the left-most (min) node or nil if tree is empty.

##### Example

``` go
```
##### (*RedBlackTree) Map 

``` go
func (tree *RedBlackTree) Map() map[interface{}]interface{}
```

Map returns all key-value items as map.

##### Example

``` go
```
##### (*RedBlackTree) MapStrAny 

``` go
func (tree *RedBlackTree) MapStrAny() map[string]interface{}
```

MapStrAny returns all key-value items as map[string]interface{}.

##### Example

``` go
```
##### (RedBlackTree) MarshalJSON 

``` go
func (tree RedBlackTree) MarshalJSON() (jsonBytes []byte, err error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*RedBlackTree) Print 

``` go
func (tree *RedBlackTree) Print()
```

Print prints the tree to stdout.

##### Example

``` go
```
##### (*RedBlackTree) Remove 

``` go
func (tree *RedBlackTree) Remove(key interface{}) (value interface{})
```

Remove removes the node from the tree by `key`.

##### Example

``` go
```
##### (*RedBlackTree) Removes 

``` go
func (tree *RedBlackTree) Removes(keys []interface{})
```

Removes batch deletes values of the tree by `keys`.

##### Example

``` go
```
##### (*RedBlackTree) Replace 

``` go
func (tree *RedBlackTree) Replace(data map[interface{}]interface{})
```

Replace the data of the tree with given `data`.

##### Example

``` go
```
##### (*RedBlackTree) Right 

``` go
func (tree *RedBlackTree) Right() *RedBlackTreeNode
```

Right returns the right-most (max) node or nil if tree is empty.

##### Example

``` go
```
##### (*RedBlackTree) Search 

``` go
func (tree *RedBlackTree) Search(key interface{}) (value interface{}, found bool)
```

Search searches the tree with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### Example

``` go
```
##### (*RedBlackTree) Set 

``` go
func (tree *RedBlackTree) Set(key interface{}, value interface{})
```

Set inserts key-value item into the tree.

##### Example

``` go
```
##### (*RedBlackTree) SetComparator 

``` go
func (tree *RedBlackTree) SetComparator(comparator func(a, b interface{}) int)
```

SetComparator sets/changes the comparator for sorting.

##### Example

``` go
```
##### (*RedBlackTree) SetIfNotExist 

``` go
func (tree *RedBlackTree) SetIfNotExist(key interface{}, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*RedBlackTree) SetIfNotExistFunc 

``` go
func (tree *RedBlackTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*RedBlackTree) SetIfNotExistFuncLock 

``` go
func (tree *RedBlackTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*RedBlackTree) Sets 

``` go
func (tree *RedBlackTree) Sets(data map[interface{}]interface{})
```

Sets batch sets key-values to the tree.

##### Example

``` go
```
##### (*RedBlackTree) Size 

``` go
func (tree *RedBlackTree) Size() int
```

Size returns number of nodes in the tree.

##### Example

``` go
```
##### (*RedBlackTree) String 

``` go
func (tree *RedBlackTree) String() string
```

String returns a string representation of container.

##### Example

``` go
```
##### (*RedBlackTree) UnmarshalJSON 

``` go
func (tree *RedBlackTree) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*RedBlackTree) UnmarshalValue 

``` go
func (tree *RedBlackTree) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

##### Example

``` go
```
##### (*RedBlackTree) Values 

``` go
func (tree *RedBlackTree) Values() []interface{}
```

Values returns all values in asc order based on the key.

##### Example

``` go
```
#### type RedBlackTreeNode 

``` go
type RedBlackTreeNode struct {
	Key   interface{}
	Value interface{}
	// contains filtered or unexported fields
}
```

RedBlackTreeNode is a single element within the tree.