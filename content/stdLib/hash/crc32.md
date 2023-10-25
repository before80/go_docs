+++
title = "crc32"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/hash/crc32@go1.20.1



Package crc32 implements the 32-bit cyclic redundancy check, or CRC-32, checksum. See https://en.wikipedia.org/wiki/Cyclic_redundancy_check for information.

Polynomials are represented in LSB-first form also known as reversed representation.

See https://en.wikipedia.org/wiki/Mathematics_of_cyclic_redundancy_checks#Reversed_representations_and_reciprocal_polynomials for information.







## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/crc32/crc32.go;l=26)

``` go 
const (
	// IEEE is by far and away the most common CRC-32 polynomial.
	// Used by ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, ...
	IEEE = 0xedb88320

	// Castagnoli's polynomial, used in iSCSI.
	// Has better error detection characteristics than IEEE.
	// https://dx.doi.org/10.1109/26.231911
	Castagnoli = 0x82f63b78

	// Koopman's polynomial.
	// Also has better error detection characteristics than IEEE.
	// https://dx.doi.org/10.1109/DSN.2002.1028931
	Koopman = 0xeb31d82e
)
```

Predefined polynomials.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/crc32/crc32.go;l=23)

``` go 
const Size = 4
```

The size of a CRC-32 checksum in bytes.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/crc32/crc32.go;l=101)

``` go 
var IEEETable = simpleMakeTable(IEEE)
```

IEEETable is the table for the IEEE polynomial.

## 函数

#### func Checksum 

``` go 
func Checksum(data []byte, tab *Table) uint32
```

Checksum returns the CRC-32 checksum of data using the polynomial represented by the Table.

#### func ChecksumIEEE 

``` go 
func ChecksumIEEE(data []byte) uint32
```

ChecksumIEEE returns the CRC-32 checksum of data using the IEEE polynomial.

#### func New 

``` go 
func New(tab *Table) hash.Hash32
```

New creates a new hash.Hash32 computing the CRC-32 checksum using the polynomial represented by the Table. Its Sum method will lay the value out in big-endian byte order. The returned Hash32 also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

#### func NewIEEE 

``` go 
func NewIEEE() hash.Hash32
```

NewIEEE creates a new hash.Hash32 computing the CRC-32 checksum using the IEEE polynomial. Its Sum method will lay the value out in big-endian byte order. The returned Hash32 also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

#### func Update 

``` go 
func Update(crc uint32, tab *Table, p []byte) uint32
```

Update returns the result of adding the bytes in p to the crc.

## 类型

### type Table 

``` go 
type Table [256]uint32
```

Table is a 256-word table representing the polynomial for efficient processing.

#### func MakeTable 

``` go 
func MakeTable(poly uint32) *Table
```

MakeTable returns a Table constructed from the specified polynomial. The contents of this Table must not be modified.