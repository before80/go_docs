+++
title = "gipv4"
date = 2024-03-21T17:53:01+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gipv4

Package gipv4 provides useful API for IPv4 address handling.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func GetHostByName 

``` go
func GetHostByName(hostname string) (string, error)
```

GetHostByName returns the IPv4 address corresponding to a given Internet host name.

##### func GetHostsByName 

``` go
func GetHostsByName(hostname string) ([]string, error)
```

GetHostsByName returns a list of IPv4 addresses corresponding to a given Internet host name.

##### func GetIntranetIp 

``` go
func GetIntranetIp() (ip string, err error)
```

GetIntranetIp retrieves and returns the first intranet ip of current machine.

##### func GetIntranetIpArray 

``` go
func GetIntranetIpArray() (ips []string, err error)
```

GetIntranetIpArray retrieves and returns the intranet ip list of current machine.

##### func GetIpArray 

``` go
func GetIpArray() (ips []string, err error)
```

GetIpArray retrieves and returns all the ip of current host.

##### func GetMac 

``` go
func GetMac() (mac string, err error)
```

GetMac retrieves and returns the first mac address of current host.

##### func GetMacArray 

``` go
func GetMacArray() (macs []string, err error)
```

GetMacArray retrieves and returns all the mac address of current host.

##### func GetNameByAddr 

``` go
func GetNameByAddr(ipAddress string) (string, error)
```

GetNameByAddr returns the Internet host name corresponding to a given IP address.

##### func GetSegment 

``` go
func GetSegment(ip string) string
```

GetSegment returns the segment of given ip address. Eg: 192.168.2.102 -> 192.168.2

##### func Ip2long 

``` go
func Ip2long(ip string) uint32
```

Ip2long converts ip address to an uint32 integer.

##### func IsIntranet 

``` go
func IsIntranet(ip string) bool
```

IsIntranet checks and returns whether given ip an intranet ip.

Local: 127.0.0.1 A: 10.0.0.0--10.255.255.255 B: 172.16.0.0--172.31.255.255 C: 192.168.0.0--192.168.255.255

##### func Long2ip 

``` go
func Long2ip(long uint32) string
```

Long2ip converts an uint32 integer ip address to its string type address.

##### func MustGetIntranetIp 

``` go
func MustGetIntranetIp() string
```

MustGetIntranetIp performs as GetIntranetIp, but it panics if any error occurs.

##### func ParseAddress 

``` go
func ParseAddress(address string) (string, int)
```

ParseAddress parses `address` to its ip and port. Eg: 192.168.1.1:80 -> 192.168.1.1, 80

##### func Validate 

``` go
func Validate(ip string) bool
```

Validate checks whether given `ip` a valid IPv4 address.

### Types 

This section is empty.