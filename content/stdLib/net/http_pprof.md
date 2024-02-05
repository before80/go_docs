+++
title = "pprof"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/http/pprof@go1.20.1](https://pkg.go.dev/net/http/pprof@go1.20.1)

​	`pprof` 包通过其 HTTP 服务器以 pprof 可工具所期望的格式提供运行时分析数据。

​	通常，该包仅用于注册其 HTTP 处理程序的副作用。所有处理的路径都以 `/debug/pprof/` 开头。

​	要使用 pprof，请将此包链接到您的程序中：

```go
import _ "net/http/pprof"
```

​	如果您的应用程序尚未运行 HTTP 服务器，则需要启动一个。将 "net/http" 和 "log" 添加到您的导入项中，并将以下代码添加到您的主函数中：

```go
go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

​	默认情况下，所有在 [runtime/pprof.Profile](https://pkg.go.dev/runtime/pprof#Profile) 中列出的分析都可用（通过 [Handler](https://pkg.go.dev/net/http/pprof@go1.20.1#Handler)），还有此包中定义的 [Cmdline](https://pkg.go.dev/net/http/pprof@go1.20.1#Cmdline)、[Profile](https://pkg.go.dev/net/http/pprof@go1.20.1#Profile)、[Symbol](https://pkg.go.dev/net/http/pprof@go1.20.1#Symbol) 和 [Trace](https://pkg.go.dev/net/http/pprof@go1.20.1#Trace) 分析。如果您没有使用 DefaultServeMux，您将需要使用您使用的 mux 注册处理程序。

## 使用示例

​	使用 pprof 工具查看堆分析文件：

```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

​	或者查看 30 秒的 CPU 分析文件：

```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

​	或者在调用 runtime.SetBlockProfileRate 之后查看 goroutine 阻塞分析文件：

```bash
go tool pprof http://localhost:6060/debug/pprof/block
```

​	或者在调用 runtime.SetMutexProfileFraction 之后查看争用互斥锁的持有者分析文件：

```bash
go tool pprof http://localhost:6060/debug/pprof/mutex
```

​	该包还导出了一个处理程序，用于为 "go tool trace" 命令提供执行跟踪数据。要收集 5 秒的执行跟踪数据：

```bash
curl -o trace.out http://localhost:6060/debug/pprof/trace?seconds=5
go tool trace trace.out
```

​	要查看所有可用的分析文件，请在浏览器中打开 http://localhost:6060/debug/pprof/。

​	要了解该功能的实际应用，请访问官方博客：[Profiling Go Programs]({{< ref "/goBlog/2011/ProfilingGoPrograms">}})



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func Cmdline 

``` go 
func Cmdline(w http.ResponseWriter, r *http.Request)
```

​	Cmdline 函数响应运行程序的命令行，其中实参由 NUL 字节分隔。该包的初始化将其注册为 `/debug/pprof/cmdline`。

#### func Handler 

``` go 
func Handler(name string) http.Handler
```

​	Handler 函数返回一个 HTTP 处理程序，用于提供指定的分析文件。可用的分析文件可以在 [runtime/pprof.Profile](https://pkg.go.dev/runtime/pprof#Profile) 中找到。

#### func Index 

``` go 
func Index(w http.ResponseWriter, r *http.Request)
```

​	Index 函数响应请求的 pprof 格式的分析文件。例如，"/debug/pprof/heap" 会提供 "heap" 分析文件。对于请求 "/debug/pprof/"，Index 会响应一个 HTML 页面，列出可用的（所有）分析文件。

#### func Profile 

``` go 
func Profile(w http.ResponseWriter, r *http.Request)
```

​	Profile 函数响应 pprof 格式的 CPU 分析文件。分析持续时间由 GET 参数中的 seconds 指定，如果未指定，则默认为 30 秒。该包的初始化将其注册为 `/debug/pprof/profile`。

#### func Symbol 

``` go 
func Symbol(w http.ResponseWriter, r *http.Request)
```

​	Symbol 函数查找请求中列出的程序计数器，响应一个表，将程序计数器映射到函数名称。该包的初始化将其注册为 `/debug/pprof/symbol`。

#### func Trace  <- go1.5

``` go 
func Trace(w http.ResponseWriter, r *http.Request)
```

​	Trace 函数响应以二进制形式的执行跟踪。跟踪持续时间由 GET 参数中的 seconds 指定，如果未指定，则默认为 1 秒。该包的初始化将其注册为 `/debug/pprof/trace`。

## 类型

This section is empty.