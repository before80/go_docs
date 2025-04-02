+++
title = "doc/comment"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/doc/comment@go1.24.2](https://pkg.go.dev/go/doc/comment@go1.24.2)

Package comment implements parsing and reformatting of Go doc comments, (documentation comments), which are comments that immediately precede a top-level declaration of a package, const, func, type, or var.

​	comment 包实现 Go doc 注释（文档注释）的解析和重新格式化，这些注释紧接在包、const、func、type 或 var 的顶级声明之前。

Go doc comment syntax is a simplified subset of Markdown that supports links, headings, paragraphs, lists (without nesting), and preformatted text blocks. The details of the syntax are documented at https://go.dev/doc/comment.

​	Go doc 注释语法是 Markdown 的一个简化子集，支持链接、标题、段落、列表（无嵌套）和预格式文本块。语法的详细信息记录在 https://go.dev/doc/comment 中。

To parse the text associated with a doc comment (after removing comment markers), use a [Parser](https://pkg.go.dev/go/doc/comment@go1.20.1#Parser):

​	要解析与 doc 注释关联的文本（在删除注释标记后），请使用 Parser：

```go
var p comment.Parser
doc := p.Parse(text)
```

The result is a [*Doc](https://pkg.go.dev/go/doc/comment@go1.20.1#Doc). To reformat it as a doc comment, HTML, Markdown, or plain text, use a [Printer](https://pkg.go.dev/go/doc/comment@go1.20.1#Printer):

​	结果是 *Doc。要将其重新格式化为 doc 注释、HTML、Markdown 或纯文本，请使用 Printer：

```go
var pr comment.Printer
os.Stdout.Write(pr.Text(doc))
```

The [Parser](https://pkg.go.dev/go/doc/comment@go1.20.1#Parser) and [Printer](https://pkg.go.dev/go/doc/comment@go1.20.1#Printer) types are structs whose fields can be modified to customize the operations. For details, see the documentation for those types.

​	Parser 和 Printer 类型是 struct，其字段可以修改以自定义操作。有关详细信息，请参阅这些类型的文档。

Use cases that need additional control over reformatting can implement their own logic by inspecting the parsed syntax itself. See the documentation for [Doc](https://pkg.go.dev/go/doc/comment@go1.20.1#Doc), [Block](https://pkg.go.dev/go/doc/comment@go1.20.1#Block), [Text](https://pkg.go.dev/go/doc/comment@go1.20.1#Text) for an overview and links to additional types.

​	需要对重新格式化进行更多控制的用例可以通过检查解析的语法本身来实现自己的逻辑。请参阅 Doc、Block、Text 的文档以获取概述和指向其他类型的链接。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func DefaultLookupPackage

```go
func DefaultLookupPackage(name string) (importPath string, ok bool)
```

DefaultLookupPackage is the default package lookup function, used when [Parser](https://pkg.go.dev/go/doc/comment@go1.20.1#Parser).LookupPackage is nil. It recognizes names of the packages from the standard library with single-element import paths, such as math, which would otherwise be impossible to name.

​	DefaultLookupPackage 是默认包查找函数，在 Parser.LookupPackage 为 nil 时使用。它识别具有单元素导入路径的标准库的包的名称，例如 math，否则无法命名。

Note that the go/doc package provides a more sophisticated lookup based on the imports used in the current package.

​	请注意，go/doc 包根据当前包中使用的导入提供更复杂的查找。

## 类型

### type Block

```go
type Block interface {
	// contains filtered or unexported methods
}
```

A Block is block-level content in a doc comment, one of [*Code](https://pkg.go.dev/go/doc/comment@go1.20.1#Code), [*Heading](https://pkg.go.dev/go/doc/comment@go1.20.1#Heading), [*List](https://pkg.go.dev/go/doc/comment@go1.20.1#List), or [*Paragraph](https://pkg.go.dev/go/doc/comment@go1.20.1#Paragraph).

​	Block 是文档注释中的块级内容，包括 *Code、*Heading、*List 或 *Paragraph。

### type Code

```go
type Code struct {
	// Text is the preformatted text, ending with a newline character.
	// It may be multiple lines, each of which ends with a newline character.
	// It is never empty, nor does it start or end with a blank line.
	Text string
}
```

A Code is a preformatted code block.

​	Code 是预格式化的代码块。

### type Doc

```go
type Doc struct {
	// Content is the sequence of content blocks in the comment.
	Content []Block

	// Links is the link definitions in the comment.
	Links []*LinkDef
}
```

A Doc is a parsed Go doc comment.

​	Doc 是已解析的 Go 文档注释。

### type DocLink

```go
type DocLink struct {
	Text []Text // text of link

	// ImportPath, Recv, and Name identify the Go package or symbol
	// that is the link target. The potential combinations of
	// non-empty fields are:
	//  - ImportPath: a link to another package
	//  - ImportPath, Name: a link to a const, func, type, or var in another package
	//  - ImportPath, Recv, Name: a link to a method in another package
	//  - Name: a link to a const, func, type, or var in this package
	//  - Recv, Name: a link to a method in this package
	ImportPath string // import path
	Recv       string // receiver type, without any pointer star, for methods
	Name       string // const, func, type, var, or method name
}
```

A DocLink is a link to documentation for a Go package or symbol.

​	DocLink 是指向 Go 包或符号的文档的链接。

#### (*DocLink) DefaultURL 

```go
func (l *DocLink) DefaultURL(baseURL string) string
```

DefaultURL constructs and returns the documentation URL for l, using baseURL as a prefix for links to other packages.

​	DefaultURL 使用 baseURL 作为指向其他包的链接的前缀，构建并返回 l 的文档 URL。

The possible forms returned by DefaultURL are:

​	DefaultURL 返回的可能形式为：

- baseURL/ImportPath, for a link to another package
  baseURL/ImportPath，指向另一个包的链接
- baseURL/ImportPath#Name, for a link to a const, func, type, or var in another package
  baseURL/ImportPath#Name，指向另一个包中的常量、函数、类型或变量的链接
- baseURL/ImportPath#Recv.Name, for a link to a method in another package
  baseURL/ImportPath#Recv.Name，指向另一个包中的方法的链接
- \#Name, for a link to a const, func, type, or var in this package
  \#Name，指向此包中的常量、函数、类型或变量的链接
- \#Recv.Name, for a link to a method in this package
  \#Recv.Name，指向此包中的方法的链接

If baseURL ends in a trailing slash, then DefaultURL inserts a slash between ImportPath and # in the anchored forms. For example, here are some baseURL values and URLs they can generate:

​	如果 baseURL 以尾随斜杠结尾，则 DefaultURL 在锚定形式中在 ImportPath 和 # 之间插入斜杠。例如，以下是一些 baseURL 值及其可以生成的 URL：

```
"/pkg/" → "/pkg/math/#Sqrt"
"/pkg"  → "/pkg/math#Sqrt"
"/"     → "/math/#Sqrt"
""      → "/math#Sqrt"
```

### type Heading 

```go
type Heading struct {
	Text []Text // the heading text
}
```

A Heading is a doc comment heading.

​	标题是文档注释标题。

#### (*Heading) DefaultID

```go
func (h *Heading) DefaultID() string
```

DefaultID returns the default anchor ID for the heading h.

​	DefaultID 返回标题 h 的默认锚点 ID。

The default anchor ID is constructed by converting every rune that is not alphanumeric ASCII to an underscore and then adding the prefix “hdr-”. For example, if the heading text is “Go Doc Comments”, the default ID is “hdr-Go_Doc_Comments”.

​	默认锚点 ID 是通过将每个非字母数字 ASCII 字符转换为下划线，然后添加前缀“hdr-”来构建的。例如，如果标题文本是“Go Doc Comments”，则默认 ID 为“hdr-Go_Doc_Comments”。

### type Italic

```go
type Italic string
```

An Italic is a string rendered as italicized text.

​	Italic 是一个以斜体文本呈现的字符串。

### type Link

```go
type Link struct {
	Auto bool   // is this an automatic (implicit) link of a literal URL?
	Text []Text // text of link
	URL  string // target URL of link
}
```

A Link is a link to a specific URL.

​	Link 是指向特定 URL 的链接。

### type LinkDef 

```go
type LinkDef struct {
	Text string // the link text
	URL  string // the link URL
	Used bool   // whether the comment uses the definition
}
```

A LinkDef is a single link definition.

​	LinkDef 是一个单独的链接定义。

### type List

```go
type List struct {
	// Items is the list items.
	Items []*ListItem

	// ForceBlankBefore indicates that the list must be
	// preceded by a blank line when reformatting the comment,
	// overriding the usual conditions. See the BlankBefore method.
	//
	// The comment parser sets ForceBlankBefore for any list
	// that is preceded by a blank line, to make sure
	// the blank line is preserved when printing.
	ForceBlankBefore bool

	// ForceBlankBetween indicates that list items must be
	// separated by blank lines when reformatting the comment,
	// overriding the usual conditions. See the BlankBetween method.
	//
	// The comment parser sets ForceBlankBetween for any list
	// that has a blank line between any two of its items, to make sure
	// the blank lines are preserved when printing.
	ForceBlankBetween bool
}
```

A List is a numbered or bullet list. Lists are always non-empty: len(Items) > 0. In a numbered list, every Items[i].Number is a non-empty string. In a bullet list, every Items[i].Number is an empty string.

​	列表是编号或项目符号列表。列表始终非空：len(Items) > 0。在编号列表中，每个 Items[i].Number 是非空字符串。在项目符号列表中，每个 Items[i].Number 是空字符串。

#### (*List) BlankBefore

```go
func (l *List) BlankBefore() bool
```

BlankBefore reports whether a reformatting of the comment should include a blank line before the list. The default rule is the same as for [BlankBetween]: if the list item content contains any blank lines (meaning at least one item has multiple paragraphs) then the list itself must be preceded by a blank line. A preceding blank line can be forced by setting [List](https://pkg.go.dev/go/doc/comment@go1.20.1#List).ForceBlankBefore.

​	BlankBefore 报告注释的重新格式化是否应在列表前包含一个空行。默认规则与 [BlankBetween] 相同：如果列表项内容包含任何空行（意味着至少一个项目有多个段落），则列表本身必须以空行开头。可以通过设置 List.ForceBlankBefore 来强制使用前导空行。

#### (*List) BlankBetween

```go
func (l *List) BlankBetween() bool
```

BlankBetween reports whether a reformatting of the comment should include a blank line between each pair of list items. The default rule is that if the list item content contains any blank lines (meaning at least one item has multiple paragraphs) then list items must themselves be separated by blank lines. Blank line separators can be forced by setting [List](https://pkg.go.dev/go/doc/comment@go1.20.1#List).ForceBlankBetween.

​	BlankBetween 报告注释的重新格式化是否应在每对列表项之间包含一个空行。默认规则是，如果列表项内容包含任何空行（意味着至少一个项目有多个段落），则列表项本身必须用空行分隔。可以通过设置 List.ForceBlankBetween 来强制使用空行分隔符。

### type ListItem

```go
type ListItem struct {
	// Number is a decimal string in a numbered list
	// or an empty string in a bullet list.
	Number string // "1", "2", ...; "" for bullet list

	// Content is the list content.
	// Currently, restrictions in the parser and printer
	// require every element of Content to be a *Paragraph.
	Content []Block // Content of this item.
}
```

A ListItem is a single item in a numbered or bullet list.

​	ListItem 是编号或项目符号列表中的单个项目。

### type Paragraph

```go
type Paragraph struct {
	Text []Text
}
```

A Paragraph is a paragraph of text.

​	Paragraph 是一个文本段落。

### type Parser

```go
type Parser struct {
	// Words is a map of Go identifier words that
	// should be italicized and potentially linked.
	// If Words[w] is the empty string, then the word w
	// is only italicized. Otherwise it is linked, using
	// Words[w] as the link target.
	// Words corresponds to the [go/doc.ToHTML] words parameter.
	Words map[string]string

	// LookupPackage resolves a package name to an import path.
	//
	// If LookupPackage(name) returns ok == true, then [name]
	// (or [name.Sym] or [name.Sym.Method])
	// is considered a documentation link to importPath's package docs.
	// It is valid to return "", true, in which case name is considered
	// to refer to the current package.
	//
	// If LookupPackage(name) returns ok == false,
	// then [name] (or [name.Sym] or [name.Sym.Method])
	// will not be considered a documentation link,
	// except in the case where name is the full (but single-element) import path
	// of a package in the standard library, such as in [math] or [io.Reader].
	// LookupPackage is still called for such names,
	// in order to permit references to imports of other packages
	// with the same package names.
	//
	// Setting LookupPackage to nil is equivalent to setting it to
	// a function that always returns "", false.
	LookupPackage func(name string) (importPath string, ok bool)

	// LookupSym reports whether a symbol name or method name
	// exists in the current package.
	//
	// If LookupSym("", "Name") returns true, then [Name]
	// is considered a documentation link for a const, func, type, or var.
	//
	// Similarly, if LookupSym("Recv", "Name") returns true,
	// then [Recv.Name] is considered a documentation link for
	// type Recv's method Name.
	//
	// Setting LookupSym to nil is equivalent to setting it to a function
	// that always returns false.
	LookupSym func(recv, name string) (ok bool)
}
```

A Parser is a doc comment parser. The fields in the struct can be filled in before calling Parse in order to customize the details of the parsing process.

​	Parser 是一个文档注释解析器。在调用 Parse 之前，可以填充结构中的字段，以自定义解析过程的详细信息。

#### (*Parser) Parse

```go
func (p *Parser) Parse(text string) *Doc
```

Parse parses the doc comment text and returns the *Doc form. Comment markers (/* // and */) in the text must have already been removed.

​	Parse 解析文档注释文本并返回 Doc 形式。文本中的注释标记 (/ // 和 */) 必须已经删除。

### type Plain

```go
type Plain string
```

A Plain is a string rendered as plain text (not italicized).

​	Plain 是一个呈现为纯文本（非斜体）的字符串。

### type Printer

```go
type Printer struct {
	// HeadingLevel is the nesting level used for
	// HTML and Markdown headings.
	// If HeadingLevel is zero, it defaults to level 3,
	// meaning to use <h3> and ###.
	HeadingLevel int

	// HeadingID is a function that computes the heading ID
	// (anchor tag) to use for the heading h when generating
	// HTML and Markdown. If HeadingID returns an empty string,
	// then the heading ID is omitted.
	// If HeadingID is nil, h.DefaultID is used.
	HeadingID func(h *Heading) string

	// DocLinkURL is a function that computes the URL for the given DocLink.
	// If DocLinkURL is nil, then link.DefaultURL(p.DocLinkBaseURL) is used.
	DocLinkURL func(link *DocLink) string

	// DocLinkBaseURL is used when DocLinkURL is nil,
	// passed to [DocLink.DefaultURL] to construct a DocLink's URL.
	// See that method's documentation for details.
	DocLinkBaseURL string

	// TextPrefix is a prefix to print at the start of every line
	// when generating text output using the Text method.
	TextPrefix string

	// TextCodePrefix is the prefix to print at the start of each
	// preformatted (code block) line when generating text output,
	// instead of (not in addition to) TextPrefix.
	// If TextCodePrefix is the empty string, it defaults to TextPrefix+"\t".
	TextCodePrefix string

	// TextWidth is the maximum width text line to generate,
	// measured in Unicode code points,
	// excluding TextPrefix and the newline character.
	// If TextWidth is zero, it defaults to 80 minus the number of code points in TextPrefix.
	// If TextWidth is negative, there is no limit.
	TextWidth int
}
```

A Printer is a doc comment printer. The fields in the struct can be filled in before calling any of the printing methods in order to customize the details of the printing process.

​	Printer 是一个文档注释打印机。在调用任何打印方法之前，可以填充结构中的字段，以自定义打印过程的详细信息。

#### (*Printer) Comment

```go
func (p *Printer) Comment(d *Doc) []byte
```

Comment returns the standard Go formatting of the Doc, without any comment markers.

​	Comment 返回 Doc 的标准 Go 格式，没有任何注释标记。

#### (*Printer) HTML

```go
func (p *Printer) HTML(d *Doc) []byte
```

HTML returns an HTML formatting of the Doc. See the [Printer](https://pkg.go.dev/go/doc/comment@go1.20.1#Printer) documentation for ways to customize the HTML output.

​	HTML 返回 Doc 的 HTML 格式。有关自定义 HTML 输出的方法，请参阅 Printer 文档。

#### (*Printer) Markdown

```go
func (p *Printer) Markdown(d *Doc) []byte
```

Markdown returns a Markdown formatting of the Doc. See the [Printer](https://pkg.go.dev/go/doc/comment@go1.20.1#Printer) documentation for ways to customize the Markdown output.

​	Markdown 返回 Doc 的 Markdown 格式。有关自定义 Markdown 输出的方法，请参阅 Printer 文档。

#### (*Printer) Text

```go
func (p *Printer) Text(d *Doc) []byte
```

Text returns a textual formatting of the Doc. See the [Printer](https://pkg.go.dev/go/doc/comment@go1.20.1#Printer) documentation for ways to customize the text output.

​	Text 返回 Doc 的文本格式。有关自定义文本输出的方法，请参阅 Printer 文档。

### type Text

```go
type Text interface {
	// contains filtered or unexported methods
}
```

A Text is text-level content in a doc comment, one of [Plain](https://pkg.go.dev/go/doc/comment@go1.20.1#Plain), [Italic](https://pkg.go.dev/go/doc/comment@go1.20.1#Italic), [*Link](https://pkg.go.dev/go/doc/comment@go1.20.1#Link), or [*DocLink](https://pkg.go.dev/go/doc/comment@go1.20.1#DocLink).

​	Text 是文档注释中的文本级内容，包括 Plain、Italic、*Link 或 *DocLink 之一。