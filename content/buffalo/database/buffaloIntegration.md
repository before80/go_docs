+++
title = "Buffalo 集成"
date = 2024-02-04T21:13:08+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/buffalo-integration/]({{< ref "/buffalo/database/buffaloIntegration" >}})

# Buffalo Integration Buffalo 集成 

## Generate a New App 生成新应用 

When you generate a new Buffalo application you can choose the target database with the `--db-type` flag. For instance, to generate a new app with MySQL database support, you can write the following:

​	生成新的 Buffalo 应用程序时，可以使用 `--db-type` 标志选择目标数据库。例如，要生成一个支持 MySQL 数据库的新应用，可以编写以下内容：

```bash
$ buffalo new coke --db-type mysql
```

**By default, Buffalo will generate an app with PostgreSQL as the backing database.
默认情况下，Buffalo 将生成一个以 PostgreSQL 作为后端数据库的应用。**

### Skip Database Support 跳过数据库支持 

If you want to handle the database without using Pop, or if you’re building an app without database, it’s also possible to skip generation of all database components with the `--skip-pop` flag.

​	如果想在不使用 Pop 的情况下处理数据库，或者正在构建没有数据库的应用，也可以使用 `--skip-pop` 标志跳过所有数据库组件的生成。

```bash
$ buffalo new coke --skip-pop
```

## The Pop Transaction Middleware Pop 事务中间件 

Buffalo provides a Pop middleware to ease database usage within Buffalo: https://github.com/gobuffalo/buffalo-pop

​	Buffalo 提供了一个 Pop 中间件，以便在 Buffalo 中轻松使用数据库：https://github.com/gobuffalo/buffalo-pop

### Setup 设置 

This middleware is configured for you by default, if you choose to use Pop when creating a new project.

​	如果您在创建新项目时选择使用 Pop，则默认情况下会为您配置此中间件。

**actions/app.go**

```go
func App() *buffalo.App {
	if app == nil {
        // [...]

        app.Use(poptx.PopTransaction(models.DB))

        // [...]

        app.GET("/", HomeHandler)
    }

    return app
}
```

`poptx.PopTransaction(models.DB)` uses the connection to the configured database to create a new `PopTransaction` middleware. This middleware does the following:

​	 `poptx.PopTransaction(models.DB)` 使用与已配置数据库的连接来创建新的 `PopTransaction` 中间件。此中间件执行以下操作：

- Log the total duration spent during the request making database calls.
  记录在请求期间进行数据库调用的总持续时间。
- Wrap **each HTTP request** in a database transaction.
  将每个 HTTP 请求包装在数据库事务中。
- Commit **if there was no error** executing the middlewares and action; **and the response status is a 2xx or 3xx**.
  如果执行中间件和操作时没有错误；并且响应状态为 2xx 或 3xx，则提交。
- Rollback otherwise.
  否则回滚。

### Handle Transaction By Hand 手动处理事务 

If you need to handle a transaction by hand, you can skip the middleware for a given route:

​	如果您需要手动处理事务，则可以跳过给定路由的中间件：

```go
func App() *buffalo.App {
	if app == nil {
        // [...]
        txm := poptx.PopTransaction(models.DB)
        app.Use(txm)
        a.Middleware.Skip(txm, HomeHandler)

        // [...]

        app.POST("/form", FormHandler)
        app.GET("/", HomeHandler)
    }

    return app
}
```
