+++
title = "garray"
date = 2024-03-21T17:44:22+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/garray](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/garray)

Package garray provides most commonly used array containers which also support concurrent-safe/unsafe switch feature.

​	软件包 garray 提供最常用的阵列容器，这些容器还支持并发安全/不安全交换机功能。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Array

```go
type Array struct {
	// contains filtered or unexported fields
}
```

Array is a golang array with rich features. It contains a concurrent-safe/unsafe switch, which should be set when its initialization and cannot be changed then.

​	Array 是一个功能丰富的 golang 阵列。它包含一个并发安全/不安全开关，应在初始化时设置该开关，然后无法更改。

#### func New

```go
func New(safe ...bool) *Array
```

New creates and returns an empty array. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	new 创建并返回一个空数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example New

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
)

func main() {
	// A normal array.
	a := garray.New()

	// Adding items.
	for i := 0; i < 10; i++ {
		a.Append(i)
	}

	// Print the array length.
	fmt.Println(a.Len())

	// Print the array items.
	fmt.Println(a.Slice())

	// Retrieve item by index.
	fmt.Println(a.Get(6))

	// Check item existence.
	fmt.Println(a.Contains(6))
	fmt.Println(a.Contains(100))

	// Insert item before specified index.
	a.InsertAfter(9, 11)
	// Insert item after specified index.
	a.InsertBefore(10, 10)

	fmt.Println(a.Slice())

	// Modify item by index.
	a.Set(0, 100)
	fmt.Println(a.Slice())

	fmt.Println(a.At(0))

	// Search item and return its index.
	fmt.Println(a.Search(5))

	// Remove item by index.
	a.Remove(0)
	fmt.Println(a.Slice())

	// Empty the array, removes all items of it.
	fmt.Println(a.Slice())
	a.Clear()
	fmt.Println(a.Slice())

}
Output:

10
[0 1 2 3 4 5 6 7 8 9]
6 true
true
false
[0 1 2 3 4 5 6 7 8 9 10 11]
[100 1 2 3 4 5 6 7 8 9 10 11]
100
5
[1 2 3 4 5 6 7 8 9 10 11]
[1 2 3 4 5 6 7 8 9 10 11]
[]
```

#### func NewArray

```go
func NewArray(safe ...bool) *Array
```

NewArray is alias of New, please see New.

​	NewArray 是 New 的别名，请参阅 New。

#### func NewArrayFrom

```go
func NewArrayFrom(array []interface{}, safe ...bool) *Array
```

NewArrayFrom creates and returns an array with given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewArrayFrom 创建并返回一个具有给定切片 `array` 的数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### func NewArrayFromCopy

```go
func NewArrayFromCopy(array []interface{}, safe ...bool) *Array
```

NewArrayFromCopy creates and returns an array from a copy of given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewArrayFromCopy 从给定切片 `array` 的副本创建并返回一个数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### func NewArrayRange

```go
func NewArrayRange(start, end, step int, safe ...bool) *Array
```

NewArrayRange creates and returns an array by a range from `start` to `end` with step value `step`.

​	NewArrayRange 创建并返回一个范围 `start` 为 `end` to 的数组，其步骤值 `step` 为 。

#### func NewArraySize

```go
func NewArraySize(size int, cap int, safe ...bool) *Array
```

NewArraySize create and returns an array with given size and cap. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewArraySize 创建并返回一个具有给定大小和上限的数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### func NewFrom

```go
func NewFrom(array []interface{}, safe ...bool) *Array
```

NewFrom is alias of NewArrayFrom. See NewArrayFrom.

​	NewFrom 是 NewArrayFrom 的别名。请参阅 NewArrayFrom。

#### func NewFromCopy

```go
func NewFromCopy(array []interface{}, safe ...bool) *Array
```

NewFromCopy is alias of NewArrayFromCopy. See NewArrayFromCopy.

​	NewFromCopy 是 NewArrayFromCopy 的别名。请参阅 NewArrayFromCopy。

#### (*Array) Append

```go
func (a *Array) Append(value ...interface{}) *Array
```

Append is alias of PushRight, please See PushRight.

​	Append 是 PushRight 的别名，请参见 PushRight。

#### (*Array) At

```go
func (a *Array) At(index int) (value interface{})
```

At returns the value by the specified index. If the given `index` is out of range of the array, it returns `nil`.

​	at 返回指定索引的值。如果给定 `index` 的超出数组的范围，则返回 `nil` 。

#### (*Array) Chunk

```go
func (a *Array) Chunk(size int) [][]interface{}
```

Chunk splits an array into multiple arrays, the size of each array is determined by `size`. The last chunk may contain less than size elements.

​	Chunk 将一个数组拆分为多个数组，每个数组的大小由 `size` 决定。最后一个块可能包含小于 size 的元素。

##### Example Chunk

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	array := garray.NewFrom(g.Slice{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// Chunk splits an array into multiple arrays,
	// the size of each array is determined by `size`.
	// The last chunk may contain less than size elements.
	fmt.Println(array.Chunk(2))

}
Output:

[[1 2] [3 4] [5 6] [7 8] [9]]
```

#### (*Array) Clear

```go
func (a *Array) Clear() *Array
```

Clear deletes all items of current array.

​	清除将删除当前数组的所有项目。

#### (*Array) Clone

```go
func (a *Array) Clone() (newArray *Array)
```

Clone returns a new array, which is a copy of current array.

​	克隆返回一个新数组，它是当前数组的副本。

#### (*Array) Contains

```go
func (a *Array) Contains(value interface{}) bool
```

Contains checks whether a value exists in the array.

​	包含检查数组中是否存在值。

##### Example Contains

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
)

func main() {
	var array garray.StrArray
	array.Append("a")
	fmt.Println(array.Contains("a"))
	fmt.Println(array.Contains("A"))
	fmt.Println(array.ContainsI("A"))

}
Output:

true
false
true
```

#### (*Array) CountValues

```go
func (a *Array) CountValues() map[interface{}]int
```

CountValues counts the number of occurrences of all values in the array.

​	CountValues 计算数组中所有值的出现次数。

#### (*Array) DeepCopy

```go
func (a *Array) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*Array) Fill

```go
func (a *Array) Fill(startIndex int, num int, value interface{}) error
```

Fill fills an array with num entries of the value `value`, keys starting at the `startIndex` parameter.

​	Fill 使用值 `value` 的 num 个条目填充数组，键从 `startIndex` 参数开始。

#### (*Array) Filter

```go
func (a *Array) Filter(filter func(index int, value interface{}) bool) *Array
```

Filter iterates array and filters elements using custom callback function. It removes the element from array if callback function `filter` returns true, it or else does nothing and continues iterating.

​	Filter 使用自定义回调函数迭代数组和筛选元素。如果回调函数 `filter` 返回 true，它会从数组中删除该元素，否则它不执行任何操作并继续迭代。

##### Example Filter 

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/internal/empty"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	array1 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	array2 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	fmt.Printf("%#v\n", array1.Filter(func(index int, value interface{}) bool {
		return empty.IsNil(value)
	}).Slice())
	fmt.Printf("%#v\n", array2.Filter(func(index int, value interface{}) bool {
		return empty.IsEmpty(value)
	}).Slice())

}
Output:

[]interface {}{0, 1, 2, "", []interface {}{}, "john"}
[]interface {}{1, 2, "john"}
```

#### (*Array) FilterEmpty

```go
func (a *Array) FilterEmpty() *Array
```

FilterEmpty removes all empty value of the array. Values like: 0, nil, false, “”, len(slice/map/chan) == 0 are considered empty.

​	FilterEmpty 删除数组的所有空值。像 0， nil， false， “”， len（slice/map/chan） == 0 这样的值被视为空。

##### Example FilterEmpty

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	array1 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	array2 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	fmt.Printf("%#v\n", array1.FilterNil().Slice())
	fmt.Printf("%#v\n", array2.FilterEmpty().Slice())

}
Output:

[]interface {}{0, 1, 2, "", []interface {}{}, "john"}
[]interface {}{1, 2, "john"}
```

#### (*Array) FilterNil

```go
func (a *Array) FilterNil() *Array
```

FilterNil removes all nil value of the array.

​	FilterNil 删除数组的所有 nil 值。

##### Example FilterNil

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	array1 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	array2 := garray.NewFrom(g.Slice{0, 1, 2, nil, "", g.Slice{}, "john"})
	fmt.Printf("%#v\n", array1.FilterNil().Slice())
	fmt.Printf("%#v\n", array2.FilterEmpty().Slice())

}
Output:

[]interface {}{0, 1, 2, "", []interface {}{}, "john"}
[]interface {}{1, 2, "john"}
```

#### (*Array) Get

```go
func (a *Array) Get(index int) (value interface{}, found bool)
```

Get returns the value by the specified index. If the given `index` is out of range of the array, the `found` is false.

​	Get 按指定索引返回值。如果给定 `index` 的超出数组的范围，则为 `found` false。

#### (*Array) InsertAfter

```go
func (a *Array) InsertAfter(index int, values ...interface{}) error
```

InsertAfter inserts the `values` to the back of `index`.

​	InsertAfter 将 插入 `values` 到 `index` 的后面。

#### (*Array) InsertBefore

```go
func (a *Array) InsertBefore(index int, values ...interface{}) error
```

InsertBefore inserts the `values` to the front of `index`.

​	InsertBefore 将 插入 `values` 到 `index` 的前面。

#### (*Array) Interfaces

```go
func (a *Array) Interfaces() []interface{}
```

Interfaces returns current array as []interface{}.

​	Interfaces 以 []interface{} 的形式返回当前数组。

#### (*Array) IsEmpty

```go
func (a *Array) IsEmpty() bool
```

IsEmpty checks whether the array is empty.

​	IsEmpty 检查数组是否为空。

#### (*Array) Iterator

```go
func (a *Array) Iterator(f func(k int, v interface{}) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

##### Example Iterator

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	array := garray.NewArrayFrom(g.Slice{"a", "b", "c"})
	// Iterator is alias of IteratorAsc, which iterates the array readonly in ascending order
	//  with given callback function `f`.
	// If `f` returns true, then it continues iterating; or false to stop.
	array.Iterator(func(k int, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	// IteratorDesc iterates the array readonly in descending order with given callback function `f`.
	// If `f` returns true, then it continues iterating; or false to stop.
	array.IteratorDesc(func(k int, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

}
Output:

0 a
1 b
2 c
2 c
1 b
0 a
```

#### (*Array) IteratorAsc

```go
func (a *Array) IteratorAsc(f func(k int, v interface{}) bool)
```

IteratorAsc iterates the array readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序只读方式迭代数组。如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*Array) IteratorDesc

```go
func (a *Array) IteratorDesc(f func(k int, v interface{}) bool)
```

IteratorDesc iterates the array readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序迭代数组只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*Array) Join

```go
func (a *Array) Join(glue string) string
```

Join joins array elements with a string `glue`.

​	Join 使用字符串 `glue` 连接数组元素。

##### Example Join

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	array := garray.NewFrom(g.Slice{"a", "b", "c", "d"})
	fmt.Println(array.Join(","))

}
Output:

a,b,c,d
```

#### (*Array) Len

```go
func (a *Array) Len() int
```

Len returns the length of array.

​	Len 返回数组的长度。

#### (*Array) LockFunc

```go
func (a *Array) LockFunc(f func(array []interface{})) *Array
```

LockFunc locks writing by callback function `f`.

​	LockFunc 通过回调函数 `f` 锁定写入。

#### (Array) MarshalJSON

```go
func (a Array) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. Note that do not use pointer as its receiver here.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。请注意，此处不要使用指针作为其接收器。

#### (*Array) Merge

```go
func (a *Array) Merge(array interface{}) *Array
```

Merge merges `array` into current array. The parameter `array` can be any garray or slice type. The difference between Merge and Append is Append supports only specified slice type, but Merge supports more parameter types.

​	合并 `array` 合并到当前数组中。该参数 `array` 可以是任何 garray 或 slice 类型。Merge 和 Append 的区别在于 Append 仅支持指定的切片类型，但 Merge 支持更多参数类型。

##### Example Merge

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	array1 := garray.NewFrom(g.Slice{1, 2})
	array2 := garray.NewFrom(g.Slice{3, 4})
	slice1 := g.Slice{5, 6}
	slice2 := []int{7, 8}
	slice3 := []string{"9", "0"}
	fmt.Println(array1.Slice())
	array1.Merge(array1)
	array1.Merge(array2)
	array1.Merge(slice1)
	array1.Merge(slice2)
	array1.Merge(slice3)
	fmt.Println(array1.Slice())

}
Output:

[1 2]
[1 2 1 2 3 4 5 6 7 8 9 0]
```

#### (*Array) Pad

```go
func (a *Array) Pad(size int, val interface{}) *Array
```

Pad pads array to the specified length with `value`. If size is positive then the array is padded on the right, or negative on the left. If the absolute value of `size` is less than or equal to the length of the array then no padding takes place.

​	焊盘垫数组设置为指定 `value` 长度。如果大小为正数，则数组在右侧填充，或在左侧填充为负数。如果 的 `size` 绝对值小于或等于数组的长度，则不会进行填充。

#### (*Array) PopLeft

```go
func (a *Array) PopLeft() (value interface{}, found bool)
```

PopLeft pops and returns an item from the beginning of array. Note that if the array is empty, the `found` is false.

​	PopLeft 弹出并返回数组开头的项。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*Array) PopLefts

```go
func (a *Array) PopLefts(size int) []interface{}
```

PopLefts pops and returns `size` items from the beginning of array.

​	PopLefts 弹出并返回 `size` 数组开头的项。

##### Example

``` go
```

#### (*Array) PopRand

```go
func (a *Array) PopRand() (value interface{}, found bool)
```

PopRand randomly pops and return an item out of array. Note that if the array is empty, the `found` is false.

​	PopRand 随机弹出并返回数组外的项目。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*Array) PopRands

```go
func (a *Array) PopRands(size int) []interface{}
```

PopRands randomly pops and returns `size` items out of array.

​	PopRands 随机弹出并返回 `size` 数组外的项目。

#### (*Array) PopRight

```go
func (a *Array) PopRight() (value interface{}, found bool)
```

PopRight pops and returns an item from the end of array. Note that if the array is empty, the `found` is false.

​	PopRight 弹出并返回数组末尾的项。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*Array) PopRights

```go
func (a *Array) PopRights(size int) []interface{}
```

PopRights pops and returns `size` items from the end of array.

​	PopRights 弹出并返回 `size` 数组末尾的项。

##### Example

``` go
```

#### (*Array) PushLeft

```go
func (a *Array) PushLeft(value ...interface{}) *Array
```

PushLeft pushes one or multiple items to the beginning of array.

​	PushLeft 将一个或多个项推送到数组的开头。

#### (*Array) PushRight

```go
func (a *Array) PushRight(value ...interface{}) *Array
```

PushRight pushes one or multiple items to the end of array. It equals to Append.

​	PushRight 将一个或多个项推送到数组的末尾。它等于追加。

#### (*Array) RLockFunc

```go
func (a *Array) RLockFunc(f func(array []interface{})) *Array
```

RLockFunc locks reading by callback function `f`.

​	RLockFunc 通过回调函数 `f` 锁定读取。

#### (*Array) Rand

```go
func (a *Array) Rand() (value interface{}, found bool)
```

Rand randomly returns one item from array(no deleting).

​	Rand 从数组中随机返回一个项目（不删除）。

#### (*Array) Rands

```go
func (a *Array) Rands(size int) []interface{}
```

Rands randomly returns `size` items from array(no deleting).

​	Rands 从数组中随机返回 `size` 项目（不删除）。

##### Example

``` go
```

#### (*Array) Range

```go
func (a *Array) Range(start int, end ...int) []interface{}
```

Range picks and returns items by range, like array[start:end]. Notice, if in concurrent-safe usage, it returns a copy of slice; else a pointer to the underlying data.

​	范围按范围选取和返回项目，如 array[start：end]。请注意，如果在并发安全使用中，它会返回 slice 的副本;else 指向基础数据的指针。

If `end` is negative, then the offset will start from the end of array. If `end` is omitted, then the sequence will have everything from start up until the end of the array.

​	如果 `end` 为负数，则偏移量将从数组末尾开始。如果 `end` 省略，则序列将包含从数组开始到结束的所有内容。

#### (*Array) Remove

```go
func (a *Array) Remove(index int) (value interface{}, found bool)
```

Remove removes an item by index. If the given `index` is out of range of the array, the `found` is false.

​	“删除”（Remove） 按索引删除项目。如果给定 `index` 的超出数组的范围，则为 `found` false。

#### (*Array) RemoveValue

```go
func (a *Array) RemoveValue(value interface{}) bool
```

RemoveValue removes an item by value. It returns true if value is found in the array, or else false if not found.

​	RemoveValue 按值删除项目。如果在数组中找到值，则返回 true，如果未找到，则返回 false。

#### (*Array) RemoveValues

```go
func (a *Array) RemoveValues(values ...interface{})
```

RemoveValues removes multiple items by `values`.

​	RemoveValues 按 `values` 删除多个项目。

#### (*Array) Replace

```go
func (a *Array) Replace(array []interface{}) *Array
```

Replace replaces the array items by given `array` from the beginning of array.

​	Replace 将数组项替换为从数组开头开始给定 `array` 的数组项。

#### (*Array) Reverse

```go
func (a *Array) Reverse() *Array
```

Reverse makes array with elements in reverse order.

​	Reverse 使数组具有相反顺序的元素。

##### Example

``` go
```

#### (*Array) Search

```go
func (a *Array) Search(value interface{}) int
```

Search searches array by `value`, returns the index of `value`, or returns -1 if not exists.

​	搜索 搜索数 `value` 组 by 返回 的 `value` 索引，如果不存在，则返回 -1。

#### (*Array) Set

```go
func (a *Array) Set(index int, value interface{}) error
```

Set sets value to specified index.

​	将设置值设置为指定索引。

#### (*Array) SetArray

```go
func (a *Array) SetArray(array []interface{}) *Array
```

SetArray sets the underlying slice array with the given `array`.

​	SetArray 使用给定 `array` 的 .

#### (*Array) Shuffle

```go
func (a *Array) Shuffle() *Array
```

Shuffle randomly shuffles the array.

​	随机洗牌数组。

##### Example

``` go
```

#### (*Array) Slice

```go
func (a *Array) Slice() []interface{}
```

Slice returns the underlying data of array. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Slice 返回数组的基础数据。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

#### (*Array) SortFunc

```go
func (a *Array) SortFunc(less func(v1, v2 interface{}) bool) *Array
```

SortFunc sorts the array by custom function `less`.

​	SortFunc 按自定义函数 `less` 对数组进行排序。

#### (*Array) String

```go
func (a *Array) String() string
```

String returns current array as a string, which implements like json.Marshal does.

​	String 以字符串形式返回当前数组，其实现方式类似于 json。元帅做到了。

#### (*Array) SubSlice

```go
func (a *Array) SubSlice(offset int, length ...int) []interface{}
```

SubSlice returns a slice of elements from the array as specified by the `offset` and `size` parameters. If in concurrent safe usage, it returns a copy of the slice; else a pointer.

​	SubSlice 从 `offset` 数组中返回由 and `size` 参数指定的元素切片。如果处于并发安全使用状态，则返回切片的副本;else 指针。

If offset is non-negative, the sequence will start at that offset in the array. If offset is negative, the sequence will start that far from the end of the array.

​	如果偏移量为非负数，则序列将从数组中的该偏移量开始。如果偏移量为负数，则序列将从数组末尾开始。

If length is given and is positive, then the sequence will have up to that many elements in it. If the array is shorter than the length, then only the available array elements will be present. If length is given and is negative then the sequence will stop that many elements from the end of the array. If it is omitted, then the sequence will have everything from offset up until the end of the array.

​	如果给定长度并且是正数，则序列中将包含多达那么多元素。如果数组短于长度，则仅存在可用的数组元素。如果给定长度并且是负数，则序列将从数组末尾停止该许多元素。如果省略它，则序列将包含从偏移量到数组末尾的所有内容。

Any possibility crossing the left border of array, it will fail.

​	任何越过数组左边界的可能性，都会失败。

#### (*Array) Sum

```go
func (a *Array) Sum() (sum int)
```

Sum returns the sum of values in an array.

​	Sum 返回数组中值的总和。

#### (*Array) Unique

```go
func (a *Array) Unique() *Array
```

Unique uniques the array, clear repeated items. Example: [1,1,2,3,2] -> [1,2,3]

​	唯一唯一数组，清除重复项。示例：[1,1,2,3,2] -> [1,2,3]

#### (*Array) UnmarshalJSON

```go
func (a *Array) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Array) UnmarshalValue

```go
func (a *Array) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for array.

​	UnmarshalValue 是一个接口实现，用于为数组设置任何类型的值。

#### (*Array) Walk

```go
func (a *Array) Walk(f func(value interface{}) interface{}) *Array
```

Walk applies a user supplied function `f` to every item of array.

​	Walk 将用户提供的函数 `f` 应用于数组的每个项。

### type IntArray

```go
type IntArray struct {
	// contains filtered or unexported fields
}
```

IntArray is a golang int array with rich features. It contains a concurrent-safe/unsafe switch, which should be set when its initialization and cannot be changed then.

​	IntArray 是一个功能丰富的 golang int 数组。它包含一个并发安全/不安全开关，应在初始化时设置该开关，然后无法更改。

#### func NewIntArray

```go
func NewIntArray(safe ...bool) *IntArray
```

NewIntArray creates and returns an empty array. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewIntArray 创建并返回一个空数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example

``` go
```

#### func NewIntArrayFrom

```go
func NewIntArrayFrom(array []int, safe ...bool) *IntArray
```

NewIntArrayFrom creates and returns an array with given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewIntArrayFrom 创建并返回一个具有给定切片 `array` 的数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example

``` go
```

#### func NewIntArrayFromCopy

```go
func NewIntArrayFromCopy(array []int, safe ...bool) *IntArray
```

NewIntArrayFromCopy creates and returns an array from a copy of given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewIntArrayFromCopy 从给定切片 `array` 的副本创建并返回数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example

``` go
```

#### func NewIntArrayRange

```go
func NewIntArrayRange(start, end, step int, safe ...bool) *IntArray
```

NewIntArrayRange creates and returns an array by a range from `start` to `end` with step value `step`.

​	NewIntArrayRange 创建并返回一个范围 `start` 为 `end` to 的数组，其步骤值 `step` 为 。

##### Example

``` go
```

#### func NewIntArraySize

```go
func NewIntArraySize(size int, cap int, safe ...bool) *IntArray
```

NewIntArraySize create and returns an array with given size and cap. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewIntArraySize 创建并返回具有给定大小和上限的数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example

``` go
```

#### (*IntArray) Append

```go
func (a *IntArray) Append(value ...int) *IntArray
```

Append is alias of PushRight,please See PushRight.

​	Append 是 PushRight 的别名，请参见 PushRight。

##### Example

``` go
```

#### (*IntArray) At

```go
func (a *IntArray) At(index int) (value int)
```

At returns the value by the specified index. If the given `index` is out of range of the array, it returns `0`.

​	at 返回指定索引的值。如果给定 `index` 的超出数组的范围，则返回 `0` 。

##### Example

``` go
```

#### (*IntArray) Chunk

```go
func (a *IntArray) Chunk(size int) [][]int
```

Chunk splits an array into multiple arrays, the size of each array is determined by `size`. The last chunk may contain less than size elements.

​	Chunk 将一个数组拆分为多个数组，每个数组的大小由 `size` 决定。最后一个块可能包含小于 size 的元素。

##### Example

``` go
```

#### (*IntArray) Clear

```go
func (a *IntArray) Clear() *IntArray
```

Clear deletes all items of current array.

​	清除将删除当前数组的所有项目。

##### Example

``` go
```

#### (*IntArray) Clone

```go
func (a *IntArray) Clone() (newArray *IntArray)
```

Clone returns a new array, which is a copy of current array.

​	克隆返回一个新数组，它是当前数组的副本。

##### Example

``` go
```

#### (*IntArray) Contains

```go
func (a *IntArray) Contains(value int) bool
```

Contains checks whether a value exists in the array.

​	包含检查数组中是否存在值。

##### Example

``` go
```

#### (*IntArray) CountValues

```go
func (a *IntArray) CountValues() map[int]int
```

CountValues counts the number of occurrences of all values in the array.

​	CountValues 计算数组中所有值的出现次数。

##### Example

``` go
```

#### (*IntArray) DeepCopy

```go
func (a *IntArray) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*IntArray) Fill

```go
func (a *IntArray) Fill(startIndex int, num int, value int) error
```

Fill fills an array with num entries of the value `value`, keys starting at the `startIndex` parameter.

​	Fill 使用值 `value` 的 num 个条目填充数组，键从 `startIndex` 参数开始。

##### Example

``` go
```

#### (*IntArray) Filter

```go
func (a *IntArray) Filter(filter func(index int, value int) bool) *IntArray
```

Filter iterates array and filters elements using custom callback function. It removes the element from array if callback function `filter` returns true, it or else does nothing and continues iterating.

​	Filter 使用自定义回调函数迭代数组和筛选元素。如果回调函数 `filter` 返回 true，它会从数组中删除该元素，否则它不执行任何操作并继续迭代。

##### Example

``` go
```

#### (*IntArray) FilterEmpty

```go
func (a *IntArray) FilterEmpty() *IntArray
```

FilterEmpty removes all zero value of the array.

​	FilterEmpty 删除数组的所有零值。

##### Example

``` go
```

#### (*IntArray) Get

```go
func (a *IntArray) Get(index int) (value int, found bool)
```

Get returns the value by the specified index. If the given `index` is out of range of the array, the `found` is false.

​	Get 按指定索引返回值。如果给定 `index` 的超出数组的范围，则为 `found` false。

##### Example

``` go
```

#### (*IntArray) InsertAfter

```go
func (a *IntArray) InsertAfter(index int, values ...int) error
```

InsertAfter inserts the `value` to the back of `index`.

​	InsertAfter 将 插入 `value` 到 `index` 的后面。

##### Example

``` go
```

#### (*IntArray) InsertBefore

```go
func (a *IntArray) InsertBefore(index int, values ...int) error
```

InsertBefore inserts the `values` to the front of `index`.

​	InsertBefore 将 插入 `values` 到 `index` 的前面。

##### Example

``` go
```

#### (*IntArray) Interfaces

```go
func (a *IntArray) Interfaces() []interface{}
```

Interfaces returns current array as []interface{}.

​	Interfaces 以 []interface{} 的形式返回当前数组。

##### Example

``` go
```

#### (*IntArray) IsEmpty

```go
func (a *IntArray) IsEmpty() bool
```

IsEmpty checks whether the array is empty.

​	IsEmpty 检查数组是否为空。

##### Example

``` go
```

#### (*IntArray) Iterator

```go
func (a *IntArray) Iterator(f func(k int, v int) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

##### Example

``` go
```

#### (*IntArray) IteratorAsc

```go
func (a *IntArray) IteratorAsc(f func(k int, v int) bool)
```

IteratorAsc iterates the array readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序只读方式迭代数组。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*IntArray) IteratorDesc

```go
func (a *IntArray) IteratorDesc(f func(k int, v int) bool)
```

IteratorDesc iterates the array readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序迭代数组只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*IntArray) Join

```go
func (a *IntArray) Join(glue string) string
```

Join joins array elements with a string `glue`.

​	Join 使用字符串 `glue` 连接数组元素。

##### Example

``` go
```

#### (*IntArray) Len

```go
func (a *IntArray) Len() int
```

Len returns the length of array.

​	Len 返回数组的长度。

##### Example

``` go
```

#### (*IntArray) LockFunc

```go
func (a *IntArray) LockFunc(f func(array []int)) *IntArray
```

LockFunc locks writing by callback function `f`.

​	LockFunc 通过回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (IntArray) MarshalJSON

```go
func (a IntArray) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. Note that do not use pointer as its receiver here.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。请注意，此处不要使用指针作为其接收器。

##### Example

``` go
```

#### (*IntArray) Merge

```go
func (a *IntArray) Merge(array interface{}) *IntArray
```

Merge merges `array` into current array. The parameter `array` can be any garray or slice type. The difference between Merge and Append is Append supports only specified slice type, but Merge supports more parameter types.

​	合并 `array` 合并到当前数组中。该参数 `array` 可以是任何 garray 或 slice 类型。Merge 和 Append 的区别在于 Append 仅支持指定的切片类型，但 Merge 支持更多参数类型。

##### Example

``` go
```

#### (*IntArray) Pad

```go
func (a *IntArray) Pad(size int, value int) *IntArray
```

Pad pads array to the specified length with `value`. If size is positive then the array is padded on the right, or negative on the left. If the absolute value of `size` is less than or equal to the length of the array then no padding takes place.

​	焊盘垫数组设置为指定 `value` 长度。如果大小为正数，则数组在右侧填充，或在左侧填充为负数。如果 的 `size` 绝对值小于或等于数组的长度，则不会进行填充。

##### Example

``` go
```

#### (*IntArray) PopLeft

```go
func (a *IntArray) PopLeft() (value int, found bool)
```

PopLeft pops and returns an item from the beginning of array. Note that if the array is empty, the `found` is false.

​	PopLeft 弹出并返回数组开头的项。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*IntArray) PopLefts

```go
func (a *IntArray) PopLefts(size int) []int
```

PopLefts pops and returns `size` items from the beginning of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopLefts 弹出并返回 `size` 数组开头的项。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

##### Example

``` go
```

#### (*IntArray) PopRand

```go
func (a *IntArray) PopRand() (value int, found bool)
```

PopRand randomly pops and return an item out of array. Note that if the array is empty, the `found` is false.

​	PopRand 随机弹出并返回数组外的项目。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*IntArray) PopRands

```go
func (a *IntArray) PopRands(size int) []int
```

PopRands randomly pops and returns `size` items out of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopRands 随机弹出并返回 `size` 数组外的项目。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

##### Example

``` go
```

#### (*IntArray) PopRight

```go
func (a *IntArray) PopRight() (value int, found bool)
```

PopRight pops and returns an item from the end of array. Note that if the array is empty, the `found` is false.

​	PopRight 弹出并返回数组末尾的项。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*IntArray) PopRights

```go
func (a *IntArray) PopRights(size int) []int
```

PopRights pops and returns `size` items from the end of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopRights 弹出并返回 `size` 数组末尾的项。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

##### Example

``` go
```

#### (*IntArray) PushLeft

```go
func (a *IntArray) PushLeft(value ...int) *IntArray
```

PushLeft pushes one or multiple items to the beginning of array.

​	PushLeft 将一个或多个项推送到数组的开头。

##### Example

``` go
```

#### (*IntArray) PushRight

```go
func (a *IntArray) PushRight(value ...int) *IntArray
```

PushRight pushes one or multiple items to the end of array. It equals to Append.

​	PushRight 将一个或多个项推送到数组的末尾。它等于追加。

##### Example

``` go
```

#### (*IntArray) RLockFunc

```go
func (a *IntArray) RLockFunc(f func(array []int)) *IntArray
```

RLockFunc locks reading by callback function `f`.

​	RLockFunc 通过回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*IntArray) Rand

```go
func (a *IntArray) Rand() (value int, found bool)
```

Rand randomly returns one item from array(no deleting).

​	Rand 从数组中随机返回一个项目（不删除）。

##### Example

``` go
```

#### (*IntArray) Rands

```go
func (a *IntArray) Rands(size int) []int
```

Rands randomly returns `size` items from array(no deleting).

​	Rands 从数组中随机返回 `size` 项目（不删除）。

##### Example

``` go
```

#### (*IntArray) Range

```go
func (a *IntArray) Range(start int, end ...int) []int
```

Range picks and returns items by range, like array[start:end]. Notice, if in concurrent-safe usage, it returns a copy of slice; else a pointer to the underlying data.

​	范围按范围选取和返回项目，如 array[start：end]。请注意，如果在并发安全使用中，它会返回 slice 的副本;else 指向基础数据的指针。

If `end` is negative, then the offset will start from the end of array. If `end` is omitted, then the sequence will have everything from start up until the end of the array.

​	如果 `end` 为负数，则偏移量将从数组末尾开始。如果 `end` 省略，则序列将包含从数组开始到结束的所有内容。

##### Example

``` go
```

#### (*IntArray) Remove

```go
func (a *IntArray) Remove(index int) (value int, found bool)
```

Remove removes an item by index. If the given `index` is out of range of the array, the `found` is false.

​	“删除”（Remove） 按索引删除项目。如果给定 `index` 的超出数组的范围，则为 `found` false。

##### Example

``` go
```

#### (*IntArray) RemoveValue

```go
func (a *IntArray) RemoveValue(value int) bool
```

RemoveValue removes an item by value. It returns true if value is found in the array, or else false if not found.

​	RemoveValue 按值删除项目。如果在数组中找到值，则返回 true，如果未找到，则返回 false。

##### Example

``` go
```

#### (*IntArray) RemoveValues

```go
func (a *IntArray) RemoveValues(values ...int)
```

RemoveValues removes multiple items by `values`.

​	RemoveValues 按 `values` 删除多个项目。

#### (*IntArray) Replace

```go
func (a *IntArray) Replace(array []int) *IntArray
```

Replace replaces the array items by given `array` from the beginning of array.

​	Replace 将数组项替换为从数组开头开始给定 `array` 的数组项。

##### Example

``` go
```

#### (*IntArray) Reverse

```go
func (a *IntArray) Reverse() *IntArray
```

Reverse makes array with elements in reverse order.

​	Reverse 使数组具有相反顺序的元素。

##### Example

``` go
```

#### (*IntArray) Search

```go
func (a *IntArray) Search(value int) int
```

Search searches array by `value`, returns the index of `value`, or returns -1 if not exists.

​	搜索 搜索数 `value` 组 by 返回 的 `value` 索引，如果不存在，则返回 -1。

##### Example

``` go
```

#### (*IntArray) Set

```go
func (a *IntArray) Set(index int, value int) error
```

Set sets value to specified index.

​	将设置值设置为指定索引。

##### Example

``` go
```

#### (*IntArray) SetArray

```go
func (a *IntArray) SetArray(array []int) *IntArray
```

SetArray sets the underlying slice array with the given `array`.

​	SetArray 使用给定 `array` 的 .

##### Example

``` go
```

#### (*IntArray) Shuffle

```go
func (a *IntArray) Shuffle() *IntArray
```

Shuffle randomly shuffles the array.

​	随机洗牌数组。

##### Example

``` go
```

#### (*IntArray) Slice

```go
func (a *IntArray) Slice() []int
```

Slice returns the underlying data of array. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Slice 返回数组的基础数据。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

##### Example

``` go
```

#### (*IntArray) Sort

```go
func (a *IntArray) Sort(reverse ...bool) *IntArray
```

Sort sorts the array in increasing order. The parameter `reverse` controls whether sort in increasing order(default) or decreasing order.

​	排序按升序对数组进行排序。该参数 `reverse` 控制是按递增顺序（默认）排序还是按递减顺序排序。

##### Example

``` go
```

#### (*IntArray) SortFunc

```go
func (a *IntArray) SortFunc(less func(v1, v2 int) bool) *IntArray
```

SortFunc sorts the array by custom function `less`.

​	SortFunc 按自定义函数 `less` 对数组进行排序。

##### Example

``` go
```

#### (*IntArray) String

```go
func (a *IntArray) String() string
```

String returns current array as a string, which implements like json.Marshal does.

​	String 以字符串形式返回当前数组，其实现方式类似于 json。元帅做到了。

##### Example

``` go
```

#### (*IntArray) SubSlice

```go
func (a *IntArray) SubSlice(offset int, length ...int) []int
```

SubSlice returns a slice of elements from the array as specified by the `offset` and `size` parameters. If in concurrent safe usage, it returns a copy of the slice; else a pointer.

​	SubSlice 从 `offset` 数组中返回由 and `size` 参数指定的元素切片。如果处于并发安全使用状态，则返回切片的副本;else 指针。

If offset is non-negative, the sequence will start at that offset in the array. If offset is negative, the sequence will start that far from the end of the array.

​	如果偏移量为非负数，则序列将从数组中的该偏移量开始。如果偏移量为负数，则序列将从数组末尾开始。

If length is given and is positive, then the sequence will have up to that many elements in it. If the array is shorter than the length, then only the available array elements will be present. If length is given and is negative then the sequence will stop that many elements from the end of the array. If it is omitted, then the sequence will have everything from offset up until the end of the array.

​	如果给定长度并且是正数，则序列中将包含多达那么多元素。如果数组短于长度，则仅存在可用的数组元素。如果给定长度并且是负数，则序列将从数组末尾停止该许多元素。如果省略它，则序列将包含从偏移量到数组末尾的所有内容。

Any possibility crossing the left border of array, it will fail.

​	任何越过数组左边界的可能性，都会失败。

##### Example

``` go
```

#### (*IntArray) Sum

```go
func (a *IntArray) Sum() (sum int)
```

Sum returns the sum of values in an array.

​	Sum 返回数组中值的总和。

##### Example

``` go
```

#### (*IntArray) Unique

```go
func (a *IntArray) Unique() *IntArray
```

Unique uniques the array, clear repeated items. Example: [1,1,2,3,2] -> [1,2,3]

​	唯一唯一数组，清除重复项。示例：[1,1,2,3,2] -> [1,2,3]

##### Example

``` go
```

#### (*IntArray) UnmarshalJSON

```go
func (a *IntArray) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*IntArray) UnmarshalValue

```go
func (a *IntArray) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for array.

​	UnmarshalValue 是一个接口实现，用于为数组设置任何类型的值。

##### Example

``` go
```

#### (*IntArray) Walk

```go
func (a *IntArray) Walk(f func(value int) int) *IntArray
```

Walk applies a user supplied function `f` to every item of array.

​	Walk 将用户提供的函数 `f` 应用于数组的每个项。

##### Example

``` go
```

### type SortedArray

```go
type SortedArray struct {
	// contains filtered or unexported fields
}
```

SortedArray is a golang sorted array with rich features. It is using increasing order in default, which can be changed by setting it a custom comparator. It contains a concurrent-safe/unsafe switch, which should be set when its initialization and cannot be changed then.

​	SortedArray 是一个功能丰富的 golang 排序数组。默认情况下，它使用递增顺序，可以通过将其设置为自定义比较器来更改。它包含一个并发安全/不安全开关，应在初始化时设置该开关，然后无法更改。

#### func NewSortedArray

```go
func NewSortedArray(comparator func(a, b interface{}) int, safe ...bool) *SortedArray
```

NewSortedArray creates and returns an empty sorted array. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default. The parameter `comparator` used to compare values to sort in array, if it returns value < 0, means `a` < `b`; the `a` will be inserted before `b`; if it returns value = 0, means `a` = `b`; the `a` will be replaced by `b`; if it returns value > 0, means `a` > `b`; the `a` will be inserted after `b`;

​	NewSortedArray 创建并返回一个空的排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。 `comparator` 用于比较值以在数组中排序的参数，如果它返回值 < 0，则表示 `a` < `b` ;将 `a` 插入之前 `b` ;如果它返回值 = 0，则意味着 `a` = `b` ;将 `a` 替换为 `b` ;如果它返回值 > 0，则表示 `a` > `b` ;将 `a` 插入后 `b` ;

#### func NewSortedArrayFrom

```go
func NewSortedArrayFrom(array []interface{}, comparator func(a, b interface{}) int, safe ...bool) *SortedArray
```

NewSortedArrayFrom creates and returns an sorted array with given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedArrayFrom 创建并返回具有给定切片 `array` 的排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### func NewSortedArrayFromCopy

```go
func NewSortedArrayFromCopy(array []interface{}, comparator func(a, b interface{}) int, safe ...bool) *SortedArray
```

NewSortedArrayFromCopy creates and returns an sorted array from a copy of given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedArrayFromCopy 从给定切片 `array` 的副本创建并返回一个排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### func NewSortedArrayRange

```go
func NewSortedArrayRange(start, end, step int, comparator func(a, b interface{}) int, safe ...bool) *SortedArray
```

NewSortedArrayRange creates and returns an array by a range from `start` to `end` with step value `step`.

​	NewSortedArrayRange 创建并返回一个数组，其范围为 从 `start` 到 `end` 的步长值 `step` 。

#### func NewSortedArraySize

```go
func NewSortedArraySize(cap int, comparator func(a, b interface{}) int, safe ...bool) *SortedArray
```

NewSortedArraySize create and returns an sorted array with given size and cap. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedArraySize 创建并返回具有给定大小和上限的排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### (*SortedArray) Add

```go
func (a *SortedArray) Add(values ...interface{}) *SortedArray
```

Add adds one or multiple values to sorted array, the array always keeps sorted. It’s alias of function Append, see Append.

​	Add 将一个或多个值添加到排序数组中，该数组始终保持排序状态。它是函数 Append 的别名，请参阅 Append。

#### (*SortedArray) Append

```go
func (a *SortedArray) Append(values ...interface{}) *SortedArray
```

Append adds one or multiple values to sorted array, the array always keeps sorted.

​	Append 将一个或多个值添加到排序数组，该数组始终保持排序状态。

#### (*SortedArray) At

```go
func (a *SortedArray) At(index int) (value interface{})
```

At returns the value by the specified index. If the given `index` is out of range of the array, it returns `nil`.

​	at 返回指定索引的值。如果给定 `index` 的超出数组的范围，则返回 `nil` 。

#### (*SortedArray) Chunk

```go
func (a *SortedArray) Chunk(size int) [][]interface{}
```

Chunk splits an array into multiple arrays, the size of each array is determined by `size`. The last chunk may contain less than size elements.

​	Chunk 将一个数组拆分为多个数组，每个数组的大小由 `size` 决定。最后一个块可能包含小于 size 的元素。

#### (*SortedArray) Clear

```go
func (a *SortedArray) Clear() *SortedArray
```

Clear deletes all items of current array.

​	清除将删除当前数组的所有项目。

#### (*SortedArray) Clone

```go
func (a *SortedArray) Clone() (newArray *SortedArray)
```

Clone returns a new array, which is a copy of current array.

​	克隆返回一个新数组，它是当前数组的副本。

#### (*SortedArray) Contains

```go
func (a *SortedArray) Contains(value interface{}) bool
```

Contains checks whether a value exists in the array.

​	包含检查数组中是否存在值。

#### (*SortedArray) CountValues

```go
func (a *SortedArray) CountValues() map[interface{}]int
```

CountValues counts the number of occurrences of all values in the array.

​	CountValues 计算数组中所有值的出现次数。

#### (*SortedArray) DeepCopy

```go
func (a *SortedArray) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*SortedArray) Filter

```go
func (a *SortedArray) Filter(filter func(index int, value interface{}) bool) *SortedArray
```

Filter iterates array and filters elements using custom callback function. It removes the element from array if callback function `filter` returns true, it or else does nothing and continues iterating.

​	Filter 使用自定义回调函数迭代数组和筛选元素。如果回调函数 `filter` 返回 true，它会从数组中删除该元素，否则它不执行任何操作并继续迭代。

#### (*SortedArray) FilterEmpty

```go
func (a *SortedArray) FilterEmpty() *SortedArray
```

FilterEmpty removes all empty value of the array. Values like: 0, nil, false, “”, len(slice/map/chan) == 0 are considered empty.

​	FilterEmpty 删除数组的所有空值。像 0， nil， false， “”， len（slice/map/chan） == 0 这样的值被视为空。

#### (*SortedArray) FilterNil

```go
func (a *SortedArray) FilterNil() *SortedArray
```

FilterNil removes all nil value of the array.

​	FilterNil 删除数组的所有 nil 值。

#### (*SortedArray) Get

```go
func (a *SortedArray) Get(index int) (value interface{}, found bool)
```

Get returns the value by the specified index. If the given `index` is out of range of the array, the `found` is false.

​	Get 按指定索引返回值。如果给定 `index` 的超出数组的范围，则为 `found` false。

#### (*SortedArray) Interfaces

```go
func (a *SortedArray) Interfaces() []interface{}
```

Interfaces returns current array as []interface{}.

​	Interfaces 以 []interface{} 的形式返回当前数组。

#### (*SortedArray) IsEmpty

```go
func (a *SortedArray) IsEmpty() bool
```

IsEmpty checks whether the array is empty.

​	IsEmpty 检查数组是否为空。

#### (*SortedArray) Iterator

```go
func (a *SortedArray) Iterator(f func(k int, v interface{}) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

#### (*SortedArray) IteratorAsc

```go
func (a *SortedArray) IteratorAsc(f func(k int, v interface{}) bool)
```

IteratorAsc iterates the array readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序只读方式迭代数组。如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*SortedArray) IteratorDesc

```go
func (a *SortedArray) IteratorDesc(f func(k int, v interface{}) bool)
```

IteratorDesc iterates the array readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序迭代数组只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*SortedArray) Join

```go
func (a *SortedArray) Join(glue string) string
```

Join joins array elements with a string `glue`.

​	Join 使用字符串 `glue` 连接数组元素。

#### (*SortedArray) Len

```go
func (a *SortedArray) Len() int
```

Len returns the length of array.

​	Len 返回数组的长度。

#### (*SortedArray) LockFunc

```go
func (a *SortedArray) LockFunc(f func(array []interface{})) *SortedArray
```

LockFunc locks writing by callback function `f`.

​	LockFunc 通过回调函数 `f` 锁定写入。

#### (SortedArray) MarshalJSON

```go
func (a SortedArray) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. Note that do not use pointer as its receiver here.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。请注意，此处不要使用指针作为其接收器。

#### (*SortedArray) Merge

```go
func (a *SortedArray) Merge(array interface{}) *SortedArray
```

Merge merges `array` into current array. The parameter `array` can be any garray or slice type. The difference between Merge and Append is Append supports only specified slice type, but Merge supports more parameter types.

​	合并 `array` 合并到当前数组中。该参数 `array` 可以是任何 garray 或 slice 类型。Merge 和 Append 的区别在于 Append 仅支持指定的切片类型，但 Merge 支持更多参数类型。

#### (*SortedArray) PopLeft

```go
func (a *SortedArray) PopLeft() (value interface{}, found bool)
```

PopLeft pops and returns an item from the beginning of array. Note that if the array is empty, the `found` is false.

​	PopLeft 弹出并返回数组开头的项。请注意，如果数组为空，则为 `found` false。

#### (*SortedArray) PopLefts

```go
func (a *SortedArray) PopLefts(size int) []interface{}
```

PopLefts pops and returns `size` items from the beginning of array.

​	PopLefts 弹出并返回 `size` 数组开头的项。

#### (*SortedArray) PopRand

```go
func (a *SortedArray) PopRand() (value interface{}, found bool)
```

PopRand randomly pops and return an item out of array. Note that if the array is empty, the `found` is false.

​	PopRand 随机弹出并返回数组外的项目。请注意，如果数组为空，则为 `found` false。

#### (*SortedArray) PopRands

```go
func (a *SortedArray) PopRands(size int) []interface{}
```

PopRands randomly pops and returns `size` items out of array.

​	PopRands 随机弹出并返回 `size` 数组外的项目。

#### (*SortedArray) PopRight

```go
func (a *SortedArray) PopRight() (value interface{}, found bool)
```

PopRight pops and returns an item from the end of array. Note that if the array is empty, the `found` is false.

​	PopRight 弹出并返回数组末尾的项。请注意，如果数组为空，则为 `found` false。

#### (*SortedArray) PopRights

```go
func (a *SortedArray) PopRights(size int) []interface{}
```

PopRights pops and returns `size` items from the end of array.

​	PopRights 弹出并返回 `size` 数组末尾的项。

#### (*SortedArray) RLockFunc

```go
func (a *SortedArray) RLockFunc(f func(array []interface{})) *SortedArray
```

RLockFunc locks reading by callback function `f`.

​	RLockFunc 通过回调函数 `f` 锁定读取。

#### (*SortedArray) Rand

```go
func (a *SortedArray) Rand() (value interface{}, found bool)
```

Rand randomly returns one item from array(no deleting).

​	Rand 从数组中随机返回一个项目（不删除）。

#### (*SortedArray) Rands

```go
func (a *SortedArray) Rands(size int) []interface{}
```

Rands randomly returns `size` items from array(no deleting).

​	Rands 从数组中随机返回 `size` 项目（不删除）。

#### (*SortedArray) Range

```go
func (a *SortedArray) Range(start int, end ...int) []interface{}
```

Range picks and returns items by range, like array[start:end]. Notice, if in concurrent-safe usage, it returns a copy of slice; else a pointer to the underlying data.

​	范围按范围选取和返回项目，如 array[start：end]。请注意，如果在并发安全使用中，它会返回 slice 的副本;else 指向基础数据的指针。

If `end` is negative, then the offset will start from the end of array. If `end` is omitted, then the sequence will have everything from start up until the end of the array.

​	如果 `end` 为负数，则偏移量将从数组末尾开始。如果 `end` 省略，则序列将包含从数组开始到结束的所有内容。

#### (*SortedArray) Remove

```go
func (a *SortedArray) Remove(index int) (value interface{}, found bool)
```

Remove removes an item by index. If the given `index` is out of range of the array, the `found` is false.

​	“删除”（Remove） 按索引删除项目。如果给定 `index` 的超出数组的范围，则为 `found` false。

#### (*SortedArray) RemoveValue

```go
func (a *SortedArray) RemoveValue(value interface{}) bool
```

RemoveValue removes an item by value. It returns true if value is found in the array, or else false if not found.

​	RemoveValue 按值删除项目。如果在数组中找到值，则返回 true，如果未找到，则返回 false。

#### (*SortedArray) RemoveValues

```go
func (a *SortedArray) RemoveValues(values ...interface{})
```

RemoveValues removes an item by `values`.

​	RemoveValues 通过 `values` 删除项目。

#### (*SortedArray) Search

```go
func (a *SortedArray) Search(value interface{}) (index int)
```

Search searches array by `value`, returns the index of `value`, or returns -1 if not exists.

​	搜索 搜索数 `value` 组 by 返回 的 `value` 索引，如果不存在，则返回 -1。

#### (*SortedArray) SetArray

```go
func (a *SortedArray) SetArray(array []interface{}) *SortedArray
```

SetArray sets the underlying slice array with the given `array`.

​	SetArray 使用给定 `array` 的 .

#### (*SortedArray) SetComparator

```go
func (a *SortedArray) SetComparator(comparator func(a, b interface{}) int)
```

SetComparator sets/changes the comparator for sorting. It resorts the array as the comparator is changed.

​	SetComparator 设置/更改用于排序的比较器。当比较器发生变化时，它会对数组进行选取。

#### (*SortedArray) SetUnique

```go
func (a *SortedArray) SetUnique(unique bool) *SortedArray
```

SetUnique sets unique mark to the array, which means it does not contain any repeated items. It also does unique check, remove all repeated items.

​	SetUnique 为数组设置唯一标记，这意味着它不包含任何重复项。它还进行唯一检查，删除所有重复的项目。

#### (*SortedArray) Slice

```go
func (a *SortedArray) Slice() []interface{}
```

Slice returns the underlying data of array. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Slice 返回数组的基础数据。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

#### (*SortedArray) Sort

```go
func (a *SortedArray) Sort() *SortedArray
```

Sort sorts the array in increasing order. The parameter `reverse` controls whether sort in increasing order(default) or decreasing order

​	排序按升序对数组进行排序。该参数 `reverse` 控制是按递增顺序（默认）排序还是按递减顺序排序

#### (*SortedArray) String

```go
func (a *SortedArray) String() string
```

String returns current array as a string, which implements like json.Marshal does.

​	String 以字符串形式返回当前数组，其实现方式类似于 json。元帅做到了。

#### (*SortedArray) SubSlice

```go
func (a *SortedArray) SubSlice(offset int, length ...int) []interface{}
```

SubSlice returns a slice of elements from the array as specified by the `offset` and `size` parameters. If in concurrent safe usage, it returns a copy of the slice; else a pointer.

​	SubSlice 从 `offset` 数组中返回由 and `size` 参数指定的元素切片。如果处于并发安全使用状态，则返回切片的副本;else 指针。

If offset is non-negative, the sequence will start at that offset in the array. If offset is negative, the sequence will start that far from the end of the array.

​	如果偏移量为非负数，则序列将从数组中的该偏移量开始。如果偏移量为负数，则序列将从数组末尾开始。

If length is given and is positive, then the sequence will have up to that many elements in it. If the array is shorter than the length, then only the available array elements will be present. If length is given and is negative then the sequence will stop that many elements from the end of the array. If it is omitted, then the sequence will have everything from offset up until the end of the array.

​	如果给定长度并且是正数，则序列中将包含多达那么多元素。如果数组短于长度，则仅存在可用的数组元素。如果给定长度并且是负数，则序列将从数组末尾停止该许多元素。如果省略它，则序列将包含从偏移量到数组末尾的所有内容。

Any possibility crossing the left border of array, it will fail.

​	任何越过数组左边界的可能性，都会失败。

#### (*SortedArray) Sum

```go
func (a *SortedArray) Sum() (sum int)
```

Sum returns the sum of values in an array.

​	Sum 返回数组中值的总和。

#### (*SortedArray) Unique

```go
func (a *SortedArray) Unique() *SortedArray
```

Unique uniques the array, clear repeated items.

​	唯一唯一数组，清除重复项。

#### (*SortedArray) UnmarshalJSON

```go
func (a *SortedArray) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal. Note that the comparator is set as string comparator in default.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。请注意，默认情况下，比较器设置为字符串比较器。

#### (*SortedArray) UnmarshalValue

```go
func (a *SortedArray) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for array. Note that the comparator is set as string comparator in default.

​	UnmarshalValue 是一个接口实现，用于为数组设置任何类型的值。请注意，默认情况下，比较器设置为字符串比较器。

#### (*SortedArray) Walk

```go
func (a *SortedArray) Walk(f func(value interface{}) interface{}) *SortedArray
```

Walk applies a user supplied function `f` to every item of array.

​	Walk 将用户提供的函数 `f` 应用于数组的每个项。

### type SortedIntArray

```go
type SortedIntArray struct {
	// contains filtered or unexported fields
}
```

SortedIntArray is a golang sorted int array with rich features. It is using increasing order in default, which can be changed by setting it a custom comparator. It contains a concurrent-safe/unsafe switch, which should be set when its initialization and cannot be changed then.

​	SortedIntArray 是一个功能丰富的 golang 排序 int 数组。默认情况下，它使用递增顺序，可以通过将其设置为自定义比较器来更改。它包含一个并发安全/不安全开关，应在初始化时设置该开关，然后无法更改。

#### func NewSortedIntArray

```go
func NewSortedIntArray(safe ...bool) *SortedIntArray
```

NewSortedIntArray creates and returns an empty sorted array. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedIntArray 创建并返回一个空的排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### func NewSortedIntArrayComparator

```go
func NewSortedIntArrayComparator(comparator func(a, b int) int, safe ...bool) *SortedIntArray
```

NewSortedIntArrayComparator creates and returns an empty sorted array with specified comparator. The parameter `safe` is used to specify whether using array in concurrent-safety which is false in default.

​	NewSortedIntArrayComparator 创建并返回一个具有指定比较器的空排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 array，默认为 false。

#### func NewSortedIntArrayFrom

```go
func NewSortedIntArrayFrom(array []int, safe ...bool) *SortedIntArray
```

NewSortedIntArrayFrom creates and returns an sorted array with given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedIntArrayFrom 创建并返回具有给定切片 `array` 的排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### func NewSortedIntArrayFromCopy

```go
func NewSortedIntArrayFromCopy(array []int, safe ...bool) *SortedIntArray
```

NewSortedIntArrayFromCopy creates and returns an sorted array from a copy of given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedIntArrayFromCopy 从给定切片 `array` 的副本创建并返回一个排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### func NewSortedIntArrayRange

```go
func NewSortedIntArrayRange(start, end, step int, safe ...bool) *SortedIntArray
```

NewSortedIntArrayRange creates and returns an array by a range from `start` to `end` with step value `step`.

​	NewSortedIntArrayRange 创建并返回一个数组，其范围为 `start` 从 到 `end` 的步长值 `step` 。

#### func NewSortedIntArraySize

```go
func NewSortedIntArraySize(cap int, safe ...bool) *SortedIntArray
```

NewSortedIntArraySize create and returns an sorted array with given size and cap. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedIntArraySize 创建并返回具有给定大小和上限的排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### (*SortedIntArray) Add

```go
func (a *SortedIntArray) Add(values ...int) *SortedIntArray
```

Add adds one or multiple values to sorted array, the array always keeps sorted. It’s alias of function Append, see Append.

​	Add 将一个或多个值添加到排序数组中，该数组始终保持排序状态。它是函数 Append 的别名，请参阅 Append。

#### (*SortedIntArray) Append

```go
func (a *SortedIntArray) Append(values ...int) *SortedIntArray
```

Append adds one or multiple values to sorted array, the array always keeps sorted.

​	Append 将一个或多个值添加到排序数组，该数组始终保持排序状态。

#### (*SortedIntArray) At

```go
func (a *SortedIntArray) At(index int) (value int)
```

At returns the value by the specified index. If the given `index` is out of range of the array, it returns `0`.

​	at 返回指定索引的值。如果给定 `index` 的超出数组的范围，则返回 `0` 。

#### (*SortedIntArray) Chunk

```go
func (a *SortedIntArray) Chunk(size int) [][]int
```

Chunk splits an array into multiple arrays, the size of each array is determined by `size`. The last chunk may contain less than size elements.

​	Chunk 将一个数组拆分为多个数组，每个数组的大小由 `size` 决定。最后一个块可能包含小于 size 的元素。

#### (*SortedIntArray) Clear

```go
func (a *SortedIntArray) Clear() *SortedIntArray
```

Clear deletes all items of current array.

​	清除将删除当前数组的所有项目。

#### (*SortedIntArray) Clone

```go
func (a *SortedIntArray) Clone() (newArray *SortedIntArray)
```

Clone returns a new array, which is a copy of current array.

​	克隆返回一个新数组，它是当前数组的副本。

#### (*SortedIntArray) Contains

```go
func (a *SortedIntArray) Contains(value int) bool
```

Contains checks whether a value exists in the array.

​	包含检查数组中是否存在值。

#### (*SortedIntArray) CountValues

```go
func (a *SortedIntArray) CountValues() map[int]int
```

CountValues counts the number of occurrences of all values in the array.

​	CountValues 计算数组中所有值的出现次数。

#### (*SortedIntArray) DeepCopy

```go
func (a *SortedIntArray) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*SortedIntArray) Filter

```go
func (a *SortedIntArray) Filter(filter func(index int, value int) bool) *SortedIntArray
```

Filter iterates array and filters elements using custom callback function. It removes the element from array if callback function `filter` returns true, it or else does nothing and continues iterating.

​	Filter 使用自定义回调函数迭代数组和筛选元素。如果回调函数 `filter` 返回 true，它会从数组中删除该元素，否则它不执行任何操作并继续迭代。

#### (*SortedIntArray) FilterEmpty

```go
func (a *SortedIntArray) FilterEmpty() *SortedIntArray
```

FilterEmpty removes all zero value of the array.

​	FilterEmpty 删除数组的所有零值。

#### (*SortedIntArray) Get

```go
func (a *SortedIntArray) Get(index int) (value int, found bool)
```

Get returns the value by the specified index. If the given `index` is out of range of the array, the `found` is false.

​	Get 按指定索引返回值。如果给定 `index` 的超出数组的范围，则为 `found` false。

#### (*SortedIntArray) Interfaces

```go
func (a *SortedIntArray) Interfaces() []interface{}
```

Interfaces returns current array as []interface{}.

​	Interfaces 以 []interface{} 的形式返回当前数组。

#### (*SortedIntArray) IsEmpty

```go
func (a *SortedIntArray) IsEmpty() bool
```

IsEmpty checks whether the array is empty.

​	IsEmpty 检查数组是否为空。

#### (*SortedIntArray) Iterator

```go
func (a *SortedIntArray) Iterator(f func(k int, v int) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

#### (*SortedIntArray) IteratorAsc

```go
func (a *SortedIntArray) IteratorAsc(f func(k int, v int) bool)
```

IteratorAsc iterates the array readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序只读方式迭代数组。如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*SortedIntArray) IteratorDesc

```go
func (a *SortedIntArray) IteratorDesc(f func(k int, v int) bool)
```

IteratorDesc iterates the array readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序迭代数组只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

#### (*SortedIntArray) Join

```go
func (a *SortedIntArray) Join(glue string) string
```

Join joins array elements with a string `glue`.

​	Join 使用字符串 `glue` 连接数组元素。

#### (*SortedIntArray) Len

```go
func (a *SortedIntArray) Len() int
```

Len returns the length of array.

​	Len 返回数组的长度。

#### (*SortedIntArray) LockFunc

```go
func (a *SortedIntArray) LockFunc(f func(array []int)) *SortedIntArray
```

LockFunc locks writing by callback function `f`.

​	LockFunc 通过回调函数 `f` 锁定写入。

#### (SortedIntArray) MarshalJSON

```go
func (a SortedIntArray) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. Note that do not use pointer as its receiver here.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。请注意，此处不要使用指针作为其接收器。

#### (*SortedIntArray) Merge

```go
func (a *SortedIntArray) Merge(array interface{}) *SortedIntArray
```

Merge merges `array` into current array. The parameter `array` can be any garray or slice type. The difference between Merge and Append is Append supports only specified slice type, but Merge supports more parameter types.

​	合并 `array` 合并到当前数组中。该参数 `array` 可以是任何 garray 或 slice 类型。Merge 和 Append 的区别在于 Append 仅支持指定的切片类型，但 Merge 支持更多参数类型。

#### (*SortedIntArray) PopLeft

```go
func (a *SortedIntArray) PopLeft() (value int, found bool)
```

PopLeft pops and returns an item from the beginning of array. Note that if the array is empty, the `found` is false.

​	PopLeft 弹出并返回数组开头的项。请注意，如果数组为空，则为 `found` false。

#### (*SortedIntArray) PopLefts

```go
func (a *SortedIntArray) PopLefts(size int) []int
```

PopLefts pops and returns `size` items from the beginning of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopLefts 弹出并返回 `size` 数组开头的项。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

#### (*SortedIntArray) PopRand

```go
func (a *SortedIntArray) PopRand() (value int, found bool)
```

PopRand randomly pops and return an item out of array. Note that if the array is empty, the `found` is false.

​	PopRand 随机弹出并返回数组外的项目。请注意，如果数组为空，则为 `found` false。

#### (*SortedIntArray) PopRands

```go
func (a *SortedIntArray) PopRands(size int) []int
```

PopRands randomly pops and returns `size` items out of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopRands 随机弹出并返回 `size` 数组外的项目。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

#### (*SortedIntArray) PopRight

```go
func (a *SortedIntArray) PopRight() (value int, found bool)
```

PopRight pops and returns an item from the end of array. Note that if the array is empty, the `found` is false.

​	PopRight 弹出并返回数组末尾的项。请注意，如果数组为空，则为 `found` false。

#### (*SortedIntArray) PopRights

```go
func (a *SortedIntArray) PopRights(size int) []int
```

PopRights pops and returns `size` items from the end of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopRights 弹出并返回 `size` 数组末尾的项。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

#### (*SortedIntArray) RLockFunc

```go
func (a *SortedIntArray) RLockFunc(f func(array []int)) *SortedIntArray
```

RLockFunc locks reading by callback function `f`.

​	RLockFunc 通过回调函数 `f` 锁定读取。

#### (*SortedIntArray) Rand

```go
func (a *SortedIntArray) Rand() (value int, found bool)
```

Rand randomly returns one item from array(no deleting).

​	Rand 从数组中随机返回一个项目（不删除）。

#### (*SortedIntArray) Rands

```go
func (a *SortedIntArray) Rands(size int) []int
```

Rands randomly returns `size` items from array(no deleting).

​	Rands 从数组中随机返回 `size` 项目（不删除）。

#### (*SortedIntArray) Range

```go
func (a *SortedIntArray) Range(start int, end ...int) []int
```

Range picks and returns items by range, like array[start:end]. Notice, if in concurrent-safe usage, it returns a copy of slice; else a pointer to the underlying data.

​	范围按范围选取和返回项目，如 array[start：end]。请注意，如果在并发安全使用中，它会返回 slice 的副本;else 指向基础数据的指针。

If `end` is negative, then the offset will start from the end of array. If `end` is omitted, then the sequence will have everything from start up until the end of the array.

​	如果 `end` 为负数，则偏移量将从数组末尾开始。如果 `end` 省略，则序列将包含从数组开始到结束的所有内容。

#### (*SortedIntArray) Remove

```go
func (a *SortedIntArray) Remove(index int) (value int, found bool)
```

Remove removes an item by index. If the given `index` is out of range of the array, the `found` is false.

​	“删除”（Remove） 按索引删除项目。如果给定 `index` 的超出数组的范围，则为 `found` false。

#### (*SortedIntArray) RemoveValue

```go
func (a *SortedIntArray) RemoveValue(value int) bool
```

RemoveValue removes an item by value. It returns true if value is found in the array, or else false if not found.

​	RemoveValue 按值删除项目。如果在数组中找到值，则返回 true，如果未找到，则返回 false。

#### (*SortedIntArray) RemoveValues

```go
func (a *SortedIntArray) RemoveValues(values ...int)
```

RemoveValues removes an item by `values`.

​	RemoveValues 通过 `values` 删除项目。

#### (*SortedIntArray) Search

```go
func (a *SortedIntArray) Search(value int) (index int)
```

Search searches array by `value`, returns the index of `value`, or returns -1 if not exists.

​	搜索 搜索数 `value` 组 by 返回 的 `value` 索引，如果不存在，则返回 -1。

#### (*SortedIntArray) SetArray

```go
func (a *SortedIntArray) SetArray(array []int) *SortedIntArray
```

SetArray sets the underlying slice array with the given `array`.

​	SetArray 使用给定 `array` 的 .

#### (*SortedIntArray) SetUnique

```go
func (a *SortedIntArray) SetUnique(unique bool) *SortedIntArray
```

SetUnique sets unique mark to the array, which means it does not contain any repeated items. It also do unique check, remove all repeated items.

​	SetUnique 为数组设置唯一标记，这意味着它不包含任何重复项。它还进行唯一检查，删除所有重复的项目。

#### (*SortedIntArray) Slice

```go
func (a *SortedIntArray) Slice() []int
```

Slice returns the underlying data of array. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Slice 返回数组的基础数据。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

#### (*SortedIntArray) Sort

```go
func (a *SortedIntArray) Sort() *SortedIntArray
```

Sort sorts the array in increasing order. The parameter `reverse` controls whether sort in increasing order(default) or decreasing order.

​	排序按升序对数组进行排序。该参数 `reverse` 控制是按递增顺序（默认）排序还是按递减顺序排序。

#### (*SortedIntArray) String

```go
func (a *SortedIntArray) String() string
```

String returns current array as a string, which implements like json.Marshal does.

​	String 以字符串形式返回当前数组，其实现方式类似于 json。元帅做到了。

#### (*SortedIntArray) SubSlice

```go
func (a *SortedIntArray) SubSlice(offset int, length ...int) []int
```

SubSlice returns a slice of elements from the array as specified by the `offset` and `size` parameters. If in concurrent safe usage, it returns a copy of the slice; else a pointer.

​	SubSlice 从 `offset` 数组中返回由 and `size` 参数指定的元素切片。如果处于并发安全使用状态，则返回切片的副本;else 指针。

If offset is non-negative, the sequence will start at that offset in the array. If offset is negative, the sequence will start that far from the end of the array.

​	如果偏移量为非负数，则序列将从数组中的该偏移量开始。如果偏移量为负数，则序列将从数组末尾开始。

If length is given and is positive, then the sequence will have up to that many elements in it. If the array is shorter than the length, then only the available array elements will be present. If length is given and is negative then the sequence will stop that many elements from the end of the array. If it is omitted, then the sequence will have everything from offset up until the end of the array.

​	如果给定长度并且是正数，则序列中将包含多达那么多元素。如果数组短于长度，则仅存在可用的数组元素。如果给定长度并且是负数，则序列将从数组末尾停止该许多元素。如果省略它，则序列将包含从偏移量到数组末尾的所有内容。

Any possibility crossing the left border of array, it will fail.

​	任何越过数组左边界的可能性，都会失败。

#### (*SortedIntArray) Sum

```go
func (a *SortedIntArray) Sum() (sum int)
```

Sum returns the sum of values in an array.

​	Sum 返回数组中值的总和。

#### (*SortedIntArray) Unique

```go
func (a *SortedIntArray) Unique() *SortedIntArray
```

Unique uniques the array, clear repeated items.

​	唯一唯一数组，清除重复项。

#### (*SortedIntArray) UnmarshalJSON

```go
func (a *SortedIntArray) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*SortedIntArray) UnmarshalValue

```go
func (a *SortedIntArray) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for array.

​	UnmarshalValue 是一个接口实现，用于为数组设置任何类型的值。

#### (*SortedIntArray) Walk

```go
func (a *SortedIntArray) Walk(f func(value int) int) *SortedIntArray
```

Walk applies a user supplied function `f` to every item of array.

​	Walk 将用户提供的函数 `f` 应用于数组的每个项。

### type SortedStrArray

```go
type SortedStrArray struct {
	// contains filtered or unexported fields
}
```

SortedStrArray is a golang sorted string array with rich features. It is using increasing order in default, which can be changed by setting it a custom comparator. It contains a concurrent-safe/unsafe switch, which should be set when its initialization and cannot be changed then.

​	SortedStrArray 是一个具有丰富功能的 golang 排序字符串数组。默认情况下，它使用递增顺序，可以通过将其设置为自定义比较器来更改。它包含一个并发安全/不安全开关，应在初始化时设置该开关，然后无法更改。

#### func NewSortedStrArray

```go
func NewSortedStrArray(safe ...bool) *SortedStrArray
```

NewSortedStrArray creates and returns an empty sorted array. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedStrArray 创建并返回一个空的排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example

``` go
```

#### func NewSortedStrArrayComparator

```go
func NewSortedStrArrayComparator(comparator func(a, b string) int, safe ...bool) *SortedStrArray
```

NewSortedStrArrayComparator creates and returns an empty sorted array with specified comparator. The parameter `safe` is used to specify whether using array in concurrent-safety which is false in default.

​	NewSortedStrArrayComparator 创建并返回一个具有指定比较器的空排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用 array，默认为 false。

#### func NewSortedStrArrayFrom

```go
func NewSortedStrArrayFrom(array []string, safe ...bool) *SortedStrArray
```

NewSortedStrArrayFrom creates and returns an sorted array with given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedStrArrayFrom 创建并返回具有给定切片 `array` 的排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### func NewSortedStrArrayFromCopy

```go
func NewSortedStrArrayFromCopy(array []string, safe ...bool) *SortedStrArray
```

NewSortedStrArrayFromCopy creates and returns an sorted array from a copy of given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedStrArrayFromCopy 从给定切片 `array` 的副本创建并返回一个排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

#### func NewSortedStrArraySize

```go
func NewSortedStrArraySize(cap int, safe ...bool) *SortedStrArray
```

NewSortedStrArraySize create and returns an sorted array with given size and cap. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewSortedStrArraySize 创建并返回具有给定大小和上限的排序数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example

``` go
```

#### (*SortedStrArray) Add

```go
func (a *SortedStrArray) Add(values ...string) *SortedStrArray
```

Add adds one or multiple values to sorted array, the array always keeps sorted. It’s alias of function Append, see Append.

​	Add 将一个或多个值添加到排序数组中，该数组始终保持排序状态。它是函数 Append 的别名，请参阅 Append。

##### Example

``` go
```

#### (*SortedStrArray) Append

```go
func (a *SortedStrArray) Append(values ...string) *SortedStrArray
```

Append adds one or multiple values to sorted array, the array always keeps sorted.

​	Append 将一个或多个值添加到排序数组，该数组始终保持排序状态。

##### Example

``` go
```

#### (*SortedStrArray) At

```go
func (a *SortedStrArray) At(index int) (value string)
```

At returns the value by the specified index. If the given `index` is out of range of the array, it returns an empty string.

​	at 返回指定索引的值。如果给定 `index` 的超出数组的范围，则返回一个空字符串。

##### Example

``` go
```

#### (*SortedStrArray) Chunk

```go
func (a *SortedStrArray) Chunk(size int) [][]string
```

Chunk splits an array into multiple arrays, the size of each array is determined by `size`. The last chunk may contain less than size elements.

​	Chunk 将一个数组拆分为多个数组，每个数组的大小由 `size` 决定。最后一个块可能包含小于 size 的元素。

##### Example

``` go
```

#### (*SortedStrArray) Clear

```go
func (a *SortedStrArray) Clear() *SortedStrArray
```

Clear deletes all items of current array.

​	清除将删除当前数组的所有项目。

##### Example

``` go
```

#### (*SortedStrArray) Clone

```go
func (a *SortedStrArray) Clone() (newArray *SortedStrArray)
```

Clone returns a new array, which is a copy of current array.

​	克隆返回一个新数组，它是当前数组的副本。

##### Example

``` go
```

#### (*SortedStrArray) Contains

```go
func (a *SortedStrArray) Contains(value string) bool
```

Contains checks whether a value exists in the array.

​	包含检查数组中是否存在值。

##### Example

``` go
```

#### (*SortedStrArray) ContainsI

```go
func (a *SortedStrArray) ContainsI(value string) bool
```

ContainsI checks whether a value exists in the array with case-insensitively. Note that it internally iterates the whole array to do the comparison with case-insensitively.

​	ContainsI 使用不区分大小写的方式检查数组中是否存在值。请注意，它在内部迭代整个数组以不区分大小写地进行比较。

##### Example

``` go
```

#### (*SortedStrArray) CountValues

```go
func (a *SortedStrArray) CountValues() map[string]int
```

CountValues counts the number of occurrences of all values in the array.

​	CountValues 计算数组中所有值的出现次数。

##### Example

``` go
```

#### (*SortedStrArray) DeepCopy

```go
func (a *SortedStrArray) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*SortedStrArray) Filter

```go
func (a *SortedStrArray) Filter(filter func(index int, value string) bool) *SortedStrArray
```

Filter iterates array and filters elements using custom callback function. It removes the element from array if callback function `filter` returns true, it or else does nothing and continues iterating.

​	Filter 使用自定义回调函数迭代数组和筛选元素。如果回调函数 `filter` 返回 true，它会从数组中删除该元素，否则它不执行任何操作并继续迭代。

##### Example

``` go
```

#### (*SortedStrArray) FilterEmpty

```go
func (a *SortedStrArray) FilterEmpty() *SortedStrArray
```

FilterEmpty removes all empty string value of the array.

​	FilterEmpty 删除数组的所有空字符串值。

##### Example

``` go
```

#### (*SortedStrArray) Get

```go
func (a *SortedStrArray) Get(index int) (value string, found bool)
```

Get returns the value by the specified index. If the given `index` is out of range of the array, the `found` is false.

​	Get 按指定索引返回值。如果给定 `index` 的超出数组的范围，则为 `found` false。

##### Example

``` go
```

#### (*SortedStrArray) Interfaces

```go
func (a *SortedStrArray) Interfaces() []interface{}
```

Interfaces returns current array as []interface{}.

​	Interfaces 以 []interface{} 的形式返回当前数组。

##### Example

``` go
```

#### (*SortedStrArray) IsEmpty

```go
func (a *SortedStrArray) IsEmpty() bool
```

IsEmpty checks whether the array is empty.

​	IsEmpty 检查数组是否为空。

##### Example

``` go
```

#### (*SortedStrArray) Iterator

```go
func (a *SortedStrArray) Iterator(f func(k int, v string) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

##### Example

``` go
```

#### (*SortedStrArray) IteratorAsc

```go
func (a *SortedStrArray) IteratorAsc(f func(k int, v string) bool)
```

IteratorAsc iterates the array readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序只读方式迭代数组。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*SortedStrArray) IteratorDesc

```go
func (a *SortedStrArray) IteratorDesc(f func(k int, v string) bool)
```

IteratorDesc iterates the array readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序迭代数组只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*SortedStrArray) Join

```go
func (a *SortedStrArray) Join(glue string) string
```

Join joins array elements with a string `glue`.

​	Join 使用字符串 `glue` 连接数组元素。

##### Example

``` go
```

#### (*SortedStrArray) Len

```go
func (a *SortedStrArray) Len() int
```

Len returns the length of array.

​	Len 返回数组的长度。

##### Example

``` go
```

#### (*SortedStrArray) LockFunc

```go
func (a *SortedStrArray) LockFunc(f func(array []string)) *SortedStrArray
```

LockFunc locks writing by callback function `f`.

​	LockFunc 通过回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (SortedStrArray) MarshalJSON

```go
func (a SortedStrArray) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. Note that do not use pointer as its receiver here.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。请注意，此处不要使用指针作为其接收器。

##### Example

``` go
```

#### (*SortedStrArray) Merge

```go
func (a *SortedStrArray) Merge(array interface{}) *SortedStrArray
```

Merge merges `array` into current array. The parameter `array` can be any garray or slice type. The difference between Merge and Append is Append supports only specified slice type, but Merge supports more parameter types.

​	合并 `array` 合并到当前数组中。该参数 `array` 可以是任何 garray 或 slice 类型。Merge 和 Append 的区别在于 Append 仅支持指定的切片类型，但 Merge 支持更多参数类型。

##### Example

``` go
```

#### (*SortedStrArray) PopLeft

```go
func (a *SortedStrArray) PopLeft() (value string, found bool)
```

PopLeft pops and returns an item from the beginning of array. Note that if the array is empty, the `found` is false.

​	PopLeft 弹出并返回数组开头的项。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*SortedStrArray) PopLefts

```go
func (a *SortedStrArray) PopLefts(size int) []string
```

PopLefts pops and returns `size` items from the beginning of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopLefts 弹出并返回 `size` 数组开头的项。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

##### Example

``` go
```

#### (*SortedStrArray) PopRand

```go
func (a *SortedStrArray) PopRand() (value string, found bool)
```

PopRand randomly pops and return an item out of array. Note that if the array is empty, the `found` is false.

​	PopRand 随机弹出并返回数组外的项目。请注意，如果数组为空，则为 `found` false。

#### (*SortedStrArray) PopRands

```go
func (a *SortedStrArray) PopRands(size int) []string
```

PopRands randomly pops and returns `size` items out of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopRands 随机弹出并返回 `size` 数组外的项目。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

##### Example

``` go
```

#### (*SortedStrArray) PopRight

```go
func (a *SortedStrArray) PopRight() (value string, found bool)
```

PopRight pops and returns an item from the end of array. Note that if the array is empty, the `found` is false.

​	PopRight 弹出并返回数组末尾的项。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*SortedStrArray) PopRights

```go
func (a *SortedStrArray) PopRights(size int) []string
```

PopRights pops and returns `size` items from the end of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopRights 弹出并返回 `size` 数组末尾的项。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

##### Example

``` go
```

#### (*SortedStrArray) RLockFunc

```go
func (a *SortedStrArray) RLockFunc(f func(array []string)) *SortedStrArray
```

RLockFunc locks reading by callback function `f`.

​	RLockFunc 通过回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*SortedStrArray) Rand

```go
func (a *SortedStrArray) Rand() (value string, found bool)
```

Rand randomly returns one item from array(no deleting).

​	Rand 从数组中随机返回一个项目（不删除）。

##### Example

``` go
```

#### (*SortedStrArray) Rands

```go
func (a *SortedStrArray) Rands(size int) []string
```

Rands randomly returns `size` items from array(no deleting).

​	Rands 从数组中随机返回 `size` 项目（不删除）。

##### Example

``` go
```

#### (*SortedStrArray) Range

```go
func (a *SortedStrArray) Range(start int, end ...int) []string
```

Range picks and returns items by range, like array[start:end]. Notice, if in concurrent-safe usage, it returns a copy of slice; else a pointer to the underlying data.

​	范围按范围选取和返回项目，如 array[start：end]。请注意，如果在并发安全使用中，它会返回 slice 的副本;else 指向基础数据的指针。

If `end` is negative, then the offset will start from the end of array. If `end` is omitted, then the sequence will have everything from start up until the end of the array.

​	如果 `end` 为负数，则偏移量将从数组末尾开始。如果 `end` 省略，则序列将包含从数组开始到结束的所有内容。

##### Example

``` go
```

#### (*SortedStrArray) Remove

```go
func (a *SortedStrArray) Remove(index int) (value string, found bool)
```

Remove removes an item by index. If the given `index` is out of range of the array, the `found` is false.

​	“删除”（Remove） 按索引删除项目。如果给定 `index` 的超出数组的范围，则为 `found` false。

##### Example

``` go
```

#### (*SortedStrArray) RemoveValue

```go
func (a *SortedStrArray) RemoveValue(value string) bool
```

RemoveValue removes an item by value. It returns true if value is found in the array, or else false if not found.

​	RemoveValue 按值删除项目。如果在数组中找到值，则返回 true，如果未找到，则返回 false。

##### Example

``` go
```

#### (*SortedStrArray) RemoveValues

```go
func (a *SortedStrArray) RemoveValues(values ...string)
```

RemoveValues removes an item by `values`.

​	RemoveValues 通过 `values` 删除项目。

#### (*SortedStrArray) Search

```go
func (a *SortedStrArray) Search(value string) (index int)
```

Search searches array by `value`, returns the index of `value`, or returns -1 if not exists.

​	搜索 搜索数 `value` 组 by 返回 的 `value` 索引，如果不存在，则返回 -1。

##### Example

``` go
```

#### (*SortedStrArray) SetArray

```go
func (a *SortedStrArray) SetArray(array []string) *SortedStrArray
```

SetArray sets the underlying slice array with the given `array`.

​	SetArray 使用给定 `array` 的 .

##### Example

``` go
```

#### (*SortedStrArray) SetUnique

```go
func (a *SortedStrArray) SetUnique(unique bool) *SortedStrArray
```

SetUnique sets unique mark to the array, which means it does not contain any repeated items. It also do unique check, remove all repeated items.

​	SetUnique 为数组设置唯一标记，这意味着它不包含任何重复项。它还进行唯一检查，删除所有重复的项目。

##### Example

``` go
```

#### (*SortedStrArray) Slice

```go
func (a *SortedStrArray) Slice() []string
```

Slice returns the underlying data of array. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Slice 返回数组的基础数据。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

##### Example

``` go
```

#### (*SortedStrArray) Sort

```go
func (a *SortedStrArray) Sort() *SortedStrArray
```

Sort sorts the array in increasing order. The parameter `reverse` controls whether sort in increasing order(default) or decreasing order.

​	排序按升序对数组进行排序。该参数 `reverse` 控制是按递增顺序（默认）排序还是按递减顺序排序。

##### Example

``` go
```

#### (*SortedStrArray) String

```go
func (a *SortedStrArray) String() string
```

String returns current array as a string, which implements like json.Marshal does.

​	String 以字符串形式返回当前数组，其实现方式类似于 json。元帅做到了。

##### Example

``` go
```

#### (*SortedStrArray) SubSlice

```go
func (a *SortedStrArray) SubSlice(offset int, length ...int) []string
```

SubSlice returns a slice of elements from the array as specified by the `offset` and `size` parameters. If in concurrent safe usage, it returns a copy of the slice; else a pointer.

​	SubSlice 从 `offset` 数组中返回由 and `size` 参数指定的元素切片。如果处于并发安全使用状态，则返回切片的副本;else 指针。

If offset is non-negative, the sequence will start at that offset in the array. If offset is negative, the sequence will start that far from the end of the array.

​	如果偏移量为非负数，则序列将从数组中的该偏移量开始。如果偏移量为负数，则序列将从数组末尾开始。

If length is given and is positive, then the sequence will have up to that many elements in it. If the array is shorter than the length, then only the available array elements will be present. If length is given and is negative then the sequence will stop that many elements from the end of the array. If it is omitted, then the sequence will have everything from offset up until the end of the array.

​	如果给定长度并且是正数，则序列中将包含多达那么多元素。如果数组短于长度，则仅存在可用的数组元素。如果给定长度并且是负数，则序列将从数组末尾停止该许多元素。如果省略它，则序列将包含从偏移量到数组末尾的所有内容。

Any possibility crossing the left border of array, it will fail.

​	任何越过数组左边界的可能性，都会失败。

##### Example

``` go
```

#### (*SortedStrArray) Sum

```go
func (a *SortedStrArray) Sum() (sum int)
```

Sum returns the sum of values in an array.

​	Sum 返回数组中值的总和。

##### Example

``` go
```

#### (*SortedStrArray) Unique

```go
func (a *SortedStrArray) Unique() *SortedStrArray
```

Unique uniques the array, clear repeated items.

​	唯一唯一数组，清除重复项。

##### Example

``` go
```

#### (*SortedStrArray) UnmarshalJSON

```go
func (a *SortedStrArray) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*SortedStrArray) UnmarshalValue

```go
func (a *SortedStrArray) UnmarshalValue(value interface{}) (err error)
```

UnmarshalValue is an interface implement which sets any type of value for array.

​	UnmarshalValue 是一个接口实现，用于为数组设置任何类型的值。

##### Example

``` go
```

#### (*SortedStrArray) Walk

```go
func (a *SortedStrArray) Walk(f func(value string) string) *SortedStrArray
```

Walk applies a user supplied function `f` to every item of array.

​	Walk 将用户提供的函数 `f` 应用于数组的每个项。

##### Example

``` go
```

### type StrArray

```go
type StrArray struct {
	// contains filtered or unexported fields
}
```

StrArray is a golang string array with rich features. It contains a concurrent-safe/unsafe switch, which should be set when its initialization and cannot be changed then.

​	StrArray 是一个功能丰富的 golang 字符串数组。它包含一个并发安全/不安全开关，应在初始化时设置该开关，然后无法更改。

#### func NewStrArray

```go
func NewStrArray(safe ...bool) *StrArray
```

NewStrArray creates and returns an empty array. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewStrArray 创建并返回一个空数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example

``` go
```

#### func NewStrArrayFrom

```go
func NewStrArrayFrom(array []string, safe ...bool) *StrArray
```

NewStrArrayFrom creates and returns an array with given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewStrArrayFrom 创建并返回一个具有给定切片 `array` 的数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example

``` go
```

#### func NewStrArrayFromCopy

```go
func NewStrArrayFromCopy(array []string, safe ...bool) *StrArray
```

NewStrArrayFromCopy creates and returns an array from a copy of given slice `array`. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewStrArrayFromCopy 从给定切片 `array` 的副本创建并返回数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example

``` go
```

#### func NewStrArraySize

```go
func NewStrArraySize(size int, cap int, safe ...bool) *StrArray
```

NewStrArraySize create and returns an array with given size and cap. The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.

​	NewStrArraySize 创建并返回具有给定大小和上限的数组。该参数 `safe` 用于指定是否在 concurrent-safety 中使用数组，默认为 false。

##### Example

``` go
```

#### (*StrArray) Append

```go
func (a *StrArray) Append(value ...string) *StrArray
```

Append is alias of PushRight,please See PushRight.

​	Append 是 PushRight 的别名，请参见 PushRight。

##### Example

``` go
```

#### (*StrArray) At

```go
func (a *StrArray) At(index int) (value string)
```

At returns the value by the specified index. If the given `index` is out of range of the array, it returns an empty string.

​	at 返回指定索引的值。如果给定 `index` 的超出数组的范围，则返回一个空字符串。

##### Example

``` go
```

#### (*StrArray) Chunk

```go
func (a *StrArray) Chunk(size int) [][]string
```

Chunk splits an array into multiple arrays, the size of each array is determined by `size`. The last chunk may contain less than size elements.

​	Chunk 将一个数组拆分为多个数组，每个数组的大小由 `size` 决定。最后一个块可能包含小于 size 的元素。

##### Example

``` go
```

#### (*StrArray) Clear

```go
func (a *StrArray) Clear() *StrArray
```

Clear deletes all items of current array.

​	清除将删除当前数组的所有项目。

##### Example

``` go
```

#### (*StrArray) Clone

```go
func (a *StrArray) Clone() (newArray *StrArray)
```

Clone returns a new array, which is a copy of current array.

​	克隆返回一个新数组，它是当前数组的副本。

##### Example

``` go
```

#### (*StrArray) Contains

```go
func (a *StrArray) Contains(value string) bool
```

Contains checks whether a value exists in the array.

​	包含检查数组中是否存在值。

##### Example

``` go
```

#### (*StrArray) ContainsI

```go
func (a *StrArray) ContainsI(value string) bool
```

ContainsI checks whether a value exists in the array with case-insensitively. Note that it internally iterates the whole array to do the comparison with case-insensitively.

​	ContainsI 使用不区分大小写的方式检查数组中是否存在值。请注意，它在内部迭代整个数组以不区分大小写地进行比较。

##### Example

``` go
```

#### (*StrArray) CountValues

```go
func (a *StrArray) CountValues() map[string]int
```

CountValues counts the number of occurrences of all values in the array.

​	CountValues 计算数组中所有值的出现次数。

##### Example

``` go
```

#### (*StrArray) DeepCopy

```go
func (a *StrArray) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*StrArray) Fill

```go
func (a *StrArray) Fill(startIndex int, num int, value string) error
```

Fill fills an array with num entries of the value `value`, keys starting at the `startIndex` parameter.

​	Fill 使用值 `value` 的 num 个条目填充数组，键从 `startIndex` 参数开始。

##### Example

``` go
```

#### (*StrArray) Filter

```go
func (a *StrArray) Filter(filter func(index int, value string) bool) *StrArray
```

Filter iterates array and filters elements using custom callback function. It removes the element from array if callback function `filter` returns true, it or else does nothing and continues iterating.

​	Filter 使用自定义回调函数迭代数组和筛选元素。如果回调函数 `filter` 返回 true，它会从数组中删除该元素，否则它不执行任何操作并继续迭代。

##### Example

``` go
```

#### (*StrArray) FilterEmpty

```go
func (a *StrArray) FilterEmpty() *StrArray
```

FilterEmpty removes all empty string value of the array.

​	FilterEmpty 删除数组的所有空字符串值。

##### Example

``` go
```

#### (*StrArray) Get

```go
func (a *StrArray) Get(index int) (value string, found bool)
```

Get returns the value by the specified index. If the given `index` is out of range of the array, the `found` is false.

​	Get 按指定索引返回值。如果给定 `index` 的超出数组的范围，则为 `found` false。

##### Example

``` go
```

#### (*StrArray) InsertAfter

```go
func (a *StrArray) InsertAfter(index int, values ...string) error
```

InsertAfter inserts the `values` to the back of `index`.

​	InsertAfter 将 插入 `values` 到 `index` 的后面。

##### Example

``` go
```

#### (*StrArray) InsertBefore

```go
func (a *StrArray) InsertBefore(index int, values ...string) error
```

InsertBefore inserts the `values` to the front of `index`.

​	InsertBefore 将 插入 `values` 到 `index` 的前面。

##### Example

``` go
```

#### (*StrArray) Interfaces

```go
func (a *StrArray) Interfaces() []interface{}
```

Interfaces returns current array as []interface{}.

​	Interfaces 以 []interface{} 的形式返回当前数组。

##### Example

``` go
```

#### (*StrArray) IsEmpty

```go
func (a *StrArray) IsEmpty() bool
```

IsEmpty checks whether the array is empty.

​	IsEmpty 检查数组是否为空。

##### Example

``` go
```

#### (*StrArray) Iterator

```go
func (a *StrArray) Iterator(f func(k int, v string) bool)
```

Iterator is alias of IteratorAsc.

​	Iterator 是 IteratorAsc 的别名。

##### Example

``` go
```

#### (*StrArray) IteratorAsc

```go
func (a *StrArray) IteratorAsc(f func(k int, v string) bool)
```

IteratorAsc iterates the array readonly in ascending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorAsc 使用给定的回调函数 `f` 按升序只读方式迭代数组。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*StrArray) IteratorDesc

```go
func (a *StrArray) IteratorDesc(f func(k int, v string) bool)
```

IteratorDesc iterates the array readonly in descending order with given callback function `f`. If `f` returns true, then it continues iterating; or false to stop.

​	IteratorDesc 使用给定的回调函数 `f` 按降序迭代数组只读。如果 `f` 返回 true，则继续迭代;或 false 停止。

##### Example

``` go
```

#### (*StrArray) Join

```go
func (a *StrArray) Join(glue string) string
```

Join joins array elements with a string `glue`.

​	Join 使用字符串 `glue` 连接数组元素。

##### Example

``` go
```

#### (*StrArray) Len

```go
func (a *StrArray) Len() int
```

Len returns the length of array.

​	Len 返回数组的长度。

##### Example

``` go
```

#### (*StrArray) LockFunc

```go
func (a *StrArray) LockFunc(f func(array []string)) *StrArray
```

LockFunc locks writing by callback function `f`.

​	LockFunc 通过回调函数 `f` 锁定写入。

##### Example

``` go
```

#### (StrArray) MarshalJSON

```go
func (a StrArray) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. Note that do not use pointer as its receiver here.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。请注意，此处不要使用指针作为其接收器。

##### Example

``` go
```

#### (*StrArray) Merge

```go
func (a *StrArray) Merge(array interface{}) *StrArray
```

Merge merges `array` into current array. The parameter `array` can be any garray or slice type. The difference between Merge and Append is Append supports only specified slice type, but Merge supports more parameter types.

​	合并 `array` 合并到当前数组中。该参数 `array` 可以是任何 garray 或 slice 类型。Merge 和 Append 的区别在于 Append 仅支持指定的切片类型，但 Merge 支持更多参数类型。

##### Example

``` go
```

#### (*StrArray) Pad

```go
func (a *StrArray) Pad(size int, value string) *StrArray
```

Pad pads array to the specified length with `value`. If size is positive then the array is padded on the right, or negative on the left. If the absolute value of `size` is less than or equal to the length of the array then no padding takes place.

​	焊盘垫数组设置为指定 `value` 长度。如果大小为正数，则数组在右侧填充，或在左侧填充为负数。如果 的 `size` 绝对值小于或等于数组的长度，则不会进行填充。

##### Example

``` go
```

#### (*StrArray) PopLeft

```go
func (a *StrArray) PopLeft() (value string, found bool)
```

PopLeft pops and returns an item from the beginning of array. Note that if the array is empty, the `found` is false.

​	PopLeft 弹出并返回数组开头的项。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*StrArray) PopLefts

```go
func (a *StrArray) PopLefts(size int) []string
```

PopLefts pops and returns `size` items from the beginning of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopLefts 弹出并返回 `size` 数组开头的项。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

##### Example

``` go
```

#### (*StrArray) PopRand

```go
func (a *StrArray) PopRand() (value string, found bool)
```

PopRand randomly pops and return an item out of array. Note that if the array is empty, the `found` is false.

​	PopRand 随机弹出并返回数组外的项目。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*StrArray) PopRands

```go
func (a *StrArray) PopRands(size int) []string
```

PopRands randomly pops and returns `size` items out of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopRands 随机弹出并返回 `size` 数组外的项目。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

##### Example

``` go
```

#### (*StrArray) PopRight

```go
func (a *StrArray) PopRight() (value string, found bool)
```

PopRight pops and returns an item from the end of array. Note that if the array is empty, the `found` is false.

​	PopRight 弹出并返回数组末尾的项。请注意，如果数组为空，则为 `found` false。

##### Example

``` go
```

#### (*StrArray) PopRights

```go
func (a *StrArray) PopRights(size int) []string
```

PopRights pops and returns `size` items from the end of array. If the given `size` is greater than size of the array, it returns all elements of the array. Note that if given `size` <= 0 or the array is empty, it returns nil.

​	PopRights 弹出并返回 `size` 数组末尾的项。如果给定 `size` 的大小大于数组的大小，则返回数组的所有元素。请注意，如果给定 `size` <= 0 或数组为空，则返回 nil。

##### Example

``` go
```

#### (*StrArray) PushLeft

```go
func (a *StrArray) PushLeft(value ...string) *StrArray
```

PushLeft pushes one or multiple items to the beginning of array.

​	PushLeft 将一个或多个项推送到数组的开头。

##### Example

``` go
```

#### (*StrArray) PushRight

```go
func (a *StrArray) PushRight(value ...string) *StrArray
```

PushRight pushes one or multiple items to the end of array. It equals to Append.

​	PushRight 将一个或多个项推送到数组的末尾。它等于追加。

##### Example

``` go
```

#### (*StrArray) RLockFunc

```go
func (a *StrArray) RLockFunc(f func(array []string)) *StrArray
```

RLockFunc locks reading by callback function `f`.

​	RLockFunc 通过回调函数 `f` 锁定读取。

##### Example

``` go
```

#### (*StrArray) Rand

```go
func (a *StrArray) Rand() (value string, found bool)
```

Rand randomly returns one item from array(no deleting).

​	Rand 从数组中随机返回一个项目（不删除）。

##### Example

``` go
```

#### (*StrArray) Rands

```go
func (a *StrArray) Rands(size int) []string
```

Rands randomly returns `size` items from array(no deleting).

​	Rands 从数组中随机返回 `size` 项目（不删除）。

##### Example

``` go
```

#### (*StrArray) Range

```go
func (a *StrArray) Range(start int, end ...int) []string
```

Range picks and returns items by range, like array[start:end]. Notice, if in concurrent-safe usage, it returns a copy of slice; else a pointer to the underlying data.

​	范围按范围选取和返回项目，如 array[start：end]。请注意，如果在并发安全使用中，它会返回 slice 的副本;else 指向基础数据的指针。

If `end` is negative, then the offset will start from the end of array. If `end` is omitted, then the sequence will have everything from start up until the end of the array.

​	如果 `end` 为负数，则偏移量将从数组末尾开始。如果 `end` 省略，则序列将包含从数组开始到结束的所有内容。

##### Example

``` go
```

#### (*StrArray) Remove

```go
func (a *StrArray) Remove(index int) (value string, found bool)
```

Remove removes an item by index. If the given `index` is out of range of the array, the `found` is false.

​	“删除”（Remove） 按索引删除项目。如果给定 `index` 的超出数组的范围，则为 `found` false。

##### Example

``` go
```

#### (*StrArray) RemoveValue

```go
func (a *StrArray) RemoveValue(value string) bool
```

RemoveValue removes an item by value. It returns true if value is found in the array, or else false if not found.

​	RemoveValue 按值删除项目。如果在数组中找到值，则返回 true，如果未找到，则返回 false。

##### Example

``` go
```

#### (*StrArray) RemoveValues

```go
func (a *StrArray) RemoveValues(values ...string)
```

RemoveValues removes multiple items by `values`.

​	RemoveValues 按 `values` 删除多个项目。

#### (*StrArray) Replace

```go
func (a *StrArray) Replace(array []string) *StrArray
```

Replace replaces the array items by given `array` from the beginning of array.

​	Replace 将数组项替换为从数组开头开始给定 `array` 的数组项。

##### Example

``` go
```

#### (*StrArray) Reverse

```go
func (a *StrArray) Reverse() *StrArray
```

Reverse makes array with elements in reverse order.

​	Reverse 使数组具有相反顺序的元素。

##### Example

``` go
```

#### (*StrArray) Search

```go
func (a *StrArray) Search(value string) int
```

Search searches array by `value`, returns the index of `value`, or returns -1 if not exists.

​	搜索 搜索数 `value` 组 by 返回 的 `value` 索引，如果不存在，则返回 -1。

##### Example

``` go
```

#### (*StrArray) Set

```go
func (a *StrArray) Set(index int, value string) error
```

Set sets value to specified index.

​	将设置值设置为指定索引。

##### Example

``` go
```

#### (*StrArray) SetArray

```go
func (a *StrArray) SetArray(array []string) *StrArray
```

SetArray sets the underlying slice array with the given `array`.

​	SetArray 使用给定 `array` 的 .

##### Example

``` go
```

#### (*StrArray) Shuffle

```go
func (a *StrArray) Shuffle() *StrArray
```

Shuffle randomly shuffles the array.

​	随机洗牌数组。

##### Example

``` go
```

#### (*StrArray) Slice

```go
func (a *StrArray) Slice() []string
```

Slice returns the underlying data of array. Note that, if it’s in concurrent-safe usage, it returns a copy of underlying data, or else a pointer to the underlying data.

​	Slice 返回数组的基础数据。请注意，如果它处于并发安全使用状态，它将返回基础数据的副本，或者返回指向基础数据的指针。

##### Example

``` go
```

#### (*StrArray) Sort

```go
func (a *StrArray) Sort(reverse ...bool) *StrArray
```

Sort sorts the array in increasing order. The parameter `reverse` controls whether sort in increasing order(default) or decreasing order

​	排序按升序对数组进行排序。该参数 `reverse` 控制是按递增顺序（默认）排序还是按递减顺序排序

##### Example

``` go
```

#### (*StrArray) SortFunc

```go
func (a *StrArray) SortFunc(less func(v1, v2 string) bool) *StrArray
```

SortFunc sorts the array by custom function `less`.

​	SortFunc 按自定义函数 `less` 对数组进行排序。

##### Example

``` go
```

#### (*StrArray) String

```go
func (a *StrArray) String() string
```

String returns current array as a string, which implements like json.Marshal does.

​	String 以字符串形式返回当前数组，其实现方式类似于 json。元帅做到了。

##### Example

``` go
```

#### (*StrArray) SubSlice

```go
func (a *StrArray) SubSlice(offset int, length ...int) []string
```

SubSlice returns a slice of elements from the array as specified by the `offset` and `size` parameters. If in concurrent safe usage, it returns a copy of the slice; else a pointer.

​	SubSlice 从 `offset` 数组中返回由 and `size` 参数指定的元素切片。如果处于并发安全使用状态，则返回切片的副本;else 指针。

If offset is non-negative, the sequence will start at that offset in the array. If offset is negative, the sequence will start that far from the end of the array.

​	如果偏移量为非负数，则序列将从数组中的该偏移量开始。如果偏移量为负数，则序列将从数组末尾开始。

If length is given and is positive, then the sequence will have up to that many elements in it. If the array is shorter than the length, then only the available array elements will be present. If length is given and is negative then the sequence will stop that many elements from the end of the array. If it is omitted, then the sequence will have everything from offset up until the end of the array.

​	如果给定长度并且是正数，则序列中将包含多达那么多元素。如果数组短于长度，则仅存在可用的数组元素。如果给定长度并且是负数，则序列将从数组末尾停止该许多元素。如果省略它，则序列将包含从偏移量到数组末尾的所有内容。

Any possibility crossing the left border of array, it will fail.

​	任何越过数组左边界的可能性，都会失败。

##### Example

``` go
```

#### (*StrArray) Sum

```go
func (a *StrArray) Sum() (sum int)
```

Sum returns the sum of values in an array.

​	Sum 返回数组中值的总和。

##### Example

``` go
```

#### (*StrArray) Unique

```go
func (a *StrArray) Unique() *StrArray
```

Unique uniques the array, clear repeated items. Example: [1,1,2,3,2] -> [1,2,3]

​	唯一唯一数组，清除重复项。示例：[1,1,2,3,2] -> [1,2,3]

##### Example

``` go
```

#### (*StrArray) UnmarshalJSON

```go
func (a *StrArray) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*StrArray) UnmarshalValue

```go
func (a *StrArray) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for array.

​	UnmarshalValue 是一个接口实现，用于为数组设置任何类型的值。

##### Example

``` go
```

#### (*StrArray) Walk

```go
func (a *StrArray) Walk(f func(value string) string) *StrArray
```

Walk applies a user supplied function `f` to every item of array.

​	Walk 将用户提供的函数 `f` 应用于数组的每个项。

Example Walk

​	示例步行

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	var array garray.StrArray
	tables := g.SliceStr{"user", "user_detail"}
	prefix := "gf_"
	array.Append(tables...)
	// Add prefix for given table names.
	array.Walk(func(value string) string {
		return prefix + value
	})
	fmt.Println(array.Slice())

}

Output:

[gf_user gf_user_detail]
```
