+++
title = "format"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# format

https://pkg.go.dev/go/format@go1.20.1



Package format implements standard formatting of Go source.

Note that formatting of Go source code changes over time, so tools relying on consistent formatting should execute a specific version of the gofmt binary instead of using this package. That way, the formatting will be stable, and the tools won't need to be recompiled each time gofmt changes.

For example, pre-submit checks that use this package directly would behave differently depending on what Go version each developer uses, causing the check to be inherently fragile.







## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func Node 

``` go 
func Node(dst io.Writer, fset *token.FileSet, node any) error
```

Node formats node in canonical gofmt style and writes the result to dst.

The node type must be *ast.File, *printer.CommentedNode, []ast.Decl, []ast.Stmt, or assignment-compatible to ast.Expr, ast.Decl, ast.Spec, or ast.Stmt. Node does not modify node. Imports are not sorted for nodes representing partial source files (for instance, if the node is not an *ast.File or a *printer.CommentedNode not wrapping an *ast.File).

The function may return early (before the entire result is written) and return a formatting error, for instance due to an incorrect AST.

##### Example
``` go 
```

#### func Source 

``` go 
func Source(src []byte) ([]byte, error)
```

Source formats src in canonical gofmt style and returns the result or an (I/O or syntax) error. src is expected to be a syntactically correct Go source file, or a list of Go declarations or statements.

If src is a partial source file, the leading and trailing space of src is applied to the result (such that it has the same leading and trailing space as src), and the result is indented by the same amount as the first line of src containing code. Imports are not sorted for partial source files.

## 类型

This section is empty.