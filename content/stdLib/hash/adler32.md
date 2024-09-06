+++
title = "adler32"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/hash/adler32@go1.23.0](https://pkg.go.dev/hash/adler32@go1.23.0)

Package adler32 implements the Adler-32 checksum.

​	Package adler32 实现 Adler-32 校验和。

It is defined in [RFC 1950](https://rfc-editor.org/rfc/rfc1950.html):

​	它在 [RFC 1950](https://rfc-editor.org/rfc/rfc1950.html) 中定义：

```
Adler-32 is composed of two sums accumulated per byte: s1 is
the sum of all bytes, s2 is the sum of all s1 values. Both sums
are done modulo 65521. s1 is initialized to 1, s2 to zero.  The
Adler-32 checksum is stored as s2*65536 + s1 in most-
significant-byte first (network) order.
```



## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/adler32/adler32.go;l=31)

``` go 
const Size = 4
```

The size of an Adler-32 checksum in bytes.

​	Adler-32 校验和的大小（以字节为单位）。

## 变量

This section is empty.

## 函数

### func Checksum 

``` go 
func Checksum(data []byte) uint32
```

Checksum returns the Adler-32 checksum of data.

​	Checksum 返回数据的 Adler-32 校验和。

### func New

```go
func New() hash.Hash32
```

New returns a new hash.Hash32 computing the Adler-32 checksum. Its Sum method will lay the value out in big-endian byte order. The returned Hash32 also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

​	New 返回一个新的 hash.Hash32，用于计算 Adler-32 校验和。它的 Sum 方法将以大端字节顺序排列值。返回的 Hash32 还实现了 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler，用于编组和取消编组哈希的内部状态。

## 类型

This section is empty.