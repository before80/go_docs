+++
title = "语句"
date = 2023-05-17T09:59:21+08:00
weight = 12
description = ""
isCJKLanguage = true
type = "docs"
draft = false

+++
## Statements 语句

> 原文：[https://go.dev/ref/spec#Statements ](https://go.dev/ref/spec#Statements )

Statements control execution.

​	语句控制执行。

```
Statement =
	Declaration | LabeledStmt | SimpleStmt |
	GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
	FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
	DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
```

### Terminating statements 终止语句

A *terminating statement* interrupts the regular flow of control in a [block](https://go.dev/ref/spec#Blocks). The following statements are terminating:

​	终止语句中断了一个[块](../Blocks)中的常规控制流。下列语句是终止性的：

1. A ["return"](https://go.dev/ref/spec#Return_statements) or ["goto"](https://go.dev/ref/spec#Goto_statements) statement.
2. "`return`"或 "`goto`"语句。
3. A call to the built-in function [`panic`](https://go.dev/ref/spec#Handling_panics).
4. 对内置函数`panic`的调用。
5. A [block](https://go.dev/ref/spec#Blocks) in which the statement list ends in a terminating statement.
6. 语句列表以终止语句结束的块。
7. An ["if" statement](https://go.dev/ref/spec#If_statements) in which:
8. 一个 "`if`"语句，其中：
   - the "else" branch is present, and
   - 存在 "`else`"分支，并且
   - both branches are terminating statements.
   - 两个分支都是终止语句。
9. A ["for" statement](https://go.dev/ref/spec#For_statements) in which:
10. 一个 "`for`"语句，其中：
   - there are no "break" statements referring to the "for" statement, and
   - 没有 "`break`"语句引用"`for`"语句，并且
   - the loop condition is absent, and
   - 循环条件不存在，并且
   - the "for" statement does not use a range clause.
   - 这个"`for`"语句没有使用`range`子句。
11. A ["switch" statement](https://go.dev/ref/spec#Switch_statements) in which:
12. 一个 "`switch`"语句，其中：
    - there are no "break" statements referring to the "switch" statement,
    - 没有 "`break`"语句引用 "`switch`"语句
    - there is a default case, and
    - 有一个默认的分支，并且
    - the statement lists in each case, including the default, end in a terminating statement, or a possibly labeled ["fallthrough" statement](https://go.dev/ref/spec#Fallthrough_statements).
    - 每个分支下的语句列表，包括默认分支，都以一个终止语句结束，或者是一个可能标有 ["fallthrough"的语句](#fallthrough-statements-fallthrough-语句-fallthrough)。
13. A ["select" statement](https://go.dev/ref/spec#Select_statements) in which:
14. 一个 "`select`"语句，其中:
    - there are no "break" statements referring to the "select" statement, and
    - 没有"`break`"语句引用"`select`"语句，并且
    - the statement lists in each case, including the default if present, end in a terminating statement.
    - 每个分支下的语句列表，包括默认分支（如果存在），都以一个终止语句结束
15. A [labeled statement](https://go.dev/ref/spec#Labeled_statements) labeling a terminating statement.
16. 标记终止语句的[标签语句](#labeled-statements-标签语句)。

​	

All other statements are not terminating.

​	所有其他语句都不是终止性的。

A [statement list](https://go.dev/ref/spec#Blocks) ends in a terminating statement if the list is not empty and its final non-empty statement is terminating.

​	如果[语句列表](../Blocks)不是空的，并且其最后的非空语句是终止性的，则该列表以终止性语句结束。

### Empty statements 空语句

The empty statement does nothing.

​	空语句不做任何事情。

```
EmptyStmt = .
```

### Labeled statements 标签语句

A labeled statement may be the target of a `goto`, `break` or `continue` statement.

​	标签语句可以是`goto`、`break`或`continue`语句的目标。

```
LabeledStmt = Label ":" Statement .
Label       = identifier .
```

```go
Error: log.Panic("error encountered")
```



### Expression statements 表达式语句

With the exception of specific built-in functions, function and method [calls](https://go.dev/ref/spec#Calls) and [receive operations](https://go.dev/ref/spec#Receive_operator) can appear in statement context. Such statements may be parenthesized.

​	除了特定的内置函数外，函数和方法[调用](../Expressions#calls-调用)以及[接收操作](../Expressions#receive-operator-接收操作符)可以出现在语句上下文中。这样的语句可以用圆括号括起来。

```
ExpressionStmt = Expression .
```

The following built-in functions are not permitted in statement context:

​	以下内置函数不允许出现在语句上下文中：

```go 
append cap complex imag len make new real
unsafe.Add unsafe.Alignof unsafe.Offsetof unsafe.Sizeof unsafe.Slice
```

```go 
h(x+y)
f.Close()
<-ch
(<-ch)
len("foo")  // illegal if len is the built-in function
```

### Send statements 发送语句

A send statement sends a value on a channel. The channel expression's [core type](https://go.dev/ref/spec#Core_types) must be a [channel](https://go.dev/ref/spec#Channel_types), the channel direction must permit send operations, and the type of the value to be sent must be [assignable](https://go.dev/ref/spec#Assignability) to the channel's element type.

​	发送语句在通道上发送一个值。通道表达式的[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)必须是一个[通道](../Types#channel-types-通道型)，通道方向必须允许发送操作，而且要发送的值的类型必须是[可以分配](../PropertiesOfTypesAndValues#assignability-可分配性)给通道的元素类型。

```
SendStmt = Channel "<-" Expression .
Channel  = Expression .
```

Both the channel and the value expression are evaluated before communication begins. Communication blocks until the send can proceed. A send on an unbuffered channel can proceed if a receiver is ready. A send on a buffered channel can proceed if there is room in the buffer. A send on a closed channel proceeds by causing a [run-time panic](https://go.dev/ref/spec#Run_time_panics). A send on a `nil` channel blocks forever.

​	在通信开始之前，通道和值表达式都被求值。通信阻塞，直到发送（操作）可以进行。如果有接收端准备好了，那么在一个没有缓冲的通道上的发送可以继续进行。在缓冲通道上的发送可以在缓冲区有空间的情况下进行。在关闭的通道上进行发送会引起[运行时恐慌](../Run-timePanics)。在`nil`通道上的发送会永远阻塞。

```go 
ch <- 3  // send value 3 to channel ch
```

### IncDec statements  自增自减语句

The "++" and "--" statements increment or decrement their operands by the untyped [constant](https://go.dev/ref/spec#Constants) `1`. As with an assignment, the operand must be [addressable](https://go.dev/ref/spec#Address_operators) or a map index expression.

​	`++`和`--`语句通过无类型[常量](../Constants)`1`来增加或减少其操作数。与赋值一样，操作数必须是[可寻址的](../Expressions#address-operators-地址运算符)，或者是一个映射索引表达式。

> 个人注释
>
> ​	以上这句话，难道是说映射的索引表达式不可以进行寻址？验证如下：
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	m := make(map[string]int)
> 	m["age"] = 32
> 	fmt.Printf("%#v,len=%d\n", m, len(m))       // map[string]int{"age":32},len=1
> 	//fmt.Printf("%#v,%p\n", m["age"], &m["age"]) // invalid operation: cannot take address of m["age"] (map index expression of type int)
> }
> 
> ```
>
> ​	确实：映射的索引表达式不可以进行寻址！

```
IncDecStmt = Expression ( "++" | "--" ) .
```

The following [assignment statements](https://go.dev/ref/spec#Assignment_statements) are semantically equivalent:

​	下面的[赋值语句](#assignment-statements-赋值语句)在语义上是等同的：

```go
IncDec statement    Assignment
x++                 x += 1
x--                 x -= 1
```

> 个人注释
>
> ​	`浮点数`、`复数`可以进行自增和自减吗？=》可以，但要注意精度问题！
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	f1 := 1.2
> 	fmt.Printf("%T,%v\n", f1, f1) // float64,1.2
> 	f1++
> 	fmt.Printf("%T,%v\n", f1, f1) // float64,2.2
> 	f1--
> 	fmt.Printf("%T,%v\n", f1, f1) // float64,1.2000000000000002
> 
> 	f2 := float32(1.2)
> 	fmt.Printf("%T,%v\n", f2, f2) // float32,1.2
> 	f2++
> 	fmt.Printf("%T,%v\n", f2, f2) // float32,2.2
> 	f2--
> 	fmt.Printf("%T,%v\n", f2, f2) // float32,1.2
> 
> 	c1 := complex(1.2, 3.4)
> 	fmt.Printf("%T,%v\n", c1, c1) // complex128,(1.2+3.4i)
> 	c1++
> 	fmt.Printf("%T,%v\n", c1, c1) // complex128,(2.2+3.4i)
> 	c1--
> 	fmt.Printf("%T,%v\n", c1, c1) // complex128,(1.2000000000000002+3.4i)
> 
> 	c2 := complex(float32(1.2), float32(3.4))
> 	fmt.Printf("%T,%v\n", c2, c2) // complex64,(1.2+3.4i)
> 	c2++
> 	fmt.Printf("%T,%v\n", c2, c2) // complex64,(2.2+3.4i)
> 	c2--
> 	fmt.Printf("%T,%v\n", c2, c2) // complex64,(1.2+3.4i)
> }
> 
> ```
>
> 

### Assignment statements 赋值语句

An *assignment* replaces the current value stored in a [variable](https://go.dev/ref/spec#Variables) with a new value specified by an [expression](https://go.dev/ref/spec#Expressions). An assignment statement may assign a single value to a single variable, or multiple values to a matching number of variables.

​	赋值是用一个[表达式](../Expressions)指定的新值来替换存储在[变量](../Variables)中的当前值。赋值语句可以为单个变量赋值，也可以将多个值赋给匹配数量的变量。

```
Assignment = ExpressionList assign_op ExpressionList .

assign_op = [ add_op | mul_op ] "=" .
```

Each left-hand side operand must be [addressable](https://go.dev/ref/spec#Address_operators), a map index expression, or (for `=` assignments only) the [blank identifier](https://go.dev/ref/spec#Blank_identifier). Operands may be parenthesized.

​	每个左操作数必须是[可寻址的](../Expressions#address-operators-地址运算符)，或是一个映射索引表达式，或是（仅用于`=`赋值）[空白标识符，即`_`](../DeclarationsAndScope#blank-identifierr-空白标识符)。操作数可以用圆括号括起来。

```go 
x = 1
*p = f()
a[i] = 23
(k) = <-ch  // 等同于：k = <-ch
```

An *assignment operation* `x` *op*`=` `y` where *op* is a binary [arithmetic operator](https://go.dev/ref/spec#Arithmetic_operators) is equivalent to `x` `=` `x` *op* `(y)` but evaluates `x` only once. The *op*`=` construct is a single token. In assignment operations, both the left- and right-hand expression lists must contain exactly one single-valued expression, and the left-hand expression must not be the blank identifier.

​	赋值操作`x = x op y`，其中`op`是一个[二元算术运算符](../Expressions#arithmetic-operators-算术运算符)，相当于`x = x op (y)`，但只对`x`进行一次求值。`op=`结构是一个单独的token。（`op=`）在赋值操作中，左侧的表达式和右侧的表达式列表都必须正好包含一个单值表达式，并且左侧的表达式不能是空白标识符。

> 个人注释
>
> ​	只是对于`op=`的左侧的表达式不能是`_`：
>
> ```go
> package main
> 
> func main() {
> 	_ = 1
> 	_ = 2
> 	_ += 1 // cannot use _ as value or type
> }
> ```
>
> ​	`op=` 或 `=`的右侧的表达式可以是`_`吗？=> 不可以！
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	f := 1.2
> 	fmt.Printf("%T,%#v\n", f, f) // float64,1.2
> 	//f = _   // cannot use _ as value or type
> 	//f += _ // cannot use _ as value or type
> 	fmt.Printf("%T,%#v\n", f, f) // float64,1.2
> 
> 	i := 1
> 	fmt.Printf("%T,%#v\n", i, i) // int,1
> 	//i = _   // cannot use _ as value or type
> 	//i += _ // cannot use _ as value or type
> 	fmt.Printf("%T,%#v\n", i, i) // int,1
> 
> 	c := complex(1.2, 3.4)
> 	fmt.Printf("%T,%#v\n", c, c) // complex128,(1.2+3.4i)
> 	//c = _   // cannot use _ as value or type
> 	//c += _  // cannot use _ as value or type
> 	fmt.Printf("%T,%#v\n", c, c) // complex128,(1.2+3.4i)
> }
> 
> ```

```go 
a[i] <<= 2
i &^= 1<<n
```

A tuple assignment assigns the individual elements of a multi-valued operation to a list of variables. There are two forms. In the first, the right hand operand is a single multi-valued expression such as a function call, a [channel](https://go.dev/ref/spec#Channel_types) or [map](https://go.dev/ref/spec#Map_types) operation, or a [type assertion](https://go.dev/ref/spec#Type_assertions). The number of operands on the left hand side must match the number of values. For instance, if `f` is a function returning two values,

​	多元赋值将多值运算的各个元素分配给一个变量列表。有两种形式。在第一种形式中，右操作数是单个多值表达式，如一个函数调用、一个[通道](../Types#channel-types-通道型)或[映射](../Types#map-types-映射型)操作，或一个[类型断言](../Expressions#type-assertions-类型断言)。左操作数必须与值的数量相匹配。例如，如果`f`是一个返回两个值的函数，

```go 
x, y = f()
```

assigns the first value to `x` and the second to `y`. In the second form, the number of operands on the left must equal the number of expressions on the right, each of which must be single-valued, and the *n*th expression on the right is assigned to the *n*th operand on the left:

将第一个值赋给 `x`，将第二个值赋给 `y`。在第二种形式中，左操作数必须等于右表达式数量，每个表达式必须是单值的，右边的第`n`个表达式被分配给左边的第`n`个操作数：

```go 
one, two, three = '一', '二', '三'
```

The [blank identifier](https://go.dev/ref/spec#Blank_identifier) provides a way to ignore right-hand side values in an assignment:

​	[空白标识符，即`_`](../DeclarationsAndScope#blank-identifierr-空白标识符)提供了一种在赋值中忽略右值的方法：

```go 
_ = x       // 对 x 求值，但忽略它
x, _ = f()  // 对 f() 求值，但忽略它的第二个结果值
```

The assignment proceeds in two phases. First, the operands of [index expressions](https://go.dev/ref/spec#Index_expressions) and [pointer indirections](https://go.dev/ref/spec#Address_operators) (including implicit pointer indirections in [selectors](https://go.dev/ref/spec#Selectors)) on the left and the expressions on the right are all [evaluated in the usual order](https://go.dev/ref/spec#Order_of_evaluation). Second, the assignments are carried out in left-to-right order.

​	赋值分两个阶段进行。第一阶段，左边的[索引表达式](../Expressions#index-expressions-索引表达式)和[指针间接](../Expressions#address-operators-地址运算符)（包括[选择器](../Expressions#selectors-选择器)中的隐式指针间接）的操作数以及右边的表达式都按照[通常的顺序被求值](../Expressions#order-of-evaluation-求值顺序)。第二阶段，赋值是按照从左到右的顺序进行的。

```go 
a, b = b, a  // 交换 a 和 b

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // 设置 i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // 设置 x[0] = 2, i = 1

x[0], x[0] = 1, 2  // 设置 x[0] = 1, 然后设置 x[0] = 2 （故最后 x[0] == 2）

x[1], x[3] = 4, 5  // 设置 x[1] = 4, 然后设置 x[3] = 5 时 panic <= 因x[3]索引越界，引发panic!

type Point struct { x, y int }
var p *Point
x[2], p.x = 6, 7  // 设置 x[2] = 6, 然后设置 p.x = 7 时 panic <= 因 p 是 nil 指针，可以将p 是初始化为： var p *Point = &Point{} 避免panic的发生！

i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // 设置 i, x[2] = 0, x[0]
	break
}
// 在这个循环之后, i == 0 and x == []int{3, 5, 3}
```

In assignments, each value must be [assignable](https://go.dev/ref/spec#Assignability) to the type of the operand to which it is assigned, with the following special cases:

​	在赋值中，每个值必须[可以赋给](../PropertiesOfTypesAndValues#assignability-可分配性)它所赋的操作数的类型，但有以下特殊情况：

1. Any typed value may be assigned to the blank identifier.
2. 任何类型的值都可以被分配给`空白标识符`。
3. If an untyped constant is assigned to a variable of interface type or the blank identifier, the constant is first implicitly [converted](https://go.dev/ref/spec#Conversions) to its [default type](https://go.dev/ref/spec#Constants).
4. 如果无类型的常量被分配给接口类型的变量或空白标识符，那么该常量首先被隐式地[转换](../Expressions#conversions-转换)为其[默认类型](../Constants)。
5. If an untyped boolean value is assigned to a variable of interface type or the blank identifier, it is first implicitly converted to type `bool`.
6. 如果无类型的布尔值被分配给接口类型的变量或空白标识符，它首先被隐式转换为`bool`。

### If statements - if 语句

"If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.

​	"`if`"语句根据布尔表达式的值指定两个分支的条件性执行。如果表达式的值为真，则执行 "`if`"分支，否则，执行 "`else`"分支（如果存在）。

```
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
if x > max {
	x = max
}
```

The expression may be preceded by a simple statement, which executes before the expression is evaluated.

​	表达式前面可以有一个简单的语句，它在表达式被求值之前执行。

```go 
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}
```

> 个人注释
>
> ​	简单的语句之后需要跟一个求值为bool类型的条件，否则报错：
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"math/rand"
> )
> 
> func main() {
> 	if y := rand.Intn(6); y { // 报错：non-boolean condition in if statement
> 		fmt.Printf("%T,%v", y, y)
> 	}
> }
> 
> ```
>
> 

### Switch statements  - switch 语句

"Switch" statements provide multi-way execution. An expression or type is compared to the "cases" inside the "switch" to determine which branch to execute.

​	"`switch`"语句提供多路执行。表达式或类型与 "switch"内部的 "case"进行比较，以确定执行哪个分支。

```
SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .
```

There are two forms: expression switches and type switches. In an expression switch, the cases contain expressions that are compared against the value of the switch expression. In a type switch, the cases contain types that are compared against the type of a specially annotated switch expression. The switch expression is evaluated exactly once in a switch statement.

​	有两种形式：`表达式选择`和`类型选择`。在`表达式选择`中，case 包含与switch表达式的值进行比较的表达式。在`类型选择`中，case 包含类型，这些类型与特别说明的 `switch 表达式`的类型进行比较。在`switch语句`中，`switch 表达式`被精确地求值一次。

#### Expression switches 表达式选择

In an expression switch, the switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom; the first one that equals the switch expression triggers execution of the statements of the associated case; the other cases are skipped. If no case matches and there is a "default" case, its statements are executed. There can be at most one default case and it may appear anywhere in the "switch" statement. A missing switch expression is equivalent to the boolean value `true`.

​	在表达式选择中，`switch 表达式`被求值，case表达式（不需要是常量）被从左到右和从上到下求值；第一个等于switch表达式的表达式会触发相关case语句的执行；其他case被跳过。如果没有匹配的case，但有一个 "default" case，那么这个语句将被执行。"default" case 最多只能有一个，它可以出现在 "switch"语句的任何地方。缺省的`switch 表达式`等同于布尔值`true`。

```
ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
ExprCaseClause = ExprSwitchCase ":" StatementList .
ExprSwitchCase = "case" ExpressionList | "default" .
```

If the switch expression evaluates to an untyped constant, it is first implicitly [converted](https://go.dev/ref/spec#Conversions) to its [default type](https://go.dev/ref/spec#Constants). The predeclared untyped value `nil` cannot be used as a switch expression. The switch expression type must be [comparable](https://go.dev/ref/spec#Comparison_operators).

​	如果`switch 表达式`求值为无类型的常量，它首先被隐式地[转换](../Expressions#conversions-转换)为其[默认类型](../Constants)。预先声明的无类型值`nil`不能作为`switch 表达式`使用。`switch 表达式`的类型必须是[可比较的](../Expressions#comparison-operators)。

> 个人注释
>
> ​	在表达式选择中存在nil的情况：
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	switch nil {
> 	default:
> 		fmt.Println("is default")
> 	case 1:
> 		fmt.Println("is 1")
> 	case 2:
> 		fmt.Println("is 2")
> 		// cannot convert nil to type int
> 		//case nil:
> 		//	fmt.Println("is 2")
> 	}
> }
> 
> 报错：use of untyped nil in switch expression
> 
> ```
>
> ​	可以看到在 `switch nil`中，程序报错：use of untyped nil in switch expression。
>
> ​	在 `case nil`中，程序报错：use of untyped nil in switch expression。	

If a case expression is untyped, it is first implicitly [converted](https://go.dev/ref/spec#Conversions) to the type of the switch expression. For each (possibly converted) case expression `x` and the value `t` of the switch expression, `x == t` must be a valid [comparison](https://go.dev/ref/spec#Comparison_operators).

​	如果一个case 表达式是无类型的，它首先被隐式地[转换](../Expressions#conversions-转换)为`switch 表达式`的类型。对于每个（可能转换过的）case 表达式`x`和`switch 表达式`的值`t`，`x == t`必须是一个有效的[比较](../Expressions#comparison-operators-比较运算符)。

In other words, the switch expression is treated as if it were used to declare and initialize a temporary variable `t` without explicit type; it is that value of `t` against which each case expression `x` is tested for equality.

​	换句话说，`switch 表达式`被当作是用来声明和初始化一个没有明确类型的临时变量`t`；每个case表达式`x`都是用`t`的值来测试是否相等。

In a case or default clause, the last non-empty statement may be a (possibly [labeled](https://go.dev/ref/spec#Labeled_statements)) ["fallthrough" statement](https://go.dev/ref/spec#Fallthrough_statements) to indicate that control should flow from the end of this clause to the first statement of the next clause. Otherwise control flows to the end of the "switch" statement. A "fallthrough" statement may appear as the last statement of all but the last clause of an expression switch.

​	在case或default子句中，最后一个非空语句可以是一个（可能被[标记的](#labeled-statements-标签语句)）["fallthrough"语句](#fallthrough-statements-语句-fallthrough)，表示控制应该从这个子句的结尾流向下一个子句的第一个语句。否则，控制将流向 "switch"语句的末尾。"fallthrough"语句可以作为`表达式选择`中除最后一个子句之外的所有子句的最后一个语句出现。

The switch expression may be preceded by a simple statement, which executes before the expression is evaluated.

​	`switch 表达式`前面可以有一个简单的语句，它在表达式被求值之前执行。

```go 
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // 缺少 switch 表达式 意味着 "true"
case x < 0: return -x
default: return x
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}
```

> 个人注释
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	// switch 表达式中是一个无类型的常量 的情况
> 	switch 1 {
> 	default:
> 		fmt.Println("is default")
> 	case 1:
> 		fmt.Println("is 1")
> 	case 2:
> 		fmt.Println("is 2")
> 	}
> 
> 	// switch中有一个简单的语句但没有省略 switch表达式的情况
> 	switch a := 1; a {
> 	default:
> 		fmt.Println("is default")
> 	case 1:
> 		fmt.Println("is 1")
> 	case 2:
> 		fmt.Println("is 2")
> 	}
> 
> 	// switch中有一个简单的语句且省略 switch表达式的情况
> 	switch a := 1; {
> 	default:
> 		fmt.Println("is default")
> 	case a == 1:
> 		fmt.Println("is 1")
> 	case a == 2:
> 		fmt.Println("is 2")
> 	}
> }
> 
> Output:
> is 1
> is 1
> is 1
> ```
>
> 

Implementation restriction: A compiler may disallow multiple case expressions evaluating to the same constant. For instance, the current compilers disallow duplicate integer, floating point, or string constants in case expressions.

实现限制：编译器可能不允许多个case 表达式求值为同一个常量。例如，目前的编译器不允许在case表达式中出现重复的整型常量、浮点常量或字符串常量。

#### Type switches 类型选择

A type switch compares types rather than values. It is otherwise similar to an expression switch. It is marked by a special switch expression that has the form of a [type assertion](https://go.dev/ref/spec#Type_assertions) using the keyword `type` rather than an actual type:

​	`类型选择`比较的是`类型`而不是值。它与`表达式选择`类似。它由一个特殊的`switch 表达式`标记，该表达式具有[类型断言](../Expressions#type-assertions-类型断言)的形式，使用关键字`type`而不是实际的类型：

```go 
switch x.(type) {
// cases
}
```

Cases then match actual types `T` against the dynamic type of the expression `x`. As with type assertions, `x` must be of [interface type](https://go.dev/ref/spec#Interface_types), but not a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations), and each non-interface type `T` listed in a case must implement the type of `x`. The types listed in the cases of a type switch must all be [different](https://go.dev/ref/spec#Type_identity).

​	然后，case 将实际类型`T`与表达式`x`的动态类型相匹配。与类型断言一样，`x`必须是[接口类型](../Types#interface-types-接口型)，但不能是[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)，而且case 中列出的每个非接口类型`T`必须实现`x`的类型。在类型选择的case 中，列出的类型都必须是[不同的](../PropertiesOfTypesAndValues#type-identity-类型一致性)。

``` go
typeSwitchStmt  = "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClause } "}" .
TypeSwitchGuard = [ identifier ":=" ] PrimaryExpr "." "(" "type" ")" .
TypeCaseClause  = TypeSwitchCase ":" StatementList .
TypeSwitchCase  = "case" TypeList | "default" .
```

The TypeSwitchGuard may include a [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations). When that form is used, the variable is declared at the end of the TypeSwitchCase in the [implicit block](https://go.dev/ref/spec#Blocks) of each clause. In clauses with a case listing exactly one type, the variable has that type; otherwise, the variable has the type of the expression in the TypeSwitchGuard.

​	TypeSwitchGuard 可以包括一个[短变量声明](../DeclarationsAndScope#short-variable-declarations-短变量声明)。当使用这种形式时，该变量在每个子句的[隐含块](../Blocks)中的TypeSwitchCase**的末尾**被声明。在子句中，如果case正好列出了一种类型，那么变量就有这种类型；否则，变量就有TypeSwitchGuard中表达式的类型。

Instead of a type, a case may use the predeclared identifier [`nil`](https://go.dev/ref/spec#Predeclared_identifiers); that case is selected when the expression in the TypeSwitchGuard is a `nil` interface value. There may be at most one `nil` case.

​	case 可以使用预先声明的标识符[nil](../DeclarationsAndScope#predeclared-identifiers--预先声明的标识符)来代替类型；当TypeSwitchGuard中的表达式是一个`nil`接口值时，该case被选中。最多只能有一个`nil `case。

Given an expression `x` of type `interface{}`, the following type switch:

​	给定一个`interface{}`类型的表达式`x`，下面的类型选择：

```go 
switch i := x.(type) {
case nil:
	printString("x is nil")                // i 类型是 x 的类型（interface{}）
case int:
	printInt(i)                            // i 类型是 int
case float64:
	printFloat64(i)                        // i 类型是 float64 
case func(int) float64:
	printFunction(i)                       // i 类型是 func(int) float64
case bool, string:
	printString("type is bool or string")  // i 类型是 x 的类型（interface{}）
default:
	printString("don't know the type")     // i 类型是 x 的类型（interface{}）
}
```

> 个人注释
>
> ```go
> package main
> 
> import "fmt"
> 
> func JudgeType(x any) {
> 	switch i := x.(type) {
> 	default:
> 		fmt.Println("don't know the type")
> 		fmt.Printf("%T,i=%#v\n", i, i)
> 	case int8, int16, int32, int64, int:
> 		fmt.Println("is int8、int16、int32、int64、int")
> 		fmt.Printf("%T,i=%#v\n", i, i)
> 	case uint8, uint16, uint32, uint64, uint:
> 		fmt.Println("is uint8、uint16、uint32、uint64、uint")
> 		fmt.Printf("%T,i=%#v\n", i, i)
> 	case float32, float64:
> 		fmt.Println("is float32、float64")
> 		fmt.Printf("%T,i=%#v\n", i, i)
> 	case complex64, complex128:
> 		fmt.Println("is complex64、complex128")
> 		fmt.Printf("%T,i=%#v\n", i, i)
> 	}
> }
> func main() {
> 	var x interface{}
> 	JudgeType(x)
> 	fmt.Println("------------------------------")
> 	x = nil
> 	JudgeType(x)
> 	fmt.Println("------------------------------")
> 	x = int8(1)
> 	JudgeType(x)
> 	fmt.Println("------------------------------")
> 	x = uint16(1)
> 	JudgeType(x)
> 	fmt.Println("------------------------------")
> 	x = 1.2
> 	JudgeType(x)
> 	fmt.Println("------------------------------")
> 	x = complex(1.2, 3.4)
> 	JudgeType(x)
> 	fmt.Println("------------------------------")
> }
> Output:
> 
> don't know the type
> <nil>,i=<nil>
> ------------------------------
> don't know the type
> <nil>,i=<nil>
> ------------------------------
> is int8、int16、int32、int64、int
> int8,i=1
> ------------------------------
> is uint8、uint16、uint32、uint64、uint
> uint16,i=0x1
> ------------------------------
> is float32、float64
> float64,i=1.2
> ------------------------------
> is complex64、complex128
> complex128,i=(1.2+3.4i)
> ------------------------------
> 
> ```
>
> 

could be rewritten:

可以被重写为：

```go 
v := x  // x 只被求值一次
if v == nil {
	i := v    // i 类型为 x 的类型（interface{}）
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i) // i 类型是 int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i) // i 类型是 float64 
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i) // i 类型是 func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v  // i 类型是 x 的类型（interface{}）
		printString("type is bool or string")
	} else {
		i := v // i 类型是 x 的类型（interface{}）
		printString("don't know the type")
	}
}
```

A [type parameter](https://go.dev/ref/spec#Type_parameter_declarations) or a [generic type](https://go.dev/ref/spec#Type_declarations) may be used as a type in a case. If upon [instantiation](https://go.dev/ref/spec#Instantiations) that type turns out to duplicate another entry in the switch, the first matching case is chosen.

​	[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)或[泛型](../DeclarationsAndScope#type-declarations-类型声明)可以作为 case 中的一个类型。如果在[实例化](../Expressions#instantiations-实例化)时，该类型被发现与switch中的另一个case重复，则选择第一个匹配的case。

```go  hl_lines="7 7"
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

The type switch guard may be preceded by a simple statement, which executes before the guard is evaluated.

​	类型选择防护（guard ）前可以有一个简单的语句，该语句在防护（guard ）被求值前执行。

The "fallthrough" statement is not permitted in a type switch.

​	在类型选择中`不允许使用`"fallthrough"语句。

### For statements  - for 语句

A "for" statement specifies repeated execution of a block. There are three forms: The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.

​	 "`for`"语句指定重复执行一个块。有三种形式：迭代可以由单个条件、"for"子句或 "range"子句控制。

```
ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
Condition = Expression .
```

#### For statements with single condition 带有单一条件的for语句

In its simplest form, a "for" statement specifies the repeated execution of a block as long as a boolean condition evaluates to true. The condition is evaluated before each iteration. If the condition is absent, it is equivalent to the boolean value `true`.

​	在其最简单的形式中，"for"语句指定重复执行一个块，只要一个布尔条件被求值为真。该条件在每次迭代前被求值。如果条件不存在，它就等同于布尔值`true`。

```go 
for a < b {
	a *= 2
}
```

#### For statements with `for` clause 带有for子句的for语句

A "for" statement with a ForClause is also controlled by its condition, but additionally it may specify an *init* and a *post* statement, such as an assignment, an increment or decrement statement. The init statement may be a [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations), but the post statement must not.

​	带有for子句的 "for"语句也受其条件的控制，但另外它可以指定一个`init`和一个`post`语句，如一个赋值，一个增量或减量语句。init语句可以是一个[短变量声明](../DeclarationsAndScope#short-variable-declarations-短变量声明)，但post语句则不能。由init语句声明的变量会在每次迭代中重复使用。

```
ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
InitStmt = SimpleStmt .
PostStmt = SimpleStmt .
```

```go
for i := 0; i < 10; i++ {
	f(i)
}
```

If non-empty, the init statement is executed once before evaluating the condition for the first iteration; the post statement is executed after each execution of the block (and only if the block was executed). Any element of the ForClause may be empty but the [semicolons](https://go.dev/ref/spec#Semicolons) are required unless there is only a condition. If the condition is absent, it is equivalent to the boolean value `true`.

​	如果（init语句）非空，则它在求值第一个迭代的条件之前被执行一次；post语句在每次执行块之后被执行（而且只有当该块被执行时）。For子句的任何元素都可以是空的，但是[分号](../LexicalElements#semicolons-分号)是必须的，除非只有一个条件（则可以省略分号）。如果没有条件，则等同于布尔值`true`。

```go
for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }
```

Each iteration has its own separate declared variable (or variables) [[Go 1.22]({{< ref "/langSpec/Appendix#go-122">}})]. The variable used by the first iteration is declared by the init statement. The variable used by each subsequent iteration is declared implicitly before executing the post statement and initialized to the value of the previous iteration's variable at that moment.

​	每次迭代都有其自己的单独声明的变量（或变量）[[Go 1.22]({{< ref "/langSpec/Appendix#go-122">}})]。第一次迭代使用的变量由 init 语句声明。每次后续迭代使用的变量在执行 post 语句之前隐式声明，并初始化为该时刻前一次迭代变量的值。

```go
var prints []func()
for i := 0; i < 5; i++ {
	prints = append(prints, func() { println(i) })
	i++
}
for _, p := range prints {
	p()
}
```

prints

打印

```
1
3
5
```

Prior to [[Go 1.22]({{< ref "/langSpec/Appendix#go-122">}})], iterations share one set of variables instead of having their own separate variables. In that case, the example above prints

​	在 [[Go 1.22]({{< ref "/langSpec/Appendix#go-122">}})] 之前，迭代共享一组变量，而不是拥有它们自己的独立变量。在这种情况下，上面的示例打印

```
6
6
6
```

#### For statements with `range` clause  带有range子句的for语句

A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, values received on a channel, or integer values from zero to an upper limit [[Go 1.22]({{< ref "/langSpec/Appendix#go-122">}})]. For each entry it assigns *iteration values* to corresponding *iteration variables* if present and then executes the block.

​	带有 "`range`"子句的 "for"语句会迭代数组、切片、字符串或映射的所有条目，在通道上接收的值，或从零到上限的整数值 [[Go 1.22]({{< ref "/langSpec/Appendix#go-122">}})]。对于每个条目，它将迭代值分配给相应的迭代变量（如果存在），然后执行块。

```
RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
```

The expression on the right in the "range" clause is called the *range expression*, its [core type](https://go.dev/ref/spec#Core_types) must be an array, pointer to an array, slice, string, map, channel permitting [receive operations](https://go.dev/ref/spec#Receive_operator), or an integer. As with an assignment, if present the operands on the left must be [addressable](https://go.dev/ref/spec#Address_operators) or map index expressions; they denote the iteration variables. If the range expression is a channel or integer, at most one iteration variable is permitted, otherwise there may be up to two. If the last iteration variable is the [blank identifier](https://go.dev/ref/spec#Blank_identifier), the range clause is equivalent to the same clause without that identifier.

​	在 "`range`"子句中`右边的表达式`称为`范围表达式`，其[核心类型](../PropertiesOfTypesAndValues#core-types-核心类型)必须是数组、指向数组的指针、切片、字符串、映射或允许[接收操作](../Expressions#receive-operator-接收操作符)的通道。和赋值一样，如果左操作数存在的话，那么它必须是[可寻址的](../Expressions#address-operators-地址运算符)或映射索引表达式；它们（即左操作数）表示`迭代变量`。如果`range 表达式`是一个通道，最多允许一个迭代变量，其他情况下最多可以有两个迭代变量。如果最后一个迭代变量是[空白标识符，即`_`](../DeclarationsAndScope#blank-identifierr-空白标识符)，那么range 子句就等同于没有该空白标识符的相同子句。

The range expression `x` is evaluated once before beginning the loop, with one exception: if at most one iteration variable is present and `len(x)` is [constant](https://go.dev/ref/spec#Length_and_capacity), the range expression is not evaluated.

​	在开始循环之前，range 表达式`x`被求值一次，**但有一个例外**：如果最多只有一个迭代变量，并且`len(x)`是[常量](../Constants)，那么range 表达式不被求值。

> 个人注释
>
> ​	是否以下的示例满足上面所说的：“如果最多只有一个迭代变量，并且`len(x)`是[常量](../Constants)，那么range 表达式不被求值” ？TODO
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"math/rand"
> 	"time"
> )
> 
> func main() {
> 	x1 := []int{1, 2, 3, 4, 5, 6}
> 	fmt.Printf("before x1=%#v\n", x1)
> 	rand.Seed(time.Now().UnixNano())
> 
> 	for i := range x1 {
> 		// 生成随机整数，范围: [0,6)
> 		rd := rand.Intn(6)
> 		fmt.Printf("i=%d,rd=%d,x1[i]=%v\n", i, rd, x1[i])
> 		x1[rd] = rd*2 + 1
> 		fmt.Printf("x1[rd]=%v\n", x1[rd])
> 		fmt.Println("----------------------")
> 	}
> 	fmt.Printf("after x1=%#v\n", x1)
> }
> 
> 
> Output:
> before x1=[]int{1, 2, 3, 4, 5, 6}
> i=0,rd=1,x1[i]=1
> x1[rd]=3
> ----------------------
> i=1,rd=3,x1[i]=3
> x1[rd]=7
> ----------------------
> i=2,rd=2,x1[i]=3
> x1[rd]=5
> ----------------------
> i=3,rd=1,x1[i]=7
> x1[rd]=3
> ----------------------
> i=4,rd=3,x1[i]=5
> x1[rd]=7
> ----------------------
> i=5,rd=3,x1[i]=6
> x1[rd]=7
> ----------------------
> after x1=[]int{1, 3, 5, 7, 5, 6}
> ```
>
> 

Function calls on the left are evaluated once per iteration. For each iteration, iteration values are produced as follows if the respective iteration variables are present:

​	左边的函数调用在每个迭代中被求值一次（=>仍有疑问？？）。对于每个迭代，如果各自的迭代变量存在，则迭代值按以下产生：

```go
Range expression                          1st value          2nd value

range 表达式                               第一个值             第二个值

array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, <-chan E       element  e  E
```

1. For an array, pointer to array, or slice value `a`, the index iteration values are produced in increasing order, starting at element index 0. If at most one iteration variable is present, the range loop produces iteration values from 0 up to `len(a)-1` and does not index into the array or slice itself. For a `nil` slice, the number of iterations is 0.

2. 对于一个数组值、数组指针值或切片值`a`，索引迭代值按递增顺序产生，从元素索引0开始。如果最多只有一个迭代变量，range 循环产生从`0`到`len(a)-1`的迭代值，并且不对数组或切片本身进行索引。对于一个`nil` 切片，迭代次数为0。

   > 个人注释
   >
   > ```go
   > package main
   > 
   > import (
   > 	"fmt"
   > )
   > 
   > func main() {
   > 	a := []byte{'a', 'b', 'c'}
   > 	for i, v := range a {
   > 		fmt.Printf("%v,%v,%d,%x,%X,%#x,%#X,%s,%q\n", i, v, v, v, v, v, v, string(v), string(v))
   > 	}
   > 	fmt.Println("--------------------")
   > 	var s []int
   > 	for i, v := range s {
   > 		fmt.Printf("%v,%v,%d,%x,%X,%#x,%#X,%s,%q\n", i, v, v, v, v, v, v, string(v), string(v))
   > 	}
   > }
   > 
   > Output:
   > 0,97,97,61,61,0x61,0X61,a,"a"
   > 1,98,98,62,62,0x62,0X62,b,"b"
   > 2,99,99,63,63,0x63,0X63,c,"c"
   > --------------------
   > 
   > ```
   >
   > 

3. For a string value, the "range" clause iterates over the Unicode code points in the string starting at byte index 0. On successive iterations, the index value will be the index of the first byte of successive UTF-8-encoded code points in the string, and the second value, of type `rune`, will be the value of the corresponding code point. If the iteration encounters an invalid UTF-8 sequence, the second value will be `0xFFFD`, the Unicode replacement character, and the next iteration will advance a single byte in the string.

4. 对于一个字符串值，"range"子句在字符串中的Unicode码点上进行迭代，从字节索引0开始。在连续的迭代中，索引值将是字符串中连续的UTF-8编码码点的第一个字节的索引，第二个值，类型为`rune`，将是相应码点的值。`如果迭代遇到一个无效的UTF-8序列`，第二个值将是（Unicode替换字符）`0xFFFD`，下一次迭代将在字符串中推进一个字节。

   > 个人注释
   >
   > ```go
   > package main
   > 
   > import (
   > 	"fmt"
   > )
   > 
   > func main() {
   > 	str := "我爱我的\x80祖国\x80"
   > 	for i, v := range str {
   > 		fmt.Printf("%v,%v,%d,%x,%X,%#x,%#X,%s,%q,%U\n", i, v, v, v, v, v, v, string(v), string(v), v)
   > 	}
   > }
   > 
   > Output:
   > 0,25105,25105,6211,6211,0x6211,0X6211,我,"我",U+6211
   > 3,29233,29233,7231,7231,0x7231,0X7231,爱,"爱",U+7231
   > 6,25105,25105,6211,6211,0x6211,0X6211,我,"我",U+6211
   > 9,30340,30340,7684,7684,0x7684,0X7684,的,"的",U+7684
   > 12,65533,65533,fffd,FFFD,0xfffd,0XFFFD,�,"�",U+FFFD
   > 13,31062,31062,7956,7956,0x7956,0X7956,祖,"祖",U+7956
   > 16,22269,22269,56fd,56FD,0x56fd,0X56FD,国,"国",U+56FD
   > 19,65533,65533,fffd,FFFD,0xfffd,0XFFFD,�,"�",U+FFFD
   > ```
   >
   > 

5. The iteration order over maps is not specified and is not guaranteed to be the same from one iteration to the next. If a map entry that has not yet been reached is removed during iteration, the corresponding iteration value will not be produced. If a map entry is created during iteration, that entry may be produced during the iteration or may be skipped. The choice may vary for each entry created and from one iteration to the next. If the map is `nil`, the number of iterations is 0.

6. 对映射的迭代顺序没有指定，不保证每次迭代都是一样的。如果在迭代过程中删除了一个尚未到达的映射条目，将不会产生相应的迭代值。如果在迭代过程中创建了一个映射条目，该条目可能在迭代过程中被产生，也可能被跳过。对于每个创建的条目，以及从一个迭代到另一个迭代，选择可能有所不同。如果映射为`nil`，迭代次数为0。

   > 个人注释
   >
   > ```go
   > package main
   > 
   > import (
   > 	"fmt"
   > )
   > 
   > func main() {
   > 	m := map[string]int{"a": 1, "b": 2, "1a": 3, "1b": 4, "0a": 5, "0b": 6}
   > 	for k, v := range m {
   > 		fmt.Printf("%v,%v,%d,%x,%X,%#x,%#X,%s,%q\n", k, v, v, v, v, v, v, string(v), string(v))
   > 	}
   > }
   > 第一次
   > Output:
   > a,1,1,1,1,0x1,0X1,,"\x01"
   > b,2,2,2,2,0x2,0X2,,"\x02"
   > 1a,3,3,3,3,0x3,0X3,,"\x03"
   > 1b,4,4,4,4,0x4,0X4,,"\x04"
   > 0a,5,5,5,5,0x5,0X5,,"\x05"
   > 0b,6,6,6,6,0x6,0X6,,"\x06"
   > 
   > 第二次
   > Output:
   > 1a,3,3,3,3,0x3,0X3,,"\x03"
   > 1b,4,4,4,4,0x4,0X4,,"\x04"
   > 0a,5,5,5,5,0x5,0X5,,"\x05"
   > 0b,6,6,6,6,0x6,0X6,,"\x06"
   > a,1,1,1,1,0x1,0X1,,"\x01"
   > b,2,2,2,2,0x2,0X2,,"\x02"
   > ```
   >
   > 

7. For channels, the iteration values produced are the successive values sent on the channel until the channel is [closed](https://go.dev/ref/spec#Close). If the channel is `nil`, the range expression blocks forever.

8. 对于通道，产生的迭代值是通道上连续发送的值，直到通道[关闭](../Built-inFunctions#close)。如果通道为`nil`，则 range 表达式永远阻塞。

   > 个人注释
   >
   > ```go
   > package main
   > 
   > import (
   > 	"fmt"
   > 	"time"
   > )
   > 
   > func main() {
   > 	ch := make(chan int, 3)
   > 
   > 	go func() {
   > 		defer close(ch)
   > 		for i := 0; i <= 2; i++ {
   > 			ch <- i
   > 		}
   > 		time.Sleep(6 * time.Second)
   > 		fmt.Println("end sleeping 6s")
   > 		for i := 3; i <= 5; i++ {
   > 			ch <- i
   > 		}
   > 	}()
   > 
   > 	// 报错：range over ch (variable of type chan int) permits only one iteration variable
   > 	//for v, ok := range ch {
   > 	//	fmt.Printf("v=%v,ok=%T\n", v, ok)
   > 	//}
   > 	for v := range ch {
   > 		fmt.Printf("v=%v\n", v)
   > 	}
   > }
   > 
   > 
   > Output:
   > v=0
   > v=1
   > v=2
   > end sleeping 6s
   > v=3
   > v=4
   > v=5
   > 
   > ```
   >
   > 

9. For an integer value `n`, the iteration values 0 through `n-1` are produced in increasing order. If `n` <= 0, the loop does not run any iterations.

10. 对于整数值 `n` ，迭代值 0 到 `n-1` 按递增顺序生成。如果 `n` <= 0，则循环不会运行任何迭代。


The iteration variables may be declared by the "range" clause using a form of [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations) (`:=`). In this case their [scope](https://go.dev/ref/spec#Declarations_and_scope) is the block of the "for" statement and each iteration has its own new variables [[Go 1.22]({{< ref "/langSpec/Appendix#go-122">}})] (see also ["for" statements with a ForClause](https://go.dev/ref/spec#For_clause)). If the range expression is a (possibly untyped) integer expression `n`, the variable has the same type as if it was [declared](https://go.dev/ref/spec#Variable_declarations) with initialization expression `n`. Otherwise, the variables have the types of their respective iteration values.

​	迭代变量可以通过“range”子句使用短变量声明的形式（ `:=` ）来声明。在这种情况下，它们的范围是“for”语句的块，并且每次迭代都有其自己的新变量 [ [Go 1.22]({{< ref "/langSpec/Appendix#go-122">}})]（另请参阅带有 ForClause 的“for”语句）。如果范围表达式是（可能未键入的）整数表达式 `n` ，则变量的类型与使用初始化表达式 `n` 声明的类型相同。否则，变量具有其各自迭代值的数据类型。

If the iteration variables are not explicitly declared by the "range" clause, they must be preexisting. In this case, the iteration values are assigned to the respective variables as in an [assignment statement](https://go.dev/ref/spec#Assignment_statements). If the range expression is a (possibly untyped) integer expression `n`, `n` too must be [assignable](https://go.dev/ref/spec#Assignability) to the iteration variable; if there is no iteration variable, `n` must be assignable to `int`.

​	如果迭代变量未通过“range”子句显式声明，则它们必须是预先存在的。在这种情况下，迭代值将分配给各个变量，就像在赋值语句中一样。如果范围表达式是（可能未键入的）整数表达式 `n` ，则 `n` 也必须可分配给迭代变量；如果没有迭代变量，则 `n` 必须可分配给 `int` 。

```go 
var testdata *struct {
	a *[7]int
}
for i, _ := range testdata.a {
    // testdata.a is never evaluated; len(testdata.a) is constant
	// i ranges from 0 to 6
	// testdata.a 不会被求值； len(testdata.a) 是常量 <= 这里符合前面说的那个"但有一个例外"的情况吗？
	// i 范围从 0 到 6
	f(i)
}

var a [10]string
for i, s := range a {
    // type of i is int
	// type of s is string
	// s == a[i]
	// i 的类型是 int
	// s 的类型是 string
	// s == a[i]
	g(i, s)
}

var key string
var val interface{}  // element type of m is assignable to val  m 的元素类型可赋值给 val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
	h(key, val)
}
// key == last map key encountered in iteration
// val == map[key]
// key == 迭代中遇到的最后一个映射键
// val == map[key]

var ch chan Work = producer()
for w := range ch {
	doWork(w)
}

// 清空通道
// empty a channel
for range ch {}

// call f(0), f(1), ... f(9)
for i := range 10 {
	// type of i is int (default type for untyped constant 10)
	f(i)
}

// invalid: 256 cannot be assigned to uint8
var u uint8
for u = range 256 {
}
```

### Go statements  - go 语句

A "go" statement starts the execution of a function call as an independent concurrent thread of control, or *goroutine*, within the same address space.

​	`go`语句作为一个独立的并发控制线程（或称为goroutine），`在同一地址空间内`开始执行一个函数调用。

```
GoStmt = "go" Expression .
```

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for [expression statements](https://go.dev/ref/spec#Expression_statements).

​	（go语句中的）表达式`必须是一个函数或方法调用`；其不能用圆括号括起来。对内置函数的调用与[表达式语句](#expression-statements-表达式语句)一样受到限制。

> 个人注释
>
> ​	是不是不能使用 `go + 内置函数` 或 `go + 表达式语句`？
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"time"
> )
> 
> func f(x int) {
> 	fmt.Printf("x=%d\n", x)
> }
> func main() {
> 	//i := []int{0, 1}
> 	//go append(i, 2) // 报错：go discards result of append(i, 2) (value of type []int)
> 
> 	ch := make(chan int)
> 	go func() {
> 		ch <- 1
> 	}()
> 
> 	<-ch
> 	go close(ch)
> 
> 	//s := []int{0, 1}
> 	//go len(s) // 报错：go discards result of len(s) (value of type int)
> 	x := 1
>     // go + 表达式语句
> 	go f(x)
> 
> 	time.Sleep(2 * time.Second)
> 	fmt.Println("main over")
> }
> 
> ```
>
> ​	可见并不是不能使用 `go + 内置函数` 或 `go + 表达式语句`，而是使用的内置函数或表达式语句有限制，即有些内置函数、表达式语句不能使用，有些内置函数或表达式语句可以使用！

The function value and parameters are [evaluated as usual](https://go.dev/ref/spec#Calls) in the calling goroutine, but unlike with a regular call, program execution does not wait for the invoked function to complete. Instead, the function begins executing independently in a new goroutine. When the function terminates, its goroutine also terminates. If the function has any return values, they are discarded when the function completes.

​	函数值和参数在调用的goroutine中[像往常一样被求值](../Expressions#order-of-evaluation-求值顺序)，但与普通调用不同的是，程序执行不会等待被调用的函数完成。相反，该函数开始在一个新的goroutine中独立执行。当该函数终止时，其goroutine也会终止。如果该函数有任何返回值，当函数完成时，它们会被丢弃。

```go 
go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true }} (c)
```

### Select statements - select 语句

A "select" statement chooses which of a set of possible [send](https://go.dev/ref/spec#Send_statements) or [receive](https://go.dev/ref/spec#Receive_operator) operations will proceed. It looks similar to a ["switch"](https://go.dev/ref/spec#Switch_statements) statement but with the cases all referring to communication operations.

​	 "`select`"语句选择一组可能的[发送](#send-statements-发送语句)或[接收](../Expressions#receive-operator-接收操作符)操作中的一个来进行。它看起来类似于 "[switch](#switch-statements----switch-语句) "语句，但其 case 都是只涉及通信操作。

```
SelectStmt = "select" "{" { CommClause } "}" .
CommClause = CommCase ":" StatementList .
CommCase   = "case" ( SendStmt | RecvStmt ) | "default" .
RecvStmt   = [ ExpressionList "=" | IdentifierList ":=" ] RecvExpr .
RecvExpr   = Expression .
```

A case with a RecvStmt may assign the result of a RecvExpr to one or two variables, which may be declared using a [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations). The RecvExpr must be a (possibly parenthesized) receive operation. There can be at most one default case and it may appear anywhere in the list of cases.

​	一个有 RecvStmt 的 case 可以将 RecvExpr 的结果分配给一个或两个变量，这些变量可以用[短变量声明](../DeclarationsAndScope#short-variable-declarations-短变量声明)来声明。RecvExpr 必须是一个（可能是圆括号内的）接收操作。最多可以有一个default case，它可以出现在 case 列表的任何地方。

Execution of a "select" statement proceeds in several steps:

​	 "`select`"语句的执行分几个步骤进行：

1. For all the cases in the statement, the channel operands of receive operations and the channel and right-hand-side expressions of send statements are evaluated exactly once, in source order, upon entering the "select" statement. The result is a set of channels to receive from or send to, and the corresponding values to send. Any side effects in that evaluation will occur irrespective of which (if any) communication operation is selected to proceed. Expressions on the left-hand side of a RecvStmt with a short variable declaration or assignment are not yet evaluated.
2. 对于语句中的所有情况，在进入 "`select`"语句时，接收操作的通道操作数、通道、发送语句的右侧表达式（按源代码出现的顺序）被求值一次。结果是一组要接收或发送的通道，以及对应的要发送的值。无论选择哪一个（如果有）通信操作来进行，在这次求值中的任何副作用都会发生。在 RecvStmt 左侧的表达式有一个短变量声明或赋值，还没有被求值。
3. If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement blocks until at least one of the communications can proceed.
4. 如果一个或多个通信可以进行，则通过统一的伪随机选择来选择一个可以进行的通信。否则，如果有一个default case，就会选择该情况。如果没有default case，"`select`"语句就会阻塞，直到至少有一项通信可以进行。
5. Unless the selected case is the default case, the respective communication operation is executed.
6. 除非选择的情况是default case，否则将执行相应的通信操作。
7. If the selected case is a RecvStmt with a short variable declaration or an assignment, the left-hand side expressions are evaluated and the received value (or values) are assigned.
8. 如果选择的 case 是一个带有短变量声明或赋值的RecvStmt，左边的表达式会被求值，接收到的值（或多个值）被用于赋值。
9. The statement list of the selected case is executed.
10. 所选 case 语句列表被执行。

Since communication on `nil` channels can never proceed, a select with only `nil` channels and no default case blocks forever.

​	由于在`nil`通道上的通信永远不能进行，所以只有`nil`通道且没有default case 的`select`语句会永远阻塞。

```go 
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

A "return" statement in a function `F` terminates the execution of `F`, and optionally provides one or more result values. Any functions [deferred](https://go.dev/ref/spec#Defer_statements) by `F` are executed before `F` returns to its caller.

​	函数`F`中的 "`return`"语句终止了`F`的执行，并且可以选择提供一个或多个结果值。在`F`返回给它的调用者之前，任何由`F`[延迟](#defer-statements-语句-defer)的函数都会被执行。

```
ReturnStmt = "return" [ ExpressionList ] .
```

In a function without a result type, a "return" statement must not specify any result values.

​	在没有结果类型的函数中，"`return`"语句必须不指定任何结果值。

```go 
func noResult() {
	return
}
```

There are three ways to return values from a function with a result type:

​	有三种方法可以从一个有结果类型的函数中返回值：

1. The return value or values may be explicitly listed in the "return" statement. Each expression must be single-valued and [assignable](https://go.dev/ref/spec#Assignability) to the corresponding element of the function's result type.

2. 可以在 "`return`"语句中显式地列出一个或多个返回值。每个表达式必须是单值的，并且可以分配给函数的结果类型的相应元素。

   ```go 
   func simpleF() int {
   	return 2
   }
   
   func complexF1() (re float64, im float64) {
   	return -7.0, -4.0
   }
   ```

3. The expression list in the "return" statement may be a single call to a multi-valued function. The effect is as if each value returned from that function were assigned to a temporary variable with the type of the respective value, followed by a "return" statement listing these variables, at which point the rules of the previous case apply.

4. "`return`"语句中的表达式列表可能是对一个多值函数的单一调用。其效果就像从该函数返回的每个值都被分配到一个临时变量中，其类型与相应的值类型相同，随后的 "`return`"语句列出了这些变量，此时，前一种情况的规则适用。

   ```go 
   func complexF2() (re float64, im float64) {
   	return complexF1()
   }
   ```

5. The expression list may be empty if the function's result type specifies names for its [result parameters](https://go.dev/ref/spec#Function_types). The result parameters act as ordinary local variables and the function may assign values to them as necessary. The "return" statement returns the values of these variables.

6. 如果函数的结果类型为其结果参数指定了名称，该表达式列表可能是空的。结果参数作为普通的局部变量，函数可以根据需要给它们赋值。`return`语句会返回这些变量的值。

   ```go 
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

Regardless of how they are declared, all the result values are initialized to the [zero values](https://go.dev/ref/spec#The_zero_value) for their type upon entry to the function. A "return" statement that specifies results sets the result parameters before any deferred functions are executed.

​	不管它们是如何被声明的，`所有的结果值在进入函数时都被初始化为`其类型的[零值](../ProgramInitializationAndExecution#the-zero-value-零值)。指定结果的 "`return`"语句`在执行任何延迟函数之前`设置结果参数。

Implementation restriction: A compiler may disallow an empty expression list in a "return" statement if a different entity (constant, type, or variable) with the same name as a result parameter is in [scope](https://go.dev/ref/spec#Declarations_and_scope) at the place of the return.

​	实现限制：如果在返回的地方有一个与结果参数同名的不同实体（常量、类型或变量）在[作用域](../DeclarationsAndScope)内，编译器可能不允许在 "`return`"语句中出现空表达式列表。

```go 
func f(n int) (res int, err error) {
	if _, err := f(n-1); err != nil {
		return  // 无效的返回语句： err 被遮蔽了
	}
	return
}
```

> 个人注释
>
> ```go
> package main
> 
> import "fmt"
> 
> func f(n int) (res int, err error) {
> 	if _, err := f(n - 1); err != nil { // 报错： inner declaration of var err error
> 		return // 无效的返回语句： err 被遮蔽了 // 报错： result parameter err not in scope at return
> 	} //
> 	return
> }
> 
> func main() {
> 	fmt.Println(f(1))
> }
> 
> ```
>
> 

### Break statements - break 语句

A "break" statement terminates execution of the innermost ["for"](https://go.dev/ref/spec#For_statements), ["switch"](https://go.dev/ref/spec#Switch_statements), or ["select"](https://go.dev/ref/spec#Select_statements) statement within the same function.

​	 "`break`"语句可以终止同一函数中最里面的 "[for](#for-statements-for----for-语句)"、"[switch](#switch-statements-switch----switch-语句) "或 "[select](#select-statements-select---select-语句) "语句的执行。

```
BreakStmt = "break" [ Label ] .
```

If there is a label, it must be that of an enclosing "for", "switch", or "select" statement, and that is the one whose execution terminates.

​	如果有一个标签，它必须是一个封闭的 "`for`"、"`switch`"或 "`select`"语句的标签，而且是执行终止的那一个。

```go 
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

A "continue" statement begins the next iteration of the innermost enclosing ["for" loop](https://go.dev/ref/spec#For_statements) by advancing control to the end of the loop block. The "for" loop must be within the same function.

​	 "`continue`"语句通过将控制推进到循环块的末端来开始最内层的 ["for"循环](#for-statements-for----for-语句)的下一次迭代。"`for`"循环必须是在同一个函数中。

```
ContinueStmt = "continue" [ Label ] .
```

If there is a label, it must be that of an enclosing "for" statement, and that is the one whose execution advances.

​	如果有一个标签，它必须是一个封闭的 "`for`"语句的标签，而且是执行前进的那一个。

```go 
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

A "goto" statement transfers control to the statement with the corresponding label within the same function.

​	 "`goto`"语句将控制转移到同一函数中具有相应标签的语句。

```
GotoStmt = "goto" Label .
goto Error
```

Executing the "goto" statement must not cause any variables to come into [scope](https://go.dev/ref/spec#Declarations_and_scope) that were not already in scope at the point of the goto. For instance, this example:

​	执行 "`goto`"语句不能导致任何变量进入goto处的作用域之外的[作用域](../DeclarationsAndScope)。例如，这个例子：

```go 
	goto L  // BAD
	v := 3
L:
```

is erroneous because the jump to label `L` skips the creation of `v`.

是错误的，因为跳转到标签`L`时，跳过了创建`v`的过程。

A "goto" statement outside a [block](https://go.dev/ref/spec#Blocks) cannot jump to a label inside that block. For instance, this example:

​	[块](../Blocks)外的 "`goto`"语句不能跳到该块内的标签。例如，这个例子：

```go 
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

is erroneous because the label `L1` is inside the "for" statement's block but the `goto` is not.

是错误的，因为标签`L1`在 "for"语句的块内，而`goto`不在其中。

### Fallthrough statements 语句 fallthrough

A "fallthrough" statement transfers control to the first statement of the next case clause in an [expression "switch" statement](https://go.dev/ref/spec#Expression_switches). It may be used only as the final non-empty statement in such a clause.

​	在[表达式选择语句](#expression-switches-表达式选择)中，"`fallthrough`"语句将控制转移到下一个`case`子句的第一个语句。它只能作为这种子句中的最后一个非空语句使用。

```
FallthroughStmt = "fallthrough" .
```

> 个人注释
>
> ​	`fallthrough`语句难道就不能在类型选择语句中使用吗？=> 是的，不能！
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
> 	var i interface{}
> 
> 	i = 1
> 	switch i.(type) {
> 	case int:
> 		fmt.Println("is int")
> 		fallthrough  // 报错： cannot fallthrough in type switch
> 	case uint:
> 		fmt.Println("is uint")
> 	default:
> 		fmt.Println("don't know the type")
> 	}
> }
> 
> ```
>
> ​	给出`fallthrough`的示例，如下所示：
>
> ​	可见default语句所处位置对于switch中的`fallthrough`语句还是有影响的！

{{< tabpane text=true >}}

{{< tab header="default在最前面的情况" >}}

```go
package main

import (
	"fmt"
	"math/rand"
)

func GetRandInt() int {
	return rand.Intn(7)
}
func GetRank() {
	switch i := GetRandInt(); {
	default:
		fmt.Printf("------ i=%d -- over----------\n", i)
	case i == 6:
		fmt.Println("i is 6")
		fallthrough
	case i < 6:
		if i < 6 {
			fmt.Println("less 6")
		}
		fallthrough
	case i < 5:
		if i < 5 {
			fmt.Println("less 5")
		}
		fallthrough
	case i < 4:
		if i < 4 {
			fmt.Println("less 4")
		}
		fallthrough
	case i < 3:
		if i < 3 {
			fmt.Println("less 3")
		}
		fallthrough
	case i < 2:
		if i < 2 {
			fmt.Println("less 2")
		}
		fallthrough
	case i < 1:
		if i < 1 {
			fmt.Println("less 1")
		}
		//fallthrough //报错： cannot fallthrough final case in switch
	}
}

func main() {
	x := 0
	for x <= 5 {
		x++
		GetRank()
	}
}
Output:
less 6
less 5
less 4
less 3
less 2
less 6
less 6
less 5
less 4
less 3
less 2
less 6
less 5
less 4
less 6
i is 6

```

{{< /tab >}}

{{< tab header="default在最后面的情况" >}}

```go
package main

import (
	"fmt"
	"math/rand"
)

func GetRandInt() int {
	return rand.Intn(7)
}

func GetRank() {
	switch i := GetRandInt(); {
	case i == 6:
		fmt.Println("i is 6")
		fallthrough
	case i < 6:
		if i < 6 {
			fmt.Println("less 6")
		}
		fallthrough
	case i < 5:
		if i < 5 {
			fmt.Println("less 5")
		}
		fallthrough
	case i < 4:
		if i < 4 {
			fmt.Println("less 4")
		}
		fallthrough
	case i < 3:
		if i < 3 {
			fmt.Println("less 3")
		}
		fallthrough
	case i < 2:
		if i < 2 {
			fmt.Println("less 2")
		}
		fallthrough
	case i < 1:
		if i < 1 {
			fmt.Println("less 1")
		}
		fallthrough
	default:
		fmt.Printf("------ i=%d -- over----------\n", i)
	}
}

func main() {
	x := 0
	for x <= 5 {
		x++
		GetRank()
	}
}

Output:
less 6
less 5
less 4
------ i=3 -- over----------
less 6
less 5
less 4
less 3
less 2
less 1
------ i=0 -- over----------
less 6
less 5
less 4
less 3
less 2
less 1
------ i=0 -- over----------
i is 6
------ i=6 -- over----------
less 6
less 5
less 4
less 3
less 2
less 1
------ i=0 -- over----------
less 6
less 5
less 4
less 3
less 2
------ i=1 -- over----------
```

{{< /tab >}}

{{< /tabpane >}}

### Defer statements 语句 defer

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a [return statement](https://go.dev/ref/spec#Return_statements), reached the end of its [function body](https://go.dev/ref/spec#Function_declarations), or because the corresponding goroutine is [panicking](https://go.dev/ref/spec#Handling_panics).

​	"`defer`"语句调用一个函数，该函数的执行被推迟到外层函数返回的那一刻，或者是因为外层函数执行了 [return 语句](#return-statements-return----return-语句)，达到了其[函数体](../DeclarationsAndScope#function-declarations-函数声明)的末端，或者是因为相应的goroutine正在[恐慌](../Built-inFunctions#handling-panics-处理恐慌)。

```
DeferStmt = "defer" Expression .
```

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for [expression statements](https://go.dev/ref/spec#Expression_statements).

​	这个表达式必须是一个函数或方法的调用；它不能是被圆括号括起来的。对内置函数的调用与[表达式语句](#expression-statements-表达式语句)一样受到限制。

> 个人注释
>
> ```go
> package main
> 
> import (
> 	"fmt"
> )
> 
> func f(x int) {
> 	fmt.Printf("x=%d\n", x)
> }
> func main() {
> 	defer (func() { fmt.Println("over") }()) // 报错：expression in defer must not be parenthesized
> }
> 
> ```

Each time a "defer" statement executes, the function value and parameters to the call are [evaluated as usual](https://go.dev/ref/spec#Calls) and saved anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred. That is, if the surrounding function returns through an explicit [return statement](https://go.dev/ref/spec#Return_statements), deferred functions are executed *after* any result parameters are set by that return statement but *before* the function returns to its caller. If a deferred function value evaluates to `nil`, execution [panics](https://go.dev/ref/spec#Handling_panics) when the function is invoked, not when the "defer" statement is executed.

​	每次执行 "`defer`"语句时，函数值和调用的参数[像往常一样被求值](../Expressions#order-of-evaluation-求值顺序)并重新保存，但实际的函数不会被调用。相反，`被延迟函数`在外层的函数返回之前立即被调用，`其顺序与它们被延迟的顺序相反`。也就是说，如果外层的函数通过一个显式的 [return 语句](#return-statements-return----return-语句)返回，`被延迟函数`在任何结果参数被该 `reurn 语句`设置后执行，`但在外层函数返回给其调用者之前`。如果`被延迟函数`值被求值为`nil`，那么在调用该被延迟函数时会出现执行[恐慌]()（而不是当 "`defer`"语句被执行时）。

> 个人注释
>
> ​	解释下“如果`被延迟函数`值被求值为`nil`，那么在调用该被延迟函数时会出现执行[恐慌]()（而不是当 "`defer`"语句被执行时）”：
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"math/rand"
> )
> 
> func f1(i any) {
> 	fmt.Printf("i=%v\n", i)
> }
> 
> func f2() any {
> 	i := rand.Intn(101)
> 	if i >= 60 {
> 		return i
> 	} else {
> 		return nil
> 	}
> }
> 
> func main() {
> 	defer f1(f2()) // 可以执行
> 
> 	var fn func()                        // 声明一个空函数变量
> 	fmt.Printf("%T,%v\n", fn, fn == nil) // func(),true
> 	
> 	//defer fn() // 报错：panic: runtime error: invalid memory address or nil pointer dereference
> 
> 	fmt.Println("main over") // main over
> }
> 
> ```
>
> 

For instance, if the deferred function is a [function literal](https://go.dev/ref/spec#Function_literals) and the surrounding function has [named result parameters](https://go.dev/ref/spec#Function_types) that are in scope within the literal, the deferred function may access and modify the result parameters before they are returned. If the deferred function has any return values, they are discarded when the function completes. (See also the section on [handling panics](https://go.dev/ref/spec#Handling_panics).)

​	例如，如果`被延迟函数`是一个[函数字面量](../Expressions#function-literals-函数字面量)，并且外层的函数有在该字面量的作用域内的[命名结果参数](../Types#function-types-函数型)，那么该`被延迟函数`可以在这些结果参数被返回之前访问和修改它们。如果`被延迟函数`有任何返回值，这些返回值将在函数完成时被丢弃。(参见[处理恐慌](../Built-inFunctions#handling-panics-处理恐慌)一节)。

```go 
lock(l)
defer unlock(l)  // unlock 发生在外层函数返回之前

// 在外层函数返回之前，打印 3 2 1 0
for i := 0; i <= 3; i++ {
	defer fmt.Print(i)
}

// f 返回 42
func f() (result int) {
	defer func() {
		// result 会在它被 return 语句设为 6 之后再被访问
		result *= 7
	}()
	return 6
}
```

