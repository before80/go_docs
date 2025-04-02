+++
title = "fnv"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/hash/fnv@go1.24.2](https://pkg.go.dev/hash/fnv@go1.24.2)

Package fnv implements FNV-1 and FNV-1a, non-cryptographic hash functions created by Glenn Fowler, Landon Curt Noll, and Phong Vo. See https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function.

​	包 fnv 实现 FNV-1 和 FNV-1a，由 Glenn Fowler、Landon Curt Noll 和 Phong Vo 创建的非加密哈希函数。请参阅 https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function。

All the hash.Hash implementations returned by this package also implement encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

​	此程序包返回的所有 hash.Hash 实现还实现 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler，以编组和取消编组哈希的内部状态。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func New128  <- go1.9

``` go 
func New128() hash.Hash
```

New128 returns a new 128-bit FNV-1 hash.Hash. Its Sum method will lay the value out in big-endian byte order.

​	New128 返回一个新的 128 位 FNV-1 hash.Hash。它的 Sum 方法将以大端字节顺序排列值。

### func New128a  <- go1.9

``` go 
func New128a() hash.Hash
```

New128a returns a new 128-bit FNV-1a hash.Hash. Its Sum method will lay the value out in big-endian byte order.

​	New128a 返回一个新的 128 位 FNV-1a hash.Hash。它的 Sum 方法将以大端字节顺序排列值。

### func New32

```go
func New32() hash.Hash32
```

New32 returns a new 32-bit FNV-1 hash.Hash. Its Sum method will lay the value out in big-endian byte order.

​	New32 返回一个新的 32 位 FNV-1 hash.Hash。它的 Sum 方法将以大端字节顺序排列值。

### func New32a

```go
func New32a() hash.Hash32
```

New32a returns a new 32-bit FNV-1a hash.Hash. Its Sum method will lay the value out in big-endian byte order.

​	New32a 返回一个新的 32 位 FNV-1a hash.Hash。它的 Sum 方法将以大端字节顺序排列值。

### func New64

```go
func New64() hash.Hash64
```

New64 returns a new 64-bit FNV-1 hash.Hash. Its Sum method will lay the value out in big-endian byte order.

​	New64 返回一个新的 64 位 FNV-1 hash.Hash。它的 Sum 方法将以大端字节顺序排列值。

### func New64a

```go
func New64a() hash.Hash64
```

New64a returns a new 64-bit FNV-1a hash.Hash. Its Sum method will lay the value out in big-endian byte order.

​	New64a 返回一个新的 64 位 FNV-1a hash.Hash。它的 Sum 方法将以大端字节顺序排列值。

## 类型

This section is empty.