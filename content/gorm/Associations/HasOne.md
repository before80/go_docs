+++
title = "一对一"
date = 2023-10-28T14:28:02+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/has_one.html](https://gorm.io/docs/has_one.html)

## 一对一 Has One

A `has one` association sets up a one-to-one connection with another model, but with somewhat different semantics (and consequences). This association indicates that each instance of a model contains or possesses one instance of another model.

​	一个`has one`关联设置了一个与另一个模型的一对一连接，但具有一些不同的语义（和后果）。这个关联表示每个模型实例包含或拥有另一个模型的一个实例。

For example, if your application includes users and credit cards, and each user can only have one credit card.

​	例如，如果你的应用包括用户和信用卡，并且每个用户可以只有一个信用卡。

### 声明 Declare

``` go
// User有一个CreditCard，UserID是外键 User has one CreditCard, UserID is the foreign key
type User struct {
  gorm.Model
  CreditCard CreditCard
}

type CreditCard struct {
  gorm.Model
  Number string
  UserID uint
}
```

### 检索 Retrieve

``` go
// 使用预加载获取用户列表及其信用卡 Retrieve user list with eager loading credit card
func GetAll(db *gorm.DB) ([]User, error) {
  var users []User
  err := db.Model(&User{}).Preload("CreditCard").Find(&users).Error
  return users, err
}
```

## 覆盖外键 Override Foreign Key

For a `has one` relationship, a foreign key field must also exist, the owner will save the primary key of the model belongs to it into this field.

​	对于`has one`关系，外键字段也必须存在，所有者将模型所属的主键保存到该字段中。

The field’s name is usually generated with `has one` model’s type plus its `primary key`, for the above example it is `UserID`.

​	字段的名称通常是`has one`模型的类型加上其主键，例如上面的示例中的`UserID`。

When you give a credit card to the user, it will save the User’s `ID` into its `UserID` field.

​	当你给用户一张信用卡时，它会将用户的`ID`保存到其`UserID`字段中。

If you want to use another field to save the relationship, you can change it with tag `foreignKey`, e.g:

​	如果你想使用另一个字段来保存关系，你可以使用标签`foreignKey`更改它，例如：

``` go
type User struct {
  gorm.Model
  CreditCard CreditCard `gorm:"foreignKey:UserName"`
  // 使用UserName作为外键 use UserName as foreign key
}

type CreditCard struct {
  gorm.Model
  Number   string
  UserName string
}
```

## 覆盖引用 Override References

By default, the owned entity will save the `has one` model’s primary key into a foreign key, you could change to save another field’s value, like using `Name` for the below example.

​	默认情况下，拥有的实体将`has one`模型的主键保存到一个外键中，你可以更改为保存另一个字段的值，例如使用下面的示例中的`Name`。

You are able to change it with tag `references`, e.g:

​	你可以使用标签`references`更改它，例如：

``` go
type User struct {
  gorm.Model
  Name       string     `gorm:"index"`
  CreditCard CreditCard `gorm:"foreignKey:UserName;references:name"`
}

type CreditCard struct {
  gorm.Model
  Number   string
  UserName string
}
```

## 多态关联 Polymorphism Association

GORM supports polymorphism association for `has one` and `has many`, it will save owned entity’s table name into polymorphic type’s field, primary key into the polymorphic field

​	GORM支持`has one`和`has many`的多态关联，它将拥有的实体的表名保存到多态类型的字段中，主键保存到多态字段中。

``` go
type Cat struct {
  ID    int
  Name  string
  Toy   Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
  ID   int
  Name string
  Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
  ID        int
  Name      string
  OwnerID   int
  OwnerType string
}

db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
// INSERT INTO `dogs` (`name`) VALUES ("dog1")
// INSERT INTO `toys` (`name`,`owner_id`,`owner_type`) VALUES ("toy1","1","dogs")
```

You can change the polymorphic type value with tag `polymorphicValue`, for example:

​	你可以使用标签`polymorphicValue`更改多态类型值，例如：

``` go
type Dog struct {
  ID   int
  Name string
  Toy  Toy `gorm:"polymorphic:Owner;polymorphicValue:master"`
}

type Toy struct {
  ID        int
  Name      string
  OwnerID   int
  OwnerType string
}

db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
// INSERT INTO `dogs` (`name`) VALUES ("dog1")
// INSERT INTO `toys` (`name`,`owner_id`,`owner_type`) VALUES ("toy1","1","master")
```

## CRUD与Has One的关系 CRUD with Has One

Please checkout [Association Mode](https://gorm.io/docs/associations.html#Association-Mode) for working with `has one` relations

​	请查看[关联模式](https://gorm.io/docs/associations.html#Association-Mode)以处理`has one`关系

## 预加载（Eager Loading） Eager Loading

GORM allows eager loading `has one` associations with `Preload` or `Joins`, refer [Preloading (Eager loading)](https://gorm.io/docs/preload.html) for details

​	GORM允许使用`Preload`或`Joins`进行预加载`has one`关联，详情请参考[预加载（Eager loading）](https://gorm.io/docs/preload.html)。

## Self-Referential Has One

``` go
type User struct {
  gorm.Model
  Name      string
  ManagerID *uint
  Manager   *User
}
```

## FOREIGN KEY Constraints

You can setup `OnUpdate`, `OnDelete` constraints with tag `constraint`, it will be created when migrating with GORM, for example:

​	你可以使用标签`constraint`设置`OnUpdate`、`OnDelete`约束，当使用GORM进行迁移时，它将被创建，例如：

``` go
type User struct {
  gorm.Model
  CreditCard CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreditCard struct {
  gorm.Model
  Number string
  UserID uint
}
```

You are also allowed to delete selected has one associations with `Select` when deleting, checkout [Delete with Select](https://gorm.io/docs/associations.html#delete_with_select) for details

​	你还可以使用`Select`在删除时选择删除特定的has one关联，详情请参考[Delete with Select](https://gorm.io/docs/associations.html#delete_with_select)。