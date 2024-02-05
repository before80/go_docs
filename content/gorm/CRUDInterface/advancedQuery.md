+++
title = "高级查询"
date = 2023-10-28T14:26:01+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/advanced_query.html](https://gorm.io/docs/advanced_query.html)

## 智能选择字段 Smart Select Fields

GORM allows selecting specific fields with [`Select`](https://gorm.io/docs/query.html), if you often use this in your application, maybe you want to define a smaller struct for API usage which can select specific fields automatically, for example:

​	GORM 允许使用 [`Select`](https://gorm.io/docs/query.html) 来选择特定的字段，如果你在应用程序中经常使用这个功能，也许你想要为 API 定义一个较小的结构体，以便自动选择特定的字段，例如：

``` go
type User struct {
  ID     uint
  Name   string
  Age    int
  Gender string
  // 数百个字段 hundreds of fields
}

type APIUser struct {
  ID   uint
  Name string
}

// 查询时自动选择 `id` 和 `name` Select `id`, `name` automatically when querying
db.Model(&User{}).Limit(10).Find(&APIUser{})
// SELECT `id`, `name` FROM `users` LIMIT 10
```

> **NOTE** `QueryFields` mode will select by all fields’ name for current model
>
> **注意** `QueryFields` 模式将根据当前模型的所有字段名称进行选择

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  QueryFields: true,
})

db.Find(&user)
// SELECT `users`.`name`, `users`.`age`, ... FROM `users` // with this option

// 会话模式 Session Mode
db.Session(&gorm.Session{QueryFields: true}).Find(&user)
// SELECT `users`.`name`, `users`.`age`, ... FROM `users`
```

## 锁定（FOR UPDATE）Locking (FOR UPDATE)

GORM supports different types of locks, for example:

​	GORM 支持不同类型的锁，例如：

``` go
db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)
// SELECT * FROM `users` FOR UPDATE

db.Clauses(clause.Locking{
  Strength: "SHARE",
  Table: clause.Table{Name: clause.CurrentTable},
}).Find(&users)
// SELECT * FROM `users` FOR SHARE OF `users`

db.Clauses(clause.Locking{
  Strength: "UPDATE",
  Options: "NOWAIT",
}).Find(&users)
// SELECT * FROM `users` FOR UPDATE NOWAIT
```

Refer [Raw SQL and SQL Builder](https://gorm.io/docs/sql_builder.html) for more detail

​	更多详细信息请参考 [Raw SQL and SQL Builder](../rawSQLAndSQLBuilder)

## 子查询 SubQuery

A subquery can be nested within a query, GORM can generate subquery when using a `*gorm.DB` object as param

​	子查询可以在查询中使用，GORM 在使用 `*gorm.DB` 对象作为参数时会生成子查询。

``` go
db.Where("amount > (?)", db.Table("orders").Select("AVG(amount)")).Find(&orders)
// SELECT * FROM "orders" WHERE amount > (SELECT AVG(amount) FROM "orders");

subQuery := db.Select("AVG(age)").Where("name LIKE ?", "name%").Table("users")
db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)
// SELECT AVG(age) as avgage FROM `users` GROUP BY `name` HAVING AVG(age) > (SELECT AVG(age) FROM `users` WHERE name LIKE "name%")
```

## From 子查询 From SubQuery

GORM allows you using subquery in FROM clause with the method `Table`, for example:

​	GORM允许你在FROM子句中使用子查询，使用`Table`方法，例如：

``` go
db.Table("(?) as u", db.Model(&User{}).Select("name", "age")).Where("age = ?", 18).Find(&User{})
// SELECT * FROM (SELECT `name`,`age` FROM `users`) as u WHERE `age` = 18

subQuery1 := db.Model(&User{}).Select("name")
subQuery2 := db.Model(&Pet{}).Select("name")
db.Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&User{})
// SELECT * FROM (SELECT `name` FROM `users`) as u, (SELECT `name` FROM `pets`) as p
```

## Group条件 Group Conditions

Easier to write complicated SQL query with Group Conditions

​	编写复杂的 SQL 查询更容易使用分组条件。

``` go
db.Where(
  db.Where("pizza = ?", "pepperoni").Where(db.Where("size = ?", "small").Or("size = ?", "medium")),
).Or(
  db.Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge"),
).Find(&Pizza{}).Statement

// SELECT * FROM `pizzas` WHERE (pizza = "pepperoni" AND (size = "small" OR size = "medium")) OR (pizza = "hawaiian" AND size = "xlarge")
```

## IN 与多个列 IN with multiple columns

Selecting IN with multiple columns

​	选择 IN 与多个列。

``` go
db.Where("(name, age, role) IN ?", [][]interface{}{{"jinzhu", 18, "admin"}, {"jinzhu2", 19, "user"}}).Find(&users)
// SELECT * FROM users WHERE (name, age, role) IN (("jinzhu", 18, "admin"), ("jinzhu 2", 19, "user"));
```

## 命名实参 Named Argument

GORM supports named arguments with [`sql.NamedArg`](https://tip.golang.org/pkg/database/sql/#NamedArg) or `map[string]interface{}{}`, for example:

​	GORM 支持命名参数，可以使用 [sql.NamedArg]({{< ref "/stdLib/database/sql#type-namedarg----go18">}}) 或 `map[string]interface{}{}`，例如：

``` go
db.Where("name1 = @name OR name2 = @name", sql.Named("name", "jinzhu")).Find(&user)
// SELECT * FROM `users` WHERE name1 = "jinzhu" OR name2 = "jinzhu"

db.Where("name1 = @name OR name2 = @name", map[string]interface{}{"name": "jinzhu"}).First(&user)
// SELECT * FROM `users` WHERE name1 = "jinzhu" OR name2 = "jinzhu" ORDER BY `users`.`id` LIMIT 1
```

Check out [Raw SQL and SQL Builder](https://gorm.io/docs/sql_builder.html#named_argument) for more detail

​	更多详细信息请查看 [Raw SQL and SQL Builder](../rawSQLAndSQLBuilder)。

## Find To Map

GORM allows scanning results to `map[string]interface{}` or `[]map[string]interface{}`, don’t forget to specify `Model` or `Table`, for example:

​	GORM 允许扫描结果到 `map[string]interface{}` 或 `[]map[string]interface{}`，不要忘记指定 `Model` 或 `Table`，例如：

``` go
result := map[string]interface{}{}
db.Model(&User{}).First(&result, "id = ?", 1)

var results []map[string]interface{}
db.Table("users").Find(&results)
```

## FirstOrInit

Get first matched record or initialize a new instance with given conditions (only works with struct or map conditions)

​	获取第一个匹配的记录或初始化一个具有给定条件的实例（仅适用于结构或映射条件）。

``` go
// 用户未找到，请使用给定条件初始化。 User not found, initialize it with give conditions
db.FirstOrInit(&user, User{Name: "non_existing"})
// user -> User{Name: "non_existing"}

// 找到了名为jinzhu的用户。 Found user with `name` = `jinzhu`
db.Where(User{Name: "jinzhu"}).FirstOrInit(&user)
// user -> User{ID: 111, Name: "Jinzhu", Age: 18}

// 找到了名为jinzhu的用户。 Found user with `name` = `jinzhu`
db.FirstOrInit(&user, map[string]interface{}{"name": "jinzhu"})
// user -> User{ID: 111, Name: "Jinzhu", Age: 18}
```

Initialize struct with more attributes if record not found, those `Attrs` won’t be used to build the SQL query

​	如果记录未找到，使用更多属性初始化结构体，这些`Attrs`将不会被用于构建SQL查询。

``` go
// 用户未找到，根据给定条件和Attrs初始化它 User not found, initialize it with give conditions and Attrs
db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrInit(&user)
// SELECT * FROM USERS WHERE name = 'non_existing' ORDER BY id LIMIT 1;
// user -> User{Name: "non_existing", Age: 20}

// 用户未找到，根据给定条件和Attrs初始化它 User not found, initialize it with give conditions and Attrs
db.Where(User{Name: "non_existing"}).Attrs("age", 20).FirstOrInit(&user)
// SELECT * FROM USERS WHERE name = 'non_existing' ORDER BY id LIMIT 1;
// user -> User{Name: "non_existing", Age: 20}

// 找到名为`jinzhu`的用户，属性将被忽略 Found user with `name` = `jinzhu`, attributes will be ignored
db.Where(User{Name: "Jinzhu"}).Attrs(User{Age: 20}).FirstOrInit(&user)
// SELECT * FROM USERS WHERE name = jinzhu' ORDER BY id LIMIT 1;
// user -> User{ID: 111, Name: "Jinzhu", Age: 18}
```

`Assign` attributes to struct regardless it is found or not, those attributes won’t be used to build SQL query and the final data won’t be saved into database

​	无论是否找到用户，都将属性`Assign`给结构体，这些属性不会被用于构建SQL查询，并且最终的数据不会保存到数据库中。

``` go
// 用户未找到，根据给定条件和Assign属性初始化它 User not found, initialize it with give conditions and Assign attributes
db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrInit(&user)
// user -> User{Name: "non_existing", Age: 20}

// 找到名为`jinzhu`的用户，使用Assign属性更新它 Found user with `name` = `jinzhu`, update it with Assign attributes
db.Where(User{Name: "Jinzhu"}).Assign(User{Age: 20}).FirstOrInit(&user)
// SELECT * FROM USERS WHERE name = jinzhu' ORDER BY id LIMIT 1;
// user -> User{ID: 111, Name: "Jinzhu", Age: 20}
```

## FirstOrCreate

Get first matched record or create a new one with given conditions (only works with struct, map conditions), `RowsAffected` returns created/updated record’s count

​	获取第一个匹配的记录或根据给定条件创建一个新的记录（仅适用于结构体、映射条件），`RowsAffected` 返回创建/更新记录的数量

``` go
// User not found, create a new record with give conditions
result := db.FirstOrCreate(&user, User{Name: "non_existing"})
// INSERT INTO "users" (name) VALUES ("non_existing");
// user -> User{ID: 112, Name: "non_existing"}
// result.RowsAffected // => 1

// Found user with `name` = `jinzhu`
result := db.Where(User{Name: "jinzhu"}).FirstOrCreate(&user)
// user -> User{ID: 111, Name: "jinzhu", "Age": 18}
// result.RowsAffected // => 0
```

Create struct with more attributes if record not found, those `Attrs` won’t be used to build SQL query

​	如果记录未找到，使用更多属性创建结构体，这些 `Attrs` 不会用于构建 SQL 查询

``` go
// 用户未找到，根据给定条件和 Attrs 创建一个新记录 User not found, create it with give conditions and Attrs
db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrCreate(&user)
// SELECT * FROM users WHERE name = 'non_existing' ORDER BY id LIMIT 1;
// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
// user -> User{ID: 112, Name: "non_existing", Age: 20}

// 找到了名为 `jinzhu` 的用户，属性将被忽略 Found user with `name` = `jinzhu`, attributes will be ignored
db.Where(User{Name: "jinzhu"}).Attrs(User{Age: 20}).FirstOrCreate(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
// user -> User{ID: 111, Name: "jinzhu", Age: 18}
```

`Assign` attributes to the record regardless it is found or not and save them back to the database.

​	无论是否找到记录，都将属性`Assign`给记录并将其保存回数据库。

``` go
// 用户未找到，根据给定条件和 Assign 属性初始化一个新记录 User not found, initialize it with give conditions and Assign attributes
db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrCreate(&user)
// SELECT * FROM users WHERE name = 'non_existing' ORDER BY id LIMIT 1;
// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
// user -> User{ID: 112, Name: "non_existing", Age: 20}

// 找到了名为 `jinzhu` 的用户，使用 Assign 属性更新它 Found user with `name` = `jinzhu`, update it with Assign attributes
db.Where(User{Name: "jinzhu"}).Assign(User{Age: 20}).FirstOrCreate(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
// UPDATE users SET age=20 WHERE id = 111;
// user -> User{ID: 111, Name: "jinzhu", Age: 20}
```

## 优化器/索引提示 Optimizer/Index Hints

Optimizer hints allow to control the query optimizer to choose a certain query execution plan, GORM supports it with `gorm.io/hints`, e.g:

​	优化器提示允许控制查询优化器选择特定的查询执行计划，GORM 支持它，例如：

``` go
import "gorm.io/hints"

db.Clauses(hints.New("MAX_EXECUTION_TIME(10000)")).Find(&User{})
// SELECT * /*+ MAX_EXECUTION_TIME(10000) */ FROM `users`
```

Index hints allow passing index hints to the database in case the query planner gets confused.

​	索引提示允许将索引提示传递给数据库以防查询规划器混淆。

``` go
import "gorm.io/hints"

db.Clauses(hints.UseIndex("idx_user_name")).Find(&User{})
// SELECT * FROM `users` USE INDEX (`idx_user_name`)

db.Clauses(hints.ForceIndex("idx_user_name", "idx_user_id").ForJoin()).Find(&User{})
// SELECT * FROM `users` FORCE INDEX FOR JOIN (`idx_user_name`,`idx_user_id`)"
```

Refer [Optimizer Hints/Index/Comment](https://gorm.io/docs/hints.html) for more details

​	参考 [优化器提示/索引/注释]({{< ref "/gorm/AdancedTopics/hints">}}) 以获取更多详细信息

## 迭代 Iteration

GORM supports iterating through Rows

​	GORM支持遍历行

``` go
rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Rows()
defer rows.Close()

for rows.Next() {
  var user User
  // `gorm.DB`的`ScanRows`方法可以将一行扫描到结构体中 ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
  db.ScanRows(rows, &user)

  // do something
}
```

## FindInBatches

Query and process records in batch

​	分批查询和处理记录

``` go
// batch size 100
result := db.Where("processed = ?", false).FindInBatches(&results, 100, func(tx *gorm.DB, batch int) error {
  for _, result := range results {
    // batch processing found records
  }

  tx.Save(&results)

  tx.RowsAffected // number of records in this batch

  batch // Batch 1, 2, 3

  // 如果返回错误，将停止后续批次 returns error will stop future batches
  return nil
})

result.Error // 返回的错误 returned error
result.RowsAffected // 所有批次中处理过的记录数 processed records count in all batches
```

## 查询钩子 Query Hooks

GORM allows hooks `AfterFind` for a query, it will be called when querying a record, refer [Hooks](https://gorm.io/docs/hooks.html) for details

​	GORM允许在查询记录时调用`AfterFind`钩子，详情请参考[Hooks]({{< ref "/gorm/Tutorials/hooks">}})。

``` go
func (u *User) AfterFind(tx *gorm.DB) (err error) {
  if u.Role == "" {
    u.Role = "user"
  }
  return
}
```

## Pluck

Query single column from database and scan into a slice, if you want to query multiple columns, use `Select` with [`Scan`](https://gorm.io/docs/query.html#scan) instead

​	从数据库中查询单列并将其扫描到切片中，如果要查询多个列，请使用带有[Scan]({{< ref "/gorm/CRUDInterface/query#scan">}})的`Select`代替。

``` go
var ages []int64
db.Model(&users).Pluck("age", &ages)

var names []string
db.Model(&User{}).Pluck("name", &names)

db.Table("deleted_users").Pluck("name", &names)

// Distinct Pluck
db.Model(&User{}).Distinct().Pluck("Name", &names)
// SELECT DISTINCT `name` FROM `users`

// 请求多个列，使用`Scan`或`Find`像这样： Requesting more than one column, use `Scan` or `Find` like this:
db.Select("name", "age").Scan(&users)
db.Select("name", "age").Find(&users)
```

## 范围（Scopes）Scopes

`Scopes` allows you to specify commonly-used queries which can be referenced as method calls

​	`Scopes`允许你指定常用的查询，可以像方法调用一样引用它们。

``` go
func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
  return db.Where("amount > ?", 1000)
}

func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
  return db.Where("pay_mode_sign = ?", "C")
}

func PaidWithCod(db *gorm.DB) *gorm.DB {
  return db.Where("pay_mode_sign = ?", "C")
}

func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    return db.Where("status IN (?)", status)
  }
}

db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&orders)
// 查找所有信用卡订单且金额大于1000的订单 Find all credit card orders and amount greater than 1000

db.Scopes(AmountGreaterThan1000, PaidWithCod).Find(&orders)
// 查找所有COD订单且金额大于1000的订单 Find all COD orders and amount greater than 1000

db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
// 查找所有已支付、已发货且金额大于1000的订单 Find all paid, shipped orders that amount greater than 1000
```

Checkout [Scopes](https://gorm.io/docs/scopes.html) for details

​	请参阅[范围（Scopes）]({{< ref "/gorm/Tutorials/scopes">}})以获取详细信息

## Count

Get matched records count

​	获取匹配的记录数

``` go
var count int64
db.Model(&User{}).Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Count(&count)
// SELECT count(1) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'

db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
// SELECT count(1) FROM users WHERE name = 'jinzhu'; (count)

db.Table("deleted_users").Count(&count)
// SELECT count(1) FROM deleted_users;

// Count with Distinct
db.Model(&User{}).Distinct("name").Count(&count)
// SELECT COUNT(DISTINCT(`name`)) FROM `users`

db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
// SELECT count(distinct(name)) FROM deleted_users

// Count with Group
users := []User{
  {Name: "name1"},
  {Name: "name2"},
  {Name: "name3"},
  {Name: "name3"},
}

db.Model(&User{}).Group("name").Count(&count)
count // => 3
```