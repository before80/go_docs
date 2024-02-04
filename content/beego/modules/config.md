+++
title = "config 模块"
date = 2024-02-04T09:32:20+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/module/config/]({{< ref "/beego/modules/config" >}})

# Config Module 配置模块



## Parsing Configuration Files 解析配置文件

The config module is used for parsing configuration files, inspired by `database/sql`. It supports ini, json, xml and yaml files. You can install it by:

​	config 模块用于解析配置文件，灵感来自 `database/sql` 。它支持 ini、json、xml 和 yaml 文件。您可以通过以下方式安装它：

```
go get github.com/beego/beego/v2/core/config
```

If you want to parse xml or yaml, you should first install:

​	如果您想解析 xml 或 yaml，您应该首先安装：

```
go get -u github.com/beego/beego/v2/core/config/xml
```

and then import:

​	然后导入：

```
import _ "github.com/beego/beego/v2/core/config/xml"
```

## Remote configure middleware 远程配置中间件

Now we support `etcd` as the implementation.

​	现在我们支持 `etcd` 作为实现。

## Usage 用法

### Using package 使用包

In v2.x, Beego create a `globalInstance`, so that users could use `config` module directly.

​	在 v2.x 中，Beego 创建了一个 `globalInstance` ，以便用户可以直接使用 `config` 模块。

```go
val, err := config.String("mykey")
```

Beego use `ini` implementation and loads config from `config/app.conf`.

​	Beego 使用 `ini` 实现并从 `config/app.conf` 加载配置。

If the file not found or got some error, Beego outputs some warning log.

​	如果找不到文件或出现错误，Beego 会输出一些警告日志。

Or you can initialize the `globalInstance` by:

​	或者，您可以通过以下方式初始化 `globalInstance` ：

```go
_ import "github.com/beego/beego/v2/core/config/toml"
err := InitGlobalInstance("toml", "some config")
// ...
val, err := config.String("mykey")
// ...
```

### Create instance manually 手动创建实例

Initialize a parser object:

​	初始化解析器对象：

```
iniconf, err := NewConfig("ini", "testini.conf")
if err != nil {
	t.Fatal(err)
}
```

Get data from parser:

​	从解析器获取数据：

```
iniconf.String("appname")
```

### Parser methods 解析器方法

Here are the parser’s methods:

​	以下是解析器的方法：

```
// Configer defines how to get and set value from configuration raw data.
type Configer interface {
    // support section::key type in given key when using ini type.
    Set(key, val string) error

    // support section::key type in key string when using ini and json type; Int,Int64,Bool,Float,DIY are same.
    String(key string) (string, error)
    // get string slice
    Strings(key string) ([]string, error)
    Int(key string) (int, error)
    Int64(key string) (int64, error)
    Bool(key string) (bool, error)
    Float(key string) (float64, error)
    // support section::key type in key string when using ini and json type; Int,Int64,Bool,Float,DIY are same.
    DefaultString(key string, defaultVal string) string
    // get string slice
    DefaultStrings(key string, defaultVal []string) []string
    DefaultInt(key string, defaultVal int) int
    DefaultInt64(key string, defaultVal int64) int64
    DefaultBool(key string, defaultVal bool) bool
    DefaultFloat(key string, defaultVal float64) float64

    // DIY return the original value
    DIY(key string) (interface{}, error)

    GetSection(section string) (map[string]string, error)

    Unmarshaler(prefix string, obj interface{}, opt ...DecodeOption) error
    Sub(key string) (Configer, error)
    OnChange(key string, fn func(value string))
    SaveConfigFile(filename string) error
}
```

Notice:

​	注意：

1. All `Default*` methods, default value will be returned if key not found or go some error;
   所有 `Default*` 方法，如果未找到键或出现错误，将返回默认值；
2. `DIY` returns original value. When you want to use this method, you should be care of value’s type.
   `DIY` 返回原始值。当您想使用此方法时，您应该注意值的数据类型。
3. `GetSection` returns all configure items under the `section`. `section` has different meaning in different implementation.
   `GetSection` 返回 `section` 下的所有配置项。 `section` 在不同的实现中具有不同的含义。
4. `Unmarshaler` try to decode the configs to `obj`. And the parameter `prefix` is similar with `section`.
   `Unmarshaler` 尝试将配置解码为 `obj` 。参数 `prefix` 与 `section` 类似。
5. `Sub` is similar to `GetSection`, but `Sub` will wrap result as an `Configer` instance.
   `Sub` 与 `GetSection` 类似，但 `Sub` 会将结果包装为 `Configer` 实例。
6. `Onchange` is used to listen change event. But most of file-based implementations don’t support this method.
   `Onchange` 用于监听更改事件。但大多数基于文件的实现都不支持此方法。
7. `SaveConfigFile` output configs to a file.
   `SaveConfigFile` 将输出配置写入文件。
8. Some implementation support key like `a.b.c.d`, but not all implementations support it and not all of them use `.` as separator.
   某些实现支持类似 `a.b.c.d` 的键，但并非所有实现都支持它，也并非所有实现都使用 `.` 作为分隔符。

### Configuration sections 配置节

The ini file supports configuration sections. You can get values inside a section by using `section::key`.

​	ini 文件支持配置节。您可以使用 `section::key` 获取节中的值。

For example:

​	例如：

```
[demo]
key1 = "asta"
key2 = "xie"
```

You can use `iniconf.String("demo::key2")` to get the value.

​	您可以使用 `iniconf.String("demo::key2")` 获取值。

### How to Obtain Environment Variables 如何获取环境变量

After Pull Request “Support get environment variables in config #1636” was merged into the code, beego supports using environment variables in the configuration file.

​	在将 Pull Request “在配置中支持获取环境变量 #1636” 合并到代码后，beego 支持在配置文件中使用环境变量。

The format for this is `${ENVIRONMENTVARIABLE}` within the configuration file which is equivalent to `value = os.Getenv('ENVIRONMENTVARIABLE')`. Beego will only check for environment variables if the value begins with `${` and ends with `}`.

​	此格式在配置文件中为 `${ENVIRONMENTVARIABLE}` ，等效于 `value = os.Getenv('ENVIRONMENTVARIABLE')` 。Beego 仅在值以 `${` 开头并以 `}` 结尾时检查环境变量。

Additionally, a default value can be configured for the case that there is no environment variable set or the environment variable is empty. This is accomplished by using the format `${ENVVAR||defaultvalue}`, for example `${GOPATH||/home/asataxie/workspace/go}`. This `||` is used to split environment values and default values. See `/config/config_test.go` in the [beego repo](https://github.com/beego/beego) for more examples and edge cases about how these environment variables and default values are parsed.

​	此外，还可以为没有设置环境变量或环境变量为空的情况配置一个默认值。这可以通过使用格式 `${ENVVAR||defaultvalue}` 来实现，例如 `${GOPATH||/home/asataxie/workspace/go}` 。此 `||` 用于分割环境值和默认值。有关如何解析这些环境变量和默认值的更多示例和边缘情况，请参阅 beego 仓库中的 `/config/config_test.go` 。

For example:

​	例如：

```
password = ${MyPWD}
token = ${TOKEN||astaxie}
user = ${MyUser||beego}
```

If the environment variable `$TOKEN` is set, its value will be used for the `token` configuration value and `beego.AppConfig.String("token")` would return its value. If `$TOKEN` is not set, the value would then be the string `astaxie`.

​	如果设置了环境变量 `$TOKEN` ，则其值将用于 `token` 配置值，并且 `beego.AppConfig.String("token")` 将返回其值。如果未设置 `$TOKEN` ，则该值将为字符串 `astaxie` 。

**Please note**: The environment variables are only read when the configuration file is parsed, not when configuration item is obtained by a function like `beego.AppConfig.String(string)`.

​	请注意：仅在解析配置文件时才会读取环境变量，而不是通过函数（如 `beego.AppConfig.String(string)` ）获取配置项时读取环境变量。
