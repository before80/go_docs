+++
title = "Fizz"
date = 2024-02-04T21:14:06+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/fizz/](https://gobuffalo.io/documentation/database/fizz/)

# Fizz

Fizz is a common DSL for migrating databases. It tries to be as database-agnostic as possible. This is the default language used by Pop to define [database migrations](https://gobuffalo.io/documentation/database/migrations).

​	Fizz 是一种用于迁移数据库的通用 DSL。它尝试尽可能地与数据库无关。这是 Pop 用于定义数据库迁移的默认语言。

## Create a Table 创建表 

```javascript
create_table("users") {
  t.Column("email", "string", {})
  t.Column("twitter_handle", "string", {"size": 50})
  t.Column("age", "integer", {"default": 0})
  t.Column("admin", "bool", {"default": false})
  t.Column("company_id", "uuid", {"default_raw": "uuid_generate_v1()"})
  t.Column("bio", "text", {"null": true})
  t.Column("joined_at", "timestamp", {})
}

create_table("todos") {
  t.Column("user_id", "integer", {})
  t.Column("title", "string", {"size": 100})
  t.Column("details", "text", {"null": true})
  t.ForeignKey("user_id", {"users": ["id"]}, {"on_delete": "cascade"})
}
```

The `id` column doesn’t have to be an integer. For instance, you can use an [`UUID`](https://github.com/gofrs/uuid) type instead:

​	 `id` 列不必是整数。例如，您可以使用 `UUID` 类型：

```javascript
create_table("users") {
  t.Column("id", "uuid", {primary: true})
  // ...
}
```

By default, fizz will generate two `timestamp` columns: `created_at` and `updated_at`.

​	默认情况下，fizz 将生成两列 `timestamp` ： `created_at` 和 `updated_at` 。

The `t.Column` method takes the following arguments: name of the column, the type of the field, and finally the last argument is any options you want to set on that column.

​	方法采用以下参数：列的名称、字段的类型，最后，最后一个参数是您想要在该列上设置的任何选项。

#### “Common” Types: “通用”类型：

- `string`
- `text`
- `timestamp`, `time`, `datetime`
- `integer`
- `bool`
- `uuid`

Any other type passed will be passed straight through to the underlying database.

​	传递的任何其他类型都将直接传递到基础数据库。

For example for PostgreSQL you could pass `jsonb` and it will be supported, however, SQLite will yell very loudly at you if you do the same thing!

​	例如，对于 PostgreSQL，您可以传递 `jsonb` ，它将得到支持，但是，如果您对 SQLite 执行相同的操作，它会非常大声地向您发出警告！

#### Supported Options: 支持的选项：

| Option 选项          | Description 说明                                             | Example 示例                                               |
| :------------------- | :----------------------------------------------------------- | :--------------------------------------------------------- |
| `primary`            | Whether the column is the primary key. To have a composite key look [below](https://gobuffalo.io/documentation/database/fizz/#composite-primary-keys) 列是否是主键。要查看复合键，请在下面查找 | `{"primary": true}`                                        |
| `size`               | The size of the column. The default value for a string column is 255 (or 191 for MariaDB) 列的大小。字符串列的默认值为 255（或 MariaDB 的 191） | `{"size": 50}`                                             |
| `scale`, `precision` | The scale and the precision for a float column 浮点列的比例和精度 | `{"scale": 4, "precision": 2}`                             |
| `null`               | By default columns are not allowed to be `null` 默认情况下，不允许列为 `null` | `{"null": true}`                                           |
| `default`            | The default value you want for this column. By default this is `null` 您希望为此列设置的默认值。默认情况下，这是 `null` | `{"default": 0}` `{"default": false}` `{"default": "foo"}` |
| `default_raw`        | The default value defined as a database function 定义为数据库函数的默认值 | `{"default_raw": "uuid_generate_v1()"}`                    |
| `after`              | (**MySQL Only**) Add a column after another column in the table (仅限 MySQL) 在表中另一个列之后添加列 | `{"after": "created_at"}`                                  |
| `first`              | (**MySQL Only**) Add a column to the first position in the table (仅限 MySQL) 将列添加到表中的第一个位置 | `{"first": true}`                                          |

#### Disable Auto Timestamps 禁用自动时间戳 

```javascript
create_table("users") {
  t.Column("id", "uuid", {primary: true})
  // ...
  // Disable auto-creation of created_at and updated_at columns
  t.DisableTimestamps()
}
```

or

```javascript
create_table("users", {timestamps: false}) {
  t.Column("id", "uuid", {primary: true})
  // ...
}
```

## Drop a Table 删除表 

```javascript
drop_table("table_name")
```

## Rename a Table 重命名表 

```javascript
rename_table("old_table_name", "new_table_name")
```

## Add a Column 添加列 

```javascript
add_column("table_name", "column_name", "string", {})
```

See [above](https://gobuffalo.io/documentation/database/fizz/#common-types) for more details on column types and options.

​	请参阅上文以了解有关列类型和选项的更多详细信息。

## Alter a column 更改列 

```javascript
change_column("table_name", "column_name", "string", {})
```

## Rename a Column 重命名列 

```javascript
rename_column("table_name", "old_column_name", "new_column_name")
```

## Drop a Column 删除列 

```javascript
drop_column("table_name", "column_name")
```

## Composite Primary Keys 复合主键 

```javascript
t.PrimaryKey("column_1", "column_2")
```

Please note that the `t.PrimaryKey` statement MUST be after the columns definitions.

​	请注意， `t.PrimaryKey` 语句必须在列定义之后。

## Add an Index 添加索引 

#### Supported Options: 支持的选项： 

- `name` - This defaults to `table_name_column_name_idx`
  `name` - 默认为 `table_name_column_name_idx`
- `unique`

### Simple Index: 简单索引： 

```javascript
add_index("table_name", "column_name", {})
```

### Multi-Column Index: 多列索引： 

```javascript
add_index("table_name", ["column_1", "column_2"], {})
```

### Unique Index: 唯一索引：

```javascript
add_index("table_name", "column_name", {"unique": true})
```

### Index Names: 索引名称：

```javascript
add_index("table_name", "column_name", {}) # name => table_name_column_name_idx
add_index("table_name", "column_name", {"name": "custom_index_name"})
```

## Rename an Index 重命名索引 

```javascript
rename_index("table_name", "old_index_name", "new_index_name")
```

## Drop an Index 删除索引 

```javascript
drop_index("table_name", "index_name")
```

## Add a Foreign Key 添加外键 

```javascript
add_foreign_key("table_name", "field", {"ref_table_name": ["ref_column"]}, {
    "name": "optional_fk_name",
    "on_delete": "action",
    "on_update": "action",
})
```

#### Supported Options 支持的选项 

- `name` - This defaults to `table_name_ref_table_name_ref_column_name_fk`
  `name` - 默认为 `table_name_ref_table_name_ref_column_name_fk`
- `on_delete` - `CASCADE`, `SET NULL`, …
- `on_update`

**Note:** `on_update` and `on_delete` are not supported on CockroachDB yet.
注意：CockroachDB 尚不支持 `on_update` 和 `on_delete` 。

## Drop a Foreign Key 删除外键 

```javascript
drop_foreign_key("table_name", "fk_name", {"if_exists": true})
```

#### Supported Options 支持的选项 

- `if_exists` - Adds `IF EXISTS` condition
  `if_exists` - 添加 `IF EXISTS` 条件

## Raw SQL 原始 SQL 

```javascript
sql("select * from users;")
```

## Execute an External Command 执行外部命令 

Sometimes during a migration you need to shell out to an external command.

​	有时在迁移期间，您需要使用外部命令进行 shell 操作。

```javascript
exec("echo hello")
```