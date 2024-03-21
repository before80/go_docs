+++
title = "gtoml"
date = 2024-03-21T17:50:06+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gtoml

Package gtoml provides accessing and converting for TOML content.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Decode 

``` go
func Decode(v []byte) (interface{}, error)
```

##### func DecodeTo 

``` go
func DecodeTo(v []byte, result interface{}) (err error)
```

##### func Encode 

``` go
func Encode(v interface{}) ([]byte, error)
```

##### func ToJson 

``` go
func ToJson(v []byte) ([]byte, error)
```

### Types 

This section is empty.