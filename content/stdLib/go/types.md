+++
title = "types"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# types

https://pkg.go.dev/go/types@go1.20.1



Package types declares the data types and implements the algorithms for type-checking of Go packages. Use Config.Check to invoke the type checker for a package. Alternatively, create a new type checker with NewChecker and invoke it incrementally by calling Checker.Files.

Type-checking consists of several interdependent phases:

Name resolution maps each identifier (ast.Ident) in the program to the language object (Object) it denotes. Use Info.{Defs,Uses,Implicits} for the results of name resolution.

Constant folding computes the exact constant value (constant.Value) for every expression (ast.Expr) that is a compile-time constant. Use Info.Types[expr].Value for the results of constant folding.

Type inference computes the type (Type) of every expression (ast.Expr) and checks for compliance with the language specification. Use Info.Types[expr].Type for the results of type inference.

For a tutorial, see https://golang.org/s/types-tutorial.







## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/go/types/universe.go;l=38)

``` go 
var Typ = []*Basic{
	Invalid: {Invalid, 0, "invalid type"},

	Bool:          {Bool, IsBoolean, "bool"},
	Int:           {Int, IsInteger, "int"},
	Int8:          {Int8, IsInteger, "int8"},
	Int16:         {Int16, IsInteger, "int16"},
	Int32:         {Int32, IsInteger, "int32"},
	Int64:         {Int64, IsInteger, "int64"},
	Uint:          {Uint, IsInteger | IsUnsigned, "uint"},
	Uint8:         {Uint8, IsInteger | IsUnsigned, "uint8"},
	Uint16:        {Uint16, IsInteger | IsUnsigned, "uint16"},
	Uint32:        {Uint32, IsInteger | IsUnsigned, "uint32"},
	Uint64:        {Uint64, IsInteger | IsUnsigned, "uint64"},
	Uintptr:       {Uintptr, IsInteger | IsUnsigned, "uintptr"},
	Float32:       {Float32, IsFloat, "float32"},
	Float64:       {Float64, IsFloat, "float64"},
	Complex64:     {Complex64, IsComplex, "complex64"},
	Complex128:    {Complex128, IsComplex, "complex128"},
	String:        {String, IsString, "string"},
	UnsafePointer: {UnsafePointer, 0, "Pointer"},

	UntypedBool:    {UntypedBool, IsBoolean | IsUntyped, "untyped bool"},
	UntypedInt:     {UntypedInt, IsInteger | IsUntyped, "untyped int"},
	UntypedRune:    {UntypedRune, IsInteger | IsUntyped, "untyped rune"},
	UntypedFloat:   {UntypedFloat, IsFloat | IsUntyped, "untyped float"},
	UntypedComplex: {UntypedComplex, IsComplex | IsUntyped, "untyped complex"},
	UntypedString:  {UntypedString, IsString | IsUntyped, "untyped string"},
	UntypedNil:     {UntypedNil, IsUntyped, "untyped nil"},
}
```

Typ contains the predeclared *Basic types indexed by their corresponding BasicKind.

The *Basic type for Typ[Byte] will have the name "uint8". Use Universe.Lookup("byte").Type() to obtain the specific alias basic type named "byte" (and analogous for "rune").

## 函数

#### func AssertableTo 

``` go 
func AssertableTo(V *Interface, T Type) bool
```

AssertableTo reports whether a value of type V can be asserted to have type T.

The behavior of AssertableTo is unspecified in three cases:

- if T is Typ[Invalid]
- if V is a generalized interface; i.e., an interface that may only be used as a type constraint in Go code
- if T is an uninstantiated generic type

#### func AssignableTo 

``` go 
func AssignableTo(V, T Type) bool
```

AssignableTo reports whether a value of type V is assignable to a variable of type T.

The behavior of AssignableTo is unspecified if V or T is Typ[Invalid] or an uninstantiated generic type.

#### func CheckExpr  <- go1.13

``` go 
func CheckExpr(fset *token.FileSet, pkg *Package, pos token.Pos, expr ast.Expr, info *Info) (err error)
```

CheckExpr type checks the expression expr as if it had appeared at position pos of package pkg. Type information about the expression is recorded in info. The expression may be an identifier denoting an uninstantiated generic function or type.

If pkg == nil, the Universe scope is used and the provided position pos is ignored. If pkg != nil, and pos is invalid, the package scope is used. Otherwise, pos must belong to the package.

An error is returned if pos is not within the package or if the node cannot be type-checked.

Note: Eval and CheckExpr should not be used instead of running Check to compute types and values, but in addition to Check, as these functions ignore the context in which an expression is used (e.g., an assignment). Thus, top-level untyped constants will return an untyped type rather then the respective context-specific type.

#### func Comparable 

``` go 
func Comparable(T Type) bool
```

Comparable reports whether values of type T are comparable.

#### func ConvertibleTo 

``` go 
func ConvertibleTo(V, T Type) bool
```

ConvertibleTo reports whether a value of type V is convertible to a value of type T.

The behavior of ConvertibleTo is unspecified if V or T is Typ[Invalid] or an uninstantiated generic type.

#### func DefPredeclaredTestFuncs 

``` go 
func DefPredeclaredTestFuncs()
```

DefPredeclaredTestFuncs defines the assert and trace built-ins. These built-ins are intended for debugging and testing of this package only.

#### func ExprString 

``` go 
func ExprString(x ast.Expr) string
```

ExprString returns the (possibly shortened) string representation for x. Shortened representations are suitable for user interfaces but may not necessarily follow Go syntax.

#### func Id 

``` go 
func Id(pkg *Package, name string) string
```

Id returns name if it is exported, otherwise it returns the name qualified with the package path.

#### func Identical 

``` go 
func Identical(x, y Type) bool
```

Identical reports whether x and y are identical types. Receivers of Signature types are ignored.

#### func IdenticalIgnoreTags  <- go1.8

``` go 
func IdenticalIgnoreTags(x, y Type) bool
```

IdenticalIgnoreTags reports whether x and y are identical types if tags are ignored. Receivers of Signature types are ignored.

#### func Implements 

``` go 
func Implements(V Type, T *Interface) bool
```

Implements reports whether type V implements interface T.

The behavior of Implements is unspecified if V is Typ[Invalid] or an uninstantiated generic type.

#### func IsInterface 

``` go 
func IsInterface(t Type) bool
```

IsInterface reports whether t is an interface type.

#### func ObjectString 

``` go 
func ObjectString(obj Object, qf Qualifier) string
```

ObjectString returns the string form of obj. The Qualifier controls the printing of package-level objects, and may be nil.

#### func Satisfies  <- go1.20

``` go 
func Satisfies(V Type, T *Interface) bool
```

Satisfies reports whether type V satisfies the constraint T.

The behavior of Satisfies is unspecified if V is Typ[Invalid] or an uninstantiated generic type.

#### func SelectionString 

``` go 
func SelectionString(s *Selection, qf Qualifier) string
```

SelectionString returns the string form of s. The Qualifier controls the printing of package-level objects, and may be nil.

Examples:

```
"field (T) f int"
"method (T) f(X) Y"
"method expr (T) f(X) Y"
```

#### func TypeString 

``` go 
func TypeString(typ Type, qf Qualifier) string
```

TypeString returns the string representation of typ. The Qualifier controls the printing of package-level objects, and may be nil.

#### func WriteExpr 

``` go 
func WriteExpr(buf *bytes.Buffer, x ast.Expr)
```

WriteExpr writes the (possibly shortened) string representation for x to buf. Shortened representations are suitable for user interfaces but may not necessarily follow Go syntax.

#### func WriteSignature 

``` go 
func WriteSignature(buf *bytes.Buffer, sig *Signature, qf Qualifier)
```

WriteSignature writes the representation of the signature sig to buf, without a leading "func" keyword. The Qualifier controls the printing of package-level objects, and may be nil.

#### func WriteType 

``` go 
func WriteType(buf *bytes.Buffer, typ Type, qf Qualifier)
```

WriteType writes the string representation of typ to buf. The Qualifier controls the printing of package-level objects, and may be nil.

## 类型

### type ArgumentError  <- go1.18

``` go 
type ArgumentError struct {
	Index int
	Err   error
}
```

An ArgumentError holds an error associated with an argument index.

#### (*ArgumentError) Error  <- go1.18

``` go 
func (e *ArgumentError) Error() string
```

#### (*ArgumentError) Unwrap  <- go1.18

``` go 
func (e *ArgumentError) Unwrap() error
```

### type Array 

``` go 
type Array struct {
	// contains filtered or unexported fields
}
```

An Array represents an array type.

#### func NewArray 

``` go 
func NewArray(elem Type, len int64) *Array
```

NewArray returns a new array type for the given element type and length. A negative length indicates an unknown length.

#### (*Array) Elem 

``` go 
func (a *Array) Elem() Type
```

Elem returns element type of array a.

#### (*Array) Len 

``` go 
func (a *Array) Len() int64
```

Len returns the length of array a. A negative result indicates an unknown length.

#### (*Array) String 

``` go 
func (t *Array) String() string
```

#### (*Array) Underlying 

``` go 
func (t *Array) Underlying() Type
```

### type Basic 

``` go 
type Basic struct {
	// contains filtered or unexported fields
}
```

A Basic represents a basic type.

#### (*Basic) Info 

``` go 
func (b *Basic) Info() BasicInfo
```

Info returns information about properties of basic type b.

#### (*Basic) Kind 

``` go 
func (b *Basic) Kind() BasicKind
```

Kind returns the kind of basic type b.

#### (*Basic) Name 

``` go 
func (b *Basic) Name() string
```

Name returns the name of basic type b.

#### (*Basic) String 

``` go 
func (t *Basic) String() string
```

#### (*Basic) Underlying 

``` go 
func (t *Basic) Underlying() Type
```

### type BasicInfo 

``` go 
type BasicInfo int
```

BasicInfo is a set of flags describing properties of a basic type.

``` go 
const (
	IsBoolean BasicInfo = 1 << iota
	IsInteger
	IsUnsigned
	IsFloat
	IsComplex
	IsString
	IsUntyped

	IsOrdered   = IsInteger | IsFloat | IsString
	IsNumeric   = IsInteger | IsFloat | IsComplex
	IsConstType = IsBoolean | IsNumeric | IsString
)
```

Properties of basic types.

### type BasicKind 

``` go 
type BasicKind int
```

BasicKind describes the kind of basic type.

``` go 
const (
	Invalid BasicKind = iota // type is invalid

	// predeclared types
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	String
	UnsafePointer

	// types for untyped values
	UntypedBool
	UntypedInt
	UntypedRune
	UntypedFloat
	UntypedComplex
	UntypedString
	UntypedNil

	// aliases
	Byte = Uint8
	Rune = Int32
)
```

### type Builtin 

``` go 
type Builtin struct {
	// contains filtered or unexported fields
}
```

A Builtin represents a built-in function. Builtins don't have a valid type.

#### (*Builtin) Exported 

``` go 
func (obj *Builtin) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn't take into account whether the object is in a local (function) scope or not.

#### (*Builtin) Id 

``` go 
func (obj *Builtin) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

#### (*Builtin) Name 

``` go 
func (obj *Builtin) Name() string
```

Name returns the object's (package-local, unqualified) name.

#### (*Builtin) Parent 

``` go 
func (obj *Builtin) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

#### (*Builtin) Pkg 

``` go 
func (obj *Builtin) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

#### (*Builtin) Pos 

``` go 
func (obj *Builtin) Pos() token.Pos
```

Pos returns the declaration position of the object's identifier.

#### (*Builtin) String 

``` go 
func (obj *Builtin) String() string
```

#### (*Builtin) Type 

``` go 
func (obj *Builtin) Type() Type
```

Type returns the object's type.

### type Chan 

``` go 
type Chan struct {
	// contains filtered or unexported fields
}
```

A Chan represents a channel type.

#### func NewChan 

``` go 
func NewChan(dir ChanDir, elem Type) *Chan
```

NewChan returns a new channel type for the given direction and element type.

#### (*Chan) Dir 

``` go 
func (c *Chan) Dir() ChanDir
```

Dir returns the direction of channel c.

#### (*Chan) Elem 

``` go 
func (c *Chan) Elem() Type
```

Elem returns the element type of channel c.

#### (*Chan) String 

``` go 
func (t *Chan) String() string
```

#### (*Chan) Underlying 

``` go 
func (t *Chan) Underlying() Type
```

### type ChanDir 

``` go 
type ChanDir int
```

A ChanDir value indicates a channel direction.

``` go 
const (
	SendRecv ChanDir = iota
	SendOnly
	RecvOnly
)
```

The direction of a channel is indicated by one of these constants.

### type Checker 

``` go 
type Checker struct {
	*Info
	// contains filtered or unexported fields
}
```

A Checker maintains the state of the type checker. It must be created with NewChecker.

#### func NewChecker 

``` go 
func NewChecker(conf *Config, fset *token.FileSet, pkg *Package, info *Info) *Checker
```

NewChecker returns a new Checker instance for a given package. Package files may be added incrementally via checker.Files.

#### (*Checker) Files 

``` go 
func (check *Checker) Files(files []*ast.File) error
```

Files checks the provided files as part of the checker's package.

### type Config 

``` go 
type Config struct {
	// Context is the context used for resolving global identifiers. If nil, the
	// type checker will initialize this field with a newly created context.
	Context *Context

	// GoVersion describes the accepted Go language version. The string
	// must follow the format "go%d.%d" (e.g. "go1.12") or it must be
	// empty; an empty string indicates the latest language version.
	// If the format is invalid, invoking the type checker will cause a
	// panic.
	GoVersion string

	// If IgnoreFuncBodies is set, function bodies are not
	// type-checked.
	IgnoreFuncBodies bool

	// If FakeImportC is set, `import "C"` (for packages requiring Cgo)
	// declares an empty "C" package and errors are omitted for qualified
	// identifiers referring to package C (which won't find an object).
	// This feature is intended for the standard library cmd/api tool.
	//
	// Caution: Effects may be unpredictable due to follow-on errors.
	//          Do not use casually!
	FakeImportC bool

	// If Error != nil, it is called with each error found
	// during type checking; err has dynamic type Error.
	// Secondary errors (for instance, to enumerate all types
	// involved in an invalid recursive type declaration) have
	// error strings that start with a '\t' character.
	// If Error == nil, type-checking stops with the first
	// error found.
	Error func(err error)

	// An importer is used to import packages referred to from
	// import declarations.
	// If the installed importer implements ImporterFrom, the type
	// checker calls ImportFrom instead of Import.
	// The type checker reports an error if an importer is needed
	// but none was installed.
	Importer Importer

	// If Sizes != nil, it provides the sizing functions for package unsafe.
	// Otherwise SizesFor("gc", "amd64") is used instead.
	Sizes Sizes

	// If DisableUnusedImportCheck is set, packages are not checked
	// for unused imports.
	DisableUnusedImportCheck bool
	// contains filtered or unexported fields
}
```

A Config specifies the configuration for type checking. The zero value for Config is a ready-to-use default configuration.

#### (*Config) Check 

``` go 
func (conf *Config) Check(path string, fset *token.FileSet, files []*ast.File, info *Info) (*Package, error)
```

Check type-checks a package and returns the resulting package object and the first error if any. Additionally, if info != nil, Check populates each of the non-nil maps in the Info struct.

The package is marked as complete if no errors occurred, otherwise it is incomplete. See Config.Error for controlling behavior in the presence of errors.

The package is specified by a list of *ast.Files and corresponding file set, and the package path the package is identified with. The clean path must not be empty or dot (".").

### type Const 

``` go 
type Const struct {
	// contains filtered or unexported fields
}
```

A Const represents a declared constant.

#### func NewConst 

``` go 
func NewConst(pos token.Pos, pkg *Package, name string, typ Type, val constant.Value) *Const
```

NewConst returns a new constant with value val. The remaining arguments set the attributes found with all Objects.

#### (*Const) Exported 

``` go 
func (obj *Const) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn't take into account whether the object is in a local (function) scope or not.

#### (*Const) Id 

``` go 
func (obj *Const) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

#### (*Const) Name 

``` go 
func (obj *Const) Name() string
```

Name returns the object's (package-local, unqualified) name.

#### (*Const) Parent 

``` go 
func (obj *Const) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

#### (*Const) Pkg 

``` go 
func (obj *Const) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

#### (*Const) Pos 

``` go 
func (obj *Const) Pos() token.Pos
```

Pos returns the declaration position of the object's identifier.

#### (*Const) String 

``` go 
func (obj *Const) String() string
```

#### (*Const) Type 

``` go 
func (obj *Const) Type() Type
```

Type returns the object's type.

#### (*Const) Val 

``` go 
func (obj *Const) Val() constant.Value
```

Val returns the constant's value.

### type Context  <- go1.18

``` go 
type Context struct {
	// contains filtered or unexported fields
}
```

A Context is an opaque type checking context. It may be used to share identical type instances across type-checked packages or calls to Instantiate. Contexts are safe for concurrent use.

The use of a shared context does not guarantee that identical instances are deduplicated in all cases.

#### func NewContext  <- go1.18

``` go 
func NewContext() *Context
```

NewContext creates a new Context.

### type Error 

``` go 
type Error struct {
	Fset *token.FileSet // file set for interpretation of Pos
	Pos  token.Pos      // error position
	Msg  string         // error message
	Soft bool           // if set, error is "soft"
	// contains filtered or unexported fields
}
```

An Error describes a type-checking error; it implements the error interface. A "soft" error is an error that still permits a valid interpretation of a package (such as "unused variable"); "hard" errors may lead to unpredictable behavior if ignored.

#### (Error) Error 

``` go 
func (err Error) Error() string
```

Error returns an error string formatted as follows: filename:line:column: message

### type Func 

``` go 
type Func struct {
	// contains filtered or unexported fields
}
```

A Func represents a declared function, concrete method, or abstract (interface) method. Its Type() is always a *Signature. An abstract method may belong to many interfaces due to embedding.

#### func MissingMethod 

``` go 
func MissingMethod(V Type, T *Interface, static bool) (method *Func, wrongType bool)
```

MissingMethod returns (nil, false) if V implements T, otherwise it returns a missing method required by T and whether it is missing or just has the wrong type.

For non-interface types V, or if static is set, V implements T if all methods of T are present in V. Otherwise (V is an interface and static is not set), MissingMethod only checks that methods of T which are also present in V have matching types (e.g., for a type assertion x.(T) where x is of interface type V).

#### func NewFunc 

``` go 
func NewFunc(pos token.Pos, pkg *Package, name string, sig *Signature) *Func
```

NewFunc returns a new function with the given signature, representing the function's type.

#### (*Func) Exported 

``` go 
func (obj *Func) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn't take into account whether the object is in a local (function) scope or not.

#### (*Func) FullName 

``` go 
func (obj *Func) FullName() string
```

FullName returns the package- or receiver-type-qualified name of function or method obj.

#### (*Func) Id 

``` go 
func (obj *Func) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

#### (*Func) Name 

``` go 
func (obj *Func) Name() string
```

Name returns the object's (package-local, unqualified) name.

#### (*Func) Origin  <- go1.19

``` go 
func (obj *Func) Origin() *Func
```

Origin returns the canonical Func for its receiver, i.e. the Func object recorded in Info.Defs.

For synthetic functions created during instantiation (such as methods on an instantiated Named type or interface methods that depend on type arguments), this will be the corresponding Func on the generic (uninstantiated) type. For all other Funcs Origin returns the receiver.

#### (*Func) Parent 

``` go 
func (obj *Func) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

#### (*Func) Pkg 

``` go 
func (obj *Func) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

#### (*Func) Pos 

``` go 
func (obj *Func) Pos() token.Pos
```

Pos returns the declaration position of the object's identifier.

#### (*Func) Scope 

``` go 
func (obj *Func) Scope() *Scope
```

Scope returns the scope of the function's body block. The result is nil for imported or instantiated functions and methods (but there is also no mechanism to get to an instantiated function).

#### (*Func) String 

``` go 
func (obj *Func) String() string
```

#### (*Func) Type 

``` go 
func (obj *Func) Type() Type
```

Type returns the object's type.

### type ImportMode  <- go1.6

``` go 
type ImportMode int
```

ImportMode is reserved for future use.

### type Importer 

``` go 
type Importer interface {
	// Import returns the imported package for the given import path.
	// The semantics is like for ImporterFrom.ImportFrom except that
	// dir and mode are ignored (since they are not present).
	Import(path string) (*Package, error)
}
```

An Importer resolves import paths to Packages.

CAUTION: This interface does not support the import of locally vendored packages. See https://golang.org/s/go15vendor. If possible, external implementations should implement ImporterFrom.

### type ImporterFrom  <- go1.6

``` go 
type ImporterFrom interface {
	// Importer is present for backward-compatibility. Calling
	// Import(path) is the same as calling ImportFrom(path, "", 0);
	// i.e., locally vendored packages may not be found.
	// The types package does not call Import if an ImporterFrom
	// is present.
	Importer

	// ImportFrom returns the imported package for the given import
	// path when imported by a package file located in dir.
	// If the import failed, besides returning an error, ImportFrom
	// is encouraged to cache and return a package anyway, if one
	// was created. This will reduce package inconsistencies and
	// follow-on type checker errors due to the missing package.
	// The mode value must be 0; it is reserved for future use.
	// Two calls to ImportFrom with the same path and dir must
	// return the same package.
	ImportFrom(path, dir string, mode ImportMode) (*Package, error)
}
```

An ImporterFrom resolves import paths to packages; it supports vendoring per https://golang.org/s/go15vendor. Use go/importer to obtain an ImporterFrom implementation.

### type Info 

``` go 
type Info struct {
	// Types maps expressions to their types, and for constant
	// expressions, also their values. Invalid expressions are
	// omitted.
	//
	// For (possibly parenthesized) identifiers denoting built-in
	// functions, the recorded signatures are call-site specific:
	// if the call result is not a constant, the recorded type is
	// an argument-specific signature. Otherwise, the recorded type
	// is invalid.
	//
	// The Types map does not record the type of every identifier,
	// only those that appear where an arbitrary expression is
	// permitted. For instance, the identifier f in a selector
	// expression x.f is found only in the Selections map, the
	// identifier z in a variable declaration 'var z int' is found
	// only in the Defs map, and identifiers denoting packages in
	// qualified identifiers are collected in the Uses map.
	Types map[ast.Expr]TypeAndValue

	// Instances maps identifiers denoting generic types or functions to their
	// type arguments and instantiated type.
	//
	// For example, Instances will map the identifier for 'T' in the type
	// instantiation T[int, string] to the type arguments [int, string] and
	// resulting instantiated *Named type. Given a generic function
	// func F[A any](A), Instances will map the identifier for 'F' in the call
	// expression F(int(1)) to the inferred type arguments [int], and resulting
	// instantiated *Signature.
	//
	// Invariant: Instantiating Uses[id].Type() with Instances[id].TypeArgs
	// results in an equivalent of Instances[id].Type.
	Instances map[*ast.Ident]Instance

	// Defs maps identifiers to the objects they define (including
	// package names, dots "." of dot-imports, and blank "_" identifiers).
	// For identifiers that do not denote objects (e.g., the package name
	// in package clauses, or symbolic variables t in t := x.(type) of
	// type switch headers), the corresponding objects are nil.
	//
	// For an embedded field, Defs returns the field *Var it defines.
	//
	// Invariant: Defs[id] == nil || Defs[id].Pos() == id.Pos()
	Defs map[*ast.Ident]Object

	// Uses maps identifiers to the objects they denote.
	//
	// For an embedded field, Uses returns the *TypeName it denotes.
	//
	// Invariant: Uses[id].Pos() != id.Pos()
	Uses map[*ast.Ident]Object

	// Implicits maps nodes to their implicitly declared objects, if any.
	// The following node and object types may appear:
	//
	//     node               declared object
	//
	//     *ast.ImportSpec    *PkgName for imports without renames
	//     *ast.CaseClause    type-specific *Var for each type switch case clause (incl. default)
	//     *ast.Field         anonymous parameter *Var (incl. unnamed results)
	//
	Implicits map[ast.Node]Object

	// Selections maps selector expressions (excluding qualified identifiers)
	// to their corresponding selections.
	Selections map[*ast.SelectorExpr]*Selection

	// Scopes maps ast.Nodes to the scopes they define. Package scopes are not
	// associated with a specific node but with all files belonging to a package.
	// Thus, the package scope can be found in the type-checked Package object.
	// Scopes nest, with the Universe scope being the outermost scope, enclosing
	// the package scope, which contains (one or more) files scopes, which enclose
	// function scopes which in turn enclose statement and function literal scopes.
	// Note that even though package-level functions are declared in the package
	// scope, the function scopes are embedded in the file scope of the file
	// containing the function declaration.
	//
	// The following node types may appear in Scopes:
	//
	//     *ast.File
	//     *ast.FuncType
	//     *ast.TypeSpec
	//     *ast.BlockStmt
	//     *ast.IfStmt
	//     *ast.SwitchStmt
	//     *ast.TypeSwitchStmt
	//     *ast.CaseClause
	//     *ast.CommClause
	//     *ast.ForStmt
	//     *ast.RangeStmt
	//
	Scopes map[ast.Node]*Scope

	// InitOrder is the list of package-level initializers in the order in which
	// they must be executed. Initializers referring to variables related by an
	// initialization dependency appear in topological order, the others appear
	// in source order. Variables without an initialization expression do not
	// appear in this list.
	InitOrder []*Initializer
}
```

Info holds result type information for a type-checked package. Only the information for which a map is provided is collected. If the package has type errors, the collected information may be incomplete.

##### Example
``` go 
```

#### (*Info) ObjectOf 

``` go 
func (info *Info) ObjectOf(id *ast.Ident) Object
```

ObjectOf returns the object denoted by the specified id, or nil if not found.

If id is an embedded struct field, ObjectOf returns the field (*Var) it defines, not the type (*TypeName) it uses.

Precondition: the Uses and Defs maps are populated.

#### (*Info) TypeOf 

``` go 
func (info *Info) TypeOf(e ast.Expr) Type
```

TypeOf returns the type of expression e, or nil if not found. Precondition: the Types, Uses and Defs maps are populated.

### type Initializer 

``` go 
type Initializer struct {
	Lhs []*Var // var Lhs = Rhs
	Rhs ast.Expr
}
```

An Initializer describes a package-level variable, or a list of variables in case of a multi-valued initialization expression, and the corresponding initialization expression.

#### (*Initializer) String 

``` go 
func (init *Initializer) String() string
```

### type Instance  <- go1.18

``` go 
type Instance struct {
	TypeArgs *TypeList
	Type     Type
}
```

Instance reports the type arguments and instantiated type for type and function instantiations. For type instantiations, Type will be of dynamic type *Named. For function instantiations, Type will be of dynamic type *Signature.

### type Interface 

``` go 
type Interface struct {
	// contains filtered or unexported fields
}
```

An Interface represents an interface type.

##### Example
``` go 
```

#### func NewInterfaceType  <- go1.11

``` go 
func NewInterfaceType(methods []*Func, embeddeds []Type) *Interface
```

NewInterfaceType returns a new interface for the given methods and embedded types. NewInterfaceType takes ownership of the provided methods and may modify their types by setting missing receivers.

To avoid race conditions, the interface's type set should be computed before concurrent use of the interface, by explicitly calling Complete.

#### (*Interface) Complete 

``` go 
func (t *Interface) Complete() *Interface
```

Complete computes the interface's type set. It must be called by users of NewInterfaceType and NewInterface after the interface's embedded types are fully defined and before using the interface type in any way other than to form other types. The interface must not contain duplicate methods or a panic occurs. Complete returns the receiver.

Interface types that have been completed are safe for concurrent use.

##### Example
``` go 
```

#### (*Interface) EmbeddedType  <- go1.11

``` go 
func (t *Interface) EmbeddedType(i int) Type
```

EmbeddedType returns the i'th embedded type of interface t for 0 <= i < t.NumEmbeddeds().

#### (*Interface) Empty 

``` go 
func (t *Interface) Empty() bool
```

Empty reports whether t is the empty interface.

#### (*Interface) ExplicitMethod 

``` go 
func (t *Interface) ExplicitMethod(i int) *Func
```

ExplicitMethod returns the i'th explicitly declared method of interface t for 0 <= i < t.NumExplicitMethods(). The methods are ordered by their unique Id.

#### (*Interface) IsComparable  <- go1.18

``` go 
func (t *Interface) IsComparable() bool
```

IsComparable reports whether each type in interface t's type set is comparable.

#### (*Interface) IsImplicit  <- go1.18

``` go 
func (t *Interface) IsImplicit() bool
```

IsImplicit reports whether the interface t is a wrapper for a type set literal.

#### (*Interface) IsMethodSet  <- go1.18

``` go 
func (t *Interface) IsMethodSet() bool
```

IsMethodSet reports whether the interface t is fully described by its method set.

#### (*Interface) MarkImplicit  <- go1.18

``` go 
func (t *Interface) MarkImplicit()
```

MarkImplicit marks the interface t as implicit, meaning this interface corresponds to a constraint literal such as ~T or A|B without explicit interface embedding. MarkImplicit should be called before any concurrent use of implicit interfaces.

#### (*Interface) Method 

``` go 
func (t *Interface) Method(i int) *Func
```

Method returns the i'th method of interface t for 0 <= i < t.NumMethods(). The methods are ordered by their unique Id.

#### (*Interface) NumEmbeddeds 

``` go 
func (t *Interface) NumEmbeddeds() int
```

NumEmbeddeds returns the number of embedded types in interface t.

#### (*Interface) NumExplicitMethods 

``` go 
func (t *Interface) NumExplicitMethods() int
```

NumExplicitMethods returns the number of explicitly declared methods of interface t.

#### (*Interface) NumMethods 

``` go 
func (t *Interface) NumMethods() int
```

NumMethods returns the total number of methods of interface t.

#### (*Interface) String 

``` go 
func (t *Interface) String() string
```

#### (*Interface) Underlying 

``` go 
func (t *Interface) Underlying() Type
```

### type Label 

``` go 
type Label struct {
	// contains filtered or unexported fields
}
```

A Label represents a declared label. Labels don't have a type.

#### func NewLabel 

``` go 
func NewLabel(pos token.Pos, pkg *Package, name string) *Label
```

NewLabel returns a new label.

#### (*Label) Exported 

``` go 
func (obj *Label) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn't take into account whether the object is in a local (function) scope or not.

#### (*Label) Id 

``` go 
func (obj *Label) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

#### (*Label) Name 

``` go 
func (obj *Label) Name() string
```

Name returns the object's (package-local, unqualified) name.

#### (*Label) Parent 

``` go 
func (obj *Label) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

#### (*Label) Pkg 

``` go 
func (obj *Label) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

#### (*Label) Pos 

``` go 
func (obj *Label) Pos() token.Pos
```

Pos returns the declaration position of the object's identifier.

#### (*Label) String 

``` go 
func (obj *Label) String() string
```

#### (*Label) Type 

``` go 
func (obj *Label) Type() Type
```

Type returns the object's type.

### type Map 

``` go 
type Map struct {
	// contains filtered or unexported fields
}
```

A Map represents a map type.

#### func NewMap 

``` go 
func NewMap(key, elem Type) *Map
```

NewMap returns a new map for the given key and element types.

#### (*Map) Elem 

``` go 
func (m *Map) Elem() Type
```

Elem returns the element type of map m.

#### (*Map) Key 

``` go 
func (m *Map) Key() Type
```

Key returns the key type of map m.

#### (*Map) String 

``` go 
func (t *Map) String() string
```

#### (*Map) Underlying 

``` go 
func (t *Map) Underlying() Type
```

### type MethodSet 

``` go 
type MethodSet struct {
	// contains filtered or unexported fields
}
```

A MethodSet is an ordered set of concrete or abstract (interface) methods; a method is a MethodVal selection, and they are ordered by ascending m.Obj().Id(). The zero value for a MethodSet is a ready-to-use empty method set.

##### Example
``` go 
```

#### func NewMethodSet 

``` go 
func NewMethodSet(T Type) *MethodSet
```

NewMethodSet returns the method set for the given type T. It always returns a non-nil method set, even if it is empty.

#### (*MethodSet) At 

``` go 
func (s *MethodSet) At(i int) *Selection
```

At returns the i'th method in s for 0 <= i < s.Len().

#### (*MethodSet) Len 

``` go 
func (s *MethodSet) Len() int
```

Len returns the number of methods in s.

#### (*MethodSet) Lookup 

``` go 
func (s *MethodSet) Lookup(pkg *Package, name string) *Selection
```

Lookup returns the method with matching package and name, or nil if not found.

#### (*MethodSet) String 

``` go 
func (s *MethodSet) String() string
```

### type Named 

``` go 
type Named struct {
	// contains filtered or unexported fields
}
```

A Named represents a named (defined) type.

#### func NewNamed 

``` go 
func NewNamed(obj *TypeName, underlying Type, methods []*Func) *Named
```

NewNamed returns a new named type for the given type name, underlying type, and associated methods. If the given type name obj doesn't have a type yet, its type is set to the returned named type. The underlying type must not be a *Named.

#### (*Named) AddMethod 

``` go 
func (t *Named) AddMethod(m *Func)
```

AddMethod adds method m unless it is already in the method list. t must not have type arguments.

#### (*Named) Method 

``` go 
func (t *Named) Method(i int) *Func
```

Method returns the i'th method of named type t for 0 <= i < t.NumMethods().

For an ordinary or instantiated type t, the receiver base type of this method is the named type t. For an uninstantiated generic type t, each method receiver is instantiated with its receiver type parameters.

#### (*Named) NumMethods 

``` go 
func (t *Named) NumMethods() int
```

NumMethods returns the number of explicit methods defined for t.

#### (*Named) Obj 

``` go 
func (t *Named) Obj() *TypeName
```

Obj returns the type name for the declaration defining the named type t. For instantiated types, this is same as the type name of the origin type.

#### (*Named) Origin  <- go1.18

``` go 
func (t *Named) Origin() *Named
```

Origin returns the generic type from which the named type t is instantiated. If t is not an instantiated type, the result is t.

#### (*Named) SetTypeParams  <- go1.18

``` go 
func (t *Named) SetTypeParams(tparams []*TypeParam)
```

SetTypeParams sets the type parameters of the named type t. t must not have type arguments.

#### (*Named) SetUnderlying 

``` go 
func (t *Named) SetUnderlying(underlying Type)
```

SetUnderlying sets the underlying type and marks t as complete. t must not have type arguments.

#### (*Named) String 

``` go 
func (t *Named) String() string
```

#### (*Named) TypeArgs  <- go1.18

``` go 
func (t *Named) TypeArgs() *TypeList
```

TypeArgs returns the type arguments used to instantiate the named type t.

#### (*Named) TypeParams  <- go1.18

``` go 
func (t *Named) TypeParams() *TypeParamList
```

TypeParams returns the type parameters of the named type t, or nil. The result is non-nil for an (originally) generic type even if it is instantiated.

#### (*Named) Underlying 

``` go 
func (t *Named) Underlying() Type
```

### type Nil 

``` go 
type Nil struct {
	// contains filtered or unexported fields
}
```

Nil represents the predeclared value nil.

#### (*Nil) Exported 

``` go 
func (obj *Nil) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn't take into account whether the object is in a local (function) scope or not.

#### (*Nil) Id 

``` go 
func (obj *Nil) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

#### (*Nil) Name 

``` go 
func (obj *Nil) Name() string
```

Name returns the object's (package-local, unqualified) name.

#### (*Nil) Parent 

``` go 
func (obj *Nil) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

#### (*Nil) Pkg 

``` go 
func (obj *Nil) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

#### (*Nil) Pos 

``` go 
func (obj *Nil) Pos() token.Pos
```

Pos returns the declaration position of the object's identifier.

#### (*Nil) String 

``` go 
func (obj *Nil) String() string
```

#### (*Nil) Type 

``` go 
func (obj *Nil) Type() Type
```

Type returns the object's type.

### type Object 

``` go 
type Object interface {
	Parent() *Scope // scope in which this object is declared; nil for methods and struct fields
	Pos() token.Pos // position of object identifier in declaration
	Pkg() *Package  // package to which this object belongs; nil for labels and objects in the Universe scope
	Name() string   // package local object name
	Type() Type     // object type
	Exported() bool // reports whether the name starts with a capital letter
	Id() string     // object name if exported, qualified name if not exported (see func Id)

	// String returns a human-readable string of the object.
	String() string
	// contains filtered or unexported methods
}
```

An Object describes a named language entity such as a package, constant, type, variable, function (incl. methods), or label. All objects implement the Object interface.

#### func LookupFieldOrMethod 

``` go 
func LookupFieldOrMethod(T Type, addressable bool, pkg *Package, name string) (obj Object, index []int, indirect bool)
```

LookupFieldOrMethod looks up a field or method with given package and name in T and returns the corresponding *Var or *Func, an index sequence, and a bool indicating if there were any pointer indirections on the path to the field or method. If addressable is set, T is the type of an addressable variable (only matters for method lookups). T must not be nil.

The last index entry is the field or method index in the (possibly embedded) type where the entry was found, either:

1. the list of declared methods of a named type; or
2. the list of all methods (method set) of an interface type; or
3. the list of fields of a struct type.

The earlier index entries are the indices of the embedded struct fields traversed to get to the found entry, starting at depth 0.

If no entry is found, a nil object is returned. In this case, the returned index and indirect values have the following meaning:

- If index != nil, the index sequence points to an ambiguous entry (the same name appeared more than once at the same embedding level).
- If indirect is set, a method with a pointer receiver type was found but there was no pointer on the path from the actual receiver type to the method's formal receiver base type, nor was the receiver addressable.

### type Package 

``` go 
type Package struct {
	// contains filtered or unexported fields
}
```

A Package describes a Go package.

``` go 
var Unsafe *Package
```

The Unsafe package is the package returned by an importer for the import path "unsafe".

#### func NewPackage 

``` go 
func NewPackage(path, name string) *Package
```

NewPackage returns a new Package for the given package path and name. The package is not complete and contains no explicit imports.

#### (*Package) Complete 

``` go 
func (pkg *Package) Complete() bool
```

A package is complete if its scope contains (at least) all exported objects; otherwise it is incomplete.

#### (*Package) Imports 

``` go 
func (pkg *Package) Imports() []*Package
```

Imports returns the list of packages directly imported by pkg; the list is in source order.

If pkg was loaded from export data, Imports includes packages that provide package-level objects referenced by pkg. This may be more or less than the set of packages directly imported by pkg's source code.

If pkg uses cgo and the FakeImportC configuration option was enabled, the imports list may contain a fake "C" package.

#### (*Package) MarkComplete 

``` go 
func (pkg *Package) MarkComplete()
```

MarkComplete marks a package as complete.

#### (*Package) Name 

``` go 
func (pkg *Package) Name() string
```

Name returns the package name.

#### (*Package) Path 

``` go 
func (pkg *Package) Path() string
```

Path returns the package path.

#### (*Package) Scope 

``` go 
func (pkg *Package) Scope() *Scope
```

Scope returns the (complete or incomplete) package scope holding the objects declared at package level (TypeNames, Consts, Vars, and Funcs). For a nil pkg receiver, Scope returns the Universe scope.

#### (*Package) SetImports 

``` go 
func (pkg *Package) SetImports(list []*Package)
```

SetImports sets the list of explicitly imported packages to list. It is the caller's responsibility to make sure list elements are unique.

#### (*Package) SetName  <- go1.6

``` go 
func (pkg *Package) SetName(name string)
```

SetName sets the package name.

#### (*Package) String 

``` go 
func (pkg *Package) String() string
```

### type PkgName 

``` go 
type PkgName struct {
	// contains filtered or unexported fields
}
```

A PkgName represents an imported Go package. PkgNames don't have a type.

#### func NewPkgName 

``` go 
func NewPkgName(pos token.Pos, pkg *Package, name string, imported *Package) *PkgName
```

NewPkgName returns a new PkgName object representing an imported package. The remaining arguments set the attributes found with all Objects.

#### (*PkgName) Exported 

``` go 
func (obj *PkgName) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn't take into account whether the object is in a local (function) scope or not.

#### (*PkgName) Id 

``` go 
func (obj *PkgName) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

#### (*PkgName) Imported 

``` go 
func (obj *PkgName) Imported() *Package
```

Imported returns the package that was imported. It is distinct from Pkg(), which is the package containing the import statement.

#### (*PkgName) Name 

``` go 
func (obj *PkgName) Name() string
```

Name returns the object's (package-local, unqualified) name.

#### (*PkgName) Parent 

``` go 
func (obj *PkgName) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

#### (*PkgName) Pkg 

``` go 
func (obj *PkgName) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

#### (*PkgName) Pos 

``` go 
func (obj *PkgName) Pos() token.Pos
```

Pos returns the declaration position of the object's identifier.

#### (*PkgName) String 

``` go 
func (obj *PkgName) String() string
```

#### (*PkgName) Type 

``` go 
func (obj *PkgName) Type() Type
```

Type returns the object's type.

### type Pointer 

``` go 
type Pointer struct {
	// contains filtered or unexported fields
}
```

A Pointer represents a pointer type.

#### func NewPointer 

``` go 
func NewPointer(elem Type) *Pointer
```

NewPointer returns a new pointer type for the given element (base) type.

#### (*Pointer) Elem 

``` go 
func (p *Pointer) Elem() Type
```

Elem returns the element type for the given pointer p.

#### (*Pointer) String 

``` go 
func (t *Pointer) String() string
```

#### (*Pointer) Underlying 

``` go 
func (t *Pointer) Underlying() Type
```

### type Qualifier 

``` go 
type Qualifier func(*Package) string
```

A Qualifier controls how named package-level objects are printed in calls to TypeString, ObjectString, and SelectionString.

These three formatting routines call the Qualifier for each package-level object O, and if the Qualifier returns a non-empty string p, the object is printed in the form p.O. If it returns an empty string, only the object name O is printed.

Using a nil Qualifier is equivalent to using (*Package).Path: the object is qualified by the import path, e.g., "encoding/json.Marshal".

#### func RelativeTo 

``` go 
func RelativeTo(pkg *Package) Qualifier
```

RelativeTo returns a Qualifier that fully qualifies members of all packages other than pkg.

### type Scope 

``` go 
type Scope struct {
	// contains filtered or unexported fields
}
```

A Scope maintains a set of objects and links to its containing (parent) and contained (children) scopes. Objects may be inserted and looked up by name. The zero value for Scope is a ready-to-use empty scope.

##### Example
``` go 
```

``` go 
var Universe *Scope
```

The Universe scope contains all predeclared objects of Go. It is the outermost scope of any chain of nested scopes.

#### func NewScope 

``` go 
func NewScope(parent *Scope, pos, end token.Pos, comment string) *Scope
```

NewScope returns a new, empty scope contained in the given parent scope, if any. The comment is for debugging only.

#### (*Scope) Child 

``` go 
func (s *Scope) Child(i int) *Scope
```

Child returns the i'th child scope for 0 <= i < NumChildren().

#### (*Scope) Contains 

``` go 
func (s *Scope) Contains(pos token.Pos) bool
```

Contains reports whether pos is within the scope's extent. The result is guaranteed to be valid only if the type-checked AST has complete position information.

#### (*Scope) End 

``` go 
func (s *Scope) End() token.Pos
```

#### (*Scope) Innermost 

``` go 
func (s *Scope) Innermost(pos token.Pos) *Scope
```

Innermost returns the innermost (child) scope containing pos. If pos is not within any scope, the result is nil. The result is also nil for the Universe scope. The result is guaranteed to be valid only if the type-checked AST has complete position information.

#### (*Scope) Insert 

``` go 
func (s *Scope) Insert(obj Object) Object
```

Insert attempts to insert an object obj into scope s. If s already contains an alternative object alt with the same name, Insert leaves s unchanged and returns alt. Otherwise it inserts obj, sets the object's parent scope if not already set, and returns nil.

#### (*Scope) Len 

``` go 
func (s *Scope) Len() int
```

Len returns the number of scope elements.

#### (*Scope) Lookup 

``` go 
func (s *Scope) Lookup(name string) Object
```

Lookup returns the object in scope s with the given name if such an object exists; otherwise the result is nil.

#### (*Scope) LookupParent 

``` go 
func (s *Scope) LookupParent(name string, pos token.Pos) (*Scope, Object)
```

LookupParent follows the parent chain of scopes starting with s until it finds a scope where Lookup(name) returns a non-nil object, and then returns that scope and object. If a valid position pos is provided, only objects that were declared at or before pos are considered. If no such scope and object exists, the result is (nil, nil).

Note that obj.Parent() may be different from the returned scope if the object was inserted into the scope and already had a parent at that time (see Insert). This can only happen for dot-imported objects whose scope is the scope of the package that exported them.

#### (*Scope) Names 

``` go 
func (s *Scope) Names() []string
```

Names returns the scope's element names in sorted order.

#### (*Scope) NumChildren 

``` go 
func (s *Scope) NumChildren() int
```

NumChildren returns the number of scopes nested in s.

#### (*Scope) Parent 

``` go 
func (s *Scope) Parent() *Scope
```

Parent returns the scope's containing (parent) scope.

#### (*Scope) Pos 

``` go 
func (s *Scope) Pos() token.Pos
```

Pos and End describe the scope's source code extent [pos, end). The results are guaranteed to be valid only if the type-checked AST has complete position information. The extent is undefined for Universe and package scopes.

#### (*Scope) String 

``` go 
func (s *Scope) String() string
```

String returns a string representation of the scope, for debugging.

#### (*Scope) WriteTo 

``` go 
func (s *Scope) WriteTo(w io.Writer, n int, recurse bool)
```

WriteTo writes a string representation of the scope to w, with the scope elements sorted by name. The level of indentation is controlled by n >= 0, with n == 0 for no indentation. If recurse is set, it also writes nested (children) scopes.

### type Selection 

``` go 
type Selection struct {
	// contains filtered or unexported fields
}
```

A Selection describes a selector expression x.f. For the declarations:

``` go 
type T struct{ x int; E }
type E struct{}
func (e E) m() {}
var p *T
```

the following relations exist:

```
Selector    Kind          Recv    Obj    Type       Index     Indirect

p.x         FieldVal      T       x      int        {0}       true
p.m         MethodVal     *T      m      func()     {1, 0}    true
T.m         MethodExpr    T       m      func(T)    {1, 0}    false
```

#### (*Selection) Index 

``` go 
func (s *Selection) Index() []int
```

Index describes the path from x to f in x.f. The last index entry is the field or method index of the type declaring f; either:

1. the list of declared methods of a named type; or
2. the list of methods of an interface type; or
3. the list of fields of a struct type.

The earlier index entries are the indices of the embedded fields implicitly traversed to get from (the type of) x to f, starting at embedding depth 0.

#### (*Selection) Indirect 

``` go 
func (s *Selection) Indirect() bool
```

Indirect reports whether any pointer indirection was required to get from x to f in x.f.

#### (*Selection) Kind 

``` go 
func (s *Selection) Kind() SelectionKind
```

Kind returns the selection kind.

#### (*Selection) Obj 

``` go 
func (s *Selection) Obj() Object
```

Obj returns the object denoted by x.f; a *Var for a field selection, and a *Func in all other cases.

#### (*Selection) Recv 

``` go 
func (s *Selection) Recv() Type
```

Recv returns the type of x in x.f.

#### (*Selection) String 

``` go 
func (s *Selection) String() string
```

#### (*Selection) Type 

``` go 
func (s *Selection) Type() Type
```

Type returns the type of x.f, which may be different from the type of f. See Selection for more information.

### type SelectionKind 

``` go 
type SelectionKind int
```

SelectionKind describes the kind of a selector expression x.f (excluding qualified identifiers).

``` go 
const (
	FieldVal   SelectionKind = iota // x.f is a struct field selector
	MethodVal                       // x.f is a method selector
	MethodExpr                      // x.f is a method expression
)
```

### type Signature 

``` go 
type Signature struct {
	// contains filtered or unexported fields
}
```

A Signature represents a (non-builtin) function or method type. The receiver is ignored when comparing signatures for identity.

##### Example
``` go 
```

#### func NewSignatureType  <- go1.18

``` go 
func NewSignatureType(recv *Var, recvTypeParams, typeParams []*TypeParam, params, results *Tuple, variadic bool) *Signature
```

NewSignatureType creates a new function type for the given receiver, receiver type parameters, type parameters, parameters, and results. If variadic is set, params must hold at least one parameter and the last parameter's core type must be of unnamed slice or bytestring type. If recv is non-nil, typeParams must be empty. If recvTypeParams is non-empty, recv must be non-nil.

#### (*Signature) Params 

``` go 
func (s *Signature) Params() *Tuple
```

Params returns the parameters of signature s, or nil.

#### (*Signature) Recv 

``` go 
func (s *Signature) Recv() *Var
```

Recv returns the receiver of signature s (if a method), or nil if a function. It is ignored when comparing signatures for identity.

For an abstract method, Recv returns the enclosing interface either as a *Named or an *Interface. Due to embedding, an interface may contain methods whose receiver type is a different interface.

#### (*Signature) RecvTypeParams  <- go1.18

``` go 
func (s *Signature) RecvTypeParams() *TypeParamList
```

RecvTypeParams returns the receiver type parameters of signature s, or nil.

#### (*Signature) Results 

``` go 
func (s *Signature) Results() *Tuple
```

Results returns the results of signature s, or nil.

#### (*Signature) String 

``` go 
func (t *Signature) String() string
```

#### (*Signature) TypeParams  <- go1.18

``` go 
func (s *Signature) TypeParams() *TypeParamList
```

TypeParams returns the type parameters of signature s, or nil.

#### (*Signature) Underlying 

``` go 
func (t *Signature) Underlying() Type
```

#### (*Signature) Variadic 

``` go 
func (s *Signature) Variadic() bool
```

Variadic reports whether the signature s is variadic.

### type Sizes 

``` go 
type Sizes interface {
	// Alignof returns the alignment of a variable of type T.
	// Alignof must implement the alignment guarantees required by the spec.
	Alignof(T Type) int64

	// Offsetsof returns the offsets of the given struct fields, in bytes.
	// Offsetsof must implement the offset guarantees required by the spec.
	Offsetsof(fields []*Var) []int64

	// Sizeof returns the size of a variable of type T.
	// Sizeof must implement the size guarantees required by the spec.
	Sizeof(T Type) int64
}
```

Sizes defines the sizing functions for package unsafe.

#### func SizesFor  <- go1.9

``` go 
func SizesFor(compiler, arch string) Sizes
```

SizesFor returns the Sizes used by a compiler for an architecture. The result is nil if a compiler/architecture pair is not known.

Supported architectures for compiler "gc": "386", "amd64", "amd64p32", "arm", "arm64", "loong64", "mips", "mipsle", "mips64", "mips64le", "ppc64", "ppc64le", "riscv64", "s390x", "sparc64", "wasm".

### type Slice 

``` go 
type Slice struct {
	// contains filtered or unexported fields
}
```

A Slice represents a slice type.

#### func NewSlice 

``` go 
func NewSlice(elem Type) *Slice
```

NewSlice returns a new slice type for the given element type.

#### (*Slice) Elem 

``` go 
func (s *Slice) Elem() Type
```

Elem returns the element type of slice s.

#### (*Slice) String 

``` go 
func (t *Slice) String() string
```

#### (*Slice) Underlying 

``` go 
func (t *Slice) Underlying() Type
```

### type StdSizes 

``` go 
type StdSizes struct {
	WordSize int64 // word size in bytes - must be >= 4 (32bits)
	MaxAlign int64 // maximum alignment in bytes - must be >= 1
}
```

StdSizes is a convenience type for creating commonly used Sizes. It makes the following simplifying assumptions:

- The size of explicitly sized basic types (int16, etc.) is the specified size.
- The size of strings and interfaces is 2*WordSize.
- The size of slices is 3*WordSize.
- The size of an array of n elements corresponds to the size of a struct of n consecutive fields of the array's element type.
- The size of a struct is the offset of the last field plus that field's size. As with all element types, if the struct is used in an array its size must first be aligned to a multiple of the struct's alignment.
- All other types have size WordSize.
- Arrays and structs are aligned per spec definition; all other types are naturally aligned with a maximum alignment MaxAlign.

*StdSizes implements Sizes.

#### (*StdSizes) Alignof 

``` go 
func (s *StdSizes) Alignof(T Type) int64
```

#### (*StdSizes) Offsetsof 

``` go 
func (s *StdSizes) Offsetsof(fields []*Var) []int64
```

#### (*StdSizes) Sizeof 

``` go 
func (s *StdSizes) Sizeof(T Type) int64
```

### type Struct 

``` go 
type Struct struct {
	// contains filtered or unexported fields
}
```

A Struct represents a struct type.

#### func NewStruct 

``` go 
func NewStruct(fields []*Var, tags []string) *Struct
```

NewStruct returns a new struct with the given fields and corresponding field tags. If a field with index i has a tag, tags[i] must be that tag, but len(tags) may be only as long as required to hold the tag with the largest index i. Consequently, if no field has a tag, tags may be nil.

#### (*Struct) Field 

``` go 
func (s *Struct) Field(i int) *Var
```

Field returns the i'th field for 0 <= i < NumFields().

#### (*Struct) NumFields 

``` go 
func (s *Struct) NumFields() int
```

NumFields returns the number of fields in the struct (including blank and embedded fields).

#### (*Struct) String 

``` go 
func (t *Struct) String() string
```

#### (*Struct) Tag 

``` go 
func (s *Struct) Tag(i int) string
```

Tag returns the i'th field tag for 0 <= i < NumFields().

#### (*Struct) Underlying 

``` go 
func (t *Struct) Underlying() Type
```

### type Term  <- go1.18

``` go 
type Term term
```

A Term represents a term in a Union.

#### func NewTerm  <- go1.18

``` go 
func NewTerm(tilde bool, typ Type) *Term
```

NewTerm returns a new union term.

#### (*Term) String  <- go1.18

``` go 
func (t *Term) String() string
```

#### (*Term) Tilde  <- go1.18

``` go 
func (t *Term) Tilde() bool
```

#### (*Term) Type  <- go1.18

``` go 
func (t *Term) Type() Type
```

### type Tuple 

``` go 
type Tuple struct {
	// contains filtered or unexported fields
}
```

A Tuple represents an ordered list of variables; a nil *Tuple is a valid (empty) tuple. Tuples are used as components of signatures and to represent the type of multiple assignments; they are not first class types of Go.

#### func NewTuple 

``` go 
func NewTuple(x ...*Var) *Tuple
```

NewTuple returns a new tuple for the given variables.

#### (*Tuple) At 

``` go 
func (t *Tuple) At(i int) *Var
```

At returns the i'th variable of tuple t.

#### (*Tuple) Len 

``` go 
func (t *Tuple) Len() int
```

Len returns the number variables of tuple t.

#### (*Tuple) String 

``` go 
func (t *Tuple) String() string
```

#### (*Tuple) Underlying 

``` go 
func (t *Tuple) Underlying() Type
```

### type Type 

``` go 
type Type interface {
	// Underlying returns the underlying type of a type.
	Underlying() Type

	// String returns a string representation of a type.
	String() string
}
```

A Type represents a type of Go. All types implement the Type interface.

#### func Default  <- go1.8

``` go 
func Default(t Type) Type
```

Default returns the default "typed" type for an "untyped" type; it returns the incoming type for all other types. The default type for untyped nil is untyped nil.

#### func Instantiate  <- go1.18

``` go 
func Instantiate(ctxt *Context, orig Type, targs []Type, validate bool) (Type, error)
```

Instantiate instantiates the type orig with the given type arguments targs. orig must be a *Named or a *Signature type. If there is no error, the resulting Type is an instantiated type of the same kind (either a *Named or a *Signature). Methods attached to a *Named type are also instantiated, and associated with a new *Func that has the same position as the original method, but nil function scope.

If ctxt is non-nil, it may be used to de-duplicate the instance against previous instances with the same identity. As a special case, generic *Signature origin types are only considered identical if they are pointer equivalent, so that instantiating distinct (but possibly identical) signatures will yield different instances. The use of a shared context does not guarantee that identical instances are deduplicated in all cases.

If validate is set, Instantiate verifies that the number of type arguments and parameters match, and that the type arguments satisfy their corresponding type constraints. If verification fails, the resulting error may wrap an *ArgumentError indicating which type argument did not satisfy its corresponding type parameter constraint, and why.

If validate is not set, Instantiate does not verify the type argument count or whether the type arguments satisfy their constraints. Instantiate is guaranteed to not return an error, but may panic. Specifically, for *Signature types, Instantiate will panic immediately if the type argument count is incorrect; for *Named types, a panic may occur later inside the *Named API.

### type TypeAndValue 

``` go 
type TypeAndValue struct {
	Type  Type
	Value constant.Value
	// contains filtered or unexported fields
}
```

TypeAndValue reports the type and value (for constants) of the corresponding expression.

#### func Eval 

``` go 
func Eval(fset *token.FileSet, pkg *Package, pos token.Pos, expr string) (_ TypeAndValue, err error)
```

Eval returns the type and, if constant, the value for the expression expr, evaluated at position pos of package pkg, which must have been derived from type-checking an AST with complete position information relative to the provided file set.

The meaning of the parameters fset, pkg, and pos is the same as in CheckExpr. An error is returned if expr cannot be parsed successfully, or the resulting expr AST cannot be type-checked.

#### (TypeAndValue) Addressable 

``` go 
func (tv TypeAndValue) Addressable() bool
```

Addressable reports whether the corresponding expression is addressable (https://golang.org/ref/spec#Address_operators).

#### (TypeAndValue) Assignable 

``` go 
func (tv TypeAndValue) Assignable() bool
```

Assignable reports whether the corresponding expression is assignable to (provided a value of the right type).

#### (TypeAndValue) HasOk 

``` go 
func (tv TypeAndValue) HasOk() bool
```

HasOk reports whether the corresponding expression may be used on the rhs of a comma-ok assignment.

#### (TypeAndValue) IsBuiltin 

``` go 
func (tv TypeAndValue) IsBuiltin() bool
```

IsBuiltin reports whether the corresponding expression denotes a (possibly parenthesized) built-in function.

#### (TypeAndValue) IsNil 

``` go 
func (tv TypeAndValue) IsNil() bool
```

IsNil reports whether the corresponding expression denotes the predeclared value nil.

#### (TypeAndValue) IsType 

``` go 
func (tv TypeAndValue) IsType() bool
```

IsType reports whether the corresponding expression specifies a type.

#### (TypeAndValue) IsValue 

``` go 
func (tv TypeAndValue) IsValue() bool
```

IsValue reports whether the corresponding expression is a value. Builtins are not considered values. Constant values have a non- nil Value.

#### (TypeAndValue) IsVoid 

``` go 
func (tv TypeAndValue) IsVoid() bool
```

IsVoid reports whether the corresponding expression is a function call without results.

### type TypeList  <- go1.18

``` go 
type TypeList struct {
	// contains filtered or unexported fields
}
```

TypeList holds a list of types.

#### (*TypeList) At  <- go1.18

``` go 
func (l *TypeList) At(i int) Type
```

At returns the i'th type in the list.

#### (*TypeList) Len  <- go1.18

``` go 
func (l *TypeList) Len() int
```

Len returns the number of types in the list. It is safe to call on a nil receiver.

### type TypeName 

``` go 
type TypeName struct {
	// contains filtered or unexported fields
}
```

A TypeName represents a name for a (defined or alias) type.

#### func NewTypeName 

``` go 
func NewTypeName(pos token.Pos, pkg *Package, name string, typ Type) *TypeName
```

NewTypeName returns a new type name denoting the given typ. The remaining arguments set the attributes found with all Objects.

The typ argument may be a defined (Named) type or an alias type. It may also be nil such that the returned TypeName can be used as argument for NewNamed, which will set the TypeName's type as a side- effect.

#### (*TypeName) Exported 

``` go 
func (obj *TypeName) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn't take into account whether the object is in a local (function) scope or not.

#### (*TypeName) Id 

``` go 
func (obj *TypeName) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

#### (*TypeName) IsAlias  <- go1.9

``` go 
func (obj *TypeName) IsAlias() bool
```

IsAlias reports whether obj is an alias name for a type.

#### (*TypeName) Name 

``` go 
func (obj *TypeName) Name() string
```

Name returns the object's (package-local, unqualified) name.

#### (*TypeName) Parent 

``` go 
func (obj *TypeName) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

#### (*TypeName) Pkg 

``` go 
func (obj *TypeName) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

#### (*TypeName) Pos 

``` go 
func (obj *TypeName) Pos() token.Pos
```

Pos returns the declaration position of the object's identifier.

#### (*TypeName) String 

``` go 
func (obj *TypeName) String() string
```

#### (*TypeName) Type 

``` go 
func (obj *TypeName) Type() Type
```

Type returns the object's type.

### type TypeParam  <- go1.18

``` go 
type TypeParam struct {
	// contains filtered or unexported fields
}
```

A TypeParam represents a type parameter type.

#### func NewTypeParam  <- go1.18

``` go 
func NewTypeParam(obj *TypeName, constraint Type) *TypeParam
```

NewTypeParam returns a new TypeParam. Type parameters may be set on a Named or Signature type by calling SetTypeParams. Setting a type parameter on more than one type will result in a panic.

The constraint argument can be nil, and set later via SetConstraint. If the constraint is non-nil, it must be fully defined.

#### (*TypeParam) Constraint  <- go1.18

``` go 
func (t *TypeParam) Constraint() Type
```

Constraint returns the type constraint specified for t.

#### (*TypeParam) Index  <- go1.18

``` go 
func (t *TypeParam) Index() int
```

Index returns the index of the type param within its param list, or -1 if the type parameter has not yet been bound to a type.

#### (*TypeParam) Obj  <- go1.18

``` go 
func (t *TypeParam) Obj() *TypeName
```

Obj returns the type name for t.

#### (*TypeParam) SetConstraint  <- go1.18

``` go 
func (t *TypeParam) SetConstraint(bound Type)
```

SetConstraint sets the type constraint for t.

It must be called by users of NewTypeParam after the bound's underlying is fully defined, and before using the type parameter in any way other than to form other types. Once SetConstraint returns the receiver, t is safe for concurrent use.

#### (*TypeParam) String  <- go1.18

``` go 
func (t *TypeParam) String() string
```

#### (*TypeParam) Underlying  <- go1.18

``` go 
func (t *TypeParam) Underlying() Type
```

### type TypeParamList  <- go1.18

``` go 
type TypeParamList struct {
	// contains filtered or unexported fields
}
```

TypeParamList holds a list of type parameters.

#### (*TypeParamList) At  <- go1.18

``` go 
func (l *TypeParamList) At(i int) *TypeParam
```

At returns the i'th type parameter in the list.

#### (*TypeParamList) Len  <- go1.18

``` go 
func (l *TypeParamList) Len() int
```

Len returns the number of type parameters in the list. It is safe to call on a nil receiver.

### type Union  <- go1.18

``` go 
type Union struct {
	// contains filtered or unexported fields
}
```

A Union represents a union of terms embedded in an interface.

#### func NewUnion  <- go1.18

``` go 
func NewUnion(terms []*Term) *Union
```

NewUnion returns a new Union type with the given terms. It is an error to create an empty union; they are syntactically not possible.

#### (*Union) Len  <- go1.18

``` go 
func (u *Union) Len() int
```

#### (*Union) String  <- go1.18

``` go 
func (u *Union) String() string
```

#### (*Union) Term  <- go1.18

``` go 
func (u *Union) Term(i int) *Term
```

#### (*Union) Underlying  <- go1.18

``` go 
func (u *Union) Underlying() Type
```

### type Var 

``` go 
type Var struct {
	// contains filtered or unexported fields
}
```

A Variable represents a declared variable (including function parameters and results, and struct fields).

#### func NewField 

``` go 
func NewField(pos token.Pos, pkg *Package, name string, typ Type, embedded bool) *Var
```

NewField returns a new variable representing a struct field. For embedded fields, the name is the unqualified type name under which the field is accessible.

#### func NewParam 

``` go 
func NewParam(pos token.Pos, pkg *Package, name string, typ Type) *Var
```

NewParam returns a new variable representing a function parameter.

#### func NewVar 

``` go 
func NewVar(pos token.Pos, pkg *Package, name string, typ Type) *Var
```

NewVar returns a new variable. The arguments set the attributes found with all Objects.

#### (*Var) Anonymous 

``` go 
func (obj *Var) Anonymous() bool
```

Anonymous reports whether the variable is an embedded field. Same as Embedded; only present for backward-compatibility.

#### (*Var) Embedded  <- go1.11

``` go 
func (obj *Var) Embedded() bool
```

Embedded reports whether the variable is an embedded field.

#### (*Var) Exported 

``` go 
func (obj *Var) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn't take into account whether the object is in a local (function) scope or not.

#### (*Var) Id 

``` go 
func (obj *Var) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

#### (*Var) IsField 

``` go 
func (obj *Var) IsField() bool
```

IsField reports whether the variable is a struct field.

#### (*Var) Name 

``` go 
func (obj *Var) Name() string
```

Name returns the object's (package-local, unqualified) name.

#### (*Var) Origin  <- go1.19

``` go 
func (obj *Var) Origin() *Var
```

Origin returns the canonical Var for its receiver, i.e. the Var object recorded in Info.Defs.

For synthetic Vars created during instantiation (such as struct fields or function parameters that depend on type arguments), this will be the corresponding Var on the generic (uninstantiated) type. For all other Vars Origin returns the receiver.

#### (*Var) Parent 

``` go 
func (obj *Var) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

#### (*Var) Pkg 

``` go 
func (obj *Var) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

#### (*Var) Pos 

``` go 
func (obj *Var) Pos() token.Pos
```

Pos returns the declaration position of the object's identifier.

#### (*Var) String 

``` go 
func (obj *Var) String() string
```

#### (*Var) Type 

``` go 
func (obj *Var) Type() Type
```

Type returns the object's type.