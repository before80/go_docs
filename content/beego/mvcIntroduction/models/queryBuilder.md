+++
title = "查询构建器"
date = 2024-02-04T10:01:16+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/model/querybuilder/](https://beego.wiki/docs/mvc/model/querybuilder/)

# Query Builder 查询构建器



## Query Builder 查询生成器

**QueryBuilder** provides an API for convenient and fluent construction of SQL queries. It consists of a set of methods enabling developers to easily construct SQL queries without compromising readability.

&zeroWidthSpace;QueryBuilder 提供了一个 API，用于方便且流畅地构建 SQL 查询。它由一组方法组成，使开发人员能够轻松构建 SQL 查询，而不会影响可读性。

It serves as an alternative to ORM. ORM is more for simple CRUD operations, whereas QueryBuilder is for complex queries with subqueries and multi-joins.

&zeroWidthSpace;它作为 ORM 的替代品。ORM 更适用于简单的 CRUD 操作，而 QueryBuilder 适用于具有子查询和多重联接的复杂查询。

Usage example:

&zeroWidthSpace;用法示例：

```go
// User is a wrapper for result row in this example
type User struct {
	Name string
	Age  int
}
var users []User

// Get a QueryBuilder object. Takes DB driver name as parameter
// Second return value is error, ignored here
qb, _ := orm.NewQueryBuilder("mysql")

// Construct query object
qb.Select("user.name",
	"profile.age").
	From("user").
	InnerJoin("profile").On("user.id_user = profile.fk_user").
	Where("age > ?").
	OrderBy("name").Desc().
	Limit(10).Offset(0)

// export raw query string from QueryBuilder object
sql := qb.String()

// execute the raw query string
o := orm.NewOrm()
o.Raw(sql, 20).QueryRows(&users)
```

Full API interface:

&zeroWidthSpace;完整的 API 接口：

```go
type QueryBuilder interface {
	Select(fields ...string) QueryBuilder
	ForUpdate() QueryBuilder
	From(tables ...string) QueryBuilder
	InnerJoin(table string) QueryBuilder
	LeftJoin(table string) QueryBuilder
	RightJoin(table string) QueryBuilder
	On(cond string) QueryBuilder
	Where(cond string) QueryBuilder
	And(cond string) QueryBuilder
	Or(cond string) QueryBuilder
	In(vals ...string) QueryBuilder
	OrderBy(fields ...string) QueryBuilder
	Asc() QueryBuilder
	Desc() QueryBuilder
	Limit(limit int) QueryBuilder
	Offset(offset int) QueryBuilder
	GroupBy(fields ...string) QueryBuilder
	Having(cond string) QueryBuilder
	Update(tables ...string) QueryBuilder
	Set(kv ...string) QueryBuilder
	Delete(tables ...string) QueryBuilder
	InsertInto(table string, fields ...string) QueryBuilder
	Values(vals ...string) QueryBuilder
	Subquery(sub string, alias string) string
	String() string
}
```

Now we support `Postgress`, `MySQL` and `TiDB`。

&zeroWidthSpace;现在我们支持 `Postgress` 、 `MySQL` 和 `TiDB` 。