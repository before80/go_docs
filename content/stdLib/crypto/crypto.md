+++
title = "crypto"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto@go1.23.0](https://pkg.go.dev/crypto@go1.23.0)

Package crypto collects common cryptographic constants.

​	crypto 包收集常见的加密常量。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func RegisterHash 

``` go
func RegisterHash(h Hash, f func() hash.Hash)
```

RegisterHash registers a function that returns a new instance of the given hash function. This is intended to be called from the init function in packages that implement hash functions.

​	RegisterHash 注册一个函数，该函数返回给定哈希函数的新实例。这旨在从实现哈希函数的包中的 init 函数中调用。

## 类型

### type Decrypter  <- go1.5

``` go
type Decrypter interface {
	// Public returns the public key corresponding to the opaque,
	// private key.
	Public() PublicKey

	// Decrypt decrypts msg. The opts argument should be appropriate for
	// the primitive used. See the documentation in each implementation for
	// details.
	Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}
```

Decrypter is an interface for an opaque private key that can be used for asymmetric decryption operations. An example would be an RSA key kept in a hardware module.

​	Decrypter 是一个不透明私钥的接口，可用于非对称解密操作。一个示例是保存在硬件模块中的 RSA 密钥。

### type DecrypterOpts  <- go1.5

``` go
type DecrypterOpts any
```

### type Hash 

``` go
type Hash uint
```

Hash identifies a cryptographic hash function that is implemented in another package.

​	Hash 标识在另一个包中实现的加密哈希函数。

``` go
const (
	MD4         Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                         // import crypto/md5
	SHA1                        // import crypto/sha1
	SHA224                      // import crypto/sha256
	SHA256                      // import crypto/sha256
	SHA384                      // import crypto/sha512
	SHA512                      // import crypto/sha512
	MD5SHA1                     // no implementation; MD5+SHA1 used for TLS RSA
	RIPEMD160                   // import golang.org/x/crypto/ripemd160
	SHA3_224                    // import golang.org/x/crypto/sha3
	SHA3_256                    // import golang.org/x/crypto/sha3
	SHA3_384                    // import golang.org/x/crypto/sha3
	SHA3_512                    // import golang.org/x/crypto/sha3
	SHA512_224                  // import crypto/sha512
	SHA512_256                  // import crypto/sha512
	BLAKE2s_256                 // import golang.org/x/crypto/blake2s
	BLAKE2b_256                 // import golang.org/x/crypto/blake2b
	BLAKE2b_384                 // import golang.org/x/crypto/blake2b
	BLAKE2b_512                 // import golang.org/x/crypto/blake2b

)
```

#### (Hash) Available 

``` go
func (h Hash) Available() bool
```

Available reports whether the given hash function is linked into the binary.

​	Available 报告给定的哈希函数是否链接到二进制文件中。

#### (Hash) HashFunc  <- go1.4

``` go
func (h Hash) HashFunc() Hash
```

HashFunc simply returns the value of h so that Hash implements SignerOpts.

​	HashFunc 仅返回 h 的值，以便 Hash 实现 SignerOpts。

#### (Hash) New 

``` go
func (h Hash) New() hash.Hash
```

New returns a new hash.Hash calculating the given hash function. New panics if the hash function is not linked into the binary.

​	New 返回一个新的 hash.Hash，用于计算给定的哈希函数。如果哈希函数未链接到二进制文件中，New 会引发 panic。

#### (Hash) Size 

``` go
func (h Hash) Size() int
```

Size returns the length, in bytes, of a digest resulting from the given hash function. It doesn't require that the hash function in question be linked into the program.

​	Size 返回由给定哈希函数产生的摘要的长度（以字节为单位）。它不要求将有问题的哈希函数链接到程序中。

#### (Hash) String  <- go1.15

``` go
func (h Hash) String() string
```

### type PrivateKey 

``` go
type PrivateKey any
```

PrivateKey represents a private key using an unspecified algorithm.

​	PrivateKey 使用未指定的算法表示私钥。

Although this type is an empty interface for backwards compatibility reasons, all private key types in the standard library implement the following interface

​	尽管出于向后兼容性的原因，此类型是一个空接口，但标准库中的所有私钥类型都实现了以下接口

```
interface{
    Public() crypto.PublicKey
    Equal(x crypto.PrivateKey) bool
}
```

as well as purpose-specific interfaces such as Signer and Decrypter, which can be used for increased type safety within applications.

​	以及特定于用途的接口，例如 Signer 和 Decrypter，可用于提高应用程序中的类型安全性。

### type PublicKey  <- go1.2

``` go
type PublicKey any
```

PublicKey represents a public key using an unspecified algorithm.

​	PublicKey 使用未指定的算法表示公钥。

Although this type is an empty interface for backwards compatibility reasons, all public key types in the standard library implement the following interface

​	尽管出于向后兼容性的原因，此类型是一个空接口，但标准库中的所有公钥类型都实现了以下接口

```
interface{
    Equal(x crypto.PublicKey) bool
}
```

which can be used for increased type safety within applications.

​	可用于提高应用程序中的类型安全性。

### type Signer  <- go1.4

``` go
type Signer interface {
	// Public returns the public key corresponding to the opaque,
	// private key.
	Public() PublicKey

	// Sign signs digest with the private key, possibly using entropy from
	// rand. For an RSA key, the resulting signature should be either a
	// PKCS #1 v1.5 or PSS signature (as indicated by opts). For an (EC)DSA
	// key, it should be a DER-serialised, ASN.1 signature structure.
	//
	// Hash implements the SignerOpts interface and, in most cases, one can
	// simply pass in the hash function used as opts. Sign may also attempt
	// to type assert opts to other types in order to obtain algorithm
	// specific values. See the documentation in each package for details.
	//
	// Note that when a signature of a hash of a larger message is needed,
	// the caller is responsible for hashing the larger message and passing
	// the hash (as digest) and the hash function (as opts) to Sign.
	Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}
```

Signer is an interface for an opaque private key that can be used for signing operations. For example, an RSA key kept in a hardware module.

​	Signer 是一个不透明私钥的接口，可用于签名操作。例如，保存在硬件模块中的 RSA 密钥。

### type SignerOpts  <- go1.4

``` go
type SignerOpts interface {
	// HashFunc returns an identifier for the hash function used to produce
	// the message passed to Signer.Sign, or else zero to indicate that no
	// hashing was done.
	HashFunc() Hash
}
```

SignerOpts contains options for signing with a Signer.

​	SignerOpts 包含使用 Signer 进行签名的选项。