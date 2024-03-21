+++
title = "gdebug"
date = 2024-03-21T17:48:17+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/debug/gdebug

Package gdebug contains facilities for programs to debug themselves while they are running.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func BinVersion 

``` go
func BinVersion() string
```

BinVersion returns the version of current running binary. It uses ghash.BKDRHash+BASE36 algorithm to calculate the unique version of the binary.

##### func BinVersionMd5 

``` go
func BinVersionMd5() string
```

BinVersionMd5 returns the version of current running binary. It uses MD5 algorithm to calculate the unique version of the binary.

##### func Caller 

``` go
func Caller(skip ...int) (function string, path string, line int)
```

Caller returns the function name and the absolute file path along with its line number of the caller.

##### func CallerDirectory 

``` go
func CallerDirectory() string
```

CallerDirectory returns the directory of the caller.

##### func CallerFileLine 

``` go
func CallerFileLine() string
```

CallerFileLine returns the file path along with the line number of the caller.

##### func CallerFileLineShort 

``` go
func CallerFileLineShort() string
```

CallerFileLineShort returns the file name along with the line number of the caller.

##### func CallerFilePath 

``` go
func CallerFilePath() string
```

CallerFilePath returns the file path of the caller.

##### func CallerFunction 

``` go
func CallerFunction() string
```

CallerFunction returns the function name of the caller.

##### func CallerPackage 

``` go
func CallerPackage() string
```

CallerPackage returns the package name of the caller.

##### func CallerWithFilter 

``` go
func CallerWithFilter(filters []string, skip ...int) (function string, path string, line int)
```

CallerWithFilter returns the function name and the absolute file path along with its line number of the caller.

The parameter `filters` is used to filter the path of the caller.

##### func FuncName 

``` go
func FuncName(f interface{}) string
```

FuncName returns the function name of given `f`.

##### func FuncPath 

``` go
func FuncPath(f interface{}) string
```

FuncPath returns the complete function path of given `f`.

##### func GoroutineId 

``` go
func GoroutineId() int
```

GoroutineId retrieves and returns the current goroutine id from stack information. Be very aware that, it is with low performance as it uses runtime.Stack function. It is commonly used for debugging purpose.

##### func PrintStack 

``` go
func PrintStack(skip ...int)
```

PrintStack prints to standard error the stack trace returned by runtime.Stack.

##### func Stack 

``` go
func Stack(skip ...int) string
```

Stack returns a formatted stack trace of the goroutine that calls it. It calls runtime.Stack with a large enough buffer to capture the entire trace.

##### func StackWithFilter 

``` go
func StackWithFilter(filters []string, skip ...int) string
```

StackWithFilter returns a formatted stack trace of the goroutine that calls it. It calls runtime.Stack with a large enough buffer to capture the entire trace.

The parameter `filter` is used to filter the path of the caller.

##### func StackWithFilters 

``` go
func StackWithFilters(filters []string, skip ...int) string
```

StackWithFilters returns a formatted stack trace of the goroutine that calls it. It calls runtime.Stack with a large enough buffer to capture the entire trace.

The parameter `filters` is a slice of strings, which are used to filter the path of the caller.

TODO Improve the performance using debug.Stack.

### Types 

This section is empty.