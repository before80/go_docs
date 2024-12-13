+++
title = "dlv_attach"
date = 2024-12-09T08:04:55+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_attach.md](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_attach.md)
>
> 收录该文档时间： `2024-12-09T08:04:55+08:00`

## dlv attach



Attach to running process and begin debugging.

​	附加到正在运行的进程并开始调试。

### Synopsis



Attach to an already running process and begin debugging it.

​	附加到已运行的进程并开始调试。

This command will cause Delve to take control of an already running process, and begin a new debug session. When exiting the debug session you will have the option to let the process continue or kill it.

​	此命令会让 Delve 控制一个已运行的进程，并开始一个新的调试会话。在退出调试会话时，您可以选择让进程继续运行或终止它。

```
dlv attach pid [executable] [flags]
```



### Options



```
      --continue                 Continue the debugged process on start.
      							在启动时继续调试的进程。
      							
  -h, --help                     help for attach
  									显示 attach 的帮助信息。
  									
      --waitfor string           Wait for a process with a name beginning with this prefix
      							等待一个进程，其名称以此前缀开头。
      							
      --waitfor-duration float   Total time to wait for a process
      							等待进程的总时间。
      							
      --waitfor-interval float   Interval between checks of the process list, in millisecond (default 1)
      							检查进程列表的时间间隔，单位为毫秒（默认 1）。
```



### Options inherited from parent commands



```
      --accept-multiclient               Allows a headless server to accept multiple client connections via JSON-RPC or DAP.
      									允许无界面服务器通过 JSON-RPC 或 DAP 接受多个客户端连接。
      									
      --allow-non-terminal-interactive   Allows interactive sessions of Delve that don't have a terminal as stdin, stdout and stderr
      									允许 Delve 的交互会话在没有终端作为 stdin、stdout 和 stderr 的情况下运行。
      									
      --api-version int                  Selects JSON-RPC API version when headless. New clients should use v2. Can be reset via RPCServer.SetApiVersion. See Documentation/api/json-rpc/README.md. (default 1)
      									在无界面模式下选择 JSON-RPC API 版本。新客户端应使用 v2，可通过 RPCServer.SetApiVersion 重置。详见 Documentation/api/json-rpc/README.md。（默认 1）
      									
      --backend string                   Backend selection (see 'dlv help backend'). (default "default")
      									后端选择（参见 'dlv help backend'）。（默认 "default"）
      									
      --check-go-version                 Exits if the version of Go in use is not compatible (too old or too new) with the version of Delve. (default true)
      									如果所用的 Go 版本与 Delve 版本不兼容（过旧或过新）则退出。（默认 true）
      
      --headless                         Run debug server only, in headless mode. Server will accept both JSON-RPC or DAP client connections.
      									仅运行调试服务器，无界面模式。服务器将接受 JSON-RPC 或 DAP 客户端连接。
      									
      --init string                      Init file, executed by the terminal client.
      									初始化文件，由终端客户端执行。
      									
  -l, --listen string                    Debugging server listen address. Prefix with 'unix:' to use a unix domain socket. (default "127.0.0.1:0")
  											调试服务器的监听地址。使用 'unix:' 作为前缀以使用 Unix 域套接字。（默认 "127.0.0.1:0"）
  											
      --log                              Enable debugging server logging.
      									启用调试服务器日志。
      									
      --log-dest string                  Writes logs to the specified file or file descriptor (see 'dlv help log').
      									将日志写入指定的文件或文件描述符（参见 'dlv help log'）。
      									
      --log-output string                Comma separated list of components that should produce debug output (see 'dlv help log')
      									逗号分隔的组件列表，这些组件应生成调试输出（参见 'dlv help log'）。
      									
      --only-same-user                   Only connections from the same user that started this instance of Delve are allowed to connect. (default true)
      									仅允许与启动此 Delve 实例的用户相同的用户进行连接。（默认 true）
```



### SEE ALSO



- [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve is a debugger for the Go programming language.
  - [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve 是用于 Go 编程语言的调试器。
