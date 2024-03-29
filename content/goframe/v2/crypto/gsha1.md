+++
title = "gsha1"
date = 2024-03-21T17:47:08+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gsha1](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gsha1)

Package gsha1 provides useful API for SHA1 encryption algorithms.

​	软件包 gsha1 为 SHA1 加密算法提供了有用的 API。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Encrypt

```go
func Encrypt(v interface{}) string
```

Encrypt encrypts any type of variable using SHA1 algorithms. It uses package gconv to convert `v` to its bytes type.

​	Encrypt 使用 SHA1 算法对任何类型的变量进行加密。它使用包 gconv 转换为 `v` 其字节类型。

#### func EncryptFile

```go
func EncryptFile(path string) (encrypt string, err error)
```

EncryptFile encrypts file content of `path` using SHA1 algorithms.

​	EncryptFile `path` 使用 SHA1 算法对文件内容进行加密。

#### func MustEncryptFile

```go
func MustEncryptFile(path string) string
```

MustEncryptFile encrypts file content of `path` using SHA1 algorithms. It panics if any error occurs.

​	MustEncryptFile `path` 使用 SHA1 算法对文件内容进行加密。如果发生任何错误，它会崩溃。

## 类型

This section is empty.