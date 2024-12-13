+++
title = "dlv_log"
date = 2024-12-09T08:06:01+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_log.md](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_log.md)
>
> 收录该文档时间： `2024-12-09T08:06:01+08:00`

## dlv log



Help about logging flags.

​	关于日志记录标志的帮助信息。

### Synopsis



Logging can be enabled by specifying the --log flag and using the --log-output flag to select which components should produce logs.

​	可以通过指定 `--log` 标志并使用 `--log-output` 标志选择哪些组件应生成日志来启用日志记录。

The argument of --log-output must be a comma separated list of component names selected from this list:

​	`--log-output` 的参数必须是一个以逗号分隔的组件名称列表，选自以下列表：

```
debugger	Log debugger commands
				 Log 调试器命令
				 
gdbwire		Log connection to gdbserial backend
				Log 与 gdbserial 后端的连接
				
lldbout		Copy output from debugserver/lldb to standard output
				将 debugserver/lldb 的输出复制到标准输出
				
debuglineerr	Log recoverable errors reading .debug_line
					Log 读取 .debug_line 文件时的可恢复错误
					
rpc		Log all RPC messages
			Log 所有 RPC 消息
			
dap		Log all DAP messages
			Log 所有 DAP 消息
			
fncall		Log function call protocol
				Log 函数调用协议
				
minidump	Log minidump loading
				Log minidump 加载
				
stack           Log stacktracer
					Log 栈追踪
```



Additionally --log-dest can be used to specify where the logs should be written. If the argument is a number it will be interpreted as a file descriptor, otherwise as a file path. This option will also redirect the "server listening at" message in headless and dap modes.

​	此外，可以使用 `--log-dest` 来指定日志应写入的位置。如果参数是一个数字，则将其解释为文件描述符；如果是文件路径，则将其作为文件路径处理。此选项还将重定向无头模式和 DAP 模式中的“服务器监听地址”消息。

### Options



```
  -h, --help   help for log
  					显示 log 命令的帮助信息
```



### 从父命令继承的选项 Options inherited from parent commands



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
