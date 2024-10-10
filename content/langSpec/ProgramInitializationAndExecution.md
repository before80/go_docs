+++
title = "程序初始化和执行"
date = 2023-05-17T09:59:21+08:00
weight = 15
description = ""
isCJKLanguage = true
type = "docs"
draft = false
+++
## Program initialization and execution 程序初始化和执行

> 原文：[https://go.dev/ref/spec#Program_initialization_and_execution](https://go.dev/ref/spec#Program_initialization_and_execution)

### The zero value 零值

When storage is allocated for a [variable](https://go.dev/ref/spec#Variables), either through a declaration or a call of `new`, or when a new value is created, either through a composite literal or a call of `make`, and no explicit initialization is provided, the variable or value is given a default value. Each element of such a variable or value is set to the *zero value* for its type: `false` for booleans, `0` for numeric types, `""` for strings, and `nil` for pointers, functions, interfaces, slices, channels, and maps. This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified.

​	当通过声明或调用`new`为一个[变量](../Variables)分配存储空间时，或者通过复合字面量或调用`make`创建一个新的值时，如果没有提供明确的初始化，该变量或值将被赋予一个默认值。这种变量或值的每个元素都被设置为其类型的零值：布尔类型为`false`，数值型为`0`，字符串为`""`，指针、函数、接口、切片、通道和映射为`nil`。`这种初始化是递归进行的`，因此，举例来说，如果没有指定值，结构体数组的每个元素字段都将被设为零值。

These two simple declarations are equivalent:

​	这两个简单的声明是等同的：

```go 
var i int
var i int = 0
```

After

​	在

```go 
type T struct { i int; f float64; next *T }
t := new(T)
```

the following holds:

之后，以下情况成立：

```go 
t.i == 0
t.f == 0.0
t.next == nil
```

 The same would also be true after

​	同样的情况，在以下情况下也会成立

```go 
var t T
```

### Package initialization 包的初始化

Within a package, package-level variable initialization proceeds stepwise, with each step selecting the variable earliest in *declaration order* which has no dependencies on uninitialized variables.

​	在包内，包级别变量的初始化是逐步进行的，每一步都会选择声明顺序中最早的变量，该变量与未初始化的变量没有依赖关系。

More precisely, a package-level variable is considered *ready for initialization* if it is not yet initialized and either has no [initialization expression](https://go.dev/ref/spec#Variable_declarations) or its initialization expression has no *dependencies* on uninitialized variables. Initialization proceeds by repeatedly initializing the next package-level variable that is earliest in declaration order and ready for initialization, until there are no variables ready for initialization.

​	更确切地说，如果包级别变量还没有被初始化，并且没有初始化表达式，或者它的[初始化表达式](../DeclarationsAndScope#variable-declarations-变量声明)与未初始化的变量没有依赖关系，那么这个包级别变量就被认为可以被初始化。初始化的过程是重复初始化在声明顺序中最早的、准备好初始化的下一个包级别变量，直到没有准备好初始化的变量。

If any variables are still uninitialized when this process ends, those variables are part of one or more initialization cycles, and the program is not valid.

​	如果在这个过程结束时，仍有任何变量未被初始化，那么这些变量就是一个或多个初始化循环的一部分，程序是无效的。

Multiple variables on the left-hand side of a variable declaration initialized by single (multi-valued) expression on the right-hand side are initialized together: If any of the variables on the left-hand side is initialized, all those variables are initialized in the same step.

​	变量声明左侧的多个变量（由右侧的单个（多值）表达式来初始化）是一起被初始化的：如果左侧的任何一个变量被初始化，那么所有这些变量都在同一步骤中被初始化。

```go 
var x = a
var a, b = f() // a and b are initialized together, before x is initialized => a 和 b 是在 x 被初始化之前一起被初始化的
```

For the purpose of package initialization, [blank](https://go.dev/ref/spec#Blank_identifier) variables are treated like any other variables in declarations.

​	为了实现包的初始化，[空白](../DeclarationsAndScope#blank-identifier-空白标识符)变量与声明中的其他变量一样被处理。

The declaration order of variables declared in multiple files is determined by the order in which the files are presented to the compiler: Variables declared in the first file are declared before any of the variables declared in the second file, and so on. To ensure reproducible initialization behavior, build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler.

​	在多个文件中声明的变量的声明顺序由将文件呈现给编译器的顺序决定：在第一个文件中声明的变量在第二个文件中声明的任何变量之前声明，依此类推。为了确保可再现的初始化行为，鼓励构建系统按词法文件名顺序向编译器呈现属于同一包的多个文件。

Dependency analysis does not rely on the actual values of the variables, only on lexical *references* to them in the source, analyzed transitively. For instance, if a variable `x`'s initialization expression refers to a function whose body refers to variable `y` then `x` depends on `y`. Specifically:

​	依赖项分析不依赖于变量的实际值，只依赖于源文件中对它们的词法引用，并进行传递性分析。例如，如果一个变量`x`的初始化表达式引用了一个函数，该函数的主体引用了变量`y`，那么`x`就依赖于`y`：

- A reference to a variable or function is an identifier denoting that variable or function.
- 对变量或函数的引用是表示该变量或函数的标识符。
- A reference to a method `m` is a [method value](https://go.dev/ref/spec#Method_values) or [method expression](https://go.dev/ref/spec#Method_expressions) of the form `t.m`, where the (static) type of `t` is not an interface type, and the method `m` is in the [method set](https://go.dev/ref/spec#Method_sets) of `t`. It is immaterial whether the resulting function value `t.m` is invoked.
- 对一个方法`m`的引用是一个[方法值](../Expressions#method-values-方法值)或者形式为`t.m`的[方法表达式](../Expressions#method-expressions-方法表达式)，其中`t`的（静态）类型不是一个接口类型，并且方法`m`在`t`的[方法集](../PropertiesOfTypesAndValues#method-sets-方法集)中。是否调用结果函数值 `t.m` 并不重要。
- A variable, function, or method `x` depends on a variable `y` if `x`'s initialization expression or body (for functions and methods) contains a reference to `y` or to a function or method that depends on `y`.
- 如果`x`的初始化表达式或主体（对于函数和方法）包含对`y`的引用，或者对依赖于`y`的函数或方法的引用，那么变量、函数或方法`x`就依赖于变量`y`。

For example, given the declarations

​	例如，给定的声明有

```go 
var (
	a = c + b  // == 9
	b = f()    // == 4
	c = f()    // == 5
	d = 3      // == 5 after initialization has finished
)

func f() int {
	d++
	return d
}
```

the initialization order is `d`, `b`, `c`, `a`. Note that the order of subexpressions in initialization expressions is irrelevant: `a = c + b` and `a = b + c` result in the same initialization order in this example.

初始化顺序是 `d`，`b`，`c`，`a`。注意初始化表达式中子表达式的顺序是不相关的：在这个例子中`a = c + b`和`a = b + c`的初始化顺序是一样的。

Dependency analysis is performed per package; only references referring to variables, functions, and (non-interface) methods declared in the current package are considered. If other, hidden, data dependencies exists between variables, the initialization order between those variables is unspecified.

​	依赖项分析是按包进行的；只考虑引用当前包中声明的变量、函数和（非接口）方法。如果变量之间存在其他隐藏的数据依赖关系，那么这些变量之间的初始化顺序是未指定的。

For instance, given the declarations

​	例如，给定的声明有

```go 
var x = I(T{}).ab()   // x has an undetected, hidden dependency on a and b => x 存在 在 a 和 b 上的未被发现的隐藏依赖
var _ = sideEffect()  // unrelated to x, a, or b => 与 x, a, 或 b 无关
var a = b
var b = 42

type I interface      { ab() []int }
type T struct{}
func (T) ab() []int   { return []int{a, b} }
```

the variable `a` will be initialized after `b` but whether `x` is initialized before `b`, between `b` and `a`, or after `a`, and thus also the moment at which `sideEffect()` is called (before or after `x` is initialized) is not specified.

变量`a`将在`b`之后被初始化，但`x`是在`b`之前、`b`和`a`之间还是在`a`之后被初始化，以及`sideEffect()`被调用的时刻（在`x`被初始化之前或之后）都是未指定的。

Variables may also be initialized using functions named `init` declared in the package block, with no arguments and no result parameters.

​	变量也可以使用在包块中声明的（不带实参，也没有结果参数）的名为`init`的函数来初始化。

```go 
func init() { … }
```

Multiple such functions may be defined per package, even within a single source file. In the package block, the `init` identifier can be used only to declare `init` functions, yet the identifier itself is not [declared](https://go.dev/ref/spec#Declarations_and_scope). Thus `init` functions cannot be referred to from anywhere in a program.

​	每个包`可以定义多个这样的（init）函数`，`甚至在一个源文件中也可以（有多个init函数）`。在包块中，`init`标识符只能用于声明`init`函数，但标识符本身并没有被[声明](../DeclarationsAndScope)。因此，`init`函数不能在程序中的任何地方被引用。

The entire package is initialized by assigning initial values to all its package-level variables followed by calling all `init` functions in the order they appear in the source, possibly in multiple files, as presented to the compiler.

​	整个包通过为其所有包级变量分配初始值，然后按其在源代码中出现的顺序（可能在多个文件中）调用所有 `init` 函数来初始化，具体取决于编译器。

### Program initialization 程序初始化

The packages of a complete program are initialized stepwise, one package at a time. If a package has imports, the imported packages are initialized before initializing the package itself. If multiple packages import a package, the imported package will be initialized only once. The importing of packages, by construction, guarantees that there can be no cyclic initialization dependencies. More precisely:

​	完整程序的包分步初始化，一次一个包。如果包有导入，则在初始化包本身之前初始化导入的包。如果多个包导入一个包，则导入的包只初始化一次。根据结构，导入包可确保不存在循环初始化依赖项。更准确地说：

Given the list of all packages, sorted by import path, in each step the first uninitialized package in the list for which all imported packages (if any) are already initialized is [initialized](https://go.dev/ref/spec#Package_initialization). This step is repeated until all packages are initialized.

​	给定按导入路径排序的所有包的列表，在每一步中，列表中第一个未初始化的包（如果存在）的所有导入包（如果有）已经初始化，则对其进行初始化。此步骤重复执行，直到所有包都初始化。

Package initialization—variable initialization and the invocation of `init` functions—happens in a single goroutine, sequentially, one package at a time. An `init` function may launch other goroutines, which can run concurrently with the initialization code. However, initialization always sequences the `init` functions: it will not invoke the next one until the previous one has returned.

​	包初始化（变量初始化和 `init` 函数的调用）在一个 goroutine 中按顺序一次一个包地发生。 `init` 函数可以启动其他 goroutine，这些 goroutine 可以与初始化代码并发运行。但是，初始化始终对 `init` 函数进行排序：它不会调用下一个函数，直到前一个函数返回。

### Program execution 程序执行

A complete program is created by linking a single, unimported package called the *main package* with all the packages it imports, transitively. The main package must have package name `main` and declare a function `main` that takes no arguments and returns no value.

​	一个完整的程序是通过将一个名为（单个非导入包）`main`包与它所导入的所有包链接起来而创建的。主包必须有包名`main`，并声明一个不接受实参且不返回值的`main`函数。

```go 
func main() { … }
```

Program execution begins by [initializing the program](https://go.dev/ref/spec#Program_initialization) and then invoking the function `main` in package `main`. When that function invocation returns, the program exits. It does not wait for other (non-`main`) goroutines to complete.

​	程序的执行从初始化`main`包开始，然后调用函数`main`。当该函数调用返回时，程序退出。它`不会等待`其他（非`main`）goroutine完成。