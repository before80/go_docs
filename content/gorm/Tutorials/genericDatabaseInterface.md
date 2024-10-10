+++
title = "通用数据库接口 sql.DB"
date = 2023-10-28T14:31:29+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/generic_interface.html](https://gorm.io/docs/generic_interface.html)

GORM provides the method `DB` which returns a generic database interface [*sql.DB](https://pkg.go.dev/database/sql#DB) from the current `*gorm.DB`

​	GORM提供了方法`DB`，它从当前的`*gorm.DB`返回一个通用的数据库接口[*sql.DB]({{< ref "/stdLib/database/sql#type-db">}})。

``` go
// 获取通用数据库对象sql.DB以使用其函数 Get generic database object sql.DB to use its functions
sqlDB, err := db.DB()

// Ping
sqlDB.Ping()

// Close
sqlDB.Close()

// Returns database statistics
sqlDB.Stats()
```

> **NOTE** If the underlying database connection is not a `*sql.DB`, like in a transaction, it will returns error
>
> **注意** 如果底层数据库连接不是`*sql.DB`，例如在事务中，它将返回错误

## 连接池 Connection Pool

``` go
// 获取通用数据库对象sql.DB以使用其函数 Get generic database object sql.DB to use its functions
sqlDB, err := db.DB()

// SetMaxIdleConns 设置空闲连接池中的最大连接数。 SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
sqlDB.SetMaxIdleConns(10)

// SetMaxOpenConns 设置与数据库的打开连接的最大数量。 SetMaxOpenConns sets the maximum number of open connections to the database.
sqlDB.SetMaxOpenConns(100)

// SetConnMaxLifetime 设置连接的最大生命周期。 SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
sqlDB.SetConnMaxLifetime(time.Hour)
```