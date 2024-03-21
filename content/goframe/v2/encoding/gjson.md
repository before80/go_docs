+++
title = "gjson"
date = 2024-03-21T17:49:47+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gjson

Package gjson provides convenient API for JSON/XML/INI/YAML/TOML data handling.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Decode 

``` go
func Decode(data interface{}, options ...Options) (interface{}, error)
```

Decode decodes json format `data` to golang variable. The parameter `data` can be either bytes or string type.

##### Example

``` go
```
##### func DecodeTo 

``` go
func DecodeTo(data interface{}, v interface{}, options ...Options) (err error)
```

DecodeTo decodes json format `data` to specified golang variable `v`. The parameter `data` can be either bytes or string type. The parameter `v` should be a pointer type.

##### Example

``` go
```
##### func Encode 

``` go
func Encode(value interface{}) ([]byte, error)
```

Encode encodes any golang variable `value` to JSON bytes.

##### Example

``` go
```
##### func EncodeString 

``` go
func EncodeString(value interface{}) (string, error)
```

EncodeString encodes any golang variable `value` to JSON string.

##### Example

``` go
```
##### func IsValidDataType 

``` go
func IsValidDataType(dataType ContentType) bool
```

IsValidDataType checks and returns whether given `dataType` a valid data type for loading.

##### Example

``` go
```
##### func Marshal 

``` go
func Marshal(v interface{}) (marshaledBytes []byte, err error)
```

Marshal is alias of Encode in order to fit the habit of json.Marshal/Unmarshal functions.

##### Example

``` go
```
##### func MarshalIndent 

``` go
func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error)
```

MarshalIndent is alias of json.MarshalIndent in order to fit the habit of json.MarshalIndent function.

##### Example

``` go
```
##### func MustEncode 

``` go
func MustEncode(value interface{}) []byte
```

MustEncode performs as Encode, but it panics if any error occurs.

##### Example

``` go
```
##### func MustEncodeString 

``` go
func MustEncodeString(value interface{}) string
```

MustEncodeString encodes any golang variable `value` to JSON string. It panics if any error occurs.

##### Example

``` go
```
##### func Unmarshal 

``` go
func Unmarshal(data []byte, v interface{}) (err error)
```

Unmarshal is alias of DecodeTo in order to fit the habit of json.Marshal/Unmarshal functions.

##### Example

``` go
```
##### func Valid 

``` go
func Valid(data interface{}) bool
```

Valid checks whether `data` is a valid JSON data type. The parameter `data` specifies the json format data, which can be either bytes or string type.

##### Example

``` go
```
### Types 

#### type ContentType <-2.2.0

``` go
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

#### type Json 

``` go
type Json struct {
	// contains filtered or unexported fields
}
```

Json is the customized JSON struct.

##### func DecodeToJson 

``` go
func DecodeToJson(data interface{}, options ...Options) (*Json, error)
```

DecodeToJson codes json format `data` to a Json object. The parameter `data` can be either bytes or string type.

##### Example

``` go
```
##### func Load 

``` go
func Load(path string, safe ...bool) (*Json, error)
```

Load loads content from specified file `path`, and creates a Json object from its content.

##### Example

``` go
```
##### func LoadContent 

``` go
func LoadContent(data interface{}, safe ...bool) (*Json, error)
```

LoadContent creates a Json object from given content, it checks the data type of `content` automatically, supporting data content type as follows: JSON, XML, INI, YAML and TOML.

##### Example

``` go
```
##### func LoadContentType 

``` go
func LoadContentType(dataType ContentType, data interface{}, safe ...bool) (*Json, error)
```

LoadContentType creates a Json object from given type and content, supporting data content type as follows: JSON, XML, INI, YAML and TOML.

##### Example

``` go
```
##### func LoadIni 

``` go
func LoadIni(data interface{}, safe ...bool) (*Json, error)
```

LoadIni creates a Json object from given INI format content.

##### Example

``` go
```
##### func LoadJson 

``` go
func LoadJson(data interface{}, safe ...bool) (*Json, error)
```

LoadJson creates a Json object from given JSON format content.

##### Example

``` go
```
##### func LoadProperties <-2.1.0

``` go
func LoadProperties(data interface{}, safe ...bool) (*Json, error)
```

LoadProperties creates a Json object from given TOML format content.

##### func LoadToml 

``` go
func LoadToml(data interface{}, safe ...bool) (*Json, error)
```

LoadToml creates a Json object from given TOML format content.

##### Example

``` go
```
##### func LoadWithOptions <-2.1.0

``` go
func LoadWithOptions(data interface{}, options Options) (*Json, error)
```

LoadWithOptions creates a Json object from given JSON format content and options.

##### func LoadXml 

``` go
func LoadXml(data interface{}, safe ...bool) (*Json, error)
```

LoadXml creates a Json object from given XML format content.

##### Example

``` go
```
##### func LoadYaml 

``` go
func LoadYaml(data interface{}, safe ...bool) (*Json, error)
```

LoadYaml creates a Json object from given YAML format content.

##### Example

``` go
```
##### func New 

``` go
func New(data interface{}, safe ...bool) *Json
```

New creates a Json object with any variable type of `data`, but `data` should be a map or slice for data access reason, or it will make no sense.

The parameter `safe` specifies whether using this Json object in concurrent-safe context, which is false in default.

##### Example

``` go
```
##### func NewWithOptions 

``` go
func NewWithOptions(data interface{}, options Options) *Json
```

NewWithOptions creates a Json object with any variable type of `data`, but `data` should be a map or slice for data access reason, or it will make no sense.

##### Example

``` go
```
##### func NewWithTag 

``` go
func NewWithTag(data interface{}, tags string, safe ...bool) *Json
```

NewWithTag creates a Json object with any variable type of `data`, but `data` should be a map or slice for data access reason, or it will make no sense.

The parameter `tags` specifies priority tags for struct conversion to map, multiple tags joined with char ','.

The parameter `safe` specifies whether using this Json object in concurrent-safe context, which is false in default.

##### Example

``` go
```
##### (*Json) Append 

``` go
func (j *Json) Append(pattern string, value interface{}) error
```

Append appends value to the value by specified `pattern`. The target value by `pattern` should be type of slice.

##### Example

``` go
```
##### (*Json) Array 

``` go
func (j *Json) Array() []interface{}
```

Array converts current Json object to []interface{}. It returns nil if fails.

##### Example

``` go
```
##### (*Json) Contains 

``` go
func (j *Json) Contains(pattern string) bool
```

Contains checks whether the value by specified `pattern` exist.

##### Example

``` go
```
##### (*Json) Dump 

``` go
func (j *Json) Dump()
```

Dump prints current Json object with more manually readable.

##### Example

``` go
```
##### (*Json) Get 

``` go
func (j *Json) Get(pattern string, def ...interface{}) *gvar.Var
```

Get retrieves and returns value by specified `pattern`. It returns all values of current Json object if `pattern` is given ".". It returns nil if no value found by `pattern`.

We can also access slice item by its index number in `pattern` like: "list.10", "array.0.name", "array.0.1.id".

It returns a default value specified by `def` if value for `pattern` is not found.

##### Example

``` go
```
##### (*Json) GetJson 

``` go
func (j *Json) GetJson(pattern string, def ...interface{}) *Json
```

GetJson gets the value by specified `pattern`, and converts it to an un-concurrent-safe Json object.

##### Example

``` go
```
##### (*Json) GetJsonMap 

``` go
func (j *Json) GetJsonMap(pattern string, def ...interface{}) map[string]*Json
```

GetJsonMap gets the value by specified `pattern`, and converts it to a map of un-concurrent-safe Json object.

##### Example

``` go
```
##### (*Json) GetJsons 

``` go
func (j *Json) GetJsons(pattern string, def ...interface{}) []*Json
```

GetJsons gets the value by specified `pattern`, and converts it to a slice of un-concurrent-safe Json object.

##### Example

``` go
```
##### (*Json) Interface 

``` go
func (j *Json) Interface() interface{}
```

Interface returns the json value.

##### Example

``` go
```
##### (*Json) Interfaces 

``` go
func (j *Json) Interfaces() []interface{}
```

Interfaces implements interface function Interfaces().

##### Example

``` go
```
##### (*Json) IsNil 

``` go
func (j *Json) IsNil() bool
```

IsNil checks whether the value pointed by `j` is nil.

##### Example

``` go
```
##### (*Json) Len 

``` go
func (j *Json) Len(pattern string) int
```

Len returns the length/size of the value by specified `pattern`. The target value by `pattern` should be type of slice or map. It returns -1 if the target value is not found, or its type is invalid.

##### Example

``` go
```
##### (*Json) Map 

``` go
func (j *Json) Map() map[string]interface{}
```

Map converts current Json object to map[string]interface{}. It returns nil if fails.

##### Example

``` go
```
##### (*Json) MapStrAny 

``` go
func (j *Json) MapStrAny() map[string]interface{}
```

MapStrAny implements interface function MapStrAny().

##### Example

``` go
```
##### (Json) MarshalJSON 

``` go
func (j Json) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### Example

``` go
```
##### (*Json) MustAppend 

``` go
func (j *Json) MustAppend(pattern string, value interface{})
```

MustAppend performs as Append, but it panics if any error occurs.

##### Example

``` go
```
##### (*Json) MustRemove 

``` go
func (j *Json) MustRemove(pattern string)
```

MustRemove performs as Remove, but it panics if any error occurs.

##### Example

``` go
```
##### (*Json) MustSet 

``` go
func (j *Json) MustSet(pattern string, value interface{})
```

MustSet performs as Set, but it panics if any error occurs.

##### Example

``` go
```
##### (*Json) MustToIni 

``` go
func (j *Json) MustToIni() []byte
```

##### Example

``` go
```
##### (*Json) MustToIniString 

``` go
func (j *Json) MustToIniString() string
```

MustToIniString .

##### Example

``` go
```
##### (*Json) MustToJson 

``` go
func (j *Json) MustToJson() []byte
```

##### Example

``` go
```
##### (*Json) MustToJsonIndent 

``` go
func (j *Json) MustToJsonIndent() []byte
```

##### Example

``` go
```
##### (*Json) MustToJsonIndentString 

``` go
func (j *Json) MustToJsonIndentString() string
```

##### Example

``` go
```
##### (*Json) MustToJsonString 

``` go
func (j *Json) MustToJsonString() string
```

##### Example

``` go
```
##### (*Json) MustToProperties <-2.1.0

``` go
func (j *Json) MustToProperties() []byte
```

##### Example

``` go
```
##### (*Json) MustToPropertiesString <-2.1.0

``` go
func (j *Json) MustToPropertiesString() string
```

MustTopropertiesString

##### Example

``` go
```
##### (*Json) MustToToml 

``` go
func (j *Json) MustToToml() []byte
```

##### Example

``` go
```
##### (*Json) MustToTomlString 

``` go
func (j *Json) MustToTomlString() string
```

##### Example

``` go
```
##### (*Json) MustToXml 

``` go
func (j *Json) MustToXml(rootTag ...string) []byte
```

##### Example

``` go
```
##### (*Json) MustToXmlIndent 

``` go
func (j *Json) MustToXmlIndent(rootTag ...string) []byte
```

##### Example

``` go
```
##### (*Json) MustToXmlIndentString 

``` go
func (j *Json) MustToXmlIndentString(rootTag ...string) string
```

##### Example

``` go
```
##### (*Json) MustToXmlString 

``` go
func (j *Json) MustToXmlString(rootTag ...string) string
```

##### Example

``` go
```
##### (*Json) MustToYaml 

``` go
func (j *Json) MustToYaml() []byte
```

##### Example

``` go
```
##### (*Json) MustToYamlString 

``` go
func (j *Json) MustToYamlString() string
```

##### Example

``` go
```
##### (*Json) Remove 

``` go
func (j *Json) Remove(pattern string) error
```

Remove deletes value with specified `pattern`. It supports hierarchical data access by char separator, which is '.' in default.

##### Example

``` go
```
##### (*Json) Scan 

``` go
func (j *Json) Scan(pointer interface{}, mapping ...map[string]string) error
```

Scan automatically calls Struct or Structs function according to the type of parameter `pointer` to implement the converting.

##### Example

``` go
```
##### (*Json) Set 

``` go
func (j *Json) Set(pattern string, value interface{}) error
```

Set sets value with specified `pattern`. It supports hierarchical data access by char separator, which is '.' in default.

##### Example

``` go
```
##### (*Json) SetSplitChar 

``` go
func (j *Json) SetSplitChar(char byte)
```

SetSplitChar sets the separator char for hierarchical data access.

##### Example

``` go
```
##### (*Json) SetViolenceCheck 

``` go
func (j *Json) SetViolenceCheck(enabled bool)
```

SetViolenceCheck enables/disables violence check for hierarchical data access.

##### Example

``` go
```
##### (*Json) String <-2.2.6

``` go
func (j *Json) String() string
```

String returns current Json object as string.

##### (*Json) ToIni 

``` go
func (j *Json) ToIni() ([]byte, error)
```

ToIni json to ini

##### Example

``` go
```
##### (*Json) ToIniString 

``` go
func (j *Json) ToIniString() (string, error)
```

ToIniString ini to string

##### Example

``` go
```
##### (*Json) ToJson 

``` go
func (j *Json) ToJson() ([]byte, error)
```

##### Example

``` go
```
##### (*Json) ToJsonIndent 

``` go
func (j *Json) ToJsonIndent() ([]byte, error)
```

##### Example

``` go
```
##### (*Json) ToJsonIndentString 

``` go
func (j *Json) ToJsonIndentString() (string, error)
```

##### Example

``` go
```
##### (*Json) ToJsonString 

``` go
func (j *Json) ToJsonString() (string, error)
```

##### Example

``` go
```
##### (*Json) ToProperties <-2.1.0

``` go
func (j *Json) ToProperties() ([]byte, error)
```

======================================================================== properties ======================================================================== Toproperties json to properties

##### Example

``` go
```
##### (*Json) ToPropertiesString <-2.1.0

``` go
func (j *Json) ToPropertiesString() (string, error)
```

TopropertiesString properties to string

##### Example

``` go
```
##### (*Json) ToToml 

``` go
func (j *Json) ToToml() ([]byte, error)
```

##### Example

``` go
```
##### (*Json) ToTomlString 

``` go
func (j *Json) ToTomlString() (string, error)
```

##### Example

``` go
```
##### (*Json) ToXml 

``` go
func (j *Json) ToXml(rootTag ...string) ([]byte, error)
```

##### Example

``` go
```
##### (*Json) ToXmlIndent 

``` go
func (j *Json) ToXmlIndent(rootTag ...string) ([]byte, error)
```

##### Example

``` go
```
##### (*Json) ToXmlIndentString 

``` go
func (j *Json) ToXmlIndentString(rootTag ...string) (string, error)
```

##### Example

``` go
```
##### (*Json) ToXmlString 

``` go
func (j *Json) ToXmlString(rootTag ...string) (string, error)
```

##### Example

``` go
```
##### (*Json) ToYaml 

``` go
func (j *Json) ToYaml() ([]byte, error)
```

##### Example

``` go
```
##### (*Json) ToYamlIndent 

``` go
func (j *Json) ToYamlIndent(indent string) ([]byte, error)
```

##### Example

``` go
```
##### (*Json) ToYamlString 

``` go
func (j *Json) ToYamlString() (string, error)
```

##### Example

``` go
```
##### (*Json) UnmarshalJSON 

``` go
func (j *Json) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*Json) UnmarshalValue 

``` go
func (j *Json) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for Json.

##### (*Json) Var 

``` go
func (j *Json) Var() *gvar.Var
```

Var returns the json value as *gvar.Var.

##### Example

``` go
```
#### type Options 

``` go
type Options struct {
	Safe      bool        // Mark this object is for in concurrent-safe usage. This is especially for Json object creating.
	Tags      string      // Custom priority tags for decoding, eg: "json,yaml,MyTag". This is especially for struct parsing into Json object.
	Type      ContentType // Type specifies the data content type, eg: json, xml, yaml, toml, ini.
	StrNumber bool        // StrNumber causes the Decoder to unmarshal a number into an interface{} as a string instead of as a float64.
}
```

Options for Json object creating/loading.