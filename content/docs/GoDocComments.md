+++
title = "go 文档注释"
weight = 24
date = 2023-05-18T17:26:23+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Go Doc Comments - Go 文档注释

> 原文：[https://go.dev/doc/comment](https://go.dev/doc/comment)

​	"Doc comments"是指出现在顶级package、const、func、type和var声明前的注释，中间没有换行。每个导出的（大写的）名称都应该有一个文档注释。

​	[go/doc](https://go.dev/pkg/go/doc) 和 [go/doc/comment](https://go.dev/pkg/go/doc/comment) 包提供了从 Go 源代码中提取文档的能力，许多工具都利用了这一功能。[go doc](../References/CommandDocumentation/go#show-documentation-for-package-or-symbol) 命令查找并打印出指定包或符号的文档注释。(符号是顶级的 const、func、type 或 var。）Web服务器 [pkg.go.dev](https://pkg.go.dev/) 显示了公共 Go 包的文档（当其许可证允许使用时）。为该网站服务的程序是 [golang.org/x/pkgsite/cmd/pkgsite](https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite)，它也可以在本地运行以查看私有模块的文档，或者在没有互联网连接的情况下。语言服务器[gopls](https://pkg.go.dev/golang.org/x/tools/gopls)在IDE中编辑Go源文件时提供文档。

​	本页的其余部分记录了如何编写 Go 文档注释。

## Packages 包

​	每个包都应该有一个介绍该包的包注释。它提供了与包整体相关的信息，并且通常设定了对包的期望。特别是在大型的包中，包的注释对API最重要的部分做一个简单的概述，并根据需要链接到其他的文档注释，这样做会很有帮助。

​	如果包很简单，包的注释可以很简短。例如：

```
// Package path implements utility routines for manipulating slash-separated
// paths.
//
// The path package should only be used for paths separated by forward
// slashes, such as the paths in URLs. This package does not deal with
// Windows paths with drive letters or backslashes; to manipulate
// operating system paths, use the [path/filepath] package.
package path
```

​	`[path/filepath]` 中的方括号创建了一个[文档链接](#links)。

​	从这个例子中可以看出，Go doc注释使用完整的句子。对于包的注释，这意味着第一句话以 "`Package`"开始。

​	对于多文件的包，包注释应该只出现在一个源文件中。如果多个文件都有包注释，它们会被串联起来，形成整个包的一个大注释。

## Commands 命令

​	命令的包注释与此类似，但它描述的是程序的行为而不是包中的Go符号。第一句话通常以程序本身的名称开始，大写的原因是它位于句子的开头。例如，这里是[gofmt](../References/CommandDocumentation/gofmt)的包注释的简化版本：

```
/*
Gofmt formats Go programs.
It uses tabs for indentation and blanks for alignment.
Alignment assumes that an editor is using a fixed-width font.

Without an explicit path, it processes the standard input. Given a file,
it operates on that file; given a directory, it operates on all .go files in
that directory, recursively. (Files starting with a period are ignored.)
By default, gofmt prints the reformatted sources to standard output.

Usage:

    gofmt [flags] [path ...]

The flags are:

    -d
        Do not print reformatted sources to standard output.
        If a file's formatting is different than gofmt's, print diffs
        to standard output.
    -w
        Do not print reformatted sources to standard output.
        If a file's formatting is different from gofmt's, overwrite it
        with gofmt's version. If an error occurred during overwriting,
        the original file is restored from an automatic backup.

When gofmt reads from standard input, it accepts either a full Go program
or a program fragment. A program fragment must be a syntactically
valid declaration list, statement list, or expression. When formatting
such a fragment, gofmt preserves leading indentation as well as leading
and trailing spaces, so that individual sections of a Go program can be
formatted by piping them through gofmt.
*/
package main
```

​	注释的开头是用[semantic linefeeds](https://rhodesmill.org/brandon/2012/one-sentence-per-line/)写的，每一个新的句子或长的短语都是自己在一行，这可以使差异在代码和注释的发展中更容易阅读。后面的段落恰好不遵循这个惯例，是用手包起来的。只要是最适合你的代码库的就可以了。无论哪种方式，在打印文档注释文本时，`go doc`和`pkgsite`都会重新包装它。例如：

```
$ go doc gofmt
Gofmt formats Go programs. It uses tabs for indentation and blanks for
alignment. Alignment assumes that an editor is using a fixed-width font.

Without an explicit path, it processes the standard input. Given a file, it
operates on that file; given a directory, it operates on all .go files in that
directory, recursively. (Files starting with a period are ignored.) By default,
gofmt prints the reformatted sources to standard output.

Usage:

    gofmt [flags] [path ...]

The flags are:

    -d
        Do not print reformatted sources to standard output.
        If a file's formatting is different than gofmt's, print diffs
        to standard output.
...
```

​	缩进的行被视为预格式化的文本：它们不被重新包装，并在HTML和Markdown演示中以代码字体打印。(下面的[语法](#syntax)部分给出了细节。)

## Types 类型

​	类型的文档注释应该解释该类型的每个实例代表或提供什么。如果API很简单，文档注释可以非常简短。例如：

```
package zip

// A Reader serves content from a ZIP archive.
type Reader struct {
    ...
}
```

​	默认情况下，程序员应该期望一个类型在同一时间只被一个goroutine使用是安全的。如果类型提供了更强的保证，文档注释中应该说明。例如：

```
package regexp

// Regexp is the representation of a compiled regular expression.
// A Regexp is safe for concurrent use by multiple goroutines,
// except for configuration methods, such as Longest.
type Regexp struct {
    ...
}
```

​	Go类型也应该致力于使零值有一个有用的意义。如果不是很明显，就应该把这个意义记录下来。例如：

```
package bytes

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
    ...
}
```

​	对于有导出字段的结构体，文档注释或每个字段的注释应该解释每个导出字段的含义。例如，这个类型的 doc 注释解释了这些字段：

```
package io

// A LimitedReader reads from R but limits the amount of
// data returned to just N bytes. Each call to Read
// updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0.
type LimitedReader struct {
    R   Reader // underlying reader
    N   int64  // max bytes remaining
}
```

​	相比之下，这种类型的文档注释将解释留给了每个字段的注释：

```
package comment

// A Printer is a doc comment printer.
// The fields in the struct can be filled in before calling
// any of the printing methods
// in order to customize the details of the printing process.
type Printer struct {
    // HeadingLevel is the nesting level used for
    // HTML and Markdown headings.
    // If HeadingLevel is zero, it defaults to level 3,
    // meaning to use <h3> and ###.
    HeadingLevel int
    ...
}
```

​	与包（上文）和函数（下文）一样，类型的文档注释以命名声明符号的完整句子开始。一个明确的主题通常会使措辞更清晰，而且它使文本更容易被搜索，无论是在网页上还是在命令行上。例如：

```
$ go doc -all regexp | grep pairs
pairs within the input string: result[2*n:2*n+2] identifies the indexes
    FindReaderSubmatchIndex returns a slice holding the index pairs identifying
    FindStringSubmatchIndex returns a slice holding the index pairs identifying
    FindSubmatchIndex returns a slice holding the index pairs identifying the
$
```

## Funcs 函数

A func’s doc comment should explain what the function returns or, for functions called for side effects, what it does. Named arguments or results can be referred to directly in the comment, without any special syntax like backquotes. (A consequence of this convention is that names like `a`, which might be mistaken for ordinary words, are typically avoided.) For example:

​	func 的文档注释应该解释该函数返回什么，或者，对于为副作用而调用的函数，它做了什么。命名的参数或结果可以直接在注释中提及，而不需要任何特殊的语法，如反引号。(这个惯例的后果是，像a这样的名字，可能会被误认为是普通的单词，通常要避免。) 例如：

```
package strconv

// Quote returns a double-quoted Go string literal representing s.
// The returned string uses Go escape sequences (\t, \n, \xFF, \u0100)
// for control characters and non-printable characters as defined by IsPrint.
func Quote(s string) string {
    ...
}
```

和：

```
package os

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.
//
// For portability, the status code should be in the range [0, 125].
func Exit(code int) {
    ...
}
```

If a doc comment needs to explain multiple results, naming the results can make the doc comment more understandable, even if the names are not used in the body of the function. For example:

如果一个文档注释需要解释多个结果，给结果命名可以使文档注释更容易理解，即使这些名字没有在函数的主体中使用。例如

```
package io

// Copy copies from src to dst until either EOF is reached
// on src or an error occurs. It returns the total number of bytes
// written and the first error encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF.
// Because Copy is defined to read from src until EOF, it does
// not treat an EOF from Read as an error to be reported.
func Copy(dst Writer, src Reader) (n int64, err error) {
    ...
}
```

Conversely, when the results don’t need to be named in the doc comment, they are usually omitted in the code as well, like in the `Quote` example above, to avoid cluttering the presentation.

相反，当结果不需要在文档注释中命名时，通常也会在代码中省略，就像上面的引用例子一样，以避免杂乱的表述。

These rules all apply both to plain functions and to methods. For methods, using the same receiver name avoids needless variation when listing all the methods of a type:

这些规则都适用于普通函数和方法。对于方法，使用相同的接收器名称可以避免在列出一个类型的所有方法时出现不必要的变化：

```
$ go doc bytes.Buffer
package bytes // import "bytes"

type Buffer struct {
    // Has unexported fields.
}
    A Buffer is a variable-sized buffer of bytes with Read and Write methods.
    The zero value for Buffer is an empty buffer ready to use.

func NewBuffer(buf []byte) *Buffer
func NewBufferString(s string) *Buffer
func (b *Buffer) Bytes() []byte
func (b *Buffer) Cap() int
func (b *Buffer) Grow(n int)
func (b *Buffer) Len() int
func (b *Buffer) Next(n int) []byte
func (b *Buffer) Read(p []byte) (n int, err error)
func (b *Buffer) ReadByte() (byte, error)
...
```

This example also shows that top-level functions returning a type `T` or pointer `*T`, perhaps with an additional error result, are shown alongside the type `T` and its methods, under the assumption that they are `T`’s constructors.

这个例子还表明，返回类型T或指针*T的顶级函数，也许还有一个额外的错误结果，在假设它们是T的构造函数的情况下，与类型T及其方法一起显示。

By default, programmers can assume that a top-level func is safe to call from multiple goroutines; this fact need not be stated explicitly.

默认情况下，程序员可以假设一个顶层func可以从多个goroutine中安全调用；这个事实不需要明确说明。

On the other hand, as noted in the previous section, using an instance of a type in any way, including calling a method, is typically assumed to be restricted to a single goroutine at a time. If the methods that are safe for concurrent use are not documented in the type’s doc comment, they should be documented in per-method comments. For example:

另一方面，正如上一节所指出的，以任何方式使用一个类型的实例，包括调用一个方法，通常都被假定为一次只能使用一个goroutine。如果对并发使用安全的方法没有记录在类型的文档注释中，它们应该被记录在每个方法的注释中。例如：

```
package sql

// Close returns the connection to the connection pool.
// All operations after a Close will return with ErrConnDone.
// Close is safe to call concurrently with other operations and will
// block until all other operations finish. It may be useful to first
// cancel any used context and then call Close directly after.
func (c *Conn) Close() error {
    ...
}
```

Note that func and method doc comments focus on what the operation returns or does, detailing what the caller needs to know. Special cases can be particularly important to document. For example:

请注意，func和方法的文档注释集中在操作返回或做什么，详细说明调用者需要知道什么。特殊情况下的文档可能特别重要。例如：

```
package math

// Sqrt returns the square root of x.
//
// Special cases are:
//
//  Sqrt(+Inf) = +Inf
//  Sqrt(±0) = ±0
//  Sqrt(x < 0) = NaN
//  Sqrt(NaN) = NaN
func Sqrt(x float64) float64 {
    ...
}
```

Doc comments should not explain internal details such as the algorithm used in the current implementation. Those are best left to comments inside the function body. It may be appropriate to give asymptotic time or space bounds when that detail is particularly important to callers. For example:

文件注释不应该解释内部细节，如当前实现中使用的算法。这些最好留给函数正文中的注释。当这个细节对调用者特别重要时，给出渐近的时间或空间界限可能是合适的。例如：

```
package sort

// Sort sorts data in ascending order as determined by the Less method.
// It makes one call to data.Len to determine n and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func Sort(data Interface) {
    ...
}
```

Because this doc comment makes no mention of which sorting algorithm is used, it is easier to change the implementation to use a different algorithm in the future.

因为这个文档注释没有提到使用哪种排序算法，所以将来改变实现以使用不同的算法会比较容易。

## Consts 常量

Go’s declaration syntax allows grouping of declarations, in which case a single doc comment can introduce a group of related constants, with individual constants only documented by short end-of-line comments. For example:

Go 的声明语法允许对声明进行分组，在这种情况下，一个文档注释可以介绍一组相关的常量，单个常量仅由简短的行末注释记录。例如：

```
package scanner // import "text/scanner"

// The result of Scan is one of these tokens or a Unicode character.
const (
    EOF = -(iota + 1)
    Ident
    Int
    Float
    Char
    ...
)
```

Sometimes the group needs no doc comment at all. For example:

有时该组根本不需要文档注释。例如：

```
package unicode // import "unicode"

const (
    MaxRune         = '\U0010FFFF' // maximum valid Unicode code point.
    ReplacementChar = '\uFFFD'     // represents invalid code points.
    MaxASCII        = '\u007F'     // maximum ASCII value.
    MaxLatin1       = '\u00FF'     // maximum Latin-1 value.
)
```

On the other hand, ungrouped constants typically warrant a full doc comment starting with a complete sentence. For example:

另一方面，未分组的常数通常需要一个完整的文档注释，以一个完整的句子开始。例如：

```
package unicode

// Version is the Unicode edition from which the tables are derived.
const Version = "13.0.0"
```

Typed constants are displayed next to the declaration of their type and as a result often omit a const group doc comment in favor of the type’s doc comment. For example:

类型化的常量被显示在其类型声明的旁边，因此常常省略常量组的文档注释，而选择类型的文档注释。例如：

```
package syntax

// An Op is a single regular expression operator.
type Op uint8

const (
    OpNoMatch        Op = 1 + iota // matches no strings
    OpEmptyMatch                   // matches empty string
    OpLiteral                      // matches Runes sequence
    OpCharClass                    // matches Runes interpreted as range pair list
    OpAnyCharNotNL                 // matches any character except newline
    ...
)
```

(See [pkg.go.dev/regexp/syntax#Op](https://pkg.go.dev/regexp/syntax#Op) for the HTML presentation.)

(参见 pkg.go.dev/regexp/syntax#Op 的 HTML 演示。)

## Vars 变量

The conventions for variables are the same as those for constants. For example, here is a set of grouped variables:

变量的约定与常量的约定相同。例如，这里有一组分组的变量：

```
package fs

// Generic file system errors.
// Errors returned by file systems can be tested against these errors
// using errors.Is.
var (
    ErrInvalid    = errInvalid()    // "invalid argument"
    ErrPermission = errPermission() // "permission denied"
    ErrExist      = errExist()      // "file already exists"
    ErrNotExist   = errNotExist()   // "file does not exist"
    ErrClosed     = errClosed()     // "file already closed"
)
```

And a single variable:

还有一个单一的变量：

```
package unicode

// Scripts is the set of Unicode script tables.
var Scripts = map[string]*RangeTable{
    "Adlam":                  Adlam,
    "Ahom":                   Ahom,
    "Anatolian_Hieroglyphs":  Anatolian_Hieroglyphs,
    "Arabic":                 Arabic,
    "Armenian":               Armenian,
    ...
}
```

## Syntax 语法

Go doc comments are written in a simple syntax that supports paragraphs, headings, links, lists, and preformatted code blocks. To keep comments lightweight and readable in source files, there is no support for complex features like font changes or raw HTML. Markdown aficionados can view the syntax as a simplified subset of Markdown.

Go doc 的注释是用一种简单的语法编写的，支持段落、标题、链接、列表和预格式化的代码块。为了保持注释在源文件中的轻量级和可读性，不支持复杂的功能，如字体变化或原始HTML。Markdown爱好者可以把这个语法看作是Markdown的一个简化子集。

The standard formatter [gofmt](https://go.dev/cmd/gofmt) reformats doc comments to use a canonical formatting for each of these features. Gofmt aims for readability and user control over how comments are written in source code but will adjust presentation to make the semantic meaning of a particular comment clearer, analogous to reformatting `1+2 * 3` to `1 + 2*3` in ordinary source code.

标准的格式化器gofmt对文档注释进行了重新格式化，为每一个特征使用了一个规范的格式化。Gofmt的目的是使源代码中的注释具有可读性和用户控制，但会调整表现形式，使特定注释的语义更清晰，类似于将普通源代码中的1+2*3重新格式化为1+2*3。

Directive comments such as `//go:generate` are not considered part of a doc comment and are omitted from rendered documentation. Gofmt moves directive comments to the end of the doc comment, preceded by a blank line. For example:

指令性注释，如//go:generate，不被视为文档注释的一部分，在渲染的文档中被省略。Gofmt 将指令性注释移到文档注释的末尾，前面加一个空行。例如：

```
package regexp

// An Op is a single regular expression operator.
//
//go:generate stringer -type Op -trimprefix Op
type Op uint8
```

A directive comment is a line matching the regular expression `//(line |extern |export |[a-z0-9]+:[a-z0-9])`. Tools that define their own directives should use the form `//toolname:directive`.

指令注释是与正则表达式相匹配的行，//(行|外部|出口|[a-z0-9]+:[a-z0-9])。定义自己的指令的工具应该使用//toolname:directive的形式。

Gofmt removes leading and trailing blank lines in doc comments.

Gofmt删除文档注释中的前导和尾部空行。

### Paragraphs 段落

A paragraph is a span of unindented non-blank lines. We’ve already seen many examples of paragraphs.

一个段落是由无缩进的非空行组成的跨度。我们已经看到了许多段落的例子。

A pair of consecutive backticks (` U+0060) is interpreted as a Unicode left quote (“ U+201C), and a pair of consecutive single quotes (' U+0027) is interpreted as a Unicode right quote (” U+201D).

一对连续的反斜线（` U+0060）被解释为Unicode左引号（" U+201C），而一对连续的单引号（' U+0027）被解释为Unicode右引号（" U+201D）。

Gofmt preserves line breaks in paragraph text: it does not rewrap the text. This allows the use of [semantic linefeeds](https://rhodesmill.org/brandon/2012/one-sentence-per-line/), as seen earlier. Gofmt replaces duplicated blank lines between paragraphs with a single blank line. Gofmt also reformats consecutive backticks or single quotes to their Unicode interpretations.

Gofmt保留了段落文本中的换行符：它不会重新包装文本。这允许使用语义上的换行，如前面所见。Gofmt将段落之间重复的空行替换为一个空行。Gofmt还将连续的反斜线或单引号重新转换为Unicode的解释。

### Headings 标题

A heading is a line beginning with a number sign (U+0023) and then a space and the heading text. To be recognized as a heading, the line must be unindented and set off from adjacent paragraph text by blank lines.

标题是以数字符号（U+0023）开始的一行，然后是一个空格和标题文本。要被识别为标题，该行必须是无缩进的，并通过空行与相邻的段落文本分开。

For example:

例如：

```
// Package strconv implements conversions to and from string representations
// of basic data types.
//
// # Numeric Conversions
//
// The most common numeric conversions are [Atoi] (string to int) and [Itoa] (int to string).
...
package strconv
```

On the other hand:

另一方面：

```
// #This is not a heading, because there is no space.
//
// # This is not a heading,
// # because it is multiple lines.
//
// # This is not a heading,
// because it is also multiple lines.
//
// The next paragraph is not a heading, because there is no additional text:
//
// #
//
// In the middle of a span of non-blank lines,
// # this is not a heading either.
//
//     # This is not a heading, because it is indented.
```

The # syntax was added in Go 1.19. Before Go 1.19, headings were identified implicitly by single-line paragraphs satisfying certain conditions, most notably the lack of any terminating punctuation.

这个#语法是在Go 1.19中加入的。在Go 1.19之前，标题是由满足某些条件的单行段落隐含地识别出来的，最明显的是没有任何结尾的标点符号。

Gofmt reformats [lines treated as implicit headings](https://github.com/golang/proposal/blob/master/design/51082-godocfmt.md#headings) by earlier versions of Go to use # headings instead. If the reformatting is not appropriate—that is, if the line was not meant to be a heading—the easiest way to make it a paragraph is to introduce terminating punctuation such as a period or colon, or to break it into two lines.

Gofmt将早期版本的Go中被视为隐式标题的行数改成使用#标题。如果重新格式化不合适，也就是说，如果该行不是作为标题，那么使其成为段落的最简单方法是引入句号或冒号等结束性标点，或者将其分成两行。

### Links 链接

A span of unindented non-blank lines defines link targets when every line is of the form “[Text]: URL”. In other text in the same doc comment, “[Text]” represents a link to URL using the given text—in HTML, <a href="URL">Text</a>. For example:

当每一行都是"[文本].URL "的形式时，一个无缩进的非空白行的跨度定义了链接目标。URL"。在同一文档注释的其他文本中，"[文本]"代表使用给定文本的URL的链接--HTML，<a href="URL">文本</a>。例如。

```
// Package json implements encoding and decoding of JSON as defined in
// [RFC 7159]. The mapping between JSON and Go values is described
// in the documentation for the Marshal and Unmarshal functions.
//
// For an introduction to this package, see the article
// “[JSON and Go].”
//
// [RFC 7159]: https://tools.ietf.org/html/rfc7159
// [JSON and Go]: https://golang.org/doc/articles/json_and_go.html
package json
```

By keeping URLs in a separate section, this format only minimally interrupts the flow of the actual text. It also roughly matches the Markdown [shortcut reference link format](https://spec.commonmark.org/0.30/#shortcut-reference-link), without the optional title text.

通过将URL保持在一个单独的部分，这种格式只对实际文本的流程产生最小的干扰。它也大致符合Markdown的快捷参考链接格式，没有可选的标题文本。

If there is no corresponding URL declaration, then (except for doc links, described in the next section) “[Text]” is not a hyperlink, and the square brackets are preserved when displayed. Each doc comment is considered independently: link target definitions in one comment do not affect other comments.

如果没有相应的URL声明，那么（除了文档链接，在下一节中描述）"[文本]"就不是一个超链接，并且在显示时保留了方括号。每个文档注释都是独立考虑的：一个注释中的链接目标定义并不影响其他注释。

Although link target definition blocks may be interleaved with ordinary paragraphs, gofmt moves all link target definitions to the end of the doc comment, in up to two blocks: first a block containing all the link targets that are referenced in the comment, and then a block containing all the targets *not* referenced in the comment. The separate block makes unused targets easy to notice and fix (in case the links or the definitions have typos) or to delete (in case the definitions are no longer needed).

虽然链接目标定义块可以与普通段落交错排列，但gofmt将所有链接目标定义移到文档注释的末尾，最多有两个块：首先是包含注释中引用的所有链接目标的块，然后是包含注释中未引用的所有目标的块。分开的块使未使用的目标易于注意和修复（以防链接或定义有错别字）或删除（以防定义不再需要）。

Plain text that is recognized as a URL is automatically linked in HTML renderings.

被识别为URL的纯文本会在HTML渲染中自动链接。

### Doc links 文件链接

Doc links are links of the form “[Name1]” or “[Name1.Name2]” to refer to exported identifiers in the current package, or “[pkg]”, “[pkg.Name1]”, or “[pkg.Name1.Name2]” to refer to identifiers in other packages.

文件链接是"[Name1]"或"[Name1.Name2]"形式的链接，指的是当前包中导出的标识符，或者"[pkg]"、"[pkg.Name1]"或"[pkg.Name1.Name2]"指的是其他包中的标识符。

For example:

例如：

```
package bytes

// ReadFrom reads data from r until EOF and appends it to the buffer, growing
// the buffer as needed. The return value n is the number of bytes read. Any
// error except [io.EOF] encountered during the read is also returned. If the
// buffer becomes too large, ReadFrom will panic with [ErrTooLarge].
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
    ...
}
```

The bracketed text for a symbol link can include an optional leading star, making it easy to refer to pointer types, such as [*bytes.Buffer].

符号链接的括号内的文字可以包括一个可选的前导星，使其易于引用指针类型，如[*bytes.Buffer]。

When referring to other packages, “pkg” can be either a full import path or the assumed package name of an existing import. The assumed package name is either the identifier in a renamed import or else [the name assumed by goimports](https://pkg.go.dev/golang.org/x/tools/internal/imports#ImportPathToAssumedName). (Goimports inserts renamings when that assumption is not correct, so this rule should work for essentially all Go code.) For example, if the current package imports encoding/json, then “[json.Decoder]” can be written in place of “[encoding/json.Decoder]” to link to the docs for encoding/json’s Decoder. If different source files in a package import different packages using the same name, then the shorthand is ambiguous and cannot be used.

当引用其他包时，"pkg "可以是完整的导入路径，也可以是现有导入的假定包名。假设的包名要么是重命名的导入中的标识符，要么是 goimports 假设的名称。(当这个假设不正确时，Goimports 会插入重命名，所以这个规则基本上对所有 Go 代码都有效。) 例如，如果当前包导入了 encoding/json，那么可以用 "[json.Decoder]" 来代替 "[encoding/json.Decoder]" 来链接到 encoding/json 的 Decoder 文档。如果一个包中的不同源文件导入了使用相同名称的不同包，那么这个速记就会产生歧义，不能使用。

A “pkg” is only assumed to be a full import path if it starts with a domain name (a path element with a dot) or is one of the packages from the standard library (“[os]”, “[encoding/json]”, and so on). For example, `[os.File]` and `[example.com/sys.File]` are documentation links (the latter will be a broken link), but `[os/sys.File]` is not, because there is no os/sys package in the standard library.

只有当 "pkg "以域名（带点的路径元素）开头或者是标准库中的一个包（"[os]"，"[encoding/json]"，等等），才会被认为是一个完整的导入路径。例如，[os.File]和[example.com/sys.File]是文档链接（后者将是一个断开的链接），但[os/sys.File]不是，因为标准库中没有os/sys包。

To avoid problems with maps, generics, and array types, doc links must be both preceded and followed by punctuation, spaces, tabs, or the start or end of a line. For example, the text “map[ast.Expr]TypeAndValue” does not contain a doc link.

为了避免地图、泛型和数组类型的问题，文档链接必须在前面和后面都有标点符号、空格、制表符、或一行的开始或结束。例如，文本 "map[ast.Expr]TypeAndValue "并不包含文档链接。

### Lists 列表

A list is a span of indented or blank lines (which would otherwise be a code block, as described in the next section) in which the first indented line begins with a bullet list marker or a numbered list marker.

列表是一个缩进或空白行的跨度（否则就是一个代码块，如下一节所述），其中第一个缩进的行以一个子弹头列表标记或一个编号列表标记开始。

A bullet list marker is a star, plus, dash, or Unicode bullet (*, +, -, •; U+002A, U+002B, U+002D, U+2022) followed by a space or tab and then text. In a bullet list, each line beginning with a bullet list marker starts a new list item.

项目表标记是一个星号、加号、破折号或Unicode子弹（*、+、-、-；U+002A、U+002B、U+002D、U+2022），后面是一个空格或制表符，然后是文本。在弹出式列表中，每一行以弹出式列表标记开始，就是一个新的列表项目。

For example:

例如：

```
package url

// PublicSuffixList provides the public suffix of a domain. For example:
//   - the public suffix of "example.com" is "com",
//   - the public suffix of "foo1.foo2.foo3.co.uk" is "co.uk", and
//   - the public suffix of "bar.pvt.k12.ma.us" is "pvt.k12.ma.us".
//
// Implementations of PublicSuffixList must be safe for concurrent use by
// multiple goroutines.
//
// An implementation that always returns "" is valid and may be useful for
// testing but it is not secure: it means that the HTTP server for foo.com can
// set a cookie for bar.com.
//
// A public suffix list implementation is in the package
// golang.org/x/net/publicsuffix.
type PublicSuffixList interface {
    ...
}
```

A numbered list marker is a decimal number of any length followed by a period or right parenthesis, then a space or tab, and then text. In a numbered list, each line beginning with a number list marker starts a new list item. Item numbers are left as is, never renumbered.

一个编号列表标记是一个任意长度的十进制数字，后面是句号或右括号，然后是空格或制表符，最后是文本。在一个编号列表中，每一行以数字列表标记开始，就开始一个新的列表项。项目编号保持原样，从不重新编号。

For example:

例如：

```
package path

// Clean returns the shortest path name equivalent to path
// by purely lexical processing. It applies the following rules
// iteratively until no further processing can be done:
//
//  1. Replace multiple slashes with a single slash.
//  2. Eliminate each . path name element (the current directory).
//  3. Eliminate each inner .. path name element (the parent directory)
//     along with the non-.. element that precedes it.
//  4. Eliminate .. elements that begin a rooted path:
//     that is, replace "/.." by "/" at the beginning of a path.
//
// The returned path ends in a slash only if it is the root "/".
//
// If the result of this process is an empty string, Clean
// returns the string ".".
//
// See also Rob Pike, “[Lexical File Names in Plan 9].”
//
// [Lexical File Names in Plan 9]: https://9p.io/sys/doc/lexnames.html
func Clean(path string) string {
    ...
}
```

List items only contain paragraphs, not code blocks or nested lists. This avoids any space-counting subtlety as well as questions about how many spaces a tab counts for in inconsistent indentation.

列表项只包含段落，不包含代码块或嵌套列表。这避免了任何空间计算的微妙之处，也避免了关于在不一致的缩进中一个制表符算多少个空格的问题。

Gofmt reformats bullet lists to use a dash as the bullet marker, two spaces of indentation before the dash, and four spaces of indentation for continuation lines.

Gofmt重新格式化子弹头列表，使用破折号作为子弹头标记，在破折号之前缩进两个空格，并为续行缩进四个空格。

Gofmt reformats numbered lists to use a single space before the number, a period after the number, and again four spaces of indentation for continuation lines.

Gofmt重新格式化数字列表，在数字前使用一个空格，在数字后使用一个句号，并再次为续行缩进四个空格。

Gofmt preserves but does not require a blank line between a list and the preceding paragraph. It inserts a blank line between a list and the following paragraph or heading.

Gofmt 保留但不要求在列表和前段之间有空行。它在列表和下面的段落或标题之间插入一个空行。

### Code blocks 代码块

A code block is a span of indented or blank lines not starting with a bullet list marker or numbered list marker. It is rendered as preformatted text (a <pre> block in HTML).

代码块是一个缩进或空行的跨度，不以项目列表标记或编号列表标记开始。它被呈现为预格式化的文本（HTML中的<pre>块）。

Code blocks often contain Go code. For example:

代码块通常包含Go代码。例如：

```
package sort

// Search uses binary search...
//
// As a more whimsical example, this program guesses your number:
//
//  func GuessingGame() {
//      var s string
//      fmt.Printf("Pick an integer from 0 to 100.\n")
//      answer := sort.Search(100, func(i int) bool {
//          fmt.Printf("Is your number <= %d? ", i)
//          fmt.Scanf("%s", &s)
//          return s != "" && s[0] == 'y'
//      })
//      fmt.Printf("Your number is %d.\n", answer)
//  }
func Search(n int, f func(int) bool) int {
    ...
}
```

Of course, code blocks also often contain preformatted text besides code. For example:

当然，除了代码之外，代码块还经常包含预格式化的文本。例如：

```
package path

// Match reports whether name matches the shell pattern.
// The pattern syntax is:
//
//  pattern:
//      { term }
//  term:
//      '*'         matches any sequence of non-/ characters
//      '?'         matches any single non-/ character
//      '[' [ '^' ] { character-range } ']'
//                  character class (must be non-empty)
//      c           matches character c (c != '*', '?', '\\', '[')
//      '\\' c      matches character c
//
//  character-range:
//      c           matches character c (c != '\\', '-', ']')
//      '\\' c      matches character c
//      lo '-' hi   matches character c for lo <= c <= hi
//
// Match requires pattern to match all of name, not just a substring.
// The only possible returned error is [ErrBadPattern], when pattern
// is malformed.
func Match(pattern, name string) (matched bool, err error) {
    ...
}
```

Gofmt indents all lines in a code block by a single tab, replacing any other indentation the non-blank lines have in common. Gofmt also inserts a blank line before and after each code block, distinguishing the code block clearly from the surrounding paragraph text.

Gofmt将一个代码块中的所有行缩进一个制表符，替换非空白行共同拥有的任何其他缩进。Gofmt还在每个代码块前后插入一个空行，将代码块与周围的段落文本明确区分开来。

## Common mistakes and pitfalls 常见的错误和误区

The rule that any span of indented or blank lines in a doc comment is rendered as a code block dates to the earliest days of Go. Unfortunately, the lack of support for doc comments in gofmt has led to many existing comments that use indentation without meaning to create a code block.

在文档注释中，任何缩进或空白行的跨度都会被呈现为一个代码块，这一规则可以追溯到Go的最早时代。不幸的是，由于gofmt中缺乏对文档注释的支持，导致许多现有的注释在使用缩进的时候没有意义地创建一个代码块。

For example, this unindented list has always been interpreted by godoc as a three-line paragraph followed by a one-line code block:

例如，这个没有缩进的列表一直被godoc解释为一个三行的段落，后面是一个单行的代码块：

```
package http

// cancelTimerBody is an io.ReadCloser that wraps rc with two features:
// 1) On Read error or close, the stop func is called.
// 2) On Read failure, if reqDidTimeout is true, the error is wrapped and
//    marked as net.Error that hit its timeout.
type cancelTimerBody struct {
    ...
}
```

This always rendered in `go` `doc` as:

这在 go doc 中总是被渲染成：

```
cancelTimerBody is an io.ReadCloser that wraps rc with two features:
1) On Read error or close, the stop func is called. 2) On Read failure,
if reqDidTimeout is true, the error is wrapped and

    marked as net.Error that hit its timeout.
```

Similarly, the command in this comment is a one-line paragraph followed by a one-line code block:

类似地，本注释中的命令是一个单行段落，后面是一个单行代码块：

```
package smtp

// localhostCert is a PEM-encoded TLS cert generated from src/crypto/tls:
//
// go run generate_cert.go --rsa-bits 1024 --host 127.0.0.1,::1,example.com \
//     --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var localhostCert = []byte(`...`)
```

This rendered in `go` `doc` as:

这在go文档中呈现为：

```
localhostCert is a PEM-encoded TLS cert generated from src/crypto/tls:

go run generate_cert.go --rsa-bits 1024 --host 127.0.0.1,::1,example.com \

    --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
```

And this comment is a two-line paragraph (the second line is “{”), followed by a six-line indented code block and a one-line paragraph (“}”).

而这个注释是一个两行段落（第二行是"{"），后面是一个六行缩进的代码块和一个单行段落（"}"）。

```
// On the wire, the JSON will look something like this:
// {
//  "kind":"MyAPIObject",
//  "apiVersion":"v1",
//  "myPlugin": {
//      "kind":"PluginA",
//      "aOption":"foo",
//  },
// }
```

And this rendered in `go` `doc` as:

而这在go doc中呈现为：

```
On the wire, the JSON will look something like this: {

    "kind":"MyAPIObject",
    "apiVersion":"v1",
    "myPlugin": {
        "kind":"PluginA",
        "aOption":"foo",
    },

}
```

Another common mistake was an unindented Go function definition or block statement, similarly bracketed by “{” and “}”.

另一个常见的错误是没有缩进的Go函数定义或块状语句，同样是用"{"和"}"括起来的。

The introduction of doc comment reformatting in Go 1.19’s gofmt makes mistakes like these more visible by adding blank lines around the code blocks.

Go 1.19的gofmt中引入了doc注释重排，通过在代码块周围添加空行，使这类错误更加明显。

Analysis in 2022 found that only 3% of doc comments in public Go modules were reformatted at all by the draft Go 1.19 gofmt. Limiting ourselves to those comments, about 87% of gofmt’s reformattings preserved the structure that a person would infer from reading the comment; about 6% were tripped up by these kinds of unindented lists, unindented multiline shell commands, and unindented brace-delimited code blocks.

2022年的分析发现，只有3%的公共Go模块中的文档注释被Go 1.19 gofmt草案重新格式化。仅限于这些注释，大约87%的gofmt的重新格式化保留了一个人从阅读注释中推断出的结构；大约6%的注释被这些无缩进的列表、无缩进的多行shell命令和无缩进的括号分隔的代码块绊倒。

Based on this analysis, the Go 1.19 gofmt applies a few heuristics to merge unindented lines into an adjacent indented list or code block. With those adjustments, the Go 1.19 gofmt reformats the above examples to:

基于这一分析，Go 1.19的gofmt应用了一些启发式方法，将未缩进的行合并到相邻的缩进列表或代码块中。经过这些调整，Go 1.19版gofmt将上述例子改写为。

```
// cancelTimerBody is an io.ReadCloser that wraps rc with two features:
//  1. On Read error or close, the stop func is called.
//  2. On Read failure, if reqDidTimeout is true, the error is wrapped and
//     marked as net.Error that hit its timeout.

// localhostCert is a PEM-encoded TLS cert generated from src/crypto/tls:
//
//  go run generate_cert.go --rsa-bits 1024 --host 127.0.0.1,::1,example.com \
//      --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h

// On the wire, the JSON will look something like this:
//
//  {
//      "kind":"MyAPIObject",
//      "apiVersion":"v1",
//      "myPlugin": {
//          "kind":"PluginA",
//          "aOption":"foo",
//      },
//  }
```

This reformatting makes the meaning clearer as well as making the doc comments render correctly in earlier versions of Go. If the heuristic ever makes a bad decision, it can be overridden by inserting a blank line to clearly separate the paragraph text from non-paragraph text.

这种重新格式化的做法使含义更加清晰，同时也使文档注释在早期版本的Go中正确呈现。如果启发式做出了错误的决定，可以通过插入空行将段落文本和非段落文本明确分开来推翻它。

Even with these heuristics, other existing comments will need manual adjustment to correct their rendering. The most common mistake is indenting a wrapped unindented line of text. For example:

即使有了这些启发式方法，其他现有的注释也需要手动调整，以纠正它们的呈现方式。最常见的错误是缩进被包裹的未缩进的文本行。例如：

```
// TODO Revisit this design. It may make sense to walk those nodes
//      only once.

// According to the document:
// "The alignment factor (in bytes) that is used to align the raw data of sections in
//  the image file. The value should be a power of 2 between 512 and 64 K, inclusive."
```

In both of these, the last line is indented, making it a code block. The fix is to unindent the lines.

在这两个文件中，最后一行是缩进的，使其成为一个代码块。解决的方法是取消缩进行。

Another common mistake is not indenting a wrapped indented line of a list or code block. For example:

另一个常见的错误是不缩进列表或代码块中被包裹的缩进行。例如：

```
// Uses of this error model include:
//
//   - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
//   - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
```

The fix is to indent the wrapped lines.

修复方法是缩进被包裹的行。

Go doc comments do not support nested lists, so gofmt reformats

Go文档中的注释不支持嵌套的列表，所以gofmt改写了

```
// Here is a list:
//
//  - Item 1.
//    * Subitem 1.
//    * Subitem 2.
//  - Item 2.
//  - Item 3.
```

to 到

```
// Here is a list:
//
//  - Item 1.
//  - Subitem 1.
//  - Subitem 2.
//  - Item 2.
//  - Item 3.
```

Rewriting the text to avoid nested lists usually improves the documentation and is the best solution. Another potential workaround is to mix list markers, since bullet markers do not introduce list items in a numbered list, nor vice versa. For example:

重写文本以避免嵌套的列表，通常会改善文档，是最好的解决方案。另一个潜在的解决方法是混合列表标记，因为子弹头标记不能在数字列表中引入列表项，反之亦然。例如。

```
// Here is a list:
//
//  1. Item 1.
//
//     - Subitem 1.
//
//     - Subitem 2.
//
//  2. Item 2.
//
//  3. Item 3.
```