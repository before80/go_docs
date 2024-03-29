+++
title = "gbase64"
date = 2024-03-21T17:48:52+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gbase64](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gbase64)

Package gbase64 provides useful API for BASE64 encoding/decoding algorithm.

​	软件包 gbase64 为 BASE64 编码/解码算法提供了有用的 API。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Decode

```go
func Decode(data []byte) ([]byte, error)
```

Decode decodes bytes with BASE64 algorithm.

​	Decode 使用 BASE64 算法解码字节。

#### func DecodeString

```go
func DecodeString(data string) ([]byte, error)
```

DecodeString decodes string with BASE64 algorithm.

​	DecodeString 使用 BASE64 算法解码字符串。

#### func DecodeToString

```go
func DecodeToString(data string) (string, error)
```

DecodeToString decodes string with BASE64 algorithm.

​	DecodeToString 使用 BASE64 算法解码字符串。

#### func Encode

```go
func Encode(src []byte) []byte
```

Encode encodes bytes with BASE64 algorithm.

​	Encode 使用 BASE64 算法对字节进行编码。

#### func EncodeFile

```go
func EncodeFile(path string) ([]byte, error)
```

EncodeFile encodes file content of `path` using BASE64 algorithms.

​	EncodeFile `path` 使用 BASE64 算法对文件内容进行编码。

#### func EncodeFileToString

```go
func EncodeFileToString(path string) (string, error)
```

EncodeFileToString encodes file content of `path` to string using BASE64 algorithms.

​	EncodeFileToString 使用 BASE64 算法对 `path` to 字符串的文件内容进行编码。

#### func EncodeString

```go
func EncodeString(src string) string
```

EncodeString encodes string with BASE64 algorithm.

​	EncodeString 使用 BASE64 算法对字符串进行编码。

#### func EncodeToString

```go
func EncodeToString(src []byte) string
```

EncodeToString encodes bytes to string with BASE64 algorithm.

​	EncodeToString 使用 BASE64 算法将字节编码为字符串。

#### func MustDecode

```go
func MustDecode(data []byte) []byte
```

MustDecode decodes bytes with BASE64 algorithm. It panics if any error occurs.

​	MustDecode 使用 BASE64 算法解码字节。如果发生任何错误，它会崩溃。

#### func MustDecodeString

```go
func MustDecodeString(data string) []byte
```

MustDecodeString decodes string with BASE64 algorithm. It panics if any error occurs.

​	MustDecodeString 使用 BASE64 算法解码字符串。如果发生任何错误，它会崩溃。

#### func MustDecodeToString

```go
func MustDecodeToString(data string) string
```

MustDecodeToString decodes string with BASE64 algorithm. It panics if any error occurs.

​	MustDecodeToString 使用 BASE64 算法解码字符串。如果发生任何错误，它会崩溃。

#### func MustEncodeFile

```go
func MustEncodeFile(path string) []byte
```

MustEncodeFile encodes file content of `path` using BASE64 algorithms. It panics if any error occurs.

​	MustEncodeFile `path` 使用 BASE64 算法对文件内容进行编码。如果发生任何错误，它会崩溃。

#### func MustEncodeFileToString

```go
func MustEncodeFileToString(path string) string
```

MustEncodeFileToString encodes file content of `path` to string using BASE64 algorithms. It panics if any error occurs.

​	MustEncodeFileToString 使用 BASE64 算法对 `path` to 字符串的文件内容进行编码。如果发生任何错误，它会崩溃。

## 类型

This section is empty.