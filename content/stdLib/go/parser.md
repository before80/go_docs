+++
title = "parser"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/parser@go1.21.3](https://pkg.go.dev/go/parser@go1.21.3)

Package parser implements a parser for Go source files. Input may be provided in a variety of forms (see the various `Parse*` functions); the output is an abstract syntax tree (AST) representing the Go source. The parser is invoked through one of the Parse* functions.

​	parser 包为 Go 源文件实现了一个解析器。输入可以采用多种形式提供（请参阅各种 `Parse*` 函数）；输出是一个表示 Go 源的抽象语法树 (AST)。通过其中一个 Parse* 函数调用解析器。

The parser accepts a larger language than is syntactically permitted by the Go spec, for simplicity, and for improved robustness in the presence of syntax errors. For instance, in method declarations, the receiver is treated like an ordinary parameter list and thus may contain multiple entries where the spec permits exactly one. Consequently, the corresponding field in the AST (ast.FuncDecl.Recv) field is not restricted to one entry.

​	为了简单起见，并且为了提高在存在语法错误时具有更高的健壮性，解析器接受的语言比 Go 规范在语法上允许的语言更大。例如，在方法声明中，接收者被视为一个普通参数列表，因此可能包含多个条目，而规范只允许一个条目。因此，AST 中的相应字段（ast.FuncDecl.Recv）字段不限于一个条目。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func ParseDir

```go
func ParseDir(fset *token.FileSet, path string, filter func(fs.FileInfo) bool, mode Mode) (pkgs map[string]*ast.Package, first error)
```

ParseDir calls ParseFile for all files with names ending in “.go” in the directory specified by path and returns a map of package name -> package AST with all the packages found.

​	ParseDir 会对路径指定的目录中所有以 “.go” 结尾的文件调用 ParseFile，并返回一个包名 -> 包 AST 的映射，其中包含找到的所有包。

If filter != nil, only the files with fs.FileInfo entries passing through the filter (and ending in “.go”) are considered. The mode bits are passed to ParseFile unchanged. Position information is recorded in fset, which must not be nil.

​	如果 filter 不为 nil，则仅考虑通过该筛选器（并以 “.go” 结尾）的 fs.FileInfo 条目的文件。模式位会原样传递给 ParseFile。位置信息会记录在 fset 中，fset 不能为 nil。

If the directory couldn’t be read, a nil map and the respective error are returned. If a parse error occurred, a non-nil but incomplete map and the first error encountered are returned.

​	如果无法读取目录，则返回一个 nil 映射和相应的错误。如果发生解析错误，则返回一个非 nil 但不完整的映射和遇到的第一个错误。

### func ParseExpr

```go
func ParseExpr(x string) (ast.Expr, error)
```

ParseExpr is a convenience function for obtaining the AST of an expression x. The position information recorded in the AST is undefined. The filename used in error messages is the empty string.

​	ParseExpr 是一个获取表达式 x 的 AST 的便捷函数。记录在 AST 中的位置信息是未定义的。错误消息中使用的文件名是空字符串。

If syntax errors were found, the result is a partial AST (with ast.Bad* nodes representing the fragments of erroneous source code). Multiple errors are returned via a scanner.ErrorList which is sorted by source position.

​	如果找到语法错误，则结果是一个部分 AST（其中 ast.Bad* 节点表示错误源代码的片段）。多个错误通过按源位置排序的 scanner.ErrorList 返回。

### func ParseExprFrom <- go1.5

```go
func ParseExprFrom(fset *token.FileSet, filename string, src any, mode Mode) (expr ast.Expr, err error)
```

ParseExprFrom is a convenience function for parsing an expression. The arguments have the same meaning as for ParseFile, but the source must be a valid Go (type or value) expression. Specifically, fset must not be nil.

​	ParseExprFrom 是一个用于解析表达式的便捷函数。参数的含义与 ParseFile 相同，但源必须是有效的 Go（类型或值）表达式。具体来说，fset 不能为 nil。

If the source couldn’t be read, the returned AST is nil and the error indicates the specific failure. If the source was read but syntax errors were found, the result is a partial AST (with ast.Bad* nodes representing the fragments of erroneous source code). Multiple errors are returned via a scanner.ErrorList which is sorted by source position.

​	如果无法读取源，则返回的 AST 为 nil，并且错误指示具体故障。如果已读取源但发现语法错误，则结果是部分 AST（其中 ast.Bad* 节点表示错误源代码的片段）。多个错误通过 scanner.ErrorList 返回，该列表按源位置排序。

### func ParseFile

```go
func ParseFile(fset *token.FileSet, filename string, src any, mode Mode) (f *ast.File, err error)
```

ParseFile parses the source code of a single Go source file and returns the corresponding ast.File node. The source code may be provided via the filename of the source file, or via the src parameter.

​	ParseFile 解析单个 Go 源文件的源代码，并返回相应的 ast.File 节点。源代码可以通过源文件的 filename 提供，也可以通过 src 参数提供。

If src != nil, ParseFile parses the source from src and the filename is only used when recording position information. The type of the argument for the src parameter must be string, []byte, or io.Reader. If src == nil, ParseFile parses the file specified by filename.

​	如果 src 不为 nil，则 ParseFile 从 src 解析源，并且 filename 仅在记录位置信息时使用。src 参数的参数类型必须是 string、[]byte 或 io.Reader。如果 src 为 nil，则 ParseFile 解析由 filename 指定的文件。

The mode parameter controls the amount of source text parsed and other optional parser functionality. If the SkipObjectResolution mode bit is set, the object resolution phase of parsing will be skipped, causing File.Scope, File.Unresolved, and all Ident.Obj fields to be nil.

​	mode 参数控制解析的源文本量和其他可选的解析器功能。如果设置了 SkipObjectResolution 模式位，则将跳过解析的对象解析阶段，导致 File.Scope、File.Unresolved 和所有 Ident.Obj 字段变为 nil。

Position information is recorded in the file set fset, which must not be nil.

​	位置信息记录在文件集 fset 中，它不能为 nil。

If the source couldn’t be read, the returned AST is nil and the error indicates the specific failure. If the source was read but syntax errors were found, the result is a partial AST (with ast.Bad* nodes representing the fragments of erroneous source code). Multiple errors are returned via a scanner.ErrorList which is sorted by source position.

​	如果无法读取源代码，则返回的 AST 为 nil，并且错误指示具体故障。如果读取了源代码但发现了语法错误，则结果是一个部分 AST（其中 ast.Bad* 节点表示错误源代码的片段）。多个错误通过按源位置排序的 scanner.ErrorList 返回。

#### ParseFile Example

```go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	src := `package foo

import (
	"fmt"
	"time"
)

func bar() {
	fmt.Println(time.Now())
}`

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}

}
Output:


"fmt"
"time"
```

## 类型

### type Mode

```go
type Mode uint
```

A Mode value is a set of flags (or 0). They control the amount of source code parsed and other optional parser functionality.

​	Mode 值是一组标志（或 0）。它们控制解析的源代码量和其他可选的解析器功能。

``` go 
const (
	PackageClauseOnly    Mode             = 1 << iota // stop parsing after package clause
	ImportsOnly                                       // stop parsing after import declarations
	ParseComments                                     // parse comments and add them to AST
	Trace                                             // print a trace of parsed productions
	DeclarationErrors                                 // report declaration errors
	SpuriousErrors                                    // same as AllErrors, for backward-compatibility
	SkipObjectResolution                              // don't resolve identifiers to objects - see ParseFile
	AllErrors            = SpuriousErrors             // report all errors (not just the first 10 on different lines)
)
```