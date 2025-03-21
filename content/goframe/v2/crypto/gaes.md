+++
title = "gaes"
date = 2024-03-21T17:46:25+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gaes](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gaes)

Package gaes provides useful API for AES encryption/decryption algorithms.

​	软件包 gaes 为 AES 加密/解密算法提供了有用的 API。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/crypto/gaes/gaes.go#L20)

```go
const (
	// IVDefaultValue is the default value for IV.
	IVDefaultValue = "I Love Go Frame!"
)
```

## 变量

This section is empty.

## 函数

#### func Decrypt

```go
func Decrypt(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

Decrypt is alias of DecryptCBC.

​	Decrypt 是 DecryptCBC 的别名。

#### func DecryptCBC

```go
func DecryptCBC(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

DecryptCBC decrypts `cipherText` using CBC mode. Note that the key must be 16/24/32 bit length. The parameter `iv` initialization vector is unnecessary.

​	DecryptCBC `cipherText` 使用 CBC 模式解密。请注意，密钥长度必须为 16/24/32 位。参数 `iv` 初始化向量是不必要的。

#### func DecryptCFB

```go
func DecryptCFB(cipherText []byte, key []byte, unPadding int, iv ...[]byte) ([]byte, error)
```

DecryptCFB decrypts `plainText` using CFB mode. Note that the key must be 16/24/32 bit length. The parameter `iv` initialization vector is unnecessary.

​	DecryptCFB `plainText` 使用 CFB 模式进行解密。请注意，密钥长度必须为 16/24/32 位。参数 `iv` 初始化向量是不必要的。

#### func Encrypt

```go
func Encrypt(plainText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

Encrypt is alias of EncryptCBC.

​	Encrypt 是 EncryptCBC 的别名。

#### func EncryptCBC

```go
func EncryptCBC(plainText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

EncryptCBC encrypts `plainText` using CBC mode. Note that the key must be 16/24/32 bit length. The parameter `iv` initialization vector is unnecessary.

​	EncryptCBC `plainText` 使用 CBC 模式进行加密。请注意，密钥长度必须为 16/24/32 位。参数 `iv` 初始化向量是不必要的。

#### func EncryptCFB

```go
func EncryptCFB(plainText []byte, key []byte, padding *int, iv ...[]byte) ([]byte, error)
```

EncryptCFB encrypts `plainText` using CFB mode. Note that the key must be 16/24/32 bit length. The parameter `iv` initialization vector is unnecessary.

​	EncryptCFB `plainText` 使用 CFB 模式进行加密。请注意，密钥长度必须为 16/24/32 位。参数 `iv` 初始化向量是不必要的。

#### func PKCS5Padding

```go
func PKCS5Padding(src []byte, blockSize ...int) []byte
```

PKCS5Padding applies PKCS#5 padding to the source byte slice to match the given block size.

​	PKCS5Padding 将 PKCS#5 填充应用于源字节片以匹配给定的块大小。

If the block size is not provided, it defaults to 8.

​	如果未提供块大小，则默认为 8。

#### func PKCS5UnPadding

```go
func PKCS5UnPadding(src []byte, blockSize ...int) ([]byte, error)
```

PKCS5UnPadding removes PKCS#5 padding from the source byte slice based on the given block size.

​	PKCS5UnPadding 根据给定的块大小从源字节片中删除 PKCS#5 填充。

If the block size is not provided, it defaults to 8.

​	如果未提供块大小，则默认为 8。

#### func PKCS7Padding <-2.5.7

```go
func PKCS7Padding(src []byte, blockSize int) []byte
```

PKCS7Padding applies PKCS#7 padding to the source byte slice to match the given block size.

​	PKCS7Padding 将 PKCS#7 填充应用于源字节片以匹配给定的块大小。

#### func PKCS7UnPadding <-2.5.7

```go
func PKCS7UnPadding(src []byte, blockSize int) ([]byte, error)
```

PKCS7UnPadding removes PKCS#7 padding from the source byte slice based on the given block size.

​	PKCS7UnPadding 根据给定的块大小从源字节片中删除 PKCS#7 填充。

#### func ZeroPadding

```go
func ZeroPadding(cipherText []byte, blockSize int) ([]byte, int)
```

#### func ZeroUnPadding

```go
func ZeroUnPadding(plaintext []byte, unPadding int) []byte
```

## 类型

This section is empty.