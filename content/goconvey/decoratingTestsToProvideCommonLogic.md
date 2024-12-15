+++
title = "为测试装饰以提供通用逻辑"
date = 2024-12-15T11:19:50+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Decorating-tests-to-provide-common-logic](https://github.com/smartystreets/goconvey/wiki/Decorating-tests-to-provide-common-logic)
>
> 收录该文档时间： `2024-12-15T11:19:50+08:00`

# Decorating tests to provide common logic - 为测试装饰以提供通用逻辑



Michael Tiller edited this page on May 1, 2015 · [2 revisions](https://github.com/smartystreets/goconvey/wiki/Decorating-tests-to-provide-common-logic/_history)

​	Michael Tiller 于 2015 年 5 月 1 日编辑了此页面 · [2 次修订](https://github.com/smartystreets/goconvey/wiki/Decorating-tests-to-provide-common-logic/_history)

1 1go testsh

​	有时，测试套件需要为每个测试设置和清理某些特定资源或状态。通过将测试包装在辅助函数中，可以简化这一过程。

This is useful for otherwise shared common state, like database transactions, where you set up the transaction, let the test operate on it, check that it is still alive and finally perform a rollback before the next test.

​	这在共享通用状态（如数据库事务）的情况下特别有用。你可以设置事务，让测试操作它，检查事务是否仍然有效，最后在下一个测试之前执行回滚。

## 基本原则 Basic principle



The basic idea is this:

​	基本思路如下：

- Provide a function which performs the following: 提供一个函数，该函数完成以下操作：

  1. Accepts a closure `f` as a parameter which will be invoked with the initialized resource or state 接收一个闭包 `f` 作为参数，该闭包将在初始化的资源或状态上调用。

  2. Creates and returns a `func()` which performs the following: 创建并返回一个`func()`它执行以下操作：

     1. Initializes the resource 初始化资源
     2. Provides a `Reset` for the resource, if needed 为资源提供一个 `Reset`（如果需要）
     3. Executes the closure `f `执行闭包 `f`
  
- Then to use this function in tests: 然后在测试中使用这个函数：

  1. Create a closure which takes the initialized resource or state as a parameter 创建一个闭包，该闭包接收初始化的资源或状态作为参数。
  2. Pass this closure to the created helper function 将此闭包传递给创建的辅助函数。
  3. Pass the result from the helper function as the block to a `Convey` invocation 将辅助函数的结果作为 `Convey` 调用的块传递。

**Note**: When reporting failures, the failure will be reported by traversing the stack trace looking for the first location in a file that ends with `_test.go`. For this reason, it is important that calls to `So` appear in such a file. Otherwise, the error will be reported at a different (probably higher) level in the stack and it will be unclear exactly what failure has occurred.

​	**注意**：在报告失败时，错误会通过遍历堆栈跟踪查找文件名以 `_test.go` 结尾的第一个位置来报告。因此，`So` 的调用必须出现在这样的文件中。否则，错误会在堆栈的更高层级报告，导致难以明确失败的具体原因。

## Example



```go
package main

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
)

func WithTransaction(db *sql.DB, f func(tx *sql.Tx)) func() {
	return func() {
		tx, err := db.Begin()
		So(err, ShouldBeNil)

		Reset(func() {
			/* Verify that the transaction is alive by executing a command */
			/* 验证事务是否存活，执行一个命令 */
			_, err := tx.Exec("SELECT 1")
			So(err, ShouldBeNil)

			tx.Rollback()
		})

		/* Here we invoke the actual test-closure and provide the transaction */
		/* 调用实际的测试闭包，并提供事务 */
		f(tx)
	}
}

func TestUsers(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://localhost?sslmode=disable")
	if err != nil {
		panic(err)
	}

   // Convey("给定数据库中的一个用户", t, WithTransaction(db, func(tx *sql.Tx) {
	Convey("Given a user in the database", t, WithTransaction(db, func(tx *sql.Tx) {
		_, err := tx.Exec(`INSERT INTO "Users" ("id", "name") VALUES (1, 'Test User')`)
		So(err, ShouldBeNil)
		
		// Convey("尝试检索该用户应该返回该用户", func() {
		Convey("Attempting to retrieve the user should return the user", func() {
			 var name string

			 data := tx.QueryRow(`SELECT "name" FROM "Users" WHERE "id" = 1`)
			 err = data.Scan(&name)

			 So(err, ShouldBeNil)
			 So(name, ShouldEqual, "Test User")
		})
	}))
}

/* Required table to run the test: 运行测试所需的表：
CREATE TABLE "public"."Users" ( 
	"id" INTEGER NOT NULL UNIQUE, 
	"name" CHARACTER VARYING( 2044 ) NOT NULL
);
*/
```



Output:

```
=== RUN TestUsers

  Given a user in the database ✔✔
    Attempting to retrieve the user should return the user ✔✔✔

--- PASS: TestUsers (0.01 seconds)
PASS
ok      test    0.022s
```
