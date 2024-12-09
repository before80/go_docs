+++
title = "api"
date = 2024-12-09T07:54:43+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/api/README.md](https://github.com/go-delve/delve/blob/master/Documentation/api/README.md)
>
> 收录该文档时间： `2024-12-09T07:54:43+08:00`

# Server/Client API Documentation



Delve exposes two API interfaces, JSON-RPC and DAP, so that frontends other than the built-in [terminal client](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md), such as [IDEs and editors](https://github.com/go-delve/delve/blob/master/Documentation/EditorIntegration.md), can interact with Delve programmatically. The [JSON-RPC API](https://github.com/go-delve/delve/blob/master/Documentation/api/json-rpc/README.md) is used by the [terminal client](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md), and will always stay up to date in lockstep regardless of new features. The [DAP API](https://github.com/go-delve/delve/blob/master/Documentation/api/dap/README.md) is a popular generic API already in use by many [tools](https://microsoft.github.io/debug-adapter-protocol/implementors/tools/).

​	Delve 提供了两种 API 接口：JSON-RPC 和 DAP，使得除内置 [终端客户端](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md) 外的前端（如 [IDE 和编辑器](https://github.com/go-delve/delve/blob/master/Documentation/EditorIntegration.md)）可以通过编程与 Delve 进行交互。 [JSON-RPC API](https://github.com/go-delve/delve/blob/master/Documentation/api/json-rpc/README.md) 被 [终端客户端](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md) 使用，并且无论新功能如何，它将始终保持同步更新。 [DAP API](https://github.com/go-delve/delve/blob/master/Documentation/api/dap/README.md) 是一个流行的通用 API，许多 [工具](https://microsoft.github.io/debug-adapter-protocol/implementors/tools/) 已经在使用它。

## Usage



In order to run Delve in "API mode", simply invoke with one of the standard commands, providing the `--headless` flag, like so:

​	要以 "API 模式" 运行 Delve，只需使用其中一个标准命令，并提供 `--headless` 标志，如下所示：

```
$ dlv debug --headless --api-version=2 --log --log-output=debugger,dap,rpc --listen=127.0.0.1:8181
```



This will start the debugger in a non-interactive mode, listening on the specified address, and will enable logging. The logging flags as well as the server address are optional, of course.

​	这将以非交互模式启动调试器，监听指定地址，并启用日志记录。日志标志和服务器地址是可选的。

Optionally, you may also specify the `--accept-multiclient` flag if you would like to connect multiple JSON-RPC or DAP clients to the API.

​	另外，您还可以指定 `--accept-multiclient` 标志，如果希望允许多个 JSON-RPC 或 DAP 客户端连接到 API。

You can connect to the headless debugger from Delve itself using the `connect` subcommand:

​	您可以使用 `connect` 子命令从 Delve 本身连接到无头调试器：

```
$ dlv connect 127.0.0.1:8181
```



This can be useful for remote debugging.

​	这对于远程调试非常有用。

## API Interfaces



Delve has been architected in such a way as to allow multiple client/server implementations. All of the "business logic" as it were is abstracted away from the actual client/server implementations, allowing for easy implementation of new API interfaces.

​	Delve 的架构使得可以支持多个客户端/服务器实现。所有的 "业务逻辑" 都被抽象化，独立于实际的客户端/服务器实现，从而方便了新 API 接口的实现。

### 当前的 API 接口 Current API Interfaces



- [JSON-RPC](https://github.com/go-delve/delve/blob/master/Documentation/api/json-rpc/README.md)
- [DAP](https://github.com/go-delve/delve/blob/master/Documentation/api/dap/README.md)
