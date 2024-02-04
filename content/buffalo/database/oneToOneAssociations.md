+++
title = "一对一关联"
date = 2024-02-04T21:15:39+08:00
weight = 15
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/relations-one-to-one/](https://gobuffalo.io/documentation/database/relations-one-to-one/)

# One to One Associations 一对一关联 

In this chapter, you’ll learn how to write a one to one [association](https://gobuffalo.io/documentation/database/relations/) in Pop.

​	在本章中，您将学习如何在 Pop 中编写一对一关联。

## Tags 标签 

One to one associations work using a pair of tags:

​	一对一关联使用一对标签：

- `belongs_to` for the model with the foreign key.
  `belongs_to` 用于具有外键的模型。
- `has_one` for the model without the foreign key.
  `has_one` 用于没有外键的模型。

## Example 示例 

```go
// Models

type Head struct {
  ID           int
  BodyID       int        `db:"body_id"`
  Body         *Body      `belongs_to:"body"`
}

type Body struct {
  ID           int
  Head         Head       `has_one:"head"`
}
// Eager creation:
// Create a body with its head.
b := &models.Body{
    Head: models.Head{},
}

if err := tx.Eager().Create(b); err != nil {
    return err
}
// Eager fetch all bodies with their head.
bodies = &models.Bodies{}

if err := c.Eager().All(bodies); err != nil {
    log.Printf("err: %v", err)
    return
}

log.Printf("eager fetch: %v", bodies)
```

## Related Content 相关内容 

- [Associations with Pop: 1 to 1](https://blog.gobuffalo.io/associations-with-pop-1-to-1-592f02e2bdd8) - An article about 1 to 1 associations in Pop.
  使用 Pop 的关联：1 对 1 - 一篇关于 Pop 中 1 对 1 关联的文章。