+++
title = "gstructs"
date = 2024-03-21T17:57:28+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gstructs

Package gstructs provides functions for struct information retrieving.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func FieldMap 

``` go
func FieldMap(in FieldMapInput) (map[string]Field, error)
```

FieldMap retrieves and returns struct field as map[name/tag]Field from `pointer`.

The parameter `pointer` should be type of struct/*struct.

The parameter `priority` specifies the priority tag array for retrieving from high to low. If it's given `nil`, it returns map[name]Field, of which the `name` is attribute name.

The parameter `recursive` specifies whether retrieving the fields recursively if the attribute is an embedded struct.

Note that it only retrieves the exported attributes with first letter upper-case from struct.

##### func ParseTag 

``` go
func ParseTag(tag string) map[string]string
```

ParseTag parses tag string into map. For example: ParseTag(`v:"required" p:"id" d:"1"`) => map[v:required p:id d:1].

##### func TagMapField 

``` go
func TagMapField(object interface{}, priority []string) (map[string]Field, error)
```

TagMapField retrieves struct tags as map[tag]Field from `pointer`, and returns it. The parameter `object` should be either type of struct/*struct/[]struct/[]*struct.

Note that, 1. It only retrieves the exported attributes with first letter upper-case from struct. 2. The parameter `priority` should be given, it only retrieves fields that has given tag. 3. If one field has no specified tag, it uses its field name as result map key.

##### func TagMapName 

``` go
func TagMapName(pointer interface{}, priority []string) (map[string]string, error)
```

TagMapName retrieves and returns struct tags as map[tag]attribute from `pointer`.

The parameter `pointer` should be type of struct/*struct.

Note that, 1. It only retrieves the exported attributes with first letter upper-case from struct. 2. The parameter `priority` should be given, it only retrieves fields that has given tag. 3. If one field has no specified tag, it uses its field name as result map key.

### Types 

#### type Field 

``` go
type Field struct {
	Value reflect.Value       // The underlying value of the field.
	Field reflect.StructField // The underlying field of the field.

	// Retrieved tag name. It depends TagValue.
	TagName string

	// Retrieved tag value.
	// There might be more than one tags in the field,
	// but only one can be retrieved according to calling function rules.
	TagValue string
}
```

Field contains information of a struct field .

##### func Fields 

``` go
func Fields(in FieldsInput) ([]Field, error)
```

Fields retrieves and returns the fields of `pointer` as slice.

##### func TagFields 

``` go
func TagFields(pointer interface{}, priority []string) ([]Field, error)
```

TagFields retrieves and returns struct tags as []Field from `pointer`.

The parameter `pointer` should be type of struct/*struct.

Note that, 1. It only retrieves the exported attributes with first letter upper-case from struct. 2. The parameter `priority` should be given, it only retrieves fields that has given tag.

##### (*Field) IsEmbedded 

``` go
func (f *Field) IsEmbedded() bool
```

IsEmbedded returns true if the given field is an anonymous field (embedded)

##### (*Field) IsEmpty <-2.5.4

``` go
func (f *Field) IsEmpty() bool
```

IsEmpty checks and returns whether the value of this Field is empty.

##### (*Field) IsExported 

``` go
func (f *Field) IsExported() bool
```

IsExported returns true if the given field is exported.

##### (*Field) IsNil <-2.5.4

``` go
func (f *Field) IsNil(traceSource ...bool) bool
```

IsNil checks and returns whether the value of this Field is nil.

##### (*Field) Kind 

``` go
func (f *Field) Kind() reflect.Kind
```

Kind returns the reflect.Kind for Value of Field `f`.

##### (*Field) Name 

``` go
func (f *Field) Name() string
```

Name returns the name of the given field.

##### (*Field) OriginalKind 

``` go
func (f *Field) OriginalKind() reflect.Kind
```

OriginalKind retrieves and returns the original reflect.Kind for Value of Field `f`.

##### (*Field) OriginalValue <-2.6.2

``` go
func (f *Field) OriginalValue() reflect.Value
```

OriginalValue retrieves and returns the original reflect.Value of Field `f`.

##### (*Field) Tag 

``` go
func (f *Field) Tag(key string) string
```

Tag returns the value associated with key in the tag string. If there is no such key in the tag, Tag returns the empty string.

##### (*Field) TagAdditional <-2.2.4

``` go
func (f *Field) TagAdditional() string
```

TagAdditional returns the most commonly used tag `additional/ad` value of the field.

##### (*Field) TagDefault <-2.2.4

``` go
func (f *Field) TagDefault() string
```

TagDefault returns the most commonly used tag `default/d` value of the field.

##### (*Field) TagDescription <-2.2.4

``` go
func (f *Field) TagDescription() string
```

TagDescription returns the most commonly used tag `description/des/dc` value of the field.

##### (*Field) TagExample <-2.2.4

``` go
func (f *Field) TagExample() string
```

TagExample returns the most commonly used tag `example/eg` value of the field.

##### (*Field) TagIn <-2.5.5

``` go
func (f *Field) TagIn() string
```

TagIn returns the most commonly used tag `in` value of the field.

##### (*Field) TagJsonName 

``` go
func (f *Field) TagJsonName() string
```

TagJsonName returns the `json` tag name string of the field.

##### (*Field) TagLookup 

``` go
func (f *Field) TagLookup(key string) (value string, ok bool)
```

TagLookup returns the value associated with key in the tag string. If the key is present in the tag the value (which may be empty) is returned. Otherwise, the returned value will be the empty string. The ok return value reports whether the value was explicitly set in the tag string. If the tag does not have the conventional format, the value returned by Lookup is unspecified.

##### (*Field) TagMap 

``` go
func (f *Field) TagMap() map[string]string
```

TagMap returns all the tag of the field along with its value string as map.

##### (*Field) TagParam <-2.2.4

``` go
func (f *Field) TagParam() string
```

TagParam returns the most commonly used tag `param/p` value of the field.

##### (*Field) TagPriorityName <-2.6.3

``` go
func (f *Field) TagPriorityName() string
```

TagPriorityName checks and returns tag name that matches the name item in `gtag.StructTagPriority`. It or else returns attribute field Name if it doesn't have a tag name by `gtag.StructsTagPriority`.

##### (*Field) TagStr 

``` go
func (f *Field) TagStr() string
```

TagStr returns the tag string of the field.

##### (*Field) TagSummary <-2.2.4

``` go
func (f *Field) TagSummary() string
```

TagSummary returns the most commonly used tag `summary/sum/sm` value of the field.

##### (*Field) TagValid <-2.2.4

``` go
func (f *Field) TagValid() string
```

TagValid returns the most commonly used tag `valid/v` value of the field.

##### (*Field) Type 

``` go
func (f *Field) Type() Type
```

Type returns the type of the given field. Note that this Type is not reflect.Type. If you need reflect.Type, please use Field.Type().Type.

#### type FieldMapInput 

``` go
type FieldMapInput struct {
	// Pointer should be type of struct/*struct.
	// TODO this attribute name is not suitable, which would make confuse.
	Pointer interface{}

	// PriorityTagArray specifies the priority tag array for retrieving from high to low.
	// If it's given `nil`, it returns map[name]Field, of which the `name` is attribute name.
	PriorityTagArray []string

	// RecursiveOption specifies the way retrieving the fields recursively if the attribute
	// is an embedded struct. It is RecursiveOptionNone in default.
	RecursiveOption RecursiveOption
}
```

FieldMapInput is the input parameter struct type for function FieldMap.

#### type FieldsInput 

``` go
type FieldsInput struct {
	// Pointer should be type of struct/*struct.
	// TODO this attribute name is not suitable, which would make confuse.
	Pointer interface{}

	// RecursiveOption specifies the way retrieving the fields recursively if the attribute
	// is an embedded struct. It is RecursiveOptionNone in default.
	RecursiveOption RecursiveOption
}
```

FieldsInput is the input parameter struct type for function Fields.

#### type RecursiveOption <-2.2.0

``` go
type RecursiveOption int
const (
	RecursiveOptionNone          RecursiveOption = iota // No recursively retrieving fields as map if the field is an embedded struct.
	RecursiveOptionEmbedded                             // Recursively retrieving fields as map if the field is an embedded struct.
	RecursiveOptionEmbeddedNoTag                        // Recursively retrieving fields as map if the field is an embedded struct and the field has no tag.
)
```

#### type Type 

``` go
type Type struct {
	reflect.Type
}
```

Type wraps reflect.Type for additional features.

##### func StructType 

``` go
func StructType(object interface{}) (*Type, error)
```

StructType retrieves and returns the struct Type of specified struct/*struct. The parameter `object` should be either type of struct/*struct/[]struct/[]*struct.

##### (Type) FieldKeys 

``` go
func (t Type) FieldKeys() []string
```

FieldKeys returns the keys of current struct/map.

##### (Type) Signature 

``` go
func (t Type) Signature() string
```

Signature returns a unique string as this type.