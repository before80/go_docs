+++
title = "gproperties"
date = 2024-03-21T17:49:57+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gproperties](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gproperties)

Package gproperties provides accessing and converting for .properties content.

​	软件包 gproperties 提供对 .properties 内容的访问和转换。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Decode

```go
func Decode(data []byte) (res map[string]interface{}, err error)
```

Decode converts properties format to map.

​	解码将属性格式转换为映射。

#### func Encode

```go
func Encode(data map[string]interface{}) (res []byte, err error)
```

Encode converts map to properties format.

​	Encode 将映射转换为属性格式。

#### func ToJson

```go
func ToJson(data []byte) (res []byte, err error)
```

ToJson convert .properties format to JSON.

​	ToJson 将 .properties 格式转换为 JSON。

## 类型

This section is empty.