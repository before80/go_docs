+++
title = "BasicAuth"
date = 2024-02-05T09:14:15+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/basicauth]({{< ref "/fiber/API/Middleware/BasicAuth" >}})

# BasicAuth net/http 到 Fiber - net/http 中间件到 Fiber

Basic Authentication middleware for [Fiber](https://github.com/gofiber/fiber) that provides an HTTP basic authentication. It calls the next handler for valid credentials and [401 Unauthorized](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401) or a custom response for missing or invalid credentials.

​	Fiber 的基本身份验证中间件，提供 HTTP 基本身份验证。它为有效凭据调用下一个处理程序，并为缺失或无效凭据提供 401 未授权或自定义响应。

## Signatures 签名

```go
func New(config Config) fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/basicauth"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Provide a minimal config
app.Use(basicauth.New(basicauth.Config{
    Users: map[string]string{
        "john":  "doe",
        "admin": "123456",
    },
}))

// Or extend your config for customization
app.Use(basicauth.New(basicauth.Config{
    Users: map[string]string{
        "john":  "doe",
        "admin": "123456",
    },
    Realm: "Forbidden",
    Authorizer: func(user, pass string) bool {
        if user == "john" && pass == "doe" {
            return true
        }
        if user == "admin" && pass == "123456" {
            return true
        }
        return false
    },
    Unauthorized: func(c *fiber.Ctx) error {
        return c.SendFile("./unauthorized.html")
    },
    ContextUsername: "_user",
    ContextPassword: "_pass",
}))
```



## Config 配置

| Property 属性       | Type 输入                   | Description 说明                                             | Default 默认          |
| ------------------- | --------------------------- | ------------------------------------------------------------ | --------------------- |
| Next 下一步         | `func(*fiber.Ctx) bool`     | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                 |
| Users 用户          | `map[string]string`         | Users defines the allowed credentials. 用户定义允许的凭据。  | `map[string]string{}` |
| Realm 领域          | `string`                    | Realm is a string to define the realm attribute of BasicAuth. The realm identifies the system to authenticate against and can be used by clients to save credentials. 领域是一个字符串，用于定义 BasicAuth 的领域属性。领域标识要进行身份验证的系统，客户端可以使用它来保存凭据。 | `"Restricted"`        |
| Authorizer 授权者   | `func(string, string) bool` | Authorizer defines a function to check the credentials. It will be called with a username and password and is expected to return true or false to indicate approval. 授权者定义一个函数来检查凭据。它将使用用户名和密码调用，并希望返回 true 或 false 来指示批准。 | `nil`                 |
| Unauthorized 未授权 | `fiber.Handler`             | Unauthorized defines the response body for unauthorized responses. Unauthorized 定义未授权响应的响应主体。 | `nil`                 |
| ContextUsername     | `interface{}`               | ContextUsername is the key to store the username in Locals. ContextUsername 是将用户名存储在 Locals 中的键。 | `"username"`          |
| ContextPassword     | `interface{}`               | ContextPassword is the key to store the password in Locals. ContextPassword 是将密码存储在 Locals 中的键。 | `"password"`          |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next:            nil,
    Users:           map[string]string{},
    Realm:           "Restricted",
    Authorizer:      nil,
    Unauthorized:    nil,
    ContextUsername: "username",
    ContextPassword: "password",
}
```
