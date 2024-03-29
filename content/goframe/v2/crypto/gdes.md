+++
title = "gdes"
date = 2024-03-21T17:46:43+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gdes](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gdes)

### Overview 概述

Package gdes provides useful API for DES encryption/decryption algorithms.

​	软件包 gdes 为 DES 加密/解密算法提供了有用的 API。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/crypto/gdes/gdes.go#L19)

```go
const (
	NOPADDING = iota
	PKCS5PADDING
)
```

## 变量

This section is empty.

## 函数

#### func DecryptCBC

```go
func DecryptCBC(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

DecryptCBC decrypts `cipherText` using CBC mode.

​	DecryptCBC `cipherText` 使用 CBC 模式解密。

#### func DecryptCBCTriple

```go
func DecryptCBCTriple(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

DecryptCBCTriple decrypts `cipherText` using TripleDES and CBC mode.

​	DecryptCBCTriple `cipherText` 使用 TripleDES 和 CBC 模式进行解密。

#### func DecryptECB

```go
func DecryptECB(cipherText []byte, key []byte, padding int) ([]byte, error)
```

DecryptECB decrypts `cipherText` using ECB mode.

​	DecryptECB `cipherText` 使用 ECB 模式解密。

#### func DecryptECBTriple

```go
func DecryptECBTriple(cipherText []byte, key []byte, padding int) ([]byte, error)
```

DecryptECBTriple decrypts `cipherText` using TripleDES and ECB mode. The length of the `key` should be either 16 or 24 bytes.

​	DecryptECBTriple `cipherText` 使用 TripleDES 和 ECB 模式进行解密。的 `key` 长度应为 16 或 24 字节。

#### func EncryptCBC

```go
func EncryptCBC(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

EncryptCBC encrypts `plainText` using CBC mode.

​	EncryptCBC `plainText` 使用 CBC 模式进行加密。

#### func EncryptCBCTriple

```go
func EncryptCBCTriple(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

EncryptCBCTriple encrypts `plainText` using TripleDES and CBC mode.

​	EncryptCBCTriple `plainText` 使用 TripleDES 和 CBC 模式进行加密。

#### func EncryptECB

```go
func EncryptECB(plainText []byte, key []byte, padding int) ([]byte, error)
```

EncryptECB encrypts `plainText` using ECB mode.

​	EncryptECB `plainText` 使用 ECB 模式进行加密。

#### func EncryptECBTriple

```go
func EncryptECBTriple(plainText []byte, key []byte, padding int) ([]byte, error)
```

EncryptECBTriple encrypts `plainText` using TripleDES and ECB mode. The length of the `key` should be either 16 or 24 bytes.

​	EncryptECBTriple `plainText` 使用 TripleDES 和 ECB 模式进行加密。的 `key` 长度应为 16 或 24 字节。

#### func Padding

```go
func Padding(text []byte, padding int) ([]byte, error)
```

#### func PaddingPKCS5

```go
func PaddingPKCS5(text []byte, blockSize int) []byte
```

#### func UnPadding

```go
func UnPadding(text []byte, padding int) ([]byte, error)
```

#### func UnPaddingPKCS5

```go
func UnPaddingPKCS5(text []byte) []byte
```

## 类型

This section is empty.