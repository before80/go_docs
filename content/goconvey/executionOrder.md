+++
title = "Execution order - 执行顺序"
date = 2024-12-15T11:20:56+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Execution-order](https://github.com/smartystreets/goconvey/wiki/Execution-order)
>
> 收录该文档时间： `2024-12-15T11:20:56+08:00`

# Execution order - 执行顺序



Jones edited this page on Aug 12, 2021 · [4 revisions](https://github.com/smartystreets/goconvey/wiki/Execution-order/_history)

​	Jones 于 2021 年 8 月 12 日编辑了此页面 · [4 次修订](https://github.com/smartystreets/goconvey/wiki/Execution-order/_history)

As an extension of the FAQ, this document answers some important questions about the execution model of GoConvey, like:

​	作为常见问题 (FAQ) 的扩展，本文档回答了一些关于 GoConvey 执行模型的重要问题，例如：

1. How do I define a "Setup" method to be run before each test? 如何定义一个“Setup”方法在每个测试前运行？
2. Why don't my nested tests run in sequential order? 为什么我的嵌套测试没有按顺序执行？

These questions, surprisingly, are related. Here's an eloquent explanation from one of [GoConvey's users](https://github.com/smartystreets/goconvey/issues/111) that sheds light on these questions:

​	这些问题出乎意料地是相关的。以下是 [GoConvey 用户](https://github.com/smartystreets/goconvey/issues/111) 的一段精彩解释，阐明了这些问题：

> Consider, for example, the pseudocode:
>
> ​	考虑以下伪代码：

```
Convey A
    So 1
    Convey B
        So 2
    Convey C
        So 3
```



> I Initially thought would execute as `A1B2C3`, in other words, sequentially. As you all know, this is actually executed first as `A1B2` and then `A1C3`. Once I realized this, I actually realized the power of goconvey, because it allows you two write n tests with log(n) statements. This "tree-based" behavioral testing eliminates so much duplicated setup code and is so much easier to read for completeness (versus pages of unit tests) while still allowing for very well isolated tests (for each branch in the tree).
>
> ​	起初我认为它会按顺序执行，即 `A1B2C3`。然而，实际的执行顺序是 `A1B2` 和 `A1C3`。一旦我意识到这一点，就明白了 GoConvey 的强大之处，因为它允许你用 log(n) 的代码写出 n 个测试。这种“基于树”的行为测试消除了大量重复的设置代码，同时使测试的完整性更容易理解（与一大堆单元测试相比），并且仍然可以很好地隔离每个分支的测试。

In the pseudocode above, `Convey A` serves as a "Setup" method for `Convey B` and `Convey C` and is run separately for each. Here's a more complex example:

​	在上述伪代码中，`Convey A` 作为 `Convey B` 和 `Convey C` 的“Setup”方法，分别单独运行。以下是一个更复杂的示例：

```
Convey A
    So 1
    Convey B
        So 2
        Convey Q
        	So 9
    Convey C
        So 3
```



Can you guess what the output would be?

​	你能猜出输出是什么吗？

`A1B2Q9A1C3` is the correct answer.

​	正确答案是：`A1B2Q9A1C3`。

You're welcome to peruse the [tests in the GoConvey project itself](https://github.com/smartystreets/goconvey/blob/master/convey/isolated_execution_test.go) that document this behavior.

​	你可以查看 [GoConvey 项目本身的测试](https://github.com/smartystreets/goconvey/blob/master/convey/isolated_execution_test.go)，这些测试记录了这种行为。

**注意事项 Gotchas**

Remember that every Convey() call in Go creates a new scope. You should use Foo **=** &Bar{} in order to assign a new value to a previous declared variable. Using foo **:=** &Bar{} creates a new variable in the current scope. Example:

​	请记住，GoConvey 中的每个 `Convey()` 调用都会创建一个新作用域。要为之前声明的变量赋值，请使用 `Foo = &Bar{}`。使用 `foo := &Bar{}` 会在当前作用域中创建一个新变量。例如：

```go
    Convey("Setup", func() {
        foo := &Bar{}
        Convey("This creates a new variable foo in this scope", func() {
            foo := &Bar{}
        }
        Convey("This assigns a new value to the previous declared foo", func() {
            foo = &Bar{}
        }
    }
```



If you're wondering about how to achieve "tear-down" functionality, see the page on the [Reset](https://github.com/smartystreets/goconvey/wiki/Reset) function.

​	如果您想知道如何实现“拆卸”（tear-down）功能，请参阅关于 [Reset](https://github.com/smartystreets/goconvey/wiki/Reset) 函数的页面。
