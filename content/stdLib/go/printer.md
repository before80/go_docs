+++
title = "printer"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# printer

https://pkg.go.dev/go/printer@go1.20.1



Package printer implements printing of AST nodes.







## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func Fprint 

``` go 
func Fprint(output io.Writer, fset *token.FileSet, node any) error
```

Fprint "pretty-prints" an AST node to output. It calls Config.Fprint with default settings. Note that gofmt uses tabs for indentation but spaces for alignment; use format.Node (package go/format) for output that matches gofmt.

##### Example
``` go 
```

## 类型

### type CommentedNode 

``` go 
type CommentedNode struct {
	Node     any // *ast.File, or ast.Expr, ast.Decl, ast.Spec, or ast.Stmt
	Comments []*ast.CommentGroup
}
```

A CommentedNode bundles an AST node and corresponding comments. It may be provided as argument to any of the Fprint functions.

### type Config 

``` go 
type Config struct {
	Mode     Mode // default: 0
	Tabwidth int  // default: 8
	Indent   int  // default: 0 (all code is indented at least by this much)
}
```

A Config node controls the output of Fprint.

#### (*Config) Fprint 

``` go 
func (cfg *Config) Fprint(output io.Writer, fset *token.FileSet, node any) error
```

Fprint "pretty-prints" an AST node to output for a given configuration cfg. Position information is interpreted relative to the file set fset. The node type must be *ast.File, *CommentedNode, []ast.Decl, []ast.Stmt, or assignment-compatible to ast.Expr, ast.Decl, ast.Spec, or ast.Stmt.

### type Mode 

``` go 
type Mode uint
```

A Mode value is a set of flags (or 0). They control printing.

``` go 
const (
	RawFormat Mode = 1 << iota // do not use a tabwriter; if set, UseSpaces is ignored
	TabIndent                  // use tabs for indentation independent of UseSpaces
	UseSpaces                  // use spaces instead of tabs for alignment
	SourcePos                  // emit //line directives to preserve original source positions
)
```