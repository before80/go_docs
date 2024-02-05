+++
title = "编写驱动"
date = 2023-10-28T14:37:09+08:00
weight = 13
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/write_driver.html](https://gorm.io/docs/write_driver.html)

## 编写新驱动 Write new driver

GORM provides official support for `sqlite`, `mysql`, `postgres`, `sqlserver`.

​	GORM 为 `sqlite`、`mysql`、`postgres`、`sqlserver` 提供了官方支持。

Some databases may be compatible with the `mysql` or `postgres` dialect, in which case you could just use the dialect for those databases.

​	一些数据库可能与 `mysql` 或 `postgres` 方言兼容，在这种情况下，您可以使用这些数据库的方言。

For others, you can create a new driver, it needs to implement [the dialect interface](https://pkg.go.dev/gorm.io/gorm?tab=doc#Dialector).

​	对于其他数据库，您可以创建一个新驱动程序，它需要实现 [方言接口](https://pkg.go.dev/gorm.io/gorm?tab=doc#Dialector)。

``` go
type Dialector interface {
  Name() string
  Initialize(*DB) error
  Migrator(db *DB) Migrator
  DataTypeOf(*schema.Field) string
  DefaultValueOf(*schema.Field) clause.Expression
  BindVarTo(writer clause.Writer, stmt *Statement, v interface{})
  QuoteTo(clause.Writer, string)
  Explain(sql string, vars ...interface{}) string
}
```

Checkout the [MySQL Driver](https://github.com/go-gorm/mysql) as example

​	以 [MySQL 驱动程序](https://github.com/go-gorm/mysql)为例