+++
title = "gproperties"
date = 2024-03-21T17:49:57+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gproperties

Package gproperties provides accessing and converting for .properties content.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Decode 

``` go
func Decode(data []byte) (res map[string]interface{}, err error)
```

Decode converts properties format to map.

##### func Encode 

``` go
func Encode(data map[string]interface{}) (res []byte, err error)
```

Encode converts map to properties format.

##### func ToJson 

``` go
func ToJson(data []byte) (res []byte, err error)
```

ToJson convert .properties format to JSON.

### Types 

This section is empty.