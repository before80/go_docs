+++
title = "避免SQL注入风险"
date = 2023-05-17T15:03:14+08:00
weight = 9
description = ""
isCJKLanguage = true
draft = false
+++
# Avoiding SQL injection risk - 避免SQL注入风险

> 原文：[https://go.dev/doc/database/sql-injection](https://go.dev/doc/database/sql-injection)

​	你可以通过提供SQL参数值作为`sql`包的函数实参来避免SQL注入风险。`sql`包中的许多函数为SQL语句和用于该语句的参数值提供了参数（其他函数为预处理语句和参数提供一个参数）。

​	下面的例子中的代码使用`?` 符号作为`id`参数的占位符，该参数是作为函数实参提供的：

```go linenums="1"
// Correct format for executing an SQL statement with parameters. => 执行带参数的SQL语句的正确格式。
rows, err := db.Query("SELECT * FROM user WHERE id = ?", id)
```

​	执行数据库操作的`sql`包函数从你提供的实参中创建预处理语句。在运行时，`sql`包将SQL语句变成一个预处理语句，并将其与独立的参数一起发送。

注意：参数占位符因你所使用的`DBMS`和驱动而不同。例如，`Postgres`的[pq driver](https://pkg.go.dev/github.com/lib/pq)接受`$1`这样的占位符形式，而不是 `?`。

​	你可能会想使用`fmt`包中的一个函数来把SQL语句组合成一个包含参数的字符串——比如这样：

```go linenums="1"
// SECURITY RISK! => 安全风险!
rows, err := db.Query(fmt.Sprintf("SELECT * FROM user WHERE id = %s", id))
```

​	**这是不安全的!** 当你这样做时，Go会组装整个SQL语句，用参数值替换`%s`格式的动词，然后再将完整的语句发送给DBMS。**这会带来[SQL注入](https://en.wikipedia.org/wiki/SQL_injection)的风险**，因为代码的调用者可能会发送一个意外的SQL代码片段作为`id`参数。该代码片段可能以不可预测的方式完成SQL语句，对你的应用程序造成危险。

​	例如，通过传递某个`%s`值，你可能会得到如下语句，这可能会返回你数据库中的所有用户记录：

```mysql
SELECT * FROM user WHERE id = 1 OR 1=1;
```