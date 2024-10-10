+++
title = "迁移"
date = 2024-02-04T21:13:55+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/migrations/]({{< ref "/buffalo/database/migrations" >}})

# Migrations 迁移 

Software maintenance is a hard task, and you’ll probably need to patch your database to add, modify or remove some fields. The way to handle that with Pop is to use **migrations**.

​	软件维护是一项艰巨的任务，您可能需要修补数据库以添加、修改或删除某些字段。使用 Pop 处理此问题的方法是使用迁移。

You can create new migrations using `fizz`, a custom language describing the database changes in the most database-agnostic way; or use SQL statements if you prefer.

​	您可以使用 `fizz` 创建新迁移，这是一种自定义语言，以最与数据库无关的方式描述数据库更改；或者如果您愿意，可以使用 SQL 语句。

## Writing Migrations 编写迁移 

**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So every time you want to use a command from `soda`, just execute `buffalo pop` instead.
Buffalo 用户须知： `soda` 命令嵌入到 `buffalo` 命令中，位于 `pop` 命名空间之后。因此，每次您想使用 `soda` 中的命令时，只需执行 `buffalo pop` 即可。

### Fizz Migrations Fizz 迁移 

The `soda` command will generate SQL migrations (both the up and down) files for you.

​	 `soda` 命令将为您生成 SQL 迁移（向上和向下）文件。

```bash
$ soda generate fizz name_of_migration
```

Running this command will generate the **empty** following files:

​	运行此命令将生成以下空文件：

```bash
./migrations/20220706213354_name_of_migration.up.fizz
./migrations/20220706213354_name_of_migration.down.fizz
```

The generated files are `fizz` files. Pop uses [Fizz](https://github.com/gobuffalo/fizz/blob/master/README.md) to generate migrations that are both easy to work with and work across multiple types of databases.

​	生成的文件是 `fizz` 文件。Pop 使用 Fizz 生成易于使用且适用于多种类型数据库的迁移。

Further info about this command can be found by using the `--help` flag:

​	有关此命令的更多信息，可以使用 `--help` 标志找到：

```bash
$ soda g migration --help
Generates Up/Down migrations for your database using fizz.

Usage:
  soda generate fizz [name] [flags]

Aliases:
  fizz, migration

Flags:
  -h, --help   help for fizz

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```

By default, the migration will create an UUID `id` that serves as the primary key, as well as `created_at` and `updated_at` datetime columns, so there is no need to create your own. These are the default, but you can override them if you want.
默认情况下，迁移将创建一个用作主键的 UUID `id` ，以及 `created_at` 和 `updated_at` datetime 列，因此无需创建自己的。这些是默认值，但您可以根据需要覆盖它们。

### SQL Migrations SQL 迁移 

If you don’t want to use Fizz, or you have a complicated query you want to execute, you can use SQL.

​	如果您不想使用 Fizz，或者您有想要执行的复杂查询，则可以使用 SQL。

To generate a new **empty** migration, use the following command:

​	要生成一个新的空迁移，请使用以下命令：

```bash
$ soda generate sql name_of_migration
```

Running this command will generate the following files:

​	运行此命令将生成以下文件：

```bash
./migrations/20220706213354_name_of_migration.up.sql
./migrations/20220706213354_name_of_migration.down.sql
```

Further info about this command can be found by using the `--help` flag:

​	有关此命令的更多信息，可以使用 `--help` 标志找到：

```bash
$ soda g sql --help
Generates Up/Down migrations for your database using SQL.

Usage:
  soda generate sql [name] [flags]

Flags:
  -h, --help   help for sql

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```

## Running Migrations 运行迁移 

**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So every time you want to use a command from `soda`, just execute `buffalo pop` instead.
Buffalo 用户须知： `soda` 命令嵌入到 `buffalo` 命令中，位于 `pop` 命名空间之后。因此，每次您想使用 `soda` 中的命令时，只需执行 `buffalo pop` 即可。

### Apply Migrations 应用迁移 

Once migrations have been created they can be run with either of the following commands:

​	创建迁移后，可以使用以下任一命令运行它们：

```bash
$ soda migrate
$ soda migrate up
```

Both commands are identical, one is shorter to type! Migrations will be run in sequential order.

​	这两个命令是相同的，一个更短，更容易输入！迁移将按顺序运行。

### Rollback a Migration 回滚迁移 

If you want to rollback the last applied migration, use the following command:

​	如果您想回滚上次应用的迁移，请使用以下命令：

```bash
$ soda migrate down
```

------

More information about the migration command be found by running:

​	有关迁移命令的更多信息，可以通过运行以下命令找到：

```bash
$ soda migrate --help
Runs migrations against your database.

Usage:
  soda migrate [flags]
  soda migrate [command]

Aliases:
  migrate, m

Available Commands:
  down        Apply one or more of the 'down' migrations.
  status      Displays the status of all migrations.
  up          Apply one or more of the 'up' migrations.

Flags:
  -h, --help   help for migrate

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")

Use "soda migrate [command] --help" for more information about a command.
```

## Targeting a Database 针对数据库 

Since Pop [v4.4.0](https://github.com/gobuffalo/pop/releases/tag/v4.4.0), migrations can target a specific database, using a suffix. This allows to use commands specific to a dialect, only for a given database.

​	自 Pop v4.4.0 起，迁移可以使用后缀针对特定数据库。这允许仅针对给定数据库使用特定于方言的命令。

For instance, if you want to support both PostgreSQL and MySQL, you can create two migrations:

​	例如，如果您想同时支持 PostgreSQL 和 MySQL，则可以创建两个迁移：

- `my-migration.mysql.up.sql` and `my-migration.mysql.down.sql` will be used when migrating a MySQL database.
  `my-migration.mysql.up.sql` 和 `my-migration.mysql.down.sql` 将在迁移 MySQL 数据库时使用。
- `my-migration.postgres.up.sql` and `my-migration.postgres.down.sql` will be used when migrating a PostgreSQL database.
  `my-migration.postgres.up.sql` 和 `my-migration.postgres.down.sql` 将在迁移 PostgreSQL 数据库时使用。

If no version for the dialect can be found, Pop will fallback to the non-suffixed version, if it exists.

​	如果找不到方言的版本，Pop 将回退到非后缀版本（如果存在）。

## Custom Migrations Table 自定义迁移表 

By default, the applied migrations are tracked in the table `schema_migration`. This table is created by pop if it doesn’t exist.

​	默认情况下，已应用的迁移会在表 `schema_migration` 中进行跟踪。如果该表不存在，则由 pop 创建该表。

In some cases, though, you may want to use a different name for this table. Since pop v4.5.0, you can customize the name of this table using the `migration_table_name` option. The example below will use `migrations` as the table name:

​	但在某些情况下，您可能希望为此表使用不同的名称。自 pop v4.5.0 起，您可以使用 `migration_table_name` 选项自定义此表的名称。以下示例将使用 `migrations` 作为表名：

```yaml
development:
  dialect: "postgres"
  url: "your_db_development"
  options:
    migration_table_name: migrations
```

## Migrations Once Deployed 已部署的迁移 

When you build your app, the migrations are stored inside your binary. Your binary has a hidden `migrate` command baked in that performs the migrations, just like it does when you use `buffalo pop migrate`:

​	构建应用时，迁移存储在二进制文件中。您的二进制文件内置了一个隐藏的 `migrate` 命令，用于执行迁移，就像您使用 `buffalo pop migrate` 时一样：

```bash
$ ./myapp migrate
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213171622
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213172104
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213172249
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213173148
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171219070903
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171219071524

0.0010 seconds
```
