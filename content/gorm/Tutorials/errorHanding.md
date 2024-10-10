+++
title = "错误处理"
date = 2023-10-28T14:29:47+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/error_handling.html](https://gorm.io/docs/error_handling.html)

In Go, error handling is important.

​	在Go语言中，错误处理非常重要。

You are encouraged to do error check after any [Finisher Methods](https://gorm.io/docs/method_chaining.html#finisher_method)

​	在使用任何[Finisher Methods]({{< ref "/gorm/Tutorials/methodChaining" >}})之后，建议进行错误检查。

## 错误处理 Error Handling

Error handling in GORM is different than idiomatic Go code because of its chainable API.

​	由于GORM具有链式API，因此其错误处理与典型的Go代码不同。

If any error occurs, GORM will set `*gorm.DB`‘s `Error` field, you need to check it like this:

​	如果发生任何错误，GORM将设置`*gorm.DB`的`Error`字段，您需要像这样检查它：

``` go
if err := db.Where("name = ?", "jinzhu").First(&user).Error; err != nil {
  // error handling...
}
```

Or

或者

``` go
if result := db.Where("name = ?", "jinzhu").First(&user); result.Error != nil {
  // error handling...
}
```

## ErrRecordNotFound

GORM returns `ErrRecordNotFound` when failed to find data with `First`, `Last`, `Take`, if there are several errors happened, you can check the `ErrRecordNotFound` error with `errors.Is`, for example:

​	当使用`First`、`Last`、`Take`查找数据失败时，GORM会返回`ErrRecordNotFound`。如果有多个错误发生，您可以使用`errors.Is`检查`ErrRecordNotFound`错误，例如：

``` go
// 检查是否返回RecordNotFound错误 Check if returns RecordNotFound error
err := db.First(&user, 100).Error
errors.Is(err, gorm.ErrRecordNotFound)
```

## Dialect Translated Errors

If you would like to be able to use the dialect translated errors(like ErrDuplicatedKey), then enable the TranslateError flag when opening a db connection.

​	如果您希望能够使用方言翻译的错误（如ErrDuplicatedKey），则在打开数据库连接时启用TranslateError标志。

``` go
db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{TranslateError: true})
```