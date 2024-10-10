+++
title = "gurl"
date = 2024-03-21T17:50:13+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gurl](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gurl)

Package gurl provides useful API for URL handling.

​	Package girl 为 URL 处理提供了有用的 API。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func BuildQuery

```go
func BuildQuery(queryData url.Values) string
```

BuildQuery Generate URL-encoded query string. See http://php.net/manual/en/function.http-build-query.php.

​	BuildQuery 生成 URL 编码的查询字符串。请参见 http://php.net/manual/en/function.http-build-query.php。

#### func Decode

```go
func Decode(str string) (string, error)
```

Decode does the inverse transformation of Encode, converting each 3-byte encoded substring of the form “%AB” into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits.

​	Decode 执行 Encode 的反向变换，将“%AB”形式的每个 3 字节编码子字符串转换为十六进制解码字节0xAB。如果任何 % 后面没有两个十六进制数字，则返回错误。

#### func Encode

```go
func Encode(str string) string
```

Encode escapes the string so it can be safely placed inside an URL query.

​	Encode 转义字符串，以便可以安全地将其放置在 URL 查询中。

#### func ParseURL

```go
func ParseURL(str string, component int) (map[string]string, error)
```

ParseURL Parse an URL and return its components. -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment. See http://php.net/manual/en/function.parse-url.php.

​	ParseURL 解析 URL 并返回其组件。-1：全部;1：方案;2：主机;4：端口;8：用户;16：及格;32：路径;64：查询;128：片段。请参见 http://php.net/manual/en/function.parse-url.php。

#### func RawDecode

```go
func RawDecode(str string) (string, error)
```

RawDecode does decode the given string Decode URL-encoded strings. See http://php.net/manual/en/function.rawurldecode.php.

​	RawDecode 会解码给定的字符串 Decode URL 编码的字符串。请参见 http://php.net/manual/en/function.rawurldecode.php。

#### func RawEncode

```go
func RawEncode(str string) string
```

RawEncode does encode the given string according URL-encode according to [RFC 3986](https://rfc-editor.org/rfc/rfc3986.html). See http://php.net/manual/en/function.rawurlencode.php.

​	RawEncode 确实根据 RFC 3986 的 URL 编码对给定的字符串进行编码。请参见 http://php.net/manual/en/function.rawurlencode.php。

## 类型

This section is empty.