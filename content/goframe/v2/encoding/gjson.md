+++
title = "gjson"
date = 2024-03-21T17:49:47+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gjson](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gjson)

Package gjson provides convenient API for JSON/XML/INI/YAML/TOML data handling.

​	软件包 gjson 为 JSON/XML/INI/YAML/TOML 数据处理提供了方便的 API。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Decode

```go
func Decode(data interface{}, options ...Options) (interface{}, error)
```

Decode decodes json format `data` to golang variable. The parameter `data` can be either bytes or string type.

​	Decode 将 json 格式 `data` 解码为 golang 变量。参数 `data` 可以是字节类型，也可以是字符串类型。

##### Example

``` go
```

#### func DecodeTo

```go
func DecodeTo(data interface{}, v interface{}, options ...Options) (err error)
```

DecodeTo decodes json format `data` to specified golang variable `v`. The parameter `data` can be either bytes or string type. The parameter `v` should be a pointer type.

​	DecodeTo 将 json 格式 `data` 解码为指定的 golang 变量 `v` 。参数 `data` 可以是字节类型，也可以是字符串类型。该参数 `v` 应为指针类型。

##### Example

``` go
```

#### func Encode

```go
func Encode(value interface{}) ([]byte, error)
```

Encode encodes any golang variable `value` to JSON bytes.

​	Encode 将任何 golang 变量编码 `value` 为 JSON 字节。

##### Example

``` go
```

#### func EncodeString

```go
func EncodeString(value interface{}) (string, error)
```

EncodeString encodes any golang variable `value` to JSON string.

​	EncodeString 将任何 golang 变量编码 `value` 为 JSON 字符串。

##### Example

``` go
```

#### func IsValidDataType

```go
func IsValidDataType(dataType ContentType) bool
```

IsValidDataType checks and returns whether given `dataType` a valid data type for loading.

​	IsValidDataType 检查并返回是否给定 `dataType` 了用于加载的有效数据类型。

##### Example

``` go
```

#### func Marshal

```go
func Marshal(v interface{}) (marshaledBytes []byte, err error)
```

Marshal is alias of Encode in order to fit the habit of json.Marshal/Unmarshal functions.

​	Marshal 是 Encode 的别名，以适应 json 的习惯。元帅/非元组函数。

##### Example

``` go
```

#### func MarshalIndent

```go
func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error)
```

MarshalIndent is alias of json.MarshalIndent in order to fit the habit of json.MarshalIndent function.

​	MarshalIndent 是 json 的别名。MarshalIndent 以适应 json 的习惯。MarshalIndent 函数。

##### Example

``` go
```

#### func MustEncode

```go
func MustEncode(value interface{}) []byte
```

MustEncode performs as Encode, but it panics if any error occurs.

​	MustEncode 以 Encode 的形式执行，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### func MustEncodeString

```go
func MustEncodeString(value interface{}) string
```

MustEncodeString encodes any golang variable `value` to JSON string. It panics if any error occurs.

​	MustEncodeString 将任何 golang 变量编码 `value` 为 JSON 字符串。如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### func Unmarshal

```go
func Unmarshal(data []byte, v interface{}) (err error)
```

Unmarshal is alias of DecodeTo in order to fit the habit of json.Marshal/Unmarshal functions.

​	Unmarshal 是 DecodeTo 的别名，以适应 json 的习惯。元帅/非元组函数。

##### Example

``` go
```

#### func Valid

```go
func Valid(data interface{}) bool
```

Valid checks whether `data` is a valid JSON data type. The parameter `data` specifies the json format data, which can be either bytes or string type.

​	Valid 检查是否 `data` 为有效的 JSON 数据类型。该参数 `data` 指定 json 格式数据，可以是字节或字符串类型。

##### Example

``` go
```

## 类型

### type ContentType <-2.2.0

```go
type ContentType string
const (
	ContentTypeJson       ContentType = `json`
	ContentTypeJs         ContentType = `js`
	ContentTypeXml        ContentType = `xml`
	ContentTypeIni        ContentType = `ini`
	ContentTypeYaml       ContentType = `yaml`
	ContentTypeYml        ContentType = `yml`
	ContentTypeToml       ContentType = `toml`
	ContentTypeProperties ContentType = `properties`
)
```

### type Json

```go
type Json struct {
	// contains filtered or unexported fields
}
```

Json is the customized JSON struct.

​	Json 是自定义的 JSON 结构。

#### func DecodeToJson

```go
func DecodeToJson(data interface{}, options ...Options) (*Json, error)
```

DecodeToJson codes json format `data` to a Json object. The parameter `data` can be either bytes or string type.

​	DecodeToJson 将 json 格式 `data` 编码为 Json 对象。参数 `data` 可以是字节类型，也可以是字符串类型。

##### Example

``` go
```

#### func Load

```go
func Load(path string, safe ...bool) (*Json, error)
```

Load loads content from specified file `path`, and creates a Json object from its content.

​	Load 从指定文件 `path` 加载内容，并从其内容创建 Json 对象。

##### Example

``` go
```

#### func LoadContent

```go
func LoadContent(data interface{}, safe ...bool) (*Json, error)
```

LoadContent creates a Json object from given content, it checks the data type of `content` automatically, supporting data content type as follows: JSON, XML, INI, YAML and TOML.

​	LoadContent 从给定的内容创建一个 Json 对象，它会自动检查数据 `content` 类型，支持的数据内容类型如下：JSON、XML、INI、YAML 和 TOML。

##### Example

``` go
```

#### func LoadContentType

```go
func LoadContentType(dataType ContentType, data interface{}, safe ...bool) (*Json, error)
```

LoadContentType creates a Json object from given type and content, supporting data content type as follows: JSON, XML, INI, YAML and TOML.

​	LoadContentType 从给定的类型和内容创建一个 Json 对象，支持数据内容类型如下：JSON、XML、INI、YAML 和 TOML。

##### Example

``` go
```

#### func LoadIni

```go
func LoadIni(data interface{}, safe ...bool) (*Json, error)
```

LoadIni creates a Json object from given INI format content.

​	LoadIni 从给定的 INI 格式内容创建 Json 对象。

##### Example

``` go
```

#### func LoadJson

```go
func LoadJson(data interface{}, safe ...bool) (*Json, error)
```

LoadJson creates a Json object from given JSON format content.

​	LoadJson 从给定的 JSON 格式内容创建 Json 对象。

##### Example

``` go
```

#### func LoadProperties <-2.1.0

```go
func LoadProperties(data interface{}, safe ...bool) (*Json, error)
```

LoadProperties creates a Json object from given TOML format content.

​	LoadProperties 从给定的 TOML 格式内容创建 Json 对象。

#### func LoadToml

```go
func LoadToml(data interface{}, safe ...bool) (*Json, error)
```

LoadToml creates a Json object from given TOML format content.

​	LoadToml 从给定的 TOML 格式内容创建一个 Json 对象。

##### Example

``` go
```

#### func LoadWithOptions <-2.1.0

```go
func LoadWithOptions(data interface{}, options Options) (*Json, error)
```

LoadWithOptions creates a Json object from given JSON format content and options.

​	LoadWithOptions 从给定的 JSON 格式内容和选项创建 Json 对象。

#### func LoadXml

```go
func LoadXml(data interface{}, safe ...bool) (*Json, error)
```

LoadXml creates a Json object from given XML format content.

​	LoadXml 从给定的 XML 格式内容创建 Json 对象。

##### Example

``` go
```

#### func LoadYaml

```go
func LoadYaml(data interface{}, safe ...bool) (*Json, error)
```

LoadYaml creates a Json object from given YAML format content.

​	LoadYaml 从给定的 YAML 格式内容创建 Json 对象。

##### Example

``` go
```

#### func New

```go
func New(data interface{}, safe ...bool) *Json
```

New creates a Json object with any variable type of `data`, but `data` should be a map or slice for data access reason, or it will make no sense.

​	New 创建一个具有任何变量类型的 `data` 的 Json 对象，但 `data` 出于数据访问的原因，它应该是映射或切片，否则将毫无意义。

The parameter `safe` specifies whether using this Json object in concurrent-safe context, which is false in default.

​	该参数 `safe` 指定是否在并发安全上下文中使用此 Json 对象，默认为 false。

##### Example

``` go
```

#### func NewWithOptions

```go
func NewWithOptions(data interface{}, options Options) *Json
```

NewWithOptions creates a Json object with any variable type of `data`, but `data` should be a map or slice for data access reason, or it will make no sense.

​	NewWithOptions 创建一个具有任何变量类型的 `data` 的 Json 对象，但 `data` 出于数据访问原因，它应该是映射或切片，否则将毫无意义。

##### Example

``` go
```

#### func NewWithTag

```go
func NewWithTag(data interface{}, tags string, safe ...bool) *Json
```

NewWithTag creates a Json object with any variable type of `data`, but `data` should be a map or slice for data access reason, or it will make no sense.

​	NewWithTag 创建一个具有任何变量类型 的 `data` Json 对象，但 `data` 出于数据访问原因，它应该是映射或切片，否则将毫无意义。

The parameter `tags` specifies priority tags for struct conversion to map, multiple tags joined with char ‘,’.

​	该参数 `tags` 指定用于将结构转换为映射的优先级标签，以及使用 char '，' 连接的多个标签。

The parameter `safe` specifies whether using this Json object in concurrent-safe context, which is false in default.

​	该参数 `safe` 指定是否在并发安全上下文中使用此 Json 对象，默认为 false。

##### Example

``` go
```

#### (*Json) Append

```go
func (j *Json) Append(pattern string, value interface{}) error
```

Append appends value to the value by specified `pattern`. The target value by `pattern` should be type of slice.

​	Append 将 value 追加到指定的 `pattern` 值。目标值应 `pattern` 为切片类型。

##### Example

``` go
```

#### (*Json) Array

```go
func (j *Json) Array() []interface{}
```

Array converts current Json object to []interface{}. It returns nil if fails.

​	数组将当前 Json 对象转换为 []interface{}。如果失败，则返回 nil。

##### Example

``` go
```

#### (*Json) Contains

```go
func (j *Json) Contains(pattern string) bool
```

Contains checks whether the value by specified `pattern` exist.

​	包含检查指定的 `pattern` 值是否存在。

##### Example

``` go
```

#### (*Json) Dump

```go
func (j *Json) Dump()
```

Dump prints current Json object with more manually readable.

​	转储打印当前 Json 对象，具有更多的手动可读性。

##### Example

``` go
```

#### (*Json) Get

```go
func (j *Json) Get(pattern string, def ...interface{}) *gvar.Var
```

Get retrieves and returns value by specified `pattern`. It returns all values of current Json object if `pattern` is given “.”. It returns nil if no value found by `pattern`.

​	按指定 `pattern` 获取检索和返回值。如果 `pattern` 给定“.”，则返回当前 Json 对象的所有值。如果 未找到 的 `pattern` 值，则返回 nil。

We can also access slice item by its index number in `pattern` like: “list.10”, “array.0.name”, “array.0.1.id”.

​	我们还可以通过索引号访问切片项目， `pattern` 例如：“list.10”、“array.0.name”、“array.0.1.id”。

It returns a default value specified by `def` if value for `pattern` is not found.

​	它返回由 `def` if value for 指定的默认值 if value `pattern` is not found。

##### Example

``` go
```

#### (*Json) GetJson

```go
func (j *Json) GetJson(pattern string, def ...interface{}) *Json
```

GetJson gets the value by specified `pattern`, and converts it to an un-concurrent-safe Json object.

​	GetJson 按指定 `pattern` 获取值，并将其转换为非并发安全的 Json 对象。

##### Example

``` go
```

#### (*Json) GetJsonMap

```go
func (j *Json) GetJsonMap(pattern string, def ...interface{}) map[string]*Json
```

GetJsonMap gets the value by specified `pattern`, and converts it to a map of un-concurrent-safe Json object.

​	GetJsonMap 按 specified `pattern` 获取值，并将其转换为非并发安全 Json 对象的映射。

##### Example

``` go
```

#### (*Json) GetJsons

```go
func (j *Json) GetJsons(pattern string, def ...interface{}) []*Json
```

GetJsons gets the value by specified `pattern`, and converts it to a slice of un-concurrent-safe Json object.

​	GetJsons 按 specified `pattern` 获取值，并将其转换为非并发安全 Json 对象的切片。

##### Example

``` go
```

#### (*Json) Interface

```go
func (j *Json) Interface() interface{}
```

Interface returns the json value.

​	Interface 返回 json 值。

##### Example

``` go
```

#### (*Json) Interfaces

```go
func (j *Json) Interfaces() []interface{}
```

Interfaces implements interface function Interfaces().

​	Interfaces 实现接口函数 Interfaces（）。

##### Example

``` go
```

#### (*Json) IsNil

```go
func (j *Json) IsNil() bool
```

IsNil checks whether the value pointed by `j` is nil.

​	IsNil 检查指向的 `j` 值是否为 nil。

##### Example

``` go
```

#### (*Json) Len

```go
func (j *Json) Len(pattern string) int
```

Len returns the length/size of the value by specified `pattern`. The target value by `pattern` should be type of slice or map. It returns -1 if the target value is not found, or its type is invalid.

​	Len 按指定 `pattern` 返回值的长度/大小。目标 `pattern` 值应为切片或地图的类型。如果未找到目标值或其类型无效，则返回 -1。

##### Example

``` go
```

#### (*Json) Map

```go
func (j *Json) Map() map[string]interface{}
```

Map converts current Json object to map[string]interface{}. It returns nil if fails.

​	Map 将当前 Json 对象转换为 map[string]interface{}。如果失败，则返回 nil。

##### Example

``` go
```

#### (*Json) MapStrAny

```go
func (j *Json) MapStrAny() map[string]interface{}
```

MapStrAny implements interface function MapStrAny().

​	MapStrAny 实现接口函数 MapStrAny（）。

##### Example

``` go
```

#### (Json) MarshalJSON

```go
func (j Json) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

##### Example

``` go
```

#### (*Json) MustAppend

```go
func (j *Json) MustAppend(pattern string, value interface{})
```

MustAppend performs as Append, but it panics if any error occurs.

​	MustAppend 以 Append 的形式执行，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Json) MustRemove

```go
func (j *Json) MustRemove(pattern string)
```

MustRemove performs as Remove, but it panics if any error occurs.

​	MustRemove 以 Remove 的形式执行，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Json) MustSet

```go
func (j *Json) MustSet(pattern string, value interface{})
```

MustSet performs as Set, but it panics if any error occurs.

​	MustSet 以 Set 的形式执行，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Json) MustToIni

```go
func (j *Json) MustToIni() []byte
```

##### Example

``` go
```

#### (*Json) MustToIniString

```go
func (j *Json) MustToIniString() string
```

MustToIniString .

​	MustToIniString 中。

##### Example

``` go
```

#### (*Json) MustToJson

```go
func (j *Json) MustToJson() []byte
```

##### Example

``` go
```

#### (*Json) MustToJsonIndent

```go
func (j *Json) MustToJsonIndent() []byte
```

##### Example

``` go
```

#### (*Json) MustToJsonIndentString

```go
func (j *Json) MustToJsonIndentString() string
```

##### Example

``` go
```

#### (*Json) MustToJsonString

```go
func (j *Json) MustToJsonString() string
```

##### Example

``` go
```

#### (*Json) MustToProperties

```go
func (j *Json) MustToProperties() []byte
```

##### Example

``` go
```

#### (*Json) MustToPropertiesString

```go
func (j *Json) MustToPropertiesString() string
```

MustTopropertiesString

##### Example

``` go
```

#### (*Json) MustToToml

```go
func (j *Json) MustToToml() []byte
```

##### Example

``` go
```

#### (*Json) MustToTomlString

```go
func (j *Json) MustToTomlString() string
```

##### Example

``` go
```

#### (*Json) MustToXml

```go
func (j *Json) MustToXml(rootTag ...string) []byte
```

##### Example

``` go
```

#### (*Json) MustToXmlIndent

```go
func (j *Json) MustToXmlIndent(rootTag ...string) []byte
```

##### Example

``` go
```

#### (*Json) MustToXmlIndentString

```go
func (j *Json) MustToXmlIndentString(rootTag ...string) string
```

##### Example

``` go
```

#### (*Json) MustToXmlString

```go
func (j *Json) MustToXmlString(rootTag ...string) string
```

##### Example

``` go
```

#### (*Json) MustToYaml

```go
func (j *Json) MustToYaml() []byte
```

##### Example

``` go
```

#### (*Json) MustToYamlString

```go
func (j *Json) MustToYamlString() string
```

##### Example

``` go
```

#### (*Json) Remove

```go
func (j *Json) Remove(pattern string) error
```

Remove deletes value with specified `pattern`. It supports hierarchical data access by char separator, which is ‘.’ in default.

​	删除具有指定 `pattern` .它支持通过字符分隔符进行分层数据访问，默认为 '..。

##### Example

``` go
```

#### (*Json) Scan

```go
func (j *Json) Scan(pointer interface{}, mapping ...map[string]string) error
```

Scan automatically calls Struct or Structs function according to the type of parameter `pointer` to implement the converting.

​	Scan 根据参数类型自动调用 Struct 或 Structs 函数 `pointer` 来实现转换。

##### Example

``` go
```

#### (*Json) Set

```go
func (j *Json) Set(pattern string, value interface{}) error
```

Set sets value with specified `pattern`. It supports hierarchical data access by char separator, which is ‘.’ in default.

​	使用指定的 `pattern` 设置值。它支持通过字符分隔符进行分层数据访问，默认为 '..。

##### Example

``` go
```

#### (*Json) SetSplitChar

```go
func (j *Json) SetSplitChar(char byte)
```

SetSplitChar sets the separator char for hierarchical data access.

​	SetSplitChar 设置分层数据访问的分隔符 char。

##### Example

``` go
```

#### (*Json) SetViolenceCheck

```go
func (j *Json) SetViolenceCheck(enabled bool)
```

SetViolenceCheck enables/disables violence check for hierarchical data access.

​	SetViolenceCheck 启用/禁用分层数据访问的暴力检查。

##### Example

``` go
```

#### (*Json) String

```go
func (j *Json) String() string
```

String returns current Json object as string.

​	String 以字符串形式返回当前 Json 对象。

#### (*Json) ToIni

```go
func (j *Json) ToIni() ([]byte, error)
```

ToIni json to ini

​	ToIni json 到 ini

##### Example

``` go
```

#### (*Json) ToIniString

```go
func (j *Json) ToIniString() (string, error)
```

ToIniString ini to string

​	ToIniString ini 到字符串

##### Example

``` go
```

#### (*Json) ToJson

```go
func (j *Json) ToJson() ([]byte, error)
```

##### Example

``` go
```

#### (*Json) ToJsonIndent

```go
func (j *Json) ToJsonIndent() ([]byte, error)
```

##### Example

``` go
```

#### (*Json) ToJsonIndentString

```go
func (j *Json) ToJsonIndentString() (string, error)
```

##### Example

``` go
```

#### (*Json) ToJsonString

```go
func (j *Json) ToJsonString() (string, error)
```

##### Example

``` go
```

#### (*Json) ToProperties

```go
func (j *Json) ToProperties() ([]byte, error)
```

======================================================================== properties ======================================================================== Toproperties json to properties

​	========================================================================属性======================================================================== Toproperties json 到属性

##### Example

``` go
```

#### (*Json) ToPropertiesString

```go
func (j *Json) ToPropertiesString() (string, error)
```

TopropertiesString properties to string

​	TopropertiesString 属性到字符串

##### Example

``` go
```

#### (*Json) ToToml

```go
func (j *Json) ToToml() ([]byte, error)
```

##### Example

``` go
```

#### (*Json) ToTomlString

```go
func (j *Json) ToTomlString() (string, error)
```

##### Example

``` go
```

#### (*Json) ToXml

```go
func (j *Json) ToXml(rootTag ...string) ([]byte, error)
```

##### Example

``` go
```

#### (*Json) ToXmlIndent

```go
func (j *Json) ToXmlIndent(rootTag ...string) ([]byte, error)
```

##### Example

``` go
```

#### (*Json) ToXmlIndentString

```go
func (j *Json) ToXmlIndentString(rootTag ...string) (string, error)
```

##### Example

``` go
```

#### (*Json) ToXmlString

```go
func (j *Json) ToXmlString(rootTag ...string) (string, error)
```

##### Example

``` go
```

#### (*Json) ToYaml

```go
func (j *Json) ToYaml() ([]byte, error)
```

##### Example

``` go
```

#### (*Json) ToYamlIndent

```go
func (j *Json) ToYamlIndent(indent string) ([]byte, error)
```

##### Example

``` go
```

#### (*Json) ToYamlString

```go
func (j *Json) ToYamlString() (string, error)
```

##### Example

``` go
```

#### (*Json) UnmarshalJSON

```go
func (j *Json) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*Json) UnmarshalValue

```go
func (j *Json) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for Json.

​	UnmarshalValue 是一个接口实现，用于为 Json 设置任何类型的值。

#### (*Json) Var

```go
func (j *Json) Var() *gvar.Var
```

Var returns the json value as *gvar.Var.

​	Var 以 *gvar.Var 的形式返回 json 值。

##### Example

``` go
```

### type Options

```go
type Options struct {
	Safe      bool        // Mark this object is for in concurrent-safe usage. This is especially for Json object creating.
	Tags      string      // Custom priority tags for decoding, eg: "json,yaml,MyTag". This is especially for struct parsing into Json object.
	Type      ContentType // Type specifies the data content type, eg: json, xml, yaml, toml, ini.
	StrNumber bool        // StrNumber causes the Decoder to unmarshal a number into an interface{} as a string instead of as a float64.
}
```

Options for Json object creating/loading.

​	用于创建/加载 Json 对象的选项。