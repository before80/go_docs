+++
title = "Keyauth"
date = 2024-02-05T09:14:15+08:00
weight = 160
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/keyauth]({{< ref "/fiber/API/Middleware/Keyauth" >}})

# Keyauth

Key auth middleware provides a key based authentication.

​	密钥认证中间件提供基于密钥的认证。

## Signatures 签名

```go
func New(config ...Config) fiber.Handler
```



## Examples 示例 

```go
package main

import (
    "crypto/sha256"
    "crypto/subtle"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/keyauth"
)

var (
    apiKey = "correct horse battery staple"
)

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
    hashedAPIKey := sha256.Sum256([]byte(apiKey))
    hashedKey := sha256.Sum256([]byte(key))

    if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
        return true, nil
    }
    return false, keyauth.ErrMissingOrMalformedAPIKey
}

func main() {
    app := fiber.New()

    // note that the keyauth middleware needs to be defined before the routes are defined!
    app.Use(keyauth.New(keyauth.Config{
        KeyLookup:  "cookie:access_token",
        Validator:  validateAPIKey,
    }))

        app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Successfully authenticated!")
    })

    app.Listen(":3000")
}
```



**Test:
测试：**

```bash
# No api-key specified -> 400 missing 
curl http://localhost:3000
#> missing or malformed API Key

curl --cookie "access_token=correct horse battery staple" http://localhost:3000
#> Successfully authenticated!

curl --cookie "access_token=Clearly A Wrong Key" http://localhost:3000
#>  missing or malformed API Key
```



For a more detailed example, see also the [`github.com/gofiber/recipes`](https://github.com/gofiber/recipes) repository and specifically the `fiber-envoy-extauthz` repository and the [`keyauth example`](https://github.com/gofiber/recipes/blob/master/fiber-envoy-extauthz/authz/main.go) code.

​	有关更详细的示例，另请参阅 `github.com/gofiber/recipes` 代码库，特别是 `fiber-envoy-extauthz` 代码库和 `keyauth example` 代码。

### Authenticate only certain endpoints 仅对某些端点进行身份验证

If you want to authenticate only certain endpoints, you can use the `Config` of keyauth and apply a filter function (eg. `authFilter`) like so

​	如果您只想对某些端点进行身份验证，可以使用 keyauth 的 `Config` 并应用一个过滤器函数（例如 `authFilter` ），如下所示

```go
package main

import (
    "crypto/sha256"
    "crypto/subtle"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/keyauth"
    "regexp"
    "strings"
)

var (
    apiKey        = "correct horse battery staple"
    protectedURLs = []*regexp.Regexp{
        regexp.MustCompile("^/authenticated$"),
        regexp.MustCompile("^/auth2$"),
    }
)

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
    hashedAPIKey := sha256.Sum256([]byte(apiKey))
    hashedKey := sha256.Sum256([]byte(key))

    if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
        return true, nil
    }
    return false, keyauth.ErrMissingOrMalformedAPIKey
}

func authFilter(c *fiber.Ctx) bool {
    originalURL := strings.ToLower(c.OriginalURL())

    for _, pattern := range protectedURLs {
        if pattern.MatchString(originalURL) {
            return false
        }
    }
    return true
}

func main() {
    app := fiber.New()

    app.Use(keyauth.New(keyauth.Config{
        Next:    authFilter,
        KeyLookup: "cookie:access_token",
        Validator: validateAPIKey,
    }))

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome")
    })
    app.Get("/authenticated", func(c *fiber.Ctx) error {
        return c.SendString("Successfully authenticated!")
    })
    app.Get("/auth2", func(c *fiber.Ctx) error {
        return c.SendString("Successfully authenticated 2!")
    })

    app.Listen(":3000")
}
```



Which results in this

​	这将产生以下结果

```bash
# / does not need to be authenticated
curl http://localhost:3000
#> Welcome

# /authenticated needs to be authenticated
curl --cookie "access_token=correct horse battery staple" http://localhost:3000/authenticated
#> Successfully authenticated!

# /auth2 needs to be authenticated too
curl --cookie "access_token=correct horse battery staple" http://localhost:3000/auth2
#> Successfully authenticated 2!
```



### Specifying middleware in the handler 在处理程序中指定中间件

```go
package main

import (
    "crypto/sha256"
    "crypto/subtle"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/keyauth"
)

const (
  apiKey = "my-super-secret-key"
)

func main() {
    app := fiber.New()

    authMiddleware := keyauth.New(keyauth.Config{
        Validator:  func(c *fiber.Ctx, key string) (bool, error) {
            hashedAPIKey := sha256.Sum256([]byte(apiKey))
            hashedKey := sha256.Sum256([]byte(key))

            if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
                return true, nil
            }
            return false, keyauth.ErrMissingOrMalformedAPIKey
        },
    })

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome")
    })

    app.Get("/allowed",  authMiddleware, func(c *fiber.Ctx) error {
        return c.SendString("Successfully authenticated!")
    })

    app.Listen(":3000")
}
```



Which results in this

​	这将产生以下结果

```bash
# / does not need to be authenticated
curl http://localhost:3000
#> Welcome

# /allowed needs to be authenticated too
curl --header "Authorization: Bearer my-super-secret-key"  http://localhost:3000/allowed
#> Successfully authenticated!
```



## Config 配置

| Property 属性  | Type 输入                                | Description 说明                                             | Default 默认                               |
| -------------- | ---------------------------------------- | ------------------------------------------------------------ | ------------------------------------------ |
| Next 下一步    | `func(*fiber.Ctx) bool`                  | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                                      |
| SuccessHandler | `fiber.Handler`                          | SuccessHandler defines a function which is executed for a valid key. SuccessHandler 定义了一个针对有效密钥执行的函数。 | `nil`                                      |
| ErrorHandler   | `fiber.ErrorHandler`                     | ErrorHandler defines a function which is executed for an invalid key. ErrorHandler 定义了一个针对无效密钥执行的函数。 | `401 Invalid or expired key`               |
| KeyLookup      | `string`                                 | KeyLookup is a string in the form of "`<source>:<name>`" that is used to extract key from the request. KeyLookup 是一个 " `<source>:<name>` " 形式的字符串，用于从请求中提取密钥。 | "header:Authorization"                     |
| AuthScheme     | `string`                                 | AuthScheme to be used in the Authorization header. AuthScheme 将用于 Authorization 头。 | "Bearer"                                   |
| Validator      | `func(*fiber.Ctx, string) (bool, error)` | Validator is a function to validate the key. Validator 是一个用于验证密钥的函数。 | A function for key validation 密钥验证函数 |
| ContextKey     | `interface{}`                            | Context key to store the bearer token from the token into context. 将令牌中的持有者令牌存储到上下文的上下文键。 | "token"                                    |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    SuccessHandler: func(c *fiber.Ctx) error {
        return c.Next()
    },
    ErrorHandler: func(c *fiber.Ctx, err error) error {
        if err == ErrMissingOrMalformedAPIKey {
            return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
        }
        return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired API Key")
    },
    KeyLookup:  "header:" + fiber.HeaderAuthorization,
    AuthScheme: "Bearer",
    ContextKey: "token",
}
```
