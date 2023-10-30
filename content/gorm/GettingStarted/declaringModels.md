+++
title = "声明模型"
date = 2023-10-28T14:24:20+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/models.html](https://gorm.io/docs/models.html)

Models are normal structs with basic Go types, pointers/alias of them or custom types implementing [Scanner](https://pkg.go.dev/database/sql/?tab=doc#Scanner) and [Valuer](https://pkg.go.dev/database/sql/driver#Valuer) interfaces

​	模型是具有基本Go类型的普通结构体，或者是它们的指针/别名，或者是实现了[Scanner]({{< ref "/stdLib/database/sql#type-scanner">}})和[Valuer]({{< ref "/stdLib/database/sql_driver#type-valuer">}})接口的自定义类型。

For Example:

​	例如：

```go
type User struct {
  ID           uint
  Name         string
  Email        *string
  Age          uint8
  Birthday     *time.Time
  MemberNumber sql.NullString
  ActivatedAt  sql.NullTime
  CreatedAt    time.Time
  UpdatedAt    time.Time
}
```

## 约定 Conventions

GORM prefers convention over configuration. By default, GORM uses `ID` as primary key, pluralizes struct name to `snake_cases` as table name, `snake_case` as column name, and uses `CreatedAt`, `UpdatedAt` to track creating/updating time

​	GORM 倾向于约定优于配置。默认情况下，GORM使用`ID`作为主键，将结构体名称转换为`snake_cases`作为表名，`snake_case`作为列名，并使用`CreatedAt`、`UpdatedAt`来追踪创建/更新时间。

If you follow the conventions adopted by GORM, you’ll need to write very little configuration/code. If convention doesn’t match your requirements, [GORM allows you to configure them](https://gorm.io/docs/conventions.html)

​	如果您遵循GORM采用的约定，您只需要编写很少的配置/代码。如果约定不符合您的要求，[GORM允许您进行配置]({{< ref "/gorm/Tutorials/conventions">}})。

## gorm.Model

GORM defined a `gorm.Model` struct, which includes fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`

​	GORM定义了一个`gorm.Model`结构体，其中包含字段`ID`、`CreatedAt`、`UpdatedAt`、`DeletedAt`。

``` go
// gorm.Model 的定义 gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

You can embed it into your struct to include those fields, refer [Embedded Struct](https://gorm.io/docs/models.html#embedded_struct)

​	您可以将其嵌入到您的结构体中以包含这些字段，请参阅[嵌入式结构体](#嵌入式结构体-embedded-struct)。

## 高级功能 Advanced

### 字段级权限控制 Field-Level Permission

Exported fields have all permissions when doing CRUD with GORM, and GORM allows you to change the field-level permission with tag, so you can make a field to be read-only, write-only, create-only, update-only or ignored

​	导出的字段在进行GORM CRUD操作时具有所有权限，并且GORM允许您使用标签更改字段级权限，因此您可以使一个字段只读、只写、只创建、只更新或被忽略。

> **NOTE** ignored fields won’t be created when using GORM Migrator to create table
>
> **注意**：在使用GORM Migrator创建表时，被忽略的字段将不会被创建。

```go
type User struct {
  Name string `gorm:"<-:create"` // 允许读和创建 allow read and create
  Name string `gorm:"<-:update"` // 允许读和更新 allow read and update
  Name string `gorm:"<-"`        // 允许读和写（创建和更新） allow read and write (create and update)
  Name string `gorm:"<-:false"`  // 允许读，禁止写 allow read, disable write permission
  Name string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写） readonly (disable write permission unless it configured)
  Name string `gorm:"->;<-:create"` // 允许读和写 allow read and create
  Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读） createonly (disabled read from db)
  Name string `gorm:"-"`            // 通过 struct 读写会忽略该字段 ignore this field when write and read with struct
  Name string `gorm:"-:all"`        // 通过 struct 读写、迁移会忽略该字段 ignore this field when write, read and migrate with struct
  Name string `gorm:"-:migration"`  // 通过 struct 迁移会忽略该字段 ignore this field when migrate with struct
}
```

### 创建/更新时间追踪 Time/Unix (Milli/Nano) 秒 Creating/Updating Time/Unix (Milli/Nano) Seconds Tracking

GORM use `CreatedAt`, `UpdatedAt` to track creating/updating time by convention, and GORM will set the [current time]({{< ref "/gorm/AdancedTopics/gormConfig#nowfunc">}}) when creating/updating if the fields are defined

​	GORM使用`CreatedAt`和`UpdatedAt`按照约定来追踪创建/更新时间，如果这些字段被定义了，GORM会在创建/更新时设置当前时间。

To use fields with a different name, you can configure those fields with tag `autoCreateTime`, `autoUpdateTime`

​	如果你想使用不同名称的字段，你可以使用标签`autoCreateTime`和`autoUpdateTime`来配置这些字段。

If you prefer to save UNIX (milli/nano) seconds instead of time, you can simply change the field’s data type from `time.Time` to `int`

​	如果你更喜欢保存UNIX（毫秒/纳秒）而不是时间，你可以直接将字段的数据类型从`time.Time`更改为`int`。

```go
type User struct {
  CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充 Set to current time if it is zero on creating
  UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充 Set to current unix seconds on updating or if it is zero on creating
  Updated   int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳纳秒数填充更新时间 Use unix nano seconds as updating time
  Updated   int64 `gorm:"autoUpdateTime:milli"`// 使用时间戳毫秒数填充更新时间 Use unix milli seconds as updating time
  Created   int64 `gorm:"autoCreateTime"`      // 使用时间戳秒数填充创建时间 Use unix seconds as creating time
}
```

### 嵌入式结构体 Embedded Struct

For anonymous fields, GORM will include its fields into its parent struct, for example:

​	对于匿名字段，GORM会将其字段包含到其父结构体中，例如：

``` go
type User struct {
  gorm.Model
  Name string
}
// 等效于 equals
type User struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
}
```

For a normal struct field, you can embed it with the tag `embedded`, for example:

​	对于普通结构体字段，你可以通过使用`embedded`标签将其嵌入，例如：

```go
type Author struct {
  Name  string
  Email string
}

type Blog struct {
  ID      int
  Author  Author `gorm:"embedded"`
  Upvotes int32
}
// 等效于 equals
type Blog struct {
  ID    int64
  Name  string
  Email string
  Upvotes  int32
}
```

And you can use tag `embeddedPrefix` to add prefix to embedded fields’ db name, for example:

​	你还可以使用`embeddedPrefix`标签为嵌入式字段的数据库名称添加前缀，例如：

```go
type Blog struct {
  ID      int
  Author  Author `gorm:"embedded;embeddedPrefix:author_"`
  Upvotes int32
}
// 等效于 equals
type Blog struct {
  ID          int64
  AuthorName  string
  AuthorEmail string
  Upvotes     int32
}
```

### 字段标签 Fields Tags

Tags are optional to use when declaring models, GORM supports the following tags:

​	声明模型时使用标签是可选的，GORM支持以下标签： 

Tags are case insensitive, however `camelCase` is preferred.

​	标签不区分大小写，但是`camelCase`更受欢迎。

| **标签名** Tag Name    | **描述** Description                                         |
| :--------------------- | :----------------------------------------------------------- |
| column                 | 列的数据库名称 column db name                                |
| type                   | 列的数据类型，推荐使用兼容的通用类型，例如：bool、int、uint、float、string、time、bytes，这些类型适用于所有数据库，并且可以与其他标签一起使用，如`not null`、`size`、`autoIncrement`等指定的数据库数据类型，如`varbinary(8)`。当使用指定的数据库数据类型时，它需要是一个完整的数据库数据类型，例如：`MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT`<br />column data type, prefer to use compatible general type, e.g: bool, int, uint, float, string, time, bytes, which works for all databases, and can be used with other tags together, like `not null`, `size`, `autoIncrement`… specified database data type like `varbinary(8)` also supported, when using specified database data type, it needs to be a full database data type, for example: `MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT` |
| serializer             | 指定如何将数据序列化和反序列化到数据库的序列化器，例如：`serializer:json/gob/unixtime`<br />specifies serializer for how to serialize and deserialize data into db, e.g: `serializer:json/gob/unixtime` |
| size                   | 指定列数据大小/长度，例如：`size:256`<br />specifies column data size/length, e.g: `size:256` |
| primaryKey             | 指定列为主键<br />specifies column as primary key            |
| unique                 | 指定列为唯一<br />specifies column as unique                 |
| default                | 指定列的默认值<br />specifies column default value           |
| precision              | 指定列的精度<br />specifies column precision                 |
| scale                  | 指定列的小数位数<br />specifies column scale                 |
| not null               | 指定列为NOT NULL<br />specifies column as NOT NULL           |
| autoIncrement          | 指定列为可自动递增<br />specifies column auto incrementable  |
| autoIncrementIncrement | 自动递增步长，控制连续列值之间的间隔<br />auto increment step, controls the interval between successive column values |
| embedded               | 嵌入字段<br />embed the field                                |
| embeddedPrefix         | 嵌入式字段的名称前缀<br />column name prefix for embedded fields |
| autoCreateTime         | 创建时追踪当前时间，对于`int`字段，它将追踪unix秒，使用值`nano`/`milli`来追踪unix nano/milli秒，例如：`autoCreateTime:nano`<br />track current time when creating, for `int` fields, it will track unix seconds, use value `nano`/`milli` to track unix nano/milli seconds, e.g: `autoCreateTime:nano` |
| autoUpdateTime         | 创建/更新时追踪当前时间，对于`int`字段，它将追踪unix秒，使用值`nano`/`milli`来追踪unix nano/milli秒，例如：`autoUpdateTime:milli`<br />track current time when creating/updating, for `int` fields, it will track unix seconds, use value `nano`/`milli` to track unix nano/milli seconds, e.g: `autoUpdateTime:milli` |
| index                  | 使用选项创建索引，为多个字段创建复合索引，参考[Indexes]({{< ref "/gorm/AdancedTopics/indexes">}})以获取详细信息<br />create index with options, use same name for multiple fields creates composite indexes, refer [Indexes](https://gorm.io/docs/indexes.html) for details |
| uniqueIndex            | 与`index`相同，但创建唯一索引<br />same as `index`, but create uniqued index |
| check                  | 创建检查约束，例如：`check:age > 13`，参考[Constraints]({{< ref "/gorm/AdancedTopics/constraints">}})<br />creates check constraint, eg: `check:age > 13`, refer [Constraints](https://gorm.io/docs/constraints.html) |
| <-                     | 设置字段的写权限，`<-:create`仅创建字段，`<-:update`仅更新字段，`<-:false`无写权限，`<-`创建和更新权限<br />set field’s write permission, `<-:create` create-only field, `<-:update` update-only field, `<-:false` no write permission, `<-` create and update permission |
| ->                     | 设置字段的读权限，`->:false`无读权限<br />set field’s read permission, `->:false` no read permission |
| -                      | 忽略此字段，`-`无读/写权限，`-:migration`无迁移权限，`-:all`无读/写/迁移权限<br />ignore this field, `-` no read/write permission, `-:migration` no migrate permission, `-:all` no read/write/migrate permission |
| comment                | 在迁移时为字段添加注释<br />add comment for field when migration |

### 关联标签 Associations Tags

GORM allows configure foreign keys, constraints, many2many table through tags for Associations, check out the [Associations section](https://gorm.io/docs/associations.html#tags) for details

​	GORM允许通过关联标签配置外键、约束和多对多表。有关详细信息，请参阅[关联部分]({{< ref "/gorm/Associations">}})。