+++
title = "类型"
date = 2023-05-17T09:59:21+08:00
weight = 7
description = ""
isCJKLanguage = true
draft = false
+++
## Types 类型

> 原文：[https://go.dev/ref/spec#Types ](https://go.dev/ref/spec#Types )

​	类型决定了一组值以及特定于这些值的操作和方法。如果类型具有类型名称，则可以用类型名称表示。如果类型是泛型，则后面必须跟[类型参数](../Expressions#instantiations-实例化)。还可以使用`类型字面量`指定类型，它由现有类型组成一个类型。

```
Type      = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
TypeName  = identifier | QualifiedIdent .
TypeArgs  = "[" TypeList [ "," ] "]" .
TypeList  = Type { "," Type } .
TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
            SliceType | MapType | ChannelType .
```

​	该语言[预先声明](../DeclarationsAndScope#predeclared-identifiers--预先声明的标识符)了某些类型的名称。其他类型是通过[类型声明](../DeclarationsAndScope#type-declarations-类型声明)或[类型参数列表](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)引入的。`复合类型`：数组、结构体、指针、函数、接口、切片、映射和通道类型 —— 可以用类型字面量来构造。

​	预先声明的类型、[已定义的类型](../DeclarationsAndScope#type-declarations-类型声明)和类型参数被称为`命名类型`。如果别名声明中给出的类型是命名类型，则别名也表示一个（新的）命名类型。

### Boolean types 布尔型

​	布尔型表示由预先声明的常量`true`和`false`表示的一组布尔真值。预先声明的布尔类型是`bool`；它是一个[已定义的类型]({{<ref "/docs/References/LanguageSpecification/DeclarationsAndScope#type-definitions-类型定义">}})。

### Numeric types 数值型

​	整数类型、浮点类型或复数类型分别表示整数、浮点或复数的值的集合。它们被统称为`数值类型`。预先声明的与体系结构无关的数值类型有：

```go 
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

​	The value of an *n*-bit integer is *n* bits wide and represented using [two's complement arithmetic](https://en.wikipedia.org/wiki/Two's_complement). =>仍有疑问？？

​	一个n位整数的值是n位宽，并用[二进制补码运算法（two's complement arithmetic）](https://en.wikipedia.org/wiki/Two's_complement)表示。

​	还有一组预先声明的整数类型，它们具有特定于实现的大小：

```go 
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

​	为了避免可移植性问题，所有的数值类型都是[已定义的类型](../DeclarationsAndScope#type-definitions-类型定义)，因此除了 `byte` (`uint8`的别名)和 `rune` (`int32`的别名)之外，它们是截然不同的。 当不同的数值类型在表达式或赋值中混合使用时，需要进行显式转换。例如，int32和int不是相同类型，尽管它们在一个特定的体系结构上可能具有相同的大小。

### String types 字符串型

​	字符串类型表示字符串值的集合。字符串值是字节序列（可能为空）。`字节数`被称为字符串的`长度`，并且永远不会是负数。字符串是不可变的：一旦创建，就不可能改变字符串的内容。预先声明的字符串类型是`string`；它是一种[已定义的类型](../DeclarationsAndScope#type-declarations-类型定义)。

​	可以使用内置函数 `len` 查找字符串 `s` 的长度。如果字符串是常量，那么长度就是编译时常量。字符串的字节可以通过整数[索引](../Expressions#index-expressions-索引表达式)0到`len(s)-1`来访问。`取这样一个元素的地址是非法的`；如果`s[i]`是字符串的第`i`个字节，那么`&s[i]`是无效的。

### Array types 数组型

​	数组是单类型的元素组成的编号序列，称为元素类型。`元素的数量`被称为数组的`长度`，并且永远不会是负数。

```
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```

​	长度是数组类型的一部分；它必须求值为一个非负[常数](../Constants)，该常数可由 int 类型的值[表示](../PropertiesOfTypesAndValues#representability)。数组`a`的长度可以用内置函数`len`发现。元素可以通过整数[索引](../Expressions#index-expressions)0到`len(a)-1`进行寻址。数组类型总是一维的，但是可以组成多维类型。

```go 
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

### Slice types 切片型

​	切片是底层数组的连续段的描述符，并提供对该数组中编号的元素序列的访问。切片类型表示其元素类型的所有数组切片的集合。`元素的数量`被称为切片的`长度`，并且永远不会是负数。一个未初始化的切片的值是`nil`。

```
SliceType = "[" "]" ElementType .
```

​	切片`s`的长度可以通过内置函数`len`发现；与数组不同，它在运行过程中可能会发生变化。元素可以通过整数[索引](../Expressions#index-expressions)0到`len(s)-1`进行寻址。给定元素的切片索引可能小于底层数组中同一元素的索引。

​	切片一旦被初始化，总是与保存其元素的底层数组相关联。因此，一个切片与它的底层数组和同一数组的其他切片共享存储；相反，不同的数组总是表示不同的存储。

​	切片的底层数组可以超过切片的末端。容量是对这一范围的衡量：它是切片的长度和切片之外的数组长度之和；可以通过从原始切片切割一个新的切片来创建一个达到这个容量的切片。使用内置函数 `cap(a)`可以发现切片 `a` 的容量。

​	可以使用内置函数`make`来创建一个给定元素类型`T`的新的、初始化的切片值，该函数接受一个切片类型和指定长度和可选容量的参数。用`make`创建的切片总是分配一个新的、隐藏的数组，返回的切片值指向该数组。也就是说，执行

```go 
make([]T, length, capacity)
```

产生的切片与分配一个数组并对其进行[切片](../Expressions#slice-expressions)是一样的，所以这两个表达式是等同的：

```go 
make([]int, 50, 100)
new([100]int)[0:50]
```

​	和数组一样，切片总是一维的，但可以通过组合来构造更高维的对象。对于数组的数组，内部数组在结构上总是相同的长度；但是对于切片的切片（或切片的数组），内部长度可以动态变化。此外，`内部切片必须被单独初始化`。

### Struct types 结构体型

​		结构体是一个命名元素（称为`字段`）的序列，，每个字段都有一个名称和一个类型。字段名可以显示地指定（IdentifierList）或隐含地指定（EmbeddedField）。在一个结构体中，非[空白](../DeclarationsAndScope#blank-identifier)字段名必须是[唯一](../DeclarationsAndScope#uniqueness-of-identifiers)的。

```
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName [ TypeArgs ] .
Tag           = string_lit .
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

​	一个声明了类型但没有明确字段名的字段被称为`嵌入式字段`。嵌入字段必须被指定为一个类型名`T`或一个指向非接口类型名`*T`的指针，而且`T`本身不能是一个指针类型。未限定类型名作为字段名。

```go 
// A struct with four embedded fields of types T1, *T2, P.T3 and *P.T4
struct {
	T1        // field name is T1
	*T2       // field name is T2
	P.T3      // field name is T3
	*P.T4     // field name is T4
	x, y int  // field names are x and y
}
```

下面的声明是非法的，`因为字段名在一个结构体类型中必须是唯一的`。

```go 
struct {
	T     // conflicts with embedded field *T and *P.T
	*T    // conflicts with embedded field T and *P.T
	*P.T  // conflicts with embedded field T and *T
}
```

​	如果`x.f`是表示字段或[方法](../DeclarationsAndScope#function-declarations)`f`的合法[选择器](../Expressions#selectors )，那么结构体`x`中的嵌入式字段或方法`f`被称为（自动）提升（的字段或方法）。

​	被提升的字段与结构体中的普通字段一样，只是它们不能在结构体的[复合字面量](../Expressions#composite-literals)中作为字段名使用。

​	给定一个结构体类型`S`和一个[命名类型](../Types)`T`，提升的方法被包含在结构体的方法集中，如下所示：

- 如果`S`包含一个嵌入式字段`T`，那么`S`和`*S`的方法集都包括带有接收器`T`的提升方法，`*S`的方法集也包括带有接收器`*T`的提升方法。
- 如果`S`包含一个嵌入式字段`*T`，那么`S`和`*S`的方法集都包括带有接收器`T`或`*T`的提升方法。

​	一个字段声明后面可以有一个可选的`字符串字面量标签`，它成为相应字段声明中所有字段的属性。一个空的标签字符串等同于一个不存在标签。标签通过[反射接口](https://pkg.go.dev/reflect#StructTag)可见，并参与结构体的[类型标识](../PropertiesOfTypesAndValues#type-identity)，但在其他情况下被忽略。

```go 
struct {
	x, y float64 ""  // an empty tag string is like an absent tag
	name string  "any string is permitted as a tag"
	_    [4]byte "ceci n'est pas un champ de structure"
}

// A struct corresponding to a TimeStamp protocol buffer. 
// 一个对应于 TimeStamp 协议缓冲区的结构。
// The tag strings define the protocol buffer field numbers;
// 标签字符串定义协议缓冲区字段编号
// they follow the convention outlined by the reflect package.
// 它们遵循反射包描述的约定。
struct {
	microsec  uint64 `protobuf:"1"`
	serverIP6 uint64 `protobuf:"2"`
}
```

### Pointer types 指针型

​	指针类型表示指向给定类型（称为指针的基本类型）[变量](../Variables)的所有指针的集合。一个未初始化的指针的值是`nil`。

```go 
PointerType = "*" BaseType .
BaseType    = Type .
*Point
*[4]int
```

### Function types 函数型

​	函数类型表示具有相同参数类型和结果类型的所有函数的集合。一个函数类型的未初始化变量的值是`nil`。

```
FunctionType   = "func" Signature .
Signature      = Parameters [ Result ] .
Result         = Parameters | Type .
Parameters     = "(" [ ParameterList [ "," ] ] ")" .
ParameterList  = ParameterDecl { "," ParameterDecl } .
ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
```

​	在参数或结果的列表中，名称（IdentifierList）必须全部存在或全部不存在。如果存在（名称），每个名称代表指定类型的一个项（参数或结果），并且签名中所有非[空白](../PropertiesOfTypesAndValues#blank-identity)的名称必须是唯一的。如果不存在（名称），每个类型代表该类型的一个项。参数和结果列表总是用括号表示，但如果正好仅有一个未命名的结果，则可以写成未括号的类型。

​	在函数签名中的最后一个传入参数可以有一个`类型前缀` `...`。有这样一个参数的函数被称为 variadic （`可变参数函数`），可以用零个或多个参数来调用该函数。

```go 
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

### Interface types 接口型

​	接口类型定义了一个类型集。接口类型的变量可以存储该接口类型集中的任何类型的值。这样的类型被称为[实现了该接口](#implementing-an-interface)。未初始化的接口类型变量的值是`nil`。

```
InterfaceType  = "interface" "{" { InterfaceElem ";" } "}" .
InterfaceElem  = MethodElem | TypeElem .
MethodElem     = MethodName Signature .
MethodName     = identifier .
TypeElem       = TypeTerm { "|" TypeTerm } .
TypeTerm       = Type | UnderlyingType .
UnderlyingType = "~" Type .
```

​	接口类型由接口元素列表指定。接口元素是一个方法或一个类型元素，其中类型元素是一个或多个类型术语的联合。类型术语可以是一个单一类型，也可以是一个单一的底层类型。

#### Basic interfaces 基本接口

​	在其最基本的形式中，接口指定了一个（可能是空的）方法列表。由这样一个接口定义的类型集是实现所有这些方法的类型集，而相应的方法集则完全由这个接口指定的方法组成。那些类型集可以`完全由一个方法列表`来定义的接口被称为`基本接口`。

```go 
// A simple File interface.
interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}
```

每个显式指定的方法的名称必须是[唯一](../DeclarationsAndScope#uniqueness-of-identifiers)的，不能是[空白](../PropertiesOfTypesAndValues#blank-identity)。

```go 
interface {
	String() string
	String() string  // illegal: String not unique
	_(x int)         // illegal: method must have non-blank name
}
```

​	多个类型可以实现一个（相同的）接口。例如，如果两个类型`S1`和`S2`的方法设置为

```go 
func (p T) Read(p []byte) (n int, err error)
func (p T) Write(p []byte) (n int, err error)
func (p T) Close() error
```

(其中`T`代表`S1`或`S2`），那么`File`接口就由`S1`和`S2`实现，而不管`S1`和`S2`可能有其他方法或共享什么其他方法。

​	作为接口类型集成员的每个类型都实现了该接口。任何给定的类型都可以实现几个不同的接口。例如，所有类型都实现`空接口` （interface {}），它代表所有（非接口）类型的集合：

```go 
interface{}
```

为了方便，预先声明的类型`any`是`空接口的别名`。

​	类似地，考虑这个接口规范，它出现在定义名为 `Locker` 的接口的类型声明中：

```go 
type Locker interface {
	Lock()
	Unlock()
}
```

如果`S1`和`S2`也实现了

```go 
func (p T) Lock() { … }
func (p T) Unlock() { … }
```

他们就实现了`Locker`接口和`File`接口。

#### Embedded interfaces 嵌入式接口

​	接口`T`可以使用（可能是限定的）接口类型名称`E`作为接口元素。这就是在 `T` 中嵌入接口 `E`。`T`的类型集是由`T`的显式声明方法定义的类型集和`T`的嵌入接口的类型集的`交集`。换句话说，`T`的类型集是实现`T`的所有显式声明的方法以及`E`的所有方法的所有类型的集合。

```go 
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter's methods are Read, Write, and Close.
type ReadWriter interface {
	Reader  // includes methods of Reader in ReadWriter's method set
	Writer  // includes methods of Writer in ReadWriter's method set
}
```

在嵌入接口时，具有[相同](../DeclarationsAndScope#uniqueness-of-identifiers)名称的方法必须具有[相同](../PropertiesOfTypesAndValues#type-identity)的签名。

```go 
type ReadCloser interface {
	Reader   // includes methods of Reader in ReadCloser's method set
	Close()  // illegal: signatures of Reader.Close and Close are different
}
```

#### General interfaces 通用接口

​	在最通用的形式下，接口元素也可以是一个任意类型的术语`T`，或者是一个指定底层类型`T`的`~T`形式的术语，或者是术语`t1|t2|...|tn`的联合。与方法规范一起，这些元素能够精确地定义一个接口的类型集，如下所示：

- 空接口的类型集是`所有非接口类型的集合`。
- 非空接口的类型集是其接口元素的类型集的交集。
- 方法规范的类型集是其方法集包括该方法的`所有非接口类型的集合`。
- 非接口类型术语的类型集是仅由该类型组成的集合。
- 形式为`~T`的术语的类型集是其底层类型为`T`的所有类型的集合。
- 术语`t1|t2|...|tn`的联合体类型集是各术语类型集的联合。

​	量化 "`所有非接口类型的集合` "不仅指手头程序中声明的所有（非接口）类型，还指所有可能程序中的所有可能类型，因此是无限的。类似地，给定实现某个特定方法的`所有非接口类型的集合`，这些类型的方法集的`交集`将正好包含该方法，即使手头的程序中的所有类型总是将该方法与另一个方法配对。

通过构造，`一个接口的类型集永远不会包含一个接口类型`。

```go 
// An interface representing only the type int.
// 仅表示 int 类型的接口
interface {
	int
}

// An interface representing all types with underlying type int.
// 表示底层类型为 int 的所有类型的接口
interface {
	~int
}

// An interface representing all types with underlying type int that implement the String method.
// 表示具有实现 String() 方法 和 底层类型为 int 的所有类型的接口。
interface {
	~int
	String() string
}

// An interface representing an empty type set: there is no type that is both an int and a string.
// 表示空类型集的接口: 没有既是 int 又是 string 的类型。
interface {
	int
	string
}
```

​	在形式为`~T`的术语中，`T`的底层类型必须是它自己，而且`T`不能是一个接口。

```go 
type MyInt int

interface {
	~[]byte  // the underlying type of []byte is itself => []byte 的底层类型是其本身
	~MyInt   // illegal: the underlying type of MyInt is not MyInt => 非法的: MyInt的底层类型不是MyInt
	~error   // illegal: error is an interface => 非法的: error 是一个接口
}
```

联合元素表示类型集的联合：

```go 
// The Float interface represents all floating-point types
// (including any named types whose underlying types are
// either float32 or float64).
// Float接口表示所有的浮点类型（包括底层类型为float32或float64的任何命名类型）。
type Float interface {
	~float32 | ~float64
}
```

​	形式为`T`或`~T`的术语中的类型`T`不能是[类型参数](../DeclarationsAndScope#type-parameter-declarations)，所有非接口术语的类型集必须是成对不相交的（类型集的成对交集必须为空）。给定一个类型参数P：

```go 
interface {
	P                // illegal: P is a type parameter => 非法的: P 是一个类型参数
	int | ~P         // illegal: P is a type parameter => 非法的: P 是一个类型参数
	~int | MyInt     // illegal: the type sets for ~int and MyInt are not disjoint (~int includes MyInt)  => 非法的: ~int 和 MyInt 的类型集是相交的(~int 包括 MyInt)
	float32 | Float  // overlapping type sets but Float is an interface  => 重叠的类型集，更进一步说Float也是一个接口
}
```

实现限制：联合(有多个术语)不能包含[预先声明的标识符](../DeclarationsAndScope#predeclared-identifiers)`comparable`或指定方法的接口，或嵌入`comparable`或指定方法的接口。

​	非[基本接口](#basic-interfaces)只能作为类型约束使用，或者作为其他接口的元素作为约束使用。它们不能作为值或变量的类型，也不能作为其他非接口类型的组成部分。

```go 
var x Float                     // illegal: Float is not a basic interface => 非法: Float 不是一个基本接口

var x interface{} = Float(nil)  // illegal => 非法

type Floatish struct {
	f Float                 // illegal => 非法
}
```

接口类型 `T` 不能嵌入任何递归地包含或嵌入 `T` 的类型元素。

```go 
// illegal: Bad cannot embed itself => 非法: Bad 不能嵌入自己
type Bad interface {
	Bad
}

// illegal: Bad1 cannot embed itself using Bad2 => 非法: Bad1不能使用 Bad2嵌入自己
type Bad1 interface {
	Bad2
}
type Bad2 interface {
	Bad1
}

// illegal: Bad3 cannot embed a union containing Bad3 => Bad3 不能嵌入包含 Bad3的联合
type Bad3 interface {
	~int | ~string | Bad3
}
```

#### Implementing an interface 实现一个接口

如果类型`T`实现了接口`I`，则

- `T`不是接口，并且是`I`类型集的元素；或者
- `T`是接口，并且`T`的类型集是`I`的类型集的子集。

如果`T`实现了一个接口，那么`T`类型的值就实现了该接口。

### Map types 映射型

​	映射是一个无序的元素组，由一种类型的元素（称为`元素类型`）组成，由另一种类型的唯一键集（称为`键类型`）进行索引。一个未初始化的映射的值是`nil`。

```
MapType     = "map" "[" KeyType "]" ElementType .
KeyType     = Type .
```

​	[比较运算符](../Expressions#comparison-operators)`==`和`!=`必须为键类型的操作数完全定义；`因此键类型不能是函数、映射或切片`。如果键类型是接口类型，则必须为动态键值定义这些比较运算符；失败将导致[运行时恐慌（run-time panic）](../Run-timePanics)。

```go 
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

​	映射元素的数量被称为它的`长度`。对于一个map `m`来说，它可以用内置函数`len`来发现，并且在运行过程中可能会改变。在运行过程中可以用[赋值](../Statements#assignment-statements)添加元素，用[索引表达式](../Expressions##index-expressions)检索元素；可以用内置函数`delete`删除元素。

​	使用内置函数 `make` 创建一个新的空 map 值，它使用 map 类型和一个可选的容量提示作为参数：

```go 
make(map[string]int)
make(map[string]int, 100)
```

​	初始容量不限制其大小：映射会增长以容纳其中存储的项数，但`nil`映射除外。`nil`映射等同于空映射，`只是不能添加任何元素`。

### Channel types 通道型

​	通道为[并发执行函数](../Statements#go-statements)提供了一种机制，通过[发送](../Statements#send-statements)和[接收](../Expressions#receive-operator)指定元素类型的值进行通信。未初始化的通道的值是`nil`。

```
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .
```

​	可选的`<-`操作符指定了通道的方向：发送或接收。如果指定了方向，则该通道是定向的，否则是双向的。通过[赋值](../Statements#assignment-statements)或显式[转换](../Expressions#conversions)，通道可以被限制为仅发送或仅接收。

```go 
chan T          // can be used to send and receive values of type T => 可用于发送或接收类型为 T 的值
chan<- float64  // can only be used to send float64s => 仅用于发送 float64 类型
<-chan int      // can only be used to receive ints => 仅用于接收 int 类型
```

`<-` 操作符尽可能与最左边的 `chan` 相关联：

```go 
chan<- chan int    // same as chan<- (chan int) => 与 chan<- (chan int) 相同
chan<- <-chan int  // same as chan<- (<-chan int) =>与 chan<- (<-chan int) 相同
<-chan <-chan int  // same as <-chan (<-chan int) => 与 <-chan (<-chan int) 相同
chan (<-chan int)
```

​	可以使用内置函数 `make` 创建一个新的、初始化的 channel 值，它以channel 类型和可选的容量作为参数：

```go 
make(chan int, 100)
```

​	容量(以元素数量为单位)设置通道中缓冲区的大小。如果容量为零或没有指定，则通道是无缓冲的，只有当发送方和接收方都准备好时，通信才会成功。否则，如果缓冲区不满(可继续发送)或不是空的(可继续接收) ，通道会将数据缓冲起来，并且通信在没有阻塞的情况下成功。一个`nil`通道不能用于通信。

!!! info "引用其他书籍"
    以下摘自《Go语言精进之路》第34条 了解channel的妙用 第348页。
    与无缓冲channel 不同，带缓冲channel 可以通过带有 capacity 参数的内置make 函数创建：c:= make(chan  T, capctity)
    由于带缓冲channel 的运行时层实现带有缓冲区，因此对带有缓冲channel的发送操作在缓冲区未满、接收操作在缓冲区非空的情况下是异步的（发送或接收无需阻塞等待）。也就是说，对一个带缓冲channel，在缓冲区无数据或有数据但未满的情况下，对其进行发送操作的goroutine不会阻塞；在缓冲区已满的情况下，对其进行发送操作的goroutine会阻塞；在缓冲区为空的情况下，对其进行接收操作的goroutine亦会阻塞。



​	通道可以用内置函数`close`来关闭。[接收操作符](../Expressions#receive-operator)的多值赋值形式可以用来判断数据是否在通道关闭之前发送出去。
​	任意数量的goroutines都可以通过[发送语句](../Statements#send-statements)、[接收操作](../Expressions#receive-operator)以及对内置函数`cap`和`len`的调用，来操作一个通道。通道是一个先入先出的队列。例如，如果一个goroutine在通道上发送数据，第二个goroutine接收这些数据，那么这些数据将按照发送的顺序被接收。