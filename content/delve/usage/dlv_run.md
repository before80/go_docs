+++
title = "dlv_run"
date = 2024-12-09T08:06:27+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_run.md](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_run.md)
>
> 收录该文档时间： `2024-12-09T08:06:27+08:00`

## dlv run



Deprecated command. Use 'debug' instead.

​	已弃用的命令。请改用 `debug` 命令。

```
dlv run [flags]
```



### Options



```
  -h, --help   help for run
```



### Options inherited from parent commands



```
      --accept-multiclient               Allows a headless server to accept multiple client connections via JSON-RPC or DAP.
      									允许无头服务器通过 JSON-RPC 或 DAP 接受多个客户端连接。
      
      --allow-non-terminal-interactive   Allows interactive sessions of Delve that don't have a terminal as stdin, stdout and stderr
      									允许没有终端作为 stdin、stdout 和 stderr 的交互式会话。
      									
      --api-version int                  Selects JSON-RPC API version when headless. New clients should use v2. Can be reset via RPCServer.SetApiVersion. See Documentation/api/json-rpc/README.md. (default 1)
      									选择无头模式下的 JSON-RPC API 版本。新的客户端应使用 v2。可以通过 RPCServer.SetApiVersion 重置。请参阅文档 Documentation/api/json-rpc/README.md。（默认值为 1）
      									
      --backend string                   Backend selection (see 'dlv help backend'). (default "default")
      									后端选择（请参阅 'dlv help backend'）。默认值为 "default"
      									
      --build-flags string               Build flags, to be passed to the compiler. For example: --build-flags="-tags=integration -mod=vendor -cover -v"
      									要传递给编译器的构建标志。例如：--build-flags="-tags=integration -mod=vendor -cover -v"
      									
      --check-go-version                 Exits if the version of Go in use is not compatible (too old or too new) with the version of Delve. (default true)
      									如果正在使用的 Go 版本与 Delve 的版本不兼容（过旧或过新），则退出。（默认值为 true）
      									
      --disable-aslr                     Disables address space randomization
      									禁用地址空间随机化
      									
      --headless                         Run debug server only, in headless mode. Server will accept both JSON-RPC or DAP client connections.
      									仅运行调试服务器，在无头模式下。服务器将接受 JSON-RPC 或 DAP 客户端连接。
      									
      --init string                      Init file, executed by the terminal client.
      									初始化文件，由终端客户端执行。
      									
  -l, --listen string                    Debugging server listen address. Prefix with 'unix:' to use a unix domain socket. (default "127.0.0.1:0")
  											调试服务器监听地址。使用 'unix:' 前缀来使用 Unix 域套接字。（默认值为 "127.0.0.1:0"）
  											
      --log                              Enable debugging server logging.
      									启用调试服务器日志。
      									
      --log-dest string                  Writes logs to the specified file or file descriptor (see 'dlv help log').
      									将日志写入指定的文件或文件描述符（参见 'dlv help log'）。
      									
      --log-output string                Comma separated list of components that should produce debug output (see 'dlv help log')
      									逗号分隔的组件列表，这些组件应生成调试输出（参见 'dlv help log'）。
      									
      --only-same-user                   Only connections from the same user that started this instance of Delve are allowed to connect. (default true)
      									仅允许与启动当前 Delve 实例的同一用户连接。（默认值为 true）
      									
  -r, --redirect stringArray             Specifies redirect rules for target process (see 'dlv help redirect')
  											指定目标进程的重定向规则（参见 'dlv help redirect'）
  											
      --wd string                        Working directory for running the program.
      									运行程序的工作目录。
```



### SEE ALSO



- [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve is a debugger for the Go programming language.
  - [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve 是 Go 编程语言的调试器。
