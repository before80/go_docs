+++
title = "gcharset"
date = 2024-03-21T17:49:10+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gcharset

Package gcharset implements character-set conversion functionality.

Supported Character Set:

Chinese : GBK/GB18030/GB2312/Big5

Japanese: EUCJP/ISO2022JP/ShiftJIS

Korean : EUCKR

Unicode : UTF-8/UTF-16/UTF-16BE/UTF-16LE

Other : macintosh/IBM*/Windows*/ISO-*

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Convert 

``` go
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error)
```

Convert converts `src` charset encoding from `srcCharset` to `dstCharset`, and returns the converted string. It returns `src` as `dst` if it fails converting.

##### func Supported 

``` go
func Supported(charset string) bool
```

Supported returns whether charset `charset` is supported.

##### func ToUTF8 

``` go
func ToUTF8(srcCharset string, src string) (dst string, err error)
```

ToUTF8 converts `src` charset encoding from `srcCharset` to UTF-8 , and returns the converted string.

##### func UTF8To 

``` go
func UTF8To(dstCharset string, src string) (dst string, err error)
```

UTF8To converts `src` charset encoding from UTF-8 to `dstCharset`, and returns the converted string.

### Types 

This section is empty.