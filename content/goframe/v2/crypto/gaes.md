+++
title = "gaes"
date = 2024-03-21T17:46:25+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gaes

Package gaes provides useful API for AES encryption/decryption algorithms.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/crypto/gaes/gaes.go#L20)

``` go
const (
	// IVDefaultValue is the default value for IV.
	IVDefaultValue = "I Love Go Frame!"
)
```

### Variables 

This section is empty.

### Functions 

##### func Decrypt 

``` go
func Decrypt(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

Decrypt is alias of DecryptCBC.

##### func DecryptCBC 

``` go
func DecryptCBC(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

DecryptCBC decrypts `cipherText` using CBC mode. Note that the key must be 16/24/32 bit length. The parameter `iv` initialization vector is unnecessary.

##### func DecryptCFB 

``` go
func DecryptCFB(cipherText []byte, key []byte, unPadding int, iv ...[]byte) ([]byte, error)
```

DecryptCFB decrypts `plainText` using CFB mode. Note that the key must be 16/24/32 bit length. The parameter `iv` initialization vector is unnecessary.

##### func Encrypt 

``` go
func Encrypt(plainText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

Encrypt is alias of EncryptCBC.

##### func EncryptCBC 

``` go
func EncryptCBC(plainText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

EncryptCBC encrypts `plainText` using CBC mode. Note that the key must be 16/24/32 bit length. The parameter `iv` initialization vector is unnecessary.

##### func EncryptCFB 

``` go
func EncryptCFB(plainText []byte, key []byte, padding *int, iv ...[]byte) ([]byte, error)
```

EncryptCFB encrypts `plainText` using CFB mode. Note that the key must be 16/24/32 bit length. The parameter `iv` initialization vector is unnecessary.

##### func PKCS5Padding 

``` go
func PKCS5Padding(src []byte, blockSize ...int) []byte
```

PKCS5Padding applies PKCS#5 padding to the source byte slice to match the given block size.

If the block size is not provided, it defaults to 8.

##### func PKCS5UnPadding 

``` go
func PKCS5UnPadding(src []byte, blockSize ...int) ([]byte, error)
```

PKCS5UnPadding removes PKCS#5 padding from the source byte slice based on the given block size.

If the block size is not provided, it defaults to 8.

##### func PKCS7Padding <-2.5.7

``` go
func PKCS7Padding(src []byte, blockSize int) []byte
```

PKCS7Padding applies PKCS#7 padding to the source byte slice to match the given block size.

##### func PKCS7UnPadding <-2.5.7

``` go
func PKCS7UnPadding(src []byte, blockSize int) ([]byte, error)
```

PKCS7UnPadding removes PKCS#7 padding from the source byte slice based on the given block size.

##### func ZeroPadding 

``` go
func ZeroPadding(cipherText []byte, blockSize int) ([]byte, int)
```

##### func ZeroUnPadding 

``` go
func ZeroUnPadding(plaintext []byte, unPadding int) []byte
```

### Types 

This section is empty.