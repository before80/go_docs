+++
title = "请求体限制"
weight = 30
date = 2023-07-09T21:54:03+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# 请求体限制

> 原文：[https://echo.labstack.com/docs/middleware/body-limit](https://echo.labstack.com/docs/middleware/body-limit)

​	请求体限制中间件设置请求体的最大允许大小，如果大小超过配置的限制，则发送 "413 - Request Entity Too Large" 响应。请求体限制基于请求头中的 `Content-Length` 和实际读取的内容来确定，这使得它非常安全。

​	限制可以指定为 `4x` 或 `4xB`，其中 x 是从 K、M、G、T 或 P 中选择的倍数之一。

## Usage

```go
e := echo.New()
e.Use(middleware.BodyLimit("2M"))
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{}))
```



## Configuration

```go
BodyLimitConfig struct {
  // Skipper 定义一个用于跳过中间件的函数。
  Skipper Skipper

  // 请求体的最大允许大小，可以指定为 `4x` 或 `4xB`，
  // 其中 x 是从 K、M、G、T 或 P 中选择的倍数之一。 
  Limit string `json:"limit"`
}
```



### Default Configuration

```go
DefaultBodyLimitConfig = BodyLimitConfig{
  Skipper: DefaultSkipper,
}
```