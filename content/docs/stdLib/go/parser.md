+++
title = "parser"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# parser

https://pkg.go.dev/go/parser@go1.20.1



Package parser implements a parser for Go source files. Input may be provided in a variety of forms (see the various Parse* functions); the output is an abstract syntax tree (AST) representing the Go source. The parser is invoked through one of the Parse* functions.

The parser accepts a larger language than is syntactically permitted by the Go spec, for simplicity, and for improved robustness in the presence of syntax errors. For instance, in method declarations, the receiver is treated like an ordinary parameter list and thus may contain multiple entries where the spec permits exactly one. Consequently, the corresponding field in the AST (ast.FuncDecl.Recv) field is not restricted to one entry.







## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func ParseDir 

``` go 
func ParseDir(fset *token.FileSet, path string, filter func(fs.FileInfo) bool, mode Mode) (pkgs map[string]*ast.Package, first error)
```

ParseDir calls ParseFile for all files with names ending in ".go" in the directory specified by path and returns a map of package name -> package AST with all the packages found.

If filter != nil, only the files with fs.FileInfo entries passing through the filter (and ending in ".go") are considered. The mode bits are passed to ParseFile unchanged. Position information is recorded in fset, which must not be nil.

If the directory couldn't be read, a nil map and the respective error are returned. If a parse error occurred, a non-nil but incomplete map and the first error encountered are returned.

#### func ParseExpr 

``` go 
func ParseExpr(x string) (ast.Expr, error)
```

ParseExpr is a convenience function for obtaining the AST of an expression x. The position information recorded in the AST is undefined. The filename used in error messages is the empty string.

If syntax errors were found, the result is a partial AST (with ast.Bad* nodes representing the fragments of erroneous source code). Multiple errors are returned via a scanner.ErrorList which is sorted by source position.

#### func ParseExprFrom  <- go1.5

``` go 
func ParseExprFrom(fset *token.FileSet, filename string, src any, mode Mode) (expr ast.Expr, err error)
```

ParseExprFrom is a convenience function for parsing an expression. The arguments have the same meaning as for ParseFile, but the source must be a valid Go (type or value) expression. Specifically, fset must not be nil.

If the source couldn't be read, the returned AST is nil and the error indicates the specific failure. If the source was read but syntax errors were found, the result is a partial AST (with ast.Bad* nodes representing the fragments of erroneous source code). Multiple errors are returned via a scanner.ErrorList which is sorted by source position.

#### func ParseFile 

``` go 
func ParseFile(fset *token.FileSet, filename string, src any, mode Mode) (f *ast.File, err error)
```

ParseFile parses the source code of a single Go source file and returns the corresponding ast.File node. The source code may be provided via the filename of the source file, or via the src parameter.

If src != nil, ParseFile parses the source from src and the filename is only used when recording position information. The type of the argument for the src parameter must be string, []byte, or io.Reader. If src == nil, ParseFile parses the file specified by filename.

The mode parameter controls the amount of source text parsed and other optional parser functionality. If the SkipObjectResolution mode bit is set, the object resolution phase of parsing will be skipped, causing File.Scope, File.Unresolved, and all Ident.Obj fields to be nil.

Position information is recorded in the file set fset, which must not be nil.

If the source couldn't be read, the returned AST is nil and the error indicates the specific failure. If the source was read but syntax errors were found, the result is a partial AST (with ast.Bad* nodes representing the fragments of erroneous source code). Multiple errors are returned via a scanner.ErrorList which is sorted by source position.

##### Example
``` go 
```

## 类型

### type Mode 

``` go 
type Mode uint
```

A Mode value is a set of flags (or 0). They control the amount of source code parsed and other optional parser functionality.

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