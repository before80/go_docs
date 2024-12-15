+++
title = "skip"
date = 2024-12-15T11:36:09+08:00
weight = 15
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Skip](https://github.com/smartystreets/goconvey/wiki/Skip)
>
> 收录该文档时间： `2024-12-15T11:36:09+08:00`

# Skip



Eli Bierman edited this page on Mar 22, 2019 · [2 revisions](https://github.com/smartystreets/goconvey/wiki/Skip/_history)

​	Eli Bierman 于 2019 年 3 月 22 日编辑了此页面 · [2 次修订](https://github.com/smartystreets/goconvey/wiki/Skip/_history)

Sometimes it's nice to ignore or skip an entire scope or some assertions here and there. This is easy with GoConvey.

​	有时，需要忽略或跳过整个作用域或某些断言。使用 GoConvey 可以轻松实现这一点。

### 跳过 `Convey` 注册 Skipping `Convey` registrations



Changing a `Convey()` to `SkipConvey()` prevents the `func()` passed into that call from running. This also has the consequence of preventing any nested `Convey` registrations from running. The reporter will indicate that the registration was skipped.

​	将 `Convey()` 改为 `SkipConvey()` 可防止传入的 `func()` 执行。这也会导致任何嵌套的 `Convey` 注册无法运行，报告中将显示注册被跳过。

```
SkipConvey("Important stuff", func() {			// This func() will not be executed! 此 func() 不会执行！
    Convey("More important stuff", func() {
        So("asdf", ShouldEqual, "asdf")
    })
})
```



Using `SkipConvey()` has nearly the same effect as commenting out the test entirely. However, this is preferred over commenting out tests to avoid the usual "declared/imported but not used" errors. Usage of `SkipConvey()` is intended for temporary code alterations.

​	使用 `SkipConvey()` 的效果几乎等同于完全注释掉测试。但这比注释测试更好，可以避免出现 "声明/导入但未使用" 的错误。`SkipConvey()` 的使用主要用于临时代码修改。

### 未实现的 `Convey` 注册 Unimplemented `Convey` registrations



When composing `Convey` registrations, sometimes it's convenient to use `nil` instead of an actual `func()`. Not only does this skip the scope, but it provides an indication in the report that the registration is not complete, and that it's likely your code is missing some test coverage.

​	在编写 `Convey` 注册时，有时可以方便地使用 `nil` 替代实际的 `func()`。这不仅会跳过该作用域，还会在报告中显示注册未完成，提示代码可能缺少一些测试覆盖率。

```
Convey("Some stuff", func() {

    // This will show up as 'skipped' in the report
    // 报告中将显示为“跳过”
    Convey("Should go boink", nil)

}
```



### 跳过 `So` 断言 Skipping `So` assertions



Similar to `SkipConvey()`, changing a `So()` to `SkipSo()` prevents the execution of that assertion. The report will show that the assertion was skipped.

​	类似于 `SkipConvey()`，将 `So()` 改为 `SkipSo()` 可防止执行该断言。报告中将显示断言被跳过。

```
Convey("1 Should Equal 2", func() {
    
    // This assertion will not be executed and will show up as 'skipped' in the report
    // 此断言不会执行，且报告中将显示为“跳过”
    SkipSo(1, ShouldEqual, 2)

})
```



And like `SkipConvey`, this function is only intended for use during temporary code alterations.

​	和 `SkipConvey` 类似，此功能仅用于临时代码修改。

### 仅运行某些 `Convey` 注册 Running Only Certain `Convey` Registrations



You can use `FocusConvey` to only run certain `Convey` registrations.

​	您可以使用 `FocusConvey` 仅运行某些 `Convey` 注册。

You must mark at least one leaf `Convey` registration (where the actual assertions are) and all of its parent `Convey`s in order for it to work.

​	必须标记至少一个叶子级的 `Convey` 注册（实际包含断言的地方）及其所有父级 `Convey`，才能使其正常工作。

Let's see an example:

​	示例：

```
FocusConvey("A", func() {
    // B will not be run
    // B 不会运行
    Convey("B", nil)
    FocusConvey("C", func() {
        // Only D will be run. 
        // 只有 D 会运行。
        FocusConvey("D", func() {
        })
    })
}
```



You might want to run all subtests of a certain `Convey` registration. In that case every leaf test must be marked with `Convey`, along with all of its parents.

​	如果想运行某个 `Convey` 注册的所有子测试，每个叶子测试及其所有父级都必须标记为 `FocusConvey`。

Here's an example of a common mistake:

​	以下是一个常见错误的示例：

```
Convey("A", func() {
    // test B will still run because test D is not marked with Focus
    // 测试 B 仍然会运行，因为测试 D 未标记为 Focus
    Convey("B", nil)
    FocusConvey("C", func() {
        // Mark test D with Focus to run only test D
        // 将测试 D 标记为 Focus 以仅运行测试 D
        Convey("D", func() {
        })
    })
}
```



Read more in [the documentation on `FocusConvey`.](https://godoc.org/github.com/smartystreets/goconvey/convey#FocusConvey)

​	有关更多信息，请参阅 [`FocusConvey` 的文档](https://godoc.org/github.com/smartystreets/goconvey/convey#FocusConvey)。

### End of tutorial



Congrats, you made it! Now `go test`!

​	恭喜您完成学习！现在可以运行 `go test`！
