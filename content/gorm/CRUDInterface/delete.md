+++
title = "删除"
date = 2023-10-28T14:26:18+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/delete.html](https://gorm.io/docs/delete.html)

## 删除记录 Delete a Record

When deleting a record, the deleted value needs to have primary key or it will trigger a [Batch Delete](https://gorm.io/docs/delete.html#batch_delete), for example:

​	当删除记录时，被删除的值需要有主键，否则会触发[批量删除](https://gorm.io/docs/delete.html#batch_delete)，例如：

``` go
//  Email的ID是`10` Email's ID is `10`
db.Delete(&email)
// DELETE from emails where id = 10;

// 使用附加条件进行删除 Delete with additional conditions
db.Where("name = ?", "jinzhu").Delete(&email)
// DELETE from emails where id = 10 AND name = "jinzhu";
```

## 使用主键删除 Delete with primary key

GORM allows to delete objects using primary key(s) with inline condition, it works with numbers, check out [Query Inline Conditions](https://gorm.io/docs/query.html#inline_conditions) for details

​	GORM允许使用内联条件和主键来删除对象，它适用于数字，查看[查询内联条件](https://gorm.io/docs/query.html#inline_conditions)以获取详细信息

``` go
db.Delete(&User{}, 10)
// DELETE FROM users WHERE id = 10;

db.Delete(&User{}, "10")
// DELETE FROM users WHERE id = 10;

db.Delete(&users, []int{1,2,3})
// DELETE FROM users WHERE id IN (1,2,3);
```

## 删除钩子 Delete Hooks

GORM allows hooks `BeforeDelete`, `AfterDelete`, those methods will be called when deleting a record, refer [Hooks](https://gorm.io/docs/hooks.html) for details

​	GORM允许在删除记录时调用钩子函数`BeforeDelete`，`AfterDelete`，这些方法将在删除记录时被调用，参考[钩子](https://gorm.io/docs/hooks.html)以获取详细信息。

``` go
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
  if u.Role == "admin" {
    return errors.New("admin user not allowed to delete")
  }
  return
}
```

## 批量删除 Batch Delete

The specified value has no primary value, GORM will perform a batch delete, it will delete all matched records

​	如果指定的值没有主键，GORM将执行批量删除，它将删除所有匹配的记录

``` go
db.Where("email LIKE ?", "%jinzhu%").Delete(&Email{})
// DELETE from emails where email LIKE "%jinzhu%";

db.Delete(&Email{}, "email LIKE ?", "%jinzhu%")
// DELETE from emails where email LIKE "%jinzhu%";
```

To efficiently delete large number of records, pass a slice with primary keys to the `Delete` method.

​	要高效地删除大量记录，请将具有主键的切片传递给`Delete`方法。

``` go
var users = []User{{ID: 1}, {ID: 2}, {ID: 3}}
db.Delete(&users)
// DELETE FROM users WHERE id IN (1,2,3);

db.Delete(&users, "name LIKE ?", "%jinzhu%")
// DELETE FROM users WHERE name LIKE "%jinzhu%" AND id IN (1,2,3); 
```

### 阻止全局删除 Block Global Delete

If you perform a batch delete without any conditions, GORM WON’T run it, and will return `ErrMissingWhereClause` error

​	如果你执行一个不带任何条件的批量删除，GORM不会运行它，并返回`ErrMissingWhereClause`错误

You have to use some conditions or use raw SQL or enable `AllowGlobalUpdate` mode, for example:

​	你必须使用一些条件或使用原始SQL或启用`AllowGlobalUpdate`模式，例如：

``` go
db.Delete(&User{}).Error // gorm.ErrMissingWhereClause

db.Delete(&[]User{{Name: "jinzhu1"}, {Name: "jinzhu2"}}).Error // gorm.ErrMissingWhereClause

db.Where("1 = 1").Delete(&User{})
// DELETE FROM `users` WHERE 1=1

db.Exec("DELETE FROM users")
// DELETE FROM users

db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
// DELETE FROM users
```

### 从已删除的行中返回数据 Returning Data From Deleted Rows

Return deleted data, only works for database support Returning, for example:

​	返回已删除的数据，仅适用于数据库支持Returning的情况，例如：

``` go
// return all columns
var users []User
DB.Clauses(clause.Returning{}).Where("role = ?", "admin").Delete(&users)
// DELETE FROM `users` WHERE role = "admin" RETURNING *
// users => []User{{ID: 1, Name: "jinzhu", Role: "admin", Salary: 100}, {ID: 2, Name: "jinzhu.2", Role: "admin", Salary: 1000}}

// return specified columns
DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "salary"}}}).Where("role = ?", "admin").Delete(&users)
// DELETE FROM `users` WHERE role = "admin" RETURNING `name`, `salary`
// users => []User{{ID: 0, Name: "jinzhu", Role: "", Salary: 100}, {ID: 0, Name: "jinzhu.2", Role: "", Salary: 1000}}
```

## 软删除 Soft Delete

If your model includes a `gorm.DeletedAt` field (which is included in `gorm.Model`), it will get soft delete ability automatically!

​	如果模型包含`gorm.DeletedAt`字段（包含在`gorm.Model`中），它将自动获得软删除能力！

When calling `Delete`, the record WON’T be removed from the database, but GORM will set the `DeletedAt`‘s value to the current time, and the data is not findable with normal Query methods anymore.

​	当调用`Delete`时，记录将不会被从数据库中删除，但GORM将设置`DeletedAt`的值当前时间，并且通过正常的查询方法无法找到数据。

``` go
// user's ID is `111`
db.Delete(&user)
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

// Batch Delete
db.Where("age = ?", 20).Delete(&User{})
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// Soft deleted records will be ignored when querying
db.Where("age = 20").Find(&user)
// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
```

If you don’t want to include `gorm.Model`, you can enable the soft delete feature like:

​	如果你不想包含`gorm.Model`，你可以像下面这样启用软删除功能：

``` go
type User struct {
  ID      int
  Deleted gorm.DeletedAt
  Name    string
}
```

### 查找已软删除的记录 Find soft deleted records

You can find soft deleted records with `Unscoped`

​	你可以使用`Unscoped`查找软删除的记录

``` go
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;
```

### 永久删除 Delete permanently

You can delete matched records permanently with `Unscoped`

​	你可以永久删除匹配的记录，使用`Unscoped`

``` go
db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;
```

### 删除标志 Delete Flag

By default, `gorm.Model` uses `*time.Time` as the value for the `DeletedAt` field, and it provides other data formats support with plugin `gorm.io/plugin/soft_delete`

​	默认情况下，`gorm.Model`使用`*time.Time`作为`DeletedAt`字段的值，并提供了其他数据格式支持，使用插件`gorm.io/plugin/soft_delete`

> **INFO** when creating unique composite index for the DeletedAt field, you must use other data format like unix second/flag with plugin `gorm.io/plugin/soft_delete`‘s help, e.g:
>
> **INFO** 创建唯一复合索引以`DeletedAt`字段时，您必须使用其他数据格式，如使用插件`gorm.io/plugin/soft_delete`的帮助，例如：
>
> ``` go
> import "gorm.io/plugin/soft_delete"
> 
> type User struct {
> ID        uint
> Name      string                `gorm:"uniqueIndex:udx_name"`
> DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name"`
> }
> ```

#### Unix Second

Use unix second as delete flag

​	使用 Unix 秒数作为删除标记

``` go
import "gorm.io/plugin/soft_delete"

type User struct {
  ID        uint
  Name      string
  DeletedAt soft_delete.DeletedAt
}

// 查询 Query
SELECT * FROM users WHERE deleted_at = 0;

// 删除 Delete
UPDATE users SET deleted_at = /* current unix second */ WHERE ID = 1;
```

You can also specify to use `milli` or `nano` seconds as the value, for example:

​	你也可以指定使用 `milli` 或 `nano` 秒作为值，例如：

``` go
type User struct {
  ID    uint
  Name  string
  DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli"`
  // DeletedAt soft_delete.DeletedAt `gorm:"softDelete:nano"`
}

// Query
SELECT * FROM users WHERE deleted_at = 0;

// Delete
UPDATE users SET deleted_at = /* current unix milli second or nano second */ WHERE ID = 1;
```

#### 使用 `1` / `0` 作为删除标记 Use `1` / `0` AS Delete Flag

``` go
import "gorm.io/plugin/soft_delete"

type User struct {
  ID    uint
  Name  string
  IsDel soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

// 查询 Query
SELECT * FROM users WHERE is_del = 0;

// 删除 Delete
UPDATE users SET is_del = 1 WHERE ID = 1;
```

#### 混合模式 Mixed Mode

Mixed mode can use `0`, `1` or unix seconds to mark data as deleted or not, and save the deleted time at the same time.

​	混合模式可以使用 `0`、`1` 或 Unix 秒数来标记数据是否已删除，并同时保存删除时间。

``` go
type User struct {
  ID        uint
  Name      string
  DeletedAt time.Time
  IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"` // use `1` `0`
  // IsDel     soft_delete.DeletedAt `gorm:"softDelete:,DeletedAtField:DeletedAt"` // use `unix second`
  // IsDel     soft_delete.DeletedAt `gorm:"softDelete:nano,DeletedAtField:DeletedAt"` // use `unix nano second`
}

// 查询 Query
SELECT * FROM users WHERE is_del = 0;

// 删除 Delete
UPDATE users SET is_del = 1, deleted_at = /* current unix second */ WHERE ID = 1;
```