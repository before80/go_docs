+++
title = "godoc：文档化 go 代码"
weight = 24
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Godoc: documenting Go code - godoc：文档化 go 代码

https://go.dev/blog/godoc

Andrew Gerrand
31 March 2011

[***Note, June 2022**: For updated guidelines about documenting Go code, see "[Go Doc Comments](https://go.dev/doc/comment)."*]

[注，2022年6月：关于记录Go代码的最新指南，请参见 "Go文档注释"。］

The Go project takes documentation seriously. Documentation is a huge part of making software accessible and maintainable. Of course it must be well-written and accurate, but it also must be easy to write and to maintain. Ideally, it should be coupled to the code itself so the documentation evolves along with the code. The easier it is for programmers to produce good documentation, the better for everyone.

Go项目认真对待文档。文档是使软件可访问和可维护的一个重要部分。当然，它必须写得很好，很准确，但它也必须易于编写和维护。理想情况下，它应该与代码本身相耦合，这样文档就会随着代码的发展而发展。对程序员来说，越是容易产生好的文档，对大家来说就越好。

To that end, we have developed the [godoc](https://go.dev/cmd/godoc/) documentation tool. This article describes godoc’s approach to documentation, and explains how you can use our conventions and tools to write good documentation for your own projects.

为此，我们开发了godoc文档工具。本文描述了godoc的文档方法，并解释了您如何使用我们的惯例和工具为您自己的项目编写好的文档。

Godoc parses Go source code - including comments - and produces documentation as HTML or plain text. The end result is documentation tightly coupled with the code it documents. For example, through godoc’s web interface you can navigate from a function’s [documentation](https://go.dev/pkg/strings/#HasPrefix) to its [implementation](https://go.dev/src/strings/strings.go?s=11163:11200#L434) with one click.

Godoc解析Go源代码--包括注释--并以HTML或纯文本形式生成文档。最终的结果是文档与它所记录的代码紧密结合在一起。例如，通过godoc的网络界面，您可以一键从一个函数的文档导航到它的实现。

Godoc is conceptually related to Python’s [Docstring](https://www.python.org/dev/peps/pep-0257/) and Java’s [Javadoc](https://www.oracle.com/java/technologies/javase/javadoc-tool.html) but its design is simpler. The comments read by godoc are not language constructs (as with Docstring) nor must they have their own machine-readable syntax (as with Javadoc). Godoc comments are just good comments, the sort you would want to read even if godoc didn’t exist.

Godoc在概念上与Python的Docstring和Java的Javadoc相关，但其设计更简单。Godoc读取的注释不是语言结构（如Docstring），也不必须有自己的机器可读语法（如Javadoc）。Godoc注释就是好的注释，即使godoc不存在，您也会想读的那种。

The convention is simple: to document a type, variable, constant, function, or even a package, write a regular comment directly preceding its declaration, with no intervening blank line. Godoc will then present that comment as text alongside the item it documents. For example, this is the documentation for the `fmt` package’s [`Fprint`](https://go.dev/pkg/fmt/#Fprint) function:

惯例很简单：要记录一个类型、变量、常量、函数甚至包，在其声明前直接写一个普通的注释，中间不要有空行。然后，Godoc会将该注释作为文本呈现在它所记录的项目旁边。例如，这是fmt包的Fprint函数的文档：

```go linenums="1"
// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
```

Notice this comment is a complete sentence that begins with the name of the element it describes. This important convention allows us to generate documentation in a variety of formats, from plain text to HTML to UNIX man pages, and makes it read better when tools truncate it for brevity, such as when they extract the first line or sentence.

注意这个注释是一个完整的句子，以它所描述的元素的名字开始。这个重要的惯例使我们能够生成各种格式的文档，从纯文本到HTML再到UNIX手册页，并使其在工具为简洁而截断时，如提取第一行或第一句时，阅读效果更好。

Comments on package declarations should provide general package documentation. These comments can be short, like the [`sort`](https://go.dev/pkg/sort/) package’s brief description:

包声明的注释应该提供一般的包文档。这些注释可以很短，就像排序包的简短描述：

```go linenums="1"
// Package sort provides primitives for sorting slices and user-defined
// collections.
package sort
```

They can also be detailed like the [gob package](https://go.dev/pkg/encoding/gob/)’s overview. That package uses another convention for packages that need large amounts of introductory documentation: the package comment is placed in its own file, [doc.go](https://go.dev/src/pkg/encoding/gob/doc.go), which contains only those comments and a package clause.

它们也可以是详细的，比如gob包的概述。那个包对需要大量介绍性文档的包使用了另一个惯例：包的注释被放在它自己的文件doc.go中，其中只包含那些注释和一个包的条款。

When writing package comments of any size, keep in mind that their first sentence will appear in godoc’s [package list](https://go.dev/pkg/).

当写任何大小的包注释时，要记住它们的第一句话会出现在godoc的包列表中。

Comments that are not adjacent to a top-level declaration are omitted from godoc’s output, with one notable exception. Top-level comments that begin with the word `"BUG(who)"` are recognized as known bugs, and included in the "Bugs" section of the package documentation. The "who" part should be the user name of someone who could provide more information. For example, this is a known issue from the [bytes package](https://go.dev/pkg/bytes/#pkg-note-BUG):

与顶级声明不相邻的注释会从godoc的输出中省略，但有一个明显的例外。以 "BUG(who) "开头的顶层注释被认为是已知的错误，并包括在软件包文档的 "错误 "部分。"who "部分应该是可以提供更多信息的人的用户名。例如，这是byte包中的一个已知问题：

```
// BUG(r): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
```

Sometimes a struct field, function, type, or even a whole package becomes redundant or unnecessary, but must be kept for compatibility with existing programs. To signal that an identifier should not be used, add a paragraph to its doc comment that begins with "Deprecated:" followed by some information about the deprecation.

有时一个结构字段、函数、类型、甚至整个包变得多余或不必要，但为了与现有程序兼容，必须保留。为了表明一个标识符不应该被使用，可以在它的文档注释中添加一段话，以 "废弃的："开头，后面是一些关于废弃的信息。

There are a few formatting rules that Godoc uses when converting comments to HTML:

Godoc在将注释转换为HTML时使用了一些格式化规则：

- Subsequent lines of text are considered part of the same paragraph; you must leave a blank line to separate paragraphs.后续的几行文字被认为是同一段落的一部分；您必须留一个空行来分隔段落。
- Pre-formatted text must be indented relative to the surrounding comment text (see gob’s [doc.go](https://go.dev/src/pkg/encoding/gob/doc.go) for an example).预先格式化的文本必须相对于周围的评论文本缩进（见gob的doc.go的例子）。
- URLs will be converted to HTML links; no special markup is necessary.URL将被转换为HTML链接；不需要特殊标记。

Note that none of these rules requires you to do anything out of the ordinary.

请注意，这些规则都不要求您做任何出格的事情。

In fact, the best thing about godoc’s minimal approach is how easy it is to use. As a result, a lot of Go code, including all of the standard library, already follows the conventions.

事实上，godoc的最小化方法最好的一点是它很容易使用。因此，很多Go的代码，包括所有的标准库，都已经遵循了这些约定。

Your own code can present good documentation just by having comments as described above. Any Go packages installed inside `$GOROOT/src/pkg` and any `GOPATH` work spaces will already be accessible via godoc’s command-line and HTTP interfaces, and you can specify additional paths for indexing via the `-path` flag or just by running `"godoc ."` in the source directory. See the [godoc documentation](https://go.dev/cmd/godoc/) for more details.

您自己的代码只要有上述的注释，就可以呈现出良好的文档。任何安装在$GOROOT/src/pkg内的Go包和任何GOPATH工作空间都已经可以通过godoc的命令行和HTTP接口访问，您可以通过-path标志指定额外的路径进行索引，或者只需在源代码目录下运行 "godoc ."。更多细节见godoc文档。
