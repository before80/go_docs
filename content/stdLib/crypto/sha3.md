+++
title = "sha3"
date = 2025-04-01T13:15:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
> 原文：[https://pkg.go.dev/crypto/sha3@go1.24.2](https://pkg.go.dev/crypto/sha3@go1.24.2)

> 注意
>
> ​	从go1.24.0开始才可以使用该包。

## Overview

Package sha3 implements the SHA-3 hash algorithms and the SHAKE extendable output functions defined in FIPS 202.

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func Sum224 

```
func Sum224(data []byte) [28]byte
```

Sum224 returns the SHA3-224 hash of data.

### func Sum256 

```
func Sum256(data []byte) [32]byte
```

Sum256 returns the SHA3-256 hash of data.

### func Sum384 

```
func Sum384(data []byte) [48]byte
```

Sum384 returns the SHA3-384 hash of data.

### func Sum512 

```
func Sum512(data []byte) [64]byte
```

Sum512 returns the SHA3-512 hash of data.

### func SumSHAKE128 

```
func SumSHAKE128(data []byte, length int) []byte
```

SumSHAKE128 applies the SHAKE128 extendable output function to data and returns an output of the given length in bytes.

### func SumSHAKE256 

```
func SumSHAKE256(data []byte, length int) []byte
```

SumSHAKE256 applies the SHAKE256 extendable output function to data and returns an output of the given length in bytes.

## 类型

### type SHA3 

```
type SHA3 struct {
	// contains filtered or unexported fields
}
```

SHA3 is an instance of a SHA-3 hash. It implements [hash.Hash](https://pkg.go.dev/hash#Hash).

#### func New224 

```
func New224() *SHA3
```

New224 creates a new SHA3-224 hash.

#### func New256 

```
func New256() *SHA3
```

New256 creates a new SHA3-256 hash.

#### func New384 

```
func New384() *SHA3
```

New384 creates a new SHA3-384 hash.

#### func New512 

```
func New512() *SHA3
```

New512 creates a new SHA3-512 hash.

#### (*SHA3) AppendBinary 

```
func (s *SHA3) AppendBinary(p []byte) ([]byte, error)
```

AppendBinary implements [encoding.BinaryAppender](https://pkg.go.dev/encoding#BinaryAppender).

#### (*SHA3) BlockSize 

```
func (s *SHA3) BlockSize() int
```

BlockSize returns the hash's rate.

#### (*SHA3) MarshalBinary 

```
func (s *SHA3) MarshalBinary() ([]byte, error)
```

MarshalBinary implements [encoding.BinaryMarshaler](https://pkg.go.dev/encoding#BinaryMarshaler).

#### (*SHA3) Reset 

```
func (s *SHA3) Reset()
```

Reset resets the hash to its initial state.

#### (*SHA3) Size 

```
func (s *SHA3) Size() int
```

Size returns the number of bytes Sum will produce.

#### (*SHA3) Sum 

```
func (s *SHA3) Sum(b []byte) []byte
```

Sum appends the current hash to b and returns the resulting slice.

#### (*SHA3) UnmarshalBinary 

```
func (s *SHA3) UnmarshalBinary(data []byte) error
```

UnmarshalBinary implements [encoding.BinaryUnmarshaler](https://pkg.go.dev/encoding#BinaryUnmarshaler).

#### (*SHA3) Write 

```
func (s *SHA3) Write(p []byte) (n int, err error)
```

Write absorbs more data into the hash's state.

### type SHAKE 

```
type SHAKE struct {
	// contains filtered or unexported fields
}
```

SHAKE is an instance of a SHAKE extendable output function.

#### func NewCSHAKE128 

```
func NewCSHAKE128(N, S []byte) *SHAKE
```

NewCSHAKE128 creates a new cSHAKE128 XOF.

N is used to define functions based on cSHAKE, it can be empty when plain cSHAKE is desired. S is a customization byte string used for domain separation. When N and S are both empty, this is equivalent to NewSHAKE128.

#### func NewCSHAKE256 

```
func NewCSHAKE256(N, S []byte) *SHAKE
```

NewCSHAKE256 creates a new cSHAKE256 XOF.

N is used to define functions based on cSHAKE, it can be empty when plain cSHAKE is desired. S is a customization byte string used for domain separation. When N and S are both empty, this is equivalent to NewSHAKE256.

#### func NewSHAKE128 

```
func NewSHAKE128() *SHAKE
```

NewSHAKE128 creates a new SHAKE128 XOF.

#### func NewSHAKE256 

```
func NewSHAKE256() *SHAKE
```

NewSHAKE256 creates a new SHAKE256 XOF.

#### (*SHAKE) AppendBinary 

```
func (s *SHAKE) AppendBinary(p []byte) ([]byte, error)
```

AppendBinary implements [encoding.BinaryAppender](https://pkg.go.dev/encoding#BinaryAppender).

#### (*SHAKE) BlockSize 

```
func (s *SHAKE) BlockSize() int
```

BlockSize returns the rate of the XOF.

#### (*SHAKE) MarshalBinary 

```
func (s *SHAKE) MarshalBinary() ([]byte, error)
```

MarshalBinary implements [encoding.BinaryMarshaler](https://pkg.go.dev/encoding#BinaryMarshaler).

#### (*SHAKE) Read 

```
func (s *SHAKE) Read(p []byte) (n int, err error)
```

Read squeezes more output from the XOF.

Any call to Write after a call to Read will panic.

#### (*SHAKE) Reset 

```
func (s *SHAKE) Reset()
```

Reset resets the XOF to its initial state.

#### (*SHAKE) UnmarshalBinary 

```
func (s *SHAKE) UnmarshalBinary(data []byte) error
```

UnmarshalBinary implements [encoding.BinaryUnmarshaler](https://pkg.go.dev/encoding#BinaryUnmarshaler).

#### (*SHAKE) Write 

```
func (s *SHAKE) Write(p []byte) (n int, err error)
```

Write absorbs more data into the XOF's state.

It panics if any output has already been read.