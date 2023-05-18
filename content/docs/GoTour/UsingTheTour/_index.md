+++
title = "欢迎"
linkTitle = "欢迎"
weight = 1
date = 2023-05-17T12:10:24+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Welcome

> 原文：https://go.dev/tour/welcome/1

## Hello, 世界

​	欢迎来到[Go编程语言之旅](https://go.dev/)。

​	本指南分为多个模块，你可以通过点击页面左上方的 A Tour of Go 来访问。

​	你也可以通过点击页面右上方的菜单随时查看目录。

​	在整个指南过程中，你会发现有一系列的幻灯片和练习供你完成。

​	你可以通过以下方式浏览它们

- 上一页 "或 "`PageUp` "进入上一页。
- 下一页 "或 "`PageDown` "进入下一页。

​	本指南是互动的。现在点击`Run`按钮（或按`Shift + Enter`），在远程服务器上编译和运行程序。结果显示在代码下面。

​	这些示例程序展示了Go的不同方面。指南中的程序是你自己实验的起点。

​	编辑程序并再次运行它。

​	当你点击`Format`（快捷键：`Ctrl+Enter`）时，编辑器中的文本将使用[gofmt](https://go.dev/cmd/gofmt/)工具进行格式化。你可以通过点击`syntax`按钮来打开或关闭语法高亮。

​	当你准备继续前进时，点击下面的`右箭头`或键入`PageDown`键。

```go 
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

## Go local - Go 本地化

 本指南还有其他语言版本：

- [Brazilian Portuguese — Português do Brasil](https://go-tour-br.appspot.com/)
- [Catalan — Català](https://go-tour-ca.appspot.com/)
- [Simplified Chinese — 中文（简体）](https://tour.go-zh.org/)
- [Czech — Česky](https://go-tour-cz.appspot.com/)
- [Indonesian — Bahasa Indonesia](https://go-tour-id2.appspot.com/)
- [Japanese — 日本語](https://go-tour-jp.appspot.com/)
- [Korean — 한국어](https://go-tour-ko.appspot.com/)
- [Polish — Polski](https://go-tour-pl1.appspot.com/)
- [Thai — ภาษาไทย](https://go-tour-th.appspot.com/)

单击“下一步”按钮或键入 `PageDown` 继续。

##  Go offline (optional) - Go 离线(可选)

​	本指南也可以作为一个独立的程序，你可以在不接入互联网的情况下使用。它在你自己的机器上构建并运行代码示例。

​	要在本地运行本指南，你需要先[安装Go](../../GettingStarted/InstallingGo)，然后运行：

```
go install golang.org/x/website/tour@latest
```

​	这将在你的[GOPATH](https://go.dev/cmd/go/#hdr-GOPATH_and_Modules)的`bin`目录下放置一个`tour`二进制文件。当你运行`tour`程序时，它将打开一个网页浏览器，显示你的本地版本的`tour`。

​	当然，你可以继续通过这个网站进行学习。

##  The Go Playground

​	本指南是建立在[Go Playground](https://play.golang.org/)之上的，这是一个运行在[golang.org](https://go.dev/)服务器上的Web服务。

​	该服务接收Go程序，进行编译、链接，并在沙盒中运行该程序，然后返回输出结果。

可以在playground中运行的程序`是有限制的`：

- 在playground 中，时间从`2009-11-10 23:00:00 UTC`开始（确定这个日期的意义是给读者的一个练习）。这使得通过给程序提供确定的输出，更容易对程序进行缓存。
- 对`执行时间`以及`CPU和内存`的使用也有限制，而且程序`不能访问外部网络主机`。

该playground 使用Go的`最新稳定版本`。

阅读 "[Inside the Go Playground](../../../GoBlog/2013/InsideTheGoPlayground) "以了解更多。

```go 
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
}
```



## Congratulations 恭喜你

​	你已经完成了第一个模块的学习!

​	现在请点击`A Tour of Go`，看看你还能学到什么关于Go的知识，或者直接进入下一课。

