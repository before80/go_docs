+++
title = "viper"
type = "docs"
date = 2023-05-22T08:46:24+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# viper

> 原文：[https://pkg.go.dev/github.com/spf13/viper](https://pkg.go.dev/github.com/spf13/viper)
>
> 版本：v1.16.0
>
> 发布日期：2023.5.30
>
> github网址：[https://github.com/spf13/viper](https://github.com/spf13/viper)



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

​	可以将 Viper 视为您的应用程序所有配置需求的注册表。

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

## 向 Viper 中添加值

### 设置默认值

​	一个良好的配置系统将支持默认值。对于一个键来说，默认值不是必需的，但在没有通过配置文件、环境变量、远程配置或命令行标志设置键时非常有用。

示例：

```go
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
```

### 读取配置文件

​	Viper需要进行最小配置，以便知道在哪里查找配置文件。Viper 支持 JSON、TOML、YAML、HCL、INI、envfile 和 Java Properties 文件。Viper 可以搜索多个路径，但是当前每个 Viper 实例只支持一个配置文件。Viper 不会默认设置任何配置文件搜索路径，而是将默认的决策留给应用程序。

​	以下是如何使用 Viper 搜索并读取配置文件的示例。具体路径都不是必需的，但至少应提供一个期望找到配置文件的路径。

```go
viper.SetConfigName("config") //配置文件的名称（无扩展名）
viper.SetConfigType("yaml") // 如果配置文件名称中不包含扩展名，则必需(REQUIRED)
viper.AddConfigPath("/etc/appname/") //配置文件的搜索路径
viper.AddConfigPath("$HOME/.appname") //可多次调用以添加多个搜索路径
viper.AddConfigPath(".") // 可选地在工作目录中查找配置文件
err := viper.ReadInConfig() // 查找并读取配置文件
if err != nil { // 处理读取配置文件时的错误
	panic(fmt.Errorf("fatal error config file: %w", err))
}
```

​	您可以像这样处理未找到配置文件的特殊情况：

```go
if err := viper.ReadInConfig(); err != nil {
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// 未找到配置文件；如果需要，可以忽略错误
	} else {
		// 找到配置文件，但产生了其他错误
	}
}

// 找到配置文件并且已成功解析
```

> *注意 [自版本 1.6 起]：* 您还可以拥有没有扩展名的文件，并以编程方式指定格式。对于那些位于用户主目录下且没有任何扩展名的配置文件，例如 `.bashrc`。



### 写入配置文件

​	从配置文件中读取很有用，但有时您希望存储运行时所做的所有修改。为此，有一系列可用的命令，每个命令都有其特定的目的：

- WriteConfig： 将当前的 Viper 配置写入预定义的路径（如果存在）。如果没有预定义的路径，则报错。如果配置文件已存在，则**会**覆盖当前的配置文件。
- SafeWriteConfig：将当前的 Viper 配置写入预定义的路径。如果没有预定义的路径，则报错。如果配置文件已存在，则**不会**覆盖当前的配置文件。
- WriteConfigAs：将当前的 Viper 配置写入指定的文件路径。如果给定的文件已存在，则**会**覆盖该文件。
- SafeWriteConfigAs：将当前的 Viper 配置写入指定的文件路径。如果给定的文件已存在，则**不会**覆盖该文件。

​	根据经验，标记为 safe 的所有操作都不会覆盖任何文件，只会在文件不存在时创建，而默认行为是创建或截断（truncate）文件。

​	以下是一个小的示例部分：

```go
viper.WriteConfig() // 将当前配置写入由 'viper.AddConfigPath()' 和 'viper.SetConfigName' 设置的预定义路径
viper.SafeWriteConfig()
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") // 会报错，因为已经写入过
viper.SafeWriteConfigAs("/path/to/my/.other_config")
```

### 监听和重新读取配置文件 

​	Viper支持在应用程序运行时实时读取配置文件的能力。

​	不再需要重新启动服务器以使配置生效，使用Viper的应用程序可以在运行过程中读取配置文件的更新，而无需中断。

​	只需告诉 Viper 实例执行 `WatchConfig()`。可选地，您还可以为 Viper 提供一个在每次更改发生时运行的函数。

> ​	**确保在调用 `WatchConfig()` 之前添加所有的配置路径。**
>

```go
viper.OnConfigChange(func(e fsnotify.Event) {
	fmt.Println("Config file changed:", e.Name)
})

viper.WatchConfig()
```

### 从 io.Reader 读取配置 

​	Viper 预定义了许多配置源，如文件、环境变量、标志和远程键值存储，但您并不受限于此。您还可以实现自己所需的配置源，并将其提供给 Viper。

```go
viper.SetConfigType("yaml") // 或 viper.SetConfigType("YAML")

// 通过任何方式将此配置内容引入你的程序中。
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

viper.Get("name") // 这将返回 "steve"
```

### 设置覆盖值

​	这些可以来自命令行标志，或者来自您自己的应用程序逻辑。

```go
viper.Set("Verbose", true)
viper.Set("LogFile", LogFile)
```

### 注册和使用别名

​	别名允许多个键引用同一个值。

```go
viper.RegisterAlias("loud", "Verbose")

viper.Set("verbose", true) // 与下一行代码的结果相同
viper.Set("loud", true)   // 与上一行代码的结果相同

viper.GetBool("loud") // 返回 true
viper.GetBool("verbose") // 返回 true
```

### 使用环境变量

​	Viper完全支持环境变量。这使得12因子应用程序能够立即使用。有五个方法可用于处理环境变量：

- `AutomaticEnv()`
- `BindEnv(string...) : error`
- `SetEnvPrefix(string)`
- `SetEnvKeyReplacer(string...) *strings.Replacer`
- `AllowEmptyEnv(bool)`

> ​	在处理环境变量时，重要的是要认识到Viper将环境变量视为区分大小写。

​	Viper提供了一种机制来确保ENV变量的唯一性。通过使用`SetEnvPrefix`，您可以告诉Viper在读取环境变量时使用前缀。`BindEnv`和`AutomaticEnv`都将使用此前缀。

​	`BindEnv`接受一个或多个参数。第一个参数是键名，其余参数是要绑定到此键的环境变量的名称。如果提供了多个参数，它们将按指定的顺序优先。环境变量的名称区分大小写。如果未提供ENV变量名，则Viper将自动假设ENV变量与以下格式匹配：前缀 + "_" +键名（全部大写）。当您显式提供ENV变量名（第二个参数）时，它**不会**自动添加前缀。例如，如果第二个参数是"id"，Viper将查找ENV变量 "ID"。

​	在处理环境变量时，需要注意的一点是每次访问时都会读取其值。在调用`BindEnv`时，Viper不会固定该值。

​	`AutomaticEnv`是一个强大的辅助功能，特别是与`SetEnvPrefix`结合使用时。当（`AutomaticEnv`）被调用时，每次进行`viper.Get`请求时，Viper都会检查是否存在相应的环境变量。它将应用以下规则：如果设置了`EnvPrefix`，它将检查是否存在一个与键名相匹配的以大写形式和前缀（如果设置）作为前缀的环境变量名称。

​	`SetEnvKeyReplacer`允许您使用`strings.Replacer`对象来在一定程度上重写Env键。这在您希望在`Get()`调用中使用`-`或其他字符，但希望您的环境变量使用`_`作为分隔符时非常有用。在`viper_test.go`中可以找到使用它的示例。

​	或者，您可以使用带有`NewWithOptions`工厂函数的`EnvKeyReplacer`。与`SetEnvKeyReplacer`不同，它接受`StringReplacer`接口，允许您编写自定义字符串替换逻辑。

​	默认情况下，空环境变量被视为未设置，并将回退到下一个配置源。要将空环境变量视为已设置，请使用`AllowEmptyEnv`方法。

#### 环境变量示例

```go
SetEnvPrefix("spf") // 将自动转为大写
BindEnv("id")

os.Setenv("SPF_ID", "13") // 通常在应用程序外部进行

id := Get("id") // 13
```

### 使用标志（Flags）

​	Viper具有与标志（flags）绑定的功能。具体而言，Viper支持[Cobra]({{< ref "/thirdPkg/Cobra">}})库中使用的`Pflags`。

​	与`BindEnv`类似，当调用绑定方法时，并不会立即设置值，而是在访问时才设置。这意味着您可以在任何时候进行绑定，甚至在`init()`函数中进行绑定。

​	对于单个标志，`BindPFlag()`方法提供了此功能。

示例：

```go
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
```

​	您还可以绑定一组现有的pflags（pflag.FlagSet）：

示例：

```go
pflag.Int("flagname", 1234, "help message for flagname")

pflag.Parse()
viper.BindPFlags(pflag.CommandLine)

i := viper.GetInt("flagname") // 从viper中检索值，而不是从pflag中
```

​	在Viper中使用[pflag](https://github.com/spf13/pflag/)并不妨碍使用标准库中的[flag](https://golang.org/pkg/flag/)包。pflag包可以通过导入这些标志来处理使用flag包定义的标志。为了实现这一点，pflag包提供了一个方便的函数`AddGoFlagSet()`。

示例：

```go
package main

import (
	"flag"
	"github.com/spf13/pflag"
)

func main() {

	// 使用标准库中的 "flag" 包
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") // 从viper中检索值

	// ...
}
```

#### Flag 接口

​	Viper提供了两个Go接口，用于绑定其他标志系统（如果您不使用`Pflags`）。

​	`FlagValue`表示一个单独的标志。以下是一个实现该接口的简单示例：

``` go
type myFlag struct {}

func (f myFlag) HasChanged() bool { 
    return false 
}

func (f myFlag) Name() string { 
    return "my-flag-name" 
}

func (f myFlag) ValueString() string { 
    return "my-flag-value" 
}

func (f myFlag) ValueType() string { 
    return "string" 
}
```

​	一旦您的标志实现了该接口，您可以简单地告诉Viper绑定它：

```go
viper.BindFlagValue("my-flag-name", myFlag{})
```

​	`FlagValueSet`表示一组标志。以下是一个实现该接口的简单示例：

``` go
type myFlagSet struct {
	flags []myFlag
}

func (f myFlagSet) VisitAll(fn func(FlagValue)) {
	for _, flag := range flags {
		fn(flag)
	}
}
```

​	一旦您的标志集实现了该接口，您可以简单地告诉Viper绑定它：

```go
fSet := myFlagSet{
	flags: []myFlag{myFlag{}, myFlag{}},
}

viper.BindFlagValues("my-flags", fSet)
```

### 远程键/值存储支持

​	要在Viper中启用远程支持，请使用 `viper/remote` 包进行空白导入：

```
import _ "github.com/spf13/viper/remote"
```

​	Viper将从键/值存储（如etcd或Consul）中检索到的路径中读取配置字符串（as JSON、TOML、YAML、HCL或envfile格式）。这些值优先于默认值，但会被从磁盘、命令行标志或环境变量中检索到的配置值所覆盖。

​	Viper 使用 [crypt](https://github.com/bketelsen/crypt) 从键/值存储中检索配置，这意味着您可以将配置值加密存储，并在具有正确的 gpg  密钥环时自动解密它们。加密是可选的。

​	您可以将远程配置与本地配置一起使用，也可以独立使用。

​	`crypt`有一个命令行助手，可以用来将配置放入键/值存储中。`crypt`默认使用位于http://127.0.0.1:4001的etcd。

```sh
$ go get github.com/bketelsen/crypt/bin/crypt
$ crypt set -plaintext /config/hugo.json /Users/hugo/settings/config.json
```

​	确认您的值是否已设置：

```sh
$ crypt get -plaintext /config/hugo.json
```

​	有关如何设置加密值或使用Consul的示例，请参阅[crypt](https://bketelsen.github.io/crypt/)文档。

### 远程键/值存储示例 - 未加密

#### etcd

```go
viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，支持的扩展名有 "json"、"toml"、"yaml"、"yml"、"properties"、"props"、"prop"、"env"、"dotenv"
err := viper.ReadRemoteConfig()
```

#### etcd3

```go
viper.AddRemoteProvider("etcd3", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，支持的扩展名有 "json"、"toml"、"yaml"、"yml"、"properties"、"props"、"prop"、"env"、"dotenv"
err := viper.ReadRemoteConfig()
```

#### Consul

You need to set a key to Consul key/value storage with JSON value containing your desired config. For example, create a Consul key/value store key `MY_CONSUL_KEY` with value:

​	您需要将一个包含所需配置的JSON值的键设置到Consul键/值存储中。例如，使用值创建Consul键/值存储键`MY_CONSUL_KEY`：

​	您需要在Consul键值存储中设置一个键，其对应的JSON值包含所需的配置。例如，创建一个Consul键值存储键`MY_CONSUL_KEY`，其值为：

```json
{
    "port": 8080,
    "hostname": "myhostname.com"
}
```

```go
viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")

viper.SetConfigType("json") // 需要显式设置为json
err := viper.ReadRemoteConfig()

fmt.Println(viper.Get("port")) // 8080
fmt.Println(viper.Get("hostname")) // myhostname.com
```

#### Firestore

```go
viper.AddRemoteProvider("firestore", "google-cloud-project-id", "collection/document")

viper.SetConfigType("json") // 配置的格式: "json"、"toml"、"yaml"、"yml"

err := viper.ReadRemoteConfig()
```

​	当然，您也可以使用`SecureRemoteProvider`。

### 远程键/值存储示例 - 加密

```go
viper.AddSecureRemoteProvider("etcd","http://127.0.0.1:4001","/config/hugo.json","/etc/secrets/mykeyring.gpg")

viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，支持的扩展名有 "json"、"toml"、"yaml"、"yml"、"properties"、"props"、"prop"、"env"、"dotenv"

err := viper.ReadRemoteConfig()
```

### 监听etcd中的更改 - 未加密

```go
// 或者，您可以创建一个新的viper实例。
var runtime_viper = viper.New()

runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
runtime_viper.SetConfigType("yaml") // 因为在字节流中没有文件扩展名，支持的扩展名有 "json"、"toml"、"yaml"、"yml"、"properties"、"props"、"prop"、"env"、"dotenv"

// 第一次从远程配置中读取。
err := runtime_viper.ReadRemoteConfig()

// 反序列化配置
runtime_viper.Unmarshal(&runtime_conf)

// 打开一个goroutine无限期地监听远程更改
go func(){
	for {
		time.Sleep(time.Second * 5) // 每个请求之后的延迟

		// 目前仅已测试etcd支持
		err := runtime_viper.WatchRemoteConfig()
		if err != nil {
			log.Errorf("unable to read remote config: %v", err)
			continue
		}

         // 反序列化新配置到我们的运行时配置结构。您还可以使用通道实现一个信号，以通知系统发生了更改
		runtime_viper.Unmarshal(&runtime_conf)
	}
}()
```

## 从Viper获取值

​	在Viper中，根据值的类型，有几种获取值的方法和函数可用： 

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

​	需要注意的一点是，如果找不到给定的键，每个Get函数都会返回一个零值。要检查给定键是否存在，可以使用`IsSet()`方法。

示例：

```go
viper.GetString("logfile") // 不区分大小写的设置和获取
if viper.GetBool("verbose") {
	fmt.Println("verbose enabled")
}
```

### 访问嵌套键

​	accessor（访问器）方法还接受格式化路径来获取深层嵌套的键。例如，如果加载了以下 JSON 文件：

```json
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

​	Viper可以通过传递一个由`.`分隔的键路径来访问嵌套字段：

```go
GetString("datastore.metric.host") // 返回 "127.0.0.1"
```

​	这遵循了上述建立的优先级规则；在剩余的配置注册表中，路径的搜索将会逐级进行，直到找到为止。

​	例如，给定这个配置文件，`datastore.metric.host`和`datastore.metric.port`已经定义（并且可能被覆盖）。如果另外还在默认配置中定义了`datastore.metric.protocol`，Viper也会找到它。

​	然而，如果`datastore.metric`被立即值（immediate value）（通过标志、环境变量、`Set()`方法等）覆盖，那么`datastore.metric`的所有子键都变为未定义，它们被高优先级的配置级别"遮蔽（shadowed）"。

​	Viper可以通过在路径中使用数字来访问数组索引。例如：

```json
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
```

```go
GetInt("host.ports.1") // 返回 6029
```

​	最后，如果存在与分隔键路径匹配的键，则返回该键的值。例如：

```json
{
    "datastore.metric.host": "0.0.0.0", // <----
    
    
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

```go
GetString("datastore.metric.host") // 返回 "0.0.0.0"
```



### 提取子树（sub-tree）

​	在开发可重用模块时，将配置的子集提取出来并传递给一个模块通常很有用。这样，可以实例化多个模块，并为每个模块使用不同的配置。

​	例如，一个应用程序可能针对不同的用途使用多个不同的缓存存储：

```yaml
cache:
  cache1:
    max-items: 100
    item-size: 64
  cache2:
    max-items: 200
    item-size: 80
```

​	我们可以将缓存名称传递给一个模块（例如`NewCache("cache1")`），但这将需要奇怪的拼接来访问配置键，并且与全局配置的分离性较差。

​	因此，我们不这样做，而是将表示配置子集的Viper实例传递给构造函数：

```go
cache1Config := viper.Sub("cache.cache1")
if cache1Config == nil { // 如果找不到键，则 Sub 方法将返回 nil。
	panic("cache configuration not found")
}

cache1 := NewCache(cache1Config)
```

> **注意：**始终检查`Sub`的返回值。如果找不到键，它将返回`nil`。

​	在内部，`NewCache`函数可以直接访问`max-items`和`item-size`键。

``` go
func NewCache(v *Viper) *Cache {
	return &Cache{
		MaxItems: v.GetInt("max-items"),
		ItemSize: v.GetInt("item-size"),
	}
}
```

​	由于与主配置结构解耦，生成的代码易于测试，并且（由于同样的原因）更易于重用。

### 反序列化

​	您还可以选择将全部或特定值反序列化为结构体、映射等。

​	有两种方法可以实现这一点：

- `Unmarshal(rawVal interface{}) : error`
- `UnmarshalKey(key string, rawVal interface{}) : error`

示例：

``` go
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

​	如果要对包含点符号（默认键分隔符）的键进行配置反序列化，您需要更改分隔符：

```go
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

​	Viper还支持将值反序列化到嵌套结构体中：

```go
/*
示例配置：

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

// moduleConfig 可以在特定于模块的包中
type moduleConfig struct {
	Token string
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("unable to decode into struct, %v", err)
}
```

​	Viper在内部使用[github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure)进行值的反序列化，它默认使用`mapstructure`标签。

### 解码自定义格式

​	Viper经常被要求添加更多的值格式和解码器。例如，将以字符（点号、逗号、分号等）分隔的字符串解析为切片。

​	在Viper中，可以使用`mapstructure`解码钩子来实现这一点。

​	详细信息请参阅[此博客文章](https://sagikazarmark.hu/blog/decoding-custom-formats-with-viper/)。

### 序列化为字符串

​	您可能需要将Viper中保存的所有设置序列化为字符串，而不是将它们写入文件。您可以使用您喜欢的格式的编组器（marshaller）带上`AllSettings()`返回的配置。

```go
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

## Viper还是Vipers？

​	Viper可以直接使用，无需任何配置或初始化即可开始使用。由于大多数应用程序希望使用单个中央存储库来管理其配置，因此viper包提供了这样的功能。它类似于单例模式（singleton）。

​	在上面的所有示例中，它们演示了如何以单例模式的方式使用viper。

#### 使用多个viper

​	您还可以创建多个不同的viper实例，以供应用程序使用。每个viper实例都有自己独特的配置和数值集合。每个实例可以从不同的配置文件、键值存储等读取配置。viper包支持的所有函数也作为viper实例的方法进行了镜像。

示例：

```go
x := viper.New()
y := viper.New()

x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")

//...
```

​	在使用多个viper实例时，用户需要自行跟踪管理这些不同的viper实例。

## Q & A

### 为什么它被称为"Viper"？

A: Viper被设计为[Cobra]({{< ref "/thirdPkg/Cobra">}})的[companion（伴侣）](http://en.wikipedia.org/wiki/Viper_(G.I._Joe))。虽然它们可以完全独立运作，但是它们一起组合起来可以强大地满足你的应用程序基础需求。

### 为什么它被称为"Cobra"？

​	还有比一个[commander](http://en.wikipedia.org/wiki/Cobra_Commander)更好的名字吗？

### Viper支持区分大小写的键吗？

**tl;dr:** No.

**简短回答：** 不支持。

​	Viper从多个源中合并配置，其中许多源要么不区分大小写，要么使用与其他源（例如环境变量）不同的大小写。为了在使用多个源时提供最佳体验，决定将所有键都设置为不区分大小写。

​	已经有几次尝试实现区分大小写的功能，但不幸的是，它并不那么简单。我们可能会在[Viper v2](https://github.com/spf13/viper/issues/772)中尝试实现它，但尽管最初有些声音，但似乎并没有那么多人提出这个要求。

​	您可以通过填写这个反馈表格来表达对区分大小写功能的需求：[https://forms.gle/R6faU74qPRPAzchZ9](https://forms.gle/R6faU74qPRPAzchZ9)

### 并发读写Viper是否安全？

​	不安全，您需要自己同步对Viper的访问（例如使用`sync`包）。并发的读写操作可能会导致恐慌（panic）。

## 故障排除

​	请参阅 [TROUBLESHOOTING.md](https://github.com/spf13/viper/blob/v1.16.0/TROUBLESHOOTING.md)。



## 文档

[Rendered for](https://go.dev/about#build-context)  linux/amd64             

### 常量

This section is empty.

### 变量

[View Source](https://github.com/spf13/viper/blob/v1.16.0/viper.go#L80)

```
var RemoteConfig remoteConfigFactory
```

RemoteConfig is optional, see the remote package

[View Source](https://github.com/spf13/viper/blob/v1.16.0/viper.go#L420)

```
var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"}
```

SupportedExts are universally supported extensions.

[View Source](https://github.com/spf13/viper/blob/v1.16.0/viper.go#L423)

```
var SupportedRemoteProviders = []string{"etcd", "etcd3", "consul", "firestore"}
```

SupportedRemoteProviders are universally supported remote providers.

### 函数

#### func AddConfigPath 

``` go
func AddConfigPath(in string)
```

AddConfigPath adds a path for Viper to search for the config file in. Can be called multiple times to define multiple search paths.

#### func AddRemoteProvider 

``` go
func AddRemoteProvider(provider, endpoint, path string) error
```

AddRemoteProvider adds a remote configuration source. Remote Providers are searched in the order they are added. provider is a string value: "etcd", "etcd3", "consul" or "firestore" are currently supported. endpoint is the url. etcd requires [http://ip:port](http://ip:port/) consul requires ip:port path is the path in the k/v store to retrieve configuration To retrieve a config file called myapp.json from /configs/myapp.json you should set path to /configs and set config name (SetConfigName()) to "myapp"

#### func AddSecureRemoteProvider 

``` go
func AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
```

AddSecureRemoteProvider adds a remote configuration source. Secure Remote Providers are searched in the order they are added. provider is a string value: "etcd", "etcd3", "consul" or "firestore" are currently supported. endpoint is the url. etcd requires [http://ip:port](http://ip:port/) consul requires ip:port secretkeyring is the filepath to your openpgp secret keyring. e.g. /etc/secrets/myring.gpg path is the path in the k/v store to retrieve configuration To retrieve a config file called myapp.json from /configs/myapp.json you should set path to /configs and set config name (SetConfigName()) to "myapp" Secure Remote Providers are implemented with github.com/bketelsen/crypt

#### func AllKeys 

``` go
func AllKeys() []string
```

AllKeys returns all keys holding a value, regardless of where they are set. Nested keys are returned with a v.keyDelim separator

#### func AllSettings 

``` go
func AllSettings() map[string]interface{}
```

AllSettings merges all settings and returns them as a map[string]interface{}.

#### func AllowEmptyEnv <- 1.3.0

``` go
func AllowEmptyEnv(allowEmptyEnv bool)
```

AllowEmptyEnv tells Viper to consider set, but empty environment variables as valid values instead of falling back. For backward compatibility reasons this is false by default.

#### func AutomaticEnv 

``` go
func AutomaticEnv()
```

AutomaticEnv makes Viper check if environment variables match any of the existing keys (config, default or flags). If matching env vars are found, they are loaded into Viper.

#### func BindEnv 

``` go
func BindEnv(input ...string) error
```

BindEnv binds a Viper key to a ENV variable. ENV variables are case sensitive. If only a key is provided, it will use the env key matching the key, uppercased. If more arguments are provided, they will represent the env variable names that should bind to this key and will be taken in the specified order. EnvPrefix will be used when set when env name is not provided.

#### func BindFlagValue 

``` go
func BindFlagValue(key string, flag FlagValue) error
```

BindFlagValue binds a specific key to a FlagValue.

#### func BindFlagValues 

``` go
func BindFlagValues(flags FlagValueSet) error
```

BindFlagValues binds a full FlagValue set to the configuration, using each flag's long name as the config key.

#### func BindPFlag 

``` go
func BindPFlag(key string, flag *pflag.Flag) error
```

BindPFlag binds a specific key to a pflag (as used by cobra). Example (where serverCmd is a Cobra instance):

```
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
Viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
```

#### func BindPFlags 

``` go
func BindPFlags(flags *pflag.FlagSet) error
```

BindPFlags binds a full flag set to the configuration, using each flag's long name as the config key.

#### func ConfigFileUsed 

``` go
func ConfigFileUsed() string
```

ConfigFileUsed returns the file used to populate the config registry.

#### func Debug 

``` go
func Debug()
```

Debug prints all configuration registries for debugging purposes.

#### func DebugTo <- 1.13.0

``` go
func DebugTo(w io.Writer)
```

#### func Get 

``` go
func Get(key string) interface{}
```

Get can retrieve any value given the key to use. Get is case-insensitive for a key. Get has the behavior of returning the value associated with the first place from where it is set. Viper will check in the following order: override, flag, env, config file, key/value store, default

Get returns an interface. For a specific value use one of the Get____ methods.

#### func GetBool 

``` go
func GetBool(key string) bool
```

GetBool returns the value associated with the key as a boolean.

#### func GetDuration 

``` go
func GetDuration(key string) time.Duration
```

GetDuration returns the value associated with the key as a duration.

#### func GetFloat64 

``` go
func GetFloat64(key string) float64
```

GetFloat64 returns the value associated with the key as a float64.

#### func GetInt 

``` go
func GetInt(key string) int
```

GetInt returns the value associated with the key as an integer.

#### func GetInt32 <- 1.1.0

``` go
func GetInt32(key string) int32
```

GetInt32 returns the value associated with the key as an integer.

#### func GetInt64 

``` go
func GetInt64(key string) int64
```

GetInt64 returns the value associated with the key as an integer.

#### func GetIntSlice <- 1.5.0

``` go
func GetIntSlice(key string) []int
```

GetIntSlice returns the value associated with the key as a slice of int values.

#### func GetSizeInBytes 

``` go
func GetSizeInBytes(key string) uint
```

GetSizeInBytes returns the size of the value associated with the given key in bytes.

#### func GetString 

``` go
func GetString(key string) string
```

GetString returns the value associated with the key as a string.

#### func GetStringMap 

``` go
func GetStringMap(key string) map[string]interface{}
```

GetStringMap returns the value associated with the key as a map of interfaces.

#### func GetStringMapString 

``` go
func GetStringMapString(key string) map[string]string
```

GetStringMapString returns the value associated with the key as a map of strings.

#### func GetStringMapStringSlice 

``` go
func GetStringMapStringSlice(key string) map[string][]string
```

GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.

#### func GetStringSlice 

``` go
func GetStringSlice(key string) []string
```

GetStringSlice returns the value associated with the key as a slice of strings.

#### func GetTime 

``` go
func GetTime(key string) time.Time
```

GetTime returns the value associated with the key as time.

#### func GetUint <- 1.4.0

``` go
func GetUint(key string) uint
```

GetUint returns the value associated with the key as an unsigned integer.

#### func GetUint16 <- 1.13.0

``` go
func GetUint16(key string) uint16
```

GetUint16 returns the value associated with the key as an unsigned integer.

#### func GetUint32 <- 1.4.0

``` go
func GetUint32(key string) uint32
```

GetUint32 returns the value associated with the key as an unsigned integer.

#### func GetUint64 <- 1.4.0

``` go
func GetUint64(key string) uint64
```

GetUint64 returns the value associated with the key as an unsigned integer.

#### func InConfig 

``` go
func InConfig(key string) bool
```

InConfig checks to see if the given key (or an alias) is in the config file.

#### func IsSet 

``` go
func IsSet(key string) bool
```

IsSet checks to see if the key has been set in any of the data locations. IsSet is case-insensitive for a key.

#### func MergeConfig 

``` go
func MergeConfig(in io.Reader) error
```

MergeConfig merges a new configuration with an existing config.

#### func MergeConfigMap <- 1.3.0

``` go
func MergeConfigMap(cfg map[string]interface{}) error
```

MergeConfigMap merges the configuration from the map given with an existing config. Note that the map given may be modified.

#### func MergeInConfig 

``` go
func MergeInConfig() error
```

MergeInConfig merges a new configuration with an existing config.

#### func MustBindEnv <- 1.12.0

``` go
func MustBindEnv(input ...string)
```

MustBindEnv wraps BindEnv in a panic. If there is an error binding an environment variable, MustBindEnv will panic.

#### func OnConfigChange 

``` go
func OnConfigChange(run func(in fsnotify.Event))
```

OnConfigChange sets the event handler that is called when a config file changes.

#### func ReadConfig 

``` go
func ReadConfig(in io.Reader) error
```

ReadConfig will read a configuration file, setting existing keys to nil if the key does not exist in the file.

#### func ReadInConfig 

``` go
func ReadInConfig() error
```

ReadInConfig will discover and load the configuration file from disk and key/value stores, searching in one of the defined paths.

#### func ReadRemoteConfig 

``` go
func ReadRemoteConfig() error
```

ReadRemoteConfig attempts to get configuration from a remote source and read it in the remote configuration registry.

#### func RegisterAlias 

``` go
func RegisterAlias(alias string, key string)
```

RegisterAlias creates an alias that provides another accessor for the same key. This enables one to change a name without breaking the application.

#### func Reset 

``` go
func Reset()
```

Reset is intended for testing, will reset all to default settings. In the public interface for the viper package so applications can use it in their testing as well.

#### func SafeWriteConfig <- 1.0.1

``` go
func SafeWriteConfig() error
```

SafeWriteConfig writes current configuration to file only if the file does not exist.

#### func SafeWriteConfigAs <- 1.0.1

``` go
func SafeWriteConfigAs(filename string) error
```

SafeWriteConfigAs writes current configuration to a given filename if it does not exist.

#### func Set 

``` go
func Set(key string, value interface{})
```

Set sets the value for the key in the override register. Set is case-insensitive for a key. Will be used instead of values obtained via flags, config file, ENV, default, or key/value store.

#### func SetConfigFile 

``` go
func SetConfigFile(in string)
```

SetConfigFile explicitly defines the path, name and extension of the config file. Viper will use this and not check any of the config paths.

#### func SetConfigName 

``` go
func SetConfigName(in string)
```

SetConfigName sets name for the config file. Does not include extension.

#### func SetConfigPermissions <- 1.4.0

``` go
func SetConfigPermissions(perm os.FileMode)
```

SetConfigPermissions sets the permissions for the config file.

#### func SetConfigType 

``` go
func SetConfigType(in string)
```

SetConfigType sets the type of the configuration returned by the remote source, e.g. "json".

#### func SetDefault 

``` go
func SetDefault(key string, value interface{})
```

SetDefault sets the default value for this key. SetDefault is case-insensitive for a key. Default only used when no value is provided by the user via flag, config or ENV.

#### func SetEnvKeyReplacer 

``` go
func SetEnvKeyReplacer(r *strings.Replacer)
```

SetEnvKeyReplacer sets the strings.Replacer on the viper object Useful for mapping an environmental variable to a key that does not match it.

#### func SetEnvPrefix 

``` go
func SetEnvPrefix(in string)
```

SetEnvPrefix defines a prefix that ENVIRONMENT variables will use. E.g. if your prefix is "spf", the env registry will look for env variables that start with "SPF_".

#### func SetFs 

``` go
func SetFs(fs afero.Fs)
```

SetFs sets the filesystem to use to read configuration.

#### func SetTypeByDefaultValue 

``` go
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

``` go
func Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error
```

Unmarshal unmarshals the config into a Struct. Make sure that the tags on the fields of the structure are properly set.

#### func UnmarshalExact <- 1.6.0

``` go
func UnmarshalExact(rawVal interface{}, opts ...DecoderConfigOption) error
```

UnmarshalExact unmarshals the config into a Struct, erroring if a field is nonexistent in the destination struct.

#### func UnmarshalKey 

``` go
func UnmarshalKey(key string, rawVal interface{}, opts ...DecoderConfigOption) error
```

UnmarshalKey takes a single key and unmarshals it into a Struct.

#### func WatchConfig 

``` go
func WatchConfig()
```

WatchConfig starts watching a config file for changes.

#### func WatchRemoteConfig 

``` go
func WatchRemoteConfig() error
```

#### func WriteConfig <- 1.0.1

``` go
func WriteConfig() error
```

WriteConfig writes the current configuration to a file.

#### func WriteConfigAs <- 1.0.1

``` go
func WriteConfigAs(filename string) error
```

WriteConfigAs writes current configuration to a given filename.

### Types 

#### type ConfigFileAlreadyExistsError <- 1.6.0

``` go
type ConfigFileAlreadyExistsError string
```

ConfigFileAlreadyExistsError denotes failure to write new configuration file.

#### (ConfigFileAlreadyExistsError) Error <- 1.6.0

``` go
func (faee ConfigFileAlreadyExistsError) Error() string
```

Error returns the formatted error when configuration already exists.

#### type ConfigFileNotFoundError 

``` go
type ConfigFileNotFoundError struct {
	// contains filtered or unexported fields
}
```

ConfigFileNotFoundError denotes failing to find configuration file.

#### (ConfigFileNotFoundError) Error 

``` go
func (fnfe ConfigFileNotFoundError) Error() string
```

Error returns the formatted configuration error.

#### type ConfigMarshalError <- 1.0.1

``` go
type ConfigMarshalError struct {
	// contains filtered or unexported fields
}
```

ConfigMarshalError happens when failing to marshal the configuration.

#### (ConfigMarshalError) Error <- 1.0.1

``` go
func (e ConfigMarshalError) Error() string
```

Error returns the formatted configuration error.

#### type ConfigParseError 

``` go
type ConfigParseError struct {
	// contains filtered or unexported fields
}
```

ConfigParseError denotes failing to parse configuration file.

#### (ConfigParseError) Error 

``` go
func (pe ConfigParseError) Error() string
```

Error returns the formatted configuration error.

#### (ConfigParseError) Unwrap <- 1.16.0

``` go
func (pe ConfigParseError) Unwrap() error
```

Unwrap returns the wrapped error.

#### type DecoderConfigOption <- 1.1.0

``` go
type DecoderConfigOption func(*mapstructure.DecoderConfig)
```

A DecoderConfigOption can be passed to viper.Unmarshal to configure mapstructure.DecoderConfig options

#### func DecodeHook <- 1.1.0

``` go
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

``` go
type FlagValue interface {
	HasChanged() bool
	Name() string
	ValueString() string
	ValueType() string
}
```

FlagValue is an interface that users can implement to bind different flags to viper.

#### type FlagValueSet 

``` go
type FlagValueSet interface {
	VisitAll(fn func(FlagValue))
}
```

FlagValueSet is an interface that users can implement to bind a set of flags to viper.

#### type Logger <- 1.10.0

``` go
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

#### type Option <- 1.6.0

``` go
type Option interface {
	// contains filtered or unexported methods
}
```

Option configures Viper using the functional options paradigm popularized by Rob Pike and Dave Cheney. If you're unfamiliar with this style, see https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html and https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis.

#### func EnvKeyReplacer <- 1.6.0

``` go
func EnvKeyReplacer(r StringReplacer) Option
```

EnvKeyReplacer sets a replacer used for mapping environment variables to internal keys.

#### func IniLoadOptions <- 1.8.0

``` go
func IniLoadOptions(in ini.LoadOptions) Option
```

IniLoadOptions sets the load options for ini parsing.

#### func KeyDelimiter <- 1.6.0

``` go
func KeyDelimiter(d string) Option
```

KeyDelimiter sets the delimiter used for determining key parts. By default it's value is ".".

#### type RemoteConfigError 

``` go
type RemoteConfigError string
```

RemoteConfigError denotes encountering an error while trying to pull the configuration from the remote provider.

#### (RemoteConfigError) Error 

``` go
func (rce RemoteConfigError) Error() string
```

Error returns the formatted remote provider error

#### type RemoteProvider 

``` go
type RemoteProvider interface {
	Provider() string
	Endpoint() string
	Path() string
	SecretKeyring() string
}
```

RemoteProvider stores the configuration necessary to connect to a remote key/value store. Optional secretKeyring to unencrypt encrypted values can be provided.

#### type RemoteResponse 

``` go
type RemoteResponse struct {
	Value []byte
	Error error
}
```

#### type StringReplacer <- 1.6.0

``` go
type StringReplacer interface {
	// Replace returns a copy of s with all replacements performed.
	Replace(s string) string
}
```

StringReplacer applies a set of replacements to a string.

#### type UnsupportedConfigError 

``` go
type UnsupportedConfigError string
```

UnsupportedConfigError denotes encountering an unsupported configuration filetype.

#### (UnsupportedConfigError) Error 

``` go
func (str UnsupportedConfigError) Error() string
```

Error returns the formatted configuration error.

#### type UnsupportedRemoteProviderError 

``` go
type UnsupportedRemoteProviderError string
```

UnsupportedRemoteProviderError denotes encountering an unsupported remote provider. Currently only etcd and Consul are supported.

#### (UnsupportedRemoteProviderError) Error 

``` go
func (str UnsupportedRemoteProviderError) Error() string
```

Error returns the formatted remote provider error.

#### type Viper 

``` go
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

``` go
func GetViper() *Viper
```

GetViper gets the global Viper instance.

#### func New 

``` go
func New() *Viper
```

New returns an initialized Viper instance.

#### func NewWithOptions <- 1.6.0

``` go
func NewWithOptions(opts ...Option) *Viper
```

NewWithOptions creates a new Viper instance.

#### func Sub 

``` go
func Sub(key string) *Viper
```

Sub returns new Viper instance representing a sub tree of this instance. Sub is case-insensitive for a key.

#### (*Viper) AddConfigPath 

``` go
func (v *Viper) AddConfigPath(in string)
```

#### (*Viper) AddRemoteProvider 

``` go
func (v *Viper) AddRemoteProvider(provider, endpoint, path string) error
```

#### (*Viper) AddSecureRemoteProvider 

``` go
func (v *Viper) AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
```

#### (*Viper) AllKeys 

``` go
func (v *Viper) AllKeys() []string
```

#### (*Viper) AllSettings 

``` go
func (v *Viper) AllSettings() map[string]interface{}
```

#### (*Viper) AllowEmptyEnv <- 1.3.0

``` go
func (v *Viper) AllowEmptyEnv(allowEmptyEnv bool)
```

#### (*Viper) AutomaticEnv 

``` go
func (v *Viper) AutomaticEnv()
```

#### (*Viper) BindEnv 

``` go
func (v *Viper) BindEnv(input ...string) error
```

#### (*Viper) BindFlagValue 

``` go
func (v *Viper) BindFlagValue(key string, flag FlagValue) error
```

#### (*Viper) BindFlagValues 

``` go
func (v *Viper) BindFlagValues(flags FlagValueSet) (err error)
```

#### (*Viper) BindPFlag 

``` go
func (v *Viper) BindPFlag(key string, flag *pflag.Flag) error
```

#### (*Viper) BindPFlags 

``` go
func (v *Viper) BindPFlags(flags *pflag.FlagSet) error
```

#### (*Viper) ConfigFileUsed 

``` go
func (v *Viper) ConfigFileUsed() string
```

#### (*Viper) Debug 

``` go
func (v *Viper) Debug()
```

#### (*Viper) DebugTo <- 1.13.0

``` go
func (v *Viper) DebugTo(w io.Writer)
```

#### (*Viper) Get 

``` go
func (v *Viper) Get(key string) interface{}
```

#### (*Viper) GetBool 

``` go
func (v *Viper) GetBool(key string) bool
```

#### (*Viper) GetDuration 

``` go
func (v *Viper) GetDuration(key string) time.Duration
```

#### (*Viper) GetFloat64 

``` go
func (v *Viper) GetFloat64(key string) float64
```

#### (*Viper) GetInt 

``` go
func (v *Viper) GetInt(key string) int
```

#### (*Viper) GetInt32 <- 1.1.0

``` go
func (v *Viper) GetInt32(key string) int32
```

#### (*Viper) GetInt64 

``` go
func (v *Viper) GetInt64(key string) int64
```

#### (*Viper) GetIntSlice <- 1.5.0

``` go
func (v *Viper) GetIntSlice(key string) []int
```

#### (*Viper) GetSizeInBytes 

``` go
func (v *Viper) GetSizeInBytes(key string) uint
```

#### (*Viper) GetString 

``` go
func (v *Viper) GetString(key string) string
```

#### (*Viper) GetStringMap 

``` go
func (v *Viper) GetStringMap(key string) map[string]interface{}
```

#### (*Viper) GetStringMapString 

``` go
func (v *Viper) GetStringMapString(key string) map[string]string
```

#### (*Viper) GetStringMapStringSlice 

``` go
func (v *Viper) GetStringMapStringSlice(key string) map[string][]string
```

#### (*Viper) GetStringSlice 

``` go
func (v *Viper) GetStringSlice(key string) []string
```

#### (*Viper) GetTime 

``` go
func (v *Viper) GetTime(key string) time.Time
```

#### (*Viper) GetUint <- 1.4.0

``` go
func (v *Viper) GetUint(key string) uint
```

#### (*Viper) GetUint16 <- 1.13.0

``` go
func (v *Viper) GetUint16(key string) uint16
```

#### (*Viper) GetUint32 <- 1.4.0

``` go
func (v *Viper) GetUint32(key string) uint32
```

#### (*Viper) GetUint64 <- 1.4.0

``` go
func (v *Viper) GetUint64(key string) uint64
```

#### (*Viper) InConfig 

``` go
func (v *Viper) InConfig(key string) bool
```

#### (*Viper) IsSet 

``` go
func (v *Viper) IsSet(key string) bool
```

#### (*Viper) MergeConfig 

``` go
func (v *Viper) MergeConfig(in io.Reader) error
```

#### (*Viper) MergeConfigMap <- 1.3.0

``` go
func (v *Viper) MergeConfigMap(cfg map[string]interface{}) error
```

#### (*Viper) MergeInConfig 

``` go
func (v *Viper) MergeInConfig() error
```

#### (*Viper) MustBindEnv <- 1.12.0

``` go
func (v *Viper) MustBindEnv(input ...string)
```

#### (*Viper) OnConfigChange 

``` go
func (v *Viper) OnConfigChange(run func(in fsnotify.Event))
```

OnConfigChange sets the event handler that is called when a config file changes.

#### (*Viper) ReadConfig 

``` go
func (v *Viper) ReadConfig(in io.Reader) error
```

#### (*Viper) ReadInConfig 

``` go
func (v *Viper) ReadInConfig() error
```

#### (*Viper) ReadRemoteConfig 

``` go
func (v *Viper) ReadRemoteConfig() error
```

#### (*Viper) RegisterAlias 

``` go
func (v *Viper) RegisterAlias(alias string, key string)
```

#### (*Viper) SafeWriteConfig <- 1.0.1

``` go
func (v *Viper) SafeWriteConfig() error
```

#### (*Viper) SafeWriteConfigAs <- 1.0.1

``` go
func (v *Viper) SafeWriteConfigAs(filename string) error
```

#### (*Viper) Set 

``` go
func (v *Viper) Set(key string, value interface{})
```

#### (*Viper) SetConfigFile 

``` go
func (v *Viper) SetConfigFile(in string)
```

#### (*Viper) SetConfigName 

``` go
func (v *Viper) SetConfigName(in string)
```

#### (*Viper) SetConfigPermissions <- 1.4.0

``` go
func (v *Viper) SetConfigPermissions(perm os.FileMode)
```

#### (*Viper) SetConfigType 

``` go
func (v *Viper) SetConfigType(in string)
```

#### (*Viper) SetDefault 

``` go
func (v *Viper) SetDefault(key string, value interface{})
```

#### (*Viper) SetEnvKeyReplacer 

``` go
func (v *Viper) SetEnvKeyReplacer(r *strings.Replacer)
```

#### (*Viper) SetEnvPrefix 

``` go
func (v *Viper) SetEnvPrefix(in string)
```

#### (*Viper) SetFs 

``` go
func (v *Viper) SetFs(fs afero.Fs)
```

#### (*Viper) SetTypeByDefaultValue 

``` go
func (v *Viper) SetTypeByDefaultValue(enable bool)
```

#### (*Viper) Sub 

``` go
func (v *Viper) Sub(key string) *Viper
```

#### (*Viper) Unmarshal 

``` go
func (v *Viper) Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error
```

#### (*Viper) UnmarshalExact 

``` go
func (v *Viper) UnmarshalExact(rawVal interface{}, opts ...DecoderConfigOption) error
```

#### (*Viper) UnmarshalKey 

``` go
func (v *Viper) UnmarshalKey(key string, rawVal interface{}, opts ...DecoderConfigOption) error
```

#### (*Viper) WatchConfig 

``` go
func (v *Viper) WatchConfig()
```

WatchConfig starts watching a config file for changes.

#### (*Viper) WatchRemoteConfig 

``` go
func (v *Viper) WatchRemoteConfig() error
```

#### (*Viper) WatchRemoteConfigOnChannel 

``` go
func (v *Viper) WatchRemoteConfigOnChannel() error
```

#### (*Viper) WriteConfig <- 1.0.1

``` go
func (v *Viper) WriteConfig() error
```

#### (*Viper) WriteConfigAs <- 1.0.1

``` go
func (v *Viper) WriteConfigAs(filename string) error
```
