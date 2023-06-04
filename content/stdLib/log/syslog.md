+++
title = "syslog"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# syslog

https://pkg.go.dev/log/syslog@go1.20.1

​	syslog包提供了一个简单的接口来访问系统日志服务。它可以使用 UNIX 域套接字、UDP 或 TCP 发送消息到 syslog 守护进程。

​	只需要调用一次 Dial函数。在写入失败时，syslog 客户端将尝试重新连接服务器并重试写入。

​	syslog 包被冻结，不再接受新功能。一些外部包提供更多功能。参见：

```
https://godoc.org/?q=syslog
```

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func NewLogger 

```go 
func NewLogger(p Priority, logFlag int) (*log.Logger, error)
```

​	NewLogger函数创建一个 log.Logger，其输出被写入具有指定优先级的系统日志服务，该优先级是 syslog 设施和严重性的组合。logFlag 参数是传递给 log.New 以创建 Logger 的 flag 集。

## 类型

### type Priority 

```go 
type Priority int
```

​	Priority 是 syslog 设施和严重性的组合。例如，LOG_ALERT | LOG_FTP 从 FTP 设施发送警报严重性消息。默认的严重性为 LOG_EMERG，默认的设施为 LOG_KERN。

```go 
const (

   	// 来自 /usr/include/sys/syslog.h。
	// 在 Linux、BSD 和 OS X 上相同。
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)
const (

    //来自/usr/include/sys/syslog.h。
	// 这些在Linux、BSD和OS X上与LOG_FTP相同。  
	LOG_KERN Priority = iota << 3
	LOG_USER
	LOG_MAIL
	LOG_DAEMON
	LOG_AUTH
	LOG_SYSLOG
	LOG_LPR
	LOG_NEWS
	LOG_UUCP
	LOG_CRON
	LOG_AUTHPRIV
	LOG_FTP

	LOG_LOCAL0
	LOG_LOCAL1
	LOG_LOCAL2
	LOG_LOCAL3
	LOG_LOCAL4
	LOG_LOCAL5
	LOG_LOCAL6
	LOG_LOCAL7
)
```

### type Writer 

```go 
type Writer struct {
    // 包含已过滤或未导出的字段
}
```

​	Writer 是与 syslog 服务器的连接。

#### func Dial 

```go 
func Dial(network, raddr string, priority Priority, tag string) (*Writer, error)
```

​	Dial函数通过在指定网络上连接到地址 raddr 来建立与日志守护程序的连接。对返回的 Writer 的每次写入都会使用设施和严重性(从 priority)和标记发送日志消息。如果标记为空，则使用 os.Args[0]。如果 network 为空，则 Dial 将连接到本地 syslog 服务器。否则，请参见 net.Dial 的文档以获取 network 和 raddr 的有效值。

##### Dial Example

```go 
package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.Dial("tcp", "localhost:1234",
		syslog.LOG_WARNING|syslog.LOG_DAEMON, "demotag")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(sysLog, "This is a daemon warning with demotag.")
	sysLog.Emerg("And this is a daemon emergency with demotag.")
}

```



#### func New 

```go 
func New(priority Priority, tag string) (*Writer, error)
```

​	New 函数用于建立与系统日志守护进程的新连接。返回的 writer 对象可用于向日志守护进程发送日志消息，其中每个消息都具有给定的优先级(syslog 的设施和严重性的组合)和前缀标签。如果标签为空，则使用 os.Args[0]。

#### (*Writer) Alert 

```go 
func (w *Writer) Alert(m string) error
```

​	Alert 方法以 LOG_ALERT 严重性记录消息，忽略在 New 函数传递的严重性。

#### (*Writer) Close 

```go 
func (w *Writer) Close() error
```

​	Close 方法用于关闭与 syslog 守护进程的连接。

#### (*Writer) Crit 

```go 
func (w *Writer) Crit(m string) error
```

​	Crit 方法以 LOG_CRIT 严重性记录消息，忽略在 New 函数传递的严重性。

#### (*Writer) Debug 

```go 
func (w *Writer) Debug(m string) error
```

​	Debug 方法以 LOG_DEBUG 严重性记录消息，忽略在 New 函数传递的严重性。

#### (*Writer) Emerg 

```go 
func (w *Writer) Emerg(m string) error
```

​	Emerg 方法以 LOG_EMERG 严重性记录消息，忽略在 New 函数传递的严重性。

#### (*Writer) Err 

```go 
func (w *Writer) Err(m string) error
```

​	Err 方法以 LOG_ERR 严重性记录消息，忽略在 New 方法传递的严重性。

#### (*Writer) Info 

```go 
func (w *Writer) Info(m string) error
```

​	Info方法使用 LOG_INFO 严重性记录一条消息，忽略在 New 方法传递的严重性。

#### (*Writer) Notice 

```go 
func (w *Writer) Notice(m string) error
```

​	Notice方法使用 LOG_NOTICE 严重性记录一条消息，忽略在 New 方法传递的严重性。

#### (*Writer) Warning 

```go 
func (w *Writer) Warning(m string) error
```

​	Warning方法使用 LOG_WARNING 严重性记录一条消息，忽略在 New 方法传递的严重性。



#### (*Writer) Write 

```go 
func (w *Writer) Write(b []byte) (int, error)
```

​	Write方法将一条日志消息发送到系统日志守护程序(syslog daemon)。

## Notes

## Bugs

- 该包在 Windows 上没有实现。由于 syslog 包被冻结，因此鼓励 Windows 用户使用标准库之外的包。有关背景，请参见 [https://golang.org/issue/1108](https://golang.org/issue/1108)。
- 该包在 Plan 9 上没有实现。