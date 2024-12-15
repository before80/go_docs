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

When your Conveys have some set-up involved, you may need to tear down after or between tests. Use `Reset()` to clean up in those cases. A Convey's Reset() runs at the end of each `Convey()` within that same scope.

For example:

```go
Convey("Top-level", t, func() {

    // setup (run before each `Convey` at this scope):
    db.Open()
    db.Initialize()

    Convey("Test a query", t, func() {
        db.Query()
        // TODO: assertions here
    })

    Convey("Test inserts", t, func() {
        db.Insert()
        // TODO: assertions here
    })

    Reset(func() {
        // This reset is run after each `Convey` at the same scope.
        db.Close()
    })

})
```
