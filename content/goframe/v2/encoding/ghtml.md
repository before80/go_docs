+++
title = "ghtml"
date = 2024-03-21T17:49:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/ghtml](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/ghtml)

Package ghtml provides useful API for HTML content handling.

​	软件包 ghtml 为 HTML 内容处理提供了有用的 API。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Entities

```go
func Entities(s string) string
```

Entities encodes all HTML chars for content. Referer: http://php.net/manual/zh/function.htmlentities.php

​	实体对内容的所有 HTML 字符进行编码。推荐人： http://php.net/manual/zh/function.htmlentities.php

#### func EntitiesDecode

```go
func EntitiesDecode(s string) string
```

EntitiesDecode decodes all HTML chars for content. Referer: http://php.net/manual/zh/function.html-entity-decode.php

​	EntitiesDecode 对内容的所有 HTML 字符进行解码。推荐人： http://php.net/manual/zh/function.html-entity-decode.php

#### func SpecialChars

```go
func SpecialChars(s string) string
```

SpecialChars encodes some special chars for content, these special chars are: “&”, “<”, “>”, `"`, “’”. Referer: http://php.net/manual/zh/function.htmlspecialchars.php

​	SpecialChars 对内容的一些特殊字符进行编码，这些特殊字符是：“&”、“<”、“>”、“ `"` ”、“'”。推荐人：http://php.net/manual/zh/function.htmlspecialchars.php

#### func SpecialCharsDecode

```go
func SpecialCharsDecode(s string) string
```

SpecialCharsDecode decodes some special chars for content, these special chars are: “&”, “<”, “>”, `"`, “’”. Referer: http://php.net/manual/zh/function.htmlspecialchars-decode.php

​	SpecialCharsDecode 对内容的一些特殊字符进行解码，这些特殊字符是：“&”、“<”、“>”、“ `"` ”、“'”。推荐人： http://php.net/manual/zh/function.htmlspecialchars-decode.php

#### func SpecialCharsMapOrStruct

```go
func SpecialCharsMapOrStruct(mapOrStruct interface{}) error
```

SpecialCharsMapOrStruct automatically encodes string values/attributes for map/struct.

​	SpecialCharsMapOrStruct 自动对 map/struct 的字符串值/属性进行编码。

#### func StripTags

```go
func StripTags(s string) string
```

StripTags strips HTML tags from content, and returns only text. Referer: http://php.net/manual/zh/function.strip-tags.php

​	StripTags 从内容中剥离 HTML 标记，并仅返回文本。推荐人： http://php.net/manual/zh/function.strip-tags.php

## 类型

This section is empty.