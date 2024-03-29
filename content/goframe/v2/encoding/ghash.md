+++
title = "ghash"
date = 2024-03-21T17:49:28+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/ghash](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/ghash)

Package ghash provides some classic hash functions(uint32/uint64) in go.

​	软件包 ghash 在 go 中提供了一些经典的哈希函数（uint32/uint64）。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func AP

```go
func AP(str []byte) uint32
```

AP implements the classic AP hash algorithm for 32 bits.

​	AP 实现了 32 位的经典 AP 哈希算法。

#### func AP64

```go
func AP64(str []byte) uint64
```

AP64 implements the classic AP hash algorithm for 64 bits.

​	AP64 实现了 64 位的经典 AP 哈希算法。

#### func BKDR

```go
func BKDR(str []byte) uint32
```

BKDR implements the classic BKDR hash algorithm for 32 bits.

​	BKDR 实现了 32 位的经典 BKDR 哈希算法。

#### func BKDR64

```go
func BKDR64(str []byte) uint64
```

BKDR64 implements the classic BKDR hash algorithm for 64 bits.

​	BKDR64 实现了 64 位的经典 BKDR 哈希算法。

#### func DJB

```go
func DJB(str []byte) uint32
```

DJB implements the classic DJB hash algorithm for 32 bits.

​	DJB 实现了 32 位的经典 DJB 哈希算法。

#### func DJB64

```go
func DJB64(str []byte) uint64
```

DJB64 implements the classic DJB hash algorithm for 64 bits.

​	DJB64 实现了 64 位的经典 DJB 哈希算法。

#### func ELF

```go
func ELF(str []byte) uint32
```

ELF implements the classic ELF hash algorithm for 32 bits.

​	ELF 实现了 32 位的经典 ELF 哈希算法。

#### func ELF64

```go
func ELF64(str []byte) uint64
```

ELF64 implements the classic ELF hash algorithm for 64 bits.

​	ELF64 实现了 64 位的经典 ELF 哈希算法。

#### func JS

```go
func JS(str []byte) uint32
```

JS implements the classic JS hash algorithm for 32 bits.

​	JS 实现了 32 位的经典 JS 哈希算法。

#### func JS64

```go
func JS64(str []byte) uint64
```

JS64 implements the classic JS hash algorithm for 64 bits.

​	JS64 实现了 64 位的经典 JS 哈希算法。

#### func PJW

```go
func PJW(str []byte) uint32
```

PJW implements the classic PJW hash algorithm for 32 bits.

​	PJW 实现了 32 位的经典 PJW 哈希算法。

#### func PJW64

```go
func PJW64(str []byte) uint64
```

PJW64 implements the classic PJW hash algorithm for 64 bits.
PJW64 实现了 64 位的经典 PJW 哈希算法。

#### func RS

```go
func RS(str []byte) uint32
```

RS implements the classic RS hash algorithm for 32 bits.
RS 实现了 32 位的经典 RS 哈希算法。

#### func RS64

```go
func RS64(str []byte) uint64
```

RS64 implements the classic RS hash algorithm for 64 bits.
RS64 实现了 64 位的经典 RS 哈希算法。

#### func SDBM

```go
func SDBM(str []byte) uint32
```

SDBM implements the classic SDBM hash algorithm for 32 bits.
SDBM 实现了 32 位的经典 SDBM 哈希算法。

#### func SDBM64

```go
func SDBM64(str []byte) uint64
```

SDBM64 implements the classic SDBM hash algorithm for 64 bits.
SDBM64 实现了 64 位的经典 SDBM 哈希算法。

## 类型

This section is empty.