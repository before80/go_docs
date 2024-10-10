+++
title = "多对一"
date = 2023-10-28T14:28:11+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/has_many.html](https://gorm.io/docs/has_many.html)

## 多对一 Has Many

A `has many` association sets up a one-to-many connection with another model, unlike `has one`, the owner could have zero or many instances of models.

​	`has many`关联设置了一个一对多的连接，与`has one`不同，拥有者可以有零个或多个模型实例。

For example, if your application includes users and credit card, and each user can have many credit cards.

​	`has many`关联设置了一个一对多的连接，与`has one`不同，拥有者可以有零个或多个模型实例。

### 声明 Declare

``` go
// User有多个CreditCards，UserID是外键 User has many CreditCards, UserID is the foreign key
type User struct {
  gorm.Model
  CreditCards []CreditCard
}

type CreditCard struct {
  gorm.Model
  Number string
  UserID uint
}
```

### 检索 Retrieve

``` go
// 使用预加载（Eager loading）检索用户列表及其信用卡 Retrieve user list with eager loading credit cards
func GetAll(db *gorm.DB) ([]User, error) {
    var users []User
    err := db.Model(&User{}).Preload("CreditCards").Find(&users).Error
    return users, err
}
```

## 覆盖外键 Override Foreign Key

To define a `has many` relationship, a foreign key must exist. The default foreign key’s name is the owner’s type name plus the name of its primary key field

​	要定义一个`has many`关系，必须存在一个外键。默认的外键名称是拥有者的类型的名称加上其主键字段的名称。

For example, to define a model that belongs to `User`, the foreign key should be `UserID`.

​	例如，要定义一个属于`User`的模型，外键应该是`UserID`。

To use another field as foreign key, you can customize it with a `foreignKey` tag, e.g:

​	要使用另一个字段作为外键，可以使用带有`foreignKey`标签的自定义它，例如：

``` go
type User struct {
  gorm.Model
  CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
  gorm.Model
  Number    string
  UserRefer uint
}
```

## 覆盖引用 Override References

GORM usually uses the owner’s primary key as the foreign key’s value, for the above example, it is the `User`‘s `ID`,

When you assign credit cards to a user, GORM will save the user’s `ID` into credit cards’ `UserID` field.

​	GORM通常使用拥有者的主键作为外键的值，对于上面的示例，它是`User`的`ID`。

​	当你将信用卡分配给用户时，GORM会将用户的`ID`保存到信用卡的`UserID`字段中。

You are able to change it with tag `references`, e.g:

​	你可以使用带有`references`标签更改它，例如：

``` go
type User struct {
  gorm.Model
  MemberNumber string
  CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
}

type CreditCard struct {
  gorm.Model
  Number     string
  UserNumber string
}
```

## 多态关联 Polymorphism Association

GORM supports polymorphism association for `has one` and `has many`, it will save owned entity’s table name into polymorphic type’s field, primary key value into the polymorphic field

​	GORM支持`has one`和`has many`的多态关联，它将拥有实体的表名保存到多态类型字段中，并将主键值保存到多态字段中。

``` go
type Dog struct {
  ID   int
  Name string
  Toys []Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
  ID        int
  Name      string
  OwnerID   int
  OwnerType string
}

db.Create(&Dog{Name: "dog1", Toys: []Toy{{Name: "toy1"}, {Name: "toy2"}}})
// INSERT INTO `dogs` (`name`) VALUES ("dog1")
// INSERT INTO `toys` (`name`,`owner_id`,`owner_type`) VALUES ("toy1","1","dogs"), ("toy2","1","dogs")
```

You can change the polymorphic type value with tag `polymorphicValue`, for example:

​	你可以使用带有`polymorphicValue`标签更改多态类型值，例如：

``` go
type Dog struct {
  ID   int
  Name string
  Toys []Toy `gorm:"polymorphic:Owner;polymorphicValue:master"`
}

type Toy struct {
  ID        int
  Name      string
  OwnerID   int
  OwnerType string
}

db.Create(&Dog{Name: "dog1", Toys: []Toy{{Name: "toy1"}, {Name: "toy2"}}})
// INSERT INTO `dogs` (`name`) VALUES ("dog1")
// INSERT INTO `toys` (`name`,`owner_id`,`owner_type`) VALUES ("toy1","1","master"), ("toy2","1","master")
```

## CRUD with Has Many

Please checkout [Association Mode](https://gorm.io/docs/associations.html#Association-Mode) for working with has many relations

​	请查看[关联模式](https://gorm.io/docs/associations.html#Association-Mode)以处理具有许多关系的CRUD操作。

## Eager Loading

GORM allows eager loading has many associations with `Preload`, refer [Preloading (Eager loading)](https://gorm.io/docs/preload.html) for details

​	GORM允许使用`Preload`进行具有许多关联的预加载，请参阅[预加载（Eager loading）](https://gorm.io/docs/preload.html)以获取详细信息。

## Self-Referential Has Many

``` go
type User struct {
  gorm.Model
  Name      string
  ManagerID *uint
  Team      []User `gorm:"foreignkey:ManagerID"`
}
```

## FOREIGN KEY Constraints

You can setup `OnUpdate`, `OnDelete` constraints with tag `constraint`, it will be created when migrating with GORM, for example:

​	你可以使用带有`OnUpdate`、`OnDelete`约束的`constraint`标签在迁移时创建它们，例如：

``` go
type User struct {
  gorm.Model
  CreditCards []CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreditCard struct {
  gorm.Model
  Number string
  UserID uint
}
```

You are also allowed to delete selected has many associations with `Select` when deleting, checkout [Delete with Select](https://gorm.io/docs/associations.html#delete_with_select) for details

​	你还可以在删除时选择删除选定的许多关联，请参阅[删除带有Select](https://gorm.io/docs/associations.html#delete_with_select)以获取详细信息。