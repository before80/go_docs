+++
title = "自定义数据类型"
date = 2023-10-28T14:32:12+08:00
weight = 11
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/data_types.html](https://gorm.io/docs/data_types.html)

GORM provides few interfaces that allow users to define well-supported customized data types for GORM, takes [json](https://github.com/go-gorm/datatypes/blob/master/json.go) as an example

​	GORM提供了一些接口，允许用户为GORM定义支持的自定义数据类型。以[json](https://github.com/go-gorm/datatypes/blob/master/json.go)为例。

## 实现自定义数据类型 Implements Customized Data Type

### Scanner / Valuer

The customized data type has to implement the [Scanner](https://pkg.go.dev/database/sql#Scanner) and [Valuer](https://pkg.go.dev/database/sql/driver#Valuer) interfaces, so GORM knowns to how to receive/save it into the database

​	自定义数据类型必须实现[Scanner]({{< ref "/stdLib/database/sql#type-scanner">}})和[Valuer]({{< ref "/stdLib/database/sql_driver#type-valuer">}})接口，以便GORM知道如何将值接收/保存到数据库中。

For example:

​	例如：

``` go
type JSON json.RawMessage

// 将扫描值扫描到Jsonb中，实现sql.Scanner接口 Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
  bytes, ok := value.([]byte)
  if !ok {
    return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
  }

  result := json.RawMessage{}
  err := json.Unmarshal(bytes, &result)
  *j = JSON(result)
  return err
}

// 返回json值，实现driver.Valuer接口 Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
  if len(j) == 0 {
    return nil, nil
  }
  return json.RawMessage(j).MarshalJSON()
}
```

There are many third party packages implement the `Scanner`/`Valuer` interface, which can be used with GORM together, for example:

​	有许多第三方包实现了`扫描器`/`值器`接口，可以与GORM一起使用，例如：

``` go
import (
  "github.com/google/uuid"
  "github.com/lib/pq"
)

type Post struct {
  ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
  Title  string
  Tags   pq.StringArray `gorm:"type:text[]"`
}
```

### GormDataTypeInterface

GORM will read column’s database type from [tag](https://gorm.io/docs/models.html#tags) `type`, if not found, will check if the struct implemented interface `GormDBDataTypeInterface` or `GormDataTypeInterface` and will use its result as data type

​	GORM将从[标签]({{< ref "/gorm/GettingStarted/declaringModels#字段标签-fields-tags">}})`type`读取列的数据库类型，如果找不到，它将检查结构是否实现了接口`GormDBDataTypeInterface`或`GormDataTypeInterface`，并使用其结果作为数据类型。

``` go
type GormDataTypeInterface interface {
  GormDataType() string
}

type GormDBDataTypeInterface interface {
  GormDBDataType(*gorm.DB, *schema.Field) string
}
```

The result of `GormDataType` will be used as the general data type and can be obtained from `schema.Field`‘s field `DataType`, which might be helpful when [writing plugins](https://gorm.io/docs/write_plugins.html) or [hooks](https://gorm.io/docs/hooks.html) for example:

​	`GormDataType`的结果将用作通用数据类型，可以从`schema.Field`的字段`DataType`中获得，这可能在[编写插件]({{< ref "/gorm/AdancedTopics/writePlugins">}})或[钩子]({{< ref "/gorm/Tutorials/hooks">}})时有用，例如：

``` go
func (JSON) GormDataType() string {
  return "json"
}

type User struct {
  Attrs JSON
}

func (user User) BeforeCreate(tx *gorm.DB) {
  field := tx.Statement.Schema.LookUpField("Attrs")
  if field.DataType == "json" {
    // do something
  }
}
```

`GormDBDataType` usually returns the right data type for current driver when migrating, for example:

​	`GormDBDataType`通常在迁移时根据驱动程序名称返回正确的数据类型，例如：

``` go
func (JSON) GormDBDataType(db *gorm.DB, field *schema.Field) string {
  // 使用field.Tag和field.TagSettings获取字段的标签 use field.Tag, field.TagSettings gets field's tags
  // 请查看https://github.com/go-gorm/gorm/blob/master/schema/field.go以获取所有选项 checkout https://github.com/go-gorm/gorm/blob/master/schema/field.go for all options

  // 根据驱动程序名称返回不同的数据库类型 returns different database type based on driver name
  switch db.Dialector.Name() {
  case "mysql", "sqlite":
    return "JSON"
  case "postgres":
    return "JSONB"
  }
  return ""
}
```

If the struct hasn’t implemented the `GormDBDataTypeInterface` or `GormDataTypeInterface` interface, GORM will guess its data type from the struct’s first field, for example, will use `string` for `NullString`

​	如果结构没有实现`GormDBDataTypeInterface`或`GormDataTypeInterface`接口，GORM将从结构的第一个字段猜测其数据类型，例如，将`NullString`的数据类型设置为字符串。

``` go
type NullString struct {
  String string // 使用第一个字段的数据类型 use the first field's data type
  Valid  bool
}

type User struct {
  Name NullString // 数据类型将是字符串 data type will be string
}
```

### GormValuerInterface

GORM provides a `GormValuerInterface` interface, which can allow to create/update from SQL Expr or value based on context, for example:

​	GORM提供了一个`GormValuerInterface`接口，可以根据上下文从SQL表达式或值创建/更新，例如：

``` go
// GORM Valuer接口 GORM Valuer interface
type GormValuerInterface interface {
  GormValue(ctx context.Context, db *gorm.DB) clause.Expr
}
```

#### 从SQL表达式创建/更新 Create/Update from SQL Expr

``` go
type Location struct {
  X, Y int
}

func (loc Location) GormDataType() string {
  return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
  return clause.Expr{
    SQL:  "ST_PointFromText(?)",
    Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
  }
}

// Scan实现了sql.Scanner接口 Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
  // 从数据库驱动程序中将值扫描到结构体中 Scan a value into struct from database driver
}

type User struct {
  ID       int
  Name     string
  Location Location
}

db.Create(&User{
  Name:     "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// INSERT INTO `users` (`name`,`point`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"))

db.Model(&User{ID: 1}).Updates(User{
  Name:  "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// UPDATE `user_with_points` SET `name`="jinzhu",`location`=ST_PointFromText("POINT(100 100)") WHERE `id` = 1
```

You can also create/update with SQL Expr from map, checkout [Create From SQL Expr](https://gorm.io/docs/create.html#create_from_sql_expr) and [Update with SQL Expression](https://gorm.io/docs/update.html#update_from_sql_expr) for details

​	您还可以从映射创建/更新带有SQL表达式，请参阅[从SQL表达式创建]({{< ref "/gorm/CRUDInterface/create#从-sql-表达式上下文值器创建-create-from-sql-expressioncontext-valuer">}})和[使用SQL表达式更新]({{< ref "/gorm/CRUDInterface/update#使用-sql-表达式进行更新-update-with-sql-expression">}})以获取详细信息。

#### 基于上下文的值 Value based on Context

If you want to create or update a value depends on current context, you can also implements the `GormValuerInterface` interface, for example:

​	如果你想根据当前上下文创建或更新值，你也可以实现`GormValuerInterface`接口，例如：

``` go
type EncryptedString struct {
  Value string
}

func (es EncryptedString) GormValue(ctx context.Context, db *gorm.DB) (expr clause.Expr) {
  if encryptionKey, ok := ctx.Value("TenantEncryptionKey").(string); ok {
    return clause.Expr{SQL: "?", Vars: []interface{}{Encrypt(es.Value, encryptionKey)}}
  } else {
    db.AddError(errors.New("invalid encryption key"))
  }

  return
}
```

### 子句表达式 Clause Expression

If you want to build some query helpers, you can make a struct that implements the `clause.Expression` interface:

​	如果你想构建一些查询辅助函数，你可以创建一个实现`clause.Expression`接口的结构体：

``` go
type Expression interface {
  Build(builder Builder)
}
```

Checkout [JSON](https://github.com/go-gorm/datatypes/blob/master/json.go) and [SQL Builder](https://gorm.io/docs/sql_builder.html#clauses) for details, the following is an example of usage:

​	详细信息请查看[JSON](https://github.com/go-gorm/datatypes/blob/master/json.go)和[SQL Builder]({{< ref "/gorm/CRUDInterface/rawSQLAndSQLBuilder#子句-clauses">}})，以下是一个使用示例：

``` go
// 使用子句表达式生成SQL Generates SQL with clause Expression
db.Find(&user, datatypes.JSONQuery("attributes").HasKey("role"))
db.Find(&user, datatypes.JSONQuery("attributes").HasKey("orgs", "orga"))

// MySQL
// SELECT * FROM `users` WHERE JSON_EXTRACT(`attributes`, '$.role') IS NOT NULL
// SELECT * FROM `users` WHERE JSON_EXTRACT(`attributes`, '$.orgs.orga') IS NOT NULL

// PostgreSQL
// SELECT * FROM "user" WHERE "attributes"::jsonb ? 'role'
// SELECT * FROM "user" WHERE "attributes"::jsonb -> 'orgs' ? 'orga'

db.Find(&user, datatypes.JSONQuery("attributes").Equals("jinzhu", "name"))
// MySQL
// SELECT * FROM `user` WHERE JSON_EXTRACT(`attributes`, '$.name') = "jinzhu"

// PostgreSQL
// SELECT * FROM "user" WHERE json_extract_path_text("attributes"::json,'name') = 'jinzhu'
```

## 自定义数据类型集合 Customized Data Types Collections

We created a Github repo for customized data types collections https://github.com/go-gorm/datatypes, pull request welcome ;)

​	我们为自定义数据类型集合创建了一个Github仓库 [https://github.com/go-gorm/datatypes](https://github.com/go-gorm/datatypes)，欢迎提交pull request！