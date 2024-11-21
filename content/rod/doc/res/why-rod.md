+++
title = "为什么选择 Rod"
date = 2024-11-21T08:12:02+08:00
weight = 40
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-rod.github.io/i18n/zh-CN/#/why-rod](https://go-rod.github.io/i18n/zh-CN/#/why-rod)
>
> 收录该文档时间： `2024-11-21T08:12:02+08:00`

# 为什么选择 Rod

​	有许多不错的项目，但其中没有一个是完美的。 重要的是选择其中最适合你需求的那一个。

## 与其他库对比

### Chromedp

​	理论上，Rod 应该比 Chromedp 更快，消耗更少的内存，Rod 的架构更稳定和一致。

​	[Chromedp](https://github.com/chromedp/chromedp) 默认使用系统浏览器，如果意外升级了浏览器，这可能会导致问题。 [Chromedp](https://github.com/chromedp/chromedp) 使用不同的 Go 模块 来管理 devtools 协议版本，这可能会导致诸如 [这种问题](https://github.com/chromedp/chromedp/issues/1031)。 每个Rod 版本都将配备一个特定的 Chromium 版本及其 devtools 协议版本，并对其进行完整的单元测试。

​	[Chromedp](https://github.com/chromedp/chromedp) 为事件使用[固定大小的缓冲](https://github.com/chromedp/chromedp/blob/b56cd66/target.go#L69-L73)，这可能会在高并发情况下导致死锁。 由于 Chromedp 只使用一个事件循环，缓慢的事件处理程序可能会互相阻塞。 Rod 没有这些问题，因为它基于 [goob](https://github.com/ysmood/goob)。

​	Chromedp 会对浏览器传回的每条消息进行 JSON 解码，而 rod 则按需解码，因此 Rod 性能更好，尤其是对于大量用到网络的事件来说。

​	在 Windows 和 Mac 上，Chromedp 会在崩溃时留下浏览器僵尸进程。

​	Rod 可配置程度更高，比如甚至可以把 WebSocket 库替换成一个你喜欢的库。

​	对于直接的代码比较，见[此](https://github.com/go-rod/rod/tree/master/lib/examples/compare-chromedp)。 比较 [rod](https://github.com/go-rod/rod/tree/master/lib/examples/compare-chromedp/logic/main.go) 和 [chromedp](https://github.com/chromedp/examples/blob/master/logic/main.go) 名为 `logic` 的示例后可以发现 rod 简洁非常多。

​	使用 Chromedp，你必须使用他们冗长的类 DSL task 来处理代码逻辑。 Chromedp 使用了基层抽象来封装代码执行和参数，这使得在出现错误时很难理解他们的代码。 大量的使用接口使得在追踪问题时静态类型无用武之地。 相比之下，Rod 尽可能少的使用接口。

​	Rod 依赖更少，代码结构更简洁，有更完善的自动化测试。 你会发现给 Rod 贡献代码更加容易。 因此，与 Chromedp 相比，在未来 Rod 有潜力从社区中获得更多优秀的功能。

​	Chromedp 的另一个问题是，他们的架构基于 [DOM 节点 id](https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-NodeId)，而 puppeteer 和 rod 基于 [远程对象 id](https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-RemoteObjectId)。 因此，它不仅 [更慢](https://github.com/puppeteer/puppeteer/issues/2936) ，而且也会使得给 Chromedp 添加与运行时强耦合的功能更加困难。 例如，这个 [ticket](https://github.com/chromedp/chromedp/issues/72) 开了整整三年。 即便它现在已经关闭了，你还是不能在 iframe 内的元素上执行 js 表达式。 除此之外，Chromedp 还在内存里维护了一份全部 nodes 的[拷贝](https://github.com/chromedp/chromedp/blob/e2970556e3d05f3259c464faeed1ec0e862f0560/target.go#L375-L376)。 这会造成本地 NodeID 列表和 [DOM.documentUpdated](https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-documentUpdated) 的竞争，会导致类似 [#762](https://github.com/chromedp/chromedp/issues/762) 的问题。

### Puppeteer

​	[Puppeteer](https://github.com/puppeteer/puppeteer) 会对浏览器传回的每条消息进行 JSON 解码，而 Rod 则按需解码，因此理论上 Rod 性能更好，尤其是大量用到网络的事件时。

​	使用 puppeteer 时必须大量处理 promise/async/await，而这让设计优雅的[流式接口](https://en.wikipedia.org/wiki/Fluent_interface)变得非常困难。 端对端测试需要用到许多同步操作来模拟真人输入。 由于 Puppeteer 基于 Nodejs，所有 IO 操作都是异步的，所以通常人们会不得不输入成堆的 async/await。 忘记写 `await` 的话，调试 Promise 泄露通常会非常痛苦。 你的项目越大，这种开销就越大。

​	Rod 默认类型安全，且有更好的注释。 它对于 Devtools 协议中的所有 endpoint 都有类型绑定。

​	Rod 会尽可能禁用 domain 事件，而 puppeteer 则总是启用所有 domain 事件。 远程驱动浏览器时这会消耗大量资源。

​	Rod 对取消，超时支持的更好。 如果想要处理成千上万的页面，这至关重要。 例如，要模拟 `click` 我们需要发送数个 cdp 请求。 使用 [Promise](https://stackoverflow.com/questions/29478751/cancel-a-vanilla-ecmascript-6-promise-chain) 时不可能实现“只发送一半的 cdp 请求”，但使用 [context](https://golang.org/pkg/context/) 时则可以。

### Playwright

​	Rod 和 [Playwright](https://github.com/microsoft/playwright) 几乎是同时发布的。 对于 Playwright 来说，Rod 和 Puppeteer 之间的大多数比较仍然适用，因为 Playwright 和 Puppeteer 都是由几乎相同的贡献者维护的。

​	正如 Playwright 在他们的文档上所说的那样： "Playwright enables reliable end-to-end testing for modern web apps."，它侧重的是测试。 但是 Rod 专注于更通用的领域，页面自动化和爬虫，这使得其设计更加侧重于灵活度和性能。

​	Rod 的架构目标之一是让每个人都能更轻松的为社区贡献力量，让 Rod 成为一个纯粹的社区项目，而这也是我选择 Golang 与 MIT 许可的一大原因。 TypeScript 也是一个不错的选择，不过如果你了解过 Playwright 的设计选择的话，你会发现 [`any`](https://www.typescriptlang.org/docs/handbook/basic-types.htmvl#any) 和 [union 类型](https://www.typescriptlang.org/docs/handbook/unions-and-intersections.html#union-types)随处可见。 如果你尝试跳转到 [page.click](https://playwright.dev/#version=v1.6.2&path=docs%2Fapi.md&q=pageclickselector-options) 的源码，那些 `d.ts` 文件会让你明白 TypeScript 的现实问题。 当然 Golang 也还不够好，但它通常会比 Node.js TypeScript 带来很少的技术债。 如果我是一个不熟悉如何用代码来自动化端对端测试或网站监测的人，而你想让我选择用 Golang 还是 TypeScript 来做 QA 或 Infra 的话，我会选择 Golang。

​	他们为跨浏览器支持所做的努力令人敬畏。 但如今主要厂商大都采用 HTML5，很难说它带来的复杂度大于好处。 跨浏览器[补丁](https://github.com/microsoft/playwright/tree/master/browser_patches)将来会变成一个负担吗？ Patch 过的浏览器的安全性也是个问题。 这使得测试旧版本的 Firefox 或 Safari 也因此变得非常棘手。 但愿这不是过度设计。

### Selenium

​	[Selenium](https://www.selenium.dev/) 基于 [webdriver 协议](https://www.w3.org/TR/webdriver/) ，这一协议的功能比 [devtools 协议](https://chromedevtools.github.io/devtools-protocol)少得多。 比如说它不能处理[闭合的 shadow DOM](https://github.com/sukgu/shadow-automation-selenium/issues/7#issuecomment-563062460)。 无法将页面保存为 PDF。 不支持诸如 [Profiler](https://chromedevtools.github.io/devtools-protocol/tot/Profiler/) 或 [Performance](https://chromedevtools.github.io/devtools-protocol/tot/Performance/)等工具。

​	由于像浏览器驱动之类的额外依赖，Selenium 更难配置与维护。

​	尽管 Selenium 宣传自己有更好的跨浏览器支持，但通常很难让测试支持所有主流浏览器。

​	关于“selenium vs puppeteer”的文章有很多，你可以把 rod 当作 Golang 版的 Puppeteer。

### Cypress

​	[Cypress](https://www.cypress.io/) 的功能很有限，对于闭合的 shadow dom 或跨域 iframe 它就无可奈何了。 要了解更多详情，请阅读他们的[有关局限性的文档](https://docs.cypress.io/guides/references/trade-offs.html)。

​	如果你想要和我们合作，基于 Rod 创建一个以测试为重点的框架，从而克服 cypress 的局限性，请联系我们。

## Rod 是什么意思

​	Rod 是用于控制木偶的装置，比如说下图中褐色棍子：

![rod](why-rod_img/80178856-31cd8880-863a-11ea-83e9-64f84be3282d.png)

​	意思是说，我们是操控木偶的人，而浏览器则是木偶。 我们使用“木偶棍（rod）”来控制木偶。
