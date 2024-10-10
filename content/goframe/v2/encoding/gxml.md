+++
title = "gxml"
date = 2024-03-21T17:50:21+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gxml](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gxml)

Package gxml provides accessing and converting for XML content.

​	软件包 gxml 提供对 XML 内容的访问和转换。

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

#### func DecodeWithoutRoot

```go
func DecodeWithoutRoot(content []byte) (map[string]interface{}, error)
```

DecodeWithoutRoot parses `content` into a map, and returns the map without root level.

​	DecodeWithoutRoot 解析 `content` 为映射，并返回没有根级别的映射。

#### func Encode

```go
func Encode(m map[string]interface{}, rootTag ...string) ([]byte, error)
```

Encode encodes map `m` to an XML format content as bytes. The optional parameter `rootTag` is used to specify the XML root tag.

​	Encode 将映射 `m` 到 XML 格式的内容编码为字节。可选参数 `rootTag` 用于指定 XML 根标记。

#### func EncodeWithIndent

```go
func EncodeWithIndent(m map[string]interface{}, rootTag ...string) ([]byte, error)
```

EncodeWithIndent encodes map `m` to an XML format content as bytes with indent. The optional parameter `rootTag` is used to specify the XML root tag.

​	EncodeWithIndent 将映射 `m` 到 XML 格式的内容编码为带缩进的字节。可选参数 `rootTag` 用于指定 XML 根标记。

#### func ToJson

```go
func ToJson(content []byte) ([]byte, error)
```

ToJson converts `content` as XML format into JSON format bytes.

​	ToJson 将 `content` XML 格式转换为 JSON 格式字节。

## 类型

This section is empty.