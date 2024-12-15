+++
title = "server_api"
date = 2024-12-15T21:21:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/api](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/api)
>
> 收录该文档时间： `2024-12-15T21:21:34+08:00`

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type HTTPServer 

``` go
type HTTPServer struct {
	// contains filtered or unexported fields
}
```

#### func NewHTTPServer 

``` go
func NewHTTPServer(
	root string,
	watcher chan messaging.WatcherCommand,
	executor contract.Executor,
	status chan chan string) *HTTPServer
```

#### (*HTTPServer) Execute 

``` go
func (self *HTTPServer) Execute(response http.ResponseWriter, request *http.Request)
```

#### (*HTTPServer) Ignore 

``` go
func (self *HTTPServer) Ignore(response http.ResponseWriter, request *http.Request)
```

#### (*HTTPServer) LongPollStatus 

``` go
func (self *HTTPServer) LongPollStatus(response http.ResponseWriter, request *http.Request)
```

#### (*HTTPServer) ReceiveUpdate 

``` go
func (self *HTTPServer) ReceiveUpdate(root string, update *contract.CompleteOutput)
```

#### (*HTTPServer) Reinstate 

``` go
func (self *HTTPServer) Reinstate(response http.ResponseWriter, request *http.Request)
```

#### (*HTTPServer) Results 

``` go
func (self *HTTPServer) Results(response http.ResponseWriter, request *http.Request)
```

#### (*HTTPServer) Status 

``` go
func (self *HTTPServer) Status(response http.ResponseWriter, request *http.Request)
```

#### (*HTTPServer) TogglePause 

``` go
func (self *HTTPServer) TogglePause(response http.ResponseWriter, request *http.Request)
```

#### (*HTTPServer) Watch 

``` go
func (self *HTTPServer) Watch(response http.ResponseWriter, request *http.Request)
```
