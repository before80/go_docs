+++
title = "Go 1.22 发布！"
date = 2024-02-22T20:31:34+08:00
weight = 990
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/go1.22](https://go.dev/blog/go1.22)

# Go 1.22 is released! - Go 1.22 发布！

Eli Bendersky, on behalf of the Go team

Eli Bendersky，代表 Go 团队

6 February 2024

2024 年 2 月 6 日

Today the Go team is thrilled to release Go 1.22, which you can get by visiting the [download page](https://go.dev/dl/).

​	今天，Go 团队很高兴发布 Go 1.22，您可以访问下载页面获取该版本。

Go 1.22 comes with several important new features and improvements. Here are some of the notable changes; for the full list, refer to the [release notes](https://go.dev/doc/go1.22).

​	Go 1.22 带来了多项重要的新功能和改进。以下是一些值得注意的更改；有关完整列表，请参阅发行说明。

## Language changes 语言更改

The long-standing “for” loop gotcha with accidental sharing of loop variables between iterations is now resolved. Starting with Go 1.22, the following code will print “a”, “b”, and “c” in some order:

​	长期存在的“for”循环问题，即在迭代之间意外共享循环变量，现已解决。从 Go 1.22 开始，以下代码将按某种顺序打印“a”、“b”和“c”：

```go
func main() {
    done := make(chan bool)

    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }

    // wait for all goroutines to complete before exiting
    for _ = range values {
        <-done
    }
}
```

For more information about this change and the tooling that helps keep code from breaking accidentally, see the earlier [loop variable blog post](https://go.dev/blog/loopvar-preview).	

​	有关此更改以及有助于防止代码意外中断的工具的更多信息，请参阅较早的循环变量博客文章。

The second language change is support for ranging over integers:

​	第二个语言更改是对整数范围的支持：

```go
package main

import "fmt"

func main() {
    for i := range 10 {
        fmt.Println(10 - i)
    }
    fmt.Println("go1.22 has lift-off!")
}
```

The values of `i` in this countdown program go from 0 to 9, inclusive. For more details, please refer to [the spec](https://go.dev/ref/spec#For_range).

​	此倒计时程序中 `i` 的值从 0 到 9（包括 0 和 9）。有关更多详细信息，请参阅规范。

## Improved performance 性能改进

Memory optimization in the Go runtime improves CPU performance by 1-3%, while also reducing the memory overhead of most Go programs by around 1%.

​	Go 运行时中的内存优化将 CPU 性能提高了 1-3%，同时还将大多数 Go 程序的内存开销减少了约 1%。

In Go 1.21, [we shipped](https://go.dev/blog/pgo) profile-guided optimization (PGO) for the Go compiler and this functionality continues to improve. One of the optimizations added in 1.22 is improved devirtualization, allowing static dispatch of more interface method calls. Most programs will see improvements between 2-14% with PGO enabled.

​	在 Go 1.21 中，我们为 Go 编译器提供了概要指导优化 (PGO)，此功能还在不断改进。1.22 中添加的一项优化是改进的去虚拟化，允许静态调度更多接口方法调用。启用 PGO 后，大多数程序的性能将提高 2-14%。

## Standard library additions 标准库新增内容

- A new [math/rand/v2](https://go.dev/pkg/math/rand/v2) package provides a cleaner, more consistent API and uses higher-quality, faster pseudo-random generation algorithms. See [the proposal](https://go.dev/issue/61716) for additional details.

- 新的 math/rand/v2 包提供了一个更简洁、更一致的 API，并使用更高质量、更快的伪随机生成算法。有关其他详细信息，请参阅提案。

- The patterns used by [net/http.ServeMux](https://go.dev/pkg/net/http#ServeMux) now accept methods and wildcards.

- net/http.ServeMux 使用的模式现在接受方法和通配符。

  For example, the router accepts a pattern like `GET /task/{id}/`, which matches only `GET` requests and captures the value of the `{id}` segment in a map that can be accessed through [Request](https://go.dev/pkg/net/http#Request) values.

  例如，路由器接受 `GET /task/{id}/` 这样的模式，它只匹配 `GET` 请求，并将 `{id}` 段的值捕获到可以通过请求值访问的映射中。

- A new `Null[T]` type in [database/sql](https://go.dev/pkg/database/sql) provides a way to scan nullable columns.

- database/sql 中的新 `Null[T]` 类型提供了一种扫描可空列的方法。

- A `Concat` function was added in package [slices](https://go.dev/pkg/slices), to concatenate multiple slices of any type.

- 在 slices 包中添加了一个 `Concat` 函数，用于连接任何类型的多个切片。

------

Thanks to everyone who contributed to this release by writing code and documentation, filing bugs, sharing feedback, and testing the release candidates. Your efforts helped to ensure that Go 1.22 is as stable as possible. As always, if you notice any problems, please [file an issue](https://go.dev/issue/new).

​	感谢所有通过编写代码和文档、提交错误报告、分享反馈和测试候选版本为此次发布做出贡献的人员。你们的努力帮助确保 Go 1.22 尽可能稳定。一如既往，如果您发现任何问题，请提交问题。

Enjoy Go 1.22!

​	尽情享受 Go 1.22！