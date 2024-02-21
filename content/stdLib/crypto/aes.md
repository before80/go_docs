+++
title = "aes"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/aes@go1.21.3](https://pkg.go.dev/crypto/aes@go1.21.3)

Package aes implements AES encryption (formerly Rijndael), as defined in U.S. Federal Information Processing Standards Publication 197.

​	aes 包实现了 AES 加密（以前称为 Rijndael），如美国联邦信息处理标准出版物 197 中所定义。

The AES operations in this package are not implemented using constant-time algorithms. An exception is when running on systems with enabled hardware support for AES that makes these operations constant-time. Examples include amd64 systems using AES-NI extensions and s390x systems using Message-Security-Assist extensions. On such systems, when the result of NewCipher is passed to cipher.NewGCM, the GHASH operation used by GCM is also constant-time.

​	此包中的 AES 操作并非使用恒定时间算法实现。例外情况是在启用了对 AES 的硬件支持的系统上运行时，这使得这些操作成为恒定时间。示例包括使用 AES-NI 扩展的 amd64 系统和使用消息安全辅助扩展的 s390x 系统。在这些系统上，当 NewCipher 的结果传递给 cipher.NewGCM 时，GCM 使用的 GHASH 操作也是恒定时间的。


## 常量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/aes/cipher.go;l=15)

``` go
const BlockSize = 16
```

The AES block size in bytes.

​	AES 块大小（以字节为单位）。

## 变量

This section is empty.

## 函数

### func NewCipher 

``` go
func NewCipher(key []byte) (cipher.Block, error)
```

NewCipher creates and returns a new cipher.Block. The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.

​	NewCipher 创建并返回一个新的 cipher.Block。key 参数应为 AES 密钥，16、14 或 32 字节，以选择 AES-128、AES-192 或 AES-256。

## 类型

### type KeySizeError 

``` go
type KeySizeError int
```

#### (KeySizeError) Error 

``` go
func (k KeySizeError) Error() string
```