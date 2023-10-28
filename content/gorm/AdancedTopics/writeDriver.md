+++
title = "Write Driver"
date = 2023-10-28T14:37:09+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/write_driver.html](https://gorm.io/docs/write_driver.html)

## Write new driver

GORM provides official support for `sqlite`, `mysql`, `postgres`, `sqlserver`.

Some databases may be compatible with the `mysql` or `postgres` dialect, in which case you could just use the dialect for those databases.

For others, you can create a new driver, it needs to implement [the dialect interface](https://pkg.go.dev/gorm.io/gorm?tab=doc#Dialector).

```
type Dialector interface {
  Name() string
  Initialize(*DB) error
  Migrator(db *DB) Migrator
  DataTypeOf(*schema.Field) string
  DefaultValueOf(*schema.Field) clause.Expression
  BindVarTo(writer clause.Writer, stmt *Statement, v interface{})
  QuoteTo(clause.Writer, string)
  Explain(sql string, vars ...interface{}) string
}
```

Checkout the [MySQL Driver](https://github.com/go-gorm/mysql) as example