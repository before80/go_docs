+++
title = "gset"
date = 2024-03-21T17:45:04+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gset](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gset)

Package gset provides kinds of concurrent-safe/unsafe sets.

​	软件包 gset 提供了各种并发安全/不安全集。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type IntSet

```go
type IntSet struct {
	// contains filtered or unexported fields
}
```

#### func NewIntSet

```go
func NewIntSet(safe ...bool) *IntSet
```

NewIntSet create and returns a new set, which contains un-repeated items. The parameter `safe` is used to specify whether using set in concurrent-safety, which is false in default.

​	NewIntSet 创建并返回一个新集合，其中包含未重复的项。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 set，默认为 false。

##### Example

``` go
```

#### func NewIntSetFrom

```go
func NewIntSetFrom(items []int, safe ...bool) *IntSet
```

NewIntSetFrom returns a new set from `items`.

​	NewIntSetFrom 从 返回一个新 `items` 集。

#### (*IntSet) Add

```go
func (set *IntSet) Add(item ...int)
```

Add adds one or multiple items to the set.

​	“添加”会将一个或多个项目添加到集合中。

##### Example

``` go
```

#### (*IntSet) AddIfNotExist

```go
func (set *IntSet) AddIfNotExist(item int) bool
```

AddIfNotExist checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set, or else it does nothing and returns false.

​	AddIfNotExist 检查集合中是否存在项，将项添加到集合中，如果集合中不存在项，则返回 true，否则不执行任何操作并返回 false。

Note that, if `item` is nil, it does nothing and returns false.

​	请注意，如果 `item` 为 nil，则不执行任何操作并返回 false。

##### Example

``` go
```

#### (*IntSet) AddIfNotExistFunc

```go
func (set *IntSet) AddIfNotExistFunc(item int, f func() bool) bool
```

AddIfNotExistFunc checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set and function `f` returns true, or else it does nothing and returns false.

​	AddIfNotExistFunc 检查集合中是否存在项，将项添加到集合中，如果集合中不存在项，则返回 true，函数 `f` 返回 true，否则不执行任何操作并返回 false。

Note that, the function `f` is executed without writing lock.

​	请注意，该函数 `f` 是在没有写入锁的情况下执行的。

##### Example

``` go
```

#### (*IntSet) AddIfNotExistFuncLock

```go
func (set *IntSet) AddIfNotExistFuncLock(item int, f func() bool) bool
```

AddIfNotExistFuncLock checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set and function `f` returns true, or else it does nothing and returns false.

​	AddIfNotExistFuncLock 检查集合中是否存在项，将项添加到集合中，如果集合中不存在项，则返回 true，函数 `f` 返回 true，否则不执行任何操作并返回 false。

Note that, the function `f` is executed without writing lock.

​	请注意，该函数 `f` 是在没有写入锁的情况下执行的。

##### Example

``` go
```

#### (*IntSet) Clear

```go
func (set *IntSet) Clear()
```

Clear deletes all items of the set.

​	“清除”（Clear） 将删除集合中的所有项目。

##### Example

``` go
```

#### (*IntSet) Complement

```go
func (set *IntSet) Complement(full *IntSet) (newSet *IntSet)
```

Complement returns a new set which is the complement from `set` to `full`. Which means, all the items in `newSet` are in `full` and not in `set`.

​	补码返回一个新集合，该集合是从 `set` 到 `full` 的补码。这意味着，中 `newSet` 的所有项目都在 中 `full` ，而不是在 `set` 中。

It returns the difference between `full` and `set` if the given set `full` is not the full set of `set`.

​	如果给定的集合 `full` 不是 的全 `set` 集，则返回 和 `set` 之间的 `full` 差值。

##### Example

``` go
```

#### (*IntSet) Contains

```go
func (set *IntSet) Contains(item int) bool
```

Contains checks whether the set contains `item`.

​	包含检查集合是否包含 `item` 。

##### Example

``` go
```

#### (*IntSet) DeepCopy

```go
func (set *IntSet) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*IntSet) Diff

```go
func (set *IntSet) Diff(others ...*IntSet) (newSet *IntSet)
```

Diff returns a new set which is the difference set from `set` to `other`. Which means, all the items in `newSet` are in `set` but not in `other`.

​	Diff 返回一个新集合，该集合是 到 `set` `other` 的差值集。这意味着，中 `newSet` 的所有项目都在 中 `set` ，但不在 `other` 中。

##### Example

``` go
```

#### (*IntSet) Equal

```go
func (set *IntSet) Equal(other *IntSet) bool
```

Equal checks whether the two sets equal.

​	相等检查两组是否相等。

##### Example

``` go
```

#### (*IntSet) Intersect

```go
func (set *IntSet) Intersect(others ...*IntSet) (newSet *IntSet)
```

Intersect returns a new set which is the intersection from `set` to `other`. Which means, all the items in `newSet` are in `set` and also in `other`.

​	Intersect 返回一个新集合，该集合是 的 `set` `other` 交集。这意味着，中 `newSet` 的所有项目都在 和 `set` 中 `other` 。

##### Example

``` go
```

#### (*IntSet) IsSubsetOf

```go
func (set *IntSet) IsSubsetOf(other *IntSet) bool
```

IsSubsetOf checks whether the current set is a sub-set of `other`.

​	IsSubsetOf 检查当前集是否是 的 `other` 子集。

##### Example

``` go
```

#### (*IntSet) Iterator

```go
func (set *IntSet) Iterator(f func(v int) bool)
```

Iterator iterates the set readonly with given callback function `f`, if `f` returns true then continue iterating; or false to stop.

​	迭代器使用给定的回调函数 `f` 迭代集 readonly ，如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*IntSet) Join

```go
func (set *IntSet) Join(glue string) string
```

Join joins items with a string `glue`.

​	Join 使用字符串 `glue` 连接项目。

##### Example

``` go
```

#### (*IntSet) LockFunc

```go
func (set *IntSet) LockFunc(f func(m map[int]struct{}))
```

LockFunc locks writing with callback function `f`.

​	LockFunc 使用回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (IntSet) MarshalJSON

```go
func (set IntSet) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*IntSet) Merge

```go
func (set *IntSet) Merge(others ...*IntSet) *IntSet
```

Merge adds items from `others` sets into `set`.

​	Merge 将 `others` 集合中的项添加到 `set` .

##### Example

``` go
```

#### (*IntSet) Pop

```go
func (set *IntSet) Pop() int
```

Pop randomly pops an item from set.

​	Pop 会随机弹出集合中的项目。

##### Example

``` go
```

#### (*IntSet) Pops

```go
func (set *IntSet) Pops(size int) []int
```

Pops randomly pops `size` items from set. It returns all items if size == -1.

​	弹出会随机弹出集合 `size` 中的物品。如果大小 == -1，则返回所有项目。

##### Example

``` go
```

#### (*IntSet) RLockFunc

```go
func (set *IntSet) RLockFunc(f func(m map[int]struct{}))
```

RLockFunc locks reading with callback function `f`.

​	RLockFunc 使用回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*IntSet) Remove

```go
func (set *IntSet) Remove(item int)
```

Remove deletes `item` from set.

​	从集中删除删除 `item` 。

##### Example

``` go
```

#### (*IntSet) Size

```go
func (set *IntSet) Size() int
```

Size returns the size of the set.

​	size 返回集合的大小。

##### Example

``` go
```

#### (*IntSet) Slice

```go
func (set *IntSet) Slice() []int
```

Slice returns the an of items of the set as slice.

​	Slice 将集合的项的 an 作为 slice 返回。

##### Example

``` go
```

#### (*IntSet) String

```go
func (set *IntSet) String() string
```

String returns items as a string, which implements like json.Marshal does.

​	String 以字符串形式返回项目，其实现方式类似于 json。元帅做到了。

##### Example

``` go
```

#### (*IntSet) Sum

```go
func (set *IntSet) Sum() (sum int)
```

Sum sums items. Note: The items should be converted to int type, or you’d get a result that you unexpected.

​	总和项目。注意：这些项目应转换为 int 类型，否则您将得到意想不到的结果。

##### Example

``` go
```

#### (*IntSet) Union

```go
func (set *IntSet) Union(others ...*IntSet) (newSet *IntSet)
```

Union returns a new set which is the union of `set` and `other`. Which means, all the items in `newSet` are in `set` or in `other`.

​	Union 返回一个新集合，它是 `set` 和 `other` 的并集。这意味着，中 `newSet` 的所有项目都在 或 `set` 中 `other` 。

##### Example

``` go
```

#### (*IntSet) UnmarshalJSON

```go
func (set *IntSet) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*IntSet) UnmarshalValue

```go
func (set *IntSet) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for set.

​	UnmarshalValue 是一个接口实现，用于为 set 设置任何类型的值。

##### Example

``` go
```

#### (*IntSet) Walk

```go
func (set *IntSet) Walk(f func(item int) int) *IntSet
```

Walk applies a user supplied function `f` to every item of set.

​	Walk 将用户提供的功能 `f` 应用于集合的每个项目。

##### Example

``` go
```

### type Set

```go
type Set struct {
	// contains filtered or unexported fields
}
```

#### func New

```go
func New(safe ...bool) *Set
```

New create and returns a new set, which contains un-repeated items. The parameter `safe` is used to specify whether using set in concurrent-safety, which is false in default.

​	新建并返回一个新集合，其中包含未重复的项目。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 set，默认为 false。

#### func NewFrom

```go
func NewFrom(items interface{}, safe ...bool) *Set
```

NewFrom returns a new set from `items`. Parameter `items` can be either a variable of any type, or a slice.

​	NewFrom 返回一个新 `items` 集合。参数 `items` 可以是任何类型的变量，也可以是切片。

##### Example

``` go
```

#### func NewSet

```go
func NewSet(safe ...bool) *Set
```

NewSet create and returns a new set, which contains un-repeated items. Also see New.

​	NewSet 创建并返回一个新集合，其中包含未重复的项。另请参阅新建。

#### (*Set) Add

```go
func (set *Set) Add(items ...interface{})
```

Add adds one or multiple items to the set.

​	“添加”会将一个或多个项目添加到集合中。

#### (*Set) AddIfNotExist

```go
func (set *Set) AddIfNotExist(item interface{}) bool
```

AddIfNotExist checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set, or else it does nothing and returns false.

​	AddIfNotExist 检查集合中是否存在项，将项添加到集合中，如果集合中不存在项，则返回 true，否则不执行任何操作并返回 false。

Note that, if `item` is nil, it does nothing and returns false.

​	请注意，如果 `item` 为 nil，则不执行任何操作并返回 false。

##### Example

``` go
```

#### (*Set) AddIfNotExistFunc

```go
func (set *Set) AddIfNotExistFunc(item interface{}, f func() bool) bool
```

AddIfNotExistFunc checks whether item exists in the set, it adds the item to set and returns true if it does not exist in the set and function `f` returns true, or else it does nothing and returns false.

​	AddIfNotExistFunc 检查集合中是否存在项，将项添加到集合中，如果集合中不存在项，则返回 true，函数 `f` 返回 true，否则不执行任何操作并返回 false。

Note that, if `item` is nil, it does nothing and returns false. The function `f` is executed without writing lock.

​	请注意，如果 `item` 为 nil，则不执行任何操作并返回 false。该函数 `f` 在没有写入锁的情况下执行。

#### (*Set) AddIfNotExistFuncLock

```go
func (set *Set) AddIfNotExistFuncLock(item interface{}, f func() bool) bool
```

AddIfNotExistFuncLock checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set and function `f` returns true, or else it does nothing and returns false.

​	AddIfNotExistFuncLock 检查集合中是否存在项，将项添加到集合中，如果集合中不存在项，则返回 true，函数 `f` 返回 true，否则不执行任何操作并返回 false。

Note that, if `item` is nil, it does nothing and returns false. The function `f` is executed within writing lock.

​	请注意，如果 `item` 为 nil，则不执行任何操作并返回 false。该函数 `f` 在写入锁中执行。

#### (*Set) Clear

```go
func (set *Set) Clear()
```

Clear deletes all items of the set.

​	“清除”（Clear） 将删除集合中的所有项目。

#### (*Set) Complement

```go
func (set *Set) Complement(full *Set) (newSet *Set)
```

Complement returns a new set which is the complement from `set` to `full`. Which means, all the items in `newSet` are in `full` and not in `set`.

​	补码返回一个新集合，该集合是从 `set` 到 `full` 的补码。这意味着，中 `newSet` 的所有项目都在 中 `full` ，而不是在 `set` 中。

It returns the difference between `full` and `set` if the given set `full` is not the full set of `set`.

​	如果给定的集合 `full` 不是 的全 `set` 集，则返回 和 `set` 之间的 `full` 差值。

##### Example

``` go
```

#### (*Set) Contains

```go
func (set *Set) Contains(item interface{}) bool
```

Contains checks whether the set contains `item`.

​	包含检查集合是否包含 `item` 。

##### Example

``` go
```

#### (*Set) DeepCopy

```go
func (set *Set) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*Set) Diff

```go
func (set *Set) Diff(others ...*Set) (newSet *Set)
```

Diff returns a new set which is the difference set from `set` to `others`. Which means, all the items in `newSet` are in `set` but not in `others`.

​	Diff 返回一个新集合，该集合是 到 `set` `others` 的差值集。这意味着，中 `newSet` 的所有项目都在 中 `set` ，但不在 `others` 中。

##### Example

``` go
```

#### (*Set) Equal

```go
func (set *Set) Equal(other *Set) bool
```

Equal checks whether the two sets equal.

​	相等检查两组是否相等。

#### (*Set) Intersect

```go
func (set *Set) Intersect(others ...*Set) (newSet *Set)
```

Intersect returns a new set which is the intersection from `set` to `others`. Which means, all the items in `newSet` are in `set` and also in `others`.

​	Intersect 返回一个新集合，该集合是 的 `set` `others` 交集。这意味着，中 `newSet` 的所有项目都在 和 `set` 中 `others` 。

##### Example

``` go
```

#### (*Set) IsSubsetOf

```go
func (set *Set) IsSubsetOf(other *Set) bool
```

IsSubsetOf checks whether the current set is a sub-set of `other`.

​	IsSubsetOf 检查当前集是否是 的 `other` 子集。

##### Example

``` go
```

#### (*Set) Iterator

```go
func (set *Set) Iterator(f func(v interface{}) bool)
```

Iterator iterates the set readonly with given callback function `f`, if `f` returns true then continue iterating; or false to stop.

​	迭代器使用给定的回调函数 `f` 迭代集 readonly ，如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*Set) Join

```go
func (set *Set) Join(glue string) string
```

Join joins items with a string `glue`.

​	Join 使用字符串 `glue` 连接项目。

##### Example

``` go
```

#### (*Set) LockFunc

```go
func (set *Set) LockFunc(f func(m map[interface{}]struct{}))
```

LockFunc locks writing with callback function `f`.

​	LockFunc 使用回调函数 `f` 锁定写入。

#### (Set) MarshalJSON

```go
func (set Set) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Set) Merge

```go
func (set *Set) Merge(others ...*Set) *Set
```

Merge adds items from `others` sets into `set`.

​	Merge 将 `others` 集合中的项添加到 `set` .

#### (*Set) Pop

```go
func (set *Set) Pop() interface{}
```

Pop randomly pops an item from set.

​	Pop 会随机弹出集合中的项目。

##### Example

``` go
```

#### (*Set) Pops

```go
func (set *Set) Pops(size int) []interface{}
```

Pops randomly pops `size` items from set. It returns all items if size == -1.

​	弹出会随机弹出集合 `size` 中的物品。如果大小 == -1，则返回所有项目。

##### Example

``` go
```

#### (*Set) RLockFunc

```go
func (set *Set) RLockFunc(f func(m map[interface{}]struct{}))
```

RLockFunc locks reading with callback function `f`.

​	RLockFunc 使用回调函数 `f` 锁定读取。

#### (*Set) Remove

```go
func (set *Set) Remove(item interface{})
```

Remove deletes `item` from set.

​	从集中删除删除 `item` 。

#### (*Set) Size

```go
func (set *Set) Size() int
```

Size returns the size of the set.

​	size 返回集合的大小。

#### (*Set) Slice

```go
func (set *Set) Slice() []interface{}
```

Slice returns all items of the set as slice.

​	Slice 将集合的所有项目作为 slice 返回。

#### (*Set) String

```go
func (set *Set) String() string
```

String returns items as a string, which implements like json.Marshal does.

​	String 以字符串形式返回项目，其实现方式类似于 json。元帅做到了。

#### (*Set) Sum

```go
func (set *Set) Sum() (sum int)
```

Sum sums items. Note: The items should be converted to int type, or you’d get a result that you unexpected.

​	总和项目。注意：这些项目应转换为 int 类型，否则您将得到意想不到的结果。

#### (*Set) Union

```go
func (set *Set) Union(others ...*Set) (newSet *Set)
```

Union returns a new set which is the union of `set` and `others`. Which means, all the items in `newSet` are in `set` or in `others`.

​	Union 返回一个新集合，它是 `set` 和 `others` 的并集。这意味着，中 `newSet` 的所有项目都在 或 `set` 中 `others` 。

##### Example

``` go
```

#### (*Set) UnmarshalJSON

```go
func (set *Set) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Set) UnmarshalValue

```go
func (set *Set) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for set.

​	UnmarshalValue 是一个接口实现，用于为 set 设置任何类型的值。

#### (*Set) Walk

```go
func (set *Set) Walk(f func(item interface{}) interface{}) *Set
```

Walk applies a user supplied function `f` to every item of set.

​	Walk 将用户提供的功能 `f` 应用于集合的每个项目。

### type StrSet

```go
type StrSet struct {
	// contains filtered or unexported fields
}
```

#### func NewStrSet

```go
func NewStrSet(safe ...bool) *StrSet
```

NewStrSet create and returns a new set, which contains un-repeated items. The parameter `safe` is used to specify whether using set in concurrent-safety, which is false in default.

​	NewStrSet 创建并返回一个新集合，其中包含未重复的项。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 set，默认为 false。

##### Example

``` go
```

#### func NewStrSetFrom

```go
func NewStrSetFrom(items []string, safe ...bool) *StrSet
```

NewStrSetFrom returns a new set from `items`.

​	NewStrSetFrom 从 返回一个新 `items` 集。

##### Example

``` go
```

#### (*StrSet) Add

```go
func (set *StrSet) Add(item ...string)
```

Add adds one or multiple items to the set.

​	“添加”会将一个或多个项目添加到集合中。

##### Example

``` go
```

#### (*StrSet) AddIfNotExist

```go
func (set *StrSet) AddIfNotExist(item string) bool
```

AddIfNotExist checks whether item exists in the set, it adds the item to set and returns true if it does not exist in the set, or else it does nothing and returns false.

​	AddIfNotExist 检查集合中是否存在项，将项添加到集合中，如果集合中不存在项，则返回 true，否则不执行任何操作并返回 false。

##### Example

``` go
```

#### (*StrSet) AddIfNotExistFunc

```go
func (set *StrSet) AddIfNotExistFunc(item string, f func() bool) bool
```

AddIfNotExistFunc checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set and function `f` returns true, or else it does nothing and returns false.

​	AddIfNotExistFunc 检查集合中是否存在项，将项添加到集合中，如果集合中不存在项，则返回 true，函数 `f` 返回 true，否则不执行任何操作并返回 false。

Note that, the function `f` is executed without writing lock.

​	请注意，该函数 `f` 是在没有写入锁的情况下执行的。

##### Example

``` go
```

#### (*StrSet) AddIfNotExistFuncLock

```go
func (set *StrSet) AddIfNotExistFuncLock(item string, f func() bool) bool
```

AddIfNotExistFuncLock checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set and function `f` returns true, or else it does nothing and returns false.

​	AddIfNotExistFuncLock 检查集合中是否存在项，将项添加到集合中，如果集合中不存在项，则返回 true，函数 `f` 返回 true，否则不执行任何操作并返回 false。

Note that, the function `f` is executed without writing lock.

​	请注意，该函数 `f` 是在没有写入锁的情况下执行的。

##### Example

``` go
```

#### (*StrSet) Clear

```go
func (set *StrSet) Clear()
```

Clear deletes all items of the set.

​	“清除”（Clear） 将删除集合中的所有项目。

##### Example

``` go
```

#### (*StrSet) Complement

```go
func (set *StrSet) Complement(full *StrSet) (newSet *StrSet)
```

Complement returns a new set which is the complement from `set` to `full`. Which means, all the items in `newSet` are in `full` and not in `set`.

​	补码返回一个新集合，该集合是从 `set` 到 `full` 的补码。这意味着，中 `newSet` 的所有项目都在 中 `full` ，而不是在 `set` 中。

It returns the difference between `full` and `set` if the given set `full` is not the full set of `set`.

​	如果给定的集合 `full` 不是 的全 `set` 集，则返回 和 `set` 之间的 `full` 差值。

##### Example

``` go
```

#### (*StrSet) Contains

```go
func (set *StrSet) Contains(item string) bool
```

Contains checks whether the set contains `item`.

​	包含检查集合是否包含 `item` 。

##### Example

``` go
```

#### (*StrSet) ContainsI

```go
func (set *StrSet) ContainsI(item string) bool
```

ContainsI checks whether a value exists in the set with case-insensitively. Note that it internally iterates the whole set to do the comparison with case-insensitively.

​	ContainsI 使用不区分大小写的方式检查集合中是否存在值。请注意，它会在内部迭代整个集合，以不区分大小写地进行比较。

##### Example

``` go
```

#### (*StrSet) DeepCopy

```go
func (set *StrSet) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*StrSet) Diff

```go
func (set *StrSet) Diff(others ...*StrSet) (newSet *StrSet)
```

Diff returns a new set which is the difference set from `set` to `other`. Which means, all the items in `newSet` are in `set` but not in `other`.

​	Diff 返回一个新集合，该集合是 到 `set` `other` 的差值集。这意味着，中 `newSet` 的所有项目都在 中 `set` ，但不在 `other` 中。

##### Example

``` go
```

#### (*StrSet) Equal

```go
func (set *StrSet) Equal(other *StrSet) bool
```

Equal checks whether the two sets equal.

​	相等检查两组是否相等。

##### Example

``` go
```

#### (*StrSet) Intersect

```go
func (set *StrSet) Intersect(others ...*StrSet) (newSet *StrSet)
```

Intersect returns a new set which is the intersection from `set` to `other`. Which means, all the items in `newSet` are in `set` and also in `other`.

​	Intersect 返回一个新集合，该集合是 的 `set` `other` 交集。这意味着，中 `newSet` 的所有项目都在 和 `set` 中 `other` 。

##### Example

``` go
```

#### (*StrSet) IsSubsetOf

```go
func (set *StrSet) IsSubsetOf(other *StrSet) bool
```

IsSubsetOf checks whether the current set is a sub-set of `other`.

​	IsSubsetOf 检查当前集是否是 的 `other` 子集。

##### Example

``` go
```

#### (*StrSet) Iterator

```go
func (set *StrSet) Iterator(f func(v string) bool)
```

Iterator iterates the set readonly with given callback function `f`, if `f` returns true then continue iterating; or false to stop.

​	迭代器使用给定的回调函数 `f` 迭代集 readonly ，如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*StrSet) Join

```go
func (set *StrSet) Join(glue string) string
```

Join joins items with a string `glue`.

​	Join 使用字符串 `glue` 连接项目。

##### Example

``` go
```

#### (*StrSet) LockFunc

```go
func (set *StrSet) LockFunc(f func(m map[string]struct{}))
```

LockFunc locks writing with callback function `f`.

​	LockFunc 使用回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (StrSet) MarshalJSON

```go
func (set StrSet) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*StrSet) Merge

```go
func (set *StrSet) Merge(others ...*StrSet) *StrSet
```

Merge adds items from `others` sets into `set`.

​	Merge 将 `others` 集合中的项添加到 `set` .

##### Example

``` go
```

#### (*StrSet) Pop

```go
func (set *StrSet) Pop() string
```

Pop randomly pops an item from set.

​	Pop 会随机弹出集合中的项目。

##### Example

``` go
```

#### (*StrSet) Pops

```go
func (set *StrSet) Pops(size int) []string
```

Pops randomly pops `size` items from set. It returns all items if size == -1.

​	弹出会随机弹出集合 `size` 中的物品。如果大小 == -1，则返回所有项目。

##### Example

``` go
```

#### (*StrSet) RLockFunc

```go
func (set *StrSet) RLockFunc(f func(m map[string]struct{}))
```

RLockFunc locks reading with callback function `f`.

​	RLockFunc 使用回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*StrSet) Remove

```go
func (set *StrSet) Remove(item string)
```

Remove deletes `item` from set.

​	从集中删除删除 `item` 。

##### Example

``` go
```

#### (*StrSet) Size

```go
func (set *StrSet) Size() int
```

Size returns the size of the set.

​	size 返回集合的大小。

##### Example

``` go
```

#### (*StrSet) Slice

```go
func (set *StrSet) Slice() []string
```

Slice returns the an of items of the set as slice.

​	Slice 将集合的项的 an 作为 slice 返回。

##### Example

``` go
```

#### (*StrSet) String

```go
func (set *StrSet) String() string
```

String returns items as a string, which implements like json.Marshal does.

​	String 以字符串形式返回项目，其实现方式类似于 json。元帅做到了。

##### Example

``` go
```

#### (*StrSet) Sum

```go
func (set *StrSet) Sum() (sum int)
```

Sum sums items. Note: The items should be converted to int type, or you’d get a result that you unexpected.

​	总和项目。注意：这些项目应转换为 int 类型，否则您将得到意想不到的结果。

##### Example

``` go
```

#### (*StrSet) Union

```go
func (set *StrSet) Union(others ...*StrSet) (newSet *StrSet)
```

Union returns a new set which is the union of `set` and `other`. Which means, all the items in `newSet` are in `set` or in `other`.

​	Union 返回一个新集合，它是 `set` 和 `other` 的并集。这意味着，中 `newSet` 的所有项目都在 或 `set` 中 `other` 。

##### Example

``` go
```

#### (*StrSet) UnmarshalJSON

```go
func (set *StrSet) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*StrSet) UnmarshalValue

```go
func (set *StrSet) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for set.

​	UnmarshalValue 是一个接口实现，用于为 set 设置任何类型的值。

##### Example

``` go
```

#### (*StrSet) Walk

```go
func (set *StrSet) Walk(f func(item string) string) *StrSet
```

Walk applies a user supplied function `f` to every item of set.

​	Walk 将用户提供的功能 `f` 应用于集合的每个项目。

Example Walk

​	示例步行

Walk applies a user supplied function `f` to every item of set.

​	Walk 将用户提供的功能 `f` 应用于集合的每个项目。

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	var (
		set    gset.StrSet
		names  = g.SliceStr{"user", "user_detail"}
		prefix = "gf_"
	)
	set.Add(names...)
	// Add prefix for given table names.
	set.Walk(func(item string) string {
		return prefix + item
	})
	fmt.Println(set.Slice())

	// May Output:
	// [gf_user gf_user_detail]
}

Output:
```

