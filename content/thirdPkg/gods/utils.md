+++
title = "utils"
date = 2024-12-07T11:10:14+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/utils](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/utils)
>
> 收录该文档时间： `2024-12-07T11:10:14+08:00`

## Overview 

Package utils provides common utility functions.

​	`utils` 包提供了常用的工具函数。

Provided functionalities: 提供的功能：

- sorting  排序
- comparators 排序

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数 

#### func TimeComparator 

``` go
func TimeComparator(a, b time.Time) int
```

TimeComparator provides a basic comparison on time.Time

​	`TimeComparator` 提供对 `time.Time` 类型的基本比较。

#### func ToString 

``` go
func ToString(value interface{}) string
```

ToString converts a value to string.

​	`ToString` 将一个值转换为字符串。

## 类型 

### type Comparator 

``` go
type Comparator[T any] func(x, y T) int
```
