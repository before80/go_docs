+++
title = "设置"
date = 2023-10-28T14:32:46+08:00
weight = 14
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/settings.html](https://gorm.io/docs/settings.html)

GORM provides `Set`, `Get`, `InstanceSet`, `InstanceGet` methods allow users pass values to [hooks](https://gorm.io/docs/hooks.html) or other methods

​	GORM提供了`Set`、`Get`、`InstanceSet`和`InstanceGet`方法，允许用户向[钩子](../hooks)（hooks）或其他方法传递值。

GORM uses this for some features, like pass creating table options when migrating table.

​	GORM在迁移表时使用这些功能，例如传递创建表选项。

``` go
// 在创建表时添加表后缀 Add table suffix when creating tables
db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
```

## Set / Get

Use `Set` / `Get` pass settings to hooks methods, for example:

​	使用`Set` / `Get`将设置传递给钩子方法，例如：

``` go
type User struct {
  gorm.Model
  CreditCard CreditCard
  // ...
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
  myValue, ok := tx.Get("my_value")
  // ok => true
  // myValue => 123
}

type CreditCard struct {
  gorm.Model
  // ...
}

func (card *CreditCard) BeforeCreate(tx *gorm.DB) error {
  myValue, ok := tx.Get("my_value")
  // ok => true
  // myValue => 123
}

myValue := 123
db.Set("my_value", myValue).Create(&User{})
```

## InstanceSet / InstanceGet

Use `InstanceSet` / `InstanceGet` pass settings to current `*Statement`‘s hooks methods, for example:

​	使用`InstanceSet` / `InstanceGet`将设置传递给当前`*Statement`的钩子方法，例如：

``` go
type User struct {
  gorm.Model
  CreditCard CreditCard
  // ...
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
  myValue, ok := tx.InstanceGet("my_value")
  // ok => true
  // myValue => 123
}

type CreditCard struct {
  gorm.Model
  // ...
}

// 当创建关联时，GORM会创建一个新`*Statement`，因此无法读取其他实例的设置 When creating associations, GORM creates a new `*Statement`, so can't read other instance's settings
func (card *CreditCard) BeforeCreate(tx *gorm.DB) error {
  myValue, ok := tx.InstanceGet("my_value")
  // ok => false
  // myValue => nil
}

myValue := 123
db.InstanceSet("my_value", myValue).Create(&User{})
```