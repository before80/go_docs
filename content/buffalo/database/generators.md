+++
title = "生成器"
date = 2024-02-04T21:13:42+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/generators/]({{< ref "/buffalo/database/generators" >}})

# Generators 生成器 

```bash
$ buffalo pop g --help
Generates config, model, and migrations files.

Usage:
  buffalo-pop pop generate [command]

Aliases:
  generate, g

Available Commands:
  config      Generates a database.yml file for your project.
  fizz        Generates Up/Down migrations for your database using fizz.
  model       Generates a model for your database
  sql         Generates Up/Down migrations for your database using SQL.

Flags:
  -h, --help   help for generate

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")

Use "buffalo-pop pop generate [command] --help" for more information about a command.
```

## Migrations 迁移 

For information on generating migrations see [Migrations]({{< ref "/buffalo/database/migrations" >}}).

​	有关生成迁移的信息，请参阅迁移。
