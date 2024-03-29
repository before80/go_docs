+++
title = "gtcp"
date = 2024-03-21T17:53:39+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gtcp](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gtcp)

Package gtcp provides TCP server and client implementations.

​	软件包 gtcp 提供 TCP 服务器和客户端实现。

#### Examples 例子

- [GetFreePort
  获取自由港](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gtcp#example-GetFreePort)
- [GetFreePorts
  获取自由港](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gtcp#example-GetFreePorts)

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/gtcp/gtcp_server.go#L22)

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

#### func LoadKeyCrt

```go
func LoadKeyCrt(crtFile, keyFile string) (*tls.Config, error)
```

LoadKeyCrt creates and returns a TLS configuration object with given certificate and key files.

​	LoadKeyCrt 创建并返回具有给定证书和密钥文件的 TLS 配置对象。

#### func MustGetFreePort

```go
func MustGetFreePort() int
```

MustGetFreePort performs as GetFreePort, but it panics is any error occurs.

​	MustGetFreePort 以 GetFreePort 的形式执行，但如果发生任何错误，它就会崩溃。

#### func NewNetConn

```go
func NewNetConn(address string, timeout ...time.Duration) (net.Conn, error)
```

NewNetConn creates and returns a net.Conn with given address like “127.0.0.1:80”. The optional parameter `timeout` specifies the timeout for dialing connection.

​	NewNetConn 创建并返回一个网络。具有给定地址的 Conn，例如“127.0.0.1：80”。可选参数 `timeout` 指定拨号连接的超时时间。

#### func NewNetConnKeyCrt

```go
func NewNetConnKeyCrt(addr, crtFile, keyFile string, timeout ...time.Duration) (net.Conn, error)
```

NewNetConnKeyCrt creates and returns a TLS net.Conn with given TLS certificate and key files and address like “127.0.0.1:80”. The optional parameter `timeout` specifies the timeout for dialing connection.

​	NewNetConnKeyCrt 创建并返回 TLS 网络。具有给定的 TLS 证书和密钥文件以及地址（如“127.0.0.1：80”）的 Conn。可选参数 `timeout` 指定拨号连接的超时时间。

#### func NewNetConnTLS

```go
func NewNetConnTLS(address string, tlsConfig *tls.Config, timeout ...time.Duration) (net.Conn, error)
```

NewNetConnTLS creates and returns a TLS net.Conn with given address like “127.0.0.1:80”. The optional parameter `timeout` specifies the timeout for dialing connection.

​	NewNetConnTLS 创建并返回 TLS 网络。具有给定地址的 Conn，例如“127.0.0.1：80”。可选参数 `timeout` 指定拨号连接的超时时间。

#### func Send

```go
func Send(address string, data []byte, retry ...Retry) error
```

Send creates connection to `address`, writes `data` to the connection and then closes the connection. The optional parameter `retry` specifies the retry policy when fails in writing data.

​	发送创建连接 `address` ，写 `data` 入连接，然后关闭连接。可选参数 `retry` 指定写入数据失败时的重试策略。

#### func SendPkg

```go
func SendPkg(address string, data []byte, option ...PkgOption) error
```

SendPkg sends a package containing `data` to `address` and closes the connection. The optional parameter `option` specifies the package options for sending.

​	SendPkg 发送一个包含 `data` to `address` 的包并关闭连接。optional 参数 `option` 指定用于发送的包选项。

#### func SendPkgWithTimeout

```go
func SendPkgWithTimeout(address string, data []byte, timeout time.Duration, option ...PkgOption) error
```

SendPkgWithTimeout sends a package containing `data` to `address` with timeout limitation and closes the connection. The optional parameter `option` specifies the package options for sending.

​	SendPkgWithTimeout 发送包含 `data` 超时限制的包 `address` 并关闭连接。optional 参数 `option` 指定用于发送的包选项。

#### func SendRecv

```go
func SendRecv(address string, data []byte, length int, retry ...Retry) ([]byte, error)
```

SendRecv creates connection to `address`, writes `data` to the connection, receives response and then closes the connection.

​	SendRecv 创建与 `address` 的连接，写 `data` 入连接，接收响应，然后关闭连接。

The parameter `length` specifies the bytes count waiting to receive. It receives all buffer content and returns if `length` is -1.

​	该参数 `length` 指定等待接收的字节数。它接收所有缓冲区内容，如果 `length` 为 -1，则返回。

The optional parameter `retry` specifies the retry policy when fails in writing data.

​	可选参数 `retry` 指定写入数据失败时的重试策略。

#### func SendRecvPkg

```go
func SendRecvPkg(address string, data []byte, option ...PkgOption) ([]byte, error)
```

SendRecvPkg sends a package containing `data` to `address`, receives the response and closes the connection. The optional parameter `option` specifies the package options for sending.

​	SendRecvPkg 发送一个包含 `data` 的 `address` 包，接收响应并关闭连接。optional 参数 `option` 指定用于发送的包选项。

#### func SendRecvPkgWithTimeout

```go
func SendRecvPkgWithTimeout(address string, data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error)
```

SendRecvPkgWithTimeout sends a package containing `data` to `address`, receives the response with timeout limitation and closes the connection. The optional parameter `option` specifies the package options for sending.

​	SendRecvPkgWithTimeout 发送一个包含 `data` 的 `address` 包，接收具有超时限制的响应并关闭连接。optional 参数 `option` 指定用于发送的包选项。

#### func SendRecvWithTimeout

```go
func SendRecvWithTimeout(address string, data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error)
```

SendRecvWithTimeout does SendRecv logic with reading timeout limitation.

​	SendRecvWithTimeout 执行具有读取超时限制的 SendRecv 逻辑。

#### func SendWithTimeout

```go
func SendWithTimeout(address string, data []byte, timeout time.Duration, retry ...Retry) error
```

SendWithTimeout does Send logic with writing timeout limitation.

​	SendWithTimeout 执行具有写入超时限制的 Send 逻辑。

## 类型

### type Conn

```go
type Conn struct {
	net.Conn // Underlying TCP connection object.
	// contains filtered or unexported fields
}
```

Conn is the TCP connection object.

​	Conn 是 TCP 连接对象。

#### func NewConn

```go
func NewConn(addr string, timeout ...time.Duration) (*Conn, error)
```

NewConn creates and returns a new connection with given address.

​	NewConn 创建并返回具有给定地址的新连接。

#### func NewConnByNetConn

```go
func NewConnByNetConn(conn net.Conn) *Conn
```

NewConnByNetConn creates and returns a TCP connection object with given net.Conn object.

​	NewConnByNetConn 创建并返回具有给定网络的 TCP 连接对象。Conn 对象。

#### func NewConnKeyCrt

```go
func NewConnKeyCrt(addr, crtFile, keyFile string) (*Conn, error)
```

NewConnKeyCrt creates and returns a new TLS connection with given address and TLS certificate and key files.

​	NewConnKeyCrt 创建并返回具有给定地址、TLS 证书和密钥文件的新 TLS 连接。

#### func NewConnTLS

```go
func NewConnTLS(addr string, tlsConfig *tls.Config) (*Conn, error)
```

NewConnTLS creates and returns a new TLS connection with given address and TLS configuration.

​	NewConnTLS 创建并返回具有给定地址和 TLS 配置的新 TLS 连接。

#### (*Conn) Recv

```go
func (c *Conn) Recv(length int, retry ...Retry) ([]byte, error)
```

Recv receives and returns data from the connection.

​	Recv 从连接接收和返回数据。

Note that,

​	请注意，

1. If length = 0, which means it receives the data from current buffer and returns immediately.
   如果 length = 0，则表示它从当前缓冲区接收数据并立即返回。
2. If length < 0, which means it receives all data from connection and returns it until no data from connection. Developers should notice the package parsing yourself if you decide receiving all data from buffer.
   如果长度< 0，则表示它从连接接收所有数据并返回它，直到连接中没有数据。如果您决定从缓冲区接收所有数据，开发人员应该会注意到包会自行解析。
3. If length > 0, which means it blocks reading data from connection until length size was received. It is the most commonly used length value for data receiving.
   如果长度> 0，则表示在收到长度大小之前，它会阻止从连接读取数据。它是数据接收最常用的长度值。

#### (*Conn) RecvLine

```go
func (c *Conn) RecvLine(retry ...Retry) ([]byte, error)
```

RecvLine reads data from the connection until reads char ‘\n’. Note that the returned result does not contain the last char ‘\n’.

​	RecvLine 从连接中读取数据，直到读取 char '\n'。请注意，返回的结果不包含最后一个字符“\n”。

#### (*Conn) RecvPkg

```go
func (c *Conn) RecvPkg(option ...PkgOption) (result []byte, err error)
```

RecvPkg receives data from connection using simple package protocol.

​	RecvPkg 使用简单的包协议从连接接收数据。

#### (*Conn) RecvPkgWithTimeout

```go
func (c *Conn) RecvPkgWithTimeout(timeout time.Duration, option ...PkgOption) (data []byte, err error)
```

RecvPkgWithTimeout reads data from connection with timeout using simple package protocol.

​	RecvPkgWithTimeout 使用简单的包协议从具有超时的连接中读取数据。

#### (*Conn) RecvTill

```go
func (c *Conn) RecvTill(til []byte, retry ...Retry) ([]byte, error)
```

RecvTill reads data from the connection until reads bytes `til`. Note that the returned result contains the last bytes `til`.

​	RecvTill 从连接中读取数据，直到读取字节 `til` 。请注意，返回的结果包含最后一个字节 `til` 。

#### (*Conn) RecvWithTimeout

```go
func (c *Conn) RecvWithTimeout(length int, timeout time.Duration, retry ...Retry) (data []byte, err error)
```

RecvWithTimeout reads data from the connection with timeout.

​	RecvWithTimeout 从具有超时的连接中读取数据。

#### (*Conn) Send

```go
func (c *Conn) Send(data []byte, retry ...Retry) error
```

Send writes data to remote address.

​	将写入数据发送到远程地址。

#### (*Conn) SendPkg

```go
func (c *Conn) SendPkg(data []byte, option ...PkgOption) error
```

SendPkg send data using simple package protocol.

​	SendPkg 使用简单的包协议发送数据。

Simple package protocol: DataLength(24bit)|DataField(variant)。

​	简单包协议：DataLength（24bit）|DataField（variant）。

Note that, 1. The DataLength is the length of DataField, which does not contain the header size. 2. The integer bytes of the package are encoded using BigEndian order.

​	请注意，1.DataLength 是 DataField 的长度，它不包含标头大小。2. 包的整数字节使用 BigEndian 顺序进行编码。

#### (*Conn) SendPkgWithTimeout

```go
func (c *Conn) SendPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) (err error)
```

SendPkgWithTimeout writes data to connection with timeout using simple package protocol.

​	SendPkgWithTimeout 使用简单的包协议将数据写入具有超时的连接。

#### (*Conn) SendRecv

```go
func (c *Conn) SendRecv(data []byte, length int, retry ...Retry) ([]byte, error)
```

SendRecv writes data to the connection and blocks reading response.

​	SendRecv 将数据写入连接并阻止读取响应。

#### (*Conn) SendRecvPkg

```go
func (c *Conn) SendRecvPkg(data []byte, option ...PkgOption) ([]byte, error)
```

SendRecvPkg writes data to connection and blocks reading response using simple package protocol.

​	SendRecvPkg 将数据写入连接，并使用简单的包协议阻止读取响应。

#### (*Conn) SendRecvPkgWithTimeout

```go
func (c *Conn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error)
```

SendRecvPkgWithTimeout writes data to connection and reads response with timeout using simple package protocol.

​	SendRecvPkgWithTimeout 将数据写入连接，并使用简单的包协议读取带有超时的响应。

#### (*Conn) SendRecvWithTimeout

```go
func (c *Conn) SendRecvWithTimeout(data []byte, length int, timeout time.Duration, retry ...Retry) ([]byte, error)
```

SendRecvWithTimeout writes data to the connection and reads response with timeout.

​	SendRecvWithTimeout 将数据写入连接，并读取具有超时的响应。

#### (*Conn) SendWithTimeout

```go
func (c *Conn) SendWithTimeout(data []byte, timeout time.Duration, retry ...Retry) (err error)
```

SendWithTimeout writes data to the connection with timeout.

​	SendWithTimeout 将数据写入具有超时的连接。

#### (*Conn) SetBufferWaitRecv

```go
func (c *Conn) SetBufferWaitRecv(bufferWaitDuration time.Duration)
```

SetBufferWaitRecv sets the buffer waiting timeout when reading all data from connection. The waiting duration cannot be too long which might delay receiving data from remote address.

​	SetBufferWaitRecv 设置从连接读取所有数据时的缓冲区等待超时。等待时间不能太长，这可能会延迟从远程地址接收数据。

#### (*Conn) SetDeadline

```go
func (c *Conn) SetDeadline(t time.Time) (err error)
```

SetDeadline sets the deadline for current connection.

​	SetDeadline 设置当前连接的截止时间。

#### (*Conn) SetDeadlineRecv

```go
func (c *Conn) SetDeadlineRecv(t time.Time) (err error)
```

SetDeadlineRecv sets the deadline of receiving for current connection.

​	SetDeadlineRecv 设置当前连接的接收截止时间。

#### (*Conn) SetDeadlineSend

```go
func (c *Conn) SetDeadlineSend(t time.Time) (err error)
```

SetDeadlineSend sets the deadline of sending for current connection.

​	SetDeadlineSend 设置当前连接的发送截止时间。

### type PkgOption

```go
type PkgOption struct {
	// HeaderSize is used to mark the data length for next data receiving.
	// It's 2 bytes in default, 4 bytes max, which stands for the max data length
	// from 65535 to 4294967295 bytes.
	HeaderSize int

	// MaxDataSize is the data field size in bytes for data length validation.
	// If it's not manually set, it'll automatically be set correspondingly with the HeaderSize.
	MaxDataSize int

	// Retry policy when operation fails.
	Retry Retry
}
```

PkgOption is package option for simple protocol.

​	PkgOption 是简单协议的包选项。

### type PoolConn

```go
type PoolConn struct {
	*Conn // Underlying connection object.
	// contains filtered or unexported fields
}
```

PoolConn is a connection with pool feature for TCP. Note that it is NOT a pool or connection manager, it is just a TCP connection object.

​	PoolConn 是具有 TCP 池功能的连接。请注意，它不是池或连接管理器，它只是一个 TCP 连接对象。

#### func NewPoolConn

```go
func NewPoolConn(addr string, timeout ...time.Duration) (*PoolConn, error)
```

NewPoolConn creates and returns a connection with pool feature.

​	NewPoolConn 创建并返回具有池功能的连接。

#### (*PoolConn) Close

```go
func (c *PoolConn) Close() error
```

Close puts back the connection to the pool if it’s active, or closes the connection if it’s not active.

​	如果池处于活动状态，则关闭将连接放回池，如果连接未处于活动状态，则关闭连接。

Note that, if `c` calls Close function closing itself, `c` can not be used again.

​	需要注意的是，如果 `c` 调用关闭函数本身， `c` 则不能再次使用。

#### (*PoolConn) Recv

```go
func (c *PoolConn) Recv(length int, retry ...Retry) ([]byte, error)
```

Recv receives data from the connection.

​	Recv 从连接接收数据。

#### (*PoolConn) RecvLine

```go
func (c *PoolConn) RecvLine(retry ...Retry) ([]byte, error)
```

RecvLine reads data from the connection until reads char ‘\n’. Note that the returned result does not contain the last char ‘\n’.

​	RecvLine 从连接中读取数据，直到读取 char '\n'。请注意，返回的结果不包含最后一个字符“\n”。

#### (*PoolConn) RecvPkg

```go
func (c *PoolConn) RecvPkg(option ...PkgOption) ([]byte, error)
```

RecvPkg receives package from connection using simple package protocol. The optional parameter `option` specifies the package options for receiving.

​	RecvPkg 使用简单的包协议从连接接收包。optional 参数 `option` 指定用于接收的包选项。

#### (*PoolConn) RecvPkgWithTimeout

```go
func (c *PoolConn) RecvPkgWithTimeout(timeout time.Duration, option ...PkgOption) (data []byte, err error)
```

RecvPkgWithTimeout reads data from connection with timeout using simple package protocol.

​	RecvPkgWithTimeout 使用简单的包协议从具有超时的连接中读取数据。

#### (*PoolConn) RecvTill

```go
func (c *PoolConn) RecvTill(til []byte, retry ...Retry) ([]byte, error)
```

RecvTill reads data from the connection until reads bytes `til`. Note that the returned result contains the last bytes `til`.

​	RecvTill 从连接中读取数据，直到读取字节 `til` 。请注意，返回的结果包含最后一个字节 `til` 。

#### (*PoolConn) RecvWithTimeout

```go
func (c *PoolConn) RecvWithTimeout(length int, timeout time.Duration, retry ...Retry) (data []byte, err error)
```

RecvWithTimeout reads data from the connection with timeout.

​	RecvWithTimeout 从具有超时的连接中读取数据。

#### (*PoolConn) Send

```go
func (c *PoolConn) Send(data []byte, retry ...Retry) error
```

Send writes data to the connection. It retrieves a new connection from its pool if it fails writing data.

​	向连接发送写入数据。如果写入数据失败，它将从其池中检索新连接。

#### (*PoolConn) SendPkg

```go
func (c *PoolConn) SendPkg(data []byte, option ...PkgOption) (err error)
```

SendPkg sends a package containing `data` to the connection. The optional parameter `option` specifies the package options for sending.

​	SendPkg 向连接发送一个包含的 `data` 包。optional 参数 `option` 指定用于发送的包选项。

#### (*PoolConn) SendPkgWithTimeout

```go
func (c *PoolConn) SendPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) (err error)
```

SendPkgWithTimeout writes data to connection with timeout using simple package protocol.

​	SendPkgWithTimeout 使用简单的包协议将数据写入具有超时的连接。

#### (*PoolConn) SendRecv

```go
func (c *PoolConn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error)
```

SendRecv writes data to the connection and blocks reading response.

​	SendRecv 将数据写入连接并阻止读取响应。

#### (*PoolConn) SendRecvPkg

```go
func (c *PoolConn) SendRecvPkg(data []byte, option ...PkgOption) ([]byte, error)
```

SendRecvPkg writes data to connection and blocks reading response using simple package protocol.

​	SendRecvPkg 将数据写入连接，并使用简单的包协议阻止读取响应。

#### (*PoolConn) SendRecvPkgWithTimeout

```go
func (c *PoolConn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error)
```

SendRecvPkgWithTimeout reads data from connection with timeout using simple package protocol.

​	SendRecvPkgWithTimeout 使用简单的包协议从超时连接中读取数据。

#### (*PoolConn) SendRecvWithTimeout

```go
func (c *PoolConn) SendRecvWithTimeout(data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error)
```

SendRecvWithTimeout writes data to the connection and reads response with timeout.

​	SendRecvWithTimeout 将数据写入连接，并读取具有超时的响应。

#### (*PoolConn) SendWithTimeout

```go
func (c *PoolConn) SendWithTimeout(data []byte, timeout time.Duration, retry ...Retry) (err error)
```

SendWithTimeout writes data to the connection with timeout.

​	SendWithTimeout 将数据写入具有超时的连接。

### type Retry

```go
type Retry struct {
	Count    int           // Retry count.
	Interval time.Duration // Retry interval.
}
```

### type Server

```go
type Server struct {
	// contains filtered or unexported fields
}
```

Server is a TCP server.

​	服务器是 TCP 服务器。

#### func GetServer

```go
func GetServer(name ...interface{}) *Server
```

GetServer returns the TCP server with specified `name`, or it returns a new normal TCP server named `name` if it does not exist. The parameter `name` is used to specify the TCP server

​	GetServer 返回指定 `name` 的 TCP 服务器，或者返回一个新的普通 TCP 服务器，如果它不存在，则命名为 `name` 。该参数 `name` 用于指定 TCP 服务器

#### func NewServer

```go
func NewServer(address string, handler func(*Conn), name ...string) *Server
```

NewServer creates and returns a new normal TCP server. The parameter `name` is optional, which is used to specify the instance name of the server.

​	NewServer 创建并返回一个新的普通 TCP 服务器。该参数 `name` 是可选的，用于指定服务器的实例名称。

#### func NewServerKeyCrt

```go
func NewServerKeyCrt(address, crtFile, keyFile string, handler func(*Conn), name ...string) (*Server, error)
```

NewServerKeyCrt creates and returns a new TCP server with TLS support. The parameter `name` is optional, which is used to specify the instance name of the server.

​	NewServerKeyCrt 创建并返回支持 TLS 的新 TCP 服务器。该参数 `name` 是可选的，用于指定服务器的实例名称。

#### func NewServerTLS

```go
func NewServerTLS(address string, tlsConfig *tls.Config, handler func(*Conn), name ...string) *Server
```

NewServerTLS creates and returns a new TCP server with TLS support. The parameter `name` is optional, which is used to specify the instance name of the server.

​	NewServerTLS 创建并返回支持 TLS 的新 TCP 服务器。该参数 `name` 是可选的，用于指定服务器的实例名称。

#### (*Server) Close

```go
func (s *Server) Close() error
```

Close closes the listener and shutdowns the server.

​	关闭将关闭侦听器并关闭服务器。

#### (*Server) GetAddress

```go
func (s *Server) GetAddress() string
```

GetAddress get the listening address for server.

​	GetAddress 获取服务器的侦听地址。

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
func (s *Server) Run() (err error)
```

Run starts running the TCP Server.

​	Run 开始运行 TCP 服务器。

#### (*Server) SetAddress

```go
func (s *Server) SetAddress(address string)
```

SetAddress sets the listening address for server.

​	SetAddress 设置服务器的侦听地址。

#### (*Server) SetHandler

```go
func (s *Server) SetHandler(handler func(*Conn))
```

SetHandler sets the connection handler for server.

​	SetHandler 设置服务器的连接处理程序。

#### (*Server) SetTLSConfig

```go
func (s *Server) SetTLSConfig(tlsConfig *tls.Config)
```

SetTLSConfig sets the TLS configuration of server.

​	SetTLSConfig 设置服务器的 TLS 配置。

#### (*Server) SetTLSKeyCrt

```go
func (s *Server) SetTLSKeyCrt(crtFile, keyFile string) error
```

SetTLSKeyCrt sets the certificate and key file for TLS configuration of server.

​	SetTLSKeyCrt 设置服务器 TLS 配置的证书和密钥文件。