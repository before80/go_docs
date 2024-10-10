+++
title = "约定 Conventions"
date = 2023-10-28T14:32:37+08:00
weight = 13
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/conventions.html](https://gorm.io/docs/conventions.html)

## `ID` 作为主键 `ID` as Primary Key

GORM uses the field with the name `ID` as the table’s primary key by default.

​	GORM 默认将名为 `ID` 的字段作为表的主键。

``` go
type User struct {
  ID   string // 默认情况下，名为 `ID` 的字段将被用作主键 field named `ID` will be used as a primary field by default
  Name string
}
```

You can set other fields as primary key with tag `primaryKey`

​	您可以使用带有 `primaryKey` 标签的其他字段设置其他字段为主键。

``` go
// 将 `UUID` 字段设置为主键 Set field `UUID` as primary field
type Animal struct {
  ID     int64
  UUID   string `gorm:"primaryKey"`
  Name   string
  Age    int64
}
```

Also check out [Composite Primary Key](https://gorm.io/docs/composite_primary_key.html)

​	还可以查看 [复合主键]({{< ref "/gorm/AdancedTopics/compositePrimaryKey">}})。

## 复数形式的名称 Pluralized Table Name

GORM pluralizes struct name to `snake_cases` as table name, for struct `User`, its table name is `users` by convention

​	GORM 将结构体名称转换为 `snake_cases` 作为表名，对于结构体 `User`，其表名 convention 为 `users`。

### TableName

You can change the default table name by implementing the `Tabler` interface, for example:

​	您可以通过实现 `Tabler` 接口来更改默认表名，例如：

``` go
type Tabler interface {
  TableName() string
}

// 通过 User 覆盖表名以使用 `profiles` TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
  return "profiles"
}
```

> **NOTE** `TableName` doesn’t allow dynamic name, its result will be cached for future, to use dynamic name, you can use `Scopes`, for example:
>
> **注意** `TableName` 不允许动态命名，其结果将被缓存以供将来使用，要使用动态名称，您可以使用 `Scopes`，例如：

``` go
func UserTable(user User) func (tx *gorm.DB) *gorm.DB {
  return func (tx *gorm.DB) *gorm.DB {
    if user.Admin {
      return tx.Table("admin_users")
    }

    return tx.Table("users")
  }
}

db.Scopes(UserTable(user)).Create(&user)
```

### 临时指定名称 Temporarily specify a name

Temporarily specify table name with `Table` method, for example:

​	使用 `Table` 方法临时指定表名，例如：

``` go
// 使用结构体 User 的字段创建名为 `deleted_users` 的表 Create table `deleted_users` with struct User's fields
db.Table("deleted_users").AutoMigrate(&User{})

// 从另一个表中查询数据 Query data from another table
var deletedUsers []User
db.Table("deleted_users").Find(&deletedUsers)
// SELECT * FROM deleted_users;

db.Table("deleted_users").Where("name = ?", "jinzhu").Delete(&User{})
// DELETE FROM deleted_users WHERE name = 'jinzhu';
```

Check out [From SubQuery](https://gorm.io/docs/advanced_query.html#from_subquery) for how to use SubQuery in FROM clause

​	查看 [From SubQuery]({{< ref "/gorm/CRUDInterface/advancedQuery#from-子查询-from-subquery">}}) 了解如何在 FROM 子句中使用 SubQuery。

### NamingStrategy

GORM allows users change the default naming conventions by overriding the default `NamingStrategy`, which is used to build `TableName`, `ColumnName`, `JoinTableName`, `RelationshipFKName`, `CheckerName`, `IndexName`, Check out [GORM Config](https://gorm.io/docs/gorm_config.html#naming_strategy) for details

​	GORM 允许用户通过覆盖默认的 `NamingStrategy` 来更改默认的命名约定，该策略用于构建 `TableName`、`ColumnName`、`JoinTableName`、`RelationshipFKName`、`CheckerName`、`IndexName`，请参阅 [GORM Config]({{< ref "/gorm/AdancedTopics/gormConfig">}}) 以获取详细信息。

## 列名 Column Name

Column db name uses the field’s name’s `snake_case` by convention.

​	列的数据库名称使用字段的 `snake_case` 命名约定。

``` go
type User struct {
  ID        uint      // 列名为 `id` column name is `id`
  Name      string    // 列名为 `name` column name is `name`
  Birthday  time.Time // 列名为 `birthday` column name is `birthday`
  CreatedAt time.Time // 列名为 `created_at` column name is `created_at`
}
```

You can override the column name with tag `column` or use [`NamingStrategy`](https://gorm.io/docs/conventions.html#naming_strategy)

​	您可以使用标签 `column` 覆盖列名，或使用 [NamingStrategy]({{< ref "/gorm/Tutorials/conventions#namingstrategy">}})

``` go
type Animal struct {
  AnimalID int64     `gorm:"column:beast_id"`         // 将名称设置为 `beast_id` set name to `beast_id`
  Birthday time.Time `gorm:"column:day_of_the_beast"` // 将名称设置为 `day_of_the_beast` set name to `day_of_the_beast`
  Age      int64     `gorm:"column:age_of_the_beast"` // 将名称设置为 `age_of_the_beast` set name to `age_of_the_beast`
}
```

## 时间戳跟踪 Timestamp Tracking

### CreatedAt

For models having `CreatedAt` field, the field will be set to the current time when the record is first created if its value is zero

​	对于具有 `CreatedAt` 字段的模型，如果其值为零，则在记录首次创建时将其字段设置为当前时间。

``` go
db.Create(&user) // 将 `CreatedAt` 设置为当前时间 set `CreatedAt` to current time

user2 := User{Name: "jinzhu", CreatedAt: time.Now()}
db.Create(&user2) // user2 的 `CreatedAt` 不会改变 user2's `CreatedAt` won't be changed

// 要更改其值，可以使用 `Update` To change its value, you could use `Update`
db.Model(&user).Update("CreatedAt", time.Now())
```

You can disable the timestamp tracking by setting `autoCreateTime` tag to `false`, for example:

​	您可以通过将 `autoCreateTime` 标签设置为 `false` 禁用时间戳跟踪，例如：

``` go
type User struct {
  CreatedAt time.Time `gorm:"autoCreateTime:false"`
}
```

### UpdatedAt

For models having `UpdatedAt` field, the field will be set to the current time when the record is updated or created if its value is zero

​	对于具有 `UpdatedAt` 字段的模型，如果其值为零，则在记录更新或创建时将其字段设置为当前时间。

``` go
db.Save(&user) // 将 `UpdatedAt` 设置为当前时间 set `UpdatedAt` to current time

db.Model(&user).Update("name", "jinzhu") // 将 `UpdatedAt` 设置为当前时间 will set `UpdatedAt` to current time

db.Model(&user).UpdateColumn("name", "jinzhu") // `UpdatedAt` 不会改变 `UpdatedAt` won't be changed

user2 := User{Name: "jinzhu", UpdatedAt: time.Now()}
db.Create(&user2) // user2 在创建时不会改变 `UpdatedAt` user2's `UpdatedAt` won't be changed when creating

user3 := User{Name: "jinzhu", UpdatedAt: time.Now()}
db.Save(&user3) // user3 在更新时将 `UpdatedAt` 更改为当前时间 user3's `UpdatedAt` will change to current time when updating
```

You can disable the timestamp tracking by setting `autoUpdateTime` tag to `false`, for example:

​	您可以通过将 `autoUpdateTime` 标签设置为 `false` 禁用时间戳跟踪，例如：

``` go
type User struct {
  UpdatedAt time.Time `gorm:"autoUpdateTime:false"`
}
```

> **NOTE** GORM supports having multiple time tracking fields and track with UNIX (nano/milli) seconds, checkout [Models](https://gorm.io/docs/models.html#time_tracking) for more details
>
> **注意** GORM 支持有多个时间跟踪字段，并使用 UNIX（纳秒/毫秒）秒进行跟踪，请参阅 [Models]({{< ref "/gorm/GettingStarted/declaringModels#创建更新时间追踪-timeunix-millinano-秒-creatingupdating-timeunix-millinano-seconds-tracking">}}) 以获取更多详细信息