+++
title = "composition"
date = 2024-12-15T11:18:27+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Composition](https://github.com/smartystreets/goconvey/wiki/Composition)
>
> 收录该文档时间： `2024-12-15T11:18:27+08:00`

# Composition - 组合测试



Warren Turkal edited this page on Sep 15, 2015 · [7 revisions](https://github.com/smartystreets/goconvey/wiki/Composition/_history)

​	Warren Turkal 于 2015 年 9 月 15 日编辑了此页面 · [7 次修订](https://github.com/smartystreets/goconvey/wiki/Composition/_history)

Writing self-documenting tests is remarkably easy with GoConvey.

​	使用 GoConvey 编写自解释的测试非常简单。

### Examples



First, take a look through the [examples folder](https://github.com/smartystreets/goconvey/tree/master/examples) to get the basic idea. We'd recommend reviewing [isolated_execution_test.go](https://github.com/smartystreets/goconvey/blob/master/convey/isolated_execution_test.go) for a more thorough understanding of how you can compose test cases.

​	首先，查看 [examples 文件夹](https://github.com/smartystreets/goconvey/tree/master/examples) 来获取基本概念。我们推荐阅读 [isolated_execution_test.go](https://github.com/smartystreets/goconvey/blob/master/convey/isolated_execution_test.go)，以更深入地了解如何编写测试用例。

### Functions



See [GoDoc](http://godoc.org/github.com/smartystreets/goconvey) for exported functions and assertions. You'd be most interested in the [convey](http://godoc.org/github.com/smartystreets/goconvey/convey) package.

​	请查看 [GoDoc](http://godoc.org/github.com/smartystreets/goconvey) 获取导出的函数和断言。您可能最感兴趣的是 [convey](http://godoc.org/github.com/smartystreets/goconvey/convey) 包。

### Quick tutorial



In your test file, import needed packages:

​	在您的测试文件中导入所需的包：

```go
import(
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)
```



(Notice the dot-notation for the `convey` package, for convenience.)

​	（注意 `convey` 包的点表示法，这是为了方便。）

Since GoConvey uses `go test`, set up a Test function:

​	由于 GoConvey 使用的是 `go test`，请设置一个 Test 函数：

```go
func TestSomething(t *testing.T) {
	
}
```



To set up test cases, we use `Convey()` to define scope/context/behavior/ideas, and `So()` to make assertions. For example:

​	要设置测试用例，我们使用 `Convey()` 定义范围/上下文/行为/概念，使用 `So()` 进行断言。例如：

```go
Convey("1 should equal 1", t, func() {
	So(1, ShouldEqual, 1)
})
```



There's a working GoConvey test. Notice that we pass in the `*testing.T` object. Only the top-level calls to `Convey()` require that. For nested calls, you must omit it. For instance:	

​	这是一个有效的 GoConvey 测试。注意我们传入了 `*testing.T` 对象。仅顶级的 `Convey()` 调用需要传入它。对于嵌套的调用，必须省略它。例如：

```go
Convey("Comparing two variables", t, func() {
	myVar := "Hello, world!"

	Convey(`"Asdf" should NOT equal "qwerty"`, func() {
		So("Asdf", ShouldNotEqual, "qwerty")
	})

	Convey("myVar should not be nil", func() {
		So(myVar, ShouldNotBeNil)
	})
})
```



If you haven't yet implemented a test or scope, just set its function to `nil` to [skip](https://github.com/smartystreets/goconvey/wiki/Skip) it:

​	如果您尚未实现某个测试或范围，只需将其函数设置为 `nil` 以 [跳过](https://github.com/smartystreets/goconvey/wiki/Skip) 它：

```go
Convey("This isn't yet implemented", nil)
```



### [Next](https://github.com/smartystreets/goconvey/wiki/Assertions)



Next, you should learn about the [standard assertions](https://github.com/smartystreets/goconvey/wiki/Assertions). You may also skip ahead to [executing tests](https://github.com/smartystreets/goconvey/wiki/Execution) or to [Skip](https://github.com/smartystreets/goconvey/wiki/Skip) to make testing more convenient.

​	接下来，您应该了解 [标准断言](https://github.com/smartystreets/goconvey/wiki/Assertions)。您也可以直接跳转到 [执行测试](https://github.com/smartystreets/goconvey/wiki/Execution) 或 [跳过测试](https://github.com/smartystreets/goconvey/wiki/Skip) 以使测试更加方便。
