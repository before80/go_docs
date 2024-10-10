+++
title = "gview"
date = 2024-03-21T17:57:48+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gview](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gview)

Package gview implements a template engine based on text/template.

​	软件包 gview 实现了一个基于文本/模板的模板引擎。

Reserved template variable names: I18nLanguage: Assign this variable to define i18n language for each page.

​	保留模板变量名称： I18nLanguage：分配此变量以定义每个页面的 i18n 语言。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gview/gview_instance.go#L11)

```go
const (
	// DefaultName is the default group name for instance usage.
	DefaultName = "default"
)
```

## 变量

This section is empty.

## 函数

#### func ParseContent

```go
func ParseContent(ctx context.Context, content string, params ...Params) (string, error)
```

ParseContent parses the template content directly using the default view object and returns the parsed content.

​	ParseContent 使用默认视图对象直接分析模板内容，并返回解析后的内容。

## 类型

### type Config

```go
type Config struct {
	Paths       []string               `json:"paths"`       // Searching array for path, NOT concurrent-safe for performance purpose.
	Data        map[string]interface{} `json:"data"`        // Global template variables including configuration.
	DefaultFile string                 `json:"defaultFile"` // Default template file for parsing.
	Delimiters  []string               `json:"delimiters"`  // Custom template delimiters.
	AutoEncode  bool                   `json:"autoEncode"`  // Automatically encodes and provides safe html output, which is good for avoiding XSS.
	I18nManager *gi18n.Manager         `json:"-"`           // I18n manager for the view.
}
```

Config is the configuration object for template engine.

​	Config 是模板引擎的配置对象。

#### func DefaultConfig

```go
func DefaultConfig() Config
```

DefaultConfig creates and returns a configuration object with default configurations.

​	DefaultConfig 创建并返回具有默认配置的配置对象。

### type FuncMap

```go
type FuncMap = map[string]interface{} // FuncMap is type for custom template functions.
```

### type Option <-2.0.4

```go
type Option struct {
	File    string // Template file path in absolute or relative to searching paths.
	Content string // Template content, it ignores `File` if `Content` is given.
	Orphan  bool   // If true, the `File` is considered as a single file parsing without files recursively parsing from its folder.
	Params  Params // Template parameters map.
}
```

Option for template parsing.

​	模板解析选项。

### type Params

```go
type Params = map[string]interface{} // Params is type for template params.
```

### type View

```go
type View struct {
	// contains filtered or unexported fields
}
```

View object for template engine.

​	模板引擎的 View 对象。

#### func Instance

```go
func Instance(name ...string) *View
```

Instance returns an instance of View with default settings. The parameter `name` is the name for the instance.

​	Instance 返回具有默认设置的 View 实例。该参数 `name` 是实例的名称。

#### func New

```go
func New(path ...string) *View
```

New returns a new view object. The parameter `path` specifies the template directory path to load template files.

​	New 返回一个新的视图对象。该参数 `path` 指定用于加载模板文件的模板目录路径。

#### (*View) AddPath

```go
func (view *View) AddPath(path string) error
```

AddPath adds an absolute or relative path to the search paths.

​	AddPath 将绝对路径或相对路径添加到搜索路径。

#### (*View) Assign

```go
func (view *View) Assign(key string, value interface{})
```

Assign binds a global template variable to current view object. Note that it’s not concurrent-safe, which means it would panic if it’s called in multiple goroutines in runtime.

​	Assign 将全局模板变量绑定到当前视图对象。请注意，它不是并发安全的，这意味着如果在运行时调用多个 goroutine，它会崩溃。

#### (*View) Assigns

```go
func (view *View) Assigns(data Params)
```

Assigns binds multiple global template variables to current view object. Note that it’s not concurrent-safe, which means it would panic if it’s called in multiple goroutines in runtime.

​	Assign 将多个全局模板变量绑定到当前视图对象。请注意，它不是并发安全的，这意味着如果在运行时调用多个 goroutine，它会崩溃。

#### (*View) BindFunc

```go
func (view *View) BindFunc(name string, function interface{})
```

BindFunc registers customized global template function named `name` with given function `function` to current view object. The `name` is the function name which can be called in template content.

​	BindFunc 将自定义的全局模板函数与给定函数 `function` 一起 `name` 注册到当前视图对象。是 `name` 可以在模板内容中调用的函数名称。

#### (*View) BindFuncMap

```go
func (view *View) BindFuncMap(funcMap FuncMap)
```

BindFuncMap registers customized global template functions by map to current view object. The key of map is the template function name and the value of map is the address of customized function.

​	BindFuncMap 通过映射到当前视图对象来注册自定义的全局模板函数。map的键是模板函数名，map的值是自定义函数的地址。

#### (*View) GetDefaultFile

```go
func (view *View) GetDefaultFile() string
```

GetDefaultFile returns default template file for parsing.

​	GetDefaultFile 返回用于分析的默认模板文件。

#### (*View) Parse

```go
func (view *View) Parse(ctx context.Context, file string, params ...Params) (result string, err error)
```

Parse parses given template file `file` with given template variables `params` and returns the parsed template content.

​	使用给定的模板变量解析给定的模板文件 `file` `params` ，并返回解析的模板内容。

#### (*View) ParseContent

```go
func (view *View) ParseContent(ctx context.Context, content string, params ...Params) (string, error)
```

ParseContent parses given template content `content` with template variables `params` and returns the parsed content in []byte.

​	ParseContent 使用模板变量解析给定的 `params` 模板内容 `content` ，并以 [] 字节的形式返回解析的内容。

#### (*View) ParseDefault

```go
func (view *View) ParseDefault(ctx context.Context, params ...Params) (result string, err error)
```

ParseDefault parses the default template file with params.

​	ParseDefault 使用参数解析默认模板文件。

#### (*View) ParseOption

```go
func (view *View) ParseOption(ctx context.Context, option Option) (result string, err error)
```

ParseOption implements template parsing using Option.

​	ParseOption 使用 Option 实现模板解析。

#### (*View) SetAutoEncode

```go
func (view *View) SetAutoEncode(enable bool)
```

SetAutoEncode enables/disables automatically html encoding feature. When AutoEncode feature is enables, view engine automatically encodes and provides safe html output, which is good for avoid XSS.

​	SetAutoEncode 自动启用/禁用 html 编码功能。启用自动编码功能后，视图引擎会自动编码并提供安全的 html 输出，这有利于避免 XSS。

#### (*View) SetConfig

```go
func (view *View) SetConfig(config Config) error
```

SetConfig sets the configuration for view.

​	SetConfig 设置视图的配置。

#### (*View) SetConfigWithMap

```go
func (view *View) SetConfigWithMap(m map[string]interface{}) error
```

SetConfigWithMap set configurations with map for the view.

​	SetConfigWithMap 使用视图的地图设置配置。

#### (*View) SetDefaultFile

```go
func (view *View) SetDefaultFile(file string)
```

SetDefaultFile sets default template file for parsing.

​	SetDefaultFile 设置用于分析的默认模板文件。

#### (*View) SetDelimiters

```go
func (view *View) SetDelimiters(left, right string)
```

SetDelimiters sets customized delimiters for template parsing.

​	SetDelimiters 为模板分析设置自定义分隔符。

#### (*View) SetI18n

```go
func (view *View) SetI18n(manager *gi18n.Manager)
```

SetI18n binds i18n manager to current view engine.

​	SetI18n 将 i18n 管理器绑定到当前视图引擎。

#### (*View) SetPath

```go
func (view *View) SetPath(path string) error
```

SetPath sets the template directory path for template file search. The parameter `path` can be absolute or relative path, but absolute path is suggested.

​	SetPath 设置模板文件搜索的模板目录路径。参数 `path` 可以是绝对路径或相对路径，但建议使用绝对路径。