+++
title = "gcrc32"
date = 2024-03-21T17:46:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gcrc32

### Overview 

Package gcrc32 provides useful API for CRC32 encryption algorithms.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Encrypt 

``` go
func Encrypt(v interface{}) uint32
```

Encrypt encrypts any type of variable using CRC32 algorithms. It uses gconv package to convert `v` to its bytes type.

### Types 

This section is empty.