+++
title = "gi18n"
date = 2024-03-21T17:52:07+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/i18n/gi18n](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/i18n/gi18n)

Package gi18n implements internationalization and localization.

​	软件包 gi18n 实现国际化和本地化。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/i18n/gi18n/gi18n_instance.go#L11)

```go
const (
	// DefaultName is the default group name for instance usage.
	DefaultName = "default"
)
```

## 变量

This section is empty.

## 函数

#### func GetContent

```go
func GetContent(ctx context.Context, key string) string
```

GetContent retrieves and returns the configured content for given key and specified language. It returns an empty string if not found.

​	GetContent 检索并返回给定键和指定语言的配置内容。如果未找到，它将返回一个空字符串。

#### func LanguageFromCtx

```go
func LanguageFromCtx(ctx context.Context) string
```

LanguageFromCtx retrieves and returns language name from context. It returns an empty string if it is not set previously.

​	LanguageFromCtx 从上下文中检索并返回语言名称。如果之前未设置，则返回空字符串。

#### func SetDelimiters

```go
func SetDelimiters(left, right string)
```

SetDelimiters sets the delimiters for translator.

​	SetDelimiters 设置翻译器的分隔符。

#### func SetLanguage

```go
func SetLanguage(language string)
```

SetLanguage sets the language for translator.

​	SetLanguage 设置翻译器的语言。

#### func SetPath

```go
func SetPath(path string) error
```

SetPath sets the directory path storing i18n files.

​	SetPath 设置存储 i18n 文件的目录路径。

#### func T

```go
func T(ctx context.Context, content string) string
```

T is alias of Translate for convenience.

​	为方便起见，T 是 Translate 的别名。

#### func Tf

```go
func Tf(ctx context.Context, format string, values ...interface{}) string
```

Tf is alias of TranslateFormat for convenience.

​	为方便起见，Tf 是 TranslateFormat 的别名。

#### func Translate

```go
func Translate(ctx context.Context, content string) string
```

Translate translates `content` with configured language and returns the translated content.

​	Translate `content` 使用配置的语言进行翻译，并返回翻译后的内容。

#### func TranslateFormat

```go
func TranslateFormat(ctx context.Context, format string, values ...interface{}) string
```

TranslateFormat translates, formats and returns the `format` with configured language and given `values`.

​	TranslateFormat 翻译、格式化并返回 `format` 具有配置语言和给定 `values` 的语言。

#### func WithLanguage

```go
func WithLanguage(ctx context.Context, language string) context.Context
```

WithLanguage append language setting to the context and returns a new context.

​	WithLanguage 将语言设置追加到上下文并返回新上下文。

## 类型

### type Manager

```go
type Manager struct {
	// contains filtered or unexported fields
}
```

Manager for i18n contents, it is concurrent safe, supporting hot reload.

​	i18n内容管理器，并发安全，支持热重载。

#### func Instance

```go
func Instance(name ...string) *Manager
```

Instance returns an instance of Resource. The parameter `name` is the name for the instance.

​	Instance 返回 Resource 的实例。该参数 `name` 是实例的名称。

#### func New

```go
func New(options ...Options) *Manager
```

New creates and returns a new i18n manager. The optional parameter `option` specifies the custom options for i18n manager. It uses a default one if it’s not passed.

​	new 创建并返回一个新的 i18n 管理器。可选参数 `option` 指定 i18n 管理器的自定义选项。如果未通过，则使用默认值。

#### (*Manager) GetContent

```go
func (m *Manager) GetContent(ctx context.Context, key string) string
```

GetContent retrieves and returns the configured content for given key and specified language. It returns an empty string if not found.

​	GetContent 检索并返回给定键和指定语言的配置内容。如果未找到，它将返回一个空字符串。

#### (*Manager) SetDelimiters

```go
func (m *Manager) SetDelimiters(left, right string)
```

SetDelimiters sets the delimiters for translator.

​	SetDelimiters 设置翻译器的分隔符。

#### (*Manager) SetLanguage

```go
func (m *Manager) SetLanguage(language string)
```

SetLanguage sets the language for translator.

​	SetLanguage 设置翻译器的语言。

#### (*Manager) SetPath

```go
func (m *Manager) SetPath(path string) error
```

SetPath sets the directory path storing i18n files.

​	SetPath 设置存储 i18n 文件的目录路径。

#### (*Manager) T

```go
func (m *Manager) T(ctx context.Context, content string) string
```

T is alias of Translate for convenience.

​	为方便起见，T 是 Translate 的别名。

#### (*Manager) Tf

```go
func (m *Manager) Tf(ctx context.Context, format string, values ...interface{}) string
```

Tf is alias of TranslateFormat for convenience.

​	为方便起见，Tf 是 TranslateFormat 的别名。

#### (*Manager) Translate

```go
func (m *Manager) Translate(ctx context.Context, content string) string
```

Translate translates `content` with configured language.

​	 `content` 使用配置的语言进行翻译。

#### (*Manager) TranslateFormat

```go
func (m *Manager) TranslateFormat(ctx context.Context, format string, values ...interface{}) string
```

TranslateFormat translates, formats and returns the `format` with configured language and given `values`.

​	TranslateFormat 翻译、格式化并返回 `format` 具有配置语言和给定 `values` 的语言。

### type Options

```go
type Options struct {
	Path       string         // I18n files storage path.
	Language   string         // Default local language.
	Delimiters []string       // Delimiters for variable parsing.
	Resource   *gres.Resource // Resource for i18n files.
}
```

Options is used for i18n object configuration.

​	Options 用于 i18n 对象配置。