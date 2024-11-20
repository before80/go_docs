+++
title = "flags"
date = 2024-11-20T18:02:07+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/launcher/flags](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/launcher/flags)
>
> 收录该文档时间：`2024-11-20T18:02:07+08:00`
>
> [Version: v0.116.2](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/launcher/flags?tab=versions)

Package flags ...

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

#### type Flag 

``` go
type Flag string
```

Flag name of a command line argument of the browser, also known as command line flag or switch. List of available flags: https://peter.sh/experiments/chromium-command-line-switches

​	Flag 表示浏览器命令行参数的名称，也称为命令行标志或开关。可用标志的列表：https://peter.sh/experiments/chromium-command-line-switches

``` go
const (
	// UserDataDir https://chromium.googlesource.com/chromium/src/+/master/docs/user_data_dir.md
	UserDataDir Flag = "user-data-dir"

	// Headless mode. Whether to run browser in headless mode. A mode without visible UI.
	// Headless 模式。是否以无头模式运行浏览器（无可见 UI）。
	Headless Flag = "headless"

	// App flag.
	// App 标志。
	App Flag = "app"

	// RemoteDebuggingPort flag.
	// RemoteDebuggingPort 标志。
	RemoteDebuggingPort Flag = "remote-debugging-port"

	// NoSandbox flag.
	// NoSandbox 标志。
	NoSandbox Flag = "no-sandbox"

	// ProxyServer flag.
	// ProxyServer 标志。
	ProxyServer Flag = "proxy-server"

	// WorkingDir flag.
	// WorkingDir 标志。
	WorkingDir Flag = "rod-working-dir"

	// Env flag.
	// Env 标志。
	Env Flag = "rod-env"

	// XVFB flag.
	// XVFB 标志。
	XVFB Flag = "rod-xvfb"

	// ProfileDir flag.
	// ProfileDir 标志。
	ProfileDir = "profile-directory"

	// Preferences flag.
	// Preferences 标志。
	Preferences Flag = "rod-preferences"

	// Leakless flag.
	// Leakless 标志。
	Leakless Flag = "rod-leakless"

	// Bin is the browser executable file path. If it's empty, launcher will automatically search or download the bin.
	// Bin 是浏览器可执行文件路径。如果为空，Launcher 会自动搜索或下载该文件。
	Bin Flag = "rod-bin"

	// KeepUserDataDir flag.
	// KeepUserDataDir 标志。
	KeepUserDataDir Flag = "rod-keep-user-data-dir"

	// Arguments for the command. Such as
	//     chrome-bin http://a.com http://b.com
	// The "http://a.com" and "http://b.com" are the arguments.
	// Arguments 表示命令的参数。例如：
	//     chrome-bin http://a.com http://b.com
	// "http://a.com" 和 "http://b.com" 是参数。
	Arguments Flag = ""
)
```

TODO: we should automatically generate all the flags here.

TODO：我们应该在这里自动生成所有标志。

#### (Flag) Check <-0.112.6

``` go
func (f Flag) Check()
```

Check if the flag name is valid.

​	Check 检查标志名称是否有效。

#### (Flag) NormalizeFlag <-0.112.6

``` go
func (f Flag) NormalizeFlag() Flag
```

NormalizeFlag normalize the flag name, remove the leading dash.

​	NormalizeFlag 标准化标志名称，移除前导的`-`。
