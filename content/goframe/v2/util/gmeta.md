+++
title = "gmeta"
date = 2024-03-21T17:59:27+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gmeta](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gmeta)

Package gmeta provides embedded meta data feature for struct.

​	软件包 gmeta 为 struct 提供嵌入式元数据功能。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Data

```go
func Data(object interface{}) map[string]string
```

Data retrieves and returns all metadata from `object`.

​	数据检索并返回 中 `object` 的所有元数据。

#### func Get

```go
func Get(object interface{}, key string) *gvar.Var
```

Get retrieves and returns specified metadata by `key` from `object`.

​	Get 检索并返回指定的 `key` 元数据 from `object` 。

## 类型

### type Meta

```go
type Meta struct{}
```

Meta is used as an embedded attribute for struct to enabled metadata feature.

​	Meta 用作 struct 启用元数据功能的嵌入属性。