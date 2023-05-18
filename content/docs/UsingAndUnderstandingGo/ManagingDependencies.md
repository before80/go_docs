+++
title = "管理依赖项"
linkTitle = "管理依赖项"
weight = 20
date = 2023-05-17T15:03:14+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Managing dependencies - 管理依赖项

> 原文：[https://go.dev/doc/modules/managing-dependencies](https://go.dev/doc/modules/managing-dependencies)

​	当你的代码使用外部包时，这些包（作为模块分发）就成为了依赖。随着时间的推移，你可能需要升级它们或替换它们。Go 提供了依赖项管理工具，帮助您在纳入外部依赖时保持 Go 应用程序的安全。

​	本主题介绍了如何执行任务来管理代码中的依赖项。你可以用 Go 工具来执行其中的大多数任务。本主题还介绍了如何执行一些其他与依赖项有关的任务，您可能会发现这些任务很有用。

还请参见：

- 如果你刚开始将依赖项作为模块来使用，请参阅[Getting started tutorial （入门教程）](../../GettingStarted/TutorialGetStartedWithGo)，以获得简短的介绍。
- 使用 `go` 命令来管理依赖项有助于确保你的需求保持一致，以及 go.mod 文件内容的有效性。关于命令的参考，请看[Command go](../../References/CommandDocumentation/go)。你也可以在命令行中输入`go help command-name`来获得帮助，如`go help mod tidy`。
- 编辑 `go.mod`文件时使用的用于进行依赖项更改的 Go 命令。关于该文件的内容，请参阅 [go.mod file reference（go.mod文件参考）](../../References/gomodFileReference)。
- 让你的编辑器或 IDE 了解 Go 模块可以使管理它们的工作更容易。关于支持 Go 的编辑器的更多信息，请参阅 [Editor plugins and IDEs（编辑器插件和 IDE）](../EditorPluginsAndIDEs)。
- 本主题不介绍如何开发、发布和版本模块供他人使用。有关这方面的更多信息，请参见 [Developing and publishing modules（开发和发布模块）](../DevelopingModules/DevelopingAndPublishingModules)。

## Workflow for using and managing dependencies 使用和管理依赖项的工作流程

​	你可以通过Go工具获取并使用有用的包。在 [pkg.go.dev](https://pkg.go.dev/) 上，你可以搜索你觉得有用的包，然后使用 `go` 命令将这些包导入你自己的代码中，以调用它们的功能。

​	下面列出了最常见的依赖项管理步骤。关于每个步骤的更多信息，请参见本主题中的章节。

6. 在 [pkg.go.dev](https://pkg.go.dev/) 上[找到有用的包](#locating-and-importing-useful-packages)。

7. 在你的代码中[导入你想要的包](#locating-and-importing-useful-packages)。

8. 将你的代码添加到一个模块中以进行依赖项跟踪（如果尚未在模块中）。参见[启用依赖项跟踪](#enabling-dependency-tracking-in-your-code)

9. [将外部包添加为依赖项](#adding-a-dependency)，以便你可以管理它们。

10. 随着时间的推移，根据需要[升级或降级依赖的版本](#upgrading-or-downgrading-a-dependency)。

    

## Managing dependencies as modules 将依赖项作为模块进行管理

​	在Go中，你可以将依赖项作为包含导入的包的模块来管理。这个过程得到了以下支持：

- 用于发布模块和检索其代码的分散系统。开发者从他们自己的仓库中提供他们的模块供其他开发者使用，并以版本号进行发布。
- **包搜索引擎**和文档浏览器（[pkg.go.dev](https://pkg.go.dev/)），你可以在那里找到模块。参见[找到并导入有用的软件包](#locating-and-importing-useful-packages)。
- 模块**版本号约定**，帮助你了解一个模块的稳定性和向后兼容性保证。参见[模块版本号](../DevelopingModules/ModuleVersionNnumbering)。
- **go tools**使你更容易管理依赖项，包括获取模块的源代码、升级等等。更多内容请参见本主题的各个章节。

## Locating and importing useful packages 找到并导入有用的软件包

​	你可以在[pkg.go.dev](https://pkg.go.dev/)上搜索，以查找你觉得有用的功能包。

​	当你找到一个你想在你的代码中使用的包时，在页面顶部找到包的路径，点击复制路径按钮，将路径复制到你的剪贴板上。在你自己的代码中，将该路径粘贴到导入语句中，如下面的例子：

```go 
import "rsc.io/quote"
```

​	在你的代码导入包后，启用依赖项跟踪并使用该包的代码进行编译。更多信息，请参阅在你的代码中[启用依赖项跟踪](#enabling-dependency-tracking-in-your-code)和[添加依赖项](#adding-a-dependency)。

## Enabling dependency tracking in your code 在你的代码中启用依赖项跟踪

​	为了跟踪和管理你添加的依赖项，你首先要把你的代码放在自己的模块中。这将在你的源代码树的根目录下创建一个`go.mod`文件。你添加的依赖将被列在该文件中。

​	要将你的代码添加到自己的模块中，使用[go mod init](../../References/CommandDocumentation/go#go-mod-init)命令。例如，从命令行切换到代码的根目录，然后运行下面例子中的命令：

```shell
$ go mod init example/mymodule
```

​	`go mod init`命令的参数是你的模块的模块路径。如果可能的话，该模块路径应该是你的源代码的版本库位置。

​	如果一开始你不知道模块的最终仓库位置，请使用一个安全的替代品。这可能是你拥有的域名或你控制的其他名称（如你的公司名称），以及从模块的名称或源目录后的一个路径。更多信息，请参见[模块的命名](#naming-a-module)。

​	当你使用 Go 工具来管理依赖项时，这些工具会更新 `go.mod` 文件，以便它维护你的依赖项的当前列表。

​	当你添加依赖项时，Go 工具还会创建一个 `go.sum` 文件，其中包含你所依赖模块的校验和。Go 使用它来验证下载的模块文件的完整性，特别是对于从事你项目的其他开发人员。

​	将`go.mod`和`go.sum`文件与你的代码一起包含在你的版本库中。

​	更多信息请参见 [go.mod file reference（go.mod文件参考）](../../References/gomodFileReference)。

## Naming a module 命名一个模块

​	当你运行 `go mod init` 来创建一个用于跟踪依赖项的模块时，你会指定一个模块路径作为模块的名称。该模块路径成为该模块中包的导入路径前缀。请确保指定一个不会与其他模块的路径冲突的模块路径。

​	至少，一个模块的路径只需要表明它的来源，比如公司、作者或所有者的名称。但是路径也可以更多的描述模块是什么或做什么。

​	模块路径通常采用以下形式：

```
<prefix>/<descriptive-text>
```

- 前缀（prefix）通常是部分描述模块的字符串，例如描述其来源的字符串。这可能是：

  - Go工具可以找到该模块源代码的版本库位置（如果你要发布该模块，则则需要这个位置）。

    例如，它可能是`github.com/<project-name>/`。

    如果你认为你可能会发布该模块供他人使用，请使用此最佳实践。关于发布的更多信息，请参阅 [Developing and publishing modules（开发和发布模块）](../DevelopingModules/DevelopingAndPublishingModules)。

  - 一个你控制的名称。

    如果你不使用版本库名称，请确保选择一个你确信不会被他人使用的前缀。一个好的选择是你公司的名称。避免使用常见的术语，如`widgets`、`utilities`或`app`。

- 对于描述性文本（descriptive text），一个好的选择是项目名称。记住，包名承载了描述功能的重任。模块路径为这些包名创建一个命名空间。

**Reserved module path prefixes** 保留的模块路径前缀

​	Go 保证以下字符串不会在包名中使用。

- `test` - 你可以用`test`作为模块路径前缀，该模块的代码被设计用于本地测试另一个模块中的。

  对于作为测试的一部分而创建的模块，请使用`test`路径前缀。例如，你的测试本身可能会运行`go mod init test`，然后以某种特殊方式设置该模块，以便用Go源代码分析工具进行测试。

- `example` - 在一些Go文档中用作模块路径前缀，例如在教程中，你创建一个模块只是为了跟踪依赖项。

  请注意，Go 文档也使用 `example.com` 来说明例子可能是一个已发布的模块。

## Adding a dependency 添加一个依赖项

​	一旦你导入一个已发布的模块，你就可以使用 [go get](../../References/CommandDocumentation/go#add-dependencies-to-current-module-and-install-them) 命令将该模块作为一个依赖项来管理。

该命令做了以下工作：

- 如果需要，它会在`go.mod`文件中添加`require`指令，用于构建在命令行上命名的包所需的模块。一个`require`指令可以追踪模块所依赖的模块的最小版本。更多信息请参见 [go.mod file reference（go.mod文件参考）](../../References/gomodFileReference)。

- 如果需要，它会下载模块的源代码，这样你就可以编译依赖于它们的包。它可以从模块代理（如`proxy.golang.org`）或直接从版本控制库下载模块。源代码被缓存在本地。

  你可以设置Go工具下载模块的位置。更多信息请参见[指定模块代理服务器](#specifying-a-module-proxy-server)。

下面介绍几个例子。

- 要在你的模块中添加一个包的所有依赖项，运行类似下面的命令（"`.` "指的是当前目录中的包）：

  ```shell
  $ go get .
  ```

- 要添加一个特定的依赖项，请将其模块路径指定为命令的参数。

  ```shell
  $ go get example.com/theirmodule
  ```

​	该命令还对其下载的每个模块进行认证。这可以确保它与模块发布时没有变化。如果模块在发布后发生了变化——例如，开发者改变了提交的内容——Go工具会显示一个安全错误。这种认证检查可以保护你免受可能被篡改的模块的影响。

## Getting a specific dependency version 获取特定的依赖版本

​	你可以通过在 `go get` 命令中指定一个依赖模块的版本来获取它的特定版本。该命令会更新你`go.mod`文件中的`require`指令（当然你也可以手动更新）。

你可能想这样做，如果：

- 你想获得一个特定的预发布版本的模块以进行试用。
- 你发现你目前需要的版本不适合你，因此希望获得一个你知道可以依赖的版本。
- 你想升级或降级一个你已经需要的模块。

下面是使用[go get](../../References/CommandDocumentation/go#add-dependencies-to-current-module-and-install-them)命令的例子：

- 要获得一个特定的版本，请在模块路径后面加上`@`符号和你想要的版本：

  ```shell
  $ go get example.com/theirmodule@v1.3.4
  ```

- 要获得最新的版本，请在模块路径后加上`@latest`：

  ```shell
  $ go get example.com/theirmodule@latest
  ```

​	下面的`go.mod`文件`require`指令例子（详见[go.mod file reference（go.mod文件参考）](../../References/gomodFileReference)）说明了如何要求一个特定的版本号：

```
require example.com/theirmodule v1.3.4
```

## Discovering available updates 发现可用的更新

​	你可以检查你当前模块中已经使用的依赖项是否有更新的版本。使用`go list`命令来显示你的模块的依赖项列表，以及该模块的最新版本。一旦你发现了可用的升级版本，你就可以在你的代码中试用它们，以决定是否升级到新版本。

​	关于`go list`命令的更多信息，见[go list -m](../../References/GoModulesReference/Module-awareCommands#go-list-m)。

​	这里有几个例子。

- 列出当前模块的所有依赖模块，以及每个模块的最新版本：

  ```shell
  $ go list -m -u all
  ```

- 显示某个特定模块的最新版本：

  ```shell
  $ go list -m -u example.com/theirmodule
  ```

## Upgrading or downgrading a dependency 升级或降级一个依赖项

​	你可以使用 Go 工具来发现可用的版本，然后添加不同的版本作为依赖项，从而升级或降级依赖项模块。

1. 要发现新的版本，请使用 `go list` 命令，如 [发现可用的更新](#discovering-available-updates) 中所述。
2. 要添加一个特定的版本作为依赖项，使用 `go get` 命令，如 [获取特定的依赖项版本](#getting-a-specific-dependency-version) 中所述。

## Synchronizing your code’s dependencies 同步你代码的依赖项

​	你可以确保正在管理的代码中所有导入的包的依赖项，同时还可以删除你不再导入的包的依赖项。

​	当你对你的代码和依赖项进行更改后时，这可能很有用，可能创建了一个管理的依赖项和下载的模块的集合，这些模块不再与你代码中导入的包所特别需要的集合相匹配。

​	为了保持你的管理依赖项集的整洁，使用 `go mod tidy` 命令。该命令使用你的代码中导入的包集合，编辑你的`go.mod`文件，以添加必要但缺少的模块。它还会删除那些不提供任何相关包的未使用的模块。

​	该命令没有参数，只有一个标志，即`-v`，可以打印出被删除模块的信息。

```shell
$ go mod tidy
```

## Developing and testing against unpublished module code 针对未发布的模块代码进行开发和测试

​	你可以指定你的代码使用可能尚未发布的依赖模块。这些模块的代码可能在它们各自的版本库中，或者在这些版本库的fork中，或者在与当前模块一起使用的驱动器中。

​	你可能想在以下情况下这样做：

- 你想对一个外部模块的代码进行自己的修改，比如在 fork和/或clone之后。例如，你可能想对该模块进行修复，然后将其作为拉动请求发送给该模块的开发者。
- 你正在构建一个新的模块，但还没有发布，因此它在存储库中是不可用的，而 `go get` 命令可以访问它。

### Requiring module code in a local directory 要求本地目录中的模块代码

​	你可以指定所需模块的代码和需要它的代码在同一个本地驱动器上。你可能会发现这在以下情况下很有用：

- 开发你自己的独立模块，并希望从当前模块中进行测试。
- 修复外部模块的问题或添加功能，并希望从当前模块进行测试。(注意，你也可以从你自己的fork库中要求外部模块。更多信息，请看 [从你自己的版本库分叉中获取外部模块代码](#requiring-external-module-code-from-your-own-repository-fork-fork)）。

​	要告诉 Go 命令使用模块代码的本地副本，请使用 `go.mod` 文件中的 `replace` 指令来替换 `require` 指令中给出的模块路径。关于指令的更多信息，请参见 [go.mod file reference（go.mod文件参考）](../../References/gomodFileReference)。

​	在下面的`go.mod`文件例子中，当前模块需要外部模块`example.com/theirmodule`，使用了一个不存在的版本号（`v0.0.0-unpublished`），以确保替换正常工作。然后`replace`指令用`./theirmodule`替换原来的模块路径，这个目录与当前模块的目录处于同一级别。

```
module example.com/mymodule

go 1.16

require example.com/theirmodule v0.0.0-unpublished

replace example.com/theirmodule v0.0.0-unpublished => ../theirmodule
```

​	当设置`require`/`replace`对时，使用[go mod edit](../../References/CommandDocumentation/go#go-mod-edit)和[go get](../../References/CommandDocumentation/go#add-dependencies-to-current-module-and-install-them)命令来确保文件描述的需求保持一致。

```shell
$ go mod edit -replace=example.com/theirmodule@v0.0.0-unpublished=../theirmodule
$ go get example.com/theirmodule@v0.0.0-unpublished
```

!!! warning "注意"

	注意：当你使用`replace` 指令时，Go工具不会像[添加依赖项](#adding-a-dependency)中描述的那样对外部模块进行认证。

​	更多关于版本号的信息，请参见[模块版本号](../DevelopingModules/ModuleVersionNnumbering)。

### Requiring external module code from your own repository fork 要求从你自己的仓库fork的外部模块代码

​	当你 fork了一个外部模块的仓库时（例如修复模块代码中的问题或增加一个功能），你可以让 Go 工具使用你的fork来获取模块的源代码。这对于测试来自你自己代码的更改非常有用。(注意，你还可以要求在本地驱动器上的目录中的模块代码与需要它的模块一起使用。更多信息，请参见 [要求本地目录下的模块代码](#requiring-module-code-in-a-local-directory)）。

​	为此，你可以在`go.mod`文件中使用`replace`指令，将外部模块的原始路径替换为仓库中 fork 的路径。这样，Go工具在编译时就会使用替换路径（fork 的位置），例如，同时允许你不改变原始模块路径中的`import`语句。

​	关于`replace`指令的更多信息，请参见 [go.mod file reference（go.mod文件参考）](../../References/gomodFileReference)。

​	在下面的`go.mod`文件例子中，当前模块需要外部模块`example.com/theirmodule`。`replace`指令用`example.com/myfork/theirmodule`替换了原来的模块路径，这是一个模块自己的仓库的fork。

```
module example.com/mymodule

go 1.16

require example.com/theirmodule v1.2.3

replace example.com/theirmodule v1.2.3 => example.com/myfork/theirmodule v1.2.3-fixed
```

​	在设置`require`/`replace`对时，使用Go工具命令来确保文件描述的需求保持一致。使用 [go list](../../References/GoModulesReference/Module-awareCommands#go-list-m) 命令来获取当前模块使用的版本。然后使用 [go mod edit]([go get](../../References/CommandDocumentation/go#go-mod-edit)) 命令将所需模块替换为分叉模块：

```shell
$ go list -m example.com/theirmodule
example.com/theirmodule v1.2.3
$ go mod edit -replace=example.com/theirmodule@v1.2.3=example.com/myfork/theirmodule@v1.2.3-fixed
```

注意：当你使用`replace`指令时，Go工具不会像[添加依赖项](#adding-a-dependency)中描述的那样对外部模块进行认证。

​	更多关于版本号的信息，请参见 [模块版本号](../DevelopingModules/ModuleVersionNnumbering)。

## Getting a specific commit using a repository identifier

​	你可以使用 `go get` 命令来添加某个模块的未发布代码，这些代码来自其版本库中的某个特定提交。

​	要做到这一点，你需要使用 `go get` 命令，用 `@` 符号指定你想要的代码。当你使用 `go get` 命令时，该命令将在你的 `go.mod` 文件中添加一个 `require` 指令，该指令需要外部模块，使用基于提交细节的伪版本号。

​	下面的例子提供了一些说明。它们基于源代码位于 git 版本库中的模块。

- 要获得特定提交的模块，请附加`@commithash`的形式：

  ```shell
  $ go get example.com/theirmodule@4cf76c2
  ```

- 要获得特定分支的模块，请添加`@branchname`的形式：

  ```shell
  $ go get example.com/theirmodule@bugfixes
  ```

## Removing a dependency 移除一个依赖项

​	当你的代码不再使用某个模块的任何包时，你可以停止跟踪该模块作为一个依赖项。

​	要停止跟踪所有未使用的模块，运行[go mod tidy](../../References/CommandDocumentation/go#add-missing-and-remove-unused-modules)命令。这个命令还可能添加（构建模块中的包）所缺失的依赖。

```shell
$ go mod tidy
```

​	要移除一个特定的依赖项，使用[go get](../../References/CommandDocumentation/go#add-dependencies-to-current-module-and-install-them)命令，指定模块的模块路径并附加`@none`，如下面的例子：

```shell
$ go get example.com/theirmodule@none
```

​	`go get`命令也会降级或移除依赖于被移除模块的其他依赖项。

## Specifying a module proxy server 指定一个模块代理服务器

​	当你使用Go工具来处理模块时，这些工具默认从`proxy.golang.org`（一个由Google运营的公共模块镜像）或直接从模块的版本库下载模块。你可以指定Go工具应该使用另一个代理服务器来下载和验证模块。

​	如果你（或你的团队）已经建立或选择了要使用的不同的模块代理服务器，则可能需要执行此操作。例如，有些人建立了一个模块代理服务器，以便更好地控制依赖项的使用方式。

​	为了指定另一个模块代理服务器供Go工具使用，将`GOPROXY`环境变量设置为一个或多个服务器的URL。**Go工具将按照你指定的顺序尝试每个URL**。默认情况下，`GOPROXY`首先指定一个公共的Google运行的模块代理，然后直接从模块库中下载（如其模块路径中指定的）：

```
GOPROXY="https://proxy.golang.org,direct"
```

​	关于`GOPROXY`环境变量的更多信息，包括支持其他行为的值，见[go.mod file reference（go.mod文件参考）](../../References/gomodFileReference)。

​	你可以将该变量设置为其他模块代理服务器的URL，用**逗号或管道**来分隔URL。

- 当你使用**逗号**时，Go工具只有在当前URL返回HTTP 404或410时才会尝试列表中的下一个URL。

  ```
  GOPROXY="https://proxy.example.com,https://proxy2.example.com"
  ```

- 当你使用管道时，Go工具将尝试列表中的下一个URL，无论HTTP错误代码如何。

  ```
  GOPROXY="https://proxy.example.com|https://proxy2.example.com"
  ```

​	Go模块经常在版本控制服务器和模块代理上开发和发布，而这些服务器和代理在公共互联网上是不存在的。你可以设置`GOPRIVATE`环境变量。你可以设置`GOPRIVATE`环境变量来配置`go`命令从私人来源下载和构建模块。然后`go`命令就可以从私人来源下载和构建模块了。

`GOPRIVATE`或`GONOPROXY`环境变量可以被设置为匹配模块前缀的`glob`模式列表，这些模块前缀是私有的，不应该从任何代理处请求。例如：

```
GOPRIVATE=*.corp.example.com,*.research.example.com
```