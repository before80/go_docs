+++
title = "GORM 配置"
date = 2023-10-28T14:36:45+08:00
weight = 11
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/gorm_config.html](https://gorm.io/docs/gorm_config.html)

GORM provides Config can be used during initialization

​	GORM提供了Config可以在初始化时使用。

``` go
type Config struct {
  SkipDefaultTransaction   bool
  NamingStrategy           schema.Namer
  Logger                   logger.Interface
  NowFunc                  func() time.Time
  DryRun                   bool
  PrepareStmt              bool
  DisableNestedTransaction bool
  AllowGlobalUpdate        bool
  DisableAutomaticPing     bool
  DisableForeignKeyConstraintWhenMigrating bool
}
```

## SkipDefaultTransaction

GORM perform write (create/update/delete) operations run inside a transaction to ensure data consistency, you can disable it during initialization if it is not required

​	GORM在执行写（创建/更新/删除）操作时，会运行在一个事务中以确保数据一致性，如果在初始化时不需要，可以禁用它

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
})
```

## NamingStrategy

GORM allows users to change the naming conventions by overriding the default `NamingStrategy` which need to implements interface `Namer`

​	GORM允许用户通过覆盖默认的`NamingStrategy`来改变命名约定，需要实现接口`Namer`

``` go
type Namer interface {
  TableName(table string) string
  SchemaName(table string) string
  ColumnName(table, column string) string
  JoinTableName(table string) string
  RelationshipFKName(Relationship) string
  CheckerName(table, column string) string
  IndexName(table, column string) string
}
```

The default `NamingStrategy` also provides few options, like:

​	默认的`NamingStrategy`还提供了一些选项，例如：

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  NamingStrategy: schema.NamingStrategy{
    TablePrefix: "t_",   // table name prefix, table for `User` would be `t_users`
    SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
    NoLowerCase: true, // skip the snake_casing of names
    NameReplacer: strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
  },
})
```

## Logger

Allow to change GORM’s default logger by overriding this option, refer [Logger](https://gorm.io/docs/logger.html) for more details

​	允许通过覆盖此选项来更改GORM的默认日志记录器，有关更多详细信息，请参阅[Logger](https://gorm.io/docs/logger.html)

## NowFunc

Change the function to be used when creating a new timestamp

​	更改创建新时间戳时要使用的函数

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  NowFunc: func() time.Time {
    return time.Now().Local()
  },
})
```

## DryRun

Generate `SQL` without executing, can be used to prepare or test generated SQL, refer [Session](https://gorm.io/docs/session.html) for details

​	生成`SQL`而不执行，可用于准备或测试生成的SQL，有关详细信息，请参阅[Session](https://gorm.io/docs/session.html)

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  DryRun: false,
})
```

## PrepareStmt

`PreparedStmt` creates a prepared statement when executing any SQL and caches them to speed up future calls, refer [Session](https://gorm.io/docs/session.html) for details

​	`PreparedStmt`在执行任何SQL时创建一个预处理语句并将其缓存以加速未来的调用，有关详细信息，请参阅[Session](https://gorm.io/docs/session.html)

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  PrepareStmt: false,
})
```

## DisableNestedTransaction

When using `Transaction` method inside a db transaction, GORM will use `SavePoint(savedPointName)`, `RollbackTo(savedPointName)` to give you the nested transaction support, you could disable it by using the `DisableNestedTransaction` option, refer [Session](https://gorm.io/docs/session.html) for details

​	在使用db事务内的`Transaction`方法时，GORM将使用`SavePoint(savedPointName)`，`RollbackTo(savedPointName)`提供嵌套事务支持，您可以通过使用`DisableNestedTransaction`选项禁用它，有关详细信息，请参阅[Session](https://gorm.io/docs/session.html)

## AllowGlobalUpdate

Enable global update/delete, refer [Session](https://gorm.io/docs/session.html) for details

​	启用全局更新/删除，有关详细信息，请参阅[Session](https://gorm.io/docs/session.html)

## DisableAutomaticPing

GORM automatically ping database after initialized to check database availability, disable it by setting it to `true`

​	GORM在初始化后自动ping数据库以检查数据库可用性，通过将其设置为`true`禁用它

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  DisableAutomaticPing: true,
})
```

## DisableForeignKeyConstraintWhenMigrating

GORM creates database foreign key constraints automatically when `AutoMigrate` or `CreateTable`, disable this by setting it to `true`, refer [Migration](https://gorm.io/docs/migration.html) for details

​	GORM在`AutoMigrate`或`CreateTable`时自动创建数据库外键约束，禁用此选项，将其设置为`true`，有关详细信息，请参阅[Migration](https://gorm.io/docs/migration.html)

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  DisableForeignKeyConstraintWhenMigrating: true,
})
```