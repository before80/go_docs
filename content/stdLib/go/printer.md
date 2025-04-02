+++
title = "printer"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/printer@go1.24.2](https://pkg.go.dev/go/printer@go1.24.2)

Package printer implements printing of AST nodes.

​	Package printer 实现 AST 节点的打印。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Fprint 

``` go 
func Fprint(output io.Writer, fset *token.FileSet, node any) error
```

Fprint “pretty-prints” an AST node to output. It calls Config.Fprint with default settings. Note that gofmt uses tabs for indentation but spaces for alignment; use format.Node (package go/format) for output that matches gofmt.

​	Fprint 将 AST 节点“漂亮地打印”到输出。它使用默认设置调用 Config.Fprint。请注意，gofmt 使用制表符进行缩进，但使用空格进行对齐；使用 format.Node（包 go/format）以获得与 gofmt 匹配的输出。

#### Fprint Example

```go
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
)

func parseFunc(filename, functionname string) (fun *ast.FuncDecl, fset *token.FileSet) {
	fset = token.NewFileSet()
	if file, err := parser.ParseFile(fset, filename, nil, 0); err == nil {
		for _, d := range file.Decls {
			if f, ok := d.(*ast.FuncDecl); ok && f.Name.Name == functionname {
				fun = f
				return
			}
		}
	}
	panic("function not found")
}

func printSelf() {
	// Parse source file and extract the AST without comments for
	// this function, with position information referring to the
	// file set fset.
	funcAST, fset := parseFunc("example_test.go", "printSelf")

	// Print the function body into buffer buf.
	// The file set is provided to the printer so that it knows
	// about the original source formatting and can add additional
	// line breaks where they were present in the source.
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, funcAST.Body)

	// Remove braces {} enclosing the function body, unindent,
	// and trim leading and trailing white space.
	s := buf.String()
	s = s[1 : len(s)-1]
	s = strings.TrimSpace(strings.ReplaceAll(s, "\n\t", "\n"))

	// Print the cleaned-up body text to stdout.
	fmt.Println(s)
}

func main() {
	printSelf()

}
Output:

funcAST, fset := parseFunc("example_test.go", "printSelf")

var buf bytes.Buffer
printer.Fprint(&buf, fset, funcAST.Body)

s := buf.String()
s = s[1 : len(s)-1]
s = strings.TrimSpace(strings.ReplaceAll(s, "\n\t", "\n"))

fmt.Println(s)
```

## 类型

### type CommentedNode

```go
type CommentedNode struct {
	Node     any // *ast.File, or ast.Expr, ast.Decl, ast.Spec, or ast.Stmt
	Comments []*ast.CommentGroup
}
```

A CommentedNode bundles an AST node and corresponding comments. It may be provided as argument to any of the Fprint functions.

​	CommentedNode 捆绑了一个 AST 节点和相应的注释。它可以作为任何 Fprint 函数的参数提供。

### type Config

```go
type Config struct {
	Mode     Mode // default: 0
	Tabwidth int  // default: 8
	Indent   int  // default: 0 (all code is indented at least by this much)
}
```

A Config node controls the output of Fprint.

​	Config 节点控制 Fprint 的输出。

#### (*Config) Fprint

```go
func (cfg *Config) Fprint(output io.Writer, fset *token.FileSet, node any) error
```

Fprint “pretty-prints” an AST node to output for a given configuration cfg. Position information is interpreted relative to the file set fset. The node type must be *ast.File, *CommentedNode, []ast.Decl, []ast.Stmt, or assignment-compatible to ast.Expr, ast.Decl, ast.Spec, or ast.Stmt.

​	Fprint “漂亮地打印”一个 AST 节点，以针对给定配置 cfg 输出。位置信息相对于文件集 fset 解释。节点类型必须是 *ast.File、*CommentedNode、[]ast.Decl、[]ast.Stmt 或与 ast.Expr、ast.Decl、ast.Spec 或 ast.Stmt 兼容的赋值。

### type Mode

```go
type Mode uint
```

A Mode value is a set of flags (or 0). They control printing.

​	Mode 值是一组标志（或 0）。它们控制打印。

``` go 
const (
	RawFormat Mode = 1 << iota // do not use a tabwriter; if set, UseSpaces is ignored
	TabIndent                  // use tabs for indentation independent of UseSpaces
	UseSpaces                  // use spaces instead of tabs for alignment
	SourcePos                  // emit //line directives to preserve original source positions
)
```