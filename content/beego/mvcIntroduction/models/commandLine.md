+++
title = "命令行"
date = 2024-02-04T10:02:08+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/model/cmd/](https://beego.wiki/docs/mvc/model/cmd/)

# Command Line 命令行



## Command Line 命令行

You can call `orm.RunCommand()` after you registered models and database(s) as follows:

&zeroWidthSpace;在注册模型和数据库后，您可以按如下方式调用 `orm.RunCommand()` ：

```go
func main() {
	// orm.RegisterModel...
	// orm.RegisterDataBase...
	...
	orm.RunCommand()
}
go build main.go
./main orm
# Get help by just run it.
# If possible, go run main.go orm has the same result.
```

## Database Schema Generation 数据库架构生成

```bash
./main orm syncdb -h
Usage of orm command: syncdb:
  -db="default": DataBase alias
  -force=false: drop tables before create
  -v=false: verbose info
```

Use the `-force=1` flag to force drop tables and re-create.

&zeroWidthSpace;使用 `-force=1` 标志强制删除表并重新创建。

Use the `-v` flag to print SQL statements.

&zeroWidthSpace;使用 `-v` 标志打印 SQL 语句。

------

Use program to create tables:

&zeroWidthSpace;使用程序创建表：

```go
// Database alias.
name := "default"

// Drop table and re-create.
force := true

// Print log.
verbose := true

// Error.
err := orm.RunSyncdb(name, force, verbose)
if err != nil {
	fmt.Println(err)
}
```

Even if you do not enable `force` mode, ORM also will auto-add new fields and indexes, but you have to deal with delete operations yourself.

&zeroWidthSpace;即使您未启用 `force` 模式，ORM 也会自动添加新字段和索引，但您必须自己处理删除操作。

## Print SQL Statements 打印 SQL 语句

```bash
./main orm sqlall -h
Usage of orm command: syncdb:
  -db="default": DataBase alias name
```

Use database with alias `default` as default.

&zeroWidthSpace;使用具有别名 `default` 的数据库作为默认数据库。