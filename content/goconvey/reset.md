+++
title = "reset"
date = 2024-12-15T11:22:30+08:00
weight = 14
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Reset](https://github.com/smartystreets/goconvey/wiki/Reset)
>
> 收录该文档时间： `2024-12-15T11:22:30+08:00`

# Reset



Jian Liu edited this page on Dec 18, 2019 · [4 revisions](https://github.com/smartystreets/goconvey/wiki/Reset/_history)

​	Jian Liu 于 2019 年 12 月 18 日编辑了此页面 · [4 次修订](https://github.com/smartystreets/goconvey/wiki/Reset/_history)

When your Conveys have some set-up involved, you may need to tear down after or between tests. Use `Reset()` to clean up in those cases. A Convey's Reset() runs at the end of each `Convey()` within that same scope.

​	当您的 `Convey` 包含一些初始化设置时，可能需要在测试结束后或测试之间进行清理。在这种情况下，可以使用 `Reset()` 方法。`Reset()` 会在同一作用域内的每个 `Convey()` 结束时运行。

For example:

​	示例：

```go
Convey("Top-level", t, func() {

    // setup (run before each `Convey` at this scope):
    // 初始化（在此作用域的每个 `Convey` 之前运行）:
    db.Open()
    db.Initialize()

    Convey("Test a query", t, func() {
        db.Query()
        // TODO: assertions here
        // TODO: 在此添加断言
    })

    Convey("Test inserts", t, func() {
        db.Insert()
        // TODO: assertions here
        // TODO: 在此添加断言
    })

    Reset(func() {
        // This reset is run after each `Convey` at the same scope.
        // 此 Reset 会在同一作用域内每个 `Convey` 之后运行。
        db.Close()
    })

})
```
