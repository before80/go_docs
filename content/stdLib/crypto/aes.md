+++
title = "aes"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/crypto/aes@go1.21.3

Package aes implements AES encryption (formerly Rijndael), as defined in U.S. Federal Information Processing Standards Publication 197.

The AES operations in this package are not implemented using constant-time algorithms. An exception is when running on systems with enabled hardware support for AES that makes these operations constant-time. Examples include amd64 systems using AES-NI extensions and s390x systems using Message-Security-Assist extensions. On such systems, when the result of NewCipher is passed to cipher.NewGCM, the GHASH operation used by GCM is also constant-time.


## 常量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/aes/cipher.go;l=15)

``` go
const BlockSize = 16
```

The AES block size in bytes.

## 变量

This section is empty.

## 函数

### func NewCipher 

``` go
func NewCipher(key []byte) (cipher.Block, error)
```

NewCipher creates and returns a new cipher.Block. The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.

## 类型

### type KeySizeError 

``` go
type KeySizeError int
```

#### (KeySizeError) Error 

``` go
func (k KeySizeError) Error() string
```