+++
title = "与非模块存储库的兼容性"
date = 2023-05-17T09:59:21+08:00
weight = 7
description = ""
isCJKLanguage = true
draft = false
+++
## Compatibility with non-module repositories 与非模块存储库的兼容性

> 原文：[https://go.dev/ref/mod#non-module-compat](https://go.dev/ref/mod#non-module-compat)

​	为了确保从`GOPATH`到模块的平稳过渡，`go`命令可以通过添加`go.mod`文件，从尚未迁移到模块的存储库中以模块感知模式下载和构建包。

​	当 `go` 命令[直接](../VersionControlSystems)从存储库中下载一个给定版本的模块时，它会查找模块路径的存储库 URL，将该版本映射到存储库中的一个修订版，然后提取该修订版的存储库存档。如果模块的路径等于存储库的根路径，且存储库的根目录不包含`go.mod`文件，那么`go`命令会在模块缓存中合成一个`go.mod`文件，其中包含一个[module 指令](../gomodFiles#module-directive)，而不包含其他内容。由于合成的`go.mod`文件不包含其依赖项的[`require`指令](../gomodFiles#require-directive)，依赖这些模块的其他模块可能需要额外的`require`指令（带有`// indirect`），以确保每个依赖项在每次构建时都以相同的版本被获取。

​	当`go`命令从[proxy（代理）](../ModuleProxies#communicating-with-proxies)下载模块时，它将`go.mod`文件与其余模块内容分开下载。如果原始模块没有`go.mod`文件，那么代理将提供一个合成的`go.mod`文件。

### +incompatible versions

​	在主版本2或更高版本发布的模块必须在其模块路径上有一个匹配的[major version suffix（主版本后缀）](../ModulesPackagesAndVersions#major-version-suffixes)。例如，如果一个模块是以`v2.0.0`发布的，其路径必须有`/v2`的后缀。这允许 `go` 命令可以将一个项目的多个主版本视为不同的模块，即使它们是在同一个存储库中开发的。

​	主版本后缀的需求是在`go`命令添加模块支持时引入的，许多存储库在这之前已经将版本标记为主版本2或更高。为了保持与这些存储库的兼容性，`go`命令在没有`go.mod`文件的主版本2或更高的版本上添加一个`+incompatible`后缀。`+incompatible`表示某个版本与主版本号较低的版本属于同一个模块；因此，`go`命令可能会自动升级到较高的`+incompatible`版本，即使它可能会破坏构建。

​	考虑下面的示例需求：

```
require example.com/m v4.1.2+incompatible
```

​	版本`v4.1.2+incompatible`指的是提供`example.com/m`模块的存储库中的[semantic version tag（语义版本标签）](../Glossary#semantic-version-tag)`v4.1.2`。该模块必须在[repository root path（存储库根目录）](../Glossary#module-path)中（也就是说，存储库根路径也必须是`example.com/m`），并且不能有`go.mod`文件存在。该模块可能有主版本号较低的版本，如`v1.5.2`，`go`命令可能会自动升级这些版本到`v4.1.2+incompatible`（有关升级如何工作的信息，请参见[最小版本选择（MVS）](../MVS)）。

​	在版本`v2.0.0`被标记后迁移到模块的存储库通常应该发布一个新的主版本。在上面的例子中，作者应该创建一个路径为 `example.com/m/v5` 的模块，并发布 `v5.0.0` 版本。作者还应该更新模块中包的导入，使用前缀 `example.com/m/v5` 而不是 `example.com/m`。更详细的例子请参见[Go Modules: v2 and Beyond（Go模块：v2及以后）]({{< ref "/goBlog/2019/GoModulesV2AndBeyond" >}})。

​	注意 `+incompatible` 后缀不应该出现在存储库的标签上；像 `v4.1.2+incompatible` 这样的标签会被忽略。这个后缀只出现在 `go` 命令所使用的版本中。关于版本和标签之间的区别，请参见[Mapping versions to commits（将版本映射到提交）](../VersionControlSystems#mapping-versions-to-commits)。

​	还要注意的是，`+incompatible` 后缀可能会出现在[pseudo-versions（伪版本）](../Glossary#pseudo-version)中。例如，`v2.0.1-20200722182040-012345abcdef+incompatible`可能是一个有效的伪版本。

### Minimal module compatibility 最小的模块兼容性

​	以主版本2或更高版本发布的模块需要在其[module path（模块路径）](../Glossary#module-path)上有一个[major version suffix（主版本后缀）](../Glossary#major-version-suffix)。该模块可以在其存储库中的[major version subdirectory（主版本子目录）](../Glossary#major-version-subdirectory)下开发，也可以不在其中开发。这对于在构建 `GOPATH` 模式时在模块中导入包的包有一定的影响。

​	通常在 `GOPATH` 模式下，包被存储在与其[repository’s root path（存储库的根路径）](../Glossary#repository-root-path)相匹配的目录中，该根路径与其在存储库中的目录相关联。例如，存储库中根路径 `example.com/repo`在子目录 `sub` 的包将被存储在 `$GOPATH/src/example.com/repo/sub`，并将被导入为 `example.com/repo/sub`。

​	对于带有主版本后缀的模块，人们可能希望在`$GOPATH/src/example.com/repo/v2/sub`目录下找到包`example.com/repo/v2/sub`。这就要求在其存储库的`v2`子目录中开发该模块。`go`命令支持这一点，但不需要它（请参见[Mapping versions to commits（将版本映射到提交）](../VersionControlSystems#mapping-versions-to-commits)）。

​	如果模块不是在主版本的子目录中开发的，那么它在`GOPATH`中的目录将不包含主版本后缀，并且它的包可以在没有主版本后缀的情况下被导入。在上面的示例中，可以在目录 `$GOPATH/src/example.com/repo/sub` 中找到该包，并被导入为 `example.com/repo/sub`。

​	这给打算在模块模式和 `GOPATH` 模式下构建的包带来问题：模块模式需要后缀，而 `GOPATH` 模式不需要。

​	为了解决这个问题，在Go 1.11中加入了最小模块兼容性，并向后移植到Go 1.9.7和1.10.3。当导入路径被解析为 `GOPATH` 模式下的目录时：

（a）当解析形式为`$modpath/$vn/$dir`的导入时，其中：

- `$modpath` 是有效的模块路径。
- `$vn`是主版本后缀。
- `$dir`可能是一个是空的子目录。

（b）如果以下所有情况均为 true：

- `$modpath/$vn/$dir`包不存在于任何相关的`vendor`目录中。

- `go.mod`文件与导入文件位于同一目录中，或者位于`$GOPATH/src`根目录以下的任何父目录中。

- 不存在`$GOPATH[i]/src/$modpath/$vn/$suffix`目录（对于任何根`$GOPATH[i]`）。

- 文件`$GOPATH[d]/src/$modpath/go.mod`存在(对于某个根`$GOPATH[d]`)，并且声明模块路径为`$modpath/$vn`。


（c）那么将`$modpath/$vn/$dir`的导入解析到目录`$GOPATH[d]/src/$modpath/$dir`。

​	这个规则允许已经迁移到模块的包在`GOPATH`模式下构建时导入其他已经迁移到模块的包，即使没有使用主版本子目录。