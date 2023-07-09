+++
title = "method-override"
date = 2023-07-09T21:56:03+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Method Override

https://echo.labstack.com/docs/middleware/method-override

Method override middleware checks for the overridden method from the request and uses it instead of the original method.

INFO

For security reasons, only `POST` method can be overridden.

## Usage

```go
e.Pre(middleware.MethodOverride())
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
  Getter: middleware.MethodFromForm("_method"),
}))
```



## Configuration

```go
MethodOverrideConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Getter is a function that gets overridden method from the request.
  // Optional. Default values MethodFromHeader(echo.HeaderXHTTPMethodOverride).
  Getter MethodOverrideGetter
}
```



### Default Configuration

```go
DefaultMethodOverrideConfig = MethodOverrideConfig{
  Skipper: DefaultSkipper,
  Getter:  MethodFromHeader(echo.HeaderXHTTPMethodOverride),
}
```