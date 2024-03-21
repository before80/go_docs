+++
title = "gudp"
date = 2024-03-21T17:53:58+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gudp

Package gudp provides UDP server and client implementations.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/gudp/gudp_server.go#L21)

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
##### func MustGetFreePort <-2.2.0

``` go
func MustGetFreePort() (port int)
```

MustGetFreePort performs as GetFreePort, but it panics if any error occurs.

##### func NewNetConn 

``` go
func NewNetConn(remoteAddress string, localAddress ...string) (*net.UDPConn, error)
```

NewNetConn creates and returns a *net.UDPConn with given addresses.

##### func Send 

``` go
func Send(address string, data []byte, retry ...Retry) error
```

Send writes data to `address` using UDP connection and then closes the connection. Note that it is used for short connection usage.

##### func SendRecv 

``` go
func SendRecv(address string, data []byte, receive int, retry ...Retry) ([]byte, error)
```

SendRecv writes data to `address` using UDP connection, reads response and then closes the connection. Note that it is used for short connection usage.

### Types 

#### type Conn 

``` go
type Conn struct {
	*net.UDPConn // Underlying UDP connection.
	// contains filtered or unexported fields
}
```

Conn handles the UDP connection.

##### func NewConn 

``` go
func NewConn(remoteAddress string, localAddress ...string) (*Conn, error)
```

NewConn creates UDP connection to `remoteAddress`. The optional parameter `localAddress` specifies the local address for connection.

##### func NewConnByNetConn 

``` go
func NewConnByNetConn(udp *net.UDPConn) *Conn
```

NewConnByNetConn creates an UDP connection object with given *net.UDPConn object.

##### (*Conn) Recv 

``` go
func (c *Conn) Recv(buffer int, retry ...Retry) ([]byte, error)
```

Recv receives and returns data from remote address. The parameter `buffer` is used for customizing the receiving buffer size. If `buffer` <= 0, it uses the default buffer size, which is 1024 byte.

There's package border in UDP protocol, we can receive a complete package if specified buffer size is big enough. VERY NOTE that we should receive the complete package in once or else the leftover package data would be dropped.

##### (*Conn) RecvWithTimeout 

``` go
func (c *Conn) RecvWithTimeout(length int, timeout time.Duration, retry ...Retry) (data []byte, err error)
```

RecvWithTimeout reads data from remote address with timeout.

##### (*Conn) RemoteAddr 

``` go
func (c *Conn) RemoteAddr() net.Addr
```

RemoteAddr returns the remote address of current UDP connection. Note that it cannot use c.conn.RemoteAddr() as it is nil.

##### (*Conn) Send 

``` go
func (c *Conn) Send(data []byte, retry ...Retry) (err error)
```

Send writes data to remote address.

##### (*Conn) SendRecv 

``` go
func (c *Conn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error)
```

SendRecv writes data to connection and blocks reading response.

##### (*Conn) SendRecvWithTimeout 

``` go
func (c *Conn) SendRecvWithTimeout(data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error)
```

SendRecvWithTimeout writes data to connection and reads response with timeout.

##### (*Conn) SendWithTimeout 

``` go
func (c *Conn) SendWithTimeout(data []byte, timeout time.Duration, retry ...Retry) (err error)
```

SendWithTimeout writes data to connection with timeout.

##### (*Conn) SetBufferWaitRecv <-2.3.0

``` go
func (c *Conn) SetBufferWaitRecv(d time.Duration)
```

SetBufferWaitRecv sets the buffer waiting timeout when reading all data from connection. The waiting duration cannot be too long which might delay receiving data from remote address.

##### (*Conn) SetDeadline 

``` go
func (c *Conn) SetDeadline(t time.Time) (err error)
```

SetDeadline sets the read and write deadlines associated with the connection.

##### (*Conn) SetDeadlineRecv <-2.3.0

``` go
func (c *Conn) SetDeadlineRecv(t time.Time) (err error)
```

SetDeadlineRecv sets the read deadline associated with the connection.

##### (*Conn) SetDeadlineSend <-2.3.0

``` go
func (c *Conn) SetDeadlineSend(t time.Time) (err error)
```

SetDeadlineSend sets the deadline of sending for current connection.

#### type Retry 

``` go
type Retry struct {
	Count    int           // Max retry count.
	Interval time.Duration // Retry interval.
}
```

#### type Server 

``` go
type Server struct {
	// contains filtered or unexported fields
}
```

Server is the UDP server.

##### func GetServer 

``` go
func GetServer(name ...interface{}) *Server
```

GetServer creates and returns an UDP server instance with given name.

##### func NewServer 

``` go
func NewServer(address string, handler func(*Conn), name ...string) *Server
```

NewServer creates and returns an UDP server. The optional parameter `name` is used to specify its name, which can be used for GetServer function to retrieve its instance.

##### (*Server) Close 

``` go
func (s *Server) Close() (err error)
```

Close closes the connection. It will make server shutdowns immediately.

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
func (s *Server) Run() error
```

Run starts listening UDP connection.

##### (*Server) SetAddress 

``` go
func (s *Server) SetAddress(address string)
```

SetAddress sets the server address for UDP server.

##### (*Server) SetHandler 

``` go
func (s *Server) SetHandler(handler func(*Conn))
```

SetHandler sets the connection handler for UDP server.