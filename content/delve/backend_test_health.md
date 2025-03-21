+++
title = "后端测试健康状况"
date = 2024-12-09T08:09:11+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/backend_test_health.md](https://github.com/go-delve/delve/blob/master/Documentation/backend_test_health.md)
>
> 收录该文档时间： `2024-12-09T08:09:11+08:00`

Tests skipped by each supported backend:

​	每个支持的后端跳过的测试：

- 386 skipped = 9
  - 1 broken
  - 3 broken - cgo stacktraces  
  - 4 not implemented
  - 1 not working due to optimizations
- arm64 skipped = 1
  - 1 broken - global variable symbolication
- darwin skipped = 3
  - 2 follow exec not implemented on macOS
  - 1 waitfor implementation is delegated to debugserver
- darwin/arm64 skipped = 1
  - 1 broken - cgo stacktraces
- darwin/lldb skipped = 1
  - 1 upstream issue
- freebsd skipped = 11
  - 2 flaky
  - 2 follow exec not implemented on freebsd
  - 5 not implemented
  - 2 not working on freebsd
- linux/386 skipped = 2
  - 2 not working on linux/386
- linux/386/pie skipped = 1
  - 1 broken
- linux/ppc64le skipped = 3
  - 1 broken - cgo stacktraces
  - 2 not working on linux/ppc64le when -gcflags=-N -l is passed
- linux/ppc64le/native skipped = 1
  - 1 broken in linux ppc64le
- linux/ppc64le/native/pie skipped = 3
  - 3 broken - pie mode
- linux/riscv64 skipped = 2
  - 1 broken - cgo stacktraces
  - 1 not working on linux/riscv64
- pie skipped = 2
  - 2 upstream issue - [golang/go#29322](https://github.com/golang/go/issues/29322)
- ppc64le skipped = 12
  - 6 broken
  - 1 broken - global variable symbolication
  - 5 not implemented
- riscv64 skipped = 6
  - 2 broken
  - 1 broken - global variable symbolication
  - 3 not implemented
- windows skipped = 7
  - 1 broken
  - 2 not working on windows
  - 4 see [#2768](https://github.com/go-delve/delve/issues/2768)
- windows/arm64 skipped = 5
  - 3 broken
  - 1 broken - cgo stacktraces
  - 1 broken - step concurrent
