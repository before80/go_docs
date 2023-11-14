+++
title = "utf16"
date = 2023-05-17T09:59:21+08:00
weight = 20
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/unicode/utf16@go1.21.3

Package utf16 implements encoding and decoding of UTF-16 sequences.

​	`utf16`包实现了UTF-16序列的编码和解码。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func AppendRune  <- go1.20

``` go 
func AppendRune(a []uint16, r rune) []uint16
```

AppendRune appends the UTF-16 encoding of the Unicode code point r to the end of p and returns the extended buffer. If the rune is not a valid Unicode code point, it appends the encoding of U+FFFD.

​	`AppendRune`函数将Unicode码点r的UTF-16编码追加到`a`的末尾，并返回扩展后的缓冲区。如果`r`不是有效的Unicode码点，则追加`U+FFFD`的编码。

### func Decode 

``` go 
func Decode(s []uint16) []rune
```

Decode returns the Unicode code point sequence represented by the UTF-16 encoding s.

​	`Decode`函数返回由UTF-16编码`s`表示的Unicode码点序列。

### func DecodeRune 

``` go 
func DecodeRune(r1, r2 rune) rune
```

DecodeRune returns the UTF-16 decoding of a surrogate pair. If the pair is not a valid UTF-16 surrogate pair, DecodeRune returns the Unicode replacement code point U+FFFD.

​	`DecodeRune`函数返回代理对的UTF-16解码。如果代理对不是有效的UTF-16代理对，则`DecodeRune`返回Unicode替换码点`U+FFFD`。

### func Encode 

``` go 
func Encode(s []rune) []uint16
```

Encode returns the UTF-16 encoding of the Unicode code point sequence s.

​	`Encode`函数返回Unicode码点序列`s`的UTF-16编码。

### func EncodeRune 

``` go 
func EncodeRune(r rune) (r1, r2 rune)
```

EncodeRune returns the UTF-16 surrogate pair r1, r2 for the given rune. If the rune is not a valid Unicode code point or does not need encoding, EncodeRune returns U+FFFD, U+FFFD.

​	`EncodeRune`函数返回给定rune的UTF-16代理对`r1`、`r2`。如果`rune`不是有效的Unicode码点或不需要编码，则`EncodeRune`返回`U+FFFD`，`U+FFFD`。

### func IsSurrogate 

``` go 
func IsSurrogate(r rune) bool
```

IsSurrogate reports whether the specified Unicode code point can appear in a surrogate pair.

​	`IsSurrogate`函数报告指定的Unicode码点是否可以出现在代理对中。

## 类型

This section is empty.