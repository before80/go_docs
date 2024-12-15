+++
title = "Custom Assertions"
date = 2024-12-15T11:19:01+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Custom-Assertions](https://github.com/smartystreets/goconvey/wiki/Custom-Assertions)
>
> 收录该文档时间： `2024-12-15T11:19:01+08:00`

# Custom Assertions

aQua edited this page on Dec 6, 2018 · [5 revisions](https://github.com/smartystreets/goconvey/wiki/Custom-Assertions/_history)

Sometimes a test suite might need an assertion that is too specific to be included in the general repository. Not to worry, simply implement a function with the following signature (replace the bracketed parts and string values):

```
func should<do-something>(actual interface{}, expected ...interface{}) string {
    if <some-important-condition-is-met(actual, expected)> {
        return ""   // empty string means the assertion passed
    }
    return "<some descriptive message detailing why the assertion failed...>"
}
```



Suppose I implemented the following assertion:

```
func shouldScareGophersMoreThan(actual interface{}, expected ...interface{}) string {
    if actual == "BOO!" && expected[0] == "boo" {
        return ""
    }
    return "Ha! You'll have to get a lot friendlier with the capslock if you want to scare a gopher!"
}
```



I can then make use of the assertion function when calling the `So()` method in the tests:

```
Convey("All caps always makes text more meaningful", func() {
    So("BOO!", shouldScareGophersMoreThan, "boo")
})
```



### [Next](https://github.com/smartystreets/goconvey/wiki/Execution)



If you haven't figured out how already, it's time to [learn how to run your tests](https://github.com/smartystreets/goconvey/wiki/Execution).
