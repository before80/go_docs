+++
title = "Gzip"
weight = 80
date = 2023-07-09T21:54:52+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Gzip

> 原文：[https://echo.labstack.com/docs/middleware/gzip](https://echo.labstack.com/docs/middleware/gzip)

​	Gzip中间件使用gzip压缩方案对HTTP响应进行压缩。

## Usage

```
e.Use(middleware.Gzip())
```

## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
  Level: 5,
}))
```



> 提示
>
> ​	可以传递一个中间件跳过器（skipper）来避免对特定URL进行gzip压缩。

#### Example

```go
e := echo.New()
e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
  Skipper: func(c echo.Context) bool {
    return strings.Contains(c.Path(), "metrics") // 将 "metrics" 更改为您自己的路径
  },
}))
```



## Configuration

```go
GzipConfig struct {
  // Skipper 定义一个用于跳过中间件的函数。
  Skipper Skipper

  // Gzip 压缩级别。
  // 可选。默认值为-1。
  Level int `json:"level"`
}
```



### Default Configuration

```go
DefaultGzipConfig = GzipConfig{
  Skipper: DefaultSkipper,
  Level:   -1,
}
```