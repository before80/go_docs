+++
title = "rc4"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# rc4

https://pkg.go.dev/crypto/rc4@go1.20.1



Package rc4 implements RC4 encryption, as defined in Bruce Schneier's Applied Cryptography.

RC4 is cryptographically broken and should not be used for secure applications.




## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Cipher 

``` go
type Cipher struct {
	// contains filtered or unexported fields
}
```

A Cipher is an instance of RC4 using a particular key.

#### func NewCipher 

``` go
func NewCipher(key []byte) (*Cipher, error)
```

NewCipher creates and returns a new Cipher. The key argument should be the RC4 key, at least 1 byte and at most 256 bytes.

#### (*Cipher) Reset <- DEPRECATED

#### (*Cipher) XORKeyStream 

``` go
func (c *Cipher) XORKeyStream(dst, src []byte)
```

XORKeyStream sets dst to the result of XORing src with the key stream. Dst and src must overlap entirely or not at all.

### type KeySizeError 

``` go
type KeySizeError int
```

#### (KeySizeError) Error 

``` go
func (k KeySizeError) Error() string
```