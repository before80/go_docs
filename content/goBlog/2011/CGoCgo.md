+++
title = "C? Go? Cgo!"
weight = 26
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# C? Go? Cgo!

https://go.dev/blog/cgo

Andrew Gerrand
17 March 2011

2011年3月17日

## 简介

​	Cgo 允许 Go 包调用 C 代码。通过在 Go 源文件中加入一些特殊特性，cgo 就能输出 Go 和 C 文件，并将它们合并成一个单独的 Go 包。

​	让我们通过一个例子来了解一下。下面是一个 Go 包，它提供了两个函数 Random 和 Seed，分别调用了 C 语言的 random 和 srandom 函数。

```go
package rand

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
    return int(C.random())
}

func Seed(i int) {
    C.srandom(C.uint(i))
}
```

​	现在让我们从 import 语句开始分析。

​	rand 包引入了 C，但是在 Go 标准库中并没有这个包。这是因为 C 是一个"伪包"，是 cgo 特殊解释的名称，用于表示 C 的命名空间。

​	rand 包包含了对 C 包的四个引用：对 C.random 和 C.srandom 的调用，对 C.uint(i) 的转换，以及 import 语句。

​	Random 函数调用标准的 C 库的 random 函数，并返回其结果。在 C 语言中，random 函数返回类型为 long，cgo 将其表示为类型 C.long。在被这个包以外的 Go 代码使用前，需要将它转换成 Go 类型。这里使用了普通的 Go 类型转换：

```go
func Random() int {
    return int(C.random())
}
```

​	下面是另一种等价的函数，它使用了一个临时变量，更明确地展示了类型转换：

```go
func Random() int {
    var r C.long = C.random()
    return int(r)
}
```

​	Seed 函数相当于在转换方面做了相反的事情。它接受一个常规的 Go int 类型，将其转换成 C 语言中的 unsigned int 类型，并将其传递给 srandom 函数。

```go
func Seed(i int) {
    C.srandom(C.uint(i))
}
```

​	注意，cgo知道无符号int类型为C.uint；关于这些数字类型名称的完整列表，见[cgo文档](https://go.dev/cmd/cgo)。

​	需要注意的是，cgo 将 unsigned int 类型表示为 C.uint。关于这些数值类型名称的完整列表，可以参考 cgo 文档。

​	这个例子中唯一没有解释的细节是 import 语句上方的注释。

```go
/*
#include <stdlib.h>
*/
import "C"
```

​	Cgo 会识别这个注释。以 `#cgo` 和一个空格字符开头的所有行都会被删除；它们成为 cgo 的指令。剩余的行在编译包中的 C 部分时将被用作头文件。在本例中，这些行只是一个单独的 `#include` 语句，但几乎可以是任何 C 代码。`#cgo` 指令用于在构建包的 C 部分时提供编译器和链接器的标志。

​	有一个限制：如果程序使用任何 `//export` 指令，则注释中的 C 代码只能包括声明（`extern int f();`），而不是定义（`int f() { return 1; }`）。您可以使用 `//export` 指令使 Go 函数可被 C 代码访问。

​	`#cgo` 和 `//export` 指令在 [cgo 文档](https://go.dev/cmd/cgo/)中有详细介绍。

## 字符串和其他类型

​	与Go不同，C没有显式的字符串类型。在C中，字符串由以零结尾的char数组表示。

​	在Go和C之间进行字符串转换时，使用C.CString、C.GoString和C.GoStringN函数。这些转换会复制字符串数据。

​	下面的示例实现了一个打印函数，使用C标准库中的fputs函数将字符串写入标准输出：

```go
package print

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func Print(s string) {
    cs := C.CString(s)
    C.fputs(cs, (*C.FILE)(C.stdout))
    C.free(unsafe.Pointer(cs))
}
```

​	C代码的内存分配不会被Go的内存管理器所知道。当使用C.CString（或任何C内存分配）创建C字符串时，必须记住在使用完毕后通过调用C.free来释放内存。

​	C.CString调用返回一个指向char数组开头的指针，因此在函数退出之前，我们将其转换为unsafe.Pointer，并使用C.free释放内存分配。在cgo程序中，通常的用法是在分配后立即推迟释放（尤其是当后续的代码比单个函数调用更复杂时），如Print的以下重写：

```go
func Print(s string) {
    cs := C.CString(s)
    defer C.free(unsafe.Pointer(cs))
    C.fputs(cs, (*C.FILE)(C.stdout))
}
```

## 构建cgo程序包

​	要构建cgo程序包，只需像往常一样使用[go build](https://go.dev/cmd/go/#hdr-Compile_packages_and_dependencies)或[go install](https://go.dev/cmd/go/#hdr-Compile_and_install_packages_and_dependencies)。go工具会识别特殊的"C"导入，并自动为这些文件使用cgo。

## 更多cgo资源

​	[cgo 命令](https://go.dev/cmd/cgo/)文档介绍了C伪包和构建过程的更多细节。Go树中的[cgo示例](https://go.dev/misc/cgo/)展示了更高级的概念。

​	最后，如果您对这些内容的内部工作原理感到好奇，请查看运行时包的[cgocall.go](https://go.dev/src/runtime/cgocall.go)文件中的介绍性注释。
