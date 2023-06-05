+++
title = "环境变量"
date = 2023-05-17T09:59:21+08:00
weight = 15
description = ""
isCJKLanguage = true
draft = false
+++
## Environment variables 环境变量

> 原文：[https://go.dev/ref/mod#environment-variables](https://go.dev/ref/mod#environment-variables)

​	`go`命令中的模块行为可以使用下面列出的环境变量进行配置。这个列表只包括与模块相关的环境变量。有关`go`命令识别的所有环境变量的列表，请参阅[go help environment]({{< ref "/cmd/go#environment-variables">}})。

### GO111MODULE

​	控制`go`命令是在模块感知模式还是`GOPATH`模式下运行。有三个值可以识别：

- `off`：`go`命令忽略`go.mod`文件，并在`GOPATH`模式下运行。

- `on` (或unset): `go`命令在模块感知模式下运行，即使不存在`go.mod`文件。

- `auto`：如果当前目录或任何父目录中存在`go.mod`文件，`go`命令将以模块感知模式运行。在Go 1.15及以下版本中，这是默认的。

​	更多信息请参见[模块感知命令](../Module-awareCommands)。

### GOMODCACHE

​	`go`命令将存储下载的模块和相关文件的目录。关于此目录结构的详细信息，请参见[模块缓存](../ModuleCache)。如果未设置`GOMODCACHE`，它默认为`$GOPATH/pkg/mod`。

### GOINSECURE

​	以逗号分隔的模块路径前缀 glob 模式列表（采用 Go的 [path.Match](https://go.dev/pkg/path/#Match) 的语法），可能总是以不安全的方式获取。仅适用于被直接获取的依赖项。与`go get`上的`-insecure`标志不同，`GOINSECURE`不会禁用模块校验和数据库验证（module checksum database validation）。可以用`GOPRIVATE`或`GONOSUMDB`来实现这一点。

### GONOPROXY

​	以逗号分隔的模块路径前缀 glob 模式列表（采用Go的[path.Match](https://go.dev/pkg/path/#Match)语法），应该总是直接从版本控制存储库获取，而不是从模块代理获取。如果没有设置`GONOPROXY`，它默认为`GOPRIVATE`。参见[Privacy（隐私）](../PrivateModules#privacy)。

### GONOSUMDB

​	以逗号分隔的模块路径前缀 glob 模式列表（采用Go的[path.Match](https://go.dev/pkg/path/#Match)语法），`go` 不应使用校验数据库来验证校验。如果未设置 `GONOSUMDB`，它默认为 `GOPRIVATE`。参见[Privacy（隐私）](../PrivateModules#privacy)。

### GOPATH

​	在 `GOPATH` 模式中，`GOPATH` 变量是一个可能包含 Go 代码的目录列表。在模块感知模式下，[module cache（模块缓存）](../Glossary#module-cache)存储在第一个 `GOPATH` 目录的 `pkg/mod` 子目录中。缓存之外的模块源代码可以存储在任何目录中。如果未设置`GOPATH`，它默认为用户的主目录（home directory）的`go`子目录。

### GOPRIVATE

​	以逗号分隔的模块路径前缀的 glob 模式列表（采用Go的[path.Match](https://go.dev/pkg/path/#Match)语法），应被视认为是私有的。`GOPRIVATE`是`GONOPROXY`和`GONOSUMDB`的默认值。参见[Privacy（隐私）](../PrivateModules#privacy)。`GOPRIVATE` 也决定了一个模块是否被视为 `GOVCS` 的私有模块。

### GOPROXY

​	模块代理URL的列表，用逗号（`,`）或管道符 `|`分隔。当`go`命令查询有关模块的信息时，它会依次联系列表中的每个代理，直到收到成功的响应或终端错误。代理可能会响应404（Not Found）或410（Gone）状态，表示该模块在该服务器上不可用。

​	`go`命令的错误回退行为（error fallback behavior）是由URL之间的分隔符决定的。如果代理URL后面是逗号，`go`命令在404或410错误后会回退到下一个URL；所有其他错误都被视为终端错误。如果代理URL后面跟着管道，则`go`命令在出现任何错误(包括超时等非HTTP错误)后回退到下一个源。`GOPROXY` URLs可以有`https`、`http`或`file`等协议。如果一个URL没有协议，则假定为`https`。模块缓存可以直接作为文件代理使用：

`GOPROXY=file://$(go env GOMODCACHE)/cache/download`

有两个关键词可以用来代替代理URL：

- `off`：不允许从任何源下载模块。

- `direct`： 直接从版本控制存储库下载，而不是使用模块代理。

  

​	`GOPROXY`默认为`https://proxy.golang.org,direct`。在该配置下，`go`命令首先联系Google运行的Go模块镜像，如果该镜像没有模块，则回退到直接连接。请参见[https://proxy.golang.org/privacy](https://proxy.golang.org/privacy)了解镜像的隐私政策。可以设置 `GOPRIVATE` 和 `GONOPROXY` 环境变量，以防止使用代理下载特定模块。关于隐私代理配置的信息，请参见[Privacy（隐私）](../PrivateModules#privacy)。

​	有关如何使用代理的更多信息，请参见[Module proxies（模块代理）](../ModuleProxies)和[Resolving a package to a module（将包解析为模块）](../ModulesPackagesAndVersions#resolving-a-package-to-a-module)。

### GOSUMDB

​	标识要使用的校验和数据库的名称以及可选的公钥和 URL。例如：

```
GOSUMDB="sum.golang.org"
GOSUMDB="sum.golang.org+<publickey>"
GOSUMDB="sum.golang.org+<publickey> https://sum.golang.org"
```

​	`go`命令知道`sum.golang.org`的公钥，也知道名称为`sum.golang.google.cn`（在中国大陆可用）连接到`sum.golang.org`数据库；使用任何其他数据库都需要显式地给出公钥。URL默认为`https://`后跟数据库名。

​	`GOSUMDB`默认为`sum.golang.org`，由Google运行的Go校验和数据库（checksum database）。有关该服务的隐私政策，请参见 [https://sum.golang.org/privacy](https://sum.golang.org/privacy)。

​	如果 `GOSUMDB` 被设置为`off`，或者在调用 `go get` 时使用了 `-insecure` 标志，则不会查询校验和数据库，所有未识别的模块都会被接受，代价是放弃对所有模块进行验证的可重复下载的安全保证。绕过特定模块的校验和数据库的更好方法是使用`GOPRIVATE`或`GONOSUMDB`环境变量。

​	更多信息请参见[Authenticating modules（验证模块）](../AuthenticatingModules)和[Privacy（隐私）](../PrivateModules#privacy)。

### GOVCS

​	控制`go`命令在下载公共和私有模块（由其路径是否匹配`GOPRIVATE`中的模式来定义）或其他与glob模式相匹配的模块时可以使用的版本控制工具集。如果没有设置`GOVCS`，或者某个模块不符合`GOVCS`中的任何模式，`go`命令可以对公共模块使用`git`和`hg`，对私有模块使用任何已知的版本控制工具。具体来说，`go`命令的行为就像`GOVCS`被设置为`public:git|hg,private:all`，完整的解释请参见 [用`GOVCS`控制版本控制工具](../VersionControlSystems#controlling-version-control-tools-with-govcs)。

### GOWORK

​	`GOWORK`环境变量指示`go`命令使用提供的[`go.work`文件](#go-work-file)进入工作区模式。如果`GOWORK`被设置为`off`，工作区模式将被禁用。这可以用来在单模块模式下运行`go`命令：例如，`GOWORK=off go build .`在单模块模式下构建`.`包。如果`GOWORK`为空，`go`命令将搜索`go.work`文件，如[工作区](#workspaces)部分所述。

