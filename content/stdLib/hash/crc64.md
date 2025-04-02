+++
title = "crc64"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/hash/crc64@go1.24.2](https://pkg.go.dev/hash/crc64@go1.24.2)

Package crc64 implements the 64-bit cyclic redundancy check, or CRC-64, checksum. See https://en.wikipedia.org/wiki/Cyclic_redundancy_check for information.

​	crc64 包实现了 64 位循环冗余校验或 CRC-64 校验和。有关信息，请参阅 https://en.wikipedia.org/wiki/Cyclic_redundancy_check。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/crc64/crc64.go;l=20)

``` go 
const (
	// The ISO polynomial, defined in ISO 3309 and used in HDLC.
	ISO = 0xD800000000000000

	// The ECMA polynomial, defined in ECMA 182.
	ECMA = 0xC96C5795D7870F42
)
```

Predefined polynomials.

​	预定义多项式。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/crc64/crc64.go;l=17)

``` go 
const Size = 8
```

The size of a CRC-64 checksum in bytes.

​	CRC-64 校验和的大小（以字节为单位）。

## 变量

This section is empty.

## 函数

### func Checksum 

``` go 
func Checksum(data []byte, tab *Table) uint64
```

Checksum returns the CRC-64 checksum of data using the polynomial represented by the Table.

​	Checksum 使用 Table 表示的多项式返回数据的 CRC-64 校验和。

### func New

```go
func New(tab *Table) hash.Hash64
```

New creates a new hash.Hash64 computing the CRC-64 checksum using the polynomial represented by the Table. Its Sum method will lay the value out in big-endian byte order. The returned Hash64 also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

​	New 使用 Table 表示的多项式创建一个新的 hash.Hash64 来计算 CRC-64 校验和。它的 Sum 方法将以大端字节顺序排列值。返回的 Hash64 还实现了 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler 来编组和取消编组哈希的内部状态。

### func Update

```go
func Update(crc uint64, tab *Table, p []byte) uint64
```

Update returns the result of adding the bytes in p to the crc.

​	Update 返回将 p 中的字节添加到 crc 的结果。

## 类型

### type Table

```go
type Table [256]uint64
```

Table is a 256-word table representing the polynomial for efficient processing.

​	Table 是一个 256 字的表，表示用于高效处理的多项式。

#### func MakeTable

```go
func MakeTable(poly uint64) *Table
```

MakeTable returns a Table constructed from the specified polynomial. The contents of this Table must not be modified.

​	MakeTable 返回由指定多项式构建的 Table。此 Table 的内容不得修改。