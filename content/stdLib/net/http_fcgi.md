+++
title = "fcgi"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# fcgi

https://pkg.go.dev/net/http/fcgi@go1.20.1



​	fcgi 包实现了 FastCGI 协议。

​	请参阅 https://fast-cgi.github.io/ 以获取原始文档的非官方镜像。

​	目前仅支持响应者（responder）角色。

## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/fcgi/child.go;l=191)

``` go 
var ErrConnClosed = errors.New("fcgi: connection to web server closed")
```

​	当处理程序尝试在与 Web 服务器的连接已关闭后读取请求的主体时，Read 将返回 ErrConnClosed 错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/fcgi/child.go;l=187)

``` go 
var ErrRequestAborted = errors.New("fcgi: request aborted by web server")
```

​	当处理程序尝试读取已被 Web 服务器中止的请求的主体时，Read 将返回 ErrRequestAborted 错误。

## 函数

#### func ProcessEnv  <- go1.9

``` go 
func ProcessEnv(r *http.Request) map[string]string
```

​	ProcessEnv 函数返回与请求 `r` 相关的 FastCGI 环境变量。对于这些环境变量，没有在请求本身中进行包含的努力 —— 数据被隐藏在请求的上下文中。例如，如果请求设置了 REMOTE_USER （环境变量），那么在 `r` 的任何地方都找不到它，但它将在 ProcessEnv 的响应中包含（通过 `r` 的上下文）。

#### func Serve 

``` go 
func Serve(l net.Listener, handler http.Handler) error
```

​	Serve 函数在监听器 `l`上接受传入的 FastCGI 连接，为每个连接创建一个新的 goroutine。这些 goroutine 读取请求，然后调用处理程序来回复请求。如果 `l` 为 nil，则 Serve 从 os.Stdin 接受连接。如果 handler 为 nil，则使用 http.DefaultServeMux。

## 类型

This section is empty.