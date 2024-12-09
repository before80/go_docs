+++
title = "JSON-RPC 接口"
date = 2024-12-09T07:56:17+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/api/json-rpc/README.md](https://github.com/go-delve/delve/blob/master/Documentation/api/json-rpc/README.md)
>
> 收录该文档时间： `2024-12-09T07:56:17+08:00`

# JSON-RPC interface - JSON-RPC 接口



Delve exposes a [JSON-RPC](https://www.jsonrpc.org/specification_v1) API interface.

​	Delve 提供了一个 [JSON-RPC](https://www.jsonrpc.org/specification_v1) API 接口。

Note that this JSON-RPC interface is served over a streaming socket, *not* over HTTP.

​	请注意，此 JSON-RPC 接口是通过流式套接字提供的，*而不是*通过 HTTP。

# API versions



Delve currently supports two versions of its API. By default a headless instance of `dlv` will serve APIv1 for backward compatibility with old clients, however new clients should use APIv2 as new features will only be made available through version 2. To select APIv2 use `--api-version=2` command line argument. Clients can also select APIv2 by sending a [SetApiVersion](https://pkg.go.dev/github.com/go-delve/delve/service/rpccommon#RPCServer.SetApiVersion) request specifying `APIVersion = 2` after connecting to the headless instance.

​	Delve 当前支持两个版本的 API。默认情况下，`dlv` 的无头实例将提供 APIv1，以保持与旧客户端的兼容性，但新客户端应使用 APIv2，因为新特性仅会通过版本 2 提供。要选择 APIv2，请使用 `--api-version=2` 命令行参数。客户端还可以通过在连接到无头实例后发送一个 [SetApiVersion](https://pkg.go.dev/github.com/go-delve/delve/service/rpccommon#RPCServer.SetApiVersion) 请求，指定 `APIVersion = 2` 来选择 APIv2。

# API version 2 documentation



All the methods of the type `service/rpc2.RPCServer` can be called using JSON-RPC, the documentation for these calls is [available on godoc](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer).

​	类型为 `service/rpc2.RPCServer` 的所有方法都可以通过 JSON-RPC 调用，这些调用的文档可以在 [godoc 上查看](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer)。

Note that all exposed methods take one single input parameter (usually called `args`) of a struct type and also return a result of a struct type. Also note that the method name should be prefixed with `RPCServer.` in JSON-RPC.

​	请注意，所有公开的方法都接受一个结构体类型的单一输入参数（通常称为 `args`），并且返回一个结构体类型的结果。同样需要注意的是，方法名在 JSON-RPC 中应以 `RPCServer.` 为前缀。

# Example



Your client wants to set a breakpoint on the function `main.main`. The first step will be calling the method `FindLocation` with `Scope = api.EvalScope{ GoroutineID: -1, Frame: 0}` and `Loc = "main.main"`. The JSON-RPC request packet should look like this:

​	假设您的客户端想要在函数 `main.main` 上设置一个断点。第一步将是调用方法 `FindLocation`，并传入 `Scope = api.EvalScope{ GoroutineID: -1, Frame: 0}` 和 `Loc = "main.main"`。JSON-RPC 请求数据包应如下所示：

```
{"method":"RPCServer.FindLocation","params":[{"Scope":{"GoroutineID":-1,"Frame":0},"Loc":"main.main"}],"id":2}
```



the response packet will look like this:

​	响应数据包将如下所示：

```
{"id":2,"result":{"Locations":[{"pc":4199019,"file":"/home/a/temp/callme/callme.go","line":31,"function":{"name":"main.main","value":4198992,"type":84,"goType":0}}]},"error":null}
```



Now your client should call the method `CreateBreakpoint` and specify `4199019` (the `pc` field in the response object) as the target address:

​	现在，您的客户端应调用方法 `CreateBreakpoint` 并指定 `4199019`（响应对象中的 `pc` 字段）作为目标地址：

```
{"method":"RPCServer.CreateBreakpoint","params":[{"Breakpoint":{"addr":4199019}}],"id":3}
```



if this request is successful your client will receive the following response:

​	如果该请求成功，您的客户端将收到以下响应：

```
{"id":3,"result":{"Breakpoint":{"id":1,"name":"","addr":4199019,"file":"/home/a/temp/callme/callme.go","line":31,"functionName":"main.main","Cond":"","continue":false,"goroutine":false,"stacktrace":0,"LoadArgs":null,"LoadLocals":null,"hitCount":{},"totalHitCount":0}},"error":null}
```
