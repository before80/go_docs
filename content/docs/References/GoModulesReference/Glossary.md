+++
title = "词汇表"
date = 2023-05-17T09:59:21+08:00
weight = 16
description = ""
isCJKLanguage = true
draft = false
+++
## Glossary词汇表

> 原文：[https://go.dev/ref/mod#glossary](https://go.dev/ref/mod#glossary)

### build constraint

A condition that determines whether a Go source file is used when compiling a package. Build constraints may be expressed with file name suffixes (for example, `foo_linux_amd64.go`) or with build constraint comments (for example, `// +build linux,amd64`). See [Build Constraints](https://go.dev/pkg/go/build/#hdr-Build_Constraints).

**build constraint** （构建约束）：决定在编译包时是否使用Go源文件的条件。构建约束可以用文件名后缀来表示（例如，`foo_linux_amd64.go`），也可以用构建约束的注释来表达（例如，`// +build linux,amd64`）。参见[构建约束](https://pkg.go.dev/go/build#hdr-Build_Constraints)。

### build list

The list of module versions that will be used for a build command such as `go build`, `go list`, or `go test`. The build list is determined from the [main module’s](https://go.dev/ref/mod#glos-main-module) [`go.mod` file](https://go.dev/ref/mod#glos-go-mod-file) and `go.mod` files in transitively required modules using [minimal version selection](https://go.dev/ref/mod#glos-minimal-version-selection). The build list contains versions for all modules in the [module graph](https://go.dev/ref/mod#glos-module-graph), not just those relevant to a specific command.

build list: 将用于构建命令的模块版本列表，如go build、go list或go test。构建列表是由主模块的go.mod文件和中转所需模块的go.mod文件决定的，使用最小版本选择。构建列表包含了模块图中所有模块的版本，而不仅仅是与特定命令相关的模块。

### canonical version

 A correctly formatted [version](https://go.dev/ref/mod#glos-version) without a build metadata suffix other than `+incompatible`. For example, `v1.2.3` is a canonical version, but `v1.2.3+meta` is not.

典型版本。一个格式正确的版本，除了+incompatible之外，没有构建元数据的后缀。例如，v1.2.3是一个典型的版本，但v1.2.3+meta不是。

### current module

 Synonym for [main module](https://go.dev/ref/mod#glos-main-module).

当前模块。主模块的同义词。

### deprecated module

A module that is no longer supported by its authors (though major versions are considered distinct modules for this purpose). A deprecated module is marked with a [deprecation comment](https://go.dev/ref/mod#go-mod-file-module-deprecation) in the latest version of its [`go.mod` file](https://go.dev/ref/mod#glos-go-mod-file).

deprecated module（废弃的模块）。一个不再被其作者支持的模块（尽管主要版本在此被视为独立的模块）。被废弃的模块在其go.mod文件的最新版本中被标记为废弃注释。

### direct dependency 

A package whose path appears in an [`import` declaration](https://go.dev/ref/spec#import_declarations) in a `.go` source file for a package or test in the [main module](https://go.dev/ref/mod#glos-main-module), or the module containing such a package. (Compare [indirect dependency](https://go.dev/ref/mod#glos-indirect-dependency).)

直接依赖关系。一个软件包，其路径出现在主模块中的包或测试的.go源文件的导入声明中，或者包含这样一个包的模块。(比较间接依赖)。

### direct mode

 A setting of [environment variables](https://go.dev/ref/mod#environment-variables) that causes the `go` command to download a module directly from a [version control system](https://go.dev/ref/mod#vcs), as opposed to a [module proxy](https://go.dev/ref/mod#glos-module-proxy). `GOPROXY=direct` does this for all modules. `GOPRIVATE` and `GONOPROXY` do this for modules matching a list of patterns.

直接模式。一种环境变量的设置，使go命令直接从版本控制系统中下载模块，而不是通过模块代理。GOPROXY=direct 对所有模块都是如此。GOPRIVATE和GONOPROXY对符合模式列表的模块做此处理。

### go.mod file 

The file that defines a module’s path, requirements, and other metadata. Appears in the [module’s root directory](https://go.dev/ref/mod#glos-module-root-directory). See the section on [`go.mod` files](https://go.dev/ref/mod#go-mod-file).

go.mod 文件。定义一个模块的路径、要求和其他元数据的文件。出现在模块的根目录下。参见 go.mod 文件一节。

### go.work file 

The file that defines the set of modules to be used in a [workspace](https://go.dev/ref/mod#workspaces). See the section on [`go.work` files](https://go.dev/ref/mod#go-work-file)

go.work文件：定义工作区中使用的模块集的文件。见go.work文件一节

### import path

 A string used to import a package in a Go source file. Synonymous with [package path](https://go.dev/ref/mod#glos-package-path).

import路径。用来在 Go 源文件中导入软件包的字符串。与软件包路径同义。

### indirect dependency

 A package transitively imported by a package or test in the [main module](https://go.dev/ref/mod#glos-main-module), but whose path does not appear in any [`import` declaration](https://go.dev/ref/spec#import_declarations) in the main module; or a module that appears in the [module graph](https://go.dev/ref/mod#glos-module-graph) but does not provide any package directly imported by the main module. (Compare [direct dependency](https://go.dev/ref/mod#glos-direct-dependency).)

间接依赖。一个被主模块中的包或测试间接导入的包，但其路径没有出现在主模块的任何导入声明中；或者一个出现在模块图中的模块，但没有提供任何被主模块直接导入的包。(比较直接依赖）。

### lazy module loading

 A change in Go 1.17 that avoids loading the [module graph](https://go.dev/ref/mod#glos-module-graph) for commands that do not need it in modules that specify `go 1.17` or higher. See [Lazy module loading](https://go.dev/ref/mod#lazy-loading).

懒惰的模块加载。Go 1.17中的一个变化，在指定go 1.17或更高版本的模块中避免为不需要的命令加载模块图。参见 "懒惰模块加载"。

### main module 

The module in which the `go` command is invoked. The main module is defined by a [`go.mod` file](https://go.dev/ref/mod#glos-go-mod-file) in the current directory or a parent directory. See [Modules, packages, and versions](https://go.dev/ref/mod#modules-overview).

主模块。调用go命令的模块。主模块由当前目录或父目录中的go.mod文件定义。参见模块、包和版本。

### major version 

The first number in a semantic version (`1` in `v1.2.3`). In a release with incompatible changes, the major version must be incremented, and the minor and patch versions must be set to 0. Semantic versions with major version 0 are considered unstable.

主版本。语义版本中的第一个数字（1在v1.2.3中）。在有不兼容变化的版本中，主版本必须被递增，而次版本和补丁版本必须被设置为0。 主版本为0的语义版本被认为是不稳定的。

### major version subdirectory 

A subdirectory within a version control repository matching a module’s [major version suffix](https://go.dev/ref/mod#glos-major-version-suffix) where a module may be defined. For example, the module `example.com/mod/v2` in the repository with [root path](https://go.dev/ref/mod#glos-repository-root-path) `example.com/mod` may be defined in the repository root directory or the major version subdirectory `v2`. See [Module directories within a repository](https://go.dev/ref/mod#vcs-dir).

major version subdirectory（主要版本子目录）。版本控制库中的一个子目录，与模块的主要版本后缀相匹配，在这里可以定义一个模块。例如，根路径为 example.com/mod 的版本库中的模块 example.com/mod/v2 可以定义在版本库根目录或主要版本子目录 v2 中。参见版本库内的模块目录。

### major version suffix

 A module path suffix that matches the major version number. For example, `/v2` in `example.com/mod/v2`. Major version suffixes are required at `v2.0.0` and later and are not allowed at earlier versions. See the section on [Major version suffixes](https://go.dev/ref/mod#major-version-suffixes).

主要版本的后缀。一个与主要版本号相匹配的模块路径后缀。例如，/v2 在 example.com/mod/v2。主要版本后缀在v2.0.0及以后的版本中是必须的，在早期版本中不允许使用。参见主要版本后缀一节。

### minimal version selection (MVS)

 The algorithm used to determine the versions of all modules that will be used in a build. See the section on [Minimal version selection](https://go.dev/ref/mod#minimal-version-selection) for details.

最小版本选择（MVS）。用来确定构建中使用的所有模块的版本的算法。详见最小版本选择一节。

### minor version 

The second number in a semantic version (`2` in `v1.2.3`). In a release with new, backwards compatible functionality, the minor version must be incremented, and the patch version must be set to 0.

次要版本。语义版本中的第二个数字（V1.2.3中为2）。在具有新的、向后兼容的功能的版本中，次要版本必须被递增，而补丁版本必须被设置为0。

### module 

A collection of packages that are released, versioned, and distributed together.

模块。一组包的集合，它们被一起发布、版本化和分发。

**module cache:** A local directory storing downloaded modules, located in `GOPATH/pkg/mod`. See [Module cache](https://go.dev/ref/mod#module-cache).

模块缓存。一个存储下载模块的本地目录，位于GOPATH/pkg/mod。参见模块缓存。

### module graph 

The directed graph of module requirements, rooted at the [main module](https://go.dev/ref/mod#glos-main-module). Each vertex in the graph is a module; each edge is a version from a `require` statement in a `go.mod` file (subject to `replace` and `exclude` statements in the main module’s `go.mod` file).

模块图。模块需求的有向图，以主模块为根。图中的每个顶点都是一个模块；每条边都是来自 go.mod 文件中 require 语句的版本（受主模块的 go.mod 文件中的替换和排除语句制约）。

### module graph pruning 

A change in Go 1.17 that reduces the size of the module graph by omitting transitive dependencies of modules that specify `go 1.17` or higher. See [Module graph pruning](https://go.dev/ref/mod#graph-pruning).

模块图的修剪。Go 1.17中的一个变化，通过省略指定go 1.17或更高版本的模块的相互依赖关系来减少模块图的大小。参见模块图修剪。

### module path 

A path that identifies a module and acts as a prefix for package import paths within the module. For example, `"golang.org/x/net"`.

模块路径。一个标识模块的路径，作为模块内包导入路径的前缀。例如，"golang.org/x/net"。

### module proxy

 A web server that implements the [`GOPROXY` protocol](https://go.dev/ref/mod#goproxy-protocol). The `go` command downloads version information, `go.mod` files, and module zip files from module proxies.

模块代理。一个实现GOPROXY协议的网络服务器。go命令从模块代理处下载版本信息、go.mod文件和模块压缩文件。

### module root directory

 The directory that contains the `go.mod` file that defines a module.

模块根目录。包含定义模块的go.mod文件的目录。

### module subdirectory

 The portion of a [module path](https://go.dev/ref/mod#glos-module-path) after the [repository root path](https://go.dev/ref/mod#glos-repository-root-path) that indicates the subdirectory where the module is defined. When non-empty, the module subdirectory is also a prefix for [semantic version tags](https://go.dev/ref/mod#glos-semantic-version-tag). The module subdirectory does not include the [major version suffix](https://go.dev/ref/mod#glos-major-version-suffix), if there is one, even if the module is in a [major version subdirectory](https://go.dev/ref/mod#glos-major-version-subdirectory). See [Module paths](https://go.dev/ref/mod#module-path).

模块子目录。模块路径在版本库根路径之后的部分，表示定义该模块的子目录。非空时，模块子目录也是语义版本标签的前缀。模块子目录不包括主要版本后缀（如果有的话），即使该模块位于主要版本子目录中。参见模块路径。

### package 

A collection of source files in the same directory that are compiled together. See the [Packages section](https://go.dev/ref/spec#Packages) in the Go Language Specification.

包。在同一目录下的源文件的集合，这些文件被编译在一起。参见 Go 语言规范中的包部分。

### package path 

The path that uniquely identifies a package. A package path is a [module path](https://go.dev/ref/mod#glos-module-path) joined with a subdirectory within the module. For example `"golang.org/x/net/html"` is the package path for the package in the module `"golang.org/x/net"` in the `"html"` subdirectory. Synonym of [import path](https://go.dev/ref/mod#glos-import-path).

包路径。唯一标识一个包的路径。包的路径是一个模块的路径与模块内的子目录相连接。例如 "golang.org/x/net/html "是模块 "golang.org/x/net "中 "html "子目录下的包的路径。进口路径的同义词。

### patch version 

The third number in a semantic version (`3` in `v1.2.3`). In a release with no changes to the module’s public interface, the patch version must be incremented.

补丁版本。语义版本中的第三个数字（v1.2.3中为3）。在模块的公共接口没有变化的版本中，补丁版本必须被递增。

### pre-release version 

A version with a dash followed by a series of dot-separated identifiers immediately following the patch version, for example, `v1.2.3-beta4`. Pre-release versions are considered unstable and are not assumed to be compatible with other versions. A pre-release version sorts before the corresponding release version: `v1.2.3-pre` comes before `v1.2.3`. See also [release version](https://go.dev/ref/mod#glos-release-version).

pre-release version（预发布版本）。在补丁版本后面紧跟着一系列点分隔的标识符，带有破折号的版本，例如，v1.2.3-beta4。预发布版本被认为是不稳定的，不被认为与其他版本兼容。预发布版本会在相应的发布版本之前排序：v1.2.3-pre会在v1.2.3之前。另见发布版本。

### pseudo-version 

A version that encodes a revision identifier (such as a Git commit hash) and a timestamp from a version control system. For example, `v0.0.0-20191109021931-daa7c04131f5`. Used for [compatibility with non-module repositories](https://go.dev/ref/mod#non-module-compat) and in other situations when a tagged version is not available.

pseudo-version（伪版本）。一个编码了修订标识符（如Git提交哈希值）和版本控制系统的时间戳的版本。例如，v0.0.0-20191109021931-daa7c04131f5。 用于与非模块仓库的兼容性，以及其他无法获得标记版本的情况。

### release version 

A version without a pre-release suffix. For example, `v1.2.3`, not `v1.2.3-pre`. See also [pre-release version](https://go.dev/ref/mod#glos-pre-release-version).

发布版本。一个没有预发布后缀的版本。例如，v1.2.3，而不是v1.2.3-pre。另见预发布版本。

### repository root path 

The portion of a [module path](https://go.dev/ref/mod#glos-module-path) that corresponds to a version control repository’s root directory. See [Module paths](https://go.dev/ref/mod#module-path).

仓库根路径。模块路径中与版本控制库的根目录相对应的部分。参见模块路径。

### retracted version

 A version that should not be depended upon, either because it was published prematurely or because a severe problem was discovered after it was published. See [`retract` directive](https://go.dev/ref/mod#go-mod-file-retract).

撤回的版本。一个不应该被依赖的版本，因为它被过早地发布，或者在发布后发现了严重的问题。参见retract指令。

### semantic version tag

 A tag in a version control repository that maps a [version](https://go.dev/ref/mod#glos-version) to a specific revision. See [Mapping versions to commits](https://go.dev/ref/mod#vcs-version).

语义版本标签。版本控制库中的一个标签，将一个版本映射到一个特定的修订版。参见版本与提交的映射。

### selected version 

The version of a given module chosen by [minimal version selection](https://go.dev/ref/mod#minimal-version-selection). The selected version is the highest version for the module’s path found in the [module graph](https://go.dev/ref/mod#glos-module-graph).

选定的版本。通过最小版本选择选择的特定模块的版本。选择的版本是在模块图中发现的该模块路径的最高版本。

### vendor directory

 A directory named `vendor` that contains packages from other modules needed to build packages in the main module. Maintained with [`go mod vendor`](https://go.dev/ref/mod#go-mod-vendor). See [Vendoring](https://go.dev/ref/mod#vendoring).

vendor目录。一个名为vendor的目录，包含了其他模块的包，需要在主模块中构建包。用go mod vendor维护。参看vendoring。

### version 

An identifier for an immutable snapshot of a module, written as the letter `v` followed by a semantic version. See the section on [Versions](https://go.dev/ref/mod#versions).

版本。一个模块的不可改变的快照的标识符，写成字母v，后面是语义版本。参见 "版本 "一节。

### workspace

 A collection of modules on disk that are used as the main modules when running [minimal version selection (MVS)](https://go.dev/ref/mod#minimal-version-selection). See the section on [Workspaces](https://go.dev/ref/mod#workspaces)

workspace（工作区）。磁盘上的模块集合，在运行最小版本选择（MVS）时被用作主要模块。见关于工作区的章节