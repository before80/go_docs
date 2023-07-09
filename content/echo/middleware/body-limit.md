+++
title = "body-limit"
date = 2023-07-09T21:54:03+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Body Limit

https://echo.labstack.com/docs/middleware/body-limit

Body limit middleware sets the maximum allowed size for a request body, if the size exceeds the configured limit, it sends "413 - Request Entity Too Large" response. The body limit is determined based on both `Content-Length` request header and actual content read, which makes it super secure.

Limit can be specified as `4x` or `4xB`, where x is one of the multiple from K, M, G, T or P.

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
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Maximum allowed size for a request body, it can be specified
  // as `4x` or `4xB`, where x is one of the multiple from K, M, G, T or P.
  Limit string `json:"limit"`
}
```



### Default Configuration

```go
DefaultBodyLimitConfig = BodyLimitConfig{
  Skipper: DefaultSkipper,
}
```