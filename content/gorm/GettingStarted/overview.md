+++
title = "Overview"
date = 2023-10-28T14:24:01+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/index.html](https://gorm.io/docs/index.html)

The fantastic ORM library for Golang aims to be developer friendly.

- Full-Featured ORM
- 功能齐全的ORM
- Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism, Single-table inheritance)
- 关联（一对多、一对一、多对多、多态、单表继承）
- Hooks (Before/After Create/Save/Update/Delete/Find)
- 钩子（创建/保存/更新/删除/查找之前/之后）
- Eager loading with `Preload`, `Joins`
- 预加载（`Preload`）、联接（`Joins`）的延迟加载
- Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
- 事务、嵌套事务、保存点、回滚到已保存点
- Context, Prepared Statement Mode, DryRun Mode
- 上下文，预处理语句模式，干运行模式
- Batch Insert, FindInBatches, Find/Create with Map, CRUD with SQL Expr and Context Valuer
- 批量插入，分批查找，使用Map查找/创建，使用SQL表达式和上下文值进行CRUD操作
- SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery
- SQL构建器，Upsert，锁定，优化器/索引/注释提示，命名参数，子查询
- Composite Primary Key, Indexes, Constraints
- 复合主键、索引、约束
- Auto Migrations
- 自动迁移
- Logger
- 日志记录器
- Extendable, flexible plugin API: Database Resolver (Multiple Databases, Read/Write Splitting) / Prometheus…
- 可扩展且灵活的插件API：数据库解析器（多个数据库，读写分离）/ Prometheus等
- Every feature comes with tests
- 每个功能都有测试用例
- Developer Friendly
- 对开发人员友好

## 安装 Install

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

## 快速入门 Quick Start

```go
package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // 迁移模式 Migrate the schema
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  db.First(&product, 1) // 查找具有整数主键的产品 find product with integer primary key
  db.First(&product, "code = ?", "D42") // 查找编码为D42的产品 find product with code D42

  // 更新 - 将产品价格更新为200 Update - update product's price to 200
  db.Model(&product).Update("Price", 200)
  // 更新 - 更新多个字段 Update - update multiple fields
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.Delete(&product, 1)
}
```