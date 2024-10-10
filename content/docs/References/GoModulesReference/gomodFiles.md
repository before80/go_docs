+++
title = "go.mod 文件"
date = 2023-05-17T09:59:21+08:00
weight = 3
description = ""
isCJKLanguage = true
draft = false
+++
## go.mod files  - go.mod 文件

> 原文：[https://go.dev/ref/mod#go-mod-file](https://go.dev/ref/mod#go-mod-file)

​	模块由其根目录下的名为`go.mod`的UTF-8编码的文本文件来定义。`go.mod`文件是以行为单位的。每一行都有一个单独指令，由关键字后跟参数组成。例如：

```
module example.com/my/thing

go 1.12

require example.com/other/thing v1.0.2
require example.com/new/thing/v2 v2.3.4
exclude example.com/old/thing v1.2.3
replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
retract [v1.9.0, v1.9.5]
```

​	可以将前导关键字从相邻的行中分解出来，以创建一个块，就像Go导入中一样。

```
require (
    example.com/new/thing/v2 v2.3.4
    example.com/old/thing v1.2.3
)
```

​	`go.mod`文件被设计成人类可读，机器可写。`go`命令提供了几个改变`go.mod`文件的子命令。例如，[go get](../Module-awareCommands#go-get)可以升级或降级特定的依赖项。加载模块图的命令会在需要时[automatically update（自动更新）](#automatic-updates)`go.mod`。[go mod edit](../Module-awareCommands#go-mod-edit)可以执行低级别的编辑。Go程序可以使用`golang.org/x/mod/modfile`包，以编程方式进行同样的修改。

​	[main module（主模块）](../Glossary#main-module)和任何用本地文件路径指定的[replacement module（替换模块）](#replace-directive)都需要`go.mod`文件。然而，缺少显式的`go.mod`文件的模块仍然可以作为依赖项被[要求](#require-directive)，或者作为用模块路径和版本指定的替换模块使用；参见[Compatibility with non-module repositories（与非模块存储库的兼容性）](../CompatibilityWithNon-moduleRepositories)。

### Lexical elements 词法元素

​	当`go.mod`文件被解析时，其内容被分解为一连串的标记。有几种标记：空白（whitespace）、注释（comments）、标点符号（punctuation）、关键字（keywords）、标识符（identifiers）和字符串（strings）。

​	空白（White space）包括空格（U+0020）、制表符（U+0009）、回车符（U+000D）和换行符（U+000A）。除换行符外，其他空白字符没有任何作用，只是将原本合并的标记分开。换行符是重要的标记。

​	注释（Comments）以`//`开始并运行到行尾。`/* */` 的注释是不允许的。

​	标点符号（Punctuation）包括`(`、`)`和`=>`。

​	关键字（Keywords）用于区分`go.mod`文件中不同类型的指令。允许的关键字有`module`、`go`、`require`、`replace`、`exclude`和`retract`。

​	标识符（Identifiers）是由非空白字符组成的序列，例如模块路径或语义版本。

​	字符串（Strings）是有引号的字符序列。有两种字符串：以引号（"，U+0022）开头和结尾的解释字符串（interpreted strings）和以重音符号（`，U+0060）开头和结尾的原始字符串（raw strings）。被解释的字符串可以包含由反斜线（\, U+005C）和其他字符组成的转义序列。转义的引号（\\"）不会终止解释字符串。解释字符串（interpreted strings）的无引号值是引号之间的字符序列，每个转义序列被反斜线后面的字符取代（例如，\\"被 "取代，\n被n取代）。相比之下，原始字符串的无引号值只是重音符号之间的字符序列；反斜杠在原始字符串中没有特殊意义。

​	在`go.mod`语法中，标识符和字符串是可以互换的。

### Module paths and versions 模块路径和版本

​	`go.mod`文件中的大多数标识符和字符串都是模块路径或版本。

​	模块路径必须满足以下要求：

- 路径必须由一个或多个路径元素组成，以斜线（`/`，U+002F）分隔。路径不能以斜线开始或结束。
- 每个路径元素是一个非空字符串，由ASCII字母、ASCII数字和有限的ASCII标点符号（`-`、`.`、`_`和`~`）组成。
- 路径元素不能以点（`.`，U+002E）开始或结束。
- 路径元素前缀到第一个点不能是Windows上的保留文件名，无论大小写（`CON`、`com1`、`NuL`等）。
- 路径元素前缀到第一个点不能以一个波浪符号后跟一个或多个数字结尾（如`EXAMPL~1.COM`）。

​	如果模块路径出现在`require`指令中并且没有被替换，或者模块路径出现在`replace`指令的右侧，`go`命令可能需要下载该路径的模块，并且必须满足一些额外的要求。

​	（a）按照惯例，域名的前导路径元素(直到第一个斜线)必须只包含小写的 ASCII 字母、 ASCII 数字、点(`.`，U+002E)和破折号(`-` ，U+002D) ；它必须包含至少一个点，并且不能以破折号开始。

​	（b）对于形式为`/vN`的最后路径元素，其中`N`看起来是数字（ASCII数字和点），`N`不能以前导零开始，不能是`/v1`，也不能包含任何点。

- 对于以`gopkg.in/`开头的路径，这个要求被一个路径必须遵循 [gopkg.in](https://gopkg.in/)服务约定的要求所取代。

  

​	`go.mod`文件中的版本可以是[经典（canonical）](../Glossary#canonical version)的或非经典的。

​	经典版本以字母`v`开头，后面是符合[Semantic Versioning 2.0.0（语义化版本2.0.0）](https://semver.org/lang/zh-CN/)规范的语义版本。更多信息请参见 [Versions（版本）](../ModulesPackagesAndVersions)。

​	大多数其他标识符和字符串都可以作为非经典版本使用，但是也有一些限制，以避免文件系统、存储库和[module proxies（模块代理）](../Glossary#module-proxy)出现问题。非经典版本只允许在主模块的`go.mod`文件中使用。`go`命令在[自动更新](#automatic-updates)`go.mod`文件时，会尝试用一个等效的经典版本来替换每个非经典版本。

​	在模块路径与版本相关的地方（如`require`、`replace`和`exclude`指令），最后的路径元素必须与版本一致。参见[Major version suffixes（主版本后缀）](../ModulesPackagesAndVersions#major-version-suffixes)。

### Grammar 语法

​	`go.mod`的语法是使用Extended Backus-Naur Form (EBNF)指定的。关于EBNF语法的详细信息，请参见[Notation section in the Go Language Specification（Go语言规范中的标记法部分）](../../LanguageSpecification/Notation)。

```
GoMod = { Directive } .
Directive = ModuleDirective |
            GoDirective |
            RequireDirective |
            ExcludeDirective |
            ReplaceDirective |
            RetractDirective .
```

​	换行符、标识符和字符串分别用`newline`、`ident`和`string`表示。

​	模块路径和版本用`ModulePath`和`Version`来表示。

```
ModulePath = ident | string . /* see restrictions above */
Version = ident | string .    /* see restrictions above */
```

### module directive

​	`module`指令定义了主模块的[path（路径）](../Glossary#path)。`go.mod`文件必须恰好包含一个`module`指令。

```
ModuleDirective = "module" ( ModulePath | "(" newline ModulePath newline ")" ) newline .
```

示例：

```
module golang.org/x/net
```

#### Deprecation 废弃

​	模块可以在段落开头包含字符串 `Deprecated:`（区分大小写）的注释块中标记为已弃用。 废弃信息从冒号之后开始，一直到段落的末尾。注释可以出现在`module`指令的前面，也可以出现在同一行的后面。 

示例：

```
// Deprecated: use example.com/mod/v2 instead.
module example.com/mod
```

​	从Go 1.17开始，[go list -m -u](../Module-awareCommands#go-list-m-u)检查[build list（构建列表）](../Glossary#build-list)中所有已废弃模块的信息。[go get](../Module-awareCommands#go-get)检查构建命令行上命名的包所需的已废弃模块。

​	当`go`命令检索一个模块的废弃信息时，它从匹配`@latest`[version query（版本查询）](../Module-awareCommands#version-queries)的版本中加载`go.mod`文件，而不考虑[retractions（撤回）](#retract-directive)或[exclusions（排除）](#exclude-directive)。`go`命令从同一个`go.mod`文件中加载[retracted versions（撤回版本）](../Glossary#retracted-version)的列表。

​	为了废弃一个模块，作者可以添加一个`// Deprecated:`注释并标记一个新版本。作者可以在更高的版本中更改或删除废弃信息。

​	废弃适用于一个模块的所有次版本。高于`v2`的主版本被认为是独立的模块，因为它们的[major version suffixes（主版本后缀）](../Glossary#major-version-suffix)赋予它们不同的模块路径。

Deprecation messages are intended to inform users that the module is no longer supported and to provide migration instructions, for example, to the latest major version. Individual minor and patch versions cannot be deprecated; [`retract`](https://go.dev/ref/mod#go-mod-file-retract) may be more appropriate for that.

​	废弃信息的目的是通知用户，该模块不再被支持，并提供迁移说明，例如，迁移到最新的主版本。**单个次版本和修订版本不能被废弃（=>仍有疑问？？）**；[retract](#retract-directive)可能更适合于这种情况。

### go directive

​	`go`指令表明一个模块是以给定的Go版本的语义为基础编写的。版本必须是有效的Go发布版本：一个正整数，后面跟一个点和一个非负整数（例如，`1.9`，`1.14`）。

​	`go`指令最初是为了支持Go语言的向后不兼容的变化（见[Go 2 过渡](https://go.googlesource.com/proposal/+/master/design/28221-go2-transitions.md)）。自从引入模块以来，没有任何不兼容的语言变化，但`go`指令仍然影响到新语言特性的使用：

- 对于模块中的包，编译器会拒绝使用`go`指令指定的版本之后引入的语言特性。例如，如果模块的指令是`go 1.12`，它的包就不能使用`1_000_000`这样的数字字面量，这是在`Go 1.13`引入的。
- 如果较旧的Go版本构建了该模块的一个包并遇到了编译错误，那么错误就会指出该模块是为一个较新的Go版本编写的。例如，假设模块有`go 1.13`，一个包使用数字字面`1_000_000`。如果该包是用`Go 1.12`构建的，编译器就会注意到该代码是为`Go 1.13`编写的。

​	此外，`go`命令会根据`go`指令所指定的版本改变其行为。这有以下影响：

（a）在`go 1.14`或更高版本中，可以启用自动[vendoring](../Module-awareCommands#vendoring)。如果文件`vendor/modules.txt`存在并与`go.mod`一致，就不需要显式使用`-mod=vendor`标志。

（b）在 `go 1.16` 或更高版本中，`all`包模式只匹配由[main module（主模块）](../Glossary#main-module)中由包和测试过渡导入的包。这也是[go mod vendor](../Module-awareCommands#go-mod-vendor)自引入模块以来所保留的包的集合。在较低的版本中，`all`也包括由主模块中的包导入的包的测试、对这些包的测试等等。

（c）在`go 1.17`或更高版本：

- `go.mod`文件包括一个明确的[require指令](#require-directive)，该指令提供由主模块中的包或测试过渡地导入的任何包。(在 `go 1.16` 或更低版本，只有在[minimal version selection（最小版本选择）](../MVS)会选择不同版本的情况下，才会包含[indirect dependency（间接依赖）](../Glossary#direct-dependency)。) 这个额外的信息使得[module graph pruning（模块图的修剪）](../ModuleGraphPruning)和[lazy module loading（延迟模块加载）](../ModuleGraphPruning#lazy-module-loading)成为可能。

-  由于`// indirect`依赖可能比以前的`go`版本多得多，间接依赖被记录在`go.mod`文件中的一个独立块中。

- `go mod vendor`省略了`go.mod`和`go.sum`文件中的供应商依赖项。(这允许在`vendor`的子目录中调用`go`命令来识别正确的主模块）。

- `go mod vendor`将 go 版本从每个依赖项的 go.mod 文件中记录下来，并放在`vendor/modules.txt` 中。

  

​	`go.mod`文件最多可以包含一个`go`指令。如果没有`go`指令，大多数命令会添加一个当前Go版本的`go`指令。

​	在Go 1.17发行版中，如果`go`指令缺失，则假定`go 1.16`。

```
GoDirective = "go" GoVersion newline .
GoVersion = string | ident .  /* valid release version; see above */
```

示例：

```
go 1.14
```

### require directive

​	`require`指令声明了给定模块依赖的最低要求版本。对于每个所需的模块版本，`go`命令加载该版本的`go.mod`文件，并将该文件中的requirements 纳入其中。一旦所有requirements 被加载完，`go`命令就会使用[minimal version selection（最小版本选择 MVS）](../MVS) 来解析它们，从而产生[build list（构建列表）](../Glossary#build-list)。

​	`go`命令自动为一些requirements 添加`// indirect`注释。`// indirect`注释表示所需模块的任何包都没有被[main module（主模块）](../Glossary#main-module)中的任何包直接导入。

​	如果[go directive](#go-directive)指定了`go 1.16`或更低的版本，当所选模块的版本高于主模块的其他依赖项已经暗示（过渡地）的版本时，`go`命令会添加一个间接需求。这可能是由于显式的升级（`go get -u ./...`），移除之前施加需求的其他依赖项（`go mod tidy`），或者依赖项导入的包在其自身的`go.mod`文件中没有相应的 requirement （比如一个完全没有`go.mod`文件的依赖项）。

​	在`go 1.17`及以上版本中，`go`命令为每个模块增加了一个间接需求，提供任何被主模块中的包或测试导入（即使是[间接的](../Glossary#indirect-dependency)）或作为参数传递给`go get`的包。这些更全面的 requirements 可以支持[module graph pruning（模块图的修剪）](../ModuleGraphPruning)和[lazy module loading（延迟模块加载）](../ModuleGraphPruning#lazy-module-loading)。

```
RequireDirective = "require" ( RequireSpec | "(" newline { RequireSpec } ")" newline ) .
RequireSpec = ModulePath Version newline .
```

示例：

```
require golang.org/x/net v1.2.3

require (
    golang.org/x/crypto v1.4.5 // indirect
    golang.org/x/text v1.6.7
)
```

### exclude directive

​	`exclude`指令可以防止模块版本被`go`命令加载。

​	从Go 1.16开始，如果任何`go.mod`文件中的`require`指令所引用的版本被主模块的`go.mod`文件中的`exclude`指令所排除，该requirement 将被忽略。这可能会导致像[go get](../Module-awareCommands#go-get)和[go mod tidy](../Module-awareCommands#go-mod-tidy)这样的命令在`go.mod`中添加更高版本的新requirements ，如果合适的话，会加上一个`// indirect`注释。

​	在Go 1.16之前，如果一个排除的版本被`require`指令引用，`go`命令会列出该模块的可用版本（如[go list -m -versions](../Module-awareCommands#go-list-m)所示）并加载下一个较高的非排除版本。这可能会导致不确定的版本选择，因为下一个更高版本可能会随着时间的推移而改变。为了这个目的，发布版和预布行版都被考虑了，但伪版本没有被考虑。如果没有更高的版本，`go`命令会报告一个错误。

​	`exclude`指令只适用于主模块的`go.mod`文件，在其他模块中被忽略。详见[Minimal version selection（最小版本选择 MVS）](../MVS)。

```
ExcludeDirective = "exclude" ( ExcludeSpec | "(" newline { ExcludeSpec } ")" newline ) .
ExcludeSpec = ModulePath Version newline .
```

示例：

```
exclude golang.org/x/net v1.2.3

exclude (
    golang.org/x/crypto v1.4.5
    golang.org/x/text v1.6.7
)
```

### replace directive

​	`replace`指令用其他地方找到的内容替换某个模块的特定版本或某个模块的所有版本。替换可以指定另一个模块路径和版本，或者一个特定平台的文件路径。

​	如果一个版本出现在箭头（`=>`）的左侧，那么只有该模块的特定版本被替换，其他版本将被正常访问。如果左侧的版本被省略，则模块的所有版本都被替换。

​	如果箭头右侧的路径是一个绝对或相对路径（以`./`或`./`开头），它被解释为替换模块根目录的本地文件路径，其中必须包含一个`go.mod`文件。在这种情况下，替换版本必须被省略。

​	如果右边的路径不是本地路径，它必须是一个有效的模块路径。在这种情况下，需要一个版本号。同一模块的版本不能同时出现在构建列表中。

​	不管替换是用本地路径还是模块路径指定的，如果替换的模块有一个`go.mod`文件，它的`module`指令必须与它所替换的模块路径一致。

​	`replace`指令只适用于主模块的`go.mod`文件，在其他模块中被忽略。详见[Minimal version selection（最小版本选择 MVS）](../MVS)。

​	如果有多个主模块，所有主模块的`go.mod`文件都适用。不允许主模块间的`replace`指令发生冲突，必须在[go.work 文件的替换](../Workspaces#replace-directive)中删除或重写这些指令。

​	请注意，仅仅是`replace`指令并不能将一个模块添加到[模块图](../Glossary#module-graph)中。在主模块的 `go.mod` 文件或依赖模块的 `go.mod` 文件中，还需要一个指向被替换模块版本的 [require 指令](#require-directive)。如果不需要左侧的模块版本，`replace`指令就没有作用。

```
ReplaceDirective = "replace" ( ReplaceSpec | "(" newline { ReplaceSpec } ")" newline ) .
ReplaceSpec = ModulePath [ Version ] "=>" FilePath newline
            | ModulePath [ Version ] "=>" ModulePath Version newline .
FilePath = /* platform-specific relative or absolute file path */
```

示例：

```
replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5

replace (
    golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5
    golang.org/x/net => example.com/fork/net v1.4.5
    golang.org/x/net v1.2.3 => ./fork/net
    golang.org/x/net => ./fork/net
)
```

### retract directive

​	`retract`指令表示由`go.mod`定义的模块的某个版本或一系列版本不应该被依赖。当版本过早发布或在发布后发现严重问题时，`retract`指令就很有用。撤回的版本应该在版本控制存储库和[module proxy（模块代理）](../Glossary#module-proxy)上保持可用，以确保依赖它们的构建不会被破坏。`retract`这个词是从学术文献中借来的：被撤回的研究论文仍然可以使用，但它有问题，不应该成为未来工作的基础。

​	当一个模块的版本被撤回时，用户将不会使用[go get](../Module-awareCommands#go-get)、[go mod tidy](../Module-awareCommands#go-mod-tidy)或其他命令自动升级到该版本。依赖于撤消版本的构建应该继续工作，但是当用户用[go list -m -u](../Module-awareCommands#go-list-m)检查更新或用[go get](../Module-awareCommands#go-get)更新相关模块时，将会收到撤回的通知。

​	要撤回一个版本，模块作者应该在`go.mod`中添加一个`retract`指令，然后发布一个包含该指令的新版本。新版本必须高于其他发布或预发布的版本；也就是说，在考虑撤回之前，`@latest`[version query（版本查询）](../Module-awareCommands#version-query)应该解析到新版本。`go`命令从`go list -m -retracted $modpath@latest`（其中`$modpath`是模块路径）显示的版本中加载和应用撤回。

​	除非使用 `-retracted` 标志，否则撤回的版本会从 [go list -m -versions](../Module-awareCommands#go-list-m) 打印的版本列表中隐藏。在解析版本查询（如`@>=v1.2.3`或`@latest`）时，撤回的版本将被排除。

​	包含撤回的版本可以自行撤回。如果一个模块的最高发布版本或预发布版本自行撤回了，`@latest`查询会在排除撤回的版本后解析到一个较低的版本。

As an example, consider a case where the author of module `example.com/m` publishes version `v1.0.0` accidentally. To prevent users from upgrading to `v1.0.0`, the author can add two `retract` directives to `go.mod`, then tag `v1.0.1` with the retractions.

​	举个例子，考虑这样一种情况：模块 `example.com/m` 的作者意外地发布了 `v1.0.0` 版本。为了防止用户升级到`v1.0.0`，作者可以在`go.mod`中添加两个`retract`指令，然后用撤回标记`v1.0.1`。

```
retract (
    v1.0.0 // Published accidentally.
    v1.0.1 // Contains retractions only.
)
```

​	当用户运行`go get example.com/m@latest`时，`go`命令会读取`v1.0.1`版本的撤回，这就是现在的最高版本。`v1.0.0`和`v1.0.1`都被撤回了，所以`go`命令会升级（或降级！）到下一个最高版本，也许是`v0.9.5`。

​	`retract`指令既可以写成单一版本（如`v1.0.0`），也可以写成有上下限的封闭的版本区间，用`[` 和 `]` 界定（如`[v1.1.0, v1.2.0]`）。单一版本等同于一个上界和下界相同的区间。和其他指令一样，多个`retract`指令可以组合在一起，以`(`为界，放在一行的末尾，以`)`为界，放在单独的一行。

​	每条`retract`指令都应该有一个注释，解释撤回的理由，尽管这并不是强制性的。`go`命令可以在关于撤回版本的警告和`go list`输出中显示理由注释。理由注释可以紧接着写在`retract`指令的上方（中间没有空行），也可以写在同一行的后面。如果一条注释出现在一个块的上方，它适用于该块内所有没有自己注释的`retract`指令。一个理由注释可以跨越多行。

```
RetractDirective = "retract" ( RetractSpec | "(" newline { RetractSpec } ")" newline ) .
RetractSpec = ( Version | "[" Version "," Version "]" ) newline .
```

示例：

- 撤回`v1.0.0`到`v1.9.9`之间的所有版本：

```
retract v1.0.0
retract [v1.0.0, v1.9.9]
retract (
    v1.0.0
    [v1.0.0, v1.9.9]
)
```

- 在提前发布了`v1.0.0`版本之后，返回到 unversioned：

```
retract [v0.0.0, v1.0.1] // assuming v1.0.1 contains this retraction.
```

- 抹去包含所有伪版本和标签版本的模块：

```
retract [v0.0.0-0, v0.15.2]  // assuming v0.15.2 contains this retraction.
```

​	`retract`指令是在Go 1.16中添加的。如果[主模块](../Glossary#main-module)的`go.mod`文件中写有`retract`指令，Go 1.15及以下版本将报告错误，并忽略依赖模块的`go.mod`文件中的`retract`指令。

### Automatic updates 自动更新

​	如果`go.mod`缺少信息或不能准确反映现实，大多数命令都会报告错误。[go get](../Module-awareCommands#go-get)和[go mod tidy](../Module-awareCommands#go-mod-tidy)命令可以用来修复大多数这些问题。此外，`-mod=mod`标志可用于大多数模块感知命令（`go build`、`go test`等），指示`go`命令自动修复`go.mod`和`go.sum`中的问题。

例如，考虑这个`go.mod`文件：

```
module example.com/M

go 1.16

require (
    example.com/A v1
    example.com/B v1.0.0
    example.com/C v1.0.0
    example.com/D v1.2.3
    example.com/E dev
)

exclude example.com/D v1.2.3
```

​	用 `-mod=mod` 触发的更新将非经典版本标识符重写为[canonical（经典）](../Glossary#canonical-version)的 semver 形式，因此 `example.com/A` 的 `v1` 变成了 `v1.0.0`，而 `example.com/E` 的 `dev` 变成了 `dev` 分支上最新提交的伪版本，可能是 `v0.0.0-20180523231146-b3f5c0f6e5f1`。

​	该更新修改了requirements 以尊重排除项，因此对排除项 `example.com/D v1.2.3` 的 requirements 被更新为使用 `example.com/D` 的下一个可用版本，可能是 `v1.2.4` 或 `v1.3.0`。

​	该更新移除了多余的或误导性的 requirements。例如，如果 `example.com/A v1.0.0` 本身需要 `example.com/B v1.2.0` 和 `example.com/C v1.0.0`，那么 `go.mod` 对 `example.com/B v1.0.0` 的 requirement 是误导性的（被 `example.com/A` 对 `v1.2.0` 的 requirement 所取代），而它对 `example.com/C v1.0.0` 的 requirement 是多余的（被 `example.com/A` 对相同版本的需求暗示），因此两者都将被移除。如果主模块包含直接从`example.com/B`或`example.com/C`导入的包，那么requirement 将被保留，但会更新为实际使用的版本。

​	最后，该更新会将`go.mod`改成经典格式，这样未来的机械更改将导致最小的差异。如果只需要更改格式，`go`命令将不更新`go.mod`。

​	因为模块图定义了 import 语句的含义，任何加载包的命令也会使用`go.mod`，因此可以对其进行更新的（命令），包括`go build`、`go get`、`go install`、`go list`、`go test`、`go mod tidy`。

​	在Go 1.15及以下版本中，`-mod=mod`标志是默认启用的，所以更新会自动进行。从Go 1.16开始，`go`命令的行为就像设置了`-mod=readonly`一样：如果需要对`go.mod`进行任何更改，`go`命令会报告一个错误并建议进行修复。