+++
title = "go fmt 你的代码"
weight = 18
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# go fmt your code - go fmt 你的代码

https://go.dev/blog/gofmt

Andrew Gerrand
23 January 2013

## Introduction 简介

[Gofmt](https://go.dev/cmd/gofmt/) is a tool that automatically formats Go source code.

Gofmt是一个可以自动格式化Go源代码的工具。



Gofmt’d code is:

使用 Gofmt 的代码是：

- easier to **write**: never worry about minor formatting concerns while hacking away, 更加容易编写：不用担心小的格式问题。
- easier to **read**: when all code looks the same you need not mentally convert others' formatting style into something you can understand. 更容易阅读：当所有的代码看起来都一样时，你不需要在头脑中把别人的格式化风格转换成你能理解的东西。
- easier to **maintain**: mechanical changes to the source don’t cause unrelated changes to the file’s formatting; diffs show only the real changes. 更容易维护：对源代码的机械性修改不会导致对文件格式的不相关修改；差异只显示真正的修改。
- **uncontroversial**: never have a debate about spacing or brace position ever again! 无争议：再也不会有关于间距或括号位置的争论了!

## Format your code 格式化你的代码

We recently conducted a survey of Go packages in the wild and found that about 70% of them are formatted according to gofmt’s rules. This was more than expected - and thanks to everyone who uses gofmt - but it would be great to close the gap.

我们最近对野外的Go包进行了一次调查，发现大约70%的Go包是按照gofmt的规则进行格式化的。这比预期的要多--感谢所有使用gofmt的人--但如果能缩小这个差距就更好了。

To format your code, you can use the gofmt tool directly:

要格式化你的代码，你可以直接使用gofmt工具：

```
gofmt -w yourcode.go
```

Or you can use the “[go fmt](https://go.dev/cmd/go/#hdr-Gofmt__reformat__package_sources)” command:

或者你可以使用 "go fmt "命令：

```
go fmt path/to/your/package
```

To help keep your code in the canonical style, the Go repository contains hooks for editors and version control systems that make it easy to run gofmt on your code.

为了帮助你的代码保持规范的风格，Go 仓库包含了编辑器和版本控制系统的钩子，使你可以轻松地在代码上运行 gofmt。

For Vim users, the [Vim plugin for Go](https://github.com/fatih/vim-go) includes the :Fmt command that runs gofmt on the current buffer.

对于 Vim 用户，Vim 的 Go 插件包括 :Fmt 命令，可以在当前缓冲区运行 gofmt。

For emacs users, [go-mode.el](https://github.com/dominikh/go-mode.el) provides a gofmt-before-save hook that can be installed by adding this line to your .emacs file:

对于emacs用户，go-mode.el提供了一个gofmt-before-save钩子，可以通过在你的.emacs文件中添加这一行来安装：

```
(add-hook 'before-save-hook #'gofmt-before-save)
```

For Eclipse or Sublime Text users, the [GoClipse](https://github.com/GoClipse/goclipse) and [GoSublime](https://github.com/DisposaBoy/GoSublime) projects add a gofmt facility to those editors.

对于Eclipse或Sublime Text用户，GoClipse和GoSublime项目为这些编辑器添加了gofmt工具。

And for Git aficionados, the [misc/git/pre-commit script](https://github.com/golang/go/blob/release-branch.go1.1/misc/git/pre-commit) is a pre-commit hook that prevents incorrectly-formatted Go code from being committed. If you use Mercurial, the [hgstyle plugin](https://bitbucket.org/fhs/hgstyle/overview) provides a gofmt pre-commit hook.

对于Git爱好者来说，misc/git/pre-commit脚本是一个预提交钩子，可以防止格式不正确的Go代码被提交。如果你使用Mercurial，hgstyle插件提供了一个gofmt预提交钩子。

## Mechanical source transformation 机械源码转换

One of the greatest virtues of machine-formatted code is that it can be transformed mechanically without generating unrelated formatting noise in the diffs. Mechanical transformation is invaluable when working with large code bases, as it is both more comprehensive and less error prone than making wide-sweeping changes by hand. Indeed, when working at scale (like we do at Google) it often isn’t practical to make these kinds of changes manually.

机械格式化的代码最大的优点之一是它可以被机械地转换，而不会在差异中产生无关的格式化噪音。在处理大型代码库时，机械转换是非常有价值的，因为它比手工进行大范围的修改更全面，更不容易出错。事实上，当大规模工作时（就像我们在谷歌做的那样），手动进行这类修改往往是不现实的。

The easiest way to mechanically manipulate Go code is with gofmt’s -r flag. The flag specifies a rewrite rule of the form

机械地处理Go代码的最简单方法是使用gofmt的-r标志。该标志指定了一个重写规则，其形式为

```
pattern -> replacement
```

where both pattern and replacement are valid Go expressions. In the pattern, single-character lowercase identifiers serve as wildcards matching arbitrary sub-expressions, and those expressions are substituted for the same identifiers in the replacement.

其中模式和替换都是有效的围棋表达式。在模式中，单字符小写标识符作为通配符匹配任意的子表达式，而这些表达式在替换中被替换为相同的标识符。

For example, this[ recent change](https://go.dev/cl/7038051) to the Go core rewrote some uses of [bytes.Compare](https://go.dev/pkg/bytes/#Compare) to use the more efficient [bytes.Equal](https://go.dev/pkg/bytes/#Equal). The contributor made the change using just two gofmt invocations:

例如，最近对Go核心的这一修改，重写了byte.Compare的一些用法，以使用更有效的byte.Equal。贡献者只用了两个gofmt的调用就完成了这个改变。

```
gofmt -r 'bytes.Compare(a, b) == 0 -> bytes.Equal(a, b)'
gofmt -r 'bytes.Compare(a, b) != 0 -> !bytes.Equal(a, b)'
```

Gofmt also enables [gofix](https://go.dev/cmd/fix/), which can make arbitrarily complex source transformations. Gofix was an invaluable tool during the early days when we regularly made breaking changes to the language and libraries. For example, before Go 1 the built-in error interface didn’t exist and the convention was to use the os.Error type. When we [introduced error](https://go.dev/doc/go1.html#errors), we provided a gofix module that rewrote all references to os.Error and its associated helper functions to use error and the new [errors package](https://go.dev/pkg/errors/). It would have been daunting to attempt by hand, but with the code in a standard format it was relatively easy to prepare, execute, and review this change which touched almost all Go code in existence.

Gofmt还可以启用gofix，它可以进行任意复杂的源转换。在早期我们经常对语言和库进行破坏性的修改时，Gofix是一个非常有价值的工具。例如，在Go 1之前，内置的错误接口并不存在，惯例是使用os.Error类型。当我们引入error时，我们提供了一个gofix模块，将所有对os.Error的引用及其相关的辅助函数重写为使用error和新的error包。这本来是一个令人生畏的尝试，但由于代码采用了标准格式，准备、执行和审查这一涉及几乎所有Go代码的变化就变得相对容易。

For more about gofix, see [this article](https://blog.golang.org/introducing-gofix).

关于gofix的更多信息，请看这篇文章。
