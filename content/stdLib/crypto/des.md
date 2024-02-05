+++
title = "des"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/des@go1.21.3](https://pkg.go.dev/crypto/des@go1.21.3)

Package des implements the Data Encryption Standard (DES) and the Triple Data Encryption Algorithm (TDEA) as defined in U.S. Federal Information Processing Standards Publication 46-3.

DES is cryptographically broken and should not be used for secure applications.

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/des/cipher.go;l=15)

``` go
const BlockSize = 8
```

The DES block size in bytes.

## 变量

This section is empty.

## 函数

### func NewCipher 

``` go
func NewCipher(key []byte) (cipher.Block, error)
```

NewCipher creates and returns a new cipher.Block.

### func NewTripleDESCipher 

``` go
func NewTripleDESCipher(key []byte) (cipher.Block, error)
```

NewTripleDESCipher creates and returns a new cipher.Block.

##### NewTripleDESCipher Example

```go
package main

import (
	"crypto/des"
)

func main() {
	// NewTripleDESCipher can also be used when EDE2 is required by
	// duplicating the first 8 bytes of the 16-byte key.
	ede2Key := []byte("example key 1234")

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)

	_, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}

	// See crypto/cipher for how to use a cipher.Block for encryption and
	// decryption.
}

```



## 类型

### type KeySizeError 

``` go
type KeySizeError int
```

#### (KeySizeError) Error 

``` go
func (k KeySizeError) Error() string
```