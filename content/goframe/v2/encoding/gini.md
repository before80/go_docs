+++
title = "gini"
date = 2024-03-21T17:49:41+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gini](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gini)

Package gini provides accessing and converting for INI content.

​	软件包 gini 提供对 INI 内容的访问和转换。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Decode

```go
func Decode(data []byte) (res map[string]interface{}, err error)
```

Decode converts INI format to map.

​	Decode 将 INI 格式转换为映射。

#### func Encode

```go
func Encode(data map[string]interface{}) (res []byte, err error)
```

Encode converts map to INI format.

​	Encode 将映射转换为 INI 格式。

#### func ToJson

```go
func ToJson(data []byte) (res []byte, err error)
```

ToJson convert INI format to JSON.

​	ToJson 将 INI 格式转换为 JSON。

## 类型

This section is empty.