+++
title = "GORM Config"
date = 2023-10-28T14:36:45+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/gorm_config.html](https://gorm.io/docs/gorm_config.html)

GORM provides Config can be used during initialization

```
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

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
})
```

## NamingStrategy

GORM allows users to change the naming conventions by overriding the default `NamingStrategy` which need to implements interface `Namer`

```
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

```
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

## NowFunc

Change the function to be used when creating a new timestamp

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  NowFunc: func() time.Time {
    return time.Now().Local()
  },
})
```

## DryRun

Generate `SQL` without executing, can be used to prepare or test generated SQL, refer [Session](https://gorm.io/docs/session.html) for details

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  DryRun: false,
})
```

## PrepareStmt

`PreparedStmt` creates a prepared statement when executing any SQL and caches them to speed up future calls, refer [Session](https://gorm.io/docs/session.html) for details

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  PrepareStmt: false,
})
```

## DisableNestedTransaction

When using `Transaction` method inside a db transaction, GORM will use `SavePoint(savedPointName)`, `RollbackTo(savedPointName)` to give you the nested transaction support, you could disable it by using the `DisableNestedTransaction` option, refer [Session](https://gorm.io/docs/session.html) for details

## AllowGlobalUpdate

Enable global update/delete, refer [Session](https://gorm.io/docs/session.html) for details

## DisableAutomaticPing

GORM automatically ping database after initialized to check database availability, disable it by setting it to `true`

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  DisableAutomaticPing: true,
})
```

## DisableForeignKeyConstraintWhenMigrating

GORM creates database foreign key constraints automatically when `AutoMigrate` or `CreateTable`, disable this by setting it to `true`, refer [Migration](https://gorm.io/docs/migration.html) for details

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  DisableForeignKeyConstraintWhenMigrating: true,
})
```