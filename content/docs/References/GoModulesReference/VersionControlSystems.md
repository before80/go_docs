+++
title = "版本控制系统"
date = 2023-05-17T09:59:21+08:00
weight = 10
description = ""
isCJKLanguage = true
draft = false
+++
## Version control systems 版本控制系统

> 原文：[https://go.dev/ref/mod#vcs](https://go.dev/ref/mod#vcs)

​	`go`命令可以直接从版本控制储存库中下载模块的源代码和元数据。从[proxy（代理）](../ModuleProxies#communicating-with-proxies)下载模块通常更快，但如果代理不可用，或者代理无法访问模块的存储库（对于私有存储库经常如此），则需要直接连接到存储库。支持`Git`、`Subversion`、`Mercurial`、`Bazaar`和`Fossil`。版本控制工具必须安装在`PATH`的一个目录中，以便`go`命令使用它。

​	要从源码库而不是代理下载特定的模块，可以设置`GOPRIVATE`或`GONOPROXY`环境变量。要配置`go`命令直接从源码库下载所有模块，请将`GOPROXY`设置为`direct`。更多信息请参见[Environment variables（环境变量）](../EnvironmentVariables)。

### Finding a repository for a module path 查找模块路径的存储库

​	当`go`命令以`direct`模式下载模块时，它首先要定位包含该模块的存储库。

​	如果模块路径的末尾有一个`VCS`限定词（`.bzr`， `.fossil`， `.git`， `.hg`， `.svn`中的一个），`go`命令将使用该路径限定词之前的所有内容作为存储库的URL。例如，对于模块`example.com/foo.git/bar`，`go`命令使用`git`下载`example.com/foo.git`的存储库，期望在`bar`子目录下找到该模块。`go`命令将根据版本控制工具支持的协议来猜测要使用的协议。

​	如果模块路径没有限定词，`go`命令会向模块路径派生的URL发送带有`?go-get=1`查询字符串的HTTP `GET`请求。例如，对于模块`golang.org/x/mod`，`go`命令可能发送以下请求：

```
https://golang.org/x/mod?go-get=1 (preferred)
http://golang.org/x/mod?go-get=1  (fallback, only with GOINSECURE)
```

​	`go`命令跟随重定向，但会忽略响应状态码，因此服务器可能会以404或任何其他错误状态来响应。`GOINSECURE`环境变量可以被设置为允许回退并重定向到特定模块的未加密的HTTP。

​	服务器必须使用一个HTML文档来响应，该文档的`<head>`中包含一个`<meta>`标签。`<meta>`标签应该出现在文档的早期，以避免混淆`go`命令的受限解析器。特别是，它应该出现在任何原始JavaScript或CSS之前。`<meta>`标签必须具有以下形式：

```
<meta name="go-import" content="root-path vcs repo-url">
```

​	`root-path` 是存储库的根路径，即模块路径中与存储库根目录相对应的部分。它必须是一个前缀或与请求的模块路径完全匹配。如果不是完全匹配，则会对前缀进行另一次请求，以验证`<meta>`标签是否匹配。

​	`vcs`是版本控制系统。它必须是下表中所列的工具之一，或者是关键字`mod`，它指示`go`命令使用[GOPROXY 协议](#goproxy-protocol)从给定的URL下载模块。有关详细信息，请参阅 [直接从代理向模块提供服务](#serving-modules-directly-from-a-proxy)。

​	`repo-url`是存储库的URL。如果URL不包含方案（要么是因为模块路径有一个`VCS`限定符，要么是因为`<meta>`标签缺少一个方案），`go`命令将尝试版本控制系统支持的每一个协议。例如，对于 `Git`，`go` 命令将尝试 `https://` 然后是 `git+ssh://`。不安全的协议（如 `http://` 和 `git://`）只有在模块路径被 `GOINSECURE` 环境变量匹配的情况下才能使用。

| Name       | Command  | GOVCS default      | Secure schemes              |
| ---------- | -------- | ------------------ | --------------------------- |
| Bazaar     | `bzr`    | Private only       | `https`， `bzr+ssh`         |
| Fossil     | `fossil` | Private only       | `https`                     |
| Git        | `git`    | Public and private | `https`， `git+ssh`， `ssh` |
| Mercurial  | `hg`     | Public and private | `https`， `ssh`             |
| Subversion | `svn`    | Private only       | `https`， `svn+ssh`         |

​	作为一个例子，再次考虑`golang.org/x/mod`。`go`命令向`https://golang.org/x/mod?go-get=1` 发送一个请求。服务器响应的是一个包含标签的HTML文档：

```
<meta name="go-import" content="golang.org/x/mod git https://go.googlesource.com/mod">
```

​	从这个响应来看，`go` 命令将使用远程 URL `https://go.googlesource.com/mod` 的 `Git` 存储库。

​	`GitHub` 和其他流行的托管服务会响应所有存储库的 `?go-get=1` 查询，所以通常在这些站点托管的模块不需要进行服务器配置。

​	找到存储库的URL后，`go`命令会将存储库克隆到模块缓存中。一般来说，`go`命令尝试避免从存储库中获取不需要的数据。但是，实际使用的命令因版本控制系统而异，并可能随时间而改变。对于 `Git` 来说，`go` 命令可以在不下载提交的情况下列出大多数可用的版本。它通常会在不下载祖先提交的情况下获取提交，但这样做有时是必要的。

### Mapping versions to commits 将版本映射到提交

`go` 命令可以检出存储库中特定[canonical version（经典版本）](../Glossary#canonical-version)的模块，比如 `v1.2.3`, `v2.4.0-beta`, 或者 `v3.0.0+incompatible`。每个模块版本在存储库中都应该有一个语义上的版本标签，表明哪个版本应该被检出。

​	如果在存储库根目录或根目录的主版本子目录中定义了模块，则每个版本标签名都等于相应的版本。例如，模块 `golang.org/x/text` 被定义在其存储库的根目录中，因此版本 `v0.3.2` 在该存储库中的标签为 `v0.3.2`。这对大多数模块来说都是如此。

​	如果模块被定义在存储库的子目录中，也就是说，模块路径中的[module subdirectory（模块子目录）](../Glossary#module-subdirectory)部分不是空的，那么每个标签名必须以模块子目录为前缀，后跟斜线。例如，模块`golang.org/x/tools/gopls`被定义在根路径为`golang.org/x/tools`的存储库的`gopls`子目录中。该模块的`v0.4.0`版本必须在该存储库中具有名为`gopls/v0.4.0`的标签。

​	语义版本标签的主版本号必须与模块路径的主版本后缀（如果有的话）一致。例如，标签`v1.0.0`可能属于模块`example.com/mod`，但不属于`example.com/mod/v2`，后者会有`v2.0.0`这样的标签。

​	如果没有`go.mod`文件，并且模块在存储库根目录中，那么主版本为`v2`或更高的标签可能属于没有主版本后缀的模块。这种版本用后缀`+incompatible`来表示。版本标签本身不能有这个后缀。参见[Compatibility with non-module repositories（与非模块存储库的兼容性）](../CompatibilityWithNon-moduleRepositories)。

​	一旦标签被创建，它就不应该被删除或改变为不同的版本。版本经过[身份验证](../AuthenticatingModules)，以确保安全、可重复的构建。如果标签被修改，客户端在下载时可能会看到安全错误。即使标签被删除，其内容仍可在[module proxies（模块代理）](../Glossary#module-proxy)上使用。

### Mapping pseudo-versions to commits 将伪版本映射到提交

​	`go`命令可以在存储库中检出特定修订版的模块，该修订版编码为[pseudo-version（伪版本）](../Glossary#pseudo-version)，如`v1.3.2-0.20191109021931-daa7c04131f5`。

​	伪版本的最后12个字符（上例中的`daa7c04131f5`）表示要检出存储库中的一个修订版。它的含义取决于版本控制系统。对于`Git`和`Mercurial`，这是一个提交哈希值的前缀。对于`Subversion`，这是一个以零填充的修订号。

​	在检出一个提交之前，`go` 命令会验证时间戳（上面的 `20191109021931`）是否与提交日期相符。它还会验证基本版本（`v1.3.1`，即上例中`v1.3.2`之前的版本）是否与提交的祖先的语义版本标签相对应。这些检查确保模块作者能够完全控制伪版本与其他发布版本的比较。

​	更多信息请参见[Pseudo-versions（伪版本）](../ModulesPackagesAndVersions#pseudo-versions)。

### Mapping branches and commits to versions 将分支和提交映射到版本

​	可以使用[version query（版本查询）](../Module-awareCommands#version-queries)在特定分支、标签或修订版检出模块。

```shell
go get example.com/mod@master
```

​	`go` 命令将这些名称转换为可用于[最小版本选择（MVS）](../MVS)的[canonical versions（经典版本）](../Glossary#canonical-version)。MVS 依赖于对版本进行明确排序的能力。分支名称和修订版不能随着时间的推移可靠地进行比较，因为它们依赖于可能会更改的存储库结构。

​	如果一个修订版被标记了一个或多个语义版本标签，如`v1.2.3`，那么将使用最高有效版本的标签。`go`命令只考虑可能属于目标模块的语义版本标签；例如，对于`example.com/mod/v2`来说，`v1.5.2`标签不会被考虑，因为其主版本与模块路径的后缀不匹配。

​	如果一个修订版没有被贴上有效的语义版本标签，`go`命令将生成一个伪版本。如果修订版具有具有有效语义版本标签的祖先版本，那么最高的祖先版本将被用作伪版本基础。请参阅[Pseudo-versions（伪版本）](../ModulesPackagesAndVersions#pseudo-versions)。

### Module directories within a repository 存储库中的模块目录

​	一旦模块的存储库在特定的修订版中被检出，`go`命令必须找到包含该模块的`go.mod`文件的目录（模块的根目录）。

​	回顾一下，[module path（模块路径）](../ModulesPackagesAndVersions#module-paths)由三部分组成：存储库根路径（对应于存储库根目录）、模块子目录和主版本后缀（仅适用于以`v2`或更高版本发布的模块）。

​	对于大多数模块，模块路径等于存储库根路径，因此模块的根目录就是存储库的根目录。

​	模块有时会被定义在存储库的子目录下。这通常适用于具有多个组件的大型存储库，这些组件需要独立发布和版本化。这样的模块有望在子目录中找到，该目录与存储库根目录之后的模块路径部分相匹配。例如，假设模块 `example.com/monorepo/foo/bar` 位于根路径为 `example.com/monorepo`的存储库的中。它的`go.mod`文件必须位于`foo/bar`子目录中。

​	如果模块在主版本`v2`或更高版本发布，则其路径必须有一个[major version suffix（主版本后缀）](../ModulesPackagesAndVersions#major-version-suffixes)。带有主版本后缀的模块可以定义在两个子目录中的一个：一个带有后缀，另一个没有。例如，假设上面模块的一个新版本以`example.com/monorepo/foo/bar/v2`的路径发布。其`go.mod`文件可能位于`foo/bar`或`foo/bar/v2`中。

​	带有主版本后缀的子目录是主版本的子目录。它们可用于在单个分支上开发模块的多个主版本。当在不同的分支上进行多个主要版本的开发时，这可能是不必要的。然而，主版本的子目录有一个重要的特性：在`GOPATH`模式下，包的导入路径与`GOPATH/src`下的目录完全匹配。`go`命令在`GOPATH`模式下提供了最低限度的模块兼容性（参见[Compatibility with non-module repositories（与非模块存储库的兼容性）](../CompatibilityWithNon-moduleRepositories)），因此主版本子目录对于与`GOPATH`模式下构建的项目的兼容性来说并不总是必要的。不过不支持最低限度的模块兼容性的旧工具可能会有问题。

​	一旦`go`命令找到了模块根目录，它就会创建该目录内容的`.zip`文件，然后将该`.zip`文件解压缩到模块缓存中。关于`.zip`文件中可以包含哪些文件的细节，请参见[File path and size constraints（文件路径和大小限制）](../ModuleZipFiles#file-path-and-size-constraints)。`.zip`文件的内容在提取到模块缓存之前是经过[验证](../AuthenticatingModules)的，就像从代理下载`.zip`文件一样。

​	模块压缩文件不包括`vendor`目录的内容或任何嵌套模块（包含`go.mod`文件的子目录）。这意味着模块必须注意不要引用其目录外或其他模块中的文件。例如，[//go:embed](https://pkg.go.dev/embed#hdr-Directives)模式不能匹配嵌套模块中的文件。在文件不应包含在模块中的情况下，这种行为可以作为一种有用的变通方法。例如，如果存储库有大文件被检入`testdata`目录中，模块作者可以在`testdata`中添加一个空的`go.mod`文件，这样他们的用户就不需要下载这些文件。当然，这可能会减少用户测试其依赖项的覆盖率。

### Special case for LICENSE files - 许可证文件的特殊情况

​	当`go`命令为不在存储库根目录下的模块创建`.zip`文件时，如果该模块的根目录（与`go.mod`并列）中没有名为`LICENSE`的文件，那么`go`命令将从存储库根目录下复制名为`LICENSE`的文件（如果该文件存在于同一修订版中）。

​	这种特殊情况允许相同的 `LICENSE` 文件应用于存储库中的所有模块。这仅适用于专门命名为 `LICENSE` 的文件，而没有像 `.txt` 这样的扩展名。遗憾的是，在不破坏现有模块的加密和的情况下，无法对其进行扩展；请参见[Authenticating modules（认证模块）](../AuthenticatingModules)。其他工具和网站如 `pkg.go.dev` 可能会识别其他名称的文件。

​	还请注意，在创建模块`.zip`文件时，`go`命令不包括符号链接；请参见[File path and size constraints（文件路径和大小限制）](../ModuleZipFiles#file-path-and-size-constraints)。因此，如果存储库的根目录中没有`LICENSE`文件，作者可以在子目录中定义的模块中创建许可证文件的副本，以确保这些文件包含在模块`.zip`文件中。

### Controlling version control tools with GOVCS - 使用GOVCS控制版本控制工具

​	`go`命令能够使用`git`等版本控制命令下载模块，这对于去中心化的包生态系统至关重要，在这个系统中，代码可以从任何服务器导入。如果恶意服务器找到方法使调用的版本控制命令运行非预期的代码，这也是一个潜在的安全问题。

​	为了平衡功能和安全问题，`go`命令默认只使用`git`和`hg`从公共服务器下载代码。它将使用任何[已知的版本控制系统](#finding-a-repository-for-a-module-path)从私有服务器下载代码，私有服务器定义为那些托管与`GOPRIVATE`[环境变量](../EnvironmentVariables)匹配的包的服务器。之所以只允许使用`Git`和`Mercurial`，是因为这两个系统最关注作为不可信服务器的客户端运行的问题。相比之下，`Bazaar`、`Fossil`和`Subversion`主要用于可信的、经过验证的环境中，并没有被作为攻击面那样受到仔细检查。

​	版本控制命令的限制仅适用于使用直接版本控制访问下载代码的情况。当从代理下载模块时，`go`命令改用[GOPROXY 协议](../ModuleProxies#goproxy-protocol)，这是始终允许的。默认情况下，`go`命令使用Go模块镜像（`proxy.golang.org`）来下载公共模块，只有在私有模块或镜像拒绝为公共包提供服务（通常是出于法律原因）才会回退到版本控制。因此，客户端仍然可以默认访问从`Bazaar`、`Fossil`或`Subversion`存储库提供的公共代码，因为这些下载使用Go模块镜像，它承担了使用自定义沙箱运行版本控制命令的安全风险。

​	`GOVCS`变量可以用来更改特定模块所允许的版本控制系统。`GOVCS`变量适用于在模块感知模式和`GOPATH`模式下构建包。当使用模块时，模式与模块路径匹配。当使用 `GOPATH` 时，模式与版本控制库根目录对应的导入路径相匹配。

​	`GOVCS`变量的一般形式是一个用逗号分隔的`pattern:vcslist`规则的列表。`pattern` 是一个[glob pattern](https://go.dev/pkg/path#Match)，必须与模块或导入路径的一个或多个前导元素相匹配。`vcslist`是一个管道分隔的允许使用的版本控制命令的列表，或`all`允许使用任何已知的命令，或`off`不允许使用。请注意，如果模块与`vcslist` `off`时的模式相匹配，如果原服务器使用`mod`方案，该模块仍然可以被下载，该方案指示`go`命令使用[GOPROXY 协议](../ModuleProxies#goproxy-protocol)下载该模块。 应用列表中最早的匹配模式，即使后面的模式也可能匹配。

​	例如，请考虑：

```
GOVCS=github.com:git,evil.com:off,*:git|hg
```

​	在此设置下，模块或导入路径以`github.com/`开头的代码只能使用`git`；`evil.com`上的路径不能使用任何版本控制命令，而所有其他路径（`*`匹配所有）只能使用`git`或`hg`。

​	特殊模式`public`和`private`匹配公共和私人模块或导入路径。如果路径与`GOPRIVATE`变量相匹配，那么它就是私有路径；否则就是公共路径。

​	如果`GOVCS`变量中没有规则与某个特定的模块或导入路径相匹配，`go`命令会应用其默认规则，该规则现在可以用`GOVCS`标记法概括为`public:git|hg`，`private:all`。

​	要允许对任何包不受限制地使用任何版本控制系统，请使用：

```
GOVCS=*:all
```

​	要禁用所有版本控制的使用，使用：

```
GOVCS=*:off
```

​	[go env -w 命令]({{< ref "/cmd/go#print-go-environment-information">}})可以用来设置 `GOVCS` 变量，以便今后调用 `go` 命令。

​	`GOVCS` 是在 Go 1.16 中引入的。早期版本的Go可以对任何模块使用任何已知的版本控制工具。

