+++
title = "sha1"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# sha1

https://pkg.go.dev/crypto/sha1@go1.20.1



Package sha1 implements the SHA-1 hash algorithm as defined in [RFC 3174](https://rfc-editor.org/rfc/rfc3174.html).

SHA-1 is cryptographically broken and should not be used for secure applications.













## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/sha1/sha1.go;l=26)

```
const BlockSize = 64
```

The blocksize of SHA-1 in bytes.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/sha1/sha1.go;l=23)

```
const Size = 20
```

The size of a SHA-1 checksum in bytes.

## 变量

This section is empty.

## 函数

#### func New 

```
func New() hash.Hash
```

New returns a new hash.Hash computing the SHA1 checksum. The Hash also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

##### Example

##### Example (File)

#### func Sum  <- go1.2

```
func Sum(data []byte) [Size]byte
```

Sum returns the SHA-1 checksum of the data.

##### Example

## 类型

This section is empty.