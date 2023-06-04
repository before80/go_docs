+++
title = "教程：开始使用泛型"
weight = 11
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Tutorial: Getting started with generics - 教程：开始使用泛型

> 原文：[https://go.dev/doc/tutorial/generics](https://go.dev/doc/tutorial/generics)

​	本教程介绍Go中泛型的基本知识。通过泛型，您可以声明和使用函数或类型，这些函数或类型是为了使用调用代码提供的任何类型集而编写的。

​	在本教程中，您将声明两个简单的非泛型函数，然后在一个泛型函数中捕获相同的逻辑。

​	你将通过以下几个部分取得进展：

1. 为你的代码创建一个文件夹。
2. 添加非泛型函数。
3. 添加一个泛型函数来处理多种类型。
4. 在调用泛型函数时删除类型参数。
5. 声明一个类型约束。

注意：关于其他教程，请看[Tutorials](../Tutorials)。

> 注意：如果你愿意，你可以使用 [the Go playground in “Go dev branch” mode](https://go.dev/play/?v=gotip)来代替编辑和运行你的程序。

## 前提条件

- 安装 Go 1.18 或更高版本。有关安装说明，请参阅 [Installing Go](../InstallingGo)。
- 编辑代码的工具。任何文本编辑器都可以使用。
- 命令终端。在 Linux 和 Mac 上使用任何终端，以及在 Windows 上使用 `PowerShell` 或 `cmd`，Go 都能很好地工作。

## 为你的代码创建一个文件夹

首先，为你要写的代码创建一个文件夹。

a. 打开一个命令提示符，切换到你的主目录。

在Linux或Mac上：

```shell
$ cd
```

在Windows上：

```shell
C:\> cd %HOMEPATH%
```

​	本教程的其余部分将显示一个`$`作为提示符。你使用的命令在Windows上也会起作用。

b. 在命令提示符下，为你的代码创建一个名为`generics`的目录。

```shell
$ mkdir generics
$ cd generics
```

c. 创建一个模块来存放你的代码。

​	运行 `go mod init` 命令，给它你的新代码的模块路径。

```shell
$ go mod init example/generics
go: creating new go.mod: module example/generics
```

注意：对于生产代码，你可以根据自己的需要指定一个更具体的模块路径。更多信息，请参见[管理依赖关系](../../UsingAndUnderstandingGo/ManagingDependencies#naming-a-module)。

接下来，你将添加一些简单的代码来处理映射（maps）。

## 添加非泛型函数

​	在这一步中，你将添加两个函数，它们分别将一个映射的值相加并返回总数。

​	你要声明两个函数而不是一个，因为你要处理两种不同类型的映射：一个是存储`int64`值的，另一个是存储`float64`值的。

#### 编写代码

a. 使用你的文本编辑器，在 `generics` 目录中创建一个名为 `main.go` 的文件。你将在这个文件中写下你的Go代码。

b. 在`main.go`中，在文件的顶部粘贴以下包声明。

```go
package main
```

​	独立程序（相对于一个库）总是在`main` 包中。

c. 在包声明的下面，粘贴以下两个函数声明。

```go linenums="1"
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}
```

在这段代码中，你：

- 声明两个函数，将一个映射的值相加并返回总和。
  - `SumFloats`接收一个字符串到`float64`值的映射。
  - `SumInts`接收一个从字符串到`int64`值的映射。

d. 在`main.go`的顶部，在包声明的下面，粘贴以下`main`函数，以初始化两个映射，并在调用你在上一步声明的函数时将它们作为参数。

```go linenums="1"
func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))
}
```

在这段代码中，你：

- 初始化一个`float64`值的映射和一个`int64`值的映射，每个映射都有两个条目。

- 调用你之前声明的两个函数，以找到每个映射的值的总和。

- 打印结果。

  

e. 在`main.go`的顶部，就在包声明的下面，导入你需要的包来支持你刚刚写的代码。

第一行代码应该是这样的：

```go
package main

import "fmt"
```

f. 保存`main.go`

#### 运行代码

在包含`main.go`的目录下的命令行中，运行该代码。

```shell
$ go run .
Non-Generic Sums: 46 and 62.97
```

​	有了泛型，你可以在这里写一个函数而不是两个。接下来，你将为包含整数或浮点数的映射添加一个泛型函数。

## 添加一个泛型函数来处理多种类型

​	在这一节中，你将添加一个简单的泛型函数，它可以接收包含整数或浮点值的映射，有效地用一个简单的函数取代你刚才写的两个函数。

​	为了支持两种类型的值，这个简单的函数将需要一种方法来声明它支持哪些类型。另一方面，调用代码将需要一种方法来指定它是用整数映射还是浮点数映射来调用。

​	为了支持这一点，您将编写一个函数，除了它的普通函数参数之外，还声明`类型参数`。这些`类型参数`使函数具有泛型，使其能够处理不同类型的参数。你将用`类型参数`和普通函数参数来调用该函数。

​	每个`类型参数`都有一个`类型约束`，作为类型参数的一种`元类型（meta-type）`。每个类型约束都指定了调用代码可以为各自的类型参数使用的允许的类型参数。

​	虽然一个`类型参数的约束`通常代表一组类型，但在编译时，类型参数代表一个单一的类型 —— 调用代码中作为类型参数的类型。如果类型参数的类型不被类型参数的约束所允许，代码将无法编译。

​	请记住，类型参数必须支持泛型代码对其进行的所有操作。例如，如果函数代码试图对一个类型参数进行字符串操作（如索引），而这个类型参数的约束条件包括数字类型，那么代码将无法编译。

​	在你要写的代码中，你将使用一个允许整数或浮点数类型的约束条件。

#### 编写代码

a. 在你之前添加的两个函数下面，粘贴以下泛型函数。

```go linenums="1"
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

在这段代码中，你：

- 声明一个`SumIntsOrFloats`函数，有两个类型参数（在方括号内），`K`和`V`，以及一个使用类型参数的参数`m`，类型为`map[K]V`。该函数返回一个类型为`V`的值。
- 为`K`类型参数指定可比较的（`comparable`）类型约束。可比约束是专门为类似这样的情况而设计的，在Go中`预先声明了可比约束`。它允许任何类型的值可以作为比较运算符`==`和`!=`的操作数。 Go要求`映射键`是可比较的。因此，将`K`声明为可比较（`comparable`）是必需的，这样你就可以将`K`作为map变量的键。它还可以保证调用代码使用允许的类型作为映射键。
- 为`V`类型参数指定一个约束，该约束是两种类型的联合：`int64`和`float64`。使用`|`指定这两种类型的联合，意味着这个约束允许任何一种类型。编译器将允许这两种类型作为调用代码中的参数。
-  为`m`参数指定`map[K]V`类型，其中`K`和`V`是已经为类型参数指定的类型。注意，我们知道`map[K]V`是一个有效的map类型，因为`K`是一个可比较的类型。如果我们没有声明`K`的可比性（`comparable`），编译器会拒绝对`map[K]V`的引用。

b. 在`main.go`中，在已有的代码下面，粘贴以下代码。

```go hl_lines="2 3" linenums="1"
fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats[string, int64](ints),
    SumIntsOrFloats[string, float64](floats))
```

在这段代码中，你：

- 调用你刚才声明的泛型函数，传递你创建的每个映射。

- 指定类型参数 —— 方括号中的类型名称 —— 以明确在你调用的函数中应该取代类型参数的类型。

  ​	正如你将在下一节看到的，你通常可以在函数调用中`省略类型参数`。Go通常可以从你的代码中推断出它们。

- 打印由函数返回的和。

#### 运行代码

从包含`main.go`的目录中的命令行，运行代码。

```shell
$ go run .
Non-Generic Sums: 46 and 62.97
Generic Sums: 46 and 62.97
```

​	为了运行你的代码，在每次调用中，编译器都用该调用中指定的具体类型替换类型参数。

​	在调用你写的泛型函数时，你指定了类型参数，告诉编译器使用什么类型来代替函数的类型参数。正如你将在下一节看到的，在许多情况下，你可以`省略这些类型参数`，因为`编译器可以推断出`它们。

## 调用泛型函数时移除类型参数

​	在这一节中，你将添加一个修改版的泛型函数调用，做一个小改动以简化调用代码。你将删除类型参数，在这种情况下不需要这些参数。

​	当Go编译器可以推断出你想要使用的类型时，你可以在调用代码中`省略类型参数`。编译器会从函数参数的类型中推断出类型参数。

!!! warning "注意"

	请注意，这并不总是万能的。例如，如果你需要调用一个没有参数的泛型函数，你需要在函数调用中包含类型参数。

#### 编写代码

- 在`main.go`中，在你已有的代码下面，粘贴以下代码。

  ```go hl_lines="2 3" linenums="1"
  fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
      SumIntsOrFloats(ints),
      SumIntsOrFloats(floats))
  ```

  在这段代码中，你：

  - 调用泛型用函数，省略类型参数。

#### 运行该代码

在包含main.go的目录下的命令行中，运行该代码。

```shell
$ go run .
Non-Generic Sums: 46 and 62.97
Generic Sums: 46 and 62.97
Generic Sums, type parameters inferred: 46 and 62.97
```

​	接下来，你将进一步简化这个函数，将整数和浮点数的结合捕捉到一个可以重用的类型约束中，比如从其他代码中。

## 声明一个类型约束

​	在最后一节中，你将把之前定义的约束移到自己的接口中，这样你就可以在多个地方重复使用它。以这种方式声明约束有助于简化代码，例如当一个约束比较复杂时。

​	你把一个类型约束声明为一个接口。该约束允许任何类型实现该接口。例如，如果你声明一个具有三种方法的类型约束接口，然后在一个泛型函数中用一个类型参数来使用它，用于调用该函数的类型参数必须具有所有这些方法。

​	约束接口也可以指代特定的类型，正如你将在本节看到的那样。

#### 编写代码

a. 就在`main`上面，紧接着`import`语句，粘贴下面的代码来声明一个类型约束。

```go linenums="1"
type Number interface {
    int64 | float64
}
```

在这段代码中，你：

- 声明`Number`接口类型，作为类型约束使用。

- 在接口内声明一个`int64`和`float64`的联合。

  ​	实际上，你正在将联合从函数声明中移到一个新的类型约束中。这样，当你想把一个类型参数限制在`int64`或`float64`时，你可以使用这个`Number`类型约束，而不是写出`int64 | float64`。

b. 在你已经有的函数下面，粘贴下面的泛型 `SumNumbers`函数。

```go linenums="1"
// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

在这段代码中，你：

- 声明一个泛型函数，其逻辑与你之前声明的泛型函数相同，但用新的接口类型而不是联盟作为类型约束。和以前一样，你用类型参数来表示参数和返回类型。

c. 在`main.go`中，在你已经有的代码下面，粘贴以下代码。

```go hl_lines="2 3"
fmt.Printf("Generic Sums with Constraint: %v and %v\n",
    SumNumbers(ints),
    SumNumbers(floats))
```

在这段代码中，你：

- 对每个映射调用 SumNumbers，从每个映射的值中打印和。

  ​	与上一节一样，你在调用通用函数时`省略了类型参数`（方括号中的类型名称）。Go编译器可以从其他参数中推断出类型参数。

#### 运行代码

在包含`main.go`的目录下的命令行中，运行该代码。

```shell
$ go run .
Non-Generic Sums: 46 and 62.97
Generic Sums: 46 and 62.97
Generic Sums, type parameters inferred: 46 and 62.97
Generic Sums with Constraint: 46 and 62.97
```

## 总结

做得很好! 你刚刚向自己介绍了Go中的泛型。

建议的下一个主题：

- The [go Tour](../../GoTour)是对Go基础知识的一个很好的逐步介绍。
- 你会在[Effective Go](../../UsingAndUnderstandingGo/EffectiveGo)和[How to write Go code](../HowToWriteGoCode)中找到有用的Go最佳实践。
  

## 完整的代码

​	你可以在[Go playground](https://go.dev/play/p/apNmfVwogK0?v=gotip)上运行这个程序。在 playground 上只需点击运行按钮。

```go title="main.go" linenums="1" hl_lines="35 36"
package main

import "fmt"

type Number interface {
    int64 | float64
}

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first": 34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first": 35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))

    fmt.Printf("Generic Sums: %v and %v\n",
        SumIntsOrFloats[string, int64](ints),
        SumIntsOrFloats[string, float64](floats))

    fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
        SumIntsOrFloats(ints),
        SumIntsOrFloats(floats))

    fmt.Printf("Generic Sums with Constraint: %v and %v\n",
        SumNumbers(ints),
        SumNumbers(floats))
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```