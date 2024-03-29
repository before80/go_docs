+++
title = "gtag"
date = 2024-03-21T17:59:54+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gtag](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gtag)

Package gtag providing tag content storing for struct.

​	软件包 gtag 为 struct 提供标签内容存储。

Note that calling functions of this package is not concurrently safe, which means you cannot call them in runtime but in boot procedure.

​	请注意，调用此包的函数不是并发安全的，这意味着您不能在运行时调用它们，而是在引导过程中调用它们。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/util/gtag/gtag.go#L13)

```go
const (
	Default           = "default"      // Default value tag of struct field for receiving parameters from HTTP request.
	DefaultShort      = "d"            // Short name of Default.
	Param             = "param"        // Parameter name for converting certain parameter to specified struct field.
	ParamShort        = "p"            // Short name of Param.
	Valid             = "valid"        // Validation rule tag for struct of field.
	ValidShort        = "v"            // Short name of Valid.
	NoValidation      = "nv"           // No validation for specified struct/field.
	ORM               = "orm"          // ORM tag for ORM feature, which performs different features according scenarios.
	Arg               = "arg"          // Arg tag for struct, usually for command argument option.
	Brief             = "brief"        // Brief tag for struct, usually be considered as summary.
	Root              = "root"         // Root tag for struct, usually for nested commands management.
	Additional        = "additional"   // Additional tag for struct, usually for additional description of command.
	AdditionalShort   = "ad"           // Short name of Additional.
	Path              = `path`         // Route path for HTTP request.
	Method            = `method`       // Route method for HTTP request.
	Domain            = `domain`       // Route domain for HTTP request.
	Mime              = `mime`         // MIME type for HTTP request/response.
	Consumes          = `consumes`     // MIME type for HTTP request.
	Summary           = `summary`      // Summary for struct, usually for OpenAPI in request struct.
	SummaryShort      = `sm`           // Short name of Summary.
	SummaryShort2     = `sum`          // Short name of Summary.
	Description       = `description`  // Description for struct, usually for OpenAPI in request struct.
	DescriptionShort  = `dc`           // Short name of Description.
	DescriptionShort2 = `des`          // Short name of Description.
	Example           = `example`      // Example for struct, usually for OpenAPI in request struct.
	ExampleShort      = `eg`           // Short name of Example.
	Examples          = `examples`     // Examples for struct, usually for OpenAPI in request struct.
	ExamplesShort     = `egs`          // Short name of Examples.
	ExternalDocs      = `externalDocs` // External docs for struct, always for OpenAPI in request struct.
	ExternalDocsShort = `ed`           // Short name of ExternalDocs.
	GConv             = "gconv"        // GConv defines the converting target name for specified struct field.
	GConvShort        = "c"            // GConv defines the converting target name for specified struct field.
	Json              = "json"         // Json tag is supported by stdlib.
	Security          = "security"     // Security defines scheme for authentication. Detail to see https://swagger.io/docs/specification/authentication/
	In                = "in"           // Swagger distinguishes between the following parameter types based on the parameter location. Detail to see https://swagger.io/docs/specification/describing-parameters/
)
```

## 变量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/util/gtag/gtag.go#L54)

```go
var StructTagPriority = []string{
	GConv, Param, GConvShort, ParamShort, Json,
}
```

StructTagPriority defines the default priority tags for Map*/Struct* functions. Note that, the `gconv/param` tags are used by old version of package. It is strongly recommended using short tag `c/p` instead in the future.

​	StructTagPriority 定义 Map*/Struct* 函数的默认优先级标记。请注意，这些 `gconv/param` 标签由旧版本的包使用。强烈建议将来改用短标签 `c/p` 。

## 函数

#### func Get

```go
func Get(name string) string
```

Get retrieves and returns the stored tag content for specified name.

​	Get 检索并返回指定名称的存储标记内容。

#### func GetEnumsByType <-2.4.0

```go
func GetEnumsByType(typeName string) string
```

GetEnumsByType retrieves and returns the stored enums json by type name. The type name is like: github.com/gogf/gf/v2/encoding/gjson.ContentType

​	GetEnumsByType 按类型名称检索并返回存储的枚举 json。类型名称如下：github.com/gogf/gf/v2/encoding/gjson.ContentType

#### func GetGlobalEnums <-2.4.3

```go
func GetGlobalEnums() (string, error)
```

GetGlobalEnums retrieves and returns the global enums.

​	GetGlobalEnums 检索并返回全局枚举。

#### func Parse

```go
func Parse(content string) string
```

Parse parses and returns the content by replacing all tag name variable to its content for given `content`. Eg: gtag.Set(“demo”, “content”) Parse(`This is {demo}`) -> `This is content`.

​	Parse 解析并返回内容，方法是将 all tag name 变量替换为其给定 `content` 的内容。例如：gtag。Set（“demo”， “content”） Parse（ `This is {demo}` ） -> `This is content` .

#### func Set

```go
func Set(name, value string)
```

Set sets tag content for specified name. Note that it panics if `name` already exists.

​	设置指定名称的标签内容。请注意，如果 `name` 已经存在，它会崩溃。

##### Example

``` go
```

#### func SetGlobalEnums <-2.4.0

```go
func SetGlobalEnums(enumsJson string) error
```

SetGlobalEnums sets the global enums into package. Note that this operation is not concurrent safety.

​	SetGlobalEnums 将全局枚举设置到包中。请注意，此操作不是并发安全操作。

#### func SetOver <-2.1.0

```go
func SetOver(name, value string)
```

SetOver performs as Set, but it overwrites the old value if `name` already exists.

​	SetOver 以 Set 的形式执行，但如果 `name` 已存在，它会覆盖旧值。

#### func Sets

```go
func Sets(m map[string]string)
```

Sets sets multiple tag content by map.

​	按地图设置多个标签内容。

#### func SetsOver <-2.1.0

```go
func SetsOver(m map[string]string)
```

SetsOver performs as Sets, but it overwrites the old value if `name` already exists.

​	SetsOver 以 Sets 的形式执行，但如果 `name` 已存在旧值，它会覆盖旧值。

## 类型

This section is empty.