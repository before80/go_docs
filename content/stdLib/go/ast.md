+++
title = "ast"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/ast@go1.23.0](https://pkg.go.dev/go/ast@go1.23.0)

Package ast declares the types used to represent syntax trees for Go packages.

​	ast 包声明用于表示 Go 包的语法树的类型。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func FileExports

```go
func FileExports(src *File) bool
```

FileExports trims the AST for a Go source file in place such that only exported nodes remain: all top-level identifiers which are not exported and their associated information (such as type, initial value, or function body) are removed. Non-exported fields and methods of exported types are stripped. The File.Comments list is not changed.

​	FileExports 就地修剪 Go 源文件中的 AST，以便仅保留导出的节点：所有未导出的顶级标识符及其关联信息（例如类型、初始值或函数体）都会被移除。导出的类型的非导出字段和方法也会被剥离。File.Comments 列表不会发生变化。

FileExports reports whether there are exported declarations.

​	FileExports 报告是否存在导出的声明。

### func FilterDecl

```go
func FilterDecl(decl Decl, f Filter) bool
```

FilterDecl trims the AST for a Go declaration in place by removing all names (including struct field and interface method names, but not from parameter lists) that don’t pass through the filter f.

​	FilterDecl 就地修剪 Go 声明的 AST，方法是移除所有未通过过滤器 f 的名称（包括结构字段和接口方法名称，但不包括参数列表）。

FilterDecl reports whether there are any declared names left after filtering.

​	FilterDecl 报告过滤后是否还有任何已声明的名称。

### func FilterFile

```go
func FilterFile(src *File, f Filter) bool
```

FilterFile trims the AST for a Go file in place by removing all names from top-level declarations (including struct field and interface method names, but not from parameter lists) that don’t pass through the filter f. If the declaration is empty afterwards, the declaration is removed from the AST. Import declarations are always removed. The File.Comments list is not changed.

​	FilterFile 就地修剪 Go 文件的 AST，方法是移除所有未通过过滤器 f 的顶级声明中的名称（包括结构字段和接口方法名称，但不包括参数列表）。如果声明之后为空，则从 AST 中移除该声明。始终移除导入声明。File.Comments 列表不会发生变化。

FilterFile reports whether there are any top-level declarations left after filtering.

​	FilterFile 报告过滤后是否还有任何顶级声明。

### func FilterPackage

```go
func FilterPackage(pkg *Package, f Filter) bool
```

FilterPackage trims the AST for a Go package in place by removing all names from top-level declarations (including struct field and interface method names, but not from parameter lists) that don’t pass through the filter f. If the declaration is empty afterwards, the declaration is removed from the AST. The pkg.Files list is not changed, so that file names and top-level package comments don’t get lost.

​	FilterPackage 就地修剪 Go 包的 AST，方法是删除所有顶级声明中的名称（包括结构字段和接口方法名称，但不包括参数列表），这些名称未通过过滤器 f。如果声明之后为空，则从 AST 中删除该声明。 pkg.Files 列表不会更改，因此不会丢失文件名和顶级包注释。

FilterPackage reports whether there are any top-level declarations left after filtering.

​	FilterPackage 报告过滤后是否还有任何顶级声明。

### func Fprint

```go
func Fprint(w io.Writer, fset *token.FileSet, x any, f FieldFilter) error
```

Fprint prints the (sub-)tree starting at AST node x to w. If fset != nil, position information is interpreted relative to that file set. Otherwise positions are printed as integer values (file set specific offsets).

​	Fprint 将从 AST 节点 x 开始的（子）树打印到 w。如果 fset 不为 nil，则位置信息将相对于该文件集进行解释。否则，位置将以整数值（特定于文件集的偏移量）打印。

A non-nil FieldFilter f may be provided to control the output: struct fields for which f(fieldname, fieldvalue) is true are printed; all others are filtered from the output. Unexported struct fields are never printed.

​	可以提供一个非 nil 的 FieldFilter f 来控制输出：将打印 f(fieldname, fieldvalue) 为 true 的结构字段；所有其他字段都将从输出中过滤掉。从不打印未导出的结构字段。

### func Inspect

```go
func Inspect(node Node, f func(Node) bool)
```

Inspect traverses an AST in depth-first order: It starts by calling f(node); node must not be nil. If f returns true, Inspect invokes f recursively for each of the non-nil children of node, followed by a call of f(nil).

​	Inspect 以深度优先顺序遍历 AST：它首先调用 f(node)；node 不能为 nil。如果 f 返回 true，则 Inspect 会为 node 的每个非 nil 子项递归调用 f，然后调用 f(nil)。

#### Inspect Example 

This example demonstrates how to inspect the AST of a Go program.

​	此示例演示如何检查 Go 程序的 AST。

```go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to inspect the AST.
	src := `
package p
const c = 1.0
var X = f(3.14)*2 + c
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
		case *ast.Ident:
			s = x.Name
		}
		if s != "" {
			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})

}
Output:

src.go:2:9:	p
src.go:3:7:	c
src.go:3:11:	1.0
src.go:4:5:	X
src.go:4:9:	f
src.go:4:11:	3.14
src.go:4:17:	2
src.go:4:21:	c
```

### func IsExported

```go
func IsExported(name string) bool
```

IsExported reports whether name starts with an upper-case letter.

​	IsExported 报告名称是否以大写字母开头。

### func IsGenerated <-go1.21.0

``` go
func IsGenerated(file *File) bool
```

IsGenerated reports whether the file was generated by a program, not handwritten, by detecting the special comment described at https://go.dev/s/generatedcode.

​	IsGenerated 报告文件是否由程序生成，而不是手写，方法是检测 https://go.dev/s/generatedcode 中描述的特殊注释。

The syntax tree must have been parsed with the ParseComments flag. Example:

​	必须使用 ParseComments 标志解析语法树。示例：

```
f, err := parser.ParseFile(fset, filename, src, parser.ParseComments|parser.PackageClauseOnly)
if err != nil { ... }
gen := ast.IsGenerated(f)
```

### func NotNilFilter

```go
func NotNilFilter(_ string, v reflect.Value) bool
```

NotNilFilter returns true for field values that are not nil; it returns false otherwise.

​	NotNilFilter 对非 nil 的字段值返回 true；否则返回 false。

### func PackageExports

```go
func PackageExports(pkg *Package) bool
```

PackageExports trims the AST for a Go package in place such that only exported nodes remain. The pkg.Files list is not changed, so that file names and top-level package comments don’t get lost.

​	PackageExports 就地修剪 Go 包的 AST，以便仅保留导出的节点。pkg.Files 列表不会更改，因此不会丢失文件名和顶级包注释。

PackageExports reports whether there are exported declarations; it returns false otherwise.

​	PackageExports 报告是否存在导出的声明；否则返回 false。

### func Print

```go
func Print(fset *token.FileSet, x any) error
```

Print prints x to standard output, skipping nil fields. Print(fset, x) is the same as Fprint(os.Stdout, fset, x, NotNilFilter).

​	Print 将 x 打印到标准输出，跳过 nil 字段。Print(fset, x) 与 Fprint(os.Stdout, fset, x, NotNilFilter) 相同。

#### Print Example 

This example shows what an AST looks like when printed for debugging.

​	此示例显示了在调试时打印 AST 的样子。

```go
package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to print the AST.
	src := `
package main
func main() {
	println("Hello, World!")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)

}
Output:

     0  *ast.File {
     1  .  Package: 2:1
     2  .  Name: *ast.Ident {
     3  .  .  NamePos: 2:9
     4  .  .  Name: "main"
     5  .  }
     6  .  Decls: []ast.Decl (len = 1) {
     7  .  .  0: *ast.FuncDecl {
     8  .  .  .  Name: *ast.Ident {
     9  .  .  .  .  NamePos: 3:6
    10  .  .  .  .  Name: "main"
    11  .  .  .  .  Obj: *ast.Object {
    12  .  .  .  .  .  Kind: func
    13  .  .  .  .  .  Name: "main"
    14  .  .  .  .  .  Decl: *(obj @ 7)
    15  .  .  .  .  }
    16  .  .  .  }
    17  .  .  .  Type: *ast.FuncType {
    18  .  .  .  .  Func: 3:1
    19  .  .  .  .  Params: *ast.FieldList {
    20  .  .  .  .  .  Opening: 3:10
    21  .  .  .  .  .  Closing: 3:11
    22  .  .  .  .  }
    23  .  .  .  }
    24  .  .  .  Body: *ast.BlockStmt {
    25  .  .  .  .  Lbrace: 3:13
    26  .  .  .  .  List: []ast.Stmt (len = 1) {
    27  .  .  .  .  .  0: *ast.ExprStmt {
    28  .  .  .  .  .  .  X: *ast.CallExpr {
    29  .  .  .  .  .  .  .  Fun: *ast.Ident {
    30  .  .  .  .  .  .  .  .  NamePos: 4:2
    31  .  .  .  .  .  .  .  .  Name: "println"
    32  .  .  .  .  .  .  .  }
    33  .  .  .  .  .  .  .  Lparen: 4:9
    34  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
    35  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
    36  .  .  .  .  .  .  .  .  .  ValuePos: 4:10
    37  .  .  .  .  .  .  .  .  .  Kind: STRING
    38  .  .  .  .  .  .  .  .  .  Value: "\"Hello, World!\""
    39  .  .  .  .  .  .  .  .  }
    40  .  .  .  .  .  .  .  }
    41  .  .  .  .  .  .  .  Ellipsis: -
    42  .  .  .  .  .  .  .  Rparen: 4:25
    43  .  .  .  .  .  .  }
    44  .  .  .  .  .  }
    45  .  .  .  .  }
    46  .  .  .  .  Rbrace: 5:1
    47  .  .  .  }
    48  .  .  }
    49  .  }
    50  .  FileStart: 1:1
    51  .  FileEnd: 5:3
    52  .  Scope: *ast.Scope {
    53  .  .  Objects: map[string]*ast.Object (len = 1) {
    54  .  .  .  "main": *(obj @ 11)
    55  .  .  }
    56  .  }
    57  .  Unresolved: []*ast.Ident (len = 1) {
    58  .  .  0: *(obj @ 29)
    59  .  }
    60  .  GoVersion: ""
    61  }
```

### func SortImports

```go
func SortImports(fset *token.FileSet, f *File)
```

SortImports sorts runs of consecutive import lines in import blocks in f. It also removes duplicate imports when it is possible to do so without data loss.

​	SortImports 对 f 中导入块中连续导入行的运行进行排序。当有可能在不丢失数据的情况下执行此操作时，它还会删除重复的导入。

### func Walk

```go
func Walk(v Visitor, node Node)
```

Walk traverses an AST in depth-first order: It starts by calling v.Visit(node); node must not be nil. If the visitor w returned by v.Visit(node) is not nil, Walk is invoked recursively with visitor w for each of the non-nil children of node, followed by a call of w.Visit(nil).

​	Walk 以深度优先顺序遍历 AST：它首先调用 v.Visit(node)；node 不能为 nil。如果 v.Visit(node) 返回的访问者 w 不为 nil，则对 node 的每个非 nil 子级递归调用访问者 w，然后调用 w.Visit(nil)。

## 类型

### type ArrayType

```go
type ArrayType struct {
	Lbrack token.Pos // position of "["
	Len    Expr      // Ellipsis node for [...]T array types, nil for slice types
	Elt    Expr      // element type
}
```

An ArrayType node represents an array or slice type.

​	ArrayType 节点表示数组或切片类型。

#### (*ArrayType) End

```go
func (x *ArrayType) End() token.Pos
```

#### (*ArrayType) Pos

```go
func (x *ArrayType) Pos() token.Pos
```

### type AssignStmt

```go
type AssignStmt struct {
	Lhs    []Expr
	TokPos token.Pos   // position of Tok
	Tok    token.Token // assignment token, DEFINE
	Rhs    []Expr
}
```

An AssignStmt node represents an assignment or a short variable declaration.

​	AssignStmt 节点表示赋值或短变量声明。

#### (*AssignStmt) End 

```go
func (s *AssignStmt) End() token.Pos
```

#### (*AssignStmt) Pos 

```go
func (s *AssignStmt) Pos() token.Pos
```

### type BadDecl

```go
type BadDecl struct {
	From, To token.Pos // position range of bad declaration
}
```

A BadDecl node is a placeholder for a declaration containing syntax errors for which a correct declaration node cannot be created.

​	BadDecl 节点是包含语法错误的声明的占位符，无法为其创建正确的声明节点。

#### (*BadDecl) End 

```go
func (d *BadDecl) End() token.Pos
```

#### (*BadDecl) Pos 

```go
func (d *BadDecl) Pos() token.Pos
```

### type BadExpr 

```go
type BadExpr struct {
	From, To token.Pos // position range of bad expression
}
```

A BadExpr node is a placeholder for an expression containing syntax errors for which a correct expression node cannot be created.

​	BadExpr 节点是包含语法错误的表达式的占位符，无法为其创建正确的表达式节点。

#### (*BadExpr) End 

```go
func (x *BadExpr) End() token.Pos
```

#### (*BadExpr) Pos

```go
func (x *BadExpr) Pos() token.Pos
```

### type BadStmt

```go
type BadStmt struct {
	From, To token.Pos // position range of bad statement
}
```

A BadStmt node is a placeholder for statements containing syntax errors for which no correct statement nodes can be created.

​	BadStmt 节点是包含语法错误的语句的占位符，无法为其创建正确的语句节点。

#### (*BadStmt) End

```go
func (s *BadStmt) End() token.Pos
```

#### (*BadStmt) Pos

```go
func (s *BadStmt) Pos() token.Pos
```

### type BasicLit

```go
type BasicLit struct {
	ValuePos token.Pos   // literal position
	Kind     token.Token // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
	Value    string      // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
}
```

A BasicLit node represents a literal of basic type.

​	BasicLit 节点表示基本类型的文字。

#### (*BasicLit) End

```go
func (x *BasicLit) End() token.Pos
```

#### (*BasicLit) Pos

```go
func (x *BasicLit) Pos() token.Pos
```

### type BinaryExpr

```go
type BinaryExpr struct {
	X     Expr        // left operand
	OpPos token.Pos   // position of Op
	Op    token.Token // operator
	Y     Expr        // right operand
}
```

A BinaryExpr node represents a binary expression.

​	BinaryExpr 节点表示二元表达式。

#### (*BinaryExpr) End

```go
func (x *BinaryExpr) End() token.Pos
```

#### (*BinaryExpr) Pos

```go
func (x *BinaryExpr) Pos() token.Pos
```

### type BlockStmt 

```go
type BlockStmt struct {
	Lbrace token.Pos // position of "{"
	List   []Stmt
	Rbrace token.Pos // position of "}", if any (may be absent due to syntax error)
}
```

A BlockStmt node represents a braced statement list.

​	A BlockStmt 节点表示一个带大括号的语句列表。

#### (*BlockStmt) End 

```go
func (s *BlockStmt) End() token.Pos
```

#### (*BlockStmt) Pos 

```go
func (s *BlockStmt) Pos() token.Pos
```

### type BranchStmt

```go
type BranchStmt struct {
	TokPos token.Pos   // position of Tok
	Tok    token.Token // keyword token (BREAK, CONTINUE, GOTO, FALLTHROUGH)
	Label  *Ident      // label name; or nil
}
```

A BranchStmt node represents a break, continue, goto, or fallthrough statement.

​	BranchStmt 节点表示 break、continue、goto 或 fallthrough 语句。

#### (*BranchStmt) End 

```go
func (s *BranchStmt) End() token.Pos
```

#### (*BranchStmt) Pos 

```go
func (s *BranchStmt) Pos() token.Pos
```

### type CallExpr

```go
type CallExpr struct {
	Fun      Expr      // function expression
	Lparen   token.Pos // position of "("
	Args     []Expr    // function arguments; or nil
	Ellipsis token.Pos // position of "..." (token.NoPos if there is no "...")
	Rparen   token.Pos // position of ")"
}
```

A CallExpr node represents an expression followed by an argument list.

​	CallExpr 节点表示一个表达式，后跟一个参数列表。

#### (*CallExpr) End

```go
func (x *CallExpr) End() token.Pos
```

#### (*CallExpr) Pos 

```go
func (x *CallExpr) Pos() token.Pos
```

### type CaseClause 

```go
type CaseClause struct {
	Case  token.Pos // position of "case" or "default" keyword
	List  []Expr    // list of expressions or types; nil means default case
	Colon token.Pos // position of ":"
	Body  []Stmt    // statement list; or nil
}
```

A CaseClause represents a case of an expression or type switch statement.

​	CaseClause 表示表达式或类型 switch 语句的一个 case。

#### (*CaseClause) End 

```go
func (s *CaseClause) End() token.Pos
```

#### (*CaseClause) Pos

```go
func (s *CaseClause) Pos() token.Pos
```

### type ChanDir

```go
type ChanDir int
```

The direction of a channel type is indicated by a bit mask including one or both of the following constants.

​	通道类型的方向由一个位掩码指示，其中包括以下一个或两个常量。

```go
const (
	SEND ChanDir = 1 << iota
	RECV
)
```

### type ChanType 

```go
type ChanType struct {
	Begin token.Pos // position of "chan" keyword or "<-" (whichever comes first)
	Arrow token.Pos // position of "<-" (token.NoPos if there is no "<-")
	Dir   ChanDir   // channel direction
	Value Expr      // value type
}
```

A ChanType node represents a channel type.

​	ChanType 节点表示通道类型。

#### (*ChanType) End 

```go
func (x *ChanType) End() token.Pos
```

#### (*ChanType) Pos 

```go
func (x *ChanType) Pos() token.Pos
```

### type CommClause

```go
type CommClause struct {
	Case  token.Pos // position of "case" or "default" keyword
	Comm  Stmt      // send or receive statement; nil means default case
	Colon token.Pos // position of ":"
	Body  []Stmt    // statement list; or nil
}
```

A CommClause node represents a case of a select statement.

​	CommClause 节点表示 select 语句的一个 case。

#### (*CommClause) End 

```go
func (s *CommClause) End() token.Pos
```

#### (*CommClause) Pos 

```go
func (s *CommClause) Pos() token.Pos
```

### type Comment 

```go
type Comment struct {
	Slash token.Pos // position of "/" starting the comment
	Text  string    // comment text (excluding '\n' for //-style comments)
}
```

A Comment node represents a single //-style or /*-style comment.

​	Comment 节点表示单个 // 样式或 /* 样式注释。

The Text field contains the comment text without carriage returns (\r) that may have been present in the source. Because a comment’s end position is computed using len(Text), the position reported by End() does not match the true source end position for comments containing carriage returns.

​	Text 字段包含注释文本，不包含源代码中可能存在的回车符 (\r)。由于注释的结束位置是使用 len(Text) 计算的，因此 End() 报告的位置与包含回车符的注释的真实源结束位置不匹配。

#### (*Comment) End （*Comment）End

```go
func (c *Comment) End() token.Pos
```

#### (*Comment) Pos

```go
func (c *Comment) Pos() token.Pos
```

### type CommentGroup

```go
type CommentGroup struct {
	List []*Comment // len(List) > 0
}
```

A CommentGroup represents a sequence of comments with no other tokens and no empty lines between.

​	CommentGroup 表示一系列注释，其间没有其他标记且没有空行。

#### (*CommentGroup) End

```go
func (g *CommentGroup) End() token.Pos
```

#### (*CommentGroup) Pos

```go
func (g *CommentGroup) Pos() token.Pos
```

#### (*CommentGroup) Text

```go
func (g *CommentGroup) Text() string
```

Text returns the text of the comment. Comment markers (//, /*, and */), the first space of a line comment, and leading and trailing empty lines are removed. Comment directives like “//line” and “//go:noinline” are also removed. Multiple empty lines are reduced to one, and trailing space on lines is trimmed. Unless the result is empty, it is newline-terminated.

​	Text 返回注释的文本。注释标记（//、/* 和 */）、行注释的第一个空格以及前导和尾随空行均已删除。注释指令（如 “//line” 和 “//go:noinline”）也已删除。多个空行减少为一个，并且修剪了行上的尾随空格。除非结果为空，否则它以换行符结尾。

### type CommentMap <- go1.1

```go
type CommentMap map[Node][]*CommentGroup
```

A CommentMap maps an AST node to a list of comment groups associated with it. See NewCommentMap for a description of the association.

​	CommentMap 将 AST 节点映射到与之关联的注释组列表。有关关联的说明，请参阅 NewCommentMap。

##### Example

This example illustrates how to remove a variable declaration in a Go program while maintaining correct comment association using an ast.CommentMap.

​	这个例子说明了如何在Go程序中删除变量声明，同时使用ast.CommentMap保持正确的注释关联。

```go
package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	// src is the input for which we create the AST that we
	// are going to manipulate.
	src := `
// This is the package comment.
package main

// This comment is associated with the hello constant.
const hello = "Hello, World!" // line comment 1

// This comment is associated with the foo variable.
var foo = hello // line comment 2

// This comment is associated with the main function.
func main() {
	fmt.Println(hello) // line comment 3
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Create an ast.CommentMap from the ast.File's comments.
	// This helps keeping the association between comments
	// and AST nodes.
	cmap := ast.NewCommentMap(fset, f, f.Comments)

	// Remove the first variable declaration from the list of declarations.
	for i, decl := range f.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.VAR {
			copy(f.Decls[i:], f.Decls[i+1:])
			f.Decls = f.Decls[:len(f.Decls)-1]
			break
		}
	}

	// Use the comment map to filter comments that don't belong anymore
	// (the comments associated with the variable declaration), and create
	// the new comments list.
	f.Comments = cmap.Filter(f).Comments()

	// Print the modified AST.
	var buf strings.Builder
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}
	fmt.Printf("%s", buf.String())

}

Output:

// This is the package comment.
package main

// This comment is associated with the hello constant.
const hello = "Hello, World!" // line comment 1

// This comment is associated with the main function.
func main() {
	fmt.Println(hello) // line comment 3
}
```



#### func NewCommentMap <- go1.1

```go
func NewCommentMap(fset *token.FileSet, node Node, comments []*CommentGroup) CommentMap
```

NewCommentMap creates a new comment map by associating comment groups of the comments list with the nodes of the AST specified by node.

​	NewCommentMap 通过将注释列表的注释组与节点关联起来创建新的注释映射。

A comment group g is associated with a node n if:

​	注释组 g 与节点 n 相关联，如果：

- g starts on the same line as n ends
  g 与 n 结束在同一行开始
- g starts on the line immediately following n, and there is at least one empty line after g and before the next node
  g 在紧跟 n 的行开始，并且 g 和下一个节点之前至少有一行空行
- g starts before n and is not associated to the node before n via the previous rules
  g 在 n 之前开始，并且通过前面的规则未与 n 之前的节点相关联

NewCommentMap tries to associate a comment group to the “largest” node possible: For instance, if the comment is a line comment trailing an assignment, the comment is associated with the entire assignment rather than just the last operand in the assignment.

​	NewCommentMap 尝试将注释组与尽可能“最大”的节点相关联：例如，如果注释是尾随赋值的行注释，则注释与整个赋值相关联，而不仅仅是赋值中的最后一个操作数。

#### (CommentMap) Comments <- go1.1

```go
func (cmap CommentMap) Comments() []*CommentGroup
```

Comments returns the list of comment groups in the comment map. The result is sorted in source order.

​	Comments 返回注释映射中的注释组列表。结果按源顺序排序。

#### (CommentMap) Filter <- go1.1

```go
func (cmap CommentMap) Filter(node Node) CommentMap
```

Filter returns a new comment map consisting of only those entries of cmap for which a corresponding node exists in the AST specified by node.

​	Filter 返回一个新的注释映射，其中仅包含 cmap 的那些条目，其对应的节点存在于节点指定的 AST 中。

#### (CommentMap) String <- go1.1

```go
func (cmap CommentMap) String() string
```

#### (CommentMap) Update <- go1.1

```go
func (cmap CommentMap) Update(old, new Node) Node
```

Update replaces an old node in the comment map with the new node and returns the new node. Comments that were associated with the old node are associated with the new node.

​	Update 用新节点替换注释图中的旧节点，并返回新节点。与旧节点关联的注释与新节点关联。

### type CompositeLit

```go
type CompositeLit struct {
	Type       Expr      // literal type; or nil
	Lbrace     token.Pos // position of "{"
	Elts       []Expr    // list of composite elements; or nil
	Rbrace     token.Pos // position of "}"
	Incomplete bool      // true if (source) expressions are missing in the Elts list
}
```

A CompositeLit node represents a composite literal.

​	CompositeLit 节点表示复合字面量。

#### (*CompositeLit) End

```go
func (x *CompositeLit) End() token.Pos
```

#### (*CompositeLit) Pos

```go
func (x *CompositeLit) Pos() token.Pos
```

### type Decl

```go
type Decl interface {
	Node
	// contains filtered or unexported methods
}
```

All declaration nodes implement the Decl interface.

​	所有声明节点都实现了 Decl 接口。

### type DeclStmt

```go
type DeclStmt struct {
	Decl Decl // *GenDecl with CONST, TYPE, or VAR token
}
```

A DeclStmt node represents a declaration in a statement list.

​	DeclStmt 节点表示语句列表中的声明。

#### (*DeclStmt) End

```go
func (s *DeclStmt) End() token.Pos
```

#### (*DeclStmt) Pos

```go
func (s *DeclStmt) Pos() token.Pos
```

### type DeferStmt

```go
type DeferStmt struct {
	Defer token.Pos // position of "defer" keyword
	Call  *CallExpr
}
```

A DeferStmt node represents a defer statement.

​	DeferStmt 节点表示 defer 语句。

#### (*DeferStmt) End

```go
func (s *DeferStmt) End() token.Pos
```

#### (*DeferStmt) Pos

```go
func (s *DeferStmt) Pos() token.Pos
```

### type Ellipsis 

```go
type Ellipsis struct {
	Ellipsis token.Pos // position of "..."
	Elt      Expr      // ellipsis element type (parameter lists only); or nil
}
```

An Ellipsis node stands for the “…” type in a parameter list or the “…” length in an array type.

​	Ellipsis 节点表示参数列表中的“…”类型或数组类型中的“…”长度。

#### (*Ellipsis) End

```go
func (x *Ellipsis) End() token.Pos
```

#### (*Ellipsis) Pos 

```go
func (x *Ellipsis) Pos() token.Pos
```

### type EmptyStmt

```go
type EmptyStmt struct {
	Semicolon token.Pos // position of following ";"
	Implicit  bool      // if set, ";" was omitted in the source
}
```

An EmptyStmt node represents an empty statement. The “position” of the empty statement is the position of the immediately following (explicit or implicit) semicolon.

​	EmptyStmt 节点表示一个空语句。“位置”的空语句是紧跟其后（显式或隐式）分号的位置。

#### (*EmptyStmt) End

```go
func (s *EmptyStmt) End() token.Pos
```

#### (*EmptyStmt) Pos

```go
func (s *EmptyStmt) Pos() token.Pos
```

### type Expr

```go
type Expr interface {
	Node
	// contains filtered or unexported methods
}
```

All expression nodes implement the Expr interface.

​	所有表达式节点都实现了 Expr 接口。

### type ExprStmt

```go
type ExprStmt struct {
	X Expr // expression
}
```

An ExprStmt node represents a (stand-alone) expression in a statement list.

​	ExprStmt 节点表示语句列表中的（独立）表达式。

#### (*ExprStmt) End

```go
func (s *ExprStmt) End() token.Pos
```

#### (*ExprStmt) Pos

```go
func (s *ExprStmt) Pos() token.Pos
```

### type Field

```go
type Field struct {
	Doc     *CommentGroup // associated documentation; or nil
	Names   []*Ident      // field/method/(type) parameter names; or nil
	Type    Expr          // field/method/parameter type; or nil
	Tag     *BasicLit     // field tag; or nil
	Comment *CommentGroup // line comments; or nil
}
```

A Field represents a Field declaration list in a struct type, a method list in an interface type, or a parameter/result declaration in a signature. Field.Names is nil for unnamed parameters (parameter lists which only contain types) and embedded struct fields. In the latter case, the field name is the type name.

​	Field 表示结构类型中的 Field 声明列表、接口类型中的方法列表或签名中的参数/结果声明。对于未命名参数（仅包含类型的参数列表）和嵌入式结构字段，Field.Names 为 nil。在后一种情况下，字段名称是类型名称。

#### (*Field) End

```go
func (f *Field) End() token.Pos
```

#### (*Field) Pos

```go
func (f *Field) Pos() token.Pos
```

### type FieldFilter

```go
type FieldFilter func(name string, value reflect.Value) bool
```

A FieldFilter may be provided to Fprint to control the output.

​	可以向 Fprint 提供 FieldFilter 以控制输出。

### type FieldList

```go
type FieldList struct {
	Opening token.Pos // position of opening parenthesis/brace/bracket, if any
	List    []*Field  // field list; or nil
	Closing token.Pos // position of closing parenthesis/brace/bracket, if any
}
```

A FieldList represents a list of Fields, enclosed by parentheses, curly braces, or square brackets.

​	FieldList 表示用括号、花括号或方括号括起来的一系列 Field。

#### (*FieldList) End

```go
func (f *FieldList) End() token.Pos
```

#### (*FieldList) NumFields

```go
func (f *FieldList) NumFields() int
```

NumFields returns the number of parameters or struct fields represented by a FieldList.

​	NumFields 返回 FieldList 表示的参数或结构字段的数量。

#### (*FieldList) Pos

```go
func (f *FieldList) Pos() token.Pos
```

### type File

```go
type File struct {
	Doc     *CommentGroup // associated documentation; or nil
	Package token.Pos     // position of "package" keyword
	Name    *Ident        // package name
	Decls   []Decl        // top-level declarations; or nil

	FileStart, FileEnd token.Pos       // start and end of entire file
	Scope              *Scope          // package scope (this file only)
	Imports            []*ImportSpec   // imports in this file
	Unresolved         []*Ident        // unresolved identifiers in this file
	Comments           []*CommentGroup // list of all comments in the source file
}
```

A File node represents a Go source file.

​	File 节点表示 Go 源文件。

The Comments list contains all comments in the source file in order of appearance, including the comments that are pointed to from other nodes via Doc and Comment fields.

​	Comments 列表按出现顺序包含源文件中的所有注释，包括通过 Doc 和 Comment 字段指向的其他节点中的注释。

For correct printing of source code containing comments (using packages go/format and go/printer), special care must be taken to update comments when a File’s syntax tree is modified: For printing, comments are interspersed between tokens based on their position. If syntax tree nodes are removed or moved, relevant comments in their vicinity must also be removed (from the File.Comments list) or moved accordingly (by updating their positions). A CommentMap may be used to facilitate some of these operations.

​	为了正确打印包含注释的源代码（使用包 go/format 和 go/printer），在修改 File 的语法树时必须特别注意更新注释：为了打印，注释会根据其位置穿插在标记之间。如果删除或移动语法树节点，还必须删除其附近相关的注释（从 File.Comments 列表中删除）或相应地移动（通过更新其位置）。可以使用 CommentMap 来简化其中一些操作。

Whether and how a comment is associated with a node depends on the interpretation of the syntax tree by the manipulating program: Except for Doc and Comment comments directly associated with nodes, the remaining comments are “free-floating” (see also issues #18593, #20744).

​	注释是否以及如何与节点相关联取决于操纵程序对语法树的解释：除了与节点直接关联的 Doc 和 Comment 注释外，其余注释都是“自由浮动”（另请参阅问题 #18593、#20744）。

#### func MergePackageFiles

```go
func MergePackageFiles(pkg *Package, mode MergeMode) *File
```

MergePackageFiles creates a file AST by merging the ASTs of the files belonging to a package. The mode flags control merging behavior.

​	MergePackageFiles 通过合并属于包的文件的 AST 来创建文件 AST。模式标志控制合并行为。

#### (*File) End

```go
func (f *File) End() token.Pos
```

End returns the end of the last declaration in the file. (Use FileEnd for the end of the entire file.)

​	End 返回文件中最后一个声明的结尾。（对于整个文件的结尾，请使用 FileEnd。）

#### (*File) Pos

```go
func (f *File) Pos() token.Pos
```

Pos returns the position of the package declaration. (Use FileStart for the start of the entire file.)

​	Pos 返回包声明的位置。（对于整个文件的开头，请使用 FileStart。）

### type Filter

```go
type Filter func(string) bool
```

### type ForStmt

```go
type ForStmt struct {
	For  token.Pos // position of "for" keyword
	Init Stmt      // initialization statement; or nil
	Cond Expr      // condition; or nil
	Post Stmt      // post iteration statement; or nil
	Body *BlockStmt
}
```

A ForStmt represents a for statement.

​	ForStmt 表示 for 语句。

#### (*ForStmt) End

```go
func (s *ForStmt) End() token.Pos
```

#### (*ForStmt) Pos

```go
func (s *ForStmt) Pos() token.Pos
```

### type FuncDecl

```go
type FuncDecl struct {
	Doc  *CommentGroup // associated documentation; or nil
	Recv *FieldList    // receiver (methods); or nil (functions)
	Name *Ident        // function/method name
	Type *FuncType     // function signature: type and value parameters, results, and position of "func" keyword
	Body *BlockStmt    // function body; or nil for external (non-Go) function
}
```

A FuncDecl node represents a function declaration.

​	FuncDecl 节点表示函数声明。

#### (*FuncDecl) End

```go
func (d *FuncDecl) End() token.Pos
```

#### (*FuncDecl) Pos

```go
func (d *FuncDecl) Pos() token.Pos
```

### type FuncLit

```go
type FuncLit struct {
	Type *FuncType  // function type
	Body *BlockStmt // function body
}
```

A FuncLit node represents a function literal.

​	FuncLit 节点表示函数字面量。

#### (*FuncLit) End

```go
func (x *FuncLit) End() token.Pos
```

#### (*FuncLit) Pos

```go
func (x *FuncLit) Pos() token.Pos
```

### type FuncType

```go
type FuncType struct {
	Func       token.Pos  // position of "func" keyword (token.NoPos if there is no "func")
	TypeParams *FieldList // type parameters; or nil
	Params     *FieldList // (incoming) parameters; non-nil
	Results    *FieldList // (outgoing) results; or nil
}
```

A FuncType node represents a function type.

​	FuncType 节点表示函数类型。

#### (*FuncType) End 

```go
func (x *FuncType) End() token.Pos
```

#### (*FuncType) Pos

```go
func (x *FuncType) Pos() token.Pos
```

### type GenDecl

```go
type GenDecl struct {
	Doc    *CommentGroup // associated documentation; or nil
	TokPos token.Pos     // position of Tok
	Tok    token.Token   // IMPORT, CONST, TYPE, or VAR
	Lparen token.Pos     // position of '(', if any
	Specs  []Spec
	Rparen token.Pos // position of ')', if any
}
```

A GenDecl node (generic declaration node) represents an import, constant, type or variable declaration. A valid Lparen position (Lparen.IsValid()) indicates a parenthesized declaration.

​	GenDecl 节点（泛型声明节点）表示导入、常量、类型或变量声明。有效的 Lparen 位置（Lparen.IsValid()）表示括号声明。

Relationship between Tok value and Specs element type:

​	Tok 值与 Specs 元素类型之间的关系：

```
token.IMPORT  *ImportSpec
token.CONST   *ValueSpec
token.TYPE    *TypeSpec
token.VAR     *ValueSpec
```

#### (*GenDecl) End 

```go
func (d *GenDecl) End() token.Pos
```

#### (*GenDecl) Pos 

```go
func (d *GenDecl) Pos() token.Pos
```

### type GoStmt

```go
type GoStmt struct {
	Go   token.Pos // position of "go" keyword
	Call *CallExpr
}
```

A GoStmt node represents a go statement.

​	GoStmt 节点表示 go 语句。

#### (*GoStmt) End 

```go
func (s *GoStmt) End() token.Pos
```

#### (*GoStmt) Pos

```go
func (s *GoStmt) Pos() token.Pos
```

### type Ident

```go
type Ident struct {
	NamePos token.Pos // identifier position
	Name    string    // identifier name
	Obj     *Object   // denoted object; or nil
}
```

An Ident node represents an identifier.

​	Ident 节点表示标识符。

#### func NewIdent

```go
func NewIdent(name string) *Ident
```

NewIdent creates a new Ident without position. Useful for ASTs generated by code other than the Go parser.

​	NewIdent 创建一个没有位置的新 Ident。对于由 Go 解析器以外的代码生成的 AST 很有用。

#### (*Ident) End

```go
func (x *Ident) End() token.Pos
```

#### (*Ident) IsExported

```go
func (id *Ident) IsExported() bool
```

IsExported reports whether id starts with an upper-case letter.

​	IsExported 报告 id 是否以大写字母开头。

#### (*Ident) Pos

```go
func (x *Ident) Pos() token.Pos
```

#### (*Ident) String

```go
func (id *Ident) String() string
```

### type IfStmt

```go
type IfStmt struct {
	If   token.Pos // position of "if" keyword
	Init Stmt      // initialization statement; or nil
	Cond Expr      // condition
	Body *BlockStmt
	Else Stmt // else branch; or nil
}
```

An IfStmt node represents an if statement.

​	IfStmt 节点表示 if 语句。

#### (*IfStmt) End

```go
func (s *IfStmt) End() token.Pos
```

#### (*IfStmt) Pos

```go
func (s *IfStmt) Pos() token.Pos
```

### type ImportSpec

```go
type ImportSpec struct {
	Doc     *CommentGroup // associated documentation; or nil
	Name    *Ident        // local package name (including "."); or nil
	Path    *BasicLit     // import path
	Comment *CommentGroup // line comments; or nil
	EndPos  token.Pos     // end of spec (overrides Path.Pos if nonzero)
}
```

An ImportSpec node represents a single package import.

​	ImportSpec 节点表示单个包导入。

#### (*ImportSpec) End

```go
func (s *ImportSpec) End() token.Pos
```

#### (*ImportSpec) Pos

```go
func (s *ImportSpec) Pos() token.Pos
```

### type Importer

```go
type Importer func(imports map[string]*Object, path string) (pkg *Object, err error)
```

An Importer resolves import paths to package Objects. The imports map records the packages already imported, indexed by package id (canonical import path). An Importer must determine the canonical import path and check the map to see if it is already present in the imports map. If so, the Importer can return the map entry. Otherwise, the Importer should load the package data for the given path into a new *Object (pkg), record pkg in the imports map, and then return pkg.

​	Importer 将导入路径解析为包对象。imports 映射记录已导入的包，按包 ID（规范导入路径）编制索引。Importer 必须确定规范导入路径并检查映射，以查看它是否已存在于 imports 映射中。如果是，Importer 可以返回映射条目。否则，Importer 应将给定路径的包数据加载到新的 *Object (pkg) 中，将 pkg 记录在 imports 映射中，然后返回 pkg。

### type IncDecStmt

```go
type IncDecStmt struct {
	X      Expr
	TokPos token.Pos   // position of Tok
	Tok    token.Token // INC or DEC
}
```

An IncDecStmt node represents an increment or decrement statement.

​	IncDecStmt 节点表示增量或减量语句。

#### (*IncDecStmt) End

```go
func (s *IncDecStmt) End() token.Pos
```

#### (*IncDecStmt) Pos

```go
func (s *IncDecStmt) Pos() token.Pos
```

### type IndexExpr

```go
type IndexExpr struct {
	X      Expr      // expression
	Lbrack token.Pos // position of "["
	Index  Expr      // index expression
	Rbrack token.Pos // position of "]"
}
```

An IndexExpr node represents an expression followed by an index.

​	IndexExpr 节点表示表达式后跟索引。

#### (*IndexExpr) End

```go
func (x *IndexExpr) End() token.Pos
```

#### (*IndexExpr) Pos

```go
func (x *IndexExpr) Pos() token.Pos
```

### type IndexListExpr <- go1.18

```go
type IndexListExpr struct {
	X       Expr      // expression
	Lbrack  token.Pos // position of "["
	Indices []Expr    // index expressions
	Rbrack  token.Pos // position of "]"
}
```

An IndexListExpr node represents an expression followed by multiple indices.

​	IndexListExpr 节点表示表达式后跟多个索引。

#### (*IndexListExpr) End <- go1.18 

```go
func (x *IndexListExpr) End() token.Pos
```

#### (*IndexListExpr) Pos <- go1.18

```go
func (x *IndexListExpr) Pos() token.Pos
```

### type InterfaceType

```go
type InterfaceType struct {
	Interface  token.Pos  // position of "interface" keyword
	Methods    *FieldList // list of embedded interfaces, methods, or types
	Incomplete bool       // true if (source) methods or types are missing in the Methods list
}
```

An InterfaceType node represents an interface type.

​	InterfaceType 节点表示接口类型。

#### (*InterfaceType) End

```go
func (x *InterfaceType) End() token.Pos
```

#### (*InterfaceType) Pos

```go
func (x *InterfaceType) Pos() token.Pos
```

### type KeyValueExpr

```go
type KeyValueExpr struct {
	Key   Expr
	Colon token.Pos // position of ":"
	Value Expr
}
```

A KeyValueExpr node represents (key : value) pairs in composite literals.

​	KeyValueExpr 节点表示复合字面量中的 (key : value) 对。

#### (*KeyValueExpr) End

```go
func (x *KeyValueExpr) End() token.Pos
```

#### (*KeyValueExpr) Pos

```go
func (x *KeyValueExpr) Pos() token.Pos
```

### type LabeledStmt

```go
type LabeledStmt struct {
	Label *Ident
	Colon token.Pos // position of ":"
	Stmt  Stmt
}
```

A LabeledStmt node represents a labeled statement.

​	LabeledStmt 节点表示带标签的语句。

#### (*LabeledStmt) End

```go
func (s *LabeledStmt) End() token.Pos
```

#### (*LabeledStmt) Pos

```go
func (s *LabeledStmt) Pos() token.Pos
```

### type MapType

```go
type MapType struct {
	Map   token.Pos // position of "map" keyword
	Key   Expr
	Value Expr
}
```

A MapType node represents a map type.

​	MapType 节点表示映射类型。

#### (*MapType) End

```go
func (x *MapType) End() token.Pos
```

#### (*MapType) Pos

```go
func (x *MapType) Pos() token.Pos
```

### type MergeMode

```go
type MergeMode uint
```

The MergeMode flags control the behavior of MergePackageFiles.

​	MergeMode 标志控制 MergePackageFiles 的行为。

```go
const (
	// If set, duplicate function declarations are excluded.
	FilterFuncDuplicates MergeMode = 1 << iota
	// If set, comments that are not associated with a specific
	// AST node (as Doc or Comment) are excluded.
	FilterUnassociatedComments
	// If set, duplicate import declarations are excluded.
	FilterImportDuplicates
)
```

### type Node

```go
type Node interface {
	Pos() token.Pos // position of first character belonging to the node
	End() token.Pos // position of first character immediately after the node
}
```

All node types implement the Node interface.

​	所有节点类型都实现 Node 接口。

### type ObjKind 

```go
type ObjKind int
```

ObjKind describes what an object represents.

​	ObjKind 描述对象表示什么。

```go
const (
	Bad ObjKind = iota // for error handling
	Pkg                // package
	Con                // constant
	Typ                // type
	Var                // variable
	Fun                // function or method
	Lbl                // label
)
```

The list of possible Object kinds.

​	可能的对象种类的列表。

#### (ObjKind) String

```go
func (kind ObjKind) String() string
```

### type Object

```go
type Object struct {
	Kind ObjKind
	Name string // declared name
	Decl any    // corresponding Field, XxxSpec, FuncDecl, LabeledStmt, AssignStmt, Scope; or nil
	Data any    // object-specific data; or nil
	Type any    // placeholder for type information; may be nil
}
```

An Object describes a named language entity such as a package, constant, type, variable, function (incl. methods), or label.

​	对象描述命名的语言实体，例如包、常量、类型、变量、函数（包括方法）或标签。

The Data fields contains object-specific data:

​	Data 字段包含特定于对象的数据：

```
Kind    Data type         Data value
Pkg     *Scope            package scope
Con     int               iota for the respective declaration
```

#### func NewObj

```go
func NewObj(kind ObjKind, name string) *Object
```

NewObj creates a new object of a given kind and name.

​	NewObj 创建一个给定类型和名称的新对象。

#### (*Object) Pos

```go
func (obj *Object) Pos() token.Pos
```

Pos computes the source position of the declaration of an object name. The result may be an invalid position if it cannot be computed (obj.Decl may be nil or not correct).

​	Pos 计算对象名称声明的源位置。如果无法计算，结果可能是一个无效的位置（obj.Decl 可能为 nil 或不正确）。

### type Package

```go
type Package struct {
	Name    string             // package name
	Scope   *Scope             // package scope across all files
	Imports map[string]*Object // map of package id -> package object
	Files   map[string]*File   // Go source files by filename
}
```

A Package node represents a set of source files collectively building a Go package.

​	Package 节点表示一组源文件，共同构建一个 Go 包。

#### func NewPackage

```go
func NewPackage(fset *token.FileSet, files map[string]*File, importer Importer, universe *Scope) (*Package, error)
```

NewPackage creates a new Package node from a set of File nodes. It resolves unresolved identifiers across files and updates each file’s Unresolved list accordingly. If a non-nil importer and universe scope are provided, they are used to resolve identifiers not declared in any of the package files. Any remaining unresolved identifiers are reported as undeclared. If the files belong to different packages, one package name is selected and files with different package names are reported and then ignored. The result is a package node and a scanner.ErrorList if there were errors.

​	NewPackage 从一组 File 节点创建新的 Package 节点。它解析文件中的未解析标识符，并相应地更新每个文件的 Unresolved 列表。如果提供了非 nil 的导入程序和 universe 作用域，则使用它们来解析未在任何包文件中声明的标识符。任何剩余的未解析标识符都将报告为未声明。如果文件属于不同的包，则选择一个包名称，并报告具有不同包名称的文件，然后忽略它们。结果是一个包节点和一个 scanner.ErrorList（如果有错误）。

#### (*Package) End

```go
func (p *Package) End() token.Pos
```

#### (*Package) Pos

```go
func (p *Package) Pos() token.Pos
```

### type ParenExpr

```go
type ParenExpr struct {
	Lparen token.Pos // position of "("
	X      Expr      // parenthesized expression
	Rparen token.Pos // position of ")"
}
```

A ParenExpr node represents a parenthesized expression.

​	ParenExpr 节点表示括号表达式。

#### (*ParenExpr) End

```go
func (x *ParenExpr) End() token.Pos
```

#### (*ParenExpr) Pos

```go
func (x *ParenExpr) Pos() token.Pos
```

### type RangeStmt

```go
type RangeStmt struct {
	For        token.Pos   // position of "for" keyword
	Key, Value Expr        // Key, Value may be nil
	TokPos     token.Pos   // position of Tok; invalid if Key == nil
	Tok        token.Token // ILLEGAL if Key == nil, ASSIGN, DEFINE
	Range      token.Pos   // position of "range" keyword
	X          Expr        // value to range over
	Body       *BlockStmt
}
```

A RangeStmt represents a for statement with a range clause.

​	RangeStmt 表示带有范围子句的 for 语句。

#### (*RangeStmt) End

```go
func (s *RangeStmt) End() token.Pos
```

#### (*RangeStmt) Pos

```go
func (s *RangeStmt) Pos() token.Pos
```

### type ReturnStmt

```go
type ReturnStmt struct {
	Return  token.Pos // position of "return" keyword
	Results []Expr    // result expressions; or nil
}
```

A ReturnStmt node represents a return statement.

​	ReturnStmt 节点表示 return 语句。

#### (*ReturnStmt) End

```go
func (s *ReturnStmt) End() token.Pos
```

#### (*ReturnStmt) Pos 

```go
func (s *ReturnStmt) Pos() token.Pos
```

### type Scope

```go
type Scope struct {
	Outer   *Scope
	Objects map[string]*Object
}
```

A Scope maintains the set of named language entities declared in the scope and a link to the immediately surrounding (outer) scope.

​	范围维护在范围内声明的命名语言实体集以及与紧邻的（外部）范围的链接。

#### func NewScope

```go
func NewScope(outer *Scope) *Scope
```

NewScope creates a new scope nested in the outer scope.

​	NewScope 在外部范围内创建一个新的嵌套范围。

#### (*Scope) Insert 

```go
func (s *Scope) Insert(obj *Object) (alt *Object)
```

Insert attempts to insert a named object obj into the scope s. If the scope already contains an object alt with the same name, Insert leaves the scope unchanged and returns alt. Otherwise it inserts obj and returns nil.

​	Insert 尝试将命名对象 obj 插入范围 s。如果范围已经包含具有相同名称的对象 alt，则 Insert 将范围保持不变并返回 alt。否则，它将插入 obj 并返回 nil。

#### (*Scope) Lookup

```go
func (s *Scope) Lookup(name string) *Object
```

Lookup returns the object with the given name if it is found in scope s, otherwise it returns nil. Outer scopes are ignored.

​	如果在范围 s 中找到具有给定名称的对象，则 Lookup 返回该对象，否则返回 nil。忽略外部范围。

#### (*Scope) String 

```go
func (s *Scope) String() string
```

Debugging support

​	调试支持

### type SelectStmt

```go
type SelectStmt struct {
	Select token.Pos  // position of "select" keyword
	Body   *BlockStmt // CommClauses only
}
```

A SelectStmt node represents a select statement.

​	SelectStmt 节点表示一个 select 语句。

#### (*SelectStmt) End 

```go
func (s *SelectStmt) End() token.Pos
```

#### (*SelectStmt) Pos

```go
func (s *SelectStmt) Pos() token.Pos
```

### type SelectorExpr

```go
type SelectorExpr struct {
	X   Expr   // expression
	Sel *Ident // field selector
}
```

A SelectorExpr node represents an expression followed by a selector.

​	SelectorExpr 节点表示一个表达式，后跟一个选择器。

#### (*SelectorExpr) End 

```go
func (x *SelectorExpr) End() token.Pos
```

#### (*SelectorExpr) Pos 

```go
func (x *SelectorExpr) Pos() token.Pos
```

### type SendStmt

```go
type SendStmt struct {
	Chan  Expr
	Arrow token.Pos // position of "<-"
	Value Expr
}
```

A SendStmt node represents a send statement.

​	SendStmt 节点表示一个 send 语句。

#### (*SendStmt) End

```go
func (s *SendStmt) End() token.Pos
```

#### (*SendStmt) Pos

```go
func (s *SendStmt) Pos() token.Pos
```

### type SliceExpr

```go
type SliceExpr struct {
	X      Expr      // expression
	Lbrack token.Pos // position of "["
	Low    Expr      // begin of slice range; or nil
	High   Expr      // end of slice range; or nil
	Max    Expr      // maximum capacity of slice; or nil
	Slice3 bool      // true if 3-index slice (2 colons present)
	Rbrack token.Pos // position of "]"
}
```

A SliceExpr node represents an expression followed by slice indices.

​	SliceExpr 节点表示一个表达式，后跟切片索引。

#### (*SliceExpr) End

```go
func (x *SliceExpr) End() token.Pos
```

#### (*SliceExpr) Pos

```go
func (x *SliceExpr) Pos() token.Pos
```

### type Spec

```go
type Spec interface {
	Node
	// contains filtered or unexported methods
}
```

The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.

​	Spec 类型表示 *ImportSpec、*ValueSpec 和 *TypeSpec 中的任何一个。

### type StarExpr

```go
type StarExpr struct {
	Star token.Pos // position of "*"
	X    Expr      // operand
}
```

A StarExpr node represents an expression of the form “*” Expression. Semantically it could be a unary “*” expression, or a pointer type.

​	StarExpr 节点表示形式为 “” Expression 的表达式。语义上它可以是单一的 “” 表达式，或指针类型。

#### (*StarExpr) End

```go
func (x *StarExpr) End() token.Pos
```

#### (*StarExpr) Pos

```go
func (x *StarExpr) Pos() token.Pos
```

### type Stmt

```go
type Stmt interface {
	Node
	// contains filtered or unexported methods
}
```

All statement nodes implement the Stmt interface.

​	所有语句节点都实现了 Stmt 接口。

### type StructType

```go
type StructType struct {
	Struct     token.Pos  // position of "struct" keyword
	Fields     *FieldList // list of field declarations
	Incomplete bool       // true if (source) fields are missing in the Fields list
}
```

A StructType node represents a struct type.

​	StructType 节点表示结构类型。

#### (*StructType) End

```go
func (x *StructType) End() token.Pos
```

#### (*StructType) Pos

```go
func (x *StructType) Pos() token.Pos
```

### type SwitchStmt

```go
type SwitchStmt struct {
	Switch token.Pos  // position of "switch" keyword
	Init   Stmt       // initialization statement; or nil
	Tag    Expr       // tag expression; or nil
	Body   *BlockStmt // CaseClauses only
}
```

A SwitchStmt node represents an expression switch statement.

​	SwitchStmt 节点表示表达式 switch 语句。

#### (*SwitchStmt) End

```go
func (s *SwitchStmt) End() token.Pos
```

#### (*SwitchStmt) Pos

```go
func (s *SwitchStmt) Pos() token.Pos
```

### type TypeAssertExpr

```go
type TypeAssertExpr struct {
	X      Expr      // expression
	Lparen token.Pos // position of "("
	Type   Expr      // asserted type; nil means type switch X.(type)
	Rparen token.Pos // position of ")"
}
```

A TypeAssertExpr node represents an expression followed by a type assertion.

​	TypeAssertExpr 节点表示一个表达式，后跟一个类型断言。

#### (*TypeAssertExpr) End

```go
func (x *TypeAssertExpr) End() token.Pos
```

#### (*TypeAssertExpr) Pos

```go
func (x *TypeAssertExpr) Pos() token.Pos
```

### type TypeSpec

```go
type TypeSpec struct {
	Doc        *CommentGroup // associated documentation; or nil
	Name       *Ident        // type name
	TypeParams *FieldList    // type parameters; or nil
	Assign     token.Pos     // position of '=', if any
	Type       Expr          // *Ident, *ParenExpr, *SelectorExpr, *StarExpr, or any of the *XxxTypes
	Comment    *CommentGroup // line comments; or nil
}
```

A TypeSpec node represents a type declaration (TypeSpec production).

​	TypeSpec 节点表示类型声明（TypeSpec 生成）。

#### (*TypeSpec) End

```go
func (s *TypeSpec) End() token.Pos
```

#### (*TypeSpec) Pos 

```go
func (s *TypeSpec) Pos() token.Pos
```

### type TypeSwitchStmt

```go
type TypeSwitchStmt struct {
	Switch token.Pos  // position of "switch" keyword
	Init   Stmt       // initialization statement; or nil
	Assign Stmt       // x := y.(type) or y.(type)
	Body   *BlockStmt // CaseClauses only
}
```

A TypeSwitchStmt node represents a type switch statement.

​	TypeSwitchStmt节点表示类型switch语句。

#### (*TypeSwitchStmt) End 

```go
func (s *TypeSwitchStmt) End() token.Pos
```

#### (*TypeSwitchStmt) Pos 

```go
func (s *TypeSwitchStmt) Pos() token.Pos
```

### type UnaryExpr

```go
type UnaryExpr struct {
	OpPos token.Pos   // position of Op
	Op    token.Token // operator
	X     Expr        // operand
}
```

A UnaryExpr node represents a unary expression. Unary “*” expressions are represented via StarExpr nodes.

​	UnaryExpr节点表示一元表达式。一元“*”表达式通过StarExpr节点表示。

#### (*UnaryExpr) End

```go
func (x *UnaryExpr) End() token.Pos
```

#### (*UnaryExpr) Pos 

```go
func (x *UnaryExpr) Pos() token.Pos
```

### type ValueSpec

```go
type ValueSpec struct {
	Doc     *CommentGroup // associated documentation; or nil
	Names   []*Ident      // value names (len(Names) > 0)
	Type    Expr          // value type; or nil
	Values  []Expr        // initial values; or nil
	Comment *CommentGroup // line comments; or nil
}
```

A ValueSpec node represents a constant or variable declaration (ConstSpec or VarSpec production).

​	A ValueSpec 节点表示常量或变量声明（ConstSpec 或 VarSpec 生成）。

#### (*ValueSpec) End 

```go
func (s *ValueSpec) End() token.Pos
```

#### (*ValueSpec) Pos 

```go
func (s *ValueSpec) Pos() token.Pos
```

### type Visitor

```go
type Visitor interface {
	Visit(node Node) (w Visitor)
}
```

A Visitor’s Visit method is invoked for each node encountered by Walk. If the result visitor w is not nil, Walk visits each of the children of node with the visitor w, followed by a call of w.Visit(nil).

​	Walk 遍历到的每个节点都会调用 Visitor 的 Visit 方法。如果结果访问者 w 不为 nil，Walk 会使用访问者 w 遍历节点的每个子节点，然后调用 w.Visit(nil)。