+++
title = "回调"
date = 2024-02-04T21:14:52+08:00
weight = 12
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/database/callbacks/](https://gobuffalo.io/documentation/database/callbacks/)

# Callbacks 回调 

Pop provides a means to execute code before and after database operations. This is done by defining specific methods for your models.

​	Pop 提供了一种在数据库操作之前和之后执行代码的方法。这是通过为您的模型定义特定方法来完成的。

For example, to hash a user password you may want to define the following method:

​	例如，要对用户密码进行哈希处理，您可能需要定义以下方法：

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password string
}

func (u *User) BeforeCreate(tx *pop.Connection) error {
  hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  if err != nil {
    return errors.WithStack(err)
  }

  u.Password = string(hash)

  return nil
}
```

In the above example, when the connection’s `Save` method is called with a `User`, the `BeforeCreate` method will be called before writing to the database.

​	在上面的示例中，当连接的 `Save` 方法使用 `User` 调用时， `BeforeCreate` 方法将在写入数据库之前调用。

The available callbacks include:

​	可用的回调包括：

- [BeforeSave](https://godoc.org/github.com/gobuffalo/pop#BeforeSaveable)
- [BeforeCreate](https://godoc.org/github.com/gobuffalo/pop#BeforeCreateable)
- [BeforeUpdate](https://godoc.org/github.com/gobuffalo/pop#BeforeUpdateable)
- [BeforeDestroy](https://godoc.org/github.com/gobuffalo/pop#BeforeDestroyable)
- [BeforeValidate](https://godoc.org/github.com/gobuffalo/pop#BeforeValidateable)
- [AfterSave](https://godoc.org/github.com/gobuffalo/pop#AfterSaveable)
- [AfterCreate](https://godoc.org/github.com/gobuffalo/pop#AfterCreateable)
- [AfterUpdate](https://godoc.org/github.com/gobuffalo/pop#AfterUpdateable)
- [AfterDestroy](https://godoc.org/github.com/gobuffalo/pop#AfterDestroyable)
- [AfterFind](https://godoc.org/github.com/gobuffalo/pop#AfterFindable)

## Related Content 相关内容 

- [Models](https://gobuffalo.io/documentation/database/models) - Define a database model.
  模型 - 定义数据库模型。