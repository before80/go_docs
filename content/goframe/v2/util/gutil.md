+++
title = "gutil"
date = 2024-03-21T18:00:04+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gutil](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gutil)

### Overview 概述

Package gutil provides utility functions.

​	软件包 gutil 提供实用函数。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func ComparatorByte

```go
func ComparatorByte(a, b interface{}) int
```

ComparatorByte provides a basic comparison on byte.

​	ComparatorByte 提供了对字节的基本比较。

#### func ComparatorFloat32

```go
func ComparatorFloat32(a, b interface{}) int
```

ComparatorFloat32 provides a basic comparison on float32.

​	ComparatorFloat32 提供了对 float32 的基本比较。

#### func ComparatorFloat64

```go
func ComparatorFloat64(a, b interface{}) int
```

ComparatorFloat64 provides a basic comparison on float64.

​	ComparatorFloat64 提供了对 float64 的基本比较。

#### func ComparatorInt

```go
func ComparatorInt(a, b interface{}) int
```

ComparatorInt provides a basic comparison on int.

​	ComparatorInt 提供了对 int 的基本比较。

#### func ComparatorInt16

```go
func ComparatorInt16(a, b interface{}) int
```

ComparatorInt16 provides a basic comparison on int16.

​	ComparatorInt16 提供了对 int16 的基本比较。

#### func ComparatorInt32

```go
func ComparatorInt32(a, b interface{}) int
```

ComparatorInt32 provides a basic comparison on int32.

​	ComparatorInt32 提供了对 int32 的基本比较。

#### func ComparatorInt64

```go
func ComparatorInt64(a, b interface{}) int
```

ComparatorInt64 provides a basic comparison on int64.

​	ComparatorInt64 提供了对 int64 的基本比较。

#### func ComparatorInt8

```go
func ComparatorInt8(a, b interface{}) int
```

ComparatorInt8 provides a basic comparison on int8.

​	ComparatorInt8 提供了对 int8 的基本比较。

#### func ComparatorRune

```go
func ComparatorRune(a, b interface{}) int
```

ComparatorRune provides a basic comparison on rune.

​	ComparatorRune 提供了对符文的基本比较。

#### func ComparatorString

```go
func ComparatorString(a, b interface{}) int
```

ComparatorString provides a fast comparison on strings.

​	ComparatorString 提供字符串的快速比较。

#### func ComparatorTime

```go
func ComparatorTime(a, b interface{}) int
```

ComparatorTime provides a basic comparison on time.Time.

​	ComparatorTime 提供了基本的时间比较。时间。

#### func ComparatorUint

```go
func ComparatorUint(a, b interface{}) int
```

ComparatorUint provides a basic comparison on uint.

​	ComparatorUint 提供了对 uint 的基本比较。

#### func ComparatorUint16

```go
func ComparatorUint16(a, b interface{}) int
```

ComparatorUint16 provides a basic comparison on uint16.

​	ComparatorUint16 提供了对 uint16 的基本比较。

#### func ComparatorUint32

```go
func ComparatorUint32(a, b interface{}) int
```

ComparatorUint32 provides a basic comparison on uint32.

​	ComparatorUint32 提供了对 uint32 的基本比较。

#### func ComparatorUint64

```go
func ComparatorUint64(a, b interface{}) int
```

ComparatorUint64 provides a basic comparison on uint64.

​	ComparatorUint64 提供了对 uint64 的基本比较。

#### func ComparatorUint8

```go
func ComparatorUint8(a, b interface{}) int
```

ComparatorUint8 provides a basic comparison on uint8.

​	ComparatorUint8 提供了对 uint8 的基本比较。

#### func Copy <-2.1.0

```go
func Copy(src interface{}) (dst interface{})
```

Copy returns a deep copy of v.

​	Copy 返回 v 的深层副本。

Copy is unable to copy unexported fields in a struct (lowercase field names). Unexported fields can’t be reflected by the Go runtime and therefore they can’t perform any data copies.

​	Copy 无法复制结构中未导出的字段（小写字段名称）。Go 运行时无法反映未导出的字段，因此它们无法执行任何数据复制。

#### func Dump

```go
func Dump(values ...interface{})
```

Dump prints variables `values` to stdout with more manually readable.

​	转储将变量打印 `values` 到 stdout，具有更多的手动可读性。

#### func DumpJson <-2.4.2

```go
func DumpJson(jsonContent string)
```

DumpJson pretty dumps json content to stdout.

​	DumpJson 漂亮地将 json 内容转储到 stdout。

#### func DumpTo

```go
func DumpTo(writer io.Writer, value interface{}, option DumpOption)
```

DumpTo writes variables `values` as a string in to `writer` with more manually readable

​	DumpTo 将 `values` 变量作为字符串写入 to `writer` ，具有更多的手动可读性

#### func DumpWithOption

```go
func DumpWithOption(value interface{}, option DumpOption)
```

DumpWithOption returns variables `values` as a string with more manually readable.

​	DumpWithOption `values` 以字符串形式返回变量，具有更多的手动可读性。

#### func DumpWithType

```go
func DumpWithType(values ...interface{})
```

DumpWithType acts like Dump, but with type information. Also see Dump.

​	DumpWithType 的行为类似于 Dump，但具有类型信息。另请参阅转储。

#### func FillStructWithDefault <-2.5.4

```go
func FillStructWithDefault(structPtr interface{}) error
```

FillStructWithDefault fills attributes of pointed struct with tag value from `default/d` tag . The parameter `structPtr` should be either type of *struct/[]*struct.

​	FillStructWithDefault 用 tag 中的 `default/d` 标记值填充指向结构的属性。参数 `structPtr` 应为 *struct/[]*struct 类型之一。

#### func GetOrDefaultAny <-2.2.0

```go
func GetOrDefaultAny(def interface{}, param ...interface{}) interface{}
```

GetOrDefaultAny checks and returns value according whether parameter `param` available. It returns `param[0]` if it is available, or else it returns `def`.

​	GetOrDefaultAny 根据参数 `param` 是否可用检查并返回值。如果它可用，它将返回 `param[0]` ，否则它将返回 `def` 。

#### func GetOrDefaultStr <-2.2.0

```go
func GetOrDefaultStr(def string, param ...string) string
```

GetOrDefaultStr checks and returns value according whether parameter `param` available. It returns `param[0]` if it is available, or else it returns `def`.

​	GetOrDefaultStr 根据参数 `param` 是否可用检查并返回值。如果它可用，它将返回 `param[0]` ，否则它将返回 `def` 。

#### func Go <-2.5.3

```go
func Go(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
	recoverFunc func(ctx context.Context, exception error),
)
```

Go creates a new asynchronous goroutine function with specified recover function.

​	Go 使用指定的 recover 函数创建一个新的异步 goroutine 函数。

The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`. If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will throw to parent goroutine.

​	当在执行 期间 `goroutineFunc` 出现任何恐慌时， `recoverFunc` 将调用该参数。如果 `recoverFunc` 给定 nil，它会忽略来自 `goroutineFunc` 的恐慌，并且不会向父 goroutine 抛出恐慌。

But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.

​	但是，请注意，如果 `recoverFunc` 也引发恐慌，这种恐慌将被抛给父 goroutine。

#### func IsEmpty

```go
func IsEmpty(value interface{}) bool
```

IsEmpty checks given `value` empty or not. It returns false if `value` is: integer(0), bool(false), slice/map(len=0), nil; or else returns true.

​	IsEmpty 检查是否为 `value` 空。如果 `value` 为 integer（0）， bool（false）， slice/map（len=0）， nil，则返回 false;否则返回 true。

#### func IsTypeOf <-2.6.4

```go
func IsTypeOf(value, valueInExpectType interface{}) bool
```

IsTypeOf checks and returns whether the type of `value` and `valueInExpectType` equal.

​	IsTypeOf 检查并返回 和 `value` `valueInExpectType` 的类型是否相等。

#### func ItemValue

```go
func ItemValue(item interface{}, key interface{}) (value interface{}, found bool)
```

ItemValue retrieves and returns its value of which name/attribute specified by `key`. The parameter `item` can be type of map/*map/struct/*struct.

​	ItemValue 检索并返回其由 `key` 指定的名称/属性的值。参数 `item` 类型可以是 map/*map/struct/*struct。

#### func Keys

```go
func Keys(mapOrStruct interface{}) (keysOrAttrs []string)
```

Keys retrieves and returns the keys from given map or struct.

​	键从给定的映射或结构中检索并返回键。

#### func ListItemValues

```go
func ListItemValues(list interface{}, key interface{}, subKey ...interface{}) (values []interface{})
```

ListItemValues retrieves and returns the elements of all item struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice.

​	ListItemValues 检索并返回所有项目 struct/map 的元素，并带有键 `key` 。请注意，该参数 `list` 应为包含 map 或 struct 元素的切片类型，否则它将返回一个空切片。

The parameter `list` supports types like: []map[string]interface{} []map[string]sub-map []struct []struct:sub-struct Note that the sub-map/sub-struct makes sense only if the optional parameter `subKey` is given.

​	该参数 `list` 支持以下类型： []map[string]interface{} []map[string]sub-map []struct []struct：sub-struct 请注意，仅当给定可选参数 `subKey` 时，sub-map/sub-struct 才有意义。

#### func ListItemValuesUnique

```go
func ListItemValuesUnique(list interface{}, key string, subKey ...interface{}) []interface{}
```

ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice.

​	ListItemValuesUnique 检索并返回所有带有键 `key` 的结构/映射的唯一元素。请注意，该参数 `list` 应为包含 map 或 struct 元素的切片类型，否则它将返回一个空切片。

#### func ListToMapByKey

```go
func ListToMapByKey(list []map[string]interface{}, key string) map[string]interface{}
```

ListToMapByKey converts `list` to a map[string]interface{} of which key is specified by `key`. Note that the item value may be type of slice.

​	ListToMapByKey `list` 转换为 map[string]interface{}，其键由 `key` 指定。请注意，项目值可以是切片类型。

#### func MapContains

```go
func MapContains(data map[string]interface{}, key string) (ok bool)
```

MapContains checks whether map `data` contains `key`.

​	MapContains 检查 map `data` 是否包含 `key` 。

#### func MapContainsPossibleKey

```go
func MapContainsPossibleKey(data map[string]interface{}, key string) bool
```

MapContainsPossibleKey checks if the given `key` is contained in given map `data`. It checks the key ignoring cases and symbols.

​	MapContainsPossibleKey 检查给定的地图中是否包含给定 `key` 的地图 `data` 。它检查键，忽略大小写和符号。

Note that this function might be of low performance.

​	请注意，此函数的性能可能较低。

#### func MapCopy

```go
func MapCopy(data map[string]interface{}) (copy map[string]interface{})
```

MapCopy does a shallow copy from map `data` to `copy` for most commonly used map type map[string]interface{}.

​	MapCopy 对最常用的地图类型 map[string]interface{} 执行从 map 到 `copy` 的 `data` 浅层复制。

#### func MapDelete

```go
func MapDelete(data map[string]interface{}, keys ...string)
```

MapDelete deletes all `keys` from map `data`.

​	MapDelete 从 map `data` 中删除所有 `keys` 。

#### func MapMerge

```go
func MapMerge(dst map[string]interface{}, src ...map[string]interface{})
```

MapMerge merges all map from `src` to map `dst`.

​	MapMerge 将所有 map 合并到 `src` map `dst` 。

#### func MapMergeCopy

```go
func MapMergeCopy(src ...map[string]interface{}) (copy map[string]interface{})
```

MapMergeCopy creates and returns a new map which merges all map from `src`.

​	MapMergeCopy 创建并返回一个新地图，该地图合并了 中的所有 `src` 地图。

#### func MapOmitEmpty

```go
func MapOmitEmpty(data map[string]interface{})
```

MapOmitEmpty deletes all empty values from given map.

​	MapOmitEmpty 从给定地图中删除所有空值。

#### func MapPossibleItemByKey

```go
func MapPossibleItemByKey(data map[string]interface{}, key string) (foundKey string, foundValue interface{})
```

MapPossibleItemByKey tries to find the possible key-value pair for given key ignoring cases and symbols.

​	MapPossibleItemByKey 尝试为给定的键查找可能的键值对，忽略大小写和符号。

Note that this function might be of low performance.

​	请注意，此函数的性能可能较低。

#### func MapToSlice

```go
func MapToSlice(data interface{}) []interface{}
```

MapToSlice converts map to slice of which all keys and values are its items. Eg: {“K1”: “v1”, “K2”: “v2”} => [“K1”, “v1”, “K2”, “v2”]

​	MapToSlice 将 map 转换为 slice，其中所有键和值都是其项。例如： {“K1”： “v1”， “K2”： “v2”} => [“K1”， “v1”， “K2”， “v2”]

#### func SliceCopy

```go
func SliceCopy(slice []interface{}) []interface{}
```

SliceCopy does a shallow copy of slice `data` for most commonly used slice type []interface{}.

​	SliceCopy 对最常用的切片类型 []interface{} 执行切片 `data` 的浅拷贝。

#### func SliceDelete

```go
func SliceDelete(slice []interface{}, index int) (newSlice []interface{})
```

SliceDelete deletes an element at `index` and returns the new slice. It does nothing if the given `index` is invalid.

​	SliceDelete 删除元素 `index` 并返回新切片。如果给定 `index` 无效，则它不执行任何操作。

#### func SliceInsertAfter <-2.3.2

```go
func SliceInsertAfter(slice []interface{}, index int, values ...interface{}) (newSlice []interface{})
```

SliceInsertAfter inserts the `values` to the back of `index` and returns a new slice.

​	SliceInsertAfter 将 插入 `values` 到切 `index` 片的背面并返回一个新切片。

##### Example

``` go
```

#### func SliceInsertBefore <-2.3.2

```go
func SliceInsertBefore(slice []interface{}, index int, values ...interface{}) (newSlice []interface{})
```

SliceInsertBefore inserts the `values` to the front of `index` and returns a new slice.

​	SliceInsertBefore 将 插入 `values` 到前面 `index` 并返回一个新切片。

##### Example

``` go
```

#### func SliceToMap

```go
func SliceToMap(slice interface{}) map[string]interface{}
```

SliceToMap converts slice type variable `slice` to `map[string]interface{}`. Note that if the length of `slice` is not an even number, it returns nil. Eg: [“K1”, “v1”, “K2”, “v2”] => {“K1”: “v1”, “K2”: “v2”} [“K1”, “v1”, “K2”] => nil

​	SliceToMap 将切片类型变量 `slice` 转换为 `map[string]interface{}` 。请注意，如果 的 `slice` 长度不是偶数，则返回 nil。例如： [“K1”， “v1”， “K2”， “v2”] => {“K1”： “v1”， “K2”： “v2”} [“K1”， “v1”， “K2”] => nil

#### func SliceToMapWithColumnAsKey

```go
func SliceToMapWithColumnAsKey(slice interface{}, key interface{}) map[interface{}]interface{}
```

SliceToMapWithColumnAsKey converts slice type variable `slice` to `map[interface{}]interface{}` The value of specified column use as the key for returned map. Eg: SliceToMapWithColumnAsKey([{“K1”: “v1”, “K2”: 1}, {“K1”: “v2”, “K2”: 2}], “K1”) => {“v1”: {“K1”: “v1”, “K2”: 1}, “v2”: {“K1”: “v2”, “K2”: 2}} SliceToMapWithColumnAsKey([{“K1”: “v1”, “K2”: 1}, {“K1”: “v2”, “K2”: 2}], “K2”) => {1: {“K1”: “v1”, “K2”: 1}, 2: {“K1”: “v2”, “K2”: 2}}

​	SliceToMapWithColumnAsKey 将切片类型变量 `slice` 转换为 `map[interface{}]interface{}` 指定列的值，用作返回映射的键。例如： SliceToMapWithColumnAsKey（[{“K1”： “v1”， “K2”： 1}， {“K1”： “v2”， “K2”： 2}]， “K1”） => {“v1”： {“K1”： “v1”， “K2”： 1}， “v2”： {“K1”： “v2”， “K2”： 2}} SliceToMapWithColumnAsKey（[{“K1”： “v1”， “K2”： 1}， {“K1”： “v2”， “K2”： 2}]， “K2”） => {1： {“K1”： “v1”， “K2”： 1}， 2： {“K1”： “v2”， “K2”： 2}}

#### func StructToSlice

```go
func StructToSlice(data interface{}) []interface{}
```

StructToSlice converts struct to slice of which all keys and values are its items. Eg: {“K1”: “v1”, “K2”: “v2”} => [“K1”, “v1”, “K2”, “v2”]

​	StructToSlice 将 struct 转换为切片，其中所有键和值都是其项。例如： {“K1”： “v1”， “K2”： “v2”}=> [“K1”， “v1”， “K2”， “v2”]

#### func Throw

```go
func Throw(exception interface{})
```

Throw throws out an exception, which can be caught be TryCatch or recover.

​	Throw 抛出一个异常，可以被 TryCatch 捕获或恢复。

#### func Try

```go
func Try(ctx context.Context, try func(ctx context.Context)) (err error)
```

Try implements try… logistics using internal panic…recover. It returns error if any exception occurs, or else it returns nil.

​	尝试实现尝试...物流使用内部恐慌...恢复。如果发生任何异常，它将返回错误，否则返回 nil。

#### func TryCatch

```go
func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error))
```

TryCatch implements `try...catch..`. logistics using internal `panic...recover`. It automatically calls function `catch` if any exception occurs and passes the exception as an error. If `catch` is given nil, it ignores the panic from `try` and no panic will throw to parent goroutine.

​	TryCatch 实现 `try...catch..` .物流使用内部 `panic...recover` .如果发生任何异常，它会自动调用函数 `catch` ，并将异常作为错误传递。如果 `catch` 给定 nil，它会忽略来自 `try` 的恐慌，并且不会向父 goroutine 抛出恐慌。

But, note that, if function `catch` also throws panic, the current goroutine will panic.

​	但是，请注意，如果函数 `catch` 也引发恐慌，则当前的 goroutine 将恐慌。

#### func Values

```go
func Values(mapOrStruct interface{}) (values []interface{})
```

Values retrieves and returns the values from given map or struct.

​	Values 从给定的映射或结构中检索并返回值。

## 类型

### type Comparator

```go
type Comparator func(a, b interface{}) int
```

Comparator is a function that compare a and b, and returns the result as int.

​	Comparator 是一个比较 a 和 b 的函数，并将结果返回为 int。

Should return a number:

​	应返回一个数字：

```
negative , if a < b
zero     , if a == b
positive , if a > b
```

### type DumpOption

```go
type DumpOption struct {
	WithType     bool // WithType specifies dumping content with type information.
	ExportedOnly bool // Only dump Exported fields for structs.
}
```

DumpOption specifies the behavior of function Export.

​	DumpOption 指定函数 Export 的行为。

### type OriginTypeAndKindOutput <-2.3.0

```go
type OriginTypeAndKindOutput = reflection.OriginTypeAndKindOutput
```

#### func OriginTypeAndKind <-2.3.0

```go
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput)
```

OriginTypeAndKind retrieves and returns the original reflect type and kind.

​	OriginTypeAndKind 检索并返回原始反射类型和种类。

### type OriginValueAndKindOutput <-2.3.0

```go
type OriginValueAndKindOutput = reflection.OriginValueAndKindOutput
```

#### func OriginValueAndKind <-2.3.0

```go
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput)
```

OriginValueAndKind retrieves and returns the original reflect value and kind.

​	OriginValueAndKind 检索并返回原始的 reflect 值和种类。