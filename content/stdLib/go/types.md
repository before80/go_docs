+++
title = "types"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/types@go1.21.3](https://pkg.go.dev/go/types@go1.21.3)

Package types declares the data types and implements the algorithms for type-checking of Go packages. Use Config.Check to invoke the type checker for a package. Alternatively, create a new type checker with NewChecker and invoke it incrementally by calling Checker.Files.

​	types 包声明数据类型并实现 Go 包类型检查的算法。使用 Config.Check 调用包的类型检查器。或者，使用 NewChecker 创建一个新的类型检查器，并通过调用 Checker.Files 增量调用它。

Type-checking consists of several interdependent phases:

​	类型检查包含几个相互依赖的阶段：

Name resolution maps each identifier (ast.Ident) in the program to the language object (Object) it denotes. Use Info.{Defs,Uses,Implicits} for the results of name resolution.

​	名称解析将程序中的每个标识符 (ast.Ident) 映射到它表示的语言对象 (Object)。使用 Info.{Defs,Uses,Implicits} 获取名称解析的结果。

Constant folding computes the exact constant value (constant.Value) for every expression (ast.Expr) that is a compile-time constant. Use Info.Types[expr].Value for the results of constant folding.

​	常量折叠为每个编译时常量表达式 (ast.Expr) 计算精确的常量值 (constant.Value)。使用 Info.Types[expr].Value 获取常量折叠的结果。

Type inference computes the type (Type) of every expression (ast.Expr) and checks for compliance with the language specification. Use Info.Types[expr].Type for the results of type inference.

​	类型推断计算每个表达式 (ast.Expr) 的类型 (Type)，并检查是否符合语言规范。使用 Info.Types[expr].Type 获取类型推断的结果。

For a tutorial, see https://golang.org/s/types-tutorial.

​	有关教程，请参阅 https://golang.org/s/types-tutorial。

## 常量

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/go/types/universe.go;l=38)

```go
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

​	Typ 包含由其对应的 BasicKind 索引的前置声明的 *Basic 类型。

The *Basic type for Typ[Byte] will have the name “uint8”. Use Universe.Lookup(“byte”).Type() to obtain the specific alias basic type named “byte” (and analogous for “rune”).

​	Typ[Byte] 的基本类型将命名为“uint8”。使用 Universe.Lookup(“byte”).Type() 来获取名为“byte”的特定别名基本类型（对于“rune”也是类似的）。

## 函数

### func AssertableTo

```go
func AssertableTo(V *Interface, T Type) bool
```

AssertableTo reports whether a value of type V can be asserted to have type T.

​	AssertableTo 报告类型 V 的值是否可以断言为类型 T。

The behavior of AssertableTo is unspecified in three cases:

​	AssertableTo 的行为在三种情况下未指定：

- if T is Typ[Invalid]
  如果 T 是 Typ[Invalid]
- if V is a generalized interface; i.e., an interface that may only be used as a type constraint in Go code
  如果 V 是泛型接口；即，只能在 Go 代码中用作类型约束的接口
- if T is an uninstantiated generic type
  如果 T 是未实例化的泛型类型

### func AssignableTo

```go
func AssignableTo(V, T Type) bool
```

AssignableTo reports whether a value of type V is assignable to a variable of type T.

​	AssignableTo 报告类型 V 的值是否可赋值给类型 T 的变量。

The behavior of AssignableTo is unspecified if V or T is Typ[Invalid] or an uninstantiated generic type.

​	如果 V 或 T 是 Typ[Invalid] 或未实例化的泛型类型，则 AssignableTo 的行为未指定。

### func CheckExpr <- go1.13

```go
func CheckExpr(fset *token.FileSet, pkg *Package, pos token.Pos, expr ast.Expr, info *Info) (err error)
```

CheckExpr type checks the expression expr as if it had appeared at position pos of package pkg. Type information about the expression is recorded in info. The expression may be an identifier denoting an uninstantiated generic function or type.

​	CheckExpr 类型检查表达式 expr，就好像它出现在包 pkg 的位置 pos。有关表达式的类型信息记录在 info 中。该表达式可以是表示未实例化泛型函数或类型的标识符。

If pkg == nil, the Universe scope is used and the provided position pos is ignored. If pkg != nil, and pos is invalid, the package scope is used. Otherwise, pos must belong to the package.

​	如果 pkg == nil，则使用 Universe 作用域，并忽略提供的位置 pos。如果 pkg != nil，并且 pos 无效，则使用包作用域。否则，pos 必须属于该包。

An error is returned if pos is not within the package or if the node cannot be type-checked.

​	如果 pos 不在包内或无法对节点进行类型检查，则会返回错误。

Note: Eval and CheckExpr should not be used instead of running Check to compute types and values, but in addition to Check, as these functions ignore the context in which an expression is used (e.g., an assignment). Thus, top-level untyped constants will return an untyped type rather then the respective context-specific type.

​	注意：Eval 和 CheckExpr 不应代替运行 Check 来计算类型和值，而应与 Check 一起使用，因为这些函数忽略表达式使用的上下文（例如，赋值）。因此，顶级无类型常量将返回无类型类型，而不是各自的特定于上下文的类型。

### func Comparable

```go
func Comparable(T Type) bool
```

Comparable reports whether values of type T are comparable.

​	Comparable 报告类型 T 的值是否可比较。

### func ConvertibleTo 

```go
func ConvertibleTo(V, T Type) bool
```

ConvertibleTo reports whether a value of type V is convertible to a value of type T.

​	ConvertibleTo 报告类型 V 的值是否可转换为类型 T 的值。

The behavior of ConvertibleTo is unspecified if V or T is Typ[Invalid] or an uninstantiated generic type.

​	如果 V 或 T 是 Typ[Invalid] 或未实例化的泛型类型，则 ConvertibleTo 的行为未指定。

### func DefPredeclaredTestFuncs

```go
func DefPredeclaredTestFuncs()
```

DefPredeclaredTestFuncs defines the assert and trace built-ins. These built-ins are intended for debugging and testing of this package only.

​	DefPredeclaredTestFuncs 定义 assert 和 trace 内置函数。这些内置函数仅用于调试和测试此软件包。

### func ExprString

```go
func ExprString(x ast.Expr) string
```

ExprString returns the (possibly shortened) string representation for x. Shortened representations are suitable for user interfaces but may not necessarily follow Go syntax.

​	ExprString 返回 x 的（可能已缩短的）字符串表示形式。缩短的表示形式适用于用户界面，但不一定遵循 Go 语法。

### func Id

```go
func Id(pkg *Package, name string) string
```

Id returns name if it is exported, otherwise it returns the name qualified with the package path.

​	如果导出，Id 返回名称，否则返回使用包路径限定的名称。

### func Identical

```go
func Identical(x, y Type) bool
```

Identical reports whether x and y are identical types. Receivers of Signature types are ignored.

​	Identical 报告 x 和 y 是否是相同的类型。Signature 类型的接收者被忽略。

### func IdenticalIgnoreTags <- go1.8

```go
func IdenticalIgnoreTags(x, y Type) bool
```

IdenticalIgnoreTags reports whether x and y are identical types if tags are ignored. Receivers of Signature types are ignored.

​	IdenticalIgnoreTags 报告在忽略标签的情况下 x 和 y 是否为相同类型。Signature 类型的接收者被忽略。

### func Implements

```go
func Implements(V Type, T *Interface) bool
```

Implements reports whether type V implements interface T.

​	Implements 报告类型 V 是否实现接口 T。

The behavior of Implements is unspecified if V is Typ[Invalid] or an uninstantiated generic type.

​	如果 V 是 Typ[Invalid] 或未实例化的泛型类型，则 Implements 的行为未指定。

### func IsInterface

```go
func IsInterface(t Type) bool
```

IsInterface reports whether t is an interface type.

​	IsInterface 报告 t 是否为接口类型。

### func ObjectString

```go
func ObjectString(obj Object, qf Qualifier) string
```

ObjectString returns the string form of obj. The Qualifier controls the printing of package-level objects, and may be nil.

​	ObjectString 返回 obj 的字符串形式。Qualifier 控制包级对象的打印，可以为 nil。

### func Satisfies <- go1.20

```go
func Satisfies(V Type, T *Interface) bool
```

Satisfies reports whether type V satisfies the constraint T.

​	Satisfies 报告类型 V 是否满足约束 T。

The behavior of Satisfies is unspecified if V is Typ[Invalid] or an uninstantiated generic type.

​	如果 V 是 Typ[Invalid] 或未实例化的泛型类型，则 Satisfies 的行为未指定。

### func SelectionString

```go
func SelectionString(s *Selection, qf Qualifier) string
```

SelectionString returns the string form of s. The Qualifier controls the printing of package-level objects, and may be nil.

​	SelectionString 返回 s 的字符串形式。Qualifier 控制包级对象的打印，可以为 nil。

Examples:

​	示例：

```
"field (T) f int"
"method (T) f(X) Y"
"method expr (T) f(X) Y"
```

### func TypeString

```go
func TypeString(typ Type, qf Qualifier) string
```

TypeString returns the string representation of typ. The Qualifier controls the printing of package-level objects, and may be nil.

​	TypeString 返回 typ 的字符串表示形式。Qualifier 控制包级对象的打印，可以为 nil。

### func WriteExpr

```go
func WriteExpr(buf *bytes.Buffer, x ast.Expr)
```

WriteExpr writes the (possibly shortened) string representation for x to buf. Shortened representations are suitable for user interfaces but may not necessarily follow Go syntax.

​	WriteExpr 将 x 的（可能缩短的）字符串表示形式写入 buf。缩短的表示形式适用于用户界面，但不一定遵循 Go 语法。

### func WriteSignature

```go
func WriteSignature(buf *bytes.Buffer, sig *Signature, qf Qualifier)
```

WriteSignature writes the representation of the signature sig to buf, without a leading “func” keyword. The Qualifier controls the printing of package-level objects, and may be nil.

​	WriteSignature 将签名 sig 的表示形式写入 buf，不带前导“func”关键字。Qualifier 控制包级对象的打印，可以为 nil。

### func WriteType

```go
func WriteType(buf *bytes.Buffer, typ Type, qf Qualifier)
```

WriteType writes the string representation of typ to buf. The Qualifier controls the printing of package-level objects, and may be nil.

​	WriteType 将 typ 的字符串表示形式写入 buf。Qualifier 控制包级对象的打印，可以为 nil。

## 类型

### type ArgumentError <- go1.18

```go
type ArgumentError struct {
	Index int
	Err   error
}
```

An ArgumentError holds an error associated with an argument index.

​	ArgumentError 持有一个与参数索引关联的错误。

#### (*ArgumentError) Error <- go1.18

```go
func (e *ArgumentError) Error() string
```

#### (*ArgumentError) Unwrap <- go1.18

```go
func (e *ArgumentError) Unwrap() error
```

### type Array

```go
type Array struct {
	// contains filtered or unexported fields
}
```

An Array represents an array type.

​	Array 表示一个数组类型。

#### func NewArray

```go
func NewArray(elem Type, len int64) *Array
```

NewArray returns a new array type for the given element type and length. A negative length indicates an unknown length.

​	NewArray 为给定的元素类型和长度返回一个新的数组类型。负长度表示未知长度。

#### (*Array) Elem

```go
func (a *Array) Elem() Type
```

Elem returns element type of array a.

​	Elem 返回数组 a 的元素类型。

#### (*Array) Len

```go
func (a *Array) Len() int64
```

Len returns the length of array a. A negative result indicates an unknown length.

​	Len 返回数组 a 的长度。负结果表示未知长度。

#### (*Array) String 

```go
func (t *Array) String() string
```

#### (*Array) Underlying 

```go
func (t *Array) Underlying() Type
```

### type Basic

```go
type Basic struct {
	// contains filtered or unexported fields
}
```

A Basic represents a basic type.

​	基本表示基本类型。

#### (*Basic) Info

```go
func (b *Basic) Info() BasicInfo
```

Info returns information about properties of basic type b.

​	Info 返回有关基本类型 b 的属性的信息。

#### (*Basic) Kind 

```go
func (b *Basic) Kind() BasicKind
```

Kind returns the kind of basic type b.

​	类型返回基本类型 b 的类型。

#### (*Basic) Name 

```go
func (b *Basic) Name() string
```

Name returns the name of basic type b.

​	名称返回基本类型 b 的名称。

#### (*Basic) String

```go
func (t *Basic) String() string
```

#### (*Basic) Underlying 

```go
func (t *Basic) Underlying() Type
```

### type BasicInfo

```go
type BasicInfo int
```

BasicInfo is a set of flags describing properties of a basic type.

​	BasicInfo 是描述基本类型属性的一组标志。

```go
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

​	基本类型的属性。

### type BasicKind 

```go
type BasicKind int
```

BasicKind describes the kind of basic type.

​	BasicKind 描述基本类型的种类。

```go
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

```go
type Builtin struct {
	// contains filtered or unexported fields
}
```

A Builtin represents a built-in function. Builtins don’t have a valid type.

​	Builtin 表示内置函数。Builtin 没有有效类型。

#### (*Builtin) Exported

```go
func (obj *Builtin) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn’t take into account whether the object is in a local (function) scope or not.

​	Exported 报告对象是否导出（以大写字母开头）。它不考虑对象是否在本地（函数）作用域中。

#### (*Builtin) Id

```go
func (obj *Builtin) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

​	Id 是 Id(obj.Pkg(), obj.Name()) 的包装器。

#### (*Builtin) Name 

```go
func (obj *Builtin) Name() string
```

Name returns the object’s (package-local, unqualified) name.

​	Name 返回对象（包本地，不合格）的名称。

#### (*Builtin) Parent

```go
func (obj *Builtin) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

​	Parent 返回声明对象的范围。对于方法和结构字段，结果为 nil。

#### (*Builtin) Pkg 

```go
func (obj *Builtin) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

​	Pkg 返回对象所属的包。对于 Universe 范围内的标签和对象，结果为 nil。

#### (*Builtin) Pos 

```go
func (obj *Builtin) Pos() token.Pos
```

Pos returns the declaration position of the object’s identifier.

​	Pos 返回对象标识符的声明位置。

#### (*Builtin) String 

```go
func (obj *Builtin) String() string
```

#### (*Builtin) Type

```go
func (obj *Builtin) Type() Type
```

Type returns the object’s type.

​	Type 返回对象类型。

### type Chan

```go
type Chan struct {
	// contains filtered or unexported fields
}
```

A Chan represents a channel type.

​	A Chan 表示一个通道类型。

#### func NewChan

```go
func NewChan(dir ChanDir, elem Type) *Chan
```

NewChan returns a new channel type for the given direction and element type.

​	NewChan 为给定的方向和元素类型返回一个新的通道类型。

#### (*Chan) Dir 

```go
func (c *Chan) Dir() ChanDir
```

Dir returns the direction of channel c.

​	Dir 返回通道 c 的方向。

#### (*Chan) Elem 

```go
func (c *Chan) Elem() Type
```

Elem returns the element type of channel c.

​	Elem 返回通道 c 的元素类型。

#### (*Chan) String 

```go
func (t *Chan) String() string
```

#### (*Chan) Underlying 

```go
func (t *Chan) Underlying() Type
```

### type ChanDir 

```go
type ChanDir int
```

A ChanDir value indicates a channel direction.

​	A ChanDir 值指示通道方向。

```go
const (
	SendRecv ChanDir = iota
	SendOnly
	RecvOnly
)
```

The direction of a channel is indicated by one of these constants.

​	通道的方向由以下常量之一指示。

### type Checker

```go
type Checker struct {
	*Info
	// contains filtered or unexported fields
}
```

A Checker maintains the state of the type checker. It must be created with NewChecker.

​	检查器维护类型检查器的状态。它必须使用 NewChecker 创建。

#### func NewChecker

```go
func NewChecker(conf *Config, fset *token.FileSet, pkg *Package, info *Info) *Checker
```

NewChecker returns a new Checker instance for a given package. Package files may be added incrementally via checker.Files.

​	NewChecker 为给定程序包返回一个新的 Checker 实例。程序包文件可以通过 checker.Files 逐步添加。

#### (*Checker) Files 

```go
func (check *Checker) Files(files []*ast.File) error
```

Files checks the provided files as part of the checker’s package.

​	文件检查程序作为检查程序包的一部分检查提供的文件。

### type Config

```go
type Config struct {
	// Context is the context used for resolving global identifiers. If nil, the
	// type checker will initialize this field with a newly created context.
	Context *Context

	// GoVersion describes the accepted Go language version. The string must
	// start with a prefix of the form "go%d.%d" (e.g. "go1.20", "go1.21rc1", or
	// "go1.21.0") or it must be empty; an empty string disables Go language
	// version checks. If the format is invalid, invoking the type checker will
	// result in an error.
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

​	Config 指定类型检查的配置。Config 的零值是即用型默认配置。

#### (*Config) Check 

```go
func (conf *Config) Check(path string, fset *token.FileSet, files []*ast.File, info *Info) (*Package, error)
```

Check type-checks a package and returns the resulting package object and the first error if any. Additionally, if info != nil, Check populates each of the non-nil maps in the Info struct.

​	Check 类型检查一个包并返回结果包对象和第一个错误（如果有）。此外，如果 info 不为 nil，Check 会填充 Info 结构中每个非 nil 的映射。

The package is marked as complete if no errors occurred, otherwise it is incomplete. See Config.Error for controlling behavior in the presence of errors.

​	如果未发生错误，则将包标记为已完成，否则为未完成。有关在存在错误时控制行为的信息，请参阅 Config.Error。

The package is specified by a list of *ast.Files and corresponding file set, and the package path the package is identified with. The clean path must not be empty or dot (".").

​	包由 *ast.Files 和相应的文件集的列表以及包的标识路径指定。干净的路径不能是空或点（“.”）。

### type Const

```go
type Const struct {
	// contains filtered or unexported fields
}
```

A Const represents a declared constant.

​	Const 表示已声明的常量。

#### func NewConst

```go
func NewConst(pos token.Pos, pkg *Package, name string, typ Type, val constant.Value) *Const
```

NewConst returns a new constant with value val. The remaining arguments set the attributes found with all Objects.

​	NewConst 返回一个值为 val 的新常量。其余参数设置所有对象中找到的属性。

#### (*Const) Exported

```go
func (obj *Const) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn’t take into account whether the object is in a local (function) scope or not.

​	Exported 报告对象是否导出（以大写字母开头）。它不考虑对象是否在本地（函数）作用域中。

#### (*Const) Id

```go
func (obj *Const) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

​	Id 是 Id(obj.Pkg(), obj.Name()) 的包装器。

#### (*Const) Name

```go
func (obj *Const) Name() string
```

Name returns the object’s (package-local, unqualified) name.

​	Name 返回对象（包本地，不合格）的名称。

#### (*Const) Parent

```go
func (obj *Const) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

​	Parent 返回声明对象的作用域。对于方法和结构字段，结果为 nil。

#### (*Const) Pkg

```go
func (obj *Const) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

​	Pkg 返回对象所属的包。对于 Universe 范围内的标签和对象，结果为 nil。

#### (*Const) Pos

```go
func (obj *Const) Pos() token.Pos
```

Pos returns the declaration position of the object’s identifier.

​	Pos 返回对象标识符的声明位置。

#### (*Const) String

```go
func (obj *Const) String() string
```

#### (*Const) Type 

```go
func (obj *Const) Type() Type
```

Type returns the object’s type.

​	Type 返回对象类型。

#### (*Const) Val

```go
func (obj *Const) Val() constant.Value
```

Val returns the constant’s value.

​	Val 返回常量值。

### type Context <- go1.18

```go
type Context struct {
	// contains filtered or unexported fields
}
```

A Context is an opaque type checking context. It may be used to share identical type instances across type-checked packages or calls to Instantiate. Contexts are safe for concurrent use.

​	上下文是不透明的类型检查上下文。它可用于在类型检查的包或对 Instantiate 的调用之间共享相同的类型实例。上下文可安全地用于并发使用。

The use of a shared context does not guarantee that identical instances are deduplicated in all cases.

​	使用共享上下文并不能保证在所有情况下都对相同的实例进行去重。

#### func NewContext <- go1.18

```go
func NewContext() *Context
```

NewContext creates a new Context.

​	NewContext 创建一个新的 Context。

### type Error 

```go
type Error struct {
	Fset *token.FileSet // file set for interpretation of Pos
	Pos  token.Pos      // error position
	Msg  string         // error message
	Soft bool           // if set, error is "soft"
	// contains filtered or unexported fields
}
```

An Error describes a type-checking error; it implements the error interface. A “soft” error is an error that still permits a valid interpretation of a package (such as “unused variable”); “hard” errors may lead to unpredictable behavior if ignored.

​	错误描述了一个类型检查错误；它实现了错误接口。 “软”错误是一个仍然允许对包进行有效解释的错误（例如“未使用变量”）；如果忽略“硬”错误，可能会导致不可预测的行为。

#### (Error) Error 

```go
func (err Error) Error() string
```

Error returns an error string formatted as follows: filename:line:column: message

​	错误返回一个格式如下所示的错误字符串：文件名：行：列：消息

### type Func 

```go
type Func struct {
	// contains filtered or unexported fields
}
```

A Func represents a declared function, concrete method, or abstract (interface) method. Its Type() is always a *Signature. An abstract method may belong to many interfaces due to embedding.

​	Func 表示已声明的函数、具体方法或抽象（接口）方法。其 Type() 始终是 *Signature。抽象方法由于嵌入可能属于多个接口。

#### func MissingMethod

```go
func MissingMethod(V Type, T *Interface, static bool) (method *Func, wrongType bool)
```

MissingMethod returns (nil, false) if V implements T, otherwise it returns a missing method required by T and whether it is missing or just has the wrong type.

​	如果 V 实现 T，则 MissingMethod 返回 (nil, false)，否则它返回 T 所需的缺失方法，以及它是缺失还是仅类型错误。

For non-interface types V, or if static is set, V implements T if all methods of T are present in V. Otherwise (V is an interface and static is not set), MissingMethod only checks that methods of T which are also present in V have matching types (e.g., for a type assertion x.(T) where x is of interface type V).

​	对于非接口类型 V，或者如果设置了 static，则当 T 的所有方法都存在于 V 中时，V 实现 T。否则（V 是接口且未设置 static），MissingMethod 仅检查 T 中也存在于 V 中的方法是否具有匹配的类型（例如，对于类型断言 x.(T)，其中 x 为接口类型 V）。

#### func NewFunc

```go
func NewFunc(pos token.Pos, pkg *Package, name string, sig *Signature) *Func
```

NewFunc returns a new function with the given signature, representing the function’s type.

​	NewFunc 返回具有给定签名的函数的新函数，表示函数的类型。

#### (*Func) Exported

```go
func (obj *Func) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn’t take into account whether the object is in a local (function) scope or not.

​	Exported 报告对象是否已导出（以大写字母开头）。它不考虑对象是否在本地（函数）范围内。

#### (*Func) FullName 

```go
func (obj *Func) FullName() string
```

FullName returns the package- or receiver-type-qualified name of function or method obj.

​	FullName 返回函数或方法 obj 的包或接收器类型限定名称。

#### (*Func) Id 

```go
func (obj *Func) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

​	Id 是 Id(obj.Pkg(), obj.Name()) 的包装器。

#### (*Func) Name

```go
func (obj *Func) Name() string
```

Name returns the object’s (package-local, unqualified) name.

​	Name 返回对象（包本地，不合格）的名称。

#### (*Func) Origin <- go1.19

```go
func (obj *Func) Origin() *Func
```

Origin returns the canonical Func for its receiver, i.e. the Func object recorded in Info.Defs.

​	Origin 返回其接收者的规范 Func，即 Info.Defs 中记录的 Func 对象。

For synthetic functions created during instantiation (such as methods on an instantiated Named type or interface methods that depend on type arguments), this will be the corresponding Func on the generic (uninstantiated) type. For all other Funcs Origin returns the receiver.

​	对于在实例化期间创建的合成函数（例如，实例化 Named 类型上的方法或依赖于类型参数的接口方法），这将是泛型（未实例化）类型上的相应 Func。对于所有其他 Func，Origin 返回接收者。

#### (*Func) Parent 

```go
func (obj *Func) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

​	Parent 返回声明对象的作用域。对于方法和结构字段，结果为 nil。

#### (*Func) Pkg 

```go
func (obj *Func) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

​	Pkg 返回对象所属的包。对于 Universe 范围内的标签和对象，结果为 nil。

#### (*Func) Pos

```go
func (obj *Func) Pos() token.Pos
```

Pos returns the declaration position of the object’s identifier.

​	Pos 返回对象标识符的声明位置。

#### (*Func) Scope 

```go
func (obj *Func) Scope() *Scope
```

Scope returns the scope of the function’s body block. The result is nil for imported or instantiated functions and methods (but there is also no mechanism to get to an instantiated function).

​	Scope 返回函数主体块的范围。对于导入或实例化的函数和方法，结果为 nil（但也没有机制可以访问实例化的函数）。

#### (*Func) String 

```go
func (obj *Func) String() string
```

#### (*Func) Type 

```go
func (obj *Func) Type() Type
```

Type returns the object’s type.

​	Type 返回对象类型。

### type ImportMode <- go1.6

```go
type ImportMode int
```

ImportMode is reserved for future use.

​	ImportMode 保留供将来使用。

### type Importer

```go
type Importer interface {
	// Import returns the imported package for the given import path.
	// The semantics is like for ImporterFrom.ImportFrom except that
	// dir and mode are ignored (since they are not present).
	Import(path string) (*Package, error)
}
```

An Importer resolves import paths to Packages.

​	导入器将导入路径解析为包。

CAUTION: This interface does not support the import of locally vendored packages. See https://golang.org/s/go15vendor. If possible, external implementations should implement ImporterFrom.

​	注意：此接口不支持导入本地供应商包。请参阅 https://golang.org/s/go15vendor。如果可能，外部实现应实现 ImporterFrom。

### type ImporterFrom <- go1.6

```go
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

​	ImporterFrom 将导入路径解析为包；它支持 https://golang.org/s/go15vendor 中的供应商。使用 go/importer 获取 ImporterFrom 实现。

### type Info

```go
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

​	Info 为类型检查包保存结果类型信息。仅收集提供映射的信息。如果包有类型错误，则收集的信息可能不完整。

#### Example

ExampleInfo prints various facts recorded by the type checker in a types.Info struct: definitions of and references to each named object, and the type, value, and mode of every expression in the package.

​	ExampleInfo 在 types.Info 结构中打印类型检查器记录的各种事实：每个命名对象的定义和引用，以及包中每个表达式的类型、值和模式。

```go
// Parse a single source file.
const input = `
package fib

type S string

var a, b, c = len(b), S(c), "hello"

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) - fib(x-2)
}`
// We need a specific fileset in this test below for positions.
// Cannot use typecheck helper.
fset := token.NewFileSet()
f := mustParse(fset, input)

// Type-check the package.
// We create an empty map for each kind of input
// we're interested in, and Check populates them.
info := types.Info{
	Types: make(map[ast.Expr]types.TypeAndValue),
	Defs:  make(map[*ast.Ident]types.Object),
	Uses:  make(map[*ast.Ident]types.Object),
}
var conf types.Config
pkg, err := conf.Check("fib", fset, []*ast.File{f}, &info)
if err != nil {
	log.Fatal(err)
}

// Print package-level variables in initialization order.
fmt.Printf("InitOrder: %v\n\n", info.InitOrder)

// For each named object, print the line and
// column of its definition and each of its uses.
fmt.Println("Defs and Uses of each named object:")
usesByObj := make(map[types.Object][]string)
for id, obj := range info.Uses {
	posn := fset.Position(id.Pos())
	lineCol := fmt.Sprintf("%d:%d", posn.Line, posn.Column)
	usesByObj[obj] = append(usesByObj[obj], lineCol)
}
var items []string
for obj, uses := range usesByObj {
	sort.Strings(uses)
	item := fmt.Sprintf("%s:\n  defined at %s\n  used at %s",
		types.ObjectString(obj, types.RelativeTo(pkg)),
		fset.Position(obj.Pos()),
		strings.Join(uses, ", "))
	items = append(items, item)
}
sort.Strings(items) // sort by line:col, in effect
fmt.Println(strings.Join(items, "\n"))
fmt.Println()

fmt.Println("Types and Values of each expression:")
items = nil
for expr, tv := range info.Types {
	var buf strings.Builder
	posn := fset.Position(expr.Pos())
	tvstr := tv.Type.String()
	if tv.Value != nil {
		tvstr += " = " + tv.Value.String()
	}
	// line:col | expr | mode : type = value
	fmt.Fprintf(&buf, "%2d:%2d | %-19s | %-7s : %s",
		posn.Line, posn.Column, exprString(fset, expr),
		mode(tv), tvstr)
	items = append(items, buf.String())
}
sort.Strings(items)
fmt.Println(strings.Join(items, "\n"))


Output:

InitOrder: [c = "hello" b = S(c) a = len(b)]

Defs and Uses of each named object:
builtin len:
  defined at -
  used at 6:15
func fib(x int) int:
  defined at fib:8:6
  used at 12:20, 12:9
type S string:
  defined at fib:4:6
  used at 6:23
type int:
  defined at -
  used at 8:12, 8:17
type string:
  defined at -
  used at 4:8
var b S:
  defined at fib:6:8
  used at 6:19
var c string:
  defined at fib:6:11
  used at 6:25
var x int:
  defined at fib:8:10
  used at 10:10, 12:13, 12:24, 9:5

Types and Values of each expression:
 4: 8 | string              | type    : string
 6:15 | len                 | builtin : func(fib.S) int
 6:15 | len(b)              | value   : int
 6:19 | b                   | var     : fib.S
 6:23 | S                   | type    : fib.S
 6:23 | S(c)                | value   : fib.S
 6:25 | c                   | var     : string
 6:29 | "hello"             | value   : string = "hello"
 8:12 | int                 | type    : int
 8:17 | int                 | type    : int
 9: 5 | x                   | var     : int
 9: 5 | x < 2               | value   : untyped bool
 9: 9 | 2                   | value   : int = 2
10:10 | x                   | var     : int
12: 9 | fib                 | value   : func(x int) int
12: 9 | fib(x - 1)          | value   : int
12: 9 | fib(x-1) - fib(x-2) | value   : int
12:13 | x                   | var     : int
12:13 | x - 1               | value   : int
12:15 | 1                   | value   : int = 1
12:20 | fib                 | value   : func(x int) int
12:20 | fib(x - 2)          | value   : int
12:24 | x                   | var     : int
12:24 | x - 2               | value   : int
12:26 | 2                   | value   : int = 2
```

#### (*Info) ObjectOf 

```go
func (info *Info) ObjectOf(id *ast.Ident) Object
```

ObjectOf returns the object denoted by the specified id, or nil if not found.

​	ObjectOf 返回由指定 id 表示的对象，如果未找到，则返回 nil。

If id is an embedded struct field, ObjectOf returns the field (*Var) it defines, not the type (*TypeName) it uses.

​	如果 id 是嵌入式结构字段，ObjectOf 返回它定义的字段 (*Var)，而不是它使用的类型 (*TypeName)。

Precondition: the Uses and Defs maps are populated.

​	前提条件：Uses 和 Defs 映射已填充。

#### (*Info) TypeOf 

```go
func (info *Info) TypeOf(e ast.Expr) Type
```

TypeOf returns the type of expression e, or nil if not found. Precondition: the Types, Uses and Defs maps are populated.

​	TypeOf 返回表达式 e 的类型，如果未找到，则返回 nil。先决条件：已填充 Types、Uses 和 Defs 映射。

### type Initializer

```go
type Initializer struct {
	Lhs []*Var // var Lhs = Rhs
	Rhs ast.Expr
}
```

An Initializer describes a package-level variable, or a list of variables in case of a multi-valued initialization expression, and the corresponding initialization expression.

​	初始化器描述了一个包级变量，或者在多值初始化表达式的情况下描述了一个变量列表，以及相应的初始化表达式。

#### (*Initializer) String

```go
func (init *Initializer) String() string
```

### type Instance <- go1.18

```go
type Instance struct {
	TypeArgs *TypeList
	Type     Type
}
```

Instance reports the type arguments and instantiated type for type and function instantiations. For type instantiations, Type will be of dynamic type *Named. For function instantiations, Type will be of dynamic type *Signature.

​	实例报告类型参数和类型和函数实例化的实例化类型。对于类型实例化，Type 将为动态类型 *Named。对于函数实例化，Type 将为动态类型 *Signature。

### type Interface 

```go
type Interface struct {
	// contains filtered or unexported fields
}
```

An Interface represents an interface type.

​	接口表示一个接口类型。

#### func NewInterface <-DEPRECATED

```go
func NewInterface(methods []*Func, embeddeds []*Named) *Interface
```

NewInterface returns a new interface for the given methods and embedded types. NewInterface takes ownership of the provided methods and may modify their types by setting missing receivers.

​	NewInterface 为给定方法和嵌入式类型返回一个新接口。NewInterface 拥有所提供方法的所有权，并且可以通过设置缺失的接收器来修改其类型。

Deprecated: Use NewInterfaceType instead which allows arbitrary embedded types.

​	已弃用：改用 NewInterfaceType，它允许任意嵌入式类型。

#### func NewInterfaceType <- go1.11

```go
func NewInterfaceType(methods []*Func, embeddeds []Type) *Interface
```

NewInterfaceType returns a new interface for the given methods and embedded types. NewInterfaceType takes ownership of the provided methods and may modify their types by setting missing receivers.

​	NewInterfaceType 为给定方法和嵌入式类型返回一个新接口。NewInterfaceType 拥有所提供的这些方法的所有权，并且可以通过设置缺失的接收器来修改它们类型。

To avoid race conditions, the interface’s type set should be computed before concurrent use of the interface, by explicitly calling Complete.

​	为了避免竞争条件，接口的类型集应该在并发使用接口之前通过显式调用 Complete 来计算。

#### (*Interface) Complete

```go
func (t *Interface) Complete() *Interface
```

Complete computes the interface’s type set. It must be called by users of NewInterfaceType and NewInterface after the interface’s embedded types are fully defined and before using the interface type in any way other than to form other types. The interface must not contain duplicate methods or a panic occurs. Complete returns the receiver.

​	Complete 计算接口的类型集。在接口的嵌入式类型完全定义之后，并且在以形成其他类型以外的任何方式使用接口类型之前，NewInterfaceType 和 NewInterface 的用户必须调用它。接口中不得包含重复的方法，否则会发生 panic。Complete 返回接收器。

Interface types that have been completed are safe for concurrent use.

​	已完成的接口类型可安全地并发使用。

#### (*Interface) Embedded <-DEPRECATED 

```go
func (t *Interface) Embedded(i int) *Named
```

Embedded returns the i’th embedded defined (*Named) type of interface t for 0 <= i < t.NumEmbeddeds(). The result is nil if the i’th embedded type is not a defined type.

​	Embedded 返回接口 t 的第 i 个嵌入式已定义的 (*Named) 类型，其中 0 <= i < t.NumEmbeddeds()。如果第 i 个嵌入式类型不是已定义类型，则结果为 nil。

Deprecated: Use EmbeddedType which is not restricted to defined (*Named) types.

​	已弃用：使用 EmbeddedType，它不限于已定义的 (*Named) 类型。

#### (*Interface) EmbeddedType <- go1.11

```go
func (t *Interface) EmbeddedType(i int) Type
```

EmbeddedType returns the i’th embedded type of interface t for 0 <= i < t.NumEmbeddeds().

​	EmbeddedType 返回接口 t 的第 i 个嵌入式类型，其中 0 <= i < t.NumEmbeddeds()。

#### (*Interface) Empty 

```go
func (t *Interface) Empty() bool
```

Empty reports whether t is the empty interface.

​	Empty 报告 t 是否为空接口。

#### (*Interface) ExplicitMethod

```go
func (t *Interface) ExplicitMethod(i int) *Func
```

ExplicitMethod returns the i’th explicitly declared method of interface t for 0 <= i < t.NumExplicitMethods(). The methods are ordered by their unique Id. 

​	ExplicitMethod 返回接口 t 的第 i 个显式声明的方法，其中 0 <= i < t.NumExplicitMethods()。这些方法按其唯一 Id 排序。

#### (*Interface) IsComparable <- go1.18

```go
func (t *Interface) IsComparable() bool
```

IsComparable reports whether each type in interface t’s type set is comparable. 

​	IsComparable 报告接口 t 的类型集中每个类型是否可比较。

#### (*Interface) IsImplicit <- go1.18

```go
func (t *Interface) IsImplicit() bool
```

IsImplicit reports whether the interface t is a wrapper for a type set literal.

​	IsImplicit 报告接口 t 是否是类型集文字的包装器。

#### (*Interface) IsMethodSet <- go1.18

```go
func (t *Interface) IsMethodSet() bool
```

IsMethodSet reports whether the interface t is fully described by its method set.

​	IsMethodSet 报告接口 t 是否由其方法集完全描述。

#### (*Interface) MarkImplicit <- go1.18

```go
func (t *Interface) MarkImplicit()
```

MarkImplicit marks the interface t as implicit, meaning this interface corresponds to a constraint literal such as ~T or A|B without explicit interface embedding. MarkImplicit should be called before any concurrent use of implicit interfaces.

​	MarkImplicit 将接口 t 标记为隐式，这意味着此接口对应于约束文字，例如 ~T 或 A|B，而没有显式接口嵌入。应在隐式接口的任何并发使用之前调用 MarkImplicit。

#### (*Interface) Method 

```go
func (t *Interface) Method(i int) *Func
```

Method returns the i’th method of interface t for 0 <= i < t.NumMethods(). The methods are ordered by their unique Id.

​	方法返回接口 t 的第 i 个方法，其中 0 <= i < t.NumMethods()。这些方法按其唯一 Id 排序。

#### (*Interface) NumEmbeddeds

```go
func (t *Interface) NumEmbeddeds() int
```

NumEmbeddeds returns the number of embedded types in interface t.

​	NumEmbeddeds 返回接口 t 中的嵌入式类型的数量。

#### (*Interface) NumExplicitMethods

```go
func (t *Interface) NumExplicitMethods() int
```

NumExplicitMethods returns the number of explicitly declared methods of interface t.

​	NumExplicitMethods 返回接口 t 中显式声明的方法数。

#### (*Interface) NumMethods 

```go
func (t *Interface) NumMethods() int
```

NumMethods returns the total number of methods of interface t.

​	NumMethods 返回接口 t 的方法总数。

#### (*Interface) String

```go
func (t *Interface) String() string
```

#### (*Interface) Underlying

```go
func (t *Interface) Underlying() Type
```

### type Label

```go
type Label struct {
	// contains filtered or unexported fields
}
```

A Label represents a declared label. Labels don’t have a type.

​	标签表示已声明的标签。标签没有类型。

#### func NewLabel

```go
func NewLabel(pos token.Pos, pkg *Package, name string) *Label
```

NewLabel returns a new label.

​	NewLabel 返回一个新标签。

#### (*Label) Exported

```go
func (obj *Label) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn’t take into account whether the object is in a local (function) scope or not.

​	Exported 报告对象是否已导出（以大写字母开头）。它不考虑对象是否在本地（函数）范围内。

#### (*Label) Id 

```go
func (obj *Label) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

​	Id 是 Id(obj.Pkg(), obj.Name()) 的包装器。

#### (*Label) Name

```go
func (obj *Label) Name() string
```

Name returns the object’s (package-local, unqualified) name.

​	Name 返回对象（包本地、不合格）的名称。

#### (*Label) Parent 

```go
func (obj *Label) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

​	Parent 返回声明对象的作用域。对于方法和结构字段，结果为 nil。

#### (*Label) Pkg

```go
func (obj *Label) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

​	Pkg 返回对象所属的包。对于 Universe 范围内的标签和对象，结果为 nil。

#### (*Label) Pos

```go
func (obj *Label) Pos() token.Pos
```

Pos returns the declaration position of the object’s identifier.

​	Pos 返回对象标识符的声明位置。

#### (*Label) String 

```go
func (obj *Label) String() string
```

#### (*Label) Type

```go
func (obj *Label) Type() Type
```

Type returns the object’s type.

​	Type 返回对象类型。

### type Map

```go
type Map struct {
	// contains filtered or unexported fields
}
```

A Map represents a map type.

​	Map 表示映射类型。

#### func NewMap

```go
func NewMap(key, elem Type) *Map
```

NewMap returns a new map for the given key and element types.

​	NewMap 返回一个具有给定键和元素类型的新映射。

#### (*Map) Elem

```go
func (m *Map) Elem() Type
```

Elem returns the element type of map m.

​	Elem 返回映射 m 的元素类型。

#### (*Map) Key

```go
func (m *Map) Key() Type
```

Key returns the key type of map m.

​	Key 返回映射 m 的键类型。

#### (*Map) String 

```go
func (t *Map) String() string
```

#### (*Map) Underlying

```go
func (t *Map) Underlying() Type
```

### type MethodSet 

```go
type MethodSet struct {
	// contains filtered or unexported fields
}
```

A MethodSet is an ordered set of concrete or abstract (interface) methods; a method is a MethodVal selection, and they are ordered by ascending m.Obj().Id(). The zero value for a MethodSet is a ready-to-use empty method set. 重试 错误原因

##### Example

ExampleMethodSet prints the method sets of various types.

​	ExampleMethodSet 打印各种类型的函数集。

```go
package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	// Parse a single source file.
	const input = `
package temperature
import "fmt"
type Celsius float64
func (c Celsius) String() string  { return fmt.Sprintf("%g°C", c) }
func (c *Celsius) SetF(f float64) { *c = Celsius(f - 32 / 9 * 5) }

type S struct { I; m int }
type I interface { m() byte }
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "celsius.go", input, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Type-check a package consisting of this file.
	// Type information for the imported packages
	// comes from $GOROOT/pkg/$GOOS_$GOOARCH/fmt.a.
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("temperature", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print the method sets of Celsius and *Celsius.
	celsius := pkg.Scope().Lookup("Celsius").Type()
	for _, t := range []types.Type{celsius, types.NewPointer(celsius)} {
		fmt.Printf("Method set of %s:\n", t)
		mset := types.NewMethodSet(t)
		for i := 0; i < mset.Len(); i++ {
			fmt.Println(mset.At(i))
		}
		fmt.Println()
	}

	// Print the method set of S.
	styp := pkg.Scope().Lookup("S").Type()
	fmt.Printf("Method set of %s:\n", styp)
	fmt.Println(types.NewMethodSet(styp))

}
Output:

Method set of temperature.Celsius:
method (temperature.Celsius) String() string

Method set of *temperature.Celsius:
method (*temperature.Celsius) SetF(f float64)
method (*temperature.Celsius) String() string

Method set of temperature.S:
MethodSet {}
```

#### func NewMethodSet

```go
func NewMethodSet(T Type) *MethodSet
```

NewMethodSet returns the method set for the given type T. It always returns a non-nil method set, even if it is empty.

​	NewMethodSet 返回给定类型 T 的方法集。即使方法集为空，它也始终返回一个非空方法集。

#### (*MethodSet) At

```go
func (s *MethodSet) At(i int) *Selection
```

At returns the i’th method in s for 0 <= i < s.Len().

​	At 返回 s 中的第 i 个方法，其中 0 <= i < s.Len()。

#### (*MethodSet) Len

```go
func (s *MethodSet) Len() int
```

Len returns the number of methods in s.

​	Len 返回 s 中的方法数。

#### (*MethodSet) Lookup

```go
func (s *MethodSet) Lookup(pkg *Package, name string) *Selection
```

Lookup returns the method with matching package and name, or nil if not found.

​	Lookup 返回具有匹配包和名称的方法，如果未找到，则返回 nil。

#### (*MethodSet) String 

```go
func (s *MethodSet) String() string
```

### type Named

```go
type Named struct {
	// contains filtered or unexported fields
}
```

A Named represents a named (defined) type.

​	Named 表示一个命名（已定义）的类型。

#### func NewNamed

```go
func NewNamed(obj *TypeName, underlying Type, methods []*Func) *Named
```

NewNamed returns a new named type for the given type name, underlying type, and associated methods. If the given type name obj doesn’t have a type yet, its type is set to the returned named type. The underlying type must not be a *Named.

​	NewNamed 为给定的类型名称、基础类型和关联方法返回一个新的命名类型。如果给定的类型名称 obj 还没有类型，则其类型将设置为返回的命名类型。基础类型不能是 *Named。

#### (*Named) AddMethod 

```go
func (t *Named) AddMethod(m *Func)
```

AddMethod adds method m unless it is already in the method list. t must not have type arguments.

​	AddMethod 添加方法 m，除非它已在方法列表中。t 不得具有类型参数。

#### (*Named) Method 

```go
func (t *Named) Method(i int) *Func
```

Method returns the i’th method of named type t for 0 <= i < t.NumMethods().

​	方法返回命名类型 t 的第 i 个方法，其中 0 <= i < t.NumMethods()。

For an ordinary or instantiated type t, the receiver base type of this method is the named type t. For an uninstantiated generic type t, each method receiver is instantiated with its receiver type parameters.

​	对于普通或实例化类型 t，此方法的接收器基类型是命名类型 t。对于未实例化的泛型类型 t，每个方法接收器都使用其接收器类型参数进行实例化。

#### (*Named) NumMethods 

```go
func (t *Named) NumMethods() int
```

NumMethods returns the number of explicit methods defined for t.

​	NumMethods 返回为 t 定义的显式方法的数量。

#### (*Named) Obj

```go
func (t *Named) Obj() *TypeName
```

Obj returns the type name for the declaration defining the named type t. For instantiated types, this is same as the type name of the origin type.

​	Obj 返回定义命名类型 t 的声明的类型名称。对于实例化类型，这与原始类型的类型名称相同。

#### (*Named) Origin <- go1.18 

```go
func (t *Named) Origin() *Named
```

Origin returns the generic type from which the named type t is instantiated. If t is not an instantiated type, the result is t.

​	Origin 返回命名类型 t 实例化的泛型类型。如果 t 不是实例化类型，则结果为 t。

#### (*Named) SetTypeParams <- go1.18 

```go
func (t *Named) SetTypeParams(tparams []*TypeParam)
```

SetTypeParams sets the type parameters of the named type t. t must not have type arguments.

​	SetTypeParams 设置命名类型 t 的类型参数。t 不得具有类型参数。

#### (*Named) SetUnderlying 

```go
func (t *Named) SetUnderlying(underlying Type)
```

SetUnderlying sets the underlying type and marks t as complete. t must not have type arguments.

​	SetUnderlying 设置基础类型并将 t 标记为已完成。t 不得具有类型参数。

#### (*Named) String 

```go
func (t *Named) String() string
```

#### (*Named) TypeArgs <- go1.18

```go
func (t *Named) TypeArgs() *TypeList
```

TypeArgs returns the type arguments used to instantiate the named type t.

​	TypeArgs 返回用于实例化命名类型 t 的类型参数。

#### (*Named) TypeParams <- go1.18 

```go
func (t *Named) TypeParams() *TypeParamList
```

TypeParams returns the type parameters of the named type t, or nil. The result is non-nil for an (originally) generic type even if it is instantiated.

​	TypeParams 返回命名类型 t 的类型参数，或 nil。即使实例化，结果对于（最初的）泛型类型也是非 nil。

#### (*Named) Underlying 

```go
func (t *Named) Underlying() Type
```

### type Nil 

```go
type Nil struct {
	// contains filtered or unexported fields
}
```

Nil represents the predeclared value nil.

​	Nil 表示预声明值 nil。

#### (*Nil) Exported 

```go
func (obj *Nil) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn’t take into account whether the object is in a local (function) scope or not.

​	导出报告对象是否导出（以大写字母开头）。它不考虑对象是否在本地（函数）范围内。

#### (*Nil) Id

```go
func (obj *Nil) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

​	Id 是 Id(obj.Pkg(), obj.Name()) 的包装器。

#### (*Nil) Name

```go
func (obj *Nil) Name() string
```

Name returns the object’s (package-local, unqualified) name.

​	Name 返回对象（包本地、不合格）的名称。

#### (*Nil) Parent 

```go
func (obj *Nil) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

​	Parent 返回声明对象的作用域。对于方法和结构字段，结果为 nil。

#### (*Nil) Pkg

```go
func (obj *Nil) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

​	Pkg 返回对象所属的包。对于 Universe 范围内的标签和对象，结果为 nil。

#### (*Nil) Pos

```go
func (obj *Nil) Pos() token.Pos
```

Pos returns the declaration position of the object’s identifier.

​	Pos 返回对象标识符的声明位置。

#### (*Nil) String

```go
func (obj *Nil) String() string
```

#### (*Nil) Type 

```go
func (obj *Nil) Type() Type
```

Type returns the object’s type.

​	Type 返回对象类型。

### type Object

```go
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

​	对象描述一个命名的语言实体，例如包、常量、类型、变量、函数（包括方法）或标签。所有对象都实现 Object 接口。

#### func LookupFieldOrMethod

```go
func LookupFieldOrMethod(T Type, addressable bool, pkg *Package, name string) (obj Object, index []int, indirect bool)
```

LookupFieldOrMethod looks up a field or method with given package and name in T and returns the corresponding *Var or *Func, an index sequence, and a bool indicating if there were any pointer indirections on the path to the field or method. If addressable is set, T is the type of an addressable variable (only matters for method lookups). T must not be nil.

​	LookupFieldOrMethod 查找 T 中具有给定包和名称的字段或方法，并返回相应的 *Var 或 *Func、索引序列和一个布尔值，指示在字段或方法的路径上是否有任何指针间接。如果设置了可寻址，则 T 是可寻址变量的类型（仅对方法查找很重要）。T 不能为 nil。

The last index entry is the field or method index in the (possibly embedded) type where the entry was found, either:

​	最后一个索引条目是找到条目的（可能嵌入的）类型中的字段或方法索引，可能是：

1. the list of declared methods of a named type; or
   已命名类型的声明方法列表；或
2. the list of all methods (method set) of an interface type; or
   接口类型的全部方法列表（方法集）；或
3. the list of fields of a struct type.
   结构类型字段列表。

The earlier index entries are the indices of the embedded struct fields traversed to get to the found entry, starting at depth 0.

​	较早的索引条目是从深度 0 开始，遍历以获取找到的条目而穿越的嵌入式结构字段的索引。

If no entry is found, a nil object is returned. In this case, the returned index and indirect values have the following meaning:

​	如果未找到条目，则返回 nil 对象。在这种情况下，返回的索引和间接值具有以下含义：

- If index != nil, the index sequence points to an ambiguous entry (the same name appeared more than once at the same embedding level).
  如果索引！= nil，则索引序列指向一个不明确的条目（同一名称在同一嵌入级别出现多次）。
- If indirect is set, a method with a pointer receiver type was found but there was no pointer on the path from the actual receiver type to the method’s formal receiver base type, nor was the receiver addressable.
  如果设置了 indirect，则找到一个具有指针接收器类型的方法，但从实际接收器类型到方法的正式接收器基本类型没有指针，接收器也不是可寻址的。

### type Package

```go
type Package struct {
	// contains filtered or unexported fields
}
```

A Package describes a Go package.

​	Package 描述了一个 Go 包。

```go
var Unsafe *Package
```

The Unsafe package is the package returned by an importer for the import path “unsafe”.

​	Unsafe 包是导入程序为导入路径“unsafe”返回的包。

#### func NewPackage

```go
func NewPackage(path, name string) *Package
```

NewPackage returns a new Package for the given package path and name. The package is not complete and contains no explicit imports.

​	NewPackage 为给定的包路径和名称返回一个新包。该包不完整，不包含任何显式导入。

#### (*Package) Complete

```go
func (pkg *Package) Complete() bool
```

A package is complete if its scope contains (at least) all exported objects; otherwise it is incomplete.

​	如果包的范围包含（至少）所有导出的对象，则该包是完整的；否则，该包是不完整的。

#### (*Package) GoVersion <-go1.21.0

```go
func (pkg *Package) GoVersion() string
```

GoVersion returns the minimum Go version required by this package. If the minimum version is unknown, GoVersion returns the empty string. Individual source files may specify a different minimum Go version, as reported in the [go/ast.File.GoVersion](https://pkg.go.dev/go/ast#File.GoVersion) field.

​	GoVersion 返回此包所需的最低 Go 版本。如果最低版本未知，GoVersion 返回空字符串。各个源文件可能会指定不同的最低 Go 版本，如 go/ast.File.GoVersion 字段中所述。

#### (*Package) Imports

```go
func (pkg *Package) Imports() []*Package
```

Imports returns the list of packages directly imported by pkg; the list is in source order.

​	Imports 返回 pkg 直接导入的包列表；该列表按源顺序排列。

If pkg was loaded from export data, Imports includes packages that provide package-level objects referenced by pkg. This may be more or less than the set of packages directly imported by pkg’s source code.

​	如果 pkg 是从导出数据加载的，则 Imports 包括提供 pkg 引用的包级对象的包。这可能多于或少于 pkg 的源代码直接导入的包集。

If pkg uses cgo and the FakeImportC configuration option was enabled, the imports list may contain a fake “C” package.

​	如果 pkg 使用 cgo 并且启用了 FakeImportC 配置选项，则导入列表可能包含一个假的“C”包。

#### (*Package) MarkComplete

```go
func (pkg *Package) MarkComplete()
```

MarkComplete marks a package as complete.

​	MarkComplete 将一个包标记为已完成。

#### (*Package) Name

```go
func (pkg *Package) Name() string
```

Name returns the package name.

​	Name 返回包名称。

#### (*Package) Path

```go
func (pkg *Package) Path() string
```

Path returns the package path.

​	Path 返回包路径。

#### (*Package) Scope

```go
func (pkg *Package) Scope() *Scope
```

Scope returns the (complete or incomplete) package scope holding the objects declared at package level (TypeNames, Consts, Vars, and Funcs). For a nil pkg receiver, Scope returns the Universe scope.

​	Scope 返回包含在包级别声明的对象（TypeName、Consts、Vars 和 Funcs）的（完整或不完整）包范围。对于 nil pkg 接收器，Scope 返回 Universe 范围。

#### (*Package) SetImports

```go
func (pkg *Package) SetImports(list []*Package)
```

SetImports sets the list of explicitly imported packages to list. It is the caller’s responsibility to make sure list elements are unique.

​	SetImports 将显式导入的包列表设置为 list。由调用者负责确保列表元素是唯一的。

#### (*Package) SetName <- go1.6

```go
func (pkg *Package) SetName(name string)
```

SetName sets the package name.

​	SetName 设置包名称。

#### (*Package) String

```go
func (pkg *Package) String() string
```

### type PkgName

```go
type PkgName struct {
	// contains filtered or unexported fields
}
```

A PkgName represents an imported Go package. PkgNames don’t have a type.

​	PkgName 表示导入的 Go 包。PkgName 没有类型。

#### func NewPkgName

```go
func NewPkgName(pos token.Pos, pkg *Package, name string, imported *Package) *PkgName
```

NewPkgName returns a new PkgName object representing an imported package. The remaining arguments set the attributes found with all Objects.

​	NewPkgName 返回一个表示导入包的新 PkgName 对象。其余参数设置在所有对象中找到的属性。

#### (*PkgName) Exported

```go
func (obj *PkgName) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn’t take into account whether the object is in a local (function) scope or not.

​	Exported 报告对象是否导出（以大写字母开头）。它不考虑对象是否在本地（函数）范围内。

#### (*PkgName) Id

```go
func (obj *PkgName) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

​	Id 是 Id(obj.Pkg(), obj.Name()) 的包装器。

#### (*PkgName) Imported 

```go
func (obj *PkgName) Imported() *Package
```

Imported returns the package that was imported. It is distinct from Pkg(), which is the package containing the import statement.

​	Imported 返回导入的包。它不同于 Pkg()，后者是包含 import 语句的包。

#### (*PkgName) Name

```go
func (obj *PkgName) Name() string
```

Name returns the object’s (package-local, unqualified) name.

​	Name 返回对象（包本地、不合格）的名称。

#### (*PkgName) Parent 

```go
func (obj *PkgName) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

​	Parent 返回声明对象的作用域。对于方法和结构字段，结果为 nil。

#### (*PkgName) Pkg

```go
func (obj *PkgName) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

​	Pkg 返回对象所属的包。对于 Universe 范围内的标签和对象，结果为 nil。

#### (*PkgName) Pos

```go
func (obj *PkgName) Pos() token.Pos
```

Pos returns the declaration position of the object’s identifier.

​	Pos 返回对象标识符的声明位置。

#### (*PkgName) String

```go
func (obj *PkgName) String() string
```

#### (*PkgName) Type 

```go
func (obj *PkgName) Type() Type
```

Type returns the object’s type.

​	Type 返回对象类型。

### type Pointer 

```go
type Pointer struct {
	// contains filtered or unexported fields
}
```

A Pointer represents a pointer type.

​	指针表示指针类型。

#### func NewPointer

```go
func NewPointer(elem Type) *Pointer
```

NewPointer returns a new pointer type for the given element (base) type.

​	NewPointer 返回给定元素（基本）类型的新的指针类型。

#### (*Pointer) Elem

```go
func (p *Pointer) Elem() Type
```

Elem returns the element type for the given pointer p.

​	Elem 返回给定指针 p 的元素类型。

#### (*Pointer) String

```go
func (t *Pointer) String() string
```

#### (*Pointer) Underlying

```go
func (t *Pointer) Underlying() Type
```

### type Qualifier 

```go
type Qualifier func(*Package) string
```

A Qualifier controls how named package-level objects are printed in calls to TypeString, ObjectString, and SelectionString.

​	限定符控制在对 TypeString、ObjectString 和 SelectionString 调用时如何打印命名的包级对象。

These three formatting routines call the Qualifier for each package-level object O, and if the Qualifier returns a non-empty string p, the object is printed in the form p.O. If it returns an empty string, only the object name O is printed.

​	这三个格式化例程为每个包级对象 O 调用限定符，如果限定符返回非空字符串 p，则对象以 p.O 的形式打印。如果它返回空字符串，则只打印对象名称 O。

Using a nil Qualifier is equivalent to using (*Package).Path: the object is qualified by the import path, e.g., “encoding/json.Marshal”.

​	使用 nil 限定符等同于使用 (*Package).Path：对象由导入路径限定，例如，“encoding/json.Marshal”。

#### func RelativeTo

```go
func RelativeTo(pkg *Package) Qualifier
```

RelativeTo returns a Qualifier that fully qualifies members of all packages other than pkg.

​	RelativeTo 返回一个限定符，该限定符完全限定除 pkg 之外的所有包的成员。

### type Scope 

```go
type Scope struct {
	// contains filtered or unexported fields
}
```

A Scope maintains a set of objects and links to its containing (parent) and contained (children) scopes. Objects may be inserted and looked up by name. The zero value for Scope is a ready-to-use empty scope.

​	作用域维护一组对象及其包含的（父）和包含的（子）作用域的链接。可以按名称插入和查找对象。Scope 的零值是一个可立即使用的空作用域。

##### Example 

ExampleScope prints the tree of Scopes of a package created from a set of parsed files.

​	ExampleScope 打印由一组已解析文件创建的包的范围树。

```go
// Parse the source files for a package.
fset := token.NewFileSet()
var files []*ast.File
for _, src := range []string{
	`package main
import "fmt"
func main() {
	freezing := FToC(-18)
	fmt.Println(freezing, Boiling) }
`,
	`package main
import "fmt"
type Celsius float64
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func FToC(f float64) Celsius { return Celsius(f - 32 / 9 * 5) }
const Boiling Celsius = 100
func Unused() { {}; {{ var x int; _ = x }} } // make sure empty block scopes get printed
`,
} {
	files = append(files, mustParse(fset, src))
}

// Type-check a package consisting of these files.
// Type information for the imported "fmt" package
// comes from $GOROOT/pkg/$GOOS_$GOOARCH/fmt.a.
conf := types.Config{Importer: importer.Default()}
pkg, err := conf.Check("temperature", fset, files, nil)
if err != nil {
	log.Fatal(err)
}

// Print the tree of scopes.
// For determinism, we redact addresses.
var buf strings.Builder
pkg.Scope().WriteTo(&buf, 0, true)
rx := regexp.MustCompile(` 0x[a-fA-F\d]*`)
fmt.Println(rx.ReplaceAllString(buf.String(), ""))


Output:

package "temperature" scope {
.  const temperature.Boiling temperature.Celsius
.  type temperature.Celsius float64
.  func temperature.FToC(f float64) temperature.Celsius
.  func temperature.Unused()
.  func temperature.main()
.  main scope {
.  .  package fmt
.  .  function scope {
.  .  .  var freezing temperature.Celsius
.  .  }
.  }
.  main scope {
.  .  package fmt
.  .  function scope {
.  .  .  var c temperature.Celsius
.  .  }
.  .  function scope {
.  .  .  var f float64
.  .  }
.  .  function scope {
.  .  .  block scope {
.  .  .  }
.  .  .  block scope {
.  .  .  .  block scope {
.  .  .  .  .  var x int
.  .  .  .  }
.  .  .  }
.  .  }
.  }
}
var Universe *Scope
```

The Universe scope contains all predeclared objects of Go. It is the outermost scope of any chain of nested scopes.

​	Universe 范围包含 Go 的所有预声明对象。它是任何嵌套范围链的最外层范围。

#### func NewScope

```go
func NewScope(parent *Scope, pos, end token.Pos, comment string) *Scope
```

NewScope returns a new, empty scope contained in the given parent scope, if any. The comment is for debugging only.

​	NewScope 返回一个新的空范围，该范围包含在给定的父范围中（如果有）。注释仅用于调试。

#### (*Scope) Child

```go
func (s *Scope) Child(i int) *Scope
```

Child returns the i’th child scope for 0 <= i < NumChildren().

​	Child 返回 0 <= i < NumChildren() 的第 i 个子范围。

#### (*Scope) Contains

```go
func (s *Scope) Contains(pos token.Pos) bool
```

Contains reports whether pos is within the scope’s extent. The result is guaranteed to be valid only if the type-checked AST has complete position information.

​	Contains 报告 pos 是否在范围的范围内。仅当类型检查的 AST 具有完整的位置信息时，才能保证结果有效。

#### (*Scope) End

```go
func (s *Scope) End() token.Pos
```

#### (*Scope) Innermost

```go
func (s *Scope) Innermost(pos token.Pos) *Scope
```

Innermost returns the innermost (child) scope containing pos. If pos is not within any scope, the result is nil. The result is also nil for the Universe scope. The result is guaranteed to be valid only if the type-checked AST has complete position information.

​	Innermost 返回包含 pos 的最内部（子）作用域。如果 pos 不在任何作用域内，则结果为 nil。对于 Universe 作用域，结果也为 nil。仅当经过类型检查的 AST 具有完整的位置信息时，才能保证结果有效。

#### (*Scope) Insert 

```go
func (s *Scope) Insert(obj Object) Object
```

Insert attempts to insert an object obj into scope s. If s already contains an alternative object alt with the same name, Insert leaves s unchanged and returns alt. Otherwise it inserts obj, sets the object’s parent scope if not already set, and returns nil.

​	Insert 尝试将对象 obj 插入范围 s。如果 s 已包含具有相同名称的备用对象 alt，则 Insert 使 s 保持不变并返回 alt。否则，它会插入 obj，设置对象的父范围（如果尚未设置），并返回 nil。

#### (*Scope) Len 

```go
func (s *Scope) Len() int
```

Len returns the number of scope elements.

​	Len 返回作用域元素的数量。

#### (*Scope) Lookup 

```go
func (s *Scope) Lookup(name string) Object
```

Lookup returns the object in scope s with the given name if such an object exists; otherwise the result is nil.

​	Lookup 返回作用域 s 中具有给定名称的对象（如果存在这样的对象）；否则结果为 nil。

#### (*Scope) LookupParent 

```go
func (s *Scope) LookupParent(name string, pos token.Pos) (*Scope, Object)
```

LookupParent follows the parent chain of scopes starting with s until it finds a scope where Lookup(name) returns a non-nil object, and then returns that scope and object. If a valid position pos is provided, only objects that were declared at or before pos are considered. If no such scope and object exists, the result is (nil, nil).

​	LookupParent 遵循从 s 开始的范围父链，直到找到一个范围，其中 Lookup(name) 返回一个非空对象，然后返回该范围和对象。如果提供了有效位置 pos，则只考虑在 pos 处或之前声明的对象。如果不存在这样的范围和对象，则结果为 (nil, nil)。

Note that obj.Parent() may be different from the returned scope if the object was inserted into the scope and already had a parent at that time (see Insert). This can only happen for dot-imported objects whose scope is the scope of the package that exported them.

​	请注意，如果对象被插入到该范围中并且当时已经有一个父对象，则 obj.Parent() 可能与返回的范围不同（请参阅 Insert）。这仅适用于点导入的对象，其范围是导出它们的包的范围。

#### (*Scope) Names 

```go
func (s *Scope) Names() []string
```

Names returns the scope’s element names in sorted order.

​	Names 返回按排序顺序排列的范围的元素名称。

#### (*Scope) NumChildren 

```go
func (s *Scope) NumChildren() int
```

NumChildren returns the number of scopes nested in s.

​	NumChildren 返回 s 中嵌套的范围数。

#### (*Scope) Parent

```go
func (s *Scope) Parent() *Scope
```

Parent returns the scope’s containing (parent) scope.

​	Parent 返回范围的包含（父）范围。

#### (*Scope) Pos 

```go
func (s *Scope) Pos() token.Pos
```

Pos and End describe the scope’s source code extent [pos, end). The results are guaranteed to be valid only if the type-checked AST has complete position information. The extent is undefined for Universe and package scopes.

​	Pos 和 End 描述范围的源代码范围 [pos, end)。仅当经过类型检查的 AST 具有完整的位置信息时，才能保证结果有效。Universe 和包范围的范围是未定义的。

#### (*Scope) String 

```go
func (s *Scope) String() string
```

String returns a string representation of the scope, for debugging.

​	String 返回一个字符串表示形式的范围，用于调试。

#### (*Scope) WriteTo 

```go
func (s *Scope) WriteTo(w io.Writer, n int, recurse bool)
```

WriteTo writes a string representation of the scope to w, with the scope elements sorted by name. The level of indentation is controlled by n >= 0, with n == 0 for no indentation. If recurse is set, it also writes nested (children) scopes.

​	WriteTo 将作用域的字符串表示形式写入 w，作用域元素按名称排序。缩进级别由 n >= 0 控制，n == 0 表示不缩进。如果设置了 recurse，它还会写入嵌套（子）作用域。

### type Selection 

```go
type Selection struct {
	// contains filtered or unexported fields
}
```

A Selection describes a selector expression x.f. For the declarations:

​	选择描述选择器表达式 x.f。对于声明：

```go
type T struct{ x int; E }
type E struct{}
func (e E) m() {}
var p *T
```

the following relations exist:

​	存在以下关系：

```
Selector    Kind          Recv    Obj    Type       Index     Indirect

p.x         FieldVal      T       x      int        {0}       true
p.m         MethodVal     *T      m      func()     {1, 0}    true
T.m         MethodExpr    T       m      func(T)    {1, 0}    false
```

#### (*Selection) Index

```go
func (s *Selection) Index() []int
```

Index describes the path from x to f in x.f. The last index entry is the field or method index of the type declaring f; either:

​	索引描述了 x.f 中从 x 到 f 的路径。最后一个索引条目是声明 f 的类型的字段或方法索引；或者：

1. the list of declared methods of a named type; or
   已命名类型的已声明方法列表；或
2. the list of methods of an interface type; or
   接口类型的列表；或
3. the list of fields of a struct type.
   结构类型的字段列表。

The earlier index entries are the indices of the embedded fields implicitly traversed to get from (the type of) x to f, starting at embedding depth 0.

​	较早的索引条目是从（x 的类型）到 f 的隐式遍历的嵌入字段的索引，从嵌入深度 0 开始。

#### (*Selection) Indirect 

```go
func (s *Selection) Indirect() bool
```

Indirect reports whether any pointer indirection was required to get from x to f in x.f.

​	间接报告是否需要任何指针间接才能从 x.f 中的 x 获取 f。

#### (*Selection) Kind

```go
func (s *Selection) Kind() SelectionKind
```

Kind returns the selection kind.

​	Kind 返回选择类型。

#### (*Selection) Obj

```go
func (s *Selection) Obj() Object
```

Obj returns the object denoted by x.f; a *Var for a field selection, and a *Func in all other cases.

​	Obj 返回 x.f 表示的对象；对于字段选择，返回 *Var，在所有其他情况下返回 *Func。

#### (*Selection) Recv 

```go
func (s *Selection) Recv() Type
```

Recv returns the type of x in x.f.

​	Recv 返回 x.f 中 x 的类型。

#### (*Selection) String 

```go
func (s *Selection) String() string
```

#### (*Selection) Type

```go
func (s *Selection) Type() Type
```

Type returns the type of x.f, which may be different from the type of f. See Selection for more information.

​	Type 返回 x.f 的类型，该类型可能与 f 的类型不同。有关更多信息，请参阅 Selection。

### type SelectionKind

```go
type SelectionKind int
```

SelectionKind describes the kind of a selector expression x.f (excluding qualified identifiers).

​	SelectionKind 描述选择器表达式 x.f 的类型（不包括限定标识符）。

```go
const (
	FieldVal   SelectionKind = iota // x.f is a struct field selector
	MethodVal                       // x.f is a method selector
	MethodExpr                      // x.f is a method expression
)
```

### type Signature

```go
type Signature struct {
	// contains filtered or unexported fields
}
```

A Signature represents a (non-builtin) function or method type. The receiver is ignored when comparing signatures for identity.

​	签名表示（非内置）函数或方法类型。在比较签名以确定标识时，忽略接收器。

#### func NewSignature <-DEPRECATED

```go
func NewSignature(recv *Var, params, results *Tuple, variadic bool) *Signature
```

NewSignature returns a new function type for the given receiver, parameters, and results, either of which may be nil. If variadic is set, the function is variadic, it must have at least one parameter, and the last parameter must be of unnamed slice type. 重试 错误原因

Deprecated: Use NewSignatureType instead which allows for type parameters. 重试 错误原因

#### func NewSignatureType <- go1.18

```go
func NewSignatureType(recv *Var, recvTypeParams, typeParams []*TypeParam, params, results *Tuple, variadic bool) *Signature
```

NewSignatureType creates a new function type for the given receiver, receiver type parameters, type parameters, parameters, and results. If variadic is set, params must hold at least one parameter and the last parameter’s core type must be of unnamed slice or bytestring type. If recv is non-nil, typeParams must be empty. If recvTypeParams is non-empty, recv must be non-nil.

​	NewSignatureType 为给定的接收器、接收器类型参数、类型参数、参数和结果创建一个新的函数类型。如果设置了 variadic，则 params 必须至少包含一个参数，并且最后一个参数的核心类型必须是未命名切片或字节串类型。如果 recv 为非 nil，则 typeParams 必须为空。如果 recvTypeParams 为非空，则 recv 必须为非 nil。

#### (*Signature) Params 

```go
func (s *Signature) Params() *Tuple
```

Params returns the parameters of signature s, or nil.

​	Params 返回签名 s 的参数，或 nil。

#### (*Signature) Recv

```go
func (s *Signature) Recv() *Var
```

Recv returns the receiver of signature s (if a method), or nil if a function. It is ignored when comparing signatures for identity.

​	Recv 返回签名 s 的接收者（如果是方法）或 nil（如果是函数）。在比较签名以确定标识时，它会被忽略。

For an abstract method, Recv returns the enclosing interface either as a *Named or an *Interface. Due to embedding, an interface may contain methods whose receiver type is a different interface.

​	对于抽象方法，Recv 返回封闭接口，作为 *Named 或 *Interface。由于嵌入，接口可能包含其接收者类型是不同接口的方法。

#### (*Signature) RecvTypeParams <- go1.18

```go
func (s *Signature) RecvTypeParams() *TypeParamList
```

RecvTypeParams returns the receiver type parameters of signature s, or nil.

​	RecvTypeParams 返回签名 s 的接收器类型参数，或 nil。

#### (*Signature) Results 

```go
func (s *Signature) Results() *Tuple
```

Results returns the results of signature s, or nil.

​	Results 返回签名 s 的结果，或 nil。

#### (*Signature) String 

```go
func (t *Signature) String() string
```

#### (*Signature) TypeParams <- go1.18 

```go
func (s *Signature) TypeParams() *TypeParamList
```

TypeParams returns the type parameters of signature s, or nil.

​	TypeParams 返回签名 s 的类型参数，或 nil。

#### (*Signature) Underlying 

```go
func (t *Signature) Underlying() Type
```

#### (*Signature) Variadic

```go
func (s *Signature) Variadic() bool
```

Variadic reports whether the signature s is variadic.

​	Variadic 报告签名 s 是否是变参的。

### type Sizes

```go
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

​	Sizes 定义包 unsafe 的大小调整函数。

#### func SizesFor <- go1.9

```go
func SizesFor(compiler, arch string) Sizes
```

SizesFor returns the Sizes used by a compiler for an architecture. The result is nil if a compiler/architecture pair is not known.

​	SizesFor 返回编译器针对某个架构使用的 Sizes。如果编译器/架构对未知，则结果为 nil。

Supported architectures for compiler “gc”: “386”, “amd64”, “amd64p32”, “arm”, “arm64”, “loong64”, “mips”, “mipsle”, “mips64”, “mips64le”, “ppc64”, “ppc64le”, “riscv64”, “s390x”, “sparc64”, “wasm”.

​	编译器“gc”支持的架构：“386”、“amd64”、“amd64p32”、“arm”、“arm64”、“loong64”、“mips”、“mipsle”、“mips64”、“mips64le”、“ppc64”、“ppc64le”、“riscv64”、“s390x”、“sparc64”、“wasm”。

### type Slice

```go
type Slice struct {
	// contains filtered or unexported fields
}
```

A Slice represents a slice type.

​	A Slice 表示一个切片类型。

#### func NewSlice

```go
func NewSlice(elem Type) *Slice
```

NewSlice returns a new slice type for the given element type.

​	NewSlice 返回给定元素类型的新切片类型。

#### (*Slice) Elem

```go
func (s *Slice) Elem() Type
```

Elem returns the element type of slice s. 重试 错误原因

#### (*Slice) String

```go
func (t *Slice) String() string
```

#### (*Slice) Underlying 

```go
func (t *Slice) Underlying() Type
```

### type StdSizes

```go
type StdSizes struct {
	WordSize int64 // word size in bytes - must be >= 4 (32bits)
	MaxAlign int64 // maximum alignment in bytes - must be >= 1
}
```

StdSizes is a convenience type for creating commonly used Sizes. It makes the following simplifying assumptions:

​	StdSizes 是一种用于创建常用大小的便捷类型。它做出了以下简化假设：

- The size of explicitly sized basic types (int16, etc.) is the specified size.
  显式大小的基本类型（int16 等）的大小是指定的大小。
- The size of strings and interfaces is 2*WordSize.
  字符串和接口的大小是 2*WordSize。
- The size of slices is 3*WordSize.
  切片的大小是 3*WordSize。
- The size of an array of n elements corresponds to the size of a struct of n consecutive fields of the array’s element type.
  包含 n 个元素的数组的大小对应于 n 个连续字段的结构的大小，这些字段属于数组的元素类型。
- The size of a struct is the offset of the last field plus that field’s size. As with all element types, if the struct is used in an array its size must first be aligned to a multiple of the struct’s alignment.
  结构的大小是最后一个字段的偏移量加上该字段的大小。与所有元素类型一样，如果结构用于数组，则必须先将其大小调整为结构对齐方式的倍数。
- All other types have size WordSize.
  所有其他类型的大小为 WordSize。
- Arrays and structs are aligned per spec definition; all other types are naturally aligned with a maximum alignment MaxAlign.
  数组和结构按照规范定义对齐；所有其他类型都以最大对齐方式 MaxAlign 自然对齐。

`*StdSizes` implements Sizes.

​	`*StdSizes` 实现 Sizes。

#### (*StdSizes) Alignof

```go
func (s *StdSizes) Alignof(T Type) int64
```

#### (*StdSizes) Offsetsof

```go
func (s *StdSizes) Offsetsof(fields []*Var) []int64
```

#### (*StdSizes) Sizeof

```go
func (s *StdSizes) Sizeof(T Type) int64
```

### type Struct

```go
type Struct struct {
	// contains filtered or unexported fields
}
```

A Struct represents a struct type.

​	结构体表示一个结构体类型。

#### func NewStruct

```go
func NewStruct(fields []*Var, tags []string) *Struct
```

NewStruct returns a new struct with the given fields and corresponding field tags. If a field with index i has a tag, tags[i] must be that tag, but len(tags) may be only as long as required to hold the tag with the largest index i. Consequently, if no field has a tag, tags may be nil.

​	NewStruct 返回具有给定字段和相应字段标记的新结构。如果索引为 i 的字段具有标记，则 tags[i] 必须是该标记，但 len(tags) 可能仅与容纳具有最大索引 i 的标记所需的一样长。因此，如果没有字段具有标记，则 tags 可能为 nil。

#### (*Struct) Field

```go
func (s *Struct) Field(i int) *Var
```

Field returns the i’th field for 0 <= i < NumFields(). 重试 错误原因

#### (*Struct) NumFields

```go
func (s *Struct) NumFields() int
```

NumFields returns the number of fields in the struct (including blank and embedded fields).

​	NumFields 返回结构中的字段数（包括空白字段和嵌入字段）。

#### (*Struct) String

```go
func (t *Struct) String() string
```

#### (*Struct) Tag

```go
func (s *Struct) Tag(i int) string
```

Tag returns the i’th field tag for 0 <= i < NumFields().

​	标签返回 0 <= i < NumFields() 的第 i 个字段标签。

#### (*Struct) Underlying

```go
func (t *Struct) Underlying() Type
```

### type Term <- go1.18

```go
type Term term
```

A Term represents a term in a Union.

​	术语表示联合中的一个术语。

#### func NewTerm <- go1.18

```go
func NewTerm(tilde bool, typ Type) *Term
```

NewTerm returns a new union term.

​	NewTerm 返回一个新的联合项。

#### (*Term) String <- go1.18

```go
func (t *Term) String() string
```

#### (*Term) Tilde <- go1.18

```go
func (t *Term) Tilde() bool
```

#### (*Term) Type <- go1.18

```go
func (t *Term) Type() Type
```

### type Tuple

```go
type Tuple struct {
	// contains filtered or unexported fields
}
```

A Tuple represents an ordered list of variables; a nil *Tuple is a valid (empty) tuple. Tuples are used as components of signatures and to represent the type of multiple assignments; they are not first class types of Go.

​	元组表示一个有序变量列表；nil *Tuple 是一个有效的（空）元组。元组用作签名组件和表示多个赋值的类型；它们不是 Go 的一等类型。

#### func NewTuple

```go
func NewTuple(x ...*Var) *Tuple
```

NewTuple returns a new tuple for the given variables.

​	NewTuple 返回给定变量的新元组。

#### (*Tuple) At 

```go
func (t *Tuple) At(i int) *Var
```

At returns the i’th variable of tuple t.

​	At 返回元组 t 的第 i 个变量。

#### (*Tuple) Len 

```go
func (t *Tuple) Len() int
```

Len returns the number variables of tuple t.

​	Len 返回元组 t 的变量数。

#### (*Tuple) String 

```go
func (t *Tuple) String() string
```

#### (*Tuple) Underlying 

```go
func (t *Tuple) Underlying() Type
```

### type Type

```go
type Type interface {
	// Underlying returns the underlying type of a type.
	Underlying() Type

	// String returns a string representation of a type.
	String() string
}
```

A Type represents a type of Go. All types implement the Type interface.

​	类型表示 Go 的一种类型。所有类型都实现了 Type 接口。

#### func Default <- go1.8

```go
func Default(t Type) Type
```

Default returns the default “typed” type for an “untyped” type; it returns the incoming type for all other types. The default type for untyped nil is untyped nil.

​	Default 返回“无类型”类型的默认“类型化”类型；它返回所有其他类型的传入类型。无类型 nil 的默认类型是无类型 nil。

#### func Instantiate <- go1.18

```go
func Instantiate(ctxt *Context, orig Type, targs []Type, validate bool) (Type, error)
```

Instantiate instantiates the type orig with the given type arguments targs. orig must be a `*Named` or a `*Signature` type. If there is no error, the resulting Type is an instantiated type of the same kind (either a `*Named` or a `*Signature`). Methods attached to a `*Named` type are also instantiated, and associated with a new *Func that has the same position as the original method, but nil function scope.

​	Instantiate 用给定的类型参数 targs 实例化类型 orig。orig 必须是 `*Named` 或 `*Signature` 类型。如果没有错误，则结果类型是相同类型的实例化类型（`*Named` 或 `*Signature`）。附加到 `*Named` 类型的函数也会被实例化，并与具有与原始函数相同位置的新 `*Func` 相关联，但函数范围为 nil。

If ctxt is non-nil, it may be used to de-duplicate the instance against previous instances with the same identity. As a special case, generic *Signature origin types are only considered identical if they are pointer equivalent, so that instantiating distinct (but possibly identical) signatures will yield different instances. The use of a shared context does not guarantee that identical instances are deduplicated in all cases.

​	如果 ctxt 为非 nil，则可以使用它来根据具有相同标识的先前实例对实例进行去重。作为特例，仅当通用 `*Signature` 原始类型指针等价时，才认为它们是相同的，因此实例化不同的（但可能相同的）签名将产生不同的实例。使用共享上下文并不能保证在所有情况下都对相同的实例进行去重。

If validate is set, Instantiate verifies that the number of type arguments and parameters match, and that the type arguments satisfy their corresponding type constraints. If verification fails, the resulting error may wrap an *ArgumentError indicating which type argument did not satisfy its corresponding type parameter constraint, and why.

​	如果设置了 validate，则 Instantiate 会验证类型参数和参数的数量是否匹配，以及类型参数是否满足其对应的类型约束。如果验证失败，则结果错误可能会包装一个 *ArgumentError，指示哪个类型参数不满足其对应的类型参数约束，以及原因。

If validate is not set, Instantiate does not verify the type argument count or whether the type arguments satisfy their constraints. Instantiate is guaranteed to not return an error, but may panic. Specifically, for *Signature types, Instantiate will panic immediately if the type argument count is incorrect; for *Named types, a panic may occur later inside the *Named API.

​	如果未设置验证，Instantiate 不会验证类型参数计数或类型参数是否满足其约束。Instantiate 保证不会返回错误，但可能会引发 panic。具体来说，对于 *Signature 类型，如果类型参数计数不正确，Instantiate 会立即引发 panic；对于 *Named 类型，panic 可能会稍后在 *Named API 内部发生。

### type TypeAndValue 

```go
type TypeAndValue struct {
	Type  Type
	Value constant.Value
	// contains filtered or unexported fields
}
```

TypeAndValue reports the type and value (for constants) of the corresponding expression.

​	TypeAndValue 报告相应表达式的类型和值（对于常量）。

#### func Eval

```go
func Eval(fset *token.FileSet, pkg *Package, pos token.Pos, expr string) (_ TypeAndValue, err error)
```

Eval returns the type and, if constant, the value for the expression expr, evaluated at position pos of package pkg, which must have been derived from type-checking an AST with complete position information relative to the provided file set.

​	Eval 返回类型，如果常量，则返回在包 pkg 的位置 pos 处计算的表达式 expr 的值，该包必须是从相对于所提供的文件集的完整位置信息进行类型检查的 AST 派生的。

The meaning of the parameters fset, pkg, and pos is the same as in CheckExpr. An error is returned if expr cannot be parsed successfully, or the resulting expr AST cannot be type-checked.

​	fset、pkg 和 pos 参数的含义与 CheckExpr 中的相同。如果无法成功解析 expr，或者无法对生成的 expr AST 进行类型检查，则会返回错误。

#### (TypeAndValue) Addressable

```go
func (tv TypeAndValue) Addressable() bool
```

Addressable reports whether the corresponding expression is addressable (https://golang.org/ref/spec#Address_operators).

​	Addressable 报告相应的表达式是否可寻址 ( https://golang.org/ref/spec#Address_operators)。

#### (TypeAndValue) Assignable

```go
func (tv TypeAndValue) Assignable() bool
```

Assignable reports whether the corresponding expression is assignable to (provided a value of the right type).

​	Assignable 报告相应的表达式是否可赋值给（提供一个正确类型的数值）。

#### (TypeAndValue) HasOk

```go
func (tv TypeAndValue) HasOk() bool
```

HasOk reports whether the corresponding expression may be used on the rhs of a comma-ok assignment.

​	HasOk 报告相应的表达式是否可以在逗号-ok 赋值的右侧使用。

#### (TypeAndValue) IsBuiltin

```go
func (tv TypeAndValue) IsBuiltin() bool
```

IsBuiltin reports whether the corresponding expression denotes a (possibly parenthesized) built-in function.

​	IsBuiltin 报告相应的表达式是否表示（可能带括号的）内置函数。

#### (TypeAndValue) IsNil

```go
func (tv TypeAndValue) IsNil() bool
```

IsNil reports whether the corresponding expression denotes the predeclared value nil.

​	IsNil 报告相应的表达式是否表示预先声明的值 nil。

#### (TypeAndValue) IsType

```go
func (tv TypeAndValue) IsType() bool
```

IsType reports whether the corresponding expression specifies a type.

​	IsType 报告相应的表达式是否指定了一个类型。

#### (TypeAndValue) IsValue

```go
func (tv TypeAndValue) IsValue() bool
```

IsValue reports whether the corresponding expression is a value. Builtins are not considered values. Constant values have a non- nil Value.

​	IsValue 报告相应的表达式是否为值。内置函数不被视为值。常量值具有非 nil 值。

#### (TypeAndValue) IsVoid

```go
func (tv TypeAndValue) IsVoid() bool
```

IsVoid reports whether the corresponding expression is a function call without results.

​	IsVoid 报告相应的表达式是否是一个没有结果的函数调用。

### type TypeList <- go1.18

```go
type TypeList struct {
	// contains filtered or unexported fields
}
```

TypeList holds a list of types.

​	TypeList 持有一系列类型。

#### (*TypeList) At <- go1.18

```go
func (l *TypeList) At(i int) Type
```

At returns the i’th type in the list.

​	At 返回列表中的第 i 个类型。

#### (*TypeList) Len <- go1.18

```go
func (l *TypeList) Len() int
```

Len returns the number of types in the list. It is safe to call on a nil receiver.

​	Len 返回列表中的类型数。对 nil 接收器调用是安全的。

### type TypeName

```go
type TypeName struct {
	// contains filtered or unexported fields
}
```

A TypeName represents a name for a (defined or alias) type.

​	TypeName 表示（已定义或别名）类型的名称。

#### func NewTypeName

```go
func NewTypeName(pos token.Pos, pkg *Package, name string, typ Type) *TypeName
```

NewTypeName returns a new type name denoting the given typ. The remaining arguments set the attributes found with all Objects.

​	NewTypeName 返回一个新的类型名称，表示给定的 typ。其余参数设置在所有对象中找到的属性。

The typ argument may be a defined (Named) type or an alias type. It may also be nil such that the returned TypeName can be used as argument for NewNamed, which will set the TypeName’s type as a side- effect.

​	typ 参数可以是已定义的（命名）类型或别名类型。它也可以是 nil，以便返回的 TypeName 可用作 NewNamed 的参数，这将把 TypeName 的类型设置为副作用。

#### (*TypeName) Exported

```go
func (obj *TypeName) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn’t take into account whether the object is in a local (function) scope or not.

​	Exported 报告对象是否已导出（以大写字母开头）。它不考虑对象是否在本地（函数）范围内。

#### (*TypeName) Id

```go
func (obj *TypeName) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

​	Id 是 Id(obj.Pkg(), obj.Name()) 的包装器。

#### (*TypeName) IsAlias <- go1.9

```go
func (obj *TypeName) IsAlias() bool
```

IsAlias reports whether obj is an alias name for a type.

​	IsAlias 报告 obj 是否是类型的别名。

#### (*TypeName) Name

```go
func (obj *TypeName) Name() string
```

Name returns the object’s (package-local, unqualified) name.

​	Name 返回对象（包本地、不合格）的名称。

#### (*TypeName) Parent

```go
func (obj *TypeName) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

​	Parent 返回声明对象的作用域。对于方法和结构字段，结果为 nil。

#### (*TypeName) Pkg

```go
func (obj *TypeName) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

​	Pkg 返回对象所属的包。对于 Universe 范围内的标签和对象，结果为 nil。

#### (*TypeName) Pos

```go
func (obj *TypeName) Pos() token.Pos
```

Pos returns the declaration position of the object’s identifier.

​	Pos 返回对象标识符的声明位置。

#### (*TypeName) String

```go
func (obj *TypeName) String() string
```

#### (*TypeName) Type

```go
func (obj *TypeName) Type() Type
```

Type returns the object’s type.

​	Type 返回对象类型。

### type TypeParam <- go1.18

```go
type TypeParam struct {
	// contains filtered or unexported fields
}
```

A TypeParam represents a type parameter type.

​	A TypeParam 表示类型参数类型。

#### func NewTypeParam <- go1.18

```go
func NewTypeParam(obj *TypeName, constraint Type) *TypeParam
```

NewTypeParam returns a new TypeParam. Type parameters may be set on a Named or Signature type by calling SetTypeParams. Setting a type parameter on more than one type will result in a panic.

​	NewTypeParam 返回一个新的 TypeParam。可以通过调用 SetTypeParams 在 Named 或 Signature 类型上设置类型参数。在多个类型上设置类型参数将导致恐慌。

The constraint argument can be nil, and set later via SetConstraint. If the constraint is non-nil, it must be fully defined.

​	约束参数可以为 nil，并通过 SetConstraint 稍后设置。如果约束不是 nil，则必须完全定义。

#### (*TypeParam) Constraint <- go1.18

```go
func (t *TypeParam) Constraint() Type
```

Constraint returns the type constraint specified for t.

​	约束返回为 t 指定的类型约束。

#### (*TypeParam) Index <- go1.18

```go
func (t *TypeParam) Index() int
```

Index returns the index of the type param within its param list, or -1 if the type parameter has not yet been bound to a type.

​	Index 返回类型参数在其参数列表中的索引，如果类型参数尚未绑定到类型，则返回 -1。

#### (*TypeParam) Obj <- go1.18

```go
func (t *TypeParam) Obj() *TypeName
```

Obj returns the type name for t.

​	Obj 返回 t 的类型名称。

#### (*TypeParam) SetConstraint <- go1.18

```go
func (t *TypeParam) SetConstraint(bound Type)
```

SetConstraint sets the type constraint for t.

​	SetConstraint 为 t 设置类型约束。

It must be called by users of NewTypeParam after the bound’s underlying is fully defined, and before using the type parameter in any way other than to form other types. Once SetConstraint returns the receiver, t is safe for concurrent use.

​	在完全定义了 bound 的基础之后，并且在以形成其他类型以外的任何方式使用类型参数之前，NewTypeParam 的用户必须调用它。一旦 SetConstraint 返回接收器，t 就可安全地并发使用。

#### (*TypeParam) String <- go1.18 

```go
func (t *TypeParam) String() string
```

#### (*TypeParam) Underlying <- go1.18

```go
func (t *TypeParam) Underlying() Type
```

### type TypeParamList <- go1.18

```go
type TypeParamList struct {
	// contains filtered or unexported fields
}
```

TypeParamList holds a list of type parameters.

​	TypeParamList 包含一个类型参数列表。

#### (*TypeParamList) At <- go1.18

```go
func (l *TypeParamList) At(i int) *TypeParam
```

At returns the i’th type parameter in the list.

​	At 返回列表中的第 i 个类型参数。

#### (*TypeParamList) Len <- go1.18

```go
func (l *TypeParamList) Len() int
```

Len returns the number of type parameters in the list. It is safe to call on a nil receiver.

​	Len 返回列表中的类型参数数量。对 nil 接收器调用是安全的。

### type Union <- go1.18

```go
type Union struct {
	// contains filtered or unexported fields
}
```

A Union represents a union of terms embedded in an interface.

​	联合表示嵌入在接口中的项的联合。

#### func NewUnion <- go1.18

```go
func NewUnion(terms []*Term) *Union
```

NewUnion returns a new Union type with the given terms. It is an error to create an empty union; they are syntactically not possible.

​	NewUnion 返回具有给定项的新 Union 类型。创建空联合是一种错误；它们在语法上是不可能的。

#### (*Union) Len <- go1.18 

```go
func (u *Union) Len() int
```

#### (*Union) String <- go1.18 

```go
func (u *Union) String() string
```

#### (*Union) Term <- go1.18 

```go
func (u *Union) Term(i int) *Term
```

#### (*Union) Underlying <- go1.18 

```go
func (u *Union) Underlying() Type
```

### type Var

```go
type Var struct {
	// contains filtered or unexported fields
}
```

A Variable represents a declared variable (including function parameters and results, and struct fields).

​	变量表示已声明的变量（包括函数参数和结果，以及结构字段）。

#### func NewField

```go
func NewField(pos token.Pos, pkg *Package, name string, typ Type, embedded bool) *Var
```

NewField returns a new variable representing a struct field. For embedded fields, the name is the unqualified type name under which the field is accessible.

​	NewField 返回一个表示结构字段的新变量。对于嵌入字段，名称是字段可访问的非限定类型名称。

#### func NewParam

```go
func NewParam(pos token.Pos, pkg *Package, name string, typ Type) *Var
```

NewParam returns a new variable representing a function parameter.

​	NewParam 返回一个表示函数参数的新变量。

#### func NewVar

```go
func NewVar(pos token.Pos, pkg *Package, name string, typ Type) *Var
```

NewVar returns a new variable. The arguments set the attributes found with all Objects.

​	NewVar 返回一个新变量。参数设置所有对象中发现的属性。

#### (*Var) Anonymous

```go
func (obj *Var) Anonymous() bool
```

Anonymous reports whether the variable is an embedded field. Same as Embedded; only present for backward-compatibility.

​	Anonymous 报告变量是否为嵌入字段。与 Embedded 相同；仅出于向后兼容性而存在。

#### (*Var) Embedded <- go1.11

```go
func (obj *Var) Embedded() bool
```

Embedded reports whether the variable is an embedded field.

​	Embedded 报告变量是否为嵌入字段。

#### (*Var) Exported

```go
func (obj *Var) Exported() bool
```

Exported reports whether the object is exported (starts with a capital letter). It doesn’t take into account whether the object is in a local (function) scope or not.

​	Exported 报告对象是否已导出（以大写字母开头）。它不考虑对象是否在本地（函数）作用域中。

#### (*Var) Id 

```go
func (obj *Var) Id() string
```

Id is a wrapper for Id(obj.Pkg(), obj.Name()).

​	Id 是 Id(obj.Pkg(), obj.Name()) 的包装器。

#### (*Var) IsField

```go
func (obj *Var) IsField() bool
```

IsField reports whether the variable is a struct field.

​	IsField 报告变量是否为结构字段。

#### (*Var) Name

```go
func (obj *Var) Name() string
```

Name returns the object’s (package-local, unqualified) name.

​	Name 返回对象（包本地、不合格）的名称。

#### (*Var) Origin <- go1.19

```go
func (obj *Var) Origin() *Var
```

Origin returns the canonical Var for its receiver, i.e. the Var object recorded in Info.Defs.

​	Origin 返回其接收器的规范 Var，即 Info.Defs 中记录的 Var 对象。

For synthetic Vars created during instantiation (such as struct fields or function parameters that depend on type arguments), this will be the corresponding Var on the generic (uninstantiated) type. For all other Vars Origin returns the receiver.

​	对于在实例化期间创建的合成 Var（例如，取决于类型参数的结构字段或函数参数），这将是泛型（未实例化）类型上的相应 Var。对于所有其他 Var，Origin 返回接收器。

#### (*Var) Parent 

```go
func (obj *Var) Parent() *Scope
```

Parent returns the scope in which the object is declared. The result is nil for methods and struct fields.

​	Parent 返回声明对象的作用域。对于方法和结构字段，结果为 nil。

#### (*Var) Pkg

```go
func (obj *Var) Pkg() *Package
```

Pkg returns the package to which the object belongs. The result is nil for labels and objects in the Universe scope.

​	Pkg 返回对象所属的包。对于 Universe 范围内的标签和对象，结果为 nil。

#### (*Var) Pos

```go
func (obj *Var) Pos() token.Pos
```

Pos returns the declaration position of the object’s identifier.

​	Pos 返回对象标识符的声明位置。

#### (*Var) String

```go
func (obj *Var) String() string
```

#### (*Var) Type 

```go
func (obj *Var) Type() Type
```

Type returns the object’s type.

​	Type 返回对象类型。