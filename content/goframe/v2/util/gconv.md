+++
title = "gconv"
date = 2024-03-21T17:59:20+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gconv](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gconv)

Package gconv implements powerful and convenient converting functionality for any types of variables.

​	软件包 gconv 为任何类型的变量实现了强大而方便的转换功能。

This package should keep much less dependencies with other packages.

​	此包应保留与其他包的依赖性要少得多。

## 常量

This section is empty.

## 变量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/util/gconv/gconv.go#L29)

```go
var (

	// StructTagPriority defines the default priority tags for Map*/Struct* functions.
	// Note that, the `gconv/param` tags are used by old version of package.
	// It is strongly recommended using short tag `c/p` instead in the future.
	StructTagPriority = gtag.StructTagPriority
)
```

## 函数

#### func Bool

```go
func Bool(any interface{}) bool
```

Bool converts `any` to bool. It returns false if `any` is: false, “”, 0, “false”, “off”, “no”, empty slice/map.

​	布尔转换为 `any` 布尔。如果 `any` 为以下情况，则返回 false：false、“”、0、“false”、“off”、“no”、空切片/映射。

#### func Byte

```go
func Byte(any interface{}) byte
```

Byte converts `any` to byte.

​	Byte `any` 转换为 byte。

#### func Bytes

```go
func Bytes(any interface{}) []byte
```

Bytes converts `any` to []byte.

​	字节转换为 `any` []byte。

#### func Convert

```go
func Convert(fromValue interface{}, toTypeName string, extraParams ...interface{}) interface{}
```

Convert converts the variable `fromValue` to the type `toTypeName`, the type `toTypeName` is specified by string.

​	Convert 将变量 `fromValue` 转换为类型 `toTypeName` ，类型 `toTypeName` 由字符串指定。

The optional parameter `extraParams` is used for additional necessary parameter for this conversion. It supports common basic types conversion as its conversion based on type name string.

​	可选参数 `extraParams` 用于此转换的其他必要参数。它支持常见的基本类型转换，因为它基于类型名称字符串进行转换。

#### func ConvertWithRefer <-2.5.4

```go
func ConvertWithRefer(fromValue interface{}, referValue interface{}, extraParams ...interface{}) interface{}
```

ConvertWithRefer converts the variable `fromValue` to the type referred by value `referValue`.

​	ConvertWithRefer 将变量 `fromValue` 转换为 value `referValue` 所引用的类型。

The optional parameter `extraParams` is used for additional necessary parameter for this conversion. It supports common basic types conversion as its conversion based on type name string.

​	可选参数 `extraParams` 用于此转换的其他必要参数。它支持常见的基本类型转换，因为它基于类型名称字符串进行转换。

#### func Duration

```go
func Duration(any interface{}) time.Duration
```

Duration converts `any` to time.Duration. If `any` is string, then it uses time.ParseDuration to convert it. If `any` is numeric, then it converts `any` as nanoseconds.

​	持续时间 `any` 转换为时间。期间。如果 `any` 是字符串，则它使用时间。ParseDuration 进行转换。如果 `any` 是数字，则转换为 `any` 纳秒。

#### func Float32

```go
func Float32(any interface{}) float32
```

Float32 converts `any` to float32.

​	Float32 `any` 转换为 float32。

#### func Float32s

```go
func Float32s(any interface{}) []float32
```

Float32s converts `any` to []float32.

​	Float32s `any` 转换为 []float32。

#### func Float64

```go
func Float64(any interface{}) float64
```

Float64 converts `any` to float64.

​	Float64 `any` 转换为 float64。

#### func Float64s

```go
func Float64s(any interface{}) []float64
```

Float64s converts `any` to []float64.

​	Float64s `any` 转换为 []float64。

#### func Floats

```go
func Floats(any interface{}) []float64
```

Floats converts `any` to []float64.

​	浮点数转换为 `any` []float64。

#### func GTime

```go
func GTime(any interface{}, format ...string) *gtime.Time
```

GTime converts `any` to *gtime.Time. The parameter `format` can be used to specify the format of `any`. It returns the converted value that matched the first format of the formats slice. If no `format` given, it converts `any` using gtime.NewFromTimeStamp if `any` is numeric, or using gtime.StrToTime if `any` is string.

​	GTime 转换为 `any` *gtime。时间。该参数 `format` 可用于指定 的 `any` 格式。它返回与格式切片的第一个格式匹配的转换值。如果没有 `format` 给出，它将 `any` 使用 gtime 进行转换。NewFromTimeStamp 如果 `any` 为数值，则使用 gtime。如果 StrToTime `any` 是字符串。

#### func Int

```go
func Int(any interface{}) int
```

Int converts `any` to int.

​	Int `any` 转换为 int。

#### func Int16

```go
func Int16(any interface{}) int16
```

Int16 converts `any` to int16.

​	Int16 `any` 转换为 int16。

#### func Int32

```go
func Int32(any interface{}) int32
```

Int32 converts `any` to int32.

​	Int32 转换为 `any` int32。

#### func Int32s

```go
func Int32s(any interface{}) []int32
```

Int32s converts `any` to []int32.

​	Int32s `any` 转换为 []int32。

#### func Int64

```go
func Int64(any interface{}) int64
```

Int64 converts `any` to int64.

​	Int64 转换为 `any` int64。

#### func Int64s

```go
func Int64s(any interface{}) []int64
```

Int64s converts `any` to []int64.

​	Int64s `any` 转换为 []int64。

#### func Int8

```go
func Int8(any interface{}) int8
```

Int8 converts `any` to int8.

​	Int8 `any` 转换为 int8。

#### func Interfaces

```go
func Interfaces(any interface{}) []interface{}
```

Interfaces converts `any` to []interface{}.

​	接口转换为 `any` []interface{}。

#### func Ints

```go
func Ints(any interface{}) []int
```

Ints converts `any` to []int.

​	Ints `any` 转换为 []int。

#### func Map

```go
func Map(value interface{}, option ...MapOption) map[string]interface{}
```

Map converts any variable `value` to map[string]interface{}. If the parameter `value` is not a map/struct/*struct type, then the conversion will fail and returns nil.

​	Map 将任何变量 `value` 转换为 map[string]interface{}。如果参数 `value` 不是 map/struct/*struct 类型，则转换将失败并返回 nil。

If `value` is a struct/*struct object, the second parameter `tags` specifies the most priority tags that will be detected, otherwise it detects the tags in order of: gconv, json, field name.

​	如果 `value` 是 struct/*struct 对象，则第二个参数 `tags` 指定将检测的最高优先级标记，否则它将按以下顺序检测标记：gconv、json、字段名称。

#### func MapDeep

```go
func MapDeep(value interface{}, tags ...string) map[string]interface{}
```

MapDeep does Map function recursively, which means if the attribute of `value` is also a struct/*struct, calls Map function on this attribute converting it to a map[string]interface{} type variable. Deprecated: used Map instead.

​	MapDeep 以递归方式执行 Map 函数，这意味着如果 的 `value` 属性也是 struct/*struct，则在此属性上调用 Map 函数，将其转换为 map[string]interface{} 类型变量。已弃用：改用地图。

#### func MapStrStr

```go
func MapStrStr(value interface{}, option ...MapOption) map[string]string
```

MapStrStr converts `value` to map[string]string. Note that there might be data copy for this map type converting.

​	MapStrStr `value` 转换为 map[string]string。请注意，此地图类型转换可能存在数据副本。

#### func MapStrStrDeep

```go
func MapStrStrDeep(value interface{}, tags ...string) map[string]string
```

MapStrStrDeep converts `value` to map[string]string recursively. Note that there might be data copy for this map type converting. Deprecated: used MapStrStr instead.

​	MapStrStrDeep 以递归方式 `value` 转换为 map[string]string。请注意，此地图类型转换可能存在数据副本。已弃用：改用 MapStrStr。

#### func MapToMap

```go
func MapToMap(params interface{}, pointer interface{}, mapping ...map[string]string) error
```

MapToMap converts any map type variable `params` to another map type variable `pointer` using reflect. See doMapToMap.

​	MapToMap `pointer` 使用 reflect 将任何地图类型变量 `params` 转换为另一个地图类型变量。请参阅 doMapToMap。

#### func MapToMaps

```go
func MapToMaps(params interface{}, pointer interface{}, mapping ...map[string]string) error
```

MapToMaps converts any slice type variable `params` to another map slice type variable `pointer`. See doMapToMaps.

​	MapToMaps 将任何切片类型变量 `params` 转换为另一个地图切片类型变量 `pointer` 。请参阅 doMapToMaps。

#### func Maps

```go
func Maps(value interface{}, option ...MapOption) []map[string]interface{}
```

Maps converts `value` to []map[string]interface{}. Note that it automatically checks and converts json string to []map if `value` is string/[]byte.

​	映射转换为 `value` []map[string]interface{}。请注意，如果 `value` string 是 string/[]byte，它会自动检查并将 json 字符串转换为 []map。

#### func MapsDeep

```go
func MapsDeep(value interface{}, tags ...string) []map[string]interface{}
```

MapsDeep converts `value` to []map[string]interface{} recursively.

​	MapsDeep 以递归方式 `value` 转换为 []map[string]interface{}。

TODO completely implement the recursive converting for all types. Deprecated: used Maps instead.

​	TODO完全实现了所有类型的递归转换。已弃用：改用地图。

#### func PtrAny <-2.2.1

```go
func PtrAny(any interface{}) *interface{}
```

PtrAny creates and returns an interface{} pointer variable to this value.

​	PtrAny 创建并返回指向此值的 interface{} 指针变量。

#### func PtrBool <-2.2.1

```go
func PtrBool(any interface{}) *bool
```

PtrBool creates and returns a bool pointer variable to this value.

​	PtrBool 创建并返回指向此值的布尔指针变量。

#### func PtrFloat32 <-2.2.1

```go
func PtrFloat32(any interface{}) *float32
```

PtrFloat32 creates and returns a float32 pointer variable to this value.

​	PtrFloat32 创建并返回此值的 float32 指针变量。

#### func PtrFloat64 <-2.2.1

```go
func PtrFloat64(any interface{}) *float64
```

PtrFloat64 creates and returns a float64 pointer variable to this value.

​	PtrFloat64 创建并返回此值的 float64 指针变量。

#### func PtrInt <-2.2.1

```go
func PtrInt(any interface{}) *int
```

PtrInt creates and returns an int pointer variable to this value.

​	PtrInt 创建并返回指向此值的 int 指针变量。

#### func PtrInt16 <-2.2.1

```go
func PtrInt16(any interface{}) *int16
```

PtrInt16 creates and returns an int16 pointer variable to this value.

​	PtrInt16 创建并返回指向此值的 int16 指针变量。

#### func PtrInt32 <-2.2.1

```go
func PtrInt32(any interface{}) *int32
```

PtrInt32 creates and returns an int32 pointer variable to this value.

​	PtrInt32 创建并返回此值的 int32 指针变量。

#### func PtrInt64 <-2.2.1

```go
func PtrInt64(any interface{}) *int64
```

PtrInt64 creates and returns an int64 pointer variable to this value.

​	PtrInt64 创建并返回指向此值的 int64 指针变量。

#### func PtrInt8 <-2.2.1

```go
func PtrInt8(any interface{}) *int8
```

PtrInt8 creates and returns an int8 pointer variable to this value.

​	PtrInt8 创建并返回此值的 int8 指针变量。

#### func PtrString <-2.2.1

```go
func PtrString(any interface{}) *string
```

PtrString creates and returns a string pointer variable to this value.

​	PtrString 创建并返回指向此值的字符串指针变量。

#### func PtrUint <-2.2.1

```go
func PtrUint(any interface{}) *uint
```

PtrUint creates and returns an uint pointer variable to this value.

​	PtrUint 创建并返回指向此值的 uint 指针变量。

#### func PtrUint16 <-2.2.1

```go
func PtrUint16(any interface{}) *uint16
```

PtrUint16 creates and returns an uint16 pointer variable to this value.

​	PtrUint16 创建并返回指向此值的 uint16 指针变量。

#### func PtrUint32 <-2.2.1

```go
func PtrUint32(any interface{}) *uint32
```

PtrUint32 creates and returns an uint32 pointer variable to this value.

​	PtrUint32 创建并返回指向此值的 uint32 指针变量。

#### func PtrUint64 <-2.2.1

```go
func PtrUint64(any interface{}) *uint64
```

PtrUint64 creates and returns an uint64 pointer variable to this value.

​	PtrUint64 创建并返回此值的 uint64 指针变量。

#### func PtrUint8 <-2.2.1

```go
func PtrUint8(any interface{}) *uint8
```

PtrUint8 creates and returns an uint8 pointer variable to this value.

​	PtrUint8 创建并返回指向此值的 uint8 指针变量。

#### func RegisterConverter <-2.5.2

```go
func RegisterConverter(fn interface{}) (err error)
```

RegisterConverter to register custom converter. It must be registered before you use this custom converting feature. It is suggested to do it in boot procedure of the process.

​	RegisterConverter 注册自定义转换器。在使用此自定义转换功能之前，必须先注册它。建议在进程的引导过程中执行此操作。

Note:

​	注意：

1. The parameter `fn` must be defined as pattern `func(T1) (T2, error)`. It will convert type `T1` to type `T2`.
   `fn` 该参数必须定义为 pattern `func(T1) (T2, error)` 。它会将 type `T1` 转换为 type `T2` 。
2. The `T1` should not be type of pointer, but the `T2` should be type of pointer.
   不应 `T1` 是指针的类型，而 `T2` 应该是指针的类型。

#### func Rune

```go
func Rune(any interface{}) rune
```

Rune converts `any` to rune.

​	符文转换为 `any` 符文。

#### func Runes

```go
func Runes(any interface{}) []rune
```

Runes converts `any` to []rune.

​	符文转换为 `any` []符文。

#### func Scan

```go
func Scan(srcValue interface{}, dstPointer interface{}, paramKeyToAttrMap ...map[string]string) (err error)
```

Scan automatically checks the type of `pointer` and converts `params` to `pointer`. It supports `pointer` in type of `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` for converting.

​	扫描会自动检查 的 `pointer` 类型并转换为 `params` `pointer` 。它支持 `pointer` `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` 转换类型。

TODO change `paramKeyToAttrMap` to `ScanOption` to be more scalable; add `DeepCopy` option for `ScanOption`.

​	TODO 更改 `paramKeyToAttrMap` 为 `ScanOption` 更具可扩展性;的 `ScanOption` add `DeepCopy` 选项。

#### func ScanList

```go
func ScanList(structSlice interface{}, structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error)
```

ScanList converts `structSlice` to struct slice which contains other complex struct attributes. Note that the parameter `structSlicePointer` should be type of *[]struct/*[]*struct.

​	ScanList `structSlice` 转换为包含其他复杂结构属性的结构切片。请注意，参数 `structSlicePointer` 的类型应为 []struct/[]*struct。

Usage example 1: Normal attribute struct relation:

​	使用示例1：法线属性结构关系：

```go
type EntityUser struct {
    Uid  int
    Name string
}

type EntityUserDetail struct {
    Uid     int
    Address string
}

type EntityUserScores struct {
    Id     int
    Uid    int
    Score  int
    Course string
}

type Entity struct {
    User       *EntityUser
    UserDetail *EntityUserDetail
    UserScores []*EntityUserScores
}
```

var users []*Entity var userRecords = EntityUser{Uid: 1, Name:“john”} var detailRecords = EntityUser{Uid: 1, Address: “chengdu”} var scoresRecords = EntityUser{Id: 1, Uid: 1, Score: 100, Course: “math”} ScanList(userRecords, &users, “User”) ScanList(userRecords, &users, “User”, “uid”) ScanList(detailRecords, &users, “UserDetail”, “User”, “uid:Uid”) ScanList(scoresRecords, &users, “UserScores”, “User”, “uid:Uid”) ScanList(scoresRecords, &users, “UserScores”, “User”, “uid”)

​	var users []*实体 var userRecords = EntityUser{Uid： 1， Name：“john”} var detailRecords = EntityUser{Uid： 1， 地址：“chengdu”} var scoresRecords = EntityUser{Id： 1， Uid： 1， Score： 100， 课程： “math”} ScanList（userRecords， &users， “用户”） ScanList（userRecords， &users， “用户”， “uid”） ScanList（detailRecords， &users， “UserDetail”， “User”， “uid：Uid”） ScanList（scoresRecords， &users， “UserScores”， “User”， “uid：Uid”） ScanList（scoresRecords， &users， “UserScores”， “UserScores”， “User”， “User”， “uid”）

Usage example 2: Embedded attribute struct relation:

​	使用示例2：嵌入属性结构关系：

```go
type EntityUser struct {
	   Uid  int
	   Name string
}

type EntityUserDetail struct {
	   Uid     int
	   Address string
}

type EntityUserScores struct {
	   Id    int
	   Uid   int
	   Score int
}

type Entity struct {
	   EntityUser
	   UserDetail EntityUserDetail
	   UserScores []EntityUserScores
}
```

var userRecords = EntityUser{Uid: 1, Name:“john”} var detailRecords = EntityUser{Uid: 1, Address: “chengdu”} var scoresRecords = EntityUser{Id: 1, Uid: 1, Score: 100, Course: “math”} ScanList(userRecords, &users) ScanList(detailRecords, &users, “UserDetail”, “uid”) ScanList(scoresRecords, &users, “UserScores”, “uid”)

​	var userRecords = EntityUser{Uid： 1， Name：“john”} var detailRecords = EntityUser{Uid： 1， Address： “chengdu”} var scoresRecords = EntityUser{Id： 1， Uid： 1， Score： 100， Course： “math”} ScanList（userRecords， &users） ScanList（detailRecords， &users， “UserDetail”， “uid”） ScanList（scoresRecords， &users， “UserScores”， “uid”）

The parameters “User/UserDetail/UserScores” in the example codes specify the target attribute struct that current result will be bound to.

​	示例代码中的参数“User/UserDetail/UserScores”指定当前结果将绑定到的目标属性结构。

The “uid” in the example codes is the table field name of the result, and the “Uid” is the relational struct attribute name - not the attribute name of the bound to target. In the example codes, it’s attribute name “Uid” of “User” of entity “Entity”. It automatically calculates the HasOne/HasMany relationship with given `relation` parameter.

​	示例代码中的“uid”是结果的表字段名称，“Uid”是关系结构属性名称，而不是绑定到目标的属性名称。在示例代码中，它是实体“Entity”的“User”的属性名称“Uid”。它自动计算给定 `relation` 参数的 HasOne/HasMany 关系。

See the example or unit testing cases for clear understanding for this function.

​	请参阅示例或单元测试用例，以清楚地了解此函数。

#### func SliceAny

```go
func SliceAny(any interface{}) []interface{}
```

SliceAny is alias of Interfaces.

​	SliceAny 是 Interfaces 的别名。

#### func SliceFloat

```go
func SliceFloat(any interface{}) []float64
```

SliceFloat is alias of Floats.

​	SliceFloat 是 Floats 的别名。

#### func SliceFloat32

```go
func SliceFloat32(any interface{}) []float32
```

SliceFloat32 is alias of Float32s.

​	SliceFloat32 是 Float32s 的别名。

#### func SliceFloat64

```go
func SliceFloat64(any interface{}) []float64
```

SliceFloat64 is alias of Float64s.

​	SliceFloat64 是 Float64s 的别名。

#### func SliceInt

```go
func SliceInt(any interface{}) []int
```

SliceInt is alias of Ints.

​	SliceInt 是 Ints 的别名。

#### func SliceInt32

```go
func SliceInt32(any interface{}) []int32
```

SliceInt32 is alias of Int32s.

​	SliceInt32 是 Int32 的别名。

#### func SliceInt64

```go
func SliceInt64(any interface{}) []int64
```

SliceInt64 is alias of Int64s.

​	SliceInt64 是 Int64 的别名。

#### func SliceMap

```go
func SliceMap(any interface{}, option ...MapOption) []map[string]interface{}
```

SliceMap is alias of Maps.

​	SliceMap 是 Maps 的别名。

#### func SliceMapDeep

```go
func SliceMapDeep(any interface{}) []map[string]interface{}
```

SliceMapDeep is alias of MapsDeep. Deprecated: used SliceMap instead.

​	SliceMapDeep 是 MapsDeep 的别名。已弃用：改用 SliceMap。

#### func SliceStr

```go
func SliceStr(any interface{}) []string
```

SliceStr is alias of Strings.

​	SliceStr 是 Strings 的别名。

#### func SliceStruct

```go
func SliceStruct(params interface{}, pointer interface{}, mapping ...map[string]string) (err error)
```

SliceStruct is alias of Structs.

​	SliceStruct 是 Structs 的别名。

#### func SliceUint

```go
func SliceUint(any interface{}) []uint
```

SliceUint is alias of Uints.

​	SliceUint 是 Uints 的别名。

#### func SliceUint32

```go
func SliceUint32(any interface{}) []uint32
```

SliceUint32 is alias of Uint32s.

​	SliceUint32 是 Uint32s 的别名。

#### func SliceUint64

```go
func SliceUint64(any interface{}) []uint64
```

SliceUint64 is alias of Uint64s.

​	SliceUint64 是 Uint64s 的别名。

#### func String

```go
func String(any interface{}) string
```

String converts `any` to string. It’s most commonly used converting function.

​	字符串转换为 `any` 字符串。这是最常用的转换功能。

#### func Strings

```go
func Strings(any interface{}) []string
```

Strings converts `any` to []string.

​	字符串转换为 `any` []string。

#### func Struct

```go
func Struct(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error)
```

Struct maps the params key-value pairs to the corresponding struct object’s attributes. The third parameter `mapping` is unnecessary, indicating the mapping rules between the custom key name and the attribute name(case-sensitive).

​	struct 将参数键值对映射到相应的 struct 对象的属性。第三个参数 `mapping` 是不必要的，表示自定义键名和属性名（区分大小写）之间的映射规则。

Note:

​	注意：

1. The `params` can be any type of map/struct, usually a map.
   可以 `params` 是任何类型的映射/结构，通常是映射。
2. The `pointer` should be type of *struct/**struct, which is a pointer to struct object or struct pointer.
   `pointer` 应为 *struct/**struct 的类型，它是指向 struct 对象或 struct 指针的指针。
3. Only the public attributes of struct object can be mapped.
   只能映射 struct 对象的公共属性。
4. If `params` is a map, the key of the map `params` can be lowercase. It will automatically convert the first letter of the key to uppercase in mapping procedure to do the matching. It ignores the map key, if it does not match.
   如果 `params` 是地图，则地图 `params` 的键可以是小写的。它会在映射过程中自动将键的第一个字母转换为大写字母以进行匹配。如果映射键不匹配，它将忽略该映射键。

#### func StructTag

```go
func StructTag(params interface{}, pointer interface{}, priorityTag string) (err error)
```

StructTag acts as Struct but also with support for priority tag feature, which retrieves the specified tags for `params` key-value items to struct attribute names mapping. The parameter `priorityTag` supports multiple tags that can be joined with char ‘,’.

​	StructTag 充当 Struct，但也支持优先级标记功能，该功能将 `params` 键值项的指定标签检索到结构属性名称映射。该参数 `priorityTag` 支持多个标签，这些标签可以与 char '，' 连接。

#### func Structs

```go
func Structs(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error)
```

Structs converts any slice to given struct slice. Also see Scan, Struct.

​	Structs 将任何切片转换为给定的结构切片。另请参阅 Scan、Struct。

#### func StructsTag

```go
func StructsTag(params interface{}, pointer interface{}, priorityTag string) (err error)
```

StructsTag acts as Structs but also with support for priority tag feature, which retrieves the specified tags for `params` key-value items to struct attribute names mapping. The parameter `priorityTag` supports multiple tags that can be joined with char ‘,’.

​	StructsTag 充当 Structs，但也支持优先级标记功能，该功能将 `params` 键值项的指定标签检索到结构属性名称映射。该参数 `priorityTag` 支持多个标签，这些标签可以与 char '，' 连接。

#### func Time

```go
func Time(any interface{}, format ...string) time.Time
```

Time converts `any` to time.Time.

​	时间 `any` 转换为时间。时间。

#### func Uint

```go
func Uint(any interface{}) uint
```

Uint converts `any` to uint.

​	Uint 转换为 `any` uint。

#### func Uint16

```go
func Uint16(any interface{}) uint16
```

Uint16 converts `any` to uint16.

​	Uint16 `any` 转换为 uint16。

#### func Uint32

```go
func Uint32(any interface{}) uint32
```

Uint32 converts `any` to uint32.

​	Uint32 `any` 转换为 uint32。

#### func Uint32s

```go
func Uint32s(any interface{}) []uint32
```

Uint32s converts `any` to []uint32.

​	Uint32s `any` 转换为 []uint32。

#### func Uint64

```go
func Uint64(any interface{}) uint64
```

Uint64 converts `any` to uint64.

​	Uint64 `any` 转换为 uint64。

#### func Uint64s

```go
func Uint64s(any interface{}) []uint64
```

Uint64s converts `any` to []uint64.

​	Uint64s `any` 转换为 []uint64。

#### func Uint8

```go
func Uint8(any interface{}) uint8
```

Uint8 converts `any` to uint8.

​	Uint8 `any` 转换为 uint8。

#### func Uints

```go
func Uints(any interface{}) []uint
```

Uints converts `any` to []uint.

​	Uints `any` 转换为 []uint。

#### func UnsafeBytesToStr

```go
func UnsafeBytesToStr(b []byte) string
```

UnsafeBytesToStr converts []byte to string without memory copy. Note that, if you completely sure you will never use `b` variable in the feature, you can use this unsafe function to implement type conversion in high performance.

​	UnsafeBytesToStr 将 []byte 转换为不带内存副本的字符串。请注意，如果您完全确定永远不会在功能中使用 `b` 变量，则可以使用此不安全的函数来实现高性能的类型转换。

#### func UnsafeStrToBytes

```go
func UnsafeStrToBytes(s string) []byte
```

UnsafeStrToBytes converts string to []byte without memory copy. Note that, if you completely sure you will never use `s` variable in the feature, you can use this unsafe function to implement type conversion in high performance.

​	UnsafeStrToBytes 将字符串转换为 []字节，而无需内存复制。请注意，如果您完全确定永远不会在功能中使用 `s` 变量，则可以使用此不安全的函数来实现高性能的类型转换。

## 类型

### type MapOption <-2.6.0

```go
type MapOption struct {
	// Deep marks doing Map function recursively, which means if the attribute of given converting value
	// is also a struct/*struct, it automatically calls Map function on this attribute converting it to
	// a map[string]interface{} type variable.
	Deep bool

	// OmitEmpty ignores the attributes that has json `omitempty` tag.
	OmitEmpty bool

	// Tags specifies the converted map key name by struct tag name.
	Tags []string
}
```

MapOption specifies the option for map converting.

​	MapOption 指定地图转换选项。