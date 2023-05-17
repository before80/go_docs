+++
title = "语句"
date = 2023-05-17T09:59:21+08:00
weight = 12
description = ""
isCJKLanguage = true
draft = false
+++
## Statements 语句

> 原文：[https://go.dev/ref/spec#Statements ](https://go.dev/ref/spec#Statements )

语句控制执行。

```
Statement =
	Declaration | LabeledStmt | SimpleStmt |
	GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
	FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
	DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
```

### Terminating statements 终止语句

​	终止语句中断了一个[块](../Blocks)中的常规控制流。下列语句是终止性的：

1. "`return`"或 "`goto` "语句。
2. 对内置函数`panic`的调用。
3. 语句列表以终止语句结束的块。
4. 一个 "`if` "语句，其中：
   - 存在 "`else` "分支，并且
   - 两个分支都是终止语句。
5. 一个 "`for` "语句，其中：
   - 没有 "`break` "语句引用"`for` "语句，并且
   - 循环条件不存在，并且
   - 这个"`for`"语句没有使用`range`子句。
6. 一个 "`switch` "语句，其中：
   - 没有 "`break` "语句引用 "`switch` "语句
   - 有一个默认的分支，并且
   - 每个分支下的语句列表，包括默认分支，都以一个终止语句结束，或者是一个可能标有 ["fallthrough "的语句](#fallthrough-statements-fallthrough)。
7. 一个 "`select`"语句，其中:
   - 没有"`break`"语句引用"`select`"语句，并且
   - 每个分支下的语句列表，包括默认分支（如果存在），都以一个终止语句结束
8. 标记终止语句的[标签语句](#labeled-statements)。

​	所有其他语句都不是终止性的。

​	如果[语句列表](../Blocks)不是空的，并且其最后的非空语句是终止性的，则该列表以终止性语句结束。

### Empty statements 空语句

空语句不做任何事情。

```
EmptyStmt = .
```

### Labeled statements 标签语句

标签语句可以是`goto`、`break`或`continue`语句的目标。

```
LabeledStmt = Label ":" Statement .
Label       = identifier .
Error: log.Panic("error encountered")
```

### Expression statements 表达式语句

​	除了特定的内置函数外，函数和方法[调用](../Expressions#calls)以及[接收操作](../Expressions#receive-operator )可以出现在语句上下文中。这样的语句可以用圆括号括起来。

```
ExpressionStmt = Expression .
```

以下内置函数不允许出现在语句上下文中：

```go linenums="1"
append cap complex imag len make new real
unsafe.Add unsafe.Alignof unsafe.Offsetof unsafe.Sizeof unsafe.Slice
```

```go linenums="1"
h(x+y)
f.Close()
<-ch
(<-ch)
len("foo")  // illegal if len is the built-in function
```

### Send statements 发送语句

​	发送语句在通道上发送一个值。通道表达式的[核心类型](../PropertiesOfTypesAndValues#core-types)必须是一个[通道](../Types#channel-types)，通道方向必须允许发送操作，而且要发送的值的类型必须是[可以分配](../PropertiesOfTypesAndValues#assignability)给通道的元素类型。

```
SendStmt = Channel "<-" Expression .
Channel  = Expression .
```

​	在通信开始之前，通道和值表达式都被求值。通信阻塞，直到发送（操作）可以进行。如果有接收端准备好了，那么在一个没有缓冲的通道上的发送可以继续进行。在缓冲通道上的发送可以在缓冲区有空间的情况下进行。在关闭的通道上进行发送会引起[运行时恐慌](../Run-timePanics)。在`nil`通道上的发送会永远阻塞。

```go linenums="1"
ch <- 3  // send value 3 to channel ch
```

### IncDec statements  自增自减语句

​	`++`和`--`语句通过无类型[常量](../Constants)`1`来增加或减少其操作数。与赋值一样，操作数必须是[可寻址的](../Expressions#address-operators)，或者是一个映射索引表达式。

```
IncDecStmt = Expression ( "++" | "--" ) .
```

下面的[赋值语句](#assignment-statements)在语义上是等同的：

```
IncDec statement    Assignment
x++                 x += 1
x--                 x -= 1
```

### Assignment statements 赋值语句

​	赋值是用一个[表达式](../#Expressions)指定的新值来替换存储在[变量](../Variables)中的当前值。赋值语句可以为单个变量赋值，也可以将多个值赋给匹配数量的变量。

```
Assignment = ExpressionList assign_op ExpressionList .

assign_op = [ add_op | mul_op ] "=" .
```

​	每个左操作数必须是[可寻址的](../Expressions#address-operators)，或是一个映射索引表达式，或是（仅对`=`赋值）[空白标识符，即`_`](../DeclarationsAndScope#blank-identifier)。操作数可以用圆括号括起来。

```go linenums="1"
x = 1
*p = f()
a[i] = 23
(k) = <-ch  // same as: k = <-ch
```

​	赋值操作`x` op`=` `y`，其中`op`是一个[二元算术运算符](../Expressions#arithmetic-operators)，相当于`x` `=` `x` op `(y)`，但只对`x`进行一次求值。op`=`结构是一个单一的标记。在赋值操作中，左表达式和右表达式列表都必须正好包含一个单值表达式，并且左表达式不能是空白标识符。

```go linenums="1"
a[i] <<= 2
i &^= 1<<n
```

​	多元赋值将多值运算的各个元素分配给一个变量列表。有两种形式。在第一种形式中，右操作数是单个多值表达式，如一个函数调用、一个[通道](../Types#channel-types)或[映射](../Types#map-types)操作，或一个[类型断言](../Expressions#type-assertions)。左操作数必须与值的数量相匹配。例如，如果`f`是一个返回两个值的函数，

```go linenums="1"
x, y = f()
```

将第一个值赋给 `x`，将第二个值赋给 `y`。在第二种形式中，左操作数必须等于右表达式数量，每个表达式必须是单值的，右边的第`n`个表达式被分配给左边的第`n`个操作数：

```go linenums="1"
one, two, three = '一', '二', '三'
```

[空白标识符，即`_`](../DeclarationsAndScope#blank-identifier)提供了一种在赋值中忽略右值的方法：

```go linenums="1"
_ = x       // evaluate x but ignore it => 对 x 求值，但忽略它
x, _ = f()  // evaluate f() but ignore second result value => 对 f() 求值，但忽略它的第二个结果值
```

​	赋值分两个阶段进行。第一阶段，左边的[索引表达式](../Expressions#index-expressions)和[指针间接](../Expressions#address-operators)（包括[选择器](../Expressions#selectors)中的隐式指针间接）的操作数以及右边的表达式都按照[通常的顺序被求值](../Expressions#order-of-evaluation)。第二阶段，赋值是按照从左到右的顺序进行的。

```go linenums="1"
a, b = b, a  // exchange a and b => 交换 a 和 b

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1

x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end) => 设 x[0] = 1, 然后设 x[0] = 2 （故最后 x[0] == 2）

x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5. => 设 x[1] = 4, 然后设 x[3] = 5 时 panic

type Point struct { x, y int }
var p *Point
x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7 => 设 x[2] = 6, 然后设 p.x = 7 时 panic

i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // set i, x[2] = 0, x[0]
	break
}
// after this loop, i == 0 and x == []int{3, 5, 3}
```

​	在赋值中，每个值必须[可以赋给](../PropertiesOfTypesAndValues#assignability)它所赋的操作数的类型，但有以下特殊情况：

1. 任何类型的值都可以被分配给`空白标识符`。
2. 如果无类型的常量被分配给接口类型的变量或空白标识符，那么该常量首先被隐式地[转换](../Expressions#conversions)为其[默认类型](../Constants)。
3. 如果无类型的布尔值被分配给接口类型的变量或空白标识符，它首先被隐式转换为`bool`。

### If statements - if 语句

​	"`if` "语句根据布尔表达式的值指定两个分支的条件性执行。如果表达式的值为真，则执行 "`if` "分支，否则，执行 "`else` "分支（如果存在）。

```
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
if x > max {
	x = max
}
```

表达式前面可以有一个简单的语句，它在表达式被求值之前执行。

```go linenums="1"
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}
```

### Switch statements  - switch 语句

​	"`switch` "语句提供多路执行。表达式或类型与 "switch "内部的 "case "进行比较，以确定执行哪个分支。

```
SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .
```

​	有两种形式：`表达式开关`和`类型开关`。在`表达式开关`中，case 包含与开关表达式的值进行比较的表达式。在`类型开关`中，case 包含类型，这些类型与特别说明的表达式开关的类型进行比较。在`switch语句`中，`switch 表达式`被精确地求值一次。

#### Expression switches 表达式开关

​	在表达式开关中，`switch 表达式`被求值，case表达式（不需要是常量）被从左到右和从上到下求值；第一个等于switch表达式的表达式会触发相关case语句的执行；其他case被跳过。如果没有匹配的case，但有一个 "default" case，那么这个语句将被执行。"default" case 最多只能有一个，它可以出现在 "switch "语句的任何地方。缺省的`switch 表达式`等同于布尔值`true`。

```
ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
ExprCaseClause = ExprSwitchCase ":" StatementList .
ExprSwitchCase = "case" ExpressionList | "default" .
```

​	如果`switch 表达式`求值为无类型的常量，它首先被隐式地[转换](../Expressions#conversions)为其[默认类型](../Constants)。预先声明的无类型值`nil`不能作为`switch 表达式`使用。`switch 表达式`的类型必须是[可比较的](../Expressions#comparison-operators)。

​	如果一个case 表达式是无类型的，它首先被隐式地[转换](../Expressions#conversions)为`switch 表达式`的类型。对于每个（可能转换过的）case 表达式`x`和`switch 表达式`的值`t`，`x == t`必须是一个有效的[比较](../Expressions#comparison-operators)。

​	换句话说，`switch 表达式`被当作是用来声明和初始化一个没有明确类型的临时变量`t`；每个case表达式`x`都是用`t`的值来测试是否相等。

​	在case或default子句中，最后一个非空语句可以是一个（可能被[标记的](#labeled-statements)）["fallthrough "语句](#fallthrough-statements)，表示控制应该从这个子句的结尾流向下一个子句的第一个语句。否则，控制将流向 "switch "语句的末尾。"fallthrough"语句可以作为`表达式开关`中除最后一个子句之外的所有子句的最后一个语句出现。

​	`switch 表达式`前面可以有一个简单的语句，它在表达式被求值之前执行。

```go linenums="1"
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // missing switch expression means "true" => 缺少 switch 表达式 意味着 "true"
case x < 0: return -x
default: return x
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}
```

实现限制：编译器可能不允许多个case 表达式求值为同一个常量。例如，目前的编译器不允许在case表达式中出现重复的整型常量、浮点常量或字符串常量。

#### Type switches 类型开关

​	`类型开关`比较的是`类型`而不是值。它与`表达式开关`类似。它由一个特殊的`switch 表达式`标记，该表达式具有[类型断言](../Expressions#type-assertions)的形式，使用关键字`type`而不是实际的类型：

```go linenums="1"
switch x.(type) {
// cases
}
```

​	然后，case 将实际类型`T`与表达式`x`的动态类型相匹配。与类型断言一样，`x`必须是[接口类型](../Types#interface-types)，但不是[类型参数](../DeclarationsAndScope#type-parameter-declarations)，而且case 中列出的每个非接口类型`T`必须实现`x`的类型。在类型开关的case 中，列出的类型都必须是[不同的](../PropertiesOfTypesAndValues#type-identity)。

```
TypeSwitchStmt  = "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClause } "}" .
TypeSwitchGuard = [ identifier ":=" ] PrimaryExpr "." "(" "type" ")" .
TypeCaseClause  = TypeSwitchCase ":" StatementList .
TypeSwitchCase  = "case" TypeList | "default" .
```

​	TypeSwitchGuard 可能包括一个[短变量声明](../DeclarationsAndScope#short-variable-declarations)。当使用这种形式时，该变量在每个子句的[隐含块](../Blocks)中的TypeSwitchCase的末尾被声明。在子句中，如果case正好列出了一种类型，那么变量就有这种类型；否则，变量就有TypeSwitchGuard中表达式的类型。

​	case 可以使用预先声明的标识符[nil](../DeclarationsAndScope#predeclared-identifiers)来代替类型；当TypeSwitchGuard中的表达式是一个`nil`接口值时，该case被选中。最多只能有一个`nil `case。

​	给定一个`interface{}`类型的表达式`x`，下面的类型开关：

```go linenums="1"
switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{}) => i 类型是 x 的类型（interface{}）
case int:
	printInt(i)                            // type of i is int => i 类型是 int
case float64:
	printFloat64(i)                        // type of i is float64 => i 类型是 float64 
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64 => i 类型是 func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{}) => i 类型是 x 的类型（interface{}）
default:
	printString("don't know the type")     // type of i is type of x (interface{}) => i 类型是 x 的类型（interface{}）
}
```

可以重写：

```go linenums="1"
v := x  // x is evaluated exactly once => x 只被求值一次
if v == nil {
	i := v                                 // type of i is type of x (interface{})
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i)                            // type of i is int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)                        // type of i is float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)                       // type of i is func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // type of i is type of x (interface{})
		printString("type is bool or string")
	} else {
		i := v                         // type of i is type of x (interface{})
		printString("don't know the type")
	}
}
```

​	[类型参数](../DeclarationsAndScope#type-parameter-declarations)或[泛型](../DeclarationsAndScope#type-declarations)可以作为 case 中的一个类型。如果在[实例化](../Expressions#instantiations)时，该类型被发现与开关中的另一个条目重复，则选择第一个匹配的case。

```go linenums="1" hl_lines="7 7"
func f[P any](x any) int {
	switch x.(type) {
	case P:
		return 0
	case string:
		return 1
	case []P:
		return 2
	case []byte:
		return 3
	default:
		return 4
	}
}

var v1 = f[string]("foo")   // v1 == 0
var v2 = f[byte]([]byte{})  // v2 == 2
```

​	类型开关防护（guard ）前可以有一个简单的语句，该语句在防护（guard ）被求值前执行。

​	在类型开关中`不允许使用` "fallthrough "语句。

### For statements  - for 语句

​	 "`for` "语句指定重复执行一个块。有三种形式：迭代可以由单个条件、"for "子句或 "range "子句控制。

```
ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
Condition = Expression .
```

#### For statements with single condition 带有单一条件的for语句

​	在其最简单的形式中，"for "语句指定重复执行一个块，只要一个布尔条件被求值为真。该条件在每次迭代前被求值。如果条件不存在，它就等同于布尔值`true`。

```go linenums="1"
for a < b {
	a *= 2
}
```

#### For statements with `for` clause 带有for子句的for语句

​	带有for子句的 "for "语句也受其条件的控制，但另外它可以指定一个`init`和一个`post`语句，如一个赋值，一个增量或减量语句。init语句可以是一个[短变量声明](../DeclarationsAndScope#short-variable-declarations)，但post语句则不能。由init语句声明的变量会在每次迭代中重复使用。

```
ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
InitStmt = SimpleStmt .
PostStmt = SimpleStmt .
for i := 0; i < 10; i++ {
	f(i)
}
```

​	如果（init语句）非空，则它在求值第一个迭代的条件之前被执行一次；post语句在每次执行块之后被执行（而且只有当该块被执行时）。For子句的任何元素都可以是空的，但是（除非只有一个条件）[分号](../LexicalElements#semicolons)是必须的。如果没有条件，则等同于布尔值`true`。=> 仍有疑问？？

```
for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }
```

#### For statements with `range` clause  带有range子句的for语句

​	带有 "`range`"子句的 "for "语句遍历一个数组、切片、字符串或映射的所有条目，或在一个通道上收到的值。对于每个条目，如果存在的话，它将迭代值分配给相应的迭代变量，然后执行该块。

```
RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
```

​	在 "`range`"子句中`右边的表达式`称为`范围表达式`，其[核心类型](../PropertiesOfTypesAndValues#core-types)必须是数组、指向数组的指针、切片、字符串、映射或允许[接收操作](../Expressions#receive-operator )的通道。和赋值一样，如果左操作数存在的话，那么它必须是[可寻址的](../Expressions#address-operators)或映射索引表达式；它们（即左操作数）表示`迭代变量`。如果`range 表达式`是一个通道，最多允许一个迭代变量，其他情况下最多可以有两个迭代变量。如果最后一个迭代变量是[空白标识符，即`_`](../DeclarationsAndScope#blank-identifier)，那么range 子句就等同于没有该空白标识符的相同子句。

​	在开始循环之前，range 表达式`x`被求值一次，**但有一个例外**：如果最多只有一个迭代变量，并且`len(x)`是[常量](../Constants)，那么range 表达式不被求值。

Function calls on the left are evaluated once per iteration. For each iteration, iteration values are produced as follows if the respective iteration variables are present:

​	左边的函数调用在每个迭代中被求值一次（=>仍有疑问？？）。对于每个迭代，如果各自的迭代变量存在，则迭代值按以下产生：

```
Range expression                          1st value          2nd value
range 表达式                               第一个值           第二

array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, <-chan E       element  e  E
```

1. 对于一个数组值、数组指针值或切片值`a`，索引迭代值按递增顺序产生，从元素索引0开始。如果最多只有一个迭代变量，range 循环产生从`0`到`len(a)-1`的迭代值，并且不对数组或切片本身进行索引。对于一个`nil` 切片，迭代次数为0。
2. 对于一个字符串值，"range "子句在字符串中的Unicode码点上进行迭代，从字节索引0开始。在连续的迭代中，索引值将是字符串中连续的UTF-8编码码点的第一个字节的索引，第二个值，类型为`rune`，将是相应码点的值。`如果迭代遇到一个无效的UTF-8序列`，第二个值将是（Unicode替换字符）`0xFFFD`，下一次迭代将在字符串中推进一个字节。
3. 对映射的迭代顺序没有指定，不保证每次迭代都是一样的。如果在迭代过程中删除了一个尚未到达的映射条目，将不会产生相应的迭代值。如果在迭代过程中创建了一个映射条目，该条目可能在迭代过程中被产生，也可能被跳过。对于每个创建的条目，以及从一个迭代到另一个迭代，选择可能有所不同。如果映射为`nil`，迭代次数为0。
4. 对于通道，产生的迭代值是通道上连续发送的值，直到通道[关闭](../Built-inFunctions#close)。如果通道为`nil`，则 range 表达式永远阻塞。

​	迭代值被分配给各自的迭代变量，就像在[赋值语句](#assignment-statements)中一样。

​	迭代变量可以由 "range "子句使用[短变量声明](../DeclarationsAndScope#short-variable-declarations)的形式（`:=`）来声明。在这种情况下，它们的类型被设置为各自的迭代值的类型，它们的[作用域](../DeclarationsAndScope)是 "for "语句的块；它们在每个迭代中被重复使用。如果迭代变量是在 "for "语句之外声明的，执行后它们的值将是最后一次迭代的值。

```go linenums="1"
var testdata *struct {
	a *[7]int
}
for i, _ := range testdata.a {
	// testdata.a is never evaluated; len(testdata.a) is constant => testdata.a 不会被求值； len(testdata.a) 是常量 <= 这里符合前面说的那个“但有一个例外”的情况
	// i ranges from 0 to 6 => i 范围从 0 到 6
	f(i)
}

var a [10]string
for i, s := range a {
	// type of i is int
	// type of s is string
	// s == a[i]
	g(i, s)
}

var key string
var val interface{}  // element type of m is assignable to val => m 的元素类型可赋值给 val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
	h(key, val)
}
// key == last map key encountered in iteration
// val == map[key]

var ch chan Work = producer()
for w := range ch {
	doWork(w)
}

// empty a channel => 清空通道
for range ch {}
```

### Go statements  - go 语句

​	`go`语句作为一个独立的并发控制线程（或称为goroutine），`在同一地址空间内`开始执行一个函数调用。

```
GoStmt = "go" Expression .
```

​	（go语句中的）表达式`必须是一个函数或方法调用`；其不能用圆括号括起来。对内置函数的调用与[表达式语句](#expression-statements)一样受到限制。

​	函数值和参数在调用的goroutine中[像往常一样被求值](../Expressions#order-of-evaluation)，但与普通调用不同的是，程序执行不会等待被调用的函数完成。相反，该函数开始在一个新的goroutine中独立执行。当函数终止时，其goroutine也会终止。如果该函数有任何返回值，当函数完成时，它们会被丢弃。

```go linenums="1"
go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true }} (c)
```

### Select statements - select 语句

​	 "`select` "语句选择一组可能的[发送](#send-statements)或[接收](../Expressions#receive-operator)操作中的一个来进行。它看起来类似于 "[switch](#switch-statements) "语句，但其 case 都是只涉及通信操作。

```
SelectStmt = "select" "{" { CommClause } "}" .
CommClause = CommCase ":" StatementList .
CommCase   = "case" ( SendStmt | RecvStmt ) | "default" .
RecvStmt   = [ ExpressionList "=" | IdentifierList ":=" ] RecvExpr .
RecvExpr   = Expression .
```

​	一个有 RecvStmt 的 case 可以将 RecvExpr 的结果分配给一个或两个变量，这些变量可以用[短变量声明](../DeclarationsAndScope#short-variable-declarations)来声明。RecvExpr 必须是一个（可能是圆括号内的）接收操作。最多可以有一个default case，它可以出现在 case 列表的任何地方。

​	 "`select` "语句的执行分几个步骤进行：

1. 对于语句中的所有情况，在进入 "`select`"语句时，接收操作的通道操作数、通道、发送语句的右侧表达式（按源代码出现的顺序）被求值一次。结果是一组要接收或发送的通道，以及对应的要发送的值。无论选择哪一个（如果有）通信操作来进行，在这次求值中的任何副作用都会发生。在 RecvStmt 左侧的表达式有一个短变量声明或赋值，还没有被求值。
2. 如果一个或多个通信可以进行，则通过统一的伪随机选择来选择一个可以进行的通信。否则，如果有一个default case，就会选择该情况。如果没有default case，"`select`"语句就会阻塞，直到至少有一项通信可以进行。
3. 除非选择的情况是default case，否则将执行相应的通信操作。
4. 如果选择的 case 是一个带有短变量声明或赋值的RecvStmt，左边的表达式会被求值，接收到的值（或多个值）被用于赋值。
5. 所选 case 语句列表被执行。

​	由于在`nil`通道上的通信永远不能进行，所以只有`nil`通道且没有default case 的`select`语句会永远阻塞。

```go linenums="1"
var a []int
var c, c1, c2, c3, c4 chan int
var i1, i2 int
select {
case i1 = <-c1:
	print("received ", i1, " from c1\n")
case c2 <- i2:
	print("sent ", i2, " to c2\n")
case i3, ok := (<-c3):  // same as: i3, ok := <-c3
	if ok {
		print("received ", i3, " from c3\n")
	} else {
		print("c3 is closed\n")
	}
case a[f()] = <-c4:
	// same as:
	// case t := <-c4
	//	a[f()] = t
default:
	print("no communication\n")
}

for {  // send random sequence of bits to c
	select {
	case c <- 0:  // note: no statement, no fallthrough, no folding of cases
	case c <- 1:
	}
}

select {}  // block forever
```

### Return statements  - return 语句

​	函数`F`中的 "`return`"语句终止了`F`的执行，并且可以选择提供一个或多个结果值。在`F`返回给它的调用者之前，任何由`F`[延迟](#defer-statements)的函数都会被执行。

```
ReturnStmt = "return" [ ExpressionList ] .
```

​	在没有结果类型的函数中，"`return`"语句必须不指定任何结果值。

```go linenums="1"
func noResult() {
	return
}
```

有三种方法可以从一个有结果类型的函数中返回值：

1. 可以在 "`return`"语句中明确列出一个或多个返回值。每个表达式必须是单值的，并且可以分配给函数的结果类型的相应元素。

   ```go linenums="1"
   func simpleF() int {
   	return 2
   }
   
   func complexF1() (re float64, im float64) {
   	return -7.0, -4.0
   }
   ```
   
2. "`return`"语句中的表达式列表可能是对一个多值函数的单一调用。其效果就像从该函数返回的每个值都被分配到一个临时变量中，其类型为相应的值，随后的 "`return`"语句列出了这些变量，此时，前一种情况的规则适用。

   ```go linenums="1"
   func complexF2() (re float64, im float64) {
   	return complexF1()
   }
   ```
   
3. 如果函数的结果类型为其结果参数指定了名称，表达式列表可能是空的。结果参数作为普通的局部变量，函数可以根据需要给它们赋值。`return`语句会返回这些变量的值。

   ```go linenums="1"
   func complexF3() (re float64, im float64) {
   	re = 7.0
   	im = 4.0
   	return
   }
   
   func (devnull) Write(p []byte) (n int, _ error) {
   	n = len(p)
   	return
   }
   ```

​	不管它们是如何被声明的，`所有的结果值在进入函数时都被初始化为`其类型的[零值](../ProgramInitializationAndExecution#the-zero-value)。指定结果的 "`return`"语句`在执行任何延迟函数之前`设置结果参数。

实现限制：如果在返回的地方有一个与结果参数同名的不同实体（常量、类型或变量）在[作用域](../DeclarationsAndScope)内，编译器可能不允许在 "`return`"语句中出现空表达式列表。

```go linenums="1"
func f(n int) (res int, err error) {
	if _, err := f(n-1); err != nil {
		return  // invalid return statement: err is shadowed => 无效的返回语句： err 被遮蔽了
	}
	return
}
```

### Break statements - break 语句

​	 "`break` "语句可以终止同一函数中最里面的 "[for](#for-statements-for)"、"[switch](#switch-statements-switch) "或 "[select](#select-statements-select) "语句的执行。

```
BreakStmt = "break" [ Label ] .
```

​	如果有一个标签，它必须是一个封闭的 "`for`"、"`switch` "或 "`select` "语句的标签，而且是执行终止的那一个。

```go linenums="1"
OuterLoop:
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			switch a[i][j] {
			case nil:
				state = Error
				break OuterLoop
			case item:
				state = Found
				break OuterLoop
			}
		}
	}
```

### Continue statements  - continue 语句

​	 "`continue` "语句通过将控制推进到循环块的末端来开始最内层的 ["for "循环](#for-statements-for)的下一次迭代。"`for` "循环必须是在同一个函数中。

```
ContinueStmt = "continue" [ Label ] .
```

​	如果有一个标签，它必须是一个封闭的 "`for` "语句的标签，而且是执行前进的那一个。

```go linenums="1"
RowLoop:
	for y, row := range rows {
		for x, data := range row {
			if data == endOfRow {
				continue RowLoop
			}
			row[x] = data + bias(x, y)
		}
	}
```

### Goto statements 语句 goto

​	 "`goto` "语句将控制转移到同一函数中具有相应标签的语句。

```
GotoStmt = "goto" Label .
goto Error
```

​	执行 "`goto` "语句不能导致任何变量进入goto处的作用域之外的[作用域](../DeclarationsAndScope)。例如，这个例子：

```go linenums="1"
	goto L  // BAD
	v := 3
L:
```

是错误的，因为跳转到标签`L`时，跳过了创建`v`的过程。

​	[块](../Blocks)外的 "`goto` "语句不能跳到该块内的标签。例如，这个例子：

```go linenums="1"
if n%2 == 1 {
	goto L1
}
for n > 0 {
	f()
	n--
L1:
	f()
	n--
}
```

是错误的，因为标签`L1`在 "for "语句的块内，而`goto`不在其中。

### Fallthrough statements 语句 fallthrough

​	在[表达式开关语句](#expression-switches)中，"`fallthrough` "语句将控制转移到下一个`case`子句的第一个语句。它只能作为这种子句中的最后一个非空语句使用。

```
FallthroughStmt = "fallthrough" .
```

### Defer statements 语句 defer

​	"`defer`"语句调用一个函数，该函数的执行被推迟到外层函数返回的那一刻，或者是因为外层函数执行了 [return 语句](#return-statements-return)，达到了其[函数体](../DeclarationsAndScope#function-declarations)的末端，或者是因为相应的goroutine正在[恐慌](../Built-inFunctions#handling-panics)。

```
DeferStmt = "defer" Expression .
```

这个表达式必须是一个函数或方法的调用；它不能是被圆括号括起来的。对内置函数的调用与[表达式语句](#expression-statements)一样受到限制。

​	每次执行 "`defer` "语句时，函数值和调用的参数[像往常一样被求值](../Expressions#order-of-evaluation)并重新保存，但实际的函数不会被调用。相反，`被延迟函数`在外层的函数返回之前立即被调用，`其顺序与它们被延迟的顺序相反`。也就是说，如果外层的函数通过一个明确的 [return 语句](#return-statements-return)返回，`被延迟函数`在任何结果参数被该 `reurn 语句`设置后执行，`但在外层函数返回给其调用者之前`。如果`被延迟函数`值被求值为`nil`，那么在调用该被延迟函数时会出现执行[恐慌]()（而不是当 "`defer` "语句被执行时）。

​	例如，如果`被延迟函数`是一个[函数字面量](../Expressions#function-literals)，并且外层的函数有在该字面量的作用域内的[命名结果参数](../Types#function-types)，那么该`被延迟函数`可以在这些结果参数被返回之前访问和修改它们。如果`被延迟函数`有任何返回值，这些返回值将在函数完成时被丢弃。(参见[处理恐慌](../Built-inFunctions#handling-panics)一节)。

```go linenums="1"
lock(l)
defer unlock(l)  // unlocking happens before surrounding function returns => unlock 发生在外层函数返回之前

// prints 3 2 1 0 before surrounding function returns
// => 在外层函数返回之前，打印 3 2 1 0
for i := 0; i <= 3; i++ {
	defer fmt.Print(i)
}

// f returns 42
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement => 结果会在它被 return 语句设为 6 之后再被访问
		result *= 7
	}()
	return 6
}
```

