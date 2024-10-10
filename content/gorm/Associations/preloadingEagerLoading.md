+++
title = "预加载 (Eager Loading)"
date = 2023-10-28T14:29:04+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/preload.html](https://gorm.io/docs/preload.html)

## 预加载 Preload

GORM allows eager loading relations in other SQL with `Preload`, for example:

​	GORM允许在SQL中使用`Preload`来预加载其他表的关联数据，例如：

``` go
type User struct {
  gorm.Model
  Username string
  Orders   []Order
}

type Order struct {
  gorm.Model
  UserID uint
  Price  float64
}

// 当查找用户时预加载订单 Preload Orders when find users
db.Preload("Orders").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4);

db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4); // has many
// SELECT * FROM profiles WHERE user_id IN (1,2,3,4); // has one
// SELECT * FROM roles WHERE id IN (4,5,6); // belongs to
```

## 连接预加载 Joins Preloading

`Preload` loads the association data in a separate query, `Join Preload` will loads association data using left join, for example:

​	`Preload`会将关联数据加载在一个单独的查询中，`Join Preload`会使用左连接来加载关联数据，例如：

``` go
db.Joins("Company").Joins("Manager").Joins("Account").First(&user, 1)
db.Joins("Company").Joins("Manager").Joins("Account").First(&user, "users.name = ?", "jinzhu")
db.Joins("Company").Joins("Manager").Joins("Account").Find(&users, "users.id IN ?", []int{1,2,3,4,5})
```

有条件的Join Join with conditions

``` go
db.Joins("Company", DB.Where(&Company{Alive: true})).Find(&users)
// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id` AND `Company`.`alive` = true;
```

Join嵌套模型 Join nested model

``` go
db.Joins("Manager").Joins("Manager.Company").Find(&users)
// SELECT "users"."id","users"."created_at","users"."updated_at","users"."deleted_at","users"."name","users"."age","users"."birthday","users"."company_id","users"."manager_id","users"."active","Manager"."id" AS "Manager__id","Manager"."created_at" AS "Manager__created_at","Manager"."updated_at" AS "Manager__updated_at","Manager"."deleted_at" AS "Manager__deleted_at","Manager"."name" AS "Manager__name","Manager"."age" AS "Manager__age","Manager"."birthday" AS "Manager__birthday","Manager"."company_id" AS "Manager__company_id","Manager"."manager_id" AS "Manager__manager_id","Manager"."active" AS "Manager__active","Manager__Company"."id" AS "Manager__Company__id","Manager__Company"."name" AS "Manager__Company__name" FROM "users" LEFT JOIN "users" "Manager" ON "users"."manager_id" = "Manager"."id" AND "Manager"."deleted_at" IS NULL LEFT JOIN "companies" "Manager__Company" ON "Manager"."company_id" = "Manager__Company"."id" WHERE "users"."deleted_at" IS NULL
```

> **NOTE** `Join Preload` works with one-to-one relation, e.g: `has one`, `belongs to`
>
> **注意** `Join Preload`只适用于一对一关系，例如：`has one`，`belongs to`

## 预加载所有 Preload All

`clause.Associations` can work with `Preload` similar like `Select` when creating/updating, you can use it to `Preload` all associations, for example:

​	`clause.Associations`可以与`Preload`类似地用于创建/更新时预加载所有关联，例如：

``` go
type User struct {
  gorm.Model
  Name       string
  CompanyID  uint
  Company    Company
  Role       Role
  Orders     []Order
}

db.Preload(clause.Associations).Find(&users)
```

`clause.Associations` won’t preload nested associations, but you can use it with [Nested Preloading](https://gorm.io/docs/preload.html#nested_preloading) together, e.g:

​	`clause.Associations`不会预加载嵌套关联，但是你可以与[嵌套预加载](https://gorm.io/docs/preload.html#nested_preloading)一起使用，例如：

``` go
db.Preload("Orders.OrderItems.Product").Preload(clause.Associations).Find(&users)
```

## 带条件的预加载 Preload with conditions

GORM allows Preload associations with conditions, it works similar to [Inline Conditions](https://gorm.io/docs/query.html#inline_conditions)

​	GORM允许使用带条件的预加载，类似于[内联条件](https://gorm.io/docs/query.html#inline_conditions)，例如：

``` go
// 预加载订单并应用状态条件 Preload Orders with conditions
db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4) AND state NOT IN ('cancelled');

db.Where("state = ?", "active").Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
// SELECT * FROM users WHERE state = 'active';
// SELECT * FROM orders WHERE user_id IN (1,2) AND state NOT IN ('cancelled');
```

## 自定义预加载SQL Custom Preloading SQL

You are able to custom preloading SQL by passing in `func(db *gorm.DB) *gorm.DB`, for example:

​	你可以使用自定义的预加载SQL通过传入一个函数来实现，例如：

``` go
db.Preload("Orders", func(db *gorm.DB) *gorm.DB {
  return db.Order("orders.amount DESC")
}).Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4) order by orders.amount DESC;
```

## 嵌套预加载 Nested Preloading

GORM supports nested preloading, for example:

​	GORM支持嵌套预加载，例如：

``` go
db.Preload("Orders.OrderItems.Product").Preload("CreditCard").Find(&users)

// 自定义预加载条件以处理“Orders” Customize Preload conditions for `Orders`
// GORM也不会预加载未匹配的order的OrderItems然后 And GORM won't preload unmatched order's OrderItems then
db.Preload("Orders", "state = ?", "paid").Preload("Orders.OrderItems").Find(&users)
```

## 嵌入式预加载 Embedded Preloading

Embedded Preloading is used for [Embedded Struct](https://gorm.io/docs/models.html#embedded_struct), especially the
same struct. The syntax for Embedded Preloading is similar to Nested Preloading, they are divided by dot.

​	嵌入式预加载是用于嵌入式结构体的，特别是这种相同的结构体。嵌入式预加载的语法与嵌套预加载相同，它们由点分隔。

For example:

​	例如：

``` go
type Address struct {
  CountryID int
  Country   Country
}

type Org struct {
  PostalAddress   Address `gorm:"embedded;embeddedPrefix:postal_address_"`
  VisitingAddress Address `gorm:"embedded;embeddedPrefix:visiting_address_"`
  Address         struct {
    ID int
    Address
  }
}

// 只预加载Org.Address和Org.Address.Country Only preload Org.Address and Org.Address.Country
db.Preload("Address.Country")  // "Address" is has_one, "Country" is belongs_to (nested association)

// 只预加载Org.VisitingAddress（嵌入式） Only preload Org.VisitingAddress
db.Preload("PostalAddress.Country") // "PostalAddress.Country" is belongs_to (embedded association)

// 只预加载Org.NestedAddress（嵌入式） Only preload Org.NestedAddress
db.Preload("NestedAddress.Address.Country") // "NestedAddress.Address.Country" is belongs_to (embedded association)

// 所有预加载都包括"Address"，但不包括"Address.Country"，因为不会预加载嵌套关联。 All preloaded include "Address" but exclude "Address.Country", because it won't preload nested associations.
db.Preload(clause.Associations)
```

We can omit embedded part when there is no ambiguity.

​	我们可以省略嵌入式部分，如果没有歧义的话。例如：

``` go
type Address struct {
  CountryID int
  Country   Country
}

type Org struct {
  Address Address `gorm:"embedded"`
}

db.Preload("Address.Country")
db.Preload("Country") //  省略了"Address"部分，因为我们知道没有歧义 omit "Address" because there is no ambiguity
```

> **NOTE** `Embedded Preload` only works with `belongs to` relation.
> Values of other relations are the same in database, we can’t distinguish them.
>
> **注意** `嵌入式预加载` 仅适用于 `belongs to` 关系。 数据库中其他关系的值相同，我们无法区分它们。

