+++
title = "gtcp"
date = 2024-03-21T17:53:39+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gtcp

Package gtcp provides TCP server and client implementations.

#### Examples 

- [GetFreePort](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gtcp#example-GetFreePort)
- [GetFreePorts](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gtcp#example-GetFreePorts)

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/gtcp/gtcp_server.go#L22)

``` go
const (
	// FreePortAddress marks the server listens using random free port.
	FreePortAddress = ":0"
)
```

### Variables 

This section is empty.

### Functions 

##### func GetFreePort 

``` go
func GetFreePort() (port int, err error)
```

GetFreePort retrieves and returns a port that is free.

##### Example

``` go
```
##### func GetFreePorts 

``` go
func GetFreePorts(count int) (ports []int, err error)
```

GetFreePorts retrieves and returns specified number of ports that are free.

##### Example

``` go
```
##### func LoadKeyCrt 

``` go
func LoadKeyCrt(crtFile, keyFile string) (*tls.Config, error)
```

LoadKeyCrt creates and returns a TLS configuration object with given certificate and key files.

##### func MustGetFreePort 

``` go
func MustGetFreePort() int
```

MustGetFreePort performs as GetFreePort, but it panics is any error occurs.

##### func NewNetConn 

``` go
func NewNetConn(address string, timeout ...time.Duration) (net.Conn, error)
```

NewNetConn creates and returns a net.Conn with given address like "127.0.0.1:80". The optional parameter `timeout` specifies the timeout for dialing connection.

##### func NewNetConnKeyCrt 

``` go
func NewNetConnKeyCrt(addr, crtFile, keyFile string, timeout ...time.Duration) (net.Conn, error)
```

NewNetConnKeyCrt creates and returns a TLS net.Conn with given TLS certificate and key files and address like "127.0.0.1:80". The optional parameter `timeout` specifies the timeout for dialing connection.

##### func NewNetConnTLS 

``` go
func NewNetConnTLS(address string, tlsConfig *tls.Config, timeout ...time.Duration) (net.Conn, error)
```

NewNetConnTLS creates and returns a TLS net.Conn with given address like "127.0.0.1:80". The optional parameter `timeout` specifies the timeout for dialing connection.

##### func Send 

``` go
func Send(address string, data []byte, retry ...Retry) error
```

Send creates connection to `address`, writes `data` to the connection and then closes the connection. The optional parameter `retry` specifies the retry policy when fails in writing data.

##### func SendPkg 

``` go
func SendPkg(address string, data []byte, option ...PkgOption) error
```

SendPkg sends a package containing `data` to `address` and closes the connection. The optional parameter `option` specifies the package options for sending.

##### func SendPkgWithTimeout 

``` go
func SendPkgWithTimeout(address string, data []byte, timeout time.Duration, option ...PkgOption) error
```

SendPkgWithTimeout sends a package containing `data` to `address` with timeout limitation and closes the connection. The optional parameter `option` specifies the package options for sending.

##### func SendRecv 

``` go
func SendRecv(address string, data []byte, length int, retry ...Retry) ([]byte, error)
```

SendRecv creates connection to `address`, writes `data` to the connection, receives response and then closes the connection.

The parameter `length` specifies the bytes count waiting to receive. It receives all buffer content and returns if `length` is -1.

The optional parameter `retry` specifies the retry policy when fails in writing data.

##### func SendRecvPkg 

``` go
func SendRecvPkg(address string, data []byte, option ...PkgOption) ([]byte, error)
```

SendRecvPkg sends a package containing `data` to `address`, receives the response and closes the connection. The optional parameter `option` specifies the package options for sending.

##### func SendRecvPkgWithTimeout 

``` go
func SendRecvPkgWithTimeout(address string, data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error)
```

SendRecvPkgWithTimeout sends a package containing `data` to `address`, receives the response with timeout limitation and closes the connection. The optional parameter `option` specifies the package options for sending.

##### func SendRecvWithTimeout 

``` go
func SendRecvWithTimeout(address string, data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error)
```

SendRecvWithTimeout does SendRecv logic with reading timeout limitation.

##### func SendWithTimeout 

``` go
func SendWithTimeout(address string, data []byte, timeout time.Duration, retry ...Retry) error
```

SendWithTimeout does Send logic with writing timeout limitation.

### Types 

#### type Conn 

``` go
type Conn struct {
	net.Conn // Underlying TCP connection object.
	// contains filtered or unexported fields
}
```

Conn is the TCP connection object.

##### func NewConn 

``` go
func NewConn(addr string, timeout ...time.Duration) (*Conn, error)
```

NewConn creates and returns a new connection with given address.

##### func NewConnByNetConn 

``` go
func NewConnByNetConn(conn net.Conn) *Conn
```

NewConnByNetConn creates and returns a TCP connection object with given net.Conn object.

##### func NewConnKeyCrt 

``` go
func NewConnKeyCrt(addr, crtFile, keyFile string) (*Conn, error)
```

NewConnKeyCrt creates and returns a new TLS connection with given address and TLS certificate and key files.

##### func NewConnTLS 

``` go
func NewConnTLS(addr string, tlsConfig *tls.Config) (*Conn, error)
```

NewConnTLS creates and returns a new TLS connection with given address and TLS configuration.

##### (*Conn) Recv 

``` go
func (c *Conn) Recv(length int, retry ...Retry) ([]byte, error)
```

Recv receives and returns data from the connection.

Note that,

1. If length = 0, which means it receives the data from current buffer and returns immediately.
2. If length < 0, which means it receives all data from connection and returns it until no data from connection. Developers should notice the package parsing yourself if you decide receiving all data from buffer.
3. If length > 0, which means it blocks reading data from connection until length size was received. It is the most commonly used length value for data receiving.

##### (*Conn) RecvLine 

``` go
func (c *Conn) RecvLine(retry ...Retry) ([]byte, error)
```

RecvLine reads data from the connection until reads char '\n'. Note that the returned result does not contain the last char '\n'.

##### (*Conn) RecvPkg 

``` go
func (c *Conn) RecvPkg(option ...PkgOption) (result []byte, err error)
```

RecvPkg receives data from connection using simple package protocol.

##### (*Conn) RecvPkgWithTimeout 

``` go
func (c *Conn) RecvPkgWithTimeout(timeout time.Duration, option ...PkgOption) (data []byte, err error)
```

RecvPkgWithTimeout reads data from connection with timeout using simple package protocol.

##### (*Conn) RecvTill 

``` go
func (c *Conn) RecvTill(til []byte, retry ...Retry) ([]byte, error)
```

RecvTill reads data from the connection until reads bytes `til`. Note that the returned result contains the last bytes `til`.

##### (*Conn) RecvWithTimeout 

``` go
func (c *Conn) RecvWithTimeout(length int, timeout time.Duration, retry ...Retry) (data []byte, err error)
```

RecvWithTimeout reads data from the connection with timeout.

##### (*Conn) Send 

``` go
func (c *Conn) Send(data []byte, retry ...Retry) error
```

Send writes data to remote address.

##### (*Conn) SendPkg 

``` go
func (c *Conn) SendPkg(data []byte, option ...PkgOption) error
```

SendPkg send data using simple package protocol.

Simple package protocol: DataLength(24bit)|DataField(variant)。

Note that, 1. The DataLength is the length of DataField, which does not contain the header size. 2. The integer bytes of the package are encoded using BigEndian order.

##### (*Conn) SendPkgWithTimeout 

``` go
func (c *Conn) SendPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) (err error)
```

SendPkgWithTimeout writes data to connection with timeout using simple package protocol.

##### (*Conn) SendRecv 

``` go
func (c *Conn) SendRecv(data []byte, length int, retry ...Retry) ([]byte, error)
```

SendRecv writes data to the connection and blocks reading response.

##### (*Conn) SendRecvPkg 

``` go
func (c *Conn) SendRecvPkg(data []byte, option ...PkgOption) ([]byte, error)
```

SendRecvPkg writes data to connection and blocks reading response using simple package protocol.

##### (*Conn) SendRecvPkgWithTimeout 

``` go
func (c *Conn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error)
```

SendRecvPkgWithTimeout writes data to connection and reads response with timeout using simple package protocol.

##### (*Conn) SendRecvWithTimeout 

``` go
func (c *Conn) SendRecvWithTimeout(data []byte, length int, timeout time.Duration, retry ...Retry) ([]byte, error)
```

SendRecvWithTimeout writes data to the connection and reads response with timeout.

##### (*Conn) SendWithTimeout 

``` go
func (c *Conn) SendWithTimeout(data []byte, timeout time.Duration, retry ...Retry) (err error)
```

SendWithTimeout writes data to the connection with timeout.

##### (*Conn) SetBufferWaitRecv <-2.3.0

``` go
func (c *Conn) SetBufferWaitRecv(bufferWaitDuration time.Duration)
```

SetBufferWaitRecv sets the buffer waiting timeout when reading all data from connection. The waiting duration cannot be too long which might delay receiving data from remote address.

##### (*Conn) SetDeadline 

``` go
func (c *Conn) SetDeadline(t time.Time) (err error)
```

SetDeadline sets the deadline for current connection.

##### (*Conn) SetDeadlineRecv <-2.3.0

``` go
func (c *Conn) SetDeadlineRecv(t time.Time) (err error)
```

SetDeadlineRecv sets the deadline of receiving for current connection.

##### (*Conn) SetDeadlineSend <-2.3.0

``` go
func (c *Conn) SetDeadlineSend(t time.Time) (err error)
```

SetDeadlineSend sets the deadline of sending for current connection.

#### type PkgOption 

``` go
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

#### type PoolConn 

``` go
type PoolConn struct {
	*Conn // Underlying connection object.
	// contains filtered or unexported fields
}
```

PoolConn is a connection with pool feature for TCP. Note that it is NOT a pool or connection manager, it is just a TCP connection object.

##### func NewPoolConn 

``` go
func NewPoolConn(addr string, timeout ...time.Duration) (*PoolConn, error)
```

NewPoolConn creates and returns a connection with pool feature.

##### (*PoolConn) Close 

``` go
func (c *PoolConn) Close() error
```

Close puts back the connection to the pool if it's active, or closes the connection if it's not active.

Note that, if `c` calls Close function closing itself, `c` can not be used again.

##### (*PoolConn) Recv 

``` go
func (c *PoolConn) Recv(length int, retry ...Retry) ([]byte, error)
```

Recv receives data from the connection.

##### (*PoolConn) RecvLine 

``` go
func (c *PoolConn) RecvLine(retry ...Retry) ([]byte, error)
```

RecvLine reads data from the connection until reads char '\n'. Note that the returned result does not contain the last char '\n'.

##### (*PoolConn) RecvPkg 

``` go
func (c *PoolConn) RecvPkg(option ...PkgOption) ([]byte, error)
```

RecvPkg receives package from connection using simple package protocol. The optional parameter `option` specifies the package options for receiving.

##### (*PoolConn) RecvPkgWithTimeout 

``` go
func (c *PoolConn) RecvPkgWithTimeout(timeout time.Duration, option ...PkgOption) (data []byte, err error)
```

RecvPkgWithTimeout reads data from connection with timeout using simple package protocol.

##### (*PoolConn) RecvTill 

``` go
func (c *PoolConn) RecvTill(til []byte, retry ...Retry) ([]byte, error)
```

RecvTill reads data from the connection until reads bytes `til`. Note that the returned result contains the last bytes `til`.

##### (*PoolConn) RecvWithTimeout 

``` go
func (c *PoolConn) RecvWithTimeout(length int, timeout time.Duration, retry ...Retry) (data []byte, err error)
```

RecvWithTimeout reads data from the connection with timeout.

##### (*PoolConn) Send 

``` go
func (c *PoolConn) Send(data []byte, retry ...Retry) error
```

Send writes data to the connection. It retrieves a new connection from its pool if it fails writing data.

##### (*PoolConn) SendPkg 

``` go
func (c *PoolConn) SendPkg(data []byte, option ...PkgOption) (err error)
```

SendPkg sends a package containing `data` to the connection. The optional parameter `option` specifies the package options for sending.

##### (*PoolConn) SendPkgWithTimeout 

``` go
func (c *PoolConn) SendPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) (err error)
```

SendPkgWithTimeout writes data to connection with timeout using simple package protocol.

##### (*PoolConn) SendRecv 

``` go
func (c *PoolConn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error)
```

SendRecv writes data to the connection and blocks reading response.

##### (*PoolConn) SendRecvPkg 

``` go
func (c *PoolConn) SendRecvPkg(data []byte, option ...PkgOption) ([]byte, error)
```

SendRecvPkg writes data to connection and blocks reading response using simple package protocol.

##### (*PoolConn) SendRecvPkgWithTimeout 

``` go
func (c *PoolConn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error)
```

SendRecvPkgWithTimeout reads data from connection with timeout using simple package protocol.

##### (*PoolConn) SendRecvWithTimeout 

``` go
func (c *PoolConn) SendRecvWithTimeout(data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error)
```

SendRecvWithTimeout writes data to the connection and reads response with timeout.

##### (*PoolConn) SendWithTimeout 

``` go
func (c *PoolConn) SendWithTimeout(data []byte, timeout time.Duration, retry ...Retry) (err error)
```

SendWithTimeout writes data to the connection with timeout.

#### type Retry 

``` go
type Retry struct {
	Count    int           // Retry count.
	Interval time.Duration // Retry interval.
}
```

#### type Server 

``` go
type Server struct {
	// contains filtered or unexported fields
}
```

Server is a TCP server.

##### func GetServer 

``` go
func GetServer(name ...interface{}) *Server
```

GetServer returns the TCP server with specified `name`, or it returns a new normal TCP server named `name` if it does not exist. The parameter `name` is used to specify the TCP server

##### func NewServer 

``` go
func NewServer(address string, handler func(*Conn), name ...string) *Server
```

NewServer creates and returns a new normal TCP server. The parameter `name` is optional, which is used to specify the instance name of the server.

##### func NewServerKeyCrt 

``` go
func NewServerKeyCrt(address, crtFile, keyFile string, handler func(*Conn), name ...string) (*Server, error)
```

NewServerKeyCrt creates and returns a new TCP server with TLS support. The parameter `name` is optional, which is used to specify the instance name of the server.

##### func NewServerTLS 

``` go
func NewServerTLS(address string, tlsConfig *tls.Config, handler func(*Conn), name ...string) *Server
```

NewServerTLS creates and returns a new TCP server with TLS support. The parameter `name` is optional, which is used to specify the instance name of the server.

##### (*Server) Close 

``` go
func (s *Server) Close() error
```

Close closes the listener and shutdowns the server.

##### (*Server) GetAddress <-2.1.0

``` go
func (s *Server) GetAddress() string
```

GetAddress get the listening address for server.

##### (*Server) GetListenedAddress <-2.2.0

``` go
func (s *Server) GetListenedAddress() string
```

GetListenedAddress retrieves and returns the address string which are listened by current server.

##### (*Server) GetListenedPort <-2.2.0

``` go
func (s *Server) GetListenedPort() int
```

GetListenedPort retrieves and returns one port which is listened to by current server.

##### (*Server) Run 

``` go
func (s *Server) Run() (err error)
```

Run starts running the TCP Server.

##### (*Server) SetAddress 

``` go
func (s *Server) SetAddress(address string)
```

SetAddress sets the listening address for server.

##### (*Server) SetHandler 

``` go
func (s *Server) SetHandler(handler func(*Conn))
```

SetHandler sets the connection handler for server.

##### (*Server) SetTLSConfig 

``` go
func (s *Server) SetTLSConfig(tlsConfig *tls.Config)
```

SetTLSConfig sets the TLS configuration of server.

##### (*Server) SetTLSKeyCrt 

``` go
func (s *Server) SetTLSKeyCrt(crtFile, keyFile string) error
```

SetTLSKeyCrt sets the certificate and key file for TLS configuration of server.