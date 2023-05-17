+++
title = "go 命令首页表格"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Command Documentation - 命令文档

> 原文：[https://go.dev/doc/cmd](https://go.dev/doc/cmd)

​	有一套程序用于构建和处理 Go 源代码。这套程序通常由[go](../go)程序调用，而不是直接运行。

​	运行这些程序最常见的方式是作为 `go` 程序的一个子命令，例如 `go fmt`。以这种方式运行，该命令对完整的 Go 源代码包进行操作，`go` 使用适合于包级别处理的参数调用底层二进制文件。

​	这些程序也可以作为独立的二进制文件运行，带有未做修改的参数，使用`go tool`子命令，如`go tool cgo`。对于大多数命令来说，主要用于调试。有些命令，如`pprof`，只能通过`go tool`子命令来访问。

​	最后，`fmt`和`godoc`命令被安装成常规的二进制文件，称为`gofmt`和`godoc`，因为它们经常被引用。

​	点击链接以获得更多的文档、调用方法和使用细节。

| Name             | Synopsis 简述                                                |
| ---------------- | ------------------------------------------------------------ |
| [go](./go)       | `go` 程序管理 Go 源代码并运行这里列出的其他命令。有关使用细节，请参见该命令的文档。 |
| [cgo](./cgo)     | `cgo` 可以创建调用 C 代码的 Go 包。                          |
| [cover](./cover) | `cover` 是一个用于创建和分析由 "`go test -coverprofile`" 生成的覆盖率配置文件的程序。 |
| [fix](./fix)     | `fix`找到使用语言和库的旧特性的 Go 程序，并使用新特性重写它们。 |
| [fmt](./gofmt)   | `fmt` 格式化 Go 包，它也可以作为独立的 [gofmt](./gofmt) 命令使用，具有更多通用选项。 |
| [godoc](./godoc) | `godoc` 提取并生成 Go 包的文档。                             |
| [vet](./vet)     | `vet` 检查 Go 源代码并报告可疑的结构，例如实参与格式字符串不一致的 `Printf` 调用。 |

This is an abridged list. See the [full command reference](https://go.dev/cmd/) for documentation of the compilers and more.

​	这是一个简略的列表。请参阅完整的命令参考，了解编译器的文档和更多。