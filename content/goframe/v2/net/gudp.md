+++
title = "gudp"
date = 2024-03-21T17:53:58+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gudp](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gudp)

Package gudp provides UDP server and client implementations.

​	软件包 gudp 提供 UDP 服务器和客户端实现。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/gudp/gudp_server.go#L21)

```go
const (
	// FreePortAddress marks the server listens using random free port.
	FreePortAddress = ":0"
)
```

## 变量

This section is empty.

## 函数

#### func GetFreePort

```go
func GetFreePort() (port int, err error)
```

GetFreePort retrieves and returns a port that is free.

​	GetFreePort 检索并返回一个空闲端口。

##### Example

``` go
```

#### func GetFreePorts

```go
func GetFreePorts(count int) (ports []int, err error)
```

GetFreePorts retrieves and returns specified number of ports that are free.

​	GetFreePorts 检索并返回指定数量的可用端口。

##### Example

``` go
```

#### func MustGetFreePort <-2.2.0

```go
func MustGetFreePort() (port int)
```

MustGetFreePort performs as GetFreePort, but it panics if any error occurs.

​	MustGetFreePort 以 GetFreePort 的形式执行，但如果发生任何错误，它会崩溃。

#### func NewNetConn

```go
func NewNetConn(remoteAddress string, localAddress ...string) (*net.UDPConn, error)
```

NewNetConn creates and returns a *net.UDPConn with given addresses.

​	NewNetConn 创建并返回一个 *net。具有给定地址的 UDPConn。

#### func Send

```go
func Send(address string, data []byte, retry ...Retry) error
```

Send writes data to `address` using UDP connection and then closes the connection. Note that it is used for short connection usage.

​	使用 UDP 连接发送写入数据， `address` 然后关闭连接。请注意，它用于短连接使用。

#### func SendRecv

```go
func SendRecv(address string, data []byte, receive int, retry ...Retry) ([]byte, error)
```

SendRecv writes data to `address` using UDP connection, reads response and then closes the connection. Note that it is used for short connection usage.

​	SendRecv 使用 UDP 连接将数据写入 `address` ，读取响应，然后关闭连接。请注意，它用于短连接使用。

## 类型

### type Conn

```go
type Conn struct {
	*net.UDPConn // Underlying UDP connection.
	// contains filtered or unexported fields
}
```

Conn handles the UDP connection.

​	Conn 处理 UDP 连接。

#### func NewConn

```go
func NewConn(remoteAddress string, localAddress ...string) (*Conn, error)
```

NewConn creates UDP connection to `remoteAddress`. The optional parameter `localAddress` specifies the local address for connection.

​	NewConn 创建与 `remoteAddress` 的 UDP 连接。可选参数 `localAddress` 指定用于连接的本地地址。

#### func NewConnByNetConn

```go
func NewConnByNetConn(udp *net.UDPConn) *Conn
```

NewConnByNetConn creates an UDP connection object with given *net.UDPConn object.

​	NewConnByNetConn 使用给定的 *net 创建 UDP 连接对象。UDPConn 对象。

#### (*Conn) Recv

```go
func (c *Conn) Recv(buffer int, retry ...Retry) ([]byte, error)
```

Recv receives and returns data from remote address. The parameter `buffer` is used for customizing the receiving buffer size. If `buffer` <= 0, it uses the default buffer size, which is 1024 byte.

​	Recv 从远程地址接收和返回数据。该参数 `buffer` 用于自定义接收缓冲区大小。如果 `buffer` <= 0，则使用默认缓冲区大小，即 1024 字节。

There’s package border in UDP protocol, we can receive a complete package if specified buffer size is big enough. VERY NOTE that we should receive the complete package in once or else the leftover package data would be dropped.

​	UDP协议中有包边框，如果指定的缓冲区大小足够大，我们可以收到一个完整的包。非常注意，我们应该一次性收到完整的包，否则剩余的包数据将被丢弃。

#### (*Conn) RecvWithTimeout

```go
func (c *Conn) RecvWithTimeout(length int, timeout time.Duration, retry ...Retry) (data []byte, err error)
```

RecvWithTimeout reads data from remote address with timeout.

​	RecvWithTimeout 使用超时从远程地址读取数据。

#### (*Conn) RemoteAddr

```go
func (c *Conn) RemoteAddr() net.Addr
```

RemoteAddr returns the remote address of current UDP connection. Note that it cannot use c.conn.RemoteAddr() as it is nil.

​	RemoteAddr 返回当前 UDP 连接的远程地址。请注意，它不能使用 c.conn.RemoteAddr（），因为它是 nil。

#### (*Conn) Send

```go
func (c *Conn) Send(data []byte, retry ...Retry) (err error)
```

Send writes data to remote address.

​	将写入数据发送到远程地址。

#### (*Conn) SendRecv

```go
func (c *Conn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error)
```

SendRecv writes data to connection and blocks reading response.

​	SendRecv 将数据写入连接并阻止读取响应。

#### (*Conn) SendRecvWithTimeout

```go
func (c *Conn) SendRecvWithTimeout(data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error)
```

SendRecvWithTimeout writes data to connection and reads response with timeout.

​	SendRecvWithTimeout 将数据写入连接，并在超时时时读取响应。

#### (*Conn) SendWithTimeout

```go
func (c *Conn) SendWithTimeout(data []byte, timeout time.Duration, retry ...Retry) (err error)
```

SendWithTimeout writes data to connection with timeout.

​	SendWithTimeout 将数据写入 timeout 的连接。

#### (*Conn) SetBufferWaitRecv

```go
func (c *Conn) SetBufferWaitRecv(d time.Duration)
```

SetBufferWaitRecv sets the buffer waiting timeout when reading all data from connection. The waiting duration cannot be too long which might delay receiving data from remote address.

​	SetBufferWaitRecv 设置从连接读取所有数据时的缓冲区等待超时。等待时间不能太长，这可能会延迟从远程地址接收数据。

#### (*Conn) SetDeadline

```go
func (c *Conn) SetDeadline(t time.Time) (err error)
```

SetDeadline sets the read and write deadlines associated with the connection.

​	SetDeadline 设置与连接关联的读取和写入截止时间。

#### (*Conn) SetDeadlineRecv

```go
func (c *Conn) SetDeadlineRecv(t time.Time) (err error)
```

SetDeadlineRecv sets the read deadline associated with the connection.

​	SetDeadlineRecv 设置与连接关联的读取截止时间。

#### (*Conn) SetDeadlineSend

```go
func (c *Conn) SetDeadlineSend(t time.Time) (err error)
```

SetDeadlineSend sets the deadline of sending for current connection.

​	SetDeadlineSend 设置当前连接的发送截止时间。

### type Retry

```go
type Retry struct {
	Count    int           // Max retry count.
	Interval time.Duration // Retry interval.
}
```

### type Server

```go
type Server struct {
	// contains filtered or unexported fields
}
```

Server is the UDP server.

​	服务器是 UDP 服务器。

#### func GetServer

```go
func GetServer(name ...interface{}) *Server
```

GetServer creates and returns an UDP server instance with given name.

​	GetServer 创建并返回具有给定名称的 UDP 服务器实例。

#### func NewServer

```go
func NewServer(address string, handler func(*Conn), name ...string) *Server
```

NewServer creates and returns an UDP server. The optional parameter `name` is used to specify its name, which can be used for GetServer function to retrieve its instance.

​	NewServer 创建并返回 UDP 服务器。可选参数 `name` 用于指定其名称，该名称可用于 GetServer 函数检索其实例。

#### (*Server) Close

```go
func (s *Server) Close() (err error)
```

Close closes the connection. It will make server shutdowns immediately.

​	关闭关闭连接。它将立即关闭服务器。

#### (*Server) GetListenedAddress

```go
func (s *Server) GetListenedAddress() string
```

GetListenedAddress retrieves and returns the address string which are listened by current server.

​	GetListenedAddress 检索并返回当前服务器侦听的地址字符串。

#### (*Server) GetListenedPort

```go
func (s *Server) GetListenedPort() int
```

GetListenedPort retrieves and returns one port which is listened to by current server.

​	GetListenedPort 检索并返回当前服务器侦听的一个端口。

#### (*Server) Run

```go
func (s *Server) Run() error
```

Run starts listening UDP connection.

​	Run 开始侦听 UDP 连接。

#### (*Server) SetAddress

```go
func (s *Server) SetAddress(address string)
```

SetAddress sets the server address for UDP server.

​	SetAddress 设置 UDP 服务器的服务器地址。

#### (*Server) SetHandler

```go
func (s *Server) SetHandler(handler func(*Conn))
```

SetHandler sets the connection handler for UDP server.

​	SetHandler 设置 UDP 服务器的连接处理程序。