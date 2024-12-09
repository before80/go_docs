+++
title = "dlv_dap"
date = 2024-12-09T08:05:37+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_dap.md](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_dap.md)
>
> 收录该文档时间： `2024-12-09T08:05:37+08:00`

## dlv dap



Starts a headless TCP server communicating via Debug Adaptor Protocol (DAP).

​	启动一个通过调试适配器协议（DAP）通信的无界面 TCP 服务器。

### Synopsis



Starts a headless TCP server communicating via Debug Adaptor Protocol (DAP).

​	启动一个通过调试适配器协议（DAP）通信的无界面 TCP 服务器。

The server is always headless and requires a DAP client like VS Code to connect and request a binary to be launched or a process to be attached to. The following modes can be specified via the client's launch config:

​	该服务器始终是无界面的，需要一个 DAP 客户端（如 VS Code）连接并请求启动一个二进制文件或附加到一个进程。客户端的启动配置可以指定以下模式：

- launch + exec (executes precompiled binary, like 'dlv exec')
  - launch + exec（执行预编译的二进制文件，如 `dlv exec`）

- launch + debug (builds and launches, like 'dlv debug')
  - launch + debug（构建并启动程序，如 `dlv debug`）

- launch + test (builds and tests, like 'dlv test')
  - launch + test（构建并测试程序，如 `dlv test`）

- launch + replay (replays an rr trace, like 'dlv replay')
  - launch + replay（重播 rr 跟踪，如 `dlv replay`）

- launch + core (replays a core dump file, like 'dlv core')
  - launch + core（重播核心转储文件，如 `dlv core`）

- attach + local (attaches to a running process, like 'dlv attach')
  - attach + local（附加到正在运行的进程，如 `dlv attach`）


Program and output binary paths will be interpreted relative to dlv's working directory.

​	程序和输出二进制文件路径将相对于 dlv 的工作目录解释。

This server does not accept multiple client connections (--accept-multiclient). Use 'dlv [command] --headless' instead and a DAP client with attach + remote config. While --continue is not supported, stopOnEntry launch/attach attribute can be used to control if execution is resumed at the start of the debug session.

​	此服务器不接受多个客户端连接（`--accept-multiclient`）。使用 `dlv [command] --headless` 代替，并使用带有 attach + remote 配置的 DAP 客户端。虽然不支持 `--continue`，但可以使用 `stopOnEntry` 启动/附加属性来控制是否在调试会话开始时恢复执行。

The --client-addr flag is a special flag that makes the server initiate a debug session by dialing in to the host:port where a DAP client is waiting. This server process will exit when the debug session ends.

​	`--client-addr` 标志是一个特殊标志，它使服务器通过拨入等待 DAP 客户端的 `host:port` 来启动调试会话。调试会话结束时，服务器进程将退出。

```
dlv dap [flags]
```



### Options



```
      --client-addr string   Address where the DAP client is waiting for the DAP server to dial in. Prefix with 'unix:' to use a unix domain socket.
      								DAP 客户端等待 DAP 服务器拨入的地址。使用 'unix:' 前缀来使用 Unix 域套接字。
  -h, --help                 help for dap
  									显示 dap 的帮助信息
```



### 从父命令继承的选项 Options inherited from parent commands



```
      --check-go-version    Exits if the version of Go in use is not compatible (too old or too new) with the version of Delve. (default true)
      							如果正在使用的 Go 版本与 Delve 的版本不兼容（过旧或过新），则退出。（默认值为 true）
      							
      --disable-aslr        Disables address space randomization
      							禁用地址空间随机化
      							
  -l, --listen string       Debugging server listen address. Prefix with 'unix:' to use a unix domain socket. (default "127.0.0.1:0")
  								调试服务器监听地址。使用 'unix:' 前缀来使用 Unix 域套接字。（默认值为 "127.0.0.1:0"）
  								
      --log                 Enable debugging server logging.
      							启用调试服务器日志。
      							
      --log-dest string     Writes logs to the specified file or file descriptor (see 'dlv help log').
      							将日志写入指定的文件或文件描述符（参见 'dlv help log'）。
      							
      --log-output string   Comma separated list of components that should produce debug output (see 'dlv help log')
      							逗号分隔的组件列表，这些组件应生成调试输出（参见 'dlv help log'）。
      							
      --only-same-user      Only connections from the same user that started this instance of Delve are allowed to connect. (default true)
      							仅允许与启动当前 Delve 实例的同一用户连接。（默认值为 true）
```



### SEE ALSO



- [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve is a debugger for the Go programming language.
  - [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve 是用于 Go 编程语言的调试器。
