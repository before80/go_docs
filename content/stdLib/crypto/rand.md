+++
title = "rand"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# rand

https://pkg.go.dev/crypto/rand@go1.20.1



Package rand implements a cryptographically secure random number generator.












## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/rand/rand.go;l=20)

```
var Reader io.Reader
```

Reader is a global, shared instance of a cryptographically secure random number generator.

On Linux, FreeBSD, Dragonfly and Solaris, Reader uses getrandom(2) if available, /dev/urandom otherwise. On OpenBSD and macOS, Reader uses getentropy(2). On other Unix-like systems, Reader reads from /dev/urandom. On Windows systems, Reader uses the RtlGenRandom API. On Wasm, Reader uses the Web Crypto API.

## 函数

#### func Int 

```
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
```

Int returns a uniform random value in [0, max). It panics if max <= 0.

#### func Prime 

```
func Prime(rand io.Reader, bits int) (*big.Int, error)
```

Prime returns a number of the given bit length that is prime with high probability. Prime will return error for any error returned by rand.Read or if bits < 2.

#### func Read 

```
func Read(b []byte) (n int, err error)
```

Read is a helper function that calls Reader.Read using io.ReadFull. On return, n == len(b) if and only if err == nil.

##### Example

## 类型

This section is empty.