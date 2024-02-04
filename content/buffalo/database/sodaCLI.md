+++
title = "Soda CLI"
date = 2024-02-04T21:12:13+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/soda/]({{< ref "/buffalo/database/sodaCLI" >}})

# Soda CLI

Pop helps you to manage database connections, but it also provides `soda`, a small CLI toolbox to manage your database. It can help you to create a new database, drop existing ones, and so on.

​	Pop 帮助您管理数据库连接，但它还提供 `soda` ，一个用于管理数据库的小型 CLI 工具箱。它可以帮助您创建新数据库、删除现有数据库，等等。

**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So every time you want to use a command from `soda`, just execute `buffalo pop` instead. You don’t need to install `soda` CLI.
Buffalo 用户注意： `soda` 命令已嵌入到 `buffalo` 命令中，位于 `pop` 命名空间之后。因此，每次您想使用 `soda` 中的命令时，只需执行 `buffalo pop` 即可。您无需安装 `soda` CLI。

## Installing CLI Support 安装 CLI 支持 

### From a release archive 从发行存档 

Pre-compiled archives contain Soda **with SQLite support**.

​	预编译存档包含带有 SQLite 支持的 Soda。

Download the appropriate version for your platform from [Pop releases](https://github.com/gobuffalo/pop/releases).

​	从 Pop 发行版下载适用于您平台的相应版本。

Place it somewhere in your `PATH`, and ensure the `soda` binary is executable.

​	将其放在 `PATH` 中的某个位置，并确保 `soda` 二进制文件可执行。

### Homebrew (macOS) Homebrew（macOS）

```console
$ brew install gobuffalo/tap/pop
```

### From source 从源代码 

For go version 1.16 and later,

​	对于 go 版本 1.16 及更高版本，

**Without** sqlite 3 support:

​	不支持 sqlite 3：

```console
$ go install github.com/gobuffalo/pop/v6/soda@latest
```

**With** sqlite 3 support (requires GCC or equivalent C compiler):

​	支持 sqlite 3（需要 GCC 或同等 C 编译器）：

```console
$ go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest
```

If you’re not building your code with `buffalo build`, you’ll also have to pass `-tags sqlite` to `go build` when building your program.

​	如果您没有使用 `buffalo build` 构建代码，则在构建程序时还必须将 `-tags sqlite` 传递给 `go build` 。

## Creating Databases 创建数据库 

Once the `database.yml` has been configured with the appropriate settings, and the database server is running, Soda can create all of the databases in the `database.yml` file with a simple command:

​	一旦使用适当的设置配置了 `database.yml` ，并且数据库服务器正在运行，Soda 就可以使用一个简单命令创建 `database.yml` 文件中的所有数据库：

```console
$ soda create -a
```

You can also create just one of the configured databases by using the `-e` flag and the name of the database:

​	您还可以使用 `-e` 标志和数据库名称来创建其中一个已配置的数据库：

```console
$ soda create -e test
```

## Dropping Databases 删除数据库 

Soda can drop all of your databases, should you want to, with one command:

​	如果您愿意，Soda 可以使用一个命令删除您的所有数据库：

```console
$ soda drop -a
```

You can also drop just one of the configured databases by using the `-e` flag and the name of the database:

​	您还可以使用 `-e` 标志和数据库名称来删除其中一个已配置的数据库：

```console
$ soda drop -e test
```
