+++
title = "模块、包和版本"
date = 2023-05-17T09:59:21+08:00
weight = 2
description = ""
isCJKLanguage = true
draft = false
+++
## Modules, packages, and versions 模块、包和版本

> 原文：https://go.dev/ref/mod#modules-overview

​	模块是发布、版本化和一起分发的包的集合。模块可以直接从版本控制存储库或模块代理服务器上下载。

​	模块由[模块路径](#module-paths)来识别，该路径在 [go.mod 文件](../gomodFiles)中声明，同时还有关于该模块的依赖项信息。模块根目录是包含`go.mod`文件的目录。主模块是包含调用`go`命令的目录的模块。

​	模块内的每个包是同一目录下的源文件的集合，这些文件被编译在一起。包路径是与包含包（相对于模块根目录）的子目录相连接的模块路径。例如，模块 "`golang.org/x/net`"在 "`html`"目录下包含一个包。那个包的路径是 "`golang.org/x/net/html`"。

### Module paths 模块路径

​	模块路径是模块的规范名称，在模块的 [go.mod 文件](../gomodFiles)中用[module指令](../gomodFiles#module-directive)声明。模块的路径是该模块中包路径的前缀。

​	模块路径应该同时描述模块的作用和查找模块的位置。通常，模块路径由存储库根路径、存储库中的目录（通常为空）和主版本后缀（仅适用于主版本2或更高版本）组成。

- 存储库根路径是模块路径中与开发模块的版本控制存储库的根目录相对应的部分。大多数模块都定义在其存储库的根目录下，所以这通常是整个路径。例如，`golang.org/x/net`是同名模块的存储库根目录。关于 `go` 命令如何使用从模块路径派生的 HTTP 请求来定位存储库的信息，请参见[Finding a repository for a module path（寻找模块路径的存储库）](../VersionControlSystems#finding-a-repository-for-a-module-path)。
- 如果模块没有定义在存储库的根目录中，模块子目录是模块路径中命名目录的一部分，不包括主版本后缀。这也可以作为语义版本标签的前缀。例如，模块`golang.org/x/tools/gopls`位于根目录`golang.org/x/tools`的`gopls`子目录中，所以它的模块子目录是`gopls`。参见[Mapping versions to commits](../VersionControlSystems#mapping-versions-to-commits)和[Module directories within a repository](../VersionControlSystems#module-directories-within-a-repository)。
- 如果模块是在主版本2或更高版本发布的，模块路径必须以主版本后缀结尾，如`/v2`。这可能是也可能不是子目录名称的一部分。例如，路径为 `golang.org/x/repo/sub/v2` 的模块可能在 `golang.org/x/repo` 仓库的 `/sub` 或 `/sub/v2` 子目录下。

​	如果模块可能被其他模块所依赖，就必须遵循这些规则，以便`go`命令能够找到并下载该模块。对于模块路径中允许的字符也有一些[lexical restrictions（词法限制）](../gomodFiles#module-paths-and-versions)。

### Versions 版本

​	一个版本标识了模块的不可改变的快照，它可以是一个发布版或预发布版。每个版本以字母`v`开头，后面是语义版本。关于如何对版本进行格式化、解释和比较的细节，请参见[Semantic Versioning 2.0.0（语义化版本2.0.0）](https://semver.org/lang/zh-CN/)。

​	简而言之，语义版本由三个非负整数（从左到右分别为：主版本号，次版本号，修订版本号）组成，中间用点分开。修订版本号后面可以有一个以连字符（`-`或`+`）开头的可选预发布字符串。预发布字符串或补丁版本后面可以有一个以加号开头的构建元数据字符串。例如，`v0.0.0`、`v1.12.134`、`v8.0.5-pre`和`v2.0.9+meta`是有效的版本。

​	版本的每一部分都指示该版本是否稳定，以及是否与以前的版本兼容。

- 在对模块的公共接口或文档功能进行了向后不兼容的更改（例如，在删除了一个包）之后，[major version（主版本号）](../Glossary#major-version)必须被递增，次版本号和修订版本号必须被设置为零。

- 在进行向后兼容的更改（例如，增加了一个新的功能）之后，[minor version（次版本号）](../Glossary#minor-version)必须被递增，修订版本号必须被设置为零。

- [patch version（修订版本号）](../Glossary#patch-version)必须在不影响模块的公共接口的更改（例如错误修复或优化）之后被递增。

- [pre-release （预发布）](../Glossary#pre-release-version)后缀表示版本是预发布版本。预发布版本在相应的发布版本之前排序。例如，`v1.2.3-pre`排在`v1.2.3`之前。

- 在比较版本时，build metadata的后缀被忽略。带有build metadata的标签在版本控制存储库中被忽略，但build metadata在 [go.mod 文件](../gomodFiles)中指定的版本中被保留下来。后缀`+incompatible`表示在迁移到模块版本主版本2或更高版本之前发布的版本（参见[与非模块存储库的兼容性](../CompatibilityWithNon-moduleRepositories)）。

  

​	如果一个版本的主版本号是`0`，或者它有一个预发布的后缀，则被认为是不稳定的。不稳定的版本不受兼容性要求的限制。例如，`v0.2.0` 可能与 `v0.1.0` 不兼容，`v1.5.0-beta` 可能与 `v1.5.0` 不兼容。

​	Go可以使用不遵循这些约定的标签、分支或修订版来访问在版本控制系统的模块。然而，在主模块中，`go`命令会自动将不遵循此标准的修订版名称转换为规范的版本号。在这个过程中，`go`命令也会删除build metadata的后缀（除了`+incompatible`）。这可能会产生一个[pseudo-version（伪版本号）](../Glossary#pseudo-version)，一个编码修订版标识符（如Git提交的哈希值）的预先发布的版本号和一个版本控制系统的时间戳。例如，命令`go get golang.org/x/net@daa7c041`将把提交的哈希值`daa7c041`转换为伪版本`v0.0.0-20191109021931-daa7c04131f5`。 在主模块之外需要规范的版本，如果在 [go.mod 文件](../gomodFiles)中出现类似`master`的非规范版本，`go`命令将报告错误。

### Pseudo-versions 伪版本号

​	伪版本号是一个特殊格式的[pre-release version（预发布版本号）](../Glossary#pre-release-version)，它对版本控制存储库中的特定版本信息进行编码。例如，`v0.0.0-20191109021931-daa7c04131f5`就是一个伪版本号。

​	伪版本号可以指那些没有语义版本标签可用的修订版。例如，在开发分支上创建版本标签之前，可以使用它们来测试提交的内容。

每个伪版本都有三个部分：

- 基础版本前缀（`base version prefix`）（`vX.0.0`或`vX.Y.Z-0`），该前缀来自于该修订版之前的语义版本标签，如果没有这样的标签，则是`vX.0.0`。
- 时间戳（`timestamp`）（`yyyymmddhhmmss`），这是修订版创建的UTC时间。在Git中，这是提交的时间，而不是作者时间。
- 修订版标识符（`revision identifier`）（`abcdefabcdef`），它是提交哈希值的一个12个字符的前缀，或者在Subversion中，是一个零填充的修订号。

​	每个伪版本可以是三种形式之一，取决于基础版本（base version）。这些形式保证了伪版本比其基础版本高，但比下一个标记的版本低。

- `vX.0.0-yyyymmddhhmmss-abcdefabcdef`在没有已知基础版本时使用。与所有版本一样，主版本`X`必须与模块的[major version suffix（主版本后缀）](../Glossary#major-version-suffix)相匹配。
- `vX.Y.Z-pre.0.yyyymmddhhmmss-abcdefabcdef`在基础版本是预发布版本（如 `vX.Y.Z-pre`）时使用。
- `vX.Y.(Z+1)-0.yyyymmddhhmmss-abcdefabcdef`在基础版本是发布版本（如`vX.Y.Z`）时使用。例如，如果基础版本是`v1.2.3`，那么伪版本可能是`v1.2.4-0.20191109021931-daa7c04131f5`。

​	通过使用不同的基础版本，多个伪版本可以引用相同的提交。这发生在写完伪版本后，低版本被标记的时候。

​	这些形式给了伪版本两个有用的属性：

- 具有已知基础版本的伪版本排序高于这些版本，但低于其他预发布的后续版本。
- 具有相同基础版本前缀的伪版本按时间顺序排序。

​	`go`命令进行了一些检查，以确保模块作者能够控制伪版本与其他版本的比较，并确保伪版本引用的修订版是实际属于模块提交历史的一部分。

- 如果指定了基础版本（base version），就必须有相应的语义版本标签，该标签是伪版本所描述的修订版的祖先。这可以防止开发者使用伪版本来绕过[minimal version selection（最小版本）](../Glossary#minimal-version-selection)的选择，该伪版本比所有标签的版本（比如`v1.999.999-9999999999-daa7c04131f5`）都要高。
- 时间戳（timestamp）必须与修订版的时间戳匹配。这可以防止攻击者用无限数量的其他相同伪版本来 flooding [module proxies （模块代理）](../Glossary#module-proxy)。这也可以防止模块使用者更改版本的相对顺序。
- 修订版必须是模块存储库的一个分支或标签的祖先。这可以防止攻击者引用未经批准的更改或拉取请求。

​	伪版本从不需要手工输入。许多命令接受提交哈希值或分支名称，并将其自动翻译成伪版本（或标签版本（如果可用））。比如说：

```
go get example.com/mod@master
go list -m -json example.com/mod@abcd1234
```

### Major version suffixes 主版本的后缀

​	从主版本2开始，模块路径必须有一个主版本后缀，如`/v2`，与主版本相匹配。例如，如果模块的路径`example.com/mod`是`v1.0.0`版本，那么它的路径`example.com/mod/v2`必须是`v2.0.0`版本。

​	主版本后缀实现了[import compatibility rule（导入兼容性规则）](https://research.swtch.com/vgo-import)：

> 如果旧包和新包有相同的导入路径，则新包必须向后兼容旧包。

​	根据定义，模块的新的主版本中的包不向后兼容之前的主版本中的相应包。因此，从`v2`开始，包需要新的导入路径。这是通过在模块路径上添加一个主版本的后缀来实现的。由于模块路径是模块内每个包的导入路径的前缀，在模块路径上添加主版本后缀为每个不兼容的版本提供了一个不同的导入路径。

​	主版本后缀在主版本`v0`或`v1`时是不允许的。没有必要在`v0`和`v1`之间改变模块路径，因为`v0`版本是不稳定的，没有兼容性保证。此外，对于大多数模块来说，`v1`是向后兼容最后一个`v0`版本的；`v1`版本是对兼容性的一种承诺，而不是表明与`v0`相比有不兼容的变化。

​	作为一种特殊情况，以`gopkg.in/`开始的模块路径必须始终有一个主版本后缀，即使是`v0`和`v1`。 后缀必须以**点**开始，而不是斜线（例如，`gopkg.in/yaml.v2`）。

​	主版本后缀允许模块的多个主版本共存于同一个构建中。这可能是由于[diamond dependency problem（钻石依赖性问题）](https://research.swtch.com/vgo-import#dependency_story)而需要的。通常，如果一个模块在两个不同的版本中被传递依赖关系所需要，则将使用更高的版本。然而，如果这两个版本不兼容，那么这两个版本都不能满足所有的客户。由于不兼容的版本必须有不同的主版本号，因此由于主版本后缀，它们也必须具有不同的模块路径。这就解决了冲突：具有不同后缀的模块被视为独立的模块，它们的包 —— 即使是相对于它们的模块根的同一子目录下的包 —— 也是不同的。

​	许多Go项目在迁移到模块之前（也许在模块被引入之前）发布了`v2`或更高的版本，但没有使用主版本后缀。这些版本使用`+incompatible`的构建标签进行注释（例如，`v2.0.0+incompatible`）。更多信息请参见[Compatibility with non-module repositories（与非模块存储库的兼容性）](../CompatibilityWithNon-moduleRepositories)。

### Resolving a package to a module 将一个包解析为模块

​	当 `go` 命令使用[package path（包路径）](../Glossary#package-path)加载包时，它需要确定哪个模块提供该包。

​	`go`命令首先在[build list（构建列表）](../Glossary#build-list)中搜索路径为包路径前缀的模块。例如，如果包 `example.com/a/b` 被导入，并且模块 `example.com/a` 在构建列表中，`go` 命令将在目录 `b` 中检查 `example.com/a` 是否包含该包。在一个目录中至少要有一个扩展名为 `.go` 的文件，才能被视为一个包。[Build constraints（构建限制）](https://pkg.go.dev/go/build#hdr-Build_Constraints)并不适用这个目的。如果在构建列表中正好有模块提供了该包，那么就使用该模块。如果没有模块提供该包，或者有两个或更多的模块提供该包，`go`命令会报告一个错误。`-mod=mod`标志指示`go`命令尝试寻找提供缺失包的新模块，并更新`go.mod`和`go.sum`。[go get](../Module-awareCommands#go-get)和[go mod tidy](../Module-awareCommands#go-mod-tidy)命令会自动完成这项工作。

​	当`go`命令查找一个包路径的新模块时，它会检查`GOPROXY`环境变量，它是一个用逗号分隔的代理URL列表或关键词`direct`或`off`。代理URL表示`go`命令应该使用[GOPROXY 协议](../ModuleProxies#goproxy-protocol)与模块代理联系。`direct`表示`go`命令应该与[一个版本控制系统通信](../VersionControlSystems)。`GOPRIVATE`和`GONOPROXY`[environment variables（环境变量）](../EnvironmentVariables)也可以用来控制这种行为。

​	对于`GOPROXY`列表中的每个条目，`go`命令请求可能提供包（即包路径的每个前缀）的每个模块路径的最新版本。对于每个成功请求的模块路径，`go`命令将下载最新版本的模块，并检查该模块是否包含请求的包。如果一个或多个模块包含所请求的包，则使用路径最长的模块。如果找到一个或多个模块，但没有一个包含所请求的包，则报告错误。如果没有找到模块，`go`命令将尝试`GOPROXY`列表中的下一个条目。如果没有剩余的条目，则报告错误。

​	例如，假设`go`命令正在寻找一个提供包`golang.org/x/net/html`的模块，并且`GOPROXY`被设置为`https://corp.example.com,https://proxy.golang.org`。`go`命令可能会发出以下请求：

- 向  `https://corp.example.com/` （并行）:

  - 请求`golang.org/x/net/html`的最新版本
  - 请求`golang.org/x/net`的最新版本
  - 请求`golang.org/x`的最新版本
  - 请求`golang.org`的最新版本
  
- 如果向 `https://corp.example.com/`的所有请求都以 `404` 或 `410`失败告终，则向  `https://proxy.golang.org/` :

  - 请求`golang.org/x/net/html`的最新版本
  - 请求`golang.org/x/net`的最新版本
  - 请求`golang.org/x`的最新版本
  - 请求`golang.org`的最新版本

​	在找到合适的模块后，`go` 命令将使用新模块的路径和版本向主模块的 [go.mod 文件](../gomodFiles)添加一个新的[requirement（需求）](../gomodFiles#require-directive)。这样可以确保将来加载相同的包时，会使用相同版本的模块。如果解析后的包没有被主模块中的包导入，那么新的需求将会有一个`// indirect`的注释。