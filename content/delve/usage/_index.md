+++
title = "用法"
date = 2024-12-09T08:04:37+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/usage/README.md](https://github.com/go-delve/delve/blob/master/Documentation/usage/README.md)
>
> 收录该文档时间： `2024-12-09T08:04:37+08:00`

# Using Delve

You can invoke Delve in [multiple ways](dlv.md), depending on your usage needs. Delve makes every attempt to be user-friendly, ensuring the user has to do the least amount of work possible to begin debugging their program.

​	您可以根据需要通过[多种方式](dlv.md)调用 Delve。Delve 力求为用户提供便利，确保用户以最少的操作即可开始调试程序。

The [available commands](dlv.md) can be grouped into the following categories:

​	[可用命令](dlv.md)可分为以下几类：

*  Specify target and start debugging with the default [terminal interface](../cli/README.md): 指定目标并通过默认的[终端界面](../cli/README.md)开始调试：
   * [dlv debug [package]](dlv_debug.md)
   * [dlv test [package]](dlv_test.md)
   * [dlv exec \<exe\>](dlv_exec.md)
   * [dlv attach \<pid\>](dlv_attach.md)
   * [dlv core \<exe\> \<core\>](dlv_core.md)
   * [dlv replay \<rr trace\> ](dlv_replay.md)
* Trace target program execution 跟踪目标程序执行：
   * [dlv trace [package] \<regexp\>](dlv_trace.md)
* Start a headless backend server only and connect with an external [frontend client](../EditorIntegration.md): 仅启动无界面后端服务器并通过外部[前端客户端](../EditorIntegration.md)连接：
   * [dlv **--headless** \<command\> \<target\> \<args\> ](../api/ClientHowto.md#spawning-the-backend)
      * starts a server, enters a debug session for the specified target and waits to accept a client connection over JSON-RPC or DAP 启动服务器，进入指定目标的调试会话，并通过 JSON-RPC 或 DAP 等待客户端连接
      * `<command>` can be any of `debug`, `test`, `exec`, `attach`, `core` or `replay`
      * if `--headless` flag is not specified the default [terminal client](../cli/README.md) will be automatically started instead 如果未指定 `--headless` 标志，将自动启动默认的[终端客户端](../cli/README.md)
      * compatible with [dlv connect](dlv_connect.md), [VS Code Go](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#remote-debugging), [GoLand](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#attach-to-a-process-on-a-remote-machine) 兼容 [dlv connect](dlv_connect.md)、[VS Code Go](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#remote-debugging)、[GoLand](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#attach-to-a-process-on-a-remote-machine)
   * [dlv dap](dlv_dap.md)
      * starts a DAP-only server and waits for a DAP client connection to specify the target and arguments 启动仅支持 DAP 的服务器，并等待 DAP 客户端连接以指定目标和参数
      * compatible with [VS Code Go](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#remote-debugging) 兼容 [VS Code Go](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#remote-debugging)
      * NOT compatible with [dlv connect](dlv_connect.md), [GoLand](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#attach-to-a-process-on-a-remote-machine) 不兼容 [dlv connect](dlv_connect.md)、[GoLand](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#attach-to-a-process-on-a-remote-machine)
   * [dlv connect \<addr\>](dlv_connect.md)
      * starts a [terminal interface client](../cli/README.md) and connects it to a running headless server over JSON-RPC 启动[终端界面客户端](../cli/README.md)并通过 JSON-RPC 连接到运行中的无界面服务器
* Help information 帮助信息：
   * [dlv help [command]](dlv.md)
   * [dlv log](dlv_log.md)
   * [dlv backend](dlv_backend.md)
   * [dlv redirect](dlv_redirect.md)
   * [dlv version](dlv_version.md)

The above list may be incomplete. Refer to the auto-generated [complete usage document](dlv.md) to further explore all available commands.

​	上述列表可能不完整。请参阅自动生成的[完整使用文档](dlv.md)以进一步探索所有可用命令。

## 环境变量 Environment variables

Delve also reads the following environment variables:

​	Delve 还会读取以下环境变量：

* `$DELVE_EDITOR` is used by the `edit` command (if it isn't set the `$EDITOR` variable is used instead) 
  * `$DELVE_EDITOR` 由 `edit` 命令使用（如果未设置，则使用 `$EDITOR` 变量）
* `$DELVE_PAGER` is used by commands that emit large output (if it isn't set the `$PAGER` variable is used instead, if neither is set `more` is used)
  * `$DELVE_PAGER` 由输出大量内容的命令使用（如果未设置，则使用 `$PAGER` 变量，如果两者都未设置，则使用 `more`）
* `$TERM` is used to decide whether or not ANSI escape codes should be used for colorized output
  * `$TERM` 用于决定是否使用 ANSI 转义码来输出带颜色的内容
* `$DELVE_DEBUGSERVER_PATH` is used to locate the debugserver executable on macOS
  * `$DELVE_DEBUGSERVER_PATH` 用于在 macOS 上定位 debugserver 可执行文件