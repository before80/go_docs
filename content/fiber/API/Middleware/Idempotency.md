+++
title = "Idempotency"
date = 2024-02-05T09:14:15+08:00
weight = 150
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/idempotency]({{< ref "/fiber/API/Middleware/Idempotency" >}})

# Idempotency 幂等性

Idempotency middleware for [Fiber](https://github.com/gofiber/fiber) allows for fault-tolerant APIs where duplicate requests — for example due to networking issues on the client-side — do not erroneously cause the same action performed multiple times on the server-side.

​	Fiber 的幂等性中间件允许容错 API，其中重复请求（例如由于客户端的网络问题）不会错误地导致在服务器端多次执行相同的操作。

Refer to https://datatracker.ietf.org/doc/html/draft-ietf-httpapi-idempotency-key-header-02 for a better understanding.

​	有关更深入的了解，请参阅 https://datatracker.ietf.org/doc/html/draft-ietf-httpapi-idempotency-key-header-02。

## Signatures 签名

```go
func New(config ...Config) fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/idempotency"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

### Default Config 默认配置 

```go
app.Use(idempotency.New())
```



### Custom Config 自定义配置 

```go
app.Use(idempotency.New(idempotency.Config{
    Lifetime: 42 * time.Minute,
    // ...
}))
```



### Config 配置

| Property 属性       | Type 输入               | Description 说明                                             | Default 默认                                        |
| ------------------- | ----------------------- | ------------------------------------------------------------ | --------------------------------------------------- |
| Next 下一步         | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | A function for safe methods 安全方法的函数          |
| Lifetime 生存期     | `time.Duration`         | Lifetime is the maximum lifetime of an idempotency key. 生存期是幂等键的最大生存期。 | 30 * time.Minute                                    |
| KeyHeader           | `string`                | KeyHeader is the name of the header that contains the idempotency key. KeyHeader 是包含幂等键的标头的名称。 | "X-Idempotency-Key"                                 |
| KeyHeaderValidate   | `func(string) error`    | KeyHeaderValidate defines a function to validate the syntax of the idempotency header. KeyHeaderValidate 定义一个函数来验证幂等性标头的语法。 | A function for UUID validation 用于 UUID 验证的函数 |
| KeepResponseHeaders | `[]string`              | KeepResponseHeaders is a list of headers that should be kept from the original response. KeepResponseHeaders 是应从原始响应中保留的标头列表。 | nil (keep all headers) nil（保留所有标头）          |
| Lock                | `Locker`                | Lock locks an idempotency key. Lock 锁定幂等性密钥。         | An in-memory locker 内存中储物柜                    |
| Storage             | `fiber.Storage`         | Storage stores response data by idempotency key. Storage 按幂等性密钥存储响应数据。 | An in-memory storage 内存中存储                     |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next: func(c *fiber.Ctx) bool {
        // Skip middleware if the request was done using a safe HTTP method
        return fiber.IsMethodSafe(c.Method())
    },

    Lifetime: 30 * time.Minute,

    KeyHeader: "X-Idempotency-Key",
    KeyHeaderValidate: func(k string) error {
        if l, wl := len(k), 36; l != wl { // UUID length is 36 chars
            return fmt.Errorf("%w: invalid length: %d != %d", ErrInvalidIdempotencyKey, l, wl)
        }

        return nil
    },

    KeepResponseHeaders: nil,

    Lock: nil, // Set in configDefault so we don't allocate data here.

    Storage: nil, // Set in configDefault so we don't allocate data here.
}
```
