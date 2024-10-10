+++
title = "运行时恐慌"
date = 2023-05-17T09:59:21+08:00
weight = 17
description = ""
isCJKLanguage = true
type = "docs"
draft = false
+++
## Run-time panics 运行时恐慌

> 原文：[https://go.dev/ref/spec#Run-time_panics](https://go.dev/ref/spec#Run-time_panics)

Execution errors such as attempting to index an array out of bounds trigger a *run-time panic* equivalent to a call of the built-in function [`panic`](https://go.dev/ref/spec#Handling_panics) with a value of the implementation-defined interface type `runtime.Error`. That type satisfies the predeclared interface type [`error`](https://go.dev/ref/spec#Errors). The exact error values that represent distinct run-time error conditions are unspecified.

​	执行错误，如试图对一个数组进行超界索引，会触发运行时恐慌，它等同于带由实现所定义的接口类型 `runtime.Error` 的值来对内置函数 [panic](../Built-inFunctions#handling-panics-处理恐慌) 的调用。这个类型满足预先声明的接口类型[error](../Errors)。表示不同的运行时错误条件的确切的错误值是未指定的。

```go 
package runtime

type Error interface {
	error
	// and perhaps other methods
}
```