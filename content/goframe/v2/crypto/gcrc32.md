+++
title = "gcrc32"
date = 2024-03-21T17:46:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gcrc32](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/crypto/gcrc32)

### Overview 概述

Package gcrc32 provides useful API for CRC32 encryption algorithms.

​	软件包 gcrc32 为 CRC32 加密算法提供了有用的 API。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Encrypt

```go
func Encrypt(v interface{}) uint32
```

Encrypt encrypts any type of variable using CRC32 algorithms. It uses gconv package to convert `v` to its bytes type.

​	Encrypt 使用 CRC32 算法对任何类型的变量进行加密。它使用 gconv 包转换为 `v` 其字节类型。

## 类型

This section is empty.