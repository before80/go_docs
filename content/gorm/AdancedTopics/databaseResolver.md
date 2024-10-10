+++
title = "DBResolver"
date = 2023-10-28T14:34:34+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/dbresolver.html](https://gorm.io/docs/dbresolver.html)

DBResolver adds multiple databases support to GORM, the following features are supported:

​	DBResolver为GORM提供了多个数据库支持，以下是支持的功能：

- Multiple sources, replicas
- 多数据源、副本
- Read/Write Splitting
- 读写分离
- Automatic connection switching based on the working table/struct
- 根据工作表/结构自动切换连接
- Manual connection switching
- 手动切换连接
- Sources/Replicas load balancing
- 数据源/副本负载均衡
- Works for RAW SQL
- 支持RAW SQL
- Transaction
- 事务

https://github.com/go-gorm/dbresolver

## 用法 Usage

``` go
import (
  "gorm.io/gorm"
  "gorm.io/plugin/dbresolver"
  "gorm.io/driver/mysql"
)

db, err := gorm.Open(mysql.Open("db1_dsn"), &gorm.Config{})

db.Use(dbresolver.Register(dbresolver.Config{
  // 使用 `db2` 作为数据源，`db3`、`db4` 作为副本 use `db2` as sources, `db3`, `db4` as replicas
  Sources:  []gorm.Dialector{mysql.Open("db2_dsn")},
  Replicas: []gorm.Dialector{mysql.Open("db3_dsn"), mysql.Open("db4_dsn")},
  // 数据源/副本负载均衡策略 sources/replicas load balancing policy
  Policy: dbresolver.RandomPolicy{},
  // 在日志中打印数据源/副本模式 print sources/replicas mode in logger
  TraceResolverMode: true,
}).Register(dbresolver.Config{
  // 使用 `db1` 作为默认连接（数据库的默认连接），`db5` 用于 `User`、`Address` 的副本 use `db1` as sources (DB's default connection), `db5` as replicas for `User`, `Address`
  Replicas: []gorm.Dialector{mysql.Open("db5_dsn")},
}, &User{}, &Address{}).Register(dbresolver.Config{
  // 使用 `db6`、`db7` 作为数据源，`db8` 作为 `orders`、`Product` 的副本 use `db6`, `db7` as sources, `db8` as replicas for `orders`, `Product`
  Sources:  []gorm.Dialector{mysql.Open("db6_dsn"), mysql.Open("db7_dsn")},
  Replicas: []gorm.Dialector{mysql.Open("db8_dsn")},
}, "orders", &Product{}, "secondary"))
```

## 自动连接切换 Automatic connection switching

DBResolver will automatically switch connection based on the working table/struct

​	DBResolver 根据正在工作的表/结构自动切换连接。

For RAW SQL, DBResolver will extract the table name from the SQL to match the resolver, and will use `sources` unless the SQL begins with `SELECT` (excepts `SELECT... FOR UPDATE`), for example:

​	对于RAW SQL，DBResolver将从SQL中提取表名以匹配解析器，除非SQL以`SELECT`开头（除了`SELECT... FOR UPDATE`），例如：

``` go
// `User` 解析器示例 `User` Resolver Examples
db.Table("users").Rows() // replicas `db5`
db.Model(&User{}).Find(&AdvancedUser{}) // replicas `db5`
db.Exec("update users set name = ?", "jinzhu") // sources `db1`
db.Raw("select name from users").Row().Scan(&name) // replicas `db5`
db.Create(&user) // sources `db1`
db.Delete(&User{}, "name = ?", "jinzhu") // sources `db1`
db.Table("users").Update("name", "jinzhu") // sources `db1`

// 全局解析器示例 Global Resolver Examples
db.Find(&Pet{}) // replicas `db3`/`db4`
db.Save(&Pet{}) // sources `db2`

// Orders 解析器示例 Orders Resolver Examples
db.Find(&Order{}) // replicas `db8`
db.Table("orders").Find(&Report{}) // replicas `db8`
```

## 读写分离 Read/Write Splitting

Read/Write splitting with DBResolver based on the current used [GORM callbacks](https://gorm.io/docs/write_plugins.html).

​	基于当前使用的GORM回调进行读写分离的DBResolver。

For `Query`, `Row` callback, will use `replicas` unless `Write` mode specified

​	对于`Query`、`Row`回调，将使用副本，除非指定了`Write`模式。

For `Raw` callback, statements are considered read-only and will use `replicas` if the SQL starts with `SELECT`

​	对于`Raw`回调，语句被认为是只读的，如果SQL以`SELECT`开头，则将使用副本。

## 手动连接切换 Manual connection switching

``` go
// 使用写模式：从数据源 `db1` 读取用户 Use Write Mode: read user from sources `db1`
db.Clauses(dbresolver.Write).First(&user)

// 指定解析器：从 `secondary` 的副本中读取用户：db8 Specify Resolver: read user from `secondary`'s replicas: db8
db.Clauses(dbresolver.Use("secondary")).First(&user)

// 指定解析器和写模式：从 `secondary` 的数据源中读取用户：db6 或 db7  Specify Resolver and Write Mode: read user from `secondary`'s sources: db6 or db7
db.Clauses(dbresolver.Use("secondary"), dbresolver.Write).First(&user)
```

## 事务 Transaction

When using transaction, DBResolver will keep using the transaction and won’t switch to sources/replicas based on configuration

​	在使用事务时，DBResolver将保持使用事务并不会根据配置切换到数据源/副本。

But you can specifies which DB to use before starting a transaction, for example:

​	但是，您可以在开始事务之前指定要使用的数据库，例如：

``` go
// 根据默认副本数据库启动事务 Start transaction based on default replicas db
tx := DB.Clauses(dbresolver.Read).Begin()

// 根据默认数据源数据库启动事务 Start transaction based on default sources db
tx := DB.Clauses(dbresolver.Write).Begin()

// 根据 `secondary` 的数据源启动事务 Start transaction based on `secondary`'s sources
tx := DB.Clauses(dbresolver.Use("secondary"), dbresolver.Write).Begin()
```

## 负载均衡 Load Balancing

GORM supports load balancing sources/replicas based on policy, the policy should be a struct implements following interface:

​	GORM支持基于策略的负载均衡数据源/副本，策略应实现以下接口：

``` go
type Policy interface {
  Resolve([]gorm.ConnPool) gorm.ConnPool
}
```

Currently only the `RandomPolicy` implemented and it is the default option if no other policy specified.

​	目前只有实现了`RandomPolicy`的策略，如果没有指定其他策略，则为默认选项。

## 连接池 Connection Pool

``` go
db.Use(
  dbresolver.Register(dbresolver.Config{ /* xxx */ }).
  SetConnMaxIdleTime(time.Hour).
  SetConnMaxLifetime(24 * time.Hour).
  SetMaxIdleConns(100).
  SetMaxOpenConns(200)
)
```