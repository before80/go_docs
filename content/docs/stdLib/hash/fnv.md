+++
title = "fnv"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# fnv

https://pkg.go.dev/hash/fnv@go1.20.1



Package fnv implements FNV-1 and FNV-1a, non-cryptographic hash functions created by Glenn Fowler, Landon Curt Noll, and Phong Vo. See https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function.

All the hash.Hash implementations returned by this package also implement encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [New128](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/fnv/fnv.go;l=71)  <- go1.9

``` go linenums="1"
func New128() hash.Hash
```

New128 returns a new 128-bit FNV-1 hash.Hash. Its Sum method will lay the value out in big-endian byte order.

#### func [New128a](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/fnv/fnv.go;l=80)  <- go1.9

``` go linenums="1"
func New128a() hash.Hash
```

New128a returns a new 128-bit FNV-1a hash.Hash. Its Sum method will lay the value out in big-endian byte order.

#### func [New32](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/fnv/fnv.go;l=43) 

``` go linenums="1"
func New32() hash.Hash32
```

New32 returns a new 32-bit FNV-1 hash.Hash. Its Sum method will lay the value out in big-endian byte order.

#### func [New32a](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/fnv/fnv.go;l=50) 

``` go linenums="1"
func New32a() hash.Hash32
```

New32a returns a new 32-bit FNV-1a hash.Hash. Its Sum method will lay the value out in big-endian byte order.

#### func [New64](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/fnv/fnv.go;l=57) 

``` go linenums="1"
func New64() hash.Hash64
```

New64 returns a new 64-bit FNV-1 hash.Hash. Its Sum method will lay the value out in big-endian byte order.

#### func [New64a](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/fnv/fnv.go;l=64) 

``` go linenums="1"
func New64a() hash.Hash64
```

New64a returns a new 64-bit FNV-1a hash.Hash. Its Sum method will lay the value out in big-endian byte order.

## 类型

This section is empty.