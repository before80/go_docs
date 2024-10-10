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

```sh
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

```go
var RemoteConfig remoteConfigFactory
```

​	RemoteConfig变量是可选的，参见remote包

[View Source](https://github.com/spf13/viper/blob/v1.16.0/viper.go#L420)

```go
var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"}
```

​	SupportedExts变量是通用支持的文件扩展名。

[View Source](https://github.com/spf13/viper/blob/v1.16.0/viper.go#L423)

```go
var SupportedRemoteProviders = []string{"etcd", "etcd3", "consul", "firestore"}
```

​	SupportedRemoteProviders变量是通用支持的远程提供者（remote providers）。

### 函数

#### func AddConfigPath 

``` go
func AddConfigPath(in string)
```

​	AddConfigPath函数添加要在其中搜索配置文件的路径。可以多次调用以定义多个搜索路径。

#### func AddRemoteProvider 

``` go
func AddRemoteProvider(provider, endpoint, path string) error
```

​	AddRemoteProvider函数添加远程配置源。远程提供者按添加顺序进行搜索。

- provider是一个字符串值："etcd"、"etcd3"、"consul"或"firestore"目前受支持。

- endpoint是URL。etcd需要`http://ip:port` 。consul需要`ip:port` 。

- path是在k/v存储中检索配置的路径。

​	要从`/configs/myapp.json`检索名为`myapp.json`的配置文件，应将`path`设置为`/configs`，并设置配置名（`SetConfigName()`）为"`myapp`"。

#### func AddSecureRemoteProvider 

``` go
func AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
```

​	AddSecureRemoteProvider函数添加远程安全配置源（remote configuration source.）。安全远程提供者（Secure Remote Providers）按添加顺序进行搜索。

- provider是一个字符串值："etcd"、"etcd3"、"consul"或"firestore"目前受支持。

- endpoint是URL。etcd需要`http://ip:port`， consul需要`ip:port` 

- secretkeyring是您的openpgp密钥环的文件路径。例如，`/etc/secrets/myring.gpg` 

- path是在k/v存储中检索配置的路径。

  

​	要从`/configs/myapp.json`检索名为`myapp.json`的配置文件，应将路径设置为`/configs`，并设置配置名（`SetConfigName()`）为"`myapp`"。Secure Remote Providers 是使用`github.com/bketelsen/crypt`实现的。

#### func AllKeys 

``` go
func AllKeys() []string
```

​	AllKeys函数返回所有具有值的键，无论它们在何处设置。嵌套的键使用`v.keyDelim`分隔符返回。

#### func AllSettings 

``` go
func AllSettings() map[string]interface{}
```

​	AllSettings函数合并所有设置并将它们作为一个`map[string]interface{}`返回。

#### func AllowEmptyEnv <- 1.3.0

``` go
func AllowEmptyEnv(allowEmptyEnv bool)
```

AllowEmptyEnv tells Viper to consider set, but empty environment variables as valid values instead of falling back. For backward compatibility reasons this is false by default.

​	AllowEmptyEnv函数告诉Viper将设置但为空的环境变量视为有效值，而不是返回默认值。出于向后兼容性的原因，默认值为false。（falling back 该怎么翻译？？）

#### func AutomaticEnv 

``` go
func AutomaticEnv()
```

AutomaticEnv makes Viper check if environment variables match any of the existing keys (config, default or flags). If matching env vars are found, they are loaded into Viper.

​	AutomaticEnv函数使Viper检查环境变量是否与任何现有键（配置、默认或标志）匹配。如果找到匹配的环境变量，则将它们加载到Viper中。

#### func BindEnv 

``` go
func BindEnv(input ...string) error
```

BindEnv binds a Viper key to a ENV variable. ENV variables are case sensitive. If only a key is provided, it will use the env key matching the key, uppercased. If more arguments are provided, they will represent the env variable names that should bind to this key and will be taken in the specified order. EnvPrefix will be used when set when env name is not provided.

​	BindEnv函数将 Viper 键绑定到环境变量。环境变量对大小写敏感。如果只提供键名，则它将使用与键名匹配且大写的环境变量键。如果提供更多参数，它们将表示应绑定到该键的环境变量名称，并按指定顺序进行处理。如果设置了 EnvPrefix，当未提供环境变量名称时将使用该前缀。

#### func BindFlagValue 

``` go
func BindFlagValue(key string, flag FlagValue) error
```

​	BindFlagValue函数将特定的key 绑定到FlagValue。

#### func BindFlagValues 

``` go
func BindFlagValues(flags FlagValueSet) error
```

BindFlagValues binds a full FlagValue set to the configuration, using each flag's long name as the config key.

​	BindFlagValues函数将完整的FlagValue集合绑定到配置，使用每个标志的长名称作为配置键。

#### func BindPFlag 

``` go
func BindPFlag(key string, flag *pflag.Flag) error
```

BindPFlag binds a specific key to a pflag (as used by cobra). Example (where serverCmd is a Cobra instance):

​	BindPFlag函数将特定的键绑定到pflag（由cobra使用）。示例（其中`serverCmd`是Cobra实例）：

```go
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
Viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
```

#### func BindPFlags 

``` go
func BindPFlags(flags *pflag.FlagSet) error
```

BindPFlags binds a full flag set to the configuration, using each flag's long name as the config key.

​	BindPFlags函数将完整的标志集合绑定到配置，使用每个标志的长名称作为配置键。

#### func ConfigFileUsed 

``` go
func ConfigFileUsed() string
```

ConfigFileUsed returns the file used to populate the config registry.

​	ConfigFileUsed函数返回用于填充配置注册表的文件。

#### func Debug 

``` go
func Debug()
```

​	Debug函数打印所有配置注册表以进行调试目的。

#### func DebugTo <- 1.13.0

``` go
func DebugTo(w io.Writer)
```

#### func Get 

``` go
func Get(key string) interface{}
```

​	Get函数可以通过使用键来检索任何值。对于键，Get不区分大小写。Get的行为是返回与它设置的第一个位置关联的值。Viper将按照以下顺序进行检查：覆盖（override）、标志（flag）、环境变量（env）、配置文件（config ）、键/值存储（key/value store）、默认值（default）。

​	Get函数返回一个interface{}类型的值。要获取特定的值，请使用`Get____`方法之一。

#### func GetBool 

``` go
func GetBool(key string) bool
```

​	GetBool 函数将与键关联的值转换为`bool`后返回。



#### func GetDuration 

``` go
func GetDuration(key string) time.Duration
```

​	GetDuration函数将与键关联的值转换为time.Duration后返回。

#### func GetFloat64 

``` go
func GetFloat64(key string) float64
```

​	GetFloat64函数将与键关联的值转换为float64后返回。

#### func GetInt 

``` go
func GetInt(key string) int
```

​	GetInt函数将与键关联的值转换为int后返回。

#### func GetInt32 <- 1.1.0

``` go
func GetInt32(key string) int32
```

​	GetInt32函数将与键关联的值转换为int32后返回。

#### func GetInt64 

``` go
func GetInt64(key string) int64
```

​	GetInt64函数将与键关联的值转换为int64后返回。

#### func GetIntSlice <- 1.5.0

``` go
func GetIntSlice(key string) []int
```

​	GetIntSlice函数将与键关联的值转换为`[]int`后返回。

#### func GetSizeInBytes 

``` go
func GetSizeInBytes(key string) uint
```

​	GetSizeInBytes函数将与键关联的值转换为`uint`后返回。

#### func GetString 

``` go
func GetString(key string) string
```

​	GetString函数将与键关联的值转换为`string`后返回。

#### func GetStringMap 

``` go
func GetStringMap(key string) map[string]interface{}
```

​	GetStringMap函数将与键关联的值转换为`map[string]interface{}`后返回。

#### func GetStringMapString 

``` go
func GetStringMapString(key string) map[string]string
```

​	GetStringMapString函数将与键关联的值转换为`map[string]string`后返回。

#### func GetStringMapStringSlice 

``` go
func GetStringMapStringSlice(key string) map[string][]string
```

​	GetStringMapStringSlice函数将与键关联的值转换为`map[string][]string`后返回。

#### func GetStringSlice 

``` go
func GetStringSlice(key string) []string
```

​	GetStringSlice函数将与键关联的值转换为`[]string`后返回。

#### func GetTime 

``` go
func GetTime(key string) time.Time
```

​	GetTime函数将与键关联的值转换为`time.Time`后返回。

#### func GetUint <- 1.4.0

``` go
func GetUint(key string) uint
```

​	GetUint函数将与键关联的值转换为`uint`后返回。

#### func GetUint16 <- 1.13.0

``` go
func GetUint16(key string) uint16
```

​	GetUint16函数将与键关联的值转换为`uint16`后返回。

#### func GetUint32 <- 1.4.0

``` go
func GetUint32(key string) uint32
```

​	GetUint32函数将与键关联的值转换为`uint32`后返回。

#### func GetUint64 <- 1.4.0

``` go
func GetUint64(key string) uint64
```

​	GetUint64函数将与键关联的值转换为`uint64`后返回。

#### func InConfig 

``` go
func InConfig(key string) bool
```

​	InConfig函数检查给定的键（或别名）是否在配置文件中。

#### func IsSet 

``` go
func IsSet(key string) bool
```

​	IsSet 检查键是否在任何数据位置被设置。对于键而言，IsSet函数不区分大小写。

#### func MergeConfig 

``` go
func MergeConfig(in io.Reader) error
```

​	MergeConfig函数将新配置与现有配置合并。

#### func MergeConfigMap <- 1.3.0

``` go
func MergeConfigMap(cfg map[string]interface{}) error
```

​	MergeConfigMap函数将给定的映射中的配置与现有配置合并。**请注意，给定的映射可能会被修改。**

#### func MergeInConfig 

``` go
func MergeInConfig() error
```

​	MergeInConfig 函数将新的配置与现有配置合并。

#### func MustBindEnv <- 1.12.0

``` go
func MustBindEnv(input ...string) { 
    v.MustBindEnv(input...) 
}

var v *Viper

func init() {
	v = New()
}

func (v *Viper) MustBindEnv(input ...string) {
	if err := v.BindEnv(input...); err != nil {
		panic(fmt.Sprintf("error while binding environment variable: %v", err))
	}
}
```

​	MustBindEnv 函数将 BindEnv 方法包装在 panic 中（如上函数的定义所示）。如果绑定环境变量时出现错误，MustBindEnv函数将引发 panic。

#### func OnConfigChange 

``` go
func OnConfigChange(run func(in fsnotify.Event))
```

​	OnConfigChange 函数设置（在配置文件发生更改时调用的）事件处理程序。

#### func ReadConfig 

``` go
func ReadConfig(in io.Reader) error
```

​	ReadConfig 函数会读取一个配置文件，并将不存在于该文件中的键设置为 nil。

#### func ReadInConfig 

``` go
func ReadInConfig() error { 
    return v.ReadInConfig() 
}

func (v *Viper) ReadInConfig() error {
	v.logger.Info("attempting to read in config file")
	filename, err := v.getConfigFile()
	if err != nil {
		return err
	}

	if !stringInSlice(v.getConfigType(), SupportedExts) {
		return UnsupportedConfigError(v.getConfigType())
	}

	v.logger.Debug("reading file", "file", filename)
	file, err := afero.ReadFile(v.fs, filename)
	if err != nil {
		return err
	}

	config := make(map[string]interface{})

	err = v.unmarshalReader(bytes.NewReader(file), config)
	if err != nil {
		return err
	}

	v.config = config
	return nil
}
```

ReadInConfig will discover and load the configuration file from disk and key/value stores, searching in one of the defined paths.

​	ReadInConfig 函数会从磁盘和键值存储中查找和加载（在已定义的路径中搜索）配置文件。

#### func ReadRemoteConfig 

``` go
func ReadRemoteConfig() error
```

ReadRemoteConfig attempts to get configuration from a remote source and read it in the remote configuration registry.

​	ReadRemoteConfig 函数试图从远程源获取配置并在远程配置注册表中读取。

#### func RegisterAlias 

``` go
func RegisterAlias(alias string, key string)
```

​	RegisterAlias 函数创建一个别名，为同一个键提供另一个访问器。这允许更改名称而不会破坏应用程序。

#### func Reset 

``` go
func Reset()
```

​	Reset 函数用于测试，将所有设置重置为默认设置。它在 viper 包的公共接口中，因此应用程序也可以在测试中使用它。

#### func SafeWriteConfig <- 1.0.1

``` go
func SafeWriteConfig() error
```

​	SafeWriteConfig 函数仅在文件不存在时将当前配置写入文件。

#### func SafeWriteConfigAs <- 1.0.1

``` go
func SafeWriteConfigAs(filename string) error
```

​	SafeWriteConfigAs 函数将当前配置写入给定的文件名（如果该文件不存在）。

#### func Set 

``` go
func Set(key string, value interface{})
```

​	Set 函数在覆盖注册表中设置键的值。Set函数对于键不区分大小写。将用于替代通过标志、配置文件、环境变量、默认值或键值存储获取的值。

​	Set 函数在**覆盖注册（override register）**中**设置**键的值。对于键，Set 函数不区分大小写。它将替代通过标志、配置文件、环境变量、默认值或键值存储获取的值。

#### func SetConfigFile 

``` go
func SetConfigFile(in string) { 
    v.SetConfigFile(in) 
}

func (v *Viper) SetConfigFile(in string) {
	if in != "" {
		v.configFile = in
	}
}
```

SetConfigFile explicitly defines the path, name and extension of the config file. Viper will use this and not check any of the config paths.

​	SetConfigFile 函数显式定义配置文件的路径、名称和扩展名。Viper 将使用此配置，并且不会检查任何配置路径。

#### func SetConfigName 

``` go
func SetConfigName(in string) { 
    v.SetConfigName(in) 
}

func (v *Viper) SetConfigName(in string) {
	if in != "" {
		v.configName = in
		v.configFile = ""
	}
}
```

​	SetConfigName 函数设置配置文件的名称。不包括扩展名。

#### func SetConfigPermissions <- 1.4.0

``` go
func SetConfigPermissions(perm os.FileMode) {
    v.SetConfigPermissions(perm) 
}

func (v *Viper) SetConfigPermissions(perm os.FileMode) {
	v.configPermissions = perm.Perm()
}
```

​	SetConfigPermissions 函数设置配置文件的权限。

#### func SetConfigType 

``` go
func SetConfigType(in string) { 
    v.SetConfigType(in) 
}

func (v *Viper) SetConfigType(in string) {
	if in != "" {
		v.configType = in
	}
}
```

​	SetConfigType 函数设置从远程源返回的配置的类型，例如 "`json`"。

#### func SetDefault 

``` go
func SetDefault(key string, value interface{}) { 
    v.SetDefault(key, value) 
}

func (v *Viper) SetDefault(key string, value interface{}) {
    // 如果传入了别名，则设置适当的默认值。
	key = v.realKey(strings.ToLower(key))
	value = toCaseInsensitiveValue(value)

	path := strings.Split(key, v.keyDelim)
	lastKey := strings.ToLower(path[len(path)-1])
	deepestMap := deepSearch(v.defaults, path[0:len(path)-1])

    // 设置最内层的值
	deepestMap[lastKey] = value
}
```

​	SetDefault函数为此键设置默认值。对于键来说，SetDefault函数是不区分大小写的。仅当用户未通过标志（flag）、配置（config）或环境变量（ENV）提供值时才使用默认值。

#### func SetEnvKeyReplacer 

``` go
func SetEnvKeyReplacer(r *strings.Replacer) {
    v.SetEnvKeyReplacer(r) 
}

func (v *Viper) SetEnvKeyReplacer(r *strings.Replacer) {
	v.envKeyReplacer = r
}
```

​	SetEnvKeyReplacer函数在viper对象上设置strings.Replacer。这对于将环境变量映射到与其不匹配的键非常有用。

#### func SetEnvPrefix 

``` go
func SetEnvPrefix(in string) { 
    v.SetEnvPrefix(in) 
}

func (v *Viper) SetEnvPrefix(in string) {
	if in != "" {
		v.envPrefix = in
	}
}
```

​	SetEnvPrefix 函数定义环境变量将使用的前缀。例如，如果前缀为 "spf"，则环境变量注册表将查找以 "SPF_" 开头的环境变量。

#### func SetFs 

``` go
func SetFs(fs afero.Fs) { 
    v.SetFs(fs) 
}

func (v *Viper) SetFs(fs afero.Fs) {
	v.fs = fs
}
```

​	SetFs函数设置要用于读取配置的文件系统。

#### func SetTypeByDefaultValue 

``` go
func SetTypeByDefaultValue(enable bool) { 
    v.SetTypeByDefaultValue(enable) 
}

func (v *Viper) SetTypeByDefaultValue(enable bool) {
	v.typeByDefValue = enable
}
```

SetTypeByDefaultValue enables or disables the inference of a key value's type when the Get function is used based upon a key's default value as opposed to the value returned based on the normal fetch logic.

​	SetTypeByDefaultValue 函数用于在使用 Get 函数时基于键的默认值来推断键值类型，而不是基于常规获取逻辑返回的值。该函数可以启用或禁用此行为。

​	例如，如果一个键的默认值为 `[]string{}`，并且通过环境变量将相同的键设置为 `"a b c"`，当使用 Get 函数调用该键时，如果键的类型是通过默认值推断的，则会返回一个字符串切片。因此，Get 函数将返回：

```go
[]string {"a", "b", "c"}
```

​	否则，Get 函数将返回：

```go
"a b c"
```

#### func Unmarshal 

``` go
func Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error {
	return v.Unmarshal(rawVal, opts...)
}

func (v *Viper) Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error {
	return decode(v.AllSettings(), defaultDecoderConfig(rawVal, opts...))
}
```

​	Unmarshal 函数将配置反序列化为一个结构体。请确保该结构体字段上的标签设置正确。

#### func UnmarshalExact <- 1.6.0

``` go
func UnmarshalExact(rawVal interface{}, opts ...DecoderConfigOption) error {
	return v.UnmarshalExact(rawVal, opts...)
}

func (v *Viper) UnmarshalExact(rawVal interface{}, opts ...DecoderConfigOption) error {
	config := defaultDecoderConfig(rawVal, opts...)
	config.ErrorUnused = true

	return decode(v.AllSettings(), config)
}
```

UnmarshalExact unmarshals the config into a Struct, erroring if a field is nonexistent in the destination struct.

​	UnmarshalExact 函数将配置反序列化为一个结构体，如果字段不存在于目标结构体中，则会产生错误。

#### func UnmarshalKey 

``` go
func UnmarshalKey(key string, rawVal interface{}, opts ...DecoderConfigOption) error {
	return v.UnmarshalKey(key, rawVal, opts...)
}

func (v *Viper) UnmarshalKey(key string, rawVal interface{}, opts ...DecoderConfigOption) error {
	return decode(v.Get(key), defaultDecoderConfig(rawVal, opts...))
}
```

UnmarshalKey takes a single key and unmarshals it into a Struct.

​	UnmarshalKey 函数接受一个键，并将其反序列化为一个结构体。

#### func WatchConfig 

``` go
func WatchConfig() {
    v.WatchConfig() 
}

// WatchConfig starts watching a config file for changes.
// WatchConfig 开始监视配置文件的变化。
func (v *Viper) WatchConfig() {
	initWG := sync.WaitGroup{}
	initWG.Add(1)
	go func() {
		watcher, err := newWatcher()
		if err != nil {
			v.logger.Error(fmt.Sprintf("failed to create watcher: %s", err))
			os.Exit(1)
		}
		defer watcher.Close()
        
        // 我们必须监视整个目录，以便跨平台捕捉重命名/原子保存
		filename, err := v.getConfigFile()
		if err != nil {
			v.logger.Error(fmt.Sprintf("get config file: %s", err))
			initWG.Done()
			return
		}

		configFile := filepath.Clean(filename)
		configDir, _ := filepath.Split(configFile)
		realConfigFile, _ := filepath.EvalSymlinks(filename)

		eventsWG := sync.WaitGroup{}
		eventsWG.Add(1)
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok { // 'Events' 通道已关闭
						eventsWG.Done()
						return
					}
					currentConfigFile, _ := filepath.EvalSymlinks(filename)

                    // 我们只关心以下情况的配置文件：
					// 1 - 配置文件被修改或创建
					// 2 - 配置文件的真实路径发生变化（例如：k8s ConfigMap 替换）
					if (filepath.Clean(event.Name) == configFile &&
						(event.Has(fsnotify.Write) || event.Has(fsnotify.Create))) ||
						(currentConfigFile != "" && currentConfigFile != realConfigFile) {
						realConfigFile = currentConfigFile
						err := v.ReadInConfig()
						if err != nil {
							v.logger.Error(fmt.Sprintf("read config file: %s", err))
						}
						if v.onConfigChange != nil {
							v.onConfigChange(event)
						}
					} else if filepath.Clean(event.Name) == configFile && event.Has(fsnotify.Remove) {
						eventsWG.Done()
						return
					}

				case err, ok := <-watcher.Errors:
					if ok { // 'Errors' 通道未关闭
						v.logger.Error(fmt.Sprintf("watcher error: %s", err))
					}
					eventsWG.Done()
					return
				}
			}
		}()
		watcher.Add(configDir)
		initWG.Done()   // 此 go 协程中的监视初始化完成，所以父协程可以继续执行...
		eventsWG.Wait() // 现在，等待此 go 协程中的事件循环结束...
	}()
	initWG.Wait() // 确保上述 go 协程完全结束后再返回
}
```

WatchConfig starts watching a config file for changes.

​	WatchConfig 函数开始监视配置文件的更改。

#### func WatchRemoteConfig 

``` go
func WatchRemoteConfig() error {
    return v.WatchRemoteConfig() 
}
func (v *Viper) WatchRemoteConfig() error {
	return v.watchKeyValueConfig()
}

// 检索第一个找到的远程配置。
func (v *Viper) watchKeyValueConfig() error {
	if len(v.remoteProviders) == 0 {
		return RemoteConfigError("No Remote Providers")
	}

	for _, rp := range v.remoteProviders {
		val, err := v.watchRemoteConfig(rp)
		if err != nil {
			v.logger.Error(fmt.Errorf("watch remote config: %w", err).Error())

			continue
		}
		v.kvstore = val
		return nil
	}
	return RemoteConfigError("No Files Found")
}
```

#### func WriteConfig <- 1.0.1

``` go
func WriteConfig() error { 
    return v.WriteConfig() 
}

func (v *Viper) WriteConfig() error {
	filename, err := v.getConfigFile()
	if err != nil {
		return err
	}
	return v.writeConfig(filename, true)
}
```

​	WriteConfig 函数将当前配置写入一个文件。

#### func WriteConfigAs <- 1.0.1

``` go
func WriteConfigAs(filename string) error { 
    return v.WriteConfigAs(filename) 
}

func (v *Viper) WriteConfigAs(filename string) error {
	return v.writeConfig(filename, true)
}

func (v *Viper) writeConfig(filename string, force bool) error {
	v.logger.Info("attempting to write configuration to file")

	var configType string

	ext := filepath.Ext(filename)
	if ext != "" && ext != filepath.Base(filename) {
		configType = ext[1:]
	} else {
		configType = v.configType
	}
	if configType == "" {
		return fmt.Errorf("config type could not be determined for %s", filename)
	}

	if !stringInSlice(configType, SupportedExts) {
		return UnsupportedConfigError(configType)
	}
	if v.config == nil {
		v.config = make(map[string]interface{})
	}
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	if !force {
		flags |= os.O_EXCL
	}
	f, err := v.fs.OpenFile(filename, flags, v.configPermissions)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := v.marshalWriter(f, configType); err != nil {
		return err
	}

	return f.Sync()
}
```

​	WriteConfigAs函数将当前配置写入一个给定的文件名中。

### 类型

#### type ConfigFileAlreadyExistsError <- 1.6.0

``` go
type ConfigFileAlreadyExistsError string
```

​	ConfigFileAlreadyExistsError类型表示写入新配置文件失败。

##### (ConfigFileAlreadyExistsError) Error <- 1.6.0

``` go
func (faee ConfigFileAlreadyExistsError) Error() string {
	return fmt.Sprintf("Config File %q Already Exists", string(faee))
}
```

Error returns the formatted error when configuration already exists.

​	Error方法返回配置已存在时的格式化错误信息。（？？不是配置文件已存在的错误？？）

#### type ConfigFileNotFoundError 

``` go
type ConfigFileNotFoundError struct {
	name, locations string
}
```

​	ConfigFileNotFoundError结构体表示无法找到配置文件。

##### (ConfigFileNotFoundError) Error 

``` go
func (fnfe ConfigFileNotFoundError) Error() string {
	return fmt.Sprintf("Config File %q Not Found in %q", fnfe.name, fnfe.locations)
}
```

​	Error方法返回格式化的配置错误信息。

#### type ConfigMarshalError <- 1.0.1

``` go
type ConfigMarshalError struct {
	err error
}
```

​	ConfigMarshalError结构体表示无法将配置进行系列化（Marshal）时发生错误。

##### (ConfigMarshalError) Error <- 1.0.1

``` go
func (e ConfigMarshalError) Error() string {
	return fmt.Sprintf("While marshaling config: %s", e.err.Error())
}
```

​	Error方法返回格式化的配置错误信息。

#### type ConfigParseError 

``` go
type ConfigParseError struct {
	err error
}
```

​	ConfigParseError结构体表示解析配置文件失败。

##### (ConfigParseError) Error 

``` go
func (pe ConfigParseError) Error() string {
	return fmt.Sprintf("While parsing config: %s", pe.err.Error())
}
```

​	Error方法返回格式化的配置错误信息。

##### (ConfigParseError) Unwrap <- 1.16.0

``` go
func (pe ConfigParseError) Unwrap() error {
	return pe.err
}
```

​	Unwrap方法返回包装的错误。

#### type DecoderConfigOption <- 1.1.0

``` go
type DecoderConfigOption func(*mapstructure.DecoderConfig)
```

​	可以将DecoderConfigOption传递给viper.Unmarshal，以配置mapstructure.DecoderConfig选项。

##### func DecodeHook <- 1.1.0

``` go
func DecodeHook(hook mapstructure.DecodeHookFunc) DecoderConfigOption
```

​	DecodeHook函数返回一个DecoderConfigOption，用于覆盖默认的DecoderConfig.DecodeHook值，其默认值为：

```go
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

​	FlagValue是一个接口，用户可以实现该接口来将不同的标志（flag）绑定到viper上。

#### type FlagValueSet 

``` go
type FlagValueSet interface {
	VisitAll(fn func(FlagValue))
}
```

​	FlagValueSet是一个接口，用户可以实现该接口来将一组标志（flag）绑定到viper上。

#### type Logger <- 1.10.0

``` go
type Logger interface {
    // Trace方法记录Trace事件。
	//
	// 比Debug事件提供更细粒度的信息。
	// 不支持此级别的日志记录器应该回退到Debug。
	Trace(msg string, keyvals ...interface{})
	
    // Debug方法记录Debug事件。
	//
	// 一系列详细信息事件。
	// 在调试系统时很有用。
	Debug(msg string, keyvals ...interface{})
	
    // Info方法记录Info事件。
	//
	// 关于系统内部发生的一般信息。
	Info(msg string, keyvals ...interface{})

    // Warn方法记录Warn(ing)事件。
	//
    // 应该查看的非致命（non-critical）事件。
	Warn(msg string, keyvals ...interface{})

    // Error方法记录Error事件。
	//
    // 需要立即注意的致命事件（critical events）。
	// 日志记录器通常在Error级别以上提供Fatal和Panic级别，
	// 但退出和恐慌超出了日志记录库的范围。
	Error(msg string, keyvals ...interface{})
}
```

​	Logger接口是各种日志记录用例和实践的统一接口，包括： 

- 分级日志记录
- 结构化日志记录

#### type Option <- 1.6.0

``` go
type Option interface {
	apply(v *Viper)
}
```

​	Option接口使用Rob Pike和Dave Cheney提倡的功能选项范式（functional options paradigm）配置Viper。如果您对此风格不熟悉，请参阅[https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)和[https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)。

##### func EnvKeyReplacer <- 1.6.0

``` go
func EnvKeyReplacer(r StringReplacer) Option {
	return optionFunc(func(v *Viper) {
		v.envKeyReplacer = r
	})
}

type optionFunc func(v *Viper)
```

EnvKeyReplacer sets a replacer used for mapping environment variables to internal keys.

​	EnvKeyReplacer函数设置用于将环境变量映射到内部键的替换器。

##### func IniLoadOptions <- 1.8.0

``` go
func IniLoadOptions(in ini.LoadOptions) Option
```

​	IniLoadOptions函数设置ini解析的加载选项。

##### func KeyDelimiter <- 1.6.0

``` go
func KeyDelimiter(d string) Option
```

KeyDelimiter sets the delimiter used for determining key parts. By default it's value is ".".

​	KeyDelimiter函数设置用于确定键部分的分隔符。默认值为"."。

#### type RemoteConfigError 

``` go
type RemoteConfigError string
```

RemoteConfigError denotes encountering an error while trying to pull the configuration from the remote provider.

​	RemoteConfigError类型表示在尝试从远程提供者获取配置时遇到错误。

##### (RemoteConfigError) Error 

``` go
func (rce RemoteConfigError) Error() string
```

​	Error方法返回格式化的远程提供者错误。

#### type RemoteProvider 

``` go
type RemoteProvider interface {
	Provider() string
	Endpoint() string
	Path() string
	SecretKeyring() string
}
```

​	RemoteProvider接口存储连接到远程键/值存储所需的配置。可提供secretKeyring以解密加密的值。

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
	// Replace 返回执行所有替换操作后的 s 的副本。
	Replace(s string) string
}
```

​	StringReplacer 接口对字符串应用一组替换。

#### type UnsupportedConfigError 

``` go
type UnsupportedConfigError string
```

​	UnsupportedConfigError 类型表示遇到不支持的配置文件类型。

##### (UnsupportedConfigError) Error 

``` go
func (str UnsupportedConfigError) Error() string {
	return fmt.Sprintf("Unsupported Config Type %q", string(str))
}
```

​	Error 方法返回格式化的配置错误。

#### type UnsupportedRemoteProviderError 

``` go
type UnsupportedRemoteProviderError string
```

​	UnsupportedRemoteProviderError 类型表示遇到不支持的远程提供者。目前只支持 etcd 和 Consul。

##### (UnsupportedRemoteProviderError) Error 

``` go
func (str UnsupportedRemoteProviderError) Error() string
```

​	Error 方法返回格式化的远程提供者错误。

#### type Viper 

``` go
type Viper struct {    
    // 分隔符，用于一次性访问嵌套值的键列表
	keyDelim string
    
    // 用于查找配置文件的路径集合
	configPaths []string

    // 用于读取配置的文件系统
	fs afero.Fs

    // 用于搜索配置的远程提供者集合
	remoteProviders []*defaultRemoteProvider

    // 在路径中查找的文件的名称
	configName        string
	configFile        string
	configType        string
	configPermissions os.FileMode
	envPrefix         string

    // ini 解析的特定命令
	iniLoadOptions ini.LoadOptions

	automaticEnvApplied bool
	envKeyReplacer      StringReplacer
	allowEmptyEnv       bool

	parents        []string
	config         map[string]interface{}
	override       map[string]interface{}
	defaults       map[string]interface{}
	kvstore        map[string]interface{}
	pflags         map[string]FlagValue
	env            map[string][]string
	aliases        map[string]string
	typeByDefValue bool

	onConfigChange func(fsnotify.Event)

	logger Logger

    // TODO: 应该使用互斥锁（mutex）进行保护
	encoderRegistry *encoding.EncoderRegistry
	decoderRegistry *encoding.DecoderRegistry
}
```

Viper is a prioritized configuration registry. It maintains a set of configuration sources, fetches values to populate those, and provides them according to the source's priority. The priority of the sources is the following: 1. overrides 2. flags 3. env. variables 4. config file 5. key/value store 6. defaults

​	Viper结构体是一个优先级配置注册表。它维护一组配置源，获取值以填充这些源，并根据源的优先级提供值。源的优先级如下：

1. 覆盖 
2. 标志
3. 境变量 ‘
4. 配置文件 
5. 键/值存储 
6. 默认值

​	例如，如果从以下源加载了值：

```json
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

​	结果配置将具有以下值：

```json
{
	"secret": "somesecretkey",
	"user": "root",
	"endpoint": "https://localhost"
}
```

​	注意：Vipers（对象）在并发的Get()和Set()操作中不是线程安全的。

##### func GetViper 

``` go
func GetViper() *Viper
```

​	GetViper函数获取全局 Viper 实例。

##### func New 

``` go
func New() *Viper
```

​	New 函数返回一个初始化后的 Viper 实例。

##### func NewWithOptions <- 1.6.0

``` go
func NewWithOptions(opts ...Option) *Viper
```

​	NewWithOptions函数创建一个新的 Viper 实例。

##### func Sub 

``` go
func Sub(key string) *Viper
```

Sub returns new Viper instance representing a sub tree of this instance. Sub is case-insensitive for a key.

​	Sub 函数返回一个表示当前实例子树的新Viper实例。Sub函数对于键是不区分大小写的。

##### (*Viper) AddConfigPath 

``` go
func (v *Viper) AddConfigPath(in string)
```

##### (*Viper) AddRemoteProvider 

``` go
func (v *Viper) AddRemoteProvider(provider, endpoint, path string) error
```

##### (*Viper) AddSecureRemoteProvider 

``` go
func (v *Viper) AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
```

##### (*Viper) AllKeys 

``` go
func (v *Viper) AllKeys() []string
```

##### (*Viper) AllSettings 

``` go
func (v *Viper) AllSettings() map[string]interface{}
```

##### (*Viper) AllowEmptyEnv <- 1.3.0

``` go
func (v *Viper) AllowEmptyEnv(allowEmptyEnv bool)
```

##### (*Viper) AutomaticEnv 

``` go
func (v *Viper) AutomaticEnv()
```

##### (*Viper) BindEnv 

``` go
func (v *Viper) BindEnv(input ...string) error
```

##### (*Viper) BindFlagValue 

``` go
func (v *Viper) BindFlagValue(key string, flag FlagValue) error
```

##### (*Viper) BindFlagValues 

``` go
func (v *Viper) BindFlagValues(flags FlagValueSet) (err error)
```

##### (*Viper) BindPFlag 

``` go
func (v *Viper) BindPFlag(key string, flag *pflag.Flag) error
```

##### (*Viper) BindPFlags 

``` go
func (v *Viper) BindPFlags(flags *pflag.FlagSet) error
```

##### (*Viper) ConfigFileUsed 

``` go
func (v *Viper) ConfigFileUsed() string
```

##### (*Viper) Debug 

``` go
func (v *Viper) Debug()
```

##### (*Viper) DebugTo <- 1.13.0

``` go
func (v *Viper) DebugTo(w io.Writer)
```

##### (*Viper) Get 

``` go
func (v *Viper) Get(key string) interface{}
```

##### (*Viper) GetBool 

``` go
func (v *Viper) GetBool(key string) bool
```

##### (*Viper) GetDuration 

``` go
func (v *Viper) GetDuration(key string) time.Duration
```

##### (*Viper) GetFloat64 

``` go
func (v *Viper) GetFloat64(key string) float64
```

##### (*Viper) GetInt 

``` go
func (v *Viper) GetInt(key string) int
```

##### (*Viper) GetInt32 <- 1.1.0

``` go
func (v *Viper) GetInt32(key string) int32
```

##### (*Viper) GetInt64 

``` go
func (v *Viper) GetInt64(key string) int64
```

##### (*Viper) GetIntSlice <- 1.5.0

``` go
func (v *Viper) GetIntSlice(key string) []int
```

##### (*Viper) GetSizeInBytes 

``` go
func (v *Viper) GetSizeInBytes(key string) uint
```

##### (*Viper) GetString 

``` go
func (v *Viper) GetString(key string) string
```

##### (*Viper) GetStringMap 

``` go
func (v *Viper) GetStringMap(key string) map[string]interface{}
```

##### (*Viper) GetStringMapString 

``` go
func (v *Viper) GetStringMapString(key string) map[string]string
```

##### (*Viper) GetStringMapStringSlice 

``` go
func (v *Viper) GetStringMapStringSlice(key string) map[string][]string
```

##### (*Viper) GetStringSlice 

``` go
func (v *Viper) GetStringSlice(key string) []string
```

##### (*Viper) GetTime 

``` go
func (v *Viper) GetTime(key string) time.Time
```

##### (*Viper) GetUint <- 1.4.0

``` go
func (v *Viper) GetUint(key string) uint
```

##### (*Viper) GetUint16 <- 1.13.0

``` go
func (v *Viper) GetUint16(key string) uint16
```

##### (*Viper) GetUint32 <- 1.4.0

``` go
func (v *Viper) GetUint32(key string) uint32
```

##### (*Viper) GetUint64 <- 1.4.0

``` go
func (v *Viper) GetUint64(key string) uint64
```

##### (*Viper) InConfig 

``` go
func (v *Viper) InConfig(key string) bool
```

##### (*Viper) IsSet 

``` go
func (v *Viper) IsSet(key string) bool
```

##### (*Viper) MergeConfig 

``` go
func (v *Viper) MergeConfig(in io.Reader) error
```

##### (*Viper) MergeConfigMap <- 1.3.0

``` go
func (v *Viper) MergeConfigMap(cfg map[string]interface{}) error
```

##### (*Viper) MergeInConfig 

``` go
func (v *Viper) MergeInConfig() error
```

##### (*Viper) MustBindEnv <- 1.12.0

``` go
func (v *Viper) MustBindEnv(input ...string)
```

##### (*Viper) OnConfigChange 

``` go
func (v *Viper) OnConfigChange(run func(in fsnotify.Event))
```

OnConfigChange sets the event handler that is called when a config file changes.

​	OnConfigChange 方法设置在配置文件更改时调用的事件处理程序。

##### (*Viper) ReadConfig 

``` go
func (v *Viper) ReadConfig(in io.Reader) error
```

##### (*Viper) ReadInConfig 

``` go
func (v *Viper) ReadInConfig() error
```

##### (*Viper) ReadRemoteConfig 

``` go
func (v *Viper) ReadRemoteConfig() error
```

##### (*Viper) RegisterAlias 

``` go
func (v *Viper) RegisterAlias(alias string, key string)
```

##### (*Viper) SafeWriteConfig <- 1.0.1

``` go
func (v *Viper) SafeWriteConfig() error
```

##### (*Viper) SafeWriteConfigAs <- 1.0.1

``` go
func (v *Viper) SafeWriteConfigAs(filename string) error
```

##### (*Viper) Set 

``` go
func (v *Viper) Set(key string, value interface{})
```

##### (*Viper) SetConfigFile 

``` go
func (v *Viper) SetConfigFile(in string)
```

##### (*Viper) SetConfigName 

``` go
func (v *Viper) SetConfigName(in string)
```

##### (*Viper) SetConfigPermissions <- 1.4.0

``` go
func (v *Viper) SetConfigPermissions(perm os.FileMode)
```

##### (*Viper) SetConfigType 

``` go
func (v *Viper) SetConfigType(in string)
```

##### (*Viper) SetDefault 

``` go
func (v *Viper) SetDefault(key string, value interface{})
```

##### (*Viper) SetEnvKeyReplacer 

``` go
func (v *Viper) SetEnvKeyReplacer(r *strings.Replacer)
```

##### (*Viper) SetEnvPrefix 

``` go
func (v *Viper) SetEnvPrefix(in string)
```

##### (*Viper) SetFs 

``` go
func (v *Viper) SetFs(fs afero.Fs)
```

##### (*Viper) SetTypeByDefaultValue 

``` go
func (v *Viper) SetTypeByDefaultValue(enable bool)
```

##### (*Viper) Sub 

``` go
func (v *Viper) Sub(key string) *Viper
```

##### (*Viper) Unmarshal 

``` go
func (v *Viper) Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error
```

##### (*Viper) UnmarshalExact 

``` go
func (v *Viper) UnmarshalExact(rawVal interface{}, opts ...DecoderConfigOption) error
```

##### (*Viper) UnmarshalKey 

``` go
func (v *Viper) UnmarshalKey(key string, rawVal interface{}, opts ...DecoderConfigOption) error
```

##### (*Viper) WatchConfig 

``` go
func (v *Viper) WatchConfig()
```

WatchConfig starts watching a config file for changes.

​	WatchConfig 方法开始监视配置文件的更改。

##### (*Viper) WatchRemoteConfig 

``` go
func (v *Viper) WatchRemoteConfig() error
```

##### (*Viper) WatchRemoteConfigOnChannel 

``` go
func (v *Viper) WatchRemoteConfigOnChannel() error
```

##### (*Viper) WriteConfig <- 1.0.1

``` go
func (v *Viper) WriteConfig() error
```

##### (*Viper) WriteConfigAs <- 1.0.1

``` go
func (v *Viper) WriteConfigAs(filename string) error
```
