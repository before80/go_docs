+++
title = "gipv4"
date = 2024-03-21T17:53:01+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gipv4](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gipv4)

Package gipv4 provides useful API for IPv4 address handling.

​	软件包 gipv4 为 IPv4 地址处理提供了有用的 API。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func GetHostByName

```go
func GetHostByName(hostname string) (string, error)
```

GetHostByName returns the IPv4 address corresponding to a given Internet host name.

​	GetHostByName 返回与给定 Internet 主机名对应的 IPv4 地址。

#### func GetHostsByName

```go
func GetHostsByName(hostname string) ([]string, error)
```

GetHostsByName returns a list of IPv4 addresses corresponding to a given Internet host name.

​	GetHostsByName 返回与给定 Internet 主机名对应的 IPv4 地址列表。

#### func GetIntranetIp

```go
func GetIntranetIp() (ip string, err error)
```

GetIntranetIp retrieves and returns the first intranet ip of current machine.

​	GetIntranetIp 检索并返回当前计算机的第一个 Intranet IP。

#### func GetIntranetIpArray

```go
func GetIntranetIpArray() (ips []string, err error)
```

GetIntranetIpArray retrieves and returns the intranet ip list of current machine.

​	GetIntranetIpArray 检索并返回当前计算机的 Intranet IP 列表。

#### func GetIpArray

```go
func GetIpArray() (ips []string, err error)
```

GetIpArray retrieves and returns all the ip of current host.

​	GetIpArray 检索并返回当前主机的所有 ip。

#### func GetMac

```go
func GetMac() (mac string, err error)
```

GetMac retrieves and returns the first mac address of current host.

​	GetMac 检索并返回当前主机的第一个 mac 地址。

#### func GetMacArray

```go
func GetMacArray() (macs []string, err error)
```

GetMacArray retrieves and returns all the mac address of current host.

​	GetMacArray 检索并返回当前主机的所有 MAC 地址。

#### func GetNameByAddr

```go
func GetNameByAddr(ipAddress string) (string, error)
```

GetNameByAddr returns the Internet host name corresponding to a given IP address.

​	GetNameByAddr 返回与给定 IP 地址对应的 Internet 主机名。

#### func GetSegment

```go
func GetSegment(ip string) string
```

GetSegment returns the segment of given ip address. Eg: 192.168.2.102 -> 192.168.2

​	GetSegment 返回给定 IP 地址的段。例如：192.168.2.102 -> 192.168.2

#### func Ip2long

```go
func Ip2long(ip string) uint32
```

Ip2long converts ip address to an uint32 integer.

​	Ip2long 将 ip 地址转换为 uint32 整数。

#### func IsIntranet

```go
func IsIntranet(ip string) bool
```

IsIntranet checks and returns whether given ip an intranet ip.

​	IsIntranet 检查并返回给定的 ip 是否为 Intranet IP。

Local: 127.0.0.1 A: 10.0.0.0–10.255.255.255 B: 172.16.0.0–172.31.255.255 C: 192.168.0.0–192.168.255.255

​	本地：127.0.0.1 答：10.0.0.0–10.255.255.255 B：172.16.0.0–172.31.255.255 C：192.168.0.0–192.168.255.255

#### func Long2ip

```go
func Long2ip(long uint32) string
```

Long2ip converts an uint32 integer ip address to its string type address.

​	Long2ip 将 uint32 整数 IP 地址转换为其字符串类型地址。

#### func MustGetIntranetIp

```go
func MustGetIntranetIp() string
```

MustGetIntranetIp performs as GetIntranetIp, but it panics if any error occurs.

​	MustGetIntranetIp 作为 GetIntranetIp 执行，但如果发生任何错误，它会崩溃。

#### func ParseAddress

```go
func ParseAddress(address string) (string, int)
```

ParseAddress parses `address` to its ip and port. Eg: 192.168.1.1:80 -> 192.168.1.1, 80

​	ParseAddress 解析 `address` 为其 IP 和端口。例如：192.168.1.1：80 -> 192.168.1.1， 80

#### func Validate

```go
func Validate(ip string) bool
```

Validate checks whether given `ip` a valid IPv4 address.

​	验证检查是否给定 `ip` 了有效的 IPv4 地址。

## 类型

This section is empty.