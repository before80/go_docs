+++
title = "表达式"
date = 2023-05-17T09:59:21+08:00
weight = 11
description = ""
isCJKLanguage = true
type = "docs"
math = true
draft = false

+++
## Expressions 表达式

> 原文：[https://go.dev/ref/spec#Expressions](https://go.dev/ref/spec#Expressions )

An expression specifies the computation of a value by applying operators and functions to operands.

​	表达式通过将运算符和函数应用于操作数来规定值的计算。

### Operands 操作数

Operands denote the elementary values in an expression. An operand may be a literal, a (possibly [qualified](https://go.dev/ref/spec#Qualified_identifiers)) non-[blank](https://go.dev/ref/spec#Blank_identifier) identifier denoting a [constant](https://go.dev/ref/spec#Constant_declarations), [variable](https://go.dev/ref/spec#Variable_declarations), or [function](https://go.dev/ref/spec#Function_declarations), or a parenthesized expression.

​	操作数表示表达式中的基本值。操作数可以是一个字面量，一个表示[常量](../DeclarationsAndScope#constant-declarations)、[变量](../DeclarationsAndScope#variable-declarations-变量声明)或函数的（可以是[限定的](#qualified-identifiers-限定标识符)）非[空白](../DeclarationsAndScope#blank-identifier-空白标识符)标识符，或者一对圆括号内的表达式。

```
Operand     = Literal | OperandName [ TypeArgs ] | "(" Expression ")" .
Literal     = BasicLit | CompositeLit | FunctionLit .
BasicLit    = int_lit | float_lit | imaginary_lit | rune_lit | string_lit .
OperandName = identifier | QualifiedIdent .
```

An operand name denoting a [generic function](https://go.dev/ref/spec#Function_declarations) may be followed by a list of [type arguments](https://go.dev/ref/spec#Instantiations); the resulting operand is an [instantiated](https://go.dev/ref/spec#Instantiations) function.

​	表示泛型函数的操作数名称后面可以跟一个[类型实参](#instantiations-实例化)列表；产生的操作数是一个[实例化过的](#instantiations-实例化)函数。

The [blank identifier](https://go.dev/ref/spec#Blank_identifier) may appear as an operand only on the left-hand side of an [assignment statement](https://go.dev/ref/spec#Assignment_statements).

​	[空白标识符（即`_`）](../DeclarationsAndScope#blank-identifier-空白标识符)只能在[赋值语句](../Statements#assignment-statements-赋值语句)的左侧作为操作数出现。

Implementation restriction: A compiler need not report an error if an operand's type is a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations) with an empty [type set](https://go.dev/ref/spec#Interface_types). Functions with such type parameters cannot be [instantiated](https://go.dev/ref/spec#Instantiations); any attempt will lead to an error at the instantiation site.

​	实现限制：若操作数的类型是具有空[类型集](../Types#interface-types-接口型)的[类型形参](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)，则编译器不必报告错误。具有这种类型形参的函数不能被[实例化](#instantiations-实例化)；任何尝试都会导致实例化处的错误。

### Qualified identifiers 限定标识符

A *qualified identifier* is an identifier qualified with a package name prefix. Both the package name and the identifier must not be [blank](https://go.dev/ref/spec#Blank_identifier).

​	限定标识符是以包名作为前缀限定的标识符。包名和标识符都不能是[空白标识符（即`_`）](../DeclarationsAndScope#blank-identifier-空白标识符)。

```
QualifiedIdent = PackageName "." identifier .
```

A qualified identifier accesses an identifier in a different package, which must be [imported](https://go.dev/ref/spec#Import_declarations). The identifier must be [exported](https://go.dev/ref/spec#Exported_identifiers) and declared in the [package block](https://go.dev/ref/spec#Blocks) of that package.

​	限定标识符可以在不同包中访问一个标识符，但该标识符所在的包必须已经被[导入](../Packages#import-declarations-导入声明)。该标识符必须[可被导出](../DeclarationsAndScope#exported-identifiers-可导出的标识符)并在该包的[package block](../Blocks)中声明。

```go 
math.Sin // denotes the Sin function in package math 
		 //	=> 表示 math 包中的 Sin 函数
```

### Composite literals 复合字面量

Composite literals construct new composite values each time they are evaluated. They consist of the type of the literal followed by a brace-bound list of elements. Each element may optionally be preceded by a corresponding key.

​	复合字面量每次被求值时都会构造新的复合值。`它们由字面量的类型和一个由花括号组成的元素列表组成`。每个元素可以选择在前面加上一个相应的键。

```go
CompositeLit  = LiteralType LiteralValue .
LiteralType   = StructType | ArrayType | "[" "..." "]" ElementType |
                SliceType | MapType | TypeName [ TypeArgs ] .
LiteralValue  = "{" [ ElementList [ "," ] ] "}" .
ElementList   = KeyedElement { "," KeyedElement } .
KeyedElement  = [ Key ":" ] Element .
Key           = FieldName | Expression | LiteralValue .
FieldName     = identifier .
Element       = Expression | LiteralValue .
```

The LiteralType's [core type](https://go.dev/ref/spec#Core_types) `T` must be a struct, array, slice, or map type (the syntax enforces this constraint except when the type is given as a TypeName). The types of the elements and keys must be [assignable](https://go.dev/ref/spec#Assignability) to the respective field, element, and key types of type `T`; there is no additional conversion. The key is interpreted as a field name for struct literals, an index for array and slice literals, and a key for map literals. For map literals, all elements must have a key. It is an error to specify multiple elements with the same field name or constant key value. For non-constant map keys, see the section on [evaluation order](https://go.dev/ref/spec#Order_of_evaluation).

​	LiteralType 的[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)`T`必须是一个结构体、数组、切片或映射类型（语法会强制执行这个约束，除非当类型是作为TypeName给出时）。元素和键的类型必须[可分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`T`类型的对应字段、元素和键类型；不需要进行额外的转换。这里的键被解释为结构体字面量的字段名、数组字面量或切片字面量的索引、映射字面量的键。对于映射字面量，所有的元素必须有一个键。用相同的字段名或常量键值指定多个元素是错误的。对于非常量的映射键，请参见关于[求值顺序](#order-of-evaluation-求值顺序)的章节。

For struct literals the following rules apply:

​	对于结构体字面量来说，以下规则适用：

- A key must be a field name declared in the struct type.
- 键必须是结构体类型中声明的字段名。
- An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared.
- 不包含任何键的元素列表必须按照字段的声明顺序为每个结构体字段列出一个元素。
- If any element has a key, every element must have a key.
- 如果任何元素有一个键，那么每个元素都必须有一个键。
- An element list that contains keys does not need to have an element for each struct field. Omitted fields get the zero value for that field.
- 包含键的元素列表不需要每个结构体字段都有一个元素。省略的字段将获得该字段类型的零值。
- A literal may omit the element list; such a literal evaluates to the zero value for its type.
- `字面量可以省略元素列表；这样的字面量相当对其类型的求值为零值`。
- It is an error to specify an element for a non-exported field of a struct belonging to a different package.
- 为属于不在同一个包中的结构体（即该结构体在其他包中定义）的非可导出字段指定元素是错误的。

Given the declarations 

​	给定声明：

```go 
type Point3D struct { x, y, z float64 }
type Line struct { p, q Point3D }
```

one may write

我们可以这样写：

```go 
// zero value for Point3D 
//=> Point3D 的零值
origin := Point3D{} 

// zero value for line.q.x 
//=> line.q.x 的零值
line := Line{origin, Point3D{y: -4, z: 12.3}}  
```

For array and slice literals the following rules apply:

​	对于数组字面量和切片字面量，以下规则适用：

- Each element has an associated integer index marking its position in the array.
- 每个元素都有一个相关的整数索引，标记其在数组中的位置。
- An element with a key uses the key as its index. The key must be a non-negative constant [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; and if it is typed it must be of [integer type](https://go.dev/ref/spec#Numeric_types).
- 带键的元素使用该键作为其索引。键必须是一个可由`int`类型的值[表示的](../PropertiesOfTypesAndValues#representability-可表示性)非负常数；如果它是有类型的，则它必须是[整数类型](../Types#numeric-types-数值型)。
- An element without a key uses the previous element's index plus one. If the first element has no key, its index is zero.
- 不带键的元素使用前一个元素的索引加1。如果第一个元素没有键，它的索引是0。

[Taking the address](https://go.dev/ref/spec#Address_operators) of a composite literal generates a pointer to a unique [variable](https://go.dev/ref/spec#Variables) initialized with the literal's value.

​	对一个复合字面量[取址](#address-operators-地址运算符)会产生一个指向唯一[变量](../Variables)的指针，该变量用字面量的值初始化。

```go 
var pointer *Point3D = &Point3D{y: 1000}
```

Note that the [zero value](https://go.dev/ref/spec#The_zero_value) for a slice or map type is not the same as an initialized but empty value of the same type. Consequently, taking the address of an empty slice or map composite literal does not have the same effect as allocating a new slice or map value with [new](https://go.dev/ref/spec#Allocation).

​	请注意，切片或映射类型的[零值](../ProgramInitializationAndExecution#the-zero-value-零值)与同一类型的初始化过但为空的值不同。因此，获取一个空切片或空映射复合字面量的地址与用[new](../Built-inFunctions#allocation-分配)分配一个新的切片或映射值的效果不同。

```go 
// p1 points to an initialized, empty slice with value []int{} and length 0 
//=> p1 指向一个值为 []int{} 且 长度为 0 的初始化过的空切片
p1 := &[]int{}    

// p2 points to an uninitialized slice with value nil and length 0 
//=> p2 指向一个值为 nil 且其长度为 0 的未初始化的切片
p2 := new([]int)  
```

The length of an array literal is the length specified in the literal type. If fewer elements than the length are provided in the literal, the missing elements are set to the zero value for the array element type. It is an error to provide elements with index values outside the index range of the array. The notation `...` specifies an array length equal to the maximum element index plus one.

​	数组字面量的长度是字面量类型中指定的长度。如果在字面量上提供的元素少于长度，缺少的元素将被设置为数组元素类型的零值。若提供的元素的索引值超出了数组的索引范围，将导致错误。标记法`...`指定的数组长度等于最大元素索引加1。

```go 
buffer := [10]string{}             // len(buffer) == 10
intSet := [6]int{1, 2, 3, 5}       // len(intSet) == 6
days := [...]string{"Sat", "Sun"}  // len(days) == 2
```

A slice literal describes the entire underlying array literal. Thus the length and capacity of a slice literal are the maximum element index plus one. A slice literal has the form

​	切片字面量描述了整个底层数组字面量。因此，切片字面量的长度和容量是最大元素索引加1。切片字面量的形式是：

```go 
[]T{x1, x2, … xn}
```

and is shorthand for a slice operation applied to an array:

是对数组进行切片操作的简写：

```go 
tmp := [n]T{x1, x2, … xn}
tmp[0 : n]
```

Within a composite literal of array, slice, or map type `T`, elements or map keys that are themselves composite literals may elide the respective literal type if it is identical to the element or key type of `T`. Similarly, elements or keys that are addresses of composite literals may elide the `&T` when the element or key type is `*T`.

​	在一个数组、切片或映射类型`T`的复合字面量中，如果本身是复合字面量的**元素或映射键**与`T`的元素或键类型相同，则可以省略（**元素或映射键的**）相应的字面量类型。同样，当元素或键类型为`*T`时，作为复合字面量地址的元素或键可以省略`&T`。

```go 
[...]Point{{1.5, -3.5}, {0, 0}}     // same as [...]Point{Point{1.5, -3.5}, Point{0, 0}}

[][]int{{1, 2, 3}, {4, 5}}          // same as [][]int{[]int{1, 2, 3}, []int{4, 5}}

[][]Point{{{0, 1}, {1, 2}}}         // same as [][]Point{[]Point{Point{0, 1}, Point{1, 2}}}

map[string]Point{"orig": {0, 0}}    // same as map[string]Point{"orig": Point{0, 0}}

map[Point]string{{0, 0}: "orig"}    // same as map[Point]string{Point{0, 0}: "orig"}

type PPoint *Point
[2]*Point{{1.5, -3.5}, {}}          // same as [2]*Point{&Point{1.5, -3.5}, &Point{}}
[2]PPoint{{1.5, -3.5}, {}}          // same as [2]PPoint{PPoint(&Point{1.5, -3.5}), PPoint(&Point{})}
```

A parsing ambiguity arises when a composite literal using the TypeName form of the LiteralType appears as an operand between the [keyword](https://go.dev/ref/spec#Keywords) and the opening brace of the block of an "if", "for", or "switch" statement, and the composite literal is not enclosed in parentheses, square brackets, or curly braces. In this rare case, the opening brace of the literal is erroneously parsed as the one introducing the block of statements. To resolve the ambiguity, the composite literal must appear within parentheses.

​	当使用LiteralType的TypeName形式的复合字面量`作为操作数`出现在[关键字](../LexicalElements#keywords-关键字)和 "`if`"、"`for` "或 "`switch` "等语句块的`左花括号`之间，并且复合字面量没有被括在圆括号、方括号或花括号中时，会出现解析歧义。在这种罕见的情况下，字面量的左花括号被错误地解析为引入语句块的左花括号。为了解决这个问题，复合字面量`必须出现在圆括号`内。

```go 
if x == (T{a,b,c}[i]) { … }
if (x == T{a,b,c}[i]) { … }
```

Examples of valid array, slice, and map literals: 

​	有效的数组、切片和映射字面量的例子：

```go 
// list of prime numbers => 质数列表
primes := []int{2, 3, 5, 7, 9, 2147483647}

// vowels[ch] is true if ch is a vowel 
//=> 当 ch 是 元音时，vowels[ch] 为 真
vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

// the array [10]float32{-1, 0, 0, 0, -0.1, -0.1, 0, 0, 0, -1}
filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}

// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
noteFrequency := map[string]float32{
	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	"G0": 24.50, "A0": 27.50, "B0": 30.87,
}
```

### Function literals 函数字面量

A function literal represents an anonymous [function](https://go.dev/ref/spec#Function_declarations). Function literals cannot declare type parameters.

​	函数字面量表示一个匿名[函数](../DeclarationsAndScope#function-declarations-函数声明)。函数字面量不能声明`类型参数`。

``` go
functionLit = "func" Signature FunctionBody .
```

```go 
func(a, b int, z float64) bool { return a*b < int(z) }
```

A function literal can be assigned to a variable or invoked directly.

​	函数字面量可以被分配给一个变量或直接调用。

```go 
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```

Function literals are *closures*: they may refer to variables defined in a surrounding function. Those variables are then shared between the surrounding function and the function literal, and they survive as long as they are accessible.

​	函数字面量可以是`闭包`：它们可以引用外层函数中定义的变量。然后，这些变量在外层的函数和函数字面量之间共享，并且只要它们可以被访问，它们就可以一直存在。

### Primary expressions 主表达式

Primary expressions are the operands for unary and binary expressions.

​	主表达式是一元、二元表达式的操作数。

```go
PrimaryExpr =
	Operand |
	Conversion |
	MethodExpr |
	PrimaryExpr Selector |
	PrimaryExpr Index |
	PrimaryExpr Slice |
	PrimaryExpr TypeAssertion |
	PrimaryExpr Arguments .

Selector       = "." identifier .
Index          = "[" Expression "]" .
Slice          = "[" [ Expression ] ":" [ Expression ] "]" |
                 "[" [ Expression ] ":" Expression ":" Expression "]" .
TypeAssertion  = "." "(" Type ")" .
Arguments      = "(" [ ( ExpressionList | Type [ "," ExpressionList ] ) [ "..." ] [ "," ] ] ")" .
```

```go 
x
2
(s + ".txt")
f(3.1415, true)
Point{1, 2}
m["foo"]
s[i : j + 1]
obj.color
f.p[i].x()
```

### Selectors 选择器

For a [primary expression](https://go.dev/ref/spec#Primary_expressions) `x` that is not a [package name](https://go.dev/ref/spec#Package_clause), the *selector expression*

​	对于[主表达式](#primary-expressions-主表达式)`x`（不是[包名](../Packages#package-clause-包子句)）来说，选择器表达式：

```go 
x.f
```

denotes the field or method `f` of the value `x` (or sometimes `*x`; see below). The identifier `f` is called the (field or method) *selector*; it must not be the [blank identifier](https://go.dev/ref/spec#Blank_identifier). The type of the selector expression is the type of `f`. If `x` is a package name, see the section on [qualified identifiers](https://go.dev/ref/spec#Qualified_identifiers). 

表示值`x`（有时是`*x`；见下文）的字段或方法`f`。标识符`f`被称为（字段或方法）`选择器`；它不能是[空白标识符](../DeclarationsAndScope#blank-identifierr-空白标识符)。选择器表达式的类型是`f`的类型。若`x`是包名，请参见关于[限定标识符](#qualified-identifiers-限定标识符)一节。

A selector `f` may denote a field or method `f` of a type `T`, or it may refer to a field or method `f` of a nested [embedded field](https://go.dev/ref/spec#Struct_types) of `T`. The number of embedded fields traversed to reach `f` is called its *depth* in `T`. The depth of a field or method `f` declared in `T` is zero. The depth of a field or method `f` declared in an embedded field `A` in `T` is the depth of `f` in `A` plus one.

​	选择器`f`可以表示类型`T`的`f`字段或`f`方法，也可以指代`T`的[嵌入字段](../Types#struct-types-结构体型)或嵌入方法`f`。在`T`的一个嵌入字段`A`中声明的字段或方法`f`的深度是`A`中`f`的深度加1。

The following rules apply to selectors:

​	以下规则适用于选择器：

1. For a value `x` of type `T` or `*T` where `T` is not a pointer or interface type, `x.f` denotes the field or method at the shallowest depth in `T` where there is such an `f`. If there is not exactly [one `f`](https://go.dev/ref/spec#Uniqueness_of_identifiers) with shallowest depth, the selector expression is illegal.

2. 对于类型为`T`或`*T`（`T`不是指针或接口类型）的值`x`，`x.f`表示`T`中存在这样一个最浅深度的字段或方法`f`。如果不是恰好有[仅有一个](../DeclarationsAndScope#uniqueness-of-identifiers-标识符的唯一性)`f`在最浅深度的话，那么选择器表达式是非法的。

   > 个人注释
   >
   > ​	这里额的[仅有一个](../DeclarationsAndScope#uniqueness-of-identifiers-标识符的唯一性)是说在最浅深度！
   >
   > ```go
   > package main
   > 
   > import "fmt"
   > 
   > type A struct {
   > 	Name string
   > }
   > 
   > type B struct {
   > 	A
   > 	Name string
   > }
   > 
   > func main() {
   > 	st := B{A: A{Name: "a"}, Name: "b"}
   > 	fmt.Printf("%+v\n", st)     // {A:{Name:a} Name:b}
   > 	fmt.Printf("%#v\n", st)     // main.B{A:main.A{Name:"a"}, Name:"b"}
   > 	fmt.Printf("%v\n", st.Name) // b
   > }
   > 
   > ```
   >
   > 

3. For a value `x` of type `I` where `I` is an interface type, `x.f` denotes the actual method with name `f` of the dynamic value of `x`. If there is no method with name `f` in the [method set](https://go.dev/ref/spec#Method_sets) of `I`, the selector expression is illegal.

4. 对于接口类型`I`的值`x`，`x.f`表示动态值`x`的名为`f`的实际方法。如果在`I`的[方法集](../PropertiesOfTypesAndValues#method-sets-方法集)中没有名为`f`的方法，那么选择器表达式是非法的。

   > 个人注释
   >
   > ```go
   > package main
   > 
   > import "fmt"
   > 
   > type Adder interface {
   > 	Add(int, int) int
   > }
   > 
   > type Op struct {
   > }
   > 
   > func (o Op) Add(a int, b int) (total int) {
   > 	return a + b
   > }
   > 
   > func main() {
   > 	var i Adder
   > 	a := Op{}
   > 	fmt.Printf("%+v\n", i) //<nil>
   > 	fmt.Printf("%#v\n", a) //main.Op{}
   > 	i = a
   > 	fmt.Printf("%v\n", i.Add(1, 2)) // 3
   > 	//fmt.Printf("%v\n", i.Sub(1, 2)) // i.Sub undefined (type Adder has no field or method Sub)
   > }
   > 
   > ```
   >
   > 

5. As an exception, if the type of `x` is a [defined](https://go.dev/ref/spec#Type_definitions) pointer type and `(*x).f` is a valid selector expression denoting a field (but not a method), `x.f` is shorthand for `(*x).f`.

6. 作为例外，如果`x`的类型是一个[已定义的](../DeclarationsAndScope#type-definitions-类型定义)指针类型，并且`(*x).f`是一个有效的表示一个字段（不是一个方法）的选择器表达式，那么`x.f`是`(*x).f`的简写。

   > 个人注释
   >
   > ```go
   > package main
   > 
   > import "fmt"
   > 
   > type Person struct {
   > 	Name *string
   > 	Age  int
   > }
   > 
   > func main() {
   > 	name1 := "zlongx1"
   > 	name2 := "zlongx2"
   > 	v := Person{Name: &name1, Age: 31}
   > 	p := &Person{Name: &name2, Age: 32}
   > 
   > 	fmt.Printf("%#v\n", v)         // main.Person{Name:(*string)(0xc00004e270)}
   > 	fmt.Printf("%#v\n", p)         // &main.Person{Name:(*string)(0xc00004e280)}
   > 	fmt.Printf("%#v\n", *p)        // main.Person{Name:(*string)(0xc00004e280)}
   > 	fmt.Printf("%#v\n", v.Name)    // (*string)(0xc00004e270)
   > 	fmt.Printf("%#v\n", p.Name)    // (*string)(0xc00004e280)
   > 	fmt.Printf("%#v\n", (*p).Name) // (*string)(0xc00004e280)
   > 	fmt.Printf("%#v\n", v.Age)     // 31
   > 	fmt.Printf("%#v\n", p.Age)     // 32
   > 	fmt.Printf("%#v\n", (*p).Age)  // 32
   > }
   > ```
   >
   > 

7. In all other cases, `x.f` is illegal.

8. 在所有其它情况下，`x.f`是非法的。

9. If `x` is of pointer type and has the value `nil` and `x.f` denotes a struct field, assigning to or evaluating `x.f` causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics). 

10. 如果`x`是值为`nil`的指针类型，并且`x.f`表示一个结构体字段，那么赋值或计算`x.f`会引起[运行时恐慌](../Run-timePanics)。

   > 个人注释
   >
   > ```go
   > package main
   > 
   > import "fmt"
   > 
   > type Person struct {
   > 	Name string
   > 	Age  int
   > }
   > 
   > func main() {
   > 	var a Person
   > 	var b *Person
   > 	fmt.Printf("%#v\n", a) //main.Person{Name:"", Age:0}
   > 	fmt.Printf("%#v\n", b) //(*main.Person)(nil)
   > 
   > 	fmt.Printf("Name=\"%v\"\n", a.Name) //Name=""
   > 	a.Name = "a"
   > 	fmt.Printf("Name=\"%v\"\n", a.Name) //Name="a"
   > 	//b.Name = "b" // runtime error: invalid memory address or nil pointer dereference
   > 	//fmt.Printf("Name=\"%v\"\n", b.Name) //panic: runtime error: invalid memory address or nil pointer dereference
   > }
   > 
   > ```
   >
   > 

11. If `x` is of interface type and has the value `nil`, [calling](https://go.dev/ref/spec#Calls) or [evaluating](https://go.dev/ref/spec#Method_values) the method `x.f` causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics).

12. 如果`x`是值为`nil`的接口类型，那么[调用](#calls-调用)或[计值](#method-values-方法值)`x.f`方法会引起[运行时恐慌](../Run-timePanics)。

    > 个人注释
    >
    > ```go
    > package main
    > 
    > import "fmt"
    > 
    > type Adder interface {
    > 	Add(int, int) int
    > }
    > 
    > func main() {
    > 	var i Adder
    > 	fmt.Printf("%#v\n", i)          //<nil>
    > 	fmt.Printf("%v\n", i.Add(1, 2)) //panic: runtime error: invalid memory address or nil pointer dereference
    > }
    > 
    > ```
    >
    > 

For example, given the declarations:

​	例如，给定声明：

```go 
type T0 struct {
	x int
}

func (*T0) M0()

type T1 struct {
	y int
}

func (T1) M1()

type T2 struct {
	z int
	T1
	*T0
}

func (*T2) M2()

type Q *T2

var t T2     // with t.T0 != nil => 假定 t.T0 != nil
var p *T2    // with p != nil and (*p).T0 != nil => 假定 p != nil 并且 (*p).T0 != nil
var q Q = p
```

one may write: 

则我们可以这样写：

```go 
t.z          // t.z
t.y          // t.T1.y
t.x          // (*t.T0).x

p.z          // (*p).z
p.y          // (*p).T1.y
p.x          // (*(*p).T0).x

q.x          // (*(*q).T0).x        (*q).x is a valid field selector => (*q).x 是一个有效的字段选择器

p.M0()       // ((*p).T0).M0()      M0 expects *T0 receiver
p.M1()       // ((*p).T1).M1()      M1 expects T1 receiver
p.M2()       // p.M2()              M2 expects *T2 receiver
t.M2()       // (&t).M2()           M2 expects *T2 receiver, see section on Calls
```

but the following is invalid:

但下面的内容是无效的：

```go 
q.M0()       // (*q).M0 is valid but not a field selector => (*q).M0是有效的，但不是字段选择器
```

> 个人注释
>
> ​	以上示例的完整代码如下：
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"reflect"
> )
> 
> type T0 struct {
> 	x int
> }
> 
> func (*T0) M0() string {
> 	return "in M0 func"
> }
> 
> type T1 struct {
> 	y int
> }
> 
> func (T1) M1() string {
> 	return "in M1 func"
> }
> 
> type T2 struct {
> 	z int
> 	T1
> 	*T0
> }
> 
> func (*T2) M2() string {
> 	return "in M2 func"
> }
> 
> type Q *T2
> 
> func DumpMethodSet(i interface{}) {
> 	v := reflect.TypeOf(i)
> 	elemTyp := v.Elem()
> 
> 	n := elemTyp.NumMethod()
> 	if n == 0 {
> 		fmt.Printf("%s's method set is empty!\n", elemTyp)
> 	}
> 
> 	fmt.Printf("%s's method set:\n", elemTyp)
> 	for j := 0; j < n; j++ {
> 		fmt.Println("-", elemTyp.Method(j).Name)
> 	}
> 	fmt.Println()
> }
> 
> func main() {
> 	var t T2 = T2{T0: &T0{}}
> 	var p *T2 = &T2{T0: &T0{}}
> 	var q Q = p
> 
> 	fmt.Printf("t=%+v\n", t) // t={z:0 T1:{y:0} T0:0xc00001a0a8}
> 	fmt.Printf("p=%+v\n", p) // p=&{z:0 T1:{y:0} T0:0xc00001a0c0}
> 	fmt.Printf("q=%+v\n", q) // q=&{z:0 T1:{y:0} T0:0xc00001a0c0}
> 	//main.T2's method set:
> 	//- M0
> 	//- M1
> 	DumpMethodSet(&t)
> 
> 	//*main.T2's method set:
> 	//- M0
> 	//- M1
> 	//- M2
> 	DumpMethodSet(&p)
> 
> 	//main.Q's method set is empty!
> 	//main.Q's method set:
> 	DumpMethodSet(&q)
> 
> 	//main.T2's method set:
> 	//- M0
> 	//- M1
> 	DumpMethodSet(&(*q))
> 
> 	fmt.Printf("t.z=%v\n", t.z) // t.z=0
> 	fmt.Printf("t.y=%v\n", t.y) // t.y=0
> 	fmt.Printf("t.x=%v\n", t.x) // t.x=0
> 
> 	fmt.Printf("p.z=%v\n", p.z) // p.z=0
> 	fmt.Printf("p.y=%v\n", p.y) // p.y=0
> 	fmt.Printf("p.x=%v\n", p.x) // p.x=0
> 
> 	fmt.Printf("q.z=%v\n", q.z) // q.z=0
> 	fmt.Printf("q.y=%v\n", q.y) // q.y=0
> 	fmt.Printf("q.x=%v\n", q.x) // q.x=0
> 
> 	fmt.Printf("t.M0()=%v\n", t.M0()) // t.M0()=in M0 func
> 	fmt.Printf("t.M1()=%v\n", t.M1()) // t.M1()=in M1 func
> 	fmt.Printf("t.M2()=%v\n", t.M2()) // t.M2()=in M2 func
> 
> 	fmt.Printf("p.M0()=%v\n", p.M0()) // p.M0()=in M0 func
> 	fmt.Printf("p.M1()=%v\n", p.M1()) // p.M1()=in M1 func
> 	fmt.Printf("p.M2()=%v\n", p.M2()) // p.M2()=in M2 func
> 
> 	//fmt.Printf("q.M0()=%v\n", q.M0()) // q.M0 undefined (type Q has no field or method M0)
> 	//fmt.Printf("q.M1()=%v\n", q.M1()) // q.M1 undefined (type Q has no field or method M1)
> 	//fmt.Printf("q.M2()=%v\n", q.M2()) // q.M2 undefined (type Q has no field or method M2)
> 
> 	fmt.Printf("(*q).M0()=%v\n", (*q).M0()) // (*q).M0()=in M0 func
> 	fmt.Printf("(*q).M1()=%v\n", (*q).M1()) // (*q).M1()=in M1 func
> 	fmt.Printf("(*q).M2()=%v\n", (*q).M2()) // (*q).M2()=in M2 func
> }
> 
> ```
>
> 

### Method expressions 方法表达式

If `M` is in the [method set](https://go.dev/ref/spec#Method_sets) of type `T`, `T.M` is a function that is callable as a regular function with the same arguments as `M` prefixed by an additional argument that is the receiver of the method.

​	如果`M`在类型`T`的[方法集](../PropertiesOfTypesAndValues#method-sets-方法集)中，那么`T.M`是一个可以作为普通函数来调用的函数，其实参与`M`相同，不过其前缀有一个额外的（作为该方法的接收器的）实参。

```go
MethodExpr    = ReceiverType "." MethodName .
ReceiverType  = Type .
```

Consider a struct type `T` with two methods, `Mv`, whose receiver is of type `T`, and `Mp`, whose receiver is of type `*T`.

​	考虑一个结构体类型`T`，它有两个方法：`Mv`，其接收器类型为`T`，和`Mp`，其接收器类型为`*T`。

```go 
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // 值类型的接收器
func (tp *T) Mp(f float32) float32 { return 1 }  // 指针类型的接收器

var t T
```

The expression

​	表达式：

```go 
T.Mv
```

yields a function equivalent to `Mv` but with an explicit receiver as its first argument; it has signature

生成一个与`Mv`等价的函数，但是它的第一个实参是一个显式的接收器；它具有以下签名：

```go 
func(tv T, a int) int
```

That function may be called normally with an explicit receiver, so these five invocations are equivalent:

​	该函数可以在带有一个显式接收器的情况下正常调用，因此如下这五种调用是等同的：

```go 
t.Mv(7)
T.Mv(t, 7)
(T).Mv(t, 7)
f1 := T.Mv; f1(t, 7)
f2 := (T).Mv; f2(t, 7)
```

Similarly, the expression

​	类似地，表达式

```go 
(*T).Mp
```

yields a function value representing `Mp` with signature

生成一个代表`Mp`的函数值，它的签名是：

```go 
func(tp *T, f float32) float32
```

For a method with a value receiver, one can derive a function with an explicit pointer receiver, so

​	对于一个`带值接收器`的方法，可以`推导出`一个带有显式指针接收器的函数，因此

```go 
(*T).Mv
```

yields a function value representing `Mv` with signature

生成了一个代表`Mv`的函数值，它的签名是：

```go 
func(tv *T, a int) int
```

Such a function indirects through the receiver to create a value to pass as the receiver to the underlying method; the method does not overwrite the value whose address is passed in the function call.

​	这样的函数通过接收器`间接地`创建了一个值，作为接收器传递给底层方法；该方法不会覆盖（其地址在函数调用中被传递的）那个值。=>该怎么翻译？？

​	这样的一个函数间接地通过其接收器创建了（用来作为接收器传递给其底层方法的）一个值；该（底层）方法不会覆盖这个值（因这个值的地址在这个函数调用才会被传递）

The final case, a value-receiver function for a pointer-receiver method, is illegal because pointer-receiver methods are not in the method set of the value type.

​	最后一种情况，将一个带指针接收器的方法`当做`一个带值接收器的函数，是非法的，因为指针接收器的方法不在值类型的方法集中。

Function values derived from methods are called with function call syntax; the receiver is provided as the first argument to the call. That is, given `f := T.Mv`, `f` is invoked as `f(t, 7)` not `t.f(7)`. To construct a function that binds the receiver, use a [function literal](https://go.dev/ref/spec#Function_literals) or [method value](https://go.dev/ref/spec#Method_values).

​	从方法中推导出来的函数值是用`函数调用语法`来调用的；接收器被作为调用的第一个实参来提供。也就是说，`f := T.Mv`中的`f`是作为`f(t, 7)`被调用，而不是`t.f(7)`。要构造一个绑定接收器的函数，可以使用[函数字面量](#function-literals-函数字面量)或[方法值](#method-values-方法值)。

It is legal to derive a function value from a method of an interface type. The resulting function takes an explicit receiver of that interface type.

​	从接口类型的方法中推导出来函数值是合法的。这样的函数需要一个该接口类型的显式接收器。

> 个人注释
>
> ​	以上示例的完整代码如下：TODO
>
> ```go
> package main
> 
> import "fmt"
> 
> type T struct {
> 	a int
> }
> 
> func (tv T) Mv(a int) int {
> 	return a + tv.a
> }
> 
> func (tp *T) Mp(f float32) float32 {
> 	return f + float32(tp.a)
> }
> 
> func funcMv1(tv T, a int) int {
> 	return tv.a + a
> }
> 
> func funcMv2(tv *T, a int) int {
> 	return tv.a + a
> }
> 
> func funcMp(tp *T, a float32) float32 {
> 	return float32(tp.a) + a
> }
> 
> func main() {
> 	var tv T = T{}
> 	var tp *T = &T{}
> 
> 	tv.a = 1
> 	fmt.Printf("tv.Mv(7)=%v\n", tv.Mv(7))               // tv.Mv(7)=8
> 	fmt.Printf("T.Mv(tv, 7)=%v\n", T.Mv(tv, 7))         // T.Mv(tv, 7)=8
> 	fmt.Printf("(T).Mv(tv, 7)=%v\n", (T).Mv(tv, 7))     // (T).Mv(tv, 7)=8
> 	fmt.Printf("(*T).Mv(&tv, 7)=%v\n", (*T).Mv(&tv, 7)) // (*T).Mv(&tv, 7)=8
> 	fmt.Printf("tv.Mp(7)=%v\n", tv.Mp(7))               // tv.Mp(7)=8
> 
> 	fmt.Println("-- a=1 ---------------------------------")
> 
> 	tv.a = 2
> 	f1 := T.Mv
> 	fmt.Printf("f1(tv, 7)=%v\n", f1(tv, 7))           // f1(tv, 7)=9
> 	fmt.Printf("funcMv1(tv, 7)=%v\n", funcMv1(tv, 7)) // funcMv1(tv, 7)=9
> 
> 	fmt.Println("-- a=2 ---------------------------------")
> 
> 	tv.a = 3
> 	f2 := (*T).Mv
> 	fmt.Printf("f2(&tv, 7)=%v\n", f2(&tv, 7))           // f2(&tv, 7)=10
> 	fmt.Printf("funcMv2(&tv, 7)=%v\n", funcMv2(&tv, 7)) // funcMv2(&tv, 7)=10
> 
> 	fmt.Println("-- a=3 ---------------------------------")
> 
> 	tv.a = 4
> 	tp.a = 4
> 	fmt.Printf("tp.Mv(7)=%v\n", tp.Mv(7))               // tp.Mv(7)=11
> 	fmt.Printf("tp.Mp(7)=%v\n", tp.Mp(7))               // tp.Mp(7)=11
> 	fmt.Printf("(*T).Mp(tp, 7)=%v\n", (*T).Mp(tp, 7))   // (*T).Mp(tp, 7)=11
> 	fmt.Printf("(*T).Mv(tp, 7)=%v\n", (*T).Mv(tp, 7))   // (*T).Mv(tp, 7)=11
> 	fmt.Printf("(*T).Mv(&tv, 7)=%v\n", (*T).Mv(&tv, 7)) // (*T).Mv(&tv, 7)=11
> 
> 	fmt.Println("-- a=4 ---------------------------------")
> 
> 	tv.a = 5
> 	tp.a = 5
> 	f3 := (*T).Mp
> 	fmt.Printf("f3(&tv, 7)=%v\n", f3(&tv, 7))         // f3(&tv, 7)=12
> 	fmt.Printf("f3(tp, 7)=%v\n", f3(tp, 7))           // f3(tp, 7)=12
> 	fmt.Printf("funcMp(&tv, 7)=%v\n", funcMp(&tv, 7)) // funcMp(&tv, 7)=12
> 	fmt.Printf("funcMp(tp, 7)=%v\n", funcMp(tp, 7))   // funcMp(tp, 7)=12
> 
> 	fmt.Println("-- a=5 ---------------------------------")
> }
> 
> ```
>
> 

### Method values 方法值

If the expression `x` has static type `T` and `M` is in the [method set](https://go.dev/ref/spec#Method_sets) of type `T`, `x.M` is called a *method value*. The method value `x.M` is a function value that is callable with the same arguments as a method call of `x.M`. The expression `x` is evaluated and saved during the evaluation of the method value; the saved copy is then used as the receiver in any calls, which may be executed later.

​	如果表达式`x`有静态类型`T`，并且`M`在类型`T`的[方法集](../PropertiesOfTypesAndValues#method-sets-方法集)中，那么`x.M`被称为一个`方法值`。方法值`x.M`是一个可调用的函数值，其实参与`x.M`的方法调用相同。表达式`x`在方法值的求值过程中被求值 、 保存；然后保存的副本被用作任何调用中的接收器上，这些调用可能在以后执行。

```go 
type S struct { *T }
type T int
func (t T) M() { print(t) }

t := new(T)
s := S{T: t}
f := t.M                    // 接收器 *t 被求值，并被保存在 f 中
g := s.M                    // 接收器 *(s.T) 被求值，并被保存在 g 中
*t = 42                     // 不会影响 保存在 f 和 g 中的接收器
```

The type `T` may be an interface or non-interface type.

​	类型`T`既可以是接口类型，也可以是非接口类型。

As in the discussion of [method expressions](https://go.dev/ref/spec#Method_expressions) above, consider a struct type `T` with two methods, `Mv`, whose receiver is of type `T`, and `Mp`, whose receiver is of type `*T`.

​	如同上面对[方法表达式](#method-expressions-方法表达式)的讨论，考虑一个结构体类型`T`，它有两个方法：`Mv`，其接收器类型为`T`，和`Mp`，其接收器类型为`*T`。

```go 
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // 值接收器
func (tp *T) Mp(f float32) float32 { return 1 }  // 指针接收器

var t T
var pt *T
func makeT() T
```

The expression

​	表达式

```go 
t.Mv
```

yields a function value of type

​	生成一个类型如下的函数值：

```go 
func(int) int
```

These two invocations are equivalent:

​	这两种调用是等价的：

```go 
t.Mv(7)
f := t.Mv; f(7)
```

Similarly, the expression

​	同样地，表达式

```go 
pt.Mp
```

yields a function value of type

​	生成一个类型如下的函数值：

```go 
func(float32) float32
```

As with [selectors](https://go.dev/ref/spec#Selectors), a reference to a non-interface method with a value receiver using a pointer will automatically dereference that pointer: `pt.Mv` is equivalent to `(*pt).Mv`.

​	和[选择器](#selectors-选择器)一样，若对以值作为接收器的非接口方法，使用指针来引用，则（Go语言）将自动解除对该指针的引用：`pt.Mv`等同于`(*pt).Mv`。

As with [method calls](https://go.dev/ref/spec#Calls), a reference to a non-interface method with a pointer receiver using an addressable value will automatically take the address of that value: `t.Mp` is equivalent to `(&t).Mp`.

​	和[方法调用](#calls-调用)一样，若对以指针作为接收器的非接口方法，使用可寻址的值来引用，则（Go语言）将自动获取该值的地址：`t.Mp`等同于`(&t).Mp`。

```go 
f := t.Mv; f(7)   // like t.Mv(7)
f = pt.Mp; f(7)  // like pt.Mp(7)
f = pt.Mv; f(7)  // like (*pt).Mv(7)
f = t.Mp; f(7)   // like (&t).Mp(7)
f = makeT().Mp   // 无效的：makeT() 的结果是 不可寻址的	
```

> ​	个人注释
>
> ​	以上示例的完整代码如下：
>
> ```go
> 
> ```
>
> 

Although the examples above use non-interface types, it is also legal to create a method value from a value of interface type.

​	尽管上面的例子使用了非接口类型，但从接口类型的值中创建一个方法值也是合法的。

```go 
var i interface { M(int) } = myVal
f := i.M; f(7)  // like i.M(7)
```

> 个人注释
>
> ​	以上示例的完整代码如下：
>
> ```go
> package main
> 
> import "fmt"
> 
> type MyVal struct{}
> 
> func (m MyVal) M(a int) int {
> 	return a
> }
> 
> func main() {
> 	myVal := MyVal{}
> 	var i interface{ M(int) int } = myVal
> 
> 	f := i.M
> 	fmt.Printf("%v\n", f(7))   // 7
> 	fmt.Printf("%v\n", i.M(7)) // 7
> }
> 
> ```
>
> 

### Index expressions 索引表达式

A primary expression of the form

​	若主表达式的形式是：

```go 
a[x]
```

denotes the element of the array, pointer to array, slice, string or map `a` indexed by `x`. The value `x` is called the *index* or *map key*, respectively. The following rules apply:

则表示可用`x`来检索的数组`a`、数组指针`a`、切片`a`、字符串`a`或映射`a`的元素，`x`分别被称为`索引`或`映射键`。以下规则适用：

If `a` is neither a map nor a type parameter:

​	如果`a`既不是映射也不是类型参数：

- the index `x` must be an untyped constant or its [core type](https://go.dev/ref/spec#Core_types) must be an [integer](https://go.dev/ref/spec#Numeric_types)
- 索引`x`必须是一个无类型的常量，或者其[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)必须是[整数类型](../Types#numeric-types-数值型)
- a constant index must be non-negative and [representable](https://go.dev/ref/spec#Representability) by a value of type `int`
- 常量索引必须是非负且可以用`int`类型的值来[表示](../PropertiesOfTypesAndValues#representability-可表示性)
- a constant index that is untyped is given type `int`
- 无类型常量索引会被赋予`int`类型。
- the index `x` is *in range* if `0 <= x < len(a)`, otherwise it is *out of range*
- 如果`0 <= x < len(a)`，则索引`x`在范围内，否则就超出了范围。

For `a` of [array type](https://go.dev/ref/spec#Array_types) `A`:

​	对于数组类型`A`的`a`：

- a [constant](https://go.dev/ref/spec#Constants) index must be in range
- [常量](../Constants)索引必须在范围内
- if `x` is out of range at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs
- 如果`x`在运行时超出了范围，就会发生[运行时恐慌](../Run-timePanics)
- a[x]` is the array element at index `x` and the type of `a[x]` is the element type of `A
- `a[x]`是索引为`x`的数组元素，`a[x]`的类型是`A`的元素类型。

For `a` of [pointer](https://go.dev/ref/spec#Pointer_types) to array type:

​	对于数组类型[指针](../Types#pointer-types-指针型)的`a`：

- a[x]` is shorthand for `(*a)[x]
- `a[x]` 是 `(*a)[x]`的简写

For `a` of [slice type](https://go.dev/ref/spec#Slice_types) `S`:

​	对于[切片类型](../Types#slice-types-切片型)`S`的`a`：

- if `x` is out of range at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs
- 如果`x`在运行时超出了范围，就会发生[运行时恐慌](../Run-timePanics)
- a[x]` is the slice element at index `x` and the type of `a[x]` is the element type of `S
- `a[x]`是索引`x`处的切片元素，`a[x]`的类型是`S`的元素类型。

For `a` of [string type](https://go.dev/ref/spec#String_types):

​	对于[字符串类型](../Types#string-types)的`a`：

- a [constant](https://go.dev/ref/spec#Constants) index must be in range if the string `a` is also constant
- 如果字符串`a`是常量，那么[常量](../Constants)索引必须在范围内
- if `x` is out of range at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs
- 如果`x`在运行时超出了范围，就会发生[运行时恐慌](../Run-timePanics)
- `a[x]`  is the non-constant byte value at index `x` and the type of `a[x]` is `byte
- `a[x]`是索引`x`处的非常量字节值，`a[x]`的类型是`byte`。
- `a[x]` may not be assigned to
- `a[x]`不能被赋值

For `a` of [map type](https://go.dev/ref/spec#Map_types) `M`:

​	对于[映射类型](../Types#map-types-映射型)为`M`的`a`：

- `x`'s type must be [assignable](https://go.dev/ref/spec#Assignability) to the key type of `M
- `x`的类型必须可以被[分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`M`的键类型
- if the map contains an entry with key `x`, `a[x]` is the map element with key `x` and the type of `a[x]` is the element type of `M`
- 如果映射中有键`x`的项，那么`a[x]`就是键`x`的映射元素，`a[x]`的类型就是`M`的元素类型
- if the map is `nil` or does not contain such an entry, `a[x]` is the [zero value](https://go.dev/ref/spec#The_zero_value) for the element type of `M`
- 如果映射为`nil`或者不包含任何项，`a[x]`是`M`的元素类型的[零值](../ProgramInitializationAndExecution#the-zero-value-零值)。

For `a` of [type parameter type](https://go.dev/ref/spec#Type_parameter_declarations) `P`:

​	对于[参数类型](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)为`P`的`a`：

- The index expression `a[x]` must be valid for values of all types in `P`'s type set.
- 索引表达式`a[x]`必须对`P`的类型集中的所有类型的值有效。
- The element types of all types in `P`'s type set must be identical. In this context, the element type of a string type is `byte`.
- `P`的类型集中所有类型的元素类型必须是相同的。在此上下文中，字符串类型的元素类型是`byte`。
- If there is a map type in the type set of `P`, all types in that type set must be map types, and the respective key types must be all identical.
- 如果在`P`的类型集中有一个映射类型，那么该类型集中的所有类型必须是映射类型，且对应的键类型必须都是一致的。
- `a[x]` is the array, slice, or string element at index `x`, or the map element with key `x` of the type argument that `P` is instantiated with, and the type of `a[x]` is the type of the (identical) element types.
- `a[x]`是索引为`x`的数组、切片或字符串元素，或者`P`实例化的类型实参中键为`x`的映射元素，`a[x]`的类型是（一致的）元素类型的类型。
- `a[x]` may not be assigned to if `P`'s type set includes string types.
- 如果`P`的类型集包括字符串类型，则`a[x]`不能再被赋值。

Otherwise `a[x]` is illegal.

​	否则`a[x]`是非法的。

An index expression on a map `a` of type `map[K]V` used in an [assignment statement](https://go.dev/ref/spec#Assignment_statements) or initialization of the special form

​	若将类型为`map[K]V`的映射`a`上的索引表达式使用在[赋值语句](../Statements#assignment-statements-赋值语句)或特殊格式的初始化中：

```go 
v, ok = a[x]
v, ok := a[x]
var v, ok = a[x]
```

yields an additional untyped boolean value. The value of `ok` is `true` if the key `x` is present in the map, and `false` otherwise.

将产生一个额外的`无类型`布尔值。如果键`x`存在于映射中，`ok`的值为`true`，否则为`false`。

Assigning to an element of a `nil` map causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics).

​	若给`nil`映射的元素赋值，将导致[运行时恐慌](../Run-timePanics)。

### Slice expressions 切片表达式

Slice expressions construct a substring or slice from a string, array, pointer to array, or slice. There are two variants: a simple form that specifies a low and high bound, and a full form that also specifies a bound on the capacity.

​	切片表达式从一个字符串、数组、数组指针或切片中构造一个子串或切片。有两种变体：一种是指定低位和高位边界的简单形式，另一种是同时也指定容量的完整形式。

> 个人注释
>
> ```c
> package main
> 
> import "fmt"
> 
> func main() {
> 	str1 := "abcdefghijklmn"
> 	s1 := str1[0:3]
> 	//s2 := str[0:3:4] // invalid operation: 3-index slice of string
> 	fmt.Printf("%T,str1=%v\n", str1, str1) // string,str1=abcdefghijklmn
> 	fmt.Printf("%T,s1=%v\n", s1, s1)       // string,s1=abc
> 
> 	arr1 := [...]int{0, 1, 2, 3, 4, 5}
> 	s20 := arr1[0:3]
> 	s21 := arr1[0:3:5]
> 	fmt.Printf("%T,arr1=%v,len=%d,cap=%d\n", arr1, arr1, len(arr1), cap(arr1)) // [6]int,arr1=[0 1 2 3 4 5],len=6,cap=6
> 	fmt.Printf("%T,s20=%v,len=%d,cap=%d\n", s20, s20, len(s20), cap(s20))      // []int,s20=[0 1 2],len=3,cap=6
> 	fmt.Printf("%T,s21=%v,len=%d,cap=%d\n", s21, s21, len(s21), cap(s21))      // []int,s21=[0 1 2],len=3,cap=5
> 
> 	sli := []int{0, 1, 2, 3, 4, 5}
> 	s30 := sli[0:3]
> 	s31 := sli[0:3:5]
> 	fmt.Printf("%T,sli=%v,len=%d,cap=%d\n", sli, sli, len(sli), cap(sli)) // []int,sli=[0 1 2 3 4 5],len=6,cap=6
> 	fmt.Printf("%T,s30=%v,len=%d,cap=%d\n", s30, s30, len(s30), cap(s30)) // []int,s30=[0 1 2],len=3,cap=6
> 	fmt.Printf("%T,s31=%v,len=%d,cap=%d\n", s31, s31, len(s31), cap(s31)) // []int,s31=[0 1 2],len=3,cap=6
> 
> 	arr2 := &[...]int{0, 1, 2, 3, 4, 5}
> 	s40 := arr2[0:3]
> 	s41 := arr2[0:3:5]
> 	s40_1 := (*arr2)[0:3]
> 	s41_1 := (*arr2)[0:3:5]
> 	fmt.Printf("%T,arr2=%v,len=%d,cap=%d\n", arr2, arr2, len(arr2), cap(arr2))      // *[6]int,arr2=&[0 1 2 3 4 5],len=6,cap=6
> 	fmt.Printf("%T,s40=%v,len=%d,cap=%d\n", s40, s40, len(s40), cap(s40))           // []int,s40=[0 1 2],len=3,cap=6
> 	fmt.Printf("%T,s41=%v,len=%d,cap=%d\n", s41, s41, len(s41), cap(s41))           // []int,s41=[0 1 2],len=3,cap=5
> 	fmt.Printf("%T,s40_1=%v,len=%d,cap=%d\n", s40_1, s40_1, len(s40_1), cap(s40_1)) // []int,s40_1=[0 1 2],len=3,cap=6
> 	fmt.Printf("%T,s41_1=%v,len=%d,cap=%d\n", s41_1, s41_1, len(s41_1), cap(s41_1)) // []int,s41_0=[0 1 2],len=3,cap=5
> 
> 	str2 := "我爱我的祖国"
> 	s2 := str2[0:3]
> 	s3 := str2[0:4]
> 	s4 := str2[0:5]
> 	//s5 := str2[0:3:4]                      // invalid operation: 3-index slice of string
> 	fmt.Printf("%T,str2=%v\n", str2, str2) // string,str2=我爱我的祖国
> 	fmt.Printf("%T,s2=%v\n", s2, s2)       // string,s2=我
> 	fmt.Printf("%T,s3=%v\n", s3, s3)       // string,s3=我�
> 	fmt.Printf("%T,s4=%v\n", s4, s4)       // string,s4=我��
> 
> 	s50 := "I love you!"[0:6]
> 	fmt.Printf("%T,s50=%v\n", s50, s50) // string,s50=I love
> 
> 	var nilSli []int
> 	s60 := nilSli[0:]
> 	fmt.Printf("%T,nilSli=%v,len=%d,cap=%d,%t\n", nilSli, nilSli, len(nilSli), cap(nilSli), nilSli == nil) // []int,nilSli=[],len=0,cap=0,true
> 	fmt.Printf("%T,s60=%v,len=%d,cap=%d,%t\n", s60, s60, len(s60), cap(s60), s60 == nil)                   // []int,s60=[],len=0,cap=0,true
> 
> }
> 
> ```
>
> 

#### Simple slice expressions 简单切片表达式

The primary expression

​	主表达式

```go 
a[low : high]
```

constructs a substring or slice. The [core type](https://go.dev/ref/spec#Core_types) of `a` must be a string, array, pointer to array, slice, or a [`bytestring`](https://go.dev/ref/spec#Core_types). The *indices* `low` and `high` select which elements of operand `a` appear in the result. The result has indices starting at 0 and length equal to `high` - `low`. After slicing the array `a`

构造了一个子字符串或切片。`a`的[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)必须是字符串、数组、数组指针、切片或者[bytestring](../PropertiesOfTypesAndValues#core-types-核心类型)。`low`和`high`所在的索引选择了哪些元素显示在操作数`a`的结果中。若结果的索引从0开始，则长度等于`high` 减去 `low`。在对数组`a`进行切片后

```go 
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]
```

the slice `s` has type `[]int`, length 3, capacity 4, and elements

切片`s`有类型`[]int`，长度3，容量4，以及元素

```go 
s[0] == 2
s[1] == 3
s[2] == 4
```

For convenience, any of the indices may be omitted. A missing `low` index defaults to zero; a missing `high` index defaults to the length of the sliced operand:

​	为方便起见，任何索引都可以被省略。缺少的`low`索引默认为0；缺少的`high`索引默认为被切片的操作数的长度：

```go 
a[2:]  // same as a[2 : len(a)]
a[:3]  // same as a[0 : 3]
a[:]   // same as a[0 : len(a)]
```

If `a` is a pointer to an array, `a[low : high]` is shorthand for `(*a)[low : high]`. 

​	如果`a`是一个数组指针，则`a[low:high]`是`(*a)[low:high]`的简写。

For arrays or strings, the indices are *in range* if `0` <= `low` <= `high` <= `len(a)`, otherwise they are *out of range*. For slices, the upper index bound is the slice capacity `cap(a)` rather than the length. A [constant](https://go.dev/ref/spec#Constants) index must be non-negative and [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; for arrays or constant strings, constant indices must also be in range. If both indices are constant, they must satisfy `low <= high`. If the indices are out of range at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs.

​	对于数组或字符串，如果`0 <= low <= high <= len(a)`，则索引在范围内，否则就超出了范围。对于切片，索引的上限是切片的容量`cap(a)`，而不是长度。[常量](../Constants)索引必须是非负的，并且可以用`int`类型的值来[表示](../PropertiesOfTypesAndValues#representability-可表示性)；对于数组或字符串常量，常量索引也必须在范围内。如果两个索引都是常量，它们必须满足`low <= high`。如果索引在运行时超出范围，就会发生[运行时恐慌](../Run-timePanics)。

Except for [untyped strings](https://go.dev/ref/spec#Constants), if the sliced operand is a string or slice, the result of the slice operation is a non-constant value of the same type as the operand. For untyped string operands the result is a non-constant value of type `string`. If the sliced operand is an array, it must be [addressable](https://go.dev/ref/spec#Address_operators) and the result of the slice operation is a slice with the same element type as the array.

​	除了[无类型字符串](../Constants)外：

- 如果被切片的操作数是字符串或切片，则切片的操作结果是一个与操作数相同类型的非常量值。

- 如果被切片的操作数是无类型的字符串，则切片的操作结果是一个`string`类型的非常量值。

- 如果被切片的操作数是（必须[可被寻址](#address-operators-地址运算符-地址运算符)的）数组，则切片的操作结果是一个与数组的元素类型一致的切片。

If the sliced operand of a valid slice expression is a `nil` slice, the result is a `nil` slice. Otherwise, if the result is a slice, it shares its underlying array with the operand.

​	如果有效切片表达式的切片操作数是`nil`切片，那么切片的操作结果就是一个`nil`切片。否则，如果切片的操作结果是一个切片，则它与操作数共享底层数组。

```go 
var a [10]int
s1 := a[3:7]   // s1 的底层数组是数组 a；&s1[2] == &a[5]
s2 := s1[1:4]  // s2 的底层数组是 s1 的底层数组 a； &s2[1] == &a[5]
s2[1] = 42     // s2[1] == s1[2] == a[5] == 42；它们都指向相同的底层数组元素
```

#### Full slice expressions 完整的切片表达式

The primary expression

​	主表达式

```go 
a[low : high : max]
```

constructs a slice of the same type, and with the same length and elements as the simple slice expression `a[low : high]`. Additionally, it controls the resulting slice's capacity by setting it to `max - low`. Only the first index may be omitted; it defaults to 0. The [core type](https://go.dev/ref/spec#Core_types) of `a` must be an array, pointer to array, or slice (but not a string). After slicing the array `a`

构造了一个与简单切片表达式`a[low : high]`相同类型的切片，并且具有相同的长度和元素。此外，它通过将结果切片设置为`max 减去 low`的容量。`a`的[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)必须是数组，数组指针，或者切片（但不是字符串）。在对数组`a`进行切分后

```go 
a := [5]int{1, 2, 3, 4, 5}
t := a[1:3:5]
```

the slice `t` has type `[]int`, length 2, capacity 4, and elements

切片`t`有类型`[]int`，长度2，容量4，以及元素

```go 
t[0] == 2
t[1] == 3
```

As for simple slice expressions, if `a` is a pointer to an array, `a[low : high : max]` is shorthand for `(*a)[low : high : max]`. If the sliced operand is an array, it must be [addressable](https://go.dev/ref/spec#Address_operators).

​	与简单切片表达式一样，如果`a`是一个数组指针，则`a[low:high:max]`是`(*a)[low:high:max]`的简写。如果切片的操作数是一个数组，它必须是[可被寻址的](#address-operators-地址运算符)。

> 个人注释
>
> ​	给出切片的操作数是数组，但不能被寻址的示例：
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	s := ([...]int{0, 1, 2, 3})[0:2] // invalid operation: ([...]int{…}) (value of type [4]int) (slice of unaddressable value)
> 	fmt.Println(s)
> }
> ```
>
> 

The indices are *in range* if `0 <= low <= high <= max <= cap(a)`, otherwise they are *out of range*. A [constant](https://go.dev/ref/spec#Constants) index must be non-negative and [representable](https://go.dev/ref/spec#Representability) by a value of type `int`; for arrays, constant indices must also be in range. If multiple indices are constant, the constants that are present must be in range relative to each other. If the indices are out of range at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs.

​	如果`0 <= low <= high <= max <= cap(a)`，则索引就在范围内，否则就超出了范围。[常量](../Constants)索引必须是非负数，并且可以用`int`类型的值来[表示](../PropertiesOfTypesAndValues#representability-可表示性)；对于数组，常量索引也必须在范围内。如果多个索引是常量，那么出现的常量必须在相对于彼此的范围内。如果索引在运行时超出了范围，就会发生[运行时恐慌](../Run-timePanics)。

> 个人注释
>
> ​	完整的切片表达式又称：扩展的切片表达式，是在go1.2引入的，为的是限制新生成的切片的容量。
>
> ​	`限制新生成的切片容量有什么好处？`给出示例说明下：
>
> ​	（1）从数组中生成新的切片：
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	a1 := [6]int{0, 1, 2, 3, 4, 5}
> 	a2 := [6]int{0, 1, 2, 3, 4, 5}
> 
> 	s1 := a1[0:3]
> 	s2 := a2[0:3:3]
> 	fmt.Printf("a1=%#v\n", a1)
> 	fmt.Printf("s1=%#v\n", s1)
> 	fmt.Printf("a2=%#v\n", a2)
> 	fmt.Printf("s2=%#v\n", s2)
> 
> 	fmt.Println("- apppend 22--------------------------------")
> 	s1 = append(s1, 22)
> 	s2 = append(s2, 22)
> 	fmt.Printf("a1=%#v\n", a1)
> 	fmt.Printf("s1=%#v\n", s1)
> 	fmt.Printf("a2=%#v\n", a2)
> 	fmt.Printf("s2=%#v\n", s2)
> 
> 	fmt.Println("- apppend 33--------------------------------")
> 	s1 = append(s1, 33)
> 	s2 = append(s2, 33)
> 	fmt.Printf("a1=%#v\n", a1)
> 	fmt.Printf("s1=%#v\n", s1)
> 	fmt.Printf("a2=%#v\n", a2)
> 	fmt.Printf("s2=%#v\n", s2)
> }
> 
> Output:
> 
> a1=[6]int{0, 1, 2, 3, 4, 5}
> s1=[]int{0, 1, 2}
> a2=[6]int{0, 1, 2, 3, 4, 5}
> s2=[]int{0, 1, 2}
> - apppend 22--------------------------------
> a1=[6]int{0, 1, 2, 22, 4, 5}
> s1=[]int{0, 1, 2, 22}
> a2=[6]int{0, 1, 2, 3, 4, 5}
> s2=[]int{0, 1, 2, 22}
> - apppend 33--------------------------------
> a1=[6]int{0, 1, 2, 22, 33, 5}
> s1=[]int{0, 1, 2, 22, 33}
> a2=[6]int{0, 1, 2, 3, 4, 5}
> s2=[]int{0, 1, 2, 22, 33}
> 
> ```
>
> ​	可以看到，s2在使用`arr2[0:3:3]`后，对切片s2进行的两次`append`操作，都没有对原数组中的元素造成影响！
>
> ​	（2）从切片中生成新的切片：
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	sa := []int{0, 1, 2, 3, 4, 5}
> 	sb := []int{0, 1, 2, 3, 4, 5}
> 
> 	s1 := sa[0:3]
> 	s2 := sb[0:3:3]
> 	fmt.Printf("sa=%#v\n", sa)
> 	fmt.Printf("s1=%#v\n", s1)
> 	fmt.Printf("sb=%#v\n", sb)
> 	fmt.Printf("s2=%#v\n", s2)
> 
> 	fmt.Println("- apppend 22--------------------------------")
> 	s1 = append(s1, 22)
> 	s2 = append(s2, 22)
> 	fmt.Printf("sa=%#v\n", sa)
> 	fmt.Printf("s1=%#v\n", s1)
> 	fmt.Printf("sb=%#v\n", sb)
> 	fmt.Printf("s2=%#v\n", s2)
> 
> 	fmt.Println("- apppend 33--------------------------------")
> 	s1 = append(s1, 33)
> 	s2 = append(s2, 33)
> 	fmt.Printf("sa=%#v\n", sa)
> 	fmt.Printf("s1=%#v\n", s1)
> 	fmt.Printf("sb=%#v\n", sb)
> 	fmt.Printf("s2=%#v\n", s2)
> }
> 
> Output:
> 
> sa=[]int{0, 1, 2, 3, 4, 5}
> s1=[]int{0, 1, 2}
> sb=[]int{0, 1, 2, 3, 4, 5}
> s2=[]int{0, 1, 2}
> - apppend 22--------------------------------
> sa=[]int{0, 1, 2, 22, 4, 5}
> s1=[]int{0, 1, 2, 22}
> sb=[]int{0, 1, 2, 3, 4, 5}
> s2=[]int{0, 1, 2, 22}
> - apppend 33--------------------------------
> sa=[]int{0, 1, 2, 22, 33, 5}
> s1=[]int{0, 1, 2, 22, 33}
> sb=[]int{0, 1, 2, 3, 4, 5}
> s2=[]int{0, 1, 2, 22, 33}
> ```
>
> ​	也可以看到，s2在使用`sb[0:3:3]`后，对切片s2进行的两次`append`操作，都没有对原切片中的元素造成影响！
>
> ​	这么说，只要切片的容量改变（无论是否是使用简单切片表达式，还是完整的切片表达式），就与原数组或原切片没有了联系。
>
> 

### Type assertions 类型断言

For an expression `x` of [interface type](https://go.dev/ref/spec#Interface_types), but not a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), and a type `T`, the primary expression

​	对于[接口类型](../Types#interface-ypes-接口型)但非[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)的表达式`x`和类型`T`的主表达式

```go 
x.(T)
```

asserts that `x` is not `nil` and that the value stored in `x` is of type `T`. The notation `x.(T)` is called a *type assertion*.

断言了`x`不是`nil`，并且`x`中存储的值是`T`类型。标记法`x.(T)`被称为`类型断言`。

More precisely, if `T` is not an interface type, `x.(T)` asserts that the dynamic type of `x` is [identical](https://go.dev/ref/spec#Type_identity) to the type `T`. In this case, `T` must [implement](https://go.dev/ref/spec#Method_sets) the (interface) type of `x`; otherwise the type assertion is invalid since it is not possible for `x` to store a value of type `T`. If `T` is an interface type, `x.(T)` asserts that the dynamic type of `x` [implements](https://go.dev/ref/spec#Implementing_an_interface) the interface `T`.

​	更确切地说，如果`T`不是接口类型，则`x.(T)`断言`x`的动态类型与`T`的类型[一致](../PropertiesOfTypesAndValues#type-identity-类型一致性)。在这种情况下，`T`必须实现`x`的（接口）类型；否则类型断言是无效的，因为对于`x`来说存储`T`类型的值是不可能的。如果`T`是一个接口类型，则`x.(T)` 断言`x`的动态类型[实现](../Types#implementing-an-interface-实现一个接口)了接口`T`。

> 个人注释	
>
> ```go
> package main
> 
> import "fmt"
> 
> type T1 interface{ Set() }
> 
> type St1 struct{}
> 
> func (s St1) Set() {
> 	fmt.Println("called Set")
> }
> 
> type MyInt int
> 
> func main() {
> 	// T 是接口的情况
> 
> 	// 断言无效
> 	var i1 interface{ Set() }
> 	v, ok := i1.(interface{ Set() })
> 	fmt.Printf("%T,%#v,%t\n", i1, v, ok) // <nil>,<nil>,false
> 
> 	// 断言无效
> 	var i2 T1
> 	v, ok = i2.(T1)
> 	fmt.Printf("%T,%#v,%t\n", i2, v, ok) // <nil>,<nil>,false
> 
> 	//var i3 St1
> 	//v, ok = i3.(T1) // invalid operation: i3 (variable of type St1) is not an interface
> 	//fmt.Printf("%T,%#v,%t", i3, v, ok)
> 
> 	//i4 := St1{}
> 	//v, ok = i4.(T1) // invalid operation: i4 (variable of type St1) is not an interface
> 	//fmt.Printf("%T,%#v,%t", i4, v, ok)
> 
> 	//断言： i5 的动态类型实现了 T1 接口
> 	var i5 interface{}
> 	i5 = St1{}
> 	v, ok = i5.(T1)
> 	fmt.Printf("%T,%#v,%t\n", i5, v, ok) // main.St1,main.St1{},true
> 
> 	// T 不是接口的情况
> 
> 	//i20 := MyInt(10)
> 	//v, ok = i20.(int) // invalid operation: i20 (variable of type MyInt) is not an interface
> 	//fmt.Printf("%T,%#v,%t\n", i20, v, ok)
> 
> 	//var i21 interface{}
> 	//i21 = MyInt(10)
> 	//v, ok = i21.(int) // cannot use i21.(int) (value of type int) as interface{Set()} value in assignment: int does not implement interface{Set()} (missing method Set)
> 	//fmt.Printf("%T,%#v,%t\n", i21, v, ok)
> 
> 	//i22 := St1{}
> 	//v, ok = i22.(St1) // invalid operation: i22 (variable of type St1) is not an interface
> 	//fmt.Printf("%T,%#v,%t\n", i22, v, ok)
> 
> 	//断言： i23 的动态类型与 St1 类型一致
> 	var i23 interface{}
> 	i23 = St1{}
> 	v, ok = i23.(St1)
> 	fmt.Printf("%T,%#v,%t\n", i23, v, ok) // main.St1,main.St1{},true
> 
> 	// 断言无效
> 	var i24 T1
> 	v, ok = i24.(St1)
> 	fmt.Printf("%T,%#v,%t\n", i24, v, ok) // <nil>,main.St1{},false
> 
> 	//断言： i25 的动态类型与 St1 类型一致
> 	var i25 T1
> 	i25 = St1{}
> 	v, ok = i25.(St1)
> 	fmt.Printf("%T,%#v,%t\n", i25, v, ok) // main.St1,main.St1{},true
> 
> }
> 
> ```
>
> 

If the type assertion holds, the value of the expression is the value stored in `x` and its type is `T`. If the type assertion is false, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs. In other words, even though the dynamic type of `x` is known only at run time, the type of `x.(T)` is known to be `T` in a correct program.

​	如果类型断言成立，表达式的值就是存储在`x`中的值，其类型是`T`。 如果类型断言不成立，就会发生[运行时恐慌](../Run-timePanics)。换句话说，尽管`x`的动态类型只有在运行时才知道，但在一个正确的程序中，`x.(T)`的类型是已知的，是`T`。

```go 
var x interface{} = 7          // x 有动态类型 int 以及值 7
i := x.(int)                   // i 有类型 int 以及值 7

type I interface { m() }

func f(y I) {
	s := y.(string)        // 非法的：string 没有实现 I （缺少m方法）
	r := y.(io.Reader)     // r 有类型 io.Reader，并且 y 的动态类型必须同时实现 I 和 io.Reader
	…
}
```

A type assertion used in an [assignment statement](https://go.dev/ref/spec#Assignment_statements) or initialization of the special form 

​	在[赋值语句](../Statements#assignment-statements-赋值语句)或特殊格式的初始化中使用的类型断言

```go 
v, ok = x.(T)
v, ok := x.(T)
var v, ok = x.(T)
var v, ok interface{} = x.(T) // v 的动态类型是 T， ok 的动态类型是 bool
```

yields an additional untyped boolean value. The value of `ok` is `true` if the assertion holds. Otherwise it is `false` and the value of `v` is the [zero value](https://go.dev/ref/spec#The_zero_value) for type `T`. No [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs in this case.

将产生一个额外的无类型布尔值。如果断言成立，`ok`的值为`true`。否则为`false`，并且`v`的值是`T`类型的[零值](../ProgramInitializationAndExecution#the-zero-value-零值)。在这种情况下`不会`发生[运行时恐慌](../Run-timePanics)。

### Calls 调用

Given an expression `f` with a [core type](https://go.dev/ref/spec#Core_types) `F` of [function type](https://go.dev/ref/spec#Function_types),

​	给定一个表达式`f`，其[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)为[函数类型](../Types#function-types-函数型)`F`,

```go 
f(a1, a2, … an)
```

calls `f` with arguments `a1, a2, … an`. Except for one special case, arguments must be single-valued expressions [assignable](https://go.dev/ref/spec#Assignability) to the parameter types of `F` and are evaluated before the function is called. The type of the expression is the result type of `F`. A method invocation is similar but the method itself is specified as a selector upon a value of the receiver type for the method.

带实参`a1, a2, ... an`调用了`f`。除了一种特殊情况以外，实参必须是[可分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`F`的参数类型的单值表达式，并且在函数被调用之前被求值。该表达式的类型是`F`的结果类型。方法调用是类似的，但是方法本身被指定为一个（在该方法的接收器类型的值上的）选择器。

```go 
math.Atan2(x, y)  // 函数调用
var pt *Point
pt.Scale(3.5)     // 带接收器 pt 的方法调用
```

If `f` denotes a generic function, it must be [instantiated](https://go.dev/ref/spec#Instantiations) before it can be called or used as a function value.

​	如果`f`表示一个泛型函数，在它被调用或作为函数值使用之前，必须将其[实例化](#instantiations-实例化)。

In a function call, the function value and arguments are evaluated in [the usual order](https://go.dev/ref/spec#Order_of_evaluation). After they are evaluated, the parameters of the call are passed by value to the function and the called function begins execution. The return parameters of the function are passed by value back to the caller when the function returns.

​	在函数调用中，函数值和实参以[通常的顺序](#order-of-evaluation-求值顺序)被求值。在它们被求值之后，调用的参数被按值传递给函数，然后被调用的函数开始执行。当函数返回时，函数的返回参数按值传递给调用者。

Calling a `nil` function value causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics).

​	调用一个`nil`的函数值会引起[运行时恐慌](../Run-timePanics)。

> 个人注释
>
> ​	给出一个调用`nil`函数值的示例：
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	var f func()
> 	fmt.Println(f) // <nil>
> 	
> 	f() // panic: runtime error: invalid memory address or nil pointer dereference
> }
> 
> ```
>
> 

As a special case, if the return values of a function or method `g` are equal in number and individually assignable to the parameters of another function or method `f`, then the call `f(g(*parameters_of_g*))` will invoke `f` after binding the return values of `g` to the parameters of `f` in order. The call of `f` must contain no parameters other than the call of `g`, and `g` must have at least one return value. If `f` has a final `...` parameter, it is assigned the return values of `g` that remain after assignment of regular parameters.

​	作为一个特例，如果一个函数或方法的返回值`g`在数量上相等，并且可以单独分配给另一个函数或方法`f`的参数，那么调用`f(g(parameters_of_g))`将在把`g`的返回值依次绑定到`f`的参数后再调用`f`。对`f`的调用必须不包含对`g`的调用以外的参数，并且`g`必须至少有一个返回值。如果`f`有一个在最后的`...`参数，它将被分配给`g`的（在分配完普通参数后所剩余的）返回值。

```go 
func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}

func Join(s, t string) string {
	return s + t
}

...

if Join(Split(value, len(value)/2)) != value {
	log.Panic("test fails")
}
```

A method call `x.m()` is valid if the [method set](https://go.dev/ref/spec#Method_sets) of (the type of) `x` contains `m` and the argument list can be assigned to the parameter list of `m`. If `x` is [addressable](https://go.dev/ref/spec#Address_operators) and `&x`'s method set contains `m`, `x.m()` is shorthand for `(&x).m()`: 

​	如果`x`的[方法集](../PropertiesOfTypesAndValues#method-sets-方法集)（`x`的类型）包含`m`，并且实参列表可以分配给`m`的参数列表，那么方法调用`x.m()`是有效的。如果`x`是[可寻址的](#address-operators-地址运算符)，并且`&x`的方法集包含`m`，则`x.m()`就是`(&x).m()`的简写：

```go 
var p Point
p.Scale(3.5)
```

There is no distinct method type and there are no method literals.

​	没有明确的方法类型，也`没有方法字面量`。

> 个人注释
>
> ​	怎么理解在go语言规范中的：“没有明确的方法类型，也`没有方法字面量`。” ？
>
> ​	以下是ChatGPT的回答：
>
> ​	在Go语言规范中的这句话是指在Go语言中，方法不像函数那样具有明确的类型，并且无法直接创建方法的字面量。
>
> ​	在Go语言中，方法是与特定类型关联的函数，它们通过在函数签名前面添加接收者来定义。方法只能被与其关联的类型的实例调用，因为方法的接收者隐式地传递给方法。
>
> ​	举个例子：
>
> ```go
> goCopy codetype Rectangle struct {
>     width  int
>     height int
> }
> 
> func (r Rectangle) Area() int {
>     return r.width * r.height
> }
> ```
>
> ​	在上述示例中，我们定义了一个`Rectangle`结构体类型和一个与之关联的方法`Area`。方法`Area`计算矩形的面积。
>
> ​	在Go语言规范中提到的"没有明确的方法类型"意味着方法不像函数那样具有独立的类型。例如，我们不能将一个方法赋值给一个函数变量或将其作为参数传递给其他函数。
>
> ​	另外，规范中提到的"没有方法字面量"意味着我们不能像函数那样直接创建表示方法的字面量。也就是说，我们不能像下面这样直接创建方法的值：
>
> ```go
> goCopy codefunc (r Rectangle) Area() int {
>     return r.width * r.height
> }
> 
> // 错误的方法字面量
> methodLiteral := func (r Rectangle) Area() int {
>     return r.width * r.height
> }
> ```
>
> ​	上面的代码是错误的，因为方法字面量是不被允许的。
>
> ​	总之，方法是与类型关联的函数，无法像函数那样拥有明确的类型。同时，我们也不能直接创建方法的字面量。方法只能通过与其关联的类型的实例进行调用。

### Passing arguments to `...` parameters 向...参数传递实参

If `f` is [variadic](https://go.dev/ref/spec#Function_types) with a final parameter `p` of type `...T`, then within `f` the type of `p` is equivalent to type `[]T`. If `f` is invoked with no actual arguments for `p`, the value passed to `p` is `nil`. Otherwise, the value passed is a new slice of type `[]T` with a new underlying array whose successive elements are the actual arguments, which all must be [assignable](https://go.dev/ref/spec#Assignability) to `T`. The length and capacity of the slice is therefore the number of arguments bound to `p` and may differ for each call site.

​	如果`f`是带有一个`...T`类型的位置在最后的参数`p`的可变函数，那么在`f`内部，`p`的类型等同于`[]T`类型。如果`f`被调用时没有给`p`的实参，传递给`p`的值是`nil`。否则，传递的值是一个新的`[]T`类型的切片，这个切片带有一个新的底层数组，这个底层数组的连续元素作为实参，并且这些实参都必须[可分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`T`。因此，切片的长度和容量等于绑定到`p`的实参的数量，并且对每次调用（实参数量）都可能有所不同。

Given the function and calls

​	给出函数和调用

```go 
func Greeting(prefix string, who ...string)
Greeting("nobody")
Greeting("hello:", "Joe", "Anna", "Eileen")
```

within `Greeting`, `who` will have the value `nil` in the first call, and `[]string{"Joe", "Anna", "Eileen"}` in the second.

​	在`Greeting`函数第一次被调用时，`who`的值为`nil`，在第二次被调用时，`who`的值为`[]string{"Joe", "Anna", "Eileen"}`。

If the final argument is assignable to a slice type `[]T` and is followed by `...`, it is passed unchanged as the value for a `...T` parameter. In this case no new slice is created.

​		如果最后一个实参可以分配给切片类型`[]T`，并且紧跟着`...`，那么它将按原样传递作为 `...T` 参数的值。在这种情况下，不会创建新的切片。

Given the slice `s` and call

​	给定切片`s`并调用

```go 
s := []string{"James", "Jasmine"}
Greeting("goodbye:", s...)
```

within `Greeting`, `who` will have the same value as `s` with the same underlying array.

在`Greeting`函数中，`who`将拥有与`s`相同的底层数组的值。

### Instantiations 实例化

A generic function or type is *instantiated* by substituting *type arguments* for the type parameters [[Go 1.18]({{< ref "/langSpec/Appendix#go-118">}})]. Instantiation proceeds in two steps:

​	`泛型函数`或`泛型`是通过用`类型实参`替换`类型参数`而被实例化的。实例化分两步进行：

1. Each type argument is substituted for its corresponding type parameter in the generic declaration. This substitution happens across the entire function or type declaration, including the type parameter list itself and any types in that list.
2. 在泛型声明中，每个类型参数都被替换为其对应的类型实参。这种替换发生在整个函数或类型声明中，包括类型参数列表本身和该列表中的每个类型。
3. After substitution, each type argument must [satisfy](https://go.dev/ref/spec#Satisfying_a_type_constraint) the [constraint](https://go.dev/ref/spec#Type_parameter_declarations) (instantiated, if necessary) of the corresponding type parameter. Otherwise instantiation fails.
4. 替换之后，每个类型实参必须[实现](../Types#interface-types-接口型)相应类型参数的[约束](../DeclarationsAndScope#type-constraints-类型约束)（若有需要则实例化它）。否则实例化就会失败。

Instantiating a type results in a new non-generic [named type](https://go.dev/ref/spec#Types); instantiating a function produces a new non-generic function.

​	实例化一个（泛型）类型会生成一个新的非泛型的[命名类型](../Types)；实例化一个（泛型）函数会产生一个新的非泛型的函数。

```go 
type parameter list    type arguments    after substitution
  类型参数列表             类型实参            替换后

[P any]                int               int 实现了 any
[S ~[]E, E any]        []int, int        []int 实现了 ~[]int, int 实现了 any
[P io.Writer]          string            非法的: string 没有实现 io.Writer
```

When using a generic function, type arguments may be provided explicitly, or they may be partially or completely [inferred](https://go.dev/ref/spec#Type_inference) from the context in which the function is used. Provided that they can be inferred, type argument lists may be omitted entirely if the function is:

​	在使用泛型函数时，可以显式提供类型参数，也可以从函数使用的上下文中部分或完全推断出类型参数。如果可以推断出类型参数，则如果函数是：

- [called](https://go.dev/ref/spec#Calls) with ordinary arguments,
- 使用普通实参调用，
- [assigned](https://go.dev/ref/spec#Assignment_statements) to a variable with a known type
- 分配给具有已知类型的变量
- [passed as an argument](https://go.dev/ref/spec#Calls) to another function, or
- 作为另一个函数的参数传递，或
- [returned as a result](https://go.dev/ref/spec#Return_statements).
- 返回作为结果。

In all other cases, a (possibly partial) type argument list must be present. If a type argument list is absent or partial, all missing type arguments must be inferrable from the context in which the function is used.

​	在所有其他情况下，必须存在（可能部分）类型实参列表。如果类型实参列表缺失或部分，则所有缺失的类型实参都必须可以从函数使用的上下文中推断出来。

```go
// sum returns the sum (concatenation, for strings) of its arguments.
func sum[T ~int | ~float64 | ~string](x... T) T { … }

x := sum                       // illegal: the type of x is unknown
intSum := sum[int]             // intSum has type func(x... int) int
a := intSum(2, 3)              // a has value 5 of type int
b := sum[float64](2.0, 3)      // b has value 5.0 of type float64
c := sum(b, -1)                // c has value 4.0 of type float64

type sumFunc func(x... string) string
var f sumFunc = sum            // same as var f sumFunc = sum[string]
f = sum                        // same as f = sum[string]
```

A partial type argument list cannot be empty; at least the first argument must be present. The list is a prefix of the full list of type arguments, leaving the remaining arguments to be inferred. Loosely speaking, type arguments may be omitted from "right to left".

​	（泛型函数的）部分类型实参列表不能是空的；至少第一个（类型）实参必须存在。该列表是完整的类型实参列表的前缀，剩下的实参需要被推断。宽泛地说，类型实参可以从 "从右到左"省略。

```go 
func apply[S ~[]E, E any](s S, f(E) E) S { … }

f0 := apply[]                  // 非法的：类型实参列表不能为空
f1 := apply[[]int]             // S 的类型实参被明确提供了，E 的类型实参则需要被推断
f2 := apply[[]string, string]  // 两个类型实参都被明确提供了

var bytes []byte
r := apply(bytes, func(byte) byte { … })  // 两个类型实参 都需要从函数实参中被推断出来
```

For a generic type, all type arguments must always be provided explicitly.

​	对于`泛型`，`所有的类型实参必须都被显式提供`。

### Type inference 类型推断

A use of a generic function may omit some or all type arguments if they can be *inferred* from the context within which the function is used, including the constraints of the function's type parameters. Type inference succeeds if it can infer the missing type arguments and [instantiation](https://go.dev/ref/spec#Instantiations) succeeds with the inferred type arguments. Otherwise, type inference fails and the program is invalid.

​	如果可以从函数的使用上下文推断出泛型函数的一些或所有类型实参，包括函数类型参数的约束，则可以使用泛型函数时省略这些参数。如果可以推断出缺少的类型实参并且实例化使用推断出的类型参数成功，则类型推断成功。否则，类型推断失败，程序无效。

Type inference uses the type relationships between pairs of types for inference: For instance, a function argument must be [assignable](https://go.dev/ref/spec#Assignability) to its respective function parameter; this establishes a relationship between the type of the argument and the type of the parameter. If either of these two types contains type parameters, type inference looks for the type arguments to substitute the type parameters with such that the assignability relationship is satisfied. Similarly, type inference uses the fact that a type argument must [satisfy](https://go.dev/ref/spec#Satisfying_a_type_constraint) the constraint of its respective type parameter.

​	类型推断使用类型对之间的类型关系进行推断：例如，函数实参必须可赋值给其各自的函数参数；这在实参类型和参数类型之间建立了关系。如果这两个类型中的任何一个都包含类型参数，则类型推断会查找类型实参以用满足可赋值关系的类型参数替换它们。类似地，类型推断使用类型实参必须满足其各自类型参数的约束这一事实。

Each such pair of matched types corresponds to a *type equation* containing one or multiple type parameters, from one or possibly multiple generic functions. Inferring the missing type arguments means solving the resulting set of type equations for the respective type parameters.

​	每一对匹配的类型对应一个类型方程，其中包含一个或多个类型参数，来自一个或多个泛型函数。推断缺失的类型参数意味着为各个类型参数求解所得的类型方程组。

For example, given 

​	例如，给定

```
// dedup returns a copy of the argument slice with any duplicate entries removed.
func dedup[S ~[]E, E comparable](S) S { … }

type Slice []int
var s Slice
s = dedup(s)   // same as s = dedup[Slice, int](s)
```

the variable `s` of type `Slice` must be assignable to the function parameter type `S` for the program to be valid. To reduce complexity, type inference ignores the directionality of assignments, so the type relationship between `Slice` and `S` can be expressed via the (symmetric) type equation `Slice ≡A S` (or `S ≡A Slice` for that matter), where the `A` in `≡A` indicates that the LHS and RHS types must match per assignability rules (see the section on [type unification](https://go.dev/ref/spec#Type_unification) for details). Similarly, the type parameter `S` must satisfy its constraint `~[]E`. This can be expressed as `S ≡C ~[]E` where `X ≡C Y` stands for "`X` satisfies constraint `Y`". These observations lead to a set of two equations

​	变量 `s` 的类型 `Slice` 必须可赋值给函数参数类型 `S` ，程序才有效。为了降低复杂性，类型推断忽略了赋值的方向性，因此 `Slice` 和 `S` 之间的关系可以通过（对称）类型方程 `Slice ≡A S` （或者 `S ≡A Slice` ）来表示，其中 `≡A` 中的 `A` 表示 LHS 和 RHS 类型必须根据可赋值性规则匹配（有关详细信息，请参阅类型统一部分）。同样，类型参数 `S` 必须满足其约束 `~[]E` 。这可以表示为 `S ≡C ~[]E` ，其中 `X ≡C Y` 表示“ `X` 满足约束 `Y` ”。这些观察结果产生了一组两个方程

```
	Slice ≡A S      (1)
	S     ≡C ~[]E   (2)
```

which now can be solved for the type parameters `S` and `E`. From (1) a compiler can infer that the type argument for `S` is `Slice`. Similarly, because the underlying type of `Slice` is `[]int` and `[]int` must match `[]E` of the constraint, a compiler can infer that `E` must be `int`. Thus, for these two equations, type inference infers

​	现在可以为类型参数 `S` 和 `E` 求解。从 (1) 中，编译器可以推断出 `S` 的类型参数为 `Slice` 。类似地，因为 `Slice` 的底层类型为 `[]int` ，并且 `[]int` 必须匹配约束的 `[]E` ，所以编译器可以推断出 `E` 必须为 `int` 。因此，对于这两个方程，类型推断推断出

```
	S ➞ Slice
	E ➞ int
```

Given a set of type equations, the type parameters to solve for are the type parameters of the functions that need to be instantiated and for which no explicit type arguments is provided. These type parameters are called *bound* type parameters. For instance, in the `dedup` example above, the type parameters `S` and `E` are bound to `dedup`. An argument to a generic function call may be a generic function itself. The type parameters of that function are included in the set of bound type parameters. The types of function arguments may contain type parameters from other functions (such as a generic function enclosing a function call). Those type parameters may also appear in type equations but they are not bound in that context. Type equations are always solved for the bound type parameters only.

​	给定一组类型方程，要解决的类型参数是需要实例化的函数的类型参数，并且没有为此提供显式类型参数。这些类型参数称为绑定类型参数。例如，在上面的 `dedup` 示例中，类型参数 `S` 和 `E` 绑定到 `dedup` 。对泛型函数调用的参数本身可能是一个泛型函数。该函数的类型参数包含在绑定类型参数集中。函数参数的类型可能包含来自其他函数的类型参数（例如包含函数调用的泛型函数）。这些类型参数也可能出现在类型方程中，但它们在此上下文中没有绑定。类型方程始终仅针对绑定类型参数求解。

Type inference supports calls of generic functions and assignments of generic functions to (explicitly function-typed) variables. This includes passing generic functions as arguments to other (possibly also generic) functions, and returning generic functions as results. Type inference operates on a set of equations specific to each of these cases. The equations are as follows (type argument lists are omitted for clarity):

​	类型推断支持调用泛型函数并将泛型函数赋值给（显式函数类型的）变量。这包括将泛型函数作为参数传递给其他（可能也是泛型）函数，并将泛型函数作为结果返回。类型推断针对每种情况的一组特定方程进行操作。方程式如下（为清楚起见，省略了类型参数列表）：

- For a function call `f(a0, a1, …)` where `f` or a function argument `ai` is a generic function:

- 对于函数调用 `f(a0, a1, …)` ，其中 `f` 或函数参数 `ai` 是泛型函数：
  Each pair `(ai, pi)` of corresponding function arguments and parameters where `ai` is not an [untyped constant](https://go.dev/ref/spec#Constants) yields an equation `typeof(pi) ≡A typeof(ai)`. 

  每个函数实参和参数的对应对 `(ai, pi)` （其中 `ai` 不是未类型化的常量）生成一个方程 `typeof(pi) ≡A typeof(ai)` 。
  If `ai` is an untyped constant `cj`, and `typeof(pi)` is a bound type parameter `Pk`, the pair `(cj, Pk)` is collected separately from the type equations. 

  如果 `ai` 是未类型化的常量 `cj` ，并且 `typeof(pi)` 是绑定类型参数 `Pk` ，则对 `(cj, Pk)` 从类型方程中单独收集。

- For an assignment `v = f` of a generic function `f` to a (non-generic) variable `v` of function type:

- 对于将泛型函数 `f` 赋值给函数类型（非泛型）变量 `v` ：
  `typeof(v) ≡A typeof(f)`.

- For a return statement `return …, f, … `where `f` is a generic function returned as a result to a (non-generic) result variable `r` of function type:

- 对于返回语句 `return …, f, … `（其中 `f` 是作为结果返回给函数类型（非泛型）结果变量 `r` 的泛型函数）：
  `typeof(r) ≡A typeof(f)`.

Additionally, each type parameter `Pk` and corresponding type constraint `Ck` yields the type equation `Pk ≡C Ck`.

​	此外，每个类型参数 `Pk` 和相应的类型约束 `Ck` 生成类型方程 `Pk ≡C Ck` 。

Type inference gives precedence to type information obtained from typed operands before considering untyped constants. Therefore, inference proceeds in two phases:

​	类型推断优先考虑从类型化操作数获得的类型信息，然后再考虑非类型化常量。因此，推断分两个阶段进行：

1. The type equations are solved for the bound type parameters using [type unification](https://go.dev/ref/spec#Type_unification). If unification fails, type inference fails.
2. 使用类型统一来解决绑定类型参数的类型方程。如果统一失败，则类型推断失败。
3. For each bound type parameter `Pk` for which no type argument has been inferred yet and for which one or more pairs `(cj, Pk)` with that same type parameter were collected, determine the [constant kind](https://go.dev/ref/spec#Constant_expressions) of the constants `cj` in all those pairs the same way as for [constant expressions](https://go.dev/ref/spec#Constant_expressions). The type argument for `Pk` is the [default type](https://go.dev/ref/spec#Constants) for the determined constant kind. If a constant kind cannot be determined due to conflicting constant kinds, type inference fails.
4. 对于尚未推断出类型参数的每个绑定类型参数 `Pk` ，以及为此收集了一个或多个具有相同类型参数的配对 `(cj, Pk)` ，以与常量表达式相同的方式确定所有这些配对中常量 `cj` 的常量类型。 `Pk` 的类型参数是确定常量类型的默认类型。如果由于常量类型冲突而无法确定常量类型，则类型推断失败。

If not all type arguments have been found after these two phases, type inference fails.

​	如果在完成这两个阶段后仍未找到所有类型实参，则类型推断失败。

If the two phases are successful, type inference determined a type argument for each bound type parameter:

​	如果这两个阶段都成功，类型推断将为每个受限类型参数确定一个类型参数：

```
	Pk ➞ Ak
```

A type argument `Ak` may be a composite type, containing other bound type parameters `Pk` as element types (or even be just another bound type parameter). In a process of repeated simplification, the bound type parameters in each type argument are substituted with the respective type arguments for those type parameters until each type argument is free of bound type parameters.

​	类型实参 `Ak` 可以是复合类型，包含其他绑定类型参数 `Pk` 作为元素类型（甚至可以只是另一个绑定类型参数）。在重复简化的过程中，每个类型参数中的绑定类型参数都用这些类型参数的相应类型参数替换，直到每个类型参数都不包含绑定类型参数。

If type arguments contain cyclic references to themselves through bound type parameters, simplification and thus type inference fails. Otherwise, type inference succeeds.

​	如果类型实参通过绑定类型参数包含对自身的循环引用，则简化和类型推断将失败。否则，类型推断将成功。

​	

#### Type unification 类型联合

Type inference solves type equations through *type unification*. Type unification recursively compares the LHS and RHS types of an equation, where either or both types may be or contain bound type parameters, and looks for type arguments for those type parameters such that the LHS and RHS match (become identical or assignment-compatible, depending on context). To that effect, type inference maintains a map of bound type parameters to inferred type arguments; this map is consulted and updated during type unification. Initially, the bound type parameters are known but the map is empty. During type unification, if a new type argument `A` is inferred, the respective mapping `P ➞ A` from type parameter to argument is added to the map. Conversely, when comparing types, a known type argument (a type argument for which a map entry already exists) takes the place of its corresponding type parameter. As type inference progresses, the map is populated more and more until all equations have been considered, or until unification fails. Type inference succeeds if no unification step fails and the map has an entry for each type parameter.

​	类型推断通过类型统一来解决类型方程。类型统一递归地比较方程的 LHS 和 RHS 类型，其中任一类型或同时两种类型可能为绑定类型参数或包含绑定类型参数，并查找这些类型参数的类型参数，以便 LHS 和 RHS 匹配（根据上下文，变为相同或可赋值兼容）。为此，类型推断维护一个从绑定类型参数到推断类型参数的映射；在类型统一期间查阅并更新此映射。最初，绑定类型参数是已知的，但映射是空的。在类型统一期间，如果推断出新的类型参数 `A` ，则将从类型参数到参数的相应映射 `P ➞ A` 添加到映射中。相反，在比较类型时，已知类型参数（已存在映射项的类型参数）取代其对应的类型参数。随着类型推断的进行，映射会逐渐填充，直到考虑所有方程或直到统一失败。如果没有任何统一步骤失败并且映射为每个类型参数都有一个条目，则类型推断成功。

For example, given the type equation with the bound type parameter `P`

​	例如，给定具有绑定类型参数 `P` 的类型方程

```
	[10]struct{ elem P, list []P } ≡A [10]struct{ elem string; list []string }
```

type inference starts with an empty map. Unification first compares the top-level structure of the LHS and RHS types. Both are arrays of the same length; they unify if the element types unify. Both element types are structs; they unify if they have the same number of fields with the same names and if the field types unify. The type argument for `P` is not known yet (there is no map entry), so unifying `P` with `string` adds the mapping `P ➞ string` to the map. Unifying the types of the `list` field requires unifying `[]P` and `[]string` and thus `P` and `string`. Since the type argument for `P` is known at this point (there is a map entry for `P`), its type argument `string` takes the place of `P`. And since `string` is identical to `string`, this unification step succeeds as well. Unification of the LHS and RHS of the equation is now finished. Type inference succeeds because there is only one type equation, no unification step failed, and the map is fully populated.

​	类型推断从一个空映射开始。统一首先比较 LHS 和 RHS 类型的顶级结构。两者都是长度相同的数组；如果元素类型统一，它们就会统一。两种元素类型都是结构；如果它们具有相同数量的具有相同名称的字段，并且字段类型统一，则它们会统一。对于 `P` 的类型参数尚不知道（没有映射条目），因此将 `P` 与 `string` 统一会将映射 `P ➞ string` 添加到映射中。统一 `list` 字段的类型需要统一 `[]P` 和 `[]string` ，从而统一 `P` 和 `string` 。由于此时已知 `P` 的类型参数（ `P` 有一个映射条目），因此其类型参数 `string` 取代了 `P` 。并且由于 `string` 与 `string` 相同，因此此统一步骤也成功。现在，方程的 LHS 和 RHS 的统一已完成。类型推断成功，因为只有一个类型方程，没有统一步骤失败，并且映射已完全填充。

Unification uses a combination of *exact* and *loose* unification depending on whether two types have to be [identical](https://go.dev/ref/spec#Type_identity), [assignment-compatible](https://go.dev/ref/spec#Assignability), or only structurally equal. The respective [type unification rules](https://go.dev/ref/spec#Type_unification_rules) are spelled out in detail in the [Appendix](https://go.dev/ref/spec#Appendix).

​	统一使用精确和松散统一的组合，具体取决于两种类型是否必须相同、是否兼容赋值或仅在结构上相等。相应的类型统一规则在附录中详细说明。

For an equation of the form `X ≡A Y`, where `X` and `Y` are types involved in an assignment (including parameter passing and return statements), the top-level type structures may unify loosely but element types must unify exactly, matching the rules for assignments.

​	对于形式为 `X ≡A Y` 的方程，其中 `X` 和 `Y` 是参与赋值的类型（包括参数传递和返回语句），顶级类型结构可能松散地统一，但元素类型必须完全统一，与赋值规则相匹配。

For an equation of the form `P ≡C C`, where `P` is a type parameter and `C` its corresponding constraint, the unification rules are bit more complicated:

​	对于形式为 `P ≡C C` 的方程，其中 `P` 是类型参数， `C` 是其对应的约束，统一规则要复杂一些：

- If `C` has a [core type](https://go.dev/ref/spec#Core_types) `core(C)` and `P` has a known type argument `A`, `core(C)` and `A` must unify loosely. If `P` does not have a known type argument and `C` contains exactly one type term `T` that is not an underlying (tilde) type, unification adds the mapping `P ➞ T` to the map.
- 如果 `C` 具有核心类型 `core(C)` ， `P` 具有已知类型参数 `A` ，则 `core(C)` 和 `A` 必须松散地统一。如果 `P` 没有已知类型参数，并且 `C` 恰好包含一个不是底层（波浪号）类型的类型项 `T` ，则统一会将映射 `P ➞ T` 添加到映射中。
- If `C` does not have a core type and `P` has a known type argument `A`, `A` must have all methods of `C`, if any, and corresponding method types must unify exactly.
- 如果 `C` 没有核心类型， `P` 具有已知类型参数 `A` ，则 `A` 必须具有 `C` 的所有方法（如果有），并且相应的方法类型必须完全统一。

When solving type equations from type constraints, solving one equation may infer additional type arguments, which in turn may enable solving other equations that depend on those type arguments. Type inference repeats type unification as long as new type arguments are inferred.

​	在从类型约束中求解类型方程时，求解一个方程可能会推断出其他类型参数，进而可能能够求解依赖于这些类型参数的其他方程。只要推断出新的类型参数，类型推断就会重复类型统一。

### Operators 操作符

Operators combine operands into expressions.

​	操作符将操作数组合成表达式。

```
Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
```

Comparisons are discussed [elsewhere](https://go.dev/ref/spec#Comparison_operators). For other binary operators, the operand types must be [identical](https://go.dev/ref/spec#Type_identity) unless the operation involves shifts or untyped [constants](https://go.dev/ref/spec#Constants). For operations involving constants only, see the section on [constant expressions](https://go.dev/ref/spec#Constant_expressions).

​	**比较操作符**将在[其他地方](#comparison-operators-比较运算符)讨论。对于其他二元运算符，除非操作涉及移位或无类型的[常量](../Constants)，操作数类型必须是[一致的](../PropertiesOfTypesAndValues#type-identity-类型一致性)。对于只涉及常量的操作，请参见[常量表达式](#constant-expressions-常量表达式)部分。

Except for shift operations, if one operand is an untyped [constant](https://go.dev/ref/spec#Constants) and the other operand is not, the constant is implicitly [converted](https://go.dev/ref/spec#Conversions) to the type of the other operand.

​	除了**移位操作符**之外，如果一个操作数是`无类型`常量，而另一个操作数不是，那么该常量将被隐式地[转换](#conversions-转换)为另一个操作数的类型。

The right operand in a shift expression must have [integer type](https://go.dev/ref/spec#Numeric_types) [[Go 1.13]({{< ref "/langSpec/Appendix#go-113">}})] or be an untyped constant [representable](https://go.dev/ref/spec#Representability) by a value of type `uint`. If the left operand of a non-constant shift expression is an untyped constant, it is first implicitly converted to the type it would assume if the shift expression were replaced by its left operand alone.

​	移位表达式中的`右操作数`必须是[整数类型](../Types#numeric-types-数值型)，或者是可以用`uint`类型的值[表示](../PropertiesOfTypesAndValues#representability-可表示性)的`无类型`常量。如果一个非常量移位表达式的`左操作数`是一个`无类型`常量，那么它首先被隐式地转换为假设移位表达式被其左操作数单独替换时的类型。

```go 
var a [1024]byte
var s uint = 33

// 以下示例的结果是针对 64-bits ints给出的。
var i = 1<<s  // 1 has type int => 1 拥有 int 类型，i == 8589934592
var j int32 = 1<<s // 1 拥有 int32 类型；j == 0
var k = uint64(1<<s) // 1 拥有 uint64 类型；k == 1 << 33, k == 8589934592
var m int = 1.0<<s // 1.0 拥有 int 类型；m == 1 << 33, m == 8589934592 
var n = 1.0<<s == j  // 1.0 拥有 int32 类型；n == true
var o = 1<<s == 2<<s // 1 和 2 拥有 int 类型；o == false
var p = 1<<s == 1<<33 // 1 拥有 int 类型；p == true
var u = 1.0<<s        // 非法的：1.0 拥有 float64 类型，不能移位
var u1 = 1.0<<s != 0  // 非法的： 1.0 拥有 float64 类型，不能移位
var u2 = 1<<s != 1.0  // 非法的： 1 拥有 float64 类型，不能移位
var v1 float32 = 1<<s // 非法的：1 拥有 float32 类型，不能移位
var v2 = string(1<<s) // 非法的： 1 被转换成 string 类型，不能移位
var w int64 = 1.0<<33 // 1.0 << 33 是一个常量移位表达式；w == 1 << 33, w == 8589934592
var x = a[1.0<<s] // panics：1.0 拥有 int 类型，但 1 << 33 溢出了数组的边界
var b = make([]byte, 1.0<<s)   // 1.0 拥有 int 类型；len(b) == 1 << 33

// 以下示例的结果是针对 32-bits ints给出。 这意味着移位将会溢出。
var mm int = 1.0<<s  // 1.0 拥有 int 类型；mm == 0
var oo = 1<<s == 2<<s  // 1 和 2 拥有 int 类型； oo == true
var pp = 1<<s == 1<<33 // 非法的：1 拥有 int 类型，但 1 << 33 溢出了 int 的范围
var xx = a[1.0<<s]  // 1.0 拥有 int 类型；xx == a[0] 
var bb = make([]byte, 1.0<<s)  // 1.0 拥有 int 类型；len(bb) == 0
```



#### Operator precedence 优先级运算符

Unary operators have the highest precedence. As the `++` and `--` operators form statements, not expressions, they fall outside the operator hierarchy. As a consequence, statement `*p++` is the same as `(*p)++`.

​	一元运算符的优先级最高。由于`++`和`--`运算符构成的是语句，而不是表达式，因此它们不属于运算符等级体系。因此，语句`*p++`与`(*p)++`相同。

There are five precedence levels for binary operators. Multiplication operators bind strongest, followed by addition operators, comparison operators, `&&` (logical AND), and finally `||` (logical OR):

​	二元运算符有五个优先级别。`乘法运算符`绑定最强，其次是`加法运算符`、比较运算符、`&&`（逻辑AND），最后是`||`（逻辑OR）。

```go 
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
```

Binary operators of the same precedence associate from left to right. For instance, `x / y * z` is the same as `(x / y) * z`.

​	优先级相同的二元运算符按从左到右的顺序结合。例如，`x / y * z`等同于`(x / y)* z`。

```go 
+x
23 + 3*x[i]
x <= f()
^a >> b
f() || g()
x == y+1 && <-chanInt > 0
```

### Arithmetic operators 算术运算符

Arithmetic operators apply to numeric values and yield a result of the same type as the first operand. The four standard arithmetic operators (`+`, `-`, `*`, `/`) apply to [integer](https://go.dev/ref/spec#Numeric_types), [floating-point](https://go.dev/ref/spec#Numeric_types), and [complex](https://go.dev/ref/spec#Numeric_types) types; `+` also applies to [strings](https://go.dev/ref/spec#String_types). The bitwise logical and shift operators apply to integers only.

​	算术运算符适用于数字值，产生的结果`与第一个操作数的类型`相同。四个标准的算术运算符（`+`、`-`、`*`、`/`）适用于[整型](../Types#numeric-types-数值型)、[浮点型](../Types#numeric-types-数值型)和[复数型](../Types#numeric-types-数值型)；`+`也适用于[字符串](../Types#string-types-字符串型)。位逻辑运算符和移位运算符只适用于整型。

```go
+    sum        （和）            integers, floats, complex values, strings
-    difference （差）            integers, floats, complex values
*    product    （积）            integers, floats, complex values
/    quotient   （商）            integers, floats, complex values
%    remainder  （余）            integers

&    bitwise AND                 integers
|    bitwise OR                  integers
^    bitwise XOR                 integers
&^   bit clear (AND NOT)         integers

<<   left shift                  integer << integer >= 0
>>   right shift                 integer >> integer >= 0
```

If the operand type is a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), the operator must apply to each type in that type set. The operands are represented as values of the type argument that the type parameter is [instantiated](https://go.dev/ref/spec#Instantiations) with, and the operation is computed with the precision of that type argument. For example, given the function:

​	如果操作数类型是[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)，那么操作数必须适用于该类型集中的每个类型。操作数被表示为类型参数被[实例化](#instantiations-实例化)的类型实参的值，并且操作以该类型实参的精度进行计算。例如，给定一个函数：

```go 
func dotProduct[F ~float32|~float64](v1, v2 []F) F {
	var s F
	for i, x := range v1 {
		y := v2[i]
		s += x * y
	}
	return s
}
```

the product `x * y` and the addition `s += x * y` are computed with `float32` or `float64` precision, respectively, depending on the type argument for `F`.

`x * y`的乘积 和 `s += x * y`的加法分别以`float32`或`float64`精度计算，这取决于`F`的类型实参。

#### Integer operators 整数运算符

For two integer values `x` and `y`, the integer quotient `q = x / y` and remainder `r = x % y` satisfy the following relationships:

​	对于两个整数值`x`和`y`，它们的整数商`q = x / y`和余数`r = x % y`满足以下关系：

```
x = q*y + r  and  |r| < |y|
```

with `x / y` truncated towards zero (["truncated division"](https://en.wikipedia.org/wiki/Modulo_operation)).

with `x / y`被截断到零（"[截断除法](https://en.wikipedia.org/wiki/Modulo_operation)"）。

```go
 x     y     x / y     x % y
 5     3       1         2
-5     3      -1        -2
 5    -3      -1         2
-5    -3       1        -2
```

The one exception to this rule is that if the dividend `x` is the most negative value for the int type of `x`, the quotient `q = x / -1` is equal to `x` (and `r = 0`) due to two's-complement [integer overflow](https://go.dev/ref/spec#Integer_overflow):

​	这条规则有一个例外：如果）`被除数`（dividend`x`是`x`的`int`类型的最负值，那么商`q = x / -1`就等于`x`（而`r = 0`），这是由于二元补码的[整数溢出](#integer-overflow-整数溢出)：

```go
                         x, q
int8                     -128
int16                  -32768
int32             -2147483648
int64    -9223372036854775808
```

If the divisor is a [constant](https://go.dev/ref/spec#Constants), it must not be zero. If the divisor is zero at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs. If the dividend is non-negative and the divisor is a constant power of 2, the division may be replaced by a right shift, and computing the remainder may be replaced by a bitwise AND operation:

​	如果`除数`（divisor）是一个[常量](../Constants)，那么它一定不能为零。如果除数在运行时为零，就会发生[运行时恐慌](../Run-timePanics)。如果除数是非负数，并且除数是2的常数幂，除法可以用`右移`来代替，计算余数可以用`按位与`操作来代替。

```go
 x     x / 4     x % 4     x >> 2     x & 3
 11      2         3         2          3
-11     -2        -3        -3          1
```

The shift operators shift the left operand by the shift count specified by the right operand, which must be non-negative. If the shift count is negative at run time, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs. The shift operators implement arithmetic shifts if the left operand is a signed integer and logical shifts if it is an unsigned integer. There is no upper limit on the shift count. Shifts behave as if the left operand is shifted `n` times by 1 for a shift count of `n`. As a result, `x << 1` is the same as `x*2` and `x >> 1` is the same as `x/2` but truncated towards negative infinity.

​	移位运算符通过右操作数指定的`移位计数`对左操作数进行移位，`移位计数`必须为非负数。如果`移位计数`在运行时为负数，就会发生[运行时恐慌](../Run-timePanics)。如果`左操作数`是`有符号整数`，移位操作符实现`算术移位`；如果是`无符号整数`，则实现`逻辑移位`。`移位计数`没有上限。`移位计数`为`n`的移位行为就像左操作数被`1`移了`n`次。因此，`x<<1`与`x*2`相同，`x>>1`与`x/2`相同（但向右移位被截断到负无穷大）。

For integer operands, the unary operators `+`, `-`, and `^` are defined as follows:

​	对于整数操作数，一元运算符`+`、`-`和`^`的定义如下：

```go
+x                          is 0 + x
-x    negation              is 0 - x
^x    bitwise complement    is m ^ x  with m = "all bits set to 1" for unsigned x
                                      and  m = -1 for signed x
```

#### Integer overflow 整数溢出

For [unsigned integer](https://go.dev/ref/spec#Numeric_types) values, the operations `+`, `-`, `*`, and `<<` are computed modulo \\(2^n\\), where *n* is the bit width of the unsigned integer's type. Loosely speaking, these unsigned integer operations discard high bits upon overflow, and programs may rely on "wrap around".

​	对于[无符号整型值](../Types#numeric-types-数值型)，`+`、`-`、`*`和`<<`运算是以 \\(2^n\\)为模来计算的，其中`n`是无符号整型的位宽。广义上讲，这些无符号整型操作`在溢出时丢弃高位`，程序可以依靠 "`wrap around`"。

For signed integers, the operations `+`, `-`, `*`, `/`, and `<<` may legally overflow and the resulting value exists and is deterministically defined by the signed integer representation, the operation, and its operands. Overflow does not cause a [run-time panic](https://go.dev/ref/spec#Run_time_panics). A compiler may not optimize code under the assumption that overflow does not occur. For instance, it may not assume that `x < x + 1` is always true.

​	对于`有符号整型值`，`+`、`-`、`*`、`/`和`<<`运算`可以合法地溢出`，其产生的值是存在的，并且可以被有符号整型表示法、其操作和操作数明确地定义。溢出不会引起[运行时恐慌](../Run-timePanics)。在假设不发生溢出的情况下，编译器可能不会优化代码。例如，它不会假设`x<x+1`总是真的。

#### Floating-point operators 浮点运算符

For floating-point and complex numbers, `+x` is the same as `x`, while `-x` is the negation of `x`. The result of a floating-point or complex division by zero is not specified beyond the IEEE-754 standard; whether a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs is implementation-specific.

​	对于浮点数和复数，`+x`与`x`相同，而`-x`是负的`x`。浮点数或复数除以0的结果，在IEEE-754标准中没有规定；是否会发生[运行时恐慌](../Run-timePanics)是由具体实现决定的。

An implementation may combine multiple floating-point operations into a single fused operation, possibly across statements, and produce a result that differs from the value obtained by executing and rounding the instructions individually. An explicit [floating-point type](https://go.dev/ref/spec#Numeric_types) [conversion](https://go.dev/ref/spec#Conversions) rounds to the precision of the target type, preventing fusion that would discard that rounding.

​	某些实现可能会将多个浮点运算合并为一个单一的融合运算，可能会跨越语句，产生的结果与单独执行和舍入指令得到的值不同。显式的[浮点类型转换](#conversions)是按照目标类型的精度进行舍入的，这样就可以避免融合时放弃舍入的做法。=> 仍有疑问？？

For instance, some architectures provide a "fused multiply and add" (FMA) instruction that computes `x*y + z` without rounding the intermediate result `x*y`. These examples show when a Go implementation can use that instruction:

​	例如，一些体系架构提供了一个 "`fused multiply and add`"（`FMA`）指令，其在计算`x*y+z`时，不对中间结果`x*y`进行舍入。这些例子显示了Go的实现何时可以使用该指令：

```go 
// FMA allowed for computing r, because x*y is not explicitly rounded: 
// => FMA 允许被用来计算 r, 因为 x*y 不会被显式地进行舍入：
r  = x*y + z
r  = z;   r += x*y
t  = x*y; r = t + z
*p = x*y; r = *p + z
r  = x*y + float64(z)

// FMA disallowed for computing r, because it would omit rounding of x*y:
// => FMA 不允许被用来计算 r, 因为它会省略 x*y 的舍入：
r  = float64(x*y) + z
r  = z; r += float64(x*y)
t  = float64(x*y); r = t + z
```

#### String concatenation 字符串连接

Strings can be concatenated using the `+` operator or the `+=` assignment operator:

​	字符串可以使用`+`运算符或`+=`赋值运算符进行连接：

```go
s := "hi" + string(c)
s += " and good bye"
```

String addition creates a new string by concatenating the operands.

​	字符串加法通过连接操作数创建一个新的字符串。

### Comparison operators 比较运算符

Comparison operators compare two operands and yield an untyped boolean value.

​	**比较运算符**比较两个操作数，并产生一个无类型布尔值。

```
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

In any comparison, the first operand must be [assignable](https://go.dev/ref/spec#Assignability) to the type of the second operand, or vice versa.

​	在任何比较中，第一个操作数必须是[可分配](../PropertiesOfTypesAndValues#assignability-可分配性)给第二个操作数的类型，反之亦然。

The equality operators `==` and `!=` apply to operands of *comparable* types. The ordering operators `<`, `<=`, `>`, and `>=` apply to operands of *ordered* types. These terms and the result of the comparisons are defined as follows:

​	**相等运算符**`==`和`!=`适用于可比较的操作数。**排序运算符**`<`, `<=`, `>`, 和`>=`适用于被排序的操作数。这些术语和比较结果的定义如下：

- Boolean types are comparable. Two boolean values are equal if they are either both `true` or both `false`.
- 布尔值是可比较的。如果两个布尔值都是`true`或者都是`false`，那么它们是相等的。
- Integer types are comparable and ordered. Two integer values are compared in the usual way.
- 按照通常的方式，整数值是可比较的并且是可排序的。
- Floating-point types are comparable and ordered. Two floating-point values are compared as defined by the IEEE-754 standard.
- 按照**IEEE-754标准**的定义，浮点值是可比较的并且是可排序的。
- Complex types are comparable. Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.
- 复数值是可比较的。如果`real(u) == real(v)`和`imag(u) == imag(v)`，则这两个复数值`u`和`v`是相等的。
- String types are comparable and ordered. Two string values are compared lexically byte-wise.
- 字符串类型值是可比较和有序的。两个字符串值按字节顺序进行比较。
- Pointer types are comparable. Two pointer values are equal if they point to the same variable or if both have value `nil`. Pointers to distinct [zero-size](https://go.dev/ref/spec#Size_and_alignment_guarantees) variables may or may not be equal.
- 指针值是可比较的。如果两个指针值指向同一个变量，或者两个指针值都是`nil`，则它们的值是相等的。指向不同的[零尺寸](../SystemConsiderations#size-and-alignment-guarantees-大小和对齐保证)变量的指针值可能相等，也可能不相等。
- Channel types are comparable. Two channel values are equal if they were created by the same call to [`make`](https://go.dev/ref/spec#Making_slices_maps_and_channels) or if both have value `nil`.
- 通道值是可比较的。如果两个通道值是由相同的 [make](../Built-inFunctions#making-slices-maps-and-channels-制作切片映射和通道)调用创建的，或者它们的值都为`nil`，则它们的值是相等的。
- Interface types that are not type parameters are comparable. Two interface values are equal if they have [identical](https://go.dev/ref/spec#Type_identity) dynamic types and equal dynamic values or if both have value `nil`.
- 接口值是可比较的。如果两个接口值有[一致的](../PropertiesOfTypesAndValues#type-identity-类型一致性)动态类型和相同的动态值，或者两者的值都是`nil`，则它们的值是相等的。
- A value `x` of non-interface type `X` and a value `t` of interface type `T` can be compared if type `X` is comparable and `X` [implements](https://go.dev/ref/spec#Implementing_an_interface) `T`. They are equal if `t`'s dynamic type is identical to `X` and `t`'s dynamic value is equal to `x`.
- 非接口类型 `X` 的值 `x` 和接口类型 `T` 的值 `t` ，在 `X` 类型的值是可比较的并且 `X` [实现](../Types#implementing-an-interface-实现一个接口) `T` 时是可比较的。如果 `t` 的动态类型等于 `X`，且 `t` 的动态值等于 `x`，则它们是相等的。
- Struct types are comparable if all their field types are comparable. Two struct values are equal if their corresponding non-[blank](https://go.dev/ref/spec#Blank_identifier) field values are equal. The fields are compared in source order, and comparison stops as soon as two field values differ (or all fields have been compared).
- 如果结构体类型的所有字段类型都是可比较的，那么结构体类型的值是可比较的。如果两个结构体对应的非空白字段值相等，那么这两个结构体值相等。字段按源代码顺序进行比较，并在两个字段值不同时停止比较（或已比较完所有字段）。
- Array types are comparable if their array element types are comparable. Two array values are equal if their corresponding element values are equal. The elements are compared in ascending index order, and comparison stops as soon as two element values differ (or all elements have been compared).
- 如果它们的数组元素类型是可比较的，那么数组类型的值是可比较的。如果它们对应的元素值相等，两个数组值相等。元素按升序索引进行比较，并在两个元素值不同时停止比较（或已比较完所有元素）。
- Type parameters are comparable if they are strictly comparable (see below).
- 如果类型参数是严格可比较的（见下文），则类型参数的值是可比较的。 

​	A comparison of two interface values with identical dynamic types causes a [run-time panic](https://go.dev/ref/spec#Run_time_panics) if that type is not comparable. This behavior applies not only to direct interface value comparisons but also when comparing arrays of interface values or structs with interface-valued fields.

​	对两个动态类型相同的接口值进行比较，如果它们的类型值不具有可比性，则会引起[运行时恐慌](../Run-timePanics)。这种行为不仅适用于直接的接口值比较，也适用于比较接口值的数组或带有接口类型字段的结构体。

Slice, map, and function types are not comparable. However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`. Comparison of pointer, channel, and interface values to `nil` is also allowed and follows from the general rules above.

​	`切片值、映射值和函数值是不可比较的`。然而，作为一种特殊情况，切片值、映射值或函数值可以与预先声明的标识符`nil`比较。指针值、通道值和接口值与`nil`的比较也是允许的，并遵循上述的通用规则。

```go 
const c = 3 < 4 // c 是无类型的布尔常量 true

type MyBool bool
var x, y int
var (
	// 比较的结果为一个无类型的布尔值。
	// 使用通用赋值规则。
	b3        = x == y // b3 拥有 bool 类型
	b4 bool   = x == y // b4 拥有 bool 类型
	b5 MyBool = x == y // b5 拥有 MyBool 类型
)
```

A type is *strictly comparable* if it is comparable and not an interface type nor composed of interface types. Specifically:

​	如果一个类型可比较，并且不是接口类型，也不是由接口类型组成，则该类型是严格可比较的。具体来说：

- Boolean, numeric, string, pointer, and channel types are strictly comparable.
- Struct types are strictly comparable if all their field types are strictly comparable.
- Array types are strictly comparable if their array element types are strictly comparable.
- Type parameters are strictly comparable if all types in their type set are strictly comparable.

### Logical operators 逻辑运算符

Logical operators apply to [boolean](https://go.dev/ref/spec#Boolean_types) values and yield a result of the same type as the operands. The left operand is evaluated, and then the right if the condition requires it. 

​	逻辑运算符适用于[布尔](../Types#boolean-types-布尔型)值，并产生一个与操作数相同类型的结果。右操作数是按条件进行求值的。

```go
&&    conditional AND    p && q  is  "if p then q else false"
||    conditional OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"
```

### Address operators 地址运算符

For an operand `x` of type `T`, the address operation `&x` generates a pointer of type `*T` to `x`. The operand must be *addressable*, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array. As an exception to the addressability requirement, `x` may also be a (possibly parenthesized) [composite literal](https://go.dev/ref/spec#Composite_literals). If the evaluation of `x` would cause a [run-time panic](https://go.dev/ref/spec#Run_time_panics), then the evaluation of `&x` does too. 

​	对于类型为`T`的操作数`x`，寻址操作`&x`生成一个类型为`*T`的指针指向`x`。该操作数`x`必须是可寻址的，也就是说，它要么是一个变量、指针间接引用（pointer indirection）或`对切片的索引操作（slice indexing operation，是一个名词）`；要么是一个可寻址结构体操作数的字段选择器；要么是一个可寻址数组的数组索引操作。作为可寻址要求的一个例外，`x`也可以是一个（可能带括号的）[复合字面量](#composite-literals-复合字面量)。如果对`x`的求值会引起[运行时恐慌](../Run-timePanics)，那么对`&x`的求值也会引起[运行时恐慌](../Run-timePanics)。

For an operand `x` of pointer type `*T`, the pointer indirection `*x` denotes the [variable](https://go.dev/ref/spec#Variables) of type `T` pointed to by `x`. If `x` is `nil`, an attempt to evaluate `*x` will cause a [run-time panic](https://go.dev/ref/spec#Run_time_panics). 

​	对于指针类型`*T`的操作数`x`，指针间接引用`*x`表示指向`x`的类型`T`的[变量]()，如果`x`是`nil`，试图求值`*x`将导致[运行时恐慌](../Run-timePanics)。

```go 
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // 导致一个 run-time panic
&*x  // 导致一个 run-time panic
```

### Receive operator 接收操作符

For an operand `ch` whose [core type](https://go.dev/ref/spec#Core_types) is a [channel](https://go.dev/ref/spec#Channel_types), the value of the receive operation `<-ch` is the value received from the channel `ch`. The channel direction must permit receive operations, and the type of the receive operation is the element type of the channel. The expression blocks until a value is available. Receiving from a `nil` channel blocks forever. A receive operation on a [closed](https://go.dev/ref/spec#Close) channel can always proceed immediately, yielding the element type's [zero value](https://go.dev/ref/spec#The_zero_value) after any previously sent values have been received. 

​	对于[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)为[通道](../Types#channel-types-通道型)的操作数`ch`，接收操作`<-ch`的值是从通道`ch`中接收的值，通道方向必须允许接收操作，接收操作的类型是通道的元素类型。这个表达式会阻塞，直到有一个可用的值。从一个 `nil`的通道接收时，将永远阻塞。在一个[已经关闭](../Built-inFunctions#close)的通道上的接收操作总是可以立即进行，并在任何先前发送的值被接收后，产生一个该元素类型的[零值](../ProgramInitializationAndExecution#the-zero-value-零值)。

```go 
v1 := <-ch
v2 = <-ch
f(<-ch)
<-strobe  // 等待，直到时钟脉冲并且丢弃接收值
```

A receive expression used in an [assignment statement](https://go.dev/ref/spec#Assignment_statements) or initialization of the special form 

​	在[赋值语句](../Statements#assignment-statement-赋值语句)或特殊形式的初始化中使用的一个接收表达式

```go 
x, ok = <-ch
x, ok := <-ch
var x, ok = <-ch
var x, ok T = <-ch
```

yields an additional untyped boolean result reporting whether the communication succeeded. The value of `ok` is `true` if the value received was delivered by a successful send operation to the channel, or `false` if it is a zero value generated because the channel is closed and empty. 

将产生一个额外的无类型布尔值结果，报告通信是否成功。如果收到的值是由一个成功的发送操作传递给通道的，那么`ok`的值为`true`，如果通道已关闭且为空，生成的值为零值，则 `ok` 的值为 `false`。

### Conversions 转换

A conversion changes the [type](https://go.dev/ref/spec#Types) of an expression to the type specified by the conversion. A conversion may appear literally in the source, or it may be *implied* by the context in which an expression appears. 

​	转换将表达式的[类型](../Types)改变为转换所指定的类型。转换可以出现在源文件中的字面量上，也可以隐含在由表达式所在的上下文。

An *explicit* conversion is an expression of the form `T(x)` where `T` is a type and `x` is an expression that can be converted to type `T`. 

​	式转换是形如 `T(x)` 的表达式，其中`T`是一个类型，`x`是可以被转换为`T`类型的表达式。

```
Conversion = Type "(" Expression [ "," ] ")" .
```

If the type starts with the operator `*` or `<-`, or if the type starts with the keyword `func` and has no result list, it must be parenthesized when necessary to avoid ambiguity: 

​	如果类型以运算符`*`或`<-`开头，或者如果类型以关键字`func`开头，并且没有结果列表，那么在必要时必须用`圆括号`括起来，以避免产生歧义：

```go 
*Point(p)        // 等同于 *(Point(p))
(*Point)(p)      // p 被转换为 *Point
<-chan int(c)    // 等同于 <-(chan int(c))
(<-chan int)(c)  // c 被转换为 <-chan int
func()(x)        // 函数签名 func() x
(func())(x)      // x 被转换为 func()
(func() int)(x)  // x 被转换为 func() int
func() int(x)    // x 被转换为 func() int   （无歧义 unambiguous）
```

A [constant](https://go.dev/ref/spec#Constants) value `x` can be converted to type `T` if `x` is [representable](https://go.dev/ref/spec#Representability) by a value of `T`. As a special case, an integer constant `x` can be explicitly converted to a [string type](https://go.dev/ref/spec#String_types) using the [same rule](https://go.dev/ref/spec#Conversions_to_and_from_a_string_type) as for non-constant `x`. 

​	一个[常量](../Constants)值`x`可以被转换为`T`类型，如果`x`可以用`T`的一个值来[表示](../PropertiesOfTypesAndValues#representability-可表示性)的话。作为一种特殊情况，可以使用 与 非常量`x`[相同的规则](#conversions-to-and-from-a-string-type-与字符串类型的转换)显式地将整数常量`x`转换为[字符串类型](../Types#string-types-字符串型)。

Converting a constant to a type that is not a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations) yields a typed constant. 

​	将常量转换为非[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)的类型，会生成一个有类型的常量。

```go 
uint(iota)               // uint 类型的 iota 值 
float32(2.718281828)     // float32 类型的 2.718281828
complex128(1)            // complex128 类型的 1.0 + 0.0i
float32(0.49999999)      // 0.5 of type float32 =>  float32 类型的 0.5 
float64(-1e-1000)        // float64 类型的0.0 
string('x')              // string 类型的 "x" 
string(0x266c)           // string 类型的 "♬"
myString("foo" + "bar")  // myString 类型的 "foobar"
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil) // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
int(1.2)   	  // 非法的：1.2 不能被 int 表示
string(65.0)  // 非法的：65.0 不是整数常量
```

Converting a constant to a type parameter yields a *non-constant* value of that type, with the value represented as a value of the type argument that the type parameter is [instantiated](https://go.dev/ref/spec#Instantiations) with. For example, given the function: 

​	将常量转换为一个类型参数会生成一个该类型的非常量值（non-constant value），该值表示为类型参数[实例化](#instantiations-实例化)时所带的类型实参的值。例如，给定一个（泛型）函数：

``` go
func f[P ~float32|~float64]() {
	… P(1.1) …
}
```

the conversion `P(1.1)` results in a non-constant value of type `P` and the value `1.1` is represented as a `float32` or a `float64` depending on the type argument for `f`. Accordingly, if `f` is instantiated with a `float32` type, the numeric value of the expression `P(1.1) + 1.2` will be computed with the same precision as the corresponding non-constant `float32` addition. 

转换`P(1.1)`的结果是一个`P`类型的非常量值（non-constant value），而值`1.1`被表示为`float32`或`float64`，这取决于`f`的类型参数。因此，如果`f`被实例化为`float32`类型，那么表达式`P(1.1)+1.2`的数值会用与非常量`float32`加法相同的精度进行计算。

A non-constant value `x` can be converted to type `T` in any of these cases: 

​	在以下情况下，非常量值 `x` 可以转换为类型 `T`：

- `x` is [assignable](https://go.dev/ref/spec#Assignability) to `T`. 
- `x`可以被[分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`T`。
- ignoring struct tags (see below), `x`'s type and `T` are not [type parameters](https://go.dev/ref/spec#Type_parameter_declarations) but have [identical](https://go.dev/ref/spec#Type_identity) [underlying types](https://go.dev/ref/spec#Underlying_types). 
- 忽略结构体标签（见下文），`x`的类型和`T`不是[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)，但有[一致的](../PropertiesOfTypesAndValues#type-identity-类型一致性)[底层类型](../Types)。
- ignoring struct tags (see below), `x`'s type and `T` are pointer types that are not [named types](https://go.dev/ref/spec#Types), and their pointer base types are not type parameters but have identical underlying types. 
- 忽略结构体标签（见下文），`x`的类型和`T`都是指针类型，且它们不是[命名类型](../Types)，它们的指针基类型不是类型参数，但有一致的底层类型。
- `x`'s type and `T` are both integer or floating point types. 
- `x`的类型和`T`都是整型或浮点型。
- `x`'s type and `T` are both complex types. 
- `x`的类型和`T`都是复数类型。
- `x` is an integer or a slice of bytes or runes and `T` is a string type. 
- `x`是一个整型、字节型、符文型的切片，`T`是一个字符串类型。
- `x` is a string and `T` is a slice of bytes or runes. 
- `x`是一个字符串类型，`T`是一个字节型、符文型的切片。
- `x` is a slice, `T` is an array [[Go 1.20]({{< ref "/langSpec/Appendix#go-120">}})] or a pointer to an array [[Go 1.17]({{< ref "/langSpec/Appendix#go-117">}})], and the slice and array types have [identical](https://go.dev/ref/spec#Type_identity) element types. 
- `x` 是切片， `T` 是数组[[Go 1.20]({{< ref "/langSpec/Appendix#go-120">}})] 或数组指针[[Go 1.17]({{< ref "/langSpec/Appendix#go-117">}})]，切片和数组类型具有[一致的](../PropertiesOfTypesAndValues#type-identity-类型一致性)元素类型。

Additionally, if `T` or `x`'s type `V` are type parameters, `x` can also be converted to type `T` if one of the following conditions applies:

​	此外，如果`T`或`x`的类型`V`是类型参数，如果满足以下条件之一，`x`也可以被转换为`T`类型：

- Both `V` and `T` are type parameters and a value of each type in `V`'s type set can be converted to each type in `T`'s type set.
- `V`和`T`都是类型参数，并且`V`的类型集中的每个类型的值都可以转换为`T`的类型集中的每个类型。
- Only `V` is a type parameter and a value of each type in `V`'s type set can be converted to `T`.
- 只有`V`是一个类型参数，并且`V`的类型集中的每个类型的值都可以转换为`T`。
- Only `T` is a type parameter and `x` can be converted to each type in `T`'s type set.
- 只有`T`是一个类型参数，并且`x`可以转换为`T`的类型集中的每个类型。

[Struct tags](https://go.dev/ref/spec#Struct_types) are ignored when comparing struct types for identity for the purpose of conversion:

​	为了转换的目的，在比较结构体类型的是否一致时，[结构体标签](../Types#struct-types-结构体型)被忽略：

```go 
type Person struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

var data *struct {
	Name    string `json:"name"`
	Address *struct {
		Street string `json:"street"`
		City   string `json:"city"`
	} `json:"address"`
}

var person = (*Person)(data)  // 忽略标签，这些的底层类型是一致的
```

Specific rules apply to (non-constant) conversions between numeric types or to and from a string type. These conversions may change the representation of `x` and incur a run-time cost. All other conversions only change the type but not the representation of `x`.

​	数值类型之间或与字符串类型之间的（非常量）转换有特殊的规则。这些转换可能会改变`x`的表示，并产生运行时开销。而所有其他的转换只改变`x`的类型而不改变其表示。

There is no linguistic mechanism to convert between pointers and integers. The package [`unsafe`](https://go.dev/ref/spec#Package_unsafe) implements this functionality under restricted circumstances.

​	在指针和整数之间没有语言机制可以直接进行转换。[unsafe]({{< ref "/stdLib/unsafe">}})包在受限制的情况下实现了这个功能。

#### Conversions between numeric types 数值型之间的转换

For the conversion of non-constant numeric values, the following rules apply:

​	对于非常量数值的转换，有以下特定规则：

1. When converting between [integer types](https://go.dev/ref/spec#Numeric_types), if the value is a signed integer, it is sign extended to implicit infinite precision; otherwise it is zero extended. It is then truncated to fit in the result type's size. For example, if `v := uint16(0x10F0)`, then `uint32(int8(v)) == 0xFFFFFFF0`. The conversion always yields a valid value; there is no indication of overflow.

2. 当在整型之间转换时，如果数值是有符号的[整型](../Types#numeric-types-数值型)，则进行符号扩展以达到隐式的无限精度；否则它被零扩展。然后，它被截断以适应结果类型的大小。例如，如果`v := uint16(0x10F0)`，那么`uint32(int8(v)) == 0xFFFFFFF0`。该转换总是产生一个有效的值；没有溢出的迹象。

   > 个人注释
   >
   > ```go
   > package main
   > 
   > import "fmt"
   > 
   > func main() {
   > 	v := uint16(0x10F0)
   > 	fmt.Printf("%#X\n", v)               // 0X10F0
   > 	fmt.Printf("%#X\n", int8(v))         // -0X10
   > 	fmt.Printf("%#X\n", uint32(int8(v))) // 0XFFFFFFF0
   > }
   > ```

   

3. When converting a [floating-point number](https://go.dev/ref/spec#Numeric_types) to an integer, the fraction is discarded (truncation towards zero).

4. 当把[浮点型](../Types#numeric-types-数值型)数值转换为整型时，小数会被丢弃（向零截断）。

   > 个人注释
   >
   > ```go
   > package main
   > 
   > import "fmt"
   > 
   > func main() {
   > 	v := 1.23
   > 	fmt.Printf("%v\n", v)       // 1.23
   > 	fmt.Printf("%v\n", int8(v)) // 1
   > }
   > 
   > ```
   >
   > 

5. When converting an integer or floating-point number to a floating-point type, or a [complex number](https://go.dev/ref/spec#Numeric_types) to another complex type, the result value is rounded to the precision specified by the destination type. For instance, the value of a variable `x` of type `float32` may be stored using additional precision beyond that of an IEEE-754 32-bit number, but float32(x) represents the result of rounding `x`'s value to 32-bit precision. Similarly, `x + 0.1` may use more than 32 bits of precision, but `float32(x + 0.1)` does not.

6. 当将一个整型或浮点型数值转换为浮点型，或将一个[复数型](../Types#numeric-types-数值型)数值转换为另一个复数类型时，结果值被舍入到目标类型所指定的精度。例如，`float32`类型的变量`x`的值可能会使用超出IEEE-754 32位数的额外精度来存储，但是`float32(x)`表示将`x`的值舍入到`32`位精度的结果。同样地，`x + 0.1`可能使用超过`32`位的精度，但是`float32(x + 0.1)`则不会。

   > 个人注释
   >
   > ​	请给出如何将一个复数型数值转换为另一个复数类型的示例：
   >
   > ```go
   > package main
   > 
   > import "fmt"
   > 
   > func main() {
   > 	c1 := complex(1.2, 3.0)
   > 	c2 := complex(float32(real(c1)), float32(imag(c1)))
   > 	fmt.Printf("%T,%v\n", c1, c1) // complex128,(1.2+3i)
   > 	fmt.Printf("%T,%v\n", c2, c2) // complex64,(1.2+3i)
   > }
   > 
   > ```
   > ​	如何使用`math.big`来解决浮点数计算问题？TODO

   

In all non-constant conversions involving floating-point or complex values, if the result type cannot represent the value the conversion succeeds but the result value is implementation-dependent.	

​	在所有涉及浮点值或复数值的非常量转换中，如果结果类型不能表示该值，转换仍会成功，但结果值取决于实现。

#### Conversions to and from a string type 与字符串类型的转换

1. Converting a slice of bytes to a string type yields a string whose successive bytes are the elements of the slice.

2. 将字节切片转换为字符串类型会产生一个字符串，其连续字节是切片的元素。 

   ```go 
   string([]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'})   // "hellø"
   string([]byte{})                                     // ""
   string([]byte(nil))                                  // ""
   
   type bytes []byte
   string(bytes{'h', 'e', 'l', 'l', '\xc3', '\xb8'})    // "hellø"
   
   type myByte byte
   string([]myByte{'w', 'o', 'r', 'l', 'd', '!'})       // "world!"
   myString([]myByte{'\xf0', '\x9f', '\x8c', '\x8d'})   // "🌍"
   ```

3. Converting a slice of runes to a string type yields a string that is the concatenation of the individual rune values converted to strings.

4. 将符文切片转换为字符串类型会产生一个字符串，该字符串是转换为字符串的各个符文值的连接。

   ```go 
   string([]rune{0x767d, 0x9d6c, 0x7fd4})   // "\u767d\u9d6c\u7fd4" == "白鵬翔"
   string([]rune{})                         // ""
   string([]rune(nil))                      // ""
   
   type runes []rune
   string(runes{0x767d, 0x9d6c, 0x7fd4})    // "\u767d\u9d6c\u7fd4" == "白鵬翔"
   
   type myRune rune
   string([]myRune{0x266b, 0x266c})         // "\u266b\u266c" == "♫♬"
   myString([]myRune{0x1f30e})              // "\U0001f30e" == "🌎"
   ```

5. Converting a value of a string type to a slice of bytes type yields a non-nil slice whose successive elements are the bytes of the string.

6. 将字符串类型的值转换为字节切片类型会产生一个非零切片，其连续元素是字符串的字节。

   ```go 
   []byte("hellø")             // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
   []byte("")                  // []byte{}
   
   bytes("hellø")              // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
   
   []myByte("world!")          // []myByte{'w', 'o', 'r', 'l', 'd', '!'}
   []myByte(myString("🌏"))    // []myByte{'\xf0', '\x9f', '\x8c', '\x8f'}
   ```

7. Converting a value of a string type to a slice of runes type yields a slice containing the individual Unicode code points of the string.

8. 将字符串类型的值转换为符文类型切片会产生一个切片，其中包含字符串的各个 Unicode 代码点。

   ```go 
   []rune(myString("白鵬翔"))   // []rune{0x767d, 0x9d6c, 0x7fd4}
   []rune("")                  // []rune{}
   
   runes("白鵬翔")              // []rune{0x767d, 0x9d6c, 0x7fd4}
   
   []myRune("♫♬")              // []myRune{0x266b, 0x266c}
   []myRune(myString("🌐"))    // []myRune{0x1f310}
   ```

9. Finally, for historical reasons, an integer value may be converted to a string type. This form of conversion yields a string containing the (possibly multi-byte) UTF-8 representation of the Unicode code point with the given integer value. Values outside the range of valid Unicode code points are converted to `"\uFFFD"`.

10. 最后，出于历史原因，整数值可以转换为字符串类型。这种形式的转换会产生一个字符串，其中包含具有给定整数值的 Unicode 代码点的（可能为多字节）UTF-8 表示形式。超出有效 Unicode 代码点范围的值将转换为 `"\uFFFD"` 。

   ```go 
   string('a')          // "a"
   string(65)           // "A"
   string('\xf8')       // "\u00f8" == "ø" == "\xc3\xb8"
   string(-1)           // "\ufffd" == "\xef\xbf\xbd"
   
   type myString string
   myString('\u65e5')   // "\u65e5" == "日" == "\xe6\x97\xa5"
   ```

Note: This form of conversion may eventually be removed from the language. The [`go vet`](https://go.dev/pkg/cmd/vet) tool flags certain integer-to-string conversions as potential errors. Library functions such as [`utf8.AppendRune`](https://go.dev/pkg/unicode/utf8#AppendRune) or [`utf8.EncodeRune`](https://go.dev/pkg/unicode/utf8#EncodeRune) should be used instead.

​	注意：这种形式的转换最终可能会从语言中删除。 `go vet` 工具将某些整数到字符串的转换标记为潜在错误。应改用库函数，例如 `utf8.AppendRune` 或 `utf8.EncodeRune` 。

#### Conversions from slice to array pointer 从切片到数组指针的转换

Converting a slice to an array yields an array containing the elements of the underlying array of the slice. Similarly, converting a slice to an array pointer yields a pointer to the underlying array of the slice. In both cases, if the [length](https://go.dev/ref/spec#Length_and_capacity) of the slice is less than the length of the array, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs.

​	将切片转换为数组会产生一个数组，其中包含切片的底层数组的元素。类似地，将切片转换为数组指针会产生一个指向切片的底层数组的指针。在这两种情况下，如果切片长度小于数组长度，则会发生运行时恐慌。

```go 
s := make([]byte, 2, 4)

a0 := [0]byte(s)
a1 := [1]byte(s[1:])     // a1[0] == s[1]
a2 := [2]byte(s)         // a2[0] == s[0]
a4 := [4]byte(s)         // panics: len([4]byte) > len(s)

s0 := (*[0]byte)(s)      // s0 != nil
s1 := (*[1]byte)(s[1:])  // &s1[0] == &s[1]
s2 := (*[2]byte)(s)      // &s2[0] == &s[0]
s4 := (*[4]byte)(s)      // panics: len([4]byte) > len(s)

var t []string
t0 := [0]string(t)       // ok for nil slice t
t1 := (*[0]string)(t)    // t1 == nil
t2 := (*[1]string)(t)    // panics: len([1]string) > len(t)

u := make([]byte, 0)
u0 := (*[0]byte)(u)      // u0 != nil
```

> 个人注释
>
> ​	给出以上示例的完整示例：
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	s := make([]byte, 2, 4)
> 	s0 := (*[0]byte)(s)            // s0 != nil
> 	fmt.Printf("%T,%#v\n", s0, s0) // *[0]uint8,&[0]uint8{}
> 	s1 := (*[1]byte)(s[1:])        // &s1[0] == &s[1]
> 	fmt.Printf("%T,%#v\n", s1, s1) // *[1]uint8,&[1]uint8{0x0}
> 	s2 := (*[2]byte)(s)            // &s2[0] == &s[0]
> 	fmt.Printf("%T,%#v\n", s2, s2) // *[2]uint8,&[2]uint8{0x0, 0x0}
> 	// panic: runtime error: cannot convert slice with length 2 to array or pointer to array with length 4
> 	//s4 := (*[4]byte)(s)            // panics: len([4]byte) > len(s)
> 	//fmt.Printf("%T,%#v\n", s4, s4) //
> 
> 	var t []string
> 	t0 := (*[0]string)(t)          // t0 == nil
> 	fmt.Printf("%T,%#v\n", t0, t0) // *[0]string,(*[0]string)(nil)
> 	// panic: runtime error: cannot convert slice with length 0 to array or pointer to array with length 1
> 	//t1 := (*[1]string)(t)          // panics: len([1]string) > len(t)
> 	//fmt.Printf("%T,%#v\n", t1, t1) //
> 
> 	u := make([]byte, 0)
> 	u0 := (*[0]byte)(u)            // u0 != nil
> 	fmt.Printf("%T,%#v\n", u0, u0) //*[0]uint8,&[0]uint8{}
> }
> 
> ```

> 个人注释
>
> ​	以上可以将切片转为数组指针，那么是否可以将切片转换为数组？=》 可以
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	s := make([]byte, 2, 4)
> 	s0 := ([0]byte)(s)             // s0 != nil
> 	fmt.Printf("%T,%#v\n", s0, s0) // [0]uint8,[0]uint8{}
> 	s1 := ([1]byte)(s[1:])         // &s1[0] == &s[1]
> 	fmt.Printf("%T,%#v\n", s1, s1) // [1]uint8,[1]uint8{0x0}
> 	s2 := ([2]byte)(s)             // &s2[0] == &s[0]
> 	fmt.Printf("%T,%#v\n", s2, s2) // [2]uint8,[2]uint8{0x0, 0x0}
> 	// panic: runtime error: cannot convert slice with length 2 to array or pointer to array with length 4
> 	//s4 := ([4]byte)(s)             // panics: len([4]byte) > len(s)
> 	//fmt.Printf("%T,%#v\n", s4, s4) //
> 
> 	var t []string
> 	t0 := ([0]string)(t)           // t0 == nil
> 	fmt.Printf("%T,%#v\n", t0, t0) // [0]string,[0]string{}
> 	// panic: runtime error: cannot convert slice with length 0 to array or pointer to array with length 1
> 	t1 := ([1]string)(t)           // panics: len([1]string) > len(t)
> 	fmt.Printf("%T,%#v\n", t1, t1) //
> 
> 	u := make([]byte, 0)
> 	u0 := ([0]byte)(u)             // u0 != nil
> 	fmt.Printf("%T,%#v\n", u0, u0) //[0]uint8,[0]uint8{}
> }
> 
> ```
>
> 

### Constant expressions 常量表达式

Constant expressions may contain only [constant](https://go.dev/ref/spec#Constants) operands and are evaluated at compile time.

​	常量表达式可以只包含[常量](../Constants)操作数，并在编译时进行求值。

Untyped boolean, numeric, and string constants may be used as operands wherever it is legal to use an operand of boolean, numeric, or string type, respectively.

​	无类型的布尔、数值和字符串常量可以在需要布尔、数值或字符串类型操作数的地方使用。

A constant [comparison](https://go.dev/ref/spec#Comparison_operators) always yields an untyped boolean constant. If the left operand of a constant [shift expression](https://go.dev/ref/spec#Operators) is an untyped constant, the result is an integer constant; otherwise it is a constant of the same type as the left operand, which must be of [integer type](https://go.dev/ref/spec#Numeric_types).

​	常量[比较](#comparison-operators-比较运算符)总是产生一个无类型的布尔常量。如果常量[移位表达式](#operators-操作符)的左操作数是一个无类型的常量，那么结果就是一个整型常量；否则就是一个与左操作数相同类型的常量（左操作数必须是[整型](../Types#numeric-types-数值型)）。

Any other operation on untyped constants results in an untyped constant of the same kind; that is, a boolean, integer, floating-point, complex, or string constant. If the untyped operands of a binary operation (other than a shift) are of different kinds, the result is of the operand's kind that appears later in this list: integer, rune, floating-point, complex. For example, an untyped integer constant divided by an untyped complex constant yields an untyped complex constant.

​	对无类型常量的任何其他操作都会得到一个相同类型的无类型常量，也就是布尔、整数、浮点、复数或字符串常量。如果一个二元运算（除移位外）的无类型操作数是不同种类的，那么结果就是出现在如下列表的操作数类型：整数，符文，浮点，复数。例如，一个无类型的整数常量除以一个无类型的复数常量，得到一个无类型的复数常量。

```go 
const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0         // d == 8     (untyped integer constant)
const e = 1.0 << 3         // e == 8     (untyped integer constant)
const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar"    // h == true  (untyped boolean constant)
const j = true             // j == true  (untyped boolean constant)
const k = 'w' + 1          // k == 'x'   (untyped rune constant)
const l = "hi"             // l == "hi"  (untyped string constant)
const m = string(k)        // m == "x"   (type string)
const Σ = 1 - 0.707i       //            (untyped complex constant)
const Δ = Σ + 2.0e-4       //            (untyped complex constant)
const Φ = iota*1i - 1/1i   //            (untyped complex constant)
```

Applying the built-in function `complex` to untyped integer, rune, or floating-point constants yields an untyped complex constant.

​	将内置函数 `complex` 应用于无类型的整数、符文或浮点常量，可以得到一个无类型的复数常量。

```go 
const ic = complex(0, c)   // ic == 3.75i  (untyped complex constant) 类型是complex128
const iΘ = complex(0, Θ)   // iΘ == 1i     (untyped complex constant) 类型是complex128
```

Constant expressions are always evaluated exactly; intermediate values and the constants themselves may require precision significantly larger than supported by any predeclared type in the language. The following are legal declarations:

​	`常量表达式总是被精确地求值`；中间值和常量本身可能需要比语言中任何预先声明的类型`所支持的精度大得多`。以下是合法的声明：

```go 
const Huge = 1 << 100         // Huge == 1267650600228229401496703205376  (untyped integer constant)
const Four int8 = Huge >> 98  // Four == 4 (type int8)
```

The divisor of a constant division or remainder operation must not be zero:

​	常量除法或取余操作的`除数一定不能为零`。

```go 
3.14 / 0.0   // illegal: division by zero
```

The values of *typed* constants must always be accurately [representable](https://go.dev/ref/spec#Representability) by values of the constant type. The following constant expressions are illegal:

​	类型常量的值必须总是可以准确地由常量类型的值来[表示](../PropertiesOfTypesAndValues#representability-可表示性)。下面的常量表达式是非法的：

```go 
uint(-1)     // -1 不能作为 uint 来表示
int(3.14)    // 3.14 不能作为 int 来表示
int64(Huge)  // 1267650600228229401496703205376 不能作为 int64 来表示
Four * 300   // 操作数 300 不能作为 int8（Four的类型） 来表示
Four * 100   // 乘积 400 不能作为 int8（Four的类型） 来表示
```

The mask used by the unary bitwise complement operator `^` matches the rule for non-constants: the mask is all 1s for unsigned constants and -1 for signed and untyped constants.

​	一元按位补运算符`^`使用的掩码符合非常量的规则：对于无符号常量来说是所有(掩码)位都是`1`，对于有符号和无类型的常量来说是`-1`。=> 仍有疑问？？

```go 
^1         // 无类型的整数常量，等于 -2
uint8(^1)  // 非法的: 相当于 uint8(-2)， -2 不能被 uint8 所表示
^uint8(1)  // 无类型的 uint8 常量， 相当于 0xFF ^ uint8(1) = uint8(0xFE)
int8(^1)   // 相当于 int8(-2)
^int8(1)   // 相当于 -1 ^ int8(1) = -2
```

Implementation restriction: A compiler may use rounding while computing untyped floating-point or complex constant expressions; see the implementation restriction in the section on [constants](https://go.dev/ref/spec#Constants). This rounding may cause a floating-point constant expression to be invalid in an integer context, even if it would be integral when calculated using infinite precision, and vice versa.

实现限制：编译器在计算无类型浮点或复数常量表达式时可能会使用舍入，请参见[常量](../Constants)部分的实现限制。这种舍入可能会导致浮点常量表达式在整数上下文中无效，即使它在使用无限精度计算时是整数，反之亦然。

### Order of evaluation 求值顺序

At package level, [initialization dependencies](https://go.dev/ref/spec#Package_initialization) determine the evaluation order of individual initialization expressions in [variable declarations](https://go.dev/ref/spec#Variable_declarations). Otherwise, when evaluating the [operands](https://go.dev/ref/spec#Operands) of an expression, assignment, or [return statement](https://go.dev/ref/spec#Return_statements), all function calls, method calls, [receive operations](https://go.dev/ref/spec#Receive operator), and [binary logical operations](https://go.dev/ref/spec#Logical_operators) are evaluated in lexical left-to-right order.

​	在包级别上，[初始化依赖项](../ProgramInitializationAndExecution#package-initialization-包的初始化)决定了[变量声明](../DeclarationsAndScope#variable-declarations-变量声明)中各个初始化表达式的求值顺序。除此之外，在求值表达式、赋值或[返回语句](../Statements#return-statements----return-语句)的[操作数](#operands-操作数)时，所有的函数调用、方法调用和通信操作都是按词法从左到右的顺序求值的。

For example, in the (function-local) assignment

​	例如，在（函数内部）赋值语句中

```go 
y[f()], ok = g(h(), i()+x[j()], <-c), k()
```

the function calls and communication happen in the order `f()`, `h()` (if `z` evaluates to false), `i()`, `j()`, `<-c`, `g()`, and `k()`. However, the order of those events compared to the evaluation and indexing of `x` and the evaluation of `y` and `z` is not specified, except as required lexically. For instance, `g` cannot be called before its arguments are evaluated.

函数调用和通信发生的顺序是`f()`, `h()`, `i(),` `j()`, `<-c`, `g()`, 和`k()`。然而，与`x`的求值和索引以及`y`的求值相比，这些事件的顺序没有被指定。

```go 
a := 1
f := func() int { a++; return a }
x := []int{a, f()} // x 可以是 [1, 2] 或是 [2, 2]： a 和 f() 的求值顺序没有被指定
m := map[int]int{a: 1, a: 2} // m 可以是 {2: 1} 或是 {2: 2}： 两个映射赋值的求值顺序没有被指定
n := map[int]int{a: f()} // n 可以是 {2: 3} 或是 {3: 3}： 键和值的求值顺序没有被指定
```

At package level, initialization dependencies override the left-to-right rule for individual initialization expressions, but not for operands within each expression:

​	在包级别上，对于独立的初始化表达式来说，初始化依赖项会覆盖其从左到右的求值规则，但不覆盖每个表达式中的操作数：

```go 
var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int        { return c }
func g() int        { return a }
func sqr(x int) int { return x*x }

// functions u and v are independent of all other variables and functions
// => 函数 u 和 v 独立于其它所有的变量和函数
```

The function calls happen in the order `u()`, `sqr()`, `v()`, `f()`, `v()`, and `g()`.

​	函数调用按照`u()`、`sqr()`、`v()`、`f()`、`v()`、`g()`的顺序发生。

Floating-point operations within a single expression are evaluated according to the associativity of the operators. Explicit parentheses affect the evaluation by overriding the default associativity. In the expression `x + (y + z)` the addition `y + z` is performed before adding `x`.

​	单个表达式中的浮点运算是按照运算符的结合性来求值的。显式的括号会通过覆盖默认的结合性来影响求值。在表达式`x + (y + z)`中，加法`y + z`会在加`x`之前进行。