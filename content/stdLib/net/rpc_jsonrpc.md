+++
title = "rpc/jsonrpc"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/rpc/jsonrpc@go1.23.0](https://pkg.go.dev/net/rpc/jsonrpc@go1.23.0)

Package jsonrpc implements a JSON-RPC 1.0 ClientCodec and ServerCodec for the rpc package. For JSON-RPC 2.0 support, see https://godoc.org/?q=json-rpc+2.0

​	Package jsonrpc 实现了一个 JSON-RPC 1.0 ClientCodec 和 ServerCodec，用于 rpc 包。有关 JSON-RPC 2.0 支持，请参阅 [https://godoc.org/?q=json-rpc+2.0](https://godoc.org/?q=json-rpc+2.0)

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Dial 

``` go 
func Dial(network, address string) (*rpc.Client, error)
```

Dial connects to a JSON-RPC server at the specified network address.

​	Dial 连接到指定网络地址的 JSON-RPC 服务器。

### func NewClient

```go
func NewClient(conn io.ReadWriteCloser) *rpc.Client
```

NewClient returns a new rpc.Client to handle requests to the set of services at the other end of the connection.

​	NewClient 返回一个新的 rpc.Client 来处理连接另一端的服务集的请求。

### func NewClientCodec

```go
func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec
```

NewClientCodec returns a new rpc.ClientCodec using JSON-RPC on conn.

​	NewClientCodec 使用 conn 上的 JSON-RPC 返回一个新的 rpc.ClientCodec。

### func NewServerCodec

```go
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec
```

NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.

​	NewServerCodec 使用 conn 上的 JSON-RPC 返回一个新的 rpc.ServerCodec。

### func ServeConn

```go
func ServeConn(conn io.ReadWriteCloser)
```

ServeConn runs the JSON-RPC server on a single connection. ServeConn blocks, serving the connection until the client hangs up. The caller typically invokes ServeConn in a go statement.

​	ServeConn 在单个连接上运行 JSON-RPC 服务器。ServeConn 会阻塞，为连接提供服务，直到客户端挂断。调用者通常在 go 语句中调用 ServeConn。

## 类型

This section is empty.