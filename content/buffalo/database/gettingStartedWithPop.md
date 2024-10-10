+++
title = "Pop入门"
date = 2024-02-04T21:12:02+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/pop/]({{< ref "/buffalo/database/gettingStartedWithPop" >}})

# Getting Started with Pop Pop入门

The [pop](https://godoc.org/github.com/gobuffalo/pop) package is included with Buffalo by default, but you can use it outside of Buffalo. It wraps the absolutely amazing https://github.com/jmoiron/sqlx library, cleans up some of the common patterns and work flows usually associated with dealing with databases in Go.

​	pop包默认包含在Buffalo中，但您可以在Buffalo之外使用它。它包装了非常棒的https://github.com/jmoiton/sqlx库，清理了一些通常与在Go中使用数据库相关的常见模型和工作流。

**Pop makes it easy to do CRUD operations with basic ORM functionality, run migrations, and build/execute queries.
Pop可以轻松地使用基本的 CRUD 操作，运行迁移，并构建/运行SQL语句。**

Pop, by default, follows conventions that were influenced by the ActiveRecord Ruby gem. What does this mean?

​	Pop默认遵循受ActiveRecord Ruby宝石启发的惯例。这是什么含义？

- Tables must have an “id” column and a corresponding “ID” field on the struct being used.
  表中，列名和结构体中的字段名，都应遵循“id”和“ID”的对应规则。
- If there is a timestamp column named `created_at`, and a `CreatedAt time.Time` attribute on the struct, it will be set with the current time when the record is created.
  如果表中有一个时间戳列名是 `created_at` ，并且结构体中有一个 `CreatedAt time.Time` 属性，则在创建时间时，该属性会设置为当时的时间。
- If there is a timestamp column named `updated_at`, and a `UpdatedAt time.Time` attribute on the struct, it will be set with the current time when the record is updated.
  如果表中有一个时间戳列名是 `updated_at` ，并且结构体中有一个 `UpdatedAt time.Time` 属性，则在创建时间时，该属性会设置为当时的时间。
- Default database table names are lowercase, plural, and underscored versions of the struct name. Examples: `User{}` is “users”, `FooBar{}` is “foo_bars”, etc…
  默认的数据库表名是结构体名的复数形式，并且以小写字母和下划线表示。示例： `User{}` 是“users”、 `FooBar{}` 是“foo_bars”等…

Buffalo has a deep integration with Pop, and it’ll help you to generate all the stuff you need to get started. You can still use another package if you want, but you’ll be by yourself. :)

​	Buffalo 与 Pop 深度集成，它将帮助您生成开始所需的所有内容。如果您愿意，仍然可以使用其他软件包，但您将独自一人。 :)

## Supported Databases 支持的数据库 

Pop supports the following databases:

​	Pop 支持以下数据库：

- [PostgreSQL](https://www.postgresql.org/) (>= 9.3)
- [CockroachDB](https://www.cockroachlabs.com/) (>= 2.1.0)
- [MySQL](https://www.mysql.com/) (>= 5.7)
- [SQLite3](https://sqlite.org/) (>= 3.x)

## Installation 安装 

```bash
$ go get github.com/gobuffalo/pop/...
```

## Next Steps 后续步骤 

- [CLI Soda]({{< ref "/buffalo/database/sodaCLI" >}}) - Install the Soda CLI.
  CLI Soda - 安装 Soda CLI。
- [Configuration]({{< ref "/buffalo/database/databaseConfiguration" >}}) - Configure your database connections.
  配置 - 配置您的数据库连接。
