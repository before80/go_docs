+++
title = "doc"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/doc@go1.21.3](https://pkg.go.dev/go/doc@go1.21.3)

Package doc extracts source code documentation from a Go AST.

​	 doc 包从 Go AST 中提取源代码文档。

## 常量

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/go/doc/synopsis.go;l=47)

```go
var IllegalPrefixes = []string{
	"copyright",
	"all rights",
	"author",
}
```

IllegalPrefixes is a list of lower-case prefixes that identify a comment as not being a doc comment. This helps to avoid misinterpreting the common mistake of a copyright notice immediately before a package statement as being a doc comment.

​	IllegalPrefixes 是一个由小写前缀组成的列表，这些前缀标识注释不是文档注释。这有助于避免误解包语句前紧跟的版权声明是文档注释的常见错误。

## 函数

### func IsPredeclared <- go1.8

```go
func IsPredeclared(s string) bool
```

IsPredeclared reports whether s is a predeclared identifier.

​	IsPredeclared 报告 s 是否是预声明标识符。

### func Synopsis <-DEPRECATED

```
func Synopsis(text string) string
```

Synopsis returns a cleaned version of the first sentence in text.

​	Synopsis 返回文本中第一句话的已清理版本。

Deprecated: New programs should use [Package.Synopsis](https://pkg.go.dev/go/doc@go1.21.3#Package.Synopsis) instead, which handles links in text properly.

​	已弃用：新程序应使用 Package.Synopsis，它可以正确处理文本中的链接。

### func ToHTML<-DEPRECATED 

```go
func ToHTML(w io.Writer, text string, words map[string]string)
```

ToHTML converts comment text to formatted HTML.

​	ToHTML 将注释文本转换为格式化的 HTML。

Deprecated: ToHTML cannot identify documentation links in the doc comment, because they depend on knowing what package the text came from, which is not included in this API.

​	已弃用：ToHTML 无法识别文档注释中的文档链接，因为它们依赖于知道文本来自哪个包，而这不在此 API 中包含。

Given the *[doc.Package](https://pkg.go.dev/go/doc@go1.21.3#Package) p where text was found, ToHTML(w, text, nil) can be replaced by:

​	给定找到文本的 * doc.Package p，ToHTML(w, text, nil) 可以替换为：

```go
w.Write(p.HTML(text))
```

which is in turn shorthand for:

​	这是以下内容的简写：

```go
w.Write(p.Printer().HTML(p.Parser().Parse(text)))
```

If words may be non-nil, the longer replacement is:

​	如果单词可能为非 nil，则较长的替换内容为：

```go
parser := p.Parser()
parser.Words = words
w.Write(p.Printer().HTML(parser.Parse(d)))
```

### func ToText <-DEPRECATED 

```go
func ToText(w io.Writer, text string, prefix, codePrefix string, width int)
```

ToText converts comment text to formatted text.

​	ToText 将注释文本转换为格式化文本。

Deprecated: ToText cannot identify documentation links in the doc comment, because they depend on knowing what package the text came from, which is not included in this API.

​	已弃用：ToText 无法识别文档注释中的文档链接，因为它们依赖于知道文本来自哪个包，而这不在此 API 中包含。

Given the *[doc.Package](https://pkg.go.dev/go/doc@go1.21.3#Package) p where text was found, ToText(w, text, “”, “\t”, 80) can be replaced by:

​	给定发现文本的 *doc.Package p，ToText(w, text, “”, “\t”, 80) 可以替换为：

```go
w.Write(p.Text(text))
```

In the general case, ToText(w, text, prefix, codePrefix, width) can be replaced by:

​	在一般情况下，ToText(w, text, prefix, codePrefix, width) 可以替换为：

```go
d := p.Parser().Parse(text)
pr := p.Printer()
pr.TextPrefix = prefix
pr.TextCodePrefix = codePrefix
pr.TextWidth = width
w.Write(pr.Text(d))
```

See the documentation for [Package.Text](https://pkg.go.dev/go/doc@go1.21.3#Package.Text) and [comment.Printer.Text](https://pkg.go.dev/go/doc/comment#Printer.Text) for more details.

​	有关更多详细信息，请参阅 Package.Text 和 comment.Printer.Text 的文档。

## 类型

### type Example

```go
type Example struct {
	Name        string // name of the item being exemplified (including optional suffix)
	Suffix      string // example suffix, without leading '_' (only populated by NewFromFiles)
	Doc         string // example function doc string
	Code        ast.Node
	Play        *ast.File // a whole program version of the example
	Comments    []*ast.CommentGroup
	Output      string // expected output
	Unordered   bool
	EmptyOutput bool // expect empty output
	Order       int  // original source code order
}
```

An Example represents an example function found in a test source file.

​	Example 表示在测试源文件中找到的示例函数。

#### func Examples

```go
func Examples(testFiles ...*ast.File) []*Example
```

Examples returns the examples found in testFiles, sorted by Name field. The Order fields record the order in which the examples were encountered. The Suffix field is not populated when Examples is called directly, it is only populated by NewFromFiles for examples it finds in _test.go files.

​	示例返回在 testFiles 中找到的示例，按 Name 字段排序。Order 字段记录遇到示例的顺序。直接调用 Examples 时不会填充 Suffix 字段，仅当 NewFromFiles 在 _test.go 文件中找到示例时才会填充该字段。

Playable Examples must be in a package whose name ends in “_test”. An Example is “playable” (the Play field is non-nil) in either of these circumstances:

​	可播放示例必须位于名称以“_test”结尾的包中。示例在以下任一情况下为“可播放”（Play 字段为非 nil）：

- The example function is self-contained: the function references only identifiers from other packages (or predeclared identifiers, such as “int”) and the test file does not include a dot import.
  示例函数是独立的：该函数仅引用其他包的标识符（或预声明的标识符，例如“int”），并且测试文件不包含点导入。
- The entire test file is the example: the file contains exactly one example function, zero test, fuzz test, or benchmark function, and at least one top-level function, type, variable, or constant declaration other than the example function.
  整个测试文件都是示例：该文件恰好包含一个示例函数、零个测试、模糊测试或基准函数，以及除示例函数之外的至少一个顶级函数、类型、变量或常量声明。

### type Filter

```go
type Filter func(string) bool
```

### type Func

```go
type Func struct {
	Doc  string
	Name string
	Decl *ast.FuncDecl

	// methods
	// (for functions, these fields have the respective zero value)
	Recv  string // actual   receiver "T" or "*T" possibly followed by type parameters [P1, ..., Pn]
	Orig  string // original receiver "T" or "*T"
	Level int    // embedding level; 0 means not embedded

	// Examples is a sorted list of examples associated with this
	// function or method. Examples are extracted from _test.go files
	// provided to NewFromFiles.
	Examples []*Example
}
```

Func is the documentation for a func declaration.

​	Func 是 func 声明的文档。

### type Mode

```go
type Mode int
```

Mode values control the operation of New and NewFromFiles.

​	Mode 值控制 New 和 NewFromFiles 的操作。

```go
const (
	// AllDecls says to extract documentation for all package-level
	// declarations, not just exported ones.
	AllDecls Mode = 1 << iota

	// AllMethods says to show all embedded methods, not just the ones of
	// invisible (unexported) anonymous fields.
	AllMethods

	// PreserveAST says to leave the AST unmodified. Originally, pieces of
	// the AST such as function bodies were nil-ed out to save memory in
	// godoc, but not all programs want that behavior.
	PreserveAST
)
```

### type Note <- go1.1

```go
type Note struct {
	Pos, End token.Pos // position range of the comment containing the marker
	UID      string    // uid found with the marker
	Body     string    // note body text
}
```

A Note represents a marked comment starting with “MARKER(uid): note body”. Any note with a marker of 2 or more upper case [A-Z] letters and a uid of at least one character is recognized. The “:” following the uid is optional. Notes are collected in the Package.Notes map indexed by the notes marker.

​	注释表示以“MARKER(uid)：注释正文”开头的标记注释。识别任何标记为 2 个或更多大写 [A-Z] 字母且 uid 至少为一个字符的注释。uid 后面的“：”是可选的。注释收集在 Package.Notes 映射中，由注释标记索引。

### type Package

```go
type Package struct {
	Doc        string
	Name       string
	ImportPath string
	Imports    []string
	Filenames  []string
	Notes      map[string][]*Note

	// Deprecated: For backward compatibility Bugs is still populated,
	// but all new code should use Notes instead.
	Bugs []string

	// declarations
	Consts []*Value
	Types  []*Type
	Vars   []*Value
	Funcs  []*Func

	// Examples is a sorted list of examples associated with
	// the package. Examples are extracted from _test.go files
	// provided to NewFromFiles.
	Examples []*Example
	// contains filtered or unexported fields
}
```

Package is the documentation for an entire package.

​	Package 是整个包的文档。

#### func New

```go
func New(pkg *ast.Package, importPath string, mode Mode) *Package
```

New computes the package documentation for the given package AST. New takes ownership of the AST pkg and may edit or overwrite it. To have the Examples fields populated, use NewFromFiles and include the package’s _test.go files.

​	New 计算给定包 AST 的包文档。New 拥有 AST pkg 并可以编辑或覆盖它。要填充 Examples 字段，请使用 NewFromFiles 并包含包的 _test.go 文件。

#### func NewFromFiles <- go1.14

```go
func NewFromFiles(fset *token.FileSet, files []*ast.File, importPath string, opts ...any) (*Package, error)
```

NewFromFiles computes documentation for a package.

​	NewFromFiles 计算包的文档。

The package is specified by a list of *ast.Files and corresponding file set, which must not be nil. NewFromFiles uses all provided files when computing documentation, so it is the caller’s responsibility to provide only the files that match the desired build context. “go/build”.Context.MatchFile can be used for determining whether a file matches a build context with the desired GOOS and GOARCH values, and other build constraints. The import path of the package is specified by importPath.

​	包由一组 *ast.Files 和相应的非空文件集指定。NewFromFiles 在计算文档时使用所有提供的文件，因此由调用者负责仅提供与所需构建上下文匹配的文件。“go/build”.Context.MatchFile 可用于确定文件是否与具有所需 GOOS 和 GOARCH 值以及其他构建约束的构建上下文匹配。包的导入路径由 importPath 指定。

Examples found in _test.go files are associated with the corresponding type, function, method, or the package, based on their name. If the example has a suffix in its name, it is set in the Example.Suffix field. Examples with malformed names are skipped.

​	在 _test.go 文件中找到的示例与相应的类型、函数、方法或包相关联，具体取决于它们的名称。如果示例的名称中有后缀，则将其设置在 Example.Suffix 字段中。具有格式错误名称的示例将被跳过。

Optionally, a single extra argument of type Mode can be provided to control low-level aspects of the documentation extraction behavior.

​	或者，可以提供一个类型为 Mode 的单一额外参数来控制文档提取行为的底层方面。

NewFromFiles takes ownership of the AST files and may edit them, unless the PreserveAST Mode bit is on.

​	NewFromFiles 拥有 AST 文件的所有权，并且可能会编辑它们，除非 PreserveAST Mode 位已打开。

##### NewFromFiles Example 

```go
package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
)

func main() {
	// src and test are two source files that make up
	// a package whose documentation will be computed.
	const src = `
// This is the package comment.
package p

import "fmt"

// This comment is associated with the Greet function.
func Greet(who string) {
	fmt.Printf("Hello, %s!\n", who)
}
`
	const test = `
package p_test

// This comment is associated with the ExampleGreet_world example.
func ExampleGreet_world() {
	Greet("world")
}
`

	// Create the AST by parsing src and test.
	fset := token.NewFileSet()
	files := []*ast.File{
		mustParse(fset, "src.go", src),
		mustParse(fset, "src_test.go", test),
	}

	// Compute package documentation with examples.
	p, err := doc.NewFromFiles(fset, files, "example.com/p")
	if err != nil {
		panic(err)
	}

	fmt.Printf("package %s - %s", p.Name, p.Doc)
	fmt.Printf("func %s - %s", p.Funcs[0].Name, p.Funcs[0].Doc)
	fmt.Printf(" ⤷ example with suffix %q - %s", p.Funcs[0].Examples[0].Suffix, p.Funcs[0].Examples[0].Doc)

}

func mustParse(fset *token.FileSet, filename, src string) *ast.File {
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return f
}
Output:

package p - This is the package comment.
func Greet - This comment is associated with the Greet function.
 ⤷ example with suffix "world" - This comment is associated with the ExampleGreet_world example.
```

#### (*Package) Filter

```go
func (p *Package) Filter(f Filter)
```

Filter eliminates documentation for names that don’t pass through the filter f. TODO(gri): Recognize “Type.Method” as a name.

​	Filter 消除了不通过过滤器 f 的名称的文档。TODO(gri)：将“Type.Method”识别为一个名称。

#### (*Package) HTML <- go1.19

```go
func (p *Package) HTML(text string) []byte
```

HTML returns formatted HTML for the doc comment text.

​	HTML 为文档注释文本返回格式化的 HTML。

To customize details of the HTML, use [Package.Printer](https://pkg.go.dev/go/doc@go1.20.1#Package.Printer) to obtain a [comment.Printer](https://pkg.go.dev/go/doc/comment#Printer), and configure it before calling its HTML method.

​	要自定义 HTML 的详细信息，请使用 Package.Printer 获取 comment.Printer，并在调用其 HTML 方法之前对其进行配置。

#### (*Package) Markdown <- go1.19

```go
func (p *Package) Markdown(text string) []byte
```

Markdown returns formatted Markdown for the doc comment text.

​	Markdown 返回格式化的 Markdown 文档注释文本。

To customize details of the Markdown, use [Package.Printer](https://pkg.go.dev/go/doc@go1.20.1#Package.Printer) to obtain a [comment.Printer](https://pkg.go.dev/go/doc/comment#Printer), and configure it before calling its Markdown method.

​	要自定义 Markdown 的详细信息，请使用 Package.Printer 获取 comment.Printer，并在调用其 Markdown 方法之前对其进行配置。

#### (*Package) Parser <- go1.19

```go
func (p *Package) Parser() *comment.Parser
```

Parser returns a doc comment parser configured for parsing doc comments from package p. Each call returns a new parser, so that the caller may customize it before use.

​	Parser 返回一个文档注释解析器，该解析器配置为解析包 p 中的文档注释。每次调用都会返回一个新的解析器，以便调用者在使用前对其进行自定义。

#### (*Package) Printer <- go1.19

```go
func (p *Package) Printer() *comment.Printer
```

Printer returns a doc comment printer configured for printing doc comments from package p. Each call returns a new printer, so that the caller may customize it before use.

​	Printer 返回一个文档注释打印机，该打印机配置为打印包 p 中的文档注释。每次调用都会返回一个新的打印机，以便调用者在使用前对其进行自定义。

#### (*Package) Synopsis <- go1.19

```go
func (p *Package) Synopsis(text string) string
```

Synopsis returns a cleaned version of the first sentence in text. That sentence ends after the first period followed by space and not preceded by exactly one uppercase letter, or at the first paragraph break. The result string has no \n, \r, or \t characters and uses only single spaces between words. If text starts with any of the IllegalPrefixes, the result is the empty string.

​	Synopsis 返回文本中第一句话的清理版本。该句子在第一个句号后结束，后面跟一个空格，前面没有一个大写字母，或者在第一个段落中断处结束。结果字符串没有 \n、\r 或 \t 字符，并且单词之间只使用单个空格。如果文本以任何 IllegalPrefixes 开头，则结果为空字符串。

#### (*Package) Text <- go1.19

```go
func (p *Package) Text(text string) []byte
```

Text returns formatted text for the doc comment text, wrapped to 80 Unicode code points and using tabs for code block indentation.

​	Text 返回格式化的文档注释文本，换行至 80 个 Unicode 代码点，并使用制表符进行代码块缩进。

To customize details of the formatting, use [Package.Printer](https://pkg.go.dev/go/doc@go1.20.1#Package.Printer) to obtain a [comment.Printer](https://pkg.go.dev/go/doc/comment#Printer), and configure it before calling its Text method.

​	要自定义格式化的详细信息，请使用 Package.Printer 获取 comment.Printer，并在调用其 Text 方法之前对其进行配置。

### type Type

```go
type Type struct {
	Doc  string
	Name string
	Decl *ast.GenDecl

	// associated declarations
	Consts  []*Value // sorted list of constants of (mostly) this type
	Vars    []*Value // sorted list of variables of (mostly) this type
	Funcs   []*Func  // sorted list of functions returning this type
	Methods []*Func  // sorted list of methods (including embedded ones) of this type

	// Examples is a sorted list of examples associated with
	// this type. Examples are extracted from _test.go files
	// provided to NewFromFiles.
	Examples []*Example
}
```

Type is the documentation for a type declaration.

​	Type 是类型声明的文档。

### type Value 

```go
type Value struct {
	Doc   string
	Names []string // var or const names in declaration order
	Decl  *ast.GenDecl
	// contains filtered or unexported fields
}
```

Value is the documentation for a (possibly grouped) var or const declaration.

​	Value 是（可能分组的）var 或 const 声明的文档。