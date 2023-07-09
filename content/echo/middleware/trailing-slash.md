+++
title = "trailing-slash"
date = 2023-07-09T21:58:43+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Trailing Slash

https://echo.labstack.com/docs/middleware/trailing-slash

## Add Trailing Slash

Add trailing slash middleware adds a trailing slash to the request URI.

### Usage

```go
e := echo.New()
e.Pre(middleware.AddTrailingSlash())
```



## Remove Trailing Slash

Remove trailing slash middleware removes a trailing slash from the request URI.

### Usage

```go
e := echo.New()
e.Pre(middleware.RemoveTrailingSlash())
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
  RedirectCode: http.StatusMovedPermanently,
}))
```



Example above will add a trailing slash to the request URI and redirect with `301 - StatusMovedPermanently`.

## xxxxxxxxxx5 1DefaultTimeoutConfig = TimeoutConfig{2    Skipper:      DefaultSkipper,3    Timeout:      0,4    ErrorHandler: nil,5}go

```go
TrailingSlashConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Status code to be used when redirecting the request.
  // Optional, but when provided the request is redirected using this code.
  RedirectCode int `json:"redirect_code"`
}
```



### Default Configuration

```go
DefaultTrailingSlashConfig = TrailingSlashConfig{
  Skipper: DefaultSkipper,
}
```