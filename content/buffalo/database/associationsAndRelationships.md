+++
title = "关联和关系"
date = 2024-02-04T21:15:20+08:00
weight = 14
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/relations/]({{< ref "/buffalo/database/associationsAndRelationships" >}})

# Associations and Relationships 关联和关系 

Associations are the Pop way to define **a relation between two objects in the database**. In this chapter, you’ll learn how to define associations using struct tags; and how to manipulate them with the `Eager()` modifier.

​	关联是定义数据库中两个对象之间关系的流行方式。在本章中，您将学习如何使用结构体标签定义关联；以及如何使用 `Eager()` 修饰符操作它们。

{{< vimeo "253683926">}}

## Example 示例 

```go
type User struct {
  ID           uuid.UUID
  Email        string
  Password     string
  Books        Books     `has_many:"books" order_by:"title asc"`
  FavoriteSong Song      `has_one:"song" fk_id:"u_id"`
  Houses       Addresses `many_to_many:"users_addresses"`
}

type Book struct {
  ID      uuid.UUID
  Title   string
  Isbn    string
  User    User        `belongs_to:"user"`
  UserID  uuid.UUID
}

type Song struct {
  ID      uuid.UUID
  Title   string
  UserID  uuid.UUID   `db:"u_id"`
}

type Address struct {
  ID           uuid.UUID
  Street       string
  HouseNumber  int
}

type Books []Book
type Addresses []Address
```

## Available Struct Tags 可用的结构体标签 

Using the above [example](https://gobuffalo.io/documentation/database/relations/#example) code below is a list of available struct tags and how to use them.

​	使用上面的示例代码，下面列出了可用的结构体标签以及如何使用它们。

- `has_many`: This tag is used to describe [one-to-many](https://en.wikipedia.org/wiki/One-to-many_(data_model)) relationships in the database. In the example, `User` type defines a one-to-many relation with `Books` slice type through the use of `has_many` tag, meaning a `User` can own many `Books`. When querying to the database, Pop will load all records from the `books` table that have a column named `user_id`, or the column specified with `fk_id` that matches the `User.ID` value.

  ​	 `has_many` ：此标签用于描述数据库中的一对多关系。在示例中， `User` 类型通过使用 `has_many` 标签定义了与 `Books` 切片类型的一对多关系，这意味着 `User` 可以拥有许多 `Books` 。在查询数据库时，Pop 将加载 `books` 表中所有具有名为 `user_id` 的列的记录，或使用 `fk_id` 指定的列，该列与 `User.ID` 值匹配。

- `belongs_to`: This tag is used to describe the owner in the relationship. An owner represents a highly coupled dependency between the model and the target association field where `belongs_to` tag was defined. This tag is mostly used to indicate that model owns its “existence” to the association field with `belongs_to`. In the example above, `Book` type use `belongs_to` to indicate that it is owned by a `User` type. When querying to the database, Pop will load a record from the `users` table with `id` that matches with `Book.UserID` value.

  ​	 `belongs_to` ：此标记用于描述关系中的所有者。所有者表示模型与目标关联字段之间的高度耦合依赖关系，其中定义了 `belongs_to` 标记。此标记主要用于指示模型拥有其与 `belongs_to` 关联字段的“存在”。在上面的示例中， `Book` 类型使用 `belongs_to` 指示它归 `User` 类型所有。在查询数据库时，Pop 将从 `users` 表中加载记录，其中 `id` 与 `Book.UserID` 值匹配。

- `has_one`: This tag is used to describe [one-to-one](https://en.wikipedia.org/wiki/One-to-one_(data_model)) relationships in the database. In the example above, there is only one `FavoriteSong` within all songs records that `User` type like the most. When querying to the database, Pop will load a record from the `songs` table that have a column named `user_id`, or the column specified with `fk_id` that matches the `User.ID` field value.

  ​	 `has_one` ：此标记用于描述数据库中的一对一关系。在上面的示例中，所有歌曲记录中只有一个 `FavoriteSong` 最喜欢的 `User` 类型。在查询数据库时，Pop 将从 `songs` 表中加载一条记录，该记录具有名为 `user_id` 的列，或使用 `fk_id` 指定的列，该列与 `User.ID` 字段值匹配。

- `many_to_many`: This tag is used to describe [many-to-many](https://en.wikipedia.org/wiki/Many-to-many_(data_model)) relationships in the database. In the example above, the relationship between `User` type and `Addresses` slice type exists to indicate an `User` can own many `Houses` and a `House` can be owned by many `Users`. It is important to notice that value for `many_to_many` tag is the associative table that connects both sides in the relationship; in the example above this value is defined as `users_addresses`. When querying to the database, Pop will load all records from the `addresses` table through the associative table `users_addresses`. Table `users_addresses` **MUST** define `address_id` and `user_id` columns to match `User.ID` and `Address.ID` field values. You can also define a `fk_id` tag that will be used in the target association i.e. `addresses` table.

  ​	 `many_to_many` ：此标记用于描述数据库中的多对多关系。在上面的示例中， `User` 类型和 `Addresses` 切片类型之间的关系存在，以指示一个 `User` 可以拥有许多 `Houses` ，而一个 `House` 可以被许多 `Users` 拥有。请注意， `many_to_many` 标记的值是连接关系中双方的关联表；在上面的示例中，此值定义为 `users_addresses` 。在查询数据库时，Pop 将通过关联表 `users_addresses` 从 `addresses` 表加载所有记录。表 `users_addresses` 必须定义 `address_id` 和 `user_id` 列以匹配 `User.ID` 和 `Address.ID` 字段值。您还可以定义一个 `fk_id` 标记，该标记将用于目标关联中，即 `addresses` 表。

- `fk_id`: This tag can be used to define the column name in the target association that matches model ID. In the example above, `Song` has a column named `u_id` that references the id of the `users` table. When loading `FavoriteSong`, `u_id` column will be used instead of `user_id`.

  ​	 `fk_id` ：此标记可用于定义目标关联中与模型 ID 匹配的列名。在上面的示例中， `Song` 有一个名为 `u_id` 的列，该列引用 `users` 表的 ID。在加载 `FavoriteSong` 时，将使用 `u_id` 列而不是 `user_id` 。

- `order_by`: This tag can be used in combination with `has_many` and `many_to_many` tags to indicate the order for the association when loading. The format to use is `order_by:"<column_name> <asc | desc>"`

  ​	 `order_by` ：此标签可与 `has_many` 和 `many_to_many` 标签结合使用，以指示加载时的关联顺序。要使用的格式为 `order_by:"<column_name> <asc | desc>"`

## Loading Associations 加载关联 

Pop currently provides two modes for loading associations; each mode will affect the way pop loads associations and queries to the database.

​	Pop 目前提供两种加载关联的模式；每种模式都会影响 pop 加载关联和查询数据库的方式。

[Eager](https://gobuffalo.io/documentation/database/relations/#eager-mode). Default mode. By enabling this mode, pop will perform “n” queries for every association defined in the model. This means more hits to the database in order to not affect memory use.

​	Eager。默认模式。通过启用此模式，pop 将对模型中定义的每个关联执行“n”个查询。这意味着更多地访问数据库，以便不影响内存使用。

[EagerPreload](https://gobuffalo.io/documentation/database/relations/#eagerpreload-mode). Optional mode. By enabling this mode, pop will perform one query for every association defined in the model. This mode will hit the database with a reduced frequency by sacrifing more memory space.

​	EagerPreload。可选模式。通过启用此模式，pop 将对模型中定义的每个关联执行一个查询。此模式将以牺牲更多内存空间为代价，减少访问数据库的频率。

- `pop.SetEagerMode`: Pop allows enabling any of these modes globally which will affect **ALL** queries handle performance. Use `EagerDefault` or `EagerPreload` as parameter to activate any of these modes.

  ​	 `pop.SetEagerMode` ：Pop 允许全局启用这些模式中的任何一种，这将影响所有查询处理性能。使用 `EagerDefault` 或 `EagerPreload` 作为参数来激活这些模式中的任何一种。

- `tx.EagerPreload | q.EagerPreload`: Pop allows developers to take control in which situations they want Pop to perform any of these modes when necessary. This method will activate `EagerPreload` mode only for the query in action.

  ​	 `tx.EagerPreload | q.EagerPreload` ：Pop 允许开发人员控制在必要时希望 Pop 执行这些模式中的哪一种的情况。此方法仅对正在执行的查询激活 `EagerPreload` 模式。

- `tx.Eager | q.Eager`: Pop allows developers to take control in which situations they want Pop to perform any of these modes when necessary. This method will activate `Eager` mode only for the query in action.

  ​	 `tx.Eager | q.Eager` ：Pop 允许开发人员控制在必要时希望 Pop 在哪些情况下执行这些模式之一。此方法仅对操作中的查询激活 `Eager` 模式。

## Eager Mode 急切模式 

The [`pop.Connection.Eager()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Eager) method tells Pop to load the associations for a model once that model is loaded from the database. This mode will perform “n” queries for every association defined in the model.

​	 `pop.Connection.Eager()` 方法告诉 Pop 在从数据库加载模型后立即加载该模型的关联。此模式将对模型中定义的每个关联执行“n”个查询。

```go
for i := 0; i < 3; i++ {
  user := User{ID: i + 1}
  tx.Create(&user)
}

for i := 0; i < 3; i++ {
  book := Book{UserID: i +1}
  tx.Create(&book)
}
u := Users{}
err := tx.Eager().All(&u)  // loads all associations for every user registered, i.e Books, Houses and FavoriteSong
```

`Eager` mode will:

​	 `Eager` 模式将：

1. Load all users.
   加载所有用户。

```text
 SELECT * FROM users;
```

1. Iterate on every user and load its associations:
   迭代每个用户并加载其关联：

```erb
 SELECT * FROM books WHERE user_id=1)
 SELECT * FROM books WHERE user_id=2)
 SELECT * FROM books WHERE user_id=3)
```

## EagerPreload Mode EagerPreload 模式 

The [`pop.Connection.EagerPreload()`](https://github.com/gobuffalo/pop/pull/146/files#diff-f49e947ec94f65964b0845af2b62845aR180) method tells Pop to load the associations for a model once that model is loaded from the database. This mode will hit the database with a reduced frequency by sacrifing more memory space.

​	 `pop.Connection.EagerPreload()` 方法告诉 Pop 在从数据库加载模型后立即加载该模型的关联。此模式将通过牺牲更多内存空间来减少与数据库的交互频率。

```go
for i := 0; i < 3; i++ {
  user := User{ID: i + 1}
  tx.Create(&user)
}

for i := 0; i < 3; i++ {
  book := Book{UserID: i +1}
  tx.Create(&book)
}
u := Users{}
err := tx.EagerPreload().All(&u)  // loads all associations for every user registered, i.e Books, Houses and FavoriteSong
```

`EagerPreload` mode will:

​	 `EagerPreload` 模式将：

1. Load all users.
   加载所有用户。

```erb
 SELECT * FROM users;
```

1. Load associations for all users in one single query.
   在单个查询中加载所有用户的关联。

```erb
  SELECT * FROM books WHERE user_id IN (1,2,3))
```

## Load Specific Associations 加载特定关联 

By default `Eager` and `EagerPreload` will load all the assigned associations for the model. To specify which associations should be loaded you can pass in the names of those fields to the `Eager` or `EagerPreload` methods and only those associations will be loaded.

​	默认情况下， `Eager` 和 `EagerPreload` 将加载模型的所有已分配关联。要指定应加载哪些关联，可以将这些字段的名称传递给 `Eager` 或 `EagerPreload` 方法，并且只会加载这些关联。

```go
err  = tx.Eager("Books").Where("name = 'Mark'").All(&u) // load only Books association for user with name 'Mark'.
// OR
err  = tx.EagerPreload("Books").Where("name = 'Mark'").All(&u) // load only Books association for user with name 'Mark'.
```

Pop also allows you to eager load nested associations by using the `.` character to concatenate them. Take a look at the example below.

​	Pop 还允许您使用 `.` 字符连接嵌套关联以急切加载它们。请看下面的例子。

```go
// will load all Books for u and for every Book will load the user which will be the same as u.
tx.Eager("Books.User").First(&u)
// OR
tx.EagerPreload("Books.User").First(&u)
// will load all Books for u and for every Book will load all Writers and for every writer will load the Book association.
tx.Eager("Books.Writers.Book").First(&u)
// OR
tx.EagerPreload("Books.Writers.Book").First(&u)
// will load all Books for u and for every Book will load all Writers. And Also it will load the favorite song for user.
tx.Eager("Books.Writers").Eager("FavoriteSong").First(&u)
// OR
tx.EagerPreload("Books.Writers").EagerPreload("FavoriteSong").First(&u)
```

## Loading Associations for an Existing Model 加载现有模型的关联 

The [`pop.Connection.Load()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Load) method takes a model struct, that has already been populated from the database, and an optional list of associations to load.

​	 `pop.Connection.Load()` 方法采用一个模型结构，该结构已从数据库中填充，以及一个要加载的关联的可选列表。

```go
tx.Load(&u) // load all associations for user, i.e Books, Houses and FavoriteSong
tx.Load(&u, "Books") // load only the Books associations for user
```

The `Load` method will not retrieve the `User` from the database, only its associations.

​	 `Load` 方法不会从数据库中检索 `User` ，只会检索其关联。

## Flat Nested Creation 扁平嵌套创建 

Pop allows you to create the models and their associations with other models in one step by default. You no longer need to create every association separately anymore. Pop will even create join table records for `many_to_many` associations.

​	默认情况下，Pop 允许您一步创建模型及其与其他模型的关联。您不再需要分别创建每个关联。Pop 甚至会为 `many_to_many` 关联创建联接表记录。

Assuming the following pieces of pseudo-code:

​	假设以下伪代码片段：

```go
book := Book{Title: "Pop Book", Description: "Pop Book", Isbn: "PB1"}
tx.Create(&book)
song := Song{Title: "Don't know the title"}
tx.Create(&song)
addr := Address{HouseNumber: 1, Street: "Golang"}
tx.Create(&addr)

user := User{
  Name: "Mark Bates",
  Books: Books{Book{ID: book.ID}},
  FavoriteSong: song,
  Houses: Addresses{
    addr,
  },
}
err := tx.Create(&user)
```

1. It will notice `Books` is a `has_many` association and it will realize that to actually update each book it will need to get the `User ID` first. So, it proceeds to store first `User` data so it can retrieve an **ID** and then use that ID to fill `UserID` field in every `Book` in `Books`. It updates all affected books in the database using their `ID`s to target them.

   ​	它会注意到 `Books` 是一个 `has_many` 关联，并且它会意识到要实际更新每本书，它需要首先获取 `User ID` 。因此，它会继续首先存储 `User` 数据，以便它可以检索一个 ID，然后使用该 ID 来填充 `UserID` 字段中的每个 `Book` 在 `Books` 中。它使用它们的 `ID` s 来定位它们，从而更新数据库中所有受影响的书籍。

2. `FavoriteSong` is a `has_one` association and it uses same logic described in `has_many` association. Since `User` data was previously saved before updating all affected books, it already knows that `User` has got an `ID` so it fills its `UserID` field with that value and `FavoriteSong` is then updated in the database.

   ​	 `FavoriteSong` 是一个 `has_one` 关联，并且它使用在 `has_many` 关联中描述的相同逻辑。由于在更新所有受影响的书籍之前已经保存了 `User` 数据，因此它已经知道 `User` 已经获得了 `ID` ，因此它用该值填充其 `UserID` 字段，然后在数据库中更新 `FavoriteSong` 。

3. `Houses` in this example is a `many_to_many` relationship and it will have to deal with two tables in this case: `users` and `addresses`. Because `User` was already stored, it already has its `ID`. It will then use the `ID`s passed with the `Addresses` to create the coresponding entries in the join table.

   ​	本例中的 `Houses` 是一个 `many_to_many` 关系，并且在这种情况下它将必须处理两个表： `users` 和 `addresses` 。因为 `User` 已经存储，所以它已经有了它的 `ID` 。然后，它将使用与 `Addresses` 一起传递的 `ID` s 来在联接表中创建相应的条目。

For a `belongs_to` association like shown in the example below, it fills its `UserID` field before being saved in the database.

​	对于如下例所示的 `belongs_to` 关联，它会在保存到数据库之前填充其 `UserID` 字段。

```go
book := Book{
   Title:      "Pop Book",
   Description: "Pop Book",
   Isbn:        "PB1",
   User: user,
}
tx.Create(&book)
```

## Eager Creation 急切创建 

Pop also allows you to create models and embed the creation of their associations in one step as well.

​	Pop 还允许您创建模型并在一个步骤中嵌入其关联的创建。

Assuming the following pieces of pseudo-code:

​	假设以下伪代码片段：

```go
user := User{
  Name: "Mark Bates",
  Books: Books{{Title: "Pop Book", Description: "Pop Book", Isbn: "PB1"}},
  FavoriteSong: Song{Title: "Don't know the title"},
  Houses: Addresses{
    Address{HouseNumber: 1, Street: "Golang"},
  },
}
err := tx.Eager().Create(&user)
```

1. It will notice `Books` is a `has_many` association and it will realize that to actually store every book it will need to get the `User ID` first. So, it proceeds to first store/create the `User` data so it can retrieve an **ID** and then use that ID to fill the `UserID` field in every `Book` in `Books`. Later it stores all books in the database.

   ​	它会注意到 `Books` 是一个 `has_many` 关联，并且它会意识到要实际存储每本书，它需要先获取 `User ID` 。因此，它会首先存储/创建 `User` 数据，以便它可以检索一个 ID，然后使用该 ID 来填充 `Books` 中的每个 `Book` 中的 `UserID` 字段。稍后，它会将所有书籍存储在数据库中。

2. `FavoriteSong` is a `has_one` association and it uses same logic described in the `has_many` association. Since `User` data was previously saved before creating all books, it already knows that `User` has got an `ID` so it fills its `UserID` field with that value and `FavoriteSong` is then stored in the database.

   ​	 `FavoriteSong` 是一个 `has_one` 关联，并且它使用 `has_many` 关联中描述的相同逻辑。由于在创建所有书籍之前已经保存了 `User` 数据，因此它已经知道 `User` 有一个 `ID` ，因此它用该值填充其 `UserID` 字段，然后将 `FavoriteSong` 存储在数据库中。

3. `Houses` in this example is a `many_to_many` relationship and it will have to deal with two tables, in this case: `users` and `addresses`. It will need to store all addresses first in the `addresses` table before saving them in the many to many(join) table. Because `User` was already stored, it already has an `ID`. * This is a special case to deal with, since this behavior is different from all other associations, it is solved by implementing the `AssociationCreatableStatement` interface, all other associations by default implement the `AssociationCreatable` interface.

   ​	 `Houses` 在此示例中是一个 `many_to_many` 关系，它将处理两个表，在本例中为： `users` 和 `addresses` 。它需要先将所有地址存储在 `addresses` 表中，然后再将它们保存在多对多（联接）表中。因为 `User` 已存储，所以它已具有 `ID` 。* 这是一个需要处理的特殊情况，因为此行为与所有其他关联不同，通过实现 `AssociationCreatableStatement` 接口来解决，默认情况下所有其他关联都实现 `AssociationCreatable` 接口。

For a `belongs_to` association like shown in the example below, it will need to first create the `User` to retrieve its **ID** value and then fill its `UserID` field before being saved in the database.

​	对于如下例所示的 `belongs_to` 关联，它需要首先创建 `User` 以检索其 ID 值，然后在将其保存在数据库中之前填充其 `UserID` 字段。

```go
book := Book{
   Title:      "Pop Book",
   Description: "Pop Book",
   Isbn:        "PB1",
   User: User{
        Name: nulls.NewString("Larry"),
   },
}
tx.Eager().Create(&book)
```

In the case where you feed the eager create with associated models that already exist, it will, instead of creating duplicates of them or updating the contents of them, simply create/update the associations with them.

​	在您使用已存在的关联模型提供急切创建的情况下，它不会创建它们的副本或更新它们的内容，而是简单地创建/更新与它们的关联。

## Next Steps 后续步骤 

- [One to one relations
  一对一关系]({{< ref "/buffalo/database/oneToOneAssociations" >}})
- [One to many relations
  一对多关系]({{< ref "/buffalo/database/oneToManyAssociations" >}})
