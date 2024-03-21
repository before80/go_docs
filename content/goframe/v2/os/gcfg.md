+++
title = "gcfg"
date = 2024-03-21T17:54:46+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcfg

Package gcfg provides reading, caching and managing for configuration.

#### Examples 

- [Config.GetWithCmd](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcfg#example-Config.GetWithCmd)
- [Config.GetWithEnv](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcfg#example-Config.GetWithEnv)

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gcfg/gcfg.go#L27)

``` go
const (
	DefaultInstanceName   = "config" // DefaultName is the default instance name for instance usage.
	DefaultConfigFileName = "config" // DefaultConfigFile is the default configuration file name.
)
```

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type Adapter 

``` go
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

#### type AdapterContent <-2.5.3

``` go
type AdapterContent struct {
	// contains filtered or unexported fields
}
```

AdapterContent implements interface Adapter using content. The configuration content supports the coding types as package `gjson`.

##### func NewAdapterContent <-2.5.3

``` go
func NewAdapterContent(content ...string) (*AdapterContent, error)
```

NewAdapterContent returns a new configuration management object using custom content. The parameter `content` specifies the default configuration content for reading.

##### (*AdapterContent) Available <-2.5.3

``` go
func (a *AdapterContent) Available(ctx context.Context, resource ...string) (ok bool)
```

Available checks and returns the backend configuration service is available. The optional parameter `resource` specifies certain configuration resource.

Note that this function does not return error as it just does simply check for backend configuration service.

##### (*AdapterContent) Data <-2.5.3

``` go
func (a *AdapterContent) Data(ctx context.Context) (data map[string]interface{}, err error)
```

Data retrieves and returns all configuration data in current resource as map. Note that this function may lead lots of memory usage if configuration data is too large, you can implement this function if necessary.

##### (*AdapterContent) Get <-2.5.3

``` go
func (a *AdapterContent) Get(ctx context.Context, pattern string) (value interface{}, err error)
```

Get retrieves and returns value by specified `pattern` in current resource. Pattern like: "x.y.z" for map item. "x.0.y" for slice item.

##### (*AdapterContent) SetContent <-2.5.3

``` go
func (a *AdapterContent) SetContent(content string) error
```

SetContent sets customized configuration content for specified `file`. The `file` is unnecessary param, default is DefaultConfigFile.

#### type AdapterFile 

``` go
type AdapterFile struct {
	// contains filtered or unexported fields
}
```

AdapterFile implements interface Adapter using file.

##### func NewAdapterFile 

``` go
func NewAdapterFile(file ...string) (*AdapterFile, error)
```

NewAdapterFile returns a new configuration management object. The parameter `file` specifies the default configuration file name for reading.

##### (*AdapterFile) AddPath 

``` go
func (a *AdapterFile) AddPath(directoryPaths ...string) (err error)
```

AddPath adds an absolute or relative `directory` path to the search paths.

Note that this parameter is paths to a directories not files.

##### (*AdapterFile) Available 

``` go
func (a *AdapterFile) Available(ctx context.Context, fileName ...string) bool
```

Available checks and returns whether configuration of given `file` is available.

##### (*AdapterFile) Clear 

``` go
func (a *AdapterFile) Clear()
```

Clear removes all parsed configuration files content cache, which will force reload configuration content from file.

##### (*AdapterFile) ClearContent 

``` go
func (a *AdapterFile) ClearContent()
```

ClearContent removes all global configuration contents.

##### (*AdapterFile) Data 

``` go
func (a *AdapterFile) Data(ctx context.Context) (data map[string]interface{}, err error)
```

Data retrieves and returns all configuration data as map type.

##### (*AdapterFile) Dump 

``` go
func (a *AdapterFile) Dump()
```

Dump prints current Json object with more manually readable.

##### (*AdapterFile) Get 

``` go
func (a *AdapterFile) Get(ctx context.Context, pattern string) (value interface{}, err error)
```

Get retrieves and returns value by specified `pattern`. It returns all values of current Json object if `pattern` is given empty or string ".". It returns nil if no value found by `pattern`.

We can also access slice item by its index number in `pattern` like: "list.10", "array.0.name", "array.0.1.id".

It returns a default value specified by `def` if value for `pattern` is not found.

##### (*AdapterFile) GetContent 

``` go
func (a *AdapterFile) GetContent(file ...string) string
```

GetContent returns customized configuration content for specified `file`. The `file` is unnecessary param, default is DefaultConfigFile.

##### (*AdapterFile) GetFileName 

``` go
func (a *AdapterFile) GetFileName() string
```

GetFileName returns the default configuration file name.

##### (*AdapterFile) GetFilePath 

``` go
func (a *AdapterFile) GetFilePath(fileName ...string) (filePath string, err error)
```

GetFilePath returns the absolute configuration file path for the given filename by `file`. If `file` is not passed, it returns the configuration file path of the default name. It returns an empty `path` string and an error if the given `file` does not exist.

##### (*AdapterFile) GetPaths <-2.0.5

``` go
func (a *AdapterFile) GetPaths() []string
```

GetPaths returns the searching directory path array of current configuration manager.

##### (*AdapterFile) MustGet 

``` go
func (a *AdapterFile) MustGet(ctx context.Context, pattern string) *gvar.Var
```

MustGet acts as function Get, but it panics if error occurs.

##### (*AdapterFile) RemoveContent 

``` go
func (a *AdapterFile) RemoveContent(file ...string)
```

RemoveContent removes the global configuration with specified `file`. If `name` is not passed, it removes configuration of the default group name.

##### (*AdapterFile) Set <-2.0.5

``` go
func (a *AdapterFile) Set(pattern string, value interface{}) error
```

Set sets value with specified `pattern`. It supports hierarchical data access by char separator, which is '.' in default. It is commonly used for updates certain configuration value in runtime. Note that, it is not recommended using `Set` configuration at runtime as the configuration would be automatically refreshed if underlying configuration file changed.

##### (*AdapterFile) SetContent 

``` go
func (a *AdapterFile) SetContent(content string, file ...string)
```

SetContent sets customized configuration content for specified `file`. The `file` is unnecessary param, default is DefaultConfigFile.

##### (*AdapterFile) SetFileName 

``` go
func (a *AdapterFile) SetFileName(name string)
```

SetFileName sets the default configuration file name.

##### (*AdapterFile) SetPath 

``` go
func (a *AdapterFile) SetPath(directoryPath string) (err error)
```

SetPath sets the configuration `directory` path for file search. The parameter `path` can be absolute or relative `directory` path, but absolute `directory` path is strongly recommended.

Note that this parameter is a path to a directory not a file.

##### (*AdapterFile) SetViolenceCheck 

``` go
func (a *AdapterFile) SetViolenceCheck(check bool)
```

SetViolenceCheck sets whether to perform hierarchical conflict checking. This feature needs to be enabled when there is a level symbol in the key name. It is off in default.

Note that, turning on this feature is quite expensive, and it is not recommended allowing separators in the key names. It is best to avoid this on the application side.

#### type Config 

``` go
type Config struct {
	// contains filtered or unexported fields
}
```

Config is the configuration management object.

##### func Instance 

``` go
func Instance(name ...string) *Config
```

Instance returns an instance of Config with default settings. The parameter `name` is the name for the instance. But very note that, if the file "name.toml" exists in the configuration directory, it then sets it as the default configuration file. The toml file type is the default configuration file type.

##### func New 

``` go
func New() (*Config, error)
```

New creates and returns a Config object with default adapter of AdapterFile.

##### func NewWithAdapter 

``` go
func NewWithAdapter(adapter Adapter) *Config
```

NewWithAdapter creates and returns a Config object with given adapter.

##### (*Config) Available 

``` go
func (c *Config) Available(ctx context.Context, resource ...string) (ok bool)
```

Available checks and returns the configuration service is available. The optional parameter `pattern` specifies certain configuration resource.

It returns true if configuration file is present in default AdapterFile, or else false. Note that this function does not return error as it just does simply check for backend configuration service.

##### (*Config) Data 

``` go
func (c *Config) Data(ctx context.Context) (data map[string]interface{}, err error)
```

Data retrieves and returns all configuration data as map type.

##### (*Config) Get 

``` go
func (c *Config) Get(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error)
```

Get retrieves and returns value by specified `pattern`. It returns all values of current Json object if `pattern` is given empty or string ".". It returns nil if no value found by `pattern`.

It returns a default value specified by `def` if value for `pattern` is not found.

##### (*Config) GetAdapter 

``` go
func (c *Config) GetAdapter() Adapter
```

GetAdapter returns the adapter of current Config object.

##### (*Config) GetWithCmd 

``` go
func (c *Config) GetWithCmd(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error)
```

GetWithCmd returns the configuration value specified by pattern `pattern`. If the configuration value does not exist, then it retrieves and returns the command line option specified by `key`. It returns the default value `def` if none of them exists.

Fetching Rules: Command line arguments are in lowercase format, eg: gf.package.variable.

##### Example

``` go
```
##### (*Config) GetWithEnv 

``` go
func (c *Config) GetWithEnv(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error)
```

GetWithEnv returns the configuration value specified by pattern `pattern`. If the configuration value does not exist, then it retrieves and returns the environment value specified by `key`. It returns the default value `def` if none of them exists.

Fetching Rules: Environment arguments are in uppercase format, eg: GF_PACKAGE_VARIABLE.

##### Example

``` go
```
##### (*Config) MustData 

``` go
func (c *Config) MustData(ctx context.Context) map[string]interface{}
```

MustData acts as function Data, but it panics if error occurs.

##### (*Config) MustGet 

``` go
func (c *Config) MustGet(ctx context.Context, pattern string, def ...interface{}) *gvar.Var
```

MustGet acts as function Get, but it panics if error occurs.

##### (*Config) MustGetWithCmd 

``` go
func (c *Config) MustGetWithCmd(ctx context.Context, pattern string, def ...interface{}) *gvar.Var
```

MustGetWithCmd acts as function GetWithCmd, but it panics if error occurs.

##### (*Config) MustGetWithEnv 

``` go
func (c *Config) MustGetWithEnv(ctx context.Context, pattern string, def ...interface{}) *gvar.Var
```

MustGetWithEnv acts as function GetWithEnv, but it panics if error occurs.

##### (*Config) SetAdapter 

``` go
func (c *Config) SetAdapter(adapter Adapter)
```

SetAdapter sets the adapter of current Config object.