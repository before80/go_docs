+++
title = "rpc"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/rpc@go1.23.0](https://pkg.go.dev/net/rpc@go1.23.0)

Package rpc provides access to the exported methods of an object across a network or other I/O connection. A server registers an object, making it visible as a service with the name of the type of the object. After registration, exported methods of the object will be accessible remotely. A server may register multiple objects (services) of different types but it is an error to register multiple objects of the same type.

​	`rpc`包提供了通过网络或其他I/O连接访问对象导出方法的方式。服务器注册一个对象，使其作为一个带有该对象类型名称的服务可见。注册后，对象的导出方法将可以远程访问。服务器可以注册多个不同类型的对象(服务)，但注册多个相同类型的对象是错误的。

Only methods that satisfy these criteria will be made available for remote access; other methods will be ignored:

​	仅满足以下条件的方法才可供远程访问，其他方法将被忽略：

- the method's type is exported.
- 方法的类型是导出的。
- the method is exported.
- 方法是导出的。
- the method has two arguments, both exported (or builtin) types.
- 方法有两个实参，均为导出(或内置)类型。 
- the method's second argument is a pointer.
- 方法的第二个实参是一个指针。 
- the method has return type error.
- 方法有返回类型错误。

In effect, the method must look schematically like

​	实际上，方法的原型应该看起来像这样：

``` go 
func (t *T) MethodName(argType T1, replyType *T2) error
```

where T1 and T2 can be marshaled by encoding/gob. These requirements apply even if a different codec is used. (In the future, these requirements may soften for custom codecs.)

其中T1和T2可以通过encoding/gob编组。即使使用不同的编解码器，这些要求也适用。(未来，这些要求可能会放宽以用于自定义编解码器。)

The method's first argument represents the arguments provided by the caller; the second argument represents the result parameters to be returned to the caller. The method's return value, if non-nil, is passed back as a string that the client sees as if created by errors.New. If an error is returned, the reply parameter will not be sent back to the client.

​	方法的第一个参数表示调用者提供的参数；第二个参数表示要返回给调用者的结果参数。如果方法的返回值非空，则作为字符串传回给客户端，客户端将看到它如同由errors.New创建。如果返回一个错误，则回复参数不会发送回客户端。

The server may handle requests on a single connection by calling ServeConn. More typically it will create a network listener and call Accept or, for an HTTP listener, HandleHTTP and http.Serve.

​	服务器可以通过调用ServeConn在单个连接上处理请求。更典型的是，它将创建一个网络监听器并调用Accept或对于HTTP监听器，则为HandleHTTP和http.Serve。

A client wishing to use the service establishes a connection and then invokes NewClient on the connection. The convenience function Dial (DialHTTP) performs both steps for a raw network connection (an HTTP connection). The resulting Client object has two methods, Call and Go, that specify the service and method to call, a pointer containing the arguments, and a pointer to receive the result parameters.

​	希望使用服务的客户端建立连接，然后在连接上调用NewClient。便捷函数Dial(DialHTTP)为原始网络连接(HTTP连接)执行两个步骤。生成的Client对象有两个方法，Call和Go，指定要调用的服务和方法，一个指向参数的指针以及一个指向接收结果参数的指针。

The Call method waits for the remote call to complete while the Go method launches the call asynchronously and signals completion using the Call structure's Done channel.

​	`Call`方法等待远程调用完成，而Go方法以异步方式启动调用，并使用Call结构体的Done通道表示完成。

Unless an explicit codec is set up, package encoding/gob is used to transport the data.

​	除非设置了显式编解码器，否则使用encoding/gob包传输数据。

Here is a simple example. A server wishes to export an object of type Arith:

​	以下是一个简单的示例。服务器希望导出Arith类型的对象：

``` go 
package server

import "errors"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
```

The server calls (for HTTP service):

​	服务器调用(用于 HTTP 服务)：

```go 
arith := new(Arith)
rpc.Register(arith)
rpc.HandleHTTP()
l, e := net.Listen("tcp", ":1234")
if e != nil {
	log.Fatal("listen error:", e)
}
go http.Serve(l, nil)
```

At this point, clients can see a service "Arith" with methods "Arith.Multiply" and "Arith.Divide". To invoke one, a client first dials the server:

​	此时，客户端可以看到一个名为"Arith"的服务，其具有"Arith.Multiply"和"Arith.Divide"方法。要调用其中一个，客户端首先要拨号到服务器：

```go 
client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
if err != nil {
	log.Fatal("dialing:", err)
}
```

Then it can make a remote call:

​	然后可以进行远程调用：

```go 
// 同步调用
args := &server.Args{7,8}
var reply int
err = client.Call("Arith.Multiply", args, &reply)
if err != nil {
	log.Fatal("arith error:", err)
}
fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
```

or

或

```go 
// 异步调用
quotient := new(Quotient)
divCall := client.Go("Arith.Divide", args, quotient, nil)
replyCall := <-divCall.Done	//  将等于 divCall
// 检查错误，打印等等
```

A server implementation will often provide a simple, type-safe wrapper for the client.

​	服务器实现通常会为客户端提供一个简单、类型安全的包装器。

The net/rpc package is frozen and is not accepting new features.

​	`net/rpc` 包已经被冻结，不再接受新功能。



## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/rpc/server.go;l=143)

``` go 
const (
	// HandleHTTP使用的默认值
	DefaultRPCPath   = "/_goRPC_"
	DefaultDebugPath = "/debug/rpc"
)
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/rpc/server.go;l=202)

``` go 
var DefaultServer = NewServer()
```

DefaultServer is the default instance of *Server.

​	DefaultServer是`*Server`的默认实例。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/rpc/client.go;l=26)

``` go 
var ErrShutdown = errors.New("connection is shut down")
```

## 函数

### func Accept 

``` go 
func Accept(lis net.Listener)
```

Accept accepts connections on the listener and serves requests to DefaultServer for each incoming connection. Accept blocks; the caller typically invokes it in a go statement.

​	`Accept`函数在侦听器上接受连接，并为每个传入连接为DefaultServer提供请求。 Accept将阻塞； 调用者通常在go语句中调用它。

### func HandleHTTP 

``` go 
func HandleHTTP()
```

HandleHTTP registers an HTTP handler for RPC messages to DefaultServer on DefaultRPCPath and a debugging handler on DefaultDebugPath. It is still necessary to invoke http.Serve(), typically in a go statement.

​	`HandleHTTP`函数在DefaultRPCPath上为RPC消息注册HTTP处理程序到DefaultServer，并在DefaultDebugPath上注册调试处理程序。 仍然需要调用http.Serve()，通常在go语句中调用。

### func Register 

``` go 
func Register(rcvr any) error
```

Register publishes the receiver's methods in the DefaultServer.

​	`Register`函数在DefaultServer中发布接收器的方法。

### func RegisterName 

``` go 
func RegisterName(name string, rcvr any) error
```

RegisterName is like Register but uses the provided name for the type instead of the receiver's concrete type.

​	`RegisterName`函数类似于Register函数，但使用提供的名称替换接收器的具体类型。

### func ServeCodec 

``` go 
func ServeCodec(codec ServerCodec)
```

ServeCodec is like ServeConn but uses the specified codec to decode requests and encode responses.

​	`ServeCodec`函数类似于ServeConn函数，但使用指定的编解码器解码请求和编码响应。

### func ServeConn 

``` go 
func ServeConn(conn io.ReadWriteCloser)
```

ServeConn runs the DefaultServer on a single connection. ServeConn blocks, serving the connection until the client hangs up. The caller typically invokes ServeConn in a go statement. ServeConn uses the gob wire format (see package gob) on the connection. To use an alternate codec, use ServeCodec. See NewClient's comment for information about concurrent access.

​	`ServeConn`函数在单个连接上运行 DefaultServer。ServeConn函数会阻塞，直到客户端挂断连接。通常调用者会在go语句中调用ServeConn。ServeConn函数在连接上使用gob wire格式(请参见package gob)。要使用替代编解码器，请使用ServeCodec函数。有关并发访问的信息，请参阅NewClient函数的注释。

### func ServeRequest 

``` go 
func ServeRequest(codec ServerCodec) error
```

ServeRequest is like ServeCodec but synchronously serves a single request. It does not close the codec upon completion.

​	`ServeRequest`函数类似于ServeCodec函数，但是同步服务一个单一请求。完成后不会关闭编解码器。

## 类型

### type Call 

``` go 
type Call struct {
	ServiceMethod string     // The name of the service and method to call.
	Args          any        // The argument to the function (*struct).
	Reply         any        // The reply from the function (*struct).
	Error         error      // After completion, the error status.
	Done          chan *Call // Receives *Call when Go is complete.
}
```

Call represents an active RPC.

​	`Call`代表一个活动的RPC。

### type Client 

``` go 
type Client struct {
	// contains filtered or unexported fields
}
```

Client represents an RPC Client. There may be multiple outstanding Calls associated with a single Client, and a Client may be used by multiple goroutines simultaneously.

​	`Client`代表一个RPC客户端。一个Client可能与多个Outstanding Calls相关联，并且一个Client可能被多个goroutine同时使用。

#### func Dial 

``` go 
func Dial(network, address string) (*Client, error)
```

Dial connects to an RPC server at the specified network address.

​	`Dial`函数连接到指定网络地址的RPC服务器。

#### func DialHTTP 

``` go 
func DialHTTP(network, address string) (*Client, error)
```

DialHTTP connects to an HTTP RPC server at the specified network address listening on the default HTTP RPC path.

​	`DialHTTP`函数连接到指定网络地址上的HTTP RPC服务器，监听默认的HTTP RPC路径。

#### func DialHTTPPath 

``` go 
func DialHTTPPath(network, address, path string) (*Client, error)
```

DialHTTPPath connects to an HTTP RPC server at the specified network address and path.

​	`DialHTTPPath`函数连接到指定网络地址和路径的HTTP RPC服务器。

#### func NewClient 

``` go 
func NewClient(conn io.ReadWriteCloser) *Client
```

NewClient returns a new Client to handle requests to the set of services at the other end of the connection. It adds a buffer to the write side of the connection so the header and payload are sent as a unit.

​	`NewClient`函数返回一个新的Client，以处理连接另一端的服务集合的请求。它在连接的写入端添加一个缓冲区，以便标头和负载作为一个单元发送。

The read and write halves of the connection are serialized independently, so no interlocking is required. However each half may be accessed concurrently so the implementation of conn should protect against concurrent reads or concurrent writes.

​	连接的读取和写入部分独立进行序列化，因此不需要交错。但是，每个部分可能会同时访问，因此conn的实现应保护并发读取或并发写入。

#### func NewClientWithCodec 

``` go 
func NewClientWithCodec(codec ClientCodec) *Client
```

NewClientWithCodec is like NewClient but uses the specified codec to encode requests and decode responses.

​	`NewClientWithCodec`函数与 `NewClient`函数类似，但使用指定的编解码器对请求进行编码和响应进行解码。

#### (*Client) Call 

``` go 
func (client *Client) Call(serviceMethod string, args any, reply any) error
```

Call invokes the named function, waits for it to complete, and returns its error status.

​	`Call`方法调用指定的函数，等待它完成，并返回其错误状态。

#### (*Client) Close 

``` go 
func (client *Client) Close() error
```

Close calls the underlying codec's Close method. If the connection is already shutting down, ErrShutdown is returned.

​	`Close`方法调用底层编解码器的 `Close` 方法。如果连接已经在关闭，则返回 `ErrShutdown`。

#### (*Client) Go 

``` go 
func (client *Client) Go(serviceMethod string, args any, reply any, done chan *Call) *Call
```

Go invokes the function asynchronously. It returns the Call structure representing the invocation. The done channel will signal when the call is complete by returning the same Call object. If done is nil, Go will allocate a new channel. If non-nil, done must be buffered or Go will deliberately crash.

​	`Go`方法异步地调用函数。它返回表示调用的 `Call` 结构体。当调用完成时，`done` 信道会通过返回相同的 `Call` 对象来发出信号。如果 `done` 是 nil，则 Go 会分配一个新的信道。如果 `done` 不是 nil，则必须是有缓冲的，否则 Go 将会故意崩溃。

### type ClientCodec 

``` go 
type ClientCodec interface {
	WriteRequest(*Request, any) error
	ReadResponseHeader(*Response) error
	ReadResponseBody(any) error

	Close() error
}
```

A ClientCodec implements writing of RPC requests and reading of RPC responses for the client side of an RPC session. The client calls WriteRequest to write a request to the connection and calls ReadResponseHeader and ReadResponseBody in pairs to read responses. The client calls Close when finished with the connection. ReadResponseBody may be called with a nil argument to force the body of the response to be read and then discarded. See NewClient's comment for information about concurrent access.

​	`ClientCodec` 实现了 RPC 会话的客户端一侧的 RPC 请求的写入和响应的读取。客户端调用 `WriteRequest` 将请求写入连接，调用 `ReadResponseHeader` 和 `ReadResponseBody` 成对地读取响应。客户端在完成连接时调用 `Close`。可以将 `ReadResponseBody` 调用为 nil 参数，以强制读取响应体并随后丢弃。有关并发访问的信息，请参见 `NewClient` 的注释。

### type Request 

``` go 
type Request struct {
	ServiceMethod string // format: "Service.Method"
	Seq           uint64 // 由客户端选择的序列号
	// contains filtered or unexported fields
}
```

Request is a header written before every RPC call. It is used internally but documented here as an aid to debugging, such as when analyzing network traffic.

​	`Request` 是每个 RPC 调用前写入的头。它在内部使用，但在此文档中作为调试辅助工具进行记录，例如在分析网络流量时使用。

### type Response 

``` go 
type Response struct {
	ServiceMethod string // 与 Request 相同
	Seq           uint64 // 与请求的序列号相同
	Error         string // 错误(如果有)。
	// contains filtered or unexported fields
}
```

Response is a header written before every RPC return. It is used internally but documented here as an aid to debugging, such as when analyzing network traffic.

​	`Response` 是每个 RPC 返回前写入的头。它在内部使用，但在此文档中作为调试辅助工具进行记录，例如在分析网络流量时使用。

### type Server 

``` go 
type Server struct {
	// contains filtered or unexported fields
}
```

Server represents an RPC Server.

​	`Server` 表示一个 RPC 服务器。

#### func NewServer 

``` go 
func NewServer() *Server
```

NewServer returns a new Server.

​	`NewServer`函数返回一个新的 `Server`。

#### (*Server) Accept 

``` go 
func (server *Server) Accept(lis net.Listener)
```

Accept accepts connections on the listener and serves requests for each incoming connection. Accept blocks until the listener returns a non-nil error. The caller typically invokes Accept in a go statement.

​	`Accept`方法在侦听器上接受连接并为每个传入的连接提供请求服务。Accept 阻塞，直到侦听器返回非 nil 错误。通常在 go 语句中调用 Accept。

#### (*Server) HandleHTTP 

``` go 
func (server *Server) HandleHTTP(rpcPath, debugPath string)
```

HandleHTTP registers an HTTP handler for RPC messages on rpcPath, and a debugging handler on debugPath. It is still necessary to invoke http.Serve(), typically in a go statement.

​	`HandleHTTP`方法为 rpcPath 上的 RPC 消息注册一个 HTTP 处理程序，并在 debugPath 上注册一个调试处理程序。仍然需要调用 http.Serve()，通常在 go 语句中。

#### (*Server) Register 

``` go 
func (server *Server) Register(rcvr any) error
```

Register publishes in the server the set of methods of the receiver value that satisfy the following conditions:

​	`Register`方法在服务器中发布接收器值的方法集，这些方法满足以下条件：

- exported method of exported type
- 导出类型的导出方法
- two arguments, both of exported type
- 两个实参，均为导出类型
- the second argument is a pointer
- 第二个实参是一个指针
- one return value, of type error
- 一个返回类型为 error 的返回值

It returns an error if the receiver is not an exported type or has no suitable methods. It also logs the error using package log. The client accesses each method using a string of the form "Type.Method", where Type is the receiver's concrete type.

​	如果接收器不是导出类型或没有合适的方法，则返回错误。它还使用包记录记录错误。客户端使用"Type.Method"格式的字符串访问每个方法，其中 `Type` 是接收器的具体类型。

#### (*Server) RegisterName 

``` go 
func (server *Server) RegisterName(name string, rcvr any) error
```

RegisterName is like Register but uses the provided name for the type instead of the receiver's concrete type.

​	`RegisterName`方法类似于 `Register`方法，但使用提供的名称替换接收器的具体类型。

#### (*Server) ServeCodec 

``` go 
func (server *Server) ServeCodec(codec ServerCodec)
```

ServeCodec is like ServeConn but uses the specified codec to decode requests and encode responses.

​	`ServeCodec`方法类似 `ServeConn`方法，但使用指定的编解码器解码请求和编码响应。

#### (*Server) ServeConn 

``` go 
func (server *Server) ServeConn(conn io.ReadWriteCloser)
```

ServeConn runs the server on a single connection. ServeConn blocks, serving the connection until the client hangs up. The caller typically invokes ServeConn in a go statement. ServeConn uses the gob wire format (see package gob) on the connection. To use an alternate codec, use ServeCodec. See NewClient's comment for information about concurrent access.

​	`ServeConn`方法在单个连接上运行服务器。`ServeConn`方法阻塞，直到客户端挂断连接为止。调用方通常在一个 go 语句中调用 `ServeConn`方法。`ServeConn`方法在连接上使用 gob wire 格式(参见 gob 包)。要使用其他编解码器，请使用 `ServeCodec`方法。有关并发访问的信息，请参见 NewClient 的注释。

#### (*Server) ServeHTTP 

``` go 
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)
```

ServeHTTP implements an http.Handler that answers RPC requests.

​	`ServeHTTP` 实现了一个 `http.Handler`，用于响应 RPC 请求。

#### (*Server) ServeRequest 

``` go 
func (server *Server) ServeRequest(codec ServerCodec) error
```

ServeRequest is like ServeCodec but synchronously serves a single request. It does not close the codec upon completion.

​	`ServeRequest`方法类似 `ServeCodec`，但同步地提供单个请求服务。完成后不关闭编解码器。

### type ServerCodec 

``` go 
type ServerCodec interface {
	ReadRequestHeader(*Request) error
	ReadRequestBody(any) error
	WriteResponse(*Response, any) error

	// Close can be called multiple times and must be idempotent.
    // Close 可以被多次调用，必须是幂等的。
	Close() error
}
```

A ServerCodec implements reading of RPC requests and writing of RPC responses for the server side of an RPC session. The server calls ReadRequestHeader and ReadRequestBody in pairs to read requests from the connection, and it calls WriteResponse to write a response back. The server calls Close when finished with the connection. ReadRequestBody may be called with a nil argument to force the body of the request to be read and discarded. See NewClient's comment for information about concurrent access.

​	`ServerCodec` 实现了读取 RPC 请求和写入 RPC 响应的功能，用于 RPC 会话的服务器端。服务器以成对的方式调用 ReadRequestHeader 和 ReadRequestBody 从连接中读取请求，并调用 WriteResponse 写入响应。服务器在完成连接后调用 Close。可以使用 nil 参数调用 ReadRequestBody，以强制读取并丢弃请求的正文。有关并发访问的信息，请参见 NewClient 的注释。

### type ServerError 

``` go 
type ServerError string
```

ServerError represents an error that has been returned from the remote side of the RPC connection.

​	`ServerError` 表示从 RPC 连接的远程一侧返回的错误。

#### (ServerError) Error 

``` go 
func (e ServerError) Error() string
```