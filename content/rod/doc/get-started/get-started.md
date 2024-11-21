+++
title = "开始使用 Rod"
date = 2024-11-21T08:04:51+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-rod.github.io/i18n/zh-CN/#/get-started/README](https://go-rod.github.io/i18n/zh-CN/#/get-started/README)
>
> 收录该文档时间： `2024-11-21T08:08:39+08:00`

# 开始使用 Rod

## 依赖

​	[Golang](https://golang.org/) 是唯一的依赖，你甚至完全不需要了解 HTML。

​	如果你从未使用过 Golang，可以依照[这篇文档](https://golang.org/doc/install)安装 Golang，并通过[这个网站](https://tour.golang.org/welcome)在几个小时内掌握它。

## 第一个程序

​	让我们使用 Rod 来打开一个网页并获取它的截图。 首先创建 "main.go"，并在其中输入以下内容：

```go
package main

import "github.com/go-rod/rod"

func main() {
    page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")
    page.MustWaitStable().MustScreenshot("a.png")
}
```

​	`rod.New` 用于创建浏览器对象，而 `MustConnect` 则会启动并连接到浏览器。 `MustPage` 会创建一个页面对象（类似于浏览器中的一个标签页）。 `MustWaitStable` 等到页面几乎没有变化。 `MustScreenshot` 会获取页面的截图。

​	创建一个 module：

```bash
go env -w GOPROXY=https://goproxy.io,direct
go mod init learn-rod
go mod tidy
```

​	运行这个 module：

```bash
go run .
```

​	程序会输出如下的一张截图 "a.png"。

![first-program](get-started_img/first-program.png)

## xxxxxxxxxx15 1func main() {2    _, err := page.Element("a")3    handleError(err)4}5​6func handleError(err error) {7    var evalErr *rod.EvalError8    if errors.Is(err, context.DeadlineExceeded) { // 超时错误9        fmt.Println("timeout err")10    } else if errors.As(err, &evalErr) { // eval 错误11        fmt.Println(evalErr.LineNumber)12    } else if err != nil {13        fmt.Println("can't handle", err)14    }15}go

​	对于有经验的开发者，可以跳过这里的所有内容、阅读[这个文件](https://github.com/go-rod/rod/blob/main/examples_test.go)。

​	默认情况下，Rod 会禁用浏览器的 UI 来最大化性能。 但开发自动化任务时我们通常更加关心调试的难易程度。 Rod 提供了很多用于提升调试体验帮助函数。

​	在我们再次运行模块之前，让我们稍微修改代码，以便于调试：

```go
package main

import (
    "time"

    "github.com/go-rod/rod"
)

func main() {
    page := rod.New().NoDefaultDevice().MustConnect().MustPage("https://www.wikipedia.org/")
    page.MustWindowFullscreen()
    page.MustWaitStable().MustScreenshot("a.png")
    time.Sleep(time.Hour)
}
```

​	`NoDefaultDevice` 和 `MustWindowFullscreen` 将最大化页面视图和浏览器窗口，使其更容易调试。 我们在代码结尾添加了 `time.Sleep(time.Hour)` ，这样程序就不会在肉眼能察觉前太快退出。

​	让我们用 `-rod` 命令行参数再次运行模块：

```bash
go run . -rod=show
```

​	`show` 选项的意思是“在前景中显示浏览器界面”。 现在你应该像这样看到浏览器：

![show](get-started_img/show.png)

​	要停止程序，让我们回到终端，然后按键盘上的 [CTRL + C](https://en.wikipedia.org/wiki/Control-C)。

## 输入与点击

​	让我们控制浏览器来搜索关键词“earth”。 一个网站可能有许多输入框和按钮。 我们需要告诉程序它需要操控其中的哪一个。 通常我们会使用 [Devtools](https://developers.google.com/web/tools/chrome-devtools/) 来帮助定位我们想要控制的元素。 让我们在 `-rod` 参数重添加一个新的配置来启用 Devtools，现在命令变成了：

```bash
go run . -rod=show,devtools
```

​	运行上面的命令，将鼠标移动到输入框，在上面点击右键，然后在弹出的菜单中点击“审查元素”：

![inspect](get-started_img/inspect.png)

​	你会看到如下的 `<input id="searchInput`。

![input](get-started_img/input.png)

​	如上图所示，右击复制 [css 选择器](https://go-rod.github.io/i18n/zh-CN/#/css-selector)。 剪贴板中的内容会变成“#searchInput”。 我们之后会使用它来定位用于输入关键字的元素。 现在“main.go”中的内容变为：

```go
package main

import (
    "time"

    "github.com/go-rod/rod"
)

func main() {
    browser := rod.New().MustConnect().NoDefaultDevice()
    page := browser.MustPage("https://www.wikipedia.org/").MustWindowFullscreen()

    page.MustElement("#searchInput").MustInput("earth")

    page.MustWaitStable().MustScreenshot("a.png")
    time.Sleep(time.Hour)
}
```

​	我们使用 `MustElement` 与先前从 Devtools 面板复制的选择器来获取我们想要控制的元素。 `MustElement` 会自动等待直到元素出现为止，所以我们不需要在它之前使用 `MustWaitStable`。 然后我们调用 `MustInput` 来输入关键词“earth”。 再次运行“main.go”后你会看到如下的结果：

![after-input](get-started_img/after-input.png)

​	让我们用类似的方法，右击搜索按钮，复制它的选择器：

![search-btn](get-started_img/search-btn.png)

![search-btn-selector](get-started_img/search-btn-selector.png)

​	然后添加代码来点击这个搜索按钮。 现在“main.go”的内容是：

```go
package main

import (
    "time"

    "github.com/go-rod/rod"
)

func main() {
    browser := rod.New().MustConnect().NoDefaultDevice()
    page := browser.MustPage("https://www.wikipedia.org/").MustWindowFullscreen()

    page.MustElement("#searchInput").MustInput("earth")
    page.MustElement("#search-form > fieldset > button").MustClick()

    page.MustWaitStable().MustScreenshot("a.png")
    time.Sleep(time.Hour)
}
```

​	如果我们重新运行这个模块，“a.png”会显示搜索结果：

![earth-page](get-started_img/earth-page.png)

## 慢动作和可视化跟踪

​	自动化操作对人眼来说太快了，调试时我们通常会启用慢动作和可视化跟踪。 让我们用些额外的配置来运行这个模块：

```bash
go run . -rod="show,slow=1s,trace"
```

​	现在每次操作都会在执行前等待 1 秒。 在页面上，你会看到 Rod 生成的如下的可视化跟踪：

![trace](get-started_img/trace.png)

​	如图所示，Rod 会在搜索按钮上创建一个虚拟的鼠标光标。

​	在控制台中，你会看到如下的跟踪日志：

```txt
[rod] 2020/11/11 11:11:11 [eval] {"js":"rod.element","params":["#searchInput"]}
[rod] 2020/11/11 11:11:11 [eval] {"js":"rod.visible","this":"input#searchInput"}
[rod] 2020/11/11 11:11:11 [input] scroll into view
[rod] 2020/11/11 11:11:11 [input] input earth
[rod] 2020/11/11 11:11:11 [eval] {"js":"rod.element","params":["#search-form > fieldset > button"]}
[rod] 2020/11/11 11:11:11 [eval] {"js":"rod.visible","this":"button.pure-button.pure-button-primary-progressive"}
[rod] 2020/11/11 11:11:11 [input] scroll into view
[rod] 2020/11/11 11:11:11 [input] left click
```

## 命令行选项以外的其它选项

​	命令行参数只是一些常用方法的快捷方式。 你也可以在代码中手动设置，比如“slow”可以通过 `rod.New().SlowMotion(2 * time.Second)` 这样的代码来实现。

## 获取文本内容

Rod 提供了许多方便的方法来获取页面中的内容。

让我们试着来获取关于 Earth 的说明，依然和先前一样通过 Devtools 来复制 CSS 选择器：

![get-text](get-started_img/get-text.png)

​	我们使用的方法是 `MustText`，下面是它的完整代码：

```go
package main

import (
    "fmt"

    "github.com/go-rod/rod"
)

func main() {
    page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")

    page.MustElement("#searchInput").MustInput("earth")
    page.MustElement("#search-form > fieldset > button").MustClick()

    el := page.MustElement("#mw-content-text > div.mw-parser-output > p:nth-child(6)")
    fmt.Println(el.MustText())
}
```

​	如果我们重新运行该模块，我们应该看到控制台输出类似：

```txt
Earth is the third planet from the Sun and the only astronomical object known to harbor life.
...
```

## 获取图片内容

​	与获取文本内容一样，我们也可以从页面中获取图像。 让我们找到 Earth 图像的 CSS 选择器，并使用 `MustResource` 获取图像的二进制数据：

![get-image](get-started_img/get-image.png)

​	完整的代码如下：

```go
package main

import (
    "github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/utils"
)

func main() {
    page := rod.New().MustConnect().MustPage("https://www.wikipedia.org/")

    page.MustElement("#searchInput").MustInput("earth")
    page.MustElement("#search-form > fieldset > button").MustClick()

    el := page.MustElement("#mw-content-text > div.mw-parser-output > table.infobox > tbody > tr:nth-child(1) > td > a > img")
    _ = utils.OutputFile("b.png", el.MustResource())
}
```

​	输出文件“b.png”应为：

![earth](get-started_img/earth.png)
