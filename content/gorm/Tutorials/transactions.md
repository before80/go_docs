+++
title = "事务"
date = 2023-10-28T14:30:45+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/transactions.html](https://gorm.io/docs/transactions.html)

## 禁用默认事务 Disable Default Transaction

GORM perform write (create/update/delete) operations run inside a transaction to ensure data consistency, you can disable it during initialization if it is not required, you will gain about 30%+ performance improvement after that

​	GORM在执行写（创建/更新/删除）操作时，会运行在事务中以确保数据一致性。如果在初始化过程中不需要，可以在此处禁用它，之后将获得大约30%+的性能提升。

``` go
// 全局禁用 Globally disable
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
})

// 连续会话模式 Continuous session mode
tx := db.Session(&Session{SkipDefaultTransaction: true})
tx.First(&user, 1)
tx.Find(&users)
tx.Model(&user).Update("Age", 18)
```

## 事务 Transaction

To perform a set of operations within a transaction, the general flow is as below.

​	要执行一组操作的事务，一般流程如下。

``` go
db.Transaction(func(tx *gorm.DB) error {
  // 在此点之后的数据库操作中使用'tx'（而不是'db'）进行一些数据库操作 do some database operations in the transaction (use 'tx' from this point, not 'db')
  if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
    // 返回任何错误都将回滚 return any error will rollback
    return err
  }

  if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
    return err
  }

  // 返回nil将提交整个事务 return nil will commit the whole transaction
  return nil
})
```

### 嵌套事务 Nested Transactions

GORM supports nested transactions, you can rollback a subset of operations performed within the scope of a larger transaction, for example:

​	GORM支持嵌套事务，您可以回滚在更大事务范围内执行的操作的一部分，例如：

``` go
db.Transaction(func(tx *gorm.DB) error {
  tx.Create(&user1)

  tx.Transaction(func(tx2 *gorm.DB) error {
    tx2.Create(&user2)
    return errors.New("rollback user2") // Rollback user2
  })

  tx.Transaction(func(tx2 *gorm.DB) error {
    tx2.Create(&user3)
    return nil
  })

  return nil
})

// 提交 user1, user3 Commit user1, user3
```

## 手动控制事务 Control the transaction manually

Gorm supports calling transaction control functions (commit / rollback) directly, for example:

​	Gorm支持直接调用事务控制函数（提交 / 回滚），例如：

``` go
// 开始一个事务 begin a transaction
tx := db.Begin()

// 在此点之后的数据库操作中使用'tx'（而不是'db'）进行一些数据库操作 do some database operations in the transaction (use 'tx' from this point, not 'db')
tx.Create(...)

// ...

// 发生错误时的回滚事务 rollback the transaction in case of error
tx.Rollback()

// 或者提交事务 Or commit the transaction
tx.Commit()
```

### 一个具体的例子  A Specific Example

``` go
func CreateAnimals(db *gorm.DB) error {
  // 注意一旦你进入事务，就使用tx作为数据库句柄 Note the use of tx as the database handle once you are within a transaction
  tx := db.Begin()
  defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()

  if err := tx.Error; err != nil {
    return err
  }

  if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
     tx.Rollback()
     return err
  }

  if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
     tx.Rollback()
     return err
  }

  return tx.Commit().Error
}
```

## SavePoint, RollbackTo

GORM provides `SavePoint`, `RollbackTo` to save points and roll back to a savepoint, for example:

​	GORM提供`SavePoint`和`RollbackTo`来保存点和回滚到保存点，例如：

``` go
tx := db.Begin()
tx.Create(&user1)

tx.SavePoint("sp1")
tx.Create(&user2)
tx.RollbackTo("sp1") // Rollback user2

tx.Commit() // Commit user1
```