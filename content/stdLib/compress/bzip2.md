+++
title = "bzip2"
date = 2023-05-17T09:59:21+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/compress/bzip2@go1.23.0](https://pkg.go.dev/compress/bzip2@go1.23.0)

Package bzip2 implements bzip2 decompression.

​	bzip2 包实现 bzip2 的解压。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func NewReader 

``` go 
func NewReader(r io.Reader) io.Reader
```

NewReader returns an io.Reader which decompresses bzip2 data from r. If r does not also implement io.ByteReader, the decompressor may read more data than necessary from r.

​	NewReader 返回一个 io.Reader，它从 r 中解压 bzip2 数据。如果 r 也没有实现 io.ByteReader，则解压器可能会从 r 中读取多余的数据。

## 类型

### type StructuralError 

``` go 
type StructuralError string
```

A StructuralError is returned when the bzip2 data is found to be syntactically invalid.

​	当发现 bzip2 数据在语法上不正确时，将返回一个StructuralError。

#### (StructuralError) Error 

``` go 
func (s StructuralError) Error() string
```