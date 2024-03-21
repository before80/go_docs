+++
title = "gutil"
date = 2024-03-21T18:00:04+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gutil

### Overview 

Package gutil provides utility functions.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func ComparatorByte 

``` go
func ComparatorByte(a, b interface{}) int
```

ComparatorByte provides a basic comparison on byte.

##### func ComparatorFloat32 

``` go
func ComparatorFloat32(a, b interface{}) int
```

ComparatorFloat32 provides a basic comparison on float32.

##### func ComparatorFloat64 

``` go
func ComparatorFloat64(a, b interface{}) int
```

ComparatorFloat64 provides a basic comparison on float64.

##### func ComparatorInt 

``` go
func ComparatorInt(a, b interface{}) int
```

ComparatorInt provides a basic comparison on int.

##### func ComparatorInt16 

``` go
func ComparatorInt16(a, b interface{}) int
```

ComparatorInt16 provides a basic comparison on int16.

##### func ComparatorInt32 

``` go
func ComparatorInt32(a, b interface{}) int
```

ComparatorInt32 provides a basic comparison on int32.

##### func ComparatorInt64 

``` go
func ComparatorInt64(a, b interface{}) int
```

ComparatorInt64 provides a basic comparison on int64.

##### func ComparatorInt8 

``` go
func ComparatorInt8(a, b interface{}) int
```

ComparatorInt8 provides a basic comparison on int8.

##### func ComparatorRune 

``` go
func ComparatorRune(a, b interface{}) int
```

ComparatorRune provides a basic comparison on rune.

##### func ComparatorString 

``` go
func ComparatorString(a, b interface{}) int
```

ComparatorString provides a fast comparison on strings.

##### func ComparatorTime 

``` go
func ComparatorTime(a, b interface{}) int
```

ComparatorTime provides a basic comparison on time.Time.

##### func ComparatorUint 

``` go
func ComparatorUint(a, b interface{}) int
```

ComparatorUint provides a basic comparison on uint.

##### func ComparatorUint16 

``` go
func ComparatorUint16(a, b interface{}) int
```

ComparatorUint16 provides a basic comparison on uint16.

##### func ComparatorUint32 

``` go
func ComparatorUint32(a, b interface{}) int
```

ComparatorUint32 provides a basic comparison on uint32.

##### func ComparatorUint64 

``` go
func ComparatorUint64(a, b interface{}) int
```

ComparatorUint64 provides a basic comparison on uint64.

##### func ComparatorUint8 

``` go
func ComparatorUint8(a, b interface{}) int
```

ComparatorUint8 provides a basic comparison on uint8.

##### func Copy <-2.1.0

``` go
func Copy(src interface{}) (dst interface{})
```

Copy returns a deep copy of v.

Copy is unable to copy unexported fields in a struct (lowercase field names). Unexported fields can't be reflected by the Go runtime and therefore they can't perform any data copies.

##### func Dump 

``` go
func Dump(values ...interface{})
```

Dump prints variables `values` to stdout with more manually readable.

##### func DumpJson <-2.4.2

``` go
func DumpJson(jsonContent string)
```

DumpJson pretty dumps json content to stdout.

##### func DumpTo 

``` go
func DumpTo(writer io.Writer, value interface{}, option DumpOption)
```

DumpTo writes variables `values` as a string in to `writer` with more manually readable

##### func DumpWithOption 

``` go
func DumpWithOption(value interface{}, option DumpOption)
```

DumpWithOption returns variables `values` as a string with more manually readable.

##### func DumpWithType 

``` go
func DumpWithType(values ...interface{})
```

DumpWithType acts like Dump, but with type information. Also see Dump.

##### func FillStructWithDefault <-2.5.4

``` go
func FillStructWithDefault(structPtr interface{}) error
```

FillStructWithDefault fills attributes of pointed struct with tag value from `default/d` tag . The parameter `structPtr` should be either type of *struct/[]*struct.

##### func GetOrDefaultAny <-2.2.0

``` go
func GetOrDefaultAny(def interface{}, param ...interface{}) interface{}
```

GetOrDefaultAny checks and returns value according whether parameter `param` available. It returns `param[0]` if it is available, or else it returns `def`.

##### func GetOrDefaultStr <-2.2.0

``` go
func GetOrDefaultStr(def string, param ...string) string
```

GetOrDefaultStr checks and returns value according whether parameter `param` available. It returns `param[0]` if it is available, or else it returns `def`.

##### func Go <-2.5.3

``` go
func Go(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
	recoverFunc func(ctx context.Context, exception error),
)
```

Go creates a new asynchronous goroutine function with specified recover function.

The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`. If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will throw to parent goroutine.

But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.

##### func IsEmpty 

``` go
func IsEmpty(value interface{}) bool
```

IsEmpty checks given `value` empty or not. It returns false if `value` is: integer(0), bool(false), slice/map(len=0), nil; or else returns true.

##### func IsTypeOf <-2.6.4

``` go
func IsTypeOf(value, valueInExpectType interface{}) bool
```

IsTypeOf checks and returns whether the type of `value` and `valueInExpectType` equal.

##### func ItemValue 

``` go
func ItemValue(item interface{}, key interface{}) (value interface{}, found bool)
```

ItemValue retrieves and returns its value of which name/attribute specified by `key`. The parameter `item` can be type of map/*map/struct/*struct.

##### func Keys 

``` go
func Keys(mapOrStruct interface{}) (keysOrAttrs []string)
```

Keys retrieves and returns the keys from given map or struct.

##### func ListItemValues 

``` go
func ListItemValues(list interface{}, key interface{}, subKey ...interface{}) (values []interface{})
```

ListItemValues retrieves and returns the elements of all item struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice.

The parameter `list` supports types like: []map[string]interface{} []map[string]sub-map []struct []struct:sub-struct Note that the sub-map/sub-struct makes sense only if the optional parameter `subKey` is given.

##### func ListItemValuesUnique 

``` go
func ListItemValuesUnique(list interface{}, key string, subKey ...interface{}) []interface{}
```

ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice.

##### func ListToMapByKey 

``` go
func ListToMapByKey(list []map[string]interface{}, key string) map[string]interface{}
```

ListToMapByKey converts `list` to a map[string]interface{} of which key is specified by `key`. Note that the item value may be type of slice.

##### func MapContains 

``` go
func MapContains(data map[string]interface{}, key string) (ok bool)
```

MapContains checks whether map `data` contains `key`.

##### func MapContainsPossibleKey 

``` go
func MapContainsPossibleKey(data map[string]interface{}, key string) bool
```

MapContainsPossibleKey checks if the given `key` is contained in given map `data`. It checks the key ignoring cases and symbols.

Note that this function might be of low performance.

##### func MapCopy 

``` go
func MapCopy(data map[string]interface{}) (copy map[string]interface{})
```

MapCopy does a shallow copy from map `data` to `copy` for most commonly used map type map[string]interface{}.

##### func MapDelete 

``` go
func MapDelete(data map[string]interface{}, keys ...string)
```

MapDelete deletes all `keys` from map `data`.

##### func MapMerge 

``` go
func MapMerge(dst map[string]interface{}, src ...map[string]interface{})
```

MapMerge merges all map from `src` to map `dst`.

##### func MapMergeCopy 

``` go
func MapMergeCopy(src ...map[string]interface{}) (copy map[string]interface{})
```

MapMergeCopy creates and returns a new map which merges all map from `src`.

##### func MapOmitEmpty 

``` go
func MapOmitEmpty(data map[string]interface{})
```

MapOmitEmpty deletes all empty values from given map.

##### func MapPossibleItemByKey 

``` go
func MapPossibleItemByKey(data map[string]interface{}, key string) (foundKey string, foundValue interface{})
```

MapPossibleItemByKey tries to find the possible key-value pair for given key ignoring cases and symbols.

Note that this function might be of low performance.

##### func MapToSlice 

``` go
func MapToSlice(data interface{}) []interface{}
```

MapToSlice converts map to slice of which all keys and values are its items. Eg: {"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]

##### func SliceCopy 

``` go
func SliceCopy(slice []interface{}) []interface{}
```

SliceCopy does a shallow copy of slice `data` for most commonly used slice type []interface{}.

##### func SliceDelete 

``` go
func SliceDelete(slice []interface{}, index int) (newSlice []interface{})
```

SliceDelete deletes an element at `index` and returns the new slice. It does nothing if the given `index` is invalid.

##### func SliceInsertAfter <-2.3.2

``` go
func SliceInsertAfter(slice []interface{}, index int, values ...interface{}) (newSlice []interface{})
```

SliceInsertAfter inserts the `values` to the back of `index` and returns a new slice.

##### Example

``` go
```
##### func SliceInsertBefore <-2.3.2

``` go
func SliceInsertBefore(slice []interface{}, index int, values ...interface{}) (newSlice []interface{})
```

SliceInsertBefore inserts the `values` to the front of `index` and returns a new slice.

##### Example

``` go
```
##### func SliceToMap 

``` go
func SliceToMap(slice interface{}) map[string]interface{}
```

SliceToMap converts slice type variable `slice` to `map[string]interface{}`. Note that if the length of `slice` is not an even number, it returns nil. Eg: ["K1", "v1", "K2", "v2"] => {"K1": "v1", "K2": "v2"} ["K1", "v1", "K2"] => nil

##### func SliceToMapWithColumnAsKey 

``` go
func SliceToMapWithColumnAsKey(slice interface{}, key interface{}) map[interface{}]interface{}
```

SliceToMapWithColumnAsKey converts slice type variable `slice` to `map[interface{}]interface{}` The value of specified column use as the key for returned map. Eg: SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K1") => {"v1": {"K1": "v1", "K2": 1}, "v2": {"K1": "v2", "K2": 2}} SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K2") => {1: {"K1": "v1", "K2": 1}, 2: {"K1": "v2", "K2": 2}}

##### func StructToSlice 

``` go
func StructToSlice(data interface{}) []interface{}
```

StructToSlice converts struct to slice of which all keys and values are its items. Eg: {"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]

##### func Throw 

``` go
func Throw(exception interface{})
```

Throw throws out an exception, which can be caught be TryCatch or recover.

##### func Try 

``` go
func Try(ctx context.Context, try func(ctx context.Context)) (err error)
```

Try implements try... logistics using internal panic...recover. It returns error if any exception occurs, or else it returns nil.

##### func TryCatch 

``` go
func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error))
```

TryCatch implements `try...catch..`. logistics using internal `panic...recover`. It automatically calls function `catch` if any exception occurs and passes the exception as an error. If `catch` is given nil, it ignores the panic from `try` and no panic will throw to parent goroutine.

But, note that, if function `catch` also throws panic, the current goroutine will panic.

##### func Values 

``` go
func Values(mapOrStruct interface{}) (values []interface{})
```

Values retrieves and returns the values from given map or struct.

### Types 

#### type Comparator 

``` go
type Comparator func(a, b interface{}) int
```

Comparator is a function that compare a and b, and returns the result as int.

Should return a number:

```
negative , if a < b
zero     , if a == b
positive , if a > b
```

#### type DumpOption 

``` go
type DumpOption struct {
	WithType     bool // WithType specifies dumping content with type information.
	ExportedOnly bool // Only dump Exported fields for structs.
}
```

DumpOption specifies the behavior of function Export.

#### type OriginTypeAndKindOutput <-2.3.0

``` go
type OriginTypeAndKindOutput = reflection.OriginTypeAndKindOutput
```

##### func OriginTypeAndKind <-2.3.0

``` go
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput)
```

OriginTypeAndKind retrieves and returns the original reflect type and kind.

#### type OriginValueAndKindOutput <-2.3.0

``` go
type OriginValueAndKindOutput = reflection.OriginValueAndKindOutput
```

##### func OriginValueAndKind <-2.3.0

``` go
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput)
```

OriginValueAndKind retrieves and returns the original reflect value and kind.