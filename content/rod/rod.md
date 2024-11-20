+++
title = "rod"
date = 2024-11-20T18:01:04+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/go-rod/rod](https://pkg.go.dev/github.com/go-rod/rod)
>
> 收录该文档时间：`2024-11-20T18:02:07+08:00`
>
> [Version: v0.116.2](https://pkg.go.dev/github.com/go-rod/rod?tab=versions)

Rod is a high-level driver directly based on [DevTools Protocol](https://chromedevtools.github.io/devtools-protocol). It's designed for web automation and scraping for both high-level and low-level use, senior developers can use the low-level packages and functions to easily customize or build up their own version of Rod, the high-level functions are just examples to build a default version of Rod.

​	Rod 是一个基于 [DevTools Protocol](https://chromedevtools.github.io/devtools-protocol) 的高级驱动。它专为高层和低层的网页自动化和抓取任务而设计。高级开发人员可以使用低层包和函数轻松定制或构建他们自己的 Rod 版本，高层函数只是构建默认 Rod 版本的示例。

**Features**

- Chained context design, intuitive to timeout or cancel the long-running task
  - 链式上下文设计，便于超时或取消长时间运行的任务

- Auto-wait elements to be ready
  - 自动等待元素准备好

- Debugging friendly, auto input tracing, remote monitoring headless browser
  - 调试友好，自动输入跟踪，远程监控无头浏览器

- Thread-safe for all operations
  - 所有操作均线程安全

- Automatically find or download [browser](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher)
  - 自动查找或下载 [浏览器](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher)

- High-level helpers like WaitStable, WaitRequestIdle, HijackRequests, WaitDownload, etc
  - 高级辅助功能，如 `WaitStable`、`WaitRequestIdle`、`HijackRequests`、`WaitDownload` 等

- Two-step WaitEvent design, never miss an event ([how it works](https://github.com/ysmood/goob))
  - 两步式事件等待设计，永不丢失事件（[工作原理](https://github.com/ysmood/goob)）

- Correctly handles nested iframes or shadow DOMs
  - 正确处理嵌套 iframe 或 shadow DOM

- No zombie browser process after the crash ([how it works](https://github.com/ysmood/leakless))
  - 崩溃后无僵尸浏览器进程（[工作原理](https://github.com/ysmood/leakless)）

- [CI](https://github.com/go-rod/rod/actions) enforced 100% test coverage
  - [CI](https://github.com/go-rod/rod/actions) 强制 100% 测试覆盖率


**Examples**

Please check the [examples_test.go](https://github.com/go-rod/rod/blob/v0.116.2/examples_test.go) file first, then check the [examples](https://github.com/go-rod/rod/blob/v0.116.2/lib/examples) folder.

​	请先查看 [examples_test.go](https://github.com/go-rod/rod/blob/v0.116.2/examples_test.go) 文件，然后查看 [examples](https://github.com/go-rod/rod/blob/v0.116.2/lib/examples) 文件夹。

For more detailed examples, please search the unit tests. Such as the usage of method `HandleAuth`, you can search all the `*_test.go` files that contain `HandleAuth`, for example, use Github online [search in repository](https://github.com/go-rod/rod/search?q=HandleAuth&unscoped_q=HandleAuth). You can also search the GitHub [issues](https://github.com/go-rod/rod/issues) or [discussions](https://github.com/go-rod/rod/discussions), a lot of usage examples are recorded there.

​	有关更详细的示例，请搜索单元测试。例如，使用 `HandleAuth` 方法时，可以搜索所有包含 `HandleAuth` 的 `*_test.go` 文件。例如，使用 GitHub 在线 [仓库搜索](https://github.com/go-rod/rod/search?q=HandleAuth&unscoped_q=HandleAuth)。也可以搜索 GitHub 的 [问题](https://github.com/go-rod/rod/issues) 或 [讨论](https://github.com/go-rod/rod/discussions)，其中记录了许多用法示例。

[Here](https://github.com/go-rod/rod/blob/v0.116.2/lib/examples/compare-chromedp) is a comparison of the examples between rod and Chromedp.

​	[这里](https://github.com/go-rod/rod/blob/v0.116.2/lib/examples/compare-chromedp) 是 Rod 和 Chromedp 示例的比较。

If you have questions, please raise an [issues](https://github.com/go-rod/rod/issues)/[discussions](https://github.com/go-rod/rod/discussions) or join the [chat room](https://discord.gg/CpevuvY).

​	如果有疑问，请提交 [问题](https://github.com/go-rod/rod/issues)/[讨论](https://github.com/go-rod/rod/discussions) 或加入 [聊天室](https://discord.gg/CpevuvY)。

**Join us**

Your help is more than welcome! Even just open an issue to ask a question may greatly help others.

​	欢迎您的帮助！即使只是打开一个问题来提问，也可能极大地帮助其他人。

Please read [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html) before you ask questions.

​	在提问前，请阅读 [如何聪明地提问](http://www.catb.org/~esr/faqs/smart-questions.html)。

We use Github Projects to manage tasks, you can see the priority and progress of the issues [here](https://github.com/go-rod/rod/projects).

​	我们使用 GitHub Projects 来管理任务，可以在 [这里](https://github.com/go-rod/rod/projects) 查看问题的优先级和进展。

If you want to contribute please read the [Contributor Guide](https://github.com/go-rod/rod/blob/v0.116.2/.github/CONTRIBUTING.md).

​	如果想要贡献，请阅读 [贡献指南](https://github.com/go-rod/rod/blob/v0.116.2/.github/CONTRIBUTING.md)。

# Overview 

Package rod is a high-level driver directly based on DevTools Protocol.

​	`rod` 包是一个直接基于 DevTools Protocol 的高级驱动。

## Example (基础 Basic)

This example opens https://github.com/, searches for "git", and then gets the header element which gives the description for Git.

​	该示例打开 https://github.com/，搜索 "git"，然后获取用于描述 Git 的标题元素。

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func main() {
	// Launch a new browser with default options, and connect to it.
    // 使用默认选项启动一个新浏览器，并连接到它。
	browser := rod.New().MustConnect()

	// Even you forget to close, rod will close it after main process ends.
    // 即使忘记关闭，Rod 也会在主进程结束后关闭它。
	defer browser.MustClose()

	// Create a new page
    // 创建一个新页面
	page := browser.MustPage("https://github.com").MustWaitStable()

	// Trigger the search input with hotkey "/"
    // 使用快捷键 "/" 触发搜索输入框
	page.Keyboard.MustType(input.Slash)

	// We use css selector to get the search input element and input "git"
    // 使用 CSS 选择器获取搜索输入元素并输入 "git"
	page.MustElement("#query-builder-test").MustInput("git").MustType(input.Enter)

	// Wait until css selector get the element then get the text content of it.
    // 等待直到 CSS 选择器找到元素，然后获取其文本内容
	text := page.MustElementR("span", "most widely used").MustText()

	fmt.Println(text)

	// Get all input elements. Rod supports query elements by css selector, xpath, and regex.
	// For more detailed usage, check the query_test.go file.
    // 获取所有输入元素。Rod 支持通过 CSS 选择器、XPath 和正则表达式查询元素。
	// 更详细的用法请查看 query_test.go 文件。
	fmt.Println("Found", len(page.MustElements("input")), "input elements")

	// Eval js on the page
    // 在页面上执行 JS
	page.MustEval(`() => console.log("hello world")`)

	// Pass parameters as json objects to the js function. This MustEval will result 3
    // 将参数作为 JSON 对象传递给 JS 函数。此 MustEval 将返回 3
	fmt.Println("1 + 2 =", page.MustEval(`(a, b) => a + b`, 1, 2).Int())

	// When eval on an element, "this" in the js is the current DOM element.
    // 当在元素上执行 JS 时，JS 中的 "this" 是当前 DOM 元素。
	fmt.Println(page.MustElement("title").MustEval(`() => this.innerText`).String())

}
Output:

Git is the most widely used version control system.
Found 10 input elements
1 + 2 = 3
Repository search results · GitHub
```
## Example (上下文与每个事件 Context_and_EachEvent)

``` go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func main() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://github.com").MustWaitLoad()

	page, cancel := page.WithCancel()

	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	// It's a blocking method, it will wait until the context is cancelled
    // 这是一个阻塞方法，会等待直到上下文被取消
	page.EachEvent(func(_ *proto.PageLifecycleEvent) {})()

	if page.GetContext().Err() == context.Canceled {
		fmt.Println("cancelled")
	}
}
Output:
```
## Example (上下文与超时 Context_and_timeout)

Rod use https://golang.org/pkg/context to handle cancellations for IO blocking operations, most times it's timeout. Context will be recursively passed to all sub-methods. For example, methods like Page.Context(ctx) will return a clone of the page with the ctx, all the methods of the returned page will use the ctx if they have IO blocking operations. [Page.Timeout](https://pkg.go.dev/github.com/go-rod/rod#Page.Timeout) or [Page.WithCancel](https://pkg.go.dev/github.com/go-rod/rod#Page.WithCancel) is just a shortcut for Page.Context. Of course, Browser or Element works the same way.

​	Rod 使用 [context](https://golang.org/pkg/context) 处理 IO 阻塞操作的取消，大多数情况下是超时。上下文将递归传递给所有子方法。例如，`Page.Context(ctx)` 方法返回一个带有 `ctx` 的页面克隆，如果返回页面的方法涉及 IO 阻塞操作，它们将使用该 `ctx`。[Page.Timeout](https://pkg.go.dev/github.com/go-rod/rod#Page.Timeout) 或 [Page.WithCancel](https://pkg.go.dev/github.com/go-rod/rod#Page.WithCancel) 是 `Page.Context` 的快捷方式。当然，浏览器或元素的工作方式相同。

``` go
package main

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://github.com")

	page.
		// Set a 5-second timeout for all chained methods
    	// 为所有链式方法设置 5 秒超时
		Timeout(5 * time.Second).

		// The total time for MustWaitLoad and MustElement must be less than 5 seconds
    	// MustWaitLoad 和 MustElement 的总耗时必须少于 5 秒
		MustWaitLoad().
		MustElement("title").

		// Methods after CancelTimeout won't be affected by the 5-second timeout
    	// CancelTimeout 后的方法将不受 5 秒超时的影响
		CancelTimeout().

		// Set a 10-second timeout for all chained methods
    	// 为所有链式方法设置 10 秒超时
		Timeout(10 * time.Second).

		// Panics if it takes more than 10 seconds
    	// 如果耗时超过 10 秒则会 panic
		MustText()

	// The two code blocks below are basically the same:
    // 下面两段代码块的功能基本相同：
	{
		page.Timeout(5 * time.Second).MustElement("a").CancelTimeout()
	}
	{
		// Use this way you can customize your own way to cancel long-running task
        // 这种方式允许自定义取消长时间运行的任务
		page, cancel := page.WithCancel()
		go func() {
			time.Sleep(time.Duration(rand.Int())) // cancel after randomly time 随机时间后取消
			cancel()
		}()
		page.MustElement("a")
	}
}
Output:
```
## Example (自定义浏览器启动 Customize_browser_launch)

Shows how we can further customize the browser with the launcher library. Usually you use launcher lib to set the browser's command line flags (switches). Doc for flags: https://peter.sh/experiments/chromium-command-line-switches

​	展示如何使用 launcher 库进一步自定义浏览器。通常使用 launcher 库设置浏览器的命令行标志（开关）。标志的文档：https://peter.sh/experiments/chromium-command-line-switches

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	url := launcher.New().
		Proxy("127.0.0.1:8080").     // set flag "--proxy-server=127.0.0.1:8080" 设置标志 "--proxy-server=127.0.0.1:8080"
		Delete("use-mock-keychain"). // delete flag "--use-mock-keychain" 删除标志 "--use-mock-keychain"
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	// So that we don't have to self issue certs for MITM
    // 让浏览器忽略证书错误，适合中间人攻击 (MITM) 测试
	browser.MustIgnoreCertErrors(true)

	// Adding authentication to the proxy, for the next auth request.
	// We use CLI tool "mitmproxy --proxyauth user:pass" as an example.
	go browser.MustHandleAuth("user", "pass")()

	// mitmproxy needs a cert config to support https. We use http here instead,
	// for example
    // mitmproxy 需要证书配置以支持 https，这里使用 http 作为示例
	fmt.Println(browser.MustPage("https://mdn.dev/").MustElement("title").MustText())
}
Output:
```
## Example (自定义重试策略 Customize_retry_strategy)

Shows how to change the retry/polling options that is used to query elements. This is useful when you want to customize the element query retry logic.

​	展示如何更改查询元素时使用的重试/轮询选项。这在需要自定义元素查询重试逻辑时非常有用。

``` go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
)

func main() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://github.com")

	// sleep for 0.5 seconds before every retry
    // 每次重试前休眠 0.5 秒
	sleeper := func() utils.Sleeper {
		return func(context.Context) error {
			time.Sleep(time.Second / 2)
			return nil
		}
	}
	el, _ := page.Sleeper(sleeper).Element("input")
	fmt.Println(el.MustProperty("name"))

	// If sleeper is nil page.ElementE will query without retrying.
	// If nothing found it will return an error.
    // 如果 Sleeper 为 nil，page.ElementE 将不重试查询。
	// 如果未找到内容，将返回错误。
	el, err := page.Sleeper(rod.NotFoundSleeper).Element("input")
	if errors.Is(err, &rod.ElementNotFoundError{}) {
		fmt.Println("element not found")
	} else if err != nil {
		panic(err)
	}

	fmt.Println(el.MustProperty("name"))

}
Output:

type
type
```
## Example (直接使用 CDP - Direct_cdp)

When rod doesn't have a feature that you need. You can easily call the cdp to achieve it. List of cdp API: https://github.com/go-rod/rod/tree/main/lib/proto

​	当 Rod 没有提供您需要的功能时，可以直接调用 CDP（Chrome DevTools Protocol）实现。CDP API 列表：https://github.com/go-rod/rod/tree/main/lib/proto

``` go
package main

import (
	"context"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func main() {
	page := rod.New().MustConnect().MustPage()

	// Rod doesn't have a method to enable AD blocking,
	// but you can call cdp interface directly to achieve it.
	// Rod 没有直接提供启用广告拦截的方法，
	// 但可以通过直接调用 CDP 接口实现。
    
	// The two code blocks below are equal to enable AD blocking
	// 以下两段代码块等效于启用广告拦截
	{
		_ = proto.PageSetAdBlockingEnabled{
			Enabled: true,
		}.Call(page)
	}

	{
		// Interact with the cdp JSON API directly
        // 直接与 CDP JSON API 交互
		_, _ = page.Call(context.TODO(), "", "Page.setAdBlockingEnabled", map[string]bool{
			"enabled": true,
		})
	}
}
Output:
```
## Example (禁用无头模式进行调试 Disable_headless_to_debug)

Shows how to disable headless mode and debug. Rod provides a lot of debug options, you can set them with setter methods or use environment variables. Doc for environment variables: https://pkg.go.dev/github.com/go-rod/rod/lib/defaults

​	展示如何禁用无头模式并进行调试。Rod 提供了许多调试选项，可以通过 setter 方法设置，也可以使用环境变量。环境变量文档：https://pkg.go.dev/github.com/go-rod/rod/lib/defaults

``` go
package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
)

func main() {
	// Headless runs the browser on foreground, you can also use flag "-rod=show"
	// Devtools opens the tab in each new tab opened automatically
    // 使用 Headless(false) 在前台运行浏览器，也可以使用标志 "-rod=show"
	// 使用 Devtools(true) 自动打开新标签页的开发者工具
	l := launcher.New().
		Headless(false).
		Devtools(true)

	defer l.Cleanup()

	url := l.MustLaunch()

	// Trace shows verbose debug information for each action executed
	// SlowMotion is a debug related function that waits 2 seconds between
	// each action, making it easier to inspect what your code is doing.
    // Trace 显示每个操作的详细调试信息
	// SlowMotion 是一个调试相关功能，在每个操作之间等待 2 秒，
	// 便于检查代码的执行情况
	browser := rod.New().
		ControlURL(url).
		Trace(true).
		SlowMotion(2 * time.Second).
		MustConnect()

	// ServeMonitor plays screenshots of each tab. This feature is extremely
	// useful when debugging with headless mode.
	// You can also enable it with flag "-rod=monitor"
    // ServeMonitor 可以为每个标签页播放截图。
	// 在无头模式调试时非常有用。也可以通过标志 "-rod=monitor" 启用
	launcher.Open(browser.ServeMonitor(""))

	defer browser.MustClose()

	page := browser.MustPage("https://github.com/")

	page.MustElement("input").MustInput("git").MustType(input.Enter)

	text := page.MustElement(".codesearch-results p").MustText()

	fmt.Println(text)

	utils.Pause() // pause goroutine 暂停 goroutine
}
Output:
```
## Example (下载文件 Download_file)

​	展示如何通过页面元素触发文件下载，并保存到本地。

``` go
package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
)

func main() {
	browser := rod.New().MustConnect()
	page := browser.MustPage("https://file-examples.com/index.php/sample-documents-download/sample-pdf-download/")

    // 等待下载开始
	wait := browser.MustWaitDownload()

    // 点击下载链接
	page.MustElementR("a", "DOWNLOAD SAMPLE PDF FILE").MustClick()

	_ = utils.OutputFile("t.pdf", wait())
}
Output:
```
## Example (错误处理 Error_handling)

We use "Must" prefixed functions to write example code. But in production you may want to use the no-prefix version of them. About why we use "Must" as the prefix, it's similar to https://golang.org/pkg/regexp/#MustCompile

​	示例代码使用 "Must" 前缀的函数，但在生产中，可能需要使用没有前缀的版本。关于为什么使用 "Must" 作为前缀，可以参考：https://golang.org/pkg/regexp/#MustCompile

``` go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://mdn.dev")

	// We use Go's standard way to check error types, no magic.
    // 使用 Go 标准方式检查错误类型，无魔法操作
	check := func(err error) {
		var evalErr *rod.EvalError
		if errors.Is(err, context.DeadlineExceeded) { // timeout error 超时错误
			fmt.Println("timeout err")
		} else if errors.As(err, &evalErr) { // eval error 执行错误
			fmt.Println(evalErr.LineNumber)
		} else if err != nil {
			fmt.Println("can't handle", err)
		}
	}

	// The two code blocks below are doing the same thing in two styles:
    // 以下两段代码块以不同风格实现相同功能：

	// The block below is better for debugging or quick scripting. We use panic to short-circuit logics.
	// So that we can take advantage of fluent interface (https://en.wikipedia.org/wiki/Fluent_interface)
	// and fail-fast (https://en.wikipedia.org/wiki/Fail-fast).
	// This style will reduce code, but it may also catch extra errors (less consistent and precise).
    // 下面的代码块更适合调试或快速脚本。我们使用 panic 来中断逻辑，
	// 这样我们可以利用流畅接口(https://en.wikipedia.org/wiki/Fluent_interface)和快速失败 (https://en.wikipedia.org/wiki/Fail-fast)。
	// 这种风格减少了代码量，但也可能捕获额外的错误（一致性和准确性较低）。
	{
		err := rod.Try(func() {
			fmt.Println(page.MustElement("a").MustHTML()) // use "Must" prefixed functions
		})
		check(err)
	}

	// The block below is better for production code. It's the standard way to handle errors.
	// Usually, this style is more consistent and precise.
    // 下面的代码块更适合生产代码。这是标准的错误处理方式。
	// 通常，这种风格更一致和准确。
	{
		el, err := page.Element("a")
		if err != nil {
			check(err)
			return
		}
		html, err := el.HTML()
		if err != nil {
			check(err)
			return
		}
		fmt.Println(html)
	}
}
Output:
```
## Example (复用远程对象 Eval_reuse_remote_object)

Shows how to share a remote object reference between two Eval.

​	展示如何在两个 `Eval` 中共享远程对象引用。

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	page := rod.New().MustConnect().MustPage()

	fn := page.MustEvaluate(rod.Eval(`() => Math.random`).ByObject())

	res := page.MustEval(`f => f()`, fn)

	// print a random number
    // 打印一个随机数
	fmt.Println(res.Num())
}
Output:
```
## Example (处理事件 Handle_events)

Shows how to listen for events.	

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func main() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage()

	done := make(chan struct{})

	// Listen for all events of console output.   
    // 监听所有控制台输出事件。
	go page.EachEvent(func(e *proto.RuntimeConsoleAPICalled) {
		if e.Type == proto.RuntimeConsoleAPICalledTypeLog {
			fmt.Println(page.MustObjectsToJSON(e.Args))
			close(done)
		}
	})()

	wait := page.WaitEvent(&proto.PageLoadEventFired{})
	page.MustNavigate("https://mdn.dev")
	wait()

	// EachEvent allows us to achieve the same functionality as above.
    // EachEvent 允许实现与上面相同的功能。
    
	if false {
		// Subscribe events before they happen, run the "wait()" to start consuming
		// the events. We can return an optional stop signal to unsubscribe events.
        // 在事件发生前订阅事件，调用 "wait()" 开始消费事件。
		// 可以返回一个可选的停止信号来取消订阅事件。
		wait := page.EachEvent(func(_ *proto.PageLoadEventFired) (stop bool) {
			return true
		})
		page.MustNavigate("https://mdn.dev")
		wait()
	}

	// Or the for-loop style to handle events to do the same thing above.
    // 或者使用 for 循环的方式处理事件，效果相同。
	if false {
		page.MustNavigate("https://mdn.dev")

		for msg := range page.Event() {
			e := proto.PageLoadEventFired{}
			if msg.Load(&e) {
				break
			}
		}
	}

	page.MustEval(`() => console.log("hello", "world")`)

	<-done

}
Output:

[hello world]
```
## Example (劫持请求 Hijack_requests)

Shows how to intercept requests and modify both the request and the response. The entire process of hijacking one request:

​	展示如何拦截请求并修改请求和响应。劫持一个请求的完整过程如下：

```
browser --req-> rod ---> server ---> rod --res-> browser
```

The `--req->` and `--res->` are the parts that can be modified.

​	`--req->` 和 `--res->` 部分可以被修改。

``` go
package main

import (
	"fmt"
	"net/http"

	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	router := browser.HijackRequests()
	defer router.MustStop()

	router.MustAdd("*.js", func(ctx *rod.Hijack) {
		// Here we update the request's header. Rod gives functionality to
		// change or update all parts of the request. Refer to the documentation
		// for more information.
        // 在此处更新请求的头部。Rod 提供了修改请求所有部分的功能。
		// 参考文档以获取更多信息。
		ctx.Request.Req().Header.Set("My-Header", "test")

		// LoadResponse runs the default request to the destination of the request.
		// Not calling this will require you to mock the entire response.
		// This can be done with the SetXxx (Status, Header, Body) functions on the
		// ctx.Response struct.
        // LoadResponse 执行请求的默认目标。
		// 如果不调用此方法，则需要完全模拟响应。
		// 可以通过 ctx.Response 的 SetXxx (Status, Header, Body) 方法完成。
		_ = ctx.LoadResponse(http.DefaultClient, true)

		// Here we append some code to every js file.
		// The code will update the document title to "hi"
        // 在每个 JS 文件中追加一些代码。
		// 代码将更新文档标题为 "hi"。
		ctx.Response.SetBody(ctx.Response.Body() + "\n document.title = 'hi' ")
	})

	go router.Run()

	browser.MustPage("https://go-rod.github.io").MustWait(`() => document.title === 'hi'`)

	fmt.Println("done")

}
Output:

done
```
## Example (加载扩展 Load_extension)

``` go
package main

import (
	"fmt"
	"path/filepath"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	extPath, _ := filepath.Abs("fixtures/chrome-extension")

	u := launcher.New().
		// Must use abs path for an extension
    	// 必须使用扩展的绝对路径
		Set("load-extension", extPath).
		// Headless mode doesn't support extension yet.
    	// 无头模式尚不支持扩展。
		// 原因 Reason: https://bugs.chromium.org/p/chromium/issues/detail?id=706008#c5
		// You can use XVFB to get rid of it: https://github.com/go-rod/rod/blob/main/lib/examples/launch-managed/main.go
    	// 可以使用 XVFB 解决：https://github.com/go-rod/rod/blob/main/lib/examples/launch-managed/main.go
		Headless(false).
		MustLaunch()

	page := rod.New().ControlURL(u).MustConnect().MustPage("http://mdn.dev")

	page.MustWait(`() => document.title === 'test-extension'`)

	fmt.Println("ok")

	// Skip
}	
Output:
```
## Example (记录 CDP 流量 Log_cdp_traffic)

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/cdp"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
)

func main() {
	cdp := cdp.New().
		// Here we can customize how to log the requests, responses, and events transferred between Rod and the browser.
    	// 在这里可以自定义如何记录在 Rod 和浏览器之间传输的请求、响应和事件。
		Logger(utils.Log(func(args ...interface{}) {
			switch v := args[0].(type) {
			case *cdp.Request:
				fmt.Printf("id: %d", v.ID)
			}
		})).
		Start(cdp.MustConnectWS(launcher.New().MustLaunch()))

	rod.New().Client(cdp).MustConnect().MustPage("http://mdn.dev")
}
Output:
```
## Example (页面PDF化 - Page_pdf)

``` go
package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://github.com").MustWaitLoad()

	// simple version
    // 简单版本
	page.MustPDF("my.pdf")

	// customized version
    // 自定义版本	
	pdf, _ := page.PDF(&proto.PagePrintToPDF{
		PaperWidth:  gson.Num(8.5),
		PaperHeight: gson.Num(11),
		PageRanges:  "1-3",
	})
	_ = utils.OutputFile("my.pdf", pdf)
}
Output:

```
## Example (页面截图 Page_screenshot)

``` go
package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://github.com").MustWaitLoad()

	// simple version
    // 简单版本
	page.MustScreenshot("my.png")

	// customization version
    // 自定义版本
	img, _ := page.Screenshot(true, &proto.PageCaptureScreenshot{
		Format:  proto.PageCaptureScreenshotFormatJpeg,
		Quality: gson.Int(90),
		Clip: &proto.PageViewport{
			X:      0,
			Y:      0,
			Width:  300,
			Height: 200,
			Scale:  1,
		},
		FromSurface: true,
	})
	_ = utils.OutputFile("my.jpg", img)
}
Output:

```
## Example (页面滚动截图 Page_scroll_screenshot)

``` go
package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

func main() {
	browser := rod.New().MustConnect()

	// capture entire browser viewport, returning jpg with quality=90
    // 捕获整个浏览器视口，返回 JPG 格式，质量为 90
	img, err := browser.MustPage("https://desktop.github.com/").MustWaitStable().ScrollScreenshot(&rod.ScrollScreenshotOptions{
		Format:  proto.PageCaptureScreenshotFormatJpeg,
		Quality: gson.Int(90),
	})
	if err != nil {
		panic(err)
	}

	_ = utils.OutputFile("my.jpg", img)
}
Output:
```
## Example (竞争选择器 Race_selectors)

Show how to handle multiple results of an action. Such as when you login a page, the result can be success or wrong password.

​	展示如何处理操作的多个结果。例如登录页面时，可能的结果是登录成功或密码错误。

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func main() {
	const username = ""
	const password = ""

	browser := rod.New().MustConnect()

	page := browser.MustPage("https://leetcode.com/accounts/login/")

	page.MustElement("#id_login").MustInput(username)
	page.MustElement("#id_password").MustInput(password).MustType(input.Enter)

	// It will keep retrying until one selector has found a match
    // 它会不断重试，直到找到一个匹配的选择器
	elm := page.Race().Element(".nav-user-icon-base").MustHandle(func(e *rod.Element) {
		// print the username after successful login
        // 登录成功后打印用户名
		fmt.Println(*e.MustAttribute("title"))
	}).Element("[data-cy=sign-in-error]").MustDo()

	if elm.MustMatches("[data-cy=sign-in-error]") {
		// when wrong username or password
        // 如果用户名或密码错误
		panic(elm.MustText())
	}
}
Output:
```
## Example (搜索 Search)

Example_search shows how to use Search to get element inside nested iframes or shadow DOMs. It works the same as https://developers.google.com/web/tools/chrome-devtools/dom#search

​	展示如何使用 `Search` 获取嵌套 iframe 或 shadow DOM 内的元素。用法与 [Chrome DevTools DOM 搜索](https://developers.google.com/web/tools/chrome-devtools/dom#search) 相同。

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe")

	// Click the zoom-in button of the OpenStreetMap
    // 点击 OpenStreetMap 的放大按钮
	page.MustSearch(".leaflet-control-zoom-in").MustClick()

	fmt.Println("done")

}
Output:

done
```
## Example (状态 States)

Shows how to update the state of the current page. In this example we enable the network domain.

​	展示如何更新当前页面的状态。在此示例中，我们启用网络域。

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func main() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage()

	// LoadState detects whether the network domain is enabled or not.
    // LoadState 检测网络域是否已启用。
	fmt.Println(page.LoadState(&proto.NetworkEnable{}))

	_ = proto.NetworkEnable{}.Call(page)

	// Check if the network domain is successfully enabled.
    // 检查网络域是否成功启用。
	fmt.Println(page.LoadState(&proto.NetworkEnable{}))

}
Output:

false
true
```
## Example (等待动画完成 Wait_for_animation)

Rod uses mouse cursor to simulate clicks, so if a button is moving because of animation, the click may not work as expected. We usually use WaitStable to make sure the target isn't changing anymore.

​	Rod 使用鼠标光标模拟点击，因此如果按钮因动画而移动，点击可能无法如预期工作。我们通常使用 `WaitStable` 确保目标不再变化。

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://getbootstrap.com/docs/4.0/components/modal/")

	page.MustWaitLoad().MustElement("[data-target='#exampleModalLive']").MustClick()

	saveBtn := page.MustElementR("#exampleModalLive button", "Close")

	// Here, WaitStable will wait until the button's position and size become stable.
    // 此处，`WaitStable` 会等待按钮的位置和大小稳定。
	saveBtn.MustWaitStable().MustClick().MustWaitInvisible()

	fmt.Println("done")

}
Output:

done
```
## Example (等待请求完成 Wait_for_request)

When you want to wait for an ajax request to complete, this example will be useful.

​	当需要等待一个 AJAX 请求完成时，此示例会非常有用。

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://www.wikipedia.org/").MustWaitLoad()

	// Start to analyze request events
    // 开始分析请求事件
	wait := page.MustWaitRequestIdle()

	// This will trigger the search ajax request
    // 这将触发搜索的 AJAX 请求
	page.MustElement("#searchInput").MustClick().MustInput("lisp")

	// Wait until there's no active requests
    // 等待直到没有活动请求
	wait()

	// We want to make sure that after waiting, there are some autocomplete
	// suggestions available.
    // 确保等待后，有一些自动完成的建议可用。
	fmt.Println(len(page.MustElements(".suggestion-link")) > 0)

}
Output:

true
```
## 常量

This section is empty.

## 变量 

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/utils.go#L67)

``` go
var DefaultLogger = log.New(os.Stdout, "[rod] ", log.LstdFlags)
```

DefaultLogger for rod.

​	`DefaultLogger` 是 Rod 的默认日志记录器。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/utils.go#L76)

``` go
var DefaultSleeper = func() utils.Sleeper {
	return utils.BackoffSleeper(100*time.Millisecond, time.Second, nil)
}
```

DefaultSleeper generates the default sleeper for retry, it uses backoff to grow the interval. The growth looks like:

​	`DefaultSleeper` 生成默认的重试等待器，使用回退机制增长间隔。增长形式如下：

```go
A(0) = 100ms, A(n) = A(n-1) * random[1.9, 2.1), A(n) < 1s
```

Why the default is not RequestAnimationFrame or DOM change events is because of if a retry never ends it can easily flood the program. But you can always easily config it into what you want.

​	为什么默认值不是 `RequestAnimationFrame` 或 DOM 变化事件？因为如果重试永远不会结束，很容易使程序过载。但您始终可以轻松配置成所需的形式。

## 函数 

## func NotFoundSleeper <- 0.88.9

``` go
func NotFoundSleeper() utils.Sleeper
```

NotFoundSleeper returns ErrElementNotFound on the first call.

​	`NotFoundSleeper` 在首次调用时返回 `ErrElementNotFound`。

## func Try <- 0.46.0

``` go
func Try(fn func()) (err error)
```

Try try fn with recover, return the panic as rod.ErrTry.

​	`Try` 尝试运行 `fn` 并捕获 `panic`，将其作为 `rod.ErrTry` 返回。

## 类型

### type Browser 

``` go
type Browser struct {
	// BrowserContextID is the id for incognito window
    // BrowserContextID 是无痕窗口的 ID
	BrowserContextID proto.BrowserBrowserContextID
	// contains filtered or unexported fields
}
```

Browser represents the browser. It doesn't depends on file system, it should work with remote browser seamlessly. To check the env var you can use to quickly enable options from CLI, check here: https://pkg.go.dev/github.com/go-rod/rod/lib/defaults

​	`Browser` 表示浏览器。它不依赖于文件系统，应当能够无缝运行于远程浏览器环境中。要快速启用 CLI 中的选项，可以查看环境变量说明：[defaults](https://pkg.go.dev/github.com/go-rod/rod/lib/defaults)。

### Example (浏览器池 Pool)

We can use [rod.BrowserPool](https://pkg.go.dev/github.com/go-rod/rod#BrowserPool) to concurrently control and reuse browsers.

​	我们可以使用 [rod.BrowserPool](https://pkg.go.dev/github.com/go-rod/rod#BrowserPool) 来并发控制和重用浏览器实例。

``` go
package main

import (
	"fmt"
	"sync"

	"github.com/go-rod/rod"
)

func main() {
	// Create a new browser pool with a limit of 3
    // 创建一个最大容量为 3 的浏览器池
	pool := rod.NewBrowserPool(3)

	// Create a function that returns a new browser instance
    // 创建一个返回新浏览器实例的函数
	create := func() *rod.Browser {
		browser := rod.New().MustConnect()
		return browser
	}

	// Use the browser instances in separate goroutines
    // 在不同的 goroutine 中使用浏览器实例
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Get a browser instance from the pool
            // 从浏览器池中获取一个浏览器实例
			browser := pool.MustGet(create)

			// Put the instance back to the pool after we're done,
			// so the instance can be reused by other goroutines.
            // 操作完成后，将实例放回浏览器池中
			// 这样其他 goroutine 就可以重用该实例
			defer pool.Put(browser)

			// Use the browser instance
            // 使用浏览器实例
			page := browser.MustPage("https://www.google.com")
			fmt.Println(page.MustInfo().Title)
		}()
	}

	// Wait for all the goroutines to finish
    // 等待所有 goroutine 完成
	wg.Wait()

	// Cleanup the pool by closing all the browser instances
    // 清理浏览器池，关闭所有浏览器实例
	pool.Cleanup(func(p *rod.Browser) {
		p.MustClose()
	})
}
Output:
```
#### func New 

``` go
func New() *Browser
```

New creates a controller. DefaultDevice to emulate is set to [devices.LaptopWithMDPIScreen](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/devices#LaptopWithMDPIScreen).Landscape(), it will change the default user-agent and can make the actual view area smaller than the browser window on headful mode, you can use [Browser.NoDefaultDevice](https://pkg.go.dev/github.com/go-rod/rod#Browser.NoDefaultDevice) to disable it.

​	`New` 创建一个浏览器控制器。默认模拟的设备是 [devices.LaptopWithMDPIScreen](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/devices#LaptopWithMDPIScreen).Landscape()，这会更改默认的 `user-agent`，并在启用界面模式时使实际的视图区域小于浏览器窗口。可以使用 [Browser.NoDefaultDevice](https://pkg.go.dev/github.com/go-rod/rod#Browser.NoDefaultDevice) 禁用该默认设置。

#### (*Browser) Call 

``` go
func (b *Browser) Call(ctx context.Context, sessionID, methodName string, params interface{}) (res []byte, err error)
```

Call implements the [proto.Client](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#Client) to call raw cdp interface directly.

​	`Call` 实现了 [proto.Client](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#Client)，可直接调用底层的 CDP 接口。

#### (*Browser) CancelTimeout 

``` go
func (b *Browser) CancelTimeout() *Browser
```

CancelTimeout cancels the current timeout context and returns a clone with the parent context.

​	`CancelTimeout` 取消当前超时上下文，并返回一个包含父上下文的克隆。

#### (*Browser) Client 

``` go
func (b *Browser) Client(c CDPClient) *Browser
```

Client set the cdp client.

​	`Client` 设置 CDP 客户端。

#### (*Browser) Close 

``` go
func (b *Browser) Close() error
```

Close the browser.

​	`Close` 关闭浏览器。

#### (*Browser) Connect 

``` go
func (b *Browser) Connect() error
```

Connect to the browser and start to control it. If fails to connect, try to launch a local browser, if local browser not found try to download one.

​	`Connect` 连接到浏览器并开始控制。如果连接失败，尝试启动本地浏览器；如果本地浏览器未找到，则尝试下载。

#### (*Browser) Context 

``` go
func (b *Browser) Context(ctx context.Context) *Browser
```

Context returns a clone with the specified ctx for chained sub-operations.

​	`Context` 返回一个包含指定 `ctx` 的克隆，用于链接的子操作。

#### (*Browser) ControlURL 

``` go
func (b *Browser) ControlURL(url string) *Browser
```

ControlURL set the url to remote control browser.

​	`ControlURL` 设置用于远程控制浏览器的 URL。

#### (*Browser) DefaultDevice <- 0.71.0

``` go
func (b *Browser) DefaultDevice(d devices.Device) *Browser
```

DefaultDevice sets the default device for new page to emulate in the future. Default is [devices.LaptopWithMDPIScreen](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/devices#LaptopWithMDPIScreen). Set it to [devices.Clear](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/devices#Clear) to disable it.

​	`DefaultDevice` 设置新页面默认模拟的设备。默认值为 [devices.LaptopWithMDPIScreen](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/devices#LaptopWithMDPIScreen)。设置为 [devices.Clear](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/devices#Clear) 可禁用设备模拟。

#### (*Browser) DisableDomain 

``` go
func (b *Browser) DisableDomain(sessionID proto.TargetSessionID, req proto.Request) (restore func())
```

DisableDomain and returns a restore function to restore previous state.

​	`DisableDomain` 禁用指定域，并返回一个恢复函数以恢复先前的状态。

#### (*Browser) EachEvent 

``` go
func (b *Browser) EachEvent(callbacks ...interface{}) (wait func())
```

EachEvent is similar to [Page.EachEvent](https://pkg.go.dev/github.com/go-rod/rod#Page.EachEvent), but catches events of the entire browser.

​	`EachEvent` 类似于 [Page.EachEvent](https://pkg.go.dev/github.com/go-rod/rod#Page.EachEvent)，但会捕获整个浏览器的事件。

#### (*Browser) EnableDomain 

``` go
func (b *Browser) EnableDomain(sessionID proto.TargetSessionID, req proto.Request) (restore func())
```

EnableDomain and returns a restore function to restore previous state.

​	`EnableDomain` 启用指定域，并返回一个恢复函数以恢复先前的状态。

#### (*Browser) Event 

``` go
func (b *Browser) Event() <-chan *Message
```

Event of the browser.

​	`Event` 获取浏览器的事件通道。

#### (*Browser) GetContext 

``` go
func (b *Browser) GetContext() context.Context
```

GetContext of current instance.

​	`GetContext` 获取当前实例的上下文。

#### (*Browser) GetCookies <- 0.71.0

``` go
func (b *Browser) GetCookies() ([]*proto.NetworkCookie, error)
```

GetCookies from the browser.

​	`GetCookies` 从浏览器获取所有的 Cookie。

#### (*Browser) HandleAuth 

``` go
func (b *Browser) HandleAuth(username, password string) func() error
```

HandleAuth for the next basic HTTP authentication. It will prevent the popup that requires user to input user name and password. Ref: https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication

​	`HandleAuth` 处理下一个基本 HTTP 验证。它会防止弹出需要用户输入用户名和密码的窗口。参考：[MDN HTTP Authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication)。

#### (*Browser) HijackRequests 

``` go
func (b *Browser) HijackRequests() *HijackRouter
```

HijackRequests same as Page.HijackRequests, but can intercept requests of the entire browser.

​	`HijackRequests` 类似于 `Page.HijackRequests`，但可以拦截整个浏览器的请求。

#### (*Browser) IgnoreCertErrors <- 0.61.3

``` go
func (b *Browser) IgnoreCertErrors(enable bool) error
```

IgnoreCertErrors switch. If enabled, all certificate errors will be ignored.

​	`IgnoreCertErrors` 开关。如果启用，将忽略所有证书错误。

#### (*Browser) Incognito 

``` go
func (b *Browser) Incognito() (*Browser, error)
```

Incognito creates a new incognito browser.

​	`Incognito` 创建一个新的无痕浏览器实例。

#### (*Browser) LoadState 

``` go
func (b *Browser) LoadState(sessionID proto.TargetSessionID, method proto.Request) (has bool)
```

LoadState into the method, sessionID can be empty.

​	`LoadState` 加载指定方法的状态。`sessionID` 可以为空。

#### (*Browser) Logger <- 0.70.0

``` go
func (b *Browser) Logger(l utils.Logger) *Browser
```

Logger overrides the default log functions for tracing.

​	`Logger` 覆盖默认日志函数，用于跟踪操作。

#### (*Browser) Monitor <- 0.70.0

``` go
func (b *Browser) Monitor(url string) *Browser
```

Monitor address to listen if not empty. Shortcut for [Browser.ServeMonitor](https://pkg.go.dev/github.com/go-rod/rod#Browser.ServeMonitor).

​	`Monitor` 设置监听地址（如果不为空）。这是 [Browser.ServeMonitor](https://pkg.go.dev/github.com/go-rod/rod#Browser.ServeMonitor) 的快捷方式。

#### (*Browser) MustClose <- 0.50.0

``` go
func (b *Browser) MustClose()
```

MustClose is similar to [Browser.Close](https://pkg.go.dev/github.com/go-rod/rod#Browser.Close).

​	`MustClose` 是 [Browser.Close](https://pkg.go.dev/github.com/go-rod/rod#Browser.Close) 的简化版本。

#### (*Browser) MustConnect <- 0.50.0

``` go
func (b *Browser) MustConnect() *Browser
```

MustConnect is similar to [Browser.Connect](https://pkg.go.dev/github.com/go-rod/rod#Browser.Connect).

​	`MustConnect` 是 [Browser.Connect](https://pkg.go.dev/github.com/go-rod/rod#Browser.Connect) 的简化版本。

#### (*Browser) MustGetCookies <- 0.71.0

``` go
func (b *Browser) MustGetCookies() []*proto.NetworkCookie
```

MustGetCookies is similar to [Browser.GetCookies](https://pkg.go.dev/github.com/go-rod/rod#Browser.GetCookies).

​	`MustGetCookies` 是 [Browser.GetCookies](https://pkg.go.dev/github.com/go-rod/rod#Browser.GetCookies) 的简化版本。

#### (*Browser) MustHandleAuth <- 0.50.0

``` go
func (b *Browser) MustHandleAuth(username, password string) (wait func())
```

MustHandleAuth is similar to [Browser.HandleAuth](https://pkg.go.dev/github.com/go-rod/rod#Browser.HandleAuth).

​	`MustHandleAuth` 是 [Browser.HandleAuth](https://pkg.go.dev/github.com/go-rod/rod#Browser.HandleAuth) 的简化版本。

#### (*Browser) MustIgnoreCertErrors <- 0.61.3

``` go
func (b *Browser) MustIgnoreCertErrors(enable bool) *Browser
```

MustIgnoreCertErrors is similar to [Browser.IgnoreCertErrors](https://pkg.go.dev/github.com/go-rod/rod#Browser.IgnoreCertErrors).

​	`MustIgnoreCertErrors` 是 [Browser.IgnoreCertErrors](https://pkg.go.dev/github.com/go-rod/rod#Browser.IgnoreCertErrors) 的简化版本。

#### (*Browser) MustIncognito <- 0.50.0

``` go
func (b *Browser) MustIncognito() *Browser
```

MustIncognito is similar to [Browser.Incognito](https://pkg.go.dev/github.com/go-rod/rod#Browser.Incognito).

​	`MustIncognito` 是 [Browser.Incognito](https://pkg.go.dev/github.com/go-rod/rod#Browser.Incognito) 的简化版本。

#### (*Browser) MustPage <- 0.50.0

``` go
func (b *Browser) MustPage(url ...string) *Page
```

MustPage is similar to [Browser.Page](https://pkg.go.dev/github.com/go-rod/rod#Browser.Page). The url list will be joined by "/".

​	`MustPage` 是 [Browser.Page](https://pkg.go.dev/github.com/go-rod/rod#Browser.Page) 的简化版本。URL 列表将通过 “/” 拼接。

#### (*Browser) MustPageFromTargetID <- 0.50.0

``` go
func (b *Browser) MustPageFromTargetID(targetID proto.TargetTargetID) *Page
```

MustPageFromTargetID is similar to [Browser.PageFromTargetID].

​	`MustPageFromTargetID` 是 [Browser.PageFromTargetID](https://pkg.go.dev/github.com/go-rod/rod#Browser.PageFromTargetID) 的简化版本。

#### (*Browser) MustPages <- 0.50.0

``` go
func (b *Browser) MustPages() Pages
```

MustPages is similar to [Browser.Pages](https://pkg.go.dev/github.com/go-rod/rod#Browser.Pages).

​	`MustPages` 是 [Browser.Pages](https://pkg.go.dev/github.com/go-rod/rod#Browser.Pages) 的简化版本。

#### (*Browser) MustSetCookies <- 0.71.0

``` go
func (b *Browser) MustSetCookies(cookies ...*proto.NetworkCookie) *Browser
```

MustSetCookies is similar to [Browser.SetCookies](https://pkg.go.dev/github.com/go-rod/rod#Browser.SetCookies). If the len(cookies) is 0 it will clear all the cookies.

​	`MustSetCookies` 是 [Browser.SetCookies](https://pkg.go.dev/github.com/go-rod/rod#Browser.SetCookies) 的简化版本。如果 `cookies` 长度为 0，则会清除所有的 Cookie。

#### (*Browser) MustVersion <- 0.107.0

``` go
func (b *Browser) MustVersion() *proto.BrowserGetVersionResult
```

MustVersion is similar to [Browser.Version](https://pkg.go.dev/github.com/go-rod/rod#Browser.Version).

​	`MustVersion` 是 [Browser.Version](https://pkg.go.dev/github.com/go-rod/rod#Browser.Version) 的简化版本。

#### (*Browser) MustWaitDownload <- 0.83.0

``` go
func (b *Browser) MustWaitDownload() func() []byte
```

MustWaitDownload is similar to [Browser.WaitDownload](https://pkg.go.dev/github.com/go-rod/rod#Browser.WaitDownload). It will read the file into bytes then remove the file.

​	`MustWaitDownload` 是 [Browser.WaitDownload](https://pkg.go.dev/github.com/go-rod/rod#Browser.WaitDownload) 的简化版本。它会将文件读取为字节后删除文件。

#### (*Browser) NoDefaultDevice <- 0.81.1

``` go
func (b *Browser) NoDefaultDevice() *Browser
```

NoDefaultDevice is the same as [Browser.DefaultDevice](https://pkg.go.dev/github.com/go-rod/rod#Browser.DefaultDevice)(devices.Clear).

​	`NoDefaultDevice` 等价于 [Browser.DefaultDevice](https://pkg.go.dev/github.com/go-rod/rod#Browser.DefaultDevice)(devices.Clear)。

#### (*Browser) Page 

``` go
func (b *Browser) Page(opts proto.TargetCreateTarget) (p *Page, err error)
```

Page creates a new browser tab. If opts.URL is empty, the default target will be "about:blank".

​	`Page` 创建一个新的浏览器标签页。如果 `opts.URL` 为空，默认目标将为 `about:blank`。

#### (*Browser) PageFromSession <- 0.74.0

``` go
func (b *Browser) PageFromSession(sessionID proto.TargetSessionID) *Page
```

PageFromSession is used for low-level debugging.

​	`PageFromSession` 用于低级调试。

#### (*Browser) PageFromTarget <- 0.50.0

``` go
func (b *Browser) PageFromTarget(targetID proto.TargetTargetID) (*Page, error)
```

PageFromTarget gets or creates a Page instance.

​	`PageFromTarget` 获取或创建一个 `Page` 实例。

#### (*Browser) Pages 

``` go
func (b *Browser) Pages() (Pages, error)
```

Pages retrieves all visible pages.

​	`Pages` 获取所有可见的页面。

#### (*Browser) RemoveState <- 0.74.0

``` go
func (b *Browser) RemoveState(key interface{})
```

RemoveState a state.

​	`RemoveState` 移除一个状态。

#### (*Browser) ServeMonitor 

``` go
func (b *Browser) ServeMonitor(host string) string
```

ServeMonitor starts the monitor server. The reason why not to use "chrome://inspect/#devices" is one target cannot be driven by multiple controllers.

​	`ServeMonitor` 启动监控服务器。未使用 "chrome://inspect/#devices" 的原因是一个目标不能被多个控制器驱动。

#### (*Browser) SetCookies <- 0.71.0

``` go
func (b *Browser) SetCookies(cookies []*proto.NetworkCookieParam) error
```

SetCookies to the browser. If the cookies is nil it will clear all the cookies.

​	`SetCookies` 将 Cookie 设置到浏览器中。如果 `cookies` 为 `nil`，它将清除所有 Cookie。

#### (*Browser) Sleeper <- 0.50.0

``` go
func (b *Browser) Sleeper(sleeper func() utils.Sleeper) *Browser
```

Sleeper returns a clone with the specified sleeper for chained sub-operations.

​	`Sleeper` 返回一个包含指定 `sleeper` 的克隆，用于链式子操作。

#### (*Browser) SlowMotion <- 0.77.0

``` go
func (b *Browser) SlowMotion(delay time.Duration) *Browser
```

SlowMotion set the delay for each control action, such as the simulation of the human inputs.

​	`SlowMotion` 设置每个控制操作的延迟，例如模拟人类输入时的操作。

#### (*Browser) Timeout 

``` go
func (b *Browser) Timeout(d time.Duration) *Browser
```

Timeout returns a clone with the specified total timeout of all chained sub-operations.

​	`Timeout` 返回一个包含指定总超时时间的克隆，用于所有链式子操作。

#### (*Browser) Trace 

``` go
func (b *Browser) Trace(enable bool) *Browser
```

Trace enables/disables the visual tracing of the input actions on the page.

​	`Trace` 启用或禁用页面上输入操作的可视化跟踪。

#### (*Browser) Version <- 0.107.0

``` go
func (b *Browser) Version() (*proto.BrowserGetVersionResult, error)
```

Version info of the browser.

​	`Version` 返回浏览器的版本信息。

#### (*Browser) WaitDownload <- 0.83.0

``` go
func (b *Browser) WaitDownload(dir string) func() (info *proto.PageDownloadWillBegin)
```

WaitDownload returns a helper to get the next download file. The file path will be:

​	`WaitDownload` 返回一个辅助工具，用于获取下一个下载的文件。文件路径将为：

```
filepath.Join(dir, info.GUID)
```

#### (*Browser) WaitEvent 

``` go
func (b *Browser) WaitEvent(e proto.Event) (wait func())
```

WaitEvent waits for the next event for one time. It will also load the data into the event object.

​	`WaitEvent` 等待下一个事件，仅等待一次。它还会将数据加载到事件对象中。

#### (*Browser) WithCancel <- 0.69.0

``` go
func (b *Browser) WithCancel() (*Browser, func())
```

WithCancel returns a clone with a context cancel function.

​	`WithCancel` 返回一个克隆，并带有上下文取消函数。

#### (*Browser) WithPanic <- 0.100.0

``` go
func (b *Browser) WithPanic(fail func(interface{})) *Browser
```

WithPanic returns a browser clone with the specified panic function. The fail must stop the current goroutine's execution immediately, such as use [runtime.Goexit](https://pkg.go.dev/runtime#Goexit) or panic inside it.

​	`WithPanic` 返回一个包含指定异常处理函数的浏览器克隆。`fail` 必须立即停止当前 Goroutine 的执行，例如使用 [runtime.Goexit](https://pkg.go.dev/runtime#Goexit) 或在内部触发 `panic`。

### type CDPClient <- 0.70.0

``` go
type CDPClient interface {
	Event() <-chan *cdp.Event
	Call(ctx context.Context, sessionID, method string, params interface{}) ([]byte, error)
}
```

CDPClient is usually used to make rod side-effect free. Such as proxy all IO of rod.

​	`CDPClient` 通常用于使 Rod 无副作用。例如代理 Rod 的所有 IO 操作。

### type CoveredError <- 0.114.8

``` go
type CoveredError struct {
	*Element
}
```

CoveredError error.

​	`CoveredError` 表示一个错误。

#### (*CoveredError) Error <- 0.114.8

``` go
func (e *CoveredError) Error() string
```

Error ...

​	`Error` 返回错误描述。

#### (*CoveredError) Is <- 0.114.8

``` go
func (e *CoveredError) Is(err error) bool
```

Is interface.

​	`Is` 检查错误接口。

#### (*CoveredError) Unwrap <- 0.114.8

``` go
func (e *CoveredError) Unwrap() error
```

Unwrap ...

​	`Unwrap` 返回嵌套的错误。

### type Element 

``` go
type Element struct {
	Object *proto.RuntimeRemoteObject
	// contains filtered or unexported fields
}
```

Element represents the DOM element.

​	`Element` 表示 DOM 元素。

#### (*Element) Attribute 

``` go
func (el *Element) Attribute(name string) (*string, error)
```

Attribute of the DOM object. Attribute vs Property: https://stackoverflow.com/questions/6003819/what-is-the-difference-between-properties-and-attributes-in-html

​	`Attribute` 返回 DOM 对象的属性值。关于属性和属性值的区别：https://stackoverflow.com/questions/6003819/what-is-the-difference-between-properties-and-attributes-in-html

#### (*Element) BackgroundImage <- 0.76.6

``` go
func (el *Element) BackgroundImage() ([]byte, error)
```

BackgroundImage returns the css background-image of the element.

​	`BackgroundImage` 返回元素的 CSS 背景图像。

#### (*Element) Blur 

``` go
func (el *Element) Blur() error
```

Blur removes focus from the element.

​	`Blur` 从元素中移除焦点。

#### (*Element) Call <- 0.70.0

``` go
func (el *Element) Call(ctx context.Context, sessionID, methodName string, params interface{}) (res []byte, err error)
```

Call implements the [proto.Client](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#Client).

​	`Call` 实现了 [proto.Client](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#Client)。

#### (*Element) CancelTimeout 

``` go
func (el *Element) CancelTimeout() *Element
```

CancelTimeout cancels the current timeout context and returns a clone with the parent context.

​	`CancelTimeout` 取消当前超时上下文并返回一个具有父上下文的克隆。

#### (*Element) CanvasToImage <- 0.45.1

``` go
func (el *Element) CanvasToImage(format string, quality float64) ([]byte, error)
```

CanvasToImage get image data of a canvas. The default format is image/png. The default quality is 0.92. doc: https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/toDataURL

​	`CanvasToImage` 获取画布的图像数据。默认格式为 `image/png`，默认质量为 `0.92`。文档：https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/toDataURL

#### (*Element) Click 

``` go
func (el *Element) Click(button proto.InputMouseButton, clickCount int) error
```

Click will press then release the button just like a human. Before the action, it will try to scroll to the element, hover the mouse over it, wait until the it's interactable and enabled.

​	`Click` 模拟人类操作，按下并释放鼠标按钮。在操作之前，它会尝试滚动到元素、将鼠标悬停在元素上，并等待元素可交互且已启用。

#### (*Element) ContainsElement <- 0.48.0

``` go
func (el *Element) ContainsElement(target *Element) (bool, error)
```

ContainsElement check if the target is equal or inside the element.

​	`ContainsElement` 检查目标元素是否等于或包含在当前元素中。

#### (*Element) Context 

``` go
func (el *Element) Context(ctx context.Context) *Element
```

Context returns a clone with the specified ctx for chained sub-operations.

​	`Context` 返回一个包含指定上下文的克隆，用于链式子操作。

#### (*Element) Describe 

``` go
func (el *Element) Describe(depth int, pierce bool) (*proto.DOMNode, error)
```

Describe the current element. The depth is the maximum depth at which children should be retrieved, defaults to 1, use -1 for the entire subtree or provide an integer larger than 0. The pierce decides whether or not iframes and shadow roots should be traversed when returning the subtree. The returned [proto.DOMNode.NodeID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMNode.NodeID) will always be empty, because NodeID is not stable (when [proto.DOMDocumentUpdated](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMDocumentUpdated) is fired all NodeID on the page will be reassigned to another value) we don't recommend using the NodeID, instead, use the [proto.DOMBackendNodeID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMBackendNodeID) to identify the element.

​	`Describe` 描述当前元素。`depth` 指定返回子节点的最大深度，默认为 1。使用 `-1` 返回整个子树，或提供大于 0 的整数。`pierce` 决定是否穿透 iframe 和 shadow DOM。返回的 [proto.DOMNode.NodeID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMNode.NodeID) 始终为空，因为 `NodeID` 不稳定（当触发 [proto.DOMDocumentUpdated](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMDocumentUpdated) 时，页面上的所有 `NodeID` 将重新分配为其他值）。建议使用 [proto.DOMBackendNodeID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMBackendNodeID) 代替 `NodeID` 来标识元素。

#### (*Element) Disabled <- 0.112.5

``` go
func (el *Element) Disabled() (bool, error)
```

Disabled checks if the element is disabled.

​	`Disabled` 检查元素是否被禁用。

#### (*Element) Element 

``` go
func (el *Element) Element(selector string) (*Element, error)
```

Element returns the first child that matches the css selector.

​	`Element` 返回第一个匹配 CSS 选择器的子元素。

#### (*Element) ElementByJS 

``` go
func (el *Element) ElementByJS(opts *EvalOptions) (*Element, error)
```

ElementByJS returns the element from the return value of the js.

​	`ElementByJS` 根据 JS 的返回值获取元素。

#### (*Element) ElementR <- 0.57.0

``` go
func (el *Element) ElementR(selector, jsRegex string) (*Element, error)
```

ElementR returns the first child element that matches the css selector and its text matches the jsRegex.

​	`ElementR` 返回第一个匹配 CSS 选择器且文本匹配 JS 正则表达式的子元素。

#### (*Element) ElementX 

``` go
func (el *Element) ElementX(xPath string) (*Element, error)
```

ElementX returns the first child that matches the XPath selector.

​	`ElementX` 返回第一个匹配 XPath 选择器的子元素。

#### (*Element) Elements 

``` go
func (el *Element) Elements(selector string) (Elements, error)
```

Elements returns all elements that match the css selector.

​	`Elements` 返回所有匹配 CSS 选择器的元素。

#### (*Element) ElementsByJS 

``` go
func (el *Element) ElementsByJS(opts *EvalOptions) (Elements, error)
```

ElementsByJS returns the elements from the return value of the js.

​	`ElementsByJS` 根据 JS 的返回值获取元素集合。

#### (*Element) ElementsX 

``` go
func (el *Element) ElementsX(xpath string) (Elements, error)
```

ElementsX returns all elements that match the XPath selector.

​	`ElementsX` 返回所有匹配 XPath 选择器的元素。

#### (*Element) Equal <- 0.85.7

``` go
func (el *Element) Equal(elm *Element) (bool, error)
```

Equal checks if the two elements are equal.

​	`Equal` 检查两个元素是否相等。

#### (*Element) Eval 

``` go
func (el *Element) Eval(js string, params ...interface{}) (*proto.RuntimeRemoteObject, error)
```

Eval is a shortcut for [Element.Evaluate](https://pkg.go.dev/github.com/go-rod/rod#Element.Evaluate) with AwaitPromise, ByValue and AutoExp set to true.

​	`Eval` 是 [Element.Evaluate](https://pkg.go.dev/github.com/go-rod/rod#Element.Evaluate) 的简化版本，默认启用 `AwaitPromise`、`ByValue` 和 `AutoExp`。

#### (*Element) Evaluate <- 0.67.0

``` go
func (el *Element) Evaluate(opts *EvalOptions) (*proto.RuntimeRemoteObject, error)
```

Evaluate is just a shortcut of [Page.Evaluate](https://pkg.go.dev/github.com/go-rod/rod#Page.Evaluate) with This set to current element.

​	`Evaluate` 是 [Page.Evaluate](https://pkg.go.dev/github.com/go-rod/rod#Page.Evaluate) 的简化版本，默认将 `This` 设置为当前元素。

#### (*Element) Focus 

``` go
func (el *Element) Focus() error
```

Focus sets focus on the specified element. Before the action, it will try to scroll to the element.

​	`Focus` 将焦点设置到指定元素上。在操作之前，它会尝试滚动到元素。

#### (*Element) Frame 

``` go
func (el *Element) Frame() (*Page, error)
```

Frame creates a page instance that represents the iframe.

​	`Frame` 创建一个表示 iframe 的页面实例。

#### (*Element) GetContext 

``` go
func (el *Element) GetContext() context.Context
```

GetContext of current instance.

​	`GetContext` 返回当前实例的上下文。

#### (*Element) GetSessionID <- 0.72.0

``` go
func (el *Element) GetSessionID() proto.TargetSessionID
```

GetSessionID interface.

​	`GetSessionID` 返回会话 ID。

#### (*Element) GetXPath <- 0.109.3

``` go
func (el *Element) GetXPath(optimized bool) (string, error)
```

GetXPath returns the xpath of the element.

​	`GetXPath` 返回元素的 XPath 表达式。

#### (*Element) HTML 

``` go
func (el *Element) HTML() (string, error)
```

HTML of the element.

​	`HTML` 返回元素的 HTML 内容。

#### (*Element) Has 

``` go
func (el *Element) Has(selector string) (bool, *Element, error)
```

Has an element that matches the css selector.

​	`Has` 检查是否存在匹配指定 CSS 选择器的子元素。

#### (*Element) HasR <- 0.61.0

``` go
func (el *Element) HasR(selector, jsRegex string) (bool, *Element, error)
```

HasR returns true if a child element that matches the css selector and its text matches the jsRegex.

​	`HasR` 返回是否存在匹配指定 CSS 选择器且其文本匹配 JS 正则表达式的子元素。

#### (*Element) HasX 

``` go
func (el *Element) HasX(selector string) (bool, *Element, error)
```

HasX an element that matches the XPath selector.

​	`HasX` 检查是否存在匹配指定 XPath 选择器的子元素。

#### (*Element) Hover <- 0.49.1

``` go
func (el *Element) Hover() error
```

Hover the mouse over the center of the element. Before the action, it will try to scroll to the element and wait until it's interactable.

​	`Hover` 将鼠标悬停在元素的中心。在操作之前，它会尝试滚动到元素，并等待元素可交互。

#### (*Element) Input 

``` go
func (el *Element) Input(text string) error
```

Input focuses on the element and input text to it. Before the action, it will scroll to the element, wait until it's visible, enabled and writable. To empty the input you can use something like

​	`Input` 将焦点设置到元素上并输入文本。在操作之前，它会滚动到元素并等待其可见、启用且可写入。如果需要清空输入，可以使用类似以下代码：

```go
el.SelectAllText().MustInput("")
```

#### (*Element) InputColor <- 0.114.3

``` go
func (el *Element) InputColor(color string) error
```

InputColor focuses on the element and inputs a color string to it. Before the action, it will scroll to the element, wait until it's visible, enabled and writable.

​	`InputColor` 将焦点设置到元素上并输入颜色字符串。在操作之前，它会滚动到元素并等待其可见、启用且可写入。

#### (*Element) InputTime <- 0.79.2

``` go
func (el *Element) InputTime(t time.Time) error
```

InputTime focuses on the element and input time to it. Before the action, it will scroll to the element, wait until it's visible, enabled and writable. It will wait until the element is visible, enabled and writable.

​	`InputTime` 将焦点设置到元素上并输入时间。在操作之前，它会滚动到元素并等待其可见、启用且可写入。

#### (*Element) Interactable <- 0.66.0

``` go
func (el *Element) Interactable() (pt *proto.Point, err error)
```

Interactable checks if the element is interactable with cursor. The cursor can be mouse, finger, stylus, etc. If not interactable err will be ErrNotInteractable, such as when covered by a modal,.

​	`Interactable` 检查元素是否可与鼠标、手指或触控笔等光标进行交互。如果元素不可交互，`err` 将返回 `ErrNotInteractable`，例如当元素被弹窗遮挡时。

#### (*Element) KeyActions <- 0.107.0

``` go
func (el *Element) KeyActions() (*KeyActions, error)
```

KeyActions is similar with Page.KeyActions. Before the action, it will try to scroll to the element and focus on it.

​	`KeyActions` 类似于 `Page.KeyActions`。在操作之前，它会尝试滚动到元素并将焦点设置到该元素上。

#### (*Element) Matches <- 0.45.0

``` go
func (el *Element) Matches(selector string) (bool, error)
```

Matches checks if the element can be selected by the css selector.

​	`Matches` 检查元素是否可通过指定的 CSS 选择器选中。

#### (*Element) MoveMouseOut <- 0.97.13

``` go
func (el *Element) MoveMouseOut() error
```

MoveMouseOut of the current element.

​	`MoveMouseOut` 将鼠标移出当前元素。

#### (*Element) MustAttribute <- 0.50.0

``` go
func (el *Element) MustAttribute(name string) *string
```

MustAttribute is similar to [Element.Attribute](https://pkg.go.dev/github.com/go-rod/rod#Element.Attribute).

​	`MustAttribute` 是 [Element.Attribute](https://pkg.go.dev/github.com/go-rod/rod#Element.Attribute) 的简化版本。

#### (*Element) MustBackgroundImage <- 0.76.6

``` go
func (el *Element) MustBackgroundImage() []byte
```

MustBackgroundImage is similar to [Element.BackgroundImage](https://pkg.go.dev/github.com/go-rod/rod#Element.BackgroundImage).

​	`MustBackgroundImage` 是 [Element.BackgroundImage](https://pkg.go.dev/github.com/go-rod/rod#Element.BackgroundImage) 的简化版本。

#### (*Element) MustBlur <- 0.50.0

``` go
func (el *Element) MustBlur() *Element
```

MustBlur is similar to [Element.Blur](https://pkg.go.dev/github.com/go-rod/rod#Element.Blur).

​	`MustBlur` 是 [Element.Blur](https://pkg.go.dev/github.com/go-rod/rod#Element.Blur) 的简化版本。

#### (*Element) MustCanvasToImage <- 0.50.0

``` go
func (el *Element) MustCanvasToImage() []byte
```

MustCanvasToImage is similar to [Element.CanvasToImage](https://pkg.go.dev/github.com/go-rod/rod#Element.CanvasToImage).

​	`MustCanvasToImage` 是 [Element.CanvasToImage](https://pkg.go.dev/github.com/go-rod/rod#Element.CanvasToImage) 的简化版本。

#### (*Element) MustClick <- 0.50.0

``` go
func (el *Element) MustClick() *Element
```

MustClick is similar to [Element.Click](https://pkg.go.dev/github.com/go-rod/rod#Element.Click).

​	`MustClick` 是 [Element.Click](https://pkg.go.dev/github.com/go-rod/rod#Element.Click) 的简化版本。

#### (*Element) MustContainsElement <- 0.50.0

``` go
func (el *Element) MustContainsElement(target *Element) bool
```

MustContainsElement is similar to [Element.ContainsElement](https://pkg.go.dev/github.com/go-rod/rod#Element.ContainsElement).

​	`MustContainsElement` 是 [Element.ContainsElement](https://pkg.go.dev/github.com/go-rod/rod#Element.ContainsElement) 的简化版本。

#### (*Element) MustDescribe <- 0.50.0

``` go
func (el *Element) MustDescribe() *proto.DOMNode
```

MustDescribe is similar to [Element.Describe](https://pkg.go.dev/github.com/go-rod/rod#Element.Describe).

​	`MustDescribe` 是 [Element.Describe](https://pkg.go.dev/github.com/go-rod/rod#Element.Describe) 的简化版本。

#### (*Element) MustDisabled <- 0.112.5

``` go
func (el *Element) MustDisabled() bool
```

MustDisabled is similar to [Element.Disabled](https://pkg.go.dev/github.com/go-rod/rod#Element.Disabled).

​	`MustDisabled` 是 [Element.Disabled](https://pkg.go.dev/github.com/go-rod/rod#Element.Disabled) 的简化版本。

#### (*Element) MustDoubleClick <- 0.111.0

``` go
func (el *Element) MustDoubleClick() *Element
```

MustDoubleClick is similar to [Element.Click](https://pkg.go.dev/github.com/go-rod/rod#Element.Click).

​	`MustDoubleClick` 是 [Element.Click](https://pkg.go.dev/github.com/go-rod/rod#Element.Click) 的简化版本，用于双击操作。

#### (*Element) MustElement <- 0.50.0

``` go
func (el *Element) MustElement(selector string) *Element
```

MustElement is similar to [Element.Element](https://pkg.go.dev/github.com/go-rod/rod#Element.Element).

​	`MustElement` 是 [Element.Element](https://pkg.go.dev/github.com/go-rod/rod#Element.Element) 的简化版本。

#### (*Element) MustElementByJS <- 0.50.0

``` go
func (el *Element) MustElementByJS(js string, params ...interface{}) *Element
```

MustElementByJS is similar to [Element.ElementByJS](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementByJS).

​	`MustElementByJS` 是 [Element.ElementByJS](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementByJS) 的简化版本。

#### (*Element) MustElementR <- 0.57.0

``` go
func (el *Element) MustElementR(selector, regex string) *Element
```

MustElementR is similar to [Element.ElementR](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementR).	

​	`MustElementR` 是 [Element.ElementR](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementR) 的简化版本。

#### (*Element) MustElementX <- 0.50.0

``` go
func (el *Element) MustElementX(xpath string) *Element
```

MustElementX is similar to [Element.ElementX](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementX).

​	`MustElementX` 是 [Element.ElementX](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementX) 的简化版本。

#### (*Element) MustElements <- 0.50.0

``` go
func (el *Element) MustElements(selector string) Elements
```

MustElements is similar to [Element.Elements](https://pkg.go.dev/github.com/go-rod/rod#Element.Elements).

​	`MustElements` 是 [Element.Elements](https://pkg.go.dev/github.com/go-rod/rod#Element.Elements) 的简化版本。

#### (*Element) MustElementsByJS <- 0.50.0

``` go
func (el *Element) MustElementsByJS(js string, params ...interface{}) Elements
```

MustElementsByJS is similar to [Element.ElementsByJS](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementsByJS).

​	`MustElementsByJS` 是 [Element.ElementsByJS](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementsByJS) 的简化版本。

#### (*Element) MustElementsX <- 0.50.0

``` go
func (el *Element) MustElementsX(xpath string) Elements
```

MustElementsX is similar to [Element.ElementsX](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementsX).

​	`MustElementsX` 是 [Element.ElementsX](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementsX) 的简化版本。

#### (*Element) MustEqual <- 0.85.7

``` go
func (el *Element) MustEqual(elm *Element) bool
```

MustEqual is similar to [Element.Equal](https://pkg.go.dev/github.com/go-rod/rod#Element.Equal).

​	`MustEqual` 是 [Element.Equal](https://pkg.go.dev/github.com/go-rod/rod#Element.Equal) 的简化版本。

#### (*Element) MustEval <- 0.50.0

``` go
func (el *Element) MustEval(js string, params ...interface{}) gson.JSON
```

MustEval is similar to [Element.Eval](https://pkg.go.dev/github.com/go-rod/rod#Element.Eval).

​	`MustEval` 是 [Element.Eval](https://pkg.go.dev/github.com/go-rod/rod#Element.Eval) 的简化版本。

#### (*Element) MustFocus <- 0.50.0

``` go
func (el *Element) MustFocus() *Element
```

MustFocus is similar to [Element.Focus](https://pkg.go.dev/github.com/go-rod/rod#Element.Focus).

​	`MustFocus` 是 [Element.Focus](https://pkg.go.dev/github.com/go-rod/rod#Element.Focus) 的简化版本。

#### (*Element) MustFrame <- 0.55.1

``` go
func (el *Element) MustFrame() *Page
```

MustFrame is similar to [Element.Frame](https://pkg.go.dev/github.com/go-rod/rod#Element.Frame).

​	`MustFrame` 是 [Element.Frame](https://pkg.go.dev/github.com/go-rod/rod#Element.Frame) 的简化版本。

#### (*Element) MustGetXPath <- 0.109.3

``` go
func (el *Element) MustGetXPath(optimized bool) string
```

MustGetXPath is similar to [Element.GetXPath](https://pkg.go.dev/github.com/go-rod/rod#Element.GetXPath).

​	`MustGetXPath` 是 [Element.GetXPath](https://pkg.go.dev/github.com/go-rod/rod#Element.GetXPath) 的简化版本。

#### (*Element) MustHTML <- 0.50.0

``` go
func (el *Element) MustHTML() string
```

MustHTML is similar to [Element.HTML](https://pkg.go.dev/github.com/go-rod/rod#Element.HTML).

​	`MustHTML` 是 [Element.HTML](https://pkg.go.dev/github.com/go-rod/rod#Element.HTML) 的简化版本。

#### (*Element) MustHas <- 0.50.0

``` go
func (el *Element) MustHas(selector string) bool
```

MustHas is similar to [Element.Has](https://pkg.go.dev/github.com/go-rod/rod#Element.Has).

​	`MustHas` 是 [Element.Has](https://pkg.go.dev/github.com/go-rod/rod#Element.Has) 的简化版本。

#### (*Element) MustHasR <- 0.61.0

``` go
func (el *Element) MustHasR(selector, regex string) bool
```

MustHasR is similar to [Element.HasR](https://pkg.go.dev/github.com/go-rod/rod#Element.HasR).

​	`MustHasR` 是 [Element.HasR](https://pkg.go.dev/github.com/go-rod/rod#Element.HasR) 的简化版本。

#### (*Element) MustHasX <- 0.50.0

``` go
func (el *Element) MustHasX(selector string) bool
```

MustHasX is similar to [Element.HasX](https://pkg.go.dev/github.com/go-rod/rod#Element.HasX).

​	`MustHasX` 是 [Element.HasX](https://pkg.go.dev/github.com/go-rod/rod#Element.HasX) 的简化版本。

#### (*Element) MustHover <- 0.50.0

``` go
func (el *Element) MustHover() *Element
```

MustHover is similar to [Element.Hover](https://pkg.go.dev/github.com/go-rod/rod#Element.Hover).

​	`MustHover` 是 [Element.Hover](https://pkg.go.dev/github.com/go-rod/rod#Element.Hover) 的简化版本。

#### (*Element) MustInput <- 0.50.0

``` go
func (el *Element) MustInput(text string) *Element
```

MustInput is similar to [Element.Input](https://pkg.go.dev/github.com/go-rod/rod#Element.Input).

​	`MustInput` 是 [Element.Input](https://pkg.go.dev/github.com/go-rod/rod#Element.Input) 的简化版本。

#### (*Element) MustInputColor <- 0.114.3

``` go
func (el *Element) MustInputColor(color string) *Element
```

MustInputColor is similar to [Element.InputColor](https://pkg.go.dev/github.com/go-rod/rod#Element.InputColor).

​	`MustInputColor` 是 [Element.InputColor](https://pkg.go.dev/github.com/go-rod/rod#Element.InputColor) 的简化版本。

#### (*Element) MustInputTime <- 0.79.2

``` go
func (el *Element) MustInputTime(t time.Time) *Element
```

MustInputTime is similar to [Element.Input](https://pkg.go.dev/github.com/go-rod/rod#Element.Input).

​	`MustInputTime` 是 [Element.InputTime](https://pkg.go.dev/github.com/go-rod/rod#Element.InputTime) 的简化版本。

#### (*Element) MustInteractable <- 0.66.0

``` go
func (el *Element) MustInteractable() bool
```

MustInteractable is similar to [Element.Interactable](https://pkg.go.dev/github.com/go-rod/rod#Element.Interactable).

​	`MustInteractable` 是 [Element.Interactable](https://pkg.go.dev/github.com/go-rod/rod#Element.Interactable) 的简化版本。

#### (*Element) MustKeyActions <- 0.107.0

``` go
func (el *Element) MustKeyActions() *KeyActions
```

MustKeyActions is similar to [Element.KeyActions](https://pkg.go.dev/github.com/go-rod/rod#Element.KeyActions).

​	`MustKeyActions ` 是 [Element.KeyActions](https://pkg.go.dev/github.com/go-rod/rod#Element.KeyActions) 的简化版本。

#### (*Element) MustMatches <- 0.50.0

``` go
func (el *Element) MustMatches(selector string) bool
```

MustMatches is similar to [Element.Matches](https://pkg.go.dev/github.com/go-rod/rod#Element.Matches).

​	`MustMatches` 是 [Element.Matches](https://pkg.go.dev/github.com/go-rod/rod#Element.Matches) 的简化版本。

#### (*Element) MustMoveMouseOut <- 0.97.13

``` go
func (el *Element) MustMoveMouseOut() *Element
```

MustMoveMouseOut is similar to [Element.MoveMouseOut](https://pkg.go.dev/github.com/go-rod/rod#Element.MoveMouseOut).

​	`MustMoveMouseOut` 是 [Element.MoveMouseOut](https://pkg.go.dev/github.com/go-rod/rod#Element.MoveMouseOut) 的简化版本。

#### (*Element) MustNext <- 0.50.0

``` go
func (el *Element) MustNext() *Element
```

MustNext is similar to [Element.Next](https://pkg.go.dev/github.com/go-rod/rod#Element.Next).

​	`MustNext` 是 [Element.Next](https://pkg.go.dev/github.com/go-rod/rod#Element.Next) 的简化版本。

#### (*Element) MustParent <- 0.50.0

``` go
func (el *Element) MustParent() *Element
```

MustParent is similar to [Element.Parent](https://pkg.go.dev/github.com/go-rod/rod#Element.Parent).

​	`MustParent` 是 [Element.Parent](https://pkg.go.dev/github.com/go-rod/rod#Element.Parent) 的简化版本。

#### (*Element) MustParents <- 0.50.0

``` go
func (el *Element) MustParents(selector string) Elements
```

MustParents is similar to [Element.Parents](https://pkg.go.dev/github.com/go-rod/rod#Element.Parents).

​	`MustParents` 是 [Element.Parents](https://pkg.go.dev/github.com/go-rod/rod#Element.Parents) 的简化版本。

#### (*Element) MustPrevious <- 0.50.0

``` go
func (el *Element) MustPrevious() *Element
```

MustPrevious is similar to [Element.Previous](https://pkg.go.dev/github.com/go-rod/rod#Element.Previous).

​	`MustPrevious` 是 [Element.Previous](https://pkg.go.dev/github.com/go-rod/rod#Element.Previous) 的简化版本。

#### (*Element) MustProperty <- 0.50.0

``` go
func (el *Element) MustProperty(name string) gson.JSON
```

MustProperty is similar to [Element.Property](https://pkg.go.dev/github.com/go-rod/rod#Element.Property).

​	`MustProperty` 是 [Element.Property](https://pkg.go.dev/github.com/go-rod/rod#Element.Property) 的简化版本。

#### (*Element) MustRelease <- 0.50.0

``` go
func (el *Element) MustRelease()
```

MustRelease is similar to [Element.Release](https://pkg.go.dev/github.com/go-rod/rod#Element.Release).

​	`MustRelease` 是 [Element.Release](https://pkg.go.dev/github.com/go-rod/rod#Element.Release) 的简化版本。

#### (*Element) MustRemove <- 0.66.0

``` go
func (el *Element) MustRemove()
```

MustRemove is similar to [Element.Remove](https://pkg.go.dev/github.com/go-rod/rod#Element.Remove).

​	`MustRemove` 是 [Element.Remove](https://pkg.go.dev/github.com/go-rod/rod#Element.Remove) 的简化版本。

#### (*Element) MustResource <- 0.50.0

``` go
func (el *Element) MustResource() []byte
```

MustResource is similar to [Element.Resource](https://pkg.go.dev/github.com/go-rod/rod#Element.Resource).

​	`MustResource` 是 [Element.Resource](https://pkg.go.dev/github.com/go-rod/rod#Element.Resource) 的简化版本。

#### (*Element) MustScreenshot <- 0.50.0

``` go
func (el *Element) MustScreenshot(toFile ...string) []byte
```

MustScreenshot is similar to [Element.Screenshot](https://pkg.go.dev/github.com/go-rod/rod#Element.Screenshot).

​	`MustScreenshot` 是 [Element.Screenshot](https://pkg.go.dev/github.com/go-rod/rod#Element.Screenshot) 的简化版本。

#### (*Element) MustScrollIntoView <- 0.50.0

``` go
func (el *Element) MustScrollIntoView() *Element
```

MustScrollIntoView is similar to [Element.ScrollIntoView](https://pkg.go.dev/github.com/go-rod/rod#Element.ScrollIntoView).

​	`MustScrollIntoView` 是 [Element.ScrollIntoView](https://pkg.go.dev/github.com/go-rod/rod#Element.ScrollIntoView) 的简化版本。

#### (*Element) MustSelect <- 0.50.0

``` go
func (el *Element) MustSelect(selectors ...string) *Element
```

MustSelect is similar to [Element.Select](https://pkg.go.dev/github.com/go-rod/rod#Element.Select).

​	`MustSelect` 是 [Element.Select](https://pkg.go.dev/github.com/go-rod/rod#Element.Select) 的简化版本。

#### (*Element) MustSelectAllText <- 0.50.0

``` go
func (el *Element) MustSelectAllText() *Element
```

MustSelectAllText is similar to [Element.SelectAllText](https://pkg.go.dev/github.com/go-rod/rod#Element.SelectAllText).

​	`MustSelectAllText` 是 [Element.SelectAllText](https://pkg.go.dev/github.com/go-rod/rod#Element.SelectAllText) 的简化版本。

#### (*Element) MustSelectText <- 0.50.0

``` go
func (el *Element) MustSelectText(regex string) *Element
```

MustSelectText is similar to [Element.SelectText](https://pkg.go.dev/github.com/go-rod/rod#Element.SelectText).

​	`MustSelectText` 是 [Element.SelectText](https://pkg.go.dev/github.com/go-rod/rod#Element.SelectText) 的简化版本。

#### (*Element) MustSetFiles <- 0.50.0

``` go
func (el *Element) MustSetFiles(paths ...string) *Element
```

MustSetFiles is similar to [Element.SetFiles](https://pkg.go.dev/github.com/go-rod/rod#Element.SetFiles).

​	`MustSetFiles` 是 [Element.SetFiles](https://pkg.go.dev/github.com/go-rod/rod#Element.SetFiles) 的简化版本。

#### (*Element) MustShadowRoot <- 0.50.0

``` go
func (el *Element) MustShadowRoot() *Element
```

MustShadowRoot is similar to [Element.ShadowRoot](https://pkg.go.dev/github.com/go-rod/rod#Element.ShadowRoot).

​	`MustShadowRoot` 是 [Element.ShadowRoot](https://pkg.go.dev/github.com/go-rod/rod#Element.ShadowRoot) 的简化版本。

#### (*Element) MustShape <- 0.66.0

``` go
func (el *Element) MustShape() *proto.DOMGetContentQuadsResult
```

MustShape is similar to [Element.Shape](https://pkg.go.dev/github.com/go-rod/rod#Element.Shape).

​	`MustShape` 是 [Element.Shape](https://pkg.go.dev/github.com/go-rod/rod#Element.Shape) 的简化版本。

#### (*Element) MustTap <- 0.61.4

``` go
func (el *Element) MustTap() *Element
```

MustTap is similar to [Element.Tap](https://pkg.go.dev/github.com/go-rod/rod#Element.Tap).

​	`MustTap` 是 [Element.Tap](https://pkg.go.dev/github.com/go-rod/rod#Element.Tap) 的简化版本。

#### (*Element) MustText <- 0.50.0

``` go
func (el *Element) MustText() string
```

MustText is similar to [Element.Text](https://pkg.go.dev/github.com/go-rod/rod#Element.Text).

​	`MustText` 是 [Element.Text](https://pkg.go.dev/github.com/go-rod/rod#Element.Text) 的简化版本。

#### (*Element) MustType <- 0.107.0

``` go
func (el *Element) MustType(keys ...input.Key) *Element
```

MustType is similar to [Element.Type](https://pkg.go.dev/github.com/go-rod/rod#Element.Type).

​	`MustType` 是 [Element.Type](https://pkg.go.dev/github.com/go-rod/rod#Element.Type) 的简化版本。

#### (*Element) MustVisible <- 0.50.0

``` go
func (el *Element) MustVisible() bool
```

MustVisible is similar to [Element.Visible](https://pkg.go.dev/github.com/go-rod/rod#Element.Visible).

​	`MustVisible` 是 [Element.Visible](https://pkg.go.dev/github.com/go-rod/rod#Element.Visible) 的简化版本。

#### (*Element) MustWait <- 0.50.0

``` go
func (el *Element) MustWait(js string, params ...interface{}) *Element
```

MustWait is similar to [Element.Wait](https://pkg.go.dev/github.com/go-rod/rod#Element.Wait).

​	`MustWait` 是 [Element.Wait](https://pkg.go.dev/github.com/go-rod/rod#Element.Wait) 的简化版本。

#### (*Element) MustWaitEnabled <- 0.84.1

``` go
func (el *Element) MustWaitEnabled() *Element
```

MustWaitEnabled is similar to [Element.WaitEnabled](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitEnabled).

​	`MustWaitEnabled` 是 [Element.WaitEnabled](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitEnabled) 的简化版本。

#### (*Element) MustWaitInteractable <- 0.88.0

``` go
func (el *Element) MustWaitInteractable() *Element
```

MustWaitInteractable is similar to [Element.WaitInteractable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitInteractable).

​	`MustWaitInteractable` 是 [Element.WaitInteractable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitInteractable) 的简化版本。

#### (*Element) MustWaitInvisible <- 0.50.0

``` go
func (el *Element) MustWaitInvisible() *Element
```

MustWaitInvisible is similar to [Element.WaitInvisible](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitInvisible)..

​	`MustWaitInvisible` 是 [Element.WaitInvisible](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitInvisible) 的简化版本。

#### (*Element) MustWaitLoad <- 0.50.0

``` go
func (el *Element) MustWaitLoad() *Element
```

MustWaitLoad is similar to [Element.WaitLoad](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitLoad).

​	`MustWaitLoad` 是 [Element.WaitLoad](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitLoad) 的简化版本。

#### (*Element) MustWaitStable <- 0.50.0

``` go
func (el *Element) MustWaitStable() *Element
```

MustWaitStable is similar to [Element.WaitStable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitStable).

​	`MustWaitStable` 是 [Element.WaitStable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitStable) 的简化版本。

#### (*Element) MustWaitVisible <- 0.50.0

``` go
func (el *Element) MustWaitVisible() *Element
```

MustWaitVisible is similar to [Element.WaitVisible](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitVisible).

​	`MustWaitVisible` 是 [Element.WaitVisible](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitVisible) 的简化版本。

#### (*Element) MustWaitWritable <- 0.84.1

``` go
func (el *Element) MustWaitWritable() *Element
```

MustWaitWritable is similar to [Element.WaitWritable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitWritable).

​	`MustWaitWritable` 是 [Element.WaitWritable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitWritable) 的简化版本。

#### (*Element) Next 

``` go
func (el *Element) Next() (*Element, error)
```

Next returns the next sibling element in the DOM tree.

​	`Next` 返回当前元素在 DOM 树中的下一个兄弟元素。

#### (*Element) Overlay <- 0.88.0

``` go
func (el *Element) Overlay(msg string) (removeOverlay func())
```

Overlay msg on the element.

​	`Overlay` 在元素上显示消息。

#### (*Element) Page <- 0.101.7

``` go
func (el *Element) Page() *Page
```

Page of the element.

​	`Page` 返回该元素所属的页面。

#### (*Element) Parent 

``` go
func (el *Element) Parent() (*Element, error)
```

Parent returns the parent element in the DOM tree.

​	`Parent` 返回 DOM 树中的父元素。

#### (*Element) Parents 

``` go
func (el *Element) Parents(selector string) (Elements, error)
```

Parents that match the selector.

​	`Parents` 返回匹配选择器的所有父元素。

#### (*Element) Previous 

``` go
func (el *Element) Previous() (*Element, error)
```

Previous returns the previous sibling element in the DOM tree.

​	`Previous` 返回 DOM 树中的上一个兄弟元素。

#### (*Element) Property 

``` go
func (el *Element) Property(name string) (gson.JSON, error)
```

Property of the DOM object. Property vs Attribute: https://stackoverflow.com/questions/6003819/what-is-the-difference-between-properties-and-attributes-in-html

​	`Property` 返回 DOM 对象的属性。有关属性与属性值的区别请参考：[Attribute vs Property](https://stackoverflow.com/questions/6003819/what-is-the-difference-between-properties-and-attributes-in-html)。

#### (*Element) Release 

``` go
func (el *Element) Release() error
```

Release is a shortcut for [Page.Release](https://pkg.go.dev/github.com/go-rod/rod#Page.Release) current element.

​	`Release` 是 [Page.Release](https://pkg.go.dev/github.com/go-rod/rod#Page.Release) 的快捷方式，用于释放当前元素。

#### (*Element) Remove <- 0.66.0

``` go
func (el *Element) Remove() error
```

Remove the element from the page.

​	`Remove` 从页面中移除该元素。

#### (*Element) Resource 

``` go
func (el *Element) Resource() ([]byte, error)
```

Resource returns the "src" content of current element. Such as the jpg of `<img src="a.jpg">`.

​	`Resource` 返回当前元素的 "src" 内容，例如 `<img src="a.jpg">` 中的图片内容。

#### (*Element) Screenshot 

``` go
func (el *Element) Screenshot(format proto.PageCaptureScreenshotFormat, quality int) ([]byte, error)
```

Screenshot of the area of the element.

​	`Screenshot` 捕获当前元素区域的截图。

#### (*Element) ScrollIntoView 

``` go
func (el *Element) ScrollIntoView() error
```

ScrollIntoView scrolls the current element into the visible area of the browser window if it's not already within the visible area.

​	`ScrollIntoView` 将当前元素滚动到浏览器窗口的可视区域。

#### (*Element) Select 

``` go
func (el *Element) Select(selectors []string, selected bool, t SelectorType) error
```

Select the children option elements that match the selectors. Before the action, it will scroll to the element, wait until it's visible. If no option matches the selectors, it will return [ErrElementNotFound].

​	`Select` 选择匹配选择器的子选项元素。在操作之前，它会滚动到该元素并等待其可见。如果没有选项匹配选择器，则返回 [ErrElementNotFound](https://pkg.go.dev/github.com/go-rod/rod#ErrElementNotFound)。

#### (*Element) SelectAllText 

``` go
func (el *Element) SelectAllText() error
```

SelectAllText selects all text Before the action, it will try to scroll to the element and focus on it.

​	`SelectAllText` 选择元素中的所有文本。在操作之前，它会尝试滚动到该元素并聚焦。

#### (*Element) SelectText 

``` go
func (el *Element) SelectText(regex string) error
```

SelectText selects the text that matches the regular expression. Before the action, it will try to scroll to the element and focus on it.

​	`SelectText` 选择匹配正则表达式的文本。在操作之前，它会尝试滚动到该元素并聚焦。

#### (*Element) SetFiles 

``` go
func (el *Element) SetFiles(paths []string) error
```

SetFiles of the current file input element.

​	`SetFiles` 设置当前文件输入元素的文件路径。

#### (*Element) ShadowRoot 

``` go
func (el *Element) ShadowRoot() (*Element, error)
```

ShadowRoot returns the shadow root of this element.

​	`ShadowRoot` 返回该元素的 shadow root。

#### (*Element) Shape <- 0.66.0

``` go
func (el *Element) Shape() (*proto.DOMGetContentQuadsResult, error)
```

Shape of the DOM element content. The shape is a group of 4-sides polygons. A 4-sides polygon is not necessary a rectangle. 4-sides polygons can be apart from each other. For example, we use 2 4-sides polygons to describe the shape below:

​	`Shape` 获取 DOM 元素内容的形状。形状由多个四边形组成。四边形不一定是矩形，可能相互分离。例如，下面的形状由两个四边形描述：

​	

```
  ____________          ____________
 /        ___/    =    /___________/    +     _________
/________/                                   /________/
```

#### (*Element) Sleeper <- 0.50.0

``` go
func (el *Element) Sleeper(sleeper func() utils.Sleeper) *Element
```

Sleeper returns a clone with the specified sleeper for chained sub-operations.

​	`Sleeper` 返回一个克隆体，并为后续链式操作指定 `sleeper`。

#### (*Element) String <- 0.88.0

``` go
func (el *Element) String() string
```

String interface.

​	`String` 接口。

#### (*Element) Tap <- 0.61.4

``` go
func (el *Element) Tap() error
```

Tap will scroll to the button and tap it just like a human. Before the action, it will try to scroll to the element and wait until it's interactable and enabled.

​	`Tap` 模拟点击按钮操作。在执行前，会尝试滚动到该元素并等待其变为可交互且启用。

#### (*Element) Text 

``` go
func (el *Element) Text() (string, error)
```

Text that the element displays.

​	`Text` 返回元素的显示文本。

#### (*Element) Timeout 

``` go
func (el *Element) Timeout(d time.Duration) *Element
```

Timeout returns a clone with the specified total timeout of all chained sub-operations.

​	`Timeout` 返回一个克隆体，并为所有链式操作设置总超时时间。

#### (*Element) Type <- 0.107.0

``` go
func (el *Element) Type(keys ...input.Key) error
```

Type is similar with Keyboard.Type. Before the action, it will try to scroll to the element and focus on it.

​	`Type` 是 [Keyboard.Type](https://pkg.go.dev/github.com/go-rod/rod#Keyboard.Type) 的简化版本。在执行前，会尝试滚动到该元素并聚焦。

#### (*Element) Visible 

``` go
func (el *Element) Visible() (bool, error)
```

Visible returns true if the element is visible on the page.

​	`Visible` 检查元素是否在页面上可见。如果可见，返回 `true`。

#### (*Element) Wait 

``` go
func (el *Element) Wait(opts *EvalOptions) error
```

Wait until the js returns true.

​	`Wait` 等待 JavaScript 表达式返回 `true`。

#### (*Element) WaitEnabled <- 0.84.1

``` go
func (el *Element) WaitEnabled() error
```

WaitEnabled until the element is not disabled. Doc for readonly: https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/readonly

​	`WaitEnabled` 等待元素不再禁用。有关 `readonly` 的文档：https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/readonly。

#### (*Element) WaitInteractable <- 0.88.0

``` go
func (el *Element) WaitInteractable() (pt *proto.Point, err error)
```

WaitInteractable waits for the element to be interactable. It will try to scroll to the element on each try.

​	`WaitInteractable` 等待元素变为可交互状态。在每次尝试时，它会滚动到该元素。

#### (*Element) WaitInvisible 

``` go
func (el *Element) WaitInvisible() error
```

WaitInvisible until the element invisible.

​	`WaitInvisible` 等待元素变为不可见状态。

#### (*Element) WaitLoad <- 0.49.0

``` go
func (el *Element) WaitLoad() error
```

WaitLoad for element like ` <img>`.

​	`WaitLoad` 等待类似 `<img>` 的元素加载完成。

#### (*Element) WaitStable 

``` go
func (el *Element) WaitStable(d time.Duration) error
```

WaitStable waits until no shape or position change for d duration. Be careful, d is not the max wait timeout, it's the least stable time. If you want to set a timeout you can use the [Element.Timeout](https://pkg.go.dev/github.com/go-rod/rod#Element.Timeout) function.

​	`WaitStable` 等待元素在 `d` 时间内没有形状或位置的变化。注意，`d` 不是最大等待超时，而是最短稳定时间。如需设置超时，可以使用 [Element.Timeout](https://pkg.go.dev/github.com/go-rod/rod#Element.Timeout)。

#### (*Element) WaitStableRAF <- 0.84.1

``` go
func (el *Element) WaitStableRAF() error
```

WaitStableRAF waits until no shape or position change for 2 consecutive animation frames. If you want to wait animation that is triggered by JS not CSS, you'd better use [Element.WaitStable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitStable). About animation frame: https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame

​	`WaitStableRAF` 等待连续两个动画帧内没有形状或位置的变化。如果要等待由 JavaScript 而非 CSS 触发的动画，建议使用 [Element.WaitStable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitStable)。关于动画帧的文档：https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame。

#### (*Element) WaitVisible 

``` go
func (el *Element) WaitVisible() error
```

WaitVisible until the element is visible.

​	`WaitVisible` 等待元素变为可见状态。

#### (*Element) WaitWritable <- 0.84.1

``` go
func (el *Element) WaitWritable() error
```

WaitWritable until the element is not readonly. Doc for disabled: https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled

​	`WaitWritable` 等待元素变为可写（非只读）状态。有关 `disabled` 的文档：https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled。

#### (*Element) WithCancel <- 0.69.0

``` go
func (el *Element) WithCancel() (*Element, func())
```

WithCancel returns a clone with a context cancel function.

​	`WithCancel` 返回一个克隆体，并提供用于取消上下文的函数。

#### (*Element) WithPanic <- 0.100.0

``` go
func (el *Element) WithPanic(fail func(interface{})) *Element
```

WithPanic returns an element clone with the specified panic function. The fail must stop the current goroutine's execution immediately, such as use [runtime.Goexit](https://pkg.go.dev/runtime#Goexit) or panic inside it.

​	`WithPanic` 返回一个包含指定 `panic` 函数的元素克隆体。`fail` 必须立即停止当前 goroutine 的执行，例如使用 [runtime.Goexit](https://pkg.go.dev/runtime#Goexit) 或触发 `panic`。

### type ElementNotFoundError <- 0.114.8

``` go
type ElementNotFoundError struct{}
```

ElementNotFoundError error.

​	`ElementNotFoundError` 错误。

#### (*ElementNotFoundError) Error <- 0.114.8

``` go
func (e *ElementNotFoundError) Error() string
```

### type Elements 

``` go
type Elements []*Element
```

Elements provides some helpers to deal with element list.

​	`Elements` 提供了一些处理元素列表的辅助功能。

#### (Elements) Empty 

``` go
func (els Elements) Empty() bool
```

Empty returns true if the list is empty.

​	`Empty` 如果列表为空，返回 `true`。

#### (Elements) First 

``` go
func (els Elements) First() *Element
```

First returns the first element, if the list is empty returns nil.

​	`First` 返回列表中的第一个元素。如果列表为空，返回 `nil`。

#### (Elements) Last 

``` go
func (els Elements) Last() *Element
```

Last returns the last element, if the list is empty returns nil.

### type EvalError <- 0.114.8

``` go
type EvalError struct {
	*proto.RuntimeExceptionDetails
}
```

EvalError error.

#### (*EvalError) Error <- 0.114.8

``` go
func (e *EvalError) Error() string
```

#### (*EvalError) Is <- 0.114.8

``` go
func (e *EvalError) Is(err error) bool
```

Is interface.

### type EvalOptions <- 0.50.0

``` go
type EvalOptions struct {
	// If enabled the eval result will be a plain JSON value.
	// If disabled the eval result will be a reference of a remote js object.
	ByValue bool

	AwaitPromise bool

	// ThisObj represents the "this" object in the JS
	ThisObj *proto.RuntimeRemoteObject

	// JS function definition to execute.
	JS string

	// JSArgs represents the arguments that will be passed to JS.
	// If an argument is [*proto.RuntimeRemoteObject] type, the corresponding remote object will be used.
	// Or it will be passed as a plain JSON value.
	// When an arg in the args is a *js.Function, the arg will be cached on the page's js context.
	// When the arg.Name exists in the page's cache, it reuse the cache without sending
	// the definition to the browser again.
	// Useful when you need to eval a huge js expression many times.
	JSArgs []interface{}

	// Whether execution should be treated as initiated by user in the UI.
	UserGesture bool
}
```

EvalOptions for Page.Evaluate.

### func Eval <- 0.67.0

``` go
func Eval(js string, args ...interface{}) *EvalOptions
```

Eval creates a [EvalOptions](https://pkg.go.dev/github.com/go-rod/rod#EvalOptions) with ByValue set to true.

#### (*EvalOptions) ByObject <- 0.50.0

``` go
func (e *EvalOptions) ByObject() *EvalOptions
```

ByObject disables ByValue.

#### (*EvalOptions) ByPromise <- 0.74.0

``` go
func (e *EvalOptions) ByPromise() *EvalOptions
```

ByPromise enables AwaitPromise.

#### (*EvalOptions) ByUser <- 0.64.0

``` go
func (e *EvalOptions) ByUser() *EvalOptions
```

ByUser enables UserGesture.

#### (*EvalOptions) String <- 0.88.0

``` go
func (e *EvalOptions) String() string
```

String interface.

#### (*EvalOptions) This <- 0.50.0

``` go
func (e *EvalOptions) This(obj *proto.RuntimeRemoteObject) *EvalOptions
```

This set the obj as ThisObj.

### type ExpectElementError <- 0.114.8

``` go
type ExpectElementError struct {
	*proto.RuntimeRemoteObject
}
```

ExpectElementError error.

#### (*ExpectElementError) Error <- 0.114.8

``` go
func (e *ExpectElementError) Error() string
```

#### (*ExpectElementError) Is <- 0.114.8

``` go
func (e *ExpectElementError) Is(err error) bool
```

Is interface.

### type ExpectElementsError <- 0.114.8

``` go
type ExpectElementsError struct {
	*proto.RuntimeRemoteObject
}
```

ExpectElementsError error.

#### (*ExpectElementsError) Error <- 0.114.8

``` go
func (e *ExpectElementsError) Error() string
```

#### (*ExpectElementsError) Is <- 0.114.8

``` go
func (e *ExpectElementsError) Is(err error) bool
```

Is interface.

### type Hijack 

``` go
type Hijack struct {
	Request  *HijackRequest
	Response *HijackResponse
	OnError  func(error)

	// Skip to next handler
	Skip bool

	// CustomState is used to store things for this context
	CustomState interface{}
	// contains filtered or unexported fields
}
```

Hijack context.

#### (*Hijack) ContinueRequest <- 0.42.0

``` go
func (h *Hijack) ContinueRequest(cq *proto.FetchContinueRequest)
```

ContinueRequest without hijacking. The RequestID will be set by the router, you don't have to set it.

#### (*Hijack) LoadResponse 

``` go
func (h *Hijack) LoadResponse(client *http.Client, loadBody bool) error
```

LoadResponse will send request to the real destination and load the response as default response to override.

#### (*Hijack) MustLoadResponse <- 0.50.0

``` go
func (h *Hijack) MustLoadResponse()
```

MustLoadResponse is similar to [Hijack.LoadResponse](https://pkg.go.dev/github.com/go-rod/rod#Hijack.LoadResponse).

### type HijackRequest 

``` go
type HijackRequest struct {
	// contains filtered or unexported fields
}
```

HijackRequest context.

#### (*HijackRequest) Body 

``` go
func (ctx *HijackRequest) Body() string
```

Body of the request, devtools API doesn't support binary data yet, only string can be captured.

#### (*HijackRequest) Header 

``` go
func (ctx *HijackRequest) Header(key string) string
```

Header via a key.

#### (*HijackRequest) Headers 

``` go
func (ctx *HijackRequest) Headers() proto.NetworkHeaders
```

Headers of request.

#### (*HijackRequest) IsNavigation <- 0.97.1

``` go
func (ctx *HijackRequest) IsNavigation() bool
```

IsNavigation determines whether the request is a navigation request.

#### (*HijackRequest) JSONBody 

``` go
func (ctx *HijackRequest) JSONBody() gson.JSON
```

JSONBody of the request.

#### (*HijackRequest) Method 

``` go
func (ctx *HijackRequest) Method() string
```

Method of the request.

#### (*HijackRequest) Req <- 0.52.0

``` go
func (ctx *HijackRequest) Req() *http.Request
```

Req returns the underlying http.Request instance that will be used to send the request.

#### (*HijackRequest) SetBody 

``` go
func (ctx *HijackRequest) SetBody(obj interface{}) *HijackRequest
```

SetBody of the request, if obj is []byte or string, raw body will be used, else it will be encoded as json.

#### (*HijackRequest) SetContext <- 0.57.1

``` go
func (ctx *HijackRequest) SetContext(c context.Context) *HijackRequest
```

SetContext of the underlying http.Request instance.

#### (*HijackRequest) Type <- 0.49.1

``` go
func (ctx *HijackRequest) Type() proto.NetworkResourceType
```

Type of the resource.

#### (*HijackRequest) URL 

``` go
func (ctx *HijackRequest) URL() *url.URL
```

URL of the request.

### type HijackResponse 

``` go
type HijackResponse struct {
	RawResponse *http.Response
	// contains filtered or unexported fields
}
```

HijackResponse context.

#### (*HijackResponse) Body 

``` go
func (ctx *HijackResponse) Body() string
```

Body of the payload.

#### (*HijackResponse) Fail <- 0.48.1

``` go
func (ctx *HijackResponse) Fail(reason proto.NetworkErrorReason) *HijackResponse
```

Fail request.

#### (*HijackResponse) Headers 

``` go
func (ctx *HijackResponse) Headers() http.Header
```

Headers returns the clone of response headers. If you want to modify the response headers use HijackResponse.SetHeader .

#### (*HijackResponse) Payload <- 0.52.0

``` go
func (ctx *HijackResponse) Payload() *proto.FetchFulfillRequest
```

Payload to respond the request from the browser.

#### (*HijackResponse) SetBody 

``` go
func (ctx *HijackResponse) SetBody(obj interface{}) *HijackResponse
```

SetBody of the payload, if obj is []byte or string, raw body will be used, else it will be encoded as json.

#### (*HijackResponse) SetHeader 

``` go
func (ctx *HijackResponse) SetHeader(pairs ...string) *HijackResponse
```

SetHeader of the payload via key-value pairs.

### type HijackRouter 

``` go
type HijackRouter struct {
	// contains filtered or unexported fields
}
```

HijackRouter context.

#### (*HijackRouter) Add 

``` go
func (r *HijackRouter) Add(pattern string, resourceType proto.NetworkResourceType, handler func(*Hijack)) error
```

Add a hijack handler to router, the doc of the pattern is the same as "proto.FetchRequestPattern.URLPattern".

#### (*HijackRouter) MustAdd <- 0.50.0

``` go
func (r *HijackRouter) MustAdd(pattern string, handler func(*Hijack)) *HijackRouter
```

MustAdd is similar to [HijackRouter.Add](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.Add).

#### (*HijackRouter) MustRemove <- 0.50.0

``` go
func (r *HijackRouter) MustRemove(pattern string) *HijackRouter
```

MustRemove is similar to [HijackRouter.Remove](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.Remove).

#### (*HijackRouter) MustStop <- 0.50.0

``` go
func (r *HijackRouter) MustStop()
```

MustStop is similar to [HijackRouter.Stop](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.Stop).

#### (*HijackRouter) Remove 

``` go
func (r *HijackRouter) Remove(pattern string) error
```

Remove handler via the pattern.

#### (*HijackRouter) Run 

``` go
func (r *HijackRouter) Run()
```

Run the router, after you call it, you shouldn't add new handler to it.

#### (*HijackRouter) Stop 

``` go
func (r *HijackRouter) Stop() error
```

Stop the router.

### type InvisibleShapeError <- 0.114.8

``` go
type InvisibleShapeError struct {
	*Element
}
```

InvisibleShapeError error.

#### (*InvisibleShapeError) Error <- 0.114.8

``` go
func (e *InvisibleShapeError) Error() string
```

Error ...

#### (*InvisibleShapeError) Is <- 0.114.8

``` go
func (e *InvisibleShapeError) Is(err error) bool
```

Is interface.

#### (*InvisibleShapeError) Unwrap <- 0.114.8

``` go
func (e *InvisibleShapeError) Unwrap() error
```

Unwrap ...

### type KeyAction <- 0.107.0

``` go
type KeyAction struct {
	Type KeyActionType
	Key  input.Key
}
```

KeyAction to perform.

### type KeyActionType <- 0.107.0

``` go
type KeyActionType int
```

KeyActionType enum.

``` go
const (
	KeyActionPress KeyActionType = iota
	KeyActionRelease
	KeyActionTypeKey
)
```

KeyActionTypes.

### type KeyActions <- 0.107.0

``` go
type KeyActions struct {
	Actions []KeyAction
	// contains filtered or unexported fields
}
```

KeyActions to simulate.

#### (*KeyActions) Do <- 0.107.0

``` go
func (ka *KeyActions) Do() (err error)
```

Do the actions.

#### (*KeyActions) MustDo <- 0.107.0

``` go
func (ka *KeyActions) MustDo()
```

MustDo is similar to [KeyActions.Do](https://pkg.go.dev/github.com/go-rod/rod#KeyActions.Do).

#### (*KeyActions) Press <- 0.107.0

``` go
func (ka *KeyActions) Press(keys ...input.Key) *KeyActions
```

Press keys is guaranteed to have a release at the end of actions.

#### (*KeyActions) Release <- 0.107.0

``` go
func (ka *KeyActions) Release(keys ...input.Key) *KeyActions
```

Release keys.

#### (*KeyActions) Type <- 0.107.0

``` go
func (ka *KeyActions) Type(keys ...input.Key) *KeyActions
```

Type will release the key immediately after the pressing.

### type Keyboard 

``` go
type Keyboard struct {
	sync.Mutex
	// contains filtered or unexported fields
}
```

Keyboard represents the keyboard on a page, it's always related the main frame.

#### (*Keyboard) MustType <- 0.107.0

``` go
func (k *Keyboard) MustType(key ...input.Key) *Keyboard
```

MustType is similar to [Keyboard.Type](https://pkg.go.dev/github.com/go-rod/rod#Keyboard.Type).

#### (*Keyboard) Press 

``` go
func (k *Keyboard) Press(key input.Key) error
```

Press the key down. To input characters that are not on the keyboard, such as Chinese or Japanese, you should use method like [Page.InsertText](https://pkg.go.dev/github.com/go-rod/rod#Page.InsertText).

#### (*Keyboard) Release <- 0.107.0

``` go
func (k *Keyboard) Release(key input.Key) error
```

Release the key.

#### (*Keyboard) Type <- 0.107.0

``` go
func (k *Keyboard) Type(keys ...input.Key) (err error)
```

Type releases the key after the press.

### type Message <- 0.74.0

``` go
type Message struct {
	SessionID proto.TargetSessionID
	Method    string
	// contains filtered or unexported fields
}
```

Message represents a cdp.Event.

#### (*Message) Load <- 0.74.0

``` go
func (msg *Message) Load(e proto.Event) bool
```

Load data into e, returns true if e matches the event type.

### type Mouse 

``` go
type Mouse struct {
	sync.Mutex
	// contains filtered or unexported fields
}
```

Mouse represents the mouse on a page, it's always related the main frame.

#### (*Mouse) Click 

``` go
func (m *Mouse) Click(button proto.InputMouseButton, clickCount int) error
```

Click the button. It's the combination of [Mouse.Down](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Down) and [Mouse.Up](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Up).

#### (*Mouse) Down 

``` go
func (m *Mouse) Down(button proto.InputMouseButton, clickCount int) error
```

Down holds the button down.

#### (*Mouse) MoveAlong <- 0.112.0

``` go
func (m *Mouse) MoveAlong(guide func() (proto.Point, bool)) error
```

MoveAlong the guide function. Every time the guide function is called it should return the next mouse position, return true to stop. Read the source code of [Mouse.MoveLinear](https://pkg.go.dev/github.com/go-rod/rod#Mouse.MoveLinear) as an example to use this method.

#### (*Mouse) MoveLinear <- 0.112.0

``` go
func (m *Mouse) MoveLinear(to proto.Point, steps int) error
```

MoveLinear to the absolute position with the given steps. Such as move from (0,0) to (6,6) with 3 steps, the mouse will first move to (2,2) then (4,4) then (6,6).

#### (*Mouse) MoveTo <- 0.112.0

``` go
func (m *Mouse) MoveTo(p proto.Point) error
```

MoveTo the absolute position.

#### (*Mouse) MustClick <- 0.50.0

``` go
func (m *Mouse) MustClick(button proto.InputMouseButton) *Mouse
```

MustClick is similar to [Mouse.Click](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Click).

#### (*Mouse) MustDown <- 0.50.0

``` go
func (m *Mouse) MustDown(button proto.InputMouseButton) *Mouse
```

MustDown is similar to [Mouse.Down](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Down).

#### (*Mouse) MustMoveTo <- 0.112.0

``` go
func (m *Mouse) MustMoveTo(x, y float64) *Mouse
```

MustMoveTo is similar to [Mouse.Move].

#### (*Mouse) MustScroll <- 0.50.0

``` go
func (m *Mouse) MustScroll(x, y float64) *Mouse
```

MustScroll is similar to [Mouse.Scroll](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Scroll).

#### (*Mouse) MustUp <- 0.50.0

``` go
func (m *Mouse) MustUp(button proto.InputMouseButton) *Mouse
```

MustUp is similar to [Mouse.Up](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Up).

#### (*Mouse) Position <- 0.112.0

``` go
func (m *Mouse) Position() proto.Point
```

Position of current cursor.

#### (*Mouse) Scroll 

``` go
func (m *Mouse) Scroll(offsetX, offsetY float64, steps int) error
```

Scroll the relative offset with specified steps.

#### (*Mouse) Up 

``` go
func (m *Mouse) Up(button proto.InputMouseButton, clickCount int) error
```

Up releases the button.

### type NavigationError <- 0.114.8

``` go
type NavigationError struct {
	Reason string
}
```

NavigationError error.

#### (*NavigationError) Error <- 0.114.8

``` go
func (e *NavigationError) Error() string
```

#### (*NavigationError) Is <- 0.114.8

``` go
func (e *NavigationError) Is(err error) bool
```

Is interface.

### type NoPointerEventsError <- 0.114.8

``` go
type NoPointerEventsError struct {
	*Element
}
```

NoPointerEventsError error.

#### (*NoPointerEventsError) Error <- 0.114.8

``` go
func (e *NoPointerEventsError) Error() string
```

Error ...

#### (*NoPointerEventsError) Is <- 0.114.8

``` go
func (e *NoPointerEventsError) Is(err error) bool
```

Is interface.

#### (*NoPointerEventsError) Unwrap <- 0.114.8

``` go
func (e *NoPointerEventsError) Unwrap() error
```

Unwrap ...

### type NoShadowRootError <- 0.114.8

``` go
type NoShadowRootError struct {
	*Element
}
```

NoShadowRootError error.

#### (*NoShadowRootError) Error <- 0.114.8

``` go
func (e *NoShadowRootError) Error() string
```

Error ...

#### (*NoShadowRootError) Is <- 0.114.8

``` go
func (e *NoShadowRootError) Is(err error) bool
```

Is interface.

### type NotInteractableError <- 0.114.8

``` go
type NotInteractableError struct{}
```

NotInteractableError error. Check the doc of Element.Interactable for details.

#### (*NotInteractableError) Error <- 0.114.8

``` go
func (e *NotInteractableError) Error() string
```

### type ObjectNotFoundError <- 0.114.8

``` go
type ObjectNotFoundError struct {
	*proto.RuntimeRemoteObject
}
```

ObjectNotFoundError error.

#### (*ObjectNotFoundError) Error <- 0.114.8

``` go
func (e *ObjectNotFoundError) Error() string
```

#### (*ObjectNotFoundError) Is <- 0.114.8

``` go
func (e *ObjectNotFoundError) Is(err error) bool
```

Is interface.

### type Page 

``` go
type Page struct {
	// TargetID is a unique ID for a remote page.
	// It's usually used in events sent from the browser to tell which page an event belongs to.
	TargetID proto.TargetTargetID

	// FrameID is a unique ID for a browsing context.
	// Usually, different FrameID means different javascript execution context.
	// Such as an iframe and the page it belongs to will have the same TargetID but different FrameIDs.
	FrameID proto.PageFrameID

	// SessionID is a unique ID for a page attachment to a controller.
	// It's usually used in transport layer to tell which page to send the control signal.
	// A page can attached to multiple controllers, the browser uses it distinguish controllers.
	SessionID proto.TargetSessionID

	// devices
	Mouse    *Mouse
	Keyboard *Keyboard
	Touch    *Touch
	// contains filtered or unexported fields
}
```

Page represents the webpage. We try to hold as less states as possible. When a page is closed by Rod or not all the ongoing operations an events on it will abort.

##### Example (Pool)

``` go
```
#### (*Page) Activate <- 0.86.3

``` go
func (p *Page) Activate() (*Page, error)
```

Activate (focuses) the page.

#### (*Page) AddScriptTag 

``` go
func (p *Page) AddScriptTag(url, content string) error
```

AddScriptTag to page. If url is empty, content will be used.

#### (*Page) AddStyleTag 

``` go
func (p *Page) AddStyleTag(url, content string) error
```

AddStyleTag to page. If url is empty, content will be used.

#### (*Page) Browser <- 0.101.7

``` go
func (p *Page) Browser() *Browser
```

Browser of the page.

#### (*Page) Call <- 0.70.0

``` go
func (p *Page) Call(ctx context.Context, sessionID, methodName string, params interface{}) (res []byte, err error)
```

Call implements the [proto.Client](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#Client).

#### (*Page) CancelTimeout 

``` go
func (p *Page) CancelTimeout() *Page
```

CancelTimeout cancels the current timeout context and returns a clone with the parent context.

#### (*Page) CaptureDOMSnapshot <- 0.113.0

``` go
func (p *Page) CaptureDOMSnapshot() (domSnapshot *proto.DOMSnapshotCaptureSnapshotResult, err error)
```

CaptureDOMSnapshot Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened. `Documents` The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document. `Strings` Shared string table that all string properties refer to with indexes. Normally use `Strings` is enough.

#### (*Page) Close 

``` go
func (p *Page) Close() error
```

Close tries to close page, running its beforeunload hooks, if has any.

#### (*Page) Context 

``` go
func (p *Page) Context(ctx context.Context) *Page
```

Context returns a clone with the specified ctx for chained sub-operations.

#### (*Page) Cookies 

``` go
func (p *Page) Cookies(urls []string) ([]*proto.NetworkCookie, error)
```

Cookies returns the page cookies. By default it will return the cookies for current page. The urls is the list of URLs for which applicable cookies will be fetched.

#### (*Page) DisableDomain 

``` go
func (p *Page) DisableDomain(method proto.Request) (restore func())
```

DisableDomain and returns a restore function to restore previous state.

#### (*Page) EachEvent 

``` go
func (p *Page) EachEvent(callbacks ...interface{}) (wait func())
```

EachEvent of the specified event types, if any callback returns true the wait function will resolve, The type of each callback is (? means optional):

``` go
func(proto.Event, proto.TargetSessionID?) bool?
```

You can listen to multiple event types at the same time like:

```
browser.EachEvent(func(a *proto.A) {}, func(b *proto.B) {})
```

Such as subscribe the events to know when the navigation is complete or when the page is rendered. Here's an example to dismiss all dialogs/alerts on the page:

```
go page.EachEvent(func(e *proto.PageJavascriptDialogOpening) {
    _ = proto.PageHandleJavaScriptDialog{ Accept: false, PromptText: ""}.Call(page)
})()
```

#### (*Page) Element 

``` go
func (p *Page) Element(selector string) (*Element, error)
```

Element retries until an element in the page that matches the CSS selector, then returns the matched element.

#### (*Page) ElementByJS 

``` go
func (p *Page) ElementByJS(opts *EvalOptions) (*Element, error)
```

ElementByJS returns the element from the return value of the js function. If sleeper is nil, no retry will be performed. By default, it will retry until the js function doesn't return null. To customize the retry logic, check the examples of Page.Sleeper.

#### (*Page) ElementFromNode <- 0.47.0

``` go
func (p *Page) ElementFromNode(node *proto.DOMNode) (*Element, error)
```

ElementFromNode creates an Element from the node, [proto.DOMNodeID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMNodeID) or [proto.DOMBackendNodeID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMBackendNodeID) must be specified.

#### (*Page) ElementFromObject <- 0.47.0

``` go
func (p *Page) ElementFromObject(obj *proto.RuntimeRemoteObject) (*Element, error)
```

ElementFromObject creates an Element from the remote object id.

#### (*Page) ElementFromPoint <- 0.48.0

``` go
func (p *Page) ElementFromPoint(x, y int) (*Element, error)
```

ElementFromPoint creates an Element from the absolute point on the page. The point should include the window scroll offset.

#### (*Page) ElementR <- 0.57.0

``` go
func (p *Page) ElementR(selector, jsRegex string) (*Element, error)
```

ElementR retries until an element in the page that matches the css selector and it's text matches the jsRegex, then returns the matched element.

#### (*Page) ElementX 

``` go
func (p *Page) ElementX(xPath string) (*Element, error)
```

ElementX retries until an element in the page that matches one of the XPath selectors, then returns the matched element.

#### (*Page) Elements 

``` go
func (p *Page) Elements(selector string) (Elements, error)
```

Elements returns all elements that match the css selector.

#### (*Page) ElementsByJS 

``` go
func (p *Page) ElementsByJS(opts *EvalOptions) (Elements, error)
```

ElementsByJS returns the elements from the return value of the js.

#### (*Page) ElementsX 

``` go
func (p *Page) ElementsX(xpath string) (Elements, error)
```

ElementsX returns all elements that match the XPath selector.

#### (*Page) Emulate <- 0.42.1

``` go
func (p *Page) Emulate(device devices.Device) error
```

Emulate the device, such as iPhone9. If device is devices.Clear, it will clear the override.

#### (*Page) EnableDomain 

``` go
func (p *Page) EnableDomain(method proto.Request) (restore func())
```

EnableDomain and returns a restore function to restore previous state.

#### (*Page) Eval 

``` go
func (p *Page) Eval(js string, args ...interface{}) (*proto.RuntimeRemoteObject, error)
```

Eval is a shortcut for [Page.Evaluate](https://pkg.go.dev/github.com/go-rod/rod#Page.Evaluate) with AwaitPromise, ByValue set to true.

#### (*Page) EvalOnNewDocument <- 0.44.0

``` go
func (p *Page) EvalOnNewDocument(js string) (remove func() error, err error)
```

EvalOnNewDocument Evaluates given script in every frame upon creation (before loading frame's scripts).

#### (*Page) Evaluate <- 0.67.0

``` go
func (p *Page) Evaluate(opts *EvalOptions) (res *proto.RuntimeRemoteObject, err error)
```

Evaluate js on the page.

#### (*Page) Event <- 0.70.2

``` go
func (p *Page) Event() <-chan *Message
```

Event of the page.

#### (*Page) Expose <- 0.49.1

``` go
func (p *Page) Expose(name string, fn func(gson.JSON) (interface{}, error)) (stop func() error, err error)
```

Expose fn to the page's window object with the name. The exposure survives reloads. Call stop to unbind the fn.

#### (*Page) ExposeHelpers <- 0.85.1

``` go
func (p *Page) ExposeHelpers(list ...*js.Function)
```

ExposeHelpers helper functions to page's js context so that we can use the Devtools' console to debug them.

#### (*Page) GetContext 

``` go
func (p *Page) GetContext() context.Context
```

GetContext of current instance.

#### (*Page) GetNavigationHistory <- 0.116.2

``` go
func (p *Page) GetNavigationHistory() (*proto.PageGetNavigationHistoryResult, error)
```

GetNavigationHistory get navigation history.

#### (*Page) GetResource <- 0.76.6

``` go
func (p *Page) GetResource(url string) ([]byte, error)
```

GetResource content by the url. Such as image, css, html, etc. Use the [proto.PageGetResourceTree](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#PageGetResourceTree) to list all the resources.

#### (*Page) GetSessionID <- 0.72.0

``` go
func (p *Page) GetSessionID() proto.TargetSessionID
```

GetSessionID interface.

#### (*Page) GetWindow 

``` go
func (p *Page) GetWindow() (*proto.BrowserBounds, error)
```

GetWindow position and size info.

#### (*Page) HTML <- 0.94.0

``` go
func (p *Page) HTML() (string, error)
```

HTML of the page.

#### (*Page) HandleDialog 

``` go
func (p *Page) HandleDialog() (
	wait func() *proto.PageJavascriptDialogOpening,
	handle func(*proto.PageHandleJavaScriptDialog) error,
)
```

HandleDialog accepts or dismisses next JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload). Because modal dialog will block js, usually you have to trigger the dialog in another goroutine. For example:

```
wait, handle := page.MustHandleDialog()
go page.MustElement("button").MustClick()
wait()
handle(true, "")
```

#### (*Page) HandleFileDialog <- 0.109.0

``` go
func (p *Page) HandleFileDialog() (func([]string) error, error)
```

HandleFileDialog return a functions that waits for the next file chooser dialog pops up and returns the element for the event.

#### (*Page) Has 

``` go
func (p *Page) Has(selector string) (bool, *Element, error)
```

Has an element that matches the css selector.

#### (*Page) HasR <- 0.61.0

``` go
func (p *Page) HasR(selector, jsRegex string) (bool, *Element, error)
```

HasR an element that matches the css selector and its display text matches the jsRegex.

#### (*Page) HasX 

``` go
func (p *Page) HasX(selector string) (bool, *Element, error)
```

HasX an element that matches the XPath selector.

#### (*Page) HijackRequests 

``` go
func (p *Page) HijackRequests() *HijackRouter
```

HijackRequests creates a new router instance for requests hijacking. When use Fetch domain outside the router should be stopped. Enabling hijacking disables page caching, but such as 304 Not Modified will still work as expected. The entire process of hijacking one request:

```
browser --req-> rod ---> server ---> rod --res-> browser
```

The --req-> and --res-> are the parts that can be modified.

#### (*Page) Info <- 0.42.1

``` go
func (p *Page) Info() (*proto.TargetTargetInfo, error)
```

Info of the page, such as the URL or title of the page.

#### (*Page) InsertText <- 0.107.0

``` go
func (p *Page) InsertText(text string) error
```

InsertText is like pasting text into the page.

#### (*Page) IsIframe 

``` go
func (p *Page) IsIframe() bool
```

IsIframe tells if it's iframe.

#### (*Page) KeyActions <- 0.107.0

``` go
func (p *Page) KeyActions() *KeyActions
```

KeyActions simulates the type actions on a physical keyboard. Useful when input shortcuts like ctrl+enter .

#### (*Page) LoadState 

``` go
func (p *Page) LoadState(method proto.Request) (has bool)
```

LoadState into the method.

#### (*Page) MustActivate <- 0.86.3

``` go
func (p *Page) MustActivate() *Page
```

MustActivate is similar to [Page.Activate](https://pkg.go.dev/github.com/go-rod/rod#Page.Activate).

#### (*Page) MustAddScriptTag <- 0.50.0

``` go
func (p *Page) MustAddScriptTag(url string) *Page
```

MustAddScriptTag is similar to [Page.AddScriptTag](https://pkg.go.dev/github.com/go-rod/rod#Page.AddScriptTag).

#### (*Page) MustAddStyleTag <- 0.50.0

``` go
func (p *Page) MustAddStyleTag(url string) *Page
```

MustAddStyleTag is similar to [Page.AddStyleTag](https://pkg.go.dev/github.com/go-rod/rod#Page.AddStyleTag).

#### (*Page) MustCaptureDOMSnapshot <- 0.113.0

``` go
func (p *Page) MustCaptureDOMSnapshot() (domSnapshot *proto.DOMSnapshotCaptureSnapshotResult)
```

MustCaptureDOMSnapshot is similar to [Page.CaptureDOMSnapshot](https://pkg.go.dev/github.com/go-rod/rod#Page.CaptureDOMSnapshot).

#### (*Page) MustClose <- 0.50.0

``` go
func (p *Page) MustClose()
```

MustClose is similar to [Page.Close](https://pkg.go.dev/github.com/go-rod/rod#Page.Close).

#### (*Page) MustCookies <- 0.50.0

``` go
func (p *Page) MustCookies(urls ...string) []*proto.NetworkCookie
```

MustCookies is similar to [Page.Cookies](https://pkg.go.dev/github.com/go-rod/rod#Page.Cookies).

#### (*Page) MustElement <- 0.50.0

``` go
func (p *Page) MustElement(selector string) *Element
```

MustElement is similar to [Page.Element](https://pkg.go.dev/github.com/go-rod/rod#Page.Element).

#### (*Page) MustElementByJS <- 0.50.0

``` go
func (p *Page) MustElementByJS(js string, params ...interface{}) *Element
```

MustElementByJS is similar to [Page.ElementByJS](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementByJS).

#### (*Page) MustElementFromNode <- 0.50.0

``` go
func (p *Page) MustElementFromNode(node *proto.DOMNode) *Element
```

MustElementFromNode is similar to [Page.ElementFromNode](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementFromNode).

#### (*Page) MustElementFromPoint <- 0.50.0

``` go
func (p *Page) MustElementFromPoint(left, top int) *Element
```

MustElementFromPoint is similar to [Page.ElementFromPoint](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementFromPoint).

#### (*Page) MustElementR <- 0.57.0

``` go
func (p *Page) MustElementR(selector, jsRegex string) *Element
```

MustElementR is similar to [Page.ElementR](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementR).

#### (*Page) MustElementX <- 0.50.0

``` go
func (p *Page) MustElementX(xPath string) *Element
```

MustElementX is similar to [Page.ElementX](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementX).

#### (*Page) MustElements <- 0.50.0

``` go
func (p *Page) MustElements(selector string) Elements
```

MustElements is similar to [Page.Elements](https://pkg.go.dev/github.com/go-rod/rod#Page.Elements).

#### (*Page) MustElementsByJS <- 0.50.0

``` go
func (p *Page) MustElementsByJS(js string, params ...interface{}) Elements
```

MustElementsByJS is similar to [Page.ElementsByJS](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementsByJS).

#### (*Page) MustElementsX <- 0.50.0

``` go
func (p *Page) MustElementsX(xpath string) Elements
```

MustElementsX is similar to [Page.ElementsX](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementsX).

#### (*Page) MustEmulate <- 0.50.0

``` go
func (p *Page) MustEmulate(device devices.Device) *Page
```

MustEmulate is similar to [Page.Emulate](https://pkg.go.dev/github.com/go-rod/rod#Page.Emulate).

#### (*Page) MustEval <- 0.50.0

``` go
func (p *Page) MustEval(js string, params ...interface{}) gson.JSON
```

MustEval is similar to [Page.Eval](https://pkg.go.dev/github.com/go-rod/rod#Page.Eval).

#### (*Page) MustEvalOnNewDocument <- 0.50.0

``` go
func (p *Page) MustEvalOnNewDocument(js string)
```

MustEvalOnNewDocument is similar to [Page.EvalOnNewDocument](https://pkg.go.dev/github.com/go-rod/rod#Page.EvalOnNewDocument).

#### (*Page) MustEvaluate <- 0.67.0

``` go
func (p *Page) MustEvaluate(opts *EvalOptions) *proto.RuntimeRemoteObject
```

MustEvaluate is similar to [Page.Evaluate](https://pkg.go.dev/github.com/go-rod/rod#Page.Evaluate).

#### (*Page) MustExpose <- 0.50.0

``` go
func (p *Page) MustExpose(name string, fn func(gson.JSON) (interface{}, error)) (stop func())
```

MustExpose is similar to [Page.Expose](https://pkg.go.dev/github.com/go-rod/rod#Page.Expose).

#### (*Page) MustGetWindow <- 0.50.0

``` go
func (p *Page) MustGetWindow() *proto.BrowserBounds
```

MustGetWindow is similar to [Page.GetWindow](https://pkg.go.dev/github.com/go-rod/rod#Page.GetWindow).

#### (*Page) MustHTML <- 0.94.0

``` go
func (p *Page) MustHTML() string
```

MustHTML is similar to [Page.HTML](https://pkg.go.dev/github.com/go-rod/rod#Page.HTML).

#### (*Page) MustHandleDialog <- 0.50.0

``` go
func (p *Page) MustHandleDialog() (wait func() *proto.PageJavascriptDialogOpening, handle func(bool, string))
```

MustHandleDialog is similar to [Page.HandleDialog](https://pkg.go.dev/github.com/go-rod/rod#Page.HandleDialog).

#### (*Page) MustHandleFileDialog <- 0.109.0

``` go
func (p *Page) MustHandleFileDialog() func(...string)
```

MustHandleFileDialog is similar to [Page.HandleFileDialog](https://pkg.go.dev/github.com/go-rod/rod#Page.HandleFileDialog).

#### (*Page) MustHas <- 0.50.0

``` go
func (p *Page) MustHas(selector string) bool
```

MustHas is similar to [Page.Has](https://pkg.go.dev/github.com/go-rod/rod#Page.Has).

#### (*Page) MustHasR <- 0.61.0

``` go
func (p *Page) MustHasR(selector, regex string) bool
```

MustHasR is similar to [Page.HasR](https://pkg.go.dev/github.com/go-rod/rod#Page.HasR).

#### (*Page) MustHasX <- 0.50.0

``` go
func (p *Page) MustHasX(selector string) bool
```

MustHasX is similar to [Page.HasX](https://pkg.go.dev/github.com/go-rod/rod#Page.HasX).

#### (*Page) MustInfo <- 0.50.0

``` go
func (p *Page) MustInfo() *proto.TargetTargetInfo
```

MustInfo is similar to [Page.Info](https://pkg.go.dev/github.com/go-rod/rod#Page.Info).

#### (*Page) MustInsertText <- 0.107.0

``` go
func (p *Page) MustInsertText(text string) *Page
```

MustInsertText is similar to [Page.InsertText](https://pkg.go.dev/github.com/go-rod/rod#Page.InsertText).

#### (*Page) MustNavigate <- 0.50.0

``` go
func (p *Page) MustNavigate(url string) *Page
```

MustNavigate is similar to [Page.Navigate](https://pkg.go.dev/github.com/go-rod/rod#Page.Navigate).

#### (*Page) MustNavigateBack <- 0.61.4

``` go
func (p *Page) MustNavigateBack() *Page
```

MustNavigateBack is similar to [Page.NavigateBack](https://pkg.go.dev/github.com/go-rod/rod#Page.NavigateBack).

#### (*Page) MustNavigateForward <- 0.61.4

``` go
func (p *Page) MustNavigateForward() *Page
```

MustNavigateForward is similar to [Page.NavigateForward](https://pkg.go.dev/github.com/go-rod/rod#Page.NavigateForward).

#### (*Page) MustObjectToJSON <- 0.50.0

``` go
func (p *Page) MustObjectToJSON(obj *proto.RuntimeRemoteObject) gson.JSON
```

MustObjectToJSON is similar to [Page.ObjectToJSON](https://pkg.go.dev/github.com/go-rod/rod#Page.ObjectToJSON).

#### (*Page) MustObjectsToJSON <- 0.50.0

``` go
func (p *Page) MustObjectsToJSON(list []*proto.RuntimeRemoteObject) gson.JSON
```

MustObjectsToJSON is similar to [Page.ObjectsToJSON].

#### (*Page) MustPDF <- 0.50.0

``` go
func (p *Page) MustPDF(toFile ...string) []byte
```

MustPDF is similar to [Page.PDF](https://pkg.go.dev/github.com/go-rod/rod#Page.PDF). If the toFile is "", it Page.will save output to "tmp/pdf" folder, time as the file name.

#### (*Page) MustRelease <- 0.50.0

``` go
func (p *Page) MustRelease(obj *proto.RuntimeRemoteObject) *Page
```

MustRelease is similar to [Page.Release](https://pkg.go.dev/github.com/go-rod/rod#Page.Release).

#### (*Page) MustReload <- 0.61.4

``` go
func (p *Page) MustReload() *Page
```

MustReload is similar to [Page.Reload](https://pkg.go.dev/github.com/go-rod/rod#Page.Reload).

#### (*Page) MustResetNavigationHistory <- 0.116.2

``` go
func (p *Page) MustResetNavigationHistory() *Page
```

MustResetNavigationHistory is similar to [Page.ResetNavigationHistory](https://pkg.go.dev/github.com/go-rod/rod#Page.ResetNavigationHistory).

#### (*Page) MustScreenshot <- 0.50.0

``` go
func (p *Page) MustScreenshot(toFile ...string) []byte
```

MustScreenshot is similar to [Page.Screenshot](https://pkg.go.dev/github.com/go-rod/rod#Page.Screenshot). If the toFile is "", it Page.will save output to "tmp/screenshots" folder, time as the file name.

#### (*Page) MustScreenshotFullPage <- 0.50.0

``` go
func (p *Page) MustScreenshotFullPage(toFile ...string) []byte
```

MustScreenshotFullPage is similar to [Page.ScreenshotFullPage]. If the toFile is "", it Page.will save output to "tmp/screenshots" folder, time as the file name.

#### (*Page) MustScrollScreenshot <- 0.116.2

``` go
func (p *Page) MustScrollScreenshot(toFile ...string) []byte
```

MustScrollScreenshot is similar to [Page.ScrollScreenshot](https://pkg.go.dev/github.com/go-rod/rod#Page.ScrollScreenshot). If the toFile is "", it Page.will save output to "tmp/screenshots" folder, time as the file name.

#### (*Page) MustSearch <- 0.50.0

``` go
func (p *Page) MustSearch(query string) *Element
```

MustSearch is similar to [Page.Search](https://pkg.go.dev/github.com/go-rod/rod#Page.Search). It only returns the first element in the search result.

#### (*Page) MustSetBlockedURLs <- 0.112.3

``` go
func (p *Page) MustSetBlockedURLs(urls ...string) *Page
```

MustSetBlockedURLs is similar to [Page.SetBlockedURLs](https://pkg.go.dev/github.com/go-rod/rod#Page.SetBlockedURLs).

#### (*Page) MustSetCookies <- 0.50.0

``` go
func (p *Page) MustSetCookies(cookies ...*proto.NetworkCookieParam) *Page
```

MustSetCookies is similar to [Page.SetCookies](https://pkg.go.dev/github.com/go-rod/rod#Page.SetCookies). If the len(cookies) is 0 it will clear all the cookies.

#### (*Page) MustSetDocumentContent <- 0.104.0

``` go
func (p *Page) MustSetDocumentContent(html string) *Page
```

MustSetDocumentContent is similar to [Page.SetDocumentContent](https://pkg.go.dev/github.com/go-rod/rod#Page.SetDocumentContent).

#### (*Page) MustSetExtraHeaders <- 0.50.0

``` go
func (p *Page) MustSetExtraHeaders(dict ...string) (cleanup func())
```

MustSetExtraHeaders is similar to [Page.SetExtraHeaders](https://pkg.go.dev/github.com/go-rod/rod#Page.SetExtraHeaders).

#### (*Page) MustSetUserAgent <- 0.50.0

``` go
func (p *Page) MustSetUserAgent(req *proto.NetworkSetUserAgentOverride) *Page
```

MustSetUserAgent is similar to [Page.SetUserAgent](https://pkg.go.dev/github.com/go-rod/rod#Page.SetUserAgent).

#### (*Page) MustSetViewport <- 0.64.0

``` go
func (p *Page) MustSetViewport(width, height int, deviceScaleFactor float64, mobile bool) *Page
```

MustSetViewport is similar to [Page.SetViewport](https://pkg.go.dev/github.com/go-rod/rod#Page.SetViewport).

#### (*Page) MustSetWindow <- 0.64.0

``` go
func (p *Page) MustSetWindow(left, top, width, height int) *Page
```

MustSetWindow is similar to [Page.SetWindow](https://pkg.go.dev/github.com/go-rod/rod#Page.SetWindow).

#### (*Page) MustStopLoading <- 0.50.0

``` go
func (p *Page) MustStopLoading() *Page
```

MustStopLoading is similar to [Page.StopLoading](https://pkg.go.dev/github.com/go-rod/rod#Page.StopLoading).

#### (*Page) MustTriggerFavicon <- 0.113.2

``` go
func (p *Page) MustTriggerFavicon() *Page
```

MustTriggerFavicon is similar to [PageTriggerFavicon].

#### (*Page) MustWait <- 0.50.0

``` go
func (p *Page) MustWait(js string, params ...interface{}) *Page
```

MustWait is similar to [Page.Wait](https://pkg.go.dev/github.com/go-rod/rod#Page.Wait).

#### (*Page) MustWaitDOMStable <- 0.114.0

``` go
func (p *Page) MustWaitDOMStable() *Page
```

MustWaitDOMStable is similar to [Page.WaitDOMStable](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitDOMStable).

#### (*Page) MustWaitElementsMoreThan <- 0.97.3

``` go
func (p *Page) MustWaitElementsMoreThan(selector string, num int) *Page
```

MustWaitElementsMoreThan is similar to [Page.WaitElementsMoreThan](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitElementsMoreThan).

#### (*Page) MustWaitIdle <- 0.50.0

``` go
func (p *Page) MustWaitIdle() *Page
```

MustWaitIdle is similar to [Page.WaitIdle](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitIdle).

#### (*Page) MustWaitLoad <- 0.50.0

``` go
func (p *Page) MustWaitLoad() *Page
```

MustWaitLoad is similar to [Page.WaitLoad](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitLoad).

#### (*Page) MustWaitNavigation <- 0.63.2

``` go
func (p *Page) MustWaitNavigation() func()
```

MustWaitNavigation is similar to [Page.WaitNavigation](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitNavigation).

#### (*Page) MustWaitOpen <- 0.50.0

``` go
func (p *Page) MustWaitOpen() (wait func() (newPage *Page))
```

MustWaitOpen is similar to [Page.WaitOpen](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitOpen).

#### (*Page) MustWaitRequestIdle <- 0.50.0

``` go
func (p *Page) MustWaitRequestIdle(excludes ...string) (wait func())
```

MustWaitRequestIdle is similar to [Page.WaitRequestIdle](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitRequestIdle).

#### (*Page) MustWaitStable <- 0.113.0

``` go
func (p *Page) MustWaitStable() *Page
```

MustWaitStable is similar to [Page.WaitStable](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitStable).

#### (*Page) MustWindowFullscreen <- 0.50.0

``` go
func (p *Page) MustWindowFullscreen() *Page
```

MustWindowFullscreen is similar to [Page.WindowFullscreen].

#### (*Page) MustWindowMaximize <- 0.50.0

``` go
func (p *Page) MustWindowMaximize() *Page
```

MustWindowMaximize is similar to [Page.WindowMaximize].

#### (*Page) MustWindowMinimize <- 0.50.0

``` go
func (p *Page) MustWindowMinimize() *Page
```

MustWindowMinimize is similar to [Page.WindowMinimize].

#### (*Page) MustWindowNormal <- 0.50.0

``` go
func (p *Page) MustWindowNormal() *Page
```

MustWindowNormal is similar to [Page.WindowNormal].

#### (*Page) Navigate 

``` go
func (p *Page) Navigate(url string) error
```

Navigate to the url. If the url is empty, "about:blank" will be used. It will return immediately after the server responds the http header.

#### (*Page) NavigateBack <- 0.61.4

``` go
func (p *Page) NavigateBack() error
```

NavigateBack history.

#### (*Page) NavigateForward <- 0.61.4

``` go
func (p *Page) NavigateForward() error
```

NavigateForward history.

#### (*Page) ObjectToJSON 

``` go
func (p *Page) ObjectToJSON(obj *proto.RuntimeRemoteObject) (gson.JSON, error)
```

ObjectToJSON by object id.

#### (*Page) Overlay 

``` go
func (p *Page) Overlay(left, top, width, height float64, msg string) (remove func())
```

Overlay a rectangle on the main frame with specified message.

#### (*Page) PDF 

``` go
func (p *Page) PDF(req *proto.PagePrintToPDF) (*StreamReader, error)
```

PDF prints page as PDF.

#### (*Page) Race <- 0.57.0

``` go
func (p *Page) Race() *RaceContext
```

Race creates a context to race selectors.

#### (*Page) Release 

``` go
func (p *Page) Release(obj *proto.RuntimeRemoteObject) error
```

Release the remote object. Usually, you don't need to call it. When a page is closed or reloaded, all remote objects will be released automatically. It's useful if the page never closes or reloads.

#### (*Page) Reload <- 0.61.4

``` go
func (p *Page) Reload() error
```

Reload page.

#### (*Page) ResetNavigationHistory <- 0.116.2

``` go
func (p *Page) ResetNavigationHistory() error
```

ResetNavigationHistory reset history.

#### (*Page) Screenshot 

``` go
func (p *Page) Screenshot(fullPage bool, req *proto.PageCaptureScreenshot) ([]byte, error)
```

Screenshot captures the screenshot of current page.

#### (*Page) ScrollScreenshot <- 0.114.7

``` go
func (p *Page) ScrollScreenshot(opt *ScrollScreenshotOptions) ([]byte, error)
```

ScrollScreenshot Scroll screenshot does not adjust the size of the viewport, but achieves it by scrolling and capturing screenshots in a loop, and then stitching them together. Note that this method also has a flaw: when there are elements with fixed positioning on the page (usually header navigation components), these elements will appear repeatedly, you can set the FixedTop parameter to optimize it.

Only support png and jpeg format yet, webP is not supported because no suitable processing library was found in golang.

#### (*Page) Search <- 0.47.0

``` go
func (p *Page) Search(query string) (*SearchResult, error)
```

Search for the given query in the DOM tree until the result count is not zero, before that it will keep retrying. The query can be plain text or css selector or xpath. It will search nested iframes and shadow doms too.

#### (*Page) SetBlockedURLs <- 0.112.3

``` go
func (p *Page) SetBlockedURLs(urls []string) error
```

SetBlockedURLs For some requests that do not want to be triggered, such as some dangerous operations, delete, quit logout, etc. Wildcards ('*') are allowed, such as ["*/api/logout/*","delete"]. NOTE: if you set empty pattern "", it will block all requests.

#### (*Page) SetCookies 

``` go
func (p *Page) SetCookies(cookies []*proto.NetworkCookieParam) error
```

SetCookies is similar to Browser.SetCookies .

#### (*Page) SetDocumentContent <- 0.104.0

``` go
func (p *Page) SetDocumentContent(html string) error
```

SetDocumentContent sets the page document html content.

#### (*Page) SetExtraHeaders 

``` go
func (p *Page) SetExtraHeaders(dict []string) (func(), error)
```

SetExtraHeaders whether to always send extra HTTP headers with the requests from this page.

#### (*Page) SetUserAgent 

``` go
func (p *Page) SetUserAgent(req *proto.NetworkSetUserAgentOverride) error
```

SetUserAgent (browser brand, accept-language, etc) of the page. If req is nil, a default user agent will be used, a typical mac chrome.

#### (*Page) SetViewport <- 0.62.0

``` go
func (p *Page) SetViewport(params *proto.EmulationSetDeviceMetricsOverride) error
```

SetViewport overrides the values of device screen dimensions.

#### (*Page) SetWindow <- 0.62.0

``` go
func (p *Page) SetWindow(bounds *proto.BrowserBounds) error
```

SetWindow location and size.

#### (*Page) Sleeper 

``` go
func (p *Page) Sleeper(sleeper func() utils.Sleeper) *Page
```

Sleeper returns a clone with the specified sleeper for chained sub-operations.

#### (*Page) StopLoading 

``` go
func (p *Page) StopLoading() error
```

StopLoading forces the page stop navigation and pending resource fetches.

#### (*Page) String <- 0.88.0

``` go
func (p *Page) String() string
```

String interface.

#### (*Page) Timeout 

``` go
func (p *Page) Timeout(d time.Duration) *Page
```

Timeout returns a clone with the specified total timeout of all chained sub-operations.

#### (*Page) TriggerFavicon <- 0.113.2

``` go
func (p *Page) TriggerFavicon() error
```

TriggerFavicon supports when browser in headless mode to trigger favicon's request. Pay attention to this function only supported when browser in headless mode, if you call it in no-headless mode, it will raise an error with the message "browser is no-headless".

#### (*Page) Wait 

``` go
func (p *Page) Wait(opts *EvalOptions) error
```

Wait until the js returns true.

#### (*Page) WaitDOMStable <- 0.114.0

``` go
func (p *Page) WaitDOMStable(d time.Duration, diff float64) error
```

WaitDOMStable waits until the change of the DOM tree is less or equal than diff percent for d duration. Be careful, d is not the max wait timeout, it's the least stable time. If you want to set a timeout you can use the [Page.Timeout](https://pkg.go.dev/github.com/go-rod/rod#Page.Timeout) function.

#### (*Page) WaitElementsMoreThan <- 0.97.3

``` go
func (p *Page) WaitElementsMoreThan(selector string, num int) error
```

WaitElementsMoreThan waits until there are more than num elements that match the selector.

#### (*Page) WaitEvent 

``` go
func (p *Page) WaitEvent(e proto.Event) (wait func())
```

WaitEvent waits for the next event for one time. It will also load the data into the event object.

#### (*Page) WaitIdle 

``` go
func (p *Page) WaitIdle(timeout time.Duration) (err error)
```

WaitIdle waits until the next window.requestIdleCallback is called.

#### (*Page) WaitLoad 

``` go
func (p *Page) WaitLoad() error
```

WaitLoad waits for the `window.onload` event, it returns immediately if the event is already fired.

#### (*Page) WaitNavigation <- 0.63.2

``` go
func (p *Page) WaitNavigation(name proto.PageLifecycleEventName) func()
```

WaitNavigation wait for a page lifecycle event when navigating. Usually you will wait for [proto.PageLifecycleEventNameNetworkAlmostIdle](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#PageLifecycleEventNameNetworkAlmostIdle).

#### (*Page) WaitOpen 

``` go
func (p *Page) WaitOpen() func() (*Page, error)
```

WaitOpen waits for the next new page opened by the current one.

#### (*Page) WaitRepaint <- 0.84.1

``` go
func (p *Page) WaitRepaint() error
```

WaitRepaint waits until the next repaint. Doc: https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame

#### (*Page) WaitRequestIdle 

``` go
func (p *Page) WaitRequestIdle(
	d time.Duration,
	includes, excludes []string,
	excludeTypes []proto.NetworkResourceType,
) func()
```

WaitRequestIdle returns a wait function that waits until no request for d duration. Be careful, d is not the max wait timeout, it's the least idle time. If you want to set a timeout you can use the [Page.Timeout](https://pkg.go.dev/github.com/go-rod/rod#Page.Timeout) function. Use the includes and excludes regexp list to filter the requests by their url.

#### (*Page) WaitStable <- 0.113.0

``` go
func (p *Page) WaitStable(d time.Duration) error
```

WaitStable waits until the page is stable for d duration.

#### (*Page) WithCancel <- 0.69.0

``` go
func (p *Page) WithCancel() (*Page, func())
```

WithCancel returns a clone with a context cancel function.

#### (*Page) WithPanic <- 0.100.0

``` go
func (p *Page) WithPanic(fail func(interface{})) *Page
```

WithPanic returns a page clone with the specified panic function. The fail must stop the current goroutine's execution immediately, such as use [runtime.Goexit](https://pkg.go.dev/runtime#Goexit) or panic inside it.

### type PageCloseCanceledError <- 0.114.8

``` go
type PageCloseCanceledError struct{}
```

PageCloseCanceledError error.

#### (*PageCloseCanceledError) Error <- 0.114.8

``` go
func (e *PageCloseCanceledError) Error() string
```

### type PageNotFoundError <- 0.114.8

``` go
type PageNotFoundError struct{}
```

PageNotFoundError error.

#### (*PageNotFoundError) Error <- 0.114.8

``` go
func (e *PageNotFoundError) Error() string
```

### type Pages 

``` go
type Pages []*Page
```

Pages provides some helpers to deal with page list.

#### (Pages) Empty <- 0.53.0

``` go
func (ps Pages) Empty() bool
```

Empty returns true if the list is empty.

#### (Pages) Find 

``` go
func (ps Pages) Find(selector string) (*Page, error)
```

Find the page that has the specified element with the css selector.

#### (Pages) FindByURL 

``` go
func (ps Pages) FindByURL(jsRegex string) (*Page, error)
```

FindByURL returns the page that has the url that matches the jsRegex.

#### (Pages) First <- 0.53.0

``` go
func (ps Pages) First() *Page
```

First returns the first page, if the list is empty returns nil.

#### (Pages) Last <- 0.53.0

``` go
func (ps Pages) Last() *Page
```

Last returns the last page, if the list is empty returns nil.

#### (Pages) MustFind <- 0.50.3

``` go
func (ps Pages) MustFind(selector string) *Page
```

MustFind is similar to [Browser.Find].

#### (Pages) MustFindByURL <- 0.50.0

``` go
func (ps Pages) MustFindByURL(regex string) *Page
```

MustFindByURL is similar to [Page.FindByURL].

### type Pool <- 0.116.2

``` go
type Pool[T any] chan *T
```

Pool is used to thread-safely limit the number of elements at the same time. It's a common practice to use a channel to limit concurrency, it's not special for rod. This helper is more like an example to use Go Channel. Reference: https://golang.org/doc/effective_go#channels

### func NewBrowserPool <- 0.101.7

``` go
func NewBrowserPool(limit int) Pool[Browser]
```

NewBrowserPool instance.

### func NewPagePool <- 0.73.2

``` go
func NewPagePool(limit int) Pool[Page]
```

NewPagePool instance.

### func NewPool <- 0.116.2

``` go
func NewPool[T any](limit int) Pool[T]
```

NewPool instance.

#### (Pool[T]) Cleanup <- 0.116.2

``` go
func (p Pool[T]) Cleanup(iteratee func(*T))
```

Cleanup helper.

#### (Pool[T]) Get <- 0.116.2

``` go
func (p Pool[T]) Get(create func() (*T, error)) (elem *T, err error)
```

Get a elem from the pool, allow error. Use the [Pool[T].Put] to make it reusable later.

#### (Pool[T]) MustGet <- 0.116.2

``` go
func (p Pool[T]) MustGet(create func() *T) *T
```

MustGet an elem from the pool. Use the [Pool[T].Put] to make it reusable later.

#### (Pool[T]) Put <- 0.116.2

``` go
func (p Pool[T]) Put(elem *T)
```

Put an elem back to the pool.

### type RaceContext <- 0.57.0

``` go
type RaceContext struct {
	// contains filtered or unexported fields
}
```

RaceContext stores the branches to race.

#### (*RaceContext) Do <- 0.57.0

``` go
func (rc *RaceContext) Do() (*Element, error)
```

Do the race.

#### (*RaceContext) Element <- 0.57.0

``` go
func (rc *RaceContext) Element(selector string) *RaceContext
```

Element is similar to [Page.Element](https://pkg.go.dev/github.com/go-rod/rod#Page.Element).

#### (*RaceContext) ElementByJS <- 0.57.0

``` go
func (rc *RaceContext) ElementByJS(opts *EvalOptions) *RaceContext
```

ElementByJS is similar to [Page.ElementByJS](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementByJS).

#### (*RaceContext) ElementFunc <- 0.107.1

``` go
func (rc *RaceContext) ElementFunc(fn func(*Page) (*Element, error)) *RaceContext
```

ElementFunc takes a custom function to determine race success.

#### (*RaceContext) ElementR <- 0.57.0

``` go
func (rc *RaceContext) ElementR(selector, regex string) *RaceContext
```

ElementR is similar to [Page.ElementR](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementR).

#### (*RaceContext) ElementX <- 0.57.0

``` go
func (rc *RaceContext) ElementX(selector string) *RaceContext
```

ElementX is similar to [Page.ElementX](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementX).

#### (*RaceContext) Handle <- 0.81.0

``` go
func (rc *RaceContext) Handle(callback func(*Element) error) *RaceContext
```

Handle adds a callback function to the most recent chained selector. The callback function is run, if the corresponding selector is present first, in the Race condition.

#### (*RaceContext) MustDo <- 0.57.0

``` go
func (rc *RaceContext) MustDo() *Element
```

MustDo is similar to [RaceContext.Do](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.Do).

#### (*RaceContext) MustElementByJS <- 0.57.0

``` go
func (rc *RaceContext) MustElementByJS(js string, params []interface{}) *RaceContext
```

MustElementByJS is similar to [RaceContext.ElementByJS](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.ElementByJS).

#### (*RaceContext) MustHandle <- 0.81.0

``` go
func (rc *RaceContext) MustHandle(callback func(*Element)) *RaceContext
```

MustHandle is similar to [RaceContext.Handle](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.Handle).

#### (*RaceContext) Search <- 0.112.0

``` go
func (rc *RaceContext) Search(query string) *RaceContext
```

Search is similar to [Page.Search](https://pkg.go.dev/github.com/go-rod/rod#Page.Search).

### type ScrollScreenshotOptions <- 0.114.7

``` go
type ScrollScreenshotOptions struct {
	// Format (optional) Image compression format (defaults to png).
	Format proto.PageCaptureScreenshotFormat `json:"format,omitempty"`

	// Quality (optional) Compression quality from range [0..100] (jpeg only).
	Quality *int `json:"quality,omitempty"`

	// FixedTop (optional) The number of pixels to skip from the top.
	// It is suitable for optimizing the screenshot effect when there is a fixed
	// positioning element at the top of the page.
	FixedTop float64

	// FixedBottom (optional) The number of pixels to skip from the bottom.
	FixedBottom float64

	// WaitPerScroll until no animation (default is 300ms)
	WaitPerScroll time.Duration
}
```

ScrollScreenshotOptions is the options for the ScrollScreenshot.

### type SearchResult <- 0.97.0

``` go
type SearchResult struct {
	*proto.DOMPerformSearchResult

	// First element in the search result
	First *Element
	// contains filtered or unexported fields
}
```

SearchResult handler.

#### (*SearchResult) All <- 0.97.0

``` go
func (s *SearchResult) All() (Elements, error)
```

All returns all elements.

#### (*SearchResult) Get <- 0.97.0

``` go
func (s *SearchResult) Get(i, l int) (Elements, error)
```

Get l elements at the index of i from the remote search result.

#### (*SearchResult) Release <- 0.97.0

``` go
func (s *SearchResult) Release()
```

Release the remote search result.

### type SelectorType <- 0.68.0

``` go
type SelectorType string
```

SelectorType enum.

``` go
const (
	// SelectorTypeRegex type.
	SelectorTypeRegex SelectorType = "regex"
	// SelectorTypeCSSSector type.
	SelectorTypeCSSSector SelectorType = "css-selector"
	// SelectorTypeText type.
	SelectorTypeText SelectorType = "text"
)
```

### type StreamReader <- 0.63.0

``` go
type StreamReader struct {
	Offset *int
	// contains filtered or unexported fields
}
```

StreamReader for browser data stream.

### func NewStreamReader <- 0.63.0

``` go
func NewStreamReader(c proto.Client, h proto.IOStreamHandle) *StreamReader
```

NewStreamReader instance.

#### (*StreamReader) Close <- 0.102.0

``` go
func (sr *StreamReader) Close() error
```

Close the stream, discard any temporary backing storage.

#### (*StreamReader) Read <- 0.63.0

``` go
func (sr *StreamReader) Read(p []byte) (n int, err error)
```

### type Touch <- 0.61.1

``` go
type Touch struct {
	// contains filtered or unexported fields
}
```

Touch presents a touch device, such as a hand with fingers, each finger is a [proto.InputTouchPoint](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#InputTouchPoint). Touch events is stateless, we use the struct here only as a namespace to make the API style unified.

#### (*Touch) Cancel <- 0.61.1

``` go
func (t *Touch) Cancel() error
```

Cancel touch action.

#### (*Touch) End <- 0.61.1

``` go
func (t *Touch) End() error
```

End touch action.

#### (*Touch) Move <- 0.61.1

``` go
func (t *Touch) Move(points ...*proto.InputTouchPoint) error
```

Move touch points. Use the [proto.InputTouchPoint.ID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#InputTouchPoint.ID) (Touch.identifier) to track points. Doc: https://developer.mozilla.org/en-US/docs/Web/API/Touch_events

#### (*Touch) MustCancel <- 0.61.1

``` go
func (t *Touch) MustCancel() *Touch
```

MustCancel is similar to [Touch.Cancel](https://pkg.go.dev/github.com/go-rod/rod#Touch.Cancel).

#### (*Touch) MustEnd <- 0.61.1

``` go
func (t *Touch) MustEnd() *Touch
```

MustEnd is similar to [Touch.End](https://pkg.go.dev/github.com/go-rod/rod#Touch.End).

#### (*Touch) MustMove <- 0.61.1

``` go
func (t *Touch) MustMove(points ...*proto.InputTouchPoint) *Touch
```

MustMove is similar to [Touch.Move](https://pkg.go.dev/github.com/go-rod/rod#Touch.Move).

#### (*Touch) MustStart <- 0.61.1

``` go
func (t *Touch) MustStart(points ...*proto.InputTouchPoint) *Touch
```

MustStart is similar to [Touch.Start](https://pkg.go.dev/github.com/go-rod/rod#Touch.Start).

#### (*Touch) MustTap <- 0.61.1

``` go
func (t *Touch) MustTap(x, y float64) *Touch
```

MustTap is similar to [Touch.Tap](https://pkg.go.dev/github.com/go-rod/rod#Touch.Tap).

#### (*Touch) Start <- 0.61.1

``` go
func (t *Touch) Start(points ...*proto.InputTouchPoint) error
```

Start a touch action.

#### (*Touch) Tap <- 0.61.1

``` go
func (t *Touch) Tap(x, y float64) error
```

Tap dispatches a touchstart and touchend event.

### type TraceType <- 0.59.0

``` go
type TraceType string
```

TraceType for logger.

``` go
const (
	// TraceTypeWaitRequestsIdle type.
	TraceTypeWaitRequestsIdle TraceType = "wait requests idle"

	// TraceTypeWaitRequests type.
	TraceTypeWaitRequests TraceType = "wait requests"

	// TraceTypeQuery type.
	TraceTypeQuery TraceType = "query"

	// TraceTypeWait type.
	TraceTypeWait TraceType = "wait"

	// TraceTypeInput type.
	TraceTypeInput TraceType = "input"
)
```

#### (TraceType) String <- 0.88.0

``` go
func (t TraceType) String() string
```

String interface.

### type TryError <- 0.114.8

``` go
type TryError struct {
	Value interface{}
	Stack string
}
```

TryError error.

#### (*TryError) Error <- 0.114.8

``` go
func (e *TryError) Error() string
```

#### (*TryError) Is <- 0.114.8

``` go
func (e *TryError) Is(err error) bool
```

Is interface.

#### (*TryError) Unwrap <- 0.114.8

``` go
func (e *TryError) Unwrap() error
```

Unwrap stdlib interface.
