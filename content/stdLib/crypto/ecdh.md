+++
title = "ecdh"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/ecdh@go1.24.2](https://pkg.go.dev/crypto/ecdh@go1.24.2)

Package ecdh implements Elliptic Curve Diffie-Hellman over NIST curves and Curve25519. 

​	 ecdh 包实现了 NIST 曲线和 Curve25519 上的椭圆曲线 Diffie-Hellman。


## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Curve 

``` go
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

``` go
func P256() Curve
```

P256 returns a Curve which implements NIST P-256 (FIPS 186-3, section D.2.3), also known as secp256r1 or prime256v1.

​	P256 返回一个实现 NIST P-256（FIPS 186-3，D.2.3 节）的曲线，也称为 secp256r1 或 prime256v1。

Multiple invocations of this function will return the same value, which can be used for equality checks and switch statements.

​	多次调用此函数将返回相同的值，该值可用于相等性检查和 switch 语句。

#### func P384 

``` go
func P384() Curve
```

P384 returns a Curve which implements NIST P-384 (FIPS 186-3, section D.2.4), also known as secp384r1.

​	P384 返回一个实现 NIST P-384（FIPS 186-3，D.2.4 节）的曲线，也称为 secp384r1。

Multiple invocations of this function will return the same value, which can be used for equality checks and switch statements.

​	多次调用此函数将返回相同的值，该值可用于相等性检查和 switch 语句。

#### func P521 

``` go
func P521() Curve
```

P521 returns a Curve which implements NIST P-521 (FIPS 186-3, section D.2.5), also known as secp521r1.

​	P521 返回一个实现 NIST P-521（FIPS 186-3，D.2.5 节）的曲线，也称为 secp521r1。

Multiple invocations of this function will return the same value, which can be used for equality checks and switch statements.

​	多次调用此函数将返回相同的值，该值可用于相等性检查和 switch 语句。

#### func X25519 

``` go
func X25519() Curve
```

X25519 returns a Curve which implements the X25519 function over Curve25519 ([RFC 7748, Section 5](https://rfc-editor.org/rfc/rfc7748.html#section-5)).

​	X25519 返回一个在 Curve25519（[RFC 7748, Section 5](https://rfc-editor.org/rfc/rfc7748.html#section-5)）上实现 X25519 函数的曲线。

Multiple invocations of this function will return the same value, so it can be used for equality checks and switch statements.

​	多次调用此函数将返回相同的值，因此可将其用于相等性检查和 switch 语句。

### type PrivateKey 

``` go
type PrivateKey struct {
	// contains filtered or unexported fields
}
```

PrivateKey is an ECDH private key, usually kept secret.

​	PrivateKey 是 ECDH 私钥，通常保持为机密。

These keys can be parsed with [crypto/x509.ParsePKCS8PrivateKey](https://pkg.go.dev/crypto/x509#ParsePKCS8PrivateKey) and encoded with [crypto/x509.MarshalPKCS8PrivateKey](https://pkg.go.dev/crypto/x509#MarshalPKCS8PrivateKey). For NIST curves, they then need to be converted with [crypto/ecdsa.PrivateKey.ECDH](https://pkg.go.dev/crypto/ecdsa#PrivateKey.ECDH) after parsing.

​	可以使用 [crypto/x509.ParsePKCS8PrivateKey]({{< ref "/stdLib/crypto/x509#func-parsepkcs8privatekey">}}) 解析这些密钥，并使用 [crypto/x509.MarshalPKCS8PrivateKey]({{< ref "/stdLib/crypto/x509#func-marshalpkcs8privatekey----go110">}}) 对其进行编码。对于 NIST 曲线，解析后需要使用 [crypto/ecdsa.PrivateKey.ECDH]({{< ref "/stdLib/crypto/ecdsa#publickey-ecdh----go120">}}) 对其进行转换。

#### (*PrivateKey) Bytes 

``` go
func (k *PrivateKey) Bytes() []byte
```

Bytes returns a copy of the encoding of the private key.

​	Bytes 返回私钥编码的副本。

#### (*PrivateKey) Curve 

``` go
func (k *PrivateKey) Curve() Curve
```

#### (*PrivateKey) ECDH 

``` go
func (k *PrivateKey) ECDH(remote *PublicKey) ([]byte, error)
```

ECDH performs a ECDH exchange and returns the shared secret.

​	ECDH 执行 ECDH 交换并返回共享密钥。

For NIST curves, this performs ECDH as specified in SEC 1, Version 2.0, Section 3.3.1, and returns the x-coordinate encoded according to SEC 1, Version 2.0, Section 2.3.5. The result is never the point at infinity.

​	对于 NIST 曲线，此操作按照 SEC 1 版本 2.0 第 3.3.1 节中的规定执行 ECDH，并返回根据 SEC 1 版本 2.0 第 2.3.5 节编码的 x 坐标。结果永远不是无穷远点。

For X25519, this performs ECDH as specified in [RFC 7748, Section 6.1](https://rfc-editor.org/rfc/rfc7748.html#section-6.1). If the result is the all-zero value, ECDH returns an error.

​	对于 X25519，此操作按照 [RFC 7748, Section 6.1](https://rfc-editor.org/rfc/rfc7748.html#section-6.1)中的规定执行 ECDH。如果结果为全零值，则 ECDH 返回错误。

#### (*PrivateKey) Equal 

``` go
func (k *PrivateKey) Equal(x crypto.PrivateKey) bool
```

Equal returns whether x represents the same private key as k.

​	Equal 返回 x 是否表示与 k 相同的私钥。

Note that there can be equivalent private keys with different encodings which would return false from this check but behave the same way as inputs to ECDH.

​	请注意，可能存在具有不同编码的等效私钥，这些私钥将从此检查返回 false，但作为 ECDH 的输入表现出相同的方式。

This check is performed in constant time as long as the key types and their curve match.

​	只要密钥类型及其曲线匹配，此检查就会在恒定时间内执行。

#### (*PrivateKey) Public 

``` go
func (k *PrivateKey) Public() crypto.PublicKey
```

Public implements the implicit interface of all standard library private keys. See the docs of crypto.PrivateKey.

​	Public 实现所有标准库私钥的隐式接口。请参阅 crypto.PrivateKey 的文档。

#### (*PrivateKey) PublicKey 

``` go
func (k *PrivateKey) PublicKey() *PublicKey
```

### type PublicKey 

``` go
type PublicKey struct {
	// contains filtered or unexported fields
}
```

PublicKey is an ECDH public key, usually a peer's ECDH share sent over the wire.

​	PublicKey 是 ECDH 公钥，通常是通过网络发送的对等方 ECDH 共享。

These keys can be parsed with [crypto/x509.ParsePKIXPublicKey](https://pkg.go.dev/crypto/x509#ParsePKIXPublicKey) and encoded with [crypto/x509.MarshalPKIXPublicKey](https://pkg.go.dev/crypto/x509#MarshalPKIXPublicKey). For NIST curves, they then need to be converted with [crypto/ecdsa.PublicKey.ECDH](https://pkg.go.dev/crypto/ecdsa#PublicKey.ECDH) after parsing.

​	可以使用 [crypto/x509.ParsePKIXPublicKey]({{< ref "/stdLib/crypto/x509#func-parsepkixpublickey">}}) 解析这些密钥，并使用 [crypto/x509.MarshalPKIXPublicKey]({{< ref "/stdLib/crypto/x509#func-marshalpkixpublickey">}}) 对其进行编码。对于 NIST 曲线，在解析后需要使用 [crypto/ecdsa.PublicKey.ECDH]({{< ref "/stdLib/crypto/ecdsa#privatekey-ecdh----go120">}}) 对其进行转换。



#### (*PublicKey) Bytes 

``` go
func (k *PublicKey) Bytes() []byte
```

Bytes returns a copy of the encoding of the public key.

​	Bytes 返回公钥编码的副本。

#### (*PublicKey) Curve 

``` go
func (k *PublicKey) Curve() Curve
```

#### (*PublicKey) Equal 

``` go
func (k *PublicKey) Equal(x crypto.PublicKey) bool
```

Equal returns whether x represents the same public key as k.

​	Equal 返回 x 是否表示与 k 相同的公钥。

Note that there can be equivalent public keys with different encodings which would return false from this check but behave the same way as inputs to ECDH.

​	请注意，可能存在具有不同编码的等效公钥，这些公钥将从此检查返回 false，但作为 ECDH 的输入表现出相同的方式。

This check is performed in constant time as long as the key types and their curve match.

​	只要密钥类型及其曲线匹配，此检查就会在恒定时间内执行。