+++
title = "recover"
weight = 170
date = 2023-07-09T21:57:18+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Recover

https://echo.labstack.com/docs/middleware/recover

Recover middleware recovers from panics anywhere in the chain, prints stack trace and handles the control to the centralized [HTTPErrorHandler](https://echo.labstack.com/docs/customization#http-error-handler).

## Usage

```go
e.Use(middleware.Recover())
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
  StackSize: 1 << 10, // 1 KB
  LogLevel:  log.ERROR,
}))
```



Example above uses a `StackSize` of 1 KB, `LogLevel` of error and default values for `DisableStackAll` and `DisablePrintStack`.

## Configuration

```go
// LogErrorFunc defines a function for custom logging in the middleware.
LogErrorFunc func(c echo.Context, err error, stack []byte) error

RecoverConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Size of the stack to be printed.
  // Optional. Default value 4KB.
  StackSize int `yaml:"stack_size"`

  // DisableStackAll disables formatting stack traces of all other goroutines
  // into buffer after the trace for the current goroutine.
  // Optional. Default value false.
  DisableStackAll bool `yaml:"disable_stack_all"`

  // DisablePrintStack disables printing stack trace.
  // Optional. Default value as false.
  DisablePrintStack bool `yaml:"disable_print_stack"`

  // LogLevel is log level to printing stack trace.
  // Optional. Default value 0 (Print).
  LogLevel log.Lvl

  // LogErrorFunc defines a function for custom logging in the middleware.
  // If it's set you don't need to provide LogLevel for config.
  LogErrorFunc LogErrorFunc
}
```



### Default Configuration

```go
DefaultRecoverConfig = RecoverConfig{
  Skipper:           DefaultSkipper,
  StackSize:         4 << 10, // 4 KB
  DisableStackAll:   false,
  DisablePrintStack: false,
  LogLevel:          0,
  LogErrorFunc:      nil,
}
```