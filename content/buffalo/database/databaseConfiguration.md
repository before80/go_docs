+++
title = "数据库配置"
date = 2024-02-04T21:12:29+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/configuration/]({{< ref "/buffalo/database/databaseConfiguration" >}})

# Database Configuration 数据库配置 

Pop configuration is managed by a `database.yml` file, located at the root of your project. This file is generated for you if you use Buffalo – if you choose to use Pop – and contains a basic configuration for the database you selected on generation with the `--db-type` flag. PostgreSQL is considered as the default.

​	Pop 配置由位于项目根目录的 `database.yml` 文件管理。如果您使用 Buffalo（如果您选择使用 Pop），则会为您生成此文件，其中包含使用 `--db-type` 标志在生成时选择的数据库的基本配置。PostgreSQL 被视为默认值。

Here is a sample configuration generated for a new app based on PostgreSQL:

​	以下是为基于 PostgreSQL 的新应用生成的示例配置：

```yaml
development:
  dialect: postgres
  database: myapp_development
  user: postgres
  password: postgres
  host: 127.0.0.1
  pool: 5

test:
  url: {{envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"}}

production:
  url: {{envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_production"}}
```

You can see three connections defined:

​	您可以看到定义了三个连接：

- `development` is the one used when your app runs on dev mode.
  `development` 是在您的应用在开发模式下运行时使用的连接。
- `test` serves to run the integration tests.
  `test` 用于运行集成测试。
- `production` is the config you’ll use on the final app, on the server.
  `production` 是您将在服务器上的最终应用中使用的配置。

Of course, you can configure any new connection you want, but Buffalo won’t pick them by default.

​	当然，您可以配置所需的任何新连接，但 Buffalo 不会默认选择它们。

## Generator 生成器 

**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So every time you want to use a command from `soda`, just execute `buffalo pop` instead.
Buffalo 用户须知： `soda` 命令嵌入到 `buffalo` 命令中，位于 `pop` 命名空间之后。因此，每次您想使用 `soda` 中的命令时，只需执行 `buffalo pop` 即可。

You can generate a default configuration file using the init command:

​	您可以使用 init 命令生成默认配置文件：

```bash
$ soda g config
```

The default will generate a `database.yml` file in the current directory for a PostgreSQL database. You can override the type of database using the `-t` flag and passing in any of the supported database types: `postgres`, `cockroach`, `mysql`, or `sqlite3`.

​	默认情况下，将在当前目录中为 PostgreSQL 数据库生成一个 `database.yml` 文件。您可以使用 `-t` 标志并传入任何受支持的数据库类型来覆盖数据库类型： `postgres` 、 `cockroach` 、 `mysql` 或 `sqlite3` 。

## Config File Location 配置文件位置 

The Pop configuration file – `database.yml` – can be found either:

​	Pop 配置文件 – `database.yml` – 可以位于以下位置：

- At your project root (default).
  在您的项目根目录（默认）。
- In the `config/` directory, at your project root.
  在您的项目根目录的 `config/` 目录中。

If you want to put your config file in another location, you can use the [`AddLookupPaths`](https://godoc.org/github.com/gobuffalo/pop#AddLookupPaths).

​	如果您想将配置文件放在其他位置，可以使用 `AddLookupPaths` 。

You can also customize the file name:

​	您还可以自定义文件名：

```go
pop.ConfigName = "my_pop_config.yml"
```

## Env vs Detailed Configuration 环境变量与详细配置 

Note that the `database.yml` file is also a Go template, so you can use Go template syntax. There are two special functions that are included, `env` and `envOr`.
请注意， `database.yml` 文件也是一个 Go 模板，因此您可以使用 Go 模板语法。包含两个特殊函数， `env` 和 `envOr` 。

As you can see, you have two ways to configure a new connection:

​	如您所见，您可以通过两种方式配置新连接：

- The one used by the `development` connection is the most detailed. It allows you to set each available parameter, one by one.
  `development` 连接使用的方式最为详细。它允许您逐个设置每个可用参数。
- The one used by the `test` and `production` connections is a bit different: it uses a variable (see the `{{ }}` marks?) to set the value, and the `envOr` helper.
  `test` 和 `production` 连接使用的方式略有不同：它使用变量（请参阅 `{{ }}` 标记？）来设置值，以及 `envOr` 帮助器。

The `envOr` helper tries to get a value from an environment variable, and default to the second value. For instance:

​	 `envOr` 帮助器尝试从环境变量中获取值，并默认为第二值。例如：

```yaml
envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"
```

Tries to get the `TEST_DATABASE_URL` value from environment, and defaults to `postgres://postgres:postgres@127.0.0.1:5432/myapp_test`.

​	尝试从环境中获取 `TEST_DATABASE_URL` 值，并默认为 `postgres://postgres:postgres@127.0.0.1:5432/myapp_test` 。

This way, you can provide a default value for development purposes, and allow to reconfigure the database settings from an environment variable!

​	这样，您可以为开发目的提供默认值，并允许从环境变量重新配置数据库设置！

The `url` param for a connection will override any other connection param. Make sure you set all the settings you want from the URL string.
连接的 `url` 参数将覆盖任何其他连接参数。请确保从 URL 字符串中设置所需的所有设置。

For additional details, check the documentation for [github.com/gobuffalo/pop](https://github.com/gobuffalo/pop).

​	有关其他详细信息，请查看github.com/gobuaffalo/pop的文档。

**Make sure you have configured this file properly before working with Pop!
在使用 Pop 之前，请确保已正确配置此文件！**

## Available Options 可用选项 

### database 数据库 

The name of the database to use.

​	要使用的数据库的名称。

### dialect 方言 

The database dialect to use with the connection. Accepted values are:

​	与连接一起使用的数据库方言。可接受的值是：

- MySQL driver: “mysql”
  MySQL 驱动程序：“mysql”
- PostgreSQL driver: “postgres”, “postgresql” or “pg”
  PostgreSQL 驱动程序：“postgres”、“postgresql”或“pg”
- Cockroach driver: “cockroach”, “cockroachdb” or “crdb”
  Cockroach 驱动程序：“cockroach”、“cockroachdb”或“crdb”
- SQLite driver: “sqlite” or “sqlite3”
  SQLite 驱动程序：“sqlite”或“sqlite3”

### driver 驱动程序 

Since **4.11.2**
自 4.11.2 起



Use this option to customize the database driver and override the default one used by Pop.

​	使用此选项自定义数据库驱动程序并覆盖 Pop 使用的默认驱动程序。

Here is the list of the default SQL drivers:

​	以下是默认 SQL 驱动程序的列表：

- MySQL: [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
  MySQL：github.com/go-sql-driver/mysql
- PostgreSQL: [github.com/lib/pq](https://github.com/lib/pq)
  PostgreSQL：github.com/lib/pq
- Cockroach DB: [github.com/cockroachdb/cockroach-go/crdb](https://github.com/cockroachdb/cockroach-go/tree/master/crdb)
  Cockroach DB：github.com/cockroachdb/cockroach-go/crdb
- SQLite: [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
  SQLite：github.com/mattn/go-sqlite3

### encoding 编码 

Since **4.6.0**
自 4.6.0 起



This option is currently only supported by the **mysql dialect**. This encoding will be used to create the database (if you create it using `soda`), and as the `collation` parameter for the connection string. If this option is omitted, the default value is `utf8mb4_general_ci`.

​	此选项目前仅受 mysql 方言支持。此编码将用于创建数据库（如果您使用 `soda` 创建），并用作连接字符串的 `collation` 参数。如果省略此选项，则默认值为 `utf8mb4_general_ci` 。

```yaml
development:
  dialect: mysql
  options:
    encoding: "utf8_general_ci"
```

### host 主机 

The database host address to connect to.

​	要连接到的数据库主机地址。

### password 密码 

The password for the user you use to connect to the database.

​	用于连接到数据库的用户的密码。

### port 端口 

The database host port for the database.

​	数据库的主机端口。

**Defaults**:

​	默认值：

| Driver 驱动程序 | Port 端口 |
| :-------------- | :-------- |
| PostgreSQL      | 5432      |
| MySQL           | 3306      |
| Cockroach       | 26257     |

### user 用户 

The user to use to connect to the database.

​	用于连接到数据库的用户。
