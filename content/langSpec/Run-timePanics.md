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

​	执行错误，如尝试对数组进行越界索引，会触发一个**运行时panic**，相当于调用内置函数[panic](../Built-inFunctions#handling-panics-处理恐慌)，并传入实现定义的接口类型`runtime.Error`的值。该类型满足预声明的接口类型[error](../Errors)。表示不同运行时错误条件的确切错误值未指定。

```go 
package runtime

type Error interface {
	error
	// and perhaps other methods
}
```