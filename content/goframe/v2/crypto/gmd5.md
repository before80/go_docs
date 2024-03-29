+++
title = "gmd5"
date = 2024-03-21T17:46:49+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gmd5](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gmd5)

Package gmd5 provides useful API for MD5 encryption algorithms.

​	软件包 gmd5 为 MD5 加密算法提供了有用的 API。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Encrypt

```go
func Encrypt(data interface{}) (encrypt string, err error)
```

Encrypt encrypts any type of variable using MD5 algorithms. It uses gconv package to convert `v` to its bytes type.

​	Encrypt 使用 MD5 算法对任何类型的变量进行加密。它使用 gconv 包转换为 `v` 其字节类型。

#### func EncryptBytes

```go
func EncryptBytes(data []byte) (encrypt string, err error)
```

EncryptBytes encrypts `data` using MD5 algorithms.

​	EncryptBytes `data` 使用 MD5 算法进行加密。

#### func EncryptFile

```go
func EncryptFile(path string) (encrypt string, err error)
```

EncryptFile encrypts file content of `path` using MD5 algorithms.

​	EncryptFile `path` 使用 MD5 算法对文件内容进行加密。

#### func EncryptString

```go
func EncryptString(data string) (encrypt string, err error)
```

EncryptString encrypts string `data` using MD5 algorithms.

​	EncryptString 使用 MD5 算法对字符串 `data` 进行加密。

#### func MustEncrypt

```go
func MustEncrypt(data interface{}) string
```

MustEncrypt encrypts any type of variable using MD5 algorithms. It uses gconv package to convert `v` to its bytes type. It panics if any error occurs.

​	MustEncrypt 使用 MD5 算法加密任何类型的变量。它使用 gconv 包转换为 `v` 其字节类型。如果发生任何错误，它会崩溃。

#### func MustEncryptBytes

```go
func MustEncryptBytes(data []byte) string
```

MustEncryptBytes encrypts `data` using MD5 algorithms. It panics if any error occurs.

​	MustEncryptBytes `data` 使用 MD5 算法进行加密。如果发生任何错误，它会崩溃。

#### func MustEncryptFile

```go
func MustEncryptFile(path string) string
```

MustEncryptFile encrypts file content of `path` using MD5 algorithms. It panics if any error occurs.

​	MustEncryptFile `path` 使用 MD5 算法对文件内容进行加密。如果发生任何错误，它会崩溃。

#### func MustEncryptString

```go
func MustEncryptString(data string) string
```

MustEncryptString encrypts string `data` using MD5 algorithms. It panics if any error occurs.

​	MustEncryptString 使用 MD5 算法对字符串 `data` 进行加密。如果发生任何错误，它会崩溃。

## 类型

This section is empty.