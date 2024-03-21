+++
title = "gmap"
date = 2024-03-21T17:44:36+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gmap

### Overview 

Package gmap provides most commonly used map container which also support concurrent-safe/unsafe switch feature.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type AnyAnyMap 

``` go
type AnyAnyMap struct {
	// contains filtered or unexported fields
}
```

AnyAnyMap wraps map type `map[interface{}]interface{}` and provides more map features.

##### func NewAnyAnyMap 

``` go
func NewAnyAnyMap(safe ...bool) *AnyAnyMap
```

NewAnyAnyMap creates and returns an empty hash map. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewAnyAnyMapFrom 

``` go
func NewAnyAnyMapFrom(data map[interface{}]interface{}, safe ...bool) *AnyAnyMap
```

NewAnyAnyMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

##### Example

``` go
```
##### (*AnyAnyMap) Clear 

``` go
func (m *AnyAnyMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

##### Example

``` go
```
##### (*AnyAnyMap) Clone 

``` go
func (m *AnyAnyMap) Clone(safe ...bool) *AnyAnyMap
```

Clone returns a new hash map with copy of current map data.

##### Example

``` go
```
##### (*AnyAnyMap) Contains 

``` go
func (m *AnyAnyMap) Contains(key interface{}) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

##### Example

``` go
```
##### (*AnyAnyMap) DeepCopy <-2.1.0

``` go
func (m *AnyAnyMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*AnyAnyMap) Diff <-2.5.1

``` go
func (m *AnyAnyMap) Diff(other *AnyAnyMap) (addedKeys, removedKeys, updatedKeys []interface{})
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

##### (*AnyAnyMap) FilterEmpty 

``` go
func (m *AnyAnyMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.

##### Example

``` go
```
##### (*AnyAnyMap) FilterNil 

``` go
func (m *AnyAnyMap) FilterNil()
```

FilterNil deletes all key-value pair of which the value is nil.

##### Example

``` go
```
##### (*AnyAnyMap) Flip 

``` go
func (m *AnyAnyMap) Flip()
```

Flip exchanges key-value of the map to value-key.

##### Example

``` go
```
##### (*AnyAnyMap) Get 

``` go
func (m *AnyAnyMap) Get(key interface{}) (value interface{})
```

Get returns the value by given `key`.

##### Example

``` go
```
##### (*AnyAnyMap) GetOrSet 

``` go
func (m *AnyAnyMap) GetOrSet(key interface{}, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*AnyAnyMap) GetOrSetFunc 

``` go
func (m *AnyAnyMap) GetOrSetFunc(key interface{}, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*AnyAnyMap) GetOrSetFuncLock 

``` go
func (m *AnyAnyMap) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*AnyAnyMap) GetVar 

``` go
func (m *AnyAnyMap) GetVar(key interface{}) *gvar.Var
```

GetVar returns a Var with the value by given `key`. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*AnyAnyMap) GetVarOrSet 

``` go
func (m *AnyAnyMap) GetVarOrSet(key interface{}, value interface{}) *gvar.Var
```

GetVarOrSet returns a Var with result from GetOrSet. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*AnyAnyMap) GetVarOrSetFunc 

``` go
func (m *AnyAnyMap) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a Var with result from GetOrSetFunc. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*AnyAnyMap) GetVarOrSetFuncLock 

``` go
func (m *AnyAnyMap) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*AnyAnyMap) IsEmpty 

``` go
func (m *AnyAnyMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

##### Example

``` go
```
##### (*AnyAnyMap) IsSubOf <-2.3.3

``` go
func (m *AnyAnyMap) IsSubOf(other *AnyAnyMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

##### (*AnyAnyMap) Iterator 

``` go
func (m *AnyAnyMap) Iterator(f func(k interface{}, v interface{}) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*AnyAnyMap) Keys 

``` go
func (m *AnyAnyMap) Keys() []interface{}
```

Keys returns all keys of the map as a slice.

##### Example

``` go
```
##### (*AnyAnyMap) LockFunc 

``` go
func (m *AnyAnyMap) LockFunc(f func(m map[interface{}]interface{}))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

##### Example

``` go
```
##### (*AnyAnyMap) Map 

``` go
func (m *AnyAnyMap) Map() map[interface{}]interface{}
```

Map returns the underlying data map. Note that, if it's in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

##### Example

``` go
```
##### (*AnyAnyMap) MapCopy 

``` go
func (m *AnyAnyMap) MapCopy() map[interface{}]interface{}
```

MapCopy returns a shallow copy of the underlying data of the hash map.

##### Example

``` go
```
##### (*AnyAnyMap) MapStrAny 

``` go
func (m *AnyAnyMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

##### Example

``` go
```
##### (AnyAnyMap) MarshalJSON 

``` go
func (m AnyAnyMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*AnyAnyMap) Merge 

``` go
func (m *AnyAnyMap) Merge(other *AnyAnyMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

##### Example

``` go
```
##### (*AnyAnyMap) Pop 

``` go
func (m *AnyAnyMap) Pop() (key, value interface{})
```

Pop retrieves and deletes an item from the map.

##### Example

``` go
```
##### (*AnyAnyMap) Pops 

``` go
func (m *AnyAnyMap) Pops(size int) map[interface{}]interface{}
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

##### Example

``` go
```
##### (*AnyAnyMap) RLockFunc 

``` go
func (m *AnyAnyMap) RLockFunc(f func(m map[interface{}]interface{}))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

##### Example

``` go
```
##### (*AnyAnyMap) Remove 

``` go
func (m *AnyAnyMap) Remove(key interface{}) (value interface{})
```

Remove deletes value from map by given `key`, and return this deleted value.

##### Example

``` go
```
##### (*AnyAnyMap) Removes 

``` go
func (m *AnyAnyMap) Removes(keys []interface{})
```

Removes batch deletes values of the map by keys.

##### Example

``` go
```
##### (*AnyAnyMap) Replace 

``` go
func (m *AnyAnyMap) Replace(data map[interface{}]interface{})
```

Replace the data of the map with given `data`.

##### Example

``` go
```
##### (*AnyAnyMap) Search 

``` go
func (m *AnyAnyMap) Search(key interface{}) (value interface{}, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### Example

``` go
```
##### (*AnyAnyMap) Set 

``` go
func (m *AnyAnyMap) Set(key interface{}, value interface{})
```

Set sets key-value to the hash map.

##### Example

``` go
```
##### (*AnyAnyMap) SetIfNotExist 

``` go
func (m *AnyAnyMap) SetIfNotExist(key interface{}, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*AnyAnyMap) SetIfNotExistFunc 

``` go
func (m *AnyAnyMap) SetIfNotExistFunc(key interface{}, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*AnyAnyMap) SetIfNotExistFuncLock 

``` go
func (m *AnyAnyMap) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*AnyAnyMap) Sets 

``` go
func (m *AnyAnyMap) Sets(data map[interface{}]interface{})
```

Sets batch sets key-values to the hash map.

##### Example

``` go
```
##### (*AnyAnyMap) Size 

``` go
func (m *AnyAnyMap) Size() int
```

Size returns the size of the map.

##### Example

``` go
```
##### (*AnyAnyMap) String 

``` go
func (m *AnyAnyMap) String() string
```

String returns the map as a string.

##### Example

``` go
```
##### (*AnyAnyMap) UnmarshalJSON 

``` go
func (m *AnyAnyMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*AnyAnyMap) UnmarshalValue 

``` go
func (m *AnyAnyMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

##### Example

``` go
```
##### (*AnyAnyMap) Values 

``` go
func (m *AnyAnyMap) Values() []interface{}
```

Values returns all values of the map as a slice.

##### Example

``` go
```
#### type HashMap 

``` go
type HashMap = AnyAnyMap // HashMap is alias of AnyAnyMap.
```

#### type IntAnyMap 

``` go
type IntAnyMap struct {
	// contains filtered or unexported fields
}
```

IntAnyMap implements map[int]interface{} with RWMutex that has switch.

##### func NewIntAnyMap 

``` go
func NewIntAnyMap(safe ...bool) *IntAnyMap
```

NewIntAnyMap returns an empty IntAnyMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewIntAnyMapFrom 

``` go
func NewIntAnyMapFrom(data map[int]interface{}, safe ...bool) *IntAnyMap
```

NewIntAnyMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

##### Example

``` go
```
##### (*IntAnyMap) Clear 

``` go
func (m *IntAnyMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

##### Example

``` go
```
##### (*IntAnyMap) Clone 

``` go
func (m *IntAnyMap) Clone() *IntAnyMap
```

Clone returns a new hash map with copy of current map data.

##### Example

``` go
```
##### (*IntAnyMap) Contains 

``` go
func (m *IntAnyMap) Contains(key int) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

##### Example

``` go
```
##### (*IntAnyMap) DeepCopy <-2.1.0

``` go
func (m *IntAnyMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*IntAnyMap) Diff <-2.5.1

``` go
func (m *IntAnyMap) Diff(other *IntAnyMap) (addedKeys, removedKeys, updatedKeys []int)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

##### (*IntAnyMap) FilterEmpty 

``` go
func (m *IntAnyMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.

##### Example

``` go
```
##### (*IntAnyMap) FilterNil 

``` go
func (m *IntAnyMap) FilterNil()
```

FilterNil deletes all key-value pair of which the value is nil.

##### Example

``` go
```
##### (*IntAnyMap) Flip 

``` go
func (m *IntAnyMap) Flip()
```

Flip exchanges key-value of the map to value-key.

##### Example

``` go
```
##### (*IntAnyMap) Get 

``` go
func (m *IntAnyMap) Get(key int) (value interface{})
```

Get returns the value by given `key`.

##### Example

``` go
```
##### (*IntAnyMap) GetOrSet 

``` go
func (m *IntAnyMap) GetOrSet(key int, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*IntAnyMap) GetOrSetFunc 

``` go
func (m *IntAnyMap) GetOrSetFunc(key int, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

##### Example

``` go
```
##### (*IntAnyMap) GetOrSetFuncLock 

``` go
func (m *IntAnyMap) GetOrSetFuncLock(key int, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*IntAnyMap) GetVar 

``` go
func (m *IntAnyMap) GetVar(key int) *gvar.Var
```

GetVar returns a Var with the value by given `key`. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*IntAnyMap) GetVarOrSet 

``` go
func (m *IntAnyMap) GetVarOrSet(key int, value interface{}) *gvar.Var
```

GetVarOrSet returns a Var with result from GetVarOrSet. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*IntAnyMap) GetVarOrSetFunc 

``` go
func (m *IntAnyMap) GetVarOrSetFunc(key int, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a Var with result from GetOrSetFunc. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*IntAnyMap) GetVarOrSetFuncLock 

``` go
func (m *IntAnyMap) GetVarOrSetFuncLock(key int, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*IntAnyMap) IsEmpty 

``` go
func (m *IntAnyMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

##### Example

``` go
```
##### (*IntAnyMap) IsSubOf <-2.3.3

``` go
func (m *IntAnyMap) IsSubOf(other *IntAnyMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

##### (*IntAnyMap) Iterator 

``` go
func (m *IntAnyMap) Iterator(f func(k int, v interface{}) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*IntAnyMap) Keys 

``` go
func (m *IntAnyMap) Keys() []int
```

Keys returns all keys of the map as a slice.

##### Example

``` go
```
##### (*IntAnyMap) LockFunc 

``` go
func (m *IntAnyMap) LockFunc(f func(m map[int]interface{}))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

##### Example

``` go
```
##### (*IntAnyMap) Map 

``` go
func (m *IntAnyMap) Map() map[int]interface{}
```

Map returns the underlying data map. Note that, if it's in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

##### Example

``` go
```
##### (*IntAnyMap) MapCopy 

``` go
func (m *IntAnyMap) MapCopy() map[int]interface{}
```

MapCopy returns a copy of the underlying data of the hash map.

##### Example

``` go
```
##### (*IntAnyMap) MapStrAny 

``` go
func (m *IntAnyMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

##### Example

``` go
```
##### (IntAnyMap) MarshalJSON 

``` go
func (m IntAnyMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*IntAnyMap) Merge 

``` go
func (m *IntAnyMap) Merge(other *IntAnyMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

##### Example

``` go
```
##### (*IntAnyMap) Pop 

``` go
func (m *IntAnyMap) Pop() (key int, value interface{})
```

Pop retrieves and deletes an item from the map.

##### Example

``` go
```
##### (*IntAnyMap) Pops 

``` go
func (m *IntAnyMap) Pops(size int) map[int]interface{}
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

##### Example

``` go
```
##### (*IntAnyMap) RLockFunc 

``` go
func (m *IntAnyMap) RLockFunc(f func(m map[int]interface{}))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

##### Example

``` go
```
##### (*IntAnyMap) Remove 

``` go
func (m *IntAnyMap) Remove(key int) (value interface{})
```

Remove deletes value from map by given `key`, and return this deleted value.

##### Example

``` go
```
##### (*IntAnyMap) Removes 

``` go
func (m *IntAnyMap) Removes(keys []int)
```

Removes batch deletes values of the map by keys.

##### Example

``` go
```
##### (*IntAnyMap) Replace 

``` go
func (m *IntAnyMap) Replace(data map[int]interface{})
```

Replace the data of the map with given `data`.

##### Example

``` go
```
##### (*IntAnyMap) Search 

``` go
func (m *IntAnyMap) Search(key int) (value interface{}, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### Example

``` go
```
##### (*IntAnyMap) Set 

``` go
func (m *IntAnyMap) Set(key int, val interface{})
```

Set sets key-value to the hash map.

##### Example

``` go
```
##### (*IntAnyMap) SetIfNotExist 

``` go
func (m *IntAnyMap) SetIfNotExist(key int, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*IntAnyMap) SetIfNotExistFunc 

``` go
func (m *IntAnyMap) SetIfNotExistFunc(key int, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*IntAnyMap) SetIfNotExistFuncLock 

``` go
func (m *IntAnyMap) SetIfNotExistFuncLock(key int, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*IntAnyMap) Sets 

``` go
func (m *IntAnyMap) Sets(data map[int]interface{})
```

Sets batch sets key-values to the hash map.

##### Example

``` go
```
##### (*IntAnyMap) Size 

``` go
func (m *IntAnyMap) Size() int
```

Size returns the size of the map.

##### Example

``` go
```
##### (*IntAnyMap) String 

``` go
func (m *IntAnyMap) String() string
```

String returns the map as a string.

##### Example

``` go
```
##### (*IntAnyMap) UnmarshalJSON 

``` go
func (m *IntAnyMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*IntAnyMap) UnmarshalValue 

``` go
func (m *IntAnyMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

##### Example

``` go
```
##### (*IntAnyMap) Values 

``` go
func (m *IntAnyMap) Values() []interface{}
```

Values returns all values of the map as a slice.

##### Example

``` go
```
#### type IntIntMap 

``` go
type IntIntMap struct {
	// contains filtered or unexported fields
}
```

IntIntMap implements map[int]int with RWMutex that has switch.

##### func NewIntIntMap 

``` go
func NewIntIntMap(safe ...bool) *IntIntMap
```

NewIntIntMap returns an empty IntIntMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewIntIntMapFrom 

``` go
func NewIntIntMapFrom(data map[int]int, safe ...bool) *IntIntMap
```

NewIntIntMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

##### Example

``` go
```
##### (*IntIntMap) Clear 

``` go
func (m *IntIntMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

##### Example

``` go
```
##### (*IntIntMap) Clone 

``` go
func (m *IntIntMap) Clone() *IntIntMap
```

Clone returns a new hash map with copy of current map data.

##### Example

``` go
```
##### (*IntIntMap) Contains 

``` go
func (m *IntIntMap) Contains(key int) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

##### Example

``` go
```
##### (*IntIntMap) DeepCopy <-2.1.0

``` go
func (m *IntIntMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*IntIntMap) Diff <-2.5.1

``` go
func (m *IntIntMap) Diff(other *IntIntMap) (addedKeys, removedKeys, updatedKeys []int)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

##### (*IntIntMap) FilterEmpty 

``` go
func (m *IntIntMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.

##### Example

``` go
```
##### (*IntIntMap) Flip 

``` go
func (m *IntIntMap) Flip()
```

Flip exchanges key-value of the map to value-key.

##### Example

``` go
```
##### (*IntIntMap) Get 

``` go
func (m *IntIntMap) Get(key int) (value int)
```

Get returns the value by given `key`.

##### Example

``` go
```
##### (*IntIntMap) GetOrSet 

``` go
func (m *IntIntMap) GetOrSet(key int, value int) int
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*IntIntMap) GetOrSetFunc 

``` go
func (m *IntIntMap) GetOrSetFunc(key int, f func() int) int
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

##### Example

``` go
```
##### (*IntIntMap) GetOrSetFuncLock 

``` go
func (m *IntIntMap) GetOrSetFuncLock(key int, f func() int) int
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*IntIntMap) IsEmpty 

``` go
func (m *IntIntMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

##### Example

``` go
```
##### (*IntIntMap) IsSubOf <-2.3.3

``` go
func (m *IntIntMap) IsSubOf(other *IntIntMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

##### (*IntIntMap) Iterator 

``` go
func (m *IntIntMap) Iterator(f func(k int, v int) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*IntIntMap) Keys 

``` go
func (m *IntIntMap) Keys() []int
```

Keys returns all keys of the map as a slice.

##### Example

``` go
```
##### (*IntIntMap) LockFunc 

``` go
func (m *IntIntMap) LockFunc(f func(m map[int]int))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

##### Example

``` go
```
##### (*IntIntMap) Map 

``` go
func (m *IntIntMap) Map() map[int]int
```

Map returns the underlying data map. Note that, if it's in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

##### Example

``` go
```
##### (*IntIntMap) MapCopy 

``` go
func (m *IntIntMap) MapCopy() map[int]int
```

MapCopy returns a copy of the underlying data of the hash map.

##### Example

``` go
```
##### (*IntIntMap) MapStrAny 

``` go
func (m *IntIntMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

##### Example

``` go
```
##### (IntIntMap) MarshalJSON 

``` go
func (m IntIntMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*IntIntMap) Merge 

``` go
func (m *IntIntMap) Merge(other *IntIntMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

##### Example

``` go
```
##### (*IntIntMap) Pop 

``` go
func (m *IntIntMap) Pop() (key, value int)
```

Pop retrieves and deletes an item from the map.

##### Example

``` go
```
##### (*IntIntMap) Pops 

``` go
func (m *IntIntMap) Pops(size int) map[int]int
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

##### Example

``` go
```
##### (*IntIntMap) RLockFunc 

``` go
func (m *IntIntMap) RLockFunc(f func(m map[int]int))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

##### Example

``` go
```
##### (*IntIntMap) Remove 

``` go
func (m *IntIntMap) Remove(key int) (value int)
```

Remove deletes value from map by given `key`, and return this deleted value.

##### Example

``` go
```
##### (*IntIntMap) Removes 

``` go
func (m *IntIntMap) Removes(keys []int)
```

Removes batch deletes values of the map by keys.

##### Example

``` go
```
##### (*IntIntMap) Replace 

``` go
func (m *IntIntMap) Replace(data map[int]int)
```

Replace the data of the map with given `data`.

##### Example

``` go
```
##### (*IntIntMap) Search 

``` go
func (m *IntIntMap) Search(key int) (value int, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### Example

``` go
```
##### (*IntIntMap) Set 

``` go
func (m *IntIntMap) Set(key int, val int)
```

Set sets key-value to the hash map.

##### Example

``` go
```
##### (*IntIntMap) SetIfNotExist 

``` go
func (m *IntIntMap) SetIfNotExist(key int, value int) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*IntIntMap) SetIfNotExistFunc 

``` go
func (m *IntIntMap) SetIfNotExistFunc(key int, f func() int) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*IntIntMap) SetIfNotExistFuncLock 

``` go
func (m *IntIntMap) SetIfNotExistFuncLock(key int, f func() int) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*IntIntMap) Sets 

``` go
func (m *IntIntMap) Sets(data map[int]int)
```

Sets batch sets key-values to the hash map.

##### Example

``` go
```
##### (*IntIntMap) Size 

``` go
func (m *IntIntMap) Size() int
```

Size returns the size of the map.

##### Example

``` go
```
##### (*IntIntMap) String 

``` go
func (m *IntIntMap) String() string
```

String returns the map as a string.

##### Example

``` go
```
##### (*IntIntMap) UnmarshalJSON 

``` go
func (m *IntIntMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*IntIntMap) UnmarshalValue 

``` go
func (m *IntIntMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

##### Example

``` go
```
##### (*IntIntMap) Values 

``` go
func (m *IntIntMap) Values() []int
```

Values returns all values of the map as a slice.

##### Example

``` go
```
#### type IntStrMap 

``` go
type IntStrMap struct {
	// contains filtered or unexported fields
}
```

IntStrMap implements map[int]string with RWMutex that has switch.

##### func NewIntStrMap 

``` go
func NewIntStrMap(safe ...bool) *IntStrMap
```

NewIntStrMap returns an empty IntStrMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

##### func NewIntStrMapFrom 

``` go
func NewIntStrMapFrom(data map[int]string, safe ...bool) *IntStrMap
```

NewIntStrMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

##### (*IntStrMap) Clear 

``` go
func (m *IntStrMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

##### (*IntStrMap) Clone 

``` go
func (m *IntStrMap) Clone() *IntStrMap
```

Clone returns a new hash map with copy of current map data.

##### (*IntStrMap) Contains 

``` go
func (m *IntStrMap) Contains(key int) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

##### (*IntStrMap) DeepCopy <-2.1.0

``` go
func (m *IntStrMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*IntStrMap) Diff <-2.5.1

``` go
func (m *IntStrMap) Diff(other *IntStrMap) (addedKeys, removedKeys, updatedKeys []int)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

##### (*IntStrMap) FilterEmpty 

``` go
func (m *IntStrMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.

##### (*IntStrMap) Flip 

``` go
func (m *IntStrMap) Flip()
```

Flip exchanges key-value of the map to value-key.

##### (*IntStrMap) Get 

``` go
func (m *IntStrMap) Get(key int) (value string)
```

Get returns the value by given `key`.

##### (*IntStrMap) GetOrSet 

``` go
func (m *IntStrMap) GetOrSet(key int, value string) string
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### (*IntStrMap) GetOrSetFunc 

``` go
func (m *IntStrMap) GetOrSetFunc(key int, f func() string) string
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

##### (*IntStrMap) GetOrSetFuncLock 

``` go
func (m *IntStrMap) GetOrSetFuncLock(key int, f func() string) string
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### (*IntStrMap) IsEmpty 

``` go
func (m *IntStrMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

##### (*IntStrMap) IsSubOf <-2.3.3

``` go
func (m *IntStrMap) IsSubOf(other *IntStrMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

##### (*IntStrMap) Iterator 

``` go
func (m *IntStrMap) Iterator(f func(k int, v string) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### (*IntStrMap) Keys 

``` go
func (m *IntStrMap) Keys() []int
```

Keys returns all keys of the map as a slice.

##### (*IntStrMap) LockFunc 

``` go
func (m *IntStrMap) LockFunc(f func(m map[int]string))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

##### (*IntStrMap) Map 

``` go
func (m *IntStrMap) Map() map[int]string
```

Map returns the underlying data map. Note that, if it's in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

##### (*IntStrMap) MapCopy 

``` go
func (m *IntStrMap) MapCopy() map[int]string
```

MapCopy returns a copy of the underlying data of the hash map.

##### (*IntStrMap) MapStrAny 

``` go
func (m *IntStrMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

##### (IntStrMap) MarshalJSON 

``` go
func (m IntStrMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*IntStrMap) Merge 

``` go
func (m *IntStrMap) Merge(other *IntStrMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

##### (*IntStrMap) Pop 

``` go
func (m *IntStrMap) Pop() (key int, value string)
```

Pop retrieves and deletes an item from the map.

##### (*IntStrMap) Pops 

``` go
func (m *IntStrMap) Pops(size int) map[int]string
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

##### (*IntStrMap) RLockFunc 

``` go
func (m *IntStrMap) RLockFunc(f func(m map[int]string))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

##### (*IntStrMap) Remove 

``` go
func (m *IntStrMap) Remove(key int) (value string)
```

Remove deletes value from map by given `key`, and return this deleted value.

##### (*IntStrMap) Removes 

``` go
func (m *IntStrMap) Removes(keys []int)
```

Removes batch deletes values of the map by keys.

##### (*IntStrMap) Replace 

``` go
func (m *IntStrMap) Replace(data map[int]string)
```

Replace the data of the map with given `data`.

##### (*IntStrMap) Search 

``` go
func (m *IntStrMap) Search(key int) (value string, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### (*IntStrMap) Set 

``` go
func (m *IntStrMap) Set(key int, val string)
```

Set sets key-value to the hash map.

##### (*IntStrMap) SetIfNotExist 

``` go
func (m *IntStrMap) SetIfNotExist(key int, value string) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### (*IntStrMap) SetIfNotExistFunc 

``` go
func (m *IntStrMap) SetIfNotExistFunc(key int, f func() string) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### (*IntStrMap) SetIfNotExistFuncLock 

``` go
func (m *IntStrMap) SetIfNotExistFuncLock(key int, f func() string) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### (*IntStrMap) Sets 

``` go
func (m *IntStrMap) Sets(data map[int]string)
```

Sets batch sets key-values to the hash map.

##### (*IntStrMap) Size 

``` go
func (m *IntStrMap) Size() int
```

Size returns the size of the map.

##### (*IntStrMap) String 

``` go
func (m *IntStrMap) String() string
```

String returns the map as a string.

##### (*IntStrMap) UnmarshalJSON 

``` go
func (m *IntStrMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*IntStrMap) UnmarshalValue 

``` go
func (m *IntStrMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

##### (*IntStrMap) Values 

``` go
func (m *IntStrMap) Values() []string
```

Values returns all values of the map as a slice.

#### type ListMap 

``` go
type ListMap struct {
	// contains filtered or unexported fields
}
```

ListMap is a map that preserves insertion-order.

It is backed by a hash table to store values and doubly-linked list to store ordering.

Structure is not thread safe.

Reference: http://en.wikipedia.org/wiki/Associative_array

##### func NewListMap 

``` go
func NewListMap(safe ...bool) *ListMap
```

NewListMap returns an empty link map. ListMap is backed by a hash table to store values and doubly-linked list to store ordering. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewListMapFrom 

``` go
func NewListMapFrom(data map[interface{}]interface{}, safe ...bool) *ListMap
```

NewListMapFrom returns a link map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

##### Example

``` go
```
##### (*ListMap) Clear 

``` go
func (m *ListMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

##### Example

``` go
```
##### (*ListMap) Clone 

``` go
func (m *ListMap) Clone(safe ...bool) *ListMap
```

Clone returns a new link map with copy of current map data.

##### Example

``` go
```
##### (*ListMap) Contains 

``` go
func (m *ListMap) Contains(key interface{}) (ok bool)
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

##### Example

``` go
```
##### (*ListMap) DeepCopy <-2.1.0

``` go
func (m *ListMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*ListMap) FilterEmpty 

``` go
func (m *ListMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty.

##### Example

``` go
```
##### (*ListMap) Flip 

``` go
func (m *ListMap) Flip()
```

Flip exchanges key-value of the map to value-key.

##### Example

``` go
```
##### (*ListMap) Get 

``` go
func (m *ListMap) Get(key interface{}) (value interface{})
```

Get returns the value by given `key`.

##### Example

``` go
```
##### (*ListMap) GetOrSet 

``` go
func (m *ListMap) GetOrSet(key interface{}, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*ListMap) GetOrSetFunc 

``` go
func (m *ListMap) GetOrSetFunc(key interface{}, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*ListMap) GetOrSetFuncLock 

``` go
func (m *ListMap) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the map.

##### Example

``` go
```
##### (*ListMap) GetVar 

``` go
func (m *ListMap) GetVar(key interface{}) *gvar.Var
```

GetVar returns a Var with the value by given `key`. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*ListMap) GetVarOrSet 

``` go
func (m *ListMap) GetVarOrSet(key interface{}, value interface{}) *gvar.Var
```

GetVarOrSet returns a Var with result from GetVarOrSet. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*ListMap) GetVarOrSetFunc 

``` go
func (m *ListMap) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a Var with result from GetOrSetFunc. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*ListMap) GetVarOrSetFuncLock 

``` go
func (m *ListMap) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*ListMap) IsEmpty 

``` go
func (m *ListMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

##### Example

``` go
```
##### (*ListMap) Iterator 

``` go
func (m *ListMap) Iterator(f func(key, value interface{}) bool)
```

Iterator is alias of IteratorAsc.

##### Example

``` go
```
##### (*ListMap) IteratorAsc 

``` go
func (m *ListMap) IteratorAsc(f func(key interface{}, value interface{}) bool)
```

IteratorAsc iterates the map readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*ListMap) IteratorDesc 

``` go
func (m *ListMap) IteratorDesc(f func(key interface{}, value interface{}) bool)
```

IteratorDesc iterates the map readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*ListMap) Keys 

``` go
func (m *ListMap) Keys() []interface{}
```

Keys returns all keys of the map as a slice in ascending order.

##### Example

``` go
```
##### (*ListMap) Map 

``` go
func (m *ListMap) Map() map[interface{}]interface{}
```

Map returns a copy of the underlying data of the map.

##### Example

``` go
```
##### (*ListMap) MapStrAny 

``` go
func (m *ListMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

##### Example

``` go
```
##### (ListMap) MarshalJSON 

``` go
func (m ListMap) MarshalJSON() (jsonBytes []byte, err error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*ListMap) Merge 

``` go
func (m *ListMap) Merge(other *ListMap)
```

Merge merges two link maps. The `other` map will be merged into the map `m`.

##### Example

``` go
```
##### (*ListMap) Pop 

``` go
func (m *ListMap) Pop() (key, value interface{})
```

Pop retrieves and deletes an item from the map.

##### Example

``` go
```
##### (*ListMap) Pops 

``` go
func (m *ListMap) Pops(size int) map[interface{}]interface{}
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

##### Example

``` go
```
##### (*ListMap) Remove 

``` go
func (m *ListMap) Remove(key interface{}) (value interface{})
```

Remove deletes value from map by given `key`, and return this deleted value.

##### Example

``` go
```
##### (*ListMap) Removes 

``` go
func (m *ListMap) Removes(keys []interface{})
```

Removes batch deletes values of the map by keys.

##### Example

``` go
```
##### (*ListMap) Replace 

``` go
func (m *ListMap) Replace(data map[interface{}]interface{})
```

Replace the data of the map with given `data`.

##### Example

``` go
```
##### (*ListMap) Search 

``` go
func (m *ListMap) Search(key interface{}) (value interface{}, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### Example

``` go
```
##### (*ListMap) Set 

``` go
func (m *ListMap) Set(key interface{}, value interface{})
```

Set sets key-value to the map.

##### Example

``` go
```
##### (*ListMap) SetIfNotExist 

``` go
func (m *ListMap) SetIfNotExist(key interface{}, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*ListMap) SetIfNotExistFunc 

``` go
func (m *ListMap) SetIfNotExistFunc(key interface{}, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*ListMap) SetIfNotExistFuncLock 

``` go
func (m *ListMap) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the map.

##### Example

``` go
```
##### (*ListMap) Sets 

``` go
func (m *ListMap) Sets(data map[interface{}]interface{})
```

Sets batch sets key-values to the map.

##### Example

``` go
```
##### (*ListMap) Size 

``` go
func (m *ListMap) Size() (size int)
```

Size returns the size of the map.

##### Example

``` go
```
##### (*ListMap) String 

``` go
func (m *ListMap) String() string
```

String returns the map as a string.

##### Example

``` go
```
##### (*ListMap) UnmarshalJSON 

``` go
func (m *ListMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*ListMap) UnmarshalValue 

``` go
func (m *ListMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

##### Example

``` go
```
##### (*ListMap) Values 

``` go
func (m *ListMap) Values() []interface{}
```

Values returns all values of the map as a slice.

##### Example

``` go
```
#### type Map 

``` go
type Map = AnyAnyMap // Map is alias of AnyAnyMap.
```

##### func New 

``` go
func New(safe ...bool) *Map
```

New creates and returns an empty hash map. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewFrom 

``` go
func NewFrom(data map[interface{}]interface{}, safe ...bool) *Map
```

NewFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewHashMap 

``` go
func NewHashMap(safe ...bool) *Map
```

NewHashMap creates and returns an empty hash map. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewHashMapFrom 

``` go
func NewHashMapFrom(data map[interface{}]interface{}, safe ...bool) *Map
```

NewHashMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

##### Example

``` go
```
#### type StrAnyMap 

``` go
type StrAnyMap struct {
	// contains filtered or unexported fields
}
```

StrAnyMap implements map[string]interface{} with RWMutex that has switch.

##### func NewStrAnyMap 

``` go
func NewStrAnyMap(safe ...bool) *StrAnyMap
```

NewStrAnyMap returns an empty StrAnyMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewStrAnyMapFrom 

``` go
func NewStrAnyMapFrom(data map[string]interface{}, safe ...bool) *StrAnyMap
```

NewStrAnyMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

##### Example

``` go
```
##### (*StrAnyMap) Clear 

``` go
func (m *StrAnyMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

##### Example

``` go
```
##### (*StrAnyMap) Clone 

``` go
func (m *StrAnyMap) Clone() *StrAnyMap
```

Clone returns a new hash map with copy of current map data.

##### Example

``` go
```
##### (*StrAnyMap) Contains 

``` go
func (m *StrAnyMap) Contains(key string) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

##### Example

``` go
```
##### (*StrAnyMap) DeepCopy <-2.1.0

``` go
func (m *StrAnyMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*StrAnyMap) Diff <-2.5.1

``` go
func (m *StrAnyMap) Diff(other *StrAnyMap) (addedKeys, removedKeys, updatedKeys []string)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

##### (*StrAnyMap) FilterEmpty 

``` go
func (m *StrAnyMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.

##### Example

``` go
```
##### (*StrAnyMap) FilterNil 

``` go
func (m *StrAnyMap) FilterNil()
```

FilterNil deletes all key-value pair of which the value is nil.

##### Example

``` go
```
##### (*StrAnyMap) Flip 

``` go
func (m *StrAnyMap) Flip()
```

Flip exchanges key-value of the map to value-key.

##### Example

``` go
```
##### (*StrAnyMap) Get 

``` go
func (m *StrAnyMap) Get(key string) (value interface{})
```

Get returns the value by given `key`.

##### Example

``` go
```
##### (*StrAnyMap) GetOrSet 

``` go
func (m *StrAnyMap) GetOrSet(key string, value interface{}) interface{}
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*StrAnyMap) GetOrSetFunc 

``` go
func (m *StrAnyMap) GetOrSetFunc(key string, f func() interface{}) interface{}
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*StrAnyMap) GetOrSetFuncLock 

``` go
func (m *StrAnyMap) GetOrSetFuncLock(key string, f func() interface{}) interface{}
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*StrAnyMap) GetVar 

``` go
func (m *StrAnyMap) GetVar(key string) *gvar.Var
```

GetVar returns a Var with the value by given `key`. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*StrAnyMap) GetVarOrSet 

``` go
func (m *StrAnyMap) GetVarOrSet(key string, value interface{}) *gvar.Var
```

GetVarOrSet returns a Var with result from GetVarOrSet. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*StrAnyMap) GetVarOrSetFunc 

``` go
func (m *StrAnyMap) GetVarOrSetFunc(key string, f func() interface{}) *gvar.Var
```

GetVarOrSetFunc returns a Var with result from GetOrSetFunc. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*StrAnyMap) GetVarOrSetFuncLock 

``` go
func (m *StrAnyMap) GetVarOrSetFuncLock(key string, f func() interface{}) *gvar.Var
```

GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock. The returned Var is un-concurrent safe.

##### Example

``` go
```
##### (*StrAnyMap) IsEmpty 

``` go
func (m *StrAnyMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

##### Example

``` go
```
##### (*StrAnyMap) IsSubOf <-2.3.3

``` go
func (m *StrAnyMap) IsSubOf(other *StrAnyMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

##### (*StrAnyMap) Iterator 

``` go
func (m *StrAnyMap) Iterator(f func(k string, v interface{}) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*StrAnyMap) Keys 

``` go
func (m *StrAnyMap) Keys() []string
```

Keys returns all keys of the map as a slice.

##### Example

``` go
```
##### (*StrAnyMap) LockFunc 

``` go
func (m *StrAnyMap) LockFunc(f func(m map[string]interface{}))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

##### Example

``` go
```
##### (*StrAnyMap) Map 

``` go
func (m *StrAnyMap) Map() map[string]interface{}
```

Map returns the underlying data map. Note that, if it's in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

##### Example

``` go
```
##### (*StrAnyMap) MapCopy 

``` go
func (m *StrAnyMap) MapCopy() map[string]interface{}
```

MapCopy returns a copy of the underlying data of the hash map.

##### Example

``` go
```
##### (*StrAnyMap) MapStrAny 

``` go
func (m *StrAnyMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

##### Example

``` go
```
##### (StrAnyMap) MarshalJSON 

``` go
func (m StrAnyMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*StrAnyMap) Merge 

``` go
func (m *StrAnyMap) Merge(other *StrAnyMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

##### Example

``` go
```
##### (*StrAnyMap) Pop 

``` go
func (m *StrAnyMap) Pop() (key string, value interface{})
```

Pop retrieves and deletes an item from the map.

##### Example

``` go
```
##### (*StrAnyMap) Pops 

``` go
func (m *StrAnyMap) Pops(size int) map[string]interface{}
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

##### Example

``` go
```
##### (*StrAnyMap) RLockFunc 

``` go
func (m *StrAnyMap) RLockFunc(f func(m map[string]interface{}))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

##### Example

``` go
```
##### (*StrAnyMap) Remove 

``` go
func (m *StrAnyMap) Remove(key string) (value interface{})
```

Remove deletes value from map by given `key`, and return this deleted value.

##### Example

``` go
```
##### (*StrAnyMap) Removes 

``` go
func (m *StrAnyMap) Removes(keys []string)
```

Removes batch deletes values of the map by keys.

##### Example

``` go
```
##### (*StrAnyMap) Replace 

``` go
func (m *StrAnyMap) Replace(data map[string]interface{})
```

Replace the data of the map with given `data`.

##### Example

``` go
```
##### (*StrAnyMap) Search 

``` go
func (m *StrAnyMap) Search(key string) (value interface{}, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### Example

``` go
```
##### (*StrAnyMap) Set 

``` go
func (m *StrAnyMap) Set(key string, val interface{})
```

Set sets key-value to the hash map.

##### Example

``` go
```
##### (*StrAnyMap) SetIfNotExist 

``` go
func (m *StrAnyMap) SetIfNotExist(key string, value interface{}) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*StrAnyMap) SetIfNotExistFunc 

``` go
func (m *StrAnyMap) SetIfNotExistFunc(key string, f func() interface{}) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*StrAnyMap) SetIfNotExistFuncLock 

``` go
func (m *StrAnyMap) SetIfNotExistFuncLock(key string, f func() interface{}) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*StrAnyMap) Sets 

``` go
func (m *StrAnyMap) Sets(data map[string]interface{})
```

Sets batch sets key-values to the hash map.

##### Example

``` go
```
##### (*StrAnyMap) Size 

``` go
func (m *StrAnyMap) Size() int
```

Size returns the size of the map.

##### Example

``` go
```
##### (*StrAnyMap) String 

``` go
func (m *StrAnyMap) String() string
```

String returns the map as a string.

##### Example

``` go
```
##### (*StrAnyMap) UnmarshalJSON 

``` go
func (m *StrAnyMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*StrAnyMap) UnmarshalValue 

``` go
func (m *StrAnyMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

##### Example

``` go
```
##### (*StrAnyMap) Values 

``` go
func (m *StrAnyMap) Values() []interface{}
```

Values returns all values of the map as a slice.

##### Example

``` go
```
#### type StrIntMap 

``` go
type StrIntMap struct {
	// contains filtered or unexported fields
}
```

StrIntMap implements map[string]int with RWMutex that has switch.

##### func NewStrIntMap 

``` go
func NewStrIntMap(safe ...bool) *StrIntMap
```

NewStrIntMap returns an empty StrIntMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewStrIntMapFrom 

``` go
func NewStrIntMapFrom(data map[string]int, safe ...bool) *StrIntMap
```

NewStrIntMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

##### Example

``` go
```
##### (*StrIntMap) Clear 

``` go
func (m *StrIntMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

##### Example

``` go
```
##### (*StrIntMap) Clone 

``` go
func (m *StrIntMap) Clone() *StrIntMap
```

Clone returns a new hash map with copy of current map data.

##### Example

``` go
```
##### (*StrIntMap) Contains 

``` go
func (m *StrIntMap) Contains(key string) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

##### Example

``` go
```
##### (*StrIntMap) DeepCopy <-2.1.0

``` go
func (m *StrIntMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*StrIntMap) Diff <-2.5.1

``` go
func (m *StrIntMap) Diff(other *StrIntMap) (addedKeys, removedKeys, updatedKeys []string)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

##### (*StrIntMap) FilterEmpty 

``` go
func (m *StrIntMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.

##### Example

``` go
```
##### (*StrIntMap) Flip 

``` go
func (m *StrIntMap) Flip()
```

Flip exchanges key-value of the map to value-key.

##### Example

``` go
```
##### (*StrIntMap) Get 

``` go
func (m *StrIntMap) Get(key string) (value int)
```

Get returns the value by given `key`.

##### Example

``` go
```
##### (*StrIntMap) GetOrSet 

``` go
func (m *StrIntMap) GetOrSet(key string, value int) int
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*StrIntMap) GetOrSetFunc 

``` go
func (m *StrIntMap) GetOrSetFunc(key string, f func() int) int
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*StrIntMap) GetOrSetFuncLock 

``` go
func (m *StrIntMap) GetOrSetFuncLock(key string, f func() int) int
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*StrIntMap) IsEmpty 

``` go
func (m *StrIntMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

##### Example

``` go
```
##### (*StrIntMap) IsSubOf <-2.3.3

``` go
func (m *StrIntMap) IsSubOf(other *StrIntMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

##### (*StrIntMap) Iterator 

``` go
func (m *StrIntMap) Iterator(f func(k string, v int) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*StrIntMap) Keys 

``` go
func (m *StrIntMap) Keys() []string
```

Keys returns all keys of the map as a slice.

##### Example

``` go
```
##### (*StrIntMap) LockFunc 

``` go
func (m *StrIntMap) LockFunc(f func(m map[string]int))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

##### Example

``` go
```
##### (*StrIntMap) Map 

``` go
func (m *StrIntMap) Map() map[string]int
```

Map returns the underlying data map. Note that, if it's in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

##### Example

``` go
```
##### (*StrIntMap) MapCopy 

``` go
func (m *StrIntMap) MapCopy() map[string]int
```

MapCopy returns a copy of the underlying data of the hash map.

##### Example

``` go
```
##### (*StrIntMap) MapStrAny 

``` go
func (m *StrIntMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

##### Example

``` go
```
##### (StrIntMap) MarshalJSON 

``` go
func (m StrIntMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*StrIntMap) Merge 

``` go
func (m *StrIntMap) Merge(other *StrIntMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

##### Example

``` go
```
##### (*StrIntMap) Pop 

``` go
func (m *StrIntMap) Pop() (key string, value int)
```

Pop retrieves and deletes an item from the map.

##### Example

``` go
```
##### (*StrIntMap) Pops 

``` go
func (m *StrIntMap) Pops(size int) map[string]int
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

##### Example

``` go
```
##### (*StrIntMap) RLockFunc 

``` go
func (m *StrIntMap) RLockFunc(f func(m map[string]int))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

##### Example

``` go
```
##### (*StrIntMap) Remove 

``` go
func (m *StrIntMap) Remove(key string) (value int)
```

Remove deletes value from map by given `key`, and return this deleted value.

##### Example

``` go
```
##### (*StrIntMap) Removes 

``` go
func (m *StrIntMap) Removes(keys []string)
```

Removes batch deletes values of the map by keys.

##### Example

``` go
```
##### (*StrIntMap) Replace 

``` go
func (m *StrIntMap) Replace(data map[string]int)
```

Replace the data of the map with given `data`.

##### Example

``` go
```
##### (*StrIntMap) Search 

``` go
func (m *StrIntMap) Search(key string) (value int, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### Example

``` go
```
##### (*StrIntMap) Set 

``` go
func (m *StrIntMap) Set(key string, val int)
```

Set sets key-value to the hash map.

##### Example

``` go
```
##### (*StrIntMap) SetIfNotExist 

``` go
func (m *StrIntMap) SetIfNotExist(key string, value int) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*StrIntMap) SetIfNotExistFunc 

``` go
func (m *StrIntMap) SetIfNotExistFunc(key string, f func() int) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*StrIntMap) SetIfNotExistFuncLock 

``` go
func (m *StrIntMap) SetIfNotExistFuncLock(key string, f func() int) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*StrIntMap) Sets 

``` go
func (m *StrIntMap) Sets(data map[string]int)
```

Sets batch sets key-values to the hash map.

##### Example

``` go
```
##### (*StrIntMap) Size 

``` go
func (m *StrIntMap) Size() int
```

Size returns the size of the map.

##### Example

``` go
```
##### (*StrIntMap) String 

``` go
func (m *StrIntMap) String() string
```

String returns the map as a string.

##### Example

``` go
```
##### (*StrIntMap) UnmarshalJSON 

``` go
func (m *StrIntMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*StrIntMap) UnmarshalValue 

``` go
func (m *StrIntMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

##### Example

``` go
```
##### (*StrIntMap) Values 

``` go
func (m *StrIntMap) Values() []int
```

Values returns all values of the map as a slice.

##### Example

``` go
```
#### type StrStrMap 

``` go
type StrStrMap struct {
	// contains filtered or unexported fields
}
```

StrStrMap implements map[string]string with RWMutex that has switch.

##### func NewStrStrMap 

``` go
func NewStrStrMap(safe ...bool) *StrStrMap
```

NewStrStrMap returns an empty StrStrMap object. The parameter `safe` is used to specify whether using map in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewStrStrMapFrom 

``` go
func NewStrStrMapFrom(data map[string]string, safe ...bool) *StrStrMap
```

NewStrStrMapFrom creates and returns a hash map from given map `data`. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside.

##### Example

``` go
```
##### (*StrStrMap) Clear 

``` go
func (m *StrStrMap) Clear()
```

Clear deletes all data of the map, it will remake a new underlying data map.

##### Example

``` go
```
##### (*StrStrMap) Clone 

``` go
func (m *StrStrMap) Clone() *StrStrMap
```

Clone returns a new hash map with copy of current map data.

##### Example

``` go
```
##### (*StrStrMap) Contains 

``` go
func (m *StrStrMap) Contains(key string) bool
```

Contains checks whether a key exists. It returns true if the `key` exists, or else false.

##### Example

``` go
```
##### (*StrStrMap) DeepCopy <-2.1.0

``` go
func (m *StrStrMap) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*StrStrMap) Diff <-2.5.1

``` go
func (m *StrStrMap) Diff(other *StrStrMap) (addedKeys, removedKeys, updatedKeys []string)
```

Diff compares current map `m` with map `other` and returns their different keys. The returned `addedKeys` are the keys that are in map `m` but not in map `other`. The returned `removedKeys` are the keys that are in map `other` but not in map `m`. The returned `updatedKeys` are the keys that are both in map `m` and `other` but their values and not equal (`!=`).

##### (*StrStrMap) FilterEmpty 

``` go
func (m *StrStrMap) FilterEmpty()
```

FilterEmpty deletes all key-value pair of which the value is empty. Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.

##### Example

``` go
```
##### (*StrStrMap) Flip 

``` go
func (m *StrStrMap) Flip()
```

Flip exchanges key-value of the map to value-key.

##### Example

``` go
```
##### (*StrStrMap) Get 

``` go
func (m *StrStrMap) Get(key string) (value string)
```

Get returns the value by given `key`.

##### Example

``` go
```
##### (*StrStrMap) GetOrSet 

``` go
func (m *StrStrMap) GetOrSet(key string, value string) string
```

GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*StrStrMap) GetOrSetFunc 

``` go
func (m *StrStrMap) GetOrSetFunc(key string, f func() string) string
```

GetOrSetFunc returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

##### Example

``` go
```
##### (*StrStrMap) GetOrSetFuncLock 

``` go
func (m *StrStrMap) GetOrSetFuncLock(key string, f func() string) string
```

GetOrSetFuncLock returns the value by key, or sets value with returned value of callback function `f` if it does not exist and then returns this value.

GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*StrStrMap) IsEmpty 

``` go
func (m *StrStrMap) IsEmpty() bool
```

IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.

##### Example

``` go
```
##### (*StrStrMap) IsSubOf <-2.3.3

``` go
func (m *StrStrMap) IsSubOf(other *StrStrMap) bool
```

IsSubOf checks whether the current map is a sub-map of `other`.

##### (*StrStrMap) Iterator 

``` go
func (m *StrStrMap) Iterator(f func(k string, v string) bool)
```

Iterator iterates the hash map readonly with custom callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

##### Example

``` go
```
##### (*StrStrMap) Keys 

``` go
func (m *StrStrMap) Keys() []string
```

Keys returns all keys of the map as a slice.

##### Example

``` go
```
##### (*StrStrMap) LockFunc 

``` go
func (m *StrStrMap) LockFunc(f func(m map[string]string))
```

LockFunc locks writing with given callback function `f` within RWMutex.Lock.

##### Example

``` go
```
##### (*StrStrMap) Map 

``` go
func (m *StrStrMap) Map() map[string]string
```

Map returns the underlying data map. Note that, if it's in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

##### Example

``` go
```
##### (*StrStrMap) MapCopy 

``` go
func (m *StrStrMap) MapCopy() map[string]string
```

MapCopy returns a copy of the underlying data of the hash map.

##### Example

``` go
```
##### (*StrStrMap) MapStrAny 

``` go
func (m *StrStrMap) MapStrAny() map[string]interface{}
```

MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.

##### Example

``` go
```
##### (StrStrMap) MarshalJSON 

``` go
func (m StrStrMap) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*StrStrMap) Merge 

``` go
func (m *StrStrMap) Merge(other *StrStrMap)
```

Merge merges two hash maps. The `other` map will be merged into the map `m`.

##### Example

``` go
```
##### (*StrStrMap) Pop 

``` go
func (m *StrStrMap) Pop() (key, value string)
```

Pop retrieves and deletes an item from the map.

##### Example

``` go
```
##### (*StrStrMap) Pops 

``` go
func (m *StrStrMap) Pops(size int) map[string]string
```

Pops retrieves and deletes `size` items from the map. It returns all items if size == -1.

##### Example

``` go
```
##### (*StrStrMap) RLockFunc 

``` go
func (m *StrStrMap) RLockFunc(f func(m map[string]string))
```

RLockFunc locks reading with given callback function `f` within RWMutex.RLock.

##### Example

``` go
```
##### (*StrStrMap) Remove 

``` go
func (m *StrStrMap) Remove(key string) (value string)
```

Remove deletes value from map by given `key`, and return this deleted value.

##### Example

``` go
```
##### (*StrStrMap) Removes 

``` go
func (m *StrStrMap) Removes(keys []string)
```

Removes batch deletes values of the map by keys.

##### Example

``` go
```
##### (*StrStrMap) Replace 

``` go
func (m *StrStrMap) Replace(data map[string]string)
```

Replace the data of the map with given `data`.

##### Example

``` go
```
##### (*StrStrMap) Search 

``` go
func (m *StrStrMap) Search(key string) (value string, found bool)
```

Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.

##### Example

``` go
```
##### (*StrStrMap) Set 

``` go
func (m *StrStrMap) Set(key string, val string)
```

Set sets key-value to the hash map.

##### Example

``` go
```
##### (*StrStrMap) SetIfNotExist 

``` go
func (m *StrStrMap) SetIfNotExist(key string, value string) bool
```

SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*StrStrMap) SetIfNotExistFunc 

``` go
func (m *StrStrMap) SetIfNotExistFunc(key string, f func() string) bool
```

SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

##### Example

``` go
```
##### (*StrStrMap) SetIfNotExistFuncLock 

``` go
func (m *StrStrMap) SetIfNotExistFuncLock(key string, f func() string) bool
```

SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true. It returns false if `key` exists, and `value` would be ignored.

SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that it executes function `f` with mutex.Lock of the hash map.

##### Example

``` go
```
##### (*StrStrMap) Sets 

``` go
func (m *StrStrMap) Sets(data map[string]string)
```

Sets batch sets key-values to the hash map.

##### Example

``` go
```
##### (*StrStrMap) Size 

``` go
func (m *StrStrMap) Size() int
```

Size returns the size of the map.

##### Example

``` go
```
##### (*StrStrMap) String 

``` go
func (m *StrStrMap) String() string
```

String returns the map as a string.

##### Example

``` go
```
##### (*StrStrMap) UnmarshalJSON 

``` go
func (m *StrStrMap) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*StrStrMap) UnmarshalValue 

``` go
func (m *StrStrMap) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for map.

##### Example

``` go
```
##### (*StrStrMap) Values 

``` go
func (m *StrStrMap) Values() []string
```

Values returns all values of the map as a slice.

##### Example

``` go
```
#### type TreeMap 

``` go
type TreeMap = gtree.RedBlackTree
```

TreeMap based on red-black tree, alias of RedBlackTree.

##### func NewTreeMap 

``` go
func NewTreeMap(comparator func(v1, v2 interface{}) int, safe ...bool) *TreeMap
```

NewTreeMap instantiates a tree map with the custom comparator. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewTreeMapFrom 

``` go
func NewTreeMapFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *TreeMap
```

NewTreeMapFrom instantiates a tree map with the custom comparator and `data` map. Note that, the param `data` map will be set as the underlying data map(no deep copy), there might be some concurrent-safe issues when changing the map outside. The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

##### Example

``` go
```








<iframe allowtransparency="true" frameborder="0" scrolling="no" class="sk_ui" src="chrome-extension://gfbliohnnapiefjpjlpjnehglfpaknnc/pages/frontend.html" title="Surfingkeys" style="left: 0px; bottom: 0px; width: 1555px; height: 0px; z-index: 2147483647;"></iframe>