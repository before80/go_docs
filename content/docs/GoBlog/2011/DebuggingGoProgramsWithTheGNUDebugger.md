+++
title = "用GNU调试器调试 go 程序"
weight = 6
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Debugging Go programs with the GNU Debugger - 用GNU调试器调试 go 程序

https://go.dev/blog/debug-gdb

Andrew Gerrand
30 October 2011

Last year we [reported](https://blog.golang.org/2010/11/debugging-go-code-status-report.html) that Go’s [gc](https://go.dev/cmd/gc/)/[ld](https://go.dev/cmd/6l/) toolchain produces DWARFv3 debugging information that can be read by the GNU Debugger (GDB). Since then, work has continued steadily on improving support for debugging Go code with GDB. Among the improvements are the ability to inspect goroutines and to print native Go data types, including structs, slices, strings, maps, interfaces, and channels.

去年我们报道了Go的gc/ld工具链产生的DWARFv3调试信息可以被GNU调试器（GDB）读取。从那时起，我们一直在不断地改进对用GDB调试Go代码的支持。在这些改进中，包括检查goroutines和打印本地Go数据类型的能力，包括结构、片断、字符串、地图、接口和通道。

To learn more about Go and GDB, see the [Debugging with GDB](https://go.dev/doc/debugging_with_gdb.html) article.

要了解更多关于Go和GDB的信息，请看用GDB调试的文章。
