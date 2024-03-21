+++
title = "gvar"
date = 2024-03-21T17:45:43+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gvar

Package gvar provides an universal variable type, like generics.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type MapOption <-2.6.0

``` go
type MapOption = gconv.MapOption
```

MapOption specifies the option for map converting.

#### type Var 

``` go
type Var struct {
	// contains filtered or unexported fields
}
```

Var is an universal variable type implementer.

##### func New 

``` go
func New(value interface{}, safe ...bool) *Var
```

New creates and returns a new Var with given `value`. The optional parameter `safe` specifies whether Var is used in concurrent-safety, which is false in default.

##### (*Var) Array 

``` go
func (v *Var) Array() []interface{}
```

Array is alias of Interfaces.

##### Example

``` go
```
##### (*Var) Bool 

``` go
func (v *Var) Bool() bool
```

Bool converts and returns `v` as bool.

##### Example

``` go
```
##### (*Var) Bytes 

``` go
func (v *Var) Bytes() []byte
```

Bytes converts and returns `v` as []byte.

##### Example

``` go
```
##### (*Var) Clone 

``` go
func (v *Var) Clone() *Var
```

Clone does a shallow copy of current Var and returns a pointer to this Var.

##### Example

``` go
```
##### (*Var) Copy <-2.1.0

``` go
func (v *Var) Copy() *Var
```

Copy does a deep copy of current Var and returns a pointer to this Var.

##### (*Var) DeepCopy <-2.1.0

``` go
func (v *Var) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*Var) Duration 

``` go
func (v *Var) Duration() time.Duration
```

Duration converts and returns `v` as time.Duration. If value of `v` is string, then it uses time.ParseDuration for conversion.

##### Example

``` go
```
##### (*Var) Float32 

``` go
func (v *Var) Float32() float32
```

Float32 converts and returns `v` as float32.

##### Example

``` go
```
##### (*Var) Float32s 

``` go
func (v *Var) Float32s() []float32
```

Float32s converts and returns `v` as []float32.

##### Example

``` go
```
##### (*Var) Float64 

``` go
func (v *Var) Float64() float64
```

Float64 converts and returns `v` as float64.

##### (*Var) Float64s 

``` go
func (v *Var) Float64s() []float64
```

Float64s converts and returns `v` as []float64.

##### Example

``` go
```
##### (*Var) Floats 

``` go
func (v *Var) Floats() []float64
```

Floats is alias of Float64s.

##### Example

``` go
```
##### (*Var) GTime 

``` go
func (v *Var) GTime(format ...string) *gtime.Time
```

GTime converts and returns `v` as *gtime.Time. The parameter `format` specifies the format of the time string using gtime, eg: Y-m-d H:i:s.

##### Example

``` go
```
##### (*Var) Int 

``` go
func (v *Var) Int() int
```

Int converts and returns `v` as int.

##### Example

``` go
```
##### (*Var) Int16 

``` go
func (v *Var) Int16() int16
```

Int16 converts and returns `v` as int16.

##### (*Var) Int32 

``` go
func (v *Var) Int32() int32
```

Int32 converts and returns `v` as int32.

##### (*Var) Int64 

``` go
func (v *Var) Int64() int64
```

Int64 converts and returns `v` as int64.

##### (*Var) Int64s 

``` go
func (v *Var) Int64s() []int64
```

Int64s converts and returns `v` as []int64.

##### Example

``` go
```
##### (*Var) Int8 

``` go
func (v *Var) Int8() int8
```

Int8 converts and returns `v` as int8.

##### (*Var) Interface 

``` go
func (v *Var) Interface() interface{}
```

Interface is alias of Val.

##### Example

``` go
```
##### (*Var) Interfaces 

``` go
func (v *Var) Interfaces() []interface{}
```

Interfaces converts and returns `v` as []interfaces{}.

##### Example

``` go
```
##### (*Var) Ints 

``` go
func (v *Var) Ints() []int
```

Ints converts and returns `v` as []int.

##### Example

``` go
```
##### (*Var) IsEmpty 

``` go
func (v *Var) IsEmpty() bool
```

IsEmpty checks whether `v` is empty.

##### Example

``` go
```
##### (*Var) IsFloat 

``` go
func (v *Var) IsFloat() bool
```

IsFloat checks whether `v` is type of float.

##### Example

``` go
```
##### (*Var) IsInt 

``` go
func (v *Var) IsInt() bool
```

IsInt checks whether `v` is type of int.

##### Example

``` go
```
##### (*Var) IsMap 

``` go
func (v *Var) IsMap() bool
```

IsMap checks whether `v` is type of map.

##### Example

``` go
```
##### (*Var) IsNil 

``` go
func (v *Var) IsNil() bool
```

IsNil checks whether `v` is nil.

##### Example

``` go
```
##### (*Var) IsSlice 

``` go
func (v *Var) IsSlice() bool
```

IsSlice checks whether `v` is type of slice.

##### Example

``` go
```
##### (*Var) IsStruct 

``` go
func (v *Var) IsStruct() bool
```

IsStruct checks whether `v` is type of struct.

##### Example

``` go
```
##### (*Var) IsUint 

``` go
func (v *Var) IsUint() bool
```

IsUint checks whether `v` is type of uint.

##### Example

``` go
```
##### (*Var) ListItemValues 

``` go
func (v *Var) ListItemValues(key interface{}) (values []interface{})
```

ListItemValues retrieves and returns the elements of all item struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice.

##### Example

``` go
```
##### (*Var) ListItemValuesUnique 

``` go
func (v *Var) ListItemValuesUnique(key string) []interface{}
```

ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice.

##### Example

``` go
```
##### (*Var) Map 

``` go
func (v *Var) Map(option ...MapOption) map[string]interface{}
```

Map converts and returns `v` as map[string]interface{}.

##### Example

``` go
```
##### (*Var) MapDeep 

``` go
func (v *Var) MapDeep(tags ...string) map[string]interface{}
```

MapDeep converts and returns `v` as map[string]interface{} recursively. Deprecated: used Map instead.

##### Example

``` go
```
##### (*Var) MapStrAny 

``` go
func (v *Var) MapStrAny(option ...MapOption) map[string]interface{}
```

MapStrAny is like function Map, but implements the interface of MapStrAny.

##### Example

``` go
```
##### (*Var) MapStrStr 

``` go
func (v *Var) MapStrStr(option ...MapOption) map[string]string
```

MapStrStr converts and returns `v` as map[string]string.

##### Example

``` go
```
##### (*Var) MapStrStrDeep 

``` go
func (v *Var) MapStrStrDeep(tags ...string) map[string]string
```

MapStrStrDeep converts and returns `v` as map[string]string recursively. Deprecated: used MapStrStr instead.

##### Example

``` go
```
##### (*Var) MapStrVar 

``` go
func (v *Var) MapStrVar(option ...MapOption) map[string]*Var
```

MapStrVar converts and returns `v` as map[string]Var.

##### Example

``` go
```
##### (*Var) MapStrVarDeep 

``` go
func (v *Var) MapStrVarDeep(tags ...string) map[string]*Var
```

MapStrVarDeep converts and returns `v` as map[string]*Var recursively. Deprecated: used MapStrVar instead.

##### Example

``` go
```
##### (*Var) MapToMap 

``` go
func (v *Var) MapToMap(pointer interface{}, mapping ...map[string]string) (err error)
```

MapToMap converts any map type variable `params` to another map type variable `pointer`. See gconv.MapToMap.

##### Example

``` go
```
##### (*Var) MapToMaps 

``` go
func (v *Var) MapToMaps(pointer interface{}, mapping ...map[string]string) (err error)
```

MapToMaps converts any map type variable `params` to another map type variable `pointer`. See gconv.MapToMaps.

##### Example

``` go
```
##### (*Var) MapToMapsDeep 

``` go
func (v *Var) MapToMapsDeep(pointer interface{}, mapping ...map[string]string) (err error)
```

MapToMapsDeep converts any map type variable `params` to another map type variable `pointer` recursively. See gconv.MapToMapsDeep.

##### Example

``` go
```
##### (*Var) Maps 

``` go
func (v *Var) Maps(option ...MapOption) []map[string]interface{}
```

Maps converts and returns `v` as map[string]string. See gconv.Maps.

##### Example

``` go
```
##### (*Var) MapsDeep 

``` go
func (v *Var) MapsDeep(tags ...string) []map[string]interface{}
```

MapsDeep converts `value` to []map[string]interface{} recursively. Deprecated: used Maps instead.

##### Example

``` go
```
##### (Var) MarshalJSON 

``` go
func (v Var) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*Var) Scan 

``` go
func (v *Var) Scan(pointer interface{}, mapping ...map[string]string) error
```

Scan automatically checks the type of `pointer` and converts `params` to `pointer`. It supports `pointer` with type of `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` for converting.

See gconv.Scan.

##### Example

``` go
```
##### (*Var) Set 

``` go
func (v *Var) Set(value interface{}) (old interface{})
```

Set sets `value` to `v`, and returns the old value.

##### Example

``` go
```
##### (*Var) Slice 

``` go
func (v *Var) Slice() []interface{}
```

Slice is alias of Interfaces.

##### Example

``` go
```
##### (*Var) String 

``` go
func (v *Var) String() string
```

String converts and returns `v` as string.

##### Example

``` go
```
##### (*Var) Strings 

``` go
func (v *Var) Strings() []string
```

Strings converts and returns `v` as []string.

##### Example

``` go
```
##### (*Var) Struct 

``` go
func (v *Var) Struct(pointer interface{}, mapping ...map[string]string) error
```

Struct maps value of `v` to `pointer`. The parameter `pointer` should be a pointer to a struct instance. The parameter `mapping` is used to specify the key-to-attribute mapping rules.

##### Example

``` go
```
##### (*Var) Structs 

``` go
func (v *Var) Structs(pointer interface{}, mapping ...map[string]string) error
```

Structs converts and returns `v` as given struct slice.

##### Example

``` go
```
##### (*Var) Time 

``` go
func (v *Var) Time(format ...string) time.Time
```

Time converts and returns `v` as time.Time. The parameter `format` specifies the format of the time string using gtime, eg: Y-m-d H:i:s.

##### Example

``` go
```
##### (*Var) Uint 

``` go
func (v *Var) Uint() uint
```

Uint converts and returns `v` as uint.

##### Example

``` go
```
##### (*Var) Uint16 

``` go
func (v *Var) Uint16() uint16
```

Uint16 converts and returns `v` as uint16.

##### (*Var) Uint32 

``` go
func (v *Var) Uint32() uint32
```

Uint32 converts and returns `v` as uint32.

##### (*Var) Uint64 

``` go
func (v *Var) Uint64() uint64
```

Uint64 converts and returns `v` as uint64.

##### (*Var) Uint64s 

``` go
func (v *Var) Uint64s() []uint64
```

Uint64s converts and returns `v` as []uint64.

##### Example

``` go
```
##### (*Var) Uint8 

``` go
func (v *Var) Uint8() uint8
```

Uint8 converts and returns `v` as uint8.

##### (*Var) Uints 

``` go
func (v *Var) Uints() []uint
```

Uints converts and returns `v` as []uint.

##### Example

``` go
```
##### (*Var) UnmarshalJSON 

``` go
func (v *Var) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*Var) UnmarshalValue 

``` go
func (v *Var) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for Var.

##### Example

``` go
```
##### (*Var) Val 

``` go
func (v *Var) Val() interface{}
```

Val returns the current value of `v`.

##### Example

``` go
```
##### (*Var) Vars 

``` go
func (v *Var) Vars() []*Var
```

Vars converts and returns `v` as []Var.

##### Example

``` go
```
#### type Vars <-2.3.0

``` go
type Vars []*Var
```

Vars is a slice of *Var.

##### (Vars) Float32s <-2.3.0

``` go
func (vs Vars) Float32s() (s []float32)
```

Float32s converts and returns `vs` as []float32.

##### (Vars) Float64s <-2.3.0

``` go
func (vs Vars) Float64s() (s []float64)
```

Float64s converts and returns `vs` as []float64.

##### (Vars) Int16s <-2.3.0

``` go
func (vs Vars) Int16s() (s []int16)
```

Int16s converts and returns `vs` as []int16.

##### (Vars) Int32s <-2.3.0

``` go
func (vs Vars) Int32s() (s []int32)
```

Int32s converts and returns `vs` as []int32.

##### (Vars) Int64s <-2.3.0

``` go
func (vs Vars) Int64s() (s []int64)
```

Int64s converts and returns `vs` as []int64.

##### (Vars) Int8s <-2.3.0

``` go
func (vs Vars) Int8s() (s []int8)
```

Int8s converts and returns `vs` as []int8.

##### (Vars) Interfaces <-2.3.0

``` go
func (vs Vars) Interfaces() (s []interface{})
```

Interfaces converts and returns `vs` as []interface{}.

##### (Vars) Ints <-2.3.0

``` go
func (vs Vars) Ints() (s []int)
```

Ints converts and returns `vs` as []Int.

##### (Vars) Scan <-2.3.0

``` go
func (vs Vars) Scan(pointer interface{}, mapping ...map[string]string) error
```

Scan converts `vs` to []struct/[]*struct.

##### (Vars) Strings <-2.3.0

``` go
func (vs Vars) Strings() (s []string)
```

Strings converts and returns `vs` as []string.

##### (Vars) Uint16s <-2.3.0

``` go
func (vs Vars) Uint16s() (s []uint16)
```

Uint16s converts and returns `vs` as []uint16.

##### (Vars) Uint32s <-2.3.0

``` go
func (vs Vars) Uint32s() (s []uint32)
```

Uint32s converts and returns `vs` as []uint32.

##### (Vars) Uint64s <-2.3.0

``` go
func (vs Vars) Uint64s() (s []uint64)
```

Uint64s converts and returns `vs` as []uint64.

##### (Vars) Uint8s <-2.3.0

``` go
func (vs Vars) Uint8s() (s []uint8)
```

Uint8s converts and returns `vs` as []uint8.

##### (Vars) Uints <-2.3.0

``` go
func (vs Vars) Uints() (s []uint)
```

Uints converts and returns `vs` as []uint.