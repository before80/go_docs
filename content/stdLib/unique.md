+++
title = "unique"
date = 2024-09-06T11:48:45+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/unique@go1.23.0](https://pkg.go.dev/unique@go1.23.0)

> 注意
>
> ​	从go1.23.0开始才可以使用该包。

## Overview 

The unique package provides facilities for canonicalizing ("interning") comparable values.

​	`unique` 包提供了用于规范化（“驻留”）可比较值的功能。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Handle 

``` go
type Handle[T comparable] struct {
	// contains filtered or unexported fields
}
```

Handle is a globally unique identity for some value of type T.

​	`Handle` 是某个类型为 T 的值的全局唯一标识。

Two handles compare equal exactly if the two values used to create the handles would have also compared equal. The comparison of two handles is trivial and typically much more efficient than comparing the values used to create them.

​	当且仅当创建 `Handle` 的两个值相等时，两个 `Handle` 才相等。比较两个 `Handle` 的操作非常简单，通常比比较创建它们的值更加高效。

#### func Make 

``` go
func Make[T comparable](value T) Handle[T]
```

Make returns a globally unique handle for a value of type T. Handles are equal if and only if the values used to produce them are equal.

​	`Make` 返回一个值类型为 T 的全局唯一 `Handle`。当且仅当用于生成 `Handle` 的值相等时，`Handle` 才相等。

#### (Handle[T]) Value 

``` go
func (h Handle[T]) Value() T
```

Value returns a shallow copy of the T value that produced the Handle.

​	`Value` 返回生成 `Handle` 的 T 值的浅拷贝。
