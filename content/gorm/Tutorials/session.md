+++
title = "Session"
date = 2023-10-28T14:30:17+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文：[https://gorm.io/docs/session.html](https://gorm.io/docs/session.html)

GORM provides `Session` method, which is a [`New Session Method`](https://gorm.io/docs/method_chaining.html), it allows to create a new session mode with configuration:

​	GORM提供了一个`Session`方法，它是一个[新会话方法](../methodChaining)，允许使用配置创建一个新的会话模式：

``` go
// Session Configuration
type Session struct {
  DryRun                   bool
  PrepareStmt              bool
  NewDB                    bool
  Initialized              bool
  SkipHooks                bool
  SkipDefaultTransaction   bool
  DisableNestedTransaction bool
  AllowGlobalUpdate        bool
  FullSaveAssociations     bool
  QueryFields              bool
  Context                  context.Context
  Logger                   logger.Interface
  NowFunc                  func() time.Time
  CreateBatchSize          int
}
```

## DryRun

Generates `SQL` without executing. It can be used to prepare or test generated SQL, for example:

​	生成`SQL`而不执行。它可以用于准备或测试生成的SQL，例如：

``` go
// 会话模式 session mode
stmt := db.Session(&Session{DryRun: true}).First(&user, 1).Statement
stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
stmt.Vars         //=> []interface{}{1}

// 全局模式下的DryRun globally mode with DryRun
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{DryRun: true})

// 不同的数据库生成不同的SQL different databases generate different SQL
stmt := db.Find(&user, 1).Statement
stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 // PostgreSQL
stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = ?  // MySQL
stmt.Vars         //=> []interface{}{1}
```

To generate the final SQL, you could use following code:

​	要生成最终的SQL，可以使用以下代码：

``` go
// 注意：SQL并不总是安全的执行，GORM仅用于日志记录，它可能导致SQL注入 NOTE: the SQL is not always safe to execute, GORM only uses it for logs, it might cause SQL injection
db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
// SELECT * FROM `users` WHERE `id` = 1
```

## PrepareStmt

`PreparedStmt` creates prepared statements when executing any SQL and caches them to speed up future calls, for example:

​	`PreparedStmt`在执行任何SQL时创建预处理语句并缓存它们以加速未来的调用，例如：

``` go
// 全局模式下，所有数据库操作将创建预处理语句并缓存它们 globally mode, all DB operations will create prepared statements and cache them
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  PrepareStmt: true,
})

// 会话模式 session mode
tx := db.Session(&Session{PrepareStmt: true})
tx.First(&user, 1)
tx.Find(&users)
tx.Model(&user).Update("Age", 18)

// 返回预处理语句管理器 returns prepared statements manager
stmtManger, ok := tx.ConnPool.(*PreparedStmtDB)

// 关闭当前会话的预处理语句 close prepared statements for *current session*
stmtManger.Close()

// 当前会话的预处理SQL prepared SQL for *current session*
stmtManger.PreparedSQL // => []string{}

// 当前数据库连接池（所有会话）的预处理语句 prepared statements for current database connection pool (all sessions)
stmtManger.Stmts // map[string]*sql.Stmt

for sql, stmt := range stmtManger.Stmts {
  sql  // 预处理SQL prepared SQL
  stmt // 预处理语句 prepared statement
  stmt.Close() // 关闭预处理语句 close the prepared statement
}
```

## NewDB

Create a new DB without conditions with option `NewDB`, for example:

​	创建一个没有条件的新DB，使用选项`NewDB`，例如：

``` go
tx := db.Where("name = ?", "jinzhu").Session(&gorm.Session{NewDB: true})

tx.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1

tx.First(&user, "id = ?", 10)
// SELECT * FROM users WHERE id = 10 ORDER BY id

// 不使用选项`NewDB` Without option `NewDB`
tx2 := db.Where("name = ?", "jinzhu").Session(&gorm.Session{})
tx2.First(&user)
// SELECT * FROM users WHERE name = "jinzhu" ORDER BY id
```

## Initialized

Create a new initialized DB, which is not Method Chain/Goroutine Safe anymore, refer [Method Chaining](https://gorm.io/docs/method_chaining.html)

​	创建一个初始化的DB，它不再是Method Chain/Goroutine Safe anymore，参考[Method Chaining](../methodChaining)

``` go
tx := db.Session(&gorm.Session{Initialized: true})
```

## Skip Hooks

If you want to skip `Hooks` methods, you can use the `SkipHooks` session mode, for example:

​	如果你想要跳过`Hooks`方法，你可以使用`SkipHooks`会话模式，例如：

``` go
DB.Session(&gorm.Session{SkipHooks: true}).Create(&user)

DB.Session(&gorm.Session{SkipHooks: true}).Create(&users)

DB.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)

DB.Session(&gorm.Session{SkipHooks: true}).Find(&user)

DB.Session(&gorm.Session{SkipHooks: true}).Delete(&user)

DB.Session(&gorm.Session{SkipHooks: true}).Model(User{}).Where("age > ?", 18).Updates(&user)
```

## DisableNestedTransaction

When using `Transaction` method inside a DB transaction, GORM will use `SavePoint(savedPointName)`, `RollbackTo(savedPointName)` to give you the nested transaction support. You can disable it by using the `DisableNestedTransaction` option, for example:

​	当在DB事务内使用`Transaction`方法时，GORM将使用`SavePoint(savedPointName)`，`RollbackTo(savedPointName)`来给你提供嵌套事务支持。你可以通过设置此选项为true来启用它，例如：

``` go
db.Session(&gorm.Session{
  DisableNestedTransaction: true,
}).CreateInBatches(&users, 100)
```

## AllowGlobalUpdate

GORM doesn’t allow global update/delete by default, will return `ErrMissingWhereClause` error. You can set this option to true to enable it, for example:

​	GORM默认不允许全局更新/删除，将返回`ErrMissingWhereClause`错误。您可以将此选项设置为true以启用它，例如：

``` go
db.Session(&gorm.Session{
  AllowGlobalUpdate: true,
}).Model(&User{}).Update("name", "jinzhu")
// UPDATE users SET `name` = "jinzhu"
```

## FullSaveAssociations

GORM will auto-save associations and its reference using [Upsert](https://gorm.io/docs/create.html#upsert) when creating/updating a record. If you want to update associations’ data, you should use the `FullSaveAssociations` mode, for example:

​	GORM在创建/更新记录时将自动使用[Upsert]({{< ref "/gorm/CRUDInterface/create#upsert--on-conflict">}})保存关联及其引用。如果你想更新关联的数据，你应该使用`FullSaveAssociations`模式，例如：

``` go
db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
// ...
// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY SET address1=VALUES(address1);
// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY SET email=VALUES(email);
// ...
```

## Context

With the `Context` option, you can set the `Context` for following SQL operations, for example:

​	使用`Context`选项，你可以为后续的SQL操作设置`Context`，例如：

``` go
timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second)
tx := db.Session(&Session{Context: timeoutCtx})

tx.First(&user) // 使用timeoutCtx查询 query with context timeoutCtx
tx.Model(&user).Update("role", "admin") // 使用timeoutCtx更新 update with context timeoutCtx
```

GORM also provides shortcut method `WithContext`, here is the definition:

​	GORM还提供了快捷方法`WithContext`，这里是定义：

``` go
func (db *DB) WithContext(ctx context.Context) *DB {
  return db.Session(&Session{Context: ctx})
}
```

## 日志记录器 Logger

Gorm allows customizing built-in logger with the `Logger` option, for example:

​	Gorm允许使用`Logger`选项自定义内置日志记录器，例如：

``` go
newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
              logger.Config{
                SlowThreshold: time.Second,
                LogLevel:      logger.Silent,
                Colorful:      false,
              })
db.Session(&Session{Logger: newLogger})

db.Session(&Session{Logger: logger.Default.LogMode(logger.Silent)})
```

Checkout [Logger](https://gorm.io/docs/logger.html) for more details.

​	更多详细信息请查看[日志记录器]({{< ref "/gorm/Tutorials/logger">}})。

## NowFunc

`NowFunc` allows changing the function to get current time of GORM, for example:

​	`NowFunc`允许更改获取GORM当前时间的函数，例如：

``` go
db.Session(&Session{
  NowFunc: func() time.Time {
    return time.Now().Local()
  },
})
```

## 调试 Debug

`Debug` is a shortcut method to change session’s `Logger` to debug mode, here is the definition:

​	`Debug`是一个快捷方法，用于将会话的`Logger`更改为调试模式，以下是定义：

``` go
func (db *DB) Debug() (tx *DB) {
  return db.Session(&Session{
    Logger:         db.Logger.LogMode(logger.Info),
  })
}
```

## QueryFields

Select by fields

​	按字段选择

``` go
db.Session(&gorm.Session{QueryFields: true}).Find(&user)
// SELECT `users`.`name`, `users`.`age`, ... FROM `users` // with this option
// SELECT * FROM `users` // without this option
```

## CreateBatchSize

Default batch size

​	默认批处理大小

``` go
users = [5000]User{{Name: "jinzhu", Pets: []Pet{pet1, pet2, pet3}}...}

db.Session(&gorm.Session{CreateBatchSize: 1000}).Create(&users)
// INSERT INTO users xxx (5 batches)
// INSERT INTO pets xxx (15 batches)
```