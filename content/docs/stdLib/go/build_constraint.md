+++
title = "contraint"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# constraint

https://pkg.go.dev/go/build/constraint@go1.20.1



Package constraint implements parsing and evaluation of build constraint lines. See https://golang.org/cmd/go/#hdr-Build_constraints for documentation about build constraints themselves.

This package parses both the original "// +build" syntax and the "//go:build" syntax that was added in Go 1.17. See https://golang.org/design/draft-gobuild for details about the "//go:build" syntax.



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [IsGoBuild](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=161) 

``` go linenums="1"
func IsGoBuild(line string) bool
```

IsGoBuild reports whether the line of text is a "//go:build" constraint. It only checks the prefix of the text, not that the expression itself parses.

#### func [IsPlusBuild](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=350) 

``` go linenums="1"
func IsPlusBuild(line string) bool
```

IsPlusBuild reports whether the line of text is a "// +build" constraint. It only checks the prefix of the text, not that the expression itself parses.

#### func [PlusBuildLines](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=451) 

``` go linenums="1"
func PlusBuildLines(x Expr) ([]string, error)
```

PlusBuildLines returns a sequence of "// +build" lines that evaluate to the build expression x. If the expression is too complex to convert directly to "// +build" lines, PlusBuildLines returns an error.

## 类型

### type [AndExpr](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=76) 

``` go linenums="1"
type AndExpr struct {
	X, Y Expr
}
```

An AndExpr represents the expression X && Y.

#### (*AndExpr) [Eval](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=82) 

``` go linenums="1"
func (x *AndExpr) Eval(ok func(tag string) bool) bool
```

#### (*AndExpr) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=89) 

``` go linenums="1"
func (x *AndExpr) String() string
```

### type [Expr](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=21) 

``` go linenums="1"
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

#### func [Parse](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=149) 

``` go linenums="1"
func Parse(line string) (Expr, error)
```

Parse parses a single build constraint line of the form "//go:build ..." or "// +build ..." and returns the corresponding boolean expression.

### type [NotExpr](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=54) 

``` go linenums="1"
type NotExpr struct {
	X Expr
}
```

A NotExpr represents the expression !X (the negation of X).

#### (*NotExpr) [Eval](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=60) 

``` go linenums="1"
func (x *NotExpr) Eval(ok func(tag string) bool) bool
```

#### (*NotExpr) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=64) 

``` go linenums="1"
func (x *NotExpr) String() string
```

### type [OrExpr](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=106) 

``` go linenums="1"
type OrExpr struct {
	X, Y Expr
}
```

An OrExpr represents the expression X || Y.

#### (*OrExpr) [Eval](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=112) 

``` go linenums="1"
func (x *OrExpr) Eval(ok func(tag string) bool) bool
```

#### (*OrExpr) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=119) 

``` go linenums="1"
func (x *OrExpr) String() string
```

### type [SyntaxError](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=136) 

``` go linenums="1"
type SyntaxError struct {
	Offset int    // byte offset in input where error was detected
	Err    string // description of error
}
```

A SyntaxError reports a syntax error in a parsed build expression.

#### (*SyntaxError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=141) 

``` go linenums="1"
func (e *SyntaxError) Error() string
```

### type [TagExpr](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=37) 

``` go linenums="1"
type TagExpr struct {
	Tag string // for example, "linux" or "cgo"
}
```

A TagExpr is an Expr for the single tag Tag.

#### (*TagExpr) [Eval](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=43) 

``` go linenums="1"
func (x *TagExpr) Eval(ok func(tag string) bool) bool
```

#### (*TagExpr) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/constraint/expr.go;l=47) 

``` go linenums="1"
func (x *TagExpr) String() string
```