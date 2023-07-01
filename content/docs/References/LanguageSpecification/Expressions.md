+++
title = "表达式"
date = 2023-05-17T09:59:21+08:00
weight = 11
description = ""
isCJKLanguage = true
draft = false

+++
## Expressions 表达式

> 原文：[https://go.dev/ref/spec#Expressions](https://go.dev/ref/spec#Expressions )

表达式通过将运算符和函数应用于操作数来规定值的计算。

### Operands 操作数

​	操作数表示表达式中的基本值。操作数可以是一个字面量，一个表示[常量](../DeclarationsAndScope#constant-declarations)、[变量](../DeclarationsAndScope#variable-declarations-变量声明)或函数的（可以是[限定的](#qualified-identifiers-限定标识符)）非[空白](../DeclarationsAndScope#blank-identifier-空白标识符)标识符，或者一对圆括号内的表达式。

```
Operand     = Literal | OperandName [ TypeArgs ] | "(" Expression ")" .
Literal     = BasicLit | CompositeLit | FunctionLit .
BasicLit    = int_lit | float_lit | imaginary_lit | rune_lit | string_lit .
OperandName = identifier | QualifiedIdent .
```

​	表示泛型函数的操作数名称后面可以跟一个[类型实参](#instantiations-实例化)列表；产生的操作数是一个[实例化过的](#instantiations-实例化)函数。

​	[空白标识符（即`_`）](../DeclarationsAndScope#blank-identifier-空白标识符)只能在[赋值语句](../Statements#assignment-statements-赋值语句)的左侧作为操作数出现。

实现限制：若操作数的类型是具有空[类型集](../Types#interface-types-接口型)的[类型形参](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)，则编译器不必报告错误。具有这种类型形参的函数不能被[实例化](#instantiations-实例化)；任何尝试都会导致实例化处的错误。

### Qualified identifiers 限定标识符

​	限定标识符是以包名作为前缀限定的标识符。包名和标识符都不能是[空白标识符（即`_`）](../DeclarationsAndScope#blank-identifier-空白标识符)。

```
QualifiedIdent = PackageName "." identifier .
```

​	限定标识符可以在不同包中访问一个标识符，但该标识符所在的包必须已经被[导入](../Packages#import-declarations-导入声明)。该标识符必须[可被导出](../DeclarationsAndScope#exported-identifiers-可导出的标识符)并在该包的[package block](../Blocks)中声明。

```go 
math.Sin // denotes the Sin function in package math 
		 //	=> 表示 math 包中的 Sin 函数
```

### Composite literals 复合字面量

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

​	LiteralType 的[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)`T`必须是一个结构体、数组、切片或映射类型（语法会强制执行这个约束，除非当类型是作为TypeName给出时）。元素和键的类型必须[可分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`T`类型的对应字段、元素和键类型；不需要进行额外的转换。

​	这里的键被解释为结构体字面量的字段名、数组字面量或切片字面量的索引、映射字面量的键。

​	对于映射字面量，所有的元素必须有一个键。用相同的字段名或常量键值指定多个元素是错误的。对于非常量的映射键，请参见关于[求值顺序](#order-of-evaluation-求值顺序)的章节。

​	对于结构体字面量来说，以下规则适用：

- 键必须是结构体类型中声明的字段名。
- 不包含任何键的元素列表必须按照字段的声明顺序为每个结构体字段列出一个元素。
- 如果任何元素有一个键，那么每个元素都必须有一个键。
- 包含键的元素列表不需要每个结构体字段都有一个元素。省略的字段将获得该字段类型的零值。
- `字面量可以省略元素列表；这样的字面量相当对其类型的求值为零值`。
- 为属于不在同一个包中的结构体（即该结构体在其他包中定义）的非可导出字段指定元素是错误的。

给定的声明：

```go 
type Point3D struct { x, y, z float64 }
type Line struct { p, q Point3D }
```

我们可以这样写：

```go 
// zero value for Point3D 
//=> Point3D 的零值
origin := Point3D{} 

// zero value for line.q.x 
//=> line.q.x 的零值
line := Line{origin, Point3D{y: -4, z: 12.3}}  
```

对于数组字面量和切片字面量，以下规则适用：

- 每个元素都有一个相关的整数索引，标记其在数组中的位置。
- 带键的元素使用该键作为其索引。键必须是一个可由`int`类型的值[表示的](../PropertiesOfTypesAndValues#representability-可表示性)非负常数；如果它是有类型的，则它必须是[整数类型](../Types#numeric-types-数值型)。
- 不带键的元素使用前一个元素的索引加1。如果第一个元素没有键，它的索引是0。

​	对一个复合字面量[取址](#address-operators-地址运算符)会产生一个指向唯一[变量](../Variables)的指针，该变量用字面量的值初始化。

```go 
var pointer *Point3D = &Point3D{y: 1000}
```

请注意，切片或映射类型的[零值](../ProgramInitializationAndExecution#the-zero-value-零值)与同一类型的初始化过但为空的值不同。因此，获取一个空切片或空映射复合字面量的地址与用[new](../Built-inFunctions#allocation-分配)分配一个新的切片或映射值的效果不同。

```go 
// p1 points to an initialized, empty slice with value []int{} and length 0 
//=> p1 指向一个值为 []int{} 且 长度为 0 的初始化过的空切片
p1 := &[]int{}    

// p2 points to an uninitialized slice with value nil and length 0 
//=> p2 指向一个值为 nil 且其长度为 0 的未初始化的切片
p2 := new([]int)  
```

​	数组字面量的长度是字面量类型中指定的长度。如果在字面量上提供的元素少于长度，缺少的元素将被设置为数组元素类型的零值。若提供的元素的索引值超出了数组的索引范围，将导致错误。标记法`...`指定的数组长度等于最大元素索引加1。

```go 
buffer := [10]string{}             // len(buffer) == 10
intSet := [6]int{1, 2, 3, 5}       // len(intSet) == 6
days := [...]string{"Sat", "Sun"}  // len(days) == 2
```

​	切片字面量描述了整个底层数组字面量。因此，切片字面量的长度和容量是最大元素索引加1。切片字面量的形式是：

```go 
[]T{x1, x2, … xn}
```

是对数组进行切片操作的简写：

```go 
tmp := [n]T{x1, x2, … xn}
tmp[0 : n]
```

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

​	当使用LiteralType的TypeName形式的复合字面量`作为操作数`出现在[关键字](../LexicalElements#keywords-关键字)和 "`if`"、"`for` "或 "`switch` "等语句块的`左花括号`之间，并且复合字面量没有被括在圆括号、方括号或花括号中时，会出现解析歧义。在这种罕见的情况下，字面量的左花括号被错误地解析为引入语句块的左花括号。为了解决这个问题，复合字面量`必须出现在圆括号`内。

```go 
if x == (T{a,b,c}[i]) { … }
if (x == T{a,b,c}[i]) { … }
```

有效的数组、切片和映射字面量的例子：

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

​	函数字面量表示一个匿名[函数](../DeclarationsAndScope#function-declarations-函数声明)。函数字面量不能声明`类型参数`。

``` go
functionLit = "func" Signature FunctionBody .
```

```go 
func(a, b int, z float64) bool { return a*b < int(z) }
```

​	函数字面量可以被分配给一个变量或直接调用。

```go 
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```

​	函数字面量可以是`闭包`：它们可以引用外层函数中定义的变量。然后，这些变量在外层的函数和函数字面量之间共享，并且只要它们可以被访问，它们就可以一直存在。

### Primary expressions 主表达式

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

​	对于[主表达式](#primary-expressions-主表达式)`x`（不是[包名](../Packages#package-clause-包子句)）来说，选择器表达式：

```go 
x.f
```

表示值`x`（有时是`*x`；见下文）的字段或方法`f`。标识符`f`被称为（字段或方法）`选择器`；它不能是[空白标识符](../DeclarationsAndScope#blank-identifierr-空白标识符)。选择器表达式的类型是`f`的类型。若`x`是包名，请参见关于[限定标识符](#qualified-identifiers-限定标识符)一节。

​	选择器`f`可以表示类型`T`的`f`字段或`f`方法，也可以指代`T`的[嵌入字段](../Types#struct-types-结构体型)或嵌入方法`f`。在`T`的一个嵌入字段`A`中声明的字段或方法`f`的深度是`A`中`f`的深度加1。

​	以下规则适用于选择器：

1. 对于类型为`T`或`*T`（`T`不是指针或接口类型）的值`x`，`x.f`表示`T`中存在这样一个最浅深度的字段或方法`f`。如果不是恰好有[仅有一个](../DeclarationsAndScope#uniqueness-of-identifiers-标识符的唯一性)`f`在最浅深度的话，那么选择器表达式是非法的。

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

2. 对于接口类型`I`的值`x`，`x.f`表示动态值`x`的名为`f`的实际方法。如果在`I`的[方法集](../PropertiesOfTypesAndValues#method-sets-方法集)中没有名为`f`的方法，那么选择器表达式是非法的。

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

3. 作为例外，如果`x`的类型是一个[已定义的](../DeclarationsAndScope#type-definitions-类型定义)指针类型，并且`(*x).f`是一个有效的表示一个字段（不是一个方法）的选择器表达式，那么`x.f`是`(*x).f`的简写。

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

4. 在所有其它情况下，`x.f`是非法的。

5. 如果`x`是值为`nil`的指针类型，并且`x.f`表示一个结构体字段，那么赋值或计算`x.f`会引起[运行时恐慌](../Run-timePanics)。

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

6. 如果`x`是值为`nil`的接口类型，那么[调用](#calls-调用)或[计值](#method-values-方法值)`x.f`方法会引起[运行时恐慌](../Run-timePanics)。

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

例如，给定声明：

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

​	如果`M`在类型`T`的[方法集](../PropertiesOfTypesAndValues#method-sets-方法集)中，那么`T.M`是一个可以作为普通函数来调用的函数，其实参与`M`相同，不过其前缀有一个额外的（作为该方法的接收器的）实参。

```go
MethodExpr    = ReceiverType "." MethodName .
ReceiverType  = Type .
```

​	考虑一个结构体类型`T`，它有两个方法：`Mv`，其接收器类型为`T`，和`Mp`，其接收器类型为`*T`。

```go 
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // 值类型的接收器
func (tp *T) Mp(f float32) float32 { return 1 }  // 指针类型的接收器

var t T
```

​	表达式：

```go 
T.Mv
```

生成一个与`Mv`等价的函数，但是它的第一个实参是一个显式的接收器；它具有以下签名：

```go 
func(tv T, a int) int
```

​	该函数可以在带有一个显式接收器的情况下正常调用，因此如下这五种调用是等同的：

```go 
t.Mv(7)
T.Mv(t, 7)
(T).Mv(t, 7)
f1 := T.Mv; f1(t, 7)
f2 := (T).Mv; f2(t, 7)
```

类似地，表达式

```go 
(*T).Mp
```

生成一个代表`Mp`的函数值，它的签名是：

```go 
func(tp *T, f float32) float32
```

​	对于一个`带值接收器`的方法，可以`推导出`一个带有显式指针接收器的函数，因此

```go 
(*T).Mv
```

生成了一个代表`Mv`的函数值，它的签名是：

```go 
func(tv *T, a int) int
```

Such a function indirects through the receiver to create a value to pass as the receiver to the underlying method; the method does not overwrite the value whose address is passed in the function call.

​	这样的函数通过接收器`间接地`创建了一个值，作为接收器传递给底层方法；该方法不会覆盖（其地址在函数调用中被传递的）那个值。=>该怎么翻译？？

​	这样的一个函数间接地通过其接收器创建了（用来作为接收器传递给其底层方法的）一个值；该（底层）方法不会覆盖这个值（因这个值的地址在这个函数调用才会被传递）

​	最后一种情况，将一个带指针接收器的方法`当做`一个带值接收器的函数，是非法的，因为指针接收器的方法不在值类型的方法集中。

​	从方法中推导出来的函数值是用`函数调用语法`来调用的；接收器被作为调用的第一个实参来提供。也就是说，`f := T.Mv`中的`f`是作为`f(t, 7)`被调用，而不是`t.f(7)`。要构造一个绑定接收器的函数，可以使用[函数字面量](#function-literals-函数字面量)或[方法值](#method-values-方法值)。

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

​	类型`T`既可以是接口类型，也可以是非接口类型。

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

​	表达式

```go 
t.Mv
```

​	生成一个类型如下的函数值：

```go 
func(int) int
```

​	这两种调用是等价的：

```go 
t.Mv(7)
f := t.Mv; f(7)
```

​	同样地，表达式

```go 
pt.Mp
```

​	生成一个类型如下的函数值：

```go 
func(float32) float32
```

​	和[选择器](#selectors-选择器)一样，若对以值作为接收器的非接口方法，使用指针来引用，则（Go语言）将自动解除对该指针的引用：`pt.Mv`等同于`(*pt).Mv`。

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

若主表达式的形式是：

```go 
a[x]
```

则表示可用`x`来检索的数组`a`、数组指针`a`、切片`a`、字符串`a`或映射`a`的元素，`x`分别被称为`索引`或`映射键`。以下规则适用：

如果`a`既不是映射也不是类型参数：

- 索引`x`必须是一个无类型的常量，或者其[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)必须是[整数类型](../Types#numeric-types-数值型)
- 常量索引必须是非负且可以用`int`类型的值来[表示](../PropertiesOfTypesAndValues#representability-可表示性)
- 无类型常量索引会被赋予`int`类型。
- 如果`0 <= x < len(a)`，则索引`x`在范围内，否则就超出了范围。

对于数组类型`A`的`a`：

- [常量](../Constants)索引必须在范围内
- 如果`x`在运行时超出了范围，就会发生[运行时恐慌](../Run-timePanics)
- `a[x]`是索引为`x`的数组元素，`a[x]`的类型是`A`的元素类型。

对于数组类型[指针](../Types#pointer-types-指针型)的`a`：

- `a[x]` 是 `(*a)[x]`的简写

对于[切片类型](../Types#slice-types-切片型)`S`的`a`：

- 如果`x`在运行时超出了范围，就会发生[运行时恐慌](../Run-timePanics)
- `a[x]`是索引`x`处的切片元素，`a[x]`的类型是`S`的元素类型。

For `a` of [string type](https://go.dev/ref/spec#String_types):

对于[字符串类型](../Types#string-types)的`a`：

- 如果字符串`a`是常量，那么[常量](../Constants)索引必须在范围内
- 如果`x`在运行时超出了范围，就会发生[运行时恐慌](../Run-timePanics)
- `a[x]`是索引`x`处的非常量字节值，`a[x]`的类型是`byte`。
- `a[x]`不能被赋值

对于[映射类型](../Types#map-types-映射型)为`M`的`a`：

- `x`的类型必须可以被[分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`M`的键类型
- 如果映射中有键`x`的项，那么`a[x]`就是键`x`的映射元素，`a[x]`的类型就是`M`的元素类型
- 如果映射为`nil`或者不包含任何项，`a[x]`是`M`的元素类型的[零值](../ProgramInitializationAndExecution#the-zero-value-零值)。

对于[参数类型](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)为`P`的`a`：

- 索引表达式`a[x]`必须对`P`的类型集中的所有类型的值有效。
- `P`的类型集中所有类型的元素类型必须是相同的。在此上下文中，字符串类型的元素类型是`byte`。
- 如果在`P`的类型集中有一个映射类型，那么该类型集中的所有类型必须是映射类型，且对应的键类型必须都是一致的。
- `a[x]`是索引为`x`的数组、切片或字符串元素，或者`P`实例化的类型实参中键为`x`的映射元素，`a[x]`的类型是（一致的）元素类型的类型。
- 如果`P`的类型集包括字符串类型，则`a[x]`不能再被赋值。

否则`a[x]`是非法的。

​	若将类型为`map[K]V`的映射`a`上的索引表达式使用在[赋值语句](../Statements#assignment-statements-赋值语句)或特殊格式的初始化中：

```go 
v, ok = a[x]
v, ok := a[x]
var v, ok = a[x]
```

将产生一个额外的`无类型`布尔值。如果键`x`存在于映射中，`ok`的值为`true`，否则为`false`。

​	若给`nil`映射的元素赋值，将导致[运行时恐慌](../Run-timePanics)。

### Slice expressions 切片表达式

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

主表达式

```go 
a[low : high]
```

构造了一个子字符串或切片。`a`的[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)必须是字符串、数组、数组指针、切片或者[bytestring](../PropertiesOfTypesAndValues#core-types-核心类型)。`low`和`high`所在的索引选择了哪些元素显示在操作数`a`的结果中。若结果的索引从0开始，则长度等于`high` 减去 `low`。在对数组`a`进行切片后

```go 
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]
```

切片`s`有类型`[]int`，长度3，容量4，以及元素

```go 
s[0] == 2
s[1] == 3
s[2] == 4
```

​	为方便起见，任何索引都可以被省略。缺少的`low`索引默认为0；缺少的`high`索引默认为被切片的操作数的长度：

```go 
a[2:]  // same as a[2 : len(a)]
a[:3]  // same as a[0 : 3]
a[:]   // same as a[0 : len(a)]
```

​	如果`a`是一个数组指针，则`a[low:high]`是`(*a)[low:high]`的简写。

​	对于数组或字符串，如果`0 <= low <= high <= len(a)`，则索引在范围内，否则就超出了范围。对于切片，索引的上限是切片的容量`cap(a)`，而不是长度。[常量](../Constants)索引必须是非负的，并且可以用`int`类型的值来[表示](../PropertiesOfTypesAndValues#representability-可表示性)；对于数组或字符串常量，常量索引也必须在范围内。如果两个索引都是常量，它们必须满足`low <= high`。如果索引在运行时超出范围，就会发生[运行时恐慌](../Run-timePanics)。

​	除了[无类型字符串](../Constants)外：

- 如果被切片的操作数是字符串或切片，则切片的操作结果是一个与操作数相同类型的非常量值。

- 如果被切片的操作数是无类型的字符串，则切片的操作结果是一个`string`类型的非常量值。

- 如果被切片的操作数是（必须[可被寻址](#address-operators-地址运算符-地址运算符)的）数组，则切片的操作结果是一个与数组的元素类型一致的切片。

​	如果有效切片表达式的切片操作数是`nil`切片，那么切片的操作结果就是一个`nil`切片。否则，如果切片的操作结果是一个切片，则它与操作数共享底层数组。

```go 
var a [10]int
s1 := a[3:7]   // s1 的底层数组是数组 a；&s1[2] == &a[5]
s2 := s1[1:4]  // s2 的底层数组是 s1 的底层数组 a； &s2[1] == &a[5]
s2[1] = 42     // s2[1] == s1[2] == a[5] == 42；它们都指向相同的底层数组元素
```

#### Full slice expressions 完整的切片表达式

主表达式

```go 
a[low : high : max]
```

构造了一个与简单切片表达式`a[low : high]`相同类型的切片，并且具有相同的长度和元素。此外，它通过将结果切片设置为`max 减去 low`的容量。`a`的[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)必须是数组，数组指针，或者切片（但不是字符串）。在对数组`a`进行切分后

```go 
a := [5]int{1, 2, 3, 4, 5}
t := a[1:3:5]
```

切片`t`有类型`[]int`，长度2，容量4，以及元素

```go 
t[0] == 2
t[1] == 3
```

​	与简单切片表达式一样，如果`a`是一个数组指针，则`a[low:high:max]`是`(*a)[low:high:max]`的简写。如果切片的操作数是一个数组，它必须是[可被寻址的](#address-operators-地址运算符)。

​	如果`0 <= low <= high <= max <= cap(a)`，则索引就在范围内，否则就超出了范围。[常量](../Constants)索引必须是非负数，并且可以用`int`类型的值来[表示](../PropertiesOfTypesAndValues#representability-可表示性)；对于数组，常量索引也必须在范围内。如果多个索引是常量，那么出现的常量必须在相对于彼此的范围内。如果索引在运行时超出了范围，就会发生[运行时恐慌](../Run-timePanics)。

### Type assertions 类型断言

​	对于[接口类型](../Types#interface-ypes-接口型)但非[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)的表达式`x`和类型`T`的主表达式

```go 
x.(T)
```

断言了`x`不是`nil`，并且`x`中存储的值是`T`类型。标记法`x.(T)`被称为`类型断言`。

​	更确切地说，如果`T`不是接口类型，则`x.(T)`断言`x`的动态类型与`T`的类型[一致](../PropertiesOfTypesAndValues#type-identity-类型一致性)。在这种情况下，`T`必须实现`x`的（接口）类型；否则类型断言是无效的，因为对于`x`来说存储`T`类型的值是不可能的。如果`T`是一个接口类型，则`x.(T)` 断言`x`的动态类型[实现](../Types#implementing-an-interface-实现一个接口)了接口`T`。

If the type assertion holds, the value of the expression is the value stored in `x` and its type is `T`. If the type assertion is false, a [run-time panic](https://go.dev/ref/spec#Run_time_panics) occurs. In other words, even though the dynamic type of `x` is known only at run time, the type of `x.(T)` is known to be `T` in a correct program.

​	如果类型断言成立，表达式的值就是存储在`x`中的值，其类型是`T`。 如果类型断言不成立，就会发生[运行时恐慌](../Run-timePanics)。换句话说，尽管`x`的动态类型只有在运行时才知道，但在一个正确的程序中，`x.(T)`的类型是已知的，是`T`。=> 仍有疑问？？

```go 
var x interface{} = 7          // x has dynamic type int and value 7 => x 有动态类型 int 以及值 7
i := x.(int)                   // i has type int and value 7 => i 有类型 int 和 值 7

type I interface { m() }

func f(y I) {
	s := y.(string)        // illegal: string does not implement I (missing method m) => 非法的：string 没有实现 I （缺少m方法）
	r := y.(io.Reader)     // r has type io.Reader and the dynamic type of y must implement both I and io.Reader => r 有类型 io.Reader，并且 y 的动态类型必须同时实现 I 和 io.Reader
	…
}
```

​	在[赋值语句](../Statements#assignment-statements-赋值语句)或特殊格式的初始化中使用的类型断言

```go 
v, ok = x.(T)
v, ok := x.(T)
var v, ok = x.(T)
var v, ok interface{} = x.(T) // dynamic types of v and ok are T and bool => v 的动态类型是 T， ok 的动态类型是 bool
```

将产生一个额外的无类型布尔值。如果断言成立，`ok`的值为`true`。否则为`false`，并且`v`的值是`T`类型的[零值](../ProgramInitializationAndExecution#the-zero-value-零值)。在这种情况下`不会`发生[运行时恐慌](../Run-timePanics)。

### Calls 调用

给定一个表达式`f`，其[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)为[函数类型](../Types#function-types-函数型)`F`,

```go 
f(a1, a2, … an)
```

带实参`a1, a2, ... an`调用了`f`。除了一种特殊情况以外，实参必须是[可分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`F`的参数类型的单值表达式，并且在函数被调用之前被求值。该表达式的类型是`F`的结果类型。方法调用是类似的，但是方法本身被指定为一个（在该方法的接收器类型的值上的）选择器。

```go 
math.Atan2(x, y)  // function call => 函数调用
var pt *Point
pt.Scale(3.5)     // method call with receiver pt => 带接收器 pt 的方法调用
```

​	如果`f`表示一个泛型函数，在它被调用或作为函数值使用之前，必须将其[实例化](#instantiations-实例化)。

​	在函数调用中，函数值和实参以[通常的顺序](#order-of-evaluation-求值顺序)被求值。在它们被求值之后，调用的参数被按值传递给函数，然后被调用的函数开始执行。当函数返回时，函数的返回参数按值传递给调用者。

​	调用一个`nil`的函数值会引起[运行时恐慌](../Run-timePanics)。

​	作为一个特例，如果一个函数或方法的返回值`g`在数量上相等，并且可以单独分配给另一个函数或方法`f`的参数，那么调用`f(g(parameters_of_g))`将在把`g`的返回值依次绑定到`f`的参数后再调用`f`。对`f`的调用必须不包含对`g`的调用以外的参数，并且`g`必须至少有一个返回值。如果`f`有一个在最后的`...`参数，它将被分配给`g`的（在分配完普通参数后所剩余的）返回值。

```go 
func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}

func Join(s, t string) string {
	return s + t
}

if Join(Split(value, len(value)/2)) != value {
	log.Panic("test fails")
}
```

​	如果`x`的[方法集](../PropertiesOfTypesAndValues#method-sets-方法集)（`x`的类型）包含`m`，并且实参列表可以分配给`m`的参数列表，那么方法调用`x.m()`是有效的。如果`x`是[可寻址的](#address-operators-地址运算符)，并且`&x`的方法集包含`m`，则`x.m()`就是`(&x).m()`的简写：

```go 
var p Point
p.Scale(3.5)
```

没有明确的方法类型，也`没有方法字面量`。

### Passing arguments to `...` parameters 向...参数传递实参

​	如果`f`是带有一个`...T`类型的位置在最后的参数`p`的可变函数，那么在`f`内部，`p`的类型等同于`[]T`类型。如果`f`被调用时没有给`p`的实参，传递给`p`的值是`nil`。否则，传递的值是一个新的`[]T`类型的切片，这个切片带有一个新的底层数组，这个底层数组的连续元素作为实参，并且这些实参都必须[可分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`T`。因此，切片的长度和容量等于绑定到`p`的实参的数量，并且对每次调用（实参数量）都可能有所不同。

给出函数和调用

```go 
func Greeting(prefix string, who ...string)
Greeting("nobody")
Greeting("hello:", "Joe", "Anna", "Eileen")
```

​	在`Greeting`函数第一次被调用时，`who`的值为`nil`，在第二次被调用时，`who`的值为`[]string{"Joe", "Anna", "Eileen"}`。

​		如果最后一个实参可以分配给切片类型`[]T`，并且后面是`...`，那么它就会在不改变值的情况下传递一个`...T`参数。在这种情况下，不会创建新的切片。

给出切片`s`并调用

```go 
s := []string{"James", "Jasmine"}
Greeting("goodbye:", s...)
```

在`Greeting`函数中，`who`将拥有与`s`相同的底层数组的值。

### Instantiations 实例化

​	`泛型函数`或`泛型`是通过用`类型实参`替换`类型参数`而被实例化的。实例化分两步进行：

1. 在泛型声明中，每个类型参数都被替换为其对应的类型实参。这种替换发生在整个函数或类型声明中，包括类型参数列表本身和该列表中的每个类型。
2. 替换之后，每个类型实参必须[实现](../Types#interface-types-接口型)相应类型参数的[约束](../DeclarationsAndScope#type-constraints-类型约束)（若有需要则实例化它）。否则实例化就会失败。

​	实例化一个类型会产生一个新的非泛型的[命名类型](../Types)；实例化一个函数会产生一个新的非泛型的函数。

```go 
type parameter list    type arguments    after substitution
  类型参数列表             类型实参            替换后

[P any]                int               int implements any
[S ~[]E, E any]        []int, int        []int implements ~[]int, int implements any
[P io.Writer]          string            illegal: string doesn't implement io.Writer
```

​	对于泛型函数，可以明确地提供类型实参，也可以靠部分或完整地[推断](#type-inference)出它们。非[调用](#calls-调用)的泛型函数需要一个类型实参列表用于实例化；如果该列表是部分的，那么所有剩余的类型实参必须是可推断的。被调用的泛型函数可以提供一份（可能是部分的）类型实参列表，（如果省略的类型实参可以从普通（非类型）函数参数中推断出来）也可以完全省略。

```go 
func min[T ~int|~float64](x, y T) T { … }

f := min                   // illegal: min must be instantiated with type arguments when used without being called => 非法的：在非调用情况下使用 min 时，min 必须用类型实参实例化
minInt := min[int]         // minInt has type func(x, y int) int => minInt 类型为 func(x, y int) int   <= 非调用的情况
a := minInt(2, 3)          // a has value 2 of type int => a 的类型为 int，值为 2 
b := min[float64](2.0, 3)  // b has value 2.0 of type float64 => b 的类型为 float64，值为 2.0 
c := min(b, -1)            // c has value -1.0 of type float64 => c 的类型为 float64，值为 -1.0 
```

​	部分类型实参列表不能是空的；至少第一个（类型）实参必须存在。该列表是完整的类型实参列表的前缀，剩下的参数需要推断。宽泛地说，类型实参可以从 "从右到左"省略。

```go 
func apply[S ~[]E, E any](s S, f(E) E) S { … }

f0 := apply[]                  // illegal: type argument list cannot be empty => 非法的：类型实参列表不能为空
f1 := apply[[]int]             // type argument for S explicitly provided, type argument for E inferred => S 的类型实参被明确提供了，E 的类型实参则需要被推断
f2 := apply[[]string, string]  // both type arguments explicitly provided => 两个类型实参都被明确提供了

var bytes []byte
r := apply(bytes, func(byte) byte { … })  // both type arguments inferred from the function arguments => 两个类型实参 都需要从函数实参中被推断出来
```

对于`泛型，所有的类型实参必须都被明确提供`。

### Type inference 类型推断

​	缺失的函数类型实参可以通过一系列的步骤来推断，如下所述。每个步骤都试图使用已知的信息来推断额外的类型实参。一旦所有的类型实参都是已知的，类型推断就会停止。在类型推断完成后，仍然有必要将所有类型实参替换为类型参数，并验证每个类型实参是否[实现](../Types#interface-types-接口型)了相关的约束；推断的类型实参有可能无法实现约束，在这种情况下，实例化就会失败。

类型推断是基于：

- [类型参数列表](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)
- 使用已知类型实参(如果有的话)初始化过的用于替换的映射 `M`
- （可能是空的）普通函数实参列表（仅在函数调用的情况下）。

然后进行以下步骤：

1. 对所有`有类型`普通函数实参应用[函数实参类型推断](#function-argument-type-inference-函数实参类型推断)
6. 应用[约束类型推断](#constraint-type-inference-约束类型推断)
7. 使用每个`无类型`函数实参的默认类型，对所有`无类型`普通函数实参应用[函数实参类型推断](#function-argument-type-inference-函数实参类型推断)
8. 应用[约束类型推断](#constraint-type-inference-约束类型推断)

​	如果没有普通函数实参或无类型函数实参数，则跳过相应的步骤。如果前一步没有推断出任何新的类型实参，则跳过[约束类型推断](#constraint-type-inference-约束类型推断)，但如果有缺失的类型实参，则（这个步骤：即[约束类型推断](#constraint-type-inference-约束类型推断)）至少要运行一次。

​	替换映射`M`贯穿所有的步骤，每个步骤都可以向`M`添加条目。一旦`M`为每个类型参数提供了一个类型实参或者推断步骤失败，这个过程就会停止。如果推断步骤失败，或者在最后一步之后`M`仍然缺少类型实参，则类型推断失败。

#### Type unification 类型联合

​	类型推断是基于类型联合的。单一的联合步骤适用于一个[替换映射](#type-inference-类型推断)和两种类型，其中一个或两个可能是或包含类型参数。替换映射跟踪已知的（显式提供的或已经推断出的）类型实参：该映射包含每个类型参数`P`和相应的已知类型实参`A`的一个条目`P`→`A`。在联合过程中，已知的类型实参在比较类型时取代了它们对应的类型参数。联合过程是寻找使两个类型等同的替换映射条目的过程。

​	对于联合来说，如果两个类型不包含当前类型参数列表中的任何类型参数，或者它们是忽略了通道方向的通道类型，或者它们的底层类型是等同的，那么它们（这两个类型）就是一致的。

Unification works by comparing the structure of pairs of types: their structure disregarding type parameters must be identical, and types other than type parameters must be equivalent. A type parameter in one type may match any complete subtype in the other type; each successful match causes an entry to be added to the substitution map. If the structure differs, or types other than type parameters are not equivalent, unification fails.

​	联合是通过比较类型对的结构来实现的：它们的结构在不考虑类型参数时必须是一致的，而类型参数以外的类型必须是等同的。一个类型中的类型参数可以匹配另一个类型中的任何完整的子类型；每个成功的匹配都会将一个条目添加到替换映射中。如果结构不同，或者除类型参数外的其他类型不等同，那么联合就会失败。=> 仍有疑问？？

For example, if `T1` and `T2` are type parameters, `[]map[int]bool` can be unified with any of the following:

​	例如，如果`T1`和`T2`是类型参数，`[]map[int]bool`可以与以下任何一种类型联合：=> 仍有疑问？？

```go 
[]map[int]bool   // types are identical => 类型是一致的
T1               // adds T1 → []map[int]bool to substitution map => 添加 T1 → []map[int]bool 到 替换映射中
[]T1             // adds T1 → map[int]bool to substitution map => 添加 T1 → map[int]bool 到 替换映射中
[]map[T1]T2      // adds T1 → int and T2 → bool to substitution map => 添加 T1 → int 和 T2 → bool 到 替换映射中
```

On the other hand, `[]map[int]bool` cannot be unified with any of 

另一方面，`[]map[int]bool`不能与以下任何一个联合起来  => 仍有疑问？？

```go 
int              // int is not a slice => int 不是切片
struct{}         // a struct is not a slice => 结构体不是切片
[]struct{}       // a struct is not a map => 结构体不是映射
[]map[T1]string  // map element types don't match => 映射元素类型不能匹配
```

​	作为这个通用规则的一个例外，由于[定义类型](../DeclarationsAndScope#type-definitions-类型定义)`D`和类型字面量`L`从来都是`不等同`的，联合将`D`的底层类型与`L`进行比较。例如，给定定义类型

```go 
type Vector []float64
```

和类型字面量`[]E`，联合过程会将`[]float64`与`[]E`进行比较，并在替换映射中添加一个条目`E`→`float64`。

#### Function argument type inference 函数实参类型推断

Function argument type inference infers type arguments from function arguments: if a function parameter is declared with a type `T` that uses type parameters, [unifying](https://go.dev/ref/spec#Type_unification) the type of the corresponding function argument with `T` may infer type arguments for the type parameters used by `T`.

​	函数实参类型推断从函数实参中推断出类型实参：如果函数参数声明时带有使用了类型参数的类型`T`，那么将对应函数实参的类型与`T`进行[联合](#type-unification-类型联合)，可能推断出被`T`所使用的类型参数的类型实参。=>仍有疑问？？

例如，给定[泛型函数](../DeclarationsAndScope#function-declarations-函数声明)

```go 
func scale[Number ~int64|~float64|~complex128](v []Number, s Number) []Number
```

和调用

```go 
var vector []float64
scaledVector := scale(vector, 42)
```

`Number`的类型实参，可以通过联合`vector`的类型与对应的参数类型中推断出：`[]float64`和`[]Number`在结构上匹配，且`float64`与`Number`匹配。这就把`Number`→`float64`这个条目添加到[替换映射](#type-inference-类型推断)中。在第一轮`函数实参类型推断`中，无类型实参，比如这里的第二个函数实参`42`，会被忽略，只有在还有未解决的类型参数时才会考虑。

​	推断发生在两个独立的阶段；每个阶段在一个特定的（参数，实参）对的列表上操作：

1. 列表`Lt`包含所有使用了类型参数的参数类型和`有类型的`函数实参所组成的（参数，实参）对。
2. 列表`Lu`包含所有剩下的参数类型为单一类型参数的对。在这个列表中，各自的函数实参是`无类型的`。

任何其他的（参数，实参）对都被忽略。

​	根据结构，列表`Lu`中对的实参是无类型的常量（或无类型的布尔比较结果）。而且由于无类型值的[默认类型](../Constants)总是预先声明的非复合类型，它们永远不能与复合类型相匹配，所以仅考虑作为单一类型参数的参数类型就足够了。

​	每个列表都在一个独立的阶段中被处理：

1. 在第一阶段，`Lt`中的每各配对的参数和实参类型被联合。如果一个配对的联合成功了，它可能会产生新的条目，并被添加到替换映射`M`中。
2. 第二阶段考虑列表`Lu`中的条目。此阶段，将忽略已为其确定类型实参的类型参数。对于剩下的每一对，参数类型（这是一个单一的类型参数）和相应的无类型实参的[默认类型](../Constants)进行联合。如果联合失败，则类型推断失败。

​	当联合过程成功时，即使所有的类型实参在最后一个列表元素被处理之前就被推断出来了，对每个列表的处理还是会继续进行，直到所有的列表元素都被考虑。

示例：

```go 
func min[T ~int|~float64](x, y T) T

var x int
min(x, 2.0)    // T is int, inferred from typed argument x; 2.0 is assignable to int => T 是 int 类型，它是从有类型的实参 x 中推断出来的，2.0 是 可分配给 int 的
min(1.0, 2.0)  // T is float64, inferred from default type for 1.0 and matches default type for 2.0 => T 是 float64 类型，它是从 1.0 的默认类型中推断出来的，且与 2.0 的默认类型匹配
min(1.0, 2)    // illegal: default type float64 (for 1.0) doesn't match default type int (for 2) => 非法的：1.0 的默认类型 float64 与 2 的默认类型 int 不匹配
```

​	在例子`min(1.0, 2)`中，处理函数实参`1.0`会产生替换映射条目`T` → `float64`。因为处理过程一直持续到所有无类型的实参被考虑，所以报告了一个错误。这确保了类型推断不依赖于无类型实参的顺序。

#### Constraint type inference 约束类型推断

Constraint type inference infers type arguments by considering type constraints. If a type parameter `P` has a constraint with a [core type](https://go.dev/ref/spec#Core_types) `C`, [unifying](https://go.dev/ref/spec#Type_unification) `P` with `C` may infer additional type arguments, either the type argument for `P`, or if that is already known, possibly the type arguments for type parameters used in `C`.

​	约束类型推断通过考虑类型约束来推断类型实参。如果类型参数`P`有一个[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)`C`的约束，将`P`与`C`[联合](#type-unification-类型联合)起来可能会推断出额外的类型实参，要么是`P`的类型实参，（如果这个是已知，则）要么可能是`C`中使用的类型参数的类型实参。=>仍有疑问？？已知指的是什么，是 类型参数P有一个核心类型C的约束？

例如，考虑具有类型参数`List`和`Elem`的类型参数列表：

```go 
[List ~[]Elem, Elem any]
```

约束类型推断可以从`List`的类型实参推断出`Elem`的类型，因为`Elem`是`List`的核心类型`[]Elem`中的一个类型参数。如果类型实参是`Bytes`类型：

```go 
type Bytes []byte
```

将`Bytes`的底层类型与核心类型联合起来，就意味着将`[]byte`与`[]Elem`联合起来。这一联合成功了，并产生了`Elem`→`byte`的[替换映射](#type-inference-类型推断)条目。因此，在这个例子中，约束类型推断可以从第一个类型实参推断出第二个类型实参。

Using the core type of a constraint may lose some information: In the (unlikely) case that the constraint's type set contains a single [defined type](https://go.dev/ref/spec#Type_definitions) `N`, the corresponding core type is `N`'s underlying type rather than `N` itself. In this case, constraint type inference may succeed but instantiation will fail because the inferred type is not in the type set of the constraint. Thus, constraint type inference uses the *adjusted core type* of a constraint: if the type set contains a single type, use that type; otherwise use the constraint's core type. 

​	使用约束的核心类型可能会丢失一些信息。在（不太可能的）情况下，约束的类型集包含单一的[定义类型](../DeclarationsAndScope#type-definitions-类型定义)`N`，相应的核心类型是`N`的底层类型而不是`N`本身。在这种情况下，约束的类型推断可能会成功，但实例化会失败，因为推断的类型不在约束的类型集中。因此，约束类型推断使用调整后的约束的核心类型：如果类型集包含一个单一的类型，则使用该类型；否则使用约束的核心类型。=> 仍有疑问？？

​	通常，约束类型推断分两个阶段进行。从一个给定的替换映射`M`开始

1. 对于所有带调整过的核心类型的类型参数，将该类型参数与该类型联合起来。如果任何联合失败，约束类型推断就会失败。
2. 此时，`M`中的一些条目可能将类型参数映射到其他类型参数或包含类型参数的类型。对于`M`中每个条目`P`→`A`，其中`A`是或包含类型参数`Q`，而`M`中存在条目`Q`→`B`，用`A`中相应的`B`替换这些`Q`。在无法进一步替换时，则停止。

​	约束类型推断的结果是从类型参数`P`到类型实参`A`的最终替换映射`M`，其中在任何`A`中都没有出现类型参数`P`。

例如，给定类型参数列表

```go 
[A any, B []C, C *A]
```

以及为类型参数`A`提供的单一类型实参`int`，初始替换映射`M`包含条目`A`→`int`。

​	在`第一阶段`，类型参数`B`和`C`与它们各自约束的核心类型联合起来。这就在`M`中加入了`B`→`[]C`和`C`→`*A`的条目。

此时，`M`中存在两个条目，其右手边是或包含类型参数，而`M`中存在其他条目：`[]C`和`*A`。在`第二阶段`，这些类型参数被替换成它们各自的类型。这发生在哪个顺序上并不重要。从第一阶段后的`M`的状态开始：

```
A` → `int`, `B` → `[]C`, `C` → `*A
```

用`int`替换→右边的`A`。

```
A` → `int`, `B` → `[]C`, `C` → `*int
```

用`*int`替换→右边的`C`。

```
A` → `int`, `B` → `[]*int`, `C` → `*int
```

此时，不可能再进一步替换，该替换映射已满。因此，`M`代表类型参数到类型实参列表的最终映射。

### Operators 操作符

操作符将操作数组合成表达式。

```
Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
```

​	比较操作符将在[其他地方](#comparison-operators-比较运算符)讨论。对于其他二元运算符，操作数类型必须是[一致的](../PropertiesOfTypesAndValues#type-identity-类型一致性)，除非操作涉及移位或无类型的[常量](../Constants)。对于只涉及常量的操作，请参见[常量表达式](#constant-expressions-常量表达式)部分。

​	除了移位操作之外，如果一个操作数是`无类型`常量，而另一个操作数不是，那么该常量将被隐式地[转换](#conversions-转换)为另一个操作数的类型。

​	移位表达式中的`右操作数`必须是[整数类型](../Types#numeric-types-数值型)，或者是可以用`uint`类型的值[表示](../PropertiesOfTypesAndValues#representability-可表示性)的`无类型`常量。如果一个非常量移位表达式的`左操作数`是一个`无类型`常量，那么它首先被隐式地转换为假设移位表达式被其左操作数单独替换时的类型。

```go 
var a [1024]byte
var s uint = 33

// The results of the following examples are given for 64-bit ints.
// 以下示例的结果是针对 64-bits 整型给出的。
var i = 1<<s                   // 1 has type int => 1 拥有 int 类型 ，i == 8589934592
var j int32 = 1<<s             // 1 has type int32; j == 0 => 1 拥有 int32 类型；j == 0
var k = uint64(1<<s)           // 1 has type uint64; k == 1 <<33 => 1 拥有 uint64 类型；k == 1 << 33, k == 8589934592
var m int = 1.0<<s             // 1.0 has type int; m == 1<<33 => 1.0 拥有 int 类型；m == 1 << 33, m == 8589934592 
var n = 1.0<<s == j            // 1.0 has type int32; n == true => 1.0 拥有 int32 类型；n == true
var o = 1<<s == 2<<s           // 1 and 2 have type int; o == false => 1 和 2 拥有 int 类型；o == false
var p = 1<<s == 1<<33          // 1 has type int; p == true => 1 拥有 int 类型；p == true
var u = 1.0<<s                 // illegal: 1.0 has type float64, cannot shift => 非法的：1.0 拥有 float64 类型，但不能移位
var u1 = 1.0<<s != 0           // illegal: 1.0 has type float64, cannot shift => 非法的： 1.0 拥有 float64 类型，但不能移位
var u2 = 1<<s != 1.0           // illegal: 1 has type float64, cannot shift => 非法的： 1 拥有 float64 类型，但不能移位
var v1 float32 = 1<<s          // illegal: 1 has type float32, cannot shift => 非法的：1 拥有 float32 类型，但不能移位
var v2 = string(1<<s)          // illegal: 1 is converted to a string, cannot shift => 非法的： 1 被转换成 string 类型，但不能移位
var w int64 = 1.0<<33          // 1.0<<33 is a constant shift expression; w == 1<<33 => 1.0 << 33 是一个常量移位表达式；w == 1 << 33, w == 8589934592
var x = a[1.0<<s]              // panics: 1.0 has type int, but 1<<33 overflows array bounds => 恐慌：1.0 拥有 int 类型，但 1 << 33 溢出了数组的边界
var b = make([]byte, 1.0<<s)   // 1.0 has type int; len(b) == 1<<33 => 1.0 拥有 int 类型；len(b) == 1 << 33

// The results of the following examples are given for 32-bit ints,
// which means the shifts will overflow.
// 以下示例的结果是针对 32-bits 整型给出。 这意味着移位将会溢出。
var mm int = 1.0<<s            // 1.0 has type int; mm == 0 => 1.0 拥有 int 类型；mm == 0
var oo = 1<<s == 2<<s          // 1 and 2 have type int; oo == true => 1 和 2 拥有 int 类型； oo == true
var pp = 1<<s == 1<<33         // illegal: 1 has type int, but 1<<33 overflows int => 非法的：1 拥有 int 类型，但 1 << 33 溢出了 int 的范围
var xx = a[1.0<<s]             // 1.0 has type int; xx == a[0] => 1.0 拥有 int 类型；xx == a[0] 
var bb = make([]byte, 1.0<<s)  // 1.0 has type int; len(bb) == 0 => 1.0 拥有 int 类型；len(bb) == 0
```

=> 仍有疑问？？

#### Operator precedence 优先级运算符

​	一元运算符的优先级最高。由于`++`和`--`运算符构成的是语句，而不是表达式，因此它们不属于运算符等级体系。因此，语句`*p++`与`(*p)++`相同。

​	二元运算符有五个优先级别。`乘法运算符`绑定最强，其次是`加法运算符`、比较运算符、`&&`（逻辑AND），最后是`||`（逻辑OR）。

```go 
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
```

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

​	算术运算符适用于数字值，产生的结果`与第一个操作数的类型`相同。四个标准的算术运算符（`+`、`-`、`*`、`/`）适用于[整型](../Types#numeric-types-数值型)、[浮点型](../Types#numeric-types-数值型)和[复数型](../Types#numeric-types-数值型)；`+`也适用于[字符串](../Types#string-types-字符串型)。位逻辑运算符和移位运算符只适用于整型。

```
+    sum        （和）            integers, floats, complex values, strings
-    difference （差）            integers, floats, complex values
*    product    （积）            integers, floats, complex values
/    quotient   （商）            integers, floats, complex values
%    remainder  （余）            integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << integer >= 0
>>   right shift            integer >> integer >= 0
```

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

`x * y`的乘积 和 `s += x * y`的加法分别以`float32`或`float64`精度计算，这取决于`F`的类型实参。

#### Integer operators 整数运算符

对于两个整数值`x`和`y`，它们的整数商`q = x / y`和余数`r = x % y`满足以下关系：

```
x = q*y + r  and  |r| < |y|
```

with `x / y`被截断到零（"[截断除法](https://en.wikipedia.org/wiki/Modulo_operation)"）。

```
 x     y     x / y     x % y
 5     3       1         2
-5     3      -1        -2
 5    -3      -1         2
-5    -3       1        -2
```

​	这条规则有一个例外：如果）`被除数`（dividend`x`是`x`的`int`类型的最负值，那么商`q = x / -1`就等于`x`（而`r = 0`），这是由于二元补码的[整数溢出](#integer-overflow-整数溢出)：

```
                         x, q
int8                     -128
int16                  -32768
int32             -2147483648
int64    -9223372036854775808
```

​	如果`除数`（divisor）是一个[常量](../Constants)，那么它一定不能为零。如果除数在运行时为零，就会发生[运行时恐慌](../Run-timePanics)。如果除数是非负数，并且除数是2的常数幂，除法可以用`右移`来代替，计算余数可以用`按位与`操作来代替。

```
 x     x / 4     x % 4     x >> 2     x & 3
 11      2         3         2          3
-11     -2        -3        -3          1
```

​	移位运算符通过右操作数指定的`移位计数`对左操作数进行移位，`移位计数`必须为非负数。如果`移位计数`在运行时为负数，就会发生[运行时恐慌](../Run-timePanics)。如果`左操作数`是`有符号整数`，移位操作符实现`算术移位`；如果是`无符号整数`，则实现`逻辑移位`。`移位计数`没有上限。`移位计数`为`n`的移位行为就像左操作数被`1`移了`n`次。因此，`x<<1`与`x*2`相同，`x>>1`与`x/2`相同（但向右移位被截断到负无穷大）。

对于整数操作数，一元运算符`+`、`-`和`^`的定义如下：

```
+x                          is 0 + x
-x    negation              is 0 - x
^x    bitwise complement    is m ^ x  with m = "all bits set to 1" for unsigned x
                                      and  m = -1 for signed x
```

#### Integer overflow 整数溢出

​	对于[无符号整型值](../Types#numeric-types-数值型)，`+`、`-`、`*`和`<<`运算是以`2n`为模来计算的，其中`n`是无符号整型的位宽。广义上讲，这些无符号整型操作`在溢出时丢弃高位`，程序可以依靠 "`wrap around`"。

​	对于`有符号整型值`，`+`、`-`、`*`、`/`和`<<`运算`可以合法地溢出`，其产生的值是存在的，并且可以被有符号整型表示法、其操作和操作数明确地定义。溢出不会引起[运行时恐慌](../Run-timePanics)。在假设不发生溢出的情况下，编译器可能不会优化代码。例如，它不会假设`x<x+1`总是真的。

#### Floating-point operators 浮点运算符

​	对于浮点数和复数，`+x`与`x`相同，而`-x`是负的`x`。浮点数或复数除以0的结果，在IEEE-754标准中没有规定；是否会发生[运行时恐慌](../Run-timePanics)是由具体实现决定的。

An implementation may combine multiple floating-point operations into a single fused operation, possibly across statements, and produce a result that differs from the value obtained by executing and rounding the instructions individually. An explicit [floating-point type](https://go.dev/ref/spec#Numeric_types) [conversion](https://go.dev/ref/spec#Conversions) rounds to the precision of the target type, preventing fusion that would discard that rounding.

​	某些实现可能会将多个浮点运算合并为一个单一的融合运算，可能会跨越语句，产生的结果与单独执行和舍入指令得到的值不同。明确的[浮点类型转换](#conversions)是按照目标类型的精度进行舍入的，这样就可以避免融合时放弃舍入的做法。=> 仍有疑问？？

​	例如，一些体系架构提供了一个 "`fused multiply and add`"（`FMA`）指令，其在计算`x*y+z`时，不对中间结果`x*y`进行舍入。这些例子显示了Go的实现何时可以使用该指令：

```go 
// FMA allowed for computing r, because x*y is not explicitly rounded: 
// => FMA 允许被用来计算 r, 因为 x*y 不会被明确地进行舍入：
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

​	字符串可以使用`+`运算符或`+=`赋值运算符进行连接：

```
s := "hi" + string(c)
s += " and good bye"
```

​	字符串加法通过连接操作数创建一个新的字符串。

### Comparison operators 比较运算符

​	比较运算符比较两个操作数，并产生一个无类型布尔值。

```
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

​	在任何比较中，第一个操作数必须是[可分配](../PropertiesOfTypesAndValues#assignability-可分配性)给第二个操作数的类型，反之亦然。

​	相等运算符`==`和`!=`适用于可比较的操作数。排序运算符`<`, `<=`, `>`, 和`>=`适用于被排序的操作数。这些术语和比较结果的定义如下：

- 布尔值是可比较的。如果两个布尔值都是`true`或者都是`false`，那么它们是相等的。
- 按照通常的方式，整数值是可比较的并且是可排序的。
- 按照IEEE-754标准的定义，浮点值是可比较的并且是可排序的。
- 复数值是可比较的。如果`real(u) == real(v)`和`imag(u) == imag(v)`，则这两个复数值`u`和`v`是相等的。
- 字符串值是可（按字节顺序）比较的并且是可（按字节顺序）排序的。
- 指针值是可比较的。如果两个指针值指向同一个变量，或者两个指针值都是`nil`，则它们的值是相等的。指向不同的[零尺寸](../SystemConsiderations#size-and-alignment-guarantees-大小和对齐保证)变量的指针值可能相等，也可能不相等。
- 通道值是可比较的。如果两个通道是由同一个调用[make](../Built-inFunctions#making-slices-maps-and-channels-制作切片映射和通道)创建的，或者它们的值都为`nil`，则它们的值是相等的。
- 接口值是可比较的。如果两个接口值有[一致的](../PropertiesOfTypesAndValues#type-identity-类型一致性)动态类型和相同的动态值，或者两者的值都是`nil`，则它们的值是相等的。
- 非接口类型 `X` 的值 `x` 和接口类型 `T` 的值 `t` ，在 `X` 类型的值是可比较的并且 `X` [实现](../Types#implementing-an-interface-实现一个接口) `T` 时是可比较的。如果 `t` 的动态类型等于 `X`，且 `t` 的动态值等于 `x`，则它们是相等的。
- 如果结构体值的所有字段都是可比较的，那么结构体值就是可比较的。如果两个结构体值对应的非[空白](../DeclarationsAndScope#blank-identifier-空白标识符)字段相等，那么它们就是相等的。
- 如果数组元素类型的值是可比较的，那么数组值是可比较的。如果两个数组的对应元素是相等的，那么这两个数组值就是相等的。

​	对两个动态类型相同的接口值进行比较，如果它们的类型值不具有可比性，则会引起[运行时恐慌](../Run-timePanics)。这种行为不仅适用于直接的接口值比较，也适用于比较接口值的数组或带有接口值字段的结构体。

​	`切片值、映射值和函数值是不可比较的`。然而，作为一种特殊情况，切片值、映射值或函数值可以与预先声明的标识符`nil`比较。指针值、通道值和接口值与`nil`的比较也是允许的，并遵循上述的通用规则。

```go 
const c = 3 < 4            // c is the untyped boolean constant true => c 是无类型的布尔常量 true

type MyBool bool
var x, y int
var (
	// The result of a comparison is an untyped boolean. => 比较的结果为一个无类型的布尔值。
	// The usual assignment rules apply. => 使用通用赋值规则。
	b3        = x == y // b3 has type bool => b3 拥有 bool 类型
	b4 bool   = x == y // b4 has type bool => b4 拥有 bool 类型
	b5 MyBool = x == y // b5 has type MyBool => b5 拥有 MyBool 类型
)
```

### Logical operators 逻辑运算符

​	逻辑运算符适用于[布尔](../Types#boolean-types-布尔型)值，并产生一个与操作数相同类型的结果。右操作数是按条件进行求值的。

```
&&    conditional AND    p && q  is  "if p then q else false"
||    conditional OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"
```

### Address operators 地址运算符

​	对于类型为`T`的操作数`x`，寻址操作`&x`产生一个类型为`*T`的指针指向`x`。该操作数`x`必须是可寻址的，也就是说，它要么是一个变量、指针间接（pointer indirection）或`对切片的索引操作（slice indexing operation，是一个名词）`；要么是一个可寻址结构体操作数的字段选择器；要么是一个可寻址数组的数组索引操作。作为可寻址要求的一个例外，`x`也可以是一个（可能是括号内的）[复合字面量](#composite-literals-复合字面量)。如果对`x`的求值会引起[运行时恐慌](../Run-timePanics)，那么对`&x`的求值也会引起[运行时恐慌](../Run-timePanics)。

​	对于指针类型`*T`的操作数`x`，指针间接`*x`表示`x`所指向的`T`类型的[变量]()，如果`x`是`nil`，试图求值`*x`将导致[运行时恐慌](../Run-timePanics)。

```go 
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // causes a run-time panic => 导致一个 运行时恐慌
&*x  // causes a run-time panic => 导致一个 运行时恐慌
```

### Receive operator 接收操作符

​	对于[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)为[通道](../Types#channel-types-通道型)的操作数`ch`，接收操作`<-ch`的值是从通道`ch`中接收的值，通道方向必须允许接收操作，接收操作的类型是通道的元素类型。这个表达式会阻塞，直到有一个可用的值。从一个 `nil`的通道接收时，将永远阻塞。在一个[已经关闭](../Built-inFunctions#close)的通道上的接收操作总是可以立即进行，并在任何先前发送的值被接收后，产生一个该元素类型的[零值](../ProgramInitializationAndExecution#the-zero-value-零值)。

```go 
v1 := <-ch
v2 = <-ch
f(<-ch)
<-strobe  // wait until clock pulse and discard received value => 等待，直到时钟脉冲并且丢弃接收值
```

​	在[赋值语句](../Statements#assignment-statement-赋值语句)或特殊形式的初始化中使用的一个接收表达式

```go 
x, ok = <-ch
x, ok := <-ch
var x, ok = <-ch
var x, ok T = <-ch
```

将产生一个额外的无类型布尔值结果，报告通信是否成功。如果收到的值是由一个成功的发送操作传递给通道的，那么`ok`的值为`true`，如果是一个因为通道关闭且空而产生的零值，则为`false`。

### Conversions 转换

​	转换将表达式的[类型](../Types)改变为转换所指定的类型。转换可以出现在源文件中的字面量上，也可以隐含在由表达式所在的上下文。

​	显式转换是`T(x)`形式的表达式，其中`T`是一个类型，`x`是可以被转换为`T`类型的表达式。

```
Conversion = Type "(" Expression [ "," ] ")" .
```

​	如果类型以运算符`*`或`<-`开头，或者如果类型以关键字`func`开头，并且没有结果列表，那么在必要时必须用`圆括号`括起来，以避免产生歧义：

```go 
*Point(p)        // same as *(Point(p))
(*Point)(p)      // p is converted to *Point
<-chan int(c)    // same as <-(chan int(c))
(<-chan int)(c)  // c is converted to <-chan int
func()(x)        // function signature func() x
(func())(x)      // x is converted to func()
(func() int)(x)  // x is converted to func() int
func() int(x)    // x is converted to func() int (unambiguous)
```

​	一个[常量](../Constants)值`x`可以被转换为`T`类型，如果`x`可以用`T`的一个值来[表示](../PropertiesOfTypesAndValues#representability-可表示性)的话。作为一种特殊情况，可以使用 与 非常量`x`[相同的规则](#conversions-to-and-from-a-string-type-与字符串类型的转换)显式地将整数常量`x`转换为[字符串类型](../Types#string-types-字符串型)。

​	将常量转换为非[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)的类型，会得到一个有类型的常量。

```go 
uint(iota)               // iota value of type uint => uint 类型的 iota 值 
float32(2.718281828)     // 2.718281828 of type float32 => float32 类型的 2.718281828
complex128(1)            // 1.0 + 0.0i of type complex128 =>  complex128 类型的 1.0 + 0.0i
float32(0.49999999)      // 0.5 of type float32 =>  float32 类型的 0.5 
float64(-1e-1000)        // 0.0 of type float64 =>  float64 类型的0.0 
string('x')              // "x" of type string =>  string 类型的 "x" 
string(0x266c)           // "♬" of type string => string 类型的 "♬"
myString("foo" + "bar")  // "foobar" of type myString => myString 类型的 "foobar"
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant => 不是常量：[]byte{'a'} 不是常量
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type => 不是常量：nil 不是常量，*int 不是布尔、数值、字符串类型
int(1.2)                 // illegal: 1.2 cannot be represented as an int => 非法的：1.2 不能被 int 表示
string(65.0)             // illegal: 65.0 is not an integer constant => 非法的：65.0 不是整数常量
```

​	将常量转换为一个类型参数会产生一个该类型的非常量值，该值表示为类型参数[实例化](#instantiations-实例化)时所带的类型实参的值。例如，给定一个函数：

``` go
func f[P ~float32|~float64]() {
	… P(1.1) …
}
```

转换`P(1.1)`的结果是一个`P`类型的非常量值，而值`1.1`被表示为`float32`或`float64`，这取决于`f`的类型参数。因此，如果`f`被实例化为`float32`类型，那么表达式`P(1.1)+1.2`的数值会用与非常量`float32`加法相同的精度进行计算。

非常量值`x`可以在以下的任何情况下被转换为`T`类型：

- `x`可以被[分配](../PropertiesOfTypesAndValues#assignability-可分配性)给`T`。
- 忽略结构体标签（见下文），`x`的类型和`T`不是[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)，但有[一致的](../PropertiesOfTypesAndValues#type-identity-类型一致性)[底层类型](../Types)。
- 忽略结构体标签（见下文），`x`的类型和`T`是指针类型，不是[命名类型](../Types)，它们的指针基类型不是类型参数，但有一致的底层类型。
- `x`的类型和`T`都是整型或浮点型。
- `x`的类型和`T`都是复数类型。
- `x`是一个整型、字节型、符文型的切片，`T`是一个字符串类型。
- `x`是一个字符串类型，`T`是一个字节型、符文型的切片。
- `x`是一个切片，`T`是一个指向数组的指针，而且切片和数组的类型有[一致的](../PropertiesOfTypesAndValues#type-identity-类型一致性)元素类型。

​	此外，如果`T`或`x`的类型`V`是类型参数，如果满足以下条件之一，`x`也可以被转换为`T`类型：

- `V`和`T`都是类型参数，并且`V`的类型集中的每个类型的值都可以转换为`T`的类型集中的每个类型。
- 只有`V`是一个类型参数，并且`V`的类型集中的每个类型的值都可以转换为`T`。
- 只有`T`是一个类型参数，并且`x`可以转换为`T`的类型集中的每个类型。

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

var person = (*Person)(data)  // ignoring tags, the underlying types are identical => 忽略标签，底层类型是一致的
```

​	数值类型之间或与字符串类型之间的（非常量）转换有特殊的规则。这些转换可能会改变`x`的表示，并产生运行时间成本。所有其他的转换只改变`x`的类型而不改变其表示。

​	没有语言机制可以在`指针和整型之间进行转换`。[unsafe]({{< ref "/stdLib/unsafe">}})包在受限制的情况下实现了这个功能。

#### Conversions between numeric types 数值型之间的转换

对于非常量数值的转换，适用以下规则：

1. 当在整型之间转换时，如果数值是有符号的[整型](../Types#numeric-types-数值型)，那么它被符号位扩展到隐式的无限精度；否则它被零扩展。然后，它被截断以适应结果类型的大小。例如，如果`v := uint16(0x10F0)`，那么`uint32(int8(v)) == 0xFFFFFFF0`。该转换总是产生一个有效的值；没有溢出的迹象。
2. 当把[浮点型](../Types#numeric-types-数值型)数值转换为整型时，小数会被丢弃（向零截断）。
3. 当将一个整型或浮点型数值转换为浮点型，或将一个[复数型](../Types#numeric-types-数值型)数值转换为另一个复数类型时，结果值被舍入到目标类型所指定的精度。例如，`float32`类型的变量`x`的值可能会使用超出IEEE-754 32位数的额外精度来存储，但是`float32(x)`表示将`x`的值舍入到`32`位精度的结果。同样地，`x + 0.1`可能使用超过`32`位的精度，但是`float32(x + 0.1)`则不会。

​	在所有涉及浮点值或复数值的非常量转换中，如果结果类型不能表示该值，转换仍会成功，但结果值取决于实现。

#### Conversions to and from a string type 与字符串类型的转换

1. 将有符号或无符号的整型值转换为字符串类型，会产生一个包含该整型值的UTF-8表示的字符串。在有效的Unicode码点范围之外的值会被转换为`"\uFFFD"`。 

   ```go 
   string('a')       // "a"
   string(-1)        // "\ufffd" == "\xef\xbf\xbd"
   string(0xf8)      // "\u00f8" == "ø" == "\xc3\xb8"
   
   type myString string
   myString(0x65e5)  // "\u65e5" == "日" == "\xe6\x97\xa5"
   ```
   
2. 将字节切片转换为字符串类型，可以得到一个字符串，其连续的字节是该切片的元素。

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
   
3. 将符文切片转换为字符串类型，可以得到一个字符串，即转换为字符串的各个符文值的连接。

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
   
4. 将字符串类型的值转换为字节类型的切片，得到一个切片，其连续的元素是字符串的字节。

   ```go 
   []byte("hellø")             // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
   []byte("")                  // []byte{}
   
   bytes("hellø")              // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
   
   []myByte("world!")          // []myByte{'w', 'o', 'r', 'l', 'd', '!'}
   []myByte(myString("🌏"))    // []myByte{'\xf0', '\x9f', '\x8c', '\x8f'}
   ```
   
5. 将字符串类型的值转换为符文类型的切片，会得到一个包含该字符串的各个Unicode码点的切片。

   ```go 
   []rune(myString("白鵬翔"))   // []rune{0x767d, 0x9d6c, 0x7fd4}
   []rune("")                  // []rune{}
   
   runes("白鵬翔")              // []rune{0x767d, 0x9d6c, 0x7fd4}
   
   []myRune("♫♬")              // []myRune{0x266b, 0x266c}
   []myRune(myString("🌐"))    // []myRune{0x1f310}
   ```

#### Conversions from slice to array pointer 从切片到数组指针的转换

​	将切片转换为数组指针，会得到一个指向切片底层数组的指针。如果切片的[长度](../Built-inFunctions#length-and-capacity-长度和容量)小于数组的长度，就会发生[运行时恐慌](../Run-timePanics)。

```go 
s := make([]byte, 2, 4)
s0 := (*[0]byte)(s)      // s0 != nil
s1 := (*[1]byte)(s[1:])  // &s1[0] == &s[1]
s2 := (*[2]byte)(s)      // &s2[0] == &s[0]
s4 := (*[4]byte)(s)      // panics: len([4]byte) > len(s)

var t []string
t0 := (*[0]string)(t)    // t0 == nil
t1 := (*[1]string)(t)    // panics: len([1]string) > len(t)

u := make([]byte, 0)
u0 := (*[0]byte)(u)      // u0 != nil
```

### Constant expressions 常量表达式

​	常量表达式可以只包含[常量](../Constants)操作数，并在编译时进行求值。

​	无类型的布尔、数值和字符串常量可以作为操作数使用，只要合法地分别使用布尔、数值或字符串类型的操作数。

​	常量[比较](#comparison-operators-比较运算符)总是产生一个无类型的布尔常量。如果常量[移位表达式](#operators-操作符)的左操作数是一个无类型的常量，那么结果就是一个整型常量；否则就是一个与左操作数相同类型的常量（左操作数必须是[整型](../Types#numeric-types-数值型)）。

​	任何其他对无类型常量的操作都会产生一个相同类型的无类型常量，也就是布尔、整数、浮点、复数或字符串常量。如果一个二元运算（除移位外）的无类型操作数是不同种类的，那么结果就是出现在如下列表的操作数类型：整数，符文，浮点，复数。例如，一个无类型的整数常量除以一个无类型的复数常量，得到一个无类型的复数常量。

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

​	将内置函数 `complex` 应用于无类型的整数、符文或浮点常量，可以得到一个无类型的复数常量。

```go 
const ic = complex(0, c)   // ic == 3.75i  (untyped complex constant)
const iΘ = complex(0, Θ)   // iΘ == 1i     (type complex128)
```

​	`常量表达式总是被精确地求值`；中间值和常量本身可能需要比语言中任何预先声明的类型`所支持的精度大得多`。以下是合法的声明：

```go 
const Huge = 1 << 100         // Huge == 1267650600228229401496703205376  (untyped integer constant)
const Four int8 = Huge >> 98  // Four == 4                                (type int8)
```

常量除法或取余操作的`除数一定不能为零`。

```go 
3.14 / 0.0   // illegal: division by zero
```

​	类型常量的值必须总是可以准确地由常量类型的值来[表示](../PropertiesOfTypesAndValues#representability-可表示性)。下面的常量表达式是非法的：

```go 
uint(-1)     // -1 cannot be represented as a uint => -1 不能作为 uint 来表示
int(3.14)    // 3.14 cannot be represented as an int => 3.14 不能作为 int 来表示
int64(Huge)  // 1267650600228229401496703205376 cannot be represented as an int64 => 1267650600228229401496703205376 不能作为 int64 来表示
Four * 300   // operand 300 cannot be represented as an int8 (type of Four) => 操作数 300 不能作为 int8（Four的类型） 来表示
Four * 100   // product 400 cannot be represented as an int8 (type of Four) => 乘积 400 不能作为 int8（Four的类型） 来表示
```

​	一元按位补运算符`^`使用的掩码符合非常量的规则：对于无符号常量来说是所有(掩码)位都是`1`，对于有符号和无类型的常量来说是`-1`。=> 仍有疑问？？

```go 
^1         // untyped integer constant, equal to -2 => 无类型的整数常量，等于 -2
uint8(^1)  // illegal: same as uint8(-2), -2 cannot be represented as a uint8 => 非法的: 相当于 uint8(-2)， -2 不能被 uint8 所表示
^uint8(1)  // typed uint8 constant, same as 0xFF ^ uint8(1) = uint8(0xFE) => 无类型的 uint8 常量， 相当于 0xFF ^ uint8(1) = uint8(0xFE)
int8(^1)   // same as int8(-2) => 相当于 int8(-2)
^int8(1)   // same as -1 ^ int8(1) = -2 => 相当于 -1 ^ int8(1) = -2
```

实现限制：编译器在计算无类型浮点或复数常量表达式时可能会使用舍入，请参见[常量](../Constants)部分的实现限制。这种舍入可能会导致浮点常量表达式在整数上下文中无效，即使它在使用无限精度计算时是整数，反之亦然。

### Order of evaluation 求值顺序

​	在包级别上，[初始化依赖关系](../ProgramInitializationAndExecution#package-initialization-包的初始化)决定了[变量声明](../DeclarationsAndScope#variable-declarations-变量声明)中各个初始化表达式的求值顺序。除此之外，在求值表达式、赋值或[返回语句](../Statements#return-statements----return-语句)的[操作数](#operands-操作数)时，所有的函数调用、方法调用和通信操作都是按词法从左到右的顺序求值的。

例如，在（函数局部）赋值中

```go 
y[f()], ok = g(h(), i()+x[j()], <-c), k()
```

函数调用和通信发生的顺序是`f()`, `h()`, `i(),` `j()`, `<-c`, `g()`, 和`k()`。然而，与`x`的求值和索引以及`y`的求值相比，这些事件的顺序没有被指定。

```go 
a := 1
f := func() int { a++; return a }
x := []int{a, f()}            // x may be [1, 2] or [2, 2]: evaluation order between a and f() is not specified => x 可以是 [1, 2] 或是 [2, 2]： a 和 f() 的求值顺序没有被指定
m := map[int]int{a: 1, a: 2}  // m may be {2: 1} or {2: 2}: evaluation order between the two map assignments is not specified => m 可以是 {2: 1} 或是 {2: 2}： 两个映射赋值的求值顺序没有被指定
n := map[int]int{a: f()}      // n may be {2: 3} or {3: 3}: evaluation order between the key and the value is not specified => n 可以是 {2: 3} 或是 {3: 3}： 键和值的求值顺序没有被指定
```

​	在包级别上，对于独立的初始化表达式来说，初始化依赖关系会覆盖其从左到右的求值规则，但不覆盖每个表达式中的操作数：

```go 
var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int        { return c }
func g() int        { return a }
func sqr(x int) int { return x*x }

// functions u and v are independent of all other variables and functions
// => 函数 u 和 v 独立于其它所有的变量和函数
```

函数调用按照`u()`、`sqr()`、`v()`、`f()`、`v()`、`g()`的顺序发生。

​	单个表达式中的浮点运算是按照运算符的结合性来求值的。显式的括号会通过覆盖默认的结合性来影响求值。在表达式`x + (y + z)`中，加法`y + z`会在加`x`之前进行。