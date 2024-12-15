+++
title = "Home - 首页"
date = 2024-12-15T11:17:39+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki](https://github.com/smartystreets/goconvey/wiki)
>
> 收录该文档时间： `2024-12-15T11:17:39+08:00`

# Home - 首页



Elmar Hinz edited this page on Feb 5, 2023 · [16 revisions](https://github.com/smartystreets/goconvey/wiki/Home/_history)

​	Elmar Hinz 于 2023 年 2 月 5 日编辑了此页面 · [16 次修订](https://github.com/smartystreets/goconvey/wiki/Home/_history)

Welcome to GoConvey, a yummy testing tool for gophers.

​	欢迎使用 GoConvey，一款专为 gopher 打造的优秀测试工具。

**[Documentation & tutorial](https://github.com/smartystreets/goconvey/wiki/Documentation)**

### 主要功能 Main Features



- Integrates with `go test`
  - 与 `go test` 无缝集成

- Readable, colorized console output
  - 可读的、彩色的控制台输出

- Fully-automatic web UI
  - 全自动 Web UI

- Huge suite of regression tests
  - 大量回归测试

- Test code generator
  - 测试代码生成器


View a **[comprehensive table of all features](https://github.com/smartystreets/goconvey/wiki/Features-Table)** compared to other Go testing tools.

​	查看与其他 Go 测试工具对比的 **[完整功能表](https://github.com/smartystreets/goconvey/wiki/Features-Table)**。

### 25 秒快速入门 Get going in 25 seconds



1. In your terminal: 在终端中运行以下命令：

```sh
# make sure your GOPATH is set

$ cd <project path>
$ go get github.com/smartystreets/goconvey
$ go install github.com/smartystreets/goconvey
$ $GOPATH/bin/goconvey
```



1. In your browser: 在浏览器中访问：

```
http://localhost:8080
```



If you have existing Go tests, they will run automatically and the results will appear in your browser.

​	如果您已有现成的 Go 测试，它们将自动运行，结果会显示在您的浏览器中。

### 您的第一个 GoConvey 测试 Your first GoConvey test



Open any `_test.go` file and put this in it, customizing your package declaration:

​	打开任意 `_test.go` 文件，按需修改包声明，并添加以下内容：

```
package package_name

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegerStuff(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
```



Save the file, then glance over at your browser window, and you'll see that the new tests have already been run.

​	保存文件后，查看浏览器窗口，您会看到新测试已经运行。

Change the assertion (the line with `So()`) to make the test fail, then see the output change in your browser.

​	修改断言（包含 `So()` 的那一行）使测试失败，然后在浏览器中观察输出的变化。

You can also run tests from the terminal as usual, with `go test`. If you want the tests to run automatically in the terminal, check out [the auto-test script](https://github.com/smartystreets/goconvey/wiki/Auto-test).

​	您也可以像往常一样在终端中运行测试，使用 `go test` 命令即可。如果希望测试在终端中自动运行，可以查看 [自动测试脚本](https://github.com/smartystreets/goconvey/wiki/Auto-test)。

### 必读内容 Required Reading



If I could ensure that every GoConvey user read only one bit of code from this repository it would be the [isolated execution tests](https://github.com/smartystreets/goconvey/blob/master/convey/isolated_execution_test.go). Those tests are the very best documentation for the GoConvey execution model.

​	如果只能确保每位 GoConvey 用户阅读本仓库中的一段代码，那一定是 [隔离执行测试](https://github.com/smartystreets/goconvey/blob/master/convey/isolated_execution_test.go)。这些测试是 GoConvey 执行模型的最佳文档。

### 完整文档 Full Documentation



See the [documentation index](https://github.com/smartystreets/goconvey/wiki/Documentation) for details about assertions, writing tests, execution, etc.

​	查看 [文档索引](https://github.com/smartystreets/goconvey/wiki/Documentation)，了解断言、编写测试、执行等相关内容。
