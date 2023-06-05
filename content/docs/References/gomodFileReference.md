+++
title = "go.mod 文件参考"
date = 2023-05-17T09:59:21+08:00
weight = 2
description = ""
isCJKLanguage = true
draft = false

+++
# go.mod file reference - go.mod文件参考

> 原文：[https://go.dev/doc/modules/gomod-ref](https://go.dev/doc/modules/gomod-ref)

​	每个 Go 模块都由一个 `go.mod` 文件定义，该文件描述了该模块的属性，包括它对其他模块和 Go 版本的依赖项。

​	这些属性包括：

- 当前模块的`模块路径`。这应该是Go工具可以下载该模块的位置，例如模块代码的存储库位置。当与模块的版本号结合时，它可以作为唯一的标识符。它也是该模块中所有包的包路径的前缀。关于Go如何定位模块的更多信息，请参阅[Go模块参考](../GoModulesReference/Introduction)。
- 当前模块所需的`Go的最小版本`。
- 当前模块所需的其他模块的最低版本列表。
- 指令，可选择用其他模块版本或本地目录`替换 (replace)`所需模块，或`排除 （exclude）`所需模块的特定版本。

​	当您运行`go mod init`命令时，Go会生成一个`go.mod`文件。下面的例子创建了一个`go.mod`文件，将模块的模块路径设置为`example/mymodule`。

```
$ go mod init example/mymodule
```

​	使用`go`命令来管理依赖项。这些命令确保您的`go.mod`文件中描述的需求保持一致，并且`go.mod`文件的内容是有效的。这些命令包括`go get`和`go mod tidy`以及`go mod edit`命令。

​	关于`go`命令的参考，请看[go 命令]({{< ref "/cmd/go">}})。您可以通过输入`go help command-name`从命令行获得帮助，如`go help mod tidy`。

**参见**：

- Go 工具在您使用它们管理依赖项时对 `go.mod` 文件进行修改。更多信息请参见[管理依赖项](../../UsingAndUnderstandingGo/ManagingDependencies)。
- 有关`go.mod`文件的更多细节和限制，请参见[Go模块参考](../GoModulesReference/Introduction)。

## Example 示例

一个`go.mod`文件包括以下例子中的指令。这些在本主题的其他部分有描述。

```
module example.com/mymodule

go 1.14

require (
    example.com/othermodule v1.2.3
    example.com/thismodule v1.2.3
    example.com/thatmodule v1.2.3
)

replace example.com/thatmodule => ../thatmodule
exclude example.com/thismodule v1.3.0
```

## module 模块

​	声明模块的模块路径，它是模块的唯一标识符（与模块的版本号结合）。模块路径成为该模块包含的所有包的导入前缀。

​	更多信息请参见 `Go Modules Reference` 中的 [module 指令](../GoModulesReference/gomodFiles#module-directive)。

### Syntax 语法

```
module module-path
```

- module-path

  模块的模块路径，通常是 Go 工具可以下载该模块的存储库位置。对于版本`v2`及以后的模块，该值必须以主版本号结尾，如`/v2`。

### Examples 示例

下面的例子用`example.com`代替可以下载该模块的存储库域。

- `v0`或`v1`模块的模块声明：

  ```
  module example.com/mymodule
  ```

- `v2`模块的模块路径：

  ```
  module example.com/mymodule/v2
  ```


### Notes 注意事项

​	模块路径必须唯一标识您的模块。对于大多数模块，路径是一个URL，`go`命令可以在其中找到代码（或重定向到代码）。对于那些不会被直接下载的模块，模块路径可以是一些您能控制的名字，以确保唯一性。前缀`example/`也被保留下来，用于像这样的例子中。

​	更多细节，请参见[管理依赖项](../../UsingAndUnderstandingGo/ManagingDependencies)。

​	在实践中，模块路径通常是模块源的版本库域和版本库中的模块代码的路径。`go`命令在下载模块版本时依赖这种形式，来代表模块用户解决依赖关系。

​	即使您一开始不打算让您的模块供其他代码使用，使用它的存储库路径也是一种最佳做法，可以帮助您避免在以后发布模块时不得不重命名它。

​	如果一开始您不知道模块的最终存储库位置，可以考虑暂时使用一个安全的替代品，比如您拥有的域名或您控制的名称（如您的公司名称），以及模块名称或源目录的后续路径。更多信息，请参见[管理依赖项](../../UsingAndUnderstandingGo/ManagingDependencies)。

​	例如，如果您在`stringtools`目录下开发，您的临时模块路径可能是`<company-name>/stringtools`，如下面的例子，其中`company-name`是您公司的名字：

```
go mod init <company-name>/stringtools
```

## go

​	表示该模块是按照指令指定的Go版本的语义来编写的。

​	更多信息请参见 Go 模块参考中的 [go 指令](../GoModulesReference/gomodFiles#go-directive)。

### Syntax 语法

```
go minimum-go-version
```

- minimum-go-version

  编译本模块中的包所需的最小Go版本。

### Examples 示例

- 模块必须在1.14或更高版本的Go上运行：

  ```
  go 1.14
  ```

### Notes 注意事项

​	`go`指令最初是为了支持Go语言的向后不兼容的变化（见[Go 2 过渡](https://go.dev/design/28221-go2-transitions)）。自从引入模块以来，没有任何不兼容的语言变化，但`go`指令仍会影响新的语言特性的使用：

- 对于模块内的包，编译器会拒绝使用`go`指令指定的版本之后引入的语言特性。例如，如果一个模块的指令是`go 1.12`，它的包就不能使用`1_000_000`这样的数字字面量，这是在`Go 1.13（版本）`引入的。
- 如果一个较早的Go版本构建了该模块的一个包并遇到了编译错误，那么错误就会指出该模块是为一个较新的Go版本编写的。例如，假设一个模块是`go 1.13`，其中一个包使用数字字面量`1_000_000`。如果该包是用`Go 1.12 （版本）`构建的，编译器就会注意到该代码是为`Go 1.13（版本）`编写的。

此外，`go`命令会根据`go`指令所指定的版本改变其行为。这有以下影响：

（1）在`go 1.14`或更高版本中，可以启用自动[vendoring](../GoModulesReference/Module-awareCommands#vendoring)  。如果文件`vendor/modules.txt`存在并且与`go.mod`一致，就不需要显式使用`-mod=vendor`标志。

（2）在 `go 1.16` 或更高版本中，`all`包模式只匹配由[主模块（main module）](../GoModulesReference/Glossary)中的包和测试导入的包。这也是[go mod vendor](../GoModulesReference/Module-awareCommands#go-mod-vendor)自引入模块以来所保留的包的集合。在较低的版本中，`all`也包括由主模块中的包导入的包的测试、对这些包的测试等等。

（3）在go 1.17或更高版本

- `go.mod`文件包括一个明确的[require指令](../GoModulesReference/gomodFiles#require-directive)，为每个模块提供由主模块中的包或测试过渡导入的任何包。(在 go 1.16 或更低版本中，只有在[最小版本选择](../GoModulesReference/MVS)会选择不同版本的情况下，才会包括间接依赖。） 这个额外的信息使得[模块图的修剪](../GoModulesReference/ModuleGraphPruning)和[延迟模块加载](../GoModulesReference/ModuleGraphPruning#lazy-module-loading)成为可能。

- 由于`// indirect`依赖关系可能比以前的`go`版本多得多，间接依赖关系被记录在`go.mod`文件中的一个独立块中。

- `go mod vendor`省略了`go.mod`和`go.sum`文件中的供应商依赖关系。(这允许在`vendor`的子目录中调用`go`命令来识别正确的主模块）。

- `go mod vendor`在`vendor/modules.txt`中记录每个依赖关系的`go.mod`文件的`go`版本。


​	一个`go.mod`文件最多可以包含一个`go`指令。如果没有`go`指令，大多数命令会添加一个当前Go版本的`go`指令。

## require 

​	将模块声明为当前模块的依赖项，并指定所需模块的最小版本。

​	更多信息请参见 Go 模块参考中的 [require 指令](../GoModulesReference/gomodFiles#require-directive)。

### Syntax 语法

```
require module-path module-version
```

- module-path

  模块的模块路径，通常是模块源的存储库域和模块名称的连接。对于版本v2及以后的模块，该值必须以主要版本号结尾，如`/v2`。

- module-version

  模块的版本。这可以是一个发布的版本号，如v1.2.3，或Go生成的伪版本号，如v0.0.0-20200921210052-fa0125251cc4。

### Examples 示例

- 要求发布版本v1.2.3：

  ```
  require example.com/othermodule v1.2.3
  ```

- 通过使用 Go 工具生成的伪版本号来要求其存储库中尚未标记的版本。

  ```
  require example.com/othermodule v0.0.0-20200921210052-fa0125251cc4
  ```


### Notes 注意事项

​	当您运行 `go` 命令时，例如 `go get`，Go 会为包含导入包的每个模块插入 `require` 指令。当一个模块还没有在其存储库中被标记时，Go 会分配一个它在运行命令时生成的伪版本号。

​	您可以通过使用[replace指令](#replace)让 Go 从其存储库以外的地方要求一个模块。

​	更多关于版本号的信息，请参见[模块版本号](../../UsingAndUnderstandingGo/DevelopingModules/ModuleVersionNumbering)。

​	关于管理依赖项的更多信息，请参见下文：

- [Adding a dependency（添加一个依赖关系）](../../UsingAndUnderstandingGo/ManagingDependencies#adding-a-dependency)

- [Getting a specific dependency version（获取一个特定的依赖关系版本）](../../UsingAndUnderstandingGo/ManagingDependencies#getting-a-specific-dependency-version)

- [Discovering available updates（发现可用的更新）](../../UsingAndUnderstandingGo/ManagingDependencies#discovering-available-updates)

- [Upgrading or downgrading a dependency（升级或降级一个依赖项）](../../UsingAndUnderstandingGo/ManagingDependencies#upgrading-or-downgrading-a-dependency)

- [Synchronizing your code’s dependencies（同步您的代码的依赖项）](../../UsingAndUnderstandingGo/ManagingDependencies#synchronizing-your-code-s-dependencies)

  


## replace 替换

​	将特定版本(或所有版本)的模块内容替换为另一个模块版本或本地目录。Go工具在解析依赖项时将使用替换路径。

​	更多信息请参见 Go 模块参考中的 [replace 指令](../GoModulesReference/gomodFiles#replace-directive)。

### Syntax 语法

```
replace module-path [module-version] => replacement-path [replacement-version]
```

- `module-path`

  要替换的模块的模块路径。

- `module-version`

  可选的。要替换的特定版本。如果这个版本号被省略，该模块的所有版本都会被替换成箭头右侧的内容。

- `replacement-path`

  Go应该寻找所需模块的路径。这可以是一个模块路径，也可以是文件系统中与替换模块相关的目录的路径。如果这是一个模块路径，则必须指定一个替换版本（`replacement-version`）的值。如果这是一个本地路径，则不能使用替换版本（`replacement-version`）的值。

- `replacement-version`

  替换模块的版本。只有当`replacement-path`是一个模块路径（不是本地目录）时，才可以指定替换版本。

### Examples 示例

- 用模块存储库的一个分叉来替换

  在下面的例子中，example.com/othermodule的任何版本都被替换成指定的分叉代码

  ```
  require example.com/othermodule v1.2.3
  
  replace example.com/othermodule => example.com/myfork/othermodule v1.2.3-fixed
  ```

  ​	当另一个模块路径替换一个模块时，不用更改要替换的模块中包的导入语句。

  ​	关于使用分叉的模块代码副本的更多信息，请参阅[Requiring external module code from your own repository fork（从您自己的存储库分叉中请求外部模块代码）](../../UsingAndUnderstandingGo/ManagingDependencies#requiring-external-module-code-from-your-own-repository-fork)。

- 用不同的版本号替换

  下面的例子指定使用`v1.2.3`版本，而不是该模块的任何其他版本。

  ```
  require example.com/othermodule v1.2.2
  
  replace example.com/othermodule => example.com/othermodule v1.2.3
  ```

  下面的例子用同一模块的`v1.2.3`版本替换了模块`v1.2.5`版本。

  ```
  replace example.com/othermodule v1.2.5 => example.com/othermodule v1.2.3
  ```

- 用本地代码替换

  下面的示例指定应使用本地目录替换模块的所有版本。

  ```
  require example.com/othermodule v1.2.3
  
  replace example.com/othermodule => ../othermodule
  ```
  
  下面的示例指定只能使用本地目录替换 `v1.2.5`。

  ```
  require example.com/othermodule v1.2.5
  
  replace example.com/othermodule v1.2.5 => ../othermodule
  ```
  
  ​	关于使用模块代码的本地副本的更多信息，请参见[Requiring module code in a local directory](../../UsingAndUnderstandingGo/ManagingDependencies#requiring-module-code-in-a-local-directory)。

### Notes 注意事项

​	当您想让Go使用另一个路径来查找模块的源代码时，可以使用`replace`指令来暂时用另一个值来替换模块的路径值。这样做的效果是将 Go 的模块搜索重定向到替换的位置。您不必更改包的导入路径以使用替换路径。

​	在构建当前模块时，使用`exclude`和`replace`指令来控制构建时的依赖项解析。**这些指令在依赖于当前模块的模块中会被忽略。**

​	`replace`指令在以下情况下很有用：

- 您正在开发一个新的模块，其代码还没有进入存储库。您想用一个本地版本的客户端进行测试。
- 您发现了一个依赖项的问题，克隆了这个依赖项的存储库，并且正在使用本地存储库测试一个修复程序。

> ​	请注意，单独的`replace`指令并不能将一个模块添加到[module graph（模块图）](../GoModulesReference/Glossary#module-graph)中。在主模块的 `go.mod` 文件或依赖模块的 `go.mod` 文件中，还需要一个指向被替换模块版本的 [require 指令](#require)。如果您没有一个特定的版本要替换，您可以使用一个假版本，就像下面的示例。注意，这将破坏依赖于您的模块的模块，因为`replace`指令仅应用于主模块。

```
require example.com/mod v0.0.0-replace

replace example.com/mod v0.0.0-replace => ./mod
```

​	关于替换所需模块的更多信息，包括使用Go工具进行更改，请参见：

- [Requiring external module code from your own repository fork（从您自己的版本库分叉中请求外部模块代码）](../../UsingAndUnderstandingGo/ManagingDependencies#requiring-external-module-code-from-your-own-repository-fork)
- [Requiring module code in a local directory（要求本地目录中的模块代码）](../../UsingAndUnderstandingGo/ManagingDependencies#requiring-module-code-in-a-local-directory)

更多关于版本号的信息，参见[模块版本号](../../UsingAndUnderstandingGo/DevelopingModules/ModuleVersionNumbering)。

## exclude 排除

​	指定要从当前模块的依赖关系图中排除的模块或模块版本。

​	更多信息请参见 Go Modules Reference 中的 [exclude 指令](../GoModulesReference/gomodFiles#exclude-directive)。

### Syntax 语法

```
exclude module-path module-version
```

- module-path

  要排除的模块的模块路径。

- module-version

  要排除的特定版本。

### Example 示例

- 排除 example.com/theirmodule 版本 `v1.3.0`

  ```
  exclude example.com/theirmodule v1.3.0
  ```


### Notes 注意事项

​	使用`exclude`指令可以排除一个间接需要但由于某种原因无法加载的模块的特定版本。例如，您可以用它来排除一个有无效校验和的模块的版本。

​	在构建当前模块（您正在构建的主模块）时，使用`exclude`和`replace`指令来控制构建时的依赖项解析。**这些指令在依赖于当前模块的模块中会被忽略**。

​	您可以使用[go mod edit](../GoModulesReference/Module-awareCommands#go-mod-edit)命令来排除一个模块，如下面的示例。

```
go mod edit -exclude=example.com/theirmodule@v1.3.0
```

​	更多关于版本号的信息，请参见[模块版本号](../../UsingAndUnderstandingGo/DevelopingModules/ModuleVersionNumbering)。

## retract 撤回

​	指示不应依赖`go.mod`定义的模块的版本或版本范围。当版本提前发布或在发布后发现严重问题时，`retract`指令很有用。

​	更多信息请参见 Go Modules Reference 中的 [retract 指令](../GoModulesReference/gomodFiles#retract-directive)。

### Syntax 语法

```
retract version // rationale
retract [version-low,version-high] // rationale
```

- version

  要撤回的单个版本。

- version-low

  要撤回的版本范围的下限。

- version-high

  要撤回的版本范围的上限。版本-低和版本-高都包括在这个范围内。

- rationale

  解释撤回的可选注释。可以在发给用户的消息中显示。

### Example 示例

- Retracting a single version 撤回单个版本

  ```
  retract v1.1.0 // Published accidentally.
  ```

- Retracting a range of versions 撤回一系列的版本

  ```
  retract [v1.0.0,v1.0.5] // Build broken on some platforms.
  ```

### Notes 注意事项

​	使用`retract`指令来指示您的模块的前一个版本不应该被使用。用户不会通过`go get`、`go mod tidy`或其他命令自动升级到撤回的版本。用户不会在`go list -m -u`中看到撤回的版本是可用的更新。

​	撤回的版本应该保持可用，以便已经依赖它们的用户能够构建他们的软件包。即使撤回的版本已经从源码库中删除，它仍然可以在[proxy.golang.org](https://proxy.golang.org/)这样的镜像中使用。当用户在相关模块上运行 `go get` 或 `go list -m -u` 时，他们可能会被通知依赖被撤回的版本。

​	`go`命令通过读取模块最新版本的`go.mod`文件中的`retract`指令来发现撤回的版本。最新的版本是，按优先顺序排列：

1. 它的最高发布版本，如果有的话；
2. 它的最高预发布版本，如果有的话；
3. 储存库默认分支的尖端（tip）的伪版本。

​	当您添加一个撤回内容时，**您几乎总是需要标记一个新的、更高的版本**，这样命令就会在模块的最新版本中看到它。

​	您可以发布一个版本，该版本的唯一目的是发出撤回的信号。在这种情况下，新版本也可能撤回自己。

​	例如，如果您不小心标记了`v1.0.0`，您可以用以下指令标记`v1.0.1`：

```
retract v1.0.0 // Published accidentally. 已意外地发布。
retract v1.0.1 // Contains retraction only. 仅包含撤回。
```

​	遗憾的是，一旦版本发布，就无法更改。如果稍后在不同的提交中标记了`v1.0.0`，`go`命令可能会在`go.sum`或[校验数据库](../GoModulesReference/AuthenticatingModules#checksum-database)中检测到一个不匹配的和。

​	模块的撤回版本通常不会出现在 `go list -m -versions` 的输出中，但您可以使用 `-retracted` 来显示它们。更多信息请参见 Go Modules Reference 中的 [go list -m](../GoModulesReference/Module-awareCommands#go-list-m)。