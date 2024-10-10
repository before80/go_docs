+++
title = "作用域"
date = 2024-02-04T21:15:01+08:00
weight = 13
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/scoping/]({{< ref "/buffalo/database/scoping" >}})

# Scoping 作用域 

Scoping is a way to structure your DB calls, when it needs the same “base” query. Let’s say you want to create a book store: the store provides books for everyone, but some special editions are reserved to customers with a registered account. It means that for the whole store, you’ll need to filter the books, so the “guest” customers can only see the restricted list of books.

​	作用域是一种构建数据库调用（当它需要相同的“基本”查询时）的方式。假设您想创建一个书店：该商店为每个人提供书籍，但某些特别版本仅限于拥有注册帐户的客户。这意味着对于整个商店，您需要过滤书籍，以便“访客”客户只能看到受限的书籍列表。

## The Usual Way 通常的方式 

A “naive” way can be writing each full query.

​	一种“天真的”方式可以是编写每个完整查询。

```go
type Book struct {
    ID         uuid.UUID `json:"id" db:"id"`
    Label      string    `json:"label" db:"label"`
    Restricted bool      `json:"is_restricted" db:"is_restricted"`
}

type Books []Book
// Get available books list
books := Books{}
tx := c.Value("tx").(*pop.Connection)

var err error

if !registeredAccount {
    err = tx.Where("is_restricted = false").All(&books)
} else {
    // Create an empty query
    err = tx.All(&books)
}

if err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", books)
}

// Get a specific book
book := Book{}
bookID := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
tx := c.Value("tx").(*pop.Connection)

var err error

if !registeredAccount {
    err = tx.Where("is_restricted = false AND id = ?", bookID).First(&book)
} else {
    err = tx.Find(&book, bookID)
}

if err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", book)
}
```

## The Scoped Way 作用域方式 

The scope factorizes the common part of the query:

​	作用域将查询的公共部分分解为：

```go
type Book struct {
    ID         uuid.UUID `json:"id" db:"id"`
    Label      string    `json:"label" db:"label"`
    Restricted bool      `json:"is_restricted" db:"is_restricted"`
}

type Books []Book
// restrictedScope defines a base query which shares the common constraint.
func restrictedScope(registeredAccount bool) pop.ScopeFunc {
  return func(q *pop.Query) *pop.Query {
    if !registeredAccount {
      return q
    }
    return q.Where("is_restricted = false")
  }
}
// Get available books list
books := Books{}

if err := tx.Scope(restrictedScope(registeredAccount)).All(&books); err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", books)
}

// Get a specific book
book := Book{}
bookID := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
tx := c.Value("tx").(*pop.Connection)

var err error

if err := tx.Scope(restrictedScope(registeredAccount)).Find(&book, bookID) != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", book)
}
```

See how we factorized the common restriction for each query, using the `restrictedScope` function?

​	您看到我们如何使用 `restrictedScope` 函数分解每个查询的公共限制了吗？
