+++
title = "约束"
date = 2023-10-28T14:35:55+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/constraints.html](https://gorm.io/docs/constraints.html)

GORM allows create database constraints with tag, constraints will be created when [AutoMigrate or CreateTable with GORM](https://gorm.io/docs/migration.html)

​	GORM 允许使用标签创建数据库约束，当执行 [AutoMigrate 或 CreateTable 与 GORM](https://gorm.io/docs/migration.html) 时，将创建这些约束。

## CHECK Constraint

Create CHECK constraints with tag `check`

​	使用 `check` 标签创建 CHECK 约束：

``` go
type UserIndex struct {
  Name  string `gorm:"check:name_checker,name <> 'jinzhu'"`
  Name2 string `gorm:"check:name <> 'jinzhu'"`
  Name3 string `gorm:"check:,name <> 'jinzhu'"`
}
```

## Index Constraint

Checkout [Database Indexes](https://gorm.io/docs/indexes.html)

​	查看 [数据库索引](https://gorm.io/docs/indexes.html)。

## Foreign Key Constraint

GORM will creates foreign keys constraints for associations, you can disable this feature during initialization:

​	GORM 将为关联关系创建外键约束，您可以在初始化过程中禁用此功能：

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  DisableForeignKeyConstraintWhenMigrating: true,
})
```

GORM allows you setup FOREIGN KEY constraints’s `OnDelete`, `OnUpdate` option with tag `constraint`, for example:

​	GORM 允许您使用 `constraint` 标签的 `OnDelete`、`OnUpdate` 选项设置 FOREIGN KEY 约束，例如：

``` go
type User struct {
  gorm.Model
  CompanyID  int
  Company    Company    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
  CreditCard CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreditCard struct {
  gorm.Model
  Number string
  UserID uint
}

type Company struct {
  ID   int
  Name string
}
```

