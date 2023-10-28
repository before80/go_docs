+++
title = "Constraints"
date = 2023-10-28T14:35:55+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/constraints.html](https://gorm.io/docs/constraints.html)

GORM allows create database constraints with tag, constraints will be created when [AutoMigrate or CreateTable with GORM](https://gorm.io/docs/migration.html)

## CHECK Constraint

Create CHECK constraints with tag `check`

```
type UserIndex struct {
  Name  string `gorm:"check:name_checker,name <> 'jinzhu'"`
  Name2 string `gorm:"check:name <> 'jinzhu'"`
  Name3 string `gorm:"check:,name <> 'jinzhu'"`
}
```

## Index Constraint

Checkout [Database Indexes](https://gorm.io/docs/indexes.html)

## Foreign Key Constraint

GORM will creates foreign keys constraints for associations, you can disable this feature during initialization:

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  DisableForeignKeyConstraintWhenMigrating: true,
})
```

GORM allows you setup FOREIGN KEY constraints’s `OnDelete`, `OnUpdate` option with tag `constraint`, for example:

```
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