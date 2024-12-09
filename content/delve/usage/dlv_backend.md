+++
title = "dlv_backend"
date = 2024-12-09T08:05:05+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_backend.md](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_backend.md)
>
> 收录该文档时间： `2024-12-09T08:05:05+08:00`

## dlv backend



Help about the --backend flag.

​	关于 `--backend` 标志的帮助信息。

### Synopsis



The --backend flag specifies which backend should be used, possible values are:

​	`--backend` 标志指定要使用的后端，可能的值为：

```
default		Uses lldb on macOS, native everywhere else.
				在 macOS 上使用 lldb，在其他平台使用 native。
				
native		Native backend.
				使用原生后端。
				
lldb		Uses lldb-server or debugserver.
				使用 lldb-server 或 debugserver。
				
rr		Uses mozilla rr (https://github.com/mozilla/rr).
			使用 mozilla rr (https://github.com/mozilla/rr)。
```



Some backends can be configured using environment variables:

​	某些后端可以通过环境变量进行配置：

- DELVE_DEBUGSERVER_PATH specifies the path of the debugserver executable for the lldb backend
  - `DELVE_DEBUGSERVER_PATH` 指定 lldb 后端的 debugserver 可执行文件路径。

- DELVE_RR_RECORD_FLAGS specifies additional flags used when calling 'rr record'
  - `DELVE_RR_RECORD_FLAGS` 指定调用 `rr record` 时使用的附加标志。

- DELVE_RR_REPLAY_FLAGS specifies additional flags used when calling 'rr replay'
  - `DELVE_RR_REPLAY_FLAGS` 指定调用 `rr replay` 时使用的附加标志。


### Options



```
  -h, --help   help for backend
  				显示 backend 的帮助信息。
```



### 从父命令继承的选项 Options inherited from parent commands



```
      --accept-multiclient               Allows a headless server to accept multiple client connections via JSON-RPC or DAP.
      										允许无界面服务器通过 JSON-RPC 或 DAP 接受多个客户端连接。
      										
      --allow-non-terminal-interactive   Allows interactive sessions of Delve that don't have a terminal as stdin, stdout and stderr
      										允许 Delve 的交互会话在没有终端作为 stdin、stdout 和 stderr 的情况下运行。
      										
      --api-version int                  Selects JSON-RPC API version when headless. New clients should use v2. Can be reset via RPCServer.SetApiVersion. See Documentation/api/json-rpc/README.md. (default 1)
      										在无界面模式下选择 JSON-RPC API 版本。新客户端应使用 v2，可通过 RPCServer.SetApiVersion 重置。详见 Documentation/api/json-rpc/README.md。（默认 1）
      										
      --backend string                   Backend selection (see 'dlv help backend'). (default "default")
      										后端选择（参见 'dlv help backend'）。（默认 "default"）
      										
      --build-flags string               Build flags, to be passed to the compiler. For example: --build-flags="-tags=integration -mod=vendor -cover -v"
      										编译标志，将传递给编译器。例如：--build-flags="-tags=integration -mod=vendor -cover -v"
      										
      --check-go-version                 Exits if the version of Go in use is not compatible (too old or too new) with the version of Delve. (default true)
      										如果所用的 Go 版本与 Delve 版本不兼容（过旧或过新）则退出。（默认 true）
      
      --disable-aslr                     Disables address space randomization
      										禁用地址空间随机化。
      										
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
      										
  -r, --redirect stringArray             Specifies redirect rules for target process (see 'dlv help redirect')
  											指定目标进程的重定向规则（参见 'dlv help redirect'）。
  											
      --wd string                        Working directory for running the program.
      										运行程序的工作目录。
```



### SEE ALSO



- [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve is a debugger for the Go programming language.
  - [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve 是用于 Go 编程语言的调试器。
