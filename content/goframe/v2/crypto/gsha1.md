+++
title = "gsha1"
date = 2024-03-21T17:47:08+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gsha1

Package gsha1 provides useful API for SHA1 encryption algorithms.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Encrypt 

``` go
func Encrypt(v interface{}) string
```

Encrypt encrypts any type of variable using SHA1 algorithms. It uses package gconv to convert `v` to its bytes type.

##### func EncryptFile 

``` go
func EncryptFile(path string) (encrypt string, err error)
```

EncryptFile encrypts file content of `path` using SHA1 algorithms.

##### func MustEncryptFile 

``` go
func MustEncryptFile(path string) string
```

MustEncryptFile encrypts file content of `path` using SHA1 algorithms. It panics if any error occurs.

### Types 

This section is empty.