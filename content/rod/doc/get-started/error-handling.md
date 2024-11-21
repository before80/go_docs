+++
title = "异常处理"
date = 2024-11-21T08:05:42+08:00
weight = 30
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-rod.github.io/i18n/zh-CN/#/error-handling](https://go-rod.github.io/i18n/zh-CN/#/error-handling)
>
> 收录该文档时间： `2024-11-21T08:08:39+08:00`

# 异常处理

​	在前面的章节中，我们看到了很多带 `Must` 前缀的方法，如 `MustNavigate`、`MustElement` 等。 它们都有不带前缀的版本，比如 `Navigate`、`Element` 等。 它们之间的主要区别在于如何处理异常。 这不是 Rod 特有的，你可以在标准库中找到类似设计，如 [regex.MustCompile](https://golang.org/pkg/regexp/#MustCompile)。

​	形如 `MustNavigate` 与 `MustElement` 的方法通常在示例代码或快速脚本中使用。 它们适用于冒烟测试、站点监控、端到端测试等任务。 对于有很多不确定性的任务，比如网络抓取，无前缀版本将是一个更好的选择。

​	带前缀的版本只是无前缀版本加上异常检查。 下面是 `MustElement` 的源代码。 可以看到，它只是调用了 `Element`，并多加了几行代码，在 `err` 不是 nil 时 panic。

```go
// Page ...
type Page rod.Page

// MustElement ...
func (p *Page) MustElement(selector string) *rod.Element {
    el, err := (*rod.Page)(p).Element(selector)
    if err != nil {
        panic(err)
    }
    return el
}
```

## 获取错误值

​	下面两段代码虽风格不同，但做的事情几乎一致。

​	这种风格是 Go 处理异常的标准方法：

```go
page := rod.New().MustConnect().MustPage("https://example.com")

el, err := page.Element("a")
if err != nil {
    panic(err)
}
html, err := el.HTML()
if err != nil {
    panic(err)
}
fmt.Println(html)
```

​	我们可以使用 `rod.Try` 捕获带有 `Must` 前缀的方法——`MustElement` 和 `MustHTML`——抛出的异常。 这种风格通常需要更少的代码，但可能会捕获到额外的异常：

```go
page := rod.New().MustConnect().MustPage("https://example.com")

err := rod.Try(func() {
    fmt.Println(page.MustElement("a").MustHTML())
})
panic(err)
```

## 检查异常类型

​	我们使用 Go 的标准方法来检查异常类型（没有魔法）。

​	将上面代码中的 `panic` 替换为 `handleError` ：

```go
func main() {
    _, err := page.Element("a")
    handleError(err)
}

func handleError(err error) {
    var evalErr *rod.EvalError
    if errors.Is(err, context.DeadlineExceeded) { // 超时错误
        fmt.Println("timeout err")
    } else if errors.As(err, &evalErr) { // eval 错误
        fmt.Println(evalErr.LineNumber)
    } else if err != nil {
        fmt.Println("can't handle", err)
    }
}
```
