+++
title = "FAQ"
date = 2024-12-15T11:21:02+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/FAQ](https://github.com/smartystreets/goconvey/wiki/FAQ)
>
> 收录该文档时间： `2024-12-15T11:21:02+08:00`

# FAQ - 常见问题解答



A. Svensson edited this page on Aug 20, 2015 · [15 revisions](https://github.com/smartystreets/goconvey/wiki/FAQ/_history)

​	A. Svensson 于 2015 年 8 月 20 日编辑了此页面 · [15 次修订](https://github.com/smartystreets/goconvey/wiki/FAQ/_history)

#### 什么是 GoConvey？What is GoConvey?



Basically, GoConvey is an extension of the built-in Go test tool. It facilitates [Behavior-driven Development (BDD)](https://en.wikipedia.org/wiki/Behavior-driven_development) in Go, though this is not the only way to use it. Many people continue to write traditional Go tests but prefer [GoConvey's web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) for reporting test results.

​	基本上，GoConvey 是对内置 Go 测试工具的扩展。它支持 [行为驱动开发 (BDD)](https://en.wikipedia.org/wiki/Behavior-driven_development)，但这并不是唯一的用途。许多人仍然编写传统的 Go 测试，但更喜欢使用 [GoConvey 的 Web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) 来报告测试结果。

There are two main parts to GoConvey:

​	GoConvey 主要包括两个部分：

1. A comprehensive BDD framework 一个全面的 BDD 框架
2. A server & web UI 一个服务器和 Web UI

Both parts are optional and can be used independently according to your workflow.

​	这两个部分是可选的，可以根据您的工作流程独立使用。

#### 我可以在 `go test` 中使用 GoConvey 吗？Can I use GoConvey with `go test`?



Yes, that's the point!

​	可以，这正是它的目的！

#### 为什么我的嵌套 `Convey` 块以一种奇怪的顺序执行？Why do my nested `Convey` blocks execute in a strange order?



Please read the [isolated execution tests](https://github.com/smartystreets/goconvey/blob/master/convey/isolated_execution_test.go). They are the best form of documentation for the execution model of GoConvey, which is very powerful but not apparent at first.

​	请阅读 [isolated execution tests](https://github.com/smartystreets/goconvey/blob/master/convey/isolated_execution_test.go)。它们是 GoConvey 执行模型的最佳文档，虽然非常强大，但起初可能不太明显。

#### 什么是 [Web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI)？What is [the web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI)?



It's test results in your browser. It updates automatically as files in the watched directories are saved or changed. See the [web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) wiki page for more information.

​	这是显示测试结果的浏览器界面。它会在监视的目录中的文件保存或更改时自动更新。有关详细信息，请参阅 [Web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) 的维基页面。

#### 如何让测试自动运行？How can I get tests to run automatically?



If you're using the [web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) with `goconvey` (the server) to watch tests in your browser, then tests already run automatically when `.go` files are changed.

​	如果您在浏览器中使用 `goconvey`（服务器）和 [Web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) 来监视测试，则在更改 `.go` 文件时测试会自动运行。

For running tests in your terminal, check out how to use [the auto-run.py auto-test script](https://github.com/smartystreets/goconvey/wiki/Auto-test). (It's really easy!)

​	要在终端中运行测试，请查看如何使用 [auto-run.py 自动测试脚本](https://github.com/smartystreets/goconvey/wiki/Auto-test)。（非常简单！）

#### [Web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) 是否适用于传统的 Go 测试？Does the [web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) work with traditional Go tests?



Yep! If you haven't ported all your tests over to GoConvey, your traditional Go test cases will still be run and the results will be reported in the browser.

​	是的！即使您没有将所有测试移植到 GoConvey，传统的 Go 测试用例仍会运行，结果会在浏览器中报告。

#### 如何让调试输出显示在断言旁边，而不是在函数级别？How can I get debug output to appear next to my assertions, rather than up at the function level?



Use `convey.Print` or `convey.Printf` or `convey.Println` just as you would from the `fmt` package. This will cause the output to show up by your assertions rather than up at a higher level.

​	使用 `convey.Print` 或 `convey.Printf` 或 `convey.Println`，就像使用 `fmt` 包一样。这将使输出显示在断言旁，而不是在更高的级别。

#### 如何强制测试在失败后继续执行？How can I force tests to continue executing even after a failure?



By default, a test failure or panic causes future tests in that scope to halt. To have tests continue running, you can pass in a FailureMode in a `Convey()` call:

​	默认情况下，测试失败或 panic 会导致该范围内的后续测试停止运行。要让测试继续运行，可以在 `Convey()` 调用中传入一个 FailureMode：

```go
Convey("A", t, FailureContinues, func() {
    // ...
})
```



All nested Conveys will inherit that setting. You can also set the default FailureMode in an `init()` function with `SetDefaultFailureMode()`, like so:

​	所有嵌套的 Convey 都将继承此设置。您还可以在 `init()` 函数中使用 `SetDefaultFailureMode()` 设置默认 FailureMode，如下所示：

```go
func init() {
    SetDefaultFailureMode(FailureContinues)
}
```



#### 为什么我无法打开测试文件的覆盖率报告（404 Not Found）？Why can't I open a coverage report (404 Not Found) of a tested file?



You have to make sure that the package you are testing lives inside your $GOPATH.

​	您需要确保正在测试的包位于您的 `$GOPATH` 内。

#### GoConvey 受到支持吗？Is GoConvey supported?



Not in the commercial sense of the word, no. Even though it is "sponsored" by SmartyStreets (in that a couple of devs were given company time to work on it), it comes as-is, so "use-at-your-own-risk" -- though honestly you'll probably quite enjoy it. :)

​	从商业意义上讲，不是。尽管 SmartyStreets 提供了“赞助”（即开发者可以利用公司时间来开发它），但它是“按原样提供”的，因此是“使用自负风险”——不过，老实说，您可能会非常喜欢它。

But yes, in the open-source sense of the word, it is supported, meaning: it's not a defunct project. It's very much alive and well. Feel free to submit a pull request with contributions!

​	但从开源的意义上来说，是的，这仍然是一个活跃的项目。欢迎提交拉取请求贡献代码！

#### GoConvey 是从哪里开始的？Where did GoConvey start?



When SmartyStreets decided to scrap its .NET code (which was pretty much everything) in favor of Go, all the developers started learning and practicing with Go. While Go's built-in `go test` command and standard `testing` package excited us, in practice it left us wanting.

​	当 SmartyStreets 决定放弃其 .NET 代码（几乎所有代码）并转向 Go 时，所有开发者都开始学习并实践 Go。虽然 Go 的内置 `go test` 命令和标准 `testing` 包让我们很兴奋，但在实践中仍有不足。

After our first project re-write in Go, using the standard library tests, we found that our tests didn't clearly document what our code was doing, and how it should behave. Assertions weren't clear because they were backward.

​	在用标准库测试完成第一个 Go 项目重写后，我们发现测试并未清晰地记录代码的行为。断言表达方式不够直观。

On top of that (which is solved by other, more lightweight libraries), we thought a browser tab with test results displayed visually would be really cool and it sounded fun. So originally, GoConvey was written for internal use at SmartyStreets. Then we decided to make it kind of our gift to the Golang community.

​	除此之外（其他一些轻量库也解决了这一问题），我们认为用一个浏览器标签直观地显示测试结果会很酷，听起来也很有趣。所以最初，GoConvey 是为内部使用而写的。后来我们决定将其作为一份礼物献给 Go 社区。

#### 你们会添加某些功能吗？Will you add this or that feature?



Maybe, but GoConvey works well enough for us at SmartyStreets as-is. You're welcome to do it.

​	可能，但 GoConvey 已经满足了 SmartyStreets 的需求。您可以自行添加。

#### 如何贡献代码？How do I contribute?



See our [contributors](https://github.com/smartystreets/goconvey/wiki/For-Contributors) page.

​	请查看我们的 [贡献者页面](https://github.com/smartystreets/goconvey/wiki/For-Contributors)。

#### 如果我是测试新手，你推荐阅读什么？If I'm new to testing, what do you recommend reading?



- [Unit Testing](http://en.wikipedia.org/wiki/Unit_testing)
- [Test-Driven Development (TDD)](http://en.wikipedia.org/wiki/Test-driven_development)
- [Behavior-Driven Development (BDD)](http://en.wikipedia.org/wiki/Behavior-driven_development)
- [BDD vs. TDD](http://stackoverflow.com/questions/2509/what-are-the-primary-differences-between-tdd-and-bdd)
- [Laws of TDD](http://butunclebob.com/ArticleS.UncleBob.TheThreeRulesOfTdd)
- [Integration Testing](http://en.wikipedia.org/wiki/Integration_testing)

#### 为什么创建 GoConvey？（`go test` 还不够好吗？）Why did we create this? (Isn't `go test` good enough?)



We weren't satisifed with the built-in GoLang test tools. No, actually we were overjoyed that the language came with something built-in. And we liked `go test` enough that rather than create something from the ground up we decided to integrate with `go test` directly. We were just used to something much more descriptive and that facilitated testing large systems withing a lot of boiler plate code.

​	我们对 Go 的内置测试工具并不完全满意。准确地说，我们很高兴语言自带了测试功能。而且我们非常喜欢 `go test`，以至于决定直接与 `go test` 集成，而不是从头开始创建工具。我们习惯于更具描述性的工具，它们更适合大系统的测试并减少了样板代码。

#### 为什么叫 GoConvey？Why is it called GoConvey?



We've used a few different BDD tools before, each having its own take on the language you should use to specify the behavior of the system under test. "Given, When, Then" vs. "Establish, Because, It" from BDD and "Arrange, Act, Assert" from TDD are a few examples. In the end, you use the testing tool to specify or "convey" what the system should do. So, the main function you'll use to write specifications is named `Convey`. The language you use to specify your system is up to you, although we usually use the "Given, When, Then" style.

​	我们使用过几种不同的 BDD 工具，每种工具对测试语言有自己的定义，例如 "Given, When, Then"（BDD），"Establish, Because, It"（BDD），或 "Arrange, Act, Assert"（TDD）。最终，您使用测试工具来描述或“传达”（Convey）系统的行为。所以，您用来编写测试的主要函数被命名为 `Convey`。具体使用哪种语言取决于您自己，我们通常使用 "Given, When, Then" 的风格。

#### 我是否需要避免在代码中生成某些特定的输出？Should I be careful to not produce certain output in my code?



Well, that's an obscure question if I've ever heard one, but I'm glad you asked. Don't ever output either of these patterns:

​	这是一个很少见但非常重要的问题。请避免输出以下模式：

1. `>->->OPEN-JSON->->->`
2. `<-<-<-CLOSE-JSON<-<-<`

Those patterns are used to delimit blocks of JSON so that the web server can parse test output correctly.

​	这些模式被用作 JSON 块的分隔符，以便 Web 服务器正确解析测试输出。
