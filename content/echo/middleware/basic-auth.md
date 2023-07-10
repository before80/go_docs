+++
title = "基本认证"
weight = 10
date = 2023-07-09T21:53:36+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Basic Auth - 基本认证

https://echo.labstack.com/docs/middleware/basic-auth

Basic auth middleware provides an HTTP basic authentication.

基本认证中间件提供了HTTP基本认证。 

- For valid credentials it calls the next handler.
- For missing or invalid credentials, it sends "401 - Unauthorized" response.
- 对于有效的凭证，它调用下一个处理程序。
- 对于缺失或无效的凭证，它发送"401 - 未经授权"的响应。

## 使用方法

```go
e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
    // Be careful to use constant time comparison to prevent timing attacks
    // 请注意使用恒定时间比较来防止时序攻击
    if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
        subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
        return true, nil
    }
    return false, nil
}))
```



## 自定义配置

### 使用方法

```go
e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}))
```



## 配置

```go
BasicAuthConfig struct {
  // Skipper 定义一个跳过中间件的函数。
  Skipper Skipper

  // Validator 是一个用于验证 BasicAuth 凭证的函数。
  // 必需。
  Validator BasicAuthValidator

  // Realm 是定义 BasicAuth 的域(realm)属性的字符串。
  // 默认值为"Restricted"。
  Realm string
}
```



### 默认配置

```go
DefaultBasicAuthConfig = BasicAuthConfig{
    Skipper: DefaultSkipper,
}
```



