+++
title = "gstructs"
date = 2024-03-21T17:57:28+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gstructs](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gstructs)

Package gstructs provides functions for struct information retrieving.

​	软件包 gstructs 提供用于结构信息检索的函数。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func FieldMap

```go
func FieldMap(in FieldMapInput) (map[string]Field, error)
```

FieldMap retrieves and returns struct field as map[name/tag]Field from `pointer`.

​	FieldMap 检索并返回 struct 字段作为 map[name/tag]Field from `pointer` .

The parameter `pointer` should be type of struct/*struct.

​	参数 `pointer` 应为 struct/*struct 类型。

The parameter `priority` specifies the priority tag array for retrieving from high to low. If it’s given `nil`, it returns map[name]Field, of which the `name` is attribute name.

​	该参数 `priority` 指定从高到低检索的优先级标记数组。如果给定 `nil` ，则返回 map[name]Field，其中 `name` is 属性名称。

The parameter `recursive` specifies whether retrieving the fields recursively if the attribute is an embedded struct.

​	该参数 `recursive` 指定如果属性是嵌入结构，则是否以递归方式检索字段。

Note that it only retrieves the exported attributes with first letter upper-case from struct.

​	请注意，它仅从 struct 中检索首字母大写的导出属性。

#### func ParseTag

```go
func ParseTag(tag string) map[string]string
```

ParseTag parses tag string into map. For example: ParseTag(`v:"required" p:"id" d:"1"`) => map[v:required p:id d:1].

​	ParseTag 将标签字符串解析为地图。例如：ParseTag（ `v:"required" p:"id" d:"1"` ） => map[v：required p：id d：1]。

#### func TagMapField

```go
func TagMapField(object interface{}, priority []string) (map[string]Field, error)
```

TagMapField retrieves struct tags as map[tag]Field from `pointer`, and returns it. The parameter `object` should be either type of struct/*struct/[]struct/[]*struct.

​	TagMapField 从 `pointer` 中检索 map[tag]Field 作为 struct 标签，并返回它。参数 `object` 应为 struct/*struct/[]struct/[]*struct 的任一类型。

Note that, 1. It only retrieves the exported attributes with first letter upper-case from struct. 2. The parameter `priority` should be given, it only retrieves fields that has given tag. 3. If one field has no specified tag, it uses its field name as result map key.

​	请注意，1.它仅从 struct 中检索首字母大写的导出属性。2. `priority` 应该给出参数，它只检索已给出标签的字段。3. 如果一个字段没有指定的标签，则使用其字段名称作为结果映射键。

#### func TagMapName

```go
func TagMapName(pointer interface{}, priority []string) (map[string]string, error)
```

TagMapName retrieves and returns struct tags as map[tag]attribute from `pointer`.

​	TagMapName 检索并返回结构标签作为 map[tag]属性。 `pointer`

The parameter `pointer` should be type of struct/*struct.

​	参数 `pointer` 应为 struct/*struct 类型。

Note that, 1. It only retrieves the exported attributes with first letter upper-case from struct. 2. The parameter `priority` should be given, it only retrieves fields that has given tag. 3. If one field has no specified tag, it uses its field name as result map key.

​	请注意，1.它仅从 struct 中检索首字母大写的导出属性。2. `priority` 应该给出参数，它只检索已给出标签的字段。3. 如果一个字段没有指定的标签，则使用其字段名称作为结果映射键。

## 类型

### type Field

```go
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

​	字段包含结构字段的信息。

#### func Fields

```go
func Fields(in FieldsInput) ([]Field, error)
```

Fields retrieves and returns the fields of `pointer` as slice.

​	字段检索并返回 as slice 的 `pointer` 字段。

#### func TagFields

```go
func TagFields(pointer interface{}, priority []string) ([]Field, error)
```

TagFields retrieves and returns struct tags as []Field from `pointer`.

​	TagFields 检索结构标签并将其返回为 []Field from `pointer` .

The parameter `pointer` should be type of struct/*struct.

​	参数 `pointer` 应为 struct/*struct 类型。

Note that, 1. It only retrieves the exported attributes with first letter upper-case from struct. 2. The parameter `priority` should be given, it only retrieves fields that has given tag.

​	请注意，1.它仅从 struct 中检索首字母大写的导出属性。2. `priority` 应该给出参数，它只检索已给出标签的字段。

#### (*Field) IsEmbedded

```go
func (f *Field) IsEmbedded() bool
```

IsEmbedded returns true if the given field is an anonymous field (embedded)

​	如果给定字段是匿名字段（嵌入式），则 IsEmbedded 返回 true

#### (*Field) IsEmpty

```go
func (f *Field) IsEmpty() bool
```

IsEmpty checks and returns whether the value of this Field is empty.

​	IsEmpty 检查并返回此字段的值是否为空。

#### (*Field) IsExported

```go
func (f *Field) IsExported() bool
```

IsExported returns true if the given field is exported.

​	如果导出给定字段，则 IsExported 返回 true。

#### (*Field) IsNil

```go
func (f *Field) IsNil(traceSource ...bool) bool
```

IsNil checks and returns whether the value of this Field is nil.

​	IsNil 检查并返回此字段的值是否为 nil。

#### (*Field) Kind

```go
func (f *Field) Kind() reflect.Kind
```

Kind returns the reflect.Kind for Value of Field `f`.

​	Kind 返回反射。字段 `f` 值的种类。

#### (*Field) Name

```go
func (f *Field) Name() string
```

Name returns the name of the given field.

​	Name 返回给定字段的名称。

#### (*Field) OriginalKind

```go
func (f *Field) OriginalKind() reflect.Kind
```

OriginalKind retrieves and returns the original reflect.Kind for Value of Field `f`.

​	OriginalKind 检索并返回原始反射。字段 `f` 值的种类。

#### (*Field) OriginalValue

```go
func (f *Field) OriginalValue() reflect.Value
```

OriginalValue retrieves and returns the original reflect.Value of Field `f`.

​	OriginalValue 检索并返回原始反射。字段 `f` 的值。

#### (*Field) Tag

```go
func (f *Field) Tag(key string) string
```

Tag returns the value associated with key in the tag string. If there is no such key in the tag, Tag returns the empty string.

​	Tag 返回与标签字符串中的键关联的值。如果标签中没有此类键，则 Tag 返回空字符串。

#### (*Field) TagAdditional

```go
func (f *Field) TagAdditional() string
```

TagAdditional returns the most commonly used tag `additional/ad` value of the field.

​	TagAdditional 返回字段最常用的标记 `additional/ad` 值。

#### (*Field) TagDefault

```go
func (f *Field) TagDefault() string
```

TagDefault returns the most commonly used tag `default/d` value of the field.

​	TagDefault 返回字段最常用的标记 `default/d` 值。

#### (*Field) TagDescription

```go
func (f *Field) TagDescription() string
```

TagDescription returns the most commonly used tag `description/des/dc` value of the field.

​	TagDescription 返回字段最常用的标记 `description/des/dc` 值。

#### (*Field) TagExample

```go
func (f *Field) TagExample() string
```

TagExample returns the most commonly used tag `example/eg` value of the field.

​	TagExample 返回字段最常用的标记 `example/eg` 值。

#### (*Field) TagIn

```go
func (f *Field) TagIn() string
```

TagIn returns the most commonly used tag `in` value of the field.

​	TagIn 返回字段最常用的标记 `in` 值。

#### (*Field) TagJsonName

```go
func (f *Field) TagJsonName() string
```

TagJsonName returns the `json` tag name string of the field.

​	TagJsonName 返回字段的 `json` 标记名称字符串。

#### (*Field) TagLookup

```go
func (f *Field) TagLookup(key string) (value string, ok bool)
```

TagLookup returns the value associated with key in the tag string. If the key is present in the tag the value (which may be empty) is returned. Otherwise, the returned value will be the empty string. The ok return value reports whether the value was explicitly set in the tag string. If the tag does not have the conventional format, the value returned by Lookup is unspecified.

​	TagLookup 返回与标记字符串中的键关联的值。如果标记中存在键，则返回值（可能为空）。否则，返回的值将为空字符串。ok 返回值报告是否在标记字符串中显式设置了该值。如果标记没有常规格式，则 Lookup 返回的值未指定。

#### (*Field) TagMap

```go
func (f *Field) TagMap() map[string]string
```

TagMap returns all the tag of the field along with its value string as map.

​	TagMap 返回字段的所有标记及其值字符串作为 map。

#### (*Field) TagParam

```go
func (f *Field) TagParam() string
```

TagParam returns the most commonly used tag `param/p` value of the field.

​	TagParam 返回字段最常用的标记 `param/p` 值。

#### (*Field) TagPriorityName

```go
func (f *Field) TagPriorityName() string
```

TagPriorityName checks and returns tag name that matches the name item in `gtag.StructTagPriority`. It or else returns attribute field Name if it doesn’t have a tag name by `gtag.StructsTagPriority`.

​	TagPriorityName 检查并返回与 中 `gtag.StructTagPriority` 的名称项匹配的标记名称。如果它没有 by 的 `gtag.StructsTagPriority` 标签名称，则返回属性字段 Name。

#### (*Field) TagStr

```go
func (f *Field) TagStr() string
```

TagStr returns the tag string of the field.

​	TagStr 返回字段的标记字符串。

#### (*Field) TagSummary

```go
func (f *Field) TagSummary() string
```

TagSummary returns the most commonly used tag `summary/sum/sm` value of the field.

​	TagSummary 返回字段最常用的标记 `summary/sum/sm` 值。

#### (*Field) TagValid

```go
func (f *Field) TagValid() string
```

TagValid returns the most commonly used tag `valid/v` value of the field.

​	TagValid 返回字段最常用的标记 `valid/v` 值。

#### (*Field) Type

```go
func (f *Field) Type() Type
```

Type returns the type of the given field. Note that this Type is not reflect.Type. If you need reflect.Type, please use Field.Type().Type.

​	Type 返回给定字段的类型。请注意，此类型不反映。类型。如果你需要反思。类型，请使用 Field.Type（）。类型。

### type FieldMapInput

```go
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

​	FieldMapInput 是函数 FieldMap 的输入参数结构类型。

### type FieldsInput

```go
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

​	FieldsInput 是函数 Fields 的输入参数结构类型。

### type RecursiveOption <-2.2.0

```go
type RecursiveOption int
const (
	RecursiveOptionNone          RecursiveOption = iota // No recursively retrieving fields as map if the field is an embedded struct.
	RecursiveOptionEmbedded                             // Recursively retrieving fields as map if the field is an embedded struct.
	RecursiveOptionEmbeddedNoTag                        // Recursively retrieving fields as map if the field is an embedded struct and the field has no tag.
)
```

### type Type

```go
type Type struct {
	reflect.Type
}
```

Type wraps reflect.Type for additional features.

​	类型换行反映。键入其他功能。

#### func StructType

```go
func StructType(object interface{}) (*Type, error)
```

StructType retrieves and returns the struct Type of specified struct/*struct. The parameter `object` should be either type of struct/*struct/[]struct/[]*struct.

​	StructType 检索并返回指定 struct/*struct 的 struct Type。参数 `object` 应为 struct/*struct/[]struct/[]*struct 的任一类型。

#### (Type) FieldKeys

```go
func (t Type) FieldKeys() []string
```

FieldKeys returns the keys of current struct/map.

​	FieldKeys 返回当前 struct/map 的键。

#### (Type) Signature

```go
func (t Type) Signature() string
```

Signature returns a unique string as this type.

​	Signature 返回一个唯一的字符串作为此类型。