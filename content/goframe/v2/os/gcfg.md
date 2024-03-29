+++
title = "gcfg"
date = 2024-03-21T17:54:46+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcfg](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcfg)

Package gcfg provides reading, caching and managing for configuration.

​	软件包 gcfg 提供读取、缓存和管理配置。

#### Examples 例子

- [Config.GetWithCmd](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcfg#example-Config.GetWithCmd)
- [Config.GetWithEnv](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcfg#example-Config.GetWithEnv)

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gcfg/gcfg.go#L27)

```go
const (
	DefaultInstanceName   = "config" // DefaultName is the default instance name for instance usage.
	DefaultConfigFileName = "config" // DefaultConfigFile is the default configuration file name.
)
```

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Adapter

```go
type Adapter interface {
	// Available checks and returns the backend configuration service is available.
	// The optional parameter `resource` specifies certain configuration resource.
	//
	// Note that this function does not return error as it just does simply check for
	// backend configuration service.
	Available(ctx context.Context, resource ...string) (ok bool)

	// Get retrieves and returns value by specified `pattern` in current resource.
	// Pattern like:
	// "x.y.z" for map item.
	// "x.0.y" for slice item.
	Get(ctx context.Context, pattern string) (value interface{}, err error)

	// Data retrieves and returns all configuration data in current resource as map.
	// Note that this function may lead lots of memory usage if configuration data is too large,
	// you can implement this function if necessary.
	Data(ctx context.Context) (data map[string]interface{}, err error)
}
```

Adapter is the interface for configuration retrieving.

​	适配器是用于配置检索的接口。

### type AdapterContent <-2.5.3

```go
type AdapterContent struct {
	// contains filtered or unexported fields
}
```

AdapterContent implements interface Adapter using content. The configuration content supports the coding types as package `gjson`.

​	AdapterContent 使用 content 实现接口 Adapter。配置内容支持将编码类型作为 package `gjson` 。

#### func NewAdapterContent <-2.5.3

```go
func NewAdapterContent(content ...string) (*AdapterContent, error)
```

NewAdapterContent returns a new configuration management object using custom content. The parameter `content` specifies the default configuration content for reading.

​	NewAdapterContent 使用自定义内容返回新的配置管理对象。该参数 `content` 指定用于读取的默认配置内容。

#### (*AdapterContent) Available

```go
func (a *AdapterContent) Available(ctx context.Context, resource ...string) (ok bool)
```

Available checks and returns the backend configuration service is available. The optional parameter `resource` specifies certain configuration resource.

​	available 检查并返回后端配置服务是否可用。可选参数 `resource` 指定特定配置资源。

Note that this function does not return error as it just does simply check for backend configuration service.

​	请注意，此函数不会返回错误，因为它只是检查后端配置服务。

#### (*AdapterContent) Data

```go
func (a *AdapterContent) Data(ctx context.Context) (data map[string]interface{}, err error)
```

Data retrieves and returns all configuration data in current resource as map. Note that this function may lead lots of memory usage if configuration data is too large, you can implement this function if necessary.

​	数据检索当前资源中的所有配置数据，并将其作为映射返回。请注意，如果配置数据过大，此函数可能会导致大量内存使用，如有必要，可以实现此函数。

#### (*AdapterContent) Get

```go
func (a *AdapterContent) Get(ctx context.Context, pattern string) (value interface{}, err error)
```

Get retrieves and returns value by specified `pattern` in current resource. Pattern like: “x.y.z” for map item. “x.0.y” for slice item.

​	按当前资源中指定的 `pattern` 方式获取检索和返回值。类似图案：“x.y.z”表示地图项。“x.0.y”表示切片项。

#### (*AdapterContent) SetContent

```go
func (a *AdapterContent) SetContent(content string) error
```

SetContent sets customized configuration content for specified `file`. The `file` is unnecessary param, default is DefaultConfigFile.

​	SetContent 为指定的 `file` . `file` 这是不必要的参数，默认值为 DefaultConfigFile。

### type AdapterFile

```go
type AdapterFile struct {
	// contains filtered or unexported fields
}
```

AdapterFile implements interface Adapter using file.

​	AdapterFile 使用 file 实现接口 Adapter。

#### func NewAdapterFile

```go
func NewAdapterFile(file ...string) (*AdapterFile, error)
```

NewAdapterFile returns a new configuration management object. The parameter `file` specifies the default configuration file name for reading.

​	NewAdapterFile 返回新的配置管理对象。该参数 `file` 指定用于读取的默认配置文件名。

#### (*AdapterFile) AddPath

```go
func (a *AdapterFile) AddPath(directoryPaths ...string) (err error)
```

AddPath adds an absolute or relative `directory` path to the search paths.

​	AddPath 将绝对路径或相对 `directory` 路径添加到搜索路径。

Note that this parameter is paths to a directories not files.

​	请注意，此参数是目录的路径，而不是文件。

#### (*AdapterFile) Available

```go
func (a *AdapterFile) Available(ctx context.Context, fileName ...string) bool
```

Available checks and returns whether configuration of given `file` is available.

​	available 检查并返回给定 `file` 的配置是否可用。

#### (*AdapterFile) Clear

```go
func (a *AdapterFile) Clear()
```

Clear removes all parsed configuration files content cache, which will force reload configuration content from file.

​	Clear 将删除所有已分析的配置文件内容缓存，这将强制从文件中重新加载配置内容。

#### (*AdapterFile) ClearContent

```go
func (a *AdapterFile) ClearContent()
```

ClearContent removes all global configuration contents.

​	ClearContent 删除所有全局配置内容。

#### (*AdapterFile) Data

```go
func (a *AdapterFile) Data(ctx context.Context) (data map[string]interface{}, err error)
```

Data retrieves and returns all configuration data as map type.

​	数据检索并返回所有配置数据作为映射类型。

#### (*AdapterFile) Dump

```go
func (a *AdapterFile) Dump()
```

Dump prints current Json object with more manually readable.

​	转储打印当前 Json 对象，具有更多的手动可读性。

#### (*AdapterFile) Get

```go
func (a *AdapterFile) Get(ctx context.Context, pattern string) (value interface{}, err error)
```

Get retrieves and returns value by specified `pattern`. It returns all values of current Json object if `pattern` is given empty or string “.”. It returns nil if no value found by `pattern`.

​	按指定 `pattern` 获取检索和返回值。如果 `pattern` 给定空或字符串“.”，则返回当前 Json 对象的所有值。如果 未找到 的 `pattern` 值，则返回 nil。

We can also access slice item by its index number in `pattern` like: “list.10”, “array.0.name”, “array.0.1.id”.

​	我们还可以通过索引号访问切片项目， `pattern` 例如：“list.10”、“array.0.name”、“array.0.1.id”。

It returns a default value specified by `def` if value for `pattern` is not found.

​	它返回由 `def` if value for 指定的默认值 if value `pattern` is not found。

#### (*AdapterFile) GetContent

```go
func (a *AdapterFile) GetContent(file ...string) string
```

GetContent returns customized configuration content for specified `file`. The `file` is unnecessary param, default is DefaultConfigFile.

​	GetContent 返回指定 `file` 的自定义配置内容。 `file` 这是不必要的参数，默认值为 DefaultConfigFile。

#### (*AdapterFile) GetFileName

```go
func (a *AdapterFile) GetFileName() string
```

GetFileName returns the default configuration file name.

​	GetFileName 返回默认配置文件名。

#### (*AdapterFile) GetFilePath

```go
func (a *AdapterFile) GetFilePath(fileName ...string) (filePath string, err error)
```

GetFilePath returns the absolute configuration file path for the given filename by `file`. If `file` is not passed, it returns the configuration file path of the default name. It returns an empty `path` string and an error if the given `file` does not exist.

​	GetFilePath 返回给定文件名的 `file` 绝对配置文件路径。如果 `file` 未传递，则返回默认名称的配置文件路径。如果给定 `file` 的字符串不存在，则返回一个空 `path` 字符串和一个错误。

#### (*AdapterFile) GetPaths

```go
func (a *AdapterFile) GetPaths() []string
```

GetPaths returns the searching directory path array of current configuration manager.

​	GetPaths 返回当前配置管理器的搜索目录路径数组。

#### (*AdapterFile) MustGet

```go
func (a *AdapterFile) MustGet(ctx context.Context, pattern string) *gvar.Var
```

MustGet acts as function Get, but it panics if error occurs.

​	MustGet 充当函数 Get，但如果发生错误，它会崩溃。

#### (*AdapterFile) RemoveContent

```go
func (a *AdapterFile) RemoveContent(file ...string)
```

RemoveContent removes the global configuration with specified `file`. If `name` is not passed, it removes configuration of the default group name.

​	RemoveContent 删除具有指定 `file` .如果 `name` 未传递，则删除默认组名称的配置。

#### (*AdapterFile) Set

```go
func (a *AdapterFile) Set(pattern string, value interface{}) error
```

Set sets value with specified `pattern`. It supports hierarchical data access by char separator, which is ‘.’ in default. It is commonly used for updates certain configuration value in runtime. Note that, it is not recommended using `Set` configuration at runtime as the configuration would be automatically refreshed if underlying configuration file changed.

​	使用指定的 `pattern` 设置值。它支持通过字符分隔符进行分层数据访问，默认为 '..。它通常用于更新运行时中的某些配置值。请注意，不建议在运行时使用 `Set` 配置，因为如果基础配置文件发生更改，配置将自动刷新。

#### (*AdapterFile) SetContent

```go
func (a *AdapterFile) SetContent(content string, file ...string)
```

SetContent sets customized configuration content for specified `file`. The `file` is unnecessary param, default is DefaultConfigFile.

​	SetContent 为指定的 `file` . `file` 这是不必要的参数，默认值为 DefaultConfigFile。

#### (*AdapterFile) SetFileName

```go
func (a *AdapterFile) SetFileName(name string)
```

SetFileName sets the default configuration file name.

​	SetFileName 设置默认配置文件名。

#### (*AdapterFile) SetPath

```go
func (a *AdapterFile) SetPath(directoryPath string) (err error)
```

SetPath sets the configuration `directory` path for file search. The parameter `path` can be absolute or relative `directory` path, but absolute `directory` path is strongly recommended.

​	SetPath 设置文件搜索的配置 `directory` 路径。参数 `path` 可以是绝对路径或相对 `directory` 路径，但强烈建议使用绝对 `directory` 路径。

Note that this parameter is a path to a directory not a file.

​	请注意，此参数是目录的路径，而不是文件。

#### (*AdapterFile) SetViolenceCheck

```go
func (a *AdapterFile) SetViolenceCheck(check bool)
```

SetViolenceCheck sets whether to perform hierarchical conflict checking. This feature needs to be enabled when there is a level symbol in the key name. It is off in default.

​	SetViolenceCheck 设置是否执行分层冲突检查。当键名称中有级别符号时，需要启用此功能。默认情况下，它处于关闭状态。

Note that, turning on this feature is quite expensive, and it is not recommended allowing separators in the key names. It is best to avoid this on the application side.

​	请注意，启用此功能的成本非常高，并且不建议在键名中允许分隔符。最好在应用程序方面避免这种情况。

### type Config

```go
type Config struct {
	// contains filtered or unexported fields
}
```

Config is the configuration management object.

​	Config 是配置管理对象。

#### func Instance

```go
func Instance(name ...string) *Config
```

Instance returns an instance of Config with default settings. The parameter `name` is the name for the instance. But very note that, if the file “name.toml” exists in the configuration directory, it then sets it as the default configuration file. The toml file type is the default configuration file type.

​	Instance 返回具有默认设置的 Config 实例。该参数 `name` 是实例的名称。但需要注意的是，如果文件“name.toml”存在于配置目录中，则会将其设置为默认配置文件。toml 文件类型是默认配置文件类型。

#### func New

```go
func New() (*Config, error)
```

New creates and returns a Config object with default adapter of AdapterFile.

​	New 创建并返回一个默认适配器为 AdapterFile 的 Config 对象。

#### func NewWithAdapter

```go
func NewWithAdapter(adapter Adapter) *Config
```

NewWithAdapter creates and returns a Config object with given adapter.

​	NewWithAdapter 创建并返回具有给定适配器的 Config 对象。

#### (*Config) Available

```go
func (c *Config) Available(ctx context.Context, resource ...string) (ok bool)
```

Available checks and returns the configuration service is available. The optional parameter `pattern` specifies certain configuration resource.

​	available 检查并返回配置服务是否可用。可选参数 `pattern` 指定特定配置资源。

It returns true if configuration file is present in default AdapterFile, or else false. Note that this function does not return error as it just does simply check for backend configuration service.

​	如果配置文件存在于默认 AdapterFile 中，则返回 true，否则返回 false。请注意，此函数不会返回错误，因为它只是检查后端配置服务。

#### (*Config) Data

```go
func (c *Config) Data(ctx context.Context) (data map[string]interface{}, err error)
```

Data retrieves and returns all configuration data as map type.

​	数据检索并返回所有配置数据作为映射类型。

#### (*Config) Get

```go
func (c *Config) Get(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error)
```

Get retrieves and returns value by specified `pattern`. It returns all values of current Json object if `pattern` is given empty or string “.”. It returns nil if no value found by `pattern`.

​	按指定 `pattern` 获取检索和返回值。如果 `pattern` 给定空或字符串“.”，则返回当前 Json 对象的所有值。如果 未找到 的 `pattern` 值，则返回 nil。

It returns a default value specified by `def` if value for `pattern` is not found.

​	它返回由 `def` if value for 指定的默认值 if value `pattern` is not found。

#### (*Config) GetAdapter

```go
func (c *Config) GetAdapter() Adapter
```

GetAdapter returns the adapter of current Config object.

​	GetAdapter 返回当前 Config 对象的适配器。

#### (*Config) GetWithCmd

```go
func (c *Config) GetWithCmd(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error)
```

GetWithCmd returns the configuration value specified by pattern `pattern`. If the configuration value does not exist, then it retrieves and returns the command line option specified by `key`. It returns the default value `def` if none of them exists.

​	GetWithCmd 返回 pattern `pattern` 指定的配置值。如果配置值不存在，则检索并返回 指定的 `key` 命令行选项。如果它们都不存在，则返回默认值 `def` 。

Fetching Rules: Command line arguments are in lowercase format, eg: gf.package.variable.

​	获取规则：命令行参数采用小写格式，例如：gf.package.variable。

##### Example

``` go
```

#### (*Config) GetWithEnv

```go
func (c *Config) GetWithEnv(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error)
```

GetWithEnv returns the configuration value specified by pattern `pattern`. If the configuration value does not exist, then it retrieves and returns the environment value specified by `key`. It returns the default value `def` if none of them exists.

​	GetWithEnv 返回 pattern `pattern` 指定的配置值。如果配置值不存在，则检索并返回 指定的 `key` 环境值。如果它们都不存在，则返回默认值 `def` 。

Fetching Rules: Environment arguments are in uppercase format, eg: GF_PACKAGE_VARIABLE.

​	获取规则：环境参数为大写格式，例如：GF_PACKAGE_VARIABLE。

##### Example

``` go
```

#### (*Config) MustData

```go
func (c *Config) MustData(ctx context.Context) map[string]interface{}
```

MustData acts as function Data, but it panics if error occurs.

​	MustData 充当函数 Data，但如果发生错误，它会崩溃。

#### (*Config) MustGet

```go
func (c *Config) MustGet(ctx context.Context, pattern string, def ...interface{}) *gvar.Var
```

MustGet acts as function Get, but it panics if error occurs.

​	MustGet 充当函数 Get，但如果发生错误，它会崩溃。

#### (*Config) MustGetWithCmd

```go
func (c *Config) MustGetWithCmd(ctx context.Context, pattern string, def ...interface{}) *gvar.Var
```

MustGetWithCmd acts as function GetWithCmd, but it panics if error occurs.

​	MustGetWithCmd 充当函数 GetWithCmd，但如果发生错误，它会崩溃。

#### (*Config) MustGetWithEnv

```go
func (c *Config) MustGetWithEnv(ctx context.Context, pattern string, def ...interface{}) *gvar.Var
```

MustGetWithEnv acts as function GetWithEnv, but it panics if error occurs.

​	MustGetWithEnv 充当函数 GetWithEnv，但如果发生错误，它会崩溃。

#### (*Config) SetAdapter

```go
func (c *Config) SetAdapter(adapter Adapter)
```

SetAdapter sets the adapter of current Config object.

​	SetAdapter 设置当前 Config 对象的适配器。