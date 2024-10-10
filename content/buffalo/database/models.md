+++
title = "模型"
date = 2024-02-04T21:13:19+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/models/]({{< ref "/buffalo/database/models" >}})

# Models 模型 

Pop, as an ORM, allows you to translate database tables into Go structs. This way, you can manipulate Go structs instead of writing SQL statements. The Go code managing this part is named “models”, as a reference to the MVC architecture.

​	Pop 作为 ORM，允许您将数据库表转换为 Go 结构。这样，您可以操作 Go 结构，而无需编写 SQL 语句。管理此部分的 Go 代码被命名为“模型”，作为对 MVC 架构的引用。

In this chapter, you’ll learn how to work with models by hand; and how to improve your workflow using the provided generators.

​	在本章中，您将学习如何手动使用模型；以及如何使用提供的生成器改进您的工作流。

## The Models Directory 模型目录 

Pop model files are stored in the `models` directory, at your project root (see [the Directory Structure]({{< ref "/buffalo/gettingStarted/directoryStructure" >}}) chapter for more info about the Buffalo way to organize your files).

​	Pop 模型文件存储在 `models` 目录中，位于您的项目根目录（有关组织文件的 Buffalo 方式的更多信息，请参阅目录结构章节）。

This directory contains:

​	此目录包含：

- A `models.go` file, which defines the common parts for every defined model. It also contains a pointer to the configured connection. Remember the code is your own, so you can place whatever you like here.
  `models.go` 文件，它定义了每个已定义模型的公共部分。它还包含指向已配置连接的指针。请记住，代码是您自己的，因此您可以在这里放置您喜欢的任何内容。
- Model definition files, one for each model (so one per database table you want to access this way).
  模型定义文件，每个模型一个（因此每个数据库表一个，您希望通过这种方式访问）。

## Define a Simple Model 定义简单模型 

A model file defines a mapping for the database table, validation methods and Pop callbacks if you want to add more model-related logic.

​	模型文件定义数据库表的映射、验证方法和 Pop 回调（如果您想添加更多与模型相关的逻辑）。

Let’s take the following SQL table definition, and write a matching structure:

​	让我们采用以下 SQL 表定义，并编写一个匹配的结构：

```sql
CREATE TABLE sodas (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    label character varying(255)
);

ALTER TABLE sodas ADD CONSTRAINT sodas_pkey PRIMARY KEY (id);
```

We’ll start by creating a new file in the `models` directory, called `soda.go` (the convention used here is to take the singular form of the word). In this file, we’ll create the structure for the `sodas` table (the structure is singular too, since it will contain a single line of the table):

​	我们将首先在 `models` 目录中创建一个新文件，名为 `soda.go` （此处使用的约定是采用单词的单数形式）。在此文件中，我们将创建 `sodas` 表的结构（结构也是单数，因为它将包含表的一行）：

```go
package models

import (
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
)

type Soda struct {
	ID                   uuid.UUID    `db:"id"`
	CreatedAt            time.Time    `db:"created_at"`
	UpdatedAt            time.Time    `db:"updated_at"`
	Label                nulls.String `db:"label"`
}
```

That’s it! You don’t need anything else to work with Pop! Note, for each table field, we defined a `pop` tag matching the field name, but it’s not required. If you don’t provide a name, Pop will use the name of the struct field to generate one.

​	就是这样！您无需任何其他内容即可使用 Pop！请注意，对于每个表字段，我们定义了一个与字段名称匹配的 `pop` 标记，但这不是必需的。如果您不提供名称，Pop 将使用结构字段的名称生成一个名称。

## Using the Generator 使用生成器 

**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So everytime you want to use a command from `soda`, just execute `buffalo pop` instead.
Buffalo 用户请注意： `soda` 命令嵌入到 `buffalo` 命令中，位于 `pop` 命名空间之后。因此，每次您想使用 `soda` 中的命令时，只需执行 `buffalo pop` 即可。

Writing the files by hand is not the most efficient way to work. Soda (and Buffalo, if you followed the chapter about Soda) provides a generator to help you:

​	手动编写文件并不是最有效的工作方式。Soda（如果您遵循了有关 Soda 的章节，则还包括 Buffalo）提供了一个生成器来帮助您：

```bash
$ soda g model --help
Generates a model for your database

Usage:
  soda generate model [name] [flags]

Aliases:
  model, m


Flags:
 -h, --help                    help for model
      --migration-type string   sets the type of migration files for model (sql or fizz) (default "fizz")
      --models-path string      the path the model will be created in (default "models")
  -s, --skip-migration   Skip creating a new fizz migration for this model.
      --struct-tag string       sets the struct tags for model (xml/json/jsonapi) (default "json")

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```

You can remove generated model by running:

​	您可以通过运行以下命令来移除生成的模型：

```bash
$ soda destroy model [name]
```

Or in short form:

​	或简短形式：

```bash
$ soda d m [name]
```

## Nulls Handling Null 值处理 

If you need to store `NULL` values in your table, you’ll have to use special types: for instance, you can’t store a `NULL` value if your type is `int`.

​	如果需要在表中存储 `NULL` 值，则必须使用特殊类型：例如，如果类型是 `int` ，则无法存储 `NULL` 值。

The [Go standard library](https://golang.org/pkg/database/sql) provides special types for that use case, like [`sql.NullBool`](https://golang.org/pkg/database/sql/#NullBool) or [`sql.NullInt64`](https://golang.org/pkg/database/sql/#NullInt64).

​	Go 标准库为该用例提供了特殊类型，例如 `sql.NullBool` 或 `sql.NullInt64` 。

If you need more than what the standard library offers, you can use the [gobuffalo/nulls](https://github.com/gobuffalo/nulls) package which provides more nulls types and a better handling for JSON serialization and unserialization.

​	如果需要超出标准库提供的功能，可以使用 gobuffalo/nulls 包，该包提供了更多 null 类型，并更好地处理 JSON 序列化和反序列化。

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password nulls.String
}
```

## Customize Models 自定义模型 

### Mapping Model Fields 映射模型字段 

By default when trying to map a struct to a database table, Pop, will use the name of the field in the struct as the name of the column in the database.

​	默认情况下，当尝试将结构映射到数据库表时，Pop 会使用结构中的字段名称作为数据库中列的名称。

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password string
}
```

With the above struct it is assumed the column names in the database are `ID`, `Email`, and `Password`.

​	对于上述结构，假定数据库中的列名为 `ID` 、 `Email` 和 `Password` 。

These column names can be changed by using the `db` struct tag.

​	可以使用 `db` 结构标签来更改这些列名。

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}
```

Now the columns names are expected to be `id`, `email`, and `password`.

​	现在，列名预计为 `id` 、 `email` 和 `password` 。

This is very similar to how [form binding]({{< ref "/buffalo/requestHandling/requestBinding" >}}) works.

​	这与表单绑定工作方式非常相似。

Any types can be used that adhere to the [Scanner](https://golang.org/pkg/database/sql/#Scanner) and [Valuer](https://golang.org/pkg/database/sql/driver/#Valuer) interfaces, however, so that you don’t have to write these yourself it is recommended you stick with the following types:

​	任何遵守 Scanner 和 Valuer 接口的类型都可以使用，但是，为了避免您自己编写这些类型，建议您坚持使用以下类型：

| Base type 基本类型     | Nullable 可为 null | Slice/Array 切片/数组 |
| :--------------------- | :----------------- | :-------------------- |
| int                    | nulls.Int          | slices.Int            |
| int32                  | nulls.Int32        | ——                    |
| int64                  | nulls.Int64        | ——                    |
| uint32                 | nulls.UInt32       | ——                    |
| float32                | nulls.Float32      | ——                    |
| float, float64         | nulls.Float64      | slices.Float          |
| bool                   | nulls.Bool         | ——                    |
| []byte                 | nulls.ByteSlice    | ——                    |
| string                 | nulls.String       | slices.String         |
| uuid.UUID              | nulls.UUID         | slices.UUID           |
| time.Time              | nulls.Time         | ——                    |
| map[string]interface{} | ———                | slices.Map            |

**Note**: Any `slices.Map` typed fields will need to be initialized before `Bind`ing or accessing.

​	注意：任何 `slices.Map` 类型字段在 `Bind` 或访问之前都需要初始化。

```go
widget := &models.Widget{Data: slices.Map{}}
```

### Read Only Fields 只读字段 

It is often necessary to read a field from a database, but not want to write that field to the database. This can be done using the `rw` struct tag.

​	通常需要从数据库中读取字段，但不想将该字段写入数据库。可以使用 `rw` 结构标签来完成此操作。

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"r"`
}
```

In this example all fields will be read **from** the database and all fields, **except** for `Password` will be able to write to the database.

​	在此示例中，所有字段都将从数据库中读取，并且除 `Password` 之外的所有字段都能够写入数据库。

### Write Only Fields 只写字段 

Write only fields are the reverse of read only fields. These are fields that you want to write to the database, but never retrieve. Again, this makes use of the `rw` struct tag.

​	只写字段与只读字段相反。这些是您想要写入数据库但从不检索的字段。同样，这也利用了 `rw` 结构标签。

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"w"`
}
```

### Skipping Model Fields 跳过模型字段 

Sometimes you need to let Pop know that certain field should not be stored in the database table. Perhaps it’s just a field you use in-memory or other logical reason related with the application you’re building.

​	有时您需要让 Pop 知道某些字段不应存储在数据库表中。也许它只是您在内存中使用的字段或与您正在构建的应用程序相关的其他逻辑原因。

The way you let Pop know about this is by using the `db` struct tag on your model and setting it to be `-` like the following example:

​	您可以通过在模型上使用 `db` 结构标签并将其设置为 `-` 来让 Pop 知道这一点，如下例所示：

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"-"`
}
```

As you may see the `Password` field is marked as `db:"-"` that means Pop will neither ***store\*** nor ***retrieve\*** this field from the database.

​	正如您所见， `Password` 字段被标记为 `db:"-"` ，这意味着 Pop 既不会存储也不会从数据库中检索此字段。

### Changing the Select Clause for a Column 更改列的选择子 

The default, when trying to build the `select` query for a struct is to use all of the field names to build a query.

​	尝试为结构构建 `select` 查询时的默认行为是使用所有字段名称来构建查询。

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}
```

The resulting `select` statement would look like this:

​	生成的 `select` 语句将如下所示：

```sql
select id, email, password from users
```

We can change the statement for a column using the `select` tag.

​	我们可以使用 `select` 标记更改列的语句。

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" select:"password as p"`
}
```

The resulting `select` statement would look like this:

​	生成的 `select` 语句将如下所示：

```sql
select id, email, password as p from users
```

### Using a Custom Table Name 使用自定义表名 

Sometimes, you’ll have to work with an existing schema, with the table names non-matching the Pop conventions. You can override this behavior, and provide a custom table name by implementing the [`TableNameAble`](https://godoc.org/github.com/gobuffalo/pop#TableNameAble) interface:

​	有时，您必须使用现有架构，其中表名与 Pop 约定不匹配。您可以覆盖此行为，并通过实现 `TableNameAble` 接口提供自定义表名：

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}

// TableName overrides the table name used by Pop.
func (u User) TableName() string {
  return "my_users"
}
```

It is recommended to use a value receiver over a pointer receiver if the struct is used as a value anywhere in the code.

​	如果结构在代码中的任何位置用作值，则建议使用值接收器而不是指针接收器。

```go
// recommended:
func (u User) TableName() string {

// can cause issues:
func (u *User) TableName() string {
```

### UNIX Timestamps UNIX 时间戳 

Since **v4.7.0**
自 v4.7.0 起



If you define the `CreatedAt` and `UpdatedAt` fields in your model struct (and they are created by default when you use the model generator), Pop will manage them for you. It means when you create a new entity in the database, the `CreatedAt` field will be set to the current datetime, and `UpdatedAt` will be set each time you update an existing entity.

​	如果您在模型结构中定义 `CreatedAt` 和 `UpdatedAt` 字段（当您使用模型生成器时，它们默认创建），Pop 将为您管理它们。这意味着当您在数据库中创建新实体时， `CreatedAt` 字段将设置为当前日期时间， `UpdatedAt` 将在每次更新现有实体时设置。

These fields are defined as time.Time, but now you can define them as `int` and handle them as UNIX timestamps.

​	这些字段定义为 time.Time，但现在您可以将它们定义为 `int` 并将它们作为 UNIX 时间戳处理。

```go
type User struct {
  ID        int    `db:"id"`
  CreatedAt int    `db:"created_at"`
  UpdatedAt int    `db:"updated_at"`
  FirstName string `db:"first_name"`
  LastName  string `db:"last_name"`
}
```

If you use fizz migrations, make sure to define these fields by yourself, and disable the default datetime timestamps:

​	如果您使用 fizz 迁移，请务必自己定义这些字段，并禁用默认日期时间戳：

```go
create_table("users") {
  t.Column("id", "int", {primary: true})
  t.Column("created_at", "int")
  t.Column("updated_at", "int")
  t.Column("first_name", "string")
  t.Column("last_name", "string")
  t.DisableTimestamps()
}
```

## Views Models 视图模型 

A [view](https://en.wikipedia.org/wiki/View_(SQL)) is a database collection object which stores the result of a query. Since this object acts as a read-only table, you can map it with Pop models just like a table.

​	视图是存储查询结果的数据库集合对象。由于此对象充当只读表，因此您可以像表一样将其与 Pop 模型进行映射。

If you want to use a model with more than one table, defining a view is probably the best solution for you.

​	如果您想使用具有多个表的模型，定义视图可能是您的最佳解决方案。

### Example 示例 

The following example uses the PostgreSQL syntax. We’ll start by creating two tables:

​	以下示例使用 PostgreSQL 语法。我们将首先创建两个表：

```sql
-- Create a sodas table
CREATE TABLE sodas (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    provider_id uuid NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE sodas ADD CONSTRAINT sodas_pkey PRIMARY KEY (id);

-- Create a providers table
CREATE TABLE providers (
    id uuid NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE providers ADD CONSTRAINT providers_pkey PRIMARY KEY (id);

-- Create a foreign key between the two tables
ALTER TABLE sodas ADD FOREIGN KEY (provider_id) REFERENCES providers(id);
```

Then create a view from the two tables:

​	然后从这两个表中创建一个视图：

```sql
CREATE VIEW sodas_with_providers AS
SELECT s.id, s.created_at, s.updated_at, p.label AS provider_label, s.label
FROM sodas s
LEFT JOIN providers p ON p.id = s.provider_id;
```

Since the view is considered as a table by Pop, let’s finish by declaring a new model:

​	由于 Pop 将视图视为表，因此我们最后声明一个新模型：

```go
type SodasWithProvider struct {
	ID            uuid.UUID `db:"id" rw:"r"`
	CreatedAt     time.Time `db:"created_at" rw:"r"`
	UpdatedAt     time.Time `db:"updated_at" rw:"r"`
	Label         string    `db:"label" rw:"r"`
	ProviderLabel string    `db:"provider_label" rw:"r"`
}
```

As we learned in this chapter, each attribute on the structure has a read-only tag `rw:"r"`. Since a view is a read-only object, it prevents any writing operation before hitting the database.

​	正如我们在本章中学到的，结构上的每个属性都有一个只读标记 `rw:"r"` 。由于视图是一个只读对象，因此它可以防止在访问数据库之前进行任何写入操作。

## Related Content 相关内容 

- [Migrations]({{< ref "/buffalo/database/migrations" >}}) - Write database migrations.
  迁移 - 编写数据库迁移。
- [Querying]({{< ref "/buffalo/database/querying" >}}) - Query data from your database.
  查询 - 从数据库中查询数据。
