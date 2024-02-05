+++
title = "数据库索引"
date = 2023-10-28T14:35:45+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/indexes.html](https://gorm.io/docs/indexes.html)

GORM allows create database index with tag `index`, `uniqueIndex`, those indexes will be created when [AutoMigrate or CreateTable with GORM](https://gorm.io/docs/migration.html)

​	GORM允许使用标签`index`、`uniqueIndex`创建数据库索引，这些索引将在[AutoMigrate或使用GORM创建表](https://gorm.io/docs/migration.html)时创建。

## 索引标签 Index Tag

GORM accepts lots of index settings, like `class`, `type`, `where`, `comment`, `expression`, `sort`, `collate`, `option`

​	GORM接受许多索引设置，如`class`、`type`、`where`、`comment`、`expression`、`sort`、`collate`、`option`等。

Check the following example for how to use it

​	请参阅以下示例以了解如何使用它：

``` go
type User struct {
  Name  string `gorm:"index"`
  Name2 string `gorm:"index:idx_name,unique"`
  Name3 string `gorm:"index:,sort:desc,collate:utf8,type:btree,length:10,where:name3 != 'jinzhu'"`
  Name4 string `gorm:"uniqueIndex"`
  Age   int64  `gorm:"index:,class:FULLTEXT,comment:hello \\, world,where:age > 10"`
  Age2  int64  `gorm:"index:,expression:ABS(age)"`
}

// MySQL选项 MySQL option
type User struct {
  Name string `gorm:"index:,class:FULLTEXT,option:WITH PARSER ngram INVISIBLE"`
}

// PostgreSQL选项 PostgreSQL option
type User struct {
  Name string `gorm:"index:,option:CONCURRENTLY"`
}
```

### uniqueIndex

tag `uniqueIndex` works similar like `index`, it equals to `index:,unique`

​	标签`uniqueIndex`的工作方式与`index`相似，等于`index:,unique`。

``` go
type User struct {
  Name1 string `gorm:"uniqueIndex"`
  Name2 string `gorm:"uniqueIndex:idx_name,sort:desc"`
}
```

## 复合索引 Composite Indexes

Use same index name for two fields will creates composite indexes, for example:

​	使用两个字段具有相同名称的索引将创建复合索引，例如：

``` go
// 创建一个名为`idx_member`的复合索引，包含列`name`和`number` create composite index `idx_member` with columns `name`, `number`
type User struct {
  Name   string `gorm:"index:idx_member"`
  Number string `gorm:"index:idx_member"`
}
```

### 字段优先级 Fields Priority

The column order of a composite index has an impact on its performance so it must be chosen carefully

​	复合索引的列顺序对其性能有很大影响，因此必须谨慎选择。

You can specify the order with the `priority` option, the default priority value is `10`, if priority value is the same, the order will be based on model struct’s field index

​	您可以使用`priority`选项指定顺序，默认优先级值为`10`，如果优先级值相同，则根据模型结构字段的索引顺序。

``` go
type User struct {
  Name   string `gorm:"index:idx_member"`
  Number string `gorm:"index:idx_member"`
}
// 列顺序：name, number  column order: name, number

type User struct {
  Name   string `gorm:"index:idx_member,priority:2"`
  Number string `gorm:"index:idx_member,priority:1"`
}
// 列顺序：number, name  column order: number, name

type User struct {
  Name   string `gorm:"index:idx_member,priority:12"`
  Number string `gorm:"index:idx_member"`
}
// 列顺序：number, name  column order: number, name
```

### 共享复合索引 Shared composite indexes

If you are creating shared composite indexes with an embedding struct, you can’t specify the index name, as embedding the struct more than once results in the duplicated index name in db.

​	如果您使用嵌入结构体创建共享复合索引，您不能指定索引名称，因为嵌入结构体多次出现会导致数据库中重复的索引名称。

In this case, you can use index tag `composite`, it means the id of the composite index. All fields which have the same composite id of the struct are put together to the same index, just like the original rule. But the improvement is it lets the most derived/embedding struct generates the name of index by NamingStrategy. For example:

​	在这种情况下，您可以使用索引标签`composite`，表示复合索引的id。所有具有相同复合id的结构体的字段都将放在同一个索引上，就像原始规则一样。但是，改进是它让最派生/嵌入的结构体通过NamingStrategy生成索引的名称。例如：

``` go
type Foo struct {
  IndexA int `gorm:"index:,unique,composite:myname"`
  IndexB int `gorm:"index:,unique,composite:myname"`
}
```

If the table Foo is created, the name of composite index will be `idx_foo_myname`.

​	如果表Foo被创建，复合索引的名称将是`idx_foo_myname`。

``` go
type Bar0 struct {
  Foo
}

type Bar1 struct {
  Foo
}
```

Respectively, the name of composite index is `idx_bar0_myname` and `idx_bar1_myname`.

​	分别，复合索引的名称将是`idx_bar0_myname`和`idx_bar1_myname`。

`composite` only works if not specify the name of index.

​	`composite`仅在未指定索引名称时起作用。

## 多个索引 Multiple indexes

A field accepts multiple `index`, `uniqueIndex` tags that will create multiple indexes on a field

一个字段可以接受多个`index`、`uniqueIndex`标签，这将在一个字段上创建多个索引。

``` go
type UserIndex struct {
  OID          int64  `gorm:"index:idx_id;index:idx_oid,unique"`
  MemberNumber string `gorm:"index:idx_id"`
}
```