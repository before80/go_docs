+++
title = "声明和作用域"
date = 2023-05-17T09:59:21+08:00
weight = 10
description = ""
isCJKLanguage = true
draft = false
+++
## Declarations and scope 声明和作用域

> 原文：[https://go.dev/ref/spec#Declarations_and_scope](https://go.dev/ref/spec#Declarations_and_scope)

​	声明将一个非[空白](#blank-identifier-空白标识符)的标识符绑定到一个[常量](../Constants)、[类型](../Types)、[类型参数](#type-parameter-declarations-类型参数声明)、[变量](../Variables)、[函数](#function-declarations-函数声明)、[标签](../Statements#labeled-statements-标签语句)或[包](../Pachages#import-declarations-导入声明)。程序中的每个标识符都必须被声明。没有一个标识符可以在同一个块中声明两次，也没有一个标识符可以同时在文件块和包块中声明。

​	[空白标识符（即`_`）](#blank-identifier-空白标识符)可以像其他标识符一样在声明中使用，但它不引入绑定，因此不被声明。在包块中，标识符`init`只能用于[init函数](../ProgramInitializationAndExecution#package-initialization-包的初始化)的声明，和空白标识符一样，它不会引入一个新的绑定。

```
Declaration   = ConstDecl | TypeDecl | VarDecl .
TopLevelDecl  = Declaration | FunctionDecl | MethodDecl .
```

​	声明标识符的作用域是源文本的范围，其中标识符表示指定的常量、类型、变量、函数、标签或包。

​	Go在词法使用[块](../Blocks)来确定作用域：

1. [预先声明的标识符](../DeclarationsAndScope#predeclared-identifiers--预先声明的标识符)的作用域是`universe block`。
2. 表示常量、类型、变量或函数（但不是方法）的标识符在顶层（任何函数之外）声明的作用域是`package block`。
3. 导入的包的包名的作用域是在导入声明的文件的`file block`。
4. 表示方法接收器、函数参数或结果变量的标识符的作用域是函数体。
5. 表示函数的类型参数或由方法接收器声明的标识符的作用域从函数名称之后开始，在函数体的末端结束。
6. 表示类型的类型参数的标识符的作用域从类型的名称之后开始，在TypeSpec的末尾结束。=>仍有疑问？？
7. 在函数中声明的常量或变量标识符的作用域从ConstSpec或VarSpec（ShortVarDecl用于短变量声明）的末尾开始，在最里面的包含块的末尾结束。
8. 在函数中声明的类型标识符的作用域从TypeSpec中的标识符开始，在最内层的包含块的末端结束。

​	在块中声明的标识符可以在内部块中被重新声明。当内部声明的标识符在作用域内时，它表示内部声明所声明的实体。

​	[包语句](../Packages#package-clause-包子句)不是一个声明；包名不出现在任何作用域中。它的目的是识别属于同一[包](../Packages)的文件，并为导入声明指定默认的包名。

### Label scopes 标签作用域

​	标签由[标签语句](../Statements#labeled-statements-标签语句)声明，用于 "[break](../Statements#break-statements---break-语句)"、"[continue](../Statements#continue-statements----continue-语句) "和 "[goto](../Statements#goto-statements-语句-goto) "语句中。定义一个从不使用的标签是非法的。与其他标识符不同的是，标签没有块作用域，也不会与不是标签的标识符冲突。标签的作用域是声明它的函数的主体，不包括任何嵌套函数的主体。

### Blank identifier 空白标识符

​	空白标识符由下划线字符`_`表示。它是一个`匿名的占位符`，而不是普通的（非空白）标识符，在声明中具有特殊的意义，可以作为[操作数](../Expresssions#operands-操作数)，也可以在[赋值语句](../Statements#assignment-statements-赋值语句)中。

### Predeclared identifiers  预先声明的标识符

下列标识符是在[universe block](../Blocks)中隐式声明的：

```
Types:
	any bool byte comparable
	complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap close complex copy delete imag len
	make new panic print println real recover
```

### Exported identifiers 可导出的标识符

​	标识符可以被导出，以允许从另一个包访问它。一个标识符在以下两种情况下被导出：

1. 标识符的第一个字符是一个Unicode大写字母（Unicode字符类别Lu）；并且
2. 该标识符在[package block](../Blocks)中被声明，或者它是一个字段名或方法名。

所有其他标识符都不被导出。

### Uniqueness of identifiers 标识符的唯一性

​	给定一组标识符，如果一个标识符与这组标识符中的每一个都不同，则称为唯一。如果两个标识符的拼写不同，或者它们出现在不同的[包](../Packages)中并且不被[导出](#exported-identifiers-可导出的标识符)，那么它们就是不同的。否则，它们是相同的。

### Constant declarations 常量声明

​	常量声明将一列标识符（常量的名称）与一列[常量表达式](../Expressions#constant-expressions-常量表达式)的值绑定。标识符的数量必须等于表达式的数量，左边的第n个标识符被绑定到右边的第n个表达式的值上。

```
ConstDecl      = "const" ( ConstSpec | "(" { ConstSpec ";" } ")" ) .
ConstSpec      = IdentifierList [ [ Type ] "=" ExpressionList ] .

IdentifierList = identifier { "," identifier } .
ExpressionList = Expression { "," Expression } .
```

​	如果有指定类型，则所有常量都采用该指定的类型，并且表达式必须[可以赋值](../PropertiesOfTypesAndValues#assignability-可分配性)给该类型，该类型不能是一个类型参数。

​	如果类型被省略，则常量采取相应表达式的类型。

​	如果表达式的值是`无类型的（untyped）`[常量](../Constants)，那么声明的常量仍然是无类型，并且常量标识符表示该`无类型的`常量值。例如，如果表达式是一个浮点字面量，那么常量标识符就表示一个浮点常量，即使字面量的小数部分是0。

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

### Iota

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

​	根据定义，在同一个ConstSpec中（可以认为是同一行中）多次使用`iota`都有相同的值：

```go 
const (
	bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0  (iota == 0)
	bit1, mask1                           // bit1 == 2, mask1 == 1  (iota == 1)
	_, _                                  //                        (iota == 2, unused)
	bit3, mask3                           // bit3 == 8, mask3 == 7  (iota == 3)
)
```

最后这个例子利用了最后一个非空表达式列表的[隐式重复](#constant-declarations-常量声明)。

### Type declarations 类型声明

​	类型声明将一个标识符，即类型名称，与一个类型绑定。类型声明有两种形式：`别名声明`和`类型定义`。

```
TypeDecl = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
TypeSpec = AliasDecl | TypeDef .
```

#### Alias declarations 别名声明

别名声明将一个标识符绑定到给定的类型上。

```
AliasDecl = identifier "=" Type .
```

在标识符的[作用域](#declarations-and-scope-声明和作用域)内，它作为该类型的别名。

```go 
type (
	nodeList = []*Node  // nodeList and []*Node are identical types => nodeList 和[]*Node 类型是一致的
	Polar    = polar    // Polar and polar denote identical types => Polar 和 polar 表示一致的类型
)
```

#### Type definitions 类型定义

​	类型定义创建了一个新的、不同的类型，其[底层类型](../Types)和操作与给定的类型相同，并将标识符，即类型名称绑定到它。

```
TypeDef = identifier [ TypeParameters ] Type .
```

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

类型定义可用于定义不同的布尔型、数值型或字符串型，并为它们绑定方法：

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

​	如果类型定义指定了[类型参数](#type-parameter-declarations-类型参数声明)，那么这个`类型名称`表示一个`泛型`。`泛型`在使用时必须被[实例化](../Expressions#instantiations-实例化)。

```go  hl_lines="1 1"
type List[T any] struct {
	next  *List[T]
	value T
}
```

In a type definition the given type cannot be a type parameter.

在一个类型定义中，给定的类型不能是一个`类型参数`。

```go 
type T[P any] P    // illegal: P is a type parameter 
//=> 非法的：P 是一个类型参数

func f[T any]() {
	type L T   // illegal: T is a type parameter declared by the enclosing function 
    //=> 非法的：T 是封闭函数所声明的一个类型参数
}
```

​	泛型也可以有与之相关的[方法](#method-declarations-方法声明)。在这种情况下，方法接收器必须声明与`泛型定义`中存在的相同数量的类型参数。=>仍有疑问？？

```go 
// The method Len returns the number of elements in the linked list l. 
//=> 方法 Len 会返回链接列表 l 中的元素的数量
func (l *List[T]) Len() int  { … }
```

### Type parameter declarations 类型参数声明

​	类型参数列表`声明`了一个泛型函数或类型声明的类型参数。类型参数列表看起来和普通的函数参数列表一样，除了类型参数名称必须全部出现，并且列表被括在`方括号`中，而不是`花括号`中。

```
TypeParameters  = "[" TypeParamList [ "," ] "]" .
TypeParamList   = TypeParamDecl { "," TypeParamDecl } .
TypeParamDecl   = IdentifierList TypeConstraint .
```

​	列表中所有非空白的名字必须是唯一的。每个名字都声明了一个类型参数，这是一个新的且不同的[命名类型](../Types)，作为声明中一个（到目前为止）未知类型的占位符。类型参数在泛型函数或类型[实例化](../Expressions#instantiations-实例化)时被替换为`类型实参（type argument）`。

```
[P any]
[S interface{ ~[]byte|string }]
[S ~[]E, E any]
[P Constraint[int]]
[_ any]
```

​	就像每个普通的函数参数都有一个参数类型一样，每个类型参数也有一个相应的被称为其[类型约束](#type-constraints-类型约束)的（元）类型。

​	当泛型的类型参数列表声明了一个带有约束条件`C`的单一类型参数`P`，从而使文本`P C`构成一个有效的表达式时，就会出现解析歧义：

```go 
type T[P *C] …
type T[P (C)] …
type T[P *C|Q] …
…
```

​	在这些罕见的情况下，类型参数列表很难与表达式进行区别，导致该类型声明被解析为一个数组类型声明。为了解决这种歧义，可将约束嵌入到一个[接口](../Types#interface-types-接口型)中或者在尾部使用逗号：

```go 
type T[P interface{*C}] …
type T[P *C,] …
```

​	类型参数也可以通过与泛型相关的[方法声明](#method-declarations-方法声明)的接收器规范来声明。

#### Type constraints 类型约束

​	类型约束是一个[接口](../Types#interface-types-接口型)，该接口定义了对应的`类型参数`所允许的一组`类型实参`，并控制该类型参数的值所支持的操作。

```
TypeConstraint = TypeElem .
```

​	如果约束是一个形式为`interface{E}`的接口字面量，其中`E`是一个嵌入的类型元素（不是方法），在类型参数列表中，为了方便起见，可以省略参数列表中封闭的`interface{ … }`：

```go 
[T []P]                      // = [T interface{[]P}]
[T ~int]                     // = [T interface{~int}]
[T int|string]               // = [T interface{int|string}]
type Constraint ~int         // illegal: ~int is not inside a type parameter list 
//=> 非法的： ~int 不在一个类型参数列表中
```

​	[预先声明的接口类型](../Types#interface-types-接口型)： `comparable` ，表示所有非接口类型的集合，这些类型是[可比较的](../Expressions#comparison-operators-比较运算符)。具体来说，如果一个类型`T`实现了`comparable`：

- `T`不是一个接口类型并且`T`支持操作`==`和`!=`；或者
- `T`是一个接口类型，并且`T`的[类型集](../Types#interface-types-接口型)中的每个类型都实现了`comparable`。

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

​	`comparable`这一接口和（直接或间接）嵌入`comparable`的接口只能作为`类型约束`使用。它们不能成为值或变量的类型，或其他非接口类型的组成部分。

### Variable declarations 变量声明

​	变量声明创建一个或多个变量，为它们绑定相应的标识符，并为每个变量设定一个类型和一个初始值。

```go 
VarDecl     = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
VarSpec     = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
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

​	如果变量声明时给出的是表达式列表，则变量将按照[赋值语句](../statements#assignment-statements-赋值语句)的规则用表达式进行初始化。否则，每个变量被初始化为其[零值](../ProgramInitializationAndExecution#the-zero-value-零值)。

​	如果变量声明时提供了类型，则每个变量都被指定为该类型。否则，每个变量都被设定为赋值中相应的初始化值的类型。如果该值是一个`无类型的`常量，它首先被隐式[转换](../Expressions#conversions-转换)为其[默认类型](../Constants)；如果它是一个`无类型的`布尔值，它首先被隐式转换为`bool`。预先声明的值`nil`不能用来初始化一个没有明确类型的变量。

```go 
var d = math.Sin(0.5)  // d is float64
var i = 42             // i is int
var t, ok = x.(T)      // t is T, ok is bool
var n = nil            // illegal 非法的
```

实现限制：在[函数体](#function-declarations-函数声明)中声明一个变量，若该变量从未被使用，编译器可以认为它是非法的。

### Short variable declarations 短变量声明

短变量声明使用的语法：

```
ShortVarDecl = IdentifierList ":=" ExpressionList .
```

它是带有初始化表达式但没有类型的常规[变量声明](#variable-declarations-变量声明)的简写：

```
"var" IdentifierList "=" ExpressionList .
i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w, _ := os.Pipe()  // os.Pipe() returns a connected pair of Files and an error, if any
_, y, _ := coord(p)   // coord() returns three values; only interested in y coordinate
```

​	与常规变量声明不同，短变量声明可以重新声明变量，前提是这些变量最初是在同一个块（如果该块是函数体，则是参数列表）中`以相同的类型声明的`，并且`至少有一个`非[空白](#blank-identifier-空白标识符)变量是新的。因此，重复声明只能出现在一个多变量的短声明中。重复声明并没有引入一个新的变量；它只是给原来的变量分配了一个新的值。`:=`左侧的非空白变量名必须是[唯一](#uniqueness-of-identifiers-标识符的唯一性)的。

```
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
x, y, x := 1, 2, 3                        // illegal: x repeated on left side of :=
```

​	短变量声明只能出现在函数内部。在某些情况下，如 "[if](../Statements#if-statements---if-语句)"、"[for](../Statements#for-statements----for-语句) "或 "[switch](../Statements#switch-statements----switch-语句) "语句的初始化语句中，它们可以用来声明局部临时变量。

### Function declarations 函数声明

函数声明将标识符（函数名称）绑定到函数。

```
FunctionDecl = "func" FunctionName [ TypeParameters ] Signature [ FunctionBody ] .
FunctionName = identifier .
FunctionBody = Block .
```

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

​	如果函数声明中指定了[类型参数](#type-parameter-declarations-类型参数声明)，那么函数名就表示一个`泛型函数`。在被调用或作为值使用之前，泛型函数必须先被实例化。

```go 
func min[T ~int|~float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}
```

​	没有类型参数的函数声明可以省略函数体。这样的声明提供了一个在Go外部实现的函数的签名，比如一个汇编程序。

```go 
func flushICache(begin, end uintptr)  // implemented externally => 由外部实现
```

### Method declarations 方法声明

​	方法是带有接收器的[函数](#function-declarationss-函数声明)。方法声明将一个标识符，即方法名称，绑定到一个方法上，并将该方法与接收器的基本类型联系起来。

```
MethodDecl = "func" Receiver MethodName Signature [ FunctionBody ] .
Receiver   = Parameters .
```

​	接收器是通过方法名前面的一个额外的参数部分指定的。该参数部分必须声明一个非可变参数，即接收器。它的类型必须是一个[已定义](#type-declarations-类型定义)类型`T`或者一个指向已定义类型`T`的指针，后面可能是一个用方括号括起来的类型参数名称列表`[P1, P2, ...]`。`T`被称为接收器的`基本类型`。`接收器的基类型不能是一个指针或接口类型`，`并且它必须在与方法相同的包中定义`。这个定义过程称为将该方法与其接收器基本类型绑定，该方法名只在`T`或`*T`类型的[选择器](../Expressions#selectors-选择器)中可见。

​	一个非[空白](#blank-identifier-空白标识符)的接收器标识符在方法签名中必须是[唯一](#uniqueness-of-identifiers-标识符的唯一性)的。如果接收器的值没有在方法体中被引用，它的标识符可以在声明中被省略。这一规则与函数和方法的普通参数类似。

​	对于一个基本类型，绑定到它的非空白名称必须是唯一的。如果基本类型是一个[结构类型](../Types#struct-types-结构体型)，非空白的方法名和字段名必须是唯一的。

给出定义类型`Point`，其声明：

```go 
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
```

绑定了方法`Length`和`Scale`，接收器类型为`*Point`，对应基本类型`Point`。

If the receiver base type is a [generic type](https://go.dev/ref/spec#Type_declarations), the receiver specification must declare corresponding type parameters for the method to use. This makes the receiver type parameters available to the method. Syntactically, this type parameter declaration looks like an [instantiation](https://go.dev/ref/spec#Instantiations) of the receiver base type: the type arguments must be identifiers denoting the type parameters being declared, one for each type parameter of the receiver base type. The type parameter names do not need to match their corresponding parameter names in the receiver base type definition, and all non-blank parameter names must be unique in the receiver parameter section and the method signature. The receiver type parameter constraints are implied by the receiver base type definition: corresponding type parameters have corresponding constraints.

​	如果接收器的基本类型是一个[泛型](../DeclarationsAndScope#type-declarations-类型声明)，接收器规范必须为要使用的方法声明相应的类型形参。这使得接收器的类型形参对该方法可用。从语法上讲，这个类型形参声明看起来就像接收器基本类型的实例化：类型实参必须是表示被声明的类型参数的标识符，接收器基本类型的每个类型形参各有一个。`类型形参名无需匹配接收器基本类型定义中对应的形参名`，并且所有非空白形参名在接收器形参部分和方法签名中必须是唯一的。接收器类型形参的约束是由接收器基本类型定义所隐含的：相应的类型形参有相应的约束。=> 仍有疑问？？

```go 
type Pair[A, B any] struct {
	a A
	b B
}

func (p Pair[A, B]) Swap() Pair[B, A]  { … }  // receiver declares A, B 
//=> 接收器声明了 A，B
func (p Pair[First, _]) First() First  { … }  // receiver declares First, corresponds to A in Pair 
//=> 接收器声明了 First， 对应 Pair 中的 A
```

