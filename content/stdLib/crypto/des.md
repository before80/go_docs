+++
title = "des"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# des



Package des implements the Data Encryption Standard (DES) and the Triple Data Encryption Algorithm (TDEA) as defined in U.S. Federal Information Processing Standards Publication 46-3.

DES is cryptographically broken and should not be used for secure applications.













## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/des/cipher.go;l=15)

```
const BlockSize = 8
```

The DES block size in bytes.

## 变量

This section is empty.

## 函数

#### func NewCipher 

```
func NewCipher(key []byte) (cipher.Block, error)
```

NewCipher creates and returns a new cipher.Block.

#### func NewTripleDESCipher 

```
func NewTripleDESCipher(key []byte) (cipher.Block, error)
```

NewTripleDESCipher creates and returns a new cipher.Block.

##### Example

## 类型

### type KeySizeError 

```
type KeySizeError int
```

#### (KeySizeError) Error 

```
func (k KeySizeError) Error() string
```