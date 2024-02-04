+++
title = "事务"
date = 2024-02-04T10:01:37+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/model/transaction/](https://beego.wiki/docs/mvc/model/transaction/)

# Transaction 事务



## Transaction 事务

There are two ways to handle transaction in Beego. One is closure:

&zeroWidthSpace;在 Beego 中有两种处理事务的方法。一种是闭包：

```go
// Beego will manage the transaction's lifecycle
// if the @param task return error, the transaction will be rollback
// or the transaction will be committed
err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
	// data
	user := new(User)
	user.Name = "test_transaction"

	// insert data
	// Using txOrm to execute SQL
	_, e := txOrm.Insert(user)
	// if e != nil the transaction will be rollback
	// or it will be committed
	return e
})
```

In this way, the first parameter is `task`, all DB operation should be inside the task.

&zeroWidthSpace;这种方式中，第一个参数是 `task` ，所有数据库操作都应该在任务内部。

If the task return error, Beego rollback the transaction.

&zeroWidthSpace;如果任务返回错误，Beego 会回滚事务。

We recommend you to use this way.

&zeroWidthSpace;我们建议您使用这种方式。

Another way is that users handle transaction manually:

&zeroWidthSpace;另一种方式是用户手动处理事务：

```go
	o := orm.NewOrm()
	to, err := o.Begin()
	if err != nil {
		logs.Error("start the transaction failed")
		return
	}

	user := new(User)
	user.Name = "test_transaction"

	// do something with to. to is an instance of TxOrm

	// insert data
	// Using txOrm to execute SQL
	_, err = to.Insert(user)

	if err != nil {
		logs.Error("execute transaction's sql fail, rollback.", err)
		err = to.Rollback()
		if err != nil {
			logs.Error("roll back transaction failed", err)
		}
	} else {
		err = to.Commit()
		if err != nil {
			logs.Error("commit transaction failed.", err)
		}
	}
o := orm.NewOrm()
to, err := o.Begin()

// outside the txn
o.Insert(xxx)

// inside the txn
to.Insert(xxx)
```