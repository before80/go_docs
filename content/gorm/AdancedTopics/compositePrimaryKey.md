+++
title = "Composite Primary Key"
date = 2023-10-28T14:36:19+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/composite_primary_key.html](https://gorm.io/docs/composite_primary_key.html)

Set multiple fields as primary key creates composite primary key, for example:

``` go
type Product struct {
  ID           string `gorm:"primaryKey"`
  LanguageCode string `gorm:"primaryKey"`
  Code         string
  Name         string
}
```

**Note** integer `PrioritizedPrimaryField` enables `AutoIncrement` by default, to disable it, you need to turn off `autoIncrement` for the int fields:

``` go
type Product struct {
  CategoryID uint64 `gorm:"primaryKey;autoIncrement:false"`
  TypeID     uint64 `gorm:"primaryKey;autoIncrement:false"`
}
```