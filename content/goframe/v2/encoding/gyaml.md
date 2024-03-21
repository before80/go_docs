+++
title = "gyaml"
date = 2024-03-21T17:50:29+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gyaml

Package gyaml provides accessing and converting for YAML content.

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

##### func DecodeTo 

``` go
func DecodeTo(value []byte, result interface{}) (err error)
```

DecodeTo parses `content` into `result`.

##### func Encode 

``` go
func Encode(value interface{}) (out []byte, err error)
```

Encode encodes `value` to an YAML format content as bytes.

##### func EncodeIndent 

``` go
func EncodeIndent(value interface{}, indent string) (out []byte, err error)
```

EncodeIndent encodes `value` to an YAML format content with indent as bytes.

##### func ToJson 

``` go
func ToJson(content []byte) (out []byte, err error)
```

ToJson converts `content` to JSON format content.

### Types 

This section is empty.