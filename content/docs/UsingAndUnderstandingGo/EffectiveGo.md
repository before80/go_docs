+++
title = "Effective Go"
linkTitle = "Effective Go"
weight = 20
date = 2023-05-17T15:03:14+08:00
description = ""
isCJKLanguage = true
draft = false

+++
# Effective Go

> 原文：[https://go.dev/doc/effective_go](https://go.dev/doc/effective_go)

## 简介

​	Go是一种新的语言。虽然它借鉴了现有语言的思想，但它具有不寻常的特性，使得有效的Go程序在性质上不同于其亲属编写的程序。将C++或Java程序直接翻译成Go不太可能产生令人满意的结果——Java 程序是用 Java 而不是 Go 编写的。另一方面，从Go的角度来考虑问题可能会产生一个成功的但完全不同的程序。换句话说，要写好Go，了解它的属性和习语（idioms）是很重要的。了解 Go 编程的既定约定也很重要，例如命名、格式化、程序构造等等，这样您写的程序才能让其他 Go 程序员容易理解。

​	本文档给出了编写清晰、惯用的 Go 代码的技巧。它是对[语言规范]({{< ref "/langSpec/Introduction">}})、[Tour of Go](../../GoTour/UsingTheTour/welcome)和[How to Write Go Code](../../GettingStarted/HowToWriteGoCode)的补充，您应该首先阅读这些内容。

​	`2022年1月添加的注释`：**本文档是为2009年发布的Go编写的，此后没有进行过重大更新**。虽然它是了解如何使用语言本身的好指南，但由于语言的稳定性，它**几乎没有提到库**，也**没有提到Go生态系统自编写以来的重大变化**，如构建系统、测试、模块和多态性。**目前还没有计划对其进行更新**，因为已经发生了太多的事情，而且有一大批不断增加的文档、博客和书籍对现代Go的使用做了很好的描述。**Effective Go仍然是有用的，但读者应该明白它远不是一个完整的指南**。请参阅[issue 28782](https://github.com/golang/go/issues/28782)，了解相关情况。

### 示例

​	[Go package sources](https://go.dev/src/)不仅是作为核心库，也是作为如何使用该语言的示例。此外，许多包都包含了可以工作的、独立的可执行的示例，您可以直接从 [golang.org](https://golang.org/)网站上运行，比如这个[示例](https://go.dev/pkg/strings/#example_Map)（如果需要，点击 "Example"一词来打开它）。如果您对如何处理一个问题或如何实现某些东西有疑问，库中的文档、代码和示例可以提供答案、想法和背景。

## 格式化

​	格式化问题是最有争议但最不重要的问题。人们可以适应不同的格式风格，但最好是不必这样做，如果每个人都遵循相同的风格，就会花费更少的时间来讨论此问题。问题是如何在没有长篇指导手册的情况下实现这个理想。

​	在Go中，我们采取了一种不同寻常的方法，让机器来处理大多数的格式化问题。`gofmt`程序（也可以用`go fmt`，它在包级别而不是源文件级别操作）读取Go程序，并以缩进和垂直对齐的标准风格发出（emit）源代码，保留并在必要时重新格式化注释。如果您想知道如何处理一些新的布局情况，请运行`gofmt`；如果答案似乎不对，请重新排列您的程序（或提交关于`gofmt`的bug），而不是绕过它。

​	举例来说，没有必要花时间排列结构体字段上的注释。`Gofmt`会帮您做到这一点。给出的声明：

``` go 
type T struct {
    name string // name of the object
    value int // its value
}
```

`gofmt`会对齐列：

``` go 
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

​	标准包中的所有Go代码都已经用`gofmt`格式化过了。

​	一些格式化的细节仍然保留。非常简单：

- 缩进（Indentation）

  我们使用制表符来缩进，`gofmt`默认会发出（emit）制表符。只有在必须时才使用空格。

- 行的长度（Line length）

  Go没有行长限制。不要担心会超出打孔卡。如果一行太长，请换行并使用额外的制表符进行缩进。

- 圆括号（Parentheses）

  Go需要的圆括号比C和Java少：控制结构（if、for、switch）的语法中没有圆括号。另外，运算符的优先级层次更短，更清晰，因此

  ```go
  x<<8 + y<<16
  ```

  意思就像空格所示，不像其他语言那样。

## 注释

​	Go 提供了 C风格 `/* */` 块注释 和 C++风格 `//` 行注释。行注释是常规（norm）注释；块注释主要作为包注释出现，但在表达式内部或禁用大量代码时也很有用。

​	如果在顶级声明之前出现没有换行符的注释，则认为该注释记录了该声明本身。这些"文档注释（doc comments） "是给定 Go 包或命令的主要文档。有关文档注释的更多信息，请参见[Go Doc Comments](../../GoDocComments)。

## 命名

​	在 Go 中，命名与任何其他语言一样重要。它们甚至具有语义效果：名称在包外部的可见性取决于其第一个字符是否大写。因此，值得花费一点时间谈论 Go 程序中的命名约定。

### 包名

​	当一个包被导入后，包名就成为内容的访问器。在

​	当导入一个包时，包名称变成该内容的访问器。在执行：

``` go 
import "bytes"
```

之后，导入的包可以使用 `byte.Buffer`。如果每个使用该包的人都可以使用相同的名称来引用其内容，这将很有帮助，这意味着包名称应该很好：简短，简洁，意义明确。**按照惯例，包名称应为小写的单词名称；不需要使用下划线或混合大小写。** 应该更注重简洁，因为使用您的包的每个人都将输入该名称。并且不必担心引用次序冲突。包名称仅是导入的默认名称；它不需要在所有源代码中唯一，并且在冲突的罕见情况下，导入包可以选择不同的名称在本地使用。无论如何，混淆是罕见的，因为导入中的文件名确定正在使用的包。

​	另一个约定是包名是其源代码目录的基本名称。`src/encoding/base64`中的包以 `"encoding/base64"`被导入，其包名是`base64`，而不是`encoding_base64`，也不是`encodingBase64`。

​	另一个惯例是包名称是其源目录的基本名称；在 `src/encoding/base64` 中导入的包被称为 `"encoding/base64"`，但名称是 `base64`，而不是 `encoding_base64` 或 `encodingBase64`。

​	包的导入方将使用该名称来引用其内容，因此，包中导出的名称可以利用这一点避免重复。 （不要使用 `import .` 表示法，它可以简化必须在它们正在测试的包之外运行的测试，但在其他情况下应避免使用。）例如，在 `bufio` 包中称为 `Reader` 的缓冲读取器类型，而不是 `BufReader`，因为用户将其视为 `bufio.Reader`，这是一个清晰、简洁的名称。此外，由于导入的实体始终带有其包名称，所以 `bufio.Reader` 不会与 `io.Reader` 冲突。同样，用于创建 `ring.Ring` 新实例的函数——这是 Go 中构造函数的定义——通常被称为 `NewRing`，但由于`Ring`是包中唯一导出的类型，并且包叫做`ring`，所以它只被称为`New`，客户端将其视为`ring.New`。使用包结构来帮助您选择良好的名称。

​	另一个简短的例子是`once.Do`；`once.Do(setup)`读起来很好，不需要写成`once.DoOrWaitUntilDone(setup)`。长名称并不能自动使事情更易读。一个有用的文档注释往往比一个额外的长名称更有价值。

### Getters

​	Go不提供getter和setter的自动支持。自己提供getter和setter没有问题，而且通常是适当的，但在getter的名称中添加`Get`既不符合惯例，也不必要。如果您有一个名为`owner`（小写，未公开）的字段，getter方法应该被称为`Owner`（大写，公开），而不是`GetOwner`。大写名称的使用为导出提供了钩子，以区分字段和方法。如果需要setter函数，它很可能被称为`SetOwner`。在实践中，这两个名称都很好读：

``` go 
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

### 接口名称

​	按照惯例，单一方法的接口通过方法名称加上`-er`后缀或类似的修改来构建一个代理名词：`Reader`、`Writer`、`Formatter`、`CloseNotifier`等等。

​	有许多这样的名称，遵循它们和它们捕获的函数名称是很有成效的。`Read`、`Write`、`Close`、`Flush`、`String`等等具有规范的签名和含义。为避免混淆，除非具有相同的签名和含义，否则不要给您的方法赋予这些名称之一。相反，如果您的类型实现了一个与已知类型上的方法相同含义的方法，请使用相同的名称和签名，将您的字符串转换方法称为`String`而不是`ToString`。

### 驼峰式命名

​	最后，Go 的惯例是使用 `MixedCaps` 或 `mixedCaps` 而不是下划线来编写多单词名称。

## 分号

​	与 C 一样，Go 的正式语法使用分号来终止语句，但与 C 不同的是，这些分号不会出现在源代码中。相反，词法分析器在扫描源代码时会自动插入分号，因此输入文本大多没有分号。

​	规则如下。如果换行符之前的最后一个标记是标识符（包括像 `int` 和 `float64` 这样的单词）、基本字面量（如数字或字符串常量）或以下任一标记：

``` go 
break continue fallthrough return ++ -- ) }
```

​	词法分析器（lexer）总是在标记之后插入分号。可以总结为，"如果换行符出现在可能结束语句的标记之后，就插入分号"。

​	分号也可以省略在一个右花括号之前，因此像这样的语句

``` go 
go func() { for { dst <- <-src } }()
```

不需要分号。惯用的 Go 程序只在诸如 for 循环的条件中使用分号，用于分隔初始化器、条件和增减量元素。您若在一行中写多个语句，也需要用分号隔开。

​	分号插入规则的一个后果是，您不能将控制结构（`if`、`for`、`switch` 或 `select`）的左花括号放在下一行。如果您这样做，会在花括号之前插入一个分号，这可能会导致意外的效果。应该这样写：

``` go 
if i < f() {
    g()
}
```

而不是这样写：

``` go 
if i < f()  // wrong!
{           // wrong!
    g()
}
```

## 控制结构

​	Go的控制结构与C语言的控制结构有关，但在一些重要方面有所不同。没有`do`或`while`循环，只有略微通用的`for`；`switch`更加灵活；`if`和`switch`接受一个可选的初始化语句，就像`for`一样；`break`和`continue`语句接受一个可选的标签，以确定中断或继续的内容；还有一些新的控制结构，包括一个`类型选择（type switch）`和一个多路通信复用器（multiplexer），`select`。语法也略有不同：没有圆括号（parenthese），主体必须始终以花括号（brace）为界。

### If

​	在Go中，一个简单的`if`看起来像这样：

``` go 
if x > 0 {
    return y
}
```

强制的花括号鼓励将简单的`if`语句写在多行上。这样做是很好的风格，特别是当语句的主体包含控制语句，如`return`或`break`。

​	由于`if`和`switch`接受初始化语句，因此经常看到用它来设置局部变量。

``` go 
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

​	在Go库中，您会发现当`if`语句不流入下一条语句时——也就是说，主体以`break`、`continue`、`goto`或`return`结束——不必要的`else`被省略了。

``` go 
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)
```

​	这是一个常见情况的例子，代码必须防范一系列的错误条件。如果满足条件的控制流顺着页面运行，在出现错误时消除错误情况，那么代码读起来就很容易。由于错误情况往往在`return`语句中结束，因此之后的代码不需要`else`语句。

``` go 
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

### Redeclaration and reassignment 重新声明和重新赋值

​	提示：上一节中最后一个例子演示了`:=`短声明形式如何工作的一个细节。调用`os.Open`的声明是这样的

``` go 
f, err := os.Open(name)
```

这个语句声明了两个变量，`f`和`err`。几行之后，对`f.Stat`的调用是这样的

``` go 
d, err := f.Stat()
```

​	这看起来好像是声明了`d`和`err`。但是请注意，`err`出现在两个语句中。这种重复是合法的：`err`在第一条语句中声明，**但在第二条语句中只是重新赋值**。这意味着对`f.Stat`的调用使用了上面已经声明的`err`变量，只是给它一个新的值。

​	在`:=`声明中，即使变量`v`已经被声明，也可以出现，条件是：

- 本次声明与已声明的`v`在同一作用域内（如果`v`已经在外层作用域中声明过，则本次声明将创建一个新变量  :material-information:）。

- 本次初始化中的对应值是可分配给`v`的（即需要注意值的类型），并且
- 本次声明中至少有一个变量被创建（即本次声明的中新声明的）。

​	这个不寻常的属性是纯粹的实用主义，它使得在一个长的`if-else`链中使用一个单一的`err`值很容易。您会看到它经常被使用。

> ​	这里值得注意的是，在Go中，函数参数和返回值的作用域与函数主体相同，尽管它们在词法上出现在包围函数体的花括号之外。

### For

​	Go 的 `for` 循环与 C 的类似，但不一样。它统一了`for`和`while`，没有`do-while`。有三种形式，其中只有一种有分号。

``` go 
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```

​	短声明使我们很容易在循环中直接声明索引变量。

``` go 
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

​	如果您在一个数组、切片、字符串或映射上进行循环，或者从一个通道中读取，一个`range`子句可以管理循环。

``` go 
for key, value := range oldMap {
    newMap[key] = value
}
```

如果您只需要遍历中的第一个项（键或索引），去掉第二个就行：

``` go 
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```

​	如果您只需要遍历中的第二个项（值），使用空白标识符（即`_`），来丢弃第一个项：

``` go 
sum := 0
for _, value := range array {
    sum += value
}
```

​	空白标识符有很多用途，在[后面的章节](#the-blank-identifier)中会介绍。

​	对于字符串，`range`为您做了更多的工作，通过解析`UTF-8`来分解出各个Unicode码点。错误的编码将会占用一个字节并用符文（rune）`U+FFFD`来替换。(名称(带有相关的内建类型)`rune`，是Go对单个Unicode码点的称谓。详见[语言规范]({{< ref "/langSpec/LexicalElements#rune-literals-rune">}})）。循环

``` go 
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding => \x80 是一个非法的 UTF-8编码（字符）
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```

打印：

```
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```

​	最后，**Go没有逗号运算符**，**`++`和`--`是语句而不是表达式**。因此，如果您想在一个`for`中使用多个变量，您应该采用平行赋值（虽然它会拒绝`++`和`--`）。

``` go 
// Reverse a => 反转 a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

### Switch

​	Go的`switch`比C的`switch`更通用。其表达式无需为常量，甚至无需为整数，`case` 从上到下进行求值，直到找到匹配的`case`，如果`switch`没有表达式，则它将匹配`true`。因此，可以将`if`-`else`-`if`-`else`链写成一个`switch，而且这也更符合 Go 的风格。

``` go 
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

​	`switch` 并不会自动下溯，但`case`可以用逗号分隔的列表呈现。

``` go 
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

Although they are not nearly as common in Go as some other C-like languages, `break` statements can be used to terminate a `switch` early. Sometimes, though, it's necessary to break out of a surrounding loop, not the switch, and in Go that can be accomplished by putting a label on the loop and "breaking" to that label. This example shows both uses.

​	尽管它们在Go中并不像其他类 C语言那样常见，但`break`语句可以用来提前终止`switch`。有时候也必须打破层层的循环，而不仅仅是`switch`，在Go中可以通过在循环上加一个标签并 "breaking"该标签来实现。这个例子展示了这两种用法。

``` go 
Loop:
    for n := 0; n < len(src); n += size {
        switch {
        case src[n] < sizeOne:
            if validateOnly {
                break
            }
            size = 1
            update(src[n])

        case src[n] < sizeTwo:
            if n+1 >= len(src) {
                err = errShortInput
                break Loop
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]<<shift)
        }
    }
```

​	当然，`continue`语句也接受一个可选的标签，不过它只能在循环中使用。

​	作为本节的结束，下面是一个使用两个`switch`语句的字节片比较例程。

``` go 
// Compare returns an integer comparing the two byte slices,lexicographically.
// => 比较两个字节型切片，返回一个整数，按字典顺序。
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
// => 如果a == b，结果为0；如果a < b，结果为-1；如果a > b，结果为+1。
func Compare(a, b []byte) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
    switch {
    case len(a) > len(b):
        return 1
    case len(a) < len(b):
        return -1
    }
    return 0
}
```

### 类型选择 Type switch

​	`switch`也可以用来判断接口变量的动态类型。这样的`type switch`使用类型断言的语法，圆括号内有关键字`type`。如果`switch`在表达式中声明了一个变量，那么该变量将在每个子句中具有对应的类型。在每一个 `case` 子句中，重复利用该变量名字也是惯常的做法，实际上这是在每一个 `case` 子句中，分别声明一个拥有相同名字，但类型不同的新变量。

``` go 
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t) // %T 打印任何类型的 t 的 类型名
case bool:
    fmt.Printf("boolean %t\n", t) // t 是 bool 类型
case int:
    fmt.Printf("integer %d\n", t) // t 是 int 类型
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t 是 *bool 类型
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t 是 *int 类型
}
```

## Functions 函数

### Multiple return values 多重返回值

​	Go的一个不寻常的特点是，函数和方法可以返回多个值。这种形式可以用来改进C语言程序中的一些笨拙的习惯：带内错误返回（例如用 `-1` 表示 `EOF`），以及修改按地址传递的实参。

​	在C语言中，写入操作发生的错误会用一个负数标记，而错误代码被隐藏在一个易失性位置。在Go中，`Write`可以返回一个计数以及一个错误："是的，您写了一些字节，但不是全部，因为您填满了设备"。`os`包中的文件的`Write`方法的签名是：

``` go 
func (file *File) Write(b []byte) (n int, err error)
```

正如文档所述，它返回写入的字节数，当`n != len(b)`时返回一个非`nil`的`error`。这是一种常见的风格，更多的例子见错误处理一节。

​	我们可以采用一种简单的方法来避免为模拟引用参数而传入指针。下面是一个简单的函数，从一个字节切片的某个位置抓取一个数字，并返回该数值和下一个位置。

``` go 
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}
```

​	您可以用下面这样的（代码）来扫描一个输入切片`b`中的数字：

``` go 
    for i := 0; i < len(b); {
        x, i = nextInt(b, i)
        fmt.Println(x)
    }
```

### Named result parameters 命名的结果参数

​	Go函数的返回或结果 "参数 "可以被命名并作为常规变量使用，**就像传入参数一样**。**当命名后，它们在函数开始时被初始化为其类型的零值**；如果函数执行没有参数的`return`语句，则结果形参的当前值被用作返回值。

​	这些名称不是强制性的，但它们可以使代码更短、更清晰：它们就是文档。如果我们给 `nextInt` 的结果命名，那么哪个返回的是 `int` 就很明显了。

``` go 
func nextInt(b []byte, pos int) (value, nextPos int) {
```

​	由于被命名的结果已经初始化，且已经关联至无参数的返回，它们就能让代码简单而清晰。下面是`io.ReadFull`的一个版本，就是很好地使用了它们：

``` go 
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

### Defer

​	Go的`defer`语句预设了一个函数调用（即**推迟执行**函数），该函数会在执行 `defer` 的函数返回之前立即执行。这是一种不寻常但有效的方法，可以处理诸如资源必须被释放的情况，而不管一个函数采取哪种路径返回。典型的例子是解锁`mutex`和关闭文件。

``` go 
// Contents returns the file's contents as a string. => Contents将文件的内容作为一个字符串返回。
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later. => append稍后讨论。
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```

​	推迟对`Close`这样的函数的调用有两个好处。首先，它保证您永远不会忘记关闭文件，如果您以后又为该函数添加了新的返回路径时， 这种情况往往就会发生。其次，它意味着关闭位于打开附近，这比把它放在函数的最后要清楚得多。

​	被延迟函数的实参（如果函数是方法的话，还包括接收器）在**推迟**执行时就会求值，而不是在**调用**执行时。这样不仅无需担心变量值在函数执行时被改变， 同时还意味着单个已推迟的调用可推迟多个函数的执行。这里有一个简单的例子。

``` go 
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```

​	被推迟的函数是按`后进先出（LIFO）`的顺序执行，所以这段代码会导致函数返回时打印出`4 3 2 1 0`。一个更合理的例子是通过程序追踪函数执行的简单方法。我们可以这样写几个简单的追踪程序：

``` go 
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
    trace("a")
    defer untrace("a")
    // do something....
}
```

​	我们可以充分利用这个特点，即被推迟函数的实参在`defer`执行时就会求值。追踪例程可以为反追踪例程设置实参。这个例子：

``` go 
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```

打印

```
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

​	对于习惯了其他语言的块级资源管理的程序员来说，`defer`可能看起来很奇怪，但它最有趣和强大的应用正是来自于**它不是基于块而是基于函数**的事实。在关于`panic`和`recover`的章节中，我们将看到它的另一个可能性的例子。

## Data 数据

### Allocation with `new` 用new进行分配

​	Go有两个分配原语，即内置函数`new`和`make`。它们做不同的事情，适用于不同的类型，这可能会令人困惑，但规则很简单。我们先来谈谈`new`。这是一个分配内存的内置函数，但与其他一些语言中的同名函数不同，它并不初始化内存，只是将其置零。也就是说，`new(T)`为一个类型为`T`的新项分配了已置零的内存空间，并返回其地址，即一个类型为`*T`的值。用Go的术语来说，它**返回**一个指向新分配的`T`类型的零值的**指针**。

​	由于`new`返回的内存是已置零，所以在设计数据结构时，安排每种类型的零值无需进一步初始化就可以使用是很有帮助的。这意味着数据结构的用户可以用`new`创建一个数据结构并直接开始工作。例如，`bytes.Buffer`的文档指出，""零值的 `Buffer` 就是已准备就绪的缓冲区。" 同样，`sync.Mutex`也没有一个显式的构造函数或`Init`方法。相反，零值的 `sync.Mutex` 就已经被定义为已解锁的互斥锁了。

​	"零值属性" 可以带来各种好处。考虑一下这个类型声明。

``` go 
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
```

`SyncedBuffer`类型的值也是在分配或声明时就可以立即使用。后续代码中， `p` 和 `v` 无需进一步处理即可正确工作。

``` go 
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

### Constructors and composite literals 构造函数和复合字面量

​	有时零值还不够好，这时需要一个初始化构造函数，就像这个源自`os`包的例子。

``` go 
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

​	这里面有很多繁文缛节。**我们可以使用复合字面量来简化它**，复合字面是一个表达式，每次求值都会创建一个新的实例。

``` go 
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

注意，与C语言不同的是，`返回局部变量的地址是完全可以的`；与该变量相关的存储（数据）在函数返回后仍然存在。事实上，获取一个复合字面量的地址在每次求值时都会分配一个新的实例，所以我们可以将最后两行合并起来。

``` go 
    return &File{fd, name, nil, 0}
```

​	复合字面量的字段必须按顺序全部列出。但如果以 **字段**`:`**值** 对的形式明确地标出元素，初始化字段时就可以按任何顺序出现，未给出的字段值将赋予零值。因此我们可以用如下形式

``` go 
    return &File{fd: fd, name: name}
```

​	少数情况下，如果复合字面量不包含任何字段，它就会为该类型创建一个零值。`new(File)` 和 `&File{}` 的表达式是等价的。

​	复合字面量同样可用于创建数组、切片和映射，字段标签（field labels）可以是索引或映射的键值。在这些例子中，不管`Enone`、`Eio`和`Einval`的值是什么，只要它们的标签不同，初始化就会正常进行。

``` go 
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
```

### Allocation with `make `用make进行分配

​	回到内存分配上。内置函数`make(T, args)`的目的与`new(T)`不同。`它只用于创建切片、映射和通道`，并返回一个初始化（**非置零**）的`T`（而不是`*T`）类型的值。出现这种用差异的原因在于，**这三种类型本质上为引用数据类型，它们在使用前必须初始化**。例如，切片是一个包含指向（在一个数组内）数据的指针、长度和容量的三个项的描述符，在这些项被初始化之前，这个切片为`nil`。对于切片、映射和通道，`make`初始化内部数据结构，并准备好使用的值。例如，

``` go 
make([]int, 10, 100)
```

分配了一个具有100个`int`的数组，然后创建了一个长度为10，容量为 100 并指向该数组中前 10 个元素的切片结构。(当创建一个切片时，容量可以省略；更多信息请看切片部分)。相反，`new([]int)`返回一个指向新分配的、已置零的切片结构的指针，即一个指向`nil`切片值的指针。

​	这些例子说明了`new`和`make`的区别。

``` go 
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful => 分配切片结构；*p == nil；很少用到
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints => 切片 v 现在引用了一个具有 100 个 int 元素的新数组

// Unnecessarily complex: 没必要的复杂用法
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic: 惯用法
v := make([]int, 100)
```

记住，`make`只适用于映射、切片和通道，且不返回指针。要获得一个显式的指针，需要用`new`来分配或者显式地获取一个变量的地址。

### Arrays 数组

​	在规划内存的详细布局时，数组是很有用的，有时还能避免过多的内存分配，但它们主要用作切片的构件，这将是下一节的主题。为了给这一主题打下基础，这里有一些关于数组的话。

​	在Go和C中，数组的工作方式有很大的不同：

- 数组是值。将一个数组分配给另一个数组会复制所有的元素。
- 特别是，如果您把一个数组传递给一个函数，它将接收到一个数组的副本，而不是一个指针。
- 数组的大小是其类型的一部分。类型`[10]int`和`[20]int`是不同的。

​	数组为值属性可能很有用，但也代价高昂；如果您想获得类似C语言那样的行为和效率，您可以传递一个指向数组的指针。

``` go 
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // Note the explicit address-of operator => 注意这里显示的 &运算符
```

​	但这并不是 Go 的惯用法，切片才是。

### Slices 切片

​	切片通过对数组进行封装，为数据序列提供了一个更通用、更强大、更方便的接口。除了矩阵变换这类需要明确维度的情况外，Go 中的大部分数组编程都是通过切片来完成的。

Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array. A `Read` function can therefore accept a slice argument rather than a pointer and a count; the length within the slice sets an upper limit of how much data to read. Here is the signature of the `Read` method of the `File` type in package `os`:

​	切片持有对底层数组的引用，如果您把一个切片赋予另一个切片，两者都会引用同一个数组。如果某个函数接受一个切片参数，那么它对切片中的元素所做的改变对调用者来说是可见的，类似于传递一个指向底层数组的指针。因此，`Read`函数可以接受一个切片实参，而不是一个指针和一个计数；切片中的长度决定了可读取数据的上限。下面是`os`包中文件类型的读取方法的签名：

``` go 
func (f *File) Read(buf []byte) (n int, err error)
```

​	该方法返回读取的字节数和一个错误值（如果有的话）。要读入一个更大的缓冲区`buf`的前32个字节，可以对缓冲区进行切片（这里作为动词使用）。

``` go 
    n, err := f.Read(buf[0:32])
```

​	这样的切片是很常见的，而且很有效。若不谈效率，下面的切片也会读取缓冲区的前32个字节。

``` go 
    var n int
    var err error
    for i := 0; i < 32; i++ {
        nbytes, e := f.Read(buf[i:i+1])  // Read one byte.
        n += nbytes
        if nbytes == 0 || e != nil {
            err = e
            break
        }
    }
```

​	切片的长度可以改变，只要它仍然符合底层数组的限制；只需将它赋予其自身的切片即可。一个切片的容量，可以通过内置函数`cap`获得，它将给出该切片可能取得的最大长度。以下有一个将数据追加到切片的函数。如果数据超过了容量，切片将被重新分配。返回值即为所得的切片。该函数利用`len`和`cap`在应用于`nil` 切片时是合法的，它将返回0。

``` go 
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate => 重新分配
        // Allocate double what's needed, for future growth. => 为未来的增长,两倍分配所需的长度.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type. => copy 函数是预先声明的，适用于任何切片类型。
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
```

​	最终我们必须返回切片，因为尽管`Append`可以修改`slice`的元素，但`slice`本身（其运行时数据结构包含指针、长度和容量）是通过值传递的。

​	向切片中追加的想法非常有用，因此有专门的内置函数 `append`。为了理解这个函数的设计，我们需要更多的信息，我们将稍后再介绍它。

### Two-dimensional slices 二维切片

​	Go 的数组和切片都是一维的。要创建相当于二维数组或切片，必须定义一个数组的数组或切片的切片，像这样：

``` go 
type Transform [3][3]float64  // A 3x3 array, really an array of arrays.
type LinesOfText [][]byte     // A slice of byte slices.
```

​	由于切片是可变长度的，因此有可能让每个内部的切片都是不同的长度。这可能是一种常见的情况，就像我们的`LinesOfText`例子：每行都有其自己的长度：

``` go 
text := LinesOfText{
    []byte("Now is the time"),
    []byte("for all good gophers"),
    []byte("to bring some fun to the party."),
}
```

Sometimes it's necessary to allocate a 2D slice, a situation that can arise when processing scan lines of pixels, for instance. There are two ways to achieve this. One is to allocate each slice independently; the other is to allocate a single array and point the individual slices into it. Which to use depends on your application. If the slices might grow or shrink, they should be allocated independently to avoid overwriting the next line; if not, it can be more efficient to construct the object with a single allocation. For reference, here are sketches of the two methods. First, a line at a time:

​	有时，有必要分配一个二维切片，例如，在处理像素的扫描行时，这种情况就会发生。有两种方式可以实现这一点。一种是独立地分配每个切片；另一种是分配一个数组，将各个切片指向它。使用哪种方法取决于您的应用。如果切片可能会增长或缩小，则它们应该独立分配，以避免覆盖下一行；如果不会，用单次分配构建对象可能更有效率。以下是这两种方法的大概代码，仅供参考。首先是一次分配一行。

``` go 
// Allocate the top-level slice. => 分配底层切片
picture := make([][]uint8, YSize) // One row per unit of y. => 每 y 个单元一行
// Loop over the rows, allocating the slice for each row.
for i := range picture {
    picture[i] = make([]uint8, XSize)
}
```

现在是作为一个分配，对行进行切片：

``` go 
// Allocate the top-level slice, the same as before. => 分配底层切片， 和上面的一样
picture := make([][]uint8, YSize) // One row per unit of y. => 每 y 个单元一行
// Allocate one large slice to hold all the pixels. => 分配一个大一点的切片用来容纳所有的像素
pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8. => 指定类型[]uint8, 即使图片是 [][]uint8。
// Loop over the rows, slicing each row from the front of the remaining pixels slice. => 循环遍历行，从剩余像素切片的前面对每一行进行切片。
for i := range picture {
    picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```

### Maps 映射

​	映射是一种方便而强大的内置数据结构，它将一种类型的值（键）与另一种类型的值（元素或值）关联起来。其`键`可以是任何相等性操作符支持的类型，如`整数`、`浮点数`和`复数`、`字符串`、`指针`、`接口`（只要其动态类型支持相等性判断）、`结构体`和`数组`。切片不能被用作映射的键，因为它们的相等性还未定义。与切片一样，映射也持有对一个底层数据结构的引用。**若将映射传入函数中，并更改了该映射的内容，则此修改对调用者同样可见**。

​	映射可以使用一般的复合字面量语法和冒号分隔的键值对来构建，所以在初始化过程中很容易构建它们。	

``` go 
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
```

​	赋值和获取映射值在语法上看起来就像对数组和切片做同样的事情，只是索引无需是一个整数。

``` go 
offset := timeZone["EST"]
```

​	试图通过映射中不存在的键来获取值，就会返回与该映射中项的类型对应的零值。例如，如果映射包含整数，查找一个不存在的键将返回`0`。集合可以被实现为一个值类型为`bool`的映射。将该映射中的项置为 `true` 可将该值放入集合中，此后通过简单的索引操作即可判断是否存在。

``` go 
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map => 若 person不在映射中，则返回 false
    fmt.Println(person, "was at the meeting")
}
```

​	有时您需要区分某项是不存在还是其值为零值。是有一个 "`UTC` "的条目，还是因为它根本就不在映射中，所以是`0`？您可以用一种`多重赋值`的形式进行区分。

``` go  hl_lines="3 3"
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

​	由于明显的原因，这被称为 "逗号ok "惯用法。在这个例子中，如果`tz`存在，`seconds`将被适当地设置，`ok`将为`true`；如果不存在，`seconds`将被设置为零，`ok`将为`false`。下面是一个函数，它把它和一个很好的错误报告放在一起：

``` go 
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

​	若仅需判断映射中是否存在某项而不关心实际的值，您可以用[空白标识符](#the-blank-identifier)（`_`）来代替该值的一般变量。

``` go 
_, present := timeZone[tz]
```

​	要删除映射中的项，请使用`delete`内置函数，它以映射及要被删除的键为实参。即使该键已经不在映射中，此操作也是安全的。

``` go 
delete(timeZone, "PDT")  // Now on Standard Time
```

### Printing 打印

​	Go中的格式化打印使用类似于C的`printf`系列的风格，但更丰富，更通用。这些函数存在于`fmt`包中，且函数名首字母均为大写：`fmt.Printf`, `fmt.Fprintf`, `fmt.Sprintf`等等。字符串函数（`Sprintf`等）会返回一个字符串，而非填充给定的缓冲区。

​	您无需提供一个格式字符串。对于`Printf`、`Fprintf`和`Sprintf`中的每一个，都分别有对应另外的函数，例如`Print`和`Println`。这些函数不接受格式字符串，而是为每个实参生成一种默认格式。`Println`版本的函数还在实参之间插入一个空白，并在输出中附加一个换行符，而`Print`版本的函数仅在两边的操作数都不是字符串的情况下添加空白。在这个例子中，每一行都产生相同的输出。

``` go 
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
```

​	`fmt.Fprint` 一类的格式化打印函数可接受任何实现了 `io.Writer` 接口的对象作为第一个实参；变量 `os.Stdout` 和 `os.Stderr` 是人们熟悉的实例。

​	从这里开始，就与 C 有些不同了。首先，像 `%d` 这样的数值格式并不接受表示符号或大小的标记， 打印例程会根据实参的类型来决定这些属性。

``` go 
var x uint64 = 1<<64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
```

打印

```
18446744073709551615 ffffffffffffffff; -1 -1
```

​	如果您只想得到默认的转换，比如整数的十进制，您可以使用通用格式`%v`（代表 "值"）；其结果与 `Print`和`Println`的输出完全相同。此外，这种格式可以打印任何数值，甚至是数组、切片、结构体和映射。下面是上一节中定义的时区映射的打印语句。

``` go 
fmt.Printf("%v\n", timeZone)  // or just fmt.Println(timeZone)
```

即可得到输出：

```
map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
```

​	对于映射，`Printf`一类的函数会按照键的字典顺序排序输出。

​	在打印结构体时，修改后的格式`%+v`对结构体的字段进行注解，对于任何值，替代格式`%#v`以完整的Go语法打印出该值。

``` go 
type T struct {
    a int
    b float64
    c string
}
t := &T{ 7, -2.35, "abc\tdef" }
fmt.Printf("%v\n", t)
fmt.Printf("%+v\n", t)
fmt.Printf("%#v\n", t)
fmt.Printf("%#v\n", timeZone)
```

prints 打印

```
&{7 -2.35 abc   def}
&{a:7 b:-2.35 c:abc     def}
&main.T{a:7, b:-2.35, c:"abc\tdef"}
map[string]int{"CST":-21600, "EST":-18000, "MST":-25200, "PST":-28800, "UTC":0}
```

(请注意其中的 & 符号）当遇到`string`或`[]byte`类型的值时，可使用 `%q` 产生带引号的字符串。而格式 `%#q` 会尽可能使用反引号。(`%q`格式也可用于整数和符文，它会产生一个单引号符文常量)。另外，`%x`也适用于字符串、字节数组、字节切片以及整数，生成一个很长的十六进制字符串，并且在格式中加入空格（`% x`），它还会在字节之间插入空格。

​	另一种实用的格式是`%T`，它打印出某个值的类型。

``` go 
fmt.Printf("%T\n", timeZone)
```

打印

``` go 
map[string]int
```

​	如果您想控制自定义类型的默认格式，只需要在该类型上定义一个具有`String() string`签名的方法。对于我们的简单类型`T`，可能看起来像这样。

``` go 
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)
```

打印格式为

```
7/-2.35/"abc\tdef"
```

(如果您需要像指向 `T` 的指针那样打印类型 `T` 的**值**，那么`String`的接收器必须是值类型的；上面这个例子（中的接收器）使用一个指针，因为这对结构体类型来说更加有效和惯用。更多信息请参见下面关于[指针与值接收器](#pointers-vs-values)的部分）。

​	我们的`String`方法能够调用`Sprintf`，因为打印例程是完全可重入的，并可以用这种方式进行封装。不过，关于这种方法有一个重要的细节需要知道：**请勿通过调用`Sprintf`的方式来构造`String`方法，这样它会无限递归您的`String`方法**。如果`Sprintf`调用试图将接收器直接打印成字符串，而该字符串又将再次调用该方法，则会发生这种情况。这是一个常见且容易犯的错误，正如本例所示。

``` go 
type MyString string

func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", m) // Error: will recur forever. => 错误：会无限递归
}
```

​	这也很容易解决：将实参转换为基本字符串类型，该实参没有这个方法。

``` go 
type MyString string
func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion. =>  可以：注意转换
}
```

​	在[初始化一节](#initialization)，我们将看到另一种避免这种递归的技术。

​	另一种打印技术是将打印例程的实参直接传递给另一个这样的例程。`Printf`的签名为它的最后一个参数使用了`...interface{}`类型，这样格式的后面就能出现（任意类型之一的）任意数量的参数。

``` go 
func Printf(format string, v ...interface{}) (n int, err error) {
```

​	在函数`Printf`中，`v`的行为就像一个`[]interface{}`类型的变量，**但如果它被传递给另一个变参函数，它的行为就像一个普通的实参列表**。以下是我们之前用过的 `log.Println` 的实现。它直接将实参传递给 `fmt.Sprintln` 来进行实际格式化。

``` go  hl_lines="4 4"
// Println prints to the standard logger in the manner of fmt.Println.
// => Println 通过 fmt.Println 的方式将日志打印到标准记录器
func Println(v ...interface{}) {
    std.Output(2, fmt.Sprintln(v...))  // Output takes parameters (int, string) => Output 接收参数 (int, string)
}
```

​	我们在嵌套调用 `Sprintln` 的 `v` 后面写上 `...` 来告诉编译器把 `v` 当作一个实参列表；否则它就会把 `v` 作为一个单一的切片实参来传递。

​	还有很多关于打印知识点没有提及。详情请参见[fmt]({{< ref "/stdLib/fmt" >}})包的`godoc`文档。

​	顺便说一下，**`...`形参可指定具体的类型**，例如`...int`用于`min`函数，该函数选择整数列表中的最小值。

``` go  hl_lines="1"
func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // largest int
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}
```

### Append 追加

​	现在我们有了解释`append`内置函数的设计所需的缺失部分。`append`的签名与我们上面的自定义`Append`函数不同。大致来说，它是这样的：

``` go 
func append(slice []T, elements ...T) []T
```

其中`T`是任何给定类型的占位符。**实际上，您不能在Go中写一个由调用者决定类型`T`的函数**。**这就是为什么`append`是内置的：它需要编译器的支持**。

​	`append`所做的是将元素追加到切片的末尾并返回结果。结果需要被返回，原因与我们手写的`Append`一样，即底层数组可能会改变。这个简单的例子

``` go 
x := []int{1,2,3}
x = append(x, 4, 5, 6)
fmt.Println(x)
```

打印出`[1 2 3 4 5 6]`。所以`append`的工作方式有点像`Printf`，接收任意数量的实参。

​	但如果我们要像 `Append` 那样将一个切片追加到另一个切片中呢？很简单：在调用处使用`...`，就像我们在上面调用`Output`时那样。以下代码片段的输出与上一个相同。

``` go  hl_lines="3 3"
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)
fmt.Println(x)
```

​	如果没有那个`...`，它就会由于类型错误而无法编译；因为`y`不是int类型。

## Initialization 初始化

​	虽然从表面上看Go与C或C++的初始化没有什么不同，但Go的初始化功能更强大。在初始化过程中，不仅可以构建复杂的结构，还能正确处理不同包对象间的初始化顺序。

### Constants 常量

​	Go中的常量就是constant。它们在编译时被创建，即便它们可能是函数中定义的局部变量，常量只能是数字、字符（符文）、字符串或布尔值。由于编译时的限制，定义它们的表达式必须也是可被编译器求值的常量表达式。例如，`1<<3`是一个常量表达式，而`math.Sin(math.Pi/4)`则不是，因为对`math.Sin`的函数调用在运行时才会发生。

​	在Go中，枚举常量是使用`iota`枚举器创建的。由于`iota`可以是表达式的一部分，而且表达式可以隐式地重复，这样也就更容易构建复杂的值的集合了。

``` go 
type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier => 通过赋予空白标识符来忽略第一个值
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```

​	由于可将 `String` 之类的方法附加在用户定义的类型上， 因此它就为打印时自动格式化任意值提供了可能性。虽然您会看到它最常被应用于结构体，但这种技术对标量类型也很有用，如`ByteSize`等浮点类型。

``` go 
func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}
```

表达式`YB`打印为`1.00YB`，而`ByteSize(1e13)`打印为`9.09TB`。

​	这里使用`Sprintf`来实现`ByteSize`的`String`方法是安全的（不会无限递归），这倒不是因为类型转换，而是它以 `%f` 调用了 `Sprintf`，它并不是一个字符串格式。**`Sprintf`只有在需要字符串时才会调用`String`方法**，而`%f`需要的是一个浮点值。

### Variables 变量

​	变量可以像常量一样被初始化，而且可以初始化为一个可在运行时得出结果的普通表达式。

``` go 
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
```

### The init function - init 函数

​	最后，每个源文件都可以定义自己的无参数（niladic） `init`函数来设置任何需要的状态。(实际上每个文件可以有多个`init`函数。)而它的结束就意味着初始化结束： **只有该包中的所有变量声明都通过它们的初始化器求值后 `init` 才会被调用**， **而包中的变量只有在所有已导入的包都被初始化后才会被求值**。

​	除了那些不能被表示成声明的初始化外，`init` 函数还常被用在程序真正开始执行前，检验或校正程序的状态。

``` go 
func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line. => gopath 可通过命令行中的 --gopath 标记覆盖掉。
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}
```

## Methods 方法

### Pointers vs. Values 指针与值

​	正如 `ByteSize` 那样，我们可以为任何已命名的类型（除了指针或接口）定义方法； 接收器可不必为结构体。

​	在上面关于切片的讨论中，我们写了一个`Append`函数。我们可以把它定义为切片上的方法。要做到这一点，我们首先声明一个命名的类型，我们可以将该方法与之绑定，然后使该方法的接收器成为该类型的值。

``` go 
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // Body exactly the same as the Append function defined above. => 主体与上面定义的Append函数完全相同。
}
```

​	这仍然需要该方法返回更新后的切片。为了消除这种不便，我们可通过重新定义该方法， 将一个指向 `ByteSlice` 的指针作为该方法的接收器， 这样该方法就能重写调用者提供的切片了。

``` go 
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body as above, without the return.
    *p = slice
}
```

​	事实上，我们可以做得更好。如果我们修改我们的函数，使它看起来像一个标准的`Write`方法，像这样：

``` go 
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // Again as above.
    *p = slice
    return len(data), nil
}
```

then the type `*ByteSlice` satisfies the standard interface `io.Writer`, which is handy. For instance, we can print into one.

那么`*ByteSlice`类型就满足标准接口`io.Writer`，这会很实用。例如，我们可以通过打印将内容写入。

``` go 
var b ByteSlice
fmt.Fprintf(&b, "This hour has %d days\n", 7)
```

​	我们传递一个`ByteSlice`的地址，因为只有`*ByteSlice`才满足`io.Writer`。以指针或值为接收器的区别在于：值方法可通过指针和值调用， 而指针方法只能通过指针来调用。

​	之所以会有这条规则是因为指针方法可以修改接收器；通过值调用它们会导致方法接收到该值的副本，故任何修改都会被丢弃。因此，Go语言不允许这种错误。不过，有一个方便的例外。**当值是可寻址的，那么Go语言通过自动插入取地址操作符来处理在值上调用指针方法的常见情况**。在我们的例子中，变量`b`是可寻址的，所以我们可以只用`b.Write`来调用它的`Write`方法。编译器将为我们把它重写成`(&b).Write`。

​	顺便说一下，在字节切片上使用`Write`的想法是实现`bytes.Buffer`的核心。

## 接口和其他类型

### 接口

​	在Go中，接口提供了一种指定对象行为的方式：如果某个东西可以做这个，那么它就可以在这里使用。我们已经见过许多简单的示例了；通过实现 `String` 方法，我们可以自定义打印函数，而`Fprintf`可以将生成输出到任何具有`Write`方法的地方。在 Go 代码中， 仅包含一两种方法的接口很常见，且其名称通常来自于实现它的方法， 如 `io.Writer` 就是实现了 `Write` 的一类对象。

​	每种类型都能实现多个接口。例如一个实现了 `sort.Interface` 接口的集合就可通过 `sort` 包中的例程进行排序。该接口包括了 `Len()`、`Less(i, j int) bool` 以及 `Swap(i, j int)`，另外，该集合仍然可以有一个自定义的格式化器。 以下特意构建的例子 `Sequence` 就同时满足这两种情况。

``` go 
type Sequence []int

// sort.Interface所需的方法。
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

//  Copy方法返回Sequence的副本。
func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

// 用于打印的方法-在打印之前对元素进行排序。
func (s Sequence) String() string {
    s = s.Copy() // 创建副本；不要覆盖实参本身
    sort.Sort(s)
    str := "["
    for i, elem := range s { // 循环的时间复杂度是O(N²);将在下一个例子中修复它。
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```

### 转换

​	`Sequence` 的 `String` 方法重新实现了 `Sprint` 为切片实现的功能。（它还具有O(N²)的复杂度，这是很差的。）若我们在调用 `Sprint` 之前将 `Sequence` 转换为纯粹的 `[]int`，就能共享已实现的功能。

``` go 
func (s Sequence) String() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
```

​	该方法是通过类型转换技术，在 `String` 方法中安全调用 `Sprintf` 的另个一例子。若我们忽略类型名的话，这两种类型（`Sequence` 和 `[]int`）其实是相同的，因此在二者之间进行转换是合法的。 转换过程并不会创建新值，它只是暂时让现有的值看起来有个新类型而已。 （还有些合法转换则会创建新值，如从整数转换为浮点数等。）

​	在Go程序中，为访问不同的方法集而进行类型转换的情况非常常见。例如，我们可以使用现有的`sort.IntSlice`类型类型来简化整个示例：

``` go  hl_lines="6 6"
type Sequence []int

// 打印方法-在打印之前对元素进行排序
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
```

​	现在，不必让 `Sequence` 实现多个接口（排序和打印）， 我们可通过将数据项转换为多种类型（`Sequence`、`sort.IntSlice`和`[]int`）来使用相应的功能。这在实践中比较少见，但往往却很有效。

### 接口转换和类型断言

​	[类型选择](#type-switch)是一种类型转换形式：它们接受一种接口，在开关 （switch）中根据其判断选择对应的情况（case），并在某种意义上将其转换为该种类型。以下代码为`fmt.Printf`通过使用类型选择将一个值变成一个字符串的简化版本。若它已经为字符串，我们想要接口持有的实际字符串值，若它有 String 方法，我们想要调用该方法所得的结果。

``` go 
type Stringer interface {
    String() string
}

var value interface{} // value 由调用者提供
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

​	第一个`case`获取了一个具体的值；第二个`case`将该接口转换为另一个接口。这种方式对于混合类型来说非常完美。

​	如果我们只关心一种类型呢？如果我们知道这个值持有一个`string`类型，而我们只想提取它？ 只需一种情况的类型选择就行，但它需要**类型断言**。类型断言接受一个接口值并从中提取一个指定的显式类型的值。这种语法借鉴自类型选择开头的子句，但它需要一个显式的类型名， 而非 `type` 关键字：

``` go 
value.(typeName)
```

而其结果则是具有静态类型 `typeName` 的新值。该类型必须是该接口所持有的具体类型，或者是该值可以被转换为的第二种接口类型。为了提取我们知道在该值中的字符串，我们可以写：

``` go 
str := value.(string)
```

> 以下给出示例来解释下：该类型必须是该接口所持有的具体类型，或者是该值可以被转换为的第二种接口类型。
>
> ```go 
> type Animal interface {
>     MakeSound() string
> }
> 
> type Dog struct {}
> 
> func (d Dog) MakeSound() string {
>     return "Bark!"
> }
> 
> type Cat struct {}
> 
> func (c Cat) MakeSound() string {
>     return "Meow!"
> }
> 
> func MakeSpecificSound(a Animal) {
>     if d, ok := a.(Dog); ok {
>         fmt.Println(d.MakeSound())
>     } else if c, ok := a.(Cat); ok {
>         fmt.Println(c.MakeSound())
>     } else {
>         fmt.Println("Unknown animal type!")
>     }
> }
> 
> ```
>
> 

​	但是如果结果发现值不包含字符串，程序就会因运行时错误而崩溃。为了防止这种情况的发生，可以使用"comma, ok"惯用测试它能安全地判断该值是否为字符串：

``` go 
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

​	如果类型断言失败，`str`仍然存在，并且是字符串类型，但它将拥有零值，即空字符串。

​	作为（"comma, ok"）能力的说明，这里有一个`if-else`语句，相当于本节开头的类型选择。

``` go 
if str, ok := value.(string); ok {
    return str
} else if str, ok := value.(Stringer); ok {
    return str.String()
}
```

### 通用性

​	如果一个类型只是为了实现一个接口而存在，并且永远不会有除了这个接口以外的导出方法，那么就没必要导出这个类型本身。仅导出接口可以清晰地表明该值除了接口描述的行为之外没有其他有趣的行为。它还避免了在每个常见方法的实例上重复文档的必要。

​	在这种情况下，构造函数应该返回一个接口值而非实现的类型。举个例子，在 hash 库中，`crc32.NewIEEE`和`adler32.New`都返回接口类型`hash.Hash32`。在Go程序中用`CRC-32算法`代替`Adler-32`，只需要改变构造函数的调用；其余的代码不受算法改变的影响。

​	在这种情况下，构造函数应该返回一个接口值而不是实现的类型。例如，在hash 库中，`crc32.NewIEEE` 和 `adler32.New` 都返回接口类型 `hash.Hash32`。在 Go 程序中，将 `CRC-32` 算法替换为 `Adler-32` 算法仅需要更改构造函数调用；其余代码不受算法更改的影响。

​	类似的方法也允许将各种`crypto`包中的流式加密算法与它们连接在一起的块加密算法分离开来。在 `crypto/cipher` 包中，`Block` 接口指定了块加密的行为，它提供单个数据块的加密。然后，与 `bufio` 包类似，任何实现了该接口的加密包都能被用于构造以 Stream 为接口表示的流加密，而无需知道块加密的细节。

​	`crypto/cipher` 接口如下：

``` go 
type Block interface {
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
}

type Stream interface {
    XORKeyStream(dst, src []byte)
}
```

​	以下是计数器模式（CTR）流的定义，它将块加密转换为流加密；请注意块加密的细节已被抽象掉：

``` go 
// NewCTR returns a Stream that encrypts/decrypts using the given Block in
// counter mode. The length of iv must be the same as the Block's block size.
func NewCTR(block Block, iv []byte) Stream
```

​	`NewCTR` 的应用并不仅限于特定的加密算法和数据源，它适用于任何对 `Block` 接口和 `Stream` 的实现。因为它们返回接口值，所以将 CTR 加密替换为其他加密模式是一种局部化的变更。构造函数的调用过程必须被修改， 但由于其周围的代码只能将它看做 `Stream`，因此它们不会注意到其中的区别。

### 接口和方法

​	由于几乎任何类型都能添加方法，因此几乎任何类型都能满足一个接口。一个很直观的例子就是 `http` 包中定义的 `Handler` 接口。任何实现了 `Handler` 的对象都能够处理 HTTP 请求。

``` go 
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

​	`ResponseWriter` 本身是一个接口，它提供了返回给客户端响应所需的方法。这些方法包括标准的 `Write` 方法，因此 `http.ResponseWriter` 可以在任何需要 `io.Writer` 的地方使用。`Request` 是一个包含来自客户端的请求的解析表示的结构体。

​	为简单起见，我们假设所有的 HTTP 请求都是 GET 方法，而忽略 POST 方法， 这种简化不会影响处理程序的建立方式。这里有个短小却完整的处理程序实现， 它用于记录某个页面被访问的次数。

``` go 
// 简单的计数器服务器。
type Counter struct {
    n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctr.n++
    fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
```

(与我们的主题保持一致，注意`Fprintf`如何打印到`http.ResponseWriter`。) 在一个真正的服务器中，对 `ctr.n` 的访问将需要防止并发访问。请参见`sync`和`atomic`包以获取建议。

​	作为参考，这里演示了如何将这样一个服务器添加到 URL 树的一个节点上。

``` go 
import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)
```

​	但是为什么要将`Counter`定义为一个结构体呢？只需要一个整数即可。（接收器必须为指针，增量操作对于调用者才可见。）

``` go 
// 简单的计数器服务器。
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    *ctr++
    fmt.Fprintf(w, "counter = %d\n", *ctr)
}
```

​	当页面被访问时，怎样通知您的程序去更新一些内部状态呢？为 Web 页面绑定个通道吧。

``` go 
// 每次浏览该通道都会发送一个提醒。
// （可能需要带缓冲的通道。）
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ch <- req
    fmt.Fprint(w, "notification sent")
}
```

​	最后，假设我们想在`/args`上展示调用服务器二进制文件时使用的参数。编写一个函数来打印这些参数是很容易的。

``` go 
func ArgServer() {
    fmt.Println(os.Args)
}
```

​	我们如何将它转换为 HTTP 服务器呢？我们可以将 `ArgServer` 实现为某种可忽略值的方法，不过还有种更简单的方法。 既然我们可以为除指针和接口以外的任何类型定义方法，同样也能为一个函数写一个方法。 `http` 包里有这样的代码：

``` go 
// HandlerFunc 类型是一个适配器，
// 它允许将普通函数用做HTTP处理程序。
// 若 f 是个具有适当签名的函数，
// HandlerFunc(f) 就是个调用 f 的处理程序对象。
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP 调用f(w, req).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}
```

​	`HandlerFunc`是一个具有`ServeHTTP`方法的类型，因此该类型的值可以用作HTTP请求的处理程序。看看该方法的实现：接收器是一个函数`f`，而该方法调用`f`。这可能看起来有些奇怪，但与接收器是通道并且方法在通道上发送的情况并没有太大的区别。

​	为了使`ArgServer`成为一个HTTP服务器，我们首先要修改它，使其具有正确的签名。

``` go 
// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}
```

​	`ArgServer`现在具有与`HandlerFunc`相同的签名，所以它可以被转换为该类型以访问其方法，就像我们将`Sequence`转换为`IntSlice`以访问`IntSlice.Sort`一样。设置它的代码很简洁：

``` go 
http.Handle("/args", http.HandlerFunc(ArgServer))
```

​	当有人访问页面`/args`时，在该页面安装的处理程序具有值`ArgServer`和类型`HandlerFunc`。HTTP服务器将调用该类型的`ServeHTTP`方法，并以`ArgServer`作为接收器，接着调用`ArgServer`（通过`HandlerFunc.ServeHTTP`内的`f(w, req)`调用）。然后，参数将被显示出来。

​	在本节中，我们从结构体、整数、通道和函数制作了一个HTTP服务器，这是因为接口只是方法集，可以为（几乎）任何类型定义方法。

## The blank identifier 空白标识符

We've mentioned the blank identifier a couple of times now, in the context of [`for` `range` loops](https://go.dev/doc/effective_go#for) and [maps](https://go.dev/doc/effective_go#maps). The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly. It's a bit like writing to the Unix `/dev/null` file: it represents a write-only value to be used as a place-holder where a variable is needed but the actual value is irrelevant. It has uses beyond those we've seen already.

我们已经在for range循环和map的背景下提到过几次空白标识符。空白标识符可以被分配或声明为任何类型的任何值，其值会被无害地丢弃。它有点像写到Unix的/dev/null文件：它代表一个只写的值，在需要变量但实际值不相关的情况下作为一个占位符使用。它的用途超出了我们已经看到的那些。

### The blank identifier in multiple assignment 多重赋值中的空白标识符

The use of a blank identifier in a `for` `range` loop is a special case of a general situation: multiple assignment.

在for range循环中使用空白标识符是一般情况下的一个特殊情况：多重赋值。

If an assignment requires multiple values on the left side, but one of the values will not be used by the program, a blank identifier on the left-hand-side of the assignment avoids the need to create a dummy variable and makes it clear that the value is to be discarded. For instance, when calling a function that returns a value and an error, but only the error is important, use the blank identifier to discard the irrelevant value.

如果一个赋值在左边需要多个值，但其中一个值不会被程序使用，在赋值的左边使用空白标识符可以避免创建一个虚拟变量，并清楚地表明该值将被丢弃。例如，当调用一个返回一个值和一个错误的函数，但只有错误是重要的，使用空白标识符来丢弃不相关的值。

``` go 
if _, err := os.Stat(path); os.IsNotExist(err) {
    fmt.Printf("%s does not exist\n", path)
}
```

Occasionally you'll see code that discards the error value in order to ignore the error; this is terrible practice. Always check error returns; they're provided for a reason.

偶尔您会看到为了忽略错误而丢弃错误值的代码；这是很糟糕的做法。一定要检查错误返回；它们的出现是有原因的。

``` go 
// Bad! This code will crash if path does not exist.
fi, _ := os.Stat(path)
if fi.IsDir() {
    fmt.Printf("%s is a directory\n", path)
}
```

### Unused imports and variables 未使用的导入和变量

It is an error to import a package or to declare a variable without using it. Unused imports bloat the program and slow compilation, while a variable that is initialized but not used is at least a wasted computation and perhaps indicative of a larger bug. When a program is under active development, however, unused imports and variables often arise and it can be annoying to delete them just to have the compilation proceed, only to have them be needed again later. The blank identifier provides a workaround.

导入一个包或声明一个变量而不使用它是一个错误。未使用的导入会使程序变得臃肿，并使编译速度变慢，而一个被初始化但未使用的变量至少是一个浪费的计算，并可能表明一个更大的错误。然而，当一个程序处于活跃的开发阶段时，未使用的导入和变量经常出现，为了让编译继续进行而删除它们是很烦人的，但以后又会再次需要它们。空白标识符提供了一个变通办法。

This half-written program has two unused imports (`fmt` and `io`) and an unused variable (`fd`), so it will not compile, but it would be nice to see if the code so far is correct.

这个写了一半的程序有两个未使用的导入（fmt和io）和一个未使用的变量（fd），所以它不会被编译，但是看看到目前为止的代码是否正确也不错。

``` go 
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
}
```

To silence complaints about the unused imports, use a blank identifier to refer to a symbol from the imported package. Similarly, assigning the unused variable `fd` to the blank identifier will silence the unused variable error. This version of the program does compile.

为了消除对未使用的导入的抱怨，使用一个空白标识符来引用导入包中的符号。同样地，将未使用的变量fd分配给空白标识符，将使未使用的变量错误不再出现。这个版本的程序确实可以编译。

``` go 
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader    // For debugging; delete when done.

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
```

By convention, the global declarations to silence import errors should come right after the imports and be commented, both to make them easy to find and as a reminder to clean things up later.

按照约定，用于消除导入错误的全局声明应该紧跟在导入之后，并加以注释，这既是为了让它们容易被发现，也是为了提醒人们以后要清理好。

### Import for side effect 导入的副作用

An unused import like `fmt` or `io` in the previous example should eventually be used or removed: blank assignments identify code as a work in progress. But sometimes it is useful to import a package only for its side effects, without any explicit use. For example, during its `init` function, the `net/http/pprof` package registers HTTP handlers that provide debugging information. It has an exported API, but most clients need only the handler registration and access the data through a web page. To import the package only for its side effects, rename the package to the blank identifier:

像前面例子中的fmt或io这样的未使用的导入，最终应该被使用或删除：空白的赋值表明代码正在进行中。但有时导入一个包只是为了它的副作用，而没有任何明确的用途，是很有用的。例如，在其初始函数中，net/http/pprof 包注册了提供调试信息的 HTTP 处理程序。它有一个导出的API，但大多数客户端只需要注册处理程序，并通过网页访问数据。要想只为它的副作用导入包，请将包重命名为空白标识符：

``` go 
import _ "net/http/pprof"
```

This form of import makes clear that the package is being imported for its side effects, because there is no other possible use of the package: in this file, it doesn't have a name. (If it did, and we didn't use that name, the compiler would reject the program.)

这种形式的导入清楚地表明该包是为了其副作用而被导入的，因为该包没有其他可能的用途：在这个文件中，它没有一个名字。(如果它有，而我们没有使用这个名字，编译器会拒绝这个程序）。

### Interface checks 接口检查

As we saw in the discussion of [interfaces](https://go.dev/doc/effective_go#interfaces_and_types) above, a type need not declare explicitly that it implements an interface. Instead, a type implements the interface just by implementing the interface's methods. In practice, most interface conversions are static and therefore checked at compile time. For example, passing an `*os.File` to a function expecting an `io.Reader` will not compile unless `*os.File` implements the `io.Reader` interface.

正如我们在上面关于接口的讨论中看到的，一个类型不需要明确声明它实现了一个接口。相反，一个类型只是通过实现该接口的方法来实现该接口。在实践中，大多数接口的转换都是静态的，因此在编译时进行检查。例如，将一个*os.File传递给一个期望有io.Reader的函数，除非*os.File实现了io.Reader接口，否则不会被编译。

Some interface checks do happen at run-time, though. One instance is in the `encoding/json` package, which defines a `Marshaler` interface. When the JSON encoder receives a value that implements that interface, the encoder invokes the value's marshaling method to convert it to JSON instead of doing the standard conversion. The encoder checks this property at run time with a [type assertion](https://go.dev/doc/effective_go#interface_conversions) like:

不过，有些接口检查确实发生在运行时。一个例子是在编码/json包中，它定义了一个Marshaler接口。当JSON编码器收到一个实现该接口的值时，编码器会调用该值的marshaling方法将其转换为JSON，而不是做标准转换。编码器在运行时用一个类型断言来检查这个属性，比如：

``` go 
m, ok := val.(json.Marshaler)
```

If it's necessary only to ask whether a type implements an interface, without actually using the interface itself, perhaps as part of an error check, use the blank identifier to ignore the type-asserted value:

如果只需要询问一个类型是否实现了一个接口，而没有实际使用该接口本身，也许是作为错误检查的一部分，使用空白标识符来忽略类型断言的值：

``` go 
if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
```

One place this situation arises is when it is necessary to guarantee within the package implementing the type that it actually satisfies the interface. If a type—for example, `json.RawMessage`—needs a custom JSON representation, it should implement `json.Marshaler`, but there are no static conversions that would cause the compiler to verify this automatically. If the type inadvertently fails to satisfy the interface, the JSON encoder will still work, but will not use the custom implementation. To guarantee that the implementation is correct, a global declaration using the blank identifier can be used in the package:

这种情况出现的一个地方是，当有必要在实现该类型的包中保证它确实满足接口。如果一个类型——例如json.RawMessage——需要一个自定义的JSON表示法，它应该实现json.Marshaler，但是没有静态转换可以使编译器自动验证这一点。如果该类型无意中未能满足接口，JSON编码器仍将工作，但不会使用自定义的实现。为了保证实现的正确性，可以在包中使用一个使用空白标识符的全局声明：

``` go 
var _ json.Marshaler = (*RawMessage)(nil)
```

In this declaration, the assignment involving a conversion of a `*RawMessage` to a `Marshaler` requires that `*RawMessage` implements `Marshaler`, and that property will be checked at compile time. Should the `json.Marshaler` interface change, this package will no longer compile and we will be on notice that it needs to be updated.

在这个声明中，涉及到将*RawMessage转换为Marshaler的赋值要求*RawMessage实现Marshaler，并且该属性将在编译时被检查。如果json.Marshaler接口发生变化，这个包将不再编译，我们将注意到它需要被更新。

The appearance of the blank identifier in this construct indicates that the declaration exists only for the type checking, not to create a variable. Don't do this for every type that satisfies an interface, though. By convention, such declarations are only used when there are no static conversions already present in the code, which is a rare event.

在这个结构中出现的空白标识符表明，这个声明只是为了进行类型检查而存在，而不是为了创建一个变量。不过，不要对每个满足接口的类型都这样做。根据约定，只有在代码中没有静态转换时才会使用这种声明，而这是一种罕见的情况。

## 嵌入

​	Go语言并没有提供传统的基于类型的子类概念，但它可以通过在结构体或接口中嵌入类型来"借用"实现的一部分。

​	接口嵌入非常简单。我们之前提到了`io.Reader`和`io.Writer`接口; 以下是它们的定义。

``` go 
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

​	`io`包还导出了几个其他接口，用于指定可以实现多个这样的方法的对象。例如，有一个`io.ReadWriter`，一个包含`Read`和`Write`方法的接口。我们可以通过显式列出这两种方法来指定`io.ReadWriter`，但通过嵌入这两个接口来形成新接口更容易且更具启示性，就像这样：

``` go 
// ReadWriter是将Reader和Writer接口结合在一起的接口。
type ReadWriter interface {
    Reader
    Writer
}
```

​	这就是它看起来的样子：`ReadWriter`可以像`Reader`和`Writer`一样执行操作；它是嵌入接口们的并集。只有接口可以嵌入到接口中。

​	相同的基本思想也适用于结构体，但影响更为深远。`bufio`包有两个结构体类型`bufio.Reader`和`bufio.Writer`，它们分别实现了来自`io`包的相应接口。而且`bufio`还实现了一个缓冲reader/writer，它通过使用嵌入将reader 和writer组合成一个结构体来实现：它列出了结构体内的类型，但不给它们字段名。

``` go 
// ReadWriter存储Reader和Writer的指针。
// 它实现了io.ReadWriter。
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
```

​	嵌入的元素是结构体的指针，在使用之前必须初始化为指向有效的结构体。`ReadWriter`结构体可以编写为：

``` go 
type ReadWriter struct {
    reader *Reader
    writer *Writer
}
```

​	但是为了促进字段的方法并满足`io`接口，我们还需要提供转发方法，就像这样：

``` go 
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
    return rw.reader.Read(p)
}
```

​	通过直接嵌入结构体，我们避免了这种繁琐的工作。嵌入类型的方法会自动继承，这意味着`bufio.ReadWriter`不仅具有`bufio.Reader`和`bufio.Writer`的方法，还满足了所有三个接口：`io.Reader`，`io.Writer`和`io.ReadWriter`。

​	嵌入与子类化的一个重要区别是，当我们嵌入类型时，该类型的方法成为外部类型的方法，但是当它们被调用时，方法的接收器是内部类型，而不是外部类型。在我们的例子中，当`bufio.ReadWriter`的`Read`方法被调用时，它的效果与上面写出的转发方法完全相同；接收器是`ReadWriter`的`reader`字段，而不是`ReadWriter`本身。

​	嵌入也可以是一个简单的便利。这个例子展示了一个嵌入字段和一个常规命名字段。

``` go 
type Job struct {
    Command string
    *log.Logger
}
```

​	现在，`Job`类型具有`*log.Logger`的`Print`，`Printf`，`Println`和其他方法。当然，我们可以给`Logger`一个字段名，但没有必要这样做。现在，一旦初始化，我们就可以用`Job`记录：

``` go 
job.Println("starting now...")
```

​	`Logger`是`Job`结构的常规字段，因此我们可以在`Job`的构造函数中按照通常的方式进行初始化，例如：

``` go 
func NewJob(command string, logger *log.Logger) *Job {
    return &Job{command, logger}
}
```

或者使用复合字面值，例如：

``` go 
job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
```

​	如果我们需要直接引用嵌入字段，则字段名称作为字段名，忽略包限定符，就像在我们的`ReadWriter`结构的`Read`方法中一样。在这里，如果我们需要访问`Job`变量`job`的`*log.Logger`，则会写成`job.Logger`，这对于我们想完善`Logger`的方法很有用。

``` go 
func (job *Job) Printf(format string, args ...interface{}) {
    job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}
```

​	嵌入类型引入了名称冲突的问题，但解决它们的规则很简单。首先，一个名为`X`的字段或方法会覆盖在类型的更深嵌套部分中出现的任何其他`X`项。如果`log.Logger`包含一个名为`Command`的字段或方法，则`Job`的`Command`字段将覆盖它。

​	其次，如果在相同的嵌套级别上出现相同的名称，通常会出现错误；如果Job结构包含另一个名为`Logger`的字段或方法，则嵌入`log.Logger`将是错误的。但是，如果在程序类型定义之外从未提到重复的名称，则没有问题。这种限定提供了一些保护，以防从外部嵌入的类型进行更改；如果添加了与另一个子类型中的另一个字段冲突的字段，但若两个字段都从未被使用，则没有问题。（不好理解，以下给出ChatGPT给出的示例）

> 下面是一个示例，演示了嵌入类型中出现相同名称时，只要名称没有被使用，程序就不会有问题：
> ```go 
> package main
> 
> import "fmt"
> 
> type A struct {
>     Field int
> }
> 
> type B struct {
>     Field int
> }
> 
> type C struct {
>     A
>     B
> }
> 
> func main() {
>     c := C{A{1}, B{2}}
>     fmt.Println(c.A.Field) // 输出 1
>     fmt.Println(c.B.Field) // 输出 2
> }
> 
> 
> ```
>
> 在这个例子中，我们定义了三个类型 `A`、`B` 和 `C`。类型 `C` 嵌入了类型 `A` 和类型 `B`，这两个类型中都包含一个名为 `Field` 的字段。然而，在程序中，我们只是用 `A.Field` 和 `B.Field` 分别访问了这两个字段，而没有直接使用 `Field` 这个名称。因此，即使 `A` 和 `B` 中的 `Field` 字段有名称冲突，程序仍然能够正确编译和运行。
>
> 这个例子说明了 Effective Go 中提到的规则，即只要冲突的名称没有被使用到，就不会有问题。但需要注意的是，这并不代表名称冲突是可以被忽略的，因为一旦使用了相同名称的字段或方法，就会导致编译错误或运行错误。
>
> 
>



## 并发

### 通过通信共享 

​	并发编程是一个广泛的主题，在这里只有一些与 Go 有关的亮点。

​	在许多环境中，并发编程由于需要实现对共享变量的正确访问而变得困难。Go 鼓励采用一种不同的方法，在其中共享值通过通道传递，实际上，不由执行的分离线程主动共享。任何时候只有一个 goroutine 可以访问该值。由于设计原因，数据竞争不会发生。为了鼓励这种思考方式，我们将其简化为一个口号：

> Do not communicate by sharing memory; instead, share memory by communicating.
>
> 不要通过共享内存来通信；相反，通过通信来共享内存。

​	这种方法可能会走得太远。例如，引用计数最好通过在整数变量周围放置互斥锁来完成。但作为一种高级方法，使用通道来控制访问，使编写清晰、正确的程序更容易。

​	从某种角度来看，这种模型可以理解为在一个 CPU 上运行的典型单线程程序。它不需要同步原语。现在再运行另一个实例；它也不需要同步。现在让这两个实例通信；如果通信是同步器，那么就没有其他同步的需要。例如，Unix 管道完全符合这个模型。虽然 Go 的并发方法起源于 Hoare 的通信顺序进程 (CSP)，但它也可以看作是 Unix 管道的类型安全的概括。

### Goroutines

​	它们被称为 goroutines，因为现有的术语——线程、协程、进程等——传达了不准确的内涵。goroutine 有一个简单的模型：它是一个与同一地址空间中的其他 goroutine 并发执行的函数。它是轻量级的，成本几乎只有栈空间的分配。栈开始很小，所以它们很便宜，并通过根据需要分配（和释放）堆存储来增长。

​	Goroutines 被多路复用到多个操作系统线程中，因此如果其中一个线程阻塞，例如在等待 I/O 时，其他线程继续运行。它们的设计隐藏了许多线程创建和管理的复杂性。

​	在一个函数或方法调用之前加上 `go` 关键字以在新的 goroutine 中运行该调用。当调用完成时，goroutine 静默退出。（效果类似于 Unix shell 的`&`符号，用于在后台运行命令。）

``` go 
go list.Sort()  // 并发运行 list.Sort；不等待它。
```

​	函数字面量在goroutine调用中非常方便。

``` go 
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // 注意该括号 - 必须调用该函数。
}
```

​	在Go中，函数字面量是闭包：实现上确保函数引用的变量在其活动期间存活。

​	这些示例并不太实用，因为这些函数无法发出完成信号。为此，我们需要使用通道。

### 通道

​	与映射一样，通道是使用`make`分配的，而生成的值充当底层数据结构的引用。如果提供了可选的整数参数，则会为通道设置缓冲区大小。默认值为零，表示无缓冲或同步通道。

``` go 
ci := make(chan int)            // 整数无缓冲通道
cj := make(chan int, 0)         // 整数无缓冲通道
cs := make(chan *os.File, 100)  // 100个指向文件的缓冲通道
```

​	无缓冲通道结合了通信 - 交换值 - 与同步 —— 确保两个计算（goroutines）处于已知状态。

​	有很多使用通道的好习惯。这里有一个开始的例子。在前一节中，我们在后台启动了一个排序。通道可以让启动的goroutine等待排序完成。

``` go 
c := make(chan int)  // 分配一个通道。
// 在goroutine中启动排序；当它完成时，向通道发出信号。
go func() {
    list.Sort()
    c <- 1  // 发送信号；值不重要。
}()
doSomethingForAWhile()
<-c   // 等待排序完成；忽略已发送的值。
```

​	接收器始终阻塞，直到有数据可接收。如果通道无缓冲，则发送器将阻塞，直到接收器接收到该值。如果通道具有缓冲区，则发送器仅阻塞，直到该值已复制到缓冲区；如果缓冲区已满，则这意味着等待，直到某个接收器检索到一个值。

A buffered channel can be used like a semaphore, for instance to limit throughput. In this example, incoming requests are passed to `handle`, which sends a value into the channel, processes the request, and then receives a value from the channel to ready the "semaphore" for the next consumer. The capacity of the channel buffer limits the number of simultaneous calls to `process`.

缓冲通道可以像信号灯一样使用，例如用来限制吞吐量。在这个例子中，传入的请求被传递给handle，handle向通道发送一个值，处理请求，然后从通道接收一个值，为下一个消费者准备 "信号"。通道缓冲区的容量限制了同时调用处理的数量。

​	缓冲通道可用作信号量，例如限制吞吐量。在此示例中，传入的请求传递给`handle`，handle将一个值发送到通道中，处理请求，然后从通道中接收一个值以准备好下一个使用者的"信号量"。通道缓冲区的容量限制了对进程的同时调用数量。

``` go 
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // 等待活动队列排空。
    process(r)  // 可能需要很长时间。
    <-sem       // 完成；启用下一个请求运行。
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // 不等待handle完成。
    }
}
```

​	一旦`MaxOutstanding`个处理程序正在执行，任何进一步的处理程序都会尝试发送到已填充的通道缓冲区并被阻塞，直到现有的处理程序完成并从缓冲区接收。

​	但这种设计存在问题：即使只有`MaxOutstanding`中的一部分可以运行，`Serve`还是会为每个传入的请求创建一个新的goroutine。结果，如果请求过于频繁，程序可能会消耗无限的资源。我们可以通过更改`Serve`来限制goroutine的创建来解决这个问题。下面是一个显而易见的解决方案，但要注意它有一个bug，我们随后会修复：

``` go 
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req) // 有Bug;请参见下面的解释。
            <-sem
        }()
    }
}
```

​	bug在于，在Go for循环中，循环变量在每次迭代中都会被重用，因此`req`变量会在所有goroutine之间共享。这不是我们想要的。我们需要确保对于每个goroutine，`req`都是唯一的。下面是一种将`req`的值作为参数传递给goroutine中的闭包的方法：

``` go 
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func(req *Request) {
            process(req)
            <-sem
        }(req)
    }
}
```

​	将此版本与之前的版本进行比较，查看闭包声明和运行方式的差异。另一个解决方案是仅创建一个具有相同名称的新变量，如以下示例所示：

``` go 
func Serve(queue chan *Request) {
    for req := range queue {
        req := req // 为goroutine创建req的新实例。
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}
```

写成

``` go 
req := req
```

可能看起来有些奇怪，但在Go中这样做是合法且惯用的。您将获得一个具有相同名称的新变量的新版本，它在本地有意遮蔽了循环变量，但对于每个goroutine都是唯一的。

​	回到编写服务器的一般问题上，另一种管理资源的好方法是启动一定数量的处理goroutine，它们都从请求通道中读取。goroutine的数量限制了对`process`的同时调用次数。此`Serve`函数还接受一个通道，在该通道上将告诉它退出；在启动goroutine后，它会阻塞接收该通道。

``` go 
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // 启动处理程序
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // 等待被告知退出。
}
```

### 通道的通道

​	Go的一个最重要的特性是通道是一种一等值，可以像其他值一样分配和传递。这种特性的常见用途是实现安全的并行复用。

​	在前一节的示例中，`handle`是一个理想化的处理程序，但我们没有定义它处理的类型。如果该类型包含一个通道来回复，每个客户端都可以提供自己的答案路径。这是`Request`类型的示意定义。

``` go 
type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}
```

​	客户端提供一个函数和它的参数，以及一个请求对象内部的通道，用于接收答案。

``` go 
func sum(a []int) (s int) {
    for _, v := range a {
        s += v
    }
    return
}

request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
// Send request
clientRequests <- request
// Wait for response.
fmt.Printf("answer: %d\n", <-request.resultChan)
```

​	在服务器端，处理函数是唯一需要更改的内容。

``` go 
func handle(queue chan *Request) {
    for req := range queue {
        req.resultChan <- req.f(req.args)
    }
}
```

​	显然，要使其变得更加现实，需要做更多的工作，但是这段代码是一个速率受限、并行且非阻塞的RPC系统的框架，而且没有一个互斥锁。

### 并行

​	这些想法的另一个应用是将计算并行化到多个CPU核心上。如果计算可以分成可以独立执行的不同部分，那么它就可以被并行化，并使用信道来指示每个部分何时完成。

​	假设我们需要对一组项的向量执行一个昂贵的操作，而且每个项的操作值是独立的，就像这个理想化的例子：

``` go 
type Vector []float64

// 对v[i]、v[i+1] … 一直到v[n-1]执行操作。
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
        v[i] += u.Op(v[i])
    }
    c <- 1    // 指示该部分已完成
}
```

​	我们在一个循环中独立地启动每个部分，每个CPU一个部分。它们可以以任何顺序完成，但这无关紧要；我们只需通过在启动所有goroutine后从通道中取出信号来计算完成信号的数量。

``` go 
const numCPU = 4 // CPU核心数

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // 可选的缓冲区大小。
    for i := 0; i < numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    // 取出信道中的信号。
    for i := 0; i < numCPU; i++ {
        <-c    // 等待一个任务完成
    }
    // 全部完成。
}
```

​	与其为numCPU创建一个常量值，我们可以询问运行时适当的值是多少。函数`runtime.NumCPU`返回机器中硬件CPU核心的数量，因此我们可以编写：

``` go 
var numCPU = runtime.NumCPU()
```

​	还有一个函数`runtime.GOMAXPROCS`，它报告（或设置）Go程序可以同时运行的用户指定的核心数。默认值为`runtime.NumCPU`的值，但可以通过设置类似命名的shell环境变量或调用带有正整数参数的函数来覆盖。调用它时，如果使用零，则只是查询该值。因此，如果我们想遵守用户的资源请求，我们应该编写：

``` go 
var numCPU = runtime.GOMAXPROCS(0)
```

​	一定要注意不要混淆并发的概念 - 结构化一个程序为独立执行的组件 - 和并行化的概念 - 在多个CPU上并行执行计算以提高效率。尽管Go的并发特性可以使一些问题易于结构化为并行计算，但Go是一种并发语言，而不是一种并行语言，不是所有并行化问题都适合Go的模型。有关区别的讨论，请参见[这一博客文章]({{< ref "/goBlog/2013/ConcurrencyIsNotParallelism" >}})中引用的演讲。

### 一个泄漏的缓冲区

​	并发编程的工具甚至可以使非并发的想法更容易表达。以下是一个从RPC包中抽象出来的示例。客户端 goroutine 循环从某些来源（例如网络）接收数据。为避免分配和释放缓冲区，它保持一个空闲列表，并使用缓冲的通道来表示它。如果通道为空，则会分配一个新的缓冲区。一旦消息缓冲区准备好，它就会被发送到 `serverChan` 上的服务器。

``` go 
var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
    for {
        var b *Buffer
        // 如果有可用的缓冲区，则获取一个；如果没有，则分配一个新的。
        select {
        case b = <-freeList:
            // 已获取一个；无需进行更多操作。
        default:
            // 没有空闲的，所以分配一个新的。
            b = new(Buffer)
        }
        load(b)              // 从网络读取下一条消息。
        serverChan <- b      // 发送到服务器。
    }
}
```

​	服务器循环从客户端接收每条消息，处理它并将缓冲区返回到空闲列表中。

``` go 
func server() {
    for {
        b := <-serverChan    // 等待任务。
        process(b)
        // 如果有空间，则重用缓冲区。
        select {
        case freeList <- b:
            // 缓冲区在空闲列表上；无需进行更多操作。
        default:
            // 空闲列表已满，继续执行。
        }
    }
}
```

​	客户端尝试从 `freeList` 中检索缓冲区；如果没有可用的，则分配一个新的。服务器将 `b` 发送到 `freeList` 上，除非列表已满，在这种情况下，缓冲区将被丢弃以供垃圾回收器回收。（在 `select` 语句中的`default`子句在没有其他 case 准备好的情况下执行，这意味着`selects`永远不会阻塞。）这个实现只用了几行代码就建立了一个泄漏的桶式空闲列表，依靠缓冲的通道和垃圾回收器进行簿记(bookkeeping)。

## 错误

​	库程序常常需要向调用者返回某种错误指示。如前所述，Go 的多值返回使得返回详细的错误描述与常规返回值一样容易。使用这个特性提供详细的错误信息是一个好的编程风格。例如，正如我们将看到的那样，`os.Open` 不仅在失败时返回一个 nil 指针，它还返回一个错误值，描述出了问题所在。

​	按照惯例，错误类型为 `error`，它是一个简单的内置接口。

``` go 
type error interface {
    Error() string
}
```

​	一个库的作者可以自由地在底层实现这个接口，使用更丰富的模型，不仅可以看到错误，还可以提供一些上下文信息。正如前面提到的，除了通常的`*os.File`返回值外，`os.Open`还返回一个错误值。如果文件成功打开，错误将为`nil`，但是当出现问题时，它将包含一个`os.PathError`：

``` go 
// PathError 记录错误和导致错误的操作和文件路径。
type PathError struct {
    Op string    // "open", "unlink", etc.
    Path string  // 关联的文件。
    Err error    // 系统调用返回的错误。
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```

​	`PathError`的`Error`方法生成像这样的字符串：

``` go 
open /etc/passwx: no such file or directory
```

​	这样的错误，包括有问题的文件名、操作和触发它的操作系统错误，即使在远离引起错误的调用处打印，也很有用；它比纯粹的"no such file or directory"更具信息性。

​	在可能的情况下，错误字符串应该标识它们的来源，比如通过具有命名操作或生成错误的包的前缀。例如，在`image`包中，由于未知格式而导致解码错误的字符串表示为"image: unknown format"。

​	关心精确错误详情的调用方可以使用类型选择或类型断言查找特定错误并提取详细信息。对于`PathErrors`，这可能包括检查内部`Err`字段以进行可恢复的故障。

``` go 
for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles()  // Recover some space.
        continue
    }
    return
}
```

​	这里的第二个`if`语句是另一个类型断言。如果失败，`ok`将为false，`e`将为`nil`。如果成功，`ok`将为true，这意味着错误的类型是`*os.PathError`，那么`e`也是这样的，我们可以检查更多关于错误的信息。

### Panic

The usual way to report an error to a caller is to return an `error` as an extra return value. The canonical `Read` method is a well-known instance; it returns a byte count and an `error`. But what if the error is unrecoverable? Sometimes the program simply cannot continue.

向调用者报告错误的通常方法是返回一个错误作为额外的返回值。典型的Read方法是一个著名的例子；它返回一个字节数和一个错误。但是如果错误是无法恢复的呢？有时程序根本无法继续。

​	向调用者报告错误的通常方式是作为额外的返回值返回`error`。经典的 `Read` 方法是一个众所周知的例子，它返回一个字节数和一个错误。但是如果错误是不可恢复的呢？有时程序就是不能继续执行。

For this purpose, there is a built-in function `panic` that in effect creates a run-time error that will stop the program (but see the next section). The function takes a single argument of arbitrary type—often a string—to be printed as the program dies. It's also a way to indicate that something impossible has happened, such as exiting an infinite loop.

为此，有一个内置的函数panic，它实际上创造了一个运行时错误，将停止程序（但见下一节）。该函数需要一个任意类型的参数——通常是一个字符串——在程序死亡时被打印出来。这也是一种表示发生了不可能的事情的方法，比如退出一个无限循环。

​	为此，有一个内置函数 `panic`，实际上创建一个运行时错误，将停止程序（但请参见下一节）。该函数接受一个任意类型的单个参数，通常是一个字符串，用于在程序停止时打印。这也是一种指示发生了不可能的事情（比如退出一个无限循环）的方法。

``` go 
// A toy implementation of cube root using Newton's method.
// 一个使用牛顿法的立方根玩具实现。
func CubeRoot(x float64) float64 {
    z := x/3   // 任意初始值
    for i := 0; i < 1e6; i++ {
        prevz := z
        z -= (z*z*z-x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    // A million iterations has not converged; something is wrong. 百万次迭代没有收敛；有些问题。
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}
```

​	这只是一个示例，但真正的库函数应该避免 `panic`。如果问题可以被掩盖或绕过，让事情继续运行总是比将整个程序关闭更好。一个可能的反例是在初始化期间：如果库确实不能设置自身，可能会出现 panic，可以这么说。

``` go 
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

### Recover

​	当调用`panic`时（包括隐式调用，例如索引超出切片界限或类型断言失败的运行时错误），它会立即停止当前函数的执行，并开始解开the goroutine的栈，同时运行任何延迟的函数。如果这种解开栈的操作到达the goroutines的顶部，则程序将停止。但是，可以使用内置的`recover`函数来重新获得the goroutine的控制权并恢复正常执行。

​	调用`recover`会停止解开栈的操作，并返回传递给`panic`的参数。因为解开栈时运行的唯一代码是在延迟的函数中，所以`recover`仅在延迟的函数中有用。

​	`recover`的一个应用是在服务器内部关闭一个失败的goroutine，而不会杀死其他正在执行的goroutines。

``` go 
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

​	在这个例子中，如果 `do(work)` 发生恐慌，结果将被记录，并且 goroutine 将在不干扰其他 goroutine 的情况下干净地退出。在延迟闭包中不需要执行任何其他操作；调用 `recover` 将完全处理该条件。

​	因为 `recover` 总是返回 `nil`，除非直接从延迟函数中调用，所以延迟代码可以调用库例程，这些库例程本身使用 `panic` 和 `recover` 而不会失败。例如，`safelyDo` 中的延迟函数在调用 `recover` 之前可能会调用一个日志记录函数，而该日志记录代码将在恐慌状态下不受影响地运行。

With our recovery pattern in place, the `do` function (and anything it calls) can get out of any bad situation cleanly by calling `panic`. We can use that idea to simplify error handling in complex software. Let's look at an idealized version of a `regexp` package, which reports parsing errors by calling `panic` with a local error type. Here's the definition of `Error`, an `error` method, and the `Compile` function.

有了我们的恢复模式，do函数（以及它调用的任何东西）可以通过调用panic干净利落地摆脱任何糟糕的情况。我们可以用这个想法来简化复杂软件中的错误处理。让我们看看一个理想化版本的regexp包，它通过调用本地错误类型的panic来报告解析错误。这里有Error的定义，一个错误方法，以及Compile函数。

​	有了我们的恢复模式，`do` 函数（以及它调用的任何内容）都可以通过调用 `panic` 干净地摆脱任何不良情况。我们可以利用这个想法简化复杂软件中的错误处理。让我们来看一个`regexp`包的理想化版本，它通过使用一个本地的错误类型，通过调用 `panic` 来报告解析错误。这是 `Error`、`Error` 方法和 `Compile` 函数的定义。

``` go 
// Error 是解析错误的类型；它满足 error 接口。
type Error string
func (e Error) Error() string {
    return string(e)
}

// error 是 *Regexp 的一个方法，它通过 panic 来报告解析错误。
func (regexp *Regexp) error(err string) {
    panic(Error(err))
}

// Compile 返回正则表达式的解析表示。
func Compile(str string) (regexp *Regexp, err error) {
    regexp = new(Regexp)
    // 如果出现解析错误，doParse 将 panic。
    defer func() {
        if e := recover(); e != nil {
            regexp = nil    // 清空返回值。
            err = e.(Error) // 如果不是解析错误，则重新 panic。
        }
    }()
    return regexp.doParse(str), nil
}
```

​	如果 `doParse` 发生恐慌，恢复块将把返回值设置为 `nil`——延迟函数可以修改命名返回值。然后，它将在将 `err` 赋值时检查问题是否为解析错误，方法是断言它具有本地类型 `Error`。如果不是，则类型断言将失败，导致运行时错误，其将继续栈展开，就好像没有中断一样。这个检查意味着如果发生了一些意外情况，比如越界，代码将失败，即使我们使用 `panic` 和 `recover` 来处理解析错误也是如此。

​	有了错误处理，`error` 方法（因为它是绑定到一个类型的方法，因此它具有与内置 `error` 类型相同的名称是可以的，甚至是自然的）使报告解析错误变得容易，而不必手动展开解析栈：

``` go 
if pos == 0 {
    re.error("'*' illegal at start of expression")
}
```

​	虽然这种模式很有用，但应该仅在一个包内使用。`Parse`函数将其内部的`panic`调用转换为`error`值，不会向客户端暴露`panics`。这是一个值得遵循的好规则。

​	顺便说一下，如果出现实际错误，这种重新panic的惯用语会更改panic的值。但是，原始失败和新的失败都会在崩溃报告中呈现，因此问题的根本原因仍然可见。因此，这种简单的重新panic方法通常足够了——毕竟这是一个崩溃——但如果您只想显示原始值，您可以编写更多的代码来过滤意外问题，并使用原始错误重新panic。这留给读者作为一个练习。

## A web server 一个web服务器

Let's finish with a complete Go program, a web server. This one is actually a kind of web re-server. Google provides a service at `chart.apis.google.com` that does automatic formatting of data into charts and graphs. It's hard to use interactively, though, because you need to put the data into the URL as a query. The program here provides a nicer interface to one form of data: given a short piece of text, it calls on the chart server to produce a QR code, a matrix of boxes that encode the text. That image can be grabbed with your cell phone's camera and interpreted as, for instance, a URL, saving you typing the URL into the phone's tiny keyboard.

让我们用一个完整的Go程序来结束，一个网络服务器。这个实际上是一种网络再服务器。Google在chart.apis.google.com上提供了一个服务，可以将数据自动格式化为图表和图形。不过，它很难交互使用，因为您需要把数据作为查询放入URL。这里的程序为一种形式的数据提供了一个更好的接口：给定一个简短的文本，它调用图表服务器来产生一个QR码，一个编码文本的方框矩阵。该图像可以用手机的摄像头抓取，并解释为，例如，一个URL，省得您在手机的小键盘上输入URL。

Here's the complete program. An explanation follows.

这里是完整的程序。下面是一个解释。

``` go 
package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
    flag.Parse()
    http.Handle("/", http.HandlerFunc(QR))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func QR(w http.ResponseWriter, req *http.Request) {
    templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
```

The pieces up to `main` should be easy to follow. The one flag sets a default HTTP port for our server. The template variable `templ` is where the fun happens. It builds an HTML template that will be executed by the server to display the page; more about that in a moment.

到main为止的部分应该很容易理解。一个标志是为我们的服务器设置一个默认的HTTP端口。模板变量templ是最有趣的地方。它建立了一个HTML模板，将由服务器执行，以显示页面；稍后会有更多关于这个的内容。

The `main` function parses the flags and, using the mechanism we talked about above, binds the function `QR` to the root path for the server. Then `http.ListenAndServe` is called to start the server; it blocks while the server runs.

主函数解析标志，并使用我们上面谈到的机制，将函数QR与服务器的根路径绑定。然后调用http.ListenAndServe来启动服务器；当服务器运行时，它就会阻塞。

`QR` just receives the request, which contains form data, and executes the template on the data in the form value named `s`.

QR只是接收包含表单数据的请求，并在名为s的表单值中的数据上执行模板。

The template package `html/template` is powerful; this program just touches on its capabilities. In essence, it rewrites a piece of HTML text on the fly by substituting elements derived from data items passed to `templ.Execute`, in this case the form value. Within the template text (`templateStr`), double-brace-delimited pieces denote template actions. The piece from `{{if .}}` to `{{end}}` executes only if the value of the current data item, called `.` (dot), is non-empty. That is, when the string is empty, this piece of the template is suppressed.

模板包html/template很强大；这个程序只是触及了它的功能。从本质上讲，它通过替换从传递给templ.Execute的数据项中得到的元素，在本例中是表单值，来临时重写一段HTML文本。在模板文本（templateStr）中，以双括号分隔的部分表示模板动作。从{{if .}}到{{end}}的部分只有在当前数据项的值（称为.（点））非空时才会执行。也就是说，当字符串为空时，模板的这一块被抑制。

The two snippets `{{.}}` say to show the data presented to the template—the query string—on the web page. The HTML template package automatically provides appropriate escaping so the text is safe to display.

这两个切片{{.}}说的是要在网页上显示呈现给模板的数据——查询字符串。HTML模板包会自动提供适当的转义，这样文本就可以安全地显示。

The rest of the template string is just the HTML to show when the page loads. If this is too quick an explanation, see the [documentation](https://go.dev/pkg/html/template/) for the template package for a more thorough discussion.

模板字符串的其余部分只是在页面加载时显示的HTML。如果这解释得太快，请看模板包的文档以获得更全面的讨论。

And there you have it: a useful web server in a few lines of code plus some data-driven HTML text. Go is powerful enough to make a lot happen in a few lines.

就这样：几行代码加上一些数据驱动的HTML文本，就有了一个有用的网络服务器。Go的功能足够强大，几行代码就能实现很多事情。