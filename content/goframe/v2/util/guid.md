+++
title = "guid"
date = 2024-03-21T17:59:59+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/guid

Package guid provides simple and high performance unique id generation functionality.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func S 

``` go
func S(data ...[]byte) string
```

S creates and returns a global unique string in 32 bytes that meets most common usages without strict UUID algorithm. It returns a unique string using default unique algorithm if no `data` is given.

The specified `data` can be no more than 2 parts. No matter how long each of the `data` size is, each of them will be hashed into 7 bytes as part of the result. If given `data` parts is less than 2, the leftover size of the result bytes will be token by random string.

The returned string is composed with: 1. Default: MACHash(7) + PID(4) + TimestampNano(12) + Sequence(3) + RandomString(6) 2. CustomData: DataHash(7/14) + TimestampNano(12) + Sequence(3) + RandomString(3/10)

Note that：

1. The returned length is fixed to 32 bytes for performance purpose.
2. The custom parameter `data` composed should have unique attribute in your business scenario.

### Types 

This section is empty.