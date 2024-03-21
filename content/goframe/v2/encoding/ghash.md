+++
title = "ghash"
date = 2024-03-21T17:49:28+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/ghash

Package ghash provides some classic hash functions(uint32/uint64) in go.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func AP 

``` go
func AP(str []byte) uint32
```

AP implements the classic AP hash algorithm for 32 bits.

##### func AP64 

``` go
func AP64(str []byte) uint64
```

AP64 implements the classic AP hash algorithm for 64 bits.

##### func BKDR 

``` go
func BKDR(str []byte) uint32
```

BKDR implements the classic BKDR hash algorithm for 32 bits.

##### func BKDR64 

``` go
func BKDR64(str []byte) uint64
```

BKDR64 implements the classic BKDR hash algorithm for 64 bits.

##### func DJB 

``` go
func DJB(str []byte) uint32
```

DJB implements the classic DJB hash algorithm for 32 bits.

##### func DJB64 

``` go
func DJB64(str []byte) uint64
```

DJB64 implements the classic DJB hash algorithm for 64 bits.

##### func ELF 

``` go
func ELF(str []byte) uint32
```

ELF implements the classic ELF hash algorithm for 32 bits.

##### func ELF64 

``` go
func ELF64(str []byte) uint64
```

ELF64 implements the classic ELF hash algorithm for 64 bits.

##### func JS 

``` go
func JS(str []byte) uint32
```

JS implements the classic JS hash algorithm for 32 bits.

##### func JS64 

``` go
func JS64(str []byte) uint64
```

JS64 implements the classic JS hash algorithm for 64 bits.

##### func PJW 

``` go
func PJW(str []byte) uint32
```

PJW implements the classic PJW hash algorithm for 32 bits.

##### func PJW64 

``` go
func PJW64(str []byte) uint64
```

PJW64 implements the classic PJW hash algorithm for 64 bits.

##### func RS 

``` go
func RS(str []byte) uint32
```

RS implements the classic RS hash algorithm for 32 bits.

##### func RS64 

``` go
func RS64(str []byte) uint64
```

RS64 implements the classic RS hash algorithm for 64 bits.

##### func SDBM 

``` go
func SDBM(str []byte) uint32
```

SDBM implements the classic SDBM hash algorithm for 32 bits.

##### func SDBM64 

``` go
func SDBM64(str []byte) uint64
```

SDBM64 implements the classic SDBM hash algorithm for 64 bits.

### Types 

This section is empty.