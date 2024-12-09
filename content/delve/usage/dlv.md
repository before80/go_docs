+++
title = "dlv"
date = 2024-12-09T08:04:45+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md)
>
> 收录该文档时间： `2024-12-09T08:04:45+08:00`

## dlv



Delve is a debugger for the Go programming language.

​	Delve 是用于 Go 编程语言的调试器。

### Synopsis



Delve is a source level debugger for Go programs.

​	Delve 是一个面向 Go 程序的源代码级调试器。

Delve enables you to interact with your program by controlling the execution of the process, evaluating variables, and providing information of thread / goroutine state, CPU register state and more.

​	Delve 允许您通过控制进程执行、评估变量以及提供线程/协程状态、CPU 寄存器状态等信息与程序进行交互。

The goal of this tool is to provide a simple yet powerful interface for debugging Go programs.

​	此工具的目标是为 Go 程序调试提供一个简单但强大的界面。

Pass flags to the program you are debugging using `--`, for example:

​	可以使用 `--` 将标志传递给您正在调试的程序，例如：

```
dlv exec ./hello -- server --config conf/config.toml
```

### Options



```
  -h, --help   help for dlv
```



### SEE ALSO



- [dlv attach](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_attach.md) - Attach to running process and begin debugging.
  - [dlv attach](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_attach.md) - 附加到正在运行的进程并开始调试。
- [dlv connect](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_connect.md) - Connect to a headless debug server with a terminal client.
  - [dlv connect](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_connect.md) - 使用终端客户端连接到无界面调试服务器。
- [dlv core](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_core.md) - Examine a core dump.
  - [dlv core](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_core.md) - 检查核心转储文件。
- [dlv dap](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_dap.md) - Starts a headless TCP server communicating via Debug Adaptor Protocol (DAP).
  - [dlv dap](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_dap.md) - 启动一个使用调试适配器协议（DAP）通信的无界面 TCP 服务器。
- [dlv debug](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_debug.md) - Compile and begin debugging main package in current directory, or the package specified.
  - [dlv debug](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_debug.md) - 编译并开始调试当前目录的主程序包或指定的程序包。
- [dlv exec](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_exec.md) - Execute a precompiled binary, and begin a debug session.
  - [dlv exec](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_exec.md) - 执行预编译的二进制文件，并开始调试会话。
- [dlv replay](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_replay.md) - Replays a rr trace.
  - [dlv replay](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_replay.md) - 重放 rr 跟踪。
- [dlv test](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_test.md) - Compile test binary and begin debugging program.
  - [dlv test](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_test.md) - 编译测试二进制文件并开始调试程序。
- [dlv trace](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_trace.md) - Compile and begin tracing program.
  - [dlv trace](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_trace.md) - 编译并开始跟踪程序。
- [dlv version](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_version.md) - Prints version.
  - [dlv version](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_version.md) - 显示版本信息。
- [dlv log](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_log.md) - Help about logging flags
  - [dlv log](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_log.md) - 关于日志标志的帮助信息。
- [dlv backend](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_backend.md) - Help about the `--backend` flag
  - [dlv backend](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_backend.md) - 关于 `--backend` 标志的帮助信息。