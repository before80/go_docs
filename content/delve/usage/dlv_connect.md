+++
title = "dlv_connect"
date = 2024-12-09T08:05:19+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_connect.md](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_connect.md)
>
> 收录该文档时间： `2024-12-09T08:05:19+08:00`

## dlv connect



Connect to a headless debug server with a terminal client.

​	使用终端客户端连接到无界面调试服务器。

### Synopsis



Connect to a running headless debug server with a terminal client. Prefix with 'unix:' to use a unix domain socket.

​	使用终端客户端连接到正在运行的无界面调试服务器。使用 'unix:' 前缀来使用 Unix 域套接字。

```
dlv connect addr [flags]
```



### Options



```
  -h, --help   help for connect
  					显示 connect 的帮助信息。
```



### 从父命令继承的选项 Options inherited from parent commands



```
      --backend string      Backend selection (see 'dlv help backend'). (default "default")
      							后端选择（参见 'dlv help backend'）。（默认 "default"）
      							
      --init string         Init file, executed by the terminal client.
      							初始化文件，由终端客户端执行。
      							
      --log                 Enable debugging server logging.
      							启用调试服务器日志。
      							
      --log-dest string     Writes logs to the specified file or file descriptor (see 'dlv help log').
      							将日志写入指定的文件或文件描述符（参见 'dlv help log'）。
      							
      --log-output string   Comma separated list of components that should produce debug output (see 'dlv help log')
      							逗号分隔的组件列表，这些组件应生成调试输出（参见 'dlv help log'）。
```



### SEE ALSO



- [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve is a debugger for the Go programming language.
  - [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve 是用于 Go 编程语言的调试器。
