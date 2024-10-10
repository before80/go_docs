+++
title = "trace"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# trace

> 原文：[https://pkg.go.dev/cmd/trace@go1.19.3](https://pkg.go.dev/cmd/trace@go1.19.3)

### Overview 概述

Trace is a tool for viewing trace files.

​	trace 是一个查看跟踪文件的工具。

Trace files can be generated with:

​	跟踪文件可以通过以下方式生成：

- runtime/trace.Start

- net/http/pprof package

- go test -trace


Example usage: Generate a trace file with 'go test':

​	使用实例：用'go test'生成一个跟踪文件：

```
go test -trace trace.out pkg
```

View the trace in a web browser:

​	在网络浏览器中查看跟踪：

```
go tool trace trace.out
```

Generate a pprof-like profile from the trace:

​	从跟踪中生成一个类似于pprof的分析文件：

```
go tool trace -pprof=TYPE trace.out > TYPE.pprof
```

Supported profile types are:

​	支持的分析文件类型有：

- net: network blocking profile
- net：网络阻断分析文件
- sync: synchronization blocking profile
- sync：同步阻断分析文件
- syscall: syscall blocking profile
- syscall：系统调用阻塞分析文件
- sched: scheduler latency profile
- sched：调度器延迟分析文件
  
  

Then, you can use the pprof tool to analyze the profile:

​	然后，您可以使用pprof工具来剖析分析文件：

```
go tool pprof TYPE.pprof
```

Note that while the various profiles available when launching 'go tool trace' work on every browser, the trace viewer itself (the 'view trace' page) comes from the Chrome/Chromium project and is only actively tested on that browser.

​	请注意，虽然在启动“go tool trace”时可用的各种分析文件在所有浏览器中均有效，但跟踪查看器本身（“查看跟踪”页面）来自 Chrome/Chromium 项目，并且仅在此浏览器上经过主动测试。

