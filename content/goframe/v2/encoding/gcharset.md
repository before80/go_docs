+++
title = "gcharset"
date = 2024-03-21T17:49:10+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gcharset](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gcharset)

Package gcharset implements character-set conversion functionality.

​	Package gcharset 实现了字符集转换功能。

Supported Character Set:

​	支持的字符集：

Chinese : GBK/GB18030/GB2312/Big5

​	中文 ： GBK/GB18030/GB2312/Big5

Japanese: EUCJP/ISO2022JP/ShiftJIS

​	日语：EUCJP/ISO2022JP/ShiftJIS

Korean : EUCKR

​	韩语 ： EUCKR

Unicode : UTF-8/UTF-16/UTF-16BE/UTF-16LE

​	统一代码：UTF-8/UTF-16/UTF-16BE/UTF-16LE

Other : macintosh/IBM*/Windows*/ISO-*

​	其他 ： macintosh/IBM*/Windows*/ISO-*

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Convert

```go
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error)
```

Convert converts `src` charset encoding from `srcCharset` to `dstCharset`, and returns the converted string. It returns `src` as `dst` if it fails converting.

​	Convert 将 `src` 字符集编码从 `srcCharset` 转换为 `dstCharset` ，并返回转换后的字符串。它返回 `src` ，就好像 `dst` 转换失败一样。

#### func Supported

```go
func Supported(charset string) bool
```

Supported returns whether charset `charset` is supported.

​	Supported 返回是否支持 charset `charset` 。

#### func ToUTF8

```go
func ToUTF8(srcCharset string, src string) (dst string, err error)
```

ToUTF8 converts `src` charset encoding from `srcCharset` to UTF-8 , and returns the converted string.

​	ToUTF8 将 `src` 字符集编码从 `srcCharset` 转换为 UTF-8 ，并返回转换后的字符串。

#### func UTF8To

```go
func UTF8To(dstCharset string, src string) (dst string, err error)
```

UTF8To converts `src` charset encoding from UTF-8 to `dstCharset`, and returns the converted string.

​	UTF8To 将 `src` 字符集编码从 UTF-8 转换为 `dstCharset` ，并返回转换后的字符串。

## 类型

This section is empty.