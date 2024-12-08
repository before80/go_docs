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

Provided functionalities: - sorting - comparators

### Index 

- [func TimeComparator(a, b time.Time) int](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/utils#TimeComparator)
- [func ToString(value interface{}) string](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/utils#ToString)
- [type Comparator](https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha/utils#Comparator)

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

#### func ToString 

``` go
func ToString(value interface{}) string
```

ToString converts a value to string.

## 类型 

### type Comparator 

``` go
type Comparator[T any] func(x, y T) int
```
