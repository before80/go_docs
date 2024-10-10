+++
title = "Body Dump"
weight = 20
date = 2023-07-09T21:53:47+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Body Dump

> 原文：[https://echo.labstack.com/docs/middleware/body-dump](https://echo.labstack.com/docs/middleware/body-dump)

Body dump middleware captures the request and response payload and calls the registered handler. Generally used for debugging/logging purpose. Avoid using it if your request/response payload is huge e.g. file upload/download, but if you still need to, add an exception for your endpoints in the skipper function.

​	请求体转储中间件捕获请求和响应的有效载荷，并调用注册的处理程序。通常用于调试/日志记录的目的。如果您的请求/响应有效载荷很大，例如文件上传/下载，请避免使用它，但如果仍然需要，请在跳过函数中为您的端点添加例外。

## Usage

```go
e := echo.New()
e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
}))
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{}))
```



## Configuration

```go
BodyDumpConfig struct {
  // Skipper 定义一个跳过中间件的函数。
  Skipper Skipper

  // Handler 接收请求和响应的有效载荷。
  // 必需。
  Handler BodyDumpHandler
}
```



### Default Configuration

```go
DefaultBodyDumpConfig = BodyDumpConfig{
  Skipper: DefaultSkipper,
}
```