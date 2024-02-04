+++
title = "CRUD 操作"
date = 2024-02-04T10:00:25+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/model/object/]({{< ref "/beego/mvcIntroduction/models/crudOperations" >}})

# CRUD Operations -  CRUD 操作



## CRUD of Object 对象的 CRUD

If the value of the primary key is already known, `Read`, `Insert`, `Update`, `Delete` can be used to manipulate the object.

​	如果主键的值已知， `Read` 、 `Insert` 、 `Update` 、 `Delete` 可用于操作对象。

```go
o := orm.NewOrm()
user := new(User)
user.Name = "slene"

fmt.Println(o.Insert(user))

user.Name = "Your"
fmt.Println(o.Update(user))
fmt.Println(o.Read(user))
fmt.Println(o.Delete(user))
```

To query the object by conditions see [Query in advance](https://beego.wiki/docs/mvc/model/object/query.md#all)

​	要按条件查询对象，请参阅预先查询

## Read 读取

```go
o := orm.NewOrm()
user := User{Id: 1}

err := o.Read(&user)

if err == orm.ErrNoRows {
	fmt.Println("No result found.")
} else if err == orm.ErrMissPK {
	fmt.Println("No primary key found.")
} else {
	fmt.Println(user.Id, user.Name)
}
```

Read uses primary key by default. But it can use other fields as well:

​	读取默认情况下使用主键。但它也可以使用其他字段：

```go
user := User{Name: "slene"}
err := o.Read(&user, "Name")
...
```

Other fields of the object are set to the default value according to the field type.

​	对象的其它字段根据字段类型设置为默认值。

For detailed single object query, see [One](https://beego.wiki/docs/mvc/model/object/query.md#one)

​	有关详细的单个对象查询，请参阅一个

## ReadOrCreate 读取或创建

Try to read a row from the database, or insert one if it doesn’t exist.

​	尝试从数据库中读取一行，或在不存在时插入一行。

At least one condition field must be supplied, multiple condition fields are also supported.

​	必须提供至少一个条件字段，也支持多个条件字段。

```go
o := orm.NewOrm()
user := User{Name: "slene"}
// Three return values：Is Created，Object Id，Error
if created, id, err := o.ReadOrCreate(&user, "Name"); err == nil {
	if created {
		fmt.Println("New Insert an object. Id:", id)
	} else {
		fmt.Println("Get an object. Id:", id)
	}
}
```

## Insert 插入

The first return value is auto inc Id value.

​	第一个返回值是自动增长的 Id 值。

```go
o := orm.NewOrm()
var user User
user.Name = "slene"
user.IsActive = true

id, err := o.Insert(&user)
if err == nil {
	fmt.Println(id)
}
```

After creation, it will assign values for auto fields.

​	创建后，它将为自动字段分配值。

## InsertMulti

Insert multiple objects in one api.

​	在一次 API 中插入多个对象。

Like sql statement:

​	类似于 sql 语句：

```
insert into table (name, age) values("slene", 28),("astaxie", 30),("unknown", 20)
```

The 1st param is the number of records to insert in one bulk statement. The 2nd param is models slice.

​	第一个参数是在一个批量语句中要插入的记录数。第二个参数是模型切片。

The return value is the number of successfully inserted rows.

​	返回值是成功插入的行数。

```go
users := []User{
	{Name: "slene"},
	{Name: "astaxie"},
	{Name: "unknown"},
	...
}
successNums, err := o.InsertMulti(100, users)
```

When bulk is equal to 1, then models will be inserted one by one.

​	当批量等于 1 时，模型将被逐个插入。

## Update

The first return value is the number of affected rows.

​	第一个返回值是受影响的行数。

```go
o := orm.NewOrm()
user := User{Id: 1}
if o.Read(&user) == nil {
	user.Name = "MyName"
	if num, err := o.Update(&user); err == nil {
		fmt.Println(num)
	}
}
```

Update updates all fields by default. You can update specified fields:

​	默认情况下，Update 会更新所有字段。您可以更新指定字段：

```go
// Only update Name
o.Update(&user, "Name")
// Update multiple fields
// o.Update(&user, "Field1", "Field2", ...)
...
```

For detailed object update, see [One](https://beego.wiki/docs/mvc/model/object/query.md#one)

​	有关详细的对象更新，请参阅 One

## Delete

The first return value is the number of affected rows.

​	第一个返回值是受影响的行数。

```go
o := orm.NewOrm()
if num, err := o.Delete(&User{Id: 1}); err == nil {
	fmt.Println(num)
}
```

Delete will also manipulate reverse relationships. E.g.: `Post` has a foreign key to `User`. If on_delete is set to `cascade`, `Post` will be deleted while delete `User`.

​	Delete 还会处理反向关系。例如： `Post` 具有指向 `User` 的外键。如果 on_delete 设置为 `cascade` ，则在删除 `User` 时将删除 `Post` 。

After deleting, it will clean up values for auto fields.

​	删除后，它将清理自动字段的值。

**Changed in 1.0.3** After deleting, it will **not** clean up values for auto fields.

​	在 1.0.3 中更改删除后，它不会清理自动字段的值。
