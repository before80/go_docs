+++
title = "gbase64"
date = 2024-03-21T17:48:52+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gbase64

Package gbase64 provides useful API for BASE64 encoding/decoding algorithm.



### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Decode 

``` go
func Decode(data []byte) ([]byte, error)
```

Decode decodes bytes with BASE64 algorithm.

##### func DecodeString 

``` go
func DecodeString(data string) ([]byte, error)
```

DecodeString decodes string with BASE64 algorithm.

##### func DecodeToString 

``` go
func DecodeToString(data string) (string, error)
```

DecodeToString decodes string with BASE64 algorithm.

##### func Encode 

``` go
func Encode(src []byte) []byte
```

Encode encodes bytes with BASE64 algorithm.

##### func EncodeFile 

``` go
func EncodeFile(path string) ([]byte, error)
```

EncodeFile encodes file content of `path` using BASE64 algorithms.

##### func EncodeFileToString 

``` go
func EncodeFileToString(path string) (string, error)
```

EncodeFileToString encodes file content of `path` to string using BASE64 algorithms.

##### func EncodeString 

``` go
func EncodeString(src string) string
```

EncodeString encodes string with BASE64 algorithm.

##### func EncodeToString 

``` go
func EncodeToString(src []byte) string
```

EncodeToString encodes bytes to string with BASE64 algorithm.

##### func MustDecode 

``` go
func MustDecode(data []byte) []byte
```

MustDecode decodes bytes with BASE64 algorithm. It panics if any error occurs.

##### func MustDecodeString 

``` go
func MustDecodeString(data string) []byte
```

MustDecodeString decodes string with BASE64 algorithm. It panics if any error occurs.

##### func MustDecodeToString 

``` go
func MustDecodeToString(data string) string
```

MustDecodeToString decodes string with BASE64 algorithm. It panics if any error occurs.

##### func MustEncodeFile 

``` go
func MustEncodeFile(path string) []byte
```

MustEncodeFile encodes file content of `path` using BASE64 algorithms. It panics if any error occurs.

##### func MustEncodeFileToString 

``` go
func MustEncodeFileToString(path string) string
```

MustEncodeFileToString encodes file content of `path` to string using BASE64 algorithms. It panics if any error occurs.

### Types 

This section is empty.