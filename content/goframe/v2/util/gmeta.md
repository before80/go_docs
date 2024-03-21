+++
title = "gmeta"
date = 2024-03-21T17:59:27+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gmeta

Package gmeta provides embedded meta data feature for struct.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Data 

``` go
func Data(object interface{}) map[string]string
```

Data retrieves and returns all metadata from `object`.

##### func Get 

``` go
func Get(object interface{}, key string) *gvar.Var
```

Get retrieves and returns specified metadata by `key` from `object`.

### Types 

#### type Meta 

``` go
type Meta struct{}
```

Meta is used as an embedded attribute for struct to enabled metadata feature.