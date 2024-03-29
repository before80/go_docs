+++
title = "goai"
date = 2024-03-21T17:53:15+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/goai](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/goai)

Package goai implements and provides document generating for OpenApi specification.

​	软件包 goai 实现并提供 OpenApi 规范的文档生成。

https://editor.swagger.io/

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/goai/goai.go#L40)

```go
const (
	TypeInteger    = `integer`
	TypeNumber     = `number`
	TypeBoolean    = `boolean`
	TypeArray      = `array`
	TypeString     = `string`
	TypeFile       = `file`
	TypeObject     = `object`
	FormatInt32    = `int32`
	FormatInt64    = `int64`
	FormatDouble   = `double`
	FormatByte     = `byte`
	FormatBinary   = `binary`
	FormatDate     = `date`
	FormatDateTime = `date-time`
	FormatPassword = `password`
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/goai/goai.go#L58)

```go
const (
	ParameterInHeader = `header`
	ParameterInPath   = `path`
	ParameterInQuery  = `query`
	ParameterInCookie = `cookie`
)
```

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type AddInput

```go
type AddInput struct {
	Path   string      // Path specifies the custom path if this is not configured in Meta of struct tag.
	Prefix string      // Prefix specifies the custom route path prefix, which will be added with the path tag in Meta of struct tag.
	Method string      // Method specifies the custom HTTP method if this is not configured in Meta of struct tag.
	Object interface{} // Object can be an instance of struct or a route function.
}
```

AddInput is the structured parameter for function OpenApiV3.Add.

​	AddInput 是函数 OpenApiV3.Add 的结构化参数。

### type Callback

```go
type Callback map[string]*Path
```

Callback is specified by OpenAPI/Swagger standard version 3.0.

​	回调由 OpenAPI/Swagger 标准版本 3.0 指定。

### type CallbackRef

```go
type CallbackRef struct {
	Ref   string
	Value *Callback
}
```

#### (CallbackRef) MarshalJSON

```go
func (r CallbackRef) MarshalJSON() ([]byte, error)
```

### type Callbacks

```go
type Callbacks map[string]*CallbackRef
```

### type Components

```go
type Components struct {
	Schemas         Schemas         `json:"schemas,omitempty"`
	Parameters      ParametersMap   `json:"parameters,omitempty"`
	Headers         Headers         `json:"headers,omitempty"`
	RequestBodies   RequestBodies   `json:"requestBodies,omitempty"`
	Responses       Responses       `json:"responses,omitempty"`
	SecuritySchemes SecuritySchemes `json:"securitySchemes,omitempty"`
	Examples        Examples        `json:"examples,omitempty"`
	Links           Links           `json:"links,omitempty"`
	Callbacks       Callbacks       `json:"callbacks,omitempty"`
}
```

Components is specified by OpenAPI/Swagger standard version 3.0.

​	组件由 OpenAPI/Swagger 标准版本 3.0 指定。

### type Config

```go
type Config struct {
	ReadContentTypes        []string    // ReadContentTypes specifies the default MIME types for consuming if MIME types are not configured.
	WriteContentTypes       []string    // WriteContentTypes specifies the default MIME types for producing if MIME types are not configured.
	CommonRequest           interface{} // Common request structure for all paths.
	CommonRequestDataField  string      // Common request field name to be replaced with certain business request structure. Eg: `Data`, `Request.`.
	CommonResponse          interface{} // Common response structure for all paths.
	CommonResponseDataField string      // Common response field name to be replaced with certain business response structure. Eg: `Data`, `Response.`.
	IgnorePkgPath           bool        // Ignores package name for schema name.
}
```

Config provides extra configuration feature for OpenApiV3 implements.

​	Config 为 OpenApiV3 实现提供了额外的配置功能。

### type Contact

```go
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}
```

Contact is specified by OpenAPI/Swagger standard version 3.0.

​	联系人由 OpenAPI/Swagger 标准版本 3.0 指定。

### type Content

```go
type Content map[string]MediaType
```

Content is specified by OpenAPI/Swagger 3.0 standard.

​	内容由 OpenAPI/Swagger 3.0 标准指定。

### type Discriminator

```go
type Discriminator struct {
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty"`
}
```

Discriminator is specified by OpenAPI/Swagger standard version 3.0.

​	鉴别器由 OpenAPI/Swagger 标准版本 3.0 指定。

### type Encoding

```go
type Encoding struct {
	ContentType   string  `json:"contentType,omitempty"`
	Headers       Headers `json:"headers,omitempty"`
	Style         string  `json:"style,omitempty"`
	Explode       *bool   `json:"explode,omitempty"`
	AllowReserved bool    `json:"allowReserved,omitempty"`
}
```

Encoding is specified by OpenAPI/Swagger 3.0 standard.

​	编码由 OpenAPI/Swagger 3.0 标准指定。

### type Example

```go
type Example struct {
	Summary       string      `json:"summary,omitempty"`
	Description   string      `json:"description,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	ExternalValue string      `json:"externalValue,omitempty"`
}
```

Example is specified by OpenAPI/Swagger 3.0 standard.

​	示例由 OpenAPI/Swagger 3.0 标准指定。

### type ExampleRef

```go
type ExampleRef struct {
	Ref   string
	Value *Example
}
```

#### (ExampleRef) MarshalJSON

```go
func (r ExampleRef) MarshalJSON() ([]byte, error)
```

### type Examples

```go
type Examples map[string]*ExampleRef
```

### type ExternalDocs

```go
type ExternalDocs struct {
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}
```

ExternalDocs is specified by OpenAPI/Swagger standard version 3.0.

​	ExternalDocs 由 OpenAPI/Swagger 标准版本 3.0 指定。

#### (*ExternalDocs) UnmarshalValue

```go
func (ed *ExternalDocs) UnmarshalValue(value interface{}) error
```

### type Header

```go
type Header struct {
	Parameter
}
```

Header is specified by OpenAPI/Swagger 3.0 standard. See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#headerObject

​	标头由 OpenAPI/Swagger 3.0 标准指定。请参阅 https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#headerObject

### type HeaderRef

```go
type HeaderRef struct {
	Ref   string
	Value *Header
}
```

#### (HeaderRef) MarshalJSON

```go
func (r HeaderRef) MarshalJSON() ([]byte, error)
```

### type Headers

```go
type Headers map[string]HeaderRef
```

### type Info

```go
type Info struct {
	Title          string   `json:"title"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version"`
}
```

Info is specified by OpenAPI/Swagger standard version 3.0.

​	信息由 OpenAPI/Swagger 标准版本 3.0 指定。

### type License

```go
type License struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}
```

License is specified by OpenAPI/Swagger standard version 3.0.

​	许可证由 OpenAPI/Swagger 标准版本 3.0 指定。

### type Link

```go
type Link struct {
	OperationID  string                 `json:"operationId,omitempty"`
	OperationRef string                 `json:"operationRef,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
	Server       *Server                `json:"server,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty"`
}
```

Link is specified by OpenAPI/Swagger standard version 3.0.

​	链接由 OpenAPI/Swagger 标准版本 3.0 指定。

### type LinkRef

```go
type LinkRef struct {
	Ref   string
	Value *Link
}
```

#### (LinkRef) MarshalJSON

```go
func (r LinkRef) MarshalJSON() ([]byte, error)
```

### type Links

```go
type Links map[string]LinkRef
```

### type MediaType

```go
type MediaType struct {
	Schema   *SchemaRef           `json:"schema,omitempty"`
	Example  interface{}          `json:"example,omitempty"`
	Examples Examples             `json:"examples,omitempty"`
	Encoding map[string]*Encoding `json:"encoding,omitempty"`
}
```

MediaType is specified by OpenAPI/Swagger 3.0 standard.

​	MediaType 由 OpenAPI/Swagger 3.0 标准指定。

### type OAuthFlow

```go
type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes"`
}
```

### type OAuthFlows

```go
type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty"`
}
```

### type OpenApiV3

```go
type OpenApiV3 struct {
	Config       Config                `json:"-"`
	OpenAPI      string                `json:"openapi"`
	Components   Components            `json:"components,omitempty"`
	Info         Info                  `json:"info"`
	Paths        Paths                 `json:"paths"`
	Security     *SecurityRequirements `json:"security,omitempty"`
	Servers      *Servers              `json:"servers,omitempty"`
	Tags         *Tags                 `json:"tags,omitempty"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty"`
}
```

OpenApiV3 is the structure defined from: https://swagger.io/specification/ https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md

​	OpenApiV3 是定义如下的结构： https://swagger.io/specification/ https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md

#### func New

```go
func New() *OpenApiV3
```

New creates and returns an OpenApiV3 implements object.

​	New 创建并返回 OpenApiV3 实现对象。

#### (*OpenApiV3) Add

```go
func (oai *OpenApiV3) Add(in AddInput) error
```

Add adds an instance of struct or a route function to OpenApiV3 definition implements.

​	Add 将 struct 的实例或路由函数添加到 OpenApiV3 定义实现中。

#### (OpenApiV3) String

```go
func (oai OpenApiV3) String() string
```

### type Operation

```go
type Operation struct {
	Tags         []string              `json:"tags,omitempty"`
	Summary      string                `json:"summary,omitempty"`
	Description  string                `json:"description,omitempty"`
	OperationID  string                `json:"operationId,omitempty"`
	Parameters   Parameters            `json:"parameters,omitempty"`
	RequestBody  *RequestBodyRef       `json:"requestBody,omitempty"`
	Responses    Responses             `json:"responses"`
	Deprecated   bool                  `json:"deprecated,omitempty"`
	Callbacks    *Callbacks            `json:"callbacks,omitempty"`
	Security     *SecurityRequirements `json:"security,omitempty"`
	Servers      *Servers              `json:"servers,omitempty"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty"`
	XExtensions  XExtensions           `json:"-"`
}
```

Operation represents “operation” specified by OpenAPI/Swagger 3.0 standard.

​	Operation 表示 OpenAPI/Swagger 3.0 标准指定的“操作”。

#### (Operation) MarshalJSON

```go
func (o Operation) MarshalJSON() ([]byte, error)
```

### type Parameter

```go
type Parameter struct {
	Name            string      `json:"name,omitempty"`
	In              string      `json:"in,omitempty"`
	Description     string      `json:"description,omitempty"`
	Style           string      `json:"style,omitempty"`
	Explode         *bool       `json:"explode,omitempty"`
	AllowEmptyValue bool        `json:"allowEmptyValue,omitempty"`
	AllowReserved   bool        `json:"allowReserved,omitempty"`
	Deprecated      bool        `json:"deprecated,omitempty"`
	Required        bool        `json:"required,omitempty"`
	Schema          *SchemaRef  `json:"schema,omitempty"`
	Example         interface{} `json:"example,omitempty"`
	Examples        *Examples   `json:"examples,omitempty"`
	Content         *Content    `json:"content,omitempty"`
	XExtensions     XExtensions `json:"-"`
}
```

Parameter is specified by OpenAPI/Swagger 3.0 standard. See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterObject

​	参数由 OpenAPI/Swagger 3.0 标准指定。请参阅 https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterObject

#### (Parameter) MarshalJSON

```go
func (p Parameter) MarshalJSON() ([]byte, error)
```

### type ParameterRef

```go
type ParameterRef struct {
	Ref   string
	Value *Parameter
}
```

#### (ParameterRef) MarshalJSON

```go
func (r ParameterRef) MarshalJSON() ([]byte, error)
```

### type Parameters

```go
type Parameters []ParameterRef
```

Parameters is specified by OpenAPI/Swagger 3.0 standard.

​	参数由 OpenAPI/Swagger 3.0 标准指定。

### type ParametersMap

```go
type ParametersMap map[string]*ParameterRef
```

### type Path

```go
type Path struct {
	Ref         string      `json:"$ref,omitempty"`
	Summary     string      `json:"summary,omitempty"`
	Description string      `json:"description,omitempty"`
	Connect     *Operation  `json:"connect,omitempty"`
	Delete      *Operation  `json:"delete,omitempty"`
	Get         *Operation  `json:"get,omitempty"`
	Head        *Operation  `json:"head,omitempty"`
	Options     *Operation  `json:"options,omitempty"`
	Patch       *Operation  `json:"patch,omitempty"`
	Post        *Operation  `json:"post,omitempty"`
	Put         *Operation  `json:"put,omitempty"`
	Trace       *Operation  `json:"trace,omitempty"`
	Servers     Servers     `json:"servers,omitempty"`
	Parameters  Parameters  `json:"parameters,omitempty"`
	XExtensions XExtensions `json:"-"`
}
```

Path is specified by OpenAPI/Swagger standard version 3.0.

​	路径由 OpenAPI/Swagger 标准版本 3.0 指定。

#### (Path) MarshalJSON

```go
func (p Path) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

### type Paths

```go
type Paths map[string]Path
```

Paths are specified by OpenAPI/Swagger standard version 3.0.

​	路径由 OpenAPI/Swagger 标准版本 3.0 指定。

### type RequestBodies

```go
type RequestBodies map[string]*RequestBodyRef
```

### type RequestBody

```go
type RequestBody struct {
	Description string  `json:"description,omitempty"`
	Required    bool    `json:"required,omitempty"`
	Content     Content `json:"content,omitempty"`
}
```

RequestBody is specified by OpenAPI/Swagger 3.0 standard.

​	RequestBody 由 OpenAPI/Swagger 3.0 标准指定。

### type RequestBodyRef

```go
type RequestBodyRef struct {
	Ref   string
	Value *RequestBody
}
```

#### (RequestBodyRef) MarshalJSON

```go
func (r RequestBodyRef) MarshalJSON() ([]byte, error)
```

### type Response

```go
type Response struct {
	Description string      `json:"description"`
	Headers     Headers     `json:"headers,omitempty"`
	Content     Content     `json:"content,omitempty"`
	Links       Links       `json:"links,omitempty"`
	XExtensions XExtensions `json:"-"`
}
```

Response is specified by OpenAPI/Swagger 3.0 standard.

​	响应由 OpenAPI/Swagger 3.0 标准指定。

#### (Response) MarshalJSON

```go
func (r Response) MarshalJSON() ([]byte, error)
```

### type ResponseRef

```go
type ResponseRef struct {
	Ref   string
	Value *Response
}
```

#### (ResponseRef) MarshalJSON

```go
func (r ResponseRef) MarshalJSON() ([]byte, error)
```

### type Responses

```go
type Responses map[string]ResponseRef
```

Responses is specified by OpenAPI/Swagger 3.0 standard.

​	响应由 OpenAPI/Swagger 3.0 标准指定。

### type Schema

```go
type Schema struct {
	OneOf                SchemaRefs     `json:"oneOf,omitempty"`
	AnyOf                SchemaRefs     `json:"anyOf,omitempty"`
	AllOf                SchemaRefs     `json:"allOf,omitempty"`
	Not                  *SchemaRef     `json:"not,omitempty"`
	Type                 string         `json:"type,omitempty"`
	Title                string         `json:"title,omitempty"`
	Format               string         `json:"format,omitempty"`
	Description          string         `json:"description,omitempty"`
	Enum                 []interface{}  `json:"enum,omitempty"`
	Default              interface{}    `json:"default,omitempty"`
	Example              interface{}    `json:"example,omitempty"`
	ExternalDocs         *ExternalDocs  `json:"externalDocs,omitempty"`
	UniqueItems          bool           `json:"uniqueItems,omitempty"`
	ExclusiveMin         bool           `json:"exclusiveMinimum,omitempty"`
	ExclusiveMax         bool           `json:"exclusiveMaximum,omitempty"`
	Nullable             bool           `json:"nullable,omitempty"`
	ReadOnly             bool           `json:"readOnly,omitempty"`
	WriteOnly            bool           `json:"writeOnly,omitempty"`
	AllowEmptyValue      bool           `json:"allowEmptyValue,omitempty"`
	XML                  interface{}    `json:"xml,omitempty"`
	Deprecated           bool           `json:"deprecated,omitempty"`
	Min                  *float64       `json:"minimum,omitempty"`
	Max                  *float64       `json:"maximum,omitempty"`
	MultipleOf           *float64       `json:"multipleOf,omitempty"`
	MinLength            uint64         `json:"minLength,omitempty"`
	MaxLength            *uint64        `json:"maxLength,omitempty"`
	Pattern              string         `json:"pattern,omitempty"`
	MinItems             uint64         `json:"minItems,omitempty"`
	MaxItems             *uint64        `json:"maxItems,omitempty"`
	Items                *SchemaRef     `json:"items,omitempty"`
	Required             []string       `json:"required,omitempty"`
	Properties           Schemas        `json:"properties,omitempty"`
	MinProps             uint64         `json:"minProperties,omitempty"`
	MaxProps             *uint64        `json:"maxProperties,omitempty"`
	AdditionalProperties *SchemaRef     `json:"additionalProperties,omitempty"`
	Discriminator        *Discriminator `json:"discriminator,omitempty"`
	XExtensions          XExtensions    `json:"-"`
	ValidationRules      string         `json:"-"`
}
```

Schema is specified by OpenAPI/Swagger 3.0 standard.

​	架构由 OpenAPI/Swagger 3.0 标准指定。

#### (*Schema) Clone

```go
func (s *Schema) Clone() *Schema
```

Clone only clones necessary attributes. TODO clone all attributes, or improve package deepcopy.

​	仅克隆必要的属性。TODO 克隆所有属性，或改进包 deepcopy。

#### (Schema) MarshalJSON

```go
func (s Schema) MarshalJSON() ([]byte, error)
```

### type SchemaRef

```go
type SchemaRef struct {
	Ref   string
	Value *Schema
}
```

#### (SchemaRef) MarshalJSON

```go
func (r SchemaRef) MarshalJSON() ([]byte, error)
```

### type SchemaRefs

```go
type SchemaRefs []SchemaRef
```

### type Schemas

```go
type Schemas struct {
	// contains filtered or unexported fields
}
```

#### (*Schemas) Clone

```go
func (s *Schemas) Clone() Schemas
```

#### (*Schemas) Get

```go
func (s *Schemas) Get(name string) *SchemaRef
```

#### (*Schemas) Iterator

```go
func (s *Schemas) Iterator(f func(key string, ref SchemaRef) bool)
```

#### (*Schemas) Map

```go
func (s *Schemas) Map() map[string]SchemaRef
```

#### (Schemas) MarshalJSON

```go
func (s Schemas) MarshalJSON() ([]byte, error)
```

#### (*Schemas) Removes

```go
func (s *Schemas) Removes(names []interface{})
```

#### (*Schemas) Set

```go
func (s *Schemas) Set(name string, ref SchemaRef)
```

### type SecurityRequirement

```go
type SecurityRequirement map[string][]string
```

### type SecurityRequirements

```go
type SecurityRequirements []SecurityRequirement
```

### type SecurityScheme

```go
type SecurityScheme struct {
	Type             string      `json:"type,omitempty"`
	Description      string      `json:"description,omitempty"`
	Name             string      `json:"name,omitempty"`
	In               string      `json:"in,omitempty"`
	Scheme           string      `json:"scheme,omitempty"`
	BearerFormat     string      `json:"bearerFormat,omitempty"`
	Flows            *OAuthFlows `json:"flows,omitempty"`
	OpenIdConnectUrl string      `json:"openIdConnectUrl,omitempty"`
}
```

### type SecuritySchemeRef

```go
type SecuritySchemeRef struct {
	Ref   string
	Value *SecurityScheme
}
```

#### (SecuritySchemeRef) MarshalJSON

```go
func (r SecuritySchemeRef) MarshalJSON() ([]byte, error)
```

### type SecuritySchemes

```go
type SecuritySchemes map[string]SecuritySchemeRef
```

### type Server

```go
type Server struct {
	URL         string                     `json:"url"`
	Description string                     `json:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty"`
}
```

Server is specified by OpenAPI/Swagger standard version 3.0.

​	服务器由 OpenAPI/Swagger 标准版本 3.0 指定。

### type ServerVariable

```go
type ServerVariable struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
}
```

ServerVariable is specified by OpenAPI/Swagger standard version 3.0.

​	ServerVariable 由 OpenAPI/Swagger 标准版本 3.0 指定。

### type Servers

```go
type Servers []Server
```

Servers is specified by OpenAPI/Swagger standard version 3.0.

​	服务器由 OpenAPI/Swagger 标准版本 3.0 指定。

### type Tag

```go
type Tag struct {
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
}
```

Tag is specified by OpenAPI/Swagger 3.0 standard.

​	标签由 OpenAPI/Swagger 3.0 标准指定。

### type Tags

```go
type Tags []Tag
```

Tags is specified by OpenAPI/Swagger 3.0 standard.

​	标签由 OpenAPI/Swagger 3.0 标准指定。

### type XExtensions

```go
type XExtensions map[string]string
```

XExtensions stores the `x-` custom extensions.

​	XExtensions 存储 `x-` 自定义扩展。