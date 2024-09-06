+++
title = "rand"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/rand@go1.23.0](https://pkg.go.dev/crypto/rand@go1.23.0)

Package rand implements a cryptographically secure random number generator.

​	Package rand 实现了一个密码安全随机数生成器。


## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/rand/rand.go;l=20)

``` go
var Reader io.Reader
```

Reader is a global, shared instance of a cryptographically secure random number generator.

​	Reader 是一个密码安全随机数生成器的全局共享实例。

On Linux, FreeBSD, Dragonfly and Solaris, Reader uses getrandom(2) if available, /dev/urandom otherwise. On OpenBSD and macOS, Reader uses getentropy(2). On other Unix-like systems, Reader reads from /dev/urandom. On Windows systems, Reader uses the RtlGenRandom API. On Wasm, Reader uses the Web Crypto API.

​	在 Linux、FreeBSD、Dragonfly 和 Solaris 上，Reader 在可用时使用 getrandom(2)，否则使用 /dev/urandom。在 OpenBSD 和 macOS 上，Reader 使用 getentropy(2)。在其他类 Unix 系统上，Reader 从 /dev/urandom 读取。在 Windows 系统上，Reader 使用 RtlGenRandom API。在 Wasm 上，Reader 使用 Web Crypto API。

## 函数

### func Int 

``` go
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
```

​	Int 返回 [0, max) 中的均匀随机值。如果 max <= 0，则会引发 panic。

### func Prime

```go
func Prime(rand io.Reader, bits int) (*big.Int, error)
```

Prime returns a number of the given bit length that is prime with high probability. Prime will return error for any error returned by rand.Read or if bits < 2.

​	Prime 返回给定位长的数字，该数字很可能是素数。对于 rand.Read 返回的任何错误或 bits < 2，Prime 将返回错误。

### func Read

```go
func Read(b []byte) (n int, err error)
```

Read is a helper function that calls Reader.Read using io.ReadFull. On return, n == len(b) if and only if err == nil.

​	Read 是一个帮助函数，它使用 io.ReadFull 调用 Reader.Read。返回时，当且仅当 err == nil 时，n == len(b)。

#### Read Example

```go
package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

func main() {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(bytes.Equal(b, make([]byte, c)))

}
Output:

false
```



## 类型

This section is empty.