+++
title = "gview"
date = 2024-03-21T17:57:48+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gview

Package gview implements a template engine based on text/template.

Reserved template variable names: I18nLanguage: Assign this variable to define i18n language for each page.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gview/gview_instance.go#L11)

``` go
const (
	// DefaultName is the default group name for instance usage.
	DefaultName = "default"
)
```

### Variables 

This section is empty.

### Functions 

##### func ParseContent 

``` go
func ParseContent(ctx context.Context, content string, params ...Params) (string, error)
```

ParseContent parses the template content directly using the default view object and returns the parsed content.

### Types 

#### type Config 

``` go
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

##### func DefaultConfig 

``` go
func DefaultConfig() Config
```

DefaultConfig creates and returns a configuration object with default configurations.

#### type FuncMap 

``` go
type FuncMap = map[string]interface{} // FuncMap is type for custom template functions.
```

#### type Option <-2.0.4

``` go
type Option struct {
	File    string // Template file path in absolute or relative to searching paths.
	Content string // Template content, it ignores `File` if `Content` is given.
	Orphan  bool   // If true, the `File` is considered as a single file parsing without files recursively parsing from its folder.
	Params  Params // Template parameters map.
}
```

Option for template parsing.

#### type Params 

``` go
type Params = map[string]interface{} // Params is type for template params.
```

#### type View 

``` go
type View struct {
	// contains filtered or unexported fields
}
```

View object for template engine.

##### func Instance 

``` go
func Instance(name ...string) *View
```

Instance returns an instance of View with default settings. The parameter `name` is the name for the instance.

##### func New 

``` go
func New(path ...string) *View
```

New returns a new view object. The parameter `path` specifies the template directory path to load template files.

##### (*View) AddPath 

``` go
func (view *View) AddPath(path string) error
```

AddPath adds an absolute or relative path to the search paths.

##### (*View) Assign 

``` go
func (view *View) Assign(key string, value interface{})
```

Assign binds a global template variable to current view object. Note that it's not concurrent-safe, which means it would panic if it's called in multiple goroutines in runtime.

##### (*View) Assigns 

``` go
func (view *View) Assigns(data Params)
```

Assigns binds multiple global template variables to current view object. Note that it's not concurrent-safe, which means it would panic if it's called in multiple goroutines in runtime.

##### (*View) BindFunc 

``` go
func (view *View) BindFunc(name string, function interface{})
```

BindFunc registers customized global template function named `name` with given function `function` to current view object. The `name` is the function name which can be called in template content.

##### (*View) BindFuncMap 

``` go
func (view *View) BindFuncMap(funcMap FuncMap)
```

BindFuncMap registers customized global template functions by map to current view object. The key of map is the template function name and the value of map is the address of customized function.

##### (*View) GetDefaultFile 

``` go
func (view *View) GetDefaultFile() string
```

GetDefaultFile returns default template file for parsing.

##### (*View) Parse 

``` go
func (view *View) Parse(ctx context.Context, file string, params ...Params) (result string, err error)
```

Parse parses given template file `file` with given template variables `params` and returns the parsed template content.

##### (*View) ParseContent 

``` go
func (view *View) ParseContent(ctx context.Context, content string, params ...Params) (string, error)
```

ParseContent parses given template content `content` with template variables `params` and returns the parsed content in []byte.

##### (*View) ParseDefault 

``` go
func (view *View) ParseDefault(ctx context.Context, params ...Params) (result string, err error)
```

ParseDefault parses the default template file with params.

##### (*View) ParseOption <-2.0.4

``` go
func (view *View) ParseOption(ctx context.Context, option Option) (result string, err error)
```

ParseOption implements template parsing using Option.

##### (*View) SetAutoEncode 

``` go
func (view *View) SetAutoEncode(enable bool)
```

SetAutoEncode enables/disables automatically html encoding feature. When AutoEncode feature is enables, view engine automatically encodes and provides safe html output, which is good for avoid XSS.

##### (*View) SetConfig 

``` go
func (view *View) SetConfig(config Config) error
```

SetConfig sets the configuration for view.

##### (*View) SetConfigWithMap 

``` go
func (view *View) SetConfigWithMap(m map[string]interface{}) error
```

SetConfigWithMap set configurations with map for the view.

##### (*View) SetDefaultFile 

``` go
func (view *View) SetDefaultFile(file string)
```

SetDefaultFile sets default template file for parsing.

##### (*View) SetDelimiters 

``` go
func (view *View) SetDelimiters(left, right string)
```

SetDelimiters sets customized delimiters for template parsing.

##### (*View) SetI18n 

``` go
func (view *View) SetI18n(manager *gi18n.Manager)
```

SetI18n binds i18n manager to current view engine.

##### (*View) SetPath 

``` go
func (view *View) SetPath(path string) error
```

SetPath sets the template directory path for template file search. The parameter `path` can be absolute or relative path, but absolute path is suggested.