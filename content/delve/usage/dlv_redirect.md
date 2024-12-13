+++
title = "dlv_redirect"
date = 2024-12-09T08:06:10+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_redirect.md](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_redirect.md)
>
> 收录该文档时间： `2024-12-09T08:06:10+08:00`

## dlv redirect



Help about file redirection.

​	关于文件重定向的帮助信息。

### Synopsis



The standard file descriptors of the target process can be controlled using the '-r' and '--tty' arguments.

​	可以使用 `-r` 和 `--tty` 参数控制目标进程的标准文件描述符。

The --tty argument allows redirecting all standard descriptors to a terminal, specified as an argument to --tty.

​	`--tty` 参数允许将所有标准描述符重定向到指定的终端，该终端作为 `--tty` 的参数。

The syntax for '-r' argument is:

​	`-r` 参数的语法为：

```
	-r [source:]destination
```



Where source is one of 'stdin', 'stdout' or 'stderr' and destination is the path to a file. If the source is omitted stdin is used implicitly.

​	其中，`source` 可以是 'stdin'、'stdout' 或 'stderr'，而 `destination` 是文件的路径。如果省略了 `source`，则默认使用 `stdin`。

File redirects can also be changed using the 'restart' command.

​	文件重定向也可以通过 `restart` 命令来更改。

### Options



```
  -h, --help   help for redirect
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
      									后端选择（参见 'dlv help backend'）。 （默认值为 "default"）
      									
      --build-flags string               Build flags, to be passed to the compiler. For example: --build-flags="-tags=integration -mod=vendor -cover -v"
      									编译器的构建标志。例如：--build-flags="-tags=integration -mod=vendor -cover -v"
      									
      --check-go-version                 Exits if the version of Go in use is not compatible (too old or too new) with the version of Delve. (default true)
      									如果正在使用的 Go 版本与 Delve 的版本不兼容（过旧或过新），则退出。（默认值为 true）
      									
      --disable-aslr                     Disables address space randomization
      									禁用地址空间随机化
      									
      --headless                         Run debug server only, in headless mode. Server will accept both JSON-RPC or DAP client connections.
      									仅运行调试服务器，在无头模式下。服务器将接受 JSON-RPC 或 DAP 客户端连接。
      									
      --init string                      Init file, executed by the terminal client.
      									执行的初始化文件，由终端客户端执行。
      									
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
  											指定目标进程的重定向规则（参见 'dlv help redirect'）。
  											
      --wd string                        Working directory for running the program.
      									运行程序的工作目录。
```



### SEE ALSO



- [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve is a debugger for the Go programming language.
  - [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve 是 Go 编程语言的调试器。
