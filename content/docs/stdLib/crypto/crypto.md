+++
title = "crypto"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# crypto

https://pkg.go.dev/crypto@go1.20.1



Package crypto collects common cryptographic constants.








  
  
  
  





## 常量 [¶](https://pkg.go.dev/crypto@go1.20.1#pkg-constants)

This section is empty.

## 变量

This section is empty.

## 函数

#### func [RegisterHash](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=145) [¶](https://pkg.go.dev/crypto@go1.20.1#RegisterHash)

```
func RegisterHash(h Hash, f func() hash.Hash)
```

RegisterHash registers a function that returns a new instance of the given hash function. This is intended to be called from the init function in packages that implement hash functions.

## 类型

### type [Decrypter](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=212) [¶](https://pkg.go.dev/crypto@go1.20.1#Decrypter)added in go1.5

```
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

### type [DecrypterOpts](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=223) [¶](https://pkg.go.dev/crypto@go1.20.1#DecrypterOpts)added in go1.5

```
type DecrypterOpts any
```

### type [Hash](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=16) [¶](https://pkg.go.dev/crypto@go1.20.1#Hash)

```
type Hash uint
```

Hash identifies a cryptographic hash function that is implemented in another package.

```
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

#### (Hash) [Available](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=138) [¶](https://pkg.go.dev/crypto@go1.20.1#Hash.Available)

```
func (h Hash) Available() bool
```

Available reports whether the given hash function is linked into the binary.

#### (Hash) [HashFunc](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=19) [¶](https://pkg.go.dev/crypto@go1.20.1#Hash.HashFunc)added in go1.4

```
func (h Hash) HashFunc() Hash
```

HashFunc simply returns the value of h so that Hash implements SignerOpts.

#### (Hash) [New](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=127) [¶](https://pkg.go.dev/crypto@go1.20.1#Hash.New)

```
func (h Hash) New() hash.Hash
```

New returns a new hash.Hash calculating the given hash function. New panics if the hash function is not linked into the binary.

#### (Hash) [Size](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=116) [¶](https://pkg.go.dev/crypto@go1.20.1#Hash.Size)

```
func (h Hash) Size() int
```

Size returns the length, in bytes, of a digest resulting from the given hash function. It doesn't require that the hash function in question be linked into the program.

#### (Hash) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=23) [¶](https://pkg.go.dev/crypto@go1.20.1#Hash.String)added in go1.15

```
func (h Hash) String() string
```

### type [PrivateKey](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=176) [¶](https://pkg.go.dev/crypto@go1.20.1#PrivateKey)

```
type PrivateKey any
```

PrivateKey represents a private key using an unspecified algorithm.

Although this type is an empty interface for backwards compatibility reasons, all private key types in the standard library implement the following interface

```
interface{
    Public() crypto.PublicKey
    Equal(x crypto.PrivateKey) bool
}
```

as well as purpose-specific interfaces such as Signer and Decrypter, which can be used for increased type safety within applications.

### type [PublicKey](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=162) [¶](https://pkg.go.dev/crypto@go1.20.1#PublicKey)added in go1.2

```
type PublicKey any
```

PublicKey represents a public key using an unspecified algorithm.

Although this type is an empty interface for backwards compatibility reasons, all public key types in the standard library implement the following interface

```
interface{
    Equal(x crypto.PublicKey) bool
}
```

which can be used for increased type safety within applications.

### type [Signer](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=180) [¶](https://pkg.go.dev/crypto@go1.20.1#Signer)added in go1.4

```
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

### type [SignerOpts](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/crypto.go;l=202) [¶](https://pkg.go.dev/crypto@go1.20.1#SignerOpts)added in go1.4

```
type SignerOpts interface {
	// HashFunc returns an identifier for the hash function used to produce
	// the message passed to Signer.Sign, or else zero to indicate that no
	// hashing was done.
	HashFunc() Hash
}
```

SignerOpts contains options for signing with a Signer.