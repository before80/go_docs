+++
title = "自定义浏览器启动"
date = 2024-11-21T08:06:18+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-rod.github.io/i18n/zh-CN/#/custom-launch](https://go-rod.github.io/i18n/zh-CN/#/custom-launch)
>
> 收录该文档时间： `2024-11-21T08:08:39+08:00`

# 自定义浏览器启动

## 连接到正在运行的浏览器

​	查找您的浏览器的可执行路径，例如 macOS 运行：

```bash
"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" --headless --remote-debugging-port=9222
```

​	它将输出类似于：

```txt
DevTools listening on ws://127.0.0.1:9222/devtools/browser/4dcf09f2-ba2b-463a-8ff5-90d27c6cc913
```

上面的 `ws://127.0.0.1:9222/devtools/browser/4dcf09f2-ba2b-463a-8ff5-90d27c6cc913` 就是控制浏览器的接口：

```go
package main

import (
    "github.com/go-rod/rod"
)

func main() {
    u := "ws://127.0.0.1:9222/devtools/browser/4dcf09f2-ba2b-463a-8ff5-90d27c6cc913"
    rod.New().ControlURL(u).MustConnect().MustPage("https://example.com")
}
```

## launcher 库

​	由于上面的工作流经常被使用，我们抽象出 `launcher` 库来简化浏览器的启动。 例如自动下载或搜索浏览器可执行程序， 添加或删除浏览器可执行程序的命令行参数等。

​	因此，上述的手动启动和代码变成：

```go
func main() {
    u := launcher.New().Bin("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome").MustLaunch()
    rod.New().ControlURL(u).MustConnect().MustPage("https://example.com")
}
```

​	我们可以使用帮助函数 `launcher.LookPath` 来获取浏览器的可执行文件路径，上面的代码等价于：

```go
func main() {
    path, _ := launcher.LookPath()
    u := launcher.New().Bin(path).MustLaunch()
    rod.New().ControlURL(u).MustConnect().MustPage("https://example.com")
}
```

​	如果 `ControlURL` 未设置， `MustConnect` 将自动运行 `launcher.New().MustLaunch()`。 默认情况下，launcher 将自动下载并使用固定版本的浏览器，以保证浏览器 的行为一致性。 所以您可以将上述代码简化为：

```go
func main() {
    rod.New().MustConnect().MustPage("https://example.com")
}
```

## 增加或删除选项

​	可以使用 `Set` 和 `Delete` 来修改浏览器的启动参数（标志）：

```go
package main

import (
    "github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/launcher"
)

func main() {
    u := launcher.New().
        Set("user-data-dir", "path").
        Set("headless").
        Delete("--headless").
        MustLaunch()

    rod.New().ControlURL(u).MustConnect().MustPage("https://example.com")
}
```

​	前缀可选，例如 `headless` 和 `--headless` 相同。

​	由于类似 `user-data-dir`、`proxy-server`、`headless` 的选项经常会用到，我们为它们写了一些 helper，所以上面的代码可以改成这样：

```go
func main() {
    u := launcher.New().
        UserDataDir("path").
        Headless(true).
        Headless(false).
        MustLaunch()

    rod.New().ControlURL(u).MustConnect().MustPage("https://example.com")
}
```

​	所有可用的选项：[链接](https://peter.sh/experiments/chromium-command-line-switches)。

​	阅读 API 文档以获取更多信息：[链接](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher)。

## 清理

​	默认情况下，浏览器会创建一个 [user-data-dir](https://chromium.googlesource.com/chromium/src/+/master/docs/user_data_dir.md) 来保存用户数据，比如 cookie，缓存等。 Rod 提供了一个帮助函数 `Cleanup` 在浏览器完全关闭后来删除它：

```go
func main() {
    l := launcher.New().
        Headless(false).
        Devtools(true)

    defer l.Cleanup()
}
```

## 远程管理启动器

​	对于生产环境的爬虫系统，我们通常会把爬虫和浏览器拆分到不同的集群，从而使它们能够独立扩容。 Rod 提供模块 `launcher.Manager` 来远程管理启动器。 通过它我们可以远程启动用自定义启动参数浏览器。 它的用例在 [这里](https://github.com/go-rod/rod/blob/main/lib/launcher/rod-manager/main.go)。

​	因为很难在某些Linux发行版上正确安装chromium， Rod 提供了一个 docker image 来支持跨平台。 下面是一个用例：

1. 运行 rod 镜像 `docker run -p 7317:7317 ghcr.io/go-rod/rod`
2. 打开另一个终端，并运行类似这个[示例](https://github.com/go-rod/rod/blob/main/lib/examples/launch-managed/main.go)中的代码

​	它对于常见的自然语言的截图和字体进行过[调优](https://github.com/go-rod/rod/blob/main/lib/docker/Dockerfile)。 每个容器可以同时启动多个浏览器。 当控制链接断开后管理器会自动删除 [user-data-dir](https://chromium.googlesource.com/chromium/src/+/master/docs/user_data_dir.md) 。

## 用户模式

​	当您登录到您的 github 帐户时，您想要重新使用登录会话来完成自动化任务。 您可以使用 `launcher.NewUserMode` 启动您的常规用户浏览器。 Rod 将就像一个浏览器插件：

```go
wsURL := launcher.NewUserMode().MustLaunch()
rod.New().ControlURL(wsURL).MustConnect().NoDefaultDevice()
```

​	这里是一个更详细的示例： [代码示例](https://github.com/go-rod/rod/blob/main/lib/examples/use-rod-like-chrome-extension/main.go)。

## 底层 API

​	如果你想要控制启动过程中的每个步骤，比如说禁用自动下载、使用系统默认浏览器，见此[示例文件](https://github.com/go-rod/rod/blob/main/lib/launcher/example_test.go)。
