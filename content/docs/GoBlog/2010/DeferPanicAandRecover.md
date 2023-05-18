+++
title = "Defer, Panic, and Recover"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Defer, Panic, and Recover

https://go.dev/blog/defer-panic-and-recover

Andrew Gerrand
4 August 2010

2010 年 8 月 4 日

​	Go 有常见的控制流机制：if、for、switch、goto。它还有 go 语句，用于在一个独立的 goroutine 中运行代码。在这里，我想讨论一些不太常见的：defer、panic 和 recover。

​	defer 语句会将一个函数调用推迟到一个列表上。保存的函数调用列表会在包围的函数返回后执行。Defer 通常用于简化执行各种清理操作的函数。

​	例如，我们来看一个打开两个文件并将一个文件的内容复制到另一个文件的函数：

```go linenums="1"
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }

    written, err = io.Copy(dst, src)
    dst.Close()
    src.Close()
    return
}
```

​	这个函数是可行的，但是有一个 bug。如果调用 os.Create 失败，函数将在没有关闭源文件的情况下返回。可以很容易地通过在第二个 return 语句前放置 src.Close 的调用来解决这个问题，但如果函数更加复杂，问题可能就不那么容易被注意和解决了。通过引入 defer 语句，我们可以确保文件始终被关闭：

```go linenums="1"
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```

​	defer 语句使我们可以在打开文件后立即考虑关闭每个文件，确保无论函数中有多少个 return 语句，文件都将被关闭。defer语句允许我们在打开每个文件后立即考虑关闭它，保证无论函数中的返回语句有多少，文件都会被关闭。

​	defer 语句的行为是直接和可预测的。有三条简单的规则： 

（a）延迟执行的函数的参数在 defer 语句执行时计算。

在这个例子中，“i”表达式在 Println 调用被延迟时计算。延迟的调用会在函数返回后打印“0”。

```go linenums="1"
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```

（b）延迟函数调用在包围函数返回后按照后进先出的顺序执行。 

下面这个函数将输出“3210”：

```go linenums="1"
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
```

（c）延迟函数可以读取和修改返回值。 

​	在下面这个例子中，一个延迟函数在包围函数返回后将返回值 i 加 1。因此，这个函数返回 2：

```go linenums="1"
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```

​	这对于修改函数的错误返回值非常方便，我们将在稍后看到一个例子。

​	`panic` 是一个内置函数，它停止普通的控制流并开始 panic（恐慌）。当函数 F 调用 panic 时，F 的执行停止，F 中的任何延迟函数都按正常顺序执行，然后 F 返回给它的调用者。对于调用者来说，F 的行为就像一个 panic 调用。这个过程会一直沿着调用栈向上走，直到当前 goroutine 中的所有函数都返回，此时程序就会崩溃。Panic 可以通过直接调用 panic 来发起，也可以由运行时错误（例如数组越界）引起。

​	`recover` 是一个内置函数，它重新获得了正在 panic 的 goroutine 的控制权。Recover 只有在延迟函数中有用。在正常执行期间，调用 recover 会返回 nil，没有其他效果。如果当前 goroutine 正在 panic，调用 recover 将捕获传递给 panic 的值并恢复正常执行。

​	下面是一个演示 panic 和 defer 机制的示例程序：

```go linenums="1"
package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
```

​	函数 g 接受一个 int 类型的参数 i，如果 i 大于 3，就会 panic；否则它将用 i+1 作为参数调用自身。函数 f 延迟一个函数，该函数调用 recover 并打印恢复的值（如果不为 nil）。在阅读下面的内容之前，试着想象一下这个程序的输出会是什么。

​	该程序的输出将是：

```
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f.
```

​	如果我们从 f 函数中删除延迟函数，恐慌就无法恢复并到达 goroutine 的调用堆栈顶部，从而终止程序。修改后的程序将输出：

```
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
panic: 4

panic PC=0x2a9cd8
[stack trace omitted]
```

​	关于**panic**和**recover**的真实示例，请参见 Go 标准库中的 [json 包](https://go.dev/pkg/encoding/json/)。它使用一组递归函数对接口进行编码。如果在遍历值时发生错误，则调用 panic 来解开栈并返回适当的错误值（请参见 [encode.go](https://go.dev/src/pkg/encoding/json/encode.go) 中 encodeState 类型的 'error' 和 'marshal' 方法）。

​	Go 库的约定是，即使包内部使用 panic，其外部 API 仍然会呈现明确的错误返回值。

​	除了之前给出的 file.Close 示例之外，defer 的其他用途包括释放互斥锁：

```go linenums="1"
mu.Lock()
defer mu.Unlock()
```

打印一个页脚。

```go linenums="1"
printHeader()
defer printFooter()
```

等等。

​	总之，defer 语句（带或不带 panic 和 recover）提供了一种不同寻常且强大的控制流机制。它可用于模拟其他编程语言中的专用结构实现的一些特性。尝试一下。
