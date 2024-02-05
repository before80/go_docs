+++
title = "复合主键"
date = 2023-10-28T14:36:19+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/composite_primary_key.html](https://gorm.io/docs/composite_primary_key.html)

Set multiple fields as primary key creates composite primary key, for example:

​	将多个字段设置为主键，创建复合主键，例如：

``` go
type Product struct {
  ID           string `gorm:"primaryKey"`
  LanguageCode string `gorm:"primaryKey"`
  Code         string
  Name         string
}
```

> **Note** integer `PrioritizedPrimaryField` enables `AutoIncrement` by default, to disable it, you need to turn off `autoIncrement` for the int fields:
>
> **注意**：整数`PrioritizedPrimaryField`默认启用`AutoIncrement`，要禁用它，您需要为int字段关闭`autoIncrement`：

``` go
type Product struct {
  CategoryID uint64 `gorm:"primaryKey;autoIncrement:false"`
  TypeID     uint64 `gorm:"primaryKey;autoIncrement:false"`
}
```