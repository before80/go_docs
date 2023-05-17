+++
title = "net"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# net

[https://pkg.go.dev/net@go1.20.1](https://pkg.go.dev/net@go1.20.1)

​	net包提供了一个可移植的网络 I/O 接口，包括 TCP/IP、UDP、域名解析和 Unix 域套接字。

​	虽然该包提供了对低级网络原语的访问，但大多数客户端只需要 Dial、Listen 和 Accept 函数以及相关的 Conn 和 Listener 接口所提供的基本接口。crypto/tls 包使用相同的接口和类似的 Dial 和 Listen 函数。

​	Dial 函数连接到服务器：

```
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
	// 处理错误
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
// ...
```

​	Listen 函数创建服务器：

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

#### Name 解析 

​	解析域名的方法(无论是间接使用 Dial 等函数还是直接使用 LookupHost 和 LookupAddr 等函数)因操作系统而异。

​	在 Unix 系统上，解析器有两个选项来解析名称。它可以使用纯 Go 解析器，该解析器将 DNS 请求直接发送到 /etc/resolv.conf 中列出的服务器，或者它可以使用基于 cgo 的解析器，该解析器调用 C 库例程，例如 getaddrinfo 和 getnameinfo。

​	默认情况下，使用纯 Go 解析器，因为阻塞的 DNS 请求只会消耗一个 goroutine，而阻塞的 C 调用会消耗一个操作系统线程。当 cgo 可用时，在多种条件下将改为使用基于 cgo 的解析器：在不允许程序直接进行 DNS 请求的系统上(OS X)，当存在 LOCALDOMAIN 环境变量时(即使为空)，当 RES_OPTIONS 或 HOSTALIASES 环境变量非空时，当 ASR_CONFIG 环境变量非空时(仅在 OpenBSD 上)，当 /etc/resolv.conf 或 /etc/nsswitch.conf 指定使用 Go 解析器未实现的功能时，以及正在查找的名称以 .local 结尾或是 mDNS 名称时。

​	可以通过将 GODEBUG 环境变量(参见 package runtime)的 netdns 值设置为 go 或 cgo 来覆盖解析器决策，例如：

```
export GODEBUG=netdns=go    # force pure Go resolver
export GODEBUG=netdns=cgo   # force native resolver (cgo, win32)
```

​	在构建 Go 源代码树时，也可以通过设置 netgo 或 netcgo 构建标签来强制执行决策。

​	数字netdns设置，例如GODEBUG=netdns=1，会导致解析器打印有关其决策的调试信息。为了在打印调试信息的同时强制使用特定的解析器，请将两个设置连接起来，如GODEBUG=netdns=go+1。

​	在Plan 9上，解析器总是访问/net/cs和/net/dns。

​	在早期的Go 1.18.x及之前的Windows版本中，解析器总是使用C库函数，例如GetAddrInfo和DnsQuery。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=21)

``` go linenums="1"
const (
	IPv4len = 4
	IPv6len = 16
)
```

IP地址长度(以字节为单位)。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=100)

``` go linenums="1"
var (
	IPv4bcast     = IPv4(255, 255, 255, 255) // 有限广播
	IPv4allsys    = IPv4(224, 0, 0, 1)       // 所有系统
	IPv4allrouter = IPv4(224, 0, 0, 2)       // 所有路由器
	IPv4zero      = IPv4(0, 0, 0, 0)         // 全部为零
)
```

知名IPv4地址。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=108)

``` go linenums="1"
var (
	IPv6zero                   = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6unspecified            = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6loopback               = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	IPv6interfacelocalallnodes = IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallnodes      = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallrouters    = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)
```

知名IPv6地址。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=118)

``` go linenums="1"
var DefaultResolver = &Resolver{}
```

​	DefaultResolver 是由包级别的 Lookup 函数和没有指定 Resolver 的 Dialer 使用的解析器。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=665)

``` go linenums="1"
var ErrClosed error = errClosed
```

​	ErrClosed 是在网络连接上进行 I/O 调用时返回的错误，该连接已经关闭，或者在 I/O 完成之前被另一个 goroutine 关闭。这可能会包装在另一个错误中，并且通常应使用 errors.Is(err，net.ErrClosed)进行测试。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=408)

``` go linenums="1"
var (
	ErrWriteToConnected = errors.New("use of WriteTo with pre-connected connection")
)
```

​	OpError 中包含的各种错误。

## 函数

#### func [JoinHostPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ipsock.go;l=235) 

``` go linenums="1"
func JoinHostPort(host, port string) string
```

​	JoinHostPort函数将 host 和 port 组合成形如 "host:port" 的网络地址。如果 host 包含冒号，如字面上的 IPv6 地址，则 JoinHostPort函数返回 "[host]:port"。

​	有关 host 和 port 参数的说明，请参见 func Dial。

#### func [LookupAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=645) 

``` go linenums="1"
func LookupAddr(addr string) (names []string, err error)
```

​	LookupAddr函数根据给定的地址执行反向查找，返回映射到该地址的名称列表。

​	返回的名称已验证为格式正确的表示格式域名。如果响应包含无效的名称，则会过滤掉这些记录，并在剩余结果(如果有)旁边返回错误。

​	使用主机 C 库解析程序时，最多只会返回一个结果。要绕过主机解析程序，请使用自定义解析程序。

​	LookupAddr函数在内部使用 context.Background；要指定上下文，请使用 Resolver.LookupAddr。

#### func [LookupCNAME](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=456) 

``` go linenums="1"
func LookupCNAME(host string) (cname string, err error)
```

​	LookupCNAME函数返回给定主机的规范名称。不关心规范名称的调用方可以直接调用 LookupHost函数或 LookupIP函数；两者都会在查找中解析规范名称。

​	规范名称是在跟随零个或多个 CNAME 记录之后的最终名称。如果主机不包含 DNS "CNAME" 记录，则 LookupCNAME 不会返回错误，只要主机解析到地址记录即可。

​	返回的规范名称已验证为格式正确的表示格式域名。

​	LookupCNAME函数在内部使用 context.Background；要指定上下文，请使用 Resolver.LookupCNAME。

#### func [LookupHost](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=176) 

``` go linenums="1"
func LookupHost(host string) (addrs []string, err error)
```

​	LookupHost函数使用本地解析器查找给定的主机名。它返回该主机的地址列表。

​	LookupHost函数在内部使用context.Background；要指定上下文，请使用Resolver.LookupHost。

#### func [LookupPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=414) 

``` go linenums="1"
func LookupPort(network, service string) (port int, err error)
```

​	LookupPort函数查找给定网络和服务的端口。

​	LookupPort函数在内部使用context.Background；要指定上下文，请使用Resolver.LookupPort。

#### func [LookupTXT](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=624) 

``` go linenums="1"
func LookupTXT(name string) ([]string, error)
```

​	LookupTXT函数返回给定域名的DNS TXT记录。

​	LookupTXT函数在内部使用context.Background；要指定上下文，请使用Resolver.LookupTXT。

#### func [ParseCIDR](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=741) 

``` go linenums="1"
func ParseCIDR(s string) (IP, *IPNet, error)
```

​	ParseCIDR函数将s解析为CIDR表示法的IP地址和前缀长度，例如"192.0.2.0/24"或"2001:db8::/32"，如RFC 4632和RFC 4291中所定义。

​	它返回IP地址和由IP和前缀长度隐含的网络。例如，ParseCIDR("192.0.2.1/24")返回IP地址192.0.2.1和网络192.0.2.0/24。

##### ParseCIDR Example
``` go linenums="1"
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

#### func [Pipe](https://cs.opensource.google/go/go/+/go1.20.1:src/net/pipe.go;l=113) 

``` go linenums="1"
func Pipe() (Conn, Conn)
```

​	Pipe函数创建一个同步的内存中全双工网络连接；两端都实现了Conn接口。一个端口的读取与另一个端口的写入匹配，直接在两者之间复制数据；没有内部缓冲。

#### func [SplitHostPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ipsock.go;l=164) 

``` go linenums="1"
func SplitHostPort(hostport string) (host, port string, err error)
```

​	SplitHostPort函数将"`host:port`"、"`host%zone:port`"、"`[host]:port`"或"`[host％zone]:port`"形式的网络地址分成主机或host％zone和端口。

​	hostport中的文字IPv6地址必须用方括号括起来，如"`[::1]:80`"、"`[::1％lo0]:80`"。

​	有关hostport参数和host和port结果的说明，请参见func Dial。

## 类型

### type [Addr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=105) 

``` go linenums="1"
type Addr interface {
	Network() string // 网络的名称(例如 "tcp"、"udp")
	String() string  // 地址的字符串表示形式(例如 "192.0.2.1:25"、"[2001:db8::1]:80")
}
```

​	Addr接口表示网络终端地址。

​	Network 和 String 两个方法一般返回可以作为 Dial 函数参数的字符串，但字符串的确切格式和含义取决于实现。

#### func [InterfaceAddrs](https://cs.opensource.google/go/go/+/go1.20.1:src/net/interface.go;l=118) 

``` go linenums="1"
func InterfaceAddrs() ([]Addr, error)
```

​	InterfaceAddrs函数返回系统的单播接口地址列表。

​	返回的列表不包含关联接口的标识，使用 Interfaces 和 Interface.Addrs 获取更多详细信息。

### type [AddrError](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=554) 

``` go linenums="1"
type AddrError struct {
	Err  string
	Addr string
}
```

#### (*AddrError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=559) 

``` go linenums="1"
func (e *AddrError) Error() string
```

#### (*AddrError) [Temporary](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=571) 

``` go linenums="1"
func (e *AddrError) Temporary() bool
```

#### (*AddrError) [Timeout](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=570) 

``` go linenums="1"
func (e *AddrError) Timeout() bool
```

### type [Buffers](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=712)  <- go1.8

``` go linenums="1"
type Buffers [][]byte
```

​	Buffers 包含零个或多个要写入的字节序列。

​	在某些机器上，对于某些类型的连接，这可以被优化为操作系统特定的批量写操作(例如 "writev")。

#### (*Buffers) [Read](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=747)  <- go1.8

``` go linenums="1"
func (v *Buffers) Read(p []byte) (n int, err error)
```

​	Read方法从缓冲区读取。

​	Read方法为 Buffers 实现了 io.Reader。

​	Read方法修改了切片 v 和 v[i]，其中 0 <= i < len(v)，但不修改 v[i][j]，其中 i 和 j 为任意值。

#### (*Buffers) [WriteTo](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=725)  <- go1.8

``` go linenums="1"
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error)
```

​	WriteTo方法将缓冲区的内容写入 w。

​	WriteTo方法为 Buffers 实现了 io.WriterTo。

​	WriteTo方法修改了切片 v 和 v[i]，其中 0 <= i < len(v)，但不修改 v[i][j]，其中 i 和 j 为任意值。

### type [Conn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=113) 

``` go linenums="1"
type Conn interface {
	// Read 从连接中读取数据。
	// Read 可以设置超时时间，在固定时间限制后超时并返回错误；
    // 参见 SetDeadline 和 	SetReadDeadline。
	Read(b []byte) (n int, err error)

	// Write 向连接中写入数据。
	// Write 可以设置超时时间，在固定时间限制后超时并返回错误；
    // 参见 SetDeadline 和 SetWriteDeadline。
	Write(b []byte) (n int, err error)

	// Close 关闭连接。
	// 任何被阻塞的 Read 或 Write 操作都将取消阻塞并返回错误。
	Close() error

	// LocalAddr 返回本地网络地址(如果已知)。
	LocalAddr() Addr

	// RemoteAddr 返回远程网络地址(如果已知)。
	RemoteAddr() Addr

	// SetDeadline 设置与连接关联的读取和写入截止时间。
    // 它等效于调用 SetReadDeadline 和 SetWriteDeadline 两个方法。
	//
	// 截止时间是一个绝对时间，在此之后 I/O 操作将失败而不是阻塞。
	// 截止时间适用于所有未来和挂起的 I/O，而不仅仅是立即跟随 Read 或 Write 调用的操作。
	// 如果超过截止时间，可以通过将截止时间设置为未来来刷新连接。
	//
	// 如果超过截止时间，
    // 对 Read 或 Write 或其他 I/O 方法的调用将返回包含 os.ErrDeadlineExceeded 的错误。
	// 可以使用 errors.Is(err, os.ErrDeadlineExceeded) 测试此错误。
	// 错误的 Timeout 方法将返回 true，但请注意，
    // 即使截止时间尚未超过，也有其他可能导致 Timeout 方法返回 true 的错误。
	//
	// 空值 t 表示 I/O 操作将不会超时。
	SetDeadline(t time.Time) error

	// SetReadDeadline 设置未来 Read 调用和任何当前被阻塞的 Read 调用的截止时间。
	// 空值 t 表示 Read 不会超时。
	SetReadDeadline(t time.Time) error

	// SetWriteDeadline 设置未来 Write 调用和任何当前被阻塞的 Write 调用的截止时间。
	// 即使写入超时，它也可能返回 n > 0，表示某些数据已成功写入。
	// 空值 t 表示 Write 不会超时。
	SetWriteDeadline(t time.Time) error
}
```

​	Conn 是一个通用的面向流的网络连接。

​	多个 goroutine 可以同时调用 Conn 上的方法。

#### func [Dial](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dial.go;l=332) 

``` go linenums="1"
func Dial(network, address string) (Conn, error)
```

​	Dial函数连接到指定网络上的地址。

​	已知的网络有 "tcp"、"tcp4"(仅限 IPv4)、"tcp6"(仅限 IPv6)、"udp"、"udp4"(仅限 IPv4)、"udp6"(仅限 IPv6)、"ip"、"ip4"(仅限 IPv4)、"ip6"(仅限 IPv6)、"unix"、"unixgram" 和 "unixpacket"。

​	对于 TCP 和 UDP 网络，地址的格式为 "host:port"。主机必须是一个字面 IP 地址，或者是可以解析为 IP 地址的主机名。端口必须是字面端口号或服务名称。如果主机是字面 IPv6 地址，则必须用方括号括起来，如 "[2001:db8::1]:80" 或 "[fe80::1%zone]:80"。区域指定了字面 IPv6 地址的作用域，如 RFC 4007 中定义的那样。函数 JoinHostPort 和 SplitHostPort 以这种形式操作一对主机和端口。在使用 TCP 时，如果主机解析为多个 IP 地址，则 Dial 将按顺序尝试每个 IP 地址，直到其中一个成功。

​	示例：

```
Dial("tcp", "golang.org:http")
Dial("tcp", "192.0.2.1:http")
Dial("tcp", "198.51.100.1:80")
Dial("udp", "[2001:db8::1]:domain")
Dial("udp", "[fe80::1%lo0]:53")
Dial("tcp", ":80")
```

​	对于 IP 网络，网络必须是 "ip"、"ip4" 或 "ip6"，后跟冒号和字面协议号或协议名称，地址的格式为 "host"。主机必须是一个字面 IP 地址或带区域的字面 IPv6 地址。每个操作系统如何处理非知名协议号(如 "0" 或 "255")取决于操作系统。

​	示例：

```
Dial("ip4:1", "192.0.2.1")
Dial("ip6:ipv6-icmp", "2001:db8::1")
Dial("ip6:58", "fe80::1%lo0")
```

​	对于 TCP、UDP 和 IP 网络，如果主机为空或为未指定的字面 IP 地址，如 ":80"、"0.0.0.0:80" 或 "[::]:80" 用于 TCP 和 UDP、""、"0.0.0.0" 或 "::" 用于 IP，则假定本地系统。

​	对于 Unix 网络，地址必须是一个文件系统路径。

#### func [DialTimeout](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dial.go;l=347) 

``` go linenums="1"
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
```

​	DialTimeout函数类似于Dial，但带有超时设置。

​	超时包括名称解析(如果需要)。当使用 TCP 时，如果地址参数中的主机名解析为多个 IP 地址，则超时时间将分布在每个连续的 dial 上，以使每个 dial 分配适当的时间比例进行连接。

​	有关网络和地址参数的说明，请参见 func Dial。

#### func [FileConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/file.go;l=21) 

``` go linenums="1"
func FileConn(f *os.File) (c Conn, err error)
```

​	FileConn函数返回与打开文件 f 对应的网络连接的副本。当结束时，调用者负责关闭 f。关闭 c 不会影响 f，关闭 f 也不会影响 c。

### type [DNSConfigError](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=610) 

``` go linenums="1"
type DNSConfigError struct {
	Err error
}
```

​	DNSConfigError 表示读取计算机 DNS 配置的错误。(不再使用；为了向后兼容而保留。)

#### (*DNSConfigError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=615) 

``` go linenums="1"
func (e *DNSConfigError) Error() string
```

#### (*DNSConfigError) [Temporary](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=617) 

``` go linenums="1"
func (e *DNSConfigError) Temporary() bool
```

#### (*DNSConfigError) [Timeout](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=616) 

``` go linenums="1"
func (e *DNSConfigError) Timeout() bool
```

#### (*DNSConfigError) [Unwrap](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=614)  <- go1.13

``` go linenums="1"
func (e *DNSConfigError) Unwrap() error
```

### type [DNSError](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=625) 

``` go linenums="1"
type DNSError struct {
	Err         string // 错误描述
	Name        string // 查找的名称
	Server      string // 使用的服务器
	IsTimeout   bool   // 如果为真，则已超时；并非所有超时都设置了此项
	IsTemporary bool   // 如果为真，则错误是暂时的；并非所有错误都设置了此项
	IsNotFound  bool   // 如果为真，则找不到主机
}
```

​	DNSError 表示 DNS 查找错误。

#### (*DNSError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=634) 

``` go linenums="1"
func (e *DNSError) Error() string
```

#### (*DNSError) [Temporary](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=654) 

``` go linenums="1"
func (e *DNSError) Temporary() bool
```

​	Temporary方法报告DNS错误是否已知为临时错误。这并非总是已知的；DNS查找可能由于临时错误而失败，并返回一个DNSError，其中Temporary返回false。

#### (*DNSError) [Timeout](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=649) 

``` go linenums="1"
func (e *DNSError) Timeout() bool
```

​	Timeout方法报告DNS查找是否已知已超时。这并非总是已知的；DNS查找可能由于超时而失败，并返回一个DNSError，其中Timeout返回false。

### type [Dialer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dial.go;l=27)  <- go1.1

``` go linenums="1"
type Dialer struct {
	// Timeout是拨号等待连接完成的最大时间。
    // 如果Deadline也设置了，可能会更早失败。
	//
	// 默认情况下没有超时。
	//
	// 在使用TCP并拨打具有多个IP地址的主机名时，
    // 超时可能会在它们之间分配。
	//
	// 有或没有超时，操作系统可能会强制执行自己的较早超时。
    // 例如，TCP超时通常为约3分钟。
	Timeout time.Duration

	// Deadline是绝对时间点，超过这个时间点，拨号将失败。
    // 如果设置了Timeout，它可能会更早失败。
	// 零表示没有截止日期，或者依赖于操作系统，如Timeout选项。
	Deadline time.Time

	// LocalAddr是在拨打地址时使用的本地地址。
    // 地址必须是适用于被拨打的网络的兼容类型。
	// 如果为nil，则自动选择本地地址。
	LocalAddr Addr

	// DualStack先前启用了RFC 6555快速回退支持，
    // 也称为" Happy Eyeballs"，
    // 在其中如果IPv6似乎配置不正确且挂起，则很快尝试IPv4。
	//
	// 已弃用：默认情况下启用了快速回退。
    // 要禁用，请将FallbackDelay设置为负值。
	DualStack bool

	// FallbackDelay指定等待RFC 6555快速回退连接生成的时间长度。
    // 也就是说，在假定IPv6配置不正确并回退到IPv4之前，
    // 等待IPv6成功的时间量。
	//
	// 如果为零，则使用默认延迟300毫秒。
	// 负值禁用快速回退支持。
	FallbackDelay time.Duration

	// KeepAlive指定活动网络连接的保持活动探测之间的间隔时间。
	// 如果为零，并且协议和操作系统支持，
    // 则使用默认值发送保持活动探测(当前为15秒)。
    // 不支持保持活动的网络协议或操作系统会忽略此字段。
	// 如果为负值，则禁用保持活动探测。
	KeepAlive time.Duration

	// Resolver是可选的，它指定要使用的替代解析程序。
	Resolver *Resolver

	// Cancel是一个可选的通道，其关闭指示应取消拨号。
    // 不是所有类型的拨号都支持取消。
	//
	// 已弃用：改用DialContext。
	Cancel <-chan struct{}

	// 如果Control不为nil，则在创建网络连接但尚未拨打时调用它。
	//
	// 传递给Control方法的网络和地址参数不一定是传递给Dial的参数。
    // 例如，
    // 将"tcp"传递给Dial将导致使用"tcp4"或"tcp6"调用Control函数。
	//
	// 如果ControlContext不为nil，则忽略Control。
	Control func(network, address string, c syscall.RawConn) error

	// 如果ControlContext不为nil，
    // 在实际拨号之前创建网络连接时会调用它。
	//
	// 传递给Control方法的网络和地址参数不一定是传递给Dial的参数。
    // 例如，
    // 向Dial传递"tcp"将导致Control函数被调用时传递"tcp4" 或 "tcp6"。
	//
	// 如果ControlContext不为nil，则忽略Control。
	ControlContext func(ctx context.Context, network, address string, c syscall.RawConn) error
}
```

​	Dialer包含用于连接到地址的选项。

​	每个字段的零值等效于不使用该选项进行拨号。因此，使用Dialer的零值进行拨号等效于只调用Dial函数。

​	并发调用Dialer的方法是安全的。

##### Example
``` go linenums="1"
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

##### Example(Unix)
``` go linenums="1"
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

#### (*Dialer) [Dial](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dial.go;l=366)  <- go1.1

``` go linenums="1"
func (d *Dialer) Dial(network, address string) (Conn, error)
```

​	Dial方法在指定网络上连接到地址。

​	有关网络和地址参数的描述，请参见func Dial。

​	Dial在内部使用context.Background；要指定上下文，请使用DialContext。

#### (*Dialer) [DialContext](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dial.go;l=388)  <- go1.7

``` go linenums="1"
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)
```

​	DialContext方法在指定网络上使用提供的上下文连接到地址。

​	提供的上下文必须是非nil的。如果上下文在连接完成之前到期，则会返回错误。一旦成功连接，上下文的任何到期都不会影响连接。

​	使用TCP时，如果地址参数中的主机解析为多个网络地址，则任何拨号超时(来自d.Timeout或ctx)将分布在每个连续的拨号之间，以便每个拨号在适当的时间内完成。例如，如果主机有4个IP地址，并且超时时间为1分钟，则在尝试下一个地址之前，将为每个单个地址分配15秒钟的时间。

​	有关网络和地址参数的描述，请参见func Dial。

### type [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=397) 

``` go linenums="1"
type Error interface {
	error
	Timeout() bool // 错误是否超时？

	// Deprecated: 临时错误没有明确定义。
	// 大多数"临时"错误是超时错误，极少数例外情况是出人意料的。
	// 不要使用这个方法。
	Temporary() bool
}
```

​	Error接口表示网络错误。

### type [Flags](https://cs.opensource.google/go/go/+/go1.20.1:src/net/interface.go;l=39) 

``` go linenums="1"
type Flags uint
const (
	FlagUp           Flags = 1 << iota // 接口已被管理员开启
	FlagBroadcast                      // 接口支持广播访问能力
	FlagLoopback                       // 接口是一个回环接口
	FlagPointToPoint                   // 接口属于点对点连接
	FlagMulticast                      // 接口支持组播访问能力
	FlagRunning                        // 接口处于运行状态
)
```

#### (Flags) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/interface.go;l=59) 

``` go linenums="1"
func (f Flags) String() string
```

### type [HardwareAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mac.go;l=10) 

``` go linenums="1"
type HardwareAddr []byte
```

​	HardwareAddr表示一个物理硬件地址。

#### func [ParseMAC](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mac.go;l=39) 

``` go linenums="1"
func ParseMAC(s string) (hw HardwareAddr, err error)
```

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

#### (HardwareAddr) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mac.go;l=12) 

``` go linenums="1"
func (a HardwareAddr) String() string
```

### type [IP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=35) 

``` go linenums="1"
type IP []byte
```

​	IP是单个IP地址，是字节片。此包中的函数接受4字节(IPv4)或16字节(IPv6)片作为输入。

​	请注意，在本文档中，将IP地址称为IPv4地址或IPv6地址是地址的语义属性，而不仅仅是字节片的长度：16字节的片仍然可以是IPv4地址。

##### IP Example
``` go linenums="1"
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

#### func [IPv4](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=51) 

``` go linenums="1"
func IPv4(a, b, c, d byte) IP
```

​	IPv4函数返回IPv4地址a.b.c.d的IP地址(以16字节形式)。

##### IPv4 Example
``` go linenums="1"
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

#### func [LookupIP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=196) 

``` go linenums="1"
func LookupIP(host string) ([]IP, error)
```

​	LookupIP函数使用本地解析器查找主机，返回该主机的IPv4和IPv6地址切片。

#### func [ParseIP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=707) 

``` go linenums="1"
func ParseIP(s string) IP
```

​	ParseIP函数将 s 解析为一个 IP 地址并返回结果。字符串 s 可以是 IPv4 的点分十进制表示("192.0.2.1")、IPv6 的十六进制表示("`2001:db8::68`")或 IPv4 映射的 IPv6 表示形式("`::ffff:192.0.2.1`")。如果 s 不是有效的文本表示形式的 IP 地址，则 ParseIP 返回 nil。

##### ParseIP Example
``` go linenums="1"
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

#### (IP) [DefaultMask](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=246) 

``` go linenums="1"
func (ip IP) DefaultMask() IPMask
```

​	DefaultMask方法返回 IP 地址 ip 的默认 IP 掩码。只有 IPv4 地址有默认掩码；如果 ip 不是有效的 IPv4 地址，则 DefaultMask 返回 nil。

##### DefaultMask Example
``` go linenums="1"
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

#### (IP) [Equal](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=432) 

``` go linenums="1"
func (ip IP) Equal(x IP) bool
```

​	Equal方法报告 ip 和 x 是否是相同的 IP 地址。IPv4 地址和相同的 IPv6 地址被认为是相等的。

##### Equal Example
``` go linenums="1"
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

#### (IP) [IsGlobalUnicast](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=190) 

``` go linenums="1"
func (ip IP) IsGlobalUnicast() bool
```

​	IsGlobalUnicast方法报告 ip 是否是全局单播地址。

​	全局单播地址的标识使用 [RFC 1122](https://rfc-editor.org/rfc/rfc1122.html)、[RFC 4632](https://rfc-editor.org/rfc/rfc4632.html) 和 [RFC 4291](https://rfc-editor.org/rfc/rfc4291.html) 中定义的地址类型标识，但 IPv4 定向广播地址除外。即使 ip 在 IPv4 私有地址空间或本地 IPv6 单播地址空间中，它也返回 true。

##### IsGlobalUnicast Example
``` go linenums="1"
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

#### (IP) [IsInterfaceLocalMulticast](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=160) 

``` go linenums="1"
func (ip IP) IsInterfaceLocalMulticast() bool
```

​	IsInterfaceLocalMulticast方法报告 ip 是否是接口本地组播地址。

##### IsInterfaceLocalMulticast Example
``` go linenums="1"
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

#### (IP) [IsLinkLocalMulticast](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=166) 

``` go linenums="1"
func (ip IP) IsLinkLocalMulticast() bool
```

​	IsLinkLocalMulticast方法报告 ip 是否是链路本地组播地址。

##### IsLinkLocalMulticast Example
``` go linenums="1"
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

#### (IP) [IsLinkLocalUnicast](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=175) 

``` go linenums="1"
func (ip IP) IsLinkLocalUnicast() bool
```

​	IsLinkLocalUnicast方法报告IP地址ip是否为链路本地单播地址。

##### IsLinkLocalUnicast Example
``` go linenums="1"
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

#### (IP) [IsLoopback](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=124) 

``` go linenums="1"
func (ip IP) IsLoopback() bool
```

​	IsLoopback方法报告IP地址ip是否为回环地址。

##### IsLoopback Example
``` go linenums="1"
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

#### (IP) [IsMulticast](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=151) 

``` go linenums="1"
func (ip IP) IsMulticast() bool
```

​	IsMulticast方法报告IP地址ip是否为多播地址。

##### IsMulticast Example
``` go linenums="1"
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

#### (IP) [IsPrivate](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=133)  <- go1.17

``` go linenums="1"
func (ip IP) IsPrivate() bool
```

​	IsPrivate方法报告IP地址ip是否为私有地址，根据[RFC 1918](https://rfc-editor.org/rfc/rfc1918.html)(IPv4地址)和[RFC 4193](https://rfc-editor.org/rfc/rfc4193.html)(IPv6地址)。

##### IsPrivate Example
``` go linenums="1"
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

#### (IP) [IsUnspecified](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=119) 

``` go linenums="1"
func (ip IP) IsUnspecified() bool
```

​	IsUnspecified方法报告IP地址ip是否为未指定地址，即IPv4地址"0.0.0.0"或IPv6地址"::"。

##### IsUnspecified Example
``` go linenums="1"
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

#### (IP) [MarshalText](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=403)  <- go1.2

``` go linenums="1"
func (ip IP) MarshalText() ([]byte, error)
```

​	MarshalText方法实现encoding.TextMarshaler接口。编码与String返回的编码相同，唯一的区别是当len(ip)为零时，它会返回一个空切片。

#### (IP) [Mask](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=270) 

``` go linenums="1"
func (ip IP) Mask(mask IPMask) IP
```

​	Mask方法返回使用掩码mask对IP地址ip进行掩码后的结果。

##### Mask Example
``` go linenums="1"
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

#### (IP) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=313) 

``` go linenums="1"
func (ip IP) String() string
```

String方法返回IP地址ip的字符串形式。它返回以下4种形式之一：

- "`<nil>`"，如果ip长度为0
- 点分十进制("192.0.2.1")，如果ip是IPv4或IP4映射IPv6地址
- 符合[RFC 5952](https://rfc-editor.org/rfc/rfc5952.html)的IPv6地址("2001:db8::1")，如果ip是有效的IPv6地址
- 如果没有其他情况，则是ip的十六进制形式，不包含标点符号



##### String Example
``` go linenums="1"
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

#### (IP) [To16](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=226) 

``` go linenums="1"
func (ip IP) To16() IP
```

​	To16 方法将IP地址ip转换为16字节表示。如果ip不是IP地址(长度不正确)，To16 返回nil。

##### To16 Example
``` go linenums="1"
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

#### (IP) [To4](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=211) 

``` go linenums="1"
func (ip IP) To4() IP
```

​	To4方法将IPv4地址ip转换为4字节表示。如果ip不是IPv4地址，则To4返回nil。

#### (*IP) [UnmarshalText](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=415)  <- go1.2

``` go linenums="1"
func (ip *IP) UnmarshalText(text []byte) error
```

​	UnmarshalText方法实现encoding.TextUnmarshaler接口。IP地址应该以ParseIP接受的形式出现。

### type [IPAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=31) 

``` go linenums="1"
type IPAddr struct {
	IP   IP
	Zone string // IPv6有作用域的寻址区域
}
```

​	IPAddr结构体表示IP端点的地址。

#### func [ResolveIPAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=77) 

``` go linenums="1"
func ResolveIPAddr(network, address string) (*IPAddr, error)
```

​	ResolveIPAddr函数返回IP端点的地址。

​	网络必须是一个IP网络名称。

​	如果address参数中的主机不是字面上的IP地址，则ResolveIPAddr将地址解析为IP端点地址。否则，它将解析地址为一个字面上的IP地址。地址参数可以使用主机名，但不建议这样做，因为它最多只返回主机名的一个IP地址。

​	有关网络和地址参数的描述，请参见func Dial。

#### (*IPAddr) [Network](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=37) 

``` go linenums="1"
func (a *IPAddr) Network() string
```

​	Network方法返回地址的网络名称，即"ip"。

#### (*IPAddr) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=39) 

``` go linenums="1"
func (a *IPAddr) String() string
```

### type [IPConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=99) 

``` go linenums="1"
type IPConn struct {
	// 包含过滤或未导出字段
}
```

​	IPConn结构体是IP网络连接的Conn和PacketConn接口的实现。

#### func [DialIP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=211) 

``` go linenums="1"
func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error)
```

​	DialIP函数类似于IP网络的Dial函数。

​	网络必须是一个IP网络名称；有关详细信息，请参见func Dial。

​	如果laddr为nil，则会自动选择本地地址。如果raddr的IP字段为nil或未指定的IP地址，则假定为本地系统。

#### func [ListenIP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=230) 

``` go linenums="1"
func ListenIP(network string, laddr *IPAddr) (*IPConn, error)
```

​	ListenIP函数类似于IP网络的ListenPacket函数。

​	网络必须是一个IP网络名称；有关详细信息，请参见func Dial。

​	如果laddr的IP字段为nil或未指定的IP地址，则ListenIP会侦听本地系统的所有可用IP地址，但不包括多播IP地址。

#### (*IPConn) [Close](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=203) 

``` go linenums="1"
func (c *IPConn) Close() error
```

​	Close方法关闭连接。

#### (*IPConn) [File](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=298) 

``` go linenums="1"
func (c *IPConn) File() (f *os.File, err error)
```

​	File方法返回基础os.File的副本。调用者有责任在完成后关闭f。关闭c不会影响f，关闭f也不会影响c。

​	返回的os.File的文件描述符与连接的不同。尝试使用此副本更改原始文件的属性可能会或可能不会产生预期的效果。

#### (*IPConn) [LocalAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=217) 

``` go linenums="1"
func (c *IPConn) LocalAddr() Addr
```

​	LocalAddr方法返回本地网络地址。返回的Addr由LocalAddr的所有调用共享，因此不要修改它。

#### (*IPConn) [Read](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=179) 

``` go linenums="1"
func (c *IPConn) Read(b []byte) (int, error)
```

​	Read方法实现Conn Read方法。

#### (*IPConn) [ReadFrom](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=125) 

``` go linenums="1"
func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)
```

​	ReadFrom方法实现PacketConn ReadFrom 方法。

#### (*IPConn) [ReadFromIP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=113) 

``` go linenums="1"
func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)
```

​	ReadFromIP方法类似于ReadFrom但返回IPAddr。

#### (*IPConn) [ReadMsgIP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=146)  <- go1.1

``` go linenums="1"
func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)
```

​	ReadMsgIP方法从c读取消息，将有效负载复制到b中，将相关的带外数据复制到oob中。它返回复制到b中的字节数，复制到oob中的字节数，消息设置的标志以及消息的源地址。

​	golang.org/x/net/ipv4 和 golang.org/x/net/ipv6 包可用于操作oob中的IP级套接字选项。

#### (*IPConn) [RemoteAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=227) 

``` go linenums="1"
func (c *IPConn) RemoteAddr() Addr
```

​	RemoteAddr方法返回远程网络地址。返回的Addr由RemoteAddr的所有调用共享，因此不要修改它。

#### (*IPConn) [SetDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=235) 

``` go linenums="1"
func (c *IPConn) SetDeadline(t time.Time) error
```

​	SetDeadline方法实现Conn SetDeadline方法。

#### (*IPConn) [SetReadBuffer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=269) 

``` go linenums="1"
func (c *IPConn) SetReadBuffer(bytes int) error
```

​	SetReadBuffer方法设置与连接相关联的操作系统接收缓冲区的大小。

#### (*IPConn) [SetReadDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=246) 

``` go linenums="1"
func (c *IPConn) SetReadDeadline(t time.Time) error
```

​	SetReadDeadline方法实现Conn SetReadDeadline方法。

#### (*IPConn) [SetWriteBuffer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=281) 

``` go linenums="1"
func (c *IPConn) SetWriteBuffer(bytes int) error
```

​	SetWriteBuffer方法设置与连接相关联的操作系统传输缓冲区的大小。

#### (*IPConn) [SetWriteDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=257) 

``` go linenums="1"
func (c *IPConn) SetWriteDeadline(t time.Time) error
```

​	SetWriteDeadline方法实现Conn SetWriteDeadline方法。

#### (*IPConn) [SyscallConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=105)  <- go1.9

``` go linenums="1"
func (c *IPConn) SyscallConn() (syscall.RawConn, error)
```

​	SyscallConn方法返回原始网络连接。这实现了syscall.Conn接口。

#### (*IPConn) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=191) 

``` go linenums="1"
func (c *IPConn) Write(b []byte) (int, error)
```

​	Write方法实现Conn Write方法。

#### (*IPConn) [WriteMsgIP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=191)  <- go1.1

``` go linenums="1"
func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)
```

​	WriteMsgIP方法通过c向addr写入消息，将有效负载从b复制，将相关的带外数据从oob复制。它返回写入的有效负载和带外字节数。

​	golang.org/x/net/ipv4 和 golang.org/x/net/ipv6 包可用于操作oob中的IP级套接字选项。

#### (*IPConn) [WriteTo](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=170) 

``` go linenums="1"
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)
```

​	WriteTo方法实现PacketConn WriteTo方法。

#### (*IPConn) [WriteToIP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/iprawsock.go;l=158) 

``` go linenums="1"
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)
```

​	WriteToIP方法类似于WriteTo但接受IPAddr。

### type [IPMask](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=41) 

``` go linenums="1"
type IPMask []byte
```

​	IPMask是一个比特掩码，可用于操作IP地址，用于IP寻址和路由。

​	有关详细信息，请参见类型IPNet和函数ParseCIDR。

#### func [CIDRMask](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=77) 

``` go linenums="1"
func CIDRMask(ones, bits int) IPMask
```

​	CIDRMask函数返回一个IPMask，它由"ones"个1位组成，后跟0位，总长度为"bits"位。对于这种形式的掩码，CIDRMask是IPMask.Size的反函数。

##### CIDRMask Example
``` go linenums="1"
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

#### func [IPv4Mask](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=65) 

``` go linenums="1"
func IPv4Mask(a, b, c, d byte) IPMask
```

​	IPv4Mask函数返回IPv4掩码a.b.c.d的IP掩码(以4字节形式)。

##### IPv4Mask Example
``` go linenums="1"
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

#### (IPMask) [Size](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=481) 

``` go linenums="1"
func (m IPMask) Size() (ones, bits int)
```

​	Size方法返回掩码中前导1位和总位数。如果掩码不在规范形式——1位后面是0位——那么Size将返回0, 0。

#### (IPMask) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=490) 

``` go linenums="1"
func (m IPMask) String() string
```

​	String方法以无标点符号的十六进制形式返回m。

### type [IPNet](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=44) 

``` go linenums="1"
type IPNet struct {
	IP   IP     // 网络号
	Mask IPMask // 网络掩码
}
```

​	IPNet表示IP网络。

#### (*IPNet) [Contains](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=521) 

``` go linenums="1"
func (n *IPNet) Contains(ip IP) bool
```

​	Contains方法报告网络是否包含ip。

#### (*IPNet) [Network](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=539) 

``` go linenums="1"
func (n *IPNet) Network() string
```

​	Network方法返回地址的网络名称，"ip+net"。

#### (*IPNet) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/ip.go;l=547) 

``` go linenums="1"
func (n *IPNet) String() string
```

​	String方法返回n的CIDR表示形式，例如"192.0.2.0/24"或"2001:db8::/48"，如[RFC 4632](https://rfc-editor.org/rfc/rfc4632.html)和[RFC 4291](https://rfc-editor.org/rfc/rfc4291.html)所定义。如果掩码不在规范形式中，则返回的字符串由IP地址和斜杠字符和表示为十六进制形式且无标点符号的掩码组成，例如"198.51.100.0/c000ff00"。

### type [Interface](https://cs.opensource.google/go/go/+/go1.20.1:src/net/interface.go;l=31) 

``` go linenums="1"
type Interface struct {
	Index        int          // 正整数，从1开始，0永远不使用
	MTU          int          // 最大传输单元
	Name         string       // 例如"en0"，"lo0"，"eth0.100"
	HardwareAddr HardwareAddr // IEEE MAC-48、EUI-48和EUI-64形式
	Flags        Flags        // 例如FlagUp、FlagLoopback、FlagMulticast
}
```

​	Interface结构体表示网络接口名称和索引之间的映射。它还表示网络接口设施信息。

#### func [InterfaceByIndex](https://cs.opensource.google/go/go/+/go1.20.1:src/net/interface.go;l=131) 

``` go linenums="1"
func InterfaceByIndex(index int) (*Interface, error)
```

​	InterfaceByIndex函数返回指定索引的接口。

​	在Solaris上，它返回共享逻辑数据链路的逻辑网络接口之一；要获得更精确的结果，请使用InterfaceByName。

#### func [InterfaceByName](https://cs.opensource.google/go/go/+/go1.20.1:src/net/interface.go;l=156) 

``` go linenums="1"
func InterfaceByName(name string) (*Interface, error)
```

​	InterfaceByName函数返回指定名称的接口。

#### func [Interfaces](https://cs.opensource.google/go/go/+/go1.20.1:src/net/interface.go;l=102) 

``` go linenums="1"
func Interfaces() ([]Interface, error)
```

​	Interfaces函数返回系统的网络接口列表。

#### (*Interface) [Addrs](https://cs.opensource.google/go/go/+/go1.20.1:src/net/interface.go;l=77) 

``` go linenums="1"
func (ifi *Interface) Addrs() ([]Addr, error)
```

​	Addrs方法返回特定接口的单播接口地址列表。

#### (*Interface) [MulticastAddrs](https://cs.opensource.google/go/go/+/go1.20.1:src/net/interface.go;l=90) 

``` go linenums="1"
func (ifi *Interface) MulticastAddrs() ([]Addr, error)
```

​	MulticastAddrs方法返回特定接口的多播、组播地址列表。

### type [InvalidAddrError](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=579) 

``` go linenums="1"
type InvalidAddrError string
```

#### (InvalidAddrError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=581) 

``` go linenums="1"
func (e InvalidAddrError) Error() string
```

#### (InvalidAddrError) [Temporary](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=583) 

``` go linenums="1"
func (e InvalidAddrError) Temporary() bool
```

#### (InvalidAddrError) [Timeout](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=582) 

``` go linenums="1"
func (e InvalidAddrError) Timeout() bool
```

### type [ListenConfig](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dial.go;l=600)  <- go1.11

``` go linenums="1"
type ListenConfig struct {
	// 如果 Control 不是 nil，
    // 则在创建网络连接但在将其绑定到操作系统之前调用它。
	//
	// 传递给 Control 方法的网络和地址参数未必是传递给 Listen 的参数。
	// 例如，将 "tcp" 传递给 Listen 
    // 将导致 Control 函数使用 "tcp4" 或 "tcp6" 调用。
	Control func(network, address string, c syscall.RawConn) error

	// KeepAlive 指定此侦听器接受的网络连接的保活期。
	// 如果为零，则如果协议和操作系统支持保活，则启用保活。
    // 不支持保活的网络协议或操作系统将忽略此字段。
	// 如果为负，则禁用保活。
	KeepAlive time.Duration
}
```

​	ListenConfig结构体包含用于监听地址的选项。

#### (*ListenConfig) [Listen](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dial.go;l=622)  <- go1.11

``` go linenums="1"
func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)
```

​	Listen方法在本地网络地址上进行侦听。

​	有关网络和地址参数的描述，请参见 func Listen。

#### (*ListenConfig) [ListenPacket](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dial.go;l=652)  <- go1.11

``` go linenums="1"
func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error)
```

​	ListenPacket方法在本地网络地址上进行侦听。

​	有关网络和地址参数的描述，请参见 func ListenPacket。

### type [Listener](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=384) 

``` go linenums="1"
type Listener interface {
	// Accept 等待并返回到监听器的下一个连接。
	Accept() (Conn, error)

	// Close 关闭监听器。
	// 任何已阻止的 Accept 操作将解除阻塞并返回错误。
	Close() error

	// Addr 返回监听器的网络地址。
	Addr() Addr
}
```

​	Listener接口是面向流协议的通用网络监听器。

​	多个 goroutine 可同时调用 Listener 上的方法。

##### Listener Example
``` go linenums="1"
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

#### func [FileListener](https://cs.opensource.google/go/go/+/go1.20.1:src/net/file.go;l=33) 

``` go linenums="1"
func FileListener(f *os.File) (ln Listener, err error)
```

​	FileListener函数返回与已打开文件 f 相应的网络监听器的副本。调用者有责任在完成时关闭 ln。关闭 ln 不会影响 f，关闭 f 不会影响 ln。

#### func [Listen](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dial.go;l=707) 

``` go linenums="1"
func Listen(network, address string) (Listener, error)
```

​	Listen函数在本地网络地址上进行监听。

​	网络必须是 "tcp"、"tcp4"、"tcp6"、"unix" 或 "unixpacket"。

​	对于 TCP 网络，如果地址参数中的主机为空或为字面未指定 IP 地址，则 Listen 会侦听本地系统的所有可用单播和任播 IP 地址。要仅使用 IPv4，请使用网络 "tcp4"。地址可以使用主机名，但不建议这样做，因为它将仅为主机的一个 IP 地址创建侦听器。如果地址参数中的端口为空或为 "0"，例如 "127.0.0.1:" 或 "[::1]:0"，则会自动选择一个端口号。可以使用 Listener 的 Addr 方法来查找所选端口。

​	有关网络和地址参数的描述，请参见 func Dial。

​	Listen 在内部使用 context.Background；要指定上下文，请使用 ListenConfig.Listen。

### type [MX](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dnsclient.go;l=204) 

``` go linenums="1"
type MX struct {
	Host string
	Pref uint16
}
```

​	MX结构体表示单个 DNS MX 记录。

#### func [LookupMX](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=549) 

``` go linenums="1"
func LookupMX(name string) ([]*MX, error)
```

​	LookupMX函数按优先级排序返回给定域名的 DNS MX 记录。

​	返回的邮件服务器名称已验证为格式正确的演示格式域名。如果响应包含无效名称，则这些记录将被过滤掉，并将在剩余结果(如果有)旁边返回错误。

​	LookupMX 在内部使用 context.Background；要指定上下文，请使用 Resolver.LookupMX。

### type [NS](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dnsclient.go;l=226)  <- go1.1

``` go linenums="1"
type NS struct {
	Host string
}
```

​	NS结构体表示单个 DNS NS 记录。

#### func [LookupNS](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=589)  <- go1.1

``` go linenums="1"
func LookupNS(name string) ([]*NS, error)
```

​	LookupNS 函数返回给定域名的 DNS NS 记录。

​	返回的名称服务器名称已验证为格式正确的演示格式域名。如果响应包含无效名称，则这些记录将被过滤掉，并将在剩余结果(如果有)旁边返回错误。

​	LookupNS 在内部使用 context.Background；要指定上下文，请使用 Resolver.LookupNS。

### type [OpError](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=444) 

``` go linenums="1"
type OpError struct {
	// Op(操作)是引起错误的操作，例如"read"或"write"。
	Op string

	// Net(网络)是发生此错误的网络类型，例如"tcp"或"udp6"。
	Net string

	// 对于涉及远程网络连接的操作，
    // 如Dial、Read或Write，Source是相应的本地网络地址。
	Source Addr

	// Addr(地址)是发生此错误的网络地址。
	// 对于本地操作，如Listen或SetDeadline，
    // Addr是正在操作的本地端点的地址。
	// 对于涉及远程网络连接的操作，
    // 如Dial、Read或Write，Addr是该连接的远程地址。
	Addr Addr

	// Err(错误)是操作期间发生的错误。
	// 如果错误为nil，则Error方法会引发panic。
	Err error
}
```

​	OpError结构体是通常由net包中的函数返回的错误类型。它描述了错误的操作、网络类型和地址。

#### (*OpError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=473) 

``` go linenums="1"
func (e *OpError) Error() string
```

#### (*OpError) [Temporary](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=524) 

``` go linenums="1"
func (e *OpError) Temporary() bool
```

#### (*OpError) [Timeout](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=511) 

``` go linenums="1"
func (e *OpError) Timeout() bool
```

#### (*OpError) [Unwrap](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=471)  <- go1.13

``` go linenums="1"
func (e *OpError) Unwrap() error
```

### type [PacketConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=309) 

``` go linenums="1"
type PacketConn interface {
	// ReadFrom 从连接中读取一个数据包，将有效载荷复制到 p 中。
	// 它返回复制到 p 中的字节数以及在数据包上的返回地址。
	// 它返回读取的字节数(0 <= n <= len(p))和任何遇到的错误。
	// 调用者应该始终在考虑错误 err 之前处理返回的 n > 0 字节。
	// ReadFrom 可以在固定的时间限制后超时并返回一个错误；
    // 参见 SetDeadline 和 SetReadDeadline。
	ReadFrom(p []byte) (n int, addr Addr, err error)

	// WriteTo 将有效载荷为 p 的数据包写入 addr。
	// WriteTo 可以在固定的时间限制后超时并返回一个错误；
    // 参见 SetDeadline 和 SetWriteDeadline。
	// 在面向数据包的连接中，写入超时很少发生。
	WriteTo(p []byte, addr Addr) (n int, err error)

	// Close 关闭连接。
	// 任何被阻塞的 ReadFrom 或 WriteTo 操作都将被取消阻塞并返回错误。
	Close() error

	// LocalAddr 返回本地网络地址(如果已知)。
	LocalAddr() Addr

	// SetDeadline 设置与连接关联的读取和写入截止时间。
	// 它相当于调用 SetReadDeadline 和 SetWriteDeadline。
	//
	// 截止时间是绝对时间，超过这个时间，I/O 操作将失败而不是阻塞。
	// 截止时间适用于所有未来和挂起的 I/O，
    // 而不仅仅是即将到来的 Read 或 Write 调用。
	// 在超过截止时间后，可以通过将未来的截止时间设置为刷新连接。
	//
	// 如果超过截止时间，调用 Read 或 Write 或其他 I/O 方法
    // 将返回一个错误，该错误包装了 os.ErrDeadlineExceeded。
	// 可以使用 errors.Is(err, os.ErrDeadlineExceeded) 进行测试。
	// 错误的 Timeout 方法将返回 true，
    // 但请注意，即使未超过截止时间，也可能有其他可能的错误，
    // Timeout 方法将返回 true。
	//
	// 空值 t 表示 I/O 操作不会超时。
	SetDeadline(t time.Time) error

	// SetReadDeadline 设置未来的 ReadFrom 调用的截止时间和
    // 任何当前被阻塞的 ReadFrom 调用。
	// t 的零值表示 ReadFrom 不会超时。
	SetReadDeadline(t time.Time) error

	// SetWriteDeadline 设置未来的 WriteTo 调用的截止时间和
    // 任何当前被阻塞的 WriteTo 调用。
	// 即使写入超时，它也可能返回 n > 0，表示某些数据已成功写入。
	// t 的零值表示 WriteTo 不会超时。
	SetWriteDeadline(t time.Time) error
}
```

​	PacketConn接口是一个通用的面向数据包的网络连接。

​	多个 goroutine 可以同时调用 PacketConn 上的方法。

#### func [FilePacketConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/file.go;l=45) 

``` go linenums="1"
func FilePacketConn(f *os.File) (c PacketConn, err error)
```

​	FilePacketConn函数返回与打开的文件 f 对应的数据包网络连接的副本。当使用完毕时，调用方负责关闭 f。关闭 c 不影响 f，关闭 f 也不影响 c。

#### func [ListenPacket](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dial.go;l=737) 

``` go linenums="1"
func ListenPacket(network, address string) (PacketConn, error)
```

​	ListenPacket函数在本地网络地址上进行公告。

​	网络必须是 "udp"、"udp4"、"udp6"、"unixgram" 或 IP 传输。IP 传输为 "ip"、"ip4" 或 "ip6"，后跟冒号和字面协议号或协议名称，例如 "ip:1" 或 "ip:icmp"。

​	对于 UDP 和 IP 网络，如果 address 参数中的主机为空或为字面未指定的 IP 地址，则 ListenPacket 会在本地系统的所有可用 IP 地址上监听，但不包括多播 IP 地址。要仅使用 IPv4，请使用网络 "udp4" 或 "ip4:proto"。地址可以使用主机名，但不建议这样做，因为这将仅为主机的一个 IP 地址创建侦听器。如果地址参数中的端口为空或为 "0"，例如 "127.0.0.1:" 或 "[::1]:0"，则会自动选择一个端口号。可以使用 PacketConn 的 LocalAddr 方法来发现所选端口。

​	有关网络和地址参数的描述，请参见 func Dial。

​	ListenPacket 在内部使用 context.Background；要指定上下文，请使用 ListenConfig.ListenPacket。

### type [ParseError](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=540) 

``` go linenums="1"
type ParseError struct {
	// Type 是预期的字符串类型，例如
	// "IP 地址"、"CIDR 地址"。
	Type string

	// Text 是格式不正确的文本字符串。
	Text string
}
```

​	ParseError结构体是文字网络地址解析器的错误类型。

#### (*ParseError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=549) 

``` go linenums="1"
func (e *ParseError) Error() string
```

#### (*ParseError) [Temporary](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=552)  <- go1.17

``` go linenums="1"
func (e *ParseError) Temporary() bool
```

#### (*ParseError) [Timeout](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=551)  <- go1.17

``` go linenums="1"
func (e *ParseError) Timeout() bool
```

### type [Resolver](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=123)  <- go1.8

``` go linenums="1"
type Resolver struct {
	// PreferGo 控制在可用平台上是否优先使用 Go 的内置 DNS 解析器。
	// 它等效于设置 GODEBUG=netdns=go，但仅限于此解析器。
	PreferGo bool

	// StrictErrors 控制使用 Go 的内置解析器时临时错误的行为
	//(包括超时、套接字错误和 SERVFAIL)。
    // 对于由多个子查询(例如 A+AAAA 地址查找或遍历 DNS 搜索列表)
    // 组成的查询，此选项会导致这些错误中止整个查询，
    // 而不是返回部分结果。默认情况下未启用此选项，
    // 因为它可能会影响与处理 AAAA 查询不正确的解析器的兼容性。
	StrictErrors bool

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

​	Resolver 查找名称和数字。

​	nil 的 *Resolver 等效于零 Resolver。

#### (*Resolver) [LookupAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=655)  <- go1.8

``` go linenums="1"
func (r *Resolver) LookupAddr(ctx context.Context, addr string) ([]string, error)
```

​	LookupAddr方法执行给定地址的反向查找，返回映射到该地址的名称列表。

​	返回的名称经过验证，以确保格式正确的呈现格式域名。如果响应包含无效名称，则这些记录将被过滤掉，并在剩余结果(如果有)的同时返回错误。

#### (*Resolver) [LookupCNAME](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=473)  <- go1.8

``` go linenums="1"
func (r *Resolver) LookupCNAME(ctx context.Context, host string) (string, error)
```

​	LookupCNAME方法返回给定主机的规范名称。不关心规范名称的调用方可以直接调用 LookupHost 或 LookupIP；它们都会在查找过程中解析规范名称。

​	规范名称是在遵循零个或多个 CNAME 记录后的最终名称。如果主机不包含 DNS "CNAME" 记录，LookupCNAME 不会返回错误，只要主机解析为地址记录即可。

​	返回的规范名称经过验证，以确保格式正确的呈现格式域名。

#### (*Resolver) [LookupHost](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=182)  <- go1.8

``` go linenums="1"
func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error)
```

​	LookupHost方法使用本地解析器查找给定的主机。它返回该主机的地址片段。

#### (*Resolver) [LookupIP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=218)  <- go1.15

``` go linenums="1"
func (r *Resolver) LookupIP(ctx context.Context, network, host string) ([]IP, error)
```

​	LookupIP方法使用本地解析器查找给定网络的主机。它返回该主机指定类型的 IP 地址的片段。network 必须是 "ip"、"ip4" 或 "ip6" 之一。

#### (*Resolver) [LookupIPAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=210)  <- go1.8

``` go linenums="1"
func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error)
```

​	LookupIPAddr方法使用本地解析器查找主机。它返回该主机的 IPv4 和 IPv6 地址的片段。

#### (*Resolver) [LookupMX](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=559)  <- go1.8

``` go linenums="1"
func (r *Resolver) LookupMX(ctx context.Context, name string) ([]*MX, error)
```

​	LookupMX方法返回给定域名的 DNS MX 记录，并按优先级排序。

​	返回的邮件服务器名称经过验证，以确保格式正确的呈现格式域名。如果响应包含无效名称，则这些记录将被过滤掉，并在剩余结果(如果有)的同时返回错误。

#### (*Resolver) [LookupNS](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=599)  <- go1.8

``` go linenums="1"
func (r *Resolver) LookupNS(ctx context.Context, name string) ([]*NS, error)
```

​	LookupNS方法返回给定域名的 DNS NS 记录。

​	返回的名称服务器名称已验证为格式正确的演示格式域名。如果响应包含无效名称，则这些记录将被过滤，并将返回一个错误以及剩余结果(如果有)。

#### (*Resolver) [LookupNetIP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=248)  <- go1.18

``` go linenums="1"
func (r *Resolver) LookupNetIP(ctx context.Context, network, host string) ([]netip.Addr, error)
```

​	LookupNetIP方法使用本地解析器查找主机。它返回指定网络类型的主机 IP 地址切片。网络类型必须是"ip"、"ip4"或"ip6"。

#### (*Resolver) [LookupPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=419)  <- go1.8

``` go linenums="1"
func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error)
```

​	LookupPort方法查找给定网络和服务的端口。

#### (*Resolver) [LookupSRV](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=516)  <- go1.8

``` go linenums="1"
func (r *Resolver) LookupSRV(ctx context.Context, service, proto, name string) (string, []*SRV, error)
```

​	LookupSRV方法尝试解析给定服务、协议和域名的 SRV 查询。proto 为"tcp"或"udp"。返回的记录按优先级排序，并在优先级内按权重随机排序。

​	LookupSRV 根据 [RFC 2782](https://rfc-editor.org/rfc/rfc2782.html)构造要查找的 DNS 名称。也就是说，它查找 _service._proto.name。为了适应在非标准名称下发布 SRV 记录的服务，如果服务和 proto 都为空字符串，则 LookupSRV 直接查找 name。

​	返回的服务名称已验证为格式正确的演示格式域名。如果响应包含无效名称，则这些记录将被过滤，并将返回一个错误以及剩余结果(如果有)。



#### (*Resolver) [LookupTXT](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=629)  <- go1.8

``` go linenums="1"
func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error)
```

​	LookupTXT方法返回给定域名的 DNS TXT 记录。

### type [SRV](https://cs.opensource.google/go/go/+/go1.20.1:src/net/dnsclient.go;l=150) 

``` go linenums="1"
type SRV struct {
	Target   string
	Port     uint16
	Priority uint16
	Weight   uint16
}
```

​	SRV 表示单个 DNS SRV 记录。

#### func [LookupSRV](https://cs.opensource.google/go/go/+/go1.20.1:src/net/lookup.go;l=498) 

``` go linenums="1"
func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)
```

​	LookupSRV函数尝试解析给定服务、协议和域名的 SRV 查询。proto 是 "tcp" 或 "udp"。返回的记录按优先级排序，并在优先级内按权重随机分布。

​	LookupSRV 构造了遵循 [RFC 2782](https://rfc-editor.org/rfc/rfc2782.html)的 DNS 名称进行查找。也就是说，它查找 `_service._proto.name`。为了适应发布非标准名称的 SRV 记录的服务，如果 service 和 proto 都是空字符串，则 LookupSRV 直接查找 name。

​	返回的服务名已验证为格式正确的表示格式的域名。如果响应包含无效名称，则这些记录将被过滤掉，并且将返回错误以及剩余结果(如果有)。

### type [TCPAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=21) 

``` go linenums="1"
type TCPAddr struct {
	IP   IP
	Port int
	Zone string // IPv6 范围限定符
}
```

​	TCPAddr结构体表示 TCP 端点的地址。

#### func [ResolveTCPAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=84) 

``` go linenums="1"
func ResolveTCPAddr(network, address string) (*TCPAddr, error)
```

​	ResolveTCPAddr函数返回 TCP 端点的地址。

​	network 必须是 TCP 网络名称。

​	如果 address 参数中的主机不是字面 IP 地址，或者端口不是字面端口号，则 ResolveTCPAddr 将地址解析为 TCP 端点的地址。否则，它将地址解析为字面 IP 地址和端口号对。address 参数可以使用主机名，但不建议这样做，因为它最多只会返回主机名的一个 IP 地址。

​	有关 network 和 address 参数的说明，请参见 func Dial。

#### func [TCPAddrFromAddrPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=102)  <- go1.18

``` go linenums="1"
func TCPAddrFromAddrPort(addr netip.AddrPort) *TCPAddr
```

​	TCPAddrFromAddrPort函数将 addr 转换为 TCPAddr。如果 addr.IsValid() 为 false，则返回的 TCPAddr 将包含一个空的 IP 字段，表示不指定地址族的未指定地址。

#### (*TCPAddr) [AddrPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=32)  <- go1.18

``` go linenums="1"
func (a *TCPAddr) AddrPort() netip.AddrPort
```

​	AddrPort方法将 TCPAddr a 转换为 netip.AddrPort。

​	如果 a.Port 无法适应 uint16，则会静默截断。

​	如果 a 为 nil，则返回零值。

#### (*TCPAddr) [Network](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=42) 

``` go linenums="1"
func (a *TCPAddr) Network() string
```

​	Network方法返回地址的网络名称，即"tcp"。

#### (*TCPAddr) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=44) 

``` go linenums="1"
func (a *TCPAddr) String() string
```

### type [TCPConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=112) 

``` go linenums="1"
type TCPConn struct {
	// 包含过滤或未公开的字段
}
```

​	TCPConn结构体是TCP网络连接的Conn接口实现。

#### func [DialTCP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=242) 

``` go linenums="1"
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
```

​	DialTCP函数用于TCP网络连接。

​	network参数必须是TCP网络的名称，有关详细信息，请参见func Dial。

​	如果laddr是nil，则会自动选择本地地址。如果raddr的IP字段为nil或未指定IP地址，则会假定为本地系统。

#### (*TCPConn) [Close](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=203) 

``` go linenums="1"
func (c *TCPConn) Close() error
```

​	Close方法关闭连接。

#### (*TCPConn) [CloseRead](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=139) 

``` go linenums="1"
func (c *TCPConn) CloseRead() error
```

​	CloseRead方法关闭TCP连接的读取侧。大多数调用方应该使用Close。

#### (*TCPConn) [CloseWrite](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=151) 

``` go linenums="1"
func (c *TCPConn) CloseWrite() error
```

​	CloseWrite方法关闭TCP连接的写入侧。大多数调用方应该使用Close。

#### (*TCPConn) [File](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=298) 

``` go linenums="1"
func (c *TCPConn) File() (f *os.File, err error)
```

​	File方法返回基础os.File的副本。调用方有责任在完成后关闭f。关闭c不会影响f，关闭f也不会影响c。

​	返回的os.File的文件描述符与连接的不同。尝试使用此副本更改原始文件的属性可能会或可能不会产生预期的效果。

#### (*TCPConn) [LocalAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=217) 

``` go linenums="1"
func (c *TCPConn) LocalAddr() Addr
```

​	LocalAddr方法返回本地网络地址。返回的Addr被所有调用LocalAddr共享，因此请勿修改它。

#### (*TCPConn) [Read](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=179) 

``` go linenums="1"
func (c *TCPConn) Read(b []byte) (int, error)
```

Read方法实现了Conn的Read方法。

#### (*TCPConn) [ReadFrom](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=126) 

``` go linenums="1"
func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)
```

​	ReadFrom方法实现了io.ReaderFrom的ReadFrom方法。

#### (*TCPConn) [RemoteAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=227) 

``` go linenums="1"
func (c *TCPConn) RemoteAddr() Addr
```

​	RemoteAddr方法返回远程网络地址。返回的Addr由所有RemoteAddr调用共享，因此请不要修改它。

#### (*TCPConn) [SetDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=235) 

``` go linenums="1"
func (c *TCPConn) SetDeadline(t time.Time) error
```

​	SetDeadline方法实现了Conn的SetDeadline方法。

#### (*TCPConn) [SetKeepAlive](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=185) 

``` go linenums="1"
func (c *TCPConn) SetKeepAlive(keepalive bool) error
```

​	SetKeepAlive方法设置操作系统是否应在连接上发送保持活动消息。

#### (*TCPConn) [SetKeepAlivePeriod](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=196)  <- go1.2

``` go linenums="1"
func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error
```

​	SetKeepAlivePeriod方法设置保持活动消息之间的时间间隔。

#### (*TCPConn) [SetLinger](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=173) 

``` go linenums="1"
func (c *TCPConn) SetLinger(sec int) error
```

​	SetLinger方法设置在仍有等待发送或确认的数据的连接上执行Close时的行为。

​	如果sec < 0(默认值)，则操作系统在后台完成发送数据。

​	如果sec == 0，则操作系统丢弃任何未发送或未确认的数据。

​	如果sec > 0，则数据像sec < 0一样在后台发送。在某些操作系统上，经过sec秒后，任何剩余的未发送数据可能会被丢弃。

#### (*TCPConn) [SetNoDelay](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=210) 

``` go linenums="1"
func (c *TCPConn) SetNoDelay(noDelay bool) error
```

​	SetNoDelay方法控制操作系统是否应推迟数据包传输，以期望发送较少的数据包(Nagle算法)。默认值为true(无延迟)，这意味着在Write之后尽快发送数据。

#### (*TCPConn) [SetReadBuffer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=269) 

``` go linenums="1"
func (c *TCPConn) SetReadBuffer(bytes int) error
```

​	SetReadBuffer方法设置与连接关联的操作系统接收缓冲区的大小。

#### (*TCPConn) [SetReadDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=246) 

``` go linenums="1"
func (c *TCPConn) SetReadDeadline(t time.Time) error
```

​	SetReadDeadline方法实现Conn SetReadDeadline方法。

#### (*TCPConn) [SetWriteBuffer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=281) 

``` go linenums="1"
func (c *TCPConn) SetWriteBuffer(bytes int) error
```

​	SetWriteBuffer方法设置与连接关联的操作系统传输缓冲区的大小。

#### (*TCPConn) [SetWriteDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=257) 

``` go linenums="1"
func (c *TCPConn) SetWriteDeadline(t time.Time) error
```

​	SetWriteDeadline方法实现Conn SetWriteDeadline方法。

#### (*TCPConn) [SyscallConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=118)  <- go1.9

``` go linenums="1"
func (c *TCPConn) SyscallConn() (syscall.RawConn, error)
```

​	SyscallConn方法返回一个原始的网络连接。这实现了syscall.Conn接口。

#### (*TCPConn) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=191) 

``` go linenums="1"
func (c *TCPConn) Write(b []byte) (int, error)
```

​	Write方法实现了Conn接口的Write方法。

### type [TCPListener](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=261) 

``` go linenums="1"
type TCPListener struct {
	// contains filtered or unexported fields
}
```

​	TCPListener结构体是TCP网络侦听器。客户端通常应该使用类型为Listener的变量，而不是假定为TCP。

#### func [ListenTCP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=360) 

``` go linenums="1"
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
```

​	ListenTCP函数像TCP网络的Listen一样工作。

​	网络必须是TCP网络名称;有关详细信息，请参阅func Dial。

​	如果laddr的IP字段为nil或未指定的IP地址，则ListenTCP将侦听本地系统的所有可用单播和任播IP地址。如果laddr的Port字段为0，则自动选择端口号。

#### (*TCPListener) [Accept](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=293) 

``` go linenums="1"
func (l *TCPListener) Accept() (Conn, error)
```

​	Accept方法实现了 Listener 接口的 Accept 方法；它等待下一个呼叫并返回一个通用的 Conn。

#### (*TCPListener) [AcceptTCP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=280) 

``` go linenums="1"
func (l *TCPListener) AcceptTCP() (*TCPConn, error)
```

​	AcceptTCP方法接受下一个传入的呼叫并返回新连接。

#### (*TCPListener) [Addr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=319) 

``` go linenums="1"
func (l *TCPListener) Addr() Addr
```

​	Addr方法返回监听器的网络地址，即 *TCPAddr。返回的 Addr 在所有调用 Addr 的地方都是共享的，因此不要修改它。

#### (*TCPListener) [Close](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=306) 

``` go linenums="1"
func (l *TCPListener) Close() error
```

​	Close方法停止在 TCP 地址上的监听。已经接受的连接不会关闭。

#### (*TCPListener) [File](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=340) 

``` go linenums="1"
func (l *TCPListener) File() (f *os.File, err error)
```

​	File方法返回底层的 os.File 的副本。调用者有责任在完成后关闭 f。关闭 l 不影响 f，关闭 f 不影响 l。

​	返回的 os.File 的文件描述符与连接的不同。尝试使用此副本更改原始连接的属性可能会或可能不会产生预期的效果。

#### (*TCPListener) [SetDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=323) 

``` go linenums="1"
func (l *TCPListener) SetDeadline(t time.Time) error
```

​	SetDeadline方法设置监听器关联的截止日期。零时间值禁用截止日期。

#### (*TCPListener) [SyscallConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/tcpsock.go;l=271)  <- go1.10

``` go linenums="1"
func (l *TCPListener) SyscallConn() (syscall.RawConn, error)
```

​	SyscallConn方法返回原始网络连接。这实现了 syscall.Conn 接口。

​	返回的 RawConn 仅支持调用 Control。Read 和 Write 返回一个错误。

### type [UDPAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=24) 

``` go linenums="1"
type UDPAddr struct {
	IP   IP
	Port int
	Zone string // IPv6 范围地址区域
}
```

​	UDPAddr结构体表示 UDP 端点的地址。

#### func [ResolveUDPAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=87) 

``` go linenums="1"
func ResolveUDPAddr(network, address string) (*UDPAddr, error)
```

​	ResolveUDPAddr函数返回 UDP 终点地址。

​	network 必须是 UDP 网络名称。

​	如果 address 参数中的主机不是字面 IP 地址或端口不是字面端口号，则 ResolveUDPAddr 将地址解析为 UDP 终点地址。否则，它将地址解析为字面 IP 地址和端口号的一对。地址参数可以使用主机名，但不建议这样做，因为它最多只会返回主机名的一个 IP 地址。

​	有关网络和地址参数的说明，请参见 func Dial。

#### func [UDPAddrFromAddrPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=105)  <- go1.18

``` go linenums="1"
func UDPAddrFromAddrPort(addr netip.AddrPort) *UDPAddr
```

​	UDPAddrFromAddrPort函数将 addr 转换为 UDPAddr。如果 addr.IsValid() 为 false，则返回的 UDPAddr 将包含一个 nil IP 字段，表示未指定地址簇的未指定地址。

#### (*UDPAddr) [AddrPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=35)  <- go1.18

``` go linenums="1"
func (a *UDPAddr) AddrPort() netip.AddrPort
```

​	AddrPort方法将 UDPAddr a 转换为 netip.AddrPort。

​	如果 a.Port 不适合 uint16，则它将被静默截断。

​	如果 a 为 nil，则返回零值。

#### (*UDPAddr) [Network](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=45) 

``` go linenums="1"
func (a *UDPAddr) Network() string
```

​	Network方法返回地址的网络名称，即 "udp"。

#### (*UDPAddr) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=47) 

``` go linenums="1"
func (a *UDPAddr) String() string
```

### type [UDPConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=122) 

``` go linenums="1"
type UDPConn struct {
	// contains filtered or unexported fields
}
```

​	UDPConn结构体是 UDP 网络连接的 Conn 和 PacketConn 接口实现。

#### func [DialUDP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=290) 

``` go linenums="1"
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
```

​	DialUDP函数用于 UDP 网络的 Dial。

​	network 必须是 UDP 网络名称；有关详细信息，请参见 func Dial。

​	如果 laddr 为 nil，则自动选择本地地址。如果 raddr 的 IP 字段为 nil 或未指定 IP 地址，则假定为本地系统。

#### func [ListenMulticastUDP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=353) 

``` go linenums="1"
func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
```

​	ListenMulticastUDP函数用于 UDP 网络的 ListenPacket，但它接受特定网络接口上的组地址。

​	network 必须是 UDP 网络名称；有关详细信息，请参见 func Dial。

​	ListenMulticastUDP 监听本地系统的所有可用 IP 地址，包括组播 IP 地址。如果 ifi 为 nil，则 ListenMulticastUDP 使用系统分配的多播接口，尽管这不被推荐，因为分配取决于平台，有时可能需要路由配置。如果 gaddr 的 Port 字段为 0，则自动选择端口号。

​	ListenMulticastUDP 只是简单、小型应用的方便。对于一般用途，有 golang.org/x/net/ipv4 和 golang.org/x/net/ipv6 包。

​	请注意，ListenMulticastUDP将在IPPROTO_IP下将IP_MULTICAST_LOOP套接字选项设置为0，以禁用组播数据包的回送。

#### func [ListenUDP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=316) 

``` go linenums="1"
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
```

​	ListenUDP函数的行为类似于UDP网络的ListenPacket。

​	网络必须是UDP网络名称；有关详细信息，请参见Dial函数。

​	如果laddr的IP字段为nil或未指定的IP地址，则ListenUDP会在本地系统的所有可用IP地址上进行侦听，但不包括多播IP地址。如果laddr的Port字段为0，则自动选择端口号。

#### (*UDPConn) [Close](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=203) 

``` go linenums="1"
func (c *UDPConn) Close() error
```

​	Close方法关闭连接。

#### (*UDPConn) [File](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=298) 

``` go linenums="1"
func (c *UDPConn) File() (f *os.File, err error)
```

​	File方法返回底层os.File的副本。调用者有责任在完成后关闭f。关闭c不会影响f，关闭f也不会影响c。

​	返回的os.File的文件描述符与连接的不同。尝试使用此副本更改原始的属性可能会产生预期或非预期的影响。

#### (*UDPConn) [LocalAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=217) 

``` go linenums="1"
func (c *UDPConn) LocalAddr() Addr
```

​	LocalAddr方法返回本地网络地址。返回的Addr由LocalAddr的所有调用共享，因此不要修改它。

#### (*UDPConn) [Read](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=179) 

``` go linenums="1"
func (c *UDPConn) Read(b []byte) (int, error)
```

​	Read方法实现Conn Read方法。

#### (*UDPConn) [ReadFrom](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=157) 

``` go linenums="1"
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)
```

​	ReadFrom方法实现PacketConn ReadFrom方法。

#### (*UDPConn) [ReadFromUDP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=136) 

``` go linenums="1"
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)
```

​	ReadFromUDP方法的行为类似于ReadFrom，但返回一个UDPAddr。

#### (*UDPConn) [ReadFromUDPAddrPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=171)  <- go1.18

``` go linenums="1"
func (c *UDPConn) ReadFromUDPAddrPort(b []byte) (n int, addr netip.AddrPort, err error)
```

​	ReadFromUDPAddrPort方法的行为类似于ReadFrom，但返回一个netip.AddrPort。

​	如果c绑定到未指定的地址，则返回的netip.AddrPort的地址可能是一个IPv4映射的IPv6地址。使用netip.Addr.Unmap获取不带IPv6前缀的地址。

#### (*UDPConn) [ReadMsgUDP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=189)  <- go1.1

``` go linenums="1"
func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)
```

​	ReadMsgUDP方法从c读取消息，将有效负载复制到b中，并将关联的带外数据复制到oob中。它返回复制到b中的字节数，复制到oob中的字节数，设置在消息上的标志以及消息的源地址。

​	可以使用golang.org/x/net/ipv4和golang.org/x/net/ipv6包来操作oob中的IP级套接字选项。

#### (*UDPConn) [ReadMsgUDPAddrPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=199)  <- go1.18

``` go linenums="1"
func (c *UDPConn) ReadMsgUDPAddrPort(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error)
```

​	ReadMsgUDPAddrPort方法与ReadMsgUDP方法类似，但返回的是netip.AddrPort而不是UDPAddr。

#### (*UDPConn) [RemoteAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=227) 

``` go linenums="1"
func (c *UDPConn) RemoteAddr() Addr
```

​	RemoteAddr方法返回远程网络地址。返回的 Addr 在所有 RemoteAddr 的调用中共享，因此不要修改它。

#### (*UDPConn) [SetDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=235) 

``` go linenums="1"
func (c *UDPConn) SetDeadline(t time.Time) error
```

​	SetDeadline方法实现了 Conn SetDeadline 方法。

#### (*UDPConn) [SetReadBuffer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=269) 

``` go linenums="1"
func (c *UDPConn) SetReadBuffer(bytes int) error
```

​	SetReadBuffer方法设置与连接关联的操作系统接收缓冲区的大小。

#### (*UDPConn) [SetReadDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=246) 

``` go linenums="1"
func (c *UDPConn) SetReadDeadline(t time.Time) error
```

​	SetReadDeadline方法实现了 Conn SetReadDeadline 方法。

#### (*UDPConn) [SetWriteBuffer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=281) 

``` go linenums="1"
func (c *UDPConn) SetWriteBuffer(bytes int) error
```

​	SetWriteBuffer方法设置与连接关联的操作系统传输缓冲区的大小。

#### (*UDPConn) [SetWriteDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=257) 

``` go linenums="1"
func (c *UDPConn) SetWriteDeadline(t time.Time) error
```

​	SetWriteDeadline方法实现了 Conn SetWriteDeadline 方法。

#### (*UDPConn) [SyscallConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=128)  <- go1.9

``` go linenums="1"
func (c *UDPConn) SyscallConn() (syscall.RawConn, error)
```

​	SyscallConn方法返回原始网络连接。这实现了 syscall.Conn 接口。

#### (*UDPConn) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=191) 

``` go linenums="1"
func (c *UDPConn) Write(b []byte) (int, error)
```

​	Write方法实现了 Conn Write 方法。

#### (*UDPConn) [WriteMsgUDP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=258)  <- go1.1

``` go linenums="1"
func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)
```

​	WriteMsgUDP方法向 addr 写入消息，如果 c 未连接，则通过 c 写入，否则通过 c 的远程地址写入(此时 addr 必须为 nil)。从 b 复制有效载荷，从 oob 复制关联的带外数据。返回写入的有效载荷和带外字节数。

​	包 golang.org/x/net/ipv4 和 golang.org/x/net/ipv6 可用于操作 oob 中的 IP 级 socket 选项。

#### (*UDPConn) [WriteMsgUDPAddrPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=270)  <- go1.18

``` go linenums="1"
func (c *UDPConn) WriteMsgUDPAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error)
```

​	WriteMsgUDPAddrPort方法与WriteMsgUDP方法类似，但是接受netip.AddrPort而不是UDPAddr。

#### (*UDPConn) [WriteTo](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=235) 

``` go linenums="1"
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)
```

​	WriteTo方法实现了PacketConn WriteTo方法。

##### WriteTo Example
``` go linenums="1"
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

#### (*UDPConn) [WriteToUDP](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=211) 

``` go linenums="1"
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
```

​	WriteToUDP方法类似于WriteTo，但是需要一个UDPAddr参数。

#### (*UDPConn) [WriteToUDPAddrPort](https://cs.opensource.google/go/go/+/go1.20.1:src/net/udpsock.go;l=223)  <- go1.18

``` go linenums="1"
func (c *UDPConn) WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error)
```

​	WriteToUDPAddrPort方法类似于WriteTo，但是需要一个netip.AddrPort参数。

### type [UnixAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=22) 

``` go linenums="1"
type UnixAddr struct {
	Name string
	Net  string
}
```

​	UnixAddr结构体表示Unix域套接字终端的地址。

#### func [ResolveUnixAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=57) 

``` go linenums="1"
func ResolveUnixAddr(network, address string) (*UnixAddr, error)
```

​	ResolveUnixAddr函数返回Unix域套接字终端的地址。

​	network参数必须是Unix网络名称。

​	有关network和address参数的说明，请参见func Dial。

#### (*UnixAddr) [Network](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=29) 

``` go linenums="1"
func (a *UnixAddr) Network() string
```

​	Network方法返回地址的网络名称，即"unix"、"unixgram"或"unixpacket"。

#### (*UnixAddr) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=33) 

``` go linenums="1"
func (a *UnixAddr) String() string
```

### type [UnixConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=68) 

``` go linenums="1"
type UnixConn struct {
	// 包含过滤或未公开的字段
}
```

​	UnixConn结构体是用于Unix域套接字连接的Conn接口的实现。

#### func [DialUnix](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=203) 

``` go linenums="1"
func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)
```

​	DialUnix函数类似于Unix网络的Dial函数。

​	network参数必须是Unix网络名称；有关详细信息，请参见func Dial。

​	如果laddr非nil，则用作连接的本地地址。

#### func [ListenUnixgram](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=337) 

``` go linenums="1"
func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)
```

​	ListenUnixgram函数类似于Unix网络的ListenPacket函数。

​	network参数必须是"unixgram"。

#### (*UnixConn) [Close](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=203) 

``` go linenums="1"
func (c *UnixConn) Close() error
```

​	Close方法关闭连接。

#### (*UnixConn) [CloseRead](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=83)  <- go1.1

``` go linenums="1"
func (c *UnixConn) CloseRead() error
```

​	CloseRead方法关闭Unix域连接的读取端。大多数调用方应该只使用Close。

#### (*UnixConn) [CloseWrite](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=95)  <- go1.1

``` go linenums="1"
func (c *UnixConn) CloseWrite() error
```

​	CloseWrite方法关闭Unix域连接的写入端。大多数调用方应该只使用Close。

#### (*UnixConn) [File](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=298) 

``` go linenums="1"
func (c *UnixConn) File() (f *os.File, err error)
```

​	File方法返回底层os.File的副本。调用方有责任在完成后关闭f。关闭c不会影响f，关闭f也不会影响c。

​	返回的os.File的文件描述符与连接的不同。尝试使用此副本更改原始文件的属性可能会或可能不会产生所需的效果。

#### (*UnixConn) [LocalAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=217) 

``` go linenums="1"
func (c *UnixConn) LocalAddr() Addr
```

​	LocalAddr方法返回本地网络地址。返回的Addr由LocalAddr的所有调用共享，因此不要修改它。

#### (*UnixConn) [Read](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=179) 

``` go linenums="1"
func (c *UnixConn) Read(b []byte) (int, error)
```

​	Read方法实现了Conn Read方法。

#### (*UnixConn) [ReadFrom](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=118) 

``` go linenums="1"
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)
```

​	ReadFrom方法实现了PacketConn ReadFrom方法。

#### (*UnixConn) [ReadFromUnix](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=106) 

``` go linenums="1"
func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error)
```

​	ReadFromUnix方法类似于ReadFrom，但返回一个UnixAddr。

#### (*UnixConn) [ReadMsgUnix](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=139) 

``` go linenums="1"
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)
```

​	ReadMsgUnix方法从c读取消息，将有效负载复制到b，将相关的带外数据复制到oob。它返回复制到b的字节数，复制到oob的字节数，消息上设置的标志以及消息的源地址。

​	请注意，如果len(b) == 0且len(oob) > 0，则此函数仍将从连接中读取(并丢弃)1个字节。

#### (*UnixConn) [RemoteAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=227) 

``` go linenums="1"
func (c *UnixConn) RemoteAddr() Addr
```

​	RemoteAddr方法返回远程网络地址。返回的Addr由RemoteAddr的所有调用共享，因此不要修改它。

#### (*UnixConn) [SetDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=235) 

``` go linenums="1"
func (c *UnixConn) SetDeadline(t time.Time) error
```

​	SetDeadline方法实现了Conn SetDeadline方法。

#### (*UnixConn) [SetReadBuffer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=269) 

``` go linenums="1"
func (c *UnixConn) SetReadBuffer(bytes int) error
```

​	SetReadBuffer方法设置与连接关联的操作系统接收缓冲区的大小。

#### (*UnixConn) [SetReadDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=246) 

``` go linenums="1"
func (c *UnixConn) SetReadDeadline(t time.Time) error
```

​	SetReadDeadline方法实现了Conn SetReadDeadline方法。

#### (*UnixConn) [SetWriteBuffer](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=281) 

``` go linenums="1"
func (c *UnixConn) SetWriteBuffer(bytes int) error
```

​	SetWriteBuffer方法设置与连接关联的操作系统传输缓冲区的大小。

#### (*UnixConn) [SetWriteDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=257) 

``` go linenums="1"
func (c *UnixConn) SetWriteDeadline(t time.Time) error
```

​	SetWriteDeadline方法实现了Conn SetWriteDeadline方法。

#### (*UnixConn) [SyscallConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=74)  <- go1.9

``` go linenums="1"
func (c *UnixConn) SyscallConn() (syscall.RawConn, error)
```

​	SyscallConn方法返回原始网络连接。这实现了syscall.Conn接口。

#### (*UnixConn) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=191) 

``` go linenums="1"
func (c *UnixConn) Write(b []byte) (int, error)
```

​	Write方法实现了Conn Write方法。

#### (*UnixConn) [WriteMsgUnix](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=184) 

``` go linenums="1"
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)
```

​	WriteMsgUnix方法向addr通过c写入消息，从b复制有效负载，从oob复制相关的带外数据。它返回已写入有效负载和带外字节数。

​	请注意，如果len(b) == 0且len(oob) > 0，则此函数仍将向连接写入1个字节。

#### (*UnixConn) [WriteTo](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=163) 

``` go linenums="1"
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error)
```

​	WriteTo方法实现了PacketConn WriteTo方法。

#### (*UnixConn) [WriteToUnix](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=151) 

``` go linenums="1"
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error)
```

​	WriteToUnix方法类似于WriteTo，但需要一个UnixAddr参数。

### type [UnixListener](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=220) 

``` go linenums="1"
type UnixListener struct {
	// 包含过滤或未公开的字段
}
```

​	UnixListener结构体是Unix域套接字侦听器。客户端通常应使用类型为Listener的变量，而不是假定Unix域套接字。

#### func [ListenUnix](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=317) 

``` go linenums="1"
func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)
```

​	ListenUnix函数类似于Unix网络的Listen函数。

​	network必须是"unix"或"unixpacket"。

#### (*UnixListener) [Accept](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=256) 

``` go linenums="1"
func (l *UnixListener) Accept() (Conn, error)
```

​	Accept方法实现了Listener接口中的Accept方法。返回的连接将是*UnixConn类型。

#### (*UnixListener) [AcceptUnix](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=243) 

``` go linenums="1"
func (l *UnixListener) AcceptUnix() (*UnixConn, error)
```

​	AcceptUnix方法接受下一个传入的调用并返回新的连接。

#### (*UnixListener) [Addr](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=282) 

``` go linenums="1"
func (l *UnixListener) Addr() Addr
```

​	Addr方法返回侦听器的网络地址。返回的Addr由所有Addr调用共享，因此不要修改它。

#### (*UnixListener) [Close](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=269) 

``` go linenums="1"
func (l *UnixListener) Close() error
```

​	Close方法停止侦听Unix地址。已接受的连接不会关闭。

#### (*UnixListener) [File](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=303) 

``` go linenums="1"
func (l *UnixListener) File() (f *os.File, err error)
```

​	File方法返回底层os.File的副本。调用者有责任在完成后关闭f。关闭l不影响f，关闭f也不影响l。

​	返回的os.File的文件描述符与连接的不同。尝试使用此副本更改原始文件的属性可能会产生预期的效果，也可能不会。

#### (*UnixListener) [SetDeadline](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=286) 

``` go linenums="1"
func (l *UnixListener) SetDeadline(t time.Time) error
```

​	SetDeadline方法设置与侦听器关联的截止日期。零时间值禁用截止日期。

#### (*UnixListener) [SetUnlinkOnClose](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock_posix.go;l=215)  <- go1.8

``` go linenums="1"
func (l *UnixListener) SetUnlinkOnClose(unlink bool)
```

​	SetUnlinkOnClose方法设置在关闭侦听器时是否应从文件系统中删除底层套接字文件。

​	默认行为是仅在包net创建它时取消连接套接字文件。也就是说，当使用调用Listen或ListenUnix创建侦听器和底层套接字文件时，默认情况下，关闭侦听器将删除套接字文件。但是，如果通过调用FileListener创建侦听器来使用已经存在的套接字文件，则默认情况下，关闭侦听器不会删除套接字文件。

#### (*UnixListener) [SyscallConn](https://cs.opensource.google/go/go/+/go1.20.1:src/net/unixsock.go;l=234)  <- go1.10

``` go linenums="1"
func (l *UnixListener) SyscallConn() (syscall.RawConn, error)
```

​	SyscallConn方法返回原始网络连接。这实现了syscall.Conn接口。

​	返回的RawConn仅支持调用Control。读取和写入返回错误。

### type [UnknownNetworkError](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=573) 

``` go linenums="1"
type UnknownNetworkError string
```

#### (UnknownNetworkError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=575) 

``` go linenums="1"
func (e UnknownNetworkError) Error() string
```

#### (UnknownNetworkError) [Temporary](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=577) 

``` go linenums="1"
func (e UnknownNetworkError) Temporary() bool
```

#### (UnknownNetworkError) [Timeout](https://cs.opensource.google/go/go/+/go1.20.1:src/net/net.go;l=576) 

``` go linenums="1"
func (e UnknownNetworkError) Timeout() bool
```

## Notes

## 缺陷

- 在JS和Windows上，未实现FileConn、FileListener和FilePacketConn函数。
- 在JS上，未实现与Interface相关的方法和函数。
- 在AIX、DragonFly BSD、NetBSD、OpenBSD、Plan 9和Solaris上，未实现Interface的MulticastAddrs方法。
- 在每个POSIX平台上，使用ReadFrom或ReadFromIP方法从"ip4"网络读取可能不会返回完整的IPv4数据包，包括其标头，即使有空间可用。即使Read或ReadMsgIP可以返回完整的数据包，这也可能发生。因此，如果接收完整数据包很重要，则建议不使用这些方法。
- Go 1兼容性指南使我们无法更改这些方法的行为。请改用Read或ReadMsgIP。
- 在JS和Plan 9上，未实现与IPConn相关的方法和函数。
- 在Windows上，未实现IPConn的File方法。
- 在DragonFly BSD和OpenBSD上，侦听"tcp"和"udp"网络不会同时侦听IPv4和IPv6连接。这是因为IPv4流量不会被路由到IPv6套接字——如果要支持两个地址家族，则需要两个单独的套接字。有关详细信息，请参见inet6(4)。
- 在Windows上，syscall.RawConn的Write方法无法与运行时的网络轮询器集成。它不能等待连接变为可写，并且不尊重期限。如果用户提供的回调返回false，则Write方法将立即失败。
- 在JS和Plan 9上，未实现syscall.RawConn的Control、Read和Write方法。
- 在JS和Windows上，未实现TCPConn和TCPListener的File方法。
- 在Plan 9上，未实现UDPConn的ReadMsgUDP和WriteMsgUDP方法。
- 在Windows上，未实现UDPConn的File方法。
- 在JS上，未实现与UDPConn相关的方法和函数。
- 在JS和Plan 9上，未实现与UnixConn和UnixListener相关的方法和函数。
- 在Windows上，与UnixConn和UnixListener相关的方法和函数不适用于"unixgram"和"unixpacket"。