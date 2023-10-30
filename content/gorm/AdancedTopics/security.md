+++
title = "安全"
date = 2023-10-28T14:36:30+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/security.html](https://gorm.io/docs/security.html)

GORM uses the `database/sql`‘s argument placeholders to construct the SQL statement, which will automatically escape arguments to avoid SQL injection

​	GORM 使用 `database/sql` 的参数占位符来构建 SQL 语句，这将自动转义参数以避免 SQL 注入

> **NOTE** The SQL from Logger is not fully escaped like the one executed, be careful when copying and executing it in SQL console
>
> **注意** Logger 中的 SQL 与执行的 SQL 不完全转义，在使用 SQL 控制台复制和执行时要小心

## 查询条件 Query Condition

User’s input should be only used as an argument, for example:

​	用户输入应该仅用作参数，例如：

``` go
userInput := "jinzhu;drop table users;"

// 安全，将被转义 safe, will be escaped
db.Where("name = ?", userInput).First(&user)

// SQL 注入 SQL injection
db.Where(fmt.Sprintf("name = %v", userInput)).First(&user)
```

## 内联条件 Inline Condition

``` go
// 将被转义 will be escaped
db.First(&user, "name = ?", userInput)

// SQL 注入 SQL injection
db.First(&user, fmt.Sprintf("name = %v", userInput))
```

When retrieving objects with number primary key by user’s input, you should check the type of variable.

​	当根据用户输入检索具有数字主键的对象时，您应检查变量的类型。

``` go
userInputID := "1=1;drop table users;"
// 安全，返回错误 safe, return error
id,err := strconv.Atoi(userInputID)
if err != nil {
    return error
}
db.First(&user, id)

// SQL 注入 SQL injection
db.First(&user, userInputID)
// SELECT * FROM users WHERE 1=1;drop table users;
```

## SQL 注入方法 SQL injection Methods

To support some features, some inputs are not escaped, be careful when using user’s input with those methods

​	为了支持某些功能，一些输入没有被转义，在使用这些方法时要小心

``` go
db.Select("name; drop table users;").First(&user)
db.Distinct("name; drop table users;").First(&user)

db.Model(&user).Pluck("name; drop table users;", &names)

db.Group("name; drop table users;").First(&user)

db.Group("name").Having("1 = 1;drop table users;").First(&user)

db.Raw("select name from users; drop table users;").First(&user)

db.Exec("select name from users; drop table users;")

db.Order("name; drop table users;").First(&user)
```

The general rule to avoid SQL injection is don’t trust user-submitted data, you can perform whitelist validation to test user input against an existing set of known, approved, and defined input, and when using user’s input, only use them as an argument.

​	避免 SQL 注入的一般规则是不要信任用户提交的数据，您可以对用户输入进行白名单验证，以测试用户输入是否与现有已知、批准和定义的输入集相匹配，并在使用用户输入时仅将它们用作参数。