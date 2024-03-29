+++
title = "gyaml"
date = 2024-03-21T17:50:29+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gyaml](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gyaml)

Package gyaml provides accessing and converting for YAML content.

​	Package gyaml 提供对 YAML 内容的访问和转换。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Decode

```go
func Decode(content []byte) (map[string]interface{}, error)
```

Decode parses `content` into and returns as map.

​	解码解 `content` 析为映射并返回。

#### func DecodeTo

```go
func DecodeTo(value []byte, result interface{}) (err error)
```

DecodeTo parses `content` into `result`.

​	DecodeTo 解析 `content` 为 `result` .

#### func Encode

```go
func Encode(value interface{}) (out []byte, err error)
```

Encode encodes `value` to an YAML format content as bytes.

​	编码 `value` 编码为 YAML 格式的内容，以字节表示。

#### func EncodeIndent

```go
func EncodeIndent(value interface{}, indent string) (out []byte, err error)
```

EncodeIndent encodes `value` to an YAML format content with indent as bytes.

​	EncodeIndent 编码 `value` 为 YAML 格式的内容，缩进为字节。

#### func ToJson

```go
func ToJson(content []byte) (out []byte, err error)
```

ToJson converts `content` to JSON format content.

​	ToJson `content` 转换为 JSON 格式的内容。

## 类型

This section is empty.