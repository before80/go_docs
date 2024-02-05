+++
title = "timeout"
weight = 240
date = 2023-07-09T21:58:27+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Timeout

> 原文：[https://echo.labstack.com/docs/middleware/timeout](https://echo.labstack.com/docs/middleware/timeout)

Timeout middleware is used to timeout at a long running operation within a predefined period.

## Usage

```go
e.Use(middleware.Timeout())
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
  Skipper: Skipper,
  ErrorHandler: func(err error, e echo.Context) error {
      // you can handle your error here, the returning error will be 
      // passed down the middleware chain
      return err
  },
  Timeout: 30*time.Second,
}))
```



## Configuration

```go
// TimeoutConfig defines the config for Timeout middleware.
TimeoutConfig struct {
    // Skipper defines a function to skip middleware.
    Skipper Skipper
    // ErrorHandler defines a function which is executed for a timeout
    // It can be used to define a custom timeout error
    ErrorHandler TimeoutErrorHandlerWithContext
    // Timeout configures a timeout for the middleware, defaults to 0 for no timeout
    Timeout time.Duration
}
```



`TimeoutErrorHandlerWithContext` is responsible for handling the errors when a timeout happens

```go
// TimeoutErrorHandlerWithContext is an error handler that is used 
// with the timeout middleware so we can handle the error 
// as we see fit
TimeoutErrorHandlerWithContext func(error, echo.Context) error
```



### Default Configuration*

```go
DefaultTimeoutConfig = TimeoutConfig{
    Skipper:      DefaultSkipper,
    Timeout:      0,
    ErrorHandler: nil,
}
```