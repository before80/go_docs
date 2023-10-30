+++
title = "方法链"
date = 2023-10-28T14:30:07+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/method_chaining.html](https://gorm.io/docs/method_chaining.html)

GORM allows method chaining, so you can write code like this:

​	GORM允许方法链，因此你可以这样写代码：

``` go
db.Where("name = ?", "jinzhu").Where("age = ?", 18).First(&user)
```

There are three kinds of methods in GORM: `Chain Method`, `Finisher Method`, `New Session Method`.

​	GORM有三种方法：`Chain Method`，`Finisher Method`和`New Session Method`。

After a `Chain method`, `Finisher Method`, GORM returns an initialized `*gorm.DB` instance, which is NOT safe to reuse anymore, or new generated SQL might be polluted by the previous conditions, for example:

​	在`Chain method`之后，`Finisher Method`，GORM返回一个初始化的`*gorm.DB`实例，这个实例不能再被重用，或者新的SQL可能会受到之前条件的污染，例如：

``` go
queryDB := DB.Where("name = ?", "jinzhu")

queryDB.Where("age > ?", 10).First(&user)
// SELECT * FROM users WHERE name = "jinzhu" AND age > 10

queryDB.Where("age > ?", 20).First(&user2)
// SELECT * FROM users WHERE name = "jinzhu" AND age > 10 AND age > 20
```

In order to reuse a initialized `*gorm.DB` instance, you can use a `New Session Method` to create a shareable `*gorm.DB`, e.g:

​	为了重用一个初始化的`*gorm.DB`实例，你可以使用`New Session Method`创建一个可共享的`*gorm.DB`，例如：

``` go
queryDB := DB.Where("name = ?", "jinzhu").Session(&gorm.Session{})

queryDB.Where("age > ?", 10).First(&user)
// SELECT * FROM users WHERE name = "jinzhu" AND age > 10

queryDB.Where("age > ?", 20).First(&user2)
// SELECT * FROM users WHERE name = "jinzhu" AND age > 20
```

## Chain方法 Chain Method

Chain methods are methods to modify or add `Clauses` to current `Statement`, like:

​	Chain方法是用于修改或添加`Clauses`到当前`Statement`的方法，例如：

`Where`, `Select`, `Omit`, `Joins`, `Scopes`, `Preload`, `Raw` (`Raw` can’t be used with other chainable methods to build SQL)…

​	`Where`，`Select`，`Omit`，`Joins`，`Scopes`，`Preload`，`Raw`（`Raw`不能与其他可链接的方法一起构建SQL）…

Here is [the full lists](https://github.com/go-gorm/gorm/blob/master/chainable_api.go), also check out the [SQL Builder](https://gorm.io/docs/sql_builder.html) for more details about `Clauses`.

​	这里有一个[完整的列表](https://github.com/go-gorm/gorm/blob/master/chainable_api.go)，也可以查看[SQL Builder]({{< ref "/gorm/CRUDInterface/rawSQLAndSQLBuilder">}})以获取更多关于`Clauses`的详细信息。

## Finisher方法 Finisher Method

Finishers are immediate methods that execute registered callbacks, which will generate and execute SQL, like those methods:

​	Finishers是立即执行已注册回调的方法，它们将生成并执行SQL，例如这些方法：

`Create`, `First`, `Find`, `Take`, `Save`, `Update`, `Delete`, `Scan`, `Row`, `Rows`…

​	`Create`，`First`，`Find`，`Take`，`Save`，`Update`，`Delete`，`Scan`，`Row`，`Rows`…

Check out [the full lists](https://github.com/go-gorm/gorm/blob/master/finisher_api.go) here.

​	请查看[完整列表](https://github.com/go-gorm/gorm/blob/master/finisher_api.go)。

## New Session Method

GORM defined `Session`, `WithContext`, `Debug` methods as `New Session Method`, refer [Session](https://gorm.io/docs/session.html) for more details.

​	GORM定义了`Session`，`WithContext`，`Debug`方法作为`New Session Method`，参考[Session](https://gorm.io/docs/session.html)以获取更多详细信息。

After a `Chain method`, `Finisher Method`, GORM returns an initialized `*gorm.DB` instance, which is NOT safe to reuse anymore, you should use a `New Session Method` to mark the `*gorm.DB` as shareable.

​	在`Chain method`，`Finisher Method`之后，GORM返回一个初始化的`*gorm.DB`实例，这个实例不能再被重用，你应该使用`New Session Method`将其标记为可共享。

Let’s explain it with examples:

​	让我们通过示例来解释：

示例1 ：Example 1:

``` go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// db是一个新初始化的`*gorm.DB`实例，它是安全的可以重用的 db is a new initialized `*gorm.DB`, which is safe to reuse

db.Where("name = ?", "jinzhu").Where("age = ?", 18).Find(&users)
// Where("name = ?", "jinzhu")`是第一个链方法调用，它创建了一个初始化的`*gorm.Statement`实例，也就是`*gorm.DB` `Where("name = ?", "jinzhu")` is the first chain method call, it will create an initialized `*gorm.DB` instance, aka `*gorm.Statement`
// `Where("age = ?", 18)`是第二个链方法调用，它重用了上面的`*gorm.Statement`，向其添加了新的条件`age = 18` `Where("age = ?", 18)` is the second chain method call, it reuses the above `*gorm.Statement`, adds new condition `age = 18` to it
// `Find(&users)`是一个完成器方法，它执行已注册的查询回调，生成并运行以下SQL： `Find(&users)` is a finisher method, it executes registered Query Callbacks, which generates and runs the following SQL:
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18;

db.Where("name = ?", "jinzhu2").Where("age = ?", 20).Find(&users)
// `Where("name = ?", "jinzhu2")`也是第一个链方法调用，它创建了一个新的`*gorm.Statement` `Where("name = ?", "jinzhu2")` is also the first chain method call, it creates a new `*gorm.Statement`
// `Where("age = ?", 20)`重用了上面的`Statement`，并向其添加了条件 `Where("age = ?", 20)` reuses the above `Statement`, and add conditions to it
// `Find(&users)`是一个完成器方法，它执行已注册的查询回调，生成并运行以下SQL： `Find(&users)` is a finisher method, it executes registered Query Callbacks, generates and runs the following SQL:
// SELECT * FROM users WHERE name = 'jinzhu2' AND age = 20;

db.Find(&users)
// `Find(&users)`是一个完成器方法调用，它也创建了一个新的`Statement`并执行了已注册的查询回调，生成并运行以下SQL： `Find(&users)` is a finisher method call, it also creates a new `Statement` and executes registered Query Callbacks, generates and runs the following SQL:
// SELECT * FROM users;
```

（不好的）示例2： (Bad) Example 2:

``` go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// db是一个新初始化的*gorm.DB实例，它是安全的可以重用的 db is a new initialized *gorm.DB, which is safe to reuse

tx := db.Where("name = ?", "jinzhu")
// `Where("name = ?", "jinzhu")`返回一个初始化的`*gorm.Statement`实例后在链方法`Where`之后，这是不安全的，不能重用 `Where("name = ?", "jinzhu")` returns an initialized `*gorm.Statement` instance after chain method `Where`, which is NOT safe to reuse

// good case
tx.Where("age = ?", 18).Find(&users)
// `tx.Where("age = ?", 18)`使用上述的`*gorm.Statement`，向其添加了新条件 `tx.Where("age = ?", 18)` use the above `*gorm.Statement`, adds new condition to it
// `Find(&users)`是一个完成器方法调用，它执行已注册的查询回调，生成并运行以下SQL： `Find(&users)` is a finisher method call, it executes registered Query Callbacks, generates and runs the following SQL:
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18

// bad case
tx.Where("age = ?", 28).Find(&users)
// `tx.Where("age = ?", 28)`也使用上述的`*gorm.Statement`，并继续添加条件 `tx.Where("age = ?", 28)` also use the above `*gorm.Statement`, and keep adding conditions to it
// 所以下面的生成的SQL被之前的条件污染了： So the following generated SQL is polluted by the previous conditions:
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18 AND age = 28;
```

示例3：Example 3:

``` go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// db是一个新初始化的*gorm.DB实例，它是安全的可以重用的 db is a new initialized *gorm.DB, which is safe to reuse

tx := db.Where("name = ?", "jinzhu").Session(&gorm.Session{})
tx := db.Where("name = ?", "jinzhu").WithContext(context.Background())
tx := db.Where("name = ?", "jinzhu").Debug()
// `Session`, `WithContext`, `Debug`返回`*gorm.DB`标记为安全可以重用，基于它的新初始化的`*gorm.Statement`保持当前条件 `Session`, `WithContext`, `Debug` returns `*gorm.DB` marked as safe to reuse, newly initialized `*gorm.Statement` based on it keeps current conditions

// good case
tx.Where("age = ?", 18).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18

// good case
tx.Where("age = ?", 28).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 28;
```