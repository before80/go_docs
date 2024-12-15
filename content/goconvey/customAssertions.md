+++
title = "Custom Assertions - 自定义断言"
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

# Custom Assertions - 自定义断言

aQua edited this page on Dec 6, 2018 · [5 revisions](https://github.com/smartystreets/goconvey/wiki/Custom-Assertions/_history)

​	aQua 于 2018 年 12 月 6 日编辑了此页面 · [5 次修订](https://github.com/smartystreets/goconvey/wiki/Custom-Assertions/_history)

Sometimes a test suite might need an assertion that is too specific to be included in the general repository. Not to worry, simply implement a function with the following signature (replace the bracketed parts and string values):

​	有时，测试套件可能需要一个过于特定的断言，无法包含在通用的库中。别担心，只需实现一个具有以下签名的函数（替换括号内的部分和字符串值）：

```go
func should<do-something>(actual interface{}, expected ...interface{}) string {
    if <some-important-condition-is-met(actual, expected)> {
        return ""   // empty string means the assertion passed 空字符串表示断言通过
    }
    return "<some descriptive message detailing why the assertion failed...>"
    //  return "<一些描述性消息，详细说明为什么断言失败...>"
}
```



Suppose I implemented the following assertion:

​	假设我实现了以下断言：

```go
func shouldScareGophersMoreThan(actual interface{}, expected ...interface{}) string {
    if actual == "BOO!" && expected[0] == "boo" {
        return ""
    }
    return "Ha! You'll have to get a lot friendlier with the capslock if you want to scare a gopher!"
    // return "哈！如果你想吓唬地鼠，得多用些大写字母！"
}
```



I can then make use of the assertion function when calling the `So()` method in the tests:

​	然后我可以在测试中调用 `So()` 方法时使用这个断言函数：

```go
// Convey("全大写总能让文字更有意义", func() {
Convey("All caps always makes text more meaningful", func() {
    So("BOO!", shouldScareGophersMoreThan, "boo")
})
```



### [Next](https://github.com/smartystreets/goconvey/wiki/Execution)



If you haven't figured out how already, it's time to [learn how to run your tests](https://github.com/smartystreets/goconvey/wiki/Execution).

​	如果你还没有学会，现在是时候[学习如何运行你的测试](https://github.com/smartystreets/goconvey/wiki/Execution)了。
