+++
title = "gmd5"
date = 2024-03-21T17:46:49+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gmd5

Package gmd5 provides useful API for MD5 encryption algorithms.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Encrypt 

``` go
func Encrypt(data interface{}) (encrypt string, err error)
```

Encrypt encrypts any type of variable using MD5 algorithms. It uses gconv package to convert `v` to its bytes type.

##### func EncryptBytes 

``` go
func EncryptBytes(data []byte) (encrypt string, err error)
```

EncryptBytes encrypts `data` using MD5 algorithms.

##### func EncryptFile 

``` go
func EncryptFile(path string) (encrypt string, err error)
```

EncryptFile encrypts file content of `path` using MD5 algorithms.

##### func EncryptString 

``` go
func EncryptString(data string) (encrypt string, err error)
```

EncryptString encrypts string `data` using MD5 algorithms.

##### func MustEncrypt 

``` go
func MustEncrypt(data interface{}) string
```

MustEncrypt encrypts any type of variable using MD5 algorithms. It uses gconv package to convert `v` to its bytes type. It panics if any error occurs.

##### func MustEncryptBytes 

``` go
func MustEncryptBytes(data []byte) string
```

MustEncryptBytes encrypts `data` using MD5 algorithms. It panics if any error occurs.

##### func MustEncryptFile 

``` go
func MustEncryptFile(path string) string
```

MustEncryptFile encrypts file content of `path` using MD5 algorithms. It panics if any error occurs.

##### func MustEncryptString 

``` go
func MustEncryptString(data string) string
```

MustEncryptString encrypts string `data` using MD5 algorithms. It panics if any error occurs.

### Types 

This section is empty.