+++
title = "高级查询"
date = 2024-02-04T10:00:41+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/model/query/]({{< ref "/beego/mvcIntroduction/models/advancedQueries" >}})

# Advanced Queries 高级查询



## Advanced Queries 高级查询

ORM uses **QuerySeter** to organize queries. Every method that returns **QuerySeter** will give you a new **QuerySeter** object.

​	ORM 使用 QuerySeter 来组织查询。每个返回 QuerySeter 的方法都会给你一个新的 QuerySeter 对象。

Basic Usage:

​	基本用法：

```go
o := orm.NewOrm()

// Get a QuerySeter object. User is table name
qs := o.QueryTable("user")

// Can also use object as table name
user := new(User)
qs = o.QueryTable(user) // return a QuerySeter
```

## expr

expr describes fields and SQL operators in `QuerySeter`.

​	expr 在 `QuerySeter` 中描述字段和 SQL 运算符。

Field combination orders are decided by the relationship of tables. For example, `User` has a foreign key to `Profile`, so if you want to use `Profile.Age` as the condition, you have to use the expression `Profile__Age`. Note that the separator is double under scores `__`. `Expr` can also append operators at the end to execute related SQL. For example, `Profile__Age__gt` represents condition query `Profile.Age > 18`.

​	字段组合顺序由表的关联决定。例如， `User` 有一个外键指向 `Profile` ，所以如果你想使用 `Profile.Age` 作为条件，你必须使用表达式 `Profile__Age` 。请注意，分隔符是双下划线 `__` 。 `Expr` 也可以在末尾追加运算符来执行相关的 SQL。例如， `Profile__Age__gt` 表示条件查询 `Profile.Age > 18` 。

Comments below describe SQL statements that are similar to the expr, but may not be the exact generated results.

​	下面的注释描述了与 expr 相似的 SQL 语句，但可能不是确切的生成结果。

```go
qs.Filter("id", 1) // WHERE id = 1
qs.Filter("profile__age", 18) // WHERE profile.age = 18
qs.Filter("Profile__Age", 18) // key name and field name are both valid
qs.Filter("profile__age", 18) // WHERE profile.age = 18
qs.Filter("profile__age__gt", 18) // WHERE profile.age > 18
qs.Filter("profile__age__gte", 18) // WHERE profile.age >= 18
qs.Filter("profile__age__in", 18, 20) // WHERE profile.age IN (18, 20)

qs.Filter("profile__age__in", 18, 20).Exclude("profile__lt", 1000)
// WHERE profile.age IN (18, 20) AND NOT profile_id < 1000
```

## Operators 运算符

The supported operators:

​	支持的操作符：

- [exact](https://beego.wiki/docs/mvc/model/query/#exact) / [iexact](https://beego.wiki/docs/mvc/model/query/#iexact) equal to
  exact / iexact 等于
- [contains](https://beego.wiki/docs/mvc/model/query/#contains) / [icontains](https://beego.wiki/docs/mvc/model/query/#icontains) contains
  contains / icontains 包含
- [gt / gte](#gt / gte) greater than / greater than or equal to
  [gt / gte](#gt / gte) 大于 / 大于或等于
- [lt / lte](#lt / lte) less than / less than or equal to
  [lt / lte](#lt / lte) 小于 / 小于或等于
- [startswith](https://beego.wiki/docs/mvc/model/query/#startswith) / [istartswith](https://beego.wiki/docs/mvc/model/query/#istartswith) starts with
  startswith / istartswith 以...开头
- [endswith](https://beego.wiki/docs/mvc/model/query/#endswith) / [iendswith](https://beego.wiki/docs/mvc/model/query/#iendswith) ends with
  endswith / iendswith 以...结尾
- [in](https://beego.wiki/docs/mvc/model/query/#in)
- [isnull
  isnull 以 开头的操作符忽略大小写。 exact](https://beego.wiki/docs/mvc/model/query/#isnull)

The operators that start with `i` ignore case.

### exact

Default values of Filter, Exclude and Condition expr

​	Filter、Exclude 和 Condition expr 的默认值

```go
qs.Filter("name", "slene") // WHERE name = 'slene'
qs.Filter("name__exact", "slene") // WHERE name = 'slene'
// using = , case sensitive or not is depending on which collation database table is used
qs.Filter("profile", nil) // WHERE profile_id IS NULL
```

### iexact

```go
qs.Filter("name__iexact", "slene")
// WHERE name LIKE 'slene'
// Case insensitive, will match any name that equals to 'slene'
```

### contains

```go
qs.Filter("name__contains", "slene")
// WHERE name LIKE BINARY '%slene%'
// Case sensitive, only match name that contains 'slene'
```

### icontains

```go
qs.Filter("name__icontains", "slene")
// WHERE name LIKE '%slene%'
// Case insensitive, will match any name that contains 'slene'
```

### in

```go
qs.Filter("profile__age__in", 17, 18, 19, 20)
// WHERE profile.age IN (17, 18, 19, 20)
```

### gt / gte

```go
qs.Filter("profile__age__gt", 17)
// WHERE profile.age > 17

qs.Filter("profile__age__gte", 18)
// WHERE profile.age >= 18
```

### lt / lte

```go
qs.Filter("profile__age__lt", 17)
// WHERE profile.age < 17

qs.Filter("profile__age__lte", 18)
// WHERE profile.age <= 18
```

### startswith

```go
qs.Filter("name__startswith", "slene")
// WHERE name LIKE BINARY 'slene%'
// Case sensitive, only match name that starts with 'slene'
```

### istartswith

```go
qs.Filter("name__istartswith", "slene")
// WHERE name LIKE 'slene%'
// Case insensitive, will match any name that starts with 'slene'
```

### endswith

```go
qs.Filter("name__endswith", "slene")
// WHERE name LIKE BINARY '%slene'
// Case sensitive, only match name that ends with 'slene'
```

### iendswith

```go
qs.Filter("name__iendswith", "slene")
// WHERE name LIKE '%slene'
// Case insensitive, will match any name that ends with 'slene'
```

### isnull

```go
qs.Filter("profile__isnull", true)
qs.Filter("profile_id__isnull", true)
// WHERE profile_id IS NULL

qs.Filter("profile__isnull", false)
// WHERE profile_id IS NOT NULL
```

## Advanced Query API 高级查询 API

QuerySeter is the API of advanced queries. Here are its methods:

​	QuerySeter 是高级查询的 API。以下是其方法：

```go
type QuerySeter interface {
// add condition expression to QuerySeter.
// for example:
//	filter by UserName == 'slene'
//	qs.Filter("UserName", "slene")
//	sql : left outer join profile on t0.id1==t1.id2 where t1.age == 28
//	Filter("profile__Age", 28)
// 	 // time compare
//	qs.Filter("created", time.Now())
Filter(string, ...interface{}) QuerySeter
// add raw sql to querySeter.
// for example:
// qs.FilterRaw("user_id IN (SELECT id FROM profile WHERE age>=18)")
// //sql-> WHERE user_id IN (SELECT id FROM profile WHERE age>=18)
FilterRaw(string, string) QuerySeter
// add NOT condition to querySeter.
// have the same usage as Filter
Exclude(string, ...interface{}) QuerySeter
// set condition to QuerySeter.
// sql's where condition
//	cond := orm.NewCondition()
//	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)
//	//sql-> WHERE T0.`profile_id` IS NOT NULL AND NOT T0.`Status` IN (?) OR T1.`age` >  2000
//	num, err := qs.SetCond(cond1).Count()
SetCond(*Condition) QuerySeter
// get condition from QuerySeter.
// sql's where condition
//  cond := orm.NewCondition()
//  cond = cond.And("profile__isnull", false).AndNot("status__in", 1)
//  qs = qs.SetCond(cond)
//  cond = qs.GetCond()
//  cond := cond.Or("profile__age__gt", 2000)
//  //sql-> WHERE T0.`profile_id` IS NOT NULL AND NOT T0.`Status` IN (?) OR T1.`age` >  2000
//  num, err := qs.SetCond(cond).Count()
GetCond() *Condition
// add LIMIT value.
// args[0] means offset, e.g. LIMIT num,offset.
// if Limit <= 0 then Limit will be set to default limit ,eg 1000
// if QuerySeter doesn't call Limit, the sql's Limit will be set to default limit, eg 1000
//  for example:
//	qs.Limit(10, 2)
//	// sql-> limit 10 offset 2
Limit(limit interface{}, args ...interface{}) QuerySeter
// add OFFSET value
// same as Limit function's args[0]
Offset(offset interface{}) QuerySeter
// add GROUP BY expression
// for example:
//	qs.GroupBy("id")
GroupBy(exprs ...string) QuerySeter
// add ORDER expression.
// "column" means ASC, "-column" means DESC.
// for example:
//	qs.OrderBy("-status")
OrderBy(exprs ...string) QuerySeter
// add FORCE INDEX expression.
// for example:
//	qs.ForceIndex(`idx_name1`,`idx_name2`)
// ForceIndex, UseIndex , IgnoreIndex are mutually exclusive
ForceIndex(indexes ...string) QuerySeter
// add USE INDEX expression.
// for example:
//	qs.UseIndex(`idx_name1`,`idx_name2`)
// ForceIndex, UseIndex , IgnoreIndex are mutually exclusive
UseIndex(indexes ...string) QuerySeter
// add IGNORE INDEX expression.
// for example:
//	qs.IgnoreIndex(`idx_name1`,`idx_name2`)
// ForceIndex, UseIndex , IgnoreIndex are mutually exclusive
IgnoreIndex(indexes ...string) QuerySeter
// set relation model to query together.
// it will query relation models and assign to parent model.
// for example:
//	// will load all related fields use left join .
// 	qs.RelatedSel().One(&user)
//	// will  load related field only profile
//	qs.RelatedSel("profile").One(&user)
//	user.Profile.Age = 32
RelatedSel(params ...interface{}) QuerySeter
// Set Distinct
// for example:
//  o.QueryTable("policy").Filter("Groups__Group__Users__User", user).
//    Distinct().
//    All(&permissions)
Distinct() QuerySeter
// set FOR UPDATE to query.
// for example:
//  o.QueryTable("user").Filter("uid", uid).ForUpdate().All(&users)
ForUpdate() QuerySeter
// return QuerySeter execution result number
// for example:
//	num, err = qs.Filter("profile__age__gt", 28).Count()
Count() (int64, error)
// check result empty or not after QuerySeter executed
// the same as QuerySeter.Count > 0
Exist() bool
// execute update with parameters
// for example:
//	num, err = qs.Filter("user_name", "slene").Update(Params{
//		"Nums": ColValue(Col_Minus, 50),
//	}) // user slene's Nums will minus 50
//	num, err = qs.Filter("UserName", "slene").Update(Params{
//		"user_name": "slene2"
//	}) // user slene's  name will change to slene2
Update(values Params) (int64, error)
// delete from table
// for example:
//	num ,err = qs.Filter("user_name__in", "testing1", "testing2").Delete()
// 	//delete two user  who's name is testing1 or testing2
Delete() (int64, error)
// return a insert queryer.
// it can be used in times.
// example:
// 	i,err := sq.PrepareInsert()
// 	num, err = i.Insert(&user1) // user table will add one record user1 at once
//	num, err = i.Insert(&user2) // user table will add one record user2 at once
//	err = i.Close() //don't forget call Close
PrepareInsert() (Inserter, error)
// query all data and map to containers.
// cols means the columns when querying.
// for example:
//	var users []*User
//	qs.All(&users) // users[0],users[1],users[2] ...
All(container interface{}, cols ...string) (int64, error)
// query one row data and map to containers.
// cols means the columns when querying.
// for example:
//	var user User
//	qs.One(&user) //user.UserName == "slene"
One(container interface{}, cols ...string) error
// query all data and map to []map[string]interface.
// expres means condition expression.
// it converts data to []map[column]value.
// for example:
//	var maps []Params
//	qs.Values(&maps) //maps[0]["UserName"]=="slene"
Values(results *[]Params, exprs ...string) (int64, error)
// query all data and map to [][]interface
// it converts data to [][column_index]value
// for example:
//	var list []ParamsList
//	qs.ValuesList(&list) // list[0][1] == "slene"
ValuesList(results *[]ParamsList, exprs ...string) (int64, error)
// query all data and map to []interface.
// it's designed for one column record set, auto change to []value, not [][column]value.
// for example:
//	var list ParamsList
//	qs.ValuesFlat(&list, "UserName") // list[0] == "slene"
ValuesFlat(result *ParamsList, expr string) (int64, error)
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
}
```

- Every API call that returns **QuerySeter** will give you a new **QuerySeter** object. It won’t affect the previous object.

  ​	每个返回 QuerySeter 的 API 调用都会为您提供一个新的 QuerySeter 对象。它不会影响以前的对象。

- Advanced queries use `Filter` and `Exclude` to do conditional queries. There are two filter rules - contain and exclude

  ​	高级查询使用 `Filter` 和 `Exclude` 来执行条件查询。有两个过滤规则 - 包含和不包含

### Filter 过滤

Used to filter the result for the **include conditions**.

​	用于过滤包含条件的结果。

Use `AND` to connect multiple filters:

​	使用 `AND` 连接多个过滤器：

```go
qs.Filter("profile__isnull", true).Filter("name", "slene")
// WHERE profile_id IS NULL AND name = 'slene'
```

### Exclude 不包含

Used to filter the result for the **exclude conditions**.

​	用于过滤不包含条件的结果。

Use `NOT` to exclude condition Use `AND` to connect multiple filters:

​	使用 `NOT` 排除条件 使用 `AND` 连接多个过滤器：

```go
qs.Exclude("profile__isnull", true).Filter("name", "slene")
// WHERE NOT profile_id IS NULL AND name = 'slene'
```

### SetCond

Custom conditions:

​	自定义条件：

```go
cond := NewCondition()
cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)

qs := orm.QueryTable("user")
qs = qs.SetCond(cond1)
// WHERE ... AND ... AND NOT ... OR ...

cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
qs = qs.SetCond(cond2).Count()
// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )
```

### Limit

Limit maximum returned lines. The second param can set `Offset`

​	限制返回的最大行数。第二个参数可以设置 `Offset`

```go
var DefaultRowsLimit = 1000 // The default limit of ORM is 1000

// LIMIT 1000

qs.Limit(10)
// LIMIT 10

qs.Limit(10, 20)
// LIMIT 10 OFFSET 20

qs.Limit(-1)
// no limit

qs.Limit(-1, 100)
// LIMIT 18446744073709551615 OFFSET 100
// 18446744073709551615 is 1<<64 - 1. Used to set the condition which is no limit but with offset
```

### Offset

Set offset lines:

​	设置偏移行：

```go
qs.Offset(20)
// LIMIT 1000 OFFSET 20
```

### GroupBy

```go
qs.GroupBy("id", "age")
// GROUP BY id,age
```

### OrderBy

Param uses **expr**

​	Param 使用 expr

Using `-` at the beginning of expr stands for order by `DESC`

​	在 expr 的开头使用 `-` 表示按 `DESC` 排序

```go
qs.OrderBy("id", "-profile__age")
// ORDER BY id ASC, profile.age DESC

qs.OrderBy("-profile__age", "profile")
// ORDER BY profile.age DESC, profile_id ASC
```

### ForceIndex

Forcing DB to use the index.

​	强制数据库使用索引。

You need to check your DB whether it support this feature.

​	您需要检查您的数据库是否支持此功能。

```go
qs.ForceIndex(`idx_name1`,`idx_name2`)
```

### UseIndex

Suggest DB to user the index.

​	建议数据库使用索引。

You need to check your DB whether it support this feature.

​	您需要检查您的数据库是否支持此功能。

```go
qs.UseIndex(`idx_name1`,`idx_name2`)
```

### IgnoreIndex

Make DB ignore the index

​	使数据库忽略索引

You need to check your DB whether it support this feature.

​	您需要检查您的数据库是否支持此功能。

```go
qs.IgnoreIndex(`idx_name1`,`idx_name2`)
```

### Distinct

Same as `distinct` statement in sql, return only distinct (different) values

​	与 sql 中的 `distinct` 语句相同，仅返回不同的值

```go
qs.Distinct()
// SELECT DISTINCT
```

### RelatedSel

Relational queries. Param uses **expr**

​	关系查询。Param 使用 expr

```go
var DefaultRelsDepth = 5 // RelatedSel will query for maximum 5 level by default

qs := o.QueryTable("post")

qs.RelatedSel()
// INNER JOIN user ... LEFT OUTER JOIN profile ...

qs.RelatedSel("user")
// INNER JOIN user ... 
// Only query the fields set by expr

// For fields with null attribute will use LEFT OUTER JOIN
```

### Count

Return line count based on the current query

​	根据当前查询返回行数

```go
cnt, err := o.QueryTable("user").Count() // SELECT COUNT(*) FROM USER
fmt.Printf("Count Num: %s, %s", cnt, err)
```

### Exist

```go
exist := o.QueryTable("user").Filter("UserName", "Name").Exist()
fmt.Printf("Is Exist: %s", exist)
```

### Update

Execute batch updating based on the current query

​	根据当前查询执行批量更新

```go
num, err := o.QueryTable("user").Filter("name", "slene").Update(orm.Params{
	"name": "astaxie",
})
fmt.Printf("Affected Num: %s, %s", num, err)
// SET name = "astaixe" WHERE name = "slene"
```

Atom operation add field:

​	原子操作添加字段：

```go
// Assume there is a nums int field in user struct
num, err := o.QueryTable("user").Update(orm.Params{
	"nums": orm.ColValue(orm.Col_Add, 100),
})
// SET nums = nums + 100
```

orm.ColValue supports:

​	orm.ColValue 支持：

```go
Col_Add      // plus
Col_Minus    // minus 
Col_Multiply // multiply 
Col_Except   // divide
```

### Delete

Execute batch deletion based on the current query

​	根据当前查询执行批量删除

```go
num, err := o.QueryTable("user").Filter("name", "slene").Delete()
fmt.Printf("Affected Num: %s, %s", num, err)
// DELETE FROM user WHERE name = "slene"
```

### PrepareInsert

Use a prepared statement to increase inserting speed with multiple inserts.

​	使用预处理语句，通过多次插入提高插入速度。

```go
var users []*User
...
qs := o.QueryTable("user")
i, _ := qs.PrepareInsert()
for _, user := range users {
	id, err := i.Insert(user)
	if err != nil {
		...
	}
}
// PREPARE INSERT INTO user (`name`, ...) VALUES (?, ...)
// EXECUTE INSERT INTO user (`name`, ...) VALUES ("slene", ...)
// EXECUTE ...
// ...
i.Close() // Don't forget to close the statement
```

### All

Return the related ResultSet

​	返回相关的 ResultSet

Param of `All` supports *[]Type and *[]*Type

​	 `All` 的 Param 支持 *[]Type 和 *[]*Type

```go
var users []*User
num, err := o.QueryTable("user").Filter("name", "slene").All(&users)
fmt.Printf("Returned Rows Num: %s, %s", num, err)
```

All / Values / ValuesList / ValuesFlat will be limited by [Limit](https://beego.wiki/docs/mvc/model/query/#limit). 1000 lines by default.

​	所有 / 值 / 值列表 / 值平面都将受到限制。默认情况下为 1000 行。

The returned fields can be specified:

​	可以指定返回的字段：

```go
type Post struct {
	Id      int
	Title   string
	Content string
	Status  int
}

// Only return Id and Title
var posts []Post
o.QueryTable("post").Filter("Status", 1).All(&posts, "Id", "Title")
```

The other fields of the object are set to the default value of the field’s type.

​	对象的其余字段设置为字段类型的默认值。

### One

Try to return one record

​	尝试返回一条记录

```go
var user User
err := o.QueryTable("user").Filter("name", "slene").One(&user)
if err == orm.ErrMultiRows {
	// Have multiple records
	fmt.Printf("Returned Multi Rows Not One")
}
if err == orm.ErrNoRows {
	// No result 
	fmt.Printf("Not row found")
}
```

The returned fields can be specified:

​	可以指定返回的字段：

```go
// Only return Id and Title
var post Post
o.QueryTable("post").Filter("Content__istartswith", "prefix string").One(&post, "Id", "Title")
```

The other fields of the object are set to the default value of the fields’ type.

​	对象的其他字段设置为字段类型的默认值。

### Values 值

Return key => value of result set

​	返回键 => 结果集的值

key is Field name in Model. value type if string.

​	键是模型中的字段名称。值类型为字符串。

```go
var maps []orm.Params
num, err := o.QueryTable("user").Values(&maps)
if err == nil {
	fmt.Printf("Result Nums: %d\n", num)
	for _, m := range maps {
		fmt.Println(m["Id"], m["Name"])
	}
}
```

Return specific fields:

​	返回特定字段：

**TODO**: doesn’t support recursive query. **RelatedSel** return Values directly

​	TODO：不支持递归查询。RelatedSel 直接返回值

But it can specify the value needed by expr.

​	但它可以通过 expr 指定所需的值。

```go
var maps []orm.Params
num, err := o.QueryTable("user").Values(&maps, "id", "name", "profile", "profile__age")
if err == nil {
	fmt.Printf("Result Nums: %d\n", num)
	for _, m := range maps {
		fmt.Println(m["Id"], m["Name"], m["Profile"], m["Profile__Age"])
    // There is no complicated nesting data in the map
	}
}
```

### ValuesList 值列表

The result set will be stored as a slice

​	结果集将存储为切片

The order of the result is same as the Fields order in the Model definition.

​	结果的顺序与模型定义中的字段顺序相同。

The values are saved as strings.

​	这些值以字符串形式保存。

```go
var lists []orm.ParamsList
num, err := o.QueryTable("user").ValuesList(&lists)
if err == nil {
	fmt.Printf("Result Nums: %d\n", num)
	for _, row := range lists {
		fmt.Println(row)
	}
}
```

It can return specific fields by setting expr.

​	它可以通过设置 expr 返回特定字段。

```go
var lists []orm.ParamsList
num, err := o.QueryTable("user").ValuesList(&lists, "name", "profile__age")
if err == nil {
	fmt.Printf("Result Nums: %d\n", num)
	for _, row := range lists {
		fmt.Printf("Name: %s, Age: %s\m", row[0], row[1])
	}
}
```

### ValuesFlat

Only returns a single values slice of a specific field.

​	仅返回特定字段的单个值切片。

```go
var list orm.ParamsList
num, err := o.QueryTable("user").ValuesFlat(&list, "name")
if err == nil {
	fmt.Printf("Result Nums: %d\n", num)
	fmt.Printf("All User Names: %s", strings.Join(list, ", "))
}
```

## Relational Query 关系查询

Let’s see how to do a Relational Query by looking at [Model Definition]({{< ref "/beego/mvcIntroduction/models/ormUsage" >}})

​	让我们通过查看模型定义来了解如何执行关系查询

#### User and Profile is OnToOne relation User 和 Profile 是 OnToOne 关系

Query Profile by known User object:

​	通过已知的 User 对象查询 Profile：

```go
user := &User{Id: 1}
o.Read(user)
if user.Profile != nil {
	o.Read(user.Profile)
}
```

Cascaded query directly:

​	直接级联查询：

```go
user := &User{}
o.QueryTable("user").Filter("Id", 1).RelatedSel().One(user)
// Get Profile automatically
fmt.Println(user.Profile)
// Because In Profile we defined reverse relation User, Profile's User is also auto assigned. Can directly use:
fmt.Println(user.Profile.User)
```

Reverse finding Profile by User:

​	通过 User 反向查找 Profile：

```go
var profile Profile
err := o.QueryTable("profile").Filter("User__Id", 1).One(&profile)
if err == nil {
	fmt.Println(profile)
}
```

#### Post and User are ManyToOne relation. i.e.: ForeignKey is User Post 和 User 是多对一关系。即：外键是 User

```go
type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"`
}
var posts []*Post
num, err := o.QueryTable("post").Filter("User", 1).RelatedSel().All(&posts)
if err == nil {
	fmt.Printf("%d posts read\n", num)
	for _, post := range posts {
		fmt.Printf("Id: %d, UserName: %d, Title: %s\n", post.Id, post.User.UserName, post.Title)
	}
}
```

Query related User by Post.Title:

​	通过 Post.Title 查询相关 User：

While RegisterModel, ORM will create reverse relation for Post in User. So it can query directly:

​	在 RegisterModel 时，ORM 会在 User 中为 Post 创建反向关系。因此它可以直接查询：

```go
var user User
err := o.QueryTable("user").Filter("Post__Title", "The Title").Limit(1).One(&user)
if err == nil {
	fmt.Printf(user)
}
```

#### Post and Tag are ManyToMany relation Post 和 Tag 是多对多关系

After setting rel(m2m), ORM will create connecting table automatically.

​	设置 rel(m2m) 后，ORM 会自动创建连接表。

```go
type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"`
}
type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}
```

Query which post used the tag with tag name:

​	查询使用带有标签名称的标签的帖子：

```go
var posts []*Post
num, err := dORM.QueryTable("post").Filter("Tags__Tag__Name", "golang").All(&posts)
```

Query how many tags does the post have with post title:

​	查询帖子标题中包含多少个标签：

```go
var tags []*Tag
num, err := dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)
```

## Load Related Field 加载相关字段

LoadRelated is used to load relation field of model. Including all rel/reverse - one/many relation.

​	LoadRelated 用于加载模型的关系字段。包括所有 rel/reverse - one/many 关系。

Load ManyToMany relation field

​	加载多对多关系字段

```go
// load related Tags
post := Post{Id: 1}
err := o.Read(&post)
num, err := o.LoadRelated(&post, "Tags")
// Load related Posts
tag := Tag{Id: 1}
err := o.Read(&tag)
num, err := o.LoadRelated(&tag, "Posts")
```

User is the ForeignKey of Post. Load related ReverseMany

​	用户是 Post 的外键。加载相关 ReverseMany

```go
type User struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

user := User{Id: 1}
err := dORM.Read(&user)
num, err := dORM.LoadRelated(&user, "Posts")
for _, post := range user.Posts {
	//...
}
```

## Handling ManyToMany relation 处理 ManyToMany 关系

```go
// QueryM2Mer model to model query struct
// all operations are on the m2m table only, will not affect the origin model table
type QueryM2Mer interface {
	// add models to origin models when creating queryM2M.
	// example:
	// 	m2m := orm.QueryM2M(post,"Tag")
	// 	m2m.Add(&Tag1{},&Tag2{})
	//  	for _,tag := range post.Tags{}{ ... }
	// param could also be any of the follow
	// 	[]*Tag{{Id:3,Name: "TestTag1"}, {Id:4,Name: "TestTag2"}}
	//	&Tag{Id:5,Name: "TestTag3"}
	//	[]interface{}{&Tag{Id:6,Name: "TestTag4"}}
	// insert one or more rows to m2m table
	// make sure the relation is defined in post model struct tag.
	Add(...interface{}) (int64, error)
	// remove models following the origin model relationship
	// only delete rows from m2m table
	// for example:
	// tag3 := &Tag{Id:5,Name: "TestTag3"}
	// num, err = m2m.Remove(tag3)
	Remove(...interface{}) (int64, error)
	// check model is existed in relationship of origin model
	Exist(interface{}) bool
	// clean all models in related of origin model
	Clear() (int64, error)
	// count all related models of origin model
	Count() (int64, error)
}
```

Create a QueryM2Mer object

​	创建一个 QueryM2Mer 对象

```go
o := orm.NewOrm()
post := Post{Id: 1}
m2m := o.QueryM2M(&post, "Tags")
// In the first param object must have primary key
// The second param is the M2M field will work with
// API of QueryM2Mer will used to Post with id equals 1
```

### QueryM2Mer Add QueryM2Mer 添加

```go
tag := &Tag{Name: "golang"}
o.Insert(tag)

num, err := m2m.Add(tag)
if err == nil {
	fmt.Println("Added nums: ", num)
}
```

Add supports many types: Tag *Tag []*Tag []Tag []interface{}

​	添加支持多种类型：Tag *Tag []*Tag []Tag []interface{}

```go
var tags []*Tag
...
// After reading tags
...
num, err := m2m.Add(tags)
if err == nil {
	fmt.Println("Added nums: ", num)
}
// It can pass multiple params
// m2m.Add(tag1, tag2, tag3)
```

### QueryM2Mer Remove QueryM2Mer 删除

Remove tag from M2M relation:

​	从 M2M 关系中删除标签：

Remove supports many types: Tag *Tag []*Tag []Tag []interface{}

​	删除支持多种类型：Tag *Tag []*Tag []Tag []interface{}

```go
var tags []*Tag
...
// After reading tags
...
num, err := m2m.Remove(tags)
if err == nil {
	fmt.Println("Removed nums: ", num)
}
// It can pass multiple params
// m2m.Remove(tag1, tag2, tag3)
```

### QueryM2Mer Exist QueryM2Mer 存在

Test if Tag is in M2M relation

​	测试标签是否在 M2M 关系中

```go
if m2m.Exist(&Tag{Id: 2}) {
	fmt.Println("Tag Exist")
}
```

### QueryM2Mer Clear QueryM2Mer 清除

Clear all M2M relation

​	清除所有 M2M 关系

```go
nums, err := m2m.Clear()
if err == nil {
	fmt.Println("Removed Tag Nums: ", nums)
}
```

### QueryM2Mer Count QueryM2Mer 计数

Count the number of Tags

​	计算标签数

```go
nums, err := m2m.Count()
if err == nil {
	fmt.Println("Total Nums: ", nums)
}
```
