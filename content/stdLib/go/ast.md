+++
title = "ast"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# ast

https://pkg.go.dev/go/ast@go1.20.1



Package ast declares the types used to represent syntax trees for Go packages.









## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func FileExports 

``` go 
func FileExports(src *File) bool
```

FileExports trims the AST for a Go source file in place such that only exported nodes remain: all top-level identifiers which are not exported and their associated information (such as type, initial value, or function body) are removed. Non-exported fields and methods of exported types are stripped. The File.Comments list is not changed.

FileExports reports whether there are exported declarations.

#### func FilterDecl 

``` go 
func FilterDecl(decl Decl, f Filter) bool
```

FilterDecl trims the AST for a Go declaration in place by removing all names (including struct field and interface method names, but not from parameter lists) that don't pass through the filter f.

FilterDecl reports whether there are any declared names left after filtering.

#### func FilterFile 

``` go 
func FilterFile(src *File, f Filter) bool
```

FilterFile trims the AST for a Go file in place by removing all names from top-level declarations (including struct field and interface method names, but not from parameter lists) that don't pass through the filter f. If the declaration is empty afterwards, the declaration is removed from the AST. Import declarations are always removed. The File.Comments list is not changed.

FilterFile reports whether there are any top-level declarations left after filtering.

#### func FilterPackage 

``` go 
func FilterPackage(pkg *Package, f Filter) bool
```

FilterPackage trims the AST for a Go package in place by removing all names from top-level declarations (including struct field and interface method names, but not from parameter lists) that don't pass through the filter f. If the declaration is empty afterwards, the declaration is removed from the AST. The pkg.Files list is not changed, so that file names and top-level package comments don't get lost.

FilterPackage reports whether there are any top-level declarations left after filtering.

#### func Fprint 

``` go 
func Fprint(w io.Writer, fset *token.FileSet, x any, f FieldFilter) error
```

Fprint prints the (sub-)tree starting at AST node x to w. If fset != nil, position information is interpreted relative to that file set. Otherwise positions are printed as integer values (file set specific offsets).

A non-nil FieldFilter f may be provided to control the output: struct fields for which f(fieldname, fieldvalue) is true are printed; all others are filtered from the output. Unexported struct fields are never printed.

#### func Inspect 

``` go 
func Inspect(node Node, f func(Node) bool)
```

Inspect traverses an AST in depth-first order: It starts by calling f(node); node must not be nil. If f returns true, Inspect invokes f recursively for each of the non-nil children of node, followed by a call of f(nil).

##### Example
``` go 
```

#### func IsExported 

``` go 
func IsExported(name string) bool
```

IsExported reports whether name starts with an upper-case letter.

#### func NotNilFilter 

``` go 
func NotNilFilter(_ string, v reflect.Value) bool
```

NotNilFilter returns true for field values that are not nil; it returns false otherwise.

#### func PackageExports 

``` go 
func PackageExports(pkg *Package) bool
```

PackageExports trims the AST for a Go package in place such that only exported nodes remain. The pkg.Files list is not changed, so that file names and top-level package comments don't get lost.

PackageExports reports whether there are exported declarations; it returns false otherwise.

#### func Print 

``` go 
func Print(fset *token.FileSet, x any) error
```

Print prints x to standard output, skipping nil fields. Print(fset, x) is the same as Fprint(os.Stdout, fset, x, NotNilFilter).

##### Example
``` go 
```

#### func SortImports 

``` go 
func SortImports(fset *token.FileSet, f *File)
```

SortImports sorts runs of consecutive import lines in import blocks in f. It also removes duplicate imports when it is possible to do so without data loss.

#### func Walk 

``` go 
func Walk(v Visitor, node Node)
```

Walk traverses an AST in depth-first order: It starts by calling v.Visit(node); node must not be nil. If the visitor w returned by v.Visit(node) is not nil, Walk is invoked recursively with visitor w for each of the non-nil children of node, followed by a call of w.Visit(nil).

## 类型

### type ArrayType 

``` go 
type ArrayType struct {
	Lbrack token.Pos // position of "["
	Len    Expr      // Ellipsis node for [...]T array types, nil for slice types
	Elt    Expr      // element type
}
```

An ArrayType node represents an array or slice type.

#### (*ArrayType) End 

``` go 
func (x *ArrayType) End() token.Pos
```

#### (*ArrayType) Pos 

``` go 
func (x *ArrayType) Pos() token.Pos
```

### type AssignStmt 

``` go 
type AssignStmt struct {
	Lhs    []Expr
	TokPos token.Pos   // position of Tok
	Tok    token.Token // assignment token, DEFINE
	Rhs    []Expr
}
```

An AssignStmt node represents an assignment or a short variable declaration.

#### (*AssignStmt) End 

``` go 
func (s *AssignStmt) End() token.Pos
```

#### (*AssignStmt) Pos 

``` go 
func (s *AssignStmt) Pos() token.Pos
```

### type BadDecl 

``` go 
type BadDecl struct {
	From, To token.Pos // position range of bad declaration
}
```

A BadDecl node is a placeholder for a declaration containing syntax errors for which a correct declaration node cannot be created.

#### (*BadDecl) End 

``` go 
func (d *BadDecl) End() token.Pos
```

#### (*BadDecl) Pos 

``` go 
func (d *BadDecl) Pos() token.Pos
```

### type BadExpr 

``` go 
type BadExpr struct {
	From, To token.Pos // position range of bad expression
}
```

A BadExpr node is a placeholder for an expression containing syntax errors for which a correct expression node cannot be created.

#### (*BadExpr) End 

``` go 
func (x *BadExpr) End() token.Pos
```

#### (*BadExpr) Pos 

``` go 
func (x *BadExpr) Pos() token.Pos
```

### type BadStmt 

``` go 
type BadStmt struct {
	From, To token.Pos // position range of bad statement
}
```

A BadStmt node is a placeholder for statements containing syntax errors for which no correct statement nodes can be created.

#### (*BadStmt) End 

``` go 
func (s *BadStmt) End() token.Pos
```

#### (*BadStmt) Pos 

``` go 
func (s *BadStmt) Pos() token.Pos
```

### type BasicLit 

``` go 
type BasicLit struct {
	ValuePos token.Pos   // literal position
	Kind     token.Token // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
	Value    string      // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
}
```

A BasicLit node represents a literal of basic type.

#### (*BasicLit) End 

``` go 
func (x *BasicLit) End() token.Pos
```

#### (*BasicLit) Pos 

``` go 
func (x *BasicLit) Pos() token.Pos
```

### type BinaryExpr 

``` go 
type BinaryExpr struct {
	X     Expr        // left operand
	OpPos token.Pos   // position of Op
	Op    token.Token // operator
	Y     Expr        // right operand
}
```

A BinaryExpr node represents a binary expression.

#### (*BinaryExpr) End 

``` go 
func (x *BinaryExpr) End() token.Pos
```

#### (*BinaryExpr) Pos 

``` go 
func (x *BinaryExpr) Pos() token.Pos
```

### type BlockStmt 

``` go 
type BlockStmt struct {
	Lbrace token.Pos // position of "{"
	List   []Stmt
	Rbrace token.Pos // position of "}", if any (may be absent due to syntax error)
}
```

A BlockStmt node represents a braced statement list.

#### (*BlockStmt) End 

``` go 
func (s *BlockStmt) End() token.Pos
```

#### (*BlockStmt) Pos 

``` go 
func (s *BlockStmt) Pos() token.Pos
```

### type BranchStmt 

``` go 
type BranchStmt struct {
	TokPos token.Pos   // position of Tok
	Tok    token.Token // keyword token (BREAK, CONTINUE, GOTO, FALLTHROUGH)
	Label  *Ident      // label name; or nil
}
```

A BranchStmt node represents a break, continue, goto, or fallthrough statement.

#### (*BranchStmt) End 

``` go 
func (s *BranchStmt) End() token.Pos
```

#### (*BranchStmt) Pos 

``` go 
func (s *BranchStmt) Pos() token.Pos
```

### type CallExpr 

``` go 
type CallExpr struct {
	Fun      Expr      // function expression
	Lparen   token.Pos // position of "("
	Args     []Expr    // function arguments; or nil
	Ellipsis token.Pos // position of "..." (token.NoPos if there is no "...")
	Rparen   token.Pos // position of ")"
}
```

A CallExpr node represents an expression followed by an argument list.

#### (*CallExpr) End 

``` go 
func (x *CallExpr) End() token.Pos
```

#### (*CallExpr) Pos 

``` go 
func (x *CallExpr) Pos() token.Pos
```

### type CaseClause 

``` go 
type CaseClause struct {
	Case  token.Pos // position of "case" or "default" keyword
	List  []Expr    // list of expressions or types; nil means default case
	Colon token.Pos // position of ":"
	Body  []Stmt    // statement list; or nil
}
```

A CaseClause represents a case of an expression or type switch statement.

#### (*CaseClause) End 

``` go 
func (s *CaseClause) End() token.Pos
```

#### (*CaseClause) Pos 

``` go 
func (s *CaseClause) Pos() token.Pos
```

### type ChanDir 

``` go 
type ChanDir int
```

The direction of a channel type is indicated by a bit mask including one or both of the following constants.

``` go 
const (
	SEND ChanDir = 1 << iota
	RECV
)
```

### type ChanType 

``` go 
type ChanType struct {
	Begin token.Pos // position of "chan" keyword or "<-" (whichever comes first)
	Arrow token.Pos // position of "<-" (token.NoPos if there is no "<-")
	Dir   ChanDir   // channel direction
	Value Expr      // value type
}
```

A ChanType node represents a channel type.

#### (*ChanType) End 

``` go 
func (x *ChanType) End() token.Pos
```

#### (*ChanType) Pos 

``` go 
func (x *ChanType) Pos() token.Pos
```

### type CommClause 

``` go 
type CommClause struct {
	Case  token.Pos // position of "case" or "default" keyword
	Comm  Stmt      // send or receive statement; nil means default case
	Colon token.Pos // position of ":"
	Body  []Stmt    // statement list; or nil
}
```

A CommClause node represents a case of a select statement.

#### (*CommClause) End 

``` go 
func (s *CommClause) End() token.Pos
```

#### (*CommClause) Pos 

``` go 
func (s *CommClause) Pos() token.Pos
```

### type Comment 

``` go 
type Comment struct {
	Slash token.Pos // position of "/" starting the comment
	Text  string    // comment text (excluding '\n' for //-style comments)
}
```

A Comment node represents a single //-style or /*-style comment.

The Text field contains the comment text without carriage returns (\r) that may have been present in the source. Because a comment's end position is computed using len(Text), the position reported by End() does not match the true source end position for comments containing carriage returns.

#### (*Comment) End 

``` go 
func (c *Comment) End() token.Pos
```

#### (*Comment) Pos 

``` go 
func (c *Comment) Pos() token.Pos
```

### type CommentGroup 

``` go 
type CommentGroup struct {
	List []*Comment // len(List) > 0
}
```

A CommentGroup represents a sequence of comments with no other tokens and no empty lines between.

#### (*CommentGroup) End 

``` go 
func (g *CommentGroup) End() token.Pos
```

#### (*CommentGroup) Pos 

``` go 
func (g *CommentGroup) Pos() token.Pos
```

#### (*CommentGroup) Text 

``` go 
func (g *CommentGroup) Text() string
```

Text returns the text of the comment. Comment markers (//, /*, and */), the first space of a line comment, and leading and trailing empty lines are removed. Comment directives like "//line" and "//go:noinline" are also removed. Multiple empty lines are reduced to one, and trailing space on lines is trimmed. Unless the result is empty, it is newline-terminated.

### type CommentMap  <- go1.1

``` go 
type CommentMap map[Node][]*CommentGroup
```

A CommentMap maps an AST node to a list of comment groups associated with it. See NewCommentMap for a description of the association.

##### Example
``` go 
```

#### func NewCommentMap  <- go1.1

``` go 
func NewCommentMap(fset *token.FileSet, node Node, comments []*CommentGroup) CommentMap
```

NewCommentMap creates a new comment map by associating comment groups of the comments list with the nodes of the AST specified by node.

A comment group g is associated with a node n if:

- g starts on the same line as n ends
- g starts on the line immediately following n, and there is at least one empty line after g and before the next node
- g starts before n and is not associated to the node before n via the previous rules

NewCommentMap tries to associate a comment group to the "largest" node possible: For instance, if the comment is a line comment trailing an assignment, the comment is associated with the entire assignment rather than just the last operand in the assignment.

#### (CommentMap) Comments  <- go1.1

``` go 
func (cmap CommentMap) Comments() []*CommentGroup
```

Comments returns the list of comment groups in the comment map. The result is sorted in source order.

#### (CommentMap) Filter  <- go1.1

``` go 
func (cmap CommentMap) Filter(node Node) CommentMap
```

Filter returns a new comment map consisting of only those entries of cmap for which a corresponding node exists in the AST specified by node.

#### (CommentMap) String  <- go1.1

``` go 
func (cmap CommentMap) String() string
```

#### (CommentMap) Update  <- go1.1

``` go 
func (cmap CommentMap) Update(old, new Node) Node
```

Update replaces an old node in the comment map with the new node and returns the new node. Comments that were associated with the old node are associated with the new node.

### type CompositeLit 

``` go 
type CompositeLit struct {
	Type       Expr      // literal type; or nil
	Lbrace     token.Pos // position of "{"
	Elts       []Expr    // list of composite elements; or nil
	Rbrace     token.Pos // position of "}"
	Incomplete bool      // true if (source) expressions are missing in the Elts list
}
```

A CompositeLit node represents a composite literal.

#### (*CompositeLit) End 

``` go 
func (x *CompositeLit) End() token.Pos
```

#### (*CompositeLit) Pos 

``` go 
func (x *CompositeLit) Pos() token.Pos
```

### type Decl 

``` go 
type Decl interface {
	Node
	// contains filtered or unexported methods
}
```

All declaration nodes implement the Decl interface.

### type DeclStmt 

``` go 
type DeclStmt struct {
	Decl Decl // *GenDecl with CONST, TYPE, or VAR token
}
```

A DeclStmt node represents a declaration in a statement list.

#### (*DeclStmt) End 

``` go 
func (s *DeclStmt) End() token.Pos
```

#### (*DeclStmt) Pos 

``` go 
func (s *DeclStmt) Pos() token.Pos
```

### type DeferStmt 

``` go 
type DeferStmt struct {
	Defer token.Pos // position of "defer" keyword
	Call  *CallExpr
}
```

A DeferStmt node represents a defer statement.

#### (*DeferStmt) End 

``` go 
func (s *DeferStmt) End() token.Pos
```

#### (*DeferStmt) Pos 

``` go 
func (s *DeferStmt) Pos() token.Pos
```

### type Ellipsis 

``` go 
type Ellipsis struct {
	Ellipsis token.Pos // position of "..."
	Elt      Expr      // ellipsis element type (parameter lists only); or nil
}
```

An Ellipsis node stands for the "..." type in a parameter list or the "..." length in an array type.

#### (*Ellipsis) End 

``` go 
func (x *Ellipsis) End() token.Pos
```

#### (*Ellipsis) Pos 

``` go 
func (x *Ellipsis) Pos() token.Pos
```

### type EmptyStmt 

``` go 
type EmptyStmt struct {
	Semicolon token.Pos // position of following ";"
	Implicit  bool      // if set, ";" was omitted in the source
}
```

An EmptyStmt node represents an empty statement. The "position" of the empty statement is the position of the immediately following (explicit or implicit) semicolon.

#### (*EmptyStmt) End 

``` go 
func (s *EmptyStmt) End() token.Pos
```

#### (*EmptyStmt) Pos 

``` go 
func (s *EmptyStmt) Pos() token.Pos
```

### type Expr 

``` go 
type Expr interface {
	Node
	// contains filtered or unexported methods
}
```

All expression nodes implement the Expr interface.

### type ExprStmt 

``` go 
type ExprStmt struct {
	X Expr // expression
}
```

An ExprStmt node represents a (stand-alone) expression in a statement list.

#### (*ExprStmt) End 

``` go 
func (s *ExprStmt) End() token.Pos
```

#### (*ExprStmt) Pos 

``` go 
func (s *ExprStmt) Pos() token.Pos
```

### type Field 

``` go 
type Field struct {
	Doc     *CommentGroup // associated documentation; or nil
	Names   []*Ident      // field/method/(type) parameter names; or nil
	Type    Expr          // field/method/parameter type; or nil
	Tag     *BasicLit     // field tag; or nil
	Comment *CommentGroup // line comments; or nil
}
```

A Field represents a Field declaration list in a struct type, a method list in an interface type, or a parameter/result declaration in a signature. Field.Names is nil for unnamed parameters (parameter lists which only contain types) and embedded struct fields. In the latter case, the field name is the type name.

#### (*Field) End 

``` go 
func (f *Field) End() token.Pos
```

#### (*Field) Pos 

``` go 
func (f *Field) Pos() token.Pos
```

### type FieldFilter 

``` go 
type FieldFilter func(name string, value reflect.Value) bool
```

A FieldFilter may be provided to Fprint to control the output.

### type FieldList 

``` go 
type FieldList struct {
	Opening token.Pos // position of opening parenthesis/brace/bracket, if any
	List    []*Field  // field list; or nil
	Closing token.Pos // position of closing parenthesis/brace/bracket, if any
}
```

A FieldList represents a list of Fields, enclosed by parentheses, curly braces, or square brackets.

#### (*FieldList) End 

``` go 
func (f *FieldList) End() token.Pos
```

#### (*FieldList) NumFields 

``` go 
func (f *FieldList) NumFields() int
```

NumFields returns the number of parameters or struct fields represented by a FieldList.

#### (*FieldList) Pos 

``` go 
func (f *FieldList) Pos() token.Pos
```

### type File 

``` go 
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

The Comments list contains all comments in the source file in order of appearance, including the comments that are pointed to from other nodes via Doc and Comment fields.

For correct printing of source code containing comments (using packages go/format and go/printer), special care must be taken to update comments when a File's syntax tree is modified: For printing, comments are interspersed between tokens based on their position. If syntax tree nodes are removed or moved, relevant comments in their vicinity must also be removed (from the File.Comments list) or moved accordingly (by updating their positions). A CommentMap may be used to facilitate some of these operations.

Whether and how a comment is associated with a node depends on the interpretation of the syntax tree by the manipulating program: Except for Doc and Comment comments directly associated with nodes, the remaining comments are "free-floating" (see also issues #18593, #20744).

#### func MergePackageFiles 

``` go 
func MergePackageFiles(pkg *Package, mode MergeMode) *File
```

MergePackageFiles creates a file AST by merging the ASTs of the files belonging to a package. The mode flags control merging behavior.

#### (*File) End 

``` go 
func (f *File) End() token.Pos
```

End returns the end of the last declaration in the file. (Use FileEnd for the end of the entire file.)

#### (*File) Pos 

``` go 
func (f *File) Pos() token.Pos
```

Pos returns the position of the package declaration. (Use FileStart for the start of the entire file.)

### type Filter 

``` go 
type Filter func(string) bool
```

### type ForStmt 

``` go 
type ForStmt struct {
	For  token.Pos // position of "for" keyword
	Init Stmt      // initialization statement; or nil
	Cond Expr      // condition; or nil
	Post Stmt      // post iteration statement; or nil
	Body *BlockStmt
}
```

A ForStmt represents a for statement.

#### (*ForStmt) End 

``` go 
func (s *ForStmt) End() token.Pos
```

#### (*ForStmt) Pos 

``` go 
func (s *ForStmt) Pos() token.Pos
```

### type FuncDecl 

``` go 
type FuncDecl struct {
	Doc  *CommentGroup // associated documentation; or nil
	Recv *FieldList    // receiver (methods); or nil (functions)
	Name *Ident        // function/method name
	Type *FuncType     // function signature: type and value parameters, results, and position of "func" keyword
	Body *BlockStmt    // function body; or nil for external (non-Go) function
}
```

A FuncDecl node represents a function declaration.

#### (*FuncDecl) End 

``` go 
func (d *FuncDecl) End() token.Pos
```

#### (*FuncDecl) Pos 

``` go 
func (d *FuncDecl) Pos() token.Pos
```

### type FuncLit 

``` go 
type FuncLit struct {
	Type *FuncType  // function type
	Body *BlockStmt // function body
}
```

A FuncLit node represents a function literal.

#### (*FuncLit) End 

``` go 
func (x *FuncLit) End() token.Pos
```

#### (*FuncLit) Pos 

``` go 
func (x *FuncLit) Pos() token.Pos
```

### type FuncType 

``` go 
type FuncType struct {
	Func       token.Pos  // position of "func" keyword (token.NoPos if there is no "func")
	TypeParams *FieldList // type parameters; or nil
	Params     *FieldList // (incoming) parameters; non-nil
	Results    *FieldList // (outgoing) results; or nil
}
```

A FuncType node represents a function type.

#### (*FuncType) End 

``` go 
func (x *FuncType) End() token.Pos
```

#### (*FuncType) Pos 

``` go 
func (x *FuncType) Pos() token.Pos
```

### type GenDecl 

``` go 
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

Relationship between Tok value and Specs element type:

```
token.IMPORT  *ImportSpec
token.CONST   *ValueSpec
token.TYPE    *TypeSpec
token.VAR     *ValueSpec
```

#### (*GenDecl) End 

``` go 
func (d *GenDecl) End() token.Pos
```

#### (*GenDecl) Pos 

``` go 
func (d *GenDecl) Pos() token.Pos
```

### type GoStmt 

``` go 
type GoStmt struct {
	Go   token.Pos // position of "go" keyword
	Call *CallExpr
}
```

A GoStmt node represents a go statement.

#### (*GoStmt) End 

``` go 
func (s *GoStmt) End() token.Pos
```

#### (*GoStmt) Pos 

``` go 
func (s *GoStmt) Pos() token.Pos
```

### type Ident 

``` go 
type Ident struct {
	NamePos token.Pos // identifier position
	Name    string    // identifier name
	Obj     *Object   // denoted object; or nil
}
```

An Ident node represents an identifier.

#### func NewIdent 

``` go 
func NewIdent(name string) *Ident
```

NewIdent creates a new Ident without position. Useful for ASTs generated by code other than the Go parser.

#### (*Ident) End 

``` go 
func (x *Ident) End() token.Pos
```

#### (*Ident) IsExported 

``` go 
func (id *Ident) IsExported() bool
```

IsExported reports whether id starts with an upper-case letter.

#### (*Ident) Pos 

``` go 
func (x *Ident) Pos() token.Pos
```

#### (*Ident) String 

``` go 
func (id *Ident) String() string
```

### type IfStmt 

``` go 
type IfStmt struct {
	If   token.Pos // position of "if" keyword
	Init Stmt      // initialization statement; or nil
	Cond Expr      // condition
	Body *BlockStmt
	Else Stmt // else branch; or nil
}
```

An IfStmt node represents an if statement.

#### (*IfStmt) End 

``` go 
func (s *IfStmt) End() token.Pos
```

#### (*IfStmt) Pos 

``` go 
func (s *IfStmt) Pos() token.Pos
```

### type ImportSpec 

``` go 
type ImportSpec struct {
	Doc     *CommentGroup // associated documentation; or nil
	Name    *Ident        // local package name (including "."); or nil
	Path    *BasicLit     // import path
	Comment *CommentGroup // line comments; or nil
	EndPos  token.Pos     // end of spec (overrides Path.Pos if nonzero)
}
```

An ImportSpec node represents a single package import.

#### (*ImportSpec) End 

``` go 
func (s *ImportSpec) End() token.Pos
```

#### (*ImportSpec) Pos 

``` go 
func (s *ImportSpec) Pos() token.Pos
```

### type Importer 

``` go 
type Importer func(imports map[string]*Object, path string) (pkg *Object, err error)
```

An Importer resolves import paths to package Objects. The imports map records the packages already imported, indexed by package id (canonical import path). An Importer must determine the canonical import path and check the map to see if it is already present in the imports map. If so, the Importer can return the map entry. Otherwise, the Importer should load the package data for the given path into a new *Object (pkg), record pkg in the imports map, and then return pkg.

### type IncDecStmt 

``` go 
type IncDecStmt struct {
	X      Expr
	TokPos token.Pos   // position of Tok
	Tok    token.Token // INC or DEC
}
```

An IncDecStmt node represents an increment or decrement statement.

#### (*IncDecStmt) End 

``` go 
func (s *IncDecStmt) End() token.Pos
```

#### (*IncDecStmt) Pos 

``` go 
func (s *IncDecStmt) Pos() token.Pos
```

### type IndexExpr 

``` go 
type IndexExpr struct {
	X      Expr      // expression
	Lbrack token.Pos // position of "["
	Index  Expr      // index expression
	Rbrack token.Pos // position of "]"
}
```

An IndexExpr node represents an expression followed by an index.

#### (*IndexExpr) End 

``` go 
func (x *IndexExpr) End() token.Pos
```

#### (*IndexExpr) Pos 

``` go 
func (x *IndexExpr) Pos() token.Pos
```

### type IndexListExpr  <- go1.18

``` go 
type IndexListExpr struct {
	X       Expr      // expression
	Lbrack  token.Pos // position of "["
	Indices []Expr    // index expressions
	Rbrack  token.Pos // position of "]"
}
```

An IndexListExpr node represents an expression followed by multiple indices.

#### (*IndexListExpr) End  <- go1.18

``` go 
func (x *IndexListExpr) End() token.Pos
```

#### (*IndexListExpr) Pos  <- go1.18

``` go 
func (x *IndexListExpr) Pos() token.Pos
```

### type InterfaceType 

``` go 
type InterfaceType struct {
	Interface  token.Pos  // position of "interface" keyword
	Methods    *FieldList // list of embedded interfaces, methods, or types
	Incomplete bool       // true if (source) methods or types are missing in the Methods list
}
```

An InterfaceType node represents an interface type.

#### (*InterfaceType) End 

``` go 
func (x *InterfaceType) End() token.Pos
```

#### (*InterfaceType) Pos 

``` go 
func (x *InterfaceType) Pos() token.Pos
```

### type KeyValueExpr 

``` go 
type KeyValueExpr struct {
	Key   Expr
	Colon token.Pos // position of ":"
	Value Expr
}
```

A KeyValueExpr node represents (key : value) pairs in composite literals.

#### (*KeyValueExpr) End 

``` go 
func (x *KeyValueExpr) End() token.Pos
```

#### (*KeyValueExpr) Pos 

``` go 
func (x *KeyValueExpr) Pos() token.Pos
```

### type LabeledStmt 

``` go 
type LabeledStmt struct {
	Label *Ident
	Colon token.Pos // position of ":"
	Stmt  Stmt
}
```

A LabeledStmt node represents a labeled statement.

#### (*LabeledStmt) End 

``` go 
func (s *LabeledStmt) End() token.Pos
```

#### (*LabeledStmt) Pos 

``` go 
func (s *LabeledStmt) Pos() token.Pos
```

### type MapType 

``` go 
type MapType struct {
	Map   token.Pos // position of "map" keyword
	Key   Expr
	Value Expr
}
```

A MapType node represents a map type.

#### (*MapType) End 

``` go 
func (x *MapType) End() token.Pos
```

#### (*MapType) Pos 

``` go 
func (x *MapType) Pos() token.Pos
```

### type MergeMode 

``` go 
type MergeMode uint
```

The MergeMode flags control the behavior of MergePackageFiles.

``` go 
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

``` go 
type Node interface {
	Pos() token.Pos // position of first character belonging to the node
	End() token.Pos // position of first character immediately after the node
}
```

All node types implement the Node interface.

### type ObjKind 

``` go 
type ObjKind int
```

ObjKind describes what an object represents.

``` go 
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

#### (ObjKind) String 

``` go 
func (kind ObjKind) String() string
```

### type Object 

``` go 
type Object struct {
	Kind ObjKind
	Name string // declared name
	Decl any    // corresponding Field, XxxSpec, FuncDecl, LabeledStmt, AssignStmt, Scope; or nil
	Data any    // object-specific data; or nil
	Type any    // placeholder for type information; may be nil
}
```

An Object describes a named language entity such as a package, constant, type, variable, function (incl. methods), or label.

The Data fields contains object-specific data:

```
Kind    Data type         Data value
Pkg     *Scope            package scope
Con     int               iota for the respective declaration
```

#### func NewObj 

``` go 
func NewObj(kind ObjKind, name string) *Object
```

NewObj creates a new object of a given kind and name.

#### (*Object) Pos 

``` go 
func (obj *Object) Pos() token.Pos
```

Pos computes the source position of the declaration of an object name. The result may be an invalid position if it cannot be computed (obj.Decl may be nil or not correct).

### type Package 

``` go 
type Package struct {
	Name    string             // package name
	Scope   *Scope             // package scope across all files
	Imports map[string]*Object // map of package id -> package object
	Files   map[string]*File   // Go source files by filename
}
```

A Package node represents a set of source files collectively building a Go package.

#### func NewPackage 

``` go 
func NewPackage(fset *token.FileSet, files map[string]*File, importer Importer, universe *Scope) (*Package, error)
```

NewPackage creates a new Package node from a set of File nodes. It resolves unresolved identifiers across files and updates each file's Unresolved list accordingly. If a non-nil importer and universe scope are provided, they are used to resolve identifiers not declared in any of the package files. Any remaining unresolved identifiers are reported as undeclared. If the files belong to different packages, one package name is selected and files with different package names are reported and then ignored. The result is a package node and a scanner.ErrorList if there were errors.

#### (*Package) End 

``` go 
func (p *Package) End() token.Pos
```

#### (*Package) Pos 

``` go 
func (p *Package) Pos() token.Pos
```

### type ParenExpr 

``` go 
type ParenExpr struct {
	Lparen token.Pos // position of "("
	X      Expr      // parenthesized expression
	Rparen token.Pos // position of ")"
}
```

A ParenExpr node represents a parenthesized expression.

#### (*ParenExpr) End 

``` go 
func (x *ParenExpr) End() token.Pos
```

#### (*ParenExpr) Pos 

``` go 
func (x *ParenExpr) Pos() token.Pos
```

### type RangeStmt 

``` go 
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

#### (*RangeStmt) End 

``` go 
func (s *RangeStmt) End() token.Pos
```

#### (*RangeStmt) Pos 

``` go 
func (s *RangeStmt) Pos() token.Pos
```

### type ReturnStmt 

``` go 
type ReturnStmt struct {
	Return  token.Pos // position of "return" keyword
	Results []Expr    // result expressions; or nil
}
```

A ReturnStmt node represents a return statement.

#### (*ReturnStmt) End 

``` go 
func (s *ReturnStmt) End() token.Pos
```

#### (*ReturnStmt) Pos 

``` go 
func (s *ReturnStmt) Pos() token.Pos
```

### type Scope 

``` go 
type Scope struct {
	Outer   *Scope
	Objects map[string]*Object
}
```

A Scope maintains the set of named language entities declared in the scope and a link to the immediately surrounding (outer) scope.

#### func NewScope 

``` go 
func NewScope(outer *Scope) *Scope
```

NewScope creates a new scope nested in the outer scope.

#### (*Scope) Insert 

``` go 
func (s *Scope) Insert(obj *Object) (alt *Object)
```

Insert attempts to insert a named object obj into the scope s. If the scope already contains an object alt with the same name, Insert leaves the scope unchanged and returns alt. Otherwise it inserts obj and returns nil.

#### (*Scope) Lookup 

``` go 
func (s *Scope) Lookup(name string) *Object
```

Lookup returns the object with the given name if it is found in scope s, otherwise it returns nil. Outer scopes are ignored.

#### (*Scope) String 

``` go 
func (s *Scope) String() string
```

Debugging support

### type SelectStmt 

``` go 
type SelectStmt struct {
	Select token.Pos  // position of "select" keyword
	Body   *BlockStmt // CommClauses only
}
```

A SelectStmt node represents a select statement.

#### (*SelectStmt) End 

``` go 
func (s *SelectStmt) End() token.Pos
```

#### (*SelectStmt) Pos 

``` go 
func (s *SelectStmt) Pos() token.Pos
```

### type SelectorExpr 

``` go 
type SelectorExpr struct {
	X   Expr   // expression
	Sel *Ident // field selector
}
```

A SelectorExpr node represents an expression followed by a selector.

#### (*SelectorExpr) End 

``` go 
func (x *SelectorExpr) End() token.Pos
```

#### (*SelectorExpr) Pos 

``` go 
func (x *SelectorExpr) Pos() token.Pos
```

### type SendStmt 

``` go 
type SendStmt struct {
	Chan  Expr
	Arrow token.Pos // position of "<-"
	Value Expr
}
```

A SendStmt node represents a send statement.

#### (*SendStmt) End 

``` go 
func (s *SendStmt) End() token.Pos
```

#### (*SendStmt) Pos 

``` go 
func (s *SendStmt) Pos() token.Pos
```

### type SliceExpr 

``` go 
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

#### (*SliceExpr) End 

``` go 
func (x *SliceExpr) End() token.Pos
```

#### (*SliceExpr) Pos 

``` go 
func (x *SliceExpr) Pos() token.Pos
```

### type Spec 

``` go 
type Spec interface {
	Node
	// contains filtered or unexported methods
}
```

The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.

### type StarExpr 

``` go 
type StarExpr struct {
	Star token.Pos // position of "*"
	X    Expr      // operand
}
```

A StarExpr node represents an expression of the form "*" Expression. Semantically it could be a unary "*" expression, or a pointer type.

#### (*StarExpr) End 

``` go 
func (x *StarExpr) End() token.Pos
```

#### (*StarExpr) Pos 

``` go 
func (x *StarExpr) Pos() token.Pos
```

### type Stmt 

``` go 
type Stmt interface {
	Node
	// contains filtered or unexported methods
}
```

All statement nodes implement the Stmt interface.

### type StructType 

``` go 
type StructType struct {
	Struct     token.Pos  // position of "struct" keyword
	Fields     *FieldList // list of field declarations
	Incomplete bool       // true if (source) fields are missing in the Fields list
}
```

A StructType node represents a struct type.

#### (*StructType) End 

``` go 
func (x *StructType) End() token.Pos
```

#### (*StructType) Pos 

``` go 
func (x *StructType) Pos() token.Pos
```

### type SwitchStmt 

``` go 
type SwitchStmt struct {
	Switch token.Pos  // position of "switch" keyword
	Init   Stmt       // initialization statement; or nil
	Tag    Expr       // tag expression; or nil
	Body   *BlockStmt // CaseClauses only
}
```

A SwitchStmt node represents an expression switch statement.

#### (*SwitchStmt) End 

``` go 
func (s *SwitchStmt) End() token.Pos
```

#### (*SwitchStmt) Pos 

``` go 
func (s *SwitchStmt) Pos() token.Pos
```

### type TypeAssertExpr 

``` go 
type TypeAssertExpr struct {
	X      Expr      // expression
	Lparen token.Pos // position of "("
	Type   Expr      // asserted type; nil means type switch X.(type)
	Rparen token.Pos // position of ")"
}
```

A TypeAssertExpr node represents an expression followed by a type assertion.

#### (*TypeAssertExpr) End 

``` go 
func (x *TypeAssertExpr) End() token.Pos
```

#### (*TypeAssertExpr) Pos 

``` go 
func (x *TypeAssertExpr) Pos() token.Pos
```

### type TypeSpec 

``` go 
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

#### (*TypeSpec) End 

``` go 
func (s *TypeSpec) End() token.Pos
```

#### (*TypeSpec) Pos 

``` go 
func (s *TypeSpec) Pos() token.Pos
```

### type TypeSwitchStmt 

``` go 
type TypeSwitchStmt struct {
	Switch token.Pos  // position of "switch" keyword
	Init   Stmt       // initialization statement; or nil
	Assign Stmt       // x := y.(type) or y.(type)
	Body   *BlockStmt // CaseClauses only
}
```

A TypeSwitchStmt node represents a type switch statement.

#### (*TypeSwitchStmt) End 

``` go 
func (s *TypeSwitchStmt) End() token.Pos
```

#### (*TypeSwitchStmt) Pos 

``` go 
func (s *TypeSwitchStmt) Pos() token.Pos
```

### type UnaryExpr 

``` go 
type UnaryExpr struct {
	OpPos token.Pos   // position of Op
	Op    token.Token // operator
	X     Expr        // operand
}
```

A UnaryExpr node represents a unary expression. Unary "*" expressions are represented via StarExpr nodes.

#### (*UnaryExpr) End 

``` go 
func (x *UnaryExpr) End() token.Pos
```

#### (*UnaryExpr) Pos 

``` go 
func (x *UnaryExpr) Pos() token.Pos
```

### type ValueSpec 

``` go 
type ValueSpec struct {
	Doc     *CommentGroup // associated documentation; or nil
	Names   []*Ident      // value names (len(Names) > 0)
	Type    Expr          // value type; or nil
	Values  []Expr        // initial values; or nil
	Comment *CommentGroup // line comments; or nil
}
```

A ValueSpec node represents a constant or variable declaration (ConstSpec or VarSpec production).

#### (*ValueSpec) End 

``` go 
func (s *ValueSpec) End() token.Pos
```

#### (*ValueSpec) Pos 

``` go 
func (s *ValueSpec) Pos() token.Pos
```

### type Visitor 

``` go 
type Visitor interface {
	Visit(node Node) (w Visitor)
}
```

A Visitor's Visit method is invoked for each node encountered by Walk. If the result visitor w is not nil, Walk visits each of the children of node with the visitor w, followed by a call of w.Visit(nil).