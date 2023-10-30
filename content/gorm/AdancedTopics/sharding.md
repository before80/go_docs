+++
title = "分片"
date = 2023-10-28T14:34:44+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/sharding.html](https://gorm.io/docs/sharding.html)

Sharding plugin using SQL parser and replace for splits large tables into smaller ones, redirects Query into sharding tables. Give you a high performance database access.

​	分片插件使用SQL解析器和替换将大表拆分为小表，将查询重定向到分片表。为您提供高性能的数据库访问。

https://github.com/go-gorm/sharding

## 功能 Features

- Non-intrusive design. Load the plugin, specify the config, and all done.
- 无侵入式设计。加载插件，指定配置，即可完成。
- Lighting-fast. No network based middlewares, as fast as Go.
- 快速。没有基于网络的中间件，速度与Go相当。
- Multiple database support. PostgreSQL tested, MySQL and SQLite is coming.
- 支持多个数据库。已测试PostgreSQL，即将支持MySQL和SQLite。
- Allows you custom the Primary Key generator (Built in keygen, Sequence, Snowflake …).
- 支持多个数据库。已测试PostgreSQL，即将支持MySQL和SQLite。

## 用法 Usage

Config the sharding middleware, register the tables which you want to shard. See [Godoc](https://pkg.go.dev/github.com/go-gorm/sharding) for config details.

​	配置分片中间件，注册要分片的表。查看[Godoc](https://pkg.go.dev/github.com/go-gorm/sharding)以获取配置详细信息。

``` go
import (
  "fmt"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/sharding"
)

dsn := "postgres://localhost:5432/sharding-db?sslmode=disable"
db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}))

db.Use(sharding.Register(sharding.Config{
    ShardingKey:         "user_id",
    NumberOfShards:      64,
    PrimaryKeyGenerator: sharding.PKSnowflake,
}, "orders").Register(sharding.Config{
    ShardingKey:         "user_id",
    NumberOfShards:      256,
    PrimaryKeyGenerator: sharding.PKSnowflake,
    // This case for show up give notifications, audit_logs table use same sharding rule.
}, Notification{}, AuditLog{}))
```

Use the db session as usual. Just note that the query should have the `Sharding Key` when operate sharding tables.

​	像往常一样使用db会话。只需注意，操作分片表时查询应该包含`分片键`。

``` go
// Gorm创建示例，这将插入到orders_02 Gorm create example, this will insert to orders_02
db.Create(&Order{UserID: 2})
// sql: INSERT INTO orders_2 ...

// 显示已使用Raw SQL插入，这将插入到orders_03 Show have use Raw SQL to insert, this will insert into orders_03
db.Exec("INSERT INTO orders(user_id) VALUES(?)", int64(3))

// 这将抛出ErrMissingShardingKey错误，因为没有包含分片键。 This will throw ErrMissingShardingKey error, because there not have sharding key presented.
db.Create(&Order{Amount: 10, ProductID: 100})
fmt.Println(err)

// 查找，这将将查询重定向到orders_02 Find, this will redirect query to orders_02
var orders []Order
db.Model(&Order{}).Where("user_id", int64(2)).Find(&orders)
fmt.Printf("%#v\n", orders)

// Raw SQL也支持 Raw SQL also supported
db.Raw("SELECT * FROM orders WHERE user_id = ?", int64(3)).Scan(&orders)
fmt.Printf("%#v\n", orders)

// 这将抛出ErrMissingShardingKey错误，因为WHERE条件没有包含分片键 This will throw ErrMissingShardingKey error, because WHERE conditions not included sharding key
err = db.Model(&Order{}).Where("product_id", "1").Find(&orders).Error
fmt.Println(err)

// 更新和删除与创建和查询类似 Update and Delete are similar to create and query
db.Exec("UPDATE orders SET product_id = ? WHERE user_id = ?", 2, int64(3))
err = db.Exec("DELETE FROM orders WHERE product_id = 3").Error
fmt.Println(err) // ErrMissingShardingKey
```