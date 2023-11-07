+++
title = "template"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
https://pkg.go.dev/text/template@go1.21.3

Package template implements data-driven templates for generating textual output.

​	`template` 包实现了基于数据驱动的模板，用于生成文本输出。

To generate HTML output, see package html/template, which has the same interface as this package but automatically secures HTML output against certain attacks.

​	要生成 HTML 输出，请参考 html/template 包，它具有与此包相同的接口，但自动防止某些攻击来保护 HTML 输出。

Templates are executed by applying them to a data structure. Annotations in the template refer to elements of the data structure (typically a field of a struct or a key in a map) to control execution and derive values to be displayed. Execution of the template walks the structure and sets the cursor, represented by a period '.' and called "dot", to the value at the current location in the structure as execution proceeds.

​	模板通过将其应用于数据结构来执行。模板中的注释引用数据结构的元素（通常是结构体的字段或映射中的键），用于控制执行并派生要显示的值。模板的执行会遍历数据结构并设置游标，由句点 '`.`' 表示，称为 "dot"，在执行过程中将其设置为结构中当前位置的值。

The input text for a template is UTF-8-encoded text in any format. "Actions"--data evaluations or control structures--are delimited by "{{" and "}}"; all text outside actions is copied to the output unchanged.

​	模板的输入文本可以是任何格式的 UTF-8 编码文本。"操作"（数据求值或控制结构）由 "{{" 和 "}}" 分隔；所有在操作外的文本都将原样复制到输出中。

Once parsed, a template may be executed safely in parallel, although if parallel executions share a Writer the output may be interleaved.

​	一旦解析完成，模板可以在并发环境中安全地执行，但如果并发执行共享一个 Writer，则输出可能会交错。

Here is a trivial example that prints "17 items are made of wool".

​	下面是一个简单的示例，打印出 "17 items are made of wool"。

``` go 
type Inventory struct {
	Material string
	Count    uint
}
sweaters := Inventory{"wool", 17}
tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
if err != nil { panic(err) }
err = tmpl.Execute(os.Stdout, sweaters)
if err != nil { panic(err) }
```

More intricate examples appear below.

​	下面是更复杂的示例。

## 文本和空格 Text and spaces

By default, all text between actions is copied verbatim when the template is executed. For example, the string " items are made of " in the example above appears on standard output when the program is run.

​	默认情况下，在操作之间的所有文本在模板执行时都会被原样复制。例如，在上面的示例中，字符串 " items are made of" 在程序运行时会显示在标准输出中。

However, to aid in formatting template source code, if an action's left delimiter (by default "{{") is followed immediately by a minus sign and white space, all trailing white space is trimmed from the immediately preceding text. Similarly, if the right delimiter ("}}") is preceded by white space and a minus sign, all leading white space is trimmed from the immediately following text. In these trim markers, the white space must be present: "{{- 3}}" is like "{{3}}" but trims the immediately preceding text, while "{{-3}}" parses as an action containing the number -3.

​	然而，为了帮助格式化模板源代码，如果一个操作的`左定界符`（默认为 "{{"）紧跟着一个减号和空格，那么紧随其后的所有尾随空白都会被从前面的文本中去除。类似地，如果`右定界符`（"}}"）之前有空白和减号，那么紧随其后的所有前导空白都会被从后面的文本中去除。在这些修剪标记中，空白必须存在："{{- 3}}" 类似于 "{{3}}"，但会去除紧随其前面的文本的空白，而 "{{-3}}" 解析为包含数字 -3 的操作。

For instance, when executing the template whose source is

​	例如，当执行源代码为以下模板时：

```
"{{23 -}} < {{- 45}}"
```

the generated output would be

生成的输出将会是：

```
"23<45"
```

For this trimming, the definition of white space characters is the same as in Go: space, horizontal tab, carriage return, and newline.

​	对于这种修剪操作，空白字符的定义与 Go 语言相同：空格、水平制表符、回车和换行符。

## 操作 Actions

Here is the list of actions. "Arguments" and "pipelines" are evaluations of data, defined in detail in the corresponding sections that follow.

​	以下是操作的列表。"实参"和"管道"是数据的求值，在接下来的相应章节中有详细定义。

```
{{/* a comment */}}
{{- /* a comment with white space trimmed from preceding and following text */ -}}
	A comment; discarded. May contain newlines.
	Comments do not nest and must start and end at the
	delimiters, as shown here.
	注释；被丢弃。可以包含换行符。
	注释不嵌套，必须以示例中所示的分隔符开头和结尾。

{{pipeline}}
	The default textual representation (the same as would be
	printed by fmt.Print) of the value of the pipeline is copied
	to the output.
	管道值的默认文本表示（与使用 fmt.Print 打印的相同）被复制到输出。

{{if pipeline}} T1 {{end}}
	If the value of the pipeline is empty, no output is generated;
	otherwise, T1 is executed. The empty values are false, 0, any
	nil pointer or interface value, and any array, slice, map, or
	string of length zero.
	Dot is unaffected.
	如果管道的值为空，则不生成输出；
	否则，执行 T1。空值包括 false、0、任何 nil 指针或接口值，以及长度为零的数组、切片、映射或字符串。

{{if pipeline}} T1 {{else}} T0 {{end}}
	If the value of the pipeline is empty, T0 is executed;
	otherwise, T1 is executed. Dot is unaffected.
	如果管道的值为空，则执行 T0；否则，执行 T1。Dot 不受影响。

{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
	To simplify the appearance of if-else chains, the else action
	of an if may include another if directly; the effect is exactly
	the same as writing
	为了简化 if-else 链的外观，if 的 else 动作可以直接包含另一个 if；效果相同于使用如下的编写
{{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}

{{range pipeline}} T1 {{end}}
	The value of the pipeline must be an array, slice, map, or channel.
	If the value of the pipeline has length zero, nothing is output;
	otherwise, dot is set to the successive elements of the array,
	slice, or map and T1 is executed. If the value is a map and the
	keys are of basic type with a defined order, the elements will be
	visited in sorted key order.
	管道的值必须是数组、切片、映射或通道。
	如果管道的值长度为零，则不输出任何内容；
	否则，dot 依次设置为数组、切片或映射的元素，然后执行 T1。
	如果值是映射且键是具有定义顺序的基本类型，则元素将按排序的键顺序访问。

{{range pipeline}} T1 {{else}} T0 {{end}}
	The value of the pipeline must be an array, slice, map, or channel.
	If the value of the pipeline has length zero, dot is unaffected and
	T0 is executed; otherwise, dot is set to the successive elements
	of the array, slice, or map and T1 is executed.
	管道的值必须是数组、切片、映射或通道。
	如果管道的值长度为零，dot 不受影响并执行 T0；
	否则，dot 依次设置为数组、切片或映射的元素，然后执行 T1。

{{break}}
	The innermost {{range pipeline}} loop is ended early, stopping the
	current iteration and bypassing all remaining iterations.
	最内层的 {{range pipeline}} 循环被提前结束，停止当前迭代并跳过所有剩余的迭代。
	

{{continue}}
	The current iteration of the innermost {{range pipeline}} loop is
	stopped, and the loop starts the next iteration.
	最内层的 {{range pipeline}} 循环的当前迭代被停止，循环开始下一次迭代。

{{template "name"}}
	The template with the specified name is executed with nil data.
	执行具有指定名称的模板，数据为 nil。

{{template "name" pipeline}}
	The template with the specified name is executed with dot set
	to the value of the pipeline.
	执行具有指定名称的模板，dot 设置为管道的值。

{{block "name" pipeline}} T1 {{end}}
	A block is shorthand for defining a template
	块是定义模板的缩写
		{{define "name"}} T1 {{end}}
	and then executing it in place
	然后在此处执行它
		{{template "name" pipeline}}
	The typical use is to define a set of root templates that are
	then customized by redefining the block templates within.
	典型用法是定义一组根模板，然后通过重新定义块模板进行定制。

{{with pipeline}} T1 {{end}}
	If the value of the pipeline is empty, no output is generated;
	otherwise, dot is set to the value of the pipeline and T1 is
	executed.
	如果管道的值为空，则不生成输出；否则，dot 设置为管道的值，然后执行 T1。

{{with pipeline}} T1 {{else}} T0 {{end}}
	If the value of the pipeline is empty, dot is unaffected and T0
	is executed; otherwise, dot is set to the value of the pipeline
	and T1 is executed.
	如果管道的值为空，dot 不受影响并执行 T0；否则，dot 设置为管道的值，然后执行 T1。
```

## 实参 Arguments

An argument is a simple value, denoted by one of the following.

​	实参是一个简单的值，由以下之一表示。

- A boolean, string, character, integer, floating-point, imaginary or complex constant in Go syntax. These behave like Go's untyped constants. Note that, as in Go, whether a large integer constant overflows when assigned or passed to a function can depend on whether the host machine's ints are 32 or 64 bits.
- 一个布尔值、字符串、字符、整数、浮点数、虚数或复数常量，采用 Go 语法。它们的行为类似于 Go 的无类型常量。需要注意的是，与 Go 语言中一样，当大整数常量被赋值或传递给函数时，是否会溢出取决于主机机器的 int 类型是 32 位还是 64 位。
- The keyword nil, representing an untyped Go nil.
- 关键字 nil，表示无类型的 Go nil。
- The character '.' (period): . The result is the value of dot.
- 字符 '.'（句点）：. 结果是 dot 的值。
- A variable name, which is a (possibly empty) alphanumeric string preceded by a dollar sign, such as `$piOver2` or `$` The result is the value of the variable. Variables are described below.
- 变量名，是一个（可能为空的）由美元符号前缀的字母数字字符串，例如 `$piOver2` 或 `$`。结果是变量的值。变量将在下面进行描述。
- The name of a field of the data, which must be a struct, preceded by a period, such as `.Field` The result is the value of the field. Field invocations may be chained: `.Field1.Field2` . Fields can also be evaluated on variables, including chaining: `$x.Field1.Field2`
- 数据的字段名，必须是一个结构体，以句点（.）开头，例如 `.Field`。结果是字段的值。字段调用可以链式调用：`.Field1.Field2` 。字段还可以在变量上进行求值，包括链式调用：`$x.Field1.Field2`
- The name of a key of the data, which must be a map, preceded by a period, such as .Key The result is the map element value indexed by the key. Key invocations may be chained and combined with fields to any depth: `.Field1.Key1.Field2.Key2` Although the key must be an alphanumeric identifier, unlike with field names they do not need to start with an upper case letter. Keys can also be evaluated on variables, including chaining: `$x.key1.key2`
- 数据的键名，必须是一个映射，以句点（.）开头，例如 .Key。结果是由键索引的映射元素的值。键调用可以链式调用，并与字段组合到任意深度：`.Field1.Key1.Field2.Key2` 尽管键必须是字母数字标识符，**但与字段名不同，它们不需要以大写字母开头**。键还可以在变量上进行求值，包括链式调用：`$x.key1.key2`
- The name of a niladic method of the data, preceded by a period, such as .Method The result is the value of invoking the method with dot as the receiver, `dot.Method()`. Such a method must have one return value (of any type) or two return values, the second of which is an error. If it has two and the returned error is non-nil, execution terminates and an error is returned to the caller as the value of Execute. Method invocations may be chained and combined with fields and keys to any depth: `.Field1.Key1.Method1.Field2.Key2.Method2` Methods can also be evaluated on variables, including chaining: `$x.Method1.Field`
- 数据的零参数方法名称，以句点（.）开头，例如 `.Method`。结果是使用 dot 作为接收器调用方法的值，`dot.Method()`。这样的方法必须具有一个返回值（任意类型）或两个返回值，第二个返回值是一个error。如果有两个返回值且返回的error非nil，则执行终止，并将error作为 Execute 的返回值返回给调用者。方法调用可以链式调用，并与字段和键组合到任意深度：`.Field1.Key1.Method1.Field2.Key2.Method2` 方法还可以在变量上进行求值，包括链式调用：`$x.Method1.Field`
- The name of a niladic function, such as `fun` The result is the value of invoking the function, `fun()`. The return types and values behave as in methods. Functions and function names are described below.
- 零参数函数的名称，例如 `fun`。结果是调用函数的值，`fun()`。返回类型和值的行为与方法相同。函数和函数名将在下面进行描述。
- A parenthesized instance of one the above, for grouping. The result may be accessed by a field or map key invocation. `print (.F1 arg1) (.F2 arg2) (.StructValuedMethod "arg").Field`
- 上述内容的括号实例，用于分组。结果可以通过字段或映射键调用进行访问。例如 `print (.F1 arg1) (.F2 arg2) (.StructValuedMethod "arg").Field`

Arguments may evaluate to any type; if they are pointers the implementation automatically indirects to the base type when required. If an evaluation yields a function value, such as a function-valued field of a struct, the function is not invoked automatically, but it can be used as a truth value for an `if` action and the like. To invoke it, use the `call` function, defined below.

​	实参可以求值为任何类型；如果它们是指针，实现会在需要时自动对其基本类型进行间接引用。如果一个求值结果是函数值，例如结构体的函数值字段，该函数不会自动调用，但它可以用作`if`操作等的真值。要调用它，请使用下文中定义的`call`函数。

## 管道 Pipelines 

A pipeline is a possibly chained sequence of "commands". A command is a simple value (argument) or a function or method call, possibly with multiple arguments:

​	管道是一个可以被链式组合的"命令"序列。命令可以是一个简单的值（实参）或一个函数或方法调用，可以带有多个实参：

```
Argument
	The result is the value of evaluating the argument.
	结果是对实参进行求值的值。
.Method [Argument...]
	The method can be alone or the last element of a chain but,
	unlike methods in the middle of a chain, it can take arguments.
	The result is the value of calling the method with the
	arguments:
	该方法可以独立存在，也可以作为链的最后一个元素，
	但与链中间的方法不同，它可以接受实参。
	结果是使用实参调用该方法的值：
		dot.Method(Argument1, etc.)
functionName [Argument...]
	The result is the value of calling the function associated
	with the name:
	结果是调用与名称相关联的函数的值：
		function(Argument1, etc.)
	Functions and function names are described below.
	函数和函数名称将在下文描述。
```

A pipeline may be "chained" by separating a sequence of commands with pipeline characters '|'. In a chained pipeline, the result of each command is passed as the last argument of the following command. The output of the final command in the pipeline is the value of the pipeline.

​	通过使用管道字符 '|' 将一系列命令连接起来，可以"链式"组合管道。在链式管道中，每个命令的结果都作为下一个命令的最后一个实参传递。管道中最后一个命令的输出就是管道的值。

The output of a command will be either one value or two values, the second of which has type error. If that second value is present and evaluates to non-nil, execution terminates and the error is returned to the caller of `Execute`.

​	命令的输出将是一个值或两个值，其中第二个值的类型为 error。如果第二个值存在且求值为非nil，则执行终止，并将error返回给 `Execute` 的调用者。

## 变量

A pipeline inside an action may initialize a variable to capture the result. The initialization has syntax

​	在操作内部的管道中可以初始化一个变量来捕获结果。初始化的语法如下：

```
$variable := pipeline
```

where $variable is the name of the variable. An action that declares a variable produces no output.

其中 `$variable` 是变量的名称。声明变量的动作不会产生输出。

Variables previously declared can also be assigned, using the syntax

​	之前声明的变量也可以被赋值，使用的语法是：

```
$variable = pipeline
```

If a "range" action initializes a variable, the variable is set to the successive elements of the iteration. Also, a "range" may declare two variables, separated by a comma:

​	如果"range"操作初始化了一个变量，该变量将被设置为迭代的连续元素。此外，"range"还可以声明两个变量，用逗号分隔：

```
range $index, $element := pipeline
```

in which case `$index` and `$element` are set to the successive values of the array/slice index or map key and element, respectively. Note that if there is only one variable, it is assigned the element; this is opposite to the convention in Go range clauses.

​	在这种情况下，`$index` 和 `$element` 分别设置为数组/切片索引或映射键和元素的连续值。请注意，**如果只有一个变量，它将被赋值为元素**；**这与 Go range 语句的惯例相反**。

A variable's scope extends to the "end" action of the control structure ("if", "with", or "range") in which it is declared, or to the end of the template if there is no such control structure. A template invocation does not inherit variables from the point of its invocation.

​	变量的作用域延伸到控制结构（"if"、"with" 或 "range"）的 "end" 操作，或者如果没有这样的控制结构，则作用域延伸到模板的末尾。模板调用不会从其调用点继承变量。

When execution begins, $ is set to the data argument passed to Execute, that is, to the starting value of dot.

​	当执行开始时，`$` 设置为传递给 `Execute` 的数据实参，即 dot 的初始值。

Here are some example one-line templates demonstrating pipelines and variables. All produce the quoted word "output":

​	以下是一些示例的一行模板，演示了管道和变量。它们都会产生引号中的单词 "output"。

```
{{"\"output\""}}
	A string constant.
	一个字符串常量。
{{`"output"`}}
	A raw string constant.
	一个原始字符串常量。
{{printf "%q" "output"}}
	A function call.
	一个函数调用。
{{"output" | printf "%q"}}
	A function call whose final argument comes from the previous command.
	一个函数调用，其最后一个参数来自前一个命令。
{{printf "%q" (print "out" "put")}}
	A parenthesized argument.
	一个带括号的参数。
{{"put" | printf "%s%s" "out" | printf "%q"}}
	A more elaborate call.
	一个更复杂的调用。
{{"output" | printf "%s" | printf "%q"}}
	A longer chain.
	一个更长的链。
{{with "output"}}{{printf "%q" .}}{{end}}
	A with action using dot.
	使用点的with操作。
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
	A with action that creates and uses a variable.
	创建和使用变量的with操作。
{{with $x := "output"}}{{printf "%q" $x}}{{end}}
	A with action that uses the variable in another action.
	在另一个操作中使用变量的with操作。
{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
	The same, but pipelined.
	相同，但使用管道。
```

## 函数

During execution functions are found in two function maps: first in the template, then in the global function map. By default, no functions are defined in the template but the Funcs method can be used to add them.

​	在执行过程中，函数可以在两个函数映射中找到：首先在模板中，然后在全局函数映射中。默认情况下，模板中未定义任何函数，但可以使用 `Funcs` 方法来添加函数。

Predefined global functions are named as follows.

​	预定义的全局函数的名称如下所示。

```
and
	Returns the boolean AND of its arguments by returning the
	first empty argument or the last argument. That is,
	"and x y" behaves as "if x then y else x."
	Evaluation proceeds through the arguments left to right
	and returns when the result is determined.
	返回其实参的布尔AND，通过返回第一个空实参或最后一个实参来实现。
	也就是说，"and x y"的行为类似于"if x then y else x"。
	求值按从左到右的实参进行，当确定结果时返回。
call
	Returns the result of calling the first argument, which
	must be a function, with the remaining arguments as parameters.
	Thus "call .X.Y 1 2" is, in Go notation, dot.X.Y(1, 2) where
	Y is a func-valued field, map entry, or the like.
	The first argument must be the result of an evaluation
	that yields a value of function type (as distinct from
	a predefined function such as print). The function must
	return either one or two result values, the second of which
	is of type error. If the arguments don't match the function
	or the returned error value is non-nil, execution stops.
	返回调用第一个实参（必须是函数）并将其余实参作为实参的结果。
	因此，"call .X.Y 1 2" 在Go表示中等同于dot.X.Y(1, 2)，
	其中Y是一个函数值字段、映射条目或类似的东西。
	第一个实参必须是通过求值获得的函数类型值（与预定义函数（如print）不同）。
	函数必须返回一个或两个结果值，第二个结果值的类型是error。
	如果实参不匹配函数或返回的错误值不为nil，则执行停止。
html
	Returns the escaped HTML equivalent of the textual
	representation of its arguments. This function is unavailable
	in html/template, with a few exceptions.
	返回其实参文本表示的转义HTML等效项。在html/template中，此函数不可用，有几个例外。
index
	Returns the result of indexing its first argument by the
	following arguments. Thus "index x 1 2 3" is, in Go syntax,
	x[1][2][3]. Each indexed item must be a map, slice, or array.
	返回通过以下实参索引其第一个实参的结果。
	因此，"index x 1 2 3" 在Go语法中表示x[1][2][3]。
	每个被索引的项必须是映射、切片或数组。
slice
	slice returns the result of slicing its first argument by the
	remaining arguments. Thus "slice x 1 2" is, in Go syntax, x[1:2],
	while "slice x" is x[:], "slice x 1" is x[1:], and "slice x 1 2 3"
	is x[1:2:3]. The first argument must be a string, slice, or array.
	slice返回通过其第一个实参和剩余实参切片的结果。
	因此，"slice x 1 2" 在Go语法中表示x[1:2]，
	而"slice x" 表示x[:]，"slice x 1" 表示x[1:]，
	而"slice x 1 2 3" 表示x[1:2:3]。第一个实参必须是字符串、切片或数组。
js
	Returns the escaped JavaScript equivalent of the textual
	representation of its arguments.
	返回其实参文本表示的转义JavaScript等效项。
len
	Returns the integer length of its argument.
	返回其实参的整数长度。
not
	Returns the boolean negation of its single argument.
	返回其单个实参的布尔否定。
or
	Returns the boolean OR of its arguments by returning the
	first non-empty argument or the last argument, that is,
	"or x y" behaves as "if x then x else y".
	Evaluation proceeds through the arguments left to right
	and returns when the result is determined.
	返回其实参的布尔OR，通过返回第一个非空实参或最后一个实参来实现，
	也就是说，"or x y"的行为类似于"if x then x else y"。
	求值按从左到右的实参进行，当确定结果时返回。
print
	An alias for fmt.Sprint
	fmt.Sprint的别名
printf
	An alias for fmt.Sprintf
	fmt.Sprintf的别名
println
	An alias for fmt.Sprintln
	fmt.Sprintln的别名
urlquery
	Returns the escaped value of the textual representation of
	its arguments in a form suitable for embedding in a URL query.
	This function is unavailable in html/template, with a few
	exceptions.
	返回其实参文本表示的转义值，适合嵌入URL查询中。
	在html/template中，此函数不可用，有几个例外。
```

The boolean functions take any zero value to be false and a non-zero value to be true.

​	布尔函数将任何零值视为 false，非零值视为 true。

There is also a set of binary comparison operators defined as functions:

​	还有一组作为函数定义的二元比较运算符：

```
eq
	Returns the boolean truth of arg1 == arg2
	返回arg1 == arg2的布尔真值
ne
	Returns the boolean truth of arg1 != arg2
	返回arg1 != arg2的布尔真值
lt
	Returns the boolean truth of arg1 < arg2	
	返回arg1 < arg2的布尔真值
le
	Returns the boolean truth of arg1 <= arg2
	返回arg1 <= arg2的布尔真值
gt
	Returns the boolean truth of arg1 > arg2
	返回arg1 > arg2的布尔真值
ge
	Returns the boolean truth of arg1 >= arg2
	返回arg1 >= arg2的布尔真值
```

For simpler multi-way equality tests, eq (only) accepts two or more arguments and compares the second and subsequent to the first, returning in effect

​	对于更简单的多路相等性测试，eq（仅限）接受两个或多个实参，并将第二个及后续实参与第一个实参进行比较，从而返回

```
arg1==arg2 || arg1==arg3 || arg1==arg4 ...
```

(Unlike with || in Go, however, eq is a function call and all the arguments will be evaluated.)

（与 Go 中的 || 不同，eq 是一个函数调用，所有实参都将被求值。）

The comparison functions work on any values whose type Go defines as comparable. For basic types such as integers, the rules are relaxed: size and exact type are ignored, so any integer value, signed or unsigned, may be compared with any other integer value. (The arithmetic value is compared, not the bit pattern, so all negative integers are less than all unsigned integers.) However, as usual, one may not compare an int with a float32 and so on.

​	比较函数适用于 Go 定义为可比较的任何值。对于基本类型（例如整数），规则被放宽：大小和精确类型被忽略，因此任何整数值（有符号或无符号）可以与任何其他整数值进行比较。（比较的是算术值，而不是位模式，因此所有负整数都小于所有无符号整数。）但是，通常情况下，不能将 int 与 float32 等进行比较。

### 关联模板 Associated templates  

Each template is named by a string specified when it is created. Also, each template is associated with zero or more other templates that it may invoke by name; such associations are transitive and form a name space of templates.

​	每个模板都有一个在创建时指定的字符串名称。此外，每个模板与零个或多个其他模板相关联，可以通过名称调用这些模板；这种关联是传递性的，并形成模板的名称空间。

A template may use a template invocation to instantiate another associated template; see the explanation of the "template" action above. The name must be that of a template associated with the template that contains the invocation.

​	模板可以使用模板调用来实例化另一个关联的模板；请参考上面的"template"操作的解释。名称必须是包含调用的模板所关联的模板的名称。

### 嵌套模板定义 Nested template definitions 

When parsing a template, another template may be defined and associated with the template being parsed. Template definitions must appear at the top level of the template, much like global variables in a Go program.

​	在解析模板时，可以定义另一个模板并将其与正在解析的模板关联。模板定义必须出现在模板的顶层，类似于Go程序中的全局变量。 

The syntax of such definitions is to surround each template declaration with a "define" and "end" action.

​	这种定义的语法是用"define"和"end"操作将每个模板声明括起来。 

The define action names the template being created by providing a string constant. Here is a simple example:

​	通过提供一个字符串常量，`define`操作命名了正在创建的模板。以下是一个简单的示例：

```
{{define "T1"}}ONE{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
{{template "T3"}}
```

This defines two templates, T1 and T2, and a third T3 that invokes the other two when it is executed. Finally it invokes T3. If executed this template will produce the text

​	这段代码定义了两个模板T1和T2，以及第三个模板T3，当T3被执行时，它会调用T1和T2。最后，代码调用了T3。如果执行这个模板，将会产生以下文本：

```
ONE TWO
```

By construction, a template may reside in only one association. If it's necessary to have a template addressable from multiple associations, the template definition must be parsed multiple times to create distinct *Template values, or must be copied with the Clone or AddParseTree method.

​	根据构建，一个模板只能存在于一个关联中。如果需要从多个关联中访问一个模板，必须多次解析模板定义以创建不同的`*Template`值，或者必须使用`Clone`或`AddParseTree`方法进行复制。

Parse may be called multiple times to assemble the various associated templates; see the ParseFiles and ParseGlob functions and methods for simple ways to parse related templates stored in files.

​	可以多次调用`Parse`来组装各种关联的模板；有关解析存储在文件中的相关模板的简便方法，请参阅`ParseFiles`和`ParseGlob`函数和方法。

A template may be executed directly or through ExecuteTemplate, which executes an associated template identified by name. To invoke our example above, we might write,

​	一个模板可以直接执行，也可以通过`ExecuteTemplate`执行由名称标识的相关模板。要调用我们上面的示例，我们可以这样写：

```go
err := tmpl.Execute(os.Stdout, "no data needed")
if err != nil {
	log.Fatalf("execution failed: %s", err)
}
```

or to invoke a particular template explicitly by name,

或者通过名称显式调用特定的模板：

```go
err := tmpl.ExecuteTemplate(os.Stdout, "T2", "no data needed")
if err != nil {
	log.Fatalf("execution failed: %s", err)
}
```




## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func HTMLEscape 

``` go 
func HTMLEscape(w io.Writer, b []byte)
```

HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.

​	`HTMLEscape`函数向`w`写入`b`的纯文本数据的等效转义HTML。

### func HTMLEscapeString 

``` go 
func HTMLEscapeString(s string) string
```

HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.

​	`HTMLEscapeString`函数返回`s`的纯文本数据的等效转义HTML。

### func HTMLEscaper 

``` go 
func HTMLEscaper(args ...any) string
```

HTMLEscaper returns the escaped HTML equivalent of the textual representation of its arguments.

​	`HTMLEscaper`函数返回其实参的文本表示形式的等效转义HTML。

### func IsTrue  <- go1.6

``` go 
func IsTrue(val any) (truth, ok bool)
```

IsTrue reports whether the value is 'true', in the sense of not the zero of its type, and whether the value has a meaningful truth value. This is the definition of truth used by if and other such actions.

​	`IsTrue`函数报告值（val）是否为'true'，即不是其类型的零值，以及该值是否具有有意义的真值。 这是`if`和其他此类操作中使用的真值定义。

### func JSEscape 

``` go 
func JSEscape(w io.Writer, b []byte)
```

JSEscape writes to w the escaped JavaScript equivalent of the plain text data b.

​	`JSEscape`函数向`w`写入`b`的纯文本数据的等效转义JavaScript。

### func JSEscapeString 

``` go 
func JSEscapeString(s string) string
```

JSEscapeString returns the escaped JavaScript equivalent of the plain text data s.

​	`JSEscapeString`函数返回`s`的纯文本数据的等效转义JavaScript。

### func JSEscaper 

``` go 
func JSEscaper(args ...any) string
```

JSEscaper returns the escaped JavaScript equivalent of the textual representation of its arguments.

​	`JSEscaper`函数返回其实参的文本表示形式的等效转义JavaScript。

### func URLQueryEscaper 

``` go 
func URLQueryEscaper(args ...any) string
```

URLQueryEscaper returns the escaped value of the textual representation of its arguments in a form suitable for embedding in a URL query.

​	`URLQueryEscaper`函数返回其实参的文本表示形式的转义值，适合于嵌入URL查询中。

## 类型

### type ExecError  <- go1.6

``` go 
type ExecError struct {
	Name string // Name of template.
	Err  error  // Pre-formatted error.
}
```

ExecError is the custom error type returned when Execute has an error evaluating its template. (If a write error occurs, the actual error is returned; it will not be of type ExecError.)

​	`ExecError` 是在执行模板时发生错误时返回的自定义错误类型。（如果发生写入错误，则返回实际错误；它不会是 ExecError 类型。）

#### (ExecError) Error  <- go1.6

``` go 
func (e ExecError) Error() string
```

#### (ExecError) Unwrap  <- go1.13

``` go 
func (e ExecError) Unwrap() error
```

### type FuncMap 

``` go 
type FuncMap map[string]any
```

FuncMap is the type of the map defining the mapping from names to functions. Each function must have either a single return value, or two return values of which the second has type error. In that case, if the second (error) return value evaluates to non-nil during execution, execution terminates and `Execute` returns that error.

​	`FuncMap` 是定义从名称到函数的映射的映射类型。每个函数必须具有单个返回值或两个返回值，其中第二个返回值的类型为 error。在这种情况下，如果第二个（error）返回值在执行期间求值为非 nil，则执行终止，并且 `Execute` 返回该错误。

Errors returned by Execute wrap the underlying error; call errors.As to uncover them.

​	`Execute` 返回的错误会封装底层错误；可以使用 `errors.As` 解析它们。

When template execution invokes a function with an argument list, that list must be assignable to the function's parameter types. Functions meant to apply to arguments of arbitrary type can use parameters of type interface{} or of type reflect.Value. Similarly, functions meant to return a result of arbitrary type can return interface{} or reflect.Value.

​	当模板执行调用具有参数列表的函数时，该列表必须可分配给函数的参数类型。 意在应用于任意类型参数的函数可以使用类型为`interface{}`或`reflect.Value`的参数。 同样，意在返回任意类型结果的函数可以返回`interface{}`或`reflect.Value`。

### type Template 

``` go 
type Template struct {
	*parse.Tree
	// contains filtered or unexported fields
}
```

Template is the representation of a parsed template. The *parse.Tree field is exported only for use by html/template and should be treated as unexported by all other clients.

​	Template 是已解析模板的表示形式。`*parse.Tree` 字段只对 `html/template` 使用者公开，其他所有客户端应将其视为未导出的字段。

#### Template Example
``` go 
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Prepare some data to insert into the template.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}

//Output:
//
//Dear Aunt Mildred,
//
//It was a pleasure to see you at the wedding.
//Thank you for the lovely bone china tea set.
//
//Best wishes,
//Josie
//
//Dear Uncle John,
//
//It is a shame you couldn't make it to the wedding.
//Thank you for the lovely moleskin pants.

//Best wishes,
//Josie
//
//Dear Cousin Rodney,
//
//It is a shame you couldn't make it to the wedding.
//
//Best wishes,
//Josie
```

#### Template Example (Block)
``` go 
//package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}

//Output:
//
//Names:
//- Gamora
//- Groot
//- Nebula
//- Rocket
//- Star-Lord
//Names: Gamora, Groot, Nebula, Rocket, Star-Lord
```

#### Template Example (Func)

This example demonstrates a custom function to process template text. It installs the strings.Title function and uses it to Make Title Text Look Good In Our Template's Output.

​	这个示例演示了处理模板文本的自定义函数。它安装了 `strings.Title` 函数并使用它来使标题文本在我们的模板输出中看起来好看。

``` go 
package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	// First we create a FuncMap with which to register the function.
    // 首先，我们创建一个FuncMap以注册函数。
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
        // “title”是模板文本中函数的名称。
		"title": strings.Title,
	}

	// A simple template definition to test our function.
    // 一个用于测试函数的简单模板定义
	// We print the input text several ways:
    //  我们用几种方式打印输入文本：
	// - the original
	// - title-cased
	// - title-cased and then printed with %q
	// - printed with %q and then title-cased.
	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
`

	// Create a template, add the function map, and parse the text.
    // 创建一个模板，添加函数映射，解析文本。
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
    // 运行模板以验证输出。
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}

//Output:
//
//Input: "the go programming language"
//Output 0: The Go Programming Language
//Output 1: "The Go Programming Language"
//Output 2: "The Go Programming Language"
```

#### Template Example(Glob)

Here we demonstrate loading a set of templates from a directory.

​	在这里，我们演示如何从目录中加载一组模板。

``` go 
package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// templateFile defines the contents of a template to be stored in a file, for testing.
// templateFile 定义了存储在文件中的模板内容，用于测试。
type templateFile struct {
	name     string
	contents string
}

func createTestDir(files []templateFile) string {
	dir, err := os.MkdirTemp("", "template")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func main() {
	// Here we create a temporary directory and populate it with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
    // 在这里，我们创建一个临时目录，
    // 并向其中添加示例模板定义文件；
    // 通常，模板文件已经存在于程序已知的某个位置。
	dir := createTestDir([]templateFile{
		// T0.tmpl is a plain template file that just invokes T1.
        // T0.tmpl 是一个简单的模板文件，它只是调用 T1。
		{"T0.tmpl", `T0 invokes T1: ({{template "T1"}})`},
		// T1.tmpl defines a template, T1 that invokes T2.
        // T1.tmpl 定义了一个模板 T1，这个模板会调用 T2。
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		// T2.tmpl defines a template T2.
        // T2.tmpl 定义了模板 T2。
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})
	// Clean up after the test; another quirk of running as an example.
    // 在测试之后进行清理；这是作为示例运行的另一个特点。
	defer os.RemoveAll(dir)

	// pattern is the glob pattern used to find all the template files.
    // pattern 是用于找到所有模板文件的Glob模式。
	pattern := filepath.Join(dir, "*.tmpl")

	// Here starts the example proper.
	// T0.tmpl is the first name matched, so it becomes the starting template,
	// the value returned by ParseGlob.
    // 以下是示例的开始部分：
    // T0.tmpl 是第一个匹配到的文件名，
    // 因此它成为起始模板，这是由 ParseGlob 返回的值。
	tmpl := template.Must(template.ParseGlob(pattern))

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

//Output:
//
//T0 invokes T1: (T1 invokes T2: (This is T2))
```

#### Template Example (Helpers) 

This example demonstrates one way to share some templates and use them in different contexts. In this variant we add multiple driver templates by hand to an existing bundle of templates.

``` go 
package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// templateFile defines the contents of a template to be stored in a file, for testing.
type templateFile struct {
	name     string
	contents string
}

func createTestDir(files []templateFile) string {
	dir, err := os.MkdirTemp("", "template")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func main() {
	// Here we create a temporary directory and populate it with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	dir := createTestDir([]templateFile{
		// T1.tmpl defines a template, T1 that invokes T2.
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		// T2.tmpl defines a template T2.
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})
	// Clean up after the test; another quirk of running as an example.
	defer os.RemoveAll(dir)

	// pattern is the glob pattern used to find all the template files.
	pattern := filepath.Join(dir, "*.tmpl")

	// Here starts the example proper.
	// Load the helpers.
	templates := template.Must(template.ParseGlob(pattern))
	// Add one driver template to the bunch; we do this with an explicit template definition.
	_, err := templates.Parse("{{define `driver1`}}Driver 1 calls T1: ({{template `T1`}})\n{{end}}")
	if err != nil {
		log.Fatal("parsing driver1: ", err)
	}
	// Add another driver template.
	_, err = templates.Parse("{{define `driver2`}}Driver 2 calls T2: ({{template `T2`}})\n{{end}}")
	if err != nil {
		log.Fatal("parsing driver2: ", err)
	}
	// We load all the templates before execution. This package does not require
	// that behavior but html/template's escaping does, so it's a good habit.
	err = templates.ExecuteTemplate(os.Stdout, "driver1", nil)
	if err != nil {
		log.Fatalf("driver1 execution: %s", err)
	}
	err = templates.ExecuteTemplate(os.Stdout, "driver2", nil)
	if err != nil {
		log.Fatalf("driver2 execution: %s", err)
	}
}

//Output:
//
//Driver 1 calls T1: (T1 invokes T2: (This is T2))
//Driver 2 calls T2: (This is T2)
```

#### Template Example (Share)

This example demonstrates how to use one group of driver templates with distinct sets of helper templates.

``` go 
package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// templateFile defines the contents of a template to be stored in a file, for testing.
type templateFile struct {
	name     string
	contents string
}

func createTestDir(files []templateFile) string {
	dir, err := os.MkdirTemp("", "template")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func main() {
	// Here we create a temporary directory and populate it with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	dir := createTestDir([]templateFile{
		// T0.tmpl is a plain template file that just invokes T1.
		{"T0.tmpl", "T0 ({{.}} version) invokes T1: ({{template `T1`}})\n"},
		// T1.tmpl defines a template, T1 that invokes T2. Note T2 is not defined
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
	})
	// Clean up after the test; another quirk of running as an example.
	defer os.RemoveAll(dir)

	// pattern is the glob pattern used to find all the template files.
	pattern := filepath.Join(dir, "*.tmpl")

	// Here starts the example proper.
	// Load the drivers.
	drivers := template.Must(template.ParseGlob(pattern))

	// We must define an implementation of the T2 template. First we clone
	// the drivers, then add a definition of T2 to the template name space.

	// 1. Clone the helper set to create a new name space from which to run them.
	first, err := drivers.Clone()
	if err != nil {
		log.Fatal("cloning helpers: ", err)
	}
	// 2. Define T2, version A, and parse it.
	_, err = first.Parse("{{define `T2`}}T2, version A{{end}}")
	if err != nil {
		log.Fatal("parsing T2: ", err)
	}

	// Now repeat the whole thing, using a different version of T2.
	// 1. Clone the drivers.
	second, err := drivers.Clone()
	if err != nil {
		log.Fatal("cloning drivers: ", err)
	}
	// 2. Define T2, version B, and parse it.
	_, err = second.Parse("{{define `T2`}}T2, version B{{end}}")
	if err != nil {
		log.Fatal("parsing T2: ", err)
	}

	// Execute the templates in the reverse order to verify the
	// first is unaffected by the second.
	err = second.ExecuteTemplate(os.Stdout, "T0.tmpl", "second")
	if err != nil {
		log.Fatalf("second execution: %s", err)
	}
	err = first.ExecuteTemplate(os.Stdout, "T0.tmpl", "first")
	if err != nil {
		log.Fatalf("first: execution: %s", err)
	}

}

//Output:
//
//T0 (second version) invokes T1: (T1 invokes T2: (T2, version B))
//T0 (first version) invokes T1: (T1 invokes T2: (T2, version A))
```

#### func Must 

``` go 
func Must(t *Template, err error) *Template
```

Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil. It is intended for use in variable initializations such as

​	`Must` 是一个辅助函数，它包装了对返回(*Template, error)的函数的调用，如果错误不为 nil，它会引发 panic。它旨在用于变量初始化等，例如：

``` go 
var t = template.Must(template.New("name").Parse("text"))
```

#### func New 

``` go 
func New(name string) *Template
```

New allocates a new, undefined template with the given name.

​	`New`函数分配一个具有给定名称的新的未定义模板。

#### func ParseFS  <- go1.16

``` go 
func ParseFS(fsys fs.FS, patterns ...string) (*Template, error)
```

ParseFS is like ParseFiles or ParseGlob but reads from the file system fsys instead of the host operating system's file system. It accepts a list of glob patterns. (Note that most file names serve as glob patterns matching only themselves.)

​	`ParseFS`函数类似于`ParseFiles`函数或`ParseGlob`函数，但是它从文件系统`fsys`读取而不是从宿主操作系统的文件系统读取。它接受一个glob模式的列表。(注意，大多数文件名作为glob模式只匹配它们自己。)

#### func ParseFiles 

``` go 
func ParseFiles(filenames ...string) (*Template, error)
```

ParseFiles creates a new Template and parses the template definitions from the named files. The returned template's name will have the base name and parsed contents of the first file. There must be at least one file. If an error occurs, parsing stops and the returned `*Template` is nil.

​	`ParseFiles`函数创建一个新的Template并从指定文件中解析模板定义。返回的模板名称将具有第一个文件的基本名称和解析内容。必须至少有一个文件。如果发生错误，解析将停止并且返回的`*Template`为nil。

When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results. For instance, `ParseFiles("a/foo", "b/foo")` stores "b/foo" as the template named "foo", while "a/foo" is unavailable.

​	在解析不同目录中具有相同名称的多个文件时，最后提到的文件将是结果文件。例如，`ParseFiles("a/foo", "b/foo")`将"b/foo"存储为名为"foo"的模板，而"a/foo"将无法使用。

#### func ParseGlob 

``` go 
func ParseGlob(pattern string) (*Template, error)
```

ParseGlob creates a new Template and parses the template definitions from the files identified by the pattern. The files are matched according to the semantics of filepath.Match, and the pattern must match at least one file. The returned template will have the (base) name and (parsed) contents of the first file matched by the pattern. ParseGlob is equivalent to calling ParseFiles with the list of files matched by the pattern.

​	`ParseGlob`函数创建一个新的Template并从由模式标识的文件中解析模板定义。文件根据`filepath.Match`的语义进行匹配，并且模式必须匹配至少一个文件。返回的模板将具有由模式匹配的第一个文件的基本名称和解析内容。`ParseGlob`函数等同于使用由模式匹配的文件列表调用`ParseFiles`函数。

When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results.

​	在解析不同目录中具有相同名称的多个文件时，最后提到的文件将是结果文件。

#### (*Template) AddParseTree 

``` go 
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)
```

AddParseTree associates the argument parse tree with the template t, giving it the specified name. If the template has not been defined, this tree becomes its definition. If it has been defined and already has that name, the existing definition is replaced; otherwise a new template is created, defined, and returned.

​	`AddParseTree` 方法将实参解析树（`tree`）与模板 `t` 相关联，并为其指定名称。 如果模板未定义，则此树（`tree`）成为其定义。 如果已定义并且已经具有该名称，则现有定义将被替换； 否则将创建一个新的模板，定义并返回它。

#### (*Template) Clone 

``` go 
func (t *Template) Clone() (*Template, error)
```

Clone returns a duplicate of the template, including all associated templates. The actual representation is not copied, but the name space of associated templates is, so further calls to Parse in the copy will add templates to the copy but not to the original. Clone can be used to prepare common templates and use them with variant definitions for other templates by adding the variants after the clone is made.

​	`Clone` 方法返回模板的副本，包括所有关联的模板。 实际的表示形式不会被复制，但关联模板的名称空间会被复制，因此在副本中进一步调用 `Parse` 方法将向副本添加模板而不是原始模板。 `Clone` 方法可用于准备常见的模板，并在克隆后添加变体定义以将其用于其他具有变体定义的模板。

#### (*Template) DefinedTemplates  <- go1.5

``` go 
func (t *Template) DefinedTemplates() string
```

DefinedTemplates returns a string listing the defined templates, prefixed by the string "; defined templates are: ". If there are none, it returns the empty string. For generating an error message here and in html/template.

​	`DefinedTemplates` 方法返回一个字符串，列出已定义的模板，以 "; defined templates are: " 为前缀。 如果没有，则返回空字符串。 用于生成错误消息和 html/template 中的错误消息。

#### (*Template) Delims 

``` go 
func (t *Template) Delims(left, right string) *Template
```

Delims sets the action delimiters to the specified strings, to be used in subsequent calls to Parse, ParseFiles, or ParseGlob. Nested template definitions will inherit the settings. An empty delimiter stands for the corresponding default: {{ or }}. The return value is the template, so calls can be chained.

​	`Delims` 方法将操作定界符设置为指定的字符串，以在后续调用 `Parse`、`ParseFiles` 或 `ParseGlob` 中使用。 嵌套的模板定义将继承这些设置。 空定界符代表相应的默认值：{{ 或 }}。 返回值是模板本身，因此可以链式调用。

#### (*Template) Execute 

``` go 
func (t *Template) Execute(wr io.Writer, data any) error
```

Execute applies a parsed template to the specified data object, and writes the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel, although if parallel executions share a Writer the output may be interleaved.

​	`Execute` 方法将已解析的模板应用于指定的数据对象，并将输出写入 `wr`。 如果执行模板或写入其输出时发生错误，则执行停止，但部分结果可能已经写入输出写入器。 模板可以安全地并行执行，但如果并行执行共享一个 Writer，则输出可能会交错。

If data is a reflect.Value, the template applies to the concrete value that the reflect.Value holds, as in fmt.Print.

​	如果 `data` 是 `reflect.Value`，则模板适用于 `reflect.Value` 持有的具体值，就像 `fmt.Print` 一样。

#### (*Template) ExecuteTemplate 

``` go 
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error
```

ExecuteTemplate applies the template associated with t that has the given name to the specified data object and writes the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel, although if parallel executions share a Writer the output may be interleaved.

​	`ExecuteTemplate` 方法将与 `t` 关联且具有给定名称的模板应用于指定的数据对象，并将输出写入 `wr`。如果在执行模板或写入输出时发生错误，则执行将停止，但部分结果可能已经写入输出写入器。可以安全地并行执行模板，但如果并行执行共享一个 Writer，则输出可能会交错。

#### (*Template) Funcs 

``` go 
func (t *Template) Funcs(funcMap FuncMap) *Template
```

Funcs adds the elements of the argument map to the template's function map. It must be called before the template is parsed. It panics if a value in the map is not a function with appropriate return type or if the name cannot be used syntactically as a function in a template. It is legal to overwrite elements of the map. The return value is the template, so calls can be chained.

​	`Funcs` 方法将实参映射中的元素添加到模板的函数映射中。 它必须在模板解析之前被调用。 如果映射中的值不是具有适当返回类型的函数，或者如果名称不能在语法上用作模板中的函数，它会引发panic。 覆盖映射中的元素是合法的。 返回值是模板本身，因此可以链式调用。

#### (*Template) Lookup 

``` go 
func (t *Template) Lookup(name string) *Template
```

Lookup returns the template with the given name that is associated with t. It returns nil if there is no such template or the template has no definition.

​	`Lookup` 方法返回与 `t` 关联且具有给定名称的模板。如果没有这样的模板或模板没有定义，则返回 nil。

#### (*Template) Name 

``` go 
func (t *Template) Name() string
```

Name returns the name of the template.

​	`Name` 方法返回模板的名称。

#### (*Template) New 

``` go 
func (t *Template) New(name string) *Template
```

New allocates a new, undefined template associated with the given one and with the same delimiters. The association, which is transitive, allows one template to invoke another with a {{template}} action.

​	`New` 方法分配一个新的未定义模板，并与给定模板关联，并具有相同的定界符。这种关联是传递性的，允许一个模板使用 `{{template}}` 操作调用另一个模板。

Because associated templates share underlying data, template construction cannot be done safely in parallel. Once the templates are constructed, they can be executed in parallel.

​	由于相关联的模板共享底层数据，因此无法安全地在并行中构建模板。 一旦构建了模板，它们就可以并行执行。

#### (*Template) Option  <- go1.5

``` go 
func (t *Template) Option(opt ...string) *Template
```

Option sets options for the template. Options are described by strings, either a simple string or "key=value". There can be at most one equals sign in an option string. If the option string is unrecognized or otherwise invalid, Option panics.

​	`Option` 方法为模板设置选项。选项由字符串描述，可以是简单字符串或"key=value" 。选项字符串中最多只能有一个等号。如果选项字符串无法识别或无效，`Option` 方法会引发panic。

Known options:

已知选项：

​	missingkey: Control the behavior during execution if a map is indexed with a key that is not present in the map.

​	missingkey: 控制在执行过程中，如果使用在映射中不存在的键对映射进行索引的行为。

```
"missingkey=default" or "missingkey=invalid"
	The default behavior: Do nothing and continue execution.
	If printed, the result of the index operation is the string
	"<no value>".
	默认行为：不执行任何操作并继续执行。
	如果打印，索引操作的结果是字符串 "<no value>"。
"missingkey=zero"
	The operation returns the zero value for the map type's element.
	该操作返回 map 类型元素的零值。
"missingkey=error"
	Execution stops immediately with an error.
	立即因错误而停止执行。
```

#### (*Template) Parse 

``` go 
func (t *Template) Parse(text string) (*Template, error)
```

Parse parses text as a template body for t. Named template definitions ({{define ...}} or {{block ...}} statements) in text define additional templates associated with t and are removed from the definition of t itself.

​	`Parse` 方法将`text` 解析为 `t` 的模板主体。`text` 中的命名模板定义（`{{define ...}}` 或 `{{block ...}}` 语句）定义与 `t` 相关联的其他模板，并从 `t` 本身的定义中删除。

Templates can be redefined in successive calls to Parse. A template definition with a body containing only white space and comments is considered empty and will not replace an existing template's body. This allows using Parse to add new named template definitions without overwriting the main template body.

​	可以连续调用 `Parse` 方法来重新定义模板。如果模板定义的主体只包含空白和注释，则认为它是空的，并且不会替换现有模板的主体。这允许使用 `Parse` 方法添加新的命名模板定义，而不会覆盖主模板主体。

#### (*Template) ParseFS  <- go1.16

``` go 
func (t *Template) ParseFS(fsys fs.FS, patterns ...string) (*Template, error)
```

ParseFS is like ParseFiles or ParseGlob but reads from the file system fsys instead of the host operating system's file system. It accepts a list of glob patterns. (Note that most file names serve as glob patterns matching only themselves.)

​	`ParseFS` 方法类似于 `ParseFiles` 方法或 `ParseGlob`方法，但它从文件系统 `fsys` 读取而不是主机操作系统的文件系统。它接受一个 glob 模式列表。（请注意，大多数文件名仅作为匹配自己的 glob 模式。）

#### (*Template) ParseFiles 

``` go 
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
```

ParseFiles parses the named files and associates the resulting templates with t. If an error occurs, parsing stops and the returned template is nil; otherwise it is t. There must be at least one file. Since the templates created by ParseFiles are named by the base names of the argument files, t should usually have the name of one of the (base) names of the files. If it does not, depending on t's contents before calling ParseFiles, t.Execute may fail. In that case use t.ExecuteTemplate to execute a valid template.

​	`ParseFiles`方法解析指定的文件，并将生成的模板与t关联。如果发生错误，解析将停止，并返回的模板将为nil；否则它是`t`。必须至少有一个文件。由于由`ParseFiles`方法创建的模板是由实参文件的基本名称命名的，因此`t`通常应该有文件的一个（基本）名称。如果没有，取决于在调用`ParseFiles`方法之前`t`的内容，`t.Execute`可能会失败。在这种情况下，请使用`t.ExecuteTemplate`来执行有效的模板。

When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results.

​	在解析具有相同名称的不同目录中的多个文件时，最后提到的文件将是结果文件。

#### (*Template) ParseGlob 

``` go 
func (t *Template) ParseGlob(pattern string) (*Template, error)
```

ParseGlob parses the template definitions in the files identified by the pattern and associates the resulting templates with t. The files are matched according to the semantics of filepath.Match, and the pattern must match at least one file. ParseGlob is equivalent to calling t.ParseFiles with the list of files matched by the pattern.

​	`ParseGlob` 方法解析由模式标识的文件中定义的模板，并将结果模板与 `t` 相关联。根据 `filepath.Match` 的语义匹配文件，模式必须至少匹配一个文件。`ParseGlob` 方法等同于使用模式匹配的文件列表调用 `t.ParseFiles`方法。

When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results.

​	在解析具有相同名称的不同目录中的多个文件时，最后提到的文件将是结果文件

#### (*Template) Templates 

``` go 
func (t *Template) Templates() []*Template
```

Templates returns a slice of defined templates associated with t.

​	`Templates`方法 返回与 `t` 关联的已定义模板的切片。