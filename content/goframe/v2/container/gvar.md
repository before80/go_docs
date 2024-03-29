+++
title = "gvar"
date = 2024-03-21T17:45:43+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gvar](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gvar)

Package gvar provides an universal variable type, like generics.

​	软件包 gvar 提供了一个通用的变量类型，如泛型。

## 常量

This section is empty.## 变量

This section is empty.## 函数

This section is empty.## 类型

### type MapOption <-2.6.0

```go
type MapOption = gconv.MapOption
```

MapOption specifies the option for map converting.

​	MapOption 指定地图转换选项。

### type Var

```go
type Var struct {
	// contains filtered or unexported fields
}
```

Var is an universal variable type implementer.

​	Var 是一个通用的变量类型实现器。

#### func New

```go
func New(value interface{}, safe ...bool) *Var
```

New creates and returns a new Var with given `value`. The optional parameter `safe` specifies whether Var is used in concurrent-safety, which is false in default.

​	new 创建并返回一个给定 `value` .可选参数 `safe` 指定是否在 concurrent-safety 中使用 Var，默认为 false。

#### (*Var) Array 

```go
func (v *Var) Array() []interface{}
```

Array is alias of Interfaces.

​	Array 是 Interfaces 的别名。

##### Example

``` go
```

#### (*Var) Bool

```go
func (v *Var) Bool() bool
```

Bool converts and returns `v` as bool.

​	Bool 转换并返回 `v` 为 bool。

##### Example

``` go
```

#### (*Var) Bytes 

```go
func (v *Var) Bytes() []byte
```

Bytes converts and returns `v` as []byte.

​	Bytes 转换并返回 `v` 为 []byte。

##### Example

``` go
```

#### (*Var) Clone 

```go
func (v *Var) Clone() *Var
```

Clone does a shallow copy of current Var and returns a pointer to this Var.

​	克隆执行当前 Var 的浅拷贝，并返回指向此 Var 的指针。

##### Example

``` go
```

#### (*Var) Copy

```go
func (v *Var) Copy() *Var
```

Copy does a deep copy of current Var and returns a pointer to this Var.

​	Copy 对当前 Var 进行深层复制，并返回指向此 Var 的指针。

#### (*Var) DeepCopy

```go
func (v *Var) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*Var) Duration 

```go
func (v *Var) Duration() time.Duration
```

Duration converts and returns `v` as time.Duration. If value of `v` is string, then it uses time.ParseDuration for conversion.

​	持续时间将转换并返回 `v` 为时间。期间。如果 value of `v` 是字符串，则它使用时间。用于转换的 ParseDuration。

##### Example

``` go
```

#### (*Var) Float32 

```go
func (v *Var) Float32() float32
```

Float32 converts and returns `v` as float32.

​	Float32 转换并返回 `v` 为 float32。

##### Example

``` go
```

#### (*Var) Float32s

```go
func (v *Var) Float32s() []float32
```

Float32s converts and returns `v` as []float32.

​	Float32s 转换并返回 `v` 为 []float32。

##### Example

``` go
```

#### (*Var) Float64

```go
func (v *Var) Float64() float64
```

Float64 converts and returns `v` as float64.

​	Float64 转换并返回 `v` 为 float64。

#### (*Var) Float64s

```go
func (v *Var) Float64s() []float64
```

Float64s converts and returns `v` as []float64.

​	Float64s 转换并返回 `v` 为 []float64。

##### Example

``` go
```

#### (*Var) Floats

```go
func (v *Var) Floats() []float64
```

Floats is alias of Float64s.

​	Floats 是 Float64s 的别名。

##### Example

``` go
```

#### (*Var) GTime

```go
func (v *Var) GTime(format ...string) *gtime.Time
```

GTime converts and returns `v` as *gtime.Time. The parameter `format` specifies the format of the time string using gtime, eg: Y-m-d H:i:s.

​	GTime 转换并返回 `v` 为 *gtime。时间。该参数使用 gtime `format` 指定时间字符串的格式，例如：Y-m-d H：i：s。

##### Example

``` go
```

#### (*Var) Int

```go
func (v *Var) Int() int
```

Int converts and returns `v` as int.

​	Int 转换并返回 `v` 为 int。

##### Example

``` go
```

#### (*Var) Int16

```go
func (v *Var) Int16() int16
```

Int16 converts and returns `v` as int16.

​	Int16 转换并返回 `v` 为 int16。

#### (*Var) Int32

```go
func (v *Var) Int32() int32
```

Int32 converts and returns `v` as int32.

​	Int32 转换并返回 `v` 为 int32。

#### (*Var) Int64

```go
func (v *Var) Int64() int64
```

Int64 converts and returns `v` as int64.

​	Int64 转换并返回 `v` 为 int64。

#### (*Var) Int64s

```go
func (v *Var) Int64s() []int64
```

Int64s converts and returns `v` as []int64.

​	Int64s 转换并返回 `v` 为 []int64。

##### Example

``` go
```

#### (*Var) Int8

```go
func (v *Var) Int8() int8
```

Int8 converts and returns `v` as int8.

​	Int8 转换并返回 `v` 为 int8。

#### (*Var) Interface

```go
func (v *Var) Interface() interface{}
```

Interface is alias of Val.

​	Interface 是 Val 的别名。

##### Example

``` go
```

#### (*Var) Interfaces

```go
func (v *Var) Interfaces() []interface{}
```

Interfaces converts and returns `v` as []interfaces{}.

​	接口转换并返回 `v` 为 []interfaces{}。

##### Example

``` go
```

#### (*Var) Ints

```go
func (v *Var) Ints() []int
```

Ints converts and returns `v` as []int.

​	Ints 转换并返回 `v` 为 []int。

##### Example

``` go
```

#### (*Var) IsEmpty

```go
func (v *Var) IsEmpty() bool
```

IsEmpty checks whether `v` is empty.

​	IsEmpty 检查是否 `v` 为空。

##### Example

``` go
```

#### (*Var) IsFloat

```go
func (v *Var) IsFloat() bool
```

IsFloat checks whether `v` is type of float.

​	IsFloat 检查是否 `v` 是浮点数的类型。

##### Example

``` go
```

#### (*Var) IsInt

```go
func (v *Var) IsInt() bool
```

IsInt checks whether `v` is type of int.

​	IsInt 检查 int 类型是否 `v` 为 int。

##### Example

``` go
```

#### (*Var) IsMap

```go
func (v *Var) IsMap() bool
```

IsMap checks whether `v` is type of map.

​	IsMap 检查是否 `v` 是 map 的类型。

##### Example

``` go
```

#### (*Var) IsNil

```go
func (v *Var) IsNil() bool
```

IsNil checks whether `v` is nil.

​	IsNil 检查是否 `v` 为 nil。

##### Example

``` go
```

#### (*Var) IsSlice

```go
func (v *Var) IsSlice() bool
```

IsSlice checks whether `v` is type of slice.

​	IsSlice 检查是否 `v` 是切片的类型。

##### Example

``` go
```

#### (*Var) IsStruct

```go
func (v *Var) IsStruct() bool
```

IsStruct checks whether `v` is type of struct.

​	IsStruct 检查是否 `v` 是结构体的类型。

##### Example

``` go
```

#### (*Var) IsUint

```go
func (v *Var) IsUint() bool
```

IsUint checks whether `v` is type of uint.

​	IsUint 检查是否 `v` 为 uint 的类型。

##### Example

``` go
```

#### (*Var) ListItemValues

```go
func (v *Var) ListItemValues(key interface{}) (values []interface{})
```

ListItemValues retrieves and returns the elements of all item struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice.

​	ListItemValues 检索并返回所有项目 struct/map 的元素，并带有键 `key` 。请注意，该参数 `list` 应为包含 map 或 struct 元素的切片类型，否则它将返回一个空切片。

##### Example

``` go
```

#### (*Var) ListItemValuesUnique

```go
func (v *Var) ListItemValuesUnique(key string) []interface{}
```

ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice.

​	ListItemValuesUnique 检索并返回所有带有键 `key` 的结构/映射的唯一元素。请注意，该参数 `list` 应为包含 map 或 struct 元素的切片类型，否则它将返回一个空切片。

##### Example

``` go
```

#### (*Var) Map

```go
func (v *Var) Map(option ...MapOption) map[string]interface{}
```

Map converts and returns `v` as map[string]interface{}.

​	Map 转换并返回 `v` 为 map[string]interface{}。

##### Example

``` go
```

#### (*Var) MapDeep

```go
func (v *Var) MapDeep(tags ...string) map[string]interface{}
```

MapDeep converts and returns `v` as map[string]interface{} recursively. Deprecated: used Map instead.

​	MapDeep 以递归方式转换并返回 `v` 为 map[string]interface{}。已弃用：改用地图。

##### Example

``` go
```

#### (*Var) MapStrAny

```go
func (v *Var) MapStrAny(option ...MapOption) map[string]interface{}
```

MapStrAny is like function Map, but implements the interface of MapStrAny.

​	MapStrAny 类似于函数 Map，但实现了 MapStrAny 的接口。

##### Example

``` go
```

#### (*Var) MapStrStr

```go
func (v *Var) MapStrStr(option ...MapOption) map[string]string
```

MapStrStr converts and returns `v` as map[string]string.

​	MapStrStr 转换并返回 `v` 为 map[string]string。

##### Example

``` go
```

#### (*Var) MapStrStrDeep

```go
func (v *Var) MapStrStrDeep(tags ...string) map[string]string
```

MapStrStrDeep converts and returns `v` as map[string]string recursively. Deprecated: used MapStrStr instead.

​	MapStrStrDeep 以递归方式转换并返回 `v` 为 map[string]string。已弃用：改用 MapStrStr。

##### Example

``` go
```

#### (*Var) MapStrVar

```go
func (v *Var) MapStrVar(option ...MapOption) map[string]*Var
```

MapStrVar converts and returns `v` as map[string]Var.

​	MapStrVar 转换并返回 `v` 为 map[string]Var。

##### Example

``` go
```

#### (*Var) MapStrVarDeep

```go
func (v *Var) MapStrVarDeep(tags ...string) map[string]*Var
```

MapStrVarDeep converts and returns `v` as map[string]*Var recursively. Deprecated: used MapStrVar instead.

​	MapStrVarDeep 以递归方式转换并返回 `v` 为 map[string]*Var。已弃用：改用 MapStrVar。

##### Example

``` go
```

#### (*Var) MapToMap

```go
func (v *Var) MapToMap(pointer interface{}, mapping ...map[string]string) (err error)
```

MapToMap converts any map type variable `params` to another map type variable `pointer`. See gconv.MapToMap.

​	MapToMap 将任何地图类型变量 `params` 转换为另一个地图类型变量 `pointer` 。请参见 gconv。MapToMap的。

##### Example

``` go
```

#### (*Var) MapToMaps

```go
func (v *Var) MapToMaps(pointer interface{}, mapping ...map[string]string) (err error)
```

MapToMaps converts any map type variable `params` to another map type variable `pointer`. See gconv.MapToMaps.

​	MapToMaps将任何地图类型变量 `params` 转换为另一个地图类型变量 `pointer` 。请参见 gconv。MapToMaps。

##### Example

``` go
```

#### (*Var) MapToMapsDeep

```go
func (v *Var) MapToMapsDeep(pointer interface{}, mapping ...map[string]string) (err error)
```

MapToMapsDeep converts any map type variable `params` to another map type variable `pointer` recursively. See gconv.MapToMapsDeep.

​	MapToMapsDeep 以递归方式将任何地图类型变量 `params` 转换为另一个地图类型变量 `pointer` 。请参见 gconv。MapToMapsDeep。

##### Example

``` go
```

#### (*Var) Maps

```go
func (v *Var) Maps(option ...MapOption) []map[string]interface{}
```

Maps converts and returns `v` as map[string]string. See gconv.Maps.

​	Maps 转换并返回 `v` 为 map[string]string。请参见 gconv。地图。

##### Example

``` go
```

#### (*Var) MapsDeep

```go
func (v *Var) MapsDeep(tags ...string) []map[string]interface{}
```

MapsDeep converts `value` to []map[string]interface{} recursively. Deprecated: used Maps instead.

​	MapsDeep 以递归方式 `value` 转换为 []map[string]interface{}。已弃用：改用地图。

##### Example

``` go
```

#### (Var) MarshalJSON

```go
func (v Var) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*Var) Scan

```go
func (v *Var) Scan(pointer interface{}, mapping ...map[string]string) error
```

Scan automatically checks the type of `pointer` and converts `params` to `pointer`. It supports `pointer` with type of `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` for converting.

​	扫描会自动检查 的 `pointer` 类型并转换为 `params` `pointer` 。它支持 `pointer` `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` 用于转换的类型。

See gconv.Scan.

​	请参阅 gconv.Scan。

##### Example

``` go
```

#### (*Var) Set

```go
func (v *Var) Set(value interface{}) (old interface{})
```

Set sets `value` to `v`, and returns the old value.

​	将 sets `value` 设置为 `v` ，并返回旧值。

##### Example

``` go
```

#### (*Var) Slice

```go
func (v *Var) Slice() []interface{}
```

Slice is alias of Interfaces.

​	Slice 是 Interfaces 的别名。

##### Example

``` go
```

#### (*Var) String

```go
func (v *Var) String() string
```

String converts and returns `v` as string.

​	String 转换并返回 `v` 为字符串。

##### Example

``` go
```

#### (*Var) Strings

```go
func (v *Var) Strings() []string
```

Strings converts and returns `v` as []string.

​	Strings 转换并返回 `v` 为 []string。

##### Example

``` go
```

#### (*Var) Struct

```go
func (v *Var) Struct(pointer interface{}, mapping ...map[string]string) error
```

Struct maps value of `v` to `pointer`. The parameter `pointer` should be a pointer to a struct instance. The parameter `mapping` is used to specify the key-to-attribute mapping rules.

​	结构映射值为 `v` `pointer` 。该参数 `pointer` 应是指向结构实例的指针。该参数 `mapping` 用于指定键到属性的映射规则。

##### Example

``` go
```

#### (*Var) Structs

```go
func (v *Var) Structs(pointer interface{}, mapping ...map[string]string) error
```

Structs converts and returns `v` as given struct slice.

​	structs 转换并返回 `v` 给定的结构切片。

##### Example

``` go
```

#### (*Var) Time

```go
func (v *Var) Time(format ...string) time.Time
```

Time converts and returns `v` as time.Time. The parameter `format` specifies the format of the time string using gtime, eg: Y-m-d H:i:s.

​	时间转换并返回 `v` 为时间。时间。该参数使用 gtime `format` 指定时间字符串的格式，例如：Y-m-d H：i：s。

##### Example

``` go
```

#### (*Var) Uint

```go
func (v *Var) Uint() uint
```

Uint converts and returns `v` as uint.

​	Uint 转换并返回 `v` 为 uint。

##### Example

``` go
```

#### (*Var) Uint16

```go
func (v *Var) Uint16() uint16
```

Uint16 converts and returns `v` as uint16.

​	Uint16 转换并返回 `v` 为 uint16。

#### (*Var) Uint32

```go
func (v *Var) Uint32() uint32
```

Uint32 converts and returns `v` as uint32.

​	Uint32 转换并返回 `v` 为 uint32。

#### (*Var) Uint64

```go
func (v *Var) Uint64() uint64
```

Uint64 converts and returns `v` as uint64.

​	Uint64 转换并返回 `v` 为 uint64。

#### (*Var) Uint64s

```go
func (v *Var) Uint64s() []uint64
```

Uint64s converts and returns `v` as []uint64.

​	Uint64s 转换并返回 `v` 为 []uint64。

##### Example

``` go
```

#### (*Var) Uint8

```go
func (v *Var) Uint8() uint8
```

Uint8 converts and returns `v` as uint8.

​	Uint8 转换并返回 `v` 为 uint8。

#### (*Var) Uints

```go
func (v *Var) Uints() []uint
```

Uints converts and returns `v` as []uint.

​	Uints 转换并返回 `v` 为 []uint。

##### Example

``` go
```

#### (*Var) UnmarshalJSON

```go
func (v *Var) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*Var) UnmarshalValue

```go
func (v *Var) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for Var.

​	UnmarshalValue 是一个接口实现，用于为 Var 设置任何类型的值。

##### Example

``` go
```

#### (*Var) Val

```go
func (v *Var) Val() interface{}
```

Val returns the current value of `v`.

​	Val 返回 的 `v` 当前值。

##### Example

``` go
```

#### (*Var) Vars

```go
func (v *Var) Vars() []*Var
```

Vars converts and returns `v` as []Var.

​	Vars 转换并返回 `v` 为 []Var。

##### Example

``` go
```

### type Vars <-2.3.0

```go
type Vars []*Var
```

Vars is a slice of *Var.

​	Vars 是 *Var 的切片。

#### (Vars) Float32s <-2.3.0

```go
func (vs Vars) Float32s() (s []float32)
```

Float32s converts and returns `vs` as []float32.

​	Float32s 转换并返回 `vs` 为 []float32。

#### (Vars) Float64s <-2.3.0

```go
func (vs Vars) Float64s() (s []float64)
```

Float64s converts and returns `vs` as []float64.

​	Float64s 转换并返回 `vs` 为 []float64。

#### (Vars) Int16s <-2.3.0

```go
func (vs Vars) Int16s() (s []int16)
```

Int16s converts and returns `vs` as []int16.

​	Int16s 转换并返回 `vs` 为 []int16。

#### (Vars) Int32s <-2.3.0

```go
func (vs Vars) Int32s() (s []int32)
```

Int32s converts and returns `vs` as []int32.

​	Int32s 转换并返回 `vs` 为 []int32。

#### (Vars) Int64s <-2.3.0

```go
func (vs Vars) Int64s() (s []int64)
```

Int64s converts and returns `vs` as []int64.

​	Int64s 转换并返回 `vs` 为 []int64。

#### (Vars) Int8s <-2.3.0

```go
func (vs Vars) Int8s() (s []int8)
```

Int8s converts and returns `vs` as []int8.

​	Int8s 转换并返回 `vs` 为 []int8。

#### (Vars) Interfaces

```go
func (vs Vars) Interfaces() (s []interface{})
```

Interfaces converts and returns `vs` as []interface{}.

​	接口转换并返回 `vs` 为 []interface{}。

#### (Vars) Ints

```go
func (vs Vars) Ints() (s []int)
```

Ints converts and returns `vs` as []Int.

​	Ints 转换并返回 `vs` 为 []Int。

#### (Vars) Scan

```go
func (vs Vars) Scan(pointer interface{}, mapping ...map[string]string) error
```

Scan converts `vs` to []struct/[]*struct.

​	扫描转换为 `vs` []struct/[]*struct。

#### (Vars) Strings

```go
func (vs Vars) Strings() (s []string)
```

Strings converts and returns `vs` as []string.

​	Strings 转换并返回 `vs` 为 []string。

#### (Vars) Uint16s <-2.3.0

```go
func (vs Vars) Uint16s() (s []uint16)
```

Uint16s converts and returns `vs` as []uint16.

​	Uint16s 转换并返回 `vs` 为 []uint16。

#### (Vars) Uint32s <-2.3.0

```go
func (vs Vars) Uint32s() (s []uint32)
```

Uint32s converts and returns `vs` as []uint32.

​	Uint32s 转换并返回 `vs` 为 []uint32。

#### (Vars) Uint64s <-2.3.0

```go
func (vs Vars) Uint64s() (s []uint64)
```

Uint64s converts and returns `vs` as []uint64.

​	Uint64s 转换并返回 `vs` 为 []uint64。

#### (Vars) Uint8s <-2.3.0

```go
func (vs Vars) Uint8s() (s []uint8)
```

Uint8s converts and returns `vs` as []uint8.

​	Uint8s 转换并返回 `vs` 为 []uint8。

#### (Vars) Uints

```go
func (vs Vars) Uints() (s []uint)
```

Uints converts and returns `vs` as []uint.

​	Uints 转换并返回 `vs` 为 []uint。