+++
title = "crc64"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/hash/crc64@go1.20.1



Package crc64 implements the 64-bit cyclic redundancy check, or CRC-64, checksum. See https://en.wikipedia.org/wiki/Cyclic_redundancy_check for information.



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

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/crc64/crc64.go;l=17)

``` go 
const Size = 8
```

The size of a CRC-64 checksum in bytes.

## 变量

This section is empty.

## 函数

#### func Checksum 

``` go 
func Checksum(data []byte, tab *Table) uint64
```

Checksum returns the CRC-64 checksum of data using the polynomial represented by the Table.

#### func New 

``` go 
func New(tab *Table) hash.Hash64
```

New creates a new hash.Hash64 computing the CRC-64 checksum using the polynomial represented by the Table. Its Sum method will lay the value out in big-endian byte order. The returned Hash64 also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

#### func Update 

``` go 
func Update(crc uint64, tab *Table, p []byte) uint64
```

Update returns the result of adding the bytes in p to the crc.

## 类型

### type Table 

``` go 
type Table [256]uint64
```

Table is a 256-word table representing the polynomial for efficient processing.

#### func MakeTable 

``` go 
func MakeTable(poly uint64) *Table
```

MakeTable returns a Table constructed from the specified polynomial. The contents of this Table must not be modified.