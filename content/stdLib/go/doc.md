+++
title = "doc"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/go/doc@go1.21.3

Package doc extracts source code documentation from a Go AST.

## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/go/doc/synopsis.go;l=47)

``` go 
var IllegalPrefixes = []string{
	"copyright",
	"all rights",
	"author",
}
```

IllegalPrefixes is a list of lower-case prefixes that identify a comment as not being a doc comment. This helps to avoid misinterpreting the common mistake of a copyright notice immediately before a package statement as being a doc comment.

## 函数

### func IsPredeclared  <- go1.8

``` go 
func IsPredeclared(s string) bool
```

IsPredeclared reports whether s is a predeclared identifier.

### func Synopsis <-DEPRECATED

```
func Synopsis(text string) string
```

Synopsis returns a cleaned version of the first sentence in text.

Deprecated: New programs should use [Package.Synopsis](https://pkg.go.dev/go/doc@go1.21.3#Package.Synopsis) instead, which handles links in text properly.

### func ToHTML<-DEPRECATED

```go
func ToHTML(w io.Writer, text string, words map[string]string)
```

ToHTML converts comment text to formatted HTML.

Deprecated: ToHTML cannot identify documentation links in the doc comment, because they depend on knowing what package the text came from, which is not included in this API.

Given the *[doc.Package](https://pkg.go.dev/go/doc@go1.21.3#Package) p where text was found, ToHTML(w, text, nil) can be replaced by:

```go
w.Write(p.HTML(text))
```

which is in turn shorthand for:

```go
w.Write(p.Printer().HTML(p.Parser().Parse(text)))
```

If words may be non-nil, the longer replacement is:

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

Deprecated: ToText cannot identify documentation links in the doc comment, because they depend on knowing what package the text came from, which is not included in this API.

Given the *[doc.Package](https://pkg.go.dev/go/doc@go1.21.3#Package) p where text was found, ToText(w, text, "", "\t", 80) can be replaced by:

```go
w.Write(p.Text(text))
```

In the general case, ToText(w, text, prefix, codePrefix, width) can be replaced by:

```go
d := p.Parser().Parse(text)
pr := p.Printer()
pr.TextPrefix = prefix
pr.TextCodePrefix = codePrefix
pr.TextWidth = width
w.Write(pr.Text(d))
```

See the documentation for [Package.Text](https://pkg.go.dev/go/doc@go1.21.3#Package.Text) and [comment.Printer.Text](https://pkg.go.dev/go/doc/comment#Printer.Text) for more details.

## 类型

### type Example 

``` go 
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

#### func Examples 

``` go 
func Examples(testFiles ...*ast.File) []*Example
```

Examples returns the examples found in testFiles, sorted by Name field. The Order fields record the order in which the examples were encountered. The Suffix field is not populated when Examples is called directly, it is only populated by NewFromFiles for examples it finds in _test.go files.

Playable Examples must be in a package whose name ends in "_test". An Example is "playable" (the Play field is non-nil) in either of these circumstances:

- The example function is self-contained: the function references only identifiers from other packages (or predeclared identifiers, such as "int") and the test file does not include a dot import.
- The entire test file is the example: the file contains exactly one example function, zero test, fuzz test, or benchmark function, and at least one top-level function, type, variable, or constant declaration other than the example function.

### type Filter 

``` go 
type Filter func(string) bool
```

### type Func 

``` go 
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

### type Mode 

``` go 
type Mode int
```

Mode values control the operation of New and NewFromFiles.

``` go 
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

### type Note  <- go1.1

``` go 
type Note struct {
	Pos, End token.Pos // position range of the comment containing the marker
	UID      string    // uid found with the marker
	Body     string    // note body text
}
```

A Note represents a marked comment starting with "MARKER(uid): note body". Any note with a marker of 2 or more upper case [A-Z] letters and a uid of at least one character is recognized. The ":" following the uid is optional. Notes are collected in the Package.Notes map indexed by the notes marker.

### type Package 

``` go 
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

#### func New 

``` go 
func New(pkg *ast.Package, importPath string, mode Mode) *Package
```

New computes the package documentation for the given package AST. New takes ownership of the AST pkg and may edit or overwrite it. To have the Examples fields populated, use NewFromFiles and include the package's _test.go files.

#### func NewFromFiles  <- go1.14

``` go 
func NewFromFiles(fset *token.FileSet, files []*ast.File, importPath string, opts ...any) (*Package, error)
```

NewFromFiles computes documentation for a package.

The package is specified by a list of *ast.Files and corresponding file set, which must not be nil. NewFromFiles uses all provided files when computing documentation, so it is the caller's responsibility to provide only the files that match the desired build context. "go/build".Context.MatchFile can be used for determining whether a file matches a build context with the desired GOOS and GOARCH values, and other build constraints. The import path of the package is specified by importPath.

Examples found in _test.go files are associated with the corresponding type, function, method, or the package, based on their name. If the example has a suffix in its name, it is set in the Example.Suffix field. Examples with malformed names are skipped.

Optionally, a single extra argument of type Mode can be provided to control low-level aspects of the documentation extraction behavior.

NewFromFiles takes ownership of the AST files and may edit them, unless the PreserveAST Mode bit is on.

##### NewFromFiles  Example
``` go 
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

``` go 
func (p *Package) Filter(f Filter)
```

Filter eliminates documentation for names that don't pass through the filter f. TODO(gri): Recognize "Type.Method" as a name.

#### (*Package) HTML  <- go1.19

``` go 
func (p *Package) HTML(text string) []byte
```

HTML returns formatted HTML for the doc comment text.

To customize details of the HTML, use [Package.Printer](https://pkg.go.dev/go/doc@go1.20.1#Package.Printer) to obtain a [comment.Printer](https://pkg.go.dev/go/doc/comment#Printer), and configure it before calling its HTML method.

#### (*Package) Markdown  <- go1.19

``` go 
func (p *Package) Markdown(text string) []byte
```

Markdown returns formatted Markdown for the doc comment text.

To customize details of the Markdown, use [Package.Printer](https://pkg.go.dev/go/doc@go1.20.1#Package.Printer) to obtain a [comment.Printer](https://pkg.go.dev/go/doc/comment#Printer), and configure it before calling its Markdown method.

#### (*Package) Parser  <- go1.19

``` go 
func (p *Package) Parser() *comment.Parser
```

Parser returns a doc comment parser configured for parsing doc comments from package p. Each call returns a new parser, so that the caller may customize it before use.

#### (*Package) Printer  <- go1.19

``` go 
func (p *Package) Printer() *comment.Printer
```

Printer returns a doc comment printer configured for printing doc comments from package p. Each call returns a new printer, so that the caller may customize it before use.

#### (*Package) Synopsis  <- go1.19

``` go 
func (p *Package) Synopsis(text string) string
```

Synopsis returns a cleaned version of the first sentence in text. That sentence ends after the first period followed by space and not preceded by exactly one uppercase letter, or at the first paragraph break. The result string has no \n, \r, or \t characters and uses only single spaces between words. If text starts with any of the IllegalPrefixes, the result is the empty string.

#### (*Package) Text  <- go1.19

``` go 
func (p *Package) Text(text string) []byte
```

Text returns formatted text for the doc comment text, wrapped to 80 Unicode code points and using tabs for code block indentation.

To customize details of the formatting, use [Package.Printer](https://pkg.go.dev/go/doc@go1.20.1#Package.Printer) to obtain a [comment.Printer](https://pkg.go.dev/go/doc/comment#Printer), and configure it before calling its Text method.

### type Type 

``` go 
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

### type Value 

``` go 
type Value struct {
	Doc   string
	Names []string // var or const names in declaration order
	Decl  *ast.GenDecl
	// contains filtered or unexported fields
}
```

Value is the documentation for a (possibly grouped) var or const declaration.