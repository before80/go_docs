+++
title = "私有模块"
date = 2023-05-17T09:59:21+08:00
weight = 12
description = ""
isCJKLanguage = true
draft = false
+++
## Private modules 私有模块

> 原文：[https://go.dev/ref/mod#private-modules](https://go.dev/ref/mod#private-modules)

​	Go模块经常是在版本控制服务器和模块代理上开发和发布的，这些模块在公共互联网上是不可用的。`go` 命令可以从私有资源中下载和构建模块，不过通常需要一些配置。

​	下面的环境变量可以用来配置对私有模块的访问。详情请参见[Environment variables（环境变量）](../EnvironmentVariables)。有关控制发送到公共服务器的信息的信息，也请参见[Privacy（隐私）](#privacy)。

- `GOPROXY` —— 模块代理URL的列表。`go` 命令将尝试按顺序从每个服务器下载模块。关键字 `direct` 指示 `go` 命令从其开发的版本控制存储库中下载模块，而不是使用代理。
- `GOPRIVATE` —— 应该被视为私有的模块路径前缀的 glob 模式列表。作为 `GONOPROXY` 和 `GONOSUMDB` 的默认值。
- `GONOPROXY` —— 不应从代理下载的模块路径前缀的 glob 模式列表。`go` 命令将从开发模块的版本控制存储库中下载匹配的模块，而不考虑 `GOPROXY`。
- `GONOSUMDB` —— 不应使用公共校验和数据库 [sum.golang.org](https://sum.golang.org/) 进行检查的模块路径前缀的 glob 模式列表。
- `GOINSECURE` —— 可以通过HTTP和其他不安全协议检索的模块路径前缀的 glob 模式列表。

​	这些变量可以在开发环境中设置（例如，在 `.profile` 文件中），也可以用 [go env -w](../../CommandDocumentation/go#print-go-environment-information) 永久设置。

​	本节的其余部分描述了提供访问私有模块代理和版本控制存储库的常见模式。

### Private proxy serving all modules 为所有模块提供服务的私有代理

​	为所有模块（公共和私有）提供服务的中央私有代理服务器为管理员提供了最大程度的控制，并且对单个开发人员来说需要的配置最少。

​	要配置`go`命令以使用这样的服务器，请设置以下环境变量，将`https://proxy.corp.example.com`替换为您的代理URL，将`corp.example.com`替换为您的模块前缀。

```
GOPROXY=https://proxy.corp.example.com
GONOSUMDB=corp.example.com
```

​	`GOPROXY` 设置指示 `go` 命令仅从 `https://proxy.corp.example.com` 下载模块；`go` 命令不会连接到其他代理或版本控制存储库。

​	`GONOSUMDB` 设置指示 `go` 命令不使用公共校验和数据库来验证路径以 `corp.example.com` 开头的模块。

​	以这种配置运行的代理可能需要对私有版本控制服务器进行读取访问。它还需要访问公共互联网来下载公共模块的新版本。

​	有几种现有的`GOPROXY`服务器实现可以以这种方式使用。最小的实现将从[module cache（模块缓存）](../Glossary#module-cache)目录中提供文件，并使用[go mod download](../gomodFiles#go-mod-download)（带有适当的配置）来检索丢失的模块。

### Private proxy serving private modules 为私有模块提供服务的私有代理

​	私人代理服务器可以为私人模块提供服务，而不需要为公开的模块提供服务。`go`命令可以被配置为在私人服务器上无法使用的模块时回退到公共源。

​	要配置`go`命令以这种方式工作，请设置以下环境变量，将`https://proxy.corp.example.com`替换为代理URL，将`corp.example.com`替换为模块前缀：

```
GOPROXY=https://proxy.corp.example.com,https://proxy.golang.org,direct
GONOSUMDB=corp.example.com
```

​	`GOPROXY` 设置指示 `go` 命令首先尝试从 `https://proxy.corp.example.com` 下载模块。如果该服务器的回应是404 (Not Found)或410 (Gone)，`go`命令将回退到`https://proxy.golang.org`，在之后才是直接连接到存储库。

​	`GONOSUMDB` 设置指示 `go` 命令不使用公共校验和数据库来验证路径以 `corp.example.com` 开头的模块。

​	请注意，在这种配置下使用的代理仍然可以控制对公共模块的访问，即使它不为公共模块提供服务。如果代理以404或410以外的错误状态响应请求，`go`命令将不会回退到`GOPROXY`列表的后面条目。例如，代理可能会对具有不合适许可证或具有已知安全漏洞的模块响应403 (Forbidden) 。

### Direct access to private modules 直接访问私有模块

​	`go`命令可以被配置为绕过公共代理，并直接从版本控制服务器下载私有模块。当无法运行私有代理服务器时，这非常有用。

​	要配置 `go` 命令以这种方式工作，请设置 `GOPRIVATE`，将私有模块前缀替换为 `corp.example.com` ：

```
GOPRIVATE=corp.example.com
```

​	在这种情况下，不需要更改`GOPROXY`变量。它的默认值是`https://proxy.golang.org,direct`，它指示`go` 命令首先尝试从`https://proxy.golang.org`，然后在代理响应404 (Not Found)或410 (Gone)时回退到直接连接。

​	`GOPRIVATE`设置指示`go` 命令不要连接到代理或以`corp.example.com`开头的模块的校验和数据库。

​	可能仍需要内部HTTP服务器来将模块路径解析为存储库URL。例如，当 `go` 命令下载模块 `corp.example.com/mod` 时，它将向 `https://corp.example.com/mod?go-get=1`发送一个 GET 请求，并在响应中查找存储库的 URL。要避免这种需求，请确保每个私有模块路径都有一个`VCS`后缀（如`.git`）来标记存储库根的前缀。例如，当 `go` 命令下载模块 `corp.example.com/repo.git/mod` 时，它将在 `https://corp.example.com/repo.git` 或 `ssh://corp.example.com/repo.git`中克隆 Git 存储库，而不需要发出额外的请求。

​	开发人员将需要对包含私有模块的存储库具有读取权限。这可以在全局`VCS`配置文件中配置，如`.gitconfig`。好将`VCS`工具配置为不需要交互式身份验证提示。默认情况下，在调用Git时，`go` 命令通过设置`GIT_TERMINAL_PROMPT=0`来禁用交互式提示，但它会遵守显式设置。

### Passing credentials to private proxies 将凭据传递给私有代理

​	`go` 命令在与代理服务器通信时支持 HTTP[basic authentication（基本身份验证）](https://en.wikipedia.org/wiki/Basic_access_authentication)。

​	凭证可以在[.netrc 文件](https://www.gnu.org/software/inetutils/manual/html_node/The-_002enetrc-file.html)中指定。例如，包含以下几行的`.netrc`文件将配置`go` 命令用给定的用户名和密码连接到`proxy.corp.example.com`机器上。

```
machine proxy.corp.example.com
login jrgopher
password hunter2
```

​	文件的位置可以通过`NETRC`环境变量来设置。如果没有设置`NETRC`，`go` 命令将读取`$HOME/.netrc`(在类UNIX平台上)，或读取`%USERPROFILE%\_netrc`(在Windows上)。

​	`.netrc`中的字段用空格、制表符和换行符分隔。遗憾的是，这些字符不能用于用户名或密码。还要注意的是，计算机名不能是一个完整的URL，因此不能为同一台机器上的不同路径指定不同的用户名和密码。

​	另外，凭证可以直接在`GOPROXY`的URL中指定。例如：

```
GOPROXY=https://jrgopher:hunter2@proxy.corp.example.com
```

​	采取这种方法时要小心：环境变量可能会出现在shell历史记录和日志中。

### Passing credentials to private repositories 将凭证传递给私有存储库

`go` 命令可以直接从版本控制储存库中下载模块。如果没有使用私有代理，这对私有模块是必要的。参见[Direct access to private modules（直接访问私有模块）](#direct-access-to-private-modules)的配置。

​	`go` 命令在直接下载模块时运行`git`等版本控制工具。这些工具执行它们自己的身份验证，所以您可能需要在工具特定的配置文件（如`.gitconfig`）中配置凭证。

​	为了确保这项工作顺利进行，请确保`go` 命令使用正确的存储库URL，并且版本控制工具不需要交互式输入密码。`go` 命令优先使用 `https://` URL，而不是`ssh://`等其他方案，除非在[查找存储库 URL](../VersionControlSystems#finding-a-repository-for-a-module-path) 时指定了方案。特别是对于 GitHub 存储库，`go` 命令假定为 `https://`。

​	对于大多数服务器，您可以将客户端配置为通过HTTP进行身份验证。例如，GitHub 支持使用 [OAuth 个人访问令牌作为 HTTP 密码](https://docs.github.com/en/free-pro-team@latest/github/extending-github/git-automation-with-oauth-tokens)。您可以将HTTP密码存储在`.netrc`文件中，就像[将凭证传递给私有代理](#passing-credentials-to-private-proxies)时一样。

​	另外，您也可以将`https://` URL重写成另一种方案。例如，在`.gitconfig`中：

```
[url "git@github.com:"]
    insteadOf = https://github.com/
```

​	更多信息，请参见[为什么 "go get "在克隆存储库时使用HTTPS？](../../../FAQ_En_Zh#why-does-go-get-use-https-when-cloning-a-repository-go-get-https)

### Privacy 隐私

​	`go` 命令可以从模块代理服务器和版本控制系统下载模块和元数据。环境变量`GOPROXY`控制使用哪些服务器。环境变量`GOPRIVATE`和`GONOPROXY`控制从代理服务器获取哪些模块。

​	`GOPROXY`的默认值是：

```
https://proxy.golang.org,direct
```

​	在此设置下，当`go` 命令下载模块或模块元数据时，它将首先向`proxy.golang.org`发送请求，这是一个由谷歌（[privacy policy（隐私策略）](https://proxy.golang.org/privacy)）运营的公共模块代理。有关在每个请求中发送哪些信息的详细信息，请参见[GOPROXY 协议](../ModuleProxies#goproxy-protocol)。`go` 命令不会传输个人身份信息，但它会传输所请求的完整模块路径。如果代理以404 (Not Found)或410 (Gone) 状态响应，`go` 命令将尝试直接连接到提供该模块的版本控制系统。详见[Version control systems（版本控制系统）](../VersionControlSystems)。

​	`GOPRIVATE` 或 `GONOPROXY` 环境变量可以被设置为与模块前缀匹配的 glob 模式列表，这些模块前缀是私有的，不应该从任何代理请求。例如：

```
GOPRIVATE=*.corp.example.com,*.research.example.com
```

​	`GOPRIVATE`只是作为`GONOPROXY`和`GONOSUMDB`的默认值，因此没有必要设置`GONOPROXY`，除非`GONOSUMDB`应该有一个不同的值。当模块路径与`GONOPROXY`匹配时，`go` 命令将忽略该模块的`GOPROXY`，并直接从其版本控制存储库中获取它。当没有代理为私有模块提供服务时，这很有用。参见[Direct access to private modules（直接访问私有模块）](#direct-access-to-private-modules)。

​	如果有一个[受信任的代理为所有模块提供服务](#private-proxy-serving-all-modules)，那么`GONOPROXY`就不应该被设置。例如，如果 `GOPROXY` 被设置为一个源，那么`go` 命令将不会从其他源下载模块。在这种情况下仍应设置`GONOSUMDB`。

```
GOPROXY=https://proxy.corp.example.com
GONOSUMDB=*.corp.example.com,*.research.example.com
```

​	如果有一个[受信任的代理仅为私有模块提供服务](#private-proxy-serving-private-modules)，那么不应该设置`GONOPROXY`，但必须注意确保该代理以正确的状态码进行响应。例如，请考虑以下配置：

```
GOPROXY=https://proxy.corp.example.com,https://proxy.golang.org
GONOSUMDB=*.corp.example.com,*.research.example.com
```

​	假设由于输入错误，开发人员试图下载一个不存在的模块。

```
go mod download corp.example.com/secret-product/typo@latest
```

​	`go` 命令首先从`proxy.corp.example.com`请求此模块。如果该代理响应404 (Not Found)或410 (Gone)，那么`go` 命令将回退到`proxy.golang.org`，在请求URL中传输`secret-product`路径。如果私人代理响应任何其他错误码，那么`go` 命令将打印错误，并且不会回退到其他源。

​	除了代理之外，`go` 命令还可以连接到校验和数据库，以验证`go.sum`中没有列出的模块的加密散列。`GOSUMDB`环境变量设置校验和数据库的名称、URL和公钥。`GOSUMDB` 的默认值是 `sum.golang.org`，它是由 Google （[privacy policy（隐私策略）](https://sum.golang.org/privacy)）运营的公共校验和数据库。关于每次请求所传送的内容，请参见[Checksum database（校验和数据库）](../AuthenticatingModules#checksum-database)。与代理一样，`go` 命令不会传输个人身份信息，但它会传输所请求的完整模块路径，而且校验和数据库无法计算非公共模块的校验和。

​	可以将`GONOSUMDB`环境变量设置为指示哪些模块是私有的并且不应从校验和数据库请求的模式。 `GOPRIVATE`作为`GONOSUMDB`和`GONOPROXY`的默认值，因此没有必要设置`GONOSUMDB`，除非`GONOPROXY`应该有不同的值。

​	代理可以[mirror the checksum database（镜像校验和数据库）](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#proxying-a-checksum-database)。如果`GOPROXY`中的代理执行此操作，`go` 命令将不会直接连接到校验和数据库。

​	`GOSUMDB`可以被设置为`off`，以完全禁止使用校验和数据库。使用此设置时，`go` 命令将不会对下载的模块进行身份验证，除非它们已经在`go.sum`中。请参阅[Authenticating modules（验证模块）](../AuthenticatingModules)。