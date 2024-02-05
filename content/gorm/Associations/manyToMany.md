+++
title = "多对多"
date = 2023-10-28T14:28:27+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/many_to_many.html](https://gorm.io/docs/many_to_many.html)

## 多对多 Many To Many

Many to Many add a join table between two models.

​	多对多在两个模型之间添加一个连接表。

For example, if your application includes users and languages, and a user can speak many languages, and many users can speak a specified language.

​	例如，如果你的应用包括用户和语言，用户可以说多种语言，也可以说多种语言。

``` go
// User 有且属于多种语言，`user_languages` 是连接表 User has and belongs to many languages, `user_languages` is the join table
type User struct {
  gorm.Model
  Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
  gorm.Model
  Name string
}
```

When using GORM `AutoMigrate` to create a table for `User`, GORM will create the join table automatically

​	当使用 GORM `AutoMigrate` 为 `User` 创建表时，GORM 会自动创建连接表

## 反向引用 Back-Reference

### 声明 Declare

``` go
// User 有且属于多种语言，使用 `user_languages` 作为连接表 User has and belongs to many languages, use `user_languages` as join table
type User struct {
  gorm.Model
  Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
  gorm.Model
  Name string
  Users []*User `gorm:"many2many:user_languages;"`
}
```

### 检索 Retrieve

``` go
//  检索带有预加载语言的用户列表 Retrieve user list with eager loading languages
func GetAllUsers(db *gorm.DB) ([]User, error) {
  var users []User
  err := db.Model(&User{}).Preload("Languages").Find(&users).Error
  return users, err
}

// 检索带有预加载用户的编程语言列表 Retrieve language list with eager loading users
func GetAllLanguages(db *gorm.DB) ([]Language, error) {
  var languages []Language
  err := db.Model(&Language{}).Preload("Users").Find(&languages).Error
  return languages, err
}
```

## 覆盖外键 Override Foreign Key

For a `many2many` relationship, the join table owns the foreign key which references two models, for example:

​	对于 `many2many` 关系，连接表拥有引用两个模型的外键，例如：

``` go
type User struct {
  gorm.Model
  Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
  gorm.Model
  Name string
}

// Join Table: user_languages
//   foreign key: user_id, reference: users.id
//   foreign key: language_id, reference: languages.id
```

To override them, you can use tag `foreignKey`, `references`, `joinForeignKey`, `joinReferences`, not necessary to use them together, you can just use one of them to override some foreign keys/references

​	要覆盖它们，可以使用标签 `foreignKey`、`references`、`joinForeignKey`、`joinReferences`，不一定需要同时使用它们，只需使用其中一个来覆盖某些外键/引用即可。

``` go
type User struct {
  gorm.Model
  Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:ProfileRefer"`
  Refer    uint      `gorm:"index:,unique"`
}

type Profile struct {
  gorm.Model
  Name      string
  UserRefer uint `gorm:"index:,unique"`
}

// Which creates join table: user_profiles
//   foreign key: user_refer_id, reference: users.refer
//   foreign key: profile_refer, reference: profiles.user_refer
```

> **NOTE:**
> Some databases only allow create database foreign keys that reference on a field having unique index, so you need to specify the `unique index` tag if you are creating database foreign keys when migrating
>
> **注意：** 一些数据库只允许创建指向具有唯一索引的字段的数据库外键，因此如果您在迁移时创建数据库外键，您需要指定 `unique index` 标签。

## 自引用多对多 Self-Referential Many2Many

Self-referencing many2many relationship

​	自引用多对多关系

``` go
type User struct {
  gorm.Model
  Friends []*User `gorm:"many2many:user_friends"`
}

// Which creates join table: user_friends
//   foreign key: user_id, reference: users.id
//   foreign key: friend_id, reference: users.id
```

## 预加载（Eager loading）Eager Loading

GORM allows eager loading has many associations with `Preload`, refer [Preloading (Eager loading)](https://gorm.io/docs/preload.html) for details

​	GORM 允许使用 `Preload` 进行预加载具有许多关联的 `many2many`。有关详细信息，请参阅 [预加载（Eager loading）](https://gorm.io/docs/preload.html)。

## 与 many2many 关系的 CRUD CRUD with Many2Many

Please checkout [Association Mode](https://gorm.io/docs/associations.html#Association-Mode) for working with many2many relations

​	请查看 [Association Mode](https://gorm.io/docs/associations.html#Association-Mode) 以处理 many2many 关系。

## Customize JoinTable

`JoinTable` can be a full-featured model, like having `Soft Delete`，`Hooks` supports and more fields, you can set it up with `SetupJoinTable`, for example:

​	`JoinTable` 可以是一个完整的功能模型，如具有 `Soft Delete`、`Hooks` 支持和更多字段的模型，您可以使用 `SetupJoinTable` 设置它，例如：

> **NOTE:**
> Customized join table’s foreign keys required to be composited primary keys or composited unique index
>
> **注意：** 自定义连接表的外键要求是复合主键或复合唯一索引。

``` go
type Person struct {
  ID        int
  Name      string
  Addresses []Address `gorm:"many2many:person_addressses;"`
}

type Address struct {
  ID   uint
  Name string
}

type PersonAddress struct {
  PersonID  int `gorm:"primaryKey"`
  AddressID int `gorm:"primaryKey"`
  CreatedAt time.Time
  DeletedAt gorm.DeletedAt
}

func (PersonAddress) BeforeCreate(db *gorm.DB) error {
  // ...
}

// Change model Person's field Addresses' join table to PersonAddress
// PersonAddress must defined all required foreign keys or it will raise error
err := db.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{})
```

## FOREIGN KEY Constraints

You can setup `OnUpdate`, `OnDelete` constraints with tag `constraint`, it will be created when migrating with GORM, for example:

​	您可以使用标签 `constraint` 设置 `OnUpdate`、`OnDelete` 约束，它在迁移时由 GORM 创建，例如：

``` go
type User struct {
  gorm.Model
  Languages []Language `gorm:"many2many:user_speaks;"`
}

type Language struct {
  Code string `gorm:"primarykey"`
  Name string
}

// CREATE TABLE `user_speaks` (`user_id` integer,`language_code` text,PRIMARY KEY (`user_id`,`language_code`),CONSTRAINT `fk_user_speaks_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,CONSTRAINT `fk_user_speaks_language` FOREIGN KEY (`language_code`) REFERENCES `languages`(`code`) ON DELETE SET NULL ON UPDATE CASCADE);
```

You are also allowed to delete selected many2many relations with `Select` when deleting, checkout [Delete with Select](https://gorm.io/docs/associations.html#delete_with_select) for details

## Composite Foreign Keys

If you are using [Composite Primary Keys](https://gorm.io/docs/composite_primary_key.html) for your models, GORM will enable composite foreign keys by default

​	如果您使用 [Composite Primary Keys](https://gorm.io/docs/composite_primary_key.html) 为您的模型，GORM 默认情况下会启用复合外键。

You are allowed to override the default foreign keys, to specify multiple foreign keys, just separate those keys’ name by commas, for example:

​	您还可以覆盖默认的外键，以指定多个外键，只需用逗号分隔这些键的名称即可，例如：

``` go
type Tag struct {
  ID     uint   `gorm:"primaryKey"`
  Locale string `gorm:"primaryKey"`
  Value  string
}

type Blog struct {
  ID         uint   `gorm:"primaryKey"`
  Locale     string `gorm:"primaryKey"`
  Subject    string
  Body       string
  Tags       []Tag `gorm:"many2many:blog_tags;"`
  LocaleTags []Tag `gorm:"many2many:locale_blog_tags;ForeignKey:id,locale;References:id"`
  SharedTags []Tag `gorm:"many2many:shared_blog_tags;ForeignKey:id;References:id"`
}

// Join Table: blog_tags
//   foreign key: blog_id, reference: blogs.id
//   foreign key: blog_locale, reference: blogs.locale
//   foreign key: tag_id, reference: tags.id
//   foreign key: tag_locale, reference: tags.locale

// Join Table: locale_blog_tags
//   foreign key: blog_id, reference: blogs.id
//   foreign key: blog_locale, reference: blogs.locale
//   foreign key: tag_id, reference: tags.id

// Join Table: shared_blog_tags
//   foreign key: blog_id, reference: blogs.id
//   foreign key: tag_id, reference: tags.id
```

Also check out [Composite Primary Keys](https://gorm.io/docs/composite_primary_key.html)

​	请查看[复合主键](https://gorm.io/docs/composite_primary_key.html)