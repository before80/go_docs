+++
title = "查询"
date = 2023-10-28T14:25:43+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/query.html](https://gorm.io/docs/query.html)

## 检索单个对象 Retrieving a single object

GORM provides `First`, `Take`, `Last` methods to retrieve a single object from the database, it adds `LIMIT 1` condition when querying the database, and it will return the error `ErrRecordNotFound` if no record is found.

​	GORM提供了`First`、`Take`和`Last`方法来从数据库中检索单个对象。当查询数据库时，它会添加`LIMIT 1`条件，如果没有找到记录，它将返回错误`ErrRecordNotFound`。

``` go
// 根据主键顺序获取第一条记录 Get the first record ordered by primary key
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

// 获取一条记录，没有指定顺序 Get one record, no specified order
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// 获取最后一条记录，按主键降序排序 Get last record, ordered by primary key desc
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

result := db.First(&user)
result.RowsAffected // 返回找到的记录数 returns count of records found
result.Error        // 返回错误或nil returns error or nil

// 检查错误ErrRecordNotFound check error ErrRecordNotFound
errors.Is(result.Error, gorm.ErrRecordNotFound)
```

> If you want to avoid the `ErrRecordNotFound` error, you could use `Find` like `db.Limit(1).Find(&user)`, the `Find` method accepts both struct and slice data
>
> ​	如果你想避免`ErrRecordNotFound`错误，你可以使用`Find`方法，例如`db.Limit(1).Find(&user)`。`Find`方法接受结构体和切片数据作为参数。

> Using `Find` without a limit for single object `db.Find(&user)` will query the full table and return only the first object which is not performant and nondeterministic
>
> ​	对于单个对象的查找，如果不使用限制条件直接使用`Find`方法，如`db.Find(&user)`，将会查询整个表并仅返回第一个对象，这既不高效也不具有确定性。

The `First` and `Last` methods will find the first and last record (respectively) as ordered by primary key. They only work when a pointer to the destination struct is passed to the methods as argument or when the model is specified using `db.Model()`. Additionally, if no primary key is defined for relevant model, then the model will be ordered by the first field. For example:

​	`First` 和 `Last` 方法将分别找到按主键排序的第一个和最后一个记录。它们仅在将目标结构体的指针作为参数传递给方法时有效，或者使用 `db.Model()` 指定模型时有效。此外，如果相关模型没有定义主键，则模型将按照第一个字段进行排序。例如：

``` go
var user User
var users []User

// 工作是因为传递了目标结构体 works because destination struct is passed in
db.First(&user)
// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

// 使用 `db.Model()` 指定模型 works because model is specified using `db.Model()`
result := map[string]interface{}{}
db.Model(&User{}).First(&result)
// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

// 不起作用 doesn't work
result := map[string]interface{}{}
db.Table("users").First(&result)

// 使用 Take works with Take
result := map[string]interface{}{}
db.Table("users").Take(&result)

// 没有定义主键，结果将按照第一个字段（即，`Code`）排序 no primary key defined, results will be ordered by first field (i.e., `Code`)
type Language struct {
  Code string
  Name string
}
db.First(&Language{})
// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1
```

### 使用主键检索对象 Retrieving objects with primary key

Objects can be retrieved using primary key by using [Inline Conditions](https://gorm.io/docs/query.html#inline_conditions) if the primary key is a number. When working with strings, extra care needs to be taken to avoid SQL Injection; check out [Security](https://gorm.io/docs/security.html) section for details.

​	可以使用主键通过内联条件检索对象，如果主键是数字。当处理字符串时，需要特别注意以避免 SQL 注入；请参阅 [Security](https://gorm.io/docs/security.html) 部分以获取详细信息。

``` go
db.First(&user, 10)
// SELECT * FROM users WHERE id = 10;

db.First(&user, "10")
// SELECT * FROM users WHERE id = 10;

db.Find(&users, []int{1,2,3})
// SELECT * FROM users WHERE id IN (1,2,3);
```

If the primary key is a string (for example, like a uuid), the query will be written as follows:

​	如果主键是字符串（例如，类似于 uuid），查询将如下所示：

``` go
db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
```

When the destination object has a primary value, the primary key will be used to build the condition, for example:

​	当目标对象具有主值时，主键将用于构建条件，例如：

``` go
var user = User{ID: 10}
db.First(&user)
// SELECT * FROM users WHERE id = 10;

var result User
db.Model(User{ID: 10}).First(&result)
// SELECT * FROM users WHERE id = 10;
```

> **NOTE:** If you use gorm’s specific field types like `gorm.DeletedAt`, it will run a different query for retrieving object/s.
>
> **注意：** 如果你使用gorm的特定字段类型，如`gorm.DeletedAt`，它将运行一个不同的查询以检索对象/s。

``` go
type User struct {
  ID           string `gorm:"primarykey;size:16"`
  Name         string `gorm:"size:24"`
  DeletedAt    gorm.DeletedAt `gorm:"index"`
}

var user = User{ID: 15}
db.First(&user)
//  SELECT * FROM `users` WHERE `users`.`id` = '15' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
```

## 检索所有对象 Retrieving all objects

``` go
// 获取所有记录 Get all records
result := db.Find(&users)
// SELECT * FROM users;

result.RowsAffected // 返回找到的记录数，等于`len(users)` returns found records count, equals `len(users)`
result.Error        // 返回错误 returns error
```

## 条件 Conditions

### 字符串条件 String Conditions

``` go
// 获取第一个匹配的记录 Get first matched record
db.Where("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

// 获取所有匹配的记录 Get all matched records
db.Where("name <> ?", "jinzhu").Find(&users)
// SELECT * FROM users WHERE name <> 'jinzhu';

// IN
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)
// SELECT * FROM users WHERE name LIKE '%jin%';

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)
// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
```

> If the object’s primary key has been set, then condition query wouldn’t cover the value of primary key but use it as a ‘and’ condition. For example:
>
> ​	如果对象的主键已设置，则条件查询不会覆盖主键的值，而是将其用作"and"条件。例如：
>
> ``` go
> var user = User{ID: 10}
> db.Where("id = ?", 20).First(&user)
> // SELECT * FROM users WHERE id = 10 and id = 20 ORDER BY id ASC LIMIT 1
> ``` go
> 
> This query would give `record not found` Error. So set the primary key attribute such as `id` to nil before you want to use the variable such as `user` to get new value from database.

### 结构体 & 映射条件 Struct & Map Conditions

``` go
// 结构体 Struct
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

// 映射 Map
db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

// 主键切片 Slice of primary keys
db.Where([]int64{20, 21, 22}).Find(&users)
// SELECT * FROM users WHERE id IN (20, 21, 22);
```

> **NOTE** When querying with struct, GORM will only query with non-zero fields, that means if your field’s value is `0`, `''`, `false` or other [zero values](https://tour.golang.org/basics/12), it won’t be used to build query conditions, for example:
>
> **注意** 当使用结构体进行查询时，GORM只会查询非零字段，这意味着如果字段的值是`0`、`''`、`false`或其他零值，它将不会被用于构建查询条件，例如：

``` go
db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu";
```

To include zero values in the query conditions, you can use a map, which will include all key-values as query conditions, for example:

​	要包含零值在查询条件中，可以使用映射，这将包含所有键值作为查询条件，例如：

``` go
db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
```

For more details, see [Specify Struct search fields](https://gorm.io/docs/query.html#specify_search_fields).

​	有关更多详细信息，请参阅[指定结构体搜索字段](#指定结构体搜索字段-specify-struct-search-fields)。

### 指定结构体搜索字段 Specify Struct search fields

When searching with struct, you can specify which particular values from the struct to use in the query conditions by passing in the relevant field name or the dbname to `Where()`, for example:

​	当使用结构体进行搜索时，可以通过将相关字段名称或dbname传递给`Where()`来指定要在查询条件中使用的结构体的哪些特定值，例如：

``` go
db.Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

db.Where(&User{Name: "jinzhu"}, "Age").Find(&users)
// SELECT * FROM users WHERE age = 0;
```

### 内联条件 Inline Condition

Query conditions can be inlined into methods like `First` and `Find` in a similar way to `Where`.

​	查询条件可以像`Where`方法一样内联到`First`和`Find`等方法中。

``` go
// 如果主键是整数类型之外的非整数类型，则通过主键获取 Get by primary key if it were a non-integer type
db.First(&user, "id = ?", "string_primary_key")
// SELECT * FROM users WHERE id = 'string_primary_key';

// 纯SQL Plain SQL
db.Find(&user, "name = ?", "jinzhu")
// SELECT * FROM users WHERE name = "jinzhu";

db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

// 结构体 Struct
db.Find(&users, User{Age: 20})
// SELECT * FROM users WHERE age = 20;

// 映射 Map
db.Find(&users, map[string]interface{}{"age": 20})
// SELECT * FROM users WHERE age = 20;
```

### Not条件 Not Conditions

Build NOT conditions, works similar to `Where`

​	构建非条件，与`Where`类似。

``` go
db.Not("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

// Not In
db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

// Struct
db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

// Not In slice of primary keys
db.Not([]int64{1,2,3}).First(&user)
// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
```

### Or条件 Or Conditions

``` go
db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

// Struct
db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

// Map
db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);
```

For more complicated SQL queries. please also refer to [Group Conditions in Advanced Query](https://gorm.io/docs/advanced_query.html#group_conditions).

​	对于更复杂的SQL查询，请还参考[高级查询中的分组条件](https://gorm.io/docs/advanced_query.html#group_conditions)。

## 选择特定字段 Selecting Specific Fields

`Select` allows you to specify the fields that you want to retrieve from database. Otherwise, GORM will select all fields by default.

​	`Select` 允许你指定要从数据库检索的字段。否则，GORM 将默认选择所有字段。

``` go
db.Select("name", "age").Find(&users)
// SELECT name, age FROM users;

db.Select([]string{"name", "age"}).Find(&users)
// SELECT name, age FROM users;

db.Table("users").Select("COALESCE(age,?)", 42).Rows()
// SELECT COALESCE(age,'42') FROM users;
```

Also check out [Smart Select Fields](https://gorm.io/docs/advanced_query.html#smart_select)    

​	也请查看 [Smart Select Fields](https://gorm.io/docs/advanced_query.html#smart_select)。

## Order

Specify order when retrieving records from the database

​	在从数据库检索记录时指定顺序。

``` go
db.Order("age desc, name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;

// 多个顺序 Multiple orders
db.Order("age desc").Order("name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;

db.Clauses(clause.OrderBy{
  Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
}).Find(&User{})
// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
```

## Limit & Offset

`Limit` specify the max number of records to retrieve

​	`Limit` 指定要检索的最大记录数

`Offset` specify the number of records to skip before starting to return the records

​	`Offset` 指定在开始返回记录之前要跳过的记录数

``` go
db.Limit(3).Find(&users)
// SELECT * FROM users LIMIT 3;

// Cancel limit condition with -1
db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
// SELECT * FROM users LIMIT 10; (users1)
// SELECT * FROM users; (users2)

db.Offset(3).Find(&users)
// SELECT * FROM users OFFSET 3;

db.Limit(10).Offset(5).Find(&users)
// SELECT * FROM users OFFSET 5 LIMIT 10;

// Cancel offset condition with -1
db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
// SELECT * FROM users OFFSET 10; (users1)
// SELECT * FROM users; (users2)
```

Refer to [Pagination](https://gorm.io/docs/scopes.html#pagination) for details on how to make a paginator

​	关于如何制作分页器，请参阅 [Pagination](https://gorm.io/docs/scopes.html#pagination)。

## Group By & Having

``` go
type result struct {
  Date  time.Time
  Total int
}

db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1


db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
defer rows.Close()
for rows.Next() {
  ...
}

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
defer rows.Close()
for rows.Next() {
  ...
}

type Result struct {
  Date  time.Time
  Total int64
}
db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)
```

## Distinct

Selecting distinct values from the model

​	从模型中选择不重复的值

``` go
db.Distinct("name", "age").Order("name, age desc").Find(&results)
```

`Distinct` works with [`Pluck`](https://gorm.io/docs/advanced_query.html#pluck) and [`Count`](https://gorm.io/docs/advanced_query.html#count) too

​	`Distinct` 可以与 [`Pluck`](https://gorm.io/docs/advanced_query.html#pluck) 和 [`Count`](https://gorm.io/docs/advanced_query.html#count) 一起使用。

## Joins

Specify Joins conditions

​	指定连接条件

``` go
type result struct {
  Name  string
  Email string
}

db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
for rows.Next() {
  ...
}

db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

// multiple joins with parameter
db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
```

### Joins 预加载 Joins Preloading

You can use `Joins` eager loading associations with a single SQL, for example:

​	你可以使用 `Joins` 在单个 SQL 中预加载关联，例如：

``` go
db.Joins("Company").Find(&users)
// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;

// inner join
db.InnerJoins("Company").Find(&users)
// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` INNER JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;
```

带条件的Join Join with conditions

``` go
db.Joins("Company", db.Where(&Company{Alive: true})).Find(&users)
// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id` AND `Company`.`alive` = true;
```

For more details, please refer to [Preloading (Eager Loading)](https://gorm.io/docs/preload.html).

​	更多详细信息，请参阅 [预加载（预加载）](https://gorm.io/docs/preload.html)。

### Joins 派生表 Joins a Derived Table

You can also use `Joins` to join a derived table.

​	你也可以使用 `Joins` 来连接派生表。

``` go
type User struct {
  Id  int
  Age int
}

type Order struct {
  UserId     int
  FinishedAt *time.Time
}

query := db.Table("order").Select("MAX(order.finished_at) as latest").Joins("left join user user on order.user_id = user.id").Where("user.age > ?", 18).Group("order.user_id")
db.Model(&Order{}).Joins("join (?) q on order.finished_at = q.latest", query).Scan(&results)
// SELECT `order`.`user_id`,`order`.`finished_at` FROM `order` join (SELECT MAX(order.finished_at) as latest FROM `order` left join user user on order.user_id = user.id WHERE user.age > 18 GROUP BY `order`.`user_id`) q on order.finished_at = q.latest
```

## Scan

Scanning results into a struct works similarly to the way we use `Find`

​	将扫描结果存入结构体的工作方式类似于我们使用`Find`的方式。

``` go
type Result struct {
  Name string
  Age  int
}

var result Result
db.Table("users").Select("name", "age").Where("name = ?", "Antonio").Scan(&result)

// Raw SQL
db.Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)
```