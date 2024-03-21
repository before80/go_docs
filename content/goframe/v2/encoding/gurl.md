+++
title = "gurl"
date = 2024-03-21T17:50:13+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gurl

Package gurl provides useful API for URL handling.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func BuildQuery 

``` go
func BuildQuery(queryData url.Values) string
```

BuildQuery Generate URL-encoded query string. See http://php.net/manual/en/function.http-build-query.php.

##### func Decode 

``` go
func Decode(str string) (string, error)
```

Decode does the inverse transformation of Encode, converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits.

##### func Encode 

``` go
func Encode(str string) string
```

Encode escapes the string so it can be safely placed inside an URL query.

##### func ParseURL 

``` go
func ParseURL(str string, component int) (map[string]string, error)
```

ParseURL Parse an URL and return its components. -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment. See http://php.net/manual/en/function.parse-url.php.

##### func RawDecode 

``` go
func RawDecode(str string) (string, error)
```

RawDecode does decode the given string Decode URL-encoded strings. See http://php.net/manual/en/function.rawurldecode.php.

##### func RawEncode 

``` go
func RawEncode(str string) string
```

RawEncode does encode the given string according URL-encode according to [RFC 3986](https://rfc-editor.org/rfc/rfc3986.html). See http://php.net/manual/en/function.rawurlencode.php.

### Types 

This section is empty.