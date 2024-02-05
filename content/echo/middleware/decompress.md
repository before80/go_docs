+++
title = "解压缩"
weight = 70
date = 2023-07-09T21:54:44+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Decompress - 解压缩

> 原文：[https://echo.labstack.com/docs/middleware/decompress](https://echo.labstack.com/docs/middleware/decompress)

​	解压缩中间件会在Content-Encoding标头设置为gzip时对HTTP请求进行解压缩。

> 注意
>
> ​	请求体将在内存中被解压缩，并在请求的整个生命周期中占用内存空间（包括垃圾回收）。

## Usage

```go
e.Use(middleware.Decompress())
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.DecompressWithConfig(middleware.DecompressConfig{
  Skipper: Skipper
}))
```



## Configuration

```go
DecompressConfig struct {
  // Skipper定义一个用于跳过中间件的函数。
  Skipper Skipper
}
```



### Default Configuration

```go
DefaultDecompressConfig = DecompressConfig{
  Skipper: DefaultSkipper,
}
```