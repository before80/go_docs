+++
title = "模块代理"
date = 2023-05-17T09:59:21+08:00
weight = 9
description = ""
isCJKLanguage = true
draft = false
+++
## Module proxies 模块代理

> 原文：[https://go.dev/ref/mod#module-proxy](https://go.dev/ref/mod#module-proxy)

### GOPROXY protocol

​	模块代理是一个HTTP服务器，可以响应下面指定路径的`GET`请求。这些请求没有查询参数，也不需要特定的头，因此即使是一个从固定文件系统（包括`file://` URL）提供服务的网站也可以成为一个模块代理。

​	成功的HTTP响应必须有状态码200 (OK)。遵循重定向(3xx)。状态码为4xx和5xx的响应被视为错误。错误码 404 (Not Found) 和410 (Gone)表示请求的模块或版本在代理上不可用，但可能在其他地方找到。错误响应的内容类型应该是`text/plain`，`charset`是`utf-8`或`us-ascii`。

​	`go`命令可以配置为使用`GOPROXY`环境变量联系代理或源码管理服务器，该环境变量接受一个代理URL列表。列表中可以包括关键词`direct`或`off`（详见[Environment variables（环境变量）]()）。列表中的元素可以用逗号（`,`）或管道符（`|`）分隔，这决定了错误回退行为。当URL后跟一个逗号时，`go`命令只有在出现404 (Not Found)或410 (Gone) 响应后才会回退到后面的源。当URL后跟一个管道符时，`go`命令在任何错误（包括超时等非HTTP错误）发生后都会回退到后面的源。这种错误处理行为让代理充当未知模块的守门员。例如，对于不在批准列表上的模块，代理可以响应403 (Forbidden)错误（请参见 [Private proxy serving private modules（为私有模块提供服务的私有代理）](../PrivateModules#private-proxy-serving-private-modules)）。

​	下表列出了模块代理必须响应的查询。对于每个路径，`$base`是代理URL的路径部分，`$module`是模块路径，`$version`是版本。例如，如果代理的URL是`https://example.com/mod`，并且客户端正在为 `v0.3.2`版本的模块 `golang.org/x/text` 请求 `go.mod` 文件，那么客户端将为`https://example.com/mod/golang.org/x/text/@v/v0.3.2.mod`发送一个`GET`请求。

​	为了避免从不区分大小写的文件系统中提供服务时出现歧义，`$module`和`$version`元素采用大小写编码，将每个大写字母替换为感叹号后跟相应的小写字母。这允许模块 `example.com/M` 和 `example.com/m` 同时存储在磁盘上，因为前者被编码为 `example.com/!m`。



| Path 路径 | Description 描述 |
| --------- | ---------------- |
|           |                  |

#### `$base/$module/@v/list`

​	以纯文本形式返回给定模块的已知版本列表，每行一个。该列表不应包括伪版本。



#### `$base/$module/@v/$version.info`

​	返回有关模块的特定版本的JSON格式的元数据。响应必须是与下面的GO数据结构体相对应的JSON对象：

```go 
type Info struct {
    Version string    // version string
    Time    time.Time // commit time
}
```

​	`Version`字段是必须的，并且必须包含一个有效的、[canonical version（经典的版本）](../Glossary#canonical-version)（请参见[Versions（版本）](../ModulesPackagesAndVersions#versions)）。请求路径中的`$version`不需要是相同的版本，甚至不需要是有效的版本；此端点可以用来查找分支名称或修订标识符的版本。但是，如果`$version`是一个经典的版本，其主版本与`$module`兼容，那么成功响应中的 `Version` 字段必须是相同的。

​	`Time`字段是可选的。如果存在，它必须是一个RFC 3339格式的字符串。它表示版本创建的时间。未来可能会增加更多的字段，因此保留其他的名字。



#### `$base/$module/@v/$version.mod`

​	返回模块的特定版本的`go.mod`文件。如果该模块在请求的版本中没有`go.mod`文件，则必须返回一个仅包含请求模块路径的`module`语句的文件。否则，必须返回原始的、未经修改的`go.mod`文件。



#### `$base/$module/@v/$version.zip`

​	返回包含模块特定版本内容的zip文件。有关此zip文件必须如何格式化的详细信息，请参阅[Module zip files（模块zip文件）](../ModuleZipFiles)。



#### `$base/$module/@latest`

​	以与`$base/$module/@v/$version.info`相同格式返回有关模块的最新已知版本的JSON格式元数据。如果`$base/$module/@v/list`为空或者列出的版本不合适，则最新版本应该是`go`命令应该使用的模块版本。此端点是可选的，模块代理不需要实现它。



​	在解析模块的最新版本时，`go`命令将请求`$base/$module/@v/list`，如果没有找到合适的版本，则请求`$base/$module/@latest`。`go`命令按顺序优先选择：语义上最高的发布版本，语义上最高的预发布版本，以及时间上最新的伪版本。在Go 1.12和更早的版本中，`go`命令认为`$base/$module/@v/list`中的伪版本是预发布版本，但从Go 1.13开始不再是这样了。

​	模块代理必须始终为`$base/$module/$version.mod`和`$base/$module/$version.zip`查询的成功响应提供相同的内容。该内容使用`go.sum` 文件进行[cryptographically authenticated（加密身份验证）](../AuthenticatingModules)，默认情况下使用[checksum database（校验和数据库）](../AuthenticatingModules#checksum-database)。

​	`go`命令将它从模块代理下载的大部分内容缓存在`$GOPATH/pkg/mod/cache/download`的模块缓存中。即使是直接从版本控制系统中下载，`go`命令也会合成显式的`info`、`mod`和`zip`文件，并将它们存储在此目录中，就像它直接从代理那里下载一样。缓存的布局与代理的URL空间相同，因此将`$GOPATH/pkg/mod/cache/download`服务于（或复制到）`https://example.com/proxy`，将可以让用户通过设置`GOPROXY`为`https://example.com/proxy`，来访问缓存的模块版本。

### Communicating with proxies 与代理通信

​	`go`命令可以从[module proxy（模块代理）](../Glossary#module-proxy)处下载模块的源代码和元数据。`GOPROXY`环境变量可以用来配置`go`命令可以连接哪些代理，以及它是否可以直接与[version control systems（版本控制系统）](../VersionControlSystems)通信。下载的模块数据被保存在[module cache（模块缓存）](../Glossary#module-cache)中。`go`命令只有在需要缓存中没有的信息时才会联系代理。

​	[GOPROXY 协议](#goproxy-protocol)部分描述了可能被发送到`GOPROXY`服务器的请求。不过，了解`go`命令何时发出这些请求也很有帮助。例如，`go build`遵循以下步骤：

- 通过读取`go.mod`文件并执行[最小版本选择（MVS）](../MVS)来计算[build list（构建列表）](../Glossary#build-list)。
- 读取命令行上命名的包和及其导入的包。
- 如果构建列表中的任何模块都没有提供某个包，则寻找提供该包的模块。将最新版本的模块需求添加到`go.mod`，然后重新开始（这些步骤）。
- 在加载完所有内容之后构建包。

When the `go` command computes the build list, it loads the `go.mod` file for each module in the [module graph](https://go.dev/ref/mod#glos-module-graph). If a `go.mod` file is not in the cache, the `go` command will download it from the proxy using a `$module/@v/$version.mod` request (where `$module` is the module path and `$version` is the version). These requests can be tested with a tool like `curl`. For example, the command below downloads the `go.mod` file for `golang.org/x/mod` at version `v0.2.0`:

​	当 `go` 命令计算构建列表时，它为[module graph（模块图）](../Glossary#module-graph)中的每个模块加载 `go.mod` 文件。如果`go.mod`文件不在缓存中，`go`命令将使用`$module/@v/$version.mod`请求（其中`$module`是模块路径，`$version`是版本）从代理中下载它。这些请求可以用`curl`这样的工具来测试。例如，下面的命令下载版本为`v0.2.0`的`golang.org/x/mod`的`go.mod`文件。

```shell
$ curl https://proxy.golang.org/golang.org/x/mod/@v/v0.2.0.mod

module golang.org/x/mod

go 1.12

require (
    golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
    golang.org/x/tools v0.0.0-20191119224855-298f0cb1881e
    golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898
)
```

​	为了加载包，`go`命令需要提供该包的模块的源代码。模块源代码以`.zip`文件的形式发布，这些文件被解压缩到模块缓存中。如果模块`.zip`不在缓存中，`go`命令将使用`$module/@v/$version.zip`请求下载它。

```shell
$ curl -O https://proxy.golang.org/golang.org/x/mod/@v/v0.2.0.zip
$ unzip -l v0.2.0.zip | head
Archive:  v0.2.0.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
     1479  00-00-1980 00:00   golang.org/x/mod@v0.2.0/LICENSE
     1303  00-00-1980 00:00   golang.org/x/mod@v0.2.0/PATENTS
      559  00-00-1980 00:00   golang.org/x/mod@v0.2.0/README
       21  00-00-1980 00:00   golang.org/x/mod@v0.2.0/codereview.cfg
      214  00-00-1980 00:00   golang.org/x/mod@v0.2.0/go.mod
     1476  00-00-1980 00:00   golang.org/x/mod@v0.2.0/go.sum
     5224  00-00-1980 00:00   golang.org/x/mod@v0.2.0/gosumcheck/main.go
```

注意，`.mod`和`.zip`的请求是分开的，尽管`go.mod`文件通常包含在`.zip`文件中。`go`命令可能需要为许多不同的模块下载`go.mod`文件，而`.mod`文件要比`.zip`文件小得多。此外，如果一个Go项目没有`go.mod`文件，代理将提供一个仅包含[module 指令](../gomodFiles#module-directive)的合成`go.mod`文件。合成的`go.mod`文件是由`go`命令从[version control system（版本控制系统）](../VersionControlSystems)下载时生成的。

​	如果`go`命令需要加载一个构建列表中任何模块都没有提供的包，它将尝试查找一个提供该包的新模块。[Resolving a package to a module（将包解析为模块）](../ModulesPackagesAndVersions#resolving-a-package-to-a-module)一节描述了这个过程。总之，`go`命令会请求每个可能包含该包的模块路径的最新版本信息。例如，对于包`golang.org/x/net/html`，`go`命令会试图查找`golang.org/x/net/html`、`golang.org/x/net`、`golang.org/x/`和`golang.org`等模块的最新版本。只有`golang.org/x/net`实际存在并提供该包，因此`go`命令使用该模块的最新版本。如果有多个模块提供该包，`go`命令将使用路径最长的模块。

​	当`go`命令请求某个模块的最新版本时，它首先发送一个`$module/@v/list`的请求。如果列表是空的或者没有一个返回的版本可以使用，它将发送对`$module/@latest`的请求。一旦选择了一个版本，`go`命令就会发送对`$module/@v/$version.info`的元数据请求。然后它可能会发送`$module/@v/$version.mod`和`$module/@v/$version.zip`请求来加载`go.mod`文件和源代码。

```shell
$ curl https://proxy.golang.org/golang.org/x/mod/@v/list
v0.1.0
v0.2.0

$ curl https://proxy.golang.org/golang.org/x/mod/@v/v0.2.0.info
{"Version":"v0.2.0","Time":"2020-01-02T17:33:45Z"}
```

​	在下载一个`.mod`或`.zip`文件后，`go`命令会计算一个加密哈希值，并检查它是否与主模块的`go.sum`文件中的哈希值相匹配。如果哈希值不在`go.sum`中，默认情况下，`go`命令会从[checksum database（校验数据库）](../AuthenticatingModules#checksum-database)中检索它。如果计算出的哈希值不匹配，`go`命令会报告一个安全错误，并且不会将该文件安装到模块缓存中。`GOPRIVATE`和`GONOSUMDB`[环境变量](../EnvironmentVariables)可以用来禁止对特定模块的校验数据库的请求。`GOSUMDB`环境变量也可以被设置为`off`，以完全禁止对校验数据库的请求。更多信息请参见[Authenticating modules（验证模块）](../AuthenticatingModules)。请注意，为`.info`请求返回的版本列表和版本元数据不经过身份验证，并且可能会随着时间的推移而改变。

### Serving modules directly from a proxy 直接从代理向模块提供服务

​	大多数模块都是从版本控制存储库中开发和提供的。在[direct mode（直接模式）](../Glossary#direct-mode)下，`go`命令用版本控制工具下载这样的模块（见[Version control systems（版本控制系统）](../VersionControlSystems)）。还可以直接从模块代理提供模块。这对那些希望在不暴露其版本控制服务器的情况下提供模块服务的组织以及使用`go`命令不支持的版本控制工具的组织来说非常有用。

​	当`go`命令以直接模式下载模块时，它首先根据模块路径用HTTP GET请求查找模块服务器的URL。它在HTML响应中查找一个名为`go-import`的`<meta>`标签。该标签的内容必须包含[repository root path（存储库根路径）](../Glossary#repository-root-path)、版本控制系统和 URL，并以空格隔开。详见[Finding a repository for a module path（查找模块路径的存储库）](../VersionControlSystems#finding-a-repository-for-a-module-path)。

​	如果版本控制系统是`mod`，`go`命令使用[GOPROXY 协议](#goproxy-protocol)从给定的URL下载模块。

​	例如，假设`go`命令试图下载版本为`v1.0.0`的模块`example.com/gopher`。它向`https://example.com/gopher?go-get=1`发送请求。服务器使用包含以下标签的HTML文档进行响应：

```html
<meta name="go-import" content="example.com/gopher mod https://modproxy.example.com">
```

​	根据这个响应，`go`命令通过发送`https://modproxy.example.com/example.com/gopher/@v/v1.0.0.info`、`v1.0.0.mod`和`v1.0.0.zip`的请求来下载该模块。

> 注意，在`GOPATH`模式下，不能用`go get`下载从代理处直接提供的模块。