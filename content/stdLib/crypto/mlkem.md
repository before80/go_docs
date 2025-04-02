+++
title = "mlkem"
date = 2025-04-01T13:15:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
> 原文：[https://pkg.go.dev/crypto/mlkem@go1.24.2](https://pkg.go.dev/crypto/mlkem@go1.24.2)

> 注意
>
> ​	从go1.24.0开始才可以使用该包。

## Overview

Package mlkem implements the quantum-resistant key encapsulation method ML-KEM (formerly known as Kyber), as specified in [NIST FIPS 203](https://doi.org/10.6028/NIST.FIPS.203).

Most applications should use the ML-KEM-768 parameter set, as implemented by [DecapsulationKey768](https://pkg.go.dev/crypto/mlkem@go1.24.2#DecapsulationKey768) and [EncapsulationKey768](https://pkg.go.dev/crypto/mlkem@go1.24.2#EncapsulationKey768).

### Example

```go
package main

import (
	"crypto/mlkem"
	"log"
)

func main() {
	// Alice generates a new key pair and sends the encapsulation key to Bob.
	dk, err := mlkem.GenerateKey768()
	if err != nil {
		log.Fatal(err)
	}
	encapsulationKey := dk.EncapsulationKey().Bytes()

	// Bob uses the encapsulation key to encapsulate a shared secret, and sends
	// back the ciphertext to Alice.
	ciphertext := Bob(encapsulationKey)

	// Alice decapsulates the shared secret from the ciphertext.
	sharedSecret, err := dk.Decapsulate(ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	// Alice and Bob now share a secret.
	_ = sharedSecret
}

func Bob(encapsulationKey []byte) (ciphertext []byte) {
	// Bob encapsulates a shared secret using the encapsulation key.
	ek, err := mlkem.NewEncapsulationKey768(encapsulationKey)
	if err != nil {
		log.Fatal(err)
	}
	sharedSecret, ciphertext := ek.Encapsulate()

	// Alice and Bob now share a secret.
	_ = sharedSecret

	// Bob sends the ciphertext to Alice.
	return ciphertext
}

```



## 常量

```go
const (
	// SharedKeySize is the size of a shared key produced by ML-KEM.
	SharedKeySize = 32

	// SeedSize is the size of a seed used to generate a decapsulation key.
	SeedSize = 64

	// CiphertextSize768 is the size of a ciphertext produced by ML-KEM-768.
	CiphertextSize768 = 1088

	// EncapsulationKeySize768 is the size of an ML-KEM-768 encapsulation key.
	EncapsulationKeySize768 = 1184

	// CiphertextSize1024 is the size of a ciphertext produced by ML-KEM-1024.
	CiphertextSize1024 = 1568

	// EncapsulationKeySize1024 is the size of an ML-KEM-1024 encapsulation key.
	EncapsulationKeySize1024 = 1568
)
```



## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type DecapsulationKey1024 

```go
type DecapsulationKey1024 struct {
	// contains filtered or unexported fields
}
```

DecapsulationKey1024 is the secret key used to decapsulate a shared key from a ciphertext. It includes various precomputed values.

#### func GenerateKey1024 

```go
func GenerateKey1024() (*DecapsulationKey1024, error)
```

GenerateKey1024 generates a new decapsulation key, drawing random bytes from the default crypto/rand source. The decapsulation key must be kept secret.

#### func NewDecapsulationKey1024 

```go
func NewDecapsulationKey1024(seed []byte) (*DecapsulationKey1024, error)
```

NewDecapsulationKey1024 expands a decapsulation key from a 64-byte seed in the "d || z" form. The seed must be uniformly random.

#### (*DecapsulationKey1024) Bytes 

```go
func (dk *DecapsulationKey1024) Bytes() []byte
```

Bytes returns the decapsulation key as a 64-byte seed in the "d || z" form.

The decapsulation key must be kept secret.

#### (*DecapsulationKey1024) Decapsulate 

```go
func (dk *DecapsulationKey1024) Decapsulate(ciphertext []byte) (sharedKey []byte, err error)
```

Decapsulate generates a shared key from a ciphertext and a decapsulation key. If the ciphertext is not valid, Decapsulate returns an error.

The shared key must be kept secret.

#### (*DecapsulationKey1024) EncapsulationKey 

```go
func (dk *DecapsulationKey1024) EncapsulationKey() *EncapsulationKey1024
```

EncapsulationKey returns the public encapsulation key necessary to produce ciphertexts.

### type DecapsulationKey768 

```go
type DecapsulationKey768 struct {
	// contains filtered or unexported fields
}
```

DecapsulationKey768 is the secret key used to decapsulate a shared key from a ciphertext. It includes various precomputed values.

#### func GenerateKey768 

```go
func GenerateKey768() (*DecapsulationKey768, error)
```

GenerateKey768 generates a new decapsulation key, drawing random bytes from the default crypto/rand source. The decapsulation key must be kept secret.

#### func NewDecapsulationKey768 

```
func NewDecapsulationKey768(seed []byte) (*DecapsulationKey768, error)
```

NewDecapsulationKey768 expands a decapsulation key from a 64-byte seed in the "d || z" form. The seed must be uniformly random.

#### (*DecapsulationKey768) Bytes 

```go
func (dk *DecapsulationKey768) Bytes() []byte
```

Bytes returns the decapsulation key as a 64-byte seed in the "d || z" form.

The decapsulation key must be kept secret.

#### (*DecapsulationKey768) Decapsulate 

```go
func (dk *DecapsulationKey768) Decapsulate(ciphertext []byte) (sharedKey []byte, err error)
```

Decapsulate generates a shared key from a ciphertext and a decapsulation key. If the ciphertext is not valid, Decapsulate returns an error.

The shared key must be kept secret.

#### (*DecapsulationKey768) EncapsulationKey 

```go
func (dk *DecapsulationKey768) EncapsulationKey() *EncapsulationKey768
```

EncapsulationKey returns the public encapsulation key necessary to produce ciphertexts.

### type EncapsulationKey1024 

```go
type EncapsulationKey1024 struct {
	// contains filtered or unexported fields
}
```

An EncapsulationKey1024 is the public key used to produce ciphertexts to be decapsulated by the corresponding DecapsulationKey1024.

#### func NewEncapsulationKey1024 

```go
func NewEncapsulationKey1024(encapsulationKey []byte) (*EncapsulationKey1024, error)
```

NewEncapsulationKey1024 parses an encapsulation key from its encoded form. If the encapsulation key is not valid, NewEncapsulationKey1024 returns an error.

#### (*EncapsulationKey1024) Bytes 

```go
func (ek *EncapsulationKey1024) Bytes() []byte
```

Bytes returns the encapsulation key as a byte slice.

#### (*EncapsulationKey1024) Encapsulate 

```go
func (ek *EncapsulationKey1024) Encapsulate() (sharedKey, ciphertext []byte)
```

Encapsulate generates a shared key and an associated ciphertext from an encapsulation key, drawing random bytes from the default crypto/rand source.

The shared key must be kept secret.

### type EncapsulationKey768 

```go
type EncapsulationKey768 struct {
	// contains filtered or unexported fields
}
```

An EncapsulationKey768 is the public key used to produce ciphertexts to be decapsulated by the corresponding DecapsulationKey768.

#### func NewEncapsulationKey768 

```go
func NewEncapsulationKey768(encapsulationKey []byte) (*EncapsulationKey768, error)
```

NewEncapsulationKey768 parses an encapsulation key from its encoded form. If the encapsulation key is not valid, NewEncapsulationKey768 returns an error.

#### (*EncapsulationKey768) Bytes 

```go
func (ek *EncapsulationKey768) Bytes() []byte
```

Bytes returns the encapsulation key as a byte slice.

#### (*EncapsulationKey768) Encapsulate 

```go
func (ek *EncapsulationKey768) Encapsulate() (sharedKey, ciphertext []byte)
```

Encapsulate generates a shared key and an associated ciphertext from an encapsulation key, drawing random bytes from the default crypto/rand source.

The shared key must be kept secret.