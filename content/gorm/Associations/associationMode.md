+++
title = "关联模式"
date = 2023-10-28T14:28:42+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/associations.html](https://gorm.io/docs/associations.html)

## 自动创建/更新 Auto Create/Update

GORM will auto-save associations and its reference using [Upsert](https://gorm.io/docs/create.html#upsert) when creating/updating a record.

​	GORM 在创建或更新记录时会自动保存关联及其引用，使用 Upsert（如果存在则更新，否则插入）功能。

``` go
user := User{
  Name:            "jinzhu",
  BillingAddress:  Address{Address1: "Billing Address - Address 1"},
  ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
  Emails:          []Email{
    {Email: "jinzhu@example.com"},
    {Email: "jinzhu-2@example.com"},
  },
  Languages:       []Language{
    {Name: "ZH"},
    {Name: "EN"},
  },
}

db.Create(&user)
// BEGIN TRANSACTION;
// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY DO NOTHING;
// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY DO NOTHING;
// INSERT INTO "languages" ("name") VALUES ('ZH'), ('EN') ON DUPLICATE KEY DO NOTHING;
// INSERT INTO "user_languages" ("user_id","language_id") VALUES (111, 1), (111, 2) ON DUPLICATE KEY DO NOTHING;
// COMMIT;

db.Save(&user)
```

If you want to update associations’s data, you should use the `FullSaveAssociations` mode:

​	如果你想更新关联的数据，你应该使用 `FullSaveAssociations` 模式：

``` go
db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
// ...
// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY SET address1=VALUES(address1);
// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY SET email=VALUES(email);
// ...
```

## 跳过自动创建/更新 Skip Auto Create/Update

To skip the auto save when creating/updating, you can use `Select` or `Omit`, for example:

​	为了在创建/更新时跳过自动保存，你可以使用 `Select` 或 `Omit`，例如：

``` go
user := User{
  Name:            "jinzhu",
  BillingAddress:  Address{Address1: "Billing Address - Address 1"},
  ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
  Emails:          []Email{
    {Email: "jinzhu@example.com"},
    {Email: "jinzhu-2@example.com"},
  },
  Languages:       []Language{
    {Name: "ZH"},
    {Name: "EN"},
  },
}

db.Select("Name").Create(&user)
// INSERT INTO "users" (name) VALUES ("jinzhu", 1, 2);

db.Omit("BillingAddress").Create(&user)
// Skip create BillingAddress when creating a user

db.Omit(clause.Associations).Create(&user)
// Skip all associations when creating a user
```

> **NOTE:**
> For many2many associations, GORM will upsert the associations before creating the join table references, if you want to skip the upserting of associations, you could skip it like:
>
> **注意：** 对于 many2many 关联，GORM 将在创建 join 表引用之前插入关联，如果你想要跳过关联的插入，你可以像这样跳过它：
>
> ```go
> db.Omit("Languages.*").Create(&user)
> ```
>
> The following code will skip the creation of the association and its references
>
> ​	下面的代码将跳过创建关联及其引用
>
> ```go
> db.Omit("Languages").Create(&user)
> ```

## Select/Omit Association fields

``` go
user := User{
  Name:            "jinzhu",
  BillingAddress:  Address{Address1: "Billing Address - Address 1", Address2: "addr2"},
  ShippingAddress: Address{Address1: "Shipping Address - Address 1", Address2: "addr2"},
}

// Create user and his BillingAddress, ShippingAddress
// When creating the BillingAddress only use its address1, address2 fields and omit others
db.Select("BillingAddress.Address1", "BillingAddress.Address2").Create(&user)

db.Omit("BillingAddress.Address2", "BillingAddress.CreatedAt").Create(&user)
```

## Association Mode

Association Mode contains some commonly used helper methods to handle relationships

``` go
// Start Association Mode
var user User
db.Model(&user).Association("Languages")
// `user` is the source model, it must contains primary key
// `Languages` is a relationship's field name
// If the above two requirements matched, the AssociationMode should be started successfully, or it should return error
db.Model(&user).Association("Languages").Error
```

### Find Associations

Find matched associations

``` go
db.Model(&user).Association("Languages").Find(&languages)
```

Find associations with conditions

``` go
codes := []string{"zh-CN", "en-US", "ja-JP"}
db.Model(&user).Where("code IN ?", codes).Association("Languages").Find(&languages)

db.Model(&user).Where("code IN ?", codes).Order("code desc").Association("Languages").Find(&languages)
```

### Append Associations

Append new associations for `many to many`, `has many`, replace current association for `has one`, `belongs to`

``` go
db.Model(&user).Association("Languages").Append([]Language{languageZH, languageEN})

db.Model(&user).Association("Languages").Append(&Language{Name: "DE"})

db.Model(&user).Association("CreditCard").Append(&CreditCard{Number: "411111111111"})
```

### Replace Associations

Replace current associations with new ones

​	将当前关联替换为新关联

``` go
db.Model(&user).Association("Languages").Replace([]Language{languageZH, languageEN})

db.Model(&user).Association("Languages").Replace(Language{Name: "DE"}, languageEN)
```

### 删除关联 Delete Associations

Remove the relationship between source & arguments if exists, only delete the reference, won’t delete those objects from DB.

​	如果存在，则删除源和参数之间的关系，只删除引用，不会从数据库中删除这些对象。

``` go
db.Model(&user).Association("Languages").Delete([]Language{languageZH, languageEN})
db.Model(&user).Association("Languages").Delete(languageZH, languageEN)
```

### 清除关联 Clear Associations

Remove all reference between source & association, won’t delete those associations

​	删除源和关联之间的所有引用，不会删除这些关联

``` go
db.Model(&user).Association("Languages").Clear()
```

### 计数关联 Count Associations

Return the count of current associations

​	返回当前关联的计数

``` go
db.Model(&user).Association("Languages").Count()

// 带条件计数 Count with conditions
codes := []string{"zh-CN", "en-US", "ja-JP"}
db.Model(&user).Where("code IN ?", codes).Association("Languages").Count()
```

### 批量数据 Batch Data

Association Mode supports batch data, e.g:

​	关联模式支持批量数据，例如：

``` go
// 查找所有用户的所有角色 Find all roles for all users
db.Model(&users).Association("Role").Find(&roles)

// 从所有用户的团队中删除用户A Delete User A from all user's team
db.Model(&users).Association("Team").Delete(&userA)

// 获取所有用户团队的不同计数 Get distinct count of all users' teams
db.Model(&users).Association("Team").Count()

// 对于`Append`、`Replace`，使用批量数据时，参数的长度需要等于数据的长度或否则会返回错误 For `Append`, `Replace` with batch data, the length of the arguments needs to be equal to the data's length or else it will return an error
var users = []User{user1, user2, user3}
// 例如：我们有3个用户，将userA添加到user1的团队，将userB添加到user2的团队，将userA、userB和userC添加到user3的团队 e.g: we have 3 users, Append userA to user1's team, append userB to user2's team, append userA, userB and userC to user3's team
db.Model(&users).Association("Team").Append(&userA, &userB, &[]User{userA, userB, userC})
// 重置user1的团队为userA，重置user2的团队为userB，重置user3的团队为userA、userB和userC Reset user1's team to userA，reset user2's team to userB, reset user3's team to userA, userB and userC
db.Model(&users).Association("Team").Replace(&userA, &userB, &[]User{userA, userB, userC})
```

## 删除关联记录 Delete Association Record

By default, `Replace`/`Delete`/`Clear` in `gorm.Association` only delete the reference,
that is, set old associations’s foreign key to null.

​	默认情况下，`gorm.Association`中的`Replace`/`Delete`/`Clear`仅删除引用，即将旧关联的外键设置为null。

You can delete those objects with `Unscoped` (it has nothing to do with `ManyToMany`).

​	你可以使用`Unscoped`（它与`ManyToMany`无关）来删除它们。

How to delete is decided by `gorm.DB`.

​	如何删除由`gorm.DB`决定。

``` go
// 软删除 Soft delete
// UPDATE `languages` SET `deleted_at`= ...
db.Model(&user).Association("Languages").Unscoped().Clear()

// 永久删除 Delete permanently
// DELETE FROM `languages` WHERE ...
db.Unscoped().Model(&item).Association("Languages").Unscoped().Clear()
```

## 使用Select删除 Delete with Select

You are allowed to delete selected has one/has many/many2many relations with `Select` when deleting records, for example:

​	在删除记录时，可以使用`Select`来删除选定的一对一/一对多/多对多关系，例如：

``` go
// 在删除用户时删除用户的帐户 delete user's account when deleting user
db.Select("Account").Delete(&user)

// 在删除用户时删除用户的Orders、CreditCards关系 delete user's Orders, CreditCards relations when deleting user
db.Select("Orders", "CreditCards").Delete(&user)

// 在删除用户时删除用户的一对一/一对多/多对多关系 delete user's has one/many/many2many relations when deleting user
db.Select(clause.Associations).Delete(&user)

// 在删除用户时删除每个用户的帐户 delete each user's account when deleting users
db.Select("Account").Delete(&users)
```

> **NOTE:**
> Associations will only be deleted if the deleting records’s primary key is not zero, GORM will use those primary keys as conditions to delete selected associations
>
> **注意：** 只有在删除记录的主键不为零时，关联才会被删除。GORM将使用这些主键作为条件来删除选定的关联
>
> ``` go
> // 不起作用 DOESN'T WORK
> db.Select("Account").Where("name = ?", "jinzhu").Delete(&User{})
> // 将删除所有名为`jinzhu`的用户，但这些用户的帐户不会被删除 will delete all user with name `jinzhu`, but those user's account won't be deleted
> 
> db.Select("Account").Where("name = ?", "jinzhu").Delete(&User{ID: 1})
> // 将删除名为`jinzhu`且id为`1`的用户，以及用户`1`的帐户将被删除 will delete the user with name = `jinzhu` and id = `1`, and user `1`'s account will be deleted
> 
> db.Select("Account").Delete(&User{ID: 1})
> // 将删除id为`1`的用户，以及用户`1`的帐户将被删除 will delete the user with id = `1`, and user `1`'s account will be deleted
> ```

## 关联标签 Association Tags

| Tag              | Description                                                  |
| :--------------- | :----------------------------------------------------------- |
| foreignKey       | 指定当前模型用作联接表外键的列名<br />Specifies column name of the current model that is used as a foreign key to the join table |
| references       | 指定引用表的列名，该列映射到联接表的外键<br />Specifies column name of the reference’s table that is mapped to the foreign key of the join table |
| polymorphic      | 指定多态类型，如模型名称<br />Specifies polymorphic type such as model name |
| polymorphicValue | 指定多态值，默认表名<br />Specifies polymorphic value, default table name |
| many2many        | 指定联接表名<br />Specifies join table name                  |
| joinForeignKey   | 指定联接表中映射到当前表的外键列名<br />Specifies foreign key column name of join table that maps to the current table |
| joinReferences   | 指定联接表中映射到引用表的外键列名<br />Specifies foreign key column name of join table that maps to the reference’s table |
| constraint       | 关系约束，例如：`OnUpdate`、`OnDelete`<br />Relations constraint, e.g: `OnUpdate`,`OnDelete` |

