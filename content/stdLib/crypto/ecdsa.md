+++
title = "ecdsa"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
> 原文：[https://pkg.go.dev/crypto/ecdsa@go1.24.2](https://pkg.go.dev/crypto/ecdsa@go1.24.2)

Package ecdsa implements the Elliptic Curve Digital Signature Algorithm, as defined in FIPS 186-4 and SEC 1, Version 2.0.

​	ecdsa 包实现了椭圆曲线数字签名算法，如 FIPS 186-4 和 SEC 1（版本 2.0）中所定义。

Signatures generated by this package are not deterministic, but entropy is mixed with the private key and the message, achieving the same level of security in case of randomness source failure.

​	此包生成的签名不是确定性的，但熵与私钥和消息混合，在随机性来源发生故障的情况下实现了相同级别的安全性。

## Example

```go
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	msg := "hello, world"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", sig)

	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)
}

```

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Sign 

``` go
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)
```

Sign signs a hash (which should be the result of hashing a larger message) using the private key, priv. If the hash is longer than the bit-length of the private key’s curve order, the hash will be truncated to that length. It returns the signature as a pair of integers. Most applications should use SignASN1 instead of dealing directly with r, s.

​	Sign 使用私钥 priv 对哈希（应该是对较大消息进行哈希的结果）进行签名。如果哈希长度超过私钥曲线阶的比特长度，则哈希将被截断为该长度。它将签名作为一对整数返回。大多数应用程序应使用 SignASN1，而不是直接处理 r、s。

### func SignASN1 <- go1.15

```go
func SignASN1(rand io.Reader, priv *PrivateKey, hash []byte) ([]byte, error)
```

SignASN1 signs a hash (which should be the result of hashing a larger message) using the private key, priv. If the hash is longer than the bit-length of the private key’s curve order, the hash will be truncated to that length. It returns the ASN.1 encoded signature.

​	SignASN1 使用私钥 priv 对哈希（应该是对较大消息进行哈希的结果）进行签名。如果哈希长度超过私钥曲线阶的比特长度，则哈希将被截断为该长度。它返回 ASN.1 编码的签名。

### func Verify

```go
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool
```

Verify verifies the signature in r, s of hash using the public key, pub. Its return value records whether the signature is valid. Most applications should use VerifyASN1 instead of dealing directly with r, s.

​	Verify 验证使用公钥 pub 对哈希的签名 r、s。其返回值记录签名是否有效。大多数应用程序应使用 VerifyASN1，而不是直接处理 r、s。

### func VerifyASN1 <- go1.15

```go
func VerifyASN1(pub *PublicKey, hash, sig []byte) bool
```

VerifyASN1 verifies the ASN.1 encoded signature, sig, of hash using the public key, pub. Its return value records whether the signature is valid.

​	VerifyASN1 验证使用公钥 pub 对哈希的 ASN.1 编码签名 sig。其返回值记录签名是否有效。

## 类型

### type PrivateKey

```go
type PrivateKey struct {
	PublicKey
	D *big.Int
}
```

PrivateKey represents an ECDSA private key.

​	PrivateKey 表示 ECDSA 私钥。

#### func GenerateKey

```go
func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error)
```

GenerateKey generates a public and private key pair.

​	GenerateKey 生成公钥和私钥对。

#### (*PrivateKey) ECDH <- go1.20

```go
func (k *PrivateKey) ECDH() (*ecdh.PrivateKey, error)
```

ECDH returns k as a [ecdh.PrivateKey](https://pkg.go.dev/crypto/ecdh#PrivateKey). It returns an error if the key is invalid according to the definition of [ecdh.Curve.NewPrivateKey](https://pkg.go.dev/crypto/ecdh#Curve.NewPrivateKey), or if the Curve is not supported by crypto/ecdh.

​	ECDH 将 k 作为 ecdh.PrivateKey 返回。如果密钥根据 ecdh.Curve.NewPrivateKey 的定义无效，或者 Curve 不受 crypto/ecdh 支持，则返回错误。

#### (*PrivateKey) Equal <- go1.15

```go
func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool
```

Equal reports whether priv and x have the same value.

​	Equal 报告 priv 和 x 是否具有相同的值。

See PublicKey.Equal for details on how Curve is compared.

​	有关如何比较 Curve 的详细信息，请参阅 PublicKey.Equal。

#### (*PrivateKey) Public <- go1.4

```go
func (priv *PrivateKey) Public() crypto.PublicKey
```

Public returns the public key corresponding to priv.

​	Public 返回与 priv 对应的公钥。

#### (*PrivateKey) Sign <- go1.4

```go
func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
```

Sign signs digest with priv, reading randomness from rand. The opts argument is not currently used but, in keeping with the crypto.Signer interface, should be the hash function used to digest the message.

​	Sign 使用 priv 对摘要进行签名，从 rand 中读取随机数。opts 参数当前未使用，但为了与 crypto.Signer 接口保持一致，应为用于对消息进行摘要处理的哈希函数。

This method implements crypto.Signer, which is an interface to support keys where the private part is kept in, for example, a hardware module. Common uses can use the SignASN1 function in this package directly.

​	此方法实现了 crypto.Signer，这是一个接口，用于支持私有部分保存在（例如）硬件模块中的密钥。常见用法可以直接使用此包中的 SignASN1 函数。

### type PublicKey

```go
type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}
```

PublicKey represents an ECDSA public key.

​	PublicKey 表示 ECDSA 公钥。

#### (*PublicKey) ECDH <- go1.20

```go
func (k *PublicKey) ECDH() (*ecdh.PublicKey, error)
```

ECDH returns k as a [ecdh.PublicKey](https://pkg.go.dev/crypto/ecdh#PublicKey). It returns an error if the key is invalid according to the definition of [ecdh.Curve.NewPublicKey](https://pkg.go.dev/crypto/ecdh#Curve.NewPublicKey), or if the Curve is not supported by crypto/ecdh.

​	ECDH 将 k 返回为 ecdh.PublicKey。如果密钥根据 ecdh.Curve.NewPublicKey 的定义无效，或者 Curve 不受 crypto/ecdh 支持，则返回错误。

#### (*PublicKey) Equal <- go1.15

```go
func (pub *PublicKey) Equal(x crypto.PublicKey) bool
```

Equal reports whether pub and x have the same value.

​	Equal 报告 pub 和 x 是否具有相同的值。

Two keys are only considered to have the same value if they have the same Curve value. Note that for example elliptic.P256() and elliptic.P256().Params() are different values, as the latter is a generic not constant time implementation.

​	只有当两个密钥具有相同的曲线值时，才认为它们具有相同的值。请注意，例如 elliptic.P256() 和 elliptic.P256().Params() 是不同的值，因为后者是一个通用的非恒定时间实现。