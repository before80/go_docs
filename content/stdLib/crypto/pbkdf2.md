+++
title = "pbkdf2"
date = 2025-04-01T13:15:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
> 原文：[https://pkg.go.dev/crypto/fips140@go1.24.2](https://pkg.go.dev/crypto/fips140@go1.24.2)

> 注意
>
> ​	从go1.24.0开始才可以使用该包。

## Overview

Package pbkdf2 implements the key derivation function PBKDF2 as defined in [RFC 8018](https://rfc-editor.org/rfc/rfc8018.html) (PKCS #5 v2.1).

A key derivation function is useful when encrypting data based on a password or any other not-fully-random data. It uses a pseudorandom function to derive a secure encryption key based on the password.

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func Key

```go
func Key[Hash hash.Hash](h func() Hash, password string, salt []byte, iter, keyLength int) ([]byte, error)
```

Key derives a key from the password, salt and iteration count, returning a []byte of length keyLength that can be used as cryptographic key. The key is derived based on the method described as PBKDF2 with the HMAC variant using the supplied hash function.

For example, to use a HMAC-SHA-1 based PBKDF2 key derivation function, you can get a derived key for e.g. AES-256 (which needs a 32-byte key) by doing:

```go
dk := pbkdf2.Key(sha1.New, []byte("some password"), salt, 4096, 32)
```

Remember to get a good random salt. At least 8 bytes is recommended by the RFC.

Using a higher iteration count will increase the cost of an exhaustive search but will also make derivation proportionally slower.

keyLength must be a positive integer between 1 and `(2^32 - 1) * h.Size().` Setting keyLength to a value outside of this range will result in an error.

## 类型

This section is empty.