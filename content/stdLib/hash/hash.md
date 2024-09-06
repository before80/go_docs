+++
title = "hash"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/hash@go1.23.0](https://pkg.go.dev/hash@go1.23.0)

Package hash provides interfaces for hash functions.

 	hash 包提供了哈希函数的接口。

## Example (BinaryMarshaler)
``` go 
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding"
	"fmt"
	"log"
)

func main() {
	const (
		input1 = "The tunneling gopher digs downwards, "
		input2 = "unaware of what he will find."
	)

	first := sha256.New()
	first.Write([]byte(input1))

	marshaler, ok := first.(encoding.BinaryMarshaler)
	if !ok {
		log.Fatal("first does not implement encoding.BinaryMarshaler")
	}
	state, err := marshaler.MarshalBinary()
	if err != nil {
		log.Fatal("unable to marshal hash:", err)
	}

	second := sha256.New()

	unmarshaler, ok := second.(encoding.BinaryUnmarshaler)
	if !ok {
		log.Fatal("second does not implement encoding.BinaryUnmarshaler")
	}
	if err := unmarshaler.UnmarshalBinary(state); err != nil {
		log.Fatal("unable to unmarshal hash:", err)
	}

	first.Write([]byte(input2))
	second.Write([]byte(input2))

	fmt.Printf("%x\n", first.Sum(nil))
	fmt.Println(bytes.Equal(first.Sum(nil), second.Sum(nil)))
}

Output:

57d51a066f3a39942649cd9a76c77e97ceab246756ff3888659e6aa5a07f4a52
true
```

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Hash 

``` go 
type Hash interface {
	// Write (via the embedded io.Writer interface) adds more data to the running hash.
	// It never returns an error.
	io.Writer

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte

	// Reset resets the Hash to its initial state.
	Reset()

	// Size returns the number of bytes Sum will return.
	Size() int

	// BlockSize returns the hash's underlying block size.
	// The Write method must be able to accept any amount
	// of data, but it may operate more efficiently if all writes
	// are a multiple of the block size.
	BlockSize() int
}
```

Hash is the common interface implemented by all hash functions.

​	Hash 是所有哈希函数实现的通用接口。

Hash implementations in the standard library (e.g. hash/crc32 and crypto/sha256) implement the encoding.BinaryMarshaler and encoding.BinaryUnmarshaler interfaces. Marshaling a hash implementation allows its internal state to be saved and used for additional processing later, without having to re-write the data previously written to the hash. The hash state may contain portions of the input in its original form, which users are expected to handle for any possible security implications.

​	标准库中的哈希实现（例如 hash/crc32 和 crypto/sha256）实现了 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler 接口。对哈希实现进行编组允许保存其内部状态并稍后用于其他处理，而无需重新写入先前写入哈希的数据。哈希状态可能包含其原始形式的部分输入，用户应处理任何可能的安全影响。

Compatibility: Any future changes to hash or crypto packages will endeavor to maintain compatibility with state encoded using previous versions. That is, any released versions of the packages should be able to decode data written with any previously released version, subject to issues such as security fixes. See the Go compatibility document for background: https://golang.org/doc/go1compat

​	兼容性：对 hash 或 crypto 包的任何未来更改都将努力保持与使用以前版本编码的状态的兼容性。也就是说，任何已发布版本的包都应该能够解码使用任何以前发布的版本编写的的数据，但需视安全修复等问题而定。有关背景信息，请参阅 Go 兼容性文档：[https://golang.org/doc/go1compat](https://golang.org/doc/go1compat)

### type Hash32

```go
type Hash32 interface {
	Hash
	Sum32() uint32
}
```

Hash32 is the common interface implemented by all 32-bit hash functions.

​	Hash32 是所有 32 位哈希函数实现的通用接口。

### type Hash64

```go
type Hash64 interface {
	Hash
	Sum64() uint64
}
```

Hash64 is the common interface implemented by all 64-bit hash functions.

​	Hash64 是所有 64 位哈希函数实现的通用接口。