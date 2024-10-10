+++
title = "Encrypt Cookie"
date = 2024-02-05T09:14:15+08:00
weight = 70
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/encryptcookie]({{< ref "/fiber/API/Middleware/EncryptCookie" >}})

# Encrypt Cookie 加密 Cookie

Encrypt Cookie is a middleware for [Fiber](https://github.com/gofiber/fiber) that secures your cookie values through encryption.

​	Encrypt Cookie 是一个中间件，用于通过加密保护您的 cookie 值。

NOTE
注意

This middleware encrypts cookie values and not the cookie names.

​	此中间件加密 cookie 值，而不是 cookie 名称。

## Signatures 签名

```go
// Intitializes the middleware
func New(config ...Config) fiber.Handler

// Returns a random 32 character long string
func GenerateKey() string
```



## Examples 示例 

To use the Encrypt Cookie middleware, first, import the middleware package as part of the Fiber web framework:

​	要使用 Encrypt Cookie 中间件，首先，将中间件包作为 Fiber Web 框架的一部分导入：

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/encryptcookie"
)
```



Once you've imported the middleware package, you can use it inside your Fiber app:

​	导入中间件包后，您可以在 Fiber 应用中使用它：

```go
// Provide a minimal configuration
app.Use(encryptcookie.New(encryptcookie.Config{
    Key: "secret-thirty-2-character-string",
}))

// Retrieve the encrypted cookie value
app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("value=" + c.Cookies("test"))
})

// Create an encrypted cookie
app.Post("/", func(c *fiber.Ctx) error {
    c.Cookie(&fiber.Cookie{
        Name:  "test",
        Value: "SomeThing",
    })
    return nil
})
```



NOTE
注意

`Key` must be a 32 character string. It's used to encrypt the values, so make sure it is random and keep it secret. You can run `openssl rand -base64 32` or call `encryptcookie.GenerateKey()` to create a random key for you. Make sure not to set `Key` to `encryptcookie.GenerateKey()` because that will create a new key every run.

​	 `Key` 必须是一个 32 个字符的字符串。它用于加密值，所以请确保它是随机的并保密。您可以运行 `openssl rand -base64 32` 或调用 `encryptcookie.GenerateKey()` 为您创建随机密钥。请确保不要将 `Key` 设置为 `encryptcookie.GenerateKey()` ，因为这会每次运行都创建一个新密钥。

## Config 配置

| Property 属性 | Type 输入                                           | Description 说明                                             | Default 默认                                        |
| ------------- | --------------------------------------------------- | ------------------------------------------------------------ | --------------------------------------------------- |
| Next 下一步   | `func(*fiber.Ctx) bool`                             | A function to skip this middleware when returned true. 当返回 true 时，用于跳过此中间件的函数。 | `nil`                                               |
| Except        | `[]string`                                          | Array of cookie keys that should not be encrypted. 不应加密的 cookie 密钥数组。 | `[]`                                                |
| Key           | `string`                                            | A base64-encoded unique key to encode & decode cookies. Required. Key length should be 32 characters. 用于编码和解码 cookie 的 base64 编码唯一密钥。必需。密钥长度应为 32 个字符。 | (No default, required field) （无默认值，必填字段） |
| Encryptor     | `func(decryptedString, key string) (string, error)` | A custom function to encrypt cookies. 用于加密 cookie 的自定义函数。 | `EncryptCookie`                                     |
| Decryptor     | `func(encryptedString, key string) (string, error)` | A custom function to decrypt cookies. 用于解密 cookie 的自定义函数。 | `DecryptCookie`                                     |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next:      nil,
    Except:    []string{},
    Key:       "",
    Encryptor: EncryptCookie,
    Decryptor: DecryptCookie,
}
```



## Usage With Other Middlewares That Reads Or Modify Cookies 与读取或修改 Cookie 的其他中间件一起使用

Place the encryptcookie middleware before any other middleware that reads or modifies cookies. For example, if you are using the CSRF middleware, ensure that the encryptcookie middleware is placed before it. Failure to do so may prevent the CSRF middleware from reading the encrypted cookie.

​	将 encryptcookie 中间件放在任何其他读取或修改 cookie 的中间件之前。例如，如果您正在使用 CSRF 中间件，请确保 encryptcookie 中间件放在它之前。否则，可能会阻止 CSRF 中间件读取加密的 cookie。

You may also choose to exclude certain cookies from encryption. For instance, if you are using the CSRF middleware with a frontend framework like Angular, and the framework reads the token from a cookie, you should exclude that cookie from encryption. This can be achieved by adding the cookie name to the Except array in the configuration:

​	您还可以选择将某些 cookie 排除在加密之外。例如，如果您将 CSRF 中间件与 Angular 等前端框架一起使用，并且该框架从 cookie 中读取令牌，则应将该 cookie 排除在加密之外。这可以通过将 cookie 名称添加到配置中的 Except 数组中来实现：

```go
app.Use(encryptcookie.New(encryptcookie.Config{
    Key:    "secret-thirty-2-character-string",
    Except: []string{csrf.ConfigDefault.CookieName}, // exclude CSRF cookie
}))
app.Use(csrf.New(csrf.Config{
    KeyLookup:      "header:" + csrf.HeaderName,
    CookieSameSite: "Lax",
    CookieSecure:   true,
    CookieHTTPOnly: false,
}))
```
