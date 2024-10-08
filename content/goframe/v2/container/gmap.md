+++
title = "gmap"
date = 2024-03-21T17:44:36+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gmap](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gmap)

### Overview 概述

Package gmap provides most commonly used map container which also support concurrent-safe/unsafe switch feature.

​	软件包 gmap 提供了最常用的地图容器，它还支持并发安全/不安全交换功能。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type AnyAnyMap

```go
type AnyAnyMap struct {
	// contains filtered or unexported fields
}
```

AnyAnyMap wraps map type `map[interface{}]interface{}` and provides more map features.

​	AnyAnyMap 包装地图类型 `map[interface{}]interface{}` 并提供更多地图功能。

#### func NewAnyAnyMap

```go
func NewAnyAnyMap(safe ...bool) *AnyAnyMap
```

NewAnyAnyMap creates and returns an empty hash map. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

​	NewAnyAnyMap 创建并返回一个空的哈希映射。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 map，默认为 false。

##### Example

``` go
```

#### func NewAnyAnyMapFrom

```go
func NewAnyAnyMapFrom(data map[interface{}]interface{}, safe ...bool) *AnyAnyMap
```

NewAnyAnyMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

​	NewAnyAnyMapFrom 创建并返回来自给定映射 `data` 的哈希映射。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。

##### Example

``` go
```

#### (*AnyAnyMap) Clear

```go
func (m *AnyAnyMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

​	清除会删除地图的所有数据，它将重新制作一个新的基础数据地图。

##### Example

``` go
```

#### (*AnyAnyMap) Clone

```go
func (m *AnyAnyMap) Clone(safe ...bool) *AnyAnyMap
```

Clone returns a new hash map with copy of current map data.

​	克隆将返回一个新的哈希映射，其中包含当前映射数据的副本。

##### Example

``` go
```

#### (*AnyAnyMap) Contains

```go
func (m *AnyAnyMap) Contains(key interface{}) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

​	包含检查密钥是否存在。如果存在， `key` 则返回 true，否则返回 false。

##### Example

``` go
```

#### (*AnyAnyMap) DeepCopy

```go
func (m *AnyAnyMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*AnyAnyMap) Diff

```go
func (m *AnyAnyMap) Diff(other *AnyAnyMap) (addedKeys, removedKeys, updatedKeys []interface{})
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

​	Diff 将当前映射 `m` 与映射 `other` 进行比较，并返回它们的不同键。返回 `addedKeys` 的键在 map `m` 中但不在 map `other` 中。返回 `removedKeys` 的键在 map `other` 中但不在 map `m` 中。返回 `updatedKeys` 的键既在 map `m` 中， `other` 又在 但它们的值中且不相等 （ `!=` ）。

#### (*AnyAnyMap) FilterEmpty

```go
func (m *AnyAnyMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, “”, len(slice/map/chan) == 0 are considered empty.

​	FilterEmpty 删除所有值为空的键值对。像 0， nil， false， “”， len（slice/map/chan） == 0 这样的值被视为空。

##### Example

``` go
```

#### (*AnyAnyMap) FilterNil

```go
func (m *AnyAnyMap) FilterNil()
```

FilterNil deletes all key-value pair of which the value is nil.

​	FilterNil 删除值为 nil 的所有键值对。

##### Example

``` go
```

#### (*AnyAnyMap) Flip

```go
func (m *AnyAnyMap) Flip()
```

Flip exchanges key-value of the map to value-key.

​	Flip 将映射的键值交换为值键。

##### Example

``` go
```

#### (*AnyAnyMap) Get

```go
func (m *AnyAnyMap) Get(key interface{}) (value interface{})
```

Get returns the value by given `key`.

​	Get 返回给定 `key` 的值。

##### Example

``` go
```

#### (*AnyAnyMap) GetOrSet

```go
func (m *AnyAnyMap) GetOrSet(key interface{}, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

##### Example

``` go
```

#### (*AnyAnyMap) GetOrSetFunc

```go
func (m *AnyAnyMap) GetOrSetFunc(key interface{}, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

##### Example

``` go
```

#### (*AnyAnyMap) GetOrSetFuncLock

```go
func (m *AnyAnyMap) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*AnyAnyMap) GetVar

```go
func (m *AnyAnyMap) GetVar(key interface{}) *gvar.Var
```

GetVar returns a Var with the value by given `key`. The returned Var is un-concurrent safe.

​	GetVar 返回一个 Var，其值为 给定 `key` 。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*AnyAnyMap) GetVarOrSet

```go
func (m *AnyAnyMap) GetVarOrSet(key interface{}, value interface{}) *gvar.Var
```

GetVarOrSet returns a Var with result from GetOrSet. The returned Var is un-concurrent safe.

​	GetVarOrSet 返回一个 Var，其中包含来自 GetOrSet 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*AnyAnyMap) GetVarOrSetFunc

```go
func (m *AnyAnyMap) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a Var with result from GetOrSetFunc. The returned Var is un-concurrent safe.

​	GetVarOrSetFunc 返回一个 Var，其中包含来自 GetOrSetFunc 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*AnyAnyMap) GetVarOrSetFuncLock

```go
func (m *AnyAnyMap) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock. The returned Var is un-concurrent safe.

​	GetVarOrSetFuncLock 返回一个 Var，其中包含来自 GetOrSetFuncLock 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*AnyAnyMap) IsEmpty

```go
func (m *AnyAnyMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

​	IsEmpty 检查地图是否为空。如果 map 为空，则返回 true，否则返回 false。

##### Example

``` go
```

#### (*AnyAnyMap) IsSubOf

```go
func (m *AnyAnyMap) IsSubOf(other *AnyAnyMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

​	IsSubOf 检查当前映射是否是 的 `other` 子映射。

#### (*AnyAnyMap) Iterator

```go
func (m *AnyAnyMap) Iterator(f func(k interface{}, v interface{}) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	迭代器使用自定义回调函数 `f` 迭代哈希映射只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*AnyAnyMap) Keys

```go
func (m *AnyAnyMap) Keys() []interface{}
```

Keys returns all keys of the map as a slice.

​	Keys 将地图的所有键作为切片返回。

##### Example

``` go
```

#### (*AnyAnyMap) LockFunc

```go
func (m *AnyAnyMap) LockFunc(f func(m map[interface{}]interface{}))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

​	LockFunc 在 RWMutex.Lock 中使用给定的回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (*AnyAnyMap) Map

```go
func (m *AnyAnyMap) Map() map[interface{}]interface{}
```

Map returns the underlying data map. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Map 返回基础数据映射。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

##### Example

``` go
```

#### (*AnyAnyMap) MapCopy

```go
func (m *AnyAnyMap) MapCopy() map[interface{}]interface{}
```

MapCopy returns a shallow copy of the underlying data of the hash map.

​	MapCopy 返回哈希映射基础数据的浅层副本。

##### Example

``` go
```

#### (*AnyAnyMap) MapStrAny

```go
func (m *AnyAnyMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回地图基础数据的副本。

##### Example

``` go
```

#### (AnyAnyMap) MarshalJSON

```go
func (m AnyAnyMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*AnyAnyMap) Merge

```go
func (m *AnyAnyMap) Merge(other *AnyAnyMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

​	合并合并两个哈希映射。 `other` 地图将合并到地图 `m` 中。

##### Example

``` go
```

#### (*AnyAnyMap) Pop

```go
func (m *AnyAnyMap) Pop() (key, value interface{})
```

Pop retrieves and deletes an item from the map.

​	Pop 从地图中检索和删除项目。

##### Example

``` go
```

#### (*AnyAnyMap) Pops

```go
func (m *AnyAnyMap) Pops(size int) map[interface{}]interface{}
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

​	Pops 从地图中检索和删除 `size` 项目。如果大小 == -1，则返回所有项目。

##### Example

``` go
```

#### (*AnyAnyMap) RLockFunc

```go
func (m *AnyAnyMap) RLockFunc(f func(m map[interface{}]interface{}))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

​	RLockFunc 在 RWMutex.RLock 中使用给定的回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*AnyAnyMap) Remove

```go
func (m *AnyAnyMap) Remove(key interface{}) (value interface{})
```

Remove deletes value from map by given `key`, and return this deleted value.

​	按给定 `key` 从映射中删除删除值，并返回此删除值。

##### Example

``` go
```

#### (*AnyAnyMap) Removes

```go
func (m *AnyAnyMap) Removes(keys []interface{})
```

Removes batch deletes values of the map by keys.

​	删除按键批量删除映射的值。

##### Example

``` go
```

#### (*AnyAnyMap) Replace

```go
func (m *AnyAnyMap) Replace(data map[interface{}]interface{})
```

Replace the data of the map with given `data`.

​	将地图的数据替换为给定 `data` 的 .

##### Example

``` go
```

#### (*AnyAnyMap) Search

```go
func (m *AnyAnyMap) Search(key interface{}) (value interface{}, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的 .如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

##### Example

``` go
```

#### (*AnyAnyMap) Set

```go
func (m *AnyAnyMap) Set(key interface{}, value interface{})
```

Set sets key-value to the hash map.

​	Set 将 key-value 设置为哈希映射。

##### Example

``` go
```

#### (*AnyAnyMap) SetIfNotExist

```go
func (m *AnyAnyMap) SetIfNotExist(key interface{}, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*AnyAnyMap) SetIfNotExistFunc

```go
func (m *AnyAnyMap) SetIfNotExistFunc(key interface{}, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*AnyAnyMap) SetIfNotExistFuncLock

```go
func (m *AnyAnyMap) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*AnyAnyMap) Sets

```go
func (m *AnyAnyMap) Sets(data map[interface{}]interface{})
```

Sets batch sets key-values to the hash map.

​	将批处理设置键值设置为哈希映射。

##### Example

``` go
```

#### (*AnyAnyMap) Size

```go
func (m *AnyAnyMap) Size() int
```

Size returns the size of the map.

​	Size 返回地图的大小。

##### Example

``` go
```

#### (*AnyAnyMap) String

```go
func (m *AnyAnyMap) String() string
```

String returns the map as a string.

​	String 以字符串形式返回映射。

##### Example

``` go
```

#### (*AnyAnyMap) UnmarshalJSON

```go
func (m *AnyAnyMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*AnyAnyMap) UnmarshalValue

```go
func (m *AnyAnyMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

​	UnmarshalValue 是一个接口实现，用于为 map 设置任何类型的值。

##### Example

``` go
```

#### (*AnyAnyMap) Values

```go
func (m *AnyAnyMap) Values() []interface{}
```

Values returns all values of the map as a slice.

​	Values 以切片的形式返回地图的所有值。

##### Example

``` go
```

### type HashMap

```go
type HashMap = AnyAnyMap // HashMap is alias of AnyAnyMap.
```

### type IntAnyMap

```go
type IntAnyMap struct {
	// contains filtered or unexported fields
}
```

IntAnyMap implements map[int]interface{} with RWMutex that has switch.

​	IntAnyMap 使用具有 switch 的 RWMutex 实现 map[int]interface{}。

#### func NewIntAnyMap

```go
func NewIntAnyMap(safe ...bool) *IntAnyMap
```

NewIntAnyMap returns an empty IntAnyMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

​	NewIntAnyMap 返回一个空的 IntAnyMap 对象。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 map，默认为 false。

##### Example

``` go
```

#### func NewIntAnyMapFrom

```go
func NewIntAnyMapFrom(data map[int]interface{}, safe ...bool) *IntAnyMap
```

NewIntAnyMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

​	NewIntAnyMapFrom 从给定的映射 `data` 创建并返回哈希映射。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。

##### Example

``` go
```

#### (*IntAnyMap) Clear

```go
func (m *IntAnyMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

​	清除会删除地图的所有数据，它将重新制作一个新的基础数据地图。

##### Example

``` go
```

#### (*IntAnyMap) Clone

```go
func (m *IntAnyMap) Clone() *IntAnyMap
```

Clone returns a new hash map with copy of current map data.

​	克隆将返回一个新的哈希映射，其中包含当前映射数据的副本。

##### Example

``` go
```

#### (*IntAnyMap) Contains

```go
func (m *IntAnyMap) Contains(key int) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

​	包含检查密钥是否存在。如果存在， `key` 则返回 true，否则返回 false。

##### Example

``` go
```

#### (*IntAnyMap) DeepCopy

```go
func (m *IntAnyMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*IntAnyMap) Diff

```go
func (m *IntAnyMap) Diff(other *IntAnyMap) (addedKeys, removedKeys, updatedKeys []int)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

​	Diff 将当前映射 `m` 与映射 `other` 进行比较，并返回它们的不同键。返回 `addedKeys` 的键在 map `m` 中但不在 map `other` 中。返回 `removedKeys` 的键在 map `other` 中但不在 map `m` 中。返回 `updatedKeys` 的键既在 map `m` 中， `other` 又在 但它们的值中且不相等 （ `!=` ）。

#### (*IntAnyMap) FilterEmpty

```go
func (m *IntAnyMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, “”, len(slice/map/chan) == 0 are considered empty.

​	FilterEmpty 删除所有值为空的键值对。像 0， nil， false， “”， len（slice/map/chan） == 0 这样的值被视为空。

##### Example

``` go
```

#### (*IntAnyMap) FilterNil

```go
func (m *IntAnyMap) FilterNil()
```

FilterNil deletes all key-value pair of which the value is nil.

​	FilterNil 删除值为 nil 的所有键值对。

##### Example

``` go
```

#### (*IntAnyMap) Flip

```go
func (m *IntAnyMap) Flip()
```

Flip exchanges key-value of the map to value-key.

​	Flip 将映射的键值交换为值键。

##### Example

``` go
```

#### (*IntAnyMap) Get

```go
func (m *IntAnyMap) Get(key int) (value interface{})
```

Get returns the value by given `key`.

​	Get 返回给定 `key` 的值。

##### Example

``` go
```

#### (*IntAnyMap) GetOrSet

```go
func (m *IntAnyMap) GetOrSet(key int, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

##### Example

``` go
```

#### (*IntAnyMap) GetOrSetFunc

```go
func (m *IntAnyMap) GetOrSetFunc(key int, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在）并返回此值。

##### Example

``` go
```

#### (*IntAnyMap) GetOrSetFuncLock

```go
func (m *IntAnyMap) GetOrSetFuncLock(key int, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在）并返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*IntAnyMap) GetVar

```go
func (m *IntAnyMap) GetVar(key int) *gvar.Var
```

GetVar returns a Var with the value by given `key`. The returned Var is un-concurrent safe.

​	GetVar 返回一个 Var，其值为 给定 `key` 。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*IntAnyMap) GetVarOrSet

```go
func (m *IntAnyMap) GetVarOrSet(key int, value interface{}) *gvar.Var
```

GetVarOrSet returns a Var with result from GetVarOrSet. The returned Var is un-concurrent safe.

​	GetVarOrSet 返回一个 Var，其中包含来自 GetVarOrSet 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*IntAnyMap) GetVarOrSetFunc

```go
func (m *IntAnyMap) GetVarOrSetFunc(key int, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a Var with result from GetOrSetFunc. The returned Var is un-concurrent safe.

​	GetVarOrSetFunc 返回一个 Var，其中包含来自 GetOrSetFunc 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*IntAnyMap) GetVarOrSetFuncLock

```go
func (m *IntAnyMap) GetVarOrSetFuncLock(key int, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock. The returned Var is un-concurrent safe.

​	GetVarOrSetFuncLock 返回一个 Var，其中包含来自 GetOrSetFuncLock 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*IntAnyMap) IsEmpty

```go
func (m *IntAnyMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

​	IsEmpty 检查地图是否为空。如果 map 为空，则返回 true，否则返回 false。

##### Example

``` go
```

#### (*IntAnyMap) IsSubOf

```go
func (m *IntAnyMap) IsSubOf(other *IntAnyMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

​	IsSubOf 检查当前映射是否是 的 `other` 子映射。

#### (*IntAnyMap) Iterator

```go
func (m *IntAnyMap) Iterator(f func(k int, v interface{}) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	迭代器使用自定义回调函数 `f` 迭代哈希映射只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*IntAnyMap) Keys

```go
func (m *IntAnyMap) Keys() []int
```

Keys returns all keys of the map as a slice.

​	Keys 将地图的所有键作为切片返回。

##### Example

``` go
```

#### (*IntAnyMap) LockFunc

```go
func (m *IntAnyMap) LockFunc(f func(m map[int]interface{}))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

​	LockFunc 在 RWMutex.Lock 中使用给定的回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (*IntAnyMap) Map

```go
func (m *IntAnyMap) Map() map[int]interface{}
```

Map returns the underlying data map. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Map 返回基础数据映射。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

##### Example

``` go
```

#### (*IntAnyMap) MapCopy

```go
func (m *IntAnyMap) MapCopy() map[int]interface{}
```

MapCopy returns a copy of the underlying data of the hash map.

​	MapCopy 返回哈希映射的基础数据的副本。

##### Example

``` go
```

#### (*IntAnyMap) MapStrAny

```go
func (m *IntAnyMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回地图基础数据的副本。

##### Example

``` go
```

#### (IntAnyMap) MarshalJSON

```go
func (m IntAnyMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*IntAnyMap) Merge

```go
func (m *IntAnyMap) Merge(other *IntAnyMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

​	合并合并两个哈希映射。 `other` 地图将合并到地图 `m` 中。

##### Example

``` go
```

#### (*IntAnyMap) Pop

```go
func (m *IntAnyMap) Pop() (key int, value interface{})
```

Pop retrieves and deletes an item from the map.

​	Pop 从地图中检索和删除项目。

##### Example

``` go
```

#### (*IntAnyMap) Pops

```go
func (m *IntAnyMap) Pops(size int) map[int]interface{}
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

​	Pops 从地图中检索和删除 `size` 项目。如果大小 == -1，则返回所有项目。

##### Example

``` go
```

#### (*IntAnyMap) RLockFunc

```go
func (m *IntAnyMap) RLockFunc(f func(m map[int]interface{}))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

​	RLockFunc 在 RWMutex.RLock 中使用给定的回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*IntAnyMap) Remove

```go
func (m *IntAnyMap) Remove(key int) (value interface{})
```

Remove deletes value from map by given `key`, and return this deleted value.

​	按给定 `key` 从映射中删除删除值，并返回此删除值。

##### Example

``` go
```

#### (*IntAnyMap) Removes

```go
func (m *IntAnyMap) Removes(keys []int)
```

Removes batch deletes values of the map by keys.

​	删除按键批量删除映射的值。

##### Example

``` go
```

#### (*IntAnyMap) Replace

```go
func (m *IntAnyMap) Replace(data map[int]interface{})
```

Replace the data of the map with given `data`.

​	将地图的数据替换为给定 `data` 的 .

##### Example

``` go
```

#### (*IntAnyMap) Search

```go
func (m *IntAnyMap) Search(key int) (value interface{}, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的 .如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

##### Example

``` go
```

#### (*IntAnyMap) Set

```go
func (m *IntAnyMap) Set(key int, val interface{})
```

Set sets key-value to the hash map.

​	Set 将 key-value 设置为哈希映射。

##### Example

``` go
```

#### (*IntAnyMap) SetIfNotExist

```go
func (m *IntAnyMap) SetIfNotExist(key int, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*IntAnyMap) SetIfNotExistFunc

```go
func (m *IntAnyMap) SetIfNotExistFunc(key int, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*IntAnyMap) SetIfNotExistFuncLock

```go
func (m *IntAnyMap) SetIfNotExistFuncLock(key int, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*IntAnyMap) Sets

```go
func (m *IntAnyMap) Sets(data map[int]interface{})
```

Sets batch sets key-values to the hash map.

​	将批处理设置键值设置为哈希映射。

##### Example

``` go
```

#### (*IntAnyMap) Size

```go
func (m *IntAnyMap) Size() int
```

Size returns the size of the map.

​	Size 返回地图的大小。

##### Example

``` go
```

#### (*IntAnyMap) String

```go
func (m *IntAnyMap) String() string
```

String returns the map as a string.

​	String 以字符串形式返回映射。

##### Example

``` go
```

#### (*IntAnyMap) UnmarshalJSON

```go
func (m *IntAnyMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*IntAnyMap) UnmarshalValue

```go
func (m *IntAnyMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

​	UnmarshalValue 是一个接口实现，用于为 map 设置任何类型的值。

##### Example

``` go
```

#### (*IntAnyMap) Values

```go
func (m *IntAnyMap) Values() []interface{}
```

Values returns all values of the map as a slice.

​	Values 以切片的形式返回地图的所有值。

##### Example

``` go
```

### type IntIntMap

```go
type IntIntMap struct {
	// contains filtered or unexported fields
}
```

IntIntMap implements map[int]int with RWMutex that has switch.

​	IntIntMap 使用具有 switch 的 RWMutex 实现 map[int]int。

#### func NewIntIntMap

```go
func NewIntIntMap(safe ...bool) *IntIntMap
```

NewIntIntMap returns an empty IntIntMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

​	NewIntIntMap 返回一个空的 IntIntMap 对象。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 map，默认为 false。

##### Example

``` go
```

#### func NewIntIntMapFrom

```go
func NewIntIntMapFrom(data map[int]int, safe ...bool) *IntIntMap
```

NewIntIntMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

​	NewIntIntMapFrom 从给定映射 `data` 创建并返回哈希映射。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。

##### Example

``` go
```

#### (*IntIntMap) Clear

```go
func (m *IntIntMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

​	清除会删除地图的所有数据，它将重新制作一个新的基础数据地图。

##### Example

``` go
```

#### (*IntIntMap) Clone

```go
func (m *IntIntMap) Clone() *IntIntMap
```

Clone returns a new hash map with copy of current map data.

​	克隆将返回一个新的哈希映射，其中包含当前映射数据的副本。

##### Example

``` go
```

#### (*IntIntMap) Contains

```go
func (m *IntIntMap) Contains(key int) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

​	包含检查密钥是否存在。如果存在， `key` 则返回 true，否则返回 false。

##### Example

``` go
```

#### (*IntIntMap) DeepCopy

```go
func (m *IntIntMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*IntIntMap) Diff

```go
func (m *IntIntMap) Diff(other *IntIntMap) (addedKeys, removedKeys, updatedKeys []int)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

​	Diff 将当前映射 `m` 与映射 `other` 进行比较，并返回它们的不同键。返回 `addedKeys` 的键在 map `m` 中但不在 map `other` 中。返回 `removedKeys` 的键在 map `other` 中但不在 map `m` 中。返回 `updatedKeys` 的键既在 map `m` 中， `other` 又在 但它们的值中且不相等 （ `!=` ）。

#### (*IntIntMap) FilterEmpty

```go
func (m *IntIntMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, “”, len(slice/map/chan) == 0 are considered empty.

​	FilterEmpty 删除所有值为空的键值对。像 0， nil， false， “”， len（slice/map/chan） == 0 这样的值被视为空。

##### Example

``` go
```

#### (*IntIntMap) Flip

```go
func (m *IntIntMap) Flip()
```

Flip exchanges key-value of the map to value-key.

​	Flip 将映射的键值交换为值键。

##### Example

``` go
```

#### (*IntIntMap) Get

```go
func (m *IntIntMap) Get(key int) (value int)
```

Get returns the value by given `key`.

​	Get 返回给定 `key` 的值。

##### Example

``` go
```

#### (*IntIntMap) GetOrSet

```go
func (m *IntIntMap) GetOrSet(key int, value int) int
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

##### Example

``` go
```

#### (*IntIntMap) GetOrSetFunc

```go
func (m *IntIntMap) GetOrSetFunc(key int, f func() int) int
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在）并返回此值。

##### Example

``` go
```

#### (*IntIntMap) GetOrSetFuncLock

```go
func (m *IntIntMap) GetOrSetFuncLock(key int, f func() int) int
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在）并返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*IntIntMap) IsEmpty

```go
func (m *IntIntMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

​	IsEmpty 检查地图是否为空。如果 map 为空，则返回 true，否则返回 false。

##### Example

``` go
```

#### (*IntIntMap) IsSubOf

```go
func (m *IntIntMap) IsSubOf(other *IntIntMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

​	IsSubOf 检查当前映射是否是 的 `other` 子映射。

#### (*IntIntMap) Iterator

```go
func (m *IntIntMap) Iterator(f func(k int, v int) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	迭代器使用自定义回调函数 `f` 迭代哈希映射只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*IntIntMap) Keys

```go
func (m *IntIntMap) Keys() []int
```

Keys returns all keys of the map as a slice.

​	Keys 将地图的所有键作为切片返回。

##### Example

``` go
```

#### (*IntIntMap) LockFunc

```go
func (m *IntIntMap) LockFunc(f func(m map[int]int))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

​	LockFunc 在 RWMutex.Lock 中使用给定的回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (*IntIntMap) Map

```go
func (m *IntIntMap) Map() map[int]int
```

Map returns the underlying data map. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Map 返回基础数据映射。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

##### Example

``` go
```

#### (*IntIntMap) MapCopy

```go
func (m *IntIntMap) MapCopy() map[int]int
```

MapCopy returns a copy of the underlying data of the hash map.

​	MapCopy 返回哈希映射的基础数据的副本。

##### Example

``` go
```

#### (*IntIntMap) MapStrAny

```go
func (m *IntIntMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回地图基础数据的副本。

##### Example

``` go
```

#### (IntIntMap) MarshalJSON

```go
func (m IntIntMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*IntIntMap) Merge

```go
func (m *IntIntMap) Merge(other *IntIntMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

​	合并合并两个哈希映射。 `other` 地图将合并到地图 `m` 中。

##### Example

``` go
```

#### (*IntIntMap) Pop

```go
func (m *IntIntMap) Pop() (key, value int)
```

Pop retrieves and deletes an item from the map.

​	Pop 从地图中检索和删除项目。

##### Example

``` go
```

#### (*IntIntMap) Pops

```go
func (m *IntIntMap) Pops(size int) map[int]int
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

​	Pops 从地图中检索和删除 `size` 项目。如果大小 == -1，则返回所有项目。

##### Example

``` go
```

#### (*IntIntMap) RLockFunc

```go
func (m *IntIntMap) RLockFunc(f func(m map[int]int))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

​	RLockFunc 在 RWMutex.RLock 中使用给定的回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*IntIntMap) Remove

```go
func (m *IntIntMap) Remove(key int) (value int)
```

Remove deletes value from map by given `key`, and return this deleted value.

​	按给定 `key` 从映射中删除删除值，并返回此删除值。

##### Example

``` go
```

#### (*IntIntMap) Removes

```go
func (m *IntIntMap) Removes(keys []int)
```

Removes batch deletes values of the map by keys.

​	删除按键批量删除映射的值。

##### Example

``` go
```

#### (*IntIntMap) Replace

```go
func (m *IntIntMap) Replace(data map[int]int)
```

Replace the data of the map with given `data`.

​	将地图的数据替换为给定 `data` 的 .

##### Example

``` go
```

#### (*IntIntMap) Search

```go
func (m *IntIntMap) Search(key int) (value int, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的 .如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

##### Example

``` go
```

#### (*IntIntMap) Set

```go
func (m *IntIntMap) Set(key int, val int)
```

Set sets key-value to the hash map.

​	Set 将 key-value 设置为哈希映射。

##### Example

``` go
```

#### (*IntIntMap) SetIfNotExist

```go
func (m *IntIntMap) SetIfNotExist(key int, value int) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*IntIntMap) SetIfNotExistFunc

```go
func (m *IntIntMap) SetIfNotExistFunc(key int, f func() int) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*IntIntMap) SetIfNotExistFuncLock

```go
func (m *IntIntMap) SetIfNotExistFuncLock(key int, f func() int) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*IntIntMap) Sets

```go
func (m *IntIntMap) Sets(data map[int]int)
```

Sets batch sets key-values to the hash map.

​	将批处理设置键值设置为哈希映射。

##### Example

``` go
```

#### (*IntIntMap) Size

```go
func (m *IntIntMap) Size() int
```

Size returns the size of the map.

​	Size 返回地图的大小。

##### Example

``` go
```

#### (*IntIntMap) String

```go
func (m *IntIntMap) String() string
```

String returns the map as a string.

​	String 以字符串形式返回映射。

##### Example

``` go
```

#### (*IntIntMap) UnmarshalJSON

```go
func (m *IntIntMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*IntIntMap) UnmarshalValue

```go
func (m *IntIntMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

​	UnmarshalValue 是一个接口实现，用于为 map 设置任何类型的值。

##### Example

``` go
```

#### (*IntIntMap) Values

```go
func (m *IntIntMap) Values() []int
```

Values returns all values of the map as a slice.

​	Values 以切片的形式返回地图的所有值。

##### Example

``` go
```

### type IntStrMap

```go
type IntStrMap struct {
	// contains filtered or unexported fields
}
```

IntStrMap implements map[int]string with RWMutex that has switch.

​	IntStrMap 使用具有 switch 的 RWMutex 实现 map[int]字符串。

#### func NewIntStrMap

```go
func NewIntStrMap(safe ...bool) *IntStrMap
```

NewIntStrMap returns an empty IntStrMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

​	NewIntStrMap 返回一个空的 IntStrMap 对象。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 map，默认为 false。

#### func NewIntStrMapFrom

```go
func NewIntStrMapFrom(data map[int]string, safe ...bool) *IntStrMap
```

NewIntStrMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

​	NewIntStrMapFrom 从给定映射 `data` 创建并返回哈希映射。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。

#### (*IntStrMap) Clear

```go
func (m *IntStrMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

​	清除会删除地图的所有数据，它将重新制作一个新的基础数据地图。

#### (*IntStrMap) Clone

```go
func (m *IntStrMap) Clone() *IntStrMap
```

Clone returns a new hash map with copy of current map data.

​	克隆将返回一个新的哈希映射，其中包含当前映射数据的副本。

#### (*IntStrMap) Contains

```go
func (m *IntStrMap) Contains(key int) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

​	包含检查密钥是否存在。如果存在， `key` 则返回 true，否则返回 false。

#### (*IntStrMap) DeepCopy

```go
func (m *IntStrMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*IntStrMap) Diff

```go
func (m *IntStrMap) Diff(other *IntStrMap) (addedKeys, removedKeys, updatedKeys []int)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

​	Diff 将当前映射 `m` 与映射 `other` 进行比较，并返回它们的不同键。返回 `addedKeys` 的键在 map `m` 中但不在 map `other` 中。返回 `removedKeys` 的键在 map `other` 中但不在 map `m` 中。返回 `updatedKeys` 的键既在 map `m` 中， `other` 又在 但它们的值中且不相等 （ `!=` ）。

#### (*IntStrMap) FilterEmpty

```go
func (m *IntStrMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, “”, len(slice/map/chan) == 0 are considered empty.

​	FilterEmpty 删除所有值为空的键值对。像 0， nil， false， “”， len（slice/map/chan） == 0 这样的值被视为空。

#### (*IntStrMap) Flip

```go
func (m *IntStrMap) Flip()
```

Flip exchanges key-value of the map to value-key.

​	Flip 将映射的键值交换为值键。

#### (*IntStrMap) Get

```go
func (m *IntStrMap) Get(key int) (value string)
```

Get returns the value by given `key`.

​	Get 返回给定 `key` 的值。

#### (*IntStrMap) GetOrSet

```go
func (m *IntStrMap) GetOrSet(key int, value string) string
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

#### (*IntStrMap) GetOrSetFunc

```go
func (m *IntStrMap) GetOrSetFunc(key int, f func() string) string
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在）并返回此值。

#### (*IntStrMap) GetOrSetFuncLock

```go
func (m *IntStrMap) GetOrSetFuncLock(key int, f func() string) string
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在）并返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

#### (*IntStrMap) IsEmpty

```go
func (m *IntStrMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

​	IsEmpty 检查地图是否为空。如果 map 为空，则返回 true，否则返回 false。

#### (*IntStrMap) IsSubOf

```go
func (m *IntStrMap) IsSubOf(other *IntStrMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

​	IsSubOf 检查当前映射是否是 的 `other` 子映射。

#### (*IntStrMap) Iterator

```go
func (m *IntStrMap) Iterator(f func(k int, v string) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	迭代器使用自定义回调函数 `f` 迭代哈希映射只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*IntStrMap) Keys

```go
func (m *IntStrMap) Keys() []int
```

Keys returns all keys of the map as a slice.

​	Keys 将地图的所有键作为切片返回。

#### (*IntStrMap) LockFunc

```go
func (m *IntStrMap) LockFunc(f func(m map[int]string))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

​	LockFunc 在 RWMutex.Lock 中使用给定的回调函数 `f` 锁定写入。

#### (*IntStrMap) Map

```go
func (m *IntStrMap) Map() map[int]string
```

Map returns the underlying data map. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Map 返回基础数据映射。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

#### (*IntStrMap) MapCopy

```go
func (m *IntStrMap) MapCopy() map[int]string
```

MapCopy returns a copy of the underlying data of the hash map.

​	MapCopy 返回哈希映射的基础数据的副本。

#### (*IntStrMap) MapStrAny

```go
func (m *IntStrMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回地图基础数据的副本。

#### (IntStrMap) MarshalJSON

```go
func (m IntStrMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*IntStrMap) Merge

```go
func (m *IntStrMap) Merge(other *IntStrMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

​	合并合并两个哈希映射。 `other` 地图将合并到地图 `m` 中。

#### (*IntStrMap) Pop

```go
func (m *IntStrMap) Pop() (key int, value string)
```

Pop retrieves and deletes an item from the map.

​	Pop 从地图中检索和删除项目。

#### (*IntStrMap) Pops

```go
func (m *IntStrMap) Pops(size int) map[int]string
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

​	Pops 从地图中检索和删除 `size` 项目。如果大小 == -1，则返回所有项目。

#### (*IntStrMap) RLockFunc

```go
func (m *IntStrMap) RLockFunc(f func(m map[int]string))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

​	RLockFunc 在 RWMutex.RLock 中使用给定的回调函数 `f` 锁定读取。

#### (*IntStrMap) Remove

```go
func (m *IntStrMap) Remove(key int) (value string)
```

Remove deletes value from map by given `key`, and return this deleted value.

​	按给定 `key` 从映射中删除删除值，并返回此删除值。

#### (*IntStrMap) Removes

```go
func (m *IntStrMap) Removes(keys []int)
```

Removes batch deletes values of the map by keys.

​	删除按键批量删除映射的值。

#### (*IntStrMap) Replace

```go
func (m *IntStrMap) Replace(data map[int]string)
```

Replace the data of the map with given `data`.

​	将地图的数据替换为给定 `data` 的 .

#### (*IntStrMap) Search

```go
func (m *IntStrMap) Search(key int) (value string, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的 .如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

#### (*IntStrMap) Set

```go
func (m *IntStrMap) Set(key int, val string)
```

Set sets key-value to the hash map.

​	Set 将 key-value 设置为哈希映射。

#### (*IntStrMap) SetIfNotExist

```go
func (m *IntStrMap) SetIfNotExist(key int, value string) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

#### (*IntStrMap) SetIfNotExistFunc

```go
func (m *IntStrMap) SetIfNotExistFunc(key int, f func() string) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

#### (*IntStrMap) SetIfNotExistFuncLock

```go
func (m *IntStrMap) SetIfNotExistFuncLock(key int, f func() string) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

#### (*IntStrMap) Sets

```go
func (m *IntStrMap) Sets(data map[int]string)
```

Sets batch sets key-values to the hash map.

​	将批处理设置键值设置为哈希映射。

#### (*IntStrMap) Size

```go
func (m *IntStrMap) Size() int
```

Size returns the size of the map.

​	Size 返回地图的大小。

#### (*IntStrMap) String

```go
func (m *IntStrMap) String() string
```

String returns the map as a string.

​	String 以字符串形式返回映射。

#### (*IntStrMap) UnmarshalJSON

```go
func (m *IntStrMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*IntStrMap) UnmarshalValue

```go
func (m *IntStrMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

​	UnmarshalValue 是一个接口实现，用于为 map 设置任何类型的值。

#### (*IntStrMap) Values

```go
func (m *IntStrMap) Values() []string
```

Values returns all values of the map as a slice.

​	Values 以切片的形式返回地图的所有值。

### type ListMap

```go
type ListMap struct {
	// contains filtered or unexported fields
}
```

ListMap is a map that preserves insertion-order.

​	ListMap 是保留插入顺序的映射。

It is backed by a hash table to store values and doubly-linked list to store ordering.

​	它由用于存储值的哈希表和用于存储排序的双向链表提供支持。

Structure is not thread safe.

​	结构不是线程安全的。

Reference: http://en.wikipedia.org/wiki/Associative_array

​	参考资料： http://en.wikipedia.org/wiki/Associative_array

#### func NewListMap

```go
func NewListMap(safe ...bool) *ListMap
```

NewListMap returns an empty link map. ListMap is backed by a hash table to store values and doubly-linked list to store ordering. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

​	NewListMap 返回一个空的链接映射。ListMap 由用于存储值的哈希表和用于存储排序的双链列表提供支持。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 map，默认为 false。

##### Example

``` go
```

#### func NewListMapFrom

```go
func NewListMapFrom(data map[interface{}]interface{}, safe ...bool) *ListMap
```

NewListMapFrom returns a link map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

​	NewListMapFrom 从给定的 map `data` 返回链接映射。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。

##### Example

``` go
```

#### (*ListMap) Clear

```go
func (m *ListMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

​	清除会删除地图的所有数据，它将重新制作一个新的基础数据地图。

##### Example

``` go
```

#### (*ListMap) Clone

```go
func (m *ListMap) Clone(safe ...bool) *ListMap
```

Clone returns a new link map with copy of current map data.

​	克隆将返回包含当前地图数据副本的新链接地图。

##### Example

``` go
```

#### (*ListMap) Contains

```go
func (m *ListMap) Contains(key interface{}) (ok bool)
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

​	包含检查密钥是否存在。如果存在， `key` 则返回 true，否则返回 false。

##### Example

``` go
```

#### (*ListMap) DeepCopy

```go
func (m *ListMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*ListMap) FilterEmpty

```go
func (m *ListMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty.

​	FilterEmpty 删除所有值为空的键值对。

##### Example

``` go
```

#### (*ListMap) Flip

```go
func (m *ListMap) Flip()
```

Flip exchanges key-value of the map to value-key.

​	Flip 将映射的键值交换为值键。

##### Example

``` go
```

#### (*ListMap) Get

```go
func (m *ListMap) Get(key interface{}) (value interface{})
```

Get returns the value by given `key`.

​	Get 返回给定 `key` 的值。

##### Example

``` go
```

#### (*ListMap) GetOrSet

```go
func (m *ListMap) GetOrSet(key interface{}, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

##### Example

``` go
```

#### (*ListMap) GetOrSetFunc

```go
func (m *ListMap) GetOrSetFunc(key interface{}, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

##### Example

``` go
```

#### (*ListMap) GetOrSetFuncLock

```go
func (m *ListMap) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。地图的锁定。

##### Example

``` go
```

#### (*ListMap) GetVar

```go
func (m *ListMap) GetVar(key interface{}) *gvar.Var
```

GetVar returns a Var with the value by given `key`. The returned Var is un-concurrent safe.

​	GetVar 返回一个 Var，其值为 给定 `key` 。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*ListMap) GetVarOrSet

```go
func (m *ListMap) GetVarOrSet(key interface{}, value interface{}) *gvar.Var
```

GetVarOrSet returns a Var with result from GetVarOrSet. The returned Var is un-concurrent safe.

​	GetVarOrSet 返回一个 Var，其中包含来自 GetVarOrSet 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*ListMap) GetVarOrSetFunc

```go
func (m *ListMap) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a Var with result from GetOrSetFunc. The returned Var is un-concurrent safe.

​	GetVarOrSetFunc 返回一个 Var，其中包含来自 GetOrSetFunc 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*ListMap) GetVarOrSetFuncLock

```go
func (m *ListMap) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock. The returned Var is un-concurrent safe.

​	GetVarOrSetFuncLock 返回一个 Var，其中包含来自 GetOrSetFuncLock 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*ListMap) IsEmpty

```go
func (m *ListMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

​	IsEmpty 检查地图是否为空。如果 map 为空，则返回 true，否则返回 false。

##### Example

``` go
```

#### (*ListMap) Iterator

```go
func (m *ListMap) Iterator(f func(key, value interface{}) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

##### Example

``` go
```

#### (*ListMap) IteratorAsc

```go
func (m *ListMap) IteratorAsc(f func(key interface{}, value interface{}) bool)
```

IteratorAsc iterates the map readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序迭代映射只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*ListMap) IteratorDesc

```go
func (m *ListMap) IteratorDesc(f func(key interface{}, value interface{}) bool)
```

IteratorDesc iterates the map readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序迭代映射只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*ListMap) Keys

```go
func (m *ListMap) Keys() []interface{}
```

Keys returns all keys of the map as a slice in ascending order.

​	Keys 按升序将地图的所有键作为切片返回。

##### Example

``` go
```

#### (*ListMap) Map

```go
func (m *ListMap) Map() map[interface{}]interface{}
```

Map returns a copy of the underlying data of the map.

​	Map 返回地图基础数据的副本。

##### Example

``` go
```

#### (*ListMap) MapStrAny

```go
func (m *ListMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回地图基础数据的副本。

##### Example

``` go
```

#### (ListMap) MarshalJSON

```go
func (m ListMap) MarshalJSON() (jsonBytes []byte, err error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*ListMap) Merge

```go
func (m *ListMap) Merge(other *ListMap)
```

Merge merges two link maps. The `other` map will be merged into the map `m`.

​	合并合并两个链接映射。 `other` 地图将合并到地图 `m` 中。

##### Example

``` go
```

#### (*ListMap) Pop

```go
func (m *ListMap) Pop() (key, value interface{})
```

Pop retrieves and deletes an item from the map.

​	Pop 从地图中检索和删除项目。

##### Example

``` go
```

#### (*ListMap) Pops

```go
func (m *ListMap) Pops(size int) map[interface{}]interface{}
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

​	Pops 从地图中检索和删除 `size` 项目。如果大小 == -1，则返回所有项目。

##### Example

``` go
```

#### (*ListMap) Remove

```go
func (m *ListMap) Remove(key interface{}) (value interface{})
```

Remove deletes value from map by given `key`, and return this deleted value.

​	按给定 `key` 从映射中删除删除值，并返回此删除值。

##### Example

``` go
```

#### (*ListMap) Removes

```go
func (m *ListMap) Removes(keys []interface{})
```

Removes batch deletes values of the map by keys.

​	删除按键批量删除映射的值。

##### Example

``` go
```

#### (*ListMap) Replace

```go
func (m *ListMap) Replace(data map[interface{}]interface{})
```

Replace the data of the map with given `data`.

​	将地图的数据替换为给定 `data` 的 .

##### Example

``` go
```

#### (*ListMap) Search

```go
func (m *ListMap) Search(key interface{}) (value interface{}, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的 .如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

##### Example

``` go
```

#### (*ListMap) Set

```go
func (m *ListMap) Set(key interface{}, value interface{})
```

Set sets key-value to the map.

​	Set 将键值设置为映射。

##### Example

``` go
```

#### (*ListMap) SetIfNotExist

```go
func (m *ListMap) SetIfNotExist(key interface{}, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*ListMap) SetIfNotExistFunc

```go
func (m *ListMap) SetIfNotExistFunc(key interface{}, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*ListMap) SetIfNotExistFuncLock

```go
func (m *ListMap) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。地图的锁定。

##### Example

``` go
```

#### (*ListMap) Sets

```go
func (m *ListMap) Sets(data map[interface{}]interface{})
```

Sets batch sets key-values to the map.

​	将批处理设置键值设置为映射。

##### Example

``` go
```

#### (*ListMap) Size

```go
func (m *ListMap) Size() (size int)
```

Size returns the size of the map.

​	Size 返回地图的大小。

##### Example

``` go
```

#### (*ListMap) String

```go
func (m *ListMap) String() string
```

String returns the map as a string.

​	String 以字符串形式返回映射。

##### Example

``` go
```

#### (*ListMap) UnmarshalJSON

```go
func (m *ListMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*ListMap) UnmarshalValue

```go
func (m *ListMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

​	UnmarshalValue 是一个接口实现，用于为 map 设置任何类型的值。

##### Example

``` go
```

#### (*ListMap) Values

```go
func (m *ListMap) Values() []interface{}
```

Values returns all values of the map as a slice.

​	Values 以切片的形式返回地图的所有值。

##### Example

``` go
```

### type Map

```go
type Map = AnyAnyMap // Map is alias of AnyAnyMap.
```

#### func New

```go
func New(safe ...bool) *Map
```

New creates and returns an empty hash map. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

​	New 创建并返回一个空的哈希映射。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 map，默认为 false。

##### Example

``` go
```

#### func NewFrom

```go
func NewFrom(data map[interface{}]interface{}, safe ...bool) *Map
```

NewFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

​	NewFrom 从给定的映射 `data` 创建并返回哈希映射。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 tree，默认为 false。

##### Example

``` go
```

#### func NewHashMap

```go
func NewHashMap(safe ...bool) *Map
```

NewHashMap creates and returns an empty hash map. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

​	NewHashMap 创建并返回一个空的哈希映射。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 map，默认为 false。

##### Example

``` go
```

#### func NewHashMapFrom

```go
func NewHashMapFrom(data map[interface{}]interface{}, safe ...bool) *Map
```

NewHashMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

​	NewHashMapFrom 从给定的映射 `data` 创建并返回哈希映射。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 tree，默认为 false。

##### Example

``` go
```

### type StrAnyMap

```go
type StrAnyMap struct {
	// contains filtered or unexported fields
}
```

StrAnyMap implements map[string]interface{} with RWMutex that has switch.

​	StrAnyMap 使用具有 switch 的 RWMutex 实现 map[string]interface{}。

#### func NewStrAnyMap

```go
func NewStrAnyMap(safe ...bool) *StrAnyMap
```

NewStrAnyMap returns an empty StrAnyMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

​	NewStrAnyMap 返回一个空的 StrAnyMap 对象。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 map，默认为 false。

##### Example

``` go
```

#### func NewStrAnyMapFrom

```go
func NewStrAnyMapFrom(data map[string]interface{}, safe ...bool) *StrAnyMap
```

NewStrAnyMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

​	NewStrAnyMapFrom 从给定的映射 `data` 创建并返回哈希映射。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。

##### Example

``` go
```

#### (*StrAnyMap) Clear

```go
func (m *StrAnyMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

​	清除会删除地图的所有数据，它将重新制作一个新的基础数据地图。

##### Example

``` go
```

#### (*StrAnyMap) Clone

```go
func (m *StrAnyMap) Clone() *StrAnyMap
```

Clone returns a new hash map with copy of current map data.

​	克隆将返回一个新的哈希映射，其中包含当前映射数据的副本。

##### Example

``` go
```

#### (*StrAnyMap) Contains

```go
func (m *StrAnyMap) Contains(key string) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

​	包含检查密钥是否存在。如果存在， `key` 则返回 true，否则返回 false。

##### Example

``` go
```

#### (*StrAnyMap) DeepCopy

```go
func (m *StrAnyMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*StrAnyMap) Diff

```go
func (m *StrAnyMap) Diff(other *StrAnyMap) (addedKeys, removedKeys, updatedKeys []string)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

​	Diff 将当前映射 `m` 与映射 `other` 进行比较，并返回它们的不同键。返回 `addedKeys` 的键在 map `m` 中但不在 map `other` 中。返回 `removedKeys` 的键在 map `other` 中但不在 map `m` 中。返回 `updatedKeys` 的键既在 map `m` 中， `other` 又在 但它们的值中且不相等 （ `!=` ）。

#### (*StrAnyMap) FilterEmpty

```go
func (m *StrAnyMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, “”, len(slice/map/chan) == 0 are considered empty.

​	FilterEmpty 删除所有值为空的键值对。像 0， nil， false， “”， len（slice/map/chan） == 0 这样的值被视为空。

##### Example

``` go
```

#### (*StrAnyMap) FilterNil

```go
func (m *StrAnyMap) FilterNil()
```

FilterNil deletes all key-value pair of which the value is nil.

​	FilterNil 删除值为 nil 的所有键值对。

##### Example

``` go
```

#### (*StrAnyMap) Flip

```go
func (m *StrAnyMap) Flip()
```

Flip exchanges key-value of the map to value-key.

​	Flip 将映射的键值交换为值键。

##### Example

``` go
```

#### (*StrAnyMap) Get

```go
func (m *StrAnyMap) Get(key string) (value interface{})
```

Get returns the value by given `key`.

​	Get 返回给定 `key` 的值。

##### Example

``` go
```

#### (*StrAnyMap) GetOrSet

```go
func (m *StrAnyMap) GetOrSet(key string, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

##### Example

``` go
```

#### (*StrAnyMap) GetOrSetFunc

```go
func (m *StrAnyMap) GetOrSetFunc(key string, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

##### Example

``` go
```

#### (*StrAnyMap) GetOrSetFuncLock

```go
func (m *StrAnyMap) GetOrSetFuncLock(key string, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*StrAnyMap) GetVar

```go
func (m *StrAnyMap) GetVar(key string) *gvar.Var
```

GetVar returns a Var with the value by given `key`. The returned Var is un-concurrent safe.

​	GetVar 返回一个 Var，其值为 给定 `key` 。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*StrAnyMap) GetVarOrSet

```go
func (m *StrAnyMap) GetVarOrSet(key string, value interface{}) *gvar.Var
```

GetVarOrSet returns a Var with result from GetVarOrSet. The returned Var is un-concurrent safe.

​	GetVarOrSet 返回一个 Var，其中包含来自 GetVarOrSet 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*StrAnyMap) GetVarOrSetFunc

```go
func (m *StrAnyMap) GetVarOrSetFunc(key string, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a Var with result from GetOrSetFunc. The returned Var is un-concurrent safe.

​	GetVarOrSetFunc 返回一个 Var，其中包含来自 GetOrSetFunc 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*StrAnyMap) GetVarOrSetFuncLock

```go
func (m *StrAnyMap) GetVarOrSetFuncLock(key string, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock. The returned Var is un-concurrent safe.

​	GetVarOrSetFuncLock 返回一个 Var，其中包含来自 GetOrSetFuncLock 的结果。返回的 Var 是非并发安全的。

##### Example

``` go
```

#### (*StrAnyMap) IsEmpty

```go
func (m *StrAnyMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

​	IsEmpty 检查地图是否为空。如果 map 为空，则返回 true，否则返回 false。

##### Example

``` go
```

#### (*StrAnyMap) IsSubOf

```go
func (m *StrAnyMap) IsSubOf(other *StrAnyMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

​	IsSubOf 检查当前映射是否是 的 `other` 子映射。

#### (*StrAnyMap) Iterator

```go
func (m *StrAnyMap) Iterator(f func(k string, v interface{}) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	迭代器使用自定义回调函数 `f` 迭代哈希映射只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*StrAnyMap) Keys

```go
func (m *StrAnyMap) Keys() []string
```

Keys returns all keys of the map as a slice.

​	Keys 将地图的所有键作为切片返回。

##### Example

``` go
```

#### (*StrAnyMap) LockFunc

```go
func (m *StrAnyMap) LockFunc(f func(m map[string]interface{}))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

​	LockFunc 在 RWMutex.Lock 中使用给定的回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (*StrAnyMap) Map

```go
func (m *StrAnyMap) Map() map[string]interface{}
```

Map returns the underlying data map. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Map 返回基础数据映射。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

##### Example

``` go
```

#### (*StrAnyMap) MapCopy

```go
func (m *StrAnyMap) MapCopy() map[string]interface{}
```

MapCopy returns a copy of the underlying data of the hash map.

​	MapCopy 返回哈希映射的基础数据的副本。

##### Example

``` go
```

#### (*StrAnyMap) MapStrAny

```go
func (m *StrAnyMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回地图基础数据的副本。

##### Example

``` go
```

#### (StrAnyMap) MarshalJSON

```go
func (m StrAnyMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*StrAnyMap) Merge

```go
func (m *StrAnyMap) Merge(other *StrAnyMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

​	合并合并两个哈希映射。 `other` 地图将合并到地图 `m` 中。

##### Example

``` go
```

#### (*StrAnyMap) Pop

```go
func (m *StrAnyMap) Pop() (key string, value interface{})
```

Pop retrieves and deletes an item from the map.

​	Pop 从地图中检索和删除项目。

##### Example

``` go
```

#### (*StrAnyMap) Pops

```go
func (m *StrAnyMap) Pops(size int) map[string]interface{}
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

​	Pops 从地图中检索和删除 `size` 项目。如果大小 == -1，则返回所有项目。

##### Example

``` go
```

#### (*StrAnyMap) RLockFunc

```go
func (m *StrAnyMap) RLockFunc(f func(m map[string]interface{}))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

​	RLockFunc 在 RWMutex.RLock 中使用给定的回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*StrAnyMap) Remove

```go
func (m *StrAnyMap) Remove(key string) (value interface{})
```

Remove deletes value from map by given `key`, and return this deleted value.

​	按给定 `key` 从映射中删除删除值，并返回此删除值。

##### Example

``` go
```

#### (*StrAnyMap) Removes

```go
func (m *StrAnyMap) Removes(keys []string)
```

Removes batch deletes values of the map by keys.

​	删除按键批量删除映射的值。

##### Example

``` go
```

#### (*StrAnyMap) Replace

```go
func (m *StrAnyMap) Replace(data map[string]interface{})
```

Replace the data of the map with given `data`.

​	将地图的数据替换为给定 `data` 的 .

##### Example

``` go
```

#### (*StrAnyMap) Search

```go
func (m *StrAnyMap) Search(key string) (value interface{}, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的 .如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

##### Example

``` go
```

#### (*StrAnyMap) Set

```go
func (m *StrAnyMap) Set(key string, val interface{})
```

Set sets key-value to the hash map.

​	Set 将 key-value 设置为哈希映射。

##### Example

``` go
```

#### (*StrAnyMap) SetIfNotExist

```go
func (m *StrAnyMap) SetIfNotExist(key string, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*StrAnyMap) SetIfNotExistFunc

```go
func (m *StrAnyMap) SetIfNotExistFunc(key string, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*StrAnyMap) SetIfNotExistFuncLock

```go
func (m *StrAnyMap) SetIfNotExistFuncLock(key string, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*StrAnyMap) Sets

```go
func (m *StrAnyMap) Sets(data map[string]interface{})
```

Sets batch sets key-values to the hash map.

​	将批处理设置键值设置为哈希映射。

##### Example

``` go
```

#### (*StrAnyMap) Size

```go
func (m *StrAnyMap) Size() int
```

Size returns the size of the map.

​	Size 返回地图的大小。

##### Example

``` go
```

#### (*StrAnyMap) String

```go
func (m *StrAnyMap) String() string
```

String returns the map as a string.

​	String 以字符串形式返回映射。

##### Example

``` go
```

#### (*StrAnyMap) UnmarshalJSON

```go
func (m *StrAnyMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*StrAnyMap) UnmarshalValue

```go
func (m *StrAnyMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

​	UnmarshalValue 是一个接口实现，用于为 map 设置任何类型的值。

##### Example

``` go
```

#### (*StrAnyMap) Values

```go
func (m *StrAnyMap) Values() []interface{}
```

Values returns all values of the map as a slice.

​	Values 以切片的形式返回地图的所有值。

##### Example

``` go
```

### type StrIntMap

```go
type StrIntMap struct {
	// contains filtered or unexported fields
}
```

StrIntMap implements map[string]int with RWMutex that has switch.

​	StrIntMap 使用具有 switch 的 RWMutex 实现 map[string]int。

#### func NewStrIntMap

```go
func NewStrIntMap(safe ...bool) *StrIntMap
```

NewStrIntMap returns an empty StrIntMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

​	NewStrIntMap 返回一个空的 StrIntMap 对象。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 map，默认为 false。

##### Example

``` go
```

#### func NewStrIntMapFrom

```go
func NewStrIntMapFrom(data map[string]int, safe ...bool) *StrIntMap
```

NewStrIntMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

​	NewStrIntMapFrom 从给定的映射 `data` 创建并返回哈希映射。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。

##### Example

``` go
```

#### (*StrIntMap) Clear

```go
func (m *StrIntMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

​	清除会删除地图的所有数据，它将重新制作一个新的基础数据地图。

##### Example

``` go
```

#### (*StrIntMap) Clone

```go
func (m *StrIntMap) Clone() *StrIntMap
```

Clone returns a new hash map with copy of current map data.

​	克隆将返回一个新的哈希映射，其中包含当前映射数据的副本。

##### Example

``` go
```

#### (*StrIntMap) Contains

```go
func (m *StrIntMap) Contains(key string) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

​	包含检查密钥是否存在。如果存在， `key` 则返回 true，否则返回 false。

##### Example

``` go
```

#### (*StrIntMap) DeepCopy

```go
func (m *StrIntMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*StrIntMap) Diff

```go
func (m *StrIntMap) Diff(other *StrIntMap) (addedKeys, removedKeys, updatedKeys []string)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

​	Diff 将当前映射 `m` 与映射 `other` 进行比较，并返回它们的不同键。返回 `addedKeys` 的键在 map `m` 中但不在 map `other` 中。返回 `removedKeys` 的键在 map `other` 中但不在 map `m` 中。返回 `updatedKeys` 的键既在 map `m` 中， `other` 又在 但它们的值中且不相等 （ `!=` ）。

#### (*StrIntMap) FilterEmpty

```go
func (m *StrIntMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, “”, len(slice/map/chan) == 0 are considered empty.

​	FilterEmpty 删除所有值为空的键值对。像 0， nil， false， “”， len（slice/map/chan） == 0 这样的值被视为空。

##### Example

``` go
```

#### (*StrIntMap) Flip

```go
func (m *StrIntMap) Flip()
```

Flip exchanges key-value of the map to value-key.

​	Flip 将映射的键值交换为值键。

##### Example

``` go
```

#### (*StrIntMap) Get

```go
func (m *StrIntMap) Get(key string) (value int)
```

Get returns the value by given `key`.

​	Get 返回给定 `key` 的值。

##### Example

``` go
```

#### (*StrIntMap) GetOrSet

```go
func (m *StrIntMap) GetOrSet(key string, value int) int
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

##### Example

``` go
```

#### (*StrIntMap) GetOrSetFunc

```go
func (m *StrIntMap) GetOrSetFunc(key string, f func() int) int
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

##### Example

``` go
```

#### (*StrIntMap) GetOrSetFuncLock

```go
func (m *StrIntMap) GetOrSetFuncLock(key string, f func() int) int
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*StrIntMap) IsEmpty

```go
func (m *StrIntMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

​	IsEmpty 检查地图是否为空。如果 map 为空，则返回 true，否则返回 false。

##### Example

``` go
```

#### (*StrIntMap) IsSubOf

```go
func (m *StrIntMap) IsSubOf(other *StrIntMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

​	IsSubOf 检查当前映射是否是 的 `other` 子映射。

#### (*StrIntMap) Iterator

```go
func (m *StrIntMap) Iterator(f func(k string, v int) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	迭代器使用自定义回调函数 `f` 迭代哈希映射只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*StrIntMap) Keys

```go
func (m *StrIntMap) Keys() []string
```

Keys returns all keys of the map as a slice.

​	Keys 将地图的所有键作为切片返回。

##### Example

``` go
```

#### (*StrIntMap) LockFunc

```go
func (m *StrIntMap) LockFunc(f func(m map[string]int))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

​	LockFunc 在 RWMutex.Lock 中使用给定的回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (*StrIntMap) Map

```go
func (m *StrIntMap) Map() map[string]int
```

Map returns the underlying data map. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Map 返回基础数据映射。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

##### Example

``` go
```

#### (*StrIntMap) MapCopy

```go
func (m *StrIntMap) MapCopy() map[string]int
```

MapCopy returns a copy of the underlying data of the hash map.

​	MapCopy 返回哈希映射的基础数据的副本。

##### Example

``` go
```

#### (*StrIntMap) MapStrAny

```go
func (m *StrIntMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回地图基础数据的副本。

##### Example

``` go
```

#### (StrIntMap) MarshalJSON

```go
func (m StrIntMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*StrIntMap) Merge

```go
func (m *StrIntMap) Merge(other *StrIntMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

​	合并合并两个哈希映射。 `other` 地图将合并到地图 `m` 中。

##### Example

``` go
```

#### (*StrIntMap) Pop

```go
func (m *StrIntMap) Pop() (key string, value int)
```

Pop retrieves and deletes an item from the map.

​	Pop 从地图中检索和删除项目。

##### Example

``` go
```

#### (*StrIntMap) Pops

```go
func (m *StrIntMap) Pops(size int) map[string]int
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

​	Pops 从地图中检索和删除 `size` 项目。如果大小 == -1，则返回所有项目。

##### Example

``` go
```

#### (*StrIntMap) RLockFunc

```go
func (m *StrIntMap) RLockFunc(f func(m map[string]int))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

​	RLockFunc 在 RWMutex.RLock 中使用给定的回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*StrIntMap) Remove

```go
func (m *StrIntMap) Remove(key string) (value int)
```

Remove deletes value from map by given `key`, and return this deleted value.

​	按给定 `key` 从映射中删除删除值，并返回此删除值。

##### Example

``` go
```

#### (*StrIntMap) Removes

```go
func (m *StrIntMap) Removes(keys []string)
```

Removes batch deletes values of the map by keys.

​	删除按键批量删除映射的值。

##### Example

``` go
```

#### (*StrIntMap) Replace

```go
func (m *StrIntMap) Replace(data map[string]int)
```

Replace the data of the map with given `data`.

​	将地图的数据替换为给定 `data` 的 .

##### Example

``` go
```

#### (*StrIntMap) Search

```go
func (m *StrIntMap) Search(key string) (value int, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的 .如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

##### Example

``` go
```

#### (*StrIntMap) Set

```go
func (m *StrIntMap) Set(key string, val int)
```

Set sets key-value to the hash map.

​	Set 将 key-value 设置为哈希映射。

##### Example

``` go
```

#### (*StrIntMap) SetIfNotExist

```go
func (m *StrIntMap) SetIfNotExist(key string, value int) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*StrIntMap) SetIfNotExistFunc

```go
func (m *StrIntMap) SetIfNotExistFunc(key string, f func() int) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*StrIntMap) SetIfNotExistFuncLock

```go
func (m *StrIntMap) SetIfNotExistFuncLock(key string, f func() int) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*StrIntMap) Sets

```go
func (m *StrIntMap) Sets(data map[string]int)
```

Sets batch sets key-values to the hash map.

​	将批处理设置键值设置为哈希映射。

##### Example

``` go
```

#### (*StrIntMap) Size

```go
func (m *StrIntMap) Size() int
```

Size returns the size of the map.

​	Size 返回地图的大小。

##### Example

``` go
```

#### (*StrIntMap) String

```go
func (m *StrIntMap) String() string
```

String returns the map as a string.

​	String 以字符串形式返回映射。

##### Example

``` go
```

#### (*StrIntMap) UnmarshalJSON

```go
func (m *StrIntMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*StrIntMap) UnmarshalValue

```go
func (m *StrIntMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

​	UnmarshalValue 是一个接口实现，用于为 map 设置任何类型的值。

##### Example

``` go
```

#### (*StrIntMap) Values

```go
func (m *StrIntMap) Values() []int
```

Values returns all values of the map as a slice.

​	Values 以切片的形式返回地图的所有值。

##### Example

``` go
```

### type StrStrMap

```go
type StrStrMap struct {
	// contains filtered or unexported fields
}
```

StrStrMap implements map[string]string with RWMutex that has switch.

​	StrStrMap 使用具有 switch 的 RWMutex 实现 map[string]string。

#### func NewStrStrMap

```go
func NewStrStrMap(safe ...bool) *StrStrMap
```

NewStrStrMap returns an empty StrStrMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

​	NewStrStrMap 返回一个空的 StrStrMap 对象。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 map，默认为 false。

##### Example

``` go
```

#### func NewStrStrMapFrom

```go
func NewStrStrMapFrom(data map[string]string, safe ...bool) *StrStrMap
```

NewStrStrMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

​	NewStrStrMapFrom 从给定的映射 `data` 创建并返回哈希映射。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。

##### Example

``` go
```

#### (*StrStrMap) Clear

```go
func (m *StrStrMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

​	清除会删除地图的所有数据，它将重新制作一个新的基础数据地图。

##### Example

``` go
```

#### (*StrStrMap) Clone

```go
func (m *StrStrMap) Clone() *StrStrMap
```

Clone returns a new hash map with copy of current map data.

​	克隆将返回一个新的哈希映射，其中包含当前映射数据的副本。

##### Example

``` go
```

#### (*StrStrMap) Contains

```go
func (m *StrStrMap) Contains(key string) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

​	包含检查密钥是否存在。如果存在， `key` 则返回 true，否则返回 false。

##### Example

``` go
```

#### (*StrStrMap) DeepCopy

```go
func (m *StrStrMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*StrStrMap) Diff

```go
func (m *StrStrMap) Diff(other *StrStrMap) (addedKeys, removedKeys, updatedKeys []string)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

​	Diff 将当前映射 `m` 与映射 `other` 进行比较，并返回它们的不同键。返回 `addedKeys` 的键在 map `m` 中但不在 map `other` 中。返回 `removedKeys` 的键在 map `other` 中但不在 map `m` 中。返回 `updatedKeys` 的键既在 map `m` 中， `other` 又在 但它们的值中且不相等 （ `!=` ）。

#### (*StrStrMap) FilterEmpty

```go
func (m *StrStrMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, “”, len(slice/map/chan) == 0 are considered empty.

​	FilterEmpty 删除所有值为空的键值对。像 0， nil， false， “”， len（slice/map/chan） == 0 这样的值被视为空。

##### Example

``` go
```

#### (*StrStrMap) Flip

```go
func (m *StrStrMap) Flip()
```

Flip exchanges key-value of the map to value-key.

​	Flip 将映射的键值交换为值键。

##### Example

``` go
```

#### (*StrStrMap) Get

```go
func (m *StrStrMap) Get(key string) (value string)
```

Get returns the value by given `key`.

​	Get 返回给定 `key` 的值。

##### Example

``` go
```

#### (*StrStrMap) GetOrSet

```go
func (m *StrStrMap) GetOrSet(key string, value string) string
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

​	GetOrSet 按键返回值，如果值不存在，则使用 given `value` 设置值，然后返回此值。

##### Example

``` go
```

#### (*StrStrMap) GetOrSetFunc

```go
func (m *StrStrMap) GetOrSetFunc(key string, f func() string) string
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFunc 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

##### Example

``` go
```

#### (*StrStrMap) GetOrSetFuncLock

```go
func (m *StrStrMap) GetOrSetFuncLock(key string, f func() string) string
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

​	GetOrSetFuncLock 按键返回值，或者使用回调函数 `f` 的返回值设置值（如果不存在），然后返回此值。

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*StrStrMap) IsEmpty

```go
func (m *StrStrMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

​	IsEmpty 检查地图是否为空。如果 map 为空，则返回 true，否则返回 false。

##### Example

``` go
```

#### (*StrStrMap) IsSubOf

```go
func (m *StrStrMap) IsSubOf(other *StrStrMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

​	IsSubOf 检查当前映射是否是 的 `other` 子映射。

#### (*StrStrMap) Iterator

```go
func (m *StrStrMap) Iterator(f func(k string, v string) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	迭代器使用自定义回调函数 `f` 迭代哈希映射只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*StrStrMap) Keys

```go
func (m *StrStrMap) Keys() []string
```

Keys returns all keys of the map as a slice.

​	Keys 将地图的所有键作为切片返回。

##### Example

``` go
```

#### (*StrStrMap) LockFunc

```go
func (m *StrStrMap) LockFunc(f func(m map[string]string))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

​	LockFunc 在 RWMutex.Lock 中使用给定的回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (*StrStrMap) Map

```go
func (m *StrStrMap) Map() map[string]string
```

Map returns the underlying data map. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Map 返回基础数据映射。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

##### Example

``` go
```

#### (*StrStrMap) MapCopy

```go
func (m *StrStrMap) MapCopy() map[string]string
```

MapCopy returns a copy of the underlying data of the hash map.

​	MapCopy 返回哈希映射的基础数据的副本。

##### Example

``` go
```

#### (*StrStrMap) MapStrAny

```go
func (m *StrStrMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

​	MapStrAny 以 map[string]interface{} 的形式返回地图基础数据的副本。

##### Example

``` go
```

#### (StrStrMap) MarshalJSON

```go
func (m StrStrMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*StrStrMap) Merge

```go
func (m *StrStrMap) Merge(other *StrStrMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

​	合并合并两个哈希映射。 `other` 地图将合并到地图 `m` 中。

##### Example

``` go
```

#### (*StrStrMap) Pop

```go
func (m *StrStrMap) Pop() (key, value string)
```

Pop retrieves and deletes an item from the map.

​	Pop 从地图中检索和删除项目。

##### Example

``` go
```

#### (*StrStrMap) Pops

```go
func (m *StrStrMap) Pops(size int) map[string]string
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

​	Pops 从地图中检索和删除 `size` 项目。如果大小 == -1，则返回所有项目。

##### Example

``` go
```

#### (*StrStrMap) RLockFunc

```go
func (m *StrStrMap) RLockFunc(f func(m map[string]string))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

​	RLockFunc 在 RWMutex.RLock 中使用给定的回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*StrStrMap) Remove

```go
func (m *StrStrMap) Remove(key string) (value string)
```

Remove deletes value from map by given `key`, and return this deleted value.

​	按给定 `key` 从映射中删除删除值，并返回此删除值。

##### Example

``` go
```

#### (*StrStrMap) Removes

```go
func (m *StrStrMap) Removes(keys []string)
```

Removes batch deletes values of the map by keys.

​	删除按键批量删除映射的值。

##### Example

``` go
```

#### (*StrStrMap) Replace

```go
func (m *StrStrMap) Replace(data map[string]string)
```

Replace the data of the map with given `data`.

​	将地图的数据替换为给定 `data` 的 .

##### Example

``` go
```

#### (*StrStrMap) Search

```go
func (m *StrStrMap) Search(key string) (value string, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

​	搜索 使用给定 `key` 的 .如果找到键，则第二个返回参数 `found` 为 true，否则为 false。

##### Example

``` go
```

#### (*StrStrMap) Set

```go
func (m *StrStrMap) Set(key string, val string)
```

Set sets key-value to the hash map.

​	Set 将 key-value 设置为哈希映射。

##### Example

``` go
```

#### (*StrStrMap) SetIfNotExist

```go
func (m *StrStrMap) SetIfNotExist(key string, value string) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	如果 不存在 `key` ，则 SetIfNotExist 设置为 `value` 映射，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*StrStrMap) SetIfNotExistFunc

```go
func (m *StrStrMap) SetIfNotExistFunc(key string, f func() string) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFunc 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

##### Example

``` go
```

#### (*StrStrMap) SetIfNotExistFuncLock

```go
func (m *StrStrMap) SetIfNotExistFuncLock(key string, f func() string) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

​	SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。如果 `key` 存在，则返回 false，并将 `value` 被忽略。

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

​	SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的不同之处在于它使用互斥锁执行函数 `f` 。哈希映射的锁定。

##### Example

``` go
```

#### (*StrStrMap) Sets

```go
func (m *StrStrMap) Sets(data map[string]string)
```

Sets batch sets key-values to the hash map.

​	将批处理设置键值设置为哈希映射。

##### Example

``` go
```

#### (*StrStrMap) Size

```go
func (m *StrStrMap) Size() int
```

Size returns the size of the map.

​	Size 返回地图的大小。

##### Example

``` go
```

#### (*StrStrMap) String

```go
func (m *StrStrMap) String() string
```

String returns the map as a string.

​	String 以字符串形式返回映射。

##### Example

``` go
```

#### (*StrStrMap) UnmarshalJSON

```go
func (m *StrStrMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*StrStrMap) UnmarshalValue

```go
func (m *StrStrMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

​	UnmarshalValue 是一个接口实现，用于为 map 设置任何类型的值。

##### Example

``` go
```

#### (*StrStrMap) Values

```go
func (m *StrStrMap) Values() []string
```

Values returns all values of the map as a slice.

​	Values 以切片的形式返回地图的所有值。

##### Example

``` go
```

### type TreeMap

```go
type TreeMap = gtree.RedBlackTree
```

TreeMap based on red-black tree, alias of RedBlackTree.

​	基于红黑树的树状图，RedBlackTree的别名。

#### func NewTreeMap

```go
func NewTreeMap(comparator func(v1, v2 interface{}) int, safe ...bool) *TreeMap
```

NewTreeMap instantiates a tree map with the custom comparator. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

​	NewTreeMap 使用自定义比较器实例化树状图。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 tree，默认为 false。

##### Example

``` go
```

#### func NewTreeMapFrom

```go
func NewTreeMapFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *TreeMap
```

NewTreeMapFrom instantiates a tree map with the custom comparator and `data` map. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

​	NewTreeMapFrom 使用自定义比较器和 `data` 映射实例化树状图。请注意，参数 `data` 映射将设置为底层数据映射（无深度拷贝），在外部更改映射时可能会出现一些并发安全问题。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 tree，默认为 false。

##### Example

``` go
```








<iframe allowtransparency="true" frameborder="0" scrolling="no" class="sk_ui" src="chrome-extension://gfbliohnnapiefjpjlpjnehglfpaknnc/pages/frontend.html" title="Surfingkeys" style="left: 0px; bottom: 0px; width: 1555px; height: 0px; z-index: 2147483647;"></iframe>