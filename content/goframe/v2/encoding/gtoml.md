+++
title = "gtoml"
date = 2024-03-21T17:50:06+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gtoml](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gtoml)

Package gtoml provides accessing and converting for TOML content.

​	软件包 gtoml 提供对 TOML 内容的访问和转换。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Decode

```go
func Decode(v []byte) (interface{}, error)
```

#### func DecodeTo

```go
func DecodeTo(v []byte, result interface{}) (err error)
```

#### func Encode

```go
func Encode(v interface{}) ([]byte, error)
```

#### func ToJson

```go
func ToJson(v []byte) ([]byte, error)
```

## 类型

This section is empty.