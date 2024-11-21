+++
title = "模拟 / Emulation"
date = 2024-11-21T08:08:39+08:00
weight = 50
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-rod.github.io/i18n/zh-CN/#/emulation](https://go-rod.github.io/i18n/zh-CN/#/emulation)
>
> 收录该文档时间： `2024-11-21T08:08:39+08:00`

# 模拟 / Emulation

​	Rod 提供了多种方法来模拟页面环境。

## 设备

​	要同时为页面设置视区、User-Agent、方向等，可以使用预定义的设备：

```go
page.MustEmulate(devices.IPhone6or7or8Plus)
```

​	或者定义你自己的设备：

```go
page.MustEmulate(devices.Device{
  Title:          "iPhone 4",
  Capabilities:   []string{"touch", "mobile"},
  UserAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X)",
  AcceptLanguage: "en",
  Screen: devices.Screen{
    DevicePixelRatio: 2,
    Horizontal: devices.ScreenSize{
      Width:  480,
      Height: 320,
    },
    Vertical: devices.ScreenSize{
      Width:  320,
      Height: 480,
    },
  },
})
```

​	见预定义设备的代码，每个字段的意思都显而易见。

​	还可以通过 [Browser.DefaultDevice](https://pkg.go.dev/github.com/go-rod/rod#Browser.DefaultDevice) 来为所有页面设置默认设备。

​	设备模拟默认会被起用（[Devices.LaptopWithMDPIScreen](https://github.com/go-rod/rod/blob/bc44c39c9b4352c15d00bef6f6a1071205d2c388/lib/devices/list.go#L616) 会被使用），这会覆盖某些浏览器的默认设定，这么做是为了稳定的一致性（比如有助于复现测试结果）。

​	你可以通过给 `Browser.DefaultDevice` 传特殊的 *Clear* 设备来禁用设备模拟功能。

```go
browser.DefaultDevice(devices.Clear)
```

​	或者你也可以直接使用 [Browser.NoDefaultDevice](https://pkg.go.dev/github.com/go-rod/rod#Browser.NoDefaultDevice) 帮助函数。

## User Agent

​	使用 [Page.SetUserAgent](https://pkg.go.dev/github.com/go-rod/rod#Page.SetUserAgent) 为特定页面指定 User Agent。

## 视区

​	使用 [Page.SetViewport](https://pkg.go.dev/github.com/go-rod/rod#Page.SetViewport) 为特定页面指定视区。

## 语言和时区

​	可以使用 launch env 为所有页面设置：

```go
u := launcher.New().Env(append(os.Environ(), "TZ=America/New_York")...).MustLaunch()
rod.New().ControlURL(u).MustConnect()
```

​	或者可以使用 [EmulationSetTimezoneOverride](https://pkg.go.dev/github.com/go-rod/rod/lib/proto#EmulationSetTimezoneOverride) 或 [EmulationSetLocaleOverride](https://pkg.go.dev/github.com/go-rod/rod/lib/proto#EmulationSetLocaleOverride) 为特定页面设置：

```go
page := browser.MustPage()
_ = proto.EmulationSetTimezoneOverride{TimezoneID: "America/New_York"}.Call(page)
```

## 权限

​	使用 [BrowserGrantPermissions](https://pkg.go.dev/github.com/go-rod/rod/lib/proto#BrowserGrantPermissions)

## 地理位置

​	使用 [EmulationSetGeolocationOverride](https://pkg.go.dev/github.com/go-rod/rod/lib/proto#EmulationSetGeolocationOverride)

## 配色方案和媒体

​	使用 [EmulationSetEmulatedMedia](https://pkg.go.dev/github.com/go-rod/rod/lib/proto#EmulationSetEmulatedMedia)

```go
page := browser.MustPage()
_ = proto.EmulationSetEmulatedMedia{
    Media: "screen",
    Features: []*proto.EmulationMediaFeature{
        {Name: "prefers-color-scheme", Value: "dark"},
    },
}.Call(page)
```

## 防止机器人检测

​	通常最好让无头浏览器对页面完全透明，以使页面无法判断它是被人或机器人控制的。 在某些情况下，某些页面可以使用客户端 js 检测页面是否由人或机器人控制， 此 web WebGL 、WebDriver 或 http 请求头。 您可以手工写一个 js lib 来隐藏所有的痕迹，或者使用 [stealth](https://github.com/go-rod/stealth) 库: [代码示例](https://github.com/go-rod/stealth/blob/master/examples_test.go)。

​	如果 `stealth` 库不起作用，您可以用 `launcher.NewUserMode`: [用户模式](https://go-rod.github.io/i18n/zh-CN/#/custom-launch?id=user-mode)

​	您可以使用诸如 [https://bot.sannysoft.com](https://bot.sannysoft.com/) 等工具来测试您的配置。

## 浏览器指纹

​	浏览器指纹不是机器人检测。 它使用各种技巧来收集唯一的浏览器特征来识别浏览器。 网站可以使用它来跟踪用户，即使用户没有登录，它也被广泛用来标记无头爬虫。 例如，不同的用户通常会在他们的操作系统上安装不同的字体，我们可以使用它来区分不同的用户。 另一个例子是使用 canvas 来渲染文本，不同用户通常拥有不同的 GPU，图形驱动，或 OS，它们都会影响渲染图像的结果。

​	通常您可以通过启动多个浏览器实例来携带不同的指纹。 如果您为了节省内存和 CPU 而使用单个浏览器，您必须手动覆盖画布、字体等 API。

​	您可以使用开源项目，例如 [FingerprintJS](https://github.com/fingerprintjs/fingerprintjs/) 来测试您的配置。
