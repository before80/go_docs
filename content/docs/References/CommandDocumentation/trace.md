+++
title = "trace"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# trace

> 原文：[https://pkg.go.dev/cmd/trace@go1.19.3](https://pkg.go.dev/cmd/trace@go1.19.3)

### Overview 概述

Trace is a tool for viewing trace files.

Trace是一个查看跟踪文件的工具。

Trace files can be generated with:

跟踪文件可以通过以下方式生成：

- runtime/trace.Start

- net/http/pprof package

- go test -trace

  

Example usage: Generate a trace file with 'go test':

使用实例。用'go test'生成一个跟踪文件：

```
go test -trace trace.out pkg
```

View the trace in a web browser:

在网络浏览器中查看跟踪：

```
go tool trace trace.out
```

Generate a pprof-like profile from the trace:

从跟踪中生成一个类似于pprof的配置文件：

```
go tool trace -pprof=TYPE trace.out > TYPE.pprof
```

Supported profile types are:

支持的配置文件类型有：

- net: network blocking profile
- sync: synchronization blocking profile
- syscall: syscall blocking profile
- sched: scheduler latency profile
- net：网络阻断配置文件
  sync：同步阻断配置文件
  syscall：系统调用阻塞配置文件
  sched：调度器延迟配置文件

Then, you can use the pprof tool to analyze the profile:

然后，你可以使用pprof工具来分析配置文件：

```
go tool pprof TYPE.pprof
```

Note that while the various profiles available when launching 'go tool trace' work on every browser, the trace viewer itself (the 'view trace' page) comes from the Chrome/Chromium project and is only actively tested on that browser.

请注意，虽然启动 "go tool trace "时可用的各种配置文件在每个浏览器上都能工作，但跟踪查看器本身（"查看跟踪 "页面）来自Chrome/Chromium项目，并且只在该浏览器上积极测试。

