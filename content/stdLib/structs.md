+++
title = "structs"
date = 2024-09-06T10:46:23+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/structs@go1.23.0](https://pkg.go.dev/structs@go1.23.0)

> 注意
>
> ​	从go1.23.0开始才可以使用该包。

## Overview 

Package structs defines marker types that can be used as struct fields to modify the properties of a struct.

​	`structs` 包定义了标记类型，这些类型可以用作结构体字段来修改结构体的属性。

By convention, a marker type should be used as the type of a field named "_", placed at the beginning of a struct type definition.

​	按照惯例，标记类型应作为名为 "_" 的字段的类型，并放置在结构体类型定义的开头。

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数

This section is empty.

## 类型

### type HostLayout

```go
type HostLayout struct {
	// contains filtered or unexported fields
}
```

HostLayout marks a struct as using host memory layout. A struct with a field of type HostLayout will be laid out in memory according to host expectations, generally following the host's C ABI.

​	`HostLayout` 用于标记一个结构体采用主机内存布局。包含 `HostLayout` 类型字段的结构体将在内存中根据主机的预期进行布局，通常遵循主机的 C ABI（应用二进制接口）。

HostLayout does not affect layout within any other struct-typed fields of the containing struct, nor does it affect layout of structs containing the struct marked as host layout.

​	`HostLayout` 不会影响包含该结构体的其他结构体类型字段的布局，也不会影响包含标记为主机布局的结构体的布局。

By convention, HostLayout should be used as the type of a field named "_", placed at the beginning of the struct type definition.

​	按照惯例，`HostLayout` 应作为名为 "_" 的字段的类型，并放置在结构体类型定义的开头。
