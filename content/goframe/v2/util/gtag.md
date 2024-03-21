+++
title = "gtag"
date = 2024-03-21T17:59:54+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gtag

Package gtag providing tag content storing for struct.

Note that calling functions of this package is not concurrently safe, which means you cannot call them in runtime but in boot procedure.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/util/gtag/gtag.go#L13)

``` go
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

### Variables 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/util/gtag/gtag.go#L54)

``` go
var StructTagPriority = []string{
	GConv, Param, GConvShort, ParamShort, Json,
}
```

StructTagPriority defines the default priority tags for Map*/Struct* functions. Note that, the `gconv/param` tags are used by old version of package. It is strongly recommended using short tag `c/p` instead in the future.

### Functions 

##### func Get 

``` go
func Get(name string) string
```

Get retrieves and returns the stored tag content for specified name.

##### func GetEnumsByType <-2.4.0

``` go
func GetEnumsByType(typeName string) string
```

GetEnumsByType retrieves and returns the stored enums json by type name. The type name is like: github.com/gogf/gf/v2/encoding/gjson.ContentType

##### func GetGlobalEnums <-2.4.3

``` go
func GetGlobalEnums() (string, error)
```

GetGlobalEnums retrieves and returns the global enums.

##### func Parse 

``` go
func Parse(content string) string
```

Parse parses and returns the content by replacing all tag name variable to its content for given `content`. Eg: gtag.Set("demo", "content") Parse(`This is {demo}`) -> `This is content`.

##### func Set 

``` go
func Set(name, value string)
```

Set sets tag content for specified name. Note that it panics if `name` already exists.

##### Example

``` go
```
##### func SetGlobalEnums <-2.4.0

``` go
func SetGlobalEnums(enumsJson string) error
```

SetGlobalEnums sets the global enums into package. Note that this operation is not concurrent safety.

##### func SetOver <-2.1.0

``` go
func SetOver(name, value string)
```

SetOver performs as Set, but it overwrites the old value if `name` already exists.

##### func Sets 

``` go
func Sets(m map[string]string)
```

Sets sets multiple tag content by map.

##### func SetsOver <-2.1.0

``` go
func SetsOver(m map[string]string)
```

SetsOver performs as Sets, but it overwrites the old value if `name` already exists.

### Types 

This section is empty.