+++
title = "使用原生 SQL 查询"
date = 2024-02-04T10:00:58+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/model/rawsql/](https://beego.wiki/docs/mvc/model/rawsql/)

# Raw SQL to query 使用 Raw SQL 查询



## Raw SQL to query 查询原始 SQL

- Using Raw SQL to query doesn’t require an ORM definition
  使用 Raw SQL 查询不需要 ORM 定义
- Multiple databases support `?` as placeholders and auto convert.
  多个数据库支持 `?` 作为占位符并自动转换。
- The params of query support Model Struct, Slice and Array
  查询的参数支持模型结构、切片和数组

```go
ids := []int{1, 2, 3}
p.Raw("SELECT name FROM user WHERE id IN (?, ?, ?)", ids)
```

Create a **RawSeter**

&zeroWidthSpace;创建 RawSeter

```go
o := NewOrm()
var r RawSeter
r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")
// RawSeter raw query seter
// create From Ormer.Raw
// for example:
//  sql := fmt.Sprintf("SELECT %sid%s,%sname%s FROM %suser%s WHERE id = ?",Q,Q,Q,Q,Q,Q)
//  rs := Ormer.Raw(sql, 1)
type RawSeter interface {
	// execute sql and get result
	Exec() (sql.Result, error)
	// query data and map to container
	// for example:
	//	var name string
	//	var id int
	//	rs.QueryRow(&id,&name) // id==2 name=="slene"
	QueryRow(containers ...interface{}) error

	// query data rows and map to container
	//	var ids []int
	//	var names []int
	//	query = fmt.Sprintf("SELECT 'id','name' FROM %suser%s", Q, Q)
	//	num, err = dORM.Raw(query).QueryRows(&ids,&names) // ids=>{1,2},names=>{"nobody","slene"}
	QueryRows(containers ...interface{}) (int64, error)
	SetArgs(...interface{}) RawSeter
	// query data to []map[string]interface
	// see QuerySeter's Values
	Values(container *[]Params, cols ...string) (int64, error)
	// query data to [][]interface
	// see QuerySeter's ValuesList
	ValuesList(container *[]ParamsList, cols ...string) (int64, error)
	// query data to []interface
	// see QuerySeter's ValuesFlat
	ValuesFlat(container *ParamsList, cols ...string) (int64, error)
	// query all rows into map[string]interface with specify key and value column name.
	// keyCol = "name", valueCol = "value"
	// table data
	// name  | value
	// total | 100
	// found | 200
	// to map[string]interface{}{
	// 	"total": 100,
	// 	"found": 200,
	// }
	RowsToMap(result *Params, keyCol, valueCol string) (int64, error)
	// query all rows into struct with specify key and value column name.
	// keyCol = "name", valueCol = "value"
	// table data
	// name  | value
	// total | 100
	// found | 200
	// to struct {
	// 	Total int
	// 	Found int
	// }
	RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error)

	// return prepared raw statement for used in times.
	// for example:
	// 	pre, err := dORM.Raw("INSERT INTO tag (name) VALUES (?)").Prepare()
	// 	r, err := pre.Exec("name1") // INSERT INTO tag (name) VALUES (`name1`)
	Prepare() (RawPreparer, error)
}
```

#### Exec

Run sql query and return [sql.Result](http://gowalker.org/database/sql#Result) object

&zeroWidthSpace;运行 SQL 查询并返回 sql.Result 对象

```go
res, err := o.Raw("UPDATE user SET name = ?", "your").Exec()
if err == nil {
	num, _ := res.RowsAffected()
	fmt.Println("mysql row affected nums: ", num)
}
```

#### QueryRow

QueryRow and QueryRows support high-level sql mapper.

&zeroWidthSpace;QueryRow 和 QueryRows 支持高级 SQL 映射器。

Supports struct:

&zeroWidthSpace;支持结构体：

```go
type User struct {
	Id   int
	Name string
}

var user User
err := o.Raw("SELECT id, name FROM user WHERE id = ?", 1).QueryRow(&user)
```

> from Beego 1.1.0 remove multiple struct support [ISSUE 384](https://github.com/beego/beego/issues/384)
>
> &zeroWidthSpace;从 Beego 1.1.0 中移除对多个结构体支持的问题 384

#### QueryRows

QueryRows supports the same mapping rules as QueryRow but all of them are slice.

&zeroWidthSpace;QueryRows 支持与 QueryRow 相同的映射规则，但它们都是切片。

```go
type User struct {
	Id   int
	Name string
}

var users []User
num, err := o.Raw("SELECT id, name FROM user WHERE id = ?", 1).QueryRows(&users)
if err == nil {
	fmt.Println("user nums: ", num)
}
```

> from Beego 1.1.0 remove multiple struct support [ISSUE 384](https://github.com/beego/beego/issues/384)
>
> &zeroWidthSpace;从 Beego 1.1.0 中移除对多个结构体支持的问题 384

#### SetArgs

Changing args param in Raw(sql, args…) can return a new RawSeter.

&zeroWidthSpace;在 Raw(sql, args…) 中更改 args 参数可以返回一个新的 RawSeter。

It can reuse the same SQL query but different params.

&zeroWidthSpace;它可以重用相同的 SQL 查询，但参数不同。

```go
res, err := r.SetArgs("arg1", "arg2").Exec()
res, err := r.SetArgs("arg1", "arg2").Exec()
...
```

#### Values / ValuesList / ValuesFlat

The resultSet values returned by Raw SQL query are `string`. NULL field will return empty string ``

&zeroWidthSpace;Raw SQL 查询返回的结果集值是 `string` 。NULL 字段将返回空字符串 ``

> from Beego 1.1.0 Values, ValuesList, ValuesFlat. The returned fields can be specified. Generally you don’t need to specify. Because the field names are already defined in your SQL.
>
> &zeroWidthSpace;从 Beego 1.1.0 Values、ValuesList、ValuesFlat。可以指定返回的字段。通常您无需指定。因为字段名称已在您的 SQL 中定义。

#### Values 值

The key => value pairs of resultSet:

&zeroWidthSpace;结果集的键值对：

```go
var maps []orm.Params
num, err := o.Raw("SELECT user_name FROM user WHERE status = ?", 1).Values(&maps)
if err == nil && num > 0 {
	fmt.Println(maps[0]["user_name"]) // slene
}
```

#### ValuesList 值列表

slice of resultSet

&zeroWidthSpace;结果集切片

```go
var lists []orm.ParamsList
num, err := o.Raw("SELECT user_name FROM user WHERE status = ?", 1).ValuesList(&lists)
if err == nil && num > 0 {
	fmt.Println(lists[0][0]) // slene
}
```

#### ValuesFlat

Return slice of a single field:

&zeroWidthSpace;返回单个字段的切片：

```go
var list orm.ParamsList
num, err := o.Raw("SELECT id FROM user WHERE id < ?", 10).ValuesFlat(&list)
if err == nil && num > 0 {
	fmt.Println(list) // []{"1","2","3",...}
}
```

#### RowsToMap

SQL query results

&zeroWidthSpace;SQL 查询结果

| name 名称    | value 值 |
| ------------ | -------- |
| total 总数   | 100      |
| found 已找到 | 200      |

map rows results to map

&zeroWidthSpace;将映射行结果映射到映射

```go
res := make(orm.Params)
nums, err := o.Raw("SELECT name, value FROM options_table").RowsToMap(&res, "name", "value")
// res is a map[string]interface{}{
//	"total": 100,
//	"found": 200,
// }
```

#### RowsToStruct

SQL query results

&zeroWidthSpace;SQL 查询结果

| name 名称    | value 值 |
| ------------ | -------- |
| total 总计   | 100      |
| found 已找到 | 200      |

map rows results to struct

&zeroWidthSpace;将映射行结果映射到结构

```go
type Options struct {
	Total int
	Found int
}

res := new(Options)
nums, err := o.Raw("SELECT name, value FROM options_table").RowsToStruct(res, "name", "value")
fmt.Println(res.Total) // 100
fmt.Println(res.Found) // 200
```

> support name conversion: snake -> camel, eg: SELECT user_name … to your struct field UserName.
>
> &zeroWidthSpace;支持名称转换：snake -> camel，例如：SELECT user_name … 到您的结构字段 UserName。

#### Prepare 准备

Prepare once and exec multiple times to improve the speed of batch execution.

&zeroWidthSpace;一次准备，多次执行，以提高批处理执行速度。

```go
p, err := o.Raw("UPDATE user SET name = ? WHERE name = ?").Prepare()
res, err := p.Exec("testing", "slene")
res, err  = p.Exec("testing", "astaxie")
...
...
p.Close() // Don't forget to close the prepare.
```