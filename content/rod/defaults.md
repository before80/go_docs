+++
title = "defaults"
date = 2024-11-20T18:01:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/go-rod/rod/lib/defaults](https://pkg.go.dev/github.com/go-rod/rod/lib/defaults)
>
> 收录该文档时间：`2024-11-20T18:02:07+08:00`
>
> [Version: v0.116.2](https://pkg.go.dev/github.com/go-rod/rod/lib/defaults?tab=versions)

Package defaults of commonly used options parsed from environment. Check ResetWith for details.

​	解析自环境的常用选项默认值包。详细信息请参阅 ResetWith。

## 常量

This section is empty.

## 变量 

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L48)

``` go
var Bin string
```

Bin is the default of launcher.Launcher.Bin . Option name is "bin".

​	Bin 是 `launcher.Launcher.Bin` 的默认值。选项名为 "bin"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L64)

``` go
var CDP utils.Logger
```

CDP is the default of cdp.Client.Logger Option name is "cdp".

​	CDP 是 `cdp.Client.Logger` 的默认值。选项名为 "cdp"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L36)

``` go
var Devtools bool
```

Devtools is the default of launcher.Launcher.Devtools . Option name is "devtools".

​	Devtools 是 `launcher.Launcher.Devtools` 的默认值。选项名为 "devtools"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L40)

``` go
var Dir string
```

Dir is the default of launcher.Launcher.UserDataDir . Option name is "dir".

​	Dir 是 `launcher.Launcher.UserDataDir` 的默认值。选项名为 "dir"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L56)

``` go
var LockPort int
```

LockPort is the default of launcher.Browser.LockPort Option name is "lock".

​	LockPort 是 `launcher.Browser.LockPort` 的默认值。选项名为 "lock"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L28)

``` go
var Monitor string
```

Monitor is the default of rod.Browser.ServeMonitor . Option name is "monitor".

​	Monitor 是 `rod.Browser.ServeMonitor` 的默认值。选项名为 "monitor"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L44)

``` go
var Port string
```

Port is the default of launcher.Launcher.RemoteDebuggingPort . Option name is "port".

​	Port 是 `launcher.Launcher.RemoteDebuggingPort` 的默认值。选项名为 "port"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L52)

``` go
var Proxy string
```

Proxy is the default of launcher.Launcher.Proxy Option name is "proxy".

​	Proxy 是 `launcher.Launcher.Proxy` 的默认值。选项名为 "proxy"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L32)

``` go
var Show bool
```

Show is the default of launcher.Launcher.Headless . Option name is "show".

​	Show 是 `launcher.Launcher.Headless` 的默认值。选项名为 "show"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L24)

``` go
var Slow time.Duration
```

Slow is the default of rod.Browser.SlowMotion . The format is same as https://golang.org/pkg/time/#ParseDuration Option name is "slow".

​	Slow 是 `rod.Browser.SlowMotion` 的默认值。格式与 [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration) 相同。选项名为 "slow"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L19)

``` go
var Trace bool
```

Trace is the default of rod.Browser.Trace . Option name is "trace".

​	Trace 是 `rod.Browser.Trace` 的默认值。选项名为 "trace"。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L60)

``` go
var URL string
```

URL is the default websocket url for remote control a browser. Option name is "url".

​	URL 是用于远程控制浏览器的默认 WebSocket URL。选项名为 "url"。

## 函数 

## func Reset <- 0.47.0

``` go
func Reset()
```

Reset all flags to their init values.

​	将所有标志重置为其初始化值。

## func ResetWith <- 0.106.0

``` go
func ResetWith(options string)
```

ResetWith options and "-rod" command line flag. It will be called in an init() , so you don't have to call it manually. It will try to load the cli flag "-rod" and then the options, the later override the former. If you want to disable the global cli argument flag, set env DISABLE_ROD_FLAG. Values are separated by commas, key and value are separated by "=". For example:

​	使用选项和 `-rod` 命令行标志重置。它将在 `init()` 中调用，因此无需手动调用。此函数会尝试加载 CLI 标志 `-rod`，然后加载选项，后者会覆盖前者。如果希望禁用全局 CLI 参数标志，请设置环境变量 `DISABLE_ROD_FLAG`。值用逗号分隔，键和值用 `=` 分隔。例如：

```
go run main.go -rod=show
go run main.go -rod show,trace,slow=1s,monitor
go run main.go --rod="slow=1s,dir=path/has /space,monitor=:9223"
```

## 类型

This section is empty.
