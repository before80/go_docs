+++
title = "性能"
date = 2023-10-28T14:31:43+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/performance.html](https://gorm.io/docs/performance.html)

GORM optimizes many things to improve the performance, the default performance should be good for most applications, but there are still some tips for how to improve it for your application.

​	GORM 优化了许多方面以提高性能，默认的性能对于大多数应用程序来说应该是不错的，但仍然有一些方法可以帮助您提高您的应用程序的性能。

## 禁用默认事务 Disable Default Transaction

GORM performs write (create/update/delete) operations inside a transaction to ensure data consistency, which is bad for performance, you can disable it during initialization

​	GORM 在执行写（创建/更新/删除）操作时会在一个事务中进行以确保数据一致性，这对性能不利，您可以在初始化时禁用它

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
})
```

## 缓存预处理语句 Caches Prepared Statement

Creates a prepared statement when executing any SQL and caches them to speed up future calls

​	在执行任何 SQL 时创建一个预处理语句并将其缓存以加速未来的调用

``` go
// 全局模式 Globally mode
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  PrepareStmt: true,
})

// 会话模式 Session mode
tx := db.Session(&Session{PrepareStmt: true})
tx.First(&user, 1)
tx.Find(&users)
tx.Model(&user).Update("Age", 18)
```

> **NOTE** Also refer how to enable interpolateparams for MySQL to reduce roundtrip https://github.com/go-sql-driver/mysql#interpolateparams
>
> **注意** 还可以参考如何为 MySQL 启用 interpolateparams 以减少往返 [https://github.com/go-sql-driver/mysql#interpolateparams](https://github.com/go-sql-driver/mysql#interpolateparams)

### 带有预处理语句的 SQL 构建器 SQL Builder with PreparedStmt

Prepared Statement works with RAW SQL also, for example:

​	预处理语句也可以与原始 SQL 一起使用，例如：

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  PrepareStmt: true,
})

db.Raw("select sum(age) from users where role = ?", "admin").Scan(&age)
```

You can also use GORM API to prepare SQL with [DryRun Mode](https://gorm.io/docs/session.html), and execute it with prepared statement later, checkout [Session Mode](https://gorm.io/docs/session.html) for details

​	您还可以使用 GORM API 准备 SQL，并在 [DryRun 模式下](../session) 使用预处理语句稍后执行它，有关详细信息请参阅 [会话模式](../session)

## 选择字段 Select Fields

By default GORM select all fields when querying, you can use `Select` to specify fields you want

​	默认情况下，当查询时 GORM 会选择所有字段，您可以使用 `Select` 指定要选择的字段

``` go
db.Select("Name", "Age").Find(&Users{})
```

Or define a smaller API struct to use the [smart select fields feature](https://gorm.io/docs/advanced_query.html)

​	或者定义一个较小的 API 结构体以使用 [智能选择字段功能]({{< ref "/gorm/CRUDInterface/advancedQuery">}})

``` go
type User struct {
  ID     uint
  Name   string
  Age    int
  Gender string
  // hundreds of fields
}

type APIUser struct {
  ID   uint
  Name string
}

// Select `id`, `name` automatically when query
db.Model(&User{}).Limit(10).Find(&APIUser{})
// SELECT `id`, `name` FROM `users` LIMIT 10
```

## 迭代 / FindInBatches Iteration / FindInBatches

Query and process records with iteration or in batches

​	使用迭代或分批查询和处理记录

## 索引提示 Index Hints

[Index](https://gorm.io/docs/indexes.html) is used to speed up data search and SQL query performance. `Index Hints` gives the optimizer information about how to choose indexes during query processing, which gives the flexibility to choose a more efficient execution plan than the optimizer

​	[索引]({{< ref "/gorm/AdancedTopics/indexes">}}) 用于加速数据搜索和 SQL 查询性能。`Index Hints` 向优化器提供关于如何在查询处理过程中选择索引的信息，从而允许选择比优化器更高效的执行计划

``` go
import "gorm.io/hints"

db.Clauses(hints.UseIndex("idx_user_name")).Find(&User{})
// SELECT * FROM `users` USE INDEX (`idx_user_name`)

db.Clauses(hints.ForceIndex("idx_user_name", "idx_user_id").ForJoin()).Find(&User{})
// SELECT * FROM `users` FORCE INDEX FOR JOIN (`idx_user_name`,`idx_user_id`)"

db.Clauses(
  hints.ForceIndex("idx_user_name", "idx_user_id").ForOrderBy(),
  hints.IgnoreIndex("idx_user_name").ForGroupBy(),
).Find(&User{})
// SELECT * FROM `users` FORCE INDEX FOR ORDER BY (`idx_user_name`,`idx_user_id`) IGNORE INDEX FOR GROUP BY (`idx_user_name`)"
```

## 读写分离 Read/Write Splitting

Increase data throughput through read/write splitting, check out [Database Resolver](https://gorm.io/docs/dbresolver.html)

​	通过读写分离增加数据吞吐量，请参阅 [数据库解析器]({{< ref "/gorm/AdancedTopics/databaseResolver">}})