+++
title = "maphash"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/hash/maphash@go1.23.0](https://pkg.go.dev/hash/maphash@go1.23.0)

Package maphash provides hash functions on byte sequences. These hash functions are intended to be used to implement hash tables or other data structures that need to map arbitrary strings or byte sequences to a uniform distribution on unsigned 64-bit integers. Each different instance of a hash table or data structure should use its own Seed.

​	maphash 包提供字节序列的哈希函数。这些哈希函数旨在用于实现哈希表或其他需要将任意字符串或字节序列映射到无符号 64 位整数的均匀分布的数据结构。哈希表或数据结构的每个不同实例都应使用其自己的种子。

The hash functions are not cryptographically secure. (See crypto/sha256 and crypto/sha512 for cryptographic use.)

​	哈希函数不是加密安全的。（有关加密用途，请参阅 crypto/sha256 和 crypto/sha512。）

## Example 示例

```go
package main

import (
	"fmt"
	"hash/maphash"
)

func main() {
	// The zero Hash value is valid and ready to use; setting an
	// initial seed is not necessary.
	var h maphash.Hash

	// Add a string to the hash, and print the current hash value.
	h.WriteString("hello, ")
	fmt.Printf("%#x\n", h.Sum64())

	// Append additional data (in the form of a byte array).
	h.Write([]byte{'w', 'o', 'r', 'l', 'd'})
	fmt.Printf("%#x\n", h.Sum64())

	// Reset discards all data previously added to the Hash, without
	// changing its seed.
	h.Reset()

	// Use SetSeed to create a new Hash h2 which will behave
	// identically to h.
	var h2 maphash.Hash
	h2.SetSeed(h.Seed())

	h.WriteString("same")
	h2.WriteString("same")
	fmt.Printf("%#x == %#x\n", h.Sum64(), h2.Sum64())
}

Output:
```

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func Bytes <- go1.19

```go
func Bytes(seed Seed, b []byte) uint64
```

Bytes returns the hash of b with the given seed.

​	Bytes 使用给定的种子返回 b 的哈希值。

Bytes is equivalent to, but more convenient and efficient than:

​	Bytes 等同于以下内容，但更方便且更高效：

```go
var h Hash
h.SetSeed(seed)
h.Write(b)
return h.Sum64()
```

### func String <- go1.19

```go
func String(seed Seed, s string) uint64
```

String returns the hash of s with the given seed.

​	String 使用给定的种子返回 s 的哈希值。

String is equivalent to, but more convenient and efficient than:

​	String 等同于以下内容，但更方便且更高效：

```go
var h Hash
h.SetSeed(seed)
h.WriteString(s)
return h.Sum64()
```

## 类型

### type Hash 

```go
type Hash struct {
	// contains filtered or unexported fields
}
```

A Hash computes a seeded hash of a byte sequence.

​	Hash 计算字节序列的种子哈希值。

The zero Hash is a valid Hash ready to use. A zero Hash chooses a random seed for itself during the first call to a Reset, Write, Seed, or Sum64 method. For control over the seed, use SetSeed.

​	零 Hash 是一个有效的 Hash，随时可以使用。零 Hash 会在首次调用 Reset、Write、Seed 或 Sum64 方法时为自己选择一个随机种子。若要控制种子，请使用 SetSeed。

The computed hash values depend only on the initial seed and the sequence of bytes provided to the Hash object, not on the way in which the bytes are provided. For example, the three sequences

​	计算出的哈希值仅取决于初始种子和提供给 Hash 对象的字节序列，而与提供字节的方式无关。例如，以下三个序列

```
h.Write([]byte{'f','o','o'})
h.WriteByte('f'); h.WriteByte('o'); h.WriteByte('o')
h.WriteString("foo")
```

all have the same effect.

​	都具有相同的效果。

Hashes are intended to be collision-resistant, even for situations where an adversary controls the byte sequences being hashed.

​	哈希旨在具有抗冲突性，即使在对手控制要进行哈希处理的字节序列的情况下也是如此。

A Hash is not safe for concurrent use by multiple goroutines, but a Seed is. If multiple goroutines must compute the same seeded hash, each can declare its own Hash and call SetSeed with a common Seed.

​	哈希不适用于多个 goroutine 并发使用，但种子可以。如果多个 goroutine 必须计算相同的种子哈希，则每个 goroutine 都可以声明自己的哈希并使用公共种子调用 SetSeed。

#### (*Hash) BlockSize

```go
func (h *Hash) BlockSize() int
```

BlockSize returns h’s block size.

​	BlockSize 返回 h 的块大小。

#### (*Hash) Reset

```go
func (h *Hash) Reset()
```

Reset discards all bytes added to h. (The seed remains the same.)

​	Reset 丢弃添加到 h 的所有字节。（种子保持不变。）

#### (*Hash) Seed

```go
func (h *Hash) Seed() Seed
```

Seed returns h’s seed value.

​	Seed 返回 h 的种子值。

#### (*Hash) SetSeed

```go
func (h *Hash) SetSeed(seed Seed)
```

SetSeed sets h to use seed, which must have been returned by MakeSeed or by another Hash’s Seed method. Two Hash objects with the same seed behave identically. Two Hash objects with different seeds will very likely behave differently. Any bytes added to h before this call will be discarded.

​	SetSeed 将 h 设置为使用种子，该种子必须由 MakeSeed 或另一个 Hash 的 Seed 方法返回。具有相同种子的两个 Hash 对象的行为完全相同。具有不同种子的两个 Hash 对象的行为很可能不同。在此调用之前添加到 h 的任何字节都将被丢弃。

#### (*Hash) Size

```go
func (h *Hash) Size() int
```

Size returns h’s hash value size, 8 bytes.

​	Size 返回 h 的哈希值大小，为 8 个字节。

#### (*Hash) Sum

```go
func (h *Hash) Sum(b []byte) []byte
```

Sum appends the hash’s current 64-bit value to b. It exists for implementing hash.Hash. For direct calls, it is more efficient to use Sum64.

​	Sum 将哈希的当前 64 位值追加到 b。它用于实现 hash.Hash。对于直接调用，使用 Sum64 更有效。

#### (*Hash) Sum64

```go
func (h *Hash) Sum64() uint64
```

Sum64 returns h’s current 64-bit value, which depends on h’s seed and the sequence of bytes added to h since the last call to Reset or SetSeed.

​	Sum64 返回 h 的当前 64 位值，该值取决于 h 的种子和自上次调用 Reset 或 SetSeed 以来添加到 h 的字节序列。

All bits of the Sum64 result are close to uniformly and independently distributed, so it can be safely reduced by using bit masking, shifting, or modular arithmetic.

​	Sum64 结果的所有位都接近均匀且独立分布，因此可以使用位掩码、移位或模运算安全地对其进行缩减。

#### (*Hash) Write

```go
func (h *Hash) Write(b []byte) (int, error)
```

Write adds b to the sequence of bytes hashed by h. It always writes all of b and never fails; the count and error result are for implementing io.Writer.

​	Write 将 b 添加到 h 哈希的字节序列中。它始终写入所有 b 且永不失败；count 和 error 结果用于实现 io.Writer。

#### (*Hash) WriteByte 

```go
func (h *Hash) WriteByte(b byte) error
```

WriteByte adds b to the sequence of bytes hashed by h. It never fails; the error result is for implementing io.ByteWriter.

​	WriteByte 将 b 添加到 h 哈希的字节序列中。它绝不会失败；错误结果用于实现 io.ByteWriter。

#### (*Hash) WriteString

```go
func (h *Hash) WriteString(s string) (int, error)
```

WriteString adds the bytes of s to the sequence of bytes hashed by h. It always writes all of s and never fails; the count and error result are for implementing io.StringWriter.

​	WriteString 将 s 的字节添加到 h 哈希的字节序列中。它始终写入所有 s 且绝不会失败；count 和 error 结果用于实现 io.StringWriter。

### type Seed

```go
type Seed struct {
	// contains filtered or unexported fields
}
```

A Seed is a random value that selects the specific hash function computed by a Hash. If two Hashes use the same Seeds, they will compute the same hash values for any given input. If two Hashes use different Seeds, they are very likely to compute distinct hash values for any given input.

​	Seed 是一个随机值，用于选择 Hash 计算的特定哈希函数。如果两个 Hash 使用相同的 Seed，它们将为任何给定的输入计算相同的哈希值。如果两个 Hash 使用不同的 Seed，它们很可能会为任何给定的输入计算不同的哈希值。

A Seed must be initialized by calling MakeSeed. The zero seed is uninitialized and not valid for use with Hash’s SetSeed method.

​	必须通过调用 MakeSeed 来初始化 Seed。零种子未初始化，不能与 Hash 的 SetSeed 方法一起使用。

Each Seed value is local to a single process and cannot be serialized or otherwise recreated in a different process.

​	每个 Seed 值都属于单个进程，不能在其他进程中序列化或以其他方式重新创建。

#### func MakeSeed

```go
func MakeSeed() Seed
```

MakeSeed returns a new random seed.

​	MakeSeed 返回一个新的随机种子。