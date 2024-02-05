+++
title = "更新"
date = 2023-10-28T14:26:11+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/update.html](https://gorm.io/docs/update.html)

## 保存所有字段 Save All Fields

`Save` will save all fields when performing the Updating SQL

​	`Save` 在执行更新 SQL 时会保存所有字段。

``` go
db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)
// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
```

`Save` is a combination function. If save value does not contain primary key, it will execute `Create`, otherwise it will execute `Update` (with all fields).

​	`Save` 是一个组合函数。如果保存值不包含主键，它将执行 `Create`，否则它将执行 `Update`（带有所有字段）。

``` go
db.Save(&User{Name: "jinzhu", Age: 100})
// INSERT INTO `users` (`name`,`age`,`birthday`,`update_at`) VALUES ("jinzhu",100,"0000-00-00 00:00:00","0000-00-00 00:00:00")

db.Save(&User{ID: 1, Name: "jinzhu", Age: 100})
// UPDATE `users` SET `name`="jinzhu",`age`=100,`birthday`="0000-00-00 00:00:00",`update_at`="0000-00-00 00:00:00" WHERE `id` = 1
```

> **NOTE** Don’t use `Save` with `Model`, it’s an **Undefined Behavior**.
>
> **注意** 不要使用 `Save` 与 `Model`，它是一个未定义的行为。

## 更新单个列 Update single column

When updating a single column with `Update`, it needs to have any conditions or it will raise error `ErrMissingWhereClause`, checkout [Block Global Updates](https://gorm.io/docs/update.html#block_global_updates) for details.
When using the `Model` method and its value has a primary value, the primary key will be used to build the condition, for example:

​	当使用 `Update` 更新单个列时，需要具有任何条件或否则将引发错误 `ErrMissingWhereClause`，请参阅 [Block Global Updates](https://gorm.io/docs/update.html#block_global_updates) 以获取详细信息。当使用 `Model` 方法且其值具有主键时，主键将用于构建条件，例如：

``` go
// 带条件的更新 Update with conditions
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

// 用户 ID 为 `111`： User's ID is `111`:
db.Model(&user).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// 带条件和模型值的更新 Update with conditions and model value
db.Model(&user).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
```

## 更新多个列 Updates multiple columns

`Updates` supports updating with `struct` or `map[string]interface{}`, when updating with `struct` it will only update non-zero fields by default

​	`Updates` 支持使用 `struct` 或 `map[string]interface{}` 进行更新，当使用 `struct` 时，默认只更新非零字段。

``` go
// 使用 `struct` 更新属性，将只更新非零字段 Update attributes with `struct`, will only update non-zero fields
db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// 使用 `map` 更新属性 Update attributes with `map`
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
```

> **NOTE** When updating with struct, GORM will only update non-zero fields. You might want to use `map` to update attributes or use `Select` to specify fields to update
>
> **注意** 当使用 `struct` 进行更新时，GORM 只会更新非零字段。你可能想要使用 `map` 来更新属性，或者使用 `Select` 来指定要更新的字段。

## 更新选定字段 Update Selected Fields

If you want to update selected fields or ignore some fields when updating, you can use `Select`, `Omit`

​	如果你想在更新时选择特定的字段或在更新时忽略某些字段，你可以使用 `Select`、`Omit`。

``` go
// 使用 Map 选择字段 Select with Map
// 用户 ID 是 `111`：User's ID is `111`:
db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello' WHERE id=111;

db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// 使用 Struct（选择零值字段） Select with Struct (select zero value fields)
db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
// UPDATE users SET name='new_name', age=0 WHERE id=111;

// 选择所有字段（选择所有字段，包括零值字段） Select all fields (select all fields include zero value fields)
db.Model(&user).Select("*").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})

// 选择所有字段但省略 Role（选择所有字段，包括零值字段） Select all fields but omit Role (select all fields include zero value fields)
db.Model(&user).Select("*").Omit("Role").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})
```

## 更新钩子 Update Hooks

GORM allows the hooks `BeforeSave`, `BeforeUpdate`, `AfterSave`, `AfterUpdate`. Those methods will be called when updating a record, refer [Hooks](https://gorm.io/docs/hooks.html) for details

​	GORM 允许在更新记录时调用钩子方法 `BeforeSave`、`BeforeUpdate`、`AfterSave`、`AfterUpdate`。有关详细信息，请参阅 [Hooks]({{< ref "/gorm/Tutorials/hooks">}})。

``` go
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
  if u.Role == "admin" {
    return errors.New("admin user not allowed to update")
  }
  return
}
```

## 批量更新 Batch Updates

If we haven’t specified a record having a primary key value with `Model`, GORM will perform a batch update

​	如果我们没有在 `Model` 中指定具有主键值的记录，GORM 将执行批量更新。

``` go
// 使用 struct 更新 Update with struct
db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

// 使用 map 更新 Update with map
db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
```

### 阻止全局更新 Block Global Updates

If you perform a batch update without any conditions, GORM WON’T run it and will return `ErrMissingWhereClause` error by default

​	如果你执行了一个没有任何条件的批量更新，GORM 默认不会运行它，并会返回 `ErrMissingWhereClause` 错误。

You have to use some conditions or use raw SQL or enable the `AllowGlobalUpdate` mode, for example:

​	你必须使用一些条件，或者使用原始 SQL，或者启用 `AllowGlobalUpdate` 模式，例如：

``` go
db.Model(&User{}).Update("name", "jinzhu").Error // gorm.ErrMissingWhereClause

db.Model(&User{}).Where("1 = 1").Update("name", "jinzhu")
// UPDATE users SET `name` = "jinzhu" WHERE 1=1

db.Exec("UPDATE users SET name = ?", "jinzhu")
// UPDATE users SET name = "jinzhu"

db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("name", "jinzhu")
// UPDATE users SET `name` = "jinzhu"
```

### 更新的记录数 Updated Records Count

Get the number of rows affected by a update

​	获取更新操作影响的行数

``` go
// 使用 `RowsAffected` 获取更新记录数 Get updated records count with `RowsAffected`
result := db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

result.RowsAffected // 返回更新记录数 returns updated records count
result.Error        // 返回更新错误 returns updating error
```

## 高级 Advanced

### 使用 SQL 表达式进行更新 Update with SQL Expression

GORM allows updating a column with a SQL expression, e.g:

​	GORM 允许使用 SQL 表达式更新列，例如：

``` go
// product's ID is `3`
db.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

db.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

db.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3;

db.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3 AND quantity > 1;
```

And GORM also allows updating with SQL Expression/Context Valuer with [Customized Data Types](https://gorm.io/docs/data_types.html#gorm_valuer_interface), e.g:

​	GORM 还允许使用 SQL 表达式/上下文值器与 [自定义数据类型]({{< ref "/gorm/Tutorials/customizeDataTypes#gormvaluerinterface">}}) 进行更新，例如：

``` go
// 从自定义数据类型创建 Create from customized data type
type Location struct {
  X, Y int
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
  return clause.Expr{
    SQL:  "ST_PointFromText(?)",
    Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
  }
}

db.Model(&User{ID: 1}).Updates(User{
  Name:  "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// UPDATE `user_with_points` SET `name`="jinzhu",`location`=ST_PointFromText("POINT(100 100)") WHERE `id` = 1
```

### 使用子查询进行更新 Update from SubQuery

Update a table by using SubQuery

​	使用子查询更新表

``` go
db.Model(&user).Update("company_name", db.Model(&Company{}).Select("name").Where("companies.id = users.company_id"))
// UPDATE "users" SET "company_name" = (SELECT name FROM companies WHERE companies.id = users.company_id);

db.Table("users as u").Where("name = ?", "jinzhu").Update("company_name", db.Table("companies as c").Select("name").Where("c.id = u.company_id"))

db.Table("users as u").Where("name = ?", "jinzhu").Updates(map[string]interface{}{"company_name": db.Table("companies as c").Select("name").Where("c.id = u.company_id")})
```

### 不使用Hooks/时间追踪 Without Hooks/Time Tracking

If you want to skip `Hooks` methods and don’t track the update time when updating, you can use `UpdateColumn`, `UpdateColumns`, it works like `Update`, `Updates`

​	如果你想跳过`Hooks`方法并在更新时不追踪更新时间，可以使用`UpdateColumn`、`UpdateColumns`，它们的行为类似于`Update`、`Updates`。

``` go
// 更新单个列 Update single column
db.Model(&user).UpdateColumn("name", "hello")
// UPDATE users SET name='hello' WHERE id = 111;

// 更新多个列 Update multiple columns
db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE id = 111;

// 更新选定的列 Update selected columns
db.Model(&user).Select("name", "age").UpdateColumns(User{Name: "hello", Age: 0})
// UPDATE users SET name='hello', age=0 WHERE id = 111;   
```

### 从修改的行返回数据 Returning Data From Modified Rows

Returning changed data only works for databases which support Returning, for example:

​	仅适用于支持Returning的数据库，例如：

``` go
// 返回所有列 return all columns
var users []User
db.Model(&users).Clauses(clause.Returning{}).Where("role = ?", "admin").Update("salary", gorm.Expr("salary * ?", 2))
// UPDATE `users` SET `salary`=salary * 2,`updated_at`="2021-10-28 17:37:23.19" WHERE role = "admin" RETURNING *
// users => []User{{ID: 1, Name: "jinzhu", Role: "admin", Salary: 100}, {ID: 2, Name: "jinzhu.2", Role: "admin", Salary: 1000}}

// 返回指定列 return specified columns
db.Model(&users).Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "salary"}}}).Where("role = ?", "admin").Update("salary", gorm.Expr("salary * ?", 2))
// UPDATE `users` SET `salary`=salary * 2,`updated_at`="2021-10-28 17:37:23.19" WHERE role = "admin" RETURNING `name`, `salary`
// users => []User{{ID: 0, Name: "jinzhu", Role: "", Salary: 100}, {ID: 0, Name: "jinzhu.2", Role: "", Salary: 1000}}
```

### 检查字段是否已更改？Check Field has changed?

GORM provides the `Changed` method which could be used in **Before Update Hooks**, it will return whether the field has changed or not.

​	GORM提供了`Changed`方法，可以在**更新前Hooks**中使用，它将返回字段是否已更改。

The `Changed` method only works with methods `Update`, `Updates`, and it only checks if the updating value from `Update` / `Updates` equals the model value. It will return true if it is changed and not omitted

​	`Changed`方法仅在`Update`、`Updates`方法中使用，它只检查从`Update` / `Updates`更新的值是否等于模型值。如果已更改且未省略，则返回`true`。

``` go
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
  // if Role changed
  if tx.Statement.Changed("Role") {
    return errors.New("role not allowed to change")
  }

  if tx.Statement.Changed("Name", "Admin") { // if Name or Role changed
    tx.Statement.SetColumn("Age", 18)
  }

  // if any fields changed
  if tx.Statement.Changed() {
    tx.Statement.SetColumn("RefreshedAt", time.Now())
  }
  return nil
}

db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(map[string]interface{"name": "jinzhu2"})
// Changed("Name") => true
db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(map[string]interface{"name": "jinzhu"})
// Changed("Name") => false, `Name` not changed
db.Model(&User{ID: 1, Name: "jinzhu"}).Select("Admin").Updates(map[string]interface{
  "name": "jinzhu2", "admin": false,
})
// Changed("Name") => false, `Name` not selected to update

db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(User{Name: "jinzhu2"})
// Changed("Name") => true
db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(User{Name: "jinzhu"})
// Changed("Name") => false, `Name` not changed
db.Model(&User{ID: 1, Name: "jinzhu"}).Select("Admin").Updates(User{Name: "jinzhu2"})
// Changed("Name") => false, `Name` not selected to update
```

### 更改更新值 Change Updating Values

To change updating values in Before Hooks, you should use `SetColumn` unless it is a full update with `Save`, for example:

​	要在Before Hooks中更改更新值，您应该使用`SetColumn`，除非它是完整的更新（例如：`Save`），否则不要使用它：

``` go
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
  if pw, err := bcrypt.GenerateFromPassword(user.Password, 0); err == nil {
    tx.Statement.SetColumn("EncryptedPassword", pw)
  }

  if tx.Statement.Changed("Code") {
    user.Age += 20
    tx.Statement.SetColumn("Age", user.Age)
  }
}

db.Model(&user).Update("Name", "jinzhu")
```