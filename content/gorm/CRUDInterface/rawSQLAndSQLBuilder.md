+++
title = "SQL Builder"
date = 2023-10-28T14:26:51+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/sql_builder.html](https://gorm.io/docs/sql_builder.html)

## 原始SQL Raw SQL

Query Raw SQL with `Scan`

​	使用`Scan`查询原始SQL

``` go
type Result struct {
  ID   int
  Name string
  Age  int
}

var result Result
db.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&result)

db.Raw("SELECT id, name, age FROM users WHERE name = ?", "jinzhu").Scan(&result)

var age int
db.Raw("SELECT SUM(age) FROM users WHERE role = ?", "admin").Scan(&age)

var users []User
db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "jinzhu", 20).Scan(&users)
```

`Exec` with Raw SQL

​	使用原始SQL执行`Exec`

``` go
db.Exec("DROP TABLE users")
db.Exec("UPDATE orders SET shipped_at = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})

// Exec with SQL Expression
db.Exec("UPDATE users SET money = ? WHERE name = ?", gorm.Expr("money * ? + ?", 10000, 1), "jinzhu")
```

> **NOTE** GORM allows cache prepared statement to increase performance, checkout [Performance](https://gorm.io/docs/performance.html) for details
>
> **注意** GORM允许缓存预编译语句以提高性能，请参阅[性能]({{< ref "/gorm/Tutorials/performance">}})以获取详细信息

## 命名实参 Named Argument

GORM supports named arguments with [`sql.NamedArg`](https://tip.golang.org/pkg/database/sql/#NamedArg), `map[string]interface{}{}` or struct, for example:

​	GORM支持使用[sql.NamedArg]({{< ref "/stdLib/database/sql#type-namedarg----go18">}})、`map[string]interface{}{}`或结构体作为命名参数，例如：

``` go
db.Where("name1 = @name OR name2 = @name", sql.Named("name", "jinzhu")).Find(&user)
// SELECT * FROM `users` WHERE name1 = "jinzhu" OR name2 = "jinzhu"

db.Where("name1 = @name OR name2 = @name", map[string]interface{}{"name": "jinzhu2"}).First(&result3)
// SELECT * FROM `users` WHERE name1 = "jinzhu2" OR name2 = "jinzhu2" ORDER BY `users`.`id` LIMIT 1

// 使用原始SQL的命名实参 Named Argument with Raw SQL
db.Raw("SELECT * FROM users WHERE name1 = @name OR name2 = @name2 OR name3 = @name",
   sql.Named("name", "jinzhu1"), sql.Named("name2", "jinzhu2")).Find(&user)
// SELECT * FROM users WHERE name1 = "jinzhu1" OR name2 = "jinzhu2" OR name3 = "jinzhu1"

db.Exec("UPDATE users SET name1 = @name, name2 = @name2, name3 = @name",
   sql.Named("name", "jinzhunew"), sql.Named("name2", "jinzhunew2"))
// UPDATE users SET name1 = "jinzhunew", name2 = "jinzhunew2", name3 = "jinzhunew"

db.Raw("SELECT * FROM users WHERE (name1 = @name AND name3 = @name) AND name2 = @name2",
   map[string]interface{}{"name": "jinzhu", "name2": "jinzhu2"}).Find(&user)
// SELECT * FROM users WHERE (name1 = "jinzhu" AND name3 = "jinzhu") AND name2 = "jinzhu2"

type NamedArgument struct {
  Name string
  Name2 string
}

db.Raw("SELECT * FROM users WHERE (name1 = @Name AND name3 = @Name) AND name2 = @Name2",
   NamedArgument{Name: "jinzhu", Name2: "jinzhu2"}).Find(&user)
// SELECT * FROM users WHERE (name1 = "jinzhu" AND name3 = "jinzhu") AND name2 = "jinzhu2"
```

## DryRun模式 DryRun Mode

Generate `SQL` and its arguments without executing, can be used to prepare or test generated SQL, Checkout [Session](https://gorm.io/docs/session.html) for details

​	生成`SQL`及其参数而不执行，可用于准备或测试生成的SQL，请参阅[Session]({{< ref "/gorm/Tutorials/session">}})以获取详细信息

``` go
stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
stmt.Vars         //=> []interface{}{1}
```

## ToSQL

Returns generated `SQL` without executing.

​	返回生成的`SQL`而不执行。

GORM uses the database/sql’s argument placeholders to construct the SQL statement, which will automatically escape arguments to avoid SQL injection, but the generated SQL don’t provide the safety guarantees, please only use it for debugging.

​	GORM使用database/sql的参数占位符来构建SQL语句，这将自动转义参数以避免SQL注入，但生成的SQL不提供安全保证，请仅用于调试。

``` go
sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
  return tx.Model(&User{}).Where("id = ?", 100).Limit(10).Order("age desc").Find(&[]User{})
})
sql //=> SELECT * FROM "users" WHERE id = 100 AND "users"."deleted_at" IS NULL ORDER BY age desc LIMIT 10
```

## `Row` & `Rows`

Get result as `*sql.Row`

​	获取结果为`*sql.Row`

``` go
// 使用GORM API构建SQL Use GORM API build SQL
row := db.Table("users").Where("name = ?", "jinzhu").Select("name", "age").Row()
row.Scan(&name, &age)

// 使用Raw SQL Use Raw SQL
row := db.Raw("select name, age, email from users where name = ?", "jinzhu").Row()
row.Scan(&name, &age, &email)
```

Get result as `*sql.Rows`

​	获取结果为`*sql.Rows`

``` go
// 使用GORM API构建SQL Use GORM API build SQL
rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows()
defer rows.Close()
for rows.Next() {
  rows.Scan(&name, &age, &email)

  // do something
}

// Raw SQL
rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()
defer rows.Close()
for rows.Next() {
  rows.Scan(&name, &age, &email)

  // do something
}
```

Checkout [FindInBatches](https://gorm.io/docs/advanced_query.html) for how to query and process records in batch

​	查看[FindInBatches]({{< ref "/gorm/CRUDInterface/advancedQuery#findinbatches">}})以如何查询和处理记录批次。

Checkout [Group Conditions](https://gorm.io/docs/advanced_query.html#group_conditions) for how to build complicated SQL Query

​	查看[Group Conditions]({{< ref "/gorm/CRUDInterface/advancedQuery#group条件-group-conditions">}})以如何构建复杂的SQL查询。

## 将`*sql.Rows`扫描到结构体中 Scan `*sql.Rows` into struct

Use `ScanRows` to scan a row into a struct, for example:

​	使用`ScanRows`将一行扫描到结构体中，例如：

``` go
rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
defer rows.Close()

var user User
for rows.Next() {
  // ScanRows scan a row into user
  db.ScanRows(rows, &user)

  // do something
}
```

## 连接 Connection

Run mutliple SQL in same db tcp connection (not in a transaction)

​	在同一数据库的TCP连接中运行多个SQL（不是事务）

``` go
db.Connection(func(tx *gorm.DB) error {
  tx.Exec("SET my.role = ?", "admin")

  tx.First(&User{})
})
```

## 高级 Advanced

### 子句 Clauses

GORM uses SQL builder generates SQL internally, for each operation, GORM creates a `*gorm.Statement` object, all GORM APIs add/change `Clause` for the `Statement`, at last, GORM generated SQL based on those clauses

​	GORM使用SQL构建器生成SQL内部，对于每个操作，GORM创建一个`*gorm.Statement`对象，所有GORM API为`Statement`添加/更改`Clause`，最后，GORM根据这些子句生成SQL。

For example, when querying with `First`, it adds the following clauses to the `Statement`	

​	例如，当使用`First`查询时，它会向`Statement`添加以下子句。

``` go
var limit = 1
clause.Select{Columns: []clause.Column{{Name: "*"}}}
clause.From{Tables: []clause.Table{{Name: clause.CurrentTable}}}
clause.Limit{Limit: &limit}
clause.OrderBy{Columns: []clause.OrderByColumn{
  {
    Column: clause.Column{
      Table: clause.CurrentTable,
      Name:  clause.PrimaryKey,
    },
  },
}}
```

Then GORM build finally querying SQL in the `Query` callbacks like:

​	然后GORM在`Query`回调中构建最终的查询SQL，如下所示：

``` go
Statement.Build("SELECT", "FROM", "WHERE", "GROUP BY", "ORDER BY", "LIMIT", "FOR")
```

Which generate SQL:

​	生成的SQL如下：

``` go
SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
```

You can define your own `Clause` and use it with GORM, it needs to implements [Interface](https://pkg.go.dev/gorm.io/gorm/clause?tab=doc#Interface)

​	你可以定义自己的`Clause`并使用它与GORM一起使用，需要实现[Interface](https://pkg.go.dev/gorm.io/gorm/clause?tab=doc#Interface)

Check out [examples](https://github.com/go-gorm/gorm/tree/master/clause) for reference

​	查看[示例](https://github.com/go-gorm/gorm/tree/master/clause)以获取参考

### 子句构建器 Clause Builder

For different databases, Clauses may generate different SQL, for example:

​	对于不同的数据库，子句可能生成不同的SQL，例如：

``` go
db.Offset(10).Limit(5).Find(&users)
// 为SQL Server生成的 Generated for SQL Server
// SELECT * FROM "users" OFFSET 10 ROW FETCH NEXT 5 ROWS ONLY
// 为MySQL生成的 Generated for MySQL
// SELECT * FROM `users` LIMIT 5 OFFSET 10
```

Which is supported because GORM allows database driver register Clause Builder to replace the default one, take the [Limit](https://github.com/go-gorm/sqlserver/blob/512546241200023819d2e7f8f2f91d7fb3a52e42/sqlserver.go#L45) as example

​	这是由于GORM允许数据库驱动程序注册子句构建器以替换默认的构建器，请参阅[Limit](https://github.com/go-gorm/sqlserver/blob/512546241200023819d2e7f8f2f91d7fb3a52e42/sqlserver.go#L45)作为示例

### 子句选项 Clause Options

GORM defined [Many Clauses](https://github.com/go-gorm/gorm/tree/master/clause), and some clauses provide advanced options can be used for your application

​	GORM定义了[许多子句](https://github.com/go-gorm/gorm/tree/master/clause)，其中一些子句提供了高级选项，可以满足您的应用程序需求

Although most of them are rarely used, if you find GORM public API can’t match your requirements, may be good to check them out, for example:

​	尽管它们中的大多数很少使用，但如果GORM公共API不能满足您的需求，也许值得检查一下，例如：

``` go
db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&user)
// INSERT IGNORE INTO users (name,age...) VALUES ("jinzhu",18...);
```

### StatementModifier

GORM provides interface [StatementModifier](https://pkg.go.dev/gorm.io/gorm?tab=doc#StatementModifier) allows you modify statement to match your requirements, take [Hints](https://gorm.io/docs/hints.html) as example

​	GORM提供了接口[StatementModifier](https://pkg.go.dev/gorm.io/gorm?tab=doc#StatementModifier)允许您修改语句以满足您的需求，请参阅[提示](https://gorm.io/docs/hints.html)作为示例

``` go
import "gorm.io/hints"

db.Clauses(hints.New("hint")).Find(&User{})
// SELECT * /*+ hint */ FROM `users`
```