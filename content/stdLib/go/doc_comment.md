+++
title = "doc/comment"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/doc/comment@go1.21.3](https://pkg.go.dev/go/doc/comment@go1.21.3)

Package comment implements parsing and reformatting of Go doc comments, (documentation comments), which are comments that immediately precede a top-level declaration of a package, const, func, type, or var.

Go doc comment syntax is a simplified subset of Markdown that supports links, headings, paragraphs, lists (without nesting), and preformatted text blocks. The details of the syntax are documented at https://go.dev/doc/comment.

To parse the text associated with a doc comment (after removing comment markers), use a [Parser](https://pkg.go.dev/go/doc/comment@go1.20.1#Parser):

``` go 
var p comment.Parser
doc := p.Parse(text)
```

The result is a [*Doc](https://pkg.go.dev/go/doc/comment@go1.20.1#Doc). To reformat it as a doc comment, HTML, Markdown, or plain text, use a [Printer](https://pkg.go.dev/go/doc/comment@go1.20.1#Printer):

``` go 
var pr comment.Printer
os.Stdout.Write(pr.Text(doc))
```

The [Parser](https://pkg.go.dev/go/doc/comment@go1.20.1#Parser) and [Printer](https://pkg.go.dev/go/doc/comment@go1.20.1#Printer) types are structs whose fields can be modified to customize the operations. For details, see the documentation for those types.

Use cases that need additional control over reformatting can implement their own logic by inspecting the parsed syntax itself. See the documentation for [Doc](https://pkg.go.dev/go/doc/comment@go1.20.1#Doc), [Block](https://pkg.go.dev/go/doc/comment@go1.20.1#Block), [Text](https://pkg.go.dev/go/doc/comment@go1.20.1#Text) for an overview and links to additional types.



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func DefaultLookupPackage 

``` go 
func DefaultLookupPackage(name string) (importPath string, ok bool)
```

DefaultLookupPackage is the default package lookup function, used when [Parser](https://pkg.go.dev/go/doc/comment@go1.20.1#Parser).LookupPackage is nil. It recognizes names of the packages from the standard library with single-element import paths, such as math, which would otherwise be impossible to name.

Note that the go/doc package provides a more sophisticated lookup based on the imports used in the current package.

## 类型

### type Block 

``` go 
type Block interface {
	// contains filtered or unexported methods
}
```

A Block is block-level content in a doc comment, one of [*Code](https://pkg.go.dev/go/doc/comment@go1.20.1#Code), [*Heading](https://pkg.go.dev/go/doc/comment@go1.20.1#Heading), [*List](https://pkg.go.dev/go/doc/comment@go1.20.1#List), or [*Paragraph](https://pkg.go.dev/go/doc/comment@go1.20.1#Paragraph).

### type Code 

``` go 
type Code struct {
	// Text is the preformatted text, ending with a newline character.
	// It may be multiple lines, each of which ends with a newline character.
	// It is never empty, nor does it start or end with a blank line.
	Text string
}
```

A Code is a preformatted code block.

### type Doc 

``` go 
type Doc struct {
	// Content is the sequence of content blocks in the comment.
	Content []Block

	// Links is the link definitions in the comment.
	Links []*LinkDef
}
```

A Doc is a parsed Go doc comment.

### type DocLink 

``` go 
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

#### (*DocLink) DefaultURL 

``` go 
func (l *DocLink) DefaultURL(baseURL string) string
```

DefaultURL constructs and returns the documentation URL for l, using baseURL as a prefix for links to other packages.

The possible forms returned by DefaultURL are:

- baseURL/ImportPath, for a link to another package
- baseURL/ImportPath#Name, for a link to a const, func, type, or var in another package
- baseURL/ImportPath#Recv.Name, for a link to a method in another package
- \#Name, for a link to a const, func, type, or var in this package
- \#Recv.Name, for a link to a method in this package

If baseURL ends in a trailing slash, then DefaultURL inserts a slash between ImportPath and # in the anchored forms. For example, here are some baseURL values and URLs they can generate:

```
"/pkg/" → "/pkg/math/#Sqrt"
"/pkg"  → "/pkg/math#Sqrt"
"/"     → "/math/#Sqrt"
""      → "/math#Sqrt"
```

### type Heading 

``` go 
type Heading struct {
	Text []Text // the heading text
}
```

A Heading is a doc comment heading.

#### (*Heading) DefaultID 

``` go 
func (h *Heading) DefaultID() string
```

DefaultID returns the default anchor ID for the heading h.

The default anchor ID is constructed by converting every rune that is not alphanumeric ASCII to an underscore and then adding the prefix "hdr-". For example, if the heading text is "Go Doc Comments", the default ID is "hdr-Go_Doc_Comments".

### type Italic 

``` go 
type Italic string
```

An Italic is a string rendered as italicized text.

### type Link 

``` go 
type Link struct {
	Auto bool   // is this an automatic (implicit) link of a literal URL?
	Text []Text // text of link
	URL  string // target URL of link
}
```

A Link is a link to a specific URL.

### type LinkDef 

``` go 
type LinkDef struct {
	Text string // the link text
	URL  string // the link URL
	Used bool   // whether the comment uses the definition
}
```

A LinkDef is a single link definition.

### type List 

``` go 
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

#### (*List) BlankBefore 

``` go 
func (l *List) BlankBefore() bool
```

BlankBefore reports whether a reformatting of the comment should include a blank line before the list. The default rule is the same as for [BlankBetween]: if the list item content contains any blank lines (meaning at least one item has multiple paragraphs) then the list itself must be preceded by a blank line. A preceding blank line can be forced by setting [List](https://pkg.go.dev/go/doc/comment@go1.20.1#List).ForceBlankBefore.

#### (*List) BlankBetween 

``` go 
func (l *List) BlankBetween() bool
```

BlankBetween reports whether a reformatting of the comment should include a blank line between each pair of list items. The default rule is that if the list item content contains any blank lines (meaning at least one item has multiple paragraphs) then list items must themselves be separated by blank lines. Blank line separators can be forced by setting [List](https://pkg.go.dev/go/doc/comment@go1.20.1#List).ForceBlankBetween.

### type ListItem 

``` go 
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

### type Paragraph 

``` go 
type Paragraph struct {
	Text []Text
}
```

A Paragraph is a paragraph of text.

### type Parser 

``` go 
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

#### (*Parser) Parse 

``` go 
func (p *Parser) Parse(text string) *Doc
```

Parse parses the doc comment text and returns the *Doc form. Comment markers (/* // and */) in the text must have already been removed.

### type Plain 

``` go 
type Plain string
```

A Plain is a string rendered as plain text (not italicized).

### type Printer 

``` go 
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

#### (*Printer) Comment 

``` go 
func (p *Printer) Comment(d *Doc) []byte
```

Comment returns the standard Go formatting of the Doc, without any comment markers.

#### (*Printer) HTML 

``` go 
func (p *Printer) HTML(d *Doc) []byte
```

HTML returns an HTML formatting of the Doc. See the [Printer](https://pkg.go.dev/go/doc/comment@go1.20.1#Printer) documentation for ways to customize the HTML output.

#### (*Printer) Markdown 

``` go 
func (p *Printer) Markdown(d *Doc) []byte
```

Markdown returns a Markdown formatting of the Doc. See the [Printer](https://pkg.go.dev/go/doc/comment@go1.20.1#Printer) documentation for ways to customize the Markdown output.

#### (*Printer) Text 

``` go 
func (p *Printer) Text(d *Doc) []byte
```

Text returns a textual formatting of the Doc. See the [Printer](https://pkg.go.dev/go/doc/comment@go1.20.1#Printer) documentation for ways to customize the text output.

### type Text 

``` go 
type Text interface {
	// contains filtered or unexported methods
}
```

A Text is text-level content in a doc comment, one of [Plain](https://pkg.go.dev/go/doc/comment@go1.20.1#Plain), [Italic](https://pkg.go.dev/go/doc/comment@go1.20.1#Italic), [*Link](https://pkg.go.dev/go/doc/comment@go1.20.1#Link), or [*DocLink](https://pkg.go.dev/go/doc/comment@go1.20.1#DocLink).