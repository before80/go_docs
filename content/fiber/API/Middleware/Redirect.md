+++
title = "Redirect"
date = 2024-02-05T09:14:15+08:00
weight = 230
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/redirect]({{< ref "/fiber/API/Middleware/Redirect" >}})

# Redirect 重定向

Redirection middleware for Fiber.

​	Fiber 的重定向中间件。

## Signatures 签名

```go
func New(config ...Config) fiber.Handler
```



## Examples 示例 

```go
package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/redirect"
)

func main() {
  app := fiber.New()
  
  app.Use(redirect.New(redirect.Config{
    Rules: map[string]string{
      "/old":   "/new",
      "/old/*": "/new/$1",
    },
    StatusCode: 301,
  }))
  
  app.Get("/new", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })
  app.Get("/new/*", func(c *fiber.Ctx) error {
    return c.SendString("Wildcard: " + c.Params("*"))
  })
  
  app.Listen(":3000")
}
```



**Test:
测试：**

```curl
curl http://localhost:3000/old
curl http://localhost:3000/old/hello
```



## Config 配置

| Property 属性 | Type 输入               | Description 说明                                             | Default 默认                          |
| ------------- | ----------------------- | ------------------------------------------------------------ | ------------------------------------- |
| Next 下一步   | `func(*fiber.Ctx) bool` | Filter defines a function to skip middleware. Filter 定义了一个跳过中间件的函数。 | `nil`                                 |
| Rules 规则    | `map[string]string`     | Rules defines the URL path rewrite rules. The values captured in asterisk can be retrieved by index e.g. $1, $2 and so on. 规则定义了 URL 路径重写规则。星号中捕获的值可以通过索引检索，例如 $1、$2 等。 | Required 必需                         |
| StatusCode    | `int`                   | The status code when redirecting. This is ignored if Redirect is disabled. 重定向时的状态代码。如果禁用了重定向，则忽略此项。 | 302 Temporary Redirect 302 临时重定向 |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    StatusCode: fiber.StatusFound,
}
```
