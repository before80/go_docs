+++
title = "DAP 接口"
date = 2024-12-09T07:55:04+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/api/dap/README.md](https://github.com/go-delve/delve/blob/master/Documentation/api/dap/README.md)
>
> 收录该文档时间： `2024-12-09T07:55:04+08:00`

# DAP Interface - DAP 接口



Delve exposes a [DAP](https://microsoft.github.io/debug-adapter-protocol/overview) API interface.

​	Delve 提供了一个 [DAP](https://microsoft.github.io/debug-adapter-protocol/overview) API 接口。

This interface is served over a streaming TCP socket using `dlv` server in one of the two headless modes:

​	该接口通过流式 TCP 套接字提供，使用 `dlv` 服务器在以下两种无头模式下之一运行：

1. [`dlv dap`](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_dap.md) - starts a single-use DAP-only server that waits for a client to specify launch/attach configuration for starting the debug session. [`dlv dap`](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_dap.md) - 启动一个一次性使用的仅 DAP 服务器，等待客户端指定启动/附加配置以开始调试会话。

2. `dlv --headless <command> <debuggee>` - starts a general server, enters a debug session for the specified debuggee and waits for a `JSON-RPC` or a `DAP` remote-attach client to begin interactive debugging. Can be used in multi-client mode with the following options: `dlv --headless <command> <debuggee>` - 启动一个通用服务器，进入指定调试目标的调试会话，并等待一个 `JSON-RPC` 或 `DAP` 远程附加客户端开始交互式调试。可以通过以下选项以多客户端模式使用：
   - `--accept-multiclient` - use to support connections from multiple clients
     - `--accept-multiclient` - 用于支持来自多个客户端的连接
   - `--continue` - use to resume debuggee execution as soon as server session starts
     - `--continue` - 用于在服务器会话启动后立即恢复调试目标的执行

See [Launch and Attach Configurations](https://github.com/go-delve/delve/blob/master/Documentation/api/dap/README.md#launch-and-attach-configurations) for more usage details of these two options.

​	有关这两个选项的更多使用详情，请参见 [Launch and Attach Configurations](https://github.com/go-delve/delve/blob/master/Documentation/api/dap/README.md#launch-and-attach-configurations)。

The primary user of this mode is [VS Code Go](https://github.com/golang/vscode-go). Please see its detailed [debugging documentation](https://github.com/golang/vscode-go/blob/master/docs/debugging.md) for additional information.

​	此模式的主要用户是 [VS Code Go](https://github.com/golang/vscode-go)。有关更多信息，请参见其详细的 [调试文档](https://github.com/golang/vscode-go/blob/master/docs/debugging.md)。

## 调试适配器协议 Debug Adapter Protocol



[DAP](https://microsoft.github.io/debug-adapter-protocol/specification) is a general debugging protocol supported by many [tools](https://microsoft.github.io/debug-adapter-protocol/implementors/tools/) and [programming languages](https://microsoft.github.io/debug-adapter-protocol/implementors/adapters/). We tailored it to Go specifics, such as mapping [threads request](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Threads) to communicate goroutines and [exceptionInfo request](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_ExceptionInfo) to support panics and fatal errors.

​	[DAP](https://microsoft.github.io/debug-adapter-protocol/specification) 是一个广泛支持的通用调试协议，许多 [工具](https://microsoft.github.io/debug-adapter-protocol/implementors/tools/) 和 [编程语言](https://microsoft.github.io/debug-adapter-protocol/implementors/adapters/) 都支持它。我们根据 Go 语言的特点定制了该协议，例如将 [线程请求](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Threads) 映射到通信 goroutine，并将 [exceptionInfo 请求](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_ExceptionInfo) 用于支持 panic 和致命错误。

See [dap.Server.handleRequest](https://github.com/go-delve/delve/search?q=handleRequest) and capabilities set in [dap.Server.onInitializeRequest](https://github.com/go-delve/delve/search?q=onInitializeRequest) for an up-to-date list of supported requests and options.

​	请参见 [dap.Server.handleRequest](https://github.com/go-delve/delve/search?q=handleRequest) 和 [dap.Server.onInitializeRequest](https://github.com/go-delve/delve/search?q=onInitializeRequest) 中设置的能力，以获取支持的请求和选项的最新列表。

## 启动和附加配置 Launch and Attach Configurations



In addition to the general [DAP spec](https://microsoft.github.io/debug-adapter-protocol/specification), the server supports the following implementation-specific configuration options for starting the debug session:

​	除了通用的 [DAP 规范](https://microsoft.github.io/debug-adapter-protocol/specification)，该服务器还支持以下特定实现的配置选项，用于启动调试会话：

| request                                                      | mode 模式            | required 必需项 | optional 可选项 |         |         |      |            |            |         |         |                                                              |
| ------------------------------------------------------------ | -------------------- | --------------- | --------------- | ------- | ------- | ---- | ---------- | ---------- | ------- | ------- | ------------------------------------------------------------ |
| launch [godoc](https://pkg.go.dev/github.com/go-delve/delve/service/dap#LaunchConfig) | debug                | program         | dlvCwd          | env     | backend | args | cwd        | buildFlags | output  | noDebug | substitutePath stopOnEntry stackTraceDepth showGlobalVariables showRegisters showPprofLabels hideSystemGoroutines goroutineFilters |
| test                                                         | program              | dlvCwd          | env             | backend | args    | cwd  | buildFlags | output     | noDebug |         |                                                              |
| exec                                                         | program              | dlvCwd          | env             | backend | args    | cwd  |            |            | noDebug |         |                                                              |
| core                                                         | program corefilePath | dlvCwd          | env             |         |         |      |            |            |         |         |                                                              |
| replay                                                       | traceDirPath         | dlvCwd          | env             |         |         |      |            |            |         |         |                                                              |
| attach [godoc](https://pkg.go.dev/github.com/go-delve/delve/service/dap#AttachConfig) | local                | processId       |                 |         | backend |      |            |            |         |         |                                                              |
| remote                                                       |                      |                 |                 |         |         |      |            |            |         |         |                                                              |

Not all of the configurations are supported by each of the two available DAP servers:

​	并非所有配置都由这两个可用的 DAP 服务器支持：

| request | "mode":                               | `dlv dap`     | `dlv --headless` | Description                                                  | Typical Client Usage 典型客户端使用                          |
| ------- | ------------------------------------- | ------------- | ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| launch  | "debug" "test" "exec" "replay" "core" | supported     | NOT supported    | 告诉 `dlv dap` 服务器启动指定的目标并开始调试。 Tells the `dlv dap` server to launch the specified target and start debugging it. | 客户端将为用户启动 `dlv dap` 服务器，或者允许用户指定外部（即远程）服务器的 `host:port`。 The client would launch the `dlv dap` server for the user or allow them to specify `host:port` of an external (a.k.a. remote) server. |
| attach  | "local"                               | supported     | NOT supported    | 告诉 `dlv dap` 服务器附加到本地现有进程。 Tells the `dlv dap` server to attach to an existing process local to the server. |                                                              |
| attach  | "remote"                              | NOT supported | supported        | 告诉 `dlv --headless` 服务器，预计它已经在命令行调用中调试指定的目标。 Tells the `dlv --headless` server that it is expected to already be debugging a target specified as part of its command-line invocation. | 客户端将期望指定外部（即远程）服务器的 `host:port`，用户已经使用目标 [命令和参数](https://github.com/go-delve/delve/blob/master/Documentation/usage/README.md) 启动该服务器。 The client would expect `host:port` specification of an external (a.k.a. remote) server that the user already started with target [command and args](https://github.com/go-delve/delve/blob/master/Documentation/usage/README.md). |

## 断开连接与关闭 Disconnect and Shutdown



### 单客户端模式 Single-Client Mode



When used with `dlv dap` or `dlv --headless --accept-multiclient=false` (default), the DAP server will shut itself down at the end of the debug session, when the client sends a [disconnect request](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Disconnect). If the debuggee was launched, it will be taken down as well. If the debuggee was attached to, `terminateDebuggee` option will be respected.

​	当与 `dlv dap` 或 `dlv --headless --accept-multiclient=false`（默认设置）一起使用时，DAP 服务器将在调试会话结束时关闭，当客户端发送 [断开连接请求](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Disconnect) 时。如果调试目标已启动，它也将被关闭。如果附加到调试目标，则会尊重 `terminateDebuggee` 选项。

When the program terminates, we send a [terminated event](https://microsoft.github.io/debug-adapter-protocol/specification#Events_Terminated), which is expected to trigger a [disconnect request](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Disconnect) from the client for a session and a server shutdown. The [restart request](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Restart) is not yet supported.

​	当程序终止时，我们会发送 [terminated event](https://microsoft.github.io/debug-adapter-protocol/specification#Events_Terminated)，预计会触发客户端的 [断开连接请求](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Disconnect) 以关闭会话并关闭服务器。暂不支持 [重启请求](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Restart)。

The server also shuts down in case of a client connection error or SIGTERM signal, taking down a launched process, but letting an attached process continue.

​	如果发生客户端连接错误或 SIGTERM 信号，服务器也会关闭，终止启动的进程，但允许附加的进程继续运行。

Pressing Ctrl-C on the terminal where a headless server is running sends SIGINT to the debuggee, foregrounded in headless mode to support debugging interactive programs.

​	在运行无头服务器的终端按下 Ctrl-C 会向调试目标发送 SIGINT，前台支持交互式程序调试。

### 多客户端模式 Multi-Client Mode



When used with `dlv --headless --accept-multiclient=true`, the DAP server will honor the multi-client mode when a client [disconnects](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Disconnect) or client connection fails. The server will remain running and ready for a new client connection, and the debuggee will remain in whatever state it was at the time of disconnect - running or halted. Once [`suspendDebuggee`](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Disconnect) option is supported by frontends like VS Code ([vscode/issues/134412](https://github.com/microsoft/vscode/issues/134412)), we will update the server to offer this as a way to specify debuggee state on disconnect.

​	当与 `dlv --headless --accept-multiclient=true` 一起使用时，DAP 服务器将在客户端 [断开连接](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Disconnect) 或客户端连接失败时保持多客户端模式。服务器将继续运行并准备接受新的客户端连接，调试目标将保持在断开连接时的状态 —— 运行或暂停。一旦前端如 VS Code ([vscode/issues/134412](https://github.com/microsoft/vscode/issues/134412)) 支持 [`suspendDebuggee`](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Disconnect) 选项，我们将更新服务器以提供此选项，在断开连接时指定调试目标的状态。

The client may request full shutdown of the server and the debuggee with [`terminateDebuggee`](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Disconnect) option.

​	客户端可以通过 [`terminateDebuggee`](https://microsoft.github.io/debug-adapter-protocol/specification#Requests_Disconnect) 选项请求完全关闭服务器和调试目标。

The server shuts down in response to a SIGTERM signal, taking down a launched process, but letting an attached process continue.

​	服务器会响应 SIGTERM 信号，终止启动的进程，但允许附加的进程继续运行。

Pressing Ctrl-C on the terminal where a headless server is running sends SIGINT to the debuggee, foregrounded in headless mode to support debugging interactive programs.

​	在运行无头服务器的终端按下 Ctrl-C 会向调试目标发送 SIGINT，前台支持交互式程序调试。

## 调试器输出 Debugger Output



The debugger always logs one of the following on start-up to stdout:

​	调试器在启动时总是将以下内容之一记录到标准输出（stdout）：

- `dlv dap`:
  - `DAP server listening at: <host>:<port>`
  
- `dlv --headless`:
  - `API server listening at: <host>:<port>`

This can be used to confirm that server start-up succeeded.

​	这可以用来确认服务器启动成功。

The server uses [output events](https://microsoft.github.io/debug-adapter-protocol/specification#Events_Output) to communicate errors and select status messages to the client. For example:

​	服务器使用 [输出事件](https://microsoft.github.io/debug-adapter-protocol/specification#Events_Output) 来与客户端通信错误和选择状态消息。例如：

```
Step interrupted by a breakpoint. Use 'Continue' to resume the original step command.
invalid command: Unable to step while the previous step is interrupted by a breakpoint.
Use 'Continue' to resume the original step command.
Detaching and terminating target process
```



More detailed logging can be enabled with `--log --log-output=dap` as part of the [`dlv` command](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md). It will record the server-side DAP message traffic. For example,

​	可以通过在 [`dlv` 命令](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) 中使用 `--log --log-output=dap` 启用更详细的日志记录。它将记录服务器端的 DAP 消息流量。例如：

```
2022-01-04T00:27:57-08:00 debug layer=dap [<- from client]{"seq":1,"type":"request","command":"initialize","arguments":{"clientID":"vscode","clientName":"Visual Studio Code","adapterID":"go","locale":"en-us","linesStartAt1":true,"columnsStartAt1":true,"pathFormat":"path","supportsVariableType":true,"supportsVariablePaging":true,"supportsRunInTerminalRequest":true,"supportsMemoryReferences":true,"supportsProgressReporting":true,"supportsInvalidatedEvent":true}}
2022-01-04T00:27:57-08:00 debug layer=dap [-> to client]{"seq":0,"type":"response","request_seq":1,"success":true,"command":"initialize","body":{"supportsConfigurationDoneRequest":true,"supportsFunctionBreakpoints":true,"supportsConditionalBreakpoints":true,"supportsEvaluateForHovers":true,"supportsSetVariable":true,"supportsExceptionInfoRequest":true,"supportTerminateDebuggee":true,"supportsDelayedStackTraceLoading":true,"supportsLogPoints":true,"supportsDisassembleRequest":true,"supportsClipboardContext":true,"supportsSteppingGranularity":true,"supportsInstructionBreakpoints":true}}
2022-01-04T00:27:57-08:00 debug layer=dap [<- from client]{"seq":2,"type":"request","command":"launch","arguments":{"name":"Launch file","type":"go","request":"launch","mode":"debug","program":"./temp.go","hideSystemGoroutines":true,"__buildDir":"/Users/polina/go/src","__sessionId":"2ad0f0c1-a1fd-4fff-9fff-b8bc9a933fe5"}}
2022-01-04T00:27:57-08:00 debug layer=dap parsed launch config: {
	"mode": "debug",
	"program": "./temp.go",
	"backend": "default",
	"stackTraceDepth": 50,
	"hideSystemGoroutines": true
}
...
```



This logging is written to stderr and is not forwarded via [output events](https://microsoft.github.io/debug-adapter-protocol/specification#Events_Output).

​	这些日志记录会写入标准错误（stderr），并不会通过 [输出事件](https://microsoft.github.io/debug-adapter-protocol/specification#Events_Output) 转发。

## 调试目标输出 Debuggee Output



Debuggee's stdout and stderr are written to stdout and stderr respectfully and are not forwarded via [output events](https://microsoft.github.io/debug-adapter-protocol/specification#Events_Output).

​	调试目标的标准输出（stdout）和标准错误（stderr）分别写入标准输出和标准错误，并不会通过 [输出事件](https://microsoft.github.io/debug-adapter-protocol/specification#Events_Output) 转发。

## Versions



The initial DAP support was released in [v1.6.1](https://github.com/go-delve/delve/releases/tag/v1.6.1) with many additional improvements in subsequent versions. The [remote attach](https://github.com/go-delve/delve/issues/2328) support was added in [v1.7.3](https://github.com/go-delve/delve/releases/tag/v1.7.3).

​	最初的 DAP 支持在 [v1.6.1](https://github.com/go-delve/delve/releases/tag/v1.6.1) 中发布，并在后续版本中进行了许多改进。 [远程附加](https://github.com/go-delve/delve/issues/2328) 支持是在 [v1.7.3](https://github.com/go-delve/delve/releases/tag/v1.7.3) 中添加的。

The DAP API changes are backward-compatible as all new features are opt-in only. To update to a new [DAP version](https://microsoft.github.io/debug-adapter-protocol/changelog) and import a new DAP feature into delve, one must first update the [go-dap](https://github.com/google/go-dap) dependency.

​	DAP API 的更改是向后兼容的，因为所有新特性都是选择性启用的。要更新到新的 [DAP 版本](https://microsoft.github.io/debug-adapter-protocol/changelog) 并将新的 DAP 特性导入到 delve 中，首先必须更新 [go-dap](https://github.com/google/go-dap) 依赖项。
