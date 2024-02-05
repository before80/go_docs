+++
title = "属于关系"
date = 2023-10-28T14:27:48+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/belongs_to.html](https://gorm.io/docs/belongs_to.html)

## 属于关系 Belongs To

A `belongs to` association sets up a one-to-one connection with another model, such that each instance of the declaring model “belongs to” one instance of the other model.

​	一个`belongs to`关联设置了一个模型与另一个模型之间的一对一连接，使得声明模型的每个实例“属于”另一个模型的一个实例。

For example, if your application includes users and companies, and each user can be assigned to exactly one company, the following types represent that relationship. Notice here that, on the `User` object, there is both a `CompanyID` as well as a `Company`. By default, the `CompanyID` is implicitly used to create a foreign key relationship between the `User` and `Company` tables, and thus must be included in the `User` struct in order to fill the `Company` inner struct.

​	例如，如果你的应用包括用户和公司，并且每个用户可以被分配给一个公司，那么以下类型表示这种关系。请注意，在`User`对象上，有一个`CompanyID`以及一个`Company`。默认情况下，`CompanyID`隐式地用于创建`User`和`Company`表之间的外键关系，因此必须包含在`User`结构体中才能填充`Company`内部结构体。

``` go
// `User`属于`Company`，`CompanyID`是外键 `User` belongs to `Company`, `CompanyID` is the foreign key
type User struct {
  gorm.Model
  Name      string
  CompanyID int
  Company   Company
}

type Company struct {
  ID   int
  Name string
}
```

Refer to [Eager Loading](https://gorm.io/docs/belongs_to.html#Eager-Loading) for details on populating the inner struct.

​	有关填充内部结构的详细信息，请参阅[预加载（预加载）](https://gorm.io/docs/preload.html)。

## 覆盖外键 Override Foreign Key

To define a belongs to relationship, the foreign key must exist, the default foreign key uses the owner’s type name plus its primary field name.

​	要定义一个属于关系，外键必须存在，默认的外键使用拥有者的类型名称加上其主字段名称。

For the above example, to define the `User` model that belongs to `Company`, the foreign key should be `CompanyID` by convention

​	要定义一个属于关系，外键必须存在，默认的外键使用拥有者的类型名称加上其主字段名称。

GORM provides a way to customize the foreign key, for example:

​	GORM提供了一个自定义外键的方法，例如：

``` go
type User struct {
  gorm.Model
  Name         string
  CompanyRefer int
  Company      Company `gorm:"foreignKey:CompanyRefer"`
  // 使用CompanyRefer作为外键 use CompanyRefer as foreign key
}

type Company struct {
  ID   int
  Name string
}
```

## 覆盖引用 Override References

For a belongs to relationship, GORM usually uses the owner’s primary field as the foreign key’s value, for the above example, it is `Company`‘s field `ID`.

​	对于一个属于关系，GORM通常使用拥有者的主字段作为外键的值，对于上面的示例，它是`Company`的字段`ID`。

When you assign a user to a company, GORM will save the company’s `ID` into the user’s `CompanyID` field.

​	当你将用户分配给一家公司时，GORM将公司的`ID`保存到用户的`CompanyID`字段中。

You are able to change it with tag `references`, e.g:

​	你可以使用标签`references`更改它，例如：

``` go
type User struct {
  gorm.Model
  Name      string
  CompanyID string
  Company   Company `gorm:"references:Code"` // use Code as references
}

type Company struct {
  ID   int
  Code string
  Name string
}
```

> **NOTE** GORM usually guess the relationship as `has one` if override foreign key name already exists in owner’s type, we need to specify `references` in the `belongs to` relationship.
>
> **注意** GORM通常在拥有者的类型中已经存在覆盖外键名称的关系时猜测关系为`has one`，我们需要在`belongs to`关系中指定`references`。

``` go
type User struct {
  gorm.Model
  Name      string
  CompanyID string
  Company   Company `gorm:"references:CompanyID"` // 使用Company.CompanyID作为引用 use Company.CompanyID as references
}

type Company struct {
  CompanyID   int
  Code        string
  Name        string
}
```

## 与属于关系的CRUD操作 CRUD with Belongs To

Please checkout [Association Mode](https://gorm.io/docs/associations.html#Association-Mode) for working with belongs to relations

​	请查看[关联模式](https://gorm.io/docs/associations.html#Association-Mode)以处理属于关系

## 预加载（预加载） Eager Loading

GORM allows eager loading belongs to associations with `Preload` or `Joins`, refer [Preloading (Eager loading)](https://gorm.io/docs/preload.html) for details

​	GORM允许通过`Preload`或`Joins`预加载属于关系，请参阅[预加载（预加载）](https://gorm.io/docs/preload.html)以获取详细信息

## FOREIGN KEY约束 FOREIGN KEY Constraints

You can setup `OnUpdate`, `OnDelete` constraints with tag `constraint`, it will be created when migrating with GORM, for example:

​	你可以使用带有标签`constraint`的`OnUpdate`和`OnDelete`约束来设置`OnUpdate`、`OnDelete`约束，当使用GORM进行迁移时，它将被创建，例如：

``` go
type User struct {
  gorm.Model
  Name      string
  CompanyID int
  Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Company struct {
  ID   int
  Name string
}
```