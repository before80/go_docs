+++
title = "gconv"
date = 2024-03-21T17:59:20+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gconv

Package gconv implements powerful and convenient converting functionality for any types of variables.

This package should keep much less dependencies with other packages.

### Constants 

This section is empty.

### Variables 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/util/gconv/gconv.go#L29)

``` go
var (

	// StructTagPriority defines the default priority tags for Map*/Struct* functions.
	// Note that, the `gconv/param` tags are used by old version of package.
	// It is strongly recommended using short tag `c/p` instead in the future.
	StructTagPriority = gtag.StructTagPriority
)
```

### Functions 

##### func Bool 

``` go
func Bool(any interface{}) bool
```

Bool converts `any` to bool. It returns false if `any` is: false, "", 0, "false", "off", "no", empty slice/map.

##### func Byte 

``` go
func Byte(any interface{}) byte
```

Byte converts `any` to byte.

##### func Bytes 

``` go
func Bytes(any interface{}) []byte
```

Bytes converts `any` to []byte.

##### func Convert 

``` go
func Convert(fromValue interface{}, toTypeName string, extraParams ...interface{}) interface{}
```

Convert converts the variable `fromValue` to the type `toTypeName`, the type `toTypeName` is specified by string.

The optional parameter `extraParams` is used for additional necessary parameter for this conversion. It supports common basic types conversion as its conversion based on type name string.

##### func ConvertWithRefer <-2.5.4

``` go
func ConvertWithRefer(fromValue interface{}, referValue interface{}, extraParams ...interface{}) interface{}
```

ConvertWithRefer converts the variable `fromValue` to the type referred by value `referValue`.

The optional parameter `extraParams` is used for additional necessary parameter for this conversion. It supports common basic types conversion as its conversion based on type name string.

##### func Duration 

``` go
func Duration(any interface{}) time.Duration
```

Duration converts `any` to time.Duration. If `any` is string, then it uses time.ParseDuration to convert it. If `any` is numeric, then it converts `any` as nanoseconds.

##### func Float32 

``` go
func Float32(any interface{}) float32
```

Float32 converts `any` to float32.

##### func Float32s 

``` go
func Float32s(any interface{}) []float32
```

Float32s converts `any` to []float32.

##### func Float64 

``` go
func Float64(any interface{}) float64
```

Float64 converts `any` to float64.

##### func Float64s 

``` go
func Float64s(any interface{}) []float64
```

Float64s converts `any` to []float64.

##### func Floats 

``` go
func Floats(any interface{}) []float64
```

Floats converts `any` to []float64.

##### func GTime 

``` go
func GTime(any interface{}, format ...string) *gtime.Time
```

GTime converts `any` to *gtime.Time. The parameter `format` can be used to specify the format of `any`. It returns the converted value that matched the first format of the formats slice. If no `format` given, it converts `any` using gtime.NewFromTimeStamp if `any` is numeric, or using gtime.StrToTime if `any` is string.

##### func Int 

``` go
func Int(any interface{}) int
```

Int converts `any` to int.

##### func Int16 

``` go
func Int16(any interface{}) int16
```

Int16 converts `any` to int16.

##### func Int32 

``` go
func Int32(any interface{}) int32
```

Int32 converts `any` to int32.

##### func Int32s 

``` go
func Int32s(any interface{}) []int32
```

Int32s converts `any` to []int32.

##### func Int64 

``` go
func Int64(any interface{}) int64
```

Int64 converts `any` to int64.

##### func Int64s 

``` go
func Int64s(any interface{}) []int64
```

Int64s converts `any` to []int64.

##### func Int8 

``` go
func Int8(any interface{}) int8
```

Int8 converts `any` to int8.

##### func Interfaces 

``` go
func Interfaces(any interface{}) []interface{}
```

Interfaces converts `any` to []interface{}.

##### func Ints 

``` go
func Ints(any interface{}) []int
```

Ints converts `any` to []int.

##### func Map 

``` go
func Map(value interface{}, option ...MapOption) map[string]interface{}
```

Map converts any variable `value` to map[string]interface{}. If the parameter `value` is not a map/struct/*struct type, then the conversion will fail and returns nil.

If `value` is a struct/*struct object, the second parameter `tags` specifies the most priority tags that will be detected, otherwise it detects the tags in order of: gconv, json, field name.

##### func MapDeep 

``` go
func MapDeep(value interface{}, tags ...string) map[string]interface{}
```

MapDeep does Map function recursively, which means if the attribute of `value` is also a struct/*struct, calls Map function on this attribute converting it to a map[string]interface{} type variable. Deprecated: used Map instead.

##### func MapStrStr 

``` go
func MapStrStr(value interface{}, option ...MapOption) map[string]string
```

MapStrStr converts `value` to map[string]string. Note that there might be data copy for this map type converting.

##### func MapStrStrDeep 

``` go
func MapStrStrDeep(value interface{}, tags ...string) map[string]string
```

MapStrStrDeep converts `value` to map[string]string recursively. Note that there might be data copy for this map type converting. Deprecated: used MapStrStr instead.

##### func MapToMap 

``` go
func MapToMap(params interface{}, pointer interface{}, mapping ...map[string]string) error
```

MapToMap converts any map type variable `params` to another map type variable `pointer` using reflect. See doMapToMap.

##### func MapToMaps 

``` go
func MapToMaps(params interface{}, pointer interface{}, mapping ...map[string]string) error
```

MapToMaps converts any slice type variable `params` to another map slice type variable `pointer`. See doMapToMaps.

##### func Maps 

``` go
func Maps(value interface{}, option ...MapOption) []map[string]interface{}
```

Maps converts `value` to []map[string]interface{}. Note that it automatically checks and converts json string to []map if `value` is string/[]byte.

##### func MapsDeep 

``` go
func MapsDeep(value interface{}, tags ...string) []map[string]interface{}
```

MapsDeep converts `value` to []map[string]interface{} recursively.

TODO completely implement the recursive converting for all types. Deprecated: used Maps instead.

##### func PtrAny <-2.2.1

``` go
func PtrAny(any interface{}) *interface{}
```

PtrAny creates and returns an interface{} pointer variable to this value.

##### func PtrBool <-2.2.1

``` go
func PtrBool(any interface{}) *bool
```

PtrBool creates and returns a bool pointer variable to this value.

##### func PtrFloat32 <-2.2.1

``` go
func PtrFloat32(any interface{}) *float32
```

PtrFloat32 creates and returns a float32 pointer variable to this value.

##### func PtrFloat64 <-2.2.1

``` go
func PtrFloat64(any interface{}) *float64
```

PtrFloat64 creates and returns a float64 pointer variable to this value.

##### func PtrInt <-2.2.1

``` go
func PtrInt(any interface{}) *int
```

PtrInt creates and returns an int pointer variable to this value.

##### func PtrInt16 <-2.2.1

``` go
func PtrInt16(any interface{}) *int16
```

PtrInt16 creates and returns an int16 pointer variable to this value.

##### func PtrInt32 <-2.2.1

``` go
func PtrInt32(any interface{}) *int32
```

PtrInt32 creates and returns an int32 pointer variable to this value.

##### func PtrInt64 <-2.2.1

``` go
func PtrInt64(any interface{}) *int64
```

PtrInt64 creates and returns an int64 pointer variable to this value.

##### func PtrInt8 <-2.2.1

``` go
func PtrInt8(any interface{}) *int8
```

PtrInt8 creates and returns an int8 pointer variable to this value.

##### func PtrString <-2.2.1

``` go
func PtrString(any interface{}) *string
```

PtrString creates and returns a string pointer variable to this value.

##### func PtrUint <-2.2.1

``` go
func PtrUint(any interface{}) *uint
```

PtrUint creates and returns an uint pointer variable to this value.

##### func PtrUint16 <-2.2.1

``` go
func PtrUint16(any interface{}) *uint16
```

PtrUint16 creates and returns an uint16 pointer variable to this value.

##### func PtrUint32 <-2.2.1

``` go
func PtrUint32(any interface{}) *uint32
```

PtrUint32 creates and returns an uint32 pointer variable to this value.

##### func PtrUint64 <-2.2.1

``` go
func PtrUint64(any interface{}) *uint64
```

PtrUint64 creates and returns an uint64 pointer variable to this value.

##### func PtrUint8 <-2.2.1

``` go
func PtrUint8(any interface{}) *uint8
```

PtrUint8 creates and returns an uint8 pointer variable to this value.

##### func RegisterConverter <-2.5.2

``` go
func RegisterConverter(fn interface{}) (err error)
```

RegisterConverter to register custom converter. It must be registered before you use this custom converting feature. It is suggested to do it in boot procedure of the process.

Note:

1. The parameter `fn` must be defined as pattern `func(T1) (T2, error)`. It will convert type `T1` to type `T2`.
2. The `T1` should not be type of pointer, but the `T2` should be type of pointer.

##### func Rune 

``` go
func Rune(any interface{}) rune
```

Rune converts `any` to rune.

##### func Runes 

``` go
func Runes(any interface{}) []rune
```

Runes converts `any` to []rune.

##### func Scan 

``` go
func Scan(srcValue interface{}, dstPointer interface{}, paramKeyToAttrMap ...map[string]string) (err error)
```

Scan automatically checks the type of `pointer` and converts `params` to `pointer`. It supports `pointer` in type of `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` for converting.

TODO change `paramKeyToAttrMap` to `ScanOption` to be more scalable; add `DeepCopy` option for `ScanOption`.

##### func ScanList 

``` go
func ScanList(structSlice interface{}, structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error)
```

ScanList converts `structSlice` to struct slice which contains other complex struct attributes. Note that the parameter `structSlicePointer` should be type of *[]struct/*[]*struct.

Usage example 1: Normal attribute struct relation:

``` go
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

var users []*Entity var userRecords = EntityUser{Uid: 1, Name:"john"} var detailRecords = EntityUser{Uid: 1, Address: "chengdu"} var scoresRecords = EntityUser{Id: 1, Uid: 1, Score: 100, Course: "math"} ScanList(userRecords, &users, "User") ScanList(userRecords, &users, "User", "uid") ScanList(detailRecords, &users, "UserDetail", "User", "uid:Uid") ScanList(scoresRecords, &users, "UserScores", "User", "uid:Uid") ScanList(scoresRecords, &users, "UserScores", "User", "uid")

Usage example 2: Embedded attribute struct relation:

``` go
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

var userRecords = EntityUser{Uid: 1, Name:"john"} var detailRecords = EntityUser{Uid: 1, Address: "chengdu"} var scoresRecords = EntityUser{Id: 1, Uid: 1, Score: 100, Course: "math"} ScanList(userRecords, &users) ScanList(detailRecords, &users, "UserDetail", "uid") ScanList(scoresRecords, &users, "UserScores", "uid")

The parameters "User/UserDetail/UserScores" in the example codes specify the target attribute struct that current result will be bound to.

The "uid" in the example codes is the table field name of the result, and the "Uid" is the relational struct attribute name - not the attribute name of the bound to target. In the example codes, it's attribute name "Uid" of "User" of entity "Entity". It automatically calculates the HasOne/HasMany relationship with given `relation` parameter.

See the example or unit testing cases for clear understanding for this function.

##### func SliceAny 

``` go
func SliceAny(any interface{}) []interface{}
```

SliceAny is alias of Interfaces.

##### func SliceFloat 

``` go
func SliceFloat(any interface{}) []float64
```

SliceFloat is alias of Floats.

##### func SliceFloat32 

``` go
func SliceFloat32(any interface{}) []float32
```

SliceFloat32 is alias of Float32s.

##### func SliceFloat64 

``` go
func SliceFloat64(any interface{}) []float64
```

SliceFloat64 is alias of Float64s.

##### func SliceInt 

``` go
func SliceInt(any interface{}) []int
```

SliceInt is alias of Ints.

##### func SliceInt32 

``` go
func SliceInt32(any interface{}) []int32
```

SliceInt32 is alias of Int32s.

##### func SliceInt64 

``` go
func SliceInt64(any interface{}) []int64
```

SliceInt64 is alias of Int64s.

##### func SliceMap 

``` go
func SliceMap(any interface{}, option ...MapOption) []map[string]interface{}
```

SliceMap is alias of Maps.

##### func SliceMapDeep 

``` go
func SliceMapDeep(any interface{}) []map[string]interface{}
```

SliceMapDeep is alias of MapsDeep. Deprecated: used SliceMap instead.

##### func SliceStr 

``` go
func SliceStr(any interface{}) []string
```

SliceStr is alias of Strings.

##### func SliceStruct 

``` go
func SliceStruct(params interface{}, pointer interface{}, mapping ...map[string]string) (err error)
```

SliceStruct is alias of Structs.

##### func SliceUint 

``` go
func SliceUint(any interface{}) []uint
```

SliceUint is alias of Uints.

##### func SliceUint32 

``` go
func SliceUint32(any interface{}) []uint32
```

SliceUint32 is alias of Uint32s.

##### func SliceUint64 

``` go
func SliceUint64(any interface{}) []uint64
```

SliceUint64 is alias of Uint64s.

##### func String 

``` go
func String(any interface{}) string
```

String converts `any` to string. It's most commonly used converting function.

##### func Strings 

``` go
func Strings(any interface{}) []string
```

Strings converts `any` to []string.

##### func Struct 

``` go
func Struct(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error)
```

Struct maps the params key-value pairs to the corresponding struct object's attributes. The third parameter `mapping` is unnecessary, indicating the mapping rules between the custom key name and the attribute name(case-sensitive).

Note:

1. The `params` can be any type of map/struct, usually a map.
2. The `pointer` should be type of *struct/**struct, which is a pointer to struct object or struct pointer.
3. Only the public attributes of struct object can be mapped.
4. If `params` is a map, the key of the map `params` can be lowercase. It will automatically convert the first letter of the key to uppercase in mapping procedure to do the matching. It ignores the map key, if it does not match.

##### func StructTag 

``` go
func StructTag(params interface{}, pointer interface{}, priorityTag string) (err error)
```

StructTag acts as Struct but also with support for priority tag feature, which retrieves the specified tags for `params` key-value items to struct attribute names mapping. The parameter `priorityTag` supports multiple tags that can be joined with char ','.

##### func Structs 

``` go
func Structs(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error)
```

Structs converts any slice to given struct slice. Also see Scan, Struct.

##### func StructsTag 

``` go
func StructsTag(params interface{}, pointer interface{}, priorityTag string) (err error)
```

StructsTag acts as Structs but also with support for priority tag feature, which retrieves the specified tags for `params` key-value items to struct attribute names mapping. The parameter `priorityTag` supports multiple tags that can be joined with char ','.

##### func Time 

``` go
func Time(any interface{}, format ...string) time.Time
```

Time converts `any` to time.Time.

##### func Uint 

``` go
func Uint(any interface{}) uint
```

Uint converts `any` to uint.

##### func Uint16 

``` go
func Uint16(any interface{}) uint16
```

Uint16 converts `any` to uint16.

##### func Uint32 

``` go
func Uint32(any interface{}) uint32
```

Uint32 converts `any` to uint32.

##### func Uint32s 

``` go
func Uint32s(any interface{}) []uint32
```

Uint32s converts `any` to []uint32.

##### func Uint64 

``` go
func Uint64(any interface{}) uint64
```

Uint64 converts `any` to uint64.

##### func Uint64s 

``` go
func Uint64s(any interface{}) []uint64
```

Uint64s converts `any` to []uint64.

##### func Uint8 

``` go
func Uint8(any interface{}) uint8
```

Uint8 converts `any` to uint8.

##### func Uints 

``` go
func Uints(any interface{}) []uint
```

Uints converts `any` to []uint.

##### func UnsafeBytesToStr 

``` go
func UnsafeBytesToStr(b []byte) string
```

UnsafeBytesToStr converts []byte to string without memory copy. Note that, if you completely sure you will never use `b` variable in the feature, you can use this unsafe function to implement type conversion in high performance.

##### func UnsafeStrToBytes 

``` go
func UnsafeStrToBytes(s string) []byte
```

UnsafeStrToBytes converts string to []byte without memory copy. Note that, if you completely sure you will never use `s` variable in the feature, you can use this unsafe function to implement type conversion in high performance.

### Types 

#### type MapOption <-2.6.0

``` go
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