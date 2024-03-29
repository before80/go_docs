+++
title = "gredis"
date = 2024-03-21T17:47:42+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/debug/gdebug](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/debug/gdebug)

Package gdebug contains facilities for programs to debug themselves while they are running.

​	软件包 gdebug 包含程序在运行时自行调试的工具。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func BinVersion

```go
func BinVersion() string
```

BinVersion returns the version of current running binary. It uses ghash.BKDRHash+BASE36 algorithm to calculate the unique version of the binary.

​	BinVersion 返回当前正在运行的二进制文件的版本。它使用 ghash。BKDRHash+BASE36 算法来计算二进制文件的唯一版本。

#### func BinVersionMd5

```go
func BinVersionMd5() string
```

BinVersionMd5 returns the version of current running binary. It uses MD5 algorithm to calculate the unique version of the binary.

​	BinVersionMd5 返回当前正在运行的二进制文件的版本。它使用 MD5 算法来计算二进制文件的唯一版本。

#### func Caller

```go
func Caller(skip ...int) (function string, path string, line int)
```

Caller returns the function name and the absolute file path along with its line number of the caller.

​	调用方返回函数名称和绝对文件路径以及调用方的行号。

#### func CallerDirectory

```go
func CallerDirectory() string
```

CallerDirectory returns the directory of the caller.

​	CallerDirectory 返回调用方的目录。

#### func CallerFileLine

```go
func CallerFileLine() string
```

CallerFileLine returns the file path along with the line number of the caller.

​	CallerFileLine 返回文件路径以及调用方的行号。

#### func CallerFileLineShort

```go
func CallerFileLineShort() string
```

CallerFileLineShort returns the file name along with the line number of the caller.

​	CallerFileLineShort 返回调用方的文件名和行号。

#### func CallerFilePath

```go
func CallerFilePath() string
```

CallerFilePath returns the file path of the caller.

​	CallerFilePath 返回调用方的文件路径。

#### func CallerFunction

```go
func CallerFunction() string
```

CallerFunction returns the function name of the caller.

​	CallerFunction 返回调用方的函数名称。

#### func CallerPackage

```go
func CallerPackage() string
```

CallerPackage returns the package name of the caller.

​	CallerPackage 返回调用方的包名称。

#### func CallerWithFilter

```go
func CallerWithFilter(filters []string, skip ...int) (function string, path string, line int)
```

CallerWithFilter returns the function name and the absolute file path along with its line number of the caller.

​	CallerWithFilter 返回函数名称和绝对文件路径以及调用方的行号。

The parameter `filters` is used to filter the path of the caller.

​	该参数 `filters` 用于筛选调用方的路径。

#### func FuncName

```go
func FuncName(f interface{}) string
```

FuncName returns the function name of given `f`.

​	FuncName 返回给定 `f` 的函数名称。

#### func FuncPath

```go
func FuncPath(f interface{}) string
```

FuncPath returns the complete function path of given `f`.

​	FuncPath 返回给定 `f` 的完整函数路径。

#### func GoroutineId

```go
func GoroutineId() int
```

GoroutineId retrieves and returns the current goroutine id from stack information. Be very aware that, it is with low performance as it uses runtime.Stack function. It is commonly used for debugging purpose.

​	GoroutineId 从堆栈信息中检索并返回当前 goroutine ID。请注意，它在使用运行时性能较低。堆栈功能。它通常用于调试目的。

#### func PrintStack

```go
func PrintStack(skip ...int)
```

PrintStack prints to standard error the stack trace returned by runtime.Stack.

​	PrintStack 将运行时返回的堆栈跟踪打印为标准错误。叠。

#### func Stack

```go
func Stack(skip ...int) string
```

Stack returns a formatted stack trace of the goroutine that calls it. It calls runtime.Stack with a large enough buffer to capture the entire trace.

​	Stack 返回调用它的 goroutine 的格式化堆栈跟踪。它调用运行时。使用足够大的缓冲区堆叠以捕获整个跟踪。

#### func StackWithFilter

```go
func StackWithFilter(filters []string, skip ...int) string
```

StackWithFilter returns a formatted stack trace of the goroutine that calls it. It calls runtime.Stack with a large enough buffer to capture the entire trace.

​	StackWithFilter 返回调用它的 goroutine 的格式化堆栈跟踪。它调用运行时。使用足够大的缓冲区堆叠以捕获整个跟踪。

The parameter `filter` is used to filter the path of the caller.

​	该参数 `filter` 用于筛选调用方的路径。

#### func StackWithFilters

```go
func StackWithFilters(filters []string, skip ...int) string
```

StackWithFilters returns a formatted stack trace of the goroutine that calls it. It calls runtime.Stack with a large enough buffer to capture the entire trace.

​	StackWithFilters 返回调用它的 goroutine 的格式化堆栈跟踪。它调用运行时。使用足够大的缓冲区堆叠以捕获整个跟踪。

The parameter `filters` is a slice of strings, which are used to filter the path of the caller.

​	该参数 `filters` 是字符串的切片，用于筛选调用方的路径。

TODO Improve the performance using debug.Stack.

​	TODO 使用调试提高性能。叠。

## 类型

This section is empty.