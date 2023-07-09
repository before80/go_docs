+++
title = "decompress"
date = 2023-07-09T21:54:44+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Decompress

https://echo.labstack.com/docs/middleware/decompress

Decompress middleware decompresses HTTP request if Content-Encoding header is set to gzip.

NOTE

The body will be decompressed in memory and consume it for the lifetime of the request (and garbage collection).

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
  // Skipper defines a function to skip middleware.
  Skipper Skipper
}
```



### Default Configuration

```go
DefaultDecompressConfig = DecompressConfig{
  Skipper: DefaultSkipper,
}
```