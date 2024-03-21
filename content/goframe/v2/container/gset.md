+++
title = "gset"
date = 2024-03-21T17:45:04+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gset

Package gset provides kinds of concurrent-safe/unsafe sets.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type IntSet 

``` go
type IntSet struct {
	// contains filtered or unexported fields
}
```

##### func NewIntSet 

``` go
func NewIntSet(safe ...bool) *IntSet
```

NewIntSet create and returns a new set, which contains un-repeated items. The parameter `safe` is used to specify whether using set in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewIntSetFrom 

``` go
func NewIntSetFrom(items []int, safe ...bool) *IntSet
```

NewIntSetFrom returns a new set from `items`.

##### (*IntSet) Add 

``` go
func (set *IntSet) Add(item ...int)
```

Add adds one or multiple items to the set.

##### Example

``` go
```
##### (*IntSet) AddIfNotExist 

``` go
func (set *IntSet) AddIfNotExist(item int) bool
```

AddIfNotExist checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set, or else it does nothing and returns false.

Note that, if `item` is nil, it does nothing and returns false.

##### Example

``` go
```
##### (*IntSet) AddIfNotExistFunc 

``` go
func (set *IntSet) AddIfNotExistFunc(item int, f func() bool) bool
```

AddIfNotExistFunc checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set and function `f` returns true, or else it does nothing and returns false.

Note that, the function `f` is executed without writing lock.

##### Example

``` go
```
##### (*IntSet) AddIfNotExistFuncLock 

``` go
func (set *IntSet) AddIfNotExistFuncLock(item int, f func() bool) bool
```

AddIfNotExistFuncLock checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set and function `f` returns true, or else it does nothing and returns false.

Note that, the function `f` is executed without writing lock.

##### Example

``` go
```
##### (*IntSet) Clear 

``` go
func (set *IntSet) Clear()
```

Clear deletes all items of the set.

##### Example

``` go
```
##### (*IntSet) Complement 

``` go
func (set *IntSet) Complement(full *IntSet) (newSet *IntSet)
```

Complement returns a new set which is the complement from `set` to `full`. Which means, all the items in `newSet` are in `full` and not in `set`.

It returns the difference between `full` and `set` if the given set `full` is not the full set of `set`.

##### Example

``` go
```
##### (*IntSet) Contains 

``` go
func (set *IntSet) Contains(item int) bool
```

Contains checks whether the set contains `item`.

##### Example

``` go
```
##### (*IntSet) DeepCopy <-2.1.0

``` go
func (set *IntSet) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*IntSet) Diff 

``` go
func (set *IntSet) Diff(others ...*IntSet) (newSet *IntSet)
```

Diff returns a new set which is the difference set from `set` to `other`. Which means, all the items in `newSet` are in `set` but not in `other`.

##### Example

``` go
```
##### (*IntSet) Equal 

``` go
func (set *IntSet) Equal(other *IntSet) bool
```

Equal checks whether the two sets equal.

##### Example

``` go
```
##### (*IntSet) Intersect 

``` go
func (set *IntSet) Intersect(others ...*IntSet) (newSet *IntSet)
```

Intersect returns a new set which is the intersection from `set` to `other`. Which means, all the items in `newSet` are in `set` and also in `other`.

##### Example

``` go
```
##### (*IntSet) IsSubsetOf 

``` go
func (set *IntSet) IsSubsetOf(other *IntSet) bool
```

IsSubsetOf checks whether the current set is a sub-set of `other`.

##### Example

``` go
```
##### (*IntSet) Iterator 

``` go
func (set *IntSet) Iterator(f func(v int) bool)
```

Iterator iterates the set readonly with given callback function `f`, if `f` returns true then continue iterating; or false to stop.

##### Example

``` go
```
##### (*IntSet) Join 

``` go
func (set *IntSet) Join(glue string) string
```

Join joins items with a string `glue`.

##### Example

``` go
```
##### (*IntSet) LockFunc 

``` go
func (set *IntSet) LockFunc(f func(m map[int]struct{}))
```

LockFunc locks writing with callback function `f`.

##### Example

``` go
```
##### (IntSet) MarshalJSON 

``` go
func (set IntSet) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*IntSet) Merge 

``` go
func (set *IntSet) Merge(others ...*IntSet) *IntSet
```

Merge adds items from `others` sets into `set`.

##### Example

``` go
```
##### (*IntSet) Pop 

``` go
func (set *IntSet) Pop() int
```

Pop randomly pops an item from set.

##### Example

``` go
```
##### (*IntSet) Pops 

``` go
func (set *IntSet) Pops(size int) []int
```

Pops randomly pops `size` items from set. It returns all items if size == -1.

##### Example

``` go
```
##### (*IntSet) RLockFunc 

``` go
func (set *IntSet) RLockFunc(f func(m map[int]struct{}))
```

RLockFunc locks reading with callback function `f`.

##### Example

``` go
```
##### (*IntSet) Remove 

``` go
func (set *IntSet) Remove(item int)
```

Remove deletes `item` from set.

##### Example

``` go
```
##### (*IntSet) Size 

``` go
func (set *IntSet) Size() int
```

Size returns the size of the set.

##### Example

``` go
```
##### (*IntSet) Slice 

``` go
func (set *IntSet) Slice() []int
```

Slice returns the an of items of the set as slice.

##### Example

``` go
```
##### (*IntSet) String 

``` go
func (set *IntSet) String() string
```

String returns items as a string, which implements like json.Marshal does.

##### Example

``` go
```
##### (*IntSet) Sum 

``` go
func (set *IntSet) Sum() (sum int)
```

Sum sums items. Note: The items should be converted to int type, or you'd get a result that you unexpected.

##### Example

``` go
```
##### (*IntSet) Union 

``` go
func (set *IntSet) Union(others ...*IntSet) (newSet *IntSet)
```

Union returns a new set which is the union of `set` and `other`. Which means, all the items in `newSet` are in `set` or in `other`.

##### Example

``` go
```
##### (*IntSet) UnmarshalJSON 

``` go
func (set *IntSet) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*IntSet) UnmarshalValue 

``` go
func (set *IntSet) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for set.

##### Example

``` go
```
##### (*IntSet) Walk 

``` go
func (set *IntSet) Walk(f func(item int) int) *IntSet
```

Walk applies a user supplied function `f` to every item of set.

##### Example

``` go
```
#### type Set 

``` go
type Set struct {
	// contains filtered or unexported fields
}
```

##### func New 

``` go
func New(safe ...bool) *Set
```

New create and returns a new set, which contains un-repeated items. The parameter `safe` is used to specify whether using set in concurrent-safety, which is false in default.

##### func NewFrom 

``` go
func NewFrom(items interface{}, safe ...bool) *Set
```

NewFrom returns a new set from `items`. Parameter `items` can be either a variable of any type, or a slice.

##### Example

``` go
```
##### func NewSet 

``` go
func NewSet(safe ...bool) *Set
```

NewSet create and returns a new set, which contains un-repeated items. Also see New.

##### (*Set) Add 

``` go
func (set *Set) Add(items ...interface{})
```

Add adds one or multiple items to the set.

##### (*Set) AddIfNotExist 

``` go
func (set *Set) AddIfNotExist(item interface{}) bool
```

AddIfNotExist checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set, or else it does nothing and returns false.

Note that, if `item` is nil, it does nothing and returns false.

##### Example

``` go
```
##### (*Set) AddIfNotExistFunc 

``` go
func (set *Set) AddIfNotExistFunc(item interface{}, f func() bool) bool
```

AddIfNotExistFunc checks whether item exists in the set, it adds the item to set and returns true if it does not exist in the set and function `f` returns true, or else it does nothing and returns false.

Note that, if `item` is nil, it does nothing and returns false. The function `f` is executed without writing lock.

##### (*Set) AddIfNotExistFuncLock 

``` go
func (set *Set) AddIfNotExistFuncLock(item interface{}, f func() bool) bool
```

AddIfNotExistFuncLock checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set and function `f` returns true, or else it does nothing and returns false.

Note that, if `item` is nil, it does nothing and returns false. The function `f` is executed within writing lock.

##### (*Set) Clear 

``` go
func (set *Set) Clear()
```

Clear deletes all items of the set.

##### (*Set) Complement 

``` go
func (set *Set) Complement(full *Set) (newSet *Set)
```

Complement returns a new set which is the complement from `set` to `full`. Which means, all the items in `newSet` are in `full` and not in `set`.

It returns the difference between `full` and `set` if the given set `full` is not the full set of `set`.

##### Example

``` go
```
##### (*Set) Contains 

``` go
func (set *Set) Contains(item interface{}) bool
```

Contains checks whether the set contains `item`.

##### Example

``` go
```
##### (*Set) DeepCopy <-2.1.0

``` go
func (set *Set) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*Set) Diff 

``` go
func (set *Set) Diff(others ...*Set) (newSet *Set)
```

Diff returns a new set which is the difference set from `set` to `others`. Which means, all the items in `newSet` are in `set` but not in `others`.

##### Example

``` go
```
##### (*Set) Equal 

``` go
func (set *Set) Equal(other *Set) bool
```

Equal checks whether the two sets equal.

##### (*Set) Intersect 

``` go
func (set *Set) Intersect(others ...*Set) (newSet *Set)
```

Intersect returns a new set which is the intersection from `set` to `others`. Which means, all the items in `newSet` are in `set` and also in `others`.

##### Example

``` go
```
##### (*Set) IsSubsetOf 

``` go
func (set *Set) IsSubsetOf(other *Set) bool
```

IsSubsetOf checks whether the current set is a sub-set of `other`.

##### Example

``` go
```
##### (*Set) Iterator 

``` go
func (set *Set) Iterator(f func(v interface{}) bool)
```

Iterator iterates the set readonly with given callback function `f`, if `f` returns true then continue iterating; or false to stop.

##### (*Set) Join 

``` go
func (set *Set) Join(glue string) string
```

Join joins items with a string `glue`.

##### Example

``` go
```
##### (*Set) LockFunc 

``` go
func (set *Set) LockFunc(f func(m map[interface{}]struct{}))
```

LockFunc locks writing with callback function `f`.

##### (Set) MarshalJSON 

``` go
func (set Set) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Set) Merge 

``` go
func (set *Set) Merge(others ...*Set) *Set
```

Merge adds items from `others` sets into `set`.

##### (*Set) Pop 

``` go
func (set *Set) Pop() interface{}
```

Pop randomly pops an item from set.

##### Example

``` go
```
##### (*Set) Pops 

``` go
func (set *Set) Pops(size int) []interface{}
```

Pops randomly pops `size` items from set. It returns all items if size == -1.

##### Example

``` go
```
##### (*Set) RLockFunc 

``` go
func (set *Set) RLockFunc(f func(m map[interface{}]struct{}))
```

RLockFunc locks reading with callback function `f`.

##### (*Set) Remove 

``` go
func (set *Set) Remove(item interface{})
```

Remove deletes `item` from set.

##### (*Set) Size 

``` go
func (set *Set) Size() int
```

Size returns the size of the set.

##### (*Set) Slice 

``` go
func (set *Set) Slice() []interface{}
```

Slice returns all items of the set as slice.

##### (*Set) String 

``` go
func (set *Set) String() string
```

String returns items as a string, which implements like json.Marshal does.

##### (*Set) Sum 

``` go
func (set *Set) Sum() (sum int)
```

Sum sums items. Note: The items should be converted to int type, or you'd get a result that you unexpected.

##### (*Set) Union 

``` go
func (set *Set) Union(others ...*Set) (newSet *Set)
```

Union returns a new set which is the union of `set` and `others`. Which means, all the items in `newSet` are in `set` or in `others`.

##### Example

``` go
```
##### (*Set) UnmarshalJSON 

``` go
func (set *Set) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Set) UnmarshalValue 

``` go
func (set *Set) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for set.

##### (*Set) Walk 

``` go
func (set *Set) Walk(f func(item interface{}) interface{}) *Set
```

Walk applies a user supplied function `f` to every item of set.

#### type StrSet 

``` go
type StrSet struct {
	// contains filtered or unexported fields
}
```

##### func NewStrSet 

``` go
func NewStrSet(safe ...bool) *StrSet
```

NewStrSet create and returns a new set, which contains un-repeated items. The parameter `safe` is used to specify whether using set in concurrent-safety, which is false in default.

##### Example

``` go
```
##### func NewStrSetFrom 

``` go
func NewStrSetFrom(items []string, safe ...bool) *StrSet
```

NewStrSetFrom returns a new set from `items`.

##### Example

``` go
```
##### (*StrSet) Add 

``` go
func (set *StrSet) Add(item ...string)
```

Add adds one or multiple items to the set.

##### Example

``` go
```
##### (*StrSet) AddIfNotExist 

``` go
func (set *StrSet) AddIfNotExist(item string) bool
```

AddIfNotExist checks whether item exists in the set, it adds the item to set and returns true if it does not exist in the set, or else it does nothing and returns false.

##### Example

``` go
```
##### (*StrSet) AddIfNotExistFunc 

``` go
func (set *StrSet) AddIfNotExistFunc(item string, f func() bool) bool
```

AddIfNotExistFunc checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set and function `f` returns true, or else it does nothing and returns false.

Note that, the function `f` is executed without writing lock.

##### Example

``` go
```
##### (*StrSet) AddIfNotExistFuncLock 

``` go
func (set *StrSet) AddIfNotExistFuncLock(item string, f func() bool) bool
```

AddIfNotExistFuncLock checks whether item exists in the set, it adds the item to set and returns true if it does not exists in the set and function `f` returns true, or else it does nothing and returns false.

Note that, the function `f` is executed without writing lock.

##### Example

``` go
```
##### (*StrSet) Clear 

``` go
func (set *StrSet) Clear()
```

Clear deletes all items of the set.

##### Example

``` go
```
##### (*StrSet) Complement 

``` go
func (set *StrSet) Complement(full *StrSet) (newSet *StrSet)
```

Complement returns a new set which is the complement from `set` to `full`. Which means, all the items in `newSet` are in `full` and not in `set`.

It returns the difference between `full` and `set` if the given set `full` is not the full set of `set`.

##### Example

``` go
```
##### (*StrSet) Contains 

``` go
func (set *StrSet) Contains(item string) bool
```

Contains checks whether the set contains `item`.

##### Example

``` go
```
##### (*StrSet) ContainsI 

``` go
func (set *StrSet) ContainsI(item string) bool
```

ContainsI checks whether a value exists in the set with case-insensitively. Note that it internally iterates the whole set to do the comparison with case-insensitively.

##### Example

``` go
```
##### (*StrSet) DeepCopy <-2.1.0

``` go
func (set *StrSet) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*StrSet) Diff 

``` go
func (set *StrSet) Diff(others ...*StrSet) (newSet *StrSet)
```

Diff returns a new set which is the difference set from `set` to `other`. Which means, all the items in `newSet` are in `set` but not in `other`.

##### Example

``` go
```
##### (*StrSet) Equal 

``` go
func (set *StrSet) Equal(other *StrSet) bool
```

Equal checks whether the two sets equal.

##### Example

``` go
```
##### (*StrSet) Intersect 

``` go
func (set *StrSet) Intersect(others ...*StrSet) (newSet *StrSet)
```

Intersect returns a new set which is the intersection from `set` to `other`. Which means, all the items in `newSet` are in `set` and also in `other`.

##### Example

``` go
```
##### (*StrSet) IsSubsetOf 

``` go
func (set *StrSet) IsSubsetOf(other *StrSet) bool
```

IsSubsetOf checks whether the current set is a sub-set of `other`.

##### Example

``` go
```
##### (*StrSet) Iterator 

``` go
func (set *StrSet) Iterator(f func(v string) bool)
```

Iterator iterates the set readonly with given callback function `f`, if `f` returns true then continue iterating; or false to stop.

##### Example

``` go
```
##### (*StrSet) Join 

``` go
func (set *StrSet) Join(glue string) string
```

Join joins items with a string `glue`.

##### Example

``` go
```
##### (*StrSet) LockFunc 

``` go
func (set *StrSet) LockFunc(f func(m map[string]struct{}))
```

LockFunc locks writing with callback function `f`.

##### Example

``` go
```
##### (StrSet) MarshalJSON 

``` go
func (set StrSet) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*StrSet) Merge 

``` go
func (set *StrSet) Merge(others ...*StrSet) *StrSet
```

Merge adds items from `others` sets into `set`.

##### Example

``` go
```
##### (*StrSet) Pop 

``` go
func (set *StrSet) Pop() string
```

Pop randomly pops an item from set.

##### Example

``` go
```
##### (*StrSet) Pops 

``` go
func (set *StrSet) Pops(size int) []string
```

Pops randomly pops `size` items from set. It returns all items if size == -1.

##### Example

``` go
```
##### (*StrSet) RLockFunc 

``` go
func (set *StrSet) RLockFunc(f func(m map[string]struct{}))
```

RLockFunc locks reading with callback function `f`.

##### Example

``` go
```
##### (*StrSet) Remove 

``` go
func (set *StrSet) Remove(item string)
```

Remove deletes `item` from set.

##### Example

``` go
```
##### (*StrSet) Size 

``` go
func (set *StrSet) Size() int
```

Size returns the size of the set.

##### Example

``` go
```
##### (*StrSet) Slice 

``` go
func (set *StrSet) Slice() []string
```

Slice returns the an of items of the set as slice.

##### Example

``` go
```
##### (*StrSet) String 

``` go
func (set *StrSet) String() string
```

String returns items as a string, which implements like json.Marshal does.

##### Example

``` go
```
##### (*StrSet) Sum 

``` go
func (set *StrSet) Sum() (sum int)
```

Sum sums items. Note: The items should be converted to int type, or you'd get a result that you unexpected.

##### Example

``` go
```
##### (*StrSet) Union 

``` go
func (set *StrSet) Union(others ...*StrSet) (newSet *StrSet)
```

Union returns a new set which is the union of `set` and `other`. Which means, all the items in `newSet` are in `set` or in `other`.

##### Example

``` go
```
##### (*StrSet) UnmarshalJSON 

``` go
func (set *StrSet) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*StrSet) UnmarshalValue 

``` go
func (set *StrSet) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for set.

##### Example

``` go
```
##### (*StrSet) Walk 

``` go
func (set *StrSet) Walk(f func(item string) string) *StrSet
```

Walk applies a user supplied function `f` to every item of set.

Example Walk

Walk applies a user supplied function `f` to every item of set.

``` go
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

