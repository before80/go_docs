+++
title = "模型定义"
date = 2024-02-04T10:01:55+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/model/models/]({{< ref "/beego/mvcIntroduction/models/modelDefinition" >}})

# Model Definition 模型定义



## Model Definition 模型定义

Model names are used for database data conversion and [Database Schema Generation](https://beego.wiki/docs/mvc/model/models/cmd.md#database-schema-generation)

​	模型名称用于数据库数据转换和数据库模式生成

## Naming conventions 命名约定

Table name conversion consists in translating camel case used for model names to snake case for table names as follows:

​	表名转换包括将用于模型名称的驼峰式大小写转换为表名的蛇形大小写，如下所示：

```
AuthUser -> auth_user
Auth_User -> auth__user
DB_AuthUser -> d_b__auth_user
```

In other words, all is converted to lower case and `_` is the separator. Every uppercase add a separator before it, except the first one.

​	换句话说，所有内容都转换为小写， `_` 是分隔符。每个大写字母在其前面添加一个分隔符，第一个除外。

## Custom table name 自定义表名

Using `TableNameI` interface:

​	使用 `TableNameI` 接口：

```go
type User struct {
	Id int
	Name string
}

func (u *User) TableName() string {
	return "auth_user"
}
```

If you set [prefix](https://beego.wiki/docs/mvc/model/models/orm.md#registermodelwithprefix) to `prefix_`, the table name will be `prefix_auth_user`.

​	如果将前缀设置为 `prefix_` ，则表名将为 `prefix_auth_user` 。

## Custom index 自定义索引

Using `TableIndexI` interface:

​	使用 `TableIndexI` 接口：

Add index to one or more fields:

​	为一个或多个字段添加索引：

```go
type User struct {
	Id    int
	Name  string
	Email string
}

// multiple fields index
func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "Name"},
	}
}

// multiple fields unique key
func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"Name", "Email"},
	}
}
```

## Custom engine 自定义引擎

Only supports MySQL database

​	仅支持 MySQL 数据库

The default engine is the default engine of the current database engine of your mysql settings.

​	默认引擎是您 mysql 设置的当前数据库引擎的默认引擎。

Using `TableEngineI` interface:

​	使用 `TableEngineI` 接口：

```go
type User struct {
	Id    int
	Name  string
	Email string
}

// Set engine to INNODB
func (u *User) TableEngine() string {
	return "INNODB"
}
```

## Set parameters 设置参数

```go
orm:"null;rel(fk)"
```

Use `;` as the separator of multiple settings. Use `,` as the separator if a setting has multiple values.

​	使用 `;` 作为多个设置的分隔符。如果某个设置具有多个值，请使用 `,` 作为分隔符。

#### Ignore field 忽略字段

Use `-` to ignore field in the struct.

​	使用 `-` 忽略结构中的字段。

```go
type User struct {
...
	AnyField string `orm:"-"`
...
}
```

#### auto 自动

When Field type is int, int32, int64, uint, uint32 or uint64, you can set it as auto increment.

​	当字段类型为 int、int32、int64、uint、uint32 或 uint64 时，您可以将其设置为自增。

- If there is no primary key in the model definition, the field `Id` with one of the types above will be considered as auto increment key
  如果模型定义中没有主键，则类型为上述类型之一的字段 `Id` 将被视为自增键

#### pk

Set as primary key. Used for using other type field as primary key.

​	设置为主键。用于使用其他类型字段作为主键。

#### null

Fields are `NOT NULL` by default. Set null to `ALLOW NULL`.

​	字段默认值为 `NOT NULL` 。将 null 设置为 `ALLOW NULL` 。

```go
Name string `orm:"null"`
```

#### index 索引

Add index for one field

​	为一个字段添加索引

#### unique

Add unique key for one field

​	为一个字段添加唯一键

```go
Name string `orm:"unique"`
```

#### column

Set column name in db table for field.

​	为字段设置数据库表中的列名。

```go
Name string `orm:"column(user_name)"`
```

#### size 大小

Default value for string field is varchar(255).

​	字符串字段的默认值为 varchar(255)。

It will use varchar(size) after setting.

​	设置后将使用 varchar(size)。

```go
Title string `orm:"size(60)"`
```

#### digits / decimals 数字/小数

Set precision for float32 or float64.

​	为 float32 或 float64 设置精度。

```go
Money float64 `orm:"digits(12);decimals(4)"`
```

Total 12 digits, 4 digits after point. For example: `12345678.1234`

​	总共 12 位数字，小数点后 4 位数字。例如： `12345678.1234`

#### auto_now / auto_now_add

```go
Created time.Time `orm:"auto_now_add;type(datetime)"`
Updated time.Time `orm:"auto_now;type(datetime)"`
```

- auto_now: every save will update time.
  auto_now：每次保存都会更新时间。
- auto_now_add: set time at the first save
  auto_now_add：在首次保存时设置时间

This setting won’t affect massive `update`.

​	此设置不会影响大规模 `update` 。

#### type 类型

If set type as date, the field’s db type is date.

​	如果将类型设置为日期，则字段的数据库类型为日期。

```go
Created time.Time `orm:"auto_now_add;type(date)"`
```

If set type as datetime, the field’s db type is datetime.

​	如果将类型设置为日期时间，则字段的数据库类型为日期时间。

```go
Created time.Time `orm:"auto_now_add;type(datetime)"`
```

#### Time Precision 时间精度

```go
type User struct {
...
Created time.Time `orm:"type(datetime);precision(4)"`
...
}
```

#### default value 默认值

you could use it like:

​	您可以像这样使用它：

```go
import (
"github.com/beego/beego/v2/client/orm/filter/bean"
"github.com/beego/beego/v2/client/orm"
)

type DefaultValueTestEntity struct {
Id            int
Age           int `default:"12"`
AgeInOldStyle int `orm:"default(13);bee()"`
AgeIgnore     int
}

func XXX() {
    builder := bean.NewDefaultValueFilterChainBuilder(nil, true, true)
    orm.AddGlobalFilterChain(builder.FilterChain)
    o := orm.NewOrm()
    _, _ = o.Insert(&User{
        ID: 1,
        Name: "Tom",
    })
}
NewDefaultValueFilterChainBuilder`will create an instance of `DefaultValueFilterChainBuilder` In beego v1.x, the default value config looks like `orm:default(xxxx)` But the default value in 2.x is `default:xxx`, so if you want to be compatible with v1.x, please pass true as `compatibleWithOldStyle
```

​	 `NewDefaultValueFilterChainBuilder` 将创建一个 `DefaultValueFilterChainBuilder` 的实例。在 beego v1.x 中，默认值配置如下所示 `orm:default(xxxx)` 。但在 2.x 中，默认值为 `default:xxx` ，因此如果您想与 v1.x 兼容，请将 true 作为 `compatibleWithOldStyle` 传递。

#### Comment 注释

Set comment value for field.

​	为字段设置注释值。

```go
type User struct {
	...
	Status int `orm:"default(1);description(this is status)"`
	...
}
```

## Relationships 关系

#### One to one 一对一

**RelOneToOne**:

```go
type User struct {
	...
	Profile *Profile `orm:"null;rel(one);on_delete(set_null)"`
	...
}
```

The reverse relationship **RelReverseOne**:

​	反向关系 RelReverseOne:

```go
type Profile struct {
	...
	User *User `orm:"reverse(one)"`
	...
}
```

#### One to many 一对多

**RelForeignKey**:

```go
type Post struct {
	...
	User *User `orm:"rel(fk)"` // RelForeignKey relation
	...
}
```

The reverse relationship **RelReverseMany**:

​	反向关系 RelReverseMany:

```go
type User struct {
	...
	Posts []*Post `orm:"reverse(many)"` // reverse relationship of fk
	...
}
```

#### Many to many 多对多

**RelManyToMany**:

```go
type Post struct {
	...
	Tags []*Tag `orm:"rel(m2m)"` // ManyToMany relation
	...
}
```

The reverse relationship **RelReverseMany**:

​	反向关系 RelReverseMany:

```go
type Tag struct {
	...
	Posts []*Post `orm:"reverse(many)"`
	...
}
```

In this example, by default the auto-generated table name is: `post_tag`. The name of the struct in which we have `orm:"rel(m2m)"` defines the first half part, the name of the struct in which we have `orm:"reverse(many)"` defines the other half. It respects the naming conversion convention we have seen in [Naming conventions](https://beego.wiki/docs/mvc/model/models/#naming-conventions)

​	在此示例中，默认情况下自动生成的表名为： `post_tag` 。我们拥有 `orm:"rel(m2m)"` 的结构的名称定义了前半部分，我们拥有 `orm:"reverse(many)"` 的结构的名称定义了后半部分。它遵循我们在命名约定中看到的命名转换约定

##### rel_table / rel_through

This setting is for `orm:"rel(m2m)"` field:

​	此设置适用于 `orm:"rel(m2m)"` 字段：

```
rel_table       Set the auto-generated m2m connecting table name
rel_through     If you want to use custom m2m connecting table, set name by using this setting.
              Format: `project_path/current_package.ModelName`
              For example: `app/models.PostTagRel` PostTagRel table needs to have a relationship to Post table and Tag table.
```

If rel_table is set, rel_through is ignored.

​	如果设置了 rel_table，则忽略 rel_through。

You can set these as follows:

​	您可以按如下方式设置这些内容：

```
orm:"rel(m2m);rel_table(the_table_name)"
orm:"rel(m2m);rel_through(project_path/current_package.ModelName)"
```

#### on_delete

Set how to deal with field if related relationship is deleted:

​	设置在删除相关关系时如何处理字段：

```
cascade        cascade delete (default)
set_null       set to NULL. Need to set null = true
set_default    set to default value. Need to set default value.
do_nothing     do nothing. ignore.
type User struct {
	...
	Profile *Profile `orm:"null;rel(one);on_delete(set_null)"`
	...
}
type Profile struct {
	...
	User *User `orm:"reverse(one)"`
	...
}

// Set User.Profile to NULL while deleting Profile
```

#### Examples of on_delete on_delete 示例

```go
type User struct {
    Id int
    Name string
}

type Post struct {
    Id int
    Title string
    User *User `orm:"rel(fk)"`
}
```

Assume Post -> User is ManyToOne relationship by foreign key.

​	假设 Post -> User 是通过外键建立的 ManyToOne 关系。

```
o.Filter("Id", 1).Delete()
```

This will delete User with Id 1 and all his Posts.

​	这将删除 Id 为 1 的 User 及其所有 Post。

If you don’t want to delete the Posts, you need to set `set_null`

​	如果您不想删除 Post，则需要设置 `set_null`

```go
type Post struct {
    Id int
    Title string
    User *User `orm:"rel(fk);null;on_delete(set_null)"`
}
```

In this case, only set related Post.user_id to NULL while deleting.

​	在这种情况下，仅在删除时将相关的 Post.user_id 设置为 NULL。

Usually for performance purposes, it doesn’t matter to have redundant data. The massive deletion is the real problem

​	通常出于性能目的，拥有冗余数据无关紧要。真正的问题是大规模删除

```go
type Post struct {
    Id int
    Title string
    User *User `orm:"rel(fk);null;on_delete(do_nothing)"`
}
```

So just don’t change Post (ignore it) while deleting User.

​	因此，在删除用户时，不要更改帖子（忽略它）。

## Model fields mapping with database type 模型字段与数据库类型的映射

Here is the recommended database type mapping. It’s also the standard for table generation.

​	以下是推荐的数据库类型映射。它也是表生成的标准。

All the fields are **NOT NULL** by default.

​	默认情况下，所有字段均为非空。

#### MySQL

| go                                                           | mysql                           |
| :----------------------------------------------------------- | :------------------------------ |
| int, int32 - set as auto or name is `Id` int、int32 - 设置为自动或名称为 `Id` | integer AUTO_INCREMENT          |
| int64 - set as auto or name is`Id` int64 - 设置为自动或名称为 `Id` | bigint AUTO_INCREMENT           |
| uint, uint32 - set as auto or name is `Id` uint、uint32 - 设置为 auto 或名称为 `Id` | integer unsigned AUTO_INCREMENT |
| uint64 - set as auto or name is `Id` uint64 - 设置为 auto 或名称为 `Id` | bigint unsigned AUTO_INCREMENT  |
| bool                                                         | bool                            |
| string - default size 255 string - 默认大小 255              | varchar(size)                   |
| string - set type(char) string - 设置类型 (char)             | char(size)                      |
| string - set type(text) string - 设置类型 (text)             | longtext                        |
| time.Time - set type as date time.Time - 设置类型为日期      | date 日期                       |
| time.Time                                                    | datetime 日期时间               |
| byte 字节                                                    | tinyint unsigned 无符号小整数   |
| rune 符文                                                    | integer 整数                    |
| int                                                          | integer 整数                    |
| int8                                                         | tinyint 小整数                  |
| int16                                                        | smallint 短整数                 |
| int32                                                        | integer 整数                    |
| int64                                                        | bigint 大整数                   |
| uint                                                         | integer unsigned 无符号整数     |
| uint8                                                        | tinyint unsigned 无符号小整数   |
| uint16                                                       | smallint unsigned 无符号短整数  |
| uint32                                                       | integer unsigned 无符号整数     |
| uint64                                                       | bigint unsigned 无符号大整数    |
| float32 32 位浮点数                                          | double precision 双精度         |
| float64 64 位浮点数                                          | double precision 双精度         |
| float64 - set digits and decimals float64 - 设置数字和小数   | numeric(digits, decimals)       |

#### Sqlite3

| go                                                           | sqlite3                       |
| :----------------------------------------------------------- | :---------------------------- |
| int, int32, int64, uint, uint32, uint64 - set as auto or name is `Id` int, int32, int64, uint, uint32, uint64 - 设置为自动或名称为 `Id` | integer AUTOINCREMENT         |
| bool                                                         | bool                          |
| string - default size 255 字符串 - 默认大小 255              | varchar(size)                 |
| string - set type(char) 字符串 - 设置类型(char)              | character(size) 字符(大小)    |
| string - set type(text) 字符串 - 设置类型(文本)              | text 文本                     |
| time.Time - set type as date 时间.时间 - 设置类型为日期      | date 日期                     |
| time.Time                                                    | datetime 日期时间             |
| byte 字节                                                    | tinyint unsigned 无符号小整数 |
| rune 符文                                                    | integer 整数                  |
| int                                                          | integer 整数                  |
| int8                                                         | tinyint                       |
| int16                                                        | smallint                      |
| int32                                                        | integer                       |
| int64                                                        | bigint                        |
| uint                                                         | integer unsigned              |
| uint8                                                        | tinyint unsigned              |
| uint16                                                       | smallint unsigned             |
| uint32                                                       | integer unsigned              |
| uint64                                                       | bigint unsigned               |
| float32                                                      | real 实数                     |
| float64                                                      | real 实数                     |
| float64 - set digits and decimals float64 - 设置位数与小数位数 | decimal 小数                  |

#### PostgreSQL

| go                                                           | postgres                                             |
| :----------------------------------------------------------- | :--------------------------------------------------- |
| int, int32, int64, uint, uint32, uint64 - set as auto or name is `Id` int、int32、int64、uint、uint32、uint64 - 设置为自动或为 `Id` | serial 序列                                          |
| bool                                                         | bool                                                 |
| string - if not set size default text 字符串 - 如果未设置为大小，则默认文本 | varchar(size)                                        |
| string - set type(char)                                      | char(size)                                           |
| string - set type(text)                                      | text                                                 |
| string - set type(json)                                      | json                                                 |
| string - set type(jsonb)                                     | jsonb                                                |
| time.Time - set type as date                                 | date 日期                                            |
| time.Time                                                    | timestamp with time zone 带时区的timestamp           |
| byte 字节                                                    | smallint CHECK(“column” >= 0 AND “column” <= 255)    |
| rune                                                         | integer                                              |
| int                                                          | integer                                              |
| int8                                                         | smallint CHECK(“column” >= -127 AND “column” <= 128) |
| int16                                                        | smallint                                             |
| int32                                                        | integer                                              |
| int64                                                        | bigint                                               |
| uint                                                         | bigint CHECK(“column” >= 0)                          |
| uint8                                                        | smallint CHECK(“column” >= 0 AND “column” <= 255)    |
| uint16                                                       | integer CHECK(“column” >= 0)                         |
| uint32                                                       | bigint CHECK(“column” >= 0)                          |
| uint64                                                       | bigint CHECK(“column” >= 0)                          |
| float32                                                      | double precision                                     |
| float64                                                      | double precision                                     |
| float64 - set digits and decimals float64 - 设置数字和小数   | numeric(digits, decimals)                            |

## Relational fields 关系字段

It’s field type depends on related primary key.
它的字段类型取决于相关的主键。

- RelForeignKey
- RelOneToOne
- RelManyToMany
- RelReverseOne
- RelReverseMany
