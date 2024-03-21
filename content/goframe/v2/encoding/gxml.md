+++
title = "gxml"
date = 2024-03-21T17:50:21+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gxml

Package gxml provides accessing and converting for XML content.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Decode 

``` go
func Decode(content []byte) (map[string]interface{}, error)
```

Decode parses `content` into and returns as map.

##### func DecodeWithoutRoot 

``` go
func DecodeWithoutRoot(content []byte) (map[string]interface{}, error)
```

DecodeWithoutRoot parses `content` into a map, and returns the map without root level.

##### func Encode 

``` go
func Encode(m map[string]interface{}, rootTag ...string) ([]byte, error)
```

Encode encodes map `m` to an XML format content as bytes. The optional parameter `rootTag` is used to specify the XML root tag.

##### func EncodeWithIndent 

``` go
func EncodeWithIndent(m map[string]interface{}, rootTag ...string) ([]byte, error)
```

EncodeWithIndent encodes map `m` to an XML format content as bytes with indent. The optional parameter `rootTag` is used to specify the XML root tag.

##### func ToJson 

``` go
func ToJson(content []byte) ([]byte, error)
```

ToJson converts `content` as XML format into JSON format bytes.

### Types 

This section is empty.