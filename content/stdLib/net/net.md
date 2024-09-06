+++
title = "net"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net@go1.23.0](https://pkg.go.dev/net@go1.23.0)

Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.

​	`net`包提供了一个可移植的网络 I/O 接口，包括 TCP/IP、UDP、域名解析和 Unix 域套接字。

Although the package provides access to low-level networking primitives, most clients will need only the basic interface provided by the Dial, Listen, and Accept functions and the associated Conn and Listener interfaces. The crypto/tls package uses the same interfaces and similar Dial and Listen functions.

​	虽然该包提供了对低级网络原语的访问，但大多数客户端只需要 Dial、Listen 和 Accept 函数以及相关的 Conn 和 Listener 接口所提供的基本接口。crypto/tls 包使用相同的接口和类似的 Dial 和 Listen 函数。

The Dial function connects to a server:

​	`Dial` 函数连接到服务器：

```
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
	// 处理错误
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
// ...
```

The Listen function creates servers:

​	`Listen` 函数创建服务器：

```
ln, err := net.Listen("tcp", ":8080")
if err != nil {
	// handle error
}
for {
	conn, err := ln.Accept()
	if err != nil {
		// 处理错误
	}
	go handleConnection(conn)
}
```

## Name 解析  Name Resolution

The method for resolving domain names, whether indirectly with functions like Dial or directly with functions like LookupHost and LookupAddr, varies by operating system.

​	解析域名的方法(无论是间接使用 Dial 等函数还是直接使用 LookupHost 和 LookupAddr 等函数)因操作系统而异。

On Unix systems, the resolver has two options for resolving names. It can use a pure Go resolver that sends DNS requests directly to the servers listed in /etc/resolv.conf, or it can use a cgo-based resolver that calls C library routines such as getaddrinfo and getnameinfo.

​	在 Unix 系统上，解析器有两个选项来解析名称。它可以使用纯 Go 解析器，该解析器将 DNS 请求直接发送到 /etc/resolv.conf 中列出的服务器，或者它可以使用基于 cgo 的解析器，该解析器调用 C 库例程，例如 getaddrinfo 和 getnameinfo。

By default the pure Go resolver is used, because a blocked DNS request consumes only a goroutine, while a blocked C call consumes an operating system thread. When cgo is available, the cgo-based resolver is used instead under a variety of conditions: on systems that do not let programs make direct DNS requests (OS X), when the LOCALDOMAIN environment variable is present (even if empty), when the RES_OPTIONS or HOSTALIASES environment variable is non-empty, when the ASR_CONFIG environment variable is non-empty (OpenBSD only), when /etc/resolv.conf or /etc/nsswitch.conf specify the use of features that the Go resolver does not implement, and when the name being looked up ends in .local or is an mDNS name.

​	默认情况下，使用纯 Go 解析器，因为阻塞的 DNS 请求只会消耗一个 goroutine，而阻塞的 C 调用会消耗一个操作系统线程。当 cgo 可用时，在多种条件下将改为使用基于 cgo 的解析器：在不允许程序直接进行 DNS 请求的系统上(OS X)，当存在 LOCALDOMAIN 环境变量时(即使为空)，当 RES_OPTIONS 或 HOSTALIASES 环境变量非空时，当 ASR_CONFIG 环境变量非空时(仅在 OpenBSD 上)，当 /etc/resolv.conf 或 /etc/nsswitch.conf 指定使用 Go 解析器未实现的功能时，以及正在查找的名称以 .local 结尾或是 mDNS 名称时。

The resolver decision can be overridden by setting the netdns value of the GODEBUG environment variable (see package runtime) to go or cgo, as in:

​	可以通过将 GODEBUG 环境变量(参见 package runtime)的 netdns 值设置为 go 或 cgo 来覆盖解析器决策，例如：

```
export GODEBUG=netdns=go    # force pure Go resolver
export GODEBUG=netdns=cgo   # force native resolver (cgo, win32)
```

The decision can also be forced while building the Go source tree by setting the netgo or netcgo build tag.

​	在构建 Go 源代码树时，也可以通过设置 netgo 或 netcgo 构建标签来强制执行决策。

A numeric netdns setting, as in GODEBUG=netdns=1, causes the resolver to print debugging information about its decisions. To force a particular resolver while also printing debugging information, join the two settings by a plus sign, as in GODEBUG=netdns=go+1.

​	数字`netdns`设置，例如`GODEBUG=netdns=1`，会导致解析器打印有关其决策的调试信息。为了在打印调试信息的同时强制使用特定的解析器，请将两个设置连接起来，如`GODEBUG=netdns=go+1`。

On macOS, if Go code that uses the net package is built with -buildmode=c-archive, linking the resulting archive into a C program requires passing -lresolv when linking the C code.

On Plan 9, the resolver always accesses /net/cs and /net/dns.

​	在Plan 9上，解析器总是访问/net/cs和/net/dns。

On Windows, in Go 1.18.x and earlier, the resolver always used C library functions, such as GetAddrInfo and DnsQuery.

​	在早期的Go 1.18.x及之前的Windows版本中，解析器总是使用C库函数，例如`GetAddrInfo`和`DnsQuery`。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=21)

``` go 
const (
	IPv4len = 4
	IPv6len = 16
)
```

IP address lengths (bytes).

​	IP地址长度(以字节为单位)。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=100)

``` go 
var (
	IPv4bcast     = IPv4(255, 255, 255, 255) // 有限广播
	IPv4allsys    = IPv4(224, 0, 0, 1)       // 所有系统
	IPv4allrouter = IPv4(224, 0, 0, 2)       // 所有路由器
	IPv4zero      = IPv4(0, 0, 0, 0)         // 全部为零
)
```

Well-known IPv4 addresses

​	知名IPv4地址。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=108)

``` go 
var (
	IPv6zero                   = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6unspecified            = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6loopback               = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	IPv6interfacelocalallnodes = IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallnodes      = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallrouters    = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)
```

Well-known IPv6 addresses

​	知名IPv6地址。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=118)

``` go 
var DefaultResolver = &Resolver{}
```

DefaultResolver is the resolver used by the package-level Lookup functions and by Dialers without a specified Resolver.

​	`DefaultResolver` 是由包级别的 Lookup 函数和没有指定 Resolver 的 Dialer 使用的解析器。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=665)

``` go 
var ErrClosed error = errClosed
```

ErrClosed is the error returned by an I/O call on a network connection that has already been closed, or that is closed by another goroutine before the I/O is completed. This may be wrapped in another error, and should normally be tested using errors.Is(err, net.ErrClosed).

​	`ErrClosed` 是在网络连接上进行 I/O 调用时返回的错误，该连接已经关闭，或者在 I/O 完成之前被另一个 goroutine 关闭。这可能会包装在另一个错误中，并且通常应使用 errors.Is(err，net.ErrClosed)进行测试。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=408)

``` go 
var (
	ErrWriteToConnected = errors.New("use of WriteTo with pre-connected connection")
)
```

Various errors contained in OpError.

​	`OpError` 中包含的各种错误。

## 函数

### func JoinHostPort 

``` go 
func JoinHostPort(host, port string) string
```

JoinHostPort combines host and port into a network address of the form "host:port". If host contains a colon, as found in literal IPv6 addresses, then JoinHostPort returns "[host]:port".

​	`JoinHostPort`函数将 `host` 和 `port` 组合成形如 "`host:port`" 的网络地址。如果 host 包含冒号，如字面上的 IPv6 地址，则 JoinHostPort函数返回 "[host]:port"。

See func Dial for a description of the host and port parameters.

​	有关 `host` 和 `port` 参数的说明，请参见 func `Dial`。

### func LookupAddr 

``` go 
func LookupAddr(addr string) (names []string, err error)
```

LookupAddr performs a reverse lookup for the given address, returning a list of names mapping to that address.

​	LookupAddr函数根据给定的地址执行反向查找，返回映射到该地址的名称列表。

The returned names are validated to be properly formatted presentation-format domain names. If the response contains invalid names, those records are filtered out and an error will be returned alongside the remaining results, if any.

​	返回的名称已验证为格式正确的表示格式域名。如果响应包含无效的名称，则会过滤掉这些记录，并在剩余结果(如果有)旁边返回错误。

When using the host C library resolver, at most one result will be returned. To bypass the host resolver, use a custom Resolver.

​	使用主机 C 库解析程序时，最多只会返回一个结果。要绕过主机解析程序，请使用自定义解析程序。

LookupAddr uses context.Background internally; to specify the context, use Resolver.LookupAddr.

​	LookupAddr函数在内部使用 context.Background；要指定上下文，请使用 Resolver.LookupAddr。

### func LookupCNAME 

``` go 
func LookupCNAME(host string) (cname string, err error)
```

LookupCNAME returns the canonical name for the given host. Callers that do not care about the canonical name can call LookupHost or LookupIP directly; both take care of resolving the canonical name as part of the lookup.

​	LookupCNAME函数返回给定主机的规范名称。不关心规范名称的调用方可以直接调用 LookupHost函数或 LookupIP函数；两者都会在查找中解析规范名称。

A canonical name is the final name after following zero or more CNAME records. LookupCNAME does not return an error if host does not contain DNS "CNAME" records, as long as host resolves to address records.

​	规范名称是在跟随零个或多个 CNAME 记录之后的最终名称。如果主机不包含 DNS "CNAME" 记录，则 LookupCNAME 不会返回错误，只要主机解析到地址记录即可。

The returned canonical name is validated to be a properly formatted presentation-format domain name.

​	返回的规范名称已验证为格式正确的表示格式域名。

LookupCNAME uses context.Background internally; to specify the context, use Resolver.LookupCNAME.

​	LookupCNAME函数在内部使用 context.Background；要指定上下文，请使用 Resolver.LookupCNAME。

### func LookupHost 

``` go 
func LookupHost(host string) (addrs []string, err error)
```

LookupHost looks up the given host using the local resolver. It returns a slice of that host's addresses.

​	LookupHost函数使用本地解析器查找给定的主机名。它返回该主机的地址列表。

LookupHost uses context.Background internally; to specify the context, use Resolver.LookupHost.

​	LookupHost函数在内部使用context.Background；要指定上下文，请使用Resolver.LookupHost。

### func LookupPort 

``` go 
func LookupPort(network, service string) (port int, err error)
```

LookupPort looks up the port for the given network and service.

​	LookupPort函数查找给定网络和服务的端口。

LookupPort uses context.Background internally; to specify the context, use Resolver.LookupPort.

​	LookupPort函数在内部使用context.Background；要指定上下文，请使用Resolver.LookupPort。

### func LookupTXT 

``` go 
func LookupTXT(name string) ([]string, error)
```

LookupTXT returns the DNS TXT records for the given domain name.

​	LookupTXT函数返回给定域名的DNS TXT记录。

LookupTXT uses context.Background internally; to specify the context, use Resolver.LookupTXT.

​	LookupTXT函数在内部使用context.Background；要指定上下文，请使用Resolver.LookupTXT。

### func ParseCIDR 

``` go 
func ParseCIDR(s string) (IP, *IPNet, error)
```

ParseCIDR parses s as a CIDR notation IP address and prefix length, like "192.0.2.0/24" or "2001:db8::/32", as defined in [RFC 4632](https://rfc-editor.org/rfc/rfc4632.html) and [RFC 4291](https://rfc-editor.org/rfc/rfc4291.html).

​	ParseCIDR函数将s解析为CIDR表示法的IP地址和前缀长度，例如"192.0.2.0/24"或"2001:db8::/32"，如RFC 4632和RFC 4291中所定义。

It returns the IP address and the network implied by the IP and prefix length. For example, ParseCIDR("192.0.2.1/24") returns the IP address 192.0.2.1 and the network 192.0.2.0/24.

​	它返回IP地址和由IP和前缀长度隐含的网络。例如，ParseCIDR("192.0.2.1/24")返回IP地址192.0.2.1和网络192.0.2.0/24。

#### ParseCIDR Example
``` go 
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)

	ipv6Addr, ipv6Net, err := net.ParseCIDR("2001:db8:a0b:12f0::1/32")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv6Addr)
	fmt.Println(ipv6Net)

}
Output:

192.0.2.1
192.0.2.0/24
2001:db8:a0b:12f0::1
2001:db8::/32
```

### func Pipe 

``` go 
func Pipe() (Conn, Conn)
```

Pipe creates a synchronous, in-memory, full duplex network connection; both ends implement the Conn interface. Reads on one end are matched with writes on the other, copying data directly between the two; there is no internal buffering.

​	Pipe函数创建一个同步的内存中全双工网络连接；两端都实现了Conn接口。一个端口的读取与另一个端口的写入匹配，直接在两者之间复制数据；没有内部缓冲。

### func SplitHostPort 

``` go 
func SplitHostPort(hostport string) (host, port string, err error)
```

SplitHostPort splits a network address of the form "host:port", "host%zone:port", "[host]:port" or "[host%zone]:port" into host or host%zone and port.

​	SplitHostPort函数将"`host:port`"、"`host%zone:port`"、"`[host]:port`"或"`[host％zone]:port`"形式的网络地址分成主机或host％zone和端口。

A literal IPv6 address in hostport must be enclosed in square brackets, as in "[::1]:80", "[::1%lo0]:80".

​	hostport中的文字IPv6地址必须用方括号括起来，如"`[::1]:80`"、"`[::1％lo0]:80`"。

See func Dial for a description of the hostport parameter, and host and port results.

​	有关hostport参数和host和port结果的说明，请参见func Dial。

## 类型

### type Addr 

``` go 
type Addr interface {
	Network() string // 网络的名称(例如 "tcp"、"udp")
	String() string  // 地址的字符串表示形式(例如 "192.0.2.1:25"、"[2001:db8::1]:80")
}
```

Addr represents a network end point address.

​	`Addr`接口表示网络终端地址。

The two methods Network and String conventionally return strings that can be passed as the arguments to Dial, but the exact form and meaning of the strings is up to the implementation.

​	`Network` 和 `String` 两个方法一般返回可以作为 `Dial` 函数参数的字符串，但字符串的确切格式和含义取决于实现。

#### func InterfaceAddrs 

``` go 
func InterfaceAddrs() ([]Addr, error)
```

InterfaceAddrs returns a list of the system's unicast interface addresses.

​	InterfaceAddrs函数返回系统的单播接口地址列表。

The returned list does not identify the associated interface; use Interfaces and Interface.Addrs for more detail.

​	返回的列表不包含关联接口的标识，使用 Interfaces 和 Interface.Addrs 获取更多详细信息。

### type AddrError 

``` go 
type AddrError struct {
	Err  string
	Addr string
}
```

#### (*AddrError) Error 

``` go 
func (e *AddrError) Error() string
```

#### (*AddrError) Temporary 

``` go 
func (e *AddrError) Temporary() bool
```

#### (*AddrError) Timeout 

``` go 
func (e *AddrError) Timeout() bool
```

### type Buffers  <- go1.8

``` go 
type Buffers [][]byte
```

Buffers contains zero or more runs of bytes to write.

​	Buffers 包含零个或多个要写入的字节序列。

On certain machines, for certain types of connections, this is optimized into an OS-specific batch write operation (such as "writev").

​	在某些机器上，对于某些类型的连接，这可以被优化为操作系统特定的批量写操作(例如 "writev")。

#### (*Buffers) Read  <- go1.8

``` go 
func (v *Buffers) Read(p []byte) (n int, err error)
```

Read from the buffers.

​	Read方法从缓冲区读取。

Read implements io.Reader for Buffers.

​	Read方法为 Buffers 实现了 io.Reader。

Read modifies the slice v as well as v[i] for 0 <= i < len(v), but does not modify v[i][j] for any i, j.

​	Read方法修改了切片 v 和 v[i]，其中 0 <= i < len(v)，但不修改 v[i][j]，其中 i 和 j 为任意值。

#### (*Buffers) WriteTo  <- go1.8

``` go 
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error)
```

WriteTo writes contents of the buffers to w.

​	`WriteTo`方法将缓冲区的内容写入 `w`。

WriteTo implements io.WriterTo for Buffers.

​	`WriteTo`方法为 Buffers 实现了 `io.WriterTo`。

WriteTo modifies the slice v as well as v[i] for 0 <= i < len(v), but does not modify v[i][j] for any i, j.

​	`WriteTo`方法修改了切片 `v` 和 `v[i]`，其中 `0 <= i < len(v)`，但不修改 `v[i][j]`，其中 i 和 `j` 为任意值。

### type Conn 

``` go 
type Conn interface {
    // Read reads data from the connection.
	// Read can be made to time out and return an error after a fixed
	// time limit; see SetDeadline and SetReadDeadline.
	// Read 从连接中读取数据。
	// Read 可以设置超时时间，在固定时间限制后超时并返回错误；
    // 参见 SetDeadline 和 	SetReadDeadline。
	Read(b []byte) (n int, err error)

    // Write writes data to the connection.
	// Write can be made to time out and return an error after a fixed
	// time limit; see SetDeadline and SetWriteDeadline.
	// Write 向连接中写入数据。
	// Write 可以设置超时时间，在固定时间限制后超时并返回错误；
    // 参见 SetDeadline 和 SetWriteDeadline。
	Write(b []byte) (n int, err error)

    // Close closes the connection.
	// Any blocked Read or Write operations will be unblocked and return errors.
	// Close 关闭连接。
	// 任何被阻塞的 Read 或 Write 操作都将取消阻塞并返回错误。
	Close() error

    // LocalAddr returns the local network address, if known.
	// LocalAddr 返回本地网络地址(如果已知)。
	LocalAddr() Addr

    // RemoteAddr returns the remote network address, if known.
	// RemoteAddr 返回远程网络地址(如果已知)。
	RemoteAddr() Addr

    // SetDeadline sets the read and write deadlines associated
	// with the connection. It is equivalent to calling both
	// SetReadDeadline and SetWriteDeadline.
	// SetDeadline 设置与连接关联的读取和写入截止时间。
    // 它等效于调用 SetReadDeadline 和 SetWriteDeadline 两个方法。
	//
    // A deadline is an absolute time after which I/O operations
	// fail instead of blocking. The deadline applies to all future
	// and pending I/O, not just the immediately following call to
	// Read or Write. After a deadline has been exceeded, the
	// connection can be refreshed by setting a deadline in the future.
	// 截止时间是一个绝对时间，在此之后 I/O 操作将失败而不是阻塞。
	// 截止时间适用于所有未来和挂起的 I/O，而不仅仅是立即跟随 Read 或 Write 调用的操作。
	// 如果超过截止时间，可以通过将截止时间设置为未来来刷新连接。
	//
    // If the deadline is exceeded a call to Read or Write or to other
	// I/O methods will return an error that wraps os.ErrDeadlineExceeded.
	// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
	// The error's Timeout method will return true, but note that there
	// are other possible errors for which the Timeout method will
	// return true even if the deadline has not been exceeded.
	// 如果超过截止时间，
    // 对 Read 或 Write 或其他 I/O 方法的调用将返回包含 os.ErrDeadlineExceeded 的错误。
	// 可以使用 errors.Is(err, os.ErrDeadlineExceeded) 测试此错误。
	// 错误的 Timeout 方法将返回 true，但请注意，
    // 即使截止时间尚未超过，也有其他可能导致 Timeout 方法返回 true 的错误。
	//
    // A zero value for t means I/O operations will not time out.
	// 空值 t 表示 I/O 操作将不会超时。
	SetDeadline(t time.Time) error

    // SetReadDeadline sets the deadline for future Read calls
	// and any currently-blocked Read call.
	// A zero value for t means Read will not time out.
	// SetReadDeadline 设置未来 Read 调用和任何当前被阻塞的 Read 调用的截止时间。
	// 空值 t 表示 Read 不会超时。
	SetReadDeadline(t time.Time) error

    // SetWriteDeadline sets the deadline for future Write calls
	// and any currently-blocked Write call.
	// Even if write times out, it may return n > 0, indicating that
	// some of the data was successfully written.
	// A zero value for t means Write will not time out.
	// SetWriteDeadline 设置未来 Write 调用和任何当前被阻塞的 Write 调用的截止时间。
	// 即使写入超时，它也可能返回 n > 0，表示某些数据已成功写入。
	// 空值 t 表示 Write 不会超时。
	SetWriteDeadline(t time.Time) error
}
```

Conn is a generic stream-oriented network connection.

​	`Conn` 是一个通用的面向流的网络连接。

Multiple goroutines may invoke methods on a Conn simultaneously.

​	多个 goroutine 可以同时调用 `Conn` 上的方法。

#### func Dial 

``` go 
func Dial(network, address string) (Conn, error)
```

Dial connects to the address on the named network.

​	`Dial`函数连接到指定网络上的地址。

Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket".

​	已知的网络有 "tcp"、"tcp4"(仅限 IPv4)、"tcp6"(仅限 IPv6)、"udp"、"udp4"(仅限 IPv4)、"udp6"(仅限 IPv6)、"ip"、"ip4"(仅限 IPv4)、"ip6"(仅限 IPv6)、"unix"、"unixgram" 和 "unixpacket"。

For TCP and UDP networks, the address has the form "host:port". The host must be a literal IP address, or a host name that can be resolved to IP addresses. The port must be a literal port number or a service name. If the host is a literal IPv6 address it must be enclosed in square brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80". The zone specifies the scope of the literal IPv6 address as defined in [RFC 4007](https://rfc-editor.org/rfc/rfc4007.html). The functions JoinHostPort and SplitHostPort manipulate a pair of host and port in this form. When using TCP, and the host resolves to multiple IP addresses, Dial will try each IP address in order until one succeeds.

​	对于 TCP 和 UDP 网络，地址的格式为 "host:port"。主机必须是一个字面 IP 地址，或者是可以解析为 IP 地址的主机名。端口必须是字面端口号或服务名称。如果主机是字面 IPv6 地址，则必须用方括号括起来，如 "[2001:db8::1]:80" 或 "[fe80::1%zone]:80"。区域指定了字面 IPv6 地址的作用域，如 RFC 4007 中定义的那样。函数 JoinHostPort 和 SplitHostPort 以这种形式操作一对主机和端口。在使用 TCP 时，如果主机解析为多个 IP 地址，则 Dial 将按顺序尝试每个 IP 地址，直到其中一个成功。

Examples:

​	示例：

```
Dial("tcp", "golang.org:http")
Dial("tcp", "192.0.2.1:http")
Dial("tcp", "198.51.100.1:80")
Dial("udp", "[2001:db8::1]:domain")
Dial("udp", "[fe80::1%lo0]:53")
Dial("tcp", ":80")
```

For IP networks, the network must be "ip", "ip4" or "ip6" followed by a colon and a literal protocol number or a protocol name, and the address has the form "host". The host must be a literal IP address or a literal IPv6 address with zone. It depends on each operating system how the operating system behaves with a non-well known protocol number such as "0" or "255".

​	对于 IP 网络，网络必须是 "ip"、"ip4" 或 "ip6"，后跟冒号和字面协议号或协议名称，地址的格式为 "host"。主机必须是一个字面 IP 地址或带区域的字面 IPv6 地址。每个操作系统如何处理非知名协议号(如 "0" 或 "255")取决于操作系统。

Examples:

​	示例：

```
Dial("ip4:1", "192.0.2.1")
Dial("ip6:ipv6-icmp", "2001:db8::1")
Dial("ip6:58", "fe80::1%lo0")
```

For TCP, UDP and IP networks, if the host is empty or a literal unspecified IP address, as in ":80", "0.0.0.0:80" or "[::]:80" for TCP and UDP, "", "0.0.0.0" or "::" for IP, the local system is assumed.

​	对于 TCP、UDP 和 IP 网络，如果主机为空或为未指定的字面 IP 地址，如 ":80"、"0.0.0.0:80" 或 "[::]:80" 用于 TCP 和 UDP、""、"0.0.0.0" 或 "::" 用于 IP，则假定本地系统。

For Unix networks, the address must be a file system path.

​	对于 Unix 网络，地址必须是一个文件系统路径。

#### func DialTimeout 

``` go 
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
```

DialTimeout acts like Dial but takes a timeout.

​	DialTimeout函数类似于Dial，但带有超时设置。

The timeout includes name resolution, if required. When using TCP, and the host in the address parameter resolves to multiple IP addresses, the timeout is spread over each consecutive dial, such that each is given an appropriate fraction of the time to connect.

​	超时包括名称解析(如果需要)。当使用 TCP 时，如果地址参数中的主机名解析为多个 IP 地址，则超时时间将分布在每个连续的 dial 上，以使每个 dial 分配适当的时间比例进行连接。

See func Dial for a description of the network and address parameters.

​	有关网络和地址参数的说明，请参见 func Dial。

#### func FileConn 

``` go 
func FileConn(f *os.File) (c Conn, err error)
```

FileConn returns a copy of the network connection corresponding to the open file f. It is the caller's responsibility to close f when finished. Closing c does not affect f, and closing f does not affect c.

​	FileConn函数返回与打开文件 f 对应的网络连接的副本。当结束时，调用者负责关闭 f。关闭 c 不会影响 f，关闭 f 也不会影响 c。

### type DNSConfigError 

``` go 
type DNSConfigError struct {
	Err error
}
```

DNSConfigError represents an error reading the machine's DNS configuration. (No longer used; kept for compatibility.)

​	DNSConfigError 表示读取计算机 DNS 配置的错误。(不再使用；为了向后兼容而保留。)

#### (*DNSConfigError) Error 

``` go 
func (e *DNSConfigError) Error() string
```

#### (*DNSConfigError) Temporary 

``` go 
func (e *DNSConfigError) Temporary() bool
```

#### (*DNSConfigError) Timeout 

``` go 
func (e *DNSConfigError) Timeout() bool
```

#### (*DNSConfigError) Unwrap  <- go1.13

``` go 
func (e *DNSConfigError) Unwrap() error
```

### type DNSError 

``` go 
type DNSError struct {
	Err         string // 错误描述 description of the error
	Name        string // 查找的名称 name looked for
	Server      string // 使用的服务器 server used
	IsTimeout   bool   // 如果为真，则已超时；并非所有超时都设置了此项 if true, timed out; not all timeouts set this
	IsTemporary bool   // 如果为真，则错误是暂时的；并非所有错误都设置了此项 if true, error is temporary; not all errors set this
	IsNotFound  bool   // 如果为真，则找不到主机 if true, host could not be found
}
```

DNSError represents a DNS lookup error.

​	DNSError 表示 DNS 查找错误。

#### (*DNSError) Error 

``` go 
func (e *DNSError) Error() string
```

#### (*DNSError) Temporary 

``` go 
func (e *DNSError) Temporary() bool
```

Temporary reports whether the DNS error is known to be temporary. This is not always known; a DNS lookup may fail due to a temporary error and return a DNSError for which Temporary returns false.

​	Temporary方法报告DNS错误是否已知为临时错误。这并非总是已知的；DNS查找可能由于临时错误而失败，并返回一个DNSError，其中Temporary返回false。

#### (*DNSError) Timeout 

``` go 
func (e *DNSError) Timeout() bool
```

Timeout reports whether the DNS lookup is known to have timed out. This is not always known; a DNS lookup may fail due to a timeout and return a DNSError for which Timeout returns false.

​	Timeout方法报告DNS查找是否已知已超时。这并非总是已知的；DNS查找可能由于超时而失败，并返回一个DNSError，其中Timeout返回false。

### type Dialer  <- go1.1

``` go 
type Dialer struct {
    // Timeout is the maximum amount of time a dial will wait for
	// a connect to complete. If Deadline is also set, it may fail
	// earlier.
	// Timeout是拨号等待连接完成的最大时间。
    // 如果Deadline也设置了，可能会更早失败。
	//
    // The default is no timeout.
	// 默认情况下没有超时。
	//
    // When using TCP and dialing a host name with multiple IP
	// addresses, the timeout may be divided between them.
	// 在使用TCP并拨打具有多个IP地址的主机名时，
    // 超时可能会在它们之间分配。
	//
    // With or without a timeout, the operating system may impose
	// its own earlier timeout. For instance, TCP timeouts are
	// often around 3 minutes.
	// 有或没有超时，操作系统可能会强制执行自己的较早超时。
    // 例如，TCP超时通常为约3分钟。
	Timeout time.Duration

    // Deadline is the absolute point in time after which dials
	// will fail. If Timeout is set, it may fail earlier.
	// Zero means no deadline, or dependent on the operating system
	// as with the Timeout option.
	// Deadline是绝对时间点，超过这个时间点，拨号将失败。
    // 如果设置了Timeout，它可能会更早失败。
	// 零表示没有截止日期，或者依赖于操作系统，如Timeout选项。
	Deadline time.Time

    // LocalAddr is the local address to use when dialing an
	// address. The address must be of a compatible type for the
	// network being dialed.
	// If nil, a local address is automatically chosen.
	// LocalAddr是在拨打地址时使用的本地地址。
    // 地址必须是适用于被拨打的网络的兼容类型。
	// 如果为nil，则自动选择本地地址。
	LocalAddr Addr

    // DualStack previously enabled RFC 6555 Fast Fallback
	// support, also known as "Happy Eyeballs", in which IPv4 is
	// tried soon if IPv6 appears to be misconfigured and
	// hanging.
	// DualStack先前启用了RFC 6555快速回退支持，
    // 也称为" Happy Eyeballs"，
    // 在其中如果IPv6似乎配置不正确且挂起，则很快尝试IPv4。
	//
    // Deprecated: Fast Fallback is enabled by default. To
	// disable, set FallbackDelay to a negative value.
	// 已弃用：默认情况下启用了快速回退。
    // 要禁用，请将FallbackDelay设置为负值。
	DualStack bool

    // FallbackDelay specifies the length of time to wait before
	// spawning a RFC 6555 Fast Fallback connection. That is, this
	// is the amount of time to wait for IPv6 to succeed before
	// assuming that IPv6 is misconfigured and falling back to
	// IPv4.
	// FallbackDelay指定等待RFC 6555快速回退连接生成的时间长度。
    // 也就是说，在假定IPv6配置不正确并回退到IPv4之前，
    // 等待IPv6成功的时间量。
	//
    // If zero, a default delay of 300ms is used.
	// A negative value disables Fast Fallback support.
	// 如果为零，则使用默认延迟300毫秒。
	// 负值禁用快速回退支持。
	FallbackDelay time.Duration

    // KeepAlive specifies the interval between keep-alive
	// probes for an active network connection.
	// If zero, keep-alive probes are sent with a default value
	// (currently 15 seconds), if supported by the protocol and operating
	// system. Network protocols or operating systems that do
	// not support keep-alives ignore this field.
	// If negative, keep-alive probes are disabled.
	// KeepAlive指定活动网络连接的保持活动探测之间的间隔时间。
	// 如果为零，并且协议和操作系统支持，
    // 则使用默认值发送保持活动探测(当前为15秒)。
    // 不支持保持活动的网络协议或操作系统会忽略此字段。
	// 如果为负值，则禁用保持活动探测。
	KeepAlive time.Duration

    // Resolver optionally specifies an alternate resolver to use.
	// Resolver是可选的，它指定要使用的替代解析程序。
	Resolver *Resolver

    // Cancel is an optional channel whose closure indicates that
	// the dial should be canceled. Not all types of dials support
	// cancellation.
	// Cancel是一个可选的通道，其关闭指示应取消拨号。
    // 不是所有类型的拨号都支持取消。
	//
    // Deprecated: Use DialContext instead.
	// 已弃用：改用DialContext。
	Cancel <-chan struct{}

    // If Control is not nil, it is called after creating the network
	// connection but before actually dialing.
	// 如果Control不为nil，则在创建网络连接但尚未拨打时调用它。
	//
    // Network and address parameters passed to Control function are not
	// necessarily the ones passed to Dial. For example, passing "tcp" to Dial
	// will cause the Control function to be called with "tcp4" or "tcp6".
	// 传递给Control方法的网络和地址参数不一定是传递给Dial的参数。
    // 例如，
    // 将"tcp"传递给Dial将导致使用"tcp4"或"tcp6"调用Control函数。
	//
    // Control is ignored if ControlContext is not nil.
	// 如果ControlContext不为nil，则忽略Control。
	Control func(network, address string, c syscall.RawConn) error

    // If ControlContext is not nil, it is called after creating the network
	// connection but before actually dialing.
	// 如果ControlContext不为nil，
    // 在实际拨号之前创建网络连接时会调用它。
	//
    // Network and address parameters passed to ControlContext function are not
	// necessarily the ones passed to Dial. For example, passing "tcp" to Dial
	// will cause the ControlContext function to be called with "tcp4" or "tcp6".
	// 传递给Control方法的网络和地址参数不一定是传递给Dial的参数。
    // 例如，
    // 向Dial传递"tcp"将导致Control函数被调用时传递"tcp4" 或 "tcp6"。
	//
    // If ControlContext is not nil, Control is ignored.
	// 如果ControlContext不为nil，则忽略Control。
	ControlContext func(ctx context.Context, network, address string, c syscall.RawConn) error
}
```

A Dialer contains options for connecting to an address.

​	Dialer包含用于连接到地址的选项。

The zero value for each field is equivalent to dialing without that option. Dialing with the zero value of Dialer is therefore equivalent to just calling the Dial function.

​	每个字段的零值等效于不使用该选项进行拨号。因此，使用Dialer的零值进行拨号等效于只调用Dial函数。

It is safe to call Dialer's methods concurrently.

​	并发调用Dialer的方法是安全的。

#### Example
``` go 
package main

import (
	"context"
	"log"
	"net"
	"time"
)

func main() {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", "localhost:12345")
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, World!")); err != nil {
		log.Fatal(err)
	}
}

```

#### Example(Unix)
``` go 
package main

import (
	"context"
	"log"
	"net"
	"time"
)

func main() {
	// DialUnix does not take a context.Context parameter. This example shows
	// how to dial a Unix socket with a Context. Note that the Context only
	// applies to the dial operation; it does not apply to the connection once
	// it has been established.
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	d.LocalAddr = nil // if you have a local addr, add it here
	raddr := net.UnixAddr{Name: "/path/to/unix.sock", Net: "unix"}
	conn, err := d.DialContext(ctx, "unix", raddr.String())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	if _, err := conn.Write([]byte("Hello, socket!")); err != nil {
		log.Fatal(err)
	}
}

```

#### (*Dialer) Dial  <- go1.1

``` go 
func (d *Dialer) Dial(network, address string) (Conn, error)
```

Dial connects to the address on the named network.

​	Dial方法在指定网络上连接到地址。

See func Dial for a description of the network and address parameters.

​	有关网络和地址参数的描述，请参见func Dial。

Dial uses context.Background internally; to specify the context, use DialContext.

​	Dial在内部使用context.Background；要指定上下文，请使用DialContext。

#### (*Dialer) DialContext  <- go1.7

``` go 
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)
```

DialContext connects to the address on the named network using the provided context.

​	DialContext方法在指定网络上使用提供的上下文连接到地址。

The provided Context must be non-nil. If the context expires before the connection is complete, an error is returned. Once successfully connected, any expiration of the context will not affect the connection.

​	提供的上下文必须是非nil的。如果上下文在连接完成之前到期，则会返回错误。一旦成功连接，上下文的任何到期都不会影响连接。

When using TCP, and the host in the address parameter resolves to multiple network addresses, any dial timeout (from d.Timeout or ctx) is spread over each consecutive dial, such that each is given an appropriate fraction of the time to connect. For example, if a host has 4 IP addresses and the timeout is 1 minute, the connect to each single address will be given 15 seconds to complete before trying the next one.

​	使用TCP时，如果地址参数中的主机解析为多个网络地址，则任何拨号超时(来自d.Timeout或ctx)将分布在每个连续的拨号之间，以便每个拨号在适当的时间内完成。例如，如果主机有4个IP地址，并且超时时间为1分钟，则在尝试下一个地址之前，将为每个单个地址分配15秒钟的时间。

See func Dial for a description of the network and address parameters.

​	有关网络和地址参数的描述，请参见func Dial。

#### (*Dialer) MultipathTCP <-go1.21.0

``` go
func (d *Dialer) MultipathTCP() bool
```

MultipathTCP reports whether MPTCP will be used.

This method doesn't check if MPTCP is supported by the operating system or not.

#### (*Dialer) SetMultipathTCP <-go1.21.0

``` go
func (d *Dialer) SetMultipathTCP(use bool)
```

SetMultipathTCP directs the Dial methods to use, or not use, MPTCP, if supported by the operating system. This method overrides the system default and the GODEBUG=multipathtcp=... setting if any.

If MPTCP is not available on the host or not supported by the server, the Dial methods will fall back to TCP.



### type Error 

``` go 
type Error interface {
	error
	Timeout() bool // 错误是否超时？ Is the error a timeout?

    // Deprecated: Temporary errors are not well-defined.
	// Most "temporary" errors are timeouts, and the few exceptions are surprising.
	// Do not use this method.
	// Deprecated: 临时错误没有明确定义。
	// 大多数"临时"错误是超时错误，极少数例外情况是出人意料的。
	// 不要使用这个方法。
	Temporary() bool
}
```

An Error represents a network error.

​	`Error`接口表示网络错误。

### type Flags 

``` go 
type Flags uint
const (
	FlagUp           Flags = 1 << iota // 接口已被管理员开启  interface is administratively up
	FlagBroadcast                      // 接口支持广播访问能力 interface supports broadcast access capability
	FlagLoopback                       // 接口是一个回环接口 interface is a loopback interface
	FlagPointToPoint                   // 接口属于点对点连接 interface belongs to a point-to-point link
	FlagMulticast                      // 接口支持组播访问能力 interface supports multicast access capability
	FlagRunning                        // 接口处于运行状态 interface is in running state
)
```

#### (Flags) String 

``` go 
func (f Flags) String() string
```

### type HardwareAddr 

``` go 
type HardwareAddr []byte
```

A HardwareAddr represents a physical hardware address.

​	HardwareAddr表示一个物理硬件地址。

#### func ParseMAC 

``` go 
func ParseMAC(s string) (hw HardwareAddr, err error)
```

ParseMAC parses s as an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet IP over InfiniBand link-layer address using one of the following formats:

​	ParseMAC函数将s解析为IEEE 802 MAC-48、EUI-48、EUI-64或20字节IP over InfiniBand链路层地址之一，使用以下格式之一：

```
00:00:5e:00:53:01
02:00:5e:10:00:00:00:01
00:00:00:00:fe:80:00:00:00:00:00:00:02:00:5e:10:00:00:00:01
00-00-5e-00-53-01
02-00-5e-10-00-00-00-01
00-00-00-00-fe-80-00-00-00-00-00-00-02-00-5e-10-00-00-00-01
0000.5e00.5301
0200.5e10.0000.0001
0000.0000.fe80.0000.0000.0000.0200.5e10.0000.0001
```

#### (HardwareAddr) String 

``` go 
func (a HardwareAddr) String() string
```

### type IP 

``` go 
type IP []byte
```

An IP is a single IP address, a slice of bytes. Functions in this package accept either 4-byte (IPv4) or 16-byte (IPv6) slices as input.

​	IP是单个IP地址，是字节片。此包中的函数接受4字节(IPv4)或16字节(IPv6)片作为输入。

Note that in this documentation, referring to an IP address as an IPv4 address or an IPv6 address is a semantic property of the address, not just the length of the byte slice: a 16-byte slice can still be an IPv4 address.

​	请注意，在本文档中，将IP地址称为IPv4地址或IPv6地址是地址的语义属性，而不仅仅是字节片的长度：16字节的片仍然可以是IPv4地址。

#### IP Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6 := net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ipv4 := net.IPv4(10, 255, 0, 0)

	fmt.Println(ipv6.To4())
	fmt.Println(ipv4.To4())

}
Output:

<nil>
10.255.0.0
```

#### func IPv4 

``` go 
func IPv4(a, b, c, d byte) IP
```

IPv4 returns the IP address (in 16-byte form) of the IPv4 address a.b.c.d.

​	IPv4函数返回IPv4地址a.b.c.d的IP地址(以16字节形式)。

##### IPv4 Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.IPv4(8, 8, 8, 8))

}
Output:

8.8.8.8
```

#### func LookupIP 

``` go 
func LookupIP(host string) ([]IP, error)
```

LookupIP looks up host using the local resolver. It returns a slice of that host's IPv4 and IPv6 addresses.

​	`LookupIP`函数使用本地解析器查找主机，返回该主机的IPv4和IPv6地址切片。

#### func ParseIP 

``` go 
func ParseIP(s string) IP
```

ParseIP parses s as an IP address, returning the result. The string s can be in IPv4 dotted decimal ("192.0.2.1"), IPv6 ("2001:db8::68"), or IPv4-mapped IPv6 ("::ffff:192.0.2.1") form. If s is not a valid textual representation of an IP address, ParseIP returns nil.

​	`ParseIP`函数将 `s` 解析为一个 IP 地址并返回结果。字符串 `s` 可以是 IPv4 的点分十进制表示("192.0.2.1")、IPv6 的十六进制表示("`2001:db8::68`")或 IPv4 映射的 IPv6 表示形式("`::ffff:192.0.2.1`")。如果 `s` 不是有效的文本表示形式的 IP 地址，则 `ParseIP` 返回 nil。

##### ParseIP Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.ParseIP("192.0.2.1"))
	fmt.Println(net.ParseIP("2001:db8::68"))
	fmt.Println(net.ParseIP("192.0.2"))

}
Output:

192.0.2.1
2001:db8::68
<nil>
```

#### (IP) DefaultMask 

``` go 
func (ip IP) DefaultMask() IPMask
```

DefaultMask returns the default IP mask for the IP address ip. Only IPv4 addresses have default masks; DefaultMask returns nil if ip is not a valid IPv4 address.

​	`DefaultMask`方法返回 IP 地址 ip 的默认 IP 掩码。只有 IPv4 地址有默认掩码；如果 ip 不是有效的 IPv4 地址，则 DefaultMask 返回 nil。

##### DefaultMask Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("192.0.2.1")
	fmt.Println(ip.DefaultMask())

}
Output:

ffffff00
```

#### (IP) Equal 

``` go 
func (ip IP) Equal(x IP) bool
```

Equal reports whether ip and x are the same IP address. An IPv4 address and that same address in IPv6 form are considered to be equal.

​	`Equal`方法报告 ip 和 x 是否是相同的 IP 地址。IPv4 地址和相同的 IPv6 地址被认为是相等的。

##### Equal Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv4DNS := net.ParseIP("8.8.8.8")
	ipv4Lo := net.ParseIP("127.0.0.1")
	ipv6DNS := net.ParseIP("0:0:0:0:0:FFFF:0808:0808")

	fmt.Println(ipv4DNS.Equal(ipv4DNS))
	fmt.Println(ipv4DNS.Equal(ipv4Lo))
	fmt.Println(ipv4DNS.Equal(ipv6DNS))

}
Output:

true
false
true
```

#### (IP) IsGlobalUnicast 

``` go 
func (ip IP) IsGlobalUnicast() bool
```

IsGlobalUnicast reports whether ip is a global unicast address.

​	IsGlobalUnicast方法报告 ip 是否是全局单播地址。

The identification of global unicast addresses uses address type identification as defined in [RFC 1122](https://rfc-editor.org/rfc/rfc1122.html), [RFC 4632](https://rfc-editor.org/rfc/rfc4632.html) and [RFC 4291](https://rfc-editor.org/rfc/rfc4291.html) with the exception of IPv4 directed broadcast addresses. It returns true even if ip is in IPv4 private address space or local IPv6 unicast address space.

​	全局单播地址的标识使用 [RFC 1122](https://rfc-editor.org/rfc/rfc1122.html)、[RFC 4632](https://rfc-editor.org/rfc/rfc4632.html) 和 [RFC 4291](https://rfc-editor.org/rfc/rfc4291.html) 中定义的地址类型标识，但 IPv4 定向广播地址除外。即使 ip 在 IPv4 私有地址空间或本地 IPv6 单播地址空间中，它也返回 true。

##### IsGlobalUnicast Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6Global := net.ParseIP("2000::")
	ipv6UniqLocal := net.ParseIP("2000::")
	ipv6Multi := net.ParseIP("FF00::")

	ipv4Private := net.ParseIP("10.255.0.0")
	ipv4Public := net.ParseIP("8.8.8.8")
	ipv4Broadcast := net.ParseIP("255.255.255.255")

	fmt.Println(ipv6Global.IsGlobalUnicast())
	fmt.Println(ipv6UniqLocal.IsGlobalUnicast())
	fmt.Println(ipv6Multi.IsGlobalUnicast())

	fmt.Println(ipv4Private.IsGlobalUnicast())
	fmt.Println(ipv4Public.IsGlobalUnicast())
	fmt.Println(ipv4Broadcast.IsGlobalUnicast())

}
Output:

true
true
false
true
true
false
```

#### (IP) IsInterfaceLocalMulticast 

``` go 
func (ip IP) IsInterfaceLocalMulticast() bool
```

IsInterfaceLocalMulticast reports whether ip is an interface-local multicast address.

​	`IsInterfaceLocalMulticast`方法报告 ip 是否是接口本地组播地址。

##### IsInterfaceLocalMulticast Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6InterfaceLocalMulti := net.ParseIP("ff01::1")
	ipv6Global := net.ParseIP("2000::")
	ipv4 := net.ParseIP("255.0.0.0")

	fmt.Println(ipv6InterfaceLocalMulti.IsInterfaceLocalMulticast())
	fmt.Println(ipv6Global.IsInterfaceLocalMulticast())
	fmt.Println(ipv4.IsInterfaceLocalMulticast())

}
Output:

true
false
false
```

#### (IP) IsLinkLocalMulticast 

``` go 
func (ip IP) IsLinkLocalMulticast() bool
```

IsLinkLocalMulticast reports whether ip is a link-local multicast address.

​	IsLinkLocalMulticast方法报告 ip 是否是链路本地组播地址。

##### IsLinkLocalMulticast Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6LinkLocalMulti := net.ParseIP("ff02::2")
	ipv6LinkLocalUni := net.ParseIP("fe80::")
	ipv4LinkLocalMulti := net.ParseIP("224.0.0.0")
	ipv4LinkLocalUni := net.ParseIP("169.254.0.0")

	fmt.Println(ipv6LinkLocalMulti.IsLinkLocalMulticast())
	fmt.Println(ipv6LinkLocalUni.IsLinkLocalMulticast())
	fmt.Println(ipv4LinkLocalMulti.IsLinkLocalMulticast())
	fmt.Println(ipv4LinkLocalUni.IsLinkLocalMulticast())

}
Output:

true
false
true
false
```

#### (IP) IsLinkLocalUnicast 

``` go 
func (ip IP) IsLinkLocalUnicast() bool
```

IsLinkLocalUnicast reports whether ip is a link-local unicast address.

​	`IsLinkLocalUnicast`方法报告IP地址ip是否为链路本地单播地址。

##### IsLinkLocalUnicast Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6LinkLocalUni := net.ParseIP("fe80::")
	ipv6Global := net.ParseIP("2000::")
	ipv4LinkLocalUni := net.ParseIP("169.254.0.0")
	ipv4LinkLocalMulti := net.ParseIP("224.0.0.0")

	fmt.Println(ipv6LinkLocalUni.IsLinkLocalUnicast())
	fmt.Println(ipv6Global.IsLinkLocalUnicast())
	fmt.Println(ipv4LinkLocalUni.IsLinkLocalUnicast())
	fmt.Println(ipv4LinkLocalMulti.IsLinkLocalUnicast())

}
Output:

true
false
true
false
```

#### (IP) IsLoopback 

``` go 
func (ip IP) IsLoopback() bool
```

IsLoopback reports whether ip is a loopback address.

​	`IsLoopback`方法报告IP地址ip是否为回环地址。

##### IsLoopback Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6Lo := net.ParseIP("::1")
	ipv6 := net.ParseIP("ff02::1")
	ipv4Lo := net.ParseIP("127.0.0.0")
	ipv4 := net.ParseIP("128.0.0.0")

	fmt.Println(ipv6Lo.IsLoopback())
	fmt.Println(ipv6.IsLoopback())
	fmt.Println(ipv4Lo.IsLoopback())
	fmt.Println(ipv4.IsLoopback())

}
Output:

true
false
true
false
```

#### (IP) IsMulticast 

``` go 
func (ip IP) IsMulticast() bool
```

IsMulticast reports whether ip is a multicast address.

​	`IsMulticast`方法报告IP地址ip是否为多播地址。

##### IsMulticast Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6Multi := net.ParseIP("FF00::")
	ipv6LinkLocalMulti := net.ParseIP("ff02::1")
	ipv6Lo := net.ParseIP("::1")
	ipv4Multi := net.ParseIP("239.0.0.0")
	ipv4LinkLocalMulti := net.ParseIP("224.0.0.0")
	ipv4Lo := net.ParseIP("127.0.0.0")

	fmt.Println(ipv6Multi.IsMulticast())
	fmt.Println(ipv6LinkLocalMulti.IsMulticast())
	fmt.Println(ipv6Lo.IsMulticast())
	fmt.Println(ipv4Multi.IsMulticast())
	fmt.Println(ipv4LinkLocalMulti.IsMulticast())
	fmt.Println(ipv4Lo.IsMulticast())

}
Output:

true
true
false
true
true
false
```

#### (IP) IsPrivate  <- go1.17

``` go 
func (ip IP) IsPrivate() bool
```

IsPrivate reports whether ip is a private address, according to [RFC 1918](https://rfc-editor.org/rfc/rfc1918.html) (IPv4 addresses) and [RFC 4193](https://rfc-editor.org/rfc/rfc4193.html) (IPv6 addresses).

​	`IsPrivate`方法报告IP地址ip是否为私有地址，根据[RFC 1918](https://rfc-editor.org/rfc/rfc1918.html)(IPv4地址)和[RFC 4193](https://rfc-editor.org/rfc/rfc4193.html)(IPv6地址)。

##### IsPrivate Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6Private := net.ParseIP("fc00::")
	ipv6Public := net.ParseIP("fe00::")
	ipv4Private := net.ParseIP("10.255.0.0")
	ipv4Public := net.ParseIP("11.0.0.0")

	fmt.Println(ipv6Private.IsPrivate())
	fmt.Println(ipv6Public.IsPrivate())
	fmt.Println(ipv4Private.IsPrivate())
	fmt.Println(ipv4Public.IsPrivate())

}
Output:

true
false
true
false
```

#### (IP) IsUnspecified 

``` go 
func (ip IP) IsUnspecified() bool
```

IsUnspecified reports whether ip is an unspecified address, either the IPv4 address "0.0.0.0" or the IPv6 address "::".

​	`IsUnspecified`方法报告IP地址ip是否为未指定地址，即IPv4地址"0.0.0.0"或IPv6地址"::"。

##### IsUnspecified Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6Unspecified := net.ParseIP("::")
	ipv6Specified := net.ParseIP("fe00::")
	ipv4Unspecified := net.ParseIP("0.0.0.0")
	ipv4Specified := net.ParseIP("8.8.8.8")

	fmt.Println(ipv6Unspecified.IsUnspecified())
	fmt.Println(ipv6Specified.IsUnspecified())
	fmt.Println(ipv4Unspecified.IsUnspecified())
	fmt.Println(ipv4Specified.IsUnspecified())

}
Output:

true
false
true
false
```

#### (IP) MarshalText  <- go1.2

``` go 
func (ip IP) MarshalText() ([]byte, error)
```

MarshalText implements the encoding.TextMarshaler interface. The encoding is the same as returned by String, with one exception: When len(ip) is zero, it returns an empty slice.

​	`MarshalText`方法实现encoding.TextMarshaler接口。编码与String返回的编码相同，唯一的区别是当len(ip)为零时，它会返回一个空切片。

#### (IP) Mask 

``` go 
func (ip IP) Mask(mask IPMask) IP
```

Mask returns the result of masking the IP address ip with mask.

​	`Mask`方法返回使用掩码mask对IP地址ip进行掩码后的结果。

##### Mask Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv4Addr := net.ParseIP("192.0.2.1")
	// This mask corresponds to a /24 subnet for IPv4.
	ipv4Mask := net.CIDRMask(24, 32)
	fmt.Println(ipv4Addr.Mask(ipv4Mask))

	ipv6Addr := net.ParseIP("2001:db8:a0b:12f0::1")
	// This mask corresponds to a /32 subnet for IPv6.
	ipv6Mask := net.CIDRMask(32, 128)
	fmt.Println(ipv6Addr.Mask(ipv6Mask))

}
Output:

192.0.2.0
2001:db8::
```

#### (IP) String 

``` go 
func (ip IP) String() string
```

String returns the string form of the IP address ip. It returns one of 4 forms:

​	`String`方法返回IP地址ip的字符串形式。它返回以下4种形式之一：

- "<nil>", if ip has length 0
- "`<nil>`"，如果ip长度为0
- dotted decimal ("192.0.2.1"), if ip is an IPv4 or IP4-mapped IPv6 address
- 点分十进制("192.0.2.1")，如果ip是IPv4或IP4映射IPv6地址
- IPv6 conforming to [RFC 5952](https://rfc-editor.org/rfc/rfc5952.html) ("2001:db8::1"), if ip is a valid IPv6 address
- 符合[RFC 5952](https://rfc-editor.org/rfc/rfc5952.html)的IPv6地址("2001:db8::1")，如果ip是有效的IPv6地址
- the hexadecimal form of ip, without punctuation, if no other cases apply
- 如果没有其他情况，则是ip的十六进制形式，不包含标点符号



##### String Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6 := net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ipv4 := net.IPv4(10, 255, 0, 0)

	fmt.Println(ipv6.String())
	fmt.Println(ipv4.String())

}
Output:

fc00::
10.255.0.0
```

#### (IP) To16 

``` go 
func (ip IP) To16() IP
```

To16 converts the IP address ip to a 16-byte representation. If ip is not an IP address (it is the wrong length), To16 returns nil.

​	To16 方法将IP地址ip转换为16字节表示。如果ip不是IP地址(长度不正确)，To16 返回nil。

##### To16 Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6 := net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ipv4 := net.IPv4(10, 255, 0, 0)

	fmt.Println(ipv6.To16())
	fmt.Println(ipv4.To16())

}
Output:

fc00::
10.255.0.0
```

#### (IP) To4 

``` go 
func (ip IP) To4() IP
```

To4 converts the IPv4 address ip to a 4-byte representation. If ip is not an IPv4 address, To4 returns nil.

​	To4方法将IPv4地址ip转换为4字节表示。如果ip不是IPv4地址，则To4返回nil。

#### (*IP) UnmarshalText  <- go1.2

``` go 
func (ip *IP) UnmarshalText(text []byte) error
```

UnmarshalText implements the encoding.TextUnmarshaler interface. The IP address is expected in a form accepted by ParseIP.

​	UnmarshalText方法实现encoding.TextUnmarshaler接口。IP地址应该以ParseIP接受的形式出现。

### type IPAddr 

``` go 
type IPAddr struct {
	IP   IP
	Zone string // IPv6有作用域的寻址区域
}
```

IPAddr represents the address of an IP end point.

​	IPAddr结构体表示IP端点的地址。

#### func ResolveIPAddr 

``` go 
func ResolveIPAddr(network, address string) (*IPAddr, error)
```

ResolveIPAddr returns an address of IP end point.

​	ResolveIPAddr函数返回IP端点的地址。

The network must be an IP network name.

​	网络必须是一个IP网络名称。

If the host in the address parameter is not a literal IP address, ResolveIPAddr resolves the address to an address of IP end point. Otherwise, it parses the address as a literal IP address. The address parameter can use a host name, but this is not recommended, because it will return at most one of the host name's IP addresses.

​	如果address参数中的主机不是字面上的IP地址，则ResolveIPAddr将地址解析为IP端点地址。否则，它将解析地址为一个字面上的IP地址。地址参数可以使用主机名，但不建议这样做，因为它最多只返回主机名的一个IP地址。

See func Dial for a description of the network and address parameters.

​	有关网络和地址参数的描述，请参见func Dial。

#### (*IPAddr) Network 

``` go 
func (a *IPAddr) Network() string
```

Network returns the address's network name, "ip".

​	Network方法返回地址的网络名称，即"ip"。

#### (*IPAddr) String 

``` go 
func (a *IPAddr) String() string
```

### type IPConn 

``` go 
type IPConn struct {
	// 包含过滤或未导出字段
}
```

IPConn is the implementation of the Conn and PacketConn interfaces for IP network connections.

​	IPConn结构体是IP网络连接的Conn和PacketConn接口的实现。

#### func DialIP 

``` go 
func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error)
```

DialIP acts like Dial for IP networks.

​	DialIP函数类似于IP网络的Dial函数。

The network must be an IP network name; see func Dial for details.

​	网络必须是一个IP网络名称；有关详细信息，请参见func Dial。

If laddr is nil, a local address is automatically chosen. If the IP field of raddr is nil or an unspecified IP address, the local system is assumed.

​	如果laddr为nil，则会自动选择本地地址。如果raddr的IP字段为nil或未指定的IP地址，则假定为本地系统。

#### func ListenIP 

``` go 
func ListenIP(network string, laddr *IPAddr) (*IPConn, error)
```

ListenIP acts like ListenPacket for IP networks.

​	`ListenIP`函数类似于IP网络的ListenPacket函数。

The network must be an IP network name; see func Dial for details.

​	网络必须是一个IP网络名称；有关详细信息，请参见func Dial。

If the IP field of laddr is nil or an unspecified IP address, ListenIP listens on all available IP addresses of the local system except multicast IP addresses.

​	如果`laddr`的IP字段为nil或未指定的IP地址，则ListenIP会侦听本地系统的所有可用IP地址，但不包括多播IP地址。

#### (*IPConn) Close 

``` go 
func (c *IPConn) Close() error
```

Close closes the connection.

​	`Close`方法关闭连接。

#### (*IPConn) File 

``` go 
func (c *IPConn) File() (f *os.File, err error)
```

File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished. Closing c does not affect f, and closing f does not affect c.

​	`File`方法返回基础`os.File`的副本。调用者有责任在完成后关闭`f`。关闭`c`不会影响`f`，关闭`f`也不会影响`c`。

The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.

​	返回的`os.File`的文件描述符与连接的不同。尝试使用此副本更改原始文件的属性可能会或可能不会产生预期的效果。

#### (*IPConn) LocalAddr 

``` go 
func (c *IPConn) LocalAddr() Addr
```

LocalAddr returns the local network address. The Addr returned is shared by all invocations of LocalAddr, so do not modify it.

​	`LocalAddr`方法返回本地网络地址。返回的`Addr`由`LocalAddr`的所有调用共享，因此不要修改它。

#### (*IPConn) Read 

``` go 
func (c *IPConn) Read(b []byte) (int, error)
```

Read implements the Conn Read method.

​	`Read`方法实现`Conn Read`方法。

#### (*IPConn) ReadFrom 

``` go 
func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)
```

ReadFrom implements the PacketConn ReadFrom method.

​	`ReadFrom`方法实现`PacketConn ReadFrom` 方法。

#### (*IPConn) ReadFromIP 

``` go 
func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)
```

ReadFromIP acts like ReadFrom but returns an IPAddr.

​	`ReadFromIP`方法类似于`ReadFrom`但返回`IPAddr`。

#### (*IPConn) ReadMsgIP  <- go1.1

``` go 
func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)
```

ReadMsgIP reads a message from c, copying the payload into b and the associated out-of-band data into oob. It returns the number of bytes copied into b, the number of bytes copied into oob, the flags that were set on the message and the source address of the message.

​	ReadMsgIP方法从c读取消息，将有效负载复制到b中，将相关的带外数据复制到oob中。它返回复制到b中的字节数，复制到oob中的字节数，消息设置的标志以及消息的源地址。

The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be used to manipulate IP-level socket options in oob.

​	golang.org/x/net/ipv4 和 golang.org/x/net/ipv6 包可用于操作oob中的IP级套接字选项。

#### (*IPConn) RemoteAddr 

``` go 
func (c *IPConn) RemoteAddr() Addr
```

RemoteAddr returns the remote network address. The Addr returned is shared by all invocations of RemoteAddr, so do not modify it.

​	`RemoteAddr`方法返回远程网络地址。返回的`Addr`由`RemoteAddr`的所有调用共享，因此不要修改它。

#### (*IPConn) SetDeadline 

``` go 
func (c *IPConn) SetDeadline(t time.Time) error
```

SetDeadline implements the Conn SetDeadline method.

​	`SetDeadline`方法实现`Conn SetDeadline`方法。

#### (*IPConn) SetReadBuffer 

``` go 
func (c *IPConn) SetReadBuffer(bytes int) error
```

SetReadBuffer sets the size of the operating system's receive buffer associated with the connection.

​	`SetReadBuffer`方法设置与连接相关联的操作系统接收缓冲区的大小。

#### (*IPConn) SetReadDeadline 

``` go 
func (c *IPConn) SetReadDeadline(t time.Time) error
```

SetReadDeadline implements the Conn SetReadDeadline method.

​	`SetReadDeadline`方法实现Conn SetReadDeadline方法。

#### (*IPConn) SetWriteBuffer 

``` go 
func (c *IPConn) SetWriteBuffer(bytes int) error
```

SetWriteBuffer sets the size of the operating system's transmit buffer associated with the connection.

​	`SetWriteBuffer`方法设置与连接相关联的操作系统传输缓冲区的大小。

#### (*IPConn) SetWriteDeadline 

``` go 
func (c *IPConn) SetWriteDeadline(t time.Time) error
```

SetWriteDeadline implements the Conn SetWriteDeadline method.

​	`SetWriteDeadline`方法实现Conn SetWriteDeadline方法。

#### (*IPConn) SyscallConn  <- go1.9

``` go 
func (c *IPConn) SyscallConn() (syscall.RawConn, error)
```

SyscallConn returns a raw network connection. This implements the syscall.Conn interface.

​	`SyscallConn`方法返回原始网络连接。这实现了`syscall.Conn`接口。

#### (*IPConn) Write 

``` go 
func (c *IPConn) Write(b []byte) (int, error)
```

Write implements the Conn Write method.

​	`Write`方法实现Conn Write方法。

#### (*IPConn) WriteMsgIP  <- go1.1

``` go 
func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)
```

WriteMsgIP writes a message to addr via c, copying the payload from b and the associated out-of-band data from oob. It returns the number of payload and out-of-band bytes written.

​	`WriteMsgIP`方法通过c向addr写入消息，将有效负载从b复制，将相关的带外数据从oob复制。它返回写入的有效负载和带外字节数。

The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be used to manipulate IP-level socket options in oob.

​	golang.org/x/net/ipv4 和 golang.org/x/net/ipv6 包可用于操作oob中的IP级套接字选项。

#### (*IPConn) WriteTo 

``` go 
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)
```

WriteTo implements the PacketConn WriteTo method.

​	`WriteTo`方法实现PacketConn WriteTo方法。

#### (*IPConn) WriteToIP 

``` go 
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)
```

WriteToIP acts like WriteTo but takes an IPAddr.

​	`WriteToIP`方法类似于`WriteTo`但接受`IPAddr`。

### type IPMask 

``` go 
type IPMask []byte
```

An IPMask is a bitmask that can be used to manipulate IP addresses for IP addressing and routing.

​	`IPMask`是一个比特掩码，可用于操作IP地址，用于IP寻址和路由。

See type IPNet and func ParseCIDR for details.

​	有关详细信息，请参见类型`IPNet`和函数`ParseCIDR`。

#### func CIDRMask 

``` go 
func CIDRMask(ones, bits int) IPMask
```

CIDRMask returns an IPMask consisting of 'ones' 1 bits followed by 0s up to a total length of 'bits' bits. For a mask of this form, CIDRMask is the inverse of IPMask.Size.

​	`CIDRMask`函数返回一个`IPMask`，它由"ones"个1位组成，后跟0位，总长度为"bits"位。对于这种形式的掩码，`CIDRMask`是`IPMask.Size`的反函数。

##### CIDRMask Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	// This mask corresponds to a /31 subnet for IPv4.
	fmt.Println(net.CIDRMask(31, 32))

	// This mask corresponds to a /64 subnet for IPv6.
	fmt.Println(net.CIDRMask(64, 128))

}
Output:

fffffffe
ffffffffffffffff0000000000000000
```

#### func IPv4Mask 

``` go 
func IPv4Mask(a, b, c, d byte) IPMask
```

IPv4Mask returns the IP mask (in 4-byte form) of the IPv4 mask a.b.c.d.

​	`IPv4Mask`函数返回IPv4掩码a.b.c.d的IP掩码(以4字节形式)。

##### IPv4Mask Example
``` go 
package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.IPv4Mask(255, 255, 255, 0))

}
Output:

ffffff00
```

#### (IPMask) Size 

``` go 
func (m IPMask) Size() (ones, bits int)
```

Size returns the number of leading ones and total bits in the mask. If the mask is not in the canonical form--ones followed by zeros--then Size returns 0, 0.

​	`Size`方法返回掩码中前导1位和总位数。如果掩码不在规范形式——1位后面是0位——那么Size将返回0, 0。

#### (IPMask) String 

``` go 
func (m IPMask) String() string
```

String returns the hexadecimal form of m, with no punctuation.

​	`String`方法以无标点符号的十六进制形式返回`m`。

### type IPNet 

``` go 
type IPNet struct {
	IP   IP     // 网络号
	Mask IPMask // 网络掩码
}
```

An IPNet represents an IP network.

​	`IPNet`表示IP网络。

#### (*IPNet) Contains 

``` go 
func (n *IPNet) Contains(ip IP) bool
```

Contains reports whether the network includes ip.

​	`Contains`方法报告网络是否包含`ip`。

#### (*IPNet) Network 

``` go 
func (n *IPNet) Network() string
```

Network returns the address's network name, "ip+net".

​	`Network`方法返回地址的网络名称，"ip+net"。

#### (*IPNet) String 

``` go 
func (n *IPNet) String() string
```

String returns the CIDR notation of n like "192.0.2.0/24" or "2001:db8::/48" as defined in [RFC 4632](https://rfc-editor.org/rfc/rfc4632.html) and [RFC 4291](https://rfc-editor.org/rfc/rfc4291.html). If the mask is not in the canonical form, it returns the string which consists of an IP address, followed by a slash character and a mask expressed as hexadecimal form with no punctuation like "198.51.100.0/c000ff00".

​	`String`方法返回`n`的`CIDR`表示形式，例如"192.0.2.0/24"或"2001:db8::/48"，如[RFC 4632](https://rfc-editor.org/rfc/rfc4632.html)和[RFC 4291](https://rfc-editor.org/rfc/rfc4291.html)所定义。如果掩码不在规范形式中，则返回的字符串由IP地址和斜杠字符和表示为十六进制形式且无标点符号的掩码组成，例如"198.51.100.0/c000ff00"。

### type Interface 

``` go 
type Interface struct {
	Index        int          // 正整数，从1开始，0永远不使用 positive integer that starts at one, zero is never used
	MTU          int          // 最大传输单元 maximum transmission unit
	Name         string       // 例如"en0"，"lo0"，"eth0.100" e.g., "en0", "lo0", "eth0.100"
	HardwareAddr HardwareAddr // IEEE MAC-48、EUI-48和EUI-64形式 IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        Flags        // 例如FlagUp、FlagLoopback、FlagMulticast e.g., FlagUp, FlagLoopback, FlagMulticast
}
```

Interface represents a mapping between network interface name and index. It also represents network interface facility information.

​	`Interface`结构体表示网络接口名称和索引之间的映射。它还表示网络接口设施信息。

#### func InterfaceByIndex 

``` go 
func InterfaceByIndex(index int) (*Interface, error)
```

InterfaceByIndex returns the interface specified by index.

​	`InterfaceByIndex`函数返回指定索引的接口。

On Solaris, it returns one of the logical network interfaces sharing the logical data link; for more precision use InterfaceByName.

​	在Solaris上，它返回共享逻辑数据链路的逻辑网络接口之一；要获得更精确的结果，请使用`InterfaceByName`。

#### func InterfaceByName 

``` go 
func InterfaceByName(name string) (*Interface, error)
```

InterfaceByName returns the interface specified by name.

​	`InterfaceByName`函数返回指定名称的接口。

#### func Interfaces 

``` go 
func Interfaces() ([]Interface, error)
```

Interfaces returns a list of the system's network interfaces.

​	`Interfaces`函数返回系统的网络接口列表。

#### (*Interface) Addrs 

``` go 
func (ifi *Interface) Addrs() ([]Addr, error)
```

Addrs returns a list of unicast interface addresses for a specific interface.

​	`Addrs`方法返回特定接口的单播接口地址列表。

#### (*Interface) MulticastAddrs 

``` go 
func (ifi *Interface) MulticastAddrs() ([]Addr, error)
```

MulticastAddrs returns a list of multicast, joined group addresses for a specific interface.

​	`MulticastAddrs`方法返回特定接口的多播、组播地址列表。

### type InvalidAddrError 

``` go 
type InvalidAddrError string
```

#### (InvalidAddrError) Error 

``` go 
func (e InvalidAddrError) Error() string
```

#### (InvalidAddrError) Temporary 

``` go 
func (e InvalidAddrError) Temporary() bool
```

#### (InvalidAddrError) Timeout 

``` go 
func (e InvalidAddrError) Timeout() bool
```

### type ListenConfig  <- go1.11

``` go 
type ListenConfig struct {
    // If Control is not nil, it is called after creating the network
	// connection but before binding it to the operating system.
	// 如果 Control 不是 nil，
    // 则在创建网络连接但在将其绑定到操作系统之前调用它。
	//
    // Network and address parameters passed to Control method are not
	// necessarily the ones passed to Listen. For example, passing "tcp" to
	// Listen will cause the Control function to be called with "tcp4" or "tcp6".
	// 传递给 Control 方法的网络和地址参数未必是传递给 Listen 的参数。
	// 例如，将 "tcp" 传递给 Listen 
    // 将导致 Control 函数使用 "tcp4" 或 "tcp6" 调用。
	Control func(network, address string, c syscall.RawConn) error

    // KeepAlive specifies the keep-alive period for network
	// connections accepted by this listener.
	// If zero, keep-alives are enabled if supported by the protocol
	// and operating system. Network protocols or operating systems
	// that do not support keep-alives ignore this field.
	// If negative, keep-alives are disabled.
	// KeepAlive 指定此侦听器接受的网络连接的保活期。
	// 如果为零，则如果协议和操作系统支持保活，则启用保活。
    // 不支持保活的网络协议或操作系统将忽略此字段。
	// 如果为负，则禁用保活。
	KeepAlive time.Duration
}
```

ListenConfig contains options for listening to an address.

​	`ListenConfig`结构体包含用于监听地址的选项。

#### (*ListenConfig) Listen  <- go1.11

``` go 
func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)
```

Listen announces on the local network address.

​	`Listen`方法在本地网络地址上进行侦听。

See func Listen for a description of the network and address parameters.

​	有关网络和地址参数的描述，请参见 func Listen。

#### (*ListenConfig) ListenPacket  <- go1.11

``` go 
func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error)
```

ListenPacket announces on the local network address.

​	`ListenPacket`方法在本地网络地址上进行侦听。

See func ListenPacket for a description of the network and address parameters.

​	有关网络和地址参数的描述，请参见 func ListenPacket。

#### (*ListenConfig) MultipathTCP <- go1.21.0

``` go
func (lc *ListenConfig) MultipathTCP() bool
```

MultipathTCP reports whether MPTCP will be used.

This method doesn't check if MPTCP is supported by the operating system or not.

####  (*ListenConfig) SetMultipathTCP <-go1.21.0

``` go
func (lc *ListenConfig) SetMultipathTCP(use bool)
```

SetMultipathTCP directs the Listen method to use, or not use, MPTCP, if supported by the operating system. This method overrides the system default and the GODEBUG=multipathtcp=... setting if any.

If MPTCP is not available on the host or not supported by the client, the Listen method will fall back to TCP.

### type Listener 

``` go 
type Listener interface {
    // Accept waits for and returns the next connection to the listener.
	// Accept 等待并返回到监听器的下一个连接。
	Accept() (Conn, error)

    // Close closes the listener.
	// Any blocked Accept operations will be unblocked and return errors.
	// Close 关闭监听器。
	// 任何已阻止的 Accept 操作将解除阻塞并返回错误。
	Close() error

    // Addr returns the listener's network address.
	// Addr 返回监听器的网络地址。
	Addr() Addr
}
```

A Listener is a generic network listener for stream-oriented protocols.

​	`Listener`接口是面向流协议的通用网络监听器。

Multiple goroutines may invoke methods on a Listener simultaneously.

​	多个 goroutine 可同时调用 `Listener` 上的方法。

##### Listener Example
``` go 
package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

```

#### func FileListener 

``` go 
func FileListener(f *os.File) (ln Listener, err error)
```

FileListener returns a copy of the network listener corresponding to the open file f. It is the caller's responsibility to close ln when finished. Closing ln does not affect f, and closing f does not affect ln.

​	`FileListener`函数返回与已打开文件 f 相应的网络监听器的副本。调用者有责任在完成时关闭 ln。关闭 ln 不会影响 f，关闭 f 不会影响 ln。

#### func Listen 

``` go 
func Listen(network, address string) (Listener, error)
```

Listen announces on the local network address.

​	`Listen`函数在本地网络地址上进行监听。

The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".

​	网络必须是 "tcp"、"tcp4"、"tcp6"、"unix" 或 "unixpacket"。

For TCP networks, if the host in the address parameter is empty or a literal unspecified IP address, Listen listens on all available unicast and anycast IP addresses of the local system. To only use IPv4, use network "tcp4". The address can use a host name, but this is not recommended, because it will create a listener for at most one of the host's IP addresses. If the port in the address parameter is empty or "0", as in "127.0.0.1:" or "[::1]:0", a port number is automatically chosen. The Addr method of Listener can be used to discover the chosen port.

​	对于 TCP 网络，如果地址参数中的主机为空或为字面未指定 IP 地址，则 Listen 会侦听本地系统的所有可用单播和任播 IP 地址。要仅使用 IPv4，请使用网络 "tcp4"。地址可以使用主机名，但不建议这样做，因为它将仅为主机的一个 IP 地址创建侦听器。如果地址参数中的端口为空或为 "0"，例如 "127.0.0.1:" 或 "[::1]:0"，则会自动选择一个端口号。可以使用 Listener 的 Addr 方法来查找所选端口。

See func Dial for a description of the network and address parameters.

​	有关网络和地址参数的描述，请参见 func Dial。

Listen uses context.Background internally; to specify the context, use ListenConfig.Listen.

​	`Listen` 在内部使用 context.Background；要指定上下文，请使用 ListenConfig.Listen。

### type MX 

``` go 
type MX struct {
	Host string
	Pref uint16
}
```

An MX represents a single DNS MX record.

​	`MX`结构体表示单个 DNS MX 记录。

#### func LookupMX 

``` go 
func LookupMX(name string) ([]*MX, error)
```

LookupMX returns the DNS MX records for the given domain name sorted by preference.

​	`LookupMX`函数按优先级排序返回给定域名的 DNS MX 记录。

The returned mail server names are validated to be properly formatted presentation-format domain names. If the response contains invalid names, those records are filtered out and an error will be returned alongside the remaining results, if any.

​	返回的邮件服务器名称已验证为格式正确的演示格式域名。如果响应包含无效名称，则这些记录将被过滤掉，并将在剩余结果(如果有)旁边返回错误。

LookupMX uses context.Background internally; to specify the context, use Resolver.LookupMX.

​	`LookupMX` 在内部使用 `context.Background`；要指定上下文，请使用 `Resolver.LookupMX`。

### type NS  <- go1.1

``` go 
type NS struct {
	Host string
}
```

An NS represents a single DNS NS record.

​	`NS`结构体表示单个 DNS NS 记录。

#### func LookupNS  <- go1.1

``` go 
func LookupNS(name string) ([]*NS, error)
```

LookupNS returns the DNS NS records for the given domain name.

​	`LookupNS` 函数返回给定域名的 DNS NS 记录。

The returned name server names are validated to be properly formatted presentation-format domain names. If the response contains invalid names, those records are filtered out and an error will be returned alongside the remaining results, if any.

​	返回的名称服务器名称已验证为格式正确的演示格式域名。如果响应包含无效名称，则这些记录将被过滤掉，并将在剩余结果(如果有)旁边返回错误。

LookupNS uses context.Background internally; to specify the context, use Resolver.LookupNS.

​	`LookupNS` 在内部使用 `context.Background`；要指定上下文，请使用 `Resolver.LookupNS`。

### type OpError 

``` go 
type OpError struct {
    // Op is the operation which caused the error, such as
	// "read" or "write".
	// Op(操作)是引起错误的操作，例如"read"或"write"。
	Op string

    // Net is the network type on which this error occurred,
	// such as "tcp" or "udp6".
	// Net(网络)是发生此错误的网络类型，例如"tcp"或"udp6"。
	Net string

    // For operations involving a remote network connection, like
	// Dial, Read, or Write, Source is the corresponding local
	// network address.
	// 对于涉及远程网络连接的操作，
    // 如Dial、Read或Write，Source是相应的本地网络地址。
	Source Addr

    // Addr is the network address for which this error occurred.
	// For local operations, like Listen or SetDeadline, Addr is
	// the address of the local endpoint being manipulated.
	// For operations involving a remote network connection, like
	// Dial, Read, or Write, Addr is the remote address of that
	// connection.
	// Addr(地址)是发生此错误的网络地址。
	// 对于本地操作，如Listen或SetDeadline，
    // Addr是正在操作的本地端点的地址。
	// 对于涉及远程网络连接的操作，
    // 如Dial、Read或Write，Addr是该连接的远程地址。
	Addr Addr

    // Err is the error that occurred during the operation.
	// The Error method panics if the error is nil.
	// Err(错误)是操作期间发生的错误。
	// 如果错误为nil，则Error方法会引发panic。
	Err error
}
```

OpError is the error type usually returned by functions in the net package. It describes the operation, network type, and address of an error.

​	`OpError`结构体是通常由net包中的函数返回的错误类型。它描述了错误的操作、网络类型和地址。

#### (*OpError) Error 

``` go 
func (e *OpError) Error() string
```

#### (*OpError) Temporary 

``` go 
func (e *OpError) Temporary() bool
```

#### (*OpError) Timeout 

``` go 
func (e *OpError) Timeout() bool
```

#### (*OpError) Unwrap  <- go1.13

``` go 
func (e *OpError) Unwrap() error
```

### type PacketConn 

``` go 
type PacketConn interface {
    // ReadFrom reads a packet from the connection,
	// copying the payload into p. It returns the number of
	// bytes copied into p and the return address that
	// was on the packet.
	// It returns the number of bytes read (0 <= n <= len(p))
	// and any error encountered. Callers should always process
	// the n > 0 bytes returned before considering the error err.
	// ReadFrom can be made to time out and return an error after a
	// fixed time limit; see SetDeadline and SetReadDeadline.
	// ReadFrom 从连接中读取一个数据包，将有效载荷复制到 p 中。
	// 它返回复制到 p 中的字节数以及在数据包上的返回地址。
	// 它返回读取的字节数(0 <= n <= len(p))和任何遇到的错误。
	// 调用者应该始终在考虑错误 err 之前处理返回的 n > 0 字节。
	// ReadFrom 可以在固定的时间限制后超时并返回一个错误；
    // 参见 SetDeadline 和 SetReadDeadline。
	ReadFrom(p []byte) (n int, addr Addr, err error)

    // WriteTo writes a packet with payload p to addr.
	// WriteTo can be made to time out and return an Error after a
	// fixed time limit; see SetDeadline and SetWriteDeadline.
	// On packet-oriented connections, write timeouts are rare.
	// WriteTo 将有效载荷为 p 的数据包写入 addr。
	// WriteTo 可以在固定的时间限制后超时并返回一个错误；
    // 参见 SetDeadline 和 SetWriteDeadline。
	// 在面向数据包的连接中，写入超时很少发生。
	WriteTo(p []byte, addr Addr) (n int, err error)

    // Close closes the connection.
	// Any blocked ReadFrom or WriteTo operations will be unblocked and return errors.
	// Close 关闭连接。
	// 任何被阻塞的 ReadFrom 或 WriteTo 操作都将被取消阻塞并返回错误。
	Close() error

    // LocalAddr returns the local network address, if known.
	// LocalAddr 返回本地网络地址(如果已知)。
	LocalAddr() Addr

    // SetDeadline sets the read and write deadlines associated
	// with the connection. It is equivalent to calling both
	// SetReadDeadline and SetWriteDeadline.
	// SetDeadline 设置与连接关联的读取和写入截止时间。
	// 它相当于调用 SetReadDeadline 和 SetWriteDeadline。
	//
    // A deadline is an absolute time after which I/O operations
	// fail instead of blocking. The deadline applies to all future
	// and pending I/O, not just the immediately following call to
	// Read or Write. After a deadline has been exceeded, the
	// connection can be refreshed by setting a deadline in the future.
	// 截止时间是绝对时间，超过这个时间，I/O 操作将失败而不是阻塞。
	// 截止时间适用于所有未来和挂起的 I/O，
    // 而不仅仅是即将到来的 Read 或 Write 调用。
	// 在超过截止时间后，可以通过将未来的截止时间设置为刷新连接。
	//
    // If the deadline is exceeded a call to Read or Write or to other
	// I/O methods will return an error that wraps os.ErrDeadlineExceeded.
	// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
	// The error's Timeout method will return true, but note that there
	// are other possible errors for which the Timeout method will
	// return true even if the deadline has not been exceeded.
	// 如果超过截止时间，调用 Read 或 Write 或其他 I/O 方法
    // 将返回一个错误，该错误包装了 os.ErrDeadlineExceeded。
	// 可以使用 errors.Is(err, os.ErrDeadlineExceeded) 进行测试。
	// 错误的 Timeout 方法将返回 true，
    // 但请注意，即使未超过截止时间，也可能有其他可能的错误，
    // Timeout 方法将返回 true。
	//
    // An idle timeout can be implemented by repeatedly extending
	// the deadline after successful ReadFrom or WriteTo calls.
    //
    // A zero value for t means I/O operations will not time out.
	// 空值 t 表示 I/O 操作不会超时。
	SetDeadline(t time.Time) error

    // SetReadDeadline sets the deadline for future ReadFrom calls
	// and any currently-blocked ReadFrom call.
	// A zero value for t means ReadFrom will not time out.
	// SetReadDeadline 设置未来的 ReadFrom 调用的截止时间和
    // 任何当前被阻塞的 ReadFrom 调用。
	// t 的零值表示 ReadFrom 不会超时。
	SetReadDeadline(t time.Time) error

    // SetWriteDeadline sets the deadline for future WriteTo calls
	// and any currently-blocked WriteTo call.
	// Even if write times out, it may return n > 0, indicating that
	// some of the data was successfully written.
	// A zero value for t means WriteTo will not time out.
	// SetWriteDeadline 设置未来的 WriteTo 调用的截止时间和
    // 任何当前被阻塞的 WriteTo 调用。
	// 即使写入超时，它也可能返回 n > 0，表示某些数据已成功写入。
	// t 的零值表示 WriteTo 不会超时。
	SetWriteDeadline(t time.Time) error
}
```

PacketConn is a generic packet-oriented network connection.

​	`PacketConn`接口是一个通用的面向数据包的网络连接。

Multiple goroutines may invoke methods on a PacketConn simultaneously.

​	多个 goroutine 可以同时调用 `PacketConn` 上的方法。

#### func FilePacketConn 

``` go 
func FilePacketConn(f *os.File) (c PacketConn, err error)
```

FilePacketConn returns a copy of the packet network connection corresponding to the open file f. It is the caller's responsibility to close f when finished. Closing c does not affect f, and closing f does not affect c.

​	`FilePacketConn`函数返回与打开的文件 f 对应的数据包网络连接的副本。当使用完毕时，调用方负责关闭 f。关闭 c 不影响 f，关闭 f 也不影响 c。

#### func ListenPacket 

``` go 
func ListenPacket(network, address string) (PacketConn, error)
```

ListenPacket announces on the local network address.

​	`ListenPacket`函数在本地网络地址上进行公告。

The network must be "udp", "udp4", "udp6", "unixgram", or an IP transport. The IP transports are "ip", "ip4", or "ip6" followed by a colon and a literal protocol number or a protocol name, as in "ip:1" or "ip:icmp".

​	网络必须是 "udp"、"udp4"、"udp6"、"unixgram" 或 IP 传输。IP 传输为 "ip"、"ip4" 或 "ip6"，后跟冒号和字面协议号或协议名称，例如 "ip:1" 或 "ip:icmp"。

For UDP and IP networks, if the host in the address parameter is empty or a literal unspecified IP address, ListenPacket listens on all available IP addresses of the local system except multicast IP addresses. To only use IPv4, use network "udp4" or "ip4:proto". The address can use a host name, but this is not recommended, because it will create a listener for at most one of the host's IP addresses. If the port in the address parameter is empty or "0", as in "127.0.0.1:" or "[::1]:0", a port number is automatically chosen. The LocalAddr method of PacketConn can be used to discover the chosen port.

​	对于 UDP 和 IP 网络，如果 address 参数中的主机为空或为字面未指定的 IP 地址，则 ListenPacket 会在本地系统的所有可用 IP 地址上监听，但不包括多播 IP 地址。要仅使用 IPv4，请使用网络 "udp4" 或 "ip4:proto"。地址可以使用主机名，但不建议这样做，因为这将仅为主机的一个 IP 地址创建侦听器。如果地址参数中的端口为空或为 "0"，例如 "127.0.0.1:" 或 "[::1]:0"，则会自动选择一个端口号。可以使用 PacketConn 的 LocalAddr 方法来发现所选端口。

See func Dial for a description of the network and address parameters.

​	有关网络和地址参数的描述，请参见 func Dial。

ListenPacket uses context.Background internally; to specify the context, use ListenConfig.ListenPacket.

​	ListenPacket 在内部使用 context.Background；要指定上下文，请使用 ListenConfig.ListenPacket。

### type ParseError 

``` go 
type ParseError struct {
    // Type is the type of string that was expected, such as
	// "IP address", "CIDR address".
	// Type 是预期的字符串类型，例如
	// "IP 地址"、"CIDR 地址"。
	Type string

    // Text is the malformed text string.
	// Text 是格式不正确的文本字符串。
	Text string
}
```

A ParseError is the error type of literal network address parsers.

​	`ParseError`结构体是文字网络地址解析器的错误类型。

#### (*ParseError) Error 

``` go 
func (e *ParseError) Error() string
```

#### (*ParseError) Temporary  <- go1.17

``` go 
func (e *ParseError) Temporary() bool
```

#### (*ParseError) Timeout  <- go1.17

``` go 
func (e *ParseError) Timeout() bool
```

### type Resolver  <- go1.8

``` go 
type Resolver struct {
    // PreferGo controls whether Go's built-in DNS resolver is preferred
	// on platforms where it's available. It is equivalent to setting
	// GODEBUG=netdns=go, but scoped to just this resolver.
	// PreferGo 控制在可用平台上是否优先使用 Go 的内置 DNS 解析器。
	// 它等效于设置 GODEBUG=netdns=go，但仅限于此解析器。
	PreferGo bool

    // StrictErrors controls the behavior of temporary errors
	// (including timeout, socket errors, and SERVFAIL) when using
	// Go's built-in resolver. For a query composed of multiple
	// sub-queries (such as an A+AAAA address lookup, or walking the
	// DNS search list), this option causes such errors to abort the
	// whole query instead of returning a partial result. This is
	// not enabled by default because it may affect compatibility
	// with resolvers that process AAAA queries incorrectly.
	// StrictErrors 控制使用 Go 的内置解析器时临时错误的行为
	//(包括超时、套接字错误和 SERVFAIL)。
    // 对于由多个子查询(例如 A+AAAA 地址查找或遍历 DNS 搜索列表)
    // 组成的查询，此选项会导致这些错误中止整个查询，
    // 而不是返回部分结果。默认情况下未启用此选项，
    // 因为它可能会影响与处理 AAAA 查询不正确的解析器的兼容性。
	StrictErrors bool

    // Dial optionally specifies an alternate dialer for use by
	// Go's built-in DNS resolver to make TCP and UDP connections
	// to DNS services. The host in the address parameter will
	// always be a literal IP address and not a host name, and the
	// port in the address parameter will be a literal port number
	// and not a service name.
	// If the Conn returned is also a PacketConn, sent and received DNS
	// messages must adhere to RFC 1035 section 4.2.1, "UDP usage".
	// Otherwise, DNS messages transmitted over Conn must adhere
	// to RFC 7766 section 5, "Transport Protocol Selection".
	// If nil, the default dialer is used.
	// Dial 可选地指定供 Go 的内置 DNS 解析器使用的替代拨号器，
    // 以建立到 DNS 服务的 TCP 和 UDP 连接。
    // address 参数中的主机将始终是字面 IP 地址而不是主机名，
    // address 参数中的端口将是字面端口号而不是服务名称。
	// 如果返回的 Conn 也是一个 PacketConn，
    // 则发送和接收的 DNS 消息必须遵守 RFC 1035 
    // 第 4.2.1 节 "UDP usage"。
    // 否则，通过 Conn 传输的 DNS 消息必须遵守 RFC 7766 
    // 第 5 节 "Transport Protocol Selection"。
	// 如果为 nil，则使用默认拨号程序。
	Dial func(ctx context.Context, network, address string) (Conn, error)
	// 包含过滤或未公开的字段
}
```

A Resolver looks up names and numbers.

​	`Resolver` 查找名称和数字。

A nil *Resolver is equivalent to a zero Resolver.

​	nil 的 `*Resolver` 等效于零 `Resolver`。

#### (*Resolver) LookupAddr  <- go1.8

``` go 
func (r *Resolver) LookupAddr(ctx context.Context, addr string) ([]string, error)
```

LookupAddr performs a reverse lookup for the given address, returning a list of names mapping to that address.

​	`LookupAddr`方法执行给定地址的反向查找，返回映射到该地址的名称列表。

The returned names are validated to be properly formatted presentation-format domain names. If the response contains invalid names, those records are filtered out and an error will be returned alongside the remaining results, if any.

​	返回的名称经过验证，以确保格式正确的呈现格式域名。如果响应包含无效名称，则这些记录将被过滤掉，并在剩余结果(如果有)的同时返回错误。

#### (*Resolver) LookupCNAME  <- go1.8

``` go 
func (r *Resolver) LookupCNAME(ctx context.Context, host string) (string, error)
```

LookupCNAME returns the canonical name for the given host. Callers that do not care about the canonical name can call LookupHost or LookupIP directly; both take care of resolving the canonical name as part of the lookup.

​	`LookupCNAME`方法返回给定主机的规范名称。不关心规范名称的调用方可以直接调用 LookupHost 或 LookupIP；它们都会在查找过程中解析规范名称。

A canonical name is the final name after following zero or more CNAME records. LookupCNAME does not return an error if host does not contain DNS "CNAME" records, as long as host resolves to address records.

​	规范名称是在遵循零个或多个 CNAME 记录后的最终名称。如果主机不包含 DNS "CNAME" 记录，LookupCNAME 不会返回错误，只要主机解析为地址记录即可。

The returned canonical name is validated to be a properly formatted presentation-format domain name.

​	返回的规范名称经过验证，以确保格式正确的呈现格式域名。

#### (*Resolver) LookupHost  <- go1.8

``` go 
func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error)
```

LookupHost looks up the given host using the local resolver. It returns a slice of that host's addresses.

​	`LookupHost`方法使用本地解析器查找给定的主机。它返回该主机的地址片段。

#### (*Resolver) LookupIP  <- go1.15

``` go 
func (r *Resolver) LookupIP(ctx context.Context, network, host string) ([]IP, error)
```

LookupIP looks up host for the given network using the local resolver. It returns a slice of that host's IP addresses of the type specified by network. network must be one of "ip", "ip4" or "ip6".

​	`LookupIP`方法使用本地解析器查找给定网络的主机。它返回该主机指定类型的 IP 地址的片段。network 必须是 "ip"、"ip4" 或 "ip6" 之一。

#### (*Resolver) LookupIPAddr  <- go1.8

``` go 
func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error)
```

LookupIPAddr looks up host using the local resolver. It returns a slice of that host's IPv4 and IPv6 addresses.

​	`LookupIPAddr`方法使用本地解析器查找主机。它返回该主机的 IPv4 和 IPv6 地址的片段。

#### (*Resolver) LookupMX  <- go1.8

``` go 
func (r *Resolver) LookupMX(ctx context.Context, name string) ([]*MX, error)
```

LookupMX returns the DNS MX records for the given domain name sorted by preference.

​	`LookupMX`方法返回给定域名的 DNS MX 记录，并按优先级排序。

The returned mail server names are validated to be properly formatted presentation-format domain names. If the response contains invalid names, those records are filtered out and an error will be returned alongside the remaining results, if any.

​	返回的邮件服务器名称经过验证，以确保格式正确的呈现格式域名。如果响应包含无效名称，则这些记录将被过滤掉，并在剩余结果(如果有)的同时返回错误。

#### (*Resolver) LookupNS  <- go1.8

``` go 
func (r *Resolver) LookupNS(ctx context.Context, name string) ([]*NS, error)
```

LookupNS returns the DNS NS records for the given domain name.

​	`LookupNS`方法返回给定域名的 DNS NS 记录。

The returned name server names are validated to be properly formatted presentation-format domain names. If the response contains invalid names, those records are filtered out and an error will be returned alongside the remaining results, if any.

​	返回的名称服务器名称已验证为格式正确的演示格式域名。如果响应包含无效名称，则这些记录将被过滤，并将返回一个错误以及剩余结果(如果有)。

#### (*Resolver) LookupNetIP  <- go1.18

``` go 
func (r *Resolver) LookupNetIP(ctx context.Context, network, host string) ([]netip.Addr, error)
```

LookupNetIP looks up host using the local resolver. It returns a slice of that host's IP addresses of the type specified by network. The network must be one of "ip", "ip4" or "ip6".

​	`LookupNetIP`方法使用本地解析器查找主机。它返回指定网络类型的主机 IP 地址切片。网络类型必须是"ip"、"ip4"或"ip6"。

#### (*Resolver) LookupPort  <- go1.8

``` go 
func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error)
```

LookupPort looks up the port for the given network and service.

​	`LookupPort`方法查找给定网络和服务的端口。

#### (*Resolver) LookupSRV  <- go1.8

``` go 
func (r *Resolver) LookupSRV(ctx context.Context, service, proto, name string) (string, []*SRV, error)
```

LookupSRV tries to resolve an SRV query of the given service, protocol, and domain name. The proto is "tcp" or "udp". The returned records are sorted by priority and randomized by weight within a priority.

​	`LookupSRV`方法尝试解析给定服务、协议和域名的 SRV 查询。`proto` 为"tcp"或"udp"。返回的记录按优先级排序，并在优先级内按权重随机排序。

LookupSRV constructs the DNS name to look up following [RFC 2782](https://rfc-editor.org/rfc/rfc2782.html). That is, it looks up _service._proto.name. To accommodate services publishing SRV records under non-standard names, if both service and proto are empty strings, LookupSRV looks up name directly.

​	`LookupSRV` 根据 [RFC 2782](https://rfc-editor.org/rfc/rfc2782.html)构造要查找的 DNS 名称。也就是说，它查找 _service._proto.name。为了适应在非标准名称下发布 SRV 记录的服务，如果服务和 proto 都为空字符串，则 LookupSRV 直接查找 name。

The returned service names are validated to be properly formatted presentation-format domain names. If the response contains invalid names, those records are filtered out and an error will be returned alongside the remaining results, if any.

​	返回的服务名称已验证为格式正确的演示格式域名。如果响应包含无效名称，则这些记录将被过滤，并将返回一个错误以及剩余结果(如果有)。

#### (*Resolver) LookupTXT  <- go1.8

``` go 
func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error)
```

LookupTXT returns the DNS TXT records for the given domain name.

​	`LookupTXT`方法返回给定域名的 DNS TXT 记录。

### type SRV 

``` go 
type SRV struct {
	Target   string
	Port     uint16
	Priority uint16
	Weight   uint16
}
```

An SRV represents a single DNS SRV record.

​	`SRV` 表示单个 DNS SRV 记录。

#### func LookupSRV 

``` go 
func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)
```

LookupSRV tries to resolve an SRV query of the given service, protocol, and domain name. The proto is "tcp" or "udp". The returned records are sorted by priority and randomized by weight within a priority.

​	`LookupSRV`函数尝试解析给定服务、协议和域名的 SRV 查询。proto 是 "tcp" 或 "udp"。返回的记录按优先级排序，并在优先级内按权重随机分布。

LookupSRV constructs the DNS name to look up following [RFC 2782](https://rfc-editor.org/rfc/rfc2782.html). That is, it looks up _service._proto.name. To accommodate services publishing SRV records under non-standard names, if both service and proto are empty strings, LookupSRV looks up name directly.

​	`LookupSRV` 构造了遵循 [RFC 2782](https://rfc-editor.org/rfc/rfc2782.html)的 DNS 名称进行查找。也就是说，它查找 `_service._proto.name`。为了适应发布非标准名称的 SRV 记录的服务，如果 service 和 proto 都是空字符串，则 LookupSRV 直接查找 name。

The returned service names are validated to be properly formatted presentation-format domain names. If the response contains invalid names, those records are filtered out and an error will be returned alongside the remaining results, if any.

​	返回的服务名已验证为格式正确的表示格式的域名。如果响应包含无效名称，则这些记录将被过滤掉，并且将返回错误以及剩余结果(如果有)。

### type TCPAddr 

``` go 
type TCPAddr struct {
	IP   IP
	Port int
	Zone string // IPv6 范围限定符
}
```

TCPAddr represents the address of a TCP end point.

​	`TCPAddr`结构体表示 TCP 端点的地址。

#### func ResolveTCPAddr 

``` go 
func ResolveTCPAddr(network, address string) (*TCPAddr, error)
```

ResolveTCPAddr returns an address of TCP end point.

​	`ResolveTCPAddr`函数返回 TCP 端点的地址。

The network must be a TCP network name.

​	`network` 必须是 TCP 网络名称。

If the host in the address parameter is not a literal IP address or the port is not a literal port number, ResolveTCPAddr resolves the address to an address of TCP end point. Otherwise, it parses the address as a pair of literal IP address and port number. The address parameter can use a host name, but this is not recommended, because it will return at most one of the host name's IP addresses.

​	如果 address 参数中的主机不是字面 IP 地址，或者端口不是字面端口号，则 ResolveTCPAddr 将地址解析为 TCP 端点的地址。否则，它将地址解析为字面 IP 地址和端口号对。address 参数可以使用主机名，但不建议这样做，因为它最多只会返回主机名的一个 IP 地址。

See func Dial for a description of the network and address parameters.

​	有关 network 和 address 参数的说明，请参见 func Dial。

#### func TCPAddrFromAddrPort  <- go1.18

``` go 
func TCPAddrFromAddrPort(addr netip.AddrPort) *TCPAddr
```

TCPAddrFromAddrPort returns addr as a TCPAddr. If addr.IsValid() is false, then the returned TCPAddr will contain a nil IP field, indicating an address family-agnostic unspecified address.

​	`TCPAddrFromAddrPort`函数将 `addr` 转换为 `TCPAddr`。如果 `addr.IsValid()` 为 `false`，则返回的 `TCPAddr` 将包含一个空的 IP 字段，表示不指定地址族的未指定地址。

#### (*TCPAddr) AddrPort  <- go1.18

``` go 
func (a *TCPAddr) AddrPort() netip.AddrPort
```

AddrPort returns the TCPAddr a as a netip.AddrPort.

​	`AddrPort`方法将 `TCPAddr` `a` 转换为 `netip.AddrPort`。

If a.Port does not fit in a uint16, it's silently truncated.

​	如果 `a.Port` 无法适应 uint16，则会静默截断。

If a is nil, a zero value is returned.

​	如果 `a` 为 nil，则返回零值。

#### (*TCPAddr) Network 

``` go 
func (a *TCPAddr) Network() string
```

Network returns the address's network name, "tcp".

​	Network方法返回地址的网络名称，即"tcp"。

#### (*TCPAddr) String 

``` go 
func (a *TCPAddr) String() string
```

### type TCPConn 

``` go 
type TCPConn struct {
	// 包含过滤或未公开的字段
}
```

TCPConn is an implementation of the Conn interface for TCP network connections.

​	`TCPConn`结构体是TCP网络连接的Conn接口实现。

#### func DialTCP 

``` go 
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
```

DialTCP acts like Dial for TCP networks.

​	`DialTCP`函数用于TCP网络连接。

The network must be a TCP network name; see func Dial for details.

​	network参数必须是TCP网络的名称，有关详细信息，请参见func Dial。

If laddr is nil, a local address is automatically chosen. If the IP field of raddr is nil or an unspecified IP address, the local system is assumed.

​	如果`laddr`是nil，则会自动选择本地地址。如果`raddr`的IP字段为nil或未指定IP地址，则会假定为本地系统。

#### (*TCPConn) Close 

``` go 
func (c *TCPConn) Close() error
```

Close closes the connection.

​	`Close`方法关闭连接。

#### (*TCPConn) CloseRead 

``` go 
func (c *TCPConn) CloseRead() error
```

CloseRead shuts down the reading side of the TCP connection. Most callers should just use Close.

​	`CloseRead`方法关闭TCP连接的读取侧。大多数调用方应该使用`Close`。

#### (*TCPConn) CloseWrite 

``` go 
func (c *TCPConn) CloseWrite() error
```

CloseWrite shuts down the writing side of the TCP connection. Most callers should just use Close.

​	`CloseWrite`方法关闭TCP连接的写入侧。大多数调用方应该使用`Close`。

#### (*TCPConn) File 

``` go 
func (c *TCPConn) File() (f *os.File, err error)
```

File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished. Closing c does not affect f, and closing f does not affect c.

​	`File`方法返回基础`os.File`的副本。调用方有责任在完成后关闭`f`。关闭`c`不会影响`f`，关闭f也不会影响`c`。

The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.

​	返回的`os.File`的文件描述符与连接的不同。尝试使用此副本更改原始文件的属性可能会或可能不会产生预期的效果。

#### (*TCPConn) LocalAddr 

``` go 
func (c *TCPConn) LocalAddr() Addr
```

LocalAddr returns the local network address. The Addr returned is shared by all invocations of LocalAddr, so do not modify it.

​	`LocalAddr`方法返回本地网络地址。返回的Addr被所有调用LocalAddr共享，因此请勿修改它。

#### (*TCPConn) MultipathTCP <-go1.21.0

``` go
func (c *TCPConn) MultipathTCP() (bool, error)
```

MultipathTCP reports whether the ongoing connection is using MPTCP.

If Multipath TCP is not supported by the host, by the other peer or intentionally / accidentally filtered out by a device in between, a fallback to TCP will be done. This method does its best to check if MPTCP is still being used or not.

On Linux, more conditions are verified on kernels >= v5.16, improving the results.

#### (*TCPConn) Read 

``` go 
func (c *TCPConn) Read(b []byte) (int, error)
```

Read implements the Conn Read method.

​	`Read`方法实现了`Conn`的`Read`方法。

#### (*TCPConn) ReadFrom 

``` go 
func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)
```

ReadFrom implements the io.ReaderFrom ReadFrom method.

​	`ReadFrom`方法实现了`io.ReaderFrom`的`ReadFrom`方法。

#### (*TCPConn) RemoteAddr 

``` go 
func (c *TCPConn) RemoteAddr() Addr
```

RemoteAddr returns the remote network address. The Addr returned is shared by all invocations of RemoteAddr, so do not modify it.

​	`RemoteAddr`方法返回远程网络地址。返回的`Addr`由所有`RemoteAddr`调用共享，因此请不要修改它。

#### (*TCPConn) SetDeadline 

``` go 
func (c *TCPConn) SetDeadline(t time.Time) error
```

SetDeadline implements the Conn SetDeadline method.

​	`SetDeadline`方法实现了`Conn`的`SetDeadline`方法。

#### (*TCPConn) SetKeepAlive 

``` go 
func (c *TCPConn) SetKeepAlive(keepalive bool) error
```

SetKeepAlive sets whether the operating system should send keep-alive messages on the connection.

​	`SetKeepAlive`方法设置操作系统是否应在连接上发送保持活动消息。

#### (*TCPConn) SetKeepAlivePeriod  <- go1.2

``` go 
func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error
```

SetKeepAlivePeriod sets period between keep-alives.

​	`SetKeepAlivePeriod`方法设置保持活动消息之间的时间间隔。

#### (*TCPConn) SetLinger 

``` go 
func (c *TCPConn) SetLinger(sec int) error
```

SetLinger sets the behavior of Close on a connection which still has data waiting to be sent or to be acknowledged.

​	`SetLinger`方法设置在仍有等待发送或确认的数据的连接上执行`Close`时的行为。

If sec < 0 (the default), the operating system finishes sending the data in the background.

​	如果sec < 0(默认值)，则操作系统在后台完成发送数据。

If sec == 0, the operating system discards any unsent or unacknowledged data.

​	如果sec == 0，则操作系统丢弃任何未发送或未确认的数据。

If sec > 0, the data is sent in the background as with sec < 0. On some operating systems including Linux, this may cause Close to block until all data has been sent or discarded. On some operating systems after sec seconds have elapsed any remaining unsent data may be discarded.

​	如果sec > 0，则数据像sec < 0一样在后台发送。在某些操作系统上，经过sec秒后，任何剩余的未发送数据可能会被丢弃。

#### (*TCPConn) SetNoDelay 

``` go 
func (c *TCPConn) SetNoDelay(noDelay bool) error
```

SetNoDelay controls whether the operating system should delay packet transmission in hopes of sending fewer packets (Nagle's algorithm). The default is true (no delay), meaning that data is sent as soon as possible after a Write.

​	`SetNoDelay`方法控制操作系统是否应推迟数据包传输，以期望发送较少的数据包(Nagle算法)。默认值为true(无延迟)，这意味着在`Write`之后尽快发送数据。

#### (*TCPConn) SetReadBuffer 

``` go 
func (c *TCPConn) SetReadBuffer(bytes int) error
```

SetReadBuffer sets the size of the operating system's receive buffer associated with the connection.

​	`SetReadBuffer`方法设置与连接关联的操作系统接收缓冲区的大小。

#### (*TCPConn) SetReadDeadline 

``` go 
func (c *TCPConn) SetReadDeadline(t time.Time) error
```

SetReadDeadline implements the Conn SetReadDeadline method.

​	`SetReadDeadline`方法实现`Conn` `SetReadDeadline`方法。

#### (*TCPConn) SetWriteBuffer 

``` go 
func (c *TCPConn) SetWriteBuffer(bytes int) error
```

SetWriteBuffer sets the size of the operating system's transmit buffer associated with the connection.

​	`SetWriteBuffer`方法设置与连接关联的操作系统传输缓冲区的大小。

#### (*TCPConn) SetWriteDeadline 

``` go 
func (c *TCPConn) SetWriteDeadline(t time.Time) error
```

SetWriteDeadline implements the Conn SetWriteDeadline method.

​	`SetWriteDeadline`方法实现`Conn` `SetWriteDeadline`方法。

#### (*TCPConn) SyscallConn  <- go1.9

``` go 
func (c *TCPConn) SyscallConn() (syscall.RawConn, error)
```

SyscallConn returns a raw network connection. This implements the syscall.Conn interface.

​	`SyscallConn`方法返回一个原始的网络连接。这实现了`syscall.Conn`接口。

#### (*TCPConn) Write 

``` go 
func (c *TCPConn) Write(b []byte) (int, error)
```

Write implements the Conn Write method.

​	`Write`方法实现了`Conn`接口的`Write`方法。

### type TCPListener 

``` go 
type TCPListener struct {
	// contains filtered or unexported fields
}
```

TCPListener is a TCP network listener. Clients should typically use variables of type Listener instead of assuming TCP.

​	`TCPListener`结构体是TCP网络侦听器。客户端通常应该使用类型为`Listener`的变量，而不是假定为TCP。

#### func ListenTCP 

``` go 
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
```

ListenTCP acts like Listen for TCP networks.

​	`ListenTCP`函数像TCP网络的`Listen`一样工作。

The network must be a TCP network name; see func Dial for details.

​	网络必须是TCP网络名称;有关详细信息，请参阅func Dial。

If the IP field of laddr is nil or an unspecified IP address, ListenTCP listens on all available unicast and anycast IP addresses of the local system. If the Port field of laddr is 0, a port number is automatically chosen.

​	如果`laddr`的IP字段为nil或未指定的IP地址，则`ListenTCP`将侦听本地系统的所有可用单播和任播IP地址。如果`laddr`的`Port`字段为`0`，则自动选择端口号。

#### (*TCPListener) Accept 

``` go 
func (l *TCPListener) Accept() (Conn, error)
```

Accept implements the Accept method in the Listener interface; it waits for the next call and returns a generic Conn.

​	`Accept`方法实现了 `Listener` 接口的 `Accept` 方法；它等待下一个呼叫并返回一个通用的 `Conn`。

#### (*TCPListener) AcceptTCP 

``` go 
func (l *TCPListener) AcceptTCP() (*TCPConn, error)
```

AcceptTCP accepts the next incoming call and returns the new connection.

​	`AcceptTCP`方法接受下一个传入的呼叫并返回新连接。

#### (*TCPListener) Addr 

``` go 
func (l *TCPListener) Addr() Addr
```

Addr returns the listener's network address, a *TCPAddr. The Addr returned is shared by all invocations of Addr, so do not modify it.

​	`Addr`方法返回监听器的网络地址，即 `*TCPAddr`。返回的 `Addr` 在所有调用 `Addr` 的地方都是共享的，因此不要修改它。

#### (*TCPListener) Close 

``` go 
func (l *TCPListener) Close() error
```

Close stops listening on the TCP address. Already Accepted connections are not closed.

​	`Close`方法停止在 TCP 地址上的监听。已经接受的连接不会关闭。

#### (*TCPListener) File 

``` go 
func (l *TCPListener) File() (f *os.File, err error)
```

File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished. Closing l does not affect f, and closing f does not affect l.

​	`File`方法返回底层的 `os.File` 的副本。调用者有责任在完成后关闭 `f`。关闭 `l` 不影响 `f`，关闭 `f` 不影响 `l`。

The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.

​	返回的 `os.File` 的文件描述符与连接的不同。尝试使用此副本更改原始连接的属性可能会或可能不会产生预期的效果。

#### (*TCPListener) SetDeadline 

``` go 
func (l *TCPListener) SetDeadline(t time.Time) error
```

SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.

​	`SetDeadline`方法设置监听器关联的截止日期。零时间值禁用截止日期。

#### (*TCPListener) SyscallConn  <- go1.10

``` go 
func (l *TCPListener) SyscallConn() (syscall.RawConn, error)
```

SyscallConn returns a raw network connection. This implements the syscall.Conn interface.

​	`SyscallConn`方法返回原始网络连接。这实现了 `syscall.Conn` 接口。

The returned RawConn only supports calling Control. Read and Write return an error.

​	返回的 `RawConn` 仅支持调用 `Control`。`Read` 和 `Write` 返回一个错误。

### type UDPAddr 

``` go 
type UDPAddr struct {
	IP   IP
	Port int
	Zone string // IPv6 范围地址区域 IPv6 scoped addressing zone
}
```

UDPAddr represents the address of a UDP end point.

​	UDPAddr结构体表示 UDP 端点的地址。

#### func ResolveUDPAddr 

``` go 
func ResolveUDPAddr(network, address string) (*UDPAddr, error)
```

ResolveUDPAddr returns an address of UDP end point.

​	`ResolveUDPAddr`函数返回 UDP 终点地址。

The network must be a UDP network name.

​	`network` 必须是 UDP 网络名称。

If the host in the address parameter is not a literal IP address or the port is not a literal port number, ResolveUDPAddr resolves the address to an address of UDP end point. Otherwise, it parses the address as a pair of literal IP address and port number. The address parameter can use a host name, but this is not recommended, because it will return at most one of the host name's IP addresses.

​	如果 `address` 参数中的主机不是字面 IP 地址或端口不是字面端口号，则 `ResolveUDPAddr` 将地址解析为 UDP 终点地址。否则，它将地址解析为字面 IP 地址和端口号的一对。地址参数可以使用主机名，但不建议这样做，因为它最多只会返回主机名的一个 IP 地址。

See func Dial for a description of the network and address parameters.

​	有关网络和地址参数的说明，请参见 func Dial。

#### func UDPAddrFromAddrPort  <- go1.18

``` go 
func UDPAddrFromAddrPort(addr netip.AddrPort) *UDPAddr
```

UDPAddrFromAddrPort returns addr as a UDPAddr. If addr.IsValid() is false, then the returned UDPAddr will contain a nil IP field, indicating an address family-agnostic unspecified address.

​	`UDPAddrFromAddrPort`函数将 `addr` 转换为 `UDPAddr`。如果 `addr.IsValid()` 为 `false`，则返回的 `UDPAddr` 将包含一个 nil IP 字段，表示未指定地址簇的未指定地址。

#### (*UDPAddr) AddrPort  <- go1.18

``` go 
func (a *UDPAddr) AddrPort() netip.AddrPort
```

AddrPort returns the UDPAddr a as a netip.AddrPort.

​	`AddrPort`方法将 `UDPAddr` `a` 转换为 `netip.AddrPort`。

If a.Port does not fit in a uint16, it's silently truncated.

​	如果 `a.Port` 不适合 uint16，则它将被静默截断。

If a is nil, a zero value is returned.

​	如果 `a` 为 nil，则返回零值。

#### (*UDPAddr) Network 

``` go 
func (a *UDPAddr) Network() string
```

Network returns the address's network name, "udp".

​	`Network`方法返回地址的网络名称，即 "udp"。

#### (*UDPAddr) String 

``` go 
func (a *UDPAddr) String() string
```

### type UDPConn 

``` go 
type UDPConn struct {
	// contains filtered or unexported fields
}
```

UDPConn is the implementation of the Conn and PacketConn interfaces for UDP network connections.

​	`UDPConn`结构体是 UDP 网络连接的 `Conn` 和 `PacketConn` 接口实现。

#### func DialUDP 

``` go 
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
```

DialUDP acts like Dial for UDP networks.

​	`DialUDP`函数用于 UDP 网络的 `Dial`。

The network must be a UDP network name; see func Dial for details.

​	`network` 必须是 UDP 网络名称；有关详细信息，请参见 func `Dial`。

If laddr is nil, a local address is automatically chosen. If the IP field of raddr is nil or an unspecified IP address, the local system is assumed.

​	如果 `laddr` 为 nil，则自动选择本地地址。如果 `raddr` 的 IP 字段为 nil 或未指定 IP 地址，则假定为本地系统。

#### func ListenMulticastUDP 

``` go 
func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
```

ListenMulticastUDP acts like ListenPacket for UDP networks but takes a group address on a specific network interface.

​	`ListenMulticastUDP`函数用于 UDP 网络的 `ListenPacket`，但它接受特定网络接口上的组地址。

The network must be a UDP network name; see func Dial for details.

​	`network` 必须是 UDP 网络名称；有关详细信息，请参见 func `Dial`。

ListenMulticastUDP listens on all available IP addresses of the local system including the group, multicast IP address. If ifi is nil, ListenMulticastUDP uses the system-assigned multicast interface, although this is not recommended because the assignment depends on platforms and sometimes it might require routing configuration. If the Port field of gaddr is 0, a port number is automatically chosen.

​	`ListenMulticastUDP` 监听本地系统的所有可用 IP 地址，包括组播 IP 地址。如果 ifi 为 nil，则 ListenMulticastUDP 使用系统分配的多播接口，尽管这不被推荐，因为分配取决于平台，有时可能需要路由配置。如果 gaddr 的 Port 字段为 0，则自动选择端口号。

ListenMulticastUDP is just for convenience of simple, small applications. There are golang.org/x/net/ipv4 and golang.org/x/net/ipv6 packages for general purpose uses.

​	`ListenMulticastUDP` 只是简单、小型应用的方便。对于一般用途，有 golang.org/x/net/ipv4 和 golang.org/x/net/ipv6 包。

Note that ListenMulticastUDP will set the IP_MULTICAST_LOOP socket option to 0 under IPPROTO_IP, to disable loopback of multicast packets.

​	请注意，`ListenMulticastUDP`将在IPPROTO_IP下将IP_MULTICAST_LOOP套接字选项设置为0，以禁用组播数据包的回送。

#### func ListenUDP 

``` go 
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
```

ListenUDP acts like ListenPacket for UDP networks.

​	`ListenUDP`函数的行为类似于UDP网络的`ListenPacket`。

The network must be a UDP network name; see func Dial for details.

​	网络必须是UDP网络名称；有关详细信息，请参见Dial函数。

If the IP field of laddr is nil or an unspecified IP address, ListenUDP listens on all available IP addresses of the local system except multicast IP addresses. If the Port field of laddr is 0, a port number is automatically chosen.

​	如果`laddr`的IP字段为nil或未指定的IP地址，则`ListenUDP`会在本地系统的所有可用IP地址上进行侦听，但不包括多播IP地址。如果`laddr`的`Port`字段为`0`，则自动选择端口号。

#### (*UDPConn) Close 

``` go 
func (c *UDPConn) Close() error
```

Close closes the connection.

​	`Close`方法关闭连接。

#### (*UDPConn) File 

``` go 
func (c *UDPConn) File() (f *os.File, err error)
```

File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished. Closing c does not affect f, and closing f does not affect c.

​	`File`方法返回底层os.File的副本。调用者有责任在完成后关闭f。关闭c不会影响f，关闭f也不会影响c。

The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.

​	返回的`os.File`的文件描述符与连接的不同。尝试使用此副本更改原始的属性可能会产生预期或非预期的影响。

#### (*UDPConn) LocalAddr 

``` go 
func (c *UDPConn) LocalAddr() Addr
```

LocalAddr returns the local network address. The Addr returned is shared by all invocations of LocalAddr, so do not modify it.

​	`LocalAddr`方法返回本地网络地址。返回的`Addr`由`LocalAddr`的所有调用共享，因此不要修改它。

#### (*UDPConn) Read 

``` go 
func (c *UDPConn) Read(b []byte) (int, error)
```

Read implements the Conn Read method.

​	`Read`方法实现Conn `Read`方法。

#### (*UDPConn) ReadFrom 

``` go 
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)
```

ReadFrom implements the PacketConn ReadFrom method.

​	`ReadFrom`方法实现PacketConn `ReadFrom`方法。

#### (*UDPConn) ReadFromUDP 

``` go 
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)
```

ReadFromUDP acts like ReadFrom but returns a UDPAddr.

​	`ReadFromUDP`方法的行为类似于`ReadFrom`，但返回一个`UDPAddr`。

#### (*UDPConn) ReadFromUDPAddrPort  <- go1.18

``` go 
func (c *UDPConn) ReadFromUDPAddrPort(b []byte) (n int, addr netip.AddrPort, err error)
```

ReadFromUDPAddrPort acts like ReadFrom but returns a netip.AddrPort.

​	`ReadFromUDPAddrPort`方法的行为类似于ReadFrom，但返回一个`netip.AddrPort`。

If c is bound to an unspecified address, the returned netip.AddrPort's address might be an IPv4-mapped IPv6 address. Use netip.Addr.Unmap to get the address without the IPv6 prefix.

​	如果`c`绑定到未指定的地址，则返回的`netip.AddrPort`的地址可能是一个IPv4映射的IPv6地址。使用`netip.Addr.Unmap`获取不带IPv6前缀的地址。

#### (*UDPConn) ReadMsgUDP  <- go1.1

``` go 
func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)
```

ReadMsgUDP reads a message from c, copying the payload into b and the associated out-of-band data into oob. It returns the number of bytes copied into b, the number of bytes copied into oob, the flags that were set on the message and the source address of the message.

​	`ReadMsgUDP`方法从`c`读取消息，将有效负载复制到`b`中，并将关联的带外数据复制到`oob`中。它返回复制到`b`中的字节数，复制到`oob`中的字节数，设置在消息上的标志以及消息的源地址。

The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be used to manipulate IP-level socket options in oob.

​	可以使用golang.org/x/net/ipv4和golang.org/x/net/ipv6包来操作oob中的IP级套接字选项。

#### (*UDPConn) ReadMsgUDPAddrPort  <- go1.18

``` go 
func (c *UDPConn) ReadMsgUDPAddrPort(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error)
```

ReadMsgUDPAddrPort is like ReadMsgUDP but returns an netip.AddrPort instead of a UDPAddr.

​	`ReadMsgUDPAddrPort`方法与`ReadMsgUDP`方法类似，但返回的是`netip.AddrPort`而不是`UDPAddr`。

#### (*UDPConn) RemoteAddr 

``` go 
func (c *UDPConn) RemoteAddr() Addr
```

RemoteAddr returns the remote network address. The Addr returned is shared by all invocations of RemoteAddr, so do not modify it.

​	`RemoteAddr`方法返回远程网络地址。返回的 `Addr` 在所有 `RemoteAddr` 的调用中共享，因此不要修改它。

#### (*UDPConn) SetDeadline 

``` go 
func (c *UDPConn) SetDeadline(t time.Time) error
```

SetDeadline implements the Conn SetDeadline method.

​	`SetDeadline`方法实现了 Conn `SetDeadline` 方法。

#### (*UDPConn) SetReadBuffer 

``` go 
func (c *UDPConn) SetReadBuffer(bytes int) error
```

SetReadBuffer sets the size of the operating system's receive buffer associated with the connection.

​	`SetReadBuffer`方法设置与连接关联的操作系统接收缓冲区的大小。

#### (*UDPConn) SetReadDeadline 

``` go 
func (c *UDPConn) SetReadDeadline(t time.Time) error
```

SetReadDeadline implements the Conn SetReadDeadline method.

​	`SetReadDeadline`方法实现了 Conn `SetReadDeadline` 方法。

#### (*UDPConn) SetWriteBuffer 

``` go 
func (c *UDPConn) SetWriteBuffer(bytes int) error
```

SetWriteBuffer sets the size of the operating system's transmit buffer associated with the connection.

​	`SetWriteBuffer`方法设置与连接关联的操作系统传输缓冲区的大小。

#### (*UDPConn) SetWriteDeadline 

``` go 
func (c *UDPConn) SetWriteDeadline(t time.Time) error
```

SetWriteDeadline implements the Conn SetWriteDeadline method.

​	`SetWriteDeadline`方法实现了 Conn `SetWriteDeadline` 方法。

#### (*UDPConn) SyscallConn  <- go1.9

``` go 
func (c *UDPConn) SyscallConn() (syscall.RawConn, error)
```

SyscallConn returns a raw network connection. This implements the syscall.Conn interface.

​	`SyscallConn`方法返回原始网络连接。这实现了 `syscall.Conn` 接口。

#### (*UDPConn) Write 

``` go 
func (c *UDPConn) Write(b []byte) (int, error)
```

Write implements the Conn Write method.

​	Write方法实现了 Conn `Write` 方法。

#### (*UDPConn) WriteMsgUDP  <- go1.1

``` go 
func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)
```

WriteMsgUDP writes a message to addr via c if c isn't connected, or to c's remote address if c is connected (in which case addr must be nil). The payload is copied from b and the associated out-of-band data is copied from oob. It returns the number of payload and out-of-band bytes written.

​	`WriteMsgUDP`方法向 `addr` 写入消息，如果 `c` 未连接，则通过 `c` 写入，否则通过 `c` 的远程地址写入(此时 `addr` 必须为 nil)。从 `b` 复制有效载荷，从 `oob` 复制关联的带外数据。返回写入的有效载荷和带外字节数。

The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be used to manipulate IP-level socket options in oob.

​	 golang.org/x/net/ipv4 包和 golang.org/x/net/ipv6 包可用于操作 `oob` 中的 IP 级 `socket` 选项。

#### (*UDPConn) WriteMsgUDPAddrPort  <- go1.18

``` go 
func (c *UDPConn) WriteMsgUDPAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error)
```

WriteMsgUDPAddrPort is like WriteMsgUDP but takes a netip.AddrPort instead of a UDPAddr.

​	`WriteMsgUDPAddrPort`方法与`WriteMsgUDP`方法类似，但是接受`netip.AddrPort`而不是`UDPAddr`。

#### (*UDPConn) WriteTo 

``` go 
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)
```

WriteTo implements the PacketConn WriteTo method.

​	`WriteTo`方法实现了`PacketConn` `WriteTo`方法。

##### WriteTo Example
``` go 
package main

import (
	"log"
	"net"
)

func main() {
	// Unlike Dial, ListenPacket creates a connection without any
	// association with peers.
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dst, err := net.ResolveUDPAddr("udp", "192.0.2.1:2000")
	if err != nil {
		log.Fatal(err)
	}

	// The connection can write data to the desired address.
	_, err = conn.WriteTo([]byte("data"), dst)
	if err != nil {
		log.Fatal(err)
	}
}

```

#### (*UDPConn) WriteToUDP 

``` go 
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
```

WriteToUDP acts like WriteTo but takes a UDPAddr.

​	`WriteToUDP`方法类似于`WriteTo`，但是需要一个`UDPAddr`参数。

#### (*UDPConn) WriteToUDPAddrPort  <- go1.18

``` go 
func (c *UDPConn) WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error)
```

WriteToUDPAddrPort acts like WriteTo but takes a netip.AddrPort.

​	`WriteToUDPAddrPort`方法类似于`WriteTo`，但是需要一个`netip.AddrPort`参数。

### type UnixAddr 

``` go 
type UnixAddr struct {
	Name string
	Net  string
}
```

UnixAddr represents the address of a Unix domain socket end point.

​	`UnixAddr`结构体表示Unix域套接字终端的地址。

#### func ResolveUnixAddr 

``` go 
func ResolveUnixAddr(network, address string) (*UnixAddr, error)
```

ResolveUnixAddr returns an address of Unix domain socket end point.

​	`ResolveUnixAddr`函数返回Unix域套接字终端的地址。

The network must be a Unix network name.

​	`network`参数必须是Unix网络名称。

See func Dial for a description of the network and address parameters.

​	有关`network`和`address`参数的说明，请参见func Dial。

#### (*UnixAddr) Network 

``` go 
func (a *UnixAddr) Network() string
```

Network returns the address's network name, "unix", "unixgram" or "unixpacket".

​	`Network`方法返回地址的网络名称，即"unix"、"unixgram"或"unixpacket"。

#### (*UnixAddr) String 

``` go 
func (a *UnixAddr) String() string
```

### type UnixConn 

``` go 
type UnixConn struct {
	// 包含过滤或未公开的字段
}
```

UnixConn is an implementation of the Conn interface for connections to Unix domain sockets.

​	`UnixConn`结构体是用于Unix域套接字连接的`Conn`接口的实现。

#### func DialUnix 

``` go 
func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)
```

DialUnix acts like Dial for Unix networks.

​	`DialUnix`函数类似于Unix网络的`Dial`函数。

The network must be a Unix network name; see func Dial for details.

​	`network`参数必须是Unix网络名称；有关详细信息，请参见func `Dial`。

If laddr is non-nil, it is used as the local address for the connection.

​	如果`laddr`非nil，则用作连接的本地地址。

#### func ListenUnixgram 

``` go 
func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)
```

ListenUnixgram acts like ListenPacket for Unix networks.

​	`ListenUnixgram`函数类似于Unix网络的`ListenPacket`函数。

The network must be "unixgram".

​	`network`参数必须是"unixgram"。

#### (*UnixConn) Close 

``` go 
func (c *UnixConn) Close() error
```

Close closes the connection.

​	`Close`方法关闭连接。

#### (*UnixConn) CloseRead  <- go1.1

``` go 
func (c *UnixConn) CloseRead() error
```

CloseRead shuts down the reading side of the Unix domain connection. Most callers should just use Close.

​	`CloseRead`方法关闭Unix域连接的读取端。大多数调用方应该只使用`Close`。

#### (*UnixConn) CloseWrite  <- go1.1

``` go 
func (c *UnixConn) CloseWrite() error
```

CloseWrite shuts down the writing side of the Unix domain connection. Most callers should just use Close.

​	`CloseWrite`方法关闭Unix域连接的写入端。大多数调用方应该只使用`Close`。

#### (*UnixConn) File 

``` go 
func (c *UnixConn) File() (f *os.File, err error)
```

File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished. Closing c does not affect f, and closing f does not affect c.

​	`File`方法返回底层`os.File`的副本。调用方有责任在完成后关闭f。关闭c不会影响f，关闭f也不会影响c。

The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.

​	返回的`os.File`的文件描述符与连接的不同。尝试使用此副本更改原始文件的属性可能会或可能不会产生所需的效果。

#### (*UnixConn) LocalAddr 

``` go 
func (c *UnixConn) LocalAddr() Addr
```

LocalAddr returns the local network address. The Addr returned is shared by all invocations of LocalAddr, so do not modify it.

​	`LocalAddr`方法返回本地网络地址。返回的`Addr`由`LocalAddr`的所有调用共享，因此不要修改它。

#### (*UnixConn) Read 

``` go 
func (c *UnixConn) Read(b []byte) (int, error)
```

Read implements the Conn Read method.

​	`Read`方法实现了Conn `Read`方法。

#### (*UnixConn) ReadFrom 

``` go 
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)
```

ReadFrom implements the PacketConn ReadFrom method.

​	`ReadFrom`方法实现了`PacketConn` `ReadFrom`方法。

#### (*UnixConn) ReadFromUnix 

``` go 
func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error)
```

ReadFromUnix acts like ReadFrom but returns a UnixAddr.

​	`ReadFromUnix`方法类似于`ReadFrom`，但返回一个`UnixAddr`。

#### (*UnixConn) ReadMsgUnix 

``` go 
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)
```

ReadMsgUnix reads a message from c, copying the payload into b and the associated out-of-band data into oob. It returns the number of bytes copied into b, the number of bytes copied into oob, the flags that were set on the message and the source address of the message.

​	`ReadMsgUnix`方法从`c`读取消息，将有效负载复制到`b`，将相关的带外数据复制到`oob`。它返回复制到`b`的字节数，复制到`oob`的字节数，消息上设置的标志以及消息的源地址。

Note that if len(b) == 0 and len(oob) > 0, this function will still read (and discard) 1 byte from the connection.

​	请注意，如果`len(b) == 0`且`len(oob) > 0`，则此函数仍将从连接中读取(并丢弃)1个字节。

#### (*UnixConn) RemoteAddr 

``` go 
func (c *UnixConn) RemoteAddr() Addr
```

RemoteAddr returns the remote network address. The Addr returned is shared by all invocations of RemoteAddr, so do not modify it.

​	`RemoteAddr`方法返回远程网络地址。返回的`Addr`由`RemoteAddr`的所有调用共享，因此不要修改它。

#### (*UnixConn) SetDeadline 

``` go 
func (c *UnixConn) SetDeadline(t time.Time) error
```

SetDeadline implements the Conn SetDeadline method.

​	`SetDeadline`方法实现了`Conn` `SetDeadline`方法。

#### (*UnixConn) SetReadBuffer 

``` go 
func (c *UnixConn) SetReadBuffer(bytes int) error
```

SetReadBuffer sets the size of the operating system's receive buffer associated with the connection.

​	`SetReadBuffer`方法设置与连接关联的操作系统接收缓冲区的大小。

#### (*UnixConn) SetReadDeadline 

``` go 
func (c *UnixConn) SetReadDeadline(t time.Time) error
```

SetReadDeadline implements the Conn SetReadDeadline method.

​	`SetReadDeadline`方法实现了`Conn` `SetReadDeadline`方法。

#### (*UnixConn) SetWriteBuffer 

``` go 
func (c *UnixConn) SetWriteBuffer(bytes int) error
```

SetWriteBuffer sets the size of the operating system's transmit buffer associated with the connection.

​	`SetWriteBuffer`方法设置与连接关联的操作系统传输缓冲区的大小。

#### (*UnixConn) SetWriteDeadline 

``` go 
func (c *UnixConn) SetWriteDeadline(t time.Time) error
```

SetWriteDeadline implements the Conn SetWriteDeadline method.

​	`SetWriteDeadline`方法实现了`Conn` `SetWriteDeadline`方法。

#### (*UnixConn) SyscallConn  <- go1.9

``` go 
func (c *UnixConn) SyscallConn() (syscall.RawConn, error)
```

SyscallConn returns a raw network connection. This implements the syscall.Conn interface.

​	`SyscallConn`方法返回原始网络连接。这实现了`syscall.Conn`接口。

#### (*UnixConn) Write 

``` go 
func (c *UnixConn) Write(b []byte) (int, error)
```

Write implements the Conn Write method.

​	`Write`方法实现了`Conn` `Write`方法。

#### (*UnixConn) WriteMsgUnix 

``` go 
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)
```

WriteMsgUnix writes a message to addr via c, copying the payload from b and the associated out-of-band data from oob. It returns the number of payload and out-of-band bytes written.

​	`WriteMsgUnix`方法向`addr`通过`c`写入消息，从`b`复制有效负载，从`oob`复制相关的带外数据。它返回已写入有效负载和带外字节数。

Note that if len(b) == 0 and len(oob) > 0, this function will still write 1 byte to the connection.

​	请注意，如果`len(b) == 0`且`len(oob) > 0`，则此函数仍将向连接写入1个字节。

#### (*UnixConn) WriteTo 

``` go 
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error)
```

WriteTo implements the PacketConn WriteTo method.

​	`WriteTo`方法实现了`PacketConn` `WriteTo`方法。

#### (*UnixConn) WriteToUnix 

``` go 
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error)
```

WriteToUnix acts like WriteTo but takes a UnixAddr.

​	`WriteToUnix`方法类似于`WriteTo`，但需要一个`UnixAddr`参数。

### type UnixListener 

``` go 
type UnixListener struct {
	// 包含过滤或未公开的字段
}
```

UnixListener is a Unix domain socket listener. Clients should typically use variables of type Listener instead of assuming Unix domain sockets.

​	`UnixListener`结构体是Unix域套接字侦听器。客户端通常应使用类型为`Listener`的变量，而不是假定Unix域套接字。

#### func ListenUnix 

``` go 
func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)
```

ListenUnix acts like Listen for Unix networks.

​	`ListenUnix`函数类似于Unix网络的`Listen`函数。

The network must be "unix" or "unixpacket".

​	`network`必须是"unix"或"unixpacket"。

#### (*UnixListener) Accept 

``` go 
func (l *UnixListener) Accept() (Conn, error)
```

Accept implements the Accept method in the Listener interface. Returned connections will be of type *UnixConn.

​	`Accept`方法实现了`Listener`接口中的`Accept`方法。返回的连接将是`*UnixConn`类型。

#### (*UnixListener) AcceptUnix 

``` go 
func (l *UnixListener) AcceptUnix() (*UnixConn, error)
```

AcceptUnix accepts the next incoming call and returns the new connection.

​	`AcceptUnix`方法接受下一个传入的调用并返回新的连接。

#### (*UnixListener) Addr 

``` go 
func (l *UnixListener) Addr() Addr
```

Addr returns the listener's network address. The Addr returned is shared by all invocations of Addr, so do not modify it.

​	`Addr`方法返回侦听器的网络地址。返回的`Addr`由所有`Addr`调用共享，因此不要修改它。

#### (*UnixListener) Close 

``` go 
func (l *UnixListener) Close() error
```

Close stops listening on the Unix address. Already accepted connections are not closed.

​	`Close`方法停止侦听Unix地址。已接受的连接不会关闭。

#### (*UnixListener) File 

``` go 
func (l *UnixListener) File() (f *os.File, err error)
```

File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished. Closing l does not affect f, and closing f does not affect l.

​	`File`方法返回底层`os.File`的副本。调用者有责任在完成后关闭`f`。关闭`l`不影响`f`，关闭f也不影响`l`。

The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.

​	返回的`os.File`的文件描述符与连接的不同。尝试使用此副本更改原始文件的属性可能会产生预期的效果，也可能不会。

#### (*UnixListener) SetDeadline 

``` go 
func (l *UnixListener) SetDeadline(t time.Time) error
```

SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.

​	`SetDeadline`方法设置与侦听器关联的截止日期。零时间值禁用截止日期。

#### (*UnixListener) SetUnlinkOnClose  <- go1.8

``` go 
func (l *UnixListener) SetUnlinkOnClose(unlink bool)
```

SetUnlinkOnClose sets whether the underlying socket file should be removed from the file system when the listener is closed.

​	`SetUnlinkOnClose`方法设置在关闭侦听器时是否应从文件系统中删除底层套接字文件。

The default behavior is to unlink the socket file only when package net created it. That is, when the listener and the underlying socket file were created by a call to Listen or ListenUnix, then by default closing the listener will remove the socket file. but if the listener was created by a call to FileListener to use an already existing socket file, then by default closing the listener will not remove the socket file.

​	默认行为是仅在包net创建它时取消连接套接字文件。也就是说，当使用调用`Listen`或`ListenUnix`创建侦听器和底层套接字文件时，默认情况下，关闭侦听器将删除套接字文件。但是，如果通过调用`FileListener`创建侦听器来使用已经存在的套接字文件，则默认情况下，关闭侦听器不会删除套接字文件。

#### (*UnixListener) SyscallConn  <- go1.10

``` go 
func (l *UnixListener) SyscallConn() (syscall.RawConn, error)
```

SyscallConn returns a raw network connection. This implements the syscall.Conn interface.

​	`SyscallConn`方法返回原始网络连接。这实现了`syscall.Conn`接口。

The returned RawConn only supports calling Control. Read and Write return an error.

​	返回的`RawConn`仅支持调用`Control`。读取和写入返回错误。

### type UnknownNetworkError 

``` go 
type UnknownNetworkError string
```

#### (UnknownNetworkError) Error 

``` go 
func (e UnknownNetworkError) Error() string
```

#### (UnknownNetworkError) Temporary 

``` go 
func (e UnknownNetworkError) Temporary() bool
```

#### (UnknownNetworkError) Timeout 

``` go 
func (e UnknownNetworkError) Timeout() bool
```

## Notes

### Bugs

- On JS and Windows, the FileConn, FileListener and FilePacketConn functions are not implemented.

- 在JS和Windows上，未实现FileConn、FileListener和FilePacketConn函数。

- On JS, methods and functions related to Interface are not implemented.

- 在JS上，未实现与Interface相关的方法和函数。

- On AIX, DragonFly BSD, NetBSD, OpenBSD, Plan 9 and Solaris, the MulticastAddrs method of Interface is not implemented.

- 在AIX、DragonFly BSD、NetBSD、OpenBSD、Plan 9和Solaris上，未实现Interface的MulticastAddrs方法。

- On every POSIX platform, reads from the "ip4" network using the ReadFrom or ReadFromIP method might not return a complete IPv4 packet, including its header, even if there is space available. This can occur even in cases where Read or ReadMsgIP could return a complete packet. For this reason, it is recommended that you do not use these methods if it is important to receive a full packet.

- 在每个POSIX平台上，使用ReadFrom或ReadFromIP方法从"ip4"网络读取可能不会返回完整的IPv4数据包，包括其标头，即使有空间可用。即使Read或ReadMsgIP可以返回完整的数据包，这也可能发生。因此，如果接收完整数据包很重要，则建议不使用这些方法。

  The Go 1 compatibility guidelines make it impossible for us to change the behavior of these methods; use Read or ReadMsgIP instead.

  Go 1兼容性指南使我们无法更改这些方法的行为。请改用Read或ReadMsgIP。

- On JS and Plan 9, methods and functions related to IPConn are not implemented.

- 在JS和Plan 9上，未实现与IPConn相关的方法和函数。

- On Windows, the File method of IPConn is not implemented.

- 在Windows上，未实现IPConn的File方法。

- On DragonFly BSD and OpenBSD, listening on the "tcp" and "udp" networks does not listen for both IPv4 and IPv6 connections. This is due to the fact that IPv4 traffic will not be routed to an IPv6 socket - two separate sockets are required if both address families are to be supported. See inet6(4) for details.

- 在DragonFly BSD和OpenBSD上，侦听"tcp"和"udp"网络不会同时侦听IPv4和IPv6连接。这是因为IPv4流量不会被路由到IPv6套接字——如果要支持两个地址家族，则需要两个单独的套接字。有关详细信息，请参见inet6(4)。

- On Windows, the Write method of syscall.RawConn does not integrate with the runtime's network poller. It cannot wait for the connection to become writeable, and does not respect deadlines. If the user-provided callback returns false, the Write method will fail immediately.

- 在Windows上，syscall.RawConn的Write方法无法与运行时的网络轮询器集成。它不能等待连接变为可写，并且不尊重期限。如果用户提供的回调返回false，则Write方法将立即失败。

- On JS and Plan 9, the Control, Read and Write methods of syscall.RawConn are not implemented.

- 在JS和Plan 9上，未实现syscall.RawConn的Control、Read和Write方法。

- On JS and Windows, the File method of TCPConn and TCPListener is not implemented.

- 在JS和Windows上，未实现TCPConn和TCPListener的File方法。

- On Plan 9, the ReadMsgUDP and WriteMsgUDP methods of UDPConn are not implemented.

- 在Plan 9上，未实现UDPConn的ReadMsgUDP和WriteMsgUDP方法。

- On Windows, the File method of UDPConn is not implemented.

- 在Windows上，未实现UDPConn的File方法。

- On JS, methods and functions related to UDPConn are not implemented.

- 在JS上，未实现与UDPConn相关的方法和函数。

- On JS, WASIP1 and Plan 9, methods and functions related to UnixConn and UnixListener are not implemented.

- 在JS和Plan 9上，未实现与UnixConn和UnixListener相关的方法和函数。

- On Windows, methods and functions related to UnixConn and UnixListener don't work for "unixgram" and "unixpacket".

- 在Windows上，与UnixConn和UnixListener相关的方法和函数不适用于"unixgram"和"unixpacket"。