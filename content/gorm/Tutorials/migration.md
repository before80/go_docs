+++
title = "迁移"
date = 2023-10-28T14:30:56+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/migration.html](https://gorm.io/docs/migration.html)

## 自动迁移 Auto Migration

Automatically migrate your schema, to keep your schema up to date.

​	自动迁移您的模式，以保持模式的最新状态。

> **NOTE:** AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes. It will change existing column’s type if its size, precision, nullable changed. It **WON’T** delete unused columns to protect your data.
>
> **注意：** AutoMigrate将创建表、缺失的外键、约束、列和索引。如果其大小、精度、可空性发生变化，它将更改现有列的类型。它不会删除未使用的列以保护您的数据。

``` go
db.AutoMigrate(&User{})

db.AutoMigrate(&User{}, &Product{}, &Order{})

// 在创建表时添加表后缀 Add table suffix when creating tables
db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
```

> **NOTE** AutoMigrate creates database foreign key constraints automatically, you can disable this feature during initialization, for example:
>
> **注意** AutoMigrate会自动创建数据库外键约束，您可以在初始化期间禁用此功能，例如：

``` go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  DisableForeignKeyConstraintWhenMigrating: true,
})
```

## Migrator接口 Migrator Interface

GORM provides a migrator interface, which contains unified API interfaces for each database that could be used to build your database-independent migrations, for example:

​	GORM提供了一个migrator接口，该接口包含每个数据库的统一API接口，可用于构建与数据库无关的迁移，例如：

SQLite doesn’t support `ALTER COLUMN`, `DROP COLUMN`, GORM will create a new table as the one you are trying to change, copy all data, drop the old table, rename the new table

​	SQLite不支持`ALTER COLUMN`、`DROP COLUMN`，GORM将在您尝试更改的表中创建一个新表，复制所有数据，删除旧表，重命名新表

MySQL doesn’t support rename column, index for some versions, GORM will perform different SQL based on the MySQL version you are using

​	MySQL在某些版本中不支持重命名列和索引，GORM将根据您使用的MySQL版本执行不同的SQL

``` go
type Migrator interface {
  // AutoMigrate
  AutoMigrate(dst ...interface{}) error

  // Database
  CurrentDatabase() string
  FullDataTypeOf(*schema.Field) clause.Expr

  // Tables
  CreateTable(dst ...interface{}) error
  DropTable(dst ...interface{}) error
  HasTable(dst interface{}) bool
  RenameTable(oldName, newName interface{}) error
  GetTables() (tableList []string, err error)

  // Columns
  AddColumn(dst interface{}, field string) error
  DropColumn(dst interface{}, field string) error
  AlterColumn(dst interface{}, field string) error
  MigrateColumn(dst interface{}, field *schema.Field, columnType ColumnType) error
  HasColumn(dst interface{}, field string) bool
  RenameColumn(dst interface{}, oldName, field string) error
  ColumnTypes(dst interface{}) ([]ColumnType, error)
  
  // Views
  CreateView(name string, option ViewOption) error
  DropView(name string) error

  // Constraints
  CreateConstraint(dst interface{}, name string) error
  DropConstraint(dst interface{}, name string) error
  HasConstraint(dst interface{}, name string) bool

  // Indexes
  CreateIndex(dst interface{}, name string) error
  DropIndex(dst interface{}, name string) error
  HasIndex(dst interface{}, name string) bool
  RenameIndex(dst interface{}, oldName, newName string) error
}
```

### CurrentDatabase

Returns current using database name

​	返回当前使用的数据库名称

``` go
db.Migrator().CurrentDatabase()
```

### Tables

``` go
// 为`User`创建表 Create table for `User`
db.Migrator().CreateTable(&User{})

// 在创建表的SQL中为`User`添加"ENGINE=InnoDB" Append "ENGINE=InnoDB" to the creating table SQL for `User`
db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&User{})

// 检查表`User`是否存在或不存在 Check table for `User` exists or not
db.Migrator().HasTable(&User{})
db.Migrator().HasTable("users")

// 如果存在则删除表（在删除时将忽略或删除外键约束） Drop table if exists (will ignore or delete foreign key constraints when dropping)
db.Migrator().DropTable(&User{})
db.Migrator().DropTable("users")

// 将旧表重命名为新表 Rename old table to new table
db.Migrator().RenameTable(&User{}, &UserInfo{})
db.Migrator().RenameTable("users", "user_infos")
```

### Columns

``` go
type User struct {
  Name string
}

// 添加name字段 Add name field
db.Migrator().AddColumn(&User{}, "Name")
// 删除name字段 Drop name field
db.Migrator().DropColumn(&User{}, "Name")
// 修改name字段 Alter name field
db.Migrator().AlterColumn(&User{}, "Name")
// 检查列是否存在 Check column exists
db.Migrator().HasColumn(&User{}, "Name")

type User struct {
  Name    string
  NewName string
}

// 将列名重命名为新名称 Rename column to new name
db.Migrator().RenameColumn(&User{}, "Name", "NewName")
db.Migrator().RenameColumn(&User{}, "name", "new_name")

// ColumnTypes
db.Migrator().ColumnTypes(&User{}) ([]gorm.ColumnType, error)

type ColumnType interface {
  Name() string
  DatabaseTypeName() string                 // varchar
  ColumnType() (columnType string, ok bool) // varchar(64)
  PrimaryKey() (isPrimaryKey bool, ok bool)
  AutoIncrement() (isAutoIncrement bool, ok bool)
  Length() (length int64, ok bool)
  DecimalSize() (precision int64, scale int64, ok bool)
  Nullable() (nullable bool, ok bool)
  Unique() (unique bool, ok bool)
  ScanType() reflect.Type
  Comment() (value string, ok bool)
  DefaultValue() (value string, ok bool)
}
```

### Views

Create views by `ViewOption`. About `ViewOption`:

​	通过`ViewOption`创建视图。关于`ViewOption`：

- `Query` is a [subquery](https://gorm.io/docs/advanced_query.html#SubQuery), which is required.
- `Query`是一个[子查询]({{< ref "/gorm/CRUDInterface/advancedQuery#子查询-subquery">}})，这是必需的。
- If `Replace` is true, exec `CREATE OR REPLACE` otherwise exec `CREATE`.
- 如果`Replace`为true，则执行`CREATE OR REPLACE`，否则执行`CREATE`。
- If `CheckOption` is not empty, append to sql, e.g. `WITH LOCAL CHECK OPTION`.
- 如果`CheckOption`不为空，则附加到sql，例如`WITH LOCAL CHECK OPTION`。

> **NOTE** SQLite currently does not support `Replace` in `ViewOption`
>
> **注意** SQLite目前不支持`Replace`在`ViewOption`中

``` go
query := db.Model(&User{}).Where("age > ?", 20)

// 创建视图 Create View
db.Migrator().CreateView("users_pets", gorm.ViewOption{Query: query})
// CREATE VIEW `users_view` AS SELECT * FROM `users` WHERE age > 20

// 创建或替换视图 Create or Replace View
db.Migrator().CreateView("users_pets", gorm.ViewOption{Query: query, Replace: true})
// CREATE OR REPLACE VIEW `users_pets` AS SELECT * FROM `users` WHERE age > 20

// 带检查选项的创建视图 Create View With Check Option
db.Migrator().CreateView("users_pets", gorm.ViewOption{Query: query, CheckOption: "WITH CHECK OPTION"})
// CREATE VIEW `users_pets` AS SELECT * FROM `users` WHERE age > 20 WITH CHECK OPTION

// 删除视图 Drop View
db.Migrator().DropView("users_pets")
// DROP VIEW IF EXISTS "users_pets"
```

## 约束 Constraints

``` go
type UserIndex struct {
  Name  string `gorm:"check:name_checker,name <> 'jinzhu'"`
}

// 创建约束 Create constraint
db.Migrator().CreateConstraint(&User{}, "name_checker")

// 删除约束 Drop constraint
db.Migrator().DropConstraint(&User{}, "name_checker")

// 检查约束存在 Check constraint exists
db.Migrator().HasConstraint(&User{}, "name_checker")
```

创建外键约束 Create foreign keys for relations

``` go
type User struct {
  gorm.Model
  CreditCards []CreditCard
}

type CreditCard struct {
  gorm.Model
  Number string
  UserID uint
}

// 为user和credit_cards创建数据库外键 create database foreign key for user & credit_cards
db.Migrator().CreateConstraint(&User{}, "CreditCards")
db.Migrator().CreateConstraint(&User{}, "fk_users_credit_cards")
// ALTER TABLE `credit_cards` ADD CONSTRAINT `fk_users_credit_cards` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)

// 检查user和credit_cards的数据库外键是否存在 check database foreign key for user & credit_cards exists or not
db.Migrator().HasConstraint(&User{}, "CreditCards")
db.Migrator().HasConstraint(&User{}, "fk_users_credit_cards")

// 删除user和credit_cards的数据库外键 drop database foreign key for user & credit_cards
db.Migrator().DropConstraint(&User{}, "CreditCards")
db.Migrator().DropConstraint(&User{}, "fk_users_credit_cards")
```

### 索引 Indexes

``` go
type User struct {
  gorm.Model
  Name string `gorm:"size:255;index:idx_name,unique"`
}

// 为Name字段创建索引 Create index for Name field
db.Migrator().CreateIndex(&User{}, "Name")
db.Migrator().CreateIndex(&User{}, "idx_name")

// 删除Name字段的索引 Drop index for Name field
db.Migrator().DropIndex(&User{}, "Name")
db.Migrator().DropIndex(&User{}, "idx_name")

// 检查索引是否存在 Check Index exists
db.Migrator().HasIndex(&User{}, "Name")
db.Migrator().HasIndex(&User{}, "idx_name")

type User struct {
  gorm.Model
  Name  string `gorm:"size:255;index:idx_name,unique"`
  Name2 string `gorm:"size:255;index:idx_name_2,unique"`
}
// 重命名索引名称 Rename index name
db.Migrator().RenameIndex(&User{}, "Name", "Name2")
db.Migrator().RenameIndex(&User{}, "idx_name", "idx_name_2")
```

## 约束 Constraints

GORM creates constraints when auto migrating or creating table, see [Constraints](https://gorm.io/docs/constraints.html) or [Database Indexes](https://gorm.io/docs/indexes.html) for details

​	GORM在自动迁移或创建表时创建约束，详情请参考[约束]({{< ref "/gorm/AdancedTopics/constraints">}})或[数据库索引]({{< ref "/gorm/AdancedTopics/indexes">}})。

## Atlas集成 Atlas Integration

[Atlas](https://atlasgo.io/) is an open-source database migration tool that has an official integration with GORM.

​	[Atlas](https://atlasgo.io/)是一个开源的数据库迁移工具，它与GORM有一个官方的集成。

While GORM’s `AutoMigrate` feature works in most cases, at some point you many need to switch to a [versioned migrations](https://atlasgo.io/concepts/declarative-vs-versioned#versioned-migrations) strategy.

​	尽管GORM的`AutoMigrate`功能在大多数情况下都能正常工作，但在某些时候，你可能需要切换到一个版本化的迁移策略。

Once this happens, the responsibility for planning migration scripts and making sure they are in line with what GORM expects at runtime is moved to developers.

​	一旦发生这种情况，将规划迁移脚本的责任转移到开发人员是必要的。

Atlas can automatically plan database schema migrations for developers using the official [GORM Provider](https://github.com/ariga/atlas-provider-gorm). After configuring the provider you can automatically plan migrations by running:

​	Atlas可以自动为开发人员计划数据库模式迁移，使用官方的[GORM提供程序](https://github.com/ariga/atlas-provider-gorm)。配置提供程序后，可以通过运行以下命令自动规划迁移：

``` go
atlas migrate diff --env gorm
```

To learn how to use Atlas with GORM, check out the [official documentation](https://atlasgo.io/guides/orms/gorm).

​	要了解如何使用Atlas与GORM一起工作，请查看[官方文档](https://atlasgo.io/guides/orms/gorm)。

## 其他迁移工具 Other Migration Tools

To use GORM with other Go-based migration tools, GORM provides a generic DB interface that might be helpful for you.

​	要使用GORM与其他基于Go的迁移工具一起工作，GORM提供了一个通用的DB接口，这可能会对你有帮助。

``` go
// 返回`*sql.DB` returns `*sql.DB`
db.DB()
```

Refer to [Generic Interface](https://gorm.io/docs/generic_interface.html) for more details.

​	有关更多详细信息，请参阅[通用接口]({{< ref "/gorm/Tutorials/genericDatabaseInterface">}})。