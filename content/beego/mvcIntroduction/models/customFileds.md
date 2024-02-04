+++
title = "自定义字段"
date = 2024-02-04T10:02:31+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/model/custom_fields/]({{< ref "/beego/mvcIntroduction/models/customFileds" >}})

# Custom Fields 自定义字段



## Custom Fields 自定义字段

```go
TypeBooleanField = 1 << iota

// string
TypeCharField

// string
TypeTextField

// time.Time
TypeDateField
// time.Time
TypeDateTimeField

// int16
TypeSmallIntegerField
// int32
TypeIntegerField
// int64
TypeBigIntegerField
// uint16
TypePositiveSmallIntegerField
// uint32
TypePositiveIntegerField
// uint64
TypePositiveBigIntegerField

// float64
TypeFloatField
// float64
TypeDecimalField

RelForeignKey
RelOneToOne
RelManyToMany
RelReverseOne
RelReverseMany
```
