+++
title = "hash"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# hash

https://pkg.go.dev/hash@go1.20.1



Package hash provides interfaces for hash functions.

##### Example
``` go linenums="1"
```







## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type [Hash](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/hash.go;l=26) 

``` go linenums="1"
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

Hash implementations in the standard library (e.g. hash/crc32 and crypto/sha256) implement the encoding.BinaryMarshaler and encoding.BinaryUnmarshaler interfaces. Marshaling a hash implementation allows its internal state to be saved and used for additional processing later, without having to re-write the data previously written to the hash. The hash state may contain portions of the input in its original form, which users are expected to handle for any possible security implications.

Compatibility: Any future changes to hash or crypto packages will endeavor to maintain compatibility with state encoded using previous versions. That is, any released versions of the packages should be able to decode data written with any previously released version, subject to issues such as security fixes. See the Go compatibility document for background: https://golang.org/doc/go1compat

### type [Hash32](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/hash.go;l=49) 

``` go linenums="1"
type Hash32 interface {
	Hash
	Sum32() uint32
}
```

Hash32 is the common interface implemented by all 32-bit hash functions.

### type [Hash64](https://cs.opensource.google/go/go/+/go1.20.1:src/hash/hash.go;l=55) 

``` go linenums="1"
type Hash64 interface {
	Hash
	Sum64() uint64
}
```

Hash64 is the common interface implemented by all 64-bit hash functions.