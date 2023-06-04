+++
title = "工作区"
date = 2023-05-17T09:59:21+08:00
weight = 6
description = ""
isCJKLanguage = true
draft = false
+++
## Workspaces 工作区

> 原文：[https://go.dev/ref/mod#workspaces](https://go.dev/ref/mod#workspaces)

​	工作区是磁盘上模块的集合，在运行[minimal version selection (MVS) （最小版本选择（MVS））](../MVS)时这些模块被用作主模块。

​	工作区可以在 [go.work 文件](#go-work-file)中声明，该文件指定了工作区中每个模块目录的相对路径。当没有`go.work`文件存在时，工作区由包含当前目录的单个模块组成。

​	大多数处理模块的`go`子命令都是在由当前工作区决定的模块集合上操作的。`go mod init`、`go mod why`、`go mod edit`、`go mod tidy`、`go mod vendor`和`go get`总是在一个主模块上操作。

​	命令首先通过检查`GOWORK`环境变量来确定它是否处于工作区上下文中。如果`GOWORK`被设置为`off`，该命令将处于单模块上下文中。如果它是空的或者未提供，命令将在当前工作目录中搜索，然后在连续的父目录中搜索`go.work`这个文件。如果找到一个（`go.work`）文件，该命令将在该文件定义的工作区中操作；否则，工作区将只包括包含工作目录的模块。如果`GOWORK`命名一个以`.work`结尾的现有文件的路径，工作区模式将被启用。任何其他的值都是一个错误。您可以使用`go env GOWORK`命令来确定`go`命令正在使用哪个`go.work`文件。如果`go`命令没有进入工作区模式，`go env GOWORK`将为空。

### go.work files

​	工作区是由一个名为`go.work`的UTF-8编码文本文件定义的。`go.work`文件是面向行的。每行包含一个指令，由一个关键字和参数组成。例如：

```
go 1.18

use ./my/first/thing
use ./my/second/thing

replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
```

​	与`go.mod`文件一样，前导关键字可以从相邻的行中分解出来，形成一个块。

```
use (
    ./my/first/thing
    ./my/second/thing
)
```

​	`go`命令提供了几个操作`go.work`文件的子命令。[go work init](../Module-awareCommands#go-work-init)创建新的`go.work`文件。[go work use](../Module-awareCommands#go-work-use)向`go.work`文件添加模块目录。[go work edit](../Module-awareCommands#go-work-edit)执行低级别的编辑。Go 程序可以使用 [golang.org/x/mod/modfile](https://pkg.go.dev/golang.org/x/mod/modfile?tab=doc) 包，以编程方式进行相同的更改。

### Lexical elements 词汇元素

​	`go.work`文件中的词汇元素的定义方式与[go.mod 文件](../gomodFiles#lexical-elements)完全相同。

### Grammar 语法

​	`go.work`的语法由下面使用Extended Backus-Naur Form (EBNF)来指定。有关 EBNF 语法的详细信息，请参见 [Go 语言规范中的标记法部分](../../LanguageSpecification/Notation)。

```
GoWork = { Directive } .
Directive = GoDirective |
            UseDirective |
            ReplaceDirective .
```

​	换行符、标识符和字符串分别用`newline`、`ident`和`string`表示。

​	模块路径和版本用`ModulePath`和`Version`来表示。模块路径和版本的指定方式与[go.mod文件](../gomodFiles#lexical-elements)的指定方式完全相同。

```
ModulePath = ident | string . /* see restrictions above */
Version = ident | string .    /* see restrictions above */
```

### go directive

​	在一个有效的`go.work`文件中需要一个`go`指令。版本必须是有效的Go发布版本：一个正整数后跟一个点和一个非负整数（例如，`1.18`，`1.19`）。

​	`go`指令表示`go.work`文件所要使用的go工具链版本。如果`go.work`文件的格式发生了变化，未来版本的工具链将根据其指示的版本来解释该文件。

​	一个`go.work`文件最多只能包含一个`go`指令。

```
GoDirective = "go" GoVersion newline .
GoVersion = string | ident .  /* valid release version; see above */
```

示例：

```
go 1.18
```

### use directive

​	`use`将磁盘上的一个模块添加到工作区的主模块集合中。它的参数是包含该模块的`go.mod`文件的目录的相对路径。`use`指令并不添加包含在其参数目录下的子目录中的模块。这些模块可以由包含其`go.mod`文件的目录在单独的`use`指令中添加。=>仍有疑问？？这里应该是"`require`指令中添加吧"？？

```
UseDirective = "use" ( UseSpec | "(" newline { UseSpec } ")" newline ) .
UseSpec = FilePath newline .
FilePath = /* platform-specific relative or absolute file path */
```

示例：

```
use ./mymod  // example.com/mymod

use (
    ../othermod
    ./subdir/thirdmod
)
```

### replace directive

​	与`go.mod`文件中的`replace`指令类似，`go.work`文件中的`replace`指令用其他地方的内容替换一个模块的特定版本，或一个模块的所有版本。`go.work`中的通配符替换可以覆盖`go.mod`文件中特定版本的`replace`。

​	`go.work`文件中的`replace`指令会覆盖工作区模块中相同模块或模块版本的任何替换。

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