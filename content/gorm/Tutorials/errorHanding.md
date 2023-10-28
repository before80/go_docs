+++
title = "Error Handling"
date = 2023-10-28T14:29:47+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/error_handling.html](https://gorm.io/docs/error_handling.html)

In Go, error handling is important.

You are encouraged to do error check after any [Finisher Methods](https://gorm.io/docs/method_chaining.html#finisher_method)

## Error Handling

Error handling in GORM is different than idiomatic Go code because of its chainable API.

If any error occurs, GORM will set `*gorm.DB`‘s `Error` field, you need to check it like this:

```
if err := db.Where("name = ?", "jinzhu").First(&user).Error; err != nil {
  // error handling...
}
```

Or

```
if result := db.Where("name = ?", "jinzhu").First(&user); result.Error != nil {
  // error handling...
}
```

## ErrRecordNotFound

GORM returns `ErrRecordNotFound` when failed to find data with `First`, `Last`, `Take`, if there are several errors happened, you can check the `ErrRecordNotFound` error with `errors.Is`, for example:

```
// Check if returns RecordNotFound error
err := db.First(&user, 100).Error
errors.Is(err, gorm.ErrRecordNotFound)
```

## Dialect Translated Errors

If you would like to be able to use the dialect translated errors(like ErrDuplicatedKey), then enable the TranslateError flag when opening a db connection.

```
db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{TranslateError: true})
```