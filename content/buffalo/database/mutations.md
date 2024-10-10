+++
title = "突变"
date = 2024-02-04T21:14:14+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/mutations/]({{< ref "/buffalo/database/mutations" >}})

# Mutations 突变 

In this chapter, you’ll learn how to create, update and delete objects from you database using Pop.

​	在本章中，您将学习如何使用 Pop 创建、更新和删除数据库中的对象。

## Create 创建 

### Without validation 无验证 

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Create a fruit without running validations
err := c.Create(fruit)
```

### With validation 有验证 

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Run validations and create if all validations passed
vErrors, err := c.ValidateAndCreate(fruit)
```

## Update 更新 

### Without validation 无验证 

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Update a fruit without running validations
err := c.Update(fruit)
```

### With validation 有验证 

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Run validations and update if all validations passed
vErrors, err := c.ValidateAndUpdate(fruit)
```

## Save 保存 

Save checks for the ID in you model: if the ID is the zero value of the type (so for example if it’s an `int` and its value is `0`), `Save` calls `Create`. Otherwise, it calls `Update`.

​	在模型中保存 ID 的检查：如果 ID 是该类型的零值（例如，如果它是 `int` 且其值为 `0` ），则 `Save` 调用 `Create` 。否则，它调用 `Update` 。

### Without validation 无验证 

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 0}
// Create a fruit without running validations
err := c.Save(fruit)
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 1}
// Update a fruit without running validations
err := c.Save(fruit)
```

### With validation 有验证 

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 0}
// Run validations and create if all validations passed
vErrors, err := c.ValidateAndSave(fruit)
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 1}
// Run validations and update if all validations passed
vErrors, err := c.ValidateAndSave(fruit)
```

## Delete 删除 

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 1}
// Destroy the fruit
err := c.Destroy(fruit)
```

## Next Steps 后续步骤 

- [Querying]({{< ref "/buffalo/database/querying" >}}) - Fetch the data you inserted in the database.
  查询 - 获取您在数据库中插入的数据。
- [Associations and Relationships]({{< ref "/buffalo/database/associationsAndRelationships" >}}) - Handle relations between models.
  关联和关系 - 处理模型之间的关系。
