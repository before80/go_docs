+++
title = "查询"
date = 2024-02-04T21:14:28+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/querying/]({{< ref "/buffalo/database/querying" >}})

# Querying 查询 

In this chapter, you’ll learn how to retrieve data from your database using Pop.

​	在本章中，您将学习如何使用 Pop 从数据库中检索数据。

### Find By ID 按 ID 查找 

```go
user := User{}
err := models.DB.Find(&user, id)
```

### Find All 查找全部 

```go
users := []User{}
err := models.DB.All(&users)
err = models.DB.Where("id in (?)", 1, 2, 3).All(&users)
```

### Find All with Order 按顺序查找全部 

```go
// To retrieve records from the database in a specific order, you can use the Order method
users := []User{}
err := models.DB.Order("id desc").All(&users)
```

### Find Last 查找最后一条 

```go
// Last() orders by created_at
user := models.User{}
err := tx.Last(&user)
```

### Find Where 按条件查找 

```go
users := []models.User{}
query := models.DB.Where("id = 1").Where("name = 'Mark'")
err := query.All(&users)

err = tx.Where("id in (?)", 1, 2, 3).All(&users)
```

### Using `in` Queries 使用 `in` 查询 # 遗憾的是，由于各种原因，您无法在同一个 调用中同时使用 `in` 查询和 查询。

```go
err = models.DB.Where("id in (?)", 1, 2, 3).All(&users)
err = models.DB.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

Unfortunately, for a variety of reasons you can’t use an `and` query in the same `Where` call as an `in` query.

```go
// does not work:
err := tx.Where("id in (?) and foo = ?", 1, 2, 3, "bar").All(&users)

// works:
err := tx.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

### Select specific columns 选择特定列 

`Select` allows you to load specific columns from a table. Useful when you don’t want all columns from a table to be loaded in a query.

​	 `Select` 允许您从表中加载特定列。当您不希望表中的所有列都加载到查询中时，此功能非常有用。

```go
err = tx.Select("name").All(&users)
// SELECT name FROM users

err = tx.Select("max(age)").All(&users)
// SELECT max(age) FROM users

err = tx.Select("age", "name").All(&users)
// SELECT age, name FROM users
```

### Join Query 联接查询 

```go
// page: page number
// perPage: limit

roles := []models.UserRole{}

q := tx.Q()
q.LeftJoin("roles", "roles.id = user_roles.role_id")
q.LeftJoin("users u", "u.id = user_roles.user_id")
q.Where(`roles.name like ?`, name)
q.Paginate(page, perPage)

err := q.All(&roles)
```

### Count records 计数记录 

```go
query := tx.Q()
count, err := query.Count(&models.User{})
query := tx.Q()
count, err := query.Where("name = ?", "John").Count(&models.User{})
query := tx.Q()
count, err := query.CountByField(&models.User{}, "first_name")
// Equals to
count, err := query.Count(&models.User{},, "first_name")
```

### Pop to SQL 弹出到 SQL 

```go
q := tx.Q()
q.LeftJoin("roles", "roles.id = user_roles.role_id")
q.LeftJoin("users u", "u.id = user_roles.user_id")
q.Where(`roles.name like ?`, "john")
q.Paginate(1, 20)

popModel := &pop.Model{Value: models.UserRole{}}
cols := []string{"user_roles.*", "roles.name as role_name", "u.first_name", "u.last_name"}

sql, args := tx.Q().ToSQL(popModel, cols...)
```

sql

args

```sql
-- The original query is in one line
SELECT
    user_roles.*,
    roles.name as role_name,
    u.first_name,
    u.last_name
FROM
    user_roles AS user_roles
LEFT JOIN
    roles ON roles.id = user_roles.role_id
LEFT JOIN
    users u ON u.id = user_roles.user_id
WHERE
    roles.name like $1
LIMIT
    20
OFFSET
    0
```
