+++
title = "gtree"
date = 2024-03-21T17:45:11+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gtree](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gtree)

Package gtree provides concurrent-safe/unsafe tree containers.

​	软件包 gtree 提供了并发安全/不安全的树容器。

Some implements are from: https://github.com/emirpasic/gods

​	一些工具来自：https://github.com/emirpasic/gods

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type AVLTree

```go
type AVLTree struct {
	// contains filtered or unexported fields
}
```

AVLTree holds elements of the AVL tree.

​	AVLTree 保存 AVL 树的元素。

#### func NewAVLTree

```go
func NewAVLTree(comparator func(v1, v2 interface{}) int, safe ...bool) *AVLTree
```

NewAVLTree instantiates an AVL tree with the custom key comparator. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

​	NewAVLTree 使用自定义密钥比较器实例化 AVL 树。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 tree，默认为 false。

##### Example

``` go
```

#### func NewAVLTreeFrom

```go
func NewAVLTreeFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *AVLTree
```

NewAVLTreeFrom instantiates an AVL tree with the custom key comparator and data map. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

​	NewAVLTreeFrom 使用自定义键比较器和数据映射实例化 AVL 树。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 tree，默认为 false。

##### Example

``` go
```

#### (*AVLTree) Ceiling

```go
func (tree *AVLTree) Ceiling(key interface{}) (ceiling *AVLTreeNode, found bool)
```

Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling node is found. Second return parameter is true if ceiling was found, otherwise false.

​	ceiling 查找输入键的 ceiling 节点，返回 ceiling 节点，如果没有找到 ceiling 节点，则返回 nil。如果找到 ceiling，则第二个返回参数为 true，否则为 false。

Ceiling node is defined as the smallest node that is larger than or equal to the given node. A ceiling node may not be found, either because the tree is empty, or because all nodes in the tree is smaller than the given node.

​	天花板节点定义为大于或等于给定节点的最小节点。可能找不到天花板节点，因为树为空，或者树中的所有节点都小于给定节点。

Key should adhere to the comparator’s type assertion, otherwise method panics.

​	键应遵循比较器的类型断言，否则方法会崩溃。

##### Example

``` go
```

#### (*AVLTree) Clear

```go
func (tree *AVLTree) Clear()
```

Clear removes all nodes from the tree.

​	清除将从树中删除所有节点。

##### Example

``` go
```

#### (*AVLTree) Clone

```go
func (tree *AVLTree) Clone() *AVLTree
```

Clone returns a new tree with a copy of current tree.

​	克隆返回一个新树，其中包含当前树的副本。

##### Example

``` go
```

#### (*AVLTree) Contains

```go
func (tree *AVLTree) Contains(key interface{}) bool
```

Contains checks whether `key` exists in the tree.

​	包含检查树中是否 `key` 存在。

##### Example

``` go
```

#### (*AVLTree) Flip

```go
func (tree *AVLTree) Flip(comparator ...func(v1, v2 interface{}) int)
```

Flip exchanges key-value of the tree to value-key. Note that you should guarantee the value is the same type as key, or else the comparator would panic.

​	Flip 将树的键值交换为值键。请注意，应保证该值与键的类型相同，否则比较器会崩溃。

If the type of value is different with key, you pass the new `comparator`.

​	如果值的类型与键不同，则传递新的 `comparator` .

##### Example

``` go
```

#### (*AVLTree) Floor

```go
func (tree *AVLTree) Floor(key interface{}) (floor *AVLTreeNode, found bool)
```

Floor Finds floor node of the input key, return the floor node or nil if no floor node is found. Second return parameter is true if floor was found, otherwise false.

​	Floor 查找输入键的 floor 节点，返回 floor 节点，如果没有找到 floor 节点，则返回 nil。如果找到 floor，则第二个返回参数为 true，否则为 false。

Floor node is defined as the largest node that is smaller than or equal to the given node. A floor node may not be found, either because the tree is empty, or because all nodes in the tree is larger than the given node.

​	地板节点定义为小于或等于给定节点的最大节点。可能找不到楼层节点，因为树为空，或者树中的所有节点都大于给定节点。

Key should adhere to the comparator’s type assertion, otherwise method panics.

​	键应遵循比较器的类型断言，否则方法会崩溃。

##### Example

``` go
```

#### (*AVLTree) Get

```go
func (tree *AVLTree) Get(key interface{}) (value interface{})
```

Get searches the node in the tree by `key` and returns its value or nil if key is not found in tree.

​	Get 搜索树 `key` 中的节点，如果在树中找不到键，则返回其值或 nil。

##### Example

``` go
```

#### (*AVLTree) GetOrSet

```go
func (tree *AVLTree) GetOrSet(key interface{}, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

##### Example

``` go
```

#### (*AVLTree) GetOrSetFunc

```go
func (tree *AVLTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

##### Example

``` go
```

#### (*AVLTree) GetOrSetFuncLock

```go
func (tree *AVLTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*AVLTree) GetVar

```go
func (tree *AVLTree) GetVar(key interface{}) *gvar.Var
```

GetVar returns a gvar.Var with the value by given `key`. The returned gvar.Var is un-concurrent safe.

​	GetVar 返回一个 gvar。Var 的值由 给定 `key` 。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*AVLTree) GetVarOrSet

```go
func (tree *AVLTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var
```

GetVarOrSet returns a gvar.Var with result from GetVarOrSet. The returned gvar.Var is un-concurrent safe.

​	GetVarOrSet 返回 gvar。具有 GetVarOrSet 结果的 Var。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*AVLTree) GetVarOrSetFunc

```go
func (tree *AVLTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc. The returned gvar.Var is un-concurrent safe.

​	GetVarOrSetFunc 返回 gvar。具有 GetOrSetFunc 结果的 Var。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*AVLTree) GetVarOrSetFuncLock

```go
func (tree *AVLTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock. The returned gvar.Var is un-concurrent safe.

​	GetVarOrSetFuncLock 返回 gvar。具有 GetOrSetFuncLock 结果的 Var。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*AVLTree) IsEmpty

```go
func (tree *AVLTree) IsEmpty() bool
```

IsEmpty returns true if tree does not contain any nodes.

​	如果树不包含任何节点，则 IsEmpty 返回 true。

##### Example

``` go
```

#### (*AVLTree) Iterator

```go
func (tree *AVLTree) Iterator(f func(key, value interface{}) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

##### Example

``` go
```

#### (*AVLTree) IteratorAsc

```go
func (tree *AVLTree) IteratorAsc(f func(key, value interface{}) bool)
```

IteratorAsc iterates the tree readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序迭代树只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*AVLTree) IteratorAscFrom

```go
func (tree *AVLTree) IteratorAscFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorAscFrom iterates the tree readonly in ascending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAscFrom 使用给定的回调函数 `f` 按升序迭代树只读。该参数 `key` 指定用于迭代的开始条目。指定 `match` 在完全 `key` 匹配时是开始迭代，还是使用索引搜索迭代。如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*AVLTree) IteratorDesc

```go
func (tree *AVLTree) IteratorDesc(f func(key, value interface{}) bool)
```

IteratorDesc iterates the tree readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序只读迭代树。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*AVLTree) IteratorDescFrom

```go
func (tree *AVLTree) IteratorDescFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorDescFrom iterates the tree readonly in descending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDescFrom 使用给定的回调函数 `f` 按降序迭代树只读。该参数 `key` 指定用于迭代的开始条目。指定 `match` 在完全 `key` 匹配时是开始迭代，还是使用索引搜索迭代。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*AVLTree) IteratorFrom

```go
func (tree *AVLTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorFrom is alias of IteratorAscFrom.

​	IteratorFrom 是 IteratorAscFrom 的别名。

##### Example

``` go
```

#### (*AVLTree) Keys

```go
func (tree *AVLTree) Keys() []interface{}
```

Keys returns all keys in asc order.

​	Keys 按 asc 顺序返回所有键。

##### Example

``` go
```

#### (*AVLTree) Left

```go
func (tree *AVLTree) Left() *AVLTreeNode
```

Left returns the minimum element of the AVL tree or nil if the tree is empty.

​	Left 返回 AVL 树的最小元素，如果树为空，则返回 nil。

##### Example

``` go
```

#### (*AVLTree) Map

```go
func (tree *AVLTree) Map() map[interface{}]interface{}
```

Map returns all key-value items as map.

​	Map 以 map 的形式返回所有键值项。

##### Example

``` go
```

#### (*AVLTree) MapStrAny

```go
func (tree *AVLTree) MapStrAny() map[string]interface{}
```

MapStrAny returns all key-value items as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回所有键值项。

##### Example

``` go
```

#### (AVLTree) MarshalJSON

```go
func (tree AVLTree) MarshalJSON() (jsonBytes []byte, err error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*AVLTree) Print

```go
func (tree *AVLTree) Print()
```

Print prints the tree to stdout.

​	打印 将树打印到 stdout。

##### Example

``` go
```

#### (*AVLTree) Remove

```go
func (tree *AVLTree) Remove(key interface{}) (value interface{})
```

Remove removes the node from the tree by key. Key should adhere to the comparator’s type assertion, otherwise method panics.

​	Remove 按键从树中删除节点。键应遵循比较器的类型断言，否则方法会崩溃。

##### Example

``` go
```

#### (*AVLTree) Removes

```go
func (tree *AVLTree) Removes(keys []interface{})
```

Removes batch deletes values of the tree by `keys`.

​	删除 batch 删除树 `keys` 的值。

##### Example

``` go
```

#### (*AVLTree) Replace

```go
func (tree *AVLTree) Replace(data map[interface{}]interface{})
```

Replace the data of the tree with given `data`.

​	将树的数据替换为 给定 `data` 的 .

##### Example

``` go
```

#### (*AVLTree) Right

```go
func (tree *AVLTree) Right() *AVLTreeNode
```

Right returns the maximum element of the AVL tree or nil if the tree is empty.

​	Right 返回 AVL 树的最大元素，如果树为空，则返回 nil。

##### Example

``` go
```

#### (*AVLTree) Search

```go
func (tree *AVLTree) Search(key interface{}) (value interface{}, found bool)
```

Search searches the tree with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的树搜索 。如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

##### Example

``` go
```

#### (*AVLTree) Set

```go
func (tree *AVLTree) Set(key interface{}, value interface{})
```

Set inserts node into the tree.

​	将插入节点设置到树中。

##### Example

``` go
```

#### (*AVLTree) SetIfNotExist

```go
func (tree *AVLTree) SetIfNotExist(key interface{}, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*AVLTree) SetIfNotExistFunc

```go
func (tree *AVLTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*AVLTree) SetIfNotExistFuncLock

```go
func (tree *AVLTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*AVLTree) Sets

```go
func (tree *AVLTree) Sets(data map[interface{}]interface{})
```

Sets batch sets key-values to the tree.

​	将批处理设置树的键值。

##### Example

``` go
```

#### (*AVLTree) Size

```go
func (tree *AVLTree) Size() int
```

Size returns number of nodes in the tree.

​	size 返回树中的节点数。

##### Example

``` go
```

#### (*AVLTree) String

```go
func (tree *AVLTree) String() string
```

String returns a string representation of container

​	String 返回容器的字符串表示形式

##### Example

``` go
```

#### (*AVLTree) Values

```go
func (tree *AVLTree) Values() []interface{}
```

Values returns all values in asc order based on the key.

​	Values 根据键按 asc 顺序返回所有值。

##### Example

``` go
```

### type AVLTreeNode

```go
type AVLTreeNode struct {
	Key   interface{}
	Value interface{}
	// contains filtered or unexported fields
}
```

AVLTreeNode is a single element within the tree.

​	AVLTreeNode 是树中的单个元素。

#### (*AVLTreeNode) Next

```go
func (node *AVLTreeNode) Next() *AVLTreeNode
```

Next returns the next element in an inorder walk of the AVL tree.

​	Next 返回 AVL 树的无序游历中的下一个元素。

#### (*AVLTreeNode) Prev

```go
func (node *AVLTreeNode) Prev() *AVLTreeNode
```

Prev returns the previous element in an inorder walk of the AVL tree.

​	Prev 返回 AVL 树的无序遍历中的前一个元素。

### type BTree

```go
type BTree struct {
	// contains filtered or unexported fields
}
```

BTree holds elements of the B-tree.

​	BTree 包含 B 树的元素。

#### func NewBTree

```go
func NewBTree(m int, comparator func(v1, v2 interface{}) int, safe ...bool) *BTree
```

NewBTree instantiates a B-tree with `m` (maximum number of children) and a custom key comparator. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default. Note that the `m` must be greater or equal than 3, or else it panics.

​	NewBTree 实例化具有 `m` （最大子项数）和自定义键比较器的 B 树。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 tree，默认为 false。请注意，必须 `m` 大于或等于 3，否则会出现恐慌。

##### Example

``` go
```

#### func NewBTreeFrom

```go
func NewBTreeFrom(m int, comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *BTree
```

NewBTreeFrom instantiates a B-tree with `m` (maximum number of children), a custom key comparator and data map. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

​	NewBTreeFrom 实例化具有 `m` （最大子项数）、自定义键比较器和数据映射的 B 树。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 tree，默认为 false。

##### Example

``` go
```

#### (*BTree) Clear

```go
func (tree *BTree) Clear()
```

Clear removes all nodes from the tree.

​	清除将从树中删除所有节点。

##### Example

``` go
```

#### (*BTree) Clone

```go
func (tree *BTree) Clone() *BTree
```

Clone returns a new tree with a copy of current tree.

​	克隆返回一个新树，其中包含当前树的副本。

##### Example

``` go
```

#### (*BTree) Contains

```go
func (tree *BTree) Contains(key interface{}) bool
```

Contains checks whether `key` exists in the tree.

​	包含检查树中是否 `key` 存在。

##### Example

``` go
```

#### (*BTree) Get

```go
func (tree *BTree) Get(key interface{}) (value interface{})
```

Get searches the node in the tree by `key` and returns its value or nil if key is not found in tree.

​	Get 搜索树 `key` 中的节点，如果在树中找不到键，则返回其值或 nil。

##### Example

``` go
```

#### (*BTree) GetOrSet

```go
func (tree *BTree) GetOrSet(key interface{}, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

##### Example

``` go
```

#### (*BTree) GetOrSetFunc

```go
func (tree *BTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

##### Example

``` go
```

#### (*BTree) GetOrSetFuncLock

```go
func (tree *BTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*BTree) GetVar

```go
func (tree *BTree) GetVar(key interface{}) *gvar.Var
```

GetVar returns a gvar.Var with the value by given `key`. The returned gvar.Var is un-concurrent safe.

​	GetVar 返回一个 gvar。Var 的值由 给定 `key` 。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*BTree) GetVarOrSet

```go
func (tree *BTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var
```

GetVarOrSet returns a gvar.Var with result from GetVarOrSet. The returned gvar.Var is un-concurrent safe.

​	GetVarOrSet 返回 gvar。具有 GetVarOrSet 结果的 Var。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*BTree) GetVarOrSetFunc

```go
func (tree *BTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc. The returned gvar.Var is un-concurrent safe.

​	GetVarOrSetFunc 返回 gvar。具有 GetOrSetFunc 结果的 Var。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*BTree) GetVarOrSetFuncLock

```go
func (tree *BTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock. The returned gvar.Var is un-concurrent safe.

​	GetVarOrSetFuncLock 返回 gvar。具有 GetOrSetFuncLock 结果的 Var。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*BTree) Height

```go
func (tree *BTree) Height() int
```

Height returns the height of the tree.

​	Height 返回树的高度。

##### Example

``` go
```

#### (*BTree) IsEmpty

```go
func (tree *BTree) IsEmpty() bool
```

IsEmpty returns true if tree does not contain any nodes

​	如果树不包含任何节点，则 IsEmpty 返回 true

##### Example

``` go
```

#### (*BTree) Iterator

```go
func (tree *BTree) Iterator(f func(key, value interface{}) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

##### Example

``` go
```

#### (*BTree) IteratorAsc

```go
func (tree *BTree) IteratorAsc(f func(key, value interface{}) bool)
```

IteratorAsc iterates the tree readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序迭代树只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*BTree) IteratorAscFrom

```go
func (tree *BTree) IteratorAscFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorAscFrom iterates the tree readonly in ascending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAscFrom 使用给定的回调函数 `f` 按升序迭代树只读。该参数 `key` 指定用于迭代的开始条目。指定 `match` 在完全 `key` 匹配时是开始迭代，还是使用索引搜索迭代。如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*BTree) IteratorDesc

```go
func (tree *BTree) IteratorDesc(f func(key, value interface{}) bool)
```

IteratorDesc iterates the tree readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序只读迭代树。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*BTree) IteratorDescFrom

```go
func (tree *BTree) IteratorDescFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorDescFrom iterates the tree readonly in descending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDescFrom 使用给定的回调函数 `f` 按降序迭代树只读。该参数 `key` 指定用于迭代的开始条目。指定 `match` 在完全 `key` 匹配时是开始迭代，还是使用索引搜索迭代。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*BTree) IteratorFrom

```go
func (tree *BTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorFrom is alias of IteratorAscFrom.

​	IteratorFrom 是 IteratorAscFrom 的别名。

##### Example

``` go
```

#### (*BTree) Keys

```go
func (tree *BTree) Keys() []interface{}
```

Keys returns all keys in asc order.

​	Keys 按 asc 顺序返回所有键。

##### Example

``` go
```

#### (*BTree) Left

```go
func (tree *BTree) Left() *BTreeEntry
```

Left returns the left-most (min) entry or nil if tree is empty.

​	Left 返回最左边的 （min） 条目，如果 tree 为空，则返回 nil。

##### Example

``` go
```

#### (*BTree) Map

```go
func (tree *BTree) Map() map[interface{}]interface{}
```

Map returns all key-value items as map.

​	Map 以 map 的形式返回所有键值项。

##### Example

``` go
```

#### (*BTree) MapStrAny

```go
func (tree *BTree) MapStrAny() map[string]interface{}
```

MapStrAny returns all key-value items as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回所有键值项。

##### Example

``` go
```

#### (BTree) MarshalJSON

```go
func (tree BTree) MarshalJSON() (jsonBytes []byte, err error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*BTree) Print

```go
func (tree *BTree) Print()
```

Print prints the tree to stdout.

​	打印 将树打印到 stdout。

##### Example

``` go
```

#### (*BTree) Remove

```go
func (tree *BTree) Remove(key interface{}) (value interface{})
```

Remove removes the node from the tree by `key`.

​	Remove 通过 `key` 从树中删除节点。

##### Example

``` go
```

#### (*BTree) Removes

```go
func (tree *BTree) Removes(keys []interface{})
```

Removes batch deletes values of the tree by `keys`.

​	删除 batch 删除树 `keys` 的值。

##### Example

``` go
```

#### (*BTree) Replace

```go
func (tree *BTree) Replace(data map[interface{}]interface{})
```

Replace the data of the tree with given `data`.

​	将树的数据替换为 给定 `data` 的 .

##### Example

``` go
```

#### (*BTree) Right

```go
func (tree *BTree) Right() *BTreeEntry
```

Right returns the right-most (max) entry or nil if tree is empty.

​	如果树为空，则 Right 返回最右边 （max） 的条目或 nil。

##### Example

``` go
```

#### (*BTree) Search

```go
func (tree *BTree) Search(key interface{}) (value interface{}, found bool)
```

Search searches the tree with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的树搜索 。如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

##### Example

``` go
```

#### (*BTree) Set

```go
func (tree *BTree) Set(key interface{}, value interface{})
```

Set inserts key-value item into the tree.

​	Set 将键值项插入到树中。

##### Example

``` go
```

#### (*BTree) SetIfNotExist

```go
func (tree *BTree) SetIfNotExist(key interface{}, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*BTree) SetIfNotExistFunc

```go
func (tree *BTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*BTree) SetIfNotExistFuncLock

```go
func (tree *BTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*BTree) Sets

```go
func (tree *BTree) Sets(data map[interface{}]interface{})
```

Sets batch sets key-values to the tree.

​	将批处理设置树的键值。

##### Example

``` go
```

#### (*BTree) Size

```go
func (tree *BTree) Size() int
```

Size returns number of nodes in the tree.

​	size 返回树中的节点数。

##### Example

``` go
```

#### (*BTree) String

```go
func (tree *BTree) String() string
```

String returns a string representation of container (for debugging purposes)

​	String 返回容器的字符串表示形式（用于调试目的）

##### Example

``` go
```

#### (*BTree) Values

```go
func (tree *BTree) Values() []interface{}
```

Values returns all values in asc order based on the key.

​	Values 根据键按 asc 顺序返回所有值。

##### Example

``` go
```

### type BTreeEntry

```go
type BTreeEntry struct {
	Key   interface{}
	Value interface{}
}
```

BTreeEntry represents the key-value pair contained within nodes.

​	BTreeEntry 表示节点中包含的键值对。

### type BTreeNode

```go
type BTreeNode struct {
	Parent   *BTreeNode
	Entries  []*BTreeEntry // Contained keys in node
	Children []*BTreeNode  // Children nodes
}
```

BTreeNode is a single element within the tree.

​	BTreeNode 是树中的单个元素。

### type RedBlackTree

```go
type RedBlackTree struct {
	// contains filtered or unexported fields
}
```

RedBlackTree holds elements of the red-black tree.

​	RedBlackTree 包含红黑树的元素。

#### func NewRedBlackTree

```go
func NewRedBlackTree(comparator func(v1, v2 interface{}) int, safe ...bool) *RedBlackTree
```

NewRedBlackTree instantiates a red-black tree with the custom key comparator. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

​	NewRedBlackTree 使用自定义键比较器实例化红黑树。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 tree，默认为 false。

##### Example

``` go
```

#### func NewRedBlackTreeFrom

```go
func NewRedBlackTreeFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *RedBlackTree
```

NewRedBlackTreeFrom instantiates a red-black tree with the custom key comparator and `data` map. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

​	NewRedBlackTreeFrom 使用自定义键比较器和 `data` 映射实例化红黑树。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 tree，默认为 false。

##### Example

``` go
```

#### (*RedBlackTree) Ceiling

```go
func (tree *RedBlackTree) Ceiling(key interface{}) (ceiling *RedBlackTreeNode, found bool)
```

Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling node is found. Second return parameter is true if ceiling was found, otherwise false.

​	ceiling 查找输入键的 ceiling 节点，返回 ceiling 节点，如果没有找到 ceiling 节点，则返回 nil。如果找到 ceiling，则第二个返回参数为 true，否则为 false。

Ceiling node is defined as the smallest node that its key is larger than or equal to the given `key`. A ceiling node may not be found, either because the tree is empty, or because all nodes in the tree are smaller than the given node.

​	天花板节点定义为其键大于或等于给定 `key` 的最小节点。可能找不到天花板节点，因为树是空的，或者树中的所有节点都小于给定节点。

##### Example

``` go
```

#### (*RedBlackTree) Clear

```go
func (tree *RedBlackTree) Clear()
```

Clear removes all nodes from the tree.

​	清除将从树中删除所有节点。

##### Example

``` go
```

#### (*RedBlackTree) Clone

```go
func (tree *RedBlackTree) Clone() *RedBlackTree
```

Clone returns a new tree with a copy of current tree.

​	克隆返回一个新树，其中包含当前树的副本。

##### Example

``` go
```

#### (*RedBlackTree) Contains

```go
func (tree *RedBlackTree) Contains(key interface{}) bool
```

Contains checks whether `key` exists in the tree.

​	包含检查树中是否 `key` 存在。

##### Example

``` go
```

#### (*RedBlackTree) Flip

```go
func (tree *RedBlackTree) Flip(comparator ...func(v1, v2 interface{}) int)
```

Flip exchanges key-value of the tree to value-key. Note that you should guarantee the value is the same type as key, or else the comparator would panic.

​	Flip 将树的键值交换为值键。请注意，应保证该值与键的类型相同，否则比较器会崩溃。

If the type of value is different with key, you pass the new `comparator`.

​	如果值的类型与键不同，则传递新的 `comparator` .

##### Example

``` go
```

#### (*RedBlackTree) Floor

```go
func (tree *RedBlackTree) Floor(key interface{}) (floor *RedBlackTreeNode, found bool)
```

Floor Finds floor node of the input key, return the floor node or nil if no floor node is found. Second return parameter is true if floor was found, otherwise false.

​	Floor 查找输入键的 floor 节点，返回 floor 节点，如果没有找到 floor 节点，则返回 nil。如果找到 floor，则第二个返回参数为 true，否则为 false。

Floor node is defined as the largest node that its key is smaller than or equal to the given `key`. A floor node may not be found, either because the tree is empty, or because all nodes in the tree are larger than the given node.

​	Floor 节点定义为其键小于或等于给定 `key` 的最大节点。可能找不到楼层节点，因为树为空，或者树中的所有节点都大于给定节点。

##### Example

``` go
```

#### (*RedBlackTree) Get

```go
func (tree *RedBlackTree) Get(key interface{}) (value interface{})
```

Get searches the node in the tree by `key` and returns its value or nil if key is not found in tree.

​	Get 搜索树 `key` 中的节点，如果在树中找不到键，则返回其值或 nil。

##### Example

``` go
```

#### (*RedBlackTree) GetOrSet

```go
func (tree *RedBlackTree) GetOrSet(key interface{}, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

##### Example

``` go
```

#### (*RedBlackTree) GetOrSetFunc

```go
func (tree *RedBlackTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

##### Example

``` go
```

#### (*RedBlackTree) GetOrSetFuncLock

```go
func (tree *RedBlackTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*RedBlackTree) GetVar

```go
func (tree *RedBlackTree) GetVar(key interface{}) *gvar.Var
```

GetVar returns a gvar.Var with the value by given `key`. The returned gvar.Var is un-concurrent safe.

​	GetVar 返回一个 gvar。Var 的值由 给定 `key` 。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*RedBlackTree) GetVarOrSet

```go
func (tree *RedBlackTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var
```

GetVarOrSet returns a gvar.Var with result from GetVarOrSet. The returned gvar.Var is un-concurrent safe.

​	GetVarOrSet 返回 gvar。具有 GetVarOrSet 结果的 Var。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*RedBlackTree) GetVarOrSetFunc

```go
func (tree *RedBlackTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc. The returned gvar.Var is un-concurrent safe.

​	GetVarOrSetFunc 返回 gvar。具有 GetOrSetFunc 结果的 Var。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*RedBlackTree) GetVarOrSetFuncLock

```go
func (tree *RedBlackTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock. The returned gvar.Var is un-concurrent safe.

​	GetVarOrSetFuncLock 返回 gvar。具有 GetOrSetFuncLock 结果的 Var。返回的 gvar。Var 是非并发安全的。

##### Example

``` go
```

#### (*RedBlackTree) IsEmpty

```go
func (tree *RedBlackTree) IsEmpty() bool
```

IsEmpty returns true if tree does not contain any nodes.

​	如果树不包含任何节点，则 IsEmpty 返回 true。

##### Example

``` go
```

#### (*RedBlackTree) Iterator

```go
func (tree *RedBlackTree) Iterator(f func(key, value interface{}) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

##### Example

``` go
```

#### (*RedBlackTree) IteratorAsc

```go
func (tree *RedBlackTree) IteratorAsc(f func(key, value interface{}) bool)
```

IteratorAsc iterates the tree readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序迭代树只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*RedBlackTree) IteratorAscFrom

```go
func (tree *RedBlackTree) IteratorAscFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorAscFrom iterates the tree readonly in ascending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAscFrom 使用给定的回调函数 `f` 按升序迭代树只读。该参数 `key` 指定用于迭代的开始条目。指定 `match` 在完全 `key` 匹配时是开始迭代，还是使用索引搜索迭代。如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*RedBlackTree) IteratorDesc

```go
func (tree *RedBlackTree) IteratorDesc(f func(key, value interface{}) bool)
```

IteratorDesc iterates the tree readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序只读迭代树。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*RedBlackTree) IteratorDescFrom

```go
func (tree *RedBlackTree) IteratorDescFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorDescFrom iterates the tree readonly in descending order with given callback function `f`. The parameter `key` specifies the start entry for iterating. The `match` specifies whether starting iterating if the `key` is fully matched, or else using index searching iterating. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDescFrom 使用给定的回调函数 `f` 按降序迭代树只读。该参数 `key` 指定用于迭代的开始条目。指定 `match` 在完全 `key` 匹配时是开始迭代，还是使用索引搜索迭代。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*RedBlackTree) IteratorFrom

```go
func (tree *RedBlackTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool)
```

IteratorFrom is alias of IteratorAscFrom.

​	IteratorFrom 是 IteratorAscFrom 的别名。

##### Example

``` go
```

#### (*RedBlackTree) Keys

```go
func (tree *RedBlackTree) Keys() []interface{}
```

Keys returns all keys in asc order.

​	Keys 按 asc 顺序返回所有键。

##### Example

``` go
```

#### (*RedBlackTree) Left

```go
func (tree *RedBlackTree) Left() *RedBlackTreeNode
```

Left returns the left-most (min) node or nil if tree is empty.

​	如果树为空，则 left 返回最左边的 （min） 节点或 nil。

##### Example

``` go
```

#### (*RedBlackTree) Map

```go
func (tree *RedBlackTree) Map() map[interface{}]interface{}
```

Map returns all key-value items as map.

​	Map 以 map 的形式返回所有键值项。

##### Example

``` go
```

#### (*RedBlackTree) MapStrAny

```go
func (tree *RedBlackTree) MapStrAny() map[string]interface{}
```

MapStrAny returns all key-value items as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回所有键值项。

##### Example

``` go
```

#### (RedBlackTree) MarshalJSON

```go
func (tree RedBlackTree) MarshalJSON() (jsonBytes []byte, err error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*RedBlackTree) Print

```go
func (tree *RedBlackTree) Print()
```

Print prints the tree to stdout.

​	打印 将树打印到 stdout。

##### Example

``` go
```

#### (*RedBlackTree) Remove

```go
func (tree *RedBlackTree) Remove(key interface{}) (value interface{})
```

Remove removes the node from the tree by `key`.

​	Remove 通过 `key` 从树中删除节点。

##### Example

``` go
```

#### (*RedBlackTree) Removes

```go
func (tree *RedBlackTree) Removes(keys []interface{})
```

Removes batch deletes values of the tree by `keys`.

​	删除 batch 删除树 `keys` 的值。

##### Example

``` go
```

#### (*RedBlackTree) Replace

```go
func (tree *RedBlackTree) Replace(data map[interface{}]interface{})
```

Replace the data of the tree with given `data`.

​	将树的数据替换为 给定 `data` 的 .

##### Example

``` go
```

#### (*RedBlackTree) Right

```go
func (tree *RedBlackTree) Right() *RedBlackTreeNode
```

Right returns the right-most (max) node or nil if tree is empty.

​	如果 tree 为空，则 Right 返回最右边的 （max） 节点或 nil。

##### Example

``` go
```

#### (*RedBlackTree) Search

```go
func (tree *RedBlackTree) Search(key interface{}) (value interface{}, found bool)
```

Search searches the tree with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的树搜索 。如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

##### Example

``` go
```

#### (*RedBlackTree) Set

```go
func (tree *RedBlackTree) Set(key interface{}, value interface{})
```

Set inserts key-value item into the tree.

​	Set 将键值项插入到树中。

##### Example

``` go
```

#### (*RedBlackTree) SetComparator

```go
func (tree *RedBlackTree) SetComparator(comparator func(a, b interface{}) int)
```

SetComparator sets/changes the comparator for sorting.

​	SetComparator 设置/更改用于排序的比较器。

##### Example

``` go
```

#### (*RedBlackTree) SetIfNotExist

```go
func (tree *RedBlackTree) SetIfNotExist(key interface{}, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*RedBlackTree) SetIfNotExistFunc

```go
func (tree *RedBlackTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*RedBlackTree) SetIfNotExistFuncLock

```go
func (tree *RedBlackTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*RedBlackTree) Sets

```go
func (tree *RedBlackTree) Sets(data map[interface{}]interface{})
```

Sets batch sets key-values to the tree.

​	将批处理设置树的键值。

##### Example

``` go
```

#### (*RedBlackTree) Size

```go
func (tree *RedBlackTree) Size() int
```

Size returns number of nodes in the tree.

​	size 返回树中的节点数。

##### Example

``` go
```

#### (*RedBlackTree) String

```go
func (tree *RedBlackTree) String() string
```

String returns a string representation of container.

​	String 返回容器的字符串表示形式。

##### Example

``` go
```

#### (*RedBlackTree) UnmarshalJSON

```go
func (tree *RedBlackTree) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*RedBlackTree) UnmarshalValue

```go
func (tree *RedBlackTree) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

​	UnmarshalValue 是一个接口实现，用于为 map 设置任何类型的值。

##### Example

``` go
```

#### (*RedBlackTree) Values

```go
func (tree *RedBlackTree) Values() []interface{}
```

Values returns all values in asc order based on the key.

​	Values 根据键按 asc 顺序返回所有值。

##### Example

``` go
```

### type RedBlackTreeNode

```go
type RedBlackTreeNode struct {
	Key   interface{}
	Value interface{}
	// contains filtered or unexported fields
}
```

RedBlackTreeNode is a single element within the tree.

​	RedBlackTreeNode 是树中的单个元素。