+++
title = "overview"
date = 2024-02-04T09:59:56+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/model/overview/](https://beego.wiki/docs/mvc/model/overview/)

# Overview 概述



## Models － Beego ORM 模型 － Beego ORM

Beego ORM is a powerful ORM framework written in Go. It is inspired by Django ORM and SQLAlchemy.

&zeroWidthSpace;Beego ORM 是一个用 Go 编写的功能强大的 ORM 框架。它受到 Django ORM 和 SQLAlchemy 的启发。

This framework is still under development so compatibility is not guaranteed.

&zeroWidthSpace;此框架仍在开发中，因此不保证兼容性。

**Supported Database: 支持的数据库：**

- MySQL：[github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- PostgreSQL：[github.com/lib/pq](https://github.com/lib/pq)
- Sqlite3：[github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

All of the database drivers have passed the tests, but we still need your feedback and bug reports.

&zeroWidthSpace;所有数据库驱动程序都已通过测试，但我们仍然需要您的反馈和错误报告。

**ORM Features: ORM 特性：**

- Supports all the types in Go.
  支持 Go 中的所有类型。
- CRUD is easy to use.
  CRUD 易于使用。
- Auto join connection tables.
  自动连接表。
- Compatible with crossing database queries.
  兼容跨数据库查询。
- Supports raw SQL query and mapping.
  支持原始 SQL 查询和映射。
- Strict and well-covered test cases ensure the ORM’s stability.
  严格且覆盖范围广的测试用例确保了 ORM 的稳定性。

You can learn more in this documentation.

&zeroWidthSpace;您可以在此文档中了解更多信息。

**Install ORM: 安装 ORM：**

```
go get github.com/beego/beego/v2/client/orm
```

## Quickstart 快速入门

### Demo 演示

```go
package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
```

### Relation Query 关系查询

```go
type Post struct {
	Id    int    `orm:"auto"`
	Title string `orm:"size(100)"`
	User  *User  `orm:"rel(fk)"`
}

var posts []*Post
qs := o.QueryTable("post")
num, err := qs.Filter("User__Name", "slene").All(&posts)
```

### Raw SQL query 原始 SQL 查询

You can always use raw SQL to query and mapping.

&zeroWidthSpace;您可以始终使用原始 SQL 进行查询和映射。

```go
var maps []Params
num, err := o.Raw("SELECT id FROM user WHERE name = ?", "slene").Values(&maps)
if num > 0 {
	fmt.Println(maps[0]["id"])
}
```

### Transactions 事务

```go
o.Begin()
...
user := User{Name: "slene"}
id, err := o.Insert(&user)
if err == nil {
	o.Commit()
} else {
	o.Rollback()
}
```

### Debugging query log 调试查询日志

In development environment, you can enable debug mode by:

&zeroWidthSpace;在开发环境中，您可以通过以下方式启用调试模式：

```go
func main() {
	orm.Debug = true
...
```

It will output every query statement including execution, preparation and transactions.

&zeroWidthSpace;它将输出包括执行、准备和事务在内的每个查询语句。

For example: 
&zeroWidthSpace;例如：

```go
[ORM] - 2013-08-09 13:18:16 - [Queries/default] - [    db.Exec /     0.4ms] - 	[INSERT INTO `user` (`name`) VALUES (?)] - `slene`
...
```

Notes: It is not recommended to enable debug mode in a production environment.

&zeroWidthSpace;注意：不建议在生产环境中启用调试模式。

