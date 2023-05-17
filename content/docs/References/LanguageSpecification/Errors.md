+++
title = "错误"
date = 2023-05-17T09:59:21+08:00
weight = 16
description = ""
isCJKLanguage = true
draft = false
+++
## Errors 错误

> 原文：[https://go.dev/ref/spec#Errors](https://go.dev/ref/spec#Errors)

预先声明的`error`类型被定义为

```go linenums="1"
type error interface {
	Error() string
}
```

​	它是代表错误条件的常规接口，`nil`值代表没有错误。例如，可以定义一个从文件中读取数据的函数：

```go linenums="1"
func Read(f *File, b []byte) (n int, err error)
```