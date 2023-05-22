+++
title = "viper"
date = 2023-05-22T08:46:24+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# viper

> ### Viper v2 反馈
>
> ​	Viper 正朝着 v2 的方向发展，我们非常愿意听到您对于 v2 的期望。请在此处分享您的想法：[https://forms.gle/R6faU74qPRPAzchZ9](https://forms.gle/R6faU74qPRPAzchZ9)
>
> **Thank you!**

![Viper](_index_img/logo.png)

![image-20230522084952536](_index_img/image-20230522084952536.png)

**带有尖牙的 Go 配置 - Go configuration with fangs!**

​	许多 Go 项目都使用 Viper 构建，包括：

- [Hugo](http://gohugo.io/)
- [EMC RexRay](http://rexray.readthedocs.org/en/stable/)
- [Imgur’s Incus](https://github.com/Imgur/incus)
- [Nanobox](https://github.com/nanobox-io/nanobox)/[Nanopack](https://github.com/nanopack)
- [Docker Notary](https://github.com/docker/Notary)
- [BloomApi](https://www.bloomapi.com/)
- [doctl](https://github.com/digitalocean/doctl)
- [Clairctl](https://github.com/jgsqware/clairctl)
- [Mercure](https://mercure.rocks/)

## 安装

```
go get github.com/spf13/viper
```

**注意：** Viper 使用 [Go 模块](https://github.com/golang/go/wiki/Modules) 来管理依赖项。

## Viper 是什么？

​	Viper 是 Go 应用程序的完整配置解决方案，包括[12-Factor 应用程序](https://12factor.net/zh_cn/#12factors)。它被设计用于在应用程序中工作，并可以处理各种类型的配置需求和格式。它支持：

- 设置默认值
- 从 JSON、TOML、YAML、HCL、envfile 和 Java properties 配置文件中读取配置
- 实时监视和重新读取配置文件（可选）
- 从环境变量中读取配置
- 从远程配置系统（etcd 或 Consul）中读取配置并监视更改
- 从命令行标志中读取配置
- 从缓冲区中读取配置
- 设置显式值

​	可以将 Viper 视为你的应用程序所有配置需求的注册表。

## 为什么选择 Viper？

​	在构建现代应用程序时，您不希望担心配置文件格式，而是专注于构建出色的软件。Viper 就是为此而存在的。

​	Viper 为您提供以下功能：

1. 查找、加载和解析 JSON、TOML、YAML、HCL、INI、envfile 或 Java properties 格式的配置文件。
2. 提供一种机制（mechanism）来为不同的配置选项设置默认值。
3. 提供一种机制（mechanism）来为通过命令行标志指定的选项设置覆盖值。
4. 提供别名系统，以便轻松重命名参数而不会破坏现有代码。
5. 简化判断用户是提供了与默认值相同的命令行参数还是配置文件的区别。

​	Viper 使用以下优先顺序。每个条目的优先级高于它下面的条目：

- 显式调用 `Set`
- flag
- env
- config
- key/value存储
- 默认值

**重要提示：** Viper 的配置键不区分大小写。关于是否可选地进行区分大小写，目前正在进行讨论。

## Putting Values into Viper 向 Viper 中添加值

### 设置默认值

A good configuration system will support default values. A default value is not required for a key, but it’s useful in the event that a key hasn't been set via config file, environment variable, remote configuration or flag.

一个良好的配置系统将支持默认值。对于一个键来说，默认值不是必需的，但在没有通过配置文件、环境变量、远程配置或命令行标志设置键时非常有用。

示例：

```
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
```

### Reading Config Files 读取配置文件

Viper requires minimal configuration so it knows where to look for config files. Viper supports JSON, TOML, YAML, HCL, INI, envfile and Java Properties files. Viper can search multiple paths, but currently a single Viper instance only supports a single configuration file. Viper does not default to any configuration search paths leaving defaults decision to an application.

Viper 需要最小的配置来知道在哪里查找配置文件。Viper 支持 JSON、TOML、YAML、HCL、INI、envfile 和 Java Properties 文件。Viper 可以搜索多个路径，但是当前每个 Viper 实例只支持一个配置文件。Viper 不会默认设置任何配置文件搜索路径，而是将默认的决策留给应用程序。

Here is an example of how to use Viper to search for and read a configuration file. None of the specific paths are required, but at least one path should be provided where a configuration file is expected.

以下是如何使用 Viper 搜索并读取配置文件的示例。没有特定的路径是必需的，但至少应提供一个期望找到配置文件的路径。

```
viper.SetConfigName("config") // name of config file (without extension) 配置文件的名称（无扩展名）
viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name 如果配置文件名称中不包含扩展名，则必需
viper.AddConfigPath("/etc/appname/")   // path to look for the config file in 配置文件的搜索路径
viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths 可多次调用以添加多个搜索路径
viper.AddConfigPath(".")               // optionally look for config in the working directory 可选地在工作目录中查找配置文件
err := viper.ReadInConfig() // Find and read the config file  查找并读取配置文件
if err != nil { // Handle errors reading the config file 处理读取配置文件时的错误
	panic(fmt.Errorf("fatal error config file: %w", err))
}
```

You can handle the specific case where no config file is found like this:

​	您可以像这样处理未找到配置文件的特殊情况：

```
if err := viper.ReadInConfig(); err != nil {
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config file not found; ignore error if desired 未找到配置文件；如果需要，可以忽略错误
	} else {
		// Config file was found but another error was produced 找到配置文件，但产生了其他错误
	}
}

// Config file found and successfully parsed
```

*NOTE [since 1.6]:* You can also have a file without an extension and specify the format programmaticaly. For those configuration files that lie in the home of the user without any extension like `.bashrc`

*注意 [自版本 1.6 起]：* 您还可以拥有没有扩展名的文件，并以编程方式指定格式。对于那些位于用户主目录下且没有任何扩展名的配置文件，例如 `.bashrc`。



### Writing Config Files 写入配置文件

Reading from config files is useful, but at times you want to store all modifications made at run time. For that, a bunch of commands are available, each with its own purpose:

从配置文件中读取很有用，但有时您希望存储运行时所做的所有修改。为此，有一系列可用的命令，每个命令都有自己的用途：

- WriteConfig - writes the current viper configuration to the predefined path, if exists. Errors if no predefined path. Will overwrite the current config file, if it exists.
- SafeWriteConfig - writes the current viper configuration to the predefined path. Errors if no predefined path. Will not overwrite the current config file, if it exists.
- WriteConfigAs - writes the current viper configuration to the given filepath. Will overwrite the given file, if it exists.
- SafeWriteConfigAs - writes the current viper configuration to the given filepath. Will not overwrite the given file, if it exists.
- WriteConfig - 将当前的 Viper 配置写入预定义的路径（如果存在）。如果没有预定义的路径，则报错。如果配置文件已存在，则会覆盖当前的配置文件。
- SafeWriteConfig - 将当前的 Viper 配置写入预定义的路径。如果没有预定义的路径，则报错。如果配置文件已存在，则不会覆盖当前的配置文件。
- WriteConfigAs - 将当前的 Viper 配置写入指定的文件路径。如果给定的文件已存在，则会覆盖该文件。
- SafeWriteConfigAs - 将当前的 Viper 配置写入指定的文件路径。如果给定的文件已存在，则不会覆盖该文件。

As a rule of the thumb, everything marked with safe won't overwrite any file, but just create if not existent, whilst the default behavior is to create or truncate.

根据经验，标记为 safe 的所有操作都不会覆盖任何文件，只会在文件不存在时创建，而默认行为是创建或截断文件。

以下是一个小示例：

A small examples section:

```
viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName' 将当前配置写入由 'viper.AddConfigPath()' 和 'viper.SetConfigName' 设置的预定义路径
viper.SafeWriteConfig()
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") // will error since it has already been written 会报错，因为已经写入过
viper.SafeWriteConfigAs("/path/to/my/.other_config")
```

### Watching and re-reading config files 监听和重新读取配置文件

Viper supports the ability to have your application live read a config file while running.

Viper 支持在运行时实时读取配置文件的能力。

Gone are the days of needing to restart a server to have a config take effect, viper powered applications can read an update to a config file while running and not miss a beat.



不再需要重新启动服务器以使配置生效，Viper 驱动的应用程序可以在运行时读取配置文件的更新而不会出现中断。

Simply tell the viper instance to watchConfig. Optionally you can provide a function for Viper to run each time a change occurs.

只需告诉 Viper 实例执行 `WatchConfig()`。可选地，您还可以为 Viper 提供一个在每次更改发生时运行的函数。

**Make sure you add all of the configPaths prior to calling `WatchConfig()`**

**确保在调用 `WatchConfig()` 之前添加所有的配置路径。**

```
viper.OnConfigChange(func(e fsnotify.Event) {
	fmt.Println("Config file changed:", e.Name)
})
viper.WatchConfig()
```

### Reading Config from io.Reader 从 io.Reader 读取配置

Viper predefines many configuration sources such as files, environment variables, flags, and remote K/V store, but you are not bound to them. You can also implement your own required configuration source and feed it to viper.

Viper 预定义了许多配置源，如文件、环境变量、标志和远程键值存储，但您并不受限于此。您还可以实现自己所需的配置源，并将其提供给 Viper。

```
viper.SetConfigType("yaml") // or viper.SetConfigType("YAML") 或 viper.SetConfigType("YAML")

// any approach to require this configuration into your program. 任何方法将此配置需求输入到您的程序中。
var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

viper.ReadConfig(bytes.NewBuffer(yamlExample))

viper.Get("name") // this would be "steve"  这将返回 "steve"
```

### Setting Overrides 设置覆盖值

These could be from a command line flag, or from your own application logic.

这些可以来自命令行标志，或者来自您自己的应用程序逻辑。

```
viper.Set("Verbose", true)
viper.Set("LogFile", LogFile)
```

### Registering and Using Aliases 注册和使用别名

Aliases permit a single value to be referenced by multiple keys

别名允许多个键引用同一个值。

```
viper.RegisterAlias("loud", "Verbose")

viper.Set("verbose", true) // same result as next line 与下一行代码的结果相同
viper.Set("loud", true)   // same result as prior line 与上一行代码的结果相同

viper.GetBool("loud") // true 返回 true
viper.GetBool("verbose") // true 返回 true
```

### Working with Environment Variables 使用环境变量

Viper has full support for environment variables. This enables 12 factor applications out of the box. There are five methods that exist to aid working with ENV:

Viper 对环境变量提供了全面支持。这使得可以直接使用12因子应用程序。有五个方法可用于处理环境变量：

- `AutomaticEnv()`
- `BindEnv(string...) : error`
- `SetEnvPrefix(string)`
- `SetEnvKeyReplacer(string...) *strings.Replacer`
- `AllowEmptyEnv(bool)`

*When working with ENV variables, it’s important to recognize that Viper treats ENV variables as case sensitive.*

*在处理环境变量时，需要注意 Viper 将其视为区分大小写的。*

Viper provides a mechanism to try to ensure that ENV variables are unique. By using `SetEnvPrefix`, you can tell Viper to use a prefix while reading from the environment variables. Both `BindEnv` and `AutomaticEnv` will use this prefix.

`BindEnv` takes one or more parameters. The first parameter is the key name, the rest are the name of the environment variables to bind to this key. If more than one are provided, they will take precedence in the specified order. The name of the environment variable is case sensitive. If the ENV variable name is not provided, then Viper will automatically assume that the ENV variable matches the following format: prefix + "_" + the key name in ALL CAPS. When you explicitly provide the ENV variable name (the second parameter), it **does not** automatically add the prefix. For example if the second parameter is "id", Viper will look for the ENV variable "ID".

One important thing to recognize when working with ENV variables is that the value will be read each time it is accessed. Viper does not fix the value when the `BindEnv` is called.

`AutomaticEnv` is a powerful helper especially when combined with `SetEnvPrefix`. When called, Viper will check for an environment variable any time a `viper.Get` request is made. It will apply the following rules. It will check for an environment variable with a name matching the key uppercased and prefixed with the `EnvPrefix` if set.

`SetEnvKeyReplacer` allows you to use a `strings.Replacer` object to rewrite Env keys to an extent. This is useful if you want to use `-` or something in your `Get()` calls, but want your environmental variables to use `_` delimiters. An example of using it can be found in `viper_test.go`.

Alternatively, you can use `EnvKeyReplacer` with `NewWithOptions` factory function. Unlike `SetEnvKeyReplacer`, it accepts a `StringReplacer` interface allowing you to write custom string replacing logic.

By default empty environment variables are considered unset and will fall back to the next configuration source. To treat empty environment variables as set, use the `AllowEmptyEnv` method.

#### Env example

```
SetEnvPrefix("spf") // will be uppercased automatically
BindEnv("id")

os.Setenv("SPF_ID", "13") // typically done outside of the app

id := Get("id") // 13
```

### Working with Flags

Viper has the ability to bind to flags. Specifically, Viper supports `Pflags` as used in the [Cobra](https://github.com/spf13/cobra) library.

Like `BindEnv`, the value is not set when the binding method is called, but when it is accessed. This means you can bind as early as you want, even in an `init()` function.

For individual flags, the `BindPFlag()` method provides this functionality.

Example:

```
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
```

You can also bind an existing set of pflags (pflag.FlagSet):

Example:

```
pflag.Int("flagname", 1234, "help message for flagname")

pflag.Parse()
viper.BindPFlags(pflag.CommandLine)

i := viper.GetInt("flagname") // retrieve values from viper instead of pflag
```

The use of [pflag](https://github.com/spf13/pflag/) in Viper does not preclude the use of other packages that use the [flag](https://golang.org/pkg/flag/) package from the standard library. The pflag package can handle the flags defined for the flag package by importing these flags. This is accomplished by a calling a convenience function provided by the pflag package called AddGoFlagSet().

Example:

```
package main

import (
	"flag"
	"github.com/spf13/pflag"
)

func main() {

	// using standard library "flag" package
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") // retrieve value from viper

	// ...
}
```

#### Flag interfaces

Viper provides two Go interfaces to bind other flag systems if you don’t use `Pflags`.

`FlagValue` represents a single flag. This is a very simple example on how to implement this interface:

```
type myFlag struct {}
func (f myFlag) HasChanged() bool { return false }
func (f myFlag) Name() string { return "my-flag-name" }
func (f myFlag) ValueString() string { return "my-flag-value" }
func (f myFlag) ValueType() string { return "string" }
```

Once your flag implements this interface, you can simply tell Viper to bind it:

```
viper.BindFlagValue("my-flag-name", myFlag{})
```

`FlagValueSet` represents a group of flags. This is a very simple example on how to implement this interface:

```
type myFlagSet struct {
	flags []myFlag
}

func (f myFlagSet) VisitAll(fn func(FlagValue)) {
	for _, flag := range flags {
		fn(flag)
	}
}
```

Once your flag set implements this interface, you can simply tell Viper to bind it:

```
fSet := myFlagSet{
	flags: []myFlag{myFlag{}, myFlag{}},
}
viper.BindFlagValues("my-flags", fSet)
```

### Remote Key/Value Store Support

To enable remote support in Viper, do a blank import of the `viper/remote` package:

```
import _ "github.com/spf13/viper/remote"
```

Viper will read a config string (as JSON, TOML, YAML, HCL or envfile) retrieved from a path in a Key/Value store such as etcd or Consul. These values take precedence over default values, but are overridden by configuration values retrieved from disk, flags, or environment variables.

Viper uses [crypt](https://github.com/bketelsen/crypt) to retrieve configuration from the K/V store, which means that you can store your configuration values encrypted and have them automatically decrypted if you have the correct gpg keyring. Encryption is optional.

You can use remote configuration in conjunction with local configuration, or independently of it.

`crypt` has a command-line helper that you can use to put configurations in your K/V store. `crypt` defaults to etcd on http://127.0.0.1:4001.

```
$ go get github.com/bketelsen/crypt/bin/crypt
$ crypt set -plaintext /config/hugo.json /Users/hugo/settings/config.json
```

Confirm that your value was set:

```
$ crypt get -plaintext /config/hugo.json
```

See the `crypt` documentation for examples of how to set encrypted values, or how to use Consul.

### Remote Key/Value Store Example - Unencrypted

#### etcd

```
viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()
```

#### etcd3

```
viper.AddRemoteProvider("etcd3", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()
```

#### Consul

You need to set a key to Consul key/value storage with JSON value containing your desired config. For example, create a Consul key/value store key `MY_CONSUL_KEY` with value:

```
{
    "port": 8080,
    "hostname": "myhostname.com"
}
viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
viper.SetConfigType("json") // Need to explicitly set this to json
err := viper.ReadRemoteConfig()

fmt.Println(viper.Get("port")) // 8080
fmt.Println(viper.Get("hostname")) // myhostname.com
```

#### Firestore

```
viper.AddRemoteProvider("firestore", "google-cloud-project-id", "collection/document")
viper.SetConfigType("json") // Config's format: "json", "toml", "yaml", "yml"
err := viper.ReadRemoteConfig()
```

Of course, you're allowed to use `SecureRemoteProvider` also

### Remote Key/Value Store Example - Encrypted

```
viper.AddSecureRemoteProvider("etcd","http://127.0.0.1:4001","/config/hugo.json","/etc/secrets/mykeyring.gpg")
viper.SetConfigType("json") // because there is no file extension in a stream of bytes,  supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()
```

### Watching Changes in etcd - Unencrypted

```
// alternatively, you can create a new viper instance.
var runtime_viper = viper.New()

runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
runtime_viper.SetConfigType("yaml") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"

// read from remote config the first time.
err := runtime_viper.ReadRemoteConfig()

// unmarshal config
runtime_viper.Unmarshal(&runtime_conf)

// open a goroutine to watch remote changes forever
go func(){
	for {
		time.Sleep(time.Second * 5) // delay after each request

		// currently, only tested with etcd support
		err := runtime_viper.WatchRemoteConfig()
		if err != nil {
			log.Errorf("unable to read remote config: %v", err)
			continue
		}

		// unmarshal new config into our runtime config struct. you can also use channel
		// to implement a signal to notify the system of the changes
		runtime_viper.Unmarshal(&runtime_conf)
	}
}()
```

## Getting Values From Viper

In Viper, there are a few ways to get a value depending on the value’s type. The following functions and methods exist:

- `Get(key string) : interface{}`
- `GetBool(key string) : bool`
- `GetFloat64(key string) : float64`
- `GetInt(key string) : int`
- `GetIntSlice(key string) : []int`
- `GetString(key string) : string`
- `GetStringMap(key string) : map[string]interface{}`
- `GetStringMapString(key string) : map[string]string`
- `GetStringSlice(key string) : []string`
- `GetTime(key string) : time.Time`
- `GetDuration(key string) : time.Duration`
- `IsSet(key string) : bool`
- `AllSettings() : map[string]interface{}`

One important thing to recognize is that each Get function will return a zero value if it’s not found. To check if a given key exists, the `IsSet()` method has been provided.

Example:

```
viper.GetString("logfile") // case-insensitive Setting & Getting
if viper.GetBool("verbose") {
	fmt.Println("verbose enabled")
}
```

### Accessing nested keys

The accessor methods also accept formatted paths to deeply nested keys. For example, if the following JSON file is loaded:

```
{
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}
```

Viper can access a nested field by passing a `.` delimited path of keys:

```
GetString("datastore.metric.host") // (returns "127.0.0.1")
```

This obeys the precedence rules established above; the search for the path will cascade through the remaining configuration registries until found.

For example, given this configuration file, both `datastore.metric.host` and `datastore.metric.port` are already defined (and may be overridden). If in addition `datastore.metric.protocol` was defined in the defaults, Viper would also find it.

However, if `datastore.metric` was overridden (by a flag, an environment variable, the `Set()` method, …) with an immediate value, then all sub-keys of `datastore.metric` become undefined, they are “shadowed” by the higher-priority configuration level.

Viper can access array indices by using numbers in the path. For example:

```
{
    "host": {
        "address": "localhost",
        "ports": [
            5799,
            6029
        ]
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}

GetInt("host.ports.1") // returns 6029
```

Lastly, if there exists a key that matches the delimited key path, its value will be returned instead. E.g.

```
{
    "datastore.metric.host": "0.0.0.0",
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}

GetString("datastore.metric.host") // returns "0.0.0.0"
```

### Extracting a sub-tree

When developing reusable modules, it's often useful to extract a subset of the configuration and pass it to a module. This way the module can be instantiated more than once, with different configurations.

For example, an application might use multiple different cache stores for different purposes:

```
cache:
  cache1:
    max-items: 100
    item-size: 64
  cache2:
    max-items: 200
    item-size: 80
```

We could pass the cache name to a module (eg. `NewCache("cache1")`), but it would require weird concatenation for accessing config keys and would be less separated from the global config.

So instead of doing that let's pass a Viper instance to the constructor that represents a subset of the configuration:

```
cache1Config := viper.Sub("cache.cache1")
if cache1Config == nil { // Sub returns nil if the key cannot be found
	panic("cache configuration not found")
}

cache1 := NewCache(cache1Config)
```

**Note:** Always check the return value of `Sub`. It returns `nil` if a key cannot be found.

Internally, the `NewCache` function can address `max-items` and `item-size` keys directly:

```
func NewCache(v *Viper) *Cache {
	return &Cache{
		MaxItems: v.GetInt("max-items"),
		ItemSize: v.GetInt("item-size"),
	}
}
```

The resulting code is easy to test, since it's decoupled from the main config structure, and easier to reuse (for the same reason).

### Unmarshaling

You also have the option of Unmarshaling all or a specific value to a struct, map, etc.

There are two methods to do this:

- `Unmarshal(rawVal interface{}) : error`
- `UnmarshalKey(key string, rawVal interface{}) : error`

Example:

```
type config struct {
	Port int
	Name string
	PathMap string `mapstructure:"path_map"`
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("unable to decode into struct, %v", err)
}
```

If you want to unmarshal configuration where the keys themselves contain dot (the default key delimiter), you have to change the delimiter:

```
v := viper.NewWithOptions(viper.KeyDelimiter("::"))

v.SetDefault("chart::values", map[string]interface{}{
	"ingress": map[string]interface{}{
		"annotations": map[string]interface{}{
			"traefik.frontend.rule.type":                 "PathPrefix",
			"traefik.ingress.kubernetes.io/ssl-redirect": "true",
		},
	},
})

type config struct {
	Chart struct{
		Values map[string]interface{}
	}
}

var C config

v.Unmarshal(&C)
```

Viper also supports unmarshaling into embedded structs:

```
/*
Example config:

module:
    enabled: true
    token: 89h3f98hbwf987h3f98wenf89ehf
*/
type config struct {
	Module struct {
		Enabled bool

		moduleConfig `mapstructure:",squash"`
	}
}

// moduleConfig could be in a module specific package
type moduleConfig struct {
	Token string
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("unable to decode into struct, %v", err)
}
```

Viper uses [github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) under the hood for unmarshaling values which uses `mapstructure` tags by default.

### Decoding custom formats

A frequently requested feature for Viper is adding more value formats and decoders. For example, parsing character (dot, comma, semicolon, etc) separated strings into slices.

This is already available in Viper using mapstructure decode hooks.

Read more about the details in [this blog post](https://sagikazarmark.hu/blog/decoding-custom-formats-with-viper/).

### Marshalling to string

You may need to marshal all the settings held in viper into a string rather than write them to a file. You can use your favorite format's marshaller with the config returned by `AllSettings()`.

```
import (
	yaml "gopkg.in/yaml.v2"
	// ...
)

func yamlStringSettings() string {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	return string(bs)
}
```

## Viper or Vipers?

Viper comes ready to use out of the box. There is no configuration or initialization needed to begin using Viper. Since most applications will want to use a single central repository for their configuration, the viper package provides this. It is similar to a singleton.

In all of the examples above, they demonstrate using viper in its singleton style approach.

### Working with multiple vipers

You can also create many different vipers for use in your application. Each will have its own unique set of configurations and values. Each can read from a different config file, key value store, etc. All of the functions that viper package supports are mirrored as methods on a viper.

Example:

```
x := viper.New()
y := viper.New()

x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")

//...
```

When working with multiple vipers, it is up to the user to keep track of the different vipers.

## Q & A

### Why is it called “Viper”?

A: Viper is designed to be a [companion](http://en.wikipedia.org/wiki/Viper_(G.I._Joe)) to [Cobra](https://github.com/spf13/cobra). While both can operate completely independently, together they make a powerful pair to handle much of your application foundation needs.

### Why is it called “Cobra”?

Is there a better name for a [commander](http://en.wikipedia.org/wiki/Cobra_Commander)?

### Does Viper support case sensitive keys?

**tl;dr:** No.

Viper merges configuration from various sources, many of which are either case insensitive or uses different casing than the rest of the sources (eg. env vars). In order to provide the best experience when using multiple sources, the decision has been made to make all keys case insensitive.

There has been several attempts to implement case sensitivity, but unfortunately it's not that trivial. We might take a stab at implementing it in [Viper v2](https://github.com/spf13/viper/issues/772), but despite the initial noise, it does not seem to be requested that much.

You can vote for case sensitivity by filling out this feedback form: https://forms.gle/R6faU74qPRPAzchZ9

### Is it safe to concurrently read and write to a viper?

No, you will need to synchronize access to the viper yourself (for example by using the `sync` package). Concurrent reads and writes can cause a panic.

## Troubleshooting

See [TROUBLESHOOTING.md](https://github.com/spf13/viper/blob/v1.15.0/TROUBLESHOOTING.md).

Collapse ▴

## Documentation 

[Rendered for](https://go.dev/about#build-context)                   linux/amd64                   windows/amd64                   darwin/amd64                   js/wasm                

### Constants 

This section is empty.

### Variables 

[View Source](https://github.com/spf13/viper/blob/v1.15.0/viper.go#L81)

```
var RemoteConfig remoteConfigFactory
```

RemoteConfig is optional, see the remote package

[View Source](https://github.com/spf13/viper/blob/v1.15.0/viper.go#L419)

```
var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"}
```

SupportedExts are universally supported extensions.

[View Source](https://github.com/spf13/viper/blob/v1.15.0/viper.go#L422)

```
var SupportedRemoteProviders = []string{"etcd", "etcd3", "consul", "firestore"}
```

SupportedRemoteProviders are universally supported remote providers.

### Functions 

#### func AddConfigPath 

```
func AddConfigPath(in string)
```

AddConfigPath adds a path for Viper to search for the config file in. Can be called multiple times to define multiple search paths.

#### func AddRemoteProvider 

```
func AddRemoteProvider(provider, endpoint, path string) error
```

AddRemoteProvider adds a remote configuration source. Remote Providers are searched in the order they are added. provider is a string value: "etcd", "etcd3", "consul" or "firestore" are currently supported. endpoint is the url. etcd requires [http://ip:port](http://ip:port/) consul requires ip:port path is the path in the k/v store to retrieve configuration To retrieve a config file called myapp.json from /configs/myapp.json you should set path to /configs and set config name (SetConfigName()) to "myapp"

#### func AddSecureRemoteProvider 

```
func AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
```

AddSecureRemoteProvider adds a remote configuration source. Secure Remote Providers are searched in the order they are added. provider is a string value: "etcd", "etcd3", "consul" or "firestore" are currently supported. endpoint is the url. etcd requires [http://ip:port](http://ip:port/) consul requires ip:port secretkeyring is the filepath to your openpgp secret keyring. e.g. /etc/secrets/myring.gpg path is the path in the k/v store to retrieve configuration To retrieve a config file called myapp.json from /configs/myapp.json you should set path to /configs and set config name (SetConfigName()) to "myapp" Secure Remote Providers are implemented with github.com/bketelsen/crypt

#### func AllKeys 

```
func AllKeys() []string
```

AllKeys returns all keys holding a value, regardless of where they are set. Nested keys are returned with a v.keyDelim separator

#### func AllSettings 

```
func AllSettings() map[string]interface{}
```

AllSettings merges all settings and returns them as a map[string]interface{}.

#### func AllowEmptyEnv  <- v1.3.0

```
func AllowEmptyEnv(allowEmptyEnv bool)
```

AllowEmptyEnv tells Viper to consider set, but empty environment variables as valid values instead of falling back. For backward compatibility reasons this is false by default.

#### func AutomaticEnv 

```
func AutomaticEnv()
```

AutomaticEnv makes Viper check if environment variables match any of the existing keys (config, default or flags). If matching env vars are found, they are loaded into Viper.

#### func BindEnv 

```
func BindEnv(input ...string) error
```

BindEnv binds a Viper key to a ENV variable. ENV variables are case sensitive. If only a key is provided, it will use the env key matching the key, uppercased. If more arguments are provided, they will represent the env variable names that should bind to this key and will be taken in the specified order. EnvPrefix will be used when set when env name is not provided.

#### func BindFlagValue 

```
func BindFlagValue(key string, flag FlagValue) error
```

BindFlagValue binds a specific key to a FlagValue.

#### func BindFlagValues 

```
func BindFlagValues(flags FlagValueSet) error
```

BindFlagValues binds a full FlagValue set to the configuration, using each flag's long name as the config key.

#### func BindPFlag 

```
func BindPFlag(key string, flag *pflag.Flag) error
```

BindPFlag binds a specific key to a pflag (as used by cobra). Example (where serverCmd is a Cobra instance):

```
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
Viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
```

#### func BindPFlags 

```
func BindPFlags(flags *pflag.FlagSet) error
```

BindPFlags binds a full flag set to the configuration, using each flag's long name as the config key.

#### func ConfigFileUsed 

```
func ConfigFileUsed() string
```

ConfigFileUsed returns the file used to populate the config registry.

#### func Debug 

```
func Debug()
```

Debug prints all configuration registries for debugging purposes.

#### func DebugTo  <- v1.13.0

```
func DebugTo(w io.Writer)
```

#### func Get 

```
func Get(key string) interface{}
```

Get can retrieve any value given the key to use. Get is case-insensitive for a key. Get has the behavior of returning the value associated with the first place from where it is set. Viper will check in the following order: override, flag, env, config file, key/value store, default

Get returns an interface. For a specific value use one of the Get____ methods.

#### func GetBool 

```
func GetBool(key string) bool
```

GetBool returns the value associated with the key as a boolean.

#### func GetDuration 

```
func GetDuration(key string) time.Duration
```

GetDuration returns the value associated with the key as a duration.

#### func GetFloat64 

```
func GetFloat64(key string) float64
```

GetFloat64 returns the value associated with the key as a float64.

#### func GetInt 

```
func GetInt(key string) int
```

GetInt returns the value associated with the key as an integer.

#### func GetInt32  <- v1.1.0

```
func GetInt32(key string) int32
```

GetInt32 returns the value associated with the key as an integer.

#### func GetInt64 

```
func GetInt64(key string) int64
```

GetInt64 returns the value associated with the key as an integer.

#### func GetIntSlice  <- v1.5.0

```
func GetIntSlice(key string) []int
```

GetIntSlice returns the value associated with the key as a slice of int values.

#### func GetSizeInBytes 

```
func GetSizeInBytes(key string) uint
```

GetSizeInBytes returns the size of the value associated with the given key in bytes.

#### func GetString 

```
func GetString(key string) string
```

GetString returns the value associated with the key as a string.

#### func GetStringMap 

```
func GetStringMap(key string) map[string]interface{}
```

GetStringMap returns the value associated with the key as a map of interfaces.

#### func GetStringMapString 

```
func GetStringMapString(key string) map[string]string
```

GetStringMapString returns the value associated with the key as a map of strings.

#### func GetStringMapStringSlice 

```
func GetStringMapStringSlice(key string) map[string][]string
```

GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.

#### func GetStringSlice 

```
func GetStringSlice(key string) []string
```

GetStringSlice returns the value associated with the key as a slice of strings.

#### func GetTime 

```
func GetTime(key string) time.Time
```

GetTime returns the value associated with the key as time.

#### func GetUint  <- v1.4.0

```
func GetUint(key string) uint
```

GetUint returns the value associated with the key as an unsigned integer.

#### func GetUint16  <- v1.13.0

```
func GetUint16(key string) uint16
```

GetUint16 returns the value associated with the key as an unsigned integer.

#### func GetUint32  <- v1.4.0

```
func GetUint32(key string) uint32
```

GetUint32 returns the value associated with the key as an unsigned integer.

#### func GetUint64  <- v1.4.0

```
func GetUint64(key string) uint64
```

GetUint64 returns the value associated with the key as an unsigned integer.

#### func InConfig 

```
func InConfig(key string) bool
```

InConfig checks to see if the given key (or an alias) is in the config file.

#### func IsSet 

```
func IsSet(key string) bool
```

IsSet checks to see if the key has been set in any of the data locations. IsSet is case-insensitive for a key.

#### func MergeConfig 

```
func MergeConfig(in io.Reader) error
```

MergeConfig merges a new configuration with an existing config.

#### func MergeConfigMap  <- v1.3.0

```
func MergeConfigMap(cfg map[string]interface{}) error
```

MergeConfigMap merges the configuration from the map given with an existing config. Note that the map given may be modified.

#### func MergeInConfig 

```
func MergeInConfig() error
```

MergeInConfig merges a new configuration with an existing config.

#### func MustBindEnv  <- v1.12.0

```
func MustBindEnv(input ...string)
```

MustBindEnv wraps BindEnv in a panic. If there is an error binding an environment variable, MustBindEnv will panic.

#### func OnConfigChange 

```
func OnConfigChange(run func(in fsnotify.Event))
```

OnConfigChange sets the event handler that is called when a config file changes.

#### func ReadConfig 

```
func ReadConfig(in io.Reader) error
```

ReadConfig will read a configuration file, setting existing keys to nil if the key does not exist in the file.

#### func ReadInConfig 

```
func ReadInConfig() error
```

ReadInConfig will discover and load the configuration file from disk and key/value stores, searching in one of the defined paths.

#### func ReadRemoteConfig 

```
func ReadRemoteConfig() error
```

ReadRemoteConfig attempts to get configuration from a remote source and read it in the remote configuration registry.

#### func RegisterAlias 

```
func RegisterAlias(alias string, key string)
```

RegisterAlias creates an alias that provides another accessor for the same key. This enables one to change a name without breaking the application.

#### func Reset 

```
func Reset()
```

Reset is intended for testing, will reset all to default settings. In the public interface for the viper package so applications can use it in their testing as well.

#### func SafeWriteConfig  <- v1.0.1

```
func SafeWriteConfig() error
```

SafeWriteConfig writes current configuration to file only if the file does not exist.

#### func SafeWriteConfigAs  <- v1.0.1

```
func SafeWriteConfigAs(filename string) error
```

SafeWriteConfigAs writes current configuration to a given filename if it does not exist.

#### func Set 

```
func Set(key string, value interface{})
```

Set sets the value for the key in the override register. Set is case-insensitive for a key. Will be used instead of values obtained via flags, config file, ENV, default, or key/value store.

#### func SetConfigFile 

```
func SetConfigFile(in string)
```

SetConfigFile explicitly defines the path, name and extension of the config file. Viper will use this and not check any of the config paths.

#### func SetConfigName 

```
func SetConfigName(in string)
```

SetConfigName sets name for the config file. Does not include extension.

#### func SetConfigPermissions  <- v1.4.0

```
func SetConfigPermissions(perm os.FileMode)
```

SetConfigPermissions sets the permissions for the config file.

#### func SetConfigType 

```
func SetConfigType(in string)
```

SetConfigType sets the type of the configuration returned by the remote source, e.g. "json".

#### func SetDefault 

```
func SetDefault(key string, value interface{})
```

SetDefault sets the default value for this key. SetDefault is case-insensitive for a key. Default only used when no value is provided by the user via flag, config or ENV.

#### func SetEnvKeyReplacer 

```
func SetEnvKeyReplacer(r *strings.Replacer)
```

SetEnvKeyReplacer sets the strings.Replacer on the viper object Useful for mapping an environmental variable to a key that does not match it.

#### func SetEnvPrefix 

```
func SetEnvPrefix(in string)
```

SetEnvPrefix defines a prefix that ENVIRONMENT variables will use. E.g. if your prefix is "spf", the env registry will look for env variables that start with "SPF_".

#### func SetFs 

```
func SetFs(fs afero.Fs)
```

SetFs sets the filesystem to use to read configuration.

#### func SetTypeByDefaultValue 

```
func SetTypeByDefaultValue(enable bool)
```

SetTypeByDefaultValue enables or disables the inference of a key value's type when the Get function is used based upon a key's default value as opposed to the value returned based on the normal fetch logic.

For example, if a key has a default value of []string{} and the same key is set via an environment variable to "a b c", a call to the Get function would return a string slice for the key if the key's type is inferred by the default value and the Get function would return:

```
[]string {"a", "b", "c"}
```

Otherwise the Get function would return:

```
"a b c"
```

#### func Unmarshal 

```
func Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error
```

Unmarshal unmarshals the config into a Struct. Make sure that the tags on the fields of the structure are properly set.

#### func UnmarshalExact  <- v1.6.0

```
func UnmarshalExact(rawVal interface{}, opts ...DecoderConfigOption) error
```

UnmarshalExact unmarshals the config into a Struct, erroring if a field is nonexistent in the destination struct.

#### func UnmarshalKey 

```
func UnmarshalKey(key string, rawVal interface{}, opts ...DecoderConfigOption) error
```

UnmarshalKey takes a single key and unmarshals it into a Struct.

#### func WatchConfig 

```
func WatchConfig()
```

WatchConfig starts watching a config file for changes.

#### func WatchRemoteConfig 

```
func WatchRemoteConfig() error
```

#### func WriteConfig  <- v1.0.1

```
func WriteConfig() error
```

WriteConfig writes the current configuration to a file.

#### func WriteConfigAs  <- v1.0.1

```
func WriteConfigAs(filename string) error
```

WriteConfigAs writes current configuration to a given filename.

### Types 

#### type ConfigFileAlreadyExistsError  <- v1.6.0

```
type ConfigFileAlreadyExistsError string
```

ConfigFileAlreadyExistsError denotes failure to write new configuration file.

#### func (ConfigFileAlreadyExistsError) [Error](https://github.com/spf13/viper/blob/v1.15.0/viper.go#L124)  <- v1.6.0

```
func (faee ConfigFileAlreadyExistsError) Error() string
```

Error returns the formatted error when configuration already exists.

#### type ConfigFileNotFoundError 

```
type ConfigFileNotFoundError struct {
	// contains filtered or unexported fields
}
```

ConfigFileNotFoundError denotes failing to find configuration file.

#### func (ConfigFileNotFoundError) [Error](https://github.com/spf13/viper/blob/v1.15.0/viper.go#L116) 

```
func (fnfe ConfigFileNotFoundError) Error() string
```

Error returns the formatted configuration error.

#### type ConfigMarshalError  <- v1.0.1

```
type ConfigMarshalError struct {
	// contains filtered or unexported fields
}
```

ConfigMarshalError happens when failing to marshal the configuration.

#### func (ConfigMarshalError) [Error](https://github.com/spf13/viper/blob/v1.15.0/viper.go#L59)  <- v1.0.1

```
func (e ConfigMarshalError) Error() string
```

Error returns the formatted configuration error.

#### type ConfigParseError 

```
type ConfigParseError struct {
	// contains filtered or unexported fields
}
```

ConfigParseError denotes failing to parse configuration file.

#### func (ConfigParseError) [Error](https://github.com/spf13/viper/blob/v1.15.0/util.go#L30) 

```
func (pe ConfigParseError) Error() string
```

Error returns the formatted configuration error.

#### type DecoderConfigOption  <- v1.1.0

```
type DecoderConfigOption func(*mapstructure.DecoderConfig)
```

A DecoderConfigOption can be passed to viper.Unmarshal to configure mapstructure.DecoderConfig options

#### func DecodeHook  <- v1.1.0

```
func DecodeHook(hook mapstructure.DecodeHookFunc) DecoderConfigOption
```

DecodeHook returns a DecoderConfigOption which overrides the default DecoderConfig.DecodeHook value, the default is:

```
 mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	)
```

#### type FlagValue 

```
type FlagValue interface {
	HasChanged() bool
	Name() string
	ValueString() string
	ValueType() string
}
```

FlagValue is an interface that users can implement to bind different flags to viper.

#### type FlagValueSet 

```
type FlagValueSet interface {
	VisitAll(fn func(FlagValue))
}
```

FlagValueSet is an interface that users can implement to bind a set of flags to viper.

#### type Logger  <- v1.10.0

```
type Logger interface {
	// Trace logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	Trace(msg string, keyvals ...interface{})

	// Debug logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	Debug(msg string, keyvals ...interface{})

	// Info logs an Info event.
	//
	// General information about what's happening inside the system.
	Info(msg string, keyvals ...interface{})

	// Warn logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	Warn(msg string, keyvals ...interface{})

	// Error logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	Error(msg string, keyvals ...interface{})
}
```

Logger is a unified interface for various logging use cases and practices, including:

- leveled logging
- structured logging

#### type Option  <- v1.6.0

```
type Option interface {
	// contains filtered or unexported methods
}
```

Option configures Viper using the functional options paradigm popularized by Rob Pike and Dave Cheney. If you're unfamiliar with this style, see https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html and https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis.

#### func EnvKeyReplacer  <- v1.6.0

```
func EnvKeyReplacer(r StringReplacer) Option
```

EnvKeyReplacer sets a replacer used for mapping environment variables to internal keys.

#### func IniLoadOptions  <- v1.8.0

```
func IniLoadOptions(in ini.LoadOptions) Option
```

IniLoadOptions sets the load options for ini parsing.

#### func KeyDelimiter  <- v1.6.0

```
func KeyDelimiter(d string) Option
```

KeyDelimiter sets the delimiter used for determining key parts. By default it's value is ".".

#### type RemoteConfigError 

```
type RemoteConfigError string
```

RemoteConfigError denotes encountering an error while trying to pull the configuration from the remote provider.

#### func (RemoteConfigError) [Error](https://github.com/spf13/viper/blob/v1.15.0/viper.go#L106) 

```
func (rce RemoteConfigError) Error() string
```

Error returns the formatted remote provider error

#### type RemoteProvider 

```
type RemoteProvider interface {
	Provider() string
	Endpoint() string
	Path() string
	SecretKeyring() string
}
```

RemoteProvider stores the configuration necessary to connect to a remote key/value store. Optional secretKeyring to unencrypt encrypted values can be provided.

#### type RemoteResponse 

```
type RemoteResponse struct {
	Value []byte
	Error error
}
```

#### type StringReplacer  <- v1.6.0

```
type StringReplacer interface {
	// Replace returns a copy of s with all replacements performed.
	Replace(s string) string
}
```

StringReplacer applies a set of replacements to a string.

#### type UnsupportedConfigError 

```
type UnsupportedConfigError string
```

UnsupportedConfigError denotes encountering an unsupported configuration filetype.

#### func (UnsupportedConfigError) [Error](https://github.com/spf13/viper/blob/v1.15.0/viper.go#L88) 

```
func (str UnsupportedConfigError) Error() string
```

Error returns the formatted configuration error.

#### type UnsupportedRemoteProviderError 

```
type UnsupportedRemoteProviderError string
```

UnsupportedRemoteProviderError denotes encountering an unsupported remote provider. Currently only etcd and Consul are supported.

#### func (UnsupportedRemoteProviderError) [Error](https://github.com/spf13/viper/blob/v1.15.0/viper.go#L97) 

```
func (str UnsupportedRemoteProviderError) Error() string
```

Error returns the formatted remote provider error.

#### type Viper 

```
type Viper struct {
	// contains filtered or unexported fields
}
```

Viper is a prioritized configuration registry. It maintains a set of configuration sources, fetches values to populate those, and provides them according to the source's priority. The priority of the sources is the following: 1. overrides 2. flags 3. env. variables 4. config file 5. key/value store 6. defaults

For example, if values from the following sources were loaded:

```
Defaults : {
	"secret": "",
	"user": "default",
	"endpoint": "https://localhost"
}
Config : {
	"user": "root"
	"secret": "defaultsecret"
}
Env : {
	"secret": "somesecretkey"
}
```

The resulting config will have the following values:

```
{
	"secret": "somesecretkey",
	"user": "root",
	"endpoint": "https://localhost"
}
```

Note: Vipers are not safe for concurrent Get() and Set() operations.

#### func GetViper 

```
func GetViper() *Viper
```

GetViper gets the global Viper instance.

#### func New 

```
func New() *Viper
```

New returns an initialized Viper instance.

#### func NewWithOptions  <- v1.6.0

```
func NewWithOptions(opts ...Option) *Viper
```

NewWithOptions creates a new Viper instance.

#### func Sub 

```
func Sub(key string) *Viper
```

Sub returns new Viper instance representing a sub tree of this instance. Sub is case-insensitive for a key.

#### func (*Viper) AddConfigPath 

```
func (v *Viper) AddConfigPath(in string)
```

#### func (*Viper) AddRemoteProvider 

```
func (v *Viper) AddRemoteProvider(provider, endpoint, path string) error
```

#### func (*Viper) AddSecureRemoteProvider 

```
func (v *Viper) AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
```

#### func (*Viper) AllKeys 

```
func (v *Viper) AllKeys() []string
```

#### func (*Viper) AllSettings 

```
func (v *Viper) AllSettings() map[string]interface{}
```

#### func (*Viper) AllowEmptyEnv  <- v1.3.0

```
func (v *Viper) AllowEmptyEnv(allowEmptyEnv bool)
```

#### func (*Viper) AutomaticEnv 

```
func (v *Viper) AutomaticEnv()
```

#### func (*Viper) BindEnv 

```
func (v *Viper) BindEnv(input ...string) error
```

#### func (*Viper) BindFlagValue 

```
func (v *Viper) BindFlagValue(key string, flag FlagValue) error
```

#### func (*Viper) BindFlagValues 

```
func (v *Viper) BindFlagValues(flags FlagValueSet) (err error)
```

#### func (*Viper) BindPFlag 

```
func (v *Viper) BindPFlag(key string, flag *pflag.Flag) error
```

#### func (*Viper) BindPFlags 

```
func (v *Viper) BindPFlags(flags *pflag.FlagSet) error
```

#### func (*Viper) ConfigFileUsed 

```
func (v *Viper) ConfigFileUsed() string
```

#### func (*Viper) Debug 

```
func (v *Viper) Debug()
```

#### func (*Viper) DebugTo  <- v1.13.0

```
func (v *Viper) DebugTo(w io.Writer)
```

#### func (*Viper) Get 

```
func (v *Viper) Get(key string) interface{}
```

#### func (*Viper) GetBool 

```
func (v *Viper) GetBool(key string) bool
```

#### func (*Viper) GetDuration 

```
func (v *Viper) GetDuration(key string) time.Duration
```

#### func (*Viper) GetFloat64 

```
func (v *Viper) GetFloat64(key string) float64
```

#### func (*Viper) GetInt 

```
func (v *Viper) GetInt(key string) int
```

#### func (*Viper) GetInt32  <- v1.1.0

```
func (v *Viper) GetInt32(key string) int32
```

#### func (*Viper) GetInt64 

```
func (v *Viper) GetInt64(key string) int64
```

#### func (*Viper) GetIntSlice  <- v1.5.0

```
func (v *Viper) GetIntSlice(key string) []int
```

#### func (*Viper) GetSizeInBytes 

```
func (v *Viper) GetSizeInBytes(key string) uint
```

#### func (*Viper) GetString 

```
func (v *Viper) GetString(key string) string
```

#### func (*Viper) GetStringMap 

```
func (v *Viper) GetStringMap(key string) map[string]interface{}
```

#### func (*Viper) GetStringMapString 

```
func (v *Viper) GetStringMapString(key string) map[string]string
```

#### func (*Viper) GetStringMapStringSlice 

```
func (v *Viper) GetStringMapStringSlice(key string) map[string][]string
```

#### func (*Viper) GetStringSlice 

```
func (v *Viper) GetStringSlice(key string) []string
```

#### func (*Viper) GetTime 

```
func (v *Viper) GetTime(key string) time.Time
```

#### func (*Viper) GetUint  <- v1.4.0

```
func (v *Viper) GetUint(key string) uint
```

#### func (*Viper) GetUint16  <- v1.13.0

```
func (v *Viper) GetUint16(key string) uint16
```

#### func (*Viper) GetUint32  <- v1.4.0

```
func (v *Viper) GetUint32(key string) uint32
```

#### func (*Viper) GetUint64  <- v1.4.0

```
func (v *Viper) GetUint64(key string) uint64
```

#### func (*Viper) InConfig 

```
func (v *Viper) InConfig(key string) bool
```

#### func (*Viper) IsSet 

```
func (v *Viper) IsSet(key string) bool
```

#### func (*Viper) MergeConfig 

```
func (v *Viper) MergeConfig(in io.Reader) error
```

#### func (*Viper) MergeConfigMap  <- v1.3.0

```
func (v *Viper) MergeConfigMap(cfg map[string]interface{}) error
```

#### func (*Viper) MergeInConfig 

```
func (v *Viper) MergeInConfig() error
```

#### func (*Viper) MustBindEnv  <- v1.12.0

```
func (v *Viper) MustBindEnv(input ...string)
```

#### func (*Viper) OnConfigChange 

```
func (v *Viper) OnConfigChange(run func(in fsnotify.Event))
```

OnConfigChange sets the event handler that is called when a config file changes.

#### func (*Viper) ReadConfig 

```
func (v *Viper) ReadConfig(in io.Reader) error
```

#### func (*Viper) ReadInConfig 

```
func (v *Viper) ReadInConfig() error
```

#### func (*Viper) ReadRemoteConfig 

```
func (v *Viper) ReadRemoteConfig() error
```

#### func (*Viper) RegisterAlias 

```
func (v *Viper) RegisterAlias(alias string, key string)
```

#### func (*Viper) SafeWriteConfig  <- v1.0.1

```
func (v *Viper) SafeWriteConfig() error
```

#### func (*Viper) SafeWriteConfigAs  <- v1.0.1

```
func (v *Viper) SafeWriteConfigAs(filename string) error
```

#### func (*Viper) Set 

```
func (v *Viper) Set(key string, value interface{})
```

#### func (*Viper) SetConfigFile 

```
func (v *Viper) SetConfigFile(in string)
```

#### func (*Viper) SetConfigName 

```
func (v *Viper) SetConfigName(in string)
```

#### func (*Viper) SetConfigPermissions  <- v1.4.0

```
func (v *Viper) SetConfigPermissions(perm os.FileMode)
```

#### func (*Viper) SetConfigType 

```
func (v *Viper) SetConfigType(in string)
```

#### func (*Viper) SetDefault 

```
func (v *Viper) SetDefault(key string, value interface{})
```

#### func (*Viper) SetEnvKeyReplacer 

```
func (v *Viper) SetEnvKeyReplacer(r *strings.Replacer)
```

#### func (*Viper) SetEnvPrefix 

```
func (v *Viper) SetEnvPrefix(in string)
```

#### func (*Viper) SetFs 

```
func (v *Viper) SetFs(fs afero.Fs)
```

#### func (*Viper) SetTypeByDefaultValue 

```
func (v *Viper) SetTypeByDefaultValue(enable bool)
```

#### func (*Viper) Sub 

```
func (v *Viper) Sub(key string) *Viper
```

#### func (*Viper) Unmarshal 

```
func (v *Viper) Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error
```

#### func (*Viper) UnmarshalExact 

```
func (v *Viper) UnmarshalExact(rawVal interface{}, opts ...DecoderConfigOption) error
```

#### func (*Viper) UnmarshalKey 

```
func (v *Viper) UnmarshalKey(key string, rawVal interface{}, opts ...DecoderConfigOption) error
```

#### func (*Viper) WatchConfig 

```
func (v *Viper) WatchConfig()
```

WatchConfig starts watching a config file for changes.

#### func (*Viper) WatchRemoteConfig 

```
func (v *Viper) WatchRemoteConfig() error
```

#### func (*Viper) WatchRemoteConfigOnChannel 

```
func (v *Viper) WatchRemoteConfigOnChannel() error
```

#### func (*Viper) WriteConfig  <- v1.0.1

```
func (v *Viper) WriteConfig() error
```

#### func (*Viper) WriteConfigAs  <- v1.0.1

```
func (v *Viper) WriteConfigAs(filename string) error
```
