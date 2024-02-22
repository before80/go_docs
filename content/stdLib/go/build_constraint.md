+++
title = "build/contraint"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/go/build/constraint@go1.21.3](https://pkg.go.dev/go/build/constraint@go1.21.3)

Package constraint implements parsing and evaluation of build constraint lines. See https://golang.org/cmd/go/#hdr-Build_constraints for documentation about build constraints themselves.

​	Package constraint 实现构建约束行的解析和评估。有关构建约束本身的文档，请参阅 https://golang.org/cmd/go/#hdr-Build_constraints。

This package parses both the original “// +build” syntax and the “//go:build” syntax that was added in Go 1.17. See https://golang.org/design/draft-gobuild for details about the “//go:build” syntax.

​	此软件包同时解析原始的 “// +build” 语法和 Go 1.17 中添加的 “//go:build” 语法。有关 “//go:build” 语法的详细信息，请参阅 https://golang.org/design/draft-gobuild。



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func GoVersion <- go1.21.0

```go
func GoVersion(x Expr) string
```

GoVersion returns the minimum Go version implied by a given build expression. If the expression can be satisfied without any Go version tags, GoVersion returns an empty string.

​	GoVersion 返回给定构建表达式暗示的最低 Go 版本。如果表达式可以在没有任何 Go 版本标记的情况下得到满足，则 GoVersion 返回一个空字符串。

For example:

​	例如：

```go
GoVersion(linux && go1.22) = "go1.22"
GoVersion((linux && go1.22) || (windows && go1.20)) = "go1.20" => go1.20
GoVersion(linux) = ""
GoVersion(linux || (windows && go1.22)) = ""
GoVersion(!go1.22) = ""
```

GoVersion assumes that any tag or negated tag may independently be true, so that its analysis can be purely structural, without SAT solving. “Impossible” subexpressions may therefore affect the result.

​	GoVersion 假设任何标记或否定标记都可能独立为真，以便其分析可以是纯粹的结构性，而无需 SAT 求解。“不可能”的子表达式因此可能会影响结果。

For example:

​	例如：

```go
GoVersion((linux && !linux && go1.20) || go1.21) = "go1.20"
```

### func IsGoBuild

```go
func IsGoBuild(line string) bool
```

IsGoBuild reports whether the line of text is a “//go:build” constraint. It only checks the prefix of the text, not that the expression itself parses.

​	IsGoBuild 报告文本行是否为 “//go:build” 约束。它只检查文本的前缀，而不是表达式本身的解析。

### func IsPlusBuild

```go
func IsPlusBuild(line string) bool
```

IsPlusBuild reports whether the line of text is a “// +build” constraint. It only checks the prefix of the text, not that the expression itself parses.

​	IsPlusBuild 报告文本行是否为 “// +build” 约束。它只检查文本的前缀，而不是表达式本身的解析。

### func PlusBuildLines

```go
func PlusBuildLines(x Expr) ([]string, error)
```

PlusBuildLines returns a sequence of “// +build” lines that evaluate to the build expression x. If the expression is too complex to convert directly to “// +build” lines, PlusBuildLines returns an error.

​	PlusBuildLines 返回一个 “// +build” 行序列，该序列评估为构建表达式 x。如果表达式过于复杂，无法直接转换为 “// +build” 行，则 PlusBuildLines 返回一个错误。

## 类型

### type AndExpr

```go
type AndExpr struct {
	X, Y Expr
}
```

An AndExpr represents the expression X && Y.

​	AndExpr 表示表达式 X && Y。

#### (*AndExpr) Eval

```go
func (x *AndExpr) Eval(ok func(tag string) bool) bool
```

#### (*AndExpr) String

```go
func (x *AndExpr) String() string
```

### type Expr

```go
type Expr interface {
	// String returns the string form of the expression,
	// using the boolean syntax used in //go:build lines.
	String() string

	// Eval reports whether the expression evaluates to true.
	// It calls ok(tag) as needed to find out whether a given build tag
	// is satisfied by the current build configuration.
	Eval(ok func(tag string) bool) bool
	// contains filtered or unexported methods
}
```

An Expr is a build tag constraint expression. The underlying concrete type is *AndExpr, *OrExpr, *NotExpr, or *TagExpr.

​	Expr 是一个构建标记约束表达式。底层具体类型是 *AndExpr、*OrExpr、*NotExpr 或 *TagExpr。

#### func Parse

```go
func Parse(line string) (Expr, error)
```

Parse parses a single build constraint line of the form “//go:build …” or “// +build …” and returns the corresponding boolean expression.

​	Parse 解析形式为 “//go:build …” 或 “// +build …” 的单个构建约束行，并返回相应的布尔表达式。

### type NotExpr

```go
type NotExpr struct {
	X Expr
}
```

A NotExpr represents the expression !X (the negation of X).

​	NotExpr 表示表达式 !X（X 的否定）。

#### (*NotExpr) Eval

```go
func (x *NotExpr) Eval(ok func(tag string) bool) bool
```

#### (*NotExpr) String

```go
func (x *NotExpr) String() string
```

### type OrExpr

```go
type OrExpr struct {
	X, Y Expr
}
```

An OrExpr represents the expression X || Y.

​	OrExpr 表示表达式 X || Y。

#### (*OrExpr) Eval

```go
func (x *OrExpr) Eval(ok func(tag string) bool) bool
```

#### (*OrExpr) String

```go
func (x *OrExpr) String() string
```

### type SyntaxError

```go
type SyntaxError struct {
	Offset int    // byte offset in input where error was detected
	Err    string // description of error
}
```

A SyntaxError reports a syntax error in a parsed build expression.

​	SyntaxError 报告解析的构建表达式中的语法错误。

#### (*SyntaxError) Error

```go
func (e *SyntaxError) Error() string
```

### type TagExpr

```go
type TagExpr struct {
	Tag string // for example, "linux" or "cgo"
}
```

A TagExpr is an Expr for the single tag Tag.

​	TagExpr 是用于单个标记 Tag 的 Expr。

#### (*TagExpr) Eval 

``` go 
func (x *TagExpr) Eval(ok func(tag string) bool) bool
```

#### (*TagExpr) String 

``` go 
func (x *TagExpr) String() string
```