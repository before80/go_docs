+++
title = "ORM 用法"
date = 2024-02-04T10:00:11+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/model/orm/]({{< ref "/beego/mvcIntroduction/models/ormUsage" >}})

# ORM Usage - ORM 用法



## ORM Usage ORM 用法

An example of beego/orm is set out below.

​	下面列出了 beego/orm 的一个示例。

All the code samples in this section are based on this example unless otherwise stated.

​	除非另有说明，本节中的所有代码示例均基于此示例。

In v2.x, there is a big big change:

​	在 v2.x 中，有一个很大的变化：

The ORM instance should be stateless, so it’s now thread safe.

​	ORM 实例应该是无状态的，因此现在是线程安全的。

##### models.go: models.go：

```go
package main

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id          int
	Name        string
	Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // Reverse relationship (optional)
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User), new(Profile))
}
```

##### main.go

```go
package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}

func main() {

// Using default, you can use other database
	o := orm.NewOrm()

	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}
```

## Set up database 设置数据库

ORM supports three popular databases. Here are the tested drivers, you need to import them:

​	ORM 支持三个流行的数据库。以下是经过测试的驱动程序，您需要导入它们：

```go
import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)
```

#### RegisterDriver

Default databases:

​	默认数据库：

```go
orm.DRMySQL
orm.DRSqlite
orm.DRPostgres
orm.DRTiDB

// < 1.6
orm.DR_MySQL
orm.DR_Sqlite
orm.DR_Postgres
// param 1: driverName
// param 2: database type
// This mapping driverName and database type
// mysql / sqlite3 / postgres / TiDB registered by default already
orm.RegisterDriver("mysql", orm.DRMySQL)
```

#### RegisterDataBase

ORM must register a database with alias `default`.

​	ORM 必须使用别名 `default` 注册数据库。

ORM uses golang built-in connection pool.

​	ORM 使用 golang 内置连接池。

```go
// param 1:        Database alias. ORM will use it to switch database.
// param 2:        driverName
// param 3:        connection string
orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")

// param 4 (optional):  set maximum idle connections
// param 4 (optional):  set maximum connections (go >= 1.2)
maxIdle := 30
maxConn := 30
orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8", maxIdle, maxConn)
```

See [Test]({{< ref "/beego/mvcIntroduction/models/testCases" >}}) for more information on database connection strings.

​	有关数据库连接字符串的更多信息，请参阅测试。

#### SetMaxIdleConns

Set maximum idle connections according to database alias:

​	根据数据库别名设置最大空闲连接：

```go
orm.SetMaxIdleConns("default", 30)
```

#### SetMaxOpenConns

Set maximum connections (go >= 1.2) according to database alias:

​	设定数据库别名对应的连接数上限（go >= 1.2）：

```go
orm.SetMaxOpenConns("default", 30)
```

#### Timezone Config 时区配置

ORM uses time.Local by default

​	ORM 默认使用 time.Local

- used for ORM automatically created time
  用于 ORM 自动创建的时间
- convert time queried from database into ORM local time
  将从数据库查询的时间转换为 ORM 本地时间

You can change it if needed:

​	如果需要，您可以更改它：

```go
// Set to UTC time
orm.DefaultTimeLoc = time.UTC
```

ORM will get timezone of database while performing `RegisterDataBase`. When setting or getting time.Time it will convert accordingly to match system time and make sure the time is correct.

​	ORM 将在执行 `RegisterDataBase` 时获取数据库的时区。在设置或获取 time.Time 时，它将相应地进行转换以匹配系统时间并确保时间正确。

**Note:
注意：**

- In Sqlite3, set and get use UTC time by default.
  在 Sqlite3 中，默认情况下设置和获取使用 UTC 时间。
- When using `go-sql-driver` driver，please pay attention to your DSN config. From a version of `go-sql-driver` the default uses utc timezone not local. So if you use another timezone, please set it. eg: `root:root@/orm_test?charset=utf8&loc=Asia%2FShanghai` ref: [loc](https://github.com/go-sql-driver/mysql#loc) / [parseTime](https://github.com/go-sql-driver/mysql#parsetime)
  使用 `go-sql-driver` 驱动程序时，请注意您的 DSN 配置。从 `go-sql-driver` 版本开始，默认使用 utc 时区而不是本地时区。因此，如果您使用其他时区，请设置它。例如： `root:root@/orm_test?charset=utf8&loc=Asia%2FShanghai` 参考资料 / parseTime

## Registering Model 注册模型

Registering a model is mandatory if you use orm.QuerySeter for advanced queries.

​	如果您使用 orm.QuerySeter 进行高级查询，则必须注册模型。

Otherwise, you don’t need to do this if you’re using raw SQL queries and map struct only. [See Raw SQL Query]({{< ref "/beego/mvcIntroduction/models/rawSQLToQuery" >}})

​	否则，如果您使用原始 SQL 查询并仅映射结构，则无需执行此操作。请参阅原始 SQL 查询

#### RegisterModel

Register the Model you defined. The best practice is to have a single models.go file and register in it’s init function.

​	注册您定义的模型。最佳做法是在单个 models.go 文件中进行注册，并在其 init 函数中注册。

Mini models.go

```go
package main

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id   int
	name string
}

func init(){
	orm.RegisterModel(new(User))
}
```

RegisterModel can register multiple models at the same time:

​	RegisterModel 可以同时注册多个模型：

```go
orm.RegisterModel(new(User), new(Profile), new(Post))
```

For detailed struct definition, see [Model define]({{< ref "/beego/mvcIntroduction/models/modelDefinition" >}})

​	有关详细的结构定义，请参阅模型定义

#### Generate Tables 生成表

You may want Beego to automatically create your database tables. One way to do this is by using the method described in the [cli]({{< ref "/beego/mvcIntroduction/models/commandLine" >}}) documentation. Alternatively, you could choose to autogenerate your tables by including the following in your main.go file in your main block.

​	您可能希望 Beego 自动创建数据库表。执行此操作的一种方法是使用 cli 文档中描述的方法。或者，您也可以选择通过在 main 块中的 main.go 文件中包含以下内容来自动生成表。

```go
// Database alias.
name := "default"

// Drop table and re-create.
force := true

// Print log.
verbose := true

// Error.
err := orm.RunSyncdb(name, force, verbose)
if err != nil {
	fmt.Println(err)
}
```

After the initial “bee run” command, change the values of force and verbose to false. The default behavior for Beego is to add additional columns when the model is updated. You will need to manually handle dropping your columns if they are removed from your model.

​	在初始“bee run”命令后，将 force 和 verbose 的值更改为 false。Beego 的默认行为是在更新模型时添加其他列。如果从模型中删除列，您需要手动处理删除列。

#### RegisterModelWithPrefix

Using table prefix

​	使用表前缀

```go
orm.RegisterModelWithPrefix("prefix_", new(User))
```

The created table name is prefix_user

​	创建的表名为 prefix_user

#### NewOrmWithDB

You may need to manage db pools by yourself. (eg: needing two queries in one connection)

​	您可能需要自己管理数据库池。（例如：在一个连接中需要两个查询）

But you want to use awesome orm features. Voila!

​	但您想使用强大的 orm 功能。瞧！

```go
var driverName, aliasName string
// driverName name of your driver (go-sql-driver: mysql)
// aliasName custom db alias name
var db *sql.DB
...
o := orm.NewOrmWithDB(driverName, aliasName, db)
```

#### GetDB

Get *sql.DB from the registered databases. This will use `default` as default if you do not set.

​	从已注册的数据库中获取 *sql.DB。如果您未设置，这将使用 `default` 作为默认值。

```go
db, err := orm.GetDB()
if err != nil {
	fmt.Println("get default DataBase")
}

db, err := orm.GetDB("alias")
if err != nil {
	fmt.Println("get alias DataBase")
}
```

#### ResetModelCache

Reset registered models. Commonly used to write test cases.

​	重置已注册的模型。通常用于编写测试用例。

```go
orm.ResetModelCache()
```

## ORM API Usage ORM API 使用

Let’s see how to use Ormer API:

​	让我们看看如何使用 Ormer API：

```go
var o orm.Ormer
o = orm.NewOrm() // create a Ormer // While running NewOrm, it will run orm.BootStrap (only run once in the whole app lifetime) to validate the definition between models and cache it.
```

If you want to use DB transaction，we will return `TxOrm` instance [ORM Transaction]({{< ref "/beego/mvcIntroduction/models/transaction" >}})

​	如果您想使用 DB 事务，我们将返回 `TxOrm` 实例 ORM 事务

Comparing with v1.x, we designed another interface `TxOrm` to handle transaction.

​	与 v1.x 相比，我们设计了另一个接口 `TxOrm` 来处理事务。

From v1.x, we found that many users reuse global ORM instance to handle transaction. It made unpredictable result.

​	从 v1.x 开始，我们发现许多用户重复使用全局 ORM 实例来处理事务。这导致了不可预测的结果。

When you use `TxOrm`, you should drop it after ending transaction. It’s stateful object.

​	当您使用 `TxOrm` 时，您应该在结束事务后将其删除。它是状态对象。

- type Ormer interface {
  - [Read(interface{}, …string) error]({{< ref "/beego/mvcIntroduction/models/crudOperations#read" >}})
  - [ReadOrCreate(interface{}, string, …string) (bool, int64, error)]({{< ref "/beego/mvcIntroduction/models/crudOperations#readorcreate" >}})
  - [Insert(interface{}) (int64, error)]({{< ref "/beego/mvcIntroduction/models/crudOperations#insert" >}})
  - [InsertMulti(int, interface{}) (int64, error)]({{< ref "/beego/mvcIntroduction/models/crudOperations#insertmulti" >}})
  - [Update(interface{}, …string) (int64, error)
    更新(接口{}, …字符串) (int64, 错误)]({{< ref "/beego/mvcIntroduction/models/crudOperations#update" >}})
  - [Delete(interface{}) (int64, error)
    删除(接口{}) (int64, 错误)]({{< ref "/beego/mvcIntroduction/models/crudOperations#delete" >}})
  - [LoadRelated(interface{}, string, …interface{}) (int64, error)
    加载相关(接口{}, 字符串, …接口{}) (int64, 错误)]({{< ref "/beego/mvcIntroduction/models/advancedQueries#load-related-field" >}})
  - [QueryM2M(interface{}, string) QueryM2Mer
    查询M2M(接口{}, 字符串) QueryM2Mer]({{< ref "/beego/mvcIntroduction/models/advancedQueries#handling-manytomany-relation" >}})
  - [QueryTable(interface{}) QuerySeter
    查询表(接口{}) QuerySeter](https://beego.wiki/docs/mvc/model/orm/#querytable)
  - [Begin() error
    开始() 错误]({{< ref "/beego/mvcIntroduction/models/transaction" >}})
  - [Commit() error
    提交() 错误]({{< ref "/beego/mvcIntroduction/models/transaction" >}})
  - [Rollback() error
    回滚() 错误]({{< ref "/beego/mvcIntroduction/models/transaction" >}})
  - [Raw(string, …interface{}) RawSeter
    原始(字符串, …接口{}) RawSeter](https://beego.wiki/docs/mvc/model/orm/#raw)
  - [Driver() Driver
    驱动程序() 驱动程序](https://beego.wiki/docs/mvc/model/orm/#driver)
- }

#### QueryTable

Pass in a table name or a Model object and return a [QuerySeter]({{< ref "/beego/mvcIntroduction/models/advancedQueries#queryseter" >}})

​	传入一个表名或一个 Model 对象并返回一个 QuerySeter

```go
o := orm.NewOrm()
var qs orm.QuerySeter
qs = o.QueryTable("user")
// Panics if the table can't be found
```

#### NewOrmUsingDB

We remove `Using` method since some users use this method in wrong way and then met some concurrent problems.

​	我们移除了 `Using` 方法，因为一些用户错误地使用此方法，然后遇到了一些并发问题。

You can use `NewOrmUsingDB`:

​	您可以使用 `NewOrmUsingDB` ：

```go
o := orm.NewOrmUsingDB("db_name")
```

#### Raw

Use raw SQL query:

​	使用原始 SQL 查询：

Raw function will return a [RawSeter]({{< ref "/beego/mvcIntroduction/models/rawSQLToQuery" >}}) to execute a query with the SQL and params provided:

​	Raw 函数将返回一个 RawSeter，以使用提供的 SQL 和参数执行查询：

```go
o := NewOrm()
var r orm.RawSeter
r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")
```

#### Driver

The current db infomation used by ORM

​	ORM 使用的当前数据库信息

```go
type Driver interface {
	Name() string
	Type() DriverType
}
orm.RegisterDataBase("db1", "mysql", "root:root@/orm_db2?charset=utf8")
orm.RegisterDataBase("db2", "sqlite3", "data.db")

o1 := orm.NewOrmUsingDB("db1")
dr := o1.Driver()
fmt.Println(dr.Name() == "db1") // true
fmt.Println(dr.Type() == orm.DRMySQL) // true

o2 := orm.NewOrmUsingDB("db2")
dr = o2.Driver()
fmt.Println(dr.Name() == "db2") // true
fmt.Println(dr.Type() == orm.DRSqlite) // true
```

## Print out SQL queries in debugging mode 在调试模式下打印出 SQL 查询

Setting `orm.Debug` to true will print out SQL queries.

​	将 `orm.Debug` 设置为 true 将打印出 SQL 查询。

It may cause performance issues. It is not recommended to be used in a production env.

​	它可能会导致性能问题。不建议在生产环境中使用。

```go
func main() {
	orm.Debug = true
...
```

Prints to `os.Stderr` by default.

​	默认情况下打印到 `os.Stderr` 。

You can change it to your own `io.Writer`

​	您可以将其更改为您自己的 `io.Writer`

```go
var w io.Writer
...
// Use your `io.Writer`
...
orm.DebugLog = orm.NewLog(w)
```

Logs formatting

​	日志格式化

```go
[ORM] - time - [Queries/database name] - [operation/executing time] - [SQL query] - separate params with `,`  -errors 
[ORM] - 2013-08-09 13:18:16 - [Queries/default] - [    db.Exec /     0.4ms] - [INSERT INTO `user` (`name`) VALUES (?)] - `slene`
[ORM] - 2013-08-09 13:18:16 - [Queries/default] - [    db.Exec /     0.5ms] - [UPDATE `user` SET `name` = ? WHERE `id` = ?] - `astaxie`, `14`
[ORM] - 2013-08-09 13:18:16 - [Queries/default] - [db.QueryRow /     0.4ms] - [SELECT `id`, `name` FROM `user` WHERE `id` = ?] - `14`
[ORM] - 2013-08-09 13:18:16 - [Queries/default] - [    db.Exec /     0.4ms] - [INSERT INTO `post` (`user_id`,`title`,`content`) VALUES (?, ?, ?)] - `14`, `beego orm`, `powerful amazing`
[ORM] - 2013-08-09 13:18:16 - [Queries/default] - [   db.Query /     0.4ms] - [SELECT T1.`name` `User__Name`, T0.`user_id` `User`, T1.`id` `User__Id` FROM `post` T0 INNER JOIN `user` T1 ON T1.`id` = T0.`user_id` WHERE T0.`id` = ? LIMIT 1000] - `68`
[ORM] - 2013-08-09 13:18:16 - [Queries/default] - [    db.Exec /     0.4ms] - [DELETE FROM `user` WHERE `id` = ?] - `14`
[ORM] - 2013-08-09 13:18:16 - [Queries/default] - [   db.Query /     0.3ms] - [SELECT T0.`id` FROM `post` T0 WHERE T0.`user_id` IN (?) ] - `14`
[ORM] - 2013-08-09 13:18:16 - [Queries/default] - [    db.Exec /     0.4ms] - [DELETE FROM `post` WHERE `id` IN (?)] - `68`
```

The log contains all the database operations, transactions, prepare etc.

​	日志包含所有数据库操作、事务、准备等。
