+++
title = "gi18n"
date = 2024-03-21T17:52:07+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/i18n/gi18n

Package gi18n implements internationalization and localization.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/i18n/gi18n/gi18n_instance.go#L11)

``` go
const (
	// DefaultName is the default group name for instance usage.
	DefaultName = "default"
)
```

### Variables 

This section is empty.

### Functions 

##### func GetContent 

``` go
func GetContent(ctx context.Context, key string) string
```

GetContent retrieves and returns the configured content for given key and specified language. It returns an empty string if not found.

##### func LanguageFromCtx 

``` go
func LanguageFromCtx(ctx context.Context) string
```

LanguageFromCtx retrieves and returns language name from context. It returns an empty string if it is not set previously.

##### func SetDelimiters 

``` go
func SetDelimiters(left, right string)
```

SetDelimiters sets the delimiters for translator.

##### func SetLanguage 

``` go
func SetLanguage(language string)
```

SetLanguage sets the language for translator.

##### func SetPath 

``` go
func SetPath(path string) error
```

SetPath sets the directory path storing i18n files.

##### func T 

``` go
func T(ctx context.Context, content string) string
```

T is alias of Translate for convenience.

##### func Tf 

``` go
func Tf(ctx context.Context, format string, values ...interface{}) string
```

Tf is alias of TranslateFormat for convenience.

##### func Translate 

``` go
func Translate(ctx context.Context, content string) string
```

Translate translates `content` with configured language and returns the translated content.

##### func TranslateFormat 

``` go
func TranslateFormat(ctx context.Context, format string, values ...interface{}) string
```

TranslateFormat translates, formats and returns the `format` with configured language and given `values`.

##### func WithLanguage 

``` go
func WithLanguage(ctx context.Context, language string) context.Context
```

WithLanguage append language setting to the context and returns a new context.

### Types 

#### type Manager 

``` go
type Manager struct {
	// contains filtered or unexported fields
}
```

Manager for i18n contents, it is concurrent safe, supporting hot reload.

##### func Instance 

``` go
func Instance(name ...string) *Manager
```

Instance returns an instance of Resource. The parameter `name` is the name for the instance.

##### func New 

``` go
func New(options ...Options) *Manager
```

New creates and returns a new i18n manager. The optional parameter `option` specifies the custom options for i18n manager. It uses a default one if it's not passed.

##### (*Manager) GetContent 

``` go
func (m *Manager) GetContent(ctx context.Context, key string) string
```

GetContent retrieves and returns the configured content for given key and specified language. It returns an empty string if not found.

##### (*Manager) SetDelimiters 

``` go
func (m *Manager) SetDelimiters(left, right string)
```

SetDelimiters sets the delimiters for translator.

##### (*Manager) SetLanguage 

``` go
func (m *Manager) SetLanguage(language string)
```

SetLanguage sets the language for translator.

##### (*Manager) SetPath 

``` go
func (m *Manager) SetPath(path string) error
```

SetPath sets the directory path storing i18n files.

##### (*Manager) T 

``` go
func (m *Manager) T(ctx context.Context, content string) string
```

T is alias of Translate for convenience.

##### (*Manager) Tf 

``` go
func (m *Manager) Tf(ctx context.Context, format string, values ...interface{}) string
```

Tf is alias of TranslateFormat for convenience.

##### (*Manager) Translate 

``` go
func (m *Manager) Translate(ctx context.Context, content string) string
```

Translate translates `content` with configured language.

##### (*Manager) TranslateFormat 

``` go
func (m *Manager) TranslateFormat(ctx context.Context, format string, values ...interface{}) string
```

TranslateFormat translates, formats and returns the `format` with configured language and given `values`.

#### type Options 

``` go
type Options struct {
	Path       string         // I18n files storage path.
	Language   string         // Default local language.
	Delimiters []string       // Delimiters for variable parsing.
	Resource   *gres.Resource // Resource for i18n files.
}
```

Options is used for i18n object configuration.