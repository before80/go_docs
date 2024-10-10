+++
title = "一对多关联"
date = 2024-02-04T21:15:49+08:00
weight = 16
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/relations-one-to-many/]({{< ref "/buffalo/database/oneToManyAssociations" >}})

# One to Many Associations 一对多关联 

In this chapter, you’ll learn how to write a one to many [association]({{< ref "/buffalo/database/associationsAndRelationships" >}}) in Pop.

​	在本章中，您将学习如何在 Pop 中编写一对多关联。

## Tags 标签 

One to many associations work using a pair of tags:

​	一对多关联使用一对标签：

- `belongs_to` for the model with the foreign key.
  `belongs_to` 用于具有外键的模型。
- `has_many` for the model without the foreign key (the one with the slice).
  `has_many` 用于没有外键的模型（带有切片的模型）。

## Example 示例 

```go
// Models

type Fruit struct {
    ID     int   `json:"id,omitempty" db:"id"`
    TreeID int   `json:"-" db:"tree_id"`
    Tree   *Tree `json:"tree,omitempty" belongs_to:"tree"`
}
    
type Tree struct {
    ID     int     `json:"id" db:"id"`
    Name   string  `json:"name" db:"name"`
    Fruits []Fruit `json:"fruits,omitempty" has_many:"fruits"`
}
// Eager creation:
// Create an apple tree with 2 fruits.
t := &models.Tree{
    Name: "Apple tree",
    Fruits: []models.Fruit{
        {},
        {},
    },
}

if err := tx.Eager().Create(t); err != nil {
    return err
}
// Eager fetch all the trees with their fruits.
trees := &models.Trees{}

if err := c.Eager().All(trees); err != nil {
    log.Printf("err: %v", err)
    return
}

log.Printf("eager fetch: %v", trees)
```

## Custom Association Order 自定义关联顺序 

Since `has_many` is mapped to a slice, you’ll probably want to customize the order of this slice. `order_by` tag allows you to indicate the order for the association when loading it:

​	由于 `has_many` 映射到切片，因此您可能希望自定义此切片的顺序。 `order_by` 标签允许您在加载关联时指示关联的顺序：

```go
type Tree struct {
    ID     int     `json:"id" db:"id"`
    Name   string  `json:"name" db:"name"`
    Fruits []Fruit `json:"fruits,omitempty" has_many:"fruits" order_by:"id desc"`
}
```

The format to use is `order_by:"<column_name> <asc | desc>"`.

​	要使用的格式是 `order_by:"<column_name> <asc | desc>"` 。

## Customize Foreign Keys Lookup 自定义外键查找 

By default, `has_many` will fetch related records using a convention for the foreign key column. In our previous example, the `fruits` table (mapped to the `Fruit` struct) contains a `tree_id` foreign key column which references the ID of the tree the fruit is attached to.

​	默认情况下， `has_many` 将使用外键列的约定来获取相关记录。在我们的上一个示例中， `fruits` 表（映射到 `Fruit` 结构）包含一个 `tree_id` 外键列，该列引用水果所附树的 ID。

You can use the `fk_id` tag to customize this foreign key column:

​	您可以使用 `fk_id` 标记来自定义此外键列：

```go
type Tree struct {
    ID     int     `json:"id" db:"id"`
    Name   string  `json:"name" db:"name"`
    Fruits []Fruit `json:"fruits,omitempty" has_many:"fruits" fk_id:"custom_tree_id"`
}
```

Here, the relation will be looked up using the column `custom_tree_id` in the `fruits` table, instead of the default `tree_id` one.

​	在此，将使用 `fruits` 表中的列 `custom_tree_id` 来查找关系，而不是默认的 `tree_id` 列。

This can be really useful when you have structs with multiple fields pointing to the same model:

​	当您有多个字段指向同一模型的结构时，这非常有用： 与 Pop 的关联：1 对 n - 一篇关于 Pop 中 1 对 n 关联的文章。

```go
type Player struct {
    ID            int     `json:"id" db:"id"`
    Name          string  `json:"name" db:"name"`
    CurrentBandID int     `json:"current_band_id" db:"current_band_id"`
    FormerBandID  int     `json:"former_band_id" db:"former_band_id"`
}

type Band struct {
    ID             int      `json:"id" db:"id"`
    Name           string   `json:"name" db:"name"`
    CurrentPlayers []Player `json:"current_players,omitempty" has_many:"players" fk_id:"current_band_id"`
    FormerPlayers  []Player `json:"former_players,omitempty" has_many:"players" fk_id:"former_band_id"`
}
```

## Related Content 相关内容 

- [Associations with Pop: 1 to n](https://blog.gobuffalo.io/associations-with-pop-1-to-n-2fb3e1c3833f) - An article about 1 to n associations in Pop.
