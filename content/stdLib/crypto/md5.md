+++
title = "md5"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# md5

https://pkg.go.dev/crypto/md5@go1.20.1



Package md5 implements the MD5 hash algorithm as defined in [RFC 1321](https://rfc-editor.org/rfc/rfc1321.html).

MD5 is cryptographically broken and should not be used for secure applications.













## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/md5/md5.go;l=28)

```
const BlockSize = 64
```

The blocksize of MD5 in bytes.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/md5/md5.go;l=25)

```
const Size = 16
```

The size of an MD5 checksum in bytes.

## 变量

This section is empty.

## 函数

#### func New 

```
func New() hash.Hash
```

New returns a new hash.Hash computing the MD5 checksum. The Hash also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

##### Example

##### Example (File)

#### func Sum  <- go1.2

```
func Sum(data []byte) [Size]byte
```

Sum returns the MD5 checksum of the data.

##### Example

## 类型

This section is empty.