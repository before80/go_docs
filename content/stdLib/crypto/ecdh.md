+++
title = "ecdh"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# ecdh

https://pkg.go.dev/crypto/ecdh@go1.20.1



Package ecdh implements Elliptic Curve Diffie-Hellman over NIST curves and Curve25519.





  
  
  


  
  
  
  
  


  
  

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Curve 

```
type Curve interface {
	// GenerateKey generates a new PrivateKey from rand.
	GenerateKey(rand io.Reader) (*PrivateKey, error)

	// NewPrivateKey checks that key is valid and returns a PrivateKey.
	//
	// For NIST curves, this follows SEC 1, Version 2.0, Section 2.3.6, which
	// amounts to decoding the bytes as a fixed length big endian integer and
	// checking that the result is lower than the order of the curve. The zero
	// private key is also rejected, as the encoding of the corresponding public
	// key would be irregular.
	//
	// For X25519, this only checks the scalar length.
	NewPrivateKey(key []byte) (*PrivateKey, error)

	// NewPublicKey checks that key is valid and returns a PublicKey.
	//
	// For NIST curves, this decodes an uncompressed point according to SEC 1,
	// Version 2.0, Section 2.3.4. Compressed encodings and the point at
	// infinity are rejected.
	//
	// For X25519, this only checks the u-coordinate length. Adversarially
	// selected public keys can cause ECDH to return an error.
	NewPublicKey(key []byte) (*PublicKey, error)
	// contains filtered or unexported methods
}
```

#### func P256 

```
func P256() Curve
```

P256 returns a Curve which implements NIST P-256 (FIPS 186-3, section D.2.3), also known as secp256r1 or prime256v1.

Multiple invocations of this function will return the same value, which can be used for equality checks and switch statements.

#### func P384 

```
func P384() Curve
```

P384 returns a Curve which implements NIST P-384 (FIPS 186-3, section D.2.4), also known as secp384r1.

Multiple invocations of this function will return the same value, which can be used for equality checks and switch statements.

#### func P521 

```
func P521() Curve
```

P521 returns a Curve which implements NIST P-521 (FIPS 186-3, section D.2.5), also known as secp521r1.

Multiple invocations of this function will return the same value, which can be used for equality checks and switch statements.

#### func X25519 

```
func X25519() Curve
```

X25519 returns a Curve which implements the X25519 function over Curve25519 ([RFC 7748, Section 5](https://rfc-editor.org/rfc/rfc7748.html#section-5)).

Multiple invocations of this function will return the same value, so it can be used for equality checks and switch statements.

### type PrivateKey 

```
type PrivateKey struct {
	// contains filtered or unexported fields
}
```

PrivateKey is an ECDH private key, usually kept secret.

These keys can be parsed with [crypto/x509.ParsePKCS8PrivateKey](https://pkg.go.dev/crypto/x509#ParsePKCS8PrivateKey) and encoded with [crypto/x509.MarshalPKCS8PrivateKey](https://pkg.go.dev/crypto/x509#MarshalPKCS8PrivateKey). For NIST curves, they then need to be converted with [crypto/ecdsa.PrivateKey.ECDH](https://pkg.go.dev/crypto/ecdsa#PrivateKey.ECDH) after parsing.

#### (*PrivateKey) Bytes 

```
func (k *PrivateKey) Bytes() []byte
```

Bytes returns a copy of the encoding of the private key.

#### (*PrivateKey) Curve 

```
func (k *PrivateKey) Curve() Curve
```

#### (*PrivateKey) ECDH 

```
func (k *PrivateKey) ECDH(remote *PublicKey) ([]byte, error)
```

ECDH performs a ECDH exchange and returns the shared secret.

For NIST curves, this performs ECDH as specified in SEC 1, Version 2.0, Section 3.3.1, and returns the x-coordinate encoded according to SEC 1, Version 2.0, Section 2.3.5. The result is never the point at infinity.

For X25519, this performs ECDH as specified in [RFC 7748, Section 6.1](https://rfc-editor.org/rfc/rfc7748.html#section-6.1). If the result is the all-zero value, ECDH returns an error.

#### (*PrivateKey) Equal 

```
func (k *PrivateKey) Equal(x crypto.PrivateKey) bool
```

Equal returns whether x represents the same private key as k.

Note that there can be equivalent private keys with different encodings which would return false from this check but behave the same way as inputs to ECDH.

This check is performed in constant time as long as the key types and their curve match.

#### (*PrivateKey) Public 

```
func (k *PrivateKey) Public() crypto.PublicKey
```

Public implements the implicit interface of all standard library private keys. See the docs of crypto.PrivateKey.

#### (*PrivateKey) PublicKey 

```
func (k *PrivateKey) PublicKey() *PublicKey
```

### type PublicKey 

```
type PublicKey struct {
	// contains filtered or unexported fields
}
```

PublicKey is an ECDH public key, usually a peer's ECDH share sent over the wire.

These keys can be parsed with [crypto/x509.ParsePKIXPublicKey](https://pkg.go.dev/crypto/x509#ParsePKIXPublicKey) and encoded with [crypto/x509.MarshalPKIXPublicKey](https://pkg.go.dev/crypto/x509#MarshalPKIXPublicKey). For NIST curves, they then need to be converted with [crypto/ecdsa.PublicKey.ECDH](https://pkg.go.dev/crypto/ecdsa#PublicKey.ECDH) after parsing.

#### (*PublicKey) Bytes 

```
func (k *PublicKey) Bytes() []byte
```

Bytes returns a copy of the encoding of the public key.

#### (*PublicKey) Curve 

```
func (k *PublicKey) Curve() Curve
```

#### (*PublicKey) Equal 

```
func (k *PublicKey) Equal(x crypto.PublicKey) bool
```

Equal returns whether x represents the same public key as k.

Note that there can be equivalent public keys with different encodings which would return false from this check but behave the same way as inputs to ECDH.

This check is performed in constant time as long as the key types and their curve match.