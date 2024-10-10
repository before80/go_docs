+++
title = "Key Auth"
weight = 110
date = 2023-07-09T21:55:25+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Key Auth

> 原文：[https://echo.labstack.com/docs/middleware/key-auth](https://echo.labstack.com/docs/middleware/key-auth)

​	Key Auth 中间件提供基于密钥的身份验证。 

- 对于有效的密钥，调用下一个处理程序。
- 对于无效的密钥，发送 "401 - Unauthorized" 响应。
- 对于缺失的密钥，发送 "400 - Bad Request" 响应。

## Usage

```go
e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
  return key == "valid-key", nil
}))
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
  KeyLookup: "query:api-key",
  Validator: func(key string, c echo.Context) (bool, error) {
            return key == "valid-key", nil
        },
}))
```



## Configuration

```go
KeyAuthConfig struct {
  // Skipper 定义了一个用于跳过中间件的函数。
  Skipper Skipper

  // KeyLookup 是一个字符串，格式为 "<source>:<name>"，用于从请求中提取密钥。
  // 可选。默认值为 "header:Authorization"。
  // 可能的值：
  // - "header:<name>"
  // - "query:<name>"
  // - "cookie:<name>"
  // - "form:<name>"
  KeyLookup string `yaml:"key_lookup"`

  // AuthScheme 用于 Authorization 标头中的身份验证方案。
  // 可选。默认值为 "Bearer"。
  AuthScheme string

  // Validator 是一个验证密钥的函数。
  // 必需。 
  Validator KeyAuthValidator

  // ErrorHandler 定义了一个在密钥无效时执行的函数。
  // 可用于定义自定义错误。
  ErrorHandler KeyAuthErrorHandler
}
```



### Default Configuration

```go
DefaultKeyAuthConfig = KeyAuthConfig{
  Skipper:    DefaultSkipper,
  KeyLookup:  "header:" + echo.HeaderAuthorization,
  AuthScheme: "Bearer",
}
```