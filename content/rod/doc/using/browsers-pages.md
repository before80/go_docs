+++
title = "多浏览器与多页面"
date = 2024-11-21T08:09:52+08:00
weight = 90
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-rod.github.io/i18n/zh-CN/#/browsers-pages](https://go-rod.github.io/i18n/zh-CN/#/browsers-pages)
>
> 收录该文档时间： `2024-11-21T08:09:52+08:00`

# 多浏览器与多页面

​	你可以很直观的使用 Rod 同时控制多个浏览器或页面。

## 多个浏览器

​	启动多个浏览器：

```go
browser1 := rod.New().MustConnect()
browser2 := rod.New().MustConnect()
fmt.Println(browser1, browser2)
```

​	所有 API 都是线程安全的，同样适用于多个 Go routines。

​	也可以使用隐身模式启动多个浏览器：

```go
browser1 := rod.New().MustConnect()
browser2 := browser1.MustIncognito()
fmt.Println(browser1, browser2)
```

​	使用不同的启动参数启动浏览器：

```go
browser1 := rod.New().ControlURL(
    launcher.New().Headless(false).MustLaunch(),
).MustConnect()

browser2 := rod.New().ControlURL(
    launcher.New().UserDataDir("path").MustLaunch(),
).MustConnect()
fmt.Println(browser1, browser2)
```

## 多页面

​	在一个浏览器中开启多个页面：

```go
browser := rod.New().MustConnect()
page1 := browser.MustPage("http://a.com")
page2 := browser.MustPage("http://b.com")
fmt.Println(page1, page2)
```

​	如果浏览器已经开启了多个页面而且你没有它们的引用，你可以 [Browser.Pages()](https://pkg.go.dev/github.com/go-rod/rod#Browser.Pages) 来获取 [Pages](https://pkg.go.dev/github.com/go-rod/rod#Pages) 结构体，这是一个由标签页或窗口组成的数组，它拥有一些帮助函数，如 [Pages.Find()](https://pkg.go.dev/github.com/go-rod/rod#Pages.Find), [Pages.FindByURL()](https://pkg.go.dev/github.com/go-rod/rod#Pages.FindByURL)， [Pages.First()](https://pkg.go.dev/github.com/go-rod/rod#Pages.First)，等等。 一旦你获得你想要的页面的引用，你可以使用 [Page.Activate()](https://pkg.go.dev/github.com/go-rod/rod#Page.Activate) 来聚焦。 如果你点击链接打开了一个新的页面，你可以使用 [Page.WaitOpen](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitOpen) 以在新窗口开启后立即获取它的引用。

## 页面池

​	我们可以使用页面池来辅助同时控制和复用多个页面。

​	请看这个[例子](https://github.com/go-rod/rod/blob/46baf3aad803ed5cd8671aa325cbae4e297a89a4/examples_test.go#L533)。

## 浏览器池

​	Rod 中测试是管理浏览器池进行并发测试的一个好例子。 这就是为什么测试可以在数秒内跑完。 请看这里的[代码](https://github.com/go-rod/rod/blob/46baf3aad803ed5cd8671aa325cbae4e297a89a4/setup_test.go#L59)。
