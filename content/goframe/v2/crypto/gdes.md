+++
title = "gdes"
date = 2024-03-21T17:46:43+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gdes

### Overview 

Package gdes provides useful API for DES encryption/decryption algorithms.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/crypto/gdes/gdes.go#L19)

``` go
const (
	NOPADDING = iota
	PKCS5PADDING
)
```

### Variables 

This section is empty.

### Functions 

##### func DecryptCBC 

``` go
func DecryptCBC(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

DecryptCBC decrypts `cipherText` using CBC mode.

##### func DecryptCBCTriple 

``` go
func DecryptCBCTriple(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

DecryptCBCTriple decrypts `cipherText` using TripleDES and CBC mode.

##### func DecryptECB 

``` go
func DecryptECB(cipherText []byte, key []byte, padding int) ([]byte, error)
```

DecryptECB decrypts `cipherText` using ECB mode.

##### func DecryptECBTriple 

``` go
func DecryptECBTriple(cipherText []byte, key []byte, padding int) ([]byte, error)
```

DecryptECBTriple decrypts `cipherText` using TripleDES and ECB mode. The length of the `key` should be either 16 or 24 bytes.

##### func EncryptCBC 

``` go
func EncryptCBC(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

EncryptCBC encrypts `plainText` using CBC mode.

##### func EncryptCBCTriple 

``` go
func EncryptCBCTriple(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

EncryptCBCTriple encrypts `plainText` using TripleDES and CBC mode.

##### func EncryptECB 

``` go
func EncryptECB(plainText []byte, key []byte, padding int) ([]byte, error)
```

EncryptECB encrypts `plainText` using ECB mode.

##### func EncryptECBTriple 

``` go
func EncryptECBTriple(plainText []byte, key []byte, padding int) ([]byte, error)
```

EncryptECBTriple encrypts `plainText` using TripleDES and ECB mode. The length of the `key` should be either 16 or 24 bytes.

##### func Padding 

``` go
func Padding(text []byte, padding int) ([]byte, error)
```

##### func PaddingPKCS5 

``` go
func PaddingPKCS5(text []byte, blockSize int) []byte
```

##### func UnPadding 

``` go
func UnPadding(text []byte, padding int) ([]byte, error)
```

##### func UnPaddingPKCS5 

``` go
func UnPaddingPKCS5(text []byte) []byte
```

### Types 

This section is empty.