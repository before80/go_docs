+++
title = "方法重写"
weight = 130
date = 2023-07-09T21:56:03+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Method Override - 方法重写

https://echo.labstack.com/docs/middleware/method-override

​	Method Override 中间件检查请求中是否存在覆盖的方法，并使用覆盖的方法替代原始方法。

> 信息
>
> ​	出于安全原因，只有 `POST` 方法可以被覆盖。

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
  // Skipper 定义了一个用于跳过中间件的函数。
  Skipper Skipper

  // Getter 是一个从请求中获取覆盖方法的函数。
  // 可选。默认值为 MethodFromHeader(echo.HeaderXHTTPMethodOverride)。
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