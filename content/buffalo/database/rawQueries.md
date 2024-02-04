+++
title = "原生查询"
date = 2024-02-04T21:14:42+08:00
weight = 11
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/raw-queries/](https://gobuffalo.io/documentation/database/raw-queries/)

# Raw Queries 原始查询 

Sometimes you’ll need to write a custom query instead of letting Pop generate it for you. In this chapter, you’ll learn how to write raw SQL queries using Pop.

​	有时，您需要编写自定义查询，而不是让 Pop 为您生成查询。在本章中，您将学习如何使用 Pop 编写原始 SQL 查询。

## Writing a Raw Query 编写原始查询 

### Select 选择 

```go
player := Player{}
q := db.RawQuery("SELECT * FROM players WHERE id = ?", 1)
err := q.Find(&player, id)
```

### Update 更新 

```go
err := db.RawQuery("UPDATE players SET instrument = ? WHERE id = ?", "guitar", 1).Exec()
```

### Delete 删除 

```go
err := db.RawQuery("DELETE FROM players WHERE id = ?", 1).Exec()
```

## Tokens Syntax 令牌语法 

With `RawQuery`, you can continue to use the `?` tokens to secure your input values. You don’t need to use the token syntax for your underlying database.

​	使用 `RawQuery` ，您可以继续使用 `?` 令牌来保护您的输入值。您无需为基础数据库使用令牌语法。