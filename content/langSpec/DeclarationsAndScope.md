+++
title = "声明和作用域"
date = 2023-05-17T09:59:21+08:00
weight = 10
description = ""
isCJKLanguage = true
type = "docs"
draft = false
+++
## Declarations and scope 声明和作用域

> 原文：[https://go.dev/ref/spec#Declarations_and_scope](https://go.dev/ref/spec#Declarations_and_scope)

A *declaration* binds a non-[blank](https://go.dev/ref/spec#Blank_identifier) identifier to a [constant](https://go.dev/ref/spec#Constant_declarations), [type](https://go.dev/ref/spec#Type_declarations), [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), [variable](https://go.dev/ref/spec#Variable_declarations), [function](https://go.dev/ref/spec#Function_declarations), [label](https://go.dev/ref/spec#Labeled_statements), or [package](https://go.dev/ref/spec#Import_declarations). Every identifier in a program must be declared. No identifier may be declared twice in the same block, and no identifier may be declared in both the file and package block.

​	声明将一个非[空白](#blank-identifier-空白标识符)的标识符绑定到一个[常量](../Constants)、[类型](../Types)、[类型参数](#type-parameter-declarations-类型参数声明)、[变量](../Variables)、[函数](#function-declarations-函数声明)、[标签](../Statements#labeled-statements-标签语句)或[包](../Pachages#import-declarations-导入声明)。程序中的每个标识符都必须被声明。没有一个标识符可以在同一个块中声明两次，也没有一个标识符可以同时在文件块和包块中声明。

The [blank identifier](https://go.dev/ref/spec#Blank_identifier) may be used like any other identifier in a declaration, but it does not introduce a binding and thus is not declared. In the package block, the identifier `init` may only be used for [`init` function](https://go.dev/ref/spec#Package_initialization) declarations, and like the blank identifier it does not introduce a new binding.

​	[空白标识符（即`_`）](#blank-identifier-空白标识符)可以像其他标识符一样在声明中使用，但它不引入绑定，因此不被声明。在包块中，标识符`init`只能用于[init函数](../ProgramInitializationAndExecution#package-initialization-包的初始化)的声明，和空白标识符一样，它不会引入一个新的绑定。

```
Declaration   = ConstDecl | TypeDecl | VarDecl .
TopLevelDecl  = Declaration | FunctionDecl | MethodDecl .
```

The *scope* of a declared identifier is the extent of source text in which the identifier denotes the specified constant, type, variable, function, label, or package.

​	声明标识符的作用域是源文本的范围，其中标识符表示指定的常量、类型、变量、函数、标签或包。

Go is lexically scoped using [blocks](https://go.dev/ref/spec#Blocks):

​	Go在词法使用[块](../Blocks)来确定作用域：

1. The scope of a [predeclared identifier](https://go.dev/ref/spec#Predeclared_identifiers) is the universe block.
2. [预先声明的标识符](../DeclarationsAndScope#predeclared-identifiers--预先声明的标识符)的作用域是`universe block`。
3. The scope of an identifier denoting a constant, type, variable, or function (but not method) declared at top level (outside any function) is the package block. 
4. 表示常量、类型、变量或函数（但不是方法）的标识符在顶层（任何函数之外）声明的作用域是`package block`。
5. The scope of an identifier denoting a method receiver, function parameter, or result variable is the function body.
6. 导入的包的包名的作用域是在导入声明的文件的`file block`。
7. The scope of an identifier denoting a type parameter of a function or declared by a method receiver begins after the name of the function and ends at the end of the function body.
8. 表示方法接收器、函数参数或结果变量的标识符的作用域是函数体。
9. The scope of an identifier denoting a type parameter of a type begins after the name of the type and ends at the end of the TypeSpec.
10. 表示函数的类型参数或由方法接收器声明的标识符的作用域从函数名称之后开始，在函数体的末端结束。
11. The scope of an identifier denoting a type parameter of a type begins after the name of the type and ends at the end of the TypeSpec.
12. 表示类型的类型参数的标识符的作用域从类型的名称之后开始，在TypeSpec的末尾结束。=>仍有疑问？？
13. The scope of a constant or variable identifier declared inside a function begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block.
14. 在函数中声明的常量或变量标识符的作用域从ConstSpec或VarSpec（ShortVarDecl用于短变量声明）的末尾开始，在最里面的包含块的末尾结束。
15. The scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.
16. 在函数中声明的类型标识符的作用域从TypeSpec中的标识符开始，在最内层的包含块的末端结束。

An identifier declared in a block may be redeclared in an inner block. While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration.

​	在块中声明的标识符可以在内部块中被重新声明。当内部声明的标识符在作用域内时，它表示内部声明所声明的实体。

The [package clause](https://go.dev/ref/spec#Package_clause) is not a declaration; the package name does not appear in any scope. Its purpose is to identify the files belonging to the same [package](https://go.dev/ref/spec#Packages) and to specify the default package name for import declarations.

​	[包语句](../Packages#package-clause-包子句)不是一个声明；包名不出现在任何作用域中。它的目的是识别属于同一[包](../Packages)的文件，并为导入声明指定默认的包名。

### Label scopes 标签作用域

Labels are declared by [labeled statements](https://go.dev/ref/spec#Labeled_statements) and are used in the ["break"](https://go.dev/ref/spec#Break_statements), ["continue"](https://go.dev/ref/spec#Continue_statements), and ["goto"](https://go.dev/ref/spec#Goto_statements) statements. It is illegal to define a label that is never used. In contrast to other identifiers, labels are not block scoped and do not conflict with identifiers that are not labels. The scope of a label is the body of the function in which it is declared and excludes the body of any nested function.

​	标签由[标签语句](../Statements#labeled-statements-标签语句)声明，用于 "[break](../Statements#break-statements---break-语句)"、"[continue](../Statements#continue-statements----continue-语句) "和 "[goto](../Statements#goto-statements-语句-goto) "语句中。定义一个从不使用的标签是非法的。与其他标识符不同的是，标签没有块作用域，也不会与不是标签的标识符冲突。标签的作用域是声明它的函数的主体，不包括任何嵌套函数的主体。

### Blank identifier 空白标识符

The *blank identifier* is represented by the underscore character `_`. It serves as an anonymous placeholder instead of a regular (non-blank) identifier and has special meaning in [declarations](https://go.dev/ref/spec#Declarations_and_scope), as an [operand](https://go.dev/ref/spec#Operands), and in [assignment statements](https://go.dev/ref/spec#Assignment_statements).

​	空白标识符由下划线字符`_`表示。它是一个`匿名的占位符`，而不是普通的（非空白）标识符，在声明中具有特殊的意义，可以作为[操作数](../Expressions#operands-操作数)，也可以在[赋值语句](../Statements#assignment-statements-赋值语句)中。

### Predeclared identifiers  预先声明的标识符

The following identifiers are implicitly declared in the [universe block](https://go.dev/ref/spec#Blocks) [[Go 1.18]({{< ref "/langSpec/Appendix#go-118">}})] [[Go 1.21]({{< ref "/langSpec/Appendix#go-121">}})]:

​	下列标识符是在[universe block](../Blocks)中隐式声明的 [[Go 1.18]({{< ref "/langSpec/Appendix#go-118">}})] [[Go 1.21]({{< ref "/langSpec/Appendix#go-121">}})]:

``` go
types:
	any bool byte comparable
	complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap clear close complex copy delete imag len
	make max min new panic print println real recover
```

### Exported identifiers 可导出的标识符

An identifier may be *exported* to permit access to it from another package. An identifier is exported if both: 

​	标识符可以被导出，以允许从另一个包访问它。一个标识符在以下两种情况下被导出：

1. the first character of the identifier's name is a Unicode uppercase letter (Unicode character category Lu); and
2. 标识符的第一个字符是一个Unicode大写字母（Unicode字符类别Lu）；并且
3. the identifier is declared in the [package block](https://go.dev/ref/spec#Blocks) or it is a [field name](https://go.dev/ref/spec#Struct_types) or [method name](https://go.dev/ref/spec#MethodName).
4. 该标识符在[package block](../Blocks)中被声明，或者它是一个字段名或方法名。

All other identifiers are not exported.

​	所有其他标识符都不被导出。

### Uniqueness of identifiers 标识符的唯一性

Given a set of identifiers, an identifier is called *unique* if it is *different* from every other in the set. Two identifiers are different if they are spelled differently, or if they appear in different [packages](https://go.dev/ref/spec#Packages) and are not [exported](https://go.dev/ref/spec#Exported_identifiers). Otherwise, they are the same.

​	给定一组标识符，如果一个标识符与这组标识符中的每一个都不同，则称为唯一。如果两个标识符的拼写不同，或者它们出现在不同的[包](../Packages)中并且不被[导出](#exported-identifiers-可导出的标识符)，那么它们就是不同的。否则，它们是相同的。

### Constant declarations 常量声明

A constant declaration binds a list of identifiers (the names of the constants) to the values of a list of [constant expressions](https://go.dev/ref/spec#Constant_expressions). The number of identifiers must be equal to the number of expressions, and the *n*th identifier on the left is bound to the value of the *n*th expression on the right.

​	常量声明将一列标识符（常量的名称）与一列[常量表达式](../Expressions#constant-expressions-常量表达式)的值绑定。标识符的数量必须等于表达式的数量，左边的第n个标识符被绑定到右边的第n个表达式的值上。

``` go
constDecl      = "const" ( ConstSpec | "(" { ConstSpec ";" } ")" ) .
ConstSpec      = IdentifierList [ [ Type ] "=" ExpressionList ] .

IdentifierList = identifier { "," identifier } .
ExpressionList = Expression { "," Expression } .
```

If the type is present, all constants take the type specified, and the expressions must be [assignable](https://go.dev/ref/spec#Assignability) to that type, which must not be a type parameter. If the type is omitted, the constants take the individual types of the corresponding expressions. If the expression values are untyped [constants](https://go.dev/ref/spec#Constants), the declared constants remain untyped and the constant identifiers denote the constant values. For instance, if the expression is a floating-point literal, the constant identifier denotes a floating-point constant, even if the literal's fractional part is zero. 

​	如果有指定类型，则所有常量都采用该指定的类型，并且表达式必须[可以赋值](../PropertiesOfTypesAndValues#assignability-可分配性)给该类型，该类型不能是一个类型参数。如果类型被省略，则常量采取相应表达式的类型。如果表达式的值是`无类型的（untyped）`[常量](../Constants)，那么声明的常量仍然是无类型，并且常量标识符表示该`无类型的`常量值。例如，如果表达式是一个浮点字面量，那么常量标识符就表示一个浮点常量，即使字面量的小数部分是0。

```go 
const Pi float64 = 3.14159265358979323846
const zero = 0.0         // untyped floating-point constant => 无类型浮点数数常量
const (
	size int64 = 1024
	eof        = -1  // untyped integer constant => 无类型整数常量
)
const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer and string constants => 无类型整数常量 和 字符串常量
const u, v float32 = 0, 3    // u = 0.0, v = 3.0
```

Within a parenthesized `const` declaration list the expression list may be omitted from any but the first ConstSpec. Such an empty list is equivalent to the textual substitution of the first preceding non-empty expression list and its type if any. Omitting the list of expressions is therefore equivalent to repeating the previous list. The number of identifiers must be equal to the number of expressions in the previous list. Together with the [`iota` constant generator](https://go.dev/ref/spec#Iota) this mechanism permits light-weight declaration of sequential values:

​	在括号内的`const`声明列表中，除了第一个ConstSpec外的任何表达式都可以省略表达式列表。这样的空列表相当于对前面第一个非空表达式列表及其类型（如果有的话）进行文本替换。因此，`省略表达式列表等同于重复前面的列表`。标识符的数量必须等于`上一个列表`中表达式的数量。与[iota 常量生成器](#iota)一起使用，此机制允许对连续值进行轻量级声明。

```go 
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays  // this constant is not exported
)
```

### iota

Within a [constant declaration](https://go.dev/ref/spec#Constant_declarations), the predeclared identifier `iota` represents successive untyped integer [constants](https://go.dev/ref/spec#Constants). Its value is the index of the respective [ConstSpec](https://go.dev/ref/spec#ConstSpec) in that constant declaration, starting at zero. It can be used to construct a set of related constants: 

​	在[常量声明](#constant-declarations-常量声明)中，预先声明的标识符`iota`表示连续的无类型整数[常量](../Constants)。它的值是该常量声明中各个 ConstSpec 的索引，从零开始。它可以被用来构造一组相关的常量：

```go 
const (
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)

const (
	a = 1 << iota  // a == 1  (iota == 0)
	b = 1 << iota  // b == 2  (iota == 1)
	c = 3          // c == 3  (iota == 2, unused)
	d = 1 << iota  // d == 8  (iota == 3)
)

const (
	u         = iota * 42  // u == 0     (untyped integer constant)
	v float64 = iota * 42  // v == 42.0  (float64 constant)
	w         = iota * 42  // w == 84    (untyped integer constant)
)

const x = iota  // x == 0
const y = iota  // y == 0
```

By definition, multiple uses of `iota` in the same ConstSpec all have the same value:

​	根据定义，在同一个ConstSpec中（可以认为是同一行中）多次使用`iota`都有相同的值：

```go 
const (
	bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0  (iota == 0)
	bit1, mask1                           // bit1 == 2, mask1 == 1  (iota == 1)
	_, _                                  //                        (iota == 2, unused)
	bit3, mask3                           // bit3 == 8, mask3 == 7  (iota == 3)
)
```

This last example exploits the [implicit repetition](https://go.dev/ref/spec#Constant_declarations) of the last non-empty expression list.

​	最后这个例子利用了最后一个非空表达式列表的[隐式重复](#constant-declarations-常量声明)。

### Type declarations 类型声明

A type declaration binds an identifier, the *type name*, to a [type](https://go.dev/ref/spec#Types). Type declarations come in two forms: alias declarations and type definitions.

​	类型声明将一个标识符，即类型名称，与一个类型绑定。类型声明有两种形式：`别名声明`和`类型定义`。

``` go
typeDecl = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
TypeSpec = AliasDecl | TypeDef .
```

#### Alias declarations 别名声明

An alias declaration binds an identifier to the given type [[Go 1.9](https://go.dev/ref/spec#Go_1.9)].

​	别名声明将一个标识符绑定到给定的类型上。 [[Go 1.9](https://go.dev/ref/spec#Go_1.9)]

```
AliasDecl = identifier "=" Type .
```

Within the [scope](https://go.dev/ref/spec#Declarations_and_scope) of the identifier, it serves as an *alias* for the type.

​	在标识符的[作用域](#declarations-and-scope-声明和作用域)内，它作为该类型的别名。

```go 
type (
	nodeList = []*Node  // nodeList and []*Node are identical types => nodeList 和[]*Node 类型是一致的
	Polar    = polar    // Polar and polar denote identical types => Polar 和 polar 表示一致的类型
)
```

> 个人注释
>
> ​	在函数调用时，可以传入类型别名的实参吗？=> 可以
>
> ```go
> package main
> 
> import "fmt"
> 
> type A = int
> 
> func IncOne(i int) int {
> 	return i + 1
> }
> 
> func main() {
> 	var a A = 1
> 	fmt.Println(IncOne(a)) // 2
> 
> 	fmt.Println(a + 2) // 3
> }
> 
> ```
>
> 

#### Type definitions 类型定义

A type definition creates a new, distinct type with the same [underlying type](https://go.dev/ref/spec#Underlying_types) and operations as the given type and binds an identifier, the *type name*, to it.

​	类型定义创建了一个新的、不同的类型，其[底层类型](../Types)和操作与给定的类型相同，并将标识符，即类型名称绑定到它。

``` go
typeDef = identifier [ TypeParameters ] Type .
```

The new type is called a *defined type*. It is [different](https://go.dev/ref/spec#Type_identity) from any other type, including the type it is created from.

​	这个新的类型被称为`已定义类型`。它与任何其他类型（包括它创建时使用的类型）不同。

```go 
type (
	Point struct{ x, y float64 }  // Point and struct{ x, y float64 } are different types
	polar Point                   // polar and Point denote different types
)

type TreeNode struct {
	left, right *TreeNode
	value any
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}
```

A defined type may have [methods](https://go.dev/ref/spec#Method_declarations) associated with it. It does not inherit any methods bound to the given type, but the [method set](https://go.dev/ref/spec#Method_sets) of an interface type or of elements of a composite type remains unchanged:

​	一个已定义类型可以有与之相关的方法。这一新定义的类型不会继承绑定到给定类型的任何方法，但是`接口类型`或`复合类型的元素`的[方法集](../PropertiesOfTypesAndValues#method-sets-方法集)保持不变。

```go 
// A Mutex is a data type with two methods, Lock and Unlock. 
//=> Mutex 是一个数据类型，其有 Lock 和 Unlock 两个方法
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock implementation */ }
func (m *Mutex) Unlock()  { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty. 
//=> NewMutex 和 Mutex 有相同的组成部分，但 NewMutex 的方法集是空的
type NewMutex Mutex

// The method set of PtrMutex's underlying type *Mutex remains unchanged, 
//=> PtrMutex 的底层类型 *Mutex 的方法集是不会改变的
// but the method set of PtrMutex is empty. 
//=> 但 PtrMutex 的方法集是空的
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its embedded field Mutex. 
//=> *PrintableMutex 的方法集包含 Lock 和 Unlock (这两个方法都是在 Mutex 这一嵌入字段中所绑定的)
type PrintableMutex struct {
	Mutex
}

// MyBlock is an interface type that has the same method set as Block. 
//=> MyBlock 是一个接口类型，其与Block接口有着相同的方法集
type MyBlock Block
```

Type definitions may be used to define different boolean, numeric, or string types and associate methods with them:

​	类型定义可用于定义不同的布尔型、数值型或字符串型，并为它们绑定方法：

```go 
type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

func (tz TimeZone) String() string {
	return fmt.Sprintf("GMT%+dh", tz)
}
```

If the type definition specifies [type parameters](https://go.dev/ref/spec#Type_parameter_declarations), the type name denotes a *generic type*. Generic types must be [instantiated](https://go.dev/ref/spec#Instantiations) when they are used.

​	如果类型定义指定了[类型参数](#type-parameter-declarations-类型参数声明)，那么这个`类型名称`表示一个`泛型`。`泛型`在使用时必须被[实例化](../Expressions#instantiations-实例化)。

```go  hl_lines="1 1"
type List[T any] struct {
	next  *List[T]
	value T
}
```

In a type definition the given type cannot be a type parameter.

​	在一个类型定义中，给定的类型不能是一个`类型参数`。

```go 
type T[P any] P    // illegal: P is a type parameter 
//=> 非法的：P 是一个类型参数

func f[T any]() {
	type L T   // illegal: T is a type parameter declared by the enclosing function 
    //=> 非法的：T 是封闭函数所声明的一个类型参数
}
```

A generic type may also have [methods](https://go.dev/ref/spec#Method_declarations) associated with it. In this case, the method receivers must declare the same number of type parameters as present in the generic type definition.

​	泛型也可以有与之相关的[方法](#method-declarations-方法声明)。在这种情况下，方法接收器必须声明与`泛型定义`中存在的相同数量的类型参数。

```go 
// The method Len returns the number of elements in the linked list l. 
//=> 方法 Len 会返回链接列表 l 中的元素的数量
func (l *List[T]) Len() int  { … }
```

### Type parameter declarations 类型参数声明

A type parameter list declares the *type parameters* of a generic function or type declaration. The type parameter list looks like an ordinary [function parameter list](https://go.dev/ref/spec#Function_types) except that the type parameter names must all be present and the list is enclosed in square brackets rather than parentheses [[Go 1.18]({{< ref "/langSpec/Appendix#go-118">}})].

​	类型参数列表`声明`了一个泛型函数或类型声明的类型参数。类型参数列表看起来和普通的函数参数列表一样，除了类型参数名称必须全部出现，并且列表被括在`方括号`中，而不是`花括号`中。

``` go
typeParameters  = "[" TypeParamList [ "," ] "]" .
TypeParamList   = TypeParamDecl { "," TypeParamDecl } .
TypeParamDecl   = IdentifierList TypeConstraint .
```

All non-blank names in the list must be unique. Each name declares a type parameter, which is a new and different [named type](https://go.dev/ref/spec#Types) that acts as a placeholder for an (as of yet) unknown type in the declaration. The type parameter is replaced with a *type argument* upon [instantiation](https://go.dev/ref/spec#Instantiations) of the generic function or type.

​	列表中所有非空白的名字必须是唯一的。每个名字都声明了一个类型参数，这是一个新的且不同的[命名类型](../Types)，作为声明中一个（到目前为止）未知类型的占位符。类型参数在泛型函数或类型[实例化](../Expressions#instantiations-实例化)时被替换为`类型实参（type argument）`。

```go
[P any]
[S interface{ ~[]byte|string }]
[S ~[]E, E any]
[P Constraint[int]]
[_ any]
```

Just as each ordinary function parameter has a parameter type, each type parameter has a corresponding (meta-)type which is called its [*type constraint*](https://go.dev/ref/spec#Type_constraints).

​	就像每个普通的函数参数都有一个参数类型一样，每个类型参数也有一个相应的被称为其[类型约束](#type-constraints-类型约束)的（元）类型。

A parsing ambiguity arises when the type parameter list for a generic type declares a single type parameter `P` with a constraint `C` such that the text `P C` forms a valid expression:

​	当泛型的类型参数列表声明了一个带有约束条件`C`的单一类型参数`P`，从而使文本`P C`构成一个有效的表达式时，就会出现解析歧义：

```go 
type T[P *C] …
type T[P (C)] …
type T[P *C|Q] …
…
```

In these rare cases, the type parameter list is indistinguishable from an expression and the type declaration is parsed as an array type declaration. To resolve the ambiguity, embed the constraint in an [interface](https://go.dev/ref/spec#Interface_types) or use a trailing comma: 

​	在这些罕见的情况下，类型参数列表很难与表达式进行区别，导致该类型声明被解析为一个数组类型声明。为了解决这种歧义，可将约束嵌入到一个[接口](../Types#interface-types-接口型)中或者在尾部使用逗号：

```go 
type T[P interface{*C}] …
type T[P *C,] …
```

Type parameters may also be declared by the receiver specification of a [method declaration](https://go.dev/ref/spec#Method_declarations) associated with a generic type.

​	类型参数也可以通过与泛型相关的[方法声明](#method-declarations-方法声明)的接收器规范来声明。

Within a type parameter list of a generic type `T`, a type constraint may not (directly, or indirectly through the type parameter list of another generic type) refer to `T`.

​	在泛型类型 的类型参数列表中，类型约束不得（直接或间接通过另一个泛型类型的类型参数列表）引用 。

```go
type T1[P T1[P]] …                    // illegal: T1 refers to itself
type T2[P interface{ T2[int] }] …     // illegal: T2 refers to itself
type T3[P interface{ m(T3[int])}] …   // illegal: T3 refers to itself
type T4[P T5[P]] …                    // illegal: T4 refers to T5 and
type T5[P T4[P]] …                    //          T5 refers to T4

type T6[P int] struct{ f *T6[P] }     // ok: reference to T6 is not in type parameter list
```

#### Type constraints 类型约束

A *type constraint* is an [interface](https://go.dev/ref/spec#Interface_types) that defines the set of permissible type arguments for the respective type parameter and controls the operations supported by values of that type parameter [[Go 1.18]({{< ref "/langSpec/Appendix#go-118">}})]. 

​	类型约束是一个[接口](../Types#interface-types-接口型)，该接口定义了对应的`类型参数`所允许的一组`类型实参`，并控制该类型参数的值所支持的操作。

``` go
typeConstraint = TypeElem .
```

If the constraint is an interface literal of the form `interface{E}` where `E` is an embedded [type element](https://go.dev/ref/spec#Interface_types) (not a method), in a type parameter list the enclosing `interface{ … }` may be omitted for convenience:

​	如果约束是一个形式为`interface{E}`的接口字面量，其中`E`是一个嵌入的类型元素（不是方法），在类型参数列表中，为了方便起见，可以省略参数列表中封闭的`interface{ … }`：

```go 
[T []P]                      // = [T interface{[]P}]
[T ~int]                     // = [T interface{~int}]
[T int|string]               // = [T interface{int|string}]
type Constraint ~int         // illegal: ~int is not inside a type parameter list 
//=> 非法的： ~int 不在一个类型参数列表中
```

The [predeclared](https://go.dev/ref/spec#Predeclared_identifiers) [interface type](https://go.dev/ref/spec#Interface_types) `comparable` denotes the set of all non-interface types that are [strictly comparable](https://go.dev/ref/spec#Comparison_operators) [[Go 1.18]({{< ref "/langSpec/Appendix#go-118">}})].

​	[预先声明的接口类型](../Types#interface-types-接口型)： `comparable` ，表示所有非接口类型的集合，这些类型是[严格可比较的](../Expressions#comparison-operators-比较运算符)  [[Go 1.18]({{< ref "/langSpec/Appendix#go-118">}})]。 具体来说，如果一个类型`T`实现了`comparable`：

- `T`不是一个接口类型并且`T`支持操作`==`和`!=`；或者
- `T`是一个接口类型，并且`T`的[类型集](../Types#interface-types-接口型)中的每个类型都实现了`comparable`。

Even though interfaces that are not type parameters are [comparable](https://go.dev/ref/spec#Comparison_operators), they are not strictly comparable and therefore they do not implement `comparable`. However, they [satisfy](https://go.dev/ref/spec#Satisfying_a_type_constraint) `comparable`.

​	尽管非类型参数的接口可以被[比较](../Expressions#comparison-operators-比较运算符)（可能导致运行时恐慌），但它们也没有实现`comparable`。

```go 
int                          // implements comparable => 实现了 comparable
[]byte                       // does not implement comparable (slices cannot be compared) 
//=> 未实现 comparable （切片不能被比较）

interface{}                  // does not implement comparable (see above) 
//=> 未实现 comparable (见上文)

interface{ ~int | ~string }  // type parameter only: implements comparable 
//=> 仅对类型形参：实现了 comparable

interface{ comparable }      // type parameter only: implements comparable 
//=> 仅对类型形参：实现了 comparable

interface{ ~int | ~[]byte }  // type parameter only: does not implement comparable (not all types in the type set are comparable) 
//=> 仅对类型形参：未实现 comparable （不是所有类型集中的类型都可以被比较）
```

The `comparable` interface and interfaces that (directly or indirectly) embed `comparable` may only be used as type constraints. They cannot be the types of values or variables, or components of other, non-interface types.

​	`comparable`这一接口和（直接或间接）嵌入`comparable`的接口只能作为`类型约束`使用。它们不能成为值或变量的类型，或其他非接口类型的组成部分。

### Variable declarations 变量声明

A variable declaration creates one or more [variables](https://go.dev/ref/spec#Variables), binds corresponding identifiers to them, and gives each a type and an initial value.

​	变量声明创建一个或多个变量，为它们绑定相应的标识符，并为每个变量设定一个类型和一个初始值。

```
VarDecl     = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
VarSpec     = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
```

```go 
var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map lookup; only interested in "found"
```

If a list of expressions is given, the variables are initialized with the expressions following the rules for [assignment statements](https://go.dev/ref/spec#Assignment_statements). Otherwise, each variable is initialized to its [zero value](https://go.dev/ref/spec#The_zero_value).

​	如果变量声明时给出的是表达式列表，则变量将按照[赋值语句](../statements#assignment-statements-赋值语句)的规则用表达式进行初始化。否则，每个变量被初始化为其[零值](../ProgramInitializationAndExecution#the-zero-value-零值)。

If a type is present, each variable is given that type. Otherwise, each variable is given the type of the corresponding initialization value in the assignment. If that value is an untyped constant, it is first implicitly [converted](https://go.dev/ref/spec#Conversions) to its [default type](https://go.dev/ref/spec#Constants); if it is an untyped boolean value, it is first implicitly converted to type `bool`. The predeclared value `nil` cannot be used to initialize a variable with no explicit type.

​	如果变量声明时提供了类型，则每个变量都被指定为该类型。否则，每个变量都被设定为赋值中相应的初始化值的类型。如果该值是一个`无类型的`常量，它首先被隐式[转换](../Expressions#conversions-转换)为其[默认类型](../Constants)；如果它是一个`无类型的`布尔值，它首先被隐式转换为`bool`。预先声明的值`nil`不能用来初始化一个没有明确类型的变量。

```go 
var d = math.Sin(0.5)  // d is float64
var i = 42             // i is int
var t, ok = x.(T)      // t is T, ok is bool
var n = nil            // illegal 非法的
```

Implementation restriction: A compiler may make it illegal to declare a variable inside a [function body](https://go.dev/ref/spec#Function_declarations) if the variable is never used.

​	实现限制：在[函数体](#function-declarations-函数声明)中声明一个变量，若该变量从未被使用，编译器可以认为它是非法的。

### Short variable declarations 短变量声明

A *short variable declaration* uses the syntax:

​	短变量声明使用的语法：

```
ShortVarDecl = IdentifierList ":=" ExpressionList .
```

It is shorthand for a regular [variable declaration](https://go.dev/ref/spec#Variable_declarations) with initializer expressions but no types:

​	它是带有初始化表达式但没有类型的常规[变量声明](#variable-declarations-变量声明)的简写：

```
"var" IdentifierList "=" ExpressionList .
```

```
i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w, _ := os.Pipe()  // os.Pipe() returns a connected pair of Files and an error, if any
_, y, _ := coord(p)   // coord() returns three values; only interested in y coordinate
```

Unlike regular variable declarations, a short variable declaration may *redeclare* variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type, and at least one of the non-[blank](https://go.dev/ref/spec#Blank_identifier) variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original. The non-blank variable names on the left side of `:=` must be [unique](https://go.dev/ref/spec#Uniqueness_of_identifiers).

​	与常规变量声明不同，短变量声明可以重新声明变量，前提是这些变量最初是在同一个块（如果该块是函数体，则是参数列表）中`以相同的类型声明的`，并且`至少有一个`非[空白](#blank-identifier-空白标识符)变量是新的。因此，重复声明只能出现在一个多变量的短声明中。重复声明并没有引入一个新的变量；它只是给原来的变量分配了一个新的值。`:=`左侧的非空白变量名必须是[唯一](#uniqueness-of-identifiers-标识符的唯一性)的。

```
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
x, y, x := 1, 2, 3                        // illegal: x repeated on left side of :=
```

Short variable declarations may appear only inside functions. In some contexts such as the initializers for ["if"](https://go.dev/ref/spec#If_statements), ["for"](https://go.dev/ref/spec#For_statements), or ["switch"](https://go.dev/ref/spec#Switch_statements) statements, they can be used to declare local temporary variables.

​	短变量声明只能出现在函数内部。在某些情况下，如 "[if](../Statements#if-statements---if-语句)"、"[for](../Statements#for-statements----for-语句) "或 "[switch](../Statements#switch-statements----switch-语句) "语句的初始化语句中，它们可以用来声明局部临时变量。

### Function declarations 函数声明

A function declaration binds an identifier, the *function name*, to a function.

​	函数声明将标识符（函数名称）绑定到函数。

``` go
functionDecl = "func" FunctionName [ TypeParameters ] Signature [ FunctionBody ] .
FunctionName = identifier .
FunctionBody = Block .
```

If the function's [signature](https://go.dev/ref/spec#Function_types) declares result parameters, the function body's statement list must end in a [terminating statement](https://go.dev/ref/spec#Terminating_statements).

​	如果函数的[签名](../Types#function-typess-函数型)声明了结果参数，那么函数体的语句列表必须以一个[终止语句](../Statements#terminating-statements-终止语句)结束。

```go 
func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
    // invalid: missing return statement 
    //=> 无效的: 缺少 return 语句
}
```

If the function declaration specifies [type parameters](https://go.dev/ref/spec#Type_parameter_declarations), the function name denotes a *generic function*. A generic function must be [instantiated](https://go.dev/ref/spec#Instantiations) before it can be called or used as a value.

​	如果函数声明中指定了[类型参数](#type-parameter-declarations-类型参数声明)，那么函数名就表示一个`泛型函数`。在被调用或作为值使用之前，泛型函数必须先被实例化。

```go 
func min[T ~int|~float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}
```

A function declaration without type parameters may omit the body. Such a declaration provides the signature for a function implemented outside Go, such as an assembly routine.

​	没有类型参数的函数声明可以省略函数体。这样的声明提供了一个在Go外部实现的函数的签名，比如一个汇编程序。

```go 
func flushICache(begin, end uintptr)  // implemented externally => 由外部实现
```

### Method declarations 方法声明

A method is a [function](https://go.dev/ref/spec#Function_declarations) with a *receiver*. A method declaration binds an identifier, the *method name*, to a method, and associates the method with the receiver's *base type*.

​	方法是带有接收器的[函数](#function-declarationss-函数声明)。方法声明将一个标识符，即方法名称，绑定到一个方法上，并将该方法与接收器的基本类型联系起来。

```
MethodDecl = "func" Receiver MethodName Signature [ FunctionBody ] .
Receiver   = Parameters .
```

The receiver is specified via an extra parameter section preceding the method name. That parameter section must declare a single non-variadic parameter, the receiver. Its type must be a [defined](https://go.dev/ref/spec#Type_definitions) type `T` or a pointer to a defined type `T`, possibly followed by a list of type parameter names `[P1, P2, …]` enclosed in square brackets. `T` is called the receiver *base type*. A receiver base type cannot be a pointer or interface type and it must be defined in the same package as the method. The method is said to be *bound* to its receiver base type and the method name is visible only within [selectors](https://go.dev/ref/spec#Selectors) for type `T` or `*T`.

​	接收器是通过方法名前面的一个额外的参数部分指定的。该参数部分必须声明一个非可变参数，即接收器。它的类型必须是一个[已定义](#type-declarations-类型定义)类型`T`或者一个指向已定义类型`T`的指针，后面可能是一个用方括号括起来的类型参数名称列表`[P1, P2, ...]`。`T`被称为接收器的`基本类型`。`接收器的基类型不能是一个指针或接口类型`，`并且它必须在与方法相同的包中定义`。这个定义过程称为将该方法与其接收器基本类型绑定，该方法名只在`T`或`*T`类型的[选择器](../Expressions#selectors-选择器)中可见。

A non-[blank](https://go.dev/ref/spec#Blank_identifier) receiver identifier must be [unique](https://go.dev/ref/spec#Uniqueness_of_identifiers) in the method signature. If the receiver's value is not referenced inside the body of the method, its identifier may be omitted in the declaration. The same applies in general to parameters of functions and methods.

​	一个非[空白](#blank-identifier-空白标识符)的接收器标识符在方法签名中必须是[唯一](#uniqueness-of-identifiers-标识符的唯一性)的。如果接收器的值没有在方法体中被引用，它的标识符可以在声明中被省略。这一规则与函数和方法的普通参数类似。

For a base type, the non-blank names of methods bound to it must be unique. If the base type is a [struct type](https://go.dev/ref/spec#Struct_types), the non-blank method and field names must be distinct.

​	对于一个基本类型，绑定到它的非空白名称必须是唯一的。如果基本类型是一个[结构类型](../Types#struct-types-结构体型)，非空白的方法名和字段名必须是唯一的。

Given defined type `Point` the declarations

​	给出定义类型`Point`，其声明：

```go 
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
```

bind the methods `Length` and `Scale`, with receiver type `*Point`, to the base type `Point`.

绑定了方法`Length`和`Scale`，接收器类型为`*Point`，对应基本类型`Point`。

If the receiver base type is a [generic type](https://go.dev/ref/spec#Type_declarations), the receiver specification must declare corresponding type parameters for the method to use. This makes the receiver type parameters available to the method. Syntactically, this type parameter declaration looks like an [instantiation](https://go.dev/ref/spec#Instantiations) of the receiver base type: the type arguments must be identifiers denoting the type parameters being declared, one for each type parameter of the receiver base type. The type parameter names do not need to match their corresponding parameter names in the receiver base type definition, and all non-blank parameter names must be unique in the receiver parameter section and the method signature. The receiver type parameter constraints are implied by the receiver base type definition: corresponding type parameters have corresponding constraints.

​	如果接收器的基本类型是一个[泛型](../DeclarationsAndScope#type-declarations-类型声明)，接收器规范必须为要使用的方法声明相应的类型形参。这使得接收器的类型形参对该方法可用。从语法上讲，这个类型形参声明看起来就像接收器基本类型的实例化：类型实参必须是表示被声明的类型参数的标识符，接收器基本类型的每个类型形参各有一个。`类型形参名无需匹配接收器基本类型定义中对应的形参名`，并且所有非空白形参名在接收器形参部分和方法签名中必须是唯一的。接收器类型形参的约束是由接收器基本类型定义所隐含的：相应的类型形参有相应的约束。=> 仍有疑问？？

```go 
type Pair[A, B any] struct {
	a A
	b B
}

func (p Pair[A, B]) Swap() Pair[B, A]  { … }  // receiver declares A, B //=> 接收器声明了 A，B
func (p Pair[First, _]) First() First  { … }  // receiver declares First, corresponds to A in Pair  //=> 接收器声明了 First， 对应 Pair 中的 A
```

