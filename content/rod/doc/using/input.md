+++
title = "输入"
date = 2024-11-21T08:08:27+08:00
weight = 40
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-rod.github.io/i18n/zh-CN/#/input](https://go-rod.github.io/i18n/zh-CN/#/input)
>
> 收录该文档时间： `2024-11-21T08:08:27+08:00`

# 输入

​	Rod 提供了很多方法来模拟人工输入，比如鼠标点击或者键盘按键。

## 鼠标点击

​	模拟鼠标点击一个元素：

```go
// 左击
page.MustElement("button").MustClick()

// 右击
_ = page.MustElement("button").Click(proto.InputMouseButtonRight, 1)
```

## 文本输入

​	模拟输入：

```go
el := page.MustElement(`[type="text"]`)
el.MustInput("Jack")

fmt.Println(el.MustText()) // 使用 MustText 来获取文本
```

## 删除输入框中的文本

​	模拟人的行为即可。 选中所有文本，用一个空字符串替换：

```go
page.MustElement(`[type="text"]`).MustSelectAllText().MustInput("")
```

​	可以使用 `SelectText` 替换部分文本。

## 时间输入

​	支持的输入类型有 [date](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/date)、[datetime-local](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/datetime-local)、[month](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/month) 和 [time](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time)。

```go
page.MustElement(`[type="date"]`).MustInputTime(time.Now())
```

## 输入按键组合

​	例如，像人类一样输入大写 “a” 的完整动作包括：

1. 按住 `Shift` 键
2. 按住然后松开 `A` 键
3. 松开 `Shift` 键

​	可以使用 `Page.KeyAction` 或 `Element.KeyActions` 帮助函数：

```go
page.KeyActions().Press(input.ShiftLeft).Type('A').MustDo()
```

​	`KeyActions` 帮助函数将自动松开所有按键，此处 `input.ShiftLeft` 将被自动松开。

​	要模拟诸如 `CTRL + Enter` 等快捷键，可以这样做：

```go
page.KeyActions().Press(input.ControlLeft).Type(input.Enter).MustDo()
```

## 复选框

​	像真人一样点击即可：

```go
el := page.MustElement(`[type="checkbox"]`)

// check it if not checked
if !el.MustProperty("checked").Bool() {
    el.MustClick()
}
```

## 选择选项

​	选择 [` `](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select) 中的选项。

​	下面的代码会选择包含文本 "B" 或 "C" 的选项：

```go
page.MustElement("select").MustSelect("B", "C")
```

​	也可以使用正则表达式或 CSS 选择器来选择选项：

```go
_ = page.MustElement("select").Select([]string{`^B$`}, true, rod.SelectorTypeRegex)

// 设置为 false 来取消
_ = page.MustElement("select").Select([]string{`[value="c"]`}, false, rod.SelectorTypeCSSSector)
```

## 设置文件

​	使用 `SetFiles` 为[文件输入元素](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/file)设置文件：

```go
page.MustElement(`[type=file]`).MustSetFiles("a.jpg", "b.pdf")
```

## 鼠标、键盘和触摸

​	也可以使用 `page.Mouse`、`page.Keyboard` 或 `page.Touch` 模拟底层输入。 例如，可以在 Rod 的单元测试中搜索 drag 来了解如何模拟拖动。
