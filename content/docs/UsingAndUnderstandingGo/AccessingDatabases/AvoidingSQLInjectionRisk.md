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

​	你可以通过将SQL参数值作为`sql`包函数的参数来避免SQL注入风险。`sql`包中的许多函数为SQL语句和该语句参数中使用的值提供参数（其他函数为预处理语句和一个参数提供参数）。

​	以下示例代码使用`?`符号作为`id`参数的占位符，并将其作为函数实参提供：

```go 
// 执行带参数的SQL语句的正确格式。
rows, err := db.Query("SELECT * FROM user WHERE id = ?", id)
```

​	执行数据库操作的`sql`包函数将从您提供的参数中创建预处理语句。运行时，`sql`包将SQL语句转换为预处理语句并将其与单独的参数一起发送。

> 注意：参数占位符因使用的`DBMS`和驱动程序而异。例如，Postgres的[pq driver](https://pkg.go.dev/github.com/lib/pq)接受像`$1`这样的占位符形式，而不是`?`

​	您可能想要使用`fmt`包中的一个函数来组装包含参数的SQL语句字符串，如下所示：

```go 
// SECURITY RISK! => 安全风险!
rows, err := db.Query(fmt.Sprintf("SELECT * FROM user WHERE id = %s", id))
```

​	**这是不安全的！**当你这样做时，Go会先组装整个SQL语句，用参数值替换`%s`格式动词，然后将完整的语句发送给DBMS。这会导致[SQL注入](https://en.wikipedia.org/wiki/SQL_injection)风险，因为调用者的代码可以发送一个意外的SQL片段作为`id`参数。该片段可能以不可预测的方式完成SQL语句，对你的应用程序构成危险。

​	例如，通过传递某个`%s`值，您可能会得到以下内容，它可能返回您的数据库中的所有用户记录：

```mysql
SELECT * FROM user WHERE id = 1 OR 1=1;
```