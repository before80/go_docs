+++
title = "钩子"
date = 2023-10-28T14:30:28+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

## 对象生命周期 Object Life Cycle

Hooks are functions that are called before or after creation/querying/updating/deletion.

​	Hooks是在创建/查询/更新/删除之前或之后调用的函数。

If you have defined specified methods for a model, it will be called automatically when creating, updating, querying, deleting, and if any callback returns an error, GORM will stop future operations and rollback current transaction.

​	如果你为模型定义了指定的方法，它将在创建、更新、查询、删除时自动调用，如果任何回调返回错误，GORM将停止未来的操作并回滚当前事务。

The type of hook methods should be `func(*gorm.DB) error`

​	Hook方法的类型应为`func(*gorm.DB) error`

## Hooks

### 创建对象 Creating an object

可用的创建钩子 Available hooks for creating

``` go
// 开始事务 begin transaction
BeforeSave
BeforeCreate
// 保存关联之前 save before associations
// 插入数据库 insert into database
// 保存关联之后 save after associations
AfterCreate
AfterSave
// 提交或回滚事务 commit or rollback transaction
```

代码示例：Code Example:

``` go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()

  if !u.IsValid() {
    err = errors.New("can't save invalid data")
  }
  return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
  if u.ID == 1 {
    tx.Model(u).Update("role", "admin")
  }
  return
}
```

> **NOTE** Save/Delete operations in GORM are running in transactions by default, so changes made in that transaction are not visible until it is committed, if you return any error in your hooks, the change will be rollbacked
>
> **注意** GORM中的保存/删除操作默认以事务方式运行，因此在该事务中所做的更改在提交之前是不可见的，如果您的钩子中返回任何错误，更改将被回滚

``` go
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
  if !u.IsValid() {
    return errors.New("rollback invalid user")
  }
  return nil
}
```

### 更新对象 Updating an object

Available hooks for updating

​	可用的更新钩子

``` go
// begin transaction
BeforeSave
BeforeUpdate
// save before associations
// update database
// save after associations
AfterUpdate
AfterSave
// commit or rollback transaction
```

Code Example:

​	代码示例：

``` go
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
  if u.readonly() {
    err = errors.New("read only user")
  }
  return
}

// Updating data in same transaction
func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
  if u.Confirmed {
    tx.Model(&Address{}).Where("user_id = ?", u.ID).Update("verfied", true)
  }
  return
}
```

### 删除对象 Deleting an object

Available hooks for deleting

​	可用的删除钩子

``` go
// 开始事务 begin transaction
BeforeDelete
// 从数据库中删除 delete from database
AfterDelete
// 提交或回滚事务 commit or rollback transaction
```

Code Example:

​	代码示例：

``` go
// 在同一个事务中更新数据 Updating data in same transaction
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
  if u.Confirmed {
    tx.Model(&Address{}).Where("user_id = ?", u.ID).Update("invalid", false)
  }
  return
}
```

### 查询对象 Querying an object

Available hooks for querying

​	可用的查询钩子

``` go
// 从数据库加载数据 load data from database
// 预加载（预取） Preloading (eager loading)
AfterFind
```

代码示例： Code Example:

``` go
func (u *User) AfterFind(tx *gorm.DB) (err error) {
  if u.MemberShip == "" {
    u.MemberShip = "user"
  }
  return
}
```

## 修改当前操作 Modify current operation

``` go
func (u *User) BeforeCreate(tx *gorm.DB) error {
  // 通过tx.Statement修改当前操作，例如： Modify current operation through tx.Statement, e.g:
  tx.Statement.Select("Name", "Age")
  tx.Statement.AddClause(clause.OnConflict{DoNothing: true})

  // tx是带有NewDB选项的新会话模式 tx is new session mode with the `NewDB` option
  // 基于它的操作将在相同的事务内运行，但没有任何当前条件 operations based on it will run inside same transaction but without any current conditions
  var role Role
  err := tx.First(&role, "name = ?", user.Role).Error
  // SELECT * FROM roles WHERE name = "admin"
  // ...
  return err
}
```