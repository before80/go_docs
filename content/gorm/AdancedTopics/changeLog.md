+++
title = "变更记录"
date = 2023-10-28T14:37:23+08:00
weight = 14
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/changelog.html](https://gorm.io/docs/changelog.html)

## v2.0 - 2020.08

GORM 2.0 is a rewrite from scratch, it introduces some incompatible-API change and many improvements

​	GORM 2.0 是从头开始重写的，它引入了一些不兼容的API更改和许多改进。

- 性能改进 Performance Improvements
- 模块化 Modularity
- 支持上下文、批量插入、预处理语句模式、DryRun模式、预加载连接、查找到映射、从映射创建、分批查找等功能 Context, Batch Insert, Prepared Statement Mode, DryRun Mode, Join Preload, Find To Map, Create From Map, FindInBatches supports
- 支持嵌套事务/保存点/回滚到保存点 Nested Transaction/SavePoint/RollbackTo SavePoint supports
- 支持命名参数、分组条件、更新、锁定、优化器/索引/注释提示、子查询改进等功能 Named Argument, Group Conditions, Upsert, Locking, Optimizer/Index/Comment Hints supports, SubQuery improvements
- 支持完全自引用关系、连接表改进、批处理数据的关联模式 Full self-reference relationships supports, Join Table improvements, Association Mode for batch data
- 支持多字段跟踪创建/更新时间，添加对UNIX（毫秒/纳秒）的支持 Multiple fields support for tracking create/update time, which adds support for UNIX (milli/nano) seconds
- 支持字段权限：只读、仅写、仅创建、仅更新、忽略 Field permissions support: read-only, write-only, create-only, update-only, ignored
- 新的插件系统：支持多个数据库、带有插件数据库解析器的读写拆分支持、Prometheus集成等 New plugin system: multiple databases, read/write splitting support with plugin Database Resolver, prometheus integrations…
- 新的钩子API：与插件的统一接口 New Hooks API: unified interface with plugins
- 新的迁移器：允许为关系创建数据库外键、约束/检查器支持、增强索引支持 New Migrator: allows to create database foreign keys for relationships, constraints/checker support, enhanced index support
- 新的日志记录器：支持上下文、改进可扩展性 New Logger: context support, improved extensibility
- 统一的命名策略：表名、字段名、连接表名、外键、检查器、索引名规则 Unified Naming strategy: table name, field name, join table name, foreign key, checker, index name rules
- 更好的自定义数据类型支持（例如：JSON） Better customized data type support (e.g: JSON)

[GORM 2.0 Release Note](https://gorm.io/docs/v2_release_note.html)

## v1.0 - 2016.04

[GORM V1 Docs](https://v1.gorm.io/)

Breaking Changes:

- `gorm.Open` returns `*gorm.DB` instead of `gorm.DB`
- Updating will only update changed fields
- Soft Delete’s will only check `deleted_at IS NULL`
- New ToDBName logic
  Common initialisms from [golint](https://github.com/golang/lint/blob/master/lint.go#L702) like `HTTP`, `URI` was converted to lowercase, so `HTTP`‘s db name is `http`, but not `h_t_t_p`, but for some other initialisms not in the list, like `SKU`, it’s db name was `s_k_u`, this change fixed it to `sku`
- Error `RecordNotFound` has been renamed to `ErrRecordNotFound`
- `mssql` dialect has been renamed to `github.com/jinzhu/gorm/dialects/mssql`
- `Hstore` has been moved to package `github.com/jinzhu/gorm/dialects/postgres`