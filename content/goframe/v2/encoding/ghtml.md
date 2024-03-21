+++
title = "ghtml"
date = 2024-03-21T17:49:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/ghtml

Package ghtml provides useful API for HTML content handling.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Entities 

``` go
func Entities(s string) string
```

Entities encodes all HTML chars for content. Referer: http://php.net/manual/zh/function.htmlentities.php

##### func EntitiesDecode 

``` go
func EntitiesDecode(s string) string
```

EntitiesDecode decodes all HTML chars for content. Referer: http://php.net/manual/zh/function.html-entity-decode.php

##### func SpecialChars 

``` go
func SpecialChars(s string) string
```

SpecialChars encodes some special chars for content, these special chars are: "&", "<", ">", `"`, "'". Referer: http://php.net/manual/zh/function.htmlspecialchars.php

##### func SpecialCharsDecode 

``` go
func SpecialCharsDecode(s string) string
```

SpecialCharsDecode decodes some special chars for content, these special chars are: "&", "<", ">", `"`, "'". Referer: http://php.net/manual/zh/function.htmlspecialchars-decode.php

##### func SpecialCharsMapOrStruct 

``` go
func SpecialCharsMapOrStruct(mapOrStruct interface{}) error
```

SpecialCharsMapOrStruct automatically encodes string values/attributes for map/struct.

##### func StripTags 

``` go
func StripTags(s string) string
```

StripTags strips HTML tags from content, and returns only text. Referer: http://php.net/manual/zh/function.strip-tags.php

### Types 

This section is empty.